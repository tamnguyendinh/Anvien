import type Graph from "graphology";
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from "./graph-adapter";

export type GraphOrientationLabelKind = "ring" | "island";

export type GraphOrientationBounds = {
  minX: number;
  maxX: number;
  minY: number;
  maxY: number;
};

export type GraphOrientationLabel = {
  id: string;
  kind: GraphOrientationLabelKind;
  displayText: string;
  fallbackText: string;
  sourceKey: string;
  ringKey: string;
  islandKey?: string;
  anchorX: number;
  anchorY: number;
  visibleNodeCount: number;
  bounds: GraphOrientationBounds;
};

export type GraphOrientationViewportLabel = GraphOrientationLabel & {
  viewportX: number;
  viewportY: number;
  width: number;
  height: number;
  compact: boolean;
};

type LabelBucket = {
  bounds: GraphOrientationBounds;
  count: number;
  ringCenterXTotal: number;
  ringCenterYTotal: number;
  ringCenterSamples: number;
};

type ViewportLabelCandidate = GraphOrientationLabel & {
  viewportX: number;
  viewportY: number;
  width: number;
  height: number;
  compact: boolean;
};

type ViewportBox = {
  left: number;
  right: number;
  top: number;
  bottom: number;
};

export const ORIENTATION_FAR_ZOOM_RATIO = 6;
export const ORIENTATION_FAR_ZOOM_MIN_ISLAND_NODES = 3;

const MISSING_RING_KEY = "missing_app_layer";
const UNKNOWN_ISLAND_KEY = "unknown";

const APP_LAYER_RING_LABELS: Record<string, string> = {
  frontend: "Frontend",
  frontend_test: "Frontend Tests",
  generated: "Generated",
  generated_contract: "Generated Contracts",
  config: "Config",
  [MISSING_RING_KEY]: "Unclassified",
  unknown: "Unknown",
  mixed: "Mixed",
  backend: "Backend",
  backend_test: "Backend Tests",
  shared_contract: "Shared Contracts",
  api_contract: "API Contracts",
  api_shared_contract: "API Shared Contracts",
  api: "API",
  frontend_api_client: "Frontend API Client",
  api_test: "API Tests",
  cli_launcher: "CLI Launcher",
  docs: "Documentation",
};

const createEmptyBounds = (): GraphOrientationBounds => ({
  minX: Number.POSITIVE_INFINITY,
  maxX: Number.NEGATIVE_INFINITY,
  minY: Number.POSITIVE_INFINITY,
  maxY: Number.NEGATIVE_INFINITY,
});

const updateBounds = (
  bounds: GraphOrientationBounds,
  x: number,
  y: number,
): void => {
  bounds.minX = Math.min(bounds.minX, x);
  bounds.maxX = Math.max(bounds.maxX, x);
  bounds.minY = Math.min(bounds.minY, y);
  bounds.maxY = Math.max(bounds.maxY, y);
};

const createBucket = (): LabelBucket => ({
  bounds: createEmptyBounds(),
  count: 0,
  ringCenterXTotal: 0,
  ringCenterYTotal: 0,
  ringCenterSamples: 0,
});

const normalizeKey = (value: string | undefined, fallback: string): string => {
  const trimmed = typeof value === "string" ? value.trim() : "";
  return trimmed.length > 0 ? trimmed : fallback;
};

const humanizeKey = (key: string): string => {
  const normalized = key
    .replace(/[_:./-]+/g, " ")
    .replace(/([a-z0-9])([A-Z])/g, "$1 $2")
    .trim();
  if (!normalized) return "Unknown";

  return normalized
    .split(/\s+/)
    .map((part) => {
      const lower = part.toLowerCase();
      if (lower === "api") return "API";
      if (lower === "cli") return "CLI";
      if (lower === "ui") return "UI";
      return lower.charAt(0).toUpperCase() + lower.slice(1);
    })
    .join(" ");
};

export const formatRingLabel = (ringKey: string): string =>
  APP_LAYER_RING_LABELS[ringKey] ?? humanizeKey(ringKey);

export const formatIslandLabel = (islandKey: string): string => {
  if (!islandKey.startsWith("ResolutionGap:")) {
    return humanizeKey(islandKey);
  }

  const detail = islandKey.slice("ResolutionGap:".length);
  return detail ? `ResolutionGap / ${humanizeKey(detail)}` : "ResolutionGap";
};

const createRingLabel = (
  ringKey: string,
  bucket: LabelBucket,
): GraphOrientationLabel | null => {
  if (bucket.count <= 0) return null;
  const hasRingCenter = bucket.ringCenterSamples > 0;
  const centerX = hasRingCenter
    ? bucket.ringCenterXTotal / bucket.ringCenterSamples
    : (bucket.bounds.minX + bucket.bounds.maxX) / 2;
  const centerY = hasRingCenter
    ? bucket.ringCenterYTotal / bucket.ringCenterSamples
    : (bucket.bounds.minY + bucket.bounds.maxY) / 2;
  const height = Math.max(1, bucket.bounds.maxY - bucket.bounds.minY);
  const labelGap = Math.max(42, Math.min(160, height * 0.08));

  return {
    id: `ring:${ringKey}`,
    kind: "ring",
    displayText: formatRingLabel(ringKey),
    fallbackText: APP_LAYER_RING_LABELS[MISSING_RING_KEY],
    sourceKey: ringKey,
    ringKey,
    anchorX: centerX,
    anchorY: Math.min(centerY, bucket.bounds.minY - labelGap),
    visibleNodeCount: bucket.count,
    bounds: { ...bucket.bounds },
  };
};

const createIslandLabel = (
  ringKey: string,
  islandKey: string,
  bucket: LabelBucket,
): GraphOrientationLabel | null => {
  if (bucket.count <= 0) return null;
  const centerX = (bucket.bounds.minX + bucket.bounds.maxX) / 2;
  const width = Math.max(1, bucket.bounds.maxX - bucket.bounds.minX);
  const height = Math.max(1, bucket.bounds.maxY - bucket.bounds.minY);
  const radius = Math.hypot(width, height) / 2;
  const labelGap = Math.max(28, Math.min(96, radius * 0.18));
  const sourceKey = `${ringKey}:${islandKey}`;

  return {
    id: `island:${sourceKey}`,
    kind: "island",
    displayText: formatIslandLabel(islandKey),
    fallbackText: humanizeKey(UNKNOWN_ISLAND_KEY),
    sourceKey,
    ringKey,
    islandKey,
    anchorX: centerX,
    anchorY: bucket.bounds.minY - labelGap,
    visibleNodeCount: bucket.count,
    bounds: { ...bucket.bounds },
  };
};

export const buildGraphOrientationLabels = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
): GraphOrientationLabel[] => {
  const rings = new Map<string, LabelBucket>();
  const islands = new Map<string, LabelBucket>();

  graph.forEachNode((_nodeId, attributes) => {
    if (attributes.hidden) return;
    const x = Number(attributes.x);
    const y = Number(attributes.y);
    if (!Number.isFinite(x) || !Number.isFinite(y)) return;

    const ringKey = normalizeKey(attributes.appLayerRing, MISSING_RING_KEY);
    const islandKey = normalizeKey(
      attributes.islandKey ?? attributes.nodeType,
      UNKNOWN_ISLAND_KEY,
    );

    const ringBucket = rings.get(ringKey) ?? createBucket();
    updateBounds(ringBucket.bounds, x, y);
    ringBucket.count += 1;
    if (
      typeof attributes.appLayerRingCenterX === "number" &&
      typeof attributes.appLayerRingCenterY === "number" &&
      Number.isFinite(attributes.appLayerRingCenterX) &&
      Number.isFinite(attributes.appLayerRingCenterY)
    ) {
      ringBucket.ringCenterXTotal += attributes.appLayerRingCenterX;
      ringBucket.ringCenterYTotal += attributes.appLayerRingCenterY;
      ringBucket.ringCenterSamples += 1;
    }
    rings.set(ringKey, ringBucket);

    const islandSourceKey = `${ringKey}:${islandKey}`;
    const islandBucket = islands.get(islandSourceKey) ?? createBucket();
    updateBounds(islandBucket.bounds, x, y);
    islandBucket.count += 1;
    islands.set(islandSourceKey, islandBucket);
  });

  const ringLabels = [...rings.entries()]
    .map(([ringKey, bucket]) => createRingLabel(ringKey, bucket))
    .filter((label): label is GraphOrientationLabel => label !== null)
    .sort((left, right) => left.sourceKey.localeCompare(right.sourceKey));

  const islandLabels = [...islands.entries()]
    .map(([sourceKey, bucket]) => {
      const separatorIndex = sourceKey.indexOf(":");
      const ringKey = sourceKey.slice(0, separatorIndex);
      const islandKey = sourceKey.slice(separatorIndex + 1);
      return createIslandLabel(ringKey, islandKey, bucket);
    })
    .filter((label): label is GraphOrientationLabel => label !== null)
    .sort(
      (left, right) =>
        left.ringKey.localeCompare(right.ringKey) ||
        (left.islandKey ?? "").localeCompare(right.islandKey ?? ""),
    );

  return [...ringLabels, ...islandLabels];
};

export const getOrientationLabelPresentation = (
  label: Pick<GraphOrientationLabel, "kind" | "visibleNodeCount">,
  cameraRatio: number,
): { visible: boolean; compact: boolean } => {
  const farZoom =
    Number.isFinite(cameraRatio) && cameraRatio > ORIENTATION_FAR_ZOOM_RATIO;
  if (!farZoom) {
    return { visible: true, compact: false };
  }

  if (label.kind === "ring") {
    return { visible: true, compact: false };
  }

  return {
    visible: label.visibleNodeCount >= ORIENTATION_FAR_ZOOM_MIN_ISLAND_NODES,
    compact: true,
  };
};

const estimateLabelSize = (
  label: Pick<GraphOrientationLabel, "displayText" | "kind" | "visibleNodeCount">,
  compact: boolean,
): { width: number; height: number } => {
  const textUnits = Math.max(6, label.displayText.length);
  const countUnits = compact ? 0 : String(label.visibleNodeCount).length + 2;
  const width = Math.min(
    label.kind === "ring" ? 220 : 190,
    textUnits * 7 + countUnits * 6 + (label.kind === "ring" ? 38 : 32),
  );
  return {
    width,
    height: label.kind === "ring" ? 28 : 24,
  };
};

const createBox = (
  label: Pick<GraphOrientationViewportLabel, "viewportX" | "viewportY" | "width" | "height">,
): ViewportBox => ({
  left: label.viewportX - label.width / 2,
  right: label.viewportX + label.width / 2,
  top: label.viewportY - label.height / 2,
  bottom: label.viewportY + label.height / 2,
});

const clampCoordinate = (value: number, minimum: number, maximum: number): number => {
  if (minimum > maximum) return (minimum + maximum) / 2;
  return Math.max(minimum, Math.min(maximum, value));
};

const clampCandidateToViewport = (
  candidate: ViewportLabelCandidate,
  viewportWidth: number,
  viewportHeight: number,
  safeInset: number,
): ViewportLabelCandidate => ({
  ...candidate,
  viewportX: clampCoordinate(
    candidate.viewportX,
    safeInset + candidate.width / 2,
    viewportWidth - safeInset - candidate.width / 2,
  ),
  viewportY: clampCoordinate(
    candidate.viewportY,
    safeInset + candidate.height / 2,
    viewportHeight - safeInset - candidate.height / 2,
  ),
});

const createOffsetPlacements = (
  candidate: ViewportLabelCandidate,
  viewportWidth: number,
  viewportHeight: number,
  safeInset: number,
): ViewportLabelCandidate[] => {
  const stepY = candidate.height + 8;
  const stepX = Math.min(120, Math.max(72, candidate.width * 0.6));
  const offsets = [
    { x: 0, y: 0 },
    { x: 0, y: -stepY },
    { x: 0, y: stepY },
    { x: -stepX, y: 0 },
    { x: stepX, y: 0 },
    { x: -stepX, y: -stepY },
    { x: stepX, y: -stepY },
    { x: -stepX, y: stepY },
    { x: stepX, y: stepY },
    { x: 0, y: -stepY * 2 },
    { x: 0, y: stepY * 2 },
  ];

  return offsets.map((offset) =>
    clampCandidateToViewport(
      {
        ...candidate,
        viewportX: candidate.viewportX + offset.x,
        viewportY: candidate.viewportY + offset.y,
      },
      viewportWidth,
      viewportHeight,
      safeInset,
    ),
  );
};

const boxesOverlap = (
  left: ViewportBox,
  right: ViewportBox,
  padding: number,
): boolean =>
  left.left < right.right + padding &&
  left.right + padding > right.left &&
  left.top < right.bottom + padding &&
  left.bottom + padding > right.top;

const isInsideViewport = (
  box: ViewportBox,
  viewportWidth: number,
  viewportHeight: number,
  safeInset: number,
): boolean =>
  box.left >= safeInset &&
  box.right <= viewportWidth - safeInset &&
  box.top >= safeInset &&
  box.bottom <= viewportHeight - safeInset;

export const placeGraphOrientationLabels = (
  labels: GraphOrientationLabel[],
  options: {
    viewportWidth: number;
    viewportHeight: number;
    cameraRatio: number;
    project: (point: { x: number; y: number }) => { x: number; y: number };
    safeInset?: number;
    collisionPadding?: number;
  },
): GraphOrientationViewportLabel[] => {
  const safeInset = options.safeInset ?? 12;
  const collisionPadding = options.collisionPadding ?? 6;
  if (options.viewportWidth <= 0 || options.viewportHeight <= 0) return [];

  const candidates = labels
    .map((label): ViewportLabelCandidate | null => {
      const presentation = getOrientationLabelPresentation(
        label,
        options.cameraRatio,
      );
      if (!presentation.visible) return null;
      const projected = options.project({
        x: label.anchorX,
        y: label.anchorY,
      });
      if (!Number.isFinite(projected.x) || !Number.isFinite(projected.y)) {
        return null;
      }
      const size = estimateLabelSize(label, presentation.compact);
      return {
        ...label,
        viewportX: projected.x,
        viewportY: projected.y,
        width: size.width,
        height: size.height,
        compact: presentation.compact,
      };
    })
    .filter((label): label is ViewportLabelCandidate => label !== null)
    .sort((left, right) => {
      if (left.kind !== right.kind) return left.kind === "ring" ? -1 : 1;
      if (left.kind === "island") {
        return (
          right.visibleNodeCount - left.visibleNodeCount ||
          left.sourceKey.localeCompare(right.sourceKey)
        );
      }
      return left.sourceKey.localeCompare(right.sourceKey);
    });

  const placed: GraphOrientationViewportLabel[] = [];
  const placedBoxes: ViewportBox[] = [];
  for (const candidate of candidates) {
    const placements = createOffsetPlacements(
      candidate,
      options.viewportWidth,
      options.viewportHeight,
      safeInset,
    );

    let selectedPlacement: ViewportLabelCandidate | null = null;
    let selectedBox: ViewportBox | null = null;
    for (const placement of placements) {
      const box = createBox(placement);
      if (
        !isInsideViewport(
          box,
          options.viewportWidth,
          options.viewportHeight,
          safeInset,
        )
      ) {
        continue;
      }
      if (
        placedBoxes.some((placedBox) =>
          boxesOverlap(box, placedBox, collisionPadding),
        )
      ) {
        continue;
      }
      selectedPlacement = placement;
      selectedBox = box;
      break;
    }
    if (!selectedPlacement && candidate.kind === "ring") {
      const fallbackPlacement = clampCandidateToViewport(
        candidate,
        options.viewportWidth,
        options.viewportHeight,
        safeInset,
      );
      const fallbackBox = createBox(fallbackPlacement);
      if (
        !placedBoxes.some((placedBox) =>
          boxesOverlap(fallbackBox, placedBox, collisionPadding),
        )
      ) {
        selectedPlacement = fallbackPlacement;
        selectedBox = fallbackBox;
      }
    }
    if (!selectedPlacement || !selectedBox) {
      continue;
    }
    placed.push(selectedPlacement);
    placedBoxes.push(selectedBox);
  }

  return placed.sort((left, right) => {
    if (left.kind !== right.kind) return left.kind === "ring" ? -1 : 1;
    return left.sourceKey.localeCompare(right.sourceKey);
  });
};
