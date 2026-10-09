// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"github.com/gomlx/compute"
)

func (f *Function) LogicalAnd(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeLogicalAnd, "And", lhs, rhs)
}

func (f *Function) LogicalOr(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeLogicalOr, "Or", lhs, rhs)
}

func (f *Function) LogicalXor(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addBinaryOp(compute.OpTypeLogicalXor, "Xor", lhs, rhs)
}

func (f *Function) LogicalNot(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeLogicalNot, "Not", x)
}
