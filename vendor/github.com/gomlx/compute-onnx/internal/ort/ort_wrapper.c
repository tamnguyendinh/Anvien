// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

#include "onnxruntime_c_api.h"
#include <stdint.h>
#include <stdlib.h>

typedef const OrtApiBase* (*OrtGetApiBaseFn)();

const OrtApi* GetOrtApi(void* get_api_base_ptr, uint32_t version) {
    OrtGetApiBaseFn get_api_base = (OrtGetApiBaseFn)get_api_base_ptr;
    const OrtApiBase* api_base = get_api_base();
    return api_base->GetApi(version);
}

const char* GetOrtVersion(void* get_api_base_ptr) {
    OrtGetApiBaseFn get_api_base = (OrtGetApiBaseFn)get_api_base_ptr;
    const OrtApiBase* api_base = get_api_base();
    return api_base->GetVersionString();
}

OrtStatus* wrapper_CreateEnv(const OrtApi* api, OrtLoggingLevel default_logging_level, const char* logid, OrtEnv** out) {
    return api->CreateEnv(default_logging_level, logid, out);
}

void wrapper_ReleaseEnv(const OrtApi* api, OrtEnv* env) {
    api->ReleaseEnv(env);
}

OrtStatus* wrapper_CreateSessionOptions(const OrtApi* api, OrtSessionOptions** out) {
    return api->CreateSessionOptions(out);
}

void wrapper_ReleaseSessionOptions(const OrtApi* api, OrtSessionOptions* options) {
    api->ReleaseSessionOptions(options);
}

OrtStatus* wrapper_CreateSessionWithONNXData(const OrtApi* api, const OrtEnv* env, const void* model_data, size_t model_data_length, const OrtSessionOptions* options, OrtSession** out) {
    return api->CreateSessionFromArray(env, model_data, model_data_length, options, out);
}

void wrapper_ReleaseSession(const OrtApi* api, OrtSession* session) {
    api->ReleaseSession(session);
}

OrtStatus* wrapper_SessionGetInputCount(const OrtApi* api, const OrtSession* session, size_t* out) {
    return api->SessionGetInputCount(session, out);
}

OrtStatus* wrapper_SessionGetOutputCount(const OrtApi* api, const OrtSession* session, size_t* out) {
    return api->SessionGetOutputCount(session, out);
}

OrtStatus* wrapper_SessionGetInputName(const OrtApi* api, const OrtSession* session, size_t index, OrtAllocator* allocator, char** value) {
    return api->SessionGetInputName(session, index, allocator, value);
}

OrtStatus* wrapper_SessionGetOutputName(const OrtApi* api, const OrtSession* session, size_t index, OrtAllocator* allocator, char** value) {
    return api->SessionGetOutputName(session, index, allocator, value);
}

OrtStatus* wrapper_CreateCpuMemoryInfo(const OrtApi* api, OrtAllocatorType type, OrtMemType mem_type, OrtMemoryInfo** out) {
    return api->CreateCpuMemoryInfo(type, mem_type, out);
}

void wrapper_ReleaseMemoryInfo(const OrtApi* api, OrtMemoryInfo* info) {
    api->ReleaseMemoryInfo(info);
}

OrtStatus* wrapper_CreateTensorWithDataAsOrtValue(const OrtApi* api, const OrtMemoryInfo* info, void* p_data, size_t p_data_len, const int64_t* shape, size_t shape_len, ONNXTensorElementDataType type, OrtValue** out) {
    return api->CreateTensorWithDataAsOrtValue(info, p_data, p_data_len, shape, shape_len, type, out);
}

OrtStatus* wrapper_CreateTensorAsOrtValue(const OrtApi* api, OrtAllocator* allocator, const int64_t* shape, size_t shape_len, ONNXTensorElementDataType type, OrtValue** out) {
    return api->CreateTensorAsOrtValue(allocator, shape, shape_len, type, out);
}

OrtStatus* wrapper_GetTensorMutableData(const OrtApi* api, OrtValue* value, void** out) {
    return api->GetTensorMutableData(value, out);
}

void wrapper_ReleaseValue(const OrtApi* api, OrtValue* value) {
    api->ReleaseValue(value);
}

OrtStatus* wrapper_Run(const OrtApi* api, OrtSession* session, const OrtRunOptions* run_options, const char* const* input_names, const OrtValue* const* input_values, size_t input_count, const char* const* output_names, size_t output_count, OrtValue** output_values) {
    return api->Run(session, run_options, input_names, input_values, input_count, output_names, output_count, output_values);
}

const char* wrapper_GetErrorMessage(const OrtApi* api, const OrtStatus* status) {
    return api->GetErrorMessage(status);
}

void wrapper_ReleaseStatus(const OrtApi* api, OrtStatus* status) {
    api->ReleaseStatus(status);
}

OrtStatus* wrapper_GetAllocatorWithDefaultOptions(const OrtApi* api, OrtAllocator** out) {
    return api->GetAllocatorWithDefaultOptions(out);
}

OrtStatus* wrapper_CreateCUDAProviderOptions(const OrtApi* api, OrtCUDAProviderOptionsV2** out) {
    return api->CreateCUDAProviderOptions(out);
}

OrtStatus* wrapper_UpdateCUDAProviderOptions(const OrtApi* api, OrtCUDAProviderOptionsV2* cuda_options, const char* const* keys, const char* const* values, size_t num_keys) {
    return api->UpdateCUDAProviderOptions(cuda_options, keys, values, num_keys);
}

OrtStatus* wrapper_AddInitializer(const OrtApi* api, OrtSessionOptions* options, const char* name, const OrtValue* val) {
    return api->AddInitializer(options, name, val);
}

OrtStatus* wrapper_SessionOptionsAppendExecutionProvider_CUDA_V2(const OrtApi* api, OrtSessionOptions* options, const OrtCUDAProviderOptionsV2* cuda_options) {
    return api->SessionOptionsAppendExecutionProvider_CUDA_V2(options, cuda_options);
}

void wrapper_ReleaseCUDAProviderOptions(const OrtApi* api, OrtCUDAProviderOptionsV2* input) {
    api->ReleaseCUDAProviderOptions(input);
}

OrtStatus* wrapper_SessionOptionsAppendExecutionProvider_MIGraphX(const OrtApi* api, OrtSessionOptions* options, const OrtMIGraphXProviderOptions* migraphx_options) {
    return api->SessionOptionsAppendExecutionProvider_MIGraphX(options, migraphx_options);
}

int wrapper_HasMIGraphXSupport(const OrtApi* api) {
    return api->SessionOptionsAppendExecutionProvider_MIGraphX != NULL;
}

OrtStatus* wrapper_GetTensorTypeAndShape(const OrtApi* api, const OrtValue* value, OrtTensorTypeAndShapeInfo** out) {
    return api->GetTensorTypeAndShape(value, out);
}

OrtStatus* wrapper_GetTensorElementType(const OrtApi* api, const OrtTensorTypeAndShapeInfo* info, ONNXTensorElementDataType* out) {
    return api->GetTensorElementType(info, (uint32_t*)out);
}

OrtStatus* wrapper_GetDimensionsCount(const OrtApi* api, const OrtTensorTypeAndShapeInfo* info, size_t* out) {
    return api->GetDimensionsCount(info, out);
}

OrtStatus* wrapper_GetDimensions(const OrtApi* api, const OrtTensorTypeAndShapeInfo* info, int64_t* dim_values, size_t dim_values_length) {
    return api->GetDimensions(info, dim_values, dim_values_length);
}

void wrapper_ReleaseTensorTypeAndShapeInfo(const OrtApi* api, OrtTensorTypeAndShapeInfo* info) {
    api->ReleaseTensorTypeAndShapeInfo(info);
}

OrtStatus* wrapper_SetSessionLogSeverityLevel(const OrtApi* api, OrtSessionOptions* options, int session_log_severity_level) {
    return api->SetSessionLogSeverityLevel(options, session_log_severity_level);
}

// --- IoBinding wrappers ---

OrtStatus* wrapper_CreateIoBinding(const OrtApi* api, OrtSession* session, OrtIoBinding** out) {
    return api->CreateIoBinding(session, out);
}

void wrapper_ReleaseIoBinding(const OrtApi* api, OrtIoBinding* binding) {
    api->ReleaseIoBinding(binding);
}

OrtStatus* wrapper_BindInput(const OrtApi* api, OrtIoBinding* binding, const char* name, const OrtValue* val) {
    return api->BindInput(binding, name, val);
}

OrtStatus* wrapper_BindOutput(const OrtApi* api, OrtIoBinding* binding, const char* name, const OrtValue* val) {
    return api->BindOutput(binding, name, val);
}

OrtStatus* wrapper_BindOutputToDevice(const OrtApi* api, OrtIoBinding* binding, const char* name, const OrtMemoryInfo* mem_info) {
    return api->BindOutputToDevice(binding, name, mem_info);
}

OrtStatus* wrapper_RunWithBinding(const OrtApi* api, OrtSession* session, const OrtRunOptions* run_options, const OrtIoBinding* binding) {
    return api->RunWithBinding(session, run_options, binding);
}

OrtStatus* wrapper_GetBoundOutputValues(const OrtApi* api, const OrtIoBinding* binding, OrtAllocator* allocator, OrtValue*** output, size_t* output_count) {
    return api->GetBoundOutputValues(binding, allocator, output, output_count);
}

void wrapper_ClearBoundInputs(const OrtApi* api, OrtIoBinding* binding) {
    api->ClearBoundInputs(binding);
}

void wrapper_ClearBoundOutputs(const OrtApi* api, OrtIoBinding* binding) {
    api->ClearBoundOutputs(binding);
}

OrtStatus* wrapper_SynchronizeBoundInputs(const OrtApi* api, OrtIoBinding* binding) {
    return api->SynchronizeBoundInputs(binding);
}

OrtStatus* wrapper_SynchronizeBoundOutputs(const OrtApi* api, OrtIoBinding* binding) {
    return api->SynchronizeBoundOutputs(binding);
}

// --- MemoryInfo ---

OrtStatus* wrapper_CreateMemoryInfo(const OrtApi* api, const char* name, enum OrtAllocatorType type, int id, enum OrtMemType mem_type, OrtMemoryInfo** out) {
    return api->CreateMemoryInfo(name, type, id, mem_type, out);
}

// --- Allocator from session ---

OrtStatus* wrapper_CreateAllocator(const OrtApi* api, const OrtSession* session, const OrtMemoryInfo* mem_info, OrtAllocator** out) {
    return api->CreateAllocator(session, mem_info, out);
}

void wrapper_ReleaseAllocator(const OrtApi* api, OrtAllocator* allocator) {
    api->ReleaseAllocator(allocator);
}

// --- AllocatorAlloc / AllocatorFree for cudaMemcpy alternative ---

OrtStatus* wrapper_AllocatorAlloc(const OrtApi* api, OrtAllocator* allocator, size_t size, void** out) {
    return api->AllocatorAlloc(allocator, size, out);
}

OrtStatus* wrapper_AllocatorFree(const OrtApi* api, OrtAllocator* allocator, void* p) {
    return api->AllocatorFree(allocator, p);
}
