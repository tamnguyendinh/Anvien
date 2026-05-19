package resolution

import (
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
)

type AccessCandidateAuditOptions struct {
	MaxExamples int
}

type AccessCandidateAudit struct {
	Total      int                                     `json:"total"`
	Resolved   int                                     `json:"resolved"`
	Unresolved int                                     `json:"unresolved"`
	Reasons    map[string]AccessCandidateAuditBucket   `json:"reasons"`
	Languages  map[string]AccessCandidateLanguageStats `json:"languages"`
}

type AccessCandidateLanguageStats struct {
	Total      int            `json:"total"`
	Resolved   int            `json:"resolved"`
	Unresolved int            `json:"unresolved"`
	Reasons    map[string]int `json:"reasons,omitempty"`
}

type AccessCandidateAuditBucket struct {
	Count    int                           `json:"count"`
	Examples []AccessCandidateAuditExample `json:"examples,omitempty"`
}

type AccessCandidateAuditExample struct {
	FilePath  string `json:"filePath,omitempty"`
	Language  string `json:"language,omitempty"`
	Name      string `json:"name,omitempty"`
	Receiver  string `json:"receiver,omitempty"`
	Kind      string `json:"kind,omitempty"`
	StartLine int    `json:"startLine,omitempty"`
	Reason    string `json:"reason,omitempty"`
	Detail    string `json:"detail,omitempty"`
}

func AuditAccessCandidates(files []scopeir.ScopeIR, options AccessCandidateAuditOptions) (AccessCandidateAudit, error) {
	if options.MaxExamples <= 0 {
		options.MaxExamples = 50
	}
	w, err := buildWorkspace(files)
	if err != nil {
		return AccessCandidateAudit{}, err
	}
	defer w.bindingAccumulator.dispose()
	if err := w.bindingAccumulator.finalize(); err != nil {
		return AccessCandidateAudit{}, err
	}
	result := AccessCandidateAudit{
		Reasons:   map[string]AccessCandidateAuditBucket{},
		Languages: map[string]AccessCandidateLanguageStats{},
	}
	for _, ir := range w.files {
		for _, access := range ir.Accesses {
			result.Total++
			language := string(w.fileLanguages[cleanPath(access.FilePath)])
			if language == "" {
				language = string(ir.Language)
			}
			reason, resolved, detail := classifyAccessCandidate(w, access)
			if resolved {
				result.Resolved++
			} else {
				result.Unresolved++
			}
			example := AccessCandidateAuditExample{
				FilePath:  cleanPath(access.FilePath),
				Language:  language,
				Name:      access.Name,
				Receiver:  access.ExplicitReceiver,
				Kind:      string(access.Kind),
				StartLine: access.Range.StartLine,
				Reason:    reason,
				Detail:    detail,
			}
			addAccessReason(&result, language, reason, resolved, example, options.MaxExamples)
		}
	}
	ensureAccessReasonBuckets(result.Reasons,
		"resolved",
		"missing_receiver_type",
		"missing_owner_link",
		"ambiguous_owner",
		"external_library_type",
		"unsupported_syntax",
		"false_positive_candidate",
		"missing_caller",
	)
	return result, nil
}

func classifyAccessCandidate(w *workspace, access scopeir.AccessFact) (string, bool, string) {
	if strings.TrimSpace(access.Name) == "" {
		return "false_positive_candidate", false, "empty access name"
	}
	if unsupportedAccessReceiver(access.ExplicitReceiver) {
		return "unsupported_syntax", false, "receiver syntax is not modeled by resolver"
	}
	if _, ok := w.callerForScope(access.InScope); !ok {
		return "missing_caller", false, "no caller scope found for access"
	}
	ownerType, ok := w.resolveReceiverType(access.ExplicitReceiver, access.InScope)
	if !ok {
		return "missing_receiver_type", false, "receiver has no resolvable type binding or enclosing owner"
	}
	owner, ok := w.resolveName(ownerType, access.InScope, dispatchOwnerLabels())
	if !ok {
		return "external_library_type", false, "receiver type is not defined in the analyzed workspace"
	}
	members := filterDefRefsByLabel(w.ownerMembers[owner.Fact.ID][access.Name], propertyLabels())
	switch len(members) {
	case 1:
		return "resolved", true, "receiver owner and property member are unique"
	case 0:
		if hasStandalonePropertyCandidate(w, access) {
			return "missing_owner_link", false, "matching property exists but is not linked to the receiver owner"
		}
		return "false_positive_candidate", false, "receiver owner exists but matching property does not"
	default:
		return "ambiguous_owner", false, "receiver owner has multiple matching property members"
	}
}

func unsupportedAccessReceiver(receiver string) bool {
	receiver = strings.TrimSpace(receiver)
	return strings.ContainsAny(receiver, "[](){}")
}

func hasStandalonePropertyCandidate(w *workspace, access scopeir.AccessFact) bool {
	for _, candidate := range w.defsByName[access.Name] {
		if candidate.Fact.Label != scopeir.NodeProperty {
			continue
		}
		if candidate.Fact.OwnerID == "" {
			return true
		}
	}
	return false
}

func filterDefRefsByLabel(refs []defRef, labels []scopeir.NodeLabel) []defRef {
	out := make([]defRef, 0, len(refs))
	for _, ref := range refs {
		if isAnyLabel(ref.Fact.Label, labels) {
			out = append(out, ref)
		}
	}
	return out
}

func addAccessReason(result *AccessCandidateAudit, language string, reason string, resolved bool, example AccessCandidateAuditExample, maxExamples int) {
	bucket := result.Reasons[reason]
	bucket.Count++
	if len(bucket.Examples) < maxExamples {
		bucket.Examples = append(bucket.Examples, example)
	}
	result.Reasons[reason] = bucket

	stats := result.Languages[language]
	if stats.Reasons == nil {
		stats.Reasons = map[string]int{}
	}
	stats.Total++
	if resolved {
		stats.Resolved++
	} else {
		stats.Unresolved++
	}
	stats.Reasons[reason]++
	result.Languages[language] = stats
}

func ensureAccessReasonBuckets(buckets map[string]AccessCandidateAuditBucket, keys ...string) {
	for _, key := range keys {
		if _, ok := buckets[key]; !ok {
			buckets[key] = AccessCandidateAuditBucket{}
		}
	}
}
