// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package ort

/*
#include <dlfcn.h>
#include <stddef.h>
#include <string.h>

// cudaMemcpyKind enum values
#define CUDA_MEMCPY_HOST_TO_DEVICE 1
#define CUDA_MEMCPY_DEVICE_TO_HOST 2

typedef int (*cudaMemcpyFn)(void* dst, const void* src, size_t count, int kind);
typedef int (*cudaDeviceSynchronizeFn)(void);

static cudaMemcpyFn resolved_cudaMemcpy = NULL;
static cudaDeviceSynchronizeFn resolved_cudaDeviceSync = NULL;

// Resolve cudaMemcpy — try global process symbols first via dlopen(NULL), then explicitly load libcudart.
// Returns 0 on success, -1 if the symbol cannot be found.
static int resolve_cuda_symbols() {
    if (resolved_cudaMemcpy != NULL) return 0;

    // Try 1: already loaded in process symbol table (POSIX compliant using dlopen(NULL)).
    void* handle = dlopen(NULL, RTLD_NOW | RTLD_GLOBAL);
    if (handle != NULL) {
        resolved_cudaMemcpy = (cudaMemcpyFn)dlsym(handle, "cudaMemcpy");
        resolved_cudaDeviceSync = (cudaDeviceSynchronizeFn)dlsym(handle, "cudaDeviceSynchronize");
    }

    // Try 2: explicitly load libcudart shared library if not found in global process symbols.
    if (resolved_cudaMemcpy == NULL || resolved_cudaDeviceSync == NULL) {
        void* lib_handle = dlopen("libcudart.so", RTLD_NOW | RTLD_GLOBAL);
        if (lib_handle == NULL) {
            lib_handle = dlopen("libcudart.so.13", RTLD_NOW | RTLD_GLOBAL);
        }
        if (lib_handle == NULL) {
            lib_handle = dlopen("libcudart.so.12", RTLD_NOW | RTLD_GLOBAL);
        }
        if (lib_handle == NULL) {
            lib_handle = dlopen("libcudart.so.11", RTLD_NOW | RTLD_GLOBAL);
        }
        if (lib_handle != NULL) {
            if (resolved_cudaMemcpy == NULL) {
                resolved_cudaMemcpy = (cudaMemcpyFn)dlsym(lib_handle, "cudaMemcpy");
            }
            if (resolved_cudaDeviceSync == NULL) {
                resolved_cudaDeviceSync = (cudaDeviceSynchronizeFn)dlsym(lib_handle, "cudaDeviceSynchronize");
            }
        }
    }

    if (resolved_cudaMemcpy == NULL) return -1;
    return 0;
}

// Copy from GPU device memory to host memory.
// Returns CUDA error code (0 = success).
static int cuda_memcpy_d2h(void* dst, const void* src, size_t count) {
    if (resolve_cuda_symbols() != 0) return -1;
    return resolved_cudaMemcpy(dst, src, count, CUDA_MEMCPY_DEVICE_TO_HOST);
}

// Copy from host memory to GPU device memory.
// Returns CUDA error code (0 = success).
static int cuda_memcpy_h2d(void* dst, const void* src, size_t count) {
    if (resolve_cuda_symbols() != 0) return -1;
    return resolved_cudaMemcpy(dst, src, count, CUDA_MEMCPY_HOST_TO_DEVICE);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// CudaMemcpyD2H copies byteSize bytes from a CUDA device pointer to a host pointer.
// It uses dlsym to resolve cudaMemcpy from the already-loaded CUDA runtime.
func CudaMemcpyD2H(hostDst unsafe.Pointer, deviceSrc unsafe.Pointer, byteSize int) error {
	ret := C.cuda_memcpy_d2h(hostDst, deviceSrc, C.size_t(byteSize))
	if ret != 0 {
		return fmt.Errorf("cudaMemcpy D2H failed with error code %d", int(ret))
	}
	return nil
}

// CudaMemcpyH2D copies byteSize bytes from a host pointer to a CUDA device pointer.
// It uses dlsym to resolve cudaMemcpy from the already-loaded CUDA runtime.
func CudaMemcpyH2D(deviceDst unsafe.Pointer, hostSrc unsafe.Pointer, byteSize int) error {
	ret := C.cuda_memcpy_h2d(deviceDst, hostSrc, C.size_t(byteSize))
	if ret != 0 {
		return fmt.Errorf("cudaMemcpy H2D failed with error code %d", int(ret))
	}
	return nil
}
