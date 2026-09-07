// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"

	"github.com/gomlx/compute"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

func (f *Function) SelectAndScatterMax(operand compute.Value, source compute.Value, windowDimensions []int, windowStrides []int, paddings [][2]int) (compute.Value, error) {
	opNode, ok1 := operand.(*Node)
	srcNode, ok2 := source.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("inputs must be valid onnxruntime nodes")
	}

	rank := opNode.shape.Rank()
	numSpatial := rank - 2

	// Determine channel axis from windowDimensions (must be 1 for batch axis 0, and 1 for channel axis).
	var channelAxis int
	if windowDimensions[0] == 1 && windowDimensions[1] == 1 {
		channelAxis = 1
	} else if windowDimensions[0] == 1 && windowDimensions[rank-1] == 1 {
		channelAxis = rank - 1
	} else {
		return nil, errors.Wrap(compute.ErrNotImplemented, "SelectAndScatterMax requires batch and channel window dimensions to be 1")
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

	needTranspose := false
	for i, p := range inputPerm {
		if p != i {
			needTranspose = true
			break
		}
	}

	var inpVal compute.Value = opNode
	var gradVal compute.Value = srcNode
	var err error

	if needTranspose {
		inpVal, err = f.Transpose(opNode, inputPerm...)
		if err != nil {
			return nil, errors.Wrap(err, "SelectAndScatterMax: operand transpose failed")
		}
		gradVal, err = f.Transpose(srcNode, inputPerm...)
		if err != nil {
			return nil, errors.Wrap(err, "SelectAndScatterMax: source transpose failed")
		}
	}

	// 1. Run MaxPool with 2 outputs: [pooled_out, max_indices]
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
		if len(windowStrides) > origSpatialAxis {
			onnxStrides[i] = int64(windowStrides[origSpatialAxis])
		} else {
			onnxStrides[i] = int64(windowDimensions[origSpatialAxis])
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

	inpNCHW := inpVal.(*Node)
	gradNCHW := gradVal.(*Node)

	f.nodeCount++
	maxPoolYName := fmt.Sprintf("node_%d_y", f.nodeCount)
	maxPoolIdxName := fmt.Sprintf("node_%d_idx", f.nodeCount)

	maxPoolNode := &Node{
		name:        maxPoolYName,
		opType:      "MaxPool",
		inputs:      []*Node{inpNCHW},
		outputNames: []string{maxPoolYName, maxPoolIdxName},
		shape:       gradNCHW.shape,
		attributes:  attributes,
	}
	f.nodes = append(f.nodes, maxPoolNode)

	indicesNode := &Node{
		name:   maxPoolIdxName,
		inputs: []*Node{maxPoolNode},
		shape:  shapes.Make(dtypes.Int64, gradNCHW.shape.Dimensions...),
	}

	// 2. Scatter gradNCHW into zeros using indicesNode via ScatterElements (along spatial dimension)
	// Flatten spatial dimensions: NCHW -> NC(H*W)
	n := inpNCHW.shape.Dimensions[0]
	c := inpNCHW.shape.Dimensions[1]
	spatialTotal := 1
	for i := 2; i < rank; i++ {
		spatialTotal *= inpNCHW.shape.Dimensions[i]
	}

	// Flatten grad and indices to [N, C, H_out*W_out]
	gradSpatialTotal := 1
	for i := 2; i < rank; i++ {
		gradSpatialTotal *= gradNCHW.shape.Dimensions[i]
	}

	gradFlat, err := f.Reshape(gradNCHW, n, c, gradSpatialTotal)
	if err != nil {
		return nil, err
	}
	indicesFlat, err := f.Reshape(indicesNode, n, c, gradSpatialTotal)
	if err != nil {
		return nil, err
	}

	// Wrap indices using Mod to ensure they are strictly within [0, spatialTotal - 1]
	spatialTotalConst, errC := f.Constant([]int64{int64(spatialTotal)})
	if errC != nil {
		return nil, errC
	}
	indicesFlatMod, errM := f.Rem(indicesFlat, spatialTotalConst)
	if errM != nil {
		return nil, errM
	}

	// Zero destination tensor of shape [N, C, H*W] matching inpNCHW dtype
	zerosNode, err := f.MakeZeros(shapes.Make(inpNCHW.shape.DType, n, c, spatialTotal))
	if err != nil {
		return nil, err
	}

	// ScatterElements along axis 2
	f.nodeCount++
	scatterNode := &Node{
		name:   fmt.Sprintf("node_%d", f.nodeCount),
		opType: "ScatterElements",
		inputs: []*Node{zerosNode.(*Node), indicesFlatMod.(*Node), gradFlat.(*Node)},
		shape:  zerosNode.(*Node).shape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "axis",
				Type: onnx.AttributeProto_INT,
				I:    2,
			},
		},
	}
	f.nodes = append(f.nodes, scatterNode)

	// Reshape back to NCHW: [N, C, H, W]
	scatterNCHW, err := f.Reshape(scatterNode, inpNCHW.shape.Dimensions...)
	if err != nil {
		return nil, err
	}

	// 3. Transpose back to original layout if needed
	if needTranspose {
		outputPerm := make([]int, rank)
		outputPerm[0] = 0
		outputPerm[channelAxis] = 1
		for i := range numSpatial {
			origSpatialAxis := inputPerm[2+i]
			outputPerm[origSpatialAxis] = 2 + i
		}
		return f.Transpose(scatterNCHW, outputPerm...)
	}

	return scatterNCHW, nil
}
