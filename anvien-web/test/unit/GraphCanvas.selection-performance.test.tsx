import { act, fireEvent, render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { createKnowledgeGraph } from '../../src/core/graph/graph';
import type { GraphNode } from '@/generated/anvien-contracts';
import { AppStateProvider, useAppState } from '../../src/hooks/useAppState.local-runtime';

const {
  filterGraphByDepthSpy,
  filterGraphByLabelsSpy,
  sigmaRefreshSpy,
  setSigmaSelectedNodeSpy,
  setSigmaGraphSpy,
  startLayoutSpy,
  sigmaOnSpy,
  sigmaOffSpy,
  cameraOnSpy,
  cameraOffSpy,
  sigmaGraph,
  containerRefMock,
  sigmaRefMock,
} = vi.hoisted(() => {
  const sigmaGraph = {
    order: 1,
    nodes: () => ['Function:src/foo.ts:loadFoo'],
    getNodeAttributes: () => ({
      x: -100,
      y: 0,
      size: 2,
      color: '#10b981',
      label: 'loadFoo',
      nodeType: 'Function',
      appLayerRing: 'backend',
      islandKey: 'Function',
      appLayerRingCenterX: -100,
      appLayerRingCenterY: 0,
    }),
    getNodeAttribute: (_nodeId: string, attribute: string) =>
      attribute === 'nodeType' ? 'Function' : 2,
    forEachNode: (callback: (nodeId: string, attributes: any) => void) => {
      callback('Function:src/foo.ts:loadFoo', {
        x: -100,
        y: 0,
        size: 2,
        color: '#10b981',
        label: 'loadFoo',
        nodeType: 'Function',
        appLayerRing: 'backend',
        islandKey: 'Function',
        appLayerRingCenterX: -100,
        appLayerRingCenterY: 0,
      });
    },
  };
  const sigmaOnSpy = vi.fn();
  const sigmaOffSpy = vi.fn();
  const cameraOnSpy = vi.fn();
  const cameraOffSpy = vi.fn();
  const sigmaInstance = {
    getGraph: () => sigmaGraph,
    getDimensions: () => ({ width: 800, height: 600 }),
    graphToViewport: ({ x, y }: { x: number; y: number }) => ({
      x: x + 400,
      y: y + 300,
    }),
    viewportToGraph: ({ x, y }: { x: number; y: number }) => ({
      x: x - 400,
      y: y - 300,
    }),
    scaleSize: (size = 1) => size,
    getNodeDisplayData: () => ({
      x: 300,
      y: 300,
      size: 2,
      color: '#10b981',
      label: 'loadFoo',
      nodeType: 'Function',
      appLayerRing: 'backend',
      islandKey: 'Function',
    }),
    getCamera: () => ({
      getState: () => ({ x: 0, y: 0, angle: 0, ratio: 1 }),
      on: cameraOnSpy,
      off: cameraOffSpy,
    }),
    on: sigmaOnSpy,
    off: sigmaOffSpy,
    refresh: vi.fn(),
  };
  return {
    filterGraphByDepthSpy: vi.fn(),
    filterGraphByLabelsSpy: vi.fn(),
    sigmaRefreshSpy: sigmaInstance.refresh,
    setSigmaSelectedNodeSpy: vi.fn(),
    setSigmaGraphSpy: vi.fn(),
    startLayoutSpy: vi.fn(),
    sigmaOnSpy,
    sigmaOffSpy,
    cameraOnSpy,
    cameraOffSpy,
    sigmaGraph,
    containerRefMock: { current: null },
    sigmaRefMock: { current: sigmaInstance },
  };
});

vi.mock('../../src/components/QueryFAB', () => ({
  QueryFAB: () => null,
}));

vi.mock('../../src/hooks/useSigma', () => ({
  useSigma: () => ({
    containerRef: containerRefMock,
    sigmaRef: sigmaRefMock,
    setGraph: setSigmaGraphSpy,
    zoomIn: vi.fn(),
    zoomOut: vi.fn(),
    resetZoom: vi.fn(),
    focusNode: vi.fn(),
    isLayoutRunning: false,
    startLayout: startLayoutSpy,
    stopLayout: vi.fn(),
    selectedNode: null,
    setSelectedNode: setSigmaSelectedNodeSpy,
    refreshHighlights: vi.fn(),
  }),
}));

vi.mock('../../src/lib/graph-adapter', () => ({
  MAX_RENDERED_NODE_SIZE: 9,
  getMaxRenderedNodeSize: vi.fn(() => 9),
  getMinimumNodeEdgeGap: vi.fn(() => 18),
  getMinimumNodeCenterDistance: vi.fn(() => 36),
  knowledgeGraphToGraphology: vi.fn(() => sigmaGraph),
  filterGraphByDepth: filterGraphByDepthSpy,
  filterGraphByLabels: filterGraphByLabelsSpy,
}));

import { GraphCanvas } from '../../src/components/GraphCanvas';

let appState: ReturnType<typeof useAppState> | null = null;

function Harness() {
  appState = useAppState();
  return null;
}

const createFunctionNode = (
  id: string,
  name: string,
  filePath: string,
  startLine: number,
  endLine: number,
): GraphNode => ({
  id,
  label: 'Function',
  properties: {
    name,
    filePath,
    startLine,
    endLine,
  },
});

describe('GraphCanvas selection performance guards', () => {
  beforeEach(() => {
    appState = null;
    filterGraphByDepthSpy.mockReset();
    filterGraphByLabelsSpy.mockReset();
    sigmaRefreshSpy.mockReset();
    setSigmaSelectedNodeSpy.mockReset();
    setSigmaGraphSpy.mockReset();
    startLayoutSpy.mockReset();
  });

  it('does not re-run full graph filtering when selection changes and depthFilter is null', async () => {
    render(
      <AppStateProvider>
        <Harness />
        <GraphCanvas />
      </AppStateProvider>,
    );

    await waitFor(() => expect(appState).not.toBeNull());

    const graph = createKnowledgeGraph();
    const selectedNode = createFunctionNode(
      'Function:src/foo.ts:loadFoo',
      'loadFoo',
      'src/foo.ts',
      10,
      20,
    );
    graph.addNode(selectedNode);

    await act(async () => {
      appState!.setGraph(graph);
    });

    filterGraphByDepthSpy.mockClear();
    filterGraphByLabelsSpy.mockClear();
    sigmaRefreshSpy.mockClear();

    await act(async () => {
      appState!.setSelectedNode(selectedNode);
    });

    expect(filterGraphByDepthSpy).not.toHaveBeenCalled();
    expect(filterGraphByLabelsSpy).not.toHaveBeenCalled();
  });

  it('still re-applies depth filtering when selection changes and depthFilter is enabled', async () => {
    render(
      <AppStateProvider>
        <Harness />
        <GraphCanvas />
      </AppStateProvider>,
    );

    await waitFor(() => expect(appState).not.toBeNull());

    const graph = createKnowledgeGraph();
    const selectedNode = createFunctionNode(
      'Function:src/foo.ts:loadFoo',
      'loadFoo',
      'src/foo.ts',
      10,
      20,
    );
    graph.addNode(selectedNode);

    await act(async () => {
      appState!.setGraph(graph);
      appState!.setDepthFilter(2);
    });

    filterGraphByDepthSpy.mockClear();
    filterGraphByLabelsSpy.mockClear();
    sigmaRefreshSpy.mockClear();

    await act(async () => {
      appState!.setSelectedNode(selectedNode);
    });

    expect(filterGraphByDepthSpy).toHaveBeenCalledWith(
      sigmaGraph,
      selectedNode.id,
      2,
      expect.any(Array),
      expect.objectContaining({
        visibleTopologyStatuses: expect.any(Array),
      }),
      expect.objectContaining({
        visibleAppLayers: expect.any(Array),
        visibleResolutionHealthBuckets: expect.any(Array),
      }),
    );
  });

  it('does not invoke manual layout optimization during graph load', async () => {
    render(
      <AppStateProvider>
        <Harness />
        <GraphCanvas />
      </AppStateProvider>,
    );

    await waitFor(() => expect(appState).not.toBeNull());

    const graph = createKnowledgeGraph();
    graph.addNode(
      createFunctionNode(
        'Function:src/foo.ts:loadFoo',
        'loadFoo',
        'src/foo.ts',
        10,
        20,
      ),
    );

    await act(async () => {
      appState!.setGraph(graph);
    });

    expect(startLayoutSpy).not.toHaveBeenCalled();

    fireEvent.click(screen.getByRole('button', { name: 'Optimize Layout' }));

    expect(startLayoutSpy).toHaveBeenCalledTimes(1);
  });

  it('does not invoke manual layout optimization during filter changes', async () => {
    render(
      <AppStateProvider>
        <Harness />
        <GraphCanvas />
      </AppStateProvider>,
    );

    await waitFor(() => expect(appState).not.toBeNull());

    const graph = createKnowledgeGraph();
    graph.addNode(
      createFunctionNode(
        'Function:src/foo.ts:loadFoo',
        'loadFoo',
        'src/foo.ts',
        10,
        20,
      ),
    );

    await act(async () => {
      appState!.setGraph(graph);
    });

    startLayoutSpy.mockClear();

    await act(async () => {
      appState!.toggleLabelVisibility('Function');
      appState!.toggleSemanticAppLayer('backend');
      appState!.toggleResolutionConfidence('degraded');
    });

    expect(startLayoutSpy).not.toHaveBeenCalled();
  });
});
