import { describe, expect, it } from 'vitest';
import { MultiDirectedGraph } from 'graphology';
import { buildScreenNodeSpacingDiagnostics } from '../../src/lib/graph-screen-spacing';
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from '../../src/lib/graph-adapter';

const createGraph = () => {
  const graph = new MultiDirectedGraph<
    SigmaNodeAttributes,
    SigmaEdgeAttributes
  >();

  graph.addNode('left', {
    x: 0,
    y: 0,
    size: 3,
    color: '#10b981',
    label: 'left',
    nodeType: 'Function',
    filePath: 'src/left.ts',
    appLayerRing: 'frontend',
    islandKey: 'Function',
  });
  graph.addNode('right', {
    x: 12,
    y: 0,
    size: 3,
    color: '#10b981',
    label: 'right',
    nodeType: 'Function',
    filePath: 'src/right.ts',
    appLayerRing: 'frontend',
    islandKey: 'Function',
  });

  return graph;
};

const createSigma = (
  graph: ReturnType<typeof createGraph>,
  options: { positionScale: number; sizeScale: number },
) =>
  ({
    getGraph: () => graph,
    getDimensions: () => ({ width: 1280, height: 800 }),
    getCamera: () => ({
      getState: () => ({ x: 0.5, y: 0.5, ratio: 1, angle: 0 }),
    }),
    getNodeDisplayData: (nodeId: string) => graph.getNodeAttributes(nodeId),
    graphToViewport: ({ x, y }: { x: number; y: number }) => ({
      x: x * options.positionScale,
      y: y * options.positionScale,
    }),
    viewportToGraph: ({ x, y }: { x: number; y: number }) => ({
      x: x / options.positionScale,
      y: y / options.positionScale,
    }),
    scaleSize: (size = 1) => size * options.sizeScale,
  }) as any;

describe('buildScreenNodeSpacingDiagnostics', () => {
  it('passes when viewport center distance and rendered size use the same scale', () => {
    const graph = createGraph();
    const diagnostics = buildScreenNodeSpacingDiagnostics(
      createSigma(graph, { positionScale: 2, sizeScale: 2 }),
    );

    expect(diagnostics.coordinateSpace).toBe('viewport_px');
    expect(diagnostics.overlapCount).toBe(0);
    expect(diagnostics.targetGapViolationCount).toBe(0);
    expect(diagnostics.minObservedCenterDistance).toBe(24);
    expect(diagnostics.minObservedEdgeGap).toBe(12);
  });

  it('fails when graph coordinates are compressed but node size stays screen-fixed', () => {
    const graph = createGraph();
    const diagnostics = buildScreenNodeSpacingDiagnostics(
      createSigma(graph, { positionScale: 0.25, sizeScale: 1 }),
    );

    expect(diagnostics.overlapCount).toBe(1);
    expect(diagnostics.targetGapViolationCount).toBe(1);
    expect(diagnostics.minObservedCenterDistance).toBe(3);
    expect(diagnostics.minObservedEdgeGap).toBe(-3);
  });

  it('reports the required gap lower bound when nodes are farther than the local search window', () => {
    const graph = createGraph();
    const diagnostics = buildScreenNodeSpacingDiagnostics(
      createSigma(graph, { positionScale: 10, sizeScale: 1 }),
    );

    expect(diagnostics.overlapCount).toBe(0);
    expect(diagnostics.targetGapViolationCount).toBe(0);
    expect(diagnostics.minObservedCenterDistance).toBe(12);
    expect(diagnostics.minObservedEdgeGap).toBe(6);
  });
});
