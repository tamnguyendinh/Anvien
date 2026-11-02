// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"github.com/gomlx/compute"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

// RNGBitGenerator generates a tensor of the requested shape filled with random bits.
// It receives a state of shape [3]uint64 containing [seed/key, counter_low, counter_high].
// It implements SplitMix64 PRNG algorithm inside the ONNX graph.
func (f *Function) RNGBitGenerator(state compute.Value, shape shapes.Shape) (newState compute.Value, values compute.Value, err error) {
	stateNode, ok := state.(*Node)
	if !ok {
		return nil, nil, errors.New("RNGBitGenerator: state must be a valid onnxruntime node")
	}

	sliceIndex := func(idx int) (*Node, error) {
		val, err := f.Slice(stateNode, []int{idx}, []int{idx + 1}, []int{1})
		if err != nil {
			return nil, err
		}
		reshaped, err := f.Reshape(val)
		if err != nil {
			return nil, err
		}
		return reshaped.(*Node), nil
	}

	key, err := sliceIndex(0)
	if err != nil {
		return nil, nil, err
	}
	counterLow, err := sliceIndex(1)
	if err != nil {
		return nil, nil, err
	}
	counterHigh, err := sliceIndex(2)
	if err != nil {
		return nil, nil, err
	}

	size := shape.Size()
	if size == 0 {
		// Return original state and empty tensor
		return state, stateNode, nil
	}

	// 1. Generate 1D sequence of counters: iotaVal = [0, 1, 2, ..., size-1]
	iotaVal, err := f.Iota(shapes.Make(dtypes.Uint64, size), 0)
	if err != nil {
		return nil, nil, err
	}

	// counters = iotaVal + counterLow
	counters, err := f.Add(iotaVal, counterLow)
	if err != nil {
		return nil, nil, err
	}

	// 2. Perform SplitMix64 mixing step
	// x = counters ^ key
	x, err := f.BitwiseXor(counters, key)
	if err != nil {
		return nil, nil, err
	}

	// c1 = 0xbf58476d1ce4e5b9
	c1Const, err := f.Constant([]uint64{0xbf58476d1ce4e5b9})
	if err != nil {
		return nil, nil, err
	}
	// c2 = 0x94d049bb133111eb
	c2Const, err := f.Constant([]uint64{0x94d049bb133111eb})
	if err != nil {
		return nil, nil, err
	}

	step := func(val compute.Value, shift int, c compute.Value) (compute.Value, error) {
		shiftConst, err := f.Constant([]uint64{uint64(shift)})
		if err != nil {
			return nil, err
		}
		shifted, err := f.ShiftRightLogical(val, shiftConst)
		if err != nil {
			return nil, err
		}
		xored, err := f.BitwiseXor(val, shifted)
		if err != nil {
			return nil, err
		}
		return f.Mul(xored, c)
	}

	// x = (x ^ (x >> 30)) * c1
	x, err = step(x, 30, c1Const)
	if err != nil {
		return nil, nil, err
	}
	// x = (x ^ (x >> 27)) * c2
	x, err = step(x, 27, c2Const)
	if err != nil {
		return nil, nil, err
	}
	// values = x ^ (x >> 31)
	shift31Const, err := f.Constant([]uint64{31})
	if err != nil {
		return nil, nil, err
	}
	shifted31, err := f.ShiftRightLogical(x, shift31Const)
	if err != nil {
		return nil, nil, err
	}
	flatValues, err := f.BitwiseXor(x, shifted31)
	if err != nil {
		return nil, nil, err
	}

	// 3. Cast values to the requested output dtype if different from Uint64
	var castedValues compute.Value = flatValues
	if shape.DType != dtypes.Uint64 {
		castedValues, err = f.ConvertDType(flatValues, shape.DType)
		if err != nil {
			return nil, nil, err
		}
	}

	// 4. Reshape flatValues to the final shape
	dims := shape.Dimensions
	if len(dims) == 0 {
		dims = []int{1}
	}
	reshapedValues, err := f.Reshape(castedValues, dims...)
	if err != nil {
		return nil, nil, err
	}
	values = reshapedValues
	if shape.Rank() == 0 {
		// Reshape [1] back to [] (scalar)
		values, err = f.Reshape(reshapedValues)
		if err != nil {
			return nil, nil, err
		}
	}

	// 5. Update RNG state: counter_low = counter_low + size
	sizeConst, err := f.Constant([]uint64{uint64(size)})
	if err != nil {
		return nil, nil, err
	}
	newCounterLow, err := f.Add(counterLow, sizeConst)
	if err != nil {
		return nil, nil, err
	}

	reshapeTo1D := func(n *Node) (compute.Value, error) {
		return f.Reshape(n, 1)
	}
	key1D, err := reshapeTo1D(key)
	if err != nil {
		return nil, nil, err
	}
	newCounterLow1D, err := reshapeTo1D(newCounterLow.(*Node))
	if err != nil {
		return nil, nil, err
	}
	counterHigh1D, err := reshapeTo1D(counterHigh)
	if err != nil {
		return nil, nil, err
	}

	newState, err = f.Concatenate(0, key1D, newCounterLow1D, counterHigh1D)
	if err != nil {
		return nil, nil, err
	}

	return newState, values, nil
}
