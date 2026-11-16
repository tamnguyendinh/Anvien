// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"
	"io"

	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// RemapAndMarshalModel renames inputs and outputs on the ONNX ModelProto and serializes it to bytes.
func RemapAndMarshalModel(modelProto *onnx.ModelProto, inputNames []string, outputNames []string) ([]byte, error) {
	if modelProto == nil || modelProto.Graph == nil {
		return nil, errors.New("invalid ONNX model proto (nil Graph)")
	}

	graph := modelProto.Graph
	nameMap := make(map[string]string)

	// Remap input names
	for i, newName := range inputNames {
		if newName == "" {
			continue
		}
		var oldName string
		if i < len(graph.Input) {
			oldName = graph.Input[i].Name
			graph.Input[i].Name = newName
		} else {
			oldName = fmt.Sprintf("arg_%d", i)
		}
		if oldName != "" {
			nameMap[oldName] = newName
		}
	}

	// Remap output names
	for i, newName := range outputNames {
		if newName == "" {
			continue
		}
		var oldName string
		if i < len(graph.Output) {
			oldName = graph.Output[i].Name
			graph.Output[i].Name = newName
		} else {
			oldName = fmt.Sprintf("output_%d", i)
		}
		if oldName != "" {
			nameMap[oldName] = newName
		}
	}

	// Update node input/output tensor references if any names were mapped
	if len(nameMap) > 0 {
		for _, node := range graph.Node {
			for idx, inName := range node.Input {
				if replacement, ok := nameMap[inName]; ok {
					node.Input[idx] = replacement
				}
			}
			for idx, outName := range node.Output {
				if replacement, ok := nameMap[outName]; ok {
					node.Output[idx] = replacement
				}
			}
		}
	}

	modelBytes, err := proto.Marshal(modelProto)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal ONNX ModelProto")
	}

	return modelBytes, nil
}

// ParseModelProto parses ONNX model bytes and extracts input/output names and shapes.
func ParseModelProto(r io.Reader) (modelProto *onnx.ModelProto, modelBytes []byte, inputNames []string, inputShapes []shapes.Shape, outputNames []string, outputShapes []shapes.Shape, err error) {
	modelBytes, err = io.ReadAll(r)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, errors.Wrap(err, "failed to read model bytes")
	}

	modelProto = &onnx.ModelProto{}
	err = proto.Unmarshal(modelBytes, modelProto)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, errors.Wrap(err, "failed to unmarshal ONNX ModelProto")
	}

	if modelProto.Graph == nil {
		return nil, nil, nil, nil, nil, nil, errors.New("invalid ONNX model proto (nil Graph)")
	}

	graph := modelProto.Graph

	// Build map of initializers to filter out initializer names from inputs
	initializers := make(map[string]bool, len(graph.Initializer))
	for _, init := range graph.Initializer {
		initializers[init.Name] = true
	}

	for _, valInfo := range graph.Input {
		if initializers[valInfo.Name] {
			continue // Skip initializers declared in graph inputs
		}
		name := valInfo.Name
		shape := ONNXValueInfoToShape(valInfo)
		inputNames = append(inputNames, name)
		inputShapes = append(inputShapes, shape)
	}

	for _, valInfo := range graph.Output {
		name := valInfo.Name
		shape := ONNXValueInfoToShape(valInfo)
		outputNames = append(outputNames, name)
		outputShapes = append(outputShapes, shape)
	}

	return modelProto, modelBytes, inputNames, inputShapes, outputNames, outputShapes, nil
}

// ONNXValueInfoToShape converts an ONNX ValueInfoProto to a GoMLX Shape.
func ONNXValueInfoToShape(valInfo *onnx.ValueInfoProto) shapes.Shape {
	if valInfo == nil || valInfo.Type == nil {
		return shapes.Invalid()
	}
	tensorType := valInfo.Type.GetTensorType()
	if tensorType == nil {
		return shapes.Invalid()
	}

	dtype := ONNXDTypeToGoMLX(onnx.TensorProto_DataType(tensorType.ElemType))

	var dims []int
	var axisNames []string
	hasAxisNames := false
	if tensorType.Shape != nil {
		dims = make([]int, len(tensorType.Shape.Dim))
		axisNames = make([]string, len(tensorType.Shape.Dim))
		for i, d := range tensorType.Shape.Dim {
			if d.GetDimValue() > 0 {
				dims[i] = int(d.GetDimValue())
			} else {
				dims[i] = shapes.DynamicDim
				if param := d.GetDimParam(); param != "" {
					axisNames[i] = param
					hasAxisNames = true
				}
			}
		}
	}

	if hasAxisNames {
		return shapes.MakeDynamic(dtype, dims, axisNames)
	}
	return shapes.Make(dtype, dims...)
}

// ONNXDTypeToGoMLX maps an ONNX TensorProto DataType to GoMLX DType.
func ONNXDTypeToGoMLX(dt onnx.TensorProto_DataType) dtypes.DType {
	switch dt {
	case onnx.TensorProto_FLOAT:
		return dtypes.Float32
	case onnx.TensorProto_DOUBLE:
		return dtypes.Float64
	case onnx.TensorProto_INT32:
		return dtypes.Int32
	case onnx.TensorProto_INT64:
		return dtypes.Int64
	case onnx.TensorProto_BOOL:
		return dtypes.Bool
	case onnx.TensorProto_INT8:
		return dtypes.Int8
	case onnx.TensorProto_UINT8:
		return dtypes.Uint8
	case onnx.TensorProto_INT16:
		return dtypes.Int16
	case onnx.TensorProto_UINT16:
		return dtypes.Uint16
	case onnx.TensorProto_UINT32:
		return dtypes.Uint32
	case onnx.TensorProto_UINT64:
		return dtypes.Uint64
	default:
		return dtypes.InvalidDType
	}
}
