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
