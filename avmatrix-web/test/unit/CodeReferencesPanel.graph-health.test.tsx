import { act, render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import type { GraphNode } from '@/generated/avmatrix-contracts';
import { createKnowledgeGraph } from '../../src/core/graph/graph';
import { AppStateProvider, useAppState } from '../../src/hooks/useAppState.local-runtime';

const readFileMock = vi.hoisted(() =>
  vi.fn().mockResolvedValue({
    content: 'export function loadThing() {\n  return 1;\n}\n',
    startLine: 0,
  }),
);

vi.mock('../../src/services/backend-client', async (importOriginal) => ({
  ...(await importOriginal<typeof import('../../src/services/backend-client')>()),
  readFile: readFileMock,
}));

vi.mock('react-syntax-highlighter', () => ({
  Prism: ({ children }: { children: string }) => <pre>{children}</pre>,
}));

vi.mock('react-syntax-highlighter/dist/esm/styles/prism', () => ({
  vscDarkPlus: {
    'pre[class*="language-"]': {},
    'code[class*="language-"]': {},
  },
}));

import { CodeReferencesPanel } from '../../src/components/CodeReferencesPanel';

let appState: ReturnType<typeof useAppState> | null = null;

function StateCapture() {
  appState = useAppState();
  return null;
}

const makeNode = (
  id: string,
  graphHealth: GraphNode['properties']['graphHealth'],
): GraphNode => ({
  id,
  label: 'Function',
  properties: {
    name: id,
    filePath: 'src/load.ts',
    startLine: 0,
    endLine: 2,
    graphHealth,
  },
});

describe('CodeReferencesPanel Graph Health detail', () => {
  it('explains selected node health and focuses detached components', async () => {
    const selected = makeNode('comp-a', {
      topologyStatus: 'detached_component',
      countedIncoming: 0,
      countedOutgoing: 2,
      excludedEdgeCounts: { structural: 3 },
      componentId: 'component-1',
      componentSize: 2,
      componentReachableFromRoot: false,
      expectedIsolationReasons: ['test'],
      diagnostics: [{ kind: 'unresolved_reference', targetText: 'loadThing', count: 2 }],
      confidence: 'unknown',
    });
    const peer = makeNode('comp-b', {
      topologyStatus: 'detached_component',
      countedIncoming: 1,
      countedOutgoing: 1,
      componentId: 'component-1',
      componentSize: 2,
      componentReachableFromRoot: false,
      confidence: 'candidate',
    });
    const graph = createKnowledgeGraph();
    graph.addNode(selected);
    graph.addNode(peer);
    const onFocusNode = vi.fn();

    render(
      <AppStateProvider>
        <StateCapture />
        <CodeReferencesPanel onFocusNode={onFocusNode} />
      </AppStateProvider>,
    );

    act(() => {
      appState!.setGraph(graph);
      appState!.setSelectedNode(selected);
    });

    const detail = await screen.findByTestId('graph-health-node-detail');
    expect(detail).toHaveTextContent('Detached component');
    expect(detail).toHaveTextContent('Unknown');
    expect(detail).toHaveTextContent(/In\s+0/);
    expect(detail).toHaveTextContent(/Out\s+2/);
    expect(detail).toHaveTextContent(/Comp\s+2/);
    expect(detail).toHaveTextContent('Structural 3');
    expect(detail).toHaveTextContent('Test');
    expect(detail).toHaveTextContent('Unresolved reference x2: loadThing');
    expect(detail).toHaveTextContent('Detached: no accepted root reaches this counted-edge component.');
    expect(detail).toHaveTextContent('Next: Review only if this expected-isolated overlay looks wrong for the node.');

    await userEvent.click(screen.getByRole('button', { name: 'Focus component' }));

    expect(onFocusNode).toHaveBeenCalledWith('comp-a');
    expect([...appState!.highlightedNodeIds].sort()).toEqual(['comp-a', 'comp-b']);
  });
});
