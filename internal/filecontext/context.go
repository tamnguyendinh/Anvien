package filecontext

import (
	"fmt"
	"hash/fnv"
	"sort"
	"strings"
	"sync"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

const defaultSampleLimit = 10

const DerivedFileEdgesNote = "File relationship groups are projections derived from symbol and source-site graph facts; canonical graph relationships remain symbol/source-site facts."

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
	ChangedOnly         bool
	UnresolvedOnly      bool
	HighFanInOnly       bool
	HighFanOutOnly      bool
	HighFanInThreshold  int
	HighFanOutThreshold int
	ChangedPaths        map[string]struct{}
	Stale               bool
}

type FileList struct {
	Total  int           `json:"total"`
	Offset int           `json:"offset"`
	Limit  int           `json:"limit"`
	Sort   string        `json:"sort"`
	Files  []FileSummary `json:"files"`
}

type CacheKey struct {
	Repo      string
	RepoPath  string
	GraphPath string
	GraphHash string
}

type BuilderCache struct {
	mu       sync.Mutex
	builders map[CacheKey]*Builder
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

func AttachMetadata(context *FileContext, repoName string, repoPath string, graph GraphInfo) {
	if context == nil {
		return
	}
	context.Repo = repoName
	context.RepoPath = repoPath
	context.Graph = graph
	context.Quality.Stale = graph.Stale
	context.Quality.ChangedSinceAnalyze = graph.Stale
}

type FileSummary struct {
	Path                                    string `json:"path"`
	Language                                string `json:"language,omitempty"`
	Kind                                    string `json:"kind,omitempty"`
	FileRole                                string `json:"fileRole,omitempty"`
	AppLayer                                string `json:"appLayer,omitempty"`
	FunctionalArea                          string `json:"functionalArea,omitempty"`
	ParseStatus                             string `json:"parseStatus,omitempty"`
	SymbolCount                             int    `json:"symbolCount"`
	ExportedSymbolCount                     int    `json:"exportedSymbolCount"`
	InboundRefCount                         int    `json:"inboundRefCount"`
	OutboundRefCount                        int    `json:"outboundRefCount"`
	LocalRelationshipCount                  int    `json:"localRelationshipCount"`
	UnresolvedSourceSiteCount               int    `json:"unresolvedSourceSiteCount"`
	RawUnresolvedSourceSiteCount            int    `json:"rawUnresolvedSourceSiteCount"`
	ProductionUnresolvedSourceSiteCount     int    `json:"productionUnresolvedSourceSiteCount"`
	NonActionableUnresolvedSourceSiteCount  int    `json:"nonActionableUnresolvedSourceSiteCount"`
	UnknownUnresolvedSourceSiteCount        int    `json:"unknownUnresolvedSourceSiteCount"`
	DefaultVisibleUnresolvedSourceSiteCount int    `json:"defaultVisibleUnresolvedSourceSiteCount"`
	LinkedFlowCount                         int    `json:"linkedFlowCount"`
	LinkedTestCount                         int    `json:"linkedTestCount"`
	Risk                                    string `json:"risk,omitempty"`
	RawRisk                                 string `json:"rawRisk,omitempty"`
	DefaultVisibleRisk                      string `json:"defaultVisibleRisk,omitempty"`
	Stale                                   bool   `json:"stale"`
	ChangedSinceAnalyze                     bool   `json:"changedSinceAnalyze"`
}

type unresolvedBucketCounts struct {
	Raw            int
	Production     int
	NonActionable  int
	Unknown        int
	DefaultVisible int
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
	Counts   LinkedCounts `json:"counts"`
	Flows    []LinkedItem `json:"flows"`
	Routes   []LinkedItem `json:"routes"`
	MCPTools []LinkedItem `json:"mcpTools"`
	Tests    []LinkedItem `json:"tests"`
}

type LinkedCounts struct {
	Flows    int `json:"flows"`
	Routes   int `json:"routes"`
	MCPTools int `json:"mcpTools"`
	Tests    int `json:"tests"`
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

func NewBuilderCache() *BuilderCache {
	return &BuilderCache{builders: map[CacheKey]*Builder{}}
}

func (c *BuilderCache) Get(key CacheKey, g *graph.Graph) (*Builder, bool) {
	if c == nil {
		return NewBuilder(g), false
	}
	key = normalizeCacheKey(key, g)
	c.mu.Lock()
	defer c.mu.Unlock()
	if builder, ok := c.builders[key]; ok {
		return builder, true
	}
	builder := NewBuilder(g)
	c.builders[key] = builder
	return builder, false
}

func (c *BuilderCache) Invalidate(key CacheKey) {
	if c == nil {
		return
	}
	key = normalizeCacheKey(key, nil)
	c.mu.Lock()
	defer c.mu.Unlock()
	for existing := range c.builders {
		if cacheKeyMatches(key, existing) {
			delete(c.builders, existing)
		}
	}
}

func (c *BuilderCache) Clear() {
	if c == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.builders = map[CacheKey]*Builder{}
}

func (c *BuilderCache) Len() int {
	if c == nil {
		return 0
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.builders)
}

func GraphFingerprint(g *graph.Graph) string {
	if g == nil {
		return "nil"
	}
	hasher := fnv.New64a()
	writeHash(hasher, fmt.Sprintf("nodes=%d;relationships=%d;", len(g.Nodes), len(g.Relationships)))
	for _, node := range g.Nodes {
		writeHash(hasher, node.ID)
		writeHash(hasher, string(node.Label))
		writeHash(hasher, stringProperty(node, "filePath"))
		writeHash(hasher, stringProperty(node, "name"))
	}
	for _, relationship := range g.Relationships {
		writeHash(hasher, relationship.ID)
		writeHash(hasher, relationship.SourceID)
		writeHash(hasher, relationship.TargetID)
		writeHash(hasher, string(relationship.Type))
		writeHash(hasher, relationship.SourceSiteID)
		writeHash(hasher, relationship.FilePath)
	}
	return fmt.Sprintf("%016x", hasher.Sum64())
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
	linked := b.buildLinked(normalizedPath, limits.LinkedSamplesPerKind)
	kind := fileKind(fileNode)
	appLayer := stringProperty(fileNode, "appLayer")
	functionalArea := stringProperty(fileNode, "functionalArea")
	fileRole := semantic.ClassifyFileRole(normalizedPath, kind, appLayer, functionalArea)

	summary := FileSummary{
		Path:                   normalizedPath,
		Language:               stringProperty(fileNode, "language"),
		Kind:                   kind,
		FileRole:               string(fileRole.Role),
		AppLayer:               appLayer,
		FunctionalArea:         functionalArea,
		ParseStatus:            parseStatus(fileNode),
		SymbolCount:            symbolCount,
		ExportedSymbolCount:    exportedCount,
		InboundRefCount:        relationships.Counts.Inbound,
		OutboundRefCount:       relationships.Counts.Outbound,
		LocalRelationshipCount: relationships.Counts.Local,
		LinkedFlowCount:        linked.Counts.Flows,
		LinkedTestCount:        linked.Counts.Tests,
	}
	applyUnresolvedBuckets(&summary, b.unresolvedBucketCounts(normalizedPath, kind))

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
	attachFileListQuality(summaries, options)
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
	linkedFlowsByFile := map[string]map[string]struct{}{}
	linkedTestsByFile := map[string]map[string]struct{}{}
	for path, node := range b.filesByPath {
		kind := fileKind(node)
		appLayer := stringProperty(node, "appLayer")
		functionalArea := stringProperty(node, "functionalArea")
		fileRole := semantic.ClassifyFileRole(path, kind, appLayer, functionalArea)
		summary := &FileSummary{
			Path:           path,
			Language:       stringProperty(node, "language"),
			Kind:           kind,
			FileRole:       string(fileRole.Role),
			AppLayer:       appLayer,
			FunctionalArea: functionalArea,
			ParseStatus:    parseStatus(node),
			SymbolCount:    len(b.definesByFile[path]),
		}
		for _, symbolID := range b.definesByFile[path] {
			if exportedSymbol(b.nodesByID[symbolID]) {
				summary.ExportedSymbolCount++
			}
		}
		applyUnresolvedBuckets(summary, b.unresolvedBucketCounts(path, kind))
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
		if sourceFile != "" && relationship.Type == graph.RelStepInProcess && targetNode.Label == scopeir.NodeProcess {
			addStringSet(linkedFlowsByFile, sourceFile, targetNode.ID)
		}
		if sourceFile != "" && targetFile != "" && sourceFile != targetFile && b.isTestFile(sourceFile) {
			addStringSet(linkedTestsByFile, targetFile, sourceFile)
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
	for path, flows := range linkedFlowsByFile {
		if summary := aggregates[path]; summary != nil {
			summary.LinkedFlowCount = len(flows)
		}
	}
	for path, tests := range linkedTestsByFile {
		if summary := aggregates[path]; summary != nil {
			summary.LinkedTestCount = len(tests)
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

func applyUnresolvedBuckets(summary *FileSummary, counts unresolvedBucketCounts) {
	summary.UnresolvedSourceSiteCount = counts.Raw
	summary.RawUnresolvedSourceSiteCount = counts.Raw
	summary.ProductionUnresolvedSourceSiteCount = counts.Production
	summary.NonActionableUnresolvedSourceSiteCount = counts.NonActionable
	summary.UnknownUnresolvedSourceSiteCount = counts.Unknown
	summary.DefaultVisibleUnresolvedSourceSiteCount = counts.DefaultVisible
	summary.RawRisk = riskFor(counts.Raw)
	summary.DefaultVisibleRisk = riskFor(counts.DefaultVisible)
	summary.Risk = summary.DefaultVisibleRisk
}

func (b *Builder) unresolvedBucketCounts(path string, kind string) unresolvedBucketCounts {
	nodes := b.unresolvedByFile[path]
	counts := unresolvedBucketCounts{Raw: len(nodes)}
	if counts.Raw == 0 {
		return counts
	}
	if kind == "test" {
		return counts
	}
	for _, node := range nodes {
		switch {
		case isNonActionableUnresolved(node):
			counts.NonActionable++
		case isUnknownUnresolved(node):
			counts.Unknown++
			counts.DefaultVisible++
		default:
			counts.Production++
			counts.DefaultVisible++
		}
	}
	return counts
}

func isNonActionableUnresolved(node graph.Node) bool {
	return normalizedUnresolvedProperty(node, "actionability") == "non_actionable"
}

func isUnknownUnresolved(node graph.Node) bool {
	actionability := normalizedUnresolvedProperty(node, "actionability")
	classification := normalizedUnresolvedProperty(node, "classification")
	return actionability == "" || actionability == "unknown" || classification == "" || classification == "unknown"
}

func normalizedUnresolvedProperty(node graph.Node, key string) string {
	return strings.ToLower(strings.TrimSpace(stringProperty(node, key)))
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

func (b *Builder) buildLinked(path string, sampleLimit int) LinkedSummary {
	flows := map[string]LinkedItem{}
	routes := map[string]LinkedItem{}
	mcpTools := map[string]LinkedItem{}
	tests := map[string]LinkedItem{}

	for _, relationship := range b.g.Relationships {
		sourceNode, sourceOK := b.nodesByID[relationship.SourceID]
		targetNode, targetOK := b.nodesByID[relationship.TargetID]
		if !sourceOK || !targetOK {
			continue
		}
		sourceFile := b.nodeFilePath(sourceNode)
		if sourceFile == "" {
			sourceFile = normalizePath(relationship.FilePath)
		}
		targetFile := b.nodeFilePath(targetNode)
		switch relationship.Type {
		case graph.RelStepInProcess:
			if sourceFile == path && targetNode.Label == scopeir.NodeProcess {
				putLinkedItem(flows, targetNode.ID, linkedItem(targetNode, "flow", relationship, sourceNode, targetNode))
			}
		case graph.RelHandlesRoute:
			if sourceFile == path || targetFile == path {
				if item, ok := routeOrToolLinkedItem(sourceNode, targetNode, scopeir.NodeRoute, "route", relationship); ok {
					putLinkedItem(routes, item.Name, item)
				}
			}
		case graph.RelHandlesTool:
			if sourceFile == path || targetFile == path {
				if item, ok := routeOrToolLinkedItem(sourceNode, targetNode, scopeir.NodeTool, "mcp_tool", relationship); ok {
					putLinkedItem(mcpTools, item.Name, item)
				}
			}
		default:
			if targetFile == path && sourceFile != "" && sourceFile != path && b.isTestFile(sourceFile) {
				item := LinkedItem{
					Name:       sourceFile,
					Kind:       "test",
					Source:     string(relationship.Type),
					Confidence: relationshipConfidence(relationship),
					Trace:      relationship.SourceID + " -> " + relationship.TargetID,
				}
				putLinkedItem(tests, sourceFile, item)
			}
		}
	}

	return LinkedSummary{
		Counts: LinkedCounts{
			Flows:    len(flows),
			Routes:   len(routes),
			MCPTools: len(mcpTools),
			Tests:    len(tests),
		},
		Flows:    sortedLinkedItems(flows, sampleLimit),
		Routes:   sortedLinkedItems(routes, sampleLimit),
		MCPTools: sortedLinkedItems(mcpTools, sampleLimit),
		Tests:    sortedLinkedItems(tests, sampleLimit),
	}
}

func (b *Builder) isTestFile(path string) bool {
	if node, ok := b.filesByPath[path]; ok {
		return fileKind(node) == "test"
	}
	lower := strings.ToLower(path)
	return strings.Contains(lower, "_test.") || strings.Contains(lower, ".test.") || strings.Contains(lower, ".spec.")
}

func routeOrToolLinkedItem(sourceNode graph.Node, targetNode graph.Node, linkedLabel scopeir.NodeLabel, kind string, relationship graph.Relationship) (LinkedItem, bool) {
	switch {
	case targetNode.Label == linkedLabel:
		return linkedItem(targetNode, kind, relationship, sourceNode, targetNode), true
	case sourceNode.Label == linkedLabel:
		return linkedItem(sourceNode, kind, relationship, sourceNode, targetNode), true
	default:
		return LinkedItem{}, false
	}
}

func linkedItem(node graph.Node, kind string, relationship graph.Relationship, sourceNode graph.Node, targetNode graph.Node) LinkedItem {
	return LinkedItem{
		Name:       nodeName(node),
		Kind:       kind,
		Source:     string(relationship.Type),
		Confidence: relationshipConfidence(relationship),
		Trace:      sourceNode.ID + " -> " + targetNode.ID,
	}
}

func relationshipConfidence(relationship graph.Relationship) string {
	if relationship.Confidence <= 0 {
		return ""
	}
	return fmt.Sprintf("%.2f", relationship.Confidence)
}

func putLinkedItem(items map[string]LinkedItem, key string, item LinkedItem) {
	if key == "" {
		key = item.Name
	}
	if key == "" {
		return
	}
	if _, exists := items[key]; exists {
		return
	}
	items[key] = item
}

func addStringSet(groups map[string]map[string]struct{}, group string, value string) {
	if group == "" || value == "" {
		return
	}
	if groups[group] == nil {
		groups[group] = map[string]struct{}{}
	}
	groups[group][value] = struct{}{}
}

func sortedLinkedItems(items map[string]LinkedItem, limit int) []LinkedItem {
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		left, right := items[keys[i]], items[keys[j]]
		if left.Kind != right.Kind {
			return left.Kind < right.Kind
		}
		if left.Name != right.Name {
			return left.Name < right.Name
		}
		return keys[i] < keys[j]
	})
	if limit >= 0 && limit < len(keys) {
		keys = keys[:limit]
	}
	out := make([]LinkedItem, 0, len(keys))
	for _, key := range keys {
		out = append(out, items[key])
	}
	return out
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

func normalizeCacheKey(key CacheKey, g *graph.Graph) CacheKey {
	key.Repo = strings.TrimSpace(key.Repo)
	key.RepoPath = normalizePath(key.RepoPath)
	key.GraphPath = normalizePath(key.GraphPath)
	key.GraphHash = strings.TrimSpace(key.GraphHash)
	if key.GraphHash == "" && g != nil {
		key.GraphHash = GraphFingerprint(g)
	}
	return key
}

func cacheKeyMatches(pattern CacheKey, existing CacheKey) bool {
	if pattern.Repo != "" && pattern.Repo != existing.Repo {
		return false
	}
	if pattern.RepoPath != "" && pattern.RepoPath != existing.RepoPath {
		return false
	}
	if pattern.GraphPath != "" && pattern.GraphPath != existing.GraphPath {
		return false
	}
	if pattern.GraphHash != "" && pattern.GraphHash != existing.GraphHash {
		return false
	}
	return true
}

func normalizePathFromNodeID(id string) string {
	return normalizePath(strings.TrimPrefix(id, "File:"))
}

type hashWriter interface {
	Write([]byte) (int, error)
}

func writeHash(hasher hashWriter, value string) {
	_, _ = hasher.Write([]byte(value))
	_, _ = hasher.Write([]byte{0})
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
		if options.ChangedOnly && !summary.ChangedSinceAnalyze {
			continue
		}
		if options.UnresolvedOnly && unresolvedFilterCount(summary, normalizedSort(options.Sort)) == 0 {
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

func attachFileListQuality(summaries []FileSummary, options FileListOptions) {
	changedPaths := normalizeChangedPaths(options.ChangedPaths)
	for i := range summaries {
		summaries[i].Stale = options.Stale
		summaries[i].ChangedSinceAnalyze = pathMatchesSet(summaries[i].Path, changedPaths)
	}
}

func normalizeChangedPaths(paths map[string]struct{}) map[string]struct{} {
	if len(paths) == 0 {
		return nil
	}
	normalized := make(map[string]struct{}, len(paths))
	for path := range paths {
		if normalizedPath := normalizePath(path); normalizedPath != "" {
			normalized[normalizedPath] = struct{}{}
		}
	}
	return normalized
}

func pathMatchesSet(path string, paths map[string]struct{}) bool {
	if len(paths) == 0 {
		return false
	}
	normalized := normalizePath(path)
	if _, ok := paths[normalized]; ok {
		return true
	}
	for changedPath := range paths {
		if strings.HasSuffix(normalized, "/"+changedPath) || strings.HasSuffix(changedPath, "/"+normalized) {
			return true
		}
	}
	return false
}

func sortSummaries(summaries []FileSummary, sortMode string) {
	mode := normalizedSort(sortMode)
	sort.SliceStable(summaries, func(i, j int) bool {
		left, right := summaries[i], summaries[j]
		switch mode {
		case "unresolved", "raw-unresolved", "production-unresolved":
			leftCount := unresolvedSortCount(left, mode)
			rightCount := unresolvedSortCount(right, mode)
			if leftCount != rightCount {
				return leftCount > rightCount
			}
			if left.RawUnresolvedSourceSiteCount != right.RawUnresolvedSourceSiteCount {
				return left.RawUnresolvedSourceSiteCount > right.RawUnresolvedSourceSiteCount
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

func unresolvedFilterCount(summary FileSummary, sortMode string) int {
	switch sortMode {
	case "raw-unresolved", "production-unresolved":
		return unresolvedSortCount(summary, sortMode)
	default:
		return summary.DefaultVisibleUnresolvedSourceSiteCount
	}
}

func unresolvedSortCount(summary FileSummary, sortMode string) int {
	switch sortMode {
	case "raw-unresolved":
		return summary.RawUnresolvedSourceSiteCount
	case "production-unresolved":
		return summary.ProductionUnresolvedSourceSiteCount
	default:
		return summary.DefaultVisibleUnresolvedSourceSiteCount
	}
}

func normalizedSort(sortMode string) string {
	if !isSupportedSort(sortMode) {
		return "path"
	}
	switch strings.ToLower(strings.TrimSpace(sortMode)) {
	case "unresolved", "default-unresolved", "default-visible-unresolved":
		return "unresolved"
	case "raw", "raw-unresolved":
		return "raw-unresolved"
	case "production-unresolved", "actionable-unresolved":
		return "production-unresolved"
	case "fan-in", "fan-out", "symbols", "flows", "tests":
		return strings.ToLower(strings.TrimSpace(sortMode))
	default:
		return "path"
	}
}

func NormalizeFileListSort(sortMode string) string {
	return normalizedSort(sortMode)
}

func IsSupportedFileListSort(sortMode string) bool {
	return isSupportedSort(sortMode)
}

func SupportedFileListSorts() []string {
	return []string{"path", "unresolved", "raw-unresolved", "production-unresolved", "fan-in", "fan-out", "symbols", "flows", "tests"}
}

func isSupportedSort(sortMode string) bool {
	switch strings.ToLower(strings.TrimSpace(sortMode)) {
	case "", "path", "unresolved", "default-unresolved", "default-visible-unresolved", "raw", "raw-unresolved", "production-unresolved", "actionable-unresolved", "fan-in", "fan-out", "symbols", "flows", "tests":
		return true
	default:
		return false
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
