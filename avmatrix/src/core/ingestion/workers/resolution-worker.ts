import { parentPort, workerData } from 'node:worker_threads';
import {
  createReferenceResolutionContext,
  deserializeScopeResolutionIndexes,
  resolveReferenceSiteChunk,
  type ReferenceResolutionChunk,
  type ReferenceResolutionChunkResult,
  type SerializedScopeResolutionIndexes,
} from '../scope-reference-resolver.js';

interface ResolutionWorkerData {
  readonly scopeResolutionIndexes?: SerializedScopeResolutionIndexes;
}

if (parentPort === null) {
  throw new Error('resolution-worker must run inside a worker thread');
}

const data = workerData as ResolutionWorkerData;
if (data.scopeResolutionIndexes === undefined) {
  throw new Error('resolution-worker missing scopeResolutionIndexes workerData');
}

const scopes = deserializeScopeResolutionIndexes(data.scopeResolutionIndexes);
const ctx = createReferenceResolutionContext(scopes);

parentPort.on('message', (message: { readonly type?: string; readonly files?: unknown }) => {
  if (message.type !== 'sub-batch') return;
  const chunks = Array.isArray(message.files) ? (message.files as ReferenceResolutionChunk[]) : [];
  const results: ReferenceResolutionChunkResult[] = [];

  try {
    for (const chunk of chunks) results.push(resolveReferenceSiteChunk(ctx, chunk));
    parentPort!.postMessage({
      type: 'result',
      data: results.length === 1 ? results[0] : results,
    });
  } catch (err) {
    parentPort!.postMessage({
      type: 'error',
      error: err instanceof Error ? err.message : String(err),
    });
  }
});
