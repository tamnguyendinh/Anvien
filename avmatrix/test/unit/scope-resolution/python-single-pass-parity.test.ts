import { beforeAll, describe, expect, it } from 'vitest';
import type Parser from 'tree-sitter';
import { SupportedLanguages, type ParsedFile, type SymbolDefinition } from 'avmatrix-shared';
import { createKnowledgeGraph } from '../../../src/core/graph/graph.js';
import { emitReferencesToGraph } from '../../../src/core/ingestion/emit-references.js';
import { finalizeScopeModel } from '../../../src/core/ingestion/finalize-orchestrator.js';
import { pythonProvider } from '../../../src/core/ingestion/languages/python.js';
import { extractParsedFileWithStats } from '../../../src/core/ingestion/scope-extractor-bridge.js';
import { resolveScopeReferenceSites } from '../../../src/core/ingestion/scope-reference-resolver.js';
import { createParserForLanguage } from '../../../src/core/tree-sitter/parser-loader.js';

describe('Python accurate single-pass parity fixture', () => {
  let parser: Parser;

  beforeAll(async () => {
    parser = await createParserForLanguage(SupportedLanguages.Python, 'service.py');
  });

  it('emits CALLS/IMPORTS/ACCESSES/USES/INHERITS through the one graph path', () => {
    const modelsSource = `
class User:
    def save(self) -> None:
        pass

class Repo:
    def find(self, user: User) -> User:
        return user
`;
    const serviceSource = `
from .models import User, Repo

class Admin(User):
    pass

class Service:
    def __init__(self, user: User, repo: Repo):
        self.user = user
        self.repo = repo

    def run(self) -> User:
        self.repo.find(self.user)
        self.user.save()
        return self.user
`;
    const parsedFiles = [
      parseFromExistingAst(parser, modelsSource, 'src/models.py'),
      parseFromExistingAst(parser, serviceSource, 'src/service.py'),
    ];

    const indexes = finalizeScopeModel(parsedFiles, {
      hooks: {
        resolveImportTarget: (targetRaw) => (targetRaw === '.models' ? 'src/models.py' : null),
      },
      mroStrategyForFile: () => 'c3',
    });
    const resolution = resolveScopeReferenceSites(indexes);
    const graph = createKnowledgeGraph();
    addGraphNodesForParsedFiles(graph, parsedFiles);
    const emitStats = emitReferencesToGraph({
      graph,
      scopes: indexes,
      referenceIndex: resolution.referenceIndex,
    });

    const counts = edgeCounts(graph.relationships);
    expect(countFinalizedImports(indexes.imports)).toBe(2);
    expect(resolution.stats.unresolvedReferences).toBe(0);
    expect(resolution.stats.resolvedCalls).toBe(2);
    expect(resolution.stats.resolvedInheritance).toBe(1);
    expect(emitStats.finalizedImportUseEdgesEmitted).toBe(2);
    expect(counts).toMatchObject({
      ACCESSES: 4,
      CALLS: 2,
      IMPORTS: 1,
      INHERITS: 1,
    });
    expect(counts.USES).toBeGreaterThanOrEqual(6);
    expect(
      graph.relationships
        .filter((rel) => rel.resolutionSource === 'scope-resolution')
        .every((rel) => rel.fileHash !== undefined),
    ).toBe(true);
    expect(
      graph.relationships
        .filter((rel) => rel.type === 'USES' && rel.resolutionSource === 'scope-finalize')
        .map((rel) => rel.targetId)
        .sort(),
    ).toEqual(['Class:src/models.py:Repo', 'Class:src/models.py:User']);
  });
});

function parseFromExistingAst(parser: Parser, source: string, filePath: string): ParsedFile {
  const result = extractParsedFileWithStats(
    pythonProvider,
    source,
    filePath,
    SupportedLanguages.Python,
    parser.parse(source).rootNode,
  );
  expect(result.mode).toBe('ast-reused');
  expect(result.parsedFile).toBeDefined();
  return result.parsedFile!;
}

function countFinalizedImports(imports: ReadonlyMap<string, readonly unknown[]>): number {
  let total = 0;
  for (const edges of imports.values()) total += edges.length;
  return total;
}

function edgeCounts(relationships: readonly { type: string }[]): Record<string, number> {
  const counts: Record<string, number> = {};
  for (const relationship of relationships) {
    counts[relationship.type] = (counts[relationship.type] ?? 0) + 1;
  }
  return counts;
}

function addGraphNodesForParsedFiles(
  graph: ReturnType<typeof createKnowledgeGraph>,
  parsedFiles: readonly ParsedFile[],
): void {
  for (const parsed of parsedFiles) {
    graph.addNode({
      id: `File:${parsed.filePath}`,
      label: 'File',
      properties: {
        name: parsed.filePath.split('/').at(-1) ?? parsed.filePath,
        filePath: parsed.filePath,
      },
    });
    for (const def of parsed.localDefs) {
      graph.addNode({
        id: graphNodeIdForDef(def),
        label: def.type,
        properties: {
          name: simpleName(def.qualifiedName) ?? def.nodeId,
          filePath: def.filePath,
          ...(def.qualifiedName !== undefined ? { qualifiedName: def.qualifiedName } : {}),
        },
      });
    }
  }
}

function graphNodeIdForDef(def: SymbolDefinition): string {
  return `${def.type}:${def.filePath}:${def.qualifiedName ?? def.nodeId}`;
}

function simpleName(value: string | undefined): string | undefined {
  if (value === undefined) return undefined;
  const dot = value.lastIndexOf('.');
  return dot === -1 ? value : value.slice(dot + 1);
}
