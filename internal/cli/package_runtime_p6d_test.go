package cli

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
)

type packageTestFileIdentity struct {
	Path   string
	Bytes  int64
	SHA256 string
}

func TestP6DResolvePinnedWindowsNativeDirAcceptsDurableBundle(t *testing.T) {
	repoRoot := repositoryRootForPackageTest(t)
	nativeDir := filepath.Join(repoRoot, "third_party", "ladybugdb", packageLadybugVersion, "windows-x86_64")
	t.Setenv("ANVIEN_LADYBUGDB_VERSION", packageLadybugVersion)
	t.Setenv("ANVIEN_LADYBUGDB_NATIVE_DIR", nativeDir)

	got, err := resolvePinnedWindowsNativeDir(repoRoot)
	if err != nil {
		t.Fatalf("resolvePinnedWindowsNativeDir returned error: %v", err)
	}
	if got != nativeDir {
		t.Fatalf("native directory = %q, want %q", got, nativeDir)
	}
}

func TestP6DResolvePinnedWindowsNativeDirRejectsAlternateAuthority(t *testing.T) {
	repoRoot := repositoryRootForPackageTest(t)
	t.Setenv("ANVIEN_LADYBUGDB_VERSION", packageLadybugVersion)
	t.Setenv("ANVIEN_LADYBUGDB_NATIVE_DIR", filepath.Join(t.TempDir(), "windows-x86_64"))

	if _, err := resolvePinnedWindowsNativeDir(repoRoot); err == nil || !strings.Contains(err.Error(), "native authority must be") {
		t.Fatalf("alternate native authority error = %v, want explicit durable-authority failure", err)
	}
}

func TestP6DValidatePinnedWindowsNativeBundleFailsClosed(t *testing.T) {
	t.Run("partial", func(t *testing.T) {
		root := t.TempDir()
		if err := os.WriteFile(filepath.Join(root, "lbug.h"), []byte("partial"), 0o644); err != nil {
			t.Fatalf("write partial native file: %v", err)
		}
		if err := validatePinnedWindowsNativeBundle(root); err == nil || !strings.Contains(err.Error(), "exactly 3 files") {
			t.Fatalf("partial bundle error = %v, want exact-count failure", err)
		}
	})

	t.Run("tampered", func(t *testing.T) {
		root := t.TempDir()
		for _, identity := range packageWindowsNativeFiles {
			if err := os.WriteFile(filepath.Join(root, identity.Name), []byte("tampered"), 0o644); err != nil {
				t.Fatalf("write tampered %s: %v", identity.Name, err)
			}
		}
		if err := validatePinnedWindowsNativeBundle(root); err == nil || !strings.Contains(err.Error(), "identity mismatch") {
			t.Fatalf("tampered bundle error = %v, want identity failure", err)
		}
	})

	t.Run("extra", func(t *testing.T) {
		root := t.TempDir()
		copyPinnedNativeBundleForTest(t, root)
		nativeDir := filepath.Join(root, "third_party", "ladybugdb", packageLadybugVersion, "windows-x86_64")
		if err := os.WriteFile(filepath.Join(nativeDir, "lbug.hpp"), []byte("extra"), 0o644); err != nil {
			t.Fatalf("write extra native file: %v", err)
		}
		if err := validatePinnedWindowsNativeBundle(nativeDir); err == nil || !strings.Contains(err.Error(), "exactly 3 files") {
			t.Fatalf("extra bundle error = %v, want exact-count failure", err)
		}
	})
}

func TestP6DBuildGoRuntimePackageDoesNotAcceptStaleFallback(t *testing.T) {
	packageRoot := filepath.Join(t.TempDir(), "anvien")
	binRoot := filepath.Join(packageRoot, "bin")
	if err := os.MkdirAll(binRoot, 0o755); err != nil {
		t.Fatalf("mkdir bin root: %v", err)
	}
	if err := os.WriteFile(filepath.Join(packageRoot, "package.json"), []byte(`{"name":"anvien"}`), 0o644); err != nil {
		t.Fatalf("write package.json: %v", err)
	}
	if err := os.WriteFile(filepath.Join(binRoot, "anvien.exe"), []byte("stale"), 0o755); err != nil {
		t.Fatalf("write stale runtime: %v", err)
	}
	if err := os.WriteFile(filepath.Join(binRoot, "anvien-runtime.json"), []byte(`{"platform":"windows","arch":"amd64","binary":"anvien.exe"}`), 0o644); err != nil {
		t.Fatalf("write stale metadata: %v", err)
	}

	var output bytes.Buffer
	err := buildGoRuntimePackage(packageRoot, &output)
	if err == nil || !strings.Contains(err.Error(), "Go source is not available") {
		t.Fatalf("buildGoRuntimePackage error = %v, want fail-closed missing-source error; output=%q", err, output.String())
	}
}

func TestP6DEnsurePackagedRuntimeRejectsVendorManifestIdentityDrift(t *testing.T) {
	parent := shortPackageTestRoot(t)
	packageRoot := filepath.Join(parent, "anvien")
	sourceRoot := filepath.Join(packageRoot, "go-src")
	binRoot := filepath.Join(packageRoot, "bin")
	stageRealGoVendorAuthorityForPackageTest(t, sourceRoot)
	writePackageTestFile(t, sourceRoot, "cmd/anvien/main.go", "package main\n")
	writePackageTestFile(t, sourceRoot, "internal/cli/command.go", "package cli\n")
	if err := os.MkdirAll(binRoot, 0o755); err != nil {
		t.Fatalf("mkdir bin root: %v", err)
	}
	if err := os.WriteFile(filepath.Join(binRoot, "anvien.exe"), []byte("runtime"), 0o755); err != nil {
		t.Fatalf("write runtime: %v", err)
	}

	for _, test := range []struct {
		name     string
		identity string
	}{
		{name: "missing", identity: ""},
		{name: "drifted", identity: strings.Repeat("0", 64)},
	} {
		t.Run(test.name, func(t *testing.T) {
			metadata := packageRuntimeMetadata{
				Platform:             runtime.GOOS,
				Arch:                 runtime.GOARCH,
				Binary:               "anvien.exe",
				Source:               "..",
				Tags:                 []string{"ladybugdb"},
				VendorManifestSHA256: test.identity,
			}
			raw, err := json.Marshal(metadata)
			if err != nil {
				t.Fatalf("marshal metadata: %v", err)
			}
			if err := os.WriteFile(filepath.Join(binRoot, "anvien-runtime.json"), raw, 0o644); err != nil {
				t.Fatalf("write metadata: %v", err)
			}
			var output bytes.Buffer
			err = ensurePackagedRuntime(packageRoot, &output)
			if err == nil || !strings.Contains(err.Error(), "vendor manifest identity mismatch") {
				t.Fatalf("ensurePackagedRuntime error = %v, want vendor manifest identity mismatch; output=%q", err, output.String())
			}
			if strings.Contains(output.String(), "[package-runtime] using packaged Go runtime") {
				t.Fatalf("stale runtime was accepted: %q", output.String())
			}
		})
	}

	manifestSHA256, err := packageGoVendorManifestSHA256(sourceRoot)
	if err != nil {
		t.Fatalf("hash packaged vendor manifest: %v", err)
	}
	metadata := packageRuntimeMetadata{
		Platform:             runtime.GOOS,
		Arch:                 runtime.GOARCH,
		Binary:               "anvien.exe",
		Source:               "go-src",
		Tags:                 []string{"ladybugdb"},
		VendorManifestSHA256: manifestSHA256,
	}
	raw, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("marshal drift test metadata: %v", err)
	}
	if err := os.WriteFile(filepath.Join(binRoot, "anvien-runtime.json"), raw, 0o644); err != nil {
		t.Fatalf("write drift test metadata: %v", err)
	}
	manifestPath := filepath.Join(sourceRoot, "third_party", "go-vendor", "manifest.v1.json")
	restore := mutatePackageTestFile(t, manifestPath, false, func(raw []byte) []byte {
		return append(raw, '\n')
	})
	var output bytes.Buffer
	err = ensurePackagedRuntime(packageRoot, &output)
	restore()
	if err == nil || !strings.Contains(err.Error(), "vendor manifest identity mismatch") {
		t.Fatalf("ensurePackagedRuntime accepted byte-drifted manifest: err=%v output=%q", err, output.String())
	}
	if strings.Contains(output.String(), "[package-runtime] using packaged Go runtime") {
		t.Fatalf("byte-drifted manifest accepted stale runtime: %q", output.String())
	}
}

func TestP6DVerifyPackagedGoSourceVendorAuthorityFailsClosed(t *testing.T) {
	sourceRoot := filepath.Join(shortPackageTestRoot(t), "go-src")
	stageRealGoVendorAuthorityForPackageTest(t, sourceRoot)
	writePackageTestFile(t, sourceRoot, "cmd/anvien/main.go", "package main\n")
	writePackageTestFile(t, sourceRoot, "internal/cli/command.go", "package cli\n")
	var initialOutput bytes.Buffer
	if err := verifyPackageGoVendor(sourceRoot, &initialOutput); err != nil {
		t.Fatalf("initial packaged authority verification failed: %v; output=%q", err, initialOutput.String())
	}

	tests := []struct {
		name      string
		rel       string
		transform func([]byte) []byte
		remove    bool
	}{
		{name: "missing_manifest", rel: "third_party/go-vendor/manifest.v1.json", remove: true},
		{name: "tampered_manifest", rel: "third_party/go-vendor/manifest.v1.json", transform: func(raw []byte) []byte {
			return bytes.Replace(raw, []byte(`"closureContractVersion": 2`), []byte(`"closureContractVersion": 1`), 1)
		}},
		{name: "missing_vendor", rel: "vendor/modules.txt", remove: true},
		{name: "tampered_vendor", rel: "vendor/modules.txt", transform: func(raw []byte) []byte { return append(raw, []byte("# tampered\n")...) }},
		{name: "missing_patch", rel: "third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch", remove: true},
		{name: "tampered_patch", rel: "third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch", transform: func(raw []byte) []byte { return append(raw, []byte("# tampered\n")...) }},
		{name: "missing_verifier", rel: "scripts/verify-go-vendor.ps1", remove: true},
		{name: "tampered_verifier", rel: "scripts/verify-go-vendor.ps1", transform: func(raw []byte) []byte { return append(raw, []byte("\n# tampered\n")...) }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(sourceRoot, filepath.FromSlash(test.rel))
			restore := mutatePackageTestFile(t, path, test.remove, test.transform)
			var output bytes.Buffer
			err := verifyPackageGoVendor(sourceRoot, &output)
			restore()
			if err == nil {
				t.Fatalf("verifyPackageGoVendor accepted invalid %s; output=%q", test.rel, output.String())
			}
			if strings.Contains(output.String(), `"status":"PASS"`) {
				t.Fatalf("invalid packaged authority reported PASS: %q", output.String())
			}
		})
	}
}

func repositoryRootForPackageTest(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("resolve repository root: %v", err)
	}
	return root
}

func shortPackageTestRoot(t *testing.T) string {
	t.Helper()
	root := filepath.Join(repositoryRootForPackageTest(t), ".tmp")
	if err := os.MkdirAll(root, 0o755); err != nil {
		t.Fatalf("mkdir package test temp root: %v", err)
	}
	testRoot, err := os.MkdirTemp(root, "p6d-pkg-")
	if err != nil {
		t.Fatalf("create short package test root: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(testRoot); err != nil {
			t.Errorf("remove short package test root %s: %v", testRoot, err)
		}
	})
	return testRoot
}

func copyPinnedNativeBundleForTest(t *testing.T, targetRoot string) {
	t.Helper()
	repoRoot := repositoryRootForPackageTest(t)
	sourceRoot := filepath.Join(repoRoot, "third_party", "ladybugdb", packageLadybugVersion, "windows-x86_64")
	destinationRoot := filepath.Join(targetRoot, "third_party", "ladybugdb", packageLadybugVersion, "windows-x86_64")
	if err := os.MkdirAll(destinationRoot, 0o755); err != nil {
		t.Fatalf("mkdir native destination: %v", err)
	}
	for _, identity := range packageWindowsNativeFiles {
		source, err := os.Open(filepath.Join(sourceRoot, identity.Name))
		if err != nil {
			t.Fatalf("open durable native %s: %v", identity.Name, err)
		}
		destination, err := os.Create(filepath.Join(destinationRoot, identity.Name))
		if err != nil {
			_ = source.Close()
			t.Fatalf("create native fixture %s: %v", identity.Name, err)
		}
		_, copyErr := io.Copy(destination, source)
		closeDestinationErr := destination.Close()
		closeSourceErr := source.Close()
		if copyErr != nil {
			t.Fatalf("copy native fixture %s: %v", identity.Name, copyErr)
		}
		if closeDestinationErr != nil || closeSourceErr != nil {
			t.Fatalf("close native fixture %s: destination=%v source=%v", identity.Name, closeDestinationErr, closeSourceErr)
		}
	}
}

func stageRealGoVendorAuthorityForPackageTest(t *testing.T, targetRoot string) {
	t.Helper()
	repoRoot := repositoryRootForPackageTest(t)
	for _, rel := range []string{"go.mod", "go.sum", "scripts/verify-go-vendor.ps1", "vendor", "third_party/go-vendor"} {
		copyPackageTestPath(t, filepath.Join(repoRoot, filepath.FromSlash(rel)), filepath.Join(targetRoot, filepath.FromSlash(rel)))
	}
}

func copyPackageTestPath(t *testing.T, source, destination string) {
	t.Helper()
	info, err := os.Lstat(source)
	if err != nil {
		t.Fatalf("lstat package fixture source %s: %v", source, err)
	}
	if info.Mode()&os.ModeSymlink != 0 {
		t.Fatalf("package fixture source is a symlink: %s", source)
	}
	if info.IsDir() {
		if err := os.MkdirAll(destination, info.Mode().Perm()); err != nil {
			t.Fatalf("mkdir package fixture %s: %v", destination, err)
		}
		entries, err := os.ReadDir(source)
		if err != nil {
			t.Fatalf("read package fixture directory %s: %v", source, err)
		}
		for _, entry := range entries {
			copyPackageTestPath(t, filepath.Join(source, entry.Name()), filepath.Join(destination, entry.Name()))
		}
		return
	}
	if !info.Mode().IsRegular() {
		t.Fatalf("package fixture source is not a regular file: %s", source)
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		t.Fatalf("mkdir package fixture parent %s: %v", filepath.Dir(destination), err)
	}
	sourceFile, err := os.Open(source)
	if err != nil {
		t.Fatalf("open package fixture source %s: %v", source, err)
	}
	destinationFile, err := os.OpenFile(destination, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode().Perm())
	if err != nil {
		_ = sourceFile.Close()
		t.Fatalf("create package fixture destination %s: %v", destination, err)
	}
	_, copyErr := io.Copy(destinationFile, sourceFile)
	closeDestinationErr := destinationFile.Close()
	closeSourceErr := sourceFile.Close()
	if copyErr != nil || closeDestinationErr != nil || closeSourceErr != nil {
		t.Fatalf("copy package fixture %s: copy=%v destinationClose=%v sourceClose=%v", source, copyErr, closeDestinationErr, closeSourceErr)
	}
}

func inventoryPackageTreeForTest(t *testing.T, root string) []packageTestFileIdentity {
	t.Helper()
	rows := make([]packageTestFileIdentity, 0, 2048)
	var walk func(string, string)
	walk = func(current, relative string) {
		entries, err := os.ReadDir(current)
		if err != nil {
			t.Fatalf("read package tree %s: %v", current, err)
		}
		for _, entry := range entries {
			path := filepath.Join(current, entry.Name())
			rel := entry.Name()
			if relative != "" {
				rel = filepath.Join(relative, entry.Name())
			}
			info, err := os.Lstat(path)
			if err != nil {
				t.Fatalf("lstat package tree path %s: %v", path, err)
			}
			if info.Mode()&os.ModeSymlink != 0 {
				t.Fatalf("package tree contains a symlink: %s", path)
			}
			if info.IsDir() {
				walk(path, rel)
				continue
			}
			if !info.Mode().IsRegular() {
				t.Fatalf("package tree contains a non-regular file: %s", path)
			}
			file, err := os.Open(path)
			if err != nil {
				t.Fatalf("open package tree file %s: %v", path, err)
			}
			hash := sha256.New()
			bytesCopied, copyErr := io.Copy(hash, file)
			closeErr := file.Close()
			if copyErr != nil || closeErr != nil {
				t.Fatalf("hash package tree file %s: copy=%v close=%v", path, copyErr, closeErr)
			}
			rows = append(rows, packageTestFileIdentity{Path: filepath.ToSlash(rel), Bytes: bytesCopied, SHA256: fmt.Sprintf("%X", hash.Sum(nil))})
		}
	}
	walk(root, "")
	sort.Slice(rows, func(i, j int) bool { return rows[i].Path < rows[j].Path })
	return rows
}

func assertPackageTreeByteExactForTest(t *testing.T, sourceRoot, destinationRoot string) {
	t.Helper()
	sourceRows := inventoryPackageTreeForTest(t, sourceRoot)
	destinationRows := inventoryPackageTreeForTest(t, destinationRoot)
	if len(sourceRows) != len(destinationRows) {
		t.Fatalf("package tree file count differs: source=%d destination=%d", len(sourceRows), len(destinationRows))
	}
	for i := range sourceRows {
		if sourceRows[i] != destinationRows[i] {
			t.Fatalf("package tree identity differs at row %d: source=%#v destination=%#v", i, sourceRows[i], destinationRows[i])
		}
	}
}

func assertPackageFileByteExactForTest(t *testing.T, source, destination string) {
	t.Helper()
	sourceRaw, err := os.ReadFile(source)
	if err != nil {
		t.Fatalf("read package source file %s: %v", source, err)
	}
	destinationRaw, err := os.ReadFile(destination)
	if err != nil {
		t.Fatalf("read packaged file %s: %v", destination, err)
	}
	if !bytes.Equal(sourceRaw, destinationRaw) {
		t.Fatalf("packaged file differs from source: %s", destination)
	}
}

func assertPackageSourceManifestCoversTreeForTest(t *testing.T, sourceRoot string, files int, paths []string) int {
	t.Helper()
	rows := inventoryPackageTreeForTest(t, sourceRoot)
	expected := make([]string, 0, len(rows)-1)
	for _, row := range rows {
		if row.Path != "anvien-go-source.json" {
			expected = append(expected, row.Path)
		}
	}
	if files != len(paths) || files != len(expected) {
		t.Fatalf("package source manifest denominator differs: files=%d paths=%d actual=%d", files, len(paths), len(expected))
	}
	for i := range expected {
		if paths[i] != expected[i] {
			t.Fatalf("package source manifest path differs at row %d: manifest=%q actual=%q", i, paths[i], expected[i])
		}
	}
	return len(expected)
}

func mutatePackageTestFile(t *testing.T, path string, remove bool, transform func([]byte) []byte) func() {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat package test mutation target %s: %v", path, err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read package test mutation target %s: %v", path, err)
	}
	if remove {
		if err := os.Remove(path); err != nil {
			t.Fatalf("remove package test mutation target %s: %v", path, err)
		}
	} else {
		mutated := transform(append([]byte(nil), raw...))
		if bytes.Equal(mutated, raw) {
			t.Fatalf("package test mutation did not change %s", path)
		}
		if err := os.WriteFile(path, mutated, info.Mode().Perm()); err != nil {
			t.Fatalf("write package test mutation %s: %v", path, err)
		}
	}
	return func() {
		if err := os.WriteFile(path, raw, info.Mode().Perm()); err != nil {
			t.Fatalf("restore package test mutation %s: %v", path, err)
		}
	}
}
