// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build !linux

package rocm

// HasAMDGPU always returns false on non-Linux platforms where ROCm auto-detection is not yet implemented.
func HasAMDGPU() bool {
	return false
}

// CheckROCmAndMIGraphX is a no-op on non-Linux platforms.
func CheckROCmAndMIGraphX() error {
	return nil
}

// GetROCmVersion always returns "" on non-Linux platforms.
func GetROCmVersion() string {
	return ""
}

// GetROCMDirectory always returns "" on non-Linux platforms.
func GetROCMDirectory() string {
	return ""
}

// HasMigraphxExecutionProvider always returns false on non-Linux platforms.
func HasMigraphxExecutionProvider(dir string) bool {
	return false
}
