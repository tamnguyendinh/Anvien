package resolution

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

var p5cAllMeanings = []scopeir.ExportMeaning{
	scopeir.ExportMeaningValue,
	scopeir.ExportMeaningType,
	scopeir.ExportMeaningNamespace,
}

func TestP5CResolveAliasChainMatchesDirectTerminalAndRetainsProof(t *testing.T) {
	run := p5cDefinition("src/impl.ts", "run", scopeir.NodeFunction)
	parameterCount := 1
	run.ParameterCount = &parameterCount
	run.ParameterTypes = []string{"input"}
	run.Annotations = []string{"source-owned"}
	impl := p5cModule("src/impl.ts", []scopeir.DefinitionFact{run}, []scopeir.ExportFact{
		p5cLocalExport(run, scopeir.ExportDirect, "run", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	mid := p5cModule("src/mid.ts", nil, []scopeir.ExportFact{
		p5cReexport("src/mid.ts", "midRun", "run", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	barrel := p5cModule("src/barrel.ts", nil, []scopeir.ExportFact{
		p5cReexport("src/barrel.ts", "publicRun", "midRun", "./mid", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	barrelImport := p5cImport(scopeir.ImportAlias, "barrelRun", "publicRun", "./barrel", p5cAllMeanings, false)
	directImport := p5cImport(scopeir.ImportAlias, "directRun", "run", "./impl", p5cAllMeanings, false)
	consumer := p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{barrelImport, directImport}, []scopeir.CallSiteFact{
		{Name: "barrelRun", CallForm: scopeir.CallFree},
		{Name: "directRun", CallForm: scopeir.CallFree},
	})
	files := []scopeir.ScopeIR{consumer, barrel, mid, impl}

	w := p5cWorkspace(t, files)
	barrelResolved := p5cResolvedImportByLocalName(t, w, "src/app.ts", "barrelRun")
	directResolved := p5cResolvedImportByLocalName(t, w, "src/app.ts", "directRun")
	if barrelResolved.TargetDef == nil || directResolved.TargetDef == nil {
		t.Fatalf("terminal import bindings missing: barrel=%#v direct=%#v", barrelResolved, directResolved)
	}
	if barrelResolved.TargetDef.Fact.ID != run.ID || directResolved.TargetDef.Fact.ID != run.ID {
		t.Fatalf("direct/barrel terminal drift: barrel=%q direct=%q want=%q", barrelResolved.TargetDef.Fact.ID, directResolved.TargetDef.Fact.ID, run.ID)
	}

	result := w.resolveImportExport(barrelResolved.Fact, barrelResolved.TargetFiles, nil)
	if result.Outcome != exportResolutionTerminal || len(result.Candidates) != 1 {
		t.Fatalf("alias-chain result = %#v, want one terminal", result)
	}
	candidate := result.Candidates[0]
	if candidate.Terminal.Fact.ID != run.ID || len(candidate.Proofs) != 1 || len(candidate.Proofs[0].Hops) != 3 {
		t.Fatalf("alias-chain terminal/proof = %#v, want run with three export hops", candidate)
	}
	if result.Request.FilePath != "src/app.ts" || result.Request.LocalName != "barrelRun" || result.RequestedName != "publicRun" {
		t.Fatalf("consumer request provenance drifted: %#v", result)
	}
	candidate.Terminal.Fact.ParameterTypes[0] = "mutated"
	*candidate.Terminal.Fact.ParameterCount = 99
	candidate.Terminal.Fact.Annotations[0] = "mutated"
	workspaceDef := w.defsByID[run.ID]
	if workspaceDef.Fact.ParameterTypes[0] != "input" || *workspaceDef.Fact.ParameterCount != 1 || workspaceDef.Fact.Annotations[0] != "source-owned" {
		t.Fatalf("immutable result aliases workspace definition: %#v", workspaceDef.Fact)
	}

	resolved, err := Resolve(files, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	caller := requireNode(t, resolved.Graph, "Function", "src/app.ts", "caller")
	terminal := requireNode(t, resolved.Graph, "Function", "src/impl.ts", "run")
	if got := countRelationshipsBetween(resolved.Graph, graph.RelCalls, caller.ID, terminal.ID); got != 2 {
		t.Fatalf("direct/barrel CALLS to same terminal = %d, want 2", got)
	}
}

func TestP5CExplicitExportPrecedesStarAndStarExcludesDefault(t *testing.T) {
	winner := p5cDefinition("src/winner.ts", "pick", scopeir.NodeFunction)
	loser := p5cDefinition("src/loser.ts", "pick", scopeir.NodeFunction)
	defaultDef := p5cDefinition("src/loser.ts", "defaultPick", scopeir.NodeFunction)
	winnerModule := p5cModule("src/winner.ts", []scopeir.DefinitionFact{winner}, []scopeir.ExportFact{
		p5cLocalExport(winner, scopeir.ExportDirect, "pick", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	loserModule := p5cModule("src/loser.ts", []scopeir.DefinitionFact{loser, defaultDef}, []scopeir.ExportFact{
		p5cLocalExport(loser, scopeir.ExportDirect, "pick", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		p5cLocalExport(defaultDef, scopeir.ExportDefault, "default", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	barrel := p5cModule("src/barrel.ts", nil, []scopeir.ExportFact{
		p5cReexport("src/barrel.ts", "pick", "pick", "./winner", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		p5cStar("src/barrel.ts", "./loser", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	pickImport := p5cImport(scopeir.ImportNamed, "pick", "pick", "./barrel", p5cAllMeanings, false)
	defaultImport := p5cImport(scopeir.ImportNamed, "defaultPick", "default", "./barrel", p5cAllMeanings, false)
	consumer := p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{pickImport, defaultImport}, nil)
	w := p5cWorkspace(t, []scopeir.ScopeIR{consumer, barrel, winnerModule, loserModule})

	pick := p5cResolvedImportByLocalName(t, w, "src/app.ts", "pick")
	result := w.resolveImportExport(pick.Fact, pick.TargetFiles, nil)
	resolvedDef, ok := result.definition()
	if !ok || resolvedDef.Fact.ID != winner.ID {
		t.Fatalf("explicit-over-star result = %#v, want winner %q", result, winner.ID)
	}
	if len(result.Candidates[0].Proofs) != 1 || len(result.Candidates[0].Proofs[0].Hops) != 2 {
		t.Fatalf("explicit proof unexpectedly traversed star: %#v", result.Candidates[0])
	}

	missingDefault := p5cResolvedImportByLocalName(t, w, "src/app.ts", "defaultPick")
	defaultResult := w.resolveImportExport(missingDefault.Fact, missingDefault.TargetFiles, nil)
	if defaultResult.Outcome != exportResolutionMissing || len(defaultResult.Candidates) != 0 || missingDefault.TargetDef != nil {
		t.Fatalf("star synthesized default terminal: result=%#v import=%#v", defaultResult, missingDefault)
	}
}

func TestP5CSameTerminalDedupeDistinctAmbiguityAndCycles(t *testing.T) {
	t.Run("same terminal dedupe", func(t *testing.T) {
		terminal := p5cDefinition("src/impl.ts", "value", scopeir.NodeFunction)
		files := []scopeir.ScopeIR{
			p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
				p5cImport(scopeir.ImportNamed, "value", "value", "./root", p5cAllMeanings, false),
			}, nil),
			p5cModule("src/root.ts", nil, []scopeir.ExportFact{
				p5cStar("src/root.ts", "./left", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
				p5cStar("src/root.ts", "./right", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
			p5cModule("src/left.ts", nil, []scopeir.ExportFact{
				p5cReexport("src/left.ts", "value", "value", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
			p5cModule("src/right.ts", nil, []scopeir.ExportFact{
				p5cReexport("src/right.ts", "value", "value", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
			p5cModule("src/impl.ts", []scopeir.DefinitionFact{terminal}, []scopeir.ExportFact{
				p5cLocalExport(terminal, scopeir.ExportDirect, "value", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
		}
		w := p5cWorkspace(t, files)
		item := p5cResolvedImportByLocalName(t, w, "src/app.ts", "value")
		result := w.resolveImportExport(item.Fact, item.TargetFiles, nil)
		if result.Outcome != exportResolutionTerminal || len(result.Candidates) != 1 || len(result.Candidates[0].Proofs) != 2 {
			t.Fatalf("same-terminal result = %#v, want one candidate with two proofs", result)
		}
	})

	t.Run("distinct terminal ambiguity", func(t *testing.T) {
		left := p5cDefinition("src/left.ts", "value", scopeir.NodeFunction)
		right := p5cDefinition("src/right.ts", "value", scopeir.NodeFunction)
		files := []scopeir.ScopeIR{
			p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
				p5cImport(scopeir.ImportNamed, "value", "value", "./root", p5cAllMeanings, false),
			}, nil),
			p5cModule("src/root.ts", nil, []scopeir.ExportFact{
				p5cStar("src/root.ts", "./left", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
				p5cStar("src/root.ts", "./right", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
			p5cModule("src/left.ts", []scopeir.DefinitionFact{left}, []scopeir.ExportFact{
				p5cLocalExport(left, scopeir.ExportDirect, "value", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
			p5cModule("src/right.ts", []scopeir.DefinitionFact{right}, []scopeir.ExportFact{
				p5cLocalExport(right, scopeir.ExportDirect, "value", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
		}
		w := p5cWorkspace(t, files)
		item := p5cResolvedImportByLocalName(t, w, "src/app.ts", "value")
		result := w.resolveImportExport(item.Fact, item.TargetFiles, nil)
		if result.Outcome != exportResolutionAmbiguity || len(result.Candidates) != 2 || item.TargetDef != nil {
			t.Fatalf("distinct-terminal result selected a candidate: result=%#v import=%#v", result, item)
		}
		if _, selected := result.definition(); selected {
			t.Fatalf("ambiguous result returned a definition: %#v", result)
		}
	})

	t.Run("pure cycle", func(t *testing.T) {
		files := []scopeir.ScopeIR{
			p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
				p5cImport(scopeir.ImportNamed, "value", "value", "./a", p5cAllMeanings, false),
			}, nil),
			p5cModule("src/a.ts", nil, []scopeir.ExportFact{
				p5cStar("src/a.ts", "./b", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
			p5cModule("src/b.ts", nil, []scopeir.ExportFact{
				p5cStar("src/b.ts", "./a", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
		}
		w := p5cWorkspace(t, files)
		item := p5cResolvedImportByLocalName(t, w, "src/app.ts", "value")
		result := w.resolveImportExport(item.Fact, item.TargetFiles, nil)
		if result.Outcome != exportResolutionCycle || len(result.Candidates) != 0 || !p5cHasFailure(result, exportResolutionCycle) {
			t.Fatalf("pure-cycle result = %#v, want explicit cycle", result)
		}
	})

	t.Run("cycle with terminal branch", func(t *testing.T) {
		terminal := p5cDefinition("src/impl.ts", "value", scopeir.NodeFunction)
		files := []scopeir.ScopeIR{
			p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{
				p5cImport(scopeir.ImportNamed, "value", "value", "./a", p5cAllMeanings, false),
			}, nil),
			p5cModule("src/a.ts", nil, []scopeir.ExportFact{
				p5cStar("src/a.ts", "./b", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
				p5cStar("src/a.ts", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
			p5cModule("src/b.ts", nil, []scopeir.ExportFact{
				p5cStar("src/b.ts", "./a", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
			p5cModule("src/impl.ts", []scopeir.DefinitionFact{terminal}, []scopeir.ExportFact{
				p5cLocalExport(terminal, scopeir.ExportDirect, "value", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
			}),
		}
		w := p5cWorkspace(t, files)
		item := p5cResolvedImportByLocalName(t, w, "src/app.ts", "value")
		result := w.resolveImportExport(item.Fact, item.TargetFiles, nil)
		resolved, ok := result.definition()
		if !ok || resolved.Fact.ID != terminal.ID || !p5cHasFailure(result, exportResolutionCycle) {
			t.Fatalf("terminal-cycle result = %#v, want terminal plus retained cycle proof", result)
		}
	})
}

func TestP5CMeaningMismatchAndNamespaceMemberResolution(t *testing.T) {
	typeDef := p5cDefinition("src/types.ts", "OnlyType", scopeir.NodeInterface)
	typesModule := p5cModule("src/types.ts", []scopeir.DefinitionFact{typeDef}, []scopeir.ExportFact{
		p5cLocalExport(typeDef, scopeir.ExportDirect, "OnlyType", []scopeir.ExportMeaning{scopeir.ExportMeaningType}),
	})
	valueImport := p5cImport(scopeir.ImportNamed, "ValueOnlyType", "OnlyType", "./types", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}, false)
	typeImport := p5cImport(scopeir.ImportNamed, "TypeOnly", "OnlyType", "./types", []scopeir.ExportMeaning{scopeir.ExportMeaningType}, true)
	consumer := p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{valueImport, typeImport}, nil)
	w := p5cWorkspace(t, []scopeir.ScopeIR{consumer, typesModule})

	valueResolved := p5cResolvedImportByLocalName(t, w, "src/app.ts", "ValueOnlyType")
	valueResult := w.resolveImportExport(valueResolved.Fact, valueResolved.TargetFiles, nil)
	if valueResult.Outcome != exportResolutionMeaningMismatch || valueResolved.TargetDef != nil {
		t.Fatalf("value/type mismatch was repaired by name: result=%#v import=%#v", valueResult, valueResolved)
	}
	typeResolved := p5cResolvedImportByLocalName(t, w, "src/app.ts", "TypeOnly")
	if typeResolved.TargetDef == nil || typeResolved.TargetDef.Fact.ID != typeDef.ID {
		t.Fatalf("type lane did not reach terminal: %#v", typeResolved)
	}

	run := p5cDefinition("src/impl.ts", "run", scopeir.NodeFunction)
	impl := p5cModule("src/impl.ts", []scopeir.DefinitionFact{run}, []scopeir.ExportFact{
		p5cLocalExport(run, scopeir.ExportDirect, "run", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	barrel := p5cModule("src/barrel.ts", nil, []scopeir.ExportFact{
		p5cNamespace("src/barrel.ts", "api", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningNamespace}, false),
	})
	namedNamespace := p5cImport(scopeir.ImportNamed, "namedAPI", "api", "./barrel", p5cAllMeanings, false)
	directNamespace := p5cImport(scopeir.ImportNamespace, "directAPI", "", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningNamespace}, false)
	typeNamespace := p5cImport(scopeir.ImportNamespace, "typeAPI", "", "./impl", []scopeir.ExportMeaning{scopeir.ExportMeaningType}, true)
	namespaceConsumer := p5cConsumer("src/use.ts", scanner.TypeScript, []scopeir.ImportFact{namedNamespace, directNamespace, typeNamespace}, nil)
	namespaceWorkspace := p5cWorkspace(t, []scopeir.ScopeIR{namespaceConsumer, barrel, impl})
	functionScope := p5cFunctionScope("src/use.ts")

	namedTarget, namedOK := namespaceWorkspace.resolveImportedMember("namedAPI", "run", functionScope, callableLabels())
	directTarget, directOK := namespaceWorkspace.resolveImportedMember("directAPI", "run", functionScope, callableLabels())
	if !namedOK || !directOK || namedTarget.Fact.ID != run.ID || directTarget.Fact.ID != run.ID {
		t.Fatalf("namespace/direct member terminal drift: named=%#v/%t direct=%#v/%t", namedTarget, namedOK, directTarget, directOK)
	}
	memberResult, handled := namespaceWorkspace.resolveSemanticImportedMember("namedAPI", "run", functionScope, callableLabels())
	if !handled || memberResult.Outcome != exportResolutionTerminal || len(memberResult.Candidates) != 1 || len(memberResult.Candidates[0].Proofs[0].Hops) != 2 {
		t.Fatalf("namespace member proof = %#v handled=%t, want namespace+direct hops", memberResult, handled)
	}
	if _, ok := namespaceWorkspace.resolveImportedMember("typeAPI", "run", functionScope, callableLabels()); ok {
		t.Fatalf("type-only namespace resolved a value member")
	}
}

func TestP5CAmbiguousOwnerMemberPreservesNoSelectionAndEveryBranch(t *testing.T) {
	leftOwner := p5cDefinition("src/left.ts", "LeftAPI", scopeir.NodeClass)
	leftRun := p5cDefinition("src/left.ts", "run", scopeir.NodeMethod)
	leftRun.OwnerID = leftOwner.ID
	rightOwner := p5cDefinition("src/right.ts", "RightAPI", scopeir.NodeClass)

	left := p5cModule("src/left.ts", []scopeir.DefinitionFact{leftOwner, leftRun}, []scopeir.ExportFact{
		p5cLocalExport(leftOwner, scopeir.ExportDirect, "api", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	right := p5cModule("src/right.ts", []scopeir.DefinitionFact{rightOwner}, []scopeir.ExportFact{
		p5cLocalExport(rightOwner, scopeir.ExportDirect, "api", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	barrel := p5cModule("src/barrel.ts", nil, []scopeir.ExportFact{
		p5cStar("src/barrel.ts", "./left", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
		p5cStar("src/barrel.ts", "./right", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})
	apiImport := p5cImport(scopeir.ImportNamed, "api", "api", "./barrel", p5cAllMeanings, false)
	consumer := p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{apiImport}, []scopeir.CallSiteFact{
		{Name: "run", ExplicitReceiver: "api", CallForm: scopeir.CallMember},
	})
	consumer.Accesses = []scopeir.AccessFact{{
		FilePath:         "src/app.ts",
		FileHash:         "hash-src/app.ts",
		Name:             "run",
		Kind:             scopeir.AccessRead,
		Range:            scopeir.Range{StartLine: 4, EndLine: 4},
		InScope:          p5cFunctionScope("src/app.ts"),
		ExplicitReceiver: "api",
	}}
	files := []scopeir.ScopeIR{consumer, barrel, left, right}

	w := p5cWorkspace(t, files)
	memberResult, handled := w.resolveSemanticImportedMember(
		"api",
		"run",
		p5cFunctionScope("src/app.ts"),
		callableLabels(),
	)
	if !handled || memberResult.Outcome != exportResolutionAmbiguity {
		t.Fatalf("ambiguous owner member result = %#v handled=%t, want ambiguity", memberResult, handled)
	}
	if len(memberResult.Candidates) != 1 || memberResult.Candidates[0].Terminal.Fact.ID != leftRun.ID {
		t.Fatalf("member provenance candidates = %#v, want sole branch member retained without selection", memberResult.Candidates)
	}
	if _, selected := memberResult.definition(); selected {
		t.Fatalf("ambiguous owner member selected a definition: %#v", memberResult)
	}
	if _, selected := w.resolveImportedMember(
		"api",
		"run",
		p5cFunctionScope("src/app.ts"),
		callableLabels(),
	); selected {
		t.Fatalf("resolveImportedMember selected the sole member from an ambiguous owner")
	}

	branchOutcomes := make(map[string]exportResolutionOutcome)
	recordBranch := func(proof exportResolutionProof) {
		for _, hop := range proof.Hops {
			if hop.Kind == exportResolutionMemberHop && hop.MemberOwnerDefID != "" {
				branchOutcomes[hop.MemberOwnerDefID] = proof.Outcome
			}
		}
	}
	for _, candidate := range memberResult.Candidates {
		for _, proof := range candidate.Proofs {
			recordBranch(proof)
		}
	}
	for _, failure := range memberResult.Failures {
		recordBranch(failure)
	}
	if branchOutcomes[leftOwner.ID] != exportResolutionTerminal || branchOutcomes[rightOwner.ID] != exportResolutionMissing {
		t.Fatalf("owner branch provenance = %#v, want left terminal and right missing", branchOutcomes)
	}

	resolved, err := Resolve(files, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	caller := requireNode(t, resolved.Graph, "Function", "src/app.ts", "caller")
	soleMember := requireNode(t, resolved.Graph, "Method", "src/left.ts", "run")
	requireNoRelationship(t, resolved.Graph, graph.RelCalls, caller.ID, soleMember.ID)
	requireNoRelationship(t, resolved.Graph, graph.RelAccesses, caller.ID, soleMember.ID)
	if resolved.Metrics.ResolvedCalls != 0 || resolved.Metrics.ResolvedAccesses != 0 || resolved.Metrics.UnresolvedReferences != 2 {
		t.Fatalf("ambiguous member resolver metrics = %#v, want two unresolved source sites", resolved.Metrics)
	}
}

func TestP5CExplicitImportMissBlocksGlobalRescueAndPreservesImports(t *testing.T) {
	missingImport := p5cImport(scopeir.ImportNamed, "run", "run", "./barrel", p5cAllMeanings, false)
	consumer := p5cConsumer("src/app.ts", scanner.TypeScript, []scopeir.ImportFact{missingImport}, []scopeir.CallSiteFact{
		{Name: "run", CallForm: scopeir.CallFree},
	})
	barrel := p5cModule("src/barrel.ts", nil, nil)
	globalRun := p5cDefinition("src/global.ts", "run", scopeir.NodeFunction)
	globalModule := p5cModule("src/global.ts", []scopeir.DefinitionFact{globalRun}, []scopeir.ExportFact{
		p5cLocalExport(globalRun, scopeir.ExportDirect, "run", []scopeir.ExportMeaning{scopeir.ExportMeaningValue}),
	})

	result, err := Resolve([]scopeir.ScopeIR{consumer, barrel, globalModule}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	caller := requireNode(t, result.Graph, "Function", "src/app.ts", "caller")
	globalTarget := requireNode(t, result.Graph, "Function", "src/global.ts", "run")
	requireNoRelationship(t, result.Graph, graph.RelCalls, caller.ID, globalTarget.ID)
	diagnostics, ok := caller.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
	if !ok || len(diagnostics) != 1 {
		t.Fatalf("caller diagnostics = %#v, want one explicit-import miss", caller.Properties[graphhealth.DiagnosticPropertyKey])
	}
	if diagnostics[0].Note != "call target not resolved" || diagnostics[0].TargetText != "run" {
		t.Fatalf("explicit import reached global fallback: %#v", diagnostics[0])
	}
	if result.Metrics.ImportsResolved != 1 || result.Metrics.ResolvedCalls != 0 || result.Metrics.UnresolvedReferences != 1 {
		t.Fatalf("unexpected resolver metrics: %#v", result.Metrics)
	}
	if imports := result.Graph.RelationshipCountsByType()[graph.RelImports]; imports != 1 {
		t.Fatalf("syntactic IMPORTS count = %d, want unchanged 1", imports)
	}
}

func TestP5CPreservesNonSemanticGoPackageMemberResolution(t *testing.T) {
	targetRaw := "github.com/tamnguyendinh/anvien/internal/p5clegacy"
	legacyImport := p5cImport(scopeir.ImportNamed, "pkg", "pkg", targetRaw, nil, false)
	consumer := p5cConsumer("cmd/p5c/main.go", scanner.Go, []scopeir.ImportFact{legacyImport}, []scopeir.CallSiteFact{
		{Name: "Helper", ExplicitReceiver: "pkg", CallForm: scopeir.CallMember},
	})
	helper := p5cDefinition("internal/p5clegacy/helper.go", "Helper", scopeir.NodeFunction)
	target := p5cModuleWithLanguage("internal/p5clegacy/helper.go", scanner.Go, []scopeir.DefinitionFact{helper}, nil)

	result, err := Resolve([]scopeir.ScopeIR{consumer, target}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	caller := requireNode(t, result.Graph, "Function", "cmd/p5c/main.go", "caller")
	targetNode := requireNode(t, result.Graph, "Function", "internal/p5clegacy/helper.go", "Helper")
	requireRelationship(t, result.Graph, graph.RelCalls, caller.ID, targetNode.ID)
	if result.Metrics.ImportsResolved != 1 || result.Metrics.ResolvedCalls != 1 {
		t.Fatalf("legacy Go import regression: %#v", result.Metrics)
	}
}

func p5cWorkspace(t *testing.T, files []scopeir.ScopeIR) *workspace {
	t.Helper()
	w, err := buildWorkspace(files)
	if err != nil {
		t.Fatalf("buildWorkspace() error = %v", err)
	}
	t.Cleanup(func() { w.bindingAccumulator.dispose() })
	return w
}

func p5cDefinition(filePath string, name string, label scopeir.NodeLabel) scopeir.DefinitionFact {
	return scopeir.DefinitionFact{
		ID:            "def:" + filePath + ":" + string(label) + ":" + name,
		FilePath:      filePath,
		Name:          name,
		QualifiedName: name,
		Label:         label,
		Range:         scopeir.Range{StartLine: 1, EndLine: 1},
	}
}

func p5cLocalExport(def scopeir.DefinitionFact, kind scopeir.ExportKind, exportedName string, meanings []scopeir.ExportMeaning) scopeir.ExportFact {
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

func p5cReexport(filePath string, exportedName string, targetName string, targetRaw string, meanings []scopeir.ExportMeaning) scopeir.ExportFact {
	raw := targetRaw
	return scopeir.ExportFact{
		FilePath:           filePath,
		Kind:               scopeir.ExportReexport,
		ExportedName:       exportedName,
		TargetRaw:          &raw,
		TargetExportedName: targetName,
		Meanings:           append([]scopeir.ExportMeaning(nil), meanings...),
		TypeOnly:           len(meanings) == 1 && meanings[0] == scopeir.ExportMeaningType,
	}
}

func p5cStar(filePath string, targetRaw string, meanings []scopeir.ExportMeaning) scopeir.ExportFact {
	raw := targetRaw
	return scopeir.ExportFact{
		FilePath:  filePath,
		Kind:      scopeir.ExportStar,
		TargetRaw: &raw,
		Meanings:  append([]scopeir.ExportMeaning(nil), meanings...),
		TypeOnly:  len(meanings) == 1 && meanings[0] == scopeir.ExportMeaningType,
	}
}

func p5cNamespace(filePath string, exportedName string, targetRaw string, meanings []scopeir.ExportMeaning, typeOnly bool) scopeir.ExportFact {
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

func p5cImport(kind scopeir.ImportKind, localName string, importedName string, targetRaw string, meanings []scopeir.ExportMeaning, typeOnly bool) scopeir.ImportFact {
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

func p5cModule(filePath string, definitions []scopeir.DefinitionFact, exports []scopeir.ExportFact) scopeir.ScopeIR {
	return p5cModuleWithLanguage(filePath, scanner.TypeScript, definitions, exports)
}

func p5cModuleWithLanguage(filePath string, language scanner.Language, definitions []scopeir.DefinitionFact, exports []scopeir.ExportFact) scopeir.ScopeIR {
	moduleScope := p5cModuleScope(filePath)
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
		switch fact.Kind {
		case scopeir.ExportReexport:
			compatibility.Kind = scopeir.ImportReexport
			compatibility.LocalName = fact.ExportedName
			compatibility.ImportedName = fact.TargetExportedName
		case scopeir.ExportStar, scopeir.ExportNamespace:
			compatibility.Kind = scopeir.ImportWildcard
		default:
			continue
		}
		imports = append(imports, compatibility)
	}
	return scopeir.ScopeIR{
		FilePath:    filePath,
		Language:    language,
		ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{{
			ID:          moduleScope,
			Kind:        scopeir.ScopeModule,
			FilePath:    filePath,
			Range:       scopeir.Range{StartLine: 1, EndLine: 20},
			OwnedDefIDs: owned,
			Bindings:    bindings,
		}},
		Definitions: definitions,
		Exports:     exports,
		Imports:     imports,
	}
}

func p5cConsumer(filePath string, language scanner.Language, imports []scopeir.ImportFact, calls []scopeir.CallSiteFact) scopeir.ScopeIR {
	moduleScope := p5cModuleScope(filePath)
	functionScope := p5cFunctionScope(filePath)
	caller := p5cDefinition(filePath, "caller", scopeir.NodeFunction)
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

func p5cResolvedImportByLocalName(t *testing.T, w *workspace, filePath string, localName string) resolvedImport {
	t.Helper()
	filePath = cleanPath(filePath)
	for _, item := range w.imports {
		if cleanPath(item.Fact.FilePath) == filePath && item.Fact.LocalName == localName && isSemanticExportImport(item.Fact) {
			return item
		}
	}
	t.Fatalf("semantic import %s:%s not found in %#v", filePath, localName, w.imports)
	return resolvedImport{}
}

func p5cHasFailure(result exportResolutionResult, outcome exportResolutionOutcome) bool {
	for _, failure := range result.Failures {
		if failure.Outcome == outcome {
			return true
		}
	}
	return false
}

func p5cModuleScope(filePath string) string {
	return "scope:" + filePath + ":module"
}

func p5cFunctionScope(filePath string) string {
	return "scope:" + filePath + ":caller"
}

func countRelationshipsBetween(g *graph.Graph, relationshipType graph.RelationshipType, sourceID string, targetID string) int {
	count := 0
	for _, relationship := range g.Relationships {
		if relationship.Type == relationshipType && relationship.SourceID == sourceID && relationship.TargetID == targetID {
			count++
		}
	}
	return count
}
