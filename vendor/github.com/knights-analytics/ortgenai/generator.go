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

// Generator owns a model, generation parameters, and an independent generator state.
type Generator struct {
	model  *model
	params *C.OgaGeneratorParams
	ptr    *C.OgaGenerator
}

// SpeculativeStats owns a native immutable statistics snapshot.
type SpeculativeStats struct{ ptr *C.OgaSpeculativeStats }

// CreateGenerator constructs an independent generator with the requested sequence and batch limits.
func CreateGenerator(modelPath string, maxLength, batchSize int) (*Generator, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if modelPath == "" || maxLength <= 0 || batchSize <= 0 {
		return nil, errors.New("model path, positive maximum length, and positive batch size are required")
	}
	path := C.CString(modelPath)
	defer C.free(unsafe.Pointer(path))
	var nativeModel *C.OgaModel
	if err := OgaResultToError(C.CreateOgaModel(path, &nativeModel)); err != nil {
		return nil, fmt.Errorf("creating generator model: %w", err)
	}
	if nativeModel == nil {
		return nil, errors.New("generator model creation returned nil without error")
	}
	g := &Generator{model: &model{modelPtr: nativeModel}}
	options := &GenerationOptions{MaxLength: maxLength, BatchSize: batchSize}
	params, err := createGeneratorParams(g.model, options)
	if err != nil {
		g.Destroy()
		return nil, err
	}
	g.params = params
	if err := OgaResultToError(C.CreateOgaGenerator(nativeModel, params, &g.ptr)); err != nil {
		g.Destroy()
		return nil, fmt.Errorf("creating generator: %w", err)
	}
	if g.ptr == nil {
		g.Destroy()
		return nil, errors.New("generator creation returned nil without error")
	}
	return g, nil
}

func (g *Generator) check() error {
	if g == nil || g.ptr == nil || g.model == nil || g.model.modelPtr == nil {
		return errors.New("generator is destroyed")
	}
	return nil
}

func (g *Generator) AppendTokens(ids []int32) error {
	if err := g.check(); err != nil {
		return err
	}
	if len(ids) == 0 {
		return errors.New("input token IDs are empty")
	}
	if err := OgaResultToError(C.GeneratorAppendTokens(g.ptr, (*C.int32_t)(unsafe.Pointer(&ids[0])), C.size_t(len(ids)))); err != nil {
		return fmt.Errorf("appending generator input: %w", err)
	}
	return nil
}

func (g *Generator) SetInputs(inputs *NamedTensors) error {
	if err := g.check(); err != nil {
		return err
	}
	if inputs == nil || inputs.tensorsPtr == nil {
		return errors.New("named tensors are destroyed")
	}
	if err := OgaResultToError(C.GeneratorSetInputs(g.ptr, inputs.tensorsPtr)); err != nil {
		return fmt.Errorf("setting generator inputs: %w", err)
	}
	return nil
}

func (g *Generator) SetModelInput(name string, tensor *Tensor) error {
	if err := g.check(); err != nil {
		return err
	}
	if err := tensor.check(); err != nil {
		return err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	if err := OgaResultToError(C.GeneratorSetModelInput(g.ptr, cName, tensor.ptr)); err != nil {
		return fmt.Errorf("setting model input: %w", err)
	}
	return nil
}

func (g *Generator) SetHiddenStates(tensor *Tensor) error {
	if err := g.check(); err != nil {
		return err
	}
	if err := tensor.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.GeneratorSetHiddenStates(g.ptr, tensor.ptr)); err != nil {
		return fmt.Errorf("setting hidden states: %w", err)
	}
	return nil
}

func (g *Generator) GenerateNextToken() error {
	if err := g.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.GeneratorGenerateNextToken(g.ptr)); err != nil {
		return fmt.Errorf("generating next token: %w", err)
	}
	return nil
}

func (g *Generator) IsDone() (bool, error) {
	if err := g.check(); err != nil {
		return false, err
	}
	return bool(C.IsDone(g.ptr)), nil
}

func (g *Generator) TokenCount() (int, error) {
	if err := g.check(); err != nil {
		return 0, err
	}
	return int(C.GeneratorTokenCount(g.ptr)), nil
}

// NextTokens returns a copied vector of tokens produced by the latest step.
func (g *Generator) NextTokens() ([]int32, error) {
	if err := g.check(); err != nil {
		return nil, err
	}
	var data *C.int32_t
	var count C.size_t
	if err := OgaResultToError(C.GeneratorGetNextTokens(g.ptr, (**C.int32_t)(unsafe.Pointer(&data)), &count)); err != nil {
		return nil, fmt.Errorf("reading next tokens: %w", err)
	}
	if count == 0 {
		return []int32{}, nil
	}
	if data == nil {
		return nil, errors.New("generator returned nil next-token data")
	}
	return append([]int32(nil), unsafe.Slice((*int32)(unsafe.Pointer(data)), int(count))...), nil
}

// Sequence returns a copied snapshot of one generated sequence.
func (g *Generator) Sequence(index int) ([]int32, error) {
	if err := g.check(); err != nil {
		return nil, err
	}
	if index < 0 {
		return nil, errors.New("sequence index must be non-negative")
	}
	count := int(C.GeneratorGetSequenceCount(g.ptr, C.size_t(index)))
	if count == 0 {
		return []int32{}, nil
	}
	data := C.GeneratorGetSequenceData(g.ptr, C.size_t(index))
	if data == nil {
		return nil, errors.New("generator returned nil sequence data")
	}
	return append([]int32(nil), unsafe.Slice((*int32)(unsafe.Pointer(data)), count)...), nil
}

func (g *Generator) RuntimeOption(key, value string) error {
	if err := g.check(); err != nil {
		return err
	}
	cKey, cValue := C.CString(key), C.CString(value)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cValue))
	if err := OgaResultToError(C.GeneratorSetRuntimeOption(g.ptr, cKey, cValue)); err != nil {
		return fmt.Errorf("setting runtime option: %w", err)
	}
	return nil
}

func (g *Generator) RewindTo(length int) error {
	if err := g.check(); err != nil {
		return err
	}
	if length < 0 {
		return errors.New("rewind length must be non-negative")
	}
	if err := OgaResultToError(C.GeneratorRewindTo(g.ptr, C.size_t(length))); err != nil {
		return fmt.Errorf("rewinding generator: %w", err)
	}
	return nil
}

func (g *Generator) SnapshotState() error {
	if err := g.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.GeneratorSnapshotState(g.ptr)); err != nil {
		return fmt.Errorf("snapshotting generator state: %w", err)
	}
	return nil
}

func (g *Generator) tensorResult(operation string, result *C.OgaResult, tensor *C.OgaTensor) (*Tensor, error) {
	if err := OgaResultToError(result); err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}
	if tensor == nil {
		return nil, fmt.Errorf("%s returned nil tensor without error", operation)
	}
	return &Tensor{ptr: tensor}, nil
}

func (g *Generator) Input(name string) (*Tensor, error) {
	if err := g.check(); err != nil {
		return nil, err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var tensor *C.OgaTensor
	result := C.GeneratorGetInput(g.ptr, cName, &tensor)
	return g.tensorResult("getting generator input", result, tensor)
}

func (g *Generator) Output(name string) (*Tensor, error) {
	if err := g.check(); err != nil {
		return nil, err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var tensor *C.OgaTensor
	result := C.GeneratorGetOutput(g.ptr, cName, &tensor)
	return g.tensorResult("getting generator output", result, tensor)
}

func (g *Generator) Logits() (*Tensor, error) {
	if err := g.check(); err != nil {
		return nil, err
	}
	var tensor *C.OgaTensor
	result := C.GeneratorGetLogits(g.ptr, &tensor)
	return g.tensorResult("getting generator logits", result, tensor)
}

func (g *Generator) SetLogits(tensor *Tensor) error {
	if err := g.check(); err != nil {
		return err
	}
	if err := tensor.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.GeneratorSetLogits(g.ptr, tensor.ptr)); err != nil {
		return fmt.Errorf("setting generator logits: %w", err)
	}
	return nil
}

func (g *Generator) SpeculativeStats() (*SpeculativeStats, error) {
	if err := g.check(); err != nil {
		return nil, err
	}
	var stats *C.OgaSpeculativeStats
	if err := OgaResultToError(C.GeneratorGetSpeculativeStats(g.ptr, &stats)); err != nil {
		return nil, fmt.Errorf("getting speculative statistics: %w", err)
	}
	if stats == nil {
		return nil, errors.New("speculative statistics returned nil without error")
	}
	return &SpeculativeStats{ptr: stats}, nil
}

func (s *SpeculativeStats) check() error {
	if s == nil || s.ptr == nil {
		return errors.New("speculative statistics are destroyed")
	}
	return nil
}

func (s *SpeculativeStats) Count(name string) (uint64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var value C.uint64_t
	if err := OgaResultToError(C.SpeculativeStatsGetCount(s.ptr, cName, &value)); err != nil {
		return 0, fmt.Errorf("reading speculative counter: %w", err)
	}
	return uint64(value), nil
}

func (s *SpeculativeStats) AcceptanceLengthCount(length int) (uint64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	if length < 0 {
		return 0, errors.New("acceptance length must be non-negative")
	}
	var count C.uint64_t
	if err := OgaResultToError(C.SpeculativeStatsGetAcceptanceLengthCount(s.ptr, C.size_t(length), &count)); err != nil {
		return 0, fmt.Errorf("reading acceptance-length count: %w", err)
	}
	return uint64(count), nil
}

func (s *SpeculativeStats) AcceptanceLengthHistogramSize() (int, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	var size C.size_t
	if err := OgaResultToError(C.SpeculativeStatsGetAcceptanceLengthHistogramSize(s.ptr, &size)); err != nil {
		return 0, fmt.Errorf("reading acceptance histogram size: %w", err)
	}
	return int(size), nil
}

func (s *SpeculativeStats) Number(name string) (float64, error) {
	if err := s.check(); err != nil {
		return 0, err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var value C.double
	if err := OgaResultToError(C.SpeculativeStatsGetNumber(s.ptr, cName, &value)); err != nil {
		return 0, fmt.Errorf("reading speculative statistic: %w", err)
	}
	return float64(value), nil
}

func (s *SpeculativeStats) Bool(name string) (bool, error) {
	if err := s.check(); err != nil {
		return false, err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var value C.bool
	if err := OgaResultToError(C.SpeculativeStatsGetBool(s.ptr, cName, &value)); err != nil {
		return false, fmt.Errorf("reading speculative statistic: %w", err)
	}
	return bool(value), nil
}

func (s *SpeculativeStats) Destroy() {
	if s != nil && s.ptr != nil {
		C.DestroySpeculativeStats(s.ptr)
		s.ptr = nil
	}
}

func (g *Generator) Destroy() {
	if g == nil {
		return
	}
	if g.ptr != nil {
		C.DestroyOgaGenerator(g.ptr)
		g.ptr = nil
	}
	if g.params != nil {
		C.DestroyOgaGeneratorParams(g.params)
		g.params = nil
	}
	if g.model != nil {
		g.model.destroy()
		g.model = nil
	}
}

// ASRPostProcessor incrementally converts generator output into text and timestamps.
// The associated Generator must remain alive until the processor is destroyed.
type ASRPostProcessor struct {
	ptr              *C.OgaASRPostProcessor
	base             *C.OgaPostProcessor
	generator        *Generator
	sessionGenerator *generator
}

type ASRTimestamp struct {
	Text       string
	StartFrame int64
	StopFrame  int64
	StartTime  float64
	StopTime   float64
}

type ASRResult struct {
	Text     string
	Words    []ASRTimestamp
	Segments []ASRTimestamp
}

// NewASRPostProcessor creates an ASR post-processor for this generator.
func (g *Generator) NewASRPostProcessor() (*ASRPostProcessor, error) {
	if err := g.check(); err != nil {
		return nil, err
	}
	return newASRPostProcessor(g.ptr, g, nil)
}

// newASRPostProcessor bridges the session's internal generator to the same
// native post-processor used by the public Generator API.
func (g *generator) newASRPostProcessor() (*ASRPostProcessor, error) {
	if g == nil || g.generatorPtr == nil {
		return nil, errors.New("generator is destroyed")
	}
	return newASRPostProcessor(g.generatorPtr, nil, g)
}

func newASRPostProcessor(nativeGenerator *C.OgaGenerator, publicGenerator *Generator, sessionGenerator *generator) (*ASRPostProcessor, error) {
	var ptr *C.OgaASRPostProcessor
	if err := OgaResultToError(C.CreateASRPostProcessor(nativeGenerator, &ptr)); err != nil {
		return nil, fmt.Errorf("creating ASR post-processor: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("ASR post-processor creation returned nil without error")
	}
	base := C.ASRPostProcessorAsPostProcessor(ptr)
	if base == nil {
		return nil, errors.New("ASR post-processor base view returned nil")
	}
	return &ASRPostProcessor{ptr: ptr, base: base, generator: publicGenerator, sessionGenerator: sessionGenerator}, nil
}

func (p *ASRPostProcessor) check() error {
	if p == nil || p.ptr == nil || p.base == nil ||
		(p.generator == nil && p.sessionGenerator == nil) ||
		(p.generator != nil && p.generator.ptr == nil) ||
		(p.sessionGenerator != nil && p.sessionGenerator.generatorPtr == nil) {
		return errors.New("ASR post-processor is destroyed")
	}
	return nil
}

// ProcessNext consumes the generator's current output and records any completed results.
func (p *ASRPostProcessor) ProcessNext() error {
	if err := p.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.PostProcessorProcessNext(p.base)); err != nil {
		return fmt.Errorf("processing ASR output: %w", err)
	}
	return nil
}

// Finalize flushes any partial result from the processor.
func (p *ASRPostProcessor) Finalize() error {
	if err := p.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.PostProcessorFinalize(p.base)); err != nil {
		return fmt.Errorf("finalizing ASR output: %w", err)
	}
	return nil
}

// Reset clears accumulated processor state.
func (p *ASRPostProcessor) Reset() error {
	if err := p.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.PostProcessorReset(p.base)); err != nil {
		return fmt.Errorf("resetting ASR post-processor: %w", err)
	}
	return nil
}

// Results copies all currently available text and word/segment timestamps.
func (p *ASRPostProcessor) Results() ([]ASRResult, error) {
	if err := p.check(); err != nil {
		return nil, err
	}
	count := int(C.ASRPostProcessorGetResultCount(p.ptr))
	results := make([]ASRResult, count)
	for i := range results {
		result := C.ASRPostProcessorGetResult(p.ptr, C.size_t(i))
		if result == nil {
			return nil, fmt.Errorf("ASR result %d is nil", i)
		}
		if text := C.ASRPostProcessorResultGetText(result); text != nil {
			results[i].Text = C.GoString(text)
		}
		timestamps := C.ASRPostProcessorResultGetTimestampResult(result)
		if timestamps == nil {
			continue
		}
		results[i].Words = copyASRTimestamps(timestamps, C.OgaTimestampRecordKind_Words, C.ASRPostProcessorTimestampResultGetWordCount(timestamps))
		results[i].Segments = copyASRTimestamps(timestamps, C.OgaTimestampRecordKind_Segments, C.ASRPostProcessorTimestampResultGetSegmentCount(timestamps))
	}
	return results, nil
}

func copyASRTimestamps(timestamps *C.OgaASRPostProcessorTimestampResult, kind C.OgaTimestampRecordKind, count C.size_t) []ASRTimestamp {
	values := make([]ASRTimestamp, int(count))
	for i := range values {
		record := C.ASRPostProcessorTimestampResultGetTimestampRecord(timestamps, kind, C.size_t(i))
		if record == nil {
			continue
		}
		if text := C.ASRPostProcessorTimestampRecordGetText(record); text != nil {
			values[i].Text = C.GoString(text)
		}
		values[i].StartFrame = int64(C.ASRPostProcessorTimestampRecordGetStartFrame(record))
		values[i].StopFrame = int64(C.ASRPostProcessorTimestampRecordGetStopFrame(record))
		values[i].StartTime = float64(C.ASRPostProcessorTimestampRecordGetStartTime(record))
		values[i].StopTime = float64(C.ASRPostProcessorTimestampRecordGetStopTime(record))
	}
	return values
}

// Destroy releases the processor. Destroy it before its Generator.
func (p *ASRPostProcessor) Destroy() {
	if p != nil && p.base != nil {
		C.DestroyPostProcessor(p.base)
		p.ptr = nil
		p.base = nil
		p.generator = nil
		p.sessionGenerator = nil
	}
}
