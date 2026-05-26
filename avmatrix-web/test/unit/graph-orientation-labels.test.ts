import Graph from "graphology";
import { describe, expect, it } from "vitest";
import {
  buildGraphOrientationLabels,
  getOrientationLabelPresentation,
  placeGraphOrientationLabels,
  type GraphOrientationLabel,
} from "../../src/lib/graph-orientation-labels";
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from "../../src/lib/graph-adapter";

const createGraph = () =>
  new Graph<SigmaNodeAttributes, SigmaEdgeAttributes>();

const addNode = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
  id: string,
  attributes: Partial<SigmaNodeAttributes>,
) => {
  graph.addNode(id, {
    x: 0,
    y: 0,
    size: 2,
    color: "#10b981",
    label: id,
    nodeType: "Function",
    filePath: `src/${id}.ts`,
    ...attributes,
  });
};

const makeLabel = (
  id: string,
  kind: "ring" | "island",
  x: number,
  y: number,
  count = 1,
): GraphOrientationLabel => ({
  id,
  kind,
  displayText: id,
  fallbackText: "Unknown",
  sourceKey: id,
  ringKey: "backend",
  islandKey: kind === "island" ? id : undefined,
  anchorX: x,
  anchorY: y,
  visibleNodeCount: count,
  bounds: { minX: x, maxX: x, minY: y, maxY: y },
});

describe("graph orientation labels", () => {
  it("builds ring and island labels with source keys, counts, and stable anchors", () => {
    const graph = createGraph();
    addNode(graph, "a", {
      x: -20,
      y: 10,
      appLayerRing: "backend",
      islandKey: "Function",
      appLayerRingCenterX: -10,
      appLayerRingCenterY: 0,
    });
    addNode(graph, "b", {
      x: 0,
      y: 30,
      nodeType: "Method",
      appLayerRing: "backend",
      islandKey: "Function",
      appLayerRingCenterX: -10,
      appLayerRingCenterY: 0,
    });

    const labels = buildGraphOrientationLabels(graph);
    const ring = labels.find((label) => label.id === "ring:backend");
    const island = labels.find((label) => label.id === "island:backend:Function");

    expect(ring).toMatchObject({
      kind: "ring",
      displayText: "Backend",
      sourceKey: "backend",
      visibleNodeCount: 2,
      anchorX: -10,
    });
    expect(ring?.anchorY).toBeLessThan(10);
    expect(island).toMatchObject({
      kind: "island",
      displayText: "Function",
      sourceKey: "backend:Function",
      ringKey: "backend",
      islandKey: "Function",
      visibleNodeCount: 2,
      anchorX: -10,
    });
    expect(island?.anchorY).toBeLessThan(10);
  });

  it("uses hidden nodes from filters and depth filtering to update visible counts", () => {
    const graph = createGraph();
    addNode(graph, "visible", {
      x: 0,
      y: 0,
      appLayerRing: "api",
      islandKey: "Route",
    });
    addNode(graph, "hidden", {
      x: 10,
      y: 0,
      appLayerRing: "api",
      islandKey: "Route",
      hidden: true,
    });

    const labels = buildGraphOrientationLabels(graph);

    expect(labels.find((label) => label.id === "ring:api")?.visibleNodeCount).toBe(1);
    expect(labels.find((label) => label.id === "island:api:Route")?.visibleNodeCount).toBe(1);

    graph.setNodeAttribute("visible", "hidden", true);

    expect(buildGraphOrientationLabels(graph)).toEqual([]);
  });

  it("formats fallback and ResolutionGap island labels from source keys", () => {
    const graph = createGraph();
    addNode(graph, "gap", {
      nodeType: "ResolutionGap",
      appLayerRing: "custom_runtime_layer",
      islandKey: "ResolutionGap:standard_library",
    });

    const labels = buildGraphOrientationLabels(graph);

    expect(labels.find((label) => label.kind === "ring")?.displayText).toBe(
      "Custom Runtime Layer",
    );
    expect(labels.find((label) => label.kind === "island")?.displayText).toBe(
      "ResolutionGap / Standard Library",
    );
  });

  it("keeps rings visible at far zoom and simplifies small island labels", () => {
    expect(
      getOrientationLabelPresentation(makeLabel("backend", "ring", 0, 0), 20),
    ).toEqual({ visible: true, compact: false });
    expect(
      getOrientationLabelPresentation(makeLabel("Function", "island", 0, 0, 1), 20),
    ).toEqual({ visible: false, compact: true });
    expect(
      getOrientationLabelPresentation(makeLabel("Method", "island", 0, 0, 4), 20),
    ).toEqual({ visible: true, compact: true });
    expect(
      getOrientationLabelPresentation(makeLabel("Function", "island", 0, 0, 1), 1),
    ).toEqual({ visible: true, compact: false });
  });

  it("places viewport labels with deterministic overlap guardrails", () => {
    const labels = [
      makeLabel("backend", "ring", 100, 100, 10),
      makeLabel("Function", "island", 100, 100, 8),
      makeLabel("Method", "island", 260, 100, 5),
    ];

    const placed = placeGraphOrientationLabels(labels, {
      viewportWidth: 500,
      viewportHeight: 300,
      cameraRatio: 1,
      project: (point) => point,
    });

    expect(placed.map((label) => label.id)).toEqual([
      "backend",
      "Function",
      "Method",
    ]);
    expect(placed.every((label) => label.viewportX > 0 && label.viewportY > 0)).toBe(
      true,
    );
    const boxes = placed.map((label) => ({
      left: label.viewportX - label.width / 2,
      right: label.viewportX + label.width / 2,
      top: label.viewportY - label.height / 2,
      bottom: label.viewportY + label.height / 2,
    }));
    for (let leftIndex = 0; leftIndex < boxes.length; leftIndex++) {
      for (let rightIndex = leftIndex + 1; rightIndex < boxes.length; rightIndex++) {
        expect(
          boxes[leftIndex].left < boxes[rightIndex].right &&
            boxes[leftIndex].right > boxes[rightIndex].left &&
            boxes[leftIndex].top < boxes[rightIndex].bottom &&
            boxes[leftIndex].bottom > boxes[rightIndex].top,
        ).toBe(false);
      }
    }
  });

  it("clamps ring labels into the viewport instead of dropping visible rings", () => {
    const placed = placeGraphOrientationLabels(
      [makeLabel("api", "ring", -500, -500, 6)],
      {
        viewportWidth: 320,
        viewportHeight: 240,
        cameraRatio: 1,
        project: (point) => point,
      },
    );

    expect(placed).toHaveLength(1);
    expect(placed[0].viewportX).toBeGreaterThan(0);
    expect(placed[0].viewportY).toBeGreaterThan(0);
  });

  it("separates clamped ring labels on narrow viewports", () => {
    const placed = placeGraphOrientationLabels(
      [
        makeLabel("backend", "ring", 800, 110, 13),
        makeLabel("docs", "ring", 820, 115, 4),
        makeLabel("frontend", "ring", 840, 120, 10),
      ],
      {
        viewportWidth: 320,
        viewportHeight: 240,
        cameraRatio: 1,
        project: (point) => point,
      },
    );

    expect(placed).toHaveLength(3);
    const boxes = placed.map((label) => ({
      left: label.viewportX - label.width / 2,
      right: label.viewportX + label.width / 2,
      top: label.viewportY - label.height / 2,
      bottom: label.viewportY + label.height / 2,
    }));
    for (let leftIndex = 0; leftIndex < boxes.length; leftIndex++) {
      for (let rightIndex = leftIndex + 1; rightIndex < boxes.length; rightIndex++) {
        expect(
          boxes[leftIndex].left < boxes[rightIndex].right &&
            boxes[leftIndex].right > boxes[rightIndex].left &&
            boxes[leftIndex].top < boxes[rightIndex].bottom &&
            boxes[leftIndex].bottom > boxes[rightIndex].top,
        ).toBe(false);
      }
    }
  });
});
