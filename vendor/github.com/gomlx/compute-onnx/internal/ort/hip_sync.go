// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package ort

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stddef.h>

typedef int (*hipDeviceSynchronizeFn)(void);

static hipDeviceSynchronizeFn resolved_hipDeviceSync = NULL;
static int hip_sync_resolved = 0;

// Resolve hipDeviceSynchronize — try global process symbols first (the HIP runtime is
// usually already loaded by libamdhip64 pulled in by the ONNX Runtime provider library),
// then explicitly load libamdhip64.
// Returns 0 on success, -1 if the symbol cannot be found.
static int resolve_hip_sync_symbol() {
    if (hip_sync_resolved) return (resolved_hipDeviceSync != NULL) ? 0 : -1;

    void* handle = dlopen(NULL, RTLD_NOW | RTLD_GLOBAL);
    if (handle != NULL) {
        resolved_hipDeviceSync = (hipDeviceSynchronizeFn)dlsym(handle, "hipDeviceSynchronize");
    }
    if (resolved_hipDeviceSync == NULL) {
        const char* names[] = {"libamdhip64.so", "libamdhip64.so.7", "libamdhip64.so.6", "libamdhip64.so.5"};
        for (int i = 0; i < 4 && resolved_hipDeviceSync == NULL; i++) {
            void* lib_handle = dlopen(names[i], RTLD_NOW | RTLD_GLOBAL);
            if (lib_handle != NULL) {
                resolved_hipDeviceSync = (hipDeviceSynchronizeFn)dlsym(lib_handle, "hipDeviceSynchronize");
            }
        }
    }
    hip_sync_resolved = 1;
    return (resolved_hipDeviceSync != NULL) ? 0 : -1;
}

static int hip_device_synchronize() {
    if (resolve_hip_sync_symbol() != 0) return -1;
    return resolved_hipDeviceSync();
}
*/
import "C"
import (
	"runtime"

	"github.com/pkg/errors"
)

// HipDeviceSynchronize blocks until all pending HIP device work has completed.
// It resolves hipDeviceSynchronize via dlsym from the already-loaded HIP runtime.
// It returns an error if the HIP runtime is not loaded (e.g. CPU-only builds).
func HipDeviceSynchronize() error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	ret := C.hip_device_synchronize()
	if ret < 0 {
		return errors.Errorf("hipDeviceSynchronize could not be resolved from the HIP runtime (libamdhip64)")
	}
	if ret != 0 {
		return errors.Errorf("hipDeviceSynchronize failed with error code %d", int(ret))
	}
	return nil
}

// LoadHIPLibrary resolves and dynamically loads the HIP runtime (libamdhip64),
// returning whether hipDeviceSynchronize could be resolved in this process.
func LoadHIPLibrary() bool {
	return C.resolve_hip_sync_symbol() == 0
}
