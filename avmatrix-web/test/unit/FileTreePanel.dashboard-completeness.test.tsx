import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import {
  GRAPH_RELATIONSHIP_TYPES,
  NODE_LABELS,
  type GraphNode,
  type GraphRelationship,
  type NodeLabel,
  type RelationshipType,
} from '../../src/generated/avmatrix-contracts';
import type { KnowledgeGraph } from '../../src/core/graph/types';
import { getEdgeInfo } from '../../src/lib/constants';

let mockAppState: Record<string, unknown>;
let toggleLabelVisibility: ReturnType<typeof vi.fn>;
let toggleEdgeVisibility: ReturnType<typeof vi.fn>;

vi.mock('../../src/hooks/useAppState.local-runtime', () => ({
  useAppState: () => mockAppState,
}));

import { FileTreePanel } from '../../src/components/FileTreePanel';

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
    sourceId: 'node-0',
    targetId: 'node-1',
    type: type as RelationshipType,
    confidence: 1,
    reason: 'test',
  }) as GraphRelationship;

const makeGraph = (): KnowledgeGraph => {
  const nodes = [...NODE_LABELS, 'FutureNode'].map(makeNode);
  const relationships = [
    ...GRAPH_RELATIONSHIP_TYPES,
    'FUTURE_RELATIONSHIP',
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

describe('FileTreePanel dashboard completeness', () => {
  beforeEach(() => {
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

  it('renders every node label and relationship type present in the loaded graph', async () => {
    render(<FileTreePanel onFocusNode={vi.fn()} />);

    await userEvent.click(screen.getByRole('button', { name: 'Filters' }));

    for (const label of [...NODE_LABELS, 'FutureNode']) {
      expect(screen.getByTitle(`${label} (1)`)).toBeInTheDocument();
    }

    for (const type of GRAPH_RELATIONSHIP_TYPES) {
      expect(
        screen.getByTitle(`${getEdgeInfo(type).label} (1)`),
      ).toBeInTheDocument();
    }
    expect(screen.getByTitle('Future Relationship (1)')).toBeInTheDocument();
    expect(screen.getAllByText('FutureNode').length).toBeGreaterThanOrEqual(2);
    expect(screen.getAllByText('Future Relationship').length).toBeGreaterThanOrEqual(2);
    expect(screen.getByTitle('Legend node FutureNode (1)')).toBeInTheDocument();
    expect(screen.getByTitle('Legend edge Future Relationship (1)')).toBeInTheDocument();
    expect(
      screen.getByTitle('Community color set (1 communities, 1 members)'),
    ).toBeInTheDocument();
  });

  it('routes every graph-present node and relationship control through the visibility toggles', async () => {
    render(<FileTreePanel onFocusNode={vi.fn()} />);

    await userEvent.click(screen.getByRole('button', { name: 'Filters' }));

    for (const label of [...NODE_LABELS, 'FutureNode']) {
      await userEvent.click(screen.getByTitle(`${label} (1)`));
      expect(toggleLabelVisibility).toHaveBeenLastCalledWith(label);
    }

    for (const type of [...GRAPH_RELATIONSHIP_TYPES, 'FUTURE_RELATIONSHIP']) {
      await userEvent.click(
        screen.getByTitle(`${getEdgeInfo(type).label} (1)`),
      );
      expect(toggleEdgeVisibility).toHaveBeenLastCalledWith(type);
    }
  });
});
