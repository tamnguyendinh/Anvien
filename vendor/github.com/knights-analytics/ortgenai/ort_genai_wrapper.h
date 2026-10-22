#ifndef ORT_GENAI_WRAPPER_H
#define ORT_GENAI_WRAPPER_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#undef _WIN32
#include "ort_genai_c.h"

// ... However, mingw will complain if _WIN32 is *not* defined! So redefine it.
#define _WIN32

#ifdef __cplusplus
extern "C" {
#endif

typedef struct OgaResult OgaResult;
typedef struct OgaModel OgaModel;
typedef struct OgaTokenizer OgaTokenizer;
typedef struct OgaTokenizerStream OgaTokenizerStream;
typedef struct OgaSequences OgaSequences;
typedef struct OgaGenerator OgaGenerator;
typedef struct OgaGeneratorParams OgaGeneratorParams;
typedef struct OgaConfig OgaConfig;
typedef struct OgaImages OgaImages;
typedef struct OgaMultiModalProcessor OgaMultiModalProcessor;
typedef struct OgaNamedTensors OgaNamedTensors;
typedef struct OgaStringArray OgaStringArray;
typedef struct OgaEngine OgaEngine;
typedef struct OgaRequest OgaRequest;

// Function pointer typedefs for the subset of GenAI C API we wrap.
typedef OgaResult* (*PFN_OgaCreateModel)(const char*, OgaModel**);
typedef OgaResult* (*PFN_OgaCreateTokenizer)(const OgaModel*, OgaTokenizer**);
typedef OgaResult* (*PFN_OgaCreateTokenizerStream)(const OgaTokenizer*, OgaTokenizerStream**);
typedef const char* (*PFN_OgaResultGetError)(const OgaResult*);
typedef void (*PFN_OgaDestroyResult)(OgaResult*);
typedef void (*PFN_OgaDestroyModel)(OgaModel*);
typedef void (*PFN_OgaDestroyTokenizer)(OgaTokenizer*);
typedef void (*PFN_OgaDestroyTokenizerStream)(OgaTokenizerStream*);
typedef OgaResult* (*PFN_OgaTokenizerApplyChatTemplate)(const OgaTokenizer*, const char*, const char*, const char*, bool, const char**);
typedef void (*PFN_OgaDestroyString)(const char*);
typedef OgaResult* (*PFN_OgaCreateSequences)(OgaSequences**);
typedef void (*PFN_OgaDestroySequences)(OgaSequences*);
typedef OgaResult* (*PFN_OgaTokenizerEncode)(const OgaTokenizer*, const char*, OgaSequences*);
typedef OgaResult* (*PFN_OgaCreateGenerator)(const OgaModel*, const OgaGeneratorParams*, OgaGenerator**);
typedef void (*PFN_OgaDestroyGenerator)(OgaGenerator*);
typedef OgaResult* (*PFN_OgaCreateGeneratorParams)(const OgaModel*, OgaGeneratorParams**);
typedef void (*PFN_OgaDestroyGeneratorParams)(OgaGeneratorParams*);
typedef OgaResult* (*PFN_OgaGeneratorParamsSetSearchNumber)(OgaGeneratorParams*, const char*, double);
typedef OgaResult* (*PFN_OgaGeneratorParamsSetGuidance)(OgaGeneratorParams*, const char*, const char*, bool);
typedef OgaResult* (*PFN_OgaGeneratorAppendTokenSequences)(OgaGenerator*, const OgaSequences*);
typedef OgaResult* (*PFN_OgaGeneratorSetInputs)(OgaGenerator*, const OgaNamedTensors*);
typedef OgaResult* (*PFN_OgaGeneratorGenerateNextToken)(OgaGenerator*);
typedef size_t (*PFN_OgaGeneratorGetSequenceCount)(const OgaGenerator*, size_t);
typedef const int32_t* (*PFN_OgaGeneratorGetSequenceData)(const OgaGenerator*, size_t);
typedef OgaResult* (*PFN_OgaTokenizerStreamDecode)(OgaTokenizerStream*, int32_t, const char**);
typedef bool (*PFN_OgaGeneratorIsDone)(const OgaGenerator*);
typedef OgaResult* (*PFN_OgaTokenizerGetEosTokenIds)(const OgaTokenizer*, const int32_t** , size_t*);
typedef void (*PFN_OgaShutdown)();
typedef void (*PFN_OgaSetTelemetryEnabled)(bool);
typedef OgaResult* (*PFN_OgaCreateConfigFromPackageEp)(const char*, const char*, OgaConfig**);
typedef OgaResult* (*PFN_OgaTokenizerGetPadTokenId)(const OgaTokenizer*, int32_t*);
typedef OgaResult* (*PFN_OgaTokenizerGetBotTokenId)(const OgaTokenizer*, int32_t*);
typedef OgaResult* (*PFN_OgaTokenizerGetEotTokenId)(const OgaTokenizer*, int32_t*);
typedef OgaResult* (*PFN_OgaTokenizerGetBorTokenId)(const OgaTokenizer*, int32_t*);
typedef OgaResult* (*PFN_OgaTokenizerGetEorTokenId)(const OgaTokenizer*, int32_t*);

// Config-related API
typedef OgaResult* (*PFN_OgaCreateConfig)(const char*, OgaConfig**);
typedef OgaResult* (*PFN_OgaConfigClearProviders)(OgaConfig*);
typedef OgaResult* (*PFN_OgaConfigAppendProvider)(OgaConfig*, const char*);
typedef OgaResult* (*PFN_OgaConfigSetProviderOption)(OgaConfig*, const char*, const char*, const char*);
typedef OgaResult* (*PFN_OgaCreateModelFromConfig)(const OgaConfig*, OgaModel**);
typedef void (*PFN_OgaDestroyConfig)(OgaConfig*);

// Multimodal API
typedef OgaResult* (*PFN_OgaLoadImage)(const char*, OgaImages**);
typedef OgaResult* (*PFN_OgaLoadImages)(const OgaStringArray*, OgaImages**);
typedef OgaResult* (*PFN_OgaLoadImagesFromBuffers)(const void**, const size_t*, size_t, OgaImages**);
typedef void (*PFN_OgaDestroyImages)(OgaImages*);
typedef OgaResult* (*PFN_OgaCreateMultiModalProcessor)(const OgaModel*, OgaMultiModalProcessor**);
typedef void (*PFN_OgaDestroyMultiModalProcessor)(OgaMultiModalProcessor*);
typedef OgaResult* (*PFN_OgaProcessorProcessImages)(const OgaMultiModalProcessor*, const char*, const OgaImages*, OgaNamedTensors**);
typedef void (*PFN_OgaDestroyNamedTensors)(OgaNamedTensors*);
typedef OgaResult* (*PFN_OgaCreateStringArray)(OgaStringArray**);
typedef void (*PFN_OgaDestroyStringArray)(OgaStringArray*);
typedef OgaResult* (*PFN_OgaStringArrayAddString)(OgaStringArray*, const char*);
typedef OgaResult* (*PFN_OgaProcessorProcessImagesAndPrompts)(const OgaMultiModalProcessor*, const OgaStringArray*, const OgaImages*, OgaNamedTensors**);


// Aggregated API table mirroring the pattern used by OrtApi in onnxruntime_go.
typedef struct GenAiApiTable {
	PFN_OgaCreateModel        CreateModel;
	PFN_OgaResultGetError     ResultGetError;
	PFN_OgaDestroyResult      DestroyResult;
	PFN_OgaDestroyModel       DestroyModel;
	PFN_OgaCreateTokenizer    CreateTokenizer;
	PFN_OgaDestroyTokenizer   DestroyTokenizer;
	PFN_OgaCreateTokenizerStream CreateTokenizerStream;
	PFN_OgaDestroyTokenizerStream DestroyTokenizerStream;
	PFN_OgaTokenizerApplyChatTemplate  ApplyChatTemplate;
	PFN_OgaDestroyString DestroyString;
	PFN_OgaCreateSequences CreateSequences;
	PFN_OgaDestroySequences DestroySequences;
	PFN_OgaTokenizerEncode TokenizerEncode;
	PFN_OgaCreateGenerator CreateGenerator;
	PFN_OgaDestroyGenerator DestroyGenerator;
	PFN_OgaCreateGeneratorParams CreateGeneratorParams;
	PFN_OgaDestroyGeneratorParams DestroyGeneratorParams;
	PFN_OgaGeneratorParamsSetSearchNumber GeneratorParamsSetSearchNumber;
	PFN_OgaGeneratorParamsSetGuidance GeneratorParamsSetGuidance;
	PFN_OgaGeneratorAppendTokenSequences GeneratorAppendTokenSequences;
	PFN_OgaGeneratorSetInputs GeneratorSetInputs;
	PFN_OgaGeneratorGenerateNextToken GeneratorGenerateNextToken;
	PFN_OgaGeneratorGetSequenceCount GeneratorGetSequenceCount;
	PFN_OgaGeneratorGetSequenceData GeneratorGetSequenceData;
	PFN_OgaTokenizerStreamDecode TokenizerStreamDecode;
	PFN_OgaGeneratorIsDone IsDone;
    PFN_OgaTokenizerGetEosTokenIds TokenizerGetEosTokenIds;
	// Config
	PFN_OgaCreateConfig CreateConfig;
	PFN_OgaConfigClearProviders ConfigClearProviders;
	PFN_OgaConfigAppendProvider ConfigAppendProvider;
	PFN_OgaConfigSetProviderOption ConfigSetProviderOption;
	PFN_OgaCreateModelFromConfig CreateModelFromConfig;
	PFN_OgaDestroyConfig DestroyConfig;
	// Multimodal
	PFN_OgaLoadImage LoadImage;
	PFN_OgaLoadImages LoadImages;
	PFN_OgaLoadImagesFromBuffers LoadImagesFromBuffers;
	PFN_OgaDestroyImages DestroyImages;
	PFN_OgaCreateMultiModalProcessor CreateMultiModalProcessor;
	PFN_OgaDestroyMultiModalProcessor DestroyMultiModalProcessor;
	PFN_OgaProcessorProcessImages ProcessorProcessImages;
	PFN_OgaDestroyNamedTensors DestroyNamedTensors;
	PFN_OgaCreateStringArray CreateStringArray;
	PFN_OgaDestroyStringArray DestroyStringArray;
	PFN_OgaStringArrayAddString StringArrayAddString;
	PFN_OgaProcessorProcessImagesAndPrompts ProcessorProcessImagesAndPrompts;
	// Extended / Release API
	PFN_OgaShutdown Shutdown;
	PFN_OgaSetTelemetryEnabled SetTelemetryEnabled;
	PFN_OgaCreateConfigFromPackageEp CreateConfigFromPackageEp;
	PFN_OgaTokenizerGetPadTokenId TokenizerGetPadTokenId;
	PFN_OgaTokenizerGetBotTokenId TokenizerGetBotTokenId;
	PFN_OgaTokenizerGetEotTokenId TokenizerGetEotTokenId;
	PFN_OgaTokenizerGetBorTokenId TokenizerGetBorTokenId;
	PFN_OgaTokenizerGetEorTokenId TokenizerGetEorTokenId;
} GenAiApiTable;

// Sets the global function pointer table. All pointers must be non-null.
// Returns 0 on success, non-zero on failure.
int SetGenAiApi(void* createModel,
				void* resultGetError,
				void* destroyResult,
				void* destroyModel,
				void* createTokenizer,
				void* destroyTokenizer,
				void* createTokenizerStream,
				void* destroyTokenizerStream,
				void* applyChatTemplate,
				void* destroyString,
				void* createSequences,
				void* destroySequences,
			void* tokenizerEncode,
			void* createGenerator,
			void* destroyGenerator,
			void* createGeneratorParams,
			void* destroyGeneratorParams,
		void* generatorParamsSetSearchNumber,
	void* generatorAppendTokenSequences,
	void* generatorSetInputs,
	void* generatorGenerateNextToken,
	void* generatorGetSequenceCount,
	void* generatorGetSequenceData,
	void* tokenizerStreamDecode,
	void* isDone,
    void* tokenizerGetEosTokenIds,
	// Config
	void* createConfig,
	void* configClearProviders,
	void* configAppendProvider,
	void* configSetProviderOption,
	void* createModelFromConfig,
	void* destroyConfig,
	// Multimodal
	void* loadImage,
	void* loadImages,
	void* loadImagesFromBuffers,
	void* destroyImages,
	void* createMultiModalProcessor,
	void* destroyMultiModalProcessor,
	void* processorProcessImages,
	void* destroyNamedTensors,
	void* createStringArray,
	void* destroyStringArray,
	void* stringArrayAddString,
	void* processorProcessImagesAndPrompts,
	// Guidance/constrained-generation
	void* generatorParamsSetGuidance,
	// Extended API
	void* shutdown,
	void* setTelemetryEnabled,
	void* createConfigFromPackageEp,
	void* tokenizerGetPadTokenId,
	void* tokenizerGetBotTokenId,
	void* tokenizerGetEotTokenId,
	void* tokenizerGetBorTokenId,
	void* tokenizerGetEorTokenId);

int SetGenAiEngineApi(void** symbols, size_t count);

// Returns non-zero if the API table is initialized.
int GenAiApiIsInitialized(void);

// Thin wrappers that call through the function pointer table.
OgaResult* CreateOgaModel(const char* config_path, OgaModel** out);
OgaResult* CreateOgaTokenizer(const OgaModel* model, OgaTokenizer** out);
OgaResult* CreateOgaTokenizerStream(const OgaTokenizer* tokenizer, OgaTokenizerStream** out);
OgaResult* CreateOgaSequences(OgaSequences** out);
OgaResult* CreateOgaGenerator(const OgaModel* model, const OgaGeneratorParams* generatorParams, OgaGenerator** out);
OgaResult* CreateOgaGeneratorParams(const OgaModel* model,OgaGeneratorParams** out);

const char* GetOgaResultErrorString(const OgaResult* result);
void DestroyOgaResult(OgaResult* result);
void DestroyOgaModel(OgaModel* model);
void DestroyOgaTokenizer(OgaTokenizer* tokenizer);
void DestroyOgaTokenizerStream(OgaTokenizerStream* tokenizerStream);
void DestroyOgaString(const char*);
void DestroyOgaSequences(OgaSequences* sequences);
void DestroyOgaGenerator(OgaGenerator* generator);
void DestroyOgaGeneratorParams(OgaGeneratorParams* generatorParams);
OgaResult* TokenizerEncode(const OgaTokenizer* tokenizer, const char* str, OgaSequences* sequences);

OgaResult* ApplyOgaTokenizerChatTemplate(const OgaTokenizer* tokenizer, const char* input, const char* param1, const char* param2, bool flag, const char** output);

OgaResult* GeneratorParamsSetSearchNumber(OgaGeneratorParams* generatorParams, const char* name, double searchNumber);
OgaResult* GeneratorParamsSetGuidance(OgaGeneratorParams* params, const char* type, const char* data, bool enable_ff_tokens);
OgaResult* GeneratorAppendTokenSequences(OgaGenerator* generator, OgaSequences* sequences);
OgaResult* GeneratorSetInputs(OgaGenerator* generator, const OgaNamedTensors* named_tensors);
OgaResult* GeneratorGenerateNextToken(OgaGenerator* generator);
size_t GeneratorGetSequenceCount(const OgaGenerator* generator, size_t sequence_index);
const int32_t* GeneratorGetSequenceData(const OgaGenerator* generator, size_t sequence_index);
OgaResult* TokenizerStreamDecode(OgaTokenizerStream* tokenizerStream, int32_t token, const char** output);
bool IsDone(const OgaGenerator* generator);
OgaResult* OgaTokenizerGetEosTokenIds(const OgaTokenizer* tokenizer, const int32_t** eos_token_ids, size_t* token_count);

// Config thin wrappers
OgaResult* CreateOgaConfig(const char* config_path, OgaConfig** out);
OgaResult* OgaConfigClearProviders(OgaConfig* config);
OgaResult* OgaConfigAppendProvider(OgaConfig* config, const char* provider);
OgaResult* OgaConfigSetProviderOption(OgaConfig* config, const char* provider, const char* key, const char* value);
OgaResult* CreateOgaModelFromConfig(const OgaConfig* config, OgaModel** out);
void DestroyOgaConfig(OgaConfig* config);

// Multimodal thin wrappers
OgaResult* LoadOgaImage(const char* image_path, OgaImages** out);
OgaResult* LoadOgaImages(const OgaStringArray* image_paths, OgaImages** out);
OgaResult* LoadOgaImagesFromBuffers(const void** image_data, const size_t* image_data_sizes, size_t count, OgaImages** out);
void DestroyOgaImages(OgaImages* images);
OgaResult* CreateOgaMultiModalProcessor(const OgaModel* model, OgaMultiModalProcessor** out);
void DestroyOgaMultiModalProcessor(OgaMultiModalProcessor* processor);
OgaResult* ProcessOgaImages(const OgaMultiModalProcessor* processor, const char* prompt, const OgaImages* images, OgaNamedTensors** out);
void DestroyOgaNamedTensors(OgaNamedTensors* tensors);
OgaResult* CreateOgaStringArray(OgaStringArray** out);
void DestroyOgaStringArray(OgaStringArray* string_array);
OgaResult* AddStringToOgaStringArray(OgaStringArray* string_array, const char* str);
OgaResult* ProcessOgaImagesAndPrompts(const OgaMultiModalProcessor* processor,  const OgaStringArray* prompts, const OgaImages* images, OgaNamedTensors** out);

OgaResult* EngineCreate(OgaModel* model, OgaEngine** out);
void EngineDestroy(OgaEngine* engine);
OgaResult* EngineCreateEventBuffer(OgaEngine* engine, size_t capacity, OgaEngineEventBuffer** out);
void EngineDestroyEventBuffer(OgaEngineEventBuffer* buffer);
OgaResult* EngineRun(OgaEngine* engine, OgaEngineEventBuffer* buffer);
size_t EventBufferGetCount(const OgaEngineEventBuffer* buffer);
const OgaEngineEvent* EventBufferGet(const OgaEngineEventBuffer* buffer, size_t index);
OgaResult* EngineEventGetFlags(const OgaEngineEvent* event, OgaEngineEventFlags* out);
OgaResult* EngineEventGetTurnId(const OgaEngineEvent* event, uint64_t* out);
OgaResult* EngineEventGetToken(const OgaEngineEvent* event, int32_t* out);
OgaResult* EngineEventGetFinishReason(const OgaEngineEvent* event, OgaFinishReason* out);
OgaResult* EngineEventGetMatchedStopStringIndex(const OgaEngineEvent* event, int32_t* out);
OgaResult* EngineEventGetErrorCode(const OgaEngineEvent* event, OgaErrorCode* out);
OgaResult* EngineEventGetUsage(const OgaEngineEvent* event, const OgaTurnUsage** out);
OgaResult* TurnUsageGetPromptTokens(const OgaTurnUsage* usage, uint64_t* out);
OgaResult* TurnUsageGetGeneratedTokens(const OgaTurnUsage* usage, uint64_t* out);
OgaResult* TurnUsageGetCachedPromptTokens(const OgaTurnUsage* usage, uint64_t* out);
OgaResult* EngineHasPendingRequests(OgaEngine* engine, bool* out);
OgaResult* EngineCreateRequest(OgaEngine* engine, const OgaRequestOptions* options, OgaRequest** out);
OgaResult* CreateRequestOptions(OgaRequestOptions** out);
void DestroyRequestOptions(OgaRequestOptions* options);
OgaResult* RequestOptionsSetMaxSessionTokens(OgaRequestOptions* options, uint64_t max_tokens);
OgaResult* RequestCreateTurnOptions(OgaRequest* request, OgaTurnOptions** out);
void DestroyTurnOptions(OgaTurnOptions* options);
OgaResult* TurnOptionsSetMaxGeneratedTokens(OgaTurnOptions* options, uint64_t value);
OgaResult* TurnOptionsSetMinGeneratedTokens(OgaTurnOptions* options, uint64_t value);
OgaResult* TurnOptionsSetDoSample(OgaTurnOptions* options, bool value);
OgaResult* TurnOptionsSetTemperature(OgaTurnOptions* options, float value);
OgaResult* TurnOptionsSetTopP(OgaTurnOptions* options, float value);
OgaResult* TurnOptionsSetTopK(OgaTurnOptions* options, int32_t value);
OgaResult* TurnOptionsSetRepetitionPenalty(OgaTurnOptions* options, float value);
OgaResult* TurnOptionsSetNoRepeatNgramSize(OgaTurnOptions* options, int32_t value);
OgaResult* TurnOptionsSetSeed(OgaTurnOptions* options, uint64_t value);
OgaResult* TurnOptionsClearSeed(OgaTurnOptions* options);
OgaResult* TurnOptionsSetStopStrings(OgaTurnOptions* options, const OgaStringArray* strings);
OgaResult* TurnOptionsSetGuidance(OgaTurnOptions* options, const char* type, const char* data);
OgaResult* TurnOptionsClearGuidance(OgaTurnOptions* options);
OgaResult* TurnOptionsReset(OgaTurnOptions* options);
OgaResult* RequestBeginTurn(OgaRequest* request, const OgaTurnOptions* options, const int32_t* ids, uint64_t count, uint64_t* turn_id);
OgaResult* RequestCancelTurn(OgaRequest* request, uint64_t turn_id, bool* cancelled);
OgaResult* RequestRewindToStartOfTurn(OgaRequest* request, uint64_t turn_id);
OgaResult* RequestClose(OgaRequest* request);
OgaResult* RequestSetDraftTokens(OgaRequest* request, const OgaSequences* tokens);
void DestroyRequest(OgaRequest* request);
OgaResult* EngineMaxDraftTokensPerProposal(const OgaEngine* engine, size_t* out);
OgaResult* AppendTokenSequence(const int32_t* ids, size_t count, OgaSequences* sequences);
size_t SequencesCount(const OgaSequences* sequences);
size_t SequencesGetSequenceCount(const OgaSequences* sequences, size_t index);
const int32_t* SequencesGetSequenceData(const OgaSequences* sequences, size_t index);
OgaResult* TokenizerDecode(const OgaTokenizer* tokenizer, const int32_t* ids, size_t count, const char** out);
OgaResult* TokenizerToTokenId(const OgaTokenizer* tokenizer, const char* text, int32_t* out);
OgaResult* TokenizerEncodeBatch(const OgaTokenizer* tokenizer, const char** strings, size_t count, OgaTensor** out);
OgaResult* TokenizerDecodeBatch(const OgaTokenizer* tokenizer, const OgaTensor* tensor, OgaStringArray** out);
OgaResult* CreateTensorFromBuffer(void* data, const int64_t* shape, size_t rank, OgaElementType type, OgaTensor** out);
void DestroyTensor(OgaTensor* tensor);
OgaResult* TensorGetType(OgaTensor* tensor, OgaElementType* out);
OgaResult* TensorGetShapeRank(OgaTensor* tensor, size_t* out);
OgaResult* TensorGetShape(OgaTensor* tensor, int64_t* shape, size_t rank);
OgaResult* TensorGetData(OgaTensor* tensor, void** out);
OgaResult* StringArrayGetCount(const OgaStringArray* strings, size_t* out);
OgaResult* StringArrayGetString(const OgaStringArray* strings, size_t index, const char** out);
OgaResult* SetLogBool(const char* name, bool value);
OgaResult* SetLogString(const char* name, const char* value);
OgaResult* SetCurrentGpuDeviceId(int id);
OgaResult* GetCurrentGpuDeviceId(int* id);
OgaResult* CreateRuntimeSettings(OgaRuntimeSettings** out);
void DestroyRuntimeSettings(OgaRuntimeSettings* settings);
OgaResult* RuntimeSettingsSetHandle(OgaRuntimeSettings* settings, const char* name, void* handle);
OgaResult* CreateModelWithRuntimeSettings(const char* path, const OgaRuntimeSettings* settings, OgaModel** out);
OgaResult* LoadAudio(const char* path, OgaAudios** out);
OgaResult* LoadAudios(const OgaStringArray* paths, OgaAudios** out);
OgaResult* LoadAudiosFromBuffers(const void** data, const size_t* sizes, size_t count, OgaAudios** out);
void DestroyAudios(OgaAudios* audios);
OgaResult* ProcessAudios(const OgaMultiModalProcessor* processor, const char* prompt, const OgaAudios* audios, OgaNamedTensors** out);
OgaResult* ProcessAudiosAndPrompts(const OgaMultiModalProcessor* processor, const OgaStringArray* prompts, const OgaAudios* audios, OgaNamedTensors** out);
OgaResult* ProcessImagesAndAudios(const OgaMultiModalProcessor* processor, const char* prompt, const OgaImages* images, const OgaAudios* audios, OgaNamedTensors** out);
OgaResult* ProcessImagesAndAudiosAndPrompts(const OgaMultiModalProcessor* processor, const OgaStringArray* prompts, const OgaImages* images, const OgaAudios* audios, OgaNamedTensors** out);
OgaResult* CreateAdapters(const OgaModel* model, OgaAdapters** out);
void DestroyAdapters(OgaAdapters* adapters);
OgaResult* LoadAdapter(OgaAdapters* adapters, const char* path, const char* name);
OgaResult* UnloadAdapter(OgaAdapters* adapters, const char* name);
OgaResult* SetActiveAdapter(OgaGenerator* generator, OgaAdapters* adapters, const char* name);
OgaResult* CreateMtpGenerator(const OgaModel* main_model, const OgaModel* mtp_model, const OgaGeneratorParams* params, OgaMtpGenerator** out);
OgaResult* MtpGeneratorAppendTokens(OgaMtpGenerator* generator, const int32_t* ids, size_t count);
OgaResult* MtpGeneratorGenerateNextToken(OgaMtpGenerator* generator);
OgaResult* MtpGeneratorReset(OgaMtpGenerator* generator);
bool MtpGeneratorIsDone(const OgaMtpGenerator* generator);
size_t MtpGeneratorGetSequenceCount(const OgaMtpGenerator* generator);
const int32_t* MtpGeneratorGetSequenceData(const OgaMtpGenerator* generator);
size_t MtpGeneratorGetForwardCount(const OgaMtpGenerator* generator);
size_t MtpGeneratorGetAcceptCount(const OgaMtpGenerator* generator);
size_t MtpGeneratorGetTrialCount(const OgaMtpGenerator* generator);
OgaResult* MtpGeneratorGetSpeculativeStats(const OgaMtpGenerator* generator, OgaSpeculativeStats** out);
void DestroyMtpGenerator(OgaMtpGenerator* generator);
OgaResult* GeneratorParamsSetSearchBool(OgaGeneratorParams* params, const char* name, bool value);
OgaResult* GeneratorSetModelInput(OgaGenerator* generator, const char* name, OgaTensor* tensor);
OgaResult* GeneratorAppendTokens(OgaGenerator* generator, const int32_t* ids, size_t count);
size_t GeneratorTokenCount(const OgaGenerator* generator);
OgaResult* GeneratorGetNextTokens(const OgaGenerator* generator, const int32_t** out, size_t* count);
OgaResult* GeneratorSetRuntimeOption(OgaGenerator* generator, const char* key, const char* value);
OgaResult* GeneratorRewindTo(OgaGenerator* generator, size_t length);
OgaResult* GeneratorSnapshotState(OgaGenerator* generator);
OgaResult* GeneratorSetHiddenStates(OgaGenerator* generator, OgaTensor* tensor);
OgaResult* GeneratorGetInput(const OgaGenerator* generator, const char* name, OgaTensor** out);
OgaResult* GeneratorGetOutput(const OgaGenerator* generator, const char* name, OgaTensor** out);
OgaResult* GeneratorGetLogits(OgaGenerator* generator, OgaTensor** out);
OgaResult* GeneratorSetLogits(OgaGenerator* generator, OgaTensor* tensor);
OgaResult* GeneratorGetSpeculativeStats(const OgaGenerator* generator, OgaSpeculativeStats** out);
void DestroySpeculativeStats(OgaSpeculativeStats* stats);
OgaResult* SpeculativeStatsGetCount(const OgaSpeculativeStats* stats, const char* name, uint64_t* out);
OgaResult* SpeculativeStatsGetAcceptanceLengthCount(const OgaSpeculativeStats* stats, size_t length, uint64_t* out);
OgaResult* SpeculativeStatsGetAcceptanceLengthHistogramSize(const OgaSpeculativeStats* stats, size_t* out);
OgaResult* SpeculativeStatsGetNumber(const OgaSpeculativeStats* stats, const char* name, double* out);
OgaResult* SpeculativeStatsGetBool(const OgaSpeculativeStats* stats, const char* name, bool* out);
OgaResult* TokenizerUpdateOptions(OgaTokenizer* tokenizer, const char* const* keys, const char* const* values, size_t count);
OgaResult* SetLogCallback(bool enabled);
OgaResult* EngineEventGetRequest(const OgaEngineEvent* event, const OgaRequest** out);
OgaResult* EngineGetSpeculativeStats(const OgaEngine* engine, OgaSpeculativeStats** out);
OgaResult* EngineGetCapabilities(const OgaEngine* engine, OgaEngineCapabilities** out);
size_t EngineCapabilitiesGetConfiguredMaxBatchSize(const OgaEngineCapabilities* capabilities);
size_t EngineCapabilitiesGetMaxScheduledTokens(const OgaEngineCapabilities* capabilities);
uint64_t EngineCapabilitiesGetMaxRequestLength(const OgaEngineCapabilities* capabilities);
void DestroyEngineCapabilities(OgaEngineCapabilities* capabilities);


void OgaShutdown(void);
void OgaSetTelemetryEnabled(bool enabled);
OgaResult* OgaCreateConfigFromPackageEp(const char* config_path, const char* ep, OgaConfig** out);
OgaResult* OgaTokenizerGetPadTokenId(const OgaTokenizer* tokenizer, int32_t* token_id);
OgaResult* OgaTokenizerGetBotTokenId(const OgaTokenizer* tokenizer, int32_t* token_id);
OgaResult* OgaTokenizerGetEotTokenId(const OgaTokenizer* tokenizer, int32_t* token_id);
OgaResult* OgaTokenizerGetBorTokenId(const OgaTokenizer* tokenizer, int32_t* token_id);
OgaResult* OgaTokenizerGetEorTokenId(const OgaTokenizer* tokenizer, int32_t* token_id);

#ifdef __cplusplus
} // extern "C"
#endif

#endif  // ORT_GENAI_WRAPPER_H
