// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"
	"slices"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapeinference"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

func (f *Function) Reduce(x compute.Value, opType compute.OpType, axes []int, keepDims bool) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	originalDType := xNode.shape.DType
	isUnsigned := originalDType == dtypes.Uint8 || originalDType == dtypes.Uint16 || originalDType == dtypes.Uint32 || originalDType == dtypes.Uint64
	isBFloat16MinMax := (opType == compute.OpTypeReduceMin || opType == compute.OpTypeReduceMax) && originalDType == dtypes.BFloat16

	var reduceInput *Node
	if isUnsigned {
		castInput, err := f.ConvertDType(xNode, dtypes.Int64)
		if err != nil {
			return nil, err
		}
		reduceInput = castInput.(*Node)
	} else if isBFloat16MinMax {
		castInput, err := f.ConvertDType(xNode, dtypes.Float32)
		if err != nil {
			return nil, err
		}
		reduceInput = castInput.(*Node)
	} else {
		reduceInput = xNode
	}

	axesToUse := axes
	if len(axesToUse) == 0 {
		axesToUse = make([]int, reduceInput.shape.Rank())
		for i := 0; i < reduceInput.shape.Rank(); i++ {
			axesToUse[i] = i
		}
	}

	outShape, err := shapeinference.Reduce(reduceInput.shape, axesToUse)
	if err != nil {
		return nil, err
	}

	if keepDims {
		outShape.Dimensions = make([]int, reduceInput.shape.Rank())
		for axis, dim := range reduceInput.shape.Dimensions {
			reduced := slices.Contains(axesToUse, axis)
			if reduced {
				outShape.Dimensions[axis] = 1
			} else {
				outShape.Dimensions[axis] = dim
			}
		}
		if reduceInput.shape.AxisNames != nil {
			outShape.AxisNames = make([]string, reduceInput.shape.Rank())
			copy(outShape.AxisNames, reduceInput.shape.AxisNames)
		}
	}

	var ortOpType string
	switch opType {
	case compute.OpTypeReduceMin:
		ortOpType = "ReduceMin"
	case compute.OpTypeReduceMax:
		ortOpType = "ReduceMax"
	case compute.OpTypeReduceSum:
		ortOpType = "ReduceSum"
	case compute.OpTypeReduceProduct:
		ortOpType = "ReduceProd"
	default:
		return nil, errors.Errorf("unsupported reduction operation %s", opType)
	}

	axes64 := make([]int64, len(axesToUse))
	for i, a := range axesToUse {
		axes64[i] = int64(a)
	}
	axesConst, err := f.Constant(axes64, len(axesToUse))
	if err != nil {
		return nil, err
	}

	keepDimsVal := int64(0)
	if keepDims {
		keepDimsVal = 1
	}

	f.nodeCount++
	reduceNode := &Node{
		name:   fmt.Sprintf("node_%d", f.nodeCount),
		opType: ortOpType,
		inputs: []*Node{reduceInput, axesConst.(*Node)},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "keepdims",
				Type: onnx.AttributeProto_INT,
				I:    keepDimsVal,
			},
		},
	}
	f.nodes = append(f.nodes, reduceNode)

	var finalNode *Node = reduceNode
	if !keepDims && outShape.Rank() == 0 {
		newDimsConst, err := f.Constant([]int64{}, 0)
		if err != nil {
			return nil, err
		}

		f.nodeCount++
		reshapeNode := &Node{
			name:   fmt.Sprintf("node_%d", f.nodeCount),
			opType: "Reshape",
			inputs: []*Node{reduceNode, newDimsConst.(*Node)},
			shape:  outShape,
		}
		f.nodes = append(f.nodes, reshapeNode)
		finalNode = reshapeNode
	}

	if isUnsigned || isBFloat16MinMax {
		castBack, err := f.ConvertDType(finalNode, originalDType)
		if err != nil {
			return nil, err
		}
		return castBack, nil
	}

	return finalNode, nil
}

func (f *Function) ReduceLogical(x compute.Value, opType compute.OpType, axes []int, keepDims bool) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	xInt32, err := f.ConvertDType(xNode, dtypes.Int32)
	if err != nil {
		return nil, err
	}

	var reduced compute.Value
	switch opType {
	case compute.OpTypeReduceLogicalAnd:
		reduced, err = f.Reduce(xInt32, compute.OpTypeReduceProduct, axes, keepDims)
	case compute.OpTypeReduceLogicalOr:
		reduced, err = f.Reduce(xInt32, compute.OpTypeReduceMax, axes, keepDims)
	case compute.OpTypeReduceLogicalXor:
		sum, err := f.Reduce(xInt32, compute.OpTypeReduceSum, axes, keepDims)
		if err != nil {
			return nil, err
		}
		twoConst, err := f.Constant([]int32{2}, 1)
		if err != nil {
			return nil, err
		}
		div2, err := f.Div(sum, twoConst)
		if err != nil {
			return nil, err
		}
		mul2, err := f.Mul(div2, twoConst)
		if err != nil {
			return nil, err
		}
		mod2, err := f.Sub(sum, mul2)
		if err != nil {
			return nil, err
		}
		zeroConst, err := f.Constant([]int32{0}, 1)
		if err != nil {
			return nil, err
		}
		reduced, err = f.NotEqual(mod2, zeroConst)
	default:
		return nil, errors.Errorf("unsupported logical reduction type %s", opType)
	}
	if err != nil {
		return nil, err
	}

	if opType == compute.OpTypeReduceLogicalXor {
		return reduced, nil
	}
	return f.ConvertDType(reduced, dtypes.Bool)
}

func (f *Function) ReduceMin(x compute.Value, axes ...int) (compute.Value, error) {
	return f.Reduce(x, compute.OpTypeReduceMin, axes, false)
}

func (f *Function) ReduceMax(x compute.Value, axes ...int) (compute.Value, error) {
	return f.Reduce(x, compute.OpTypeReduceMax, axes, false)
}

func (f *Function) ReduceSum(x compute.Value, axes ...int) (compute.Value, error) {
	return f.Reduce(x, compute.OpTypeReduceSum, axes, false)
}

func (f *Function) ReduceProduct(x compute.Value, axes ...int) (compute.Value, error) {
	return f.Reduce(x, compute.OpTypeReduceProduct, axes, false)
}

func (f *Function) ReduceLogicalAnd(x compute.Value, axes ...int) (compute.Value, error) {
	return f.ReduceLogical(x, compute.OpTypeReduceLogicalAnd, axes, false)
}

func (f *Function) ReduceLogicalOr(x compute.Value, axes ...int) (compute.Value, error) {
	return f.ReduceLogical(x, compute.OpTypeReduceLogicalOr, axes, false)
}

func (f *Function) ReduceLogicalXor(x compute.Value, axes ...int) (compute.Value, error) {
	return f.ReduceLogical(x, compute.OpTypeReduceLogicalXor, axes, false)
}

func (f *Function) ArgMinMax(x compute.Value, axis int, outputDType dtypes.DType, isMin bool) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.ArgMinMax(xNode.shape, axis, outputDType)
	if err != nil {
		return nil, err
	}

	ortOpType := "ArgMax"
	if isMin {
		ortOpType = "ArgMin"
	}

	effInput := xNode
	effAxis := axis
	needsReshape := f.executionProvider == executionprovider.WebGPU && xNode.shape.Rank() == 1
	var int64Shape shapes.Shape
	if needsReshape {
		// Reshape 1D [N] to 2D [1, N] to avoid WebGPU WGSL bug where rank-1 indices are scalar u32 and cannot be indexed.
		reshapedIn, err := f.Reshape(xNode, 1, xNode.shape.Dimensions[0])
		if err != nil {
			return nil, err
		}
		effInput = reshapedIn.(*Node)
		effAxis = 1
		int64Shape = shapes.Make(dtypes.Int64, 1)
	} else {
		int64Shape = shapes.Make(dtypes.Int64, outShape.Dimensions...)
	}

	f.nodeCount++
	argMinMaxNode := &Node{
		name:   fmt.Sprintf("node_%d", f.nodeCount),
		opType: ortOpType,
		inputs: []*Node{effInput},
		shape:  int64Shape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "axis",
				Type: onnx.AttributeProto_INT,
				I:    int64(effAxis),
			},
			{
				Name: "keepdims",
				Type: onnx.AttributeProto_INT,
				I:    0,
			},
			{
				Name: "select_last_index",
				Type: onnx.AttributeProto_INT,
				I:    0,
			},
		},
	}
	f.nodes = append(f.nodes, argMinMaxNode)

	var res compute.Value = argMinMaxNode
	if needsReshape {
		var errReshape error
		res, errReshape = f.Reshape(argMinMaxNode)
		if errReshape != nil {
			return nil, errReshape
		}
	}

	// ONNX ArgMin/ArgMax produces INT64 output. Convert to outputDType if needed.
	if outputDType != dtypes.Int64 {
		return f.ConvertDType(res, outputDType)
	}

	return res, nil
}
