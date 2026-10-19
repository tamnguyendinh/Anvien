// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"math"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/dtypes/bfloat16"
	"github.com/gomlx/compute/dtypes/float16"
	"github.com/pkg/errors"
)

func (f *Function) Abs(x compute.Value) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("Abs: input must be a valid onnxruntime node")
	}
	if xNode.shape.DType.IsComplex() {
		return nil, errors.Wrapf(compute.ErrNotImplemented, "Abs for complex dtype %s is not implemented for ONNX Runtime backend", xNode.shape.DType)
	}
	return f.addUnaryOp(compute.OpTypeAbs, "Abs", x)
}

func (f *Function) Neg(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeNeg, "Neg", x)
}

func (f *Function) Ceil(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeCeil, "Ceil", x)
}

func (f *Function) Floor(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeFloor, "Floor", x)
}

func (f *Function) Round(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeRound, "Round", x)
}

func (f *Function) Sqrt(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeSqrt, "Sqrt", x)
}

func (f *Function) Exp(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeExp, "Exp", x)
}

func (f *Function) Log(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeLog, "Log", x)
}

func (f *Function) Log1p(x compute.Value) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	var oneConst compute.Value
	var err error
	switch xNode.shape.DType {
	case dtypes.Float32:
		oneConst, err = f.Constant([]float32{1.0})
	case dtypes.Float64:
		oneConst, err = f.Constant([]float64{1.0})
	case dtypes.Float16:
		oneConst, err = f.Constant([]float16.Float16{float16.FromFloat32(1.0)})
	case dtypes.BFloat16:
		oneConst, err = f.Constant([]bfloat16.BFloat16{bfloat16.FromFloat32(1.0)})
	default:
		return nil, errors.Errorf("Log1p: unsupported DType %s", xNode.shape.DType)
	}
	if err != nil {
		return nil, err
	}

	xPlusOne, err := f.Add(x, oneConst)
	if err != nil {
		return nil, err
	}
	return f.Log(xPlusOne)
}

func (f *Function) Cos(x compute.Value) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	if xNode.shape.DType == dtypes.BFloat16 {
		x32, err := f.ConvertDType(xNode, dtypes.Float32)
		if err != nil {
			return nil, err
		}
		cos32, err := f.Cos(x32)
		if err != nil {
			return nil, err
		}
		return f.ConvertDType(cos32, dtypes.BFloat16)
	}

	// ONNX Runtime CPU provider lacks Cos for Float64 (double).
	// We compute cos(x) = sin(x + pi/2) when DType is Float64.
	if xNode.shape.DType == dtypes.Float64 {
		piOver2Const, err := f.Constant([]float64{math.Pi / 2.0})
		if err != nil {
			return nil, err
		}
		xPlusPiOver2, err := f.Add(xNode, piOver2Const)
		if err != nil {
			return nil, err
		}
		return f.Sin(xPlusPiOver2)
	}

	return f.addUnaryOp(compute.OpTypeCos, "Cos", x)
}

func (f *Function) Sin(x compute.Value) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	if xNode.shape.DType == dtypes.BFloat16 {
		x32, err := f.ConvertDType(xNode, dtypes.Float32)
		if err != nil {
			return nil, err
		}
		sin32, err := f.Sin(x32)
		if err != nil {
			return nil, err
		}
		return f.ConvertDType(sin32, dtypes.BFloat16)
	}

	return f.addUnaryOp(compute.OpTypeSin, "Sin", x)
}

func (f *Function) Tanh(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeTanh, "Tanh", x)
}

func (f *Function) Logistic(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeLogistic, "Sigmoid", x)
}

func (f *Function) Erf(x compute.Value) (compute.Value, error) {
	xNode, ok := x.(*Node)
	if !ok {
		return nil, errors.New("input must be a valid onnxruntime node")
	}

	// ONNX Runtime CPU provider lacks Erf for Float64 (double).
	// We convert to Float32, compute Erf, and convert back to Float64 when DType is Float64.
	if xNode.shape.DType == dtypes.Float64 {
		x32, err := f.ConvertDType(xNode, dtypes.Float32)
		if err != nil {
			return nil, err
		}
		erf32, err := f.addUnaryOp(compute.OpTypeErf, "Erf", x32)
		if err != nil {
			return nil, err
		}
		return f.ConvertDType(erf32, dtypes.Float64)
	}

	return f.addUnaryOp(compute.OpTypeErf, "Erf", x)
}

func (f *Function) Sign(x compute.Value) (compute.Value, error) {
	return f.addUnaryOp(compute.OpTypeSign, "Sign", x)
}
