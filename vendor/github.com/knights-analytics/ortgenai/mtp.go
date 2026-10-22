package ortgenai

/*
#cgo CFLAGS: -O2 -g
#include "ort_genai_wrapper.h"

static OgaResult* overlayMTPConfig(void* symbol, OgaConfig* config) {
	return ((OgaResult* (*)(OgaConfig*, const char*))symbol)(
		config, "{\"model\":{\"decoder\":{\"filename\":\"mtp.onnx\"}}}");
}
*/
import "C"

import (
	"errors"
	"fmt"
	"path/filepath"
	"unsafe"
)

// MTPGenerator owns both models, generation parameters, and an MTP speculative generator.
type MTPGenerator struct {
	mainModel *C.OgaModel
	mtpModel  *C.OgaModel
	params    *C.OgaGeneratorParams
	ptr       *C.OgaMtpGenerator
}

// MTPStatistics contains copied counters from the native MTP generator.
type MTPStatistics struct {
	ForwardCount uint64
	AcceptCount  uint64
	TrialCount   uint64
}

// CreateMTPGenerator creates a greedy, single-sequence self-speculative generator.
// When both paths refer to the same directory, the MTP model uses mtp.onnx
// and its external data alongside the main model, sharing the root config.
func CreateMTPGenerator(mainModelPath, mtpModelPath string, maxLength int) (*MTPGenerator, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if mainModelPath == "" || mtpModelPath == "" {
		return nil, errors.New("main and MTP model paths are required")
	}
	if maxLength <= 0 {
		return nil, errors.New("maximum sequence length must be positive")
	}
	mainPath, mtpPath := C.CString(mainModelPath), C.CString(mtpModelPath)
	defer C.free(unsafe.Pointer(mainPath))
	defer C.free(unsafe.Pointer(mtpPath))
	generator := &MTPGenerator{}
	if err := OgaResultToError(C.CreateOgaModel(mainPath, &generator.mainModel)); err != nil {
		return nil, fmt.Errorf("creating main model: %w", err)
	}
	var mtpErr error
	if filepath.Clean(mainModelPath) == filepath.Clean(mtpModelPath) {
		mtpErr = createRootMTPModel(mtpPath, &generator.mtpModel)
	} else {
		mtpErr = OgaResultToError(C.CreateOgaModel(mtpPath, &generator.mtpModel))
	}
	if mtpErr != nil {
		generator.Destroy()
		return nil, fmt.Errorf("creating MTP model: %w", mtpErr)
	}
	if err := OgaResultToError(C.CreateOgaGeneratorParams(generator.mainModel, &generator.params)); err != nil {
		generator.Destroy()
		return nil, fmt.Errorf("creating MTP generation parameters: %w", err)
	}
	for key, value := range map[string]float64{"batch_size": 1, "max_length": float64(maxLength)} {
		name := C.CString(key)
		result := C.GeneratorParamsSetSearchNumber(generator.params, name, C.double(value))
		C.free(unsafe.Pointer(name))
		if err := OgaResultToError(result); err != nil {
			generator.Destroy()
			return nil, fmt.Errorf("setting MTP parameter %s: %w", key, err)
		}
	}
	doSample := C.CString("do_sample")
	if err := OgaResultToError(C.GeneratorParamsSetSearchBool(generator.params, doSample, C.bool(false))); err != nil {
		C.free(unsafe.Pointer(doSample))
		generator.Destroy()
		return nil, fmt.Errorf("setting MTP sampling mode: %w", err)
	}
	C.free(unsafe.Pointer(doSample))
	if err := OgaResultToError(C.CreateMtpGenerator(generator.mainModel, generator.mtpModel, generator.params, &generator.ptr)); err != nil {
		generator.Destroy()
		return nil, fmt.Errorf("creating MTP generator: %w", err)
	}
	if generator.ptr == nil {
		generator.Destroy()
		return nil, errors.New("MTP generator creation returned nil without error")
	}
	return generator, nil
}

func createRootMTPModel(path *C.char, model **C.OgaModel) error {
	symbol := createSym(genAiLibraryHandle, "OgaConfigOverlay")
	if symbol == nil {
		return errors.New("missing OgaConfigOverlay")
	}
	var config *C.OgaConfig
	if err := OgaResultToError(C.CreateOgaConfig(path, &config)); err != nil {
		return fmt.Errorf("creating root MTP config: %w", err)
	}
	if config == nil {
		return errors.New("root MTP config creation returned nil without error")
	}
	defer C.DestroyOgaConfig(config)
	if err := OgaResultToError(C.overlayMTPConfig(symbol, config)); err != nil {
		return fmt.Errorf("setting root MTP decoder filename: %w", err)
	}
	return OgaResultToError(C.CreateOgaModelFromConfig(config, model))
}

func (g *MTPGenerator) check() error {
	if g == nil || g.ptr == nil {
		return errors.New("MTP generator is destroyed")
	}
	return nil
}

// AppendTokens seeds generation with copied prompt token IDs.
func (g *MTPGenerator) AppendTokens(ids []int32) error {
	if err := g.check(); err != nil {
		return err
	}
	if len(ids) == 0 {
		return errors.New("input token IDs are empty")
	}
	if err := OgaResultToError(C.MtpGeneratorAppendTokens(g.ptr, (*C.int32_t)(unsafe.Pointer(&ids[0])), C.size_t(len(ids)))); err != nil {
		return fmt.Errorf("appending MTP input tokens: %w", err)
	}
	return nil
}

// GenerateNextToken advances the draft/verify loop once.
func (g *MTPGenerator) GenerateNextToken() error {
	if err := g.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.MtpGeneratorGenerateNextToken(g.ptr)); err != nil {
		return fmt.Errorf("generating MTP token: %w", err)
	}
	return nil
}

func (g *MTPGenerator) IsDone() (bool, error) {
	if err := g.check(); err != nil {
		return false, err
	}
	return bool(C.MtpGeneratorIsDone(g.ptr)), nil
}

// Sequence returns a copied snapshot of the committed prompt and generated tokens.
func (g *MTPGenerator) Sequence() ([]int32, error) {
	if err := g.check(); err != nil {
		return nil, err
	}
	count := int(C.MtpGeneratorGetSequenceCount(g.ptr))
	if count == 0 {
		return []int32{}, nil
	}
	data := C.MtpGeneratorGetSequenceData(g.ptr)
	if data == nil {
		return nil, errors.New("MTP generator returned nil sequence data")
	}
	return append([]int32(nil), unsafe.Slice((*int32)(unsafe.Pointer(data)), count)...), nil
}

// Statistics returns a copied snapshot of MTP counters.
func (g *MTPGenerator) Statistics() (MTPStatistics, error) {
	if err := g.check(); err != nil {
		return MTPStatistics{}, err
	}
	return MTPStatistics{
		ForwardCount: uint64(C.MtpGeneratorGetForwardCount(g.ptr)),
		AcceptCount:  uint64(C.MtpGeneratorGetAcceptCount(g.ptr)),
		TrialCount:   uint64(C.MtpGeneratorGetTrialCount(g.ptr)),
	}, nil
}

// Reset clears the current MTP request while retaining generator allocations.
func (g *MTPGenerator) Reset() error {
	if err := g.check(); err != nil {
		return err
	}
	if err := OgaResultToError(C.MtpGeneratorReset(g.ptr)); err != nil {
		return fmt.Errorf("resetting MTP generator: %w", err)
	}
	return nil
}

// Destroy releases native generator and model resources in dependency order.
func (g *MTPGenerator) Destroy() {
	if g == nil {
		return
	}
	if g.ptr != nil {
		C.DestroyMtpGenerator(g.ptr)
		g.ptr = nil
	}
	if g.params != nil {
		C.DestroyOgaGeneratorParams(g.params)
		g.params = nil
	}
	if g.mtpModel != nil {
		C.DestroyOgaModel(g.mtpModel)
		g.mtpModel = nil
	}
	if g.mainModel != nil {
		C.DestroyOgaModel(g.mainModel)
		g.mainModel = nil
	}
}
