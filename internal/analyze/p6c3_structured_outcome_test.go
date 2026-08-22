package analyze

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/resolution"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

func TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage(t *testing.T) {
	resolved := runP6BAnalyzeFixture(t, "p6b-tsstdlib-runtime")
	if len(resolved.ResolutionOutcomes) < len(resolved.TypeScriptAuthorityResults) || len(resolved.ResolutionOutcomes) == 0 {
		t.Fatalf("analyze result outcome inventory = %d, authority inventory = %d", len(resolved.ResolutionOutcomes), len(resolved.TypeScriptAuthorityResults))
	}
	bySourceSite := make(map[string]resolution.ResolutionOutcome, len(resolved.ResolutionOutcomes))
	for _, outcome := range resolved.ResolutionOutcomes {
		if _, duplicate := bySourceSite[outcome.SourceSiteID]; duplicate {
			t.Fatalf("duplicate analyze outcome for source site %q", outcome.SourceSiteID)
		}
		bySourceSite[outcome.SourceSiteID] = outcome
	}
	for _, authority := range resolved.TypeScriptAuthorityResults {
		outcome, ok := bySourceSite[authority.SourceSiteID]
		if !ok || outcome.Status != resolution.ResolutionResolvedExternal ||
			outcome.Stage != resolution.TypeScriptStandardLibraryStage ||
			outcome.Target == nil || !outcome.Target.External || outcome.Authority == nil ||
			!reflect.DeepEqual(*outcome.Authority, authority) {
			t.Fatalf("analyze external outcome/authority drift: outcome=%#v authority=%#v", outcome, authority)
		}
		requireAnalyzeP6C3RelationshipOutcome(t, resolved.Graph, outcome)
	}

	raw, err := json.Marshal(resolved)
	if err != nil {
		t.Fatalf("marshal analyze result: %v", err)
	}
	var projection struct {
		ResolutionOutcomes []resolution.ResolutionOutcome `json:"resolutionOutcomes"`
	}
	if err := json.Unmarshal(raw, &projection); err != nil {
		t.Fatalf("decode analyze result projection: %v", err)
	}
	if !reflect.DeepEqual(projection.ResolutionOutcomes, resolved.ResolutionOutcomes) {
		t.Fatalf("analyze result JSON outcome parity drift:\nwant=%#v\ngot=%#v", resolved.ResolutionOutcomes, projection.ResolutionOutcomes)
	}

	unavailable := runP6BAnalyzeFixture(t, "p6b-tsstdlib-runtime-no-lib")
	wantCapability := make(map[string]resolution.ResolutionOutcome)
	for _, outcome := range unavailable.ResolutionOutcomes {
		if outcome.Status == resolution.ResolutionCapabilityUnavailable {
			wantCapability[outcome.SourceSiteID] = outcome
		}
	}
	if len(wantCapability) == 0 {
		t.Fatalf("built no-lib analyze result has no capability outcomes: %#v", unavailable.ResolutionOutcomes)
	}
	seenGaps := make(map[string]int)
	for _, node := range unavailable.Graph.Nodes {
		if node.Label != scopeir.NodeResolutionGap {
			continue
		}
		note, _ := node.Properties["note"].(string)
		var outcome resolution.ResolutionOutcome
		if json.Unmarshal([]byte(note), &outcome) != nil || outcome.SourceSiteID == "" {
			continue
		}
		want, ok := wantCapability[outcome.SourceSiteID]
		if !ok {
			continue
		}
		if !reflect.DeepEqual(outcome, want) || node.Properties["sourceSiteStatus"] != string(resolution.ResolutionCapabilityUnavailable) {
			t.Fatalf("ResolutionGap capability outcome drift: node=%#v want=%#v got=%#v", node, want, outcome)
		}
		seenGaps[outcome.SourceSiteID]++
	}
	for sourceSiteID := range wantCapability {
		if seenGaps[sourceSiteID] != 1 {
			t.Fatalf("capability source site %q ResolutionGap rows = %d, want 1", sourceSiteID, seenGaps[sourceSiteID])
		}
	}
}

func requireAnalyzeP6C3RelationshipOutcome(t *testing.T, g *graph.Graph, want resolution.ResolutionOutcome) {
	t.Helper()
	matches := 0
	for _, relationship := range g.Relationships {
		for _, evidence := range relationship.Evidence {
			if evidence.Kind != resolution.ResolutionOutcomeEvidenceKind {
				continue
			}
			var got resolution.ResolutionOutcome
			if err := json.Unmarshal([]byte(evidence.Note), &got); err != nil {
				t.Fatalf("decode analyze relationship outcome: %v", err)
			}
			if got.SourceSiteID == want.SourceSiteID {
				matches++
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("analyze relationship outcome drift: want=%#v got=%#v", want, got)
				}
			}
		}
	}
	if matches != 1 {
		t.Fatalf("analyze relationship outcome count for %q = %d, want 1", want.SourceSiteID, matches)
	}
}

func TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus(t *testing.T) {
	result := runP6BAnalyzeFixture(t, "p6b-tsstdlib-runtime-no-lib")
	for _, authority := range result.TypeScriptAuthorityResults {
		if authority.Status != tsstdlib.LookupCapabilityUnavailable {
			continue
		}
		found := false
		for _, outcome := range result.ResolutionOutcomes {
			if outcome.SourceSiteID != authority.SourceSiteID {
				continue
			}
			found = outcome.Status == resolution.ResolutionCapabilityUnavailable &&
				outcome.Reason == string(authority.Reason) &&
				outcome.Authority != nil && reflect.DeepEqual(*outcome.Authority, authority)
		}
		if !found {
			t.Fatalf("accepted capability authority result was not retained: %#v", authority)
		}
	}
}
