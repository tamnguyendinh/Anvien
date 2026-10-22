// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build linux

package rocm

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// HasAMDGPU tries to guess if there is an actual (discrete) AMD GPU installed,
// suitable for ROCm execution. APUs (shared memory with the CPU) are excluded.
func HasAMDGPU() bool {
	return hasAMDGPU()
}

var hasAMDGPU = sync.OnceValue[bool](func() bool {
	// The AMD ROCm compute device node must be present.
	if _, err := os.Stat("/dev/kfd"); err != nil {
		return false
	}

	output, err := runRocminfo()
	if err != nil || output == "" {
		// Can't run rocminfo: assume there is no usable AMD GPU, otherwise
		// auto-installation of the ROCm runtime would be triggered and fail.
		return false
	}
	return rocminfoHasDiscreteGPU(output)
})

func runRocminfo() (string, error) {
	binPath := ""
	if _, err := exec.LookPath("rocminfo"); err == nil {
		binPath = "rocminfo"
	} else if p := filepath.Join(rocmInstallDir(), "bin", "rocminfo"); fileExists(p) {
		binPath = p
	} else {
		return "", errors.New("rocminfo not found")
	}
	cmd := exec.Command(binPath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// rocminfoHasDiscreteGPU parses `rocminfo` output and returns whether a discrete
// GPU agent (gfx*, non-APU) is found. Agents are separated by "*******" separators.
func rocminfoHasDiscreteGPU(output string) bool {
	for _, block := range strings.Split(output, "*******") {
		name := rocminfoField(block, "Name:")
		if !strings.HasPrefix(name, "gfx") {
			continue // Skip CPU agents.
		}
		if rocminfoField(block, "Memory Properties:") != "APU" {
			return true
		}
	}
	return false
}

// rocminfoField returns the first field value in the block with the given key prefix.
func rocminfoField(block, key string) string {
	for _, line := range strings.Split(block, "\n") {
		trimmed := strings.TrimLeft(line, " \t")
		if value, ok := strings.CutPrefix(trimmed, key); ok {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

// rocmInstallDir guesses where ROCm is installed: $ROCM_PATH, or inferred from
// rocminfo in PATH, falling back to /opt/rocm.
func rocmInstallDir() string {
	if dir := os.Getenv("ROCM_PATH"); dir != "" {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir
		}
	}
	if p, err := exec.LookPath("rocminfo"); err == nil {
		if real, err := filepath.EvalSymlinks(p); err == nil {
			p = real
		}
		// <root>/bin/rocminfo -> <root>
		return filepath.Dir(filepath.Dir(p))
	}
	return "/opt/rocm"
}

// GetROCmVersion returns the installed ROCm version string (e.g. "7.2.4"), or "" if not found.
// It reads <rocm-root>/.info/version (or <rocm-root>/lib/.info/version on older installs).
func GetROCmVersion() string {
	root := rocmInstallDir()
	for _, p := range []string{
		filepath.Join(root, ".info", "version"),
		filepath.Join(root, "lib", ".info", "version"),
	} {
		if data, err := os.ReadFile(p); err == nil {
			return strings.TrimSpace(string(data))
		}
	}
	return ""
}

// GetROCMDirectory returns the detected ROCm installation directory (e.g. "/opt/rocm").
func GetROCMDirectory() string {
	return rocmInstallDir()
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

// CheckROCmAndMIGraphX verifies that a discrete AMD GPU, the HIP runtime
// (libamdhip64) and MIGraphX libraries are available on this system.
func CheckROCmAndMIGraphX() error {
	if !HasAMDGPU() {
		return errors.New("no discrete AMD GPU found (/dev/kfd is missing or no gfx device reported by rocminfo)")
	}

	if !hasHIPLibrary() {
		return errors.Errorf(`HIP runtime library libamdhip64.so not found.
Make sure the AMD Radeon Software for Linux (with ROCm) is installed, e.g.:
  - Debian/Ubuntu: sudo apt install rocm-hip-libraries  # and add /opt/rocm/lib to LD_LIBRARY_PATH
  - See https://rocm.docs.amd.com/projects/install-on-linux/en/latest/install/native-install/index.html`)
	}

	if !hasMigraphxLibrary() {
		return errors.Errorf(`MIGraphX libraries not found.
The ONNX Runtime MIGraphX execution provider requires MIGraphX to be installed:
  - Debian/Ubuntu: sudo apt install migraphx migraphx-dev half
  - Verify with: dpkg -l | grep migraphx
  - See https://rocm.docs.amd.com/projects/radeon-ryzen/en/latest/docs/install/installrad/native_linux/install-migraphx.html`)
	}
	return nil
}

// hasHIPLibrary checks whether libamdhip64 can be found in the standard search paths.
func hasHIPLibrary() bool {
	dirs := librarySearchPaths()
	for _, dir := range dirs {
		matches, err := filepath.Glob(filepath.Join(dir, "libamdhip64.so*"))
		if err == nil && len(matches) > 0 {
			return true
		}
	}
	return false
}

// hasMigraphxLibrary checks whether MIGraphX shared libraries are present.
func hasMigraphxLibrary() bool {
	dirs := librarySearchPaths()
	for _, dir := range dirs {
		// Core MIGraphX libs live directly under <root>/lib (e.g. libmigraphx_c.so),
		// with the rest under <root>/lib/migraphx/lib.
		matches, err := filepath.Glob(filepath.Join(dir, "libmigraphx*.so*"))
		if err == nil && len(matches) > 0 {
			return true
		}
	}
	return false
}

// librarySearchPaths returns candidate directories for ROCm/HIP libraries:
// LD_LIBRARY_PATH entries, $ROCM_PATH/lib, /opt/rocm*/lib, plus ldconfig cache check.
func librarySearchPaths() []string {
	var dirs []string
	seen := make(map[string]bool)
	add := func(p string) {
		p = strings.TrimSpace(p)
		if p == "" || seen[p] {
			return
		}
		seen[p] = true
		dirs = append(dirs, p)
	}

	for _, entry := range strings.Split(os.Getenv("LD_LIBRARY_PATH"), string(os.PathListSeparator)) {
		add(entry)
	}
	root := rocmInstallDir()
	add(filepath.Join(root, "lib"))
	if matches, err := filepath.Glob("/opt/rocm*/lib"); err == nil {
		for _, m := range matches {
			add(m)
		}
	}
	add("/usr/lib/x86_64-linux-gnu")
	add("/usr/lib64")
	add("/usr/lib")
	return dirs
}

// HasMigraphxExecutionProvider checks if an ONNX Runtime MIGraphX provider shared library is present in the directory.
func HasMigraphxExecutionProvider(dir string) bool {
	if _, err := os.Stat(filepath.Join(dir, "libonnxruntime_providers_migraphx.so")); err == nil {
		return true
	}
	if _, err := os.Stat(filepath.Join(dir, "onnxruntime_providers_migraphx.dll")); err == nil {
		return true
	}
	return false
}
