# Tree-sitter Go Module Drift Actual Status

Title: Tree-sitter Go Module Drift
Date: 2026-07-16
Status: Implementation Complete - Pending Closure
Companion plan: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-plan.md`
Companion evidence: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-evidence.md`
Companion benchmark: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-benchmark.md`

## Purpose

This file records the real current state before implementation.

Implementation must not start until the target scope has a completed status row, evidence IDs, and a downstream plan decision.

This file does not replace `evidence.md`. It classifies current state from evidence.

## Freshness / Refresh Rules

This actual-status file is a living current-state record, not a one-time P0 snapshot.

Update this file:

- after each completed implementation slice;
- before starting the next phase if repo state changed;
- whenever evidence changes a current-state classification;
- whenever the next phase's status assumptions, next action, or work steps need updating because reality differs from the previous status.

## Scope

Target scope:

- `.github/scripts/check-tree-sitter-upgrade-readiness.py`
- `.github/workflows/tree-sitter-upgrade-readiness.yml`
- `.github/dependabot.yml`
- `go.mod`
- Tree-sitter dependency drift report behavior

Out of scope:

- LadybugDB native runtime resolver
- Parser extraction semantics
- Anvien graph schema and LadybugDB storage
- Generated agent docs
- Tree-sitter dependency version bump unless explicitly required after checker repair

## Relationship / Impact Evidence

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| `go.mod` | `E0-P0A-FD1` | 0 | Parsed config file, no graph relationships reported. | Low scope warning for graph relationship count; high practical dependency impact if edited. |
| `.github/scripts/check-tree-sitter-upgrade-readiness.py` | `E0-P0A-FD2` | N/A | Current graph/file-detail does not expose this `.github` file as a target. Direct source inspection required. | Medium CI maintenance scope. |
| `.github/workflows/tree-sitter-upgrade-readiness.yml` | `E0-P0A-FD3` | N/A | Current graph/file-detail does not expose this `.github` file as a target. Direct source inspection required. | Medium CI maintenance scope. |
| `.github/dependabot.yml` | `E0-P0A-FD4` | N/A | Current graph/file-detail does not expose this `.github` file as a target. Direct source inspection required. | Medium dependency automation scope. |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Already behaves as required. | Preserve. Add evidence or tests only if needed. |
| `partial` | Some required behavior exists, but gaps remain. | Change only the missing parts. Preserve correct parts. |
| `wrong` | Current behavior, source, or contract is incorrect. | Replace with required behavior. Record the exact reason. |
| `missing` | Required behavior, source, or contract does not exist. | Implement the missing piece only. |
| `unbound` | Surface exists but is not wired to the real source, flow, or contract. | Bind to the real source only. Preserve approved surface. |
| `fake-or-stub` | Prototype, demo, mock, fallback, or placeholder data is being used as real behavior. | Remove fake behavior or replace it with an approved truthful state. |
| `blocked` | Source, authority, contract, or required evidence is unclear. | Stop. Do not implement until resolved. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| Checker script | Reads Go module data, reports 19 Tree-sitter modules, separates upstream core lag from actionable Go module updates, and no longer reads npm package dependencies. | Read Go module parser stack, classify drift, and emit a Markdown report without crashing. | correct | N/A | `E1-P1A-SRC1`, `E1-P1B-RUN1`, `E1-P1C-RUN1`, `E3-P3A-RUN2` | preserve / validate-only |
| Workflow | Names and issue text describe Tree-sitter Go module drift and parse uppercase module status rows. | Describe Tree-sitter Go module drift and parse the new report shape. | correct | N/A | `E2-P2A-SRC1`, `E2-P2A-SRC2`, `E2-P2B-VALIDATE1` | preserve / validate-only |
| Dependabot config | Root `gomod` monitoring exists with Tree-sitter Go module grouping; stale npm Tree-sitter runtime comments are removed. | Monitor root Go modules or record an explicit no-change decision; comments must match current dependency ownership. | correct | N/A | `E2-P2B-SRC1`, `E2-P2B-SRC2`, `E2-P2B-VALIDATE1` | preserve / validate-only |
| `go.mod` parser stack | Contains current Go module parser dependencies. | Preserve as source of truth for checker inventory. | correct | 0 related files | `E0-P0A-FD1`, `E0-P0A-SRC5` | inspect-only / preserve unless later evidence requires version bump |
| LadybugDB native resolver | Auto-resolver builds and analyzes successfully with latest native runtime. | Preserve; no change for this bug. | correct | N/A | `E0-P0A-RUN2`, `E0-P0A-RUN3` | preserve-only |
| Runtime parser/analyze behavior | `anvien analyze --force` passes after package build. | Preserve; checker fix must not regress runtime analyze. | correct | N/A | `E0-P0A-RUN3`, `E0-P0A-GRAPH1` | validate-only in P3-A |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-07-16 | baseline before implementation | Tree-sitter drift checker, workflow, Dependabot, `go.mod` | initial classification | `E0-P0A-GRAPH1`, `E0-P0A-FD1..E0-P0A-FD4`, `E0-P0A-SRC1..E0-P0A-SRC5`, `E0-P0A-RUN1..E0-P0A-RUN3` | P1 must replace npm checker logic; P2 must update stale CI/config language; P3 must validate runtime remains healthy. |
| R1 | 2026-07-16 | after P1/P2/P3 validation | Checker, workflow, Dependabot | checker `wrong -> correct`; workflow `wrong -> correct`; Dependabot `partial -> correct` | `E1-P1A-SRC1`, `E1-P1B-RUN1`, `E1-P1C-RUN1`, `E2-P2A-SRC1`, `E2-P2B-SRC1`, `E3-P3A-BUILD1`, `E3-P3A-RUN2` | Closure can proceed after detect-changes, report, supervisor, and commit. |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| `.github/scripts/check-tree-sitter-upgrade-readiness.py` | `go.mod` | source-of-truth dependency inventory | P1-A/P1-B/P1-C | edit | `E0-P0A-SRC1`, `E0-P0A-SRC5` | Remove npm runtime assumption; keep report explicit. |
| `.github/workflows/tree-sitter-upgrade-readiness.yml` | `.github/scripts/check-tree-sitter-upgrade-readiness.py` | report producer consumed by workflow output and issue update logic | P2-A | edit | `E0-P0A-SRC3` | Workflow text must match checker semantics. |
| `.github/dependabot.yml` | `go.mod` | dependency update automation for parser modules | P2-B | edit | `E0-P0A-SRC4`, `E0-P0A-SRC5` | Do not remove unrelated npm/web monitoring. |
| `go.mod` | `.github/scripts/check-tree-sitter-upgrade-readiness.py` | source-of-truth module inventory | P1/P2/P3 | inspect-only / preserve-only | `E0-P0A-FD1`, `E0-P0A-SRC5` | Do not bump parser versions unless new evidence requires it. |
| LadybugDB resolver scripts | package runtime validation | unrelated healthy runtime path | P3-A | preserve-only | `E0-P0A-RUN2`, `E0-P0A-RUN3` | Do not change. |

## Detailed Findings

### Checker Script

Current state:

- The script assumes `anvien/package.json` has `dependencies.tree-sitter`.
- The package no longer has that field.
- Running the script crashes before it can produce a report.

Required state:

```text
The script inventories Tree-sitter Go modules from go.mod/go list, reports actionable Go module updates, treats upstream Tree-sitter core ahead of Go bindings as informational unless policy says otherwise, and exits with clear status.
```

Evidence:

- `E0-P0A-SRC1`: npm dependency assumption in source.
- `E0-P0A-RUN1`: reproducible `KeyError: 'dependencies'`.

Relationship and impact:

- Related file count: N/A in current Anvien file-detail.
- Relationship summary: direct source inspection required.
- Impact note: CI maintenance bug, not runtime parser failure.

Classification:

`correct`

Allowed next action:

Preserve and validate-only.

Forbidden next action:

Do not reintroduce npm package dependency inventory or hide module inventory failures.

### Workflow

Current state:

- Workflow and issue text still describe `tree-sitter@0.25` npm readiness.
- It consumes a report shape generated by the stale checker.

Required state:

```text
Workflow title, comments, issue body, and comment parsing describe Tree-sitter Go module drift.
```

Evidence:

- `E0-P0A-SRC3`: stale workflow wording and issue title.

Relationship and impact:

- Related file count: N/A in current Anvien file-detail.
- Relationship summary: direct source inspection required.
- Impact note: scheduled CI and tracking issue quality risk.

Classification:

`correct`

Allowed next action:

Preserve and validate-only.

Forbidden next action:

Do not keep old npm `tree-sitter@0.25` wording after the checker changes to Go modules.

### Dependabot Config

Current state:

- Comments describe npm Tree-sitter grammar readiness and a pinned npm runtime.
- P0 evidence did not find a root `gomod` update entry.

Required state:

```text
Dependabot monitors root Go modules or records an explicit no-change decision, and comments match the current Go module parser stack.
```

Evidence:

- `E0-P0A-SRC4`: stale npm comments.
- `E0-P0A-SRC5`: Go module parser stack in `go.mod`.

Relationship and impact:

- Related file count: N/A in current Anvien file-detail.
- Relationship summary: direct source inspection required.
- Impact note: dependency automation can silently miss parser module updates.

Classification:

`correct`

Allowed next action:

Preserve and validate-only.

Forbidden next action:

Do not remove unrelated dependency monitoring entries without evidence.

### Runtime Parser Behavior

Current state:

- Package build and analyze passed with the current parser stack.

Required state:

```text
Runtime parser/analyze behavior remains healthy after CI maintenance fixes.
```

Evidence:

- `E0-P0A-RUN2`: package runtime build passed.
- `E0-P0A-RUN3`: built binary analyze passed.
- `E0-P0A-GRAPH1`: graph metrics captured.

Relationship and impact:

- Related file count: N/A.
- Relationship summary: validation surface only.
- Impact note: preserve-only, validate in P3.

Classification:

`correct`

Allowed next action:

Preserve and validate after implementation.

Forbidden next action:

Do not change parser runtime code for this bug unless validation reveals a separate runtime issue.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | Checker source now reads Go module data. | complete / preserve |
| P1-B | Checker report now separates actionable module drift from upstream core lag. | complete / preserve |
| P1-C | Focused tests now prevent regression to npm package dependency inventory. | complete / preserve |
| P2-A | Workflow text and issue parsing now match Go module drift report. | complete / preserve |
| P2-B | Dependabot now includes root `gomod` monitoring and current comments. | complete / preserve |
| P3-A | Runtime parser path remained healthy through full build and analyze. | complete / preserve |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status.
- [x] Each status has evidence IDs.
- [x] Each target file has relationship count evidence from `file-detail` when applicable, or an explicit `N/A` because current graph lookup does not expose `.github` files.
- [x] Phase Touch Map lists plan-relevant relationship files that can affect the current phase/slice.
- [x] Phase Touch Map defines touch mode for every plan-relevant relationship unit that may be affected.
- [x] Correct parts are marked preserve-only.
- [x] Partial, missing, wrong, unbound, and fake-or-stub parts have exact next actions.
- [x] Blockers are recorded, if any.
- [x] Next phase status assumptions, next action, and work steps have been updated from this status file when needed.
- [x] Status Refresh Log has an R0 baseline row.
- [x] If implementation has started, affected Current Status Matrix rows have been refreshed from latest evidence.
- [x] If refreshed statuses changed next work, only the stale next-phase status assumptions, next action, or work steps have been updated before the next phase.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

Proceed with P1. The implementation must fix the checker as a Go-module-aware maintenance script, then align workflow and Dependabot configuration. Do not treat upstream Tree-sitter core `v0.26.11` as a runtime bug by itself.
