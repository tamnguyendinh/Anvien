// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapeinference"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

func (f *Function) getPadValue(reductionType compute.ReduceOpType, dtype dtypes.DType) (compute.Value, error) {
	switch reductionType {
	case compute.ReduceOpSum:
		return MakeScalar(f, 0, dtype)
	case compute.ReduceOpMax:
		return MakeScalar(f, dtype.LowestValue(), dtype)
	}
	return nil, errors.Wrapf(compute.ErrNotImplemented, "getPadValue for %s/%s", reductionType, dtype)
}

// reduceWindowWithSlices implements window reduction by decomposing the dilated window into
// multiple strided ONNX "Slice" operations (one slice per position in the window grid) and accumulating
// them element-wise (using f.Max for ReduceOpMax or f.Add for ReduceOpSum).
//
// Note: The number of Slice operations generated depends ONLY on the kernel window size (e.g. 3x3 = 9 slices),
// NOT on the input/image resolution. Each Slice extracts the full strided tensor in parallel on the GPU.
//
// This is used as a workaround for WebGPU, where ONNX Runtime Web's MaxPool WGSL shader has an upstream
// bug that ignores window dilations (treating them as dilation=1).
func (f *Function) reduceWindowWithSlices(
	opNode *Node,
	reductionType compute.ReduceOpType,
	windowDimensions, strides, windowDilations []int,
	paddings [][2]int,
	outShape shapes.Shape,
) (compute.Value, error) {
	rank := opNode.shape.Rank()
	effStrides := make([]int, rank)
	for i := range rank {
		if i < len(strides) && strides[i] > 0 {
			effStrides[i] = strides[i]
		} else if i < len(windowDimensions) && windowDimensions[i] > 0 {
			effStrides[i] = windowDimensions[i]
		} else {
			effStrides[i] = 1
		}
	}
	effWinDil := make([]int, rank)
	for i := range rank {
		if i < len(windowDilations) && windowDilations[i] > 0 {
			effWinDil[i] = windowDilations[i]
		} else {
			effWinDil[i] = 1
		}
	}
	effWinDims := make([]int, rank)
	for i := range rank {
		if i < len(windowDimensions) && windowDimensions[i] > 0 {
			effWinDims[i] = windowDimensions[i]
		} else {
			effWinDims[i] = 1
		}
	}

	current := opNode
	hasPadding := false
	for _, p := range paddings {
		if p[0] > 0 || p[1] > 0 {
			hasPadding = true
			break
		}
	}
	if hasPadding {
		padVal, errVal := f.getPadValue(reductionType, opNode.shape.DType)
		if errVal != nil {
			return nil, errVal
		}
		padAxes := make([]compute.PadAxis, rank)
		for i, p := range paddings {
			padAxes[i] = compute.PadAxis{Start: p[0], End: p[1]}
		}
		paddedVal, errPad := f.Pad(current, padVal, padAxes...)
		if errPad != nil {
			return nil, errPad
		}
		current = paddedVal.(*Node)
	}

	totalWindowElements := 1
	for _, w := range effWinDims {
		totalWindowElements *= w
	}

	getWindowIndices := func(idx int) []int {
		indices := make([]int, rank)
		rem := idx
		for i := rank - 1; i >= 0; i-- {
			w := effWinDims[i]
			indices[i] = rem % w
			rem /= w
		}
		return indices
	}

	axes64 := make([]int64, rank)
	steps64 := make([]int64, rank)
	for i := range rank {
		axes64[i] = int64(i)
		steps64[i] = int64(effStrides[i])
	}
	axesConst, err := f.Constant(axes64, rank)
	if err != nil {
		return nil, err
	}
	stepsConst, err := f.Constant(steps64, rank)
	if err != nil {
		return nil, err
	}

	var combined compute.Value
	for k := 0; k < totalWindowElements; k++ {
		winIdx := getWindowIndices(k)
		starts := make([]int64, rank)
		ends := make([]int64, rank)
		for i := range rank {
			offset := winIdx[i] * effWinDil[i]
			starts[i] = int64(offset)
			ends[i] = int64(offset + outShape.Dimensions[i]*effStrides[i])
		}
		startsConst, err := f.Constant(starts, rank)
		if err != nil {
			return nil, err
		}
		endsConst, err := f.Constant(ends, rank)
		if err != nil {
			return nil, err
		}

		sliceNode := f.addNode(&Node{
			opType: "Slice",
			inputs: []*Node{current, startsConst.(*Node), endsConst.(*Node), axesConst.(*Node), stepsConst.(*Node)},
			shape:  outShape,
		})

		if combined == nil {
			combined = sliceNode
		} else {
			var errCombine error
			switch reductionType {
			case compute.ReduceOpMax:
				combined, errCombine = f.Max(combined, sliceNode)
			case compute.ReduceOpSum:
				combined, errCombine = f.Add(combined, sliceNode)
			}
			if errCombine != nil {
				return nil, errCombine
			}
		}
	}

	return combined, nil
}

func (f *Function) ReduceWindow(
	operand compute.Value,
	reductionType compute.ReduceOpType,
	windowDimensions, strides, inputDilations, windowDilations []int,
	paddings [][2]int,
) (compute.Value, error) {
	opNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("operand must be a valid onnxruntime node")
	}

	outShape, err := shapeinference.ReduceWindow(opNode.shape, windowDimensions, strides, inputDilations, windowDilations, paddings)
	if err != nil {
		return nil, err
	}

	rank := opNode.shape.Rank()

	// Unsupported feature checks: ONNX pooling operators (MaxPool/AveragePool) do not support input dilations.
	for _, d := range inputDilations {
		if d > 1 {
			return nil, errors.Wrap(compute.ErrNotImplemented, "input dilations > 1 not supported by ONNX pooling operators")
		}
	}

	hasWindowDilation := false
	for _, d := range windowDilations {
		if d > 1 {
			hasWindowDilation = true
			break
		}
	}
	if hasWindowDilation && f.executionProvider == executionprovider.WebGPU {
		if f.LogSeverity() >= 0 && f.LogSeverity() <= 2 {
			klog.Infof("ReduceWindow implemented with individual slices for WebGPU")
		}
		return f.reduceWindowWithSlices(opNode, reductionType, windowDimensions, strides, windowDilations, paddings, outShape)
	}

	var ortOpType string
	switch reductionType {
	case compute.ReduceOpMax:
		ortOpType = "MaxPool"
	case compute.ReduceOpSum:
		ortOpType = "AveragePool"
	default:
		return nil, errors.Wrapf(compute.ErrNotImplemented, "ReduceWindow reduction type %s not implemented", reductionType)
	}

	// If operand rank < 3 (e.g. 1D vector), unsqueeze to 4D (1, 1, 1, L) for ONNX pooling.
	if rank < 3 {
		// Fill missing stride / windowDilation / padding defaults if nil
		effectiveWindow := windowDimensions
		effectiveStrides := strides
		if len(effectiveStrides) == 0 {
			effectiveStrides = effectiveWindow
		}
		effectiveWindowDilations := windowDilations
		if len(effectiveWindowDilations) == 0 {
			effectiveWindowDilations = make([]int, rank)
			for i := range effectiveWindowDilations {
				effectiveWindowDilations[i] = 1
			}
		}
		effectivePaddings := paddings
		if len(effectivePaddings) == 0 {
			effectivePaddings = make([][2]int, rank)
		}

		// Reshape rank -> 4D: (1, 1, ..., N)
		reshapeDims := []int{1, 1, 1, 1}
		win4D := []int{1, 1, 1, 1}
		stride4D := []int{1, 1, 1, 1}
		winDil4D := []int{1, 1, 1, 1}
		pad4D := [][2]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}}

		for i := 0; i < rank; i++ {
			dimIdx := 4 - rank + i
			reshapeDims[dimIdx] = opNode.shape.Dimensions[i]
			win4D[dimIdx] = effectiveWindow[i]
			stride4D[dimIdx] = effectiveStrides[i]
			winDil4D[dimIdx] = effectiveWindowDilations[i]
			pad4D[dimIdx] = effectivePaddings[i]
		}
		in4D, errR := f.Reshape(opNode, reshapeDims...)
		if errR != nil {
			return nil, errR
		}

		res4D, errP := f.ReduceWindow(in4D, reductionType, win4D, stride4D, nil, winDil4D, pad4D)
		if errP != nil {
			return nil, errP
		}

		// Reshape back to outShape
		return f.Reshape(res4D, outShape.Dimensions...)
	}

	numSpatial := rank - 2

	// Determine channel axis from windowDimensions (must be 1 for batch axis 0, and 1 for channel axis).
	var channelAxis int
	if windowDimensions[0] == 1 && windowDimensions[1] == 1 {
		channelAxis = 1
	} else if windowDimensions[0] == 1 && windowDimensions[rank-1] == 1 {
		channelAxis = rank - 1
	} else {
		return nil, errors.Wrap(compute.ErrNotImplemented, "ReduceWindow requires batch and channel window dimensions to be 1")
	}

	inputPerm := make([]int, rank)
	inputPerm[0] = 0           // Batch
	inputPerm[1] = channelAxis // Channel
	spatialIdx := 2
	for i := 1; i < rank; i++ {
		if i != channelAxis {
			inputPerm[spatialIdx] = i
			spatialIdx++
		}
	}

	needInputTranspose := false
	for i, p := range inputPerm {
		if p != i {
			needInputTranspose = true
			break
		}
	}

	var inpVal compute.Value = opNode
	if needInputTranspose {
		inpVal, err = f.Transpose(opNode, inputPerm...)
		if err != nil {
			return nil, errors.Wrap(err, "ReduceWindow: input transpose failed")
		}
	}

	origDType := opNode.shape.DType
	needDTypeCast := (ortOpType == "AveragePool" && !origDType.IsFloat())
	if needDTypeCast {
		castVal, errC := f.ConvertDType(inpVal, dtypes.Float32)
		if errC != nil {
			return nil, errC
		}
		inpVal = castVal
	}

	// 2. Prepare attributes for ONNX MaxPool / AveragePool (kernel_shape, pads, strides, dilations)
	kernelShape := make([]int64, numSpatial)
	for i := range numSpatial {
		origSpatialAxis := inputPerm[2+i]
		if len(windowDimensions) > origSpatialAxis {
			kernelShape[i] = int64(windowDimensions[origSpatialAxis])
		} else {
			kernelShape[i] = 1
		}
	}

	onnxPads := make([]int64, numSpatial*2)
	if len(paddings) > 0 {
		for i := range numSpatial {
			origSpatialAxis := inputPerm[2+i]
			if origSpatialAxis < len(paddings) {
				onnxPads[i] = int64(paddings[origSpatialAxis][0])            // begin
				onnxPads[i+numSpatial] = int64(paddings[origSpatialAxis][1]) // end
			}
		}
	}

	onnxStrides := make([]int64, numSpatial)
	for i := range numSpatial {
		origSpatialAxis := inputPerm[2+i]
		if len(strides) > origSpatialAxis {
			onnxStrides[i] = int64(strides[origSpatialAxis])
		} else {
			onnxStrides[i] = int64(windowDimensions[origSpatialAxis])
		}
	}

	onnxDilations := make([]int64, numSpatial)
	for i := range numSpatial {
		origSpatialAxis := inputPerm[2+i]
		if len(windowDilations) > origSpatialAxis {
			onnxDilations[i] = int64(windowDilations[origSpatialAxis])
		} else {
			onnxDilations[i] = 1
		}
	}

	attributes := []*onnx.AttributeProto{
		{
			Name: "kernel_shape",
			Type: onnx.AttributeProto_INTS,
			Ints: kernelShape,
		},
		{
			Name: "pads",
			Type: onnx.AttributeProto_INTS,
			Ints: onnxPads,
		},
		{
			Name: "strides",
			Type: onnx.AttributeProto_INTS,
			Ints: onnxStrides,
		},
	}

	if ortOpType == "MaxPool" {
		attributes = append(attributes, &onnx.AttributeProto{
			Name: "dilations",
			Type: onnx.AttributeProto_INTS,
			Ints: onnxDilations,
		})
	} else if ortOpType == "AveragePool" {
		attributes = append(attributes, &onnx.AttributeProto{
			Name: "count_include_pad",
			Type: onnx.AttributeProto_INT,
			I:    1,
		})
	}

	// 3. Intermediate NCHW output shape for Pool node
	poolOutDims := make([]int, rank)
	poolOutDims[0] = outShape.Dimensions[0]
	poolOutDims[1] = outShape.Dimensions[channelAxis]
	for i := range numSpatial {
		origSpatialAxis := inputPerm[2+i]
		poolOutDims[2+i] = outShape.Dimensions[origSpatialAxis]
	}
	poolDType := outShape.DType
	if needDTypeCast {
		poolDType = dtypes.Float32
	}
	poolOutShape := shapes.Make(poolDType, poolOutDims...)

	f.nodeCount++
	poolNode := &Node{
		name:       fmt.Sprintf("node_%d", f.nodeCount),
		opType:     ortOpType,
		inputs:     []*Node{inpVal.(*Node)},
		shape:      poolOutShape,
		attributes: attributes,
	}
	f.nodes = append(f.nodes, poolNode)

	// If reductionType is ReduceOpSum, AveragePool computes the mean over window elements.
	// Multiply by the window element count to get the sum.
	var poolResult compute.Value = poolNode
	if reductionType == compute.ReduceOpSum {
		windowElements := 1
		for _, k := range kernelShape {
			windowElements *= int(k)
		}
		if windowElements > 1 {
			scaleConst, errC := f.MakeScalar(windowElements, poolDType)
			if errC != nil {
				return nil, errC
			}
			poolResult, err = f.Mul(poolNode, scaleConst)
			if err != nil {
				return nil, err
			}
		}
	}

	if needDTypeCast {
		castBack, errC := f.ConvertDType(poolResult, origDType)
		if errC != nil {
			return nil, errC
		}
		poolResult = castBack
	}

	// 4. Transpose pool output back to target layout if needed
	outputPerm := make([]int, rank)
	outputPerm[0] = 0
	outputPerm[channelAxis] = 1
	for i := range numSpatial {
		origSpatialAxis := inputPerm[2+i]
		outputPerm[origSpatialAxis] = 2 + i
	}

	needOutputTranspose := false
	for i, p := range outputPerm {
		if p != i {
			needOutputTranspose = true
			break
		}
	}

	if needOutputTranspose {
		return f.Transpose(poolResult, outputPerm...)
	}
	return poolResult, nil
}
