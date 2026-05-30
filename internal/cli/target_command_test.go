package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/repo"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestContextAndImpactChildCommandsUseTargetTypes(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "context", "file", "src/app.go", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("context file returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("context file wrote stderr: %q", errOut)
	}
	var fileContext map[string]any
	if err := json.Unmarshal([]byte(out), &fileContext); err != nil {
		t.Fatalf("parse context file JSON: %v\n%s", err, out)
	}
	if fileContext["targetType"] != "file" || fileContext["dispatchMode"] != "explicit" {
		t.Fatalf("context file dispatch fields = %#v", fileContext)
	}
	nested, _ := fileContext["fileContext"].(map[string]any)
	summary, _ := nested["summary"].(map[string]any)
	if summary["path"] != "src/app.go" {
		t.Fatalf("context file summary = %#v", summary)
	}

	out, errOut, err = executeForTest(t, "context", "symbol", "NewServer", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("context symbol returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	var symbolContext map[string]any
	if err := json.Unmarshal([]byte(out), &symbolContext); err != nil {
		t.Fatalf("parse context symbol JSON: %v\n%s", err, out)
	}
	if symbolContext["targetType"] != "symbol" || symbolContext["dispatchMode"] != "explicit" {
		t.Fatalf("context symbol dispatch fields = %#v", symbolContext)
	}
	fileSummary, _ := symbolContext["fileSummary"].(map[string]any)
	if fileSummary["path"] != "src/app.go" {
		t.Fatalf("context symbol file summary = %#v", fileSummary)
	}

	out, errOut, err = executeForTest(t, "impact", "file", "src/app.go", "--repo", "fixture", "--direction", "downstream", "--depth", "1", "--json")
	if err != nil {
		t.Fatalf("impact file returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	var fileImpact map[string]any
	if err := json.Unmarshal([]byte(out), &fileImpact); err != nil {
		t.Fatalf("parse impact file JSON: %v\n%s", err, out)
	}
	if fileImpact["targetType"] != "file" || fileImpact["dispatchMode"] != "explicit" {
		t.Fatalf("impact file dispatch fields = %#v", fileImpact)
	}
	if impactIntFromJSON(fileImpact["impactedCount"]) == 0 {
		t.Fatalf("impact file should aggregate impacted symbols: %#v", fileImpact)
	}
	if !strings.Contains(out, "src/store.go") {
		t.Fatalf("impact file output missing affected file:\n%s", out)
	}
}

func TestQueryDetectChangesAndGraphHealthChildViews(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := initGitRepo(t)
	writeCLITestFile(t, repoPath, "src/app.go", numberedGoSource(18))
	runGit(t, repoPath, "add", "src/app.go")
	runGit(t, repoPath, "commit", "-m", "add app source")
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeFileProjectionCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "query", "files", "app", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("query files returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("query files wrote stderr: %q", errOut)
	}
	var fileQuery map[string]any
	if err := json.Unmarshal([]byte(out), &fileQuery); err != nil {
		t.Fatalf("parse query files JSON: %v\n%s", err, out)
	}
	if fileQuery["targetType"] != "files" || !strings.Contains(out, "src/app.go") {
		t.Fatalf("query files payload = %#v\n%s", fileQuery, out)
	}

	out, errOut, err = executeForTest(t, "query", "symbols", "NewServer", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("query symbols returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if !strings.Contains(out, `"targetType": "symbols"`) || !strings.Contains(out, `"fileSummary"`) {
		t.Fatalf("query symbols output missing target/file summary:\n%s", out)
	}

	out, errOut, err = executeForTest(t, "graph-health", "files", "--repo", "fixture", "--json", "--limit", "2")
	if err != nil {
		t.Fatalf("graph-health files returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if !strings.Contains(out, `"files"`) || !strings.Contains(out, `"src/app.go"`) {
		t.Fatalf("graph-health files output missing file rows:\n%s", out)
	}

	appPath := filepath.Join(repoPath, "src", "app.go")
	raw, err := os.ReadFile(appPath)
	if err != nil {
		t.Fatalf("read app source: %v", err)
	}
	lines := strings.Split(string(raw), "\n")
	lines[11] = "line12 changed"
	if err := os.WriteFile(appPath, []byte(strings.Join(lines, "\n")), 0o644); err != nil {
		t.Fatalf("modify app source: %v", err)
	}

	out, errOut, err = executeForTest(t, "detect-changes", "files", "--scope", "all", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("detect-changes files returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if !strings.Contains(out, `"targetType": "files"`) || !strings.Contains(out, `"changed_files"`) || !strings.Contains(out, `"src/app.go"`) {
		t.Fatalf("detect-changes files output missing changed file view:\n%s", out)
	}
}

func TestParentContextReportsAmbiguousFileAndSymbolTargets(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.go", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.go", "filePath": "src/app.go",
	}})
	g.AddNode(graph.Node{ID: "File:src/other.go", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "other.go", "filePath": "src/other.go",
	}})
	g.AddNode(graph.Node{ID: "Function:src/other.go:weird", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "src/app.go", "filePath": "src/other.go", "startLine": 1, "endLine": 2,
	}})
	g.AddRelationship(graph.Relationship{ID: "rel:def-weird", SourceID: "File:src/other.go", TargetID: "Function:src/other.go:weird", Type: graph.RelDefines})
	writeGroupCommandGraph(t, repoPath, g)

	out, errOut, err := executeForTest(t, "context", "src/app.go", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("context parent ambiguous returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if !strings.Contains(out, `"status": "ambiguous"`) ||
		!strings.Contains(out, `anvien context file`) ||
		!strings.Contains(out, `anvien context symbol`) {
		t.Fatalf("context parent ambiguous output missing suggestions:\n%s", out)
	}
}

func numberedGoSource(lines int) string {
	out := make([]string, 0, lines)
	for i := 1; i <= lines; i++ {
		out = append(out, "line"+strconvItoa(i))
	}
	return strings.Join(out, "\n") + "\n"
}

func strconvItoa(value int) string {
	if value == 0 {
		return "0"
	}
	digits := make([]byte, 0, 10)
	for value > 0 {
		digits = append(digits, byte('0'+value%10))
		value /= 10
	}
	for left, right := 0, len(digits)-1; left < right; left, right = left+1, right-1 {
		digits[left], digits[right] = digits[right], digits[left]
	}
	return string(digits)
}

func impactIntFromJSON(value any) int {
	number, _ := value.(float64)
	return int(number)
}
