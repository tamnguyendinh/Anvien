# Anvien Graph Identity Contract and Strict Construction Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`

## Evidence Rules

- The original problem report is the problem origin and bounded oracle source.
- Architecture recommendations inside that report are DRAFT and do not become implementation evidence without current source proof and plan acceptance.
- The causal synthesis and bounded Supervisor PASS verify the finding; they are not the originating report or an implementation design authority.
- Current source and runtime evidence determine implementation ownership.
- A pending evidence ID is a declared target, not proof.
- Every implementation slice closes only when all exact evidence IDs named by its plan Acceptance are recorded.
- Counts and timings belong in the benchmark ledger; commands, source facts, and verdicts belong here.

### Evidence ID Naming

Use `E<phase>-<item>-<kind><n>`. Each ID has one meaning and is not reused.

## E0 - P0 Evidence

Matching plan item: `P0-A`

- `E0-P0A-REPORT1` — recorded: `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` establishes the bounded identity baseline: four `time`/`now` source facts survive ScopeIR while two graph nodes remain. Only the finding and `2/4 -> 4/4` target are used here.
- `E0-P0A-VERIFY1` — recorded: `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md` classifies range/scope-free graph identity plus duplicate-node replacement as the bounded first graph divergence; `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md` passes that bounded investigation only.
- `E0-P0A-GRAPH1` — pending P0 execution: fresh `anvien analyze --force` and status evidence at the implementation HEAD.
- `E0-P0A-SOURCE1` — retained source fact; refresh required before edit: current graph definition identity uses file/name/arity-oriented inputs and lacks sufficient lexical-owner/meaning/source distinction for the bounded case.
- `E0-P0A-SOURCE2` — retained source fact; refresh required before edit: current shared range input does not yet establish the complete full/selection-range position contract required by P1-B.
- `E0-P0A-SOURCE3` — retained source fact; refresh required before edit: the current definition index can skip a later fact whose identity is already present.
- `E0-P0A-SOURCE4` — retained source fact; refresh required before edit: current graph duplicate-node handling can replace the earlier payload for the same ID.
- `E0-P0A-SOURCE5` — pending P0 execution: current end-to-end provider/ScopeIR/identity/graph flow and whether any orchestration change is actually required.
- `E0-P0A-FD1` — retained candidate count, refresh required before edit: `internal/resolution/indexes.go` has 46 related files.
- `E0-P0A-FD2` — retained candidate count, refresh required before edit: `internal/scopeir/range.go` has 227 related files.
- `E0-P0A-FD3` — retained candidate count, refresh required before edit: `internal/scopeir/definition_index.go` has 225 related files.
- `E0-P0A-FD4` — retained candidate count, refresh required before edit: `internal/graph/types.go` has 238 related files.
- `E0-P0A-FD5` — retained candidate count, refresh required before edit: `internal/resolution/emit.go` has 42 related files.
- `E0-P0A-IMPACT1` — pending: exact current upstream impact for every editable file/symbol selected by P1-B/P1-C/P1-D.
- `E0-P0A-BOUNDARY1` — recorded: `E:\cheapapp.org` is an in-place validation target; target source/worktree is preserve-only and only normal target-repository output is allowed.
- `E0-P0A-STATUS1` — recorded: current actual-status ledger marks P0 incomplete, narrows scope to identity, and keeps every production slice blocked.
- `E0-P0A-REVIEW1` — pending: unconditional Supervisor PASS for the completed P0 source/impact/status boundary.

## E1 - P1 Evidence

Matching plan items: `P1-A`, `P1-B`, `P1-C`, `P1-D`, `P1-E`

| Slice | Exact evidence ID | Required proof | Status |
|-------|-------------------|----------------|--------|
| P1-A | `E1-P1A-SOURCE1` | source/report-to-contract trace with DRAFT recommendations excluded | pending |
| P1-A | `E1-P1A-CONTRACT1` | accepted graph-accuracy identity contract and ownership boundary | pending |
| P1-A | `E1-P1A-REVIEW1` | unconditional Supervisor PASS for the contract slice | pending |
| P1-A | `E1-P1A-COMMIT1` | isolated documentation commit and known worktree | pending |
| P1-B | `E1-P1B-IMPACT1` | fresh file-detail and upstream impact for exact range/input owners | pending |
| P1-B | `E1-P1B-SOURCE1` | production range/selection/owner/meaning input change | pending |
| P1-B | `E1-P1B-BUILD1` | full build after production code and focused tests | pending |
| P1-B | `E1-P1B-TEST1` | position/owner/meaning behavior matrix | pending |
| P1-B | `E1-P1B-RUNTIME1` | nearest built provider/ScopeIR/analyzer boundary | pending |
| P1-B | `E1-P1B-REVIEW1` | unconditional Supervisor PASS | pending |
| P1-B | `E1-P1B-DETECT1` | Anvien detect-changes before commit | pending |
| P1-B | `E1-P1B-COMMIT1` | isolated slice commit | pending |
| P1-C | `E1-P1C-IMPACT1` | fresh impact for exact identity/occurrence owners | pending |
| P1-C | `E1-P1C-SOURCE1` | deterministic declaration/symbol mapping and occurrence implementation | pending |
| P1-C | `E1-P1C-BUILD1` | full build | pending |
| P1-C | `E1-P1C-TEST1` | same-name, meaning, ordering, and merge-evidence tests | pending |
| P1-C | `E1-P1C-ORACLE1` | occurrence conservation and bounded same-name oracle at the owned boundary | pending |
| P1-C | `E1-P1C-REVIEW1` | unconditional Supervisor PASS | pending |
| P1-C | `E1-P1C-DETECT1` | Anvien detect-changes | pending |
| P1-C | `E1-P1C-COMMIT1` | isolated slice commit | pending |
| P1-D | `E1-P1D-IMPACT1` | exact collision path, callers, and legitimate-operation impact | pending |
| P1-D | `E1-P1D-SOURCE1` | evidence-scoped collision/loss correction | pending |
| P1-D | `E1-P1D-BUILD1` | full build | pending |
| P1-D | `E1-P1D-TEST1` | conflict, enrichment, occurrence, and endpoint tests | pending |
| P1-D | `E1-P1D-COLLISION1` | zero silent collision/skip/replacement proof | pending |
| P1-D | `E1-P1D-REVIEW1` | unconditional Supervisor PASS | pending |
| P1-D | `E1-P1D-DETECT1` | Anvien detect-changes | pending |
| P1-D | `E1-P1D-COMMIT1` | isolated slice commit | pending |
| P1-E | `E1-P1E-BUILD1` | final full build before integration validation | pending |
| P1-E | `E1-P1E-DETERMINISM1` | matched repeated-run canonical identity comparison | pending |
| P1-E | `E1-P1E-INTEGRITY1` | occurrence conservation, collision, range, and endpoint integrity result | pending |
| P1-E | `E1-P1E-TARGET1` | bounded target `time`/`now` result `4/4` | pending |
| P1-E | `E1-P1E-BOUNDARY1` | target source/worktree preservation proof | pending |
| P1-E | `E1-P1E-REVIEW1` | unconditional Supervisor PASS | pending |
| P1-E | `E1-P1E-DETECT1` | Anvien detect-changes | pending |
| P1-E | `E1-P1E-COMMIT1` | isolated validation-slice commit | pending |

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E2-PNA-SUPERVISOR1` | independent acceptance of complete Child 01 source/diff/runtime/evidence/benchmark | pending |
| `E2-PNB-CLEANUP1` | dead-work inventory, removal, and Supervisor cleanup PASS | pending |
| `E2-PNC-DETECT1` | final detect-changes result for the accepted Child scope | pending |
| `E2-PNC-HANDOFF1` | exact accepted identity/range facts and evidence used to refresh Child 02 | pending |
| `E2-PNC-COMMIT1` | final closure commit and known worktree state | pending |
