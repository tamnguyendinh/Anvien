// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build !js

package native

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"unsafe"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	ort "github.com/gomlx/compute-onnx/internal/ort"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapes"
	"github.com/gomlx/compute/support/humanize"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

// migraphxMaxWarmUpRuns caps the number of per-shape warm-up runs (see migraphxWorkaround).
const migraphxMaxWarmUpRuns = 8

// Executable implements [compute.Executable] for ONNX Runtime native execution (CPU & CUDA).
type Executable struct {
	backend           compute.Backend
	session           *ort.DynamicAdvancedSession
	inputNames        []string
	inputShapes       []shapes.Shape
	outputNames       []string
	outputShapes      []shapes.Shape
	reusableWrappers  []OrtTensorWrapper
	executionProvider executionprovider.Type // GPU execution provider: CUDA, MIGraphX, or CPU only.

	// warmedShapes tracks which input-shape signatures have already been through the
	// MIGraphX first-eval warm-up (see migraphxWorkaround). Only used when executionProvider == MIGraphX.
	warmedShapes map[string]bool

	// Pre-allocated slices reused across Execute calls (CPU path only, single-threaded).
	cachedOrtInputs  []ort.Value
	cachedOutWraps   []OrtTensorWrapper
	cachedOrtOutputs []ort.Value

	// Mutex to protect reusableWrappers for concurrent buffer finalization.
	mu sync.Mutex

	modelProto *onnx.ModelProto
}

var _ compute.Executable = (*Executable)(nil)

// NewExecutable creates a new Native Executable.
// The executionProvider is inherited from the Backend's session: CUDA, MIGraphX, or CPU only.
func NewExecutable(backend compute.Backend, session *ort.DynamicAdvancedSession,
	inputNames []string, inputShapes []shapes.Shape,
	outputNames []string, outputShapes []shapes.Shape,
	modelProto *onnx.ModelProto, executionProvider executionprovider.Type) *Executable {

	nInputs := len(inputNames)
	nOutputs := len(outputShapes)

	e := &Executable{
		backend:           backend,
		session:           session,
		inputNames:        inputNames,
		inputShapes:       inputShapes,
		outputNames:       outputNames,
		outputShapes:      outputShapes,
		cachedOrtInputs:   make([]ort.Value, nInputs),
		cachedOutWraps:    make([]OrtTensorWrapper, nOutputs),
		cachedOrtOutputs:  make([]ort.Value, nOutputs),
		modelProto:        modelProto,
		executionProvider: executionProvider,
	}
	if executionProvider == executionprovider.MIGraphX {
		e.warmedShapes = make(map[string]bool)
	}
	runtime.SetFinalizer(e, (*Executable).Finalize)
	return e
}

// Backend returns the parent compute.Backend.
func (e *Executable) Backend() compute.Backend {
	return e.backend
}

// ModelProto returns the ONNX ModelProto struct if retain was enabled during compilation, or nil otherwise.
func (e *Executable) ModelProto() *onnx.ModelProto {
	return e.modelProto
}

func (e *Executable) Finalize() {
	if e.session != nil {
		_ = e.session.Destroy()
		e.session = nil
	}
	e.mu.Lock()
	for _, w := range e.reusableWrappers {
		if w != nil {
			_ = w.Destroy()
		}
	}
	e.reusableWrappers = nil
	e.mu.Unlock()
}

func (e *Executable) Inputs() (names []string, inputShapes []shapes.Shape) {
	return e.inputNames, e.inputShapes
}

func (e *Executable) Outputs() (outputShapes []shapes.Shape) {
	return e.outputShapes
}

func wrapperMatchesShape(w OrtTensorWrapper, sh shapes.Shape) bool {
	if w.GetDType() != sh.DType {
		return false
	}
	rwShape := w.GetShape()
	if len(rwShape) != len(sh.Dimensions) {
		return false
	}
	for i, d := range rwShape {
		if int(d) != sh.Dimensions[i] {
			return false
		}
	}
	return true
}

func (e *Executable) matchesAnyOutput(dtype dtypes.DType, ortShape ort.Shape) bool {
	for _, sh := range e.outputShapes {
		if sh.DType != dtype {
			continue
		}
		if len(sh.Dimensions) != len(ortShape) {
			continue
		}
		match := true
		for i, d := range sh.Dimensions {
			if d != int(ortShape[i]) {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func (e *Executable) recycleWrapper(w OrtTensorWrapper) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.session == nil {
		_ = w.Destroy()
		return
	}

	if e.matchesAnyOutput(w.GetDType(), w.GetShape()) && len(e.reusableWrappers) < len(e.outputShapes)*2 {
		e.reusableWrappers = append(e.reusableWrappers, w)
	} else {
		_ = w.Destroy()
	}
}

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

	defer runtime.KeepAlive(inputs)
	if e.executionProvider == executionprovider.CUDA {
		// CUDA path: outputs are bound to device memory and stay GPU-resident.
		ortInputs := make([]ort.Value, len(e.inputNames))
		var dummyWrapper OrtTensorWrapper
		if isDummyInput {
			wrapper, err := NewOrtTensorWrapper(e.inputShapes[0], []int32{0})
			if err != nil {
				return nil, err
			}
			dummyWrapper = wrapper
			ortInputs[0] = wrapper.Value()
		} else {
			for i, inp := range inputs {
				buf, ok := inp.(*Buffer)
				if !ok {
					return nil, errors.Errorf("input %d is not a valid onnxruntime buffer", i)
				}
				if buf.wrapper == nil {
					return nil, errors.Errorf("input %d is finalized", i)
				}
				ortInputs[i] = buf.wrapper.Value()
			}
		}
		result, err := e.executeCUDA(ortInputs, inputs, donate, defaultDevice)
		if dummyWrapper != nil {
			_ = dummyWrapper.Destroy()
		}
		return result, err
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	// Populate shared cachedOrtInputs.
	var dummyInputWrapper OrtTensorWrapper
	if isDummyInput {
		wrapper, err := NewOrtTensorWrapper(e.inputShapes[0], []int32{0})
		if err != nil {
			return nil, err
		}
		dummyInputWrapper = wrapper
		e.cachedOrtInputs[0] = wrapper.Value()
	} else {
		for i, inp := range inputs {
			buf, ok := inp.(*Buffer)
			if !ok {
				return nil, errors.Errorf("input %d is not a valid onnxruntime buffer", i)
			}
			if buf.wrapper == nil {
				return nil, errors.Errorf("input %d is finalized", i)
			}
			e.cachedOrtInputs[i] = buf.wrapper.Value()
		}
	}

	if dummyInputWrapper != nil {
		defer func() {
			_ = dummyInputWrapper.Destroy()
		}()
	}

	return e.executeDefault(inputs, donate, defaultDevice)
}

func (e *Executable) executeCUDA(ortInputs []ort.Value, inputs []compute.Buffer, donate []bool, defaultDevice compute.DeviceNum) ([]compute.Buffer, error) {
	defer runtime.KeepAlive(inputs)
	defer runtime.KeepAlive(ortInputs)
	ioBinding, err := e.session.CreateIoBinding()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create IoBinding")
	}
	defer ioBinding.Destroy()

	cudaMemInfo, err := ort.NewCUDAMemoryInfo()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create CUDA MemoryInfo")
	}
	defer cudaMemInfo.Destroy()

	cInputNames := e.session.CInputNames()
	cOutputNames := e.session.COutputNames()

	for i := range ortInputs {
		if err := ioBinding.BindInput(cInputNames[i], ortInputs[i]); err != nil {
			return nil, errors.Wrapf(err, "failed to bind input %d", i)
		}
	}

	for i := range e.outputShapes {
		if err := ioBinding.BindOutputToDevice(cOutputNames[i], cudaMemInfo); err != nil {
			return nil, errors.Wrapf(err, "failed to bind output %d to CUDA", i)
		}
	}

	var start time.Time
	if klog.V(1).Enabled() {
		start = time.Now()
	}
	if klog.V(2).Enabled() {
		klog.Infof("Starting IoBinding.RunWithBinding on cuda (inputs=%d, outputs=%d)...", len(ortInputs), len(e.outputShapes))
	}
	e.mu.Lock()
	err = ioBinding.RunWithBinding()
	e.mu.Unlock()
	if klog.V(1).Enabled() {
		klog.Infof("Execution (CUDA) elapsed time: %s\n", humanize.Duration(time.Since(start)))
	}
	if err != nil {
		return nil, errors.Wrap(err, "onnxruntime CUDA execution failed")
	}

	outputValues, err := ioBinding.GetBoundOutputValues()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get bound output values")
	}

	outBuffers := make([]compute.Buffer, len(e.outputShapes))
	for i, sh := range e.outputShapes {
		val := ort.WrapRawOrtValue(outputValues[i])
		actualShape := sh
		if sh.IsDynamic() {
			boundShape, err := ort.GetOrtValueShape(outputValues[i])
			if err != nil {
				return nil, errors.Wrapf(err, "failed to get shape for output %d", i)
			}
			concreteDims := make([]int, len(boundShape))
			for k, dim := range boundShape {
				concreteDims[k] = int(dim)
			}
			actualShape.Dimensions = concreteDims
		}
		ortShape := ort.NewShape(toInt64s(actualShape.Dimensions)...)
		outBuffers[i] = &Buffer{
			backend:           e.backend,
			wrapper:           NewGpuTensorWrapper(val, ortShape, actualShape.DType),
			shape:             actualShape,
			device:            defaultDevice,
			executionProvider: executionprovider.CUDA,
		}
	}

	for i, inp := range inputs {
		if i < len(donate) && donate[i] {
			buf := inp.(*Buffer)
			if buf.wrapper != nil {
				_ = buf.wrapper.Destroy()
				buf.wrapper = nil
			}
		}
	}

	return outBuffers, nil
}

// executeDefault executes the compiled graph for any non-specialized execution provider
// (CPU and MIGraphX), using caller-preallocated or ONNX Runtime-allocated output buffers.
func (e *Executable) executeDefault(inputs []compute.Buffer, donate []bool, defaultDevice compute.DeviceNum) ([]compute.Buffer, error) {
	outWrappers := e.cachedOutWraps
	ortOutputs := e.cachedOrtOutputs

	// The MIGraphX EP unreliably copies results into caller-preallocated (CPU) output
	// buffers, so for it we let ONNX Runtime allocate the outputs instead and wrap
	// them afterwards -- the same thing already done for dynamic shapes.
	preallocOutputs := e.executionProvider != executionprovider.MIGraphX

	for i, sh := range e.outputShapes {
		if sh.IsDynamic() || !preallocOutputs {
			outWrappers[i] = nil
			ortOutputs[i] = nil
			continue
		}
		matchedIdx := -1
		for idx, rw := range e.reusableWrappers {
			if wrapperMatchesShape(rw, sh) {
				matchedIdx = idx
				break
			}
		}

		if matchedIdx >= 0 {
			outWrappers[i] = e.reusableWrappers[matchedIdx]
			last := len(e.reusableWrappers) - 1
			e.reusableWrappers[matchedIdx] = e.reusableWrappers[last]
			e.reusableWrappers[last] = nil
			e.reusableWrappers = e.reusableWrappers[:last]
		} else {
			outWrappers[i] = nil
		}
	}

	for i, sh := range e.outputShapes {
		if sh.IsDynamic() || !preallocOutputs {
			ortOutputs[i] = nil
			continue
		}
		if outWrappers[i] == nil {
			var err error
			outWrappers[i], err = NewEmptyOrtTensorWrapper(sh)
			if err != nil {
				for _, w := range outWrappers {
					if w != nil {
						_ = w.Destroy()
					}
				}
				return nil, err
			}
		}
		ortOutputs[i] = outWrappers[i].Value()
	}

	var start time.Time
	if klog.V(1).Enabled() {
		start = time.Now()
	}
	if klog.V(2).Enabled() {
		klog.Infof("Starting session.Run on CPU (inputs=%d, outputs=%d)...", len(e.cachedOrtInputs), len(ortOutputs))
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if e.executionProvider == executionprovider.MIGraphX {
		if err := e.migraphxWorkaround(); err != nil {
			return nil, err
		}
	}

	err := e.session.Run(e.cachedOrtInputs, ortOutputs)
	if klog.V(1).Enabled() {
		klog.Infof("Execution (CPU) elapsed time: %s\n", humanize.Duration(time.Since(start)))
	}
	if err != nil {
		for _, w := range outWrappers {
			if w != nil {
				_ = w.Destroy()
			}
		}
		return nil, errors.Wrap(err, "onnxruntime execution failed")
	}
	if e.executionProvider == executionprovider.MIGraphX {
		// The MIGraphX EP leaves GPU work in flight; synchronize so that output
		// buffers (and inputs of the next execution) are fully materialized.
		if err := ort.HipDeviceSynchronize(); err != nil {
			for _, w := range outWrappers {
				if w != nil {
					_ = w.Destroy()
				}
			}
			return nil, errors.Wrap(err, "onnxruntime execution failed")
		}
	}

	outBuffers := make([]compute.Buffer, len(e.outputShapes))
	for i, sh := range e.outputShapes {
		actualShape := sh
		if sh.IsDynamic() || !preallocOutputs {
			var err error
			outWrappers[i], err = WrapOrtValue(ortOutputs[i], sh)
			if err != nil {
				return nil, err
			}
			actualOrtShape := outWrappers[i].GetShape()
			concreteDims := make([]int, len(actualOrtShape))
			for k, dim := range actualOrtShape {
				concreteDims[k] = int(dim)
			}
			actualShape.Dimensions = concreteDims
		}

		execBackpointer := e
		if e.executionProvider == executionprovider.CUDA {
			// CUDA buffers are never recycled via the executable.
			execBackpointer = nil
		}
		outBuffers[i] = &Buffer{
			backend:           e.backend,
			wrapper:           outWrappers[i],
			shape:             actualShape,
			device:            defaultDevice,
			executionProvider: e.executionProvider,
			executable:        execBackpointer,
		}
		outWrappers[i] = nil
	}

	for i, inp := range inputs {
		if i < len(donate) && donate[i] {
			buf := inp.(*Buffer)
			w := buf.wrapper
			buf.wrapper = nil

			if w != nil {
				_ = w.Destroy()
			}
		}
	}

	return outBuffers, nil
}

// migraphxWorkaround implements the upstream MIGraphX bug workaround: the first
// evaluations for each distinct input-shape signature return uninitialized/garbage
// outputs (the number of affected runs varies with the model). It detects new shapes
// and repeats the run until the outputs stabilize, discarding all results.
func (e *Executable) migraphxWorkaround() error {
	key := inputShapeSignature(e.cachedOrtInputs)
	if e.warmedShapes[key] {
		return nil
	}
	var prevBytes [][]byte
	for i := 0; i < migraphxMaxWarmUpRuns; i++ {
		dummyOutputs := make([]ort.Value, len(e.outputShapes))
		if err := e.session.Run(e.cachedOrtInputs, dummyOutputs); err != nil {
			return errors.Wrap(err, "migraphx warm-up run failed")
		}
		curBytes := make([][]byte, len(dummyOutputs))
		for j, o := range dummyOutputs {
			if o == nil {
				continue
			}
			if ptr, err := o.GetTensorMutableData(); err == nil {
				shape := e.outputShapes[j]
				if rtShape, serr := ort.ShapeOf(o); serr == nil && len(rtShape) > 0 {
					numel := int64(1)
					for _, d := range rtShape {
						if d > 0 {
							numel *= d
						}
					}
					shape = shapes.Make(e.outputShapes[j].DType, int(numel))
				}
				size := shape.Size() * shape.DType.Size()
				curBytes[j] = unsafe.Slice((*byte)(ptr), max(size, 1))
			}
			_ = o.Destroy()
		}
		if err := ort.HipDeviceSynchronize(); err != nil {
			return errors.Wrap(err, "migraphx warm-up synchronization failed")
		}
		stable := len(prevBytes) > 0 && len(prevBytes) == len(curBytes)
		for j := range curBytes {
			if !stable {
				break
			}
			if !bytes.Equal(curBytes[j], prevBytes[j]) {
				stable = false
			}
		}
		prevBytes = curBytes
		if stable {
			break
		}
	}
	e.warmedShapes[key] = true
	return nil
}

// inputShapeSignature returns a stable string key for the shapes of the given input values.
func inputShapeSignature(inputs []ort.Value) string {
	var sb strings.Builder
	for _, inp := range inputs {
		if t, ok := inp.(interface{ GetShape() ort.Shape }); ok {
			fmt.Fprintf(&sb, "%v;", t.GetShape())
		} else {
			sb.WriteString("?;")
		}
	}
	return sb.String()
}
