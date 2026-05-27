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
}: {
  targetNodeId: string;
  currentSelectedNodeId: string | null;
  nodeX: number;
  nodeY: number;
  nodeSize: number;
  currentCameraRatio: number;
  scaler: GraphScaleSizeScaler;
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
      }),
    },
    durationMs: GRAPH_DETAIL_FOCUS_CAMERA_DURATION_MS,
    shouldUpdateSelection: currentSelectedNodeId !== targetNodeId,
    shouldAnimateCamera: true,
  };
};
