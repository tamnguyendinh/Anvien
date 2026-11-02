package ortgenai

/*
#cgo CFLAGS: -O2 -g
#include "ort_genai_wrapper.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"sync"
	"unsafe"
)

var logCallbackMu sync.RWMutex
var logCallback func(string)

// SetLogBool sets a process-wide ORT GenAI logging option.
func SetLogBool(name string, value bool) error {
	if !IsInitialized() {
		return ErrNotInitialized
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	if err := OgaResultToError(C.SetLogBool(cName, C.bool(value))); err != nil {
		return fmt.Errorf("setting logging option: %w", err)
	}
	return nil
}

// SetLogString sets a process-wide ORT GenAI logging option.
func SetLogString(name, value string) error {
	if !IsInitialized() {
		return ErrNotInitialized
	}
	cName, cValue := C.CString(name), C.CString(value)
	defer C.free(unsafe.Pointer(cName))
	defer C.free(unsafe.Pointer(cValue))
	if err := OgaResultToError(C.SetLogString(cName, cValue)); err != nil {
		return fmt.Errorf("setting logging option: %w", err)
	}
	return nil
}

// SetLogCallback replaces the process-wide callback. Each message is copied into a Go string
// before the callback is invoked. Passing nil restores the native default logging destination.
func SetLogCallback(callback func(string)) error {
	if !IsInitialized() {
		return ErrNotInitialized
	}
	logCallbackMu.Lock()
	previous := logCallback
	logCallback = callback
	logCallbackMu.Unlock()
	if err := OgaResultToError(C.SetLogCallback(C.bool(callback != nil))); err != nil {
		logCallbackMu.Lock()
		logCallback = previous
		logCallbackMu.Unlock()
		return fmt.Errorf("setting log callback: %w", err)
	}
	return nil
}

//export goLogCallback
func goLogCallback(message *C.char, length C.size_t) {
	defer func() { _ = recover() }()
	if message == nil {
		return
	}
	text := C.GoStringN(message, C.int(length))
	logCallbackMu.RLock()
	callback := logCallback
	logCallbackMu.RUnlock()
	if callback != nil {
		callback(text)
	}
}

func clearLogCallback() {
	logCallbackMu.Lock()
	logCallback = nil
	logCallbackMu.Unlock()
}

// SetGPUDeviceID sets the process-wide active GPU device.
func SetGPUDeviceID(deviceID int) error {
	if !IsInitialized() {
		return ErrNotInitialized
	}
	if err := OgaResultToError(C.SetCurrentGpuDeviceId(C.int(deviceID))); err != nil {
		return fmt.Errorf("setting GPU device ID: %w", err)
	}
	return nil
}

// GPUDeviceID returns the process-wide active GPU device ID.
func GPUDeviceID() (int, error) {
	if !IsInitialized() {
		return 0, ErrNotInitialized
	}
	var deviceID C.int
	if err := OgaResultToError(C.GetCurrentGpuDeviceId(&deviceID)); err != nil {
		return 0, fmt.Errorf("getting GPU device ID: %w", err)
	}
	return int(deviceID), nil
}

// RuntimeSettings owns settings used when creating a model.
type RuntimeSettings struct{ ptr *C.OgaRuntimeSettings }

func CreateRuntimeSettings() (*RuntimeSettings, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	var ptr *C.OgaRuntimeSettings
	if err := OgaResultToError(C.CreateRuntimeSettings(&ptr)); err != nil {
		return nil, fmt.Errorf("creating runtime settings: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("runtime settings creation returned nil without error")
	}
	return &RuntimeSettings{ptr: ptr}, nil
}

// SetHandle configures a runtime handle. The pointer must be a native handle valid for the model's lifetime.
func (s *RuntimeSettings) SetHandle(name string, handle unsafe.Pointer) error {
	if s == nil || s.ptr == nil {
		return errors.New("runtime settings are destroyed")
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	if err := OgaResultToError(C.RuntimeSettingsSetHandle(s.ptr, cName, handle)); err != nil {
		return fmt.Errorf("setting runtime handle: %w", err)
	}
	return nil
}

// Destroy releases native runtime settings.
func (s *RuntimeSettings) Destroy() {
	if s != nil && s.ptr != nil {
		C.DestroyRuntimeSettings(s.ptr)
		s.ptr = nil
	}
}
