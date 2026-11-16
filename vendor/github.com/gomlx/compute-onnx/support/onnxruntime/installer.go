// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package onnxruntime

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

const (
	DefaultVersion = ""
	RepoURL        = "https://github.com/microsoft/onnxruntime/releases/download"
)

// GetInstallPath returns the default directory where the library will be installed.
func GetInstallPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get user home directory")
	}
	switch runtime.GOOS {
	case "linux":
		return filepath.Join(homeDir, ".local", "lib", "onnxruntime"), nil
	case "darwin":
		return filepath.Join(homeDir, "Library", "Application Support", "onnxruntime"), nil
	case "windows":
		return filepath.Join(homeDir, "AppData", "Local", "onnxruntime"), nil
	default:
		return "", errors.Errorf("platform %s/%s not supported", runtime.GOOS, runtime.GOARCH)
	}
}

// GetLibFilename returns the expected library name for the platform.
func GetLibFilename() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return "libonnxruntime.so", nil
	case "darwin":
		return "libonnxruntime.dylib", nil
	case "windows":
		return "onnxruntime.dll", nil
	default:
		return "", errors.Errorf("platform %s not supported", runtime.GOOS)
	}
}

// GetInstalledVersion returns the normalized version string of the ONNX Runtime library
// currently installed in targetDir (e.g., "1.30.0"). If targetDir is empty, it checks
// the default install directory.
// If no library is installed, it returns "" and nil.
// If the library is present but its version cannot be determined, it returns "" and nil.
func GetInstalledVersion(targetDir string) (string, error) {
	installDir := targetDir
	if installDir == "" {
		var err error
		installDir, err = GetInstallPath()
		if err != nil {
			return "", err
		}
	} else {
		installDir = filepath.Clean(installDir)
	}

	libFilename, err := GetLibFilename()
	if err != nil {
		return "", err
	}

	targetPath := filepath.Join(installDir, libFilename)
	if _, err := os.Stat(targetPath); err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", errors.Wrapf(err, "failed to stat library at %s", targetPath)
	}

	// 1. Check VERSION or .version file.
	for _, vName := range []string{"VERSION", ".version"} {
		vPath := filepath.Join(installDir, vName)
		if data, err := os.ReadFile(vPath); err == nil {
			ver := strings.TrimSpace(string(data))
			if ver != "" {
				return NormalizeVersion(ver), nil
			}
		}
	}

	// 2. Check libonnxruntime.pc pkg-config file.
	for _, pcName := range []string{"libonnxruntime.pc", filepath.Join("pkgconfig", "libonnxruntime.pc")} {
		pcPath := filepath.Join(installDir, pcName)
		if data, err := os.ReadFile(pcPath); err == nil {
			if ver := parseVersionFromPC(string(data)); ver != "" {
				return NormalizeVersion(ver), nil
			}
		}
	}

	// 3. Check resolved symlink of the library file.
	if resolved, err := filepath.EvalSymlinks(targetPath); err == nil {
		base := filepath.Base(resolved)
		if ver := parseVersionFromLibName(base); ver != "" {
			return NormalizeVersion(ver), nil
		}
	}

	return "", nil
}

func parseVersionFromPC(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Version:") {
			ver := strings.TrimSpace(strings.TrimPrefix(line, "Version:"))
			if ver != "" {
				return ver
			}
		}
	}
	return ""
}

func parseVersionFromLibName(base string) string {
	// Linux: libonnxruntime.so.<version> (e.g. libonnxruntime.so.1.30.0)
	if strings.HasPrefix(base, "libonnxruntime.so.") {
		ver := strings.TrimPrefix(base, "libonnxruntime.so.")
		parts := strings.Split(ver, ".")
		if len(parts) >= 2 {
			return ver
		}
	}
	// macOS: libonnxruntime.<version>.dylib (e.g. libonnxruntime.1.20.0.dylib)
	if strings.HasPrefix(base, "libonnxruntime.") && strings.HasSuffix(base, ".dylib") {
		ver := strings.TrimPrefix(base, "libonnxruntime.")
		ver = strings.TrimSuffix(ver, ".dylib")
		parts := strings.Split(ver, ".")
		if len(parts) >= 2 {
			return ver
		}
	}
	return ""
}

// GetDefaultCUDAVersion attempts to guess the CUDA version based on installed libraries found on standard paths,
// or via nvcc if available. It returns the major version as a string (e.g., "12", "13"). If nothing is found, it defaults to "12".
func GetDefaultCUDAVersion() string {
	if fullVer := GetCUDAFullVersion(); fullVer != "" {
		if major := cleanCudaVersion(fullVer); major != "" {
			return major
		}
	}
	return "12"
}

func normalizeCudaVersionString(v string) string {
	v = strings.ToLower(strings.TrimSpace(v))
	v = strings.TrimPrefix(v, "cuda")
	v = strings.TrimPrefix(v, "v")
	v = strings.TrimLeft(v, "-_ ")
	v = strings.ReplaceAll(v, "-", ".")
	v = strings.ReplaceAll(v, "_", ".")
	return v
}

func cleanCudaVersion(v string) string {
	v = normalizeCudaVersionString(v)
	parts := strings.Split(v, ".")
	if len(parts) > 0 {
		return parts[0]
	}
	return v
}

func parseCUDAVersionComponents(v string) (major, minor, patch, specificity int, ok bool) {
	v = normalizeCudaVersionString(v)
	if v == "" {
		return 0, 0, 0, 0, false
	}
	parts := strings.Split(v, ".")
	if len(parts) == 0 || parts[0] == "" {
		return 0, 0, 0, 0, false
	}
	maj, err := strconv.Atoi(parts[0])
	if err != nil || maj <= 0 || maj > 100 {
		return 0, 0, 0, 0, false
	}
	major = maj
	specificity = 1
	if len(parts) > 1 && parts[1] != "" {
		if min, err := strconv.Atoi(parts[1]); err == nil && min >= 0 {
			minor = min
			specificity = 2
			if len(parts) > 2 && parts[2] != "" {
				if pat, err := strconv.Atoi(parts[2]); err == nil && pat >= 0 {
					patch = pat
					specificity = 3
				}
			}
		}
	}
	return major, minor, patch, specificity, true
}

func compareCUDAVersions(a, b string) int {
	majA, minA, patA, specA, okA := parseCUDAVersionComponents(a)
	majB, minB, patB, specB, okB := parseCUDAVersionComponents(b)
	if !okA && !okB {
		return 0
	}
	if !okA {
		return -1
	}
	if !okB {
		return 1
	}
	if majA != majB {
		if majA > majB {
			return 1
		}
		return -1
	}
	if minA != minB {
		if minA > minB {
			return 1
		}
		return -1
	}
	if patA != patB {
		if patA > patB {
			return 1
		}
		return -1
	}
	if specA != specB {
		if specA > specB {
			return 1
		}
		return -1
	}
	return 0
}

func parseCudaVersionFromOutput(s string) string {
	return cleanCudaVersion(parseCUDAFullVersionFromOutput(s))
}

type githubRelease struct {
	TagName string `json:"tag_name"`
}

var knownLatestPatch = map[string]string{
	"1.27": "1.27.1",
	"1.24": "1.24.4",
	"1.29": "1.29.1",
	"1.30": "1.30.0",
}

// ResolveLatestPatchVersion queries GitHub releases to find the highest patch version for a given major.minor string (e.g. "1.27" -> "1.27.1").
func ResolveLatestPatchVersion(majorMinor string) string {
	fallback := func() string {
		if known, ok := knownLatestPatch[majorMinor]; ok {
			return known
		}
		return majorMinor + ".0"
	}

	resp, err := http.Get("https://api.github.com/repos/microsoft/onnxruntime/releases?per_page=100")
	if err != nil {
		return fallback()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fallback()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fallback()
	}

	var releases []githubRelease
	if err := json.Unmarshal(body, &releases); err != nil {
		return fallback()
	}

	prefix := "v" + majorMinor + "."
	highestPatch := -1

	for _, rel := range releases {
		tag := strings.TrimSpace(rel.TagName)
		if strings.HasPrefix(tag, prefix) {
			patchStr := strings.TrimPrefix(tag, prefix)
			patchParts := strings.Split(patchStr, "-")
			if patchVal, err := strconv.Atoi(patchParts[0]); err == nil {
				if patchVal > highestPatch {
					highestPatch = patchVal
				}
			}
		}
	}

	if highestPatch >= 0 {
		return fmt.Sprintf("%s.%d", majorMinor, highestPatch)
	}
	return fallback()
}

// NormalizeVersion cleans up user-provided version strings (e.g., "v1.29.0", "v1.29", "1.29") into standard semver "1.29.0".
// If a 2-part version is provided (e.g., "1.27" or "v1.27"), it queries GitHub releases to find the latest patch version (e.g., "1.27.1").
func NormalizeVersion(v string) string {
	v = strings.TrimSpace(v)
	v = strings.TrimPrefix(v, "v")
	v = strings.TrimPrefix(v, "V")
	if v == "" {
		return ""
	}
	parts := strings.Split(v, ".")
	if len(parts) == 1 {
		return ResolveLatestPatchVersion(v + ".0")
	}
	if len(parts) == 2 {
		return ResolveLatestPatchVersion(v)
	}
	return v
}

// GetAssetURL returns the download URL for the given version, platform, and cuda setting.
func GetAssetURL(version string, cuda bool, cudaVersion string) (string, error) {
	version = NormalizeVersion(version)
	var platform string
	var extension string

	switch runtime.GOOS {
	case "linux":
		extension = ".tgz"
		switch runtime.GOARCH {
		case "amd64":
			platform = "linux-x64"
		case "arm64":
			platform = "linux-aarch64"
		default:
			return "", errors.Errorf("unsupported linux arch: %s", runtime.GOARCH)
		}
	case "darwin":
		extension = ".tgz"
		switch runtime.GOARCH {
		case "arm64":
			platform = "osx-arm64"
		case "amd64":
			platform = "osx-x64"
		default:
			return "", errors.Errorf("unsupported osx arch: %s", runtime.GOARCH)
		}
	case "windows":
		extension = ".zip"
		switch runtime.GOARCH {
		case "amd64":
			platform = "win-x64"
		case "arm64":
			platform = "win-arm64"
		default:
			return "", errors.Errorf("unsupported windows arch: %s", runtime.GOARCH)
		}
	default:
		return "", errors.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	if cuda {
		if runtime.GOARCH != "amd64" {
			return "", errors.Errorf("CUDA is only supported on amd64/x64 architecture (got %s)", runtime.GOARCH)
		}
		if runtime.GOOS != "linux" && runtime.GOOS != "windows" {
			return "", errors.Errorf("CUDA is only supported on linux or windows OS (got %s)", runtime.GOOS)
		}

		if cudaVersion == "" {
			cudaVersion = GetDefaultCUDAVersion()
		}
		majorVer := cleanCudaVersion(cudaVersion)
		if majorVer == "" {
			majorVer = "12"
		}

		// List of candidate suffixes for CUDA flavors to try
		candidates := []string{
			fmt.Sprintf("gpu_cuda%s", majorVer),
			"gpu",
		}
		if majorVer != "12" {
			candidates = append(candidates, "gpu_cuda12")
		}
		if majorVer != "13" {
			candidates = append(candidates, "gpu_cuda13")
		}

		var attempted []string
		for _, suffix := range candidates {
			candidatePlatform := fmt.Sprintf("%s-%s", platform, suffix)
			archiveName := fmt.Sprintf("onnxruntime-%s-%s%s", candidatePlatform, version, extension)
			url := fmt.Sprintf("%s/v%s/%s", RepoURL, version, archiveName)

			// Check if URL exists
			resp, err := http.Head(url)
			if err == nil {
				resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
					return url, nil
				}
				attempted = append(attempted, fmt.Sprintf("  - %s (HTTP %s)", url, resp.Status))
			} else {
				attempted = append(attempted, fmt.Sprintf("  - %s (Error: %v)", url, err))
			}
		}

		return "", errors.Errorf(
			"failed to find a valid CUDA package URL for version %s on platform %s with CUDA version %s.\nAttempted URLs:\n%s",
			version, platform, cudaVersion, strings.Join(attempted, "\n"),
		)
	}

	archiveName := fmt.Sprintf("onnxruntime-%s-%s%s", platform, version, extension)
	url := fmt.Sprintf("%s/v%s/%s", RepoURL, version, archiveName)
	resp, err := http.Head(url)
	if err == nil {
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return "", errors.Errorf(
				"failed to find a valid package URL for version %s on platform %s.\nAttempted URL:\n  - %s (HTTP %s)",
				version, platform, url, resp.Status,
			)
		}
	} else {
		return "", errors.Wrapf(err, "failed to check package URL %s", url)
	}
	return url, nil
}

func cudaVersionFromLibFile(filePath string) string {
	bestVer := ""
	updateBest := func(candidate string) {
		if _, _, _, _, ok := parseCUDAVersionComponents(candidate); ok {
			if compareCUDAVersions(candidate, bestVer) > 0 {
				bestVer = candidate
			}
		}
	}

	checkBase := func(base string) {
		if parts := strings.Split(base, "libcudart.so."); len(parts) == 2 && parts[1] != "" {
			updateBest(parts[1])
		} else if strings.HasPrefix(base, "cudart64_") && strings.HasSuffix(base, ".dll") {
			ver := strings.TrimPrefix(base, "cudart64_")
			ver = strings.TrimSuffix(ver, ".dll")
			if len(ver) >= 2 {
				updateBest(ver[:2])
			}
		}
	}

	checkBase(filepath.Base(filePath))
	if resolved, err := filepath.EvalSymlinks(filePath); err == nil {
		checkBase(filepath.Base(resolved))
	}
	return bestVer
}

func cudaVersionFromLibDir(dir string) string {
	if dir == "" {
		return ""
	}
	bestVer := ""
	updateBest := func(candidate string) {
		if candidate != "" && compareCUDAVersions(candidate, bestVer) > 0 {
			bestVer = candidate
		}
	}

	if matches, err := filepath.Glob(filepath.Join(dir, "libcudart.so.*")); err == nil {
		for _, m := range matches {
			updateBest(cudaVersionFromLibFile(m))
		}
	}
	if matches, err := filepath.Glob(filepath.Join(dir, "cudart64_*.dll")); err == nil {
		for _, m := range matches {
			updateBest(cudaVersionFromLibFile(m))
		}
	}
	return bestVer
}

func cudaVersionFromInstallDir(cudaDir string) string {
	if cudaDir == "" {
		return ""
	}
	if _, err := os.Stat(cudaDir); err != nil {
		return ""
	}

	bestVer := ""
	updateBest := func(candidate string) {
		if candidate != "" && compareCUDAVersions(candidate, bestVer) > 0 {
			bestVer = candidate
		}
	}

	// 1. Check runtime libraries inside the CUDA directory.
	libSubdirs := []string{
		"lib64",
		filepath.Join("targets", "x86_64-linux", "lib"),
		filepath.Join("targets", "sbsa-linux", "lib"),
		filepath.Join("targets", "aarch64-linux", "lib"),
		"lib",
		"bin",
		".",
	}
	for _, sub := range libSubdirs {
		updateBest(cudaVersionFromLibDir(filepath.Join(cudaDir, sub)))
	}

	// 2. Check nvcc inside bin/.
	for _, nvccName := range []string{"nvcc", "nvcc.exe"} {
		nvccPath := filepath.Join(cudaDir, "bin", nvccName)
		if info, err := os.Stat(nvccPath); err == nil && !info.IsDir() {
			if out, err := exec.Command(nvccPath, "--version").Output(); err == nil {
				updateBest(parseCUDAFullVersionFromOutput(string(out)))
			}
		}
	}

	// 3. Check directory name and resolved symlink target.
	updateBest(parseCUDAFullVersionFromOutput(cudaDir))
	if resolved, err := filepath.EvalSymlinks(cudaDir); err == nil {
		updateBest(parseCUDAFullVersionFromOutput(resolved))
	}

	return bestVer
}

// GetCUDAFullVersion attempts to return the full CUDA version string (e.g. "13.3.29", "12.4", or "12.4.131") found on the system.
func GetCUDAFullVersion() string {
	// 1. Check explicit CUDA environment variables first (highest priority user override).
	for _, envVar := range []string{"CUDA_PATH", "CUDA_HOME", "CUDA_ROOT"} {
		if val := os.Getenv(envVar); val != "" {
			if ver := cudaVersionFromInstallDir(val); ver != "" {
				return ver
			}
			if ver := parseCUDAFullVersionFromOutput(val); ver != "" {
				return ver
			}
		}
	}

	// 2. Check LD_LIBRARY_PATH if explicitly set.
	if ldPath := os.Getenv("LD_LIBRARY_PATH"); ldPath != "" {
		bestLDVer := ""
		for _, p := range filepath.SplitList(ldPath) {
			if ver := cudaVersionFromLibDir(p); ver != "" && compareCUDAVersions(ver, bestLDVer) > 0 {
				bestLDVer = ver
			}
		}
		if bestLDVer != "" {
			return bestLDVer
		}
	}

	// 3. Discover all system CUDA installations and select the highest installed version.
	bestVer := ""
	updateBest := func(candidate string) {
		if candidate != "" && compareCUDAVersions(candidate, bestVer) > 0 {
			bestVer = candidate
		}
	}

	// 3a. Standard /usr/local/cuda and /usr/local/cuda-* installations.
	updateBest(cudaVersionFromInstallDir("/usr/local/cuda"))
	if matches, err := filepath.Glob("/usr/local/cuda-*"); err == nil {
		for _, m := range matches {
			updateBest(cudaVersionFromInstallDir(m))
		}
	}

	// 3b. Dynamic linker cache (ldconfig -p) on Linux.
	if ldconfigPath, err := exec.LookPath("ldconfig"); err == nil && ldconfigPath != "" {
		if out, err := exec.Command(ldconfigPath, "-p").Output(); err == nil {
			for _, line := range strings.Split(string(out), "\n") {
				if idx := strings.LastIndex(line, "=> "); idx != -1 {
					libPath := strings.TrimSpace(line[idx+3:])
					if strings.HasPrefix(filepath.Base(libPath), "libcudart.so") {
						updateBest(cudaVersionFromLibFile(libPath))
						updateBest(cudaVersionFromLibDir(filepath.Dir(libPath)))
					}
				}
			}
		}
	}

	// 3c. nvcc in PATH.
	if nvccPath, err := exec.LookPath("nvcc"); err == nil && nvccPath != "" {
		if out, err := exec.Command(nvccPath, "--version").Output(); err == nil {
			updateBest(parseCUDAFullVersionFromOutput(string(out)))
		}
	}

	// 3d. Standard system library directories.
	systemLibDirs := []string{
		"/usr/local/lib",
		"/usr/lib/x86_64-linux-gnu",
		"/usr/lib/aarch64-linux-gnu",
		"/usr/lib64",
		"/usr/lib",
	}
	for _, p := range systemLibDirs {
		updateBest(cudaVersionFromLibDir(p))
	}

	if bestVer != "" {
		return bestVer
	}
	return "12"
}

func parseCUDAFullVersionFromOutput(s string) string {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		// Prefer V<major>.<minor>.<patch> on the "release X.Y, VX.Y.Z" line if present.
		words := strings.Fields(line)
		for _, w := range words {
			w = strings.TrimSuffix(w, ",")
			if strings.HasPrefix(w, "V") && len(w) > 1 && unicode.IsDigit(rune(w[1])) {
				ver := w[1:]
				if _, _, _, _, ok := parseCUDAVersionComponents(ver); ok {
					return ver
				}
			}
		}
		if idx := strings.Index(line, "release "); idx != -1 {
			sub := line[idx+len("release "):]
			parts := strings.Split(sub, ",")
			if len(parts) > 0 {
				ver := strings.TrimSpace(parts[0])
				if _, _, _, _, ok := parseCUDAVersionComponents(ver); ok {
					return ver
				}
			}
		}
	}

	// Scan for cuda-<ver>, cuda_<ver>, cuda<ver>, or v<ver> in paths or strings.
	sub := strings.ToLower(s)
	bestVer := ""
	for _, prefix := range []string{"cuda-", "cuda_", "cuda", "v"} {
		search := sub
		for {
			idx := strings.Index(search, prefix)
			if idx == -1 {
				break
			}
			rest := search[idx+len(prefix):]
			search = rest
			if len(rest) == 0 || !unicode.IsDigit(rune(rest[0])) {
				continue
			}
			end := 0
			for end < len(rest) && (unicode.IsDigit(rune(rest[end])) || rest[end] == '.' || rest[end] == '-') {
				end++
			}
			candidate := strings.TrimRight(rest[:end], ".-")
			candidate = strings.ReplaceAll(candidate, "-", ".")
			if _, _, _, _, ok := parseCUDAVersionComponents(candidate); ok {
				if compareCUDAVersions(candidate, bestVer) > 0 {
					bestVer = candidate
				}
			}
		}
	}
	return bestVer
}

func parseMajorMinor(v string) (int, int, error) {
	major, minor, _, _, ok := parseCUDAVersionComponents(v)
	if !ok {
		return 0, 0, errors.Errorf("invalid CUDA version %q", v)
	}
	return major, minor, nil
}

func isCUDAVersionLE12_4(v string) bool {
	major, minor, err := parseMajorMinor(v)
	if err != nil {
		return false
	}
	if major < 12 {
		return true
	}
	if major == 12 && minor <= 4 {
		return true
	}
	return false
}

// GetLatestVersion fetches the latest release version of ONNX Runtime from GitHub.
// If cuda is true and the installed CUDA version is <= 12.4, it returns "1.27.0" to maintain compatibility.
func GetLatestVersion(cuda bool) (string, error) {
	return getLatestVersionForCUDA(cuda, "")
}

func getLatestVersionForCUDA(cuda bool, cudaVersion string) (string, error) {
	if cuda {
		cudaVer := strings.TrimSpace(cudaVersion)
		if cudaVer == "" {
			cudaVer = GetCUDAFullVersion()
		}
		if isCUDAVersionLE12_4(cudaVer) {
			klog.Infof("Current CUDA version %s (<= 12.4) is only supported for up to version v1.27 of ONNX Runtime", cudaVer)
			return "1.27.0", nil
		}
	}

	// 1. Check the 302 redirect Location header of /releases/latest without following the redirect
	// to the heavy HTML tag page (which can time out with 504 on HEAD requests).
	noRedirectClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := noRedirectClient.Head("https://github.com/microsoft/onnxruntime/releases/latest")
	if err == nil {
		resp.Body.Close()
		if loc := resp.Header.Get("Location"); loc != "" {
			if parts := strings.Split(loc, "/tag/"); len(parts) == 2 {
				return strings.TrimPrefix(strings.TrimSpace(parts[1]), "v"), nil
			}
		}
	}

	// 2. Fallback to the GitHub Releases REST API.
	apiResp, apiErr := http.Get("https://api.github.com/repos/microsoft/onnxruntime/releases/latest")
	if apiErr == nil {
		defer apiResp.Body.Close()
		if apiResp.StatusCode == http.StatusOK {
			var rel githubRelease
			if decErr := json.NewDecoder(apiResp.Body).Decode(&rel); decErr == nil && rel.TagName != "" {
				return strings.TrimPrefix(strings.TrimSpace(rel.TagName), "v"), nil
			}
		}
	}

	if err != nil {
		return "", errors.Wrap(err, "failed to fetch latest onnxruntime release from GitHub")
	}
	return "", errors.Errorf("failed to fetch latest onnxruntime release from GitHub: status code %d", resp.StatusCode)
}

// Install downloads and installs the ONNX Runtime library if not already present.
// It returns the absolute path to the main shared library file.
func Install(version string, cuda bool, cudaVersion string, targetDir string, force bool) (string, error) {
	if cudaVersion != "" {
		cuda = true
	}
	version = NormalizeVersion(version)
	if version == "" {
		version = DefaultVersion
	}
	if version == "" {
		var err error
		version, err = getLatestVersionForCUDA(cuda, cudaVersion)
		if err != nil {
			return "", err
		}
	}

	installDir := targetDir
	if installDir == "" {
		var err error
		installDir, err = GetInstallPath()
		if err != nil {
			return "", err
		}
	} else {
		installDir = filepath.Clean(installDir)
	}

	libFilename, err := GetLibFilename()
	if err != nil {
		return "", err
	}

	targetPath := filepath.Join(installDir, libFilename)
	if !force {
		if _, err := os.Stat(targetPath); err == nil {
			if installedVersion, _ := GetInstalledVersion(installDir); installedVersion != "" && installedVersion != version {
				force = true
			}
			if !force {
				if !cuda {
					return targetPath, nil
				}
				cudaLibPath := filepath.Join(installDir, "libonnxruntime_providers_cuda.so")
				if _, err := os.Stat(cudaLibPath); err == nil {
					return targetPath, nil
				}
				cudaLibPathWin := filepath.Join(installDir, "onnxruntime_providers_cuda.dll")
				if _, err := os.Stat(cudaLibPathWin); err == nil {
					return targetPath, nil
				}
			}
		}
	}

	// Needs installation
	url, err := GetAssetURL(version, cuda, cudaVersion)
	if err != nil {
		return "", err
	}

	fmt.Printf("Installing ONNX Runtime %s to %s...\n", version, installDir)
	fmt.Printf("Downloading %s ...\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "failed to download from %s", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("failed to download from %s: HTTP %s", url, resp.Status)
	}

	// Create temporary directory for extraction
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return "", errors.Wrapf(err, "failed to create directory %s", installDir)
	}

	tmpDir, err := os.MkdirTemp(installDir, ".extracting-")
	if err != nil {
		return "", errors.Wrap(err, "failed to create temp extraction directory")
	}
	defer os.RemoveAll(tmpDir)

	if strings.HasSuffix(url, ".tgz") {
		err = extractTarGz(resp.Body, tmpDir, libFilename)
	} else if strings.HasSuffix(url, ".zip") {
		err = extractZip(resp.Body, resp.ContentLength, tmpDir, libFilename)
	} else {
		err = errors.Errorf("unknown archive format for %s", url)
	}
	if err != nil {
		return "", err
	}

	// Move the extracted files into place
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		return "", errors.Wrap(err, "failed to read temp extraction directory")
	}

	if len(files) == 0 {
		return "", errors.Errorf("no library files extracted from the archive")
	}

	for _, file := range files {
		src := filepath.Join(tmpDir, file.Name())
		dest := filepath.Join(installDir, file.Name())

		// Remove existing dest if it's there
		_ = os.Remove(dest)
		if err := os.Rename(src, dest); err != nil {
			return "", errors.Wrapf(err, "failed to move extracted file %s to %s", file.Name(), dest)
		}
		fmt.Printf("Installed %s\n", dest)
	}

	// Write VERSION file to record the installed version.
	versionFile := filepath.Join(installDir, "VERSION")
	if err := os.WriteFile(versionFile, []byte(version+"\n"), 0644); err != nil {
		klog.Warningf("failed to write version file %s: %v", versionFile, err)
	}

	return targetPath, nil
}

func extractTarGz(r io.Reader, destDir, libFilename string) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return errors.Wrap(err, "failed to create gzip reader")
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "failed to read tar header")
		}

		baseName := filepath.Base(header.Name)
		// We only want the library files, which are in the 'lib/' folder
		// Typically they contain 'libonnxruntime' or the specific libFilename.
		if header.Typeflag == tar.TypeReg || header.Typeflag == tar.TypeSymlink {
			if strings.Contains(header.Name, "/lib/") && (strings.Contains(baseName, "libonnxruntime") || baseName == libFilename) {
				targetFile := filepath.Join(destDir, baseName)
				if header.Typeflag == tar.TypeSymlink {
					// We resolve/write symlinks, or write them as symlink.
					// For Go loads, symlinks are fine or copy the target directly.
					// Let's create the symlink.
					_ = os.Remove(targetFile)
					err = os.Symlink(header.Linkname, targetFile)
					if err != nil {
						return errors.Wrapf(err, "failed to create symlink %s -> %s", targetFile, header.Linkname)
					}
				} else {
					f, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
					if err != nil {
						return errors.Wrapf(err, "failed to create file %s", targetFile)
					}
					if _, err := io.Copy(f, tr); err != nil {
						f.Close()
						return errors.Wrapf(err, "failed to write file %s", targetFile)
					}
					f.Close()
				}
			}
		}
	}
	return nil
}

func extractZip(r io.Reader, contentLength int64, destDir, libFilename string) error {
	// zip.NewReader requires ReaderAt, so we must write to a temp file first.
	tmpZipFile, err := os.CreateTemp("", "onnxruntime-*.zip")
	if err != nil {
		return errors.Wrap(err, "failed to create temporary zip file")
	}
	defer func() {
		tmpZipFile.Close()
		os.Remove(tmpZipFile.Name())
	}()

	if _, err := io.Copy(tmpZipFile, r); err != nil {
		return errors.Wrap(err, "failed to save archive to temporary file")
	}

	stat, err := tmpZipFile.Stat()
	if err != nil {
		return errors.Wrap(err, "failed to stat temporary zip file")
	}

	zr, err := zip.NewReader(tmpZipFile, stat.Size())
	if err != nil {
		return errors.Wrap(err, "failed to create zip reader")
	}

	for _, file := range zr.File {
		baseName := filepath.Base(file.Name)
		if !file.FileInfo().IsDir() && (strings.Contains(file.Name, "/lib/") || strings.Contains(file.Name, "/bin/")) &&
			(strings.Contains(baseName, "onnxruntime") || baseName == libFilename) {
			targetFile := filepath.Join(destDir, baseName)
			f, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
			if err != nil {
				return errors.Wrapf(err, "failed to create file %s", targetFile)
			}
			src, err := file.Open()
			if err != nil {
				f.Close()
				return errors.Wrapf(err, "failed to open zip file entry %s", file.Name)
			}
			_, err = io.Copy(f, src)
			src.Close()
			f.Close()
			if err != nil {
				return errors.Wrapf(err, "failed to write file %s", targetFile)
			}
		}
	}
	return nil
}
