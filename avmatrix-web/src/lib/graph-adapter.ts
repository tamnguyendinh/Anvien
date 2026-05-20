import Graph, { MultiDirectedGraph } from "graphology";
import type { KnowledgeGraph } from "../core/graph/types";
import {
  COMMUNITY_COLORED_NODE_LABELS,
  DOCUMENTATION_NODE_LABEL,
  EDGE_SIZE_MULTIPLIERS,
  FILTERABLE_LABELS,
  getDisplayGraphRelationships,
  getCommunityColor,
  getEdgeInfo,
  getNodeDisplayLabel,
  getNodeColor,
  getNodeSize,
} from "./constants";
import {
  graphHealthMatchesFilters,
  getNodeGraphHealth,
  type GraphHealthFilterable,
  type GraphHealthFilterState,
} from "./graph-health-filters";

export interface SigmaNodeAttributes extends GraphHealthFilterable {
  x: number;
  y: number;
  size: number;
  color: string;
  label: string;
  nodeType: string;
  rawNodeType?: string;
  filePath: string;
  startLine?: number;
  endLine?: number;
  hidden?: boolean;
  zIndex?: number;
  highlighted?: boolean;
  mass?: number; // ForceAtlas2 mass - higher = more repulsion
  community?: number; // Community index from Leiden algorithm
  communityColor?: string; // Color assigned by community
  confidence?: string;
}

export interface SigmaEdgeAttributes {
  size: number;
  color: string;
  relationType: string;
  hidden?: boolean;
  zIndex?: number;
}

/**
 * Get node size scaled for graph density
 * Uses lower minimums to maintain hierarchy visibility even in huge graphs
 */
const getMaxScaledNodeSize = (_nodeCount: number): number => {
  return 3;
};

const getLabelScaledNodeSizeCap = (
  label: string | undefined,
  nodeCount: number,
): number => {
  if (label === "Package") {
    if (nodeCount > 20000) return 1.5;
    if (nodeCount > 5000) return 2;
    return 3;
  }
  if (label === "Section") {
    if (nodeCount > 20000) return 1;
    if (nodeCount > 5000) return 1.5;
    return 2;
  }
  return getMaxScaledNodeSize(nodeCount);
};

export const MAX_RENDERED_NODE_SIZE = 3;
export const MAX_DENSE_RENDERED_NODE_SIZE = 3;

export const getMaxRenderedNodeSize = (_nodeCount: number): number =>
  MAX_DENSE_RENDERED_NODE_SIZE;

export const capRenderedNodeSize = (
  size: number,
  nodeCount: number = 0,
): number => Math.min(size, getMaxRenderedNodeSize(nodeCount));

const communityColoredNodeLabelSet = new Set<string>(
  COMMUNITY_COLORED_NODE_LABELS,
);

const filterableLabelOrder = new Map<string, number>(
  FILTERABLE_LABELS.map((label, index) => [label, index]),
);

const stableString = (value: unknown): string =>
  typeof value === "string" ? value : value == null ? "" : String(value);

const compareStableString = (left: string, right: string): number => {
  if (left === right) return 0;
  return left < right ? -1 : 1;
};

const compareClusterLabels = (left: string, right: string): number => {
  const leftOrder = filterableLabelOrder.get(left);
  const rightOrder = filterableLabelOrder.get(right);

  if (leftOrder !== undefined && rightOrder !== undefined) {
    return leftOrder - rightOrder;
  }
  if (leftOrder !== undefined) return -1;
  if (rightOrder !== undefined) return 1;
  return compareStableString(left, right);
};

const getClusterNodeSpacing = (nodeCount: number): number => {
  if (nodeCount > 50000) return 18;
  if (nodeCount > 20000) return 20;
  if (nodeCount > 5000) return 24;
  if (nodeCount > 1000) return 30;
  return 42;
};

const GOLDEN_ANGLE = Math.PI * (3 - Math.sqrt(5));

const getStableLabelSeed = (label: string): number => {
  let hash = 2166136261;
  for (let index = 0; index < label.length; index++) {
    hash ^= label.charCodeAt(index);
    hash = Math.imul(hash, 16777619);
  }
  return hash >>> 0;
};

const getClusterIslandRadius = (
  nodeCount: number,
  nodeSpacing: number,
): number => {
  if (nodeCount <= 1) return nodeSpacing * 4;
  return nodeSpacing * Math.sqrt(nodeCount - 1) * 1.08 + nodeSpacing * 4;
};

const getIslandOffset = (
  nodeIndex: number,
  nodeSpacing: number,
  labelSeed: number,
): { x: number; y: number } => {
  if (nodeIndex === 0) {
    return { x: 0, y: 0 };
  }

  const seedRadians = ((labelSeed % 360) * Math.PI) / 180;
  const organicWave = Math.sin((nodeIndex + 1) * ((labelSeed % 997) + 1) * 0.013);
  const radius = nodeSpacing * Math.sqrt(nodeIndex) * (1 + organicWave * 0.035);
  const angle =
    nodeIndex * GOLDEN_ANGLE +
    seedRadians +
    Math.sin((nodeIndex + 1) * 0.37 + seedRadians) * 0.025;

  return {
    x: Math.cos(angle) * radius,
    y: Math.sin(angle) * radius,
  };
};

const getBalancedCircularSlots = (slotCount: number): number[] => {
  if (slotCount <= 0) return [];

  const slots = [0];
  while (slots.length < slotCount) {
    let bestSlot = -1;
    let bestDistance = -1;

    for (let candidate = 0; candidate < slotCount; candidate++) {
      if (slots.includes(candidate)) continue;

      const nearestDistance = slots.reduce((nearest, occupiedSlot) => {
        const clockwise = Math.abs(candidate - occupiedSlot);
        const circularDistance = Math.min(clockwise, slotCount - clockwise);
        return Math.min(nearest, circularDistance);
      }, Number.POSITIVE_INFINITY);

      if (nearestDistance > bestDistance) {
        bestDistance = nearestDistance;
        bestSlot = candidate;
      }
    }

    slots.push(bestSlot);
  }

  return slots;
};

const compareClusterNodeIds = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
  leftNodeId: string,
  rightNodeId: string,
): number => {
  const left = graph.getNodeAttributes(leftNodeId);
  const right = graph.getNodeAttributes(rightNodeId);

  return (
    compareStableString(stableString(left.filePath), stableString(right.filePath)) ||
    compareStableString(stableString(left.label), stableString(right.label)) ||
    compareStableString(leftNodeId, rightNodeId)
  );
};

export const applyFilterBasedClusteredLayout = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
): void => {
  const totalNodeCount = graph.order;
  if (totalNodeCount === 0) return;

  const nodeIdsByLabel = new Map<string, string[]>();
  graph.forEachNode((nodeId, attributes) => {
    const label = stableString(attributes.nodeType);
    const nodeIds = nodeIdsByLabel.get(label) ?? [];
    nodeIds.push(nodeId);
    nodeIdsByLabel.set(label, nodeIds);
  });

  const nodeSpacing = getClusterNodeSpacing(totalNodeCount);
  const clusterLabels = [...nodeIdsByLabel.keys()].sort(compareClusterLabels);

  const clusters = clusterLabels.map((label) => {
    const nodeIds = [...(nodeIdsByLabel.get(label) ?? [])].sort((left, right) =>
      compareClusterNodeIds(graph, left, right),
    );

    return {
      label,
      nodeIds,
      labelSeed: getStableLabelSeed(label),
      radius: getClusterIslandRadius(nodeIds.length, nodeSpacing),
    };
  });
  const placeCluster = (
    cluster: (typeof clusters)[number],
    centerX: number,
    centerY: number,
  ) => {
    const offsets = cluster.nodeIds.map((_, nodeIndex) =>
      getIslandOffset(nodeIndex, nodeSpacing, cluster.labelSeed),
    );
    const bounds = offsets.reduce(
      (current, offset) => ({
        minX: Math.min(current.minX, offset.x),
        maxX: Math.max(current.maxX, offset.x),
        minY: Math.min(current.minY, offset.y),
        maxY: Math.max(current.maxY, offset.y),
      }),
      {
        minX: Number.POSITIVE_INFINITY,
        maxX: Number.NEGATIVE_INFINITY,
        minY: Number.POSITIVE_INFINITY,
        maxY: Number.NEGATIVE_INFINITY,
      },
    );
    const offsetCenterX = (bounds.minX + bounds.maxX) / 2;
    const offsetCenterY = (bounds.minY + bounds.maxY) / 2;

    cluster.nodeIds.forEach((nodeId, nodeIndex) => {
      const offset = offsets[nodeIndex];
      graph.setNodeAttribute(nodeId, "x", centerX + offset.x - offsetCenterX);
      graph.setNodeAttribute(nodeId, "y", centerY + offset.y - offsetCenterY);
    });
  };

  const documentationCluster = clusters.find(
    (cluster) => cluster.label === DOCUMENTATION_NODE_LABEL,
  );
  const outerClusters = clusters.filter(
    (cluster) => cluster.label !== DOCUMENTATION_NODE_LABEL,
  );

  if (documentationCluster) {
    placeCluster(documentationCluster, 0, 0);
  }
  if (outerClusters.length === 0) return;

  const balancedSlots = getBalancedCircularSlots(outerClusters.length);
  const clustersByPlacementSize = [...outerClusters].sort(
    (left, right) =>
      right.radius - left.radius || compareClusterLabels(left.label, right.label),
  );
  const slotByClusterLabel = new Map<string, number>();
  clustersByPlacementSize.forEach((cluster, index) => {
    slotByClusterLabel.set(cluster.label, balancedSlots[index] ?? index);
  });
  const clustersBySlot = outerClusters.map((cluster) => ({
    cluster,
    slot: slotByClusterLabel.get(cluster.label) ?? 0,
  }));

  const largestOuterClusterRadius = outerClusters.reduce(
    (maximum, cluster) => Math.max(maximum, cluster.radius),
    nodeSpacing * 4,
  );
  const clusterGap = Math.max(
    nodeSpacing * 18,
    largestOuterClusterRadius * 0.18,
  );
  const centerClearance = documentationCluster
    ? documentationCluster.radius + largestOuterClusterRadius + clusterGap
    : 0;
  const minimumAngularStep =
    outerClusters.length <= 1 ? Math.PI : Math.sin(Math.PI / outerClusters.length);
  const largestAdjacentClusterSpan = clustersBySlot.reduce(
    (maximum, placedCluster) => {
      const nextPlacedCluster =
        clustersBySlot.find(
          (candidate) =>
            candidate.slot === (placedCluster.slot + 1) % outerClusters.length,
        ) ?? placedCluster;
      return Math.max(
        maximum,
        placedCluster.cluster.radius + nextPlacedCluster.cluster.radius + clusterGap,
      );
    },
    nodeSpacing * 8,
  );
  const orbitRadius =
    outerClusters.length <= 1
      ? centerClearance
      : Math.max(
          largestOuterClusterRadius * 1.4 + clusterGap,
          largestAdjacentClusterSpan / (2 * minimumAngularStep),
          centerClearance,
          nodeSpacing * 24,
        );
  const radialBand = Math.min(clusterGap * 0.35, orbitRadius * 0.06);

  outerClusters.forEach((cluster, clusterIndex) => {
    const clusterSlot = slotByClusterLabel.get(cluster.label) ?? clusterIndex;
    const clusterAngle =
      outerClusters.length <= 1
        ? 0
        : -Math.PI / 2 + (clusterSlot / outerClusters.length) * Math.PI * 2;
    const radialOffset =
      outerClusters.length <= 2 ? 0 : ((clusterIndex % 3) - 1) * radialBand;
    const clusterCenterRadius = orbitRadius + radialOffset;
    const clusterCenterX = Math.cos(clusterAngle) * clusterCenterRadius;
    const clusterCenterY = Math.sin(clusterAngle) * clusterCenterRadius;

    placeCluster(cluster, clusterCenterX, clusterCenterY);
  });
};

export const getScaledNodeSize = (
  baseSize: number,
  nodeCount: number,
  label?: string,
): number => {
  // Scale factor decreases as graph gets larger
  // But a minimum is used that preserves relative differences
  let scaledSize = baseSize;
  if (nodeCount > 50000) scaledSize = Math.max(1, baseSize * 0.4);
  else if (nodeCount > 20000) scaledSize = Math.max(1.5, baseSize * 0.5);
  else if (nodeCount > 5000) scaledSize = Math.max(2, baseSize * 0.65);
  else if (nodeCount > 1000) scaledSize = Math.max(2.5, baseSize * 0.8);

  return Math.min(scaledSize, getLabelScaledNodeSizeCap(label, nodeCount));
};

/**
 * Get mass for node type - higher mass = more repulsion in ForceAtlas2
 * Folders get MUCH higher mass so they spread out and pull their files with them
 */
const getNodeMass = (nodeType: string, nodeCount: number): number => {
  // Scale mass based on graph size
  const baseMassMultiplier = nodeCount > 5000 ? 2 : nodeCount > 1000 ? 1.5 : 1;

  const structuralRank: Record<string, number> = {
    Project: 50,
    Package: 30,
    Module: 20,
    Namespace: 20,
    Folder: 15,
    File: 3,
    Section: 2,
    Community: 2,
    Process: 4,
    Route: 4,
    Tool: 4,
  };
  const structuralMass = structuralRank[nodeType];
  if (structuralMass !== undefined) {
    return structuralMass * baseMassMultiplier;
  }
  if (communityColoredNodeLabelSet.has(nodeType)) {
    return 3 * baseMassMultiplier;
  }
  return Math.max(1, Math.min(2, getNodeSize(nodeType))) * baseMassMultiplier;
};

/**
 * Converts the KnowledgeGraph to a graphology Graph for Sigma.js
 * Nodes are positioned by node-label clusters matching the Web node filters.
 *
 * @param knowledgeGraph - The knowledge graph to convert
 * @param communityMemberships - Optional map of nodeId -> communityIndex for metadata
 */
export const knowledgeGraphToGraphology = (
  knowledgeGraph: KnowledgeGraph,
  communityMemberships?: Map<string, number>,
): Graph<SigmaNodeAttributes, SigmaEdgeAttributes> => {
  const graph = new MultiDirectedGraph<
    SigmaNodeAttributes,
    SigmaEdgeAttributes
  >();
  const nodeCount = knowledgeGraph.nodes.length;
  const displayRelationships = getDisplayGraphRelationships(
    knowledgeGraph.relationships,
  );

  knowledgeGraph.nodes.forEach((node) => {
    const displayNodeType = getNodeDisplayLabel(node);
    const baseSize = getNodeSize(displayNodeType);
    const scaledSize = getScaledNodeSize(baseSize, nodeCount, displayNodeType);
    const communityIndex = communityMemberships?.get(node.id);
    const hasCommunity = communityIndex !== undefined;
    const nodeColor = getNodeColor(displayNodeType);
    const graphHealth = getNodeGraphHealth(node);

    graph.addNode(node.id, {
      x: 0,
      y: 0,
      size: scaledSize,
      color: nodeColor,
      label: node.properties.name,
      nodeType: displayNodeType,
      rawNodeType: node.label,
      filePath: node.properties.filePath,
      startLine: node.properties.startLine,
      endLine: node.properties.endLine,
      hidden: false,
      mass: getNodeMass(displayNodeType, nodeCount),
      topologyStatus: graphHealth?.topologyStatus,
      expectedIsolationReasons: graphHealth?.expectedIsolationReasons,
      diagnostics: graphHealth?.diagnostics,
      confidence: graphHealth?.confidence,
      community: communityIndex,
      communityColor: hasCommunity
        ? getCommunityColor(communityIndex!)
        : undefined,
    });
  });

  applyFilterBasedClusteredLayout(graph);

  // Add edges with distinct colors per relationship type
  const edgeBaseSize = nodeCount > 20000 ? 0.4 : nodeCount > 5000 ? 0.6 : 1.0;

  displayRelationships.forEach((rel, index) => {
    if (graph.hasNode(rel.sourceId) && graph.hasNode(rel.targetId)) {
      const edgeInfo = getEdgeInfo(rel.type);
      const sizeMultiplier = EDGE_SIZE_MULTIPLIERS[rel.type] ?? 0.5;
      const edgeKeyBase =
        rel.id || `${rel.sourceId}->${rel.targetId}:${rel.type}:${index}`;
      let edgeKey = edgeKeyBase;
      let duplicateIndex = 1;
      while (graph.hasEdge(edgeKey)) {
        edgeKey = `${edgeKeyBase}:${duplicateIndex}`;
        duplicateIndex++;
      }

      graph.addDirectedEdgeWithKey(edgeKey, rel.sourceId, rel.targetId, {
        size: edgeBaseSize * sizeMultiplier,
        color: edgeInfo.color,
        relationType: rel.type,
      });
    }
  });

  return graph;
};

/**
 * Filter nodes by visibility - sets hidden attribute
 */
export const filterGraphByLabels = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
  visibleLabels: string[],
  graphHealthFilters?: GraphHealthFilterState,
): void => {
  graph.forEachNode((nodeId, attributes) => {
    const isVisible = visibleLabels.includes(attributes.nodeType);
    const isGraphHealthVisible = graphHealthFilters
      ? graphHealthMatchesFilters(attributes, graphHealthFilters)
      : true;
    graph.setNodeAttribute(nodeId, "hidden", !isVisible || !isGraphHealthVisible);
  });
};

/**
 * Get all nodes within N hops of a starting node
 */
export const getNodesWithinHops = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
  startNodeId: string,
  maxHops: number,
): Set<string> => {
  const visited = new Set<string>();
  const queue: { nodeId: string; depth: number }[] = [
    { nodeId: startNodeId, depth: 0 },
  ];

  while (queue.length > 0) {
    const { nodeId, depth } = queue.shift()!;

    if (visited.has(nodeId)) continue;
    visited.add(nodeId);

    if (depth < maxHops) {
      graph.forEachNeighbor(nodeId, (neighborId) => {
        if (!visited.has(neighborId)) {
          queue.push({ nodeId: neighborId, depth: depth + 1 });
        }
      });
    }
  }

  return visited;
};

/**
 * Filter nodes by depth from selected node
 */
export const filterGraphByDepth = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
  selectedNodeId: string | null,
  maxHops: number | null,
  visibleLabels: string[],
  graphHealthFilters?: GraphHealthFilterState,
): void => {
  if (maxHops === null) {
    filterGraphByLabels(graph, visibleLabels, graphHealthFilters);
    return;
  }

  if (selectedNodeId === null || !graph.hasNode(selectedNodeId)) {
    filterGraphByLabels(graph, visibleLabels, graphHealthFilters);
    return;
  }

  const nodesInRange = getNodesWithinHops(graph, selectedNodeId, maxHops);

  graph.forEachNode((nodeId, attributes) => {
    const isLabelVisible = visibleLabels.includes(attributes.nodeType);
    const isInRange = nodesInRange.has(nodeId);
    const isGraphHealthVisible = graphHealthFilters
      ? graphHealthMatchesFilters(attributes, graphHealthFilters)
      : true;
    graph.setNodeAttribute(
      nodeId,
      "hidden",
      !isLabelVisible || !isInRange || !isGraphHealthVisible,
    );
  });
};
