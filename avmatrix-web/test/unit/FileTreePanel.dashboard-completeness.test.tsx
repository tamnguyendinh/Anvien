import { fireEvent, render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { beforeEach, describe, expect, it, vi } from "vitest";
import {
  GRAPH_RELATIONSHIP_TYPES,
  NODE_LABELS,
  type GraphNode,
  type GraphRelationship,
  type NodeLabel,
  type RelationshipType,
} from "../../src/generated/avmatrix-contracts";
import type { KnowledgeGraph } from "../../src/core/graph/types";
import { getEdgeInfo } from "../../src/lib/constants";
import { DEFAULT_GRAPH_HEALTH_FILTERS } from "../../src/lib/graph-health-filters";

let mockAppState: Record<string, unknown>;
let toggleLabelVisibility: ReturnType<typeof vi.fn>;
let toggleEdgeVisibility: ReturnType<typeof vi.fn>;
let toggleGraphHealthTopologyStatus: ReturnType<typeof vi.fn>;
let toggleGraphHealthExpectedReason: ReturnType<typeof vi.fn>;
let toggleGraphHealthDiagnosticKind: ReturnType<typeof vi.fn>;
let resetGraphHealthFilters: ReturnType<typeof vi.fn>;

vi.mock("../../src/hooks/useAppState.local-runtime", () => ({
  useAppState: () => mockAppState,
}));

import { FileTreePanel } from "../../src/components/FileTreePanel";

const makeNode = (label: string, index: number): GraphNode =>
  ({
    id: `node-${index}`,
    label: label as NodeLabel,
    properties: {
      name: label,
      filePath: `src/${label.toLowerCase()}-${index}.ts`,
      graphHealth:
        index === 1
          ? {
              topologyStatus: "no_incoming",
              countedIncoming: 0,
              countedOutgoing: 1,
              componentReachableFromRoot: false,
              expectedIsolationReasons: ["test"],
              diagnostics: [{ kind: "unresolved_reference", count: 2 }],
              confidence: "unknown",
            }
          : {
              topologyStatus: "connected",
              countedIncoming: 1,
              countedOutgoing: 1,
              componentReachableFromRoot: true,
              confidence: "candidate",
            },
    },
  }) as GraphNode;

const makeRelationship = (type: string, index: number): GraphRelationship =>
  ({
    id: `rel-${index}`,
    sourceId: `node-${index}`,
    targetId: `node-${index + 1}`,
    type: type as RelationshipType,
    confidence: 1,
    reason: "test",
  }) as GraphRelationship;

const makeGraph = (): KnowledgeGraph => {
  const nodes = [...NODE_LABELS, "FutureNode"].map(makeNode);
  const relationships = [
    ...GRAPH_RELATIONSHIP_TYPES,
    "FUTURE_RELATIONSHIP",
  ].map(makeRelationship);
  return {
    nodes,
    relationships,
    nodeCount: nodes.length,
    relationshipCount: relationships.length,
    addNode: vi.fn(),
    addRelationship: vi.fn(),
  };
};

describe("FileTreePanel dashboard completeness", () => {
  beforeEach(() => {
    window.localStorage.removeItem("avmatrix.leftPanelWidth");
    const graph = makeGraph();
    toggleLabelVisibility = vi.fn();
    toggleEdgeVisibility = vi.fn();
    toggleGraphHealthTopologyStatus = vi.fn();
    toggleGraphHealthExpectedReason = vi.fn();
    toggleGraphHealthDiagnosticKind = vi.fn();
    resetGraphHealthFilters = vi.fn();
    mockAppState = {
      graph,
      visibleLabels: graph.nodes.map((node) => node.label),
      toggleLabelVisibility,
      visibleEdgeTypes: graph.relationships.map(
        (relationship) => relationship.type,
      ),
      toggleEdgeVisibility,
      graphHealthFilters: DEFAULT_GRAPH_HEALTH_FILTERS,
      toggleGraphHealthTopologyStatus,
      toggleGraphHealthExpectedReason,
      toggleGraphHealthDiagnosticKind,
      resetGraphHealthFilters,
      selectedNode: null,
      setSelectedNode: vi.fn(),
      openCodePanel: vi.fn(),
      depthFilter: null,
      setDepthFilter: vi.fn(),
    };
  });

  it("renders every node label and relationship type present in the loaded graph", async () => {
    render(<FileTreePanel onFocusNode={vi.fn()} />);

    await userEvent.click(screen.getByRole("button", { name: "Filters" }));

    for (const label of [...NODE_LABELS, "FutureNode"]) {
      expect(screen.getByTitle(`${label} (1)`)).toBeInTheDocument();
    }

    for (const type of GRAPH_RELATIONSHIP_TYPES) {
      expect(
        screen.getByTitle(`${getEdgeInfo(type).label} (1)`),
      ).toBeInTheDocument();
    }
    expect(screen.getByTitle("Future Relationship (1)")).toBeInTheDocument();
    expect(screen.getAllByText("FutureNode").length).toBeGreaterThanOrEqual(2);
    expect(
      screen.getAllByText("Future Relationship").length,
    ).toBeGreaterThanOrEqual(2);
    expect(screen.getByTitle("Legend node FutureNode (1)")).toBeInTheDocument();
    expect(
      screen.getByTitle("Legend edge Future Relationship (1)"),
    ).toBeInTheDocument();
    expect(
      screen.getByTitle("Community color set (1 communities, 1 members)"),
    ).toBeInTheDocument();
  });

  it("hides zero-count contract labels and relationships in loaded-graph mode", async () => {
    const graph: KnowledgeGraph = {
      nodes: [makeNode("File", 1)],
      relationships: [makeRelationship("CALLS", 1)],
      nodeCount: 1,
      relationshipCount: 1,
      addNode: vi.fn(),
      addRelationship: vi.fn(),
    };
    mockAppState = {
      ...mockAppState,
      graph,
      visibleLabels: ["File"],
      visibleEdgeTypes: ["CALLS"],
    };

    render(<FileTreePanel onFocusNode={vi.fn()} />);

    await userEvent.click(screen.getByRole("button", { name: "Filters" }));

    expect(screen.getByTitle("File (1)")).toBeInTheDocument();
    expect(screen.getByTitle("Calls (1)")).toBeInTheDocument();
    expect(screen.queryByTitle("Class (0)")).not.toBeInTheDocument();
    expect(screen.queryByTitle("Imports (0)")).not.toBeInTheDocument();
    expect(screen.queryByTitle("Legend node Class (0)")).not.toBeInTheDocument();
    expect(
      screen.queryByTitle("Legend edge Imports (0)"),
    ).not.toBeInTheDocument();
  });

  it("groups duplicate normalized heritage edges in counts and legend titles", async () => {
    const graph = makeGraph();
    graph.relationships.push(
      {
        id: "extends-area",
        sourceId: "node-9",
        targetId: "node-10",
        type: "EXTENDS",
        confidence: 1,
        reason: "extends",
      } as GraphRelationship,
      {
        id: "inherits-area",
        sourceId: "node-9",
        targetId: "node-10",
        type: "INHERITS",
        confidence: 1,
        reason: "scope-resolution: inherits",
      } as GraphRelationship,
    );
    mockAppState = {
      ...mockAppState,
      graph,
      visibleEdgeTypes: graph.relationships.map(
        (relationship) => relationship.type,
      ),
    };

    render(<FileTreePanel onFocusNode={vi.fn()} />);

    await userEvent.click(screen.getByRole("button", { name: "Filters" }));

    expect(
      screen.getByTitle(
        "Normalized Heritage (1; 1 grouped compatibility edges, 2 raw)",
      ),
    ).toBeInTheDocument();
    expect(
      screen.getByTitle(
        "Legend edge Normalized Heritage (1; 1 grouped compatibility edges, 2 raw)",
      ),
    ).toBeInTheDocument();
  });

  it("routes every graph-present node and relationship control through the visibility toggles", async () => {
    render(<FileTreePanel onFocusNode={vi.fn()} />);

    await userEvent.click(screen.getByRole("button", { name: "Filters" }));

    for (const label of [...NODE_LABELS, "FutureNode"]) {
      await userEvent.click(screen.getByTitle(`${label} (1)`));
      expect(toggleLabelVisibility).toHaveBeenLastCalledWith(label);
    }

    for (const type of [...GRAPH_RELATIONSHIP_TYPES, "FUTURE_RELATIONSHIP"]) {
      await userEvent.click(
        screen.getByTitle(`${getEdgeInfo(type).label} (1)`),
      );
      expect(toggleEdgeVisibility).toHaveBeenLastCalledWith(type);
    }
  });

  it("renders Graph Health filters and routes controls through app state", async () => {
    render(<FileTreePanel onFocusNode={vi.fn()} />);

    await userEvent.click(screen.getByRole("button", { name: "Filters" }));

    expect(screen.getByTestId("graph-health-filter-section")).toBeInTheDocument();
    expect(screen.getByTitle(/No incoming \(1\)/)).toBeInTheDocument();
    expect(screen.getByTitle(/Connected \(\d+\)/)).toBeInTheDocument();
    expect(screen.getByTitle(/Test \(1\)/)).toBeInTheDocument();
    expect(screen.getByTitle(/Framework entry \(0\)/)).toBeInTheDocument();
    expect(screen.getByTitle(/Unresolved reference \(2\)/)).toBeInTheDocument();
    expect(screen.getByTitle(/Unknown \(1\)/)).toBeInTheDocument();

    await userEvent.click(screen.getByTitle(/No incoming \(1\)/));
    expect(toggleGraphHealthTopologyStatus).toHaveBeenLastCalledWith("no_incoming");

    await userEvent.click(screen.getByTitle(/Test \(1\)/));
    expect(toggleGraphHealthExpectedReason).toHaveBeenLastCalledWith("test");

    await userEvent.click(screen.getByTitle(/Unresolved reference \(2\)/));
    expect(toggleGraphHealthDiagnosticKind).toHaveBeenLastCalledWith(
      "unresolved_reference",
    );

    await userEvent.click(screen.getByTitle("Reset Graph Health filters"));
    expect(resetGraphHealthFilters).toHaveBeenCalledTimes(1);
  });

  it("routes focus-depth controls through app state and warns without a selected node", async () => {
    const setDepthFilter = vi.fn();
    mockAppState = {
      ...mockAppState,
      depthFilter: null,
      selectedNode: null,
      setDepthFilter,
    };

    const { rerender } = render(<FileTreePanel onFocusNode={vi.fn()} />);

    await userEvent.click(screen.getByRole("button", { name: "Filters" }));
    await userEvent.click(screen.getByRole("button", { name: "2 hops" }));
    expect(setDepthFilter).toHaveBeenLastCalledWith(2);

    mockAppState = {
      ...mockAppState,
      depthFilter: 2,
      selectedNode: null,
      setDepthFilter,
    };
    rerender(<FileTreePanel onFocusNode={vi.fn()} />);

    expect(
      screen.getByText("Select a node to apply depth filter"),
    ).toBeInTheDocument();
    await userEvent.click(screen.getByRole("button", { name: "All" }));
    expect(setDepthFilter).toHaveBeenLastCalledWith(null);
  });

  it("resizes the left dashboard from the drag handle within bounds", async () => {
    render(<FileTreePanel onFocusNode={vi.fn()} />);

    const panel = document.querySelector(".file-tree-panel") as HTMLElement;
    const handle = screen.getByTestId("left-dashboard-resize-handle");

    expect(panel.style.width).toBe("248px");

    fireEvent.pointerDown(handle, { clientX: 248 });
    fireEvent.pointerMove(window, { clientX: 980 });
    fireEvent.pointerUp(window);
    expect(panel.style.width).toBe("480px");

    fireEvent.pointerDown(handle, { clientX: 480 });
    fireEvent.pointerMove(window, { clientX: 0 });
    fireEvent.pointerUp(window);
    expect(panel.style.width).toBe("192px");
  });
});
