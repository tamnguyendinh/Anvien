# Anvien Analyze File Classification Metrics Plan

Date: 2026-06-01

Status: Complete

Companion files:

- Evidence ledger: [2026-06-01-anvien-analyze-file-classification-metrics-evidence.md](2026-06-01-anvien-analyze-file-classification-metrics-evidence.md)
- Benchmark ledger: [2026-06-01-anvien-analyze-file-classification-metrics-benchmark.md](2026-06-01-anvien-analyze-file-classification-metrics-benchmark.md)

## Master Rules

1. Use Anvien for graph/analyze discovery and impact checks before implementation edits.
2. Refresh the graph before graph-based evidence collection.
3. Do not treat docs, manifests, configs, reports, fixtures, scripts, and static assets as parser failures.
4. Reserve user-facing `unsupported` semantics for analyzer inputs that are genuinely unsupported.
5. Keep every non-parsed file count causal: each bucket must explain why the file was not parsed into ScopeIR.
6. Preserve parser-level `ErrUnsupportedLanguage`; this plan changes repo-level analyze metrics, not parser error semantics.
7. Preserve `failed` as real execution failure, such as read errors, parser failures, extractor failures, or pipeline errors.
8. Do not hide unknowns. If Anvien cannot classify a scanned file, report it as `unknown` with sample paths.
9. Keep JSON contracts explicit. If legacy compatibility fields remain, mark them as legacy and stop using them in human-facing output.
10. Run a full build before test validation.
11. If Web UI behavior changes, include Web build/test/e2e validation.
12. Record evidence and benchmark data as each implementation slice completes.
13. Run `anvien detect-changes --repo Anvien --scope all` before implementation commits.

## Goal

Replace the current aggregate analyze file metric named `unsupported` with a causal file classification model. Users should be able to tell whether a scanned file was parsed as code, indexed as a document, kept as metadata/config/static inventory, skipped because no code extractor exists, unknown, or failed.

## Problem

Analyze currently reports a repo-level shape like:

```text
files: scanned=<n> parsed=<n> unsupported=<n> failed=<n>
```

The `unsupported` number is not a true product classification. It is a parse-phase skip bucket exposed as repo-level truth. As a result, recognized Markdown docs, JSON manifests, report files, golden fixtures, config files, scripts, and static assets can appear as "unsupported" even though they are expected non-ScopeIR inputs.

This makes a healthy analyze run look like Anvien failed to understand a large part of the repo.

## Scope

In scope:

- analyze file metric contract;
- file classification buckets for scanned files;
- parse-phase and repo-level metric reconciliation;
- CLI analyze human output;
- CLI analyze JSON output;
- analyze benchmark JSON comparison code;
- graph accuracy code that records analyze file counts;
- tests for docs/config/static/script/unknown/true-unsupported/failure cases;
- README/RUNBOOK/TESTING/CHANGELOG updates after implementation behavior is proven.

Out of scope:

- adding new language parsers;
- changing scanner inclusion rules;
- removing document indexing;
- changing `ResolutionGap` unsupported syntax diagnostics;
- treating unsupported parser language errors as execution failures;
- Web UI changes unless the analyze result is surfaced there by existing contracts.

## Requirements

1. `parsedCode` must count files that produced ScopeIR.
2. `documents` must count files handled by document classification, including Markdown and metadata-only document formats.
3. `metadataOnly` must count manifests, configs, reports, schemas, query-health suites, generated JSON fixtures, golden files, and similar structured non-code inputs.
4. `scriptNoExtractor` must count recognized script-like files without ScopeIR extraction.
5. `staticAssets` must count HTML/CSS/static report or Web assets that are indexed but not parsed as code.
6. `unsupportedLanguage` must count only recognized code-like language inputs that are not supported by a ScopeIR extractor.
7. `unknown` must count files that no classifier can explain.
8. `failed` must count real processing failures and stay separate from all classification buckets.
9. Human output must not present an aggregate `unsupported=<large number>` as the primary truth.
10. JSON output must expose causal fields. If `files.unsupported` remains for compatibility, it must mean `unsupportedLanguage`, not the old aggregate of every non-ScopeIR file.
11. Machine-readable output must expose bounded samples or grouped detail for non-zero non-parsed buckets so users can verify why files landed there.
12. The sum of primary classification buckets plus failures must reconcile to `scanned`.
13. Evidence and benchmark ledgers must be updated at the end of each completed implementation phase.

## Invariants

1. Parser-level metrics remain available for parser internals.
2. Repo-level analyze metrics describe repo file inventory, not only parse-loop behavior.
3. Classification is deterministic for the same file path, scanner language, document kind, and parser outcome.
4. A file belongs to one primary repo-level bucket.
5. Adding docs to the repo may increase `documents`, but must not increase user-facing unsupported analyzer inputs.
6. Adding JSON fixtures may increase `metadataOnly`, but must not increase user-facing unsupported analyzer inputs.
7. Real parser/extractor errors must still fail or report `failed` according to existing pipeline semantics.

## Technical Direction

Add a small shared classifier near the analyze layer. It should accept enough facts to classify a scanned file without duplicating document or parser logic:

```text
scanner.File
document kind
has code extractor
file extension/path traits
parse result or parser unsupported error
```

The classifier should produce a stable bucket enum/string and optional reason. Analyze should aggregate those buckets into `FileClassificationMetrics`. CLI, benchmark comparison, and graph accuracy summaries should consume the same metrics instead of reconstructing their own counts.

Classification precedence:

1. `failed` for real read, parse, extraction, or pipeline failures.
2. `parsedCode` for files that produced ScopeIR.
3. `documents` for files recognized by document handling.
4. `metadataOnly` for structured non-code files such as manifests, configs, reports, schemas, query suites, and golden fixtures.
5. `scriptNoExtractor` for recognized script files without ScopeIR extraction.
6. `staticAssets` for HTML, CSS, static report, or Web asset files.
7. `unsupportedLanguage` for recognized code-like language inputs with no ScopeIR extractor.
8. `unknown` for the remaining scanned files that no classifier explains.

Expected compact human shape:

```text
files: scanned=<n> parsed_code=<n> failed=<n>
indexed: documents=<n> metadata=<n> scripts=<n> static=<n>
gaps: unsupported_language=<n> unknown=<n>
```

## Definition Of Done

The plan is complete when:

1. analyze human output shows causal file buckets instead of a misleading primary `unsupported` aggregate;
2. analyze JSON output exposes the causal bucket fields;
3. existing parser unsupported-language tests still pass;
4. new analyze tests prove docs/config/static/script/fixture files do not inflate unsupported analyzer inputs;
5. current Anvien repo analyze can explain all scanned non-parsed files through causal buckets;
6. build, targeted tests, analyze smoke, graph-health summary, and detect-changes evidence are recorded;
7. benchmark ledger records before/after inventory counts;
8. user-facing docs are updated only after behavior is proven.

## Phase Checklist

- [x] [P0-A] Establish baseline and owner evidence.
  - Goal: prove the current `unsupported` number is a metric semantics problem and identify the code owners before implementation.
  - Work Steps: refresh Anvien graph; inspect analyze output; use Anvien `query` and `context` for `parseFiles`, `hasExtractor`, and `FileMetrics`; classify current non-parsed inventory from `.anvien/graph.json`; record impact seed for the main owners.
  - Implementation Gate: no code edits in this phase; only plan/evidence/benchmark docs may change.
  - Acceptance: evidence records analyze counts, owner files, current bucket breakdown, and impact seed; benchmark records inventory counts.

- [x] [P0-B] Review plan readiness against file-role rules.
  - Goal: make the plan usable as the work controller rather than a loose discussion note.
  - Work Steps: re-read plan/evidence/benchmark roles; move command results and inventory metrics out of the plan; rewrite checklist items as mini-plans with Goal, Work Steps, Implementation Gate, and Acceptance; verify current graph counts after the planning commit; re-check known `unsupported` consumers.
  - Implementation Gate: do not start code changes until this review is complete.
  - Acceptance: plan status is ready for implementation; evidence records readiness review; benchmark records initial and latest inventory snapshots.

- [x] [P1-A] Define the repo-level file classification contract.
  - Goal: create the exact output model before changing parse behavior.
  - Work Steps: add a `FileClassificationMetrics` or equivalent struct in the analyze package; define bucket names and JSON names; decide which legacy fields remain; document reconciliation rules in tests; keep parser-level `parser.Metrics.Unsupported` separate.
  - Implementation Gate: run impact for `FileMetrics` and any new exported/shared metric type before editing.
  - Acceptance: compile-time contract exists, JSON field names are stable, bucket sum rules are represented in tests or helper assertions, and no CLI output changes are made before the contract is covered.

- [x] [P1-B] Implement the shared file classifier.
  - Goal: centralize the answer to "what kind of scanned file is this?" so CLI, analyze, and later reporting do not invent separate classifications.
  - Work Steps: implement a classifier function using scanner language, extractor support, document kind, extension/path traits, and parser outcome; add table tests for Markdown, spreadsheet, JSON manifest, JSON golden fixture, Go module, YAML config, PowerShell, shell script, HTML, CSS, true unsupported code-like input, unknown extension, and failure.
  - Implementation Gate: run impact for `hasExtractor`, `documentKind`, and any classifier owner before editing; do not change scanner inclusion behavior.
  - Acceptance: every current non-parsed Anvien file class maps to a causal bucket; unknown is reachable only through an explicit fallback; true unsupported language is test-covered separately from docs/config.

- [x] [P2-A] Reconcile parse-phase and repo-level metrics.
  - Goal: stop using parse-loop skips as the repo-level unsupported truth.
  - Work Steps: aggregate classification metrics from scanned files; update `parseFiles` or its caller so parser skips and parser `ErrUnsupportedLanguage` feed the right repo-level bucket; keep `failed` semantics unchanged; preserve parser pool metrics for parser internals.
  - Implementation Gate: run impact for `parseFiles` before editing and treat CRITICAL impact as a scoping warning.
  - Acceptance: analyze result can report parsed code, documents, metadata-only, scripts, static assets, unsupported language, unknown, and failed; current repo docs/config additions no longer raise user-facing unsupported analyzer input count.

- [x] [P2-B] Update analyze CLI output and JSON output.
  - Goal: make terminal and machine-readable analyze output match the new contract.
  - Work Steps: update `internal/cli/command.go` human format; update `--json` `files` object; inspect `internal/httpapi/analyze.go`, `internal/contracts/web_ui.go`, and generated Web contracts before deciding whether Web/API surfaces are unaffected; if legacy `unsupported` remains, make it an alias of `unsupportedLanguage` only; update CLI tests that assert analyze output.
  - Implementation Gate: inspect `internal/cli/command_test.go`, `internal/httpapi/analyze.go`, `internal/contracts/web_ui.go`, and existing benchmark comparison assumptions before editing output strings or JSON shape.
  - Acceptance: human output groups parsed/failed, indexed buckets, and gaps; JSON includes causal fields and bounded bucket detail; tests assert the new output and prevent the old primary `unsupported=<large>` line from returning.

- [x] [P3-A] Update downstream metric consumers.
  - Goal: prevent old aggregate semantics from leaking through benchmark and graph accuracy surfaces.
  - Work Steps: update `internal/cli/benchmark_command.go` to compare new fields; update `internal/graphaccuracy/access_candidate.go` away from `FilesUnsupported` or mark it legacy; update any structs/tests that serialize analyze file counts.
  - Implementation Gate: run impact for each consumer symbol before editing and preserve backwards compatibility only where needed.
  - Acceptance: benchmark comparison can read new analyze JSON; graph accuracy summaries expose causal fields or clearly marked legacy fields; no consumer silently treats docs/config as unsupported code.

- [x] [P3-B] Add focused regression tests.
  - Goal: lock the behavior that caused this plan so it cannot regress.
  - Work Steps: add analyze tests with mixed repo fixtures; include docs, JSON configs, testdata/golden JSON, scripts, static assets, true unsupported code-like input, unknown input, and real failure simulation where practical; update parser tests only if parser-level semantics require explicit preservation.
  - Implementation Gate: keep fixtures small and deterministic; do not broaden tests into unrelated scanner or graph projection behavior.
  - Acceptance: targeted tests fail on the old aggregate behavior and pass with causal classification; parser unsupported-language tests still prove parser internals work.

- [x] [P4-A] Validate implementation and record evidence.
  - Goal: prove the implementation works in the real repo and did not break analyze or graph quality.
  - Work Steps: run full build; run targeted Go tests; run analyze smoke with local binary; run graph-health summary; run detect-changes before commit; record pass/fail and notable failures in evidence.
  - Implementation Gate: do not commit implementation work until build/tests/analyze smoke/detect-changes evidence is recorded.
  - Acceptance: evidence ledger contains validation commands and results, failures and handling if any, detect-changes output, and commit hash after commit.

- [x] [P4-B] Record benchmark before/after and update user docs.
  - Goal: close the plan with synchronized metrics and user-facing documentation.
  - Work Steps: update benchmark ledger with latest/final bucket counts and deltas; update README/RUNBOOK/TESTING/CHANGELOG only for the final user-facing behavior; avoid copying implementation discussion into docs; mark plan complete when all acceptance criteria are met.
  - Implementation Gate: docs updates happen after behavior is validated, not before.
  - Acceptance: benchmark ledger shows before/after classification counts; docs explain analyze output accurately; plan status and checklist reflect completion.

## Risk Notes

- `parseFiles` and `hasExtractor` sit in the analyze pipeline and have CRITICAL impact in the baseline evidence. Edits must be narrow.
- CLI output changes can break tests, scripts, and user expectations. JSON compatibility needs an explicit decision.
- Benchmark comparison currently expects `scanned`, `parsed`, `unsupported`, and `failed`; it must not silently lose comparison coverage.
- Graph accuracy summaries currently include `FilesUnsupported`; this must be renamed, split, or clearly marked legacy.
- The working tree may contain untracked docs. Analyze scans working-tree files, so readiness counts can differ from the initial baseline.
