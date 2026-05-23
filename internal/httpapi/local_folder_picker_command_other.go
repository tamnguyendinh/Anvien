//go:build !windows

package httpapi

import "os/exec"

func configureFolderPickerCommand(cmd *exec.Cmd) {}
