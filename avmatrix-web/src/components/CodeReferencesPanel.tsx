import { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import {
  Code,
  PanelLeftClose,
  PanelLeft,
  Trash2,
  X,
  Target,
  FileCode,
  MousePointerClick,
  Loader2,
} from '@/lib/lucide-icons';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism';
import { useAppState } from '../hooks/useAppState.local-runtime';
import {
  type GraphHealthDiagnostic,
  type GraphNode,
  getSyntaxLanguageFromFilename,
} from '@/generated/avmatrix-contracts';
import { NODE_COLORS } from '../lib/constants';
import {
  GRAPH_HEALTH_CONFIDENCE_DESCRIPTIONS,
  GRAPH_HEALTH_CONFIDENCE_LABELS,
  GRAPH_HEALTH_DIAGNOSTIC_DESCRIPTIONS,
  GRAPH_HEALTH_DIAGNOSTIC_LABELS,
  GRAPH_HEALTH_REASON_DESCRIPTIONS,
  GRAPH_HEALTH_REASON_LABELS,
  GRAPH_HEALTH_TOPOLOGY_DESCRIPTIONS,
  GRAPH_HEALTH_TOPOLOGY_LABELS,
  getGraphHealthNextAction,
  getNodeGraphHealth,
} from '../lib/graph-health-filters';
import { readFile, type ReadFileResult } from '../services/backend-client';

const getSyntaxLanguage = (filePath: string | undefined): string => {
  if (!filePath) return 'text';
  return getSyntaxLanguageFromFilename(filePath);
};

// Match the code theme used elsewhere in the app
const customTheme = {
  ...vscDarkPlus,
  'pre[class*="language-"]': {
    ...vscDarkPlus['pre[class*="language-"]'],
    background: '#0a0a10',
    margin: 0,
    padding: '12px 0',
    fontSize: '13px',
    lineHeight: '1.6',
  },
  'code[class*="language-"]': {
    ...vscDarkPlus['code[class*="language-"]'],
    background: 'transparent',
    fontFamily: '"JetBrains Mono", "Fira Code", monospace',
  },
};

const formatGraphHealthKey = (value: string): string =>
  value
    .split('_')
    .filter(Boolean)
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(' ');

const formatExcludedEdgeCounts = (
  counts: Record<string, number> | undefined,
): string => {
  const entries = Object.entries(counts ?? {}).filter(([, count]) => count > 0);
  if (entries.length === 0) return 'None';
  return entries
    .map(([category, count]) => `${formatGraphHealthKey(category)} ${count}`)
    .join(', ');
};

const formatDiagnosticDetail = (diagnostic: GraphHealthDiagnostic): string => {
  const count = diagnostic.count && diagnostic.count > 0 ? diagnostic.count : 1;
  const label = GRAPH_HEALTH_DIAGNOSTIC_LABELS[diagnostic.kind] ?? formatGraphHealthKey(diagnostic.kind);
  const detail = diagnostic.targetText ? `: ${diagnostic.targetText}` : '';
  return `${label} x${count}${detail}`;
};

export interface CodeReferencesPanelProps {
  onFocusNode: (nodeId: string) => void;
}

export const CodeReferencesPanel = ({ onFocusNode }: CodeReferencesPanelProps) => {
  const {
    graph,
    selectedNode,
    codeReferences,
    removeCodeReference,
    clearCodeReferences,
    setSelectedNode,
    setHighlightedNodeIds,
    codeReferenceFocus,
    projectName,
  } = useAppState();

  const nodeById = useMemo(() => {
    if (!graph) return new Map<string, GraphNode>();
    return new Map(graph.nodes.map((n) => [n.id, n]));
  }, [graph]);

  const [isCollapsed, setIsCollapsed] = useState(false);
  const [glowRefId, setGlowRefId] = useState<string | null>(null);
  const panelRef = useRef<HTMLElement | null>(null);
  const resizeRef = useRef<{ startX: number; startWidth: number } | null>(null);
  const refCardEls = useRef<Map<string, HTMLDivElement | null>>(new Map());
  const glowTimerRef = useRef<number | null>(null);

  useEffect(() => {
    return () => {
      if (glowTimerRef.current) {
        window.clearTimeout(glowTimerRef.current);
        glowTimerRef.current = null;
      }
    };
  }, []);

  const [panelWidth, setPanelWidth] = useState<number>(() => {
    try {
      const saved = window.localStorage.getItem('avmatrix.codePanelWidth');
      const parsed = saved ? parseInt(saved, 10) : NaN;
      if (!Number.isFinite(parsed)) return 560; // increased default
      return Math.max(420, Math.min(parsed, 900));
    } catch {
      return 560;
    }
  });

  useEffect(() => {
    try {
      window.localStorage.setItem('avmatrix.codePanelWidth', String(panelWidth));
    } catch {
      // ignore
    }
  }, [panelWidth]);

  const startResize = useCallback(
    (e: React.MouseEvent) => {
      e.preventDefault();
      e.stopPropagation();
      resizeRef.current = { startX: e.clientX, startWidth: panelWidth };
      document.body.style.cursor = 'col-resize';
      document.body.style.userSelect = 'none';

      const onMove = (ev: MouseEvent) => {
        const state = resizeRef.current;
        if (!state) return;
        const delta = ev.clientX - state.startX;
        const next = Math.max(420, Math.min(state.startWidth + delta, 900));
        setPanelWidth(next);
      };

      const onUp = () => {
        resizeRef.current = null;
        document.body.style.cursor = '';
        document.body.style.userSelect = '';
        window.removeEventListener('mousemove', onMove);
        window.removeEventListener('mouseup', onUp);
      };

      window.addEventListener('mousemove', onMove);
      window.addEventListener('mouseup', onUp);
    },
    [panelWidth],
  );

  const aiReferences = useMemo(
    () => codeReferences.filter((r) => r.source === 'ai'),
    [codeReferences],
  );

  // When the user clicks a citation badge in chat, focus the corresponding snippet card:
  // - expand the panel if collapsed
  // - smooth-scroll the card into view
  // - briefly glow it for discoverability
  useEffect(() => {
    if (!codeReferenceFocus) return;

    // Ensure panel is expanded
    setIsCollapsed(false);

    const { filePath, startLine, endLine } = codeReferenceFocus;
    const target =
      aiReferences.find(
        (r) => r.filePath === filePath && r.startLine === startLine && r.endLine === endLine,
      ) ?? aiReferences.find((r) => r.filePath === filePath);

    if (!target) return;

    // Double rAF: wait for collapse state + list DOM to render.
    const rafIds: number[] = [];
    const outerRafId = requestAnimationFrame(() => {
      const innerRafId = requestAnimationFrame(() => {
        const el = refCardEls.current.get(target.id);
        if (!el) return;

        el.scrollIntoView({ behavior: 'smooth', block: 'center' });
        setGlowRefId(target.id);

        if (glowTimerRef.current) {
          window.clearTimeout(glowTimerRef.current);
        }
        glowTimerRef.current = window.setTimeout(() => {
          setGlowRefId((prev) => (prev === target.id ? null : prev));
          glowTimerRef.current = null;
        }, 1200);
      });
      rafIds.push(innerRafId);
    });
    rafIds.push(outerRafId);

    return () => {
      rafIds.forEach((id) => cancelAnimationFrame(id));
    };
  }, [codeReferenceFocus, aiReferences]);

  const refsWithSnippets = useMemo(() => {
    return aiReferences.map((ref) => {
      return {
        ref,
        content: null as string | null,
        start: 0,
        end: 0,
        highlightStart: 0,
        highlightEnd: 0,
        totalLines: 0,
      };
    });
  }, [aiReferences]);

  const selectedFilePath = selectedNode?.properties?.filePath;
  const selectedIsFile = selectedNode?.label === 'File' && !!selectedFilePath;
  const showSelectedViewer = !!selectedNode && !!selectedFilePath;
  const showCitations = aiReferences.length > 0;
  const selectedGraphHealth = useMemo(
    () => (selectedNode ? getNodeGraphHealth(selectedNode) : null),
    [selectedNode],
  );
  const selectedComponentNodeIds = useMemo(() => {
    if (!graph || !selectedGraphHealth?.componentId) return [];
    return graph.nodes
      .filter((node) => getNodeGraphHealth(node)?.componentId === selectedGraphHealth.componentId)
      .map((node) => node.id);
  }, [graph, selectedGraphHealth?.componentId]);

  const handleFocusGraphHealthComponent = useCallback(() => {
    if (!selectedNode) return;
    const ids =
      selectedComponentNodeIds.length > 0 ? selectedComponentNodeIds : [selectedNode.id];
    setHighlightedNodeIds(new Set(ids));
    onFocusNode(selectedNode.id);
  }, [onFocusNode, selectedComponentNodeIds, selectedNode, setHighlightedNodeIds]);

  // Fetch file content from the server when a node with a filePath is selected.
  // For non-File nodes (functions, classes, etc.), fetch a buffer around the symbol
  // instead of the entire file. For File nodes, fetch the whole file.
  const CONTEXT_LINES = 50; // lines of context above and below the symbol

  const [fileResult, setFileResult] = useState<ReadFileResult | null>(null);
  const [isLoadingFile, setIsLoadingFile] = useState(false);
  const selectedViewerRef = useRef<HTMLDivElement>(null);

  const selectedFileContent = fileResult?.content;
  const fileStartLine = fileResult?.startLine ?? 0;

  useEffect(() => {
    if (!selectedFilePath) {
      setFileResult(null);
      return;
    }

    let cancelled = false;
    setIsLoadingFile(true);
    setFileResult(null);

    // Determine read range: full file for File nodes, buffered for symbols
    const startLine = selectedNode?.properties?.startLine as number | undefined;
    const endLine = selectedNode?.properties?.endLine as number | undefined;
    const isWholeFile = selectedIsFile || startLine === undefined;

    const options = isWholeFile
      ? { repo: projectName }
      : {
          startLine: Math.max(0, startLine - CONTEXT_LINES),
          endLine: (endLine ?? startLine) + CONTEXT_LINES,
          repo: projectName,
        };

    readFile(selectedFilePath, { ...options, repo: projectName || undefined })
      .then((result) => {
        if (!cancelled) {
          setFileResult(result);
          setIsLoadingFile(false);
        }
      })
      .catch(() => {
        if (!cancelled) {
          setFileResult(null);
          setIsLoadingFile(false);
        }
      });

    return () => {
      cancelled = true;
    };
  }, [
    selectedFilePath,
    selectedNode?.properties?.startLine,
    selectedNode?.properties?.endLine,
    selectedIsFile,
    projectName,
  ]);

  // Scroll to the selected node's startLine after content loads
  useEffect(() => {
    if (!selectedFileContent || !selectedNode?.properties?.startLine) return;
    const startLine = selectedNode.properties.startLine as number;

    // Double rAF: wait for SyntaxHighlighter to fully render before scrolling
    let cancelled = false;
    const outerRaf = requestAnimationFrame(() => {
      const innerRaf = requestAnimationFrame(() => {
        if (cancelled) return;
        const container = selectedViewerRef.current;
        if (!container) return;
        const lineEl =
          (container.querySelector(`[data-line-number="${startLine + 1}"]`) as HTMLElement) ??
          (container.querySelectorAll('.linenumber')[startLine] as HTMLElement);
        if (lineEl) {
          lineEl.scrollIntoView({ behavior: 'smooth', block: 'center' });
        } else {
          // Fallback: estimate scroll position based on line height
          const lineHeight = 20.8; // 13px font * 1.6 line-height
          container.scrollTop = Math.max(0, startLine * lineHeight - container.clientHeight / 3);
        }
      });
      rafIds.push(innerRaf);
    });
    const rafIds = [outerRaf];
    return () => {
      cancelled = true;
      rafIds.forEach((id) => cancelAnimationFrame(id));
    };
  }, [selectedFileContent, selectedNode?.properties?.startLine]);

  if (isCollapsed) {
    return (
      <aside className="workspace-shell flex h-full w-12 flex-shrink-0 flex-col items-center gap-2 border-r-[3px] border-workspace-border-default bg-workspace-surface py-3">
        <button
          onClick={() => setIsCollapsed(false)}
          className="workspace-outline-button rounded p-2 text-workspace-text-secondary hover:text-workspace-text-primary"
          title="Expand Code Panel"
        >
          <PanelLeft className="h-5 w-5" />
        </button>
        <div className="my-1 h-px w-6 bg-border-subtle" />
        {showSelectedViewer && (
          <div className="rotate-90 text-[9px] font-medium tracking-wide whitespace-nowrap text-workspace-text-primary">
            SELECTED
          </div>
        )}
        {showCitations && (
          <div className="mt-4 rotate-90 text-[9px] font-medium tracking-wide whitespace-nowrap text-workspace-text-primary">
            AI • {aiReferences.length}
          </div>
        )}
      </aside>
    );
  }

  return (
    <aside
      ref={(el) => {
        panelRef.current = el;
      }}
      className="workspace-shell relative flex h-full animate-slide-in flex-col border-r-[3px] border-workspace-border-default bg-workspace-surface/95 shadow-2xl backdrop-blur-md"
      style={{ width: panelWidth }}
    >
      <div
        onMouseDown={startResize}
        className="absolute top-0 right-0 h-full w-2 cursor-col-resize bg-transparent transition-colors hover:bg-workspace-border-strong/25"
        title="Drag to resize"
      />
      <div className="flex items-center justify-between border-b-[3px] border-workspace-border-default bg-workspace-base px-3 py-3">
        <div className="flex items-center gap-2">
          <Code className="h-4 w-4 text-workspace-text-primary" />
          <span className="font-mono text-sm font-semibold text-workspace-text-primary">
            Code Inspector
          </span>
        </div>
        <div className="flex items-center gap-1.5">
          {showCitations && (
            <button
              onClick={() => clearCodeReferences()}
              className="workspace-outline-button rounded border-error p-1.5 text-workspace-text-secondary hover:bg-error/10 hover:text-error"
              title="Clear AI citations"
            >
              <Trash2 className="h-4 w-4" />
            </button>
          )}
          <button
            onClick={() => setIsCollapsed(true)}
            className="workspace-outline-button rounded p-1.5 text-workspace-text-secondary hover:text-workspace-text-primary"
            title="Collapse Panel"
          >
            <PanelLeftClose className="h-4 w-4" />
          </button>
        </div>
      </div>

      <div className="flex min-h-0 flex-1 flex-col">
        {showSelectedViewer && (
          <div className={`${showCitations ? 'h-[42%]' : 'flex-1'} flex min-h-0 flex-col`}>
            <div className="flex items-center gap-2 border-b border-workspace-border-default bg-workspace-base px-3 py-2">
              <div className="flex items-center gap-1.5 rounded-md border-[2px] border-workspace-border-default bg-workspace-surface px-2 py-0.5">
                <MousePointerClick className="h-3 w-3 text-workspace-text-primary" />
                <span className="text-[10px] font-semibold tracking-wide text-workspace-text-primary uppercase">
                  Selected
                </span>
              </div>
              <FileCode className="ml-1 h-3.5 w-3.5 text-workspace-text-secondary" />
              <span className="flex-1 truncate font-mono text-xs text-workspace-text-primary">
                {selectedNode?.properties?.filePath?.split('/').pop() ??
                  selectedNode?.properties?.name}
              </span>
              <button
                onClick={() => setSelectedNode(null)}
                className="workspace-outline-button rounded p-1 text-workspace-text-secondary hover:text-workspace-text-primary"
                title="Clear selection"
              >
                <X className="h-4 w-4" />
              </button>
            </div>
            {selectedGraphHealth && (
              <div
                className="border-b border-workspace-border-default bg-workspace-surface px-3 py-3"
                data-testid="graph-health-node-detail"
              >
                <div className="mb-2 flex flex-wrap items-center gap-2">
                  <span
                    className="rounded border border-workspace-border-default bg-workspace-inset px-2 py-0.5 text-[10px] font-semibold tracking-wide text-workspace-text-primary uppercase"
                    title={
                      GRAPH_HEALTH_TOPOLOGY_DESCRIPTIONS[
                        selectedGraphHealth.topologyStatus
                      ]
                    }
                  >
                    {GRAPH_HEALTH_TOPOLOGY_LABELS[
                      selectedGraphHealth.topologyStatus
                    ] ?? formatGraphHealthKey(selectedGraphHealth.topologyStatus)}
                  </span>
                  <span
                    className="rounded border border-workspace-border-default bg-workspace-inset px-2 py-0.5 text-[10px] text-workspace-text-secondary"
                    title={
                      GRAPH_HEALTH_CONFIDENCE_DESCRIPTIONS[
                        selectedGraphHealth.confidence
                      ]
                    }
                  >
                    {GRAPH_HEALTH_CONFIDENCE_LABELS[
                      selectedGraphHealth.confidence
                    ] ?? formatGraphHealthKey(selectedGraphHealth.confidence)}
                  </span>
                  {selectedGraphHealth.topologyStatus === 'detached_component' &&
                    selectedGraphHealth.componentId && (
                      <button
                        onClick={handleFocusGraphHealthComponent}
                        className="workspace-outline-button rounded px-2 py-0.5 text-[10px] text-workspace-text-secondary hover:text-workspace-text-primary"
                        title="Highlight this component and focus the selected node in the graph"
                      >
                        Focus component
                      </button>
                    )}
                </div>

                <div className="grid grid-cols-3 gap-2 text-[11px] text-workspace-text-secondary">
                  <div title="Counted incoming relationships under the Graph Health edge policy">
                    In{' '}
                    <span className="font-mono text-workspace-text-primary">
                      {selectedGraphHealth.countedIncoming}
                    </span>
                  </div>
                  <div title="Counted outgoing relationships under the Graph Health edge policy">
                    Out{' '}
                    <span className="font-mono text-workspace-text-primary">
                      {selectedGraphHealth.countedOutgoing}
                    </span>
                  </div>
                  <div title="Counted-edge component size">
                    Comp{' '}
                    <span className="font-mono text-workspace-text-primary">
                      {selectedGraphHealth.componentSize ?? 1}
                    </span>
                  </div>
                </div>

                <div className="mt-2 space-y-1 text-[11px] text-workspace-text-secondary">
                  <div title="Excluded edge categories are not counted as topology wiring">
                    Excluded: {formatExcludedEdgeCounts(selectedGraphHealth.excludedEdgeCounts)}
                  </div>
                  <div>
                    Reasons:{' '}
                    {(selectedGraphHealth.expectedIsolationReasons ?? []).length > 0 ? (
                      <span className="inline-flex flex-wrap gap-1 align-middle">
                        {selectedGraphHealth.expectedIsolationReasons?.map((reason) => (
                          <span
                            key={reason}
                            className="rounded bg-workspace-inset px-1.5 py-0.5 text-[10px] text-workspace-text-primary"
                            title={GRAPH_HEALTH_REASON_DESCRIPTIONS[reason]}
                          >
                            {GRAPH_HEALTH_REASON_LABELS[reason] ?? formatGraphHealthKey(reason)}
                          </span>
                        ))}
                      </span>
                    ) : (
                      <span>None</span>
                    )}
                  </div>
                  <div>
                    Diagnostics:{' '}
                    {(selectedGraphHealth.diagnostics ?? []).length > 0 ? (
                      <span className="inline-flex flex-wrap gap-1 align-middle">
                        {selectedGraphHealth.diagnostics?.map((diagnostic, index) => (
                          <span
                            key={`${diagnostic.kind}-${index}`}
                            className="rounded bg-workspace-inset px-1.5 py-0.5 text-[10px] text-workspace-text-primary"
                            title={
                              GRAPH_HEALTH_DIAGNOSTIC_DESCRIPTIONS[diagnostic.kind] ??
                              'Graph Health diagnostic evidence.'
                            }
                          >
                            {formatDiagnosticDetail(diagnostic)}
                          </span>
                        ))}
                      </span>
                    ) : (
                      <span>None</span>
                    )}
                  </div>
                  {selectedGraphHealth.topologyStatus === 'detached_component' && (
                    <div title="Detached components have counted internal wiring but no accepted root reachability">
                      Detached: no accepted root reaches this counted-edge component.
                    </div>
                  )}
                  <div className="text-workspace-text-primary">
                    Next: {getGraphHealthNextAction(selectedGraphHealth)}
                  </div>
                </div>
              </div>
            )}
            <div ref={selectedViewerRef} className="scrollbar-thin min-h-0 flex-1 overflow-auto">
              {isLoadingFile ? (
                <div className="flex items-center justify-center gap-2 py-8 text-workspace-text-secondary">
                  <Loader2 className="h-4 w-4 animate-spin" />
                  <span className="text-sm">Loading source...</span>
                </div>
              ) : selectedFileContent ? (
                <SyntaxHighlighter
                  language={getSyntaxLanguage(selectedFilePath)}
                  style={customTheme as any}
                  showLineNumbers
                  startingLineNumber={fileStartLine + 1}
                  lineNumberStyle={{
                    minWidth: '3em',
                    paddingRight: '1em',
                    color: '#5a5a70',
                    textAlign: 'right',
                    userSelect: 'none',
                  }}
                  lineProps={(lineNumber) => {
                    const symStart = selectedNode?.properties?.startLine;
                    const symEnd = selectedNode?.properties?.endLine ?? symStart;
                    const isHighlighted =
                      typeof symStart === 'number' &&
                      lineNumber >= symStart + 1 &&
                      lineNumber <= (symEnd ?? symStart) + 1;
                    return {
                      style: {
                        display: 'block',
                        backgroundColor: isHighlighted ? 'rgba(6, 182, 212, 0.14)' : 'transparent',
                        borderLeft: isHighlighted ? '3px solid #9a7e63' : '3px solid transparent',
                        paddingLeft: '12px',
                        paddingRight: '16px',
                      },
                    };
                  }}
                  wrapLines
                >
                  {selectedFileContent}
                </SyntaxHighlighter>
              ) : (
                <div className="px-3 py-3 text-sm text-workspace-text-secondary">
                  {selectedIsFile ? (
                    <>
                      Code not available in memory for{' '}
                      <span className="font-mono">{selectedFilePath}</span>
                    </>
                  ) : (
                    <>Select a file node to preview its contents.</>
                  )}
                </div>
              )}
            </div>
          </div>
        )}

        {/* Divider between Selected viewer and AI refs (more visible) */}
        {showSelectedViewer && showCitations && (
          <div className="h-1.5 bg-gradient-to-r from-transparent via-workspace-border-default to-transparent" />
        )}

        {showCitations && (
          <div className="flex min-h-0 flex-1 flex-col">
            <div className="flex items-center gap-2 border-b border-workspace-border-default bg-workspace-base px-3 py-2">
              <div className="flex items-center rounded-md border-[2px] border-workspace-border-default bg-workspace-surface px-2 py-0.5">
                <span className="text-[10px] font-semibold tracking-wide text-workspace-text-primary uppercase">
                  AI Citations
                </span>
              </div>
              <span className="ml-1 text-xs text-workspace-text-secondary">
                {aiReferences.length} reference{aiReferences.length !== 1 ? 's' : ''}
              </span>
            </div>
            <div className="scrollbar-thin min-h-0 flex-1 space-y-3 overflow-y-auto p-3">
              {refsWithSnippets.map(
                ({ ref, content, start, highlightStart, highlightEnd, totalLines }) => {
                  const nodeColor = ref.label
                    ? (NODE_COLORS as any)[ref.label] || '#6b7280'
                    : '#6b7280';
                  const hasRange = typeof ref.startLine === 'number';
                  const startDisplay = hasRange ? (ref.startLine ?? 0) + 1 : undefined;
                  const endDisplay = hasRange ? (ref.endLine ?? ref.startLine ?? 0) + 1 : undefined;
                  const language = getSyntaxLanguage(ref.filePath);

                  const isGlowing = glowRefId === ref.id;

                  return (
                    <div
                      key={ref.id}
                      ref={(el) => {
                        refCardEls.current.set(ref.id, el);
                      }}
                      className={[
                        'overflow-hidden rounded-xl border-[2px] border-workspace-border-default bg-workspace-inset transition-all',
                        isGlowing
                          ? 'animate-pulse shadow-[0_0_0_6px_rgba(154,126,99,0.14)] ring-2 ring-workspace-border-strong/70'
                          : '',
                      ].join(' ')}
                    >
                      <div className="flex items-start gap-2 border-b border-workspace-border-default bg-workspace-base px-3 py-2">
                        <span
                          className="mt-0.5 flex-shrink-0 rounded px-2 py-0.5 text-[10px] font-semibold tracking-wide uppercase"
                          style={{ backgroundColor: nodeColor, color: '#06060a' }}
                          title={ref.label ?? 'Code'}
                        >
                          {ref.label ?? 'Code'}
                        </span>
                        <div className="min-w-0 flex-1">
                          <div className="truncate text-xs font-medium text-workspace-text-primary">
                            {ref.name ?? ref.filePath.split('/').pop() ?? ref.filePath}
                          </div>
                          <div className="truncate font-mono text-[11px] text-workspace-text-secondary">
                            {ref.filePath}
                            {startDisplay !== undefined && (
                              <span className="text-workspace-text-secondary">
                                {' '}
                                • L{startDisplay}
                                {endDisplay !== startDisplay ? `–${endDisplay}` : ''}
                              </span>
                            )}
                            {totalLines > 0 && (
                              <span className="text-workspace-text-muted">
                                {' '}
                                • {totalLines} lines
                              </span>
                            )}
                          </div>
                        </div>
                        <div className="flex items-center gap-1">
                          {ref.nodeId && (
                            <button
                              onClick={() => {
                                const nodeId = ref.nodeId!;
                                // Sync selection + focus graph
                                if (graph) {
                                  const node = nodeById.get(nodeId);
                                  if (node) setSelectedNode(node);
                                }
                                onFocusNode(nodeId);
                              }}
                              className="workspace-outline-button rounded p-1.5 text-workspace-text-secondary hover:text-workspace-text-primary"
                              title="Focus in graph"
                            >
                              <Target className="h-4 w-4" />
                            </button>
                          )}
                          <button
                            onClick={() => removeCodeReference(ref.id)}
                            className="workspace-outline-button rounded p-1.5 text-workspace-text-secondary hover:text-workspace-text-primary"
                            title="Remove"
                          >
                            <X className="h-4 w-4" />
                          </button>
                        </div>
                      </div>

                      <div className="overflow-x-auto">
                        {content ? (
                          <SyntaxHighlighter
                            language={language}
                            style={customTheme as any}
                            showLineNumbers
                            startingLineNumber={start + 1}
                            lineNumberStyle={{
                              minWidth: '3em',
                              paddingRight: '1em',
                              color: '#5a5a70',
                              textAlign: 'right',
                              userSelect: 'none',
                            }}
                            lineProps={(lineNumber) => {
                              const isHighlighted =
                                hasRange &&
                                lineNumber >= start + highlightStart + 1 &&
                                lineNumber <= start + highlightEnd + 1;
                              return {
                                style: {
                                  display: 'block',
                                  backgroundColor: isHighlighted
                                    ? 'rgba(6, 182, 212, 0.14)'
                                    : 'transparent',
                                  borderLeft: isHighlighted
                                    ? '3px solid #9a7e63'
                                    : '3px solid transparent',
                                  paddingLeft: '12px',
                                  paddingRight: '16px',
                                },
                              };
                            }}
                            wrapLines
                          >
                            {content}
                          </SyntaxHighlighter>
                        ) : (
                          <div className="px-3 py-3 text-sm text-workspace-text-secondary">
                            Code not available in memory for{' '}
                            <span className="font-mono">{ref.filePath}</span>
                          </div>
                        )}
                      </div>
                    </div>
                  );
                },
              )}
            </div>
          </div>
        )}
      </div>
    </aside>
  );
};
