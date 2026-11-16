// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package web

import (
	"syscall/js"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

// Executable implements [compute.Executable] for ONNX Runtime Web.
type Executable struct {
	backend      compute.Backend
	session      *Session
	inputNames   []string
	inputShapes  []shapes.Shape
	outputNames  []string
	outputShapes []shapes.Shape
	modelProto   *onnx.ModelProto

	// WebGPU optimization fields
	isWebGPU       bool
	gpuDevice      js.Value
	gpuInputBuffers []js.Value
	gpuInputTensors []js.Value
}

var _ compute.Executable = (*Executable)(nil)

// NewExecutable creates a new web Executable.
func NewExecutable(backend compute.Backend, session *Session,
	inputNames []string, inputShapes []shapes.Shape,
	outputNames []string, outputShapes []shapes.Shape,
	modelProto *onnx.ModelProto) *Executable {

	exec := &Executable{
		backend:      backend,
		session:      session,
		inputNames:   inputNames,
		inputShapes:  inputShapes,
		outputNames:  outputNames,
		outputShapes: outputShapes,
		modelProto:   modelProto,
	}

	// If running under WebGPU, pre-allocate static input GPU buffers
	if session != nil && session.ep == executionprovider.WebGPU {
		dev, errDev := GetWebGPUDevice()
		if errDev == nil && !dev.IsUndefined() && !dev.IsNull() {
			exec.isWebGPU = true
			exec.gpuDevice = dev
			exec.initWebGPUStaticBuffers()
		}
	}

	return exec
}

func (e *Executable) initWebGPUStaticBuffers() {
	global := js.Global()
	ortVal := global.Get("ort")
	if ortVal.IsUndefined() || ortVal.IsNull() {
		return
	}
	tensorCtor := ortVal.Get("Tensor")
	if tensorCtor.IsUndefined() || tensorCtor.IsNull() {
		return
	}
	fromGpuBuffer := tensorCtor.Get("fromGpuBuffer")
	if fromGpuBuffer.IsUndefined() || fromGpuBuffer.IsNull() {
		return
	}

	// 1. Static input buffers
	e.gpuInputBuffers = make([]js.Value, len(e.inputNames))
	e.gpuInputTensors = make([]js.Value, len(e.inputNames))

	for i, sh := range e.inputShapes {
		if sh.IsDynamic() {
			continue
		}
		byteSize := sh.Size() * sh.DType.Size()
		if byteSize == 0 || byteSize%4 != 0 {
			continue
		}
		// WebGPU buffer sizes must be aligned to 16 bytes for storage buffers
		alignedSize := (byteSize + 15) &^ 15

		bufDesc := global.Get("Object").New()
		bufDesc.Set("size", alignedSize)
		// usage: STORAGE | COPY_DST | COPY_SRC
		bufDesc.Set("usage", 0x0080|0x0008|0x0004)

		gpuBuf := e.gpuDevice.Call("createBuffer", bufDesc)
		e.gpuInputBuffers[i] = gpuBuf

		ortType, errType := ORTTypeString(sh.DType)
		if errType != nil {
			continue
		}
		opts := global.Get("Object").New()
		opts.Set("dataType", ortType)
		dimsArray := global.Get("Array").New()
		for _, d := range sh.Dimensions {
			dimsArray.Call("push", d)
		}
		opts.Set("dims", dimsArray)

		// ORT Web fromGpuBuffer supports float32, float16, int32, int64, uint32, bool (float64/bfloat16/uint64/int8/uint8/int16/uint16 are unsupported)
		if sh.DType == dtypes.Float64 || sh.DType == dtypes.BFloat16 || sh.DType == dtypes.Uint64 || sh.DType == dtypes.Int8 || sh.DType == dtypes.Uint8 || sh.DType == dtypes.Int16 || sh.DType == dtypes.Uint16 {
			continue
		}

		gpuTensor := tensorCtor.Call("fromGpuBuffer", gpuBuf, opts)
		e.gpuInputTensors[i] = gpuTensor
	}
}

// Backend returns the parent compute.Backend.
func (e *Executable) Backend() compute.Backend {
	return e.backend
}

// ModelProto returns the ONNX ModelProto struct if retain was enabled during compilation, or nil otherwise.
func (e *Executable) ModelProto() *onnx.ModelProto {
	return e.modelProto
}

// Inputs returns graph input names and shapes.
func (e *Executable) Inputs() (names []string, inputShapes []shapes.Shape) {
	return e.inputNames, e.inputShapes
}

// Outputs returns graph output shapes.
func (e *Executable) Outputs() (outputShapes []shapes.Shape) {
	return e.outputShapes
}

// Finalize releases the underlying session and GPU buffers.
func (e *Executable) Finalize() {
	if e.session != nil {
		_ = e.session.Destroy()
		e.session = nil
	}
	for _, buf := range e.gpuInputBuffers {
		if !buf.IsUndefined() && !buf.IsNull() {
			buf.Call("destroy")
		}
	}
	e.gpuInputBuffers = nil
	e.gpuInputTensors = nil
}

// Execute runs the executable with the provided inputs.
func (e *Executable) Execute(inputs []compute.Buffer, donate []bool, defaultDevice compute.DeviceNum) ([]compute.Buffer, error) {
	if e.session == nil {
		return nil, errors.New("cannot execute finalized executable")
	}

	isDummyInput := len(e.inputNames) == 1 && e.inputNames[0] == "dummy_input"
	expectedInputs := len(e.inputNames)
	if isDummyInput {
		expectedInputs = 0
	}

	if len(inputs) != expectedInputs {
		return nil, errors.Errorf("expected %d inputs, got %d", expectedInputs, len(inputs))
	}

	global := js.Global()
	feeds := global.Get("Object").New()

	if isDummyInput {
		typedArray, err := ConvertSliceToTypedArray([]int32{0}, dtypes.Int32)
		if err != nil {
			return nil, err
		}
		dummyTensor, err := CreateJSTensor(dtypes.Int32, []int{1}, typedArray)
		if err != nil {
			return nil, err
		}
		feeds.Set("dummy_input", dummyTensor)
	} else {
		for i, inp := range inputs {
			buf, ok := inp.(*Buffer)
			if !ok {
				return nil, errors.Errorf("input %d is not a valid web Buffer", i)
			}
			if buf.wrapper == nil || buf.wrapper.jsTensor.IsUndefined() {
				return nil, errors.Errorf("input %d is finalized", i)
			}

			// If WebGPU pre-allocated GPU buffer is available, upload directly via device.queue.writeBuffer
			if e.isWebGPU && i < len(e.gpuInputBuffers) && !e.gpuInputBuffers[i].IsUndefined() && !e.gpuInputBuffers[i].IsNull() &&
				i < len(e.gpuInputTensors) && !e.gpuInputTensors[i].IsUndefined() && !e.gpuInputTensors[i].IsNull() &&
				buf.shape.Equal(e.inputShapes[i]) {
				dataProp := buf.wrapper.jsTensor.Get("data")
				if !dataProp.IsUndefined() && !dataProp.IsNull() {
					queue := e.gpuDevice.Get("queue")
					queue.Call("writeBuffer", e.gpuInputBuffers[i], 0, dataProp)
					feeds.Set(e.inputNames[i], e.gpuInputTensors[i])
					continue
				}
			}

			feeds.Set(e.inputNames[i], buf.wrapper.jsTensor)
		}
	}

	results, err := e.session.Run(feeds)
	if err != nil {
		return nil, err
	}

	outBuffers := make([]compute.Buffer, len(e.outputNames))
	for i, name := range e.outputNames {
		outTensor := results.Get(name)
		if outTensor.IsUndefined() || outTensor.IsNull() {
			return nil, errors.Errorf("output tensor %q missing from session run results", name)
		}

		sh := e.outputShapes[i]
		actualShape := sh
		if sh.IsDynamic() {
			dimsVal := outTensor.Get("dims")
			numDims := dimsVal.Length()
			concreteDims := make([]int, numDims)
			for d := 0; d < numDims; d++ {
				concreteDims[d] = dimsVal.Index(d).Int()
			}
			actualShape.Dimensions = concreteDims
		}

		wrapper := NewWebTensorWrapper(outTensor, actualShape)
		outBuffers[i] = NewBuffer(e.backend, wrapper, actualShape, defaultDevice, false, e)
	}

	// Donate inputs if requested
	for i, inp := range inputs {
		if i < len(donate) && donate[i] {
			_ = inp.Finalize()
		}
	}

	return outBuffers, nil
}
