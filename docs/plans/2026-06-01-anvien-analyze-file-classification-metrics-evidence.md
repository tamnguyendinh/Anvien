# Anvien Analyze File Classification Metrics Evidence Ledger

Date: 2026-06-01

Status: Planned

Companion files:

- Plan: [2026-06-01-anvien-analyze-file-classification-metrics-plan.md](2026-06-01-anvien-analyze-file-classification-metrics-plan.md)
- Benchmark ledger: [2026-06-01-anvien-analyze-file-classification-metrics-benchmark.md](2026-06-01-anvien-analyze-file-classification-metrics-benchmark.md)

## Evidence Rules

1. Record Anvien commands used to discover the owner and baseline.
2. Keep quantitative benchmark tables in the benchmark ledger.
3. For code changes, record impact/blast-radius before edits.
4. Preserve the difference between parser-level unsupported language and repo-level file classification.
5. Record failures and their handling.
6. Record `anvien detect-changes --repo Anvien --scope all` before each implementation commit.

## E0 - Baseline Analyze Refresh

Date: 2026-06-01

Status: completed

Command:

```bash
.\anvien\bin\anvien.exe analyze --force --name Anvien --json
```

Result:

| Field | Value |
|---|---:|
| Files scanned | 807 |
| Files parsed | 596 |
| Files currently reported unsupported | 211 |
| Files failed | 0 |
| Graph nodes | 95814 |
| Graph relationships | 131151 |
| File projection files | 807 |
| File projection dependency edges | 15816 |
| File projection unresolved files | 588 |

Evidence interpretation:

- The graph refresh succeeded with no failed files.
- The issue is the meaning of the `unsupported` aggregate, not analyze failure.

## E1 - Anvien Owner Discovery

Date: 2026-06-01

Status: completed

Commands:

```bash
.\anvien\bin\anvien.exe query "unsupported files parseFiles hasExtractor analyze metrics" --repo Anvien
.\anvien\bin\anvien.exe context "parseFiles" --repo Anvien
.\anvien\bin\anvien.exe context "hasExtractor" --repo Anvien
.\anvien\bin\anvien.exe context "FileMetrics" --repo Anvien
```

Findings:

| Owner | Evidence |
|---|---|
| `internal/analyze/analyze.go` | Anvien found `parseFiles`, `hasExtractor`, and `FileMetrics` in the analyzer functional area. |
| `parseFiles` | Starts at `internal/analyze/analyze.go:785`; increments `metrics.Unsupported++` when `!hasExtractor(file.Language)`. |
| `hasExtractor` | Starts at `internal/analyze/analyze.go:874`; returns true only for languages with ScopeIR extractors or script-container extractors. |
| `FileMetrics` | `internal/analyze/analyze.go:153` currently contains only `Scanned`, `Parsed`, `Unsupported`, and `Failed`. |
| CLI output | `internal/cli/command.go:263` prints `files: scanned=%d parsed=%d unsupported=%d failed=%d`. |

Root-cause code facts:

```text
internal/analyze/analyze.go:795-797
if !hasExtractor(file.Language) {
    metrics.Unsupported++
    continue
}
```

```text
internal/analyze/analyze.go:820-823
metrics.Failed++
if errors.Is(err, parser.ErrUnsupportedLanguage) {
    metrics.Unsupported++
    continue
}
```

```text
internal/cli/command.go:263-268
files: scanned=%d parsed=%d unsupported=%d failed=%d
```

## E2 - Document Classification Evidence

Date: 2026-06-01

Status: completed

Inspected file:

```text
internal/documents/documents.go
```

Facts:

| Area | Evidence |
|---|---|
| Document phase | `Apply` iterates scanned files and selects files where `documentKind(file.Path) != ""`. |
| Markdown | `.md` and `.mdx` return `markdown`; markdown files get section/link graph enrichment. |
| Spreadsheet | `.xls`, `.xlsx`, `.csv`, `.tsv`, and related extensions return `spreadsheet`; these get document metadata. |
| Binary marker | Non-markdown documents are marked metadata/binary instead of parsed as ScopeIR. |

Interpretation:

- Markdown and spreadsheet files are recognized by Anvien.
- Counting these files as generic parser `unsupported` is a metric presentation problem.

## E3 - Current 211-File Breakdown

Date: 2026-06-01

Status: completed

Method:

- Read `.anvien/graph.json`.
- Selected `File` nodes.
- Treated languages supported by `hasExtractor` as parsed-code candidates.
- Grouped remaining file nodes by `documentKind` or extension.

Breakdown:

| Bucket | Count |
|---|---:|
| `document:markdown` | 105 |
| `document:spreadsheet` | 1 |
| `unknown-ext:.json` | 92 |
| `unknown-ext:.mod` | 3 |
| `unknown-ext:.html` | 2 |
| `unknown-ext:.ps1` | 2 |
| `unknown-ext:.cli` | 1 |
| `unknown-ext:.conf` | 1 |
| `unknown-ext:.css` | 1 |
| `unknown-ext:.sh` | 1 |
| `unknown-ext:.web` | 1 |
| `unknown-ext:.yaml` | 1 |

Top-level directory distribution:

| Top path | Count |
|---|---:|
| `reports` | 79 |
| `docs` | 62 |
| `internal` | 28 |
| `anvien-claude-plugin` | 14 |
| `anvien-web` | 7 |
| `anvien` | 3 |
| `anvien-launcher` | 3 |
| `scripts` | 2 |
| remaining root/singleton paths | 13 |

Sample non-document files:

```text
anvien-claude-plugin/hooks/hooks.json
anvien-web/index.html
anvien-web/package.json
anvien-web/src/index.css
contracts/web-ui/anvien-web-contract.schema.json
docker-compose.yaml
docker/web-nginx.conf
docs/query-health/2026-05-23-anvien-skill-system-upgrade-suite.json
go.mod
internal/providers/golang/testdata/go_scopeir_signature.golden.json
```

Interpretation:

- The current `211` is explainable.
- It should become document/config/script/static/unknown buckets, not one `unsupported` number.

## E4 - Impact Seed For Future Implementation

Date: 2026-06-01

Status: completed

Commands:

```bash
.\anvien\bin\anvien.exe impact symbol "parseFiles" --uid "Function:internal/analyze/analyze.go:parseFiles#4" --repo Anvien --direction upstream
.\anvien\bin\anvien.exe impact symbol "hasExtractor" --uid "Function:internal/analyze/analyze.go:hasExtractor#1" --repo Anvien --direction upstream
.\anvien\bin\anvien.exe impact symbol "FileMetrics" --uid "Struct:internal/analyze/analyze.go:FileMetrics" --repo Anvien --direction upstream
```

Summary:

| Target | Risk | Impact summary |
|---|---|---|
| `parseFiles` | CRITICAL | 5 impacted symbols, 4 affected files, 39 affected processes. |
| `hasExtractor` | CRITICAL | 4 impacted symbols, 3 affected files, 26 affected processes. |
| `FileMetrics` | LOW | 1 impacted symbol, 1 affected file, 0 affected processes. |

Affected files observed:

| Target | Files |
|---|---|
| `parseFiles` | `cmd/access-candidate-audit/main.go`, `internal/analyze/analyze.go`, `internal/cli/command.go`, `internal/graphaccuracy/access_candidate.go` |
| `hasExtractor` | `internal/analyze/analyze.go`, `internal/cli/command.go`, `internal/graphaccuracy/access_candidate.go` |
| `FileMetrics` | `internal/analyze/analyze.go` |

Implementation note:

- `parseFiles` and `hasExtractor` are CRITICAL blast-radius targets because they sit in the analyze pipeline. Future code edits must be narrow and validated with build, tests, analyze smoke, graph-health, and detect-changes.

## E5 - Planning Decision

Date: 2026-06-01

Status: completed

Decision:

- Create a new implementation plan instead of editing code immediately.
- Use Anvien evidence for the plan because this is graph/analyze behavior, not a doc-only commit gate.
- Do not change generated AI context or generated skill files as part of this planning slice.

Rationale:

- The user-facing problem is a metric semantics problem.
- The implementation must preserve parser-level unsupported language errors while removing misleading repo-level aggregate output.
