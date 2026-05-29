/**
 * Consolidated HTTP client for the anvien backend server.
 *
 * Replaces backend.ts, server-connection.ts, and worker HTTP helpers
 * with a single typed module. All graph queries, search, embeddings,
 * and file operations go through this client.
 */

import type {
  GraphNode,
  GraphRelationship,
  GraphResponse,
  GraphSemanticStatus,
} from '@/generated/anvien-contracts';
import {
  recordHeartbeatConnect,
  recordHeartbeatReconnect,
} from '../lib/runtime-diagnostics';

// ── Types ──────────────────────────────────────────────────────────────────

export interface BackendRepo {
  name: string;
  path: string;
  repoPath?: string; // git HEAD returns "repoPath"; older versions return "path"
  indexedAt: string;
  lastCommit?: string;
  stats?: {
    files?: number;
    nodes?: number;
    edges?: number;
    communities?: number;
    processes?: number;
  };
}

export interface EnrichedSearchResult {
  filePath: string;
  score: number;
  rank?: number;
  sources?: string[];
  nodeId?: string;
  name?: string;
  label?: string;
  startLine?: number;
  endLine?: number;
  // Enrichment (server-side)
  connections?: {
    outgoing: Array<{ name: string; type: string; confidence?: number }>;
    incoming: Array<{ name: string; type: string; confidence?: number }>;
  };
  cluster?: string;
  processes?: Array<{ id: string; label: string; step?: number; stepCount?: number }>;
}

export interface GrepResult {
  filePath: string;
  line: number;
  text: string;
}

export interface JobProgress {
  phase: string;
  percent: number;
  message: string;
}

export interface JobStatus {
  id: string;
  status: 'queued' | 'analyzing' | 'loading' | 'complete' | 'failed';
  repoPath?: string;
  repoName?: string;
  progress: JobProgress;
  error?: string;
  startedAt: number;
  completedAt?: number;
}

export interface AnalyzeCompleteData {
  repoName?: string;
  repoPath?: string;
  error?: string;
}

export class BackendError extends Error {
  constructor(
    message: string,
    public readonly status: number,
    public readonly code: 'network' | 'server' | 'client' | 'not_found',
  ) {
    super(message);
    this.name = 'BackendError';
  }
}

// ── SSE Utility ────────────────────────────────────────────────────────────

export interface SSEHandlers<T = unknown> {
  onMessage?: (data: T) => void;
  onComplete?: (data: T) => void;
  onError?: (error: string) => void;
}

/**
 * Generic SSE stream consumer using fetch + ReadableStream.
 * Returns an AbortController to cancel the stream.
 */
export function streamSSE<T = unknown>(url: string, handlers: SSEHandlers<T>): AbortController {
  const controller = new AbortController();
  let lastEventId = '';

  const connect = () => {
    if (controller.signal.aborted) return;

    (async () => {
      try {
        const headers: Record<string, string> = {};
        if (lastEventId) {
          headers['Last-Event-ID'] = lastEventId;
        }

        const response = await fetch(url, { signal: controller.signal, headers });
        if (!response.ok) {
          handlers.onError?.(`Server returned ${response.status}`);
          return;
        }

        const reader = response.body?.getReader();
        if (!reader) {
          handlers.onError?.('No response body');
          return;
        }

        const decoder = new TextDecoder();
        let buffer = '';

        while (true) {
          const { done, value } = await reader.read();
          if (done) break;

          buffer += decoder.decode(value, { stream: true });
          const lines = buffer.split('\n');
          buffer = lines.pop() || '';

          let eventType = 'message';
          for (const line of lines) {
            if (line.startsWith('id: ')) {
              lastEventId = line.slice(4).trim();
              continue;
            }
            if (line.startsWith(':')) {
              // SSE comment (heartbeat) — skip
              continue;
            }
            if (line.startsWith('event: ')) {
              eventType = line.slice(7).trim();
            } else if (line.startsWith('data: ')) {
              try {
                const parsed = JSON.parse(line.slice(6)) as T;
                if (eventType === 'complete') {
                  handlers.onComplete?.(parsed);
                  return;
                } else if (eventType === 'failed') {
                  const errData = parsed as any;
                  handlers.onError?.(errData?.error || 'Job failed');
                  return;
                } else {
                  handlers.onMessage?.(parsed);
                }
              } catch {
                // Skip malformed JSON
              }
              eventType = 'message';
            }
          }
        }

        if (!controller.signal.aborted) {
          handlers.onError?.('Stream ended before completion');
        }
      } catch (err: unknown) {
        if (err instanceof DOMException && err.name === 'AbortError') return;
        if (!controller.signal.aborted) {
          handlers.onError?.(err instanceof Error ? err.message : 'Stream error');
        }
      }
    })();
  };

  connect();
  return controller;
}

// ── Configuration ──────────────────────────────────────────────────────────

let _backendUrl = 'http://127.0.0.1:4848';

const LOCAL_BACKEND_ERROR =
  'Anvien local-only mode only supports backend URLs on 127.0.0.1, localhost, or [::1].';

const isLoopbackHostname = (hostname: string): boolean => {
  const normalized = hostname.replace(/^\[|\]$/g, '');
  return normalized === 'localhost' || normalized === '127.0.0.1' || normalized === '::1';
};

export const setBackendUrl = (url: string): void => {
  _backendUrl = normalizeServerUrl(url);
};

export const getBackendUrl = (): string => _backendUrl;

/**
 * Normalize a user-entered server URL into a loopback-only base URL suitable for
 * setBackendUrl(). Adds protocol if missing, strips trailing slashes, and strips a
 * trailing /api suffix (since all API methods append their own /api/... paths to
 * _backendUrl). Remote hosts are rejected in local-only mode.
 */
export function normalizeServerUrl(input: string): string {
  let url = input.trim().replace(/\/+$/, '');

  if (!url.startsWith('http://') && !url.startsWith('https://')) {
    if (url.startsWith('localhost') || url.startsWith('127.0.0.1') || url.startsWith('[::1]')) {
      url = `http://${url}`;
    } else {
      throw new Error(LOCAL_BACKEND_ERROR);
    }
  }

  const parsed = new URL(url);
  if (!isLoopbackHostname(parsed.hostname)) {
    throw new Error(LOCAL_BACKEND_ERROR);
  }

  const normalizedPath = parsed.pathname.replace(/\/+$/, '');
  if (normalizedPath !== '' && normalizedPath !== '/' && normalizedPath !== '/api') {
    throw new Error(
      'Anvien local-only mode expects the backend URL to point at the local server root or /api.',
    );
  }

  const canonicalHost =
    parsed.hostname === 'localhost'
      ? `127.0.0.1${parsed.port ? `:${parsed.port}` : ''}`
      : parsed.host;
  return `${parsed.protocol}//${canonicalHost}`;
}

// ── Internal Helpers ───────────────────────────────────────────────────────

const fetchFromBackend = async (
  url: string,
  init: RequestInit = {},
): Promise<Response> => {
  try {
    const response = await fetch(url, init);
    return response;
  } catch (error: unknown) {
    if (error instanceof DOMException && error.name === 'AbortError') {
      throw new BackendError('Request aborted', 0, 'network');
    }
    if (error instanceof TypeError) {
      throw new BackendError(
        `Network error reaching Anvien backend at ${_backendUrl}: ${error.message}`,
        0,
        'network',
      );
    }
    throw error;
  }
};

const assertOk = async (response: Response): Promise<void> => {
  if (response.ok) return;

  let message = response.statusText;
  try {
    const body = await response.json();
    if (body && typeof body.error === 'string') {
      message = body.error;
    } else if (body && typeof body.message === 'string') {
      message = body.message;
    }
  } catch {
    // Response body was not JSON
  }

  const code =
    response.status === 404
      ? 'not_found'
      : response.status >= 400 && response.status < 500
        ? 'client'
        : 'server';
  throw new BackendError(message, response.status, code);
};

const repoParam = (repo?: string): string => (repo ? `repo=${encodeURIComponent(repo)}` : '');

// ── API Methods ────────────────────────────────────────────────────────────

/** Server info from /api/info. */
export interface ServerInfo {
  version: string;
  launchContext: 'npx' | 'global' | 'local';
  nodeVersion: string;
}

/** Fetch server info (version, launch context). */
export const fetchServerInfo = async (): Promise<ServerInfo> => {
  const response = await fetchFromBackend(`${_backendUrl}/api/info`);
  await assertOk(response);
  return response.json() as Promise<ServerInfo>;
};

/**
 * Connect an SSE heartbeat to the backend. The browser EventSource owns
 * reconnect behavior; the app only records/displays state changes.
 *
 * - `onConnect` fires on every successful (re)connection.
 * - `onReconnecting` fires on the first observed drop — use it to show
 *   a "reconnecting" banner while keeping the current view intact.
 *
 * Returns a cleanup function that tears down the EventSource.
 */
export const connectHeartbeat = (
  onConnect: () => void,
  onReconnecting: () => void,
): (() => void) => {
  let closed = false;
  let es: EventSource | null = null;
  /** Whether we've already fired onReconnecting for the current drop. */
  let notifiedReconnecting = false;

  const markConnected = () => {
    notifiedReconnecting = false;
    recordHeartbeatConnect();
    onConnect();
  };

  const connect = () => {
    if (closed) return;
    es = new EventSource(`${_backendUrl}/api/heartbeat`);
    es.onopen = () => {
      if (!closed) {
        markConnected();
      }
    };
    es.onerror = () => {
      if (closed) return;

      if (!notifiedReconnecting) {
        notifiedReconnecting = true;
        recordHeartbeatReconnect(1);
        onReconnecting();
      }
    };
  };

  connect();

  return () => {
    closed = true;
    es?.close();
  };
};

/** Delete a repo's index and unregister it. */
export const deleteRepo = async (repoName: string): Promise<void> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/repo?repo=${encodeURIComponent(repoName)}`,
    {
      method: 'DELETE',
    },
  );
  await assertOk(response);
};

/** Probe the backend. Returns true if reachable. */
export const probeBackend = async (): Promise<boolean> => {
  try {
    const response = await fetchFromBackend(`${_backendUrl}/api/repos`);
    return response.status === 200;
  } catch {
    return false;
  }
};

/** Fetch list of indexed repositories. */
export const fetchRepos = async (): Promise<BackendRepo[]> => {
  const response = await fetchFromBackend(`${_backendUrl}/api/repos`);
  await assertOk(response);
  return response.json() as Promise<BackendRepo[]>;
};

export const fetchRepoInfo = async (
  repo?: string,
  opts?: { awaitAnalysis?: boolean },
): Promise<BackendRepo> => {
  const params = [repoParam(repo), opts?.awaitAnalysis ? 'awaitAnalysis=true' : '']
    .filter(Boolean)
    .join('&');
  const url = `${_backendUrl}/api/repo${params ? `?${params}` : ''}`;
  const response = await fetchFromBackend(url);
  await assertOk(response);
  const data = await response.json();
  return { ...data, repoPath: data.repoPath ?? data.path };
};

/** Fetch the graph (nodes + relationships). Content stripped by default. */
export const fetchGraph = async (
  repo?: string,
  opts?: {
    includeContent?: boolean;
    signal?: AbortSignal;
    onProgress?: (downloaded: number, total: number | null) => void;
  },
): Promise<GraphResponse> => {
  const params = [repoParam(repo), opts?.includeContent ? 'includeContent=true' : '', 'stream=true']
    .filter(Boolean)
    .join('&');
  const url = `${_backendUrl}/api/graph${params ? `?${params}` : ''}`;
  const response = await fetchFromBackend(url, { signal: opts?.signal });
  await assertOk(response);

  const contentType = response.headers.get('Content-Type') || '';
  if (contentType.includes('application/x-ndjson')) {
    return parseNdjsonGraphResponse(response, opts?.onProgress);
  }

  if (!opts?.onProgress || !response.body) {
    return response.json() as Promise<GraphResponse>;
  }

  // Streaming download with progress
  const contentLength = response.headers.get('Content-Length');
  const total = contentLength ? parseInt(contentLength, 10) : null;
  const reader = response.body.getReader();
  const chunks: Uint8Array[] = [];
  let downloaded = 0;

  while (true) {
    const { done, value } = await reader.read();
    if (done) break;
    chunks.push(value);
    downloaded += value.length;
    opts.onProgress(downloaded, total);
  }

  const combined = new Uint8Array(downloaded);
  let offset = 0;
  for (const chunk of chunks) {
    combined.set(chunk, offset);
    offset += chunk.length;
  }
  return JSON.parse(new TextDecoder().decode(combined));
};

const parseNdjsonGraphResponse = async (
  response: Response,
  onProgress?: (downloaded: number, total: number | null) => void,
): Promise<GraphResponse> => {
  if (!response.body) {
    throw new BackendError('No response body', response.status, 'server');
  }

  const contentLength = response.headers.get('Content-Length');
  const total = contentLength ? parseInt(contentLength, 10) : null;
  const reader = response.body.getReader();
  const decoder = new TextDecoder();
  const nodes: GraphNode[] = [];
  const relationships: GraphRelationship[] = [];
  let semanticStatus: GraphSemanticStatus | undefined;
  let buffer = '';
  let downloaded = 0;

  const parseLine = (line: string) => {
    const trimmed = line.trim();
    if (!trimmed) return;

    const record = JSON.parse(trimmed) as
      | { type: 'node'; data: GraphNode }
      | { type: 'relationship'; data: GraphRelationship }
      | { type: 'semantic_status'; data: GraphSemanticStatus }
      | { type: 'error'; error: string };

    if (record.type === 'semantic_status') {
      semanticStatus = record.data;
      return;
    }
    if (record.type === 'node') {
      nodes.push(record.data);
      return;
    }
    if (record.type === 'relationship') {
      relationships.push(record.data);
      return;
    }
    if (record.type === 'error') {
      throw new BackendError(record.error, response.status || 500, 'server');
    }
  };

  while (true) {
    const { done, value } = await reader.read();
    if (done) break;

    downloaded += value.length;
    onProgress?.(downloaded, total);
    buffer += decoder.decode(value, { stream: true });

    const lines = buffer.split('\n');
    buffer = lines.pop() || '';
    for (const line of lines) {
      parseLine(line);
    }
  }

  buffer += decoder.decode();
  parseLine(buffer);

  return { nodes, relationships, semanticStatus };
};

/** Execute a Cypher query. Returns rows. */
export const runQuery = async (
  cypher: string,
  repo?: string,
): Promise<Record<string, unknown>[]> => {
  const response = await fetchFromBackend(`${_backendUrl}/api/query`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ cypher, repo }),
  });
  await assertOk(response);
  const body = await response.json();
  return (body.result ?? body) as Record<string, unknown>[];
};

/** Search with optional enrichment and mode selection. */
export const search = async (
  query: string,
  opts?: { limit?: number; mode?: 'hybrid' | 'semantic' | 'bm25'; enrich?: boolean; repo?: string },
): Promise<EnrichedSearchResult[]> => {
  const response = await fetchFromBackend(`${_backendUrl}/api/search`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      query,
      limit: opts?.limit,
      mode: opts?.mode,
      enrich: opts?.enrich,
      repo: opts?.repo,
    }),
  });
  await assertOk(response);
  const body = await response.json();
  return (body.results ?? []) as EnrichedSearchResult[];
};

/** Grep across file contents in the indexed repo. */
export const grep = async (
  pattern: string,
  repo?: string,
  limit?: number,
): Promise<GrepResult[]> => {
  const params = [
    `pattern=${encodeURIComponent(pattern)}`,
    repoParam(repo),
    limit ? `limit=${limit}` : '',
  ]
    .filter(Boolean)
    .join('&');
  const response = await fetchFromBackend(`${_backendUrl}/api/grep?${params}`);
  await assertOk(response);
  const body = await response.json();
  return (body.results ?? []) as GrepResult[];
};

/** Result from reading a file, optionally with line range. */
export interface ReadFileResult {
  content: string;
  startLine?: number;
  endLine?: number;
  totalLines: number;
}

/** Read a file's content. Supports optional line range (0-indexed). */
export const readFile = async (
  filePath: string,
  options?: { startLine?: number; endLine?: number; repo?: string },
): Promise<ReadFileResult> => {
  const params = [
    `path=${encodeURIComponent(filePath)}`,
    repoParam(options?.repo),
    options?.startLine !== undefined ? `startLine=${options.startLine}` : '',
    options?.endLine !== undefined ? `endLine=${options.endLine}` : '',
  ]
    .filter(Boolean)
    .join('&');
  const response = await fetchFromBackend(`${_backendUrl}/api/file?${params}`);
  await assertOk(response);
  return response.json() as Promise<ReadFileResult>;
};

/** Fetch all processes for a repo. */
export const fetchProcesses = async (repo?: string): Promise<unknown> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/processes${repo ? `?${repoParam(repo)}` : ''}`,
  );
  await assertOk(response);
  return response.json();
};

/** Fetch detail for a single process. */
export const fetchProcessDetail = async (repo: string, name: string): Promise<unknown> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/process?${repoParam(repo)}&name=${encodeURIComponent(name)}`,
  );
  await assertOk(response);
  return response.json();
};

/** Fetch all clusters for a repo. */
export const fetchClusters = async (repo?: string): Promise<unknown> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/clusters${repo ? `?${repoParam(repo)}` : ''}`,
  );
  await assertOk(response);
  return response.json();
};

/** Fetch detail for a single cluster. */
export const fetchClusterDetail = async (repo: string, name: string): Promise<unknown> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/cluster?${repoParam(repo)}&name=${encodeURIComponent(name)}`,
  );
  await assertOk(response);
  return response.json();
};

// ── Analyze API ────────────────────────────────────────────────────────────

export interface AnalyzeRequest {
  path: string;
  embeddings?: boolean;
}

export interface PickLocalFolderResult {
  path: string | null;
  cancelled?: boolean;
}

/** Open the local runtime's native folder picker and return an absolute path. */
export const pickLocalFolder = async (signal?: AbortSignal): Promise<PickLocalFolderResult> => {
  const response = await fetchFromBackend(`${_backendUrl}/api/local/folder-picker`, {
    method: 'POST',
    signal,
  });
  await assertOk(response);
  return response.json() as Promise<PickLocalFolderResult>;
};

/** Start a server-side local analysis job. */
export const startAnalyze = async (
  request: AnalyzeRequest,
): Promise<{ jobId: string; status: string }> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/analyze`,
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(request),
    },
  );
  await assertOk(response);
  return response.json() as Promise<{ jobId: string; status: string }>;
};

/** Poll analysis job status. */
export const getAnalyzeStatus = async (jobId: string): Promise<JobStatus> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/analyze/${encodeURIComponent(jobId)}`,
  );
  await assertOk(response);
  return response.json() as Promise<JobStatus>;
};

/** Cancel a running analysis job. */
export const cancelAnalyze = async (jobId: string): Promise<void> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/analyze/${encodeURIComponent(jobId)}`,
    { method: 'DELETE' },
  );
  await assertOk(response);
};

/** Stream analysis progress via SSE. */
export const streamAnalyzeProgress = (
  jobId: string,
  onProgress: (progress: JobProgress) => void,
  onComplete: (data: AnalyzeCompleteData) => void,
  onError: (error: string) => void,
): AbortController => {
  return streamSSE<JobProgress>(
    `${_backendUrl}/api/analyze/${encodeURIComponent(jobId)}/progress`,
    {
      onMessage: onProgress,
      onComplete: onComplete as (data: unknown) => void,
      onError,
    },
  );
};

// ── Embed API ──────────────────────────────────────────────────────────────

/** Start server-side embedding generation. */
export const startEmbeddings = async (repo: string): Promise<{ jobId: string; status: string }> => {
  const response = await fetchFromBackend(
    `${_backendUrl}/api/embed`,
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ repo }),
    },
  );
  await assertOk(response);
  return response.json() as Promise<{ jobId: string; status: string }>;
};

/** Poll embedding job status. */
export const getEmbedStatus = async (jobId: string): Promise<JobStatus> => {
  const response = await fetchFromBackend(`${_backendUrl}/api/embed/${encodeURIComponent(jobId)}`);
  await assertOk(response);
  return response.json() as Promise<JobStatus>;
};

/** Cancel a running embedding job. */
export const cancelEmbeddings = async (jobId: string): Promise<void> => {
  const response = await fetchFromBackend(`${_backendUrl}/api/embed/${encodeURIComponent(jobId)}`, {
    method: 'DELETE',
  });
  await assertOk(response);
};

/** Stream embedding progress via SSE. */
export const streamEmbeddingProgress = (
  jobId: string,
  onProgress: (progress: JobProgress) => void,
  onComplete: (data: { repoName?: string }) => void,
  onError: (error: string) => void,
): AbortController => {
  return streamSSE<JobProgress>(`${_backendUrl}/api/embed/${encodeURIComponent(jobId)}/progress`, {
    onMessage: onProgress,
    onComplete: onComplete as (data: unknown) => void,
    onError,
  });
};

// ── Convenience: connect to server ─────────────────────────────────────────

export interface ConnectResult {
  nodes: GraphNode[];
  relationships: GraphRelationship[];
  semanticStatus?: GraphSemanticStatus;
  repoInfo: BackendRepo;
}

/**
 * Connect to a server: validate, fetch repo info, download graph.
 * Content is NOT included (use readFile/grep for file access).
 * Pass `awaitAnalysis: true` when the repo may still be queued/analyzing —
 * this enables the backend hold-queue while the repo finishes analysis.
 */
export async function connectToServer(
  url: string,
  onProgress?: (phase: string, downloaded: number, total: number | null) => void,
  signal?: AbortSignal,
  repoName?: string,
  opts?: { awaitAnalysis?: boolean },
): Promise<ConnectResult> {
  const baseUrl = normalizeServerUrl(url);
  setBackendUrl(baseUrl);

  onProgress?.('validating', 0, null);
  const repoInfo = await fetchRepoInfo(repoName, { awaitAnalysis: opts?.awaitAnalysis });
  const graphRepo = repoInfo.repoPath ?? repoInfo.path ?? repoName;

  onProgress?.('downloading', 0, null);
  const { nodes, relationships, semanticStatus } = await fetchGraph(graphRepo, {
    signal,
    onProgress: (downloaded, total) => onProgress?.('downloading', downloaded, total),
  });

  return { nodes, relationships, semanticStatus, repoInfo };
}
