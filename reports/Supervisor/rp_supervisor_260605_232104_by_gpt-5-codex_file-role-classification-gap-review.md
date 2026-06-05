# Supervisor Report: File Role Classification Gap Plan Review

Verdict: REJECT

## Metadata
- Report file: `rp_supervisor_260605_232104_by_gpt-5-codex_file-role-classification-gap-review.md`
- Review time: 260605 232104 Asia/Bangkok
- Reviewer: gpt-5-codex
- Repo/project: Anvien
- Scope reviewed: `docs/plans/2026-06-03-anvien-file-role-classification-gap`
- Claim reviewed: The plan/evidence/benchmark set is complete and the implementation commit `c063bcb feat: add backend file group classification` closes first-class `fileGroup=backend_support_model_helper` across backend, CLI/API/MCP, generated Web contracts, and Web UI verification.
- Authority used: latest user review request; `AGENTS.md`; reviewed plan/evidence/benchmark files; current source; current Anvien analyze output; current focused tests.
- Related artifacts: `docs/plans/2026-06-03-anvien-file-role-classification-gap/*`, commit `c063bcb`, current HEAD `6267b00`.

## Executive Summary
- Problem: The completed plan claims full closure of file-group classification and UI evidence.
- Decision: REJECT. The core implementation is present and focused Go/Web unit tests pass, but the completion artifact has blocking evidence gaps: the recorded Playwright command does not pass in the current repo, and the graph File-node enrichment decision contradicts the plan's own branch rule.
- Required outcome: Fix the plan/evidence or implementation so the recorded validation is reproducible and the graph-node/file-projection boundary is explicitly consistent with the approved scope.

## Blocking Findings

### [HIGH] Recorded Web e2e evidence is not reproducible with the command in the ledger
File: `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-evidence.md:604`
Issue: The evidence ledger says `npm run test:e2e -- file-map-test-unresolved.spec.ts` passes, but the same command fails in the current repo because the frontend server is not running.
Evidence: Current run of `npm run test:e2e -- file-map-test-unresolved.spec.ts` from `anvien-web` failed with `page.goto: net::ERR_CONNECTION_REFUSED at http://127.0.0.1:5228/...`. `anvien-web/package.json:16` defines `test:e2e` as only `playwright test`. `anvien-web/playwright.config.ts:20` through `anvien-web/playwright.config.ts:45` has no `webServer` stanza, and `anvien-web/package.json:10` is the separate `dev` server command for port `5228`.
Why this blocks acceptance: The plan requires Web/e2e tests for visible UI behavior changes in `2026-06-03-anvien-file-role-classification-gap-plan.md:145` through `2026-06-03-anvien-file-role-classification-gap-plan.md:147` and validates that scope in `2026-06-03-anvien-file-role-classification-gap-plan.md:356` through `2026-06-03-anvien-file-role-classification-gap-plan.md:379`. A non-reproducible e2e command is not acceptable closure evidence.
Fix direction: Either make the claimed command self-contained by adding/using a server-starting e2e workflow, or update the evidence ledger to record the exact required preconditions and rerun result.
Re-review evidence required: A fresh passing e2e run for `file-map-test-unresolved.spec.ts` with the exact command/preconditions documented in the ledger.

### [HIGH] Graph File-node enrichment decision contradicts the plan's own branch rule
File: `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-plan.md:257`
Issue: The checklist says that if File nodes already receive semantic identity properties during analyze, `fileGroup` must be added to that same path. The evidence ledger says File nodes already carry `appLayer` and `functionalArea`, but then leaves direct graph-node `fileGroup` enrichment as a follow-up.
Evidence: Plan branch rule is at `2026-06-03-anvien-file-role-classification-gap-plan.md:257` through `2026-06-03-anvien-file-role-classification-gap-plan.md:259`. Evidence ledger states current File nodes already carry `appLayer` and `functionalArea` at `2026-06-03-anvien-file-role-classification-gap-evidence.md:590` through `2026-06-03-anvien-file-role-classification-gap-evidence.md:595`. Current `.anvien/graph.json` for `internal/repo/runtime_config.go` shows `filePath`, `functionalArea`, `language`, and `name` around `.anvien/graph.json:35737`, but no `fileGroup` property on that File node. Source does add `FileSummary.FileGroup` in `internal/filecontext/context.go:115` through `internal/filecontext/context.go:121` and computes it in `internal/filecontext/context.go:391` through `internal/filecontext/context.go:399`, so the gap is graph-node persistence, not missing file-summary implementation.
Why this blocks acceptance: The reviewed plan claims the group is visible in graph/file projection/API/CLI/Web outputs in `2026-06-03-anvien-file-role-classification-gap-plan.md:51`, and its own checklist branch requires graph-node enrichment when semantic identity properties already exist on File nodes. Recording that as a follow-up does not close the approved invariant.
Fix direction: Either implement `fileGroup` enrichment on File nodes, or revise the plan/authority so file projection is explicitly the only graph-facing source for this slice and the File-node branch no longer requires enrichment.
Re-review evidence required: Fresh analyze evidence showing either File nodes contain `fileGroup`, or a corrected authority decision that explicitly narrows graph scope to file projection plus updated evidence.

### [MEDIUM] Benchmark ledger's "Latest" group counts are stale relative to current repo reality
File: `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-benchmark.md:132`
Issue: The benchmark ledger records `backend_support_model_helper` full group files as `42`, but current analyze reports `47`.
Evidence: Current `anvien analyze . --force` passed with `files scanned=1382`, `parsed_code=682`, `failed=0`, and emitted `fileProjection.group key="backend_support_model_helper" label="Backend support/model/helper files" files=47 unresolved=2531 ...`. The benchmark ledger records `Latest before corrected group implementation` as `42` at `2026-06-03-anvien-file-role-classification-gap-benchmark.md:132` through `2026-06-03-anvien-file-role-classification-gap-benchmark.md:143`.
Why this blocks acceptance: As a historical measurement this is explainable, but the artifact is still marked `Status: completed` and uses "Latest" language. Current review cannot use those counts as current acceptance evidence.
Fix direction: Update the benchmark/evidence wording to distinguish historical closure counts from current moving repo counts, or refresh the benchmark if the artifact is intended to represent current repo state.
Re-review evidence required: A refreshed benchmark row or explicit historical-count note with current analyze evidence.

## Source-Level Clearance Notes
- `internal/semantic/file_group.go`: clear for core classifier presence. It defines `FileGroupBackendSupportModelHelper` and exact label at `internal/semantic/file_group.go:11` through `internal/semantic/file_group.go:56`.
- `internal/filecontext/context.go`: clear for file-summary/projection path. It exposes `FileSummary.FileGroup` at `internal/filecontext/context.go:115` through `internal/filecontext/context.go:121`, computes group from backend semantic logic at `internal/filecontext/context.go:391` through `internal/filecontext/context.go:399`, and aggregates groups at `internal/filecontext/context.go:602` through `internal/filecontext/context.go:645`.
- `anvien-web/src/components/FileMapPanel.tsx`: clear for Web row display. It maps generated labels at `anvien-web/src/components/FileMapPanel.tsx:62` through `anvien-web/src/components/FileMapPanel.tsx:69` and renders `file.fileGroup` at `anvien-web/src/components/FileMapPanel.tsx:390` through `anvien-web/src/components/FileMapPanel.tsx:397`.
- `anvien-web/src/components/FileDetailPanel.tsx`: clear for Web detail display. It maps labels at `anvien-web/src/components/FileDetailPanel.tsx:48` through `anvien-web/src/components/FileDetailPanel.tsx:55` and renders `Group` before `Role` at `anvien-web/src/components/FileDetailPanel.tsx:574` through `anvien-web/src/components/FileDetailPanel.tsx:578`.

## Evidence Checked
Passed:
- `git branch --contains c063bcb`: current `master` contains the implementation commit.
- `anvien analyze . --force`: pass; current output emits direct `fileProjection.group key="backend_support_model_helper"` with `files=47`.
- `go test ./internal/semantic ./internal/filecontext ./internal/contracts ./internal/cli ./internal/httpapi ./internal/mcp`: pass.
- `npm test -- FileMapPanel.test.tsx FileDetailPanel.test.tsx`: pass, 2 files and 8 tests.
- Source inspection confirms backend classifier, file summary, generated contract consumers, and Web display code exist.
- Verification freshness: fresh/current for this review.

Failed:
- `npm run test:e2e -- file-map-test-unresolved.spec.ts`: fail; frontend server refused connection at `127.0.0.1:5228`.
- Graph-node enrichment check: current graph File nodes have app-layer/functional-area identity but no `fileGroup` property for the inspected target.
- Verification freshness: fresh/current for this review.

Not run:
- Full build was not rerun inside this review because the e2e evidence failure and graph-scope conflict already block acceptance. A full build had passed earlier in the current work session, but it is not used as PASS evidence for this review.

## Invariant Closure
- affected invariant: first-class file identity group `backend_support_model_helper` must be computed by backend authority and visible through claimed graph/file projection/API/CLI/Web surfaces.
- sibling surfaces checked: semantic classifier, file summary/list aggregation, CLI/API/MCP references by source scan, generated Web contract references by source scan, Web FileMap/FileDetail display, current analyze file projection output, focused Go tests, Web unit tests, claimed Web e2e command, current graph File-node properties.
- residual unverified same-invariant surfaces: direct graph File-node `fileGroup` property remains unclosed or requires authority narrowing; Web e2e validation is not currently reproducible with the documented command.

## Required Fix List For Resubmission
1. Make the e2e evidence reproducible: document/start the required Vite server or add/use an e2e command that starts it, then rerun and record the exact passing command.
2. Resolve the graph-node branch conflict: either add `fileGroup` to File node enrichment or revise the authority/plan so file projection is explicitly sufficient.
3. Refresh or clarify benchmark counts so "Latest" does not conflict with current analyze output.

## Overall Evaluation
The implementation is substantially present and the core backend/Web unit evidence is good. The artifact cannot be accepted as complete because two closure claims are not proven: the visible UI e2e evidence does not replay from the recorded command, and the graph/File-node branch in the plan is internally inconsistent with the evidence. This is a review failure of completion evidence and invariant closure, not a finding that the whole feature is absent.
