// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"

	"github.com/gomlx/compute"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/pkg/errors"
)

func isIdentityPermutation(perm []int) bool {
	for i, p := range perm {
		if p != i {
			return false
		}
	}
	return true
}

func (f *Function) scatter(
	reductionType string,
	operand, scatterIndices, updates compute.Value,
	indexVectorAxis int,
	updateWindowAxes, insertedWindowAxes, scatterAxesToOperandAxes []int,
	indicesAreSorted, uniqueIndices bool,
) (compute.Value, error) {
	operandNode, ok1 := operand.(*Node)
	indicesNode, ok2 := scatterIndices.(*Node)
	updatesNode, ok3 := updates.(*Node)
	if !ok1 || !ok2 || !ok3 {
		return nil, errors.New("scatter: inputs must be valid onnxruntime nodes")
	}

	// 1. Transpose operand so that indexed axes are first, and slice axes are last.
	operandRank := operandNode.shape.Rank()
	indexedAxes := scatterAxesToOperandAxes
	indexedAxesSet := make(map[int]bool)
	for _, ax := range indexedAxes {
		indexedAxesSet[ax] = true
	}
	var sliceAxes []int
	for i := range operandRank {
		if !indexedAxesSet[i] {
			sliceAxes = append(sliceAxes, i)
		}
	}
	operandPerm := append(indexedAxes, sliceAxes...)
	var transposedOperand compute.Value = operandNode
	if !isIdentityPermutation(operandPerm) {
		var err error
		transposedOperand, err = f.Transpose(operandNode, operandPerm...)
		if err != nil {
			return nil, err
		}
	}

	// 2. Transpose updates so that index/batch dimensions are first, and window dimensions are last.
	updatesRank := updatesNode.shape.Rank()
	updateWindowSet := make(map[int]bool)
	for _, ax := range updateWindowAxes {
		updateWindowSet[ax] = true
	}
	var updatesIndexAxes []int
	for i := range updatesRank {
		if !updateWindowSet[i] {
			updatesIndexAxes = append(updatesIndexAxes, i)
		}
	}
	updatesPerm := append(updatesIndexAxes, updateWindowAxes...)
	var transposedUpdates compute.Value = updatesNode
	if !isIdentityPermutation(updatesPerm) {
		var err error
		transposedUpdates, err = f.Transpose(updatesNode, updatesPerm...)
		if err != nil {
			return nil, err
		}
	}

	// Cast indices to int64 for standard ONNX ScatterND
	indicesInput := indicesNode
	if indicesNode.shape.DType != dtypes.Int64 {
		casted, err := f.ConvertDType(indicesNode, dtypes.Int64)
		if err != nil {
			return nil, err
		}
		indicesInput = casted.(*Node)
	}

	// 3. Create ScatterND node on the transposed inputs
	f.nodeCount++
	scatterNode := &Node{
		name:   fmt.Sprintf("node_%d", f.nodeCount),
		opType: "ScatterND",
		inputs: []*Node{transposedOperand.(*Node), indicesInput, transposedUpdates.(*Node)},
		shape:  transposedOperand.(*Node).shape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "reduction",
				Type: onnx.AttributeProto_STRING,
				S:    []byte(reductionType),
			},
		},
	}
	f.nodes = append(f.nodes, scatterNode)

	// 4. Transpose the result back to the original shape
	var finalResult compute.Value = scatterNode
	if !isIdentityPermutation(operandPerm) {
		operandInvPerm := make([]int, operandRank)
		for i, p := range operandPerm {
			operandInvPerm[p] = i
		}
		var err error
		finalResult, err = f.Transpose(scatterNode, operandInvPerm...)
		if err != nil {
			return nil, err
		}
	}

	return finalResult, nil
}

func (f *Function) ScatterMax(
	operand, scatterIndices, updates compute.Value,
	indexVectorAxis int,
	updateWindowAxes, insertedWindowAxes, scatterAxesToOperandAxes []int,
	indicesAreSorted, uniqueIndices bool,
) (compute.Value, error) {
	return f.scatter("max", operand, scatterIndices, updates, indexVectorAxis,
		updateWindowAxes, insertedWindowAxes, scatterAxesToOperandAxes, indicesAreSorted, uniqueIndices)
}

func (f *Function) ScatterMin(
	operand, scatterIndices, updates compute.Value,
	indexVectorAxis int,
	updateWindowAxes, insertedWindowAxes, scatterAxesToOperandAxes []int,
	indicesAreSorted, uniqueIndices bool,
) (compute.Value, error) {
	return f.scatter("min", operand, scatterIndices, updates, indexVectorAxis,
		updateWindowAxes, insertedWindowAxes, scatterAxesToOperandAxes, indicesAreSorted, uniqueIndices)
}

func (f *Function) ScatterSum(
	operand, scatterIndices, updates compute.Value,
	indexVectorAxis int,
	updateWindowAxes, insertedWindowAxes, scatterAxesToOperandAxes []int,
	indicesAreSorted, uniqueIndices bool,
) (compute.Value, error) {
	return f.scatter("add", operand, scatterIndices, updates, indexVectorAxis,
		updateWindowAxes, insertedWindowAxes, scatterAxesToOperandAxes, indicesAreSorted, uniqueIndices)
}
