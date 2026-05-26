package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	backendPath, logFile, err := resolveRuntime()
	if err != nil {
		fatal(err)
	}
	if logFile != nil {
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	if _, err := os.Stat(backendPath); err != nil {
		fatal(fmt.Errorf("canonical AVmatrix CLI missing: %s", backendPath))
	}

	cmd := exec.Command(backendPath, "serve", "--host", "127.0.0.1", "--port", "4848")
	cmd.Dir = filepath.Dir(backendPath)
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	cmd.SysProcAttr = hiddenProcAttr()
	log.Printf("starting Go backend: %s serve --host 127.0.0.1 --port 4848", backendPath)
	if err := cmd.Run(); err != nil {
		fatal(err)
	}
}

func resolveRuntime() (string, *os.File, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", nil, err
	}
	exePath, err = filepath.Abs(exePath)
	if err != nil {
		return "", nil, err
	}
	bundleDir := filepath.Dir(exePath)
	launcherDir := filepath.Dir(bundleDir)
	rootDir := filepath.Dir(launcherDir)
	logDir := filepath.Join(launcherDir, "logs")
	if err := os.MkdirAll(logDir, 0o755); err != nil {
		return "", nil, err
	}
	logFile, err := os.OpenFile(filepath.Join(logDir, "server-wrapper.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return "", nil, err
	}
	return filepath.Join(rootDir, "avmatrix", "bin", "avmatrix.exe"), logFile, nil
}

func hiddenProcAttr() *syscall.SysProcAttr {
	if runtime.GOOS != "windows" {
		return &syscall.SysProcAttr{}
	}
	return &syscall.SysProcAttr{HideWindow: true, CreationFlags: 0x08000000}
}

func fatal(err error) {
	log.Printf("fatal: %v", err)
	os.Exit(1)
}
