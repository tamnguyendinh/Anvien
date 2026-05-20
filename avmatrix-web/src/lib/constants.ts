import {
  GRAPH_RELATIONSHIP_TYPES,
  NODE_LABELS,
  type GraphNode,
  type GraphRelationship,
  type NodeLabel,
  type RelationshipType,
} from "@/generated/avmatrix-contracts";

export type GraphNodeLabel = NodeLabel | (string & {});
export type EdgeType = RelationshipType | (string & {});
export const DOCUMENTATION_NODE_LABEL = "Documentation" as const;
export type DisplayNodeLabel = NodeLabel | typeof DOCUMENTATION_NODE_LABEL;

type RelationshipIdentity = Pick<
  GraphRelationship,
  "sourceId" | "targetId" | "type"
>;

const UNKNOWN_NODE_COLOR = "#94a3b8";
const UNKNOWN_EDGE_COLOR = "#94a3b8";
const UNKNOWN_NODE_SIZE = 2;

// Node colors by type - slightly muted for less visual noise
export const NODE_COLORS: Record<DisplayNodeLabel, string> = {
  Project: "#a855f7", // Purple - prominent
  Package: "#8b5cf6", // Violet
  Module: "#7c3aed", // Violet darker
  Folder: "#6366f1", // Indigo
  File: "#3b82f6", // Blue
  Documentation: "#84cc16", // Lime - documentation center island
  Class: "#f59e0b", // Amber - stands out
  Function: "#10b981", // Emerald
  Method: "#14b8a6", // Teal
  Variable: "#64748b", // Slate - muted (less important)
  Interface: "#ec4899", // Pink
  Enum: "#f97316", // Orange
  Decorator: "#eab308", // Yellow
  Import: "#475569", // Slate darker - very muted
  Type: "#a78bfa", // Violet light
  CodeElement: "#64748b", // Slate - muted
  Community: "#818cf8", // Indigo light - cluster indicator
  Process: "#f43f5e", // Rose - execution flow indicator
  Section: "#60a5fa", // Blue light - structural section
  Struct: "#f59e0b", // Amber - like Class
  Trait: "#ec4899", // Pink - like Interface
  Impl: "#14b8a6", // Teal - like Method
  TypeAlias: "#a78bfa", // Violet light - like Type
  Const: "#64748b", // Slate - like Variable
  Static: "#64748b", // Slate - like Variable
  Namespace: "#7c3aed", // Violet - like Module
  Union: "#f97316", // Orange - like Enum
  Typedef: "#a78bfa", // Violet light - like Type
  Macro: "#eab308", // Yellow - like Decorator
  Property: "#64748b", // Slate - like Variable
  Record: "#f59e0b", // Amber - like Class
  Delegate: "#14b8a6", // Teal - like Method
  Annotation: "#eab308", // Yellow - like Decorator
  Constructor: "#10b981", // Emerald - like Function
  Template: "#a78bfa", // Violet light - like Type
  Route: "#f43f5e", // Rose - like Process
  Tool: "#a855f7", // Purple - like Project
};

// Node sizes by type - keep hierarchy visible without making metadata impossible to inspect.
export const NODE_SIZES: Record<DisplayNodeLabel, number> = {
  Project: 20, // Largest - root of everything
  Package: 16, // Major structural element
  Module: 13, // Important container
  Folder: 10, // Structural - clearly bigger than files
  File: 6, // Common element - smaller than folders
  Documentation: 5, // Documentation files - visible but below structural folders
  Class: 8, // Important code structure
  Function: 4, // Common code element - small
  Method: 3, // Smaller than function
  Variable: 2, // Tiny - leaf node
  Interface: 7, // Important type definition
  Enum: 5, // Type definition
  Decorator: 2, // Tiny modifier
  Import: 1.5, // Very small - usually hidden anyway
  Type: 3, // Type alias - small
  CodeElement: 2, // Generic small
  Community: 2, // Metadata node, default off but inspectable when toggled on
  Process: 4, // Execution-flow metadata, default off but inspectable when toggled on
  Section: 8, // Structural section - similar to Folder
  Struct: 8, // Like Class
  Trait: 7, // Like Interface
  Impl: 3, // Like Method
  TypeAlias: 3, // Like Type
  Const: 2, // Like Variable
  Static: 2, // Like Variable
  Namespace: 13, // Like Module
  Union: 5, // Like Enum
  Typedef: 3, // Like Type
  Macro: 2, // Like Decorator
  Property: 2, // Like Variable
  Record: 8, // Like Class
  Delegate: 3, // Like Method
  Annotation: 2, // Like Decorator
  Constructor: 4, // Like Function
  Template: 3, // Like Type
  Route: 5, // Like Enum
  Tool: 5, // Like Enum
};

export const getNodeColor = (label: string): string =>
  NODE_COLORS[label as DisplayNodeLabel] ?? UNKNOWN_NODE_COLOR;

export const getNodeSize = (label: string): number =>
  NODE_SIZES[label as DisplayNodeLabel] ?? UNKNOWN_NODE_SIZE;

// Community color palette for cluster-based coloring
export const COMMUNITY_COLORS = [
  "#ef4444", // red
  "#f97316", // orange
  "#eab308", // yellow
  "#22c55e", // green
  "#06b6d4", // cyan
  "#3b82f6", // blue
  "#8b5cf6", // violet
  "#d946ef", // fuchsia
  "#ec4899", // pink
  "#f43f5e", // rose
  "#14b8a6", // teal
  "#84cc16", // lime
];

export const getCommunityColor = (communityIndex: number): string => {
  return COMMUNITY_COLORS[communityIndex % COMMUNITY_COLORS.length];
};

// Labels to show by default (hide imports, variables, and metadata by default as they clutter)
export const DEFAULT_VISIBLE_LABELS: DisplayNodeLabel[] = [
  "Project",
  "Package",
  "Module",
  "Folder",
  "File",
  DOCUMENTATION_NODE_LABEL,
  "Class",
  "Function",
  "Method",
  "Interface",
  "Enum",
  "Type",
];

const withDocumentationLabel = (
  labels: readonly NodeLabel[],
): DisplayNodeLabel[] => {
  const result: DisplayNodeLabel[] = [...labels];
  const fileIndex = result.indexOf("File");
  result.splice(
    fileIndex >= 0 ? fileIndex + 1 : result.length,
    0,
    DOCUMENTATION_NODE_LABEL,
  );
  return result;
};

// All known filterable labels in generated graph-contract order, with the
// display-only Documentation filter kept next to File.
export const FILTERABLE_LABELS: DisplayNodeLabel[] =
  withDocumentationLabel(NODE_LABELS);

// Edge/Relation types in generated graph-contract order.
export const ALL_EDGE_TYPES: RelationshipType[] = [...GRAPH_RELATIONSHIP_TYPES];

// Keep all graph payload relationship types visible by default; users can disable noisy types.
export const DEFAULT_VISIBLE_EDGES: EdgeType[] = [...GRAPH_RELATIONSHIP_TYPES];

export const STRUCTURAL_NODE_LABELS: NodeLabel[] = [
  "Project",
  "Package",
  "Module",
  "Folder",
  "File",
  "Namespace",
  "Community",
  "Process",
  "Section",
  "Route",
  "Tool",
];

export const COMMUNITY_COLORED_NODE_LABELS: NodeLabel[] = [
  "Class",
  "Function",
  "Method",
  "Interface",
  "Struct",
  "Trait",
  "Impl",
  "Record",
  "Delegate",
  "Constructor",
];

export const EDGE_SIZE_MULTIPLIERS: Record<RelationshipType, number> = {
  CONTAINS: 0.4,
  CALLS: 0.8,
  INHERITS: 0.9,
  METHOD_OVERRIDES: 0.8,
  METHOD_IMPLEMENTS: 0.8,
  IMPORTS: 0.6,
  USES: 0.5,
  DEFINES: 0.5,
  DECORATES: 0.6,
  IMPLEMENTS: 0.9,
  EXTENDS: 1.0,
  HAS_METHOD: 0.6,
  HAS_PROPERTY: 0.6,
  ACCESSES: 0.5,
  MEMBER_OF: 0.3,
  STEP_IN_PROCESS: 0.7,
  HANDLES_ROUTE: 0.8,
  FETCHES: 0.7,
  HANDLES_TOOL: 0.8,
  ENTRY_POINT_OF: 0.8,
  WRAPS: 0.6,
  QUERIES: 0.7,
};

// Edge display info for UI
export const EDGE_INFO: Record<
  RelationshipType,
  { color: string; label: string; description?: string }
> = {
  CONTAINS: { color: "#2d5a3d", label: "Contains" },
  CALLS: { color: "#7c3aed", label: "Calls" },
  INHERITS: {
    color: "#c2410c",
    label: "Normalized Heritage",
    description:
      "Compatibility inheritance edge. When an EXTENDS or IMPLEMENTS edge exists for the same pair, the UI groups this edge instead of counting or drawing it as another source fact.",
  },
  METHOD_OVERRIDES: { color: "#ea580c", label: "Method Overrides" },
  METHOD_IMPLEMENTS: { color: "#be185d", label: "Method Implements" },
  IMPORTS: { color: "#1d4ed8", label: "Imports" },
  USES: { color: "#0891b2", label: "Uses" },
  DEFINES: { color: "#0e7490", label: "Defines" },
  DECORATES: { color: "#eab308", label: "Decorates" },
  IMPLEMENTS: { color: "#be185d", label: "Implements" },
  EXTENDS: { color: "#c2410c", label: "Extends" },
  HAS_METHOD: { color: "#0d9488", label: "Has Method" },
  HAS_PROPERTY: { color: "#64748b", label: "Has Property" },
  ACCESSES: { color: "#475569", label: "Accesses" },
  MEMBER_OF: { color: "#2563eb", label: "Member Of" },
  STEP_IN_PROCESS: { color: "#f43f5e", label: "Step In Process" },
  HANDLES_ROUTE: { color: "#db2777", label: "Handles Route" },
  FETCHES: { color: "#0284c7", label: "Fetches" },
  HANDLES_TOOL: { color: "#a855f7", label: "Handles Tool" },
  ENTRY_POINT_OF: { color: "#16a34a", label: "Entry Point Of" },
  WRAPS: { color: "#9333ea", label: "Wraps" },
  QUERIES: { color: "#ca8a04", label: "Queries" },
};

const nodeLabelOrder = new Map<string, number>(
  FILTERABLE_LABELS.map((label, index) => [label, index]),
);
const edgeTypeOrder = new Map<string, number>(
  GRAPH_RELATIONSHIP_TYPES.map((type, index) => [type, index]),
);

const compareKnownOrder =
  (order: Map<string, number>) =>
  (a: string, b: string): number => {
    const aOrder = order.get(a);
    const bOrder = order.get(b);
    if (aOrder !== undefined && bOrder !== undefined) return aOrder - bOrder;
    if (aOrder !== undefined) return -1;
    if (bOrder !== undefined) return 1;
    return a.localeCompare(b);
  };

const toDisplayLabel = (value: string): string =>
  value
    .split(/[_\s-]+/)
    .filter(Boolean)
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1).toLowerCase())
    .join(" ") || value;

export const getEdgeInfo = (type: string): { color: string; label: string } =>
  EDGE_INFO[type as RelationshipType] ?? {
    color: UNKNOWN_EDGE_COLOR,
    label: toDisplayLabel(type),
  };

const heritageCompatibilityTypes = new Set<string>(["EXTENDS", "IMPLEMENTS"]);

const relationshipPairKey = (relationship: RelationshipIdentity): string =>
  `${relationship.sourceId}\0${relationship.targetId}`;

export const getGroupedHeritageCompatibilityKeys = (
  relationships: readonly RelationshipIdentity[],
): Set<string> => {
  const canonicalPairs = new Set<string>();
  const inheritsPairs = new Set<string>();

  for (const relationship of relationships) {
    const key = relationshipPairKey(relationship);
    if (heritageCompatibilityTypes.has(relationship.type)) {
      canonicalPairs.add(key);
    } else if (relationship.type === "INHERITS") {
      inheritsPairs.add(key);
    }
  }

  const grouped = new Set<string>();
  for (const key of inheritsPairs) {
    if (canonicalPairs.has(key)) {
      grouped.add(key);
    }
  }
  return grouped;
};

export const isGroupedHeritageCompatibilityRelationship = (
  relationship: RelationshipIdentity,
  groupedKeys: ReadonlySet<string>,
): boolean =>
  relationship.type === "INHERITS" &&
  groupedKeys.has(relationshipPairKey(relationship));

export const getDisplayGraphRelationships = <T extends RelationshipIdentity>(
  relationships: readonly T[],
): T[] => {
  const groupedKeys = getGroupedHeritageCompatibilityKeys(relationships);
  return relationships.filter(
    (relationship) =>
      !isGroupedHeritageCompatibilityRelationship(relationship, groupedKeys),
  );
};

type DisplayLabelNode = { label: string } &
  Partial<Pick<GraphNode, "id">> & {
    properties?: Partial<GraphNode["properties"]>;
  };

const DOCUMENTATION_FILE_EXTENSIONS = new Set([
  ".md",
  ".mdx",
  ".rst",
  ".adoc",
  ".asciidoc",
]);

const DOCUMENTATION_FILE_NAMES = new Set([
  "agents",
  "changelog",
  "claude",
  "code_of_conduct",
  "contributing",
  "copying",
  "license",
  "notice",
  "readme",
  "security",
]);

const DOCUMENTATION_PATH_SEGMENTS = new Set([
  "doc",
  "docs",
  "documentation",
  "wiki",
]);

const toLowerPath = (value: unknown): string =>
  typeof value === "string" ? value.replace(/\\/g, "/").toLowerCase() : "";

const getPathBaseName = (path: string): string =>
  path.split("/").filter(Boolean).pop() ?? "";

const getFileExtension = (baseName: string): string => {
  const lastDotIndex = baseName.lastIndexOf(".");
  return lastDotIndex > 0 ? baseName.slice(lastDotIndex) : "";
};

const getFileStem = (baseName: string): string => {
  const lastDotIndex = baseName.lastIndexOf(".");
  return lastDotIndex > 0 ? baseName.slice(0, lastDotIndex) : baseName;
};

export const isDocumentationNode = (node: DisplayLabelNode): boolean => {
  if (node.label === DOCUMENTATION_NODE_LABEL) return true;

  const properties = node.properties ?? {};
  const path = toLowerPath(properties.filePath) || toLowerPath(properties.path);
  const name = toLowerPath(properties.name);
  const id = toLowerPath(node.id);
  const pathOrName = path || name || id;
  if (!pathOrName) return false;

  const pathSegments = pathOrName.split("/").filter(Boolean);
  if (pathSegments.some((segment) => DOCUMENTATION_PATH_SEGMENTS.has(segment))) {
    return true;
  }

  const baseName = getPathBaseName(pathOrName);
  const extension = getFileExtension(baseName);
  if (DOCUMENTATION_FILE_EXTENSIONS.has(extension)) return true;

  return DOCUMENTATION_FILE_NAMES.has(getFileStem(baseName));
};

export const getNodeDisplayLabel = (node: DisplayLabelNode): string =>
  isDocumentationNode(node) ? DOCUMENTATION_NODE_LABEL : node.label;

export const getNodeLabelCounts = (
  nodes: readonly DisplayLabelNode[],
): Map<string, number> => {
  const counts = new Map<string, number>();
  for (const node of nodes) {
    const label = getNodeDisplayLabel(node);
    counts.set(label, (counts.get(label) ?? 0) + 1);
  }
  return counts;
};

export const getRelationshipTypeCounts = (
  relationships: readonly Pick<GraphRelationship, "type">[],
): Map<string, number> => {
  const counts = new Map<string, number>();
  for (const relationship of relationships) {
    counts.set(relationship.type, (counts.get(relationship.type) ?? 0) + 1);
  }
  return counts;
};

export const getDisplayRelationshipTypeCounts = (
  relationships: readonly RelationshipIdentity[],
): Map<string, number> =>
  getRelationshipTypeCounts(getDisplayGraphRelationships(relationships));

export const getGroupedHeritageCompatibilityCount = (
  relationships: readonly RelationshipIdentity[],
  type: string,
): number => {
  if (type !== "INHERITS") return 0;
  const groupedKeys = getGroupedHeritageCompatibilityKeys(relationships);
  let count = 0;
  for (const relationship of relationships) {
    if (isGroupedHeritageCompatibilityRelationship(relationship, groupedKeys)) {
      count++;
    }
  }
  return count;
};

export const getFilterableNodeLabelsForGraph = (
  nodes: readonly DisplayLabelNode[],
): string[] =>
  Array.from(getNodeLabelCounts(nodes).keys()).sort(
    compareKnownOrder(nodeLabelOrder),
  );

export const getFilterableEdgeTypesForGraph = (
  relationships: readonly Pick<GraphRelationship, "type">[],
): string[] =>
  Array.from(getRelationshipTypeCounts(relationships).keys()).sort(
    compareKnownOrder(edgeTypeOrder),
  );
