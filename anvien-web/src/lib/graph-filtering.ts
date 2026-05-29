import Graph from "graphology";
import {
  graphHealthMatchesFilters,
  type GraphHealthFilterState,
} from "./graph-health-filters";
import {
  semanticMatchesFilters,
  type SemanticFilterState,
} from "./semantic-filters";
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from "./graph-adapter-types";

export const filterGraphByLabels = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
  visibleLabels: string[],
  graphHealthFilters?: GraphHealthFilterState,
  semanticFilters?: SemanticFilterState,
): void => {
  graph.forEachNode((nodeId, attributes) => {
    const isVisible = visibleLabels.includes(attributes.nodeType);
    const isGraphHealthVisible = graphHealthFilters
      ? graphHealthMatchesFilters(attributes, graphHealthFilters)
      : true;
    const isSemanticVisible = semanticFilters
      ? semanticMatchesFilters(attributes, semanticFilters)
      : true;
    graph.setNodeAttribute(
      nodeId,
      "hidden",
      !isVisible || !isGraphHealthVisible || !isSemanticVisible,
    );
  });
};

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

export const filterGraphByDepth = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
  selectedNodeId: string | null,
  maxHops: number | null,
  visibleLabels: string[],
  graphHealthFilters?: GraphHealthFilterState,
  semanticFilters?: SemanticFilterState,
): void => {
  if (maxHops === null) {
    filterGraphByLabels(graph, visibleLabels, graphHealthFilters, semanticFilters);
    return;
  }

  if (selectedNodeId === null || !graph.hasNode(selectedNodeId)) {
    filterGraphByLabels(graph, visibleLabels, graphHealthFilters, semanticFilters);
    return;
  }

  const nodesInRange = getNodesWithinHops(graph, selectedNodeId, maxHops);

  graph.forEachNode((nodeId, attributes) => {
    const isLabelVisible = visibleLabels.includes(attributes.nodeType);
    const isInRange = nodesInRange.has(nodeId);
    const isGraphHealthVisible = graphHealthFilters
      ? graphHealthMatchesFilters(attributes, graphHealthFilters)
      : true;
    const isSemanticVisible = semanticFilters
      ? semanticMatchesFilters(attributes, semanticFilters)
      : true;
    graph.setNodeAttribute(
      nodeId,
      "hidden",
      !isLabelVisible || !isInRange || !isGraphHealthVisible || !isSemanticVisible,
    );
  });
};

