# Versioned Persistence and Identity v2 Cutover Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Source phase: legacy `P2`

## Evidence Rules

The evidence file explains why the work is known to be correct.

It should contain:

- metadata and companion files;
- evidence rules or evidence template;
- evidence sections such as `E0`, `E1`, or sections by phase/task;
- user report or problem evidence;
- source inspection, codebase facts, and document facts;
- commands run and pass/fail result;
- impact or blast-radius evidence when code/graph behavior changes;
- implementation evidence: files changed and behavior changed;
- validation evidence: build, tests, e2e, screenshots, or traces;
- failures encountered and how they were handled;
- detect-changes before commit;
- commit hash and closure evidence.

Evidence can reference short metric traces, but long metric tables belong in the benchmark file.

### Evidence ID Naming

Use stable, phase-scoped evidence IDs so `plan.md`, `actual-status.md`, `benchmark.md`, and later agents can reference exact proof without ambiguity.

Format:

```text
E<phase>-<item>-<kind><n>
```

Rules:

- `E<phase>` matches the plan phase number: `E0` for `P0`, `E1` for `P1`, `E2` for `P2`, and so on.
- `<item>` matches the checklist item without the dash: `P0A`, `P1A`, `P2B`.
- `<kind>` is plan-local. Choose a short uppercase token that is meaningful for this repo and this plan.
- `<n>` is a 1-based sequence number within that phase item and kind.
- Keep the same `<kind>` meaning stable inside one plan.
- Do not reuse an evidence ID for different facts.
- Reference exact evidence IDs from `actual-status.md` and `benchmark.md`; avoid referencing only broad section IDs such as `E1`.
- Use ranges such as `E0-P0A-FD1..E0-P0A-FD17` only for compact inventory summaries; use exact IDs when a specific status decision depends on a specific fact.
- If nearby plans already use a clear local evidence naming style, follow that style instead of inventing a new one.

Examples only:

- `E0-P0A-SRC1`
- `E0-P0A-GRAPH1`
- `E1-P1A-ROUTE1`
- `E2-P2B-KEYBOARD1`
- `E2-P2B-DETECT1`

Evidence sections must follow the plan phases:

- `E0` corresponds to `P0`.
- `E1` corresponds to `P1`.
- `E2` corresponds to `P2`.
- Use exact evidence IDs inside each section, not broad section IDs as proof.
- Each evidence section must name the plan phase or checklist item it supports.
- Do not invent fixed evidence categories; record the evidence required by the matching plan phase.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-PLAN1`: This standard child set contains plan, evidence, benchmark, and actual-status files with slug `2026-07-28-02-versioned-persistence-and-v2-cutover`; local P1 has 42 slices mapped from legacy P2.
- `E0-P0A-SOURCE1`: Legacy source phase P2 spans source lines 899-2890, remains under frozen SHA-256 `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB`, and maps to 42 unique local slice IDs.
- `E0-P0A-GRAPH1`: `anvien analyze E:\Anvien --force --json` completed after the three-root structural commit at HEAD `55bf021f813adc8f8bb61daf57ee95ff0c8382c7`; `anvien status` reported indexed/current commit `55bf021`, analyzed at `2026-07-28T07:34:44Z`, stale `false`, and up to date. The graph contained 1,525 files, 676 parsed code files, 726 documents, 84,783 nodes, and 123,627 relationships.
- `E0-P0A-FD10`: Fresh `anvien file-detail` for this child plan at indexed/current commit `55bf021f813adc8f8bb61daf57ee95ff0c8382c7` reports parsed Markdown/docs, low risk, zero related files/relationships, zero unresolved references, and no staleness before the P2-B roadmap candidate link is added. This is the truthful committed-basis pre-link baseline, not an acceptance claim.
- `E0-P0A-GRAPH2`: After the P2-B candidate links and ledger refresh, fresh `anvien analyze E:\Anvien --force --json` completed at `2026-07-28T07:44:40Z` on indexed/current commit `55bf021f813adc8f8bb61daf57ee95ff0c8382c7`, with 1,525 files, 676 parsed code files, 726 documents, 84,783 nodes, and 123,631 relationships; `anvien status` was up to date.
- `E0-P0A-FD11`: Post-link `anvien file-detail` for this child plan at the same fresh index reports parsed Markdown/docs, low risk, one inbound roadmap `IMPORTS` relationship, zero unresolved references, and no staleness. The corresponding roadmap file reports eight outbound child-ledger imports.
- `E0-P0A-GIT1`: P2-A1 is committed at `55bf021f813adc8f8bb61daf57ee95ff0c8382c7`; the pre-correction `git status --short` contains only this untracked four-file child-02 folder and its scoped red-team report. No production, test, runtime, graph-output, repository-root, or target path is in the authoring worktree.
- `E0-P0A-GIT2`: After P2-B candidate authoring, all changed/untracked paths remain within the approved authoring/roadmap/child-02/report scope; no production, test, runtime, graph-output, repository-root, or target path appears. Child-02 remains uncommitted pending Supervisor PASS.
- `E0-P0A-STATUS1`: No production path changed between legacy baseline commit `1932359b` and current child-authoring basis `55bf021f`; the intervening accepted commits are plan/report-only. The inherited current behavior therefore remains: index versions/generations are incomplete, readers are not uniformly guarded, IDs are parsed semantically, and publication is not atomic.
- `E0-P0A-DEPENDENCY1`: Child 01's authoring set is accepted/committed at `b760d156`, and its path-preserving three-root move is committed at `55bf021f`. Those facts unblock child-02 plan authoring only. The production implementation handoff `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1` does not yet exist and remains a separate hard gate; it must bind the accepted identity authority, shadow proof, benchmark, review, and implementation-commit evidence before local P1 can open. Explicit owner authorization for implementation is also absent.
- `E0-P0A-NEXTSTATUSRULE1`: The user-required successor-freshness invariant is now explicit in this child plan's `Rules` and Pn-C: child 02 cannot close or hand off until child 03's actual-status, refresh log, and affected next actions/work steps are refreshed from the latest accepted child-02 evidence; the update is reserved as `E2-PNC-NEXTSTATUS1` and is passed cross-child only as `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-NEXTSTATUS1`. Missing or stale successor status blocks closure.
- `E0-P0A-MATRIX1`: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/index-reader-matrix.md` exists once in the preserved legacy root with SHA-256 `A6FE5148341E048425E6240B403D21F37941299319FB194A24B5CBE5E4F97409`; it is unchanged in the P2-B authoring worktree. This child is its sole future mutation owner, while authoring remains inspect-only.
- `E0-P0A-BOUNDARY1`: `E:\cheapapp.org` is an in-place validation subject only for explicitly opened target slices; no target source, report, probe, fixture, or temp artifact was used to author this child.
- `E0-P0A-SCANNER1`: The exact eight scanner omissions remain quarantined, wrong-but-out-of-scope, with zero additional omissions required.
- `E0-P0A-FD1`: inherited source P0 file-detail for `internal/analyze/analyze.go` reports 182 related files/items; analyze/build/load/snapshot orchestration; high; publication CRITICAL. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD2`: inherited source P0 file-detail for `internal/repo/types.go` reports 72 related files/items; persisted repository metadata; high; version/generation fields. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD3`: inherited source P0 file-detail for `internal/graph/types.go` reports 238 related files/items; canonical graph records; high/CRITICAL identity storage. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD4`: inherited source P0 file-detail for `internal/lbugload/csv.go` reports 19 related files/items; Ladybug schema/CSV projection; high duplicate/parity risk. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD5`: inherited source P0 file-detail for `internal/httpapi/graph.go` reports 22 related files/items; HTTP graph contract; high public contract. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD6`: inherited source P0 file-detail for `anvien-web/src/services/backend-client.ts` reports 24 related files/items; Web stream/version parsing; high public runtime. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD7`: inherited source P0 file-detail for `internal/repo/registry.go / internal/group/sync.go / internal/group/storage.go` reports bounded related files/items; global registry and group publication; high generation scope. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD8`: inherited source P0 file-detail for `internal/mcp/resource_cache.go / internal/filecontext/* / embedding stores` reports bounded related files/items; cache and embedding readers; high stale-reference scope. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD9`: inherited source P0 file-detail for `native and fallback Cypher readers` reports bounded related files/items; alternate query backends; high parity scope. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`, `P1-A1`, `P1-A2`, `P1-A3`, `P1-A4`, `P1-A5`, `P1-A6`, `P1-A7`, `P1-A8`, `P1-A9`, `P1-A10`, `P1-A11`, `P1-A12`, `P1-A13`, `P1-A14`, `P1-A15`, `P1-B`, `P1-B1`, `P1-B2`, `P1-B3`, `P1-B4`, `P1-C`, `P1-C1`, `P1-C2`, `P1-C3`, `P1-C4`, `P1-C5`, `P1-C6`, `P1-D`, `P1-D1`, `P1-D2`, `P1-E`, `P1-E1`, `P1-E2`, `P1-F`, `P1-F1`, `P1-F2`, `P1-F3`, `P1-F4`, `P1-F5`, `P1-F6`, `P1-G`

Source phase: legacy `P2`; every local evidence ID preserves its source-phase provenance through this child slug.

Every row below is reserved evidence only until its exact command, artifact, result, Supervisor verdict, detect-changes result when applicable, and commit are recorded. P1 slices never advance automatically.

| Slice | Required evidence set | Status |
|-------|-----------------------|--------|
| P1-A | `E1-P1A-IMPACT1`, `E1-P1A-SRC1`, `E1-P1A-BUILD1`, `E1-P1A-TEST1`, `E1-P1A-REVIEW1`, `E1-P1A-DETECT1`, `E1-P1A-COMMIT1` | pending |
| P1-A1 | `E1-P1A1-MATRIX1`, `E1-P1A1-MATRIXREVIEW1`, `E1-P1A1-R01..E1-P1A1-R195`, `E1-P1A1-REVIEW1`, `E1-P1A1-COMMIT1` | pending |
| P1-A2 | `E1-P1A2-IMPACT1`, `E1-P1A2-SRC1`, `E1-P1A2-BUILD1`, `E1-P1A2-TEST1`, `E1-P1A2-S0GUARD1`, `E1-P1A2-REVIEW1`, `E1-P1A2-DETECT1`, `E1-P1A2-COMMIT1` | pending |
| P1-A3 | `E1-P1A3-IMPACT1`, `E1-P1A3-SRC1`, `E1-P1A3-BUILD1`, `E1-P1A3-TEST1`, `E1-P1A3-S1GUARD1`, `E1-P1A3-REVIEW1`, `E1-P1A3-DETECT1`, `E1-P1A3-COMMIT1` | pending |
| P1-A4 | `E1-P1A4-IMPACT1`, `E1-P1A4-SRC1`, `E1-P1A4-BUILD1`, `E1-P1A4-TEST1`, `E1-P1A4-S2GUARD1`, `E1-P1A4-REVIEW1`, `E1-P1A4-DETECT1`, `E1-P1A4-COMMIT1` | pending |
| P1-A5 | `E1-P1A5-IMPACT1`, `E1-P1A5-SRC1`, `E1-P1A5-BUILD1`, `E1-P1A5-TEST1`, `E1-P1A5-S3GUARD1`, `E1-P1A5-REVIEW1`, `E1-P1A5-DETECT1`, `E1-P1A5-COMMIT1` | pending |
| P1-A6 | `E1-P1A6-IMPACT1`, `E1-P1A6-SRC1`, `E1-P1A6-BUILD1`, `E1-P1A6-TEST1`, `E1-P1A6-S4GUARD1`, `E1-P1A6-REVIEW1`, `E1-P1A6-DETECT1`, `E1-P1A6-COMMIT1` | pending |
| P1-A7 | `E1-P1A7-IMPACT1`, `E1-P1A7-SRC1`, `E1-P1A7-BUILD1`, `E1-P1A7-TEST1`, `E1-P1A7-S5GUARD1`, `E1-P1A7-REVIEW1`, `E1-P1A7-DETECT1`, `E1-P1A7-COMMIT1` | pending |
| P1-A8 | `E1-P1A8-IMPACT1`, `E1-P1A8-SRC1`, `E1-P1A8-BUILD1`, `E1-P1A8-PLAY1`, `E1-P1A8-S6GUARD1`, `E1-P1A8-REVIEW1`, `E1-P1A8-DETECT1`, `E1-P1A8-COMMIT1` | pending |
| P1-A9 | `E1-P1A9-IMPACT1`, `E1-P1A9-SRC1`, `E1-P1A9-BUILD1`, `E1-P1A9-TEST1`, `E1-P1A9-S7GUARD1`, `E1-P1A9-REVIEW1`, `E1-P1A9-DETECT1`, `E1-P1A9-COMMIT1` | pending |
| P1-A10 | `E1-P1A10-IMPACT1`, `E1-P1A10-SRC1`, `E1-P1A10-BUILD1`, `E1-P1A10-TEST1`, `E1-P1A10-S8GUARD1`, `E1-P1A10-REVIEW1`, `E1-P1A10-DETECT1`, `E1-P1A10-COMMIT1` | pending |
| P1-A11 | `E1-P1A11-IMPACT1`, `E1-P1A11-SRC1`, `E1-P1A11-BUILD1`, `E1-P1A11-TEST1`, `E1-P1A11-S9GUARD1`, `E1-P1A11-REVIEW1`, `E1-P1A11-DETECT1`, `E1-P1A11-COMMIT1` | pending |
| P1-A12 | `E1-P1A12-IMPACT1`, `E1-P1A12-SRC1`, `E1-P1A12-BUILD1`, `E1-P1A12-TEST1`, `E1-P1A12-S10GUARD1`, `E1-P1A12-REVIEW1`, `E1-P1A12-DETECT1`, `E1-P1A12-COMMIT1` | pending |
| P1-A13 | `E1-P1A13-IMPACT1`, `E1-P1A13-SRC1`, `E1-P1A13-BUILD1`, `E1-P1A13-TEST1`, `E1-P1A13-S10GUARD1`, `E1-P1A13-REVIEW1`, `E1-P1A13-DETECT1`, `E1-P1A13-COMMIT1` | pending |
| P1-A14 | `E1-P1A14-IMPACT1`, `E1-P1A14-SRC1`, `E1-P1A14-BUILD1`, `E1-P1A14-TEST1`, `E1-P1A14-S11GUARD1`, `E1-P1A14-REVIEW1`, `E1-P1A14-DETECT1`, `E1-P1A14-COMMIT1` | pending |
| P1-A15 | `E1-P1A15-IMPACT1`, `E1-P1A15-SRC1`, `E1-P1A15-BUILD1`, `E1-P1A15-TEST1`, `E1-P1A15-S11GUARD1`, `E1-P1A15-REVIEW1`, `E1-P1A15-DETECT1`, `E1-P1A15-COMMIT1` | pending |
| P1-B | `E1-P1B-IMPACT1`, `E1-P1B-SRC1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-REVIEW1`, `E1-P1B-DETECT1`, `E1-P1B-COMMIT1` | pending |
| P1-B1 | `E1-P1B1-IMPACT1`, `E1-P1B1-SRC1`, `E1-P1B1-BUILD1`, `E1-P1B1-TEST1`, `E1-P1B1-REVIEW1`, `E1-P1B1-DETECT1`, `E1-P1B1-COMMIT1` | pending |
| P1-B2 | `E1-P1B2-IMPACT1`, `E1-P1B2-SRC1`, `E1-P1B2-BUILD1`, `E1-P1B2-TEST1`, `E1-P1B2-REVIEW1`, `E1-P1B2-DETECT1`, `E1-P1B2-COMMIT1` | pending |
| P1-B3 | `E1-P1B3-IMPACT1`, `E1-P1B3-SRC1`, `E1-P1B3-BUILD1`, `E1-P1B3-TEST1`, `E1-P1B3-REVIEW1`, `E1-P1B3-DETECT1`, `E1-P1B3-COMMIT1` | pending |
| P1-B4 | `E1-P1B4-IMPACT1`, `E1-P1B4-SRC1`, `E1-P1B4-BUILD1`, `E1-P1B4-TEST1`, `E1-P1B4-REVIEW1`, `E1-P1B4-DETECT1`, `E1-P1B4-COMMIT1` | pending |
| P1-C | `E1-P1C-IMPACT1`, `E1-P1C-SRC1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` | pending |
| P1-C1 | `E1-P1C1-IMPACT1`, `E1-P1C1-SRC1`, `E1-P1C1-BUILD1`, `E1-P1C1-TEST1`, `E1-P1C1-REVIEW1`, `E1-P1C1-DETECT1`, `E1-P1C1-COMMIT1` | pending |
| P1-C2 | `E1-P1C2-IMPACT1`, `E1-P1C2-SRC1`, `E1-P1C2-BUILD1`, `E1-P1C2-TEST1`, `E1-P1C2-REVIEW1`, `E1-P1C2-DETECT1`, `E1-P1C2-COMMIT1` | pending |
| P1-C3 | `E1-P1C3-IMPACT1`, `E1-P1C3-SRC1`, `E1-P1C3-BUILD1`, `E1-P1C3-TEST1`, `E1-P1C3-REVIEW1`, `E1-P1C3-DETECT1`, `E1-P1C3-COMMIT1` | pending |
| P1-C4 | `E1-P1C4-IMPACT1`, `E1-P1C4-SRC1`, `E1-P1C4-BUILD1`, `E1-P1C4-TEST1`, `E1-P1C4-REVIEW1`, `E1-P1C4-DETECT1`, `E1-P1C4-COMMIT1` | pending |
| P1-C5 | `E1-P1C5-IMPACT1`, `E1-P1C5-SRC1`, `E1-P1C5-BUILD1`, `E1-P1C5-TEST1`, `E1-P1C5-REVIEW1`, `E1-P1C5-DETECT1`, `E1-P1C5-COMMIT1` | pending |
| P1-C6 | `E1-P1C6-IMPACT1`, `E1-P1C6-SRC1`, `E1-P1C6-BUILD1`, `E1-P1C6-TEST1`, `E1-P1C6-REVIEW1`, `E1-P1C6-DETECT1`, `E1-P1C6-COMMIT1` | pending |
| P1-D | `E1-P1D-IMPACT1`, `E1-P1D-SRC1`, `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-REVIEW1`, `E1-P1D-DETECT1`, `E1-P1D-COMMIT1` | pending |
| P1-D1 | `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`, `E1-P1D1-REVIEW1`, `E1-P1D1-DETECT1`, `E1-P1D1-COMMIT1` | pending |
| P1-D2 | `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`, `E1-P1D2-DETECT1`, `E1-P1D2-COMMIT1` | pending |
| P1-E | `E1-P1E-IMPACT1`, `E1-P1E-SRC1`, `E1-P1E-BUILD1`, `E1-P1E-TEST1`, `E1-P1E-REVIEW1`, `E1-P1E-DETECT1`, `E1-P1E-COMMIT1` | pending |
| P1-E1 | `E1-P1E1-IMPACT1`, `E1-P1E1-SRC1`, `E1-P1E1-BUILD1`, `E1-P1E1-PLAY1`, `E1-P1E1-REVIEW1`, `E1-P1E1-DETECT1`, `E1-P1E1-COMMIT1` | pending |
| P1-E2 | `E1-P1E2-BUILD1`, `E1-P1E2-S0BASE1`, `E1-P1E2-S1BASE1`, `E1-P1E2-S2BASE1`, `E1-P1E2-S3BASE1`, `E1-P1E2-S4BASE1`, `E1-P1E2-S5BASE1`, `E1-P1E2-S6BASE1`, `E1-P1E2-S7BASE1`, `E1-P1E2-S8BASE1`, `E1-P1E2-S9BASE1`, `E1-P1E2-S10BASE1`, `E1-P1E2-S11BASE1`, `E1-P1E2-MATRIX1`, `E1-P1E2-PLAY1`, `E1-P1E2-REVIEW1`, `E1-P1E2-DETECT1`, `E1-P1E2-COMMIT1` | pending |
| P1-F | `E1-P1F-IMPACT1`, `E1-P1F-SRC1`, `E1-P1F-BUILD1`, `E1-P1F-TEST1`, `E1-P1F-REVIEW1`, `E1-P1F-DETECT1`, `E1-P1F-COMMIT1` | pending |
| P1-F1 | `E1-P1F1-IMPACT1`, `E1-P1F1-SRC1`, `E1-P1F1-BUILD1`, `E1-P1F1-TEST1`, `E1-P1F1-REVIEW1`, `E1-P1F1-DETECT1`, `E1-P1F1-COMMIT1` | pending |
| P1-F2 | `E1-P1F2-IMPACT1`, `E1-P1F2-SRC1`, `E1-P1F2-BUILD1`, `E1-P1F2-TEST1`, `E1-P1F2-REVIEW1`, `E1-P1F2-DETECT1`, `E1-P1F2-COMMIT1` | pending |
| P1-F3 | `E1-P1F3-IMPACT1`, `E1-P1F3-SRC1`, `E1-P1F3-BUILD1`, `E1-P1F3-TEST1`, `E1-P1F3-REVIEW1`, `E1-P1F3-DETECT1`, `E1-P1F3-COMMIT1` | pending |
| P1-F4 | `E1-P1F4-IMPACT1`, `E1-P1F4-SRC1`, `E1-P1F4-BUILD1`, `E1-P1F4-TEST1`, `E1-P1F4-REVIEW1`, `E1-P1F4-DETECT1`, `E1-P1F4-COMMIT1` | pending |
| P1-F5 | `E1-P1F5-IMPACT1`, `E1-P1F5-SRC1`, `E1-P1F5-BUILD1`, `E1-P1F5-TEST1`, `E1-P1F5-REVIEW1`, `E1-P1F5-DETECT1`, `E1-P1F5-COMMIT1` | pending |
| P1-F6 | `E1-P1F6-BUILD1`, `E1-P1F6-FAULT1`, `E1-P1F6-RECOVERY1`, `E1-P1F6-REVIEW1`, `E1-P1F6-DETECT1`, `E1-P1F6-COMMIT1` | pending |
| P1-G | `E1-P1G-PREBASE1`, `E1-P1G-CANDIDATE1`, `E1-P1G-IMPACT1`, `E1-P1G-SRC1`, `E1-P1G-CUTOVER1`, `E1-P1G-BUILD1`, `E1-P1G-RUNTIME1`, `E1-P1G-PLAY1`, `E1-P1G-ROLLBACK1`, `E1-P1G-REVIEW1`, `E1-P1G-DETECT1`, `E1-P1G-COMMIT1` | pending |

### P1-A1 reader inventory and guard ownership

- `E1-P1A1-MATRIX1`: frozen source-derived matrix with exact path/function, truthful backend/layout, dispatcher/non-reader classification, later guard-owner slice, fixture, and expected transport failure.
- `E1-P1A1-MATRIXREVIEW1`: fresh source scan, matrix SHA-256, contiguous/unique row check, anchor existence, `rows_classified == rows_total`, `unassigned_rows == 0`, and `unlisted_readers == 0`.
- `E1-P1A1-R01..E1-P1A1-R195`: one exact classification proof per current seed row; if the fresh scan adds rows, continue the row/evidence numbering. These proofs establish source anchor, backend, surface tags, and assigned guard owner. They do not claim runtime guard PASS.

| Surface | Guard slice | Owner boundary | Runtime row-result evidence |
|---------|-------------|----------------|-----------------------------|
| `S0` | `P1-A2` | Graph JSON/repository metadata | `E1-P1A2-S0GUARD1` |
| `S1` | `P1-A3` | native Ladybug | `E1-P1A3-S1GUARD1` |
| `S2` | `P1-A4` | Go/fallback Cypher | `E1-P1A4-S2GUARD1` |
| `S3` | `P1-A5` | CLI | `E1-P1A5-S3GUARD1` |
| `S4` | `P1-A6` | MCP | `E1-P1A6-S4GUARD1` |
| `S5` | `P1-A7` | HTTP | `E1-P1A7-S5GUARD1` |
| `S6` | `P1-A8` | Web | `E1-P1A8-S6GUARD1` |
| `S7` | `P1-A9` | file-context cache | `E1-P1A9-S7GUARD1` |
| `S8` | `P1-A10` | HTTP/MCP resource cache | `E1-P1A10-S8GUARD1` |
| `S9` | `P1-A11` | embeddings | `E1-P1A11-S9GUARD1` |
| `S10` | `P1-A12` | global repository registry | `E1-P1A12-S10GUARD1` |
| `S10` | `P1-A13` | group registry/contracts | `E1-P1A13-S10GUARD1` |
| `S11` | `P1-A14` | process projections | `E1-P1A14-S11GUARD1` |
| `S11` | `P1-A15` | community/cluster projections | `E1-P1A15-S11GUARD1` |

Each guard slice records the exact matrix row IDs it executed. A parent command, router, selector, transport, or dispatcher never substitutes for an exact child row. Native and fallback query paths remain separate.

### P1-E2 pre-cutover canonical baselines

| Surface | Baseline evidence | Required comparison |
|---------|-------------------|---------------------|
| `S0` | `E1-P1E2-S0BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S1` | `E1-P1E2-S1BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S2` | `E1-P1E2-S2BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S3` | `E1-P1E2-S3BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S4` | `E1-P1E2-S4BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S5` | `E1-P1E2-S5BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S6` | `E1-P1E2-S6BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S7` | `E1-P1E2-S7BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S8` | `E1-P1E2-S8BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S9` | `E1-P1E2-S9BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S10` | `E1-P1E2-S10BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S11` | `E1-P1E2-S11BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |

- `E1-P1E2-MATRIX1`: complete built-runtime execution result with `rows_passed == rows_total`, `unlisted_readers == 0`, no skipped row, and no parent-row substitution.
- `E1-P1E2-PLAY1`: reusable Playwright JSON+Markdown, screenshots/traces, and visual inspection for the exact S6 row union against the real built Docker runtime.
- Any S0-S11 failure reopens only the responsible already-committed storage/consumer/guard slice; P1-E2 does not patch code.

### P1 publication and cutover proofs

- `E1-P1F6-FAULT1` and `E1-P1F6-RECOVERY1`: complete staging/write/fsync/CAS/restart/concurrency/lease/GC fault matrix with exact prior-state hashes, pointers, generation vectors, and queryability after every failure.
- `E1-P1G-PREBASE1`: at least five identical pre-cutover v1 runs recording analyze median, Ladybug-load median, native-Cypher p95, fallback-Cypher p95, graph size, peak RSS, and the bound corpus/config/build/machine/cache policy.
- `E1-P1G-CANDIDATE1`: at least five staged non-active v2 runs with the identical methodology and per-metric delta before active CAS.
- `E1-P1G-PLAY1`: built Docker/Web supported, mismatch, and legacy-ambiguity evidence; an old client consumes zero v2 records.
- `E1-P1G-ROLLBACK1`: active-v2 rollback restores the prior queryable generation/registry/group vector without mixed cache or embedding state.

### Preserved source traceability binding

| Local slice | Required exact local evidence IDs |
|-------------|-----------------------------------|
| P1-A | `E1-P1A-IMPACT1`, `E1-P1A-SRC1`, `E1-P1A-BUILD1`, `E1-P1A-TEST1`, `E1-P1A-REVIEW1`, `E1-P1A-DETECT1`, `E1-P1A-COMMIT1` |
| P1-A1 | `E1-P1A1-MATRIX1`, `E1-P1A1-MATRIXREVIEW1`, `E1-P1A1-R01..E1-P1A1-R195`, `E1-P1A1-REVIEW1`, `E1-P1A1-COMMIT1` |
| P1-A2 | `E1-P1A2-IMPACT1`, `E1-P1A2-SRC1`, `E1-P1A2-BUILD1`, `E1-P1A2-TEST1`, `E1-P1A2-S0GUARD1`, `E1-P1A2-REVIEW1`, `E1-P1A2-DETECT1`, `E1-P1A2-COMMIT1` |
| P1-A3 | `E1-P1A3-IMPACT1`, `E1-P1A3-SRC1`, `E1-P1A3-BUILD1`, `E1-P1A3-TEST1`, `E1-P1A3-S1GUARD1`, `E1-P1A3-REVIEW1`, `E1-P1A3-DETECT1`, `E1-P1A3-COMMIT1` |
| P1-A4 | `E1-P1A4-IMPACT1`, `E1-P1A4-SRC1`, `E1-P1A4-BUILD1`, `E1-P1A4-TEST1`, `E1-P1A4-S2GUARD1`, `E1-P1A4-REVIEW1`, `E1-P1A4-DETECT1`, `E1-P1A4-COMMIT1` |
| P1-A5 | `E1-P1A5-IMPACT1`, `E1-P1A5-SRC1`, `E1-P1A5-BUILD1`, `E1-P1A5-TEST1`, `E1-P1A5-S3GUARD1`, `E1-P1A5-REVIEW1`, `E1-P1A5-DETECT1`, `E1-P1A5-COMMIT1` |
| P1-A6 | `E1-P1A6-IMPACT1`, `E1-P1A6-SRC1`, `E1-P1A6-BUILD1`, `E1-P1A6-TEST1`, `E1-P1A6-S4GUARD1`, `E1-P1A6-REVIEW1`, `E1-P1A6-DETECT1`, `E1-P1A6-COMMIT1` |
| P1-A7 | `E1-P1A7-IMPACT1`, `E1-P1A7-SRC1`, `E1-P1A7-BUILD1`, `E1-P1A7-TEST1`, `E1-P1A7-S5GUARD1`, `E1-P1A7-REVIEW1`, `E1-P1A7-DETECT1`, `E1-P1A7-COMMIT1` |
| P1-A8 | `E1-P1A8-IMPACT1`, `E1-P1A8-SRC1`, `E1-P1A8-BUILD1`, `E1-P1A8-PLAY1`, `E1-P1A8-S6GUARD1`, `E1-P1A8-REVIEW1`, `E1-P1A8-DETECT1`, `E1-P1A8-COMMIT1` |
| P1-A9 | `E1-P1A9-IMPACT1`, `E1-P1A9-SRC1`, `E1-P1A9-BUILD1`, `E1-P1A9-TEST1`, `E1-P1A9-S7GUARD1`, `E1-P1A9-REVIEW1`, `E1-P1A9-DETECT1`, `E1-P1A9-COMMIT1` |
| P1-A10 | `E1-P1A10-IMPACT1`, `E1-P1A10-SRC1`, `E1-P1A10-BUILD1`, `E1-P1A10-TEST1`, `E1-P1A10-S8GUARD1`, `E1-P1A10-REVIEW1`, `E1-P1A10-DETECT1`, `E1-P1A10-COMMIT1` |
| P1-A11 | `E1-P1A11-IMPACT1`, `E1-P1A11-SRC1`, `E1-P1A11-BUILD1`, `E1-P1A11-TEST1`, `E1-P1A11-S9GUARD1`, `E1-P1A11-REVIEW1`, `E1-P1A11-DETECT1`, `E1-P1A11-COMMIT1` |
| P1-A12 | `E1-P1A12-IMPACT1`, `E1-P1A12-SRC1`, `E1-P1A12-BUILD1`, `E1-P1A12-TEST1`, `E1-P1A12-S10GUARD1`, `E1-P1A12-REVIEW1`, `E1-P1A12-DETECT1`, `E1-P1A12-COMMIT1` |
| P1-A13 | `E1-P1A13-IMPACT1`, `E1-P1A13-SRC1`, `E1-P1A13-BUILD1`, `E1-P1A13-TEST1`, `E1-P1A13-S10GUARD1`, `E1-P1A13-REVIEW1`, `E1-P1A13-DETECT1`, `E1-P1A13-COMMIT1` |
| P1-A14 | `E1-P1A14-IMPACT1`, `E1-P1A14-SRC1`, `E1-P1A14-BUILD1`, `E1-P1A14-TEST1`, `E1-P1A14-S11GUARD1`, `E1-P1A14-REVIEW1`, `E1-P1A14-DETECT1`, `E1-P1A14-COMMIT1` |
| P1-A15 | `E1-P1A15-IMPACT1`, `E1-P1A15-SRC1`, `E1-P1A15-BUILD1`, `E1-P1A15-TEST1`, `E1-P1A15-S11GUARD1`, `E1-P1A15-REVIEW1`, `E1-P1A15-DETECT1`, `E1-P1A15-COMMIT1` |
| P1-B | `E1-P1B-IMPACT1`, `E1-P1B-SRC1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-REVIEW1`, `E1-P1B-DETECT1`, `E1-P1B-COMMIT1` |
| P1-B1 | `E1-P1B1-IMPACT1`, `E1-P1B1-SRC1`, `E1-P1B1-BUILD1`, `E1-P1B1-TEST1`, `E1-P1B1-REVIEW1`, `E1-P1B1-DETECT1`, `E1-P1B1-COMMIT1` |
| P1-B2 | `E1-P1B2-IMPACT1`, `E1-P1B2-SRC1`, `E1-P1B2-BUILD1`, `E1-P1B2-TEST1`, `E1-P1B2-REVIEW1`, `E1-P1B2-DETECT1`, `E1-P1B2-COMMIT1` |
| P1-B3 | `E1-P1B3-IMPACT1`, `E1-P1B3-SRC1`, `E1-P1B3-BUILD1`, `E1-P1B3-TEST1`, `E1-P1B3-REVIEW1`, `E1-P1B3-DETECT1`, `E1-P1B3-COMMIT1` |
| P1-B4 | `E1-P1B4-IMPACT1`, `E1-P1B4-SRC1`, `E1-P1B4-BUILD1`, `E1-P1B4-TEST1`, `E1-P1B4-REVIEW1`, `E1-P1B4-DETECT1`, `E1-P1B4-COMMIT1` |
| P1-C | `E1-P1C-IMPACT1`, `E1-P1C-SRC1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` |
| P1-C1 | `E1-P1C1-IMPACT1`, `E1-P1C1-SRC1`, `E1-P1C1-BUILD1`, `E1-P1C1-TEST1`, `E1-P1C1-REVIEW1`, `E1-P1C1-DETECT1`, `E1-P1C1-COMMIT1` |
| P1-C2 | `E1-P1C2-IMPACT1`, `E1-P1C2-SRC1`, `E1-P1C2-BUILD1`, `E1-P1C2-TEST1`, `E1-P1C2-REVIEW1`, `E1-P1C2-DETECT1`, `E1-P1C2-COMMIT1` |
| P1-C3 | `E1-P1C3-IMPACT1`, `E1-P1C3-SRC1`, `E1-P1C3-BUILD1`, `E1-P1C3-TEST1`, `E1-P1C3-REVIEW1`, `E1-P1C3-DETECT1`, `E1-P1C3-COMMIT1` |
| P1-C4 | `E1-P1C4-IMPACT1`, `E1-P1C4-SRC1`, `E1-P1C4-BUILD1`, `E1-P1C4-TEST1`, `E1-P1C4-REVIEW1`, `E1-P1C4-DETECT1`, `E1-P1C4-COMMIT1` |
| P1-C5 | `E1-P1C5-IMPACT1`, `E1-P1C5-SRC1`, `E1-P1C5-BUILD1`, `E1-P1C5-TEST1`, `E1-P1C5-REVIEW1`, `E1-P1C5-DETECT1`, `E1-P1C5-COMMIT1` |
| P1-C6 | `E1-P1C6-IMPACT1`, `E1-P1C6-SRC1`, `E1-P1C6-BUILD1`, `E1-P1C6-TEST1`, `E1-P1C6-REVIEW1`, `E1-P1C6-DETECT1`, `E1-P1C6-COMMIT1` |
| P1-D | `E1-P1D-IMPACT1`, `E1-P1D-SRC1`, `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-REVIEW1`, `E1-P1D-DETECT1`, `E1-P1D-COMMIT1` |
| P1-D1 | `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`, `E1-P1D1-REVIEW1`, `E1-P1D1-DETECT1`, `E1-P1D1-COMMIT1` |
| P1-D2 | `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`, `E1-P1D2-DETECT1`, `E1-P1D2-COMMIT1` |
| P1-E | `E1-P1E-IMPACT1`, `E1-P1E-SRC1`, `E1-P1E-BUILD1`, `E1-P1E-TEST1`, `E1-P1E-REVIEW1`, `E1-P1E-DETECT1`, `E1-P1E-COMMIT1` |
| P1-E1 | `E1-P1E1-IMPACT1`, `E1-P1E1-SRC1`, `E1-P1E1-BUILD1`, `E1-P1E1-PLAY1`, `E1-P1E1-REVIEW1`, `E1-P1E1-DETECT1`, `E1-P1E1-COMMIT1` |
| P1-E2 | `E1-P1E2-BUILD1`, `E1-P1E2-S0BASE1`, `E1-P1E2-S1BASE1`, `E1-P1E2-S2BASE1`, `E1-P1E2-S3BASE1`, `E1-P1E2-S4BASE1`, `E1-P1E2-S5BASE1`, `E1-P1E2-S6BASE1`, `E1-P1E2-S7BASE1`, `E1-P1E2-S8BASE1`, `E1-P1E2-S9BASE1`, `E1-P1E2-S10BASE1`, `E1-P1E2-S11BASE1`, `E1-P1E2-MATRIX1`, `E1-P1E2-PLAY1`, `E1-P1E2-REVIEW1`, `E1-P1E2-DETECT1`, `E1-P1E2-COMMIT1` |
| P1-F | `E1-P1F-IMPACT1`, `E1-P1F-SRC1`, `E1-P1F-BUILD1`, `E1-P1F-TEST1`, `E1-P1F-REVIEW1`, `E1-P1F-DETECT1`, `E1-P1F-COMMIT1` |
| P1-F1 | `E1-P1F1-IMPACT1`, `E1-P1F1-SRC1`, `E1-P1F1-BUILD1`, `E1-P1F1-TEST1`, `E1-P1F1-REVIEW1`, `E1-P1F1-DETECT1`, `E1-P1F1-COMMIT1` |
| P1-F2 | `E1-P1F2-IMPACT1`, `E1-P1F2-SRC1`, `E1-P1F2-BUILD1`, `E1-P1F2-TEST1`, `E1-P1F2-REVIEW1`, `E1-P1F2-DETECT1`, `E1-P1F2-COMMIT1` |
| P1-F3 | `E1-P1F3-IMPACT1`, `E1-P1F3-SRC1`, `E1-P1F3-BUILD1`, `E1-P1F3-TEST1`, `E1-P1F3-REVIEW1`, `E1-P1F3-DETECT1`, `E1-P1F3-COMMIT1` |
| P1-F4 | `E1-P1F4-IMPACT1`, `E1-P1F4-SRC1`, `E1-P1F4-BUILD1`, `E1-P1F4-TEST1`, `E1-P1F4-REVIEW1`, `E1-P1F4-DETECT1`, `E1-P1F4-COMMIT1` |
| P1-F5 | `E1-P1F5-IMPACT1`, `E1-P1F5-SRC1`, `E1-P1F5-BUILD1`, `E1-P1F5-TEST1`, `E1-P1F5-REVIEW1`, `E1-P1F5-DETECT1`, `E1-P1F5-COMMIT1` |
| P1-F6 | `E1-P1F6-BUILD1`, `E1-P1F6-FAULT1`, `E1-P1F6-RECOVERY1`, `E1-P1F6-REVIEW1`, `E1-P1F6-DETECT1`, `E1-P1F6-COMMIT1` |
| P1-G | `E1-P1G-PREBASE1`, `E1-P1G-CANDIDATE1`, `E1-P1G-IMPACT1`, `E1-P1G-SRC1`, `E1-P1G-CUTOVER1`, `E1-P1G-BUILD1`, `E1-P1G-RUNTIME1`, `E1-P1G-PLAY1`, `E1-P1G-ROLLBACK1`, `E1-P1G-REVIEW1`, `E1-P1G-DETECT1`, `E1-P1G-COMMIT1` |

## E2 - Closure Evidence

Matching plan item(s): `Pn-A`, `Pn-B`, `Pn-C`

| Closure item | Reserved evidence | Status |
|--------------|-------------------|--------|
| Pn-A | `E2-PNA-SUP1` | pending |
| Pn-B | `E2-PNB-CLEAN1`, `E2-PNB-REVIEW1` | pending |
| Pn-C | `E2-PNC-BUILD1`, `E2-PNC-RUNTIME1`, `E2-PNC-PLAY1`, `E2-PNC-REGEN1`, `E2-PNC-DETECT1`, `E2-PNC-LEDGER1`, `E2-PNC-NEXTSTATUS1`, `E2-PNC-COMMIT1`, `E2-PNC-HANDOFF1` | pending |

## Closure Evidence

No implementation or closure evidence exists yet. Populate E2 only after every local P1 slice is accepted/committed or explicitly blocked.
