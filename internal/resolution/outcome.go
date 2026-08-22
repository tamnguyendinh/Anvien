package resolution

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

const (
	ResolutionOutcomeSchemaVersion = 1

	ResolutionStageRepository = "repository"
	ResolutionStageIntrinsic  = "intrinsic"

	ResolutionOutcomeEvidenceKind = "resolution-outcome-v1"
	intrinsicResolutionProofKind  = "language-intrinsic"
)

type ResolutionStatus string

const (
	ResolutionResolvedInternal      ResolutionStatus = "resolved_internal"
	ResolutionResolvedExternal      ResolutionStatus = "resolved_external"
	ResolutionUnresolved            ResolutionStatus = "unresolved"
	ResolutionCapabilityUnavailable ResolutionStatus = "capability_unavailable"
)

type ResolutionTarget struct {
	ID        string `json:"id"`
	GraphID   string `json:"graphId,omitempty"`
	Name      string `json:"name,omitempty"`
	Kind      string `json:"kind"`
	OwnerID   string `json:"ownerId,omitempty"`
	FilePath  string `json:"filePath,omitempty"`
	External  bool   `json:"external,omitempty"`
	Intrinsic bool   `json:"intrinsic,omitempty"`
}

type ResolutionProof struct {
	Kind     string           `json:"kind"`
	Evidence []graph.Evidence `json:"evidence,omitempty"`
}

// ResolutionOutcome is the versioned, final resolver decision for one stable
// source site. Authority is populated only when the accepted TypeScript
// declaration authority handled the site; repository and intrinsic outcomes
// remain ownership-separated from that external proof.
type ResolutionOutcome struct {
	SchemaVersion    int                        `json:"schemaVersion"`
	SourceSiteID     string                     `json:"sourceSiteId"`
	Status           ResolutionStatus           `json:"status"`
	Stage            string                     `json:"stage"`
	SiteKind         string                     `json:"siteKind"`
	FilePath         string                     `json:"filePath"`
	FileHash         string                     `json:"fileHash,omitempty"`
	Range            scopeir.Range              `json:"range"`
	RequestedName    string                     `json:"requestedName"`
	RequestedMeaning string                     `json:"requestedMeaning"`
	Language         string                     `json:"language"`
	Target           *ResolutionTarget          `json:"target,omitempty"`
	Reason           string                     `json:"reason,omitempty"`
	Proof            ResolutionProof            `json:"proof"`
	Authority        *TypeScriptAuthorityResult `json:"authority,omitempty"`
}

type resolutionOutcomeCollector struct {
	bySourceSite map[string]ResolutionOutcome
	err          error
}

func newResolutionOutcomeCollector() resolutionOutcomeCollector {
	return resolutionOutcomeCollector{bySourceSite: make(map[string]ResolutionOutcome)}
}

func (collector *resolutionOutcomeCollector) record(outcome ResolutionOutcome) (ResolutionOutcome, string, bool) {
	outcome = cloneResolutionOutcome(outcome)
	if err := validateResolutionOutcome(outcome); err != nil {
		collector.setError(err)
		return outcome, "", false
	}
	if previous, exists := collector.bySourceSite[outcome.SourceSiteID]; exists {
		if !reflect.DeepEqual(previous, outcome) {
			collector.setError(fmt.Errorf("conflicting final resolution outcome for source site %q", outcome.SourceSiteID))
		}
		encoded, err := marshalResolutionOutcome(previous)
		if err != nil {
			collector.setError(err)
		}
		return cloneResolutionOutcome(previous), encoded, false
	}
	collector.bySourceSite[outcome.SourceSiteID] = cloneResolutionOutcome(outcome)
	encoded, err := marshalResolutionOutcome(outcome)
	if err != nil {
		collector.setError(err)
		return outcome, "", false
	}
	return outcome, encoded, true
}

func (collector *resolutionOutcomeCollector) setError(err error) {
	if collector.err == nil && err != nil {
		collector.err = err
	}
}

func (collector *resolutionOutcomeCollector) finalize() ([]ResolutionOutcome, error) {
	if collector.err != nil {
		return nil, collector.err
	}
	out := make([]ResolutionOutcome, 0, len(collector.bySourceSite))
	for _, outcome := range collector.bySourceSite {
		if err := validateResolutionOutcome(outcome); err != nil {
			return nil, err
		}
		out = append(out, cloneResolutionOutcome(outcome))
	}
	sort.Slice(out, func(left int, right int) bool {
		return out[left].SourceSiteID < out[right].SourceSiteID
	})
	return out, nil
}

func cloneResolutionOutcome(outcome ResolutionOutcome) ResolutionOutcome {
	if outcome.Target != nil {
		target := *outcome.Target
		outcome.Target = &target
	}
	outcome.Proof.Evidence = append([]graph.Evidence(nil), outcome.Proof.Evidence...)
	if outcome.Authority != nil {
		authority := cloneTypeScriptAuthorityResult(*outcome.Authority)
		outcome.Authority = &authority
	}
	return outcome
}

func validateResolutionOutcome(outcome ResolutionOutcome) error {
	if outcome.SchemaVersion != ResolutionOutcomeSchemaVersion ||
		strings.TrimSpace(outcome.SourceSiteID) == "" ||
		strings.TrimSpace(outcome.Stage) == "" ||
		strings.TrimSpace(outcome.SiteKind) == "" ||
		strings.TrimSpace(outcome.FilePath) == "" ||
		outcome.Range.StartLine <= 0 ||
		outcome.Range.EndLine <= 0 ||
		strings.TrimSpace(outcome.RequestedName) == "" ||
		strings.TrimSpace(outcome.RequestedMeaning) == "" ||
		strings.TrimSpace(outcome.Language) == "" ||
		strings.TrimSpace(outcome.Proof.Kind) == "" {
		return fmt.Errorf("incomplete final resolution outcome for source site %q", outcome.SourceSiteID)
	}
	switch outcome.Status {
	case ResolutionResolvedInternal:
		if outcome.Target == nil || outcome.Reason != "" || outcome.Authority != nil {
			return fmt.Errorf("invalid resolved-internal outcome for source site %q", outcome.SourceSiteID)
		}
		if outcome.Stage != ResolutionStageRepository && outcome.Stage != ResolutionStageIntrinsic {
			return fmt.Errorf("invalid resolved-internal stage %q for source site %q", outcome.Stage, outcome.SourceSiteID)
		}
	case ResolutionResolvedExternal:
		if outcome.Target == nil || !outcome.Target.External || outcome.Reason != "" || outcome.Authority == nil || outcome.Stage != TypeScriptStandardLibraryStage {
			return fmt.Errorf("invalid resolved-external outcome for source site %q", outcome.SourceSiteID)
		}
	case ResolutionUnresolved:
		if outcome.Target != nil || strings.TrimSpace(outcome.Reason) == "" {
			return fmt.Errorf("invalid unresolved outcome for source site %q", outcome.SourceSiteID)
		}
		if outcome.Authority == nil && outcome.Stage != ResolutionStageRepository {
			return fmt.Errorf("invalid repository-unresolved stage %q for source site %q", outcome.Stage, outcome.SourceSiteID)
		}
		if outcome.Authority != nil && outcome.Stage != TypeScriptStandardLibraryStage {
			return fmt.Errorf("invalid authority-unresolved stage %q for source site %q", outcome.Stage, outcome.SourceSiteID)
		}
	case ResolutionCapabilityUnavailable:
		if outcome.Target != nil || strings.TrimSpace(outcome.Reason) == "" || outcome.Authority == nil || outcome.Stage != TypeScriptStandardLibraryStage {
			return fmt.Errorf("invalid capability-unavailable outcome for source site %q", outcome.SourceSiteID)
		}
	default:
		return fmt.Errorf("unsupported final resolution status %q for source site %q", outcome.Status, outcome.SourceSiteID)
	}
	if outcome.Target != nil {
		if strings.TrimSpace(outcome.Target.ID) == "" || strings.TrimSpace(outcome.Target.Kind) == "" {
			return fmt.Errorf("incomplete resolution target for source site %q", outcome.SourceSiteID)
		}
		if outcome.Status == ResolutionResolvedInternal && outcome.Stage == ResolutionStageIntrinsic && !outcome.Target.Intrinsic {
			return fmt.Errorf("intrinsic outcome has non-intrinsic target for source site %q", outcome.SourceSiteID)
		}
	}
	if outcome.Authority != nil {
		if err := validateTypeScriptAuthorityResult(*outcome.Authority); err != nil {
			return err
		}
		if outcome.Authority.SourceSiteID != outcome.SourceSiteID || outcome.Authority.Stage != outcome.Stage {
			return fmt.Errorf("authority/outcome identity drift for source site %q", outcome.SourceSiteID)
		}
		switch outcome.Status {
		case ResolutionResolvedExternal:
			if outcome.Authority.Status != tsstdlib.LookupResolved {
				return fmt.Errorf("external authority status drift for source site %q", outcome.SourceSiteID)
			}
		case ResolutionCapabilityUnavailable:
			if outcome.Authority.Status != tsstdlib.LookupCapabilityUnavailable {
				return fmt.Errorf("capability authority status drift for source site %q", outcome.SourceSiteID)
			}
		case ResolutionUnresolved:
			if outcome.Authority.Status != tsstdlib.LookupProfileExcluded && outcome.Authority.Status != tsstdlib.LookupMeaningMismatch {
				return fmt.Errorf("unresolved authority status drift for source site %q", outcome.SourceSiteID)
			}
		}
	}
	return nil
}

func marshalResolutionOutcome(outcome ResolutionOutcome) (string, error) {
	encoded, err := json.Marshal(outcome)
	if err != nil {
		return "", fmt.Errorf("marshal final resolution outcome for source site %q: %w", outcome.SourceSiteID, err)
	}
	return string(encoded), nil
}

func (e *emitter) languageFor(filePath string) string {
	return string(e.fileLanguages[cleanPath(filePath)])
}

func requestedMeaningForSiteKind(siteKind string) string {
	switch siteKind {
	case "type-reference", "heritage":
		return string(tsstdlib.MeaningType)
	default:
		return string(tsstdlib.MeaningValue)
	}
}

func siteKindForReference(kind ReferenceKind) string {
	switch kind {
	case ReferenceCall:
		return "call"
	case ReferenceRead, ReferenceWrite:
		return "access"
	case ReferenceTypeReference:
		return "type-reference"
	case ReferenceInherits:
		return "heritage"
	default:
		return string(kind)
	}
}

func (e *emitter) recordRepositoryResolvedOutcome(target defRef, reference Reference) {
	siteKind := siteKindForReference(reference.Kind)
	e.outcomes.record(ResolutionOutcome{
		SchemaVersion:    ResolutionOutcomeSchemaVersion,
		SourceSiteID:     reference.SourceSiteID,
		Status:           ResolutionResolvedInternal,
		Stage:            ResolutionStageRepository,
		SiteKind:         siteKind,
		FilePath:         cleanPath(reference.FilePath),
		FileHash:         reference.FileHash,
		Range:            reference.Range,
		RequestedName:    strings.TrimSpace(reference.TargetText),
		RequestedMeaning: requestedMeaningForSiteKind(siteKind),
		Language:         e.languageFor(reference.FilePath),
		Target: &ResolutionTarget{
			ID:       target.Fact.ID,
			GraphID:  target.GraphID,
			Name:     target.Fact.Name,
			Kind:     string(target.Fact.Label),
			OwnerID:  target.Fact.OwnerID,
			FilePath: cleanPath(target.Fact.FilePath),
		},
		Proof: ResolutionProof{
			Kind:     firstNonEmpty(reference.ProofKind, proofKindNone),
			Evidence: append([]graph.Evidence(nil), reference.Evidence...),
		},
	})
}

func (e *emitter) recordRepositoryUnresolvedOutcome(factFamily string, targetText string, filePath string, fileHash string, factRange scopeir.Range, reason string, proofKind string, evidence []graph.Evidence) (string, bool) {
	_, encoded, added := e.outcomes.record(ResolutionOutcome{
		SchemaVersion:    ResolutionOutcomeSchemaVersion,
		SourceSiteID:     sourceSiteID(factFamily, filePath, targetText, factRange),
		Status:           ResolutionUnresolved,
		Stage:            ResolutionStageRepository,
		SiteKind:         factFamily,
		FilePath:         cleanPath(filePath),
		FileHash:         fileHash,
		Range:            factRange,
		RequestedName:    strings.TrimSpace(targetText),
		RequestedMeaning: requestedMeaningForSiteKind(factFamily),
		Language:         e.languageFor(filePath),
		Reason:           strings.TrimSpace(reason),
		Proof: ResolutionProof{
			Kind:     firstNonEmpty(proofKind, proofKindNone),
			Evidence: append([]graph.Evidence(nil), evidence...),
		},
	})
	return encoded, added
}

func (e *emitter) recordIntrinsicTypeOutcome(annotation scopeir.TypeAnnotationFact, targetName string) {
	language := e.languageFor(annotation.FilePath)
	e.outcomes.record(ResolutionOutcome{
		SchemaVersion:    ResolutionOutcomeSchemaVersion,
		SourceSiteID:     sourceSiteID("type-reference", annotation.FilePath, annotation.Type.RawName, annotation.Range),
		Status:           ResolutionResolvedInternal,
		Stage:            ResolutionStageIntrinsic,
		SiteKind:         "type-reference",
		FilePath:         cleanPath(annotation.FilePath),
		FileHash:         annotation.FileHash,
		Range:            annotation.Range,
		RequestedName:    targetName,
		RequestedMeaning: string(tsstdlib.MeaningType),
		Language:         language,
		Target: &ResolutionTarget{
			ID:        graph.GenerateID("Intrinsic", language+"#"+string(tsstdlib.MeaningType)+"#"+targetName),
			Name:      targetName,
			Kind:      ResolutionStageIntrinsic,
			Intrinsic: true,
		},
		Proof: ResolutionProof{Kind: intrinsicResolutionProofKind},
	})
}

func (e *emitter) recordTypeScriptOutcome(site typeScriptSourceSite, record TypeScriptAuthorityResult) (ResolutionOutcome, string, bool) {
	status := ResolutionUnresolved
	var target *ResolutionTarget
	switch record.Status {
	case tsstdlib.LookupResolved:
		status = ResolutionResolvedExternal
		target = &ResolutionTarget{
			ID:       record.ResolvedSymbolID,
			GraphID:  externalSymbolGraphID(record.ResolvedSymbolID),
			Name:     record.RequestedName,
			Kind:     string(scopeir.NodeExternalSymbol),
			OwnerID:  record.ResolvedOwnerID,
			External: true,
		}
	case tsstdlib.LookupCapabilityUnavailable:
		status = ResolutionCapabilityUnavailable
	}
	reason := string(record.Reason)
	if status == ResolutionUnresolved && reason == "" {
		reason = string(record.Status)
	}
	authority := cloneTypeScriptAuthorityResult(record)
	return e.outcomes.record(ResolutionOutcome{
		SchemaVersion:    ResolutionOutcomeSchemaVersion,
		SourceSiteID:     record.SourceSiteID,
		Status:           status,
		Stage:            record.Stage,
		SiteKind:         record.SiteKind,
		FilePath:         record.FilePath,
		FileHash:         record.FileHash,
		Range:            record.Range,
		RequestedName:    record.RequestedName,
		RequestedMeaning: string(record.RequestedMeaning),
		Language:         string(scanner.TypeScript),
		Target:           target,
		Reason:           reason,
		Proof:            ResolutionProof{Kind: typeScriptExternalProofKind},
		Authority:        &authority,
	})
}

func (e *emitter) emitTypeScriptOutcomeDiagnostic(site typeScriptSourceSite, outcome ResolutionOutcome, encoded string) {
	diagnostic := graphhealth.Diagnostic{
		Kind:             graphhealth.DiagnosticUnresolvedReference,
		FactFamily:       outcome.SiteKind,
		SourceNodeID:     site.sourceGraphID,
		TargetText:       site.targetText,
		ResolutionSource: outcome.Stage,
		FilePath:         outcome.FilePath,
		FileHash:         outcome.FileHash,
		StartLine:        outcome.Range.StartLine,
		StartCol:         outcome.Range.StartCol,
		EndLine:          outcome.Range.EndLine,
		EndCol:           outcome.Range.EndCol,
		SourceSiteID:     outcome.SourceSiteID,
		SourceSiteStatus: string(outcome.Status),
		ProofKind:        outcome.Proof.Kind,
		TargetRole:       targetRoleForFactFamily(outcome.SiteKind),
		Note:             encoded,
		Source:           outcome.Stage,
	}
	if graphhealth.AppendDiagnosticToNode(e.graph, site.sourceGraphID, diagnostic) {
		e.metrics.UnresolvedReferenceDiagnostics++
		return
	}
	e.metrics.UnattributedUnresolvedReferences++
}

func projectResolutionOutcomes(g *graph.Graph, referenceIndex *ReferenceIndex, outcomes []ResolutionOutcome) error {
	bySourceSite := make(map[string]ResolutionOutcome, len(outcomes))
	encodedBySourceSite := make(map[string]string, len(outcomes))
	for _, outcome := range outcomes {
		if _, duplicate := bySourceSite[outcome.SourceSiteID]; duplicate {
			return fmt.Errorf("duplicate finalized resolution outcome for source site %q", outcome.SourceSiteID)
		}
		encoded, err := marshalResolutionOutcome(outcome)
		if err != nil {
			return err
		}
		bySourceSite[outcome.SourceSiteID] = cloneResolutionOutcome(outcome)
		encodedBySourceSite[outcome.SourceSiteID] = encoded
	}

	resolvedSites := make(map[string]struct{})
	for index := range g.Relationships {
		relationship := g.Relationships[index]
		siteIDs := mergeSourceSiteIDs(relationship.SourceSiteIDs, nil, relationship.SourceSiteID)
		for _, siteID := range siteIDs {
			outcome, ok := bySourceSite[siteID]
			if !ok {
				return fmt.Errorf("resolved relationship %q has no final outcome for source site %q", relationship.ID, siteID)
			}
			if outcome.Status != ResolutionResolvedInternal && outcome.Status != ResolutionResolvedExternal {
				return fmt.Errorf("source site %q is both resolved and %s", siteID, outcome.Status)
			}
			relationship.Evidence = mergeExportBindingEvidence(relationship.Evidence, []graph.Evidence{{
				Kind:   ResolutionOutcomeEvidenceKind,
				Weight: 1,
				Note:   encodedBySourceSite[siteID],
			}})
			resolvedSites[siteID] = struct{}{}
		}
		g.Relationships[index] = relationship
	}

	if referenceIndex != nil {
		if err := projectReferenceIndexOutcomes(referenceIndex.BySourceScope, bySourceSite, encodedBySourceSite); err != nil {
			return err
		}
		if err := projectReferenceIndexOutcomes(referenceIndex.ByTargetDef, bySourceSite, encodedBySourceSite); err != nil {
			return err
		}
	}

	diagnosticSites, err := resolutionOutcomeDiagnosticSites(g, bySourceSite, encodedBySourceSite)
	if err != nil {
		return err
	}
	for siteID := range resolvedSites {
		if _, overlap := diagnosticSites[siteID]; overlap {
			return fmt.Errorf("source site %q has both resolved and unresolved carriers", siteID)
		}
	}
	return nil
}

func projectReferenceIndexOutcomes(buckets map[string][]Reference, outcomes map[string]ResolutionOutcome, encoded map[string]string) error {
	for key, references := range buckets {
		for index := range references {
			outcome, ok := outcomes[references[index].SourceSiteID]
			if !ok {
				return fmt.Errorf("reference index entry has no final outcome for source site %q", references[index].SourceSiteID)
			}
			if outcome.Status != ResolutionResolvedInternal && outcome.Status != ResolutionResolvedExternal {
				return fmt.Errorf("reference index source site %q has non-resolved outcome %q", references[index].SourceSiteID, outcome.Status)
			}
			references[index].Evidence = mergeExportBindingEvidence(references[index].Evidence, []graph.Evidence{{
				Kind:   ResolutionOutcomeEvidenceKind,
				Weight: 1,
				Note:   encoded[references[index].SourceSiteID],
			}})
		}
		buckets[key] = references
	}
	return nil
}

func resolutionOutcomeDiagnosticSites(g *graph.Graph, outcomes map[string]ResolutionOutcome, encoded map[string]string) (map[string]struct{}, error) {
	sites := make(map[string]struct{})
	for _, node := range g.Nodes {
		value, ok := node.Properties[graphhealth.DiagnosticPropertyKey]
		if !ok {
			continue
		}
		raw, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("marshal diagnostics for node %q: %w", node.ID, err)
		}
		var diagnostics []graphhealth.Diagnostic
		if err := json.Unmarshal(raw, &diagnostics); err != nil {
			return nil, fmt.Errorf("decode diagnostics for node %q: %w", node.ID, err)
		}
		for _, diagnostic := range diagnostics {
			outcome, tracked := outcomes[diagnostic.SourceSiteID]
			if !tracked {
				continue
			}
			if outcome.Status == ResolutionResolvedInternal || outcome.Status == ResolutionResolvedExternal {
				return nil, fmt.Errorf("resolved source site %q also has an unresolved diagnostic", outcome.SourceSiteID)
			}
			if diagnostic.Note != encoded[outcome.SourceSiteID] {
				return nil, fmt.Errorf("diagnostic/outcome payload drift for source site %q", outcome.SourceSiteID)
			}
			sites[outcome.SourceSiteID] = struct{}{}
		}
	}
	return sites, nil
}
