package ortgenai

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// ASRSequenceResult associates the post-processed ASR results with the input
// conversation index passed to Session.GenerateASR.
type ASRSequenceResult struct {
	Sequence int
	Results  []ASRResult
}

// GenerateASR generates a transcription for each audio conversation and
// returns its text and any word or segment timestamps. Each conversation must
// include audio; its Messages provide the prompt expected by the model.
// Keep the conversations, audio, image, and session resources alive and
// unchanged until both returned channels are closed.
// Results are emitted as each conversation completes. The ASR post-processor
// consumes generator output incrementally and is finalized before results are
// copied to Go values.
func (s *Session) GenerateASR(ctx context.Context, conversations []MultimodalConversation, generationOptions *GenerationOptions) (<-chan ASRSequenceResult, <-chan error, error) {
	if len(conversations) == 0 {
		return nil, nil, ErrNoConversations
	}
	for i, conv := range conversations {
		if len(conv.Messages) == 0 {
			return nil, nil, fmt.Errorf("conversation %d has no messages", i)
		}
		if conv.Audios == nil || conv.Audios.audiosPtr == nil {
			return nil, nil, fmt.Errorf("conversation %d has no audio or its audios are destroyed", i)
		}
		if conv.Images != nil && conv.Images.imagesPtr == nil {
			return nil, nil, fmt.Errorf("conversation %d images are destroyed", i)
		}
	}

	opts := GenerationOptions{MaxLength: defaultMaxLength, BatchSize: 1}
	if generationOptions != nil {
		opts = *generationOptions
		if generationOptions.Temperature != nil {
			temperature := *generationOptions.Temperature
			opts.Temperature = &temperature
		}
		if generationOptions.TopP != nil {
			topP := *generationOptions.TopP
			opts.TopP = &topP
		}
		if generationOptions.Seed != nil {
			seed := *generationOptions.Seed
			opts.Seed = &seed
		}
		if generationOptions.Guidance != nil {
			guidance := *generationOptions.Guidance
			opts.Guidance = &guidance
		}
		if generationOptions.EnableThinking != nil {
			enableThinking := *generationOptions.EnableThinking
			opts.EnableThinking = &enableThinking
		}
		if opts.MaxLength <= 0 {
			opts.MaxLength = defaultMaxLength
		}
	}
	opts.BatchSize = 1
	setDefaultGenerationOptions(&opts)

	outputChan := make(chan ASRSequenceResult, len(conversations))
	errChan := make(chan error, 1)
	go func() {
		defer close(outputChan)
		defer close(errChan)
		s.mutex.Lock()
		defer s.mutex.Unlock()
		chatTokenizer, tokenizerErr := s.tokenizerForGeneration(opts.EnableThinking)
		if tokenizerErr != nil {
			sendGenerationError(errChan, tokenizerErr)
			return
		}

		for i, conv := range conversations {
			select {
			case <-ctx.Done():
				return
			default:
			}

			results, err := s.runASRConversation(chatTokenizer, ctx, conv, &opts)
			if err != nil {
				if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
					sendGenerationError(errChan, err)
				}
				return
			}
			select {
			case outputChan <- ASRSequenceResult{Sequence: i, Results: results}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return outputChan, errChan, nil
}

// runASRConversation sets up audio inputs and drives one generator. The caller
// holds the session mutex for the full run.
func (s *Session) runASRConversation(chatTokenizer *tokenizer, ctx context.Context, conv MultimodalConversation, generationOptions *GenerationOptions) ([]ASRResult, error) {
	if chatTokenizer == nil || chatTokenizer.tokenizerPtr == nil || s.model == nil || s.model.modelPtr == nil {
		return nil, errors.New("session is destroyed")
	}

	messageJSON, err := json.Marshal(conv.Messages)
	if err != nil {
		return nil, fmt.Errorf("marshaling ASR messages: %w", err)
	}
	prompt, err := chatTokenizer.ApplyChatTemplate(messageJSON, []byte("[]"), true)
	if err != nil {
		return nil, fmt.Errorf("applying ASR chat template: %w", err)
	}
	if s.processor == nil {
		if err := initMultimodalProcessor(s); err != nil {
			return nil, fmt.Errorf("initializing audio processor: %w", err)
		}
	}

	var tensors *NamedTensors
	if conv.Images != nil && conv.Images.imagesPtr != nil {
		tensors, err = s.processor.ProcessImagesAndAudios(prompt, conv.Images, conv.Audios)
	} else {
		tensors, err = s.processor.ProcessAudios(prompt, conv.Audios)
	}
	if err != nil {
		return nil, fmt.Errorf("preprocessing ASR audio: %w", err)
	}
	run := &multimodalRun{tensors: tensors}
	defer run.destroy()

	opts := *generationOptions
	opts.BatchSize = 1
	run.generator, err = s.createGenerator(&opts)
	if err != nil {
		return nil, fmt.Errorf("creating ASR generator: %w", err)
	}
	if err := run.generator.setInputs(tensors); err != nil {
		return nil, fmt.Errorf("setting ASR audio inputs: %w", err)
	}

	postProcessor, err := run.generator.newASRPostProcessor()
	if err != nil {
		return nil, err
	}
	// This defer runs before run.destroy, so the post-processor never outlives
	// the generator it references.
	defer postProcessor.Destroy()

	for !run.generator.IsDone() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		if err := run.generator.generateNextToken(); err != nil {
			return nil, err
		}
		if err := postProcessor.ProcessNext(); err != nil {
			return nil, err
		}
	}
	if err := postProcessor.Finalize(); err != nil {
		return nil, err
	}
	return postProcessor.Results()
}
