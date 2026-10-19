package ortgenai

/*
#cgo CFLAGS: -O2 -g
#include "ort_genai_wrapper.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"sort"
	"unsafe"
)

// Tokenizer owns an independent model handle and tokenizer.
type Tokenizer struct {
	modelPtr     *C.OgaModel
	tokenizerPtr *C.OgaTokenizer
}

// TokenizerStream incrementally decodes generated token IDs.
type TokenizerStream struct {
	ptr       *C.OgaTokenizerStream
	tokenizer *Tokenizer
}

// CreateTokenizer creates a tokenizer from a model directory.
func CreateTokenizer(modelPath string) (*Tokenizer, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if modelPath == "" {
		return nil, errors.New("model path is empty")
	}
	path := C.CString(modelPath)
	defer C.free(unsafe.Pointer(path))
	var model *C.OgaModel
	if err := OgaResultToError(C.CreateOgaModel(path, &model)); err != nil {
		return nil, fmt.Errorf("creating tokenizer model: %w", err)
	}
	if model == nil {
		return nil, errors.New("tokenizer model creation returned nil without error")
	}
	var tokenizer *C.OgaTokenizer
	if err := OgaResultToError(C.CreateOgaTokenizer(model, &tokenizer)); err != nil {
		C.DestroyOgaModel(model)
		return nil, fmt.Errorf("creating tokenizer: %w", err)
	}
	if tokenizer == nil {
		C.DestroyOgaModel(model)
		return nil, errors.New("tokenizer creation returned nil without error")
	}
	return &Tokenizer{modelPtr: model, tokenizerPtr: tokenizer}, nil
}

// Encode returns a copied token-ID sequence.
func (t *Tokenizer) Encode(text string) ([]int32, error) {
	if t == nil || t.tokenizerPtr == nil {
		return nil, errors.New("tokenizer is destroyed")
	}
	input := C.CString(text)
	defer C.free(unsafe.Pointer(input))
	var sequences *C.OgaSequences
	if err := OgaResultToError(C.CreateOgaSequences(&sequences)); err != nil {
		return nil, fmt.Errorf("creating token sequence: %w", err)
	}
	if sequences == nil {
		return nil, errors.New("token sequence creation returned nil without error")
	}
	defer C.DestroyOgaSequences(sequences)
	if err := OgaResultToError(C.TokenizerEncode(t.tokenizerPtr, input, sequences)); err != nil {
		return nil, fmt.Errorf("encoding text: %w", err)
	}
	if C.SequencesCount(sequences) == 0 {
		return []int32{}, nil
	}
	count := int(C.SequencesGetSequenceCount(sequences, 0))
	data := C.SequencesGetSequenceData(sequences, 0)
	if count == 0 {
		return []int32{}, nil
	}
	if data == nil {
		return nil, errors.New("tokenizer returned nil sequence data")
	}
	values := unsafe.Slice((*int32)(unsafe.Pointer(data)), count)
	return append([]int32(nil), values...), nil
}

// Decode converts token IDs to a copied UTF-8 string.
func (t *Tokenizer) Decode(ids []int32) (string, error) {
	if t == nil || t.tokenizerPtr == nil {
		return "", errors.New("tokenizer is destroyed")
	}
	var data *C.int32_t
	if len(ids) != 0 {
		data = (*C.int32_t)(unsafe.Pointer(&ids[0]))
	}
	var decoded *C.char
	if err := OgaResultToError(C.TokenizerDecode(t.tokenizerPtr, data, C.size_t(len(ids)), (**C.char)(unsafe.Pointer(&decoded)))); err != nil {
		return "", fmt.Errorf("decoding token IDs: %w", err)
	}
	if decoded == nil {
		return "", nil
	}
	text := C.GoString(decoded)
	C.DestroyOgaString(decoded)
	return text, nil
}

// TokenID converts a token spelling to its integer ID.
func (t *Tokenizer) TokenID(token string) (int32, error) {
	if t == nil || t.tokenizerPtr == nil {
		return 0, errors.New("tokenizer is destroyed")
	}
	value := C.CString(token)
	defer C.free(unsafe.Pointer(value))
	var id C.int32_t
	if err := OgaResultToError(C.TokenizerToTokenId(t.tokenizerPtr, value, &id)); err != nil {
		return 0, fmt.Errorf("converting token to ID: %w", err)
	}
	return int32(id), nil
}

// EncodeBatch returns token IDs for a batch as an owned native tensor.
func (t *Tokenizer) EncodeBatch(texts []string) (*Tensor, error) {
	if t == nil || t.tokenizerPtr == nil {
		return nil, errors.New("tokenizer is destroyed")
	}
	if len(texts) == 0 {
		return nil, errors.New("tokenizer batch is empty")
	}
	values := make([]*C.char, len(texts))
	for i, text := range texts {
		values[i] = C.CString(text)
	}
	defer func() {
		for _, value := range values {
			C.free(unsafe.Pointer(value))
		}
	}()
	var tensor *C.OgaTensor
	if err := OgaResultToError(C.TokenizerEncodeBatch(t.tokenizerPtr, (**C.char)(unsafe.Pointer(&values[0])), C.size_t(len(values)), &tensor)); err != nil {
		return nil, fmt.Errorf("batch encoding text: %w", err)
	}
	if tensor == nil {
		return nil, errors.New("batch tokenizer returned nil tensor without error")
	}
	return &Tensor{ptr: tensor}, nil
}

// DecodeBatch converts a token tensor to copied strings.
func (t *Tokenizer) DecodeBatch(tokens *Tensor) ([]string, error) {
	if t == nil || t.tokenizerPtr == nil {
		return nil, errors.New("tokenizer is destroyed")
	}
	if err := tokens.check(); err != nil {
		return nil, err
	}
	var values *C.OgaStringArray
	if err := OgaResultToError(C.TokenizerDecodeBatch(t.tokenizerPtr, tokens.ptr, &values)); err != nil {
		return nil, fmt.Errorf("batch decoding token IDs: %w", err)
	}
	if values == nil {
		return nil, errors.New("batch tokenizer returned nil strings without error")
	}
	defer C.DestroyOgaStringArray(values)
	var count C.size_t
	if err := OgaResultToError(C.StringArrayGetCount(values, &count)); err != nil {
		return nil, fmt.Errorf("reading decoded string count: %w", err)
	}
	decoded := make([]string, int(count))
	for i := range decoded {
		var value *C.char
		if err := OgaResultToError(C.StringArrayGetString(values, C.size_t(i), (**C.char)(unsafe.Pointer(&value)))); err != nil {
			return nil, fmt.Errorf("reading decoded string %d: %w", i, err)
		}
		if value != nil {
			decoded[i] = C.GoString(value)
		}
	}
	return decoded, nil
}

// UpdateOptions sets per-tokenizer string options, such as add_special_tokens or skip_special_tokens.
func (t *Tokenizer) UpdateOptions(options map[string]string) error {
	if t == nil || t.tokenizerPtr == nil {
		return errors.New("tokenizer is destroyed")
	}
	keys := make([]string, 0, len(options))
	for key := range options {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	keyPointers := make([]*C.char, len(keys))
	valuePointers := make([]*C.char, len(keys))
	defer func() {
		for i := range keys {
			C.free(unsafe.Pointer(keyPointers[i]))
			C.free(unsafe.Pointer(valuePointers[i]))
		}
	}()
	for i, key := range keys {
		keyPointers[i] = C.CString(key)
		valuePointers[i] = C.CString(options[key])
	}
	var keyPtr, valuePtr **C.char
	if len(keys) != 0 {
		keyPtr = (**C.char)(unsafe.Pointer(&keyPointers[0]))
		valuePtr = (**C.char)(unsafe.Pointer(&valuePointers[0]))
	}
	if err := OgaResultToError(C.TokenizerUpdateOptions(t.tokenizerPtr, (**C.char)(unsafe.Pointer(keyPtr)), (**C.char)(unsafe.Pointer(valuePtr)), C.size_t(len(keys)))); err != nil {
		return fmt.Errorf("updating tokenizer options: %w", err)
	}
	return nil
}

// ApplyChatTemplate applies the tokenizer's configured chat template and returns a copied string.
func (t *Tokenizer) ApplyChatTemplate(template, messages, tools string, addGenerationPrompt bool) (string, error) {
	if t == nil || t.tokenizerPtr == nil {
		return "", errors.New("tokenizer is destroyed")
	}
	cTemplate, cMessages, cTools := C.CString(template), C.CString(messages), C.CString(tools)
	defer C.free(unsafe.Pointer(cTemplate))
	defer C.free(unsafe.Pointer(cMessages))
	defer C.free(unsafe.Pointer(cTools))
	var output *C.char
	if err := OgaResultToError(C.ApplyOgaTokenizerChatTemplate(t.tokenizerPtr, cTemplate, cMessages, cTools, C.bool(addGenerationPrompt), (**C.char)(unsafe.Pointer(&output)))); err != nil {
		return "", fmt.Errorf("applying chat template: %w", err)
	}
	if output == nil {
		return "", nil
	}
	text := C.GoString(output)
	C.DestroyOgaString(output)
	return text, nil
}

// NewStream creates an incremental decoder bound to this tokenizer.
func (t *Tokenizer) NewStream() (*TokenizerStream, error) {
	if t == nil || t.tokenizerPtr == nil {
		return nil, errors.New("tokenizer is destroyed")
	}
	var ptr *C.OgaTokenizerStream
	if err := OgaResultToError(C.CreateOgaTokenizerStream(t.tokenizerPtr, &ptr)); err != nil {
		return nil, fmt.Errorf("creating tokenizer stream: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("tokenizer stream creation returned nil without error")
	}
	return &TokenizerStream{ptr: ptr, tokenizer: t}, nil
}

// Decode converts one token to its next copied stream fragment.
func (s *TokenizerStream) Decode(token int32) (string, error) {
	if s == nil || s.ptr == nil || s.tokenizer == nil || s.tokenizer.tokenizerPtr == nil {
		return "", errors.New("tokenizer stream or its tokenizer is destroyed")
	}
	var output *C.char
	if err := OgaResultToError(C.TokenizerStreamDecode(s.ptr, C.int32_t(token), (**C.char)(unsafe.Pointer(&output)))); err != nil {
		return "", fmt.Errorf("decoding token stream: %w", err)
	}
	if output == nil {
		return "", nil
	}
	return C.GoString(output), nil
}

// Destroy releases the native incremental decoder.
func (s *TokenizerStream) Destroy() {
	if s != nil && s.ptr != nil {
		C.DestroyOgaTokenizerStream(s.ptr)
		s.ptr = nil
		s.tokenizer = nil
	}
}

// Destroy releases the tokenizer and its retained model.
func (t *Tokenizer) Destroy() {
	if t == nil {
		return
	}
	if t.tokenizerPtr != nil {
		C.DestroyOgaTokenizer(t.tokenizerPtr)
		t.tokenizerPtr = nil
	}
	if t.modelPtr != nil {
		C.DestroyOgaModel(t.modelPtr)
		t.modelPtr = nil
	}
}
