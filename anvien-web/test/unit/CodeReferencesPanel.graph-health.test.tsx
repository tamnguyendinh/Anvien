import { act, render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import type { HTMLAttributes } from 'react';
import type { GraphNode, GraphRelationship } from '@/generated/anvien-contracts';
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
  Prism: ({
    children,
    lineProps,
    showLineNumbers,
    startingLineNumber = 1,
  }: {
    children: string;
    lineProps?: (lineNumber: number) => HTMLAttributes<HTMLSpanElement>;
    showLineNumbers?: boolean;
    startingLineNumber?: number;
  }) => (
    <pre>
      {children.split('\n').map((line, index) => {
        const lineNumber = startingLineNumber + index;
        return (
          <span
            key={lineNumber}
            data-testid={`syntax-line-${lineNumber}`}
            {...lineProps?.(lineNumber)}
          >
            {showLineNumbers && <span className="linenumber">{lineNumber}</span>}
            {line}
          </span>
        );
      })}
    </pre>
  ),
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
    appLayer: 'backend',
    appLayerSource: 'backend_path',
    functionalArea: 'resolution',
    functionalAreaSource: 'resolution_path',
    startLine: 0,
    endLine: 2,
    resolutionConfidence: graphHealth?.resolutionConfidence,
    resolutionGapCount: graphHealth?.resolutionGapCount,
    resolutionHealthBuckets: graphHealth?.resolutionHealthBuckets,
    graphHealth,
  },
});

const makeGapNode = (id: string): GraphNode => ({
  id,
  label: 'ResolutionGap',
  properties: {
    name: id,
    filePath: 'src/load.ts',
    appLayer: 'backend',
    functionalArea: 'resolution',
    factFamily: 'call',
    targetText: 'loadThing',
    targetRole: 'callable',
    classification: 'in_repo_unresolved',
    actionability: 'analyzer_gap',
    sourceSiteStatus: 'unresolved_local_binding',
    proofKind: 'none',
    count: 2,
  },
});

const makeGapRelationship = (sourceId: string, targetId: string): GraphRelationship => ({
  id: `${sourceId}->${targetId}`,
  sourceId,
  targetId,
  type: 'HAS_RESOLUTION_GAP',
  confidence: 1,
  reason: 'source-backed gap',
});

describe('CodeReferencesPanel Graph Health detail', () => {
  it('converts one-based exclusive graph ranges once for read, display, and highlight', async () => {
    readFileMock.mockResolvedValueOnce({
      content: ['line 9', 'line 10', 'line 11', 'line 12'].join('\n'),
      startLine: 8,
      endLine: 11,
      totalLines: 20,
    });
    const scrollIntoView = vi.fn();
    Object.defineProperty(HTMLElement.prototype, 'scrollIntoView', {
      configurable: true,
      value: scrollIntoView,
    });
    vi.stubGlobal(
      'requestAnimationFrame',
      ((callback: FrameRequestCallback) =>
        window.setTimeout(() => callback(performance.now()), 0)) as typeof requestAnimationFrame,
    );
    vi.stubGlobal(
      'cancelAnimationFrame',
      ((id: number) => window.clearTimeout(id)) as typeof cancelAnimationFrame,
    );

    const selected: GraphNode = {
      id: 'opaque-selected',
      label: 'Function',
      properties: {
        name: 'sharedName',
        filePath: 'src/load.ts',
        startLine: 10,
        startCol: 0,
        endLine: 12,
        endCol: 0,
      },
    };
    const graph = createKnowledgeGraph();
    graph.addNode(selected);

    render(
      <AppStateProvider>
        <StateCapture />
        <CodeReferencesPanel onFocusNode={vi.fn()} />
      </AppStateProvider>,
    );

    act(() => {
      appState!.setGraph(graph);
      appState!.setSelectedNode(selected);
      appState!.addCodeReference({
        filePath: 'src/load.ts',
        startLine: 10,
        startCol: 0,
        endLine: 12,
        endCol: 0,
        nodeId: selected.id,
        label: selected.label,
        name: selected.properties.name,
        source: 'ai',
      });
    });

    await screen.findByText('AI Citations');
    expect(document.body).toHaveTextContent('L10–11');
    expect(readFileMock).toHaveBeenLastCalledWith(
      'src/load.ts',
      expect.objectContaining({ startLine: 0, endLine: 60 }),
    );
    expect(screen.getByTestId('syntax-line-10')).toHaveStyle({
      backgroundColor: 'rgba(6, 182, 212, 0.14)',
    });
    expect(screen.getByTestId('syntax-line-11')).toHaveStyle({
      backgroundColor: 'rgba(6, 182, 212, 0.14)',
    });
    expect(screen.getByTestId('syntax-line-12').style.backgroundColor).not.toBe(
      'rgba(6, 182, 212, 0.14)',
    );
    expect(screen.getByTestId('syntax-line-12').style.borderLeft).toContain(
      'transparent',
    );
    await waitFor(() => expect(scrollIntoView).toHaveBeenCalled());
  });

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
      resolutionConfidence: 'degraded',
      resolutionGapCount: 2,
      resolutionHealthBuckets: { unresolved_call_target: 2, in_repo_analyzer_gap: 2 },
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
    graph.addNode(makeGapNode('gap-loadThing'));
    graph.addRelationship(makeGapRelationship(selected.id, 'gap-loadThing'));
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
    expect(detail).toHaveTextContent('App Layer Backend');
    expect(detail).toHaveTextContent('Functional Area Resolution');
    expect(detail).toHaveTextContent(/In\s+0/);
    expect(detail).toHaveTextContent(/Out\s+2/);
    expect(detail).toHaveTextContent(/Comp\s+2/);
    expect(detail).toHaveTextContent(/Resolution\s+Degraded/);
    expect(detail).toHaveTextContent(/Gaps\s+2/);
    expect(detail).toHaveTextContent('Unresolved Call Target 2');
    expect(detail).toHaveTextContent(
      'Resolution confidence: Resolution evidence is degraded by unresolved references; topology remains counted wiring evidence, not a dead-code verdict.',
    );
    expect(detail).toHaveTextContent('Structural 3');
    expect(detail).toHaveTextContent('Test');
    expect(detail).toHaveTextContent('Unresolved reference x2: loadThing');
    expect(detail).toHaveTextContent(
      'Call · target loadThing · role Callable · In-repo unresolved · Analyzer gap · status Unresolved Local Binding · proof None x2',
    );
    expect(detail).toHaveTextContent('Detached: no accepted root reaches this counted-edge component.');
    expect(detail).toHaveTextContent('Next: Review only if this expected-isolated overlay looks wrong for the node.');

    await userEvent.click(screen.getByRole('button', { name: 'Focus component' }));

    expect(onFocusNode).toHaveBeenCalledWith('comp-a');
    expect([...appState!.highlightedNodeIds].sort()).toEqual(['comp-a', 'comp-b']);
  });
});
