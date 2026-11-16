// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

// Package executionprovider defines the native GPU execution providers
// supported by the ONNX Runtime backend.
package executionprovider

import (
	"fmt"
	"strings"
)

// Type identifies the ONNX Runtime execution provider used by a backend.
type Type int

const (
	// CPU runs on CPU only (the zero value).
	CPU Type = iota
	// CUDA runs on an NVIDIA GPU via the CUDA execution provider.
	CUDA
	// MIGraphX runs on an AMD GPU via the MIGraphX execution provider. It replaces the older ROCm EP.
	MIGraphX
	// WASM runs WebAssembly CPU execution provider in browser environments.
	WASM
	// WebGPU runs on WebGPU execution provider in browser environments.
	WebGPU
	// WebNN runs on WebNN execution provider in browser environments.
	WebNN
	// WebGL runs on WebGL execution provider in browser environments.
	WebGL
)

// String returns the canonical name of the execution provider ("cpu", "cuda", "migraphx", "wasm", "webgpu", "webnn", or "webgl").
func (t Type) String() string {
	switch t {
	case CUDA:
		return "cuda"
	case MIGraphX:
		return "migraphx"
	case WASM:
		return "wasm"
	case WebGPU:
		return "webgpu"
	case WebNN:
		return "webnn"
	case WebGL:
		return "webgl"
	default:
		return "cpu"
	}
}

// IsWeb returns true if the execution provider is a web/browser execution provider (WASM, WebGPU, WebNN, WebGL).
func (t Type) IsWeb() bool {
	switch t {
	case WASM, WebGPU, WebNN, WebGL:
		return true
	default:
		return false
	}
}

// Parse maps a configuration token to an executionprovider.Type.
// It accepts the same aliases as the backend config string: "cpu"/"" for CPU,
// "cuda"/"gpu" for CUDA, "migraphx"/"rocm"/"amd" for MIGraphX,
// "wasm" for WASM, "webgpu" for WebGPU, "webnn" for WebNN, and "webgl" for WebGL.
func Parse(s string) (Type, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "", "cpu":
		return CPU, nil
	case "cuda", "gpu":
		return CUDA, nil
	case "migraphx", "rocm", "amd":
		return MIGraphX, nil
	case "wasm":
		return WASM, nil
	case "webgpu":
		return WebGPU, nil
	case "webnn":
		return WebNN, nil
	case "webgl":
		return WebGL, nil
	default:
		return CPU, fmt.Errorf("invalid execution provider %q: expected \"cpu\", \"cuda\", \"migraphx\", \"wasm\", \"webgpu\", \"webnn\", or \"webgl\"", s)
	}
}
