// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package web

import (
	"fmt"
	"strings"
	"syscall/js"
	"time"

	"github.com/gomlx/compute/dtypes"
	"github.com/gomlx/compute/dtypes/bfloat16"
	"github.com/gomlx/compute/dtypes/float16"
	"github.com/pkg/errors"
)

// Await awaits a JavaScript Promise and returns (result, error).
func Await(promise js.Value) (js.Value, error) {
	type resTuple struct {
		val js.Value
		err error
	}
	ch := make(chan resTuple, 1)

	var thenFn, catchFn js.Func
	thenFn = js.FuncOf(func(this js.Value, args []js.Value) any {
		var val js.Value
		if len(args) > 0 {
			val = args[0]
		}
		ch <- resTuple{val: val, err: nil}
		thenFn.Release()
		catchFn.Release()
		return nil
	})

	catchFn = js.FuncOf(func(this js.Value, args []js.Value) any {
		var errStr string
		if len(args) > 0 {
			errStr = args[0].Call("toString").String()
		} else {
			errStr = "unknown promise rejection"
		}
		ch <- resTuple{err: errors.New(errStr)}
		thenFn.Release()
		catchFn.Release()
		return nil
	})

	promise.Call("then", thenFn).Call("catch", catchFn)
	res := <-ch
	return res.val, res.err
}

// HasWebGPU checks if the WebGPU API (navigator.gpu) is present and an adapter is available in the browser.
func HasWebGPU() bool {
	global := js.Global()
	nav := global.Get("navigator")
	if nav.IsUndefined() || nav.IsNull() {
		return false
	}
	gpu := nav.Get("gpu")
	if gpu.IsUndefined() || gpu.IsNull() {
		return false
	}
	requestAdapter := gpu.Get("requestAdapter")
	if requestAdapter.IsUndefined() || requestAdapter.IsNull() {
		return false
	}
	promise := gpu.Call("requestAdapter")
	adapter, err := Await(promise)
	if err != nil || adapter.IsUndefined() || adapter.IsNull() {
		return false
	}
	return true
}

// GetWebGPUDevice returns the GPUDevice initialized by ORT Web or creates one if needed.
func GetWebGPUDevice() (js.Value, error) {
	global := js.Global()
	ortVal := global.Get("ort")
	if !ortVal.IsUndefined() && !ortVal.IsNull() {
		env := ortVal.Get("env")
		if !env.IsUndefined() && !env.IsNull() {
			webgpuEnv := env.Get("webgpu")
			if !webgpuEnv.IsUndefined() && !webgpuEnv.IsNull() {
				dev := webgpuEnv.Get("device")
				if !dev.IsUndefined() && !dev.IsNull() {
					return dev, nil
				}
			}
		}
	}

	// If not initialized yet by ORT, request from navigator.gpu
	nav := global.Get("navigator")
	if nav.IsUndefined() || nav.IsNull() {
		return js.Undefined(), errors.New("navigator is not available")
	}
	gpu := nav.Get("gpu")
	if gpu.IsUndefined() || gpu.IsNull() {
		return js.Undefined(), errors.New("navigator.gpu is not available")
	}
	adapterPromise := gpu.Call("requestAdapter")
	adapter, err := Await(adapterPromise)
	if err != nil || adapter.IsUndefined() || adapter.IsNull() {
		return js.Undefined(), errors.Wrap(err, "failed to get WebGPU adapter")
	}

	// Attempt requesting shader-f16 feature if supported by adapter
	var device js.Value
	features := adapter.Get("features")
	hasF16 := !features.IsUndefined() && !features.IsNull() && features.Call("has", "shader-f16").Bool()
	if hasF16 {
		reqDesc := global.Get("Object").New()
		reqFeatures := global.Get("Array").New()
		reqFeatures.Call("push", "shader-f16")
		reqDesc.Set("requiredFeatures", reqFeatures)
		devicePromise := adapter.Call("requestDevice", reqDesc)
		device, _ = Await(devicePromise)
	}

	if device.IsUndefined() || device.IsNull() {
		devicePromise := adapter.Call("requestDevice")
		device, err = Await(devicePromise)
		if err != nil || device.IsUndefined() || device.IsNull() {
			return js.Undefined(), errors.Wrap(err, "failed to get WebGPU device")
		}
	}

	// Register with ort.env.webgpu.device so ORT reuses this same device
	if !ortVal.IsUndefined() && !ortVal.IsNull() {
		env := ortVal.Get("env")
		if !env.IsUndefined() && !env.IsNull() {
			webgpuEnv := env.Get("webgpu")
			if webgpuEnv.IsUndefined() || webgpuEnv.IsNull() {
				webgpuEnv = global.Get("Object").New()
				env.Set("webgpu", webgpuEnv)
			}
			webgpuEnv.Set("device", device)
		}
	}
	return device, nil
}

// HasWebGPUFloat16 checks if WebGPU is available and the GPUDevice supports shader-f16.
func HasWebGPUFloat16() bool {
	dev, err := GetWebGPUDevice()
	if err != nil || dev.IsUndefined() || dev.IsNull() {
		return false
	}
	features := dev.Get("features")
	if features.IsUndefined() || features.IsNull() {
		return false
	}
	return features.Call("has", "shader-f16").Bool()
}

// HasWebNN checks if the WebNN API (navigator.ml) is present and available in the browser.
func HasWebNN() bool {
	global := js.Global()
	nav := global.Get("navigator")
	if nav.IsUndefined() || nav.IsNull() {
		return false
	}
	ml := nav.Get("ml")
	return !ml.IsUndefined() && !ml.IsNull()
}

// DefaultORTWebVersion is the default version of onnxruntime-web used if not specified.
// const DefaultORTWebVersion = "latest"
const DefaultORTWebVersion = "dev"

// NormalizeORTWebVersion normalizes user-provided version strings (e.g., "1.27", "v1.27", "@latest", "latest", "dev", "@dev")
// into a clean version tag or number suitable for npm CDN URLs.
func NormalizeORTWebVersion(v string) string {
	v = strings.TrimSpace(v)
	v = strings.TrimPrefix(v, "@")
	v = strings.TrimPrefix(v, "v")
	v = strings.TrimPrefix(v, "V")
	if v == "" {
		return DefaultORTWebVersion
	}
	return v
}

// EnsureORTLoaded checks if window.ort is defined. If not, it dynamically loads ort runtime via script tag or returns error.
// The version argument specifies the onnxruntime-web version to fetch from CDN (defaults to DefaultORTWebVersion if empty).
// The executionProvider argument hints which JS bundle to load (e.g. "ort.webgpu.min.js" for webgpu/webnn if desired, or standard "ort.min.js").
func EnsureORTLoaded(version string, executionProvider string) error {
	global := js.Global()
	ortVal := global.Get("ort")
	if !ortVal.IsUndefined() && !ortVal.IsNull() {
		return nil
	}

	// Try dynamically loading from CDN or local server
	doc := global.Get("document")
	if doc.IsUndefined() || doc.IsNull() {
		return errors.New("neither window.ort nor window.document is available")
	}

	normVersion := NormalizeORTWebVersion(version)
	scriptFile := "ort.min.js"

	scriptSrc := fmt.Sprintf("https://cdn.jsdelivr.net/npm/onnxruntime-web@%s/dist/%s", normVersion, scriptFile)
	wasmBasePath := fmt.Sprintf("https://cdn.jsdelivr.net/npm/onnxruntime-web@%s/dist/", normVersion)

	script := doc.Call("createElement", "script")
	script.Set("src", scriptSrc)

	loadedCh := make(chan error, 1)
	var onload, onerror js.Func
	onload = js.FuncOf(func(this js.Value, args []js.Value) any {
		loadedCh <- nil
		onload.Release()
		onerror.Release()
		return nil
	})
	onerror = js.FuncOf(func(this js.Value, args []js.Value) any {
		loadedCh <- errors.Errorf("failed to load onnxruntime-web script from %s", scriptSrc)
		onload.Release()
		onerror.Release()
		return nil
	})

	script.Set("onload", onload)
	script.Set("onerror", onerror)
	doc.Get("head").Call("appendChild", script)

	select {
	case err := <-loadedCh:
		if err != nil {
			return err
		}
	case <-time.After(15 * time.Second):
		onload.Release()
		onerror.Release()
		return errors.New("timed out waiting for onnxruntime-web script to load")
	}

	// Configure wasm paths and default logLevel
	ortVal = global.Get("ort")
	if ortVal.IsUndefined() || ortVal.IsNull() {
		return errors.New("window.ort is still undefined after script load")
	}
	env := ortVal.Get("env")
	if !env.IsUndefined() && !env.IsNull() {
		env.Set("logLevel", "error")
		wasmEnv := env.Get("wasm")
		if !wasmEnv.IsUndefined() && !wasmEnv.IsNull() {
			wasmEnv.Set("wasmPaths", wasmBasePath)
			wasmEnv.Set("numThreads", 1) // 1 thread by default for browser wasm stability
		}
	}

	return nil
}

// GetVersion returns the onnxruntime-web version string from window.ort.env.versions.web (or fallback).
func GetVersion() string {
	global := js.Global()
	ortVal := global.Get("ort")
	if !ortVal.IsUndefined() && !ortVal.IsNull() {
		env := ortVal.Get("env")
		if !env.IsUndefined() && !env.IsNull() {
			versions := env.Get("versions")
			if !versions.IsUndefined() && !versions.IsNull() {
				webVer := versions.Get("web")
				if !webVer.IsUndefined() && !webVer.IsNull() && webVer.String() != "" && webVer.String() != "undefined" {
					return webVer.String()
				}
				commonVer := versions.Get("common")
				if !commonVer.IsUndefined() && !commonVer.IsNull() && commonVer.String() != "" && commonVer.String() != "undefined" {
					return commonVer.String()
				}
			}
		}
		ver := ortVal.Get("version")
		if !ver.IsUndefined() && !ver.IsNull() && ver.String() != "" && ver.String() != "undefined" {
			return ver.String()
		}
	}
	return DefaultORTWebVersion
}

// ConvertSliceToTypedArray creates a JS TypedArray from a Go flat slice using fast CopyBytesToJS.
func ConvertSliceToTypedArray(flat any, dtype dtypes.DType) (js.Value, error) {
	global := js.Global()
	switch dtype {
	case dtypes.Float32:
		s, ok := flat.([]float32)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []float32, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Float32Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Float64:
		s, ok := flat.([]float64)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []float64, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Float64Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Int32:
		s, ok := flat.([]int32)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []int32, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Int32Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Int64:
		s, ok := flat.([]int64)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []int64, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("BigInt64Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Uint64:
		s, ok := flat.([]uint64)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []uint64, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("BigUint64Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Int8:
		s, ok := flat.([]int8)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []int8, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Int8Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Uint8, dtypes.Bool:
		switch s := flat.(type) {
		case []uint8:
			jsU8 := global.Get("Uint8Array").New(len(s))
			js.CopyBytesToJS(jsU8, s)
			return jsU8, nil
		case []bool:
			u8 := make([]byte, len(s))
			for i, v := range s {
				if v {
					u8[i] = 1
				}
			}
			jsU8 := global.Get("Uint8Array").New(len(u8))
			js.CopyBytesToJS(jsU8, u8)
			return jsU8, nil
		default:
			return js.Undefined(), errors.Errorf("expected []uint8 or []bool, got %T", flat)
		}

	case dtypes.Int16:
		s, ok := flat.([]int16)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []int16, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Int16Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Uint16:
		s, ok := flat.([]uint16)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []uint16, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Uint16Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Uint32:
		s, ok := flat.([]uint32)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []uint32, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Uint32Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.Float16:
		s, ok := flat.([]float16.Float16)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []float16.Float16, got %T", flat)
		}
		u8 := unsafeSliceBytes(s)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Uint16Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(s)), nil

	case dtypes.BFloat16:
		s, ok := flat.([]bfloat16.BFloat16)
		if !ok {
			return js.Undefined(), errors.Errorf("expected []bfloat16.BFloat16, got %T", flat)
		}
		f32 := make([]float32, len(s))
		for i, v := range s {
			f32[i] = v.Float32()
		}
		u8 := unsafeSliceBytes(f32)
		jsU8 := global.Get("Uint8Array").New(len(u8))
		js.CopyBytesToJS(jsU8, u8)
		return global.Get("Float32Array").New(jsU8.Get("buffer"), jsU8.Get("byteOffset"), len(f32)), nil

	default:
		return js.Undefined(), errors.Errorf("unsupported DType %s for ConvertSliceToTypedArray", dtype)
	}
}

// CopyTypedArrayToSlice copies data from a JS TypedArray to a Go slice using fast CopyBytesToGo.
func CopyTypedArrayToSlice(srcTypedArray js.Value, flat any, dtype dtypes.DType) error {
	global := js.Global()
	buffer := srcTypedArray.Get("buffer")
	byteOffset := srcTypedArray.Get("byteOffset").Int()
	byteLength := srcTypedArray.Get("byteLength").Int()

	jsU8 := global.Get("Uint8Array").New(buffer, byteOffset, byteLength)

	switch dtype {
	case dtypes.Float32:
		dst, ok := flat.([]float32)
		if !ok {
			return errors.Errorf("expected []float32, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Float64:
		dst, ok := flat.([]float64)
		if !ok {
			return errors.Errorf("expected []float64, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Int32:
		dst, ok := flat.([]int32)
		if !ok {
			return errors.Errorf("expected []int32, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Int64:
		dst, ok := flat.([]int64)
		if !ok {
			return errors.Errorf("expected []int64, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Uint64:
		dst, ok := flat.([]uint64)
		if !ok {
			return errors.Errorf("expected []uint64, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Int8:
		dst, ok := flat.([]int8)
		if !ok {
			return errors.Errorf("expected []int8, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Uint8:
		dst, ok := flat.([]uint8)
		if !ok {
			return errors.Errorf("expected []uint8, got %T", flat)
		}
		js.CopyBytesToGo(dst, jsU8)
		return nil

	case dtypes.Bool:
		dst, ok := flat.([]bool)
		if !ok {
			return errors.Errorf("expected []bool, got %T", flat)
		}
		u8 := make([]byte, len(dst))
		js.CopyBytesToGo(u8, jsU8)
		for i, b := range u8 {
			dst[i] = (b != 0)
		}
		return nil

	case dtypes.Int16:
		dst, ok := flat.([]int16)
		if !ok {
			return errors.Errorf("expected []int16, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Uint16:
		dst, ok := flat.([]uint16)
		if !ok {
			return errors.Errorf("expected []uint16, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Uint32:
		dst, ok := flat.([]uint32)
		if !ok {
			return errors.Errorf("expected []uint32, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.Float16:
		dst, ok := flat.([]float16.Float16)
		if !ok {
			return errors.Errorf("expected []float16.Float16, got %T", flat)
		}
		u8 := unsafeSliceBytes(dst)
		js.CopyBytesToGo(u8, jsU8)
		return nil

	case dtypes.BFloat16:
		dst, ok := flat.([]bfloat16.BFloat16)
		if !ok {
			return errors.Errorf("expected []bfloat16.BFloat16, got %T", flat)
		}
		f32 := make([]float32, len(dst))
		u8 := unsafeSliceBytes(f32)
		js.CopyBytesToGo(u8, jsU8)
		for i, v := range f32 {
			dst[i] = bfloat16.FromFloat32(v)
		}
		return nil

	default:
		return errors.Errorf("unsupported DType %s for CopyTypedArrayToSlice", dtype)
	}
}

// ORTTypeString maps a GoMLX DType to an onnxruntime-web type string.
func ORTTypeString(dtype dtypes.DType) (string, error) {
	switch dtype {
	case dtypes.Float32:
		return "float32", nil
	case dtypes.Float64:
		return "float64", nil
	case dtypes.Int32:
		return "int32", nil
	case dtypes.Int64:
		return "int64", nil
	case dtypes.Bool:
		return "bool", nil
	case dtypes.Int8:
		return "int8", nil
	case dtypes.Uint8:
		return "uint8", nil
	case dtypes.Int16:
		return "int16", nil
	case dtypes.Uint16:
		return "uint16", nil
	case dtypes.Uint32:
		return "uint32", nil
	case dtypes.Uint64:
		return "uint64", nil
	case dtypes.Float16:
		return "float16", nil
	case dtypes.BFloat16:
		return "float32", nil
	default:
		return "", errors.Errorf("unsupported dtype %s for onnxruntime-web", dtype)
	}
}

// CreateJSTensor creates an `ort.Tensor(type, data, dims)` JS object.
func CreateJSTensor(dtype dtypes.DType, dims []int, dataTypedArray js.Value) (js.Value, error) {
	ortType, err := ORTTypeString(dtype)
	if err != nil {
		return js.Undefined(), err
	}

	jsDims := make([]any, len(dims))
	for i, d := range dims {
		jsDims[i] = d
	}

	global := js.Global()
	ortVal := global.Get("ort")
	if ortVal.IsUndefined() || ortVal.IsNull() {
		return js.Undefined(), errors.New("window.ort is undefined")
	}
	tensorCtor := ortVal.Get("Tensor")
	tensor := tensorCtor.New(ortType, dataTypedArray, jsDims)
	return tensor, nil
}
