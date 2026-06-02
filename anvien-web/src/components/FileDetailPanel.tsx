import { useEffect, useMemo, useState, type ReactNode } from "react";
import {
  AlertTriangle,
  Braces,
  ChevronDown,
  ChevronRight,
  FileCode,
  GitBranch,
  Layers,
  List,
  Loader2,
  Target,
} from "@/lib/lucide-icons";
import type {
  FileContextResponse,
  FileSummary,
  FileLinkedItem,
  FileRelationshipByFileGroup,
  FileRelationshipGroup,
  FileRelationshipSample,
  FileSymbolTreeNode,
  FileUnresolvedGroup,
  FileUnresolvedSample,
} from "@/generated/anvien-contracts";
import { fetchFileContext } from "../services/backend-client";

interface FileDetailPanelProps {
  repoName?: string;
  filePath?: string | null;
  onFocusNode?: (nodeId: string) => void;
}

type RelationshipGroupWithFile = FileRelationshipGroup & { file?: string };

const SAMPLE_LIMIT = 5;

const formatKey = (value: string | undefined): string =>
  value
    ? value
        .replace(/[-_]/g, " ")
        .split(" ")
        .filter(Boolean)
        .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
        .join(" ")
    : "Unknown";

const compactCount = (value: number | undefined): string =>
  new Intl.NumberFormat("en", { notation: "compact" }).format(value ?? 0);

const countEntries = (counts: Record<string, number> | undefined) =>
  Object.entries(counts ?? {})
    .filter(([, count]) => count > 0)
    .sort((left, right) => right[1] - left[1]);

const defaultUnresolvedCount = (summary: FileSummary): number =>
  summary.defaultVisibleUnresolvedSourceSiteCount ?? summary.unresolvedSourceSiteCount ?? 0;

const isTestFile = (summary: FileSummary): boolean => summary.kind === "test";

const isDefaultVisibleUnresolvedSample = (sample: FileUnresolvedSample): boolean =>
  sample.actionability !== "non_actionable";

const defaultVisibleUnresolvedGroups = (
  groups: FileUnresolvedGroup[],
  testFile: boolean,
): FileUnresolvedGroup[] => {
  if (testFile) return [];
  return groups
    .map((group) => {
      const samples = group.samples.filter(isDefaultVisibleUnresolvedSample);
      if (samples.length === 0) return null;
      return {
        ...group,
        total: samples.length,
        samples,
      };
    })
    .filter((group): group is FileUnresolvedGroup => group !== null);
};

const unresolvedKindCounts = (groups: FileUnresolvedGroup[]): Record<string, number> => {
  const counts: Record<string, number> = {};
  for (const group of groups) {
    for (const sample of group.samples) {
      const key = sample.gapKind || "unknown";
      counts[key] = (counts[key] ?? 0) + 1;
    }
  }
  return counts;
};

const sampleLocation = (sample: FileUnresolvedSample): string => {
  const parts = [
    typeof sample.line === "number" ? `L${sample.line}` : "",
    typeof sample.column === "number" ? `C${sample.column}` : "",
  ].filter(Boolean);
  return parts.length > 0 ? parts.join(":") : "No range";
};

const relationshipSampleLabel = (sample: FileRelationshipSample): string => {
  const left = sample.sourceSymbol || sample.sourceFile || "source";
  const right = sample.targetSymbol || sample.targetFile || "target";
  return `${left} -> ${right}`;
};

const Section = ({
  icon,
  title,
  meta,
  testId,
  children,
}: {
  icon: ReactNode;
  title: string;
  meta?: string;
  testId?: string;
  children: ReactNode;
}) => (
  <section
    className="border-t border-workspace-border-default px-3 py-3"
    data-testid={testId}
  >
    <div className="mb-2 flex min-w-0 items-center gap-2">
      <span className="text-workspace-text-secondary">{icon}</span>
      <h3 className="min-w-0 flex-1 truncate text-xs font-semibold tracking-wide text-workspace-text-primary uppercase">
        {title}
      </h3>
      {meta && (
        <span className="rounded border border-workspace-border-default bg-workspace-inset px-1.5 py-0.5 font-mono text-[10px] text-workspace-text-secondary">
          {meta}
        </span>
      )}
    </div>
    {children}
  </section>
);

const EmptyLine = ({ children }: { children: ReactNode }) => (
  <div className="rounded border border-workspace-border-default bg-workspace-inset px-2 py-2 text-xs text-workspace-text-secondary">
    {children}
  </div>
);

const Stat = ({ label, value }: { label: string; value: string | number }) => (
  <div className="min-w-0 rounded border border-workspace-border-default bg-workspace-inset px-2 py-1.5">
    <div className="truncate text-[10px] font-semibold tracking-wide text-workspace-text-secondary uppercase">
      {label}
    </div>
    <div className="truncate font-mono text-sm text-workspace-text-primary">
      {value}
    </div>
  </div>
);

const Pill = ({
  label,
  value,
  tone = "neutral",
}: {
  label: string;
  value: string | number | boolean | undefined;
  tone?: "neutral" | "warning" | "good";
}) => {
  const toneClass =
    tone === "warning"
      ? "border-warning/40 bg-warning/10 text-warning"
      : tone === "good"
        ? "border-success/40 bg-success/10 text-success"
        : "border-workspace-border-default bg-workspace-inset text-workspace-text-secondary";
  return (
    <span
      className={`inline-flex min-w-0 items-center gap-1 rounded border px-1.5 py-0.5 text-[10px] ${toneClass}`}
    >
      <span className="font-semibold">{label}</span>
      <span className="min-w-0 truncate font-mono text-workspace-text-primary">
        {typeof value === "boolean" ? (value ? "yes" : "no") : value ?? "unknown"}
      </span>
    </span>
  );
};

const collectDefaultExpandedIds = (nodes: FileSymbolTreeNode[]) =>
  new Set(nodes.filter((node) => (node.children?.length ?? 0) > 0).map((node) => node.id));

const SymbolRow = ({
  node,
  depth,
  expandedIds,
  onToggle,
  onFocusNode,
}: {
  node: FileSymbolTreeNode;
  depth: number;
  expandedIds: Set<string>;
  onToggle: (nodeId: string) => void;
  onFocusNode?: (nodeId: string) => void;
}) => {
  const hasChildren = (node.children?.length ?? 0) > 0;
  const expanded = expandedIds.has(node.id);
  const range =
    typeof node.range?.startLine === "number"
      ? `L${node.range.startLine}${typeof node.range.endLine === "number" ? `-${node.range.endLine}` : ""}`
      : "No range";
  return (
    <div className="min-w-0">
      <div
        className="flex min-w-0 items-center gap-1.5 rounded px-1 py-1 text-xs hover:bg-workspace-inset"
        style={{ paddingLeft: `${Math.min(depth * 14, 42) + 4}px` }}
      >
        {hasChildren ? (
          <button
            type="button"
            onClick={() => onToggle(node.id)}
            className="workspace-outline-button rounded p-0.5 text-workspace-text-secondary hover:text-workspace-text-primary"
            aria-label={`${expanded ? "Collapse" : "Expand"} ${node.name}`}
          >
            {expanded ? (
              <ChevronDown className="h-3 w-3" />
            ) : (
              <ChevronRight className="h-3 w-3" />
            )}
          </button>
        ) : (
          <span className="h-4 w-4" />
        )}
        <span className="min-w-0 flex-1 truncate font-mono text-workspace-text-primary" title={node.signature || node.name}>
          {node.name}
        </span>
        <span className="shrink-0 text-[10px] text-workspace-text-secondary">
          {formatKey(node.kind)}
        </span>
        <span className="shrink-0 font-mono text-[10px] text-workspace-text-secondary">
          {range}
        </span>
        {onFocusNode && (
          <button
            type="button"
            onClick={() => onFocusNode(node.id)}
            className="workspace-outline-button rounded p-0.5 text-workspace-text-secondary hover:text-workspace-text-primary"
            aria-label={`Focus ${node.name}`}
            data-testid="file-detail-focus-symbol"
            title="Focus symbol in graph"
          >
            <Target className="h-3 w-3" />
          </button>
        )}
      </div>
      {hasChildren && expanded && (
        <div className="min-w-0">
          {node.children?.map((child) => (
            <SymbolRow
              key={child.id}
              node={child}
              depth={depth + 1}
              expandedIds={expandedIds}
              onToggle={onToggle}
              onFocusNode={onFocusNode}
            />
          ))}
        </div>
      )}
    </div>
  );
};

const RelationshipSamples = ({ samples }: { samples: FileRelationshipSample[] }) => {
  if (samples.length === 0) return null;
  return (
    <div className="mt-1 space-y-1">
      {samples.slice(0, SAMPLE_LIMIT).map((sample, index) => (
        <div
          key={`${sample.sourceSiteId ?? index}-${sample.relationshipKind}`}
          className="min-w-0 rounded bg-workspace-inset px-2 py-1 text-[11px] text-workspace-text-secondary"
        >
          <div className="flex min-w-0 items-center gap-1.5">
            <span className="shrink-0 rounded bg-workspace-base px-1 font-mono text-[10px] text-workspace-text-primary">
              {formatKey(sample.relationshipKind)}
            </span>
            <span className="min-w-0 truncate font-mono" title={relationshipSampleLabel(sample)}>
              {relationshipSampleLabel(sample)}
            </span>
          </div>
          {(sample.proofKind || sample.sourceSiteStatus) && (
            <div className="mt-0.5 truncate text-[10px]">
              {[sample.proofKind, sample.sourceSiteStatus].filter(Boolean).map(formatKey).join(" / ")}
            </div>
          )}
        </div>
      ))}
    </div>
  );
};

const RelationshipGroupBlock = ({ group }: { group: RelationshipGroupWithFile }) => {
  const entries = countEntries(group.counts);
  return (
    <div className="min-w-0 rounded border border-workspace-border-default bg-workspace-surface px-2 py-2">
      <div className="flex min-w-0 items-center gap-2">
        <span className="min-w-0 flex-1 truncate font-mono text-xs text-workspace-text-primary" title={group.file ?? "local"}>
          {group.file ?? "local"}
        </span>
        <span className="rounded bg-workspace-inset px-1.5 py-0.5 font-mono text-[10px] text-workspace-text-secondary">
          {compactCount(group.total)}
        </span>
      </div>
      {entries.length > 0 && (
        <div className="mt-1 flex flex-wrap gap-1">
          {entries.slice(0, 4).map(([kind, count]) => (
            <span
              key={kind}
              className="rounded bg-workspace-inset px-1.5 py-0.5 text-[10px] text-workspace-text-secondary"
            >
              {formatKey(kind)} {count}
            </span>
          ))}
        </div>
      )}
      <RelationshipSamples samples={group.samples} />
    </div>
  );
};

const RelationshipSection = ({
  title,
  groups,
  empty,
}: {
  title: string;
  groups: RelationshipGroupWithFile[];
  empty: string;
}) => (
  <div className="space-y-1.5">
    <div className="text-[10px] font-semibold tracking-wide text-workspace-text-secondary uppercase">
      {title}
    </div>
    {groups.length === 0 ? (
      <EmptyLine>{empty}</EmptyLine>
    ) : (
      <div className="space-y-1.5">
        {groups.slice(0, 6).map((group, index) => (
          <RelationshipGroupBlock
            key={`${title}-${group.file ?? "local"}-${index}`}
            group={group}
          />
        ))}
      </div>
    )}
  </div>
);

const UnresolvedSamples = ({ samples }: { samples: FileUnresolvedSample[] }) => (
  <div className="mt-1 space-y-1">
    {samples.slice(0, SAMPLE_LIMIT).map((sample, index) => (
      <div
        key={`${sample.sourceSiteId ?? sample.targetText ?? index}`}
        className="rounded bg-workspace-inset px-2 py-1 text-[11px] text-workspace-text-secondary"
      >
        <div className="flex min-w-0 items-center gap-1.5">
          <span className="shrink-0 font-mono text-[10px] text-warning">
            {sampleLocation(sample)}
          </span>
          <span className="min-w-0 truncate font-mono text-workspace-text-primary" title={sample.targetText}>
            {sample.targetText || "unknown target"}
          </span>
        </div>
        <div className="mt-0.5 truncate text-[10px]">
          {[sample.gapKind, sample.classification, sample.actionability]
            .filter(Boolean)
            .map(formatKey)
            .join(" / ")}
        </div>
      </div>
    ))}
  </div>
);

const LinkedList = ({
  label,
  count,
  items,
}: {
  label: string;
  count: number;
  items: FileLinkedItem[];
}) => (
  <div className="min-w-0 rounded border border-workspace-border-default bg-workspace-surface px-2 py-2">
    <div className="flex min-w-0 items-center gap-2">
      <span className="min-w-0 flex-1 truncate text-xs font-semibold text-workspace-text-primary">
        {label}
      </span>
      <span className="rounded bg-workspace-inset px-1.5 py-0.5 font-mono text-[10px] text-workspace-text-secondary">
        {compactCount(count)}
      </span>
    </div>
    {items.length === 0 ? (
      <div className="mt-1 text-[11px] text-workspace-text-secondary">None</div>
    ) : (
      <div className="mt-1 space-y-1">
        {items.slice(0, SAMPLE_LIMIT).map((item, index) => (
          <div
            key={`${item.name}-${index}`}
            className="min-w-0 rounded bg-workspace-inset px-2 py-1 text-[11px] text-workspace-text-secondary"
          >
            <div className="truncate font-mono text-workspace-text-primary" title={item.name}>
              {item.name}
            </div>
            <div className="truncate text-[10px]">
              {[item.kind, item.source, item.confidence, item.trace].filter(Boolean).join(" / ")}
            </div>
          </div>
        ))}
      </div>
    )}
  </div>
);

export const FileDetailPanel = ({
  repoName,
  filePath,
  onFocusNode,
}: FileDetailPanelProps) => {
  const [context, setContext] = useState<FileContextResponse | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [expandedNodeIds, setExpandedNodeIds] = useState<Set<string>>(new Set());

  useEffect(() => {
    if (!filePath) {
      setContext(null);
      setError(null);
      setIsLoading(false);
      return;
    }

    let cancelled = false;
    setIsLoading(true);
    setError(null);
    setContext(null);

    fetchFileContext(filePath, {
      repo: repoName,
      relationships: SAMPLE_LIMIT,
      unresolved: SAMPLE_LIMIT,
      linked: SAMPLE_LIMIT,
    })
      .then((nextContext) => {
        if (!cancelled) {
          setContext(nextContext);
          setExpandedNodeIds(collectDefaultExpandedIds(nextContext.symbolTree));
        }
      })
      .catch((err) => {
        if (!cancelled) {
          setError(err instanceof Error ? err.message : "Unable to load file detail.");
        }
      })
      .finally(() => {
        if (!cancelled) setIsLoading(false);
      });

    return () => {
      cancelled = true;
    };
  }, [filePath, repoName]);

  const localGroups = useMemo<RelationshipGroupWithFile[]>(() => {
    if (!context || context.relationships.local.total === 0) return [];
    return [{ ...context.relationships.local, file: "local" }];
  }, [context]);

  const outboundGroups = useMemo<RelationshipGroupWithFile[]>(
    () => ((context?.relationships.outboundByFile ?? []) as FileRelationshipByFileGroup[]).slice(0, 6),
    [context],
  );

  const inboundGroups = useMemo<RelationshipGroupWithFile[]>(
    () => ((context?.relationships.inboundByFile ?? []) as FileRelationshipByFileGroup[]).slice(0, 6),
    [context],
  );

  const toggleSymbol = (nodeId: string) => {
    setExpandedNodeIds((current) => {
      const next = new Set(current);
      if (next.has(nodeId)) next.delete(nodeId);
      else next.add(nodeId);
      return next;
    });
  };

  if (!filePath) {
    return (
      <div
        className="border-b border-workspace-border-default bg-workspace-surface"
        data-testid="file-detail-panel"
      >
        <Section
          icon={<FileCode className="h-3.5 w-3.5" />}
          title="File Detail"
        >
          <EmptyLine>Select a file to inspect file graph context.</EmptyLine>
        </Section>
      </div>
    );
  }

  if (isLoading) {
    return (
      <div
        className="border-b border-workspace-border-default bg-workspace-surface"
        data-testid="file-detail-panel"
      >
        <Section
          icon={<FileCode className="h-3.5 w-3.5" />}
          title="File Detail"
        >
          <div className="flex items-center gap-2 rounded border border-workspace-border-default bg-workspace-inset px-2 py-2 text-xs text-workspace-text-secondary">
            <Loader2 className="h-3.5 w-3.5 animate-spin" />
            Loading file detail...
          </div>
        </Section>
      </div>
    );
  }

  if (error) {
    return (
      <div
        className="border-b border-workspace-border-default bg-workspace-surface"
        data-testid="file-detail-panel"
      >
        <Section
          icon={<AlertTriangle className="h-3.5 w-3.5 text-warning" />}
          title="File Detail"
        >
          <EmptyLine>{error}</EmptyLine>
        </Section>
      </div>
    );
  }

  if (!context) {
    return (
      <div
        className="border-b border-workspace-border-default bg-workspace-surface"
        data-testid="file-detail-panel"
      >
        <Section
          icon={<FileCode className="h-3.5 w-3.5" />}
          title="File Detail"
        >
          <EmptyLine>No file detail returned.</EmptyLine>
        </Section>
      </div>
    );
  }

  const summary = context.summary;
  const quality = context.quality;
  const testFile = isTestFile(summary);
  const defaultUnresolved = defaultUnresolvedCount(summary);
  const unresolvedGroups = defaultVisibleUnresolvedGroups(context.unresolved.groups, testFile);
  const unresolvedKinds = unresolvedKindCounts(unresolvedGroups);

  return (
    <div
      className="border-b border-workspace-border-default bg-workspace-surface"
      data-testid="file-detail-panel"
    >
      <Section
        icon={<FileCode className="h-3.5 w-3.5" />}
        title="File Detail"
        meta={formatKey(summary.risk)}
        testId="file-detail-section-summary"
      >
        <div className="min-w-0 truncate font-mono text-xs text-workspace-text-primary" title={summary.path}>
          {summary.path}
        </div>
        <div className="mt-2 grid grid-cols-3 gap-1.5">
          <Stat label="Symbols" value={compactCount(summary.symbolCount)} />
          <Stat label="In" value={compactCount(summary.inboundRefCount)} />
          <Stat label="Out" value={compactCount(summary.outboundRefCount)} />
          <Stat label="Local" value={compactCount(summary.localRelationshipCount)} />
          <Stat label="Unresolved" value={compactCount(defaultUnresolved)} />
          <Stat label="Tests" value={compactCount(summary.linkedTestCount)} />
        </div>
        <div className="mt-2 flex flex-wrap gap-1">
          <Pill label="Layer" value={formatKey(summary.appLayer)} />
          <Pill label="Area" value={formatKey(summary.functionalArea)} />
          <Pill label="Kind" value={testFile ? "Test File" : formatKey(summary.kind)} />
          <Pill label="Lang" value={summary.language ?? "unknown"} />
        </div>
      </Section>

      <Section
        icon={<Layers className="h-3.5 w-3.5" />}
        title="Quality"
        testId="file-detail-section-quality"
      >
        <div className="flex flex-wrap gap-1">
          <Pill
            label="Parser"
            value={formatKey(quality.parser)}
            tone={quality.parser === "parsed" ? "good" : "warning"}
          />
          <Pill
            label="Resolution"
            value={formatKey(quality.resolutionConfidence)}
            tone={quality.resolutionConfidence === "degraded" ? "warning" : "neutral"}
          />
          {!testFile && (
            <>
              <Pill label="Calls" value={quality.unresolvedCalls} />
              <Pill label="Refs" value={quality.unresolvedRefs} />
              <Pill label="Imports" value={quality.unresolvedImports} />
            </>
          )}
          <Pill label="Generated" value={quality.generated} />
          <Pill label="Stale" value={quality.stale} tone={quality.stale ? "warning" : "neutral"} />
          <Pill
            label="Changed"
            value={quality.changedSinceAnalyze}
            tone={quality.changedSinceAnalyze ? "warning" : "neutral"}
          />
        </div>
      </Section>

      <Section
        icon={<Braces className="h-3.5 w-3.5" />}
        title="Symbol Tree"
        meta={`${summary.symbolCount} symbols`}
        testId="file-detail-section-symbol-tree"
      >
        {context.symbolTree.length === 0 ? (
          <EmptyLine>No symbols are declared in this file.</EmptyLine>
        ) : (
          <div className="space-y-0.5">
            {context.symbolTree.slice(0, 24).map((node) => (
              <SymbolRow
                key={node.id}
                node={node}
                depth={0}
                expandedIds={expandedNodeIds}
                onToggle={toggleSymbol}
                onFocusNode={onFocusNode}
              />
            ))}
          </div>
        )}
      </Section>

      <Section
        icon={<GitBranch className="h-3.5 w-3.5" />}
        title="Relationships"
        meta={`${compactCount(context.relationships.counts.samplesReturned)} samples`}
        testId="file-detail-section-relationships"
      >
        <div className="space-y-3">
          <RelationshipSection
            title="Local"
            groups={localGroups}
            empty="No local relationship samples."
          />
          <RelationshipSection
            title="Outbound"
            groups={outboundGroups}
            empty="No outbound file dependencies."
          />
          <RelationshipSection
            title="Inbound"
            groups={inboundGroups}
            empty="No inbound file dependents."
          />
        </div>
      </Section>

      <Section
        icon={<AlertTriangle className="h-3.5 w-3.5" />}
        title="Unresolved"
        meta={`${compactCount(defaultUnresolved)} sites`}
        testId="file-detail-section-unresolved"
      >
        <div className="mb-2 flex flex-wrap gap-1">
          {countEntries(unresolvedKinds)
            .slice(0, 6)
            .map(([kind, count]) => (
              <Pill key={kind} label={formatKey(kind)} value={count} />
            ))}
        </div>
        {unresolvedGroups.length === 0 ? (
          <EmptyLine>No default unresolved source-site samples.</EmptyLine>
        ) : (
          <div className="space-y-1.5">
            {unresolvedGroups.slice(0, 6).map((group, index) => (
              <div
                key={`${group.sourceSymbol ?? "unknown"}-${index}`}
                className="rounded border border-workspace-border-default bg-workspace-surface px-2 py-2"
              >
                <div className="flex min-w-0 items-center gap-2">
                  <span className="min-w-0 flex-1 truncate font-mono text-xs text-workspace-text-primary">
                    {group.sourceSymbol || "Unknown source"}
                  </span>
                  <span className="rounded bg-workspace-inset px-1.5 py-0.5 font-mono text-[10px] text-workspace-text-secondary">
                    {compactCount(group.total)}
                  </span>
                </div>
                <UnresolvedSamples samples={group.samples} />
              </div>
            ))}
          </div>
        )}
      </Section>

      <Section
        icon={<List className="h-3.5 w-3.5" />}
        title="Linked"
        testId="file-detail-section-linked"
      >
        <div className="grid grid-cols-2 gap-1.5">
          <LinkedList
            label="Flows"
            count={context.linked.counts.flows}
            items={context.linked.flows}
          />
          <LinkedList
            label="Routes"
            count={context.linked.counts.routes}
            items={context.linked.routes}
          />
          <LinkedList
            label="MCP Tools"
            count={context.linked.counts.mcpTools}
            items={context.linked.mcpTools}
          />
          <LinkedList
            label="Tests"
            count={context.linked.counts.tests}
            items={context.linked.tests}
          />
        </div>
      </Section>
    </div>
  );
};
