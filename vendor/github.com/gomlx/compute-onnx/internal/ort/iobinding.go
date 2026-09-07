// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package ort

/*
#include "onnxruntime_c_api.h"
#include <stdlib.h>

// Forward declarations
OrtStatus* wrapper_CreateIoBinding(const OrtApi* api, OrtSession* session, OrtIoBinding** out);
void wrapper_ReleaseIoBinding(const OrtApi* api, OrtIoBinding* binding);
OrtStatus* wrapper_BindInput(const OrtApi* api, OrtIoBinding* binding, const char* name, const OrtValue* val);
OrtStatus* wrapper_BindOutput(const OrtApi* api, OrtIoBinding* binding, const char* name, const OrtValue* val);
OrtStatus* wrapper_BindOutputToDevice(const OrtApi* api, OrtIoBinding* binding, const char* name, const OrtMemoryInfo* mem_info);
OrtStatus* wrapper_RunWithBinding(const OrtApi* api, OrtSession* session, const OrtRunOptions* run_options, const OrtIoBinding* binding);
OrtStatus* wrapper_GetBoundOutputValues(const OrtApi* api, const OrtIoBinding* binding, OrtAllocator* allocator, OrtValue*** output, size_t* output_count);
void wrapper_ClearBoundInputs(const OrtApi* api, OrtIoBinding* binding);
void wrapper_ClearBoundOutputs(const OrtApi* api, OrtIoBinding* binding);
OrtStatus* wrapper_SynchronizeBoundInputs(const OrtApi* api, OrtIoBinding* binding);
OrtStatus* wrapper_SynchronizeBoundOutputs(const OrtApi* api, OrtIoBinding* binding);

OrtStatus* wrapper_CreateMemoryInfo(const OrtApi* api, const char* name, enum OrtAllocatorType type, int id, enum OrtMemType mem_type, OrtMemoryInfo** out);
void wrapper_ReleaseMemoryInfo(const OrtApi* api, OrtMemoryInfo* info);

OrtStatus* wrapper_GetAllocatorWithDefaultOptions(const OrtApi* api, OrtAllocator** out);
OrtStatus* wrapper_CreateAllocator(const OrtApi* api, const OrtSession* session, const OrtMemoryInfo* mem_info, OrtAllocator** out);
void wrapper_ReleaseAllocator(const OrtApi* api, OrtAllocator* allocator);
OrtStatus* wrapper_AllocatorAlloc(const OrtApi* api, OrtAllocator* allocator, size_t size, void** out);
OrtStatus* wrapper_AllocatorFree(const OrtApi* api, OrtAllocator* allocator, void* p);

OrtStatus* wrapper_GetTensorMutableData(const OrtApi* api, OrtValue* value, void** out);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// IoBinding wraps an OrtIoBinding for binding inputs/outputs to GPU memory.
type IoBinding struct {
	binding *C.OrtIoBinding
	session *Session
}

// NewIoBinding creates a new IoBinding for the given session.
func NewIoBinding(session *Session) (*IoBinding, error) {
	var binding *C.OrtIoBinding
	status := C.wrapper_CreateIoBinding(ortApi, session.session, &binding)
	if err := statusToError(status); err != nil {
		return nil, err
	}
	return &IoBinding{binding: binding, session: session}, nil
}

func (b *IoBinding) Destroy() {
	if b.binding != nil {
		C.wrapper_ReleaseIoBinding(ortApi, b.binding)
		b.binding = nil
	}
}

// BindInput binds a Value (CPU or GPU OrtValue) as a named input.
func (b *IoBinding) BindInput(name *C.char, val Value) error {
	if val == nil || val.cValue() == nil {
		return fmt.Errorf("cannot bind nil Value for input %s", C.GoString(name))
	}
	status := C.wrapper_BindInput(ortApi, b.binding, name, val.cValue())
	return statusToError(status)
}

// BindOutput binds a pre-allocated Value as a named output.
func (b *IoBinding) BindOutput(name *C.char, val Value) error {
	status := C.wrapper_BindOutput(ortApi, b.binding, name, val.cValue())
	return statusToError(status)
}

// BindOutputToDevice tells ORT to allocate the named output on the given device.
func (b *IoBinding) BindOutputToDevice(name *C.char, memInfo *MemoryInfo) error {
	status := C.wrapper_BindOutputToDevice(ortApi, b.binding, name, memInfo.info)
	return statusToError(status)
}

// RunWithBinding executes the session with the current bindings.
func (b *IoBinding) RunWithBinding() error {
	status := C.wrapper_RunWithBinding(ortApi, b.session.session, nil, b.binding)
	return statusToError(status)
}

// SynchronizeBoundOutputs waits for all async output copies to complete.
func (b *IoBinding) SynchronizeBoundOutputs() error {
	status := C.wrapper_SynchronizeBoundOutputs(ortApi, b.binding)
	return statusToError(status)
}

// GetBoundOutputValues retrieves the output OrtValues after RunWithBinding.
// Returns raw *C.OrtValue pointers. Caller is responsible for wrapping/releasing them.
func (b *IoBinding) GetBoundOutputValues() ([]*C.OrtValue, error) {
	var allocator *C.OrtAllocator
	status := C.wrapper_GetAllocatorWithDefaultOptions(ortApi, &allocator)
	if err := statusToError(status); err != nil {
		return nil, err
	}

	var outputPtr **C.OrtValue
	var count C.size_t
	status = C.wrapper_GetBoundOutputValues(ortApi, b.binding, allocator, &outputPtr, &count)
	if err := statusToError(status); err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	// Convert the C array to a Go slice (don't free the individual OrtValues — ownership is transferred).
	values := make([]*C.OrtValue, int(count))
	cSlice := unsafe.Slice(outputPtr, int(count))
	copy(values, cSlice)

	// Free the array pointer itself (allocated by ORT via the allocator).
	C.wrapper_AllocatorFree(ortApi, allocator, unsafe.Pointer(outputPtr))

	return values, nil
}

// ClearBoundInputs clears all input bindings.
func (b *IoBinding) ClearBoundInputs() {
	C.wrapper_ClearBoundInputs(ortApi, b.binding)
}

// ClearBoundOutputs clears all output bindings.
func (b *IoBinding) ClearBoundOutputs() {
	C.wrapper_ClearBoundOutputs(ortApi, b.binding)
}
// NewCUDAMemoryInfo creates a MemoryInfo for CUDA device memory (device 0).
func NewCUDAMemoryInfo() (*MemoryInfo, error) {
	var info *C.OrtMemoryInfo
	cName := C.CString("Cuda")
	defer C.free(unsafe.Pointer(cName))
	status := C.wrapper_CreateMemoryInfo(
		ortApi,
		cName,
		C.OrtDeviceAllocator,
		C.int(0), // device ID
		C.OrtMemTypeDefault,
		&info,
	)
	if err := statusToError(status); err != nil {
		return nil, err
	}
	return &MemoryInfo{info: info}, nil
}

// CopyGPUToHost copies data from a GPU OrtValue to a host buffer.
// Uses cudaMemcpy resolved via dlsym from the already-loaded CUDA runtime.
func CopyGPUToHost(gpuValue Value, dst unsafe.Pointer, byteSize int) error {
	srcPtr, err := gpuValue.GetTensorMutableData()
	if err != nil {
		return fmt.Errorf("failed to get GPU tensor data pointer: %w", err)
	}
	return CudaMemcpyD2H(dst, srcPtr, byteSize)
}

// CopyHostToGPU copies data from a host buffer to a GPU OrtValue.
// Uses cudaMemcpy resolved via dlsym from the already-loaded CUDA runtime.
func CopyHostToGPU(gpuValue Value, src unsafe.Pointer, byteSize int) error {
	dstPtr, err := gpuValue.GetTensorMutableData()
	if err != nil {
		return fmt.Errorf("failed to get GPU tensor data pointer: %w", err)
	}
	return CudaMemcpyH2D(dstPtr, src, byteSize)
}

// ShapeOf returns the runtime shape of the given value.
func ShapeOf(v Value) (Shape, error) {
	return GetOrtValueShape(v.cValue())
}
