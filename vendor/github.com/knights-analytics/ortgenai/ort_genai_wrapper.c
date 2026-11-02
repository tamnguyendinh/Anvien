#include "ort_genai_wrapper.h"
#include <string.h>

extern void goLogCallback(char* message, size_t length);
static void logCallbackBridge(const char* message, size_t length) {
	goLogCallback((char*)message, length);
}

static GenAiApiTable g_api = {0}; // api table
static int g_initialized = 0;
static void* g_engine_api[123] = {0};
static int g_engine_initialized = 0;

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
	void* shutdown,
	void* setTelemetryEnabled,
	void* createConfigFromPackageEp,
	void* tokenizerGetPadTokenId,
	void* tokenizerGetBotTokenId,
	void* tokenizerGetEotTokenId,
	void* tokenizerGetBorTokenId,
	void* tokenizerGetEorTokenId) {
	if (g_initialized) return 0; // already initialized
	// Validate all required pointers (header comment: all must be non-null)
	if (!createModel || !resultGetError || !destroyResult || !destroyModel ||
		!createTokenizer || !destroyTokenizer || !createTokenizerStream || !destroyTokenizerStream ||
		!applyChatTemplate || !destroyString || !createSequences || !destroySequences || !tokenizerEncode ||
		!createGenerator || !destroyGenerator || !createGeneratorParams || !destroyGeneratorParams ||
		!generatorParamsSetSearchNumber || !generatorAppendTokenSequences || !generatorSetInputs || !generatorGenerateNextToken ||
		!generatorGetSequenceCount || !generatorGetSequenceData || !tokenizerStreamDecode || !isDone || !tokenizerGetEosTokenIds ||
		// Config
		!createConfig || !configClearProviders || !configAppendProvider || !configSetProviderOption || !createModelFromConfig || !destroyConfig ||
		// Multimodal
		!loadImage || !loadImages || !loadImagesFromBuffers || !destroyImages ||
		!createMultiModalProcessor || !destroyMultiModalProcessor || !processorProcessImages ||
		!destroyNamedTensors || !createStringArray || !destroyStringArray || !stringArrayAddString || !processorProcessImagesAndPrompts ||
		// Extended API
		!shutdown || !setTelemetryEnabled || !createConfigFromPackageEp ||
		!tokenizerGetPadTokenId || !tokenizerGetBotTokenId || !tokenizerGetEotTokenId ||
		!tokenizerGetBorTokenId || !tokenizerGetEorTokenId) {
		return 1;
	}
	g_api.CreateModel = (PFN_OgaCreateModel) createModel;
	g_api.ResultGetError = (PFN_OgaResultGetError) resultGetError;
	g_api.DestroyResult = (PFN_OgaDestroyResult) destroyResult;
	g_api.DestroyModel = (PFN_OgaDestroyModel) destroyModel;
	g_api.CreateTokenizer = (PFN_OgaCreateTokenizer) createTokenizer;
	g_api.DestroyTokenizer = (PFN_OgaDestroyTokenizer) destroyTokenizer;
	g_api.CreateTokenizerStream = (PFN_OgaCreateTokenizerStream) createTokenizerStream;
	g_api.DestroyTokenizerStream = (PFN_OgaDestroyTokenizerStream) destroyTokenizerStream;
	g_api.ApplyChatTemplate = (PFN_OgaTokenizerApplyChatTemplate) applyChatTemplate;
	g_api.DestroyString = (PFN_OgaDestroyString) destroyString;
	g_api.CreateSequences = (PFN_OgaCreateSequences) createSequences;
	g_api.DestroySequences = (PFN_OgaDestroySequences) destroySequences;
	g_api.TokenizerEncode = (PFN_OgaTokenizerEncode) tokenizerEncode;
	g_api.CreateGenerator = (PFN_OgaCreateGenerator) createGenerator;
	g_api.DestroyGenerator = (PFN_OgaDestroyGenerator) destroyGenerator;
	g_api.CreateGeneratorParams = (PFN_OgaCreateGeneratorParams) createGeneratorParams;
	g_api.DestroyGeneratorParams = (PFN_OgaDestroyGeneratorParams) destroyGeneratorParams;
	g_api.GeneratorParamsSetSearchNumber = (PFN_OgaGeneratorParamsSetSearchNumber) generatorParamsSetSearchNumber;
	g_api.GeneratorAppendTokenSequences = (PFN_OgaGeneratorAppendTokenSequences) generatorAppendTokenSequences;
	g_api.GeneratorSetInputs = (PFN_OgaGeneratorSetInputs) generatorSetInputs;
	g_api.GeneratorGenerateNextToken = (PFN_OgaGeneratorGenerateNextToken) generatorGenerateNextToken;
	g_api.GeneratorGetSequenceCount = (PFN_OgaGeneratorGetSequenceCount) generatorGetSequenceCount;
	g_api.GeneratorGetSequenceData = (PFN_OgaGeneratorGetSequenceData) generatorGetSequenceData;
	g_api.TokenizerStreamDecode = (PFN_OgaTokenizerStreamDecode) tokenizerStreamDecode;
	g_api.IsDone = (PFN_OgaGeneratorIsDone) isDone;
	g_api.TokenizerGetEosTokenIds = (PFN_OgaTokenizerGetEosTokenIds) tokenizerGetEosTokenIds;
	// Config
	g_api.CreateConfig = (PFN_OgaCreateConfig) createConfig;
	g_api.ConfigClearProviders = (PFN_OgaConfigClearProviders) configClearProviders;
	g_api.ConfigAppendProvider = (PFN_OgaConfigAppendProvider) configAppendProvider;
	g_api.ConfigSetProviderOption = (PFN_OgaConfigSetProviderOption) configSetProviderOption;
	g_api.CreateModelFromConfig = (PFN_OgaCreateModelFromConfig) createModelFromConfig;
	g_api.DestroyConfig = (PFN_OgaDestroyConfig) destroyConfig;
	// Multimodal
	g_api.LoadImage = (PFN_OgaLoadImage) loadImage;
	g_api.LoadImages = (PFN_OgaLoadImages) loadImages;
	g_api.LoadImagesFromBuffers = (PFN_OgaLoadImagesFromBuffers) loadImagesFromBuffers;
	g_api.DestroyImages = (PFN_OgaDestroyImages) destroyImages;
	g_api.CreateMultiModalProcessor = (PFN_OgaCreateMultiModalProcessor) createMultiModalProcessor;
	g_api.DestroyMultiModalProcessor = (PFN_OgaDestroyMultiModalProcessor) destroyMultiModalProcessor;
	g_api.ProcessorProcessImages = (PFN_OgaProcessorProcessImages) processorProcessImages;
	g_api.DestroyNamedTensors = (PFN_OgaDestroyNamedTensors) destroyNamedTensors;
	g_api.CreateStringArray = (PFN_OgaCreateStringArray) createStringArray;
	g_api.DestroyStringArray = (PFN_OgaDestroyStringArray) destroyStringArray;
	g_api.StringArrayAddString = (PFN_OgaStringArrayAddString) stringArrayAddString;
	g_api.ProcessorProcessImagesAndPrompts = (PFN_OgaProcessorProcessImagesAndPrompts) processorProcessImagesAndPrompts;
	if (generatorParamsSetGuidance) g_api.GeneratorParamsSetGuidance = (PFN_OgaGeneratorParamsSetGuidance) generatorParamsSetGuidance;
	g_api.Shutdown = (PFN_OgaShutdown) shutdown;
	g_api.SetTelemetryEnabled = (PFN_OgaSetTelemetryEnabled) setTelemetryEnabled;
	g_api.CreateConfigFromPackageEp = (PFN_OgaCreateConfigFromPackageEp) createConfigFromPackageEp;
	g_api.TokenizerGetPadTokenId = (PFN_OgaTokenizerGetPadTokenId) tokenizerGetPadTokenId;
	g_api.TokenizerGetBotTokenId = (PFN_OgaTokenizerGetBotTokenId) tokenizerGetBotTokenId;
	g_api.TokenizerGetEotTokenId = (PFN_OgaTokenizerGetEotTokenId) tokenizerGetEotTokenId;
	g_api.TokenizerGetBorTokenId = (PFN_OgaTokenizerGetBorTokenId) tokenizerGetBorTokenId;
	g_api.TokenizerGetEorTokenId = (PFN_OgaTokenizerGetEorTokenId) tokenizerGetEorTokenId;
	g_initialized = 1;
	return 0;
}

int GenAiApiIsInitialized(void) { return g_initialized; }

OgaResult* CreateOgaModel(const char* config_path, OgaModel** out) {
	if (!g_initialized || !g_api.CreateModel) return NULL;
	return g_api.CreateModel(config_path, out);
}

OgaResult* CreateOgaTokenizer(const OgaModel* model, OgaTokenizer** out) {
	if (!g_initialized || !g_api.CreateTokenizer) return NULL;
	return g_api.CreateTokenizer(model, out);
}

OgaResult* CreateOgaTokenizerStream(const OgaTokenizer* tokenizer, OgaTokenizerStream** out) {
	if (!g_initialized || !g_api.CreateTokenizerStream) return NULL;
	return g_api.CreateTokenizerStream(tokenizer, out);
}

const char* GetOgaResultErrorString(const OgaResult* result) {
	if (!g_initialized || !g_api.ResultGetError) return "GenAI API not initialized";
	return g_api.ResultGetError(result);
}

void DestroyOgaResult(OgaResult* result) {
	if (!result) return;
	if (!g_initialized || !g_api.DestroyResult) return;
	g_api.DestroyResult(result);
}

void DestroyOgaModel(OgaModel* model) {
	if (!model) return;
	if (!g_initialized || !g_api.DestroyModel) return;
	g_api.DestroyModel(model);
}

void DestroyOgaTokenizer(OgaTokenizer* tokenizer) {
	if (!tokenizer) return;
	if (!g_initialized || !g_api.DestroyTokenizer) return;
	g_api.DestroyTokenizer(tokenizer);
}

void DestroyOgaTokenizerStream(OgaTokenizerStream* tokenizerStream) {
	if (!tokenizerStream) return;
	if (!g_initialized || !g_api.DestroyTokenizerStream) return;
	g_api.DestroyTokenizerStream(tokenizerStream);
}

void DestroyOgaString(const char* str) {
	if (!str) return;
	if (!g_initialized || !g_api.DestroyString) return;
	g_api.DestroyString(str);
}

OgaResult* ApplyOgaTokenizerChatTemplate(const OgaTokenizer* tokenizer, const char* input, const char* param1, const char* param2, bool flag, const char** output) {
	if (!g_initialized || !g_api.ApplyChatTemplate) return NULL;
	return g_api.ApplyChatTemplate(tokenizer, input, param1, param2, flag, output);
}

OgaResult* CreateOgaSequences(OgaSequences** out) {
	if (!g_initialized || !g_api.CreateSequences) return NULL;
	return g_api.CreateSequences(out);
}

void DestroyOgaSequences(OgaSequences* sequences) {
	if (!sequences) return;
	if (!g_initialized || !g_api.DestroySequences) return;
	g_api.DestroySequences(sequences);
}

OgaResult* TokenizerEncode(const OgaTokenizer* tokenizer, const char* str, OgaSequences* sequences) {
	if (!g_initialized || !g_api.TokenizerEncode) return NULL;
	return g_api.TokenizerEncode(tokenizer, str, sequences);
}

OgaResult* CreateOgaGenerator(const OgaModel* model, const OgaGeneratorParams* generatorParams, OgaGenerator** out) {
	if (!g_initialized || !g_api.CreateGenerator) return NULL;
	return g_api.CreateGenerator(model, generatorParams, out);
}

void DestroyOgaGenerator(OgaGenerator* generator) {
	if (!generator) return;
	if (!g_initialized || !g_api.DestroyGenerator) return;
	g_api.DestroyGenerator(generator);
}

OgaResult* CreateOgaGeneratorParams(const OgaModel* model,OgaGeneratorParams** out) {
	if (!g_initialized || !g_api.CreateGeneratorParams) return NULL;
	return g_api.CreateGeneratorParams(model, out);
}

void DestroyOgaGeneratorParams(OgaGeneratorParams* generatorParams) {
	if (!generatorParams) return;
	if (!g_initialized || !g_api.DestroyGeneratorParams) return;
	g_api.DestroyGeneratorParams(generatorParams);
}

OgaResult* GeneratorParamsSetSearchNumber(OgaGeneratorParams* generatorParams, const char* name, double searchNumber) {
    if (!g_initialized || !g_api.GeneratorParamsSetSearchNumber) return NULL;
    return g_api.GeneratorParamsSetSearchNumber(generatorParams, name, searchNumber);
}

OgaResult* GeneratorParamsSetGuidance(OgaGeneratorParams* params, const char* type, const char* data, bool enable_ff_tokens) {
	if (!g_initialized || !g_api.GeneratorParamsSetGuidance) return NULL;
	return g_api.GeneratorParamsSetGuidance(params, type, data, enable_ff_tokens);
}

OgaResult* GeneratorAppendTokenSequences(OgaGenerator* generator, OgaSequences* sequences) {
	if (!g_initialized || !g_api.GeneratorAppendTokenSequences) return NULL;
	return g_api.GeneratorAppendTokenSequences(generator, sequences);
}

OgaResult* GeneratorSetInputs(OgaGenerator* generator, const OgaNamedTensors* named_tensors) {
	if (!g_initialized || !g_api.GeneratorSetInputs) return NULL;
	return g_api.GeneratorSetInputs(generator, named_tensors);
}

OgaResult* GeneratorGenerateNextToken(OgaGenerator* generator) {
	if (!g_initialized || !g_api.GeneratorGenerateNextToken) return NULL;
	return g_api.GeneratorGenerateNextToken(generator);
}

size_t GeneratorGetSequenceCount(const OgaGenerator* generator, size_t sequence_index) {
	if (!g_initialized || !g_api.GeneratorGetSequenceCount) return 0;
	return g_api.GeneratorGetSequenceCount(generator, sequence_index);
}

const int32_t* GeneratorGetSequenceData(const OgaGenerator* generator, size_t sequence_index) {
	if (!g_initialized || !g_api.GeneratorGetSequenceData) return NULL;
	return g_api.GeneratorGetSequenceData(generator, sequence_index);
}

OgaResult* TokenizerStreamDecode(OgaTokenizerStream* tokenizerStream, int32_t token, const char** output) {
	if (!g_initialized || !g_api.TokenizerStreamDecode) return NULL;
	return g_api.TokenizerStreamDecode(tokenizerStream, token, output);
}

bool IsDone(const OgaGenerator* generator) {
	if (!g_initialized || !g_api.IsDone) return false;
	return g_api.IsDone(generator);
}

OgaResult* OgaTokenizerGetEosTokenIds(const OgaTokenizer* tokenizer, const int32_t** eos_token_ids, size_t* token_count) {
    if (!g_initialized || !g_api.TokenizerGetEosTokenIds) return NULL;
    return g_api.TokenizerGetEosTokenIds(tokenizer, eos_token_ids, token_count);
}

// Config thin wrappers
OgaResult* CreateOgaConfig(const char* config_path, OgaConfig** out) {
	if (!g_initialized || !g_api.CreateConfig) return NULL;
	return g_api.CreateConfig(config_path, out);
}

OgaResult* OgaConfigClearProviders(OgaConfig* config) {
	if (!g_initialized || !g_api.ConfigClearProviders) return NULL;
	return g_api.ConfigClearProviders(config);
}

OgaResult* OgaConfigAppendProvider(OgaConfig* config, const char* provider) {
	if (!g_initialized || !g_api.ConfigAppendProvider) return NULL;
	return g_api.ConfigAppendProvider(config, provider);
}

OgaResult* OgaConfigSetProviderOption(OgaConfig* config, const char* provider, const char* key, const char* value) {
	if (!g_initialized || !g_api.ConfigSetProviderOption) return NULL;
	return g_api.ConfigSetProviderOption(config, provider, key, value);
}

OgaResult* CreateOgaModelFromConfig(const OgaConfig* config, OgaModel** out) {
	if (!g_initialized || !g_api.CreateModelFromConfig) return NULL;
	return g_api.CreateModelFromConfig(config, out);
}

void DestroyOgaConfig(OgaConfig* config) {
	if (!config) return;
	if (!g_initialized || !g_api.DestroyConfig) return;
	g_api.DestroyConfig(config);
}

// Multimodal thin wrappers
OgaResult* LoadOgaImage(const char* image_path, OgaImages** out) {
	if (!g_initialized || !g_api.LoadImage) return NULL;
	return g_api.LoadImage(image_path, out);
}

OgaResult* LoadOgaImages(const OgaStringArray* image_paths, OgaImages** out) {
	if (!g_initialized || !g_api.LoadImages) return NULL;
	return g_api.LoadImages(image_paths, out);
}

OgaResult* LoadOgaImagesFromBuffers(const void** image_data, const size_t* image_data_sizes, size_t count, OgaImages** out) {
	if (!g_initialized || !g_api.LoadImagesFromBuffers) return NULL;
	return g_api.LoadImagesFromBuffers(image_data, image_data_sizes, count, out);
}

void DestroyOgaImages(OgaImages* images) {
	if (!images) return;
	if (!g_initialized || !g_api.DestroyImages) return;
	g_api.DestroyImages(images);
}

OgaResult* CreateOgaMultiModalProcessor(const OgaModel* model, OgaMultiModalProcessor** out) {
	if (!g_initialized || !g_api.CreateMultiModalProcessor) return NULL;
	return g_api.CreateMultiModalProcessor(model, out);
}

void DestroyOgaMultiModalProcessor(OgaMultiModalProcessor* processor) {
	if (!processor) return;
	if (!g_initialized || !g_api.DestroyMultiModalProcessor) return;
	g_api.DestroyMultiModalProcessor(processor);
}

OgaResult* ProcessOgaImages(const OgaMultiModalProcessor* processor, const char* prompt, const OgaImages* images, OgaNamedTensors** out) {
	if (!g_initialized || !g_api.ProcessorProcessImages) return NULL;
	return g_api.ProcessorProcessImages(processor, prompt, images, out);
}

void DestroyOgaNamedTensors(OgaNamedTensors* tensors) {
	if (!tensors) return;
	if (!g_initialized || !g_api.DestroyNamedTensors) return;
	g_api.DestroyNamedTensors(tensors);
}

OgaResult* CreateOgaStringArray(OgaStringArray** out) {
	if (!g_initialized || !g_api.CreateStringArray) return NULL;
	return g_api.CreateStringArray(out);
}

void DestroyOgaStringArray(OgaStringArray* string_array) {
	if (!string_array) return;
	if (!g_initialized || !g_api.DestroyStringArray) return;
	g_api.DestroyStringArray(string_array);
}

OgaResult* AddStringToOgaStringArray(OgaStringArray* string_array, const char* str) {
	if (!g_initialized || !g_api.StringArrayAddString) return NULL;
	return g_api.StringArrayAddString(string_array, str);
}

OgaResult* ProcessOgaImagesAndPrompts(const OgaMultiModalProcessor* processor,  const OgaStringArray* prompts, const OgaImages* images, OgaNamedTensors** out) {
	if (!g_initialized || !g_api.ProcessorProcessImagesAndPrompts) return NULL;
	return g_api.ProcessorProcessImagesAndPrompts(processor, prompts, images, out);
}

int SetGenAiEngineApi(void** symbols, size_t count) {
	if (g_engine_initialized) return 0;
	if (!symbols || count != 123) return 1;
	for (size_t i = 0; i < count; ++i) {
		if (!symbols[i]) return 1;
		g_engine_api[i] = symbols[i];
	}
	g_engine_initialized = 1;
	return 0;
}

#define ENGINE_API(index, type) ((type)g_engine_api[index])
OgaResult* EngineCreate(OgaModel* m, OgaEngine** o) { return g_engine_initialized ? ENGINE_API(0, OgaResult* (*)(OgaModel*, OgaEngine**))(m, o) : NULL; }
void EngineDestroy(OgaEngine* e) { if (g_engine_initialized && e) ENGINE_API(1, void (*)(OgaEngine*))(e); }
OgaResult* EngineCreateEventBuffer(OgaEngine* e, size_t n, OgaEngineEventBuffer** o) { return g_engine_initialized ? ENGINE_API(2, OgaResult* (*)(OgaEngine*, size_t, OgaEngineEventBuffer**))(e, n, o) : NULL; }
void EngineDestroyEventBuffer(OgaEngineEventBuffer* b) { if (g_engine_initialized && b) ENGINE_API(3, void (*)(OgaEngineEventBuffer*))(b); }
OgaResult* EngineRun(OgaEngine* e, OgaEngineEventBuffer* b) { return g_engine_initialized ? ENGINE_API(4, OgaResult* (*)(OgaEngine*, OgaEngineEventBuffer*))(e, b) : NULL; }
size_t EventBufferGetCount(const OgaEngineEventBuffer* b) { return g_engine_initialized ? ENGINE_API(5, size_t (*)(const OgaEngineEventBuffer*))(b) : 0; }
const OgaEngineEvent* EventBufferGet(const OgaEngineEventBuffer* b, size_t i) { return g_engine_initialized ? ENGINE_API(6, const OgaEngineEvent* (*)(const OgaEngineEventBuffer*, size_t))(b, i) : NULL; }
OgaResult* EngineEventGetFlags(const OgaEngineEvent* e, OgaEngineEventFlags* o) { return g_engine_initialized ? ENGINE_API(7, OgaResult* (*)(const OgaEngineEvent*, OgaEngineEventFlags*))(e, o) : NULL; }
OgaResult* EngineEventGetTurnId(const OgaEngineEvent* e, uint64_t* o) { return g_engine_initialized ? ENGINE_API(8, OgaResult* (*)(const OgaEngineEvent*, uint64_t*))(e, o) : NULL; }
OgaResult* EngineEventGetToken(const OgaEngineEvent* e, int32_t* o) { return g_engine_initialized ? ENGINE_API(9, OgaResult* (*)(const OgaEngineEvent*, int32_t*))(e, o) : NULL; }
OgaResult* EngineEventGetFinishReason(const OgaEngineEvent* e, OgaFinishReason* o) { return g_engine_initialized ? ENGINE_API(10, OgaResult* (*)(const OgaEngineEvent*, OgaFinishReason*))(e, o) : NULL; }
OgaResult* EngineEventGetMatchedStopStringIndex(const OgaEngineEvent* e, int32_t* o) { return g_engine_initialized ? ENGINE_API(11, OgaResult* (*)(const OgaEngineEvent*, int32_t*))(e, o) : NULL; }
OgaResult* EngineEventGetErrorCode(const OgaEngineEvent* e, OgaErrorCode* o) { return g_engine_initialized ? ENGINE_API(12, OgaResult* (*)(const OgaEngineEvent*, OgaErrorCode*))(e, o) : NULL; }
OgaResult* EngineEventGetUsage(const OgaEngineEvent* e, const OgaTurnUsage** o) { return g_engine_initialized ? ENGINE_API(13, OgaResult* (*)(const OgaEngineEvent*, const OgaTurnUsage**))(e, o) : NULL; }
OgaResult* TurnUsageGetPromptTokens(const OgaTurnUsage* u, uint64_t* o) { return g_engine_initialized ? ENGINE_API(14, OgaResult* (*)(const OgaTurnUsage*, uint64_t*))(u, o) : NULL; }
OgaResult* TurnUsageGetGeneratedTokens(const OgaTurnUsage* u, uint64_t* o) { return g_engine_initialized ? ENGINE_API(15, OgaResult* (*)(const OgaTurnUsage*, uint64_t*))(u, o) : NULL; }
OgaResult* TurnUsageGetCachedPromptTokens(const OgaTurnUsage* u, uint64_t* o) { return g_engine_initialized ? ENGINE_API(16, OgaResult* (*)(const OgaTurnUsage*, uint64_t*))(u, o) : NULL; }
OgaResult* EngineHasPendingRequests(OgaEngine* e, bool* o) { return g_engine_initialized ? ENGINE_API(17, OgaResult* (*)(OgaEngine*, bool*))(e, o) : NULL; }
OgaResult* EngineCreateRequest(OgaEngine* e, const OgaRequestOptions* p, OgaRequest** o) { return g_engine_initialized ? ENGINE_API(18, OgaResult* (*)(OgaEngine*, const OgaRequestOptions*, OgaRequest**))(e, p, o) : NULL; }
OgaResult* CreateRequestOptions(OgaRequestOptions** o) { return g_engine_initialized ? ENGINE_API(19, OgaResult* (*)(OgaRequestOptions**))(o) : NULL; }
void DestroyRequestOptions(OgaRequestOptions* o) { if (g_engine_initialized && o) ENGINE_API(20, void (*)(OgaRequestOptions*))(o); }
OgaResult* RequestOptionsSetMaxSessionTokens(OgaRequestOptions* p, uint64_t v) { return g_engine_initialized ? ENGINE_API(21, OgaResult* (*)(OgaRequestOptions*, uint64_t))(p, v) : NULL; }
OgaResult* RequestCreateTurnOptions(OgaRequest* r, OgaTurnOptions** o) { return g_engine_initialized ? ENGINE_API(22, OgaResult* (*)(OgaRequest*, OgaTurnOptions**))(r, o) : NULL; }
void DestroyTurnOptions(OgaTurnOptions* o) { if (g_engine_initialized && o) ENGINE_API(23, void (*)(OgaTurnOptions*))(o); }
OgaResult* TurnOptionsSetMaxGeneratedTokens(OgaTurnOptions* o, uint64_t v) { return g_engine_initialized ? ENGINE_API(24, OgaResult* (*)(OgaTurnOptions*, uint64_t))(o, v) : NULL; }
OgaResult* TurnOptionsSetMinGeneratedTokens(OgaTurnOptions* o, uint64_t v) { return g_engine_initialized ? ENGINE_API(25, OgaResult* (*)(OgaTurnOptions*, uint64_t))(o, v) : NULL; }
OgaResult* TurnOptionsSetDoSample(OgaTurnOptions* o, bool v) { return g_engine_initialized ? ENGINE_API(26, OgaResult* (*)(OgaTurnOptions*, bool))(o, v) : NULL; }
OgaResult* TurnOptionsSetTemperature(OgaTurnOptions* o, float v) { return g_engine_initialized ? ENGINE_API(27, OgaResult* (*)(OgaTurnOptions*, float))(o, v) : NULL; }
OgaResult* TurnOptionsSetTopP(OgaTurnOptions* o, float v) { return g_engine_initialized ? ENGINE_API(28, OgaResult* (*)(OgaTurnOptions*, float))(o, v) : NULL; }
OgaResult* TurnOptionsSetTopK(OgaTurnOptions* o, int32_t v) { return g_engine_initialized ? ENGINE_API(29, OgaResult* (*)(OgaTurnOptions*, int32_t))(o, v) : NULL; }
OgaResult* TurnOptionsSetRepetitionPenalty(OgaTurnOptions* o, float v) { return g_engine_initialized ? ENGINE_API(30, OgaResult* (*)(OgaTurnOptions*, float))(o, v) : NULL; }
OgaResult* TurnOptionsSetNoRepeatNgramSize(OgaTurnOptions* o, int32_t v) { return g_engine_initialized ? ENGINE_API(31, OgaResult* (*)(OgaTurnOptions*, int32_t))(o, v) : NULL; }
OgaResult* TurnOptionsSetSeed(OgaTurnOptions* o, uint64_t v) { return g_engine_initialized ? ENGINE_API(32, OgaResult* (*)(OgaTurnOptions*, uint64_t))(o, v) : NULL; }
OgaResult* TurnOptionsClearSeed(OgaTurnOptions* o) { return g_engine_initialized ? ENGINE_API(33, OgaResult* (*)(OgaTurnOptions*))(o) : NULL; }
OgaResult* TurnOptionsSetStopStrings(OgaTurnOptions* o, const OgaStringArray* v) { return g_engine_initialized ? ENGINE_API(34, OgaResult* (*)(OgaTurnOptions*, const OgaStringArray*))(o, v) : NULL; }
OgaResult* TurnOptionsSetGuidance(OgaTurnOptions* o, const char* t, const char* d) { return g_engine_initialized ? ENGINE_API(35, OgaResult* (*)(OgaTurnOptions*, const char*, const char*))(o, t, d) : NULL; }
OgaResult* TurnOptionsClearGuidance(OgaTurnOptions* o) { return g_engine_initialized ? ENGINE_API(36, OgaResult* (*)(OgaTurnOptions*))(o) : NULL; }
OgaResult* TurnOptionsReset(OgaTurnOptions* o) { return g_engine_initialized ? ENGINE_API(37, OgaResult* (*)(OgaTurnOptions*))(o) : NULL; }
OgaResult* RequestBeginTurn(OgaRequest* r, const OgaTurnOptions* o, const int32_t* ids, uint64_t n, uint64_t* id) { return g_engine_initialized ? ENGINE_API(38, OgaResult* (*)(OgaRequest*, const OgaTurnOptions*, const int32_t*, uint64_t, uint64_t*))(r, o, ids, n, id) : NULL; }
OgaResult* RequestCancelTurn(OgaRequest* r, uint64_t id, bool* c) { return g_engine_initialized ? ENGINE_API(39, OgaResult* (*)(OgaRequest*, uint64_t, bool*))(r, id, c) : NULL; }
OgaResult* RequestRewindToStartOfTurn(OgaRequest* r, uint64_t id) { return g_engine_initialized ? ENGINE_API(40, OgaResult* (*)(OgaRequest*, uint64_t))(r, id) : NULL; }
OgaResult* RequestClose(OgaRequest* r) { return g_engine_initialized ? ENGINE_API(41, OgaResult* (*)(OgaRequest*))(r) : NULL; }
OgaResult* RequestSetDraftTokens(OgaRequest* r, const OgaSequences* s) { return g_engine_initialized ? ENGINE_API(42, OgaResult* (*)(OgaRequest*, const OgaSequences*))(r, s) : NULL; }
void DestroyRequest(OgaRequest* r) { if (g_engine_initialized && r) ENGINE_API(43, void (*)(OgaRequest*))(r); }
OgaResult* EngineMaxDraftTokensPerProposal(const OgaEngine* e, size_t* o) { return g_engine_initialized ? ENGINE_API(44, OgaResult* (*)(const OgaEngine*, size_t*))(e, o) : NULL; }
OgaResult* AppendTokenSequence(const int32_t* ids, size_t count, OgaSequences* sequences) { return g_engine_initialized ? ENGINE_API(45, OgaResult* (*)(const int32_t*, size_t, OgaSequences*))(ids, count, sequences) : NULL; }
size_t SequencesCount(const OgaSequences* s) { return g_engine_initialized ? ENGINE_API(46, size_t (*)(const OgaSequences*))(s) : 0; }
size_t SequencesGetSequenceCount(const OgaSequences* s, size_t i) { return g_engine_initialized ? ENGINE_API(47, size_t (*)(const OgaSequences*, size_t))(s, i) : 0; }
const int32_t* SequencesGetSequenceData(const OgaSequences* s, size_t i) { return g_engine_initialized ? ENGINE_API(48, const int32_t* (*)(const OgaSequences*, size_t))(s, i) : NULL; }
OgaResult* TokenizerDecode(const OgaTokenizer* t, const int32_t* ids, size_t n, const char** out) { return g_engine_initialized ? ENGINE_API(49, OgaResult* (*)(const OgaTokenizer*, const int32_t*, size_t, const char**))(t, ids, n, out) : NULL; }
OgaResult* TokenizerToTokenId(const OgaTokenizer* t, const char* text, int32_t* out) { return g_engine_initialized ? ENGINE_API(50, OgaResult* (*)(const OgaTokenizer*, const char*, int32_t*))(t, text, out) : NULL; }
OgaResult* TokenizerEncodeBatch(const OgaTokenizer* t, const char** values, size_t n, OgaTensor** out) { return g_engine_initialized ? ENGINE_API(51, OgaResult* (*)(const OgaTokenizer*, const char**, size_t, OgaTensor**))(t, values, n, out) : NULL; }
OgaResult* TokenizerDecodeBatch(const OgaTokenizer* t, const OgaTensor* tensor, OgaStringArray** out) { return g_engine_initialized ? ENGINE_API(52, OgaResult* (*)(const OgaTokenizer*, const OgaTensor*, OgaStringArray**))(t, tensor, out) : NULL; }
OgaResult* CreateTensorFromBuffer(void* data, const int64_t* shape, size_t rank, OgaElementType type, OgaTensor** out) { return g_engine_initialized ? ENGINE_API(53, OgaResult* (*)(void*, const int64_t*, size_t, OgaElementType, OgaTensor**))(data, shape, rank, type, out) : NULL; }
void DestroyTensor(OgaTensor* tensor) { if (g_engine_initialized && tensor) ENGINE_API(54, void (*)(OgaTensor*))(tensor); }
OgaResult* TensorGetType(OgaTensor* tensor, OgaElementType* out) { return g_engine_initialized ? ENGINE_API(55, OgaResult* (*)(OgaTensor*, OgaElementType*))(tensor, out) : NULL; }
OgaResult* TensorGetShapeRank(OgaTensor* tensor, size_t* out) { return g_engine_initialized ? ENGINE_API(56, OgaResult* (*)(OgaTensor*, size_t*))(tensor, out) : NULL; }
OgaResult* TensorGetShape(OgaTensor* tensor, int64_t* shape, size_t rank) { return g_engine_initialized ? ENGINE_API(57, OgaResult* (*)(OgaTensor*, int64_t*, size_t))(tensor, shape, rank) : NULL; }
OgaResult* TensorGetData(OgaTensor* tensor, void** out) { return g_engine_initialized ? ENGINE_API(58, OgaResult* (*)(OgaTensor*, void**))(tensor, out) : NULL; }
OgaResult* StringArrayGetCount(const OgaStringArray* strings, size_t* out) { return g_engine_initialized ? ENGINE_API(59, OgaResult* (*)(const OgaStringArray*, size_t*))(strings, out) : NULL; }
OgaResult* StringArrayGetString(const OgaStringArray* strings, size_t index, const char** out) { return g_engine_initialized ? ENGINE_API(60, OgaResult* (*)(const OgaStringArray*, size_t, const char**))(strings, index, out) : NULL; }
OgaResult* SetLogBool(const char* name, bool value) { return g_engine_initialized ? ENGINE_API(61, OgaResult* (*)(const char*, bool))(name, value) : NULL; }
OgaResult* SetLogString(const char* name, const char* value) { return g_engine_initialized ? ENGINE_API(62, OgaResult* (*)(const char*, const char*))(name, value) : NULL; }
OgaResult* SetCurrentGpuDeviceId(int id) { return g_engine_initialized ? ENGINE_API(63, OgaResult* (*)(int))(id) : NULL; }
OgaResult* GetCurrentGpuDeviceId(int* id) { return g_engine_initialized ? ENGINE_API(64, OgaResult* (*)(int*))(id) : NULL; }
OgaResult* CreateRuntimeSettings(OgaRuntimeSettings** out) { return g_engine_initialized ? ENGINE_API(65, OgaResult* (*)(OgaRuntimeSettings**))(out) : NULL; }
void DestroyRuntimeSettings(OgaRuntimeSettings* settings) { if (g_engine_initialized && settings) ENGINE_API(66, void (*)(OgaRuntimeSettings*))(settings); }
OgaResult* RuntimeSettingsSetHandle(OgaRuntimeSettings* settings, const char* name, void* handle) { return g_engine_initialized ? ENGINE_API(67, OgaResult* (*)(OgaRuntimeSettings*, const char*, void*))(settings, name, handle) : NULL; }
OgaResult* CreateModelWithRuntimeSettings(const char* path, const OgaRuntimeSettings* settings, OgaModel** out) { return g_engine_initialized ? ENGINE_API(68, OgaResult* (*)(const char*, const OgaRuntimeSettings*, OgaModel**))(path, settings, out) : NULL; }
OgaResult* LoadAudio(const char* path, OgaAudios** out) { return g_engine_initialized ? ENGINE_API(69, OgaResult* (*)(const char*, OgaAudios**))(path, out) : NULL; }
OgaResult* LoadAudios(const OgaStringArray* paths, OgaAudios** out) { return g_engine_initialized ? ENGINE_API(70, OgaResult* (*)(const OgaStringArray*, OgaAudios**))(paths, out) : NULL; }
OgaResult* LoadAudiosFromBuffers(const void** data, const size_t* sizes, size_t count, OgaAudios** out) { return g_engine_initialized ? ENGINE_API(71, OgaResult* (*)(const void**, const size_t*, size_t, OgaAudios**))(data, sizes, count, out) : NULL; }
void DestroyAudios(OgaAudios* audios) { if (g_engine_initialized && audios) ENGINE_API(72, void (*)(OgaAudios*))(audios); }
OgaResult* ProcessAudios(const OgaMultiModalProcessor* p, const char* prompt, const OgaAudios* audios, OgaNamedTensors** out) { return g_engine_initialized ? ENGINE_API(73, OgaResult* (*)(const OgaMultiModalProcessor*, const char*, const OgaAudios*, OgaNamedTensors**))(p, prompt, audios, out) : NULL; }
OgaResult* ProcessAudiosAndPrompts(const OgaMultiModalProcessor* p, const OgaStringArray* prompts, const OgaAudios* audios, OgaNamedTensors** out) { return g_engine_initialized ? ENGINE_API(74, OgaResult* (*)(const OgaMultiModalProcessor*, const OgaStringArray*, const OgaAudios*, OgaNamedTensors**))(p, prompts, audios, out) : NULL; }
OgaResult* ProcessImagesAndAudios(const OgaMultiModalProcessor* p, const char* prompt, const OgaImages* images, const OgaAudios* audios, OgaNamedTensors** out) { return g_engine_initialized ? ENGINE_API(75, OgaResult* (*)(const OgaMultiModalProcessor*, const char*, const OgaImages*, const OgaAudios*, OgaNamedTensors**))(p, prompt, images, audios, out) : NULL; }
OgaResult* ProcessImagesAndAudiosAndPrompts(const OgaMultiModalProcessor* p, const OgaStringArray* prompts, const OgaImages* images, const OgaAudios* audios, OgaNamedTensors** out) { return g_engine_initialized ? ENGINE_API(76, OgaResult* (*)(const OgaMultiModalProcessor*, const OgaStringArray*, const OgaImages*, const OgaAudios*, OgaNamedTensors**))(p, prompts, images, audios, out) : NULL; }
OgaResult* CreateAdapters(const OgaModel* model, OgaAdapters** out) { return g_engine_initialized ? ENGINE_API(77, OgaResult* (*)(const OgaModel*, OgaAdapters**))(model, out) : NULL; }
void DestroyAdapters(OgaAdapters* adapters) { if (g_engine_initialized && adapters) ENGINE_API(78, void (*)(OgaAdapters*))(adapters); }
OgaResult* LoadAdapter(OgaAdapters* adapters, const char* path, const char* name) { return g_engine_initialized ? ENGINE_API(79, OgaResult* (*)(OgaAdapters*, const char*, const char*))(adapters, path, name) : NULL; }
OgaResult* UnloadAdapter(OgaAdapters* adapters, const char* name) { return g_engine_initialized ? ENGINE_API(80, OgaResult* (*)(OgaAdapters*, const char*))(adapters, name) : NULL; }
OgaResult* SetActiveAdapter(OgaGenerator* generator, OgaAdapters* adapters, const char* name) { return g_engine_initialized ? ENGINE_API(81, OgaResult* (*)(OgaGenerator*, OgaAdapters*, const char*))(generator, adapters, name) : NULL; }
OgaResult* CreateMtpGenerator(const OgaModel* mainModel, const OgaModel* mtpModel, const OgaGeneratorParams* params, OgaMtpGenerator** out) { return g_engine_initialized ? ENGINE_API(82, OgaResult* (*)(const OgaModel*, const OgaModel*, const OgaGeneratorParams*, OgaMtpGenerator**))(mainModel, mtpModel, params, out) : NULL; }
OgaResult* MtpGeneratorAppendTokens(OgaMtpGenerator* generator, const int32_t* ids, size_t count) { return g_engine_initialized ? ENGINE_API(83, OgaResult* (*)(OgaMtpGenerator*, const int32_t*, size_t))(generator, ids, count) : NULL; }
OgaResult* MtpGeneratorGenerateNextToken(OgaMtpGenerator* generator) { return g_engine_initialized ? ENGINE_API(84, OgaResult* (*)(OgaMtpGenerator*))(generator) : NULL; }
OgaResult* MtpGeneratorReset(OgaMtpGenerator* generator) { return g_engine_initialized ? ENGINE_API(85, OgaResult* (*)(OgaMtpGenerator*))(generator) : NULL; }
bool MtpGeneratorIsDone(const OgaMtpGenerator* generator) { return g_engine_initialized ? ENGINE_API(86, bool (*)(const OgaMtpGenerator*))(generator) : false; }
size_t MtpGeneratorGetSequenceCount(const OgaMtpGenerator* generator) { return g_engine_initialized ? ENGINE_API(87, size_t (*)(const OgaMtpGenerator*))(generator) : 0; }
const int32_t* MtpGeneratorGetSequenceData(const OgaMtpGenerator* generator) { return g_engine_initialized ? ENGINE_API(88, const int32_t* (*)(const OgaMtpGenerator*))(generator) : NULL; }
size_t MtpGeneratorGetForwardCount(const OgaMtpGenerator* generator) { return g_engine_initialized ? ENGINE_API(89, size_t (*)(const OgaMtpGenerator*))(generator) : 0; }
size_t MtpGeneratorGetAcceptCount(const OgaMtpGenerator* generator) { return g_engine_initialized ? ENGINE_API(90, size_t (*)(const OgaMtpGenerator*))(generator) : 0; }
size_t MtpGeneratorGetTrialCount(const OgaMtpGenerator* generator) { return g_engine_initialized ? ENGINE_API(91, size_t (*)(const OgaMtpGenerator*))(generator) : 0; }
OgaResult* MtpGeneratorGetSpeculativeStats(const OgaMtpGenerator* generator, OgaSpeculativeStats** out) { return g_engine_initialized ? ENGINE_API(92, OgaResult* (*)(const OgaMtpGenerator*, OgaSpeculativeStats**))(generator, out) : NULL; }
void DestroyMtpGenerator(OgaMtpGenerator* generator) { if (g_engine_initialized && generator) ENGINE_API(93, void (*)(OgaMtpGenerator*))(generator); }
OgaResult* GeneratorParamsSetSearchBool(OgaGeneratorParams* params, const char* name, bool value) { return g_engine_initialized ? ENGINE_API(94, OgaResult* (*)(OgaGeneratorParams*, const char*, bool))(params, name, value) : NULL; }
OgaResult* GeneratorSetModelInput(OgaGenerator* g, const char* name, OgaTensor* tensor) { return g_engine_initialized ? ENGINE_API(95, OgaResult* (*)(OgaGenerator*, const char*, OgaTensor*))(g, name, tensor) : NULL; }
OgaResult* GeneratorAppendTokens(OgaGenerator* g, const int32_t* ids, size_t count) { return g_engine_initialized ? ENGINE_API(96, OgaResult* (*)(OgaGenerator*, const int32_t*, size_t))(g, ids, count) : NULL; }
size_t GeneratorTokenCount(const OgaGenerator* g) { return g_engine_initialized ? ENGINE_API(97, size_t (*)(const OgaGenerator*))(g) : 0; }
OgaResult* GeneratorGetNextTokens(const OgaGenerator* g, const int32_t** out, size_t* count) { return g_engine_initialized ? ENGINE_API(98, OgaResult* (*)(const OgaGenerator*, const int32_t**, size_t*))(g, out, count) : NULL; }
OgaResult* GeneratorSetRuntimeOption(OgaGenerator* g, const char* key, const char* value) { return g_engine_initialized ? ENGINE_API(99, OgaResult* (*)(OgaGenerator*, const char*, const char*))(g, key, value) : NULL; }
OgaResult* GeneratorRewindTo(OgaGenerator* g, size_t length) { return g_engine_initialized ? ENGINE_API(100, OgaResult* (*)(OgaGenerator*, size_t))(g, length) : NULL; }
OgaResult* GeneratorSnapshotState(OgaGenerator* g) { return g_engine_initialized ? ENGINE_API(101, OgaResult* (*)(OgaGenerator*))(g) : NULL; }
OgaResult* GeneratorSetHiddenStates(OgaGenerator* g, OgaTensor* tensor) { return g_engine_initialized ? ENGINE_API(102, OgaResult* (*)(OgaGenerator*, OgaTensor*))(g, tensor) : NULL; }
OgaResult* GeneratorGetInput(const OgaGenerator* g, const char* name, OgaTensor** out) { return g_engine_initialized ? ENGINE_API(103, OgaResult* (*)(const OgaGenerator*, const char*, OgaTensor**))(g, name, out) : NULL; }
OgaResult* GeneratorGetOutput(const OgaGenerator* g, const char* name, OgaTensor** out) { return g_engine_initialized ? ENGINE_API(104, OgaResult* (*)(const OgaGenerator*, const char*, OgaTensor**))(g, name, out) : NULL; }
OgaResult* GeneratorGetLogits(OgaGenerator* g, OgaTensor** out) { return g_engine_initialized ? ENGINE_API(105, OgaResult* (*)(OgaGenerator*, OgaTensor**))(g, out) : NULL; }
OgaResult* GeneratorSetLogits(OgaGenerator* g, OgaTensor* tensor) { return g_engine_initialized ? ENGINE_API(106, OgaResult* (*)(OgaGenerator*, OgaTensor*))(g, tensor) : NULL; }
OgaResult* GeneratorGetSpeculativeStats(const OgaGenerator* g, OgaSpeculativeStats** out) { return g_engine_initialized ? ENGINE_API(107, OgaResult* (*)(const OgaGenerator*, OgaSpeculativeStats**))(g, out) : NULL; }
void DestroySpeculativeStats(OgaSpeculativeStats* stats) { if (g_engine_initialized && stats) ENGINE_API(108, void (*)(OgaSpeculativeStats*))(stats); }
OgaResult* SpeculativeStatsGetCount(const OgaSpeculativeStats* stats, const char* name, uint64_t* out) { return g_engine_initialized ? ENGINE_API(109, OgaResult* (*)(const OgaSpeculativeStats*, const char*, uint64_t*))(stats, name, out) : NULL; }
OgaResult* SpeculativeStatsGetAcceptanceLengthCount(const OgaSpeculativeStats* stats, size_t length, uint64_t* out) { return g_engine_initialized ? ENGINE_API(110, OgaResult* (*)(const OgaSpeculativeStats*, size_t, uint64_t*))(stats, length, out) : NULL; }
OgaResult* SpeculativeStatsGetAcceptanceLengthHistogramSize(const OgaSpeculativeStats* stats, size_t* out) { return g_engine_initialized ? ENGINE_API(111, OgaResult* (*)(const OgaSpeculativeStats*, size_t*))(stats, out) : NULL; }
OgaResult* SpeculativeStatsGetNumber(const OgaSpeculativeStats* stats, const char* name, double* out) { return g_engine_initialized ? ENGINE_API(112, OgaResult* (*)(const OgaSpeculativeStats*, const char*, double*))(stats, name, out) : NULL; }
OgaResult* SpeculativeStatsGetBool(const OgaSpeculativeStats* stats, const char* name, bool* out) { return g_engine_initialized ? ENGINE_API(113, OgaResult* (*)(const OgaSpeculativeStats*, const char*, bool*))(stats, name, out) : NULL; }
OgaResult* TokenizerUpdateOptions(OgaTokenizer* tokenizer, const char* const* keys, const char* const* values, size_t count) { return g_engine_initialized ? ENGINE_API(114, OgaResult* (*)(OgaTokenizer*, const char* const*, const char* const*, size_t))(tokenizer, keys, values, count) : NULL; }
OgaResult* SetLogCallback(bool enabled) { return g_engine_initialized ? ENGINE_API(115, OgaResult* (*)(void (*)(const char*, size_t)))(enabled ? logCallbackBridge : NULL) : NULL; }
OgaResult* EngineEventGetRequest(const OgaEngineEvent* event, const OgaRequest** out) { return g_engine_initialized ? ENGINE_API(116, OgaResult* (*)(const OgaEngineEvent*, const OgaRequest**))(event, out) : NULL; }
OgaResult* EngineGetSpeculativeStats(const OgaEngine* engine, OgaSpeculativeStats** out) { return g_engine_initialized ? ENGINE_API(117, OgaResult* (*)(const OgaEngine*, OgaSpeculativeStats**))(engine, out) : NULL; }
OgaResult* EngineGetCapabilities(const OgaEngine* engine, OgaEngineCapabilities** out) { return g_engine_initialized ? ENGINE_API(118, OgaResult* (*)(const OgaEngine*, OgaEngineCapabilities**))(engine, out) : NULL; }
size_t EngineCapabilitiesGetConfiguredMaxBatchSize(const OgaEngineCapabilities* capabilities) { return g_engine_initialized ? ENGINE_API(119, size_t (*)(const OgaEngineCapabilities*))(capabilities) : 0; }
size_t EngineCapabilitiesGetMaxScheduledTokens(const OgaEngineCapabilities* capabilities) { return g_engine_initialized ? ENGINE_API(120, size_t (*)(const OgaEngineCapabilities*))(capabilities) : 0; }
uint64_t EngineCapabilitiesGetMaxRequestLength(const OgaEngineCapabilities* capabilities) { return g_engine_initialized ? ENGINE_API(121, uint64_t (*)(const OgaEngineCapabilities*))(capabilities) : 0; }
void DestroyEngineCapabilities(OgaEngineCapabilities* capabilities) { if (g_engine_initialized && capabilities) ENGINE_API(122, void (*)(OgaEngineCapabilities*))(capabilities); }
#undef ENGINE_API

void OgaShutdown(void) {
	if (g_initialized && g_api.Shutdown) g_api.Shutdown();
	memset(&g_api, 0, sizeof(g_api));
	memset(g_engine_api, 0, sizeof(g_engine_api));
	g_initialized = 0;
	g_engine_initialized = 0;
}

void OgaSetTelemetryEnabled(bool enabled) {
	if (!g_initialized || !g_api.SetTelemetryEnabled) return;
	g_api.SetTelemetryEnabled(enabled);
}

OgaResult* OgaCreateConfigFromPackageEp(const char* config_path, const char* ep, OgaConfig** out) {
	if (!g_initialized || !g_api.CreateConfigFromPackageEp) return NULL;
	return g_api.CreateConfigFromPackageEp(config_path, ep, out);
}

OgaResult* OgaTokenizerGetPadTokenId(const OgaTokenizer* tokenizer, int32_t* token_id) {
	if (!g_initialized || !g_api.TokenizerGetPadTokenId) return NULL;
	return g_api.TokenizerGetPadTokenId(tokenizer, token_id);
}

OgaResult* OgaTokenizerGetBotTokenId(const OgaTokenizer* tokenizer, int32_t* token_id) {
	if (!g_initialized || !g_api.TokenizerGetBotTokenId) return NULL;
	return g_api.TokenizerGetBotTokenId(tokenizer, token_id);
}

OgaResult* OgaTokenizerGetEotTokenId(const OgaTokenizer* tokenizer, int32_t* token_id) {
	if (!g_initialized || !g_api.TokenizerGetEotTokenId) return NULL;
	return g_api.TokenizerGetEotTokenId(tokenizer, token_id);
}

OgaResult* OgaTokenizerGetBorTokenId(const OgaTokenizer* tokenizer, int32_t* token_id) {
	if (!g_initialized || !g_api.TokenizerGetBorTokenId) return NULL;
	return g_api.TokenizerGetBorTokenId(tokenizer, token_id);
}

OgaResult* OgaTokenizerGetEorTokenId(const OgaTokenizer* tokenizer, int32_t* token_id) {
	if (!g_initialized || !g_api.TokenizerGetEorTokenId) return NULL;
	return g_api.TokenizerGetEorTokenId(tokenizer, token_id);
}
