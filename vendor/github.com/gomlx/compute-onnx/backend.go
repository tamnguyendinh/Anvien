// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package onnxbackend

import (
	"fmt"
	"io"
	"strings"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	"github.com/gomlx/compute-onnx/internal/graph"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/pkg/errors"
)

const (
	BackendName = "onnxruntime"
)

// Common type aliases for graph builder IR types.
type Builder = graph.Builder
type Function = graph.Function
type Node = graph.Node

// MakeScalar constructs a 0D scalar constant tensor in the given function.
func MakeScalar(f *graph.Function, value any, dtype dtypes.DType) (compute.Value, error) {
	return graph.MakeScalar(f, value, dtype)
}

// Backend represents an ONNX Runtime backed [compute.Backend].
type Backend struct {
	config             string
	version            string
	executionProvider  executionprovider.Type // Execution provider: CPU, CUDA, MIGraphX, WASM, WebGPU, WebNN, WebGL.
	migraphxCacheDir   string                 // Native MIGraphX compiled-program cache directory ("migraphx_cache_dir" config key); empty disables caching.
	webVersion         string
	logSeverity        int
	enableGraphCapture bool
	hasFloat64         bool
	hasFloat16         bool
	hasBFloat16        bool
	isFinalized        bool
	keepModelProto     bool
}

// SetKeepModelProto controls whether compiled Executable instances retain the graph *onnx.ModelProto.
// When set to true, executables compiled by this backend store their ONNX ModelProto struct,
// allowing the computation graph to be exported via SaveModel.
func (b *Backend) SetKeepModelProto(keep bool) {
	b.keepModelProto = keep
}

// KeepModelProto returns whether compiled Executable instances retain their graph *onnx.ModelProto.
func (b *Backend) KeepModelProto() bool {
	return b.keepModelProto
}

// LogSeverity returns the configured log severity level (0=verbose, 1=info, 2=warning, 3=error, 4=fatal, or -1 if not set).
func (b *Backend) LogSeverity() int {
	return b.logSeverity
}

// ExecutionProvider returns the execution provider configured for the backend.
func (b *Backend) ExecutionProvider() executionprovider.Type {
	return b.executionProvider
}

// WebVersion returns the configured web version string (e.g. "dev", "@latest", "1.27", or "").
func (b *Backend) WebVersion() string {
	return b.webVersion
}

var _ compute.Backend = (*Backend)(nil)
var _ compute.DataInterface = (*Backend)(nil)

func init() {
	if IsSupportedPlatform() {
		compute.Register(BackendName, New)
		compute.Register("onnx", New)
	}
}

// ParseGOMLXBackendEnv parses a GOMLX_BACKEND environment variable string (e.g. "onnx:cpu", "onnx", "onnxruntime:cuda,log=2").
// It verifies that the backend selection is "onnx" or "onnxruntime", strips the backend name prefix, and returns the configuration options string.
func ParseGOMLXBackendEnv(envVal string) (string, error) {
	envVal = strings.TrimSpace(envVal)
	if envVal == "" {
		return "", nil
	}
	parts := strings.SplitN(envVal, ":", 2)
	backendName := strings.ToLower(strings.TrimSpace(parts[0]))
	if backendName != "onnx" && backendName != "onnxruntime" {
		return "", errors.Errorf("invalid backend selection %q in GOMLX_BACKEND: expected backend 'onnx' or 'onnxruntime'", parts[0])
	}
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1]), nil
	}
	return "", nil
}

func (b *Backend) Name() string {
	return BackendName
}

func (b *Backend) Config() string {
	return b.config
}

func (b *Backend) String() string {
	return b.Name()
}

// Version returns the version of the underlying ONNX Runtime engine.
func (b *Backend) Version() string {
	return b.version
}

func (b *Backend) Description() string {
	verStr := ""
	if b.version != "" {
		verStr = fmt.Sprintf(" v%s", b.version)
	}
	switch b.executionProvider {
	case executionprovider.CUDA:
		return fmt.Sprintf("ONNX Runtime%s (CUDA GPU) compute backend for GoMLX", verStr)
	case executionprovider.MIGraphX:
		return fmt.Sprintf("ONNX Runtime%s (MIGraphX GPU) compute backend for GoMLX", verStr)
	case executionprovider.WASM, executionprovider.WebGPU, executionprovider.WebNN, executionprovider.WebGL:
		return fmt.Sprintf("ONNX Runtime Web%s (%s) compute backend for GoMLX", verStr, b.executionProvider.String())
	}
	return fmt.Sprintf("ONNX Runtime%s (CPU) compute backend for GoMLX", verStr)
}

func (b *Backend) NumDevices() int {
	return 1
}

func (b *Backend) DeviceDescription(deviceNum compute.DeviceNum) string {
	switch b.executionProvider {
	case executionprovider.CUDA:
		return "CUDA GPU (ONNX Runtime Default Device)"
	case executionprovider.MIGraphX:
		return "MIGraphX GPU (ONNX Runtime Default Device)"
	case executionprovider.WASM, executionprovider.WebGPU, executionprovider.WebNN, executionprovider.WebGL:
		return fmt.Sprintf("Web (%s) Default Device", b.executionProvider.String())
	}
	return "CPU (ONNX Runtime Default Device)"
}

func (b *Backend) Capabilities() compute.Capabilities {
	caps := compute.Capabilities{
		Operations:                  graph.GetSupportedOps(),
		DTypes:                      make(map[dtypes.DType]bool),
		PreferConstantsForVariables: true,
		DynamicAxes:                 true,
		DynamicShapes:               compute.DynamicShapesNative,
	}
	if b.executionProvider == executionprovider.MIGraphX {
		// The MIGraphX execution provider reliably supports only a subset of dtypes;
		// others (e.g. float64, uint8) either fall back producing incorrect results
		// or crash inside MIGraphX. Restrict the advertised capabilities accordingly,
		// so graphs using unsupported dtypes fail cleanly at compilation time.
		caps.DynamicAxes = false
		caps.DTypes[dtypes.Float32] = true
		caps.DTypes[dtypes.Int32] = true
		caps.DTypes[dtypes.Int64] = true
		return caps
	}
	caps.DTypes[dtypes.Float32] = true
	caps.DTypes[dtypes.Float64] = b.hasFloat64
	caps.DTypes[dtypes.Float16] = b.hasFloat16
	caps.DTypes[dtypes.BFloat16] = b.hasBFloat16
	caps.DTypes[dtypes.Int32] = true
	caps.DTypes[dtypes.Int64] = true
	caps.DTypes[dtypes.Bool] = true
	caps.DTypes[dtypes.Int8] = true
	caps.DTypes[dtypes.Uint8] = true
	caps.DTypes[dtypes.Int16] = true
	caps.DTypes[dtypes.Uint16] = true
	caps.DTypes[dtypes.Uint32] = true
	caps.DTypes[dtypes.Uint64] = true
	return caps
}

func (b *Backend) Finalize() {
	b.isFinalized = true
}

func (b *Backend) IsFinalized() bool {
	return b.isFinalized
}

func (b *Backend) Builder(name string) compute.Builder {
	gb := graph.NewBuilder(name, func(gb *graph.Builder) (compute.Executable, error) {
		compiled, err := graph.CompileToProto(gb)
		if err != nil {
			return nil, err
		}

		// Validate all used dtypes are supported by this backend
		caps := b.Capabilities()
		for dt := range compiled.UsedDTypes {
			if !caps.DTypes[dt] {
				return nil, errors.Wrapf(compute.ErrNotImplemented, "dtype %s is not supported by %s", dt, b.Description())
			}
		}

		return b.createExecutable(compiled.ModelBytes, compiled.InputNames, compiled.InputShapes, compiled.OutputNames, compiled.OutputShapes, compiled.Model)
	})
	gb.SetExecutionProvider(b.executionProvider)
	gb.SetLogSeverity(b.logSeverity)
	return gb
}

// onnxExecutable is an internal interface for Executable instances that have backend and ModelProto methods.
type onnxExecutable interface {
	compute.Executable
	Backend() compute.Backend
	ModelProto() *onnx.ModelProto
}

// SaveModel exports the computation graph associated with the given executable as an ONNX model file/stream.
// It verifies that both backend and executable belong to the ONNX backend package.
// If inputNames or outputNames are provided (non-empty), they rename the graph's inputs and outputs
// and update all corresponding internal node edge references before saving.
// If inputNames or outputNames are nil or empty, default or pre-existing graph names (e.g., "arg_0", "output_0") are retained.
//
// Note: The backend must have had KeepModelProto set to true prior to compiling the executable,
// otherwise SaveModel will return an error indicating that the graph proto was discarded.
func SaveModel(backend compute.Backend, executable compute.Executable, w io.Writer, inputNames []string, outputNames []string) error {
	onBackend, ok := backend.(*Backend)
	if !ok {
		return errors.New("SaveModel: backend is not an ONNX backend (*onnxbackend.Backend)")
	}

	onExec, ok := executable.(onnxExecutable)
	if !ok {
		return errors.New("SaveModel: executable is not an ONNX executable")
	}

	if onExec.Backend() != onBackend && onExec.Backend() != backend {
		return errors.New("SaveModel: executable was not created by the provided backend")
	}

	modelProto := onExec.ModelProto()
	if modelProto == nil {
		return errors.New("SaveModel: model ONNX proto not retained; backend.SetKeepModelProto(true) must be called prior to compilation")
	}

	modelBytes, err := graph.RemapAndMarshalModel(modelProto, inputNames, outputNames)
	if err != nil {
		return errors.Wrap(err, "SaveModel")
	}

	_, err = w.Write(modelBytes)
	if err != nil {
		return errors.Wrap(err, "SaveModel: failed to write model bytes to writer")
	}

	return nil
}

// LoadModel loads an ONNX model from an io.Reader and compiles it into a runnable compute.Executable.
// It parses input and output tensor shapes and data types directly from the ONNX ModelProto
// and creates an ONNX Runtime session ready for execution.
func LoadModel(backend compute.Backend, r io.Reader) (compute.Executable, error) {
	onBackend, ok := backend.(*Backend)
	if !ok {
		return nil, errors.New("LoadModel: backend is not an ONNX backend (*onnxbackend.Backend)")
	}

	modelProto, modelBytes, inputNames, inputShapes, outputNames, outputShapes, err := graph.ParseModelProto(r)
	if err != nil {
		return nil, errors.Wrap(err, "LoadModel")
	}

	return onBackend.createExecutable(modelBytes, inputNames, inputShapes, outputNames, outputShapes, modelProto)
}
