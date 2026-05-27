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

export const READABLE_GRAPH_NODE_COUNT_THRESHOLD = 1000;
export const MIN_READABLE_NODE_RADIUS_PX = 0.75;
export const MIN_READABLE_CAMERA_RATIO = 0.00004;

export type ReadableGraphCameraState = {
  x: number;
  y: number;
  ratio: number;
  angle: number;
  focusedIslandKey: string;
  focusedIslandNodeCount: number;
  previousDiagnostics: ScreenNodeSpacingDiagnostics;
};

type IslandSummary = {
  key: string;
  count: number;
  sumX: number;
  sumY: number;
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
      sumX: 0,
      sumY: 0,
    };
    island.count++;
    island.sumX += attributes.x;
    island.sumY += attributes.y;
    islands.set(key, island);
  }

  return [...islands.values()].sort(
    (left, right) => right.count - left.count || left.key.localeCompare(right.key),
  )[0] ?? null;
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
  const rawCenter = {
    x: island.sumX / island.count,
    y: island.sumY / island.count,
  };
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
  sigma.refresh();
  return true;
};
