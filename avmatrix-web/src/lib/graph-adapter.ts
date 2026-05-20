import Graph, { MultiDirectedGraph } from "graphology";
import type { KnowledgeGraph } from "../core/graph/types";
import {
  COMMUNITY_COLORED_NODE_LABELS,
  EDGE_SIZE_MULTIPLIERS,
  STRUCTURAL_NODE_LABELS,
  getDisplayGraphRelationships,
  getCommunityColor,
  getEdgeInfo,
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
const getMaxScaledNodeSize = (nodeCount: number): number => {
  if (nodeCount > 20000) return 3;
  if (nodeCount > 5000) return 4.5;
  if (nodeCount > 1000) return 6;
  return 4.5;
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

export const MAX_RENDERED_NODE_SIZE = 9;
export const MAX_DENSE_RENDERED_NODE_SIZE = 3;

export const getMaxRenderedNodeSize = (nodeCount: number): number =>
  nodeCount > 20000 ? MAX_DENSE_RENDERED_NODE_SIZE : MAX_RENDERED_NODE_SIZE;

export const capRenderedNodeSize = (
  size: number,
  nodeCount: number = 0,
): number => Math.min(size, getMaxRenderedNodeSize(nodeCount));

const communityColoredNodeLabelSet = new Set<string>(
  COMMUNITY_COLORED_NODE_LABELS,
);

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
 * Folders are positioned in a wide spread, children positioned NEAR their parents
 *
 * @param knowledgeGraph - The knowledge graph to convert
 * @param communityMemberships - Optional map of nodeId -> communityIndex for community coloring
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

  // Build parent-child map from relationships that materially improve layout.
  // Higher-priority owner/process relationships can replace broad file-level
  // grouping; community membership stays lower because cluster coloring already
  // handles that view.
  // parent -> children
  const parentToChildren = new Map<string, string[]>();
  // child -> parent
  const childToParent = new Map<string, string>();
  const childParentPriority = new Map<string, number>();

  const forwardHierarchyRelations: Record<string, number> = {
    CONTAINS: 100,
    HAS_METHOD: 95,
    HAS_PROPERTY: 95,
    DEFINES: 75,
    IMPORTS: 60,
    WRAPS: 55,
  };

  const reverseHierarchyRelations: Record<string, number> = {
    STEP_IN_PROCESS: 85,
    ENTRY_POINT_OF: 85,
    HANDLES_ROUTE: 85,
    HANDLES_TOOL: 85,
    MEMBER_OF: 50,
  };

  const addHierarchyLink = (
    parentId: string,
    childId: string,
    priority: number,
  ) => {
    const existingPriority = childParentPriority.get(childId) ?? -1;
    if (existingPriority > priority) return;

    const existingParentId = childToParent.get(childId);
    if (existingParentId && existingParentId !== parentId) {
      const siblings = parentToChildren.get(existingParentId);
      if (siblings) {
        parentToChildren.set(
          existingParentId,
          siblings.filter((id) => id !== childId),
        );
      }
    }

    if (!parentToChildren.has(parentId)) {
      parentToChildren.set(parentId, []);
    }
    const children = parentToChildren.get(parentId)!;
    if (!children.includes(childId)) {
      children.push(childId);
    }
    childToParent.set(childId, parentId);
    childParentPriority.set(childId, priority);
  };

  displayRelationships.forEach((rel) => {
    const forwardPriority = forwardHierarchyRelations[rel.type];
    if (forwardPriority !== undefined) {
      addHierarchyLink(rel.sourceId, rel.targetId, forwardPriority);
      return;
    }

    const reversePriority = reverseHierarchyRelations[rel.type];
    if (reversePriority !== undefined) {
      addHierarchyLink(rel.targetId, rel.sourceId, reversePriority);
    }
  });

  // Create node lookup
  const nodeMap = new Map(knowledgeGraph.nodes.map((n) => [n.id, n]));

  // Separate root/grouping nodes from content nodes.
  const structuralTypes = new Set<string>(STRUCTURAL_NODE_LABELS);
  const structuralNodes = knowledgeGraph.nodes.filter((n) =>
    structuralTypes.has(n.label),
  );

  // Much wider spread for structural nodes - this is the key!
  const structuralSpread = Math.sqrt(nodeCount) * 40;
  // Small jitter for children around their parent
  const childJitter = Math.sqrt(nodeCount) * 3;

  // === CLUSTER-BASED POSITIONING ===
  // Calculate cluster centers - each cluster gets a region of the graph
  const clusterCenters = new Map<number, { x: number; y: number }>();
  if (communityMemberships && communityMemberships.size > 0) {
    // Find unique community IDs
    const communities = new Set(communityMemberships.values());
    const communityCount = communities.size;
    const clusterSpread = structuralSpread * 0.8; // Clusters spread across 80% of graph

    // Position cluster centers using golden angle for even distribution
    const goldenAngle = Math.PI * (3 - Math.sqrt(5));
    let idx = 0;
    communities.forEach((communityId) => {
      const angle = idx * goldenAngle;
      const radius = clusterSpread * Math.sqrt((idx + 1) / communityCount);
      clusterCenters.set(communityId, {
        x: radius * Math.cos(angle),
        y: radius * Math.sin(angle),
      });
      idx++;
    });
  }
  // Jitter within cluster (tighter than childJitter)
  const clusterJitter = Math.sqrt(nodeCount) * 1.5;

  // Store positions for parent lookup
  const nodePositions = new Map<string, { x: number; y: number }>();

  // Position structural nodes (folders, etc.) in a wide radial pattern FIRST
  structuralNodes.forEach((node, index) => {
    // Use golden angle for even distribution
    const goldenAngle = Math.PI * (3 - Math.sqrt(5));
    const angle = index * goldenAngle;
    const radius =
      structuralSpread *
      Math.sqrt((index + 1) / Math.max(structuralNodes.length, 1));

    // Add some randomness to prevent perfect patterns
    const jitter = structuralSpread * 0.15;
    const x = radius * Math.cos(angle) + (Math.random() - 0.5) * jitter;
    const y = radius * Math.sin(angle) + (Math.random() - 0.5) * jitter;

    nodePositions.set(node.id, { x, y });

    const baseSize = getNodeSize(node.label);
    const scaledSize = getScaledNodeSize(baseSize, nodeCount, node.label);
    const graphHealth = getNodeGraphHealth(node);

    // Structural nodes keep their type-based color
    graph.addNode(node.id, {
      x,
      y,
      size: scaledSize,
      color: getNodeColor(node.label),
      label: node.properties.name,
      nodeType: node.label,
      filePath: node.properties.filePath,
      startLine: node.properties.startLine,
      endLine: node.properties.endLine,
      hidden: false,
      mass: getNodeMass(node.label, nodeCount),
      topologyStatus: graphHealth?.topologyStatus,
      expectedIsolationReasons: graphHealth?.expectedIsolationReasons,
      diagnostics: graphHealth?.diagnostics,
      confidence: graphHealth?.confidence,
    });
  });

  // Process remaining nodes in HIERARCHY ORDER (parents before children)
  // Use BFS starting from structural nodes to ensure parents are positioned first
  const addNodeWithPosition = (nodeId: string) => {
    if (graph.hasNode(nodeId)) return;

    const node = nodeMap.get(nodeId);
    if (!node) return;

    let x: number, y: number;

    // Check if this is a symbol node with a community assignment
    const communityIndex = communityMemberships?.get(nodeId);
    const symbolTypes = new Set<string>(COMMUNITY_COLORED_NODE_LABELS);
    const clusterCenter =
      communityIndex !== undefined ? clusterCenters.get(communityIndex) : null;

    if (clusterCenter && symbolTypes.has(node.label)) {
      // CLUSTER-BASED POSITIONING: Position near cluster center with tight jitter
      x = clusterCenter.x + (Math.random() - 0.5) * clusterJitter;
      y = clusterCenter.y + (Math.random() - 0.5) * clusterJitter;
    } else {
      // HIERARCHY-BASED POSITIONING: Position near parent
      const parentId = childToParent.get(nodeId);
      const parentPos = parentId ? nodePositions.get(parentId) : null;

      if (parentPos) {
        x = parentPos.x + (Math.random() - 0.5) * childJitter;
        y = parentPos.y + (Math.random() - 0.5) * childJitter;
      } else {
        // No parent found - position randomly but still spread out
        x = (Math.random() - 0.5) * structuralSpread * 0.5;
        y = (Math.random() - 0.5) * structuralSpread * 0.5;
      }
    }

    nodePositions.set(nodeId, { x, y });

    const baseSize = getNodeSize(node.label);
    const scaledSize = getScaledNodeSize(baseSize, nodeCount, node.label);

    // Check if this node has a community assignment (reuse communityIndex from above)
    const hasCommunity = communityIndex !== undefined;

    // Symbol nodes get colored by community if available
    const usesCommunityColor = hasCommunity && symbolTypes.has(node.label);
    const nodeColor = usesCommunityColor
      ? getCommunityColor(communityIndex!)
      : getNodeColor(node.label);
    const graphHealth = getNodeGraphHealth(node);

    graph.addNode(nodeId, {
      x,
      y,
      size: scaledSize,
      color: nodeColor,
      label: node.properties.name,
      nodeType: node.label,
      filePath: node.properties.filePath,
      startLine: node.properties.startLine,
      endLine: node.properties.endLine,
      hidden: false,
      mass: getNodeMass(node.label, nodeCount),
      community: communityIndex,
      communityColor: hasCommunity
        ? getCommunityColor(communityIndex!)
        : undefined,
      topologyStatus: graphHealth?.topologyStatus,
      expectedIsolationReasons: graphHealth?.expectedIsolationReasons,
      diagnostics: graphHealth?.diagnostics,
      confidence: graphHealth?.confidence,
    });
  };

  // BFS from structural nodes - this ensures parent is ALWAYS positioned before child
  const queue: string[] = [...structuralNodes.map((n) => n.id)];
  const visited = new Set<string>(queue);

  while (queue.length > 0) {
    const currentId = queue.shift()!;

    // Get children of current node and add them
    const children = parentToChildren.get(currentId) || [];
    for (const childId of children) {
      if (!visited.has(childId)) {
        visited.add(childId);
        addNodeWithPosition(childId);
        queue.push(childId); // Add to queue so its children are processed too
      }
    }
  }

  // Add any orphan nodes that weren't reached (no parent relationship)
  knowledgeGraph.nodes.forEach((node) => {
    if (!graph.hasNode(node.id)) {
      addNodeWithPosition(node.id);
    }
  });

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
