package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/repo"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestGraphHealthCommandSummaryReportComponentsAndExplain(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeDirectToolCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "graph-health", "summary", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("graph-health summary returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("graph-health summary wrote stderr: %q", errOut)
	}
	var summary graphHealthSummaryResult
	if err := json.Unmarshal([]byte(out), &summary); err != nil {
		t.Fatalf("summary output is not JSON: %v\n%s", err, out)
	}
	if summary.Summary.NodeCount != 5 || summary.Summary.CountedRelationshipCount != 3 || summary.Summary.ComponentCount != 3 {
		t.Fatalf("unexpected summary counts: %#v", summary.Summary)
	}

	out, errOut, err = executeForTest(t, "graph-health", "summary", "--repo", "fixture")
	if err != nil {
		t.Fatalf("graph-health summary table returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if !strings.Contains(out, "graphHealth.repo=") || !strings.Contains(out, "countedRelationships=3") {
		t.Fatalf("summary table missing expected fields:\n%s", out)
	}

	out, errOut, err = executeForTest(t, "graph-health", "report", "--repo", "fixture", "--limit", "1", "--json")
	if err != nil {
		t.Fatalf("graph-health report returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	var report graphHealthReportResult
	if err := json.Unmarshal([]byte(out), &report); err != nil {
		t.Fatalf("report output is not JSON: %v\n%s", err, out)
	}
	if report.ReturnedCandidates != 1 || len(report.Candidates) != 1 {
		t.Fatalf("unexpected report limit result: %#v", report.ReportResponse)
	}
	first := report.Candidates[0]
	if first.NodeID != "Function:main" || first.TriagePriority != string(graphhealth.TopologyNoIncoming) || first.TriageDimension != graphhealth.TriageDimensionTopology {
		t.Fatalf("unexpected first report candidate: %#v", first)
	}

	out, errOut, err = executeForTest(t, "graph-health", "report", "--repo", "fixture", "--limit", "1")
	if err != nil {
		t.Fatalf("graph-health report table returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if !strings.Contains(out, `graphHealth.candidate nodeId="Function:main"`) {
		t.Fatalf("report table missing expected candidate:\n%s", out)
	}

	out, errOut, err = executeForTest(t, "graph-health", "components", "--repo", "fixture", "--limit", "2", "--json")
	if err != nil {
		t.Fatalf("graph-health components returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	var components graphHealthComponentsResult
	if err := json.Unmarshal([]byte(out), &components); err != nil {
		t.Fatalf("components output is not JSON: %v\n%s", err, out)
	}
	if components.TotalComponents != 3 || components.ReturnedComponents != 2 || len(components.Components) != 2 {
		t.Fatalf("unexpected component counts: %#v", components)
	}
	if components.Components[0].NodeCount != 3 || components.Components[0].CountedEdgeCount != 3 {
		t.Fatalf("unexpected first component summary: %#v", components.Components[0])
	}

	out, errOut, err = executeForTest(t, "graph-health", "explain", "Function:main", "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("graph-health explain node returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	var nodeExplain graphHealthExplainResult
	if err := json.Unmarshal([]byte(out), &nodeExplain); err != nil {
		t.Fatalf("node explain output is not JSON: %v\n%s", err, out)
	}
	if nodeExplain.Kind != "node" || nodeExplain.NodeID != "Function:main" || nodeExplain.Health == nil {
		t.Fatalf("unexpected node explain payload: %#v", nodeExplain.ExplainResponse)
	}
	if nodeExplain.Health.TopologyStatus != graphhealth.TopologyNoIncoming || len(nodeExplain.CountedOutgoingRelationships) != 2 {
		t.Fatalf("unexpected node explain health: %#v", nodeExplain.ExplainResponse)
	}

	out, errOut, err = executeForTest(t, "graph-health", "explain", "helper", "--repo", "fixture")
	if err != nil {
		t.Fatalf("graph-health explain by name returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if !strings.Contains(out, `graphHealth.node nodeId="Function:helper"`) {
		t.Fatalf("explain by name table missing resolved node:\n%s", out)
	}

	componentID := components.Components[0].ID
	out, errOut, err = executeForTest(t, "graph-health", "explain", "--component", componentID, "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("graph-health explain component returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	var componentExplain graphHealthExplainResult
	if err := json.Unmarshal([]byte(out), &componentExplain); err != nil {
		t.Fatalf("component explain output is not JSON: %v\n%s", err, out)
	}
	if componentExplain.Kind != "component" || componentExplain.Component == nil || componentExplain.Component.NodeCount != 3 {
		t.Fatalf("unexpected component explain payload: %#v", componentExplain.ExplainResponse)
	}

	out, errOut, err = executeForTest(t, "graph-health", "explain", "missing", "--repo", "fixture")
	if err == nil {
		t.Fatalf("missing node unexpectedly passed\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), "graph node not found") {
		t.Fatalf("missing node error = %v", err)
	}
}

func TestGraphHealthCommandReportsAmbiguousNodeName(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")

	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:first", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "duplicate"}})
	g.AddNode(graph.Node{ID: "Function:second", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "duplicate"}})
	writeGroupCommandGraph(t, repoPath, g)

	out, errOut, err := executeForTest(t, "graph-health", "explain", "duplicate", "--repo", "fixture")
	if err == nil {
		t.Fatalf("ambiguous node unexpectedly passed\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), "ambiguous") || !strings.Contains(err.Error(), "Function:first") || !strings.Contains(err.Error(), "Function:second") {
		t.Fatalf("ambiguous node error = %v", err)
	}
}

func TestGraphHealthCommandRequiresFreshGraph(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")

	out, errOut, err := executeForTest(t, "graph-health", "summary", "--repo", "fixture")
	if err == nil {
		t.Fatalf("missing graph unexpectedly passed\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), "graph-health requires fresh analyze output") {
		t.Fatalf("missing graph error = %v", err)
	}
}

func TestGraphHealthCommandRejectsStaleIndexedCommit(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := initGitRepo(t)
	store := repo.NewStore(home)
	meta := repo.Meta{RepoPath: repoPath, LastCommit: "stale-commit", IndexedAt: "2026-05-26T00:00:00Z", Stats: &repo.Stats{}}
	if err := repo.SaveMeta(repo.StoragePath(repoPath), meta); err != nil {
		t.Fatalf("save meta: %v", err)
	}
	if _, err := store.Register(repoPath, meta, repo.RegisterOptions{Name: "fixture"}); err != nil {
		t.Fatalf("register repo: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(filepath.Join(repo.StoragePath(repoPath), "graph.json")), 0o755); err != nil {
		t.Fatalf("mkdir graph: %v", err)
	}
	writeDirectToolCommandGraph(t, repoPath)

	out, errOut, err := executeForTest(t, "graph-health", "summary", "--repo", "fixture")
	if err == nil {
		t.Fatalf("stale graph unexpectedly passed\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), "graph-health requires fresh analyze output") || !strings.Contains(err.Error(), "run anvien analyze --force") {
		t.Fatalf("stale graph error = %v", err)
	}
}
