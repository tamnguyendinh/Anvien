import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  connectToServer,
  fetchFileContext,
  fetchFileHotspots,
  fetchGraph,
  fetchRepoInfo,
  normalizeServerUrl,
  setBackendUrl,
} from '../../src/services/backend-client';
import { adaptCompactFileContextResponse } from '../../src/services/file-detail-adapter';

describe('normalizeServerUrl', () => {
  it('canonicalizes localhost to 127.0.0.1', () => {
    expect(normalizeServerUrl('localhost:4848')).toBe('http://127.0.0.1:4848');
  });

  it('adds http:// to 127.0.0.1', () => {
    expect(normalizeServerUrl('127.0.0.1:4848')).toBe('http://127.0.0.1:4848');
  });

  it('rejects non-local hosts', () => {
    expect(() => normalizeServerUrl('example.com')).toThrow(
      /local-only mode only supports backend URLs on 127.0.0.1, localhost, or \[::1\]/i,
    );
  });

  it('strips trailing slashes', () => {
    expect(normalizeServerUrl('http://127.0.0.1:4848/')).toBe('http://127.0.0.1:4848');
    expect(normalizeServerUrl('http://127.0.0.1:4848///')).toBe('http://127.0.0.1:4848');
  });

  it('strips /api suffix (base URL only)', () => {
    expect(normalizeServerUrl('http://127.0.0.1:4848/api')).toBe('http://127.0.0.1:4848');
  });

  it('trims whitespace', () => {
    expect(normalizeServerUrl('  127.0.0.1:4848  ')).toBe('http://127.0.0.1:4848');
  });

  it('supports IPv6 loopback', () => {
    expect(normalizeServerUrl('[::1]:4848')).toBe('http://[::1]:4848');
  });

  it('rejects non-root local paths', () => {
    expect(() => normalizeServerUrl('http://127.0.0.1:4848/anvien')).toThrow(
      /expects the backend URL to point at the local server root or \/api/i,
    );
  });

  it('rejects remote hosts with explicit protocols', () => {
    expect(() => normalizeServerUrl('https://anvien.example.com')).toThrow(
      /local-only mode only supports backend URLs on 127.0.0.1, localhost, or \[::1\]/i,
    );
  });

  it('normalizes backend URLs when set directly', () => {
    setBackendUrl('http://localhost:4848/api');
    expect(normalizeServerUrl('http://localhost:4848/api')).toBe('http://127.0.0.1:4848');
  });

  it('rejects remote URLs when set directly', () => {
    expect(() => setBackendUrl('https://anvien.example.com')).toThrow(
      /local-only mode only supports backend URLs on 127.0.0.1, localhost, or \[::1\]/i,
    );
  });
});

describe('connectToServer', () => {
  it('rejects remote backend URLs before issuing requests', async () => {
    const fetchMock = vi.fn();
    vi.stubGlobal('fetch', fetchMock);

    await expect(connectToServer('https://anvien.example.com')).rejects.toThrow(
      /local-only mode only supports backend URLs on 127.0.0.1, localhost, or \[::1\]/i,
    );
    expect(fetchMock).not.toHaveBeenCalled();
  });

  it('downloads the graph through the canonical repoPath returned by repo info', async () => {
    const canonicalRepoPath = 'F:\\two\\demo';
    const fetchMock = vi.fn((url: string) => {
      if (url.includes('/api/repo')) {
        return Promise.resolve(
          new Response(
            JSON.stringify({
              name: 'demo',
              path: canonicalRepoPath,
              repoPath: canonicalRepoPath,
              indexedAt: new Date().toISOString(),
              stats: {},
            }),
            {
              status: 200,
              headers: { 'Content-Type': 'application/json' },
            },
          ),
        );
      }
      return Promise.resolve(
        new Response('{"nodes":[],"relationships":[]}', {
          status: 200,
          headers: { 'Content-Type': 'application/json' },
        }),
      );
    });
    vi.stubGlobal('fetch', fetchMock);

    const result = await connectToServer('http://localhost:4848', undefined, undefined, 'demo');

    expect(result.repoInfo.repoPath).toBe(canonicalRepoPath);
    expect(fetchMock).toHaveBeenCalledWith(
      `http://127.0.0.1:4848/api/repo?repo=demo`,
      expect.any(Object),
    );
    expect(fetchMock).toHaveBeenCalledWith(
      `http://127.0.0.1:4848/api/graph?repo=${encodeURIComponent(canonicalRepoPath)}&stream=true`,
      expect.any(Object),
    );
    expect(
      fetchMock.mock.calls.some(
        ([url]) => typeof url === 'string' && url.includes('/api/graph?repo=demo&'),
      ),
    ).toBe(false);
  });
});

afterEach(() => {
  vi.restoreAllMocks();
  vi.useRealTimers();
});

describe('fetchGraph', () => {
  it('requests streamed graph responses from the backend', async () => {
    setBackendUrl('http://127.0.0.1:4848');

    const fetchMock = vi.fn().mockResolvedValue(
      new Response('{"nodes":[],"relationships":[]}', {
        status: 200,
        headers: {
          'Content-Type': 'application/json',
        },
      }),
    );
    vi.stubGlobal('fetch', fetchMock);

    await fetchGraph('big-repo');

    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringContaining('/api/graph?repo=big-repo&stream=true'),
      expect.any(Object),
    );
  });

  it('parses NDJSON graph streams incrementally', async () => {
    setBackendUrl('http://127.0.0.1:4848');

    const encoder = new TextEncoder();
    const stream = new ReadableStream<Uint8Array>({
      start(controller) {
        controller.enqueue(
          encoder.encode(
            [
              '{"type":"node","data":{"id":"File:src/app.ts","label":"File","properties":{"name":"app.ts","filePath":"src/app.ts"}}}\n',
              '{"type":"relationship","data":{"id":"File:src/app.ts_CONTAINS_Function:src/app.ts:main","type":"CONTAINS","sourceId":"File:src/app.ts","targetId":"Function:src/app.ts:main"}}\n',
            ].join(''),
          ),
        );
        controller.close();
      },
    });

    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue(
        new Response(stream, {
          status: 200,
          headers: {
            'Content-Type': 'application/x-ndjson',
          },
        }),
      ),
    );

    const progress = vi.fn();
    const result = await fetchGraph('big-repo', { onProgress: progress });

    expect(result.nodes).toHaveLength(1);
    expect(result.relationships).toHaveLength(1);
    expect(result.nodes[0].id).toBe('File:src/app.ts');
    expect(result.relationships[0].type).toBe('CONTAINS');
    expect(progress).toHaveBeenCalled();
  });

  it('parses NDJSON graph lines split across chunks', async () => {
    setBackendUrl('http://127.0.0.1:4848');

    const encoder = new TextEncoder();
    const stream = new ReadableStream<Uint8Array>({
      start(controller) {
        controller.enqueue(
          encoder.encode(
            '{"type":"node","data":{"id":"File:src/app.ts","label":"File","properties":{"name":"app.ts"',
          ),
        );
        controller.enqueue(
          encoder.encode(
            ',"filePath":"src/app.ts"}}}\n{"type":"relationship","data":{"id":"File:src/app.ts_CONTAINS_Function:src/app.ts:main","type":"CONTAINS","sourceId":"File:src/app.ts","targetId":"Function:src/app.ts:main"}}\n',
          ),
        );
        controller.close();
      },
    });

    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue(
        new Response(stream, {
          status: 200,
          headers: {
            'Content-Type': 'application/x-ndjson',
          },
        }),
      ),
    );

    const result = await fetchGraph('big-repo');

    expect(result.nodes).toHaveLength(1);
    expect(result.relationships).toHaveLength(1);
    expect(result.nodes[0].properties.filePath).toBe('src/app.ts');
  });

  it('throws backend errors emitted in the NDJSON stream', async () => {
    setBackendUrl('http://127.0.0.1:4848');

    const encoder = new TextEncoder();
    const stream = new ReadableStream<Uint8Array>({
      start(controller) {
        controller.enqueue(encoder.encode('{"type":"error","error":"stream failed"}\n'));
        controller.close();
      },
    });

    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue(
        new Response(stream, {
          status: 200,
          headers: {
            'Content-Type': 'application/x-ndjson',
          },
        }),
      ),
    );

    await expect(fetchGraph('big-repo')).rejects.toMatchObject({
      message: 'stream failed',
    });
  });
});

describe('fetchFileHotspots', () => {
  it('requests file projection hotspots with repo, sort, pagination, and filters', async () => {
    setBackendUrl('http://127.0.0.1:4848');

    const fetchMock = vi.fn().mockResolvedValue(
      new Response(
        JSON.stringify({
          repo: 'demo',
          total: 1,
          offset: 5,
          limit: 50,
          sort: 'fan-out',
          graph: { path: '.anvien/graph.json', stale: false },
          files: [],
        }),
        {
          status: 200,
          headers: { 'Content-Type': 'application/json' },
        },
      ),
    );
    vi.stubGlobal('fetch', fetchMock);

    const result = await fetchFileHotspots({
      repo: 'demo',
      sort: 'fan-out',
      limit: 50,
      offset: 5,
      kinds: ['source', 'test'],
      apiOnly: true,
      changedOnly: true,
      unresolvedOnly: true,
      highFanIn: true,
      highFanOut: true,
    });

    expect(result.total).toBe(1);
    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringContaining('/api/file-hotspots?'),
      {},
    );
    const url = new URL(fetchMock.mock.calls[0][0] as string);
    expect(url.searchParams.get('repo')).toBe('demo');
    expect(url.searchParams.get('sort')).toBe('fan-out');
    expect(url.searchParams.get('limit')).toBe('50');
    expect(url.searchParams.get('offset')).toBe('5');
    expect(url.searchParams.get('kind')).toBe('source,test');
    expect(url.searchParams.get('apiOnly')).toBe('true');
    expect(url.searchParams.get('changedOnly')).toBe('true');
    expect(url.searchParams.get('unresolvedOnly')).toBe('true');
    expect(url.searchParams.get('highFanIn')).toBe('true');
    expect(url.searchParams.get('highFanOut')).toBe('true');
  });
});

describe('fetchFileContext', () => {
  const compactFileDetailPayload = () => ({
    format: 'file-detail.compact',
    version: 1,
    repo: 'demo',
    graph: { path: '.anvien/graph.json', stale: false },
    target: {
      type: 'file',
      input: 'src/app.ts',
      normalizedPath: 'src/app.ts',
      dispatchMode: 'explicit',
    },
    summary: {
      path: 'src/app.ts',
      language: 'typescript',
      kind: 'source',
      symbolCount: 2,
      exportedSymbolCount: 1,
      inboundRefCount: 1,
      outboundRefCount: 1,
      localRelationshipCount: 1,
      unresolved: 1,
      linkedFlowCount: 1,
      linkedTestCount: 1,
      risk: 'medium',
      stale: false,
      changedSinceAnalyze: false,
    },
    dict: {
      files: ['src/app.ts', 'src/router.ts', 'src/app.test.ts'],
      symbols: [
        'Function:src/app.ts:main',
        'Function:src/app.ts:helper',
        'Function:src/router.ts:createRouter',
        'Function:src/app.test.ts:testMain',
      ],
      sourceSites: ['site-local', 'site-out', 'site-in', 'site-gap'],
      relationshipKinds: ['CALLS', 'USES'],
      gapKinds: ['unresolved_call'],
      classifications: ['in_repo_unresolved'],
      actionabilities: ['analyzer_gap'],
      proofKinds: ['source-site', 'none'],
      sourceSiteStatuses: ['resolved', 'unresolved_local_binding'],
      linkedKinds: ['flow', 'route', 'tool', 'test'],
    },
    tables: {
      symbols: [
        [0, null, 'main', 'function', [3, 0, 12, 0], true, 'function main()', 1, 1, 1, 1],
        [1, 0, 'helper', 'function', [6, 0, 8, 0], false, '', 1, 0, 1, 0],
      ],
      relatedFiles: [
        [
          1,
          'typescript',
          'source',
          'runtime_model',
          'backend_support_model_helper',
          'backend',
          'mcp',
          'parsed',
          4,
          0,
          'medium',
          true,
          false,
          false,
          1,
          { USES: 1 },
        ],
        [
          2,
          'typescript',
          'test',
          'test_helper',
          '',
          'api_test',
          'unknown',
          'parsed',
          2,
          0,
          'low',
          false,
          true,
          false,
          1,
          { CALLS: 1 },
        ],
      ],
      relationships: {
        counts: { local: 1, outbound: 1, inbound: 1, samplesReturned: 3 },
        local: {
          total: 1,
          counts: { CALLS: 1 },
          rows: {
            total: 1,
            returned: 1,
            omitted: 0,
            items: [[0, 0, [3, 0, 3, 10], 0, 0, 1, [6, 0, 6, 10], 0, 0, 0]],
          },
        },
        outboundByFile: [
          {
            file: 1,
            total: 1,
            counts: { USES: 1 },
            rows: {
              total: 1,
              returned: 1,
              omitted: 0,
              items: [[0, 0, [4, 0, 4, 10], 1, 1, 2, [1, 0, 1, 12], 1, 0, 0]],
            },
          },
        ],
        inboundByFile: [
          {
            file: 2,
            total: 1,
            counts: { CALLS: 1 },
            rows: {
              total: 1,
              returned: 1,
              omitted: 0,
              items: [[2, 3, [9, 0, 9, 10], 0, 0, 0, [3, 0, 3, 10], 2, 0, 0]],
            },
          },
        ],
      },
      unresolved: {
        total: 1,
        byKind: { unresolved_call: 1 },
        byClassification: { in_repo_unresolved: 1 },
        byActionability: { analyzer_gap: 1 },
        groups: [
          {
            sourceSymbol: 0,
            total: 1,
            rows: {
              total: 1,
              returned: 1,
              omitted: 0,
              items: [[11, 4, 'missingCall', 0, 0, 0, 0, 1, 3, 1]],
            },
          },
        ],
      },
      linked: {
        counts: { flows: 1, routes: 1, mcpTools: 1, tests: 1 },
        flows: {
          total: 1,
          returned: 1,
          omitted: 0,
          items: [['MCP initialize', 0, 'process', 'high', 'trace-flow']],
        },
        routes: {
          total: 1,
          returned: 1,
          omitted: 0,
          items: [['GET /api/app', 1, 'route-map', 'high', 'trace-route']],
        },
        mcpTools: {
          total: 1,
          returned: 1,
          omitted: 0,
          items: [['context', 2, 'tool-map', 'high', 'trace-tool']],
        },
        tests: {
          total: 1,
          returned: 1,
          omitted: 0,
          items: [['src/app.test.ts', 3, 'relationship', 'high', 'trace-test']],
        },
      },
    },
    quality: {
      parser: 'parsed',
      resolutionConfidence: 'clear',
      unresolvedCalls: 1,
      unresolvedRefs: 0,
      unresolvedImports: 0,
      generated: false,
      stale: false,
      changedSinceAnalyze: false,
    },
    limits: {
      relationshipSamplesPerGroup: 5,
      unresolvedSamplesPerGroup: 4,
      linkedSamplesPerKind: 3,
    },
  });

  it('requests compact selected file projection detail and adapts rows for rendering', async () => {
    setBackendUrl('http://127.0.0.1:4848');

    const fetchMock = vi.fn().mockResolvedValue(
      new Response(JSON.stringify(compactFileDetailPayload()), {
        status: 200,
        headers: { 'Content-Type': 'application/json' },
      }),
    );
    vi.stubGlobal('fetch', fetchMock);

    const result = await fetchFileContext('src/app.ts', {
      repo: 'demo',
      relationships: 5,
      unresolved: 4,
      linked: 3,
    });

    expect(result.summary.path).toBe('src/app.ts');
    const url = new URL(fetchMock.mock.calls[0][0] as string);
    expect(url.pathname).toBe('/api/file-detail');
    expect(url.searchParams.get('path')).toBe('src/app.ts');
    expect(url.searchParams.get('format')).toBe('compact');
    expect(url.searchParams.get('repo')).toBe('demo');
    expect(url.searchParams.get('relationships')).toBe('5');
    expect(url.searchParams.get('unresolved')).toBe('4');
    expect(url.searchParams.get('linked')).toBe('3');

    expect(result.sourceFormat).toBe('compact');
    expect(result.symbolTree[0].name).toBe('main');
    expect(result.symbolTree[0].children?.[0].name).toBe('helper');
    expect(result.relationships.local.samples[0].relationshipKind).toBe('CALLS');
    expect(result.relationships.outboundByFile[0].file).toBe('src/router.ts');
    expect(result.relationships.inboundByFile[0].file).toBe('src/app.test.ts');
    expect(result.unresolved.groups[0].samples[0].targetText).toBe('missingCall');
    expect(result.linked.flows[0].name).toBe('MCP initialize');
    expect(result.relatedFiles).toHaveLength(2);
    expect(result.relatedFiles[0]).toMatchObject({
      file: 'src/router.ts',
      outbound: true,
      inbound: false,
      relationshipTotal: 1,
      relationshipCounts: { USES: 1 },
    });
  });

  it('fails visibly on malformed compact file-detail rows', () => {
    const malformed = compactFileDetailPayload();
    malformed.tables.relationships.local.rows.items = [[0, 0, null, 0, 0, 1, [6, 0, 6, 10], 0, 0, 0]];

    expect(() => adaptCompactFileContextResponse(malformed)).toThrow(
      /Malformed compact file-detail response/i,
    );
  });
});

describe('fetchRepoInfo', () => {
  it('sends awaitAnalysis when repo-info should hold for an active analyze job', async () => {
    setBackendUrl('http://127.0.0.1:4848');
    const fetchMock = vi.fn().mockResolvedValue(
      new Response(
        JSON.stringify({
          name: 'Anvien',
          repoPath: 'F:\\Anvien',
          indexedAt: new Date().toISOString(),
          stats: {},
        }),
        {
          status: 200,
          headers: { 'Content-Type': 'application/json' },
        },
      ),
    );
    vi.stubGlobal('fetch', fetchMock);

    await fetchRepoInfo('F:\\Anvien', { awaitAnalysis: true });

    expect(fetchMock).toHaveBeenCalledWith(
      'http://127.0.0.1:4848/api/repo?repo=F%3A%5CAnvien&awaitAnalysis=true',
      expect.any(Object),
    );
  });

  it('does not attach an internal timeout signal to repo-info requests', async () => {
    setBackendUrl('http://127.0.0.1:4848');

    const fetchMock = vi.fn().mockResolvedValue(
      new Response(
        JSON.stringify({
          name: 'Anvien',
          repoPath: 'F:\\Anvien',
          indexedAt: new Date().toISOString(),
          stats: {},
        }),
        {
          status: 200,
          headers: { 'Content-Type': 'application/json' },
        },
      ),
    );
    vi.stubGlobal('fetch', fetchMock);

    await fetchRepoInfo('F:\\Anvien');

    expect(fetchMock).toHaveBeenCalledWith(
      'http://127.0.0.1:4848/api/repo?repo=F%3A%5CAnvien',
      {},
    );
  });
});
