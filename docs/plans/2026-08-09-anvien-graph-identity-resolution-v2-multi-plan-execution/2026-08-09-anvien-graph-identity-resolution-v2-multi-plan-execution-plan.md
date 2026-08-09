# Anvien Graph Identity Resolution v2 Multi-Plan Execution Plan

## Metadata

- Date: `2026-08-09`
- Status: `active / owner-authorized / P0 in progress`
- Plan: `docs/plans/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution-plan.md`
- Evidence: `docs/plans/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution-evidence.md`
- Benchmark: `docs/plans/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution-benchmark.md`
- Actual status: `docs/plans/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution/2026-08-09-anvien-graph-identity-resolution-v2-multi-plan-execution-actual-status.md`

## Goal

Execute the seven-child graph-identity and TypeScript-resolution campaign in the frozen dependency order. Deliver production code, direct behavior tests, evidence, measured benchmarks, per-slice commits, successor handoffs, and final cross-surface acceptance without contaminating the target repository.

## Rules

- The user's `triển khai multi plan` instruction is the current Owner authorization to activate the candidate roadmap and begin implementation.
- The seven child plans under `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/` remain the technical execution authorities; this plan coordinates order and status only.
- Complete and refresh P0 actual status before opening each child. A later child cannot repair an earlier invariant.
- Implement the exact source slice IDs and order: P1-P7, 98 implementation slices total; commit after every accepted slice.
- Run `anvien analyze E:\Anvien --force` before graph queries and run `anvien detect-changes --repo Anvien --scope all` before every implementation commit.
- Before editing an existing function, class, method, exported symbol, or shared contract, run `anvien file-detail` and upstream impact analysis and record the blast-radius warning.
- Code first; update tests only after the behavior is implemented. Tests must prove trigger/process/observable behavior and failure paths.
- Run the repository full build before final validation. If a runtime/UI slice is touched, Docker and real runtime/browser evidence are mandatory.
- Keep temporary logs and fixtures under repo-local `.tmp`; never write to `E:\cheapapp.org` or copy target files into this repo.
- Update the matching child checklist, evidence ledger, benchmark ledger, and actual-status file immediately after each slice.
- Supervisor review is an acceptance gate; a rejected slice reopens only that slice before campaign progress continues.

## Problem

The existing graph uses conflated declaration/symbol IDs, overwrites duplicate nodes, loses binding-pattern leaves, treats export visibility as a scalar flag, resolves barrels without an export table, and misclassifies ambient/external declarations. The prepared roadmap decomposes the remediation into seven dependent child plans; implementation evidence is still absent at the start of this execution.

## Scope

In scope:

- Activate the roadmap after recording the Owner authorization and preserve the legacy plan as reference-only.
- Child 01: identity contract, lossless occurrences, relationship identity, strict mutation, and shadow-v2.
- Child 02: versioned persistence, opaque readers, atomic generations, and v2 cutover.
- Child 03: recursive TypeScript binding-pattern extraction and projection.
- Child 04: first-class TypeScript export facts and syntax projection.
- Child 05: module export tables, barrel/re-export traversal, cycles, and ambiguity.
- Child 06: declaration universe, external materialization, capability modes, and truthful diagnostics.
- Child 07: determinism, parity, target boundary, downstream isolation, and performance acceptance.
- All required evidence/benchmark/status ledgers and coder/supervisor reports.

## Non-Goals

- Rewriting the approved technical order or inventing an eighth child plan.
- Copying, modifying, or storing artifacts in `E:\cheapapp.org` beyond the explicitly permitted in-place validation commands.
- Unrelated cleanup, UI redesign, transport changes, or speculative infrastructure.
- Claiming campaign completion from a build-only or stale test result.

## Requirements

1. The canonical identity contract keeps `DeclarationID` and `SymbolID` as distinct, deterministic, versioned values with explicit byte/column encoding.
2. Canonical graph construction fails closed on conflicting identity, duplicate relationship, missing endpoint, or invalid generation data.
3. Persistence publishes immutable validated generations atomically and all readers validate the versioned manifest.
4. TypeScript bindings and exports are represented as first-class facts independent of type inference, with structured diagnostics for unsupported syntax.
5. Module resolution and export resolution are separate, memoized, provenance-preserving operations with explicit ambiguity/cycle/budget outcomes.
6. External declarations come from an explicit, versioned declaration universe and remain isolated from default in-repo traversal/rename behavior.
7. Final acceptance proves deterministic output, JSON/Ladybug parity, target-specific oracles, and measured capacity/performance.

## Acceptance Criteria

- All seven child plans have refreshed actual-status baselines and their implementation checklists/evidence ledgers are updated from current repo reality.
- All 98 implementation slices either pass their gates with a commit and qualified handoff or are explicitly blocked with evidence.
- Full build and nearest-boundary behavior tests pass for the completed scope; no required validation is substituted by stale tests.
- Child 07 Supervisor review records PASS with no residual unverified same-invariant surface, or the execution plan records a concrete blocker.
- Final worktree status, detect-changes output, benchmark values, commit hashes, and target-boundary checks are recorded.

## Checklist

- [ ] P0-A: Refresh execution baseline and activate authority.
  - Goal: classify the current repo and explicitly bind the Owner authorization to the roadmap before production edits.
  - Work Steps: refresh Anvien; inspect the roadmap, legacy metadata, Child 01 status, current source inventory, and open Supervisor reports; record file-detail counts and impact warnings; update the execution actual-status and authority pointers; commit the documentation checkpoint.
  - Implementation Gate: P0 actual-status contains current evidence IDs, the authorization is recorded, the legacy pointer is unambiguous, and no production file has been edited.
  - Acceptance: the campaign has one active authority and Child 01 is the only open implementation child.

### P1: Child 01 — graph identity contract and strict graph construction

- Phase Goal: complete Child 01's 11 ordered source slices and emit a qualified handoff to Child 02.
- Phase Boundary: edit only the ownership allowlists in the Child 01 plan; production scope is exactly the files named by P1-A through P1-E; no binding/export/ambient behavior.
- Dependencies: P0-A and the activated roadmap.
- Ordered Slice List: `P1-A`, `P1-B`, `P1-C0`, `P1-C0A`, `P1-C0B`, `P1-C`, `P1-D`, `P1-D1`, `P1-D2`, `P1-D3`, `P1-E`.
- Work Steps: execute each child-authority slice in order; run impact/file-detail before edits; implement code then tests; run build and nearest behavior checks; update ledgers/status; detect changes; commit; refresh Child 02 actual-status before handoff.
- Implementation Gate: all prior Child 01 slices are accepted and committed; active v1 remains preserved until Child 02 cutover.
- Acceptance: Child 01 Pn-A/Pn-B/Pn-C pass, `E2-PNC-NEXTSTATUS1` and `E2-PNC-HANDOFF1` are recorded with the exact child slug, and the successor status is fresh.

### P2: Child 02 — versioned persistence and v2 cutover

- Phase Goal: complete the 42 Child 02 slices, migrate all readers, and atomically cut over to identity v2.
- Dependencies: qualified Child 01 handoff and fresh Child 02 P0.
- Ordered Slice List: preserve source order `P2-A` through `P2-G` exactly as listed in the Child 02 plan.
- Work Steps: execute each slice with its reader-matrix ownership; validate immutable generations, fault paths, parity, and consumer isolation; record evidence/benchmarks and commit after each slice; refresh Child 03 status.
- Implementation Gate: Child 01 handoff is accepted and all manifest/reader fields are explicit.
- Acceptance: Child 02 closure and qualified handoff pass; all readers validate the v2 generation contract.

### P3: Child 03 — TypeScript binding-pattern extraction

- Phase Goal: complete the 17 binding slices, including loop contexts and the strengthened `.map()`/shadowing gates.
- Dependencies: qualified Child 02 handoff and fresh P3 P0.
- Ordered Slice List: preserve source order `P3-A`, `P3-B`, `P3-B1`, `P3-B2`, `P3-B2A`, `P3-C`, `P3-C1`, `P3-C1A`…`P3-C1I`, `P3-C2`.
- Work Steps: implement recursive facts before projection; prove type-independent emission, structured unsupported diagnostics, six `.map()` sites, shadowing, and zero import-count delta; update ledgers/status, detect changes, commit, and refresh Child 04 status.
- Implementation Gate: Child 02 handoff and `P3-B2A` are accepted before projection slices open.
- Acceptance: binding oracle rows are complete and Child 03 handoff is qualified.

### P4: Child 04 — first-class TypeScript export semantics

- Phase Goal: complete the 15 export-fact slices and preserve the P4/P5 ownership boundary.
- Dependencies: qualified Child 03 handoff and fresh P4 P0.
- Ordered Slice List: preserve source order `P4-A`, `P4-B`, `P4-B1`, `P4-C`, `P4-C1`, `P4-C1A`…`P4-C1I`, `P4-C2`.
- Work Steps: implement module/import/export facts, meaning lanes, direct export counts, and adapters; run behavior tests and full build; commit each slice; refresh Child 05 status.
- Implementation Gate: P3 handoff is accepted and P4 facts have a single producer owner.
- Acceptance: P4 facts and direct-export oracle pass; qualified handoff is recorded.

### P5: Child 05 — module export and re-export resolution

- Phase Goal: complete the four resolver slices with terminal symbol, cycle, ambiguity, barrel, and path-accounting proofs.
- Dependencies: qualified Child 04 handoff and fresh P5 P0.
- Ordered Slice List: `P5-A`, `P5-B`, `P5-C`, `P5-D`.
- Work Steps: implement module/export tables and memoized traversal; verify zero-physical barrel and unchanged path counts; update ledgers/status, detect changes, and commit; refresh Child 06 status.
- Implementation Gate: P4 handoff is accepted and P5 derived-state ownership is explicit.
- Acceptance: resolver semantic vector passes and qualified handoff is recorded.

### P6: Child 06 — ambient/external resolution and diagnostics

- Phase Goal: complete the six declaration-universe slices with capability-aware structured outcomes and downstream isolation.
- Dependencies: qualified Child 05 handoff, inspect-only Child 02 manifest denominator, and fresh P6 P0.
- Ordered Slice List: `P6-A`, `P6-B`, `P6-C1`, `P6-C2`, `P6-C3`, `P6-D`.
- Work Steps: implement declaration loading/security/catalog/materialization/outcome projection; verify Promise/Math and negative capability fixtures; update ledgers/status, detect changes, and commit; refresh Child 07 status.
- Implementation Gate: P5 handoff is accepted and declaration capability modes are versioned.
- Acceptance: authoritative diagnostic matrix and external isolation oracle pass; qualified handoff is recorded.

### P7: Child 07 — cross-surface acceptance and target validation

- Phase Goal: run final acceptance over all six predecessor handoffs, target oracles, parity, determinism, and performance.
- Dependencies: all six qualified handoffs and fresh P7 P0.
- Ordered Slice List: `P7-A`, `P7-B`, `P7-C`.
- Work Steps: validate without repairing upstream implementation; run full build/Docker/runtime checks where required; measure benchmark targets; return failures to the owning child; update roadmap terminal status.
- Implementation Gate: all predecessor handoffs are accepted and current.
- Acceptance: terminal Supervisor PASS, no residual same-invariant gaps, and campaign closure evidence is complete.

- [ ] P8-A: Final supervisor acceptance loop.
  - Goal: independently review the complete campaign claim.
  - Work Steps: inspect source/diff, run Anvien impact/detect evidence, verify builds/tests/runtime/benchmarks, write a supervisor report, and reopen only failed slices.
  - Implementation Gate: P1-P7 are either accepted or explicitly blocked.
  - Acceptance: Supervisor PASS or a durable blocker report.
- [ ] P8-B: Remove dead work created by execution.
  - Goal: leave no stale helper, fixture, generated output, or temporary artifact from the campaign.
  - Work Steps: review final diff and status, remove only campaign-created dead artifacts, rerun relevant checks, and record cleanup evidence.
  - Implementation Gate: P8-A has no unresolved implementation rejection.
  - Acceptance: final diff contains only accepted scope artifacts.
- [ ] P8-C: Close the execution plan.
  - Goal: record final build, validation, benchmarks, detect-changes, commits, and worktree state.
  - Work Steps: run final full build, runtime/target validation, update all ledgers and roadmap, run detect-changes, commit, and verify clean/known status.
  - Implementation Gate: P8-A and P8-B pass or a blocker is recorded.
  - Acceptance: closure evidence is complete and no required work remains.

## Risk Notes

- Child 01 touches high/critical fan-in graph and resolver surfaces; blast-radius warnings are scope controls, not edit bans.
- The repository currently has no `Docs/SPEC` cluster; the accepted problem report, child plans, and newly ratified contract document are the available authority, and this limitation is recorded in P0 evidence.
- Target validation is operationally sensitive; preserve target worktree and graph hashes and never copy target artifacts.
- If a child fails its invariant gate, stop progression and reopen that child/slice.
