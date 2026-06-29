package filecontext

const (
	CompactFileContextFormat  = "file-detail.compact"
	CompactFileContextVersion = 1
)

type CompactFileContext struct {
	Format   string                   `json:"format"`
	Version  int                      `json:"version"`
	Repo     string                   `json:"repo,omitempty"`
	RepoPath string                   `json:"repoPath,omitempty"`
	Graph    GraphInfo                `json:"graph"`
	Target   Target                   `json:"target"`
	Schema   CompactFileContextSchema `json:"schema"`
	Summary  FileSummary              `json:"summary"`
	Dict     CompactDictionaries      `json:"dict"`
	Tables   CompactTables            `json:"tables"`
	Quality  QualitySignals           `json:"quality"`
	Limits   Limits                   `json:"limits"`
}

type CompactFileContextSchema struct {
	RangeTuple      []string `json:"rangeTuple"`
	SymbolRow       []string `json:"symbolRow"`
	RelationshipRow []string `json:"relationshipRow"`
	UnresolvedRow   []string `json:"unresolvedRow"`
	LinkedRow       []string `json:"linkedRow"`
}

type CompactDictionaries struct {
	Files              []string `json:"files"`
	Symbols            []string `json:"symbols"`
	SourceSites        []string `json:"sourceSites"`
	RelationshipKinds  []string `json:"relationshipKinds"`
	GapKinds           []string `json:"gapKinds"`
	Classifications    []string `json:"classifications"`
	Actionabilities    []string `json:"actionabilities"`
	ProofKinds         []string `json:"proofKinds"`
	SourceSiteStatuses []string `json:"sourceSiteStatuses"`
	LinkedKinds        []string `json:"linkedKinds"`
}

type CompactTables struct {
	Symbols       []CompactRow                `json:"symbols"`
	Relationships CompactRelationshipSections `json:"relationships"`
	Unresolved    CompactUnresolvedSummary    `json:"unresolved"`
	Linked        CompactLinkedSummary        `json:"linked"`
}

type CompactRelationshipSections struct {
	Counts         RelationshipCounts             `json:"counts"`
	Local          CompactRelationshipGroup       `json:"local"`
	OutboundByFile []CompactFileRelationshipGroup `json:"outboundByFile"`
	InboundByFile  []CompactFileRelationshipGroup `json:"inboundByFile"`
}

type CompactRelationshipGroup struct {
	Total  int            `json:"total"`
	Counts map[string]int `json:"counts,omitempty"`
	Rows   CompactRows    `json:"rows"`
}

type CompactFileRelationshipGroup struct {
	File   int            `json:"file"`
	Total  int            `json:"total"`
	Counts map[string]int `json:"counts,omitempty"`
	Rows   CompactRows    `json:"rows"`
}

type CompactUnresolvedSummary struct {
	Total            int                      `json:"total"`
	ByKind           map[string]int           `json:"byKind,omitempty"`
	ByClassification map[string]int           `json:"byClassification,omitempty"`
	ByActionability  map[string]int           `json:"byActionability,omitempty"`
	Groups           []CompactUnresolvedGroup `json:"groups"`
}

type CompactUnresolvedGroup struct {
	SourceSymbol any         `json:"sourceSymbol,omitempty"`
	Total        int         `json:"total"`
	Rows         CompactRows `json:"rows"`
}

type CompactLinkedSummary struct {
	Counts   LinkedCounts `json:"counts"`
	Flows    CompactRows  `json:"flows"`
	Routes   CompactRows  `json:"routes"`
	MCPTools CompactRows  `json:"mcpTools"`
	Tests    CompactRows  `json:"tests"`
}

type CompactRows struct {
	Total    int          `json:"total"`
	Returned int          `json:"returned"`
	Omitted  int          `json:"omitted"`
	Items    []CompactRow `json:"items"`
}

type CompactRow []any
type CompactRange [4]int

func CompactFileContextFromExpanded(context FileContext) CompactFileContext {
	builder := newCompactContextBuilder()
	builder.file(context.Summary.Path)
	builder.file(context.Target.NormalizedPath)

	tables := CompactTables{
		Symbols:       builder.symbolRows(context.SymbolTree, nil),
		Relationships: builder.relationshipSections(context.Relationships),
		Unresolved:    builder.unresolvedSummary(context.Unresolved),
		Linked:        builder.linkedSummary(context.Linked),
	}

	return CompactFileContext{
		Format:   CompactFileContextFormat,
		Version:  CompactFileContextVersion,
		Repo:     context.Repo,
		RepoPath: context.RepoPath,
		Graph:    context.Graph,
		Target:   context.Target,
		Schema:   defaultCompactFileContextSchema(),
		Summary:  context.Summary,
		Dict:     builder.dictionaries(),
		Tables:   tables,
		Quality:  context.Quality,
		Limits:   context.Limits,
	}
}

func defaultCompactFileContextSchema() CompactFileContextSchema {
	return CompactFileContextSchema{
		RangeTuple: []string{"startLine", "startColumn", "endLine", "endColumn"},
		SymbolRow: []string{
			"symbol", "parent", "name", "kind", "range", "exported", "signature",
			"local", "inbound", "outbound", "unresolved",
		},
		RelationshipRow: []string{
			"sourceFile", "sourceSymbol", "sourceRange", "relationshipKind",
			"targetFile", "targetSymbol", "targetRange", "sourceSite",
			"proofKind", "sourceSiteStatus",
		},
		UnresolvedRow: []string{
			"line", "column", "targetText", "sourceSymbol", "gapKind",
			"classification", "actionability", "proofKind", "sourceSite",
			"sourceSiteStatus",
		},
		LinkedRow: []string{"name", "kind", "source", "confidence", "trace"},
	}
}

type compactContextBuilder struct {
	files              *compactInterner
	symbols            *compactInterner
	sourceSites        *compactInterner
	relationshipKinds  *compactInterner
	gapKinds           *compactInterner
	classifications    *compactInterner
	actionabilities    *compactInterner
	proofKinds         *compactInterner
	sourceSiteStatuses *compactInterner
	linkedKinds        *compactInterner
}

func newCompactContextBuilder() *compactContextBuilder {
	return &compactContextBuilder{
		files:              newCompactInterner(),
		symbols:            newCompactInterner(),
		sourceSites:        newCompactInterner(),
		relationshipKinds:  newCompactInterner(),
		gapKinds:           newCompactInterner(),
		classifications:    newCompactInterner(),
		actionabilities:    newCompactInterner(),
		proofKinds:         newCompactInterner(),
		sourceSiteStatuses: newCompactInterner(),
		linkedKinds:        newCompactInterner(),
	}
}

func (b *compactContextBuilder) dictionaries() CompactDictionaries {
	return CompactDictionaries{
		Files:              b.files.values(),
		Symbols:            b.symbols.values(),
		SourceSites:        b.sourceSites.values(),
		RelationshipKinds:  b.relationshipKinds.values(),
		GapKinds:           b.gapKinds.values(),
		Classifications:    b.classifications.values(),
		Actionabilities:    b.actionabilities.values(),
		ProofKinds:         b.proofKinds.values(),
		SourceSiteStatuses: b.sourceSiteStatuses.values(),
		LinkedKinds:        b.linkedKinds.values(),
	}
}

func (b *compactContextBuilder) file(value string) any {
	return b.files.ref(value)
}

func (b *compactContextBuilder) symbol(value string) any {
	return b.symbols.ref(value)
}

func (b *compactContextBuilder) sourceSite(value string) any {
	return b.sourceSites.ref(value)
}

func (b *compactContextBuilder) relationshipKind(value string) any {
	return b.relationshipKinds.ref(value)
}

func (b *compactContextBuilder) gapKind(value string) any {
	return b.gapKinds.ref(value)
}

func (b *compactContextBuilder) classification(value string) any {
	return b.classifications.ref(value)
}

func (b *compactContextBuilder) actionability(value string) any {
	return b.actionabilities.ref(value)
}

func (b *compactContextBuilder) proofKind(value string) any {
	return b.proofKinds.ref(value)
}

func (b *compactContextBuilder) sourceSiteStatus(value string) any {
	return b.sourceSiteStatuses.ref(value)
}

func (b *compactContextBuilder) linkedKind(value string) any {
	return b.linkedKinds.ref(value)
}

func (b *compactContextBuilder) symbolRows(nodes []SymbolTreeNode, parent any) []CompactRow {
	rows := make([]CompactRow, 0, countSymbolTreeNodes(nodes))
	for _, node := range nodes {
		symbol := b.symbol(node.ID)
		row := CompactRow{
			symbol,
			parent,
			node.Name,
			node.Kind,
			compactRange(node.Range),
			node.Exported,
			node.Signature,
			node.RelationshipCounts.Local,
			node.RelationshipCounts.Inbound,
			node.RelationshipCounts.Outbound,
			node.RelationshipCounts.Unresolved,
		}
		rows = append(rows, row)
		rows = append(rows, b.symbolRows(node.Children, symbol)...)
	}
	return rows
}

func countSymbolTreeNodes(nodes []SymbolTreeNode) int {
	count := 0
	for _, node := range nodes {
		count++
		count += countSymbolTreeNodes(node.Children)
	}
	return count
}

func (b *compactContextBuilder) relationshipSections(sections RelationshipSections) CompactRelationshipSections {
	return CompactRelationshipSections{
		Counts:         sections.Counts,
		Local:          b.relationshipGroup(sections.Local),
		OutboundByFile: b.fileRelationshipGroups(sections.OutboundByFile),
		InboundByFile:  b.fileRelationshipGroups(sections.InboundByFile),
	}
}

func (b *compactContextBuilder) relationshipGroup(group RelationshipGroup) CompactRelationshipGroup {
	rows := b.relationshipRows(group.Samples)
	return CompactRelationshipGroup{
		Total:  group.Total,
		Counts: compactCounts(group.Counts),
		Rows:   compactRows(group.Total, rows),
	}
}

func (b *compactContextBuilder) fileRelationshipGroups(groups []FileRelationshipGroup) []CompactFileRelationshipGroup {
	out := make([]CompactFileRelationshipGroup, 0, len(groups))
	for _, group := range groups {
		rows := b.relationshipRows(group.Samples)
		fileRef := b.file(group.File)
		fileIndex, _ := fileRef.(int)
		out = append(out, CompactFileRelationshipGroup{
			File:   fileIndex,
			Total:  group.Total,
			Counts: compactCounts(group.Counts),
			Rows:   compactRows(group.Total, rows),
		})
	}
	return out
}

func (b *compactContextBuilder) relationshipRows(samples []RelationshipSample) []CompactRow {
	rows := make([]CompactRow, 0, len(samples))
	for _, sample := range samples {
		rows = append(rows, CompactRow{
			b.file(sample.SourceFile),
			b.symbol(sample.SourceSymbol),
			compactRange(sample.SourceRange),
			b.relationshipKind(sample.RelationshipKind),
			b.file(sample.TargetFile),
			b.symbol(sample.TargetSymbol),
			compactRange(sample.TargetRange),
			b.sourceSite(sample.SourceSiteID),
			b.proofKind(sample.ProofKind),
			b.sourceSiteStatus(sample.SourceSiteStatus),
		})
	}
	return rows
}

func (b *compactContextBuilder) unresolvedSummary(summary UnresolvedSummary) CompactUnresolvedSummary {
	groups := make([]CompactUnresolvedGroup, 0, len(summary.Groups))
	for _, group := range summary.Groups {
		rows := b.unresolvedRows(group.Samples)
		groups = append(groups, CompactUnresolvedGroup{
			SourceSymbol: b.symbol(group.SourceSymbol),
			Total:        group.Total,
			Rows:         compactRows(group.Total, rows),
		})
	}
	return CompactUnresolvedSummary{
		Total:            summary.Total,
		ByKind:           compactCounts(summary.ByKind),
		ByClassification: compactCounts(summary.ByClassification),
		ByActionability:  compactCounts(summary.ByActionability),
		Groups:           groups,
	}
}

func (b *compactContextBuilder) unresolvedRows(samples []UnresolvedSample) []CompactRow {
	rows := make([]CompactRow, 0, len(samples))
	for _, sample := range samples {
		rows = append(rows, CompactRow{
			sample.Line,
			sample.Column,
			sample.TargetText,
			b.symbol(sample.SourceSymbol),
			b.gapKind(sample.GapKind),
			b.classification(sample.Classification),
			b.actionability(sample.Actionability),
			b.proofKind(sample.ProofKind),
			b.sourceSite(sample.SourceSiteID),
			b.sourceSiteStatus(sample.SourceSiteStatus),
		})
	}
	return rows
}

func (b *compactContextBuilder) linkedSummary(summary LinkedSummary) CompactLinkedSummary {
	return CompactLinkedSummary{
		Counts:   summary.Counts,
		Flows:    compactRows(summary.Counts.Flows, b.linkedRows(summary.Flows)),
		Routes:   compactRows(summary.Counts.Routes, b.linkedRows(summary.Routes)),
		MCPTools: compactRows(summary.Counts.MCPTools, b.linkedRows(summary.MCPTools)),
		Tests:    compactRows(summary.Counts.Tests, b.linkedRows(summary.Tests)),
	}
}

func (b *compactContextBuilder) linkedRows(items []LinkedItem) []CompactRow {
	rows := make([]CompactRow, 0, len(items))
	for _, item := range items {
		rows = append(rows, CompactRow{
			item.Name,
			b.linkedKind(item.Kind),
			item.Source,
			item.Confidence,
			item.Trace,
		})
	}
	return rows
}

func compactRows(total int, rows []CompactRow) CompactRows {
	omitted := total - len(rows)
	if omitted < 0 {
		omitted = 0
	}
	return CompactRows{
		Total:    total,
		Returned: len(rows),
		Omitted:  omitted,
		Items:    rows,
	}
}

func compactRange(value SourceRange) CompactRange {
	return CompactRange{value.StartLine, value.StartColumn, value.EndLine, value.EndColumn}
}

type compactInterner struct {
	index map[string]int
	list  []string
}

func newCompactInterner() *compactInterner {
	return &compactInterner{index: map[string]int{}}
}

func (i *compactInterner) ref(value string) any {
	if value == "" {
		return nil
	}
	if existing, ok := i.index[value]; ok {
		return existing
	}
	next := len(i.list)
	i.index[value] = next
	i.list = append(i.list, value)
	return next
}

func (i *compactInterner) values() []string {
	return append([]string(nil), i.list...)
}
