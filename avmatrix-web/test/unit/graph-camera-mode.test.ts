import { describe, expect, it } from 'vitest';
import {
  buildDetailFocusCameraAction,
  buildOverviewCameraAction,
} from '../../src/lib/graph-camera-mode';

const createScaleSizer = () => ({
  scaleSize: (size: number, cameraRatio: number) => size / Math.sqrt(cameraRatio),
});

describe('graph camera mode', () => {
  it('keeps graph load in overview mode', () => {
    const action = buildOverviewCameraAction();

    expect(action.mode).toBe('overview');
    expect(action.durationMs).toBe(500);
  });

  it('uses detail focus mode for explicit node focus', () => {
    const action = buildDetailFocusCameraAction({
      targetNodeId: 'Function:src/app.ts:main',
      currentSelectedNodeId: null,
      nodeX: 42,
      nodeY: 24,
      nodeSize: 3,
      currentCameraRatio: 1,
      scaler: createScaleSizer(),
    });

    expect(action.mode).toBe('detail-focus');
    expect(action.shouldUpdateSelection).toBe(true);
    expect(action.shouldAnimateCamera).toBe(true);
    expect(action.cameraState).toEqual({
      x: 42,
      y: 24,
      ratio: expect.closeTo(0.140625),
    });
  });

  it('still animates camera when focusing the already selected node', () => {
    const action = buildDetailFocusCameraAction({
      targetNodeId: 'Function:src/app.ts:main',
      currentSelectedNodeId: 'Function:src/app.ts:main',
      nodeX: 42,
      nodeY: 24,
      nodeSize: 3,
      currentCameraRatio: 0.25,
      scaler: createScaleSizer(),
    });

    expect(action.shouldUpdateSelection).toBe(false);
    expect(action.shouldAnimateCamera).toBe(true);
    expect(action.cameraState.ratio).toBeCloseTo(0.140625);
  });

  it('uses the spacing-safe ratio when detail focus would otherwise violate layout spacing', () => {
    const action = buildDetailFocusCameraAction({
      targetNodeId: 'Function:src/app.ts:main',
      currentSelectedNodeId: null,
      nodeX: 42,
      nodeY: 24,
      nodeSize: 3,
      currentCameraRatio: 1,
      scaler: createScaleSizer(),
      currentRequiredCenterDistanceGraph: 48,
      minimumGraphCenterDistance: 12,
    });

    expect(action.cameraState.ratio).toBeCloseTo(0.0625);
  });
});
