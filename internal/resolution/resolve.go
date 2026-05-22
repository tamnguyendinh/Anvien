package resolution

import (
	"errors"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/graphhealth"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
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

	emitDefinitionNodes(w, e)
	emitUnresolvedHeritageDiagnostics(w, e)
	emitImportEdges(w, e)

	emitInherits := !options.DisableScopeInheritsCompatibility
	for _, item := range w.heritage {
		emitHeritageCompatibilityEdges(e, item, emitInherits)
		metrics.ResolvedInheritance++
	}

	for _, ir := range w.files {
		for _, call := range ir.Calls {
			resolveCall(w, e, call)
		}
		for _, access := range ir.Accesses {
			resolveAccess(w, e, access)
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

func resolveCall(w *workspace, e *emitter, call scopeir.CallSiteFact) {
	source, ok := sourceForScopeOrFile(w, call.InScope, call.FilePath)
	if !ok {
		e.emitUnresolvedReference(defRef{}, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, "source scope not resolved", true)
		return
	}
	var target defRef
	confidence := 1.0
	lowConfidenceFallback := false
	switch call.CallForm {
	case scopeir.CallConstructor:
		target, ok = w.resolveScopedName(call.Name, call.InScope, dispatchOwnerLabels())
		if !ok {
			target, ok = w.resolveSameFileName(call.FilePath, call.Name, dispatchOwnerLabels())
			if ok {
				confidence = 0.95
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
		if !ok {
			target, ok = w.resolveImportedMember(call.ExplicitReceiver, call.Name, call.InScope, callableLabels())
			if ok {
				confidence = 0.9
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
		if !ok {
			target, ok = w.resolveSameFileName(call.FilePath, call.Name, callableLabels())
			if ok {
				confidence = 0.95
			}
		}
		if !ok {
			target, ok = w.resolveGoSamePackageFunction(call.FilePath, call.Name, call.Arity)
			if ok {
				confidence = 0.95
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
		e.emitUnresolvedReference(source, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, "call target not resolved", true)
		return
	}
	if lowConfidenceFallback {
		e.emitUnresolvedReference(source, "call", callTargetText(call), call.FilePath, call.FileHash, call.Range, "call target matched low-confidence global fallback only", true)
		return
	}
	e.emitReference(source, target, Reference{
		FromScope:  call.InScope,
		ToDefID:    target.Fact.ID,
		FileHash:   call.FileHash,
		Range:      call.Range,
		Kind:       ReferenceCall,
		Confidence: confidence,
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

func resolveAccess(w *workspace, e *emitter, access scopeir.AccessFact) {
	source, ok := sourceForScopeOrFile(w, access.InScope, access.FilePath)
	if !ok {
		e.emitUnresolvedReference(defRef{}, "access", accessTargetText(access), access.FilePath, access.FileHash, access.Range, "source scope not resolved", true)
		return
	}
	target, ok := w.resolveMember(access.Name, access.ExplicitReceiver, access.InScope, propertyLabels())
	confidence := 1.0
	evidenceKind := "type-binding"
	if !ok {
		target, ok = w.resolveImportedMember(access.ExplicitReceiver, access.Name, access.InScope, propertyLabels())
		if ok {
			confidence = 0.9
			evidenceKind = "import-binding"
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
	e.emitReference(source, target, Reference{
		FromScope:  access.InScope,
		ToDefID:    target.Fact.ID,
		FileHash:   access.FileHash,
		Range:      access.Range,
		Kind:       kind,
		Confidence: confidence,
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
		FromScope:  annotation.InScope,
		ToDefID:    target.Fact.ID,
		FileHash:   annotation.FileHash,
		Range:      annotation.Range,
		Kind:       ReferenceTypeReference,
		Confidence: 1,
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
