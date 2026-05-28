import { useRef, useEffect, useCallback, useState } from 'react';
import Sigma from 'sigma';
import Graph from 'graphology';
import { DEFAULT_NODE_PROGRAM_CLASSES } from 'sigma/settings';
import {
  SigmaNodeAttributes,
  SigmaEdgeAttributes,
  applyFilterBasedClusteredLayout,
  capRenderedNodeSize,
} from '../lib/graph-adapter';
import type { NodeAnimation } from './useAppState.local-runtime';
import type { EdgeType } from '../lib/constants';
import { getGraphEdgeVisibilityMode } from '../lib/graph-edge-visibility-mode';
import { getSelectedContextEdgeSize } from '../lib/graph-edge-render-style';
import { buildSelectedGraphContext } from '../lib/selected-graph-context';
import {
  recordGraphInteractionMode,
  recordManualLayoutOptimizerInvocation,
} from '../lib/runtime-diagnostics';
import { READABLE_GRAPH_NODE_COUNT_THRESHOLD } from '../lib/graph-readable-camera';
import { NodeSquareProgram } from '../lib/sigma-node-square-program';
import {
  buildDetailFocusCameraAction,
  buildOverviewCameraAction,
} from '../lib/graph-camera-mode';

const DENSE_GRAPH_AMBIENT_EDGE_DIM_AMOUNT = 0.04;
const DENSE_GRAPH_AMBIENT_EDGE_SIZE_MULTIPLIER = 0.12;
const DENSE_GRAPH_AMBIENT_EDGE_MAX_SIZE = 0.12;

// Helper: Parse hex color to RGB
const hexToRgb = (hex: string): { r: number; g: number; b: number } => {
  const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
  return result
    ? {
        r: parseInt(result[1], 16),
        g: parseInt(result[2], 16),
        b: parseInt(result[3], 16),
      }
    : { r: 100, g: 100, b: 100 };
};

// Helper: RGB to hex
const rgbToHex = (r: number, g: number, b: number): string => {
  return (
    '#' +
    [r, g, b]
      .map((x) => {
        const hex = Math.max(0, Math.min(255, Math.round(x))).toString(16);
        return hex.length === 1 ? '0' + hex : hex;
      })
      .join('')
  );
};

// Dim a color by mixing with dark background (keeps color hint)
const dimColor = (hex: string, amount: number): string => {
  const rgb = hexToRgb(hex);
  const darkBg = { r: 18, g: 18, b: 28 }; // #12121c - dark background
  return rgbToHex(
    darkBg.r + (rgb.r - darkBg.r) * amount,
    darkBg.g + (rgb.g - darkBg.g) * amount,
    darkBg.b + (rgb.b - darkBg.b) * amount,
  );
};

// Brighten a color (increase luminosity)
const brightenColor = (hex: string, factor: number): string => {
  const rgb = hexToRgb(hex);
  return rgbToHex(
    rgb.r + ((255 - rgb.r) * (factor - 1)) / factor,
    rgb.g + ((255 - rgb.g) * (factor - 1)) / factor,
    rgb.b + ((255 - rgb.b) * (factor - 1)) / factor,
  );
};

interface UseSigmaOptions {
  onNodeClick?: (nodeId: string) => void;
  onNodeHover?: (nodeId: string | null) => void;
  onStageClick?: () => void;
  highlightedNodeIds?: Set<string>;
  blastRadiusNodeIds?: Set<string>;
  animatedNodes?: Map<string, NodeAnimation>;
  visibleEdgeTypes?: EdgeType[];
  // Ambient-only toggle: selected-node contextual edges may still render when off.
  areGraphLinksVisible?: boolean;
}

interface UseSigmaReturn {
  containerRef: React.RefObject<HTMLDivElement>;
  sigmaRef: React.RefObject<Sigma | null>;
  setGraph: (graph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>) => void;
  zoomIn: () => void;
  zoomOut: () => void;
  resetZoom: () => void;
  focusNode: (nodeId: string) => void;
  isLayoutRunning: boolean;
  startLayout: () => void;
  stopLayout: () => void;
  selectedNode: string | null;
  setSelectedNode: (nodeId: string | null) => void;
  refreshHighlights: () => void;
}

const capNodeReducerSize = (
  attributes: Partial<SigmaNodeAttributes>,
  nodeCount: number,
): Partial<SigmaNodeAttributes> => {
  if (typeof attributes.size === 'number') {
    attributes.size = capRenderedNodeSize(attributes.size, nodeCount);
  }
  return attributes;
};

export const useSigma = (options: UseSigmaOptions = {}): UseSigmaReturn => {
  const containerRef = useRef<HTMLDivElement>(null);
  const sigmaRef = useRef<Sigma | null>(null);
  const graphRef = useRef<Graph<
    SigmaNodeAttributes,
    SigmaEdgeAttributes
  > | null>(null);
  const selectedNodeRef = useRef<string | null>(null);
  const selectedNeighborNodeIdsRef = useRef<Set<string>>(new Set());
  const selectedDirectEdgeIdsRef = useRef<Set<string>>(new Set());
  const highlightedRef = useRef<Set<string>>(new Set());
  const blastRadiusRef = useRef<Set<string>>(new Set());
  const animatedNodesRef = useRef<Map<string, NodeAnimation>>(new Map());
  const visibleEdgeTypesRef = useRef<EdgeType[] | null>(null);
  const areGraphLinksVisibleRef = useRef(true);
  const animationFrameRef = useRef<number | null>(null);
  const [isLayoutRunning, setIsLayoutRunning] = useState(false);
  const [selectedNode, setSelectedNodeState] = useState<string | null>(null);

  useEffect(() => {
    highlightedRef.current = options.highlightedNodeIds || new Set();
    blastRadiusRef.current = options.blastRadiusNodeIds || new Set();
    animatedNodesRef.current = options.animatedNodes || new Map();
    visibleEdgeTypesRef.current = options.visibleEdgeTypes || null;
    areGraphLinksVisibleRef.current = options.areGraphLinksVisible ?? true;
    sigmaRef.current?.refresh();
  }, [
    options.highlightedNodeIds,
    options.blastRadiusNodeIds,
    options.animatedNodes,
    options.visibleEdgeTypes,
    options.areGraphLinksVisible,
  ]);

  // Animation loop for node effects
  useEffect(() => {
    if (!options.animatedNodes || options.animatedNodes.size === 0) {
      if (animationFrameRef.current) {
        cancelAnimationFrame(animationFrameRef.current);
        animationFrameRef.current = null;
      }
      return;
    }

    const animate = () => {
      sigmaRef.current?.refresh();
      animationFrameRef.current = requestAnimationFrame(animate);
    };

    animate();

    return () => {
      if (animationFrameRef.current) {
        cancelAnimationFrame(animationFrameRef.current);
        animationFrameRef.current = null;
      }
    };
  }, [options.animatedNodes]);

  const setSelectedNode = useCallback((nodeId: string | null) => {
    if (selectedNodeRef.current === nodeId) {
      return;
    }

    selectedNodeRef.current = nodeId;
    setSelectedNodeState(nodeId);

    const selectedContext = buildSelectedGraphContext(graphRef.current, nodeId);
    selectedNeighborNodeIdsRef.current = selectedContext.neighborNodeIds;
    selectedDirectEdgeIdsRef.current = selectedContext.directEdgeIds;

    const sigma = sigmaRef.current;
    if (!sigma) return;

    sigma.refresh();
  }, []);

  // Initialize Sigma ONCE
  useEffect(() => {
    if (!containerRef.current) return;
    const container = containerRef.current;

    const graph = new Graph<SigmaNodeAttributes, SigmaEdgeAttributes>();
    graphRef.current = graph;

    const sigma = new Sigma(graph, container, {
      renderLabels: true,
      labelFont: 'JetBrains Mono, monospace',
      labelSize: 11,
      labelWeight: '500',
      labelColor: { color: '#e4e4ed' },
      labelRenderedSizeThreshold: 8,
      labelDensity: 0.1,
      labelGridCellSize: 70,

      defaultNodeColor: '#6b7280',
      defaultEdgeColor: '#2a2a3a',

      // Custom hover renderer - dark background instead of white
      defaultDrawNodeHover: (context, data, settings) => {
        const label = data.label;
        if (!label) return;

        const size = settings.labelSize || 11;
        const font = settings.labelFont || 'JetBrains Mono, monospace';
        const weight = settings.labelWeight || '500';

        context.font = `${weight} ${size}px ${font}`;
        const textWidth = context.measureText(label).width;

        const nodeSize = capRenderedNodeSize(data.size || 8, graph.order);
        const x = data.x;
        const y = data.y - nodeSize - 10;
        const paddingX = 8;
        const paddingY = 5;
        const height = size + paddingY * 2;
        const width = textWidth + paddingX * 2;
        const radius = 4;

        // Dark background pill
        context.fillStyle = '#12121c';
        context.beginPath();
        context.roundRect(x - width / 2, y - height / 2, width, height, radius);
        context.fill();

        // Border matching node color
        context.strokeStyle = data.color || '#6366f1';
        context.lineWidth = 2;
        context.stroke();

        // Label text - light color
        context.fillStyle = '#f5f5f7';
        context.textAlign = 'center';
        context.textBaseline = 'middle';
        context.fillText(label, x, y);

        // Also draw a subtle glow ring around the node
        context.beginPath();
        context.arc(data.x, data.y, nodeSize + 4, 0, Math.PI * 2);
        context.strokeStyle = data.color || '#6366f1';
        context.lineWidth = 2;
        context.globalAlpha = 0.5;
        context.stroke();
        context.globalAlpha = 1;
      },

      minCameraRatio: 0.002,
      maxCameraRatio: 50,
      hideEdgesOnMove: true,
      zIndex: true,
      nodeProgramClasses: {
        ...DEFAULT_NODE_PROGRAM_CLASSES,
        square: NodeSquareProgram,
      },
      nodeHoverProgramClasses: {
        ...DEFAULT_NODE_PROGRAM_CLASSES,
        square: NodeSquareProgram,
      },

      nodeReducer: (node, data) => {
        const res = { ...data };

        if (data.hidden) {
          res.hidden = true;
          return res;
        }

        const currentSelected = selectedNodeRef.current;
        const highlighted = highlightedRef.current;
        const blastRadius = blastRadiusRef.current;
        const animatedNodes = animatedNodesRef.current;
        const hasHighlights = highlighted.size > 0;
        const hasBlastRadius = blastRadius.size > 0;
        const isQueryHighlighted = highlighted.has(node);
        const isBlastRadiusNode = blastRadius.has(node);

        // Apply animation effects FIRST (before other highlighting)
        const animation = animatedNodes.get(node);
        if (animation) {
          const now = Date.now();
          const elapsed = now - animation.startTime;
          const progress = Math.min(elapsed / animation.duration, 1);

          // Calculate animation phase (0-1-0-1... oscillation)
          const phase = (Math.sin(progress * Math.PI * 4) + 1) / 2;

          if (animation.type === 'pulse') {
            // Cyan pulse for search results
            const sizeMultiplier = 1.5 + phase * 0.8;
            res.size = (data.size || 8) * sizeMultiplier;
            res.color = phase > 0.5 ? '#06b6d4' : brightenColor('#06b6d4', 1.3);
            res.zIndex = 5;
            res.highlighted = true;
          } else if (animation.type === 'ripple') {
            // Red ripple for blast radius
            const sizeMultiplier = 1.3 + phase * 1.2;
            res.size = (data.size || 8) * sizeMultiplier;
            res.color = phase > 0.5 ? '#ef4444' : '#f87171';
            res.zIndex = 5;
            res.highlighted = true;
          } else if (animation.type === 'glow') {
            // Purple glow for highlight
            const sizeMultiplier = 1.4 + phase * 0.6;
            res.size = (data.size || 8) * sizeMultiplier;
            res.color = phase > 0.5 ? '#a855f7' : '#c084fc';
            res.zIndex = 5;
            res.highlighted = true;
          }

          return capNodeReducerSize(res, graph.order);
        }

        // Blast radius takes priority (red highlighting)
        if (hasBlastRadius && !currentSelected) {
          if (isBlastRadiusNode) {
            res.color = '#ef4444'; // Red for blast radius
            res.size = (data.size || 8) * 1.8;
            res.zIndex = 3;
            res.highlighted = true;
          } else if (isQueryHighlighted) {
            // Regular cyan highlight for non-blast-radius nodes
            res.color = '#06b6d4';
            res.size = (data.size || 8) * 1.4;
            res.zIndex = 2;
            res.highlighted = true;
          } else {
            res.color = dimColor(data.color, 0.15);
            res.size = (data.size || 8) * 0.4;
            res.zIndex = 0;
          }
          return capNodeReducerSize(res, graph.order);
        }

        if (hasHighlights && !currentSelected) {
          if (isQueryHighlighted) {
            res.color = '#06b6d4';
            res.size = (data.size || 8) * 1.6;
            res.zIndex = 2;
            res.highlighted = true;
          } else {
            res.color = dimColor(data.color, 0.2);
            res.size = (data.size || 8) * 0.5;
            res.zIndex = 0;
          }
          return capNodeReducerSize(res, graph.order);
        }

        if (currentSelected) {
          const isSelected = node === currentSelected;
          const isNeighbor = selectedNeighborNodeIdsRef.current.has(node);

          if (isSelected) {
            res.color = data.color;
            res.size = (data.size || 8) * 1.8;
            res.zIndex = 2;
            res.highlighted = true;
          } else if (isNeighbor) {
            res.color = data.color;
            res.size = (data.size || 8) * 1.3;
            res.zIndex = 1;
          } else {
            res.color = dimColor(data.color, 0.25);
            res.size = (data.size || 8) * 0.6;
            res.zIndex = 0;
          }
        }

        return capNodeReducerSize(res, graph.order);
      },

      edgeReducer: (edge, data) => {
        const res = { ...data };
        const graph = graphRef.current;

        if (!graph) {
          return res;
        }

        const [source, target] = graph.extremities(edge);
        const currentSelected = selectedNodeRef.current;
        const visibilityMode = getGraphEdgeVisibilityMode({
          areAmbientGraphLinksVisible: areGraphLinksVisibleRef.current,
          currentSelectedNodeId: currentSelected,
          sourceNodeId: source,
          targetNodeId: target,
          relationType: data.relationType,
          visibleEdgeTypes: visibleEdgeTypesRef.current,
        });

        if (visibilityMode === 'hidden') {
          res.hidden = true;
          return res;
        }

        const highlighted = highlightedRef.current;
        const blastRadius = blastRadiusRef.current;
        const hasHighlights = highlighted.size > 0 || blastRadius.size > 0; // Check BOTH sets
        const isDenseGraph = graph.order >= READABLE_GRAPH_NODE_COUNT_THRESHOLD;

        if (hasHighlights && !currentSelected) {
          // Check if nodes are in EITHER set
          const isSourceActive =
            highlighted.has(source) || blastRadius.has(source);
          const isTargetActive =
            highlighted.has(target) || blastRadius.has(target);

          const bothHighlighted = isSourceActive && isTargetActive;
          const oneHighlighted = isSourceActive || isTargetActive;

          if (bothHighlighted) {
            // If both nodes are in blast radius, use red edge
            if (blastRadius.has(source) && blastRadius.has(target)) {
              res.color = '#ef4444';
            } else {
              res.color = '#06b6d4';
            }
            res.size = Math.max(2, (data.size || 1) * 3);
            res.zIndex = 2;
          } else if (oneHighlighted) {
            res.color = dimColor('#06b6d4', 0.4);
            res.size = 1;
            res.zIndex = 1;
          } else {
            res.color = dimColor(data.color, 0.08);
            res.size = 0.2;
            res.zIndex = 0;
          }
          return res;
        }

        if (currentSelected) {
          const isConnected =
            visibilityMode === 'selected-context' &&
            selectedDirectEdgeIdsRef.current.has(edge);

          if (isConnected) {
            res.color = brightenColor(data.color, 1.5);
            res.size = getSelectedContextEdgeSize({
              baseSize: data.size || 1,
              areAmbientGraphLinksVisible: areGraphLinksVisibleRef.current,
            });
            res.zIndex = 2;
          } else {
            res.color = dimColor(data.color, 0.1);
            res.size = 0.3;
            res.zIndex = 0;
          }
        } else if (isDenseGraph && visibilityMode === 'ambient') {
          res.color = dimColor(data.color || '#2a2a3a', DENSE_GRAPH_AMBIENT_EDGE_DIM_AMOUNT);
          res.size = Math.min(
            DENSE_GRAPH_AMBIENT_EDGE_MAX_SIZE,
            (data.size || 1) * DENSE_GRAPH_AMBIENT_EDGE_SIZE_MULTIPLIER,
          );
          res.zIndex = 0;
        }

        return res;
      },
    });

    sigmaRef.current = sigma;

    sigma.on('clickNode', ({ node }) => {
      setSelectedNode(node);
      options.onNodeClick?.(node);
    });

    sigma.on('clickStage', () => {
      setSelectedNode(null);
      options.onStageClick?.();
    });

    sigma.on('enterNode', ({ node }) => {
      options.onNodeHover?.(node);
      if (containerRef.current) {
        containerRef.current.style.cursor = 'pointer';
      }
    });

    sigma.on('leaveNode', () => {
      options.onNodeHover?.(null);
      if (containerRef.current) {
        containerRef.current.style.cursor = 'grab';
      }
    });

    const handleWheel = () => {
      recordGraphInteractionMode({
        mode: 'wheel-zoom',
        targetNodeId: selectedNodeRef.current ?? undefined,
      });
    };
    const wheelListenerOptions = { passive: true, capture: true };
    container.addEventListener('wheel', handleWheel, wheelListenerOptions);

    return () => {
      container.removeEventListener('wheel', handleWheel, wheelListenerOptions);
      sigma.kill();
      sigmaRef.current = null;
      graphRef.current = null;
    };
  }, []);

  const setGraph = useCallback(
    (newGraph: Graph<SigmaNodeAttributes, SigmaEdgeAttributes>) => {
      const sigma = sigmaRef.current;
      if (!sigma) return;

      setIsLayoutRunning(false);
      graphRef.current = newGraph;
      const selectedContext = buildSelectedGraphContext(newGraph, null);
      selectedNeighborNodeIdsRef.current = selectedContext.neighborNodeIds;
      selectedDirectEdgeIdsRef.current = selectedContext.directEdgeIds;
      sigma.setGraph(newGraph);
      setSelectedNode(null);

      const overviewAction = buildOverviewCameraAction();
      recordGraphInteractionMode({ mode: 'overview' });
      sigma.getCamera().animatedReset({ duration: overviewAction.durationMs });
    },
    [setSelectedNode],
  );

  const focusNode = useCallback(
    (nodeId: string) => {
      const sigma = sigmaRef.current;
      const graph = graphRef.current;
      if (!sigma || !graph || !graph.hasNode(nodeId)) return;

      const currentSelectedNodeId = selectedNodeRef.current;
      const nodeAttrs = graph.getNodeAttributes(nodeId);
      const cameraState = sigma.getCamera().getState();
      const focusPoint = sigma.viewportToFramedGraph(
        sigma.graphToViewport({ x: nodeAttrs.x, y: nodeAttrs.y }),
      );
      const focusAction = buildDetailFocusCameraAction({
        targetNodeId: nodeId,
        currentSelectedNodeId,
        nodeX: focusPoint.x,
        nodeY: focusPoint.y,
        nodeSize: nodeAttrs.size || 1,
        currentCameraRatio: cameraState.ratio,
        scaler: sigma,
      });

      setSelectedNode(nodeId);
      recordGraphInteractionMode({
        mode: 'detail-focus',
        targetNodeId: nodeId,
      });

      sigma.getCamera().animate(focusAction.cameraState, {
        duration: focusAction.durationMs,
      });

      sigma.refresh();
    },
    [setSelectedNode],
  );

  const zoomIn = useCallback(() => {
    recordGraphInteractionMode({ mode: 'zoom-in' });
    sigmaRef.current?.getCamera().animatedZoom({ duration: 200 });
  }, []);

  const zoomOut = useCallback(() => {
    recordGraphInteractionMode({ mode: 'zoom-out' });
    sigmaRef.current?.getCamera().animatedUnzoom({ duration: 200 });
  }, []);

  const resetZoom = useCallback(() => {
    recordGraphInteractionMode({ mode: 'overview' });
    sigmaRef.current?.getCamera().animatedReset({ duration: 300 });
    setSelectedNode(null);
  }, [setSelectedNode]);

  const startLayout = useCallback(() => {
    const graph = graphRef.current;
    if (!graph || graph.order === 0) return;
    const startedAt = performance.now();
    applyFilterBasedClusteredLayout(graph);
    sigmaRef.current?.refresh();
    const finishedAt = performance.now();
    recordManualLayoutOptimizerInvocation({
      nodeCount: graph.order,
      startedAt,
      finishedAt,
    });
  }, []);

  const stopLayout = useCallback(() => {
    setIsLayoutRunning(false);
  }, []);

  const refreshHighlights = useCallback(() => {
    sigmaRef.current?.refresh();
  }, []);

  return {
    containerRef,
    sigmaRef,
    setGraph,
    zoomIn,
    zoomOut,
    resetZoom,
    focusNode,
    isLayoutRunning,
    startLayout,
    stopLayout,
    selectedNode,
    setSelectedNode,
    refreshHighlights,
  };
};
