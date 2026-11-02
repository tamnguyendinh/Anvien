// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"github.com/gomlx/compute"
)

func (f *Function) Equal(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addComparisonOp(compute.OpTypeEqual, "Equal", lhs, rhs)
}

func (f *Function) NotEqual(lhs, rhs compute.Value) (compute.Value, error) {
	eq, err := f.Equal(lhs, rhs)
	if err != nil {
		return nil, err
	}
	return f.LogicalNot(eq)
}

func (f *Function) GreaterThan(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addComparisonOp(compute.OpTypeGreaterThan, "Greater", lhs, rhs)
}

func (f *Function) GreaterOrEqual(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addComparisonOp(compute.OpTypeGreaterOrEqual, "GreaterOrEqual", lhs, rhs)
}

func (f *Function) LessThan(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addComparisonOp(compute.OpTypeLessThan, "Less", lhs, rhs)
}

func (f *Function) LessOrEqual(lhs, rhs compute.Value) (compute.Value, error) {
	return f.addComparisonOp(compute.OpTypeLessOrEqual, "LessOrEqual", lhs, rhs)
}
