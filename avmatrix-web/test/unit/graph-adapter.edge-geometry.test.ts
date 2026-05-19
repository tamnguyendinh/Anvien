import { afterEach, describe, expect, it, vi } from 'vitest';
import { createKnowledgeGraph } from '../../src/core/graph/graph';
import {
  MAX_DENSE_RENDERED_NODE_SIZE,
  MAX_RENDERED_NODE_SIZE,
  capRenderedNodeSize,
  getMaxRenderedNodeSize,
  getScaledNodeSize,
  knowledgeGraphToGraphology,
} from '../../src/lib/graph-adapter';
import type { GraphRelationship } from '../../src/generated/avmatrix-contracts';
import {
  createCallsRelationship,
  createClassNode,
  createContainsRelationship,
  createFileNode,
  createFunctionNode,
  createProcessNode,
} from '../fixtures/graph';

describe('knowledgeGraphToGraphology edge geometry', () => {
  afterEach(() => {
    vi.restoreAllMocks();
  });

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

  it('keeps package and section nodes below the generic structural cap on large graphs', () => {
    const veryLargeGraphNodeCount = 78_350;

    expect(
      getScaledNodeSize(16, veryLargeGraphNodeCount, 'Package'),
    ).toBeLessThanOrEqual(1.5);
    expect(
      getScaledNodeSize(8, veryLargeGraphNodeCount, 'Section'),
    ).toBeLessThanOrEqual(1);
    expect(getScaledNodeSize(20, veryLargeGraphNodeCount, 'Project')).toBe(3);
  });

  it('keeps structural-to-leaf node size ratios bounded across graph sizes', () => {
    for (const nodeCount of [100, 1_500, 6_000, 20_421, 60_000]) {
      const structuralSize = getScaledNodeSize(20, nodeCount);
      const leafSize = getScaledNodeSize(2, nodeCount);

      expect(structuralSize).toBeGreaterThan(leafSize);
      expect(structuralSize / leafSize).toBeLessThanOrEqual(3);
    }
  });

  it('uses process relationships as higher-priority layout parents than file definitions', () => {
    vi.spyOn(Math, 'random').mockReturnValue(0.5);

    const graph = createKnowledgeGraph();
    const fileNode = createFileNode('workflow.ts', 'src/workflow.ts');
    const functionNode = createFunctionNode('runWorkflow', 'src/workflow.ts', 1);
    const processNode = createProcessNode('proc_0_workflow', 'Workflow');
    const defines: GraphRelationship = {
      id: `${fileNode.id}_DEFINES_${functionNode.id}`,
      sourceId: fileNode.id,
      targetId: functionNode.id,
      type: 'DEFINES',
      confidence: 1,
      reason: 'test-fixture',
    };
    const stepInProcess: GraphRelationship = {
      id: `${functionNode.id}_STEP_IN_PROCESS_${processNode.id}`,
      sourceId: functionNode.id,
      targetId: processNode.id,
      type: 'STEP_IN_PROCESS',
      confidence: 1,
      reason: 'test-fixture',
    };

    graph.addNode(fileNode);
    graph.addNode(functionNode);
    graph.addNode(processNode);
    graph.addRelationship(defines);
    graph.addRelationship(stepInProcess);

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const functionAttributes = sigmaGraph.getNodeAttributes(functionNode.id);
    const processAttributes = sigmaGraph.getNodeAttributes(processNode.id);
    const fileAttributes = sigmaGraph.getNodeAttributes(fileNode.id);

    expect(functionAttributes.x).toBe(processAttributes.x);
    expect(functionAttributes.y).toBe(processAttributes.y);
    expect(functionAttributes.x).not.toBe(fileAttributes.x);
  });

  it('keeps owned properties near their owning type', () => {
    vi.spyOn(Math, 'random').mockReturnValue(0.5);

    const graph = createKnowledgeGraph();
    const fileNode = createFileNode('model.ts', 'src/model.ts');
    const classNode = createClassNode('Model', 'src/model.ts');
    const propertyNode = {
      id: 'Property:src/model.ts:Model.value',
      label: 'Property',
      properties: { name: 'value', filePath: 'src/model.ts' },
    } as const;
    const hasProperty: GraphRelationship = {
      id: `${classNode.id}_HAS_PROPERTY_${propertyNode.id}`,
      sourceId: classNode.id,
      targetId: propertyNode.id,
      type: 'HAS_PROPERTY',
      confidence: 1,
      reason: 'test-fixture',
    };

    graph.addNode(fileNode);
    graph.addNode(classNode);
    graph.addNode(propertyNode);
    graph.addRelationship(createContainsRelationship(fileNode.id, classNode.id));
    graph.addRelationship(hasProperty);

    const sigmaGraph = knowledgeGraphToGraphology(graph);
    const classAttributes = sigmaGraph.getNodeAttributes(classNode.id);
    const propertyAttributes = sigmaGraph.getNodeAttributes(propertyNode.id);

    expect(propertyAttributes.x).toBe(classAttributes.x);
    expect(propertyAttributes.y).toBe(classAttributes.y);
  });
});
