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
	"path/filepath"
	"slices"
	"time"
	"unsafe"
)

// mtpRun owns the native resources created for one conversation's MTP
// speculative run. The MTP generator owns both its main and speculative model
// handles (see MTPGenerator) and is independent of the session's model and
// tokenizer, so destroying the run must not touch any session-owned object.
type mtpRun struct {
	sequences    *sequences
	stream       *tokenizerStream
	tokenIDs     []int32 // copied prompt token IDs (used to slice the generated tail off Sequence())
	mtpGenerator *MTPGenerator
}

func (r *mtpRun) destroy() {
	if r == nil {
		return
	}
	if r.stream != nil {
		r.stream.destroy()
		r.stream = nil
	}
	if r.sequences != nil {
		r.sequences.destroy()
		r.sequences = nil
	}
	r.tokenIDs = nil
	if r.mtpGenerator != nil {
		r.mtpGenerator.Destroy()
		r.mtpGenerator = nil
	}
}

// tokenIDsFromSequences copies the first sequence's int32 token array out of a
// standalone *sequences handle. Assumed invariant: the sequences hold exactly
// one sequence (a single chat-templated prompt) at call time.
func tokenIDsFromSequences(seqs *sequences) ([]int32, error) {
	if seqs == nil || seqs.sequencesPtr == nil {
		return nil, errors.New("sequences handle is nil or destroyed")
	}
	if C.SequencesCount(seqs.sequencesPtr) < 1 {
		return []int32{}, nil
	}
	count := int(C.SequencesGetSequenceCount(seqs.sequencesPtr, 0))
	if count == 0 {
		return []int32{}, nil
	}
	data := C.SequencesGetSequenceData(seqs.sequencesPtr, 0)
	if data == nil {
		return nil, errors.New("sequences handle returned nil sequence data")
	}
	values := unsafe.Slice((*int32)(unsafe.Pointer(data)), count)
	return append([]int32(nil), values...), nil
}

// GenerateMTP generates text for a batch of text-only conversations using the
// session's MTP speculative path. The call mirrors Generate in shape: one
// output stream, one error stream, same SequenceDelta contract — so HuGo's
// downstream adapters and tests (stop-buffering, cancellation, statistics)
// can treat the two as interchangeable.
//
// messages and tools follow the same per-conversation semantics used in
// Generate: messages is a list of conversations, each a sequence of chat
// messages; tools is a flat list of tool definitions (JSON strings). All
// conversations are text-only — the [][]Message parameter type does not
// accept image/audio inputs (use GenerateMultimodal for those).
//
// The native MTP library models a greedy single-sequence draft/verify
// generator that owns its own pair of model handles. This wrapper therefore
// drives exactly one MTP generation per conversation sequentially and
// relabels the emitted sequence so index i always refers to messages[i].
//
// The speculative model is loaded from <modelPath>/MTP relative to the
// directory used by CreateSession — matching the pinned ortgenai mtp.go's
// convention.
//
// GenerationOptions.BatchSize is honored as 1 per conversation (matching
// GenerateWithImages); the pinned MTP path is greedy, so Temperature/TopP/
// Seed/Guidance are ignored. MaxLength caps the total committed length of the
// sequence (including prompt), consistent with Generate.
//
// The `Sequence` field of each SequenceDelta is the 0-based conversation index
// in messages, matching Generate.
//
// Ownership: the session mutex is held for the lifetime of the run, released
// on completion / error / cancellation, and all per-run native resources (MTP
// generator, prompt sequences, tokenizer streams) are destroyed on every path.
func (s *Session) GenerateMTP(ctx context.Context, messages [][]Message, tools []string, generationOptions *GenerationOptions) (<-chan SequenceDelta, <-chan error, error) {
	if len(messages) == 0 {
		return nil, nil, ErrNoConversations
	}
	for i, conv := range messages {
		if len(conv) == 0 {
			return nil, nil, fmt.Errorf("conversation %d has no messages", i)
		}
	}
	if s.modelPath == "" {
		return nil, nil, errors.New("session has no model path: MTP generation requires a model directory")
	}
	if generationOptions == nil {
		generationOptions = &GenerationOptions{MaxLength: defaultMaxLength}
	}
	if generationOptions.BatchSize <= 0 {
		generationOptions.BatchSize = 1
	}
	maxLength := generationOptions.MaxLength
	if maxLength <= 0 {
		maxLength = defaultMaxLength
	}

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

		mtpModelPath := filepath.Join(s.modelPath, "MTP")

		// Destroy in reverse creation order; destroy() is idempotent.
		var runs []*mtpRun
		defer func() {
			for _, run := range slices.Backward(runs) {
				run.destroy()
			}
		}()

		for i := range messages {
			select {
			case <-ctx.Done():
				return
			default:
			}
			run, setupErr := s.setupMTPRun(messages[i], toolsJSON, mtpModelPath, maxLength)
			if setupErr != nil {
				sendGenerationError(errChan, setupErr)
				return
			}
			runs = append(runs, run)
			if loopErr := s.runMTPLoop(ctx, run, i, outputChan, errChan); loopErr != nil {
				// Native errors were already relayed on errChan; a plain
				// context cancellation is not an error worth surfacing.
				return
			}
		}
	}()
	return outputChan, errChan, nil
}

// setupMTPRun builds the native MTP generator and per-conversation resources
// for one conversation. The MTP generator is created per-conversation (the
// pinned MTP API requires both model paths at construction time and loads
// its own handles), mirroring the multimodalRun model of one native
// generator per conversation. Assumes the session mutex is held.
func (s *Session) setupMTPRun(conv []Message, toolsJSON []byte, mtpModelPath string, maxLength int) (*mtpRun, error) {
	if s.tokenizer == nil || s.tokenizer.tokenizerPtr == nil || s.model == nil || s.model.modelPtr == nil {
		return nil, errors.New("session is destroyed")
	}
	run := &mtpRun{}

	msgJSON, err := json.Marshal(conv)
	if err != nil {
		run.destroy()
		return nil, fmt.Errorf("failed to marshal conversation messages: %w", err)
	}
	prompt, err := s.tokenizer.ApplyChatTemplate(msgJSON, toolsJSON, true)
	if err != nil {
		run.destroy()
		return nil, fmt.Errorf("failed to apply chat template: %w", err)
	}

	var cSequences *C.OgaSequences
	if err := OgaResultToError(C.CreateOgaSequences(&cSequences)); err != nil {
		run.destroy()
		return nil, fmt.Errorf("creating sequences: %w", err)
	}
	if cSequences == nil {
		run.destroy()
		return nil, errors.New("CreateOgaSequences returned nil without error")
	}
	run.sequences = &sequences{sequencesPtr: cSequences}

	if err := s.tokenizer.encode(prompt, run.sequences); err != nil {
		run.destroy()
		return nil, fmt.Errorf("encoding prompt: %w", err)
	}

	tokenIDs, err := tokenIDsFromSequences(run.sequences)
	if err != nil {
		run.destroy()
		return nil, fmt.Errorf("extracting token IDs from encoded prompt: %w", err)
	}
	if len(tokenIDs) == 0 {
		run.destroy()
		return nil, errors.New("encoded prompt produced no tokens")
	}
	run.tokenIDs = tokenIDs

	mtpGen, err := CreateMTPGenerator(s.modelPath, mtpModelPath, maxLength)
	if err != nil {
		run.destroy()
		return nil, fmt.Errorf("creating MTP generator: %w", err)
	}
	run.mtpGenerator = mtpGen

	if err := mtpGen.AppendTokens(tokenIDs); err != nil {
		run.destroy()
		return nil, fmt.Errorf("appending prompt tokens: %w", err)
	}

	stream, err := s.tokenizer.createTokenizerStream()
	if err != nil {
		run.destroy()
		return nil, fmt.Errorf("creating tokenizer stream: %w", err)
	}
	run.stream = stream
	return run, nil
}

// runMTPLoop drives the MTP generator until it is done or the context is
// cancelled, decoding committed tokens and emitting SequenceDelta values on
// outChan. The `Sequence` field of emitted deltas is the 0-based conversation
// index `seqOffset`. Generation errors are relayed on errChan and also
// returned. Assumes the caller already holds s.mutex and the run's prompt
// tokens have been appended; it does not destroy the run — the caller is
// responsible.
func (s *Session) runMTPLoop(ctx context.Context, run *mtpRun, seqOffset int, outChan chan<- SequenceDelta, errChan chan<- error) error {
	runStart := time.Now()
	var firstToken time.Time
	tokenCount := 0
	eosSeen := false
	prevSeqLen := len(run.tokenIDs)

	// finalize tokens/sec statistics at the end of the run (mirrors runGenerationLoop)
	defer func() {
		if !firstToken.IsZero() && tokenCount > 0 {
			dur := time.Since(firstToken).Seconds()
			if dur > 0 {
				s.statistics.CumulativeTokenDurationSeconds += dur
				s.statistics.TokensPerSecond = float64(s.statistics.CumulativeTokens) / s.statistics.CumulativeTokenDurationSeconds
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		done, doneErr := run.mtpGenerator.IsDone()
		if doneErr != nil {
			sendGenerationError(errChan, doneErr)
			return doneErr
		}
		if done {
			// Distinguish EOS from max-length the same way Generate does.
			if !eosSeen {
				sendGenerationError(errChan, MaxLengthReachedError{})
			}
			return nil
		}

		if err := run.mtpGenerator.GenerateNextToken(); err != nil {
			sendGenerationError(errChan, err)
			return err
		}

		seq, err := run.mtpGenerator.Sequence()
		if err != nil {
			sendGenerationError(errChan, err)
			return err
		}

		for j := prevSeqLen; j < len(seq); j++ {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			token := seq[j]
			if slices.Contains(s.tokenizer.EOSTokenIDs, int(token)) {
				eosSeen = true
				select {
				case outChan <- SequenceDelta{Sequence: seqOffset, EOSReached: true}:
				case <-ctx.Done():
					return ctx.Err()
				}
				return nil
			}

			decoded, decodeErr := run.stream.Decode(C.int32_t(token))
			if decodeErr != nil {
				sendGenerationError(errChan, decodeErr)
				return decodeErr
			}
			if decoded == "" {
				continue
			}

			if firstToken.IsZero() {
				firstToken = time.Now()
				prefill := firstToken.Sub(runStart).Seconds()
				s.statistics.CumulativePrefillSum += prefill
				s.statistics.CumulativePrefillCount++
				s.statistics.AvgPrefillSeconds = s.statistics.CumulativePrefillSum / float64(s.statistics.CumulativePrefillCount)
			}
			s.statistics.CumulativeTokens++
			tokenCount++

			select {
			case outChan <- SequenceDelta{Sequence: seqOffset, Token: decoded}:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
		prevSeqLen = len(seq)
	}
}
