/**
 * Web-specific graph types.
 *
 * Shared types (NodeLabel, GraphNode, etc.) should be imported
 * directly from the Go-generated Web contract at call sites.
 *
 * This file only defines web-specific additions.
 */
import type {
  GraphNode,
  GraphRelationship,
  GraphSemanticStatus,
} from '@/generated/anvien-contracts';

// Web-specific: in-memory graph container (simpler than CLI version)
export interface KnowledgeGraph {
  nodes: GraphNode[];
  relationships: GraphRelationship[];
  semanticStatus?: GraphSemanticStatus;
  nodeCount: number;
  relationshipCount: number;
  addNode: (node: GraphNode) => void;
  addRelationship: (relationship: GraphRelationship) => void;
}
