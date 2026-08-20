package resolution

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
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
	e := newEmitter(g, &metrics)

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

	metrics.GraphNodesEmitted = len(g.Nodes)
	metrics.GraphRelationshipsEmitted = len(g.Relationships)
	graphhealth.SetResolutionMetadata(g, metrics.UnresolvedReferences, metrics.UnresolvedReferenceDiagnostics, metrics.UnattributedUnresolvedReferences)
	return Result{Graph: g, ReferenceIndex: e.referenceIndex, Metrics: metrics}, nil
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
	lowConfidenceFallback := false
	bindingReceiverResolved := false
	if call.CallForm == scopeir.CallMember {
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
			proofKind = proofKindScopeBinding
		}
		if !ok {
			target, ok = w.resolveSameFileName(call.FilePath, call.Name, dispatchOwnerLabels())
			if ok {
				confidence = 0.95
				proofKind = proofKindSameFile
			}
		}
		if !ok {
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
			target, ok = w.resolveImportedMember(call.ExplicitReceiver, call.Name, call.InScope, callableLabels())
			if ok {
				confidence = 0.9
				proofKind = proofKindImportMember
			}
		}
		if !ok && call.ExplicitReceiver == "" {
			target, ok = w.resolveGlobalCallName(call.Name, callableLabels(), call.Arity)
			if ok {
				confidence = 0.5
				lowConfidenceFallback = true
			}
		}
	default:
		target, ok = w.resolveScopedName(call.Name, call.InScope, callableLabels())
		if ok {
			proofKind = proofKindScopeBinding
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
			target, ok = w.resolveGlobalCallName(call.Name, callableLabels(), call.Arity)
			if ok {
				confidence = 0.5
				lowConfidenceFallback = true
			}
		}
	}
	if !ok {
		if bindingReceiverResolved {
			return
		}
		e.emitUnresolvedReference(source, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, "call target not resolved", true)
		return
	}
	if lowConfidenceFallback {
		e.emitUnresolvedReference(source, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, "call target matched low-confidence global fallback only", true)
		return
	}
	e.emitReference(source, target, Reference{
		FromScope:        call.InScope,
		ToDefID:          target.Fact.ID,
		FilePath:         call.FilePath,
		FileHash:         call.FileHash,
		Range:            call.Range,
		Kind:             ReferenceCall,
		Confidence:       confidence,
		SourceSiteID:     sourceSiteID("call", call.FilePath, callTargetText(call), call.Range),
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        proofKind,
		TargetRole:       targetRoleCallable,
		TargetText:       callTargetText(call),
		Evidence: []graph.Evidence{{
			Kind:   callEvidenceKind(call.CallForm),
			Weight: 1,
			Note:   call.Name,
		}},
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
	if access.ExplicitReceiver == "" {
		target, ok = bindingOccurrences.resolve(w, access.Name, access.InScope)
		if ok {
			evidenceKind = proofKindScopeBinding
			proofKind = proofKindScopeBinding
			targetRole = bindingOccurrenceTargetRole
		}
	} else {
		target, ok = w.resolveMember(access.Name, access.ExplicitReceiver, access.InScope, propertyLabels())
		if !ok {
			target, ok = w.resolveImportedMember(access.ExplicitReceiver, access.Name, access.InScope, propertyLabels())
			if ok {
				confidence = 0.9
				evidenceKind = "import-binding"
				proofKind = proofKindImportMember
			}
		}
	}
	if !ok {
		e.emitUnresolvedReference(source, "access", accessTargetText(access), access.FilePath, access.FileHash, access.Range, "access target not resolved", true)
		return
	}
	kind := ReferenceRead
	if access.Kind == scopeir.AccessWrite {
		kind = ReferenceWrite
	}
	source = bindingOccurrenceReferenceSource(source, target, access.FilePath)
	e.emitReference(source, target, Reference{
		FromScope:        access.InScope,
		ToDefID:          target.Fact.ID,
		FilePath:         access.FilePath,
		FileHash:         access.FileHash,
		Range:            access.Range,
		Kind:             kind,
		Confidence:       confidence,
		SourceSiteID:     sourceSiteID("access", access.FilePath, accessTargetText(access), access.Range),
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        proofKind,
		TargetRole:       targetRole,
		TargetText:       accessTargetText(access),
		Evidence: []graph.Evidence{{
			Kind:   evidenceKind,
			Weight: 1,
			Note:   access.ExplicitReceiver + "." + access.Name,
		}},
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
	if targetName == "" || isBuiltinType(targetName) {
		return
	}
	source, ok := sourceForScopeOrFile(w, annotation.InScope, annotation.FilePath)
	if !ok {
		e.emitUnresolvedReference(defRef{}, "type-reference", annotation.Type.RawName, annotation.FilePath, annotation.FileHash, annotation.Range, "source scope not resolved", true)
		return
	}
	target, ok := w.resolveName(targetName, annotation.InScope, typeLabels())
	if !ok {
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
