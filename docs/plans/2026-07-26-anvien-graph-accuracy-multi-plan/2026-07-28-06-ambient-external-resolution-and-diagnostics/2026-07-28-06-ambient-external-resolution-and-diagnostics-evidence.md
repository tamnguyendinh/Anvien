# Anvien Ambient And External Resolution Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`

## Evidence Rules

- The originating problem report is the problem origin. Its declaration architecture is not implementation authority.
- The causal synthesis and final Supervisor report are bounded verification only; they do not establish broader declaration sources, runtime mechanism, public APIs, or remediation design.
- P6-A must use current source, config, packaging/runtime, consumer impact, and oracle evidence to select the implementation contract.
- A pending evidence ID is a required target, not proof.
- Every production slice requires impact, source, build, behavior, boundary/parity where applicable, Supervisor, detect-changes, and commit evidence.
- Record exact source sites and outcomes for `Promise`, `Math.max`, and `Math.min`; aggregate health counts cannot close the target gate.
- Long measurements belong in the benchmark ledger.

### Evidence ID Naming

Evidence IDs use `E<phase>-<item>-<kind><n>` and remain stable across all four Child 06 files. `E0` maps to P0; `E6` maps to P6 and closure.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-RULE1`: read `E:\Anvien\AGENTS.md` in full. It requires planner use, fresh graph evidence, file-detail/impact before implementation edits, code before tests, full build, truthful boundary validation, Supervisor acceptance, detect-changes, and per-slice commits.
- `E0-P0A-SKILL1`: read `.agents/skills/planner/SKILL.md` and all four planner templates in full before rewriting this four-file set.
- `E0-P0A-ORIGIN1`: read `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` in full. It is the problem origin and records the three target acceptance sites, but expressly labels its architecture DRAFT; its authority-mechanism proposal and broader source list are not accepted design.
- `E0-P0A-VERIFY1`: read `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md` in full. C5 confirms that the resolver workspace indexes target ScopeIR definitions only and lacks TypeScript ambient/lib declarations; TypeScript reports zero diagnostics at the three selected sites while Anvien records gaps. Broader policy remains unresolved.
- `E0-P0A-VERIFY2`: read `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md` in full. Verdict PASS applies only to the bounded investigation and explicitly leaves declaration policy/remediation/global semantics unresolved.
- `E0-P0A-GRAPH1`: after graph refresh in the shared workspace, `anvien status` reported indexed/current commit `238aec06d28286acc0fca0f4e6a69f9eb4ff6a49` and up-to-date status. Current Child 06 file-detail checks report graph `stale=false`, `changedSinceAnalyze=false`, analyzed at `2026-08-09T19:19:54Z`.
- `E0-P0A-SRC1`: full source read of `internal/resolution/indexes.go`. `buildWorkspace` accepts `[]scopeir.ScopeIR` and builds file, scope, definition, binding, member, import, and heritage indexes solely from those repository inputs. No declaration-universe or TypeScript standard-library input is present.
- `E0-P0A-SRC2`: full source read of `internal/resolution/resolve.go`. Resolution produces graph references or generic unresolved diagnostics. There is no external declaration target or explicit external-capability outcome contract.
- `E0-P0A-SRC3`: in current `resolve.go`, `resolveTypeAnnotation` ignores a small primitive list and otherwise calls repository workspace lookup. Member calls depend on receiver type/owner/member lookup in current indexes. This cannot resolve `Promise` or `Math` members without an authority input.
- `E0-P0A-SRC4`: full source read of `internal/graphhealth/diagnostics.go`. `classifyDiagnostic` infers from `TargetText` using Go builtin/composite/test/standard-library and qualifier maps; it does not consume a resolver-stage external outcome.
- `E0-P0A-FD1`: `anvien file-detail internal/resolution/indexes.go --repo E:\Anvien --json` reported 46 related files, 192 symbols, 164 inbound, 93 outbound, 369 local relationships, 29 linked flows, 23 linked tests, and high file risk.
- `E0-P0A-FD2`: `anvien file-detail internal/resolution/resolve.go --repo E:\Anvien --json` reported 50 related files, 40 symbols, 77 inbound, 118 outbound, 18 local relationships, 21 linked flows, 26 linked tests, and high file risk.
- `E0-P0A-FD3`: `anvien file-detail internal/graphhealth/diagnostics.go --repo E:\Anvien --json` reported 29 related files, 47 symbols, 36 inbound, 42 outbound, 22 local relationships, 8 linked flows, 14 linked tests, and high file risk.
- `E0-P0A-IMPACT1`: upstream impact for `buildWorkspace` is CRITICAL: 8 impacted symbols, 6 affected files, 5 modules, 28 processes.
- `E0-P0A-IMPACT2`: upstream impact for `resolveCall` is CRITICAL: 6 impacted symbols, 4 affected files, 4 modules, 35 processes.
- `E0-P0A-IMPACT3`: upstream impact for `resolveTypeAnnotation` is CRITICAL: 6 impacted symbols, 4 affected files, 4 modules, 35 processes.
- `E0-P0A-IMPACT4`: upstream impact for `classifyDiagnostic` is HIGH: 7 impacted symbols, 2 affected files, 3 modules, 3 processes.
- `E0-P0A-DEPEND1`: Child 06 consumes accepted Child 05 repository terminal/unresolved results and proofs. P6-A remains blocked until that predecessor's four slices, Supervisor result, ledgers, and commits close.
- `E0-P0A-SCOPE1`: exact supported TypeScript config inputs, feasible declaration-authority mechanisms, project/package scope, packaging/runtime constraints, and external side effects are unresolved P6-A decisions; no current evidence permits choosing them from document terminology.
- `E0-P0A-SCOPE2`: affected persistence/readers for new external targets/outcomes must come from the current Child 02 reader-impact inventory and fresh P6 consumer impact; no fixed transport, option, or all-reader matrix is authority.
- `E0-P0A-TARGET1`: accepted bounded denominator is exactly three sites: `Promise`, `Math.max`, and `Math.min`. Baseline is `0/3` correct external/capability outcomes; all three are currently false in-repository unresolved/analyzer-gap cases.
- `E0-P0A-BOUNDARY1`: accepted target is `E:\cheapapp.org`, HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, graph at `E:\cheapapp.org\.anvien\graph.json`, graph hash `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`, 84,807 nodes, and 114,125 relationships. Target source is preserve-only.
- `E0-P0A-STATUS1`: actual status classifies workspace/standard-library authority as missing, project/package scope and external representation as blocked, structured outcomes as missing, graph-health inference as wrong, and the target as `0/3`.

## E6 - P6 Evidence

### P6-A - Source-backed declaration-universe decision

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E6-P6A-IMPACT1` | fresh source/config/packaging/consumer owner file-detail and impact after Child 05 handoff | pending |
| `E6-P6A-SRC1` | current declaration inputs, resolver stages, config readers, packaging/runtime, and failure path inventory | pending |
| `E6-P6A-ORACLE1` | bounded plus general TypeScript differential for standard-library, unavailable, config, and language-isolation cases | pending |
| `E6-P6A-CONSUMER1` | exact graph/persistence/reader facts required for external targets/outcomes | pending |
| `E6-P6A-DECISION1` | compared mechanisms, selected behavior/mechanism, supported config, owner map, side effects, conditional P6-C1/P6-C2 decisions, and updated later-slice steps | pending |
| `E6-P6A-BUILD1` | full baseline build and real current resolver/graph-health output used by the decision | pending |
| `E6-P6A-REVIEW1` | Supervisor PASS for the source-backed decision | pending |
| `E6-P6A-COMMIT1` | isolated P6-A decision/ledger commit and worktree inventory; no production change claimed | pending |

### P6-B - TypeScript standard-library authority

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E6-P6B-IMPACT1` | fresh impact for every selected production/test/asset owner | pending |
| `E6-P6B-SRC1` | production implementation exactly matches the P6-A mechanism and behavior | pending |
| `E6-P6B-PROVENANCE1` | declaration/config/provenance or equivalent reproducibility proof required by P6-A | pending |
| `E6-P6B-BUILD1` | full build and packaging/runtime validation required by the selected design | pending |
| `E6-P6B-TEST1` | general stdlib, target names as ordinary data, config, unavailable, determinism, and language-isolation tests | pending |
| `E6-P6B-BENCH1` | measurements required by the selected mechanism; no invented metric | pending |
| `E6-P6B-REVIEW1` | Supervisor PASS for P6-B | pending |
| `E6-P6B-DETECT1` | detect-changes result before commit | pending |
| `E6-P6B-COMMIT1` | isolated P6-B commit hash | pending |

### P6-C1 - Conditional project/package declarations

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E6-P6C1-IMPACT1` | fresh exact lookup-owner impact, or evidence that no production owner is required | pending |
| `E6-P6C1-SCOPE1` | P6-A decision marking exact cases active or preserve-only | pending |
| `E6-P6C1-SRC1` | scoped production diff, or explicit no-production-diff proof | pending |
| `E6-P6C1-BUILD1` | full build for the accepted state | pending |
| `E6-P6C1-TEST1` | required present/missing/config/security behavior or preserve-only validation | pending |
| `E6-P6C1-REVIEW1` | Supervisor PASS for active/preserve-only decision | pending |
| `E6-P6C1-DETECT1` | detect-changes before implementation commit when production changed; N/A reason otherwise | pending |
| `E6-P6C1-COMMIT1` | isolated P6-C1 commit hash | pending |

### P6-C2 - External target representation

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E6-P6C2-IMPACT1` | fresh graph/persistence/reader impact for exact consumers | pending |
| `E6-P6C2-DESIGN1` | minimum representation selected from P6-A/current consumer evidence | pending |
| `E6-P6C2-SRC1` | production representation/materialization diff with external provenance | pending |
| `E6-P6C2-BUILD1` | full repository build | pending |
| `E6-P6C2-TEST1` | resolved/unavailable/internal-external separation/duplicate/provenance cases | pending |
| `E6-P6C2-PARITY1` | graph and only affected persistence/readers preserve the representation with zero repository ownership pollution | pending |
| `E6-P6C2-REVIEW1` | Supervisor PASS for P6-C2 | pending |
| `E6-P6C2-DETECT1` | detect-changes result before commit | pending |
| `E6-P6C2-COMMIT1` | isolated P6-C2 commit hash | pending |

### P6-C3 - Structured outcomes

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E6-P6C3-IMPACT1` | fresh outcome/finalizer/consumer impact | pending |
| `E6-P6C3-SRC1` | production finalizer with one outcome and no name-based override | pending |
| `E6-P6C3-STATUS1` | exact status/field set required by accepted P5/P6 cases, with stage/target/reason/proof behavior | pending |
| `E6-P6C3-BUILD1` | full repository build | pending |
| `E6-P6C3-TEST1` | all accepted statuses, precedence, one-result exclusivity, and capability outcomes | pending |
| `E6-P6C3-PARITY1` | affected persistence/readers retain equal status/stage/target/reason/proof fields | pending |
| `E6-P6C3-REVIEW1` | Supervisor PASS for P6-C3 | pending |
| `E6-P6C3-DETECT1` | detect-changes result before commit | pending |
| `E6-P6C3-COMMIT1` | isolated P6-C3 commit hash | pending |

### P6-D - Graph-health and target proof

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E6-P6D-IMPACT1` | fresh inventory/impact for every same-invariant health/output adapter | pending |
| `E6-P6D-SRC1` | mechanical outcome projection and removal of affected target-text inference | pending |
| `E6-P6D-BUILD1` | full repository build | pending |
| `E6-P6D-TEST1` | mapping/no-heuristic/three-site/affected-regression behavior | pending |
| `E6-P6D-PARITY1` | graph and affected graph-health/readers agree on exact outcomes | pending |
| `E6-P6D-TARGET1` | exact `Promise`, `Math.max`, and `Math.min` outcomes (`3/3`) and zero in-repository analyzer gaps | pending |
| `E6-P6D-ORACLE1` | independent TypeScript oracle for all three target sites and accepted environment | pending |
| `E6-P6D-BOUNDARY1` | target pre/post HEAD, worktree, source hashes, graph path, and artifact inventory | pending |
| `E6-P6D-REVIEW1` | Supervisor PASS for P6-D | pending |
| `E6-P6D-DETECT1` | detect-changes result before commit | pending |
| `E6-P6D-COMMIT1` | isolated P6-D commit hash | pending |

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E6-PNA-REVIEW1` | independent Supervisor acceptance of Child 06 decision, source, runtime, target, ledgers, benchmarks, and commits | pending |
| `E6-PNB-CLEAN1` | dead-work inventory, removal result, final diff, and Supervisor confirmation | pending |
| `E6-PNC-DETECT1` | final detect-changes evidence after accepted cleanup | pending |
| `E6-PNC-COMMITS1` | ordered P6-A/P6-B/P6-C1/P6-C2/P6-C3/P6-D commit hashes and worktree ownership | pending |
| `E6-PNC-HANDOFF1` | exact accepted outcomes/metrics supplied to Child 07 and its refreshed opening status | pending |
