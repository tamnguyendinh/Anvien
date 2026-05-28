import type { GraphScaleSizeScaler } from './graph-scale-model';
import {
  getDetailFocusCameraRatio,
  getRenderedNodeRadiusPx,
} from './graph-scale-model';

export type GraphCameraMode = 'overview' | 'detail-focus';

export type GraphOverviewCameraAction = {
  mode: 'overview';
  durationMs: number;
};

export type GraphDetailFocusCameraAction = {
  mode: 'detail-focus';
  cameraState: {
    x: number;
    y: number;
    ratio: number;
  };
  durationMs: number;
  shouldUpdateSelection: boolean;
  shouldAnimateCamera: true;
};

export const GRAPH_OVERVIEW_CAMERA_RESET_DURATION_MS = 500;
export const GRAPH_DETAIL_FOCUS_CAMERA_DURATION_MS = 400;

export const buildOverviewCameraAction = (): GraphOverviewCameraAction => ({
  mode: 'overview',
  durationMs: GRAPH_OVERVIEW_CAMERA_RESET_DURATION_MS,
});

export const buildDetailFocusCameraAction = ({
  targetNodeId,
  currentSelectedNodeId,
  nodeX,
  nodeY,
  nodeSize,
  currentCameraRatio,
  scaler,
  currentRequiredCenterDistanceGraph,
  minimumGraphCenterDistance,
}: {
  targetNodeId: string;
  currentSelectedNodeId: string | null;
  nodeX: number;
  nodeY: number;
  nodeSize: number;
  currentCameraRatio: number;
  scaler: GraphScaleSizeScaler;
  currentRequiredCenterDistanceGraph?: number;
  minimumGraphCenterDistance?: number;
}): GraphDetailFocusCameraAction => {
  const currentRenderedNodeRadiusPx = getRenderedNodeRadiusPx(
    scaler,
    nodeSize,
    currentCameraRatio,
  );

  return {
    mode: 'detail-focus',
    cameraState: {
      x: nodeX,
      y: nodeY,
      ratio: getDetailFocusCameraRatio({
        currentCameraRatio,
        currentRenderedNodeRadiusPx,
        currentRequiredCenterDistanceGraph,
        minimumGraphCenterDistance,
      }),
    },
    durationMs: GRAPH_DETAIL_FOCUS_CAMERA_DURATION_MS,
    shouldUpdateSelection: currentSelectedNodeId !== targetNodeId,
    shouldAnimateCamera: true,
  };
};
