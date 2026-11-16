// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package ort

/*
#include "onnxruntime_c_api.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"math/bits"
	"unsafe"

	"github.com/gomlx/compute-onnx/internal/pool"
)

// Arena implements a trivial arena object to speed up allocations used in CGO calls.
//
// It pre-allocates the given size in bytes in C — so it does not need to be pinned
// when using CGO and allows for fast suballocations.
// It can only be freed or reset all at once.
type Arena struct {
	next *Arena

	buf       []byte
	size      int
	current   int
	poolIndex int // index in arenaPools, -1 if not from pool
}

func (a *Arena) Next() *Arena        { return a.next }
func (a *Arena) SetNext(next *Arena) { a.next = next }

const arenaAlignBytes = 8

func NewArena(size int) *Arena {
	ptr := C.calloc(C.size_t(size), 1)
	return &Arena{
		buf:       unsafe.Slice((*byte)(ptr), size),
		size:      size,
		poolIndex: -1,
	}
}

// Free releases the C-allocated memory. After Free, the arena cannot be reused.
func (a *Arena) Free() {
	if len(a.buf) > 0 {
		C.free(unsafe.Pointer(&a.buf[0]))
		a.buf = nil
		a.size = 0
		a.current = 0
		a.poolIndex = -1
	}
}

// Reset zeroes the used portion and resets the allocation pointer.
// The underlying C memory is retained for reuse.
func (a *Arena) Reset() {
	if len(a.buf) == 0 || a.size == 0 {
		a.current = 0
		return
	}
	if a.current > 0 {
		clearSize := min(a.size, a.current)
		C.memset(unsafe.Pointer(&a.buf[0]), 0, C.size_t(clearSize))
	}
	a.current = 0
}

func (a *Arena) alloc(allocSize int) unsafe.Pointer {
	if a.current+allocSize > a.size {
		panic(fmt.Sprintf("Arena out of memory: current=%d, requested=%d, size=%d", a.current, allocSize, a.size))
	}
	ptr := unsafe.Pointer(&a.buf[a.current])
	a.current += allocSize
	a.current = (a.current + arenaAlignBytes - 1) &^ (arenaAlignBytes - 1)
	return ptr
}

func (a *Arena) AllocInt64Slice(n int) *C.int64_t {
	return (*C.int64_t)(a.alloc(n * 8))
}

func (a *Arena) AllocCharStarSlice(n int) **C.char {
	sizeOfPtr := int(unsafe.Sizeof(uintptr(0)))
	return (**C.char)(a.alloc(n * sizeOfPtr))
}

func (a *Arena) AllocOrtValueStarSlice(n int) **C.OrtValue {
	sizeOfPtr := int(unsafe.Sizeof(uintptr(0)))
	return (**C.OrtValue)(a.alloc(n * sizeOfPtr))
}

func (a *Arena) AllocCString(s string) *C.char {
	n := len(s)
	allocSize := n + 1
	ptr := a.alloc(allocSize)
	dst := unsafe.Slice((*byte)(ptr), allocSize)
	copy(dst, s)
	dst[n] = 0 // null terminator
	return (*C.char)(ptr)
}

// --- Arena Pools ---

const (
	minPooledArenaSize = 2048
	maxPooledArenaSize = 16 * 1024 * 1024
)

// ArenaPools manages pools of Arena objects with power-of-2 sizes.
type ArenaPools struct {
	pools    []*pool.Pool[Arena, *Arena]
	minShift int
	maxShift int
}

var globalArenaPools *ArenaPools

func init() {
	globalArenaPools = newArenaPools()
}

func newArenaPools() *ArenaPools {
	minShift := bits.TrailingZeros(uint(minPooledArenaSize))
	maxShift := bits.TrailingZeros(uint(maxPooledArenaSize))
	numPools := maxShift - minShift + 1

	ap := &ArenaPools{
		pools:    make([]*pool.Pool[Arena, *Arena], numPools),
		minShift: minShift,
		maxShift: maxShift,
	}

	for poolIdx := range numPools {
		poolSize := 1 << (poolIdx + minShift)
		ap.pools[poolIdx] = pool.New[Arena, *Arena](func() *Arena {
			arena := NewArena(poolSize)
			arena.poolIndex = poolIdx
			return arena
		})
	}
	return ap
}

// GetArena returns an Arena of at least targetSize bytes from the global pool.
func GetArena(targetSize int) *Arena {
	return globalArenaPools.Get(targetSize)
}

// ReturnArena returns an Arena to the global pool for reuse.
func ReturnArena(arena *Arena) {
	globalArenaPools.Return(arena)
}

func (ap *ArenaPools) Get(targetSize int) *Arena {
	if targetSize <= minPooledArenaSize {
		targetSize = minPooledArenaSize
	}

	shift := bits.Len(uint(targetSize - 1))
	if shift < ap.minShift {
		shift = ap.minShift
	}

	if shift > ap.maxShift {
		return NewArena(targetSize)
	}

	poolIndex := shift - ap.minShift
	return ap.pools[poolIndex].Get()
}

func (ap *ArenaPools) Return(arena *Arena) {
	if arena == nil || len(arena.buf) == 0 {
		return
	}

	if arena.poolIndex < 0 || arena.poolIndex >= len(ap.pools) {
		arena.Free()
		return
	}

	arena.Reset()
	ap.pools[arena.poolIndex].Put(arena)
}
