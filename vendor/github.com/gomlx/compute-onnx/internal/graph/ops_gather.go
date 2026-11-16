// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"
	"slices"

	"github.com/gomlx/compute"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

func init() {
	registerOp(compute.OpTypeGather)
}

func makeShape(dtype dtypes.DType, dims []int, axisNames []string) shapes.Shape {
	isDynamic := slices.Contains(dims, shapes.DynamicDim)
	if isDynamic {
		newNames := make([]string, len(dims))
		if len(axisNames) == len(dims) {
			copy(newNames, axisNames)
		}
		for i, d := range dims {
			if d == shapes.DynamicDim && newNames[i] == "" {
				newNames[i] = fmt.Sprintf("axis_%d", i)
			}
		}
		return shapes.MakeDynamic(dtype, dims, newNames)
	}
	return shapes.Make(dtype, dims...)
}

// Gather gathers slices from operand using coordinates specified by startIndices.
// It maps the generalized XLA Gather semantics to ONNX GatherND, Reshape, and Transpose.
func (f *Function) Gather(
	operand compute.Value,
	startIndices compute.Value,
	indexVectorAxis int,
	offsetOutputAxes []int,
	collapsedSliceAxes []int,
	startIndexMap []int,
	sliceSizes []int,
	indicesAreSorted bool,
) (compute.Value, error) {
	operandNode, ok1 := operand.(*Node)
	startIndicesNode, ok2 := startIndices.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("Gather: inputs must be valid onnxruntime nodes")
	}

	operandShape := operandNode.shape
	startIndicesShape := startIndicesNode.shape
	Q := startIndicesShape.Rank()

	// 1. Prepare indices to move the indexVectorAxis to the last axis.
	var indicesNode *Node
	if indexVectorAxis == Q {
		// Virtual indexVectorAxis of size 1. Append a dimension of size 1.
		newDims := make([]int, Q+1)
		copy(newDims, startIndicesShape.Dimensions)
		newDims[Q] = 1
		indicesNodeVal, err := f.Reshape(startIndicesNode, newDims...)
		if err != nil {
			return nil, errors.Wrap(err, "Gather: failed to reshape startIndices to append virtual vector axis")
		}
		indicesNode = indicesNodeVal.(*Node)
	} else if indexVectorAxis == Q-1 {
		// Already at the end.
		indicesNode = startIndicesNode
	} else {
		// Move indexVectorAxis to the end.
		perm := make([]int, Q)
		idx := 0
		for i := range Q {
			if i != indexVectorAxis {
				perm[idx] = i
				idx++
			}
		}
		perm[Q-1] = indexVectorAxis
		indicesNodeVal, err := f.Transpose(startIndicesNode, perm...)
		if err != nil {
			return nil, errors.Wrap(err, "Gather: failed to transpose startIndices to move vector axis to end")
		}
		indicesNode = indicesNodeVal.(*Node)
	}

	if indicesNode.shape.DType != dtypes.Int64 {
		castedIndices, err := f.ConvertDType(indicesNode, dtypes.Int64)
		if err != nil {
			return nil, errors.Wrap(err, "Gather: failed to cast indices to Int64")
		}
		indicesNode = castedIndices.(*Node)
	}

	V := indicesNode.shape.Dimensions[indicesNode.shape.Rank()-1]

	// Clamp indices to [0, dim-sliceSize] to match XLA / StableHLO semantics.
	if !operandShape.IsDynamic() {
		zeroConst, errZ := f.MakeScalar(int64(0), dtypes.Int64)
		if errZ != nil {
			return nil, errZ
		}
		maxBounds := make([]int64, V)
		for i, axis := range startIndexMap {
			dim := operandShape.Dimensions[axis]
			sliceSize := sliceSizes[axis]
			maxIdx := max(0, dim-sliceSize)
			maxBounds[i] = int64(maxIdx)
		}
		maxConst, errM := f.Constant(maxBounds, len(maxBounds))
		if errM != nil {
			return nil, errM
		}
		if indicesNode.shape.Rank() > 1 {
			bCastVal, errB := broadcastToShape(f, maxConst, indicesNode)
			if errB != nil {
				return nil, errB
			}
			maxConst = bCastVal
		}
		clampedMin, errC1 := f.Max(indicesNode, zeroConst)
		if errC1 != nil {
			return nil, errC1
		}
		clampedMax, errC2 := f.Min(clampedMin, maxConst)
		if errC2 != nil {
			return nil, errC2
		}
		indicesNode = clampedMax.(*Node)
	}

	// Check if slice sizes for mapped axes are supported by GatherND (GatherND only gathers scalar indices along mapped axes, i.e. sliceSize == 1).
	for _, axis := range startIndexMap {
		if sliceSizes[axis] != 1 {
			return nil, errors.Wrapf(compute.ErrNotImplemented, "Gather with sliceSize > 1 on mapped axis %d (sliceSize=%d) is not implemented for ONNX backend", axis, sliceSizes[axis])
		}
	}

	// 2. Prepare operand by transposing the mapped dimensions to the front.
	mappedSet := make(map[int]bool)
	for _, axis := range startIndexMap {
		mappedSet[axis] = true
	}
	var unmapped []int
	for axis := 0; axis < operandShape.Rank(); axis++ {
		if !mappedSet[axis] {
			unmapped = append(unmapped, axis)
		}
	}

	operandPerm := append(startIndexMap, unmapped...)
	transposedOperandVal, err := f.Transpose(operandNode, operandPerm...)
	if err != nil {
		return nil, errors.Wrap(err, "Gather: failed to transpose operand to place mapped axes first")
	}
	transposedOperand := transposedOperandVal.(*Node)

	// 3. Perform GatherND on the transposed operand.
	K := indicesNode.shape.Rank() - 1
	gatherNDRank := K + len(unmapped)
	gatherNDDims := make([]int, gatherNDRank)
	copy(gatherNDDims, indicesNode.shape.Dimensions[:K])
	copy(gatherNDDims[K:], transposedOperand.shape.Dimensions[V:])
	axisNames := make([]string, gatherNDRank)
	if indicesNode.shape.AxisNames != nil {
		copy(axisNames, indicesNode.shape.AxisNames[:K])
	}
	if transposedOperand.shape.AxisNames != nil {
		copy(axisNames[K:], transposedOperand.shape.AxisNames[V:])
	}
	gatherNDShape := makeShape(operandShape.DType, gatherNDDims, axisNames)

	gatherNDNode := f.addNode(&Node{
		opType: "GatherND",
		inputs: []*Node{transposedOperand, indicesNode},
		shape:  gatherNDShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "batch_dims",
				Type: onnx.AttributeProto_INT,
				I:    0,
			},
		},
	})

	// 4. Apply unmapped slicing if needed (if sliceSizes[u] < operandShape.Dimensions[u]).
	for idx, u := range unmapped {
		expectedSize := sliceSizes[u]
		actualSize := transposedOperand.shape.Dimensions[V+idx]
		if expectedSize < actualSize {
			start := make([]int, gatherNDRank)
			limit := make([]int, gatherNDRank)
			stride := make([]int, gatherNDRank)
			for i := range gatherNDRank {
				start[i] = 0
				limit[i] = gatherNDNode.shape.Dimensions[i]
				stride[i] = 1
			}
			limit[K+idx] = expectedSize
			slicedVal, err := f.Slice(gatherNDNode, start, limit, stride)
			if err != nil {
				return nil, errors.Wrap(err, "Gather: failed to slice unmapped dimension to match slice size")
			}
			gatherNDNode = slicedVal.(*Node)
		}
	}

	// 5. Reshape to B + O shape.
	// B are the batch dimensions of startIndices (excluding vector axis).
	// O are the non-collapsed slice dimensions.
	var O []int
	collapsedSet := make(map[int]bool)
	for _, axis := range collapsedSliceAxes {
		collapsedSet[axis] = true
	}
	for axis := 0; axis < operandShape.Rank(); axis++ {
		if !collapsedSet[axis] {
			O = append(O, sliceSizes[axis])
		}
	}

	B := indicesNode.shape.Dimensions[:K]
	reshapedDims := make([]int, len(B)+len(O))
	copy(reshapedDims, B)
	copy(reshapedDims[len(B):], O)

	var reshapedVal compute.Value
	isReshapeDynamic := slices.Contains(reshapedDims, shapes.DynamicDim)

	if isReshapeDynamic {
		specs := make([]compute.DynamicDimensionSpec, len(reshapedDims))
		for i := range B {
			if B[i] == shapes.DynamicDim {
				dimSize, err := f.DynamicDimensionSize(indicesNode, i)
				if err != nil {
					return nil, errors.Wrap(err, "Gather: failed to get dynamic dimension size for batch axis")
				}
				name := indicesNode.shape.AxisName(i)
				specs[i] = compute.DynamicDimensionSpec{Name: name, Value: dimSize}
			} else {
				specs[i] = compute.DynamicDimensionSpec{Static: B[i]}
			}
		}
		for i := 0; i < len(O); i++ {
			dimIdx := len(B) + i
			if O[i] == shapes.DynamicDim {
				dimSize, err := f.DynamicDimensionSize(gatherNDNode, K+i)
				if err != nil {
					return nil, errors.Wrap(err, "Gather: failed to get dynamic dimension size for slice axis")
				}
				name := gatherNDNode.shape.AxisName(K + i)
				specs[dimIdx] = compute.DynamicDimensionSpec{Name: name, Value: dimSize}
			} else {
				specs[dimIdx] = compute.DynamicDimensionSpec{Static: O[i]}
			}
		}
		var err error
		reshapedVal, err = f.DynamicReshape(gatherNDNode, specs...)
		if err != nil {
			return nil, errors.Wrap(err, "Gather: failed to dynamic reshape GatherND output")
		}
	} else {
		var err error
		reshapedVal, err = f.Reshape(gatherNDNode, reshapedDims...)
		if err != nil {
			return nil, errors.Wrap(err, "Gather: failed to reshape GatherND output to B+O shape")
		}
	}
	reshapedNode := reshapedVal.(*Node)

	// 6. Transpose to the final output shape.
	finalRank := len(reshapedDims)
	P := make([]int, finalRank)
	for i := range finalRank {
		offsetIdx := -1
		for idx, axis := range offsetOutputAxes {
			if axis == i {
				offsetIdx = idx
				break
			}
		}
		if offsetIdx != -1 {
			P[i] = len(B) + offsetIdx
		} else {
			// Find the index of this batch dimension in B.
			j := 0
			for prev := 0; prev < i; prev++ {
				isOffset := slices.Contains(offsetOutputAxes, prev)
				if !isOffset {
					j++
				}
			}
			P[i] = j
		}
	}

	finalVal, err := f.Transpose(reshapedNode, P...)
	if err != nil {
		return nil, errors.Wrap(err, "Gather: failed to transpose reshaped output to final output axes placement")
	}
	return finalVal, nil
}
