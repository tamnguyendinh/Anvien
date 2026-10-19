// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/dtypes/bfloat16"
	"github.com/gomlx/compute/dtypes/float16"
	"github.com/gomlx/compute/notimplemented"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

type Function struct {
	notimplemented.Function
	name              string
	builder           *Builder
	parent            *Function
	executionProvider executionprovider.Type
	params            []*Node
	nodes             []*Node
	returns           []*Node
	nodeCount         int
	constCache        map[string]*Node
}

var _ compute.Function = (*Function)(nil)

func NewFunction(name string, builder *Builder) *Function {
	var ep executionprovider.Type
	if builder != nil {
		ep = builder.executionProvider
	}
	return &Function{
		Function: notimplemented.Function{
			ErrFn: func(op compute.OpType) error {
				return errors.Wrapf(compute.ErrNotImplemented, "%s (%d) not implemented for ONNX Runtime backend", op, op)
			},
		},
		name:              name,
		builder:           builder,
		executionProvider: ep,
		constCache:        make(map[string]*Node),
	}
}

func (f *Function) addNode(node *Node) *Node {
	f.nodeCount++
	if node.name == "" {
		if node.opType == "Constant" {
			node.name = fmt.Sprintf("const_%d", f.nodeCount)
		} else {
			node.name = fmt.Sprintf("node_%d", f.nodeCount)
		}
	}
	f.nodes = append(f.nodes, node)
	return node
}

func (f *Function) constCacheKey(shape shapes.Shape, flat any) string {
	var sb strings.Builder
	sb.WriteString(shape.String())
	sb.WriteByte('|')
	if flat != nil {
		raw := dtypes.UnsafeByteSliceFromAny(flat)
		sb.Write(raw)
	}
	return sb.String()
}

func (f *Function) Constant(flat any, dims ...int) (compute.Value, error) {
	valType := reflect.TypeOf(flat)
	if valType.Kind() != reflect.Slice && valType.Kind() != reflect.Array {
		return nil, errors.Errorf("Constant expects flat to be a slice or array, got %T", flat)
	}
	dtype := dtypes.FromGoType(valType.Elem())
	shape := shapes.Make(dtype, dims...)

	key := f.constCacheKey(shape, flat)
	if f.constCache == nil {
		f.constCache = make(map[string]*Node)
	}
	if cached, ok := f.constCache[key]; ok {
		return cached, nil
	}

	node := &Node{
		opType:    "Constant",
		shape:     shape,
		flatValue: flat,
	}
	n := f.addNode(node)
	f.constCache[key] = n
	return n, nil
}

// MakeScalar constructs a 0D scalar constant tensor for the given value and DType.
func MakeScalar(f *Function, value any, dtype dtypes.DType) (compute.Value, error) {
	if value == nil {
		return nil, errors.New("MakeScalar value cannot be nil")
	}
	targetType := dtype.GoType()
	if targetType == nil {
		return nil, errors.Errorf("unsupported DType %s for MakeScalar", dtype)
	}

	valVal := reflect.ValueOf(value)
	if valVal.Type() == targetType {
		slice := reflect.MakeSlice(reflect.SliceOf(targetType), 1, 1)
		slice.Index(0).Set(valVal)
		return f.Constant(slice.Interface())
	}

	var flat any
	switch dtype {
	case dtypes.Float16:
		switch v := value.(type) {
		case float16.Float16:
			flat = []float16.Float16{v}
		case bfloat16.BFloat16:
			flat = []float16.Float16{float16.FromFloat32(v.Float32())}
		default:
			var f32 float32
			switch {
			case valVal.Kind() == reflect.Float32 || valVal.Kind() == reflect.Float64:
				f32 = float32(valVal.Float())
			case valVal.CanInt():
				f32 = float32(valVal.Int())
			case valVal.CanUint():
				f32 = float32(valVal.Uint())
			default:
				return nil, errors.Errorf("cannot convert %T to Float16 in MakeScalar", value)
			}
			flat = []float16.Float16{float16.FromFloat32(f32)}
		}
	case dtypes.BFloat16:
		switch v := value.(type) {
		case bfloat16.BFloat16:
			flat = []bfloat16.BFloat16{v}
		case float16.Float16:
			flat = []bfloat16.BFloat16{bfloat16.FromFloat32(v.Float32())}
		default:
			var f32 float32
			switch {
			case valVal.Kind() == reflect.Float32 || valVal.Kind() == reflect.Float64:
				f32 = float32(valVal.Float())
			case valVal.CanInt():
				f32 = float32(valVal.Int())
			case valVal.CanUint():
				f32 = float32(valVal.Uint())
			default:
				return nil, errors.Errorf("cannot convert %T to BFloat16 in MakeScalar", value)
			}
			flat = []bfloat16.BFloat16{bfloat16.FromFloat32(f32)}
		}
	case dtypes.Bool:
		switch {
		case valVal.Kind() == reflect.Bool:
			flat = []bool{valVal.Bool()}
		case valVal.CanInt():
			flat = []bool{valVal.Int() != 0}
		case valVal.CanUint():
			flat = []bool{valVal.Uint() != 0}
		default:
			return nil, errors.Errorf("cannot convert %T to Bool in MakeScalar", value)
		}
	default:
		if !valVal.Type().ConvertibleTo(targetType) {
			return nil, errors.Errorf("cannot convert %T to %s in MakeScalar", value, dtype)
		}
		converted := valVal.Convert(targetType)
		slice := reflect.MakeSlice(reflect.SliceOf(targetType), 1, 1)
		slice.Index(0).Set(converted)
		flat = slice.Interface()
	}

	return f.Constant(flat)
}

// MakeScalar constructs a 0D scalar constant tensor on Function for the given value and DType.
func (f *Function) MakeScalar(value any, dtype dtypes.DType) (compute.Value, error) {
	return MakeScalar(f, value, dtype)
}

// MakeValues constructs a constant tensor with shape filled with value by broadcasting a scalar.
func (f *Function) MakeValues(value any, shape shapes.Shape) (compute.Value, error) {
	scalar, err := f.MakeScalar(value, shape.DType)
	if err != nil {
		return nil, err
	}
	if shape.IsScalar() {
		return scalar, nil
	}
	return f.BroadcastInDim(scalar, shape, nil)
}

// MakeZeros constructs a constant tensor with shape filled with zero by broadcasting a scalar zero.
func (f *Function) MakeZeros(shape shapes.Shape) (compute.Value, error) {
	return f.MakeValues(0, shape)
}

func (f *Function) Name() string {
	return f.name
}

func (f *Function) Parent() compute.Function {
	if f.parent != nil {
		return f.parent
	}
	return nil
}

func (f *Function) Builder() compute.Builder {
	return f.builder
}

// ExecutionProvider returns the execution provider configured for the function.
func (f *Function) ExecutionProvider() executionprovider.Type {
	return f.executionProvider
}

// SetExecutionProvider sets the execution provider configured for the function.
func (f *Function) SetExecutionProvider(ep executionprovider.Type) {
	f.executionProvider = ep
}

func (f *Function) LogSeverity() int {
	if f.builder != nil {
		return f.builder.LogSeverity()
	}
	return -1
}

func (f *Function) Shape(v compute.Value) (shapes.Shape, error) {
	node, ok := v.(*Node)
	if !ok {
		return shapes.Invalid(), errors.New("value is not a valid onnxruntime node")
	}
	return node.shape, nil
}

func (f *Function) Parameter(name string, shape shapes.Shape, spec *compute.ShardingSpec) (compute.Value, error) {
	if name == "" {
		f.nodeCount++
		name = fmt.Sprintf("param_%d", f.nodeCount)
	}
	node := &Node{
		name:   name,
		opType: "Parameter",
		shape:  shape,
	}
	f.params = append(f.params, node)
	f.nodes = append(f.nodes, node)
	return node, nil
}

func (f *Function) Return(outputs []compute.Value, shardings []*compute.ShardingSpec) error {
	f.returns = make([]*Node, len(outputs))
	for i, output := range outputs {
		node, ok := output.(*Node)
		if !ok {
			return errors.New("return value is not a valid onnxruntime node")
		}
		f.returns[i] = node
	}
	return nil
}

func (f *Function) Identity(x compute.Value) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("identity input is not a valid onnxruntime node")
	}

	node := &Node{
		opType: "Identity",
		inputs: []*Node{xNode},
		shape:  xNode.shape,
	}
	return f.addNode(node), nil
}

// AddCustomNode adds a custom node to the function graph (useful for testing or custom ops).
func (f *Function) AddCustomNode(opType string, inputs []*Node, shape shapes.Shape) *Node {
	node := &Node{
		opType: opType,
		inputs: inputs,
		shape:  shape,
	}
	return f.addNode(node)
}

// Nodes returns the recorded nodes of the function.
func (f *Function) Nodes() []*Node {
	return f.nodes
}

// Params returns the parameters of the function.
func (f *Function) Params() []*Node {
	return f.params
}

// Returns returns the return nodes of the function.
func (f *Function) Returns() []*Node {
	return f.returns
}
