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
- `E0-P0A-GRAPH1` — recorded: the pre-ledger refresh at commit `7042dda8bfc02ee42b40c3a1c0ede89138439481` reported `1,557` scanned files, `676` parsed/parsed-code files, `0` failed files, `85,110` graph nodes, and `123,978` relationships; the required post-ledger refresh with the same command also reports `1,557/676/0`, while the graph inventory is `85,114` nodes and `123,982` relationships. `anvien status`/`.anvien/meta.json` agree on the commit and `stale=false`; the observed `+4/+4` inventory delta occurred across a documentation-only worktree diff, while `git diff --exit-code -- internal` confirms no production-source change. Per-document attribution of that delta was not measured.
- `E0-P0A-SOURCE1` — recorded: `internal/resolution/indexes.go:814-823` (`graphIDForDef`) derives a graph ID from label, cleaned file path, qualified/name, and optional arity; it does not consume `DefinitionFact.ID`, range, lexical scope, or a distinct meaning lane.
- `E0-P0A-SOURCE2` — recorded: `internal/scopeir/range.go:3-8` exposes only four integers; `internal/providers/tsjs/nodes.go:18-26` uses one-based lines and unchanged tree-sitter columns; no shared encoding, interval, or selection-range contract is declared. `internal/scopeir/position_index.go:51-52,95-104` treats the end coordinate as included, so P1-B must ratify semantics before changing a shared range type.
- `E0-P0A-SOURCE3` — recorded and reclassified: `internal/scopeir/definition_index.go:7-17` is first-writer-wins for duplicate `DefinitionFact.ID`, but repository search plus source-flow inspection found no production caller. Its only callers are `internal/scopeir/scope_indexes_test.go` and `internal/scopeir/legacy_p7_scope_extractor_conversion_test.go`; it is preserve-only, not the bounded production loss owner.
- `E0-P0A-SOURCE4` — recorded: `internal/graph/types.go:96-104` (`Graph.AddNode`) silently replaces an existing payload for a duplicate ID; `Graph.init` at `:270-281` also rebuilds an index without rejecting duplicate IDs already present in `Nodes`.
- `E0-P0A-SOURCE5` — recorded: the normal path is `analyze.parseFiles` (`internal/analyze/analyze.go:799-857`) → provider ScopeIR → `BuildCrossFileBinding/buildWorkspace` (`internal/resolution/resolve.go:23-48`, `internal/resolution/indexes.go:72-199`) → `ResolveBoundInto` (`internal/resolution/resolve.go:50-90`) → `emitDefinitionNodes`/`emitter.emitNode` (`internal/resolution/emit.go:31-34,136-203`) → `Graph.AddNode`; no hidden `BuildDefinitionIndex` stage is present and no orchestration edit is proven for P0.
- `E0-P0A-SOURCE6` — recorded: `DefinitionFact` (`internal/scopeir/facts.go:3-28`) has one full `Range`, `OwnerID`, `Label`, and optional `Visibility`, but no selection range, lexical-scope field, or explicit meaning mask. TSJS `addDefinition` (`internal/providers/tsjs/definitions.go:79-131`) stores the construct range and scope binding/owned-definition links, while `nodeRange` stores no identifier-token range.
- `E0-P0A-SOURCE7` — recorded: TSJS scope construction (`internal/providers/tsjs/scopes.go:54-121`) creates deterministic range-derived scope IDs and `innermostScopeID` links definitions to lexical scopes through `ScopeFact.OwnedDefIDs`; `ScopeIR.NormalizeInPlace`/`compareDefinition` sort facts but do not deduplicate them (`internal/scopeir/ir.go:69-110`, `internal/scopeir/sort_keys.go:23-35`). The source has lexical-owner input even though graph projection omits it.
- `E0-P0A-SOURCE8` — recorded: `emitDefinitionNodes` (`internal/resolution/emit.go:136-203`) emits one attempted node and `DEFINES` edge per `defsByFile` occurrence, but node properties retain only start/end lines (not columns, selection range, provider ID, or scope). `emitter.emitRelationship` (`:36-49`) is a separate relationship merge pipeline and remains outside the node-collision owner.
- `E0-P0A-SOURCE9` — recorded: existing tests assert legacy semantics that must be classified before any implementation change: `Graph.AddNode` replacement in `internal/graph/types_test.go:91-105`, `BuildDefinitionIndex` first-writer behavior in `internal/scopeir/scope_indexes_test.go:9-45` and `legacy_p7_scope_extractor_conversion_test.go:64-85`, inclusive position-index boundaries in `scope_indexes_test.go:63-93`, and intentional semantic relationship merging in `internal/resolution/resolution_test.go` (`TestResolveMergesDuplicateSemanticEdgesAndCountsUnresolved`). These are sibling contracts, not proof that the bounded identity behavior is correct.
- `E0-P0A-FD1` — recorded: fresh `anvien file-detail internal/resolution/indexes.go --repo E:\Anvien --json` reports `46` related files, `192` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD2` — recorded: fresh file-detail for `internal/scopeir/range.go` reports `227` related files, `6` symbols, risk `medium`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD3` — recorded: fresh file-detail for `internal/scopeir/definition_index.go` reports `225` related files, `14` symbols, risk `medium`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD4` — recorded: fresh file-detail for `internal/graph/types.go` reports `238` related files, `92` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD5` — recorded: fresh file-detail for `internal/resolution/emit.go` reports `42` related files, `75` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD6` — recorded: fresh file-detail for `internal/scopeir/facts.go` reports `231` related files, `119` symbols, risk `medium`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD7` — recorded: fresh file-detail for `internal/scopeir/ir.go` reports `229` related files, `29` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD8` — recorded: fresh file-detail for `internal/scopeir/sort_keys.go` reports `226` related files, `18` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD9` — recorded: fresh file-detail for `internal/providers/tsjs/definitions.go` reports `16` related files, `30` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD10` — recorded: fresh file-detail for `internal/providers/tsjs/nodes.go` reports `20` related files, `42` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD11` — recorded: fresh file-detail for `internal/providers/tsjs/scopes.go` reports `17` related files, `33` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD12` — recorded: fresh file-detail for `internal/providers/tsjs/extract.go` reports `23` related files, `33` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD13` — recorded: fresh file-detail for `internal/resolution/resolve.go` reports `50` related files, `40` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD14` — recorded: fresh file-detail for `internal/resolution/types.go` reports `31` related files, `64` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-FD15` — recorded: fresh file-detail for `internal/analyze/analyze.go` reports `182` related files, `220` symbols, risk `high`, and `stale=false/changedSinceAnalyze=false`.
- `E0-P0A-IMPACT1` — recorded exact upstream impacts with `--include-tests`: `graphIDForDef` is `CRITICAL`, `23` symbols/`13` files/`3` modules/`18` processes (`backend 6`, `backend_test 17`); `buildWorkspace` is `CRITICAL`, `41` symbols/`17` files/`8` modules/`28` processes (`backend 7`, `backend_test 33`, `cli_launcher 1`); `Range` is `CRITICAL`, `854` symbols/`143` files/`22` modules/`75` processes (`backend 503`, `backend_test 351`); `DefinitionFact` is `CRITICAL`, `586` symbols/`80` files/`23` modules/`63` processes (`backend 181`, `backend_test 405`).
- `E0-P0A-IMPACT2` — recorded supporting impacts: `Graph.AddNode` is `CRITICAL`, `296` symbols/`70` files/`11` modules/`82` processes (`backend 36`, `backend_test 179`, `api_test 81`); `emitNode` is `CRITICAL`, `6` symbols/`5` files/`2` modules/`19` processes; `emitDefinitionNodes` is `CRITICAL`, `24` symbols/`9` files/`7` modules/`35` processes; `BuildDefinitionIndex` is `LOW`, `2` test symbols/`2` test files/`0` production symbols; TSJS `nodeRange` is `MEDIUM`, `8` symbols/`4` files; `ScopeIR.NormalizeInPlace` and `compareDefinition` are `LOW` with `3` and `4` impacted symbols respectively. Each CRITICAL/HIGH value is recorded as a scope warning, not an edit prohibition.
- `E0-P0A-BOUNDARY1` — recorded: `E:\cheapapp.org` is an in-place validation target; target source/worktree is preserve-only and only normal target-repository output is allowed.
- `E0-P0A-BOUNDARY2` — recorded from independent target-boundary checks: target HEAD/branch is `a869876ab6262dacde6cd5d432d099a91852a646`/`master`; worktree has `13` pre-existing entries (`7` tracked modifications, `6` untracked, `0` staged); oracle source hash is `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C`; existing target graph hash/size is `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`/`315,569,880` bytes with `1,359` files, `84,807` nodes, and `114,125` relationships. The four bounded declarations are `time@207` in `parseTime`, `time@214` in the `latestIso` reducer callback, `now@262` in `buildEmailOperationsReport`, and `now@501` in `readEmailOperationsReport`; P0 did not run target analyze.
- `E0-P0A-STATUS1` — recorded: the actual-status ledger distinguishes source facts, production owners, preserve-only utilities, and the external target boundary; P0 is now Supervisor-accepted, P1-A is open as documentation-only, and production implementation remains gated by the ordered P1 slices.
- `E0-P0A-STATUS2` — recorded: the invariant family map and phase touch map are updated for identity, ranges/selection, lexical owner/meaning, occurrence conservation, collision/mutation, projection/endpoints, and sibling relationship/persistence surfaces; forbidden fallbacks are explicit.
- `E0-P0A-REVIEW1` — recorded: `reports/Supervisor/rp_supervisor_260810_095511_by_gpt-5-codex_child01_p0a.md` independently reviews the current P0-A source, file-detail, impact, target-boundary, ledger, and diff evidence and returns `PASS`. The report accepts only the P0 classification/documentation gate; it does not claim production repair or runtime `4/4`.

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
