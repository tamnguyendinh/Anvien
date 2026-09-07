// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package onnxbackend

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/engine/web"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

// Type aliases for web engine types.
type Executable = web.Executable
type Buffer = web.Buffer

const SaveOnFailureEnv = "GOMLX_ONNX_SAVE_ON_FAILURE"

// IsSupportedPlatform returns true when compiled for js/wasm.
func IsSupportedPlatform() bool {
	return true
}

func parseConfig(config string) (ep executionprovider.Type, logSeverity int, enableGraphCapture bool, webVersion string, err error) {
	config = strings.TrimSpace(config)
	logSeverity = -1
	if config == "" {
		if web.HasWebGPU() {
			return executionprovider.WebGPU, -1, false, "", nil
		}
		return executionprovider.WASM, -1, false, "", nil
	}
	if strings.Contains(config, ":") || strings.EqualFold(config, "onnx") || strings.EqualFold(config, "onnxruntime") {
		parsed, errEnv := ParseGOMLXBackendEnv(config)
		if errEnv == nil {
			config = parsed
		}
	}
	parts := strings.Split(config, ",")
	hasEP := false
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if strings.Contains(part, "=") {
			kv := strings.SplitN(part, "=", 2)
			key := strings.ToLower(strings.TrimSpace(kv[0]))
			val := strings.TrimSpace(kv[1])
			if key == "log" {
				var level int
				if _, errScan := fmt.Sscanf(val, "%d", &level); errScan != nil {
					return executionprovider.CPU, 0, false, "", errors.Errorf("invalid log level: %q", val)
				}
				logSeverity = max(3-level, 0)
			} else if key == "graph_capture" || key == "graphcapture" || key == "enable_graph_capture" || key == "enablegraphcapture" {
				var errBool error
				enableGraphCapture, errBool = strconv.ParseBool(val)
				if errBool != nil {
					return executionprovider.CPU, 0, false, "", errors.Wrapf(errBool, "invalid boolean value for %q: %q", key, val)
				}
			} else if key == "web_version" || key == "webversion" {
				webVersion = val
			} else {
				return executionprovider.CPU, 0, false, "", errors.Errorf("unknown configuration option %q", key)
			}
			continue
		}
		partLower := strings.ToLower(part)
		if partLower == "webgpu" || partLower == "gpu" {
			ep = executionprovider.WebGPU
			hasEP = true
		} else if partLower == "wasm" || partLower == "cpu" {
			ep = executionprovider.WASM
			hasEP = true
		} else if partLower == "webgl" {
			ep = executionprovider.WebGL
			hasEP = true
		} else if partLower == "webnn" {
			ep = executionprovider.WebNN
			hasEP = true
		} else if partLower == "graph_capture" || partLower == "graphcapture" {
			enableGraphCapture = true
		} else {
			return executionprovider.CPU, 0, false, "", errors.Errorf("unknown web backend option: %q (expected \"webgpu\", \"wasm\", \"webgl\", \"webnn\", \"graph_capture\", or \"cpu\"/\"gpu\")", part)
		}
	}
	if !hasEP {
		if web.HasWebGPU() {
			ep = executionprovider.WebGPU
		} else {
			ep = executionprovider.WASM
		}
	}
	return ep, logSeverity, enableGraphCapture, webVersion, nil
}

// New creates a new ONNX Runtime Web backend instance.
func New(config string) (compute.Backend, error) {
	ep, logSeverity, enableGraphCapture, webVersion, err := parseConfig(config)
	if err != nil {
		return nil, err
	}

	if ep == executionprovider.WebGPU && !web.HasWebGPU() {
		return nil, errors.New("WebGPU is not available or failed to acquire a GPU adapter in the current browser environment (navigator.gpu is unavailable or requestAdapter() failed; in Chrome, this may require running with GPU hardware acceleration or enabling --enable-unsafe-webgpu)")
	}

	if ep == executionprovider.WebNN && !web.HasWebNN() {
		return nil, errors.New("WebNN is not available in the current browser environment (navigator.ml is undefined); WebNN is experimental in most browsers and typically requires enabling a browser flag such as chrome://flags/#web-machine-learning-neural-network)")
	}

	if err := web.EnsureORTLoaded(webVersion, ep.String()); err != nil {
		return nil, errors.Wrap(err, "failed to initialize onnxruntime-web")
	}

	hasFloat16 := false
	hasFloat64 := true
	if ep == executionprovider.WebGPU {
		hasFloat16 = web.HasWebGPUFloat16()
	} else if ep == executionprovider.WASM || ep == executionprovider.CPU {
		hasFloat16 = true // ORT Web WebAssembly CPU runtime supports float16 models
	}

	return &Backend{
		config:             config,
		version:            web.GetVersion(),
		executionProvider:  ep,
		webVersion:         webVersion,
		logSeverity:        logSeverity,
		enableGraphCapture: enableGraphCapture,
		hasFloat64:         hasFloat64,
		hasFloat16:         hasFloat16,
		hasBFloat16:        false,
	}, nil
}

func (b *Backend) createExecutable(modelBytes []byte, inputNames []string, inputShapes []shapes.Shape,
	outputNames []string, outputShapes []shapes.Shape, modelProto *onnx.ModelProto) (compute.Executable, error) {

	session, err := web.CreateSession(modelBytes, b.executionProvider, b.logSeverity, b.enableGraphCapture, b.webVersion)
	if err != nil {
		return nil, err
	}
	var savedModelProto *onnx.ModelProto
	if b.keepModelProto {
		savedModelProto = modelProto
	}
	return web.NewExecutable(b, session, inputNames, inputShapes, outputNames, outputShapes, savedModelProto), nil
}

func (b *Backend) BufferFromFlatData(deviceNum compute.DeviceNum, flat any, shape shapes.Shape) (compute.Buffer, error) {
	typedArray, err := web.ConvertSliceToTypedArray(flat, shape.DType)
	if err != nil {
		return nil, err
	}
	jsTensor, err := web.CreateJSTensor(shape.DType, shape.Dimensions, typedArray)
	if err != nil {
		return nil, err
	}
	wrapper := web.NewWebTensorWrapper(jsTensor, shape)
	return web.NewBuffer(b, wrapper, shape, deviceNum, false, nil), nil
}

func (b *Backend) HasSharedBuffers() bool {
	return false
}

func (b *Backend) NewSharedBuffer(deviceNum compute.DeviceNum, shape shapes.Shape) (compute.Buffer, any, error) {
	return nil, nil, errors.New("shared buffers are not supported in web/wasm backend")
}
