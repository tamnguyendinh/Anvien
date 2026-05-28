import Graph from "graphology";
import {
  DOCUMENTATION_NODE_LABEL,
  FILTERABLE_LABELS,
} from "./constants";
import {
  getMinimumNodeCenterDistance,
} from "./graph-node-sizing";
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from "./graph-adapter-types";

const filterableLabelOrder = new Map<string, number>(
  FILTERABLE_LABELS.map((label, index) => [label, index]),
);

const stableString = (value: unknown): string =>
  typeof value === "string" ? value : value == null ? "" : String(value);

const compareStableString = (left: string, right: string): number => {
  if (left === right) return 0;
  return left < right ? -1 : 1;
};

const compareClusterLabels = (left: string, right: string): number => {
  const leftOrder = filterableLabelOrder.get(left);
  const rightOrder = filterableLabelOrder.get(right);

  if (leftOrder !== undefined && rightOrder !== undefined) {
    return leftOrder - rightOrder;
  }
  if (leftOrder !== undefined) return -1;
  if (rightOrder !== undefined) return 1;
  return compareStableString(left, right);
};

const getClusterNodeSpacing = (nodeCount: number): number =>
  getMinimumNodeCenterDistance(nodeCount);

const MISSING_APP_LAYER_RING = "missing_app_layer";

export const APP_LAYER_RING_ORDER = [
  "frontend",
  "frontend_test",
  "generated",
  "generated_contract",
  "config",
  MISSING_APP_LAYER_RING,
  "unknown",
  "mixed",
  "backend",
  "backend_test",
  "shared_contract",
  "api_contract",
  "api_shared_contract",
  "api",
  "frontend_api_client",
  "api_test",
  "cli_launcher",
  "docs",
] as const;

const appLayerRingOrder = new Map<string, number>(
  APP_LAYER_RING_ORDER.map((layer, index) => [layer, index]),
);

const APP_LAYER_RING_ANGLES: Record<string, number> = {
  frontend: 0,
  frontend_test: Math.PI / 6,
  generated: Math.PI / 3,
  generated_contract: (Math.PI * 4) / 9,
  config: (Math.PI * 11) / 18,
  [MISSING_APP_LAYER_RING]: (Math.PI * 25) / 36,
  unknown: (Math.PI * 31) / 36,
  mixed: (Math.PI * 3) / 4,
  backend: Math.PI,
  backend_test: (Math.PI * 41) / 36,
  shared_contract: (Math.PI * 5) / 4,
  api_contract: (Math.PI * 49) / 36,
  api_shared_contract: (Math.PI * 13) / 9,
  api: (Math.PI * 3) / 2,
  frontend_api_client: (Math.PI * 5) / 3,
  api_test: (Math.PI * 16) / 9,
  cli_launcher: (Math.PI * 17) / 9,
};

const compareAppLayerRingKeys = (left: string, right: string): number => {
  const leftOrder = appLayerRingOrder.get(left);
  const rightOrder = appLayerRingOrder.get(right);
  if (leftOrder !== undefined && rightOrder !== undefined) {
    return leftOrder - rightOrder;
  }
  if (leftOrder !== undefined) return -1;
  if (rightOrder !== undefined) return 1;
  return compareStableString(left, right);
};

const getAppLayerRingKey = (attributes: SigmaNodeAttributes): string => {
  if (attributes.appLayer) return attributes.appLayer;
  if (attributes.nodeType === DOCUMENTATION_NODE_LABEL) return "docs";
  return MISSING_APP_LAYER_RING;
};

const getNodeIslandKey = (attributes: SigmaNodeAttributes): string => {
  if (attributes.nodeType !== "ResolutionGap") {
    return attributes.nodeType;
  }

  if (
    attributes.classification === "builtin" ||
    attributes.classification === "standard_library" ||
    attributes.classification === "test_framework"
  ) {
    return ["ResolutionGap", attributes.classification].join(":");
  }

  const gapKind = stableString(attributes.gapKind);
  const factFamily = stableString(attributes.factFamily);
  const targetRole = stableString(attributes.targetRole);
  return ["ResolutionGap", gapKind || factFamily || targetRole || "unknown"].join(
    ":",
  );
};

const getComparableIslandLabel = (islandKey: string): string =>
  islandKey.startsWith("ResolutionGap:") ? "ResolutionGap" : islandKey;

const compareIslandKeys = (left: string, right: string): number =>
  compareClusterLabels(getComparableIslandLabel(left), getComparableIslandLabel(right)) ||
  compareStableString(left, right);

const getAppLayerRingAngle = (layer: string): number => {
  const knownAngle = APP_LAYER_RING_ANGLES[layer];
  if (knownAngle !== undefined) return knownAngle;

  const seed = getStableLabelSeed(layer);
  return (seed / 0xffffffff) * Math.PI * 2;
};

const getMinimumAngularSeparation = (angles: number[]): number => {
  if (angles.length <= 1) return Math.PI * 2;
  const normalized = angles
    .map((angle) => ((angle % (Math.PI * 2)) + Math.PI * 2) % (Math.PI * 2))
    .sort((left, right) => left - right);
  let minimum = Math.PI * 2;
  normalized.forEach((angle, index) => {
    const next = normalized[(index + 1) % normalized.length];
    const gap =
      index === normalized.length - 1
        ? next + Math.PI * 2 - angle
        : next - angle;
    minimum = Math.min(minimum, gap);
  });
  return Math.max(minimum, Math.PI / 36);
};

const getStableLabelSeed = (label: string): number => {
  let hash = 2166136261;
  for (let index = 0; index < label.length; index++) {
    hash ^= label.charCodeAt(index);
    hash = Math.imul(hash, 16777619);
  }
  return hash >>> 0;
};

type IslandOffset = { x: number; y: number };

type IslandOffsetBounds = {
  minX: number;
  maxX: number;
  minY: number;
  maxY: number;
};

const getIslandOffsetBounds = (offsets: IslandOffset[]): IslandOffsetBounds =>
  offsets.reduce(
    (current, offset) => ({
      minX: Math.min(current.minX, offset.x),
      maxX: Math.max(current.maxX, offset.x),
      minY: Math.min(current.minY, offset.y),
      maxY: Math.max(current.maxY, offset.y),
    }),
    {
      minX: Number.POSITIVE_INFINITY,
      maxX: Number.NEGATIVE_INFINITY,
      minY: Number.POSITIVE_INFINITY,
      maxY: Number.NEGATIVE_INFINITY,
    },
  );

const getClusterIslandRadius = (
  nodeCount: number,
  nodeSpacing: number,
  offsets: IslandOffset[] = [],
  minimumNodeCenterDistance: number = getMinimumNodeCenterDistance(nodeCount),
): number => {
  const footprintMargin = Math.max(nodeSpacing, minimumNodeCenterDistance);
  if (offsets.length === 0) return footprintMargin;

  const bounds = getIslandOffsetBounds(offsets);
  const width = bounds.maxX - bounds.minX;
  const height = bounds.maxY - bounds.minY;

  return Math.max(
    Math.hypot(width, height) / 2 + footprintMargin,
    footprintMargin,
  );
};

const HEX_AXIAL_DIRECTIONS: Array<[number, number]> = [
  [1, 0],
  [1, -1],
  [0, -1],
  [-1, 0],
  [-1, 1],
  [0, 1],
];

const getHexAxialRingCoordinates = (
  ringIndex: number,
): Array<[number, number]> => {
  const coordinates: Array<[number, number]> = [];
  let q = -ringIndex;
  let r = ringIndex;

  for (const [dq, dr] of HEX_AXIAL_DIRECTIONS) {
    for (let step = 0; step < ringIndex; step++) {
      coordinates.push([q, r]);
      q += dq;
      r += dr;
    }
  }

  return coordinates;
};

const getHexIslandOffset = (
  axialQ: number,
  axialR: number,
  cellSpacing: number,
): IslandOffset => ({
  x: cellSpacing * (axialQ + axialR / 2),
  y: cellSpacing * (Math.sqrt(3) / 2) * axialR,
});

const getIslandOffsets = (
  nodeCount: number,
  nodeSpacing: number,
  _labelSeed: number,
  minimumNodeCenterDistance: number,
): IslandOffset[] => {
  if (nodeCount <= 0) return [];

  const cellSpacing =
    Math.max(nodeSpacing, minimumNodeCenterDistance) * 1.000001;
  const offsets: IslandOffset[] = [{ x: 0, y: 0 }];
  let ringIndex = 1;

  while (offsets.length < nodeCount) {
    for (const [q, r] of getHexAxialRingCoordinates(ringIndex)) {
      offsets.push(getHexIslandOffset(q, r, cellSpacing));
      if (offsets.length === nodeCount) break;
    }
    ringIndex++;
  }

  return offsets;
};

const getBalancedCircularSlots = (slotCount: number): number[] => {
  if (slotCount <= 0) return [];

  const slots = [0];
  while (slots.length < slotCount) {
    let bestSlot = -1;
    let bestDistance = -1;

    for (let candidate = 0; candidate < slotCount; candidate++) {
      if (slots.includes(candidate)) continue;

      const nearestDistance = slots.reduce((nearest, occupiedSlot) => {
        const clockwise = Math.abs(candidate - occupiedSlot);
        const circularDistance = Math.min(clockwise, slotCount - clockwise);
        return Math.min(nearest, circularDistance);
      }, Number.POSITIVE_INFINITY);

      if (nearestDistance > bestDistance) {
        bestDistance = nearestDistance;
        bestSlot = candidate;
      }
    }

    slots.push(bestSlot);
  }

  return slots;
};

const compareClusterNodeIds = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
  leftNodeId: string,
  rightNodeId: string,
): number => {
  const left = graph.getNodeAttributes(leftNodeId);
  const right = graph.getNodeAttributes(rightNodeId);

  return (
    compareStableString(stableString(left.filePath), stableString(right.filePath)) ||
    compareStableString(stableString(left.label), stableString(right.label)) ||
    compareStableString(leftNodeId, rightNodeId)
  );
};

const getFootprintGap = (
  leftRadius: number,
  rightRadius: number,
  minimumNodeCenterDistance: number,
): number =>
  Math.max(minimumNodeCenterDistance, (leftRadius + rightRadius) * 0.12);

export const applyFilterBasedClusteredLayout = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
): void => {
  const totalNodeCount = graph.order;
  if (totalNodeCount === 0) return;

  const nodeIdsByRingAndIsland = new Map<string, Map<string, string[]>>();
  graph.forEachNode((nodeId, attributes) => {
    const ringKey = getAppLayerRingKey(attributes);
    const islandKey = getNodeIslandKey(attributes);
    const nodeIdsByIsland =
      nodeIdsByRingAndIsland.get(ringKey) ?? new Map<string, string[]>();
    const nodeIds = nodeIdsByIsland.get(islandKey) ?? [];
    nodeIds.push(nodeId);
    nodeIdsByIsland.set(islandKey, nodeIds);
    nodeIdsByRingAndIsland.set(ringKey, nodeIdsByIsland);
    graph.setNodeAttribute(nodeId, "appLayerRing", ringKey);
    graph.setNodeAttribute(nodeId, "islandKey", islandKey);
  });

  const nodeSpacing = getClusterNodeSpacing(totalNodeCount);
  const minimumNodeCenterDistance =
    getMinimumNodeCenterDistance(totalNodeCount);
  const rings = [...nodeIdsByRingAndIsland.entries()]
    .sort(([left], [right]) => compareAppLayerRingKeys(left, right))
    .map(([ringKey, nodeIdsByIsland]) => {
      const islandKeys = [...nodeIdsByIsland.keys()].sort(compareIslandKeys);
      const clusters = islandKeys.map((islandKey) => {
        const nodeIds = [...(nodeIdsByIsland.get(islandKey) ?? [])].sort(
          (left, right) => compareClusterNodeIds(graph, left, right),
        );
        const labelSeed = getStableLabelSeed(`${ringKey}:${islandKey}`);
        const offsets = getIslandOffsets(
          nodeIds.length,
          nodeSpacing,
          labelSeed,
          minimumNodeCenterDistance,
        );

        return {
          label: islandKey,
          nodeIds,
          labelSeed,
          offsets,
          radius: getClusterIslandRadius(
            nodeIds.length,
            nodeSpacing,
            offsets,
            minimumNodeCenterDistance,
          ),
        };
      });
      const largestClusterRadius = clusters.reduce(
        (maximum, cluster) => Math.max(maximum, cluster.radius),
        nodeSpacing * 4,
      );
      const minimumAngularStep =
        clusters.length <= 1 ? Math.PI : Math.sin(Math.PI / clusters.length);
      const largestAdjacentClusterSpan = clusters.reduce((maximum, cluster) => {
        const clusterIndex = clusters.indexOf(cluster);
        const nextCluster =
          clusters[(clusterIndex + 1) % clusters.length] ?? cluster;
        return Math.max(
          maximum,
          cluster.radius +
            nextCluster.radius +
            getFootprintGap(
              cluster.radius,
              nextCluster.radius,
              minimumNodeCenterDistance,
            ),
        );
      }, nodeSpacing * 8);
      const islandOrbitRadius =
        clusters.length <= 1
          ? 0
          : Math.max(
              largestAdjacentClusterSpan / (2 * minimumAngularStep),
              largestClusterRadius +
                getFootprintGap(
                  largestClusterRadius,
                  largestClusterRadius,
                  minimumNodeCenterDistance,
                ),
            );

      return {
        key: ringKey,
        clusters,
        labelSeed: getStableLabelSeed(ringKey),
        radius: islandOrbitRadius + largestClusterRadius,
        islandOrbitRadius,
      };
    });

  const placeCluster = (
    cluster: (typeof rings)[number]["clusters"][number],
    centerX: number,
    centerY: number,
    ringCenterX = centerX,
    ringCenterY = centerY,
  ) => {
    const bounds = getIslandOffsetBounds(cluster.offsets);
    const offsetCenterX = (bounds.minX + bounds.maxX) / 2;
    const offsetCenterY = (bounds.minY + bounds.maxY) / 2;

    cluster.nodeIds.forEach((nodeId, nodeIndex) => {
      const offset = cluster.offsets[nodeIndex];
      graph.setNodeAttribute(nodeId, "x", centerX + offset.x - offsetCenterX);
      graph.setNodeAttribute(nodeId, "y", centerY + offset.y - offsetCenterY);
      graph.setNodeAttribute(nodeId, "appLayerRingCenterX", ringCenterX);
      graph.setNodeAttribute(nodeId, "appLayerRingCenterY", ringCenterY);
    });
  };

  const placeRingIslands = (
    ring: (typeof rings)[number],
    centerX: number,
    centerY: number,
  ) => {
    if (ring.clusters.length === 0) return;
    if (ring.clusters.length === 1) {
      placeCluster(ring.clusters[0], centerX, centerY, centerX, centerY);
      return;
    }

    const balancedSlots = getBalancedCircularSlots(ring.clusters.length);
    const clustersByPlacementSize = [...ring.clusters].sort(
      (left, right) =>
        right.radius - left.radius || compareIslandKeys(left.label, right.label),
    );
    const slotByClusterLabel = new Map<string, number>();
    clustersByPlacementSize.forEach((cluster, index) => {
      slotByClusterLabel.set(cluster.label, balancedSlots[index] ?? index);
    });

    ring.clusters.forEach((cluster, clusterIndex) => {
      const clusterSlot = slotByClusterLabel.get(cluster.label) ?? clusterIndex;
      const clusterAngle =
        -Math.PI / 2 +
        (clusterSlot / ring.clusters.length) * Math.PI * 2;
      const clusterOrbitRadius = ring.islandOrbitRadius;
      const clusterCenterX =
        centerX + Math.cos(clusterAngle) * clusterOrbitRadius;
      const clusterCenterY =
        centerY + Math.sin(clusterAngle) * clusterOrbitRadius;

      placeCluster(cluster, clusterCenterX, clusterCenterY, centerX, centerY);
    });
  };

  const documentationRing = rings.find((ring) => ring.key === "docs");
  const outerRings = rings.filter((ring) => ring.key !== "docs");

  if (documentationRing) {
    placeRingIslands(documentationRing, 0, 0);
  }
  if (outerRings.length === 0) return;

  const largestOuterRingRadius = outerRings.reduce(
    (maximum, ring) => Math.max(maximum, ring.radius),
    nodeSpacing * 4,
  );
  const getRingGap = (leftRadius: number, rightRadius: number): number =>
    getFootprintGap(leftRadius, rightRadius, minimumNodeCenterDistance);
  const centerClearance = documentationRing
    ? documentationRing.radius +
      largestOuterRingRadius +
      getRingGap(documentationRing.radius, largestOuterRingRadius)
    : 0;
  const ringAngles = outerRings.map((ring) => getAppLayerRingAngle(ring.key));
  const minimumAngularSeparation = getMinimumAngularSeparation(ringAngles);
  const largestAdjacentRingSpan = outerRings.reduce((maximum, ring) => {
    const ringAngle = getAppLayerRingAngle(ring.key);
    const nextRing =
      outerRings
        .filter((candidate) => candidate !== ring)
        .sort(
          (left, right) =>
            Math.abs(getAppLayerRingAngle(left.key) - ringAngle) -
            Math.abs(getAppLayerRingAngle(right.key) - ringAngle),
        )[0] ?? ring;
    return Math.max(
      maximum,
      ring.radius + nextRing.radius + getRingGap(ring.radius, nextRing.radius),
    );
  }, nodeSpacing * 8);
  const orbitRadius =
    outerRings.length <= 1
      ? centerClearance
      : Math.max(
          largestAdjacentRingSpan / (2 * Math.sin(minimumAngularSeparation / 2)),
          centerClearance,
        );

  outerRings.forEach((ring) => {
    const ringAngle = outerRings.length <= 1 ? 0 : getAppLayerRingAngle(ring.key);
    const ringCenterX = Math.cos(ringAngle) * orbitRadius;
    const ringCenterY = Math.sin(ringAngle) * orbitRadius;

    placeRingIslands(ring, ringCenterX, ringCenterY);
  });
};

