// Copyright 2023-2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build !js

package onnxbackend

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/device/cuda"
	"github.com/gomlx/compute-onnx/internal/device/rocm"
	"github.com/gomlx/compute-onnx/internal/engine/native"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	ort "github.com/gomlx/compute-onnx/internal/ort"
	"github.com/gomlx/compute-onnx/support/onnxruntime"
	onnx "github.com/gomlx/compute-onnx/support/protos"
	"github.com/gomlx/compute/shapes"
	"github.com/pkg/errors"
)

// Type aliases for native engine types.
type Executable = native.Executable
type Buffer = native.Buffer

const SaveOnFailureEnv = native.SaveOnFailureEnv

var (
	initMutex        sync.Mutex
	isOrtInitialized bool
	autoInstall      = true
)

// NoAutoInstallEnv is the environment variable that, when set, disables
// automatic downloading and installation of prebuilt ONNX Runtime shared libraries.
const NoAutoInstallEnv = "GOMLX_NO_AUTO_INSTALL"

// EnableAutoInstall controls whether the backend automatically downloads and installs
// prebuilt ONNX Runtime shared libraries if none are found.
func EnableAutoInstall(enable bool) {
	initMutex.Lock()
	defer initMutex.Unlock()
	autoInstall = enable
}

func init() {
	if _, found := os.LookupEnv(NoAutoInstallEnv); found {
		autoInstall = false
	}
}

// IsSupportedPlatform returns true if the host OS and CPU architecture are supported by ONNX Runtime binaries.
func IsSupportedPlatform() bool {
	switch runtime.GOOS {
	case "linux", "darwin", "windows":
		switch runtime.GOARCH {
		case "amd64", "arm64":
			return true
		}
	}
	return false
}

func isLibraryPath(part string) bool {
	partLower := strings.ToLower(part)
	if strings.HasSuffix(partLower, ".so") || strings.HasSuffix(partLower, ".dylib") || strings.HasSuffix(partLower, ".dll") {
		return true
	}
	if strings.Contains(part, "/") || strings.Contains(part, "\\") {
		return true
	}
	if info, err := os.Stat(part); err == nil && !info.IsDir() {
		return true
	}
	return false
}

func initializeORT(executionProvider executionprovider.Type, customLibPath string) error {
	initMutex.Lock()
	defer initMutex.Unlock()
	if isOrtInitialized {
		return nil
	}

	path := customLibPath
	if path == "" {
		path = os.Getenv("ONNXRUNTIME_SHARED_LIBRARY_PATH")
		if path == "" {
			installDir, err := onnxruntime.GetInstallPath()
			if executionProvider == executionprovider.MIGraphX {
				// The AMD ROCm build installs into its own directory, so it does not
				// clobber the standard (CPU/CUDA) library.
				installDir, err = onnxruntime.GetMigraphxInstallPath()
			}
			if err == nil {
				libFilename, err := onnxruntime.GetLibFilename()
				if err == nil {
					targetPath := filepath.Join(installDir, libFilename)
					if _, err := os.Stat(targetPath); err == nil {
						useInstalled := true
						switch executionProvider {
						case executionprovider.CUDA:
							cudaLibPath := filepath.Join(installDir, "libonnxruntime_providers_cuda.so")
							if _, err := os.Stat(cudaLibPath); err != nil {
								useInstalled = false
							}
						case executionprovider.MIGraphX:
							if !rocm.HasMigraphxExecutionProvider(installDir) {
								useInstalled = false
							}
						}
						if useInstalled {
							path = targetPath
						}
					}
				}
			}
		}
	}

	if path != "" {
		if _, err := os.Stat(path); err != nil {
			return errors.Wrapf(err, "ONNX Runtime library not found at specified path %q", path)
		}
	} else {
		if !autoInstall {
			return errors.Errorf("ONNX Runtime library not found (ONNXRUNTIME_SHARED_LIBRARY_PATH is not set) and auto-installation is disabled via %s or EnableAutoInstall(false)", NoAutoInstallEnv)
		}
		var err error
		switch executionProvider {
		case executionprovider.MIGraphX:
			path, err = onnxruntime.InstallMigraphx("", "", false)
		default:
			path, err = onnxruntime.Install(onnxruntime.DefaultVersion, executionProvider == executionprovider.CUDA, "", "", false)
		}
		if err != nil {
			return errors.Wrap(err, "failed to automatically install ONNX Runtime library")
		}
	}

	ort.SetSharedLibraryPath(path)
	err := ort.InitializeEnvironment()
	if err != nil {
		return errors.Wrap(err, "failed to initialize ONNX Runtime environment")
	}
	isOrtInitialized = true
	return nil
}

// parseConfig parses the backend configuration string and returns the selected GPU
// execution provider ("cuda", "migraphx", or "" for CPU), log severity, custom ORT library path,
// and the MIGraphX compiled-program cache directory ("" to disable caching).
func parseConfig(config string) (executionProvider executionprovider.Type, logSeverity int, customLibPath string, migraphxCacheDir string, err error) {
	executionProvider = executionprovider.CPU
	hasProvider := false
	logSeverity = -1 // not set
	migraphxCacheDir = os.Getenv("GOMLX_MIGRAPHX_CACHE_DIR")

	if config == "" {
		if envVal := os.Getenv("GOMLX_BACKEND"); envVal != "" {
			var errEnv error
			config, errEnv = ParseGOMLXBackendEnv(envVal)
			if errEnv != nil {
				return executionprovider.CPU, 0, "", "", errEnv
			}
		}
	} else if strings.Contains(config, ":") || strings.EqualFold(config, "onnx") || strings.EqualFold(config, "onnxruntime") {
		parsed, errEnv := ParseGOMLXBackendEnv(config)
		if errEnv == nil {
			config = parsed
		} else if !isLibraryPath(config) && !strings.Contains(config, "=") {
			return executionprovider.CPU, 0, "", "", errEnv
		}
	}

	config = strings.TrimSpace(config)
	if config == "" {
		envPath := os.Getenv("ONNXRUNTIME_SHARED_LIBRARY_PATH")
		if envPath != "" {
			return detectGPUProvider(filepath.Dir(envPath), false), -1, "", migraphxCacheDir, nil
		}
		if installDir, installErr := onnxruntime.GetInstallPath(); installErr == nil {
			return detectGPUProvider(installDir, true), -1, "", migraphxCacheDir, nil
		}
		return detectGPUProvider("", true), -1, "", migraphxCacheDir, nil
	}

	parts := strings.SplitSeq(config, ",")
	for part := range parts {
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
				if _, err := fmt.Sscanf(val, "%d", &level); err != nil {
					return executionprovider.CPU, 0, "", "", errors.Errorf("invalid log level: %q", val)
				}
				severity := max(3-level, 0)
				logSeverity = severity
			} else if key == "migraphx_cache_dir" || key == "migraphxcachedir" {
				migraphxCacheDir = val
			} else if key == "web_version" || key == "webversion" {
				// Ignored on native desktop platform.
			} else {
				return executionprovider.CPU, 0, "", "", errors.Errorf("unknown config option: %q", key)
			}
		} else {
			partLower := strings.ToLower(part)
			switch partLower {
			case "cuda", "gpu":
				executionProvider = executionprovider.CUDA
				hasProvider = true
			case "migraphx", "rocm", "amd":
				executionProvider = executionprovider.MIGraphX
				hasProvider = true
			case "cpu":
				executionProvider = executionprovider.CPU
				hasProvider = true
			default:
				if isLibraryPath(part) {
					customLibPath = part
				} else {
					return executionprovider.CPU, 0, "", "", errors.Errorf("invalid config value %q: expected \"cpu\", \"cuda\", \"migraphx\", path to ORT library, or key=value option", part)
				}
			}
		}
	}

	if !hasProvider {
		switch {
		case customLibPath != "":
			executionProvider = detectGPUProvider(filepath.Dir(customLibPath), false)
		case os.Getenv("ONNXRUNTIME_SHARED_LIBRARY_PATH") != "":
			executionProvider = detectGPUProvider(filepath.Dir(os.Getenv("ONNXRUNTIME_SHARED_LIBRARY_PATH")), false)
		default:
			dir, err := onnxruntime.GetInstallPath()
			if err != nil {
				dir = ""
			}
			executionProvider = detectGPUProvider(dir, true)
		}
	}
	return executionProvider, logSeverity, customLibPath, migraphxCacheDir, nil
}

// detectGPUProvider auto-detects which GPU execution provider to use: "cuda" if an NVIDIA GPU
// with CUDA libraries is present, "migraphx" if a discrete AMD GPU with MIGraphX libraries is present,
// or "" for CPU only.
// If dir is non-empty, it also requires the corresponding ORT provider library to be available in that directory.
// If allowDedicatedMigraphxDir is set, an ORT library previously installed in the dedicated
// MIGraphX directory also qualifies (so that auto-installation is never triggered implicitly
// by auto-detection).
func detectGPUProvider(dir string, allowDedicatedMigraphxDir bool) executionprovider.Type {
	if cuda.HasNvidiaGPU() && (dir == "" || cuda.IsCUDALibraryAvailable(dir)) {
		return executionprovider.CUDA
	}
	if rocm.HasAMDGPU() {
		if dir == "" || rocm.HasMigraphxExecutionProvider(dir) {
			return executionprovider.MIGraphX
		}
		if allowDedicatedMigraphxDir {
			if migraphxDir, err := onnxruntime.GetMigraphxInstallPath(); err == nil && rocm.HasMigraphxExecutionProvider(migraphxDir) {
				return executionprovider.MIGraphX
			}
		}
	}
	return executionprovider.CPU
}

// New creates a new ONNX Runtime backend instance with the given configuration string.
func New(config string) (compute.Backend, error) {
	if !IsSupportedPlatform() {
		return nil, errors.Errorf("onnxruntime backend is not supported on platform %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	executionProvider, logSeverity, customLibPath, migraphxCacheDir, err := parseConfig(config)
	if err != nil {
		return nil, err
	}

	switch executionProvider {
	case executionprovider.CUDA:
		if err := cuda.CheckCUDAAndCUDNN(); err != nil {
			return nil, err
		}
	case executionprovider.MIGraphX:
		if err := rocm.CheckROCmAndMIGraphX(); err != nil {
			return nil, err
		}
	}

	err = initializeORT(executionProvider, customLibPath)
	if err != nil {
		return nil, err
	}
	return &Backend{
		config:            config,
		version:           ort.GetVersion(),
		executionProvider: executionProvider,
		migraphxCacheDir:  migraphxCacheDir,
		logSeverity:       logSeverity,
		hasFloat64:        true,
		hasFloat16:        true,
		hasBFloat16:       executionProvider == executionprovider.CUDA,
	}, nil
}

func (b *Backend) createExecutable(modelBytes []byte, inputNames []string, inputShapes []shapes.Shape,
	outputNames []string, outputShapes []shapes.Shape, modelProto *onnx.ModelProto) (compute.Executable, error) {

	var migraphxOpts *native.MIGraphXOptions
	if b.executionProvider == executionprovider.MIGraphX && b.migraphxCacheDir != "" {
		migraphxOpts = &native.MIGraphXOptions{CacheDir: b.migraphxCacheDir}
	}
	session, err := native.CreateSession(modelBytes, inputNames, inputShapes, outputNames, b.executionProvider, b.logSeverity, migraphxOpts)
	if err != nil {
		return nil, err
	}
	var savedModelProto *onnx.ModelProto
	if b.keepModelProto {
		savedModelProto = modelProto
	}
	return native.NewExecutable(b, session, inputNames, inputShapes, outputNames, outputShapes, savedModelProto, b.executionProvider), nil
}

func (b *Backend) BufferFromFlatData(deviceNum compute.DeviceNum, flat any, shape shapes.Shape) (compute.Buffer, error) {
	wrapper, err := native.NewOrtTensorWrapper(shape, flat)
	if err != nil {
		return nil, err
	}
	return native.NewBuffer(b, wrapper, shape, deviceNum, b.executionProvider, nil), nil
}

func (b *Backend) HasSharedBuffers() bool {
	return b.executionProvider != executionprovider.CUDA
}

func (b *Backend) NewSharedBuffer(deviceNum compute.DeviceNum, shape shapes.Shape) (compute.Buffer, any, error) {
	wrapper, err := native.NewEmptyOrtTensorWrapper(shape)
	if err != nil {
		return nil, nil, err
	}
	flat := wrapper.GetData()
	buf := native.NewBuffer(b, wrapper, shape, deviceNum, b.executionProvider, nil)
	return buf, flat, nil
}
