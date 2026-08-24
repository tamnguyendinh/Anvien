package graphhealth

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

func TestP6DStructuredResolutionOutcomesDriveDiagnosticPolicy(t *testing.T) {
	tests := []struct {
		name               string
		status             string
		stage              string
		target             string
		authorityStatus    tsstdlib.LookupStatus
		presetClass        string
		presetAction       string
		wantClassification string
		wantActionability  string
	}{
		{
			name:               "repository outcome overrides standard-library-looking target",
			status:             structuredResolutionStatusUnresolved,
			stage:              structuredResolutionStageRepository,
			target:             "fmt.Errorf",
			presetClass:        DiagnosticClassificationStandardLibrary,
			presetAction:       DiagnosticActionabilityNonActionable,
			wantClassification: DiagnosticClassificationInRepoUnresolved,
			wantActionability:  DiagnosticActionabilityAnalyzerGap,
		},
		{
			name:               "profile-excluded authority overrides local-looking target",
			status:             structuredResolutionStatusUnresolved,
			stage:              structuredResolutionStageTypeScriptStandardLib,
			target:             "AmbientCtor",
			authorityStatus:    tsstdlib.LookupProfileExcluded,
			presetClass:        DiagnosticClassificationInRepoUnresolved,
			presetAction:       DiagnosticActionabilityAnalyzerGap,
			wantClassification: DiagnosticClassificationStandardLibrary,
			wantActionability:  DiagnosticActionabilityNonActionable,
		},
		{
			name:               "capability outcome overrides local-looking target",
			status:             structuredResolutionStatusCapabilityUnavailable,
			stage:              structuredResolutionStageTypeScriptStandardLib,
			target:             "AmbientDisabled",
			authorityStatus:    tsstdlib.LookupCapabilityUnavailable,
			presetClass:        DiagnosticClassificationInRepoUnresolved,
			presetAction:       DiagnosticActionabilityAnalyzerGap,
			wantClassification: DiagnosticClassificationStandardLibrary,
			wantActionability:  DiagnosticActionabilityNonActionable,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			carrier := newP6DCarrier(t, test.target, test.status, test.stage, test.authorityStatus)
			carrier.diagnostic.Classification = test.presetClass
			carrier.diagnostic.Actionability = test.presetAction
			got := normalizeDiagnosticMetadata(carrier.diagnostic)
			if got.Classification != test.wantClassification || got.Actionability != test.wantActionability {
				t.Fatalf("structured policy = %s/%s, want %s/%s: %#v", got.Classification, got.Actionability, test.wantClassification, test.wantActionability, got)
			}
		})
	}
}

func TestP6DCapabilityCatalogProofStatesDriveNonActionablePolicy(t *testing.T) {
	tests := []struct {
		name         string
		proofState   tsstdlib.CatalogProofState
		reason       tsstdlib.Reason
		artifactHash string
	}{
		{
			name:         "missing catalog proof",
			proofState:   tsstdlib.CatalogProofMissing,
			reason:       tsstdlib.ReasonCatalogMissing,
			artifactHash: "",
		},
		{
			name:         "rejected catalog proof",
			proofState:   tsstdlib.CatalogProofRejected,
			reason:       tsstdlib.ReasonCatalogSchema,
			artifactHash: "attempted-artifact-hash",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			carrier := newP6DCarrier(t, "AmbientCatalogState", structuredResolutionStatusCapabilityUnavailable, structuredResolutionStageTypeScriptStandardLib, tsstdlib.LookupCapabilityUnavailable)
			authority := p6dMap(carrier.outcome, "authority")
			authority["catalogProofState"] = string(test.proofState)
			authority["reason"] = string(test.reason)
			authority["authorityHash"] = ""
			authority["catalogHash"] = ""
			authority["catalogArtifactHash"] = test.artifactHash
			carrier.outcome["reason"] = string(test.reason)
			carrier.encode(t)

			got := normalizeDiagnosticMetadata(carrier.diagnostic)
			if got.Classification != DiagnosticClassificationStandardLibrary || got.Actionability != DiagnosticActionabilityNonActionable {
				t.Fatalf("catalog proof policy = %s/%s, want standard_library/non_actionable: %#v", got.Classification, got.Actionability, got)
			}
		})
	}
}

func TestP6DStructuredResolutionOutcomesFailClosedOnMalformedOrDriftedCarriers(t *testing.T) {
	tests := []struct {
		name      string
		status    string
		authority tsstdlib.LookupStatus
		rawNote   string
		mutate    func(*p6dCarrier)
	}{
		{name: "malformed JSON", authority: tsstdlib.LookupProfileExcluded, rawNote: "{"},
		{name: "unknown schema", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.outcome["schemaVersion"] = 2 }},
		{name: "outer source-site drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.outcome["sourceSiteId"] = "site:drift" }},
		{name: "outer stage drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.outcome["stage"] = structuredResolutionStageRepository }},
		{name: "missing flat file path", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.diagnostic.FilePath = "" }},
		{name: "missing flat file hash", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.diagnostic.FileHash = "" }},
		{name: "missing flat source", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.diagnostic.Source = "" }},
		{name: "range drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { p6dMap(c.outcome, "range")["startLine"] = 99 }},
		{name: "proof drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { p6dMap(c.outcome, "proof")["kind"] = "drifted-proof" }},
		{name: "empty authority", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.outcome["authority"] = map[string]any{} }},
		{name: "array authority", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.outcome["authority"] = []any{} }},
		{name: "string authority", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { c.outcome["authority"] = "malformed" }},
		{name: "authority source-site drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { p6dMap(c.outcome, "authority")["sourceSiteId"] = "site:authority-drift" }},
		{name: "authority stage drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { p6dMap(c.outcome, "authority")["stage"] = structuredResolutionStageRepository }},
		{name: "authority file-path drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) {
			p6dMap(c.outcome, "authority")["filePath"] = "src/authority-drift.ts"
		}},
		{name: "authority file-hash drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) {
			p6dMap(c.outcome, "authority")["fileHash"] = "hash-authority-drift"
		}},
		{name: "authority range drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) {
			p6dMap(p6dMap(c.outcome, "authority"), "range")["endCol"] = 10
		}},
		{name: "authority site-kind drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) {
			p6dMap(c.outcome, "authority")["siteKind"] = "access"
		}},
		{name: "authority requested-name drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) {
			p6dMap(c.outcome, "authority")["requestedName"] = "AmbientHostileNested"
		}},
		{name: "authority requested-meaning drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) {
			p6dMap(c.outcome, "authority")["requestedMeaning"] = string(tsstdlib.MeaningType)
		}},
		{name: "authority status drift", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) {
			authority := p6dMap(c.outcome, "authority")
			authority["status"] = string(tsstdlib.LookupCapabilityUnavailable)
			authority["reason"] = string(tsstdlib.ReasonDisabledByNoLib)
		}},
		{name: "three-field-only authority", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) {
			c.outcome["authority"] = map[string]any{
				"sourceSiteId": c.diagnostic.SourceSiteID,
				"stage":        c.diagnostic.ResolutionSource,
				"status":       string(tsstdlib.LookupProfileExcluded),
			}
		}},
		{name: "authority reason missing", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { p6dMap(c.outcome, "authority")["reason"] = "" }},
		{name: "authority catalog proof missing", authority: tsstdlib.LookupProfileExcluded, mutate: func(c *p6dCarrier) { p6dMap(c.outcome, "authority")["catalogProofState"] = "" }},
		{name: "capability authority claims resolved symbol", status: structuredResolutionStatusCapabilityUnavailable, authority: tsstdlib.LookupCapabilityUnavailable, mutate: func(c *p6dCarrier) { p6dMap(c.outcome, "authority")["resolvedSymbolId"] = "external:unexpected" }},
		{name: "resolved carrier attached as diagnostic", status: structuredResolutionStatusResolvedExternal, authority: tsstdlib.LookupResolved},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			status := test.status
			if status == "" {
				status = structuredResolutionStatusUnresolved
			}
			carrier := newP6DCarrier(t, "AmbientHostile", status, structuredResolutionStageTypeScriptStandardLib, test.authority)
			carrier.diagnostic.Classification = DiagnosticClassificationStandardLibrary
			carrier.diagnostic.Actionability = DiagnosticActionabilityNonActionable
			if test.mutate != nil {
				test.mutate(&carrier)
			}
			if test.rawNote != "" {
				carrier.diagnostic.Note = test.rawNote
			} else {
				carrier.encode(t)
			}
			got := normalizeDiagnosticMetadata(carrier.diagnostic)
			if got.Classification != DiagnosticClassificationUnclassified || got.Actionability != DiagnosticActionabilityReview {
				t.Fatalf("hostile carrier policy = %s/%s, want unclassified/review: %#v", got.Classification, got.Actionability, got)
			}
		})
	}
}

func TestP6DLegacyDiagnosticWithoutStructuredMarkerKeepsHeuristicPolicy(t *testing.T) {
	got := normalizeDiagnosticMetadata(Diagnostic{
		Kind:             DiagnosticUnresolvedReference,
		TargetText:       "time.Second",
		SourceSiteStatus: "unresolved_local_binding",
		Note:             `{"message":"legacy diagnostic"}`,
	})
	if got.Classification != DiagnosticClassificationStandardLibrary || got.Actionability != DiagnosticActionabilityNonActionable {
		t.Fatalf("legacy policy = %s/%s, want standard_library/non_actionable: %#v", got.Classification, got.Actionability, got)
	}
}

func TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity(t *testing.T) {
	g := graph.New()
	const sourceNodeID = "Function:src/policy.ts:run"
	g.AddNode(graph.Node{
		ID:    sourceNodeID,
		Label: scopeir.NodeFunction,
		Properties: graph.NodeProperties{
			"name":           "run",
			"filePath":       "src/policy.ts",
			"appLayer":       "backend",
			"functionalArea": "resolution",
		},
	})

	specs := []struct {
		target          string
		status          string
		authorityStatus tsstdlib.LookupStatus
	}{
		{target: "AmbientCtor", status: structuredResolutionStatusUnresolved, authorityStatus: tsstdlib.LookupProfileExcluded},
		{target: "AmbientNamespace.member", status: structuredResolutionStatusUnresolved, authorityStatus: tsstdlib.LookupMeaningMismatch},
		{target: "AmbientDisabled", status: structuredResolutionStatusCapabilityUnavailable, authorityStatus: tsstdlib.LookupCapabilityUnavailable},
	}
	for _, spec := range specs {
		carrier := newP6DCarrier(t, spec.target, spec.status, structuredResolutionStageTypeScriptStandardLib, spec.authorityStatus)
		carrier.diagnostic.SourceNodeID = sourceNodeID
		carrier.diagnostic.Classification = DiagnosticClassificationInRepoUnresolved
		carrier.diagnostic.Actionability = DiagnosticActionabilityAnalyzerGap
		if !AppendDiagnosticToNode(g, sourceNodeID, carrier.diagnostic) {
			t.Fatalf("AppendDiagnosticToNode(%q) = false", spec.target)
		}
	}

	graphJSON := p6dRoundTripGraph(t, g)
	inputs := SourceBackedResolutionGapInputs(graphJSON)
	if len(inputs) != 3 {
		t.Fatalf("SourceBackedResolutionGapInputs() len = %d, want 3: %#v", len(inputs), inputs)
	}
	for _, input := range inputs {
		if input.Classification != DiagnosticClassificationStandardLibrary || input.Actionability != DiagnosticActionabilityNonActionable {
			t.Fatalf("gap input %q policy = %s/%s, want standard_library/non_actionable: %#v", input.SourceSiteID, input.Classification, input.Actionability, input)
		}
		gapNode := input.GraphNode()
		gapRelationship := input.GraphRelationship()
		if gapNode.Properties["classification"] != input.Classification ||
			gapNode.Properties["actionability"] != input.Actionability ||
			gapNode.Properties["note"] != input.Note ||
			gapRelationship.SourceSiteStatus != input.SourceSiteStatus ||
			gapRelationship.ProofKind != input.ProofKind ||
			gapRelationship.Evidence[0].Note != input.Note {
			t.Fatalf("gap graph projection drift for %q: input=%#v node=%#v relationship=%#v", input.SourceSiteID, input, gapNode, gapRelationship)
		}
		graphJSON.AddNode(gapNode)
		graphJSON.AddRelationship(gapRelationship)
	}

	finalGraph := p6dRoundTripGraph(t, graphJSON)
	if after := SourceBackedResolutionGapInputs(finalGraph); !reflect.DeepEqual(after, inputs) {
		t.Fatalf("Graph JSON diagnostic/input parity drift:\nwant=%#v\ngot=%#v", inputs, after)
	}
	summary := ComputeSummary(finalGraph)
	if summary.DiagnosticClassificationCounts[DiagnosticClassificationStandardLibrary] != 3 ||
		summary.DiagnosticClassificationCounts[DiagnosticClassificationInRepoUnresolved] != 0 ||
		summary.DiagnosticActionabilityCounts[DiagnosticActionabilityNonActionable] != 3 ||
		summary.DiagnosticActionabilityCounts[DiagnosticActionabilityAnalyzerGap] != 0 {
		t.Fatalf("diagnostic summary parity drift: class=%#v action=%#v", summary.DiagnosticClassificationCounts, summary.DiagnosticActionabilityCounts)
	}
	if summary.ResolutionGapNodeCount != 3 ||
		summary.HasResolutionGapRelationshipCount != 3 ||
		summary.ResolutionGapCount != 3 ||
		summary.ResolutionGapClassificationCounts[DiagnosticClassificationStandardLibrary] != 3 ||
		summary.ResolutionGapClassificationCounts[DiagnosticClassificationInRepoUnresolved] != 0 ||
		summary.ResolutionGapActionabilityCounts[DiagnosticActionabilityNonActionable] != 3 ||
		summary.ResolutionGapActionabilityCounts[DiagnosticActionabilityAnalyzerGap] != 0 ||
		summary.ResolutionHealthBucketCounts[string(ResolutionHealthUnresolvedNonActionable)] != 3 ||
		summary.ResolutionHealthBucketCounts[string(ResolutionHealthInRepoAnalyzerGap)] != 0 {
		t.Fatalf("resolution-gap summary parity drift: %#v", summary)
	}
}

type p6dCarrier struct {
	diagnostic Diagnostic
	outcome    map[string]any
}

func newP6DCarrier(t *testing.T, target string, status string, stage string, authorityStatus tsstdlib.LookupStatus) p6dCarrier {
	t.Helper()
	siteID := "SourceSite:src/policy.ts#call#" + target + "#7#2#7#9"
	proofKind := "none"
	reason := "repository target unresolved"
	if stage == structuredResolutionStageTypeScriptStandardLib {
		proofKind = "typescript-standard-library-authority"
		reason = string(p6dAuthorityReason(authorityStatus))
	}
	if status == structuredResolutionStatusResolvedExternal {
		reason = ""
	}
	outcomeRange := map[string]any{"startLine": 7, "startCol": 2, "endLine": 7, "endCol": 9}
	outcome := map[string]any{
		"schemaVersion":    structuredResolutionOutcomeSchemaVersion,
		"sourceSiteId":     siteID,
		"status":           status,
		"stage":            stage,
		"siteKind":         "call",
		"filePath":         "src/policy.ts",
		"fileHash":         "hash-policy",
		"range":            outcomeRange,
		"requestedName":    target,
		"requestedMeaning": string(tsstdlib.MeaningValue),
		"language":         "typescript",
		"reason":           reason,
		"proof":            map[string]any{"kind": proofKind},
	}
	if stage == structuredResolutionStageTypeScriptStandardLib {
		outcome["authority"] = p6dAuthority(siteID, target, authorityStatus)
	}
	if status == structuredResolutionStatusResolvedExternal {
		outcome["target"] = map[string]any{
			"id":       "tsstdlib:value:" + target,
			"kind":     string(scopeir.NodeExternalSymbol),
			"external": true,
		}
	}
	carrier := p6dCarrier{
		diagnostic: Diagnostic{
			Kind:             DiagnosticUnresolvedReference,
			FactFamily:       "call",
			TargetText:       target,
			ResolutionSource: stage,
			FilePath:         "src/policy.ts",
			FileHash:         "hash-policy",
			StartLine:        7,
			StartCol:         2,
			EndLine:          7,
			EndCol:           9,
			SourceSiteID:     siteID,
			SourceSiteStatus: status,
			ProofKind:        proofKind,
			TargetRole:       "callable",
			Count:            1,
			Source:           stage,
		},
		outcome: outcome,
	}
	carrier.encode(t)
	return carrier
}

func p6dAuthority(siteID string, target string, status tsstdlib.LookupStatus) map[string]any {
	reason := p6dAuthorityReason(status)
	authority := map[string]any{
		"sourceSiteId":        siteID,
		"stage":               structuredResolutionStageTypeScriptStandardLib,
		"filePath":            "src/policy.ts",
		"fileHash":            "hash-policy",
		"range":               map[string]any{"startLine": 7, "startCol": 2, "endLine": 7, "endCol": 9},
		"siteKind":            "call",
		"requestedName":       target,
		"requestedMeaning":    string(tsstdlib.MeaningValue),
		"status":              string(status),
		"reason":              string(reason),
		"authorityKind":       tsstdlib.AuthorityKind,
		"catalogProofState":   string(tsstdlib.CatalogProofReady),
		"authorityHash":       "authority-hash",
		"typescriptVersion":   tsstdlib.TypeScriptVersion,
		"catalogHash":         "catalog-hash",
		"catalogArtifactHash": "catalog-artifact-hash",
		"profileHash":         "profile-hash",
		"configHash":          "config-hash",
	}
	if status == tsstdlib.LookupResolved {
		authority["resolvedSymbolId"] = "tsstdlib:value:" + target
		authority["declarationRanges"] = []map[string]any{{"library": "lib.test.d.ts", "startLine": 1, "startCol": 1, "endLine": 1, "endCol": 9}}
	}
	return authority
}

func p6dAuthorityReason(status tsstdlib.LookupStatus) tsstdlib.Reason {
	switch status {
	case tsstdlib.LookupProfileExcluded:
		return tsstdlib.ReasonProfileExcludes
	case tsstdlib.LookupMeaningMismatch:
		return tsstdlib.ReasonMeaningMismatch
	case tsstdlib.LookupCapabilityUnavailable:
		return tsstdlib.ReasonDisabledByNoLib
	default:
		return ""
	}
}

func (carrier *p6dCarrier) encode(t *testing.T) {
	t.Helper()
	raw, err := json.Marshal(carrier.outcome)
	if err != nil {
		t.Fatalf("marshal P6-D carrier: %v", err)
	}
	carrier.diagnostic.Note = string(raw)
}

func p6dMap(parent map[string]any, key string) map[string]any {
	value, _ := parent[key].(map[string]any)
	return value
}

func p6dRoundTripGraph(t *testing.T, input *graph.Graph) *graph.Graph {
	t.Helper()
	raw, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("marshal Graph JSON: %v", err)
	}
	var output graph.Graph
	if err := json.Unmarshal(raw, &output); err != nil {
		t.Fatalf("unmarshal Graph JSON: %v", err)
	}
	return &output
}
