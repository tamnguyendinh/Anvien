// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"math"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/dtypes/bfloat16"
	"github.com/gomlx/compute/dtypes/float16"
	"github.com/gomlx/compute/shapeinference"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

func (f *Function) ConvertDType(operand compute.Value, targetDType dtypes.DType) (compute.Value, error) {
	xNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	outShape := xNode.shape
	outShape.DType = targetDType

	node := &Node{
		opType: "Cast",
		inputs: []*Node{xNode},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "to",
				Type: onnx.AttributeProto_INT,
				I:    int64(dtypeToONNX(targetDType)),
			},
		},
	}
	return f.addNode(node), nil
}

func (f *Function) Where(condition compute.Value, onTrue compute.Value, onFalse compute.Value) (compute.Value, error) {
	condNode, ok1 := condition.(*Node)
	onTrueNode, ok2 := onTrue.(*Node)
	onFalseNode, ok3 := onFalse.(*Node)
	if !ok1 || !ok2 || !ok3 {
		return nil, errors.New("inputs must be valid onnxruntime nodes")
	}

	outShape, err := shapeinference.Where(condNode.shape, onTrueNode.shape, onFalseNode.shape)
	if err != nil {
		return nil, err
	}

	node := &Node{
		opType: "Where",
		inputs: []*Node{condNode, onTrueNode, onFalseNode},
		shape:  outShape,
	}
	return f.addNode(node), nil
}

func (f *Function) Reshape(x compute.Value, newDimensions ...int) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.Reshape(xNode.shape, newDimensions)
	if err != nil {
		return nil, err
	}

	newDims64 := make([]int64, len(newDimensions))
	for i, d := range newDimensions {
		newDims64[i] = int64(d)
	}

	shapeConstNode, err := f.Constant(newDims64, len(newDimensions))
	if err != nil {
		return nil, err
	}

	node := &Node{
		opType: "Reshape",
		inputs: []*Node{xNode, shapeConstNode.(*Node)},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "allowzero",
				Type: onnx.AttributeProto_INT,
				I:    1,
			},
		},
	}
	return f.addNode(node), nil
}

func (f *Function) Transpose(x compute.Value, permutation ...int) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.Transpose(xNode.shape, permutation)
	if err != nil {
		return nil, err
	}

	perm64 := make([]int64, len(permutation))
	for i, p := range permutation {
		perm64[i] = int64(p)
	}

	node := &Node{
		opType: "Transpose",
		inputs: []*Node{xNode},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "perm",
				Type: onnx.AttributeProto_INTS,
				Ints: perm64,
			},
		},
	}
	return f.addNode(node), nil
}

func (f *Function) BroadcastInDim(x compute.Value, outputShape shapes.Shape, broadcastAxes []int) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	err := shapeinference.BroadcastInDim(xNode.shape, outputShape, broadcastAxes, nil)
	if err != nil {
		return nil, err
	}

	reshapeDims := make([]int, outputShape.Rank())
	for i := range reshapeDims {
		reshapeDims[i] = 1
	}
	for i, axis := range broadcastAxes {
		reshapeDims[axis] = xNode.shape.Dimensions[i]
	}

	reshaped, err := f.Reshape(xNode, reshapeDims...)
	targetDims64 := make([]int64, outputShape.Rank())
	for i, d := range outputShape.Dimensions {
		targetDims64[i] = int64(d)
	}
	targetDimsNode, err := f.Constant(targetDims64, outputShape.Rank())
	if err != nil {
		return nil, err
	}

	outShape := outputShape
	outShape.DType = xNode.shape.DType
	node := &Node{
		opType: "Expand",
		inputs: []*Node{reshaped.(*Node), targetDimsNode.(*Node)},
		shape:  outShape,
	}
	return f.addNode(node), nil
}

func (f *Function) Concatenate(axis int, inputs ...compute.Value) (compute.Value, error) {
	nodeInputs := make([]*Node, len(inputs))
	shapeInputs := make([]shapes.Shape, len(inputs))
	for i, inp := range inputs {
		n, ok := inp.(*Node)
		if !ok {
			return nil, errors.New("inputs must be valid onnxruntime nodes")
		}
		nodeInputs[i] = n
		shapeInputs[i] = n.shape
	}

	outShape, err := shapeinference.Concatenate(shapeInputs, axis)
	if err != nil {
		return nil, err
	}

	node := &Node{
		opType: "Concat",
		inputs: nodeInputs,
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "axis",
				Type: onnx.AttributeProto_INT,
				I:    int64(axis),
			},
		},
	}
	return f.addNode(node), nil
}

func (f *Function) Slice(x compute.Value, start []int, limit []int, stride []int) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.Slice(xNode.shape, start, limit, stride)
	if err != nil {
		return nil, err
	}

	rank := xNode.shape.Rank()
	starts64 := make([]int64, rank)
	ends64 := make([]int64, rank)
	axes64 := make([]int64, rank)
	steps64 := make([]int64, rank)

	for i := range rank {
		starts64[i] = int64(start[i])
		if limit[i] == shapes.DynamicDim || (xNode.shape.Dimensions[i] == shapes.DynamicDim && limit[i] <= 0) {
			ends64[i] = math.MaxInt64
		} else {
			ends64[i] = int64(limit[i])
		}
		axes64[i] = int64(i)
		steps64[i] = int64(stride[i])
	}

	startsConst, err := f.Constant(starts64, rank)
	if err != nil {
		return nil, err
	}
	endsConst, err := f.Constant(ends64, rank)
	if err != nil {
		return nil, err
	}
	axesConst, err := f.Constant(axes64, rank)
	if err != nil {
		return nil, err
	}
	stepsConst, err := f.Constant(steps64, rank)
	if err != nil {
		return nil, err
	}

	node := &Node{
		opType: "Slice",
		inputs: []*Node{xNode, startsConst.(*Node), endsConst.(*Node), axesConst.(*Node), stepsConst.(*Node)},
		shape:  outShape,
	}
	return f.addNode(node), nil
}

func (f *Function) reverseForWebGPU(xNode *Node, axes []int) (compute.Value, error) {
	// Use Gather on WebGPU to avoid WebGPU Slice clamping bug on negative step.
	currentNode := xNode
	for _, a := range axes {
		effAxis := a
		if effAxis < 0 {
			effAxis += xNode.shape.Rank()
		}
		dim := xNode.shape.Dimensions[effAxis]
		if dim <= 0 {
			// Fallback to Slice for dynamic dimensions
			startsConst, err := f.Constant([]int64{-1}, 1)
			if err != nil {
				return nil, err
			}
			endsConst, err := f.Constant([]int64{math.MinInt32}, 1)
			if err != nil {
				return nil, err
			}
			axesConst, err := f.Constant([]int64{int64(effAxis)}, 1)
			if err != nil {
				return nil, err
			}
			stepsConst, err := f.Constant([]int64{-1}, 1)
			if err != nil {
				return nil, err
			}
			node := &Node{
				opType: "Slice",
				inputs: []*Node{currentNode, startsConst.(*Node), endsConst.(*Node), axesConst.(*Node), stepsConst.(*Node)},
				shape:  currentNode.shape,
			}
			currentNode = f.addNode(node)
			continue
		}

		revIndices := make([]int64, dim)
		for j := range dim {
			revIndices[j] = int64(dim - 1 - j)
		}
		indicesConst, err := f.Constant(revIndices, dim)
		if err != nil {
			return nil, err
		}

		node := &Node{
			opType: "Gather",
			inputs: []*Node{currentNode, indicesConst.(*Node)},
			shape:  currentNode.shape,
			attributes: []*onnx.AttributeProto{
				{
					Name: "axis",
					Type: onnx.AttributeProto_INT,
					I:    int64(effAxis),
				},
			},
		}
		currentNode = f.addNode(node)
	}
	return currentNode, nil
}

func (f *Function) Reverse(x compute.Value, axes ...int) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	if len(axes) == 0 {
		return xNode, nil
	}

	if f.executionProvider == executionprovider.WebGPU {
		return f.reverseForWebGPU(xNode, axes)
	}

	// Standard Slice for other backends (CPU, CUDA, WASM CPU)
	starts := make([]int64, len(axes))
	ends := make([]int64, len(axes))
	steps := make([]int64, len(axes))
	axes64 := make([]int64, len(axes))

	for i, a := range axes {
		starts[i] = -1
		ends[i] = math.MinInt64
		steps[i] = -1
		axes64[i] = int64(a)
	}

	startsConst, err := f.Constant(starts, len(axes))
	if err != nil {
		return nil, err
	}
	endsConst, err := f.Constant(ends, len(axes))
	if err != nil {
		return nil, err
	}
	axesConst, err := f.Constant(axes64, len(axes))
	if err != nil {
		return nil, err
	}
	stepsConst, err := f.Constant(steps, len(axes))
	if err != nil {
		return nil, err
	}

	node := &Node{
		opType: "Slice",
		inputs: []*Node{xNode, startsConst.(*Node), endsConst.(*Node), axesConst.(*Node), stepsConst.(*Node)},
		shape:  xNode.shape,
	}
	return f.addNode(node), nil
}

func (f *Function) DynamicShape(operand compute.Value) (compute.Value, error) {
	xNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	rank := xNode.shape.Rank()

	if !xNode.shape.IsDynamic() {
		dims32 := make([]int32, rank)
		for i, d := range xNode.shape.Dimensions {
			dims32[i] = int32(d)
		}
		return f.Constant(dims32, rank)
	}

	shape64 := shapes.Make(dtypes.Int64, rank)
	shapeNode64 := f.addNode(&Node{
		opType: "Shape",
		inputs: []*Node{xNode},
		shape:  shape64,
	})

	return f.ConvertDType(shapeNode64, dtypes.Int32)
}

func (f *Function) DynamicDimensionSize(operand compute.Value, axis int) (compute.Value, error) {
	xNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	rank := xNode.shape.Rank()
	if axis < 0 || axis >= rank {
		return nil, errors.Errorf("axis %d out of bounds for rank %d", axis, rank)
	}

	if !xNode.shape.IsDynamic() || xNode.shape.Dimensions[axis] != shapes.DynamicDim {
		return f.Constant([]int32{int32(xNode.shape.Dimensions[axis])})
	}

	dynShape, err := f.DynamicShape(xNode)
	if err != nil {
		return nil, err
	}

	sliced, err := f.Slice(dynShape, []int{axis}, []int{axis + 1}, []int{1})
	if err != nil {
		return nil, err
	}

	return f.Reshape(sliced)
}

func (f *Function) DynamicReshape(operand compute.Value, dimensions ...compute.DynamicDimensionSpec) (compute.Value, error) {
	xNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.DynamicReshape(xNode.shape, dimensions)
	if err != nil {
		return nil, err
	}

	hasDynamicValue := false
	for _, spec := range dimensions {
		if spec.Value != nil {
			hasDynamicValue = true
			break
		}
	}

	if !outShape.IsDynamic() && !hasDynamicValue {
		return f.Reshape(xNode, outShape.Dimensions...)
	}

	parts := make([]compute.Value, len(dimensions))
	for i, spec := range dimensions {
		if spec.Value != nil {
			valNode, ok := spec.Value.(*Node)
			if !ok {
				return nil, errors.New("dimension spec Value is not a valid onnxruntime node")
			}
			v := compute.Value(valNode)
			if valNode.shape.Rank() == 0 {
				var err error
				v, err = f.Reshape(v, 1)
				if err != nil {
					return nil, err
				}
			}
			currNode := v.(*Node)
			if currNode.shape.DType != dtypes.Int64 {
				var err error
				v, err = f.ConvertDType(v, dtypes.Int64)
				if err != nil {
					return nil, err
				}
			}
			parts[i] = v
		} else if spec.Static > 0 {
			constNode, err := f.Constant([]int64{int64(spec.Static)}, 1)
			if err != nil {
				return nil, err
			}
			parts[i] = constNode
		} else {
			// Inferred dimension: -1 in ONNX Reshape
			constNode, err := f.Constant([]int64{-1}, 1)
			if err != nil {
				return nil, err
			}
			parts[i] = constNode
		}
	}

	var shapeTensorNode compute.Value
	if len(parts) == 1 {
		shapeTensorNode = parts[0]
	} else {
		var err error
		shapeTensorNode, err = f.Concatenate(0, parts...)
		if err != nil {
			return nil, err
		}
	}

	shapeNode, ok := shapeTensorNode.(*Node)
	if !ok {
		return nil, errors.New("shape tensor is not a valid onnxruntime node")
	}

	node := &Node{
		opType: "Reshape",
		inputs: []*Node{xNode, shapeNode},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "allowzero",
				Type: onnx.AttributeProto_INT,
				I:    1,
			},
		},
	}
	return f.addNode(node), nil
}

func (f *Function) DynamicBroadcastInDim(operand compute.Value, broadcastAxes []int, dimensions ...compute.DynamicDimensionSpec) (compute.Value, error) {
	xNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.DynamicBroadcastInDim(xNode.shape, broadcastAxes, dimensions, nil)
	if err != nil {
		return nil, err
	}

	outputRank := len(dimensions)
	var reshaped compute.Value
	if xNode.shape.IsDynamic() {
		reshapeSpecs := make([]compute.DynamicDimensionSpec, outputRank)
		for i := range reshapeSpecs {
			reshapeSpecs[i] = compute.DynamicDimensionSpec{Static: 1}
		}
		for operandAxis, outputAxis := range broadcastAxes {
			dim := xNode.shape.Dimensions[operandAxis]
			name := xNode.shape.AxisName(operandAxis)
			if dim == shapes.DynamicDim {
				dimVal, err := f.DynamicDimensionSize(xNode, operandAxis)
				if err != nil {
					return nil, err
				}
				reshapeSpecs[outputAxis] = compute.DynamicDimensionSpec{
					Name:  name,
					Value: dimVal,
				}
			} else {
				reshapeSpecs[outputAxis] = compute.DynamicDimensionSpec{Static: dim}
			}
		}
		reshaped, err = f.DynamicReshape(xNode, reshapeSpecs...)
		if err != nil {
			return nil, err
		}
	} else {
		reshapeDims := make([]int, outputRank)
		for i := range reshapeDims {
			reshapeDims[i] = 1
		}
		for operandAxis, outputAxis := range broadcastAxes {
			reshapeDims[outputAxis] = xNode.shape.Dimensions[operandAxis]
		}
		reshaped, err = f.Reshape(xNode, reshapeDims...)
		if err != nil {
			return nil, err
		}
	}

	parts := make([]compute.Value, outputRank)
	for i, spec := range dimensions {
		if spec.Value != nil {
			valNode, ok := spec.Value.(*Node)
			if !ok {
				return nil, errors.New("dimension spec Value is not a valid onnxruntime node")
			}
			v := compute.Value(valNode)
			if valNode.shape.Rank() == 0 {
				var err error
				v, err = f.Reshape(v, 1)
				if err != nil {
					return nil, err
				}
			}
			currNode := v.(*Node)
			if currNode.shape.DType != dtypes.Int64 {
				var err error
				v, err = f.ConvertDType(v, dtypes.Int64)
				if err != nil {
					return nil, err
				}
			}
			parts[i] = v
		} else if spec.Static > 0 {
			constNode, err := f.Constant([]int64{int64(spec.Static)}, 1)
			if err != nil {
				return nil, err
			}
			parts[i] = constNode
		} else {
			resolved := false
			for operandAxis, outputAxis := range broadcastAxes {
				if outputAxis == i && xNode.shape.Dimensions[operandAxis] == shapes.DynamicDim {
					dimVal, err := f.DynamicDimensionSize(xNode, operandAxis)
					if err != nil {
						return nil, err
					}
					dimVal64, err := f.ConvertDType(dimVal, dtypes.Int64)
					if err != nil {
						return nil, err
					}
					reshapedDim, err := f.Reshape(dimVal64, 1)
					if err != nil {
						return nil, err
					}
					parts[i] = reshapedDim
					resolved = true
					break
				}
			}
			if !resolved {
				return nil, errors.Errorf("DynamicBroadcastInDim: axis %d requires a non-nil Value node or positive Static dimension", i)
			}
		}
	}

	var shapeTensorNode compute.Value
	if len(parts) == 1 {
		shapeTensorNode = parts[0]
	} else {
		var err error
		shapeTensorNode, err = f.Concatenate(0, parts...)
		if err != nil {
			return nil, err
		}
	}

	shapeNode, ok := shapeTensorNode.(*Node)
	if !ok {
		return nil, errors.New("shape tensor is not a valid onnxruntime node")
	}

	reshapedNode, ok := reshaped.(*Node)
	if !ok {
		return nil, errors.New("reshaped is not a valid onnxruntime node")
	}

	node := &Node{
		opType: "Expand",
		inputs: []*Node{reshapedNode, shapeNode},
		shape:  outShape,
	}
	return f.addNode(node), nil
}

func (f *Function) DynamicSlice(operand compute.Value, startIndices []compute.Value, sliceDims []int) (compute.Value, error) {
	opNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("DynamicSlice: operand must be a valid onnxruntime node")
	}

	rank := opNode.shape.Rank()
	if len(startIndices) == 1 && rank > 1 {
		singleNode, ok := startIndices[0].(*Node)
		if ok && singleNode.shape.Rank() == 1 && singleNode.shape.Dimensions[0] == rank {
			unpacked := make([]compute.Value, rank)
			for i := 0; i < rank; i++ {
				sliceVal, err := f.Slice(singleNode, []int{i}, []int{i + 1}, []int{1})
				if err != nil {
					return nil, err
				}
				reshapedVal, err := f.Reshape(sliceVal)
				if err != nil {
					return nil, err
				}
				unpacked[i] = reshapedVal
			}
			startIndices = unpacked
		}
	}
	if len(startIndices) != rank {
		return nil, errors.Errorf("DynamicSlice: startIndices length (%d) must match operand rank (%d)", len(startIndices), rank)
	}
	if len(sliceDims) != rank {
		return nil, errors.Errorf("DynamicSlice: sliceDims length (%d) must match operand rank (%d)", len(sliceDims), rank)
	}

	startsList := make([]compute.Value, rank)
	endsList := make([]compute.Value, rank)
	axesList := make([]compute.Value, rank)
	stepsList := make([]compute.Value, rank)

	for i := range rank {
		startNode, ok := startIndices[i].(*Node)
		if !ok {
			return nil, errors.Errorf("DynamicSlice: startIndices[%d] must be a valid onnxruntime node", i)
		}
		dimSize := opNode.shape.Dimensions[i]
		sliceSize := sliceDims[i]
		maxStart := dimSize - sliceSize

		var clampedStart compute.Value = startNode
		if maxStart >= 0 {
			maxConst, err := f.Constant([]int64{int64(maxStart)})
			if err != nil {
				return nil, err
			}
			maxConstCasted, err := f.ConvertDType(maxConst, startNode.shape.DType)
			if err != nil {
				return nil, err
			}
			minVal, err := f.Min(startNode, maxConstCasted)
			if err != nil {
				return nil, err
			}
			zeroConst, err := f.Constant([]int64{0})
			if err != nil {
				return nil, err
			}
			zeroConstCasted, err := f.ConvertDType(zeroConst, startNode.shape.DType)
			if err != nil {
				return nil, err
			}
			maxVal, err := f.Max(minVal, zeroConstCasted)
			if err != nil {
				return nil, err
			}
			clampedStart = maxVal
		}

		startInt64, err := f.ConvertDType(clampedStart, dtypes.Int64)
		if err != nil {
			return nil, err
		}
		startReshaped, err := f.Reshape(startInt64, 1)
		if err != nil {
			return nil, err
		}
		startsList[i] = startReshaped

		sliceSizeConst, err := f.Constant([]int64{int64(sliceSize)}, 1)
		if err != nil {
			return nil, err
		}
		endVal, err := f.Add(startReshaped, sliceSizeConst)
		if err != nil {
			return nil, err
		}
		endsList[i] = endVal

		axisConst, err := f.Constant([]int64{int64(i)}, 1)
		if err != nil {
			return nil, err
		}
		axesList[i] = axisConst

		stepConst, err := f.Constant([]int64{1}, 1)
		if err != nil {
			return nil, err
		}
		stepsList[i] = stepConst
	}

	startsTensor, err := f.Concatenate(0, startsList...)
	if err != nil {
		return nil, errors.Wrap(err, "DynamicSlice: concatenating starts failed")
	}
	endsTensor, err := f.Concatenate(0, endsList...)
	if err != nil {
		return nil, errors.Wrap(err, "DynamicSlice: concatenating ends failed")
	}
	axesTensor, err := f.Concatenate(0, axesList...)
	if err != nil {
		return nil, errors.Wrap(err, "DynamicSlice: concatenating axes failed")
	}
	stepsTensor, err := f.Concatenate(0, stepsList...)
	if err != nil {
		return nil, errors.Wrap(err, "DynamicSlice: concatenating steps failed")
	}

	outShape := shapes.Make(opNode.shape.DType, sliceDims...)
	node := &Node{
		opType: "Slice",
		inputs: []*Node{opNode, startsTensor.(*Node), endsTensor.(*Node), axesTensor.(*Node), stepsTensor.(*Node)},
		shape:  outShape,
	}
	return f.addNode(node), nil
}

func (f *Function) DynamicUpdateSlice(operand, update compute.Value, startIndices []compute.Value) (compute.Value, error) {
	opNode, ok1 := operand.(*Node)
	upNode, ok2 := update.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("DynamicUpdateSlice: inputs must be valid onnxruntime nodes")
	}

	rank := opNode.shape.Rank()
	if upNode.shape.Rank() != rank {
		return nil, errors.Errorf("DynamicUpdateSlice: operand rank (%d) must match update rank (%d)", rank, upNode.shape.Rank())
	}
	if len(startIndices) == 1 && rank > 1 {
		singleNode, ok := startIndices[0].(*Node)
		if ok && singleNode.shape.Rank() == 1 && singleNode.shape.Dimensions[0] == rank {
			unpacked := make([]compute.Value, rank)
			for i := 0; i < rank; i++ {
				sliceVal, err := f.Slice(singleNode, []int{i}, []int{i + 1}, []int{1})
				if err != nil {
					return nil, err
				}
				reshapedVal, err := f.Reshape(sliceVal)
				if err != nil {
					return nil, err
				}
				unpacked[i] = reshapedVal
			}
			startIndices = unpacked
		}
	}
	if len(startIndices) != rank {
		return nil, errors.Errorf("DynamicUpdateSlice: startIndices length (%d) must match rank (%d)", len(startIndices), rank)
	}

	padsBeforeList := make([]compute.Value, rank)
	padsAfterList := make([]compute.Value, rank)
	maskParts := make([]compute.Value, 0, rank)

	for i := range rank {
		startNode, ok := startIndices[i].(*Node)
		if !ok {
			return nil, errors.Errorf("DynamicUpdateSlice: startIndices[%d] must be a valid onnxruntime node", i)
		}

		N_i := opNode.shape.Dimensions[i]
		K_i := upNode.shape.Dimensions[i]
		maxStart := N_i - K_i

		var clampedStart compute.Value = startNode
		if maxStart > 0 {
			maxConst, err := f.Constant([]int64{int64(maxStart)})
			if err != nil {
				return nil, err
			}
			maxConstCasted, err := f.ConvertDType(maxConst, startNode.shape.DType)
			if err != nil {
				return nil, err
			}
			minVal, err := f.Min(startNode, maxConstCasted)
			if err != nil {
				return nil, err
			}
			zeroConst, err := f.Constant([]int64{0})
			if err != nil {
				return nil, err
			}
			zeroConstCasted, err := f.ConvertDType(zeroConst, startNode.shape.DType)
			if err != nil {
				return nil, err
			}
			maxVal, err := f.Max(minVal, zeroConstCasted)
			if err != nil {
				return nil, err
			}
			clampedStart = maxVal
		} else {
			zeroConst, err := f.Constant([]int64{0})
			if err != nil {
				return nil, err
			}
			zeroCasted, err := f.ConvertDType(zeroConst, startNode.shape.DType)
			if err != nil {
				return nil, err
			}
			clampedStart = zeroCasted
		}

		padBefore64, err := f.ConvertDType(clampedStart, dtypes.Int64)
		if err != nil {
			return nil, err
		}
		padBeforeReshaped, err := f.Reshape(padBefore64, 1)
		if err != nil {
			return nil, err
		}
		padsBeforeList[i] = padBeforeReshaped

		totalPaddings64, err := f.Constant([]int64{int64(N_i - K_i)}, 1)
		if err != nil {
			return nil, err
		}
		padAfterReshaped, err := f.Sub(totalPaddings64, padBeforeReshaped)
		if err != nil {
			return nil, err
		}
		padsAfterList[i] = padAfterReshaped

		if K_i < N_i {
			iotaShapeDims := make([]int, rank)
			for j := range rank {
				if j == i {
					iotaShapeDims[j] = N_i
				} else {
					iotaShapeDims[j] = 1
				}
			}
			idx_i, err := f.Iota(shapes.Make(dtypes.Int32, iotaShapeDims...), i)
			if err != nil {
				return nil, err
			}
			start32, err := f.ConvertDType(clampedStart, dtypes.Int32)
			if err != nil {
				return nil, err
			}
			ge, err := f.GreaterOrEqual(idx_i, start32)
			if err != nil {
				return nil, err
			}
			kConst32, err := f.Constant([]int32{int32(K_i)})
			if err != nil {
				return nil, err
			}
			end32, err := f.Add(start32, kConst32)
			if err != nil {
				return nil, err
			}
			lt, err := f.LessThan(idx_i, end32)
			if err != nil {
				return nil, err
			}
			mask_i, err := f.LogicalAnd(ge, lt)
			if err != nil {
				return nil, err
			}
			maskParts = append(maskParts, mask_i)
		}
	}

	allPads := append(padsBeforeList, padsAfterList...)
	padsTensor, err := f.Concatenate(0, allPads...)
	if err != nil {
		return nil, errors.Wrap(err, "DynamicUpdateSlice: concatenating pads failed")
	}

	var zeroConst compute.Value
	switch upNode.shape.DType {
	case dtypes.Float32:
		zeroConst, err = f.Constant([]float32{0})
	case dtypes.Float64:
		zeroConst, err = f.Constant([]float64{0})
	case dtypes.Int32:
		zeroConst, err = f.Constant([]int32{0})
	case dtypes.Int64:
		zeroConst, err = f.Constant([]int64{0})
	case dtypes.Bool:
		zeroConst, err = f.Constant([]bool{false})
	case dtypes.Float16:
		zeroConst, err = f.Constant([]float16.Float16{float16.FromFloat32(0)})
	case dtypes.BFloat16:
		zeroConst, err = f.Constant([]bfloat16.BFloat16{bfloat16.FromFloat32(0)})
	default:
		return nil, errors.Errorf("DynamicUpdateSlice: unsupported DType %s", upNode.shape.DType)
	}
	if err != nil {
		return nil, err
	}

	zeroScalar, err := f.Reshape(zeroConst)
	if err != nil {
		return nil, err
	}

	padNode := &Node{
		opType: "Pad",
		inputs: []*Node{upNode, padsTensor.(*Node), zeroScalar.(*Node)},
		shape:  opNode.shape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "mode",
				Type: onnx.AttributeProto_STRING,
				S:    []byte("constant"),
			},
		},
	}
	paddedUpdate := f.addNode(padNode)

	if len(maskParts) == 0 {
		return paddedUpdate, nil
	}

	var fullMask compute.Value = maskParts[0]
	for i := 1; i < len(maskParts); i++ {
		fullMask, err = f.LogicalAnd(fullMask, maskParts[i])
		if err != nil {
			return nil, err
		}
	}

	fullMaskReshaped, err := broadcastToShape(f, fullMask, opNode)
	if err != nil {
		return nil, errors.Wrap(err, "DynamicUpdateSlice: mask broadcast failed")
	}

	return f.Where(fullMaskReshaped, paddedUpdate, opNode)
}

// Iota constructs a Range operator for the iotaAxis dimension and then broadcasts using Expand (BroadcastInDim).
func (f *Function) Iota(shape shapes.Shape, iotaAxis int) (compute.Value, error) {
	n := shape.Dimensions[iotaAxis]

	rangeDType := shape.DType
	castNeeded := false

	switch shape.DType {
	case dtypes.Float32, dtypes.Float64, dtypes.Int32, dtypes.Int64:
		// Natively supported by ONNX Range
	case dtypes.Float16, dtypes.BFloat16:
		rangeDType = dtypes.Float32
		castNeeded = true
	case dtypes.Int8, dtypes.Int16, dtypes.Uint8, dtypes.Uint16, dtypes.Uint32, dtypes.Uint64, dtypes.Bool:
		rangeDType = dtypes.Int64
		castNeeded = true
	default:
		return nil, errors.Errorf("unsupported DType %s for Iota", shape.DType)
	}

	startNode, err := MakeScalar(f, 0, rangeDType)
	if err != nil {
		return nil, err
	}
	limitNode, err := MakeScalar(f, n, rangeDType)
	if err != nil {
		return nil, err
	}
	deltaNode, err := MakeScalar(f, 1, rangeDType)
	if err != nil {
		return nil, err
	}

	rangeNode := &Node{
		opType: "Range",
		inputs: []*Node{startNode.(*Node), limitNode.(*Node), deltaNode.(*Node)},
		shape:  shapes.Make(rangeDType, n),
	}
	var rangeVal compute.Value = f.addNode(rangeNode)

	if castNeeded {
		rangeVal, err = f.ConvertDType(rangeVal, shape.DType)
		if err != nil {
			return nil, err
		}
	}

	return f.BroadcastInDim(rangeVal, shape, []int{iotaAxis})
}

// Pad injects padding on the start, end, or interior of the given operand.
func (f *Function) Pad(x, fillValue compute.Value, axesConfig ...compute.PadAxis) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("Pad: input must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.Pad(xNode.shape, axesConfig...)
	if err != nil {
		return nil, err
	}

	rank := xNode.shape.Rank()
	for _, cfg := range axesConfig {
		if cfg.Interior != 0 {
			return nil, errors.Wrap(compute.ErrNotImplemented, "interior padding is not supported by ONNX Pad")
		}
	}

	padsList := make([]int64, 2*rank)
	for i := range rank {
		padBefore := 0
		padAfter := 0
		if i < len(axesConfig) {
			padBefore = axesConfig[i].Start
			padAfter = axesConfig[i].End
		}
		padsList[i] = int64(padBefore)
		padsList[i+rank] = int64(padAfter)
	}

	padsConst, err := f.Constant(padsList, 2*rank)
	if err != nil {
		return nil, err
	}

	padScalar, ok := fillValue.(*Node)
	if !ok {
		return nil, errors.New("Pad: fillValue must be a valid onnxruntime node")
	}
	if padScalar.shape.Rank() != 0 {
		var errReshape error
		res, errReshape := f.Reshape(padScalar)
		if errReshape != nil {
			return nil, errReshape
		}
		padScalar = res.(*Node)
	}

	padNode := &Node{
		opType: "Pad",
		inputs: []*Node{xNode, padsConst.(*Node), padScalar},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "mode",
				Type: onnx.AttributeProto_STRING,
				S:    []byte("constant"),
			},
		},
	}
	return f.addNode(padNode), nil
}

// DynamicIota constructs an ONNX Range node for dynamic/static iotaAxis and broadcasts it to target dimensions.
func (f *Function) DynamicIota(dtype dtypes.DType, iotaAxis int, dimensions ...compute.DynamicDimensionSpec) (compute.Value, error) {
	_, err := shapeinference.DynamicIota(dtype, iotaAxis, dimensions, nil)
	if err != nil {
		return nil, err
	}

	rangeDType := dtype
	castNeeded := false

	switch dtype {
	case dtypes.Float32, dtypes.Float64, dtypes.Int32, dtypes.Int64:
		// Natively supported by ONNX Range
	case dtypes.Float16, dtypes.BFloat16:
		rangeDType = dtypes.Float32
		castNeeded = true
	case dtypes.Int8, dtypes.Int16, dtypes.Uint8, dtypes.Uint16, dtypes.Uint32, dtypes.Uint64, dtypes.Bool:
		rangeDType = dtypes.Int64
		castNeeded = true
	default:
		return nil, errors.Errorf("unsupported DType %s for DynamicIota", dtype)
	}

	startNode, err := MakeScalar(f, 0, rangeDType)
	if err != nil {
		return nil, err
	}

	axisSpec := dimensions[iotaAxis]
	var limitNode compute.Value
	if axisSpec.Value != nil {
		limitNode, err = f.ConvertDType(axisSpec.Value, rangeDType)
		if err != nil {
			return nil, err
		}
		if limitNode.(*Node).shape.Rank() != 0 {
			limitNode, err = f.Reshape(limitNode)
			if err != nil {
				return nil, err
			}
		}
	} else if axisSpec.Static > 0 {
		limitNode, err = MakeScalar(f, axisSpec.Static, rangeDType)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.Errorf("DynamicIota: axis %d requires a non-nil Value node or positive Static dimension", iotaAxis)
	}

	deltaNode, err := MakeScalar(f, 1, rangeDType)
	if err != nil {
		return nil, err
	}

	dim1D := axisSpec.Static
	var rangeShape shapes.Shape
	if axisSpec.Name != "" {
		dim1D = shapes.DynamicDim
		rangeShape = shapes.MakeDynamic(rangeDType, []int{dim1D}, []string{axisSpec.Name})
	} else {
		rangeShape = shapes.Make(rangeDType, dim1D)
	}

	rangeNode := &Node{
		opType: "Range",
		inputs: []*Node{startNode.(*Node), limitNode.(*Node), deltaNode.(*Node)},
		shape:  rangeShape,
	}
	var rangeVal compute.Value = f.addNode(rangeNode)

	if castNeeded {
		rangeVal, err = f.ConvertDType(rangeVal, dtype)
		if err != nil {
			return nil, err
		}
	}

	return f.DynamicBroadcastInDim(rangeVal, []int{iotaAxis}, dimensions...)
}

// DynamicPad injects padding on the start, end, or interior of the given operand.
func (f *Function) DynamicPad(x, fillValue compute.Value, axesConfig ...compute.DynamicPadAxis) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("DynamicPad: input must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.DynamicPad(xNode.shape, axesConfig...)
	if err != nil {
		return nil, err
	}

	rank := xNode.shape.Rank()
	padsBeforeList := make([]compute.Value, rank)
	padsAfterList := make([]compute.Value, rank)

	zeroInt64, err := f.Constant([]int64{0}, 1)
	if err != nil {
		return nil, err
	}

	for i := range rank {
		if i < len(axesConfig) {
			cfg := axesConfig[i]
			if cfg.InteriorValue != nil || cfg.Interior != 0 {
				return nil, errors.Wrap(compute.ErrNotImplemented, "interior padding is not supported by ONNX Pad")
			}

			if cfg.StartValue != nil {
				start64, err := f.ConvertDType(cfg.StartValue, dtypes.Int64)
				if err != nil {
					return nil, err
				}
				if start64.(*Node).shape.Rank() != 1 {
					start64, err = f.Reshape(start64, 1)
					if err != nil {
						return nil, err
					}
				}
				padsBeforeList[i] = start64
			} else {
				c, err := f.Constant([]int64{int64(cfg.Start)}, 1)
				if err != nil {
					return nil, err
				}
				padsBeforeList[i] = c
			}

			if cfg.EndValue != nil {
				end64, err := f.ConvertDType(cfg.EndValue, dtypes.Int64)
				if err != nil {
					return nil, err
				}
				if end64.(*Node).shape.Rank() != 1 {
					end64, err = f.Reshape(end64, 1)
					if err != nil {
						return nil, err
					}
				}
				padsAfterList[i] = end64
			} else {
				c, err := f.Constant([]int64{int64(cfg.End)}, 1)
				if err != nil {
					return nil, err
				}
				padsAfterList[i] = c
			}
		} else {
			padsBeforeList[i] = zeroInt64
			padsAfterList[i] = zeroInt64
		}
	}

	allPads := append(padsBeforeList, padsAfterList...)
	padsTensor, err := f.Concatenate(0, allPads...)
	if err != nil {
		return nil, errors.Wrap(err, "DynamicPad: concatenating pads failed")
	}

	padScalar, ok := fillValue.(*Node)
	if !ok {
		return nil, errors.New("DynamicPad: fillValue must be a valid onnxruntime node")
	}
	if padScalar.shape.Rank() != 0 {
		res, errReshape := f.Reshape(padScalar)
		if errReshape != nil {
			return nil, errReshape
		}
		padScalar = res.(*Node)
	}

	padNode := &Node{
		opType: "Pad",
		inputs: []*Node{xNode, padsTensor.(*Node), padScalar},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "mode",
				Type: onnx.AttributeProto_STRING,
				S:    []byte("constant"),
			},
		},
	}
	return f.addNode(padNode), nil
}

// CumSum implements the ONNX CumSum operation.
func (f *Function) CumSum(operand compute.Value, axis int, options compute.CumSumOptions) (compute.Value, error) {
	if options.Reverse && f.executionProvider == executionprovider.WebGPU {
		// Decompose reverse:
		// FutureWork: if WebGPU adds support for reverse, we can remove this.
		revIn, err := f.Reverse(operand, axis)
		if err != nil {
			return nil, err
		}
		forwardOptions := options
		forwardOptions.Reverse = false
		sum, err := f.CumSum(revIn, axis, forwardOptions)
		if err != nil {
			return nil, err
		}
		return f.Reverse(sum, axis)
	}

	xNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("CumSum: operand must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.CumSum(xNode.shape, axis)
	if err != nil {
		return nil, err
	}

	originalDType := xNode.shape.DType
	needCast := originalDType == dtypes.BFloat16
	var cumSumInput *Node
	if needCast {
		castInput, err := f.ConvertDType(xNode, dtypes.Float32)
		if err != nil {
			return nil, err
		}
		cumSumInput = castInput.(*Node)
	} else {
		cumSumInput = xNode
	}

	axisConst, err := f.Constant([]int64{int64(axis)})
	if err != nil {
		return nil, err
	}

	exclusiveVal := int64(0)
	if options.Exclusive {
		exclusiveVal = 1
	}
	reverseVal := int64(0)
	if options.Reverse {
		reverseVal = 1
	}

	node := &Node{
		opType: "CumSum",
		inputs: []*Node{cumSumInput, axisConst.(*Node)},
		shape:  cumSumInput.shape.Clone(),
		attributes: []*onnx.AttributeProto{
			{
				Name: "exclusive",
				Type: onnx.AttributeProto_INT,
				I:    exclusiveVal,
			},
			{
				Name: "reverse",
				Type: onnx.AttributeProto_INT,
				I:    reverseVal,
			},
		},
	}
	resNode := f.addNode(node)
	if needCast {
		castBack, err := f.ConvertDType(resNode, originalDType)
		if err != nil {
			return nil, err
		}
		resCast := castBack.(*Node)
		resCast.shape = outShape
		return resCast, nil
	}
	resNode.shape = outShape
	return resNode, nil
}
