# Anvien Analyze File Classification Metrics Plan

Date: 2026-06-01

Status: Planned

Companion files:

- Evidence ledger: [2026-06-01-anvien-analyze-file-classification-metrics-evidence.md](2026-06-01-anvien-analyze-file-classification-metrics-evidence.md)
- Benchmark ledger: [2026-06-01-anvien-analyze-file-classification-metrics-benchmark.md](2026-06-01-anvien-analyze-file-classification-metrics-benchmark.md)

## Master Rules

1. Use Anvien for graph/analyze discovery and impact checks before implementation edits.
2. Refresh the graph before graph-based evidence collection.
3. Do not treat docs, manifests, configs, reports, scripts, and static assets as parser failures.
4. Reserve `unsupported` semantics for files or languages that are genuinely unsupported by the code analyzer.
5. Keep every skipped-file count causal: each bucket must explain why a file was not parsed as ScopeIR.
6. Preserve `failed` as real execution failure, such as read errors, parse failures, or extractor errors.
7. Do not hide unknowns. If Anvien cannot classify a file, report it as `unknown` with sample paths.
8. Keep JSON contracts explicit. If a compatibility field remains, mark it as legacy and do not use it in human-facing output.
9. Run a full build before test validation.
10. If Web UI output changes, add or update e2e coverage.
11. Record evidence and benchmark data as each implementation slice completes.
12. Run `anvien detect-changes --repo Anvien --scope all` before implementation commits.

## Goal

Replace the current aggregate `unsupported` file metric with a causal file classification model so users can understand what Anvien did with every scanned file.

The target analyze result should answer:

- how many code files were parsed into ScopeIR;
- how many document files were indexed as documents;
- how many manifests, configs, reports, and fixtures were indexed as metadata-only files;
- how many scripts or static assets were intentionally not parsed as code;
- how many files are truly unsupported analyzer inputs;
- how many files are unknown or unclassified;
- how many files actually failed.

## Problem

Current analyze output reports:

```text
files: scanned=807 parsed=596 unsupported=211 failed=0
```

This is technically produced by the parse phase, but it is not a useful product-level classification. The current `unsupported=211` combines recognized documents, structured metadata, configs, fixtures, reports, scripts, and static files with truly unsupported parser inputs.

That makes the output look like Anvien failed to understand 211 files. The graph evidence shows that is not true. Most of those files are known non-code or metadata files that should not be parsed through ScopeIR at all.

## Root Cause

The root cause is metric ownership, not dead code.

Current behavior:

```text
scan phase: finds every repo file that should be indexed
document phase: enriches markdown/spreadsheet-like files as document metadata
parse phase: only parses languages with ScopeIR extractors
metrics/output: reports every file not parsed by ScopeIR as unsupported
```

The bad join happens because parse-phase skip counts are surfaced as repo-level file truth.

## Current Baseline

Fresh Anvien evidence from this repo:

```text
807 scanned
596 parsed
211 currently reported as unsupported
0 failed
```

Breakdown of the 211 files:

| Bucket | Count | Meaning |
|---|---:|---|
| `document:markdown` | 105 | Markdown docs recognized by document indexing |
| `document:spreadsheet` | 1 | Spreadsheet-like document metadata |
| `.json` | 92 | Reports, manifests, contracts, fixtures, configs, and test data |
| `.mod` | 3 | Go module manifests in repo/tooling folders |
| `.html` | 2 | Static Web/report pages |
| `.ps1` | 2 | PowerShell scripts |
| `.cli` | 1 | Dockerfile-style descriptor |
| `.conf` | 1 | Server/runtime config |
| `.css` | 1 | Static Web asset |
| `.sh` | 1 | Shell script |
| `.web` | 1 | Dockerfile-style descriptor |
| `.yaml` | 1 | YAML config |

Expected target interpretation for the same inventory:

| Target bucket | Count |
|---|---:|
| Parsed code files | 596 |
| Document files | 106 |
| Metadata/config/report/fixture files | 99 |
| Script files without ScopeIR extractor | 3 |
| Static Web/assets | 3 |
| True unsupported analyzer inputs | 0 |
| Unknown/unclassified files | 0 |
| Failed files | 0 |

## Scope

In scope:

- Add a shared file classification model for analyze metrics.
- Reconcile scan, document, parse, and file projection classifications.
- Replace human-facing `unsupported` aggregate output with causal buckets.
- Add JSON fields for the new bucketed metrics.
- Update CLI analyze output and any machine-readable analyze/benchmark surfaces that consume file metrics.
- Update graph accuracy and benchmark comparison consumers that read `FilesUnsupported`.
- Add tests proving documents/configs are not counted as unsupported analyzer failures.
- Add evidence and benchmark ledger updates after each implementation slice.
- Update user-facing docs after implementation is complete.

Out of scope:

- Adding new language parsers.
- Changing which files the scanner includes.
- Removing document indexing.
- Removing parser-level `ErrUnsupportedLanguage`; true parser unsupported still exists as a lower-level error.
- Treating `ResolutionGap` unsupported syntax diagnostics as file parse unsupported metrics.

## Classification Contract V1

Every scanned file should land in one primary classification bucket:

| Field | Meaning |
|---|---|
| `scanned` | Total files accepted by the scanner |
| `parsedCode` | Files parsed into ScopeIR by a language extractor |
| `documents` | Files indexed by document handling, including markdown and binary document metadata |
| `metadataOnly` | Structured metadata, manifests, configs, reports, contracts, fixtures, and test data |
| `scriptNoExtractor` | Script-like files that are recognized but not ScopeIR parsed |
| `staticAssets` | Static UI/report assets that are indexed as files but not parsed as code |
| `unsupportedLanguage` | Recognized code language with no ScopeIR extractor |
| `unknown` | Scanned file that no classifier can explain |
| `failed` | File that hit an actual read/parse/extraction failure |

Human-facing output should prefer grouped language:

```text
files: scanned=807 parsed_code=596 failed=0
indexed: documents=106 metadata=99 scripts=3 static=3
gaps: unsupported_language=0 unknown=0
```

JSON output may keep a legacy field during migration, but the legacy aggregate must not be the primary displayed truth.

## Implementation Plan

- [x] [P0] Capture baseline evidence with fresh Anvien analyze, Anvien owner discovery, graph-health summary, and graph snapshot classification.
- [ ] [P1] Define `FileClassificationMetrics` and bucket names in the analyze layer.
- [ ] [P2] Add a shared classifier that maps `scanner.File`, document kind, language support, extension, and parser result into one causal bucket.
- [ ] [P3] Rework `parseFiles` and analyze metric assembly so ScopeIR parser skips no longer become repo-level `unsupported` truth.
- [ ] [P4] Update CLI analyze human output and JSON output to show causal buckets.
- [ ] [P5] Update benchmark comparison and graph accuracy consumers that currently read `unsupported`.
- [ ] [P6] Add tests for docs, JSON/config/testdata, scripts, static assets, true unsupported language, unknown file, and real failure cases.
- [ ] [P7] Run full build, targeted tests, analyze smoke, graph-health summary, and detect-changes.
- [ ] [P8] Update README/RUNBOOK/TESTING/CHANGELOG only with final user-facing behavior after the implementation is complete.

## Implementation Ownership

| Area | Owner |
|---|---|
| Analyze metrics contract | `internal/analyze/analyze.go` |
| Parser-level unsupported language error | `internal/parser` |
| Document classification | `internal/documents/documents.go` |
| CLI analyze output | `internal/cli/command.go` |
| Analyze benchmark comparison | `internal/cli/benchmark_command.go` |
| Graph accuracy summaries | `internal/graphaccuracy/access_candidate.go` |
| Analyze tests | `internal/analyze/analyze_test.go` |
| Parser unsupported tests | `internal/parser/*_test.go` |

## Validation Plan

- Full build:

```bash
powershell -ExecutionPolicy Bypass -File anvien-launcher/build.ps1
```

- Targeted tests:

```bash
go test ./internal/analyze ./internal/parser ./internal/documents ./internal/cli ./internal/graphaccuracy -count=1
```

- Smoke analyze:

```bash
.\anvien\bin\anvien.exe analyze --force --name Anvien --json
```

- Graph quality:

```bash
.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json
```

- Pre-commit impact/change evidence:

```bash
.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all
```

## Success Criteria

1. Recognized docs/configs/reports/fixtures no longer inflate user-facing `unsupported`.
2. Every non-parsed file count has a causal bucket and representative sample paths.
3. Current Anvien repo baseline can be explained without an aggregate `unsupported=211`.
4. True unsupported language remains visible when it actually exists.
5. `failed=0` still means no analyze execution failure.
6. JSON, benchmark, and graph accuracy consumers use the new explicit fields or explicitly marked legacy fields.
7. Tests prevent Markdown, JSON, config, and static files from regressing back into generic unsupported counts.
