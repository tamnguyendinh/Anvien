// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build !js

package native

import (
	"unsafe"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	ort "github.com/gomlx/compute-onnx/internal/ort"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/dtypes/bfloat16"
	"github.com/gomlx/compute/dtypes/float16"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

// OrtTensorWrapper abstracts typed and GPU OrtValue containers.
type OrtTensorWrapper interface {
	GetShape() ort.Shape
	GetDType() dtypes.DType
	Destroy() error
	Value() ort.Value
	ToFlatData(flat any) error
	CopyFrom(flat any) error
	GetData() any
}

type typedTensor[T ort.TensorData] struct {
	tensor  *ort.Tensor[T]
	rawData any
}

func (t *typedTensor[T]) GetShape() ort.Shape {
	return t.tensor.GetShape()
}

func (t *typedTensor[T]) GetDType() dtypes.DType {
	var dummy T
	switch any(dummy).(type) {
	case float32:
		return dtypes.Float32
	case float64:
		return dtypes.Float64
	case int32:
		return dtypes.Int32
	case int64:
		return dtypes.Int64
	case bool:
		return dtypes.Bool
	case int8:
		return dtypes.Int8
	case uint8:
		return dtypes.Uint8
	case int16:
		return dtypes.Int16
	case uint16:
		return dtypes.Uint16
	case uint32:
		return dtypes.Uint32
	case float16.Float16:
		return dtypes.Float16
	case bfloat16.BFloat16:
		return dtypes.BFloat16
	case uint64:
		return dtypes.Uint64
	}
	return dtypes.InvalidDType
}

func (t *typedTensor[T]) Destroy() error {
	return t.tensor.Destroy()
}

func (t *typedTensor[T]) Value() ort.Value {
	return t.tensor
}

func (t *typedTensor[T]) ToFlatData(flat any) error {
	slice, ok := flat.([]T)
	if !ok {
		return errors.Errorf("flat data type %T does not match tensor type %T", flat, slice)
	}
	copy(slice, t.tensor.GetData())
	return nil
}

func (t *typedTensor[T]) CopyFrom(flat any) error {
	slice, ok := flat.([]T)
	if !ok {
		return errors.Errorf("flat data type %T does not match tensor type %T", flat, slice)
	}
	copy(t.tensor.GetData(), slice)
	return nil
}

func (t *typedTensor[T]) GetData() any {
	return t.tensor.GetData()
}

func float16ToFloat32(src []float16.Float16) []float32 {
	dst := make([]float32, len(src))
	for i, v := range src {
		dst[i] = v.Float32()
	}
	return dst
}

func float32ToFloat16(src []float32, dst []float16.Float16) {
	for i, v := range src {
		dst[i] = float16.FromFloat32(v)
	}
}

func bfloat16ToFloat32(src []bfloat16.BFloat16) []float32 {
	dst := make([]float32, len(src))
	for i, v := range src {
		dst[i] = v.Float32()
	}
	return dst
}

func float32ToBFloat16(src []float32, dst []bfloat16.BFloat16) {
	for i, v := range src {
		dst[i] = bfloat16.FromFloat32(v)
	}
}

type float16Tensor struct {
	tensor *ort.Tensor[float32]
}

func (t *float16Tensor) GetShape() ort.Shape    { return t.tensor.GetShape() }
func (t *float16Tensor) GetDType() dtypes.DType { return dtypes.Float16 }
func (t *float16Tensor) Destroy() error         { return t.tensor.Destroy() }
func (t *float16Tensor) Value() ort.Value       { return t.tensor }

func (t *float16Tensor) ToFlatData(flat any) error {
	slice, ok := flat.([]float16.Float16)
	if !ok {
		return errors.Errorf("flat data type %T does not match float16 slice", flat)
	}
	float32ToFloat16(t.tensor.GetData(), slice)
	return nil
}

func (t *float16Tensor) CopyFrom(flat any) error {
	slice, ok := flat.([]float16.Float16)
	if !ok {
		return errors.Errorf("flat data type %T does not match float16 slice", flat)
	}
	f32 := float16ToFloat32(slice)
	copy(t.tensor.GetData(), f32)
	return nil
}

func (t *float16Tensor) GetData() any {
	f32 := t.tensor.GetData()
	res := make([]float16.Float16, len(f32))
	float32ToFloat16(f32, res)
	return res
}

type bfloat16Tensor struct {
	tensor *ort.Tensor[float32]
}

func (t *bfloat16Tensor) GetShape() ort.Shape    { return t.tensor.GetShape() }
func (t *bfloat16Tensor) GetDType() dtypes.DType { return dtypes.BFloat16 }
func (t *bfloat16Tensor) Destroy() error         { return t.tensor.Destroy() }
func (t *bfloat16Tensor) Value() ort.Value       { return t.tensor }

func (t *bfloat16Tensor) ToFlatData(flat any) error {
	slice, ok := flat.([]bfloat16.BFloat16)
	if !ok {
		return errors.Errorf("flat data type %T does not match bfloat16 slice", flat)
	}
	float32ToBFloat16(t.tensor.GetData(), slice)
	return nil
}

func (t *bfloat16Tensor) CopyFrom(flat any) error {
	slice, ok := flat.([]bfloat16.BFloat16)
	if !ok {
		return errors.Errorf("flat data type %T does not match bfloat16 slice", flat)
	}
	f32 := bfloat16ToFloat32(slice)
	copy(t.tensor.GetData(), f32)
	return nil
}

func (t *bfloat16Tensor) GetData() any {
	f32 := t.tensor.GetData()
	res := make([]bfloat16.BFloat16, len(f32))
	float32ToBFloat16(f32, res)
	return res
}

func toInt64s(dims []int) []int64 {
	res := make([]int64, len(dims))
	for i, v := range dims {
		res[i] = int64(v)
	}
	return res
}

// NewOrtTensorWrapper creates an OrtTensorWrapper initialized with data.
func NewOrtTensorWrapper(shape shapes.Shape, flat any) (OrtTensorWrapper, error) {
	ortShape := ort.NewShape(toInt64s(shape.Dimensions)...)
	switch f := flat.(type) {
	case []float16.Float16:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[float16.Float16]{tensor: t, rawData: f}, nil
	case []bfloat16.BFloat16:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[bfloat16.BFloat16]{tensor: t, rawData: f}, nil
	case []float32:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[float32]{tensor: t, rawData: f}, nil
	case []float64:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[float64]{tensor: t, rawData: f}, nil
	case []int32:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[int32]{tensor: t, rawData: f}, nil
	case []int64:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[int64]{tensor: t, rawData: f}, nil
	case []bool:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[bool]{tensor: t, rawData: f}, nil
	case []int8:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[int8]{tensor: t, rawData: f}, nil
	case []uint8:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[uint8]{tensor: t, rawData: f}, nil
	case []int16:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[int16]{tensor: t, rawData: f}, nil
	case []uint16:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[uint16]{tensor: t, rawData: f}, nil
	case []uint32:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[uint32]{tensor: t, rawData: f}, nil
	case []uint64:
		t, err := ort.NewTensor(ortShape, f)
		if err != nil {
			return nil, err
		}
		return &typedTensor[uint64]{tensor: t, rawData: f}, nil
	default:
		return nil, errors.Errorf("unsupported type %T for onnxruntime tensor creation", flat)
	}
}

// NewEmptyOrtTensorWrapper creates an empty OrtTensorWrapper for the given shape.
func NewEmptyOrtTensorWrapper(shape shapes.Shape) (OrtTensorWrapper, error) {
	ortShape := ort.NewShape(toInt64s(shape.Dimensions)...)
	switch shape.DType {
	case dtypes.Float16:
		t, err := ort.NewEmptyTensor[float16.Float16](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[float16.Float16]{tensor: t}, nil
	case dtypes.BFloat16:
		t, err := ort.NewEmptyTensor[bfloat16.BFloat16](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[bfloat16.BFloat16]{tensor: t}, nil
	case dtypes.Float32:
		t, err := ort.NewEmptyTensor[float32](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[float32]{tensor: t}, nil
	case dtypes.Float64:
		t, err := ort.NewEmptyTensor[float64](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[float64]{tensor: t}, nil
	case dtypes.Int32:
		t, err := ort.NewEmptyTensor[int32](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[int32]{tensor: t}, nil
	case dtypes.Int64:
		t, err := ort.NewEmptyTensor[int64](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[int64]{tensor: t}, nil
	case dtypes.Bool:
		t, err := ort.NewEmptyTensor[bool](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[bool]{tensor: t}, nil
	case dtypes.Int8:
		t, err := ort.NewEmptyTensor[int8](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[int8]{tensor: t}, nil
	case dtypes.Uint8:
		t, err := ort.NewEmptyTensor[uint8](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[uint8]{tensor: t}, nil
	case dtypes.Int16:
		t, err := ort.NewEmptyTensor[int16](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[int16]{tensor: t}, nil
	case dtypes.Uint16:
		t, err := ort.NewEmptyTensor[uint16](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[uint16]{tensor: t}, nil
	case dtypes.Uint32:
		t, err := ort.NewEmptyTensor[uint32](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[uint32]{tensor: t}, nil
	case dtypes.Uint64:
		t, err := ort.NewEmptyTensor[uint64](ortShape)
		if err != nil {
			return nil, err
		}
		return &typedTensor[uint64]{tensor: t}, nil
	default:
		return nil, errors.Errorf("unsupported DType %s for empty onnxruntime tensor allocation", shape.DType)
	}
}

// GpuTensorWrapper wraps a raw GPU OrtValue.
type GpuTensorWrapper struct {
	val   ort.Value
	shape ort.Shape
	dtype dtypes.DType
}

// NewGpuTensorWrapper creates a GpuTensorWrapper from an existing GPU OrtValue.
func NewGpuTensorWrapper(val ort.Value, shape ort.Shape, dtype dtypes.DType) *GpuTensorWrapper {
	return &GpuTensorWrapper{val: val, shape: shape, dtype: dtype}
}

func (g *GpuTensorWrapper) GetShape() ort.Shape    { return g.shape }
func (g *GpuTensorWrapper) GetDType() dtypes.DType { return g.dtype }
func (g *GpuTensorWrapper) Value() ort.Value       { return g.val }

func (g *GpuTensorWrapper) Destroy() error {
	if g.val != nil {
		err := g.val.Destroy()
		g.val = nil
		return err
	}
	return nil
}

func (g *GpuTensorWrapper) size() int {
	s := 1
	for _, d := range g.shape {
		s *= int(d)
	}
	return s
}

func (g *GpuTensorWrapper) ToFlatData(flat any) error {
	if g.size() == 0 {
		return nil
	}
	switch f := flat.(type) {
	case []float16.Float16:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*2)
	case []bfloat16.BFloat16:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*2)
	case []float32:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*4)
	case []float64:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*8)
	case []int32:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*4)
	case []int64:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*8)
	case []bool:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*1)
	case []int8:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*1)
	case []uint8:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*1)
	case []int16:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*2)
	case []uint16:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*2)
	case []uint32:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*4)
	case []uint64:
		return ort.CopyGPUToHost(g.val, unsafe.Pointer(&f[0]), len(f)*8)
	default:
		return errors.Errorf("unsupported slice type %T for GPU ToFlatData", flat)
	}
}

func (g *GpuTensorWrapper) CopyFrom(flat any) error {
	if g.size() == 0 {
		return nil
	}
	switch f := flat.(type) {
	case []float16.Float16:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*2)
	case []bfloat16.BFloat16:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*2)
	case []float32:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*4)
	case []float64:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*8)
	case []int32:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*4)
	case []int64:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*8)
	case []bool:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*1)
	case []int8:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*1)
	case []uint8:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*1)
	case []int16:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*2)
	case []uint16:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*2)
	case []uint32:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*4)
	case []uint64:
		return ort.CopyHostToGPU(g.val, unsafe.Pointer(&f[0]), len(f)*8)
	default:
		return errors.Errorf("unsupported slice type %T for GPU CopyFrom", flat)
	}
}

func (g *GpuTensorWrapper) GetData() any {
	return nil
}

// Buffer implements compute.Buffer for ONNX Runtime native execution.
type Buffer struct {
	backend           compute.Backend
	wrapper           OrtTensorWrapper
	shape             shapes.Shape
	device            compute.DeviceNum
	executionProvider executionprovider.Type
	executable        *Executable
}

var _ compute.Buffer = (*Buffer)(nil)

// NewBuffer creates a new Buffer.
func NewBuffer(backend compute.Backend, wrapper OrtTensorWrapper, shape shapes.Shape, device compute.DeviceNum, executionProvider executionprovider.Type, exec *Executable) *Buffer {
	return &Buffer{
		backend:           backend,
		wrapper:           wrapper,
		shape:             shape,
		device:            device,
		executionProvider: executionProvider,
		executable:        exec,
	}
}

func (b *Buffer) Backend() compute.Backend {
	return b.backend
}

// IsShared returns whether the buffer supports direct host memory access.
func (b *Buffer) IsShared() bool {
	return b.executionProvider != executionprovider.CUDA
}

func (b *Buffer) Finalize() error {
	if b.wrapper != nil {
		w := b.wrapper
		b.wrapper = nil
		if b.executable != nil {
			b.executable.recycleWrapper(w)
		} else {
			if err := w.Destroy(); err != nil {
				return errors.Wrap(err, "failed to destroy onnxruntime tensor")
			}
		}
	}
	return nil
}

func (b *Buffer) Shape() (shapes.Shape, error) {
	return b.shape, nil
}

func (b *Buffer) DeviceNum() (compute.DeviceNum, error) {
	return b.device, nil
}

func (b *Buffer) ToFlatData(flat any) error {
	if b.wrapper == nil {
		return errors.New("cannot read from finalized buffer")
	}
	return b.wrapper.ToFlatData(flat)
}

func (b *Buffer) Data() (flat any, err error) {
	if !b.IsShared() {
		return nil, errors.New("direct data access not supported for GPU buffers; use ToFlatData")
	}
	if b.wrapper == nil {
		return nil, errors.New("cannot read from finalized buffer")
	}
	return b.wrapper.GetData(), nil
}

func (b *Buffer) CopyToDevice(deviceNum compute.DeviceNum) (compute.Buffer, error) {
	if b.wrapper == nil {
		return nil, errors.New("cannot copy finalized buffer")
	}
	newWrapper, err := NewEmptyOrtTensorWrapper(b.shape)
	if err != nil {
		return nil, err
	}
	flatSlice := dtypes.MakeAnySlice(b.shape.DType, b.shape.Size())
	err = b.wrapper.ToFlatData(flatSlice)
	if err != nil {
		_ = newWrapper.Destroy()
		return nil, err
	}
	err = newWrapper.CopyFrom(flatSlice)
	if err != nil {
		_ = newWrapper.Destroy()
		return nil, err
	}
	return &Buffer{
		backend:           b.backend,
		wrapper:           newWrapper,
		shape:             b.shape,
		device:            deviceNum,
		executionProvider: b.executionProvider,
	}, nil
}

// Wrapper returns the underlying OrtTensorWrapper.
func (b *Buffer) Wrapper() OrtTensorWrapper {
	return b.wrapper
}

// WrapOrtValue wraps a native ort.Value in an OrtTensorWrapper.
func WrapOrtValue(val ort.Value, shape shapes.Shape) (OrtTensorWrapper, error) {
	switch shape.DType {
	case dtypes.Float16:
		t, ok := val.(*ort.Tensor[float16.Float16])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[float16.Float16], got %T", val)
		}
		return &typedTensor[float16.Float16]{tensor: t}, nil
	case dtypes.BFloat16:
		t, ok := val.(*ort.Tensor[bfloat16.BFloat16])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[bfloat16.BFloat16], got %T", val)
		}
		return &typedTensor[bfloat16.BFloat16]{tensor: t}, nil
	case dtypes.Float32:
		t, ok := val.(*ort.Tensor[float32])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[float32], got %T", val)
		}
		return &typedTensor[float32]{tensor: t}, nil
	case dtypes.Float64:
		t, ok := val.(*ort.Tensor[float64])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[float64], got %T", val)
		}
		return &typedTensor[float64]{tensor: t}, nil
	case dtypes.Int32:
		t, ok := val.(*ort.Tensor[int32])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[int32], got %T", val)
		}
		return &typedTensor[int32]{tensor: t}, nil
	case dtypes.Int64:
		t, ok := val.(*ort.Tensor[int64])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[int64], got %T", val)
		}
		return &typedTensor[int64]{tensor: t}, nil
	case dtypes.Bool:
		t, ok := val.(*ort.Tensor[bool])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[bool], got %T", val)
		}
		return &typedTensor[bool]{tensor: t}, nil
	case dtypes.Int8:
		t, ok := val.(*ort.Tensor[int8])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[int8], got %T", val)
		}
		return &typedTensor[int8]{tensor: t}, nil
	case dtypes.Uint8:
		t, ok := val.(*ort.Tensor[uint8])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[uint8], got %T", val)
		}
		return &typedTensor[uint8]{tensor: t}, nil
	case dtypes.Int16:
		t, ok := val.(*ort.Tensor[int16])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[int16], got %T", val)
		}
		return &typedTensor[int16]{tensor: t}, nil
	case dtypes.Uint16:
		t, ok := val.(*ort.Tensor[uint16])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[uint16], got %T", val)
		}
		return &typedTensor[uint16]{tensor: t}, nil
	case dtypes.Uint32:
		t, ok := val.(*ort.Tensor[uint32])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[uint32], got %T", val)
		}
		return &typedTensor[uint32]{tensor: t}, nil
	case dtypes.Uint64:
		t, ok := val.(*ort.Tensor[uint64])
		if !ok {
			return nil, errors.Errorf("expected *ort.Tensor[uint64], got %T", val)
		}
		return &typedTensor[uint64]{tensor: t}, nil
	default:
		return nil, errors.Errorf("unsupported shape DType %s for wrapOrtValue", shape.DType)
	}
}
