// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"github.com/gomlx/compute"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/dtypes/bfloat16"
	"github.com/gomlx/compute/dtypes/float16"
	"github.com/pkg/errors"
)

func init() {
	registerOp(compute.OpTypeSchedulingBarrier)
	registerOp(compute.OpTypeOptimizationBarrier)
}

// OptimizationBarrier implements an optimization barrier.
// In ONNX Runtime, since we do not perform graph rewriting or CSE across barriers, this acts as Identity.
func (f *Function) OptimizationBarrier(operands ...compute.Value) ([]compute.Value, error) {
	if len(operands) == 0 {
		return nil, errors.New("OptimizationBarrier requires at least one operand")
	}
	outputs := make([]compute.Value, len(operands))
	for i, op := range operands {
		id, err := f.Identity(op)
		if err != nil {
			return nil, err
		}
		outputs[i] = id
	}
	return outputs, nil
}

// SchedulingBarrier introduces a scheduling barrier.
// Returned value is identity to the operand, but it is guaranteed to depend on all the dependencies.
//
// Since ONNX Runtime doesn't natively support a scheduling barrier op, we "hack" it by means of a no-op:
// taking a scalar of each dependency, multiplying by 0, and adding them to the operand.
func (f *Function) SchedulingBarrier(operand compute.Value, dependencies ...compute.Value) (compute.Value, error) {
	if len(dependencies) == 0 {
		return operand, nil
	}

	opNode, ok := operand.(*Node)
	if !ok {
		return nil, errors.New("SchedulingBarrier: operand must be a valid onnxruntime node")
	}

	opDType := opNode.shape.DType

	// Create a constant zero of the operand's dtype.
	var zeroFlat any
	switch opDType {
	case dtypes.Float32:
		zeroFlat = []float32{0}
	case dtypes.Float64:
		zeroFlat = []float64{0}
	case dtypes.Int32:
		zeroFlat = []int32{0}
	case dtypes.Int64:
		zeroFlat = []int64{0}
	case dtypes.Int16:
		zeroFlat = []int16{0}
	case dtypes.Int8:
		zeroFlat = []int8{0}
	case dtypes.Uint8:
		zeroFlat = []uint8{0}
	case dtypes.Uint16:
		zeroFlat = []uint16{0}
	case dtypes.Uint32:
		zeroFlat = []uint32{0}
	case dtypes.Uint64:
		zeroFlat = []uint64{0}
	case dtypes.Float16:
		zeroFlat = []float16.Float16{float16.FromFloat32(0)}
	case dtypes.BFloat16:
		zeroFlat = []bfloat16.BFloat16{bfloat16.FromFloat32(0)}
	case dtypes.Bool:
		zeroFlat = []bool{false}
	default:
		return nil, errors.Errorf("SchedulingBarrier: unsupported operand DType %s", opDType)
	}

	zeroConst, err := f.Constant(zeroFlat)
	if err != nil {
		return nil, errors.WithMessage(err, "SchedulingBarrier: failed to create constant zero")
	}

	accumulatedZero := zeroConst
	for _, dep := range dependencies {
		depNode, ok := dep.(*Node)
		if !ok {
			return nil, errors.New("SchedulingBarrier: dependency must be a valid onnxruntime node")
		}

		var scalarDep compute.Value
		if depNode.shape.Rank() == 0 {
			scalarDep = depNode
		} else {
			starts := make([]int, depNode.shape.Rank())
			limits := make([]int, depNode.shape.Rank())
			strides := make([]int, depNode.shape.Rank())
			for i := 0; i < depNode.shape.Rank(); i++ {
				starts[i] = 0
				limits[i] = 1
				strides[i] = 1
			}
			slice, err := f.Slice(depNode, starts, limits, strides)
			if err != nil {
				return nil, err
			}
			scalarDep, err = f.Reshape(slice)
			if err != nil {
				return nil, err
			}
		}

		if scalarDep.(*Node).shape.DType != opDType {
			scalarDep, err = f.ConvertDType(scalarDep, opDType)
			if err != nil {
				return nil, err
			}
		}

		term, err := f.Mul(scalarDep, zeroConst)
		if err != nil {
			return nil, err
		}

		accumulatedZero, err = f.Add(accumulatedZero, term)
		if err != nil {
			return nil, err
		}
	}

	operandReady, err := f.Add(operand, accumulatedZero)
	if err != nil {
		return nil, err
	}

	return operandReady, nil
}
