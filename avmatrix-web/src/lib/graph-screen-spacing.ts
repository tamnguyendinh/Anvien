import type Sigma from 'sigma';
import type Graph from 'graphology';
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from './graph-adapter';

export type ScreenNodeSpacingDiagnostics = {
  coordinateSpace: 'viewport_px';
  nodeCount: number;
  islandCount: number;
  viewportWidth: number;
  viewportHeight: number;
  cameraRatio: number;
  minRenderedRadius: number;
  maxRenderedRadius: number;
  maxRenderedDiameter: number;
  minObservedCenterDistance: number;
  minObservedEdgeGap: number;
  maxRequiredCenterDistance: number;
  overlapCount: number;
  targetGapViolationCount: number;
};

type ScreenNode = {
  x: number;
  y: number;
  radius: number;
  islandKey: string;
};

const getPairRequiredEdgeGap = (left: ScreenNode, right: ScreenNode): number =>
  Math.max(left.radius * 2, right.radius * 2);

const getPairRequiredCenterDistance = (
  left: ScreenNode,
  right: ScreenNode,
): number => left.radius + right.radius + getPairRequiredEdgeGap(left, right);

const getGridKey = (node: ScreenNode, cellSize: number): string =>
  `${Math.floor(node.x / cellSize)}:${Math.floor(node.y / cellSize)}`;

export const buildScreenNodeSpacingDiagnostics = (
  sigma: Sigma,
): ScreenNodeSpacingDiagnostics => {
  const graph = sigma.getGraph() as Graph<
    SigmaNodeAttributes,
    SigmaEdgeAttributes
  >;
  const dimensions = sigma.getDimensions();
  const cameraRatio = sigma.getCamera().getState().ratio;
  const nodes: ScreenNode[] = [];
  const islands = new Set<string>();
  let minRenderedRadius = Number.POSITIVE_INFINITY;
  let maxRenderedRadius = 0;

  for (const nodeId of graph.nodes()) {
    const attributes = graph.getNodeAttributes(nodeId);
    if (attributes.hidden) continue;

    const displayData = sigma.getNodeDisplayData(nodeId);
    if (displayData?.hidden) continue;

    const viewport = sigma.graphToViewport({
      x: attributes.x,
      y: attributes.y,
    });
    const radius = sigma.scaleSize(
      typeof displayData?.size === 'number' ? displayData.size : attributes.size,
      cameraRatio,
    );
    const islandKey = `${attributes.appLayerRing ?? 'missing_app_layer'}:${
      attributes.islandKey ?? attributes.nodeType
    }`;

    islands.add(islandKey);
    minRenderedRadius = Math.min(minRenderedRadius, radius);
    maxRenderedRadius = Math.max(maxRenderedRadius, radius);
    nodes.push({
      x: viewport.x,
      y: viewport.y,
      radius,
      islandKey,
    });
  }

  const maxRenderedDiameter = maxRenderedRadius * 2;
  const maxPossibleRequiredCenterDistance = maxRenderedRadius * 4;
  const gridCellSize = Math.max(1, maxPossibleRequiredCenterDistance);
  const neighborCellRadius =
    Math.ceil(maxPossibleRequiredCenterDistance / gridCellSize) + 1;
  const gridsByIsland = new Map<string, Map<string, ScreenNode[]>>();
  let minObservedCenterDistance = Number.POSITIVE_INFINITY;
  let minObservedEdgeGap = Number.POSITIVE_INFINITY;
  let maxRequiredCenterDistance = 0;
  let overlapCount = 0;
  let targetGapViolationCount = 0;

  for (const node of nodes) {
    const grid = gridsByIsland.get(node.islandKey) ?? new Map();
    const cellX = Math.floor(node.x / gridCellSize);
    const cellY = Math.floor(node.y / gridCellSize);

    for (
      let x = cellX - neighborCellRadius;
      x <= cellX + neighborCellRadius;
      x++
    ) {
      for (
        let y = cellY - neighborCellRadius;
        y <= cellY + neighborCellRadius;
        y++
      ) {
        const neighbors = grid.get(`${x}:${y}`);
        if (!neighbors) continue;

        for (const neighbor of neighbors) {
          const centerDistance = Math.hypot(
            node.x - neighbor.x,
            node.y - neighbor.y,
          );
          const overlapDistance = node.radius + neighbor.radius;
          const requiredCenterDistance = getPairRequiredCenterDistance(
            node,
            neighbor,
          );
          const edgeGap = centerDistance - overlapDistance;

          minObservedCenterDistance = Math.min(
            minObservedCenterDistance,
            centerDistance,
          );
          minObservedEdgeGap = Math.min(minObservedEdgeGap, edgeGap);
          maxRequiredCenterDistance = Math.max(
            maxRequiredCenterDistance,
            requiredCenterDistance,
          );
          if (centerDistance < overlapDistance) {
            overlapCount++;
          }
          if (centerDistance < requiredCenterDistance) {
            targetGapViolationCount++;
          }
        }
      }
    }

    const key = getGridKey(node, gridCellSize);
    const cell = grid.get(key) ?? [];
    cell.push(node);
    grid.set(key, cell);
    gridsByIsland.set(node.islandKey, grid);
  }

  const comparedPairs = Number.isFinite(minObservedCenterDistance);

  return {
    coordinateSpace: 'viewport_px',
    nodeCount: nodes.length,
    islandCount: islands.size,
    viewportWidth: dimensions.width,
    viewportHeight: dimensions.height,
    cameraRatio,
    minRenderedRadius: Number.isFinite(minRenderedRadius)
      ? minRenderedRadius
      : 0,
    maxRenderedRadius,
    maxRenderedDiameter,
    minObservedCenterDistance: comparedPairs
      ? minObservedCenterDistance
      : maxPossibleRequiredCenterDistance,
    minObservedEdgeGap: comparedPairs ? minObservedEdgeGap : maxRenderedDiameter,
    maxRequiredCenterDistance:
      maxRequiredCenterDistance || maxPossibleRequiredCenterDistance,
    overlapCount,
    targetGapViolationCount,
  };
};
