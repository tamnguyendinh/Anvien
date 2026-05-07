/**
 * Phase 5 of the RFC #909 ingestion lifecycle: drain `ReferenceIndex`
 * into the knowledge graph as labeled edges with `confidence` and
 * `evidence` properties (Ring 2 PKG #925).
 *
 * The resolution phase (future PR) writes `Reference` records into
 * `model.scopes.referenceSites`-derived `ReferenceIndex`; this module
 * materializes those records as `GraphRelationship`s via
 * `graph.addRelationship`. Every emitted edge carries:
 *
 *   - `type`: one of `'CALLS' | 'ACCESSES' | 'INHERITS' | 'USES'`
 *     (mapped from `Reference.kind` — `'read'` and `'write'` both route
 *     to `ACCESSES`; `'type-reference'` and `'import-use'` route to
 *     `USES`; `'call'` stays `CALLS`; `'inherits'` stays `INHERITS`).
 *   - `confidence`: the pre-computed confidence from the Reference record.
 *   - `reason`: human-readable summary (`"scope-resolution: call | confidence 0.75"`).
 *   - `evidence`: the full `ResolutionEvidence[]` trace — additive graph
 *     property (see `GraphRelationship.evidence` in avmatrix-shared),
 *     so queries that don't know about it are unaffected.
 *   - `step`: carries the reference's access-kind discriminant when
 *     available (`1` for read, `2` for write) so `ACCESSES` edges retain
 *     the read/write distinction without forcing a new edge type.
 *
 * ## Optional scope-tree flush
 *
 * When `INGESTION_EMIT_SCOPES=1` is set, this module also emits:
 *
 *   - `Scope` nodes for every `Scope` in the tree
 *   - `CONTAINS` edges from parent scope to child scope
 *   - `DEFINES` edges from scope to its `ownedDefs` members
 *   - `IMPORTS` edges from scope to `targetModuleScope` of each finalized
 *     `ImportEdge` that carries one
 *
 * Off by default — existing queries that don't know about `Scope` nodes
 * continue to work, and the storage cost is opt-in.
 *
 * ## Source-of-truth: the caller def for a reference
 *
 * A `Reference` says "some code inside `fromScope` references `toDef`".
 * The graph wants `(callerNodeId, calleeNodeId)`. We resolve the caller
 * by walking up the scope tree from `fromScope` until we find a scope
 * whose `ownedDefs` contains a Function-like def. If no such ancestor
 * exists, the edge is attributed to the first def owned by the innermost
 * ancestor scope, and if THAT produces nothing either the edge is
 * skipped (with a count returned in `EmitStats.skippedNoCaller`).
 */

import type {
  GraphNode,
  NodeLabel,
  GraphRelationship,
  ImportEdge,
  RelationshipType,
  Reference,
  ReferenceIndex,
  ResolutionEvidence,
  Scope,
  ScopeId,
  SymbolDefinition,
} from 'avmatrix-shared';
import type { KnowledgeGraph } from '../graph/types.js';
import type { ScopeResolutionIndexes } from './model/scope-resolution-indexes.js';
import { generateId } from '../../lib/utils.js';

// ─── Public API ─────────────────────────────────────────────────────────────

export interface EmitStats {
  readonly edgesEmitted: number;
  /** References dropped because no caller def could be resolved. */
  readonly skippedNoCaller: number;
  /** References dropped because `toDef` was not found in the DefIndex. */
  readonly skippedMissingTarget: number;
  /** References dropped because an equivalent graph edge already exists. */
  readonly skippedDuplicateEdge: number;
  /** Finalized file-level IMPORTS edges emitted from scope finalize output. */
  readonly finalizedImportEdgesEmitted: number;
  /** Finalized IMPORTS edges skipped because an equivalent edge already exists. */
  readonly skippedDuplicateImportEdge: number;
  /** Finalized per-symbol import-use USES edges emitted from scope finalize output. */
  readonly finalizedImportUseEdgesEmitted: number;
  /** Finalized import-use USES edges skipped because an equivalent edge already exists. */
  readonly skippedDuplicateImportUseEdge: number;
  /** Scope nodes emitted — `0` unless `INGESTION_EMIT_SCOPES=1`. */
  readonly scopeNodesEmitted: number;
  /** Scope-tree structural edges emitted — `0` unless `INGESTION_EMIT_SCOPES=1`. */
  readonly scopeEdgesEmitted: number;
}

export interface EmitReferencesInput {
  readonly graph: KnowledgeGraph;
  readonly scopes: ScopeResolutionIndexes;
  readonly referenceIndex: ReferenceIndex;
  /** Human-consumable label for the `reason` prefix. Defaults to `'scope-resolution'`. */
  readonly sourceLabel?: string;
}

/**
 * Drain `referenceIndex.bySourceScope` into graph edges.
 *
 * The scope-tree flush is controlled separately by
 * `INGESTION_EMIT_SCOPES` — callers can run `emitReferencesToGraph`
 * without scope-node emission or layer the two calls as needed.
 */
export function emitReferencesToGraph(input: EmitReferencesInput): EmitStats {
  const { graph, scopes, referenceIndex } = input;
  const sourceLabel = input.sourceLabel ?? 'scope-resolution';

  let edgesEmitted = 0;
  let skippedNoCaller = 0;
  let skippedMissingTarget = 0;
  let skippedDuplicateEdge = 0;
  const existingEdges = buildExistingEdgeKeyMap(graph);
  const graphNodeResolver = createGraphNodeResolver(graph);

  for (const [fromScope, refs] of referenceIndex.bySourceScope) {
    for (const ref of refs) {
      const targetDef = scopes.defs.get(ref.toDef);
      if (targetDef === undefined) {
        skippedMissingTarget++;
        continue;
      }
      const callerDef = resolveCallerDef(fromScope, scopes);
      if (callerDef === undefined) {
        skippedNoCaller++;
        continue;
      }
      const callerId = graphNodeResolver(callerDef);
      const targetId = graphNodeResolver(targetDef);
      if (callerId === undefined) {
        skippedNoCaller++;
        continue;
      }
      if (targetId === undefined) {
        skippedMissingTarget++;
        continue;
      }
      const mappedRelationship = buildRelationship(ref, callerId, targetId, sourceLabel);
      const edgeKey = semanticEdgeKey(mappedRelationship);
      const existing = existingEdges.get(edgeKey);
      if (existing !== undefined) {
        mergeRelationshipAudit(graph, existing, mappedRelationship);
        existingEdges.set(edgeKey, mergedRelationship(existing, mappedRelationship));
        skippedDuplicateEdge++;
        continue;
      }
      graph.addRelationship(mappedRelationship);
      existingEdges.set(edgeKey, mappedRelationship);
      edgesEmitted++;
    }
  }

  const scopeStats = isScopeEmissionEnabled()
    ? emitScopeGraph({ graph, scopes })
    : { scopeNodesEmitted: 0, scopeEdgesEmitted: 0 };
  const importStats = emitFinalizedFileImports({ graph, scopes, existingEdges, graphNodeResolver });

  return {
    edgesEmitted,
    skippedNoCaller,
    skippedMissingTarget,
    skippedDuplicateEdge,
    ...importStats,
    ...scopeStats,
  };
}

/**
 * Emit `Scope` nodes + `CONTAINS`/`DEFINES`/`IMPORTS` edges representing
 * the lexical scope tree itself. Skipped unless `INGESTION_EMIT_SCOPES=1`
 * at the public entry point; exported here for tests that want to
 * exercise the path directly.
 */
export function emitScopeGraph(input: {
  readonly graph: KnowledgeGraph;
  readonly scopes: ScopeResolutionIndexes;
}): { readonly scopeNodesEmitted: number; readonly scopeEdgesEmitted: number } {
  const { graph, scopes } = input;
  let scopeNodesEmitted = 0;
  let scopeEdgesEmitted = 0;

  for (const scope of scopes.scopeTree.byId.values()) {
    graph.addNode({
      id: scope.id,
      label: 'CodeElement' as NodeLabel, // the generic bucket for non-symbol graph nodes
      properties: {
        name: scope.kind,
        filePath: scope.filePath,
        startLine: scope.range.startLine,
        endLine: scope.range.endLine,
        description: `Scope: ${scope.kind}`,
      } as unknown as Parameters<KnowledgeGraph['addNode']>[0]['properties'],
    });
    scopeNodesEmitted++;

    if (scope.parent !== null) {
      graph.addRelationship({
        id: `rel:contains:${scope.parent}->${scope.id}`,
        sourceId: scope.parent,
        targetId: scope.id,
        type: 'CONTAINS',
        confidence: 1,
        reason: 'scope-tree parent/child',
      });
      scopeEdgesEmitted++;
    }

    for (const def of scope.ownedDefs) {
      graph.addRelationship({
        id: `rel:defines:${scope.id}->${def.nodeId}`,
        sourceId: scope.id,
        targetId: def.nodeId,
        type: 'DEFINES',
        confidence: 1,
        reason: 'scope.ownedDefs',
      });
      scopeEdgesEmitted++;
    }
  }

  for (const [scopeId, edges] of scopes.imports) {
    for (const edge of edges) {
      if (edge.targetModuleScope === undefined) continue;
      const fileHash = fileHashForScope(scopeId, scopes);
      graph.addRelationship({
        id: `rel:imports:${scopeId}->${edge.targetModuleScope}:${edge.localName}`,
        sourceId: scopeId,
        targetId: edge.targetModuleScope,
        type: 'IMPORTS',
        confidence: edge.linkStatus === 'unresolved' ? 0.5 : 1,
        reason: `import ${edge.kind} ${edge.localName}`,
        resolutionSource: 'scope-finalize',
        ...(fileHash !== undefined ? { fileHash } : {}),
        evidence: importEvidence(edge),
      });
      scopeEdgesEmitted++;
    }
  }

  return { scopeNodesEmitted, scopeEdgesEmitted };
}

function emitFinalizedFileImports(input: {
  readonly graph: KnowledgeGraph;
  readonly scopes: ScopeResolutionIndexes;
  readonly existingEdges: Map<string, GraphRelationship>;
  readonly graphNodeResolver: (def: SymbolDefinition) => string | undefined;
}): {
  readonly finalizedImportEdgesEmitted: number;
  readonly skippedDuplicateImportEdge: number;
  readonly finalizedImportUseEdgesEmitted: number;
  readonly skippedDuplicateImportUseEdge: number;
} {
  const { graph, scopes, existingEdges, graphNodeResolver } = input;
  let finalizedImportEdgesEmitted = 0;
  let skippedDuplicateImportEdge = 0;
  let finalizedImportUseEdgesEmitted = 0;
  let skippedDuplicateImportUseEdge = 0;

  for (const [scopeId, edges] of scopes.imports) {
    const sourceScope = scopes.scopeTree.getScope(scopeId);
    if (sourceScope === undefined) continue;
    for (const edge of edges) {
      if (edge.targetFile === null || edge.linkStatus === 'unresolved') continue;
      const fileHash = scopes.fileHashes.get(sourceScope.filePath);
      const importRelationship = {
        id: generateId('IMPORTS', `${sourceScope.filePath}->${edge.targetFile}`),
        sourceId: generateId('File', sourceScope.filePath),
        targetId: generateId('File', edge.targetFile),
        type: 'IMPORTS' as const,
        confidence: 1,
        reason: `scope-finalize import ${edge.kind} ${edge.localName}`,
        resolutionSource: 'scope-finalize',
        ...(fileHash !== undefined ? { fileHash } : {}),
        evidence: importEvidence(edge),
      };
      const importEdgeKey = semanticEdgeKey(importRelationship);
      const existingImport = existingEdges.get(importEdgeKey);
      if (existingImport !== undefined) {
        mergeRelationshipAudit(graph, existingImport, importRelationship);
        existingEdges.set(importEdgeKey, mergedRelationship(existingImport, importRelationship));
        skippedDuplicateImportEdge++;
      } else {
        graph.addRelationship(importRelationship);
        existingEdges.set(importEdgeKey, importRelationship);
        finalizedImportEdgesEmitted++;
      }

      if (edge.targetDefId === undefined || scopes.defs.get(edge.targetDefId) === undefined) {
        continue;
      }
      const targetDef = scopes.defs.get(edge.targetDefId);
      if (targetDef === undefined) continue;
      const mappedTargetId = graphNodeResolver(targetDef);
      if (mappedTargetId === undefined) continue;
      const importUseRelationship = {
        id: generateId(
          'USES',
          `${sourceScope.filePath}->${edge.targetDefId}:import:${edge.localName}`,
        ),
        sourceId: generateId('File', sourceScope.filePath),
        targetId: mappedTargetId,
        type: 'USES' as const,
        confidence: 1,
        reason: `scope-finalize import-use ${edge.kind} ${edge.localName}`,
        resolutionSource: 'scope-finalize',
        ...(fileHash !== undefined ? { fileHash } : {}),
        evidence: importEvidence(edge),
      };
      const importUseEdgeKey = semanticEdgeKey(importUseRelationship);
      const existingImportUse = existingEdges.get(importUseEdgeKey);
      if (existingImportUse !== undefined) {
        mergeRelationshipAudit(graph, existingImportUse, importUseRelationship);
        existingEdges.set(
          importUseEdgeKey,
          mergedRelationship(existingImportUse, importUseRelationship),
        );
        skippedDuplicateImportUseEdge++;
        continue;
      }
      graph.addRelationship(importUseRelationship);
      existingEdges.set(importUseEdgeKey, importUseRelationship);
      finalizedImportUseEdgesEmitted++;
    }
  }

  return {
    finalizedImportEdgesEmitted,
    skippedDuplicateImportEdge,
    finalizedImportUseEdgesEmitted,
    skippedDuplicateImportUseEdge,
  };
}

function fileHashForScope(scopeId: ScopeId, scopes: ScopeResolutionIndexes): string | undefined {
  const scope = scopes.scopeTree.getScope(scopeId);
  if (scope === undefined) return undefined;
  return scopes.fileHashes.get(scope.filePath);
}

function importEvidence(edge: ImportEdge): GraphRelationship['evidence'] {
  const target = edge.targetFile ?? edge.targetExportedName;
  return [
    {
      kind: 'import',
      weight: edge.linkStatus === 'unresolved' ? 0.5 : 1,
      note: `${edge.kind} ${edge.localName} -> ${target}`,
    },
  ];
}

// ─── Internal ───────────────────────────────────────────────────────────────

/** Accepted truthy values for `INGESTION_EMIT_SCOPES`. */
const TRUTHY: ReadonlySet<string> = new Set(['true', '1', 'yes']);

function isScopeEmissionEnabled(): boolean {
  const raw = process.env['INGESTION_EMIT_SCOPES'];
  if (raw === undefined) return false;
  return TRUTHY.has(raw.trim().toLowerCase());
}

/**
 * Walk up from `startScope` looking for the first ancestor scope whose
 * `ownedDefs` contains a Function-like def (Function / Method /
 * Constructor). Fall back to the innermost ancestor's first `ownedDef`
 * if none is found; return `undefined` if all ancestors have no defs.
 */
function resolveCallerDef(
  startScope: ScopeId,
  scopes: ScopeResolutionIndexes,
): SymbolDefinition | undefined {
  const tree = scopes.scopeTree;
  let current: ScopeId | null = startScope;
  const visited = new Set<ScopeId>();
  let firstOwnedFallback: SymbolDefinition | undefined;

  while (current !== null) {
    if (visited.has(current)) break;
    visited.add(current);

    const scope: Scope | undefined = tree.getScope(current);
    if (scope === undefined) break;

    // Prefer a Function-like owner.
    const fnDef = scope.ownedDefs.find((d) => isFunctionLike(d.type));
    if (fnDef !== undefined) return fnDef;

    // Stash the first owned def we see as a conservative fallback.
    if (firstOwnedFallback === undefined && scope.ownedDefs.length > 0) {
      firstOwnedFallback = scope.ownedDefs[0]!;
    }

    current = scope.parent;
  }

  return firstOwnedFallback;
}

function isFunctionLike(type: NodeLabel): boolean {
  return type === 'Function' || type === 'Method' || type === 'Constructor';
}

function createGraphNodeResolver(
  graph: KnowledgeGraph,
): (def: SymbolDefinition) => string | undefined {
  const bySemanticKey = new Map<string, string[]>();
  graph.forEachNode((node) => {
    for (const key of graphNodeSemanticKeys(node)) {
      const bucket = bySemanticKey.get(key) ?? [];
      bucket.push(node.id);
      bySemanticKey.set(key, bucket);
    }
  });

  return (def: SymbolDefinition): string | undefined => {
    if (graph.getNode(def.nodeId) !== undefined) return def.nodeId;
    for (const key of defSemanticKeys(def)) {
      const matches = bySemanticKey.get(key);
      if (matches !== undefined && matches.length === 1) return matches[0]!;
    }
    return undefined;
  };
}

function graphNodeSemanticKeys(node: GraphNode): readonly string[] {
  const filePath = stringProperty(node.properties.filePath);
  if (filePath === undefined) return [];
  const qualifiedName = stringProperty(node.properties.qualifiedName);
  const idName = graphNodeIdName(node);
  const directNames = uniqueStrings([
    stringProperty(node.properties.name),
    qualifiedName,
    idName,
    stripArityTag(idName),
    simpleName(qualifiedName),
    simpleName(stripArityTag(idName)),
  ]);
  const aliasNames = uniqueStrings([qualifiedName, idName, stripArityTag(idName)]);
  return [
    ...directNames.map((name) => semanticNodeKey(node.label, filePath, name)),
    ...semanticNodeLabels(node.label)
      .filter((label) => label !== node.label)
      .flatMap((label) => aliasNames.map((name) => semanticNodeKey(label, filePath, name))),
  ];
}

function defSemanticKeys(def: SymbolDefinition): readonly string[] {
  const directNames = uniqueStrings([
    def.qualifiedName,
    stripArityTag(def.qualifiedName),
    simpleName(def.qualifiedName),
  ]);
  const aliasNames = uniqueStrings([def.qualifiedName, stripArityTag(def.qualifiedName)]);
  return [
    ...directNames.map((name) => semanticNodeKey(def.type, def.filePath, name)),
    ...semanticNodeLabels(def.type)
      .filter((label) => label !== def.type)
      .flatMap((label) => aliasNames.map((name) => semanticNodeKey(label, def.filePath, name))),
  ];
}

function semanticNodeLabels(label: NodeLabel): readonly NodeLabel[] {
  if (label === 'Method') return ['Method', 'Function'];
  if (label === 'Function') return ['Function', 'Method'];
  return [label];
}

function graphNodeIdName(node: GraphNode): string | undefined {
  const filePath = stringProperty(node.properties.filePath);
  if (filePath === undefined) return undefined;
  const prefix = `${node.label}:${filePath}:`;
  if (!node.id.startsWith(prefix)) return undefined;
  const name = node.id.slice(prefix.length);
  return name.length > 0 ? name : undefined;
}

function stripArityTag(value: string | undefined): string | undefined {
  if (value === undefined) return undefined;
  const hash = value.indexOf('#');
  return hash === -1 ? value : value.slice(0, hash);
}

function semanticNodeKey(label: NodeLabel, filePath: string, name: string): string {
  return `${label}\0${filePath}\0${name}`;
}

function simpleName(value: string | undefined): string | undefined {
  if (value === undefined || value.length === 0) return undefined;
  const dot = value.lastIndexOf('.');
  return dot === -1 ? value : value.slice(dot + 1);
}

function stringProperty(value: unknown): string | undefined {
  return typeof value === 'string' && value.length > 0 ? value : undefined;
}

function uniqueStrings(values: readonly (string | undefined)[]): string[] {
  const out: string[] = [];
  const seen = new Set<string>();
  for (const value of values) {
    if (value === undefined || seen.has(value)) continue;
    seen.add(value);
    out.push(value);
  }
  return out;
}

function buildRelationship(
  ref: Reference,
  callerId: string,
  targetId: string,
  sourceLabel: string,
): Parameters<KnowledgeGraph['addRelationship']>[0] {
  const type = mapKindToType(ref.kind);
  const reason =
    ref.kind === 'read' || ref.kind === 'write'
      ? ref.kind
      : `${sourceLabel}: ${ref.kind} | confidence ${ref.confidence.toFixed(3)}`;
  // `step` encodes read/write discriminator for ACCESSES edges (1=read, 2=write).
  // Other kinds omit `step`.
  const step = ref.kind === 'read' ? 1 : ref.kind === 'write' ? 2 : undefined;
  return {
    id: `rel:${type}:${callerId}->${targetId}:${ref.atRange.startLine}:${ref.atRange.startCol}`,
    sourceId: callerId,
    targetId,
    type,
    confidence: ref.confidence,
    reason,
    resolutionSource: sourceLabel,
    ...(ref.fileHash !== undefined ? { fileHash: ref.fileHash } : {}),
    evidence: ref.evidence.map(serializeEvidence),
    ...(step !== undefined ? { step } : {}),
  };
}

/**
 * Map a `Reference.kind` to an existing `RelationshipType`. Read/write
 * both fold into `ACCESSES`; `type-reference` + `import-use` both fold
 * into `USES`. This keeps the graph schema additive — no new
 * RelationshipType values are introduced by this module.
 */
function mapKindToType(kind: Reference['kind']): RelationshipType {
  switch (kind) {
    case 'call':
      return 'CALLS';
    case 'read':
    case 'write':
      return 'ACCESSES';
    case 'inherits':
      return 'INHERITS';
    case 'type-reference':
    case 'import-use':
      return 'USES';
  }
}

function serializeEvidence(e: ResolutionEvidence): {
  readonly kind: string;
  readonly weight: number;
  readonly note?: string;
} {
  return {
    kind: e.kind,
    weight: e.weight,
    ...(e.note !== undefined ? { note: e.note } : {}),
  };
}

function buildExistingEdgeKeyMap(graph: KnowledgeGraph): Map<string, GraphRelationship> {
  const keys = new Map<string, GraphRelationship>();
  graph.forEachRelationship((relationship) => {
    keys.set(semanticEdgeKey(relationship), relationship);
  });
  return keys;
}

function mergeRelationshipAudit(
  graph: KnowledgeGraph,
  existing: GraphRelationship,
  incoming: GraphRelationship,
): void {
  const merged = mergedRelationship(existing, incoming);
  graph.removeRelationship(existing.id);
  graph.addRelationship(merged);
}

function mergedRelationship(
  existing: GraphRelationship,
  incoming: GraphRelationship,
): GraphRelationship {
  const existingConfidence = existing.confidence ?? 0;
  const incomingConfidence = incoming.confidence ?? 0;
  const useIncomingReason =
    existing.resolutionSource === undefined || incomingConfidence >= existingConfidence;
  return {
    ...existing,
    confidence: Math.max(existingConfidence, incomingConfidence),
    reason: useIncomingReason ? incoming.reason : existing.reason,
    ...(incoming.step !== undefined && existing.step === undefined ? { step: incoming.step } : {}),
    ...(incoming.resolutionSource !== undefined
      ? { resolutionSource: incoming.resolutionSource }
      : {}),
    ...(incoming.fileHash !== undefined ? { fileHash: incoming.fileHash } : {}),
    ...(incoming.evidence !== undefined && incoming.evidence.length > 0
      ? { evidence: incoming.evidence }
      : {}),
  };
}

function semanticEdgeKey(relationship: GraphRelationship): string {
  const accessKind =
    relationship.type === 'ACCESSES' ? `:${accessKindDiscriminator(relationship)}` : '';
  return `${relationship.sourceId}\0${relationship.targetId}\0${relationship.type}${accessKind}`;
}

function accessKindDiscriminator(relationship: GraphRelationship): string {
  if (relationship.step === 1) return 'read';
  if (relationship.step === 2) return 'write';

  const reason = relationship.reason.trim().toLowerCase();
  if (reason === 'read' || reason.includes(': read |')) return 'read';
  if (reason === 'write' || reason.includes(': write |')) return 'write';
  return 'unknown';
}
