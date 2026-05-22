import {
  GRAPH_HEALTH_CONFIDENCE_LEVELS,
  GRAPH_HEALTH_EXPECTED_ISOLATION_REASONS,
  GRAPH_HEALTH_TOPOLOGY_STATUSES,
  type GraphHealthDiagnostic,
  type GraphHealthDiagnosticActionability,
  type GraphHealthDiagnosticClassification,
  type GraphHealthConfidence,
  type GraphHealthExpectedIsolationReason,
  type GraphHealthNodeMetadata,
  type GraphHealthTopologyStatus,
  type GraphNode,
} from "@/generated/avmatrix-contracts";

export const GRAPH_HEALTH_DIAGNOSTIC_KINDS = ["unresolved_reference"] as const;

export type GraphHealthDiagnosticKind =
  | (typeof GRAPH_HEALTH_DIAGNOSTIC_KINDS)[number]
  | string;

export interface GraphHealthFilterState {
  visibleTopologyStatuses: GraphHealthTopologyStatus[];
  hiddenExpectedIsolationReasons: GraphHealthExpectedIsolationReason[];
  visibleDiagnosticKinds: GraphHealthDiagnosticKind[];
}

export interface GraphHealthFilterable {
  topologyStatus?: GraphHealthTopologyStatus;
  expectedIsolationReasons?: GraphHealthExpectedIsolationReason[];
  diagnostics?: GraphHealthDiagnostic[];
}

export const DEFAULT_GRAPH_HEALTH_FILTERS: GraphHealthFilterState = {
  visibleTopologyStatuses: [...GRAPH_HEALTH_TOPOLOGY_STATUSES],
  hiddenExpectedIsolationReasons: [],
  visibleDiagnosticKinds: [...GRAPH_HEALTH_DIAGNOSTIC_KINDS],
};

export const GRAPH_HEALTH_TOPOLOGY_LABELS: Record<
  GraphHealthTopologyStatus,
  string
> = {
  connected: "Connected",
  true_isolated: "True isolated",
  no_incoming: "No incoming",
  no_outgoing: "No outgoing",
  detached_component: "Detached component",
  unknown_connectivity: "Unknown",
};

export const GRAPH_HEALTH_TOPOLOGY_DESCRIPTIONS: Record<
  GraphHealthTopologyStatus,
  string
> = {
  connected:
    "Has counted incoming and outgoing wiring under the Graph Health edge policy.",
  true_isolated:
    "Has no counted incoming or outgoing wiring; this is a triage candidate, not a verdict.",
  no_incoming:
    "Uses other nodes but no counted in-repo wiring reaches it; inspect as an unwired candidate.",
  no_outgoing:
    "Is reached by other nodes but has no counted outgoing wiring; often normal leaf behavior.",
  detached_component:
    "Belongs to a counted-edge component that no accepted root reaches.",
  unknown_connectivity:
    "Topology could not be classified from graph metadata; valid computed graph nodes should rarely use this status.",
};

export const GRAPH_HEALTH_REASON_LABELS: Record<
  GraphHealthExpectedIsolationReason,
  string
> = {
  test: "Test",
  fixture: "Fixture",
  generated: "Generated",
  vendor: "Vendor",
  documentation: "Documentation",
  migration: "Migration",
  exported_api: "Exported API",
  framework_entry: "Framework entry",
  cli_mcp: "CLI/MCP",
};

export const GRAPH_HEALTH_REASON_DESCRIPTIONS: Record<
  GraphHealthExpectedIsolationReason,
  string
> = {
  test: "Expected-isolated overlay from test file or test-helper evidence.",
  fixture:
    "Expected-isolated overlay from fixture, snapshot, or testdata path evidence.",
  generated:
    "Expected-isolated overlay from generated-code path or metadata evidence.",
  vendor:
    "Expected-isolated overlay from vendor/dependency/build-output path evidence.",
  documentation:
    "Expected-isolated overlay from documentation or Section-node evidence.",
  migration:
    "Expected-isolated overlay from migration or database-change script evidence.",
  exported_api:
    "Public/exported surface modifier; lower priority, not automatically safe to hide.",
  framework_entry:
    "Accepted root or framework entry surface, not a dead-code verdict.",
  cli_mcp: "CLI, MCP, tool, or command surface that may be reached externally.",
};

export const GRAPH_HEALTH_DIAGNOSTIC_LABELS: Record<string, string> = {
  unresolved_reference: "Unresolved reference",
};

export const GRAPH_HEALTH_DIAGNOSTIC_DESCRIPTIONS: Record<string, string> = {
  unresolved_reference:
    "Source-backed unresolved reference; indicates analyzer or dependency uncertainty.",
};

export const GRAPH_HEALTH_DIAGNOSTIC_CLASSIFICATION_LABELS: Record<
  GraphHealthDiagnosticClassification,
  string
> = {
  builtin: "Builtin",
  standard_library: "Standard library",
  test_framework: "Test framework",
  external_library: "External library",
  in_repo_unresolved: "In-repo unresolved",
  unclassified: "Unclassified",
};

export const GRAPH_HEALTH_DIAGNOSTIC_ACTIONABILITY_LABELS: Record<
  GraphHealthDiagnosticActionability,
  string
> = {
  non_actionable: "Non-actionable",
  review: "Review",
  analyzer_gap: "Analyzer gap",
};

export const GRAPH_HEALTH_CONFIDENCE_LABELS: Record<
  GraphHealthConfidence,
  string
> = {
  candidate: "Candidate",
  expected: "Expected",
  unknown: "Unknown",
  confirmed: "Confirmed",
};

export const GRAPH_HEALTH_CONFIDENCE_DESCRIPTIONS: Record<
  GraphHealthConfidence,
  string
> = {
  candidate: "Actionable triage candidate with no expected-isolated overlay.",
  expected:
    "Explained by expected-isolated policy such as test, documentation, fixture, or entry surface.",
  unknown:
    "Resolution or analyzer uncertainty is present; inspect evidence before triage.",
  confirmed:
    "Reserved for external evidence or human review; derivation does not auto-confirm defects.",
};

export const getGraphHealthNextAction = (
  health: Pick<
    GraphHealthNodeMetadata,
    "topologyStatus" | "expectedIsolationReasons" | "diagnostics" | "confidence"
  >,
): string => {
  if (health.topologyStatus === "unknown_connectivity") {
    return "Verify graph payload completeness before using this topology status for triage.";
  }
  if (
    health.confidence === "expected" ||
    (health.expectedIsolationReasons?.length ?? 0) > 0
  ) {
    return "Review only if this expected-isolated overlay looks wrong for the node.";
  }
  if (
    health.confidence === "unknown" ||
    (health.diagnostics?.length ?? 0) > 0
  ) {
    return "Inspect diagnostic evidence separately; topology remains based on counted graph wiring.";
  }
  switch (health.topologyStatus) {
    case "no_incoming":
      return "Check routes, exports, runtime entrypoints, and external callers before treating it as unwired.";
    case "detached_component":
      return "Inspect root reachability and missing registration edges for this component.";
    case "true_isolated":
      return "Check path policy, generated/test context, and source references before pruning.";
    case "no_outgoing":
      return "Treat as a low-priority leaf unless it should create calls, accesses, or process edges.";
    case "connected":
      return "No Graph Health triage action is suggested from topology alone.";
    default:
      return "Use source, tests, routes, and runtime evidence before making a code-change decision.";
  }
};

export const getNodeGraphHealth = (
  node: GraphNode,
): GraphHealthNodeMetadata | null => {
  const structured = node.properties.graphHealth;
  if (structured) return structured;

  if (!node.properties.topologyStatus) return null;
  return {
    topologyStatus: node.properties.topologyStatus,
    countedIncoming: node.properties.countedIncoming ?? 0,
    countedOutgoing: node.properties.countedOutgoing ?? 0,
    excludedEdgeCounts: node.properties.excludedEdgeCounts,
    componentId: node.properties.componentId,
    componentSize: node.properties.componentSize,
    componentRootNodeIds: node.properties.componentRootNodeIds,
    componentReachableFromRoot:
      node.properties.componentReachableFromRoot ?? false,
    expectedIsolationReasons: node.properties.expectedIsolationReasons,
    diagnostics: node.properties.diagnostics,
    confidence: node.properties.confidence ?? "candidate",
    resolutionHealthBuckets: node.properties.resolutionHealthBuckets,
    resolutionGapCount: node.properties.resolutionGapCount,
    resolutionConfidence: node.properties.resolutionConfidence ?? "unknown",
  };
};

export const graphHealthMatchesFilters = (
  value: GraphHealthFilterable,
  filters: GraphHealthFilterState,
): boolean => {
  if (
    value.topologyStatus &&
    !filters.visibleTopologyStatuses.includes(value.topologyStatus)
  ) {
    return false;
  }

  const reasons = value.expectedIsolationReasons ?? [];
  if (
    reasons.some((reason) =>
      filters.hiddenExpectedIsolationReasons.includes(reason),
    )
  ) {
    return false;
  }

  const diagnostics = value.diagnostics ?? [];
  if (
    diagnostics.some(
      (diagnostic) =>
        diagnostic.kind &&
        GRAPH_HEALTH_DIAGNOSTIC_KINDS.includes(
          diagnostic.kind as (typeof GRAPH_HEALTH_DIAGNOSTIC_KINDS)[number],
        ) &&
        !filters.visibleDiagnosticKinds.includes(diagnostic.kind),
    )
  ) {
    return false;
  }

  return true;
};

export const graphNodeMatchesHealthFilters = (
  node: GraphNode,
  filters: GraphHealthFilterState,
): boolean => {
  const health = getNodeGraphHealth(node);
  return health ? graphHealthMatchesFilters(health, filters) : true;
};

export const getGraphHealthTopologyCounts = (
  nodes: GraphNode[],
): Map<GraphHealthTopologyStatus, number> => {
  const counts = new Map<GraphHealthTopologyStatus, number>(
    GRAPH_HEALTH_TOPOLOGY_STATUSES.map((status) => [status, 0]),
  );
  for (const node of nodes) {
    const status = getNodeGraphHealth(node)?.topologyStatus;
    if (status) {
      counts.set(status, (counts.get(status) ?? 0) + 1);
    }
  }
  return counts;
};

export const getGraphHealthExpectedReasonCounts = (
  nodes: GraphNode[],
): Map<GraphHealthExpectedIsolationReason, number> => {
  const counts = new Map<GraphHealthExpectedIsolationReason, number>(
    GRAPH_HEALTH_EXPECTED_ISOLATION_REASONS.map((reason) => [reason, 0]),
  );
  for (const node of nodes) {
    for (const reason of getNodeGraphHealth(node)?.expectedIsolationReasons ??
      []) {
      counts.set(reason, (counts.get(reason) ?? 0) + 1);
    }
  }
  return counts;
};

export const getGraphHealthDiagnosticCounts = (
  nodes: GraphNode[],
): Map<string, number> => {
  const counts = new Map<string, number>(
    GRAPH_HEALTH_DIAGNOSTIC_KINDS.map((kind) => [kind, 0]),
  );
  for (const node of nodes) {
    for (const diagnostic of getNodeGraphHealth(node)?.diagnostics ?? []) {
      if (!diagnostic.kind) continue;
      counts.set(
        diagnostic.kind,
        (counts.get(diagnostic.kind) ?? 0) +
          (diagnostic.count && diagnostic.count > 0 ? diagnostic.count : 1),
      );
    }
  }
  return counts;
};

export const getGraphHealthConfidenceCounts = (
  nodes: GraphNode[],
): Map<GraphHealthConfidence, number> => {
  const counts = new Map<GraphHealthConfidence, number>(
    GRAPH_HEALTH_CONFIDENCE_LEVELS.map((confidence) => [confidence, 0]),
  );
  for (const node of nodes) {
    const confidence = getNodeGraphHealth(node)?.confidence;
    if (confidence) {
      counts.set(confidence, (counts.get(confidence) ?? 0) + 1);
    }
  }
  return counts;
};
