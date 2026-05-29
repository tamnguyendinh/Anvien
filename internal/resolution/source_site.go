package resolution

import (
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

const (
	sourceSiteStatusResolved               = "resolved"
	sourceSiteStatusUnresolvedLocalBinding = "unresolved_local_binding"
	sourceSiteStatusUnknown                = "unknown"
	sourceSiteStatusUnsupportedSyntax      = "unsupported_syntax"

	proofKindNone                        = "none"
	proofKindScopeBinding                = "scope-binding"
	proofKindSameFile                    = "same-file"
	proofKindGoSamePackage               = "go-same-package"
	proofKindReceiverMember              = "receiver-member"
	proofKindImportMember                = "import-member"
	proofKindGlobalFallbackLowConfidence = "global-fallback-low-confidence"

	targetRoleCallable = "callable"
	targetRoleMember   = "member"
	targetRoleType     = "type"

	unresolvedNoteCallSourceFileLevel = "call source is file-level; resolved edge not emitted"
)

func sourceSiteID(factFamily string, filePath string, targetText string, factRange scopeir.Range) string {
	parts := []string{
		cleanPath(filePath),
		strings.TrimSpace(factFamily),
		strings.TrimSpace(targetText),
		intString(factRange.StartLine),
		intString(factRange.StartCol),
		intString(factRange.EndLine),
		intString(factRange.EndCol),
	}
	return graph.GenerateID("SourceSite", strings.Join(parts, "#"))
}

func targetRoleForFactFamily(factFamily string) string {
	switch factFamily {
	case "call":
		return targetRoleCallable
	case "access":
		return targetRoleMember
	case "type-reference", "heritage":
		return targetRoleType
	default:
		return ""
	}
}
