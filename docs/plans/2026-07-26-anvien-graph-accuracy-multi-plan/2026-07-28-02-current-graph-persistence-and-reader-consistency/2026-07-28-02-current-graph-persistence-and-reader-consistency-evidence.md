# Anvien Current Graph Persistence and Reader Consistency Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`

## Evidence Rules

- The original problem report is the problem origin; use its bounded findings and acceptance targets, not its DRAFT architecture proposals.
- The causal synthesis and bounded Supervisor PASS verify selected persistence/representation observations but do not define the implementation.
- Child 01 handoff and current source/runtime are the authority for corrected fields and their actual consumers.
- A reader enters the affected denominator only through exact source evidence recorded by P2-A.
- A pending evidence ID is a declared target, not proof.
- Every slice requires every exact ID named in its Acceptance; broad phase references are insufficient.
- Measured counts/timings belong in benchmark; source facts, commands, and verdicts belong here.

### Evidence ID Naming

Use `E<phase>-<item>-<kind><n>`. Do not reuse an ID for another fact.

## E0 - P0 Evidence

Matching plan item: `P0-A`

- `E0-P0A-REPORT1` — recorded: `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` is the campaign problem origin, but it does not establish a Child 02 persistence requirement. Child 02 takes its bounded representation evidence from `E0-P0A-VERIFY1` and its corrected-field contract from the accepted Child 01 handoff; the report's proposed storage design is not authority.
- `E0-P0A-VERIFY1` — recorded: `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md` observes one-to-one selected-fact cardinality with bounded representation changes across raw/Cypher/file-detail paths; `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md` passes only that bounded record.
- `E0-P0A-HANDOFF1` — pending: accepted Child 01 corrected identity/range fields, source owners, evidence, and commit.
- `E0-P0A-GRAPH1` — pending P0 execution: fresh Anvien analyze/status at Child 02 implementation HEAD.
- `E0-P0A-SOURCE1` — pending P0 execution: exact current Graph JSON/Ladybug encode/load/query flow for the accepted corrected fields.
- `E0-P0A-SOURCE2` — pending P0 execution: current repeated normal analyze and affected-reader flow, including actual artifact paths and failure behavior.
- `E0-P0A-FD1` — retained candidate count, refresh required before edit: `internal/lbugload/csv.go` has 19 related files.
- `E0-P0A-FD2` — retained candidate count, refresh required before edit: `internal/analyze/analyze.go` has 182 related files.
- `E0-P0A-FD3` — retained candidate count, refresh required before edit: `internal/httpapi/graph.go` has 22 related files.
- `E0-P0A-FD4` — retained candidate count, refresh required before edit: `anvien-web/src/services/backend-client.ts` has 24 related files.
- `E0-P0A-IMPACT1` — pending: exact current impact for persistence and reader owners selected by P2-A/P2-B/P2-C/P2-D.
- `E0-P0A-BOUNDARY1` — recorded: target source/worktree is preserve-only; validation uses normal target-repository output and stores plan/evidence artifacts in Anvien.
- `E0-P0A-STATUS1` — recorded: actual status removes campaign-wide reader assumptions, marks P0 incomplete, and leaves reader candidates inspect-only.
- `E0-P0A-REVIEW1` — pending: unconditional Supervisor PASS for Child 02 P0 scope and dependency decision.

## E2 - P2 Evidence

Matching plan items: `P2-A`, `P2-B`, `P2-C`, `P2-D`, `P2-E`

| Slice | Exact evidence ID | Required proof | Status |
|-------|-------------------|----------------|--------|
| P2-A | `E2-P2A-IMPACT1` | fresh file-detail/impact for every candidate persistence/reader owner | pending |
| P2-A | `E2-P2A-SOURCE1` | corrected-field source traces through real persistence/read flows | pending |
| P2-A | `E2-P2A-INVENTORY1` | exact symbol/path, field, impact, touch mode, and evidence per affected row | pending |
| P2-A | `E2-P2A-MATRIX1` | unique affected denominator with zero duplicate/unassigned rows and explicit unaffected exclusions | pending |
| P2-A | `E2-P2A-REVIEW1` | unconditional Supervisor PASS for inventory | pending |
| P2-A | `E2-P2A-COMMIT1` | isolated documentation/source-audit commit | pending |
| P2-B | `E2-P2B-IMPACT1` | fresh impact for exact affected Graph JSON/Ladybug owners | pending |
| P2-B | `E2-P2B-SOURCE1` | production persistence correction, if source proves needed | pending |
| P2-B | `E2-P2B-BUILD1` | full build | pending |
| P2-B | `E2-P2B-TEST1` | corrected-field round-trip/drop/closure tests | pending |
| P2-B | `E2-P2B-PARITY1` | Graph JSON/Ladybug corrected-field parity and zero dropped records | pending |
| P2-B | `E2-P2B-REVIEW1` | unconditional Supervisor PASS | pending |
| P2-B | `E2-P2B-DETECT1` | Anvien detect-changes | pending |
| P2-B | `E2-P2B-COMMIT1` | isolated slice commit | pending |
| P2-C | `E2-P2C-IMPACT1` | exact impact for P2-A affected readers | pending |
| P2-C | `E2-P2C-SOURCE1` | affected-reader production corrections only | pending |
| P2-C | `E2-P2C-BUILD1` | full build | pending |
| P2-C | `E2-P2C-TEST1` | exact affected-reader behavior and sibling regression tests | pending |
| P2-C | `E2-P2C-RUNTIME1` | nearest real boundary per affected reader, including UI evidence only when affected | pending |
| P2-C | `E2-P2C-REVIEW1` | unconditional Supervisor PASS | pending |
| P2-C | `E2-P2C-DETECT1` | Anvien detect-changes | pending |
| P2-C | `E2-P2C-COMMIT1` | isolated slice commit | pending |
| P2-D | `E2-P2D-IMPACT1` | current repeated-analyze/read owner impact before any edit | pending |
| P2-D | `E2-P2D-SOURCE1` | current behavior trace and any evidence-required production correction | pending |
| P2-D | `E2-P2D-BUILD1` | full build | pending |
| P2-D | `E2-P2D-TEST1` | unchanged/changed input plus failure/subsequent-success matrix | pending |
| P2-D | `E2-P2D-REPEAT1` | normal built repeated-analyze results and matching affected reads | pending |
| P2-D | `E2-P2D-REVIEW1` | unconditional Supervisor PASS | pending |
| P2-D | `E2-P2D-DETECT1` | Anvien detect-changes or documented no-production-change result | pending |
| P2-D | `E2-P2D-COMMIT1` | isolated implementation/validation commit | pending |
| P2-E | `E2-P2E-BUILD1` | final full build before acceptance | pending |
| P2-E | `E2-P2E-PARITY1` | independent corrected-field Graph JSON/Ladybug parity | pending |
| P2-E | `E2-P2E-READERS1` | affected readers passed/total equals exact P2-A denominator | pending |
| P2-E | `E2-P2E-REPEAT1` | independent repeated normal analyze and clear-failure acceptance | pending |
| P2-E | `E2-P2E-REVIEW1` | unconditional Supervisor PASS | pending |
| P2-E | `E2-P2E-DETECT1` | final Child 02 detect-changes result | pending |
| P2-E | `E2-P2E-COMMIT1` | isolated validation/ledger commit | pending |

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E3-PNA-SUPERVISOR1` | independent acceptance of complete Child 02 source/diff/persistence/read/runtime evidence | pending |
| `E3-PNB-CLEANUP1` | dead-work inventory, removal, and Supervisor cleanup PASS | pending |
| `E3-PNC-DETECT1` | final detect-changes result | pending |
| `E3-PNC-HANDOFF1` | exact accepted persistence fields, affected-reader guarantees, evidence, and Child 03 refresh | pending |
| `E3-PNC-COMMIT1` | final closure commit and known worktree state | pending |
