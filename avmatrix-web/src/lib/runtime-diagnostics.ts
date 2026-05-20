export interface WebRuntimeDiagnostics {
  graphConversion: {
    count: number;
    lastMs: number;
    maxMs: number;
    lastNodeCount: number;
    lastRelationshipCount: number;
    lastStartedAt: number;
    lastFinishedAt: number;
  };
  visualScale: {
    nodeCount: number;
    minNodeSize: number;
    maxNodeSize: number;
    maxRenderedNodeSizeCap: number;
    structuralToLeafRatio: number;
    maxSizeByLabel: Record<string, number>;
    recordedAt: number;
  };
  layout: {
    starts: number;
    stops: number;
    manualOptimizerInvocations: number;
    isRunning: boolean;
    lastNodeCount: number;
    lastManualOptimizerInvokedAt: number;
    lastManualOptimizerRunMs: number;
  };
  heartbeat: {
    connects: number;
    reconnects: number;
    lastConnectAt: number;
    lastReconnectAt: number;
    lastRetryAttempt: number;
  };
  reconnectBanner: {
    visible: boolean;
    shows: number;
    hides: number;
    lastChangedAt: number;
  };
}

declare global {
  interface Window {
    __AVMATRIX_WEB_DIAGNOSTICS__?: WebRuntimeDiagnostics;
    __AVMATRIX_RESET_WEB_DIAGNOSTICS__?: () => WebRuntimeDiagnostics;
  }
}

const nowMs = (): number =>
  typeof performance !== 'undefined' ? performance.now() : Date.now();

const createDiagnostics = (): WebRuntimeDiagnostics => ({
  graphConversion: {
    count: 0,
    lastMs: 0,
    maxMs: 0,
    lastNodeCount: 0,
    lastRelationshipCount: 0,
    lastStartedAt: 0,
    lastFinishedAt: 0,
  },
  visualScale: {
    nodeCount: 0,
    minNodeSize: 0,
    maxNodeSize: 0,
    maxRenderedNodeSizeCap: 0,
    structuralToLeafRatio: 0,
    maxSizeByLabel: {},
    recordedAt: 0,
  },
  layout: {
    starts: 0,
    stops: 0,
    manualOptimizerInvocations: 0,
    isRunning: false,
    lastNodeCount: 0,
    lastManualOptimizerInvokedAt: 0,
    lastManualOptimizerRunMs: 0,
  },
  heartbeat: {
    connects: 0,
    reconnects: 0,
    lastConnectAt: 0,
    lastReconnectAt: 0,
    lastRetryAttempt: 0,
  },
  reconnectBanner: {
    visible: false,
    shows: 0,
    hides: 0,
    lastChangedAt: 0,
  },
});

export const getWebRuntimeDiagnostics = (): WebRuntimeDiagnostics | null => {
  if (typeof window === 'undefined') return null;
  window.__AVMATRIX_WEB_DIAGNOSTICS__ ??= createDiagnostics();
  return window.__AVMATRIX_WEB_DIAGNOSTICS__;
};

export const resetWebRuntimeDiagnostics = (): WebRuntimeDiagnostics => {
  const diagnostics = createDiagnostics();
  if (typeof window !== 'undefined') {
    window.__AVMATRIX_WEB_DIAGNOSTICS__ = diagnostics;
    window.__AVMATRIX_RESET_WEB_DIAGNOSTICS__ = resetWebRuntimeDiagnostics;
  }
  return diagnostics;
};

export const recordGraphConversion = (input: {
  startedAt: number;
  finishedAt?: number;
  nodeCount: number;
  relationshipCount: number;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  const finishedAt = input.finishedAt ?? nowMs();
  const elapsedMs = Math.max(0, finishedAt - input.startedAt);
  diagnostics.graphConversion.count++;
  diagnostics.graphConversion.lastMs = elapsedMs;
  diagnostics.graphConversion.maxMs = Math.max(
    diagnostics.graphConversion.maxMs,
    elapsedMs,
  );
  diagnostics.graphConversion.lastNodeCount = input.nodeCount;
  diagnostics.graphConversion.lastRelationshipCount = input.relationshipCount;
  diagnostics.graphConversion.lastStartedAt = input.startedAt;
  diagnostics.graphConversion.lastFinishedAt = finishedAt;
};

export const recordVisualScale = (input: {
  nodeCount: number;
  minNodeSize: number;
  maxNodeSize: number;
  maxRenderedNodeSizeCap: number;
  structuralToLeafRatio: number;
  maxSizeByLabel?: Record<string, number>;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.visualScale = {
    maxSizeByLabel: {},
    ...input,
    recordedAt: nowMs(),
  };
};

export const recordManualLayoutOptimizerInvocation = (input: {
  nodeCount: number;
  startedAt?: number;
  finishedAt?: number;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  const finishedAt = input.finishedAt ?? nowMs();
  const startedAt = input.startedAt ?? finishedAt;
  diagnostics.layout.manualOptimizerInvocations++;
  diagnostics.layout.lastNodeCount = input.nodeCount;
  diagnostics.layout.lastManualOptimizerInvokedAt = startedAt;
  diagnostics.layout.lastManualOptimizerRunMs = Math.max(0, finishedAt - startedAt);
};

export const recordHeartbeatConnect = (): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.heartbeat.connects++;
  diagnostics.heartbeat.lastConnectAt = nowMs();
};

export const recordHeartbeatReconnect = (attempt: number): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.heartbeat.reconnects++;
  diagnostics.heartbeat.lastReconnectAt = nowMs();
  diagnostics.heartbeat.lastRetryAttempt = attempt;
};

export const recordReconnectBannerState = (visible: boolean): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics || diagnostics.reconnectBanner.visible === visible) return;
  diagnostics.reconnectBanner.visible = visible;
  diagnostics.reconnectBanner.lastChangedAt = nowMs();
  if (visible) {
    diagnostics.reconnectBanner.shows++;
  } else {
    diagnostics.reconnectBanner.hides++;
  }
};

if (typeof window !== 'undefined') {
  window.__AVMATRIX_RESET_WEB_DIAGNOSTICS__ = resetWebRuntimeDiagnostics;
  window.__AVMATRIX_WEB_DIAGNOSTICS__ ??= createDiagnostics();
}
