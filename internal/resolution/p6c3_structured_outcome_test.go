package resolution

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

func TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix(t *testing.T) {
	tests := []struct {
		name           string
		authority      *tsstdlib.Authority
		call           scopeir.CallSiteFact
		wantStatus     ResolutionStatus
		wantReason     string
		wantDiagnostic bool
	}{
		{
			name:       "resolved-external",
			authority:  p6bAuthority("es2022"),
			call:       p6bCall("max", scopeir.CallMember, "Math", 2),
			wantStatus: ResolutionResolvedExternal,
		},
		{
			name:           "profile-excluded-is-final-unresolved",
			authority:      tsstdlib.NewAuthority(".", []string{"src/main.ts"}),
			call:           p6bCall("Promise", scopeir.CallConstructor, "", 2),
			wantStatus:     ResolutionUnresolved,
			wantReason:     string(tsstdlib.ReasonProfileExcludes),
			wantDiagnostic: true,
		},
		{
			name:           "meaning-mismatch-is-final-unresolved",
			authority:      p6bAuthority("es2022"),
			call:           p6bCall("ArrayLike", scopeir.CallFree, "", 2),
			wantStatus:     ResolutionUnresolved,
			wantReason:     string(tsstdlib.ReasonMeaningMismatch),
			wantDiagnostic: true,
		},
		{
			name:           "capability-unavailable",
			authority:      p6bAuthority("no-lib"),
			call:           p6bCall("Promise", scopeir.CallConstructor, "", 2),
			wantStatus:     ResolutionCapabilityUnavailable,
			wantReason:     string(tsstdlib.ReasonDisabledByNoLib),
			wantDiagnostic: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ir := p6bTypeScriptIR()
			ir.Calls = []scopeir.CallSiteFact{test.call, test.call}
			first, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: test.authority})
			if err != nil {
				t.Fatalf("Resolve() first error = %v", err)
			}
			second, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: test.authority})
			if err != nil {
				t.Fatalf("Resolve() replay error = %v", err)
			}
			if len(first.ResolutionOutcomes) != 1 {
				t.Fatalf("final outcomes = %d, want one per duplicated stable site: %#v", len(first.ResolutionOutcomes), first.ResolutionOutcomes)
			}
			if !reflect.DeepEqual(first.ResolutionOutcomes, second.ResolutionOutcomes) {
				t.Fatalf("final outcome replay drift:\nfirst=%#v\nsecond=%#v", first.ResolutionOutcomes, second.ResolutionOutcomes)
			}
			outcome := first.ResolutionOutcomes[0]
			if outcome.SchemaVersion != ResolutionOutcomeSchemaVersion ||
				outcome.Status != test.wantStatus ||
				outcome.Stage != TypeScriptStandardLibraryStage ||
				outcome.Reason != test.wantReason ||
				outcome.Language != string(scanner.TypeScript) ||
				outcome.Authority == nil ||
				len(first.TypeScriptAuthorityResults) != 1 ||
				!reflect.DeepEqual(*outcome.Authority, first.TypeScriptAuthorityResults[0]) {
				t.Fatalf("final authority outcome = %#v, authority=%#v", outcome, first.TypeScriptAuthorityResults)
			}
			if test.wantStatus == ResolutionResolvedExternal {
				if outcome.Target == nil || !outcome.Target.External || outcome.Target.Intrinsic {
					t.Fatalf("resolved external target ownership = %#v", outcome.Target)
				}
				requireP6C3RelationshipOutcome(t, first.Graph, outcome)
				if diagnostics := p6c3DiagnosticsForSite(t, first.Graph, outcome.SourceSiteID); len(diagnostics) != 0 {
					t.Fatalf("resolved external site also has diagnostics: %#v", diagnostics)
				}
			} else {
				if outcome.Target != nil {
					t.Fatalf("non-resolved outcome fabricated target: %#v", outcome)
				}
				diagnostics := p6c3DiagnosticsForSite(t, first.Graph, outcome.SourceSiteID)
				if !test.wantDiagnostic || len(diagnostics) != 1 {
					t.Fatalf("structured outcome diagnostics = %d, want one: %#v", len(diagnostics), diagnostics)
				}
				requireP6C3DiagnosticOutcome(t, diagnostics[0], outcome)
			}
			requireP6C3NoResolvedUnresolvedOverlap(t, first.Graph)
		})
	}
}

func TestP6C3RepositoryIntrinsicAndExternalPrecedence(t *testing.T) {
	ir := p6bTypeScriptIR()
	local := scopeir.DefinitionFact{
		ID:            "def:src/main.ts#2:0:Function:parseInt",
		FilePath:      ir.FilePath,
		FileHash:      ir.FileHash,
		Name:          "parseInt",
		QualifiedName: "parseInt",
		Label:         scopeir.NodeFunction,
		Range:         scopeir.Range{StartLine: 2, EndLine: 2},
	}
	ir.Definitions = append(ir.Definitions, local)
	ir.Scopes[1].OwnedDefIDs = append(ir.Scopes[1].OwnedDefIDs, local.ID)
	ir.Scopes[1].Bindings = append(ir.Scopes[1].Bindings, scopeir.BindingFact{
		Name: local.Name, DefID: local.ID, Origin: scopeir.BindingLocal,
	})
	ir.Calls = []scopeir.CallSiteFact{
		p6bCall("parseInt", scopeir.CallFree, "", 3),
		p6bCall("DefinitelyMissing", scopeir.CallFree, "", 4),
	}
	ir.TypeAnnotations = []scopeir.TypeAnnotationFact{{
		FilePath: ir.FilePath,
		FileHash: ir.FileHash,
		Name:     "value",
		Range:    scopeir.Range{StartLine: 5, StartCol: 2, EndLine: 5, EndCol: 8},
		InScope:  p6bFunctionScope,
		Type:     scopeir.TypeRef{RawName: "string", Source: scopeir.TypeSourceAnnotation},
	}}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("es2022")})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	if len(result.ResolutionOutcomes) != 3 {
		t.Fatalf("final outcomes = %d, want repository/intrinsic/unresolved: %#v", len(result.ResolutionOutcomes), result.ResolutionOutcomes)
	}
	byName := make(map[string]ResolutionOutcome, len(result.ResolutionOutcomes))
	for _, outcome := range result.ResolutionOutcomes {
		byName[outcome.RequestedName] = outcome
	}
	repository := byName["parseInt"]
	if repository.Status != ResolutionResolvedInternal || repository.Stage != ResolutionStageRepository ||
		repository.Target == nil || repository.Target.ID != local.ID || repository.Authority != nil {
		t.Fatalf("repository result did not remain authoritative before external lookup: %#v", repository)
	}
	requireP6C3RelationshipOutcome(t, result.Graph, repository)

	intrinsic := byName["string"]
	if intrinsic.Status != ResolutionResolvedInternal || intrinsic.Stage != ResolutionStageIntrinsic ||
		intrinsic.Target == nil || !intrinsic.Target.Intrinsic || intrinsic.Authority != nil {
		t.Fatalf("intrinsic final outcome = %#v", intrinsic)
	}
	if p6c3CarrierCount(result.Graph, intrinsic.SourceSiteID) != 0 {
		t.Fatalf("intrinsic result invented graph carrier: %#v", intrinsic)
	}

	unresolved := byName["DefinitelyMissing"]
	if unresolved.Status != ResolutionUnresolved || unresolved.Stage != ResolutionStageRepository ||
		unresolved.Target != nil || unresolved.Reason != "call target not resolved" || unresolved.Authority != nil {
		t.Fatalf("repository unresolved outcome = %#v", unresolved)
	}
	diagnostics := p6c3DiagnosticsForSite(t, result.Graph, unresolved.SourceSiteID)
	if len(diagnostics) != 1 {
		t.Fatalf("repository unresolved diagnostics = %d, want one: %#v", len(diagnostics), diagnostics)
	}
	requireP6C3DiagnosticOutcome(t, diagnostics[0], unresolved)
	requireP6C3NoResolvedUnresolvedOverlap(t, result.Graph)
}

func TestP6C3RepositoryTypeNamedStringIsTerminalBeforeIntrinsic(t *testing.T) {
	ir := p6bTypeScriptIR()
	repositoryType := scopeir.DefinitionFact{
		ID:            "def:src/main.ts#2:0:TypeAlias:string",
		FilePath:      ir.FilePath,
		FileHash:      ir.FileHash,
		Name:          "string",
		QualifiedName: "string",
		Label:         scopeir.NodeTypeAlias,
		Range:         scopeir.Range{StartLine: 2, EndLine: 2},
	}
	ir.Definitions = append(ir.Definitions, repositoryType)
	ir.Scopes[0].OwnedDefIDs = append(ir.Scopes[0].OwnedDefIDs, repositoryType.ID)
	ir.Scopes[0].Bindings = append(ir.Scopes[0].Bindings, scopeir.BindingFact{
		Name: repositoryType.Name, DefID: repositoryType.ID, Origin: scopeir.BindingLocal,
	})
	annotation := scopeir.TypeAnnotationFact{
		FilePath: ir.FilePath,
		FileHash: ir.FileHash,
		Name:     "value",
		Range:    scopeir.Range{StartLine: 4, StartCol: 2, EndLine: 4, EndCol: 8},
		InScope:  p6bFunctionScope,
		Type:     scopeir.TypeRef{RawName: "string", Source: scopeir.TypeSourceAnnotation},
	}
	ir.TypeAnnotations = []scopeir.TypeAnnotationFact{annotation}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("es2022")})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	if len(result.ResolutionOutcomes) != 1 {
		t.Fatalf("repository/predeclared collision outcomes = %d, want exactly one: %#v", len(result.ResolutionOutcomes), result.ResolutionOutcomes)
	}
	outcome := result.ResolutionOutcomes[0]
	wantSiteID := sourceSiteID("type-reference", annotation.FilePath, annotation.Type.RawName, annotation.Range)
	wantTarget := &ResolutionTarget{
		ID:       repositoryType.ID,
		GraphID:  graphIDForDef(repositoryType, p6bModuleScope),
		Name:     repositoryType.Name,
		Kind:     string(scopeir.NodeTypeAlias),
		FilePath: repositoryType.FilePath,
	}
	wantEvidence := []graph.Evidence{{Kind: "scope-chain", Weight: 1, Note: annotation.Type.RawName}}
	if outcome.SchemaVersion != ResolutionOutcomeSchemaVersion ||
		outcome.SourceSiteID != wantSiteID ||
		outcome.Status != ResolutionResolvedInternal ||
		outcome.Stage != ResolutionStageRepository ||
		outcome.SiteKind != "type-reference" ||
		outcome.RequestedName != annotation.Type.RawName ||
		outcome.RequestedMeaning != string(tsstdlib.MeaningType) ||
		outcome.Language != string(scanner.TypeScript) ||
		outcome.Reason != "" ||
		outcome.Authority != nil ||
		!reflect.DeepEqual(outcome.Target, wantTarget) ||
		outcome.Proof.Kind != proofKindScopeBinding ||
		!reflect.DeepEqual(outcome.Proof.Evidence, wantEvidence) {
		t.Fatalf("repository/predeclared collision outcome = %#v, want target=%#v evidence=%#v", outcome, wantTarget, wantEvidence)
	}
	if len(result.TypeScriptAuthorityResults) != 0 {
		t.Fatalf("terminal repository type reached external authority: %#v", result.TypeScriptAuthorityResults)
	}
	requireP6C3RelationshipOutcome(t, result.Graph, outcome)
	if diagnostics := p6c3DiagnosticsForSite(t, result.Graph, wantSiteID); len(diagnostics) != 0 {
		t.Fatalf("resolved repository/predeclared collision also has diagnostics: %#v", diagnostics)
	}
	requireP6C3NoResolvedUnresolvedOverlap(t, result.Graph)
}

func TestP6C3LocalReceiverMissingMemberFinalizesCallWithoutCarrier(t *testing.T) {
	ir := p6bTypeScriptIR()
	mathRange := scopeir.Range{StartLine: 2, StartCol: 2, EndLine: 2, EndCol: 6}
	math := scopeir.DefinitionFact{
		ID:            "def:src/main.ts#2:2:Variable:Math",
		FilePath:      ir.FilePath,
		FileHash:      ir.FileHash,
		Name:          "Math",
		QualifiedName: "Math",
		Label:         scopeir.NodeVariable,
		Range:         mathRange,
	}
	ir.Definitions = append(ir.Definitions, math)
	ir.Scopes[1].OwnedDefIDs = append(ir.Scopes[1].OwnedDefIDs, math.ID)
	ir.Scopes[1].Bindings = append(ir.Scopes[1].Bindings, scopeir.BindingFact{
		Name: math.Name, DefID: math.ID, Origin: scopeir.BindingLocal,
	})
	ir.BindingLeaves = []scopeir.BindingLeafFact{{
		FilePath: math.FilePath,
		FileHash: math.FileHash,
		Name:     math.Name,
		Range:    mathRange,
	}}
	call := p6bCall("missing", scopeir.CallMember, math.Name, 3)
	ir.Calls = []scopeir.CallSiteFact{call}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	if len(result.ResolutionOutcomes) != 2 {
		t.Fatalf("receiver access plus call outcomes = %d, want exactly two stable sites: %#v", len(result.ResolutionOutcomes), result.ResolutionOutcomes)
	}
	callSiteID := sourceSiteID("call", call.FilePath, callTargetText(call), call.Range)
	accessSiteID := sourceSiteID("access", call.FilePath, math.Name, call.Range)
	bySite := make(map[string]ResolutionOutcome, len(result.ResolutionOutcomes))
	for _, outcome := range result.ResolutionOutcomes {
		if _, duplicate := bySite[outcome.SourceSiteID]; duplicate {
			t.Fatalf("duplicate final outcome for stable site %q: %#v", outcome.SourceSiteID, result.ResolutionOutcomes)
		}
		bySite[outcome.SourceSiteID] = outcome
	}
	callOutcome, ok := bySite[callSiteID]
	if !ok {
		t.Fatalf("missing final call outcome %q: %#v", callSiteID, result.ResolutionOutcomes)
	}
	wantEvidence := []graph.Evidence{{Kind: proofKindScopeBinding, Weight: 1, Note: math.Name}}
	if callOutcome.Status != ResolutionUnresolved ||
		callOutcome.Stage != ResolutionStageRepository ||
		callOutcome.SiteKind != "call" ||
		callOutcome.RequestedName != "Math.missing" ||
		callOutcome.Target != nil ||
		callOutcome.Reason != "repository receiver resolved but call member not found" ||
		callOutcome.Proof.Kind != proofKindReceiverMember ||
		!reflect.DeepEqual(callOutcome.Proof.Evidence, wantEvidence) ||
		callOutcome.Authority != nil {
		t.Fatalf("local missing-member call outcome = %#v, want evidence=%#v", callOutcome, wantEvidence)
	}
	accessOutcome, ok := bySite[accessSiteID]
	if !ok || accessOutcome.Status != ResolutionResolvedInternal ||
		accessOutcome.Stage != ResolutionStageRepository ||
		accessOutcome.Target == nil || accessOutcome.Target.ID != math.ID ||
		accessOutcome.Proof.Kind != proofKindScopeBinding {
		t.Fatalf("receiver access did not finalize independently: %#v", accessOutcome)
	}
	requireP6C3RelationshipOutcome(t, result.Graph, accessOutcome)
	if carriers := p6c3CarrierCount(result.Graph, callSiteID); carriers != 0 {
		t.Fatalf("missing-member call unexpectedly emitted optional carrier count=%d", carriers)
	}
	if diagnostics := p6c3DiagnosticsForSite(t, result.Graph, callSiteID); len(diagnostics) != 0 {
		t.Fatalf("missing-member call unexpectedly emitted diagnostics: %#v", diagnostics)
	}
	requireP6C3NoResolvedUnresolvedOverlap(t, result.Graph)
}

func TestP6C3CompatibilityDisabledHeritageFinalizesWithoutReferenceCarrier(t *testing.T) {
	const (
		filePath    = "src/heritage.ts"
		moduleScope = "scope:src/heritage.ts:module"
		baseScope   = "scope:src/heritage.ts:Base"
		childScope  = "scope:src/heritage.ts:Child"
	)
	parent := moduleScope
	base := scopeir.DefinitionFact{
		ID: "def:src/heritage.ts#1:0:Class:Base", FilePath: filePath, FileHash: "hash-heritage",
		Name: "Base", QualifiedName: "Base", Label: scopeir.NodeClass, Range: scopeir.Range{StartLine: 1, EndLine: 3},
	}
	child := scopeir.DefinitionFact{
		ID: "def:src/heritage.ts#5:0:Class:Child", FilePath: filePath, FileHash: "hash-heritage",
		Name: "Child", QualifiedName: "Child", Label: scopeir.NodeClass, Range: scopeir.Range{StartLine: 5, EndLine: 8},
	}
	heritage := scopeir.HeritageFact{
		FilePath: filePath,
		FileHash: "hash-heritage",
		Name:     base.Name,
		Kind:     scopeir.HeritageExtends,
		Range:    scopeir.Range{StartLine: 5, StartCol: 20, EndLine: 5, EndCol: 24},
		InScope:  childScope,
	}
	ir := scopeir.ScopeIR{
		FilePath: filePath, FileHash: "hash-heritage", Language: scanner.TypeScript, ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{
				ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: filePath, FileHash: "hash-heritage",
				Range: scopeir.Range{StartLine: 1, EndLine: 8},
				Bindings: []scopeir.BindingFact{
					{Name: base.Name, DefID: base.ID, Origin: scopeir.BindingLocal},
					{Name: child.Name, DefID: child.ID, Origin: scopeir.BindingLocal},
				},
			},
			{ID: baseScope, Parent: &parent, Kind: scopeir.ScopeClass, FilePath: filePath, Range: base.Range, OwnedDefIDs: []string{base.ID}},
			{ID: childScope, Parent: &parent, Kind: scopeir.ScopeClass, FilePath: filePath, Range: child.Range, OwnedDefIDs: []string{child.ID}},
		},
		Definitions: []scopeir.DefinitionFact{base, child},
		Heritage:    []scopeir.HeritageFact{heritage},
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{DisableScopeInheritsCompatibility: true})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	if len(result.ResolutionOutcomes) != 1 {
		t.Fatalf("compatibility-disabled heritage outcomes = %d, want exactly one: %#v", len(result.ResolutionOutcomes), result.ResolutionOutcomes)
	}
	outcome := result.ResolutionOutcomes[0]
	wantSiteID := sourceSiteID("heritage", heritage.FilePath, heritage.Name, heritage.Range)
	wantTarget := &ResolutionTarget{
		ID:       base.ID,
		GraphID:  graphIDForDef(base, baseScope),
		Name:     base.Name,
		Kind:     string(scopeir.NodeClass),
		FilePath: base.FilePath,
	}
	wantEvidence := []graph.Evidence{{Kind: "scope-chain", Weight: 1, Note: "extends Base"}}
	if outcome.SourceSiteID != wantSiteID ||
		outcome.Status != ResolutionResolvedInternal ||
		outcome.Stage != ResolutionStageRepository ||
		outcome.SiteKind != "heritage" ||
		outcome.RequestedName != base.Name ||
		outcome.RequestedMeaning != string(tsstdlib.MeaningType) ||
		!reflect.DeepEqual(outcome.Target, wantTarget) ||
		outcome.Proof.Kind != proofKindScopeBinding ||
		!reflect.DeepEqual(outcome.Proof.Evidence, wantEvidence) ||
		outcome.Authority != nil {
		t.Fatalf("compatibility-disabled heritage outcome = %#v, want target=%#v evidence=%#v", outcome, wantTarget, wantEvidence)
	}
	primary := 0
	compatibility := 0
	for _, relationship := range result.Graph.Relationships {
		if relationship.SourceID != graphIDForDef(child, childScope) || relationship.TargetID != graphIDForDef(base, baseScope) {
			continue
		}
		switch relationship.Type {
		case graph.RelExtends:
			primary++
		case graph.RelInherits:
			compatibility++
		}
	}
	if primary != 1 || compatibility != 0 {
		t.Fatalf("heritage carriers primary=%d compatibility=%d, want 1/0", primary, compatibility)
	}
	for _, references := range result.ReferenceIndex.BySourceScope {
		for _, reference := range references {
			if reference.Kind == ReferenceInherits || reference.SourceSiteID == wantSiteID {
				t.Fatalf("compatibility-disabled heritage leaked reference carrier: %#v", reference)
			}
		}
	}
	if carriers := p6c3CarrierCount(result.Graph, wantSiteID); carriers != 0 {
		t.Fatalf("compatibility-disabled heritage leaked source-site carrier count=%d", carriers)
	}
	if diagnostics := p6c3DiagnosticsForSite(t, result.Graph, wantSiteID); len(diagnostics) != 0 {
		t.Fatalf("resolved heritage also has diagnostics: %#v", diagnostics)
	}
	requireP6C3NoResolvedUnresolvedOverlap(t, result.Graph)
}

func TestP6C3P5ProofNestingConflictAndImmutableReplay(t *testing.T) {
	run := p6c3Definition("src/impl.ts", "run", scopeir.NodeFunction)
	impl := p6c3Module("src/impl.ts", []scopeir.DefinitionFact{run}, []scopeir.ExportFact{
		p6c3LocalExport(run, scopeir.ExportDirect, "run", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	barrel := p6c3Module("src/barrel.ts", nil, []scopeir.ExportFact{
		p6c3Namespace("src/barrel.ts", "api", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningNamespace}, false),
	})
	consumer := p6c3Consumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
		p6c3Import(scopeir.ImportNamed, "api", "api", "./barrel", p6c3AllMeanings, false),
	}, []scopeir.CallSiteFact{{Name: "run", ExplicitReceiver: "api", CallForm: scopeir.CallMember}})

	result, err := Resolve([]scopeir.ScopeIR{consumer, barrel, impl}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	var outcome ResolutionOutcome
	for _, candidate := range result.ResolutionOutcomes {
		if candidate.RequestedName == "api.run" {
			outcome = candidate
			break
		}
	}
	if outcome.SourceSiteID == "" || outcome.Status != ResolutionResolvedInternal ||
		len(p6c3EvidenceOfKind(outcome.Proof.Evidence, exportBindingTerminalEvidenceKind)) != 1 ||
		len(p6c3EvidenceOfKind(outcome.Proof.Evidence, exportBindingHopEvidenceKind)) != 2 {
		t.Fatalf("P5 terminal/hop proof was not nested losslessly: %#v", outcome)
	}
	requireP6C3RelationshipOutcome(t, result.Graph, outcome)

	failureConsumer := p6c3Consumer("src/failure.ts", scanner.TypeScript, []scopeir.ImportFact{
		p6c3Import(scopeir.ImportNamed, "api", "api", "./barrel", p6c3AllMeanings, false),
	}, []scopeir.CallSiteFact{{Name: "missing", ExplicitReceiver: "api", CallForm: scopeir.CallMember}})
	failureResult, err := Resolve([]scopeir.ScopeIR{failureConsumer, barrel, impl}, Options{})
	if err != nil {
		t.Fatalf("Resolve() P5 failure error = %v", err)
	}
	var failureOutcome ResolutionOutcome
	for _, candidate := range failureResult.ResolutionOutcomes {
		if candidate.RequestedName == "api.missing" {
			failureOutcome = candidate
			break
		}
	}
	if failureOutcome.SourceSiteID == "" || failureOutcome.Status != ResolutionUnresolved ||
		len(p6c3EvidenceOfKind(failureOutcome.Proof.Evidence, exportBindingFailureEvidenceKind)) != 1 ||
		len(p6c3EvidenceOfKind(failureOutcome.Proof.Evidence, exportBindingHopEvidenceKind)) != 1 {
		t.Fatalf("P5 failure/hop proof was not nested losslessly: %#v", failureOutcome)
	}
	failureDiagnostics := p6c3DiagnosticsForSite(t, failureResult.Graph, failureOutcome.SourceSiteID)
	if len(failureDiagnostics) != 1 {
		t.Fatalf("P5 failure diagnostics = %d, want one: %#v", len(failureDiagnostics), failureDiagnostics)
	}
	requireP6C3DiagnosticOutcome(t, failureDiagnostics[0], failureOutcome)

	collector := newResolutionOutcomeCollector()
	if _, _, added := collector.record(outcome); !added {
		t.Fatal("first valid outcome was not added")
	}
	conflict := cloneResolutionOutcome(outcome)
	conflict.Status = ResolutionUnresolved
	conflict.Target = nil
	conflict.Reason = "conflicting repository finalization"
	if _, _, added := collector.record(conflict); added {
		t.Fatal("conflicting outcome was added")
	}
	if _, err := collector.finalize(); err == nil {
		t.Fatal("conflicting outcomes for one stable site must fail closed")
	}

	immutable := newResolutionOutcomeCollector()
	immutable.record(outcome)
	first, err := immutable.finalize()
	if err != nil {
		t.Fatalf("finalize first immutable snapshot: %v", err)
	}
	first[0].Target.Name = "mutated"
	first[0].Proof.Evidence[0].Note = "mutated"
	second, err := immutable.finalize()
	if err != nil {
		t.Fatalf("finalize second immutable snapshot: %v", err)
	}
	if second[0].Target.Name == "mutated" || second[0].Proof.Evidence[0].Note == "mutated" {
		t.Fatalf("final outcome snapshot leaked caller mutation: %#v", second[0])
	}

	overlapGraph := graph.New()
	overlapGraph.AddNode(graph.Node{ID: "Function:source", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "source"}})
	overlapGraph.AddNode(graph.Node{ID: outcome.Target.GraphID, Label: scopeir.NodeFunction, Properties: graph.NodeProperties{"name": "target"}})
	overlapGraph.AddRelationship(graph.Relationship{
		ID: "rel:overlap", SourceID: "Function:source", TargetID: outcome.Target.GraphID, Type: graph.RelCalls,
		SourceSiteID: outcome.SourceSiteID, SourceSiteIDs: []string{outcome.SourceSiteID}, SourceSiteCount: 1,
	})
	encoded, err := marshalResolutionOutcome(outcome)
	if err != nil {
		t.Fatalf("marshal overlap outcome: %v", err)
	}
	if !graphhealth.AppendDiagnosticToNode(overlapGraph, "Function:source", graphhealth.Diagnostic{
		Kind: graphhealth.DiagnosticUnresolvedReference, SourceSiteID: outcome.SourceSiteID, Note: encoded,
	}) {
		t.Fatal("failed to construct overlap diagnostic")
	}
	if err := projectResolutionOutcomes(overlapGraph, nil, []ResolutionOutcome{outcome}); err == nil {
		t.Fatal("resolved+unresolved overlap must fail closed")
	}
}

func requireP6C3RelationshipOutcome(t *testing.T, g *graph.Graph, want ResolutionOutcome) {
	t.Helper()
	matches := 0
	for _, relationship := range g.Relationships {
		for _, evidence := range relationship.Evidence {
			if evidence.Kind != ResolutionOutcomeEvidenceKind {
				continue
			}
			var got ResolutionOutcome
			if err := json.Unmarshal([]byte(evidence.Note), &got); err != nil {
				t.Fatalf("decode relationship outcome evidence: %v", err)
			}
			if got.SourceSiteID != want.SourceSiteID {
				continue
			}
			matches++
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("relationship outcome drift:\nwant=%#v\ngot=%#v", want, got)
			}
		}
	}
	if matches != 1 {
		t.Fatalf("relationship outcome evidence count for %q = %d, want 1", want.SourceSiteID, matches)
	}
}

func requireP6C3DiagnosticOutcome(t *testing.T, diagnostic graphhealth.Diagnostic, want ResolutionOutcome) {
	t.Helper()
	var got ResolutionOutcome
	if err := json.Unmarshal([]byte(diagnostic.Note), &got); err != nil {
		t.Fatalf("decode diagnostic outcome: %v", err)
	}
	if !reflect.DeepEqual(got, want) || diagnostic.SourceSiteID != want.SourceSiteID {
		t.Fatalf("diagnostic outcome drift:\nwant=%#v\ngot=%#v\ndiagnostic=%#v", want, got, diagnostic)
	}
}

func p6c3DiagnosticReason(t *testing.T, diagnostic graphhealth.Diagnostic) string {
	t.Helper()
	var outcome ResolutionOutcome
	if err := json.Unmarshal([]byte(diagnostic.Note), &outcome); err != nil {
		t.Fatalf("decode diagnostic final outcome: %v", err)
	}
	return outcome.Reason
}

func p6c3DiagnosticsForSite(t *testing.T, g *graph.Graph, sourceSiteID string) []graphhealth.Diagnostic {
	t.Helper()
	var matches []graphhealth.Diagnostic
	for _, node := range g.Nodes {
		value, ok := node.Properties[graphhealth.DiagnosticPropertyKey]
		if !ok {
			continue
		}
		raw, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal diagnostics: %v", err)
		}
		var diagnostics []graphhealth.Diagnostic
		if err := json.Unmarshal(raw, &diagnostics); err != nil {
			t.Fatalf("decode diagnostics: %v", err)
		}
		for _, diagnostic := range diagnostics {
			if diagnostic.SourceSiteID == sourceSiteID {
				matches = append(matches, diagnostic)
			}
		}
	}
	return matches
}

func p6c3CarrierCount(g *graph.Graph, sourceSiteID string) int {
	count := 0
	for _, relationship := range g.Relationships {
		for _, siteID := range mergeSourceSiteIDs(relationship.SourceSiteIDs, nil, relationship.SourceSiteID) {
			if siteID == sourceSiteID {
				count++
			}
		}
	}
	return count
}

func requireP6C3NoResolvedUnresolvedOverlap(t *testing.T, g *graph.Graph) {
	t.Helper()
	resolved := make(map[string]struct{})
	for _, relationship := range g.Relationships {
		for _, siteID := range mergeSourceSiteIDs(relationship.SourceSiteIDs, nil, relationship.SourceSiteID) {
			resolved[siteID] = struct{}{}
		}
	}
	for _, node := range g.Nodes {
		value, ok := node.Properties[graphhealth.DiagnosticPropertyKey]
		if !ok {
			continue
		}
		raw, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("marshal diagnostics: %v", err)
		}
		var diagnostics []graphhealth.Diagnostic
		if err := json.Unmarshal(raw, &diagnostics); err != nil {
			t.Fatalf("decode diagnostics: %v", err)
		}
		for _, diagnostic := range diagnostics {
			if _, overlap := resolved[diagnostic.SourceSiteID]; overlap {
				t.Fatalf("source site %q has resolved and diagnostic carriers", diagnostic.SourceSiteID)
			}
		}
	}
}

var p6c3AllMeanings = []scopeir.ExportMeaning{
	scopeir.ExportMeaningValue,
	scopeir.ExportMeaningType,
	scopeir.ExportMeaningNamespace,
}

func p6c3Definition(filePath string, name string, label scopeir.NodeLabel) scopeir.DefinitionFact {
	return scopeir.DefinitionFact{
		ID:            "def:" + filePath + ":" + string(label) + ":" + name,
		FilePath:      filePath,
		Name:          name,
		QualifiedName: name,
		Label:         label,
		Range:         scopeir.Range{StartLine: 1, EndLine: 1},
	}
}

func p6c3LocalExport(def scopeir.DefinitionFact, kind scopeir.ExportKind, exportedName string, meanings []scopeir.ExportMeaning) scopeir.ExportFact {
	return scopeir.ExportFact{
		FilePath:     def.FilePath,
		Kind:         kind,
		ExportedName: exportedName,
		LocalName:    def.Name,
		LocalDefID:   def.ID,
		Meanings:     append([]scopeir.ExportMeaning(nil), meanings...),
		TypeOnly:     len(meanings) == 1 && meanings[0] == scopeir.ExportMeaningType,
		Range:        def.Range,
	}
}

func p6c3Namespace(filePath string, exportedName string, targetRaw string, meanings []scopeir.ExportMeaning, typeOnly bool) scopeir.ExportFact {
	raw := targetRaw
	return scopeir.ExportFact{
		FilePath:     filePath,
		Kind:         scopeir.ExportNamespace,
		ExportedName: exportedName,
		TargetRaw:    &raw,
		Meanings:     append([]scopeir.ExportMeaning(nil), meanings...),
		TypeOnly:     typeOnly,
	}
}

func p6c3Import(kind scopeir.ImportKind, localName string, importedName string, targetRaw string, meanings []scopeir.ExportMeaning, typeOnly bool) scopeir.ImportFact {
	raw := targetRaw
	return scopeir.ImportFact{
		Kind:              kind,
		LocalName:         localName,
		ImportedName:      importedName,
		RequestedMeanings: append([]scopeir.ExportMeaning(nil), meanings...),
		TypeOnly:          typeOnly,
		TargetRaw:         &raw,
	}
}

func p6c3Module(filePath string, definitions []scopeir.DefinitionFact, exports []scopeir.ExportFact) scopeir.ScopeIR {
	moduleScope := p6c3ModuleScope(filePath)
	owned := make([]string, 0, len(definitions))
	bindings := make([]scopeir.BindingFact, 0, len(definitions))
	for index := range definitions {
		definitions[index].FilePath = filePath
		owned = append(owned, definitions[index].ID)
		bindings = append(bindings, scopeir.BindingFact{Name: definitions[index].Name, DefID: definitions[index].ID, Origin: scopeir.BindingLocal})
	}
	imports := make([]scopeir.ImportFact, 0, len(exports))
	for index := range exports {
		exports[index].FilePath = filePath
		fact := exports[index]
		if fact.TargetRaw == nil {
			continue
		}
		compatibility := scopeir.ImportFact{FilePath: filePath, TargetRaw: fact.TargetRaw}
		if fact.Kind == scopeir.ExportNamespace {
			compatibility.Kind = scopeir.ImportWildcard
			imports = append(imports, compatibility)
		}
	}
	return scopeir.ScopeIR{
		FilePath:    filePath,
		Language:    scanner.TypeScript,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{{
			ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: filePath,
			Range: scopeir.Range{StartLine: 1, EndLine: 20}, OwnedDefIDs: owned, Bindings: bindings,
		}},
		Definitions: definitions,
		Exports:     exports,
		Imports:     imports,
	}
}

func p6c3Consumer(filePath string, language scanner.Language, imports []scopeir.ImportFact, calls []scopeir.CallSiteFact) scopeir.ScopeIR {
	moduleScope := p6c3ModuleScope(filePath)
	functionScope := p6c3FunctionScope(filePath)
	caller := p6c3Definition(filePath, "caller", scopeir.NodeFunction)
	caller.Range = scopeir.Range{StartLine: 2, EndLine: 20}
	for index := range imports {
		imports[index].FilePath = filePath
	}
	for index := range calls {
		calls[index].FilePath = filePath
		calls[index].FileHash = "hash-" + filePath
		calls[index].InScope = functionScope
		calls[index].Range = scopeir.Range{StartLine: 3 + index, EndLine: 3 + index}
	}
	return scopeir.ScopeIR{
		FilePath:    filePath,
		Language:    language,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: filePath, Range: scopeir.Range{StartLine: 1, EndLine: 20}},
			{ID: functionScope, Parent: &moduleScope, Kind: scopeir.ScopeFunction, FilePath: filePath, Range: scopeir.Range{StartLine: 2, EndLine: 20}, OwnedDefIDs: []string{caller.ID}, Bindings: []scopeir.BindingFact{{Name: caller.Name, DefID: caller.ID, Origin: scopeir.BindingLocal}}},
		},
		Definitions: []scopeir.DefinitionFact{caller},
		Imports:     imports,
		Calls:       calls,
	}
}

func p6c3ModuleScope(filePath string) string {
	return "scope:" + filePath + ":module"
}

func p6c3FunctionScope(filePath string) string {
	return "scope:" + filePath + ":caller"
}

func p6c3EvidenceOfKind(values []graph.Evidence, kind string) []graph.Evidence {
	out := make([]graph.Evidence, 0)
	for _, evidence := range values {
		if evidence.Kind == kind {
			out = append(out, evidence)
		}
	}
	return out
}
