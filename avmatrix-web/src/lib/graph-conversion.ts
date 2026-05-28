import Graph, { MultiDirectedGraph } from "graphology";
import type { KnowledgeGraph } from "../core/graph/types";
import {
  EDGE_SIZE_MULTIPLIERS,
  getDisplayGraphRelationships,
  getCommunityColor,
  getEdgeInfo,
  getNodeDisplayLabel,
  getNodeColor,
  getNodeSize,
} from "./constants";
import { getNodeGraphHealth } from "./graph-health-filters";
import { getSemanticFilterableFromNode } from "./semantic-filters";
import {
  getNodeMass,
  getScaledNodeSize,
} from "./graph-node-sizing";
import { applyFilterBasedClusteredLayout } from "./graph-layout";
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from "./graph-adapter-types";

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
    const isResolutionGap = node.label === "ResolutionGap";
    const communityIndex = communityMemberships?.get(node.id);
    const hasCommunity = communityIndex !== undefined;
    const nodeColor = getNodeColor(displayNodeType);
    const graphHealth = getNodeGraphHealth(node);
    const semantic = getSemanticFilterableFromNode(node);

    graph.addNode(node.id, {
      ...semantic,
      x: 0,
      y: 0,
      size: isResolutionGap ? 1 : scaledSize,
      type: isResolutionGap ? "square" : undefined,
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
      resolutionHealthBuckets: semantic.resolutionHealthBuckets,
      resolutionGapCount: semantic.resolutionGapCount,
      resolutionConfidence: semantic.resolutionConfidence,
      community: communityIndex,
      communityColor: hasCommunity
        ? getCommunityColor(communityIndex!)
        : undefined,
    });
  });

  applyFilterBasedClusteredLayout(graph);

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

