package graphaccuracy

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/tamnguyendinh/anvien/internal/analyze"
	"github.com/tamnguyendinh/anvien/internal/resolution"
)

type AccessCandidateAuditOptions struct {
	Repo        string
	OutPath     string
	MaxExamples int
}

type AccessCandidateAuditResult struct {
	GeneratedAt string                          `json:"generatedAt"`
	Inputs      AccessCandidateAuditInputs      `json:"inputs"`
	Analyze     AccessCandidateAnalyzeSummary   `json:"analyze"`
	Audit       resolution.AccessCandidateAudit `json:"audit"`
	Notes       []string                        `json:"notes"`
}

type AccessCandidateAuditInputs struct {
	Repo string `json:"repo"`
}

type AccessCandidateAnalyzeSummary struct {
	FilesScanned int `json:"filesScanned"`
	// FilesParsed and FilesUnsupported are legacy aliases for parsedCode and unsupportedLanguage.
	FilesParsed              int   `json:"filesParsed"`
	FilesParsedCode          int   `json:"filesParsedCode"`
	FilesDocuments           int   `json:"filesDocuments"`
	FilesMetadataOnly        int   `json:"filesMetadataOnly"`
	FilesDedicatedAnalyzer   int   `json:"filesDedicatedAnalyzer"`
	FilesScriptNoExtractor   int   `json:"filesScriptNoExtractor"`
	FilesStaticAssets        int   `json:"filesStaticAssets"`
	FilesUnsupported         int   `json:"filesUnsupported"`
	FilesUnsupportedLanguage int   `json:"filesUnsupportedLanguage"`
	FilesUnknown             int   `json:"filesUnknown"`
	FilesFailed              int   `json:"filesFailed"`
	TotalDurationMillis      int64 `json:"totalDurationMillis"`
	ResolvedAccesses         int   `json:"resolvedAccesses"`
	UnresolvedReferences     int   `json:"unresolvedReferences"`
	GraphNodes               int   `json:"graphNodes"`
	GraphRelationships       int   `json:"graphRelationships"`
	ScopeIRsRetained         int   `json:"scopeIRsRetained"`
}

func RunAccessCandidateAudit(ctx context.Context, options AccessCandidateAuditOptions) (AccessCandidateAuditResult, error) {
	if options.MaxExamples <= 0 {
		options.MaxExamples = 50
	}
	repo := options.Repo
	if strings.TrimSpace(repo) == "" {
		repo = "."
	}
	repoAbs, err := filepath.Abs(repo)
	if err != nil {
		return AccessCandidateAuditResult{}, fmt.Errorf("resolve repo: %w", err)
	}
	analyzeResult, err := analyze.Run(ctx, repoAbs, analyze.Options{})
	if err != nil {
		return AccessCandidateAuditResult{}, err
	}
	audit, err := resolution.AuditAccessCandidates(analyzeResult.ScopeIRs, resolution.AccessCandidateAuditOptions{
		MaxExamples: options.MaxExamples,
	})
	if err != nil {
		return AccessCandidateAuditResult{}, err
	}
	result := AccessCandidateAuditResult{
		GeneratedAt: time.Now().Format(time.RFC3339),
		Inputs: AccessCandidateAuditInputs{
			Repo: repoAbs,
		},
		Analyze: AccessCandidateAnalyzeSummary{
			FilesScanned:             analyzeResult.Metrics.Files.Scanned,
			FilesParsed:              analyzeResult.Metrics.Files.Parsed,
			FilesParsedCode:          analyzeResult.Metrics.Files.ParsedCode,
			FilesDocuments:           analyzeResult.Metrics.Files.Documents,
			FilesMetadataOnly:        analyzeResult.Metrics.Files.MetadataOnly,
			FilesDedicatedAnalyzer:   analyzeResult.Metrics.Files.DedicatedAnalyzer,
			FilesScriptNoExtractor:   analyzeResult.Metrics.Files.ScriptNoExtractor,
			FilesStaticAssets:        analyzeResult.Metrics.Files.StaticAssets,
			FilesUnsupported:         analyzeResult.Metrics.Files.Unsupported,
			FilesUnsupportedLanguage: analyzeResult.Metrics.Files.UnsupportedLanguage,
			FilesUnknown:             analyzeResult.Metrics.Files.Unknown,
			FilesFailed:              analyzeResult.Metrics.Files.Failed,
			TotalDurationMillis:      analyzeResult.Metrics.TotalDuration.Milliseconds(),
			ResolvedAccesses:         analyzeResult.Metrics.Resolution.ResolvedAccesses,
			UnresolvedReferences:     analyzeResult.Metrics.Resolution.UnresolvedReferences,
			GraphNodes:               len(analyzeResult.Graph.Nodes),
			GraphRelationships:       len(analyzeResult.Graph.Relationships),
			ScopeIRsRetained:         len(analyzeResult.ScopeIRs),
		},
		Audit: audit,
		Notes: []string{
			"Access candidates are audited from retained ScopeIR facts, not inferred from final graph relationships.",
			"resolved is the candidate-level classification from the same owner/member lookup model used by resolution.",
			"missing_owner_link means a matching standalone Property exists, but it is not connected to the receiver owner.",
		},
	}
	if options.OutPath != "" {
		if err := WriteAccessCandidateAuditResult(options.OutPath, result); err != nil {
			return AccessCandidateAuditResult{}, err
		}
	}
	return result, nil
}

func WriteAccessCandidateAuditResult(path string, result AccessCandidateAuditResult) error {
	raw, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal access candidate audit: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

func AccessCandidateAuditSummaryLines(result AccessCandidateAuditResult) []string {
	lines := []string{
		fmt.Sprintf("accessCandidates.total=%d resolved=%d unresolved=%d analyzeMillis=%d resolvedAccesses=%d unresolvedReferences=%d",
			result.Audit.Total,
			result.Audit.Resolved,
			result.Audit.Unresolved,
			result.Analyze.TotalDurationMillis,
			result.Analyze.ResolvedAccesses,
			result.Analyze.UnresolvedReferences,
		),
	}
	for _, language := range sortedAccessLanguageKeys(result.Audit.Languages) {
		stats := result.Audit.Languages[language]
		lines = append(lines, fmt.Sprintf("language.%s.accessCandidates=%d resolved=%d unresolved=%d",
			language,
			stats.Total,
			stats.Resolved,
			stats.Unresolved,
		))
	}
	for _, reason := range sortedAccessReasonKeys(result.Audit.Reasons) {
		lines = append(lines, fmt.Sprintf("reason.%s=%d", reason, result.Audit.Reasons[reason].Count))
	}
	return lines
}

func sortedAccessLanguageKeys(languages map[string]resolution.AccessCandidateLanguageStats) []string {
	keys := make([]string, 0, len(languages))
	for key := range languages {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedAccessReasonKeys(reasons map[string]resolution.AccessCandidateAuditBucket) []string {
	keys := make([]string, 0, len(reasons))
	for key := range reasons {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
