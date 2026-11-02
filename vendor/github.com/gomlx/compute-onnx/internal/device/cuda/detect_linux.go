// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build linux

package cuda

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// HasNvidiaGPU tries to guess if there is an actual Nvidia GPU installed.
func HasNvidiaGPU() bool {
	return hasNvidiaGPU()
}

var hasNvidiaGPU = sync.OnceValue[bool](func() bool {
	matches, err := filepath.Glob("/dev/nvidia*")
	if err == nil && len(matches) > 0 {
		return true
	}

	// Execute the nvidia-smi command if present
	_, lookErr := exec.LookPath("nvidia-smi")
	if lookErr != nil {
		return false
	}
	cmd := exec.Command("nvidia-smi")
	output, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		return false
	}
	return strings.Contains(string(output), "NVIDIA-SMI")
})

// GetCUDALibrarySearchPaths returns candidate directories for CUDA and cuDNN libraries,
// along with boolean flags indicating if ldconfig already confirmed libcudart or libcudnn.
func GetCUDALibrarySearchPaths() (paths []string, ldconfigCuda bool, ldconfigCudnn bool) {
	seen := make(map[string]bool)
	addPath := func(p string) {
		p = strings.TrimSpace(p)
		if p == "" {
			return
		}
		p = filepath.Clean(p)
		if !seen[p] {
			seen[p] = true
			paths = append(paths, p)
		}
	}

	// 1. Environment variables
	// LD_LIBRARY_PATH
	if ldLibPath := os.Getenv("LD_LIBRARY_PATH"); ldLibPath != "" {
		for _, dir := range filepath.SplitList(ldLibPath) {
			addPath(dir)
		}
	}

	// CUDA_PATH, CUDA_HOME, CUDA_ROOT
	for _, envVar := range []string{"CUDA_PATH", "CUDA_HOME", "CUDA_ROOT"} {
		if val := os.Getenv(envVar); val != "" {
			addPath(val)
			addPath(filepath.Join(val, "lib64"))
			addPath(filepath.Join(val, "lib"))
			addPath(filepath.Join(val, "targets", "x86_64-linux", "lib"))
		}
	}

	// CUDNN_PATH, CUDNN_HOME, CUDNN_ROOT
	for _, envVar := range []string{"CUDNN_PATH", "CUDNN_HOME", "CUDNN_ROOT"} {
		if val := os.Getenv(envVar); val != "" {
			addPath(val)
			addPath(filepath.Join(val, "lib64"))
			addPath(filepath.Join(val, "lib"))
			addPath(filepath.Join(val, "targets", "x86_64-linux", "lib"))
		}
	}

	// CONDA_PREFIX, VIRTUAL_ENV
	for _, envVar := range []string{"CONDA_PREFIX", "VIRTUAL_ENV"} {
		if val := os.Getenv(envVar); val != "" {
			addPath(filepath.Join(val, "lib"))
			addPath(filepath.Join(val, "lib64"))
		}
	}

	// 2. Binary paths: Look for nvcc or PATH entries
	if nvccPath, err := exec.LookPath("nvcc"); err == nil {
		binDir := filepath.Dir(nvccPath)
		baseDir := filepath.Dir(binDir)
		addPath(filepath.Join(baseDir, "lib64"))
		addPath(filepath.Join(baseDir, "lib"))
		addPath(filepath.Join(baseDir, "targets", "x86_64-linux", "lib"))
	}

	for _, p := range filepath.SplitList(os.Getenv("PATH")) {
		if strings.Contains(p, "cuda") {
			baseDir := filepath.Dir(p)
			addPath(filepath.Join(baseDir, "lib64"))
			addPath(filepath.Join(baseDir, "lib"))
			addPath(filepath.Join(baseDir, "targets", "x86_64-linux", "lib"))
		}
	}

	// 3. System dynamic linker configuration (ldconfig -p)
	if ldconfigPath, err := exec.LookPath("ldconfig"); err == nil {
		cmd := exec.Command(ldconfigPath, "-p")
		out, err := cmd.Output()
		if err == nil {
			for line := range bytes.SplitSeq(out, []byte("\n")) {
				lineStr := strings.TrimSpace(string(line))
				if idx := strings.LastIndex(lineStr, "=> "); idx != -1 {
					libPath := strings.TrimSpace(lineStr[idx+3:])
					libName := filepath.Base(libPath)
					if strings.HasPrefix(libName, "libcudart.so") {
						ldconfigCuda = true
						addPath(filepath.Dir(libPath))
					} else if strings.HasPrefix(libName, "libcudnn.so") {
						ldconfigCudnn = true
						addPath(filepath.Dir(libPath))
					}
				}
			}
		}
	}

	// 4. Standard Linux / CUDA installation directories
	standardPaths := []string{
		"/usr/local/cuda/lib64",
		"/usr/local/cuda/lib",
		"/usr/local/cuda/targets/x86_64-linux/lib",
		"/usr/lib/x86_64-linux-gnu",
		"/usr/lib/wsl/lib",
		"/usr/lib64",
		"/usr/lib",
	}
	for _, p := range standardPaths {
		addPath(p)
	}

	// Versioned /usr/local/cuda-* paths (e.g. /usr/local/cuda-12.8/lib64)
	versionedMatches, _ := filepath.Glob("/usr/local/cuda-*/lib64")
	for _, p := range versionedMatches {
		addPath(p)
	}
	versionedTargets, _ := filepath.Glob("/usr/local/cuda-*/targets/x86_64-linux/lib")
	for _, p := range versionedTargets {
		addPath(p)
	}

	return paths, ldconfigCuda, ldconfigCudnn
}

// CheckCUDAAndCUDNN verifies that the CUDA runtime (libcudart.so) and cuDNN (libcudnn.so)
// libraries are present and accessible on the Linux system.
func CheckCUDAAndCUDNN() error {
	if runtime.GOOS != "linux" {
		return nil
	}

	paths, ldconfigCuda, ldconfigCudnn := GetCUDALibrarySearchPaths()

	hasCuda := ldconfigCuda
	hasCudnn := ldconfigCudnn

	for _, dir := range paths {
		if !hasCuda {
			if matches, _ := filepath.Glob(filepath.Join(dir, "libcudart.so*")); len(matches) > 0 {
				hasCuda = true
			}
		}
		if !hasCudnn {
			if matches, _ := filepath.Glob(filepath.Join(dir, "libcudnn.so*")); len(matches) > 0 {
				hasCudnn = true
			}
		}
		if hasCuda && hasCudnn {
			break
		}
	}

	if !hasCuda && !hasCudnn {
		return errors.Errorf("CUDA runtime (libcudart.so) and cuDNN (libcudnn.so) libraries not found in library search paths (%s). "+
			"Please ensure CUDA Toolkit and cuDNN are installed, and add their 'lib64' directory to LD_LIBRARY_PATH (e.g. export LD_LIBRARY_PATH=/usr/local/cuda/lib64:$LD_LIBRARY_PATH).",
			strings.Join(paths, ":"))
	}
	if !hasCuda {
		return errors.Errorf("CUDA runtime library (libcudart.so) not found in library search paths (%s). "+
			"Please ensure CUDA Toolkit is installed and add its 'lib64' directory to LD_LIBRARY_PATH.",
			strings.Join(paths, ":"))
	}
	if !hasCudnn {
		return errors.Errorf("cuDNN library (libcudnn.so) not found in library search paths (%s). "+
			"Please ensure NVIDIA cuDNN is installed and add its library directory to LD_LIBRARY_PATH.",
			strings.Join(paths, ":"))
	}

	return nil
}

// IsCUDALibraryAvailable checks if an ONNX Runtime CUDA provider shared library is present in the directory.
func IsCUDALibraryAvailable(dir string) bool {
	cudaLibPath := filepath.Join(dir, "libonnxruntime_providers_cuda.so")
	if _, err := os.Stat(cudaLibPath); err == nil {
		return true
	}
	cudaLibPathWin := filepath.Join(dir, "onnxruntime_providers_cuda.dll")
	if _, err := os.Stat(cudaLibPathWin); err == nil {
		return true
	}
	return false
}
