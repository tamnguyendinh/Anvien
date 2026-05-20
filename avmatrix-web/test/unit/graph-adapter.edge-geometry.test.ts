import { describe, expect, it } from "vitest";
import { createKnowledgeGraph } from "../../src/core/graph/graph";
import {
  MAX_DENSE_RENDERED_NODE_SIZE,
  MAX_RENDERED_NODE_SIZE,
  capRenderedNodeSize,
  getMaxRenderedNodeSize,
  getScaledNodeSize,
  knowledgeGraphToGraphology,
} from "../../src/lib/graph-adapter";
import { DOCUMENTATION_NODE_LABEL, getNodeColor } from "../../src/lib/constants";
import type { GraphRelationship } from "../../src/generated/avmatrix-contracts";
import {
  createCallsRelationship,
  createClassNode,
  createFileNode,
  createFunctionNode,
} from "../fixtures/graph";

const createTypedNode = (
  label: string,
  index: number,
  filePath = `src/${label.toLowerCase()}/${index}.ts`,
) =>
  ({
    id: `${label}:${filePath}:${index}`,
    label,
    properties: {
      name: `${label}${index}`,
      filePath,
      startLine: index + 1,
      endLine: index + 2,
    },
  }) as const;

type ClusterBounds = {
  minX: number;
  maxX: number;
  minY: number;
  maxY: number;
};

type ClusterGeometry = ClusterBounds & {
  centerX: number;
  centerY: number;
  width: number;
  height: number;
  radius: number;
  count: number;
};

const getBoundsByType = (sigmaGraph: ReturnType<typeof knowledgeGraphToGraphology>) => {
  const boundsByType = new Map<
    string,
    ClusterBounds & { count: number }
  >();

  for (const nodeId of sigmaGraph.nodes()) {
    const attributes = sigmaGraph.getNodeAttributes(nodeId);
    const current =
      boundsByType.get(attributes.nodeType) ??
      {
        minX: Number.POSITIVE_INFINITY,
        maxX: Number.NEGATIVE_INFINITY,
        minY: Number.POSITIVE_INFINITY,
        maxY: Number.NEGATIVE_INFINITY,
        count: 0,
      };
    current.minX = Math.min(current.minX, attributes.x);
    current.maxX = Math.max(current.maxX, attributes.x);
    current.minY = Math.min(current.minY, attributes.y);
    current.maxY = Math.max(current.maxY, attributes.y);
    current.count++;
    boundsByType.set(attributes.nodeType, current);
  }

  return boundsByType;
};

const getGeometryByType = (
  sigmaGraph: ReturnType<typeof knowledgeGraphToGraphology>,
) => {
  const geometryByType = new Map<string, ClusterGeometry>();

  for (const [label, bounds] of getBoundsByType(sigmaGraph)) {
    const width = bounds.maxX - bounds.minX;
    const height = bounds.maxY - bounds.minY;
    const centerX = bounds.minX + width / 2;
    const centerY = bounds.minY + height / 2;
    geometryByType.set(label, {
      ...bounds,
      centerX,
      centerY,
      width,
      height,
      radius: Math.hypot(width, height) / 2,
    });
  }

  return geometryByType;
};

const getCircularGap = (
  left: ClusterGeometry,
  right: ClusterGeometry,
): number => {
  return (
    Math.hypot(left.centerX - right.centerX, left.centerY - right.centerY) -
    left.radius -
    right.radius
  );
};

const getAngleProgress = (x: number, y: number): number => {
  const startAngle = -Math.PI / 2;
  const fullCircle = Math.PI * 2;
  return (Math.atan2(y, x) - startAngle + fullCircle) % fullCircle;
};

describe("knowledgeGraphToGraphology edge geometry", () => {
  it("creates straight edges without curved-edge metadata", () => {
    const graph = createKnowledgeGraph();
    const fileNode = createFileNode("index.ts", "src/index.ts");
    const functionNode = createFunctionNode("main", "src/index.ts", 1);

    graph.addNode(fileNode);
    graph.addNode(functionNode);
    graph.addRelationship(
      createCallsRelationship(fileNode.id, functionNode.id),
    );

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const [edgeId] = sigmaGraph.edges();
    const edgeAttributes = sigmaGraph.getEdgeAttributes(edgeId) as Record<
      string,
      unknown
    >;

    expect(edgeId).toBeDefined();
    expect(edgeAttributes.size).toBeGreaterThan(0);
    expect(edgeAttributes).not.toHaveProperty("type");
    expect(edgeAttributes).not.toHaveProperty("curvature");
  });

  it("preserves parallel relationship types between the same source and target", () => {
    const graph = createKnowledgeGraph();
    const fileNode = createFileNode("index.ts", "src/index.ts");
    const functionNode = createFunctionNode("main", "src/index.ts", 1);
    const classNode = createClassNode("Widget", "src/index.ts");
    const propertyNode = {
      id: "Property:src/index.ts:Widget.value",
      label: "Property",
      properties: { name: "value", filePath: "src/index.ts" },
    } as const;
    const calls = createCallsRelationship(fileNode.id, functionNode.id);
    const uses: GraphRelationship = {
      ...calls,
      id: `${fileNode.id}_USES_${functionNode.id}`,
      type: "USES",
    };
    const hasProperty: GraphRelationship = {
      id: `${classNode.id}_HAS_PROPERTY_${propertyNode.id}`,
      sourceId: classNode.id,
      targetId: propertyNode.id,
      type: "HAS_PROPERTY",
      confidence: 1,
      reason: "test-fixture",
    };
    const accesses: GraphRelationship = {
      ...hasProperty,
      id: `${classNode.id}_ACCESSES_${propertyNode.id}`,
      type: "ACCESSES",
    };

    graph.addNode(fileNode);
    graph.addNode(functionNode);
    graph.addNode(classNode);
    graph.addNode(propertyNode);
    graph.addRelationship(calls);
    graph.addRelationship(uses);
    graph.addRelationship(hasProperty);
    graph.addRelationship(accesses);

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const relationTypes = sigmaGraph
      .edges()
      .map((edgeId) => sigmaGraph.getEdgeAttribute(edgeId, "relationType"))
      .sort();

    expect(sigmaGraph.multi).toBe(true);
    expect(sigmaGraph.size).toBe(4);
    expect(relationTypes).toEqual([
      "ACCESSES",
      "CALLS",
      "HAS_PROPERTY",
      "USES",
    ]);
  });

  it("collapses duplicate INHERITS compatibility edges when EXTENDS exists", () => {
    const graph = createKnowledgeGraph();
    const childNode = createClassNode("Child", "src/model.ts");
    const parentNode = createClassNode("Parent", "src/model.ts");

    graph.addNode(childNode);
    graph.addNode(parentNode);
    graph.addRelationship({
      id: "extends-child-parent",
      sourceId: childNode.id,
      targetId: parentNode.id,
      type: "EXTENDS",
      confidence: 1,
      reason: "extends",
    } as GraphRelationship);
    graph.addRelationship({
      id: "inherits-child-parent",
      sourceId: childNode.id,
      targetId: parentNode.id,
      type: "INHERITS",
      confidence: 1,
      reason: "scope-resolution: inherits",
    } as GraphRelationship);

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const relationTypes = sigmaGraph
      .edges()
      .map((edgeId) => sigmaGraph.getEdgeAttribute(edgeId, "relationType"));

    expect(sigmaGraph.size).toBe(1);
    expect(relationTypes).toEqual(["EXTENDS"]);
  });

  it("bounds scaled and rendered node sizes for very large graphs", () => {
    const largeGraphNodeCount = 20_421;
    const projectSize = getScaledNodeSize(20, largeGraphNodeCount);
    const propertySize = getScaledNodeSize(2, largeGraphNodeCount);

    expect(projectSize).toBeLessThanOrEqual(3);
    expect(propertySize).toBeGreaterThanOrEqual(1);
    expect(projectSize / propertySize).toBeLessThanOrEqual(3);
    expect(capRenderedNodeSize(100)).toBe(MAX_RENDERED_NODE_SIZE);
    expect(getMaxRenderedNodeSize(largeGraphNodeCount)).toBe(
      MAX_DENSE_RENDERED_NODE_SIZE,
    );
    expect(capRenderedNodeSize(100, largeGraphNodeCount)).toBe(
      MAX_DENSE_RENDERED_NODE_SIZE,
    );
  });

  it("keeps package and section nodes below the generic structural cap on large graphs", () => {
    const veryLargeGraphNodeCount = 78_350;

    expect(
      getScaledNodeSize(16, veryLargeGraphNodeCount, "Package"),
    ).toBeLessThanOrEqual(1.5);
    expect(
      getScaledNodeSize(8, veryLargeGraphNodeCount, "Section"),
    ).toBeLessThanOrEqual(1);
    expect(getScaledNodeSize(20, veryLargeGraphNodeCount, "Project")).toBe(3);
  });

  it("keeps structural-to-leaf node size ratios bounded across graph sizes", () => {
    for (const nodeCount of [100, 1_500, 6_000, 20_421, 60_000]) {
      const structuralSize = getScaledNodeSize(20, nodeCount);
      const leafSize = getScaledNodeSize(2, nodeCount);

      expect(structuralSize).toBeGreaterThan(leafSize);
      expect(structuralSize / leafSize).toBeLessThanOrEqual(3);
    }
  });

  it("produces deterministic filter-based clustered positions", () => {
    const graph = createKnowledgeGraph();
    const functionA = createFunctionNode("build", "src/b.ts", 1);
    const functionB = createFunctionNode("analyze", "src/a.ts", 1);
    const classNode = createClassNode("Runner", "src/runner.ts");
    const fileNode = createFileNode("index.ts", "src/index.ts");

    graph.addNode(functionA);
    graph.addNode(classNode);
    graph.addNode(fileNode);
    graph.addNode(functionB);

    const first = knowledgeGraphToGraphology(graph);
    const second = knowledgeGraphToGraphology(graph);

    for (const nodeId of first.nodes()) {
      expect(first.getNodeAttribute(nodeId, "x")).toBe(
        second.getNodeAttribute(nodeId, "x"),
      );
      expect(first.getNodeAttribute(nodeId, "y")).toBe(
        second.getNodeAttribute(nodeId, "y"),
      );
    }
  });

  it("keeps each node type in its own separated visual island", () => {
    const graph = createKnowledgeGraph();
    const nodes = [
      createFileNode("a.ts", "src/a.ts"),
      createFileNode("b.ts", "src/b.ts"),
      createClassNode("Alpha", "src/a.ts"),
      createClassNode("Beta", "src/b.ts"),
      createFunctionNode("alpha", "src/a.ts", 1),
      createFunctionNode("beta", "src/b.ts", 1),
    ];

    for (const node of nodes) graph.addNode(node);

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const geometryByType = getGeometryByType(sigmaGraph);

    const geometries = [...geometryByType.values()];
    for (let leftIndex = 0; leftIndex < geometries.length; leftIndex++) {
      for (
        let rightIndex = leftIndex + 1;
        rightIndex < geometries.length;
        rightIndex++
      ) {
        const left = geometries[leftIndex];
        const right = geometries[rightIndex];
        const gap = getCircularGap(left, right);

        expect(gap).toBeGreaterThanOrEqual(100);
      }
    }
  });

  it("keeps imbalanced node type islands visually separated with a large gutter", () => {
    const graph = createKnowledgeGraph();
    const countsByLabel = new Map([
      ["Function", 900],
      ["File", 400],
      ["Class", 100],
      ["Interface", 36],
      ["Enum", 16],
    ]);

    for (const [label, count] of countsByLabel) {
      for (let index = 0; index < count; index++) {
        graph.addNode(createTypedNode(label, index));
      }
    }

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const geometries = [...getGeometryByType(sigmaGraph).entries()];

    for (let leftIndex = 0; leftIndex < geometries.length; leftIndex++) {
      for (
        let rightIndex = leftIndex + 1;
        rightIndex < geometries.length;
        rightIndex++
      ) {
        const [leftLabel, left] = geometries[leftIndex];
        const [rightLabel, right] = geometries[rightIndex];
        const gap = getCircularGap(left, right);

        expect(
          gap,
          `${leftLabel} and ${rightLabel} should be separated enough to read as different color islands`,
        ).toBeGreaterThanOrEqual(500);
      }
    }
  });

  it("lays out medium and large clusters as two-dimensional islands instead of rails", () => {
    const graph = createKnowledgeGraph();
    const countsByLabel = new Map([
      ["Function", 196],
      ["File", 144],
      ["Class", 81],
    ]);

    for (const [label, count] of countsByLabel) {
      for (let index = 0; index < count; index++) {
        graph.addNode(createTypedNode(label, index));
      }
    }

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const geometries = getGeometryByType(sigmaGraph);

    for (const [label, geometry] of geometries) {
      const aspectRatio =
        Math.max(geometry.width, geometry.height) /
        Math.max(1, Math.min(geometry.width, geometry.height));

      expect(
        geometry.width,
        `${label} should have meaningful horizontal spread`,
      ).toBeGreaterThan(250);
      expect(
        geometry.height,
        `${label} should have meaningful vertical spread`,
      ).toBeGreaterThan(250);
      expect(
        aspectRatio,
        `${label} should not collapse into a rail-like shape`,
      ).toBeLessThanOrEqual(2);
    }
  });

  it("places node type islands on one large circular graph field", () => {
    const graph = createKnowledgeGraph();
    const countsByLabel = new Map([
      ["Project", 4],
      ["File", 64],
      ["Function", 100],
      ["Class", 36],
      ["Interface", 16],
      ["Enum", 9],
    ]);

    for (const [label, count] of countsByLabel) {
      for (let index = 0; index < count; index++) {
        graph.addNode(createTypedNode(label, index));
      }
    }

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const centerRadii = [...getGeometryByType(sigmaGraph).values()].map(
      (geometry) => Math.hypot(geometry.centerX, geometry.centerY),
    );
    const minimumRadius = Math.min(...centerRadii);
    const maximumRadius = Math.max(...centerRadii);

    expect(minimumRadius).toBeGreaterThan(0);
    expect(maximumRadius / minimumRadius).toBeLessThanOrEqual(1.2);
  });

  it("uses the node type color for every node even when community metadata exists", () => {
    const graph = createKnowledgeGraph();
    const functionA = createFunctionNode("alpha", "src/a.ts", 1);
    const functionB = createFunctionNode("beta", "src/b.ts", 1);
    const classNode = createClassNode("Runner", "src/runner.ts");

    graph.addNode(functionA);
    graph.addNode(functionB);
    graph.addNode(classNode);

    const communityMemberships = new Map<string, number>([
      [functionA.id, 0],
      [functionB.id, 1],
      [classNode.id, 2],
    ]);
    const sigmaGraph = knowledgeGraphToGraphology(graph, communityMemberships);

    expect(sigmaGraph.getNodeAttribute(functionA.id, "color")).toBe(
      getNodeColor("Function"),
    );
    expect(sigmaGraph.getNodeAttribute(functionB.id, "color")).toBe(
      getNodeColor("Function"),
    );
    expect(sigmaGraph.getNodeAttribute(classNode.id, "color")).toBe(
      getNodeColor("Class"),
    );
    expect(sigmaGraph.getNodeAttribute(functionA.id, "community")).toBe(0);
    expect(sigmaGraph.getNodeAttribute(functionB.id, "community")).toBe(1);
  });

  it("routes documentation files into the Documentation display filter", () => {
    const graph = createKnowledgeGraph();
    const readmeNode = createFileNode("README.md", "README.md");
    const guideSectionNode = createTypedNode("Section", 1, "docs/guide.md");
    const sourceFileNode = createFileNode("index.ts", "src/index.ts");
    const functionNode = createFunctionNode("main", "src/index.ts", 1);

    graph.addNode(readmeNode);
    graph.addNode(guideSectionNode as any);
    graph.addNode(sourceFileNode);
    graph.addNode(functionNode);
    graph.addRelationship({
      id: "docs-uses-main",
      sourceId: readmeNode.id,
      targetId: functionNode.id,
      type: "USES",
      confidence: 1,
      reason: "documentation-link",
    } as GraphRelationship);

    const sigmaGraph = knowledgeGraphToGraphology(graph);

    expect(sigmaGraph.getNodeAttribute(readmeNode.id, "nodeType")).toBe(
      DOCUMENTATION_NODE_LABEL,
    );
    expect(sigmaGraph.getNodeAttribute(readmeNode.id, "rawNodeType")).toBe(
      "File",
    );
    expect(sigmaGraph.getNodeAttribute(readmeNode.id, "color")).toBe(
      getNodeColor(DOCUMENTATION_NODE_LABEL),
    );
    expect(sigmaGraph.getNodeAttribute(guideSectionNode.id, "nodeType")).toBe(
      DOCUMENTATION_NODE_LABEL,
    );
    expect(sigmaGraph.getNodeAttribute(sourceFileNode.id, "nodeType")).toBe(
      "File",
    );
    expect(sigmaGraph.size).toBe(1);
    expect(sigmaGraph.getEdgeAttribute("docs-uses-main", "relationType")).toBe(
      "USES",
    );
  });

  it("places the Documentation island at the center of the outer island circle", () => {
    const graph = createKnowledgeGraph();
    const countsByLabel = new Map([
      ["File", 36],
      ["Function", 49],
      ["Class", 25],
    ]);

    for (let index = 0; index < 36; index++) {
      graph.addNode(createTypedNode("File", index, `docs/guide-${index}.md`));
    }
    for (const [label, count] of countsByLabel) {
      for (let index = 0; index < count; index++) {
        graph.addNode(createTypedNode(label, index, `src/${label}/${index}.ts`));
      }
    }

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const geometries = getGeometryByType(sigmaGraph);
    const documentation = geometries.get(DOCUMENTATION_NODE_LABEL);

    expect(documentation).toBeDefined();
    expect(Math.hypot(documentation!.centerX, documentation!.centerY)).toBeLessThan(
      1,
    );

    for (const [label, geometry] of geometries) {
      if (label === DOCUMENTATION_NODE_LABEL) continue;

      expect(
        getCircularGap(documentation!, geometry),
        `${label} should stay outside the centered Documentation island`,
      ).toBeGreaterThan(200);
      expect(Math.hypot(geometry.centerX, geometry.centerY)).toBeGreaterThan(
        documentation!.radius,
      );
    }
  });

  it("orders known clusters by filter order and appends unknown labels by label", () => {
    const graph = createKnowledgeGraph();
    const fileNode = createFileNode("index.ts", "src/index.ts");
    const zCustomNode = {
      id: "ZCustom:src/z.txt:node",
      label: "ZCustom",
      properties: { name: "z", filePath: "src/z.txt" },
    } as any;
    const aCustomNode = {
      id: "ACustom:src/a.txt:node",
      label: "ACustom",
      properties: { name: "a", filePath: "src/a.txt" },
    } as any;

    graph.addNode(zCustomNode);
    graph.addNode(fileNode);
    graph.addNode(aCustomNode);

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const fileAttributes = sigmaGraph.getNodeAttributes(fileNode.id);
    const aCustomAttributes = sigmaGraph.getNodeAttributes(aCustomNode.id);
    const zCustomAttributes = sigmaGraph.getNodeAttributes(zCustomNode.id);

    const fileProgress = getAngleProgress(fileAttributes.x, fileAttributes.y);
    const aCustomProgress = getAngleProgress(
      aCustomAttributes.x,
      aCustomAttributes.y,
    );
    const zCustomProgress = getAngleProgress(
      zCustomAttributes.x,
      zCustomAttributes.y,
    );

    expect(fileProgress).toBeLessThan(aCustomProgress);
    expect(aCustomProgress).toBeLessThan(zCustomProgress);
  });

  it("sorts nodes inside a label cluster by file path, name, then id in the island spiral", () => {
    const graph = createKnowledgeGraph();
    const bPath = createFunctionNode("alpha", "src/b.ts", 1);
    const aPathBeta = createFunctionNode("beta", "src/a.ts", 1);
    const aPathAlphaLine2 = createFunctionNode("alpha", "src/a.ts", 2);
    const aPathAlphaLine1 = createFunctionNode("alpha", "src/a.ts", 1);

    graph.addNode(bPath);
    graph.addNode(aPathBeta);
    graph.addNode(aPathAlphaLine2);
    graph.addNode(aPathAlphaLine1);

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const first = sigmaGraph.getNodeAttributes(aPathAlphaLine1.id);
    const second = sigmaGraph.getNodeAttributes(aPathAlphaLine2.id);
    const third = sigmaGraph.getNodeAttributes(aPathBeta.id);
    const fourth = sigmaGraph.getNodeAttributes(bPath.id);

    const distanceFromFirst = (node: typeof first) =>
      Math.hypot(node.x - first.x, node.y - first.y);

    expect(distanceFromFirst(first)).toBe(0);
    expect(distanceFromFirst(second)).toBeGreaterThan(30);
    expect(distanceFromFirst(third)).toBeGreaterThan(distanceFromFirst(second));
    expect(distanceFromFirst(fourth)).toBeGreaterThan(distanceFromFirst(third));
  });
});
