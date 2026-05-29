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
  graphInteraction: {
    currentMode: GraphInteractionMode;
    currentTargetNodeId: string;
    lastModeChangedAt: number;
    overviewSamples: GraphInteractionSample[];
    zoomSamples: GraphInteractionSample[];
    detailFocusSamples: GraphInteractionSample[];
    dynamicGapSamples: GraphInteractionSample[];
  };
  graphCamera: {
    recordedAt: number;
    mode: GraphInteractionMode;
    targetNodeId: string;
    cameraRatio: number;
    cameraX: number;
    cameraY: number;
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

export type GraphInteractionMode =
  | 'overview'
  | 'zoom-in'
  | 'zoom-out'
  | 'wheel-zoom'
  | 'detail-focus';

export interface GraphInteractionSample {
  recordedAt: number;
  mode: GraphInteractionMode;
  targetNodeId: string;
  coordinateSpace: 'viewport_px';
  nodeCount: number;
  islandCount: number;
  viewportWidth: number;
  viewportHeight: number;
  visibleViewportNodeCount: number;
  visibleViewportIslandCount: number;
  cameraRatio: number;
  cameraX: number;
  cameraY: number;
  minRenderedRadius: number;
  maxRenderedRadius: number;
  maxRenderedDiameter: number;
  minObservedCenterDistance: number;
  minObservedEdgeGap: number;
  maxRequiredCenterDistance: number;
  overlapCount: number;
  targetGapViolationCount: number;
}

export type GraphOverviewDiagnosticsInput = Omit<
  GraphOverviewDiagnostics,
  'recordedAt'
>;

declare global {
  interface Window {
    __ANVIEN_WEB_DIAGNOSTICS__?: WebRuntimeDiagnostics;
    __ANVIEN_RESET_WEB_DIAGNOSTICS__?: () => WebRuntimeDiagnostics;
  }
}

const nowMs = (): number =>
  typeof performance !== 'undefined' ? performance.now() : Date.now();

const GRAPH_INTERACTION_SAMPLE_LIMIT = 12;

const appendBoundedSample = (
  samples: GraphInteractionSample[],
  sample: GraphInteractionSample,
): void => {
  samples.push(sample);
  if (samples.length > GRAPH_INTERACTION_SAMPLE_LIMIT) {
    samples.splice(0, samples.length - GRAPH_INTERACTION_SAMPLE_LIMIT);
  }
};

const buildGraphInteractionSample = (
  input: Omit<WebRuntimeDiagnostics['screenNodeSpacing'], 'recordedAt'>,
  recordedAt: number,
  mode: GraphInteractionMode,
  targetNodeId: string,
): GraphInteractionSample => {
  const visibleViewportIslandCounts = input.visibleViewportIslandCounts ?? {};

  return {
    recordedAt,
    mode,
    targetNodeId,
    coordinateSpace: input.coordinateSpace,
    nodeCount: input.nodeCount,
    islandCount: input.islandCount,
    viewportWidth: input.viewportWidth,
    viewportHeight: input.viewportHeight,
    visibleViewportNodeCount: input.visibleViewportNodeCount ?? 0,
    visibleViewportIslandCount: Object.keys(visibleViewportIslandCounts).length,
    cameraRatio: input.cameraRatio,
    cameraX: input.cameraX,
    cameraY: input.cameraY,
    minRenderedRadius: input.minRenderedRadius,
    maxRenderedRadius: input.maxRenderedRadius,
    maxRenderedDiameter: input.maxRenderedDiameter,
    minObservedCenterDistance: input.minObservedCenterDistance,
    minObservedEdgeGap: input.minObservedEdgeGap,
    maxRequiredCenterDistance: input.maxRequiredCenterDistance,
    overlapCount: input.overlapCount,
    targetGapViolationCount: input.targetGapViolationCount,
  };
};

const recordGraphInteractionSample = (
  diagnostics: WebRuntimeDiagnostics,
  input: Omit<WebRuntimeDiagnostics['screenNodeSpacing'], 'recordedAt'>,
  recordedAt: number,
): void => {
  const { currentMode, currentTargetNodeId } = diagnostics.graphInteraction;
  const sample = buildGraphInteractionSample(
    input,
    recordedAt,
    currentMode,
    currentTargetNodeId,
  );

  if (currentMode === 'overview') {
    appendBoundedSample(diagnostics.graphInteraction.overviewSamples, sample);
  }
  if (
    currentMode === 'zoom-in' ||
    currentMode === 'zoom-out' ||
    currentMode === 'wheel-zoom'
  ) {
    appendBoundedSample(diagnostics.graphInteraction.zoomSamples, sample);
  }
  if (currentMode === 'detail-focus') {
    appendBoundedSample(
      diagnostics.graphInteraction.detailFocusSamples,
      sample,
    );
  }
  appendBoundedSample(diagnostics.graphInteraction.dynamicGapSamples, sample);
};

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
  graphInteraction: {
    currentMode: 'overview',
    currentTargetNodeId: '',
    lastModeChangedAt: 0,
    overviewSamples: [],
    zoomSamples: [],
    detailFocusSamples: [],
    dynamicGapSamples: [],
  },
  graphCamera: {
    recordedAt: 0,
    mode: 'overview',
    targetNodeId: '',
    cameraRatio: 0,
    cameraX: 0,
    cameraY: 0,
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
  window.__ANVIEN_WEB_DIAGNOSTICS__ ??= createDiagnostics();
  return window.__ANVIEN_WEB_DIAGNOSTICS__;
};

export const resetWebRuntimeDiagnostics = (): WebRuntimeDiagnostics => {
  const diagnostics = createDiagnostics();
  if (typeof window !== 'undefined') {
    window.__ANVIEN_WEB_DIAGNOSTICS__ = diagnostics;
    window.__ANVIEN_RESET_WEB_DIAGNOSTICS__ = resetWebRuntimeDiagnostics;
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
  const recordedAt = nowMs();
  diagnostics.screenNodeSpacing = {
    ...input,
    recordedAt,
  };
  recordGraphInteractionSample(diagnostics, input, recordedAt);
};

export const recordGraphInteractionMode = (input: {
  mode: GraphInteractionMode;
  targetNodeId?: string;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.graphInteraction.currentMode = input.mode;
  diagnostics.graphInteraction.currentTargetNodeId = input.targetNodeId ?? '';
  diagnostics.graphInteraction.lastModeChangedAt = nowMs();
};

export const recordGraphCameraSample = (input: {
  cameraRatio: number;
  cameraX: number;
  cameraY: number;
}): void => {
  const diagnostics = getWebRuntimeDiagnostics();
  if (!diagnostics) return;
  diagnostics.graphCamera = {
    ...input,
    mode: diagnostics.graphInteraction.currentMode,
    targetNodeId: diagnostics.graphInteraction.currentTargetNodeId,
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
  window.__ANVIEN_RESET_WEB_DIAGNOSTICS__ = resetWebRuntimeDiagnostics;
  window.__ANVIEN_WEB_DIAGNOSTICS__ ??= createDiagnostics();
}
