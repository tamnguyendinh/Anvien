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
	batchAxesMatch := true
	if len(lhsBatchAxes) != len(rhsBatchAxes) {
		batchAxesMatch = false
	} else {
		for i, ba := range lhsBatchAxes {
			if rhsBatchAxes[i] != ba {
				batchAxesMatch = false
				break
			}
		}
	}

	// ONNX MatMul operator only supports floating-point and 32/64-bit integer types.
	matMulDTypeSupported := accumulationDType.IsFloat() ||
		accumulationDType == dtypes.Int32 || accumulationDType == dtypes.Int64 ||
		accumulationDType == dtypes.Uint32 || accumulationDType == dtypes.Uint64

	canUseMatMul := matMulDTypeSupported && batchAxesMatch && len(lhsContractingAxes) == 1 && len(rhsContractingAxes) == 1 && lhsRank <= 2 && rhsRank <= 2

	if canUseMatMul {
		matLhs := lhsInput
		matRhs := rhsInput
		squeezeLhs := false
		squeezeRhs := false

		// If LHS is 2D and contracting axis is 0, transpose LHS to [1, 0]
		if lhsRank == 2 && len(lhsBatchAxes) == 0 && lhsContractingAxes[0] == 0 {
			transLhs, err := f.Transpose(matLhs, 1, 0)
			if err != nil {
				return nil, err
			}
			matLhs = transLhs.(*Node)
		}
		// If RHS is 2D and contracting axis is 1, transpose RHS to [1, 0]
		if rhsRank == 2 && len(rhsBatchAxes) == 0 && rhsContractingAxes[0] == 1 {
			transRhs, err := f.Transpose(matRhs, 1, 0)
			if err != nil {
				return nil, err
			}
			matRhs = transRhs.(*Node)
		}

		// Recheck contracting axes positions for general case
		curLhsContract := lhsContractingAxes[0]
		if lhsRank == 2 && len(lhsBatchAxes) == 0 && lhsContractingAxes[0] == 0 {
			curLhsContract = 1
		}
		curRhsContract := rhsContractingAxes[0]
		if rhsRank == 2 && len(rhsBatchAxes) == 0 && rhsContractingAxes[0] == 1 {
			curRhsContract = 0
		}

		if curLhsContract != matLhs.shape.Rank()-1 || (matRhs.shape.Rank() > 1 && curRhsContract != matRhs.shape.Rank()-2) {
			// Cannot simple transpose to MatMul, fallback to Einsum
			canUseMatMul = false
		} else {
			// ONNX Runtime WebGPU strictly requires 2D+ tensors for MatMul kernels.
			// If LHS is 1D [K], unsqueeze to [1, K]
			if lhsRank == 1 {
				matLhsVal, err := f.Reshape(lhsInput, 1, lhsInput.shape.Dimensions[0])
				if err != nil {
					return nil, err
				}
				matLhs = matLhsVal.(*Node)
				squeezeLhs = true
			}
			// If RHS is 1D [K], unsqueeze to [K, 1]
			if rhsRank == 1 {
				matRhsVal, err := f.Reshape(rhsInput, rhsInput.shape.Dimensions[0], 1)
				if err != nil {
					return nil, err
				}
				matRhs = matRhsVal.(*Node)
				squeezeRhs = true
			}

			// Calculate 2D matmul shape
			matOutShape := outShape
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
	if outShape.Rank() == 0 {
		newDimsConst, err := f.Constant([]int64{}, 0)
		if err != nil {
			return nil, err
		}

		f.nodeCount++
		reshapeNode := &Node{
			name:   fmt.Sprintf("node_%d", f.nodeCount),
			opType: "Reshape",
			inputs: []*Node{lastNode, newDimsConst.(*Node)},
			shape:  outShape,
		}
		f.nodes = append(f.nodes, reshapeNode)
		lastNode = reshapeNode
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
