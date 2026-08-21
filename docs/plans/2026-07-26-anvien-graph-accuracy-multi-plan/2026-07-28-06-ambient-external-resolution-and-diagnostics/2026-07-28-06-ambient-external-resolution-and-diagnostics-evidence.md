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

- `E6-P6A-IMPACT1` — recorded: before graph work, `anvien analyze --force` scanned `1,985`, parsed `743`, failed `0`, produced `116,467` nodes / `160,591` relationships at indexed/current commit `ec765debff335540c77d409ebb2c9f45e4a0a77d`, `stale=false`. Fresh related-file counts are analyze `187`, resolution types `34`, indexes `58`, resolve `60`, emit `50`, graph types `256`, scopeir kinds `250`, Ladybug CSV `23`, schema `22`, graph-health policy `31`, diagnostics `30`, ResolutionGap inputs `41`, processes `14`, MCP context `28`, impact `41`, rename `9`, HTTP graph `22`, and filecontext `46`. Fresh upstream impacts: `analyze.Run` CRITICAL `24/9/8/23`; `buildWorkspace` CRITICAL `59/24/8/23`; `resolveCall`, `resolveAccess`, `resolveTypeAnnotation` each CRITICAL `28/11/7/32`; `emitUnresolvedReference` CRITICAL `9/4/2/34`; `graph.Node` CRITICAL `1,717/273/48/428`; `scopeir.NodeLabel` CRITICAL `1,846/291/47/422`; `graphhealth.Diagnostic` CRITICAL `70/17/5/77`; `nodeCSVRow` CRITICAL `25/11/1/12`; `NodeSchema` MEDIUM `8/6/2/0`; `classifyDiagnostic` HIGH `10/4/3/0`; `SourceBackedResolutionGapInputs` CRITICAL `13/7/1/15`; `processes.Apply` CRITICAL `29/8/7/31`; `buildCallsGraph` CRITICAL `28/8/6/25`; `contextSymbolPayload` CRITICAL `1/1/1/7`; `impactItemPayload` CRITICAL `11/5/2/17`; `collectRenameChanges` CRITICAL `1/1/1/8`. Values are `symbols/files/modules/processes`; HIGH/CRITICAL are warnings, not bans.

- `E6-P6A-SRC1` — recorded: production has no semantic TypeScript config reader. `resolution.Options` contains only two compatibility booleans; `buildWorkspace` accepts only repository `[]scopeir.ScopeIR`; resolver order is repository scope/import/member before generic unresolved emission; current Diagnostic/ResolutionGap carriage has no authority/catalog/config/stage/reason/candidate contract; graph-health guesses from target text. `anvien/package.json` ships only `bin` and `go-src`, so runtime has no TypeScript/Node/node_modules authority. `anvien-web/package-lock.json` resolves dev-only TypeScript `5.9.3`, Apache-2.0, integrity `sha512-jl1vZzPDinLr9eUt3J/t7V6FgNEw9QjvBPdysz9KfQDD41fQrC2Y4vKQdiaUpFT4bXlb1RHhLpp8wtm6M5TgSw==`; its package manifest currently uses range `^5.4.5`, which is not accepted as catalog provenance without an exact generation gate.

- `E6-P6A-ORACLE1` — recorded: local TypeScript `5.9.3` contains `100` `lib*.d.ts` files / `3,141,835` bytes; observed `target ES2022` closure is `63` / `2,389,540`, explicit ES2015 `13` / `310,653`, and explicit ES5 `3` / `232,957`. `Math.max/min` are in `lib.es5.d.ts`; `Promise` type can be visible in ES5 while the runtime value is supplied by `lib.es2015.promise.d.ts`, proving meaning lanes are required. Exact compiler vectors are `10/10`: default ES2015 values fail; ES2022 passes; ES5 fails; ES2015 passes; ES5+ES2015.Promise passes Promise but fails Map/Set/Symbol; noLib fails globals; invalid lib raises compiler-option error; type-only Promise+Math pass under default and ES5 but fail under noLib. No network, install, or package script ran. Debug-only probe files are not evidence artifacts and are deleted before handoff; the immutable report records the vector contract for P6-B to reproduce as proper post-code fixtures.

- `E6-P6A-CONSUMER1` — recorded: Ladybug supports a fixed node table/pair list and skips unknown labels or relationships whose endpoint labels are unsupported; a resolver-only external identity would lose persisted parity. `processes.buildCallsGraph` currently consumes all CALLS and must exclude external endpoints. MCP context/impact use fixed projections and need explicit external provenance/non-editable semantics; rename must reject external targets before edit collection. Graph-health/ResolutionGap must consume structured outcomes. Graph JSON and HTTP/Web are generic transports/fallbacks and file-detail is generic absent contrary later impact, so those remain validate/preserve-only. Required representation is referenced-only `ExternalSymbol`; capability-unavailable stays an outcome/gap and no synthetic `IMPORTS` is emitted.

- `E6-P6A-DECISION1` — recorded in the single Architect/Planner handoff `reports/system-architect/rp_system-architect_260822_033607_by_gpt-5_p6a_declaration_universe_decision.md`, verdict `READY_FOR_SUPERVISOR` and explicitly not self-accepted. Selected mechanism is offline exact-TypeScript generation -> checked-in versioned compact DTO -> Go-embedded immutable lookup with full provenance/hash validation and no runtime Node/network/install/scripts. Supported profile is TypeScript-only zero/one root JSONC `tsconfig.json` with `target`, `lib`, `noLib`; unsupported/invalid topology is `capability_unavailable`. Repository/P5 result remains first and explicit-import failure terminal. Status set is `resolved_internal`, `resolved_external`, `unresolved`, `capability_unavailable`. P6-C1 is preserve-only; P6-C2 is active. The requested `03:30 +07` durable handoff deadline was already missed when the checkpoint resumed at `03:34:32 +07`; this is disclosed for Main/Supervisor disposition. At Main intervention `03:42:26 +07`, `docs/contracts/graph-accuracy-contract.md` had an out-of-boundary tracked `18+/0-` deviation; the lane stopped contract writes and did not revert, overwrite, stage, or commit it.

- `E6-P6A-BUILD1` — N/A/not run by explicit P6-A boundary: no production/test/runtime code changed, and the repository package lifecycle includes package scripts while this lane forbids package install/script execution. A partial command is not mislabeled a full build. The fresh real non-UI boundary is the successful `anvien analyze --force` recorded by `E6-P6A-IMPACT1`; every later production slice still requires exact lock/process preflight, full build, nearest real boundary, and regressions.

- `E6-P6A-REVIEW1` — recorded: initial independent report `reports/Supervisor/rp_supervisor_260822_041014_by_gpt-5_p6a_declaration_universe_decision.md`, `11,874` bytes / `94` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `5679FB24894F5C51E7AF4EB46FC2D8B534F93A9EB838C68CB5F3CF5FCFD66290`, verdict `REJECT`, cleared the complete architecture/source/graph/impact/package/oracle/config/P6-C1/P6-C2 invariant and rejected only the out-of-boundary contract section with stale Architect-report hash. Reject-only repair restored `docs/contracts/graph-accuracy-contract.md` to HEAD blob `2020b479f509f77a1629016526410e9025501387` with zero diff and left the four ledgers/report identities unchanged. Final report `reports/Supervisor/rp_supervisor_260822_041635_by_gpt-5_p6a_reject_only_resubmission.md`, `6,011` bytes / `83` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `39CE249E9D13F1C77FE6F61DD6E9B1D2E4000B004CE3EBBF461C762F3FA28384`, verdict `PASS`, inherits the technical clearances and records no residual same-invariant surface. The disclosed `03:30 +07` deadline miss is non-blocking.

- `E6-P6A-COMMIT1` — immediate closure invariant: commit exactly the four living Child 06 ledgers, immutable Architect report, initial Supervisor REJECT, and final reject-only PASS. Contract, production, tests, fixtures, target, Child 05, and sixteen protected Main handoffs remain outside the manifest. The resulting commit identity is the P6-B handoff anchor and is supplied directly after commit without another P6-A behavior gate.

- `E6-P6A-DEPEND1` — recorded: authoritative checkout root is `E:\Anvien`; current HEAD equals exact Child 05 four-ledger closure commit `ec765debff335540c77d409ebb2c9f45e4a0a77d`, whose sole parent is Pn-B commit `fcc44334c0f75b3b19046dc8f9f4de40eb459fa9`; Pn-A commit `b68e738d64eebea65a045afbf0b12d94dd43cbf4` and all accepted P5 implementation commits are ancestors. The closure commit changes exactly the four Child 05 ledgers. Tracked diff and index were empty before Child 06 edits; exactly fifteen protected untracked Main handoffs were present and remain out of scope. The accepted predecessor contract is immutable: module/file lookup remains separate from syntax-derived export tables and deterministic export traversal; explicit-import failure is fail-closed; terminal/hop/failure proof stays generic-first; fixed-corpus physical/resolver-emitted/persisted `IMPORTS` remains `5,072 / 5,072 / 5,088`.

P6-A report seal: `reports/system-architect/rp_system-architect_260822_033607_by_gpt-5_p6a_declaration_universe_decision.md`, `25,326` bytes / `289` lines / SHA-256 `77D5E9AC8D76D98C76D1816C8D6E69265D4AFB30367E3DA50DF3EAA3445D7BA2`.

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E6-P6A-IMPACT1` | fresh source/config/packaging/consumer owner file-detail and impact after Child 05 handoff | recorded and independently accepted |
| `E6-P6A-SRC1` | current declaration inputs, resolver stages, config readers, packaging/runtime, and failure path inventory | recorded and independently accepted |
| `E6-P6A-ORACLE1` | bounded plus general TypeScript differential for standard-library, unavailable, config, and language-isolation cases | recorded `10/10`; P6-B durable fixtures pending |
| `E6-P6A-CONSUMER1` | exact graph/persistence/reader facts required for external targets/outcomes | recorded and independently accepted |
| `E6-P6A-DECISION1` | compared mechanisms, selected behavior/mechanism, supported config, owner map, side effects, conditional P6-C1/P6-C2 decisions, and updated later-slice steps | recorded and accepted by `E6-P6A-REVIEW1` |
| `E6-P6A-BUILD1` | full baseline build and real current resolver/graph-health output used by the decision | N/A/not run by boundary; fresh real CLI analyze recorded |
| `E6-P6A-REVIEW1` | Supervisor PASS for the source-backed decision | recorded — initial narrow REJECT, exact contract-boundary repair, final reject-only PASS; residual none |
| `E6-P6A-COMMIT1` | isolated P6-A decision/ledger commit and worktree inventory; no production change claimed | immediate exact seven-path commit; resulting hash is the P6-B handoff anchor |

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
| `E6-P6C1-SCOPE1` | P6-A decision marking exact cases active or preserve-only | P6-A records preserve-only; P6-C1 closure/review pending |
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
