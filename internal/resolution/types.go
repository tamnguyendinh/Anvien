package resolution

import (
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

type Options struct {
	DisableScopeInheritsCompatibility bool
	SkipCompatibilityCrossFile        bool
	TypeScriptStandardLibrary         *tsstdlib.Authority
}

type Metrics struct {
	DefinitionsIndexed               int
	ImportsResolved                  int
	ImportUsesEmitted                int
	ResolvedReferences               int
	UnresolvedReferences             int
	UnresolvedReferenceDiagnostics   int
	UnattributedUnresolvedReferences int
	ResolvedCalls                    int
	ResolvedAccesses                 int
	ResolvedTypeReferences           int
	ResolvedExternalDeclarations     int
	ExternalCapabilityUnavailable    int
	ExternalProfileExcluded          int
	ExternalMeaningMismatches        int
	HeritageFactsIndexed             int
	ResolvedInheritance              int
	UnresolvedInheritance            int
	DuplicateEdgesMerged             int
	MethodOverridesEmitted           int
	MethodImplementsEmitted          int
	FinalizedImportsEmitted          int
	CrossFileFilesReprocessed        int
	CrossFileSkipped                 bool
	CrossFileSkipReason              string
	BindingAccumulatorFiles          int
	BindingAccumulatorEntries        int
	BindingAccumulatorFinalized      bool
	BindingAccumulatorDisposed       bool
	GraphNodesEmitted                int
	GraphRelationshipsEmitted        int
}

const TypeScriptStandardLibraryStage = "typescript_standard_library"

// TypeScriptAuthorityResult is the lossless P6-B authority result for one
// handled TypeScript source site. It is intentionally separate from graph
// materialization and the later cross-stage final outcome contract.
type TypeScriptAuthorityResult struct {
	SourceSiteID        string                     `json:"sourceSiteId"`
	Stage               string                     `json:"stage"`
	FilePath            string                     `json:"filePath"`
	FileHash            string                     `json:"fileHash,omitempty"`
	Range               scopeir.Range              `json:"range"`
	SiteKind            string                     `json:"siteKind"`
	RequestedName       string                     `json:"requestedName"`
	RequestedMeaning    tsstdlib.Meaning           `json:"requestedMeaning"`
	Status              tsstdlib.LookupStatus      `json:"status"`
	Reason              tsstdlib.Reason            `json:"reason,omitempty"`
	ResolvedSymbolID    string                     `json:"resolvedSymbolId,omitempty"`
	ResolvedOwnerID     string                     `json:"resolvedOwnerId,omitempty"`
	DeclarationRanges   []tsstdlib.Declaration     `json:"declarationRanges,omitempty"`
	AuthorityKind       string                     `json:"authorityKind"`
	CatalogProofState   tsstdlib.CatalogProofState `json:"catalogProofState"`
	AuthorityHash       string                     `json:"authorityHash"`
	TypeScriptVersion   string                     `json:"typescriptVersion"`
	CatalogHash         string                     `json:"catalogHash"`
	CatalogArtifactHash string                     `json:"catalogArtifactHash"`
	ProfileHash         string                     `json:"profileHash"`
	ConfigHash          string                     `json:"configHash"`
}

type Result struct {
	Graph                      *graph.Graph
	ReferenceIndex             ReferenceIndex
	TypeScriptAuthorityResults []TypeScriptAuthorityResult
	ResolutionOutcomes         []ResolutionOutcome
	Metrics                    Metrics
}

type BindingResult struct {
	workspace *workspace
	Metrics   Metrics
}

type ReferenceIndex struct {
	BySourceScope map[string][]Reference
	ByTargetDef   map[string][]Reference
}

func newReferenceIndex() ReferenceIndex {
	return ReferenceIndex{
		BySourceScope: make(map[string][]Reference),
		ByTargetDef:   make(map[string][]Reference),
	}
}

func (index ReferenceIndex) add(reference Reference) {
	index.BySourceScope[reference.FromScope] = append(index.BySourceScope[reference.FromScope], reference)
	index.ByTargetDef[reference.ToDefID] = append(index.ByTargetDef[reference.ToDefID], reference)
}

type ReferenceKind string

const (
	ReferenceCall          ReferenceKind = "call"
	ReferenceRead          ReferenceKind = "read"
	ReferenceWrite         ReferenceKind = "write"
	ReferenceTypeReference ReferenceKind = "type-reference"
	ReferenceInherits      ReferenceKind = "inherits"
	ReferenceImportUse     ReferenceKind = "import-use"
)

type Reference struct {
	FromScope        string
	ToDefID          string
	FilePath         string
	FileHash         string
	Range            scopeir.Range
	Kind             ReferenceKind
	Confidence       float64
	SourceSiteID     string
	SourceSiteStatus string
	ProofKind        string
	TargetRole       string
	TargetText       string
	Evidence         []graph.Evidence
}
