import Graph from 'graphology';
import type Sigma from 'sigma';
import { describe, expect, it } from 'vitest';
import { buildGraphOverviewDiagnostics } from '../../src/lib/graph-overview-diagnostics';
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from '../../src/lib/graph-adapter';

const createSigma = (
  graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>,
): Sigma =>
  ({
    getGraph: () => graph,
    getDimensions: () => ({ width: 100, height: 100 }),
    getCamera: () => ({
      getState: () => ({ x: 0.5, y: 0.5, ratio: 1, angle: 0 }),
    }),
    getNodeDisplayData: (nodeId: string) => graph.getNodeAttributes(nodeId),
    graphToViewport: (point: { x: number; y: number }) => point,
  }) as unknown as Sigma;

describe('graph overview diagnostics', () => {
  it('counts visible colors, islands, and node-type inventories from graph data', () => {
    const graph = new Graph<SigmaNodeAttributes, SigmaEdgeAttributes>();
    graph.addNode('frontend-function', {
      x: 10,
      y: 10,
      size: 2,
      color: '#22c55e',
      label: 'fn',
      nodeType: 'Function',
      rawNodeType: 'Function',
      hidden: false,
      appLayerRing: 'frontend',
      islandKey: 'Function',
    });
    graph.addNode('backend-function', {
      x: 40,
      y: 40,
      size: 2,
      color: '#3b82f6',
      label: 'fn',
      nodeType: 'Function',
      rawNodeType: 'Function',
      hidden: false,
      appLayerRing: 'backend',
      islandKey: 'Function',
    });
    graph.addNode('docs', {
      x: 140,
      y: 140,
      size: 2,
      color: '#84cc16',
      label: 'doc',
      nodeType: 'Documentation',
      rawNodeType: 'File',
      hidden: false,
      appLayerRing: 'docs',
      islandKey: 'Documentation',
    });

    const diagnostics = buildGraphOverviewDiagnostics(createSigma(graph));
    const expectedVisibleRings = ['backend', 'frontend'];
    const expectedVisibleNodeTypes = ['Function'];
    const expectedFilterRings = ['backend', 'docs', 'frontend'];
    const expectedFilterNodeTypes = ['Documentation', 'Function'];

    expect(diagnostics.nodeCount).toBe(graph.order);
    expect(diagnostics.visibleViewportNodeCount).toBe(2);
    expect(diagnostics.visibleColorCount).toBe(
      Object.keys(diagnostics.visibleColorCounts).length,
    );
    expect(diagnostics.visibleIslandCount).toBe(
      Object.keys(diagnostics.visibleIslandCounts).length,
    );
    expect(diagnostics.visibleRingInventory).toEqual(expectedVisibleRings);
    expect(diagnostics.visibleNodeTypeInventory).toEqual(
      expectedVisibleNodeTypes,
    );
    expect(diagnostics.graphRingInventory).toEqual(expectedFilterRings);
    expect(diagnostics.graphIslandInventory).toEqual([
      'backend:Function',
      'docs:Documentation',
      'frontend:Function',
    ]);
    expect(diagnostics.filterNodeTypeInventory).toEqual(expectedFilterNodeTypes);
    expect(diagnostics.dominantIslandShare).toBe(0.5);
  });
});
