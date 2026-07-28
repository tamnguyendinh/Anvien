# Anvien Graph Identity Resolution v2 Multi-Plan Authoring Plan

## Metadata

- Date: `2026-07-28`
- Status: `active / P2-A committed ce82a341 / P2-B committed a1c66865 / P2-C exact-copy Supervisor PASS / commit pending / P2-D-P2-G blocked / implementation unauthorized`
- Plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Evidence: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-evidence.md`
- Benchmark: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-benchmark.md`
- Actual status: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-actual-status.md`

## Goal

Convert the existing oversized `2026-07-26-anvien-graph-identity-resolution-v2` plan into a controlled multi-plan campaign without changing its technical intent. Keep three independent sibling roots under `docs/plans/`: the preserved legacy plan, this split-authoring plan, and a separate multi-plan campaign containing one roadmap plus seven complete implementation child plan sets. Preserve all 98 legacy implementation slices exactly once, give every child its own P0 and Pn-A/Pn-B/Pn-C lifecycle, and switch authority away from the legacy plan only after deterministic cross-plan checks and Supervisor acceptance pass.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
- Every non-terminal child plan must state in its own `Rules` that Pn-C cannot close or hand off until the next child plan's `actual-status.md` is refreshed from the latest accepted repo/runtime/evidence state, the refresh log and affected next actions/work steps are updated, and qualified evidence for that refresh is recorded. The terminal child must state that no successor exists and refresh the roadmap/campaign closure status instead.
- Run Anvien detect-changes before every implementation-slice commit when implementation work was performed.
- For public runtime or UI-facing changes, validate the real user-visible runtime with browser or Playwright evidence.
- For app/runtime validation, full build must include Docker image/container build. If Docker is missing or not run, full build is incomplete.
- Any Playwright validation must target the real built Docker/container runtime. Running Playwright against a host dev server, framework dev mode, mocked server, or source-run shortcut is not valid runtime evidence.
- If the Docker runtime cannot be built or started, the slice/plan is blocked; do not replace it with dev-server Playwright evidence.
- Playwright evidence must record the Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
- Keep the standard planner structure. These detail rules only make phase checklist items concrete enough to implement safely.
- Every implementation phase must be decomposed into multiple implementation slices that are as small as practical. A phase is a grouping and ordering container; a slice is the executable implementation unit.
- Do not implement a phase directly. Work starts from a slice ID such as `P1-A`, `P1-B`, or `P2-C`.
- Prefer many narrow slices over one broad slice. A single-slice implementation phase is allowed only when the plan explicitly states why the phase cannot be split further without creating empty or non-executable slices.
- Each implementation slice must include:
  + Goal
  + Scope Boundary
  + Non-Goals when useful
  + Pre-flight Questions
  + Work Steps
  + Implementation Gate
  + Acceptance
  + Evidence Targets
  + Actual-status Update
  + Commit Boundary
- Split planned work into separate slices when it contains more than one primary user-visible behavior, user trigger, render location, permission or visibility rule, DB write target, DB state transition, API/CLI/MCP contract, async/event/webhook flow, external side effect, cleanup/quarantine domain, behavior test target, independent acceptance gate, or independent commit boundary.
- Hidden fallback is forbidden. Prefer a visible failure over a fallback that hides a broken primary path.
- When touching DB-backed content, verify the full loop when applicable: UI input -> submit action -> DB write -> DB read after reload/new request -> correct UI render or omission. If there is no UI, replace UI steps with the real caller/consumer flow.
- Tests must prove product behavior. Delete or replace tests that only assert implementation details, helper output, static DOM existence, or mocked plumbing without proving trigger -> process -> observable result.
- If a planned item uses wording such as `and`, `also`, `then wire`, `plus update`, `both`, or `handle all`, check whether it is actually multiple slices.
- Do not write broad actionable items such as `Implement checkout, webhook, entitlement update, and billing UI`; split them into narrow slices such as `Create checkout session request`, `Persist checkout session state`, `Handle provider webhook`, `Update entitlement from webhook event`, and `Render billing status from entitlement`.
- Each slice work step must include UI flow, DB/data flow, render location, and evidence target checks. Use `N/A` with a reason when a check does not apply.
- If tests write DB rows, app state, files, queues, provider state, or other persistent data, the slice must define cleanup or quarantine before implementation.
- This is a documentation-only authoring plan. Its execution may edit only the campaign roadmap, the seven child plan sets, this authoring plan set, and the legacy plan's authority marker after acceptance; it must not edit production code, tests, runtime configuration, generated graph data, or target-repository content.
- `E:\cheapapp.org` is out of scope and must not be written, copied, moved, staged, or used as storage. No file from it may be copied into `E:\Anvien`.
- Do not create any plan, fixture, report, temporary directory, or other artifact at the Anvien repository root. The legacy plan remains at `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/`; this authoring plan belongs at `docs/plans/2026-07-28-00-multi-plan-authoring/`; the roadmap and all child plans belong at `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/`. These three plan roots are siblings and must never be nested inside one another. Debug-only temporary material belongs under `E:\Anvien\.tmp\`.
- The legacy plan and its ledgers remain the active authority until P3-B receives Supervisor PASS. Partial child-plan output is draft-only and cannot authorize implementation.
- The legacy implementation phases map one-to-one: P1 through P7 become child plans `01` through `07`. Legacy P8 is closure material and must be distributed into each child's Pn-A/Pn-B/Pn-C; it must not become child plan `08`.
- Every child is a complete standard four-file plan set. A roadmap entry or a link to the legacy plan cannot replace a child's P0, implementation phase, evidence ledger, benchmark ledger, actual-status ledger, or Pn closure.
- Every child preserves its source implementation phase name and every source slice ID. Mechanical copy/paste does not remap `P<source>-<suffix>` to a child-local prefix.
- Preserve every legacy slice's goal, scope boundary, non-goals, pre-flight questions, work-step sequence, implementation gate, acceptance fields, evidence targets, actual-status update, and commit boundary verbatim. Only standalone metadata, companion paths, P0/Pn structure, predecessor/successor links, and explicitly missing handoff structure may be added.
- All 98 legacy implementation slices must be copied exactly once with their original IDs and order. Missing, duplicate, merged, invented, normalized, or silently dropped slices block authority cutover.
- Each child ledger is independent. Cross-plan evidence references must include both the child slug and the exact evidence ID; a bare `E1` or an unqualified `E1-P1A-*` reference is invalid across child boundaries.
- `index-reader-matrix.md` has one primary responsibility and one owning child: child `02-versioned-persistence-and-v2-cutover`. Other children may link to it as inspect-only evidence but may not duplicate or independently mutate its contract.
- Every authored file must own one primary planning responsibility. A file may link to many modules, plans, or files when those links serve that one responsibility; catch-all or duplicated ledgers are forbidden.
- The roadmap coordinates order, authority, dependencies, status, and handoffs. It does not contain implementation slice bodies and does not replace any child plan.
- The original giant plan is never deleted. After and only after accepted multi-plan output exists, change its status to `superseded / reference-only` and add a direct pointer to the roadmap; preserve its historical body and companion ledgers.
- Before graph-based inspection during execution, run `anvien analyze E:\Anvien --force`. If `file-detail` cannot represent a docs plan because that file class is not indexed, record the limitation truthfully and use disk hash, Markdown structure, Git diff, and deterministic mapping checks instead of inventing a relationship count.
- Do not execute this authoring plan merely because it exists. Execution begins only after an explicit user instruction.

## Problem

The current plan is a single 5,467-line control document containing seven implementation phases, one closure phase, and 98 implementation slices. Its technical order is valid, but its size makes independent execution, P0 refresh, evidence ownership, Supervisor review, cleanup, commit boundaries, and handoff control difficult. Treating the sample campaign's number of child plans as a fixed formula would be incorrect; this campaign's natural boundary is its seven existing implementation phases.

The required transformation is not a summary or a fresh redesign. It is a lossless planning migration. Each existing phase already contains its own executable slice decomposition, so each phase becomes one complete child plan rather than being split again by arbitrary file count or combined with another phase. The campaign also needs a roadmap to coordinate dependencies without turning that roadmap into a substitute for child lifecycle sections.

Authority must remain unambiguous throughout the migration. If the legacy plan is marked superseded before all child plans are complete and verified, implementation could proceed from an incomplete plan set. If both the legacy and new plans remain active after verification, agents could update different ledgers or execute conflicting slice IDs. The conversion therefore requires a hash-bound source snapshot, a bijective slice crosswalk, independent child ledgers, and an atomic authority cutover after Supervisor PASS.

## Scope

In scope:

- Freeze the legacy plan at SHA-256 `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB` as the transformation input, or stop and re-baseline if an authorized later edit changes it before execution.
- Create one parent roadmap at `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`.
- Create exactly seven implementation child plan folders under `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/`:
  - `2026-07-28-01-graph-identity-contract-and-strict-construction`
  - `2026-07-28-02-versioned-persistence-and-v2-cutover`
  - `2026-07-28-03-typescript-binding-pattern-extraction`
  - `2026-07-28-04-typescript-export-semantics`
  - `2026-07-28-05-module-export-and-reexport-resolution`
  - `2026-07-28-06-ambient-external-resolution-and-diagnostics`
  - `2026-07-28-07-cross-surface-acceptance-and-target-validation`
- Create the four standard planner files in each child folder with matching date and slug.
- Keep the four-file authoring plan set in its own sibling root `docs/plans/2026-07-28-00-multi-plan-authoring/`; it is not a child of either the legacy plan or the resulting multi-plan.
- Map legacy P1-P7 to child `01`-`07`, respectively, and copy each complete source phase into its child without changing phase or slice IDs.
- Give each child a populated P0, its preserved source implementation phase containing the complete source-phase slice set, and tailored Pn-A/Pn-B/Pn-C closure.
- Split current status, evidence, benchmarks, dependencies, acceptance, and handoff ownership so each child can be executed and closed independently without relying on chat history.
- Validate roadmap links, companion links, slice traceability, evidence qualification, file responsibility, dependency order, and lifecycle completeness.
- After Supervisor acceptance only, mark the giant legacy plan `superseded / reference-only` and point it to the roadmap.

## Non-Goals

- No execution of graph-identity, persistence, TypeScript extraction, export, resolver, diagnostic, or target-validation implementation work.
- No production-code or test change.
- No build, Docker, runtime, browser, Playwright, or performance execution for this docs-only transformation.
- No re-audit or redesign of the technical decisions already present in the legacy plan.
- No reduction, consolidation, expansion, reprioritization, or semantic rewrite of the 98 implementation slices.
- No child plan `08` for legacy P8.
- No treating the sample campaign's ten children as a required number for this campaign.
- No deletion of the legacy plan or its ledgers.
- No copy of `E:\cheapapp.org`, no target source inspection beyond evidence already referenced by the legacy plan, and no target write.
- No move of `index-reader-matrix.md` unless a later explicit plan revision separately proves that relocation is required; this plan assigns ownership without duplicating the file.
- No commit of the plan being authored in the current request unless the user separately orders a commit.

## Requirements

### Campaign output contract

| Child | Source phase | Source slice count | Local implementation phase | Primary responsibility | Dependency |
|-------|--------------|-------------------:|----------------------------|------------------------|------------|
| `01-graph-identity-contract-and-strict-construction` | P1 | 11 | P1 | identity contract and strict graph construction | none beyond accepted roadmap/P0 |
| `02-versioned-persistence-and-v2-cutover` | P2 | 42 | P1 | versioned persistence, opaque consumers, atomic generation, and v2 cutover | child 01 |
| `03-typescript-binding-pattern-extraction` | P3 | 17 | P1 | recursive TypeScript binding-pattern extraction and projection | child 02 |
| `04-typescript-export-semantics` | P4 | 15 | P1 | first-class TypeScript export semantics and projection | child 03 |
| `05-module-export-and-reexport-resolution` | P5 | 4 | P1 | module export tables and barrel/re-export resolution | child 04 |
| `06-ambient-external-resolution-and-diagnostics` | P6 | 6 | P1 | ambient/external declaration universe and truthful diagnostics | child 05 |
| `07-cross-surface-acceptance-and-target-validation` | P7 | 3 | P1 | cross-surface acceptance, in-place target validation, and performance | children 01-06 |

### Exact source-slice inventory

- Child 01 owns 11 source slices: `P1-A`, `P1-B`, `P1-C0`, `P1-C0A`, `P1-C0B`, `P1-C`, `P1-D`, `P1-D1`, `P1-D2`, `P1-D3`, `P1-E`.
- Child 02 owns 42 source slices: `P2-A`, `P2-A1` through `P2-A15`, `P2-B`, `P2-B1` through `P2-B4`, `P2-C`, `P2-C1` through `P2-C6`, `P2-D`, `P2-D1`, `P2-D2`, `P2-E`, `P2-E1`, `P2-E2`, `P2-F`, `P2-F1` through `P2-F6`, `P2-G`.
- Child 03 owns 17 source slices: `P3-A`, `P3-B`, `P3-B1`, `P3-B2`, `P3-B2A`, `P3-C`, `P3-C1`, `P3-C1A` through `P3-C1I`, `P3-C2`.
- Child 04 owns 15 source slices: `P4-A`, `P4-B`, `P4-B1`, `P4-C`, `P4-C1`, `P4-C1A` through `P4-C1I`, `P4-C2`.
- Child 05 owns 4 source slices: `P5-A`, `P5-B`, `P5-C`, `P5-D`.
- Child 06 owns 6 source slices: `P6-A`, `P6-B`, `P6-C1`, `P6-C2`, `P6-C3`, `P6-D`.
- Child 07 owns 3 source slices: `P7-A`, `P7-B`, `P7-C`.

### Child completeness contract

- Every child folder contains `*-plan.md`, `*-evidence.md`, `*-benchmark.md`, and `*-actual-status.md` with the exact same child slug.
- Every child plan follows the current planner template and contains metadata, goal, rules, problem, scope, non-goals, requirements, acceptance criteria, checklist, and risk notes.
- Every child plan contains `P0-A`, its preserved source implementation phase with every source slice in original order, and Pn-A/Pn-B/Pn-C.
- Every mapped slice includes all required Scope Boundary fields, all 12 Pre-flight fields, ordered Work Steps with UI/data/render/Mini-QA/evidence checks, Implementation Gate, all seven Acceptance fields, Evidence Targets, Actual-status Update, and Commit Boundary.
- Every child actual-status ledger contains a real baseline classification, relationship evidence or a truthful docs/graph limitation, a touch map, a refresh log, downstream decisions, and a final P0 decision. Later execution must refresh P0 before opening that child if prior children changed repo reality.
- Every child evidence ledger uses child-local stable evidence IDs. Cross-child references are qualified by child slug and exact evidence ID.
- Every child benchmark ledger owns only measurements relevant to that child. Campaign totals belong in the roadmap or authoring benchmark, not duplicated as mutable child totals.
- Every child plan repeats the successor-freshness invariant in its own `Rules`; every non-terminal Pn-C updates the next child `actual-status.md` from the latest accepted evidence, appends its refresh-log row, updates affected next actions/work steps, and reserves a child-local proof such as `E2-PNC-NEXTSTATUS1`. A stale or missing successor actual-status update blocks closure. Child 07 records the no-successor terminal case and refreshes roadmap/campaign closure status instead.
- Every child Pn-C records its own validation, detect-changes, commit, worktree state, successor actual-status refresh (or terminal no-successor proof), and roadmap handoff. Finishing a child does not silently close the campaign.

### Authority and ownership contract

- The authoring plan set is control material for producing the campaign and is not counted among the seven implementation children.
- Directory ownership is exact: the legacy root owns the frozen source four-file set plus its existing `index-reader-matrix.md` auxiliary artifact; the authoring root owns only this standard four-file control plan; the multi-plan root owns the roadmap and numbered child plan folders. `index-reader-matrix.md` is not moved or duplicated by the structural correction; future mutation remains owned only by child 02.
- The roadmap is the sole campaign-order and active-authority index after cutover.
- A child plan is the sole execution authority for its mapped source phase after cutover.
- The legacy plan remains active until all seven children, all 28 standard child files, the roadmap, and the 98-row crosswalk pass checks and Supervisor review.
- The legacy plan's body and ledgers remain historical evidence. Only its status/pointer block may change during cutover.
- `index-reader-matrix.md` remains a single file and is owned for mutation by child 02; child 07 may validate it and other children may reference it inspect-only.
- Every output file has one primary responsibility even when it links to several modules or plan files.

## Acceptance Criteria

- Exactly one roadmap exists at the specified campaign-root path.
- Exactly three independent sibling plan roots exist at the specified paths; the legacy root contains no authoring-plan directory, roadmap, or numbered child-plan directory.
- Exactly seven implementation child folders exist with the specified slugs and order when authoring is complete; unstarted children remain absent and explicitly not authored.
- Exactly 28 standard child files exist: four per child, with matching filenames, metadata links, and H1 type.
- All authored children contain a completed P0 structure, the preserved source implementation phase, and Pn-A/Pn-B/Pn-C.
- The crosswalk contains exactly 98 source slices and 98 unchanged source-ID destinations, with zero missing, duplicate, merged, invented, normalized, or out-of-order mappings.
- Source-phase counts reproduce `11 + 42 + 17 + 15 + 4 + 6 + 3 = 98`.
- Every copied slice preserves its source text, semantics, IDs, order, and complete planner field structure; only the standalone child metadata/paths and explicitly missing handoff structure may be added.
- Legacy P8 produces closure sections in all seven children and does not produce an eighth implementation child.
- Child evidence, benchmark, and actual-status ledgers are independently usable and contain no unqualified cross-child evidence references.
- Every child plan contains the successor actual-status freshness rule in `Rules`, and every Pn-C reserves/accepts exact evidence proving the next child's actual-status refresh before handoff; child 07 records the terminal no-successor case and refreshes campaign closure status.
- `index-reader-matrix.md` has exactly one mutation owner, child 02.
- Roadmap and companion links resolve; no placeholder tokens, stale active-authority claims, or broken child dependencies remain.
- Supervisor passes the complete multi-plan set before authority cutover.
- After PASS, the legacy plan is preserved with only a `superseded / reference-only` status and roadmap pointer; before PASS, it remains active.
- Git diff proves no production source, test, runtime, graph output, repository-root artifact, or `E:\cheapapp.org` content was changed.
- This authoring plan's evidence, benchmark, actual-status, and checklist are current at closure.

## Checklist

- [x] P0-A: Complete actual status before multi-plan authoring work.
  - Goal: establish the real source-plan structure, authority state, repo basis, sample precedent, and output gap before any roadmap or child plan is authored.
  - Work Steps: refresh the Anvien graph, inspect the legacy plan and its companion files, hash and count the source plan, enumerate P1-P8 and all implementation slice IDs, inspect the proven multi-plan sample only for campaign structure, classify future roadmap/children as missing, record the docs `file-detail` limitation, and update P1-P3 from that evidence.
  - Implementation Gate: no roadmap or child plan may be created until `2026-07-28-00-multi-plan-authoring-actual-status.md` records a final P0 decision and exact evidence IDs.
  - Acceptance: actual status identifies the active legacy authority, seven source implementation phases, 98 source slices, legacy closure handling, missing roadmap/children, single-owner matrix rule, clean repo basis, and forbidden target boundary.

### P1: Freeze Transformation Contract And Campaign Skeleton

- Phase Goal: freeze a lossless source-to-child transformation contract and create the parent roadmap before child bodies are authored.
- Phase Boundary:
  - In scope: legacy source snapshot, exact phase/slice crosswalk contract, roadmap, child inventory, dependency order, authority rules, and file ownership.
  - Out of scope: child plan bodies, production code, tests, runtime artifacts, target-repository content, and legacy authority cutover.
  - Dependencies: completed P0 and an explicit user instruction to execute this authoring plan.
- Phase Implementation Rule: do not implement `P1` directly. Implement `P1-A`, verify it, record evidence, refresh actual-status, commit the docs-only slice when authorized, then continue to `P1-B`.
- Ordered Slice List:
  - P1-A: Freeze the legacy source snapshot and transformation contract.
  - P1-B: Author the campaign roadmap and exact child inventory.

- [x] P1-A: Freeze the legacy source snapshot and transformation contract.
  - Goal: establish a deterministic, lossless input contract so later authoring cannot silently omit, merge, invent, or reorder work.
  - Scope Boundary:
    - Editable: this authoring plan's evidence, benchmark, and actual-status ledgers.
    - Inspect-only: the five legacy campaign artifacts, planner templates, `AGENTS.md`, and the verified Restaurant Manager multi-plan sample.
    - Preserve-only: the legacy plan body, legacy ledgers, `index-reader-matrix.md`, production source, tests, and generated graph data.
    - Out of scope: roadmap creation, child plan creation, authority cutover, and all `E:\cheapapp.org` content.
  - Non-Goals: do not reconsider technical architecture, change source slice text, or infer extra child plans from the sample's child count.
  - Pre-flight Questions:
    - Data source: legacy plan at the recorded path and SHA-256, plus its four companion artifacts.
    - Display permission: N/A — no UI or user-visible runtime is changed.
    - DB read flow: N/A — no database is read.
    - DB write flow: N/A — no database is written.
    - Render location: Markdown ledgers in this authoring plan folder.
    - UI behavior flow: N/A — docs-only transformation control.
    - Docker runtime: N/A — no runtime behavior changes.
    - Playwright target: N/A — no browser-visible surface exists.
    - Behavior test: deterministic Markdown heading/slice inventory and hash checks.
    - Cleanup/quarantine: no temporary file is required; any debug artifact must remain under `E:\Anvien\.tmp\` and be removed before acceptance.
    - External side effects: Git worktree receives only authoring-ledger updates when execution is authorized.
    - N/A notes: UI, DB, Docker, and Playwright are inapplicable because this slice only freezes document inputs.
  - Work Steps:
    1. Refresh Anvien for `E:\Anvien`, record Git HEAD/status, verify the five source artifacts, and compare the legacy plan hash to the P0 baseline; if it differs, stop and re-baseline before continuing.
       - UI flow check: N/A — no UI flow.
       - DB/data flow check: verify file paths, bytes, line counts, SHA-256, phase headings, and source ledger presence.
       - Render location check: evidence and benchmark ledgers only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; run the nearest real boundary, the deterministic disk/Markdown inventory checks.
       - Evidence target: `E1-P1A-SNAPSHOT1`, `E1-P1A-GIT1`, and `B-P1-A`.
    2. Materialize the source-to-child transformation contract in the authoring evidence, including seven phase owners, all 98 unchanged source IDs, exact copy boundaries, P8-to-Pn distribution, and authority-cutover conditions.
       - UI flow check: N/A — no UI flow.
       - DB/data flow check: prove one source phase per child and one destination per source slice.
       - Render location check: authoring evidence and actual-status only; no roadmap yet.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; validate unique IDs and exact counts.
       - Evidence target: `E1-P1A-MAP1` and `E1-P1A-MAP2`.
  - Implementation Gate:
    - Before editing target docs, refresh Anvien and attempt the relevant file-detail command; record the truthful unsupported/not-indexed result instead of inventing relationship evidence.
    - Source hash, phase order, slice count, and authority state must be known; any unexplained source drift blocks P1-B.
  - Acceptance:
    - Source: legacy artifacts remain byte-identical and the frozen snapshot identifies all seven phases and 98 slices.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: the crosswalk contract is bijective by construction and has no target-repo data.
    - Behavior test: deterministic inventory reproduces per-phase counts `11/42/17/15/4/6/3` and total `98`.
    - Cleanup/quarantine: no plan-created temporary artifact remains outside the authoring folder.
    - Evidence IDs: `E1-P1A-SNAPSHOT1`, `E1-P1A-GIT1`, `E1-P1A-MAP1`, `E1-P1A-MAP2`.
    - Actual-status rows refreshed: legacy authority, source snapshot, phase/slice inventory, and roadmap/child readiness.
  - Evidence Targets: exact source hash, Git basis, ordered source IDs, child mapping, and legacy P8 closure mapping.
  - Actual-status Update: append a refresh row and change transformation contract from `partial` to `correct` only after all checks pass.
  - Commit Boundary: commit only this docs-only snapshot/contract slice after acceptance and Supervisor review when execution and commit are authorized.

- [x] P1-B: Author the campaign roadmap and exact child inventory.
  - Goal: create the single coordination document that fixes child order, ownership, dependencies, handoffs, status, and authority without duplicating implementation bodies.
  - Scope Boundary:
    - Editable: `2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md` and this authoring plan's ledgers.
    - Inspect-only: frozen P1-A contract, legacy plan headings, planner templates, and sample roadmap structure.
    - Preserve-only: legacy plan/ledgers, `index-reader-matrix.md`, child plan folders, production code, tests, and graph artifacts.
    - Out of scope: authoring any child plan body or marking the legacy plan superseded.
  - Non-Goals: the roadmap must not absorb implementation slices, shared mutable evidence, or a substitute P0/Pn lifecycle.
  - Pre-flight Questions:
    - Data source: accepted P1-A transformation contract and exact child inventory in Requirements.
    - Display permission: N/A — Markdown coordination only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: campaign-root roadmap.
    - UI behavior flow: N/A — no UI behavior.
    - Docker runtime: N/A — docs-only.
    - Playwright target: N/A — docs-only.
    - Behavior test: link, slug, ordering, and dependency validation.
    - Cleanup/quarantine: remove any abandoned roadmap draft created by this slice.
    - External side effects: one new Markdown roadmap and ledger updates inside `E:\Anvien`.
    - N/A notes: all runtime validation fields are inapplicable because no executable artifact changes.
  - Work Steps:
    1. Create the roadmap with campaign purpose, active-authority rules, seven-child table, dependencies, child status, slice counts, matrix owner, P0 refresh handoff, and Pn closure expectations.
       - UI flow check: N/A — no UI flow.
       - DB/data flow check: roadmap rows derive only from the frozen crosswalk.
       - Render location check: one campaign-root roadmap whose sole responsibility is coordination.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; inspect rendered Markdown structure and links.
       - Evidence target: `E1-P1B-ROADMAP1`.
    2. Validate that the roadmap names exactly seven child folders, 28 future standard files, one implementation phase per child, no child 08, and no active-authority ambiguity; update authoring ledgers and roadmap status.
       - UI flow check: N/A — no UI flow.
       - DB/data flow check: compare roadmap inventory to P1-A crosswalk and target counts.
       - Render location check: roadmap plus authoring evidence/benchmark/actual-status.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; run deterministic link/slug/count checks.
       - Evidence target: `E1-P1B-INVENTORY1` and `E1-P1B-LINK1`.
  - Implementation Gate:
    - Before editing target docs, refresh Anvien and record the docs file-detail limitation where applicable.
    - P1-A must pass; the roadmap path and seven slugs must exactly match Scope.
  - Acceptance:
    - Source: roadmap exists and legacy artifacts remain unchanged.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: roadmap has exactly seven unique child records and correct dependency edges.
    - Behavior test: links/slugs/counts pass and roadmap contains no implementation slice body.
    - Cleanup/quarantine: no duplicate roadmap or abandoned draft remains.
    - Evidence IDs: `E1-P1B-ROADMAP1`, `E1-P1B-INVENTORY1`, `E1-P1B-LINK1`.
    - Actual-status rows refreshed: roadmap, child inventory, campaign authority, and dependency contract.
  - Evidence Targets: roadmap path/hash, seven-row child inventory, dependency order, link results, and file responsibility review.
  - Actual-status Update: mark roadmap `missing -> correct`; keep seven child sets `missing` until P2.
  - Commit Boundary: commit the roadmap slice after acceptance and Supervisor review when execution and commit are authorized.

### P2: Author Seven Complete Child Plan Sets

- Phase Goal: author one complete, independent standard plan set for each legacy implementation phase while preserving every source slice and lifecycle contract.
- Phase Boundary:
  - In scope: seven named child folders, 28 standard files, phase-specific P0/evidence/benchmark/Pn content, exact source-ID copy/paste, and roadmap status/handoffs.
  - Out of scope: implementing any child, modifying production/tests/runtime, changing technical scope, or switching legacy authority.
  - Dependencies: P1-A and P1-B accepted; children are authored in numeric/source-phase order.
- Phase Implementation Rule: do not implement `P2` directly. Complete P2-A, the user-required structural correction P2-A1, then P2-B through P2-G in order. After each slice, verify it, record evidence, refresh actual-status and roadmap, and commit the docs-only scope when authorized before opening the next slice.
- Ordered Slice List:
  - P2-A: Author child 01 for legacy P1.
  - P2-A1: Rebase authored artifacts into three independent sibling plan roots.
  - P2-B: Author child 02 for legacy P2.
  - P2-C: Author child 03 for legacy P3.
  - P2-D: Author child 04 for legacy P4.
  - P2-E: Author child 05 for legacy P5.
  - P2-F: Author child 06 for legacy P6.
  - P2-G: Author child 07 for legacy P7.

- [x] P2-A: Author child 01 for legacy P1 by exact source-block copy/paste.
  - Goal: create the complete graph-identity/strict-construction child plan set with all 11 legacy P1 slices.
  - Scope Boundary:
    - Editable: the four files under `2026-07-28-01-graph-identity-contract-and-strict-construction/`, roadmap status, and authoring ledgers.
    - Inspect-only: legacy P1 text, source ledgers, frozen crosswalk, planner templates, and roadmap.
    - Preserve-only: legacy artifacts, other child folders, matrix, production code, tests, and target.
    - Out of scope: legacy P2-P8 implementation content except dependency/closure links.
  - Non-Goals: do not implement identity behavior or combine any of the 11 source slices.
  - Pre-flight Questions:
    - Data source: legacy P1 lines/blocks and P1-scoped ledger facts.
    - Display permission: N/A — plan docs only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: child 01 standard four-file set.
    - UI behavior flow: N/A — no UI behavior.
    - Docker runtime: N/A — no runtime change.
    - Playwright target: N/A — no runtime change.
    - Behavior test: structural and semantic diff of 11 mapped slices.
    - Cleanup/quarantine: delete only abandoned child-01 drafts created by this slice.
    - External side effects: four child Markdown files plus roadmap/authoring ledger updates.
    - N/A notes: runtime fields are inapplicable to docs authoring; future child content retains its own applicable validation rules.
  - Work Steps:
    1. Create child 01's plan, evidence, benchmark, and actual-status files from current templates; copy the complete source P1 blocks verbatim and add P0/Pn-A/Pn-B/Pn-C.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: copy only P1-scoped facts and ownership; do not alter source IDs.
       - Render location check: exactly four files in the child-01 folder.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; validate Markdown structure.
       - Evidence target: `E2-P2A-FILES1` and `E2-P2A-STRUCT1`.
    2. Copy `P1-A`, `P1-B`, `P1-C0`, `P1-C0A`, `P1-C0B`, `P1-C`, `P1-D`, `P1-D1`, `P1-D2`, `P1-D3`, and `P1-E` in order without changing IDs, then add only missing standalone metadata/handoff fields and update roadmap/ledgers.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: 11 source IDs copy to the same 11 source IDs in the child.
       - Render location check: child plan owns slice bodies; roadmap owns status only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; compare source/destination fields and ordering.
       - Evidence target: `E2-P2A-MAP1` and `B-P2-A`.
  - Implementation Gate:
    - Record Anvien freshness and docs relationship limitation before edits.
    - P1-B must pass; source P1 hash/block and the 11-ID inventory must match the frozen contract.
  - Acceptance:
    - Source: four standard files exist and the legacy plan remains unchanged.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: 11/11 source slices map once in original order.
    - Behavior test: child structure, required fields, links, P0, P1, and Pn all pass.
    - Cleanup/quarantine: no duplicate or partial child-01 file remains.
    - Evidence IDs: `E2-P2A-FILES1`, `E2-P2A-STRUCT1`, `E2-P2A-MAP1`, `E2-P2A-FD1`, `E2-P2A-SUP1`, `E2-P2A-REBUILD1`, `E2-P2A-SUP2`.
    - Actual-status rows refreshed: child 01, roadmap status, mapped-slice total, and next-child dependency.
  - Evidence Targets: four-file inventory, 11-row mapping, field completeness, local evidence ownership, and link validation.
  - Actual-status Update: mark child 01 `missing -> correct`; keep child 02 blocked on accepted child-01 authoring handoff.
  - Commit Boundary: commit only child 01 plus its roadmap/authoring ledger updates after acceptance and Supervisor review when authorized.

- [x] P2-A1: Rebase authored artifacts into three independent sibling plan roots.
  - Goal: correct the campaign directory ownership so the legacy plan, split-authoring plan, and resulting multi-plan are three independent sibling plans before child 02 authoring continues.
  - Scope Boundary:
    - Editable: this authoring four-file set and its path references; the roadmap; child-01 four-file set; path-only references in the unaccepted child-02 draft; Git move metadata for those artifacts.
    - Inspect-only: the legacy four-file source set, `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/index-reader-matrix.md`, existing Supervisor reports, Git history, and planner templates.
    - Preserve-only: legacy plan/ledger contents and hashes, matrix content/path, child slice semantics/evidence IDs, production source, tests, runtime, graph outputs, and `E:\cheapapp.org`.
    - Out of scope: accepting child 02, authoring children 03-07, changing implementation contracts, moving/duplicating the reader matrix, or switching authority.
  - Non-Goals: do not solve the hierarchy by nesting one plan inside another, copying files, leaving compatibility duplicates, or treating a filesystem move as implementation acceptance.
  - Pre-flight Questions:
    - Data source: the user's three-root directory contract, commit `b760d156`, current disk inventory, and frozen legacy hash.
    - Display permission: N/A — documentation structure only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: three sibling roots under `docs/plans/` and their Markdown links.
    - UI behavior flow: N/A — no UI.
    - Docker runtime: N/A — no executable change.
    - Playwright target: N/A — no runtime surface.
    - Behavior test: exact path inventory, Git rename/copy detection, companion-link resolution, relative roadmap links, legacy-root contamination count, hashes, placeholders, and scoped diff.
    - Cleanup/quarantine: remove only superseded nested directories after every file is proven present at its single destination; create no temporary plan copy.
    - External side effects: docs-only filesystem moves and path-reference edits inside `E:\Anvien`.
    - N/A notes: disk/Git/Markdown validation is the nearest real boundary.
  - Work Steps:
    1. Freeze the pre-move inventory/hashes and rebase the committed authoring set to `docs/plans/2026-07-28-00-multi-plan-authoring/`; create `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/`; move the roadmap and child 01 into it; relocate the unaccepted child-02 draft there without marking P2-B complete.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: each source file has exactly one destination; no copy/duplicate remains.
       - Render location check: legacy, authoring, and multi-plan roots are siblings directly under `docs/plans/`.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; inspect the exact disk tree and Git rename state.
       - Evidence target: `E2-P2A1-USER1`, `E2-P2A1-INVENTORY1`, `E2-P2A1-MOVE1`, `E2-P2A1-GRAPH1`, `E2-P2A1-FD1`.
    2. Update companion/source/roadmap links without changing slice semantics, run deterministic structure/link/hash/Git-scope checks, call Supervisor, and keep P2-B paused until this correction is accepted and committed.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: legacy root retains only its source four-file set plus matrix; authoring root has exactly four files; multi-plan root has one roadmap and the authored/draft child folders at their single locations.
       - Render location check: every companion and roadmap link resolves from the corrected roots.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; render/read Markdown and validate links from disk.
       - Evidence target: `E2-P2A1-STRUCT1`, `E2-P2A1-LINK1`, `E2-P2A1-DIFF1`, `E2-P2A1-SUP1`.
  - Implementation Gate:
    - User correction explicitly requires three sibling plan roots; P2-A is committed at `b760d156` and P2-B is not accepted.
    - Legacy source hash must remain `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB`; the matrix must remain byte-identical at its legacy-root path.
  - Acceptance:
    - Source: three sibling roots exist exactly once; the legacy four-file body and matrix are unchanged.
    - Runtime/UI: N/A — no executable/runtime/UI change.
    - DB/data: authoring root owns four standard files; multi-plan root owns one roadmap plus child folders; legacy root owns no nested authoring/multi-plan output.
    - Behavior test: all companion/roadmap/source links resolve, Git shows moves rather than duplicates where applicable, and no stale old path remains in active plan artifacts.
    - Cleanup/quarantine: no copied duplicate, empty obsolete nested directory, repository-root artifact, or target artifact remains.
    - Evidence IDs: `E2-P2A1-USER1`, `E2-P2A1-INVENTORY1`, `E2-P2A1-MOVE1`, `E2-P2A1-GRAPH1`, `E2-P2A1-FD1`, `E2-P2A1-STRUCT1`, `E2-P2A1-LINK1`, `E2-P2A1-DIFF1`, `E2-P2A1-SUP1`.
    - Actual-status rows refreshed: three-root ownership, roadmap location, child-01 location, child-02 draft location, and P2-B readiness.
  - Evidence Targets: before/after tree, hashes, fresh Anvien/file-detail state, Git rename state, zero nested plan roots, resolved links, scoped diff, and unconditional Supervisor PASS.
  - Actual-status Update: directory ownership `wrong -> correct`; P2-B `paused -> ready` only after acceptance/commit.
  - Commit Boundary: commit this structural correction separately; exclude the unaccepted child-02 four-file draft content from the commit, then resume P2-B.

- [x] P2-B: Author child 02 for legacy P2.
  - Current Status: exact-copy candidate created from committed Child 01 and the corrected source-ID-preserving crosswalk; deterministic validation/red-team/Supervisor PASS; committed at `a1c66865`. Historical candidate evidence does not accept this replacement candidate.
  - Goal: create the complete persistence/cutover child plan set with all 42 legacy P2 slices and sole mutation ownership of `index-reader-matrix.md`.
  - Scope Boundary:
    - Editable: the four child-02 files under the separate multi-plan root, roadmap status, authoring ledgers, and matrix ownership metadata/reference only; matrix content remains unchanged during authoring.
    - Inspect-only: legacy P2, source ledgers, `index-reader-matrix.md`, child-01 handoff, templates, and roadmap.
    - Preserve-only: legacy bodies; all accepted child-01 content; children 03-07; production/tests/runtime; graph outputs; and target.
    - Out of scope: executing persistence/cutover or duplicating the reader matrix.
  - Non-Goals: do not collapse or remap the 42 slices, give another child mutation ownership of the matrix, or reopen child-01 technical semantics.
  - Pre-flight Questions:
    - Data source: legacy P2 blocks, P2 ledger facts, reader matrix, and child-01 dependency contract.
    - Display permission: N/A — docs-only.
    - DB read flow: N/A — no live DB; future storage flows remain plan content.
    - DB write flow: N/A — no live DB.
    - Render location: child 02 standard four-file set.
    - UI behavior flow: N/A — no UI behavior.
    - Docker runtime: N/A — no runtime change.
    - Playwright target: N/A — no runtime change.
    - Behavior test: 42-slice structural/semantic mapping and single matrix owner check.
    - Cleanup/quarantine: remove abandoned child-02 drafts only.
    - External side effects: four Markdown files plus roadmap/authoring ledger updates.
    - N/A notes: runtime checks are deferred to future execution of child 02.
  - Work Steps:
    1. Verify the user-required successor-freshness rule remains exactly once in child 01 and add the same one-line operational rule to child 02 without rewriting copied source content.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: a non-terminal child cannot close until the next child actual-status, refresh log, affected next actions/work steps, and qualified proof are current.
       - Render location check: authoring contract, roadmap, child-01 plan, and child-02 plan only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; validate the exact one-line rule and confirm no child-01 technical text changed.
       - Evidence target: `E2-P2B-USER2` and `E2-P2B-HANDOFFRULE1`.
    2. Create and populate all four child-02 files with P0, the unchanged source P2 phase containing all 42 source slices, independent ledgers, and tailored Pn closure/handoff.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: assign persistence/cutover and matrix ownership only to child 02.
       - Render location check: exactly four files in child 02; the matrix remains single at `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/index-reader-matrix.md` in the preserved legacy root.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; run template/field checks.
       - Evidence target: `E2-P2B-FILES1` and `E2-P2B-STRUCT1`.
    3. Copy the complete P2 inventory in original order without changing IDs, qualify child-01 dependency and later handoffs, prove 42/42 exact source-block matches, and update roadmap/authoring ledgers.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: verify group counts A=16, B=5, C=7, D=3, E=3, F=7, G=1, total 42.
       - Render location check: child plan owns slice bodies; roadmap owns coordination; matrix has one owner.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; compare source/destination fields and matrix references.
       - Evidence target: `E2-P2B-MAP1`, `E2-P2B-MATRIX1`, and `B-P2-B`.
  - Implementation Gate:
    - Record Anvien freshness and docs relationship limitation before edits.
    - Rebuilt P2-A must be committed; the authoring crosswalk must be corrected to preserve source IDs; the full 42-ID P2 inventory and exact matrix path must match the frozen contract.
  - Acceptance:
    - Source: four standard child files exist; legacy and matrix contents remain unchanged.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: 42/42 source slices copy once with unchanged IDs; matrix has exactly one mutation owner.
    - Behavior test: child completeness, exact source-block equality, source order, qualified dependency, link checks, and the one-line successor rule pass; child-01 content remains unchanged.
    - Cleanup/quarantine: no duplicate matrix or partial child-02 artifact exists.
    - Evidence IDs: `E2-P2B-RESET1`, `E2-P2B-RFILES1`, `E2-P2B-RSTRUCT1`, `E2-P2B-RMAP1`, `E2-P2B-RVALID1`, `E2-P2B-RGRAPH1`, `E2-P2B-RFD1`, `E2-P2B-RREDTEAM1`, `E2-P2B-RSUP1`; historical `E2-P2B-*` candidate evidence is not acceptance proof for this rebuilt child.
    - Actual-status rows refreshed: child 02, matrix ownership, cumulative mapped count, and child-03 dependency.
  - Evidence Targets: successor-freshness rule/gate/evidence proof for authored children, four-file inventory, 42-row mapping, group totals, matrix single-owner proof, and field completeness.
  - Actual-status Update: mark the cross-child freshness contract `partial -> correct`; mark child 02 `missing -> correct`; set child 03 next action to receive a latest-evidence actual-status refresh before consuming child-02 cutover handoff without duplicating its contracts.
  - Commit Boundary: commit child 02, the exact child-01 successor-freshness plan/evidence correction, and roadmap/authoring ledger updates only after acceptance and Supervisor review.

- [ ] P2-C: Author child 03 for legacy P3.
  - Current Status: exact-copy four-file candidate created after Child 02 commit `a1c66865`; deterministic source-block, graph, red-team, and Supervisor review PASS; commit pending and implementation remains unauthorized.
  - Goal: create the complete TypeScript binding-pattern child plan set with all 17 legacy P3 slices.
  - Scope Boundary:
    - Editable: child-03 four-file set, roadmap status, and authoring ledgers.
    - Inspect-only: legacy P3, P3 ledger facts, child-02 handoff, templates, and roadmap.
    - Preserve-only: all other children, legacy artifacts, matrix, code/tests/runtime, and target.
    - Out of scope: export semantics and execution of binding extraction.
  - Non-Goals: do not merge binding contexts or projection-adapter slices.
  - Pre-flight Questions:
    - Data source: legacy P3 and relevant source ledgers.
    - Display permission: N/A — docs-only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: child 03 standard four-file set.
    - UI behavior flow: N/A — no UI.
    - Docker runtime: N/A — no runtime.
    - Playwright target: N/A — no runtime.
    - Behavior test: exact 17-slice mapping and full-field checks.
    - Cleanup/quarantine: remove only abandoned child-03 drafts.
    - External side effects: four docs plus roadmap/ledger updates.
    - N/A notes: future CLI/MCP/HTTP/Web validation remains inside mapped child content, not this authoring slice.
  - Work Steps:
    1. Create the four child-03 files and populate P0, the unchanged source P3 phase, independent ledgers, and Pn closure from P3-scoped source material.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: no source facts outside P3 become owned here.
       - Render location check: child-03 folder only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; run structural checks.
       - Evidence target: `E2-P2C-FILES1` and `E2-P2C-STRUCT1`.
    2. Map P3's 17 slices in order, including parameter/catch/loop contexts and all projection adapters, then update roadmap and ledgers.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: 17 source IDs copy to the same 17 source IDs.
       - Render location check: slice bodies stay in child 03; roadmap stores only status/handoff.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; compare titles, fields, order, and references.
       - Evidence target: `E2-P2C-MAP1` and `B-P2-C`.
  - Implementation Gate:
    - Record Anvien freshness/docs limitation and require accepted child-02 authoring handoff.
    - Source P3 inventory must match the frozen 17-ID contract.
  - Acceptance:
    - Source: four standard files exist and source docs remain preserved.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: 17/17 slices map exactly once.
    - Behavior test: child lifecycle, complete fields, source order, and links pass.
    - Cleanup/quarantine: no partial child-03 artifact remains.
    - Evidence IDs: `E2-P2C-FILES1`, `E2-P2C-STRUCT1`, `E2-P2C-MAP1`, `E2-P2C-VALID1`, `E2-P2C-GRAPH1`, `E2-P2C-FD1`, `E2-P2C-REDTEAM1`, `E2-P2C-SUP1`.
    - Actual-status rows refreshed: child 03, cumulative mapped count, and child-04 dependency.
  - Evidence Targets: four-file inventory, 17-row mapping, context/adapter coverage, and qualified evidence links.
  - Actual-status Update: mark child 03 `missing -> correct`; keep child 04 pending child-03 handoff.
  - Commit Boundary: commit only child 03 plus roadmap/authoring ledger updates after acceptance and Supervisor review when authorized.

- [ ] P2-D: Author child 04 for legacy P4.
  - Goal: create the complete TypeScript export-semantics child plan set with all 15 legacy P4 slices.
  - Scope Boundary:
    - Editable: child-04 four-file set, roadmap status, and authoring ledgers.
    - Inspect-only: legacy P4, relevant ledgers, child-03 handoff, templates, and roadmap.
    - Preserve-only: other children, legacy artifacts, matrix, production/tests/runtime, and target.
    - Out of scope: module export-table traversal and runtime implementation.
  - Non-Goals: do not combine extraction and projection adapters or replace export facts with visibility metadata.
  - Pre-flight Questions:
    - Data source: legacy P4 and P4-scoped ledger facts.
    - Display permission: N/A — docs-only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: child 04 standard four-file set.
    - UI behavior flow: N/A — no UI.
    - Docker runtime: N/A — no runtime.
    - Playwright target: N/A — no runtime.
    - Behavior test: exact 15-slice mapping and field completeness.
    - Cleanup/quarantine: remove abandoned child-04 drafts only.
    - External side effects: four docs plus roadmap/ledger updates.
    - N/A notes: future projection/runtime checks remain requirements inside the child plan.
  - Work Steps:
    1. Create the four child-04 files with P0, the unchanged source P4 phase, phase-owned ledgers, and tailored Pn closure.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: retain direct/default/alias/type-only/star/namespace/re-export distinctions.
       - Render location check: child-04 folder only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; validate structural fields.
       - Evidence target: `E2-P2D-FILES1` and `E2-P2D-STRUCT1`.
    2. Map all 15 P4 slices in order, preserve export semantic distinctions and adapter boundaries, then update roadmap and ledgers.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: 15 source IDs copy to the same 15 source IDs.
       - Render location check: child plan owns implementation detail; roadmap owns handoff only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; compare source/destination semantics and links.
       - Evidence target: `E2-P2D-MAP1` and `B-P2-D`.
  - Implementation Gate:
    - Record Anvien freshness/docs limitation and require accepted child-03 authoring handoff.
    - Source P4 inventory must match the frozen 15-ID contract.
  - Acceptance:
    - Source: four files exist and legacy source remains preserved.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: 15/15 slices map once without semantic collapse.
    - Behavior test: child lifecycle, field completeness, source order, and links pass.
    - Cleanup/quarantine: no partial child-04 artifact remains.
    - Evidence IDs: `E2-P2D-FILES1`, `E2-P2D-STRUCT1`, `E2-P2D-MAP1`.
    - Actual-status rows refreshed: child 04, cumulative mapped count, and child-05 dependency.
  - Evidence Targets: four-file inventory, 15-row mapping, semantic-lane preservation, and independent ledger checks.
  - Actual-status Update: mark child 04 `missing -> correct`; set child 05 to consume the export-semantics handoff.
  - Commit Boundary: commit only child 04 plus roadmap/authoring ledger updates after acceptance and Supervisor review when authorized.

- [ ] P2-E: Author child 05 for legacy P5.
  - Goal: create the complete module export-table and barrel/re-export resolution child plan set with all four legacy P5 slices.
  - Scope Boundary:
    - Editable: child-05 four-file set, roadmap status, and authoring ledgers.
    - Inspect-only: legacy P5 including its semantic vector manifest, P5 ledgers, child-04 handoff, templates, and roadmap.
    - Preserve-only: other children, source artifacts, matrix, production/tests/runtime, and target.
    - Out of scope: ambient declaration resolution and actual resolver implementation.
  - Non-Goals: do not split the authoritative P5 semantic vector away from its owning child or merge its four executable slices.
  - Pre-flight Questions:
    - Data source: legacy P5, its semantic vector manifest, and P5-scoped evidence/metrics.
    - Display permission: N/A — docs-only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: child 05 standard four-file set.
    - UI behavior flow: N/A — no UI.
    - Docker runtime: N/A — no runtime.
    - Playwright target: N/A — no runtime.
    - Behavior test: four-slice mapping plus semantic-vector ownership check.
    - Cleanup/quarantine: remove abandoned child-05 drafts only.
    - External side effects: four docs plus roadmap/ledger updates.
    - N/A notes: future resolver validation remains inside child 05.
  - Work Steps:
    1. Create all four child-05 files and place the P5 semantic vector contract with the plan that owns its resolution behavior.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: preserve vector/status/proof semantics and source provenance.
       - Render location check: child-05 plan/ledgers, not roadmap or another child.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; validate structure and ownership.
       - Evidence target: `E2-P2E-FILES1`, `E2-P2E-STRUCT1`, and `E2-P2E-VECTOR1`.
    2. Map `P5-A` through `P5-D` in order, qualify child-04/child-06 handoffs, and update roadmap/authoring ledgers.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: four source IDs copy to the same four source IDs.
       - Render location check: implementation detail stays in child 05.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; compare source/destination content and references.
       - Evidence target: `E2-P2E-MAP1` and `B-P2-E`.
  - Implementation Gate:
    - Record Anvien freshness/docs limitation and require accepted child-04 handoff.
    - Source P5 inventory and semantic vector must match the frozen source.
  - Acceptance:
    - Source: four files exist and legacy P5 remains preserved.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: 4/4 slices map exactly once and vector ownership is singular.
    - Behavior test: lifecycle, fields, order, vector content, and links pass.
    - Cleanup/quarantine: no duplicate semantic-vector contract or partial child file remains.
    - Evidence IDs: `E2-P2E-FILES1`, `E2-P2E-STRUCT1`, `E2-P2E-VECTOR1`, `E2-P2E-MAP1`.
    - Actual-status rows refreshed: child 05, cumulative mapped count, and child-06 dependency.
  - Evidence Targets: four-file inventory, four-row mapping, semantic-vector provenance, and qualified handoffs.
  - Actual-status Update: mark child 05 `missing -> correct`; set child 06 to consume its resolver-table handoff.
  - Commit Boundary: commit only child 05 plus roadmap/authoring ledger updates after acceptance and Supervisor review when authorized.

- [ ] P2-F: Author child 06 for legacy P6.
  - Goal: create the complete ambient/external declaration and truthful-diagnostics child plan set with all six legacy P6 slices.
  - Scope Boundary:
    - Editable: child-06 four-file set, roadmap status, and authoring ledgers.
    - Inspect-only: legacy P6 including its status matrix, P6 ledgers, child-05 handoff, templates, and roadmap.
    - Preserve-only: other children, source artifacts, reader matrix, production/tests/runtime, and target.
    - Out of scope: cross-surface campaign acceptance and resolver implementation.
  - Non-Goals: do not collapse candidate discovery, authorization/materialization, final outcome, and diagnostic projection.
  - Pre-flight Questions:
    - Data source: legacy P6, its authoritative status matrix, and P6-scoped evidence/metrics.
    - Display permission: N/A — docs-only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: child 06 standard four-file set.
    - UI behavior flow: N/A — no UI.
    - Docker runtime: N/A — no runtime.
    - Playwright target: N/A — no runtime.
    - Behavior test: six-slice mapping plus status-matrix ownership/coverage check.
    - Cleanup/quarantine: remove abandoned child-06 drafts only.
    - External side effects: four docs plus roadmap/ledger updates.
    - N/A notes: future resolver and diagnostic boundary tests remain within child 06.
  - Work Steps:
    1. Create the four child-06 files and place the P6 outcome/status contract with the child that owns declaration-universe and diagnostic projection behavior.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: preserve candidate, authorization, immutable-outcome, and diagnostic-projection boundaries.
       - Render location check: child-06 plan/ledgers only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; validate structure and ownership.
       - Evidence target: `E2-P2F-FILES1`, `E2-P2F-STRUCT1`, and `E2-P2F-STATUS1`.
    2. Map `P6-A`, `P6-B`, `P6-C1`, `P6-C2`, `P6-C3`, and `P6-D` in order, qualify handoffs, and update roadmap/authoring ledgers.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: six source IDs copy to the same six source IDs.
       - Render location check: detailed resolution outcomes remain in child 06.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; compare source/destination fields and status rows.
       - Evidence target: `E2-P2F-MAP1` and `B-P2-F`.
  - Implementation Gate:
    - Record Anvien freshness/docs limitation and require accepted child-05 handoff.
    - Source P6 inventory and status matrix must match the frozen source.
  - Acceptance:
    - Source: four child files exist and legacy P6 remains preserved.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: 6/6 slices map once and status semantics remain exhaustive.
    - Behavior test: lifecycle, fields, order, status contract, and links pass.
    - Cleanup/quarantine: no duplicate status contract or partial child-06 artifact remains.
    - Evidence IDs: `E2-P2F-FILES1`, `E2-P2F-STRUCT1`, `E2-P2F-STATUS1`, `E2-P2F-MAP1`.
    - Actual-status rows refreshed: child 06, cumulative mapped count, and child-07 dependency.
  - Evidence Targets: four-file inventory, six-row mapping, outcome/status preservation, and qualified handoffs.
  - Actual-status Update: mark child 06 `missing -> correct`; allow child 07 authoring to consume all implementation-child acceptance contracts.
  - Commit Boundary: commit only child 06 plus roadmap/authoring ledger updates after acceptance and Supervisor review when authorized.

- [ ] P2-G: Author child 07 for legacy P7.
  - Goal: create the complete cross-surface acceptance/target-validation child plan set with all three legacy P7 slices and explicit dependencies on children 01-06.
  - Scope Boundary:
    - Editable: child-07 four-file set, roadmap status, and authoring ledgers.
    - Inspect-only: legacy P7, P7 ledgers, acceptance contracts from children 01-06, matrix as child-02-owned evidence, templates, and roadmap.
    - Preserve-only: other children, source artifacts, matrix content, production/tests/runtime, graph outputs, and target.
    - Out of scope: running acceptance, analyzing/writing `E:\cheapapp.org`, or taking ownership of earlier implementation contracts.
  - Non-Goals: do not treat target validation as permission to alter the target or omit full campaign dependencies.
  - Pre-flight Questions:
    - Data source: legacy P7 plus qualified acceptance handoffs from children 01-06.
    - Display permission: N/A — docs-only authoring.
    - DB read flow: N/A — no live DB.
    - DB write flow: N/A — no live DB.
    - Render location: child 07 standard four-file set.
    - UI behavior flow: N/A — no UI is executed during authoring.
    - Docker runtime: N/A for authoring; future child retains exact runtime requirements from source P7.
    - Playwright target: N/A for authoring; future child records applicable real-boundary validation.
    - Behavior test: three-slice mapping, dependency completeness, and target-boundary checks.
    - Cleanup/quarantine: remove abandoned child-07 drafts only; never use target as scratch space.
    - External side effects: four docs plus roadmap/ledger updates inside Anvien only.
    - N/A notes: target/runtime validation is planned content, not an action authorized by this authoring slice.
  - Work Steps:
    1. Create the four child-07 files with P0, the unchanged source P7 phase, independent ledgers, Pn closure, all six upstream handoffs, and the unchanged in-place target boundary.
       - UI flow check: N/A — docs-only.
       - DB/data flow check: acceptance consumes qualified upstream evidence without duplicating ownership.
       - Render location check: child-07 folder; roadmap stores campaign status only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; validate structure, dependencies, and target prohibitions.
       - Evidence target: `E2-P2G-FILES1`, `E2-P2G-STRUCT1`, and `E2-P2G-DEPENDENCY1`.
    2. Map `P7-A`, `P7-B`, and `P7-C` in order, preserve exact validation/performance boundaries, distribute legacy P8 into tailored Pn, then update roadmap and ledgers.
       - UI flow check: N/A — authoring only; future runtime instructions remain intact.
       - DB/data flow check: three source IDs copy to the same three source IDs and consume six qualified handoffs.
       - Render location check: implementation/acceptance detail stays in child 07.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; compare source/destination semantics and lifecycle.
       - Evidence target: `E2-P2G-MAP1`, `E2-P2G-CLOSURE1`, and `B-P2-G`.
  - Implementation Gate:
    - Record Anvien freshness/docs limitation and require accepted authoring handoffs from children 01-06.
    - Source P7/P8 content and the three-ID inventory must match the frozen contract.
  - Acceptance:
    - Source: four child files exist; legacy plan and target remain unchanged.
    - Runtime/UI: N/A — no runtime/UI execution or claim.
    - DB/data: 3/3 source slices map once; all six upstream dependencies are qualified.
    - Behavior test: lifecycle, fields, order, target boundary, Pn distribution, and links pass.
    - Cleanup/quarantine: no target artifact, duplicate closure child, or partial child-07 file remains.
    - Evidence IDs: `E2-P2G-FILES1`, `E2-P2G-STRUCT1`, `E2-P2G-DEPENDENCY1`, `E2-P2G-MAP1`, `E2-P2G-CLOSURE1`.
    - Actual-status rows refreshed: child 07, cumulative mapped total, seven-child lifecycle count, and P3 readiness.
  - Evidence Targets: four-file inventory, three-row mapping, six handoffs, target do-not-touch proof, and seven distributed Pn sets.
  - Actual-status Update: mark child 07 `missing -> correct`; set cross-plan validation ready only if cumulative mapping equals 98.
  - Commit Boundary: commit only child 07 plus roadmap/authoring ledger updates after acceptance and Supervisor review when authorized.

### P3: Cross-Plan Traceability And Authority Cutover

- Phase Goal: prove the authored campaign is lossless, internally consistent, independently executable, and safe to make authoritative.
- Phase Boundary:
  - In scope: structural/semantic audit, exact crosswalk, links, ownership, ledgers, dependency handoffs, Supervisor review, roadmap activation, and legacy status/pointer update.
  - Out of scope: implementation of any child, rewriting legacy history, target access, production/test/runtime changes, and technical re-audit.
  - Dependencies: P2-A through P2-G accepted with 28 standard child files present.
- Phase Implementation Rule: do not implement `P3` directly. Complete P3-A and record a frozen candidate hash/inventory; then P3-B may request Supervisor review and perform the conditional authority cutover.
- Ordered Slice List:
  - P3-A: Prove cross-plan completeness, traceability, and ownership.
  - P3-B: Obtain Supervisor PASS and switch authority atomically.

- [ ] P3-A: Prove cross-plan completeness, traceability, and ownership.
  - Goal: produce deterministic proof that the candidate multi-plan set is complete, non-overlapping, correctly linked, and structurally executable.
  - Scope Boundary:
    - Editable: roadmap validation/status fields and this authoring plan's ledgers.
    - Inspect-only: all seven child sets, legacy five-artifact source set, matrix, templates, and Git diff.
    - Preserve-only: child slice semantics, legacy source, production/tests/runtime, graph data, and target.
    - Out of scope: authority cutover or any corrective change outside the plan-created artifacts.
  - Non-Goals: do not accept approximate counts or a spot-check as proof of lossless conversion.
  - Pre-flight Questions:
    - Data source: frozen source snapshot, 29 candidate campaign outputs, 98-row crosswalk, and Git diff.
    - Display permission: N/A — docs-only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: authoring evidence/benchmark/actual-status and roadmap status.
    - UI behavior flow: N/A — no UI.
    - Docker runtime: N/A — no runtime change.
    - Playwright target: N/A — no runtime change.
    - Behavior test: deterministic parser/checks for files, headings, required fields, IDs, links, references, counts, order, and diff boundaries.
    - Cleanup/quarantine: no one-off validation file; remove plan-created debug material.
    - External side effects: read-only validation plus ledger/roadmap status updates.
    - N/A notes: source/destination document equivalence is the nearest real boundary.
  - Work Steps:
    1. Validate 1 roadmap, 7 child folders, 28 standard files, matching slugs/H1/metadata, 7 P0 sections, preserved source phases P1-P7, 7 Pn-A/B/C sets, the successor-actual-status rule in every child `Rules`, required slice fields, and all links.
       - UI flow check: N/A — no UI.
       - DB/data flow check: validate document inventory and ownership graph.
       - Render location check: results in authoring evidence and benchmark only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; inspect rendered Markdown and deterministic checker output.
       - Evidence target: `E3-P3A-STRUCT1`, `E3-P3A-LINK1`, and `E3-P3A-OWNER1`.
    2. Compare the frozen 98 source slice blocks with destinations, prove exact-once/order/required-field/semantic preservation, verify qualified evidence handoffs and Git scope, then freeze candidate hashes.
       - UI flow check: N/A — no UI.
       - DB/data flow check: require 98 source, 98 destination, 98 unique mappings, zero missing, zero duplicates, zero extras.
       - Render location check: crosswalk evidence and roadmap readiness status.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; run full rather than sampled equivalence checks.
       - Evidence target: `E3-P3A-MAP1`, `E3-P3A-FIELDS1`, `E3-P3A-DIFF1`, and `B-P3-A`.
  - Implementation Gate:
    - Record current Anvien freshness and docs file-detail limitation; no graph relationship claim is required for unindexed docs.
    - All seven P2 slices must pass; any missing file, placeholder, broken link, unqualified cross-child evidence ID, or mapping mismatch blocks P3-B.
  - Acceptance:
    - Source: legacy source remains unchanged and candidate hashes/inventory are recorded.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: 98/98 bijection with zero mismatch and one owner per mutable artifact.
    - Behavior test: all structural, successor-freshness, field, link, ordering, evidence-reference, and diff-boundary checks pass.
    - Cleanup/quarantine: validation creates no persistent artifact outside approved plan ledgers.
    - Evidence IDs: `E3-P3A-STRUCT1`, `E3-P3A-LINK1`, `E3-P3A-OWNER1`, `E3-P3A-MAP1`, `E3-P3A-FIELDS1`, `E3-P3A-DIFF1`.
    - Actual-status rows refreshed: roadmap candidate, seven child sets, 98-slice mapping, ownership, and cutover readiness.
  - Evidence Targets: complete candidate hashes, exact inventories, structural results, crosswalk results, link results, ownership check, and scoped Git diff.
  - Actual-status Update: mark campaign candidate `partial -> correct` only when every acceptance check is green; otherwise record exact failed child/slice and block cutover.
  - Commit Boundary: commit the validation-ledger/roadmap readiness update after acceptance and Supervisor review when authorized; do not change legacy authority in this commit.

- [ ] P3-B: Obtain Supervisor PASS and switch authority atomically.
  - Goal: make the verified roadmap and child plans authoritative without a period of missing or dual active authority.
  - Scope Boundary:
    - Editable: roadmap authority/status block, legacy plan metadata/pointer block, and authoring ledgers.
    - Inspect-only: frozen candidate set, legacy plan body/ledgers, Supervisor report, and Git diff.
    - Preserve-only: all accepted child bodies, legacy technical/history content, legacy ledgers, matrix content, code/tests/runtime, graph artifacts, and target.
    - Out of scope: changing any implementation contract or repairing a failed candidate outside the responsible earlier slice.
  - Non-Goals: do not mark the legacy plan superseded on a conditional, partial, or failed review.
  - Pre-flight Questions:
    - Data source: P3-A candidate hashes, complete authoring ledgers, and independent Supervisor verdict.
    - Display permission: N/A — documentation authority only.
    - DB read flow: N/A — no database.
    - DB write flow: N/A — no database.
    - Render location: roadmap authority block and legacy plan metadata/pointer.
    - UI behavior flow: N/A — no UI.
    - Docker runtime: N/A — no runtime.
    - Playwright target: N/A — no runtime.
    - Behavior test: pre/post authority assertion with exactly one active plan index.
    - Cleanup/quarantine: failed candidate fixes return to owning P1/P2/P3-A slice; no ad hoc patch is hidden in cutover.
    - External side effects: docs-only authority metadata change inside Anvien.
    - N/A notes: Supervisor and deterministic document checks replace runtime QA for this docs-only slice.
  - Work Steps:
    1. Call the Supervisor skill to compare the candidate set against this authoring plan, frozen legacy source, ledgers, counts, links, ownership, Git scope, and target boundary; route any failure to its responsible slice and repeat full review.
       - UI flow check: N/A — no UI.
       - DB/data flow check: Supervisor verifies evidence and crosswalk rather than runtime data.
       - Render location check: verdict recorded in authoring evidence and roadmap candidate status.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; independent document acceptance is the nearest boundary.
       - Evidence target: `E3-P3B-SUPERVISOR1`.
    2. Only after unconditional PASS, set the roadmap active, change only the legacy plan's authority metadata/pointer to `superseded / reference-only`, verify exactly one active authority, run scoped diff/detect-changes evidence, and update ledgers.
       - UI flow check: N/A — no UI.
       - DB/data flow check: authority state changes from legacy-active/candidate-inactive to legacy-reference/roadmap-active in one accepted slice.
       - Render location check: roadmap and legacy metadata/pointer only.
       - Mini QA for each completed implementation slice (MUST): N/A for browser plugins; verify links and authority assertions from disk.
       - Evidence target: `E3-P3B-CUTOVER1`, `E3-P3B-DIFF1`, and `E3-P3B-DETECT1`.
  - Implementation Gate:
    - Refresh Anvien and record docs limitations before edits; use impact only if a code or graph surface unexpectedly enters the diff, in which case stop because scope was violated.
    - P3-A must pass and Supervisor must issue an unconditional PASS against the frozen candidate hashes.
  - Acceptance:
    - Source: legacy body/ledgers are preserved; only status/pointer changed after PASS.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: exactly one active authority exists and roadmap links to seven accepted children.
    - Behavior test: pre/post authority, link, hash, and scoped-diff checks pass.
    - Cleanup/quarantine: no failed/rejected candidate artifact is made authoritative.
    - Evidence IDs: `E3-P3B-SUPERVISOR1`, `E3-P3B-CUTOVER1`, `E3-P3B-DIFF1`, `E3-P3B-DETECT1`.
    - Actual-status rows refreshed: legacy authority, roadmap authority, child readiness, and campaign execution gate.
  - Evidence Targets: unconditional Supervisor PASS, exact metadata diff, one-authority assertion, detect-changes result, and final candidate hashes.
  - Actual-status Update: transition `legacy active -> reference-only` and `roadmap candidate -> active` only after PASS; otherwise keep legacy active and record blocker.
  - Commit Boundary: commit the atomic docs-only authority cutover after acceptance when execution and commit are authorized.

- [ ] Pn-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, changed files, generated output, and validation results before closure.
  - Work Steps:
    1. Call the supervisor skill to review the full completed plan work.
    2. If supervisor fails the work, return to the responsible implementation workflow/skill for the failed scope only.
    3. Re-run supervisor review after the fix.
    4. Repeat until supervisor passes or records a blocker.
  - Implementation Gate: all planned implementation phases must be completed or explicitly blocked before this review.
  - Acceptance: supervisor review passes, or the plan records a blocker with evidence and no closure is performed.

- [ ] Pn-B: Remove dead work created during this plan.
  - Goal: ensure the final diff contains only artifacts that still serve the accepted plan.
  - Work Steps:
    1. Review files, sections, generated output, tests, temp files, and plan artifacts created or modified during this plan.
    2. Remove or rewrite any artifact made obsolete by actual-status findings, user corrections, failed approaches, or phase status updates.
    3. Verify no rejected approach, stale placeholder, unused generated output, or dead helper artifact remains in the final diff.
    4. Call supervisor to review the dead-work cleanup.
    5. If supervisor fails the cleanup, return to the responsible implementation workflow/skill for the failed cleanup scope only, then re-run supervisor review.
  - Implementation Gate: only remove artifacts created by this plan unless the user explicitly approves broader cleanup.
  - Acceptance: final `git diff/status` contains no dead plan-created artifacts, supervisor passes the cleanup, and evidence records what was removed or preserved.

- [ ] Pn-C: Close the plan.
  - Goal: finish validation, evidence, benchmark, detect-changes, commit, and final status.
  - Work Steps:
    1. Run the required final structural, link, crosswalk, ownership, authority, and Git-scope validation. Full build is N/A because this plan changes no executable source; record that reason rather than claiming a build.
    2. Docker runtime is N/A because no app/runtime surface changes; if executable files unexpectedly appear in the diff, block closure and return to the responsible slice.
    3. Browser/Playwright validation is N/A because no public runtime or UI-facing behavior changes; validate rendered Markdown and disk links at the nearest real boundary.
    4. Regenerate only plan-derived inventories if accepted plan source requires it; do not regenerate graph/runtime outputs for docs-only closure.
    5. Run Anvien detect-changes before commit and verify that no implementation file, graph output, repository-root artifact, or target content entered scope.
    6. Record final validation, detect-changes, benchmark, candidate/legacy hashes, authority state, and commit evidence.
    7. Commit the completed docs-only scope when authorized and verify the worktree state.
  - Implementation Gate: Pn-A and Pn-B must pass or record blockers.
  - Acceptance: final evidence is recorded, required authorized commits exist, the roadmap is the sole active authority, the legacy plan is preserved reference-only, and the worktree state is known.

## Risk Notes

- Source drift: the legacy plan may change after this plan is written. Any hash mismatch requires a new baseline and slice inventory before authoring continues.
- Silent loss: manual copying can omit nested IDs such as `P1-C0A`, `P3-B2A`, or `P4-C1I`; exact-ID parsing and bijection checks are mandatory.
- Semantic drift: rewriting or normalizing prose can alter technical contracts. Copy source phase blocks verbatim and restrict additions to the minimum standalone child structure.
- Dual authority: marking both legacy and roadmap active creates conflicting execution truth. Cutover is conditional and atomic after Supervisor PASS.
- Incomplete children: a child folder or roadmap entry is not a complete child plan. Each child must have four populated files, P0, its unchanged source phase/slices, and Pn-A/B/C.
- Ledger collisions: source evidence IDs may repeat across child ledgers only as copied source facts; every cross-child reference must include the child slug and exact ID.
- Shared-file ambiguity: `index-reader-matrix.md` may be referenced broadly, but only child 02 owns mutation.
- One-file responsibility erosion: do not turn the roadmap or authoring ledger into a catch-all implementation plan; links across modules are allowed only in service of the file's single coordination/evidence/status responsibility.
- Scope contamination: `E:\cheapapp.org`, production code, tests, graph data, and repository-root paths remain untouched throughout plan authoring.
- Oversized output: child 02 remains large because its 42 existing slices belong to one coupled persistence/cutover phase. This plan preserves the user's phase-to-plan rule rather than inventing a second decomposition.
