# Anvien Test File Unresolved Separation Evidence Ledger

Date: 2026-06-02

Status: Active

Companion files:

- Plan: [2026-06-02-anvien-test-file-unresolved-separation-plan.md](2026-06-02-anvien-test-file-unresolved-separation-plan.md)
- Benchmark ledger: [2026-06-02-anvien-test-file-unresolved-separation-benchmark.md](2026-06-02-anvien-test-file-unresolved-separation-benchmark.md)

## Evidence Rules

1. Record Anvien command evidence for code/graph plan writing, plan review, and implementation slices.
2. Do not run Anvien only for doc-only commit ceremony.
3. Keep quantitative inventory and before/after counts in the benchmark ledger.
4. Record impact/blast-radius before editing graph builders, file projection, hotspot ranking, API contracts, or Web graph/file views.
5. Record generated output and UI checks only after the normal generation/build path creates them.
6. Record `anvien detect-changes --repo Anvien --scope all` before implementation commits.

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

## E0 - Plan Creation

Date: 2026-06-02

Status: recorded

Scope:

- Created the standard three-file plan set for separating test-file unresolved from default production unresolved signal.
- No implementation source files changed.
- No Anvien command was run for the initial doc-only file creation; the later code/graph plan review is recorded in E1 with Anvien evidence.

Source / command evidence:

| Check | Result |
|---|---|
| User problem statement | Test files only need to display as `Test File` and show what they test; unresolved details inside test files do not help the default production graph. |
| Prior analyze output from the current session | Top 5 unresolved hotspots were all test/e2e files, with unresolved counts from 856 to 1445. |
| Plan convention | This plan uses the standard `docs/plans/YYYY-MM-DD-<slug>/` directory with matching plan, evidence, and benchmark files. |

Impact / blast radius:

| Target | Result |
|---|---|
| Implementation code | Not run; no implementation edits in this planning step. |

Validation:

| Command | Result |
|---|---|
| File creation | Plan, evidence, and benchmark ledgers created under `docs/plans/2026-06-02-anvien-test-file-unresolved-separation/`. |

Failures / handling:

- None.

## E1 - Plan Review With Anvien

Date: 2026-06-02

Status: recorded

Scope:

- Reviewed whether the plan direction matches current code and graph behavior.
- Updated plan direction before implementation.
- No implementation source files changed.

Source / command evidence:

| Check | Result |
|---|---|
| `anvien analyze --force` | Pass. Graph refreshed with 818 files scanned, 598 parsed code files, 0 failed parses, 96,340 nodes, 131,828 relationships, 590 files with unresolved, and top 5 hotspots all test/e2e files. |
| `anvien query files "file projection unresolved hotspot ranking ResolutionGap" --repo Anvien` | Confirmed unresolved hotspot query is dominated by test/e2e files. |
| `anvien query files "test file classification kind appLayer backend_test e2e" --repo Anvien` | Confirmed file summaries already expose `kind=test` and test app layers such as `backend_test`, `api_test`, and `frontend_test`. |
| `anvien query files "web graph file map unresolved ResolutionGap node display" --repo Anvien` | Confirmed Web-facing file map/detail behavior depends on file unresolved summary fields. |
| `rg` source inspection | Found primary owners: `internal/semantic/app_layer.go`, `internal/filecontext/context.go`, CLI analyze/file-hotspots/graph-health commands, `internal/httpapi/file_context.go`, generated Web contracts, `FileMapPanel`, and `FileDetailPanel`. |

Plan review decisions:

| Decision | Evidence |
|---|---|
| Do not invent a new test-file detector first. | Existing `kind=test` and test app layers already exist in graph/file summaries. |
| P1-A should reuse/harden classification truth, not recreate it. | `filecontext.fileKind` derives `test` from app-layer values, and semantic app-layer tests already cover backend/API/frontend test paths. |
| Bucket separation must include default risk/warning semantics. | Web rows currently use `unresolvedSourceSiteCount` for warning icon, `Unres`, totals, and file detail unresolved display. |
| Web UI must not hard-code path checks. | Backend/file projection already owns classification; UI should consume backend fields. |
| Test-to-target relationships must remain visible. | `filecontext` already tracks reverse linked-test counts; plan must ensure test-file view can show tested targets too. |

Implementation evidence:

| File | Evidence |
|---|---|
| Plan | Updated Master Rules, Technical Direction, Requirements, P0-A, P1-A, P1-B, P1-C, P2-A, P2-B, P3-A, and P4-A. |
| Benchmark ledger | Updated B0 baseline to the latest analyze output and added raw/default risk separation as a target metric. |

Validation:

| Command | Result |
|---|---|
| Plan review | P0-A owner/baseline discovery is complete; implementation phases P1-A onward remain pending. |

Failures / handling:

- Initial review found baseline drift from the first plan draft; B0 was refreshed.
- Initial review found P1-A was too broad because classification already exists; P1-A was narrowed to reuse and harden existing backend truth.
