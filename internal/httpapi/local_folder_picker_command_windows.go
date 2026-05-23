//go:build windows

package httpapi

import (
	"os/exec"
	"syscall"
)

func configureFolderPickerCommand(cmd *exec.Cmd) {
	const createNoWindow = 0x08000000
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: createNoWindow,
	}
}
