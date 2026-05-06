import { describe, expect, it } from 'vitest';
import { createAnalyzeBenchmarkSnapshot } from '../../src/core/analyze/analyze-benchmark-snapshot.js';
import { buildTestGraph } from '../helpers/test-graph.js';

describe('analyze benchmark snapshot', () => {
  it('captures graph correctness and key performance metrics for benchmark comparisons', () => {
    const graph = buildTestGraph(
      [
        {
          id: 'Function:src/app.ts:run',
          label: 'Function',
          name: 'run',
          filePath: 'src/app.ts',
        },
        {
          id: 'Method:src/model.ts:save',
          label: 'Method',
          name: 'save',
          filePath: 'src/model.ts',
        },
      ],
      [
        {
          sourceId: 'Function:src/app.ts:run',
          targetId: 'Method:src/model.ts:save',
          type: 'CALLS',
          confidence: 0.95,
          reason: 'scope-resolution: call | confidence 0.950',
        },
      ],
    );

    const snapshot = createAnalyzeBenchmarkSnapshot({
      repoName: 'demo',
      repoPath: 'F:/demo',
      createdAt: '2026-05-06T00:00:00.000Z',
      label: 'baseline',
      stats: { files: 2, nodes: 2, edges: 1 },
      pipelineResult: {
        graph,
        repoPath: 'F:/demo',
        totalFileCount: 2,
        usedWorkerPool: true,
        performance: {
          phaseMs: { parse: 10, crossFile: 20, resolution: 5 },
          counters: { crossFileReprocessedFiles: 3 },
        },
      },
      performance: {
        totalWallMs: 100,
        buckets: { lbugLoad: 30 },
        pipelinePhaseMs: { parse: 10, crossFile: 20, resolution: 5 },
        ftsIndexMs: {},
        counters: {
          parseableFiles: 2,
          scopeReferenceSites: 4,
          scopeResolutionReferenceSites: 4,
          scopeResolutionResolvedReferences: 3,
          scopeResolutionUnresolvedReferences: 1,
          scopeResolutionEdgesEmitted: 3,
          crossFileReprocessedFiles: 3,
          csvRelationshipRows: 12,
        },
        bottlenecks: [],
        overheadMs: 35,
        crossFile: {
          timings: {
            readContentsMs: 7,
            processCallsParserParseMs: 9,
          },
          counters: {},
        },
      },
    });

    expect(snapshot.graph?.byRelationshipType).toEqual({ CALLS: 1 });
    expect(snapshot.keyMetrics).toMatchObject({
      totalWallMs: 100,
      parseMs: 10,
      crossFileMs: 20,
      resolutionMs: 5,
      lbugLoadMs: 30,
      parseableFiles: 2,
      scopeReferenceSites: 4,
      scopeResolutionResolvedReferences: 3,
      scopeResolutionUnresolvedReferences: 1,
      crossFileReprocessedFiles: 3,
      crossFileReadContentsMs: 7,
      crossFileProcessCallsParserParseMs: 9,
      csvRelationshipRows: 12,
    });
  });
});
