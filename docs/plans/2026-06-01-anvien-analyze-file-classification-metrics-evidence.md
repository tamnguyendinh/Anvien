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

## Evidence Template

Use this template for implementation phases:

```text
## E<n> - <Phase/Task>

Date:

Status:

Scope:

- ...

Source / command evidence:

| Check | Result |
|---|---|
| ... | ... |

Impact / blast radius:

| Target | Result |
|---|---|
| ... | ... |

Implementation evidence:

| File | Evidence |
|---|---|
| ... | ... |

Validation:

| Command | Result |
|---|---|
| ... | ... |

Failures / handling:

- ...

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | ... |

Commit:

- `<hash> <subject>`
```

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

## E6 - Readiness Review And Plan Structure Correction

Date: 2026-06-01

Status: completed

Scope:

- Re-read the plan, evidence, and benchmark files against the required file roles.
- Refresh Anvien after the planning commit so implementation readiness uses current graph evidence.
- Check current code references that still consume `unsupported`.

Source / command evidence:

| Check | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --json` | Pass. `files.scanned=813`, `files.parsed=596`, `files.unsupported=217`, `files.failed=0`; graph `95889` nodes and `131238` relationships; file projection `813` files and `15828` dependency edges. |
| Current non-parsed inventory grouping from `.anvien/graph.json` | `111` Markdown docs, `1` spreadsheet document, `92` JSON, `3` `.mod`, `2` HTML, `2` PowerShell, and one each `.cli`, `.conf`, `.css`, `.sh`, `.web`, `.yaml`. |
| `rg` for `FilesUnsupported`, `FileMetrics`, and output strings | Owners still include `internal/analyze/analyze.go`, `internal/cli/command.go`, `internal/cli/benchmark_command.go`, `internal/graphaccuracy/access_candidate.go`, and tests under `internal/analyze` / `internal/cli`. |
| `git status -sb` | Branch is ahead of origin by one plan commit. Three unrelated GitNexus benchmark plan files are untracked and were not touched. |

Decision:

- The original plan direction was correct but not detailed enough to drive implementation.
- The plan file has been rewritten so every checklist item is a mini-plan with Goal, Work Steps, Implementation Gate, and Acceptance.
- Long inventory measurements belong in the benchmark ledger; command/results and readiness facts belong in this evidence ledger.

Readiness result:

- After the structure correction, the plan is ready to implement from `[P1-A]` if the user wants to proceed.
- Implementation must start with fresh impact checks for the exact symbols/files edited, because `parseFiles` and `hasExtractor` are CRITICAL blast-radius targets.

## E7 - Plan Re-Review Tightening

Date: 2026-06-01

Status: completed

Scope:

- Re-review the plan after applying the file-role rules.
- Check whether known downstream surfaces outside CLI should be named before implementation.

Source / command evidence:

| Check | Result |
|---|---|
| `rg` over analyze/CLI/graphaccuracy/httpapi/contracts for file metrics and unsupported references | Confirmed direct old aggregate consumers in `internal/analyze/analyze.go`, `internal/cli/command.go`, `internal/cli/benchmark_command.go`, `internal/graphaccuracy/access_candidate.go`, and tests. `internal/httpapi/analyze.go` uses total scanned files for repo stats but does not directly expose the old unsupported aggregate. `internal/contracts/web_ui.go` contains `scanned-not-extracted` language coverage metadata, which must be inspected before declaring Web/contracts unaffected. |

Decision:

- Add `[P0-B]` to the plan so readiness review is tracked in the checklist.
- Add classification precedence so implementation cannot choose bucket order ad hoc.
- Tighten JSON compatibility: legacy `files.unsupported`, if kept, must mean `unsupportedLanguage`, not old non-ScopeIR aggregate.
- Expand `[P2-B]` to inspect HTTP/API and Web contract surfaces before deciding they are unaffected.

## E8 - Implementation Impact And Scope

Date: 2026-06-01

Status: completed

Scope:

- Run fresh graph/analyze evidence before editing.
- Run impact checks for the symbols that own the metric contract, parse skip behavior, and document classification reuse.

Source / command evidence:

| Check | Result |
|---|---|
| `.\anvien\bin\anvien.exe analyze --force --name Anvien --json` | Pass. `files.scanned=814`, `files.parsed=596`, `files.unsupported=218`, `files.failed=0`; graph `95904` nodes and `131259` relationships. |
| Anvien query for analyze metric owners | Confirmed `internal/analyze/analyze.go`, `internal/cli/command.go`, `internal/cli/benchmark_command.go`, `internal/graphaccuracy/access_candidate.go`, and tests as direct owners/consumers. |
| `internal/httpapi/analyze.go` inspection | HTTP analyze repo stats use total scanned files and do not expose the old unsupported aggregate. |
| `internal/contracts/web_ui.go` inspection | Web contract language coverage metadata is separate from analyze file metrics; no generated Web contract change needed. |

Impact / blast radius:

| Target | Result |
|---|---|
| `parseFiles` | CRITICAL; 5 impacted symbols, 4 affected files, 39 affected processes. |
| `hasExtractor` | CRITICAL; 4 impacted symbols, 3 affected files, 26 affected processes. |
| `FileMetrics` | LOW; 1 impacted symbol, 1 affected file, 0 affected processes. |
| `documentKind` | CRITICAL through `documents.Apply -> analyze.Run -> CLI/graphaccuracy`; scoped to adding shared `documents.Kind` without changing document selection semantics. |

Implementation boundary:

- Keep edits narrow to analyze metrics, classifier, CLI output, benchmark comparison, graphaccuracy summary, and focused tests.
- Do not change scanner inclusion behavior, document indexing behavior, parser registry behavior, or Web UI behavior.

## E9 - Classification Contract And Consumer Implementation

Date: 2026-06-01

Status: completed

Implementation evidence:

| File | Evidence |
|---|---|
| `internal/analyze/file_classification.go` | Added shared deterministic classifier, bucket constants, bounded samples, and reconciliation helper. |
| `internal/analyze/analyze.go` | Expanded `FileMetrics`; repo-level metrics now aggregate scanned files from parse outcome plus document/path/language classification; `files.unsupported` is a legacy alias of `unsupportedLanguage`. |
| `internal/documents/documents.go` | Added exported `documents.Kind` wrapper so analyze classification reuses document logic instead of duplicating it. |
| `internal/cli/command.go` | Human output now prints `parsed_code`, indexed buckets, and gap buckets; JSON output emits the full causal `FileMetrics` object. |
| `internal/cli/benchmark_command.go` | Benchmark comparison reads new causal file metric fields while still reading legacy fields. |
| `internal/graphaccuracy/access_candidate.go` | Analyze summary exposes causal fields and keeps legacy parsed/unsupported aliases. |
| `internal/analyze/file_classification_test.go` | Added regression coverage for docs, spreadsheets, JSON manifests/golden files, Go modules, YAML, PowerShell, shell, HTML, CSS, unsupported code-like input, unknown input, and failed precedence. |
| `internal/analyze/analyze_test.go`, `internal/cli/command_test.go` | Updated pipeline, benchmark, CLI human output, JSON output, and benchmark-compare assertions. |

Validation:

| Command | Result |
|---|---|
| `go test ./internal/analyze ./internal/documents ./internal/cli ./internal/graphaccuracy` | Pass. |
| `go test ./internal/analyze ./internal/documents ./internal/cli ./internal/graphaccuracy -count=1` | Pass after applicable build. |

Failures / handling:

- None in the focused implementation tests.

## E10 - Build, Analyze Smoke, And Graph Health Validation

Date: 2026-06-01

Status: completed

Validation:

| Command | Result |
|---|---|
| `go build ./...` | Failed on existing intentionally non-buildable fixture packages under `anvien/test/fixtures/...` and C fixture files; not caused by this change. |
| `go build ./cmd/... ./internal/...` | Pass. |
| `npm --prefix anvien run build` | Failed because `anvien\bin\lbug_shared.dll` was held by another running process; did not indicate compile failure. |
| `go build -trimpath -o .tmp\anvien-classification.exe .\cmd\anvien` | Pass, with existing tree-sitter Swift C warning only. |
| `.tmp\anvien-classification.exe analyze --force --name Anvien --json --benchmark-json .tmp\analyze-file-classification-final.json` | Pass. `files.scanned=816`, `parsedCode=598`, `documents=113`, `metadataOnly=99`, `scriptNoExtractor=3`, `staticAssets=3`, `unsupportedLanguage=0`, `unknown=0`, `failed=0`; graph `96159` nodes and `131644` relationships. |
| `.tmp\anvien-classification.exe graph-health summary --repo Anvien --json` | Pass. Graph is fresh at commit `81adef9`; graph `96159` nodes and `131644` relationships; file layer `816` files and `590` unresolved files. |
| `go test ./cmd/... ./internal/... -count=1` | Failed only in `internal/lbugschema` because `baseline/phase-1-contract-freeze/ladybugdb-graph-contract.json` is missing. All other packages in the run passed. |
| `$packages = go list ./cmd/... ./internal/... \| Where-Object { $_ -ne 'github.com/tamnguyendinh/anvien/internal/lbugschema' }; go test $packages -count=1` | Pass. |

Failures / handling:

- Root `go build ./...` remains unsuitable as product build validation because it includes analysis fixtures that are intentionally not buildable as packages.
- Package runtime build was blocked by an active process holding `lbug_shared.dll`; a separate `.tmp` binary was built and used for analyze smoke without stopping user/editor-owned runtime processes.
