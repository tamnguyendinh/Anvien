package ortgenai

/*
#cgo CFLAGS: -O2 -g
#include "ort_genai_wrapper.h"
*/
import "C"

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var ErrNotInitialized = fmt.Errorf("InitializeEnvironment() has either " +
	"not yet been called, or did not return successfully")

var onnxGenaiSharedLibraryPath string

var genAiEnv bool

func IsInitialized() bool {
	return genAiEnv
}

func InitializeEnvironment() error {
	if IsInitialized() {
		return fmt.Errorf("GenAI environment already initialized")
	}
	if err := InitializeGenAiLibrary(); err != nil {
		return fmt.Errorf("error initializing GenAI library: %w", err)
	}
	genAiEnv = true
	return nil
}

// SetTelemetryEnabled enables or disables telemetry collection in ORT GenAI.
func SetTelemetryEnabled(enabled bool) {
	C.OgaSetTelemetryEnabled(C.bool(enabled))
}

// DestroyEnvironment Call this function to clean up the internal onnxruntime environment when it
// is no longer required.
func DestroyEnvironment() error {
	if !IsInitialized() {
		return ErrNotInitialized
	}
	if err := platformCleanup(); err != nil {
		return fmt.Errorf("error during platform cleanup: %w", err)
	}
	genAiEnv = false
	return nil
}

func SetSharedLibraryPath(path string) {
	onnxGenaiSharedLibraryPath = path
}

// GuidanceType specifies the constrained-generation strategy passed to OgaGeneratorParamsSetGuidance.
type GuidanceType string

const (
	GuidanceTypeJSONSchema  GuidanceType = "json_schema"
	GuidanceTypeRegex       GuidanceType = "regex"
	GuidanceTypeLarkGrammar GuidanceType = "lark_grammar"
)

// Guidance configures constrained (guided) generation. Requires a recent OGA runtime.
type Guidance struct {
	Type GuidanceType
	Data string
	// EnableFFTokens speeds up generation by force-forwarding tokens that satisfy the grammar
	// without calling the model. Only valid when BatchSize=1 and beam_size=1.
	EnableFFTokens bool
}

type GenerationOptions struct {
	// MaxLength is the maximum sequence length for this generation call.
	// Values less than or equal to zero use the package default.
	MaxLength   int
	BatchSize   int
	Temperature *float64
	TopP        *float64
	Seed        *int
	Guidance    *Guidance
	// EnableThinking controls the chat template's enable_thinking argument.
	// A nil value leaves the model's default behavior unchanged.
	EnableThinking *bool
}

var defaultMaxLength = 2024

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type generator struct {
	generatorParamsPtr *C.OgaGeneratorParams
	generatorPtr       *C.OgaGenerator
}

func (g *generator) destroy() {
	if g.generatorPtr != nil {
		C.DestroyOgaGenerator(g.generatorPtr)
		g.generatorPtr = nil
	}
	if g.generatorParamsPtr != nil {
		C.DestroyOgaGeneratorParams(g.generatorParamsPtr)
		g.generatorParamsPtr = nil
	}
}

func (g *generator) setInputs(namedTensors *NamedTensors) error {
	if namedTensors == nil || namedTensors.tensorsPtr == nil {
		return errors.New("named tensors is nil")
	}
	res := C.GeneratorSetInputs(g.generatorPtr, namedTensors.tensorsPtr)
	if err := OgaResultToError(res); err != nil {
		return fmt.Errorf("GeneratorSetInputs failed: %w", err)
	}
	return nil
}

func (g *generator) addSequences(sequences *sequences) error {
	// add sequences to generator
	res := C.GeneratorAppendTokenSequences(g.generatorPtr, sequences.sequencesPtr)
	if err := OgaResultToError(res); err != nil {
		return fmt.Errorf("GeneratorAppendTokenSequences failed: %w", err)
	}
	// Sequences are no longer needed after appending; destroy to avoid leaks.
	C.DestroyOgaSequences(sequences.sequencesPtr)
	sequences.sequencesPtr = nil
	return nil
}

type tokenizer struct {
	tokenizerPtr *C.OgaTokenizer
	EOSTokenIDs  []int
	PadTokenID   int
	BotTokenID   int
	EotTokenID   int
	BorTokenID   int
	EorTokenID   int
}

func newTokenizerFromModel(model model) (tokenizer, error) {
	var cTokenizer *C.OgaTokenizer
	res := C.CreateOgaTokenizer(model.modelPtr, &cTokenizer)
	if err := OgaResultToError(res); err != nil {
		return tokenizer{}, fmt.Errorf("CreateOgaTokenizer failed: %w", err)
	}
	if cTokenizer == nil {
		return tokenizer{}, errors.New("CreateOgaTokenizer returned nil without error")
	}
	var cIDs *C.int32_t
	var cCount C.size_t

	res = C.OgaTokenizerGetEosTokenIds(cTokenizer, (**C.int32_t)(unsafe.Pointer(&cIDs)), &cCount)
	if err := OgaResultToError(res); err != nil {
		return tokenizer{}, fmt.Errorf("OgaTokenizerGetEosTokenIds failed: %w", err)
	}
	if cIDs == nil {
		return tokenizer{}, errors.New("OgaTokenizerGetEosTokenIds returned nil without error")
	}
	n := int(cCount)
	out := make([]int, n)
	arr := (*[1 << 30]C.int32_t)(unsafe.Pointer(cIDs))
	for i := range n {
		out[i] = int(arr[i])
	}

	var padID, botID, eotID, borID, eorID C.int32_t
	C.OgaTokenizerGetPadTokenId(cTokenizer, &padID)
	C.OgaTokenizerGetBotTokenId(cTokenizer, &botID)
	C.OgaTokenizerGetEotTokenId(cTokenizer, &eotID)
	C.OgaTokenizerGetBorTokenId(cTokenizer, &borID)
	C.OgaTokenizerGetEorTokenId(cTokenizer, &eorID)

	return tokenizer{
		tokenizerPtr: cTokenizer,
		EOSTokenIDs:  out,
		PadTokenID:   int(padID),
		BotTokenID:   int(botID),
		EotTokenID:   int(eotID),
		BorTokenID:   int(borID),
		EorTokenID:   int(eorID),
	}, nil
}

func (t *tokenizer) encode(prompt string, sequences *sequences) error {
	cStr := C.CString(prompt)
	defer C.free(unsafe.Pointer(cStr))
	result := C.TokenizerEncode(t.tokenizerPtr, cStr, sequences.sequencesPtr)
	if err := OgaResultToError(result); err != nil {
		return fmt.Errorf("TokenizerEncode failed: %w", err)
	}
	return nil
}

func (t *tokenizer) destroy() {
	C.DestroyOgaTokenizer(t.tokenizerPtr)
	t.tokenizerPtr = nil
}

type tokenizerStream struct {
	streamPtr *C.OgaTokenizerStream
}

func (ts *tokenizerStream) Decode(token C.int32_t) (string, error) {
	var cOutput *C.char
	result := C.TokenizerStreamDecode(ts.streamPtr, token, &cOutput)
	if err := OgaResultToError(result); err != nil {
		return "", fmt.Errorf("TokenizerStreamDecode failed: %w", err)
	}
	if cOutput == nil {
		return "", nil
	}
	decoded := C.GoString(cOutput) // don't free this one -- owned by the tokenizer stream
	return decoded, nil
}

func (ts *tokenizerStream) destroy() {
	C.DestroyOgaTokenizerStream(ts.streamPtr)
	ts.streamPtr = nil
}

type sequences struct {
	sequencesPtr *C.OgaSequences
}

func (s *sequences) destroy() {
	if s.sequencesPtr != nil {
		C.DestroyOgaSequences(s.sequencesPtr)
		s.sequencesPtr = nil
	}
}

type model struct {
	modelPtr *C.OgaModel
}

func (m *model) destroy() {
	if m.modelPtr != nil {
		C.DestroyOgaModel(m.modelPtr)
		m.modelPtr = nil
	}
}

type Session struct {
	modelPath          string // path used to load the model (retained for MTP generation)
	model              *model
	processor          *multiModalProcessor
	tokenizer          *tokenizer
	thinkingTokenizers map[bool]*tokenizer
	statistics         *Statistics
	activeAdapters     *Adapters
	activeAdapterName  string
	mutex              sync.Mutex // the C API is not thread-safe
	nativeBuffers      []unsafe.Pointer
}

// tokenizerForGeneration selects a session-owned tokenizer configured for the
// requested chat-template thinking behavior. Callers must hold s.mutex because
// tokenizer creation and option updates use the native API.
func (s *Session) tokenizerForGeneration(enableThinking *bool) (*tokenizer, error) {
	if s.tokenizer == nil || s.tokenizer.tokenizerPtr == nil {
		return nil, errors.New("session is destroyed")
	}
	if enableThinking == nil {
		return s.tokenizer, nil
	}
	if s.thinkingTokenizers != nil {
		if t := s.thinkingTokenizers[*enableThinking]; t != nil {
			return t, nil
		}
	}
	if s.model == nil || s.model.modelPtr == nil {
		return nil, errors.New("session is destroyed")
	}

	t, err := newTokenizerFromModel(*s.model)
	if err != nil {
		return nil, fmt.Errorf("creating tokenizer for thinking option: %w", err)
	}
	kwargs, err := json.Marshal(map[string]bool{"enable_thinking": *enableThinking})
	if err != nil {
		t.destroy()
		return nil, fmt.Errorf("marshaling chat template options: %w", err)
	}
	// UpdateOptions is exposed on the standalone Tokenizer wrapper. This
	// temporary wrapper only borrows t's native handle; it does not own it.
	borrowed := &Tokenizer{tokenizerPtr: t.tokenizerPtr}
	if err := borrowed.UpdateOptions(map[string]string{"chat_template_kwargs": string(kwargs)}); err != nil {
		t.destroy()
		return nil, fmt.Errorf("setting chat template options: %w", err)
	}
	if s.thinkingTokenizers == nil {
		s.thinkingTokenizers = make(map[bool]*tokenizer, 2)
	}
	s.thinkingTokenizers[*enableThinking] = &t
	return &t, nil
}

type MaxLengthReachedError struct{}

func (e MaxLengthReachedError) Error() string {
	return "generation stopped: max length reached"
}

type SequenceDelta struct {
	Sequence   int
	Token      string
	EOSReached bool
}

// Statistics captures generation performance metrics.
type Statistics struct {
	AvgPrefillSeconds float64
	TokensPerSecond   float64
	// cumulative
	CumulativePrefillSum           float64
	CumulativePrefillCount         int
	CumulativeTokens               int
	CumulativeTokenDurationSeconds float64
}

// GetStatistics returns the last computed statistics for the session.
func (s *Session) GetStatistics() *Statistics {
	return s.statistics
}

func (g *generator) IsDone() bool {
	return bool(C.IsDone(g.generatorPtr))
}

func (g *generator) generateNextToken() error {
	if g == nil || g.generatorPtr == nil {
		return errors.New("generator is destroyed")
	}
	if err := OgaResultToError(C.GeneratorGenerateNextToken(g.generatorPtr)); err != nil {
		return fmt.Errorf("generating next token: %w", err)
	}
	return nil
}

func (t *tokenizer) ApplyChatTemplate(inputMessages []byte, tools []byte, addGenerationPrompt bool) (string, error) {
	if t.tokenizerPtr == nil {
		return "", errors.New("tokenizer is not initialized")
	}
	cInput := C.CString(string(inputMessages))
	defer C.free(unsafe.Pointer(cInput))
	cTools := C.CString(string(tools))
	defer C.free(unsafe.Pointer(cTools))
	var cOutput *C.char
	res := C.ApplyOgaTokenizerChatTemplate(t.tokenizerPtr, nil, cInput, cTools, C.bool(addGenerationPrompt), &cOutput)
	if err := OgaResultToError(res); err != nil {
		return "", fmt.Errorf("ApplyOgaChatTemplate failed: %w", err)
	}
	if cOutput == nil {
		return "", errors.New("ApplyOgaChatTemplate returned nil output without error")
	}
	output := C.GoString(cOutput)
	C.DestroyOgaString(cOutput)
	return output, nil
}

func (t *tokenizer) createTokenizerStream() (*tokenizerStream, error) {
	var cStream *C.OgaTokenizerStream
	res := C.CreateOgaTokenizerStream(t.tokenizerPtr, &cStream)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("CreateOgaTokenizerStream failed: %w", err)
	}
	if cStream == nil {
		return nil, errors.New("CreateOgaTokenizerStream returned nil without error")
	}
	return &tokenizerStream{streamPtr: cStream}, nil
}

func (t *tokenizer) tokenizeMessages(messages [][]Message, tools []string) (*sequences, []*tokenizerStream, error) {
	if t.tokenizerPtr == nil {
		return nil, nil, errors.New("tokenizer is not initialized")
	}
	if len(messages) == 0 {
		return nil, nil, errors.New("no messages provided")
	}

	rawTools := make([]json.RawMessage, len(tools))
	for i, t := range tools {
		rawTools[i] = json.RawMessage(t)
	}
	toolsJSON, err := json.Marshal(rawTools)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal tools: %w", err)
	}

	var cSequences *C.OgaSequences
	res := C.CreateOgaSequences(&cSequences)
	if err := OgaResultToError(res); err != nil {
		return nil, nil, fmt.Errorf("CreateOgaSequences failed: %w", err)
	}
	if cSequences == nil {
		return nil, nil, errors.New("CreateOgaSequences returned nil without error")
	}
	sequencesInstance := &sequences{sequencesPtr: cSequences}
	tokenizerStreams := make([]*tokenizerStream, 0, len(messages))

	for _, message := range messages {
		messageJSON, err := json.Marshal(message)
		if err != nil {
			tokenizeCleanup(sequencesInstance, tokenizerStreams)
			return nil, nil, fmt.Errorf("failed to marshal input message: %w", err)
		}
		prompt, templateErr := t.ApplyChatTemplate(messageJSON, toolsJSON, true)
		if templateErr != nil {
			tokenizeCleanup(sequencesInstance, tokenizerStreams)
			return nil, nil, fmt.Errorf("failed to apply chat template: %w", templateErr)
		}
		if err = t.encode(prompt, sequencesInstance); err != nil {
			tokenizeCleanup(sequencesInstance, tokenizerStreams)
			return nil, nil, fmt.Errorf("encode failed: %w", err)
		}
		stream, err := t.createTokenizerStream()
		if err != nil {
			tokenizeCleanup(sequencesInstance, tokenizerStreams)
			return nil, nil, fmt.Errorf("createTokenizerStream failed: %w", err)
		}
		tokenizerStreams = append(tokenizerStreams, stream)
	}
	return sequencesInstance, tokenizerStreams, nil
}

func tokenizeCleanup(sequencesInstance *sequences, tokenizerStreams []*tokenizerStream) {
	sequencesInstance.destroy()
	for _, ts := range tokenizerStreams {
		if ts != nil {
			ts.destroy()
		}
	}
}

func (s *Session) createGenerator(generationOptions *GenerationOptions) (*generator, error) {
	cGeneratorParams, err := createGeneratorParams(s.model, generationOptions)
	if err != nil {
		return nil, err
	}

	// create a generator with those params
	var cGenerator *C.OgaGenerator
	res := C.CreateOgaGenerator(s.model.modelPtr, cGeneratorParams, &cGenerator)
	if err := OgaResultToError(res); err != nil {
		C.DestroyOgaGeneratorParams(cGeneratorParams)
		return nil, fmt.Errorf("CreateOgaGenerator failed: %w", err)
	}
	if cGenerator == nil {
		C.DestroyOgaGeneratorParams(cGeneratorParams)
		return nil, errors.New("CreateOgaGenerator returned nil generator without error")
	}
	if s.activeAdapters != nil {
		cName := C.CString(s.activeAdapterName)
		res = C.SetActiveAdapter(cGenerator, s.activeAdapters.ptr, cName)
		C.free(unsafe.Pointer(cName))
		if err := OgaResultToError(res); err != nil {
			C.DestroyOgaGenerator(cGenerator)
			C.DestroyOgaGeneratorParams(cGeneratorParams)
			return nil, fmt.Errorf("setting active adapter: %w", err)
		}
	}

	return &generator{
		generatorParamsPtr: cGeneratorParams,
		generatorPtr:       cGenerator,
	}, nil
}

func sendGenerationError(errChan chan<- error, err error) {
	select {
	case errChan <- err:
	default:
	}
}

// runGenerationLoop drives the native generator until it is done or the context is
// cancelled, decoding tokens and emitting SequenceDelta values on outChan. Each
// sequence of this run is relabelled by adding seqOffset so a single-conversation
// run can be positioned within a larger batch. Generation errors are relayed on
// errChan and also returned. Assumes the caller already holds s.mutex and the
// generator's sequences have been appended; it does not destroy the generator.
func (s *Session) runGenerationLoop(ctx context.Context, generator *generator, seqCount, seqOffset int, tokenizerStreams []*tokenizerStream, maxLength int, outChan chan<- SequenceDelta, errChan chan<- error) error {
	// Per-run statistics (goroutine-local to avoid races)
	runStart := time.Now()
	runFirstTokenTimes := map[int]time.Time{}
	runTokenCount := 0

	// finalize tokens/sec at the end of the run
	defer func() {
		var earliest time.Time
		for _, ft := range runFirstTokenTimes {
			if !ft.IsZero() && (earliest.IsZero() || ft.Before(earliest)) {
				earliest = ft
			}
		}
		if !earliest.IsZero() && runTokenCount > 0 {
			dur := time.Since(earliest).Seconds()
			if dur > 0 {
				s.statistics.CumulativeTokenDurationSeconds += dur
				s.statistics.TokensPerSecond = float64(s.statistics.CumulativeTokens) / s.statistics.CumulativeTokenDurationSeconds
			}
		}
	}()

	firstEmitted := make([]bool, seqCount)
	lastChar := make([]rune, seqCount)

	// Capture initial token counts per sequence to compare against at completion
	initialCounts := make([]int, seqCount)
	for i := range seqCount {
		initialCounts[i] = int(C.GeneratorGetSequenceCount(generator.generatorPtr, C.size_t(i)))
	}

	// prevCounts tracks how many tokens each sequence had before the last GenerateNextToken
	// call. When ff-tokens are enabled the call may advance a sequence by more than one
	// token; we must process every new token in order so the tokenizer stream stays in sync.
	prevCounts := make([]int, seqCount)
	copy(prevCounts, initialCounts)

	// Iterate over each sequence in the batch
	completeSequences := map[int]bool{}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if generator.IsDone() {
			// Determine if max generated tokens (excluding prompt) reached in any sequence
			reached := false
			for i := range seqCount {
				if initialCounts[i]+runTokenCount >= maxLength {
					reached = true
					break
				}
			}
			if reached {
				sendGenerationError(errChan, MaxLengthReachedError{})
			}
			return nil
		}

		result := C.GeneratorGenerateNextToken(generator.generatorPtr)
		if err := OgaResultToError(result); err != nil {
			sendGenerationError(errChan, err)
			return err
		}

		for i := range seqCount {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			if completeSequences[i] {
				continue
			}

			newCount := int(C.GeneratorGetSequenceCount(generator.generatorPtr, C.size_t(i)))
			if newCount <= prevCounts[i] {
				continue
			}
			seqData := C.GeneratorGetSequenceData(generator.generatorPtr, C.size_t(i))
			if seqData == nil {
				continue
			}
			arr := (*[1 << 30]C.int32_t)(unsafe.Pointer(seqData))

			for j := prevCounts[i]; j < newCount; j++ {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
				}

				token := arr[j]

				if slices.Contains(s.tokenizer.EOSTokenIDs, int(token)) {
					completeSequences[i] = true
					prevCounts[i] = newCount
					outChan <- SequenceDelta{Sequence: seqOffset + i, EOSReached: true}
					break
				}

				decoded, decodeErr := tokenizerStreams[i].Decode(token)
				if decodeErr != nil {
					sendGenerationError(errChan, decodeErr)
					return decodeErr
				}
				if decoded == "" {
					continue
				}
				// stats
				if runFirstTokenTimes[i].IsZero() {
					runFirstTokenTimes[i] = time.Now()
					prefill := runFirstTokenTimes[i].Sub(runStart).Seconds()
					s.statistics.CumulativePrefillSum += prefill
					s.statistics.CumulativePrefillCount++
					s.statistics.AvgPrefillSeconds = s.statistics.CumulativePrefillSum / float64(s.statistics.CumulativePrefillCount)
				}
				s.statistics.CumulativeTokens++
				runTokenCount++
				// normalization: skip leading spaces for first token, avoid repeated '.' at end
				if !firstEmitted[i] {
					trim := strings.TrimLeft(decoded, " ")
					if trim == "" {
						continue
					}
					decoded = trim
					firstEmitted[i] = true
				}
				if decoded == "." && lastChar[i] == '.' {
					continue
				}
				r := []rune(decoded)
				lastChar[i] = r[len(r)-1]

				select {
				case outChan <- SequenceDelta{Sequence: seqOffset + i, Token: decoded}:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			prevCounts[i] = newCount
		}
	}
}

// startGenerationGoroutine launches the unified generation loop and returns output and error channels.
// Assumes the session mutex is already locked; it will be unlocked inside the goroutine when done.
func (s *Session) startGenerationGoroutine(ctx context.Context, generator *generator, sequences *sequences, tensors *NamedTensors, tokenizerStreams []*tokenizerStream, seqCount int, maxLength int) (<-chan SequenceDelta, <-chan error) {
	outputChan := make(chan SequenceDelta, 1000)
	errChan := make(chan error, 1)
	go func() {
		defer close(outputChan)
		defer close(errChan)
		defer generator.destroy()
		if sequences != nil {
			defer sequences.destroy()
		}
		if tensors != nil {
			defer tensors.destroy()
		}
		defer func() {
			for _, ts := range tokenizerStreams {
				if ts != nil {
					ts.destroy()
				}
			}
		}()
		defer s.mutex.Unlock()
		_ = s.runGenerationLoop(ctx, generator, seqCount, 0, tokenizerStreams, maxLength, outputChan, errChan)
	}()
	return outputChan, errChan
}

func setDefaultGenerationOptions(generationOptions *GenerationOptions) {
	if generationOptions.Temperature == nil {
		generationOptions.Temperature = new(float64)
	}
	if generationOptions.TopP == nil {
		topP := 1.0
		generationOptions.TopP = &topP
	}
	if generationOptions.Seed == nil {
		generationOptions.Seed = new(int)
	}
}

func (s *Session) Generate(ctx context.Context, messages [][]Message, tools []string, generationOptions *GenerationOptions) (<-chan SequenceDelta, <-chan error, error) {
	s.mutex.Lock()
	if generationOptions == nil {
		generationOptions = &GenerationOptions{
			MaxLength: defaultMaxLength,
			BatchSize: len(messages),
		}
	}
	if generationOptions.BatchSize <= 0 {
		generationOptions.BatchSize = len(messages)
	}
	if generationOptions.MaxLength <= 0 {
		generationOptions.MaxLength = defaultMaxLength
	}
	setDefaultGenerationOptions(generationOptions)
	chatTokenizer, tokenizerErr := s.tokenizerForGeneration(generationOptions.EnableThinking)
	if tokenizerErr != nil {
		s.mutex.Unlock()
		return nil, nil, tokenizerErr
	}
	sequences, tokenizerStreams, tokenizeErr := chatTokenizer.tokenizeMessages(messages, tools)
	if tokenizeErr != nil {
		s.mutex.Unlock()
		return nil, nil, fmt.Errorf("TokenizeMessages failed: %w", tokenizeErr)
	}

	generator, err := s.createGenerator(generationOptions)
	if err != nil {
		sequences.destroy()
		for _, ts := range tokenizerStreams {
			ts.destroy()
		}
		s.mutex.Unlock()
		return nil, nil, err
	}
	if err = generator.addSequences(sequences); err != nil {
		sequences.destroy()
		for _, ts := range tokenizerStreams {
			ts.destroy()
		}
		generator.destroy()
		s.mutex.Unlock()
		return nil, nil, fmt.Errorf("failed to add sequences to generator: %w", err)
	}

	outputChan, errChan := s.startGenerationGoroutine(ctx, generator, sequences, nil, tokenizerStreams, len(messages), generationOptions.MaxLength)
	return outputChan, errChan, nil
}

// GenerateWithImages generates text using pre-processed named tensors (for multimodal inputs).
// Currently only supports a single prompt.
func (s *Session) GenerateWithImages(ctx context.Context, messages [][]Message, images *Images, tools []string, generationOptions *GenerationOptions) (<-chan SequenceDelta, <-chan error, error) {
	s.mutex.Lock()

	// Current limitation: underlying C API supports a single prompt with image tags
	if len(messages) != 1 {
		s.mutex.Unlock()
		return nil, nil, errors.New("GenerateWithImages currently supports only a single message set")
	}
	if generationOptions == nil {
		generationOptions = &GenerationOptions{MaxLength: defaultMaxLength}
	}
	if generationOptions.BatchSize != len(messages) {
		generationOptions.BatchSize = len(messages)
	}
	if generationOptions.MaxLength <= 0 {
		generationOptions.MaxLength = defaultMaxLength
	}
	setDefaultGenerationOptions(generationOptions)

	// Build prompt from chat template using messages[0]
	msgJSON, err := json.Marshal(messages[0])
	if err != nil {
		s.mutex.Unlock()
		return nil, nil, fmt.Errorf("failed to marshal input message: %w", err)
	}

	rawTools := make([]json.RawMessage, len(tools))
	for i, t := range tools {
		rawTools[i] = json.RawMessage(t)
	}
	toolsJSON, err := json.Marshal(rawTools)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal tools: %w", err)
	}

	chatTokenizer, tokenizerErr := s.tokenizerForGeneration(generationOptions.EnableThinking)
	if tokenizerErr != nil {
		s.mutex.Unlock()
		return nil, nil, tokenizerErr
	}
	prompt, templateErr := chatTokenizer.ApplyChatTemplate(msgJSON, toolsJSON, true)
	if templateErr != nil {
		s.mutex.Unlock()
		return nil, nil, fmt.Errorf("failed to apply chat template: %w", templateErr)
	}

	// Process images with the templated prompt (should include <|image_1|> etc.)
	if s.processor == nil {
		// Initialize multimodal processor
		err = initMultimodalProcessor(s)
		if err != nil {
			s.mutex.Unlock()
			return nil, nil, fmt.Errorf("initMultimodalProcessor failed: %w", err)
		}
	}

	tensors, err := s.processor.ProcessImages(prompt, images)
	if err != nil {
		s.mutex.Unlock()
		return nil, nil, fmt.Errorf("ProcessImages failed: %w", err)
	}

	generator, err := s.createGenerator(generationOptions)
	if err != nil {
		s.mutex.Unlock()
		return nil, nil, fmt.Errorf("failed to create generator: %w", err)
	}

	// Set the named tensors as inputs
	if err := generator.setInputs(tensors); err != nil {
		generator.destroy()
		s.mutex.Unlock()
		return nil, nil, fmt.Errorf("failed to set inputs: %w", err)
	}

	// Create tokenizer streams per sequence (align with Generate behavior)
	numSeq := generationOptions.BatchSize
	tokenizerStreams := make([]*tokenizerStream, 0, numSeq)
	for range numSeq {
		ts, err := chatTokenizer.createTokenizerStream()
		if err != nil {
			for _, t := range tokenizerStreams {
				if t != nil {
					t.destroy()
				}
			}
			generator.destroy()
			s.mutex.Unlock()
			return nil, nil, fmt.Errorf("failed to create tokenizer stream: %w", err)
		}
		tokenizerStreams = append(tokenizerStreams, ts)
	}
	outputChan, errChan := s.startGenerationGoroutine(ctx, generator, nil, tensors, tokenizerStreams, numSeq, generationOptions.MaxLength)
	return outputChan, errChan, nil
}

// MultimodalConversation pairs one conversation's messages with the media loaded
// for it. Image and audio references stay scoped to the conversation, so a batch
// may mix image, audio, image+audio, and text-only conversations and each keeps
// its own ordered media association.
type MultimodalConversation struct {
	Messages []Message
	Images   *Images
	Audios   *Audios
}

// multimodalRun owns the native resources created for one conversation.
type multimodalRun struct {
	generator *generator
	sequences *sequences
	tensors   *NamedTensors
	streams   []*tokenizerStream
}

func (r *multimodalRun) destroy() {
	if r == nil {
		return
	}
	for _, ts := range r.streams {
		if ts != nil {
			ts.destroy()
		}
	}
	r.streams = nil
	if r.tensors != nil {
		r.tensors.destroy()
		r.tensors = nil
	}
	if r.sequences != nil {
		r.sequences.destroy()
		r.sequences = nil
	}
	if r.generator != nil {
		r.generator.destroy()
		r.generator = nil
	}
}

var ErrNoConversations = errors.New("no conversations provided")

// GenerateMultimodal generates text for a batch of conversations, each carrying
// its own ordered media. Because the native multimodal API accepts exactly one
// prompt (with one flat image/audio list) per generation, the wrapper drives one
// single-prompt generation per conversation sequentially and relabels the emitted
// sequence so index i always refers to conversations[i].
//
// The native library is not thread-safe, so the session mutex is held for the
// whole batch and released (by the goroutine) on completion, error, or
// cancellation. All per-conversation native resources are destroyed on every
// path: setup errors, stream errors, cancellation, and normal completion.
func (s *Session) GenerateMultimodal(ctx context.Context, conversations []MultimodalConversation, tools []string, generationOptions *GenerationOptions) (<-chan SequenceDelta, <-chan error, error) {
	if len(conversations) == 0 {
		return nil, nil, ErrNoConversations
	}
	for i, conv := range conversations {
		if len(conv.Messages) == 0 {
			return nil, nil, fmt.Errorf("conversation %d has no messages", i)
		}
		if conv.Images != nil && conv.Images.imagesPtr == nil {
			return nil, nil, fmt.Errorf("conversation %d images are destroyed", i)
		}
		if conv.Audios != nil && conv.Audios.audiosPtr == nil {
			return nil, nil, fmt.Errorf("conversation %d audios are destroyed", i)
		}
	}
	if generationOptions == nil {
		generationOptions = &GenerationOptions{MaxLength: defaultMaxLength}
	}
	setDefaultGenerationOptions(generationOptions)
	if generationOptions.BatchSize <= 0 {
		generationOptions.BatchSize = 1
	}
	if generationOptions.MaxLength <= 0 {
		generationOptions.MaxLength = defaultMaxLength
	}
	var enableThinking *bool
	if generationOptions.EnableThinking != nil {
		value := *generationOptions.EnableThinking
		enableThinking = &value
	}
	maxLength := generationOptions.MaxLength

	rawTools := make([]json.RawMessage, len(tools))
	for i, t := range tools {
		rawTools[i] = json.RawMessage(t)
	}
	toolsJSON, err := json.Marshal(rawTools)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal tools: %w", err)
	}

	outputChan := make(chan SequenceDelta, 1000)
	errChan := make(chan error, 1)
	go func() {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		defer close(outputChan)
		defer close(errChan)
		chatTokenizer, tokenizerErr := s.tokenizerForGeneration(enableThinking)
		if tokenizerErr != nil {
			sendGenerationError(errChan, tokenizerErr)
			return
		}
		// Destroy in reverse creation order; destroy() is idempotent.
		var runs []*multimodalRun
		defer func() {
			for _, run := range slices.Backward(runs) {
				run.destroy()
			}
		}()
		for i := range conversations {
			select {
			case <-ctx.Done():
				return
			default:
			}
			run, setupErr := s.setupMultimodalRun(chatTokenizer, conversations[i], toolsJSON, generationOptions)
			if setupErr != nil {
				sendGenerationError(errChan, setupErr)
				return
			}
			runs = append(runs, run)
			if loopErr := s.runGenerationLoop(ctx, run.generator, 1, i, run.streams, maxLength, outputChan, errChan); loopErr != nil {
				// Native errors were already relayed on errChan; a plain
				// context cancellation is not an error worth surfacing.
				return
			}
		}
	}()
	return outputChan, errChan, nil
}

// setupMultimodalRun builds the native generator for one conversation. For media
// conversations the chat-templated prompt is routed through the session's
// multimodal processor (images, audio, or both); for text-only conversations the
// prompt is tokenized into sequences. Assumes the session mutex is held.
func (s *Session) setupMultimodalRun(chatTokenizer *tokenizer, conv MultimodalConversation, toolsJSON []byte, generationOptions *GenerationOptions) (*multimodalRun, error) {
	if chatTokenizer == nil || chatTokenizer.tokenizerPtr == nil || s.model == nil || s.model.modelPtr == nil {
		return nil, errors.New("session is destroyed")
	}
	run := &multimodalRun{}

	msgJSON, err := json.Marshal(conv.Messages)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal conversation messages: %w", err)
	}

	hasImages := conv.Images != nil && conv.Images.imagesPtr != nil
	hasAudios := conv.Audios != nil && conv.Audios.audiosPtr != nil

	// A per-conversation generator is single-sequence by native design.
	opts := *generationOptions
	opts.BatchSize = 1

	if !hasImages && !hasAudios {
		sequences, streams, err := chatTokenizer.tokenizeMessages([][]Message{conv.Messages}, nil)
		if err != nil {
			run.destroy()
			return nil, fmt.Errorf("tokenizing conversation messages: %w", err)
		}
		run.sequences = sequences
		run.streams = streams
		run.generator, err = s.createGenerator(&opts)
		if err != nil {
			run.destroy()
			return nil, fmt.Errorf("creating generator: %w", err)
		}
		if err = run.generator.addSequences(sequences); err != nil {
			run.destroy()
			return nil, fmt.Errorf("adding sequences to generator: %w", err)
		}
		return run, nil
	}

	prompt, templateErr := chatTokenizer.ApplyChatTemplate(msgJSON, toolsJSON, true)
	if templateErr != nil {
		return nil, fmt.Errorf("failed to apply chat template: %w", templateErr)
	}

	if s.processor == nil {
		if err = initMultimodalProcessor(s); err != nil {
			return nil, fmt.Errorf("initMultimodalProcessor failed: %w", err)
		}
	}

	var tensors *NamedTensors
	switch {
	case hasImages && hasAudios:
		tensors, err = s.processor.ProcessImagesAndAudios(prompt, conv.Images, conv.Audios)
	case hasAudios:
		tensors, err = s.processor.ProcessAudios(prompt, conv.Audios)
	default:
		tensors, err = s.processor.ProcessImages(prompt, conv.Images)
	}
	if err != nil {
		return nil, fmt.Errorf("processing multimodal inputs: %w", err)
	}
	run.tensors = tensors

	run.generator, err = s.createGenerator(&opts)
	if err != nil {
		run.destroy()
		return nil, fmt.Errorf("creating generator: %w", err)
	}
	if err = run.generator.setInputs(tensors); err != nil {
		run.destroy()
		return nil, fmt.Errorf("failed to set inputs: %w", err)
	}

	stream, streamErr := chatTokenizer.createTokenizerStream()
	if streamErr != nil {
		run.destroy()
		return nil, fmt.Errorf("failed to create tokenizer stream: %w", streamErr)
	}
	run.streams = []*tokenizerStream{stream}
	return run, nil
}

func (s *Session) Destroy() {
	if s.activeAdapters != nil {
		s.activeAdapters.Destroy()
	}
	for _, tokenizer := range s.thinkingTokenizers {
		tokenizer.destroy()
	}
	s.thinkingTokenizers = nil
	if s.tokenizer != nil {
		s.tokenizer.destroy()
		s.tokenizer = nil
	}
	if s.processor != nil {
		s.processor.destroy()
		s.processor = nil
	}
	if s.model != nil {
		s.model.destroy()
		s.model = nil
	}
	for _, buffer := range s.nativeBuffers {
		C.free(buffer)
	}
	s.nativeBuffers = nil
}

// CreateSessionWithOptions builds a GenAI config from a config directory,
// applies execution providers and options, creates the model and tokenizer, and returns a Session.
// providers: list of EP names in priority order (e.g., ["cuda"], ["NvTensorRtRtx"], ["OpenVINO"]).
// providerOptions: map of EP name -> map of key/value options.
func CreateSessionWithOptions(configDirectoryPath string, providers []string, providerOptions map[string]map[string]string) (*Session, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}

	// Create Config
	var cfg *C.OgaConfig
	cConfigPath := C.CString(configDirectoryPath)
	defer C.free(unsafe.Pointer(cConfigPath))
	res := C.CreateOgaConfig(cConfigPath, &cfg)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("CreateOgaConfig failed: %w", err)
	}
	if cfg == nil {
		return nil, errors.New("CreateOgaConfig returned nil without error")
	}
	defer C.DestroyOgaConfig(cfg)

	// Clear default providers to allow explicit configuration
	res = C.OgaConfigClearProviders(cfg)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("OgaConfigClearProviders failed: %w", err)
	}

	// Append providers and set options
	for _, providerName := range providers {
		cp := C.CString(providerName)
		res = C.OgaConfigAppendProvider(cfg, cp)
		if err := OgaResultToError(res); err != nil {
			C.free(unsafe.Pointer(cp))
			return nil, fmt.Errorf("OgaConfigAppendProvider(%s) failed: %w", providerName, err)
		}
		if opts, ok := providerOptions[providerName]; ok {
			for k, v := range opts {
				ck := C.CString(k)
				cv := C.CString(v)
				res = C.OgaConfigSetProviderOption(cfg, cp, ck, cv)
				C.free(unsafe.Pointer(ck))
				C.free(unsafe.Pointer(cv))
				if err := OgaResultToError(res); err != nil {
					C.free(unsafe.Pointer(cp))
					return nil, fmt.Errorf("OgaConfigSetProviderOption(%s,%s=%s) failed: %w", providerName, k, v, err)
				}
			}
		}
		C.free(unsafe.Pointer(cp))
	}

	// Create Model from Config
	var cModel *C.OgaModel
	res = C.CreateOgaModelFromConfig(cfg, &cModel)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("CreateOgaModelFromConfig failed: %w", err)
	}
	if cModel == nil {
		return nil, errors.New("CreateOgaModelFromConfig returned nil without error")
	}

	// Create Tokenizer
	model := model{modelPtr: cModel}
	tokenizer, err := newTokenizerFromModel(model)
	if err != nil {
		C.DestroyOgaModel(cModel)
		return nil, fmt.Errorf("newTokenizerFromModel failed: %w", err)
	}

	session := &Session{
		modelPath:  configDirectoryPath,
		model:      &model,
		tokenizer:  &tokenizer,
		statistics: &Statistics{},
	}
	return session, nil
}

func initMultimodalProcessor(session *Session) error {
	processor, errProcessor := createMultiModalProcessor(session.model)
	if errProcessor != nil {
		return fmt.Errorf("createMultiModalProcessor failed: %w", errProcessor)
	}
	session.processor = processor
	return nil
}

func CreateSession(modelPath string) (*Session, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if modelPath == "" {
		return nil, errors.New("modelPath is empty")
	}
	cPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cPath))

	var cModel *C.OgaModel
	res := C.CreateOgaModel(cPath, &cModel)
	if err := OgaResultToError(res); err != nil {
		return nil, fmt.Errorf("CreateOgaModel failed: %w", err)
	}
	if cModel == nil {
		return nil, errors.New("CreateOgaModel returned nil model without error")
	}

	model := model{modelPtr: cModel}
	tokenizer, err := newTokenizerFromModel(model)
	if err != nil {
		C.DestroyOgaModel(cModel)
		return nil, fmt.Errorf("newTokenizerFromModel failed: %w", err)
	}
	session := &Session{
		modelPath:  modelPath,
		model:      &model,
		tokenizer:  &tokenizer,
		statistics: &Statistics{},
	}
	return session, nil
}

// NonGenerativeSession runs ranking or decision models that accept and return JSON.
type NonGenerativeSession struct {
	ptr   *C.OgaNonGenerativeSession
	mutex sync.Mutex
}

// CreateNonGenerativeSession creates a non-generative session from a model config directory.
func CreateNonGenerativeSession(configDirectory string) (*NonGenerativeSession, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	path := C.CString(configDirectory)
	defer C.free(unsafe.Pointer(path))
	var config *C.OgaConfig
	if err := OgaResultToError(C.CreateOgaConfig(path, &config)); err != nil {
		return nil, fmt.Errorf("creating non-generative config: %w", err)
	}
	if config == nil {
		return nil, errors.New("non-generative config creation returned nil without error")
	}
	defer C.DestroyOgaConfig(config)
	var ptr *C.OgaNonGenerativeSession
	if err := OgaResultToError(C.CreateNonGenerativeSession(config, &ptr)); err != nil {
		return nil, fmt.Errorf("creating non-generative session: %w", err)
	}
	if ptr == nil {
		return nil, errors.New("non-generative session creation returned nil without error")
	}
	return &NonGenerativeSession{ptr: ptr}, nil
}

// RunJSON sends one JSON request and returns the copied JSON response.
func (s *NonGenerativeSession) RunJSON(request []byte) (string, error) {
	if s == nil {
		return "", errors.New("non-generative session is destroyed")
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.ptr == nil {
		return "", errors.New("non-generative session is destroyed")
	}
	if len(request) == 0 {
		return "", errors.New("non-generative request is empty")
	}
	input := C.CBytes(request)
	defer C.free(input)
	var output *C.char
	if err := OgaResultToError(C.NonGenerativeSessionRunJson(s.ptr, (*C.char)(input), C.size_t(len(request)), (**C.char)(unsafe.Pointer(&output)))); err != nil {
		if output != nil {
			C.DestroyOgaString(output)
		}
		return "", fmt.Errorf("running non-generative request: %w", err)
	}
	if output == nil {
		return "", errors.New("non-generative session returned a nil response")
	}
	response := C.GoString(output)
	C.DestroyOgaString(output)
	return response, nil
}

// Destroy releases the native session.
func (s *NonGenerativeSession) Destroy() {
	if s != nil {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	if s != nil && s.ptr != nil {
		C.DestroyNonGenerativeSession(s.ptr)
		s.ptr = nil
	}
}

// CreateSessionFromMemory builds a CPU text session from in-memory GenAI config and ONNX files.
// ONNX graph and external initializer bytes are copied into C memory and kept alive until Session.Destroy.
func CreateSessionFromMemory(configJSON []byte, modelData, externalInitializerData map[string][]byte, tokenizerConfigJSON, tokenizerJSON []byte) (*Session, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}
	if len(configJSON) == 0 {
		return nil, errors.New("in-memory model config is empty")
	}
	configBytes := C.CBytes(configJSON)
	defer C.free(configBytes)
	var config *C.OgaConfig
	if err := OgaResultToError(C.CreateConfigFromBuffer((*C.char)(configBytes), C.size_t(len(configJSON)), &config)); err != nil {
		return nil, fmt.Errorf("creating in-memory model config: %w", err)
	}
	if config == nil {
		return nil, errors.New("in-memory model config creation returned nil without error")
	}
	defer C.DestroyOgaConfig(config)

	var buffers []unsafe.Pointer
	freeBuffers := func() {
		for _, buffer := range buffers {
			C.free(buffer)
		}
		buffers = nil
	}
	defer func() {
		if buffers != nil {
			freeBuffers()
		}
	}()
	addData := func(filename string, data []byte, external bool) error {
		if filename == "" || len(data) == 0 {
			return errors.New("model data filename and bytes are required")
		}
		cFilename := C.CString(filename)
		defer C.free(unsafe.Pointer(cFilename))
		dataPtr := C.CBytes(data)
		var result *C.OgaResult
		if external {
			result = C.ConfigAddExternalInitializerData(config, cFilename, dataPtr, C.size_t(len(data)))
		} else {
			result = C.ConfigAddModelData(config, cFilename, dataPtr, C.size_t(len(data)))
		}
		if err := OgaResultToError(result); err != nil {
			C.free(dataPtr)
			return err
		}
		buffers = append(buffers, dataPtr)
		return nil
	}
	for filename, data := range modelData {
		if err := addData(filename, data, false); err != nil {
			return nil, fmt.Errorf("registering model file %q: %w", filename, err)
		}
	}
	for filename, data := range externalInitializerData {
		if err := addData(filename, data, true); err != nil {
			return nil, fmt.Errorf("registering external initializer %q: %w", filename, err)
		}
	}
	if len(tokenizerConfigJSON) != 0 || len(tokenizerJSON) != 0 {
		if len(tokenizerConfigJSON) == 0 || len(tokenizerJSON) == 0 {
			return nil, errors.New("both tokenizer config JSON and tokenizer JSON are required")
		}
		tokenizerConfig := C.CBytes(tokenizerConfigJSON)
		tokenizerData := C.CBytes(tokenizerJSON)
		result := C.ConfigSetTokenizerData(config, (*C.char)(tokenizerConfig), C.size_t(len(tokenizerConfigJSON)), (*C.char)(tokenizerData), C.size_t(len(tokenizerJSON)))
		C.free(tokenizerConfig)
		C.free(tokenizerData)
		if err := OgaResultToError(result); err != nil {
			return nil, fmt.Errorf("registering tokenizer JSON: %w", err)
		}
	}
	var nativeModel *C.OgaModel
	if err := OgaResultToError(C.CreateOgaModelFromConfig(config, &nativeModel)); err != nil {
		if nativeModel != nil {
			C.DestroyOgaModel(nativeModel)
		}
		return nil, fmt.Errorf("creating in-memory model: %w", err)
	}
	if nativeModel == nil {
		return nil, errors.New("in-memory model creation returned nil without error")
	}
	model := &model{modelPtr: nativeModel}
	tokenizer, err := newTokenizerFromModel(*model)
	if err != nil {
		model.destroy()
		return nil, fmt.Errorf("creating in-memory model tokenizer: %w", err)
	}
	session := &Session{model: model, tokenizer: &tokenizer, statistics: &Statistics{}, nativeBuffers: buffers}
	buffers = nil
	return session, nil
}
