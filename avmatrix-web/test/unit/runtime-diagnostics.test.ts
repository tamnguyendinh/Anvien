import { beforeEach, describe, expect, it } from 'vitest';
import {
  getWebRuntimeDiagnostics,
  recordGraphConversion,
  recordGraphInteractionMode,
  recordGraphOverview,
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
      visibleViewportNodeCount: 1677,
      visibleViewportIslandCounts: { 'frontend:Function': 1677 },
      cameraRatio: 1,
      cameraX: 0,
      cameraY: 0,
      viewportGraphCenterX: 0,
      viewportGraphCenterY: 0,
      viewportGraphMinX: -640,
      viewportGraphMaxX: 640,
      viewportGraphMinY: -400,
      viewportGraphMaxY: 400,
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
    expect(diagnostics?.screenNodeSpacing.visibleViewportNodeCount).toBe(1677);
    expect(diagnostics?.screenNodeSpacing.viewportWidth).toBe(1280);
    expect(diagnostics?.screenNodeSpacing.maxRenderedDiameter).toBe(6);
    expect(diagnostics?.screenNodeSpacing.minObservedEdgeGap).toBe(6);
    expect(diagnostics?.screenNodeSpacing.overlapCount).toBe(0);
    expect(diagnostics?.screenNodeSpacing.targetGapViolationCount).toBe(0);
    expect(diagnostics?.graphInteraction.overviewSamples).toHaveLength(1);
    expect(
      diagnostics?.graphInteraction.overviewSamples[0].visibleViewportIslandCount,
    ).toBe(1);
  });

  it('records bounded graph interaction samples by current camera mode', () => {
    const input = {
      coordinateSpace: 'viewport_px' as const,
      nodeCount: 20,
      islandCount: 3,
      viewportWidth: 1280,
      viewportHeight: 800,
      visibleViewportNodeCount: 15,
      visibleViewportIslandCounts: {
        'backend:Function': 5,
        'frontend:Function': 8,
        'docs:Documentation': 2,
      },
      cameraRatio: 1,
      cameraX: 0,
      cameraY: 0,
      viewportGraphCenterX: 0,
      viewportGraphCenterY: 0,
      viewportGraphMinX: -100,
      viewportGraphMaxX: 100,
      viewportGraphMinY: -100,
      viewportGraphMaxY: 100,
      minRenderedRadius: 1,
      maxRenderedRadius: 3,
      maxRenderedDiameter: 6,
      minObservedCenterDistance: 12,
      minObservedEdgeGap: 6,
      maxRequiredCenterDistance: 12,
      overlapCount: 0,
      targetGapViolationCount: 0,
    };

    recordGraphInteractionMode({ mode: 'zoom-in' });
    recordScreenNodeSpacing({ ...input, cameraRatio: 0.75 });
    recordGraphInteractionMode({
      mode: 'detail-focus',
      targetNodeId: 'Function:frontend:0',
    });
    for (let index = 0; index < 13; index++) {
      recordScreenNodeSpacing({
        ...input,
        cameraRatio: 0.5 - index * 0.01,
      });
    }

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.graphInteraction.zoomSamples).toHaveLength(1);
    expect(diagnostics?.graphInteraction.zoomSamples[0].mode).toBe('zoom-in');
    expect(diagnostics?.graphInteraction.detailFocusSamples).toHaveLength(12);
    expect(
      diagnostics?.graphInteraction.detailFocusSamples.at(-1)?.targetNodeId,
    ).toBe('Function:frontend:0');
    expect(
      diagnostics?.graphInteraction.detailFocusSamples.at(-1)
        ?.visibleViewportIslandCount,
    ).toBe(3);
    expect(diagnostics?.graphInteraction.dynamicGapSamples).toHaveLength(12);
  });

  it('records graph overview diagnostics for e2e assertions', () => {
    recordGraphOverview({
      nodeCount: 100,
      viewportWidth: 1280,
      viewportHeight: 800,
      visibleViewportNodeCount: 40,
      visibleColorCount: 4,
      visibleRingCount: 3,
      visibleIslandCount: 3,
      dominantIslandKey: 'frontend:Function',
      dominantIslandShare: 0.5,
      visibleColorCounts: {
        '#22c55e': 20,
        '#3b82f6': 10,
        '#f59e0b': 5,
        '#84cc16': 5,
      },
      visibleRingCounts: {
        frontend: 20,
        backend: 10,
        docs: 10,
      },
      visibleIslandCounts: {
        'frontend:Function': 20,
        'backend:Function': 10,
        'docs:Documentation': 10,
      },
      visibleNodeTypeCounts: { Function: 30, Documentation: 10 },
      graphRingCounts: {
        frontend: 80,
        backend: 10,
        docs: 10,
      },
      graphIslandCounts: {
        'frontend:Function': 60,
        'frontend:Method': 20,
        'backend:Function': 10,
        'docs:Documentation': 10,
      },
      graphNodeTypeCounts: {
        Function: 60,
        Method: 20,
        Route: 10,
        Documentation: 10,
      },
      visibleRingInventory: ['backend', 'docs', 'frontend'],
      visibleNodeTypeInventory: ['Documentation', 'Function'],
      graphRingInventory: ['backend', 'docs', 'frontend'],
      graphIslandInventory: [
        'backend:Function',
        'docs:Documentation',
        'frontend:Function',
        'frontend:Method',
      ],
      filterNodeTypeInventory: ['Documentation', 'Function', 'Method', 'Route'],
      cameraRatio: 1,
      cameraX: 0.5,
      cameraY: 0.5,
    });

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.graphOverview.visibleColorCount).toBe(4);
    expect(diagnostics?.graphOverview.visibleRingCount).toBe(3);
    expect(diagnostics?.graphOverview.visibleIslandCount).toBe(3);
    expect(diagnostics?.graphOverview.dominantIslandShare).toBe(0.5);
    expect(diagnostics?.graphOverview.graphRingInventory).toContain('frontend');
    expect(diagnostics?.graphOverview.graphIslandInventory).toContain(
      'frontend:Method',
    );
    expect(diagnostics?.graphOverview.filterNodeTypeInventory).toContain('Route');
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
