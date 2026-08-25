package resolution

import (
	"errors"
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

func Resolve(files []scopeir.ScopeIR, options Options) (result Result, err error) {
	return ResolveInto(nil, files, options)
}

func ResolveInto(baseGraph *graph.Graph, files []scopeir.ScopeIR, options Options) (result Result, err error) {
	binding, err := BuildCrossFileBinding(files, options)
	if err != nil {
		return Result{}, err
	}
	return ResolveBoundInto(baseGraph, binding, options)
}

func BuildCrossFileBinding(files []scopeir.ScopeIR, options Options) (result BindingResult, err error) {
	w, err := buildWorkspace(files)
	if err != nil {
		return BindingResult{}, err
	}
	w.typeScriptStandardLibrary = options.TypeScriptStandardLibrary
	defer func() {
		if err != nil {
			w.bindingAccumulator.dispose()
		}
	}()
	metrics := Metrics{
		DefinitionsIndexed:    len(w.defsByID),
		ImportsResolved:       countResolvedImports(w.imports),
		HeritageFactsIndexed:  w.heritageFacts,
		UnresolvedInheritance: w.unresolvedHeritage,
		UnresolvedReferences:  w.unresolvedHeritage,
	}
	applyCrossFileCompatibilityMetrics(options, &metrics)
	metrics.BindingAccumulatorFiles = w.bindingAccumulator.fileCount()
	metrics.BindingAccumulatorEntries = w.bindingAccumulator.total()
	if err := w.bindingAccumulator.finalize(); err != nil {
		return BindingResult{Metrics: metrics}, err
	}
	metrics.BindingAccumulatorFinalized = w.bindingAccumulator.finalized
	return BindingResult{workspace: w, Metrics: metrics}, nil
}

func ResolveBoundInto(baseGraph *graph.Graph, binding BindingResult, options Options) (result Result, err error) {
	w := binding.workspace
	if w == nil {
		return Result{}, errors.New("resolution binding result is empty")
	}
	bindingOccurrences, err := buildBindingOccurrenceIndex(w)
	if err != nil {
		return Result{}, err
	}
	g := baseGraph
	if g == nil {
		g = graph.New()
	}
	metrics := binding.Metrics
	defer func() {
		w.bindingAccumulator.dispose()
		result.Metrics.BindingAccumulatorDisposed = w.bindingAccumulator.disposed
	}()
	e := newEmitter(g, &metrics, w.fileLanguages)

	if err := emitDefinitionNodes(w, e); err != nil {
		metrics.GraphNodesEmitted = len(g.Nodes)
		metrics.GraphRelationshipsEmitted = len(g.Relationships)
		return Result{Metrics: metrics}, err
	}
	emitUnresolvedHeritageDiagnostics(w, e)
	emitImportEdges(w, e)

	emitInherits := !options.DisableScopeInheritsCompatibility
	for _, item := range w.heritage {
		emitHeritageCompatibilityEdges(e, item, emitInherits)
		metrics.ResolvedInheritance++
	}

	for _, ir := range w.files {
		for _, call := range ir.Calls {
			resolveCall(w, e, bindingOccurrences, call)
		}
		for _, access := range ir.Accesses {
			resolveAccess(w, e, bindingOccurrences, access)
		}
		for _, annotation := range ir.TypeAnnotations {
			resolveTypeAnnotation(w, e, annotation)
		}
	}
	emitMethodDispatchEdges(w, e)
	authorityResults, err := finalizeTypeScriptAuthorityResults(w.typeScriptAuthorityResults, &metrics)
	if err != nil {
		return Result{Graph: g, ReferenceIndex: e.referenceIndex, Metrics: metrics}, err
	}
	if err := emitTypeScriptExternalSymbols(e, authorityResults); err != nil {
		return Result{Graph: g, ReferenceIndex: e.referenceIndex, Metrics: metrics}, err
	}
	outcomes, err := e.outcomes.finalize()
	if err != nil {
		return Result{Graph: g, ReferenceIndex: e.referenceIndex, Metrics: metrics}, err
	}
	if err := projectResolutionOutcomes(g, &e.referenceIndex, outcomes); err != nil {
		return Result{Graph: g, ReferenceIndex: e.referenceIndex, ResolutionOutcomes: outcomes, Metrics: metrics}, err
	}

	metrics.GraphNodesEmitted = len(g.Nodes)
	metrics.GraphRelationshipsEmitted = len(g.Relationships)
	graphhealth.SetResolutionMetadata(g, metrics.UnresolvedReferences, metrics.UnresolvedReferenceDiagnostics, metrics.UnattributedUnresolvedReferences)
	return Result{
		Graph:                      g,
		ReferenceIndex:             e.referenceIndex,
		TypeScriptAuthorityResults: authorityResults,
		ResolutionOutcomes:         outcomes,
		Metrics:                    metrics,
	}, nil
}

func applyCrossFileCompatibilityMetrics(options Options, metrics *Metrics) {
	metrics.CrossFileFilesReprocessed = 0
	metrics.CrossFileSkipped = true
	if options.SkipCompatibilityCrossFile {
		metrics.CrossFileSkipReason = "disabled-by-pipeline-option"
		return
	}
	metrics.CrossFileSkipReason = "covered-by-scopeir-single-pass-resolution"
}

const bindingOccurrenceTargetRole = "binding"

type bindingOccurrenceKey struct {
	filePath       string
	name           string
	rangeValue     scopeir.Range
	selectionRange scopeir.Range
	hasSelection   bool
}

type bindingOccurrenceOwnerScope struct {
	filePath string
	scope    scopeir.ScopeFact
}

type bindingOccurrenceIndex struct {
	defIDs map[string]struct{}
}

func buildBindingOccurrenceIndex(w *workspace) (bindingOccurrenceIndex, error) {
	index := bindingOccurrenceIndex{defIDs: make(map[string]struct{})}
	definitions := make(map[bindingOccurrenceKey][]defRef)
	ownedScopes := make(map[string][]bindingOccurrenceOwnerScope)

	for _, ir := range w.files {
		for _, def := range w.defsByFile[cleanPath(ir.FilePath)] {
			if def.Fact.Label != scopeir.NodeVariable {
				continue
			}
			key := bindingOccurrenceKeyForDefinition(def.Fact)
			definitions[key] = append(definitions[key], def)
		}
		for _, scope := range ir.Scopes {
			for _, defID := range scope.OwnedDefIDs {
				ownedScopes[defID] = append(ownedScopes[defID], bindingOccurrenceOwnerScope{
					filePath: cleanPath(ir.FilePath),
					scope:    scope,
				})
			}
		}
	}

	for _, ir := range w.files {
		for _, leaf := range ir.BindingLeaves {
			key := bindingOccurrenceKeyForLeaf(leaf)
			matches := definitions[key]
			if len(matches) != 1 {
				return bindingOccurrenceIndex{}, fmt.Errorf(
					"binding occurrence %q at %s:%d:%d has %d matching variable definitions, want exactly 1",
					leaf.Name,
					cleanPath(leaf.FilePath),
					leaf.Range.StartLine,
					leaf.Range.StartCol,
					len(matches),
				)
			}
			def := matches[0]
			owners := ownedScopes[def.Fact.ID]
			if len(owners) != 1 {
				return bindingOccurrenceIndex{}, fmt.Errorf(
					"binding occurrence %q definition %q has %d lexical owners, want exactly 1",
					leaf.Name,
					def.Fact.ID,
					len(owners),
				)
			}
			owner := owners[0]
			if err := validateBindingOccurrenceOwner(leaf, def.Fact, owner); err != nil {
				return bindingOccurrenceIndex{}, err
			}
			localBindingCount := 0
			for _, binding := range owner.scope.Bindings {
				if binding.Origin == scopeir.BindingLocal && binding.Name == leaf.Name && binding.DefID == def.Fact.ID {
					localBindingCount++
				}
			}
			if localBindingCount != 1 {
				return bindingOccurrenceIndex{}, fmt.Errorf(
					"binding occurrence %q definition %q has %d matching local bindings in scope %q, want exactly 1",
					leaf.Name,
					def.Fact.ID,
					localBindingCount,
					owner.scope.ID,
				)
			}
			if _, duplicate := index.defIDs[def.Fact.ID]; duplicate {
				return bindingOccurrenceIndex{}, fmt.Errorf(
					"binding occurrence %q definition %q is projected more than once",
					leaf.Name,
					def.Fact.ID,
				)
			}
			index.defIDs[def.Fact.ID] = struct{}{}
		}
	}

	return index, nil
}

func validateBindingOccurrenceOwner(leaf scopeir.BindingLeafFact, def scopeir.DefinitionFact, owner bindingOccurrenceOwnerScope) error {
	leafFilePath := cleanPath(leaf.FilePath)
	definitionFilePath := cleanPath(def.FilePath)
	ownerFilePath := cleanPath(owner.scope.FilePath)
	if leafFilePath == "" ||
		leafFilePath != definitionFilePath ||
		leafFilePath != owner.filePath ||
		leafFilePath != ownerFilePath {
		return fmt.Errorf(
			"binding occurrence %q definition %q lexical owner scope %q has file drift: leaf=%q definition=%q ownerIR=%q ownerScope=%q",
			leaf.Name,
			def.ID,
			owner.scope.ID,
			leafFilePath,
			definitionFilePath,
			owner.filePath,
			ownerFilePath,
		)
	}
	if !rangeContains(owner.scope.Range, leaf.Range) || !rangeContains(owner.scope.Range, def.Range) {
		return fmt.Errorf(
			"binding occurrence %q definition %q lexical owner scope %q does not contain the declaration range",
			leaf.Name,
			def.ID,
			owner.scope.ID,
		)
	}
	if leaf.SelectionRange != nil && !rangeContains(owner.scope.Range, *leaf.SelectionRange) {
		return fmt.Errorf(
			"binding occurrence %q definition %q lexical owner scope %q does not contain the leaf selection range",
			leaf.Name,
			def.ID,
			owner.scope.ID,
		)
	}
	if def.SelectionRange != nil && !rangeContains(owner.scope.Range, *def.SelectionRange) {
		return fmt.Errorf(
			"binding occurrence %q definition %q lexical owner scope %q does not contain the definition selection range",
			leaf.Name,
			def.ID,
			owner.scope.ID,
		)
	}
	return nil
}

func bindingOccurrenceKeyForDefinition(def scopeir.DefinitionFact) bindingOccurrenceKey {
	key := bindingOccurrenceKey{
		filePath:   cleanPath(def.FilePath),
		name:       def.Name,
		rangeValue: def.Range,
	}
	if def.SelectionRange != nil {
		key.selectionRange = *def.SelectionRange
		key.hasSelection = true
	}
	return key
}

func bindingOccurrenceKeyForLeaf(leaf scopeir.BindingLeafFact) bindingOccurrenceKey {
	key := bindingOccurrenceKey{
		filePath:   cleanPath(leaf.FilePath),
		name:       leaf.Name,
		rangeValue: leaf.Range,
	}
	if leaf.SelectionRange != nil {
		key.selectionRange = *leaf.SelectionRange
		key.hasSelection = true
	}
	return key
}

func (index bindingOccurrenceIndex) resolve(w *workspace, name string, startScope string) (defRef, bool) {
	name = strings.TrimSpace(name)
	if name == "" {
		return defRef{}, false
	}
	target, ok := w.resolveScopedName(name, startScope, []scopeir.NodeLabel{scopeir.NodeVariable})
	if !ok {
		return defRef{}, false
	}
	_, ok = index.defIDs[target.Fact.ID]
	return target, ok
}

func bindingOccurrenceReferenceSource(source defRef, target defRef, filePath string) defRef {
	if source.GraphID != target.GraphID {
		return source
	}
	if file, ok := callerFileRef(filePath); ok {
		return file
	}
	return source
}

func emitBindingOccurrenceReference(e *emitter, source defRef, target defRef, fromScope string, filePath string, fileHash string, factRange scopeir.Range, targetText string, kind ReferenceKind) {
	source = bindingOccurrenceReferenceSource(source, target, filePath)
	e.emitReference(source, target, Reference{
		FromScope:        fromScope,
		ToDefID:          target.Fact.ID,
		FilePath:         filePath,
		FileHash:         fileHash,
		Range:            factRange,
		Kind:             kind,
		Confidence:       1,
		SourceSiteID:     sourceSiteID("access", filePath, targetText, factRange),
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        proofKindScopeBinding,
		TargetRole:       bindingOccurrenceTargetRole,
		TargetText:       targetText,
		Evidence: []graph.Evidence{{
			Kind:   proofKindScopeBinding,
			Weight: 1,
			Note:   targetText,
		}},
	})
	e.metrics.ResolvedReferences++
	e.metrics.ResolvedAccesses++
}

func sourceForScopeOrFile(w *workspace, scopeID string, filePath string) (defRef, bool) {
	if source, ok := w.callerForScope(scopeID); ok {
		return source, true
	}
	return callerFileRef(filePath)
}

func emitUnresolvedHeritageDiagnostics(w *workspace, e *emitter) {
	for _, item := range w.unresolvedHeritageFacts {
		source, ok := w.ownerForScope(item.InScope, dispatchOwnerLabels())
		note := "heritage target not resolved"
		if !ok {
			source, ok = callerFileRef(item.FilePath)
			note = "heritage owner not resolved"
		}
		if baseTypeName(item.Name) == "" {
			note = "heritage target text not modeled"
		}
		if !ok {
			e.emitUnresolvedReference(defRef{}, "heritage", item.Name, item.FilePath, item.FileHash, item.Range, note, false)
			continue
		}
		e.emitUnresolvedReference(source, "heritage", item.Name, item.FilePath, item.FileHash, item.Range, note, false)
	}
}

func resolveCall(w *workspace, e *emitter, bindingOccurrences bindingOccurrenceIndex, call scopeir.CallSiteFact) {
	source, ok := sourceForScopeOrFile(w, call.InScope, call.FilePath)
	if !ok {
		e.emitUnresolvedReference(defRef{}, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, "source scope not resolved", true)
		return
	}
	if source.Fact.Label == scopeir.NodeFile {
		e.emitUnresolvedReference(source, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, unresolvedNoteCallSourceFileLevel, true)
		return
	}
	var target defRef
	confidence := 1.0
	proofKind := ""
	var semanticResult exportResolutionResult
	hasSemanticResult := false
	lowConfidenceFallback := false
	bindingReceiverResolved := false
	externalBlocked := false
	repositoryReceiverClaimed := false
	if call.CallForm == scopeir.CallMember {
		repositoryReceiverClaimed = w.repositoryReceiverClaimed(call.ExplicitReceiver, call.InScope)
		externalBlocked = repositoryReceiverClaimed
		if receiver, resolved := bindingOccurrences.resolve(w, call.ExplicitReceiver, call.InScope); resolved {
			emitBindingOccurrenceReference(
				e,
				source,
				receiver,
				call.InScope,
				call.FilePath,
				call.FileHash,
				call.Range,
				call.ExplicitReceiver,
				ReferenceRead,
			)
			bindingReceiverResolved = true
		}
	}
	switch call.CallForm {
	case scopeir.CallConstructor:
		target, ok = w.resolveScopedName(call.Name, call.InScope, dispatchOwnerLabels())
		if ok {
			if claimed, allowed := w.explicitImportCallState(call.Name, call.InScope, dispatchOwnerLabels(), &target); claimed && !allowed {
				ok = false
				externalBlocked = true
			} else {
				proofKind = proofKindScopeBinding
			}
		}
		if !ok {
			target, ok = w.resolveSameFileName(call.FilePath, call.Name, dispatchOwnerLabels())
			if ok {
				confidence = 0.95
				proofKind = proofKindSameFile
			}
		}
		if !ok {
			claimed, _ := w.explicitImportCallState(call.Name, call.InScope, dispatchOwnerLabels(), nil)
			if claimed {
				externalBlocked = true
				break
			}
			target, ok = w.resolveGlobalCallName(call.Name, dispatchOwnerLabels(), call.Arity)
			if ok {
				confidence = 0.5
				lowConfidenceFallback = true
			}
		}
	case scopeir.CallMember:
		target, ok = w.resolveMember(call.Name, call.ExplicitReceiver, call.InScope, callableLabels())
		if ok {
			proofKind = proofKindReceiverMember
		}
		if !ok {
			target, semanticResult, ok = w.resolveImportedMemberWithProof(call.ExplicitReceiver, call.Name, call.InScope, callableLabels())
			if ok {
				hasSemanticResult = semanticResult.Outcome != ""
				confidence = 0.9
				proofKind = proofKindImportMember
			}
		}
		if !ok && (semanticResult.Outcome != "" || w.explicitImportNameClaimed(call.ExplicitReceiver, call.InScope)) {
			externalBlocked = true
		}
		if !ok && call.ExplicitReceiver == "" {
			claimed, _ := w.explicitImportCallState(call.Name, call.InScope, callableLabels(), nil)
			if claimed {
				break
			}
			target, ok = w.resolveGlobalCallName(call.Name, callableLabels(), call.Arity)
			if ok {
				confidence = 0.5
				lowConfidenceFallback = true
			}
		}
	default:
		target, ok = w.resolveScopedName(call.Name, call.InScope, callableLabels())
		if ok {
			if claimed, allowed := w.explicitImportCallState(call.Name, call.InScope, callableLabels(), &target); claimed && !allowed {
				ok = false
				externalBlocked = true
			} else {
				proofKind = proofKindScopeBinding
			}
		}
		if !ok {
			target, ok = w.resolveSameFileName(call.FilePath, call.Name, callableLabels())
			if ok {
				confidence = 0.95
				proofKind = proofKindSameFile
			}
		}
		if !ok {
			target, ok = w.resolveGoSamePackageFunction(call.FilePath, call.Name, call.Arity)
			if ok {
				confidence = 0.95
				proofKind = proofKindGoSamePackage
			}
		}
		if !ok {
			claimed, _ := w.explicitImportCallState(call.Name, call.InScope, callableLabels(), nil)
			if claimed {
				externalBlocked = true
				break
			}
			target, ok = w.resolveGlobalCallName(call.Name, callableLabels(), call.Arity)
			if ok {
				confidence = 0.5
				lowConfidenceFallback = true
			}
		}
	}
	if (!ok || lowConfidenceFallback) && !externalBlocked {
		lookup := tsstdlib.LookupResult{Status: tsstdlib.LookupNotFound}
		if call.CallForm == scopeir.CallMember && call.ExplicitReceiver != "" {
			lookup = w.lookupTypeScriptMember(call.FilePath, call.ExplicitReceiver, call.InScope, call.Name, tsstdlib.MeaningValue)
		} else {
			lookup = w.lookupTypeScriptGlobal(call.FilePath, call.Name, tsstdlib.MeaningValue)
		}
		if recordTypeScriptLookup(w, e, typeScriptSourceSite{
			kind:          "call",
			filePath:      call.FilePath,
			fileHash:      call.FileHash,
			rangeValue:    call.Range,
			targetText:    callTargetText(call),
			sourceGraphID: source.GraphID,
			fromScope:     call.InScope,
			referenceKind: ReferenceCall,
		}, lookup) {
			return
		}
	}
	if !ok {
		siteID := sourceSiteID("call", call.FilePath, callTargetText(call), call.Range)
		var outcomeEvidence []graph.Evidence
		if semanticResult.Outcome != "" {
			outcomeEvidence = appendExportBindingEvidence(nil, semanticResult, siteID)
		}
		if bindingReceiverResolved && w.typeScriptStandardLibrary == nil {
			outcomeEvidence = append(outcomeEvidence, graph.Evidence{
				Kind:   proofKindScopeBinding,
				Weight: 1,
				Note:   call.ExplicitReceiver,
			})
			e.recordRepositoryUnresolvedOutcome(
				"call",
				callTargetText(call),
				call.FilePath,
				call.FileHash,
				call.Range,
				"repository receiver resolved but call member not found",
				proofKindReceiverMember,
				outcomeEvidence,
			)
			return
		}
		e.emitUnresolvedReference(source, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, "call target not resolved", true, outcomeEvidence...)
		return
	}
	if lowConfidenceFallback {
		e.emitUnresolvedReference(source, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, "call target matched low-confidence global fallback only", true)
		return
	}
	siteID := sourceSiteID("call", call.FilePath, callTargetText(call), call.Range)
	evidence := []graph.Evidence{{
		Kind:   callEvidenceKind(call.CallForm),
		Weight: 1,
		Note:   call.Name,
	}}
	if !hasSemanticResult && proofKind == proofKindScopeBinding {
		bindingLabels := callableLabels()
		if call.CallForm == scopeir.CallConstructor {
			bindingLabels = dispatchOwnerLabels()
		}
		semanticResult, hasSemanticResult = w.retainedExportResolutionForScopedBinding(
			call.Name,
			call.InScope,
			bindingLabels,
			target,
		)
	}
	if hasSemanticResult {
		evidence = appendExportBindingEvidence(evidence, semanticResult, siteID)
	}
	e.emitReference(source, target, Reference{
		FromScope:        call.InScope,
		ToDefID:          target.Fact.ID,
		FilePath:         call.FilePath,
		FileHash:         call.FileHash,
		Range:            call.Range,
		Kind:             ReferenceCall,
		Confidence:       confidence,
		SourceSiteID:     siteID,
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        proofKind,
		TargetRole:       targetRoleCallable,
		TargetText:       callTargetText(call),
		Evidence:         evidence,
	})
	e.metrics.ResolvedReferences++
	e.metrics.ResolvedCalls++
}

func callerFileRef(filePath string) (defRef, bool) {
	filePath = cleanPath(filePath)
	if filePath == "" {
		return defRef{}, false
	}
	return defRef{
		Fact: scopeir.DefinitionFact{
			ID:       graph.GenerateID(string(scopeir.NodeFile), filePath),
			FilePath: filePath,
			Name:     filePath,
			Label:    scopeir.NodeFile,
		},
		GraphID: graph.GenerateID(string(scopeir.NodeFile), filePath),
	}, true
}

func callTargetText(call scopeir.CallSiteFact) string {
	if call.ExplicitReceiver != "" {
		return call.ExplicitReceiver + "." + call.Name
	}
	return call.Name
}

func resolveAccess(w *workspace, e *emitter, bindingOccurrences bindingOccurrenceIndex, access scopeir.AccessFact) {
	source, ok := sourceForScopeOrFile(w, access.InScope, access.FilePath)
	if !ok {
		e.emitUnresolvedReference(defRef{}, "access", accessTargetText(access), access.FilePath, access.FileHash, access.Range, "source scope not resolved", true)
		return
	}
	var target defRef
	confidence := 1.0
	evidenceKind := "type-binding"
	proofKind := proofKindReceiverMember
	targetRole := targetRoleMember
	var semanticResult exportResolutionResult
	hasSemanticResult := false
	externalBlocked := false
	referenceKind := ReferenceRead
	if access.Kind == scopeir.AccessWrite {
		referenceKind = ReferenceWrite
	}
	if access.ExplicitReceiver == "" {
		target, ok = bindingOccurrences.resolve(w, access.Name, access.InScope)
		if ok {
			evidenceKind = proofKindScopeBinding
			proofKind = proofKindScopeBinding
			targetRole = bindingOccurrenceTargetRole
		}
		if !ok && w.explicitImportNameClaimed(access.Name, access.InScope) {
			externalBlocked = true
		}
	} else {
		if w.repositoryReceiverClaimed(access.ExplicitReceiver, access.InScope) {
			externalBlocked = true
		}
		target, ok = w.resolveMember(access.Name, access.ExplicitReceiver, access.InScope, propertyLabels())
		if !ok {
			target, semanticResult, ok = w.resolveImportedMemberWithProof(access.ExplicitReceiver, access.Name, access.InScope, propertyLabels())
			if ok {
				hasSemanticResult = semanticResult.Outcome != ""
				confidence = 0.9
				evidenceKind = "import-binding"
				proofKind = proofKindImportMember
			}
		}
		if !ok && (semanticResult.Outcome != "" || (!w.hasTypeBinding(access.ExplicitReceiver, access.InScope) && w.explicitImportNameClaimed(access.ExplicitReceiver, access.InScope))) {
			externalBlocked = true
		}
	}
	if !ok {
		siteID := sourceSiteID("access", access.FilePath, accessTargetText(access), access.Range)
		var outcomeEvidence []graph.Evidence
		if semanticResult.Outcome != "" {
			outcomeEvidence = appendExportBindingEvidence(nil, semanticResult, siteID)
		}
		if !externalBlocked {
			lookup := tsstdlib.LookupResult{Status: tsstdlib.LookupNotFound}
			if access.ExplicitReceiver == "" {
				lookup = w.lookupTypeScriptGlobal(access.FilePath, access.Name, tsstdlib.MeaningValue)
			} else {
				lookup = w.lookupTypeScriptMember(access.FilePath, access.ExplicitReceiver, access.InScope, access.Name, tsstdlib.MeaningValue)
			}
			if recordTypeScriptLookup(w, e, typeScriptSourceSite{
				kind:          "access",
				filePath:      access.FilePath,
				fileHash:      access.FileHash,
				rangeValue:    access.Range,
				targetText:    accessTargetText(access),
				sourceGraphID: source.GraphID,
				fromScope:     access.InScope,
				referenceKind: referenceKind,
			}, lookup) {
				return
			}
		}
		e.emitUnresolvedReference(source, "access", accessTargetText(access), access.FilePath, access.FileHash, access.Range, "access target not resolved", true, outcomeEvidence...)
		return
	}
	kind := referenceKind
	source = bindingOccurrenceReferenceSource(source, target, access.FilePath)
	siteID := sourceSiteID("access", access.FilePath, accessTargetText(access), access.Range)
	evidence := []graph.Evidence{{
		Kind:   evidenceKind,
		Weight: 1,
		Note:   access.ExplicitReceiver + "." + access.Name,
	}}
	if !hasSemanticResult && proofKind == proofKindScopeBinding {
		semanticResult, hasSemanticResult = w.retainedExportResolutionForScopedBinding(
			access.Name,
			access.InScope,
			propertyLabels(),
			target,
		)
	}
	if hasSemanticResult {
		evidence = appendExportBindingEvidence(evidence, semanticResult, siteID)
	}
	e.emitReference(source, target, Reference{
		FromScope:        access.InScope,
		ToDefID:          target.Fact.ID,
		FilePath:         access.FilePath,
		FileHash:         access.FileHash,
		Range:            access.Range,
		Kind:             kind,
		Confidence:       confidence,
		SourceSiteID:     siteID,
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        proofKind,
		TargetRole:       targetRole,
		TargetText:       accessTargetText(access),
		Evidence:         evidence,
	})
	e.metrics.ResolvedReferences++
	e.metrics.ResolvedAccesses++
}

func accessTargetText(access scopeir.AccessFact) string {
	if access.ExplicitReceiver != "" {
		return access.ExplicitReceiver + "." + access.Name
	}
	return access.Name
}

func resolveTypeAnnotation(w *workspace, e *emitter, annotation scopeir.TypeAnnotationFact) {
	targetName := baseTypeName(annotation.Type.RawName)
	if targetName == "" {
		return
	}
	source, ok := sourceForScopeOrFile(w, annotation.InScope, annotation.FilePath)
	if !ok {
		e.emitUnresolvedReference(defRef{}, "type-reference", annotation.Type.RawName, annotation.FilePath, annotation.FileHash, annotation.Range, "source scope not resolved", true)
		return
	}
	target, ok := w.resolveName(targetName, annotation.InScope, typeLabels())
	if !ok && w.typeScriptAnnotationImportClaimed(annotation, targetName) {
		e.emitUnresolvedReference(source, "type-reference", annotation.Type.RawName, annotation.FilePath, annotation.FileHash, annotation.Range, "type target not resolved", true)
		return
	}
	if !ok && isBuiltinType(targetName) {
		e.recordIntrinsicTypeOutcome(annotation, targetName)
		return
	}
	if !ok {
		lookup := w.lookupTypeScriptAnnotation(annotation, targetName)
		if recordTypeScriptLookup(w, e, typeScriptSourceSite{
			kind:          "type-reference",
			filePath:      annotation.FilePath,
			fileHash:      annotation.FileHash,
			rangeValue:    annotation.Range,
			targetText:    annotation.Type.RawName,
			sourceGraphID: source.GraphID,
			fromScope:     annotation.InScope,
			referenceKind: ReferenceTypeReference,
		}, lookup) {
			return
		}
		e.emitUnresolvedReference(source, "type-reference", annotation.Type.RawName, annotation.FilePath, annotation.FileHash, annotation.Range, "type target not resolved", true)
		return
	}
	e.emitReference(source, target, Reference{
		FromScope:        annotation.InScope,
		ToDefID:          target.Fact.ID,
		FilePath:         annotation.FilePath,
		FileHash:         annotation.FileHash,
		Range:            annotation.Range,
		Kind:             ReferenceTypeReference,
		Confidence:       1,
		SourceSiteID:     sourceSiteID("type-reference", annotation.FilePath, annotation.Type.RawName, annotation.Range),
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        proofKindScopeBinding,
		TargetRole:       targetRoleType,
		TargetText:       annotation.Type.RawName,
		Evidence: []graph.Evidence{{
			Kind:   "scope-chain",
			Weight: 1,
			Note:   annotation.Type.RawName,
		}},
	})
	e.metrics.ResolvedReferences++
	e.metrics.ResolvedTypeReferences++
}

func (w *workspace) lookupTypeScriptAnnotation(annotation scopeir.TypeAnnotationFact, targetName string) tsstdlib.LookupResult {
	switch annotation.Type.Source {
	case scopeir.TypeSourceMethodReturn, scopeir.TypeSourceFieldAccess:
		owner, member, ok := splitQualifiedDeclarationName(annotation.Type.RawName)
		if !ok {
			return tsstdlib.LookupResult{Status: tsstdlib.LookupNotFound}
		}
		return w.lookupTypeScriptMember(annotation.FilePath, owner, annotation.InScope, member, tsstdlib.MeaningValue)
	case scopeir.TypeSourceConstructor, scopeir.TypeSourceCallReturn:
		return w.lookupTypeScriptGlobal(annotation.FilePath, targetName, tsstdlib.MeaningValue)
	default:
		return w.lookupTypeScriptGlobal(annotation.FilePath, targetName, tsstdlib.MeaningType)
	}
}

func (w *workspace) typeScriptAnnotationImportClaimed(annotation scopeir.TypeAnnotationFact, targetName string) bool {
	if w.explicitImportNameClaimed(targetName, annotation.InScope) {
		return true
	}
	switch annotation.Type.Source {
	case scopeir.TypeSourceMethodReturn, scopeir.TypeSourceFieldAccess:
		owner, _, ok := splitQualifiedDeclarationName(annotation.Type.RawName)
		return ok && w.explicitImportNameClaimed(owner, annotation.InScope)
	default:
		return false
	}
}

func splitQualifiedDeclarationName(value string) (string, string, bool) {
	value = strings.TrimSpace(value)
	index := strings.LastIndex(value, ".")
	if index <= 0 || index == len(value)-1 {
		return "", "", false
	}
	owner := strings.TrimSpace(value[:index])
	member := baseTypeName(value[index+1:])
	if owner == "" || member == "" {
		return "", "", false
	}
	return owner, member, true
}

func (w *workspace) lookupTypeScriptGlobal(filePath string, name string, meaning tsstdlib.Meaning) tsstdlib.LookupResult {
	if w.typeScriptStandardLibrary == nil || w.fileLanguages[cleanPath(filePath)] != scanner.TypeScript {
		return tsstdlib.LookupResult{Status: tsstdlib.LookupNotFound}
	}
	return w.typeScriptStandardLibrary.LookupGlobal(name, meaning)
}

func (w *workspace) lookupTypeScriptMember(filePath string, receiver string, startScope string, memberName string, memberMeaning tsstdlib.Meaning) tsstdlib.LookupResult {
	if w.typeScriptStandardLibrary == nil || w.fileLanguages[cleanPath(filePath)] != scanner.TypeScript {
		return tsstdlib.LookupResult{Status: tsstdlib.LookupNotFound}
	}
	if w.repositoryReceiverClaimed(receiver, startScope) {
		return tsstdlib.LookupResult{Status: tsstdlib.LookupNotFound}
	}
	ownerName := strings.TrimSpace(receiver)
	if ownerName == "" {
		return tsstdlib.LookupResult{Status: tsstdlib.LookupNotFound}
	}
	return w.typeScriptStandardLibrary.LookupMember(ownerName, tsstdlib.MeaningValue, memberName, memberMeaning)
}

func (w *workspace) repositoryReceiverClaimed(receiver string, startScope string) bool {
	receiver = strings.TrimSpace(receiver)
	if receiver == "" {
		return false
	}
	root := strings.TrimSpace(strings.Split(receiver, ".")[0])
	if root == "" {
		return false
	}
	if root == "this" || root == "self" || w.explicitImportNameClaimed(root, startScope) {
		return true
	}
	for scopeID := startScope; scopeID != ""; scopeID = w.parentScope(scopeID) {
		if len(w.scopeBindings[scopeID][root]) > 0 || len(w.typeBindings[scopeID][root]) > 0 {
			return true
		}
	}
	claimed := false
	forEachGlobalLookupName(root, func(name string) bool {
		if len(w.defsByName[name]) > 0 {
			claimed = true
			return false
		}
		return true
	})
	return claimed
}

func (w *workspace) explicitImportNameClaimed(name string, startScope string) bool {
	name = strings.TrimSpace(name)
	sourceFile := w.scopeFilePath(startScope)
	if name == "" || sourceFile == "" {
		return false
	}
	key := importReceiverKey{filePath: sourceFile, localName: name}
	return len(w.importClaimsByReceiver[key]) > 0
}

func (w *workspace) hasTypeBinding(name string, startScope string) bool {
	_, ok := w.lookupTypeBinding(strings.TrimSpace(name), startScope)
	return ok
}

type typeScriptSourceSite struct {
	kind          string
	filePath      string
	fileHash      string
	rangeValue    scopeir.Range
	targetText    string
	sourceGraphID string
	fromScope     string
	referenceKind ReferenceKind
}

func recordTypeScriptLookup(w *workspace, e *emitter, site typeScriptSourceSite, result tsstdlib.LookupResult) bool {
	switch result.Status {
	case tsstdlib.LookupResolved:
		e.metrics.ResolvedExternalDeclarations++
	case tsstdlib.LookupCapabilityUnavailable:
		e.metrics.ExternalCapabilityUnavailable++
	case tsstdlib.LookupProfileExcluded:
		e.metrics.ExternalProfileExcluded++
	case tsstdlib.LookupMeaningMismatch:
		e.metrics.ExternalMeaningMismatches++
	default:
		return false
	}
	siteID := sourceSiteID(site.kind, site.filePath, site.targetText, site.rangeValue)
	record := TypeScriptAuthorityResult{
		SourceSiteID:        siteID,
		Stage:               TypeScriptStandardLibraryStage,
		FilePath:            cleanPath(site.filePath),
		FileHash:            site.fileHash,
		Range:               site.rangeValue,
		SiteKind:            site.kind,
		RequestedName:       result.Name,
		RequestedMeaning:    result.Meaning,
		Status:              result.Status,
		Reason:              result.Reason,
		ResolvedSymbolID:    result.SymbolID,
		ResolvedOwnerID:     result.OwnerID,
		DeclarationRanges:   append([]tsstdlib.Declaration(nil), result.Declarations...),
		AuthorityKind:       result.AuthorityKind,
		CatalogProofState:   result.CatalogProofState,
		AuthorityHash:       result.AuthorityHash,
		TypeScriptVersion:   result.TypeScriptVersion,
		CatalogHash:         result.CatalogHash,
		CatalogArtifactHash: result.CatalogArtifactHash,
		ProfileHash:         result.ProfileHash,
		ConfigHash:          result.ConfigHash,
	}
	w.typeScriptAuthorityResults = append(w.typeScriptAuthorityResults, record)
	e.typeScriptExternalSites = append(e.typeScriptExternalSites, typeScriptExternalSite{
		SourceSiteID:  siteID,
		SourceGraphID: site.sourceGraphID,
		FromScope:     site.fromScope,
		ReferenceKind: site.referenceKind,
		TargetRole:    targetRoleForFactFamily(site.kind),
		TargetText:    site.targetText,
	})
	outcome, encoded, added := e.recordTypeScriptOutcome(site, record)
	if added && outcome.Status != ResolutionResolvedExternal {
		e.emitTypeScriptOutcomeDiagnostic(site, outcome, encoded)
	}
	return true
}

func finalizeTypeScriptAuthorityResults(records []TypeScriptAuthorityResult, metrics *Metrics) ([]TypeScriptAuthorityResult, error) {
	out := make([]TypeScriptAuthorityResult, len(records))
	for index, record := range records {
		out[index] = cloneTypeScriptAuthorityResult(record)
	}
	sort.Slice(out, func(left, right int) bool {
		return out[left].SourceSiteID < out[right].SourceSiteID
	})
	seen := make(map[string]TypeScriptAuthorityResult, len(out))
	unique := out[:0]
	counts := map[tsstdlib.LookupStatus]int{}
	for _, record := range out {
		if err := validateTypeScriptAuthorityResult(record); err != nil {
			return nil, err
		}
		if previous, duplicate := seen[record.SourceSiteID]; duplicate {
			if !reflect.DeepEqual(previous, record) {
				return nil, fmt.Errorf("conflicting TypeScript authority source site %q", record.SourceSiteID)
			}
			continue
		}
		seen[record.SourceSiteID] = record
		unique = append(unique, record)
		counts[record.Status]++
	}
	metrics.ResolvedExternalDeclarations = counts[tsstdlib.LookupResolved]
	metrics.ExternalCapabilityUnavailable = counts[tsstdlib.LookupCapabilityUnavailable]
	metrics.ExternalProfileExcluded = counts[tsstdlib.LookupProfileExcluded]
	metrics.ExternalMeaningMismatches = counts[tsstdlib.LookupMeaningMismatch]
	if len(unique) != metrics.ResolvedExternalDeclarations+
		metrics.ExternalCapabilityUnavailable+
		metrics.ExternalProfileExcluded+
		metrics.ExternalMeaningMismatches {
		return nil, fmt.Errorf(
			"TypeScript authority site/counter drift: sites=%d resolved=%d unavailable=%d excluded=%d mismatch=%d metrics=%d/%d/%d/%d",
			len(unique),
			counts[tsstdlib.LookupResolved],
			counts[tsstdlib.LookupCapabilityUnavailable],
			counts[tsstdlib.LookupProfileExcluded],
			counts[tsstdlib.LookupMeaningMismatch],
			metrics.ResolvedExternalDeclarations,
			metrics.ExternalCapabilityUnavailable,
			metrics.ExternalProfileExcluded,
			metrics.ExternalMeaningMismatches,
		)
	}
	return unique, nil
}

func cloneTypeScriptAuthorityResult(record TypeScriptAuthorityResult) TypeScriptAuthorityResult {
	record.DeclarationRanges = append([]tsstdlib.Declaration(nil), record.DeclarationRanges...)
	return record
}

func validateTypeScriptAuthorityResult(record TypeScriptAuthorityResult) error {
	if record.SourceSiteID == "" ||
		record.Stage != TypeScriptStandardLibraryStage ||
		record.FilePath == "" ||
		record.Range.StartLine <= 0 ||
		record.Range.EndLine <= 0 ||
		record.SiteKind == "" ||
		record.RequestedName == "" ||
		record.AuthorityKind != tsstdlib.AuthorityKind ||
		record.TypeScriptVersion != tsstdlib.TypeScriptVersion ||
		record.ProfileHash == "" ||
		record.ConfigHash == "" {
		return fmt.Errorf("incomplete TypeScript authority result for source site %q", record.SourceSiteID)
	}
	if err := validateTypeScriptCatalogProof(record); err != nil {
		return err
	}
	switch record.RequestedMeaning {
	case tsstdlib.MeaningValue, tsstdlib.MeaningType, tsstdlib.MeaningNamespace:
	default:
		return fmt.Errorf("invalid TypeScript authority meaning for source site %q", record.SourceSiteID)
	}
	switch record.Status {
	case tsstdlib.LookupResolved:
		if record.Reason != "" || record.ResolvedSymbolID == "" || len(record.DeclarationRanges) == 0 {
			return fmt.Errorf("incomplete resolved TypeScript authority result for source site %q", record.SourceSiteID)
		}
	case tsstdlib.LookupCapabilityUnavailable:
		if record.ResolvedSymbolID != "" || len(record.DeclarationRanges) != 0 {
			return fmt.Errorf("invalid unavailable TypeScript authority result for source site %q", record.SourceSiteID)
		}
	case tsstdlib.LookupProfileExcluded:
		if record.Reason != tsstdlib.ReasonProfileExcludes || record.ResolvedSymbolID != "" || len(record.DeclarationRanges) != 0 {
			return fmt.Errorf("invalid profile-excluded TypeScript authority result for source site %q", record.SourceSiteID)
		}
	case tsstdlib.LookupMeaningMismatch:
		if record.Reason != tsstdlib.ReasonMeaningMismatch || record.ResolvedSymbolID != "" || len(record.DeclarationRanges) != 0 {
			return fmt.Errorf("invalid meaning-mismatch TypeScript authority result for source site %q", record.SourceSiteID)
		}
	default:
		return fmt.Errorf("unhandled TypeScript authority status %q for source site %q", record.Status, record.SourceSiteID)
	}
	return nil
}

func validateTypeScriptCatalogProof(record TypeScriptAuthorityResult) error {
	switch record.CatalogProofState {
	case tsstdlib.CatalogProofReady:
		if record.AuthorityHash == "" || record.CatalogHash == "" || record.CatalogArtifactHash == "" {
			return fmt.Errorf("incomplete ready TypeScript catalog proof for source site %q", record.SourceSiteID)
		}
		switch record.Status {
		case tsstdlib.LookupResolved:
			if record.Reason != "" {
				return fmt.Errorf("invalid ready resolved TypeScript catalog proof for source site %q", record.SourceSiteID)
			}
		case tsstdlib.LookupProfileExcluded:
			if record.Reason != tsstdlib.ReasonProfileExcludes {
				return fmt.Errorf("invalid ready profile-excluded TypeScript catalog proof for source site %q", record.SourceSiteID)
			}
		case tsstdlib.LookupMeaningMismatch:
			if record.Reason != tsstdlib.ReasonMeaningMismatch {
				return fmt.Errorf("invalid ready meaning-mismatch TypeScript catalog proof for source site %q", record.SourceSiteID)
			}
		case tsstdlib.LookupCapabilityUnavailable:
			switch record.Reason {
			case tsstdlib.ReasonDisabledByNoLib,
				tsstdlib.ReasonConfigInvalid,
				tsstdlib.ReasonConfigTopology,
				tsstdlib.ReasonConfigUnreadable:
			default:
				return fmt.Errorf("invalid ready unavailable TypeScript catalog proof for source site %q", record.SourceSiteID)
			}
		default:
			return fmt.Errorf("invalid ready TypeScript catalog status %q for source site %q", record.Status, record.SourceSiteID)
		}
	case tsstdlib.CatalogProofMissing:
		if record.Status != tsstdlib.LookupCapabilityUnavailable ||
			record.Reason != tsstdlib.ReasonCatalogMissing ||
			record.AuthorityHash != "" ||
			record.CatalogHash != "" ||
			record.CatalogArtifactHash != "" {
			return fmt.Errorf("invalid missing TypeScript catalog proof for source site %q", record.SourceSiteID)
		}
	case tsstdlib.CatalogProofRejected:
		if record.Status != tsstdlib.LookupCapabilityUnavailable ||
			!isTypeScriptCatalogRejection(record.Reason) ||
			record.AuthorityHash != "" ||
			record.CatalogHash != "" ||
			record.CatalogArtifactHash == "" {
			return fmt.Errorf("invalid rejected TypeScript catalog proof for source site %q", record.SourceSiteID)
		}
	default:
		return fmt.Errorf("unknown TypeScript catalog proof state %q for source site %q", record.CatalogProofState, record.SourceSiteID)
	}
	return nil
}

func isTypeScriptCatalogRejection(reason tsstdlib.Reason) bool {
	switch reason {
	case tsstdlib.ReasonCatalogSchema,
		tsstdlib.ReasonCatalogVersion,
		tsstdlib.ReasonCatalogHash,
		tsstdlib.ReasonCatalogInputManifest:
		return true
	default:
		return false
	}
}

func countResolvedImports(imports []resolvedImport) int {
	count := 0
	for _, item := range imports {
		if len(item.TargetFiles) > 0 && item.LinkStatus != "unresolved" {
			count += len(item.TargetFiles)
		}
	}
	return count
}

func callEvidenceKind(form scopeir.CallForm) string {
	switch form {
	case scopeir.CallMember:
		return "type-binding"
	case scopeir.CallConstructor:
		return "kind-match"
	default:
		return "scope-chain"
	}
}

func isBuiltinType(name string) bool {
	switch name {
	case "string", "number", "boolean", "bool", "void", "undefined", "null", "any", "unknown", "never", "object":
		return true
	default:
		return false
	}
}
