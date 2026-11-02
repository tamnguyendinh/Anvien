// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"

	"github.com/gomlx/compute"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/shapeinference"
	"github.com/pkg/errors"
)

func (f *Function) Add(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeAdd, "Add", lhs, rhs)
}

func (f *Function) Sub(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeSub, "Sub", lhs, rhs)
}

func (f *Function) Mul(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeMul, "Mul", lhs, rhs)
}

func (f *Function) Div(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeDiv, "Div", lhs, rhs)
}

func (f *Function) Rem(lhs, rhs compute.Value) (compute.Value, error) {
	lhsNode, ok1 := lhs.(*Node)
	rhsNode, ok2 := rhs.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("inputs must be valid onnxruntime nodes")
	}
	outShape, err := shapeinference.BinaryOp(compute.OpTypeRem, lhsNode.shape, rhsNode.shape)
	if err != nil {
		return nil, err
	}
	var attrs []*onnx.AttributeProto
	if lhsNode.shape.DType.IsFloat() {
		attrs = append(attrs, &onnx.AttributeProto{
			Name: "fmod",
			Type: onnx.AttributeProto_INT,
			I:    1,
		})
	}
	node := &Node{
		opType:     "Mod",
		inputs:     []*Node{lhsNode, rhsNode},
		shape:      outShape,
		attributes: attrs,
	}
	return f.addNode(node), nil
}

func (f *Function) Max(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeMax, "Max", lhs, rhs)
}

func (f *Function) Min(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeMin, "Min", lhs, rhs)
}

func (f *Function) Pow(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypePow, "Pow", lhs, rhs)
}

func (f *Function) BitwiseXor(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeBitwiseXor, "BitwiseXor", lhs, rhs)
}

func (f *Function) BitwiseAnd(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeBitwiseAnd, "BitwiseAnd", lhs, rhs)
}

func (f *Function) BitwiseOr(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeBitwiseOr, "BitwiseOr", lhs, rhs)
}

func (f *Function) ShiftLeft(lhs, rhs compute.Value) (compute.Value, error) {
	lhsNode, ok1 := lhs.(*Node)
	rhsNode, ok2 := rhs.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("ShiftLeft: inputs must be valid onnxruntime nodes")
	}
	outShape, err := shapeinference.BinaryOp(compute.OpTypeShiftLeft, lhsNode.shape, rhsNode.shape)
	if err != nil {
		return nil, err
	}
	f.nodeCount++
	node := &Node{
		name:   fmt.Sprintf("node_%d", f.nodeCount),
		opType: "BitShift",
		inputs: []*Node{lhsNode, rhsNode},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "direction",
				Type: onnx.AttributeProto_STRING,
				S:    []byte("LEFT"),
			},
		},
	}
	f.nodes = append(f.nodes, node)
	return node, nil
}

func (f *Function) ShiftRightLogical(lhs, rhs compute.Value) (compute.Value, error) {
	lhsNode, ok1 := lhs.(*Node)
	rhsNode, ok2 := rhs.(*Node)
	if !ok1 || !ok2 {
		return nil, errors.New("ShiftRightLogical: inputs must be valid onnxruntime nodes")
	}
	outShape, err := shapeinference.BinaryOp(compute.OpTypeShiftRightLogical, lhsNode.shape, rhsNode.shape)
	if err != nil {
		return nil, err
	}
	f.nodeCount++
	node := &Node{
		name:   fmt.Sprintf("node_%d", f.nodeCount),
		opType: "BitShift",
		inputs: []*Node{lhsNode, rhsNode},
		shape:  outShape,
		attributes: []*onnx.AttributeProto{
			{
				Name: "direction",
				Type: onnx.AttributeProto_STRING,
				S:    []byte("RIGHT"),
			},
		},
	}
	f.nodes = append(f.nodes, node)
	return node, nil
}
