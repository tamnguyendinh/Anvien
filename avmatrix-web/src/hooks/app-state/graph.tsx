import { createContext, useContext, useCallback, useMemo, useState, ReactNode } from 'react';
import type {
  GraphHealthExpectedIsolationReason,
  GraphHealthTopologyStatus,
  GraphNode,
} from '@/generated/avmatrix-contracts';
import type { KnowledgeGraph } from '../../core/graph/types';
import { DEFAULT_VISIBLE_LABELS, DEFAULT_VISIBLE_EDGES, type EdgeType } from '../../lib/constants';
import {
  DEFAULT_GRAPH_HEALTH_FILTERS,
  type GraphHealthDiagnosticKind,
  type GraphHealthFilterState,
} from '../../lib/graph-health-filters';
import {
  readGraphLinksVisibilityPreference,
  writeGraphLinksVisibilityPreference,
} from '../../lib/graph-links-visibility';

interface GraphStateContextValue {
  graph: KnowledgeGraph | null;
  setGraph: (graph: KnowledgeGraph | null) => void;
  selectedNode: GraphNode | null;
  setSelectedNode: (node: GraphNode | null) => void;
  visibleLabels: string[];
  toggleLabelVisibility: (label: string) => void;
  visibleEdgeTypes: EdgeType[];
  toggleEdgeVisibility: (edgeType: EdgeType) => void;
  areGraphLinksVisible: boolean;
  setGraphLinksVisible: (visible: boolean) => void;
  toggleGraphLinksVisible: () => void;
  graphHealthFilters: GraphHealthFilterState;
  toggleGraphHealthTopologyStatus: (status: GraphHealthTopologyStatus) => void;
  toggleGraphHealthExpectedReason: (reason: GraphHealthExpectedIsolationReason) => void;
  toggleGraphHealthDiagnosticKind: (kind: GraphHealthDiagnosticKind) => void;
  resetGraphHealthFilters: () => void;
  depthFilter: number | null;
  setDepthFilter: (depth: number | null) => void;
  highlightedNodeIds: Set<string>;
  setHighlightedNodeIds: (ids: Set<string>) => void;
}

const GraphStateContext = createContext<GraphStateContextValue | null>(null);

export const GraphStateProvider = ({ children }: { children: ReactNode }) => {
  const [graph, setGraph] = useState<KnowledgeGraph | null>(null);
  const [selectedNode, setSelectedNode] = useState<GraphNode | null>(null);
  const [visibleLabels, setVisibleLabels] = useState<string[]>(DEFAULT_VISIBLE_LABELS);
  const [visibleEdgeTypes, setVisibleEdgeTypes] = useState<EdgeType[]>(DEFAULT_VISIBLE_EDGES);
  const [graphHealthFilters, setGraphHealthFilters] = useState<GraphHealthFilterState>(
    DEFAULT_GRAPH_HEALTH_FILTERS,
  );
  const [areGraphLinksVisible, setGraphLinksVisibleState] = useState<boolean>(() =>
    readGraphLinksVisibilityPreference(),
  );
  const [depthFilter, setDepthFilter] = useState<number | null>(null);
  const [highlightedNodeIds, setHighlightedNodeIds] = useState<Set<string>>(new Set());

  const toggleLabelVisibility = useCallback((label: string) => {
    setVisibleLabels((prev) =>
      prev.includes(label) ? prev.filter((l) => l !== label) : [...prev, label],
    );
  }, []);

  const toggleEdgeVisibility = useCallback((edgeType: EdgeType) => {
    setVisibleEdgeTypes((prev) =>
      prev.includes(edgeType) ? prev.filter((e) => e !== edgeType) : [...prev, edgeType],
    );
  }, []);

  const toggleGraphHealthTopologyStatus = useCallback((status: GraphHealthTopologyStatus) => {
    setGraphHealthFilters((prev) => ({
      ...prev,
      visibleTopologyStatuses: prev.visibleTopologyStatuses.includes(status)
        ? prev.visibleTopologyStatuses.filter((item) => item !== status)
        : [...prev.visibleTopologyStatuses, status],
    }));
  }, []);

  const toggleGraphHealthExpectedReason = useCallback((reason: GraphHealthExpectedIsolationReason) => {
    setGraphHealthFilters((prev) => ({
      ...prev,
      hiddenExpectedIsolationReasons: prev.hiddenExpectedIsolationReasons.includes(reason)
        ? prev.hiddenExpectedIsolationReasons.filter((item) => item !== reason)
        : [...prev.hiddenExpectedIsolationReasons, reason],
    }));
  }, []);

  const toggleGraphHealthDiagnosticKind = useCallback((kind: GraphHealthDiagnosticKind) => {
    setGraphHealthFilters((prev) => ({
      ...prev,
      visibleDiagnosticKinds: prev.visibleDiagnosticKinds.includes(kind)
        ? prev.visibleDiagnosticKinds.filter((item) => item !== kind)
        : [...prev.visibleDiagnosticKinds, kind],
    }));
  }, []);

  const resetGraphHealthFilters = useCallback(() => {
    setGraphHealthFilters(DEFAULT_GRAPH_HEALTH_FILTERS);
  }, []);

  const setGraphLinksVisible = useCallback((visible: boolean) => {
    setGraphLinksVisibleState(visible);
    writeGraphLinksVisibilityPreference(visible);
  }, []);

  const toggleGraphLinksVisible = useCallback(() => {
    setGraphLinksVisibleState((prev) => {
      const next = !prev;
      writeGraphLinksVisibilityPreference(next);
      return next;
    });
  }, []);

  const value = useMemo<GraphStateContextValue>(
    () => ({
      graph,
      setGraph,
      selectedNode,
      setSelectedNode,
      visibleLabels,
      toggleLabelVisibility,
      visibleEdgeTypes,
      toggleEdgeVisibility,
      areGraphLinksVisible,
      setGraphLinksVisible,
      toggleGraphLinksVisible,
      graphHealthFilters,
      toggleGraphHealthTopologyStatus,
      toggleGraphHealthExpectedReason,
      toggleGraphHealthDiagnosticKind,
      resetGraphHealthFilters,
      depthFilter,
      setDepthFilter,
      highlightedNodeIds,
      setHighlightedNodeIds,
    }),
    [
      graph,
      selectedNode,
      visibleLabels,
      visibleEdgeTypes,
      graphHealthFilters,
      areGraphLinksVisible,
      setGraphLinksVisible,
      toggleGraphLinksVisible,
      depthFilter,
      highlightedNodeIds,
    ],
  );

  return <GraphStateContext.Provider value={value}>{children}</GraphStateContext.Provider>;
};

export const useGraphState = (): GraphStateContextValue => {
  const ctx = useContext(GraphStateContext);
  if (!ctx) {
    throw new Error('useGraphState must be used within a GraphStateProvider');
  }
  return ctx;
};
