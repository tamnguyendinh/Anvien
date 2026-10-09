// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build !linux

package cuda

// HasNvidiaGPU always returns false on non-Linux platforms where GPU auto-detection is not yet implemented.
func HasNvidiaGPU() bool {
	return false
}

// CheckCUDAAndCUDNN is a no-op on non-Linux platforms.
func CheckCUDAAndCUDNN() error {
	return nil
}

// IsCUDALibraryAvailable checks if an ONNX Runtime CUDA provider shared library is present in the directory.
func IsCUDALibraryAvailable(dir string) bool {
	return false
}
