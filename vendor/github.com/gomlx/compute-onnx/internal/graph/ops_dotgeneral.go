// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"
	"strings"

	"github.com/gomlx/compute"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapeinference"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

// DotGeneral performs a general dot product (matrix multiplication) of lhs and rhs.
// It maps the generalized contracting and batching axes of XLA semantics to ONNX Einsum.
func (f *Function) DotGeneral(
	lhs compute.Value,
	lhsContractingAxes []int,
	lhsBatchAxes []int,
	rhs compute.Value,
	rhsContractingAxes []int,
	rhsBatchAxes []int,
	config compute.DotGeneralConfig,
) (compute.Value, error) {
	lhsNode, ok1 := lhs.(*Node)
	rhsNode, ok2 := rhs.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("DotGeneral: inputs must be valid onnxruntime nodes")
	}

	// Determine accumulation and output types
	accumulationDType := lhsNode.shape.DType
	if config.AccumulatorDType != dtypes.InvalidDType {
		accumulationDType = config.AccumulatorDType
	}

	expectedOutputDType := lhsNode.shape.DType
	if config.OutputDType != dtypes.InvalidDType {
		expectedOutputDType = config.OutputDType
	}

	// Cast inputs to accumulator type if necessary
	lhsInput := lhsNode
	rhsInput := rhsNode

	if lhsNode.shape.DType != accumulationDType {
		casted, err := f.ConvertDType(lhsNode, accumulationDType)
		if err != nil {
			return nil, err
		}
		lhsInput = casted.(*Node)
	}

	if rhsNode.shape.DType != accumulationDType {
		casted, err := f.ConvertDType(rhsNode, accumulationDType)
		if err != nil {
			return nil, err
		}
		rhsInput = casted.(*Node)
	}

	// 1. Infer output shape
	outShape, err := shapeinference.DotGeneral(
		lhsInput.shape, lhsContractingAxes, lhsBatchAxes,
		rhsInput.shape, rhsContractingAxes, rhsBatchAxes,
		config,
	)
	if err != nil {
		return nil, err
	}

	// 2. Generate Einsum equation
	charPool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charIdx := 0
	getNewChar := func() string {
		if charIdx >= len(charPool) {
			panic("DotGeneral: ran out of character pool for Einsum (rank is too high)")
		}
		c := string(charPool[charIdx])
		charIdx++
		return c
	}

	lhsRank := lhsInput.shape.Rank()
	rhsRank := rhsInput.shape.Rank()

	lhsChars := make([]string, lhsRank)
	for i := range lhsRank {
		lhsChars[i] = getNewChar()
	}

	rhsBatchMap := make(map[int]int)
	for idx, rhsAxis := range rhsBatchAxes {
		rhsBatchMap[rhsAxis] = lhsBatchAxes[idx]
	}

	rhsContractMap := make(map[int]int)
	for idx, rhsAxis := range rhsContractingAxes {
		rhsContractMap[rhsAxis] = lhsContractingAxes[idx]
	}

	rhsChars := make([]string, rhsRank)
	for i := range rhsRank {
		if lhsAxis, ok := rhsBatchMap[i]; ok {
			rhsChars[i] = lhsChars[lhsAxis]
		} else if lhsAxis, ok := rhsContractMap[i]; ok {
			rhsChars[i] = lhsChars[lhsAxis]
		} else {
			rhsChars[i] = getNewChar()
		}
	}

	outputChars := make([]string, 0)
	for _, lhsAxis := range lhsBatchAxes {
		outputChars = append(outputChars, lhsChars[lhsAxis])
	}

	lhsContractSet := make(map[int]bool)
	for _, axis := range lhsContractingAxes {
		lhsContractSet[axis] = true
	}
	lhsBatchSet := make(map[int]bool)
	for _, axis := range lhsBatchAxes {
		lhsBatchSet[axis] = true
	}
	for i := range lhsRank {
		if !lhsContractSet[i] && !lhsBatchSet[i] {
			outputChars = append(outputChars, lhsChars[i])
		}
	}

	rhsContractSet := make(map[int]bool)
	for _, axis := range rhsContractingAxes {
		rhsContractSet[axis] = true
	}
	rhsBatchSet := make(map[int]bool)
	for _, axis := range rhsBatchAxes {
		rhsBatchSet[axis] = true
	}
	for i := range rhsRank {
		if !rhsContractSet[i] && !rhsBatchSet[i] {
			outputChars = append(outputChars, rhsChars[i])
		}
	}

	// 2. Optimization: Use standard MatMul for matrix multiplication when contracting 1 axis.
	var lastNode *Node

	// ONNX MatMul operator only supports floating-point and 32/64-bit integer types.
	matMulDTypeSupported := accumulationDType.IsFloat() ||
		accumulationDType == dtypes.Int32 || accumulationDType == dtypes.Int64 ||
		accumulationDType == dtypes.Uint32 || accumulationDType == dtypes.Uint64

	canUseMatMul := false
	matLhs := lhsInput
	matRhs := rhsInput
	squeezeLhs := false
	squeezeRhs := false

	if matMulDTypeSupported && len(lhsContractingAxes) == 1 && len(rhsContractingAxes) == 1 {
		lhsContract := lhsContractingAxes[0]
		rhsContract := rhsContractingAxes[0]

		// Case A: 1D/2D Matrix Multiplication (ranks <= 2)
		if lhsRank <= 2 && rhsRank <= 2 && len(lhsBatchAxes) == len(rhsBatchAxes) {
			batchAxesMatch := true
			for i, ba := range lhsBatchAxes {
				if rhsBatchAxes[i] != ba {
					batchAxesMatch = false
					break
				}
			}
			if batchAxesMatch {
				// If LHS is 2D and contracting axis is 0, transpose LHS to [1, 0]
				if lhsRank == 2 && len(lhsBatchAxes) == 0 && lhsContract == 0 {
					transLhs, err := f.Transpose(matLhs, 1, 0)
					if err != nil {
						return nil, err
					}
					matLhs = transLhs.(*Node)
					lhsContract = 1
				}
				// If RHS is 2D and contracting axis is 1, transpose RHS to [1, 0]
				if rhsRank == 2 && len(rhsBatchAxes) == 0 && rhsContract == 1 {
					transRhs, err := f.Transpose(matRhs, 1, 0)
					if err != nil {
						return nil, err
					}
					matRhs = transRhs.(*Node)
					rhsContract = 0
				}

				if lhsContract == matLhs.shape.Rank()-1 && (matRhs.shape.Rank() <= 1 || rhsContract == matRhs.shape.Rank()-2) {
					canUseMatMul = true
					// ONNX Runtime WebGPU strictly requires 2D+ tensors for MatMul kernels.
					if lhsRank == 1 {
						matLhsVal, err := f.Reshape(matLhs, 1, matLhs.shape.Dimensions[0])
						if err != nil {
							return nil, err
						}
						matLhs = matLhsVal.(*Node)
						squeezeLhs = true
					}
					if rhsRank == 1 {
						matRhsVal, err := f.Reshape(matRhs, matRhs.shape.Dimensions[0], 1)
						if err != nil {
							return nil, err
						}
						matRhs = matRhsVal.(*Node)
						squeezeRhs = true
					}
				}
			}
		}

		// Case B: Batched Matrix Multiplication (e.g. Attention Q @ K^T and Attn @ V)
		// Both LHS and RHS have rank >= 3 with identical leading contiguous batch axes 0..B-1,
		// and each operand has exactly 2 non-batch trailing axes.
		if !canUseMatMul && len(lhsBatchAxes) > 0 && len(lhsBatchAxes) == len(rhsBatchAxes) {
			numBatch := len(lhsBatchAxes)
			batchAxesMatch := true
			for i := range numBatch {
				if lhsBatchAxes[i] != i || rhsBatchAxes[i] != i {
					batchAxesMatch = false
					break
				}
			}
			if batchAxesMatch && lhsRank == numBatch+2 && rhsRank == numBatch+2 {
				if (lhsContract == numBatch || lhsContract == numBatch+1) &&
					(rhsContract == numBatch || rhsContract == numBatch+1) {
					// For LHS, contracting axis should be the last axis (numBatch+1)
					if lhsContract == numBatch {
						perm := make([]int, lhsRank)
						for i := range numBatch {
							perm[i] = i
						}
						perm[numBatch] = numBatch + 1
						perm[numBatch+1] = numBatch
						transLhs, err := f.Transpose(matLhs, perm...)
						if err != nil {
							return nil, err
						}
						matLhs = transLhs.(*Node)
					}
					// For RHS, contracting axis should be second-to-last (numBatch)
					if rhsContract == numBatch+1 {
						perm := make([]int, rhsRank)
						for i := range numBatch {
							perm[i] = i
						}
						perm[numBatch] = numBatch + 1
						perm[numBatch+1] = numBatch
						transRhs, err := f.Transpose(matRhs, perm...)
						if err != nil {
							return nil, err
						}
						matRhs = transRhs.(*Node)
					}
					canUseMatMul = true
				}
			}
		}

		// Case C: N-D LHS x 2D RHS (Dense projection, e.g. [..., inFeat] @ [inFeat, outFeat])
		// LHS has rank >= 3, no batch axes, and contracts on the last axis.
		// RHS has rank 2, no batch axes, and contracts on axis 0 (or axis 1 after transpose).
		if !canUseMatMul && len(lhsBatchAxes) == 0 && len(rhsBatchAxes) == 0 && lhsRank >= 3 && rhsRank == 2 {
			if lhsContract == lhsRank-1 && (rhsContract == 0 || rhsContract == 1) {
				if rhsContract == 1 {
					transRhs, err := f.Transpose(matRhs, 1, 0)
					if err != nil {
						return nil, err
					}
					matRhs = transRhs.(*Node)
				}
				canUseMatMul = true
			}
		}

		if canUseMatMul {
			// Calculate matmul output shape (ONNX MatMul always produces accumulationDType).
			matOutShape := outShape.Clone()
			matOutShape.DType = accumulationDType
			if squeezeLhs || squeezeRhs {
				matDims := make([]int, 0, len(outShape.Dimensions)+2)
				for i := 0; i < lhsRank-1; i++ {
					matDims = append(matDims, lhsInput.shape.Dimensions[i])
				}
				if squeezeLhs {
					matDims = append(matDims, 1)
				}
				for i := 0; i < rhsRank; i++ {
					if i != rhsContractingAxes[0] {
						matDims = append(matDims, rhsInput.shape.Dimensions[i])
					}
				}
				if squeezeRhs {
					matDims = append(matDims, 1)
				}
				matOutShape = shapes.Make(accumulationDType, matDims...)
			}

			f.nodeCount++
			matmulNode := &Node{
				name:   fmt.Sprintf("node_%d", f.nodeCount),
				opType: "MatMul",
				inputs: []*Node{matLhs, matRhs},
				shape:  matOutShape,
			}
			f.nodes = append(f.nodes, matmulNode)
			lastNode = matmulNode

			// If we unsqueezed LHS or RHS, reshape back to outShape
			if squeezeLhs || squeezeRhs {
				if !matOutShape.Equal(outShape) {
					reshaped, err := f.Reshape(matmulNode, outShape.Dimensions...)
					if err != nil {
						return nil, err
					}
					lastNode = reshaped.(*Node)
				}
			}
		}
	}

	if !canUseMatMul {
		// 3. Fallback: Generate Einsum equation
		if accumulationDType == dtypes.BFloat16 {
			return nil, errors.Wrapf(compute.ErrNotImplemented, "ONNX doesn't support BFloat16 for Einsum: standard ONNX Einsum schema (opset 21) does not support tensor(bfloat16)")
		}

		equation := fmt.Sprintf("%s,%s->%s",
			strings.Join(lhsChars, ""),
			strings.Join(rhsChars, ""),
			strings.Join(outputChars, ""),
		)

		einsumShape := outShape
		einsumShape.DType = accumulationDType

		f.nodeCount++
		einsumNode := &Node{
			name:   fmt.Sprintf("node_%d", f.nodeCount),
			opType: "Einsum",
			inputs: []*Node{lhsInput, rhsInput},
			shape:  einsumShape,
			attributes: []*onnx.AttributeProto{
				{
					Name: "equation",
					Type: onnx.AttributeProto_STRING,
					S:    []byte(equation),
				},
			},
		}
		f.nodes = append(f.nodes, einsumNode)
		lastNode = einsumNode
	}
	if lastNode.shape.Rank() != outShape.Rank() {
		reshaped, err := f.Reshape(lastNode, outShape.Dimensions...)
		if err != nil {
			return nil, err
		}
		lastNode = reshaped.(*Node)
	}

	// Cast to final output type if necessary
	var finalVal compute.Value = lastNode
	if accumulationDType != expectedOutputDType {
		casted, err := f.ConvertDType(lastNode, expectedOutputDType)
		if err != nil {
			return nil, err
		}
		finalVal = casted
	}

	return finalVal, nil
}
