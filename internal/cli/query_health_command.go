package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/tamnguyendinh/avmatrix-go/internal/repo"
)

const defaultQueryHealthSuitePath = "docs/query-health/2026-05-22-avmatrix-app-layer-resolution-gap-suite.json"

type queryHealthSuite struct {
	SchemaVersion string            `json:"schemaVersion"`
	Suite         string            `json:"suite"`
	Description   string            `json:"description,omitempty"`
	Cases         []queryHealthCase `json:"cases"`
}

type queryHealthCase struct {
	ID                      string   `json:"id"`
	Intent                  string   `json:"intent"`
	ExpectedFiles           []string `json:"expectedFiles,omitempty"`
	ExpectedSymbols         []string `json:"expectedSymbols,omitempty"`
	ExpectedAppLayers       []string `json:"expectedAppLayers,omitempty"`
	ExpectedFunctionalAreas []string `json:"expectedFunctionalAreas,omitempty"`
	HitAt5Threshold         int      `json:"hitAt5Threshold"`
	HitAt10Threshold        int      `json:"hitAt10Threshold"`
}

type queryHealthReport struct {
	GeneratedAt string                   `json:"generatedAt"`
	Inputs      queryHealthInputs        `json:"inputs"`
	Suite       queryHealthSuiteSummary  `json:"suite"`
	Summary     queryHealthReportSummary `json:"summary"`
	Cases       []queryHealthCaseResult  `json:"cases"`
}

type queryHealthInputs struct {
	SuitePath string `json:"suitePath"`
	Repo      string `json:"repo,omitempty"`
	Limit     int    `json:"limit"`
}

type queryHealthSuiteSummary struct {
	SchemaVersion string `json:"schemaVersion"`
	Suite         string `json:"suite"`
	Description   string `json:"description,omitempty"`
}

type queryHealthReportSummary struct {
	CaseCount int `json:"caseCount"`
	Passed    int `json:"passed"`
	Failed    int `json:"failed"`
}

type queryHealthCaseResult struct {
	ID                      string                    `json:"id"`
	Intent                  string                    `json:"intent"`
	ExpectedFiles           []string                  `json:"expectedFiles,omitempty"`
	ExpectedSymbols         []string                  `json:"expectedSymbols,omitempty"`
	ExpectedAppLayers       []string                  `json:"expectedAppLayers,omitempty"`
	ExpectedFunctionalAreas []string                  `json:"expectedFunctionalAreas,omitempty"`
	ExpectedTargetCount     int                       `json:"expectedTargetCount"`
	HitAt5                  int                       `json:"hitAt5"`
	HitAt10                 int                       `json:"hitAt10"`
	HitAt5Threshold         int                       `json:"hitAt5Threshold"`
	HitAt10Threshold        int                       `json:"hitAt10Threshold"`
	Passed                  bool                      `json:"passed"`
	MatchedTargets          []queryHealthTargetMatch  `json:"matchedTargets,omitempty"`
	MissedTargets           []queryHealthTargetMiss   `json:"missedTargets,omitempty"`
	TopResults              []queryHealthActualResult `json:"topResults,omitempty"`
	NoiseReason             string                    `json:"noiseReason"`
}

type queryHealthTargetMatch struct {
	Kind     string `json:"kind"`
	Expected string `json:"expected"`
	Rank     int    `json:"rank"`
	ResultID string `json:"resultId,omitempty"`
	Name     string `json:"name,omitempty"`
	FilePath string `json:"filePath,omitempty"`
	Source   string `json:"source,omitempty"`
}

type queryHealthTargetMiss struct {
	Kind     string `json:"kind"`
	Expected string `json:"expected"`
	Reason   string `json:"reason"`
}

type queryHealthActualResult struct {
	Rank                    int            `json:"rank"`
	Source                  string         `json:"source"`
	ID                      string         `json:"id,omitempty"`
	Name                    string         `json:"name,omitempty"`
	Type                    string         `json:"type,omitempty"`
	FilePath                string         `json:"filePath,omitempty"`
	Process                 string         `json:"process,omitempty"`
	AppLayer                string         `json:"appLayer,omitempty"`
	FunctionalArea          string         `json:"functionalArea,omitempty"`
	ResolutionConfidence    string         `json:"resolutionConfidence,omitempty"`
	ResolutionGapCount      int            `json:"resolutionGapCount,omitempty"`
	ResolutionHealthBuckets map[string]int `json:"resolutionHealthBuckets,omitempty"`
}

type queryHealthQueryPayload struct {
	Query          string           `json:"query"`
	Processes      []map[string]any `json:"processes"`
	ProcessSymbols []map[string]any `json:"process_symbols"`
	Definitions    []map[string]any `json:"definitions"`
}

type queryHealthRunOptions struct {
	SuitePath string
	Repo      string
	Limit     int
}

type queryHealthRunner func(intent string, repoName string, limit int) (queryHealthQueryPayload, error)

func newQueryHealthCommand() *cobra.Command {
	var suitePath string
	var repoName string
	var outPath string
	var limit int
	var jsonOutput bool
	var failOnThreshold bool

	cmd := &cobra.Command{
		Use:   "query-health",
		Short: "Run query retrieval health benchmarks against the indexed graph",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			report, err := runQueryHealth(queryHealthRunOptions{
				SuitePath: suitePath,
				Repo:      repoName,
				Limit:     limit,
			}, runLocalQueryHealthQuery)
			if err != nil {
				return err
			}
			if outPath != "" {
				if err := writeQueryHealthReport(outPath, report); err != nil {
					return err
				}
			}
			if jsonOutput || outPath == "" {
				raw, err := json.MarshalIndent(report, "", "  ")
				if err != nil {
					return fmt.Errorf("marshal query-health report: %w", err)
				}
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", raw); err != nil {
					return err
				}
			} else {
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "wrote %s\n", outPath); err != nil {
					return err
				}
				for _, line := range queryHealthSummaryLines(report) {
					if _, err := fmt.Fprintln(cmd.OutOrStdout(), line); err != nil {
						return err
					}
				}
			}
			if failOnThreshold && report.Summary.Failed > 0 {
				return fmt.Errorf("query-health thresholds failed for %d/%d cases", report.Summary.Failed, report.Summary.CaseCount)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&suitePath, "suite", defaultQueryHealthSuitePath, "query-health suite JSON")
	cmd.Flags().StringVarP(&repoName, "repo", "r", "", "target repository")
	cmd.Flags().StringVar(&outPath, "out", "", "write query-health JSON report to this path")
	cmd.Flags().IntVarP(&limit, "limit", "l", 10, "max query processes to request from the query tool")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write full JSON report to stdout")
	cmd.Flags().BoolVar(&failOnThreshold, "fail-on-threshold", false, "return an error when any case misses its thresholds")
	return cmd
}

func runQueryHealth(options queryHealthRunOptions, runner queryHealthRunner) (queryHealthReport, error) {
	if strings.TrimSpace(options.SuitePath) == "" {
		return queryHealthReport{}, fmt.Errorf("suite path is required")
	}
	limit := options.Limit
	if limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}
	repoName, err := verifyQueryHealthFreshRepo(options.Repo)
	if err != nil {
		return queryHealthReport{}, err
	}
	suite, err := readQueryHealthSuite(options.SuitePath)
	if err != nil {
		return queryHealthReport{}, err
	}
	results := make([]queryHealthCaseResult, 0, len(suite.Cases))
	summary := queryHealthReportSummary{CaseCount: len(suite.Cases)}
	for _, testCase := range suite.Cases {
		payload, err := runner(testCase.Intent, repoName, limit)
		if err != nil {
			return queryHealthReport{}, fmt.Errorf("run query %q: %w", testCase.ID, err)
		}
		result := scoreQueryHealthCase(testCase, queryHealthActualResults(payload, limit), limit)
		results = append(results, result)
		if result.Passed {
			summary.Passed++
		} else {
			summary.Failed++
		}
	}
	return queryHealthReport{
		GeneratedAt: time.Now().Format(time.RFC3339),
		Inputs: queryHealthInputs{
			SuitePath: options.SuitePath,
			Repo:      repoName,
			Limit:     limit,
		},
		Suite: queryHealthSuiteSummary{
			SchemaVersion: suite.SchemaVersion,
			Suite:         suite.Suite,
			Description:   suite.Description,
		},
		Summary: summary,
		Cases:   results,
	}, nil
}

func readQueryHealthSuite(path string) (queryHealthSuite, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return queryHealthSuite{}, fmt.Errorf("read query-health suite %s: %w", path, err)
	}
	var suite queryHealthSuite
	if err := json.Unmarshal(raw, &suite); err != nil {
		return queryHealthSuite{}, fmt.Errorf("decode query-health suite %s: %w", path, err)
	}
	if strings.TrimSpace(suite.Suite) == "" {
		return queryHealthSuite{}, fmt.Errorf("query-health suite %s is missing suite name", path)
	}
	if len(suite.Cases) == 0 {
		return queryHealthSuite{}, fmt.Errorf("query-health suite %s has no cases", path)
	}
	for index, testCase := range suite.Cases {
		if strings.TrimSpace(testCase.ID) == "" {
			return queryHealthSuite{}, fmt.Errorf("query-health case %d is missing id", index)
		}
		if strings.TrimSpace(testCase.Intent) == "" {
			return queryHealthSuite{}, fmt.Errorf("query-health case %q is missing intent", testCase.ID)
		}
		if len(testCase.ExpectedFiles)+len(testCase.ExpectedSymbols) == 0 {
			return queryHealthSuite{}, fmt.Errorf("query-health case %q has no expected files or symbols", testCase.ID)
		}
	}
	return suite, nil
}

func verifyQueryHealthFreshRepo(repoName string) (string, error) {
	store := repo.NewEnvStore()
	entries, err := store.ListRegistered(false)
	if err != nil {
		return "", err
	}
	var entry repo.RegistryEntry
	if strings.TrimSpace(repoName) == "" {
		if len(entries) != 1 {
			return "", fmt.Errorf("Repository not found. Run: avmatrix analyze --force")
		}
		entry = entries[0]
	} else {
		resolved, err := repo.ResolveEntry(entries, repoName)
		if err != nil {
			return "", err
		}
		entry = resolved
	}
	currentCommit := repo.CurrentCommit(entry.Path)
	if currentCommit != "" && entry.LastCommit != "" && currentCommit != entry.LastCommit {
		return "", fmt.Errorf("query-health requires fresh analyze output for %s: indexed commit %s current commit %s; run avmatrix analyze --force", entry.Name, shortCommit(entry.LastCommit), shortCommit(currentCommit))
	}
	return entry.Name, nil
}

func runLocalQueryHealthQuery(intent string, repoName string, limit int) (queryHealthQueryPayload, error) {
	text, err := callLocalMCPTool("query", map[string]any{
		"query": intent,
		"repo":  emptyToNil(repoName),
		"limit": limit,
	})
	if err != nil {
		return queryHealthQueryPayload{}, err
	}
	var payload queryHealthQueryPayload
	if err := json.Unmarshal([]byte(queryHealthJSONText(text)), &payload); err != nil {
		return queryHealthQueryPayload{}, fmt.Errorf("decode query output: %w", err)
	}
	return payload, nil
}

func queryHealthJSONText(text string) string {
	if before, _, ok := strings.Cut(text, "\n\n---\n"); ok {
		return strings.TrimSpace(before)
	}
	return strings.TrimSpace(text)
}

func queryHealthActualResults(payload queryHealthQueryPayload, limit int) []queryHealthActualResult {
	processRank := make(map[string]int)
	for index, process := range payload.Processes {
		rank := index + 1
		for _, key := range []string{"Label", "label", "ID", "id"} {
			if value := mapString(process, key); value != "" {
				processRank[value] = rank
			}
		}
	}
	results := make([]queryHealthActualResult, 0, len(payload.ProcessSymbols)+len(payload.Definitions))
	for _, symbol := range payload.ProcessSymbols {
		process := firstNonEmptyMapString(symbol, "process", "Process")
		rank := processRank[process]
		if rank == 0 {
			rank = limit + 1
		}
		results = append(results, queryHealthActualResult{
			Rank:                    rank,
			Source:                  "process_symbol",
			ID:                      firstNonEmptyMapString(symbol, "id", "ID"),
			Name:                    firstNonEmptyMapString(symbol, "name", "Name"),
			Type:                    firstNonEmptyMapString(symbol, "type", "Type"),
			FilePath:                firstNonEmptyMapString(symbol, "filePath", "FilePath", "path"),
			Process:                 process,
			AppLayer:                firstNonEmptyMapString(symbol, "appLayer", "AppLayer"),
			FunctionalArea:          firstNonEmptyMapString(symbol, "functionalArea", "FunctionalArea"),
			ResolutionConfidence:    firstNonEmptyMapString(symbol, "resolutionConfidence", "ResolutionConfidence"),
			ResolutionGapCount:      mapInt(symbol, "resolutionGapCount", "ResolutionGapCount"),
			ResolutionHealthBuckets: mapStringInt(symbol, "resolutionHealthBuckets", "ResolutionHealthBuckets"),
		})
	}
	for index, definition := range payload.Definitions {
		results = append(results, queryHealthActualResult{
			Rank:                    index + 1,
			Source:                  "definition",
			ID:                      firstNonEmptyMapString(definition, "id", "ID"),
			Name:                    firstNonEmptyMapString(definition, "name", "Name"),
			Type:                    firstNonEmptyMapString(definition, "type", "Type"),
			FilePath:                firstNonEmptyMapString(definition, "filePath", "FilePath", "path"),
			AppLayer:                firstNonEmptyMapString(definition, "appLayer", "AppLayer"),
			FunctionalArea:          firstNonEmptyMapString(definition, "functionalArea", "FunctionalArea"),
			ResolutionConfidence:    firstNonEmptyMapString(definition, "resolutionConfidence", "ResolutionConfidence"),
			ResolutionGapCount:      mapInt(definition, "resolutionGapCount", "ResolutionGapCount"),
			ResolutionHealthBuckets: mapStringInt(definition, "resolutionHealthBuckets", "ResolutionHealthBuckets"),
		})
	}
	sort.SliceStable(results, func(i, j int) bool {
		if results[i].Rank != results[j].Rank {
			return results[i].Rank < results[j].Rank
		}
		if results[i].Source != results[j].Source {
			return results[i].Source < results[j].Source
		}
		return results[i].ID < results[j].ID
	})
	return results
}

func scoreQueryHealthCase(testCase queryHealthCase, actual []queryHealthActualResult, outputLimit int) queryHealthCaseResult {
	result := queryHealthCaseResult{
		ID:                      testCase.ID,
		Intent:                  testCase.Intent,
		ExpectedFiles:           append([]string{}, testCase.ExpectedFiles...),
		ExpectedSymbols:         append([]string{}, testCase.ExpectedSymbols...),
		ExpectedAppLayers:       append([]string{}, testCase.ExpectedAppLayers...),
		ExpectedFunctionalAreas: append([]string{}, testCase.ExpectedFunctionalAreas...),
		ExpectedTargetCount:     len(testCase.ExpectedFiles) + len(testCase.ExpectedSymbols),
		HitAt5Threshold:         testCase.HitAt5Threshold,
		HitAt10Threshold:        testCase.HitAt10Threshold,
		TopResults:              topQueryHealthResults(actual, outputLimit),
	}
	for _, expected := range testCase.ExpectedFiles {
		match, ok := findQueryHealthMatch("file", expected, actual)
		if ok {
			result.MatchedTargets = append(result.MatchedTargets, match)
			if match.Rank <= 5 {
				result.HitAt5++
			}
			if match.Rank <= 10 {
				result.HitAt10++
			}
			continue
		}
		result.MissedTargets = append(result.MissedTargets, queryHealthTargetMiss{Kind: "file", Expected: expected, Reason: "expected file was not returned by query top results"})
	}
	for _, expected := range testCase.ExpectedSymbols {
		match, ok := findQueryHealthMatch("symbol", expected, actual)
		if ok {
			result.MatchedTargets = append(result.MatchedTargets, match)
			if match.Rank <= 5 {
				result.HitAt5++
			}
			if match.Rank <= 10 {
				result.HitAt10++
			}
			continue
		}
		result.MissedTargets = append(result.MissedTargets, queryHealthTargetMiss{Kind: "symbol", Expected: expected, Reason: "expected function/method target was not returned by current definition or process-symbol matching"})
	}
	result.Passed = result.HitAt5 >= testCase.HitAt5Threshold && result.HitAt10 >= testCase.HitAt10Threshold
	result.NoiseReason = queryHealthNoiseReason(result, actual)
	return result
}

func findQueryHealthMatch(kind string, expected string, actual []queryHealthActualResult) (queryHealthTargetMatch, bool) {
	for _, item := range actual {
		matched := false
		switch kind {
		case "file":
			matched = queryHealthPathMatches(item.FilePath, expected)
		case "symbol":
			matched = queryHealthSymbolMatches(item, expected)
		}
		if !matched {
			continue
		}
		return queryHealthTargetMatch{
			Kind:     kind,
			Expected: expected,
			Rank:     item.Rank,
			ResultID: item.ID,
			Name:     item.Name,
			FilePath: item.FilePath,
			Source:   item.Source,
		}, true
	}
	return queryHealthTargetMatch{}, false
}

func queryHealthPathMatches(actual string, expected string) bool {
	actual = normalizeQueryHealthPath(actual)
	expected = normalizeQueryHealthPath(expected)
	if actual == "" || expected == "" {
		return false
	}
	return actual == expected || strings.HasSuffix(actual, "/"+expected)
}

func normalizeQueryHealthPath(path string) string {
	path = strings.ReplaceAll(strings.TrimSpace(path), "\\", "/")
	path = strings.Trim(path, "/")
	return strings.ToLower(path)
}

func queryHealthSymbolMatches(item queryHealthActualResult, expected string) bool {
	expected = strings.ToLower(strings.TrimSpace(expected))
	if expected == "" {
		return false
	}
	name := strings.ToLower(strings.TrimSpace(item.Name))
	id := strings.ToLower(strings.TrimSpace(item.ID))
	return name == expected || strings.Contains(id, ":"+expected) || strings.Contains(id, "#"+expected)
}

func topQueryHealthResults(actual []queryHealthActualResult, limit int) []queryHealthActualResult {
	if limit <= 0 {
		limit = 10
	}
	if len(actual) < limit {
		limit = len(actual)
	}
	return append([]queryHealthActualResult{}, actual[:limit]...)
}

func queryHealthNoiseReason(result queryHealthCaseResult, actual []queryHealthActualResult) string {
	if result.Passed {
		return "thresholds met"
	}
	missing := make([]string, 0, len(result.MissedTargets))
	for _, miss := range result.MissedTargets {
		missing = append(missing, miss.Kind+":"+miss.Expected)
	}
	topFiles := uniqueTopQueryHealthFiles(actual, 5)
	if len(topFiles) == 0 {
		return "no query results; missing " + strings.Join(missing, ", ")
	}
	return "missing " + strings.Join(missing, ", ") + "; top files: " + strings.Join(topFiles, ", ")
}

func uniqueTopQueryHealthFiles(actual []queryHealthActualResult, maxFiles int) []string {
	files := make([]string, 0, maxFiles)
	seen := map[string]bool{}
	for _, item := range actual {
		filePath := strings.TrimSpace(item.FilePath)
		if filePath == "" || seen[filePath] {
			continue
		}
		seen[filePath] = true
		files = append(files, filePath)
		if len(files) >= maxFiles {
			break
		}
	}
	return files
}

func writeQueryHealthReport(path string, report queryHealthReport) error {
	raw, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal query-health report: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

func queryHealthSummaryLines(report queryHealthReport) []string {
	lines := []string{
		fmt.Sprintf("queryHealth.suite=%s cases=%d passed=%d failed=%d", report.Suite.Suite, report.Summary.CaseCount, report.Summary.Passed, report.Summary.Failed),
	}
	for _, result := range report.Cases {
		status := "FAIL"
		if result.Passed {
			status = "PASS"
		}
		lines = append(lines, fmt.Sprintf(
			"queryHealth.case=%s status=%s hitAt5=%d/%d hitAt10=%d/%d expected=%d noise=%s",
			result.ID,
			status,
			result.HitAt5,
			result.HitAt5Threshold,
			result.HitAt10,
			result.HitAt10Threshold,
			result.ExpectedTargetCount,
			result.NoiseReason,
		))
	}
	return lines
}

func firstNonEmptyMapString(values map[string]any, keys ...string) string {
	for _, key := range keys {
		if value := mapString(values, key); value != "" {
			return value
		}
	}
	return ""
}

func mapString(values map[string]any, key string) string {
	raw, ok := values[key]
	if !ok || raw == nil {
		return ""
	}
	switch value := raw.(type) {
	case string:
		return value
	default:
		return strings.TrimSpace(fmt.Sprint(value))
	}
}

func mapInt(values map[string]any, keys ...string) int {
	for _, key := range keys {
		raw, ok := values[key]
		if !ok || raw == nil {
			continue
		}
		switch value := raw.(type) {
		case int:
			return value
		case int32:
			return int(value)
		case int64:
			return int(value)
		case float64:
			return int(value)
		}
	}
	return 0
}

func mapStringInt(values map[string]any, keys ...string) map[string]int {
	for _, key := range keys {
		raw, ok := values[key]
		if !ok || raw == nil {
			continue
		}
		out := map[string]int{}
		switch typed := raw.(type) {
		case map[string]int:
			for itemKey, count := range typed {
				out[itemKey] = count
			}
		case map[string]any:
			for itemKey, count := range typed {
				out[itemKey] = anyToInt(count)
			}
		}
		if len(out) > 0 {
			return out
		}
	}
	return nil
}

func anyToInt(value any) int {
	switch typed := value.(type) {
	case int:
		return typed
	case int32:
		return int(typed)
	case int64:
		return int(typed)
	case float64:
		return int(typed)
	default:
		return 0
	}
}
