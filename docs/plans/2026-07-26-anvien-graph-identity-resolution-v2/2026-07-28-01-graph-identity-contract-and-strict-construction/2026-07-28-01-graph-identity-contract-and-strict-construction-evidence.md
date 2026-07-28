# Graph Identity Contract and Strict Graph Construction Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Source phase: legacy `P1`

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

- `E0-P0A-PLAN1`: This standard child set contains plan, evidence, benchmark, and actual-status files with slug `2026-07-28-01-graph-identity-contract-and-strict-construction`; local P1 has 11 slices mapped from legacy P1.
- `E0-P0A-SOURCE1`: Legacy source phase P1 spans source lines 313-898, remains under frozen SHA-256 `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB`, and maps to 11 unique local slice IDs.
- `E0-P0A-GRAPH1`: Authoring uses the fresh Anvien graph recorded by the parent authoring plan at commit `7b6c8e57`: 1,511 files, 676 parsed code files, 84,617 nodes, and 123,457 relationships.
- `E0-P0A-FDPLAN1`: Post-authoring refresh at commit `c444e8c4` indexed 1,518 files, 84,702 nodes, and 123,546 relationships. Fresh file-detail classifies this plan as parsed Markdown/docs, low risk, with exactly one related file: the roadmap's inbound `IMPORTS` link; unresolved count is zero.
- `E0-P0A-STATUS1`: No production path changed between legacy baseline commit `1932359b` and roadmap commit `c444e8c4`; the inherited current behavior remains: Graph IDs omit scope/range, duplicate nodes can overwrite, and occurrence/source-site conservation is not guaranteed.
- `E0-P0A-DEPENDENCY1`: No upstream implementation child applies. Local P1-A must ratify architecture authority before later production slices, and production implementation still requires explicit owner direction. Local P1 remains blocked by campaign cutover/owner authorization, not by a missing upstream child.
- `E0-P0A-BOUNDARY1`: `E:\cheapapp.org` is an in-place validation subject only for explicitly opened target slices; no target source, report, probe, fixture, or temp artifact was used to author this child.
- `E0-P0A-SCANNER1`: The exact eight scanner omissions remain quarantined, wrong-but-out-of-scope, with zero additional omissions required.
- `E0-P0A-FD1`: inherited source P0 file-detail for `internal/graph/types.go` reports 238 related files/items; central node/relationship storage and indexing; high; Graph.AddNode/Graph.init CRITICAL. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD2`: inherited source P0 file-detail for `internal/scopeir/facts.go` reports 231 related files/items; mixed ScopeIR fact contracts; medium; dedicated owner extraction required. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD3`: inherited source P0 file-detail for `internal/scopeir/range.go` reports 227 related files/items; shared position/range contract; medium; all providers depend on encoding. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD4`: inherited source P0 file-detail for `internal/scopeir/definition_index.go` reports 225 related files/items; definition identity index; medium; currently hides duplicates. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD5`: inherited source P0 file-detail for `internal/resolution/indexes.go` reports 46 related files/items; identity construction and resolution indexes; high; several CRITICAL symbols. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD6`: inherited source P0 file-detail for `internal/resolution/emit.go` reports 42 related files/items; graph projection; high identity parity scope. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.
- `E0-P0A-FD7`: inherited source P0 file-detail for `internal/lbugload/csv.go` reports 19 related files/items; persistence projection boundary; high duplicate/closure risk. Production paths are unchanged since the legacy baseline, but this evidence must be refreshed immediately before any implementation edit.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`, `P1-B`, `P1-C0`, `P1-C0A`, `P1-C0B`, `P1-C`, `P1-D`, `P1-D1`, `P1-D2`, `P1-D3`, `P1-E`

Source phase: legacy `P1`; every local evidence ID preserves its source-phase provenance through this child slug.

| Reserved evidence | Required proof | Status |
|-------------------|----------------|--------|
| `E1-P1A-CONTRACT1` | owner-accepted identity/ownership decision record | pending |
| `E1-P1A-REVIEW1` | independent architecture/Supervisor verdict | pending |
| `E1-P1A-COMMIT1` | isolated contract-only commit and worktree state | pending |
| `E1-P1B-IMPACT1` | fresh file-detail/impact for identity/range owners | pending |
| `E1-P1B-SRC1` | production type/ownership diff | pending |
| `E1-P1B-BUILD1` | full build after production code and tests | pending |
| `E1-P1B-TEST1` | position/serialization/type-safety behavior tests | pending |
| `E1-P1B-REVIEW1` | Supervisor PASS | pending |
| `E1-P1B-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1B-COMMIT1` | isolated slice commit and worktree state | pending |
| `E1-P1C-IMPACT1` | fresh identity mapping impact | pending |
| `E1-P1C-SRC1` | canonical tuple/merge/stability implementation | pending |
| `E1-P1C-BUILD1` | full build | pending |
| `E1-P1C-TEST1` | identity, stability, overload, collision, order tests | pending |
| `E1-P1C-REVIEW1` | Supervisor PASS | pending |
| `E1-P1C-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1C-COMMIT1` | isolated slice commit and worktree state | pending |
| `E1-P1D-IMPACT1` | graph mutation/producer impact | pending |
| `E1-P1D-MAP1` | complete producer-to-operation inventory | pending |
| `E1-P1D-SRC1` | strict mutation/decode/validation diff | pending |
| `E1-P1D-BUILD1` | full build | pending |
| `E1-P1D-TEST1` | duplicate, closure, enrich, replace failure matrix | pending |
| `E1-P1D-REVIEW1` | Supervisor PASS | pending |
| `E1-P1D-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1D-COMMIT1` | isolated slice commit and worktree state | pending |
| `E1-P1E-IMPACT1` | fresh emitter/analyze-hook file-detail and upstream impact | pending |
| `E1-P1E-SRC1` | shadow v2 emitter/comparator implementation | pending |
| `E1-P1E-SHADOW1` | canonical v1/v2 shadow signature comparison | pending |
| `E1-P1E-BUILD1` | full build | pending |
| `E1-P1E-BENCH1` | five-run graph expansion/size/RSS measurements | pending |
| `E1-P1E-TARGET1` | in-memory target v2 shadow target-site run before cutover | pending |
| `E1-P1E-ORACLE1` | independent `time`/`now` 4/4 source/TS-oracle manifest | pending |
| `E1-P1E-BOUNDARY1` | target pre/post hash/worktree/graph-path/artifact boundary | pending |
| `E1-P1E-REVIEW1` | Supervisor PASS | pending |
| `E1-P1E-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1E-COMMIT1` | isolated slice commit and worktree state | pending |
| `E1-P1C0-IMPACT1` | declaration occurrence-index impact | pending |
| `E1-P1C0-SRC1` | lossless v2 occurrence index and v1 compatibility adapter | pending |
| `E1-P1C0-BUILD1` | full build | pending |
| `E1-P1C0-TEST1` | declaration conservation/collision tests | pending |
| `E1-P1C0-REVIEW1` | Supervisor PASS | pending |
| `E1-P1C0-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1C0-COMMIT1` | isolated slice commit | pending |
| `E1-P1C0A-IMPACT1` | RelationshipID/aggregation impact | pending |
| `E1-P1C0A-SRC1` | RelationshipID and source-site aggregation implementation | pending |
| `E1-P1C0A-BUILD1` | full build | pending |
| `E1-P1C0A-TEST1` | same-endpoint/multi-source-site conservation tests | pending |
| `E1-P1C0A-REVIEW1` | Supervisor PASS | pending |
| `E1-P1C0A-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1C0A-COMMIT1` | isolated slice commit | pending |
| `E1-P1C0B-IMPACT1` | decode/closure impact | pending |
| `E1-P1C0B-SRC1` | fail-closed graph decode/validation implementation | pending |
| `E1-P1C0B-BUILD1` | full build | pending |
| `E1-P1C0B-TEST1` | duplicate/conflict/endpoint/subset failure tests | pending |
| `E1-P1C0B-REVIEW1` | Supervisor PASS | pending |
| `E1-P1C0B-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1C0B-COMMIT1` | isolated slice commit | pending |
| `E1-P1D1-IMPACT1` | core producer migration impact | pending |
| `E1-P1D1-SRC1` | core producer explicit-operation diff | pending |
| `E1-P1D1-BUILD1` | full build | pending |
| `E1-P1D1-TEST1` | core producer behavior tests | pending |
| `E1-P1D1-REVIEW1` | Supervisor PASS | pending |
| `E1-P1D1-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1D1-COMMIT1` | isolated slice commit | pending |
| `E1-P1D2-IMPACT1` | resolution/projection producer migration impact | pending |
| `E1-P1D2-SRC1` | resolution/projection explicit-operation diff | pending |
| `E1-P1D2-BUILD1` | full build | pending |
| `E1-P1D2-TEST1` | resolution/projection behavior tests | pending |
| `E1-P1D2-REVIEW1` | Supervisor PASS | pending |
| `E1-P1D2-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1D2-COMMIT1` | isolated slice commit | pending |
| `E1-P1D3-IMPACT1` | ancillary producer migration impact | pending |
| `E1-P1D3-SRC1` | ancillary explicit-operation diff | pending |
| `E1-P1D3-BUILD1` | full build | pending |
| `E1-P1D3-TEST1` | ancillary behavior tests | pending |
| `E1-P1D3-REVIEW1` | Supervisor PASS | pending |
| `E1-P1D3-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1D3-COMMIT1` | isolated slice commit | pending |

### Preserved source traceability binding

| Local slice | Required exact local evidence IDs |
|-------------|-----------------------------------|
| P1-A | `E1-P1A-CONTRACT1`, `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1` |
| P1-B | `E1-P1B-IMPACT1`, `E1-P1B-SRC1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-REVIEW1`, `E1-P1B-DETECT1`, `E1-P1B-COMMIT1` |
| P1-C0 | `E1-P1C0-IMPACT1`, `E1-P1C0-SRC1`, `E1-P1C0-BUILD1`, `E1-P1C0-TEST1`, `E1-P1C0-REVIEW1`, `E1-P1C0-DETECT1`, `E1-P1C0-COMMIT1` |
| P1-C0A | `E1-P1C0A-IMPACT1`, `E1-P1C0A-SRC1`, `E1-P1C0A-BUILD1`, `E1-P1C0A-TEST1`, `E1-P1C0A-REVIEW1`, `E1-P1C0A-DETECT1`, `E1-P1C0A-COMMIT1` |
| P1-C0B | `E1-P1C0B-IMPACT1`, `E1-P1C0B-SRC1`, `E1-P1C0B-BUILD1`, `E1-P1C0B-TEST1`, `E1-P1C0B-REVIEW1`, `E1-P1C0B-DETECT1`, `E1-P1C0B-COMMIT1` |
| P1-C | `E1-P1C-IMPACT1`, `E1-P1C-SRC1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` |
| P1-D | `E1-P1D-IMPACT1`, `E1-P1D-MAP1`, `E1-P1D-SRC1`, `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-REVIEW1`, `E1-P1D-DETECT1`, `E1-P1D-COMMIT1` |
| P1-D1 | `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`, `E1-P1D1-REVIEW1`, `E1-P1D1-DETECT1`, `E1-P1D1-COMMIT1` |
| P1-D2 | `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`, `E1-P1D2-DETECT1`, `E1-P1D2-COMMIT1` |
| P1-D3 | `E1-P1D3-IMPACT1`, `E1-P1D3-SRC1`, `E1-P1D3-BUILD1`, `E1-P1D3-TEST1`, `E1-P1D3-REVIEW1`, `E1-P1D3-DETECT1`, `E1-P1D3-COMMIT1` |
| P1-E | `E1-P1E-IMPACT1`, `E1-P1E-SRC1`, `E1-P1E-SHADOW1`, `E1-P1E-TARGET1`, `E1-P1E-ORACLE1`, `E1-P1E-BOUNDARY1`, `E1-P1E-BUILD1`, `E1-P1E-BENCH1`, `E1-P1E-REVIEW1`, `E1-P1E-DETECT1`, `E1-P1E-COMMIT1` |

## E2 - Closure Evidence

Matching plan item(s): `Pn-A`, `Pn-B`, `Pn-C`

| Closure item | Reserved evidence | Status |
|--------------|-------------------|--------|
| Pn-A | `E2-PNA-SUP1` | pending |
| Pn-B | `E2-PNB-CLEAN1` | pending |
| Pn-C | `E2-PNC-BUILD1`, `E2-PNC-RUNTIME1`, `E2-PNC-DETECT1`, `E2-PNC-COMMIT1`, `E2-PNC-HANDOFF1` | pending |

## Closure Evidence

No implementation or closure evidence exists yet. Populate E2 only after every local P1 slice is accepted/committed or explicitly blocked.
