package gitignore

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const (
	StartMarker = "# anvien:start"
	EndMarker   = "# anvien:end"
)

var ManagedEntries = []string{
	"# Anvien local index and generated AI context",
	".gitignore",
	".anvien/",
	"AGENTS.md",
	"CLAUDE.md",
	"",
	"# Agent/editor local state used by Anvien integrations",
	".claude/",
	".codex/",
	".agents/",
	".grok/",
	".claude-flow/",
	".claude-plugin/",
	".history/",
	".swarm/",
	".codex-tmp/",
	".tmp/",
	"undefined/",
	"local_docs/",
	"",
	"# Local Anvien launcher/runtime generated artifacts",
	"anvien-launcher/web-dist/",
	"anvien-launcher/AnvienLauncher.exe",
	"anvien-launcher/server-bundle/",
	"anvien-launcher/logs/",
}

type Result struct {
	Path    string
	Changed bool
}

func Ensure(repoPath string) (Result, error) {
	gitignorePath := filepath.Join(repoPath, ".gitignore")
	raw, err := os.ReadFile(gitignorePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Result{}, err
	}

	next := Upsert(string(raw))
	if err == nil && next == string(raw) {
		return Result{Path: gitignorePath}, nil
	}
	if err := os.WriteFile(gitignorePath, []byte(next), 0o644); err != nil {
		return Result{}, err
	}
	return Result{Path: gitignorePath, Changed: true}, nil
}

func Upsert(text string) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	block := strings.TrimRight(ManagedBlock(), "\n")
	if strings.TrimSpace(text) == "" {
		return block + "\n"
	}

	start := strings.Index(text, StartMarker)
	if start >= 0 {
		afterStart := text[start+len(StartMarker):]
		endRelative := strings.Index(afterStart, EndMarker)
		if endRelative >= 0 {
			end := start + len(StartMarker) + endRelative + len(EndMarker)
			next := strings.TrimRight(text[:start], " \t\r\n")
			if next != "" {
				next += "\n\n"
			}
			next += block
			tail := strings.TrimLeft(text[end:], " \t\r\n")
			if tail != "" {
				next += "\n\n" + strings.TrimRight(tail, " \t\r\n")
			}
			return next + "\n"
		}
	}

	return strings.TrimRight(text, " \t\r\n") + "\n\n" + block + "\n"
}

func ManagedBlock() string {
	return StartMarker + "\n" + strings.Join(ManagedEntries, "\n") + "\n" + EndMarker + "\n"
}
