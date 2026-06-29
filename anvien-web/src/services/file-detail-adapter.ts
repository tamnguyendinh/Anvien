import type {
  CompactFileContextResponse,
  CompactFileRelationshipGroup,
  CompactRelationshipGroup,
  CompactRows,
  FileContextResponse,
  FileLinkedItem,
  FileRelationshipByFileGroup,
  FileRelationshipGroup,
  FileRelationshipSample,
  FileSourceRange,
  FileSymbolTreeNode,
  FileUnresolvedGroup,
  FileUnresolvedSample,
} from '@/generated/anvien-contracts';

export interface FileDetailRelatedFile {
  file: string;
  language?: string;
  kind?: string;
  fileRole?: string;
  fileGroup?: string;
  appLayer?: string;
  functionalArea?: string;
  parseStatus?: string;
  symbolCount: number;
  unresolved: number;
  risk?: string;
  outbound: boolean;
  inbound: boolean;
  local: boolean;
  relationshipTotal: number;
  relationshipCounts?: Record<string, number>;
}

export interface FileDetailContext extends FileContextResponse {
  sourceFormat: 'compact';
  relatedFiles: FileDetailRelatedFile[];
}

const COMPACT_FILE_DETAIL_FORMAT = 'file-detail.compact';

const malformedCompactResponse = (field: string): Error =>
  new Error(`Malformed compact file-detail response: ${field}.`);

const isRecord = (value: unknown): value is Record<string, unknown> =>
  typeof value === 'object' && value !== null && !Array.isArray(value);

const isCompactFileContextResponse = (value: unknown): value is CompactFileContextResponse =>
  isRecord(value) &&
  value.format === COMPACT_FILE_DETAIL_FORMAT &&
  isRecord(value.dict) &&
  isRecord(value.tables) &&
  isRecord(value.summary) &&
  isRecord(value.graph) &&
  isRecord(value.target) &&
  isRecord(value.quality) &&
  isRecord(value.limits);

const dictList = (
  response: CompactFileContextResponse,
  key: keyof CompactFileContextResponse['dict'],
): string[] => {
  const value = response.dict[key];
  if (!Array.isArray(value)) {
    throw malformedCompactResponse(`missing dictionary ${String(key)}`);
  }
  return value;
};

const optionalDictValue = (dict: string[], ref: unknown): string | undefined => {
  if (ref === null || ref === undefined) return undefined;
  if (typeof ref !== 'number' || !Number.isInteger(ref)) return undefined;
  return dict[ref];
};

const requiredDictValue = (dict: string[], ref: unknown, field: string): string => {
  const value = optionalDictValue(dict, ref);
  if (!value) {
    throw malformedCompactResponse(`missing ${field}`);
  }
  return value;
};

const optionalString = (value: unknown): string | undefined =>
  typeof value === 'string' && value.length > 0 ? value : undefined;

const requiredString = (value: unknown, field: string): string => {
  const text = optionalString(value);
  if (!text) {
    throw malformedCompactResponse(`missing ${field}`);
  }
  return text;
};

const requiredNumber = (value: unknown, field: string): number => {
  if (typeof value !== 'number' || Number.isNaN(value)) {
    throw malformedCompactResponse(`missing ${field}`);
  }
  return value;
};

const requiredBoolean = (value: unknown, field: string): boolean => {
  if (typeof value !== 'boolean') {
    throw malformedCompactResponse(`missing ${field}`);
  }
  return value;
};

const compactRows = (rows: CompactRows | undefined, field: string): unknown[][] => {
  if (!rows || !Array.isArray(rows.items)) {
    throw malformedCompactResponse(`missing ${field} rows`);
  }
  return rows.items;
};

const expandRange = (value: unknown, field: string): FileSourceRange => {
  if (!Array.isArray(value)) {
    throw malformedCompactResponse(`missing ${field} range`);
  }
  return {
    startLine: typeof value[0] === 'number' ? value[0] : undefined,
    startColumn: typeof value[1] === 'number' ? value[1] : undefined,
    endLine: typeof value[2] === 'number' ? value[2] : undefined,
    endColumn: typeof value[3] === 'number' ? value[3] : undefined,
  };
};

const numericCounts = (value: unknown): Record<string, number> | undefined => {
  if (!isRecord(value)) return undefined;

  const counts: Record<string, number> = {};
  for (const [key, count] of Object.entries(value)) {
    if (typeof count === 'number' && count > 0) {
      counts[key] = count;
    }
  }
  return Object.keys(counts).length > 0 ? counts : undefined;
};

const expandSymbolTree = (response: CompactFileContextResponse): FileSymbolTreeNode[] => {
  const symbols = dictList(response, 'symbols');
  const byId = new Map<string, { node: FileSymbolTreeNode; parentId?: string }>();
  const roots: FileSymbolTreeNode[] = [];

  for (const row of response.tables.symbols) {
    const id = requiredDictValue(symbols, row[0], 'symbol id');
    const parentId = optionalDictValue(symbols, row[1]);
    const node: FileSymbolTreeNode = {
      id,
      name: requiredString(row[2], 'symbol name'),
      kind: requiredString(row[3], 'symbol kind'),
      range: expandRange(row[4], 'symbol'),
      exported: requiredBoolean(row[5], 'symbol exported'),
      signature: optionalString(row[6]),
      relationshipCounts: {
        local: requiredNumber(row[7], 'symbol local count'),
        inbound: requiredNumber(row[8], 'symbol inbound count'),
        outbound: requiredNumber(row[9], 'symbol outbound count'),
        unresolved: requiredNumber(row[10], 'symbol unresolved count'),
      },
      children: [],
    };
    byId.set(id, { node, parentId });
  }

  for (const { node, parentId } of byId.values()) {
    const parent = parentId ? byId.get(parentId)?.node : undefined;
    if (parent) {
      parent.children = [...(parent.children ?? []), node];
    } else {
      roots.push(node);
    }
  }

  return roots;
};

const expandRelationshipSample = (
  response: CompactFileContextResponse,
  row: unknown[],
): FileRelationshipSample => {
  const files = dictList(response, 'files');
  const symbols = dictList(response, 'symbols');
  const relationshipKinds = dictList(response, 'relationshipKinds');
  const sourceSites = dictList(response, 'sourceSites');
  const proofKinds = dictList(response, 'proofKinds');
  const sourceSiteStatuses = dictList(response, 'sourceSiteStatuses');

  return {
    sourceFile: optionalDictValue(files, row[0]),
    sourceSymbol: optionalDictValue(symbols, row[1]),
    sourceRange: expandRange(row[2], 'relationship source'),
    relationshipKind: requiredDictValue(relationshipKinds, row[3], 'relationship kind'),
    targetFile: optionalDictValue(files, row[4]),
    targetSymbol: optionalDictValue(symbols, row[5]),
    targetRange: expandRange(row[6], 'relationship target'),
    sourceSiteId: optionalDictValue(sourceSites, row[7]),
    proofKind: optionalDictValue(proofKinds, row[8]),
    sourceSiteStatus: optionalDictValue(sourceSiteStatuses, row[9]),
  };
};

const expandRelationshipGroup = (
  response: CompactFileContextResponse,
  group: CompactRelationshipGroup,
): FileRelationshipGroup => ({
  total: group.total,
  counts: group.counts,
  samples: compactRows(group.rows, 'relationship').map((row) =>
    expandRelationshipSample(response, row),
  ),
});

const expandFileRelationshipGroup = (
  response: CompactFileContextResponse,
  group: CompactFileRelationshipGroup,
): FileRelationshipByFileGroup => ({
  ...expandRelationshipGroup(response, group),
  file: requiredDictValue(dictList(response, 'files'), group.file, 'relationship file'),
});

const expandUnresolvedSample = (
  response: CompactFileContextResponse,
  row: unknown[],
): FileUnresolvedSample => {
  const symbols = dictList(response, 'symbols');
  const gapKinds = dictList(response, 'gapKinds');
  const classifications = dictList(response, 'classifications');
  const actionabilities = dictList(response, 'actionabilities');
  const proofKinds = dictList(response, 'proofKinds');
  const sourceSites = dictList(response, 'sourceSites');
  const sourceSiteStatuses = dictList(response, 'sourceSiteStatuses');

  return {
    line: typeof row[0] === 'number' ? row[0] : undefined,
    column: typeof row[1] === 'number' ? row[1] : undefined,
    targetText: optionalString(row[2]),
    sourceSymbol: optionalDictValue(symbols, row[3]),
    gapKind: optionalDictValue(gapKinds, row[4]),
    classification: optionalDictValue(classifications, row[5]),
    actionability: optionalDictValue(actionabilities, row[6]),
    proofKind: optionalDictValue(proofKinds, row[7]),
    sourceSiteId: optionalDictValue(sourceSites, row[8]),
    sourceSiteStatus: optionalDictValue(sourceSiteStatuses, row[9]),
  };
};

const expandLinkedRows = (
  response: CompactFileContextResponse,
  rows: CompactRows,
  field: string,
): FileLinkedItem[] => {
  const linkedKinds = dictList(response, 'linkedKinds');
  return compactRows(rows, field).map((row) => ({
    name: requiredString(row[0], `${field} name`),
    kind: optionalDictValue(linkedKinds, row[1]),
    source: optionalString(row[2]),
    confidence: optionalString(row[3]),
    trace: optionalString(row[4]),
  }));
};

const expandRelatedFiles = (response: CompactFileContextResponse): FileDetailRelatedFile[] => {
  const files = dictList(response, 'files');

  return response.tables.relatedFiles.map((row) => ({
    file: requiredDictValue(files, row[0], 'related file'),
    language: optionalString(row[1]),
    kind: optionalString(row[2]),
    fileRole: optionalString(row[3]),
    fileGroup: optionalString(row[4]),
    appLayer: optionalString(row[5]),
    functionalArea: optionalString(row[6]),
    parseStatus: optionalString(row[7]),
    symbolCount: requiredNumber(row[8], 'related symbol count'),
    unresolved: requiredNumber(row[9], 'related unresolved count'),
    risk: optionalString(row[10]),
    outbound: requiredBoolean(row[11], 'related outbound flag'),
    inbound: requiredBoolean(row[12], 'related inbound flag'),
    local: requiredBoolean(row[13], 'related local flag'),
    relationshipTotal: requiredNumber(row[14], 'related relationship total'),
    relationshipCounts: numericCounts(row[15]),
  }));
};

export const adaptCompactFileContextResponse = (payload: unknown): FileDetailContext => {
  if (!isCompactFileContextResponse(payload)) {
    throw malformedCompactResponse('expected file-detail.compact format');
  }

  return {
    repo: payload.repo,
    repoPath: payload.repoPath,
    graph: payload.graph,
    target: payload.target,
    summary: payload.summary,
    symbolTree: expandSymbolTree(payload),
    relationships: {
      counts: payload.tables.relationships.counts,
      local: expandRelationshipGroup(payload, payload.tables.relationships.local),
      outboundByFile: payload.tables.relationships.outboundByFile.map((group) =>
        expandFileRelationshipGroup(payload, group),
      ),
      inboundByFile: payload.tables.relationships.inboundByFile.map((group) =>
        expandFileRelationshipGroup(payload, group),
      ),
    },
    unresolved: {
      total: payload.tables.unresolved.total,
      byKind: payload.tables.unresolved.byKind,
      byClassification: payload.tables.unresolved.byClassification,
      byActionability: payload.tables.unresolved.byActionability,
      groups: payload.tables.unresolved.groups.map<FileUnresolvedGroup>((group) => ({
        sourceSymbol: optionalDictValue(dictList(payload, 'symbols'), group.sourceSymbol),
        total: group.total,
        samples: compactRows(group.rows, 'unresolved').map((row) =>
          expandUnresolvedSample(payload, row),
        ),
      })),
    },
    linked: {
      counts: payload.tables.linked.counts,
      flows: expandLinkedRows(payload, payload.tables.linked.flows, 'flows'),
      routes: expandLinkedRows(payload, payload.tables.linked.routes, 'routes'),
      mcpTools: expandLinkedRows(payload, payload.tables.linked.mcpTools, 'mcpTools'),
      tests: expandLinkedRows(payload, payload.tables.linked.tests, 'tests'),
    },
    quality: payload.quality,
    limits: payload.limits,
    sourceFormat: 'compact',
    relatedFiles: expandRelatedFiles(payload),
  };
};
