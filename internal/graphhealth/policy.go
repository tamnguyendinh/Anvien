package graphhealth

import "github.com/tamnguyendinh/avmatrix-go/internal/graph"

// EdgePolicy defines which relationship types contribute to counted connectivity
// for topology status (incoming/outgoing).
const PolicyVersion = "graph-health-non-structural-v1"

var CountedEdgeTypes = map[graph.RelationshipType]bool{
	graph.RelCalls:            true,
	graph.RelAccesses:         true,
	graph.RelInherits:         true,
	graph.RelImplements:       true,
	graph.RelExtends:          true,
	graph.RelMethodOverrides:  true,
	graph.RelMethodImplements: true,
	graph.RelImports:          true,
	graph.RelUses:             true,
	graph.RelDecorates:        true,
	graph.RelWraps:            true,
	graph.RelQueries:          true,
	graph.RelFetches:          true,
	graph.RelStepInProcess:    true,
	graph.RelHandlesRoute:     true,
	graph.RelHandlesTool:      true,
	graph.RelEntryPointOf:     true,
}

// IsCounted reports whether the relationship type contributes to connectivity counts.
func IsCounted(t graph.RelationshipType) bool {
	return CountedEdgeTypes[t]
}

// StructuralEdgeTypes are excluded from connectivity counts (ownership/containment).
var StructuralEdgeTypes = map[graph.RelationshipType]bool{
	graph.RelContains:    true,
	graph.RelDefines:     true,
	graph.RelHasMethod:   true,
	graph.RelHasProperty: true,
	graph.RelMemberOf:    true,
}

const (
	ExcludedEdgeStructural = "structural"
	ExcludedEdgeOther      = "other"
)

// ExpectedIsolationReason values per accepted graph-health policy.
const (
	ReasonTest           = "test"
	ReasonFixture        = "fixture"
	ReasonGenerated      = "generated"
	ReasonVendor         = "vendor"
	ReasonDocumentation  = "documentation"
	ReasonMigration      = "migration"
	ReasonExportedAPI    = "exported_api" // modifier only
	ReasonFrameworkEntry = "framework_entry"
	ReasonCLIMCP         = "cli_mcp"
)

var ExpectedIsolationReasons = []string{
	ReasonTest,
	ReasonFixture,
	ReasonGenerated,
	ReasonVendor,
	ReasonDocumentation,
	ReasonMigration,
	ReasonExportedAPI,
	ReasonFrameworkEntry,
	ReasonCLIMCP,
}

// TopologyStatus values.
type TopologyStatus string

const (
	TopologyConnected    TopologyStatus = "connected"
	TopologyTrueIsolated TopologyStatus = "true_isolated"
	TopologyNoIncoming   TopologyStatus = "no_incoming"
	TopologyNoOutgoing   TopologyStatus = "no_outgoing"
	TopologyDetached     TopologyStatus = "detached_component"
	TopologyUnknown      TopologyStatus = "unknown_connectivity"
)

var TopologyStatuses = []TopologyStatus{
	TopologyConnected,
	TopologyTrueIsolated,
	TopologyNoIncoming,
	TopologyNoOutgoing,
	TopologyDetached,
	TopologyUnknown,
}

// Confidence levels.
const (
	ConfidenceCandidate = "candidate"
	ConfidenceExpected  = "expected"
	ConfidenceUnknown   = "unknown"
	ConfidenceConfirmed = "confirmed"
)

var ConfidenceLevels = []string{
	ConfidenceCandidate,
	ConfidenceExpected,
	ConfidenceUnknown,
	ConfidenceConfirmed,
}

const (
	DiagnosticUnresolvedReference = "unresolved_reference"
	DiagnosticPropertyKey         = "graphHealthDiagnostics"
)

const (
	DiagnosticClassificationBuiltin          = "builtin"
	DiagnosticClassificationStandardLibrary  = "standard_library"
	DiagnosticClassificationTestFramework    = "test_framework"
	DiagnosticClassificationExternalLibrary  = "external_library"
	DiagnosticClassificationInRepoUnresolved = "in_repo_unresolved"
	DiagnosticClassificationUnclassified     = "unclassified"
)

var DiagnosticClassifications = []string{
	DiagnosticClassificationBuiltin,
	DiagnosticClassificationStandardLibrary,
	DiagnosticClassificationTestFramework,
	DiagnosticClassificationExternalLibrary,
	DiagnosticClassificationInRepoUnresolved,
	DiagnosticClassificationUnclassified,
}

const (
	DiagnosticActionabilityNonActionable = "non_actionable"
	DiagnosticActionabilityReview        = "review"
	DiagnosticActionabilityAnalyzerGap   = "analyzer_gap"
)

var DiagnosticActionabilities = []string{
	DiagnosticActionabilityNonActionable,
	DiagnosticActionabilityReview,
	DiagnosticActionabilityAnalyzerGap,
}

// Diagnostic captures analyzer/resolution evidence attached to a node for health.
type Diagnostic struct {
	Kind             string `json:"kind"`
	FactFamily       string `json:"factFamily,omitempty"`
	SourceNodeID     string `json:"sourceNodeId,omitempty"`
	TargetText       string `json:"targetText,omitempty"`
	ResolutionSource string `json:"resolutionSource,omitempty"`
	Classification   string `json:"classification,omitempty"`
	Actionability    string `json:"actionability,omitempty"`
	FilePath         string `json:"filePath,omitempty"`
	FileHash         string `json:"fileHash,omitempty"`
	StartLine        int    `json:"startLine,omitempty"`
	StartCol         int    `json:"startCol,omitempty"`
	EndLine          int    `json:"endLine,omitempty"`
	EndCol           int    `json:"endCol,omitempty"`
	SourceSiteID     string `json:"sourceSiteId,omitempty"`
	SourceSiteStatus string `json:"sourceSiteStatus,omitempty"`
	ProofKind        string `json:"proofKind,omitempty"`
	TargetRole       string `json:"targetRole,omitempty"`
	Count            int    `json:"count,omitempty"`
	Note             string `json:"note,omitempty"`
	Source           string `json:"source,omitempty"`
}

// NodeHealth is the derived graph-health metadata attached to each node
// (populated into Node.Properties under "graphHealth" or flat keys for consumers).
type NodeHealth struct {
	TopologyStatus             TopologyStatus `json:"topologyStatus"`
	CountedIncoming            int            `json:"countedIncoming"`
	CountedOutgoing            int            `json:"countedOutgoing"`
	ExcludedEdgeCounts         map[string]int `json:"excludedEdgeCounts,omitempty"`
	ComponentID                string         `json:"componentId,omitempty"`
	ComponentSize              int            `json:"componentSize,omitempty"`
	ComponentRootNodeIDs       []string       `json:"componentRootNodeIds,omitempty"`
	ComponentReachableFromRoot bool           `json:"componentReachableFromRoot"`
	ExpectedIsolationReasons   []string       `json:"expectedIsolationReasons,omitempty"`
	Diagnostics                []Diagnostic   `json:"diagnostics,omitempty"`
	Confidence                 string         `json:"confidence"`
}

// ComponentSummary captures component-level graph-health explanation data.
type ComponentSummary struct {
	ID                string   `json:"id"`
	NodeCount         int      `json:"nodeCount"`
	CountedEdgeCount  int      `json:"countedEdgeCount"`
	Detached          bool     `json:"detached"`
	ReachableFromRoot bool     `json:"reachableFromRoot"`
	RootNodeIDs       []string `json:"rootNodeIds,omitempty"`
	SampleNodeIDs     []string `json:"sampleNodeIds,omitempty"`
}

// Summary captures graph-level inventory for consumer surfaces.
type Summary struct {
	PolicyVersion                        string             `json:"policyVersion"`
	NodeCount                            int                `json:"nodeCount"`
	CountedRelationshipCount             int                `json:"countedRelationshipCount"`
	ComponentCount                       int                `json:"componentCount"`
	DetachedComponentCount               int                `json:"detachedComponentCount"`
	RootNodeCount                        int                `json:"rootNodeCount"`
	UnresolvedReferenceCount             int                `json:"unresolvedReferenceCount"`
	SourceBackedUnresolvedReferenceCount int                `json:"sourceBackedUnresolvedReferenceCount"`
	UnattributedUnresolvedReferenceCount int                `json:"unattributedUnresolvedReferenceCount"`
	TopologyStatusCounts                 map[string]int     `json:"topologyStatusCounts"`
	ExpectedIsolationReasonCounts        map[string]int     `json:"expectedIsolationReasonCounts"`
	ConfidenceCounts                     map[string]int     `json:"confidenceCounts"`
	DiagnosticCounts                     map[string]int     `json:"diagnosticCounts"`
	DiagnosticClassificationCounts       map[string]int     `json:"diagnosticClassificationCounts"`
	DiagnosticActionabilityCounts        map[string]int     `json:"diagnosticActionabilityCounts"`
	ExcludedEdgeCounts                   map[string]int     `json:"excludedEdgeCounts"`
	LargestDetachedComponents            []ComponentSummary `json:"largestDetachedComponents,omitempty"`
}
