package resolution

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP5DProjectionEncodesRetainedDirectAndBarrelProofs(t *testing.T) {
	run := p5cDefinition("src/impl.ts", "run", scopeir.NodeFunction)
	files := []scopeir.ScopeIR{
		p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
			p5cImport(scopeir.ImportAlias, "barrelRun", "publicRun", "./barrel", p5cAllMeanings, false),
			p5cImport(scopeir.ImportAlias, "directRun", "run", "./impl", p5cAllMeanings, false),
		}, nil),
		p5cModule("src/barrel.ts", nil, []scopeir.ExportFact{
			p5cReexport("src/barrel.ts", "publicRun", "midRun", "./mid", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		}),
		p5cModule("src/mid.ts", nil, []scopeir.ExportFact{
			p5cReexport("src/mid.ts", "midRun", "run", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		}),
		p5cModule("src/impl.ts", []scopeir.DefinitionFact{run}, []scopeir.ExportFact{
			p5cLocalExport(run, scopeir.ExportDirect, "run", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		}),
	}
	w := p5cWorkspace(t, files)

	tests := []struct {
		localName     string
		requestedName string
		targetFile    string
		hops          int
	}{
		{localName: "barrelRun", requestedName: "publicRun", targetFile: "src/barrel.ts", hops: 3},
		{localName: "directRun", requestedName: "run", targetFile: "src/impl.ts", hops: 1},
	}
	for _, test := range tests {
		t.Run(test.localName, func(t *testing.T) {
			item := p5cResolvedImportByLocalName(t, w, "src/app.ts", test.localName)
			if !item.HasSemanticResult || item.TargetDef == nil || item.TargetDef.Fact.ID != run.ID {
				t.Fatalf("retained semantic result = %#v, want terminal %q", item, run.ID)
			}
			if item.SemanticResult.RequestedName != test.requestedName ||
				!reflect.DeepEqual(item.SemanticResult.TargetFiles, []string{test.targetFile}) {
				t.Fatalf("retained request = %#v, want %q via %q", item.SemanticResult, test.requestedName, test.targetFile)
			}

			siteID := "SourceSite:src/app.ts#call#" + test.localName
			generic := []graph.Evidence{{Kind: "scope-chain", Weight: 1, Note: test.localName}}
			got := appendExportBindingEvidence(generic, item.SemanticResult, siteID)
			again := appendExportBindingEvidence(generic, item.SemanticResult, siteID)
			if !reflect.DeepEqual(got, again) {
				t.Fatalf("projection is not deterministic:\nfirst=%#v\nagain=%#v", got, again)
			}
			if len(got) != 2+test.hops {
				t.Fatalf("evidence count = %d, want generic + terminal + %d hops: %#v", len(got), test.hops, got)
			}
			if got[0] != generic[0] {
				t.Fatalf("generic evidence is not first: %#v", got)
			}
			if got[1].Kind != exportBindingTerminalEvidenceKind || got[1].Weight != 1 {
				t.Fatalf("terminal evidence = %#v", got[1])
			}

			var terminal exportBindingTerminalNoteV1
			if err := json.Unmarshal([]byte(got[1].Note), &terminal); err != nil {
				t.Fatalf("decode terminal note: %v", err)
			}
			if terminal.SourceSiteID != siteID ||
				terminal.ProofOrdinal != 0 ||
				terminal.Outcome != exportResolutionTerminal ||
				terminal.TerminalKind != exportResolutionDefinitionTerminal ||
				terminal.RequestedName != test.requestedName ||
				!reflect.DeepEqual(terminal.Meanings, []scopeir.ExportMeaning{scopeir.ExportMeaningValue}) ||
				!reflect.DeepEqual(terminal.TargetFiles, []string{test.targetFile}) ||
				terminal.TerminalDefID != run.ID ||
				terminal.TerminalGraphID != item.TargetDef.GraphID ||
				terminal.TerminalFile != "src/impl.ts" ||
				terminal.NamespaceFile != "" {
				t.Fatalf("terminal note = %#v", terminal)
			}

			for hopOrdinal, evidence := range got[2:] {
				if evidence.Kind != exportBindingHopEvidenceKind || evidence.Weight != 1 {
					t.Fatalf("hop evidence %d = %#v", hopOrdinal, evidence)
				}
				var hop exportBindingHopNoteV1
				if err := json.Unmarshal([]byte(evidence.Note), &hop); err != nil {
					t.Fatalf("decode hop %d: %v", hopOrdinal, err)
				}
				if hop.SourceSiteID != siteID ||
					hop.ProofOrdinal != 0 ||
					hop.HopOrdinal != hopOrdinal ||
					hop.Outcome != exportResolutionTerminal ||
					hop.HopKind != exportResolutionExportHop ||
					hop.ExportFact.Kind == "" ||
					len(hop.ExportFact.Meanings) == 0 {
					t.Fatalf("hop note %d = %#v", hopOrdinal, hop)
				}
			}
		})
	}
}

func TestP5DProjectionEncodesOwnedMemberProof(t *testing.T) {
	owner := p5cDefinition("src/impl.ts", "API", scopeir.NodeClass)
	run := p5cDefinition("src/impl.ts", "run", scopeir.NodeMethod)
	run.OwnerID = owner.ID
	files := []scopeir.ScopeIR{
		p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
			p5cImport(scopeir.ImportNamed, "api", "publicAPI", "./barrel", p5cAllMeanings, false),
		}, nil),
		p5cModule("src/barrel.ts", nil, []scopeir.ExportFact{
			p5cReexport("src/barrel.ts", "publicAPI", "api", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		}),
		p5cModule("src/impl.ts", []scopeir.DefinitionFact{owner, run}, []scopeir.ExportFact{
			p5cLocalExport(owner, scopeir.ExportDirect, "api", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		}),
	}
	w := p5cWorkspace(t, files)
	target, result, ok := w.resolveImportedMemberWithProof(
		"api",
		"run",
		p5cFunctionScope("src/app.ts"),
		callableLabels(),
	)
	if !ok || target.Fact.ID != run.ID || result.Outcome != exportResolutionTerminal {
		t.Fatalf("member result = %#v target=%#v ok=%t", result, target, ok)
	}

	siteID := "SourceSite:src/app.ts#call#api.run"
	evidence := appendExportBindingEvidence(
		[]graph.Evidence{{Kind: "type-binding", Weight: 1, Note: "run"}},
		result,
		siteID,
	)
	hops := p5dEvidenceOfKind(evidence, exportBindingHopEvidenceKind)
	if len(hops) != 3 {
		t.Fatalf("member hop count = %d, want re-export + direct + member: %#v", len(hops), evidence)
	}
	var member exportBindingHopNoteV1
	if err := json.Unmarshal([]byte(hops[2].Note), &member); err != nil {
		t.Fatalf("decode member hop: %v", err)
	}
	if member.SourceSiteID != siteID ||
		member.ProofOrdinal != 0 ||
		member.HopOrdinal != 2 ||
		member.HopKind != exportResolutionMemberHop ||
		member.MemberOwnerDefID != owner.ID ||
		member.MemberName != "run" ||
		member.TerminalDefID != run.ID {
		t.Fatalf("member hop = %#v", member)
	}
}

func TestP5DProjectionAndCoalescingRetainFailureAndSourceSiteProofs(t *testing.T) {
	run := p5cDefinition("src/impl.ts", "run", scopeir.NodeFunction)
	files := []scopeir.ScopeIR{
		p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
			p5cImport(scopeir.ImportNamed, "run", "run", "./a", p5cAllMeanings, false),
		}, nil),
		p5cModule("src/a.ts", nil, []scopeir.ExportFact{
			p5cStar("src/a.ts", "./b", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			p5cStar("src/a.ts", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		}),
		p5cModule("src/b.ts", nil, []scopeir.ExportFact{
			p5cStar("src/b.ts", "./a", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		}),
		p5cModule("src/impl.ts", []scopeir.DefinitionFact{run}, []scopeir.ExportFact{
			p5cLocalExport(run, scopeir.ExportDirect, "run", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		}),
	}
	w := p5cWorkspace(t, files)
	item := p5cResolvedImportByLocalName(t, w, "src/app.ts", "run")
	if !item.HasSemanticResult || item.SemanticResult.Outcome != exportResolutionTerminal ||
		!p5cHasFailure(item.SemanticResult, exportResolutionCycle) {
		t.Fatalf("retained terminal/cycle result = %#v", item)
	}

	generic := []graph.Evidence{{Kind: "scope-chain", Weight: 1, Note: "run"}}
	left := appendExportBindingEvidence(generic, item.SemanticResult, "site-left")
	right := appendExportBindingEvidence(generic, item.SemanticResult, "site-right")
	merged := mergeExportBindingEvidence(left, right)
	if len(merged) == 0 || merged[0] != generic[0] {
		t.Fatalf("coalesced generic evidence not first: %#v", merged)
	}
	if again := mergeExportBindingEvidence(merged, left); !reflect.DeepEqual(again, merged) {
		t.Fatalf("exact tuple dedupe is not idempotent:\nfirst=%#v\nagain=%#v", merged, again)
	}

	failures := p5dEvidenceOfKind(merged, exportBindingFailureEvidenceKind)
	if len(failures) != 2 {
		t.Fatalf("coalesced cycle failures = %d, want one per source site: %#v", len(failures), merged)
	}
	sites := make(map[string]struct{})
	for _, evidence := range failures {
		var failure exportBindingFailureNoteV1
		if err := json.Unmarshal([]byte(evidence.Note), &failure); err != nil {
			t.Fatalf("decode failure: %v", err)
		}
		if failure.Outcome != exportResolutionCycle ||
			failure.FailureName != "run" ||
			!reflect.DeepEqual(failure.Meanings, canonicalExportMeanings(p5cAllMeanings)) {
			t.Fatalf("failure note = %#v", failure)
		}
		sites[failure.SourceSiteID] = struct{}{}
	}
	if len(sites) != 2 {
		t.Fatalf("distinct source-site failures were dropped: %#v", sites)
	}

	lastRank := -1
	for _, evidence := range merged[1:] {
		if !isExportBindingEvidence(evidence) {
			continue
		}
		rank := exportBindingEvidenceKindRank(evidence.Kind)
		if rank < lastRank {
			t.Fatalf("P5-D evidence rank order drifted: %#v", merged)
		}
		lastRank = rank
	}

	plain := appendExportBindingEvidence(generic, exportResolutionResult{}, "plain-site")
	if !reflect.DeepEqual(plain, generic) {
		t.Fatalf("no-proof path gained export evidence: %#v", plain)
	}
}

func p5dEvidenceOfKind(values []graph.Evidence, kind string) []graph.Evidence {
	out := make([]graph.Evidence, 0)
	for _, evidence := range values {
		if evidence.Kind == kind {
			out = append(out, evidence)
		}
	}
	return out
}
