// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build !js

package native

import (
	"os"
	"strings"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	ort "github.com/gomlx/compute-onnx/internal/ort"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

// SaveOnFailureEnv is the environment variable that, when set to a file path,
// instructs the ONNX Runtime backend to save the serialized ONNX model protobuf to
// that path if graph compilation / session creation fails.
const SaveOnFailureEnv = "GOMLX_ONNX_SAVE_ON_FAILURE"

// MIGraphXOptions controls optional behavior of the MIGraphX execution provider.
type MIGraphXOptions struct {
	// CacheDir, when non-empty, enables compiled-program caching: the
	// MIGraphX-compiled program (.mxr) for each model is saved to this
	// directory on first compilation and loaded from it on subsequent runs,
	// skipping the expensive graph compilation. Entries are specific to the
	// model, GPU and MIGraphX/ORT versions.
	//
	// It is implemented through the ORT_MIGRAPHX_MODEL_CACHE_PATH environment
	// variable supported by AMD's ONNX Runtime MIGraphX builds, which avoids
	// the ABI-fragile OrtMIGraphXProviderOptions struct.
	CacheDir string
}

// CreateSession creates an ONNX Runtime DynamicAdvancedSession with the given options.
// The migraphxOptions arguments is only used for MIGraphX execution provider and may be nil.
func CreateSession(modelBytes []byte, inputNames []string, inputShapes []shapes.Shape, outputNames []string, executionProvider executionprovider.Type, logSeverity int, migraphxOptions *MIGraphXOptions) (*ort.DynamicAdvancedSession, error) {
	options, err := ort.NewSessionOptions()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ONNX Runtime SessionOptions")
	}
	defer options.Destroy()

	switch executionProvider {
	case executionprovider.CUDA:
		cudaOpts, err := ort.NewCUDAProviderOptions()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create ONNX Runtime CUDAProviderOptions")
		}
		defer cudaOpts.Destroy()

		_ = cudaOpts.Update(map[string]string{
			"do_copy_in_default_stream": "1",
		})

		err = options.AppendExecutionProviderCUDA(cudaOpts)
		if err != nil {
			return nil, WrapCUDAError(errors.Wrap(err, "failed to append CUDA execution provider to SessionOptions"))
		}

	case executionprovider.MIGraphX:
		// Upstream MIGraphX bug workaround: evaluating a program with scalar
		// (0-dimensional) inputs aborts inside migraphx::program::eval
		// ("contexts.size() == 1" assertion). Such models fall back to CPU execution.
		if hasScalarInput(inputShapes) {
			// TODO: Report an issue in https://github.com/ROCm/AMDMIGraphX .
			return nil, errors.Errorf("MIGraphX execution provider does not support scalar (0-dimensional) inputs; " +
				"run this model on CPU instead (maybe by setting GOMLX_BACKEND=\"onnx:cpu\")")
		}
		migraphxProviderOptions := &ort.MIGraphXProviderOptions{DeviceID: 0}
		if migraphxOptions != nil && migraphxOptions.CacheDir != "" {
			if err := os.MkdirAll(migraphxOptions.CacheDir, 0o755); err != nil {
				return nil, WrapMIGraphXError(errors.Wrapf(err, "failed to create MIGraphX cache directory %q", migraphxOptions.CacheDir))
			}
			if err := os.Setenv("ORT_MIGRAPHX_MODEL_CACHE_PATH", migraphxOptions.CacheDir); err != nil {
				return nil, WrapMIGraphXError(errors.Wrapf(err, "failed to set ORT_MIGRAPHX_MODEL_CACHE_PATH"))
			}
			klog.V(1).Infof("MIGraphX compiled-program caching enabled in %q", migraphxOptions.CacheDir)
		}
		err := options.AppendExecutionProviderMIGraphX(migraphxProviderOptions)
		if err != nil {
			return nil, WrapMIGraphXError(errors.Wrap(err, "failed to append MIGraphX execution provider to SessionOptions"))
		}
	case executionprovider.CPU:
		// CPU only.

	default:
		return nil, errors.Errorf("unknown gpu execution provider %q: expected \"cuda\", \"migraphx\" or \"\"", executionProvider)
	}

	logSev := logSeverity
	if logSev < 0 {
		logSev = 3 // ORT_LOGGING_LEVEL_ERROR by default
	}
	err = options.SetSessionLogSeverityLevel(logSev)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set ONNX Runtime session log severity level")
	}

	session, err := ort.NewDynamicAdvancedSessionWithONNXData(modelBytes, inputNames, outputNames, options)
	if err != nil {
		if filePath := os.Getenv(SaveOnFailureEnv); filePath != "" {
			if wErr := os.WriteFile(filePath, modelBytes, 0644); wErr == nil {
				klog.Infof("Saving failed model to %q ($%s)", filePath, SaveOnFailureEnv)
			} else {
				klog.Errorf("Failed to save failed model to %q ($%s): %+v", filePath, SaveOnFailureEnv, wErr)
			}
		}
		errStr := err.Error()
		if strings.Contains(errStr, "Type 'tensor(bfloat16)'") ||
			strings.Contains(errStr, "not supported for type 'bfloat16'") ||
			strings.Contains(errStr, "Could not find an implementation") {
			return nil, errors.Wrapf(compute.ErrNotImplemented, "ONNX doesn't support operation: %s", errStr)
		}
		return nil, WrapEPError(executionProvider, errors.Wrap(err, "failed to create ONNX Runtime session"))
	}

	return session, nil
}

// hasScalarInput returns whether any of the input shapes is 0-dimensional (a scalar).
func hasScalarInput(inputShapes []shapes.Shape) bool {
	for _, sh := range inputShapes {
		if sh.Rank() == 0 {
			return true
		}
	}
	return false
}

// WrapCUDAError adds helpful troubleshooting messages to known CUDA / ORT initialization errors.
func WrapCUDAError(err error) error {
	if err == nil {
		return nil
	}
	errStr := err.Error()
	if strings.Contains(errStr, "cudaLibraryGetKernel") ||
		strings.Contains(errStr, "OrtSessionOptionsAppendExecutionProvider_Cuda: Failed to load shared library") ||
		(strings.Contains(errStr, "AppendExecutionProvider_Cuda") && strings.Contains(errStr, "Failed to load")) {
		return errors.WithMessage(err, "CUDA versions <= 12.4 are compatible only with ORT (ONNX Runtime) up to v1.27. Maybe install an earlier version of ORT ? (see github.com/gomlx/compute-onnx/cmd/onnxruntime_installer)")
	}
	return err
}

// WrapEPError dispatches to the provider-specific error wrapper for the given
// execution provider, adding helpful troubleshooting messages to known errors.
func WrapEPError(ep executionprovider.Type, err error) error {
	if err == nil {
		return nil
	}
	switch ep {
	case executionprovider.CUDA:
		return WrapCUDAError(err)
	case executionprovider.MIGraphX:
		return WrapMIGraphXError(err)
	default:
		return err
	}
}

// WrapMIGraphXError adds helpful troubleshooting messages to known MIGraphX / ORT initialization errors.
func WrapMIGraphXError(err error) error {
	if err == nil {
		return nil
	}
	errStr := err.Error()
	if strings.Contains(errStr, "Failed to load shared library") ||
		strings.Contains(errStr, "libonnxruntime_providers_migraphx") {
		return errors.WithMessage(err,
			"failed to load the ONNX Runtime MIGraphX provider library (libonnxruntime_providers_migraphx.so): "+
				"make sure the ORT library was built with MIGraphX support (e.g. AMD's onnxruntime-migraphx wheel), "+
				"and that MIGraphX + HIP libraries are in the loader path")
	}
	if strings.Contains(errStr, "was not built with MIGraphX") {
		return err
	}
	if strings.Contains(errStr, "libmigraphx") || strings.Contains(errStr, "MIGraphX") {
		return errors.WithMessage(err,
			"MIGraphX libraries not found or failed to initialize: "+
				"install with `sudo apt install migraphx migraphx-dev half` and ensure /opt/rocm/lib is in LD_LIBRARY_PATH")
	}
	return err
}
