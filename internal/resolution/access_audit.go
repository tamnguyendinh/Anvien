package resolution

import (
	"strings"

	"github.com/tamnguyendinh/anvien/internal/scopeir"
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
		"non_property_target",
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
	if _, ok := w.resolveImportedMember(access.ExplicitReceiver, access.Name, access.InScope, propertyLabels()); ok {
		return "resolved", true, "imported receiver and property member are unique"
	}
	if target, ok := w.resolveImportedMember(access.ExplicitReceiver, access.Name, access.InScope, accessCandidateRejectedLabels()); ok {
		return "non_property_target", false, "imported receiver resolves to " + string(target.Fact.Label) + ", not Property"
	}
	if w.receiverIsUnresolvedImport(access.ExplicitReceiver, access.InScope) {
		return "external_library_type", false, "receiver is imported but target is not resolved in the analyzed workspace"
	}
	ownerType, ok := w.resolveReceiverType(access.ExplicitReceiver, access.InScope)
	if !ok {
		return "missing_receiver_type", false, "receiver has no resolvable type binding or enclosing owner"
	}
	owner, ok := w.resolveMemberOwner(ownerType, access.InScope)
	if !ok {
		return "external_library_type", false, "receiver type is not defined in the analyzed workspace"
	}
	members := filterDefRefsByLabel(w.ownerMembers[owner.Fact.ID][access.Name], propertyLabels())
	switch len(members) {
	case 1:
		return "resolved", true, "receiver owner and property member are unique"
	case 0:
		if rejected := filterDefRefsByLabel(w.ownerMembers[owner.Fact.ID][access.Name], accessCandidateRejectedLabels()); len(rejected) > 0 {
			return "non_property_target", false, "receiver owner member resolves to " + string(rejected[0].Fact.Label) + ", not Property"
		}
		return "false_positive_candidate", false, "receiver owner exists but matching property does not"
	default:
		return "ambiguous_owner", false, "receiver owner has multiple matching property members"
	}
}

func accessCandidateRejectedLabels() []scopeir.NodeLabel {
	return []scopeir.NodeLabel{scopeir.NodeVariable, scopeir.NodeConst, scopeir.NodeStatic}
}

func unsupportedAccessReceiver(receiver string) bool {
	receiver = strings.TrimSpace(receiver)
	return strings.ContainsAny(receiver, "[](){}")
}

func (w *workspace) receiverIsUnresolvedImport(receiver string, startScope string) bool {
	root := accessReceiverRoot(receiver)
	if root == "" {
		return false
	}
	if _, ok := w.lookupTypeBinding(root, startScope); ok {
		return false
	}
	sourceFile := w.scopeFilePath(startScope)
	if sourceFile == "" {
		return false
	}
	for _, item := range w.imports {
		if item.LinkStatus != "unresolved" {
			continue
		}
		if cleanPath(item.Fact.FilePath) == sourceFile && item.Fact.LocalName == root {
			return true
		}
	}
	return false
}

func accessReceiverRoot(receiver string) string {
	receiver = strings.TrimSpace(receiver)
	if receiver == "" {
		return ""
	}
	if index := strings.Index(receiver, "."); index >= 0 {
		receiver = receiver[:index]
	}
	return strings.TrimSpace(receiver)
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
