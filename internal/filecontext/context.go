package filecontext

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

const defaultSampleLimit = 10

type Options struct {
	RelationshipSamplesPerGroup int
	UnresolvedSamplesPerGroup   int
	LinkedSamplesPerKind        int
}

type FileListOptions struct {
	Sort                string
	Limit               int
	Offset              int
	Kinds               []string
	AppLayer            string
	FunctionalArea      string
	APIOnly             bool
	UnresolvedOnly      bool
	HighFanInOnly       bool
	HighFanOutOnly      bool
	HighFanInThreshold  int
	HighFanOutThreshold int
}

type FileList struct {
	Total  int           `json:"total"`
	Offset int           `json:"offset"`
	Limit  int           `json:"limit"`
	Sort   string        `json:"sort"`
	Files  []FileSummary `json:"files"`
}

type GraphInfo struct {
	Path          string `json:"path,omitempty"`
	IndexedCommit string `json:"indexedCommit,omitempty"`
	CurrentCommit string `json:"currentCommit,omitempty"`
	Stale         bool   `json:"stale"`
	AnalyzedAt    string `json:"analyzedAt,omitempty"`
}

type Target struct {
	Type                string               `json:"type"`
	Input               string               `json:"input"`
	NormalizedPath      string               `json:"normalizedPath,omitempty"`
	DispatchMode        string               `json:"dispatchMode,omitempty"`
	AmbiguityCandidates []AmbiguityCandidate `json:"ambiguityCandidates,omitempty"`
}

type AmbiguityCandidate struct {
	Type       string  `json:"type"`
	ID         string  `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	File       string  `json:"file,omitempty"`
	Line       int     `json:"line,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
	Command    string  `json:"command,omitempty"`
}

type FileContext struct {
	Repo          string               `json:"repo,omitempty"`
	RepoPath      string               `json:"repoPath,omitempty"`
	Graph         GraphInfo            `json:"graph"`
	Target        Target               `json:"target"`
	Summary       FileSummary          `json:"summary"`
	SymbolTree    []SymbolTreeNode     `json:"symbolTree"`
	Relationships RelationshipSections `json:"relationships"`
	Unresolved    UnresolvedSummary    `json:"unresolved"`
	Linked        LinkedSummary        `json:"linked"`
	Quality       QualitySignals       `json:"quality"`
	Limits        Limits               `json:"limits"`
}

type FileSummary struct {
	Path                      string `json:"path"`
	Language                  string `json:"language,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	AppLayer                  string `json:"appLayer,omitempty"`
	FunctionalArea            string `json:"functionalArea,omitempty"`
	ParseStatus               string `json:"parseStatus,omitempty"`
	SymbolCount               int    `json:"symbolCount"`
	ExportedSymbolCount       int    `json:"exportedSymbolCount"`
	InboundRefCount           int    `json:"inboundRefCount"`
	OutboundRefCount          int    `json:"outboundRefCount"`
	LocalRelationshipCount    int    `json:"localRelationshipCount"`
	UnresolvedSourceSiteCount int    `json:"unresolvedSourceSiteCount"`
	LinkedFlowCount           int    `json:"linkedFlowCount"`
	LinkedTestCount           int    `json:"linkedTestCount"`
	Risk                      string `json:"risk,omitempty"`
}

type SourceRange struct {
	StartLine   int `json:"startLine,omitempty"`
	StartColumn int `json:"startColumn,omitempty"`
	EndLine     int `json:"endLine,omitempty"`
	EndColumn   int `json:"endColumn,omitempty"`
}

type SymbolRelationshipCounts struct {
	Local      int `json:"local"`
	Inbound    int `json:"inbound"`
	Outbound   int `json:"outbound"`
	Unresolved int `json:"unresolved"`
}

type SymbolTreeNode struct {
	ID                 string                   `json:"id"`
	Name               string                   `json:"name"`
	Kind               string                   `json:"kind"`
	Range              SourceRange              `json:"range"`
	Exported           bool                     `json:"exported"`
	Signature          string                   `json:"signature,omitempty"`
	RelationshipCounts SymbolRelationshipCounts `json:"relationshipCounts"`
	Children           []SymbolTreeNode         `json:"children,omitempty"`
}

type RelationshipSections struct {
	Counts         RelationshipCounts      `json:"counts"`
	Local          RelationshipGroup       `json:"local"`
	OutboundByFile []FileRelationshipGroup `json:"outboundByFile"`
	InboundByFile  []FileRelationshipGroup `json:"inboundByFile"`
}

type RelationshipCounts struct {
	Local           int `json:"local"`
	Outbound        int `json:"outbound"`
	Inbound         int `json:"inbound"`
	SamplesReturned int `json:"samplesReturned"`
}

type RelationshipGroup struct {
	Total   int                  `json:"total"`
	Counts  map[string]int       `json:"counts,omitempty"`
	Samples []RelationshipSample `json:"samples"`
}

type FileRelationshipGroup struct {
	File    string               `json:"file"`
	Total   int                  `json:"total"`
	Counts  map[string]int       `json:"counts,omitempty"`
	Samples []RelationshipSample `json:"samples"`
}

type RelationshipSample struct {
	SourceFile       string      `json:"sourceFile,omitempty"`
	SourceSymbol     string      `json:"sourceSymbol,omitempty"`
	SourceRange      SourceRange `json:"sourceRange"`
	RelationshipKind string      `json:"relationshipKind"`
	TargetFile       string      `json:"targetFile,omitempty"`
	TargetSymbol     string      `json:"targetSymbol,omitempty"`
	TargetRange      SourceRange `json:"targetRange"`
	SourceSiteID     string      `json:"sourceSiteId,omitempty"`
	ProofKind        string      `json:"proofKind,omitempty"`
	SourceSiteStatus string      `json:"sourceSiteStatus,omitempty"`
}

type UnresolvedSummary struct {
	Total            int               `json:"total"`
	ByKind           map[string]int    `json:"byKind,omitempty"`
	ByClassification map[string]int    `json:"byClassification,omitempty"`
	ByActionability  map[string]int    `json:"byActionability,omitempty"`
	Groups           []UnresolvedGroup `json:"groups"`
}

type UnresolvedGroup struct {
	SourceSymbol string             `json:"sourceSymbol,omitempty"`
	Total        int                `json:"total"`
	Samples      []UnresolvedSample `json:"samples"`
}

type UnresolvedSample struct {
	Line             int    `json:"line,omitempty"`
	Column           int    `json:"column,omitempty"`
	TargetText       string `json:"targetText,omitempty"`
	SourceSymbol     string `json:"sourceSymbol,omitempty"`
	GapKind          string `json:"gapKind,omitempty"`
	Classification   string `json:"classification,omitempty"`
	Actionability    string `json:"actionability,omitempty"`
	ProofKind        string `json:"proofKind,omitempty"`
	SourceSiteID     string `json:"sourceSiteId,omitempty"`
	SourceSiteStatus string `json:"sourceSiteStatus,omitempty"`
}

type LinkedSummary struct {
	Flows    []LinkedItem `json:"flows"`
	Routes   []LinkedItem `json:"routes"`
	MCPTools []LinkedItem `json:"mcpTools"`
	Tests    []LinkedItem `json:"tests"`
}

type LinkedItem struct {
	Name       string `json:"name"`
	Kind       string `json:"kind,omitempty"`
	Source     string `json:"source,omitempty"`
	Confidence string `json:"confidence,omitempty"`
	Trace      string `json:"trace,omitempty"`
}

type QualitySignals struct {
	Parser               string `json:"parser,omitempty"`
	ResolutionConfidence string `json:"resolutionConfidence,omitempty"`
	UnresolvedCalls      int    `json:"unresolvedCalls"`
	UnresolvedRefs       int    `json:"unresolvedRefs"`
	UnresolvedImports    int    `json:"unresolvedImports"`
	Generated            bool   `json:"generated"`
	Stale                bool   `json:"stale"`
	ChangedSinceAnalyze  bool   `json:"changedSinceAnalyze"`
}

type Limits struct {
	RelationshipSamplesPerGroup int `json:"relationshipSamplesPerGroup"`
	UnresolvedSamplesPerGroup   int `json:"unresolvedSamplesPerGroup"`
	LinkedSamplesPerKind        int `json:"linkedSamplesPerKind"`
}

type Builder struct {
	g *graph.Graph

	nodesByID         map[string]graph.Node
	filesByPath       map[string]graph.Node
	filePathsByNodeID map[string]string
	definesByFile     map[string][]string
	containsByParent  map[string][]string
	parentByChild     map[string]string
	unresolvedByFile  map[string][]graph.Node
}

func NewBuilder(g *graph.Graph) *Builder {
	if g == nil {
		g = graph.New()
	}
	builder := &Builder{g: g}
	builder.buildIndexes()
	return builder
}

func (b *Builder) BuildFileContext(path string, options Options) (FileContext, bool) {
	limits := normalizeLimits(options)
	normalizedPath := normalizePath(path)
	fileNode, ok := b.filesByPath[normalizedPath]
	if !ok {
		return FileContext{}, false
	}

	relationships := b.buildRelationships(normalizedPath, limits.RelationshipSamplesPerGroup)
	unresolved := b.buildUnresolved(normalizedPath, limits.UnresolvedSamplesPerGroup)
	symbolTree, symbolCount, exportedCount := b.buildSymbolTree(normalizedPath, unresolved)
	quality := buildQuality(fileNode, unresolved)
	linked := LinkedSummary{
		Flows:    []LinkedItem{},
		Routes:   []LinkedItem{},
		MCPTools: []LinkedItem{},
		Tests:    []LinkedItem{},
	}

	summary := FileSummary{
		Path:                      normalizedPath,
		Language:                  stringProperty(fileNode, "language"),
		Kind:                      fileKind(fileNode),
		AppLayer:                  stringProperty(fileNode, "appLayer"),
		FunctionalArea:            stringProperty(fileNode, "functionalArea"),
		ParseStatus:               parseStatus(fileNode),
		SymbolCount:               symbolCount,
		ExportedSymbolCount:       exportedCount,
		InboundRefCount:           relationships.Counts.Inbound,
		OutboundRefCount:          relationships.Counts.Outbound,
		LocalRelationshipCount:    relationships.Counts.Local,
		UnresolvedSourceSiteCount: unresolved.Total,
		LinkedFlowCount:           len(linked.Flows),
		LinkedTestCount:           len(linked.Tests),
		Risk:                      riskFor(unresolved.Total),
	}

	return FileContext{
		Target: Target{
			Type:           "file",
			Input:          path,
			NormalizedPath: normalizedPath,
			DispatchMode:   "explicit",
		},
		Summary:       summary,
		SymbolTree:    symbolTree,
		Relationships: relationships,
		Unresolved:    unresolved,
		Linked:        linked,
		Quality:       quality,
		Limits:        limits,
	}, true
}

func (b *Builder) BuildFileList(options FileListOptions) FileList {
	summaries := b.buildFileSummaries()
	summaries = filterSummaries(summaries, options)
	sortSummaries(summaries, options.Sort)

	total := len(summaries)
	offset := options.Offset
	if offset < 0 {
		offset = 0
	}
	if offset > total {
		offset = total
	}
	limit := options.Limit
	end := total
	if limit > 0 && offset+limit < end {
		end = offset + limit
	}

	return FileList{
		Total:  total,
		Offset: offset,
		Limit:  limit,
		Sort:   normalizedSort(options.Sort),
		Files:  append([]FileSummary(nil), summaries[offset:end]...),
	}
}

func (b *Builder) buildIndexes() {
	b.nodesByID = make(map[string]graph.Node, len(b.g.Nodes))
	b.filesByPath = make(map[string]graph.Node)
	b.filePathsByNodeID = make(map[string]string, len(b.g.Nodes))
	b.definesByFile = make(map[string][]string)
	b.containsByParent = make(map[string][]string)
	b.parentByChild = make(map[string]string)
	b.unresolvedByFile = make(map[string][]graph.Node)

	for _, node := range b.g.Nodes {
		b.nodesByID[node.ID] = node
		filePath := normalizePath(stringProperty(node, "filePath"))
		if filePath != "" {
			b.filePathsByNodeID[node.ID] = filePath
		}
		if node.Label == scopeir.NodeFile && filePath != "" {
			b.filesByPath[filePath] = node
		}
		if node.Label == scopeir.NodeResolutionGap && filePath != "" {
			b.unresolvedByFile[filePath] = append(b.unresolvedByFile[filePath], node)
		}
	}

	for _, relationship := range b.g.Relationships {
		switch relationship.Type {
		case graph.RelDefines:
			if _, ok := b.filesByPath[normalizePathFromNodeID(relationship.SourceID)]; ok {
				sourcePath := normalizePathFromNodeID(relationship.SourceID)
				b.definesByFile[sourcePath] = append(b.definesByFile[sourcePath], relationship.TargetID)
			}
		case graph.RelContains:
			b.containsByParent[relationship.SourceID] = append(b.containsByParent[relationship.SourceID], relationship.TargetID)
			if _, exists := b.parentByChild[relationship.TargetID]; !exists {
				b.parentByChild[relationship.TargetID] = relationship.SourceID
			}
		}
	}

	for path := range b.definesByFile {
		sort.Strings(b.definesByFile[path])
	}
	for parent := range b.containsByParent {
		sort.Strings(b.containsByParent[parent])
	}
	for path := range b.unresolvedByFile {
		sort.SliceStable(b.unresolvedByFile[path], func(i, j int) bool {
			return compareUnresolvedNodes(b.unresolvedByFile[path][i], b.unresolvedByFile[path][j]) < 0
		})
	}
}

func (b *Builder) buildFileSummaries() []FileSummary {
	aggregates := make(map[string]*FileSummary, len(b.filesByPath))
	for path, node := range b.filesByPath {
		summary := &FileSummary{
			Path:           path,
			Language:       stringProperty(node, "language"),
			Kind:           fileKind(node),
			AppLayer:       stringProperty(node, "appLayer"),
			FunctionalArea: stringProperty(node, "functionalArea"),
			ParseStatus:    parseStatus(node),
			SymbolCount:    len(b.definesByFile[path]),
			Risk:           "low",
		}
		for _, symbolID := range b.definesByFile[path] {
			if exportedSymbol(b.nodesByID[symbolID]) {
				summary.ExportedSymbolCount++
			}
		}
		summary.UnresolvedSourceSiteCount = len(b.unresolvedByFile[path])
		summary.Risk = riskFor(summary.UnresolvedSourceSiteCount)
		aggregates[path] = summary
	}

	for _, relationship := range b.g.Relationships {
		if !isProjectionRelationship(relationship.Type) {
			continue
		}
		sourceNode, sourceOK := b.nodesByID[relationship.SourceID]
		targetNode, targetOK := b.nodesByID[relationship.TargetID]
		if !sourceOK || !targetOK {
			continue
		}
		sourceFile := b.nodeFilePath(sourceNode)
		targetFile := b.nodeFilePath(targetNode)
		if sourceFile == "" {
			sourceFile = normalizePath(relationship.FilePath)
		}
		if sourceFile == "" || targetFile == "" {
			continue
		}
		if sourceFile == targetFile {
			if summary := aggregates[sourceFile]; summary != nil {
				summary.LocalRelationshipCount++
			}
			continue
		}
		if summary := aggregates[sourceFile]; summary != nil {
			summary.OutboundRefCount++
		}
		if summary := aggregates[targetFile]; summary != nil {
			summary.InboundRefCount++
		}
	}

	paths := make([]string, 0, len(aggregates))
	for path := range aggregates {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	summaries := make([]FileSummary, 0, len(paths))
	for _, path := range paths {
		summaries = append(summaries, *aggregates[path])
	}
	return summaries
}

func (b *Builder) buildSymbolTree(path string, unresolved UnresolvedSummary) ([]SymbolTreeNode, int, int) {
	defined := b.definesByFile[path]
	definedSet := make(map[string]struct{}, len(defined))
	for _, id := range defined {
		definedSet[id] = struct{}{}
	}

	unresolvedCounts := make(map[string]int)
	for _, group := range unresolved.Groups {
		if group.SourceSymbol != "" {
			unresolvedCounts[group.SourceSymbol] += group.Total
		}
	}

	counts := b.symbolRelationshipCounts(path, unresolvedCounts)
	roots := make([]string, 0, len(defined))
	for _, id := range defined {
		parent := b.parentByChild[id]
		if _, parentInFile := definedSet[parent]; parentInFile {
			continue
		}
		roots = append(roots, id)
	}
	sortSymbolIDs(roots, b.nodesByID)

	tree := make([]SymbolTreeNode, 0, len(roots))
	exported := 0
	for _, id := range roots {
		node, ok := b.nodesByID[id]
		if !ok {
			continue
		}
		item := b.buildSymbolNode(node, path, definedSet, counts)
		exported += countExported(item)
		tree = append(tree, item)
	}
	return tree, len(defined), exported
}

func (b *Builder) buildSymbolNode(node graph.Node, path string, definedSet map[string]struct{}, counts map[string]SymbolRelationshipCounts) SymbolTreeNode {
	childIDs := append([]string(nil), b.containsByParent[node.ID]...)
	filtered := childIDs[:0]
	for _, childID := range childIDs {
		child, ok := b.nodesByID[childID]
		if !ok {
			continue
		}
		childPath := normalizePath(stringProperty(child, "filePath"))
		if childPath == "" {
			childPath = path
		}
		if childPath != path {
			continue
		}
		if _, defined := definedSet[childID]; !defined {
			continue
		}
		filtered = append(filtered, childID)
	}
	sortSymbolIDs(filtered, b.nodesByID)

	children := make([]SymbolTreeNode, 0, len(filtered))
	for _, childID := range filtered {
		children = append(children, b.buildSymbolNode(b.nodesByID[childID], path, definedSet, counts))
	}

	return SymbolTreeNode{
		ID:                 node.ID,
		Name:               nodeName(node),
		Kind:               string(node.Label),
		Range:              nodeRange(node),
		Exported:           exportedSymbol(node),
		Signature:          stringProperty(node, "signature"),
		RelationshipCounts: counts[node.ID],
		Children:           children,
	}
}

func (b *Builder) buildRelationships(path string, sampleLimit int) RelationshipSections {
	localSamples := []RelationshipSample{}
	outbound := map[string][]RelationshipSample{}
	inbound := map[string][]RelationshipSample{}
	localCounts := map[string]int{}
	outboundCounts := map[string]map[string]int{}
	inboundCounts := map[string]map[string]int{}

	for _, relationship := range b.g.Relationships {
		if !isProjectionRelationship(relationship.Type) {
			continue
		}
		sourceNode, sourceOK := b.nodesByID[relationship.SourceID]
		targetNode, targetOK := b.nodesByID[relationship.TargetID]
		if !sourceOK || !targetOK {
			continue
		}
		sourceFile := b.nodeFilePath(sourceNode)
		targetFile := b.nodeFilePath(targetNode)
		if sourceFile == "" {
			sourceFile = normalizePath(relationship.FilePath)
		}
		if sourceFile == "" || targetFile == "" {
			continue
		}
		sample := relationshipSample(relationship, sourceNode, targetNode, sourceFile, targetFile)
		kind := string(relationship.Type)
		switch {
		case sourceFile == path && targetFile == path:
			localCounts[kind]++
			localSamples = append(localSamples, sample)
		case sourceFile == path:
			if outboundCounts[targetFile] == nil {
				outboundCounts[targetFile] = map[string]int{}
			}
			outboundCounts[targetFile][kind]++
			outbound[targetFile] = append(outbound[targetFile], sample)
		case targetFile == path:
			if inboundCounts[sourceFile] == nil {
				inboundCounts[sourceFile] = map[string]int{}
			}
			inboundCounts[sourceFile][kind]++
			inbound[sourceFile] = append(inbound[sourceFile], sample)
		}
	}

	sortSamples(localSamples)
	local := RelationshipGroup{
		Total:   len(localSamples),
		Counts:  localCounts,
		Samples: limitSamples(localSamples, sampleLimit),
	}
	outboundGroups := fileRelationshipGroups(outbound, outboundCounts, sampleLimit)
	inboundGroups := fileRelationshipGroups(inbound, inboundCounts, sampleLimit)

	samplesReturned := len(local.Samples)
	outboundTotal := 0
	for _, group := range outboundGroups {
		outboundTotal += group.Total
		samplesReturned += len(group.Samples)
	}
	inboundTotal := 0
	for _, group := range inboundGroups {
		inboundTotal += group.Total
		samplesReturned += len(group.Samples)
	}

	return RelationshipSections{
		Counts: RelationshipCounts{
			Local:           local.Total,
			Outbound:        outboundTotal,
			Inbound:         inboundTotal,
			SamplesReturned: samplesReturned,
		},
		Local:          local,
		OutboundByFile: outboundGroups,
		InboundByFile:  inboundGroups,
	}
}

func (b *Builder) buildUnresolved(path string, sampleLimit int) UnresolvedSummary {
	nodes := append([]graph.Node(nil), b.unresolvedByFile[path]...)
	sort.SliceStable(nodes, func(i, j int) bool {
		return compareUnresolvedNodes(nodes[i], nodes[j]) < 0
	})

	byKind := map[string]int{}
	byClassification := map[string]int{}
	byActionability := map[string]int{}
	grouped := map[string][]UnresolvedSample{}

	for _, node := range nodes {
		kind := stringProperty(node, "gapKind")
		if kind == "" {
			kind = stringProperty(node, "factFamily")
		}
		classification := stringProperty(node, "classification")
		actionability := stringProperty(node, "actionability")
		increment(byKind, kind)
		increment(byClassification, classification)
		increment(byActionability, actionability)

		sourceSymbol := b.sourceSymbolName(node)
		sample := UnresolvedSample{
			Line:             intProperty(node, "startLine"),
			Column:           intProperty(node, "startCol"),
			TargetText:       firstNonEmpty(stringProperty(node, "targetText"), stringProperty(node, "name")),
			SourceSymbol:     sourceSymbol,
			GapKind:          kind,
			Classification:   classification,
			Actionability:    actionability,
			ProofKind:        stringProperty(node, "proofKind"),
			SourceSiteID:     stringProperty(node, "sourceSiteId"),
			SourceSiteStatus: stringProperty(node, "sourceSiteStatus"),
		}
		grouped[sourceSymbol] = append(grouped[sourceSymbol], sample)
	}

	groupNames := make([]string, 0, len(grouped))
	for name := range grouped {
		groupNames = append(groupNames, name)
	}
	sort.Strings(groupNames)

	groups := make([]UnresolvedGroup, 0, len(groupNames))
	for _, name := range groupNames {
		samples := grouped[name]
		sortUnresolvedSamples(samples)
		groups = append(groups, UnresolvedGroup{
			SourceSymbol: name,
			Total:        len(samples),
			Samples:      limitUnresolvedSamples(samples, sampleLimit),
		})
	}

	return UnresolvedSummary{
		Total:            len(nodes),
		ByKind:           compactCounts(byKind),
		ByClassification: compactCounts(byClassification),
		ByActionability:  compactCounts(byActionability),
		Groups:           groups,
	}
}

func (b *Builder) symbolRelationshipCounts(path string, unresolvedCounts map[string]int) map[string]SymbolRelationshipCounts {
	counts := map[string]SymbolRelationshipCounts{}
	for _, relationship := range b.g.Relationships {
		if !isProjectionRelationship(relationship.Type) {
			continue
		}
		sourceNode, sourceOK := b.nodesByID[relationship.SourceID]
		targetNode, targetOK := b.nodesByID[relationship.TargetID]
		if !sourceOK || !targetOK {
			continue
		}
		sourceFile := b.nodeFilePath(sourceNode)
		targetFile := b.nodeFilePath(targetNode)
		if sourceFile == "" {
			sourceFile = normalizePath(relationship.FilePath)
		}
		if sourceFile == "" || targetFile == "" {
			continue
		}
		if sourceFile == path {
			value := counts[sourceNode.ID]
			if targetFile == path {
				value.Local++
			} else {
				value.Outbound++
			}
			counts[sourceNode.ID] = value
		}
		if targetFile == path {
			value := counts[targetNode.ID]
			if sourceFile == path {
				value.Local++
			} else {
				value.Inbound++
			}
			counts[targetNode.ID] = value
		}
	}
	for symbol, count := range unresolvedCounts {
		value := counts[symbol]
		value.Unresolved += count
		counts[symbol] = value
	}
	return counts
}

func (b *Builder) nodeFilePath(node graph.Node) string {
	if filePath := normalizePath(stringProperty(node, "filePath")); filePath != "" {
		return filePath
	}
	return b.filePathsByNodeID[node.ID]
}

func (b *Builder) sourceSymbolName(node graph.Node) string {
	sourceID := stringProperty(node, "sourceNodeId")
	if sourceID == "" {
		return ""
	}
	source, ok := b.nodesByID[sourceID]
	if !ok {
		return sourceID
	}
	return source.ID
}

func normalizeLimits(options Options) Limits {
	relationshipLimit := options.RelationshipSamplesPerGroup
	if relationshipLimit <= 0 {
		relationshipLimit = defaultSampleLimit
	}
	unresolvedLimit := options.UnresolvedSamplesPerGroup
	if unresolvedLimit <= 0 {
		unresolvedLimit = defaultSampleLimit
	}
	linkedLimit := options.LinkedSamplesPerKind
	if linkedLimit <= 0 {
		linkedLimit = defaultSampleLimit
	}
	return Limits{
		RelationshipSamplesPerGroup: relationshipLimit,
		UnresolvedSamplesPerGroup:   unresolvedLimit,
		LinkedSamplesPerKind:        linkedLimit,
	}
}

func normalizePath(path string) string {
	path = strings.TrimSpace(strings.ReplaceAll(path, "\\", "/"))
	path = strings.TrimPrefix(path, "./")
	return path
}

func normalizePathFromNodeID(id string) string {
	return normalizePath(strings.TrimPrefix(id, "File:"))
}

func isProjectionRelationship(relationshipType graph.RelationshipType) bool {
	switch relationshipType {
	case graph.RelDefines, graph.RelContains, graph.RelHasResolutionGap:
		return false
	default:
		return true
	}
}

func relationshipSample(relationship graph.Relationship, source graph.Node, target graph.Node, sourceFile string, targetFile string) RelationshipSample {
	return RelationshipSample{
		SourceFile:       sourceFile,
		SourceSymbol:     source.ID,
		SourceRange:      relationshipRange(relationship, source),
		RelationshipKind: string(relationship.Type),
		TargetFile:       targetFile,
		TargetSymbol:     target.ID,
		TargetRange:      nodeRange(target),
		SourceSiteID:     relationship.SourceSiteID,
		ProofKind:        relationship.ProofKind,
		SourceSiteStatus: relationship.SourceSiteStatus,
	}
}

func relationshipRange(relationship graph.Relationship, fallback graph.Node) SourceRange {
	if relationship.StartLine != 0 || relationship.EndLine != 0 || relationship.StartCol != 0 || relationship.EndCol != 0 {
		return SourceRange{
			StartLine:   relationship.StartLine,
			StartColumn: relationship.StartCol,
			EndLine:     relationship.EndLine,
			EndColumn:   relationship.EndCol,
		}
	}
	return nodeRange(fallback)
}

func nodeRange(node graph.Node) SourceRange {
	return SourceRange{
		StartLine:   intProperty(node, "startLine"),
		StartColumn: intProperty(node, "startCol"),
		EndLine:     intProperty(node, "endLine"),
		EndColumn:   intProperty(node, "endCol"),
	}
}

func nodeName(node graph.Node) string {
	return firstNonEmpty(stringProperty(node, "name"), node.ID)
}

func stringProperty(node graph.Node, key string) string {
	if node.Properties == nil {
		return ""
	}
	value, ok := node.Properties[key]
	if !ok || value == nil {
		return ""
	}
	switch typed := value.(type) {
	case string:
		return typed
	default:
		return strings.TrimSpace(fmt.Sprint(typed))
	}
}

func intProperty(node graph.Node, key string) int {
	if node.Properties == nil {
		return 0
	}
	switch value := node.Properties[key].(type) {
	case int:
		return value
	case int64:
		return int(value)
	case float64:
		return int(value)
	case float32:
		return int(value)
	default:
		return 0
	}
}

func boolProperty(node graph.Node, key string) bool {
	if node.Properties == nil {
		return false
	}
	value, _ := node.Properties[key].(bool)
	return value
}

func exportedSymbol(node graph.Node) bool {
	if boolProperty(node, "exported") {
		return true
	}
	visibility := strings.ToLower(stringProperty(node, "visibility"))
	return visibility == "public" || visibility == "exported"
}

func parseStatus(node graph.Node) string {
	if value := stringProperty(node, "parseStatus"); value != "" {
		return value
	}
	if node.Label == scopeir.NodeFile {
		return "parsed"
	}
	return ""
}

func fileKind(node graph.Node) string {
	appLayer := strings.ToLower(stringProperty(node, "appLayer"))
	switch {
	case strings.Contains(appLayer, "test"):
		return "test"
	case strings.Contains(appLayer, "generated"):
		return "generated"
	case appLayer == "docs":
		return "docs"
	case appLayer == "config":
		return "config"
	case appLayer == "":
		return "unknown"
	default:
		return "source"
	}
}

func buildQuality(node graph.Node, unresolved UnresolvedSummary) QualitySignals {
	generated := fileKind(node) == "generated" || boolProperty(node, "generated")
	calls := 0
	refs := 0
	imports := 0
	for kind, count := range unresolved.ByKind {
		lower := strings.ToLower(kind)
		switch {
		case strings.Contains(lower, "call"):
			calls += count
		case strings.Contains(lower, "import"):
			imports += count
		default:
			refs += count
		}
	}
	return QualitySignals{
		Parser:               parseStatus(node),
		ResolutionConfidence: "unknown",
		UnresolvedCalls:      calls,
		UnresolvedRefs:       refs,
		UnresolvedImports:    imports,
		Generated:            generated,
	}
}

func filterSummaries(summaries []FileSummary, options FileListOptions) []FileSummary {
	allowedKinds := map[string]struct{}{}
	for _, kind := range options.Kinds {
		kind = strings.ToLower(strings.TrimSpace(kind))
		if kind != "" {
			allowedKinds[kind] = struct{}{}
		}
	}
	appLayer := strings.ToLower(strings.TrimSpace(options.AppLayer))
	functionalArea := strings.ToLower(strings.TrimSpace(options.FunctionalArea))
	fanInThreshold := options.HighFanInThreshold
	if fanInThreshold <= 0 {
		fanInThreshold = 10
	}
	fanOutThreshold := options.HighFanOutThreshold
	if fanOutThreshold <= 0 {
		fanOutThreshold = 10
	}

	filtered := summaries[:0]
	for _, summary := range summaries {
		if len(allowedKinds) > 0 {
			if _, ok := allowedKinds[strings.ToLower(summary.Kind)]; !ok {
				continue
			}
		}
		if appLayer != "" && strings.ToLower(summary.AppLayer) != appLayer {
			continue
		}
		if functionalArea != "" && strings.ToLower(summary.FunctionalArea) != functionalArea {
			continue
		}
		if options.APIOnly && !isAPIFile(summary) {
			continue
		}
		if options.UnresolvedOnly && summary.UnresolvedSourceSiteCount == 0 {
			continue
		}
		if options.HighFanInOnly && summary.InboundRefCount < fanInThreshold {
			continue
		}
		if options.HighFanOutOnly && summary.OutboundRefCount < fanOutThreshold {
			continue
		}
		filtered = append(filtered, summary)
	}
	return append([]FileSummary(nil), filtered...)
}

func sortSummaries(summaries []FileSummary, sortMode string) {
	mode := normalizedSort(sortMode)
	sort.SliceStable(summaries, func(i, j int) bool {
		left, right := summaries[i], summaries[j]
		switch mode {
		case "unresolved":
			if left.UnresolvedSourceSiteCount != right.UnresolvedSourceSiteCount {
				return left.UnresolvedSourceSiteCount > right.UnresolvedSourceSiteCount
			}
		case "fan-in":
			if left.InboundRefCount != right.InboundRefCount {
				return left.InboundRefCount > right.InboundRefCount
			}
		case "fan-out":
			if left.OutboundRefCount != right.OutboundRefCount {
				return left.OutboundRefCount > right.OutboundRefCount
			}
		case "symbols":
			if left.SymbolCount != right.SymbolCount {
				return left.SymbolCount > right.SymbolCount
			}
		case "flows":
			if left.LinkedFlowCount != right.LinkedFlowCount {
				return left.LinkedFlowCount > right.LinkedFlowCount
			}
		case "tests":
			if left.LinkedTestCount != right.LinkedTestCount {
				return left.LinkedTestCount > right.LinkedTestCount
			}
		}
		return left.Path < right.Path
	})
}

func normalizedSort(sortMode string) string {
	switch strings.ToLower(strings.TrimSpace(sortMode)) {
	case "unresolved", "fan-in", "fan-out", "symbols", "flows", "tests":
		return strings.ToLower(strings.TrimSpace(sortMode))
	default:
		return "path"
	}
}

func isAPIFile(summary FileSummary) bool {
	layer := strings.ToLower(summary.AppLayer)
	area := strings.ToLower(summary.FunctionalArea)
	return strings.Contains(layer, "api") || strings.Contains(area, "api") || strings.Contains(area, "mcp")
}

func riskFor(unresolved int) string {
	switch {
	case unresolved >= 10:
		return "high"
	case unresolved > 0:
		return "medium"
	default:
		return "low"
	}
}

func fileRelationshipGroups(samplesByFile map[string][]RelationshipSample, countsByFile map[string]map[string]int, sampleLimit int) []FileRelationshipGroup {
	files := make([]string, 0, len(samplesByFile))
	for file := range samplesByFile {
		files = append(files, file)
	}
	sort.Strings(files)

	groups := make([]FileRelationshipGroup, 0, len(files))
	for _, file := range files {
		samples := samplesByFile[file]
		sortSamples(samples)
		groups = append(groups, FileRelationshipGroup{
			File:    file,
			Total:   len(samples),
			Counts:  compactCounts(countsByFile[file]),
			Samples: limitSamples(samples, sampleLimit),
		})
	}
	return groups
}

func limitSamples(samples []RelationshipSample, limit int) []RelationshipSample {
	if limit < 0 || limit >= len(samples) {
		return append([]RelationshipSample(nil), samples...)
	}
	return append([]RelationshipSample(nil), samples[:limit]...)
}

func limitUnresolvedSamples(samples []UnresolvedSample, limit int) []UnresolvedSample {
	if limit < 0 || limit >= len(samples) {
		return append([]UnresolvedSample(nil), samples...)
	}
	return append([]UnresolvedSample(nil), samples[:limit]...)
}

func sortSamples(samples []RelationshipSample) {
	sort.SliceStable(samples, func(i, j int) bool {
		left, right := samples[i], samples[j]
		if left.SourceFile != right.SourceFile {
			return left.SourceFile < right.SourceFile
		}
		if left.TargetFile != right.TargetFile {
			return left.TargetFile < right.TargetFile
		}
		if left.SourceRange.StartLine != right.SourceRange.StartLine {
			return left.SourceRange.StartLine < right.SourceRange.StartLine
		}
		if left.RelationshipKind != right.RelationshipKind {
			return left.RelationshipKind < right.RelationshipKind
		}
		if left.SourceSymbol != right.SourceSymbol {
			return left.SourceSymbol < right.SourceSymbol
		}
		return left.TargetSymbol < right.TargetSymbol
	})
}

func sortUnresolvedSamples(samples []UnresolvedSample) {
	sort.SliceStable(samples, func(i, j int) bool {
		left, right := samples[i], samples[j]
		if left.Line != right.Line {
			return left.Line < right.Line
		}
		if left.Column != right.Column {
			return left.Column < right.Column
		}
		if left.SourceSymbol != right.SourceSymbol {
			return left.SourceSymbol < right.SourceSymbol
		}
		return left.TargetText < right.TargetText
	})
}

func sortSymbolIDs(ids []string, nodes map[string]graph.Node) {
	sort.SliceStable(ids, func(i, j int) bool {
		return compareSymbolNodes(nodes[ids[i]], nodes[ids[j]]) < 0
	})
}

func compareSymbolNodes(left graph.Node, right graph.Node) int {
	leftRange, rightRange := nodeRange(left), nodeRange(right)
	if leftRange.StartLine != rightRange.StartLine {
		return compareInt(leftRange.StartLine, rightRange.StartLine)
	}
	if leftRange.StartColumn != rightRange.StartColumn {
		return compareInt(leftRange.StartColumn, rightRange.StartColumn)
	}
	if left.Label != right.Label {
		return strings.Compare(string(left.Label), string(right.Label))
	}
	if nodeName(left) != nodeName(right) {
		return strings.Compare(nodeName(left), nodeName(right))
	}
	return strings.Compare(left.ID, right.ID)
}

func compareUnresolvedNodes(left graph.Node, right graph.Node) int {
	leftLine, rightLine := intProperty(left, "startLine"), intProperty(right, "startLine")
	if leftLine != rightLine {
		return compareInt(leftLine, rightLine)
	}
	leftCol, rightCol := intProperty(left, "startCol"), intProperty(right, "startCol")
	if leftCol != rightCol {
		return compareInt(leftCol, rightCol)
	}
	return strings.Compare(nodeName(left), nodeName(right))
}

func countExported(node SymbolTreeNode) int {
	count := 0
	if node.Exported {
		count++
	}
	for _, child := range node.Children {
		count += countExported(child)
	}
	return count
}

func increment(counts map[string]int, key string) {
	if key == "" {
		return
	}
	counts[key]++
}

func compactCounts(counts map[string]int) map[string]int {
	if len(counts) == 0 {
		return nil
	}
	out := make(map[string]int, len(counts))
	for key, count := range counts {
		if key != "" && count > 0 {
			out[key] = count
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func compareInt(left int, right int) int {
	switch {
	case left < right:
		return -1
	case left > right:
		return 1
	default:
		return 0
	}
}
