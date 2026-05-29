import {
  COMMUNITY_COLORED_NODE_LABELS,
  getNodeSize,
} from "./constants";
import {
  GRAPH_RENDER_SCALE_POLICY,
  capRenderedNodeRadiusByPolicy,
  getPolicyMaxRenderedNodeRadiusPx,
} from "./graph-scale-model";

const getMaxScaledNodeSize = (_nodeCount: number): number => {
  return 3;
};

const getLabelScaledNodeSizeCap = (
  label: string | undefined,
  nodeCount: number,
): number => {
  if (label === "Package") {
    if (nodeCount > 20000) return 1.5;
    if (nodeCount > 5000) return 2;
    return 3;
  }
  if (label === "Section") {
    if (nodeCount > 20000) return 1;
    if (nodeCount > 5000) return 1.5;
    return 2;
  }
  return getMaxScaledNodeSize(nodeCount);
};

export const MAX_RENDERED_NODE_SIZE = 3;
export const MAX_DENSE_RENDERED_NODE_SIZE = 3;

export const getMaxRenderedNodeSize = (nodeCount: number): number =>
  getPolicyMaxRenderedNodeRadiusPx(nodeCount, GRAPH_RENDER_SCALE_POLICY);

export const capRenderedNodeSize = (
  size: number,
  nodeCount: number = 0,
): number => capRenderedNodeRadiusByPolicy(size, nodeCount);

export const getRenderedNodeRadius = (nodeCount: number = 0): number =>
  getMaxRenderedNodeSize(nodeCount);

export const getRenderedNodeDiameter = (nodeCount: number = 0): number =>
  getRenderedNodeRadius(nodeCount) * 2;

export const getMinimumNodeEdgeGap = (nodeCount: number = 0): number =>
  getRenderedNodeDiameter(nodeCount);

export const getMinimumNodeCenterDistance = (nodeCount: number = 0): number =>
  getRenderedNodeDiameter(nodeCount) + getMinimumNodeEdgeGap(nodeCount);

export const getScaledNodeSize = (
  baseSize: number,
  nodeCount: number,
  label?: string,
): number => {
  let scaledSize = baseSize;
  if (nodeCount > 50000) scaledSize = Math.max(1, baseSize * 0.4);
  else if (nodeCount > 20000) scaledSize = Math.max(1.5, baseSize * 0.5);
  else if (nodeCount > 5000) scaledSize = Math.max(2, baseSize * 0.65);
  else if (nodeCount > 1000) scaledSize = Math.max(2.5, baseSize * 0.8);

  return Math.min(scaledSize, getLabelScaledNodeSizeCap(label, nodeCount));
};

const communityColoredNodeLabelSet = new Set<string>(
  COMMUNITY_COLORED_NODE_LABELS,
);

export const getNodeMass = (nodeType: string, nodeCount: number): number => {
  const baseMassMultiplier = nodeCount > 5000 ? 2 : nodeCount > 1000 ? 1.5 : 1;

  const structuralRank: Record<string, number> = {
    Project: 50,
    Package: 30,
    Module: 20,
    Namespace: 20,
    Folder: 15,
    File: 3,
    Section: 2,
    Community: 2,
    Process: 4,
    Route: 4,
    Tool: 4,
  };
  const structuralMass = structuralRank[nodeType];
  if (structuralMass !== undefined) {
    return structuralMass * baseMassMultiplier;
  }
  if (communityColoredNodeLabelSet.has(nodeType)) {
    return 3 * baseMassMultiplier;
  }
  return Math.max(1, Math.min(2, getNodeSize(nodeType))) * baseMassMultiplier;
};

