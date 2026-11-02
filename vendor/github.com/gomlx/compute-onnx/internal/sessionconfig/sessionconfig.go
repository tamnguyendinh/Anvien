// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package sessionconfig

// Config holds runtime session options (such as threading and memory settings)
// configured via the backend configuration string.
type Config struct {
	IntraOpNumThreads      int    // Number of threads for intra-op parallelism (-1: unset/default)
	InterOpNumThreads      int    // Number of threads for inter-op parallelism (-1: unset/default)
	CpuMemArena            *bool  // Enable or disable CPU memory arena (nil: unset/default)
	MemPattern             *bool  // Enable or disable memory pattern optimization (nil: unset/default)
	ExecutionMode          string // "parallel" or "sequential" ("": unset/default)
	GraphOptimizationLevel int    // Graph optimization level (0=disable, 1=basic, 2=extended, 99=all, -1: unset/default)
	SessionClones          int    // Max concurrent session clones in the pool (<=0: unset/default)
}

// Default returns a Config with all options unset (indicating ONNX Runtime defaults).
func Default() Config {
	return Config{
		IntraOpNumThreads:      -1,
		InterOpNumThreads:      -1,
		GraphOptimizationLevel: -1,
		SessionClones:          -1,
	}
}
