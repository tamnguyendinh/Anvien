import { describe, expect, it } from 'vitest';
import { MultiDirectedGraph } from 'graphology';
import {
  buildReadableGraphCameraState,
  MIN_READABLE_NODE_RADIUS_PX,
} from '../../src/lib/graph-readable-camera';
import type {
  SigmaEdgeAttributes,
  SigmaNodeAttributes,
} from '../../src/lib/graph-adapter';

const createDenseGraph = (nodeCount: number) => {
  const graph = new MultiDirectedGraph<
    SigmaNodeAttributes,
    SigmaEdgeAttributes
  >();

  for (let index = 0; index < nodeCount; index++) {
    graph.addNode(`function-${index}`, {
      x: index * 12,
      y: 0,
      size: 3,
      color: '#8b5cf6',
      label: `function${index}`,
      nodeType: 'Function',
      filePath: `src/function-${index}.ts`,
      appLayerRing: 'frontend',
      islandKey: 'Function',
    });
  }

  graph.addNode('docs-file', {
    x: -1000,
    y: 1000,
    size: 3,
    color: '#3b82f6',
    label: 'docs-file',
    nodeType: 'File',
    filePath: 'docs/file.md',
    appLayerRing: 'docs',
    islandKey: 'File',
  });

  return graph;
};

const createSigma = (
  graph: ReturnType<typeof createDenseGraph>,
  options: { sizeScale: number },
) =>
  ({
    getGraph: () => graph,
    getDimensions: () => ({ width: 1280, height: 800 }),
    getCamera: () => ({
      getState: () => ({ x: 0.5, y: 0.5, ratio: 1, angle: 0 }),
    }),
    getNodeDisplayData: (nodeId: string) => graph.getNodeAttributes(nodeId),
    graphToViewport: ({ x, y }: { x: number; y: number }) => ({
      x: x * 0.1,
      y: y * 0.1,
    }),
    viewportToGraph: ({ x, y }: { x: number; y: number }) => ({
      x: x / 0.1,
      y: y / 0.1,
    }),
    viewportToFramedGraph: ({ x, y }: { x: number; y: number }) => ({
      x: x / 10000,
      y: y / 10000,
    }),
    scaleSize: (size = 1) => size * options.sizeScale,
  }) as any;

describe('buildReadableGraphCameraState', () => {
  it('zooms toward the densest island when dense graph nodes render too small', () => {
    const graph = createDenseGraph(1001);
    const state = buildReadableGraphCameraState(
      createSigma(graph, { sizeScale: 0.1 }),
    );

    expect(state).not.toBeNull();
    expect(state?.focusedIslandKey).toBe('frontend:Function');
    expect(state?.focusedIslandNodeCount).toBe(1001);
    expect(state?.ratio).toBeLessThan(1);
    expect(state?.ratio).toBeCloseTo(Math.pow(0.3 / MIN_READABLE_NODE_RADIUS_PX, 2));
    expect(state?.x).toBeGreaterThan(0);
  });

  it('keeps fit-to-screen camera for small graphs', () => {
    const graph = createDenseGraph(100);

    expect(
      buildReadableGraphCameraState(createSigma(graph, { sizeScale: 0.1 })),
    ).toBeNull();
  });

  it('keeps current camera when nodes are already readable', () => {
    const graph = createDenseGraph(1001);

    expect(
      buildReadableGraphCameraState(createSigma(graph, { sizeScale: 1 })),
    ).toBeNull();
  });
});
