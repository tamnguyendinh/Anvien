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
  layoutRings: {
    recordedAt: number;
    nodeCount: number;
    ringCount: number;
    ringNodeCounts: Record<string, number>;
    ringCenters: Record<string, { x: number; y: number }>;
    ringIslandCounts: Record<string, number>;
    apiBetweenBackendAndFrontend: boolean;
    docsCentered: boolean;
    sameColorIslandViolations: number;
  };
  layoutNodeSpacing: {
    recordedAt: number;
    nodeCount: number;
    islandCount: number;
    renderedRadius: number;
    renderedDiameter: number;
    requiredEdgeGap: number;
    requiredCenterDistance: number;
    minObservedCenterDistance: number;
    minObservedEdgeGap: number;
    overlapCount: number;
    targetGapViolationCount: number;
  };
  screenNodeSpacing: {
    recordedAt: number;
    coordinateSpace: 'viewport_px';
    nodeCount: number;
    islandCount: number;
    viewportWidth: number;
    viewportHeight: number;
    visibleViewportNodeCount: number;
    visibleViewportIslandCounts: Record<string, number>;
    cameraRatio: number;
    cameraX: number;
    cameraY: number;
    viewportGraphCenterX: number;
    viewportGraphCenterY: number;
    viewportGraphMinX: number;
    viewportGraphMaxX: number;
    viewportGraphMinY: number;
    viewportGraphMaxY: number;
    minRenderedRadius: number;
    maxRenderedRadius: number;
    maxRenderedDiameter: number;
    minObservedCenterDistance: number;
    minObservedEdgeGap: number;
    maxRequiredCenterDistance: number;
    overlapCount: number;
    targetGapViolationCount: number;
  };
  graphOverview: GraphOverviewDiagnostics;
  readableCamera: {
    recordedAt: number;
    applied: boolean;
    focusedIslandKey: string;
    focusedIslandNodeCount: number;
    focusCellNodeCount: number;
    focusRawX: number;
    focusRawY: number;
    cameraX: number;
    cameraY: number;
    ratio: number;
    previousMaxRenderedRadius: number;
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

export interface GraphOverviewDiagnostics {
  recordedAt: number;
  nodeCount: number;
  viewportWidth: number;
  viewportHeight: number;
  visibleViewportNodeCount: number;
  visibleColorCount: number;
  visibleRingCount: number;
  visibleIslandCount: number;
  dominantIslandKey: string;
  dominantIslandShare: number;
  visibleColorCounts: Record<string, number>;
  visibleRingCounts: Record<string, number>;
  visibleIslandCounts: Record<string, number>;
  visibleNodeTypeCounts: Record<string, number>;
  graphRingCounts: Record<string, number>;
  graphIslandCounts: Record<string, number>;
  graphNodeTypeCounts: Record<string, number>;
  visibleRingInventory: string[];
  visibleNodeTypeInventory: string[];
  graphRingInventory: string[];
  graphIslandInventory: string[];
  filterNodeTypeInventory: string[];
  cameraRatio: number;
  cameraX: number;
  cameraY: number;
}

export type GraphOverviewDiagnosticsInput = Omit<
  GraphOverviewDiagnostics,
  'recordedAt'
>;

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
  layoutRings: {
    recordedAt: 0,
    nodeCount: 0,
    ringCount: 0,
    ringNodeCounts: {},
    ringCenters: {},
    ringIslandCounts: {},
    apiBetweenBackendAndFrontend: false,
    docsCentered: false,
    sameColorIslandViolations: 0,
  },
  layoutNodeSpacing: {
    recordedAt: 0,
    nodeCount: 0,
    islandCount: 0,
    renderedRadius: 0,
    renderedDiameter: 0,
    requiredEdgeGap: 0,
    requiredCenterDistance: 0,
    minObservedCenterDistance: 0,
    minObservedEdgeGap: 0,
    overlapCount: 0,
    targetGapViolationCount: 0,
  },
  screenNodeSpacing: {
    recordedAt: 0,
    coordinateSpace: 'viewport_px',
    nodeCount: 0,
    islandCount: 0,
    viewportWidth: 0,
    viewportHeight: 0,
    visibleViewportNodeCount: 0,
    visibleViewportIslandCounts: {},
    cameraRatio: 0,
    cameraX: 0,
    cameraY: 0,
    viewportGraphCenterX: 0,
    viewportGraphCenterY: 0,
    viewportGraphMinX: 0,
    viewportGraphMaxX: 0,
    viewportGraphMinY: 0,
    viewportGraphMaxY: 0,
    minRenderedRadius: 0,
    maxRenderedRadius: 0,
    maxRenderedDiameter: 0,
    minObservedCenterDistance: 0,
    minObservedEdgeGap: 0,
    maxRequiredCenterDistance: 0,
    overlapCount: 0,
    targetGapViolationCount: 0,
  },
  graphOverview: {
    recordedAt: 0,
    nodeCount: 0,
    viewportWidth: 0,
    viewportHeight: 0,
    visibleViewportNodeCount: 0,
    visibleColorCount: 0,
    visibleRingCount: 0,
    visibleIslandCount: 0,
    dominantIslandKey: '',
    dominantIslandShare: 0,
    visibleColorCounts: {},
    visibleRingCounts: {},
    visibleIslandCounts: {},
    visibleNodeTypeCounts: {},
    graphRingCounts: {},
    graphIslandCounts: {},
    graphNodeTypeCounts: {},
    visibleRingInventory: [],
    visibleNodeTypeInventory: [],
    graphRingInventory: [],
    graphIslandInventory: [],
    filterNodeTypeInventory: [],
    cameraRatio: 0,
    cameraX: 0,
    cameraY: 0,
  },
  readableCamera: {
    recordedAt: 0,
    applied: false,
    focusedIslandKey: '',
    focusedIslandNodeCount: 0,
    focusCellNodeCount: 0,
    focusRawX: 0,
    focusRawY: 0,
    cameraX: 0,
    cameraY: 0,
    ratio: 0,
    previousMaxRenderedRadius: 0,
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

export const recordLayoutRings = (input: {
  nodeCount: number;
  ringNodeCounts: Record<string, number>;
  ringCenters: Record<string, { x: number; y: number }>;
  ringIslandCounts: Record<string, number>;
  apiBetweenBackendAndFrontend: boolean;
  docsCentered: boolean;
  sameColorIslandViolations: number;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.layoutRings = {
    ...input,
    ringCount: Object.keys(input.ringNodeCounts).length,
    recordedAt: nowMs(),
  };
};

export const recordLayoutNodeSpacing = (input: {
  nodeCount: number;
  islandCount: number;
  renderedRadius: number;
  renderedDiameter: number;
  requiredEdgeGap: number;
  requiredCenterDistance: number;
  minObservedCenterDistance: number;
  minObservedEdgeGap: number;
  overlapCount: number;
  targetGapViolationCount: number;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.layoutNodeSpacing = {
    ...input,
    recordedAt: nowMs(),
  };
};

export const recordScreenNodeSpacing = (input: {
  coordinateSpace: 'viewport_px';
  nodeCount: number;
  islandCount: number;
  viewportWidth: number;
  viewportHeight: number;
  visibleViewportNodeCount: number;
  visibleViewportIslandCounts: Record<string, number>;
  cameraRatio: number;
  cameraX: number;
  cameraY: number;
  viewportGraphCenterX: number;
  viewportGraphCenterY: number;
  viewportGraphMinX: number;
  viewportGraphMaxX: number;
  viewportGraphMinY: number;
  viewportGraphMaxY: number;
  minRenderedRadius: number;
  maxRenderedRadius: number;
  maxRenderedDiameter: number;
  minObservedCenterDistance: number;
  minObservedEdgeGap: number;
  maxRequiredCenterDistance: number;
  overlapCount: number;
  targetGapViolationCount: number;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.screenNodeSpacing = {
    ...input,
    recordedAt: nowMs(),
  };
};

export const recordGraphOverview = (
  input: GraphOverviewDiagnosticsInput,
): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.graphOverview = {
    ...input,
    recordedAt: nowMs(),
  };
};

export const recordReadableCamera = (input: {
  applied: boolean;
  focusedIslandKey: string;
  focusedIslandNodeCount: number;
  focusCellNodeCount: number;
  focusRawX: number;
  focusRawY: number;
  cameraX: number;
  cameraY: number;
  ratio: number;
  previousMaxRenderedRadius: number;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.readableCamera = {
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
