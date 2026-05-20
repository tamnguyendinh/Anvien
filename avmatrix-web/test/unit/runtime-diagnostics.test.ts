import { beforeEach, describe, expect, it } from 'vitest';
import {
  getWebRuntimeDiagnostics,
  recordGraphConversion,
  recordManualLayoutOptimizerInvocation,
  recordReconnectBannerState,
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
