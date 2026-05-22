import {
  APP_LAYERS,
  APP_LAYER_LABELS,
  GRAPH_HEALTH_DIAGNOSTIC_ACTIONABILITIES,
  GRAPH_HEALTH_DIAGNOSTIC_CLASSIFICATIONS,
  GRAPH_HEALTH_RESOLUTION_CONFIDENCE_LEVELS,
  GRAPH_HEALTH_RESOLUTION_HEALTH_BUCKETS,
  type AppLayer,
  type GraphHealthDiagnosticActionability,
  type GraphHealthDiagnosticClassification,
  type GraphHealthResolutionConfidence,
  type GraphHealthResolutionHealthBucket,
  type GraphNode,
} from "@/generated/avmatrix-contracts";
import { getNodeGraphHealth } from "./graph-health-filters";

export const RESOLUTION_GAP_FACT_FAMILIES = [
  "call",
  "access",
  "type-reference",
  "heritage",
  "external",
  "builtin",
  "test",
  "unknown",
] as const;

export type ResolutionGapFactFamily =
  | (typeof RESOLUTION_GAP_FACT_FAMILIES)[number]
  | string;

export const RESOLUTION_GAP_TARGET_ROLES = [
  "callable",
  "member",
  "type",
  "external",
  "builtin",
  "test",
  "unknown",
] as const;

export type ResolutionGapTargetRole =
  | (typeof RESOLUTION_GAP_TARGET_ROLES)[number]
  | string;

export interface SemanticFilterState {
  visibleAppLayers: AppLayer[];
  showNodesMissingAppLayer: boolean;
  visibleResolutionConfidences: GraphHealthResolutionConfidence[];
  visibleResolutionHealthBuckets: GraphHealthResolutionHealthBucket[];
  visibleResolutionGapFactFamilies: ResolutionGapFactFamily[];
  visibleResolutionGapTargetRoles: ResolutionGapTargetRole[];
  visibleResolutionGapClassifications: GraphHealthDiagnosticClassification[];
  visibleResolutionGapActionabilities: GraphHealthDiagnosticActionability[];
  visibleResolutionGapSourceAppLayers: AppLayer[];
  visibleResolutionGapTargetTexts: string[];
}

export interface SemanticFilterable {
  appLayer?: AppLayer;
  functionalArea?: string;
  resolutionConfidence?: GraphHealthResolutionConfidence;
  resolutionHealthBuckets?: Partial<Record<GraphHealthResolutionHealthBucket, number>>;
  resolutionGapCount?: number;
  factFamily?: string;
  gapKind?: string;
  targetRole?: string;
  classification?: GraphHealthDiagnosticClassification;
  actionability?: GraphHealthDiagnosticActionability;
  targetText?: string;
  sourceAppLayer?: AppLayer;
}

export interface SemanticLensRow {
  id: string;
  label: string;
  count: number;
  detail?: string;
}

export const DEFAULT_SEMANTIC_FILTERS: SemanticFilterState = {
  visibleAppLayers: [...APP_LAYERS],
  showNodesMissingAppLayer: true,
  visibleResolutionConfidences: [...GRAPH_HEALTH_RESOLUTION_CONFIDENCE_LEVELS],
  visibleResolutionHealthBuckets: [...GRAPH_HEALTH_RESOLUTION_HEALTH_BUCKETS],
  visibleResolutionGapFactFamilies: [...RESOLUTION_GAP_FACT_FAMILIES],
  visibleResolutionGapTargetRoles: [...RESOLUTION_GAP_TARGET_ROLES],
  visibleResolutionGapClassifications: [...GRAPH_HEALTH_DIAGNOSTIC_CLASSIFICATIONS],
  visibleResolutionGapActionabilities: [...GRAPH_HEALTH_DIAGNOSTIC_ACTIONABILITIES],
  visibleResolutionGapSourceAppLayers: [...APP_LAYERS],
  visibleResolutionGapTargetTexts: [],
};

const appLayerSet = new Set<string>(APP_LAYERS);
const resolutionConfidenceSet = new Set<string>(GRAPH_HEALTH_RESOLUTION_CONFIDENCE_LEVELS);
const resolutionHealthBucketSet = new Set<string>(GRAPH_HEALTH_RESOLUTION_HEALTH_BUCKETS);
const classificationSet = new Set<string>(GRAPH_HEALTH_DIAGNOSTIC_CLASSIFICATIONS);
const actionabilitySet = new Set<string>(GRAPH_HEALTH_DIAGNOSTIC_ACTIONABILITIES);

const appLayerLabels = new Map<string, (typeof APP_LAYER_LABELS)[number]>(
  APP_LAYER_LABELS.map((entry) => [entry.key, entry]),
);

const asString = (value: unknown): string | undefined =>
  typeof value === "string" && value.length > 0 ? value : undefined;

const asAppLayer = (value: unknown): AppLayer | undefined => {
  const text = asString(value);
  return text && appLayerSet.has(text) ? (text as AppLayer) : undefined;
};

const asResolutionConfidence = (
  value: unknown,
): GraphHealthResolutionConfidence | undefined => {
  const text = asString(value);
  return text && resolutionConfidenceSet.has(text)
    ? (text as GraphHealthResolutionConfidence)
    : undefined;
};

const asClassification = (
  value: unknown,
): GraphHealthDiagnosticClassification | undefined => {
  const text = asString(value);
  return text && classificationSet.has(text)
    ? (text as GraphHealthDiagnosticClassification)
    : undefined;
};

const asActionability = (
  value: unknown,
): GraphHealthDiagnosticActionability | undefined => {
  const text = asString(value);
  return text && actionabilitySet.has(text)
    ? (text as GraphHealthDiagnosticActionability)
    : undefined;
};

const asResolutionHealthBuckets = (
  value: unknown,
): Partial<Record<GraphHealthResolutionHealthBucket, number>> | undefined => {
  if (!value || typeof value !== "object") return undefined;
  const out: Partial<Record<GraphHealthResolutionHealthBucket, number>> = {};
  for (const [key, rawCount] of Object.entries(value as Record<string, unknown>)) {
    if (!resolutionHealthBucketSet.has(key)) continue;
    const count = typeof rawCount === "number" ? rawCount : Number(rawCount);
    if (Number.isFinite(count) && count > 0) {
      out[key as GraphHealthResolutionHealthBucket] = count;
    }
  }
  return Object.keys(out).length > 0 ? out : undefined;
};

export const getAppLayerLabel = (layer: AppLayer | string): string =>
  appLayerLabels.get(layer)?.webLabel ?? layer;

export const getAppLayerDescription = (layer: AppLayer | string): string =>
  appLayerLabels.get(layer)?.description ?? "Persisted App Layer value from graph data.";

export const getSemanticFilterableFromNode = (
  node: GraphNode,
): SemanticFilterable => {
  const health = getNodeGraphHealth(node);
  const properties = node.properties;

  return {
    appLayer: asAppLayer(properties.appLayer),
    functionalArea: asString(properties.functionalArea),
    resolutionConfidence:
      asResolutionConfidence(properties.resolutionConfidence) ??
      health?.resolutionConfidence,
    resolutionHealthBuckets:
      asResolutionHealthBuckets(properties.resolutionHealthBuckets) ??
      health?.resolutionHealthBuckets,
    resolutionGapCount:
      typeof properties.resolutionGapCount === "number"
        ? properties.resolutionGapCount
        : health?.resolutionGapCount,
    factFamily: asString(properties.factFamily),
    gapKind: asString(properties.gapKind),
    targetRole: asString(properties.targetRole),
    classification: asClassification(properties.classification),
    actionability: asActionability(properties.actionability),
    targetText: asString(properties.targetText),
    sourceAppLayer: asAppLayer(properties.sourceAppLayer),
  };
};

export const getAppLayerCounts = (
  nodes: GraphNode[],
): Map<AppLayer | "missing", number> => {
  const counts = new Map<AppLayer | "missing", number>(
    APP_LAYERS.map((layer) => [layer, 0]),
  );
  counts.set("missing", 0);

  for (const node of nodes) {
    const layer = asAppLayer(node.properties.appLayer);
    counts.set(layer ?? "missing", (counts.get(layer ?? "missing") ?? 0) + 1);
  }

  return counts;
};

export const getResolutionConfidenceCounts = (
  nodes: GraphNode[],
): Map<GraphHealthResolutionConfidence, number> => {
  const counts = new Map<GraphHealthResolutionConfidence, number>(
    GRAPH_HEALTH_RESOLUTION_CONFIDENCE_LEVELS.map((confidence) => [confidence, 0]),
  );
  for (const node of nodes) {
    const confidence = getSemanticFilterableFromNode(node).resolutionConfidence;
    if (confidence) counts.set(confidence, (counts.get(confidence) ?? 0) + 1);
  }
  return counts;
};

export const getResolutionHealthBucketCounts = (
  nodes: GraphNode[],
): Map<GraphHealthResolutionHealthBucket, number> => {
  const counts = new Map<GraphHealthResolutionHealthBucket, number>(
    GRAPH_HEALTH_RESOLUTION_HEALTH_BUCKETS.map((bucket) => [bucket, 0]),
  );
  for (const node of nodes) {
    const buckets = getSemanticFilterableFromNode(node).resolutionHealthBuckets ?? {};
    for (const [bucket, count] of Object.entries(buckets)) {
      counts.set(
        bucket as GraphHealthResolutionHealthBucket,
        (counts.get(bucket as GraphHealthResolutionHealthBucket) ?? 0) +
          (count ?? 0),
      );
    }
  }
  return counts;
};

const countBy = <T extends string>(
  nodes: GraphNode[],
  read: (value: SemanticFilterable) => T | undefined,
): Map<T, number> => {
  const counts = new Map<T, number>();
  for (const node of nodes) {
    const value = read(getSemanticFilterableFromNode(node));
    if (!value) continue;
    counts.set(value, (counts.get(value) ?? 0) + 1);
  }
  return counts;
};

export const getResolutionGapFactFamilyCounts = (nodes: GraphNode[]) =>
  countBy(nodes, (value) => value.factFamily);

export const getResolutionGapTargetRoleCounts = (nodes: GraphNode[]) =>
  countBy(nodes, (value) => value.targetRole);

export const getResolutionGapClassificationCounts = (nodes: GraphNode[]) =>
  countBy(nodes, (value) => value.classification);

export const getResolutionGapActionabilityCounts = (nodes: GraphNode[]) =>
  countBy(nodes, (value) => value.actionability);

export const getResolutionGapSourceAppLayerCounts = (nodes: GraphNode[]) =>
  countBy(nodes, (value) => value.sourceAppLayer);

export const getResolutionGapTargetTextCounts = (nodes: GraphNode[]) =>
  countBy(nodes, (value) => value.targetText);

export const getTopCountEntries = <T extends string>(
  counts: Map<T, number>,
  limit: number,
): Array<{ value: T; count: number }> =>
  [...counts.entries()]
    .filter(([, count]) => count > 0)
    .sort((left, right) => right[1] - left[1] || left[0].localeCompare(right[0]))
    .slice(0, limit)
    .map(([value, count]) => ({ value, count }));

const includesKnownOrOpen = <T extends string>(
  selected: readonly T[],
  value: string | undefined,
): boolean => !value || selected.includes(value as T);

const hasVisibleBucket = (
  buckets: Partial<Record<GraphHealthResolutionHealthBucket, number>> | undefined,
  visibleBuckets: readonly GraphHealthResolutionHealthBucket[],
): boolean => {
  if (!buckets || Object.keys(buckets).length === 0) return true;
  return Object.entries(buckets).some(
    ([bucket, count]) =>
      (count ?? 0) > 0 &&
      visibleBuckets.includes(bucket as GraphHealthResolutionHealthBucket),
  );
};

export const semanticMatchesFilters = (
  value: SemanticFilterable,
  filters: SemanticFilterState,
): boolean => {
  if (value.appLayer) {
    if (!filters.visibleAppLayers.includes(value.appLayer)) return false;
  } else if (!filters.showNodesMissingAppLayer) {
    return false;
  }

  if (
    value.resolutionConfidence &&
    !filters.visibleResolutionConfidences.includes(value.resolutionConfidence)
  ) {
    return false;
  }

  if (
    !hasVisibleBucket(
      value.resolutionHealthBuckets,
      filters.visibleResolutionHealthBuckets,
    )
  ) {
    return false;
  }

  if (!includesKnownOrOpen(filters.visibleResolutionGapFactFamilies, value.factFamily)) {
    return false;
  }
  if (!includesKnownOrOpen(filters.visibleResolutionGapTargetRoles, value.targetRole)) {
    return false;
  }
  if (
    !includesKnownOrOpen(
      filters.visibleResolutionGapClassifications,
      value.classification,
    )
  ) {
    return false;
  }
  if (
    !includesKnownOrOpen(
      filters.visibleResolutionGapActionabilities,
      value.actionability,
    )
  ) {
    return false;
  }
  if (
    value.sourceAppLayer &&
    !filters.visibleResolutionGapSourceAppLayers.includes(value.sourceAppLayer)
  ) {
    return false;
  }
  if (
    filters.visibleResolutionGapTargetTexts.length > 0 &&
    (!value.targetText ||
      !filters.visibleResolutionGapTargetTexts.includes(value.targetText))
  ) {
    return false;
  }

  return true;
};

const countMatching = (
  nodes: GraphNode[],
  predicate: (value: SemanticFilterable) => boolean,
): number =>
  nodes.reduce(
    (count, node) => count + (predicate(getSemanticFilterableFromNode(node)) ? 1 : 0),
    0,
  );

const topLabel = <T extends string>(
  counts: Map<T, number>,
  label: (value: T) => string = (value) => value,
): string | undefined => {
  const top = getTopCountEntries(counts, 1)[0];
  return top ? `${label(top.value)} (${top.count})` : undefined;
};

export const getResolutionLensRows = (nodes: GraphNode[]): SemanticLensRow[] => {
  const analyzerGapNodes = nodes.filter(
    (node) => getSemanticFilterableFromNode(node).actionability === "analyzer_gap",
  );
  const gapNodes = nodes.filter((node) => node.label === "ResolutionGap");
  const functionalAreaCounts = countBy(gapNodes, (value) => value.functionalArea);

  return [
    {
      id: "backend-unresolved-calls",
      label: "Backend unresolved calls",
      count: countMatching(
        gapNodes,
        (value) => value.sourceAppLayer === "backend" && value.factFamily === "call",
      ),
    },
    {
      id: "api-unresolved-handlers-contracts",
      label: "API unresolved handlers/contracts",
      count: countMatching(
        gapNodes,
        (value) =>
          value.sourceAppLayer === "api" ||
          value.sourceAppLayer === "api_contract" ||
          value.sourceAppLayer === "api_shared_contract",
      ),
    },
    {
      id: "frontend-unresolved-type-refs",
      label: "Frontend unresolved type refs",
      count: countMatching(
        gapNodes,
        (value) =>
          (value.sourceAppLayer === "frontend" ||
            value.sourceAppLayer === "frontend_api_client" ||
            value.sourceAppLayer === "frontend_test") &&
          value.factFamily === "type-reference",
      ),
    },
    {
      id: "shared-contract-analyzer-gaps",
      label: "Shared contract analyzer gaps",
      count: countMatching(
        gapNodes,
        (value) =>
          value.actionability === "analyzer_gap" &&
          (value.sourceAppLayer === "shared_contract" ||
            value.sourceAppLayer === "api_contract" ||
            value.sourceAppLayer === "api_shared_contract" ||
            value.sourceAppLayer === "generated_contract"),
      ),
    },
    {
      id: "external-unresolved-symbols",
      label: "External unresolved symbols",
      count: countMatching(
        gapNodes,
        (value) => value.classification === "external_library",
      ),
    },
    {
      id: "builtin-non-actionable",
      label: "Builtin non-actionable",
      count: countMatching(
        gapNodes,
        (value) => value.classification === "builtin",
      ),
    },
    {
      id: "standard-library-non-actionable",
      label: "Standard library non-actionable",
      count: countMatching(
        gapNodes,
        (value) => value.classification === "standard_library",
      ),
    },
    {
      id: "test-framework-non-actionable",
      label: "Test framework non-actionable",
      count: countMatching(
        gapNodes,
        (value) => value.classification === "test_framework",
      ),
    },
    {
      id: "in-repo-analyzer-gaps",
      label: "In-repo analyzer gaps",
      count: countMatching(
        gapNodes,
        (value) =>
          value.actionability === "analyzer_gap" ||
          value.classification === "in_repo_unresolved",
      ),
    },
    {
      id: "resolution-gaps-by-functional-area",
      label: "Resolution gaps by functional area",
      count: gapNodes.length,
      detail: topLabel(functionalAreaCounts),
    },
    {
      id: "top-app-layers-by-analyzer-gap",
      label: "Top app layers by analyzer gap count",
      count: analyzerGapNodes.length,
      detail: topLabel(
        getResolutionGapSourceAppLayerCounts(analyzerGapNodes),
        getAppLayerLabel,
      ),
    },
    {
      id: "top-functional-areas-by-unresolved",
      label: "Top functional areas by unresolved count",
      count: gapNodes.length,
      detail: topLabel(functionalAreaCounts),
    },
    {
      id: "top-unresolved-target-text",
      label: "Top unresolved target text",
      count: gapNodes.length,
      detail: topLabel(getResolutionGapTargetTextCounts(gapNodes)),
    },
  ];
};
