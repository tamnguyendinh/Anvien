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

let mockAppState: Record<string, unknown>;
let toggleLabelVisibility: ReturnType<typeof vi.fn>;
let toggleEdgeVisibility: ReturnType<typeof vi.fn>;

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
    mockAppState = {
      graph,
      visibleLabels: graph.nodes.map((node) => node.label),
      toggleLabelVisibility,
      visibleEdgeTypes: graph.relationships.map(
        (relationship) => relationship.type,
      ),
      toggleEdgeVisibility,
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
