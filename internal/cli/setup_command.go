package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tamnguyendinh/avmatrix-go/internal/aicontext"
	"github.com/tamnguyendinh/avmatrix-go/internal/repo"
)

const (
	setupBrand         = "AVmatrix"
	setupCommandName   = "avmatrix"
	setupMCPServerName = "avmatrix"
)

type setupResult struct {
	Configured []string
	Skipped    []string
	Errors     []string
}

type setupMCPEntry struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func newSetupCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "setup",
		Short: "Configure local MCP/runtime access for supported AI editors",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			result, err := runSetup()
			if err != nil {
				return err
			}
			printSetupResult(cmd, result)
			if len(result.Errors) > 0 {
				return errors.New("setup completed with errors")
			}
			return nil
		},
	}
}

func runSetup() (setupResult, error) {
	if err := os.MkdirAll(repo.GlobalDir(), 0o755); err != nil {
		return setupResult{}, err
	}

	result := setupResult{}
	setupCursor(&result)
	setupClaudeCode(&result)
	setupOpenCode(&result)
	setupCodex(&result)
	return result, nil
}

func printSetupResult(cmd *cobra.Command, result setupResult) {
	out := cmd.OutOrStdout()
	fmt.Fprintln(out)
	fmt.Fprintf(out, "  %s Setup\n", setupBrand)
	fmt.Fprintln(out, "  ===============")
	fmt.Fprintln(out)

	if len(result.Configured) > 0 {
		fmt.Fprintln(out, "  Configured:")
		for _, name := range result.Configured {
			fmt.Fprintf(out, "    + %s\n", name)
		}
	}
	if len(result.Skipped) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "  Skipped:")
		for _, name := range result.Skipped {
			fmt.Fprintf(out, "    - %s\n", name)
		}
	}
	if len(result.Errors) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "  Errors:")
		for _, message := range result.Errors {
			fmt.Fprintf(out, "    ! %s\n", message)
		}
	}

	fmt.Fprintln(out)
	fmt.Fprintln(out, "  Summary:")
	fmt.Fprintf(out, "    MCP configured for: %s\n", setupConfiguredMCPNames(result.Configured))
	fmt.Fprintf(out, "    Skills installed to: %s\n", setupConfiguredSkillNames(result.Configured))
	fmt.Fprintf(out, "    MCP lifecycle: %s mcp is editor-owned and may stay running while the editor or agent session is active.\n", setupCommandName)
	fmt.Fprintf(out, "    Diagnostics: %s doctor locks, %s doctor processes\n", setupCommandName, setupCommandName)
	fmt.Fprintln(out)
	fmt.Fprintln(out, "  Next steps:")
	fmt.Fprintln(out, "    1. cd into any git repo")
	fmt.Fprintf(out, "    2. Run: %s analyze\n", setupCommandName)
	fmt.Fprintln(out, "    3. Open the repo in your editor; MCP is ready and owned by that editor session.")
	fmt.Fprintln(out)
}

func setupConfiguredMCPNames(configured []string) string {
	names := make([]string, 0)
	for _, name := range configured {
		lower := strings.ToLower(name)
		if strings.Contains(lower, "skills") || strings.Contains(lower, "hooks") {
			continue
		}
		names = append(names, name)
	}
	if len(names) == 0 {
		return "none"
	}
	return strings.Join(names, ", ")
}

func setupConfiguredSkillNames(configured []string) string {
	names := make([]string, 0)
	for _, name := range configured {
		if strings.Contains(strings.ToLower(name), "skills") {
			names = append(names, name)
		}
	}
	if len(names) == 0 {
		return "none"
	}
	return strings.Join(names, ", ")
}

func setupCursor(result *setupResult) {
	cursorDir := filepath.Join(setupHomeDir(), ".cursor")
	if !setupDirExists(cursorDir) {
		result.Skipped = append(result.Skipped, "Cursor (not installed)")
		return
	}
	if err := setupWriteMCPJSON(filepath.Join(cursorDir, "mcp.json")); err != nil {
		result.Errors = append(result.Errors, "Cursor: "+err.Error())
	} else {
		result.Configured = append(result.Configured, "Cursor")
	}
	setupInstallEditorSkills(result, "Cursor", filepath.Join(cursorDir, "skills"))
}

func setupClaudeCode(result *setupResult) {
	home := setupHomeDir()
	claudeDir := filepath.Join(home, ".claude")
	if !setupDirExists(claudeDir) {
		result.Skipped = append(result.Skipped, "Claude Code (not installed)")
		return
	}
	if err := setupWriteMCPJSON(filepath.Join(home, ".claude.json")); err != nil {
		result.Errors = append(result.Errors, "Claude Code: "+err.Error())
	} else {
		result.Configured = append(result.Configured, "Claude Code")
	}
	setupInstallEditorSkills(result, "Claude Code", filepath.Join(claudeDir, "skills"))
	setupInstallClaudeHooks(result, claudeDir)
}

func setupOpenCode(result *setupResult) {
	opencodeDir := filepath.Join(setupHomeDir(), ".config", "opencode")
	if !setupDirExists(opencodeDir) {
		result.Skipped = append(result.Skipped, "OpenCode (not installed)")
		return
	}
	if err := setupWriteOpenCodeJSON(filepath.Join(opencodeDir, "opencode.json")); err != nil {
		result.Errors = append(result.Errors, "OpenCode: "+err.Error())
	} else {
		result.Configured = append(result.Configured, "OpenCode")
	}
	setupInstallEditorSkills(result, "OpenCode", filepath.Join(opencodeDir, "skill"))
}

func setupCodex(result *setupResult) {
	home := setupHomeDir()
	codexDir := filepath.Join(home, ".codex")
	if !setupDirExists(codexDir) {
		result.Skipped = append(result.Skipped, "Codex (not installed)")
		return
	}

	if err := setupRunCodexMCPAdd(); err == nil {
		result.Configured = append(result.Configured, "Codex")
	} else if err := setupUpsertCodexToml(filepath.Join(codexDir, "config.toml")); err != nil {
		result.Errors = append(result.Errors, "Codex: "+err.Error())
	} else {
		result.Configured = append(result.Configured, "Codex (MCP added to ~/.codex/config.toml)")
	}
	setupInstallEditorSkills(result, "Codex", filepath.Join(home, ".agents", "skills"))
}

func setupRunCodexMCPAdd() error {
	cmd := exec.Command("codex", "mcp", "add", setupMCPServerName, "--", setupCommandName, "mcp")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "codex", "mcp", "add", setupMCPServerName, "--", setupCommandName, "mcp")
	}
	return cmd.Run()
}

func setupWriteMCPJSON(path string) error {
	config, err := setupReadJSONObject(path)
	if err != nil {
		return err
	}
	servers, _ := config["mcpServers"].(map[string]any)
	if servers == nil {
		servers = map[string]any{}
	}
	servers[setupMCPServerName] = map[string]any{
		"command": setupCommandName,
		"args":    []any{"mcp"},
	}
	config["mcpServers"] = servers
	return setupWriteJSONObject(path, config)
}

func setupWriteOpenCodeJSON(path string) error {
	config, err := setupReadJSONObject(path)
	if err != nil {
		return err
	}
	mcp, _ := config["mcp"].(map[string]any)
	if mcp == nil {
		mcp = map[string]any{}
	}
	mcp[setupMCPServerName] = map[string]any{
		"command": setupCommandName,
		"args":    []any{"mcp"},
	}
	config["mcp"] = mcp
	return setupWriteJSONObject(path, config)
}

func setupReadJSONObject(path string) (map[string]any, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]any{}, nil
		}
		return nil, err
	}
	var config map[string]any
	if err := json.Unmarshal(raw, &config); err != nil {
		return map[string]any{}, nil
	}
	if config == nil {
		config = map[string]any{}
	}
	return config, nil
}

func setupWriteJSONObject(path string, config map[string]any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

func setupUpsertCodexToml(path string) error {
	raw, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	text := string(raw)
	if strings.Contains(text, "[mcp_servers."+setupMCPServerName+"]") {
		return nil
	}
	section := "[mcp_servers." + setupMCPServerName + "]\ncommand = " + strconvQuote(setupCommandName) + "\nargs = [\"mcp\"]\n"
	next := strings.TrimRight(text, " \t\r\n")
	if strings.TrimSpace(next) != "" {
		next += "\n\n"
	}
	next += section
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(strings.TrimRight(next, " \t\r\n")+"\n"), 0o644)
}

func setupInstallEditorSkills(result *setupResult, editor string, targetDir string) {
	installed, err := setupInstallSkillsTo(targetDir)
	if err != nil {
		result.Errors = append(result.Errors, editor+" skills: "+err.Error())
		return
	}
	if len(installed) > 0 {
		result.Configured = append(result.Configured, fmt.Sprintf("%s skills (%d skills -> %s)", editor, len(installed), setupDisplayHomePath(targetDir)))
	}
}

func setupInstallSkillsTo(targetDir string) ([]string, error) {
	return aicontext.InstallBaseSkillsTo(targetDir)
}

func setupInstallClaudeHooks(result *setupResult, claudeDir string) {
	if err := setupMergeClaudeHookSettings(filepath.Join(claudeDir, "settings.json")); err != nil {
		result.Errors = append(result.Errors, "Claude Code hooks: "+err.Error())
		return
	}
	result.Configured = append(result.Configured, "Claude Code Go hooks (PreToolUse, PostToolUse)")
}

func setupMergeClaudeHookSettings(path string) error {
	config, err := setupReadJSONObject(path)
	if err != nil {
		return err
	}
	hooks, _ := config["hooks"].(map[string]any)
	if hooks == nil {
		hooks = map[string]any{}
	}
	command := setupCommandName + " hook claude"
	hooks["PreToolUse"] = setupHookEntriesWithoutExisting(hooks["PreToolUse"], "Grep|Glob|Bash", command, "Enriching with AVmatrix graph context...")
	hooks["PostToolUse"] = setupHookEntriesWithoutExisting(hooks["PostToolUse"], "Bash", command, "Checking AVmatrix index freshness...")
	config["hooks"] = hooks
	return setupWriteJSONObject(path, config)
}

func setupHookEntriesWithoutExisting(existing any, matcher string, command string, statusMessage string) []any {
	entries := make([]any, 0)
	if list, ok := existing.([]any); ok {
		for _, item := range list {
			if !setupHookEntryContains(item, "avmatrix-hook") && !setupHookEntryContains(item, setupCommandName+" hook claude") {
				entries = append(entries, item)
			}
		}
	}
	entries = append(entries, map[string]any{
		"matcher": matcher,
		"hooks": []any{
			map[string]any{
				"type":          "command",
				"command":       command,
				"timeout":       float64(10),
				"statusMessage": statusMessage,
			},
		},
	})
	return entries
}

func setupHookEntryContains(entry any, needle string) bool {
	object, ok := entry.(map[string]any)
	if !ok {
		return false
	}
	hooks, ok := object["hooks"].([]any)
	if !ok {
		return false
	}
	for _, item := range hooks {
		hook, ok := item.(map[string]any)
		if !ok {
			continue
		}
		command, _ := hook["command"].(string)
		if strings.Contains(command, needle) {
			return true
		}
	}
	return false
}

func setupHomeDir() string {
	if runtime.GOOS == "windows" {
		if home := strings.TrimSpace(os.Getenv("USERPROFILE")); home != "" {
			return home
		}
	}
	if home := strings.TrimSpace(os.Getenv("HOME")); home != "" {
		return home
	}
	if home := strings.TrimSpace(os.Getenv("USERPROFILE")); home != "" {
		return home
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return home
}

func setupDisplayHomePath(path string) string {
	home := filepath.Clean(setupHomeDir())
	clean := filepath.Clean(path)
	if rel, err := filepath.Rel(home, clean); err == nil && rel != "." && !strings.HasPrefix(rel, "..") {
		return "~/" + filepath.ToSlash(rel)
	}
	return filepath.ToSlash(clean)
}

func setupDirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func strconvQuote(value string) string {
	raw, _ := json.Marshal(value)
	return string(raw)
}
