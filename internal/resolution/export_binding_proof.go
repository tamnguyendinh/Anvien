package resolution

import (
	"encoding/json"
	"sort"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

const (
	exportBindingTerminalEvidenceKind = "export-terminal-v1"
	exportBindingHopEvidenceKind      = "export-hop-v1"
	exportBindingFailureEvidenceKind  = "export-failure-v1"
)

type exportBindingTerminalNoteV1 struct {
	SourceSiteID    string                       `json:"sourceSiteId"`
	ProofOrdinal    int                          `json:"proofOrdinal"`
	Outcome         exportResolutionOutcome      `json:"outcome"`
	TerminalKind    exportResolutionTerminalKind `json:"terminalKind"`
	RequestedName   string                       `json:"requestedName"`
	Meanings        []scopeir.ExportMeaning      `json:"meanings"`
	TargetFiles     []string                     `json:"targetFiles"`
	TerminalDefID   string                       `json:"terminalDefId"`
	TerminalGraphID string                       `json:"terminalGraphId"`
	TerminalFile    string                       `json:"terminalFile"`
	NamespaceFile   string                       `json:"namespaceFile"`
}

type exportBindingHopNoteV1 struct {
	SourceSiteID     string                  `json:"sourceSiteId"`
	ProofOrdinal     int                     `json:"proofOrdinal"`
	HopOrdinal       int                     `json:"hopOrdinal"`
	Outcome          exportResolutionOutcome `json:"outcome"`
	ExportFact       scopeir.ExportFact      `json:"exportFact"`
	HopKind          exportResolutionHopKind `json:"hopKind"`
	FilePath         string                  `json:"filePath"`
	RequestedName    string                  `json:"requestedName"`
	TargetFile       string                  `json:"targetFile"`
	MemberOwnerDefID string                  `json:"memberOwnerDefId"`
	MemberName       string                  `json:"memberName"`
	TerminalDefID    string                  `json:"terminalDefId"`
}

type exportBindingFailureNoteV1 struct {
	SourceSiteID  string                  `json:"sourceSiteId"`
	ProofOrdinal  int                     `json:"proofOrdinal"`
	Outcome       exportResolutionOutcome `json:"outcome"`
	FailureFile   string                  `json:"failureFile"`
	FailureName   string                  `json:"failureName"`
	NamespaceFile string                  `json:"namespaceFile"`
	Meanings      []scopeir.ExportMeaning `json:"meanings"`
}

type exportBindingEvidenceKey struct {
	kind   string
	weight float64
	note   string
}

type exportBindingEvidenceOrder struct {
	kindRank     int
	proofOrdinal int
	hopOrdinal   int
	note         string
}

// retainedExportResolutionForScopedBinding returns only the semantic result
// that produced the exact scoped target. It mirrors resolveScopedName's
// nearest-scope and ambiguity behavior without performing another traversal.
func (w *workspace) retainedExportResolutionForScopedBinding(name string, startScope string, labels []scopeir.NodeLabel, target defRef) (exportResolutionResult, bool) {
	for scopeID := startScope; scopeID != ""; scopeID = w.parentScope(scopeID) {
		bindings := w.scopeBindings[scopeID][name]
		if len(bindings) == 0 {
			continue
		}
		filtered := filterBindingsByLabel(bindings, labels)
		if len(filtered) != 1 {
			return exportResolutionResult{}, false
		}
		binding := filtered[0]
		if binding.Def.GraphID != target.GraphID ||
			binding.Via == nil ||
			!binding.Via.HasSemanticResult {
			return exportResolutionResult{}, false
		}
		resolved, ok := binding.Via.SemanticResult.definition()
		if !ok || resolved.GraphID != target.GraphID {
			return exportResolutionResult{}, false
		}
		return binding.Via.SemanticResult, true
	}
	return exportResolutionResult{}, false
}

func appendExportBindingEvidence(generic []graph.Evidence, result exportResolutionResult, sourceSiteID string) []graph.Evidence {
	if sourceSiteID == "" {
		return append([]graph.Evidence(nil), generic...)
	}
	proofs := result.allProofs()
	if len(proofs) == 0 {
		return append([]graph.Evidence(nil), generic...)
	}

	projected := make([]graph.Evidence, 0, len(proofs)*2)
	for proofOrdinal, proof := range proofs {
		if proof.Outcome == exportResolutionTerminal {
			note := exportBindingTerminalNoteV1{
				SourceSiteID:    sourceSiteID,
				ProofOrdinal:    proofOrdinal,
				Outcome:         proof.Outcome,
				TerminalKind:    proof.TerminalKind,
				RequestedName:   result.RequestedName,
				Meanings:        canonicalExportMeanings(proof.Meanings),
				TargetFiles:     cloneSortedExportTablePaths(result.TargetFiles),
				TerminalDefID:   proof.Terminal.Fact.ID,
				TerminalGraphID: proof.Terminal.GraphID,
				TerminalFile:    cleanExportTablePath(proof.Terminal.Fact.FilePath),
				NamespaceFile:   cleanExportTablePath(proof.NamespaceFile),
			}
			if evidence, ok := marshalExportBindingTerminalEvidence(note); ok {
				projected = append(projected, evidence)
			}
		}

		for hopOrdinal, hop := range proof.Hops {
			fact := cloneExportTableFact(hop.Fact)
			fact.Meanings = canonicalExportMeanings(fact.Meanings)
			note := exportBindingHopNoteV1{
				SourceSiteID:     sourceSiteID,
				ProofOrdinal:     proofOrdinal,
				HopOrdinal:       hopOrdinal,
				Outcome:          proof.Outcome,
				ExportFact:       fact,
				HopKind:          hop.Kind,
				FilePath:         cleanExportTablePath(hop.FilePath),
				RequestedName:    hop.RequestedName,
				TargetFile:       cleanExportTablePath(hop.TargetFile),
				MemberOwnerDefID: hop.MemberOwnerDefID,
				MemberName:       hop.MemberName,
				TerminalDefID:    hop.TerminalDefID,
			}
			if evidence, ok := marshalExportBindingHopEvidence(note); ok {
				projected = append(projected, evidence)
			}
		}

		if proof.Outcome != exportResolutionTerminal {
			note := exportBindingFailureNoteV1{
				SourceSiteID:  sourceSiteID,
				ProofOrdinal:  proofOrdinal,
				Outcome:       proof.Outcome,
				FailureFile:   cleanExportTablePath(proof.FailureFile),
				FailureName:   proof.FailureName,
				NamespaceFile: cleanExportTablePath(proof.NamespaceFile),
				Meanings:      canonicalExportMeanings(result.RequestedMeanings),
			}
			if evidence, ok := marshalExportBindingFailureEvidence(note); ok {
				projected = append(projected, evidence)
			}
		}
	}
	sortExportBindingEvidence(projected)
	return mergeExportBindingEvidence(generic, projected)
}

func marshalExportBindingTerminalEvidence(note exportBindingTerminalNoteV1) (graph.Evidence, bool) {
	encoded, err := json.Marshal(note)
	if err != nil {
		return graph.Evidence{}, false
	}
	return graph.Evidence{Kind: exportBindingTerminalEvidenceKind, Weight: 1, Note: string(encoded)}, true
}

func marshalExportBindingHopEvidence(note exportBindingHopNoteV1) (graph.Evidence, bool) {
	encoded, err := json.Marshal(note)
	if err != nil {
		return graph.Evidence{}, false
	}
	return graph.Evidence{Kind: exportBindingHopEvidenceKind, Weight: 1, Note: string(encoded)}, true
}

func marshalExportBindingFailureEvidence(note exportBindingFailureNoteV1) (graph.Evidence, bool) {
	encoded, err := json.Marshal(note)
	if err != nil {
		return graph.Evidence{}, false
	}
	return graph.Evidence{Kind: exportBindingFailureEvidenceKind, Weight: 1, Note: string(encoded)}, true
}

// mergeExportBindingEvidence preserves the first generic evidence record and
// stable-unions the remaining generic records. P5-D records are exact-tuple
// deduplicated and sorted independently so edge coalescing is deterministic.
func mergeExportBindingEvidence(existing []graph.Evidence, incoming []graph.Evidence) []graph.Evidence {
	if len(existing) == 0 && len(incoming) == 0 {
		return nil
	}
	seen := make(map[exportBindingEvidenceKey]struct{}, len(existing)+len(incoming))
	generic := make([]graph.Evidence, 0, len(existing)+len(incoming))
	projected := make([]graph.Evidence, 0, len(existing)+len(incoming))
	appendEvidence := func(values []graph.Evidence) {
		for _, evidence := range values {
			key := exportBindingEvidenceKey{kind: evidence.Kind, weight: evidence.Weight, note: evidence.Note}
			if _, duplicate := seen[key]; duplicate {
				continue
			}
			seen[key] = struct{}{}
			if isExportBindingEvidence(evidence) {
				projected = append(projected, evidence)
				continue
			}
			generic = append(generic, evidence)
		}
	}
	appendEvidence(existing)
	appendEvidence(incoming)
	sortExportBindingEvidence(projected)
	return append(generic, projected...)
}

func isExportBindingEvidence(evidence graph.Evidence) bool {
	switch evidence.Kind {
	case exportBindingTerminalEvidenceKind, exportBindingHopEvidenceKind, exportBindingFailureEvidenceKind:
		return true
	default:
		return false
	}
}

func sortExportBindingEvidence(values []graph.Evidence) {
	sort.SliceStable(values, func(i, j int) bool {
		left := exportBindingEvidenceOrderFor(values[i])
		right := exportBindingEvidenceOrderFor(values[j])
		if left.kindRank != right.kindRank {
			return left.kindRank < right.kindRank
		}
		if left.proofOrdinal != right.proofOrdinal {
			return left.proofOrdinal < right.proofOrdinal
		}
		if left.hopOrdinal != right.hopOrdinal {
			return left.hopOrdinal < right.hopOrdinal
		}
		return left.note < right.note
	})
}

func exportBindingEvidenceOrderFor(evidence graph.Evidence) exportBindingEvidenceOrder {
	maxOrdinal := int(^uint(0) >> 1)
	order := exportBindingEvidenceOrder{
		kindRank:     exportBindingEvidenceKindRank(evidence.Kind),
		proofOrdinal: maxOrdinal,
		hopOrdinal:   maxOrdinal,
		note:         evidence.Note,
	}
	switch evidence.Kind {
	case exportBindingTerminalEvidenceKind:
		var note exportBindingTerminalNoteV1
		if json.Unmarshal([]byte(evidence.Note), &note) == nil {
			order.proofOrdinal = note.ProofOrdinal
			order.hopOrdinal = -1
		}
	case exportBindingHopEvidenceKind:
		var note exportBindingHopNoteV1
		if json.Unmarshal([]byte(evidence.Note), &note) == nil {
			order.proofOrdinal = note.ProofOrdinal
			order.hopOrdinal = note.HopOrdinal
		}
	case exportBindingFailureEvidenceKind:
		var note exportBindingFailureNoteV1
		if json.Unmarshal([]byte(evidence.Note), &note) == nil {
			order.proofOrdinal = note.ProofOrdinal
			order.hopOrdinal = -1
		}
	}
	return order
}

func exportBindingEvidenceKindRank(kind string) int {
	switch kind {
	case exportBindingTerminalEvidenceKind:
		return 0
	case exportBindingHopEvidenceKind:
		return 1
	case exportBindingFailureEvidenceKind:
		return 2
	default:
		return 3
	}
}
