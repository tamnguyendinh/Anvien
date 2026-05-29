import type Sigma from 'sigma';
import type Graph from 'graphology';
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from './graph-adapter';
import {
  buildScreenNodeSpacingDiagnostics,
  type ScreenNodeSpacingDiagnostics,
} from './graph-screen-spacing';
import { recordReadableCamera } from './runtime-diagnostics';

export const READABLE_GRAPH_NODE_COUNT_THRESHOLD = 1000;
export const MIN_READABLE_NODE_RADIUS_PX = 2;
export const MIN_READABLE_CAMERA_RATIO = 0.00004;

export type ReadableGraphCameraState = {
  x: number;
  y: number;
  ratio: number;
  angle: number;
  focusedIslandKey: string;
  focusedIslandNodeCount: number;
  previousDiagnostics: ScreenNodeSpacingDiagnostics;
  focusRawX: number;
  focusRawY: number;
  focusCellNodeCount: number;
};

type IslandSummary = {
  key: string;
  count: number;
  nodes: Array<{ x: number; y: number }>;
};

type FocusCell = {
  count: number;
  sumX: number;
  sumY: number;
};

type FocusCenter = {
  x: number;
  y: number;
  nodeCount: number;
};

const getIslandKey = (attributes: SigmaNodeAttributes): string =>
  `${attributes.appLayerRing ?? 'missing_app_layer'}:${
    attributes.islandKey ?? attributes.nodeType
  }`;

const getDensestVisibleIsland = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
): IslandSummary | null => {
  const islands = new Map<string, IslandSummary>();

  for (const nodeId of graph.nodes()) {
    const attributes = graph.getNodeAttributes(nodeId);
    if (attributes.hidden) continue;

    const key = getIslandKey(attributes);
    const island = islands.get(key) ?? {
      key,
      count: 0,
      nodes: [],
    };
    island.count++;
    island.nodes.push({ x: attributes.x, y: attributes.y });
    islands.set(key, island);
  }

  return [...islands.values()].sort(
    (left, right) => right.count - left.count || left.key.localeCompare(right.key),
  )[0] ?? null;
};

const getReadableFocusCenter = (
  sigma: Sigma,
  island: IslandSummary,
  ratio: number,
): FocusCenter => {
  const dimensions = sigma.getDimensions();
  const projectedCameraState = {
    ...sigma.getCamera().getState(),
    ratio,
  };
  const projectedTopLeft = sigma.viewportToGraph(
    { x: 0, y: 0 },
    { cameraState: projectedCameraState },
  );
  const projectedBottomRight = sigma.viewportToGraph(
    { x: dimensions.width, y: dimensions.height },
    { cameraState: projectedCameraState },
  );
  const projectedGraphWidth = Math.abs(
    projectedBottomRight.x - projectedTopLeft.x,
  );
  const projectedGraphHeight = Math.abs(
    projectedBottomRight.y - projectedTopLeft.y,
  );
  const cellWidth = Math.max(1, projectedGraphWidth);
  const cellHeight = Math.max(1, projectedGraphHeight);
  const cells = new Map<string, FocusCell>();

  for (const node of island.nodes) {
    const key = `${Math.floor(node.x / cellWidth)}:${Math.floor(
      node.y / cellHeight,
    )}`;
    const cell = cells.get(key) ?? { count: 0, sumX: 0, sumY: 0 };
    cell.count++;
    cell.sumX += node.x;
    cell.sumY += node.y;
    cells.set(key, cell);
  }

  const bestCell = [...cells.values()].sort(
    (left, right) => right.count - left.count,
  )[0];
  if (!bestCell || bestCell.count === 0) {
    const fallback = island.nodes[0] ?? { x: 0, y: 0 };
    return { ...fallback, nodeCount: 0 };
  }

  return {
    x: bestCell.sumX / bestCell.count,
    y: bestCell.sumY / bestCell.count,
    nodeCount: bestCell.count,
  };
};

export const buildReadableGraphCameraState = (
  sigma: Sigma,
): ReadableGraphCameraState | null => {
  const graph = sigma.getGraph() as Graph<
    SigmaNodeAttributes,
    SigmaEdgeAttributes
  >;
  const diagnostics = buildScreenNodeSpacingDiagnostics(sigma);

  if (
    diagnostics.nodeCount < READABLE_GRAPH_NODE_COUNT_THRESHOLD ||
    diagnostics.maxRenderedRadius >= MIN_READABLE_NODE_RADIUS_PX ||
    diagnostics.maxRenderedRadius <= 0
  ) {
    return null;
  }

  const island = getDensestVisibleIsland(graph);
  if (!island || island.count === 0) return null;

  const targetRatio =
    diagnostics.cameraRatio *
    Math.pow(diagnostics.maxRenderedRadius / MIN_READABLE_NODE_RADIUS_PX, 2);
  const ratio = Math.max(
    MIN_READABLE_CAMERA_RATIO,
    Math.min(diagnostics.cameraRatio, targetRatio),
  );
  const rawCenter = getReadableFocusCenter(sigma, island, ratio);
  const framedCenter = sigma.viewportToFramedGraph(
    sigma.graphToViewport(rawCenter),
  );

  return {
    x: framedCenter.x,
    y: framedCenter.y,
    ratio,
    angle: 0,
    focusedIslandKey: island.key,
    focusedIslandNodeCount: island.count,
    previousDiagnostics: diagnostics,
    focusRawX: rawCenter.x,
    focusRawY: rawCenter.y,
    focusCellNodeCount: rawCenter.nodeCount,
  };
};

export const applyReadableGraphCamera = (sigma: Sigma): boolean => {
  const state = buildReadableGraphCameraState(sigma);
  if (!state) return false;

  sigma.getCamera().setState({
    x: state.x,
    y: state.y,
    ratio: state.ratio,
    angle: state.angle,
  });
  recordReadableCamera({
    applied: true,
    focusedIslandKey: state.focusedIslandKey,
    focusedIslandNodeCount: state.focusedIslandNodeCount,
    focusCellNodeCount: state.focusCellNodeCount,
    focusRawX: state.focusRawX,
    focusRawY: state.focusRawY,
    cameraX: state.x,
    cameraY: state.y,
    ratio: state.ratio,
    previousMaxRenderedRadius: state.previousDiagnostics.maxRenderedRadius,
  });
  sigma.refresh();
  return true;
};
