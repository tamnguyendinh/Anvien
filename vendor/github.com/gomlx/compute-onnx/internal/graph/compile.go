// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package graph

import (
	"fmt"
	"runtime"

	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// CompiledModel contains the generated ONNX ModelProto and graph metadata.
type CompiledModel struct {
	Model        *onnx.ModelProto
	ModelBytes   []byte
	InputNames   []string
	InputShapes  []shapes.Shape
	OutputNames  []string
	OutputShapes []shapes.Shape
	UsedDTypes   map[dtypes.DType]bool
}

func ShapeToONNX(shape shapes.Shape) *onnx.TensorShapeProto {
	dims := shape.Dimensions
	onnxDims := make([]*onnx.TensorShapeProto_Dimension, len(dims))
	for i, d := range dims {
		if d == shapes.DynamicDim {
			paramName := shape.AxisName(i)
			if paramName == "" || paramName == shapes.AnonymousAxis {
				paramName = fmt.Sprintf("axis_%d", i)
			}
			onnxDims[i] = &onnx.TensorShapeProto_Dimension{
				Value: &onnx.TensorShapeProto_Dimension_DimParam{
					DimParam: paramName,
				},
			}
		} else {
			onnxDims[i] = &onnx.TensorShapeProto_Dimension{
				Value: &onnx.TensorShapeProto_Dimension_DimValue{
					DimValue: int64(d),
				},
			}
		}
	}
	return &onnx.TensorShapeProto{
		Dim: onnxDims,
	}
}

func dtypeToONNX(dt dtypes.DType) onnx.TensorProto_DataType {
	return DTypeToONNX(dt)
}

func DTypeToONNX(dt dtypes.DType) onnx.TensorProto_DataType {
	switch dt {
	case dtypes.Float32:
		return onnx.TensorProto_FLOAT
	case dtypes.Float64:
		return onnx.TensorProto_DOUBLE
	case dtypes.Float16:
		return onnx.TensorProto_FLOAT16
	case dtypes.BFloat16:
		return onnx.TensorProto_BFLOAT16
	case dtypes.Int32:
		return onnx.TensorProto_INT32
	case dtypes.Int64:
		return onnx.TensorProto_INT64
	case dtypes.Bool:
		return onnx.TensorProto_BOOL
	case dtypes.Int8:
		return onnx.TensorProto_INT8
	case dtypes.Uint8:
		return onnx.TensorProto_UINT8
	case dtypes.Int16:
		return onnx.TensorProto_INT16
	case dtypes.Uint16:
		return onnx.TensorProto_UINT16
	case dtypes.Uint32:
		return onnx.TensorProto_UINT32
	case dtypes.Uint64:
		return onnx.TensorProto_UINT64
	default:
		return onnx.TensorProto_UNDEFINED
	}
}

func constantToTensorProto(name string, shape shapes.Shape, flat any) *onnx.TensorProto {
	rawBytes := dtypes.UnsafeByteSliceFromAny(flat)
	dt := DTypeToONNX(shape.DType)

	dims := make([]int64, len(shape.Dimensions))
	for i, d := range shape.Dimensions {
		dims[i] = int64(d)
	}

	return &onnx.TensorProto{
		Dims:      dims,
		DataType:  int32(dt),
		Name:      name,
		RawData:   rawBytes,
		DocString: "Constant tensor",
	}
}

func topologicalSort(nodes []*Node, roots []*Node) []*Node {
	visited := make(map[*Node]bool)
	var order []*Node

	var visit func(n *Node)
	visit = func(n *Node) {
		if visited[n] {
			return
		}
		visited[n] = true
		for _, input := range n.inputs {
			visit(input)
		}
		order = append(order, n)
	}

	for _, root := range roots {
		visit(root)
	}
	return order
}

// CompileToProto translates a graph Builder AST into an ONNX ModelProto and serializes it to bytes.
func CompileToProto(b *Builder) (*CompiledModel, error) {
	mainFn := b.mainFn
	if mainFn == nil {
		return nil, errors.Errorf("builder %q has no main function", b.name)
	}

	if len(mainFn.returns) == 0 {
		return nil, errors.Errorf("main function in builder %q has no return values", b.name)
	}

	sortedNodes := topologicalSort(mainFn.nodes, mainFn.returns)

	var inputs []*onnx.ValueInfoProto
	var inputNames []string
	var inputShapes []shapes.Shape

	for _, param := range mainFn.params {
		inputNames = append(inputNames, param.name)
		inputShapes = append(inputShapes, param.shape)
		inputs = append(inputs, &onnx.ValueInfoProto{
			Name: param.name,
			Type: &onnx.TypeProto{
				Value: &onnx.TypeProto_TensorType{
					TensorType: &onnx.TypeProto_Tensor{
						ElemType: int32(DTypeToONNX(param.shape.DType)),
						Shape:    ShapeToONNX(param.shape),
					},
				},
			},
		})
	}

	// Workaround for ONNX Runtime limitation: graphs with 0 inputs fail session creation.
	// Inject a dummy 1-element tensor input if the function has 0 parameters.
	if len(inputs) == 0 {
		dummyName := "dummy_input"
		dummyShape := shapes.Make(dtypes.Int32, 1)
		inputNames = append(inputNames, dummyName)
		inputShapes = append(inputShapes, dummyShape)
		inputs = append(inputs, &onnx.ValueInfoProto{
			Name: dummyName,
			Type: &onnx.TypeProto{
				Value: &onnx.TypeProto_TensorType{
					TensorType: &onnx.TypeProto_Tensor{
						ElemType: int32(onnx.TensorProto_INT32),
						Shape:    ShapeToONNX(dummyShape),
					},
				},
			},
		})
	}

	var outputs []*onnx.ValueInfoProto
	var outputNames []string
	var outputShapes []shapes.Shape
	seenOutputs := make(map[string]bool)
	for i, ret := range mainFn.returns {
		retNode := ret
		if retNode.opType == "Parameter" || retNode.opType == "Constant" || seenOutputs[retNode.name] {
			// ONNX cannot have an input parameter or constant directly in graph.output without an op,
			// nor can multiple graph outputs share the same node name. Insert an Identity node.
			identNode := &Node{
				name:   fmt.Sprintf("%s_out_%d", retNode.name, i),
				opType: "Identity",
				inputs: []*Node{retNode},
				shape:  retNode.shape,
			}
			mainFn.nodes = append(mainFn.nodes, identNode)
			mainFn.returns[i] = identNode
			retNode = identNode
		}
		seenOutputs[retNode.name] = true
		outputNames = append(outputNames, retNode.name)
		outputShapes = append(outputShapes, retNode.shape)
		outputs = append(outputs, &onnx.ValueInfoProto{
			Name: retNode.name,
			Type: &onnx.TypeProto{
				Value: &onnx.TypeProto_TensorType{
					TensorType: &onnx.TypeProto_Tensor{
						ElemType: int32(DTypeToONNX(retNode.shape.DType)),
						Shape:    ShapeToONNX(retNode.shape),
					},
				},
			},
		})
	}

	sortedNodes = topologicalSort(mainFn.nodes, mainFn.returns)

	var onnxNodes []*onnx.NodeProto
	var onnxInitializers []*onnx.TensorProto

	for _, node := range sortedNodes {
		if node.opType == "Parameter" {
			continue
		}

		if node.opType == "" {
			// Sub-output node reference (e.g. 2nd output of a multi-output node like MaxPool)
			continue
		}

		if node.opType == "Constant" {
			tensorProto := constantToTensorProto(node.name, node.shape, node.flatValue)
			onnxInitializers = append(onnxInitializers, tensorProto)
			continue
		}

		if node.opType == "Identity" && len(node.inputs) == 1 && node.inputs[0].opType == "Constant" {
			tensorProto := constantToTensorProto(node.name, node.shape, node.inputs[0].flatValue)
			onnxInitializers = append(onnxInitializers, tensorProto)
			continue
		}

		nodeInputs := make([]string, len(node.inputs))
		for i, inp := range node.inputs {
			nodeInputs[i] = inp.name
		}

		nodeOutputs := node.outputNames
		if len(nodeOutputs) == 0 {
			nodeOutputs = []string{node.name}
		}

		nodeProto := &onnx.NodeProto{
			Input:     nodeInputs,
			Output:    nodeOutputs,
			OpType:    node.opType,
			Domain:    node.domain,
			Attribute: node.attributes,
		}
		onnxNodes = append(onnxNodes, nodeProto)
	}

	graph := &onnx.GraphProto{
		Name:        b.name,
		Node:        onnxNodes,
		Initializer: onnxInitializers,
		Input:       inputs,
		Output:      outputs,
	}

	model := &onnx.ModelProto{
		IrVersion: 9,
		Graph:     graph,
		OpsetImport: []*onnx.OperatorSetIdProto{
			{
				Domain:  "",
				Version: 21,
			},
			{
				Domain:  "com.microsoft",
				Version: 1,
			},
		},
	}

	modelBytes, err := proto.Marshal(model)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal ONNX ModelProto")
	}

	for _, node := range mainFn.nodes {
		node.flatValue = nil
	}
	runtime.GC()

	usedDTypes := make(map[dtypes.DType]bool)
	for _, n := range sortedNodes {
		if n.shape.DType != dtypes.InvalidDType {
			usedDTypes[n.shape.DType] = true
		}
	}

	return &CompiledModel{
		Model:        model,
		ModelBytes:   modelBytes,
		InputNames:   inputNames,
		InputShapes:  inputShapes,
		OutputNames:  outputNames,
		OutputShapes: outputShapes,
		UsedDTypes:   usedDTypes,
	}, nil
}
