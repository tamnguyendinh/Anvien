import { useEffect, useMemo, useState } from "react";
import type { FileHotspotsResponse, FileSummary } from "@/generated/anvien-contracts";
import {
  AlertTriangle,
  ArrowDown,
  FileCode,
  Filter,
  GitBranch,
  Loader2,
  Search,
  Table,
} from "@/lib/lucide-icons";
import {
  fetchFileHotspots,
  type FetchFileHotspotsOptions,
} from "../services/backend-client";

type FileMapSort =
  | "unresolved"
  | "fan-in"
  | "fan-out"
  | "symbols"
  | "flows"
  | "tests"
  | "path";

type FileKindFilter = "all" | "source" | "test" | "docs" | "config" | "generated";

interface FileMapPanelProps {
  repoName?: string;
  selectedPath?: string | null;
  onOpenFile: (path: string) => void;
}

const SORT_OPTIONS: Array<{ value: FileMapSort; label: string }> = [
  { value: "unresolved", label: "Unresolved" },
  { value: "fan-in", label: "Fan in" },
  { value: "fan-out", label: "Fan out" },
  { value: "symbols", label: "Symbols" },
  { value: "flows", label: "Flows" },
  { value: "tests", label: "Tests" },
  { value: "path", label: "Path" },
];

const KIND_OPTIONS: Array<{ value: FileKindFilter; label: string }> = [
  { value: "all", label: "All" },
  { value: "source", label: "Source" },
  { value: "test", label: "Test" },
  { value: "docs", label: "Docs" },
  { value: "config", label: "Config" },
  { value: "generated", label: "Generated" },
];

const formatSemanticKey = (value?: string): string =>
  (value || "unknown")
    .split(/[_-]+/)
    .filter(Boolean)
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(" ");

const compactCount = (value: number): string => {
  if (value >= 1000) return `${(value / 1000).toFixed(value >= 10000 ? 0 : 1)}k`;
  return String(value);
};

const defaultUnresolvedCount = (file: FileSummary): number =>
  file.defaultVisibleUnresolvedSourceSiteCount ?? file.unresolvedSourceSiteCount ?? 0;

const rawUnresolvedCount = (file: FileSummary): number =>
  file.rawUnresolvedSourceSiteCount ?? file.unresolvedSourceSiteCount ?? 0;

const testUnresolvedCount = (file: FileSummary): number =>
  file.testUnresolvedSourceSiteCount ?? 0;

const isTestFile = (file: FileSummary): boolean => file.kind === "test";

const riskClassName = (risk?: string): string => {
  switch (risk) {
    case "high":
      return "border-error text-error";
    case "medium":
      return "border-warning text-warning";
    default:
      return "border-border-subtle text-text-muted";
  }
};

export const FileMapPanel = ({
  repoName,
  selectedPath,
  onOpenFile,
}: FileMapPanelProps) => {
  const [sort, setSort] = useState<FileMapSort>("unresolved");
  const [kind, setKind] = useState<FileKindFilter>("all");
  const [searchQuery, setSearchQuery] = useState("");
  const [apiOnly, setApiOnly] = useState(false);
  const [changedOnly, setChangedOnly] = useState(false);
  const [unresolvedOnly, setUnresolvedOnly] = useState(false);
  const [highFanIn, setHighFanIn] = useState(false);
  const [highFanOut, setHighFanOut] = useState(false);
  const [data, setData] = useState<FileHotspotsResponse | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const requestOptions = useMemo<FetchFileHotspotsOptions>(
    () => ({
      repo: repoName,
      sort,
      limit: 200,
      kinds: kind === "all" ? undefined : [kind],
      apiOnly,
      changedOnly,
      unresolvedOnly,
      highFanIn,
      highFanOut,
    }),
    [apiOnly, changedOnly, highFanIn, highFanOut, kind, repoName, sort, unresolvedOnly],
  );

  useEffect(() => {
    if (!repoName) {
      setData(null);
      setError(null);
      return;
    }

    let cancelled = false;
    setIsLoading(true);
    setError(null);
    fetchFileHotspots(requestOptions)
      .then((nextData) => {
        if (!cancelled) setData(nextData);
      })
      .catch((err: unknown) => {
        if (!cancelled) {
          setError(err instanceof Error ? err.message : "Failed to load file map");
          setData(null);
        }
      })
      .finally(() => {
        if (!cancelled) setIsLoading(false);
      });

    return () => {
      cancelled = true;
    };
  }, [repoName, requestOptions]);

  const files = useMemo(() => {
    const rows = data?.files ?? [];
    const query = searchQuery.trim().toLowerCase();
    if (!query) return rows;
    return rows.filter((file) =>
      [
        file.path,
        file.kind,
        file.appLayer,
        file.functionalArea,
        file.language,
        file.risk,
      ]
        .filter(Boolean)
        .some((value) => String(value).toLowerCase().includes(query)),
    );
  }, [data?.files, searchQuery]);

  const totals = useMemo(() => {
    const rows = data?.files ?? [];
    return {
      unresolved: rows.filter((file) => defaultUnresolvedCount(file) > 0).length,
      changed: rows.filter((file) => file.changedSinceAnalyze).length,
      flows: rows.filter((file) => file.linkedFlowCount > 0).length,
      tests: rows.filter((file) => file.linkedTestCount > 0).length,
    };
  }, [data?.files]);

  return (
    <div
      className="flex min-h-0 flex-1 flex-col"
      data-testid="file-map-panel"
    >
      <div className="border-b border-border-subtle px-3 py-3">
        <div className="mb-3 flex items-center justify-between gap-3">
          <div className="min-w-0">
            <h3 className="press-eyebrow flex items-center gap-1.5 text-text-secondary">
              <Table className="h-3.5 w-3.5" />
              File Map
            </h3>
            <div className="mt-1 flex flex-wrap gap-2 font-mono text-[10px] text-text-muted">
              <span>{compactCount(data?.total ?? 0)} files</span>
              <span>{compactCount(totals.changed)} changed</span>
              <span>{compactCount(totals.unresolved)} unresolved</span>
              <span>{compactCount(totals.flows)} flows</span>
              <span>{compactCount(totals.tests)} tests</span>
              {data?.graph.stale && <span>stale</span>}
            </div>
          </div>
          {isLoading && (
            <Loader2 className="h-4 w-4 animate-spin text-text-muted" />
          )}
        </div>

        <div className="relative">
          <Search className="absolute top-1/2 left-2.5 h-3.5 w-3.5 -translate-y-1/2 text-text-muted" />
          <input
            type="text"
            placeholder="Search file map..."
            value={searchQuery}
            onChange={(event) => setSearchQuery(event.target.value)}
            className="w-full rounded border-[2px] border-border-default bg-inset py-2 pr-3 pl-8 font-mono text-xs text-text-primary placeholder:text-text-muted focus:border-border-strong focus:outline-none"
          />
        </div>

        <div className="mt-3 flex items-center gap-2">
          <ArrowDown className="h-3.5 w-3.5 text-text-muted" />
          <select
            value={sort}
            onChange={(event) => setSort(event.target.value as FileMapSort)}
            className="min-w-0 flex-1 rounded border-[2px] border-border-default bg-inset px-2 py-1.5 font-mono text-xs text-text-primary focus:border-border-strong focus:outline-none"
            aria-label="File map sort"
          >
            {SORT_OPTIONS.map((option) => (
              <option key={option.value} value={option.value}>
                {option.label}
              </option>
            ))}
          </select>
        </div>

        <div className="mt-3 flex flex-wrap gap-1.5">
          {KIND_OPTIONS.map((option) => (
            <button
              key={option.value}
              onClick={() => setKind(option.value)}
              aria-pressed={kind === option.value}
              className={`rounded px-2 py-1 font-mono text-[10px] transition-colors ${
                kind === option.value
                  ? "bg-accent text-text-inverse"
                  : "bg-base text-text-secondary hover:bg-surface hover:text-text-primary"
              }`}
            >
              {option.label}
            </button>
          ))}
        </div>

        <div className="mt-2 flex flex-wrap gap-1.5">
          {[
            {
              key: "changed",
              label: "Changed",
              active: changedOnly,
              onClick: () => setChangedOnly((value) => !value),
              testId: "file-map-filter-changed",
            },
            {
              key: "unresolved",
              label: "Unresolved",
              active: unresolvedOnly,
              onClick: () => setUnresolvedOnly((value) => !value),
              testId: "file-map-filter-unresolved",
            },
            {
              key: "api",
              label: "API",
              active: apiOnly,
              onClick: () => setApiOnly((value) => !value),
              testId: "file-map-filter-api",
            },
            {
              key: "fan-in",
              label: "High in",
              active: highFanIn,
              onClick: () => setHighFanIn((value) => !value),
              testId: "file-map-filter-high-fan-in",
            },
            {
              key: "fan-out",
              label: "High out",
              active: highFanOut,
              onClick: () => setHighFanOut((value) => !value),
              testId: "file-map-filter-high-fan-out",
            },
          ].map((filter) => (
            <button
              key={filter.key}
              onClick={filter.onClick}
              aria-pressed={filter.active}
              data-testid={filter.testId}
              className={`flex items-center gap-1 rounded px-2 py-1 font-mono text-[10px] transition-colors ${
                filter.active
                  ? "bg-base text-text-primary ring-1 ring-border-strong"
                  : "text-text-muted hover:bg-base hover:text-text-secondary"
              }`}
            >
              <Filter className="h-3 w-3" />
              {filter.label}
            </button>
          ))}
        </div>
      </div>

      {error && (
        <div className="border-b border-border-subtle px-3 py-2 text-xs text-error">
          {error}
        </div>
      )}

      <div className="scrollbar-thin min-h-0 flex-1 overflow-auto">
        <table className="w-full table-fixed border-collapse text-left">
          <thead className="sticky top-0 z-10 bg-surface">
            <tr className="border-b border-border-subtle font-mono text-[10px] text-text-muted">
              <th className="w-[46%] px-3 py-2 font-medium">Path</th>
              <th className="w-[17%] px-2 py-2 font-medium">Layer</th>
              <th className="w-[9%] px-1 py-2 text-right font-medium">Sym</th>
              <th className="w-[7%] px-1 py-2 text-right font-medium">In</th>
              <th className="w-[7%] px-1 py-2 text-right font-medium">Out</th>
              <th className="w-[7%] px-1 py-2 text-right font-medium">Unres</th>
              <th className="w-[7%] px-2 py-2 text-right font-medium">Links</th>
            </tr>
          </thead>
          <tbody>
            {files.map((file) => (
              <FileMapRow
                key={file.path}
                file={file}
                selected={selectedPath === file.path}
                onOpenFile={onOpenFile}
              />
            ))}
          </tbody>
        </table>

        {!isLoading && files.length === 0 && (
          <div className="px-3 py-6 text-center text-xs text-text-muted">
            No files matched
          </div>
        )}
      </div>
    </div>
  );
};

interface FileMapRowProps {
  file: FileSummary;
  selected: boolean;
  onOpenFile: (path: string) => void;
}

const FileMapRow = ({ file, selected, onOpenFile }: FileMapRowProps) => {
  const linkCount = file.linkedFlowCount + file.linkedTestCount;
  const defaultUnresolved = defaultUnresolvedCount(file);
  const rawUnresolved = rawUnresolvedCount(file);
  const testUnresolved = testUnresolvedCount(file);
  return (
    <tr
      data-testid="file-map-row"
      className={`border-b border-border-subtle/80 transition-colors ${
        selected ? "bg-base" : "hover:bg-base/70"
      }`}
    >
      <td className="min-w-0 px-3 py-2">
        <button
          type="button"
          onClick={() => onOpenFile(file.path)}
          title={file.path}
          className="flex w-full min-w-0 items-center gap-2 text-left"
        >
          {defaultUnresolved > 0 ? (
            <AlertTriangle className="h-3.5 w-3.5 shrink-0 text-warning" />
          ) : (
            <FileCode className="h-3.5 w-3.5 shrink-0 text-text-muted" />
          )}
          <span className="min-w-0 truncate font-mono text-[11px] text-text-primary">
            {file.path}
          </span>
          {file.changedSinceAnalyze && (
            <span className="shrink-0 rounded border border-border-subtle px-1 font-mono text-[9px] text-text-muted">
              changed
            </span>
          )}
          {isTestFile(file) && (
            <span className="shrink-0 rounded border border-border-subtle px-1 font-mono text-[9px] text-text-muted">
              Test File
            </span>
          )}
        </button>
      </td>
      <td className="min-w-0 px-2 py-2">
        <div className="min-w-0">
          <div className="truncate text-[11px] text-text-secondary">
            {formatSemanticKey(file.appLayer)}
          </div>
          <div className="truncate font-mono text-[9px] text-text-muted">
            {formatSemanticKey(file.functionalArea)}
          </div>
        </div>
      </td>
      <td className="px-1 py-2 text-right font-mono text-[10px] text-text-secondary">
        {compactCount(file.symbolCount)}
      </td>
      <td className="px-1 py-2 text-right font-mono text-[10px] text-text-secondary">
        {compactCount(file.inboundRefCount)}
      </td>
      <td className="px-1 py-2 text-right font-mono text-[10px] text-text-secondary">
        {compactCount(file.outboundRefCount)}
      </td>
      <td className="px-1 py-2 text-right font-mono text-[10px] text-text-secondary">
        {compactCount(defaultUnresolved)}
      </td>
      <td className="px-2 py-2 text-right">
        <span
          title={`flows ${file.linkedFlowCount}, tests ${file.linkedTestCount}, unresolved ${defaultUnresolved}, raw ${rawUnresolved}, test ${testUnresolved}, risk ${file.risk ?? "unknown"}, changed ${file.changedSinceAnalyze ? "yes" : "no"}`}
          className={`inline-flex min-w-8 items-center justify-end gap-1 rounded border px-1.5 py-0.5 font-mono text-[10px] ${riskClassName(file.risk)}`}
        >
          <GitBranch className="h-3 w-3" />
          {compactCount(linkCount)}
        </span>
      </td>
    </tr>
  );
};
