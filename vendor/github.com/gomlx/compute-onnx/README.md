# compute-onnx

[![Documentation](https://img.shields.io/badge/docs-gomlx.github.io-blue.svg)](https://gomlx.github.io/)
[![Sponsor GoMLX](https://img.shields.io/badge/Sponsor-GoMLX-white?logo=github&style=flat-square)](https://github.com/gomlx/gomlx/blob/main/README.md#-support-the-project)

**ONNX Runtime** based compute backend for GoMLX.

It allows GoMLX models to be executed via ONNX Runtime using:
- **Native Desktop/Server**: CPU, CUDA (NVIDIA GPU), or MIGraphX (AMD ROCm GPU).
- **WebAssembly / Browser**: WebGPU, CPU (WASM SIMD), WebNN, or WebGL in browsers (see [ORT Web Guide](docs/ort-web.md)).

It supports dynamic shapes and exporting models to `.onnx` files.

## Example Usage

To run the Adult dataset demo with the ONNX backend:

```bash
GOMLX_BACKEND=onnx go run -tags=onnx github.com/gomlx/gomlx/examples/adult/demo
```

Or targeting a specific accelerator (e.g. CUDA):

```bash
GOMLX_BACKEND=onnx:cuda go run -tags=onnx github.com/gomlx/gomlx/examples/adult/demo
```

Or on an AMD GPU (ROCm + MIGraphX):

```bash
GOMLX_BACKEND=onnx:migraphx go run -tags=onnx github.com/gomlx/gomlx/examples/adult/demo
```

## Backend Options & Configuration

Configuration can be specified in the `GOMLX_BACKEND` environment variable using `onnx:<options>` or `onnxruntime:<options>` (comma-separated).

### Accelerator Selection

- **`cpu`**: Force CPU execution.
  ```bash
  GOMLX_BACKEND=onnx:cpu
  ```
- **`cuda`** / **`gpu`**: Force CUDA GPU execution (uses ONNX Runtime CUDA Execution Provider via `OrtIoBinding`).
  ```bash
  GOMLX_BACKEND=onnx:cuda
  ```
- **`migraphx`** / **`rocm`** / **`amd`** (EXPERIMENTAL): Force AMD GPU execution (uses ONNX Runtime MIGraphX Execution Provider).
  Requires ROCm and MIGraphX to be installed (`sudo apt install migraphx migraphx-dev half`). If no ORT library with the
  MIGraphX provider is found, one is automatically installed from [AMD's manylinux wheels](https://repo.radeon.com/rocm/manylinux/)
  matching the local ROCm version — it can also be installed manually with
  `go run github.com/gomlx/compute-onnx/cmd/onnxruntime_installer -migraphx`.
  Notes:
  - Only float32/int32/int64 graphs are supported;
  - Models with scalar (0-dimensional) inputs fails (there is a bug upstream);
  - Buffers are always tranferred back to CPU (host), which makes training slower (TODO fix).
  ```bash
  GOMLX_BACKEND=onnx:migraphx
  ```
- **Custom Library Path**: Specify an explicit path to the ONNX Runtime `.so` (or `.dylib` / `.dll`) shared library file. This explicitly bypasses `ONNXRUNTIME_SHARED_LIBRARY_PATH`.
  ```bash
  GOMLX_BACKEND=onnx:/path/to/libonnxruntime.so
  ```
- **empty (default)**: Automatically detects if an NVIDIA GPU is present via `nvidia-smi` and defaults to CUDA if available, then checks for a discrete AMD GPU (ROCm/MIGraphX), otherwise falling back to CPU.
  ```bash
  GOMLX_BACKEND=onnx
  ```

### Save Model To ONNX

This allows one to export GOMLX trained (or fine-tuned) models to ONNX.

See [an example in UCI-Adult demo](https://github.com/gomlx/gomlx/blob/main/examples/adult/demo/save_onnx.go). If you have a pre-trained file in a directory called `base`:

```
GOMLX_BACKEND=onnx go run -tags=onnx ./examples/adult/demo/ -checkpoint "base" -save_onnx="/tmp/a.onnx" -vmodule=save_onnx=1
```

## ONNX Runtime Shared Libraries & Auto-Installation

The backend automatically locates or manages the required ONNX Runtime shared library (`libonnxruntime.so` / `onnxruntime.dll`):

- **Custom Library Path**: Set the `ONNXRUNTIME_SHARED_LIBRARY_PATH` environment variable or pass an explicit library path in the backend configuration (e.g. `GOMLX_BACKEND=onnx:/path/to/libonnxruntime.so`) to point directly to the shared library binary. Passing an explicit path in the configuration bypasses `ONNXRUNTIME_SHARED_LIBRARY_PATH`.
- **Auto-Installation**: If no library path is provided, the backend automatically downloads and extracts prebuilt official ONNX Runtime binaries locally (e.g. `~/.local/lib/onnxruntime/` on Linux).
- **Disabling Auto-Installation**: Set the environment variable `GOMLX_NO_AUTO_INSTALL=1` or call `onnxbackend.EnableAutoInstall(false)` programmatically to disable automatic downloads (useful for offline environments or container deployments).

## AMD ROCm / MIGraphX Environment Variables

The MIGraphX execution provider relies on a local ROCm installation:

- **`ROCM_PATH`**: Directory where ROCm is installed (defaults to `/opt/rocm`). It is used to locate `rocminfo` and the HIP/MIGraphX libraries when auto-detecting an AMD GPU and its ROCm version.
- **`GOMLX_MIGRAPHX_CACHE_DIR`**: Directory where the MIGraphX compiled-program (`.mxr`) for each model is cached, skipping the expensive graph compilation on subsequent runs. Equivalent to the `migraphx_cache_dir` config key (e.g. `GOMLX_BACKEND=onnx:migraphx,migraphx_cache_dir=/tmp/mxr`); an empty value disables caching.

## Debugging

### Saving Failed Models (`GOMLX_ONNX_SAVE_ON_FAILURE`)

If graph compilation or session creation fails in ONNX Runtime, setting the `GOMLX_ONNX_SAVE_ON_FAILURE` environment variable instructs the backend to automatically save the serialized ONNX model protobuf to the specified file path before returning the compilation error:

```bash
GOMLX_ONNX_SAVE_ON_FAILURE="/tmp/failed_model.onnx" go run -tags=onnx ...
```

This allows you to inspect the invalid graph using `onnx_printer` or [Netron](https://netron.app/) to diagnose the failure.

### Inspecting `.onnx` Files (`onnx_printer`)

This repository includes a CLI tool in `cmd/onnx_printer` to inspect and pretty-print `.onnx` model files in the terminal:

```bash
go run github.com/gomlx/compute-onnx/cmd/onnx_printer path/to/model.onnx
```

It formats input, output, and node tensor shapes using GoMLX `shapes.Shape` (including named dynamic axes) and prints each graph operation on a single line. Tensor constants and initializers are truncated to 10 elements by default (controlled via `-max_items` / `-n`).

Example usage:
```bash
# Print model details with a maximum of 5 items for constant values
go run github.com/gomlx/compute-onnx/cmd/onnx_printer -max_items 5 /tmp/model.onnx

# Read from stdin
cat /tmp/model.onnx | go run github.com/gomlx/compute-onnx/cmd/onnx_printer
```

> **Tip**: For interactive graphical visualization of ONNX models, you can open `.onnx` model files using [Netron](https://netron.app/).


### Logging & Verbosity

- **Backend Log Level (`log=<level>`)**: Configures ONNX Runtime's internal logging severity level.
  - `log=0`: Errors only (severity level 3 / ERROR)
  - `log=1`: Warnings (severity level 2 / WARNING)
  - `log=2`: Informational (severity level 1 / INFO)
  - `log=3`: Verbose (severity level 0 / VERBOSE)
  
  Example:
  ```bash
  GOMLX_BACKEND="onnx:cuda,log=2"
  ```

- **Execution Timing Log (`-vmodule=executable=1`)**: Enables per-step execution timing breakdown printed via `klog` using `humanize.Duration`.
  
  Example:
  ```bash
  GOMLX_BACKEND=onnx:cuda go run -tags=onnx github.com/gomlx/gomlx/examples/adult/demo -vmodule=executable=1
  ```

##  💖 Thanks

* [Go](golang.org)
* [ONNX](https://onnx.ai/)
* [ONNX Runtime](https://onnxruntime.ai/)
