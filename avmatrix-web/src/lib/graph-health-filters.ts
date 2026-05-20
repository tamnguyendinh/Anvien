import {
  GRAPH_HEALTH_EXPECTED_ISOLATION_REASONS,
  GRAPH_HEALTH_TOPOLOGY_STATUSES,
  type GraphHealthDiagnostic,
  type GraphHealthExpectedIsolationReason,
  type GraphHealthNodeMetadata,
  type GraphHealthTopologyStatus,
  type GraphNode,
} from '@/generated/avmatrix-contracts';

export const GRAPH_HEALTH_DIAGNOSTIC_KINDS = ['unresolved_reference'] as const;

export type GraphHealthDiagnosticKind = (typeof GRAPH_HEALTH_DIAGNOSTIC_KINDS)[number] | string;

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

export const GRAPH_HEALTH_TOPOLOGY_LABELS: Record<GraphHealthTopologyStatus, string> = {
  connected: 'Connected',
  true_isolated: 'True isolated',
  no_incoming: 'No incoming',
  no_outgoing: 'No outgoing',
  detached_component: 'Detached component',
  unknown_connectivity: 'Unknown',
};

export const GRAPH_HEALTH_REASON_LABELS: Record<GraphHealthExpectedIsolationReason, string> = {
  test: 'Test',
  fixture: 'Fixture',
  generated: 'Generated',
  vendor: 'Vendor',
  documentation: 'Documentation',
  migration: 'Migration',
  exported_api: 'Exported API',
  framework_entry: 'Framework entry',
  cli_mcp: 'CLI/MCP',
};

export const GRAPH_HEALTH_DIAGNOSTIC_LABELS: Record<string, string> = {
  unresolved_reference: 'Unresolved reference',
};

export const getNodeGraphHealth = (node: GraphNode): GraphHealthNodeMetadata | null => {
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
    componentReachableFromRoot: node.properties.componentReachableFromRoot ?? false,
    expectedIsolationReasons: node.properties.expectedIsolationReasons,
    diagnostics: node.properties.diagnostics,
    confidence: node.properties.confidence ?? 'candidate',
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
    for (const reason of getNodeGraphHealth(node)?.expectedIsolationReasons ?? []) {
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
        (counts.get(diagnostic.kind) ?? 0) + (diagnostic.count && diagnostic.count > 0 ? diagnostic.count : 1),
      );
    }
  }
  return counts;
};
