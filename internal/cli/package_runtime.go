package cli

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

type packageRuntimeMetadata struct {
	Platform             string   `json:"platform"`
	Arch                 string   `json:"arch"`
	Binary               string   `json:"binary"`
	Source               string   `json:"source"`
	Tags                 []string `json:"tags"`
	VendorManifestSHA256 string   `json:"vendorManifestSha256"`
}

const packageLadybugVersion = "v0.19.1"

type packageBuildRoots struct {
	LaneRoot    string
	CacheRoot   string
	RuntimeRoot string
}

type packageNativeFileIdentity struct {
	Name   string
	Bytes  int64
	SHA256 string
}

var packageWindowsNativeFiles = []packageNativeFileIdentity{
	{Name: "lbug.h", Bytes: 79108, SHA256: "3d5114d0863b3dab3b28bd2fec97a52e6cf669213739921a01814a5bbf5525eb"},
	{Name: "lbug_shared.lib", Bytes: 32433956, SHA256: "b18aafc0b712dc1c4cb9dd25f76c3828282d7d460627980e3a4b16efcd98a955"},
	{Name: "lbug_shared.dll", Bytes: 20230656, SHA256: "20cbd87840483a2053cff3fc2db23a86dd802b8915d86509d41a4b709624cdb7"},
}

func ensurePackagedRuntime(packageRoot string, output io.Writer) error {
	root, err := filepath.Abs(packageRoot)
	if err != nil {
		return err
	}
	outputPath := filepath.Join(root, "bin", "anvien.exe")
	metadataPath := filepath.Join(root, "bin", "anvien-runtime.json")
	stat, err := os.Stat(outputPath)
	if err != nil || stat.IsDir() {
		return fmt.Errorf("packaged Go runtime is missing: %s", outputPath)
	}
	raw, err := os.ReadFile(metadataPath)
	if err != nil {
		return fmt.Errorf("packaged Go runtime metadata is missing: %w", err)
	}
	var metadata packageRuntimeMetadata
	if err := json.Unmarshal(raw, &metadata); err != nil {
		return fmt.Errorf("packaged Go runtime metadata is invalid: %w", err)
	}
	if !platformMatches(metadata.Platform) || !archMatches(metadata.Arch) {
		return fmt.Errorf("packaged Go runtime is %s/%s, current platform is %s/%s", metadata.Platform, metadata.Arch, runtime.GOOS, runtime.GOARCH)
	}
	if metadata.Binary != filepath.Base(outputPath) {
		return fmt.Errorf("packaged Go runtime metadata names an unexpected binary: %q", metadata.Binary)
	}
	sourceRoot, err := resolvePackageSourceRoot(root)
	if err != nil {
		return err
	}
	if err := verifyPackageGoVendor(sourceRoot, output); err != nil {
		return err
	}
	vendorManifestSHA256, err := packageGoVendorManifestSHA256(sourceRoot)
	if err != nil {
		return err
	}
	if metadata.VendorManifestSHA256 == "" || !strings.EqualFold(metadata.VendorManifestSHA256, vendorManifestSHA256) {
		return fmt.Errorf("packaged Go runtime vendor manifest identity mismatch: metadata=%q source=%q", metadata.VendorManifestSHA256, vendorManifestSHA256)
	}
	if err := os.Chmod(outputPath, 0o755); err != nil {
		return err
	}
	_, err = fmt.Fprintf(output, "[package-runtime] using packaged Go runtime %s/%s with vendor manifest %s\n", metadata.Platform, metadata.Arch, vendorManifestSHA256)
	return err
}

func buildGoRuntimePackage(packageRoot string, output io.Writer) error {
	root, err := filepath.Abs(packageRoot)
	if err != nil {
		return err
	}
	sourceRoot, err := resolvePackageSourceRoot(root)
	if err != nil {
		return err
	}
	if err := verifyPackageGoVendor(sourceRoot, output); err != nil {
		return err
	}
	vendorManifestSHA256, err := packageGoVendorManifestSHA256(sourceRoot)
	if err != nil {
		return err
	}
	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("Go toolchain is required to build the packaged Anvien runtime: %w", err)
	}
	buildRoots, err := resolvePackageBuildRoots(sourceRoot)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(root, "bin")
	outputPath := filepath.Join(outputDir, "anvien.exe")
	metadataPath := filepath.Join(outputDir, "anvien-runtime.json")
	stageDir := filepath.Join(buildRoots.RuntimeRoot, "package-runtime")
	if err := assertPackageChild(buildRoots.RuntimeRoot, stageDir); err != nil {
		return err
	}
	stageOutputPath := filepath.Join(stageDir, "anvien.exe")
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(stageDir, 0o755); err != nil {
		return err
	}

	nativeDir, err := resolvePackageNativeDir(sourceRoot, buildRoots)
	if err != nil {
		return err
	}
	env := os.Environ()
	env = setEnv(env, "GOCACHE", filepath.Join(buildRoots.CacheRoot, "go-build"))
	env = setEnv(env, "GOMODCACHE", filepath.Join(buildRoots.CacheRoot, "go-mod"))
	env = setEnv(env, "GOPATH", filepath.Join(buildRoots.CacheRoot, "go-path"))
	env = setEnv(env, "GOTMPDIR", filepath.Join(buildRoots.CacheRoot, "go-tmp"))
	env = setEnv(env, "GOENV", "off")
	env = setEnv(env, "GOWORK", "off")
	env = setEnv(env, "GOFLAGS", "")
	env = setEnv(env, "GOPROXY", "off")
	env = setEnv(env, "GOSUMDB", "off")
	env = setEnv(env, "GOTOOLCHAIN", "local")
	env = setEnv(env, "GOPRIVATE", "")
	env = setEnv(env, "GONOPROXY", "none")
	env = setEnv(env, "GONOSUMDB", "none")
	env = setEnv(env, "GOINSECURE", "")
	env = setEnv(env, "GOVCS", "*:off")
	env = setEnv(env, "CGO_ENABLED", "1")
	env = setEnv(env, "CGO_CFLAGS", "-I"+nativeDir)
	env = setEnv(env, "PATH", nativeDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	env = setEnv(env, "DYLD_LIBRARY_PATH", nativeDir+string(os.PathListSeparator)+os.Getenv("DYLD_LIBRARY_PATH"))
	switch runtime.GOOS {
	case "windows":
		env = setEnv(env, "CGO_LDFLAGS", "-L"+nativeDir+" -llbug_shared")
	case "darwin":
		env = setEnv(env, "CGO_LDFLAGS", "-L"+nativeDir+" -llbug -Wl,-rpath,@loader_path")
	default:
		env = setEnv(env, "CGO_LDFLAGS", "-L"+nativeDir+" -llbug -Wl,-rpath,$ORIGIN")
		env = setEnv(env, "LD_LIBRARY_PATH", nativeDir+string(os.PathListSeparator)+os.Getenv("LD_LIBRARY_PATH"))
	}

	fmt.Fprintf(output, "[package-runtime] building Go runtime for %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Fprintf(output, "[package-runtime] Go source root: %s\n", sourceRoot)
	fmt.Fprintf(output, "[package-runtime] build lane root: %s\n", buildRoots.LaneRoot)
	fmt.Fprintf(output, "[package-runtime] build cache root: %s\n", buildRoots.CacheRoot)
	fmt.Fprintf(output, "[package-runtime] build runtime root: %s\n", buildRoots.RuntimeRoot)
	fmt.Fprintf(output, "[package-runtime] LadybugDB native runtime: %s\n", nativeDir)
	if err := runPackageCommand(output, sourceRoot, env, "go", "build", "-mod=vendor", "-tags", "ladybugdb", "-trimpath", "-ldflags=-s -w", "-o", stageOutputPath, "./cmd/anvien"); err != nil {
		return err
	}
	if err := os.Chmod(stageOutputPath, 0o755); err != nil {
		return err
	}
	if err := copyPackageFileIfExists(stageOutputPath, outputPath); err != nil {
		return err
	}
	if err := copyPackageNativeRuntime(nativeDir, outputDir); err != nil {
		return err
	}
	relativeSource, err := filepath.Rel(root, sourceRoot)
	if err != nil || strings.HasPrefix(relativeSource, "..") {
		relativeSource = sourceRoot
	}
	metadata := packageRuntimeMetadata{
		Platform:             runtime.GOOS,
		Arch:                 runtime.GOARCH,
		Binary:               "anvien.exe",
		Source:               filepath.ToSlash(relativeSource),
		Tags:                 []string{"ladybugdb"},
		VendorManifestSHA256: vendorManifestSHA256,
	}
	raw, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(metadataPath, append(raw, '\n'), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(output, "[package-runtime] wrote %s\n", outputPath)
	return nil
}

func prepareGoSourcePackage(packageRoot string, output io.Writer) error {
	root, err := filepath.Abs(packageRoot)
	if err != nil {
		return err
	}
	outputRoot := filepath.Join(root, "go-src")
	if err := assertPackageChild(root, outputRoot); err != nil {
		return err
	}
	repoRoot := filepath.Dir(root)
	if err := verifyPackageGoVendor(repoRoot, output); err != nil {
		return err
	}
	if err := os.RemoveAll(outputRoot); err != nil {
		return err
	}
	copied := make([]string, 0, 256)
	for _, rel := range []string{"go.mod", "go.sum"} {
		if err := copyPackageFile(filepath.Join(repoRoot, rel), filepath.Join(outputRoot, rel), outputRoot, &copied); err != nil {
			return err
		}
	}
	for _, rel := range []string{"cmd", "internal"} {
		if err := copyPackageGoDir(repoRoot, rel, outputRoot, &copied, "internal/aicontext/skills"); err != nil {
			return err
		}
	}
	if err := copyPackageSubtree(repoRoot, "internal/aicontext/skills", outputRoot, &copied); err != nil {
		return err
	}
	for _, rel := range []string{"scripts/ensure-ladybug-native.ps1", "scripts/ensure-ladybug-native.sh", "scripts/verify-go-vendor.ps1"} {
		if err := copyPackageFile(filepath.Join(repoRoot, rel), filepath.Join(outputRoot, rel), outputRoot, &copied); err != nil {
			return err
		}
	}
	if err := os.Chmod(filepath.Join(outputRoot, "scripts", "ensure-ladybug-native.sh"), 0o755); err != nil {
		return err
	}
	for _, rel := range []string{"vendor", "third_party/go-vendor"} {
		if err := copyPackageSubtree(repoRoot, rel, outputRoot, &copied); err != nil {
			return err
		}
	}
	nativeDir, err := resolvePinnedWindowsNativeDir(repoRoot)
	if err != nil {
		return err
	}
	for _, identity := range packageWindowsNativeFiles {
		rel := filepath.Join("third_party", "ladybugdb", packageLadybugVersion, "windows-x86_64", identity.Name)
		if err := copyPackageFile(filepath.Join(nativeDir, identity.Name), filepath.Join(outputRoot, rel), outputRoot, &copied); err != nil {
			return err
		}
	}
	sort.Strings(copied)
	manifest := map[string]any{
		"generatedBy": "anvien package prepare-go-source",
		"source":      "repo-root",
		"files":       len(copied),
		"paths":       copied,
	}
	raw, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(outputRoot, "anvien-go-source.json"), append(raw, '\n'), 0o644); err != nil {
		return err
	}
	_, err = fmt.Fprintf(output, "[prepare-go-source-package] copied %d files to %s\n", len(copied), outputRoot)
	return err
}

func resolvePackageSourceRoot(packageRoot string) (string, error) {
	repoRoot := filepath.Dir(packageRoot)
	if hasPackageGoSource(repoRoot) {
		return repoRoot, nil
	}
	packagedSourceRoot := filepath.Join(packageRoot, "go-src")
	if hasPackageGoSource(packagedSourceRoot) {
		return packagedSourceRoot, nil
	}
	return "", fmt.Errorf("Go source is not available and the packaged Go runtime does not match this platform")
}

func hasPackageGoSource(root string) bool {
	required := []string{
		"go.mod",
		"go.sum",
		"cmd/anvien/main.go",
		"internal/cli/command.go",
		"vendor/modules.txt",
		"third_party/go-vendor/manifest.v1.json",
		"scripts/verify-go-vendor.ps1",
	}
	for _, rel := range required {
		if stat, err := os.Stat(filepath.Join(root, rel)); err != nil || stat.IsDir() {
			return false
		}
	}
	return true
}

func verifyPackageGoVendor(sourceRoot string, output io.Writer) error {
	verifierPath := filepath.Join(sourceRoot, "scripts", "verify-go-vendor.ps1")
	stat, err := os.Stat(verifierPath)
	if err != nil || stat.IsDir() {
		return fmt.Errorf("Go vendor verifier is missing: %s", verifierPath)
	}
	powerShell := "pwsh"
	if runtime.GOOS == "windows" {
		powerShell = "powershell.exe"
	}
	cmd := exec.Command(powerShell, "-NoLogo", "-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-File", verifierPath, "-SourceRoot", sourceRoot, "-Json")
	cmd.Dir = sourceRoot
	cmd.Stdout = output
	cmd.Stderr = output
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Go vendor verification failed: %w", err)
	}
	return nil
}

func packageGoVendorManifestSHA256(sourceRoot string) (string, error) {
	manifestPath := filepath.Join(sourceRoot, "third_party", "go-vendor", "manifest.v1.json")
	file, err := os.Open(manifestPath)
	if err != nil {
		return "", fmt.Errorf("Go vendor manifest is missing: %w", err)
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		_ = file.Close()
		return "", err
	}
	if err := file.Close(); err != nil {
		return "", err
	}
	return strings.ToUpper(fmt.Sprintf("%x", hash.Sum(nil))), nil
}

func resolvePackageNativeDir(sourceRoot string, buildRoots packageBuildRoots) (string, error) {
	if runtime.GOOS == "windows" {
		authorityRoot := strings.TrimSpace(os.Getenv("ANVIEN_BUILD_REPO_ROOT"))
		if authorityRoot == "" {
			authorityRoot = sourceRoot
		}
		return resolvePinnedWindowsNativeDir(authorityRoot)
	}
	version := os.Getenv("ANVIEN_LADYBUGDB_VERSION")
	if strings.TrimSpace(version) == "" {
		version = "auto"
	}
	outputRoot := filepath.Join(buildRoots.CacheRoot, "ladybug-native")
	script, err := resolvePackageNativeScript(sourceRoot, "ensure-ladybug-native.sh")
	if err != nil {
		return "", err
	}
	return commandPackageOutput(sourceRoot, "bash", script, version, outputRoot)
}

func resolvePackageBuildRoots(sourceRoot string) (packageBuildRoots, error) {
	repoRoot := strings.TrimSpace(os.Getenv("ANVIEN_BUILD_REPO_ROOT"))
	if repoRoot == "" {
		repoRoot = sourceRoot
	}
	repoRoot, err := filepath.Abs(repoRoot)
	if err != nil {
		return packageBuildRoots{}, err
	}
	tempRoot := filepath.Join(repoRoot, ".tmp")
	if err := os.MkdirAll(tempRoot, 0o755); err != nil {
		return packageBuildRoots{}, err
	}

	laneRoot := strings.TrimSpace(os.Getenv("ANVIEN_BUILD_LANE_ROOT"))
	if laneRoot == "" {
		laneRoot, err = os.MkdirTemp(tempRoot, "package-runtime-")
		if err != nil {
			return packageBuildRoots{}, err
		}
	}
	laneRoot, err = filepath.Abs(laneRoot)
	if err != nil {
		return packageBuildRoots{}, err
	}
	if err := assertPackageChild(tempRoot, laneRoot); err != nil {
		return packageBuildRoots{}, fmt.Errorf("invalid package build lane root: %w", err)
	}

	cacheRoot := strings.TrimSpace(os.Getenv("ANVIEN_BUILD_CACHE_ROOT"))
	if cacheRoot == "" {
		cacheRoot = filepath.Join(laneRoot, "cache")
	}
	cacheRoot, err = filepath.Abs(cacheRoot)
	if err != nil {
		return packageBuildRoots{}, err
	}
	if err := assertPackageChild(laneRoot, cacheRoot); err != nil {
		return packageBuildRoots{}, fmt.Errorf("invalid package build cache root: %w", err)
	}

	runtimeRoot := strings.TrimSpace(os.Getenv("ANVIEN_BUILD_RUNTIME_ROOT"))
	if runtimeRoot == "" {
		runtimeRoot = filepath.Join(laneRoot, "runtime")
	}
	runtimeRoot, err = filepath.Abs(runtimeRoot)
	if err != nil {
		return packageBuildRoots{}, err
	}
	if err := assertPackageChild(laneRoot, runtimeRoot); err != nil {
		return packageBuildRoots{}, fmt.Errorf("invalid package build runtime root: %w", err)
	}
	if strings.EqualFold(cacheRoot, runtimeRoot) {
		return packageBuildRoots{}, fmt.Errorf("package build cache and runtime roots must be distinct: %s", cacheRoot)
	}

	for _, path := range []string{
		cacheRoot,
		runtimeRoot,
		filepath.Join(cacheRoot, "go-build"),
		filepath.Join(cacheRoot, "go-mod"),
		filepath.Join(cacheRoot, "go-path"),
		filepath.Join(cacheRoot, "go-tmp"),
	} {
		if err := os.MkdirAll(path, 0o755); err != nil {
			return packageBuildRoots{}, err
		}
	}
	return packageBuildRoots{LaneRoot: laneRoot, CacheRoot: cacheRoot, RuntimeRoot: runtimeRoot}, nil
}

func resolvePinnedWindowsNativeDir(authorityRoot string) (string, error) {
	version := strings.TrimSpace(os.Getenv("ANVIEN_LADYBUGDB_VERSION"))
	if version == "" {
		version = packageLadybugVersion
	}
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}
	if version != packageLadybugVersion {
		return "", fmt.Errorf("LadybugDB native authority is pinned to %s; requested %s", packageLadybugVersion, version)
	}
	authorityRoot, err := filepath.Abs(authorityRoot)
	if err != nil {
		return "", err
	}
	expected := filepath.Join(authorityRoot, "third_party", "ladybugdb", packageLadybugVersion, "windows-x86_64")
	configured := strings.TrimSpace(os.Getenv("ANVIEN_LADYBUGDB_NATIVE_DIR"))
	if configured == "" {
		configured = expected
	}
	configured, err = filepath.Abs(configured)
	if err != nil {
		return "", err
	}
	if !strings.EqualFold(filepath.Clean(configured), filepath.Clean(expected)) {
		return "", fmt.Errorf("LadybugDB native authority must be %s; requested %s", expected, configured)
	}
	if err := validatePinnedWindowsNativeBundle(configured); err != nil {
		return "", err
	}
	return configured, nil
}

func validatePinnedWindowsNativeBundle(nativeDir string) error {
	entries, err := os.ReadDir(nativeDir)
	if err != nil {
		return fmt.Errorf("LadybugDB native bundle is missing: %s: %w", nativeDir, err)
	}
	if len(entries) != len(packageWindowsNativeFiles) {
		return fmt.Errorf("LadybugDB native bundle must contain exactly %d files; found %d in %s", len(packageWindowsNativeFiles), len(entries), nativeDir)
	}
	expected := make(map[string]packageNativeFileIdentity, len(packageWindowsNativeFiles))
	for _, identity := range packageWindowsNativeFiles {
		expected[identity.Name] = identity
	}
	for _, entry := range entries {
		identity, ok := expected[entry.Name()]
		if !ok || !entry.Type().IsRegular() {
			return fmt.Errorf("LadybugDB native bundle contains an unauthorized entry: %s", filepath.Join(nativeDir, entry.Name()))
		}
		path := filepath.Join(nativeDir, entry.Name())
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		hash := sha256.New()
		bytesCopied, copyErr := io.Copy(hash, file)
		closeErr := file.Close()
		if copyErr != nil {
			return copyErr
		}
		if closeErr != nil {
			return closeErr
		}
		if bytesCopied != identity.Bytes || fmt.Sprintf("%x", hash.Sum(nil)) != identity.SHA256 {
			return fmt.Errorf("LadybugDB native input identity mismatch: %s", path)
		}
		delete(expected, entry.Name())
	}
	if len(expected) != 0 {
		return fmt.Errorf("LadybugDB native bundle is incomplete: %s", nativeDir)
	}
	return nil
}

func resolvePackageNativeScript(sourceRoot, scriptName string) (string, error) {
	repoRoot := filepath.Dir(sourceRoot)
	candidates := []string{
		filepath.Join(sourceRoot, "scripts", scriptName),
		filepath.Join(repoRoot, "scripts", scriptName),
	}
	for _, candidate := range candidates {
		if stat, err := os.Stat(candidate); err == nil && !stat.IsDir() {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("missing LadybugDB native resolver: %s", scriptName)
}

func commandPackageOutput(cwd, command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = cwd
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	raw, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%s failed: %w: %s", command, err, strings.TrimSpace(stderr.String()))
	}
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		if line := strings.TrimSpace(lines[i]); line != "" {
			return line, nil
		}
	}
	return "", fmt.Errorf("%s returned empty output", command)
}

func runPackageCommand(output io.Writer, cwd string, env []string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = cwd
	cmd.Env = env
	cmd.Stdout = output
	cmd.Stderr = output
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s failed: %w", strings.Join(append([]string{command}, args...), " "), err)
	}
	return nil
}

func copyPackageNativeRuntime(nativeDir, outputDir string) error {
	if runtime.GOOS == "windows" {
		return copyPackageFileIfExists(filepath.Join(nativeDir, "lbug_shared.dll"), filepath.Join(outputDir, "lbug_shared.dll"))
	}

	entries, err := os.ReadDir(nativeDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if runtime.GOOS == "darwin" {
			if strings.HasPrefix(name, "liblbug") && strings.HasSuffix(name, ".dylib") {
				if err := copyPackageFileIfExists(filepath.Join(nativeDir, name), filepath.Join(outputDir, name)); err != nil {
					return err
				}
			}
			continue
		}
		if name == "liblbug.so" || strings.HasPrefix(name, "liblbug.so.") {
			if err := copyPackageFileIfExists(filepath.Join(nativeDir, name), filepath.Join(outputDir, name)); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyPackageFileIfExists(source, destination string) error {
	if stat, err := os.Stat(source); err != nil || stat.IsDir() {
		return nil
	}
	raw, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	if existing, err := os.ReadFile(destination); err == nil && bytes.Equal(existing, raw) {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(destination, raw, 0o755); err != nil {
		return err
	}
	return os.Chmod(destination, 0o755)
}

func copyPackageGoDir(repoRoot, relativeDir, outputRoot string, copied *[]string, excludedDirs ...string) error {
	sourceDir := filepath.Join(repoRoot, relativeDir)
	if stat, err := os.Stat(sourceDir); err != nil || !stat.IsDir() {
		return fmt.Errorf("missing required source directory: %s", sourceDir)
	}
	stack := []string{sourceDir}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		entries, err := os.ReadDir(current)
		if err != nil {
			return err
		}
		sort.Slice(entries, func(i, j int) bool { return entries[i].Name() > entries[j].Name() })
		for _, entry := range entries {
			source := filepath.Join(current, entry.Name())
			if entry.IsDir() {
				skip, err := isPackageExcludedDir(repoRoot, source, excludedDirs)
				if err != nil {
					return err
				}
				if skip {
					continue
				}
				stack = append(stack, source)
				continue
			}
			rel, err := filepath.Rel(repoRoot, source)
			if err != nil {
				return err
			}
			goSource := strings.HasSuffix(entry.Name(), ".go") && !strings.HasSuffix(entry.Name(), "_test.go")
			if !entry.Type().IsRegular() || !goSource {
				continue
			}
			if err := copyPackageFile(source, filepath.Join(outputRoot, rel), outputRoot, copied); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyPackageSubtree(repoRoot, relativeDir, outputRoot string, copied *[]string) error {
	sourceDir := filepath.Join(repoRoot, relativeDir)
	stat, err := os.Stat(sourceDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("package subtree is not a directory: %s", sourceDir)
	}

	stack := []string{sourceDir}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		entries, err := os.ReadDir(current)
		if err != nil {
			return err
		}
		sort.Slice(entries, func(i, j int) bool { return entries[i].Name() > entries[j].Name() })
		for _, entry := range entries {
			source := filepath.Join(current, entry.Name())
			if entry.IsDir() {
				stack = append(stack, source)
				continue
			}
			if !entry.Type().IsRegular() {
				continue
			}
			rel, err := filepath.Rel(repoRoot, source)
			if err != nil {
				return err
			}
			if err := copyPackageFile(source, filepath.Join(outputRoot, rel), outputRoot, copied); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyPackageFile(source, destination, outputRoot string, copied *[]string) error {
	stat, err := os.Stat(source)
	if err != nil || stat.IsDir() {
		return fmt.Errorf("missing required source file: %s", source)
	}
	if err := assertPackageChild(outputRoot, destination); err != nil {
		return err
	}
	raw, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(destination, raw, 0o644); err != nil {
		return err
	}
	rel, err := filepath.Rel(outputRoot, destination)
	if err != nil {
		return err
	}
	*copied = append(*copied, filepath.ToSlash(rel))
	return nil
}

func isPackageExcludedDir(repoRoot, source string, excludedDirs []string) (bool, error) {
	rel, err := filepath.Rel(repoRoot, source)
	if err != nil {
		return false, err
	}
	relSlash := filepath.ToSlash(rel)
	for _, excluded := range excludedDirs {
		excludedSlash := filepath.ToSlash(filepath.Clean(excluded))
		if relSlash == excludedSlash || strings.HasPrefix(relSlash, excludedSlash+"/") {
			return true, nil
		}
	}
	return false, nil
}

func assertPackageChild(parent, child string) error {
	parentAbs, err := filepath.Abs(parent)
	if err != nil {
		return err
	}
	childAbs, err := filepath.Abs(child)
	if err != nil {
		return err
	}
	rel, err := filepath.Rel(parentAbs, childAbs)
	if err != nil {
		return err
	}
	if rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) || filepath.IsAbs(rel) {
		return fmt.Errorf("refusing to write outside package root: %s", childAbs)
	}
	return nil
}

func setEnv(env []string, key, value string) []string {
	prefix := key + "="
	for i, item := range env {
		if strings.HasPrefix(item, prefix) {
			env[i] = prefix + value
			return env
		}
	}
	return append(env, prefix+value)
}

func platformMatches(value string) bool {
	switch value {
	case "win32":
		value = "windows"
	}
	return value == runtime.GOOS
}

func archMatches(value string) bool {
	switch value {
	case "x64":
		value = "amd64"
	}
	return value == runtime.GOARCH
}
