import { describe, expect, it } from 'vitest';
import { createKnowledgeGraph } from '../../src/core/graph/graph';
import {
  MAX_RENDERED_NODE_SIZE,
  capRenderedNodeSize,
  getScaledNodeSize,
  knowledgeGraphToGraphology,
} from '../../src/lib/graph-adapter';
import type { GraphRelationship } from '../../src/generated/avmatrix-contracts';
import {
  createCallsRelationship,
  createClassNode,
  createFileNode,
  createFunctionNode,
} from '../fixtures/graph';

describe('knowledgeGraphToGraphology edge geometry', () => {
  it('creates straight edges without curved-edge metadata', () => {
    const graph = createKnowledgeGraph();
    const fileNode = createFileNode('index.ts', 'src/index.ts');
    const functionNode = createFunctionNode('main', 'src/index.ts', 1);

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
    expect(edgeAttributes).not.toHaveProperty('type');
    expect(edgeAttributes).not.toHaveProperty('curvature');
  });

  it('preserves parallel relationship types between the same source and target', () => {
    const graph = createKnowledgeGraph();
    const fileNode = createFileNode('index.ts', 'src/index.ts');
    const functionNode = createFunctionNode('main', 'src/index.ts', 1);
    const classNode = createClassNode('Widget', 'src/index.ts');
    const propertyNode = {
      id: 'Property:src/index.ts:Widget.value',
      label: 'Property',
      properties: { name: 'value', filePath: 'src/index.ts' },
    } as const;
    const calls = createCallsRelationship(fileNode.id, functionNode.id);
    const uses: GraphRelationship = {
      ...calls,
      id: `${fileNode.id}_USES_${functionNode.id}`,
      type: 'USES',
    };
    const hasProperty: GraphRelationship = {
      id: `${classNode.id}_HAS_PROPERTY_${propertyNode.id}`,
      sourceId: classNode.id,
      targetId: propertyNode.id,
      type: 'HAS_PROPERTY',
      confidence: 1,
      reason: 'test-fixture',
    };
    const accesses: GraphRelationship = {
      ...hasProperty,
      id: `${classNode.id}_ACCESSES_${propertyNode.id}`,
      type: 'ACCESSES',
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
      .map((edgeId) => sigmaGraph.getEdgeAttribute(edgeId, 'relationType'))
      .sort();

    expect(sigmaGraph.multi).toBe(true);
    expect(sigmaGraph.size).toBe(4);
    expect(relationTypes).toEqual([
      'ACCESSES',
      'CALLS',
      'HAS_PROPERTY',
      'USES',
    ]);
  });

  it('bounds scaled and rendered node sizes for very large graphs', () => {
    const largeGraphNodeCount = 20_421;
    const projectSize = getScaledNodeSize(20, largeGraphNodeCount);
    const propertySize = getScaledNodeSize(2, largeGraphNodeCount);

    expect(projectSize).toBeLessThanOrEqual(4.5);
    expect(propertySize).toBeGreaterThanOrEqual(1.5);
    expect(projectSize / propertySize).toBeLessThanOrEqual(3);
    expect(capRenderedNodeSize(100)).toBe(MAX_RENDERED_NODE_SIZE);
  });

  it('keeps structural-to-leaf node size ratios bounded across graph sizes', () => {
    for (const nodeCount of [100, 1_500, 6_000, 20_421, 60_000]) {
      const structuralSize = getScaledNodeSize(20, nodeCount);
      const leafSize = getScaledNodeSize(2, nodeCount);

      expect(structuralSize).toBeGreaterThan(leafSize);
      expect(structuralSize / leafSize).toBeLessThanOrEqual(6);
    }
  });
});
