package ortgenai

/*
#cgo CFLAGS: -O2 -g
#include "ort_genai_wrapper.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// Engine owns requests and executes them synchronously. Calls that touch an engine, its requests,
// or its event buffers must be serialized on the same goroutine pinned to an OS thread with
// runtime.LockOSThread. Engine does not create worker goroutines.
type Engine struct {
	enginePtr *C.OgaEngine
	requests  map[*C.OgaRequest]*Request
	buffers   map[*EventBuffer]struct{}
}

// RequestOptions contains request-scoped resident-session settings.
type RequestOptions struct{ MaxSessionTokens uint64 }

// TurnOptions is a reusable set of generation settings bound to one Request.
type TurnOptions struct {
	request *Request
	ptr     *C.OgaTurnOptions
}

// EventBuffer is reusable storage bound to one Engine.
type EventBuffer struct {
	engine *Engine
	ptr    *C.OgaEngineEventBuffer
}

// Request is permanently bound to its creating Engine.
type Request struct {
	engine *Engine
	ptr    *C.OgaRequest
	closed bool
}

type EventFlags uint32

const (
	EventFlagToken EventFlags = 1 << iota
	EventFlagTurnFinished
	EventFlagCapacityBlocked
	EventFlagFailed
	EventFlagRetryable
)

type FinishReason uint32

const (
	FinishReasonNone FinishReason = iota
	FinishReasonEOS
	FinishReasonStopString
	FinishReasonMaxGeneratedTokens
	FinishReasonMaxSessionTokens
	FinishReasonCancelled
	FinishReasonFailed
)

type ErrorCode uint32

const (
	ErrorCodeNone ErrorCode = iota
	ErrorCodeCapacityDeferred
	ErrorCodeExecutionCapacityExceeded
	ErrorCodeRetryableExecution
	ErrorCodeRequestUnserviceable
	ErrorCodeEngineContractFailure
	ErrorCodeEngineExecutionFailure
)

// TurnUsage is a copied snapshot; it remains valid after the next Engine.Run.
type TurnUsage struct{ PromptTokens, GeneratedTokens, CachedPromptTokens uint64 }

// Event is a copied engine event. Fields not present for an event remain at their zero value.
type Event struct {
	Flags                  EventFlags
	Request                *Request
	TurnID                 uint64
	Token                  int32
	FinishReason           FinishReason
	MatchedStopStringIndex int32
	ErrorCode              ErrorCode
	Usage                  *TurnUsage
}

// EngineCapabilities is a copied snapshot of configured engine limits.
type EngineCapabilities struct {
	ConfiguredMaxBatchSize int
	MaxScheduledTokens     int
	MaxRequestLength       uint64
}

// CreateEngine loads a model and creates an explicit, synchronous inference engine.
func CreateEngine(modelPath string) (*Engine, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if modelPath == "" {
		return nil, errors.New("model path is empty")
	}
	cPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cPath))
	var model *C.OgaModel
	if err := OgaResultToError(C.CreateOgaModel(cPath, &model)); err != nil {
		return nil, fmt.Errorf("creating model: %w", err)
	}
	if model == nil {
		return nil, errors.New("model creation returned nil without error")
	}
	defer C.DestroyOgaModel(model)
	return createEngineFromModel(model)
}

// CreateEngineWithRuntimeSettings creates an engine using owned runtime settings.
func CreateEngineWithRuntimeSettings(modelPath string, settings *RuntimeSettings) (*Engine, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if modelPath == "" {
		return nil, errors.New("model path is empty")
	}
	if settings == nil || settings.ptr == nil {
		return nil, errors.New("runtime settings are nil or destroyed")
	}
	cPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cPath))
	var model *C.OgaModel
	if err := OgaResultToError(C.CreateModelWithRuntimeSettings(cPath, settings.ptr, &model)); err != nil {
		return nil, fmt.Errorf("creating model with runtime settings: %w", err)
	}
	if model == nil {
		return nil, errors.New("model creation returned nil without error")
	}
	defer C.DestroyOgaModel(model)
	return createEngineFromModel(model)
}

func createEngineFromModel(model *C.OgaModel) (*Engine, error) {
	var ptr *C.OgaEngine
	if err := OgaResultToError(C.EngineCreate(model, &ptr)); err != nil {
		return nil, fmt.Errorf("creating engine: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("engine creation returned nil without error")
	}
	return &Engine{enginePtr: ptr, requests: make(map[*C.OgaRequest]*Request), buffers: make(map[*EventBuffer]struct{})}, nil
}

// CreateEventBuffer allocates reusable event storage for this Engine.
func (e *Engine) CreateEventBuffer(capacity int) (*EventBuffer, error) {
	if e == nil || e.enginePtr == nil {
		return nil, errors.New("engine is destroyed")
	}
	if capacity <= 0 {
		return nil, errors.New("event buffer capacity must be positive")
	}
	var ptr *C.OgaEngineEventBuffer
	if err := OgaResultToError(C.EngineCreateEventBuffer(e.enginePtr, C.size_t(capacity), &ptr)); err != nil {
		return nil, fmt.Errorf("creating event buffer: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("event buffer creation returned nil without error")
	}
	buffer := &EventBuffer{engine: e, ptr: ptr}
	e.buffers[buffer] = struct{}{}
	return buffer, nil
}

// Destroy releases the buffer and all borrowed native views it contained.
func (b *EventBuffer) Destroy() {
	if b == nil || b.ptr == nil {
		return
	}
	C.EngineDestroyEventBuffer(b.ptr)
	if b.engine != nil {
		delete(b.engine.buffers, b)
	}
	b.ptr = nil
	b.engine = nil
}

// CreateRequest creates a request permanently bound to this Engine.
func (e *Engine) CreateRequest(options *RequestOptions) (*Request, error) {
	if e == nil || e.enginePtr == nil {
		return nil, errors.New("engine is destroyed")
	}
	var nativeOptions *C.OgaRequestOptions
	if options != nil {
		if err := OgaResultToError(C.CreateRequestOptions(&nativeOptions)); err != nil {
			return nil, fmt.Errorf("creating request options: %w", err)
		}
		if nativeOptions == nil {
			return nil, errors.New("request options creation returned nil without error")
		}
		defer C.DestroyRequestOptions(nativeOptions)
		if options.MaxSessionTokens != 0 {
			if err := OgaResultToError(C.RequestOptionsSetMaxSessionTokens(nativeOptions, C.uint64_t(options.MaxSessionTokens))); err != nil {
				return nil, fmt.Errorf("setting maximum session tokens: %w", err)
			}
		}
	}
	var ptr *C.OgaRequest
	if err := OgaResultToError(C.EngineCreateRequest(e.enginePtr, nativeOptions, &ptr)); err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("request creation returned nil without error")
	}
	request := &Request{engine: e, ptr: ptr}
	e.requests[ptr] = request
	return request, nil
}

// CreateTurnOptions creates reusable generation settings for this Request.
func (r *Request) CreateTurnOptions() (*TurnOptions, error) {
	if r == nil || r.ptr == nil || r.closed {
		return nil, errors.New("request is closed or destroyed")
	}
	var ptr *C.OgaTurnOptions
	if err := OgaResultToError(C.RequestCreateTurnOptions(r.ptr, &ptr)); err != nil {
		return nil, fmt.Errorf("creating turn options: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("turn options creation returned nil without error")
	}
	return &TurnOptions{request: r, ptr: ptr}, nil
}

// Destroy releases native turn options.
func (o *TurnOptions) Destroy() {
	if o != nil && o.ptr != nil {
		C.DestroyTurnOptions(o.ptr)
		o.ptr = nil
		o.request = nil
	}
}
func (o *TurnOptions) check() error {
	if o == nil || o.ptr == nil || o.request == nil || o.request.ptr == nil || o.request.closed {
		return errors.New("turn options are destroyed or their request is closed")
	}
	return nil
}
func (o *TurnOptions) setResult(operation string, result *C.OgaResult) error {
	if err := OgaResultToError(result); err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}
	return nil
}
func (o *TurnOptions) SetMaxGeneratedTokens(v uint64) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting maximum generated tokens", C.TurnOptionsSetMaxGeneratedTokens(o.ptr, C.uint64_t(v)))
}
func (o *TurnOptions) SetMinGeneratedTokens(v uint64) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting minimum generated tokens", C.TurnOptionsSetMinGeneratedTokens(o.ptr, C.uint64_t(v)))
}
func (o *TurnOptions) SetDoSample(v bool) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting sampling mode", C.TurnOptionsSetDoSample(o.ptr, C.bool(v)))
}
func (o *TurnOptions) SetTemperature(v float32) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting temperature", C.TurnOptionsSetTemperature(o.ptr, C.float(v)))
}
func (o *TurnOptions) SetTopP(v float32) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting top-p", C.TurnOptionsSetTopP(o.ptr, C.float(v)))
}
func (o *TurnOptions) SetTopK(v int32) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting top-k", C.TurnOptionsSetTopK(o.ptr, C.int32_t(v)))
}
func (o *TurnOptions) SetRepetitionPenalty(v float32) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting repetition penalty", C.TurnOptionsSetRepetitionPenalty(o.ptr, C.float(v)))
}
func (o *TurnOptions) SetNoRepeatNgramSize(v int32) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting no-repeat n-gram size", C.TurnOptionsSetNoRepeatNgramSize(o.ptr, C.int32_t(v)))
}
func (o *TurnOptions) SetSeed(v uint64) error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("setting seed", C.TurnOptionsSetSeed(o.ptr, C.uint64_t(v)))
}
func (o *TurnOptions) ClearSeed() error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("clearing seed", C.TurnOptionsClearSeed(o.ptr))
}
func (o *TurnOptions) SetStopStrings(values []string) error {
	if err := o.check(); err != nil {
		return err
	}
	var list *C.OgaStringArray
	if err := OgaResultToError(C.CreateOgaStringArray(&list)); err != nil {
		return fmt.Errorf("creating stop-string array: %w", err)
	}
	if list == nil {
		return errors.New("stop-string array creation returned nil without error")
	}
	defer C.DestroyOgaStringArray(list)
	for _, value := range values {
		s := C.CString(value)
		result := C.AddStringToOgaStringArray(list, s)
		C.free(unsafe.Pointer(s))
		if err := OgaResultToError(result); err != nil {
			return fmt.Errorf("adding stop string: %w", err)
		}
	}
	return o.setResult("setting stop strings", C.TurnOptionsSetStopStrings(o.ptr, list))
}
func (o *TurnOptions) SetGuidance(kind GuidanceType, data string) error {
	if err := o.check(); err != nil {
		return err
	}
	t, d := C.CString(string(kind)), C.CString(data)
	defer C.free(unsafe.Pointer(t))
	defer C.free(unsafe.Pointer(d))
	return o.setResult("setting guidance", C.TurnOptionsSetGuidance(o.ptr, t, d))
}
func (o *TurnOptions) ClearGuidance() error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("clearing guidance", C.TurnOptionsClearGuidance(o.ptr))
}
func (o *TurnOptions) Reset() error {
	if err := o.check(); err != nil {
		return err
	}
	return o.setResult("resetting turn options", C.TurnOptionsReset(o.ptr))
}

// BeginTurn starts generation from token IDs and returns its unique turn ID.
func (r *Request) BeginTurn(inputIDs []int32, options ...*TurnOptions) (uint64, error) {
	if r == nil || r.ptr == nil || r.closed {
		return 0, errors.New("request is closed or destroyed")
	}
	if len(options) > 1 {
		return 0, errors.New("at most one turn options value may be supplied")
	}
	var nativeOptions *C.OgaTurnOptions
	if len(options) == 1 && options[0] != nil {
		if err := options[0].check(); err != nil {
			return 0, err
		}
		if options[0].request != r {
			return 0, errors.New("turn options belong to a different request")
		}
		nativeOptions = options[0].ptr
	}
	if len(inputIDs) == 0 {
		return 0, errors.New("turn input IDs are empty")
	}
	var ids *C.int32_t
	if len(inputIDs) > 0 {
		ids = (*C.int32_t)(unsafe.Pointer(&inputIDs[0]))
	}
	var id C.uint64_t
	if err := OgaResultToError(C.RequestBeginTurn(r.ptr, nativeOptions, ids, C.uint64_t(len(inputIDs)), &id)); err != nil {
		return 0, fmt.Errorf("beginning turn: %w", err)
	}
	return uint64(id), nil
}
func (r *Request) CancelTurn(id uint64) (bool, error) {
	if r == nil || r.ptr == nil || r.closed {
		return false, errors.New("request is closed or destroyed")
	}
	var cancelled C.bool
	if err := OgaResultToError(C.RequestCancelTurn(r.ptr, C.uint64_t(id), &cancelled)); err != nil {
		return false, fmt.Errorf("cancelling turn: %w", err)
	}
	return bool(cancelled), nil
}
func (r *Request) RewindToStartOfTurn(id uint64) error {
	if r == nil || r.ptr == nil || r.closed {
		return errors.New("request is closed or destroyed")
	}
	if err := OgaResultToError(C.RequestRewindToStartOfTurn(r.ptr, C.uint64_t(id))); err != nil {
		return fmt.Errorf("rewinding request: %w", err)
	}
	return nil
}

// Close removes this Request from scheduling; Destroy must still be called afterward.
func (r *Request) Close() error {
	if r == nil || r.ptr == nil || r.closed {
		return nil
	}
	if err := OgaResultToError(C.RequestClose(r.ptr)); err != nil {
		return fmt.Errorf("closing request: %w", err)
	}
	r.closed = true
	return nil
}

func (r *Request) SetDraftTokens(tokens []int32) error {
	if r == nil || r.ptr == nil || r.closed {
		return errors.New("request is closed or destroyed")
	}
	var sequences *C.OgaSequences
	if err := OgaResultToError(C.CreateOgaSequences(&sequences)); err != nil {
		return fmt.Errorf("creating draft sequence: %w", err)
	}
	if sequences == nil {
		return errors.New("draft sequence creation returned nil without error")
	}
	defer C.DestroyOgaSequences(sequences)
	if len(tokens) > 0 {
		if err := OgaResultToError(C.AppendTokenSequence((*C.int32_t)(unsafe.Pointer(&tokens[0])), C.size_t(len(tokens)), sequences)); err != nil {
			return fmt.Errorf("appending draft tokens: %w", err)
		}
	}
	if err := OgaResultToError(C.RequestSetDraftTokens(r.ptr, sequences)); err != nil {
		return fmt.Errorf("setting draft tokens: %w", err)
	}
	return nil
}

// Destroy releases the Request handle.
func (r *Request) Destroy() {
	if r == nil || r.ptr == nil {
		return
	}
	if !r.closed {
		_ = r.Close()
	}
	C.DestroyRequest(r.ptr)
	if r.engine != nil {
		delete(r.engine.requests, r.ptr)
	}
	r.ptr = nil
	r.engine = nil
}

func (e *Engine) MaxDraftTokensPerProposal() (int, error) {
	if e == nil || e.enginePtr == nil {
		return 0, errors.New("engine is destroyed")
	}
	var n C.size_t
	if err := OgaResultToError(C.EngineMaxDraftTokensPerProposal(e.enginePtr, &n)); err != nil {
		return 0, fmt.Errorf("getting draft token limit: %w", err)
	}
	return int(n), nil
}
func (e *Engine) HasPendingRequests() (bool, error) {
	if e == nil || e.enginePtr == nil {
		return false, errors.New("engine is destroyed")
	}
	var pending C.bool
	if err := OgaResultToError(C.EngineHasPendingRequests(e.enginePtr, &pending)); err != nil {
		return false, fmt.Errorf("checking pending requests: %w", err)
	}
	return bool(pending), nil
}

// Run advances the engine and copies all borrowed event data before returning.
func (e *Engine) Run(buffer *EventBuffer) ([]Event, error) {
	if e == nil || e.enginePtr == nil {
		return nil, errors.New("engine is destroyed")
	}
	if buffer == nil || buffer.ptr == nil || buffer.engine != e {
		return nil, errors.New("event buffer is nil, destroyed, or belongs to another engine")
	}
	if err := OgaResultToError(C.EngineRun(e.enginePtr, buffer.ptr)); err != nil {
		return nil, fmt.Errorf("running engine: %w", err)
	}
	count := int(C.EventBufferGetCount(buffer.ptr))
	events := make([]Event, 0, count)
	for i := range count {
		view := C.EventBufferGet(buffer.ptr, C.size_t(i))
		if view == nil {
			continue
		}
		var event Event
		var flags C.OgaEngineEventFlags
		if err := OgaResultToError(C.EngineEventGetFlags(view, &flags)); err != nil {
			return nil, fmt.Errorf("reading event flags: %w", err)
		}
		event.Flags = EventFlags(flags)
		var requestPtr *C.OgaRequest
		if err := OgaResultToError(C.EngineEventGetRequest(view, (**C.OgaRequest)(unsafe.Pointer(&requestPtr)))); err != nil {
			return nil, fmt.Errorf("reading event request: %w", err)
		}
		event.Request = e.requests[requestPtr]
		var id C.uint64_t
		if err := OgaResultToError(C.EngineEventGetTurnId(view, &id)); err != nil {
			return nil, fmt.Errorf("reading event turn ID: %w", err)
		}
		event.TurnID = uint64(id)
		if event.Flags&EventFlagToken != 0 {
			var token C.int32_t
			if err := OgaResultToError(C.EngineEventGetToken(view, &token)); err != nil {
				return nil, fmt.Errorf("reading event token: %w", err)
			}
			event.Token = int32(token)
		}
		if event.Flags&EventFlagTurnFinished != 0 {
			var reason C.OgaFinishReason
			var stop C.int32_t
			if err := OgaResultToError(C.EngineEventGetFinishReason(view, &reason)); err != nil {
				return nil, fmt.Errorf("reading finish reason: %w", err)
			}
			if err := OgaResultToError(C.EngineEventGetMatchedStopStringIndex(view, &stop)); err != nil {
				return nil, fmt.Errorf("reading stop-string index: %w", err)
			}
			event.FinishReason = FinishReason(reason)
			event.MatchedStopStringIndex = int32(stop)
			var usage *C.OgaTurnUsage
			if err := OgaResultToError(C.EngineEventGetUsage(view, (**C.OgaTurnUsage)(unsafe.Pointer(&usage)))); err != nil {
				return nil, fmt.Errorf("reading turn usage: %w", err)
			}
			if usage != nil {
				usageCopy := &TurnUsage{}
				var n C.uint64_t
				if err := OgaResultToError(C.TurnUsageGetPromptTokens(usage, &n)); err != nil {
					return nil, fmt.Errorf("reading prompt usage: %w", err)
				}
				usageCopy.PromptTokens = uint64(n)
				if err := OgaResultToError(C.TurnUsageGetGeneratedTokens(usage, &n)); err != nil {
					return nil, fmt.Errorf("reading generated usage: %w", err)
				}
				usageCopy.GeneratedTokens = uint64(n)
				if err := OgaResultToError(C.TurnUsageGetCachedPromptTokens(usage, &n)); err != nil {
					return nil, fmt.Errorf("reading cached prompt usage: %w", err)
				}
				usageCopy.CachedPromptTokens = uint64(n)
				event.Usage = usageCopy
			}
		}
		if event.Flags&(EventFlagFailed|EventFlagRetryable) != 0 {
			var code C.OgaErrorCode
			if err := OgaResultToError(C.EngineEventGetErrorCode(view, &code)); err != nil {
				return nil, fmt.Errorf("reading event error code: %w", err)
			}
			event.ErrorCode = ErrorCode(code)
		}
		events = append(events, event)
	}
	return events, nil
}

// SpeculativeStats returns an owned snapshot of engine-wide speculative decoding counters.
func (e *Engine) SpeculativeStats() (*SpeculativeStats, error) {
	if e == nil || e.enginePtr == nil {
		return nil, errors.New("engine is destroyed")
	}
	var ptr *C.OgaSpeculativeStats
	if err := OgaResultToError(C.EngineGetSpeculativeStats(e.enginePtr, &ptr)); err != nil {
		return nil, fmt.Errorf("getting engine speculative statistics: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("engine speculative statistics returned nil without error")
	}
	return &SpeculativeStats{ptr: ptr}, nil
}

// Capabilities returns a copied snapshot of the engine's configured runtime capabilities.
func (e *Engine) Capabilities() (EngineCapabilities, error) {
	if e == nil || e.enginePtr == nil {
		return EngineCapabilities{}, errors.New("engine is destroyed")
	}
	var ptr *C.OgaEngineCapabilities
	if err := OgaResultToError(C.EngineGetCapabilities(e.enginePtr, &ptr)); err != nil {
		return EngineCapabilities{}, fmt.Errorf("getting engine capabilities: %w", err)
	}
	if ptr == nil {
		return EngineCapabilities{}, errors.New("engine capabilities returned nil without error")
	}
	defer C.DestroyEngineCapabilities(ptr)
	return EngineCapabilities{
		ConfiguredMaxBatchSize: int(C.EngineCapabilitiesGetConfiguredMaxBatchSize(ptr)),
		MaxScheduledTokens:     int(C.EngineCapabilitiesGetMaxScheduledTokens(ptr)),
		MaxRequestLength:       uint64(C.EngineCapabilitiesGetMaxRequestLength(ptr)),
	}, nil
}

// Destroy releases requests, event buffers, and the engine in dependency order.
func (e *Engine) Destroy() {
	if e == nil || e.enginePtr == nil {
		return
	}
	for _, r := range e.requests {
		r.Destroy()
	}
	for b := range e.buffers {
		b.Destroy()
	}
	C.EngineDestroy(e.enginePtr)
	e.enginePtr = nil
}
