//go:build windows

package httpapi

import (
	"context"
	"os/exec"
	"testing"
)

func TestFolderPickerCommandRunsWithoutConsoleWindow(t *testing.T) {
	cmd := exec.CommandContext(context.Background(), "powershell.exe", "-NoProfile")
	configureFolderPickerCommand(cmd)
	if cmd.SysProcAttr == nil {
		t.Fatalf("folder picker command SysProcAttr = nil")
	}
	if !cmd.SysProcAttr.HideWindow {
		t.Fatalf("folder picker command HideWindow = false, want true")
	}
	if cmd.SysProcAttr.CreationFlags&0x08000000 == 0 {
		t.Fatalf("folder picker command CreationFlags = %#x, want CREATE_NO_WINDOW", cmd.SysProcAttr.CreationFlags)
	}
}
