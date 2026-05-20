package graphhealth

import "github.com/tamnguyendinh/avmatrix-go/internal/graph"

// EdgePolicy defines which relationship types contribute to counted connectivity
// for topology status (incoming/outgoing). This is the Phase 1 accepted policy.
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

// ExpectedIsolationReason values per Phase 1 policy.
const (
	ReasonTest            = "test"
	ReasonFixture         = "fixture"
	ReasonGenerated       = "generated"
	ReasonVendor          = "vendor"
	ReasonDocumentation   = "documentation"
	ReasonMigration       = "migration"
	ReasonExportedAPI     = "exported_api" // modifier only
	ReasonFrameworkEntry  = "framework_entry"
	ReasonCLIMCP          = "cli_mcp"
)

// TopologyStatus values (Phase 1 taxonomy).
type TopologyStatus string

const (
	TopologyConnected    TopologyStatus = "connected"
	TopologyTrueIsolated TopologyStatus = "true_isolated"
	TopologyNoIncoming   TopologyStatus = "no_incoming"
	TopologyNoOutgoing   TopologyStatus = "no_outgoing"
	TopologyDetached     TopologyStatus = "detached_component"
	TopologyUnknown      TopologyStatus = "unknown_connectivity"
)

// Confidence levels (Phase 1).
const (
	ConfidenceCandidate = "candidate"
	ConfidenceExpected  = "expected"
	ConfidenceUnknown   = "unknown"
	ConfidenceConfirmed = "confirmed"
)

// Diagnostic captures analyzer/resolution evidence attached to a node for health.
type Diagnostic struct {
	Kind   string `json:"kind"`
	Note   string `json:"note,omitempty"`
	Source string `json:"source,omitempty"`
}

// NodeHealth is the derived graph-health metadata attached to each node
// (populated into Node.Properties under "graphHealth" or flat keys for consumers).
type NodeHealth struct {
	TopologyStatus           TopologyStatus `json:"topologyStatus"`
	CountedIncoming          int            `json:"countedIncoming"`
	CountedOutgoing          int            `json:"countedOutgoing"`
	ExpectedIsolationReasons []string       `json:"expectedIsolationReasons,omitempty"`
	Diagnostics              []Diagnostic   `json:"diagnostics,omitempty"`
	Confidence               string         `json:"confidence"`
}
