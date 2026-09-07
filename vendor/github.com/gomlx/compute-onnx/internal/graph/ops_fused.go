// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"math"

	"github.com/gomlx/compute"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/dtypes/bfloat16"
	"github.com/gomlx/compute/dtypes/float16"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

func init() {
	registerOp(compute.OpTypeFusedSoftmax)
	registerOp(compute.OpTypeFusedGelu)
	registerOp(compute.OpTypeFusedLayerNorm)
	registerOp(compute.OpTypeFusedDense)
	registerOp(compute.OpTypeFusedScaledDotProductAttention)
	registerOp(compute.OpTypeFusedAttentionQKVProjection)
	registerOp(compute.OpTypeFusedQuantizedDense)
}

func broadcastToShape(f *Function, val compute.Value, targetRef *Node) (compute.Value, error) {
	valNode, ok := val.(*Node)
	if !ok {
		return nil, errors.New("broadcastToShape: val must be a valid onnxruntime node")
	}
	targetShape := targetRef.shape
	if valNode.shape.Equal(targetShape) {
		return valNode, nil
	}

	targetRank := targetShape.Rank()
	currentRank := valNode.shape.Rank()

	currVal := valNode
	if currentRank < targetRank {
		diff := targetRank - currentRank
		if valNode.shape.IsDynamic() {
			specs := make([]compute.DynamicDimensionSpec, targetRank)
			for i := range diff {
				specs[i] = compute.DynamicDimensionSpec{Static: 1}
			}
			for i := range currentRank {
				d := valNode.shape.Dimensions[i]
				if d == shapes.DynamicDim {
					dimVal, err := f.DynamicDimensionSize(valNode, i)
					if err != nil {
						return nil, err
					}
					specs[diff+i] = compute.DynamicDimensionSpec{Name: valNode.shape.AxisName(i), Value: dimVal}
				} else {
					specs[diff+i] = compute.DynamicDimensionSpec{Static: d}
				}
			}
			reshaped, err := f.DynamicReshape(valNode, specs...)
			if err != nil {
				return nil, err
			}
			currVal = reshaped.(*Node)
		} else {
			newDims := make([]int, targetRank)
			for i := range diff {
				newDims[i] = 1
			}
			for i := range currentRank {
				newDims[diff+i] = valNode.shape.Dimensions[i]
			}
			reshaped, err := f.Reshape(valNode, newDims...)
			if err != nil {
				return nil, err
			}
			currVal = reshaped.(*Node)
		}
	}

	if currVal.shape.Equal(targetShape) {
		return currVal, nil
	}

	axes := make([]int, targetRank)
	for i := range targetRank {
		axes[i] = i
	}
	if targetShape.IsDynamic() {
		specs := make([]compute.DynamicDimensionSpec, targetRank)
		for i := range targetRank {
			d := targetShape.Dimensions[i]
			if d == shapes.DynamicDim {
				dimVal, err := f.DynamicDimensionSize(targetRef, i)
				if err != nil {
					return nil, err
				}
				specs[i] = compute.DynamicDimensionSpec{Name: targetShape.AxisName(i), Value: dimVal}
			} else {
				specs[i] = compute.DynamicDimensionSpec{Static: d}
			}
		}
		return f.DynamicBroadcastInDim(currVal, axes, specs...)
	}
	return f.BroadcastInDim(currVal, targetShape, axes)
}

func alignAttentionScoreTensor(f *Function, val compute.Value, axesLayout compute.AttentionAxesLayout, qBHSD *Node, targetRef *Node) (compute.Value, error) {
	node, ok := val.(*Node)
	if !ok {
		return nil, errors.New("alignAttentionScoreTensor: val must be a valid onnxruntime node")
	}

	rank := node.shape.Rank()
	targetRank := targetRef.shape.Rank()
	if targetRank != 4 {
		return nil, errors.Errorf("alignAttentionScoreTensor: expected 4D target score shape, got rank %d (%s)", targetRank, targetRef.shape)
	}

	var currNode *Node = node
	if rank == 2 {
		// [Sq, Skv] -> [1, 1, Sq, Skv]
		reshaped, err := f.Reshape(node, 1, 1, node.shape.Dimensions[0], node.shape.Dimensions[1])
		if err != nil {
			return nil, err
		}
		currNode = reshaped.(*Node)
	} else if rank == 3 {
		// [B, Sq, Skv] -> [B, 1, Sq, Skv]
		reshaped, err := f.Reshape(node, node.shape.Dimensions[0], 1, node.shape.Dimensions[1], node.shape.Dimensions[2])
		if err != nil {
			return nil, err
		}
		currNode = reshaped.(*Node)
	} else if rank == 4 && axesLayout == compute.AttentionAxesLayoutBSHD {
		// In BSHD mode, 4D tensor can be [B, Sq, H, Skv] (BSHD) or [B, H, Sq, Skv] (BHSD).
		numHeadsQ := qBHSD.shape.Dimensions[1]
		seqLenQ := qBHSD.shape.Dimensions[2]
		d1 := node.shape.Dimensions[1]
		d2 := node.shape.Dimensions[2]

		// Transpose if it's in BSHD layout [B, Sq, H, Skv] (i.e. d1 matches Sq or d2 matches H)
		if (d1 == seqLenQ || d1 == 1) && (d2 == numHeadsQ || d2 == 1) && !(d1 == numHeadsQ && d2 == seqLenQ) {
			transposed, err := f.Transpose(node, 0, 2, 1, 3)
			if err != nil {
				return nil, err
			}
			currNode = transposed.(*Node)
		}
	}

	return broadcastToShape(f, currNode, targetRef)
}

// FusedSoftmax computes softmax along the specified axis using ONNX Softmax.
func (f *Function) FusedSoftmax(x compute.Value, axis int) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("FusedSoftmax: input must be a valid onnxruntime node")
	}
	rank := xNode.shape.Rank()
	if axis < 0 || axis >= rank {
		return nil, errors.Errorf("FusedSoftmax: axis %d out of bounds for shape rank %d", axis, rank)
	}

	node := &Node{
		opType: "Softmax",
		inputs: []*Node{xNode},
		shape:  xNode.shape,
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

// FusedGelu computes Gaussian Error Linear Unit activation using ONNX Gelu (opset 20+) or FastGelu (com.microsoft).
func (f *Function) FusedGelu(x compute.Value, exact bool) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("FusedGelu: input must be a valid onnxruntime node")
	}

	if !exact {
		node := &Node{
			domain: "com.microsoft",
			opType: "FastGelu",
			inputs: []*Node{xNode},
			shape:  xNode.shape,
		}
		return f.addNode(node), nil
	}

	node := &Node{
		opType: "Gelu",
		inputs: []*Node{xNode},
		shape:  xNode.shape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "approximate",
				Type: onnx.AttributeProto_STRING,
				S:    []byte("none"),
			},
		},
	}
	return f.addNode(node), nil
}

// FusedLayerNorm applies layer normalization over specified axes using ONNX LayerNormalization (opset 17+).
func (f *Function) FusedLayerNorm(x compute.Value, axes []int, epsilon float64, gamma, beta compute.Value) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("FusedLayerNorm: input must be a valid onnxruntime node")
	}

	rank := xNode.shape.Rank()
	if len(axes) == 0 || len(axes) > rank {
		return nil, errors.Wrapf(compute.ErrNotImplemented, "FusedLayerNorm: invalid axes %v for rank %d", axes, rank)
	}

	// ONNX LayerNormalization requires axes to form a contiguous trailing suffix: [rank-k, ..., rank-1]
	k := len(axes)
	startAxis := rank - k
	for i, a := range axes {
		normalizedAxis := a
		if normalizedAxis < 0 {
			normalizedAxis += rank
		}
		if normalizedAxis != startAxis+i {
			return nil, errors.Wrapf(compute.ErrNotImplemented, "FusedLayerNorm: ONNX LayerNormalization requires contiguous trailing axes [axis..rank-1], got axes=%v for rank %d", axes, rank)
		}
	}

	var gammaNode *Node
	if gamma != nil {
		var ok bool
		gammaNode, ok = gamma.(*Node)
		if !ok {
			return nil, errors.New("FusedLayerNorm: gamma must be a valid onnxruntime node")
		}
	} else {
		// Construct gamma as constant tensor of 1s matching normalized shape
		normDims := xNode.shape.Dimensions[startAxis:]
		size := 1
		for _, d := range normDims {
			size *= d
		}

		var constVal compute.Value
		var err error
		switch xNode.shape.DType {
		case dtypes.Float32:
			ones := make([]float32, size)
			for i := range ones {
				ones[i] = 1.0
			}
			constVal, err = f.Constant(ones, normDims...)
		case dtypes.Float64:
			ones := make([]float64, size)
			for i := range ones {
				ones[i] = 1.0
			}
			constVal, err = f.Constant(ones, normDims...)
		case dtypes.Float16:
			ones := make([]float16.Float16, size)
			for i := range ones {
				ones[i] = float16.FromFloat32(1.0)
			}
			constVal, err = f.Constant(ones, normDims...)
		case dtypes.BFloat16:
			ones := make([]bfloat16.BFloat16, size)
			for i := range ones {
				ones[i] = bfloat16.FromFloat32(1.0)
			}
			constVal, err = f.Constant(ones, normDims...)
		default:
			return nil, errors.Wrapf(compute.ErrNotImplemented, "FusedLayerNorm: unsupported dtype %s for nil gamma", xNode.shape.DType)
		}
		if err != nil {
			return nil, err
		}
		gammaNode = constVal.(*Node)
	}

	inputs := []*Node{xNode, gammaNode}

	if beta != nil {
		betaNode, ok := beta.(*Node)
		if !ok {
			return nil, errors.New("FusedLayerNorm: beta must be a valid onnxruntime node")
		}
		inputs = append(inputs, betaNode)
	}

	node := &Node{
		opType: "LayerNormalization",
		inputs: inputs,
		shape:  xNode.shape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "axis",
				Type: onnx.AttributeProto_INT,
				I:    int64(startAxis),
			},
			{
				Name: "epsilon",
				Type: onnx.AttributeProto_FLOAT,
				F:    float32(epsilon),
			},
		},
	}
	return f.addNode(node), nil
}

// FusedDense performs fused matmul + optional bias + optional activation.
func (f *Function) FusedDense(x, weight, bias compute.Value, options compute.DenseConfig) (compute.Value, error) {
	xNode, ok1 := x.(*Node)
	wNode, ok2 := weight.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("FusedDense: inputs must be valid onnxruntime nodes")
	}

	xRank := xNode.shape.Rank()
	wRank := wNode.shape.Rank()
	if xRank < 1 || wRank < 1 {
		return nil, errors.Errorf("FusedDense: x rank (%d) and weight rank (%d) must be at least 1", xRank, wRank)
	}

	var lhsContractingAxes []int
	var rhsContractingAxes []int

	if options.WeightLayout == compute.DenseLayoutOutputsInput {
		lhsContractingAxes = []int{xRank - 1}
		rhsContractingAxes = []int{wRank - 1}
	} else {
		lhsContractingAxes = []int{xRank - 1}
		rhsContractingAxes = []int{0}
	}

	// Compute MatMul / DotGeneral
	dotRes, err := f.DotGeneral(xNode, lhsContractingAxes, nil, wNode, rhsContractingAxes, nil, compute.DotGeneralConfig{})
	if err != nil {
		return nil, errors.Wrap(err, "FusedDense: matmul failed")
	}

	// Add bias if provided
	outVal := dotRes
	if bias != nil {
		biasNode, ok := bias.(*Node)
		if !ok {
			return nil, errors.New("FusedDense: bias must be a valid onnxruntime node")
		}
		dotNode := dotRes.(*Node)
		biasReshaped, err := broadcastToShape(f, biasNode, dotNode)
		if err != nil {
			return nil, errors.Wrap(err, "FusedDense: bias reshape failed")
		}
		outVal, err = f.Add(dotRes, biasReshaped)
		if err != nil {
			return nil, errors.Wrap(err, "FusedDense: bias addition failed")
		}
	}

	// Apply activation
	switch options.Activation {
	case compute.ActivationNone:
		return outVal, nil
	case compute.ActivationRelu:
		outNode, ok := outVal.(*Node)
		if !ok {
			return nil, errors.New("FusedDense: outVal is not a valid onnxruntime node")
		}
		node := &Node{
			opType: "Relu",
			inputs: []*Node{outNode},
			shape:  outNode.shape,
		}
		return f.addNode(node), nil
	case compute.ActivationGelu:
		return f.FusedGelu(outVal, true)
	case compute.ActivationSilu:
		sig, err := f.Logistic(outVal)
		if err != nil {
			return nil, err
		}
		return f.Mul(outVal, sig)
	case compute.ActivationHardSwish:
		outNode, ok := outVal.(*Node)
		if !ok {
			return nil, errors.New("FusedDense: outVal is not a valid onnxruntime node")
		}
		node := &Node{
			opType: "HardSwish",
			inputs: []*Node{outNode},
			shape:  outNode.shape,
		}
		return f.addNode(node), nil
	case compute.ActivationTanh:
		return f.Tanh(outVal)
	default:
		return nil, errors.Errorf("FusedDense: unsupported activation type %v", options.Activation)
	}
}

// FusedScaledDotProductAttention computes multi-head scaled dot-product attention.
func (f *Function) FusedScaledDotProductAttention(
	query, key, value compute.Value,
	axesLayout compute.AttentionAxesLayout,
	options *compute.ScaledDotProductAttentionConfig) (output compute.Value, statesForVJP []compute.Value, err error) {

	qNode, ok1 := query.(*Node)
	kNode, ok2 := key.(*Node)
	vNode, ok3 := value.(*Node)
	if !ok1 || !ok2 || !ok3 {
		return nil, nil, errors.New("FusedScaledDotProductAttention: inputs must be valid onnxruntime nodes")
	}

	if qNode.shape.Rank() != 4 || kNode.shape.Rank() != 4 || vNode.shape.Rank() != 4 {
		return nil, nil, errors.Errorf("FusedScaledDotProductAttention: query, key, value must be 4D tensors, got ranks %d, %d, %d",
			qNode.shape.Rank(), kNode.shape.Rank(), vNode.shape.Rank())
	}

	var qBHSD, kBHSD, vBHSD *Node
	if axesLayout == compute.AttentionAxesLayoutBSHD {
		qBHSDVal, err := f.Transpose(qNode, 0, 2, 1, 3)
		if err != nil {
			return nil, nil, err
		}
		kBHSDVal, err := f.Transpose(kNode, 0, 2, 1, 3)
		if err != nil {
			return nil, nil, err
		}
		vBHSDVal, err := f.Transpose(vNode, 0, 2, 1, 3)
		if err != nil {
			return nil, nil, err
		}
		qBHSD = qBHSDVal.(*Node)
		kBHSD = kBHSDVal.(*Node)
		vBHSD = vBHSDVal.(*Node)
	} else {
		qBHSD = qNode
		kBHSD = kNode
		vBHSD = vNode
	}

	numHeadsQ := qBHSD.shape.Dimensions[1]
	numHeadsKV := kBHSD.shape.Dimensions[1]
	if numHeadsKV < numHeadsQ {
		if numHeadsQ%numHeadsKV != 0 {
			return nil, nil, errors.Errorf("FusedScaledDotProductAttention: query heads (%d) not divisible by key heads (%d)", numHeadsQ, numHeadsKV)
		}
		repeats := numHeadsQ / numHeadsKV
		sliceK := make([]compute.Value, repeats)
		sliceV := make([]compute.Value, repeats)
		for i := range repeats {
			sliceK[i] = kBHSD
			sliceV[i] = vBHSD
		}
		kBHSDVal, err := f.Concatenate(1, sliceK...)
		if err != nil {
			return nil, nil, err
		}
		vBHSDVal, err := f.Concatenate(1, sliceV...)
		if err != nil {
			return nil, nil, err
		}
		kBHSD = kBHSDVal.(*Node)
		vBHSD = vBHSDVal.(*Node)
	}

	// Q @ K^T -> [B, H, S, Skv]
	scoresVal, err := f.DotGeneral(qBHSD, []int{3}, []int{0, 1}, kBHSD, []int{3}, []int{0, 1}, compute.DotGeneralConfig{})
	if err != nil {
		return nil, nil, errors.Wrap(err, "FusedScaledDotProductAttention: q @ k^T failed")
	}

	headDim := qBHSD.shape.Dimensions[3]
	scale := 1.0 / math.Sqrt(float64(headDim))
	if options != nil && options.Scale != 0 {
		scale = options.Scale
	}

	var scaleConst compute.Value
	switch qNode.shape.DType {
	case dtypes.Float32:
		scaleConst, err = f.Constant([]float32{float32(scale)})
	case dtypes.Float64:
		scaleConst, err = f.Constant([]float64{scale})
	case dtypes.Float16:
		scaleConst, err = f.Constant([]float16.Float16{float16.FromFloat32(float32(scale))})
	case dtypes.BFloat16:
		scaleConst, err = f.Constant([]bfloat16.BFloat16{bfloat16.FromFloat32(float32(scale))})
	default:
		return nil, nil, errors.Errorf("FusedScaledDotProductAttention: unsupported dtype %s", qNode.shape.DType)
	}
	if err != nil {
		return nil, nil, err
	}

	scoresVal, err = f.Mul(scoresVal, scaleConst)
	if err != nil {
		return nil, nil, err
	}

	// Apply QuerySeqLen and KeyValueSeqLen masking if provided
	if options != nil && (options.QuerySeqLen != nil || options.KeyValueSeqLen != nil) {
		seqLen := qBHSD.shape.Dimensions[2]
		kvLen := kBHSD.shape.Dimensions[2]
		batchSize := qBHSD.shape.Dimensions[0]

		var seqLenMask compute.Value

		if options.KeyValueSeqLen != nil {
			kvLenNode, ok := options.KeyValueSeqLen.(*Node)
			if !ok {
				return nil, nil, errors.New("FusedScaledDotProductAttention: KeyValueSeqLen must be a valid onnxruntime node")
			}
			var kvSeqIdx compute.Value
			if kBHSD.shape.IsDynamic() && kvLen == shapes.DynamicDim {
				kvLenDyn, errDyn := f.DynamicDimensionSize(kBHSD, 2)
				if errDyn != nil {
					return nil, nil, errDyn
				}
				kvSeqIdx, err = f.DynamicIota(dtypes.Int32, 3,
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Name: kBHSD.shape.AxisName(2), Value: kvLenDyn},
				)
				if err != nil {
					return nil, nil, err
				}
			} else {
				kvSeqIdx, err = f.Iota(shapes.Make(dtypes.Int32, 1, 1, 1, kvLen), 3)
				if err != nil {
					return nil, nil, err
				}
			}
			kvLen32, err := f.ConvertDType(kvLenNode, dtypes.Int32)
			if err != nil {
				return nil, nil, err
			}
			var kvLenReshaped compute.Value
			if kvLenNode.shape.IsDynamic() && batchSize == shapes.DynamicDim {
				batchDyn, errDyn := f.DynamicDimensionSize(kvLenNode, 0)
				if errDyn != nil {
					return nil, nil, errDyn
				}
				kvLenReshaped, err = f.DynamicReshape(kvLen32,
					compute.DynamicDimensionSpec{Name: kvLenNode.shape.AxisName(0), Value: batchDyn},
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Static: 1},
				)
			} else {
				kvLenReshaped, err = f.Reshape(kvLen32, batchSize, 1, 1, 1)
			}
			if err != nil {
				return nil, nil, err
			}
			validKV, err := f.LessThan(kvSeqIdx, kvLenReshaped)
			if err != nil {
				return nil, nil, err
			}
			seqLenMask = validKV
		}

		if options.QuerySeqLen != nil {
			qLenNode, ok := options.QuerySeqLen.(*Node)
			if !ok {
				return nil, nil, errors.New("FusedScaledDotProductAttention: QuerySeqLen must be a valid onnxruntime node")
			}
			var qSeqIdx compute.Value
			if qBHSD.shape.IsDynamic() && seqLen == shapes.DynamicDim {
				seqLenDyn, errDyn := f.DynamicDimensionSize(qBHSD, 2)
				if errDyn != nil {
					return nil, nil, errDyn
				}
				qSeqIdx, err = f.DynamicIota(dtypes.Int32, 2,
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Name: qBHSD.shape.AxisName(2), Value: seqLenDyn},
					compute.DynamicDimensionSpec{Static: 1},
				)
				if err != nil {
					return nil, nil, err
				}
			} else {
				qSeqIdx, err = f.Iota(shapes.Make(dtypes.Int32, 1, 1, seqLen, 1), 2)
				if err != nil {
					return nil, nil, err
				}
			}
			qLen32, err := f.ConvertDType(qLenNode, dtypes.Int32)
			if err != nil {
				return nil, nil, err
			}
			var qLenReshaped compute.Value
			if qLenNode.shape.IsDynamic() && batchSize == shapes.DynamicDim {
				batchDyn, errDyn := f.DynamicDimensionSize(qLenNode, 0)
				if errDyn != nil {
					return nil, nil, errDyn
				}
				qLenReshaped, err = f.DynamicReshape(qLen32,
					compute.DynamicDimensionSpec{Name: qLenNode.shape.AxisName(0), Value: batchDyn},
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Static: 1},
					compute.DynamicDimensionSpec{Static: 1},
				)
			} else {
				qLenReshaped, err = f.Reshape(qLen32, batchSize, 1, 1, 1)
			}
			if err != nil {
				return nil, nil, err
			}
			validQ, err := f.LessThan(qSeqIdx, qLenReshaped)
			if err != nil {
				return nil, nil, err
			}
			if seqLenMask != nil {
				seqLenMask, err = f.LogicalAnd(seqLenMask, validQ)
				if err != nil {
					return nil, nil, err
				}
			} else {
				seqLenMask = validQ
			}
		}

		if seqLenMask != nil {
			scoresNode := scoresVal.(*Node)
			maskReshaped, err := broadcastToShape(f, seqLenMask, scoresNode)
			if err != nil {
				return nil, nil, errors.Wrap(err, "FusedScaledDotProductAttention: seqLenMask reshape failed")
			}
			maskNode := maskReshaped.(*Node)

			var negInfConst compute.Value
			switch qNode.shape.DType {
			case dtypes.Float32:
				negInfConst, err = f.Constant([]float32{-1e9})
			case dtypes.Float64:
				negInfConst, err = f.Constant([]float64{-1e9})
			case dtypes.Float16:
				negInfConst, err = f.Constant([]float16.Float16{float16.FromFloat32(-10000.0)})
			case dtypes.BFloat16:
				negInfConst, err = f.Constant([]bfloat16.BFloat16{bfloat16.FromFloat32(-1e9)})
			}
			if err != nil {
				return nil, nil, err
			}
			scoresVal, err = f.Where(maskNode, scoresVal, negInfConst)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	// Apply causal mask if requested
	if options != nil && options.Causal {
		seqLen := qBHSD.shape.Dimensions[2]
		kvLen := kBHSD.shape.Dimensions[2]

		var causalMaskConst compute.Value
		switch qNode.shape.DType {
		case dtypes.Float32:
			maskFlat := make([]float32, seqLen*kvLen)
			for i := range seqLen {
				for j := range kvLen {
					if j > i {
						maskFlat[i*kvLen+j] = -1e9
					}
				}
			}
			causalMaskConst, err = f.Constant(maskFlat, 1, 1, seqLen, kvLen)
		case dtypes.Float64:
			maskFlat := make([]float64, seqLen*kvLen)
			for i := range seqLen {
				for j := range kvLen {
					if j > i {
						maskFlat[i*kvLen+j] = -1e9
					}
				}
			}
			causalMaskConst, err = f.Constant(maskFlat, 1, 1, seqLen, kvLen)
		case dtypes.Float16:
			maskFlat := make([]float16.Float16, seqLen*kvLen)
			negVal := float16.FromFloat32(-10000.0)
			for i := range seqLen {
				for j := range kvLen {
					if j > i {
						maskFlat[i*kvLen+j] = negVal
					}
				}
			}
			causalMaskConst, err = f.Constant(maskFlat, 1, 1, seqLen, kvLen)
		case dtypes.BFloat16:
			maskFlat := make([]bfloat16.BFloat16, seqLen*kvLen)
			negVal := bfloat16.FromFloat32(-1e9)
			for i := range seqLen {
				for j := range kvLen {
					if j > i {
						maskFlat[i*kvLen+j] = negVal
					}
				}
			}
			causalMaskConst, err = f.Constant(maskFlat, 1, 1, seqLen, kvLen)
		}
		if err != nil {
			return nil, nil, err
		}
		scoresVal, err = f.Add(scoresVal, causalMaskConst)
		if err != nil {
			return nil, nil, err
		}
	}

	// Apply explicit mask if provided
	if options != nil && options.Mask != nil {
		scoresNode := scoresVal.(*Node)
		maskAligned, err := alignAttentionScoreTensor(f, options.Mask, axesLayout, qBHSD, scoresNode)
		if err != nil {
			return nil, nil, errors.Wrap(err, "FusedScaledDotProductAttention: mask alignment failed")
		}
		maskNode := maskAligned.(*Node)

		if maskNode.shape.DType == dtypes.Bool {
			var negInfConst compute.Value
			switch qNode.shape.DType {
			case dtypes.Float32:
				negInfConst, err = f.Constant([]float32{-1e9})
			case dtypes.Float64:
				negInfConst, err = f.Constant([]float64{-1e9})
			case dtypes.Float16:
				negInfConst, err = f.Constant([]float16.Float16{float16.FromFloat32(-10000.0)})
			case dtypes.BFloat16:
				negInfConst, err = f.Constant([]bfloat16.BFloat16{bfloat16.FromFloat32(-1e9)})
			}
			if err != nil {
				return nil, nil, err
			}
			scoresVal, err = f.Where(maskNode, scoresVal, negInfConst)
			if err != nil {
				return nil, nil, err
			}
		} else {
			scoresVal, err = f.Add(scoresVal, maskNode)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	// Apply additive bias if provided
	if options != nil && options.Bias != nil {
		scoresNode := scoresVal.(*Node)
		biasAligned, err := alignAttentionScoreTensor(f, options.Bias, axesLayout, qBHSD, scoresNode)
		if err != nil {
			return nil, nil, errors.Wrap(err, "FusedScaledDotProductAttention: bias alignment failed")
		}
		scoresVal, err = f.Add(scoresVal, biasAligned)
		if err != nil {
			return nil, nil, err
		}
	}

	// Softmax along axis 3 (Skv)
	attnWeights, err := f.FusedSoftmax(scoresVal, 3)
	if err != nil {
		return nil, nil, errors.Wrap(err, "FusedScaledDotProductAttention: Softmax failed")
	}

	// Attn @ V -> [B, H, S, D]
	outBHSD, err := f.DotGeneral(attnWeights, []int{3}, []int{0, 1}, vBHSD, []int{2}, []int{0, 1}, compute.DotGeneralConfig{})
	if err != nil {
		return nil, nil, errors.Wrap(err, "FusedScaledDotProductAttention: attn @ v failed")
	}

	var finalOut compute.Value = outBHSD
	if axesLayout == compute.AttentionAxesLayoutBSHD {
		finalOut, err = f.Transpose(outBHSD, 0, 2, 1, 3)
		if err != nil {
			return nil, nil, err
		}
	}

	return finalOut, nil, nil
}

// FusedScaledDotProductAttentionVJP computes gradients of FusedScaledDotProductAttention.
func (f *Function) FusedScaledDotProductAttentionVJP(
	query, key, value compute.Value,
	axesLayout compute.AttentionAxesLayout,
	options *compute.ScaledDotProductAttentionConfig,
	output compute.Value, statesForVJP []compute.Value, dOutput compute.Value) (dQuery, dKey, dValue compute.Value, err error) {
	return nil, nil, nil, errors.Wrap(compute.ErrNotImplemented, "FusedScaledDotProductAttentionVJP is not implemented for onnxruntime backend")
}

// QuantizedEmbeddingLookup performs a quantized embedding lookup with on-the-fly dequantization.
func (f *Function) QuantizedEmbeddingLookup(data, indices compute.Value, dataQuantization *compute.Quantization) (compute.Value, error) {
	return nil, errors.Wrap(compute.ErrNotImplemented, "QuantizedEmbeddingLookup is not implemented for onnxruntime backend")
}

// FusedQuantizedDense performs fused dequantization + matmul + optional bias + optional activation.
func (f *Function) FusedQuantizedDense(x, weights, bias compute.Value, weightsQuantization *compute.Quantization, activation compute.ActivationType) (compute.Value, error) {
	if weightsQuantization == nil {
		return nil, errors.New("FusedQuantizedDense: weightsQuantization must not be nil")
	}

	if weightsQuantization.Scheme != compute.QuantLinear {
		return nil, errors.Wrapf(compute.ErrNotImplemented, "FusedQuantizedDense: scheme %v not supported for onnxruntime backend", weightsQuantization.Scheme)
	}

	xNode, ok1 := x.(*Node)
	wNode, ok2 := weights.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("FusedQuantizedDense: inputs must be valid onnxruntime nodes")
	}

	// Dequantize weights: floatWeight = (weights - zeroPoint) * scale
	wFloat, err := f.ConvertDType(wNode, dtypes.Float32)
	if err != nil {
		return nil, errors.Wrap(err, "FusedQuantizedDense: converting weights to float32 failed")
	}

	if weightsQuantization.ZeroPoint != nil {
		zpFloat, err := f.ConvertDType(weightsQuantization.ZeroPoint, dtypes.Float32)
		if err != nil {
			return nil, errors.Wrap(err, "FusedQuantizedDense: converting zeroPoint to float32 failed")
		}
		wNode := wFloat.(*Node)
		zpReshaped, err := broadcastToShape(f, zpFloat, wNode)
		if err != nil {
			return nil, errors.Wrap(err, "FusedQuantizedDense: zeroPoint reshape failed")
		}
		wFloatVal, err := f.Sub(wFloat, zpReshaped)
		if err != nil {
			return nil, errors.Wrap(err, "FusedQuantizedDense: sub zeroPoint failed")
		}
		wFloat = wFloatVal
	}

	if weightsQuantization.Scale != nil {
		scaleFloat, err := f.ConvertDType(weightsQuantization.Scale, dtypes.Float32)
		if err != nil {
			return nil, errors.Wrap(err, "FusedQuantizedDense: converting scale to float32 failed")
		}
		wNode := wFloat.(*Node)
		scaleReshaped, err := broadcastToShape(f, scaleFloat, wNode)
		if err != nil {
			return nil, errors.Wrap(err, "FusedQuantizedDense: scale reshape failed")
		}
		wFloatVal, err := f.Mul(wFloat, scaleReshaped)
		if err != nil {
			return nil, errors.Wrap(err, "FusedQuantizedDense: mul scale failed")
		}
		wFloat = wFloatVal
	}

	return f.FusedDense(xNode, wFloat, bias, compute.DenseConfig{
		Activation:   activation,
		WeightLayout: compute.DenseLayoutInputOutputs,
	})
}

// FusedAttentionQKVProjection performs fused Query-Key-Value projection.
func (f *Function) FusedAttentionQKVProjection(
	x, wQKV, biasQ, biasK, biasV compute.Value,
	queryDim, keyValueDim int) (query, key, value compute.Value, err error) {
	xNode, ok1 := x.(*Node)
	wNode, ok2 := wQKV.(*Node)
	if !ok1 || !ok2 {
		return nil, nil, nil, errors.New("FusedAttentionQKVProjection: inputs must be valid onnxruntime nodes")
	}

	xRank := xNode.shape.Rank()
	if xRank < 2 {
		return nil, nil, nil, errors.Errorf("FusedAttentionQKVProjection: x must be at least 2D [batch..., inFeatures], got rank %d", xRank)
	}

	xShape := xNode.shape
	inFeatures := xShape.Dimensions[xRank-1]
	batchDims := xShape.Dimensions[:xRank-1]
	flatBatch := 1
	for _, d := range batchDims {
		flatBatch *= d
	}

	var x2D *Node = xNode
	if xRank > 2 {
		xReshaped, err := f.Reshape(xNode, flatBatch, inFeatures)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: reshaping x to 2D failed")
		}
		x2D = xReshaped.(*Node)
	}

	// Compute matmul: y = x2D @ wQKV -> shape [flatBatch, queryDim + 2*keyValueDim]
	yVal, err := f.DotGeneral(x2D, []int{1}, nil, wNode, []int{0}, nil, compute.DotGeneralConfig{})
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: matmul failed")
	}

	totalDim := queryDim + 2*keyValueDim

	// Slice Q, K, V along axis 1
	qSlice, err := f.Slice(yVal, []int{0, 0}, []int{flatBatch, queryDim}, []int{1, 1})
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: slicing Q failed")
	}

	kSlice, err := f.Slice(yVal, []int{0, queryDim}, []int{flatBatch, queryDim + keyValueDim}, []int{1, 1})
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: slicing K failed")
	}

	vSlice, err := f.Slice(yVal, []int{0, queryDim + keyValueDim}, []int{flatBatch, totalDim}, []int{1, 1})
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: slicing V failed")
	}

	var queryRes, keyRes, valueRes compute.Value = qSlice, kSlice, vSlice

	if biasQ != nil {
		bqNode, ok := biasQ.(*Node)
		if !ok {
			return nil, nil, nil, errors.New("FusedAttentionQKVProjection: biasQ must be a valid onnxruntime node")
		}
		qNode := qSlice.(*Node)
		bqReshaped, err := broadcastToShape(f, bqNode, qNode)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: biasQ reshape failed")
		}
		queryRes, err = f.Add(queryRes, bqReshaped)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: adding biasQ failed")
		}
	}

	if biasK != nil {
		bkNode, ok := biasK.(*Node)
		if !ok {
			return nil, nil, nil, errors.New("FusedAttentionQKVProjection: biasK must be a valid onnxruntime node")
		}
		kNode := kSlice.(*Node)
		bkReshaped, err := broadcastToShape(f, bkNode, kNode)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: biasK reshape failed")
		}
		keyRes, err = f.Add(keyRes, bkReshaped)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: adding biasK failed")
		}
	}

	if biasV != nil {
		bvNode, ok := biasV.(*Node)
		if !ok {
			return nil, nil, nil, errors.New("FusedAttentionQKVProjection: biasV must be a valid onnxruntime node")
		}
		vNode := vSlice.(*Node)
		bvReshaped, err := broadcastToShape(f, bvNode, vNode)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: biasV reshape failed")
		}
		valueRes, err = f.Add(valueRes, bvReshaped)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: adding biasV failed")
		}
	}

	// If x was higher rank (> 2), reshape Q, K, V back to [batchDims..., queryDim/keyValueDim]
	if xRank > 2 {
		qDims := append(append([]int(nil), batchDims...), queryDim)
		kDims := append(append([]int(nil), batchDims...), keyValueDim)
		vDims := append(append([]int(nil), batchDims...), keyValueDim)

		queryRes, err = f.Reshape(queryRes, qDims...)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: reshaping Q back to original rank failed")
		}
		keyRes, err = f.Reshape(keyRes, kDims...)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: reshaping K back to original rank failed")
		}
		valueRes, err = f.Reshape(valueRes, vDims...)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "FusedAttentionQKVProjection: reshaping V back to original rank failed")
		}
	}

	return queryRes, keyRes, valueRes, nil
}
