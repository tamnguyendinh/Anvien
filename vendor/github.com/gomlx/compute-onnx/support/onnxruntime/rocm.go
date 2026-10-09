// Copyright 2026 The GoMLX Authors. SPDX-License-Identifier: Apache-2.0

package onnxruntime

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ROCmWheelRepoURL is the base URL of AMD's manylinux wheel repository, holding
// pre-built ONNX Runtime wheels with the MIGraphX execution provider enabled,
// one directory per ROCm release (e.g. .../rocm-rel-7.2.4/).
const ROCmWheelRepoURL = "https://repo.radeon.com/rocm/manylinux"

// GetDefaultROCmVersion returns the locally installed ROCm version (e.g. "7.2.4"),
// detected from $ROCM_PATH or /opt/rocm, or "" if ROCm is not installed.
func GetDefaultROCmVersion() string {
	root := os.Getenv("ROCM_PATH")
	if root == "" {
		root = "/opt/rocm"
	}
	for _, p := range []string{
		filepath.Join(root, ".info", "version"),
		filepath.Join(root, "lib", ".info", "version"),
	} {
		if data, err := os.ReadFile(p); err == nil {
			return strings.TrimSpace(string(data))
		}
	}
	return ""
}

var migraphxWheelRegexp = regexp.MustCompile(`onnxruntime_migraphx-([0-9][^/"<]*?)-cp([0-9]+)[^-]*-[^/"<>]*?\.whl`)

// FindMigraphxWheelURL lists the AMD wheel repository for the given ROCm release
// (e.g. "7.2.4") and returns the URL of the ONNX Runtime MIGraphX wheel to install.
func FindMigraphxWheelURL(rocmVersion string) (string, error) {
	if rocmVersion == "" {
		return "", errors.New("no ROCm version given or detected: cannot locate an onnxruntime-migraphx wheel")
	}
	listingURL := fmt.Sprintf("%s/rocm-rel-%s/", ROCmWheelRepoURL, rocmVersion)
	resp, err := http.Get(listingURL)
	if err != nil {
		return "", errors.Wrapf(err, "failed to list AMD wheel repository at %s", listingURL)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("failed to list AMD wheel repository at %s: HTTP %s", listingURL, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read wheel listing from %s", listingURL)
	}

	bestURL := ""
	bestCP := 0
	for _, m := range migraphxWheelRegexp.FindAllStringSubmatch(string(body), -1) {
		cpNum, _ := strconv.Atoi(m[2])
		// Prefer the highest CPython tag (the shared libraries are CPython independent,
		// but newer tags tend to be built with newer toolchains).
		if cpNum > bestCP {
			bestCP = cpNum
			bestURL = fmt.Sprintf("%s%s", listingURL, m[0])
		}
	}
	if bestURL == "" {
		return "", errors.Errorf("no onnxruntime-migraphx wheel found for ROCm %s at %s", rocmVersion, listingURL)
	}
	return bestURL, nil
}

// GetMigraphxInstallPath returns the dedicated install directory for the AMD ROCm/MIGraphX
// ONNX Runtime build (e.g. ~/.local/lib/onnxruntime-migraphx on Linux). It is kept separate
// from the standard install path so that installing it does not clobber the CPU/CUDA library
// (the AMD build is based on a different ORT version).
func GetMigraphxInstallPath() (string, error) {
	base, err := GetInstallPath()
	if err != nil {
		return "", err
	}
	return base + "-migraphx", nil
}

// InstallMigraphx downloads and installs a pre-built ONNX Runtime library with the
// MIGraphX (AMD ROCm) execution provider enabled, extracted from AMD's manylinux wheels.
// The rocmVersion parameter selects the ROCm release to fetch the matching wheel for
// (e.g. "7.2.4"); if empty, the locally installed ROCm version is detected.
// If targetDir is empty, the library is installed under the dedicated MIGraphX directory
// (see GetMigraphxInstallPath). It returns the absolute path to the main shared library file.
func InstallMigraphx(rocmVersion string, targetDir string, force bool) (string, error) {
	if rocmVersion == "" {
		rocmVersion = GetDefaultROCmVersion()
	}

	if targetDir == "" {
		var err error
		targetDir, err = GetMigraphxInstallPath()
		if err != nil {
			return "", err
		}
	} else {
		targetDir = filepath.Clean(targetDir)
	}

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", errors.Wrapf(err, "failed to create directory %s", targetDir)
	}

	libFilename, err := GetLibFilename()
	if err != nil {
		return "", err
	}

	targetPath := filepath.Join(targetDir, libFilename)
	migraphxLibPath := filepath.Join(targetDir, "libonnxruntime_providers_migraphx.so")
	if !force {
		if _, err := os.Stat(targetPath); err == nil {
			if _, err := os.Stat(migraphxLibPath); err == nil {
				return targetPath, nil
			}
		}
	}

	url, err := FindMigraphxWheelURL(rocmVersion)
	if err != nil {
		return "", err
	}
	wheelName := filepath.Base(url)

	fmt.Printf("Installing ONNX Runtime (MIGraphX EP, ROCm %s) to %s...\n", rocmVersion, targetDir)
	fmt.Printf("Downloading %s ...\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "failed to download from %s", url)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("failed to download from %s: HTTP %s", url, resp.Status)
	}

	// Wheels are zip files: extract all libonnxruntime* shared libraries from onnxruntime/capi/.
	tmpZipFile, err := os.CreateTemp("", "onnxruntime-migraphx-*.whl")
	if err != nil {
		return "", errors.Wrap(err, "failed to create temporary file")
	}
	defer func() {
		tmpZipFile.Close()
		os.Remove(tmpZipFile.Name())
	}()
	if _, err := io.Copy(tmpZipFile, resp.Body); err != nil {
		return "", errors.Wrapf(err, "failed to download %s", wheelName)
	}
	stat, err := tmpZipFile.Stat()
	if err != nil {
		return "", err
	}
	zr, err := zip.NewReader(tmpZipFile, stat.Size())
	if err != nil {
		return "", errors.Wrapf(err, "failed to open wheel %s", wheelName)
	}

	var mainLibVersioned string
	for _, file := range zr.File {
		baseName := filepath.Base(file.Name)
		dir := filepath.Dir(file.Name)
		if file.FileInfo().IsDir() ||
			!strings.Contains(dir, "capi") ||
			!strings.HasPrefix(baseName, "libonnxruntime") {
			continue
		}
		targetFile := filepath.Join(targetDir, baseName)
		f, err := os.OpenFile(targetFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
		if err != nil {
			return "", errors.Wrapf(err, "failed to create file %s", targetFile)
		}
		src, err := file.Open()
		if err != nil {
			f.Close()
			return "", errors.Wrapf(err, "failed to open zip entry %s", file.Name)
		}
		_, err = io.Copy(f, src)
		src.Close()
		f.Close()
		if err != nil {
			return "", errors.Wrapf(err, "failed to write file %s", targetFile)
		}
		fmt.Printf("Installed %s\n", targetFile)
		if strings.HasPrefix(baseName, "libonnxruntime.so.") && mainLibVersioned == "" {
			mainLibVersioned = baseName
		}
	}

	// Point the unversioned main library name at the versioned AMD library,
	// replacing any pre-existing symlink or file (e.g. a previously installed CPU/CUDA build).
	if mainLibVersioned != "" {
		existing, linkErr := os.Readlink(targetPath)
		if linkErr != nil || existing != mainLibVersioned {
			_ = os.Remove(targetPath)
			if err := os.Symlink(mainLibVersioned, targetPath); err != nil {
				return "", errors.Wrapf(err, "failed to create symlink %s -> %s", targetPath, mainLibVersioned)
			}
		}
	}
	if _, err := os.Stat(targetPath); err != nil {
		return "", errors.Errorf("no ONNX Runtime shared library found in wheel %s", wheelName)
	}
	return targetPath, nil
}
