// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package web

import (
	"unsafe"
)

// unsafeSliceBytes reinterprets any typed slice as a []byte with equivalent byte length without allocation.
func unsafeSliceBytes[T any](s []T) []byte {
	if len(s) == 0 {
		return nil
	}
	var dummy T
	elemSize := int(unsafe.Sizeof(dummy))
	byteLen := len(s) * elemSize
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), byteLen)
}
