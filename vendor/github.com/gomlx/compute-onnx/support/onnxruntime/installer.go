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

// GetDefaultCUDAVersion attempts to guess the CUDA version based on installed libraries found on standard paths,
// or via nvcc if available. It returns the major version as a string (e.g., "12"). If nothing is found, it defaults to "12".
func GetDefaultCUDAVersion() string {
	// 1. Try nvcc executable
	if nvccPath, err := exec.LookPath("nvcc"); err == nil && nvccPath != "" {
		cmd := exec.Command(nvccPath, "--version")
		if out, err := cmd.Output(); err == nil {
			if ver := parseCudaVersionFromOutput(string(out)); ver != "" {
				return ver
			}
		}
	}

	// 2. Check CUDA environment variables
	for _, envVar := range []string{"CUDA_PATH", "CUDA_HOME"} {
		if val := os.Getenv(envVar); val != "" {
			if ver := parseCudaVersionFromOutput(val); ver != "" {
				return ver
			}
		}
	}

	// 3. Search system library paths for libcudart.so.* or cudart64_*.dll
	searchPaths := []string{
		"/usr/lib/x86_64-linux-gnu",
		"/usr/local/cuda/lib64",
		"/usr/local/cuda/lib",
		"/usr/local/lib",
		"/usr/lib",
		"/usr/lib64",
	}
	if ldPath := os.Getenv("LD_LIBRARY_PATH"); ldPath != "" {
		searchPaths = append(strings.Split(ldPath, string(os.PathListSeparator)), searchPaths...)
	}

	for _, p := range searchPaths {
		if p == "" {
			continue
		}
		// Linux: libcudart.so.<version>
		if matches, err := filepath.Glob(filepath.Join(p, "libcudart.so.*")); err == nil && len(matches) > 0 {
			for _, m := range matches {
				base := filepath.Base(m)
				parts := strings.Split(base, "libcudart.so.")
				if len(parts) == 2 && parts[1] != "" {
					verParts := strings.Split(parts[1], ".")
					if len(verParts) > 0 && verParts[0] != "" {
						return verParts[0]
					}
				}
			}
		}
		// Windows: cudart64_<version>.dll or cudart64_12.dll
		if matches, err := filepath.Glob(filepath.Join(p, "cudart64_*.dll")); err == nil && len(matches) > 0 {
			for _, m := range matches {
				base := filepath.Base(m)
				base = strings.TrimPrefix(base, "cudart64_")
				base = strings.TrimSuffix(base, ".dll")
				if len(base) >= 2 {
					return base[:2]
				}
			}
		}
	}

	return "12"
}

func cleanCudaVersion(v string) string {
	v = strings.ToLower(strings.TrimSpace(v))
	v = strings.TrimPrefix(v, "cuda")
	v = strings.TrimPrefix(v, "v")
	parts := strings.Split(v, ".")
	if len(parts) > 0 {
		return parts[0]
	}
	return v
}

func parseCudaVersionFromOutput(s string) string {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if idx := strings.Index(line, "release "); idx != -1 {
			sub := line[idx+len("release "):]
			parts := strings.Split(sub, ".")
			if len(parts) > 0 && parts[0] != "" {
				if _, err := strconv.Atoi(parts[0]); err == nil {
					return parts[0]
				}
			}
		}
		words := strings.Fields(line)
		for _, w := range words {
			w = strings.TrimSuffix(w, ",")
			if strings.HasPrefix(w, "V") && len(w) > 1 && unicode.IsDigit(rune(w[1])) {
				vStr := w[1:]
				parts := strings.Split(vStr, ".")
				if len(parts) > 0 && parts[0] != "" {
					if _, err := strconv.Atoi(parts[0]); err == nil {
						return parts[0]
					}
				}
			}
		}
	}
	sub := strings.ToLower(s)
	for _, ver := range []string{"13", "12", "11"} {
		if strings.Contains(sub, "v"+ver) || strings.Contains(sub, "cuda-"+ver) || strings.Contains(sub, "cuda"+ver) {
			return ver
		}
	}
	return ""
}

type githubRelease struct {
	TagName string `json:"tag_name"`
}

var knownLatestPatch = map[string]string{
	"1.27": "1.27.1",
	"1.24": "1.24.4",
	"1.29": "1.29.0",
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

// GetCUDAFullVersion attempts to return the full CUDA version string (e.g. "12.4" or "12.4.131") found on the system.
func GetCUDAFullVersion() string {
	// 1. Try nvcc executable
	if nvccPath, err := exec.LookPath("nvcc"); err == nil && nvccPath != "" {
		cmd := exec.Command(nvccPath, "--version")
		if out, err := cmd.Output(); err == nil {
			if ver := parseCUDAFullVersionFromOutput(string(out)); ver != "" {
				return ver
			}
		}
	}

	// 2. Check CUDA environment variables
	for _, envVar := range []string{"CUDA_PATH", "CUDA_HOME"} {
		if val := os.Getenv(envVar); val != "" {
			if ver := parseCUDAFullVersionFromOutput(val); ver != "" {
				return ver
			}
		}
	}

	// 3. Search system library paths for libcudart.so.*
	searchPaths := []string{
		"/usr/lib/x86_64-linux-gnu",
		"/usr/local/cuda/lib64",
		"/usr/local/cuda/lib",
		"/usr/local/lib",
		"/usr/lib",
		"/usr/lib64",
	}
	if ldPath := os.Getenv("LD_LIBRARY_PATH"); ldPath != "" {
		searchPaths = append(strings.Split(ldPath, string(os.PathListSeparator)), searchPaths...)
	}

	for _, p := range searchPaths {
		if p == "" {
			continue
		}
		if matches, err := filepath.Glob(filepath.Join(p, "libcudart.so.*")); err == nil && len(matches) > 0 {
			for _, m := range matches {
				base := filepath.Base(m)
				parts := strings.Split(base, "libcudart.so.")
				if len(parts) == 2 && parts[1] != "" {
					return parts[1]
				}
			}
		}
	}

	return GetDefaultCUDAVersion()
}

func parseCUDAFullVersionFromOutput(s string) string {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if idx := strings.Index(line, "release "); idx != -1 {
			sub := line[idx+len("release "):]
			parts := strings.Split(sub, ",")
			if len(parts) > 0 {
				ver := strings.TrimSpace(parts[0])
				if ver != "" {
					return ver
				}
			}
		}
		words := strings.Fields(line)
		for _, w := range words {
			w = strings.TrimSuffix(w, ",")
			if strings.HasPrefix(w, "V") && len(w) > 1 && unicode.IsDigit(rune(w[1])) {
				return w[1:]
			}
		}
	}
	sub := strings.ToLower(s)
	for _, ver := range []string{"13", "12.4", "12.3", "12.2", "12.1", "12.0", "12", "11.8", "11"} {
		if strings.Contains(sub, "v"+ver) || strings.Contains(sub, "cuda-"+ver) || strings.Contains(sub, "cuda"+ver) {
			return ver
		}
	}
	return ""
}

func parseMajorMinor(v string) (int, int, error) {
	v = strings.ToLower(strings.TrimSpace(v))
	v = strings.TrimPrefix(v, "cuda")
	v = strings.TrimPrefix(v, "v")
	parts := strings.Split(v, ".")
	if len(parts) == 0 || parts[0] == "" {
		return 0, 0, errors.New("empty version")
	}
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	minor := 0
	if len(parts) > 1 && parts[1] != "" {
		m, err := strconv.Atoi(parts[1])
		if err == nil {
			minor = m
		}
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
	if cuda {
		cudaVer := GetCUDAFullVersion()
		if isCUDAVersionLE12_4(cudaVer) {
			klog.Infof("Current CUDA installed is %s (<= 12.4) and it is only supported for up to version v1.27 of ONNX Runtime", cudaVer)
			return "1.27.0", nil
		}
	}

	resp, err := http.Head("https://github.com/microsoft/onnxruntime/releases/latest")
	if err != nil {
		return "", errors.Wrap(err, "failed to fetch latest onnxruntime release from GitHub")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return "", errors.Errorf("failed to fetch latest onnxruntime release from GitHub: status code %d", resp.StatusCode)
	}

	if resp.Request == nil {
		return "", errors.New("failed to fetch latest onnxruntime release: response Request is nil")
	}

	finalURL := resp.Request.URL.String()
	// The URL is expected to be of the form: https://github.com/microsoft/onnxruntime/releases/tag/vX.Y.Z
	// We want to extract "X.Y.Z"
	parts := strings.Split(finalURL, "/tag/")
	if len(parts) != 2 {
		return "", errors.Errorf("unexpected redirect URL when fetching latest release: %s", finalURL)
	}

	version := strings.TrimPrefix(parts[1], "v")
	return version, nil
}

// Install downloads and installs the ONNX Runtime library if not already present.
// It returns the absolute path to the main shared library file.
func Install(version string, cuda bool, cudaVersion string, targetDir string, force bool) (string, error) {
	version = NormalizeVersion(version)
	if version == "" {
		version = DefaultVersion
	}
	if version == "" {
		var err error
		version, err = GetLatestVersion(cuda)
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
