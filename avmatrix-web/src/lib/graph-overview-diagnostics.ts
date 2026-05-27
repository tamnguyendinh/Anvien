import type Sigma from 'sigma';
import type Graph from 'graphology';
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from './graph-adapter';
import type { GraphOverviewDiagnosticsInput } from './runtime-diagnostics';

const getIslandKey = (attributes: SigmaNodeAttributes): string =>
  `${attributes.appLayerRing ?? 'missing_app_layer'}:${
    attributes.islandKey ?? attributes.nodeType
  }`;

const getRingKey = (attributes: SigmaNodeAttributes): string =>
  attributes.appLayerRing ?? 'missing_app_layer';

const incrementCount = (
  counts: Record<string, number>,
  key: string | undefined,
): void => {
  if (!key) return;
  counts[key] = (counts[key] ?? 0) + 1;
};

const sortedKeys = (counts: Record<string, number>): string[] =>
  Object.keys(counts).sort((left, right) => left.localeCompare(right));

export const buildGraphOverviewDiagnostics = (
  sigma: Sigma,
): GraphOverviewDiagnosticsInput => {
  const graph = sigma.getGraph() as Graph<
    SigmaNodeAttributes,
    SigmaEdgeAttributes
  >;
  const dimensions = sigma.getDimensions();
  const cameraState = sigma.getCamera().getState();
  const visibleColorCounts: Record<string, number> = {};
  const visibleRingCounts: Record<string, number> = {};
  const visibleIslandCounts: Record<string, number> = {};
  const visibleNodeTypeCounts: Record<string, number> = {};
  const graphRingCounts: Record<string, number> = {};
  const graphIslandCounts: Record<string, number> = {};
  const graphNodeTypeCounts: Record<string, number> = {};
  let visibleViewportNodeCount = 0;

  for (const nodeId of graph.nodes()) {
    const attributes = graph.getNodeAttributes(nodeId);
    if (attributes.hidden) continue;

    incrementCount(graphNodeTypeCounts, attributes.nodeType);
    incrementCount(graphRingCounts, getRingKey(attributes));
    incrementCount(graphIslandCounts, getIslandKey(attributes));

    const displayData = sigma.getNodeDisplayData(nodeId);
    if (displayData?.hidden) continue;

    const viewport = sigma.graphToViewport({
      x: attributes.x,
      y: attributes.y,
    });
    if (
      viewport.x < 0 ||
      viewport.x > dimensions.width ||
      viewport.y < 0 ||
      viewport.y > dimensions.height
    ) {
      continue;
    }

    visibleViewportNodeCount++;
    incrementCount(
      visibleColorCounts,
      (displayData?.color ?? attributes.color)?.toLowerCase(),
    );
    incrementCount(visibleRingCounts, getRingKey(attributes));
    incrementCount(visibleIslandCounts, getIslandKey(attributes));
    incrementCount(visibleNodeTypeCounts, attributes.nodeType);
  }

  const [dominantIslandKey, dominantIslandCount = 0] =
    Object.entries(visibleIslandCounts).sort(
      (left, right) => right[1] - left[1] || left[0].localeCompare(right[0]),
    )[0] ?? ['', 0];

  return {
    nodeCount: graph.order,
    viewportWidth: dimensions.width,
    viewportHeight: dimensions.height,
    visibleViewportNodeCount,
    visibleColorCount: Object.keys(visibleColorCounts).length,
    visibleRingCount: Object.keys(visibleRingCounts).length,
    visibleIslandCount: Object.keys(visibleIslandCounts).length,
    dominantIslandKey,
    dominantIslandShare:
      visibleViewportNodeCount > 0
        ? dominantIslandCount / visibleViewportNodeCount
        : 0,
    visibleColorCounts,
    visibleRingCounts,
    visibleIslandCounts,
    visibleNodeTypeCounts,
    graphRingCounts,
    graphIslandCounts,
    graphNodeTypeCounts,
    visibleRingInventory: sortedKeys(visibleRingCounts),
    visibleNodeTypeInventory: sortedKeys(visibleNodeTypeCounts),
    graphRingInventory: sortedKeys(graphRingCounts),
    graphIslandInventory: sortedKeys(graphIslandCounts),
    filterNodeTypeInventory: sortedKeys(graphNodeTypeCounts),
    cameraRatio: cameraState.ratio,
    cameraX: cameraState.x,
    cameraY: cameraState.y,
  };
};
