package cli

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCleanGoSourcePackageRemovesOnlyPackageGoSrc(t *testing.T) {
	root := t.TempDir()
	goSrc := filepath.Join(root, "go-src")
	nested := filepath.Join(goSrc, "internal")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatalf("mkdir go-src: %v", err)
	}
	if err := os.WriteFile(filepath.Join(nested, "x.go"), []byte("package x\n"), 0o644); err != nil {
		t.Fatalf("write go-src file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "package.json"), []byte(`{"name":"anvien"}`), 0o644); err != nil {
		t.Fatalf("write package.json: %v", err)
	}

	var out bytes.Buffer
	if err := cleanGoSourcePackage(root, &out); err != nil {
		t.Fatalf("cleanGoSourcePackage returned error: %v", err)
	}
	if _, err := os.Stat(goSrc); !os.IsNotExist(err) {
		t.Fatalf("go-src still exists or stat returned unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "[clean-go-source-package] removed ") {
		t.Fatalf("cleanup output missing status line: %q", out.String())
	}
	if _, err := os.Stat(filepath.Join(root, "package.json")); err != nil {
		t.Fatalf("package root file was removed: %v", err)
	}
}

func TestPackageCleanGoSourceCommandUsesWorkingDirectoryPackageRoot(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "package.json"), []byte(`{"name":"anvien"}`), 0o644); err != nil {
		t.Fatalf("write package.json: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(root, "go-src"), 0o755); err != nil {
		t.Fatalf("mkdir go-src: %v", err)
	}
	previous, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatalf("chdir package root: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(previous)
	})

	out, errOut, err := executeForTest(t, "package", "clean-go-source")
	if err != nil {
		t.Fatalf("package clean-go-source returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("package clean-go-source wrote stderr: %q", errOut)
	}
	if _, err := os.Stat(filepath.Join(root, "go-src")); !os.IsNotExist(err) {
		t.Fatalf("go-src still exists or stat returned unexpected error: %v", err)
	}
}

func TestEnsurePackagedRuntimeAcceptsCurrentPlatformMetadata(t *testing.T) {
	root := t.TempDir()
	binDir := filepath.Join(root, "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatalf("mkdir bin: %v", err)
	}
	if err := os.WriteFile(filepath.Join(binDir, "anvien.exe"), []byte("runtime"), 0o644); err != nil {
		t.Fatalf("write runtime: %v", err)
	}
	metadata := packageRuntimeMetadata{
		Platform: runtime.GOOS,
		Arch:     runtime.GOARCH,
		Binary:   "anvien.exe",
		Source:   "..",
		Tags:     []string{"ladybugdb"},
	}
	raw, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("marshal metadata: %v", err)
	}
	if err := os.WriteFile(filepath.Join(binDir, "anvien-runtime.json"), raw, 0o644); err != nil {
		t.Fatalf("write metadata: %v", err)
	}

	var out bytes.Buffer
	if err := ensurePackagedRuntime(root, &out); err != nil {
		t.Fatalf("ensurePackagedRuntime returned error: %v", err)
	}
	if !strings.Contains(out.String(), "[package-runtime] using packaged Go runtime") {
		t.Fatalf("ensure output missing status: %q", out.String())
	}
}

func TestPackageRuntimePlatformAndArchAliases(t *testing.T) {
	if !platformMatches(runtime.GOOS) {
		t.Fatalf("platformMatches(%q) returned false", runtime.GOOS)
	}
	for value, goos := range map[string]string{
		"win32":  "windows",
		"darwin": "darwin",
		"linux":  "linux",
	} {
		if got, want := platformMatches(value), runtime.GOOS == goos; got != want {
			t.Fatalf("platformMatches(%q) = %v, want %v for %s", value, got, want, runtime.GOOS)
		}
	}

	if !archMatches(runtime.GOARCH) {
		t.Fatalf("archMatches(%q) returned false", runtime.GOARCH)
	}
	for value, goarch := range map[string]string{
		"x64":   "amd64",
		"arm64": "arm64",
	} {
		if got, want := archMatches(value), runtime.GOARCH == goarch; got != want {
			t.Fatalf("archMatches(%q) = %v, want %v for %s", value, got, want, runtime.GOARCH)
		}
	}
}

func TestPrepareGoSourcePackageCopiesMinimalGoSource(t *testing.T) {
	parent := t.TempDir()
	packageRoot := filepath.Join(parent, "anvien")
	if err := os.MkdirAll(packageRoot, 0o755); err != nil {
		t.Fatalf("mkdir package root: %v", err)
	}
	if err := os.WriteFile(filepath.Join(packageRoot, "package.json"), []byte(`{"name":"anvien"}`), 0o644); err != nil {
		t.Fatalf("write package.json: %v", err)
	}
	writePackageTestFile(t, parent, "go.mod", "module example.com/anvien\n")
	writePackageTestFile(t, parent, "go.sum", "")
	writePackageTestFile(t, parent, "cmd/anvien/main.go", "package main\n")
	writePackageTestFile(t, parent, "cmd/anvien/main_test.go", "package main\n")
	writePackageTestFile(t, parent, "internal/aicontext/aicontext.go", "package aicontext\n")
	writePackageTestFile(t, parent, "internal/aicontext/skills/planner/SKILL.md", "# Planner\n")
	writePackageTestFile(t, parent, "internal/aicontext/skills/planner/assets/config.json", `{"mode":"plan"}`+"\n")
	writePackageTestFile(t, parent, "internal/cli/command.go", "package cli\n")
	writePackageTestFile(t, parent, "internal/cli/command_test.go", "package cli\n")
	writePackageTestFile(t, parent, "scripts/ensure-ladybug-native.ps1", "Write-Output native\n")
	writePackageTestFile(t, parent, "scripts/ensure-ladybug-native.sh", "#!/usr/bin/env bash\nprintf native\\n\n")
	if err := os.MkdirAll(filepath.Join(packageRoot, "go-src", "old"), 0o755); err != nil {
		t.Fatalf("mkdir old go-src: %v", err)
	}

	var out bytes.Buffer
	if err := prepareGoSourcePackage(packageRoot, &out); err != nil {
		t.Fatalf("prepareGoSourcePackage returned error: %v", err)
	}
	for _, rel := range []string{
		"go.mod",
		"go.sum",
		"cmd/anvien/main.go",
		"internal/aicontext/aicontext.go",
		"internal/aicontext/skills/planner/SKILL.md",
		"internal/aicontext/skills/planner/assets/config.json",
		"internal/cli/command.go",
		"scripts/ensure-ladybug-native.ps1",
		"scripts/ensure-ladybug-native.sh",
		"anvien-go-source.json",
	} {
		if _, err := os.Stat(filepath.Join(packageRoot, "go-src", filepath.FromSlash(rel))); err != nil {
			t.Fatalf("prepared source missing %s: %v", rel, err)
		}
	}
	for _, rel := range []string{"cmd/anvien/main_test.go", "internal/cli/command_test.go", "old"} {
		if _, err := os.Stat(filepath.Join(packageRoot, "go-src", filepath.FromSlash(rel))); !os.IsNotExist(err) {
			t.Fatalf("prepared source retained excluded path %s: %v", rel, err)
		}
	}
	var manifest struct {
		Files int      `json:"files"`
		Paths []string `json:"paths"`
	}
	raw, err := os.ReadFile(filepath.Join(packageRoot, "go-src", "anvien-go-source.json"))
	if err != nil {
		t.Fatalf("read source manifest: %v", err)
	}
	if err := json.Unmarshal(raw, &manifest); err != nil {
		t.Fatalf("parse source manifest: %v", err)
	}
	if manifest.Files != 9 {
		t.Fatalf("manifest files = %d, want 9", manifest.Files)
	}
	for _, rel := range []string{
		"internal/aicontext/skills/planner/SKILL.md",
		"internal/aicontext/skills/planner/assets/config.json",
	} {
		if !containsString(manifest.Paths, rel) {
			t.Fatalf("manifest paths missing staged skill path %s: %#v", rel, manifest.Paths)
		}
	}
	for _, rel := range []string{"cmd/anvien/main_test.go", "internal/cli/command_test.go"} {
		if containsString(manifest.Paths, rel) {
			t.Fatalf("manifest paths retained excluded path %s: %#v", rel, manifest.Paths)
		}
	}
	if !strings.Contains(out.String(), "[prepare-go-source-package] copied 9 files") {
		t.Fatalf("prepare output missing copied count: %q", out.String())
	}
}

func TestCopyPackageFileIfExistsSkipsIdenticalDestination(t *testing.T) {
	root := t.TempDir()
	source := filepath.Join(root, "native", "lbug_shared.dll")
	destination := filepath.Join(root, "bin", "lbug_shared.dll")
	raw := []byte("native-runtime")
	if err := os.MkdirAll(filepath.Dir(source), 0o755); err != nil {
		t.Fatalf("mkdir source dir: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		t.Fatalf("mkdir destination dir: %v", err)
	}
	if err := os.WriteFile(source, raw, 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	if err := os.WriteFile(destination, raw, 0o444); err != nil {
		t.Fatalf("write destination: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chmod(destination, 0o644)
	})

	if err := copyPackageFileIfExists(source, destination); err != nil {
		t.Fatalf("copyPackageFileIfExists returned error: %v", err)
	}
	got, err := os.ReadFile(destination)
	if err != nil {
		t.Fatalf("read destination: %v", err)
	}
	if !bytes.Equal(got, raw) {
		t.Fatalf("destination content changed: %q", got)
	}
}

func TestPackageJSONUsesCanonicalGoBinaryAndPackageCleanupCommand(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("..", "..", "anvien", "package.json"))
	if err != nil {
		t.Fatalf("read package.json: %v", err)
	}

	var pkg struct {
		Bin     map[string]string `json:"bin"`
		Files   []string          `json:"files"`
		Scripts map[string]string `json:"scripts"`
	}
	if err := json.Unmarshal(raw, &pkg); err != nil {
		t.Fatalf("parse package.json: %v", err)
	}

	if got := pkg.Bin["anvien"]; got != "bin/anvien.exe" {
		t.Fatalf("pkg.bin.anvien = %q", got)
	}
	if len(pkg.Bin) != 1 {
		t.Fatalf("pkg.bin should expose only anvien: %#v", pkg.Bin)
	}
	for _, want := range []string{"bin", "go-src"} {
		if !containsString(pkg.Files, want) {
			t.Fatalf("pkg.files missing %q: %#v", want, pkg.Files)
		}
	}
	if containsString(pkg.Files, "skills") {
		t.Fatalf("pkg.files should not ship package-root skills: %#v", pkg.Files)
	}
	for _, retiredFile := range []string{"dist", "scripts", "vendor"} {
		if containsString(pkg.Files, retiredFile) {
			t.Fatalf("pkg.files still ships retired legacy path %q: %#v", retiredFile, pkg.Files)
		}
	}
	if got := pkg.Scripts["build"]; got != "go run ../cmd/anvien package build-runtime" {
		t.Fatalf("pkg.scripts.build = %q", got)
	}
	if got := pkg.Scripts["prepack"]; !strings.Contains(got, "package prepare-go-source") {
		t.Fatalf("pkg.scripts.prepack missing Go prepare helper: %q", got)
	}
	if got := pkg.Scripts["postinstall"]; !strings.Contains(got, "package build-runtime") || strings.Contains(got, "scripts/") {
		t.Fatalf("pkg.scripts.postinstall should use Go package lifecycle helpers only: %q", got)
	}
	if got := pkg.Scripts["postpack"]; !strings.Contains(got, "package clean-go-source") || strings.Contains(got, "scripts/") {
		t.Fatalf("pkg.scripts.postpack = %q", got)
	}
	if strings.Contains(pkg.Scripts["postpack"], "clean-go-source-package.cjs") {
		t.Fatalf("pkg.scripts.postpack still references deleted CJS cleanup helper: %q", pkg.Scripts["postpack"])
	}
	for _, script := range []string{"test", "build"} {
		if pkg.Scripts[script] == "" {
			t.Fatalf("pkg.scripts missing %q: %#v", script, pkg.Scripts)
		}
	}
	for _, retired := range []string{"test:integration", "test:unit", "test:watch", "test:coverage"} {
		if pkg.Scripts[retired] != "" {
			t.Fatalf("pkg.scripts still contains retired legacy Vitest script %q: %q", retired, pkg.Scripts[retired])
		}
	}
	if strings.Contains(strings.ToLower(pkg.Scripts["test"]), "vitest") {
		t.Fatalf("pkg.scripts.test still references legacy Vitest: %q", pkg.Scripts["test"])
	}
	for _, script := range pkg.Scripts {
		for _, retired := range []string{"build.js", "build-go-runtime.cjs", "prepare-go-source-package.cjs", "patch-tree-sitter-swift.cjs", "build-tree-sitter-proto.cjs", "dev:ts-baseline"} {
			if strings.Contains(script, retired) {
				t.Fatalf("pkg.scripts still references retired legacy helper %q in %q", retired, script)
			}
		}
	}
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func writePackageTestFile(t *testing.T, root, rel, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
