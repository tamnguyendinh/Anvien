import {
  buildClassRegistry,
  buildFieldRegistry,
  buildMethodRegistry,
  type DefId,
  type Reference,
  type ReferenceIndex,
  type ReferenceSite,
  type RegistryContext,
  type Resolution,
} from 'avmatrix-shared';
import { performance } from 'node:perf_hooks';
import type { ScopeResolutionIndexes } from './model/scope-resolution-indexes.js';

const DEFAULT_REFERENCE_RESOLUTION_CHUNK_SIZE = 1000;

export interface ScopeReferenceResolutionOptions {
  readonly chunkSize?: number;
}

export interface ScopeReferenceResolutionStats {
  readonly totalReferenceSites: number;
  readonly chunkSize: number;
  readonly chunksResolved: number;
  readonly maxChunkReferenceSites: number;
  readonly resolvedReferences: number;
  readonly unresolvedReferences: number;
  readonly resolvedCalls: number;
  readonly resolvedAccesses: number;
  readonly resolvedTypeReferences: number;
  readonly resolvedInheritance: number;
  readonly resolvedImportUses: number;
  readonly referenceIndexSourceScopes: number;
  readonly referenceIndexTargetDefs: number;
}

export interface ScopeReferenceResolutionTimings {
  readonly referenceIndexBuildMs: number;
}

export interface ScopeReferenceResolutionResult {
  readonly referenceIndex: ReferenceIndex;
  readonly stats: ScopeReferenceResolutionStats;
  readonly timings: ScopeReferenceResolutionTimings;
}

export function resolveScopeReferenceSites(
  scopes: ScopeResolutionIndexes,
  options: ScopeReferenceResolutionOptions = {},
): ScopeReferenceResolutionResult {
  const ctx: RegistryContext = {
    scopes: scopes.scopeTree,
    defs: scopes.defs,
    qualifiedNames: scopes.qualifiedNames,
    moduleScopes: scopes.moduleScopes,
    methodDispatch: scopes.methodDispatch,
    providers: {},
  };

  const classRegistry = buildClassRegistry(ctx);
  const methodRegistry = buildMethodRegistry(ctx);
  const fieldRegistry = buildFieldRegistry(ctx);

  const refs: Reference[] = [];
  let unresolvedReferences = 0;
  let resolvedCalls = 0;
  let resolvedAccesses = 0;
  let resolvedTypeReferences = 0;
  let resolvedInheritance = 0;
  let resolvedImportUses = 0;

  const chunkSize = normalizeChunkSize(options.chunkSize);
  let chunksResolved = 0;
  let maxChunkReferenceSites = 0;

  for (let offset = 0; offset < scopes.referenceSites.length; offset += chunkSize) {
    const end = Math.min(offset + chunkSize, scopes.referenceSites.length);
    chunksResolved++;
    maxChunkReferenceSites = Math.max(maxChunkReferenceSites, end - offset);

    for (let index = offset; index < end; index++) {
      const site = scopes.referenceSites[index];
      if (site === undefined) continue;

      const resolution = bestResolutionForSite(site);
      if (resolution === undefined) {
        unresolvedReferences++;
        continue;
      }

      const fileHash = fileHashForSite(site);
      refs.push({
        fromScope: site.inScope,
        toDef: resolution.def.nodeId,
        ...(fileHash !== undefined ? { fileHash } : {}),
        atRange: site.atRange,
        kind: site.kind,
        confidence: resolution.confidence,
        evidence: resolution.evidence,
      });

      if (site.kind === 'call') resolvedCalls++;
      else if (site.kind === 'read' || site.kind === 'write') resolvedAccesses++;
      else if (site.kind === 'type-reference') resolvedTypeReferences++;
      else if (site.kind === 'inherits') resolvedInheritance++;
      else if (site.kind === 'import-use') resolvedImportUses++;
    }
  }

  const indexStart = performance.now();
  const referenceIndex = buildReferenceIndex(refs);
  const referenceIndexBuildMs = performance.now() - indexStart;

  return {
    referenceIndex,
    stats: {
      totalReferenceSites: scopes.referenceSites.length,
      chunkSize,
      chunksResolved,
      maxChunkReferenceSites,
      resolvedReferences: refs.length,
      unresolvedReferences,
      resolvedCalls,
      resolvedAccesses,
      resolvedTypeReferences,
      resolvedInheritance,
      resolvedImportUses,
      referenceIndexSourceScopes: referenceIndex.bySourceScope.size,
      referenceIndexTargetDefs: referenceIndex.byTargetDef.size,
    },
    timings: {
      referenceIndexBuildMs,
    },
  };

  function bestResolutionForSite(site: ReferenceSite): Resolution | undefined {
    if (site.kind === 'call') {
      if (site.callForm === 'constructor') {
        return (
          classRegistry.lookup(site.name, site.inScope)[0] ??
          methodRegistry.lookup(site.name, site.inScope, methodOptions(site))[0]
        );
      }
      return methodRegistry.lookup(site.name, site.inScope, methodOptions(site))[0];
    }

    if (site.kind === 'read' || site.kind === 'write') {
      return fieldRegistry.lookup(site.name, site.inScope, {
        ...(site.explicitReceiver !== undefined ? { explicitReceiver: site.explicitReceiver } : {}),
      })[0];
    }

    if (site.kind === 'type-reference' || site.kind === 'inherits') {
      return classRegistry.lookup(site.name, site.inScope)[0];
    }

    return (
      classRegistry.lookup(site.name, site.inScope)[0] ??
      methodRegistry.lookup(site.name, site.inScope, methodOptions(site))[0] ??
      fieldRegistry.lookup(site.name, site.inScope, {
        ...(site.explicitReceiver !== undefined ? { explicitReceiver: site.explicitReceiver } : {}),
      })[0]
    );
  }

  function fileHashForSite(site: ReferenceSite): string | undefined {
    const scope = scopes.scopeTree.getScope(site.inScope);
    if (scope === undefined) return undefined;
    return scopes.fileHashes.get(scope.filePath);
  }
}

function normalizeChunkSize(value: number | undefined): number {
  if (value === undefined) return DEFAULT_REFERENCE_RESOLUTION_CHUNK_SIZE;
  const chunkSize = Math.floor(value);
  return Number.isFinite(chunkSize) && chunkSize > 0
    ? chunkSize
    : DEFAULT_REFERENCE_RESOLUTION_CHUNK_SIZE;
}

function methodOptions(site: ReferenceSite) {
  return {
    ...(site.arity !== undefined ? { callsite: { arity: site.arity } } : {}),
    ...(site.explicitReceiver !== undefined ? { explicitReceiver: site.explicitReceiver } : {}),
  };
}

function buildReferenceIndex(refs: readonly Reference[]): ReferenceIndex {
  const bySourceScope = new Map<string, Reference[]>();
  const byTargetDef = new Map<DefId, Reference[]>();

  for (const ref of refs) {
    const sourceBucket = bySourceScope.get(ref.fromScope) ?? [];
    sourceBucket.push(ref);
    bySourceScope.set(ref.fromScope, sourceBucket);

    const targetBucket = byTargetDef.get(ref.toDef) ?? [];
    targetBucket.push(ref);
    byTargetDef.set(ref.toDef, targetBucket);
  }

  return {
    bySourceScope: freezeBuckets(bySourceScope),
    byTargetDef: freezeBuckets(byTargetDef),
  };
}

function freezeBuckets<K>(
  input: ReadonlyMap<K, readonly Reference[]>,
): ReadonlyMap<K, readonly Reference[]> {
  const out = new Map<K, readonly Reference[]>();
  for (const [key, refs] of input) out.set(key, Object.freeze([...refs]));
  return out;
}
