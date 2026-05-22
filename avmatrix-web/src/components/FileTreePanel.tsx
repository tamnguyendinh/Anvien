import { useState, useMemo, useCallback, useEffect, useRef } from "react";
import type { PointerEvent as ReactPointerEvent } from "react";
import {
  ChevronRight,
  ChevronDown,
  AlertTriangle,
  Folder,
  FolderOpen,
  FileCode,
  Search,
  Filter,
  PanelLeftClose,
  PanelLeft,
  Box,
  Braces,
  Variable,
  Hash,
  Target,
  List,
  AtSign,
  Type,
  Code,
  GitBranch,
  Globe,
  Layers,
  Server,
  Square,
  Table,
  Zap,
} from "@/lib/lucide-icons";
import { useAppState } from "../hooks/useAppState.local-runtime";
import {
  ALL_EDGE_TYPES,
  COMMUNITY_COLORS,
  DOCUMENTATION_NODE_LABEL,
  FILTERABLE_LABELS,
  getDisplayRelationshipTypeCounts,
  getEdgeInfo,
  getFilterableEdgeTypesForGraph,
  getFilterableNodeLabelsForGraph,
  getGroupedHeritageCompatibilityCount,
  getNodeColor,
  getNodeLabelCounts,
  getRelationshipTypeCounts,
} from "../lib/constants";
import {
  GRAPH_HEALTH_CONFIDENCE_LEVELS,
  GRAPH_HEALTH_EXPECTED_ISOLATION_REASONS,
  GRAPH_HEALTH_TOPOLOGY_STATUSES,
  type GraphHealthConfidence,
  type GraphHealthExpectedIsolationReason,
  type GraphHealthTopologyStatus,
  type GraphNode,
} from "@/generated/avmatrix-contracts";
import {
  GRAPH_HEALTH_CONFIDENCE_DESCRIPTIONS,
  GRAPH_HEALTH_CONFIDENCE_LABELS,
  GRAPH_HEALTH_DIAGNOSTIC_KINDS,
  GRAPH_HEALTH_DIAGNOSTIC_DESCRIPTIONS,
  GRAPH_HEALTH_DIAGNOSTIC_LABELS,
  GRAPH_HEALTH_REASON_DESCRIPTIONS,
  GRAPH_HEALTH_REASON_LABELS,
  GRAPH_HEALTH_TOPOLOGY_DESCRIPTIONS,
  GRAPH_HEALTH_TOPOLOGY_LABELS,
  getGraphHealthConfidenceCounts,
  getGraphHealthDiagnosticCounts,
  getGraphHealthExpectedReasonCounts,
  getGraphHealthTopologyCounts,
  type GraphHealthDiagnosticKind,
} from "../lib/graph-health-filters";

// Tree node structure
interface TreeNode {
  id: string;
  name: string;
  type: "folder" | "file";
  path: string;
  children: TreeNode[];
  graphNode?: GraphNode;
}

// Build tree from graph nodes
const buildFileTree = (nodes: GraphNode[]): TreeNode[] => {
  const root: TreeNode[] = [];
  const pathMap = new Map<string, TreeNode>();

  // Filter to only folders and files
  const fileNodes = nodes.filter(
    (n) => n.label === "Folder" || n.label === "File",
  );

  // Sort by path to ensure parents come before children
  fileNodes.sort((a, b) =>
    a.properties.filePath.localeCompare(b.properties.filePath),
  );

  fileNodes.forEach((node) => {
    const parts = node.properties.filePath.split("/").filter(Boolean);
    let currentPath = "";
    let currentLevel = root;

    parts.forEach((part: string, index: number) => {
      currentPath = currentPath ? `${currentPath}/${part}` : part;

      let existing = pathMap.get(currentPath);

      if (!existing) {
        const isLastPart = index === parts.length - 1;
        const isFile = isLastPart && node.label === "File";

        existing = {
          id: isLastPart ? node.id : currentPath,
          name: part,
          type: isFile ? "file" : "folder",
          path: currentPath,
          children: [],
          graphNode: isLastPart ? node : undefined,
        };

        pathMap.set(currentPath, existing);
        currentLevel.push(existing);
      }

      currentLevel = existing.children;
    });
  });

  return root;
};

// Tree item component
interface TreeItemProps {
  node: TreeNode;
  depth: number;
  searchQuery: string;
  onNodeClick: (node: TreeNode) => void;
  expandedPaths: Set<string>;
  toggleExpanded: (path: string) => void;
  selectedPath: string | null;
}

const TreeItem = ({
  node,
  depth,
  searchQuery,
  onNodeClick,
  expandedPaths,
  toggleExpanded,
  selectedPath,
}: TreeItemProps) => {
  const isExpanded = expandedPaths.has(node.path);
  const isSelected = selectedPath === node.path;
  const hasChildren = node.children.length > 0;

  // Filter children based on search (recursive)
  const filteredChildren = useMemo(() => {
    if (!searchQuery) return node.children;
    const searchLower = searchQuery.toLowerCase();
    const matchesSearch = (node: TreeNode, query: string): boolean => {
      if (node.name.toLowerCase().includes(query)) return true;
      return (
        node.children?.some((child) => matchesSearch(child, query)) ?? false
      );
    };
    return node.children.filter((child) => matchesSearch(child, searchLower));
  }, [node.children, searchQuery]);

  // Check if this node matches search
  const matchesSearch =
    searchQuery && node.name.toLowerCase().includes(searchQuery.toLowerCase());

  const handleClick = () => {
    if (hasChildren) {
      toggleExpanded(node.path);
    }
    onNodeClick(node);
  };

  return (
    <div>
      <button
        onClick={handleClick}
        className={`relative flex w-full items-center gap-1.5 rounded px-2 py-1.5 text-left text-sm transition-colors hover:bg-base ${isSelected ? "border-l-[3px] border-border-strong bg-base text-text-primary" : "border-l-[3px] border-transparent text-text-primary hover:text-text-primary"} ${matchesSearch ? "bg-base" : ""} `}
        style={{ paddingLeft: `${depth * 12 + 8}px` }}
      >
        {/* Expand/collapse icon */}
        {hasChildren ? (
          isExpanded ? (
            <ChevronDown className="h-3.5 w-3.5 shrink-0 text-text-secondary" />
          ) : (
            <ChevronRight className="h-3.5 w-3.5 shrink-0 text-text-secondary" />
          )
        ) : (
          <span className="w-3.5" />
        )}

        {/* Node icon */}
        {node.type === "folder" ? (
          isExpanded ? (
            <FolderOpen
              className="h-4 w-4 shrink-0"
              style={{ color: getNodeColor("Folder") }}
            />
          ) : (
            <Folder
              className="h-4 w-4 shrink-0"
              style={{ color: getNodeColor("Folder") }}
            />
          )
        ) : (
          <FileCode
            className="h-4 w-4 shrink-0"
            style={{ color: getNodeColor("File") }}
          />
        )}

        {/* Name */}
        <span
          className={`truncate font-mono text-xs ${isSelected ? "font-semibold" : "font-medium"}`}
        >
          {node.name}
        </span>
      </button>

      {/* Children */}
      {isExpanded && filteredChildren.length > 0 && (
        <div>
          {filteredChildren.map((child) => (
            <TreeItem
              key={child.id}
              node={child}
              depth={depth + 1}
              searchQuery={searchQuery}
              onNodeClick={onNodeClick}
              expandedPaths={expandedPaths}
              toggleExpanded={toggleExpanded}
              selectedPath={selectedPath}
            />
          ))}
        </div>
      )}
    </div>
  );
};

// Icon for node types
const getNodeTypeIcon = (label: string) => {
  switch (label) {
    case "Project":
    case "Tool":
      return Zap;
    case "Package":
    case "Module":
    case "Namespace":
      return Layers;
    case "Folder":
      return Folder;
    case "File":
      return FileCode;
    case "Class":
    case "Struct":
    case "Record":
      return Box;
    case "Function":
    case "Method":
    case "Constructor":
    case "Delegate":
    case "Impl":
      return Braces;
    case "Interface":
    case "Trait":
      return Hash;
    case "Enum":
    case "Union":
      return List;
    case "Type":
    case "TypeAlias":
    case "Typedef":
    case "Template":
      return Type;
    case "Decorator":
    case "Annotation":
    case "Macro":
      return AtSign;
    case "Import":
      return FileCode;
    case "Variable":
    case "Const":
    case "Static":
    case "Property":
      return Variable;
    case "CodeElement":
      return Code;
    case "Community":
      return Globe;
    case "Process":
      return GitBranch;
    case "ResolutionGap":
      return AlertTriangle;
    case "Section":
    case DOCUMENTATION_NODE_LABEL:
      return Table;
    case "Route":
      return Server;
    default:
      return Square;
  }
};

interface FileTreePanelProps {
  onFocusNode: (nodeId: string) => void;
}

const LEFT_PANEL_STORAGE_KEY = "avmatrix.leftPanelWidth";
const LEFT_PANEL_MIN_WIDTH = 192;
const LEFT_PANEL_DEFAULT_WIDTH = 248;
const LEFT_PANEL_MAX_WIDTH = 480;

const clampLeftPanelWidth = (width: number): number =>
  Math.min(LEFT_PANEL_MAX_WIDTH, Math.max(LEFT_PANEL_MIN_WIDTH, width));

const loadStoredLeftPanelWidth = (): number => {
  if (typeof window === "undefined") return LEFT_PANEL_DEFAULT_WIDTH;
  const raw = window.localStorage.getItem(LEFT_PANEL_STORAGE_KEY);
  const parsed = raw ? Number.parseInt(raw, 10) : Number.NaN;
  return Number.isFinite(parsed)
    ? clampLeftPanelWidth(parsed)
    : LEFT_PANEL_DEFAULT_WIDTH;
};

export const FileTreePanel = ({ onFocusNode }: FileTreePanelProps) => {
  const {
    graph,
    visibleLabels,
    toggleLabelVisibility,
    visibleEdgeTypes,
    toggleEdgeVisibility,
    graphHealthFilters,
    toggleGraphHealthTopologyStatus,
    toggleGraphHealthExpectedReason,
    toggleGraphHealthDiagnosticKind,
    resetGraphHealthFilters,
    selectedNode,
    setSelectedNode,
    openCodePanel,
    depthFilter,
    setDepthFilter,
  } = useAppState();

  const [isCollapsed, setIsCollapsed] = useState(false);
  const [searchQuery, setSearchQuery] = useState("");
  const [expandedPaths, setExpandedPaths] = useState<Set<string>>(new Set());
  const [activeTab, setActiveTab] = useState<"files" | "filters">("files");
  const [panelWidth, setPanelWidth] = useState(loadStoredLeftPanelWidth);
  const resizeRef = useRef<{ startX: number; startWidth: number } | null>(null);

  // Build file tree from graph
  const fileTree = useMemo(() => {
    if (!graph) return [];
    return buildFileTree(graph.nodes);
  }, [graph]);

  // Auto-expand first level on initial load
  useEffect(() => {
    if (fileTree.length > 0 && expandedPaths.size === 0) {
      const firstLevel = new Set(fileTree.map((n) => n.path));
      setExpandedPaths(firstLevel);
    }
  }, [fileTree.length]); // Only run when tree first loads

  // Auto-expand to selected file when selectedNode changes (e.g., from graph click)
  useEffect(() => {
    const path = selectedNode?.properties?.filePath;
    if (!path) return;

    // Expand all parent folders leading to this file
    const parts = path.split("/").filter(Boolean);
    const pathsToExpand: string[] = [];
    let currentPath = "";

    // Build all parent paths (exclude the last part if it's a file)
    for (let i = 0; i < parts.length - 1; i++) {
      currentPath = currentPath ? `${currentPath}/${parts[i]}` : parts[i];
      pathsToExpand.push(currentPath);
    }

    if (pathsToExpand.length > 0) {
      setExpandedPaths((prev) => {
        const next = new Set(prev);
        pathsToExpand.forEach((p) => next.add(p));
        return next;
      });
    }
  }, [selectedNode?.id]); // Trigger when selected node changes

  useEffect(() => {
    if (typeof window === "undefined") return;
    window.localStorage.setItem(LEFT_PANEL_STORAGE_KEY, String(panelWidth));
  }, [panelWidth]);

  useEffect(() => {
    const handleWindowResize = () => {
      setPanelWidth((width) => clampLeftPanelWidth(width));
    };
    window.addEventListener("resize", handleWindowResize);
    return () => window.removeEventListener("resize", handleWindowResize);
  }, []);

  const toggleExpanded = useCallback((path: string) => {
    setExpandedPaths((prev) => {
      const next = new Set(prev);
      if (next.has(path)) {
        next.delete(path);
      } else {
        next.add(path);
      }
      return next;
    });
  }, []);

  const handleNodeClick = useCallback(
    (treeNode: TreeNode) => {
      if (treeNode.graphNode) {
        // Only focus if selecting a different node
        const isSameNode = selectedNode?.id === treeNode.graphNode.id;
        setSelectedNode(treeNode.graphNode);
        openCodePanel();
        if (!isSameNode) {
          onFocusNode(treeNode.graphNode.id);
        }
      }
    },
    [setSelectedNode, openCodePanel, onFocusNode, selectedNode],
  );

  const handleResizePointerDown = useCallback(
    (event: ReactPointerEvent<HTMLButtonElement>) => {
      event.preventDefault();
      resizeRef.current = { startX: event.clientX, startWidth: panelWidth };
      document.body.style.cursor = "col-resize";
      document.body.style.userSelect = "none";

      const handlePointerMove = (moveEvent: PointerEvent) => {
        const state = resizeRef.current;
        if (!state) return;
        setPanelWidth(
          clampLeftPanelWidth(
            state.startWidth + moveEvent.clientX - state.startX,
          ),
        );
      };

      const handlePointerUp = () => {
        resizeRef.current = null;
        document.body.style.cursor = "";
        document.body.style.userSelect = "";
        window.removeEventListener("pointermove", handlePointerMove);
        window.removeEventListener("pointerup", handlePointerUp);
      };

      window.addEventListener("pointermove", handlePointerMove);
      window.addEventListener("pointerup", handlePointerUp);
    },
    [panelWidth],
  );

  const selectedPath = selectedNode?.properties.filePath || null;

  const nodeTypeItems = useMemo(() => {
    if (!graph) {
      return FILTERABLE_LABELS.map((label) => ({
        label,
        count: 0,
        color: getNodeColor(label),
      }));
    }
    const counts = getNodeLabelCounts(graph.nodes);
    return getFilterableNodeLabelsForGraph(graph.nodes).map((label) => ({
      label,
      count: counts.get(label) ?? 0,
      color: getNodeColor(label),
    }));
  }, [graph]);

  const hiddenNodeTypeLabels = useMemo(() => {
    const visibleLabelSet = new Set(visibleLabels);
    return nodeTypeItems
      .map(({ label }) => label)
      .filter((label) => !visibleLabelSet.has(label));
  }, [nodeTypeItems, visibleLabels]);

  const selectAllNodeTypes = useCallback(() => {
    hiddenNodeTypeLabels.forEach((label) => toggleLabelVisibility(label));
  }, [hiddenNodeTypeLabels, toggleLabelVisibility]);

  const edgeTypeItems = useMemo(() => {
    if (!graph) {
      return ALL_EDGE_TYPES.map((type) => ({
        type,
        count: 0,
        rawCount: 0,
        groupedCompatibilityCount: 0,
        info: getEdgeInfo(type),
      }));
    }
    const counts = getRelationshipTypeCounts(graph.relationships);
    const displayCounts = getDisplayRelationshipTypeCounts(graph.relationships);
    return getFilterableEdgeTypesForGraph(graph.relationships).map((type) => ({
      type,
      count: displayCounts.get(type) ?? 0,
      rawCount: counts.get(type) ?? 0,
      groupedCompatibilityCount: getGroupedHeritageCompatibilityCount(
        graph.relationships,
        type,
      ),
      info: getEdgeInfo(type),
    }));
  }, [graph]);

  const graphHealthTopologyItems = useMemo(() => {
    const counts = getGraphHealthTopologyCounts(graph?.nodes ?? []);
    return GRAPH_HEALTH_TOPOLOGY_STATUSES.map((status) => ({
      status,
      label: GRAPH_HEALTH_TOPOLOGY_LABELS[status],
      count: counts.get(status) ?? 0,
    }));
  }, [graph]);

  const graphHealthReasonItems = useMemo(() => {
    const counts = getGraphHealthExpectedReasonCounts(graph?.nodes ?? []);
    return GRAPH_HEALTH_EXPECTED_ISOLATION_REASONS.map((reason) => ({
      reason,
      label: GRAPH_HEALTH_REASON_LABELS[reason],
      count: counts.get(reason) ?? 0,
    }));
  }, [graph]);

  const graphHealthDiagnosticItems = useMemo(() => {
    const counts = getGraphHealthDiagnosticCounts(graph?.nodes ?? []);
    return GRAPH_HEALTH_DIAGNOSTIC_KINDS.map((kind) => ({
      kind,
      label: GRAPH_HEALTH_DIAGNOSTIC_LABELS[kind] ?? kind,
      count: counts.get(kind) ?? 0,
    }));
  }, [graph]);

  const graphHealthConfidenceItems = useMemo(() => {
    const counts = getGraphHealthConfidenceCounts(graph?.nodes ?? []);
    return GRAPH_HEALTH_CONFIDENCE_LEVELS.map((confidence) => ({
      confidence,
      label: GRAPH_HEALTH_CONFIDENCE_LABELS[confidence],
      count: counts.get(confidence) ?? 0,
    }));
  }, [graph]);

  const communityColorLegend = useMemo(() => {
    if (!graph) return null;
    const memberRelationships = graph.relationships.filter(
      (relationship) => relationship.type === "MEMBER_OF",
    );
    if (memberRelationships.length === 0) return null;

    const communityIds = new Set(
      memberRelationships.map((relationship) => relationship.targetId),
    );

    return {
      communityCount: communityIds.size,
      memberCount: memberRelationships.length,
      swatches: COMMUNITY_COLORS.slice(0, 6),
    };
  }, [graph]);

  if (isCollapsed) {
    return (
      <div className="file-tree-panel flex h-full w-12 flex-shrink-0 flex-col items-center gap-2 border-r-[3px] border-border-default bg-surface py-3">
        <button
          onClick={() => setIsCollapsed(false)}
          className="press-ghost-button rounded p-2 text-text-secondary"
          title="Expand Panel"
        >
          <PanelLeft className="h-5 w-5" />
        </button>
        <div className="my-1 h-px w-6 bg-border-subtle" />
        <button
          onClick={() => {
            setIsCollapsed(false);
            setActiveTab("files");
          }}
          className={`rounded p-2 transition-colors ${activeTab === "files" ? "bg-base text-text-primary" : "text-text-secondary hover:bg-base hover:text-text-primary"}`}
          title="File Explorer"
        >
          <Folder className="h-5 w-5" />
        </button>
        <button
          onClick={() => {
            setIsCollapsed(false);
            setActiveTab("filters");
          }}
          className={`rounded p-2 transition-colors ${activeTab === "filters" ? "bg-base text-text-primary" : "text-text-secondary hover:bg-base hover:text-text-primary"}`}
          title="Filters"
        >
          <Filter className="h-5 w-5" />
        </button>
      </div>
    );
  }

  return (
    <div
      className="file-tree-panel relative flex h-full flex-shrink-0 animate-slide-in flex-col border-r-[3px] border-border-default bg-surface"
      style={{ width: panelWidth }}
    >
      <button
        type="button"
        aria-label="Resize left dashboard"
        title="Drag to resize dashboard"
        data-testid="left-dashboard-resize-handle"
        onPointerDown={handleResizePointerDown}
        className="absolute top-0 right-[-5px] z-20 h-full w-2 cursor-col-resize bg-transparent transition-colors hover:bg-border-strong/25"
      />
      <div className="flex items-center justify-between border-b-[3px] border-border-default px-3 py-3">
        <div className="flex items-center gap-1">
          <button
            onClick={() => setActiveTab("files")}
            className={`rounded px-2 py-1 font-mono text-xs transition-colors ${
              activeTab === "files"
                ? "bg-base text-text-primary"
                : "text-text-secondary hover:bg-base hover:text-text-primary"
            }`}
          >
            Explorer
          </button>
          <button
            onClick={() => setActiveTab("filters")}
            className={`rounded px-2 py-1 font-mono text-xs transition-colors ${
              activeTab === "filters"
                ? "bg-base text-text-primary"
                : "text-text-secondary hover:bg-base hover:text-text-primary"
            }`}
          >
            Filters
          </button>
        </div>
        <button
          onClick={() => setIsCollapsed(true)}
          className="press-ghost-button rounded p-1 text-text-muted"
          title="Collapse Panel"
        >
          <PanelLeftClose className="h-4 w-4" />
        </button>
      </div>

      {activeTab === "files" && (
        <>
          {/* Search */}
          <div className="border-b border-border-subtle px-3 py-3">
            <div className="relative">
              <Search className="absolute top-1/2 left-2.5 h-3.5 w-3.5 -translate-y-1/2 text-text-muted" />
              <input
                type="text"
                placeholder="Search files..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full rounded border-[2px] border-border-default bg-inset py-2 pr-3 pl-8 font-mono text-xs text-text-primary placeholder:text-text-muted focus:border-border-strong focus:outline-none"
              />
            </div>
          </div>

          {/* File tree */}
          <div className="scrollbar-thin flex-1 overflow-y-auto py-2">
            {fileTree.length === 0 ? (
              <div className="px-3 py-4 text-center text-xs text-text-muted">
                No files loaded
              </div>
            ) : (
              fileTree.map((node) => (
                <TreeItem
                  key={node.id}
                  node={node}
                  depth={0}
                  searchQuery={searchQuery}
                  onNodeClick={handleNodeClick}
                  expandedPaths={expandedPaths}
                  toggleExpanded={toggleExpanded}
                  selectedPath={selectedPath}
                />
              ))
            )}
          </div>
        </>
      )}

      {activeTab === "filters" && (
        <div className="scrollbar-thin flex-1 overflow-y-auto p-3">
          <div className="mb-3">
            <div className="mb-2 flex items-center justify-between gap-2">
              <h3 className="press-eyebrow text-text-secondary">
                Node Types
              </h3>
              <button
                onClick={selectAllNodeTypes}
                disabled={hiddenNodeTypeLabels.length === 0}
                className="rounded px-2 py-1 font-mono text-[10px] text-text-muted transition-colors hover:bg-base hover:text-text-primary disabled:cursor-default disabled:opacity-40 disabled:hover:bg-transparent disabled:hover:text-text-muted"
                title="Select all node types"
              >
                Select all
              </button>
            </div>
            <p className="mb-3 text-[11px] text-text-muted">
              Toggle visibility of node types in the graph
            </p>
          </div>

          <div className="flex flex-col gap-1">
            {nodeTypeItems.map(({ label, count, color }) => {
              const Icon = getNodeTypeIcon(label);
              const isVisible = visibleLabels.includes(label);

              return (
                <button
                  key={label}
                  onClick={() => toggleLabelVisibility(label)}
                  title={`${label} (${count})`}
                  aria-pressed={isVisible}
                  className={`flex items-center gap-2.5 rounded px-2 py-1.5 text-left transition-colors ${
                    isVisible
                      ? "bg-base text-text-primary"
                      : "text-text-muted hover:bg-base hover:text-text-secondary"
                  } `}
                >
                  <div
                    className={`flex h-5 w-5 items-center justify-center rounded ${isVisible ? "" : "opacity-40"}`}
                    style={{ backgroundColor: `${color}20` }}
                  >
                    <Icon className="h-3 w-3" style={{ color }} />
                  </div>
                  <span className="min-w-0 flex-1 truncate text-xs">
                    {label}
                  </span>
                  <span className="font-mono text-[10px] text-text-muted">
                    {count}
                  </span>
                  <div
                    className={`h-2 w-2 rounded-full transition-colors ${isVisible ? "bg-border-strong" : "bg-border-subtle"}`}
                  />
                </button>
              );
            })}
          </div>

          {/* Graph Health Toggles */}
          <div
            className="mt-6 border-t border-border-subtle pt-4"
            data-testid="graph-health-filter-section"
          >
            <div className="mb-3 flex items-center justify-between gap-2">
              <h3 className="press-eyebrow text-text-secondary">
                <AlertTriangle className="mr-1.5 inline h-3 w-3" />
                Graph Health
              </h3>
              <button
                onClick={resetGraphHealthFilters}
                className="rounded px-2 py-1 font-mono text-[10px] text-text-muted transition-colors hover:bg-base hover:text-text-primary"
                title="Reset Graph Health filters"
              >
                Reset
              </button>
            </div>

            <div className="flex flex-col gap-1">
              {graphHealthTopologyItems.map(({ status, label, count }) => {
                const isConnectedBaseline = status === "connected";
                const isVisible =
                  graphHealthFilters.visibleTopologyStatuses.includes(
                    status as GraphHealthTopologyStatus,
                  );
                const title = `${label} (${count}) - ${GRAPH_HEALTH_TOPOLOGY_DESCRIPTIONS[status]}`;

                if (isConnectedBaseline) {
                  return (
                    <div
                      key={status}
                      title={title}
                      className="flex items-center gap-2.5 rounded px-2 py-1.5 text-left text-text-muted"
                    >
                      <div className="h-2 w-2 rounded-full bg-border-subtle" />
                      <span className="min-w-0 flex-1 truncate text-xs">
                        {label}
                      </span>
                      <span className="font-mono text-[10px]">{count}</span>
                    </div>
                  );
                }

                return (
                  <button
                    key={status}
                    onClick={() => toggleGraphHealthTopologyStatus(status)}
                    title={title}
                    aria-pressed={isVisible}
                    className={`flex items-center gap-2.5 rounded px-2 py-1.5 text-left transition-colors ${
                      isVisible
                        ? "bg-base text-text-primary"
                        : "text-text-muted hover:bg-base hover:text-text-secondary"
                    } `}
                  >
                    <div
                      className={`h-2 w-2 rounded-full ${isVisible ? "bg-border-strong" : "bg-border-subtle"}`}
                    />
                    <span className="min-w-0 flex-1 truncate text-xs">
                      {label}
                    </span>
                    <span className="font-mono text-[10px] text-text-muted">
                      {count}
                    </span>
                  </button>
                );
              })}
            </div>

            <h4 className="mt-4 mb-2 font-mono text-[10px] text-text-muted">
              Expected Isolation
            </h4>
            <div className="flex flex-col gap-1">
              {graphHealthReasonItems.map(({ reason, label, count }) => {
                const isVisible =
                  !graphHealthFilters.hiddenExpectedIsolationReasons.includes(
                    reason as GraphHealthExpectedIsolationReason,
                  );
                return (
                  <button
                    key={reason}
                    onClick={() => toggleGraphHealthExpectedReason(reason)}
                    title={`${label} (${count}) - ${GRAPH_HEALTH_REASON_DESCRIPTIONS[reason]}`}
                    aria-pressed={isVisible}
                    className={`flex items-center gap-2.5 rounded px-2 py-1.5 text-left transition-colors ${
                      isVisible
                        ? "bg-base text-text-primary"
                        : "text-text-muted hover:bg-base hover:text-text-secondary"
                    } `}
                  >
                    <div
                      className={`h-2 w-2 rounded-full ${isVisible ? "bg-border-strong" : "bg-border-subtle"}`}
                    />
                    <span className="min-w-0 flex-1 truncate text-xs">
                      {label}
                    </span>
                    <span className="font-mono text-[10px] text-text-muted">
                      {count}
                    </span>
                  </button>
                );
              })}
            </div>

            <h4 className="mt-4 mb-2 font-mono text-[10px] text-text-muted">
              Diagnostics
            </h4>
            <div className="flex flex-col gap-1">
              {graphHealthDiagnosticItems.map(({ kind, label, count }) => {
                const isVisible =
                  graphHealthFilters.visibleDiagnosticKinds.includes(
                    kind as GraphHealthDiagnosticKind,
                  );
                return (
                  <button
                    key={kind}
                    onClick={() => toggleGraphHealthDiagnosticKind(kind)}
                    title={`${label} (${count}) - ${GRAPH_HEALTH_DIAGNOSTIC_DESCRIPTIONS[kind] ?? "Graph Health diagnostic evidence."}`}
                    aria-pressed={isVisible}
                    className={`flex items-center gap-2.5 rounded px-2 py-1.5 text-left transition-colors ${
                      isVisible
                        ? "bg-base text-text-primary"
                        : "text-text-muted hover:bg-base hover:text-text-secondary"
                    } `}
                  >
                    <div
                      className={`h-2 w-2 rounded-full ${isVisible ? "bg-border-strong" : "bg-border-subtle"}`}
                    />
                    <span className="min-w-0 flex-1 truncate text-xs">
                      {label}
                    </span>
                    <span className="font-mono text-[10px] text-text-muted">
                      {count}
                    </span>
                  </button>
                );
              })}
            </div>

            <h4 className="mt-4 mb-2 font-mono text-[10px] text-text-muted">
              Confidence
            </h4>
            <div className="flex flex-col gap-1">
              {graphHealthConfidenceItems.map(({ confidence, label, count }) => (
                <div
                  key={confidence}
                  title={`${label} (${count}) - ${
                    GRAPH_HEALTH_CONFIDENCE_DESCRIPTIONS[
                      confidence as GraphHealthConfidence
                    ]
                  }`}
                  className="flex items-center gap-2.5 rounded px-2 py-1.5 text-left text-text-muted"
                >
                  <div className="h-2 w-2 rounded-full bg-border-subtle" />
                  <span className="min-w-0 flex-1 truncate text-xs">
                    {label}
                  </span>
                  <span className="font-mono text-[10px]">{count}</span>
                </div>
              ))}
            </div>
          </div>

          {/* Edge Type Toggles */}
          <div className="mt-6 border-t border-border-subtle pt-4">
            <h3 className="press-eyebrow mb-2 text-text-secondary">
              Edge Types
            </h3>
            <p className="mb-3 text-[11px] text-text-muted">
              Toggle visibility of relationship types
            </p>

            <div className="flex flex-col gap-1">
              {edgeTypeItems.map(
                ({
                  type: edgeType,
                  count,
                  rawCount,
                  groupedCompatibilityCount,
                  info,
                }) => {
                  const isVisible = visibleEdgeTypes.includes(edgeType);
                  const title =
                    groupedCompatibilityCount > 0
                      ? `${info.label} (${count}; ${groupedCompatibilityCount} grouped compatibility edges, ${rawCount} raw)`
                      : `${info.label} (${count})`;

                  return (
                    <button
                      key={edgeType}
                      onClick={() => toggleEdgeVisibility(edgeType)}
                      title={title}
                      aria-pressed={isVisible}
                      className={`flex items-center gap-2.5 rounded px-2 py-1.5 text-left transition-colors ${
                        isVisible
                          ? "bg-base text-text-primary"
                          : "text-text-muted hover:bg-base hover:text-text-secondary"
                      } `}
                    >
                      <div
                        className={`h-1.5 w-6 rounded-full ${isVisible ? "" : "opacity-40"}`}
                        style={{ backgroundColor: info.color }}
                      />
                      <span className="min-w-0 flex-1 truncate text-xs">
                        {info.label}
                      </span>
                      <span className="font-mono text-[10px] text-text-muted">
                        {count}
                      </span>
                      <div
                        className={`h-2 w-2 rounded-full transition-colors ${isVisible ? "bg-border-strong" : "bg-border-subtle"}`}
                      />
                    </button>
                  );
                },
              )}
            </div>
          </div>

          {/* Depth Filter */}
          <div className="mt-6 border-t border-border-subtle pt-4">
            <h3 className="press-eyebrow mb-2 text-text-secondary">
              <Target className="mr-1.5 inline h-3 w-3" />
              Focus Depth
            </h3>
            <p className="mb-3 text-[11px] text-text-muted">
              Show nodes within N hops of selection
            </p>

            <div className="flex flex-wrap gap-1.5">
              {[
                { value: null, label: "All" },
                { value: 1, label: "1 hop" },
                { value: 2, label: "2 hops" },
                { value: 3, label: "3 hops" },
                { value: 5, label: "5 hops" },
              ].map(({ value, label }) => (
                <button
                  key={label}
                  onClick={() => setDepthFilter(value)}
                  className={`rounded px-2 py-1 font-mono text-xs transition-colors ${
                    depthFilter === value
                      ? "bg-accent text-text-inverse"
                      : "bg-base text-text-secondary hover:bg-surface hover:text-text-primary"
                  } `}
                >
                  {label}
                </button>
              ))}
            </div>

            {depthFilter !== null && !selectedNode && (
              <p className="mt-2 text-[10px] text-warning">
                Select a node to apply depth filter
              </p>
            )}
          </div>

          {/* Legend */}
          <div className="mt-6 border-t border-border-subtle pt-4">
            <h3 className="press-eyebrow mb-3 text-text-secondary">
              Color Legend
            </h3>
            <h4 className="mb-2 font-mono text-[10px] text-text-muted">
              Node Types
            </h4>
            <div className="grid grid-cols-2 gap-2">
              {nodeTypeItems.map(({ label, count, color }) => (
                <div
                  key={label}
                  className="flex items-center gap-1.5"
                  title={`Legend node ${label} (${count})`}
                >
                  <div
                    className="h-2.5 w-2.5 rounded-full"
                    style={{ backgroundColor: color }}
                  />
                  <span className="min-w-0 truncate text-[10px] text-text-muted">
                    {label}
                  </span>
                  <span className="ml-auto font-mono text-[9px] text-text-muted">
                    {count}
                  </span>
                </div>
              ))}
            </div>

            {communityColorLegend && (
              <div
                className="mt-3 flex items-center gap-1.5"
                title={`Community color set (${communityColorLegend.communityCount} communities, ${communityColorLegend.memberCount} members)`}
              >
                <div className="flex -space-x-1">
                  {communityColorLegend.swatches.map((color) => (
                    <div
                      key={color}
                      className="h-2.5 w-2.5 rounded-full border border-surface"
                      style={{ backgroundColor: color }}
                    />
                  ))}
                </div>
                <span className="min-w-0 flex-1 truncate text-[10px] text-text-muted">
                  Community colors
                </span>
                <span className="font-mono text-[9px] text-text-muted">
                  {communityColorLegend.communityCount}
                </span>
              </div>
            )}

            <h4 className="mt-4 mb-2 font-mono text-[10px] text-text-muted">
              Edge Types
            </h4>
            <div className="flex flex-col gap-2">
              {edgeTypeItems.map(
                ({
                  type,
                  count,
                  rawCount,
                  groupedCompatibilityCount,
                  info,
                }) => {
                  const title =
                    groupedCompatibilityCount > 0
                      ? `Legend edge ${info.label} (${count}; ${groupedCompatibilityCount} grouped compatibility edges, ${rawCount} raw)`
                      : `Legend edge ${info.label} (${count})`;

                  return (
                    <div
                      key={type}
                      className="flex items-center gap-1.5"
                      title={title}
                    >
                      <div
                        className="h-1.5 w-6 rounded-full"
                        style={{ backgroundColor: info.color }}
                      />
                      <span className="min-w-0 truncate text-[10px] text-text-muted">
                        {info.label}
                      </span>
                      <span className="ml-auto font-mono text-[9px] text-text-muted">
                        {count}
                      </span>
                    </div>
                  );
                },
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
