// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/shapes"
)

// Node represents an operation or constant in the ONNX computation graph.
type Node struct {
	name        string
	opType      string
	domain      string
	inputs      []*Node
	outputNames []string
	shape       shapes.Shape
	flatValue   any // used if opType == "Constant"
	attributes  []*onnx.AttributeProto
}

// Shape returns the shape of the node's output tensor.
func (n *Node) Shape() shapes.Shape {
	return n.shape
}

// Name returns the node's name.
func (n *Node) Name() string {
	return n.name
}

// OpType returns the node's operator type.
func (n *Node) OpType() string {
	return n.opType
}
