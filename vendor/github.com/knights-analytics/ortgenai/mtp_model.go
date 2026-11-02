package ortgenai

import (
	"fmt"
	"os"
	"path/filepath"
)

// resolveMTPModelPath prefers the standalone MTP package over the root layout.
func resolveMTPModelPath(modelPath string) (string, error) {
	subfolder := filepath.Join(modelPath, "MTP")
	info, err := os.Stat(subfolder)
	if err == nil {
		if !info.IsDir() {
			return "", fmt.Errorf("MTP model path %q is not a directory", subfolder)
		}
		return subfolder, nil
	}
	if !os.IsNotExist(err) {
		return "", fmt.Errorf("checking MTP model directory: %w", err)
	}
	for _, name := range []string{"mtp.onnx", "mtp.onnx.data"} {
		path := filepath.Join(modelPath, name)
		info, err := os.Stat(path)
		if err != nil {
			return "", fmt.Errorf("MTP model requires an MTP subdirectory or root mtp.onnx and mtp.onnx.data: %w", err)
		}
		if !info.Mode().IsRegular() {
			return "", fmt.Errorf("MTP model file %q is not a regular file", path)
		}
	}
	return modelPath, nil
}
