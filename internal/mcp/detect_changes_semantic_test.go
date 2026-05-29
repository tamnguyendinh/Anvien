package mcp

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/repo"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestDetectChangesToolReturnsSemanticSummaries(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git executable is not available")
	}

	store := repo.NewStore(t.TempDir())
	repoPath := t.TempDir()
	sourceDir := filepath.Join(repoPath, "src")
	if err := os.MkdirAll(sourceDir, 0o755); err != nil {
		t.Fatalf("mkdir source: %v", err)
	}
	sourcePath := filepath.Join(sourceDir, "app.ts")
	if err := os.WriteFile(sourcePath, []byte("function main() {\n  return missing()\n}\n"), 0o644); err != nil {
		t.Fatalf("write source: %v", err)
	}
	runMCPTestGit(t, repoPath, "init")
	runMCPTestGit(t, repoPath, "config", "user.email", "test@example.com")
	runMCPTestGit(t, repoPath, "config", "user.name", "Test User")
	runMCPTestGit(t, repoPath, "add", "src/app.ts")
	runMCPTestGit(t, repoPath, "commit", "-m", "initial")

	meta := repo.Meta{
		RepoPath:   repoPath,
		IndexedAt:  "2026-05-22T00:00:00Z",
		LastCommit: "abc123",
		Stats:      &repo.Stats{},
	}
	if err := repo.SaveMeta(repo.StoragePath(repoPath), meta); err != nil {
		t.Fatalf("save meta: %v", err)
	}
	if _, err := store.Register(repoPath, meta, repo.RegisterOptions{Name: "fixture"}); err != nil {
		t.Fatalf("register repo: %v", err)
	}
	writeDetectChangesSemanticGraph(t, repoPath)

	if err := os.WriteFile(sourcePath, []byte("function main() {\n  return missing() + 1\n}\n"), 0o644); err != nil {
		t.Fatalf("modify source: %v", err)
	}

	server := NewServer(Config{Store: store})
	payload, err := server.detectChangesTool(map[string]any{"repo": "fixture", "scope": "unstaged"})
	if err != nil {
		t.Fatalf("detectChangesTool() error = %v", err)
	}

	summary := payload["summary"].(map[string]any)
	if summary["changed_count"] != 2 || summary["affected_count"] != 1 || summary["risk_level"] != "medium" {
		t.Fatalf("summary = %#v", summary)
	}
	if counts := detectChangesTestCountMap(t, payload["changedAppLayers"]); counts["backend"] != 2 {
		t.Fatalf("changedAppLayers = %#v", counts)
	}
	if counts := detectChangesTestCountMap(t, payload["changedFunctionalAreas"]); counts["resolution"] != 2 {
		t.Fatalf("changedFunctionalAreas = %#v", counts)
	}
	if counts := detectChangesTestCountMap(t, payload["affectedAppLayers"]); counts["backend"] != 1 {
		t.Fatalf("affectedAppLayers = %#v", counts)
	}

	changedSymbols := payload["changed_symbols"].([]map[string]any)
	gap := detectChangesTestFindRow(t, changedSymbols, "ResolutionGap:missing")
	if gap["resolutionGapEntity"] != true || gap["targetText"] != "missing" || gap["actionability"] != graphhealth.DiagnosticActionabilityAnalyzerGap {
		t.Fatalf("gap row = %#v", gap)
	}
	source := detectChangesTestFindRow(t, changedSymbols, "Function:main")
	if source["appLayer"] != "backend" || source["functionalArea"] != "resolution" || source["resolutionConfidence"] != graphhealth.ResolutionConfidenceDegraded {
		t.Fatalf("source row = %#v", source)
	}

	gapChanges := payload["resolutionGapChanges"].(map[string]any)
	if gapChanges["changedGapEntities"] != 1 || gapChanges["changedGapOccurrenceCount"] != 1 || gapChanges["changedSourceNodesWithGaps"] != 1 || gapChanges["totalResolutionGapCount"] != 2 {
		t.Fatalf("resolutionGapChanges = %#v", gapChanges)
	}
	if counts := detectChangesTestCountMap(t, gapChanges["topTargets"]); counts["missing"] != 1 {
		t.Fatalf("gap topTargets = %#v", counts)
	}

	healthImpact := payload["resolutionHealthImpact"].(map[string]any)
	if healthImpact["nodesWithGaps"] != 1 || healthImpact["degradedNodes"] != 1 || healthImpact["totalResolutionGapCount"] != 2 {
		t.Fatalf("resolutionHealthImpact = %#v", healthImpact)
	}
	if counts := detectChangesTestCountMap(t, healthImpact["resolutionHealthBuckets"]); counts[string(graphhealth.ResolutionHealthUnresolvedCallTarget)] != 2 {
		t.Fatalf("resolution health buckets = %#v", counts)
	}

	processes := payload["affected_processes"].([]map[string]any)
	if len(processes) != 1 || processes[0]["appLayer"] != "backend" || processes[0]["functionalArea"] != "resolution" {
		t.Fatalf("affected_processes = %#v", processes)
	}
	if counts := detectChangesTestCountMap(t, processes[0]["changedStepAppLayers"]); counts["backend"] != 1 {
		t.Fatalf("changedStepAppLayers = %#v", counts)
	}
}

func writeDetectChangesSemanticGraph(t *testing.T, repoPath string) {
	t.Helper()

	g := graph.New()
	g.AddNode(graph.Node{ID: "File:src/app.ts", Label: scopeir.NodeFile, Properties: graph.NodeProperties{
		"name": "app.ts", "filePath": "src/app.ts", "appLayer": "backend", "appLayerSource": "test_fixture", "functionalArea": "resolution", "functionalAreaSource": "test_fixture",
	}})
	g.AddNode(graph.Node{ID: "Function:main", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name":                    "main",
		"filePath":                "src/app.ts",
		"startLine":               1,
		"endLine":                 3,
		"appLayer":                "backend",
		"appLayerSource":          "test_fixture",
		"functionalArea":          "resolution",
		"functionalAreaSource":    "test_fixture",
		"topologyStatus":          string(graphhealth.TopologyConnected),
		"resolutionConfidence":    graphhealth.ResolutionConfidenceDegraded,
		"resolutionGapCount":      2,
		"resolutionHealthBuckets": map[string]int{string(graphhealth.ResolutionHealthUnresolvedCallTarget): 2, string(graphhealth.ResolutionHealthInRepoAnalyzerGap): 2},
	}})
	g.AddNode(graph.Node{ID: "ResolutionGap:missing", Label: scopeir.NodeResolutionGap, Properties: graph.NodeProperties{
		"name":                 "missing",
		"filePath":             "src/app.ts",
		"startLine":            2,
		"startCol":             10,
		"endLine":              2,
		"endCol":               18,
		"appLayer":             "backend",
		"appLayerSource":       "test_fixture",
		"functionalArea":       "resolution",
		"functionalAreaSource": "test_fixture",
		"sourceAppLayer":       "backend",
		"sourceFunctionalArea": "resolution",
		"gapKind":              string(graphhealth.ResolutionHealthUnresolvedCallTarget),
		"factFamily":           "call",
		"targetText":           "missing",
		"targetRole":           "callable",
		"sourceSiteStatus":     "unresolved_local_binding",
		"proofKind":            "none",
		"classification":       graphhealth.DiagnosticClassificationInRepoUnresolved,
		"actionability":        graphhealth.DiagnosticActionabilityAnalyzerGap,
		"count":                1,
	}})
	g.AddNode(graph.Node{ID: "Process:main", Label: scopeir.NodeProcess, Properties: graph.NodeProperties{
		"name": "MainFlow", "label": "MainFlow", "heuristicLabel": "MainFlow", "processType": "cross_community", "stepCount": 1,
		"appLayer": "backend", "appLayerSource": "test_fixture", "functionalArea": "resolution", "functionalAreaSource": "test_fixture",
	}})
	step := 1
	g.AddRelationship(graph.Relationship{ID: "rel:main-process", SourceID: "Function:main", TargetID: "Process:main", Type: graph.RelStepInProcess, Step: &step})
	g.AddRelationship(graph.Relationship{
		ID:               "rel:main-gap",
		SourceID:         "Function:main",
		TargetID:         "ResolutionGap:missing",
		Type:             graph.RelHasResolutionGap,
		SourceSiteID:     "SourceSite:src/app.ts#call#missing#2#10#2#18",
		SourceSiteStatus: "unresolved_local_binding",
		ProofKind:        "none",
		TargetRole:       "callable",
		TargetText:       "missing",
		SourceSiteCount:  1,
	})

	raw, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		t.Fatalf("marshal graph: %v", err)
	}
	graphPath := filepath.Join(repo.StoragePath(repoPath), "graph.json")
	if err := os.MkdirAll(filepath.Dir(graphPath), 0o755); err != nil {
		t.Fatalf("mkdir graph: %v", err)
	}
	if err := os.WriteFile(graphPath, append(raw, '\n'), 0o644); err != nil {
		t.Fatalf("write graph: %v", err)
	}
}

func detectChangesTestFindRow(t *testing.T, rows []map[string]any, id string) map[string]any {
	t.Helper()
	for _, row := range rows {
		if row["id"] == id {
			return row
		}
	}
	t.Fatalf("row %q not found in %#v", id, rows)
	return nil
}

func detectChangesTestCountMap(t *testing.T, value any) map[string]int {
	t.Helper()
	out := impactCountMapValue(value)
	if out == nil {
		t.Fatalf("count map is nil for %#v", value)
	}
	return out
}
