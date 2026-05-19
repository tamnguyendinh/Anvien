import { beforeEach, describe, expect, it } from 'vitest';
import {
  getWebRuntimeDiagnostics,
  recordGraphConversion,
  recordLayoutStart,
  recordLayoutStop,
  recordReconnectBannerState,
  recordVisualScale,
  resetWebRuntimeDiagnostics,
} from '../../src/lib/runtime-diagnostics';

describe('runtime diagnostics', () => {
  beforeEach(() => {
    resetWebRuntimeDiagnostics();
  });

  it('records graph conversion and layout timings', () => {
    recordGraphConversion({
      startedAt: 10,
      finishedAt: 35,
      nodeCount: 100,
      relationshipCount: 250,
    });
    recordLayoutStart({
      nodeCount: 100,
      durationBudgetMs: 20_000,
      startedAt: 40,
    });
    recordLayoutStop({
      nodeCount: 100,
      reason: 'duration-elapsed',
      runMs: 20_000,
      noverlapMs: 12,
      stoppedAt: 20_040,
    });

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.graphConversion.count).toBe(1);
    expect(diagnostics?.graphConversion.lastMs).toBe(25);
    expect(diagnostics?.graphConversion.lastNodeCount).toBe(100);
    expect(diagnostics?.graphConversion.lastRelationshipCount).toBe(250);
    expect(diagnostics?.layout.starts).toBe(1);
    expect(diagnostics?.layout.stops).toBe(1);
    expect(diagnostics?.layout.isRunning).toBe(false);
    expect(diagnostics?.layout.lastReason).toBe('duration-elapsed');
    expect(diagnostics?.layout.lastNoverlapMs).toBe(12);
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
