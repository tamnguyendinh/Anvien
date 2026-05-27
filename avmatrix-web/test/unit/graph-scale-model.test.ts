import { describe, expect, it } from 'vitest';
import {
  buildGraphScaleModel,
  getPolicyDenseRenderedNodeRadiusPx,
  getPolicyOverviewRenderedNodeRadiusPx,
  getDetailFocusCameraRatio,
  getRequiredCenterDistanceGraph,
  getRequiredCenterDistancePx,
  getRequiredEdgeGapPx,
  getRenderedNodeRadiusPx,
  measureGraphUnitToViewportPx,
} from '../../src/lib/graph-scale-model';

const createScaleSizer = () => ({
  scaleSize: (size: number, cameraRatio: number) => size / Math.sqrt(cameraRatio),
});

describe('graph scale model', () => {
  it('exposes rendered node radius caps through policy accessors', () => {
    expect(getPolicyOverviewRenderedNodeRadiusPx()).toBe(3);
    expect(getPolicyDenseRenderedNodeRadiusPx()).toBe(3);
  });

  it('derives rendered node radius through Sigma scaleSize semantics', () => {
    const scaler = createScaleSizer();

    expect(getRenderedNodeRadiusPx(scaler, 3, 1)).toBe(3);
    expect(getRenderedNodeRadiusPx(scaler, 3, 0.25)).toBe(6);
  });

  it('derives detail focus camera ratio from current rendered node radius', () => {
    expect(
      getDetailFocusCameraRatio({
        currentCameraRatio: 1,
        currentRenderedNodeRadiusPx: 3,
        targetRenderedNodeRadiusPx: 8,
      }),
    ).toBeCloseTo(0.140625);
  });

  it('enforces a one-diameter edge gap for equal-size nodes', () => {
    const radius = 3;

    expect(getRequiredEdgeGapPx(radius, radius)).toBe(6);
    expect(getRequiredCenterDistancePx(radius, radius)).toBe(12);
  });

  it('converts viewport-pixel center distance into graph units', () => {
    const requiredCenterDistancePx = 12;
    const graphUnitToViewportPx = 4;

    expect(
      getRequiredCenterDistanceGraph(
        requiredCenterDistancePx,
        graphUnitToViewportPx,
      ),
    ).toBe(3);
  });

  it('measures graph-unit to viewport-pixel scale from graph projection', () => {
    const projection = {
      graphToViewport: (point: { x: number; y: number }) => ({
        x: point.x * 5 + 20,
        y: point.y * 5 + 10,
      }),
    };

    expect(measureGraphUnitToViewportPx(projection)).toBe(5);
  });

  it('builds a scale model from camera ratio, node sizes, and projection', () => {
    const sigma = {
      ...createScaleSizer(),
      getCamera: () => ({
        getState: () => ({ x: 0.5, y: 0.5, ratio: 0.25, angle: 0 }),
      }),
      graphToViewport: (point: { x: number; y: number }) => ({
        x: point.x * 2,
        y: point.y * 2,
      }),
    };

    const model = buildGraphScaleModel(sigma, [2, 3]);

    expect(model.cameraRatio).toBe(0.25);
    expect(model.graphUnitToViewportPx).toBe(2);
    expect(model.minRenderedRadiusPx).toBe(4);
    expect(model.maxRenderedRadiusPx).toBe(6);
    expect(model.maxRenderedDiameterPx).toBe(12);
    expect(model.requiredEdgeGapPx).toBe(12);
    expect(model.requiredCenterDistancePx).toBe(24);
    expect(model.requiredCenterDistanceGraph).toBe(12);
  });
});
