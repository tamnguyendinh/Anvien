import type Sigma from 'sigma';

export type GraphScalePoint = {
  x: number;
  y: number;
};

export type GraphScaleProjection = {
  graphToViewport: (point: GraphScalePoint) => GraphScalePoint;
};

export type GraphScaleSizeScaler = {
  scaleSize: (size: number, cameraRatio: number) => number;
};

export type GraphScalePolicy = {
  maxRenderedNodeRadiusPx: number;
  maxDenseRenderedNodeRadiusPx: number;
  requiredEdgeGapDiameterMultiplier: number;
  viewportProbeGraphUnits: number;
  minimumGraphUnitToViewportPx: number;
};

export type GraphScaleModel = {
  cameraRatio: number;
  graphUnitToViewportPx: number;
  minRenderedRadiusPx: number;
  maxRenderedRadiusPx: number;
  maxRenderedDiameterPx: number;
  requiredEdgeGapPx: number;
  requiredCenterDistancePx: number;
  requiredCenterDistanceGraph: number;
};

export const GRAPH_MAX_RENDERED_NODE_RADIUS_PX = 3;
export const GRAPH_MAX_DENSE_RENDERED_NODE_RADIUS_PX = 3;
export const GRAPH_REQUIRED_EDGE_GAP_DIAMETER_MULTIPLIER = 1;
export const GRAPH_VIEWPORT_PROBE_GRAPH_UNITS = 1;
export const GRAPH_MINIMUM_GRAPH_UNIT_TO_VIEWPORT_PX = 0.000001;
export const GRAPH_DETAIL_FOCUS_RENDERED_NODE_RADIUS_PX = 8;
export const GRAPH_MINIMUM_DETAIL_FOCUS_CAMERA_RATIO = 0.002;

export const GRAPH_RENDER_SCALE_POLICY: GraphScalePolicy = {
  maxRenderedNodeRadiusPx: GRAPH_MAX_RENDERED_NODE_RADIUS_PX,
  maxDenseRenderedNodeRadiusPx: GRAPH_MAX_DENSE_RENDERED_NODE_RADIUS_PX,
  requiredEdgeGapDiameterMultiplier:
    GRAPH_REQUIRED_EDGE_GAP_DIAMETER_MULTIPLIER,
  viewportProbeGraphUnits: GRAPH_VIEWPORT_PROBE_GRAPH_UNITS,
  minimumGraphUnitToViewportPx: GRAPH_MINIMUM_GRAPH_UNIT_TO_VIEWPORT_PX,
};

export const getPolicyOverviewRenderedNodeRadiusPx = (
  policy: GraphScalePolicy = GRAPH_RENDER_SCALE_POLICY,
): number => policy.maxRenderedNodeRadiusPx;

export const getPolicyDenseRenderedNodeRadiusPx = (
  policy: GraphScalePolicy = GRAPH_RENDER_SCALE_POLICY,
): number => policy.maxDenseRenderedNodeRadiusPx;

export const getPolicyMaxRenderedNodeRadiusPx = (
  _nodeCount: number,
  policy: GraphScalePolicy = GRAPH_RENDER_SCALE_POLICY,
): number => getPolicyDenseRenderedNodeRadiusPx(policy);

export const capRenderedNodeRadiusByPolicy = (
  radiusPx: number,
  nodeCount: number,
  policy: GraphScalePolicy = GRAPH_RENDER_SCALE_POLICY,
): number => Math.min(radiusPx, getPolicyMaxRenderedNodeRadiusPx(nodeCount, policy));

export const getRenderedNodeRadiusPx = (
  scaler: GraphScaleSizeScaler,
  size: number,
  cameraRatio: number,
): number => scaler.scaleSize(size, cameraRatio);

export const getDetailFocusCameraRatio = ({
  currentCameraRatio,
  currentRenderedNodeRadiusPx,
  targetRenderedNodeRadiusPx = GRAPH_DETAIL_FOCUS_RENDERED_NODE_RADIUS_PX,
  minimumCameraRatio = GRAPH_MINIMUM_DETAIL_FOCUS_CAMERA_RATIO,
  currentRequiredCenterDistanceGraph,
  minimumGraphCenterDistance,
}: {
  currentCameraRatio: number;
  currentRenderedNodeRadiusPx: number;
  targetRenderedNodeRadiusPx?: number;
  minimumCameraRatio?: number;
  currentRequiredCenterDistanceGraph?: number;
  minimumGraphCenterDistance?: number;
}): number => {
  if (
    !Number.isFinite(currentCameraRatio) ||
    currentCameraRatio <= 0 ||
    !Number.isFinite(currentRenderedNodeRadiusPx) ||
    currentRenderedNodeRadiusPx <= 0 ||
    !Number.isFinite(targetRenderedNodeRadiusPx) ||
    targetRenderedNodeRadiusPx <= 0
  ) {
    return Math.max(minimumCameraRatio, currentCameraRatio || 1);
  }

  const readableCameraRatio = Math.max(
    minimumCameraRatio,
    currentCameraRatio *
      Math.pow(currentRenderedNodeRadiusPx / targetRenderedNodeRadiusPx, 2),
  );

  if (
    !Number.isFinite(currentRequiredCenterDistanceGraph) ||
    !Number.isFinite(minimumGraphCenterDistance) ||
    currentRequiredCenterDistanceGraph === undefined ||
    minimumGraphCenterDistance === undefined ||
    currentRequiredCenterDistanceGraph <= 0 ||
    minimumGraphCenterDistance <= 0 ||
    currentRequiredCenterDistanceGraph <= minimumGraphCenterDistance
  ) {
    return readableCameraRatio;
  }

  const spacingSafeCameraRatio =
    currentCameraRatio *
    Math.pow(minimumGraphCenterDistance / currentRequiredCenterDistanceGraph, 2);

  return Math.max(
    minimumCameraRatio,
    Math.min(readableCameraRatio, spacingSafeCameraRatio),
  );
};

export const getRequiredEdgeGapPx = (
  leftRadiusPx: number,
  rightRadiusPx: number,
  policy: GraphScalePolicy = GRAPH_RENDER_SCALE_POLICY,
): number =>
  Math.max(leftRadiusPx * 2, rightRadiusPx * 2) *
  policy.requiredEdgeGapDiameterMultiplier;

export const getRequiredCenterDistancePx = (
  leftRadiusPx: number,
  rightRadiusPx: number,
  policy: GraphScalePolicy = GRAPH_RENDER_SCALE_POLICY,
): number =>
  leftRadiusPx +
  rightRadiusPx +
  getRequiredEdgeGapPx(leftRadiusPx, rightRadiusPx, policy);

export const measureGraphUnitToViewportPx = (
  projection: GraphScaleProjection,
  probePoint: GraphScalePoint = { x: 0, y: 0 },
  policy: GraphScalePolicy = GRAPH_RENDER_SCALE_POLICY,
): number => {
  const probeDistance = policy.viewportProbeGraphUnits;
  const left = projection.graphToViewport(probePoint);
  const right = projection.graphToViewport({
    x: probePoint.x + probeDistance,
    y: probePoint.y,
  });
  const measured = Math.hypot(right.x - left.x, right.y - left.y);
  if (!Number.isFinite(measured) || measured <= 0) {
    return policy.minimumGraphUnitToViewportPx;
  }
  return Math.max(measured / probeDistance, policy.minimumGraphUnitToViewportPx);
};

export const getRequiredCenterDistanceGraph = (
  requiredCenterDistancePx: number,
  graphUnitToViewportPx: number,
): number => requiredCenterDistancePx / Math.max(graphUnitToViewportPx, 0.000001);

export const buildGraphScaleModel = (
  sigma: Pick<Sigma, 'scaleSize' | 'graphToViewport' | 'getCamera'>,
  nodeSizes: readonly number[],
  policy: GraphScalePolicy = GRAPH_RENDER_SCALE_POLICY,
  probePoint: GraphScalePoint = { x: 0, y: 0 },
): GraphScaleModel => {
  const cameraRatio = sigma.getCamera().getState().ratio;
  const renderedRadii = nodeSizes
    .map((size) => getRenderedNodeRadiusPx(sigma, size, cameraRatio))
    .filter((radius) => Number.isFinite(radius) && radius > 0);
  const minRenderedRadiusPx =
    renderedRadii.length > 0 ? Math.min(...renderedRadii) : 0;
  const maxRenderedRadiusPx =
    renderedRadii.length > 0 ? Math.max(...renderedRadii) : 0;
  const requiredEdgeGapPx = getRequiredEdgeGapPx(
    maxRenderedRadiusPx,
    maxRenderedRadiusPx,
    policy,
  );
  const requiredCenterDistancePx = getRequiredCenterDistancePx(
    maxRenderedRadiusPx,
    maxRenderedRadiusPx,
    policy,
  );
  const graphUnitToViewportPx = measureGraphUnitToViewportPx(
    sigma,
    probePoint,
    policy,
  );

  return {
    cameraRatio,
    graphUnitToViewportPx,
    minRenderedRadiusPx,
    maxRenderedRadiusPx,
    maxRenderedDiameterPx: maxRenderedRadiusPx * 2,
    requiredEdgeGapPx,
    requiredCenterDistancePx,
    requiredCenterDistanceGraph: getRequiredCenterDistanceGraph(
      requiredCenterDistancePx,
      graphUnitToViewportPx,
    ),
  };
};
