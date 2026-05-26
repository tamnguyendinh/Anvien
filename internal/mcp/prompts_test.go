package mcp

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGenerateMapPromptWithRepoUsesEscapedResourcesAndEvidenceRules(t *testing.T) {
	message := generateMapPrompt(map[string]any{"repo": "my project"})
	text := message.Content.Text

	for _, want := range []string{
		"Use the supplied repo name exactly as `my project`",
		"avmatrix://repo/my%20project/context",
		"avmatrix://repo/my%20project/clusters",
		"avmatrix://repo/my%20project/processes",
		"avmatrix://repo/my%20project/process/<process-uri>",
		"avmatrix analyze --force",
		"Write architecture claims only from resources, MCP tools, CLI commands, or Web/API output you actually read.",
		"Do not invent Mermaid nodes, edges, dependencies, layers, or ownership.",
		"Choose up to 5 representative processes",
		"Record the reason each process was selected",
		"Edit or create repository files only if the user explicitly asked for file changes.",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("generate_map prompt with repo missing %q:\n%s", want, text)
		}
	}
	if strings.Contains(text, "{name}") {
		t.Fatalf("generate_map prompt with repo contains stale placeholder:\n%s", text)
	}
}

func TestGenerateMapPromptWithoutRepoRequiresDiscoveryAndAmbiguityStop(t *testing.T) {
	result, err := getPrompt(mustPromptParams(t, "generate_map", map[string]any{}))
	if err != nil {
		t.Fatalf("getPrompt(generate_map): %v", err)
	}
	text := promptTextFromResult(t, result)
	for _, want := range []string{
		"first READ `avmatrix://repos`",
		"If exactly one repo is listed, use that repo name.",
		"indexed path clearly matches the current workspace",
		"stop and ask the user which repo to map",
		"URL-escape the exact repo name",
		"Never continue with a placeholder resource URI.",
		"avmatrix://repo/<resolved-repo-uri>/context",
		"avmatrix://repo/<resolved-repo-uri>/process/<process-uri>",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("generate_map prompt without repo missing %q:\n%s", want, text)
		}
	}
	if strings.Contains(text, "{name}") {
		t.Fatalf("generate_map prompt without repo contains stale placeholder:\n%s", text)
	}
}

func TestDetectImpactPromptUsesCurrentWorkflowLanguage(t *testing.T) {
	message := detectImpactPrompt(map[string]any{"scope": "compare", "base_ref": "main"})
	text := message.Content.Text
	for _, want := range []string{
		"detect_changes(",
		`"base_ref":"main"`,
		`"scope":"compare"`,
		"avmatrix analyze --force",
		"avmatrix detect-changes --scope compare --base-ref main --repo <repo>",
		"HIGH or CRITICAL is a blast-radius warning",
		"not an edit ban",
		"run detect-changes again before committing",
		"backed only by command output you actually read",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("detect_impact prompt missing %q:\n%s", want, text)
		}
	}
}

func mustPromptParams(t *testing.T, name string, arguments map[string]any) json.RawMessage {
	t.Helper()
	raw, err := json.Marshal(map[string]any{
		"name":      name,
		"arguments": arguments,
	})
	if err != nil {
		t.Fatalf("marshal prompt params: %v", err)
	}
	return raw
}

func promptTextFromResult(t *testing.T, result map[string]any) string {
	t.Helper()
	messages := result["messages"].([]promptMessage)
	if len(messages) != 1 {
		t.Fatalf("messages = %#v", messages)
	}
	return messages[0].Content.Text
}
