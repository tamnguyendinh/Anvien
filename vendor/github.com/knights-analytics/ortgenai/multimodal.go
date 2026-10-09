package ortgenai

/*
#cgo CFLAGS: -O2 -g
#include "ort_genai_wrapper.h"
*/
import "C"

import (
	"encoding/base64"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"unsafe"
)

// Images represents a collection of loaded images for multimodal processing.
type Images struct {
	imagesPtr *C.OgaImages
}

func (i *Images) destroy() {
	if i.imagesPtr != nil {
		C.DestroyOgaImages(i.imagesPtr)
	}
	i.imagesPtr = nil
}

// Destroy releases the images resources.
func (i *Images) Destroy() {
	i.destroy()
}

// MultiModalProcessor processes images and text together for multimodal models.
type multiModalProcessor struct {
	processorPtr *C.OgaMultiModalProcessor
}

func (p *multiModalProcessor) destroy() {
	if p.processorPtr != nil {
		C.DestroyOgaMultiModalProcessor(p.processorPtr)
	}
	p.processorPtr = nil
}

// Destroy releases the processor resources.
func (p *multiModalProcessor) Destroy() {
	p.destroy()
}

// NamedTensors represents a collection of named tensor inputs.
type NamedTensors struct {
	tensorsPtr *C.OgaNamedTensors
}

// Audios owns one or more loaded audio samples.
type Audios struct {
	audiosPtr *C.OgaAudios
}

// AudioProcessor owns a model and its multimodal processor.
type AudioProcessor struct {
	modelPtr     *C.OgaModel
	processorPtr *C.OgaMultiModalProcessor
}

// Adapters manages model adapters and must be destroyed before its source Session.
type Adapters struct {
	ptr     *C.OgaAdapters
	session *Session
}

func (nt *NamedTensors) destroy() {
	if nt.tensorsPtr != nil {
		C.DestroyOgaNamedTensors(nt.tensorsPtr)
	}
	nt.tensorsPtr = nil
}

// Destroy releases the named tensors resources.
func (nt *NamedTensors) Destroy() {
	nt.destroy()
}

func createOgaStringArray(values []string) (*C.OgaStringArray, error) {
	var stringArray *C.OgaStringArray
	if err := OgaResultToError(C.CreateOgaStringArray(&stringArray)); err != nil {
		return nil, fmt.Errorf("creating string array: %w", err)
	}
	if stringArray == nil {
		return nil, errors.New("string array creation returned nil without error")
	}
	for _, value := range values {
		cValue := C.CString(value)
		result := C.AddStringToOgaStringArray(stringArray, cValue)
		C.free(unsafe.Pointer(cValue))
		if err := OgaResultToError(result); err != nil {
			C.DestroyOgaStringArray(stringArray)
			return nil, fmt.Errorf("adding string to array: %w", err)
		}
	}
	return stringArray, nil
}

// LoadAudio loads one audio sample from a path.
func LoadAudio(path string) (*Audios, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	var ptr *C.OgaAudios
	if err := OgaResultToError(C.LoadAudio(cPath, &ptr)); err != nil {
		return nil, fmt.Errorf("loading audio: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("audio loading returned nil without error")
	}
	return &Audios{audiosPtr: ptr}, nil
}

// LoadAudios loads audio samples from paths.
func LoadAudios(paths []string) (*Audios, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if len(paths) == 0 {
		return nil, errors.New("audio paths are empty")
	}
	stringArray, err := createOgaStringArray(paths)
	if err != nil {
		return nil, err
	}
	defer C.DestroyOgaStringArray(stringArray)
	var ptr *C.OgaAudios
	if err = OgaResultToError(C.LoadAudios(stringArray, &ptr)); err != nil {
		return nil, fmt.Errorf("loading audios: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("audio loading returned nil without error")
	}
	return &Audios{audiosPtr: ptr}, nil
}

// LoadAudiosFromBuffers loads audio samples, copying the caller's bytes during the call.
func LoadAudiosFromBuffers(buffers [][]byte) (*Audios, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if len(buffers) == 0 {
		return nil, errors.New("audio buffers are empty")
	}
	data := make([]unsafe.Pointer, len(buffers))
	sizes := make([]C.size_t, len(buffers))
	for i, buffer := range buffers {
		if len(buffer) == 0 {
			for _, ptr := range data {
				C.free(ptr)
			}
			return nil, fmt.Errorf("audio buffer %d is empty", i)
		}
		data[i] = C.CBytes(buffer)
		sizes[i] = C.size_t(len(buffer))
	}
	defer func() {
		for _, ptr := range data {
			C.free(ptr)
		}
	}()
	var ptr *C.OgaAudios
	if err := OgaResultToError(C.LoadAudiosFromBuffers(&data[0], (*C.size_t)(unsafe.Pointer(&sizes[0])), C.size_t(len(data)), &ptr)); err != nil {
		return nil, fmt.Errorf("loading audio buffers: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("audio buffer loading returned nil without error")
	}
	return &Audios{audiosPtr: ptr}, nil
}

// Destroy releases loaded audio data.
func (a *Audios) Destroy() {
	if a != nil && a.audiosPtr != nil {
		C.DestroyAudios(a.audiosPtr)
		a.audiosPtr = nil
	}
}

// CreateAudioProcessor creates an audio/multimodal processor from a model directory.
func CreateAudioProcessor(modelPath string) (*AudioProcessor, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	path := C.CString(modelPath)
	defer C.free(unsafe.Pointer(path))
	var model *C.OgaModel
	if err := OgaResultToError(C.CreateOgaModel(path, &model)); err != nil {
		return nil, fmt.Errorf("creating audio processor model: %w", err)
	}
	if model == nil {
		return nil, errors.New("audio processor model creation returned nil without error")
	}
	var processor *C.OgaMultiModalProcessor
	if err := OgaResultToError(C.CreateOgaMultiModalProcessor(model, &processor)); err != nil {
		C.DestroyOgaModel(model)
		return nil, fmt.Errorf("creating audio processor: %w", err)
	}
	if processor == nil {
		C.DestroyOgaModel(model)
		return nil, errors.New("audio processor creation returned nil without error")
	}
	return &AudioProcessor{modelPtr: model, processorPtr: processor}, nil
}

// ProcessAudios converts audio and a prompt to copied-owner named tensors.
func (p *AudioProcessor) ProcessAudios(prompt string, audios *Audios) (*NamedTensors, error) {
	if p == nil || p.processorPtr == nil || audios == nil || audios.audiosPtr == nil {
		return nil, errors.New("audio processor or audios are destroyed")
	}
	cPrompt := C.CString(prompt)
	defer C.free(unsafe.Pointer(cPrompt))
	var tensors *C.OgaNamedTensors
	if err := OgaResultToError(C.ProcessAudios(p.processorPtr, cPrompt, audios.audiosPtr, &tensors)); err != nil {
		return nil, fmt.Errorf("processing audios: %w", err)
	}
	if tensors == nil {
		return nil, errors.New("audio processing returned nil tensors without error")
	}
	return &NamedTensors{tensorsPtr: tensors}, nil
}

func (p *AudioProcessor) ProcessAudiosAndPrompts(prompts []string, audios *Audios) (*NamedTensors, error) {
	if p == nil || p.processorPtr == nil || audios == nil || audios.audiosPtr == nil {
		return nil, errors.New("audio processor or audios are destroyed")
	}
	stringArray, err := createOgaStringArray(prompts)
	if err != nil {
		return nil, err
	}
	defer C.DestroyOgaStringArray(stringArray)
	var tensors *C.OgaNamedTensors
	if err := OgaResultToError(C.ProcessAudiosAndPrompts(p.processorPtr, stringArray, audios.audiosPtr, &tensors)); err != nil {
		return nil, fmt.Errorf("processing audio prompts: %w", err)
	}
	if tensors == nil {
		return nil, errors.New("audio prompt processing returned nil tensors without error")
	}
	return &NamedTensors{tensorsPtr: tensors}, nil
}

func (p *AudioProcessor) ProcessImagesAndAudios(prompt string, images *Images, audios *Audios) (*NamedTensors, error) {
	if p == nil || p.processorPtr == nil || audios == nil || audios.audiosPtr == nil {
		return nil, errors.New("audio processor or audios are destroyed")
	}
	var imagePtr *C.OgaImages
	if images != nil {
		imagePtr = images.imagesPtr
		if imagePtr == nil {
			return nil, errors.New("images are destroyed")
		}
	}
	cPrompt := C.CString(prompt)
	defer C.free(unsafe.Pointer(cPrompt))
	var tensors *C.OgaNamedTensors
	if err := OgaResultToError(C.ProcessImagesAndAudios(p.processorPtr, cPrompt, imagePtr, audios.audiosPtr, &tensors)); err != nil {
		return nil, fmt.Errorf("processing image and audio prompt: %w", err)
	}
	if tensors == nil {
		return nil, errors.New("image and audio processing returned nil tensors without error")
	}
	return &NamedTensors{tensorsPtr: tensors}, nil
}

func (p *AudioProcessor) ProcessImagesAndAudiosAndPrompts(prompts []string, images *Images, audios *Audios) (*NamedTensors, error) {
	if p == nil || p.processorPtr == nil || audios == nil || audios.audiosPtr == nil {
		return nil, errors.New("audio processor or audios are destroyed")
	}
	var imagePtr *C.OgaImages
	if images != nil {
		imagePtr = images.imagesPtr
		if imagePtr == nil {
			return nil, errors.New("images are destroyed")
		}
	}
	stringArray, err := createOgaStringArray(prompts)
	if err != nil {
		return nil, err
	}
	defer C.DestroyOgaStringArray(stringArray)
	var tensors *C.OgaNamedTensors
	if err := OgaResultToError(C.ProcessImagesAndAudiosAndPrompts(p.processorPtr, stringArray, imagePtr, audios.audiosPtr, &tensors)); err != nil {
		return nil, fmt.Errorf("processing image and audio prompts: %w", err)
	}
	if tensors == nil {
		return nil, errors.New("image and audio prompt processing returned nil tensors without error")
	}
	return &NamedTensors{tensorsPtr: tensors}, nil
}

// Destroy releases the processor and its model.
func (p *AudioProcessor) Destroy() {
	if p == nil {
		return
	}
	if p.processorPtr != nil {
		C.DestroyOgaMultiModalProcessor(p.processorPtr)
		p.processorPtr = nil
	}
	if p.modelPtr != nil {
		C.DestroyOgaModel(p.modelPtr)
		p.modelPtr = nil
	}
}

// CreateAdapters creates a manager bound to this Session's model.
func (s *Session) CreateAdapters() (*Adapters, error) {
	if s == nil || s.model == nil || s.model.modelPtr == nil {
		return nil, errors.New("session is destroyed")
	}
	var ptr *C.OgaAdapters
	if err := OgaResultToError(C.CreateAdapters(s.model.modelPtr, &ptr)); err != nil {
		return nil, fmt.Errorf("creating adapters: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("adapter creation returned nil without error")
	}
	return &Adapters{ptr: ptr, session: s}, nil
}

func (a *Adapters) Load(path, name string) error {
	if a == nil || a.ptr == nil {
		return errors.New("adapters are destroyed")
	}
	cPath, cName := C.CString(path), C.CString(name)
	defer C.free(unsafe.Pointer(cPath))
	defer C.free(unsafe.Pointer(cName))
	if err := OgaResultToError(C.LoadAdapter(a.ptr, cPath, cName)); err != nil {
		return fmt.Errorf("loading adapter: %w", err)
	}
	return nil
}

func (a *Adapters) Unload(name string) error {
	if a == nil || a.ptr == nil {
		return errors.New("adapters are destroyed")
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	if err := OgaResultToError(C.UnloadAdapter(a.ptr, cName)); err != nil {
		return fmt.Errorf("unloading adapter: %w", err)
	}
	return nil
}

func (s *Session) SetActiveAdapter(adapters *Adapters, name string) error {
	if s == nil {
		return errors.New("session is nil")
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.model == nil || s.model.modelPtr == nil {
		return errors.New("session is destroyed")
	}
	if adapters == nil || adapters.ptr == nil || adapters.session != s {
		return errors.New("adapters are destroyed or belong to another session")
	}
	if name == "" {
		return errors.New("adapter name is empty")
	}
	s.activeAdapters = adapters
	s.activeAdapterName = name
	return nil
}

func (a *Adapters) Destroy() {
	if a != nil && a.ptr != nil {
		if a.session != nil {
			a.session.mutex.Lock()
			if a.session.activeAdapters == a {
				a.session.activeAdapters = nil
				a.session.activeAdapterName = ""
			}
			a.session.mutex.Unlock()
		}
		C.DestroyAdapters(a.ptr)
		a.ptr = nil
		a.session = nil
	}
}

// Supports format: data:image/png;base64,<base64-encoded-data>.
func parseDataURI(dataURI string) ([]byte, error) {
	// Check if it starts with "data:"
	if !strings.HasPrefix(dataURI, "data:") {
		return nil, fmt.Errorf("invalid data URI: must start with 'data:'")
	}

	// Find the comma that separates metadata from data
	commaIdx := strings.Index(dataURI, ",")
	if commaIdx == -1 {
		return nil, fmt.Errorf("invalid data URI: missing comma separator")
	}

	// Extract metadata and check for base64
	metadata := dataURI[5:commaIdx] // skip "data:"
	if !strings.Contains(metadata, "base64") {
		return nil, fmt.Errorf("unsupported data URI encoding: only base64 is supported")
	}

	// Decode base64 data
	encodedData := dataURI[commaIdx+1:]
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 data: %w", err)
	}

	return decodedData, nil
}

// LoadImage loads a single image from a file path or data URI.
func LoadImage(imagePath string) (*Images, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}

	// Check if it's a data URI
	if strings.HasPrefix(imagePath, "data:") {
		imageData, err := parseDataURI(imagePath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse data URI: %w", err)
		}
		return LoadImageFromBuffer(imageData)
	}

	// Load from file path
	cPath := C.CString(imagePath)
	defer C.free(unsafe.Pointer(cPath))

	var cImages *C.OgaImages
	res := C.LoadOgaImage(cPath, &cImages)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("LoadImage failed: %w", err)
	}
	if cImages == nil {
		return nil, errors.New("LoadImage returned nil without error")
	}

	return &Images{imagesPtr: cImages}, nil
}

// LoadImageFromBuffer loads a single image from a byte buffer.
func LoadImageFromBuffer(imageData []byte) (*Images, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}

	if len(imageData) == 0 {
		return nil, errors.New("image data is empty")
	}

	// Pin the Go memory before passing to C
	var pinner runtime.Pinner
	pinner.Pin(&imageData[0])
	defer pinner.Unpin()

	// Create C array of pointers and sizes
	dataPtr := unsafe.Pointer(&imageData[0])
	dataSize := C.size_t(len(imageData))

	var cImages *C.OgaImages
	res := C.LoadOgaImagesFromBuffers(
		&dataPtr,
		&dataSize,
		1, // count = 1 for single image
		&cImages,
	)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("LoadImageFromBuffer failed: %w", err)
	}
	if cImages == nil {
		return nil, errors.New("LoadImageFromBuffer returned nil without error")
	}

	return &Images{imagesPtr: cImages}, nil
}

// LoadImages loads multiple images from file paths or data URIs.
func LoadImages(imagePaths []string) (*Images, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}

	if len(imagePaths) == 0 {
		return nil, errors.New("no image paths provided")
	}

	// Check if any paths are data URIs - if so, we need to use buffer loading
	hasDataURI := false
	for _, path := range imagePaths {
		if strings.HasPrefix(path, "data:") {
			hasDataURI = true
			break
		}
	}

	if hasDataURI {
		// Decode all images to buffers and use buffer loading
		buffers := make([][]byte, len(imagePaths))
		for i, path := range imagePaths {
			if strings.HasPrefix(path, "data:") {
				data, err := parseDataURI(path)
				if err != nil {
					return nil, fmt.Errorf("failed to parse data URI at index %d: %w", i, err)
				}
				buffers[i] = data
			} else {
				// For file paths, we'd need to read the file
				// For now, return an error if mixing data URIs with file paths
				return nil, errors.New("cannot mix data URIs with file paths in LoadImages")
			}
		}
		return LoadImagesFromBuffers(buffers)
	}

	// All are file paths - use the C API directly
	// Create OgaStringArray
	var cStringArray *C.OgaStringArray
	res := C.CreateOgaStringArray(&cStringArray)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("CreateOgaStringArray failed: %w", err)
	}
	defer C.DestroyOgaStringArray(cStringArray)

	// Add each path to the string array
	for _, path := range imagePaths {
		cPath := C.CString(path)
		res = C.AddStringToOgaStringArray(cStringArray, cPath)
		C.free(unsafe.Pointer(cPath))
		if err := OgaResultToError(res); err != nil {
			// cStringArray will be destroyed by defer C.DestroyOgaStringArray(cStringArray)
			return nil, fmt.Errorf("AddStringToOgaStringArray failed: %w", err)
		}
	}

	// Load images
	var cImages *C.OgaImages
	res = C.LoadOgaImages(cStringArray, &cImages)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("LoadImages failed: %w", err)
	}
	if cImages == nil {
		return nil, errors.New("LoadImages returned nil without error")
	}

	return &Images{imagesPtr: cImages}, nil
}

// LoadImagesFromBuffers loads multiple images from byte buffers.
func LoadImagesFromBuffers(imageBuffers [][]byte) (*Images, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}

	if len(imageBuffers) == 0 {
		return nil, errors.New("no image buffers provided")
	}

	// Create arrays for pointers and sizes
	dataPtrs := make([]unsafe.Pointer, len(imageBuffers))
	dataSizes := make([]C.size_t, len(imageBuffers))

	// Pin all buffer memory before passing to C
	var pinner runtime.Pinner
	defer pinner.Unpin()

	for i, buf := range imageBuffers {
		if len(buf) == 0 {
			return nil, fmt.Errorf("image buffer at index %d is empty", i)
		}
		pinner.Pin(&buf[0])
		dataPtrs[i] = unsafe.Pointer(&buf[0])
		dataSizes[i] = C.size_t(len(buf))
	}

	var cImages *C.OgaImages
	res := C.LoadOgaImagesFromBuffers(
		&dataPtrs[0],
		&dataSizes[0],
		C.size_t(len(imageBuffers)),
		&cImages,
	)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("LoadImagesFromBuffers failed: %w", err)
	}
	if cImages == nil {
		return nil, errors.New("LoadImagesFromBuffers returned nil without error")
	}

	return &Images{imagesPtr: cImages}, nil
}

// CreateMultiModalProcessor creates a multimodal processor from a model.
func createMultiModalProcessor(model *model) (*multiModalProcessor, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}

	if model == nil || model.modelPtr == nil {
		return nil, errors.New("model is nil")
	}

	var cProcessor *C.OgaMultiModalProcessor
	res := C.CreateOgaMultiModalProcessor(model.modelPtr, &cProcessor)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("CreateMultiModalProcessor failed: %w", err)
	}
	if cProcessor == nil {
		return nil, errors.New("CreateMultiModalProcessor returned nil without error")
	}

	return &multiModalProcessor{processorPtr: cProcessor}, nil
}

// ProcessImages processes images with a prompt and returns named tensors.
func (p *multiModalProcessor) ProcessImages(prompt string, images *Images) (*NamedTensors, error) {
	if p.processorPtr == nil {
		return nil, errors.New("processor is not initialized")
	}
	if images == nil || images.imagesPtr == nil {
		return nil, errors.New("images is nil")
	}

	// TODO: support multiple prompts, somehow this gives an error on image tag mismatch
	// // Create OgaStringArray for prompts
	// var cStringArray *C.OgaStringArray
	// res := C.CreateOgaStringArray(&cStringArray)
	// if err := OgaResultToError(res); err != nil {
	// 	return nil, fmt.Errorf("CreateOgaStringArray failed: %w", err)
	// }
	// defer C.DestroyOgaStringArray(cStringArray)

	// for _, prompt := range prompts {
	// 	cPrompt := C.CString(prompt)
	// 	res = C.AddStringToOgaStringArray(cStringArray, cPrompt)
	// 	C.free(unsafe.Pointer(cPrompt))
	// 	if err := OgaResultToError(res); err != nil {
	// 		return nil, fmt.Errorf("AddStringToOgaStringArray failed: %w", err)
	// 	}
	// }
	// res = C.ProcessOgaImagesAndPrompts(p.processorPtr, cStringArray, images.imagesPtr, &cTensors)

	var cTensors *C.OgaNamedTensors
	promptC := C.CString(prompt)
	defer C.free(unsafe.Pointer(promptC))
	res := C.ProcessOgaImages(p.processorPtr, promptC, images.imagesPtr, &cTensors)
	if err := OgaResultToError(res); err != nil {
		if cTensors != nil {
			C.DestroyOgaNamedTensors(cTensors)
		}
		return nil, fmt.Errorf("ProcessImages failed: %w", err)
	}
	if cTensors == nil {
		return nil, errors.New("ProcessImages returned nil without error")
	}
	return &NamedTensors{tensorsPtr: cTensors}, nil
}

// ProcessAudios converts audio and a prompt to copied-owner named tensors
// using the session's multimodal processor.
func (p *multiModalProcessor) ProcessAudios(prompt string, audios *Audios) (*NamedTensors, error) {
	if p == nil || p.processorPtr == nil {
		return nil, errors.New("processor is not initialized")
	}
	if audios == nil || audios.audiosPtr == nil {
		return nil, errors.New("audios is nil")
	}
	var cTensors *C.OgaNamedTensors
	promptC := C.CString(prompt)
	defer C.free(unsafe.Pointer(promptC))
	if err := OgaResultToError(C.ProcessAudios(p.processorPtr, promptC, audios.audiosPtr, &cTensors)); err != nil {
		if cTensors != nil {
			C.DestroyOgaNamedTensors(cTensors)
		}
		return nil, fmt.Errorf("ProcessAudios failed: %w", err)
	}
	if cTensors == nil {
		return nil, errors.New("ProcessAudios returned nil without error")
	}
	return &NamedTensors{tensorsPtr: cTensors}, nil
}

// ProcessImagesAndAudios converts images, audio, and a prompt to copied-owner
// named tensors using the session's multimodal processor.
func (p *multiModalProcessor) ProcessImagesAndAudios(prompt string, images *Images, audios *Audios) (*NamedTensors, error) {
	if p == nil || p.processorPtr == nil {
		return nil, errors.New("processor is not initialized")
	}
	if images == nil || images.imagesPtr == nil {
		return nil, errors.New("images is nil")
	}
	if audios == nil || audios.audiosPtr == nil {
		return nil, errors.New("audios is nil")
	}
	var cTensors *C.OgaNamedTensors
	promptC := C.CString(prompt)
	defer C.free(unsafe.Pointer(promptC))
	if err := OgaResultToError(C.ProcessImagesAndAudios(p.processorPtr, promptC, images.imagesPtr, audios.audiosPtr, &cTensors)); err != nil {
		if cTensors != nil {
			C.DestroyOgaNamedTensors(cTensors)
		}
		return nil, fmt.Errorf("ProcessImagesAndAudios failed: %w", err)
	}
	if cTensors == nil {
		return nil, errors.New("ProcessImagesAndAudios returned nil without error")
	}
	return &NamedTensors{tensorsPtr: cTensors}, nil
}
