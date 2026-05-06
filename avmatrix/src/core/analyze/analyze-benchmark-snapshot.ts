import fs from 'fs/promises';
import path from 'path';
import type { AnalyzePerformanceReport } from './analyze-metrics.js';
import {
  createGraphCorrectnessSnapshot,
  type GraphCorrectnessSnapshot,
} from './graph-correctness-snapshot.js';
import type { PipelineResult } from '../../types/pipeline.js';

export interface AnalyzeBenchmarkSnapshot {
  readonly schemaVersion: 1;
  readonly createdAt: string;
  readonly label?: string;
  readonly repoName: string;
  readonly repoPath: string;
  readonly stats: {
    readonly files?: number;
    readonly nodes?: number;
    readonly edges?: number;
    readonly communities?: number;
    readonly processes?: number;
  };
  readonly graph?: GraphCorrectnessSnapshot;
  readonly performance?: AnalyzePerformanceReport;
  readonly keyMetrics: AnalyzeBenchmarkKeyMetrics;
}

export interface AnalyzeBenchmarkKeyMetrics {
  readonly totalWallMs?: number;
  readonly phaseMs: Record<string, number>;
  readonly parseMs?: number;
  readonly crossFileMs?: number;
  readonly resolutionMs?: number;
  readonly lbugLoadMs?: number;
  readonly parseableFiles?: number;
  readonly totalParseableMB?: number;
  readonly workerCount?: number;
  readonly parseChunkCount?: number;
  readonly scopeParsedFiles?: number;
  readonly scopeReferenceSites?: number;
  readonly scopeExtractionAstReusedFiles?: number;
  readonly scopeExtractionCompatibilityFiles?: number;
  readonly scopeResolutionReferenceSites?: number;
  readonly scopeResolutionChunks?: number;
  readonly scopeResolutionResolvedReferences?: number;
  readonly scopeResolutionUnresolvedReferences?: number;
  readonly scopeResolutionEdgesEmitted?: number;
  readonly scopeResolutionDuplicateEdgesSkipped?: number;
  readonly crossFileReprocessedFiles?: number;
  readonly crossFileReadContentsMs?: number;
  readonly crossFileProcessCallsParserParseMs?: number;
  readonly csvRelationshipRows?: number;
  readonly ladybugCopyCount?: number;
}

export function createAnalyzeBenchmarkSnapshot(input: {
  readonly repoName: string;
  readonly repoPath: string;
  readonly stats: AnalyzeBenchmarkSnapshot['stats'];
  readonly pipelineResult?: PipelineResult;
  readonly performance?: AnalyzePerformanceReport;
  readonly label?: string;
  readonly createdAt?: string;
}): AnalyzeBenchmarkSnapshot {
  return {
    schemaVersion: 1,
    createdAt: input.createdAt ?? new Date().toISOString(),
    ...(input.label !== undefined ? { label: input.label } : {}),
    repoName: input.repoName,
    repoPath: input.repoPath,
    stats: input.stats,
    ...(input.pipelineResult !== undefined
      ? { graph: createGraphCorrectnessSnapshot(input.pipelineResult) }
      : {}),
    ...(input.performance !== undefined ? { performance: input.performance } : {}),
    keyMetrics: createKeyMetrics(input.performance),
  };
}

export async function writeAnalyzeBenchmarkSnapshot(
  filePath: string,
  snapshot: AnalyzeBenchmarkSnapshot,
): Promise<void> {
  await fs.mkdir(path.dirname(filePath), { recursive: true });
  await fs.writeFile(filePath, `${JSON.stringify(snapshot, null, 2)}\n`, 'utf-8');
}

function createKeyMetrics(
  performance: AnalyzePerformanceReport | undefined,
): AnalyzeBenchmarkKeyMetrics {
  const counters = performance?.counters ?? {};
  const phaseMs = performance?.pipelinePhaseMs ?? {};
  return {
    totalWallMs: performance?.totalWallMs,
    phaseMs,
    parseMs: phaseMs.parse ?? performance?.buckets.parse,
    crossFileMs: phaseMs.crossFile ?? performance?.buckets.crossFile,
    resolutionMs: phaseMs.resolution ?? performance?.buckets.resolution,
    lbugLoadMs: performance?.buckets.lbugLoad,
    parseableFiles: counters.parseableFiles,
    totalParseableMB: counters.totalParseableMB,
    workerCount: counters.workerCount,
    parseChunkCount: counters.parseChunkCount,
    scopeParsedFiles: counters.scopeParsedFiles,
    scopeReferenceSites: counters.scopeReferenceSites,
    scopeExtractionAstReusedFiles: counters.scopeExtractionAstReusedFiles,
    scopeExtractionCompatibilityFiles: counters.scopeExtractionCompatibilityFiles,
    scopeResolutionReferenceSites: counters.scopeResolutionReferenceSites,
    scopeResolutionChunks: counters.scopeResolutionChunks,
    scopeResolutionResolvedReferences: counters.scopeResolutionResolvedReferences,
    scopeResolutionUnresolvedReferences: counters.scopeResolutionUnresolvedReferences,
    scopeResolutionEdgesEmitted: counters.scopeResolutionEdgesEmitted,
    scopeResolutionDuplicateEdgesSkipped: counters.scopeResolutionDuplicateEdgesSkipped,
    crossFileReprocessedFiles: counters.crossFileReprocessedFiles,
    crossFileReadContentsMs: performance?.crossFile?.timings.readContentsMs,
    crossFileProcessCallsParserParseMs: performance?.crossFile?.timings.processCallsParserParseMs,
    csvRelationshipRows: counters.csvRelationshipRows,
    ladybugCopyCount: counters.ladybugCopyCount,
  };
}
