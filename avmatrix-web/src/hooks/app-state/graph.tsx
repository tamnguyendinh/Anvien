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
  DEFAULT_SEMANTIC_FILTERS,
  type SemanticFilterState,
} from '../../lib/semantic-filters';
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
  semanticFilters: SemanticFilterState;
  toggleSemanticAppLayer: (layer: SemanticFilterState['visibleAppLayers'][number]) => void;
  toggleSemanticMissingAppLayer: () => void;
  toggleResolutionConfidence: (
    confidence: SemanticFilterState['visibleResolutionConfidences'][number],
  ) => void;
  toggleResolutionHealthBucket: (
    bucket: SemanticFilterState['visibleResolutionHealthBuckets'][number],
  ) => void;
  toggleResolutionGapFactFamily: (family: string) => void;
  toggleResolutionGapTargetRole: (role: string) => void;
  toggleResolutionGapClassification: (
    classification: SemanticFilterState['visibleResolutionGapClassifications'][number],
  ) => void;
  toggleResolutionGapActionability: (
    actionability: SemanticFilterState['visibleResolutionGapActionabilities'][number],
  ) => void;
  toggleResolutionGapSourceAppLayer: (
    layer: SemanticFilterState['visibleResolutionGapSourceAppLayers'][number],
  ) => void;
  toggleResolutionGapTargetText: (targetText: string) => void;
  resetSemanticFilters: () => void;
  depthFilter: number | null;
  setDepthFilter: (depth: number | null) => void;
  highlightedNodeIds: Set<string>;
  setHighlightedNodeIds: (ids: Set<string>) => void;
}

const GraphStateContext = createContext<GraphStateContextValue | null>(null);

const toggleArrayValue = <T extends string>(items: T[], value: T): T[] =>
  items.includes(value) ? items.filter((item) => item !== value) : [...items, value];

export const GraphStateProvider = ({ children }: { children: ReactNode }) => {
  const [graph, setGraph] = useState<KnowledgeGraph | null>(null);
  const [selectedNode, setSelectedNode] = useState<GraphNode | null>(null);
  const [visibleLabels, setVisibleLabels] = useState<string[]>(DEFAULT_VISIBLE_LABELS);
  const [visibleEdgeTypes, setVisibleEdgeTypes] = useState<EdgeType[]>(DEFAULT_VISIBLE_EDGES);
  const [graphHealthFilters, setGraphHealthFilters] = useState<GraphHealthFilterState>(
    DEFAULT_GRAPH_HEALTH_FILTERS,
  );
  const [semanticFilters, setSemanticFilters] = useState<SemanticFilterState>(
    DEFAULT_SEMANTIC_FILTERS,
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

  const toggleSemanticAppLayer = useCallback(
    (layer: SemanticFilterState['visibleAppLayers'][number]) => {
      setSemanticFilters((prev) => ({
        ...prev,
        visibleAppLayers: toggleArrayValue(prev.visibleAppLayers, layer),
      }));
    },
    [],
  );

  const toggleSemanticMissingAppLayer = useCallback(() => {
    setSemanticFilters((prev) => ({
      ...prev,
      showNodesMissingAppLayer: !prev.showNodesMissingAppLayer,
    }));
  }, []);

  const toggleResolutionConfidence = useCallback(
    (confidence: SemanticFilterState['visibleResolutionConfidences'][number]) => {
      setSemanticFilters((prev) => ({
        ...prev,
        visibleResolutionConfidences: toggleArrayValue(
          prev.visibleResolutionConfidences,
          confidence,
        ),
      }));
    },
    [],
  );

  const toggleResolutionHealthBucket = useCallback(
    (bucket: SemanticFilterState['visibleResolutionHealthBuckets'][number]) => {
      setSemanticFilters((prev) => ({
        ...prev,
        visibleResolutionHealthBuckets: toggleArrayValue(
          prev.visibleResolutionHealthBuckets,
          bucket,
        ),
      }));
    },
    [],
  );

  const toggleResolutionGapFactFamily = useCallback((family: string) => {
    setSemanticFilters((prev) => ({
      ...prev,
      visibleResolutionGapFactFamilies: toggleArrayValue(
        prev.visibleResolutionGapFactFamilies,
        family,
      ),
    }));
  }, []);

  const toggleResolutionGapTargetRole = useCallback((role: string) => {
    setSemanticFilters((prev) => ({
      ...prev,
      visibleResolutionGapTargetRoles: toggleArrayValue(
        prev.visibleResolutionGapTargetRoles,
        role,
      ),
    }));
  }, []);

  const toggleResolutionGapClassification = useCallback(
    (classification: SemanticFilterState['visibleResolutionGapClassifications'][number]) => {
      setSemanticFilters((prev) => ({
        ...prev,
        visibleResolutionGapClassifications: toggleArrayValue(
          prev.visibleResolutionGapClassifications,
          classification,
        ),
      }));
    },
    [],
  );

  const toggleResolutionGapActionability = useCallback(
    (actionability: SemanticFilterState['visibleResolutionGapActionabilities'][number]) => {
      setSemanticFilters((prev) => ({
        ...prev,
        visibleResolutionGapActionabilities: toggleArrayValue(
          prev.visibleResolutionGapActionabilities,
          actionability,
        ),
      }));
    },
    [],
  );

  const toggleResolutionGapSourceAppLayer = useCallback(
    (layer: SemanticFilterState['visibleResolutionGapSourceAppLayers'][number]) => {
      setSemanticFilters((prev) => ({
        ...prev,
        visibleResolutionGapSourceAppLayers: toggleArrayValue(
          prev.visibleResolutionGapSourceAppLayers,
          layer,
        ),
      }));
    },
    [],
  );

  const toggleResolutionGapTargetText = useCallback((targetText: string) => {
    setSemanticFilters((prev) => ({
      ...prev,
      visibleResolutionGapTargetTexts: toggleArrayValue(
        prev.visibleResolutionGapTargetTexts,
        targetText,
      ),
    }));
  }, []);

  const resetSemanticFilters = useCallback(() => {
    setSemanticFilters(DEFAULT_SEMANTIC_FILTERS);
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
      semanticFilters,
      toggleSemanticAppLayer,
      toggleSemanticMissingAppLayer,
      toggleResolutionConfidence,
      toggleResolutionHealthBucket,
      toggleResolutionGapFactFamily,
      toggleResolutionGapTargetRole,
      toggleResolutionGapClassification,
      toggleResolutionGapActionability,
      toggleResolutionGapSourceAppLayer,
      toggleResolutionGapTargetText,
      resetSemanticFilters,
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
      semanticFilters,
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
