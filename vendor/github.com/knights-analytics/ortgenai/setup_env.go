package ortgenai

/*
#cgo linux LDFLAGS: -ldl
#include <stdio.h>
#include <stdlib.h>
#ifdef _WIN32
#include <windows.h>
#define RTLD_LAZY 0
static void* dlopen(const char* path, int flags) {
	(void)flags;
	return (void*)LoadLibraryExA(path, NULL, LOAD_WITH_ALTERED_SEARCH_PATH);
}
static void* dlsym(void* handle, const char* name) {
	return (void*)GetProcAddress((HMODULE)handle, name);
}
static int dlclose(void* handle) {
	return FreeLibrary((HMODULE)handle) ? 0 : (int)GetLastError();
}
static const char* dlerror(void) {
	static char error[64];
	snprintf(error, sizeof(error), "Windows error %lu", (unsigned long)GetLastError());
	return error;
}
#else
#include <dlfcn.h>
#endif
#include "ort_genai_wrapper.h"
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

// genAiLibraryHandle holds the dlopen handle for the GenAI shared library once loaded.
var genAiLibraryHandle unsafe.Pointer

func platformCleanup() error {
	if genAiLibraryHandle == nil {
		return nil
	}
	clearLogCallback()
	C.OgaShutdown()
	if returnCode := C.dlclose(genAiLibraryHandle); returnCode != 0 {
		return fmt.Errorf("error closing GenAI shared library: %d", int(returnCode))
	}
	genAiLibraryHandle = nil
	return nil
}

func createSym(handle unsafe.Pointer, name string) unsafe.Pointer {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.dlerror() // clear previous error per POSIX
	sym := C.dlsym(handle, cName)
	return sym
}

// InitializeGenAiLibrary loads the ONNX Runtime GenAI shared library specified by
// onnxGenaiSharedLibraryPath (or a default) so its exported symbols become available.
// The assumption is that libonnxruntime.so is available in the same folder where the libonnxruntime-genai.so is located.
func InitializeGenAiLibrary() error {
	if genAiLibraryHandle != nil {
		return fmt.Errorf("GenAI library already initialized")
	}
	libPath := onnxGenaiSharedLibraryPath
	if libPath == "" {
		libPath = "libonnxruntime-genai.so"
		if runtime.GOOS == "windows" {
			libPath = "onnxruntime-genai.dll"
		}
	}
	cName := C.CString(libPath)
	defer C.free(unsafe.Pointer(cName))
	handle := C.dlopen(cName, C.RTLD_LAZY)
	if handle == nil {
		msg := C.GoString(C.dlerror())
		return fmt.Errorf("error loading GenAI shared library %q: %v", libPath, msg)
	}

	symCreate := createSym(handle, "OgaCreateModel")
	if symCreate == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateModel")
	}

	symErr := createSym(handle, "OgaResultGetError")
	if symErr == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaResultGetError")
	}
	symDestroyRes := createSym(handle, "OgaDestroyResult")
	if symDestroyRes == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyResult")
	}
	symDestroyModel := createSym(handle, "OgaDestroyModel")
	if symDestroyModel == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyModel")
	}
	symCreateTokenizer := createSym(handle, "OgaCreateTokenizer")
	if symCreateTokenizer == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateTokenizer")
	}
	symDestroyTokenizer := createSym(handle, "OgaDestroyTokenizer")
	if symDestroyTokenizer == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyTokenizer")
	}
	symCreateTokenizerStream := createSym(handle, "OgaCreateTokenizerStream")
	if symCreateTokenizerStream == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateTokenizerStream")
	}
	symDestroyTokenizerStream := createSym(handle, "OgaDestroyTokenizerStream")
	if symDestroyTokenizerStream == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyTokenizerStream")
	}
	symApplyChatTemplate := createSym(handle, "OgaTokenizerApplyChatTemplate")
	if symApplyChatTemplate == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerApplyChatTemplate")
	}
	symDestroyString := createSym(handle, "OgaDestroyString")
	if symDestroyString == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyString")
	}
	symCreateSequence := createSym(handle, "OgaCreateSequences")
	if symCreateSequence == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateSequence")
	}
	symDestroySequence := createSym(handle, "OgaDestroySequences")
	if symDestroySequence == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroySequences")
	}
	symTokenizerEncode := createSym(handle, "OgaTokenizerEncode")
	if symTokenizerEncode == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerEncode")
	}
	symCreateGenerator := createSym(handle, "OgaCreateGenerator")
	if symCreateGenerator == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateGenerator")
	}
	symDestroyGenerator := createSym(handle, "OgaDestroyGenerator")
	if symDestroyGenerator == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyGenerator")
	}
	symCreateGeneratorParams := createSym(handle, "OgaCreateGeneratorParams")
	if symCreateGeneratorParams == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateGeneratorParams")
	}
	symDestroyGeneratorParams := createSym(handle, "OgaDestroyGeneratorParams")
	if symDestroyGeneratorParams == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyGeneratorParams")
	}
	symGeneratorParamsSetSearchNumber := createSym(handle, "OgaGeneratorParamsSetSearchNumber")
	if symGeneratorParamsSetSearchNumber == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaGeneratorParamsSetSearchNumber")
	}
	symGeneratorAppendTokenSequences := createSym(handle, "OgaGenerator_AppendTokenSequences")
	if symGeneratorAppendTokenSequences == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaGeneratorAppendTokenSequences")
	}
	symGeneratorSetInputs := createSym(handle, "OgaGenerator_SetInputs")
	if symGeneratorSetInputs == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaGeneratorSetInputs")
	}
	symGeneratorGenerateNextToken := createSym(handle, "OgaGenerator_GenerateNextToken")
	if symGeneratorGenerateNextToken == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaGeneratorGenerateNextToken")
	}
	symGeneratorGetSequenceCount := createSym(handle, "OgaGenerator_GetSequenceCount")
	if symGeneratorGetSequenceCount == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaGeneratorGetSequenceCount")
	}
	symGeneratorGetSequenceData := createSym(handle, "OgaGenerator_GetSequenceData")
	if symGeneratorGetSequenceData == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaGeneratorGetSequenceData")
	}
	symTokenizerStreamDecode := createSym(handle, "OgaTokenizerStreamDecode")
	if symTokenizerStreamDecode == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerStreamDecode")
	}
	symIsDone := createSym(handle, "OgaGenerator_IsDone")
	if symIsDone == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaGenerator_IsDone")
	}

	// EOS token ids
	symTokenizerGetEosTokenIDs := createSym(handle, "OgaTokenizerGetEosTokenIds")
	if symTokenizerGetEosTokenIDs == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerGetEosTokenIds")
	}

	// Config-related symbols
	symCreateConfig := createSym(handle, "OgaCreateConfig")
	if symCreateConfig == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateConfig")
	}
	symConfigClearProviders := createSym(handle, "OgaConfigClearProviders")
	if symConfigClearProviders == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaConfigClearProviders")
	}
	symConfigAppendProvider := createSym(handle, "OgaConfigAppendProvider")
	if symConfigAppendProvider == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaConfigAppendProvider")
	}
	symConfigSetProviderOption := createSym(handle, "OgaConfigSetProviderOption")
	if symConfigSetProviderOption == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaConfigSetProviderOption")
	}
	symCreateModelFromConfig := createSym(handle, "OgaCreateModelFromConfig")
	if symCreateModelFromConfig == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateModelFromConfig")
	}
	symDestroyConfig := createSym(handle, "OgaDestroyConfig")
	if symDestroyConfig == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyConfig")
	}

	// Multimodal symbols
	symLoadImage := createSym(handle, "OgaLoadImage")
	if symLoadImage == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaLoadImage")
	}
	symLoadImages := createSym(handle, "OgaLoadImages")
	if symLoadImages == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaLoadImages")
	}
	symLoadImagesFromBuffers := createSym(handle, "OgaLoadImagesFromBuffers")
	if symLoadImagesFromBuffers == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaLoadImagesFromBuffers")
	}
	symDestroyImages := createSym(handle, "OgaDestroyImages")
	if symDestroyImages == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyImages")
	}
	symCreateMultiModalProcessor := createSym(handle, "OgaCreateMultiModalProcessor")
	if symCreateMultiModalProcessor == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateMultiModalProcessor")
	}
	symDestroyMultiModalProcessor := createSym(handle, "OgaDestroyMultiModalProcessor")
	if symDestroyMultiModalProcessor == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyMultiModalProcessor")
	}
	symProcessorProcessImages := createSym(handle, "OgaProcessorProcessImages")
	if symProcessorProcessImages == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaProcessorProcessImages")
	}
	symDestroyNamedTensors := createSym(handle, "OgaDestroyNamedTensors")
	if symDestroyNamedTensors == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyNamedTensors")
	}
	symCreateStringArray := createSym(handle, "OgaCreateStringArray")
	if symCreateStringArray == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateStringArray")
	}
	symDestroyStringArray := createSym(handle, "OgaDestroyStringArray")
	if symDestroyStringArray == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaDestroyStringArray")
	}
	symStringArrayAddString := createSym(handle, "OgaStringArrayAddString")
	if symStringArrayAddString == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaStringArrayAddString")
	}
	symProcessorProcessImagesAndPrompts := createSym(handle, "OgaProcessorProcessImagesAndPrompts")
	if symProcessorProcessImagesAndPrompts == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaProcessorProcessImagesAndPrompts")
	}
	symGeneratorParamsSetGuidance := createSym(handle, "OgaGeneratorParamsSetGuidance")
	if symGeneratorParamsSetGuidance == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaGeneratorParamsSetGuidance")
	}

	symShutdown := createSym(handle, "OgaShutdown")
	if symShutdown == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaShutdown")
	}
	symSetTelemetry := createSym(handle, "OgaSetTelemetryEnabled")
	if symSetTelemetry == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaSetTelemetryEnabled")
	}
	symCreateConfigFromEp := createSym(handle, "OgaCreateConfigFromPackageEp")
	if symCreateConfigFromEp == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaCreateConfigFromPackageEp")
	}
	symGetPad := createSym(handle, "OgaTokenizerGetPadTokenId")
	if symGetPad == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerGetPadTokenId")
	}
	symGetBot := createSym(handle, "OgaTokenizerGetBotTokenId")
	if symGetBot == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerGetBotTokenId")
	}
	symGetEot := createSym(handle, "OgaTokenizerGetEotTokenId")
	if symGetEot == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerGetEotTokenId")
	}
	symGetBor := createSym(handle, "OgaTokenizerGetBorTokenId")
	if symGetBor == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerGetBorTokenId")
	}
	symGetEor := createSym(handle, "OgaTokenizerGetEorTokenId")
	if symGetEor == nil {
		C.dlclose(handle)
		return fmt.Errorf("missing OgaTokenizerGetEorTokenId")
	}

	engineSymbolNames := []string{
		"OgaCreateEngine", "OgaDestroyEngine", "OgaCreateEngineEventBuffer", "OgaDestroyEngineEventBuffer",
		"OgaEngineRun", "OgaEngineEventBufferGetCount", "OgaEngineEventBufferGet", "OgaEngineEventGetFlags",
		"OgaEngineEventGetTurnId", "OgaEngineEventGetToken", "OgaEngineEventGetFinishReason",
		"OgaEngineEventGetMatchedStopStringIndex", "OgaEngineEventGetErrorCode", "OgaEngineEventGetUsage",
		"OgaTurnUsageGetPromptTokens", "OgaTurnUsageGetGeneratedTokens", "OgaTurnUsageGetCachedPromptTokens",
		"OgaEngineHasPendingRequests", "OgaEngineCreateRequest", "OgaCreateRequestOptions", "OgaDestroyRequestOptions",
		"OgaRequestOptionsSetMaxSessionTokens", "OgaRequestCreateTurnOptions", "OgaDestroyTurnOptions",
		"OgaTurnOptionsSetMaxGeneratedTokens", "OgaTurnOptionsSetMinGeneratedTokens", "OgaTurnOptionsSetDoSample",
		"OgaTurnOptionsSetTemperature", "OgaTurnOptionsSetTopP", "OgaTurnOptionsSetTopK",
		"OgaTurnOptionsSetRepetitionPenalty", "OgaTurnOptionsSetNoRepeatNgramSize", "OgaTurnOptionsSetSeed",
		"OgaTurnOptionsClearSeed", "OgaTurnOptionsSetStopStrings", "OgaTurnOptionsSetGuidance",
		"OgaTurnOptionsClearGuidance", "OgaTurnOptionsReset", "OgaRequestBeginTurn", "OgaRequestCancelTurn",
		"OgaRequestRewindToStartOfTurn", "OgaRequestClose", "OgaRequestSetDraftTokens", "OgaDestroyRequest",
		"OgaEngineMaxDraftTokensPerProposal", "OgaAppendTokenSequence", "OgaSequencesCount",
		"OgaSequencesGetSequenceCount", "OgaSequencesGetSequenceData", "OgaTokenizerDecode", "OgaTokenizerToTokenId",
		"OgaTokenizerEncodeBatch", "OgaTokenizerDecodeBatch", "OgaCreateTensorFromBuffer", "OgaDestroyTensor",
		"OgaTensorGetType", "OgaTensorGetShapeRank", "OgaTensorGetShape", "OgaTensorGetData",
		"OgaStringArrayGetCount", "OgaStringArrayGetString",
		"OgaSetLogBool", "OgaSetLogString", "OgaSetCurrentGpuDeviceId", "OgaGetCurrentGpuDeviceId",
		"OgaCreateRuntimeSettings", "OgaDestroyRuntimeSettings", "OgaRuntimeSettingsSetHandle", "OgaCreateModelWithRuntimeSettings",
		"OgaLoadAudio", "OgaLoadAudios", "OgaLoadAudiosFromBuffers", "OgaDestroyAudios",
		"OgaProcessorProcessAudios", "OgaProcessorProcessAudiosAndPrompts",
		"OgaProcessorProcessImagesAndAudios", "OgaProcessorProcessImagesAndAudiosAndPrompts",
		"OgaCreateAdapters", "OgaDestroyAdapters", "OgaLoadAdapter", "OgaUnloadAdapter", "OgaSetActiveAdapter",
		"OgaCreateMtpGenerator", "OgaMtpGenerator_AppendTokens", "OgaMtpGenerator_GenerateNextToken",
		"OgaMtpGenerator_Reset", "OgaMtpGenerator_IsDone", "OgaMtpGenerator_GetSequenceCount",
		"OgaMtpGenerator_GetSequenceData", "OgaMtpGenerator_GetForwardCount",
		"OgaMtpGenerator_GetAcceptCount", "OgaMtpGenerator_GetTrialCount",
		"OgaMtpGenerator_GetSpeculativeStats", "OgaDestroyMtpGenerator",
		"OgaGeneratorParamsSetSearchBool",
		"OgaGenerator_SetModelInput", "OgaGenerator_AppendTokens", "OgaGenerator_TokenCount",
		"OgaGenerator_GetNextTokens", "OgaGenerator_SetRuntimeOption", "OgaGenerator_RewindTo",
		"OgaGenerator_SnapshotState", "OgaGenerator_SetHiddenStates", "OgaGenerator_GetInput",
		"OgaGenerator_GetOutput", "OgaGenerator_GetLogits", "OgaGenerator_SetLogits",
		"OgaGenerator_GetSpeculativeStats", "OgaDestroySpeculativeStats", "OgaSpeculativeStatsGetCount",
		"OgaSpeculativeStatsGetAcceptanceLengthCount", "OgaSpeculativeStatsGetAcceptanceLengthHistogramSize",
		"OgaSpeculativeStatsGetNumber", "OgaSpeculativeStatsGetBool",
		"OgaUpdateTokenizerOptions",
		"OgaSetLogCallback",
		"OgaEngineEventGetRequest", "OgaEngineGetSpeculativeStats",
		"OgaEngineGetCapabilities", "OgaEngineCapabilitiesGetConfiguredMaxBatchSize",
		"OgaEngineCapabilitiesGetMaxScheduledTokens", "OgaEngineCapabilitiesGetMaxRequestLength",
		"OgaDestroyEngineCapabilities",
	}
	engineSymbols := make([]unsafe.Pointer, 0, len(engineSymbolNames))
	for _, symbolName := range engineSymbolNames {
		symbol := createSym(handle, symbolName)
		if symbol == nil {
			C.dlclose(handle)
			return fmt.Errorf("missing required ORT GenAI symbol %s", symbolName)
		}
		engineSymbols = append(engineSymbols, symbol)
	}
	additionalRequiredSymbols := []string{
		"OgaEngineGetCapabilities", "OgaEngineCapabilitiesGetConfiguredMaxBatchSize",
		"OgaEngineCapabilitiesGetMaxScheduledTokens", "OgaEngineCapabilitiesGetMaxRequestLength",
		"OgaDestroyEngineCapabilities", "OgaEngineGetSpeculativeStats",
		"OgaDestroySpeculativeStats", "OgaSpeculativeStatsGetCount",
		"OgaSpeculativeStatsGetAcceptanceLengthCount", "OgaSpeculativeStatsGetAcceptanceLengthHistogramSize",
		"OgaSpeculativeStatsGetNumber", "OgaSpeculativeStatsGetBool",
		"OgaCreateTokenizerFromConfig", "OgaCreateTokenizerFromPath", "OgaUpdateTokenizerOptions",
		"OgaTokenizerGetBosTokenId", "OgaCreateTokenizerStreamFromProcessor",
		"OgaProcessorDecode", "OgaProcessorProcessAudios", "OgaProcessorProcessAudiosAndPrompts",
		"OgaProcessorProcessImagesAndAudios", "OgaProcessorProcessImagesAndAudiosAndPrompts",
		"OgaLoadAudio", "OgaLoadAudios", "OgaLoadAudiosFromBuffers", "OgaDestroyAudios",
		"OgaCreateMtpGenerator", "OgaMtpGenerator_AppendTokens", "OgaMtpGenerator_GenerateNextToken",
		"OgaMtpGenerator_Reset", "OgaMtpGenerator_IsDone", "OgaMtpGenerator_GetSequenceCount",
		"OgaMtpGenerator_GetSequenceData", "OgaMtpGenerator_GetForwardCount",
		"OgaMtpGenerator_GetAcceptCount", "OgaMtpGenerator_GetTrialCount",
		"OgaMtpGenerator_GetSpeculativeStats", "OgaDestroyMtpGenerator",
		"OgaGenerator_IsSessionTerminated", "OgaGenerator_SetModelInput", "OgaGenerator_AppendTokens",
		"OgaGenerator_TokenCount", "OgaGenerator_GetNextTokens", "OgaGenerator_SetRuntimeOption",
		"OgaGenerator_RewindTo", "OgaGenerator_SnapshotState", "OgaGenerator_SetHiddenStates",
		"OgaGenerator_GetInput", "OgaGenerator_GetOutput", "OgaGenerator_GetLogits",
		"OgaGenerator_SetLogits", "OgaGenerator_GetSpeculativeStats", "OgaCreateAdapters", "OgaDestroyAdapters", "OgaLoadAdapter",
		"OgaUnloadAdapter", "OgaSetActiveAdapter", "OgaSetLogCallback",
		"OgaConfigAddModelData", "OgaConfigRemoveModelData", "OgaConfigSetDecoderProviderOptionsHardwareDeviceType",
		"OgaConfigSetDecoderProviderOptionsHardwareDeviceId", "OgaConfigSetDecoderProviderOptionsHardwareVendorId",
		"OgaConfigClearDecoderProviderOptionsHardwareDeviceType", "OgaConfigClearDecoderProviderOptionsHardwareDeviceId",
		"OgaConfigClearDecoderProviderOptionsHardwareVendorId", "OgaConfigOverlay",
		"OgaModelGetType", "OgaModelGetDeviceType", "OgaStreamingProcessorProcess",
		"OgaStreamingProcessorFlush", "OgaCreateStreamingProcessor", "OgaDestroyStreamingProcessor",
		"OgaStreamingProcessorSetOption", "OgaStreamingProcessorGetOption",
	}
	for _, symbolName := range additionalRequiredSymbols {
		if createSym(handle, symbolName) == nil {
			C.dlclose(handle)
			return fmt.Errorf("missing required ORT GenAI symbol %s", symbolName)
		}
	}

	if rc := C.SetGenAiApi(
		symCreate, symErr, symDestroyRes, symDestroyModel, symCreateTokenizer, symDestroyTokenizer,
		symCreateTokenizerStream, symDestroyTokenizerStream, symApplyChatTemplate, symDestroyString, symCreateSequence, symDestroySequence,
		symTokenizerEncode, symCreateGenerator, symDestroyGenerator, symCreateGeneratorParams, symDestroyGeneratorParams,
		symGeneratorParamsSetSearchNumber, symGeneratorAppendTokenSequences, symGeneratorSetInputs, symGeneratorGenerateNextToken, symGeneratorGetSequenceCount,
		symGeneratorGetSequenceData, symTokenizerStreamDecode, symIsDone, symTokenizerGetEosTokenIDs, symCreateConfig, symConfigClearProviders, symConfigAppendProvider, symConfigSetProviderOption, symCreateModelFromConfig, symDestroyConfig,
		symLoadImage, symLoadImages, symLoadImagesFromBuffers, symDestroyImages, symCreateMultiModalProcessor, symDestroyMultiModalProcessor, symProcessorProcessImages, symDestroyNamedTensors, symCreateStringArray, symDestroyStringArray, symStringArrayAddString,
		symProcessorProcessImagesAndPrompts, symGeneratorParamsSetGuidance,
		symShutdown, symSetTelemetry, symCreateConfigFromEp, symGetPad, symGetBot, symGetEot, symGetBor, symGetEor,
	); rc != 0 {
		C.dlclose(handle)
		return fmt.Errorf("SetGenAiApi failed with code %d", int(rc))
	}
	if rc := C.SetGenAiEngineApi(&engineSymbols[0], C.size_t(len(engineSymbols))); rc != 0 {
		C.dlclose(handle)
		return fmt.Errorf("setting required ORT GenAI engine API failed with code %d", int(rc))
	}

	genAiLibraryHandle = handle
	return nil
}
