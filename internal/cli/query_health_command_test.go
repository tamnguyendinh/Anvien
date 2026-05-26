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
		QueryLanes:           []string{"graph_quality_discovery"},
		MatchReasons:         []string{"filePath", "semanticSurface"},
	}}
	result := scoreQueryHealthCase(suite.Cases[0], actual, 10)
	if !result.Passed || !result.ThresholdPassed || !result.ExactPassed || result.HitAt5 != 2 || result.HitAt10 != 2 {
		t.Fatalf("score result = %#v", result)
	}
	if result.MatchedTargetCount != 2 || result.MissedTargetCount != 0 {
		t.Fatalf("target coverage = matched %d missed %d", result.MatchedTargetCount, result.MissedTargetCount)
	}
	if len(result.TopResults) != 1 || result.TopResults[0].AppLayer != "backend" || result.TopResults[0].FunctionalArea != "graph_health" {
		t.Fatalf("semantic fields not preserved in top results: %#v", result.TopResults)
	}
	if len(result.TopResults[0].QueryLanes) != 1 || result.TopResults[0].QueryLanes[0] != "graph_quality_discovery" ||
		len(result.MatchedTargets[0].MatchReasons) == 0 {
		t.Fatalf("query lane evidence not preserved: %#v", result)
	}
}

func TestQueryHealthScoringSeparatesThresholdAndExactPass(t *testing.T) {
	testCase := queryHealthCase{
		ID:               "partial-match",
		Intent:           "query implementation",
		ExpectedFiles:    []string{"internal/cli/query_health_command.go"},
		ExpectedSymbols:  []string{"missingExactSymbol"},
		HitAt5Threshold:  1,
		HitAt10Threshold: 1,
	}
	actual := []queryHealthActualResult{{
		Rank:     1,
		Source:   "definition",
		ID:       "Function:internal/cli/query_health_command.go:runQueryHealth#2",
		Name:     "runQueryHealth",
		FilePath: "internal/cli/query_health_command.go",
	}}
	result := scoreQueryHealthCase(testCase, actual, 10)
	if !result.Passed || !result.ThresholdPassed {
		t.Fatalf("threshold pass should remain true for usable retrieval: %#v", result)
	}
	if result.ExactPassed {
		t.Fatalf("exact pass should be false when any expected target is missed: %#v", result)
	}
	if result.MatchedTargetCount != 1 || result.MissedTargetCount != 1 {
		t.Fatalf("target coverage = matched %d missed %d", result.MatchedTargetCount, result.MissedTargetCount)
	}
	if !strings.Contains(result.NoiseReason, "thresholds met; exact target misses") ||
		!strings.Contains(result.NoiseReason, "symbol:missingExactSymbol") {
		t.Fatalf("noise reason should distinguish exact miss: %s", result.NoiseReason)
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
	if decoded.Summary.Passed != 1 ||
		decoded.Summary.Failed != 0 ||
		decoded.Summary.ThresholdPassed != 1 ||
		decoded.Summary.ThresholdFailed != 0 ||
		decoded.Summary.ExactPassed != 1 ||
		decoded.Summary.ExactFailed != 0 ||
		decoded.Summary.MatchedTargetCount != 2 ||
		decoded.Summary.ExpectedTargetCount != 2 ||
		decoded.Cases[0].HitAt5 != 2 ||
		!decoded.Cases[0].ThresholdPassed ||
		!decoded.Cases[0].ExactPassed {
		t.Fatalf("decoded report = %#v", decoded)
	}
	if !strings.Contains(out, `"thresholdPassed"`) ||
		!strings.Contains(out, `"exactPassed"`) ||
		!strings.Contains(out, `"matchedTargetCount"`) ||
		!strings.Contains(out, `"topResults"`) ||
		!strings.Contains(out, `"Function:main"`) {
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
	if !strings.Contains(out, "queryHealth.suite=fixture-suite") ||
		!strings.Contains(out, "thresholdPassed=1") ||
		!strings.Contains(out, "exactPassed=1") ||
		!strings.Contains(out, "threshold=PASS exact=PASS") {
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

	partialSuite := filepath.Join(dir, "partial-suite.json")
	writeQueryHealthTestSuite(t, partialSuite, "src/app.ts", "missingExactSymbol", 1, 1)
	out, errOut, err = executeForTest(t, "query-health", "--suite", partialSuite, "--repo", "fixture", "--out", filepath.Join(dir, "partial.json"))
	if err != nil {
		t.Fatalf("query-health partial exact miss should pass threshold: %v\nstdout:\n%s\nstderr:\n%s", err, out, errOut)
	}
	if !strings.Contains(out, "threshold=PASS exact=FAIL") || !strings.Contains(out, "missedTargets=1") {
		t.Fatalf("partial exact miss output did not separate threshold/exact:\n%s", out)
	}
	out, errOut, err = executeForTest(t, "query-health", "--suite", partialSuite, "--repo", "fixture", "--json", "--fail-on-exact")
	if err == nil {
		t.Fatalf("query-health fail-on-exact unexpectedly passed\nstdout:\n%s\nstderr:\n%s", out, errOut)
	}
	if !strings.Contains(err.Error(), "query-health exact target coverage failed") {
		t.Fatalf("unexpected exact coverage error: %v", err)
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
