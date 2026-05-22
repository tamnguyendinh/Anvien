package cli

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/avmatrix-go/internal/repo"
)

func TestQueryHealthSuiteParsingAndScoring(t *testing.T) {
	dir := t.TempDir()
	suitePath := filepath.Join(dir, "suite.json")
	raw := `{
  "schemaVersion": "avmatrix.query_health.v1",
  "suite": "fixture-suite",
  "description": "fixture",
  "cases": [
    {
      "id": "semantic-result",
      "intent": "resolution health",
      "expectedFiles": ["internal/graphhealth/compute.go"],
      "expectedSymbols": ["ComputeSummary"],
      "expectedAppLayers": ["backend"],
      "expectedFunctionalAreas": ["graph_health"],
      "hitAt5Threshold": 2,
      "hitAt10Threshold": 2
    }
  ]
}`
	if err := os.WriteFile(suitePath, []byte(raw), 0o644); err != nil {
		t.Fatalf("write suite: %v", err)
	}
	suite, err := readQueryHealthSuite(suitePath)
	if err != nil {
		t.Fatalf("read suite: %v", err)
	}
	if suite.Suite != "fixture-suite" || len(suite.Cases) != 1 {
		t.Fatalf("suite = %#v", suite)
	}

	actual := []queryHealthActualResult{{
		Rank:                 1,
		Source:               "process_symbol",
		ID:                   "Function:internal/graphhealth/compute.go:ComputeSummary#1",
		Name:                 "ComputeSummary",
		FilePath:             "internal/graphhealth/compute.go",
		AppLayer:             "backend",
		FunctionalArea:       "graph_health",
		ResolutionConfidence: "degraded",
		ResolutionGapCount:   4,
	}}
	result := scoreQueryHealthCase(suite.Cases[0], actual, 10)
	if !result.Passed || result.HitAt5 != 2 || result.HitAt10 != 2 {
		t.Fatalf("score result = %#v", result)
	}
	if len(result.TopResults) != 1 || result.TopResults[0].AppLayer != "backend" || result.TopResults[0].FunctionalArea != "graph_health" {
		t.Fatalf("semantic fields not preserved in top results: %#v", result.TopResults)
	}
}

func TestQueryHealthScoringReportsMissingAndNoisyResults(t *testing.T) {
	testCase := queryHealthCase{
		ID:               "missing-resolution",
		Intent:           "unresolved reference diagnostic generation",
		ExpectedFiles:    []string{"internal/resolution/resolve.go"},
		ExpectedSymbols:  []string{"resolveCall"},
		HitAt5Threshold:  1,
		HitAt10Threshold: 1,
	}
	actual := []queryHealthActualResult{{
		Rank:     1,
		Source:   "process_symbol",
		ID:       "Function:avmatrix-launcher/src/main.go:hiddenProcAttr#1",
		Name:     "hiddenProcAttr",
		FilePath: "avmatrix-launcher/src/main.go",
	}}
	result := scoreQueryHealthCase(testCase, actual, 10)
	if result.Passed || result.HitAt5 != 0 || result.HitAt10 != 0 {
		t.Fatalf("score result = %#v", result)
	}
	if len(result.MissedTargets) != 2 {
		t.Fatalf("misses = %#v", result.MissedTargets)
	}
	if !strings.Contains(result.NoiseReason, "internal/resolution/resolve.go") ||
		!strings.Contains(result.NoiseReason, "avmatrix-launcher/src/main.go") {
		t.Fatalf("noise reason missing details: %s", result.NoiseReason)
	}
}

func TestQueryHealthCommandOutputsJSONTableAndFailsThreshold(t *testing.T) {
	home := t.TempDir()
	t.Setenv(repo.HomeEnvName, home)
	repoPath := t.TempDir()
	registerDirectToolCommandRepo(t, repo.NewStore(home), repoPath, "fixture")
	writeDirectToolCommandGraph(t, repoPath)

	dir := t.TempDir()
	passSuite := filepath.Join(dir, "pass-suite.json")
	writeQueryHealthTestSuite(t, passSuite, "src/app.ts", "main", 2, 2)

	out, errOut, err := executeForTest(t, "query-health", "--suite", passSuite, "--repo", "fixture", "--json")
	if err != nil {
		t.Fatalf("query-health returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("query-health wrote stderr: %q", errOut)
	}
	var decoded queryHealthReport
	if err := json.Unmarshal([]byte(out), &decoded); err != nil {
		t.Fatalf("query-health output is not JSON: %v\n%s", err, out)
	}
	if decoded.Summary.Passed != 1 || decoded.Summary.Failed != 0 || decoded.Cases[0].HitAt5 != 2 {
		t.Fatalf("decoded report = %#v", decoded)
	}
	if !strings.Contains(out, `"topResults"`) || !strings.Contains(out, `"Function:main"`) {
		t.Fatalf("query-health JSON missing expected result details:\n%s", out)
	}

	reportPath := filepath.Join(dir, "query-health-report.json")
	out, errOut, err = executeForTest(t, "query-health", "--suite", passSuite, "--repo", "fixture", "--out", reportPath)
	if err != nil {
		t.Fatalf("query-health table returned error: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if errOut != "" {
		t.Fatalf("query-health table wrote stderr: %q", errOut)
	}
	if !strings.Contains(out, "queryHealth.suite=fixture-suite") || !strings.Contains(out, "status=PASS") {
		t.Fatalf("query-health table missing summary:\n%s", out)
	}
	if _, err := os.Stat(reportPath); err != nil {
		t.Fatalf("query-health report not written: %v", err)
	}

	failSuite := filepath.Join(dir, "fail-suite.json")
	writeQueryHealthTestSuite(t, failSuite, "missing/file.go", "missingSymbol", 1, 1)
	out, errOut, err = executeForTest(t, "query-health", "--suite", failSuite, "--repo", "fixture", "--json", "--fail-on-threshold")
	if err == nil {
		t.Fatalf("query-health fail-on-threshold unexpectedly passed\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), "query-health thresholds failed") {
		t.Fatalf("unexpected threshold error: %v", err)
	}
	if !strings.Contains(out, `"passed": false`) || !strings.Contains(out, `"missedTargets"`) {
		t.Fatalf("threshold failure did not include failure report:\n%s", out)
	}
}

func writeQueryHealthTestSuite(t *testing.T, path string, expectedFile string, expectedSymbol string, hitAt5 int, hitAt10 int) {
	t.Helper()
	suite := queryHealthSuite{
		SchemaVersion: "avmatrix.query_health.v1",
		Suite:         "fixture-suite",
		Cases: []queryHealthCase{{
			ID:               "main-flow",
			Intent:           "MainFlow",
			ExpectedFiles:    []string{expectedFile},
			ExpectedSymbols:  []string{expectedSymbol},
			HitAt5Threshold:  hitAt5,
			HitAt10Threshold: hitAt10,
		}},
	}
	raw, err := json.MarshalIndent(suite, "", "  ")
	if err != nil {
		t.Fatalf("marshal suite: %v", err)
	}
	if err := os.WriteFile(path, append(raw, '\n'), 0o644); err != nil {
		t.Fatalf("write suite: %v", err)
	}
}
