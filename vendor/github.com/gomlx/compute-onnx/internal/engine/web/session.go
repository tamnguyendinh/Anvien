// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package web

import (
	"strings"
	"syscall/js"

	"github.com/gomlx/compute"
	"github.com/gomlx/compute-onnx/internal/executionprovider"
	"github.com/pkg/errors"
)

// Session wraps a JavaScript ort.InferenceSession.
type Session struct {
	jsSession js.Value
	ep        executionprovider.Type
}

// CreateSession creates an onnxruntime-web InferenceSession from model bytes.
func CreateSession(modelBytes []byte, executionProvider executionprovider.Type, logSeverity int, enableGraphCapture bool, webVersion string) (*Session, error) {
	if err := EnsureORTLoaded(webVersion, executionProvider.String()); err != nil {
		return nil, errors.Wrap(err, "failed to initialize onnxruntime-web")
	}

	global := js.Global()
	ortVal := global.Get("ort")
	if ortVal.IsUndefined() || ortVal.IsNull() {
		return nil, errors.New("window.ort is not available")
	}

	inferenceSession := ortVal.Get("InferenceSession")
	if inferenceSession.IsUndefined() || inferenceSession.IsNull() {
		return nil, errors.New("ort.InferenceSession is not available")
	}

	// Create Uint8Array for model bytes
	jsU8 := global.Get("Uint8Array").New(len(modelBytes))
	js.CopyBytesToJS(jsU8, modelBytes)

	// Create session options
	options := global.Get("Object").New()
	eps := global.Get("Array").New()
	epStr := executionProvider.String()
	if executionProvider == executionprovider.CPU || epStr == "" {
		epStr = "wasm"
	}
	eps.Call("push", epStr)
	options.Set("executionProviders", eps)
	if enableGraphCapture && executionProvider == executionprovider.WebGPU {
		options.Set("enableGraphCapture", true)
	}

	logSev := logSeverity
	if logSev < 0 {
		logSev = 3 // ERROR by default
	}
	options.Set("logSeverityLevel", logSev)

	env := ortVal.Get("env")
	if !env.IsUndefined() && !env.IsNull() {
		var levelStr string
		switch logSev {
		case 0:
			levelStr = "verbose"
		case 1:
			levelStr = "info"
		case 2:
			levelStr = "warning"
		case 3:
			levelStr = "error"
		case 4:
			levelStr = "fatal"
		default:
			levelStr = "error"
		}
		env.Set("logLevel", levelStr)
	}

	createPromise := inferenceSession.Call("create", jsU8, options)
	jsSess, err := Await(createPromise)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "not supported") || strings.Contains(errStr, "Failed to find kernel") || strings.Contains(errStr, "incompatible") {
			return nil, errors.Wrapf(compute.ErrNotImplemented, "ort.InferenceSession.create failed: %s", errStr)
		}
		return nil, errors.Wrap(err, "ort.InferenceSession.create failed")
	}

	return &Session{
		jsSession: jsSess,
		ep:        executionProvider,
	}, nil
}

// Run executes the inference session with input feeds map and returns the results map.
func (s *Session) Run(feeds js.Value) (js.Value, error) {
	if s.jsSession.IsUndefined() || s.jsSession.IsNull() {
		return js.Undefined(), errors.New("cannot run on finalized session")
	}

	runPromise := s.jsSession.Call("run", feeds)
	results, err := Await(runPromise)
	if err != nil {
		return js.Undefined(), errors.Wrap(err, "session.run failed")
	}
	return results, nil
}

// RunWithFetches executes the inference session with pre-allocated output fetches.
func (s *Session) RunWithFetches(feeds, fetches js.Value) (js.Value, error) {
	if s.jsSession.IsUndefined() || s.jsSession.IsNull() {
		return js.Undefined(), errors.New("cannot run on finalized session")
	}

	runPromise := s.jsSession.Call("run", feeds, fetches)
	results, err := Await(runPromise)
	if err != nil {
		return js.Undefined(), errors.Wrap(err, "session.run with fetches failed")
	}
	return results, nil
}

// Destroy releases the session.
func (s *Session) Destroy() error {
	s.jsSession = js.Undefined()
	return nil
}
