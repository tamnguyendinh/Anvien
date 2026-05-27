import { beforeEach, describe, expect, it } from 'vitest';
import {
  getWebRuntimeDiagnostics,
  recordGraphConversion,
  recordLayoutNodeSpacing,
  recordLayoutRings,
  recordManualLayoutOptimizerInvocation,
  recordReconnectBannerState,
  recordScreenNodeSpacing,
  recordVisualScale,
  resetWebRuntimeDiagnostics,
} from '../../src/lib/runtime-diagnostics';

describe('runtime diagnostics', () => {
  beforeEach(() => {
    resetWebRuntimeDiagnostics();
  });

  it('records graph conversion without automatic layout timings', () => {
    recordGraphConversion({
      startedAt: 10,
      finishedAt: 35,
      nodeCount: 100,
      relationshipCount: 250,
    });

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.graphConversion.count).toBe(1);
    expect(diagnostics?.graphConversion.lastMs).toBe(25);
    expect(diagnostics?.graphConversion.lastNodeCount).toBe(100);
    expect(diagnostics?.graphConversion.lastRelationshipCount).toBe(250);
    expect(diagnostics?.layout.starts).toBe(0);
    expect(diagnostics?.layout.stops).toBe(0);
    expect(diagnostics?.layout.manualOptimizerInvocations).toBe(0);
    expect(diagnostics?.layout.isRunning).toBe(false);
  });

  it('records manual layout optimizer invocations separately from layout starts', () => {
    recordManualLayoutOptimizerInvocation({
      nodeCount: 42,
      startedAt: 100,
      finishedAt: 108,
    });

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.layout.starts).toBe(0);
    expect(diagnostics?.layout.stops).toBe(0);
    expect(diagnostics?.layout.manualOptimizerInvocations).toBe(1);
    expect(diagnostics?.layout.lastNodeCount).toBe(42);
    expect(diagnostics?.layout.lastManualOptimizerInvokedAt).toBe(100);
    expect(diagnostics?.layout.lastManualOptimizerRunMs).toBe(8);
  });

  it('records visual-scale bounds for e2e assertions', () => {
    recordVisualScale({
      nodeCount: 20_000,
      minNodeSize: 1.5,
      maxNodeSize: 3,
      maxRenderedNodeSizeCap: 9,
      structuralToLeafRatio: 3,
      maxSizeByLabel: {
        Package: 1.5,
        Section: 1,
      },
    });

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.visualScale.nodeCount).toBe(20_000);
    expect(diagnostics?.visualScale.maxNodeSize).toBe(3);
    expect(diagnostics?.visualScale.maxRenderedNodeSizeCap).toBe(9);
    expect(diagnostics?.visualScale.structuralToLeafRatio).toBe(3);
    expect(diagnostics?.visualScale.maxSizeByLabel.Package).toBe(1.5);
    expect(diagnostics?.visualScale.maxSizeByLabel.Section).toBe(1);
  });

  it('records App Layer ring layout diagnostics for e2e assertions', () => {
    recordLayoutRings({
      nodeCount: 120,
      ringNodeCounts: { backend: 40, api: 20, frontend: 60 },
      ringCenters: {
        backend: { x: -100, y: 0 },
        api: { x: 0, y: 80 },
        frontend: { x: 100, y: 0 },
      },
      ringIslandCounts: { backend: 3, api: 2, frontend: 4 },
      apiBetweenBackendAndFrontend: true,
      docsCentered: true,
      sameColorIslandViolations: 0,
    });

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.layoutRings.nodeCount).toBe(120);
    expect(diagnostics?.layoutRings.ringCount).toBe(3);
    expect(diagnostics?.layoutRings.ringNodeCounts.api).toBe(20);
    expect(diagnostics?.layoutRings.ringIslandCounts.frontend).toBe(4);
    expect(diagnostics?.layoutRings.apiBetweenBackendAndFrontend).toBe(true);
    expect(diagnostics?.layoutRings.sameColorIslandViolations).toBe(0);
  });

  it('records node spacing layout diagnostics for e2e assertions', () => {
    recordLayoutNodeSpacing({
      nodeCount: 1800,
      islandCount: 1,
      renderedRadius: 3,
      renderedDiameter: 6,
      requiredEdgeGap: 6,
      requiredCenterDistance: 12,
      minObservedCenterDistance: 12,
      minObservedEdgeGap: 6,
      overlapCount: 0,
      targetGapViolationCount: 0,
    });

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.layoutNodeSpacing.nodeCount).toBe(1800);
    expect(diagnostics?.layoutNodeSpacing.islandCount).toBe(1);
    expect(diagnostics?.layoutNodeSpacing.requiredCenterDistance).toBe(12);
    expect(diagnostics?.layoutNodeSpacing.minObservedEdgeGap).toBe(6);
    expect(diagnostics?.layoutNodeSpacing.overlapCount).toBe(0);
    expect(diagnostics?.layoutNodeSpacing.targetGapViolationCount).toBe(0);
  });

  it('records screen-space node spacing diagnostics for e2e assertions', () => {
    recordScreenNodeSpacing({
      coordinateSpace: 'viewport_px',
      nodeCount: 1677,
      islandCount: 1,
      viewportWidth: 1280,
      viewportHeight: 800,
      cameraRatio: 1,
      minRenderedRadius: 0.5,
      maxRenderedRadius: 3,
      maxRenderedDiameter: 6,
      minObservedCenterDistance: 12,
      minObservedEdgeGap: 6,
      maxRequiredCenterDistance: 12,
      overlapCount: 0,
      targetGapViolationCount: 0,
    });

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.screenNodeSpacing.coordinateSpace).toBe('viewport_px');
    expect(diagnostics?.screenNodeSpacing.nodeCount).toBe(1677);
    expect(diagnostics?.screenNodeSpacing.viewportWidth).toBe(1280);
    expect(diagnostics?.screenNodeSpacing.maxRenderedDiameter).toBe(6);
    expect(diagnostics?.screenNodeSpacing.minObservedEdgeGap).toBe(6);
    expect(diagnostics?.screenNodeSpacing.overlapCount).toBe(0);
    expect(diagnostics?.screenNodeSpacing.targetGapViolationCount).toBe(0);
  });

  it('counts reconnect banner transitions without double-counting the same state', () => {
    recordReconnectBannerState(false);
    recordReconnectBannerState(true);
    recordReconnectBannerState(true);
    recordReconnectBannerState(false);

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.reconnectBanner.visible).toBe(false);
    expect(diagnostics?.reconnectBanner.shows).toBe(1);
    expect(diagnostics?.reconnectBanner.hides).toBe(1);
  });
});
