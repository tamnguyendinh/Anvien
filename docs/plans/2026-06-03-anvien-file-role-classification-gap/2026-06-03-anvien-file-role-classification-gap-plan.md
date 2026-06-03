# Anvien File Role Classification Gap Plan

Date: 2026-06-03

Status: Ready for implementation

Companion files:

- Evidence ledger: [2026-06-03-anvien-file-role-classification-gap-evidence.md](2026-06-03-anvien-file-role-classification-gap-evidence.md)
- Benchmark ledger: [2026-06-03-anvien-file-role-classification-gap-benchmark.md](2026-06-03-anvien-file-role-classification-gap-benchmark.md)

## Master Rules

1. Follow active workspace and repository instructions, including generated `AGENTS.md`; this plan records product work and validation, it does not replace repository rules.
2. Use Anvien for codebase analysis and impact checks before implementation edits.
3. Refresh graph evidence with `anvien analyze --force` before graph-based plan evidence or implementation evidence.
4. Run impact analysis before editing semantic classifiers, file projection summaries, graph-health/file-hotspot output, API contracts, or Web file views.
5. Treat HIGH/CRITICAL Anvien impact as blast-radius evidence, not an automatic edit ban.
6. This plan fixes classification semantics, not source-site resolution. Do not hide or delete raw unresolved facts just to reduce counts.
7. Keep generated `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**` as generated output only.
8. Run the full build before testing: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
9. If Web UI behavior changes, include a Web/e2e test or record why browser validation is not required.
10. Record evidence as each evidenced task completes.
11. Record benchmarkable inventory/count changes as each benchmarkable task completes.
12. Run `anvien detect-changes --repo Anvien --scope all` before implementation commits.
13. Commit each completed implementation slice after evidence and benchmark ledgers are updated.

## Goal

Add a clear file-role classification layer so files that are already recognizable as backend support/model/helper/config/adapter code are labeled that way in graph/file outputs.

The target behavior is:

- Anvien can explain the 17 raw-only unresolved files as recognized backend support/model/helper files;
- users do not interpret those files as "unknown" or "not understood";
- raw unresolved counts remain available for audit;
- default unresolved/actionable counts remain separate from non-actionable builtin/stdlib raw unresolved;
- CLI/API/Web surfaces can show concise role labels when file role is relevant.

## Problem

Current analyze/file projection output separates:

```text
unresolvedFiles=336
rawUnresolvedFiles=353
```

The 17-file difference is not a resolution failure for product code. Those files are recognized Go source files with `productionUnresolvedSourceSiteCount=0`; their raw unresolved sites are non-actionable builtins, standard library calls, or test-framework references.

The gap is semantic:

- the files have concrete identities and roles;
- current metadata has `kind`, `appLayer`, and `functionalArea`;
- there is no first-class role field that says `model`, `helper`, `config`, `storage_helper`, `adapter`, `fallback`, `test_helper`, or `analyzer_helper`;
- user-facing output can therefore make the files look less understood than they are.

## Scope

In scope:

- backend file-role taxonomy and classification rules;
- file summary fields and file-hotspot/file-context/graph-health output;
- semantic metadata/contract updates needed to carry role labels;
- Web file map/detail display only if role labels are surfaced there;
- tests proving the 17-file raw-only set is classified as known support/model/helper roles;
- benchmark tables tracking raw-only files, role coverage, and unknown-role count.

Out of scope:

- resolving Go builtins or standard library calls into real repo symbols;
- changing source-site resolution or ResolutionGap generation semantics;
- reducing raw unresolved counts by deleting graph facts;
- changing default unresolved ranking beyond role labeling;
- broad UI redesign or graph layout work;
- adding language parsers.

## Requirements

1. Add or expose a stable role classification for file summaries, preferably `fileRole`.
2. Preserve existing `kind`, `appLayer`, and `functionalArea` fields.
3. Classify the current 17 raw-only unresolved files into explicit roles with no `unknown` role remaining for that set.
4. Keep raw unresolved counts and default-visible unresolved counts separate.
5. Do not mark non-actionable builtin/stdlib raw-only files as production unresolved.
6. Role labels must be deterministic from path, package, symbol shape, and existing semantic metadata.
7. Role labels must be compact and API-safe strings, not prose.
8. If a role cannot be determined, use `unknown` and record count/sample evidence.
9. Add tests for role classification boundaries and the current file-role inventory.
10. If contracts change, regenerate Web contracts and update consumers/tests.
11. If Web UI displays role labels, use backend-provided role values instead of path-pattern checks.
12. Benchmark role coverage before and after each implementation slice.

## Invariants

1. Canonical symbol/source-site graph facts remain the source of truth.
2. File role is a semantic classification layer, not a replacement for graph relationships.
3. Raw unresolved is still raw unresolved, even when all sites are non-actionable.
4. Default-visible unresolved remains the actionable investigation signal.
5. File role must explain what the file is for; it must not claim unresolved references are fixed.
6. Existing app-layer/functional-area classifications remain valid and additive.
7. Generated contract changes must match backend source-of-truth fields.

## Technical Direction

Owner evidence points to these likely implementation areas:

- `internal/semantic/app_layer.go` owns existing app-layer classification and semantic application.
- `internal/semantic/functional_area.go` owns existing functional-area classification.
- `internal/filecontext/context.go` owns `FileSummary`, file-context aggregation, hotspot sorting fields, and unresolved bucket display counts.
- `internal/cli/command.go`, graph-health/file-hotspot commands, HTTP/MCP file summary surfaces, and Web contracts consume file summary fields.
- `anvien-web/src/components/FileMapPanel.tsx` and `anvien-web/src/components/FileDetailPanel.tsx` are likely UI consumers if role labels become visible.

Prefer a small, shared role classifier near `internal/semantic` or `internal/filecontext` instead of scattering path checks across CLI/API/Web. Candidate role values:

```text
model
contract_model
helper
storage_helper
config
adapter
fallback_adapter
test_helper
analyzer_helper
parser_model
runtime_model
unknown
```

Initial mapping examples:

| File | Target role |
|---|---|
| `internal/group/types.go` | `contract_model` |
| `internal/scopeir/facts.go` | `model` or `parser_model` |
| `internal/scopeir/range.go` | `model` or `parser_model` |
| `internal/scopeir/sort_keys.go` | `helper` |
| `internal/repo/paths.go` | `storage_helper` |
| `internal/repo/settings.go` | `config` |
| `internal/repo/runtime_config.go` | `config` |
| `internal/lbugnative/runner.go` | `adapter` |
| `internal/lbugnative/runner_default.go` | `fallback_adapter` |
| `internal/testutil/path.go` | `test_helper` |
| `internal/frameworks/frameworks.go` | `analyzer_helper` |

## Definition Of Done

The plan is complete when:

1. file summaries expose a stable role field or equivalent backend-owned classification;
2. the 17 raw-only unresolved files are classified as known backend support/model/helper roles;
3. raw-only unresolved output can be explained as non-actionable stdlib/builtin/test-framework sites plus file role, not "unknown file";
4. CLI/API/Web contract behavior is updated only where needed and covered by tests;
5. role coverage benchmarks show the raw-only 17-file set has zero unknown roles;
6. full build, focused tests, analyze smoke, graph-health/file-hotspot validation, and detect-changes evidence are recorded;
7. implementation work is committed after evidence and benchmark ledgers are updated.

## Phase Checklist

- [x] [P0-A] Establish baseline and owner evidence.
  - Goal: prove the issue is a classification gap rather than a resolution failure and identify likely code owners.
  - Work Steps: run `anvien analyze --force`; inspect current `unresolvedFiles` and `rawUnresolvedFiles`; enumerate raw-only files; inspect file-context samples for representative raw-only files; run Anvien query/context/impact for semantic classifier owners.
  - Implementation Gate: no product code edits in this phase.
  - Acceptance: evidence records current counts, raw-only file list summary, source-site classification, owner files, and CRITICAL impact warnings for classifier edits.

- [ ] [P1-A] Define the file-role contract.
  - Goal: decide the exact field name, role enum, JSON shape, and compatibility strategy before implementation.
  - Work Steps: inspect `FileSummary`, graph-health/file-hotspot JSON, HTTP/MCP file summary payloads, Web generated contracts, and semantic definitions; choose whether role belongs in `internal/semantic` or `internal/filecontext`; define role values and labels; write tests that encode role strings and unknown fallback.
  - Implementation Gate: run impact for `FileSummary` and any new/exported role type before editing.
  - Acceptance: a stable backend role contract exists, old fields are preserved, and tests prove invalid/unknown role fallback.

- [ ] [P1-B] Implement backend file-role classification.
  - Goal: classify known backend support/model/helper/config/adapter files without changing unresolved semantics.
  - Work Steps: implement a shared classifier using path, package, file name, symbol shape, and existing semantic metadata; add table tests for the 17 raw-only files plus boundary examples; keep the classifier deterministic and scoped; do not change raw/default unresolved bucket calculations.
  - Implementation Gate: no CLI/API/Web output changes until backend classification tests pass.
  - Acceptance: the current 17-file raw-only set maps to non-unknown roles, and role classification does not alter raw/default unresolved counts.

- [ ] [P2-A] Surface role classification through file summaries and CLI/graph quality output.
  - Goal: make CLI/API graph-quality users see file role where file identity is discussed.
  - Work Steps: add role fields to `FileSummary`; update file-hotspots, file-context, graph-health, analyze file projection output only where useful; update CLI tests and JSON tests; ensure raw-only files can be listed with role and non-actionable counts.
  - Implementation Gate: preserve existing JSON fields unless an explicit compatibility decision is recorded.
  - Acceptance: command output can report raw-only files with role labels, and tests cover the new field in JSON/human output where emitted.

- [ ] [P2-B] Update API contracts and Web display if role reaches Web.
  - Goal: keep generated contracts and Web consumers aligned with backend file-role metadata.
  - Work Steps: update `internal/contracts/web_ui.go` if file summary shape changes; regenerate `anvien-web/src/generated/anvien-contracts.ts`; update `FileMapPanel` and `FileDetailPanel` to use backend role labels if displayed; add unit tests and e2e validation if visible behavior changes.
  - Implementation Gate: do not add Web path-pattern role detection; use backend role only.
  - Acceptance: Web contract and consumers compile, role display is covered by tests, and e2e evidence is recorded if UI behavior changes.

- [ ] [P3-A] Validate role coverage and raw unresolved semantics.
  - Goal: prove the change solves the classification gap without hiding graph-quality evidence.
  - Work Steps: run full build; run focused Go tests; run Web tests if applicable; run `anvien analyze --force`; run `file-hotspots --sort raw-unresolved --unresolved-only --limit 0 --json`; verify raw-only file count, role coverage, and raw/default unresolved counts; run graph-health summary; record results in evidence and benchmark ledgers.
  - Implementation Gate: no commit until validation and benchmark ledgers are updated.
  - Acceptance: role coverage target is met, raw/default unresolved counts remain explainable, and validation passes or failures are recorded with handling.

- [ ] [P4-A] Detect changes, commit, and close.
  - Goal: close the implementation slice with traceable scope.
  - Work Steps: run `anvien detect-changes --repo Anvien --scope all`; review affected files/flows; record detect-changes evidence; commit the completed slice; update plan status/checklist/evidence/benchmark with final hash and closure state.
  - Implementation Gate: detect-changes must match intended classification/API/UI scope.
  - Acceptance: implementation commit exists, plan artifacts record closure evidence, and no required follow-up remains for this plan.

## Risk Notes

- `ClassifyAppLayer` and `ClassifyFunctionalArea` have CRITICAL upstream impact through analyze, CLI, graphaccuracy, and semantic application flows. Edits must be narrow and test-backed.
- Adding fields to `FileSummary` can touch CLI, MCP, HTTP/API, generated contracts, and Web consumers.
- The fix must not conflate file-role classification with source-site resolution. Builtins and stdlib calls may remain raw unresolved unless a separate resolution plan changes that behavior.
