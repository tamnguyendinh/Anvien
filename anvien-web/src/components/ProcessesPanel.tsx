/**
 * Processes Panel
 *
 * Lists all detected processes grouped by type (cross-community / intra-community).
 * Clicking a process opens the ProcessFlowModal with a flowchart.
 */

import { useState, useMemo, useCallback, useEffect } from 'react';
import {
  GitBranch,
  Search,
  Eye,
  Zap,
  Home,
  ChevronDown,
  ChevronRight,
  Lightbulb,
  Layers,
} from 'lucide-react';
import { useAppState } from '../hooks/useAppState.local-runtime';
import { ProcessFlowModal } from './ProcessFlowModal';
import type { ProcessData, ProcessStep } from '../lib/mermaid-generator';
import type { KnowledgeGraph } from '../core/graph/types';

const stringProp = (value: unknown): string | undefined =>
  typeof value === 'string' && value.trim() ? value : undefined;

const stringArrayProp = (value: unknown): string[] =>
  Array.isArray(value) ? value.filter((item): item is string => typeof item === 'string') : [];

const graphNodeLookup = (graph: KnowledgeGraph): Map<string, KnowledgeGraph['nodes'][number]> =>
  new Map(graph.nodes.map((node) => [node.id, node]));

const processStepsFromGraph = (
  graph: KnowledgeGraph,
  processId: string,
  nodeById: Map<string, KnowledgeGraph['nodes'][number]> = graphNodeLookup(graph),
): ProcessStep[] =>
  graph.relationships
    .filter((rel) => rel.type === 'STEP_IN_PROCESS' && rel.targetId === processId)
    .map((rel) => {
      const node = nodeById.get(rel.sourceId);
      return {
        id: rel.sourceId,
        name:
          stringProp(node?.properties.name) ??
          stringProp(node?.properties.label) ??
          stringProp(node?.properties.heuristicLabel) ??
          rel.sourceId,
        filePath: stringProp(node?.properties.filePath),
        stepNumber: typeof rel.step === 'number' ? rel.step : 0,
      };
    })
    .sort((a, b) => {
      if (a.stepNumber !== b.stepNumber) return a.stepNumber - b.stepNumber;
      return a.name.localeCompare(b.name);
    });

const callEdgesFromGraph = (
  graph: KnowledgeGraph,
  stepIds: Iterable<string>,
): Array<{ from: string; to: string; type: string }> => {
  const stepSet = new Set(stepIds);
  if (stepSet.size === 0) return [];

  return graph.relationships
    .filter(
      (rel) =>
        rel.type === 'CALLS' &&
        stepSet.has(rel.sourceId) &&
        stepSet.has(rel.targetId) &&
        rel.sourceId !== rel.targetId,
    )
    .map((rel) => ({ from: rel.sourceId, to: rel.targetId, type: rel.type }));
};

export const ProcessesPanel = () => {
  const { graph, setHighlightedNodeIds, highlightedNodeIds } = useAppState();
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedProcess, setSelectedProcess] = useState<ProcessData | null>(null);
  const [expandedSections, setExpandedSections] = useState<Set<string>>(
    new Set(['cross', 'intra']),
  );
  const [viewLoadingProcess, setViewLoadingProcess] = useState<string | null>(null);
  const [focusedProcessId, setFocusedProcessId] = useState<string | null>(null);

  // Extract processes from graph
  const processes = useMemo(() => {
    if (!graph) return { cross: [], intra: [] };

    const processNodes = graph.nodes.filter((n) => n.label === 'Process');

    const cross: Array<{ id: string; label: string; stepCount: number; clusters: string[] }> = [];
    const intra: Array<{ id: string; label: string; stepCount: number; clusters: string[] }> = [];

    for (const node of processNodes) {
      const item = {
        id: node.id,
        label: node.properties.heuristicLabel || node.properties.name || node.id,
        stepCount: node.properties.stepCount || 0,
        clusters: node.properties.communities || [],
      };

      if (node.properties.processType === 'cross_community') {
        cross.push(item);
      } else {
        intra.push(item);
      }
    }

    // Sort by step count (most complex first)
    cross.sort((a, b) => b.stepCount - a.stepCount);
    intra.sort((a, b) => b.stepCount - a.stepCount);

    return { cross, intra };
  }, [graph]);

  // Filter by search
  const filteredProcesses = useMemo(() => {
    if (!searchQuery.trim()) return processes;

    const query = searchQuery.toLowerCase();
    return {
      cross: processes.cross.filter((p) => p.label.toLowerCase().includes(query)),
      intra: processes.intra.filter((p) => p.label.toLowerCase().includes(query)),
    };
  }, [processes, searchQuery]);

  // Toggle section expansion
  const toggleSection = useCallback((section: string) => {
    setExpandedSections((prev) => {
      const next = new Set(prev);
      if (next.has(section)) {
        next.delete(section);
      } else {
        next.add(section);
      }
      return next;
    });
  }, []);

  // Load ALL processes and combine into one mega-diagram
  const handleViewAllProcesses = useCallback(() => {
    if (!graph) return;
    setViewLoadingProcess('all');

    try {
      const allProcessIds = [...processes.cross, ...processes.intra].map((p) => p.id);

      if (allProcessIds.length === 0) return;

      // Collect all steps from all processes
      const allStepsMap = new Map<string, ProcessStep>();
      const nodeById = graphNodeLookup(graph);
      for (const processId of allProcessIds) {
        for (const step of processStepsFromGraph(graph, processId, nodeById)) {
          if (!allStepsMap.has(step.id)) {
            allStepsMap.set(step.id, step);
          }
        }
      }

      const allSteps = Array.from(allStepsMap.values());
      const allEdges = callEdgesFromGraph(
        graph,
        allSteps.map((step) => step.id),
      );

      const combinedProcessData: ProcessData = {
        id: 'combined-all',
        label: `All Processes (${allProcessIds.length} combined)`,
        processType: 'cross_community', // Treat as cross-community for styling
        steps: allSteps,
        edges: allEdges,
        clusters: [],
      };

      setSelectedProcess(combinedProcessData);
    } catch (error) {
      console.error('Failed to load combined processes:', error);
    } finally {
      setViewLoadingProcess(null);
    }
  }, [graph, processes]);

  // Load process steps and open modal
  const handleViewProcess = useCallback(
    (processId: string, label: string, processType: string) => {
      if (!graph) return;
      setViewLoadingProcess(processId);

      try {
        const steps = processStepsFromGraph(graph, processId);
        const edges = callEdgesFromGraph(
          graph,
          steps.map((step) => step.id),
        );

        // Get clusters for this process
        const processNode = graph?.nodes.find((n) => n.id === processId);
        const clusters = stringArrayProp(processNode?.properties.communities);

        const processData: ProcessData = {
          id: processId,
          label,
          processType: processType as 'cross_community' | 'intra_community',
          steps,
          edges,
          clusters,
        };

        setSelectedProcess(processData);
      } catch (error) {
        console.error('Failed to load process steps:', error);
      } finally {
        setViewLoadingProcess(null);
      }
    },
    [graph],
  );

  // Cache for process steps (so we don't re-query when toggling focus)
  const [processStepsCache, setProcessStepsCache] = useState<Map<string, string[]>>(new Map());

  // Toggle focus for any process - loads steps on demand
  const handleToggleFocusForProcess = useCallback(
    (processId: string) => {
      // If already focused on this process, turn off
      if (focusedProcessId === processId) {
        setHighlightedNodeIds(new Set());
        setFocusedProcessId(null);
        return;
      }

      // Check if we have cached steps
      if (processStepsCache.has(processId)) {
        const stepIds = processStepsCache.get(processId)!;
        setHighlightedNodeIds(new Set(stepIds));
        setFocusedProcessId(processId);
        return;
      }

      const stepIds = graph ? processStepsFromGraph(graph, processId).map((step) => step.id) : [];
      setProcessStepsCache((prev) => new Map(prev).set(processId, stepIds));
      setHighlightedNodeIds(new Set(stepIds));
      setFocusedProcessId(processId);
    },
    [focusedProcessId, graph, processStepsCache, setHighlightedNodeIds],
  );

  // Focus in graph callback - toggles highlight (used by modal)
  const handleFocusInGraph = useCallback(
    (nodeIds: string[], processId: string) => {
      // Check if this process is already focused
      if (focusedProcessId === processId) {
        // Clear focus
        setHighlightedNodeIds(new Set());
        setFocusedProcessId(null);
      } else {
        // Set focus and cache
        setHighlightedNodeIds(new Set(nodeIds));
        setFocusedProcessId(processId);
        setProcessStepsCache((prev) => new Map(prev).set(processId, nodeIds));
      }
    },
    [focusedProcessId, setHighlightedNodeIds],
  );

  // Clear focused process when highlights are cleared externally
  useEffect(() => {
    if (highlightedNodeIds.size === 0 && focusedProcessId !== null) {
      setFocusedProcessId(null);
    }
  }, [highlightedNodeIds, focusedProcessId]);

  const totalCount = processes.cross.length + processes.intra.length;

  if (totalCount === 0) {
    return (
      <div className="flex h-full flex-col items-center justify-center p-6 text-center">
        <div className="mb-4 flex h-14 w-14 items-center justify-center rounded-xl border-[3px] border-border-default bg-base">
          <GitBranch className="h-7 w-7 text-text-muted" />
        </div>
        <p className="press-eyebrow mb-2">Process desk</p>
        <h3 className="press-title mb-2 text-2xl">No Processes Detected</h3>
        <p className="press-reading max-w-xs text-center text-text-secondary">
          Processes are execution flows traced from entry points. Load a codebase to see detected
          processes.
        </p>
      </div>
    );
  }

  return (
    <div className="flex h-full flex-col">
      <div className="border-b-[3px] border-border-default bg-base p-3">
        <div className="mb-2 flex items-center gap-2">
          <div className="press-inset flex flex-1 items-center gap-2 px-3 py-2 focus-within:border-border-strong">
            <Search className="h-4 w-4 text-text-muted" />
            <input
              type="text"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="Filter processes..."
              className="flex-1 border-none bg-transparent font-mono text-sm text-text-primary outline-none placeholder:text-text-muted"
            />
          </div>
        </div>
        <div
          className="press-eyebrow flex items-center gap-2 tracking-normal text-text-secondary normal-case"
          data-testid="process-list-loaded"
        >
          <span>{totalCount} processes detected</span>
        </div>
      </div>

      {/* Process list */}
      <div className="scrollbar-thin flex-1 overflow-y-auto">
        <div className="px-4 py-3">
          <button
            onClick={handleViewAllProcesses}
            disabled={viewLoadingProcess !== null}
            className="press-panel group flex w-full items-center gap-3 p-4 text-left transition-all hover:border-border-strong"
          >
            <div className="rounded-lg border-[2px] border-border-default bg-base p-2 transition-colors group-hover:border-border-strong">
              <Layers className="h-5 w-5 text-border-strong" />
            </div>
            <div className="flex-1">
              <h4 className="font-mono text-sm font-medium text-text-primary">Full Process Map</h4>
              <p className="font-reading text-xs text-text-secondary">
                View combined map of {totalCount} processes
              </p>
            </div>
            {viewLoadingProcess === 'all' ? (
              <span className="mr-1 font-mono text-[11px] text-border-strong">Loading...</span>
            ) : (
              <Eye className="h-4 w-4 text-text-muted group-hover:text-border-strong" />
            )}
          </button>
        </div>

        {/* Cross-Community Section */}
        {filteredProcesses.cross.length > 0 && (
          <div className="border-b border-border-subtle">
            <button
              onClick={() => toggleSection('cross')}
              className="flex w-full items-center gap-2 px-4 py-3 text-left transition-colors hover:bg-base"
            >
              {expandedSections.has('cross') ? (
                <ChevronDown className="h-4 w-4 text-text-muted" />
              ) : (
                <ChevronRight className="h-4 w-4 text-text-muted" />
              )}
              <Zap className="h-4 w-4 text-warning" />
              <span className="font-mono text-sm font-medium text-text-primary">
                Cross-Community
              </span>
              <span className="press-badge ml-auto border-border-default bg-base px-2 py-0.5 text-xs tracking-normal text-text-secondary normal-case">
                {filteredProcesses.cross.length}
              </span>
            </button>

            {expandedSections.has('cross') && (
              <div className="pb-2">
                {filteredProcesses.cross.map((process) => (
                  <ProcessItem
                    key={process.id}
                    process={process}
                    isLoading={viewLoadingProcess === process.id}
                    isSelected={selectedProcess?.id === process.id}
                    isFocused={focusedProcessId === process.id}
                    onView={() => handleViewProcess(process.id, process.label, 'cross_community')}
                    onToggleFocus={() => handleToggleFocusForProcess(process.id)}
                  />
                ))}
              </div>
            )}
          </div>
        )}

        {/* Intra-Community Section */}
        {filteredProcesses.intra.length > 0 && (
          <div>
            <button
              onClick={() => toggleSection('intra')}
              className="flex w-full items-center gap-2 px-4 py-3 text-left transition-colors hover:bg-base"
            >
              {expandedSections.has('intra') ? (
                <ChevronDown className="h-4 w-4 text-text-muted" />
              ) : (
                <ChevronRight className="h-4 w-4 text-text-muted" />
              )}
              <Home className="h-4 w-4 text-success" />
              <span className="font-mono text-sm font-medium text-text-primary">
                Intra-Community
              </span>
              <span className="press-badge ml-auto border-border-default bg-base px-2 py-0.5 text-xs tracking-normal text-text-secondary normal-case">
                {filteredProcesses.intra.length}
              </span>
            </button>

            {expandedSections.has('intra') && (
              <div className="pb-2">
                {filteredProcesses.intra.map((process) => (
                  <ProcessItem
                    key={process.id}
                    process={process}
                    isLoading={viewLoadingProcess === process.id}
                    isSelected={selectedProcess?.id === process.id}
                    isFocused={focusedProcessId === process.id}
                    onView={() => handleViewProcess(process.id, process.label, 'intra_community')}
                    onToggleFocus={() => handleToggleFocusForProcess(process.id)}
                  />
                ))}
              </div>
            )}
          </div>
        )}
      </div>

      {/* Modal */}
      <ProcessFlowModal
        process={selectedProcess}
        onClose={() => setSelectedProcess(null)}
        onFocusInGraph={handleFocusInGraph}
        isFullScreen={selectedProcess?.id === 'combined-all'}
      />
    </div>
  );
};

// Individual process item
interface ProcessItemProps {
  process: { id: string; label: string; stepCount: number; clusters: string[] };
  isLoading: boolean;
  isSelected: boolean;
  isFocused: boolean;
  onView: () => void;
  onToggleFocus: () => void;
}

const ProcessItem = ({
  process,
  isLoading,
  isSelected,
  isFocused,
  onView,
  onToggleFocus,
}: ProcessItemProps) => {
  // Determine row styling - focused gets special highlight
  const rowClass = isFocused
    ? 'bg-base border-[2px] border-border-strong'
    : isSelected
      ? 'bg-base border-[2px] border-border-default'
      : '';

  return (
    <div
      data-testid="process-row"
      className={`group mx-2 flex items-center gap-2 rounded-lg px-4 py-3 transition-all hover:bg-base ${rowClass}`}
    >
      <GitBranch className="h-4 w-4 flex-shrink-0 text-text-muted" />
      <div className="min-w-0 flex-1">
        <div className="truncate font-mono text-sm text-text-primary">{process.label}</div>
        <div className="flex items-center gap-2 font-mono text-xs text-text-secondary">
          <span>{process.stepCount} steps</span>
          {process.clusters.length > 0 && (
            <>
              <span>•</span>
              <span>{process.clusters.length} clusters</span>
            </>
          )}
        </div>
      </div>
      {/* Lightbulb icon - appears on hover, always visible when focused */}
      <button
        onClick={onToggleFocus}
        className={`rounded-md p-1.5 transition-all ${
          isFocused
            ? 'border-[2px] border-border-strong bg-base text-border-strong opacity-100'
            : 'border-[2px] border-border-default bg-surface text-text-muted opacity-0 group-hover:opacity-100 hover:border-border-strong hover:text-border-strong'
        }`}
        title={isFocused ? 'Click to remove highlight from graph' : 'Click to highlight in graph'}
        data-testid="process-highlight-button"
      >
        <Lightbulb className="h-4 w-4" />
      </button>
      <button
        onClick={onView}
        disabled={isLoading}
        data-testid="process-view-button"
        className={`flex items-center gap-1.5 rounded-md border-[2px] px-2.5 py-1.5 font-mono text-xs font-medium transition-all disabled:opacity-50 ${
          isSelected
            ? 'border-border-strong bg-accent text-text-inverse opacity-100'
            : 'border-border-default bg-base text-text-secondary opacity-0 group-hover:opacity-100 hover:border-border-strong hover:text-text-primary'
        }`}
      >
        {isLoading ? (
          <span className="animate-pulse">Loading...</span>
        ) : isSelected ? (
          <>
            <Eye className="h-3.5 w-3.5" />
            Viewing
          </>
        ) : (
          <>
            <Eye className="h-3.5 w-3.5" />
            View
          </>
        )}
      </button>
    </div>
  );
};
