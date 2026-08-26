package resolution

import (
	"fmt"
	"path"
	"reflect"
	"sort"

	"github.com/tamnguyendinh/anvien/internal/frameworks"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

const exportNodeLabel scopeir.NodeLabel = "Export"

type emitter struct {
	graph                   *graph.Graph
	diagnosticAppender      func(string, graphhealth.Diagnostic) bool
	referenceIndex          ReferenceIndex
	edgeKeys                map[string]graph.Relationship
	definitionNodesSeen     map[string]graph.Node
	typeScriptExternalSites []typeScriptExternalSite
	outcomes                resolutionOutcomeCollector
	fileLanguages           map[string]scanner.Language
	metrics                 *Metrics
	sourceLabel             string
}

func newEmitter(g *graph.Graph, metrics *Metrics, fileLanguages map[string]scanner.Language) *emitter {
	return &emitter{
		graph:               g,
		diagnosticAppender:  graphhealth.NewDiagnosticAppender(g),
		referenceIndex:      newReferenceIndex(),
		edgeKeys:            make(map[string]graph.Relationship),
		definitionNodesSeen: make(map[string]graph.Node),
		outcomes:            newResolutionOutcomeCollector(),
		fileLanguages:       fileLanguages,
		metrics:             metrics,
		sourceLabel:         "scope-resolution",
	}
}

func (e *emitter) emitNode(node graph.Node) {
	e.graph.AddNode(node)
	e.metrics.GraphNodesEmitted = len(e.graph.Nodes)
}

func (e *emitter) emitDefinitionNode(node graph.Node) error {
	incoming := node
	if emitted, ok := e.definitionNodesSeen[node.ID]; ok {
		if definitionNodesEqual(emitted, node) {
			e.metrics.GraphNodesEmitted = len(e.graph.Nodes)
			return nil
		}
		return definitionNodeCollisionError(emitted, node)
	}
	existing, ok := e.graph.GetNode(node.ID)
	if !ok {
		e.emitNode(node)
	} else {
		merged, compatible := mergeDefinitionNode(existing, node)
		if !compatible {
			return definitionNodeCollisionError(existing, node)
		}
		if !definitionNodesEqual(existing, merged) {
			e.emitNode(merged)
		}
	}
	e.definitionNodesSeen[node.ID] = incoming
	e.metrics.GraphNodesEmitted = len(e.graph.Nodes)
	return nil
}

func definitionNodesEqual(existing graph.Node, incoming graph.Node) bool {
	return existing.ID == incoming.ID &&
		existing.Label == incoming.Label &&
		reflect.DeepEqual(existing.Properties, incoming.Properties)
}

func definitionNodeCollisionError(existing graph.Node, incoming graph.Node) error {
	return fmt.Errorf(
		"definition identity collision for graph node %q: existing %q payload conflicts with incoming %q payload",
		incoming.ID,
		existing.Label,
		incoming.Label,
	)
}

func definitionNodeCompatible(existing graph.Node, incoming graph.Node) bool {
	_, compatible := mergeDefinitionNode(existing, incoming)
	return compatible
}

func mergeDefinitionNode(existing graph.Node, incoming graph.Node) (graph.Node, bool) {
	if existing.ID != incoming.ID || existing.Label != incoming.Label {
		return graph.Node{}, false
	}
	properties := make(graph.NodeProperties, len(existing.Properties)+1)
	for key, value := range existing.Properties {
		properties[key] = value
	}
	for key, incomingValue := range incoming.Properties {
		existingValue, ok := existing.Properties[key]
		if !ok {
			if key != "isExported" {
				return graph.Node{}, false
			}
			properties[key] = incomingValue
			continue
		}
		if !reflect.DeepEqual(existingValue, incomingValue) {
			return graph.Node{}, false
		}
	}
	merged := existing
	merged.Properties = properties
	return merged, true
}

func (e *emitter) emitRelationship(relationship graph.Relationship) {
	key := semanticEdgeKey(relationship)
	if existing, ok := e.edgeKeys[key]; ok {
		merged := mergeRelationship(existing, relationship)
		e.graph.ReplaceRelationship(existing.ID, merged)
		e.edgeKeys[key] = merged
		e.metrics.DuplicateEdgesMerged++
		e.metrics.GraphRelationshipsEmitted = len(e.graph.Relationships)
		return
	}
	e.graph.AddRelationship(relationship)
	e.edgeKeys[key] = relationship
	e.metrics.GraphRelationshipsEmitted = len(e.graph.Relationships)
}

func (e *emitter) emitReference(source defRef, target defRef, reference Reference) {
	if reference.SourceSiteID == "" {
		reference.SourceSiteID = sourceSiteID(siteKindForReference(reference.Kind), reference.FilePath, reference.TargetText, reference.Range)
	}
	e.recordRepositoryResolvedOutcome(target, reference)
	e.referenceIndex.add(reference)
	relType := relationshipTypeForReference(reference.Kind)
	reason := string(reference.Kind)
	if reference.Kind != ReferenceRead && reference.Kind != ReferenceWrite {
		reason = e.sourceLabel + ": " + string(reference.Kind) + " | confidence " + confidenceString(reference.Confidence)
	}
	relationship := graph.Relationship{
		ID:               "rel:" + string(relType) + ":" + source.GraphID + "->" + target.GraphID + ":" + intString(reference.Range.StartLine) + ":" + intString(reference.Range.StartCol),
		SourceID:         source.GraphID,
		TargetID:         target.GraphID,
		Type:             relType,
		Confidence:       reference.Confidence,
		Reason:           reason,
		ResolutionSource: e.sourceLabel,
		SourceSiteID:     reference.SourceSiteID,
		SourceSiteStatus: firstNonEmpty(reference.SourceSiteStatus, sourceSiteStatusResolved),
		ProofKind:        reference.ProofKind,
		TargetRole:       reference.TargetRole,
		TargetText:       reference.TargetText,
		FilePath:         cleanPath(reference.FilePath),
		FileHash:         reference.FileHash,
		StartLine:        reference.Range.StartLine,
		StartCol:         reference.Range.StartCol,
		EndLine:          reference.Range.EndLine,
		EndCol:           reference.Range.EndCol,
		Evidence:         reference.Evidence,
	}
	relationship.SourceSiteIDs = []string{relationship.SourceSiteID}
	relationship.SourceSiteCount = 1
	if reference.Kind == ReferenceRead {
		step := 1
		relationship.Step = &step
	}
	if reference.Kind == ReferenceWrite {
		step := 2
		relationship.Step = &step
	}
	e.emitRelationship(relationship)
}

func (e *emitter) emitUnresolvedReference(source defRef, factFamily string, targetText string, filePath string, fileHash string, factRange scopeir.Range, note string, incrementMetric bool, evidence ...graph.Evidence) {
	if incrementMetric {
		e.metrics.UnresolvedReferences++
	}
	status := sourceSiteStatusUnresolvedLocalBinding
	if source.GraphID == "" {
		status = sourceSiteStatusUnknown
	}
	if note == "heritage target text not modeled" || note == unresolvedNoteCallSourceFileLevel {
		status = sourceSiteStatusUnsupportedSyntax
	}
	proofKind := proofKindNone
	if note == "call target matched low-confidence global fallback only" {
		proofKind = proofKindGlobalFallbackLowConfidence
	}
	encoded, added := e.recordRepositoryUnresolvedOutcome(factFamily, targetText, filePath, fileHash, factRange, note, proofKind, evidence)
	if !added {
		return
	}
	diagnostic := graphhealth.Diagnostic{
		Kind:             graphhealth.DiagnosticUnresolvedReference,
		FactFamily:       factFamily,
		SourceNodeID:     source.GraphID,
		TargetText:       targetText,
		ResolutionSource: e.sourceLabel,
		FilePath:         cleanPath(filePath),
		FileHash:         fileHash,
		StartLine:        factRange.StartLine,
		StartCol:         factRange.StartCol,
		EndLine:          factRange.EndLine,
		EndCol:           factRange.EndCol,
		SourceSiteID:     sourceSiteID(factFamily, filePath, targetText, factRange),
		SourceSiteStatus: status,
		ProofKind:        proofKind,
		TargetRole:       targetRoleForFactFamily(factFamily),
		Note:             encoded,
		Source:           e.sourceLabel,
	}
	if e.diagnosticAppender(source.GraphID, diagnostic) {
		e.metrics.UnresolvedReferenceDiagnostics++
		return
	}
	e.metrics.UnattributedUnresolvedReferences++
}

func emitDefinitionNodes(w *workspace, e *emitter) error {
	emittedExportNodes := make(map[string]graph.Node)
	for _, ir := range w.files {
		directExportDefIDs, exportNodes, err := exportProjectionNodes(w, ir)
		if err != nil {
			return err
		}
		frameworkFacts := frameworkFactsByDefID(ir.Frameworks)
		fileID := graph.GenerateID("File", ir.FilePath)
		fileProperties := graph.NodeProperties{
			"name":     path.Base(ir.FilePath),
			"filePath": ir.FilePath,
			"language": string(ir.Language),
		}
		applyFrameworkHint(fileProperties, ir.FilePath)
		if existing, ok := e.graph.GetNode(fileID); ok && existing.Properties != nil {
			for key, value := range existing.Properties {
				if _, exists := fileProperties[key]; !exists {
					fileProperties[key] = value
				}
			}
		}
		e.emitNode(graph.Node{
			ID:         fileID,
			Label:      scopeir.NodeFile,
			Properties: fileProperties,
		})
		for _, exportNode := range exportNodes {
			if err := emitExportNode(e, emittedExportNodes, exportNode); err != nil {
				return err
			}
			e.emitRelationship(graph.Relationship{
				ID:         graph.GenerateID(string(graph.RelContains), fileID+"->"+exportNode.ID),
				SourceID:   fileID,
				TargetID:   exportNode.ID,
				Type:       graph.RelContains,
				Confidence: 1,
				Reason:     "source export fact",
			})
		}
		for _, def := range w.defsByFile[ir.FilePath] {
			props := graph.NodeProperties{
				"name":          def.Fact.Name,
				"filePath":      def.Fact.FilePath,
				"qualifiedName": def.Fact.QualifiedName,
				"startLine":     def.Fact.Range.StartLine,
				"startCol":      def.Fact.Range.StartCol,
				"endLine":       def.Fact.Range.EndLine,
				"endCol":        def.Fact.Range.EndCol,
				"language":      string(ir.Language),
			}
			if selection := def.Fact.SelectionRange; selection != nil {
				props["selectionStartLine"] = selection.StartLine
				props["selectionStartCol"] = selection.StartCol
				props["selectionEndLine"] = selection.EndLine
				props["selectionEndCol"] = selection.EndCol
			}
			applyFrameworkHint(props, def.Fact.FilePath)
			if fact, ok := frameworkFacts[def.Fact.ID]; ok {
				applyFrameworkFact(props, fact)
			}
			addNonEmpty(props, "returnType", def.Fact.ReturnType)
			addNonEmpty(props, "declaredType", def.Fact.DeclaredType)
			addNonEmpty(props, "visibility", def.Fact.Visibility)
			if ir.Language == scanner.JavaScript || ir.Language == scanner.TypeScript {
				_, exported := directExportDefIDs[def.Fact.ID]
				props["isExported"] = exported
			}
			if def.Fact.ParameterCount != nil {
				props["parameterCount"] = *def.Fact.ParameterCount
			}
			if len(def.Fact.ParameterTypes) > 0 {
				props["parameterTypes"] = append([]string(nil), def.Fact.ParameterTypes...)
			}
			if err := e.emitDefinitionNode(graph.Node{ID: def.GraphID, Label: def.Fact.Label, Properties: props}); err != nil {
				return err
			}
			e.emitRelationship(graph.Relationship{
				ID:         graph.GenerateID(string(graph.RelDefines), fileID+"->"+def.GraphID),
				SourceID:   fileID,
				TargetID:   def.GraphID,
				Type:       graph.RelDefines,
				Confidence: 1,
				Reason:     "",
			})
			if def.Fact.OwnerID != "" {
				owner, ok := w.defsByID[def.Fact.OwnerID]
				if !ok {
					continue
				}
				relType := graph.RelHasMethod
				if def.Fact.Label == scopeir.NodeProperty {
					relType = graph.RelHasProperty
				}
				e.emitRelationship(graph.Relationship{
					ID:         graph.GenerateID(string(relType), owner.GraphID+"->"+def.GraphID),
					SourceID:   owner.GraphID,
					TargetID:   def.GraphID,
					Type:       relType,
					Confidence: 1,
					Reason:     "",
				})
			}
		}
	}
	return nil
}

func exportProjectionNodes(w *workspace, ir scopeir.ScopeIR) (map[string]struct{}, []graph.Node, error) {
	directExportDefIDs := make(map[string]struct{}, len(ir.Exports))
	exportNodes := make([]graph.Node, 0, len(ir.Exports))
	for index, fact := range ir.Exports {
		factFilePath := cleanPath(fact.FilePath)
		if factFilePath == "" || factFilePath != ir.FilePath {
			return nil, nil, fmt.Errorf(
				"export fact %q at %s:%d:%d has file drift from owning ScopeIR %q",
				fact.ExportedName,
				factFilePath,
				fact.Range.StartLine,
				fact.Range.StartCol,
				ir.FilePath,
			)
		}

		localDefinitionNodeID := ""
		if fact.LocalDefID != "" {
			if fact.TargetRaw != nil {
				return nil, nil, fmt.Errorf(
					"source re-export %q at %s:%d:%d carries local definition %q",
					fact.ExportedName,
					factFilePath,
					fact.Range.StartLine,
					fact.Range.StartCol,
					fact.LocalDefID,
				)
			}
			localDefinition, ok := w.defsByID[fact.LocalDefID]
			if !ok {
				return nil, nil, fmt.Errorf(
					"local export %q at %s:%d:%d references missing definition %q",
					fact.ExportedName,
					factFilePath,
					fact.Range.StartLine,
					fact.Range.StartCol,
					fact.LocalDefID,
				)
			}
			if cleanPath(localDefinition.Fact.FilePath) != ir.FilePath {
				return nil, nil, fmt.Errorf(
					"local export %q at %s:%d:%d references cross-file definition %q in %q",
					fact.ExportedName,
					factFilePath,
					fact.Range.StartLine,
					fact.Range.StartCol,
					fact.LocalDefID,
					cleanPath(localDefinition.Fact.FilePath),
				)
			}
			localDefinitionNodeID = localDefinition.GraphID
			directExportDefIDs[fact.LocalDefID] = struct{}{}
		}

		exportNodes = append(exportNodes, exportGraphNode(fact, index, localDefinitionNodeID))
	}
	return directExportDefIDs, exportNodes, nil
}

func exportGraphNode(fact scopeir.ExportFact, index int, localDefinitionNodeID string) graph.Node {
	props := graph.NodeProperties{
		"name":               exportNodeName(fact),
		"filePath":           cleanPath(fact.FilePath),
		"kind":               string(fact.Kind),
		"exportedName":       fact.ExportedName,
		"meanings":           exportMeaningValues(fact.Meanings),
		"typeOnly":           fact.TypeOnly,
		"startLine":          fact.Range.StartLine,
		"startCol":           fact.Range.StartCol,
		"endLine":            fact.Range.EndLine,
		"endCol":             fact.Range.EndCol,
		"statementStartLine": fact.Provenance.StatementRange.StartLine,
		"statementStartCol":  fact.Provenance.StatementRange.StartCol,
		"statementEndLine":   fact.Provenance.StatementRange.EndLine,
		"statementEndCol":    fact.Provenance.StatementRange.EndCol,
		"siteKind":           fact.Provenance.SiteKind,
	}
	addNonEmpty(props, "fileHash", fact.FileHash)
	addNonEmpty(props, "localName", fact.LocalName)
	addNonEmpty(props, "localDefId", fact.LocalDefID)
	addNonEmpty(props, "localDefinitionNodeId", localDefinitionNodeID)
	addNonEmpty(props, "targetExportedName", fact.TargetExportedName)
	if fact.TargetRaw != nil {
		props["targetRaw"] = *fact.TargetRaw
	}
	if fact.SelectionRange != nil {
		props["selectionStartLine"] = fact.SelectionRange.StartLine
		props["selectionStartCol"] = fact.SelectionRange.StartCol
		props["selectionEndLine"] = fact.SelectionRange.EndLine
		props["selectionEndCol"] = fact.SelectionRange.EndCol
	}
	return graph.Node{
		ID:         exportGraphID(fact, index),
		Label:      exportNodeLabel,
		Properties: props,
	}
}

func exportGraphID(fact scopeir.ExportFact, index int) string {
	identity := "file" + graphIdentityPart(cleanPath(fact.FilePath)) +
		"ordinal" + graphIdentityPart(intString(index)) +
		"kind" + graphIdentityPart(string(fact.Kind)) +
		"name" + graphIdentityPart(fact.ExportedName) +
		"range" + graphIdentityPart(
		intString(fact.Range.StartLine)+":"+
			intString(fact.Range.StartCol)+":"+
			intString(fact.Range.EndLine)+":"+
			intString(fact.Range.EndCol),
	)
	return graph.GenerateID(string(exportNodeLabel), identity)
}

func exportNodeName(fact scopeir.ExportFact) string {
	if fact.ExportedName != "" {
		return fact.ExportedName
	}
	if fact.Kind == scopeir.ExportStar {
		return "*"
	}
	return string(fact.Kind)
}

func exportMeaningValues(meanings []scopeir.ExportMeaning) []string {
	values := make([]string, len(meanings))
	for index, meaning := range meanings {
		values[index] = string(meaning)
	}
	return values
}

func emitExportNode(e *emitter, emitted map[string]graph.Node, node graph.Node) error {
	if previous, ok := emitted[node.ID]; ok {
		if definitionNodesEqual(previous, node) {
			return nil
		}
		return fmt.Errorf("export fact identity collision for graph node %q", node.ID)
	}
	if existing, ok := e.graph.GetNode(node.ID); ok {
		if !definitionNodeCompatible(existing, node) {
			return fmt.Errorf("export fact identity collision for graph node %q", node.ID)
		}
	} else {
		e.emitNode(node)
	}
	emitted[node.ID] = node
	return nil
}

func frameworkFactsByDefID(facts []scopeir.FrameworkFact) map[string]scopeir.FrameworkFact {
	out := make(map[string]scopeir.FrameworkFact, len(facts))
	for _, fact := range facts {
		if fact.DefID == "" {
			continue
		}
		current, ok := out[fact.DefID]
		if !ok || fact.EntryPointMultiplier > current.EntryPointMultiplier {
			out[fact.DefID] = fact
		}
	}
	return out
}

func applyFrameworkHint(props graph.NodeProperties, filePath string) {
	hint, ok := frameworks.DetectFromPath(filePath)
	if !ok {
		return
	}
	props["framework"] = hint.Framework
	props["frameworkReason"] = hint.Reason
	props["frameworkEntryPointMultiplier"] = hint.EntryPointMultiplier
	props["astFrameworkReason"] = hint.Reason
	props["astFrameworkMultiplier"] = hint.EntryPointMultiplier
}

func applyFrameworkFact(props graph.NodeProperties, fact scopeir.FrameworkFact) {
	if fact.Framework != "" {
		props["framework"] = fact.Framework
	}
	props["frameworkReason"] = fact.Reason
	props["frameworkEntryPointMultiplier"] = fact.EntryPointMultiplier
	props["astFrameworkReason"] = fact.Reason
	props["astFrameworkMultiplier"] = fact.EntryPointMultiplier
}

func emitImportEdges(w *workspace, e *emitter) {
	for _, item := range w.imports {
		if len(item.TargetFiles) == 0 || item.LinkStatus == "unresolved" {
			continue
		}
		sourceFileID := graph.GenerateID("File", cleanPath(item.Fact.FilePath))
		for _, targetFile := range item.TargetFiles {
			targetFileID := graph.GenerateID("File", targetFile)
			e.emitRelationship(graph.Relationship{
				ID:               graph.GenerateID(string(graph.RelImports), cleanPath(item.Fact.FilePath)+"->"+targetFile),
				SourceID:         sourceFileID,
				TargetID:         targetFileID,
				Type:             graph.RelImports,
				Confidence:       1,
				Reason:           "scope-finalize import " + string(item.Fact.Kind) + " " + item.Fact.LocalName,
				ResolutionSource: "scope-finalize",
				FileHash:         w.fileHashes[cleanPath(item.Fact.FilePath)],
				Evidence: []graph.Evidence{{
					Kind:   "import",
					Weight: 1,
					Note:   string(item.Fact.Kind) + " " + item.Fact.LocalName + " -> " + targetFile,
				}},
			})
			e.metrics.FinalizedImportsEmitted++
		}
		if item.TargetDef == nil {
			continue
		}
		e.emitRelationship(graph.Relationship{
			ID:               graph.GenerateID(string(graph.RelUses), cleanPath(item.Fact.FilePath)+"->"+item.TargetDef.Fact.ID+":import:"+item.Fact.LocalName),
			SourceID:         sourceFileID,
			TargetID:         item.TargetDef.GraphID,
			Type:             graph.RelUses,
			Confidence:       1,
			Reason:           "scope-finalize import-use " + string(item.Fact.Kind) + " " + item.Fact.LocalName,
			ResolutionSource: "scope-finalize",
			FileHash:         w.fileHashes[cleanPath(item.Fact.FilePath)],
			Evidence: []graph.Evidence{{
				Kind:   "import",
				Weight: 1,
				Note:   string(item.Fact.Kind) + " " + item.Fact.LocalName + " -> " + item.TargetDef.Fact.Name,
			}},
		})
		e.metrics.ImportUsesEmitted++
	}
}

func emitHeritageCompatibilityEdges(e *emitter, item heritageResolution, emitInherits bool) {
	relType := graph.RelExtends
	if item.Fact.Kind == scopeir.HeritageImplements ||
		item.Fact.Kind == scopeir.HeritageTraitImpl ||
		item.Fact.Kind == scopeir.HeritageInclude ||
		item.Fact.Kind == scopeir.HeritageExtend ||
		item.Fact.Kind == scopeir.HeritagePrepend {
		relType = graph.RelImplements
	}
	e.emitRelationship(graph.Relationship{
		ID:         graph.GenerateID(string(relType), item.Owner.GraphID+"->"+item.Target.GraphID),
		SourceID:   item.Owner.GraphID,
		TargetID:   item.Target.GraphID,
		Type:       relType,
		Confidence: 1,
		Reason:     string(item.Fact.Kind),
		FileHash:   item.Fact.FileHash,
	})
	reference := Reference{
		FromScope:        item.Fact.InScope,
		ToDefID:          item.Target.Fact.ID,
		FilePath:         item.Fact.FilePath,
		FileHash:         item.Fact.FileHash,
		Range:            item.Fact.Range,
		Kind:             ReferenceInherits,
		Confidence:       1,
		SourceSiteID:     sourceSiteID("heritage", item.Fact.FilePath, item.Fact.Name, item.Fact.Range),
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        proofKindScopeBinding,
		TargetRole:       targetRoleType,
		TargetText:       item.Fact.Name,
		Evidence: []graph.Evidence{{
			Kind:   "scope-chain",
			Weight: 1,
			Note:   string(item.Fact.Kind) + " " + item.Fact.Name,
		}},
	}
	if !emitInherits {
		e.recordRepositoryResolvedOutcome(item.Target, reference)
		return
	}
	e.emitReference(item.Owner, item.Target, reference)
}

func emitMethodDispatchEdges(w *workspace, e *emitter) {
	ownMethodsByOwner := make(map[string]map[string][]defRef)
	for ownerID, members := range w.ownerMembers {
		ownMethodsByOwner[ownerID] = make(map[string][]defRef)
		for name, bucket := range members {
			for _, member := range bucket {
				if member.Fact.Label == scopeir.NodeMethod {
					ownMethodsByOwner[ownerID][name] = append(ownMethodsByOwner[ownerID][name], member)
				}
			}
		}
	}

	for _, item := range w.heritage {
		ownMethods := ownMethodsByOwner[item.Owner.Fact.ID]
		parentMethods := ownMethodsByOwner[item.Target.Fact.ID]
		for name, ownBucket := range ownMethods {
			parentBucket := parentMethods[name]
			if len(ownBucket) != 1 || len(parentBucket) != 1 {
				continue
			}
			if item.Fact.Kind == scopeir.HeritageImplements || item.Target.Fact.Label == scopeir.NodeInterface || item.Target.Fact.Label == scopeir.NodeTrait {
				e.emitRelationship(graph.Relationship{
					ID:         graph.GenerateID(string(graph.RelMethodImplements), ownBucket[0].GraphID+"->"+parentBucket[0].GraphID),
					SourceID:   ownBucket[0].GraphID,
					TargetID:   parentBucket[0].GraphID,
					Type:       graph.RelMethodImplements,
					Confidence: methodMatchConfidence(ownBucket[0], parentBucket[0]),
					Reason:     "method implements interface contract",
				})
				e.metrics.MethodImplementsEmitted++
				continue
			}
			e.emitRelationship(graph.Relationship{
				ID:         graph.GenerateID(string(graph.RelMethodOverrides), item.Owner.GraphID+"->"+parentBucket[0].GraphID),
				SourceID:   item.Owner.GraphID,
				TargetID:   parentBucket[0].GraphID,
				Type:       graph.RelMethodOverrides,
				Confidence: methodMatchConfidence(ownBucket[0], parentBucket[0]),
				Reason:     "first definition",
			})
			e.metrics.MethodOverridesEmitted++
		}
	}
}

func semanticEdgeKey(relationship graph.Relationship) string {
	accessKind := ""
	if relationship.Type == graph.RelAccesses {
		if relationship.Step != nil && *relationship.Step == 1 {
			accessKind = ":read"
		} else if relationship.Step != nil && *relationship.Step == 2 {
			accessKind = ":write"
		} else {
			accessKind = ":unknown"
		}
	}
	callName := ""
	if relationship.Type == graph.RelCalls {
		callName = ":call:" + relationshipCallName(relationship)
	}
	return relationship.SourceID + "\x00" + relationship.TargetID + "\x00" + string(relationship.Type) + accessKind + callName
}

func relationshipCallName(relationship graph.Relationship) string {
	for _, evidence := range relationship.Evidence {
		if evidence.Note != "" {
			return evidence.Note
		}
	}
	return relationship.ID
}

func mergeRelationship(existing graph.Relationship, incoming graph.Relationship) graph.Relationship {
	merged := existing
	if incoming.Confidence >= existing.Confidence {
		merged.Reason = incoming.Reason
	}
	if incoming.Confidence > merged.Confidence {
		merged.Confidence = incoming.Confidence
	}
	if merged.Step == nil && incoming.Step != nil {
		step := *incoming.Step
		merged.Step = &step
	}
	if incoming.ResolutionSource != "" {
		merged.ResolutionSource = incoming.ResolutionSource
	}
	if incoming.FileHash != "" {
		merged.FileHash = incoming.FileHash
	}
	if incoming.SourceSiteID != "" && merged.SourceSiteID == "" {
		merged.SourceSiteID = incoming.SourceSiteID
	}
	merged.SourceSiteIDs = mergeSourceSiteIDs(merged.SourceSiteIDs, incoming.SourceSiteIDs, incoming.SourceSiteID)
	if len(merged.SourceSiteIDs) > 0 {
		merged.SourceSiteCount = len(merged.SourceSiteIDs)
	} else if merged.SourceSiteCount == 0 && incoming.SourceSiteCount > 0 {
		merged.SourceSiteCount = incoming.SourceSiteCount
	}
	if incoming.SourceSiteStatus != "" {
		merged.SourceSiteStatus = incoming.SourceSiteStatus
	}
	if incoming.ProofKind != "" {
		merged.ProofKind = incoming.ProofKind
	}
	if incoming.TargetRole != "" {
		merged.TargetRole = incoming.TargetRole
	}
	if incoming.TargetText != "" {
		merged.TargetText = incoming.TargetText
	}
	if incoming.FilePath != "" {
		merged.FilePath = incoming.FilePath
	}
	if merged.StartLine == 0 || (incoming.StartLine > 0 && incoming.StartLine < merged.StartLine) {
		merged.StartLine = incoming.StartLine
		merged.StartCol = incoming.StartCol
		merged.EndLine = incoming.EndLine
		merged.EndCol = incoming.EndCol
	}
	merged.Evidence = mergeExportBindingEvidence(existing.Evidence, incoming.Evidence)
	return merged
}

func mergeSourceSiteIDs(existing []string, incoming []string, single string) []string {
	seen := make(map[string]bool, len(existing)+len(incoming)+1)
	out := make([]string, 0, len(existing)+len(incoming)+1)
	for _, value := range existing {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	for _, value := range incoming {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	if single != "" && !seen[single] {
		out = append(out, single)
	}
	return out
}

func relationshipTypeForReference(kind ReferenceKind) graph.RelationshipType {
	switch kind {
	case ReferenceCall:
		return graph.RelCalls
	case ReferenceRead, ReferenceWrite:
		return graph.RelAccesses
	case ReferenceInherits:
		return graph.RelInherits
	case ReferenceTypeReference, ReferenceImportUse:
		return graph.RelUses
	default:
		return graph.RelUses
	}
}

func methodMatchConfidence(left defRef, right defRef) float64 {
	if len(left.Fact.ParameterTypes) > 0 && len(right.Fact.ParameterTypes) > 0 {
		if stringSlicesEqual(left.Fact.ParameterTypes, right.Fact.ParameterTypes) {
			return 1
		}
		return 0.7
	}
	if left.Fact.ParameterCount != nil && right.Fact.ParameterCount != nil && *left.Fact.ParameterCount == *right.Fact.ParameterCount {
		return 1
	}
	return 0.7
}

func addNonEmpty(props graph.NodeProperties, key string, value string) {
	if value != "" {
		props[key] = value
	}
}

func stringSlicesEqual(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	leftCopy := append([]string(nil), left...)
	rightCopy := append([]string(nil), right...)
	sort.Strings(leftCopy)
	sort.Strings(rightCopy)
	for index := range leftCopy {
		if leftCopy[index] != rightCopy[index] {
			return false
		}
	}
	return true
}

func confidenceString(value float64) string {
	if value >= 1 {
		return "1.000"
	}
	if value <= 0 {
		return "0.000"
	}
	scaled := int(value*1000 + 0.5)
	return "0." + leftPad3(scaled)
}

func leftPad3(value int) string {
	if value >= 100 {
		return intString(value)
	}
	if value >= 10 {
		return "0" + intString(value)
	}
	return "00" + intString(value)
}
