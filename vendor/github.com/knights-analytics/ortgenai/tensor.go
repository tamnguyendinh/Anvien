package ortgenai

/*
#cgo CFLAGS: -O2 -g
#include "ort_genai_wrapper.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

// ElementType identifies a native tensor element representation.
type ElementType uint32

const (
	ElementTypeUndefined ElementType = iota
	ElementTypeFloat32
	ElementTypeUint8
	ElementTypeInt8
	ElementTypeUint16
	ElementTypeInt16
	ElementTypeInt32
	ElementTypeInt64
	ElementTypeString
	ElementTypeBool
	ElementTypeFloat16
	ElementTypeFloat64
	ElementTypeUint32
	ElementTypeUint64
	ElementTypeComplex64
	ElementTypeComplex128
	ElementTypeBFloat16
)

// Tensor owns its native handle. User-provided buffer memory is copied into C memory and retained
// until Destroy, as the native tensor does not take ownership of that buffer.
type Tensor struct {
	ptr         *C.OgaTensor
	ownedBuffer unsafe.Pointer
}

// NewTensorFromBuffer copies data into a tensor with the specified shape and element type.
func NewTensorFromBuffer(data []byte, shape []int64, elementType ElementType) (*Tensor, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if len(shape) == 0 {
		return nil, errors.New("tensor shape must have at least one dimension")
	}
	size, err := tensorByteSize(shape, elementType)
	if err != nil {
		return nil, err
	}
	if len(data) != size {
		return nil, fmt.Errorf("tensor buffer has %d bytes, expected %d", len(data), size)
	}
	var buffer unsafe.Pointer
	if size != 0 {
		buffer = C.CBytes(data)
	}
	if size != 0 && buffer == nil {
		return nil, errors.New("allocating tensor buffer failed")
	}
	var ptr *C.OgaTensor
	if err := OgaResultToError(C.CreateTensorFromBuffer(buffer, (*C.int64_t)(unsafe.Pointer(&shape[0])), C.size_t(len(shape)), C.OgaElementType(elementType), &ptr)); err != nil {
		C.free(buffer)
		return nil, fmt.Errorf("creating tensor: %w", err)
	}
	if ptr == nil {
		C.free(buffer)
		return nil, errors.New("tensor creation returned nil without error")
	}
	return &Tensor{ptr: ptr, ownedBuffer: buffer}, nil
}

func tensorByteSize(shape []int64, elementType ElementType) (int, error) {
	var elementSize int64
	switch elementType {
	case ElementTypeFloat32, ElementTypeInt32, ElementTypeUint32:
		elementSize = 4
	case ElementTypeUint8, ElementTypeInt8, ElementTypeBool:
		elementSize = 1
	case ElementTypeUint16, ElementTypeInt16, ElementTypeFloat16, ElementTypeBFloat16:
		elementSize = 2
	case ElementTypeInt64, ElementTypeUint64, ElementTypeFloat64, ElementTypeComplex64:
		elementSize = 8
	case ElementTypeComplex128:
		elementSize = 16
	default:
		return 0, fmt.Errorf("unsupported tensor element type %d", elementType)
	}
	count := int64(1)
	for _, dimension := range shape {
		if dimension < 0 || (dimension != 0 && count > math.MaxInt64/dimension) {
			return 0, errors.New("invalid or overflowing tensor shape")
		}
		count *= dimension
	}
	if count > int64(math.MaxInt)/elementSize {
		return 0, errors.New("tensor buffer is too large")
	}
	return int(count * elementSize), nil
}

func (t *Tensor) check() error {
	if t == nil || t.ptr == nil {
		return errors.New("tensor is destroyed")
	}
	return nil
}

// ElementType returns the tensor's native element type.
func (t *Tensor) ElementType() (ElementType, error) {
	if err := t.check(); err != nil {
		return ElementTypeUndefined, err
	}
	var elementType C.OgaElementType
	if err := OgaResultToError(C.TensorGetType(t.ptr, &elementType)); err != nil {
		return ElementTypeUndefined, fmt.Errorf("reading tensor element type: %w", err)
	}
	return ElementType(elementType), nil
}

// Shape returns a copy of the tensor dimensions.
func (t *Tensor) Shape() ([]int64, error) {
	if err := t.check(); err != nil {
		return nil, err
	}
	var rank C.size_t
	if err := OgaResultToError(C.TensorGetShapeRank(t.ptr, &rank)); err != nil {
		return nil, fmt.Errorf("reading tensor rank: %w", err)
	}
	shape := make([]int64, int(rank))
	if len(shape) != 0 {
		if err := OgaResultToError(C.TensorGetShape(t.ptr, (*C.int64_t)(unsafe.Pointer(&shape[0])), rank)); err != nil {
			return nil, fmt.Errorf("reading tensor shape: %w", err)
		}
	}
	return shape, nil
}

// CopyData returns a copy of the tensor's raw bytes.
func (t *Tensor) CopyData() ([]byte, error) {
	if err := t.check(); err != nil {
		return nil, err
	}
	shape, err := t.Shape()
	if err != nil {
		return nil, err
	}
	elementType, err := t.ElementType()
	if err != nil {
		return nil, err
	}
	size, err := tensorByteSize(shape, elementType)
	if err != nil {
		return nil, err
	}
	var data unsafe.Pointer
	if err := OgaResultToError(C.TensorGetData(t.ptr, &data)); err != nil {
		return nil, fmt.Errorf("reading tensor data: %w", err)
	}
	if size == 0 {
		return []byte{}, nil
	}
	if data == nil {
		return nil, errors.New("tensor has nil data")
	}
	return C.GoBytes(data, C.int(size)), nil
}

// Destroy releases the native tensor and any copied user buffer it retains.
func (t *Tensor) Destroy() {
	if t == nil {
		return
	}
	if t.ptr != nil {
		C.DestroyTensor(t.ptr)
		t.ptr = nil
	}
	if t.ownedBuffer != nil {
		C.free(t.ownedBuffer)
		t.ownedBuffer = nil
	}
}
