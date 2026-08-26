package graphhealth

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
)

func TestDiagnosticAppenderMatchesLegacySemantics(t *testing.T) {
	const nodeID = "File:source.go"

	t.Run("repeated writes", func(t *testing.T) {
		legacyGraph := newDiagnosticAppenderTestGraph(nodeID)
		appenderGraph := newDiagnosticAppenderTestGraph(nodeID)
		appendDiagnostic := NewDiagnosticAppender(appenderGraph)

		incoming := []Diagnostic{
			{
				Kind:             DiagnosticUnresolvedReference,
				FactFamily:       "type-reference",
				TargetText:       "Widget",
				ResolutionSource: "scope-resolution",
				FilePath:         "a.go",
				StartLine:        15,
				StartCol:         2,
				EndLine:          15,
				EndCol:           8,
				SourceSiteID:     "site-type",
				Note:             "type target missing",
			},
			{
				Kind:             DiagnosticUnresolvedReference,
				FactFamily:       "call",
				TargetText:       "stable",
				ResolutionSource: "scope-resolution",
				FilePath:         "same.go",
				StartLine:        50,
				StartCol:         1,
				EndLine:          50,
				EndCol:           7,
				SourceSiteID:     "site-stable-1",
				Note:             "stable-one",
			},
			{
				Kind:             DiagnosticUnresolvedReference,
				FactFamily:       "call",
				TargetText:       "stable",
				ResolutionSource: "scope-resolution",
				FilePath:         "same.go",
				StartLine:        50,
				StartCol:         1,
				EndLine:          50,
				EndCol:           7,
				SourceSiteID:     "site-stable-2",
				Note:             "stable-two",
			},
			{
				Kind:             DiagnosticUnresolvedReference,
				FactFamily:       "call",
				TargetText:       "anotherTarget",
				ResolutionSource: "scope-resolution",
				FilePath:         "c.go",
				StartLine:        18,
				StartCol:         3,
				EndLine:          18,
				EndCol:           16,
				SourceSiteID:     "site-unique",
				Note:             "unique call",
			},
			{
				Kind:             DiagnosticUnresolvedReference,
				FactFamily:       "call",
				TargetText:       "filledTarget",
				ResolutionSource: "scope-resolution",
				FilePath:         "b.go",
				StartLine:        7,
				StartCol:         9,
				EndLine:          7,
				EndCol:           21,
				SourceSiteID:     "site-fill",
				Count:            3,
				Note:             "legacy fill",
			},
		}
		wantLengths := []int{3, 4, 5, 6, 6}

		for index, diagnostic := range incoming {
			if !AppendDiagnosticToNode(legacyGraph, nodeID, diagnostic) {
				t.Fatalf("legacy append %d returned false", index)
			}
			if !appendDiagnostic(nodeID, diagnostic) {
				t.Fatalf("run-scoped append %d returned false", index)
			}

			legacyNode := requireGraphNode(t, legacyGraph, nodeID)
			appenderNode := requireGraphNode(t, appenderGraph, nodeID)
			legacyDiagnostics := requireNodeDiagnostics(t, legacyNode)
			appenderDiagnostics := requireNodeDiagnostics(t, appenderNode)
			if len(appenderDiagnostics) != wantLengths[index] {
				t.Fatalf("append %d diagnostic count = %d, want %d", index, len(appenderDiagnostics), wantLengths[index])
			}
			if !reflect.DeepEqual(legacyDiagnostics, appenderDiagnostics) {
				t.Fatalf("append %d diagnostics differ\nlegacy: %#v\nappender: %#v", index, legacyDiagnostics, appenderDiagnostics)
			}
			if !reflect.DeepEqual(legacyNode, appenderNode) {
				t.Fatalf("append %d graph node differs\nlegacy: %#v\nappender: %#v", index, legacyNode, appenderNode)
			}
			if appenderNode.Properties["sentinel"] != "preserved" {
				t.Fatalf("append %d did not preserve sibling properties", index)
			}
			if got, want := graphJSON(t, appenderGraph), graphJSON(t, legacyGraph); got != want {
				t.Fatalf("append %d graph JSON differs\nlegacy: %s\nappender: %s", index, want, got)
			}
		}

		diagnostics := requireNodeDiagnostics(t, requireGraphNode(t, appenderGraph, nodeID))
		wantOrder := []string{
			"fmt.Println",
			"site-fill",
			"site-unique",
			"site-stable-1",
			"site-stable-2",
			"site-type",
		}
		gotOrder := make([]string, len(diagnostics))
		for index, diagnostic := range diagnostics {
			gotOrder[index] = diagnostic.SourceSiteID
			if gotOrder[index] == "" {
				gotOrder[index] = diagnostic.TargetText
			}
		}
		if !reflect.DeepEqual(gotOrder, wantOrder) {
			t.Fatalf("stable diagnostic order = %v, want %v", gotOrder, wantOrder)
		}

		encodedLegacy := requireDiagnostic(t, diagnostics, func(diagnostic Diagnostic) bool {
			return diagnostic.TargetText == "fmt.Println"
		})
		if encodedLegacy.Classification != DiagnosticClassificationStandardLibrary ||
			encodedLegacy.Actionability != DiagnosticActionabilityNonActionable {
			t.Fatalf("encoded legacy policy = %q/%q, want %q/%q",
				encodedLegacy.Classification,
				encodedLegacy.Actionability,
				DiagnosticClassificationStandardLibrary,
				DiagnosticActionabilityNonActionable,
			)
		}

		merged := requireDiagnostic(t, diagnostics, func(diagnostic Diagnostic) bool {
			return diagnostic.SourceSiteID == "site-fill"
		})
		if merged.Count != 5 || merged.TargetText != "filledTarget" || merged.StartLine != 7 {
			t.Fatalf("merged diagnostic count/target/line = %d/%q/%d, want 5/%q/7", merged.Count, merged.TargetText, merged.StartLine, "filledTarget")
		}
		if merged.StartCol != 4 || merged.EndLine != 30 || merged.EndCol != 10 {
			t.Fatalf("merged diagnostic non-line range changed: %#v", merged)
		}
		if merged.Classification != DiagnosticClassificationUnclassified || merged.Actionability != DiagnosticActionabilityReview {
			t.Fatalf("merged legacy policy = %q/%q, want %q/%q",
				merged.Classification,
				merged.Actionability,
				DiagnosticClassificationUnclassified,
				DiagnosticActionabilityReview,
			)
		}

		appended := requireDiagnostic(t, diagnostics, func(diagnostic Diagnostic) bool {
			return diagnostic.SourceSiteID == "site-type"
		})
		if appended.SourceNodeID != nodeID || appended.Count != 1 ||
			appended.Classification != DiagnosticClassificationInRepoUnresolved ||
			appended.Actionability != DiagnosticActionabilityAnalyzerGap {
			t.Fatalf("incoming defaults/policy were not preserved: %#v", appended)
		}
	})

	t.Run("invalid and absent nodes", func(t *testing.T) {
		legacyGraph := newDiagnosticAppenderTestGraph(nodeID)
		appenderGraph := newDiagnosticAppenderTestGraph(nodeID)
		appendDiagnostic := NewDiagnosticAppender(appenderGraph)
		diagnostic := Diagnostic{Kind: DiagnosticUnresolvedReference, TargetText: "missing"}
		legacyBefore := graphJSON(t, legacyGraph)
		appenderBefore := graphJSON(t, appenderGraph)

		if AppendDiagnosticToNode(nil, nodeID, diagnostic) {
			t.Fatal("legacy nil-graph append returned true")
		}
		if NewDiagnosticAppender(nil)(nodeID, diagnostic) {
			t.Fatal("run-scoped nil-graph append returned true")
		}

		cases := []struct {
			name       string
			nodeID     string
			diagnostic Diagnostic
		}{
			{name: "blank node id", diagnostic: diagnostic},
			{name: "blank diagnostic kind", nodeID: nodeID, diagnostic: Diagnostic{TargetText: "missing"}},
			{name: "absent node", nodeID: "File:absent.go", diagnostic: diagnostic},
		}
		for _, testCase := range cases {
			t.Run(testCase.name, func(t *testing.T) {
				if AppendDiagnosticToNode(legacyGraph, testCase.nodeID, testCase.diagnostic) {
					t.Fatal("legacy append returned true")
				}
				if appendDiagnostic(testCase.nodeID, testCase.diagnostic) {
					t.Fatal("run-scoped append returned true")
				}
			})
		}

		if got := graphJSON(t, legacyGraph); got != legacyBefore {
			t.Fatalf("legacy graph changed after invalid appends\nbefore: %s\nafter: %s", legacyBefore, got)
		}
		if got := graphJSON(t, appenderGraph); got != appenderBefore {
			t.Fatalf("run-scoped graph changed after invalid appends\nbefore: %s\nafter: %s", appenderBefore, got)
		}
	})
}

func newDiagnosticAppenderTestGraph(nodeID string) *graph.Graph {
	g := graph.New()
	g.AddNode(graph.Node{
		ID:    nodeID,
		Label: "File",
		Properties: graph.NodeProperties{
			"sentinel": "preserved",
			DiagnosticPropertyKey: []any{
				map[string]any{
					"kind":             DiagnosticUnresolvedReference,
					"factFamily":       "access",
					"sourceNodeId":     nodeID,
					"targetText":       "fmt.Println",
					"resolutionSource": "legacy",
					"filePath":         "z.go",
					"startLine":        float64(12),
					"startCol":         float64(2),
					"endLine":          float64(12),
					"endCol":           float64(13),
					"count":            float64(1),
					"note":             "legacy encoded",
				},
				map[string]any{
					"kind":             DiagnosticUnresolvedReference,
					"factFamily":       "call",
					"sourceNodeId":     nodeID,
					"resolutionSource": "scope-resolution",
					"filePath":         "b.go",
					"startLine":        float64(30),
					"startCol":         float64(4),
					"endLine":          float64(30),
					"endCol":           float64(10),
					"sourceSiteId":     "site-fill",
					"count":            float64(2),
					"note":             "legacy fill",
				},
			},
		},
	})
	return g
}

func requireGraphNode(t *testing.T, g *graph.Graph, nodeID string) graph.Node {
	t.Helper()
	node, ok := g.GetNode(nodeID)
	if !ok {
		t.Fatalf("graph node %q is absent", nodeID)
	}
	return node
}

func requireNodeDiagnostics(t *testing.T, node graph.Node) []Diagnostic {
	t.Helper()
	diagnostics, ok := node.Properties[DiagnosticPropertyKey].([]Diagnostic)
	if !ok {
		t.Fatalf("diagnostic property type = %T, want []Diagnostic", node.Properties[DiagnosticPropertyKey])
	}
	return diagnostics
}

func requireDiagnostic(t *testing.T, diagnostics []Diagnostic, match func(Diagnostic) bool) Diagnostic {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if match(diagnostic) {
			return diagnostic
		}
	}
	t.Fatal("expected diagnostic was not found")
	return Diagnostic{}
}

func graphJSON(t *testing.T, g *graph.Graph) string {
	t.Helper()
	encoded, err := json.Marshal(g)
	if err != nil {
		t.Fatalf("marshal graph: %v", err)
	}
	return string(encoded)
}

func TestDecodeStructuredResolutionOutcomeCanonicalSingleInterpretationParity(t *testing.T) {
	base := Diagnostic{
		Kind:             DiagnosticUnresolvedReference,
		FactFamily:       "call",
		TargetText:       "Widget",
		ResolutionSource: structuredResolutionStageRepository,
		FilePath:         "src/widget.go",
		FileHash:         "hash-widget",
		StartLine:        7,
		StartCol:         2,
		EndLine:          7,
		EndCol:           9,
		SourceSiteID:     "SourceSite:src/widget.go#call#Widget#7#2#7#9",
		SourceSiteStatus: structuredResolutionStatusUnresolved,
		ProofKind:        "scope-resolution",
		Source:           structuredResolutionStageRepository,
	}

	fields := func() map[string]any {
		return map[string]any{
			"schemaVersion":    structuredResolutionOutcomeSchemaVersion,
			"sourceSiteId":     base.SourceSiteID,
			"status":           structuredResolutionStatusUnresolved,
			"stage":            structuredResolutionStageRepository,
			"siteKind":         base.FactFamily,
			"filePath":         base.FilePath,
			"fileHash":         base.FileHash,
			"range":            map[string]any{"startLine": base.StartLine, "startCol": base.StartCol, "endLine": base.EndLine, "endCol": base.EndCol},
			"requestedName":    base.TargetText,
			"requestedMeaning": "value",
			"language":         "go",
			"reason":           "repository target unresolved",
			"proof":            map[string]any{"kind": base.ProofKind},
		}
	}
	encode := func(value any) string {
		t.Helper()
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal test note: %v", err)
		}
		return string(encoded)
	}
	encodeIndented := func(value any) string {
		t.Helper()
		encoded, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			t.Fatalf("marshal indented test note: %v", err)
		}
		return " \n" + string(encoded) + "\n "
	}
	withNote := func(diagnostic Diagnostic, note string) Diagnostic {
		diagnostic.Note = note
		return diagnostic
	}

	canonicalNote := encode(fields())
	markerlessFields := fields()
	markerlessFields["SchemaVersion"] = markerlessFields["schemaVersion"]
	markerlessFields["Status"] = markerlessFields["status"]
	delete(markerlessFields, "schemaVersion")
	delete(markerlessFields, "status")
	markerlessNote := encode(markerlessFields)

	exactMarkerOnly := base
	exactMarkerOnly.SourceSiteStatus = ""
	unstructured := base
	unstructured.SourceSiteStatus = ""

	missingMarkers := fields()
	delete(missingMarkers, "schemaVersion")
	delete(missingMarkers, "status")
	missingSchema := fields()
	delete(missingSchema, "schemaVersion")
	typedRange := fields()
	typedRange["range"] = "not-an-object"
	unknownField := fields()
	unknownField["futureContract"] = map[string]any{"nested": true, "version": 2}
	conflictingStatus := fields()
	conflictingStatus["status"] = structuredResolutionStatusCapabilityUnavailable
	invalidProof := fields()
	invalidProof["proof"] = map[string]any{"kind": "conflicting-proof"}
	invalidAuthority := fields()
	invalidAuthority["authority"] = map[string]any{"status": "invalid", "catalogProofState": "rejected"}

	testCases := []struct {
		name       string
		diagnostic Diagnostic
	}{
		{name: "status evidence", diagnostic: withNote(base, `{"message":"legacy-shaped"}`)},
		{name: "exact marker evidence", diagnostic: withNote(exactMarkerOnly, canonicalNote)},
		{name: "both evidence sources and exact markers", diagnostic: withNote(base, canonicalNote)},
		{name: "valid markerless structured note", diagnostic: withNote(base, markerlessNote)},
		{name: "unstructured object", diagnostic: withNote(unstructured, `{"message":"legacy-shaped","Status":"not-structured"}`)},
		{name: "empty note with status evidence", diagnostic: withNote(base, " \t\n")},
		{name: "malformed JSON without status evidence", diagnostic: withNote(unstructured, `{"schemaVersion":1,"status":`)},
		{name: "malformed JSON with status evidence", diagnostic: withNote(base, `{"schemaVersion":1,"status":`)},
		{name: "non-object without status evidence", diagnostic: withNote(unstructured, `[1,2,3]`)},
		{name: "non-object with status evidence", diagnostic: withNote(base, `"structured-looking"`)},
		{name: "null without status evidence", diagnostic: withNote(unstructured, `null`)},
		{name: "null with status evidence", diagnostic: withNote(base, `null`)},
		{name: "typed top-level field error", diagnostic: withNote(base, `{"schemaVersion":"one","status":"unresolved"}`)},
		{name: "typed nested field error before later marker", diagnostic: withNote(base, encode(typedRange))},
		{name: "missing both exact markers", diagnostic: withNote(unstructured, encode(missingMarkers))},
		{name: "missing required schema with exact status", diagnostic: withNote(base, encode(missingSchema))},
		{name: "case-variant markers without status evidence", diagnostic: withNote(unstructured, markerlessNote)},
		{name: "duplicate status last valid", diagnostic: withNote(base, `{"status":"capability_unavailable",`+canonicalNote[1:])},
		{name: "duplicate status last conflicting", diagnostic: withNote(base, canonicalNote[:len(canonicalNote)-1]+`,"status":"capability_unavailable"}`)},
		{name: "duplicate typed error remains invalid", diagnostic: withNote(base, `{"schemaVersion":"one",`+canonicalNote[1:])},
		{name: "unknown fields", diagnostic: withNote(base, encode(unknownField))},
		{name: "conflicting status evidence", diagnostic: withNote(base, encode(conflictingStatus))},
		{name: "invalid proof", diagnostic: withNote(base, encode(invalidProof))},
		{name: "invalid authority", diagnostic: withNote(base, encode(invalidAuthority))},
		{name: "outer and nested whitespace", diagnostic: withNote(base, encodeIndented(fields()))},
	}

	preA003Decode := func(diagnostic Diagnostic) (structuredResolutionOutcome, bool, bool) {
		statusMarker := isStructuredResolutionStatus(strings.TrimSpace(diagnostic.SourceSiteStatus))
		note := strings.TrimSpace(diagnostic.Note)
		if note == "" {
			return structuredResolutionOutcome{}, statusMarker, false
		}

		var envelope map[string]json.RawMessage
		if err := json.Unmarshal([]byte(note), &envelope); err != nil {
			return structuredResolutionOutcome{}, statusMarker, false
		}
		_, hasSchemaVersion := envelope["schemaVersion"]
		_, hasStatus := envelope["status"]
		structured := statusMarker || hasSchemaVersion || hasStatus
		if !structured {
			return structuredResolutionOutcome{}, false, false
		}

		var outcome structuredResolutionOutcome
		if err := json.Unmarshal([]byte(note), &outcome); err != nil {
			return structuredResolutionOutcome{}, true, false
		}
		return outcome, true, validStructuredResolutionDiagnostic(diagnostic, outcome)
	}
	preA003Policy := func(diagnostic Diagnostic) (string, string, bool) {
		outcome, structured, valid := preA003Decode(diagnostic)
		if !structured {
			return "", "", false
		}
		if !valid {
			return DiagnosticClassificationUnclassified, DiagnosticActionabilityReview, true
		}
		switch outcome.Status {
		case structuredResolutionStatusUnresolved:
			switch outcome.Stage {
			case structuredResolutionStageRepository:
				return DiagnosticClassificationInRepoUnresolved, DiagnosticActionabilityAnalyzerGap, true
			case structuredResolutionStageTypeScriptStandardLib:
				return DiagnosticClassificationStandardLibrary, DiagnosticActionabilityNonActionable, true
			}
		case structuredResolutionStatusCapabilityUnavailable:
			if outcome.Stage == structuredResolutionStageTypeScriptStandardLib {
				return DiagnosticClassificationStandardLibrary, DiagnosticActionabilityNonActionable, true
			}
		}
		return DiagnosticClassificationUnclassified, DiagnosticActionabilityReview, true
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			wantOutcome, wantStructured, wantValid := preA003Decode(testCase.diagnostic)
			gotOutcome, gotStructured, gotValid := decodeStructuredResolutionOutcome(testCase.diagnostic)
			if !reflect.DeepEqual(gotOutcome, wantOutcome) || gotStructured != wantStructured || gotValid != wantValid {
				t.Fatalf("decode tuple mismatch\nwant: (%#v, %t, %t)\ngot:  (%#v, %t, %t)", wantOutcome, wantStructured, wantValid, gotOutcome, gotStructured, gotValid)
			}

			wantClassification, wantActionability, wantHandled := preA003Policy(testCase.diagnostic)
			gotClassification, gotActionability, gotHandled := structuredResolutionDiagnosticPolicy(testCase.diagnostic)
			if gotClassification != wantClassification || gotActionability != wantActionability || gotHandled != wantHandled {
				t.Fatalf("policy tuple mismatch\nwant: (%q, %q, %t)\ngot:  (%q, %q, %t)", wantClassification, wantActionability, wantHandled, gotClassification, gotActionability, gotHandled)
			}
		})
	}
}
