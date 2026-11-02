// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build linux || darwin

package ort

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func loadLibrary(libPath string) (unsafe.Pointer, error) {
	libPathC := C.CString(libPath)
	defer C.free(unsafe.Pointer(libPathC))

	// Clear dlerror
	C.dlerror()
	handle := C.dlopen(libPathC, C.RTLD_LAZY|C.RTLD_LOCAL)
	if handle == nil {
		errStr := C.GoString(C.dlerror())
		return nil, fmt.Errorf("failed to dlopen %q: %s", libPath, errStr)
	}
	return handle, nil
}

func getOrtGetApiBase(handle unsafe.Pointer) (unsafe.Pointer, error) {
	symC := C.CString("OrtGetApiBase")
	defer C.free(unsafe.Pointer(symC))

	// Clear dlerror
	C.dlerror()
	p := C.dlsym(handle, symC)
	errC := C.dlerror()
	if errC != nil {
		return nil, fmt.Errorf("failed to find symbol OrtGetApiBase: %s", C.GoString(errC))
	}
	if p == nil {
		return nil, fmt.Errorf("OrtGetApiBase symbol resolved to nil")
	}
	return p, nil
}

func closeLibrary(handle unsafe.Pointer) error {
	// Clear dlerror
	C.dlerror()
	res := C.dlclose(handle)
	if res != 0 {
		errC := C.dlerror()
		return fmt.Errorf("failed to dlclose: %s", C.GoString(errC))
	}
	return nil
}
