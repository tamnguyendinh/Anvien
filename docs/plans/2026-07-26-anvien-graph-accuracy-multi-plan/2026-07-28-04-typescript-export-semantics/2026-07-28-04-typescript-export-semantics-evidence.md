# Anvien First-Class TypeScript Export Semantics Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`

## Evidence Rules

- The problem-origin report establishes why the multi-plan exists. It explicitly labels its proposed architecture DRAFT; proposed implementation is not accepted evidence.
- The causal synthesis and final Supervisor report verify the bounded investigation only. They do not approve remediation or establish complete TypeScript module semantics.
- Current production facts require fresh source, graph, file-detail, impact, runtime, and test evidence. Historical file paths are investigation leads, not current ownership proof.
- A pending evidence ID is a reserved target, not proof.
- Each implementation slice requires the complete evidence set named in the plan: impact, source change, full build, behavior test, real boundary, Supervisor, detect-changes, and commit.
- Long counts and before/after metrics belong in the benchmark ledger.
- No target source or target-side report may be used as a plan artifact.

### Evidence ID Naming

Evidence IDs use `E<phase>-<item>-<kind><n>`. IDs are never reused for a different fact. Exact IDs, not broad section names, must be cited by actual status and benchmark rows.

## E0 - P0 Evidence

Matching plan item: `P0-A`

### Completed documentation-authority evidence

- `E0-P0A-RULE1`: on 2026-08-10, `E:\Anvien\AGENTS.md` was read in full before editing. It requires planner use for plan work, current graph refresh before graph evidence, file-detail and impact before relevant code edits, code before tests, full build before validation, Supervisor acceptance, detect-changes, and per-slice commits.
- `E0-P0A-SKILL1`: the planner skill and all four standard templates were read in full before rewriting this plan set. The rewritten files retain plan, evidence, benchmark, and actual-status roles with exact checklist/evidence mapping.
- `E0-P0A-ORIGIN1`: `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` was read in full. It identifies missing module-export metadata and proposes distinct export facts, but explicitly states that its architecture is DRAFT. This ledger uses its defect description and acceptance target, not its unapproved architecture.
- `E0-P0A-VERIFY1`: `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md` was read in full. Bounded finding C4 confirms 21 selected exported TypeScript declarations existed as definitions while zero graph rows carried export/visibility metadata because the investigated provider fact was not populated. The report expressly provides no remediation authorization.
- `E0-P0A-VERIFY2`: `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md` was read in full. Its PASS accepts the bounded causal record only, including the 21 missing metadata facts; it does not approve a fix or claim global module accuracy.
- `E0-P0A-ANVIEN1`: before these eight ledger rewrites, `anvien status` reported the self-index up to date at commit `238aec0`. Concept queries were run for binding-pattern and export ownership; their broad ranking did not establish exact export ownership or an affected-reader denominator. They are therefore not used to close P0, and graph freshness must be renewed after documentation changes.
- `E0-P0A-SCOPE1`: the corrected Child 04 scope retains export syntax, direct projection, affected persistence, and the `21/21` target gate. It excludes terminal module resolution, unrelated artifact behavior, and unrelated consumer work.

### Required current-state evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E0-P0A-GRAPH1` | fresh analyze and current HEAD/source basis after plan correction | recorded — excluded graph current at `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; `1,126/626/0`, `80,908/120,167` |
| `E0-P0A-SRC1` | full current-source inventory of export facts, direct/re-export syntax, compatibility fields, graph projection, persistence, actual consumers, and relevant tests | recorded — owner/syntax/consumer/persistence matrix below |
| `E0-P0A-FD1` | exact `file-detail` outputs and related-file counts for every candidate editable file | recorded — current/non-stale counts below |
| `E0-P0A-IMPACT1` | exact upstream impacts and blast-radius report for every candidate function/method/exported symbol | recorded — default upstream file and symbol impacts below |
| `E0-P0A-STATUS1` | completed current status matrix, touch map, R1 refresh, and Final P0 Decision | recorded — P0 complete; P4-A work steps narrowed to the four deterministic-contract owners |

### P0-A current graph and Git basis

- `E0-P0A-GRAPH1`: at clean HEAD `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`, `anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"` exited `0`: scanned `1,126`, parsed code `626`, failed `0`, graph `80,908` nodes / `120,167` relationships. `anvien status` reports indexed/current commit `3e25f9c` and up-to-date.
- Roadmap and all four Child 04 ledgers are LOW, non-stale, with related-file counts `28/1/1/1/1`; upstream file impact is LOW with `0` impacted files/flows for every document.
- Graph-backed API discovery returned no extracted route/tool definitions even though source registers graph/file-context handlers. This is recorded as a route/tool extraction limitation; HTTP/MCP classification below comes from full source reads, not a false no-surface conclusion.

### P0-A current source and syntax inventory

- `E0-P0A-SRC1`: `internal/scopeir/facts.go` contains `DefinitionFact.Visibility`, `ImportFact`, and binding-oriented `ExtractionDiagnosticFact`, but no `ExportFact`, export kind/meaning enum, export provenance contract, or export diagnostic code. `internal/scopeir/ir.go` has no export collection; `internal/scopeir/sort_keys.go` has no export comparator.
- The sole current TypeScript export syntax method is `internal/providers/tsjs/imports.go::collector.emitExportStatement`. A source-less `export_statement` returns immediately, so direct declarations, local export lists, default exports, and type-only direct/local forms emit no export fact or diagnostic. A source-bearing named re-export becomes `ImportReexport`; `export *` becomes `ImportWildcard`; namespace aliases and type-only markers are lost. `internal/providers/tsjs/definitions.go` still emits named declarations under export wrappers but never populates `DefinitionFact.Visibility`; anonymous default expressions have no definition owner.
- Current relevant tests prove only import compatibility: `internal/providers/tsjs/extract_test.go` and `internal/providers/provider_parity_test.go` assert named re-export as `ImportFact`, while the direct JavaScript export fixture asserts only the definition. `internal/scopeir/scopeir_test.go` and `internal/scopeir/testdata/scopeir.golden.json` have no export collection/fact.
- `internal/resolution/emit.go::emitDefinitionNodes` maps only `DefinitionFact.Visibility` to graph property `visibility`. No TSJS writer populates that field, and no emitter writes `isExported` for these facts. Current Cypher confirms property `visibility` is absent; `isExported` exists on `5,714` nodes but is `true` on `0` and `false` on all `5,714`.
- Graph JSON serialization is generic and preserves arbitrary node properties. Ladybug is field-specific: `internal/lbugload/csv.go::nodeCSVRow` and `internal/lbugschema/schema.go::NodeSchema` persist `isExported` only for Function/Class/Interface/CodeElement/Method; `visibility` is dropped, and other definition labels use schemas without an export column.
- Proven semantic consumers are `internal/filecontext/context.go::exportedSymbol` (`exported` or `visibility=public|exported`), `internal/graphhealth/compute.go` (`isExported` topology/reason), `internal/processes/processes.go::findEntryPoints` (`isExported` score), `internal/embeddings/text.go::NodesFromGraph` (`isExported` text/hash), and Ladybug CSV/schema. CLI/MCP/HTTP file-context surfaces consume the derived file-context result; `/api/graph` transports raw properties. Web files and generated contracts are carriers/non-consumers for this behavior.
- Child 05 remains preserve-only: `ImportFact.TargetFile`, `TargetExportedName`, `TargetModuleScope`, `TargetDefID`, `TransitiveVia`, `LinkStatus`, barrel traversal, ambiguity, cycles, and terminal/public-API outcomes are not Child 04 facts.

### P0-A file-detail evidence

`E0-P0A-FD1` uses `anvien file-detail <path> --repo E:\Anvien --json` against the fresh excluded graph. Related count excludes the target file.

| File | Symbols | Related | In / Out / Local | Unresolved | Flows / Tests | Risk |
|------|--------:|--------:|------------------:|-----------:|--------------:|------|
| `internal/scopeir/facts.go` | 164 | 245 | 1096 / 22 / 148 | 5 | 0 / 98 | MEDIUM |
| `internal/scopeir/kinds.go` | 87 | 239 | 613 / 0 / 8 | 0 | 0 / 97 | LOW |
| `internal/scopeir/ir.go` | 47 | 243 | 1094 / 34 / 75 | 138 | 3 / 96 | HIGH |
| `internal/scopeir/sort_keys.go` | 96 | 239 | 250 / 102 / 41 | 109 | 2 / 95 | HIGH |
| `internal/providers/tsjs/definitions.go` | 129 | 24 | 8 / 107 / 14 | 308 | 2 / 4 | HIGH |
| `internal/providers/tsjs/imports.go` | 29 | 17 | 7 / 26 / 1 | 52 | 1 / 4 | HIGH |
| `internal/providers/tsjs/extract.go` | 36 | 25 | 24 / 32 / 29 | 68 | 3 / 6 | HIGH |
| `internal/resolution/emit.go` | 108 | 43 | 43 / 166 / 63 | 310 | 12 / 17 | HIGH |
| `internal/lbugload/csv.go` | 192 | 21 | 55 / 53 / 92 | 200 | 2 / 8 | HIGH |
| `internal/lbugschema/schema.go` | 41 | 21 | 52 / 8 / 29 | 50 | 1 / 6 | HIGH |
| `internal/filecontext/context.go` | 552 | 44 | 337 / 88 / 565 | 542 | 51 / 8 | HIGH |
| `internal/mcp/context.go` | 165 | 28 | 48 / 111 / 50 | 261 | 24 / 2 | HIGH |

### P0-A impact and ownership decision

`E0-P0A-IMPACT1` uses default upstream impact without `--include-tests`; linked-test counts remain in `E0-P0A-FD1`.

| File | Risk | Impacted / Direct / Files / Flows |
|------|------|-----------------------------------:|
| `internal/scopeir/facts.go` | CRITICAL | 340 / 107 / 76 / 1 |
| `internal/scopeir/kinds.go` | CRITICAL | 1354 / 143 / 182 / 1 |
| `internal/scopeir/ir.go` | CRITICAL | 123 / 54 / 30 / 1 |
| `internal/scopeir/sort_keys.go` | HIGH | 23 / 20 / 3 / 1 |
| `internal/providers/tsjs/definitions.go` | MEDIUM | 10 / 10 / 2 / 1 |
| `internal/providers/tsjs/imports.go` | LOW | 1 / 1 / 1 / 0 |
| `internal/providers/tsjs/extract.go` | CRITICAL | 24 / 11 / 11 / 1 |
| `internal/resolution/emit.go` | CRITICAL | 26 / 20 / 5 / 1 |
| `internal/lbugload/csv.go` | CRITICAL | 45 / 28 / 16 / 1 |
| `internal/lbugschema/schema.go` | CRITICAL | 55 / 26 / 27 / 1 |
| `internal/filecontext/context.go` | CRITICAL | 252 / 134 / 32 / 1 |
| `internal/mcp/context.go` | CRITICAL | 48 / 39 / 12 / 1 |

Exact symbol impacts further constrain the slices: `DefinitionFact` CRITICAL `194/42/24/83` impacted/direct/modules/processes; `ScopeIR` CRITICAL `122/53/26/90`; `ExtractionDiagnosticFact` CRITICAL `138/4/24/84`; `Extract` CRITICAL `11/3/8/35`; `emitDefinitionNodes` CRITICAL `6/1/4/33`; `exportedSymbol` CRITICAL `9/3/3/24`; `nodeCSVRow` CRITICAL `3/1/1/12`; `contextMethodMetadata` CRITICAL `2/1/1/7`; `NodeSchema` LOW `3/1/2/0`. `collector.emitDefinitionKind`, `collector.emitImportKind`, `collector.emitExportStatement`, and `collector.result` are LOW with zero symbol-level downstream impact.

`E0-P0A-STATUS1`: CRITICAL/HIGH results are blast-radius warnings, not edit bans. P4-A is narrowed to the four deterministic contract owners (`facts.go`, `kinds.go`, `ir.go`, `sort_keys.go`) plus `scopeir_test.go`/golden after production code. TSJS extraction stays inspect-only until P4-B/B1; projection/persistence/compatibility consumers stay inspect-only until P4-C and receive fresh slice-local impact before edits. Current ownership is resolved, P4-A work steps are updated, and no target or production byte was changed by P0-A.

`E0-P0A-DETECT1`: fresh excluded staged graph is `1,126/626/0` with `80,913/120,172` nodes/relationships. Full and JSON staged detection are `PASS`: `29` changed documentation sections, `5` changed files, `5` affected files, LOW risk, affected processes/flows `0/0`, gap delta `0/0`, health `0/0/0`, and complete semantic fields. Changed and affected path sets equal the exact roadmap-plus-four-ledger manifest.

## E4 - P4 Evidence

P4-A source/build/test/boundary/Supervisor/detect evidence is recorded below, with isolated commit `479e8ac229a17f2f6f94be9a4d04e07d74ac4d43`. P4-B and P4-B1 are also closed below. P4-C source/build/test/boundary and independent Supervisor REVIEW1 are now recorded and `PASS`; Main-owned detect/commit remain pending, while P4-C2 and later slices remain reserved and locked.

### P4-A — export fact, meaning, and visibility boundary

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4A-IMPACT1` | fresh file-detail/impact and one-responsibility ownership for exact editable owners | recorded — exact four production owners; CRITICAL/CRITICAL/CRITICAL/HIGH blast-radius warnings; no scope expansion |
| `E4-P4A-SRC1` | smallest production export fact/meaning contract and structured unsupported diagnostic | recorded — exact six-file candidate, `570/0`, one-site fact/diagnostic contract, deterministic deep-copy JSON, zero later-slice state |
| `E4-P4A-BUILD1` | full build after production and focused tests | recorded — canonical `npm run full-build` exit `0` on accepted bytes |
| `E4-P4A-TEST1` | deterministic serialization, one-fact-per-specifier, meaning, access separation, and zero-derived-state matrices | recorded — focused `7/7` and ScopeIR package PASS |
| `E4-P4A-BOUNDARY1` | real production fact round trip and unsupported diagnostic count | recorded — nearest six-package product boundary PASS; exact contract counts below |
| `E4-P4A-REVIEW1` | independent Supervisor PASS | recorded — PASS report `1B8DEB2F...175573B2` |
| `E4-P4A-DETECT1` | pre-commit Anvien detect-changes result | recorded — exit `0`, MEDIUM closure risk, exact tracked path set, complete semantics, zero persisted health regression |
| `E4-P4A-COMMIT1` | isolated accepted slice commit | recorded — `479e8ac229a17f2f6f94be9a4d04e07d74ac4d43` |

#### `E4-P4A-IMPACT1` — fresh graph, file-detail, and blast radius

- Fresh excluded Supervisor graph on the accepted source/report boundary: `anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"` exit `0`; scanned/parsed/failed `1,128/626/0`; `81,103/120,485` nodes/relationships.
- Final-byte file-detail summaries: `facts.go` `192` symbols / `1,127` inbound / `27` outbound / `174` local / `98` linked tests; `kinds.go` `99/619/0/10/98`; `ir.go` `57/1,106/38/89/96`; `sort_keys.go` `122/252/128/54/95`; stale `false`, changedSinceAnalyze `false`.
- Upstream file impacts are `facts.go` CRITICAL `347/110/76/1`, `kinds.go` CRITICAL `1,358/143/182/1`, `ir.go` CRITICAL `124/55/30/1`, and `sort_keys.go` HIGH `27/24/3/1` impacted/direct/files/flows. The accepted edit boundary remains exactly these four production owners.

#### `E4-P4A-SRC1` — accepted source and artifact boundary

- Exact production contract: `ExportFact`, `ExportProvenance`, `ExportDiagnosticFact`, seven `ExportKind` values, three `ExportMeaning` lanes, `ScopeIR.Exports`, `ScopeIR.ExportDiagnostics`, nested owning clones, meaning sort/dedup, and total deterministic comparators over every serialized field.
- Exact candidate: four production files plus `internal/scopeir/scopeir_test.go` and `internal/scopeir/testdata/scopeir.golden.json`; `570` insertions / `0` deletions; `git diff --check -- internal/scopeir` exit `0`; staged diff remained empty before planner closure.
- Candidate SHA-256 values are `7F2B51D8...BFCA6A`, `2D80ECAB...E5D2878`, `732EE7F8...E46E728C`, `5C155B4C...A52DB5`, `ADC375EB...E6D04`, and `C9268221...4B306`; all match the Coder and Supervisor reports.
- Preserve-only TSJS extraction, resolution/projection, and Ladybug persistence owners have zero diff. New export contract types occur only in the four ScopeIR owners and `scopeir_test.go`. `DefinitionFact.Visibility` is unchanged; no compatibility writer or Child 05-derived state exists.

#### `E4-P4A-BUILD1`, `E4-P4A-TEST1`, and `E4-P4A-BOUNDARY1`

- Pre-build lock/holder gate was free with build-related holder count `0`. Canonical `npm run full-build` exited `0`; packaged CLI `1.2.8`; Web production build transformed `2,943` modules. Non-failing npm/import/chunk warnings were retained.
- Exact focused command over seven named ScopeIR tests exited `0` with `7/7` PASS. `go test ./internal/scopeir -count=1` exited `0`.
- Nearest real non-UI boundary `go test ./internal/scopeir ./internal/providers ./internal/providers/tsjs ./internal/resolution ./internal/analyze ./cmd/binding-contract-probe -count=1` exited `0` for all six packages.
- Contract measurements: source forms `7/7`; meaning lanes `3/3`; dual value/type `1/1`; explicit type-only state `1/1`; one input fact per supplied site `7/7` (`1.0`); diagnostic classes `2/2`; nested mutable members deep-copied `3/3`; access fields changed `0`; forbidden terminal/resolution/barrel/public-API JSON names `0/7`.
- `go test ./... -count=1` exited `1` from five compile/setup-negative fixture packages plus the recorded out-of-slice C#/Dart parity baseline. It is retained as failed evidence and is not used to support PASS.

#### `E4-P4A-REVIEW1` — independent acceptance

- Verdict: `PASS`.
- Report: `reports/Supervisor/rp_supervisor_260820_221058_by_gpt-5_child04_p4a_export_fact_boundary.md`.
- Identity: `16,898` bytes / `156` LF / SHA-256 `1B8DEB2F8D5F49F285BE5AA4DF817304F8A9D8DE61E112BD2FCFEECC175573B2`.
- Main independently re-read the authority, exact source/diff, reports, candidate hashes/blobs, preserve-only siblings, usage boundary, Git provenance, and evidence. No same-slice defect or unknown drift remained. Detect/stage/commit were deliberately not supplied by Supervisor.

#### `E4-P4A-DETECT1` — fresh closure graph and change detection

- Fresh excluded graph command: `anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"`; exit `0`; scanned/parsed/failed `1,130/626/0`; graph `81,132/120,514` nodes/relationships.
- `anvien detect-changes --repo E:\Anvien --scope all` and its JSON form both exited `0` on the refreshed closure boundary.
- Summary: `196` changed semantic units, `11` tracked changed files, `10` affected files, `3` affected processes, overall risk `medium`; changed app layers `backend=132`, `backend_test=41`, `docs=23`; changed functional areas `providers=173`, `documentation=23`.
- Changed path set is exactly the six accepted ScopeIR/test/golden paths plus the roadmap and four Child 04 living ledgers. The three affected processes are `NormalizeInPlace -> CompareInt`, `NormalizeInPlace -> CompareString`, and `ProbeFile -> CloneRange`; each changed step belongs to the accepted `ir.go` normalization/clone boundary.
- Resolution source-site inventory reports `52` changed gap entities / `55` occurrences: `47` builtin, `3` standard-library, and `2` in-repo-unresolved; `50` are non-actionable and `2` analyzer-gap. This is changed-site inventory, not persisted health regression: `totalResolutionGapCount=0`, `nodesWithGaps=0`, and `degradedNodes=0`.
- Semantic status is complete for both required fields: `appLayer` and `functionalArea` each cover `81,132/81,132` nodes with `0` missing values and `0` missing source nodes.
- Untracked report provenance is outside the unstaged graph changed-file set; its four exact identities are independently verified and must appear only in the final exact staged manifest. A final staged detect/path check remains a commit-operation guard and does not reopen implementation acceptance.

### P4-B — direct/default/alias/type-only extraction

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4B-IMPACT1` | exact direct-export extraction file/symbol impact | recorded — fresh excluded graph/file-detail/impact; HIGH file warnings and LOW exact-symbol impacts |
| `E4-P4B-SRC1` | production direct/default/alias/type-only extraction | recorded — exact three-file candidate; one fact per eligible site; structured diagnostics; visibility and re-export compatibility preserved |
| `E4-P4B-BUILD1` | full build | recorded — canonical `npm run full-build` exit `0` |
| `E4-P4B-TEST1` | direct/default/anonymous/alias/type-only/multi-specifier matrix and negative controls | recorded — focused TS/JS and diagnostics matrices PASS |
| `E4-P4B-BOUNDARY1` | real provider/ScopeIR direct export facts and exact counts | recorded — nearest `tsjs`, `scopeir`, and `providers` boundary PASS |
| `E4-P4B-REVIEW1` | independent Supervisor PASS | recorded — REVIEW1 REJECT sole cleanup finding closed; resubmission PASS report `rp_supervisor_260821_001554_by_gpt-5_child04_p4b_export_facts_resubmission.md` |
| `E4-P4B-DETECT1` | pre-commit Anvien detect-changes result | recorded — final detect passed before isolated commit `11a37aa8ec0320dd93258c058b088d1070aa778d`; `537` changed semantic units, `8` changed/affected files, `3` affected processes, MEDIUM risk, zero resolution-health degradation |
| `E4-P4B-COMMIT1` | isolated accepted slice commit | recorded — `11a37aa8ec0320dd93258c058b088d1070aa778d`, exact 14-file P4-B boundary, no push |

### P4-B1 — star/namespace/re-export syntax

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4B1-IMPACT1` | exact re-export syntax owner impact | recorded — fresh excluded graph/file-detail/impact; imports.go HIGH `18/18`, extract_test.go CRITICAL `63/60`, exact recovery/test symbols LOW/0 |
| `E4-P4B1-SRC1` | production named/default/star/namespace/type-only re-export syntax extraction | recorded — exact two-file candidate; token-ordered comment/trivia recovery, one immutable fact per eligible site, single compatibility derivation, no later-slice state |
| `E4-P4B1-BUILD1` | full build | recorded — fresh canonical `npm run full-build` exit `0`, runtime `1.2.8`, Web `2,943` modules, Vite `21.84s` |
| `E4-P4B1-TEST1` | syntax matrix with one fact per specifier and no terminal resolution state | recorded — focused `6/6`, full tsjs PASS, nearest `3/3`, resolution/analyze `2/2`; TS/JS comment and sibling matrices pass |
| `E4-P4B1-BOUNDARY1` | real provider/ScopeIR re-export syntax output | recorded — independent parser → Extract → ScopeIR probes pass comment-after-comma, comment-before-comma, no-comment, and newline at `2/1/2`, exact fields, no Broken, zero terminal state |
| `E4-P4B1-REVIEW1` | independent Supervisor PASS | recorded — `reports/Supervisor/rp_supervisor_260821_031004_by_gpt-5_child04_p4b1_comment_recovery_review3.md`, `11,830` bytes / `136` LF / SHA-256 `07DD5BB92F169C5923C0DBCB597F914A28E594496ACDADB55D41F13DE364421C` |
| `E4-P4B1-DETECT1` | pre-commit Anvien detect-changes result | recorded — final exit `0`; `305` changed semantic units, `7` changed/affected files, `1` affected process, MEDIUM risk; changed layers backend `219`, backend_test `61`, docs `25`; `totalResolutionGapCount=0`, `nodesWithGaps=0`, `degradedNodes=0`, complete semantic fields; final graph `1,146/626/0`, `82,079/121,780` |
| `E4-P4B1-COMMIT1` | isolated accepted slice commit | recorded — `42d167aaf28446ac0b3de479a8afefabb8d06736`; final handoff HEAD `a12e0ccb77bda7da8aad2ec29b8050b55f81bc08`; worktree/index clean; no push |

### P4-C — graph and affected persistence projection

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4C-AUTH1` | accepted predecessor closure, clean successor Git boundary, current plan-document graph evidence, and explicit authority to open only P4-C | recorded — clean HEAD `871189b8c6a4e4bb9ff538407232c913b8cf4db6`; P4-B1 source/test unchanged; fresh excluded graph `1,147/626/0`, `82,087/121,788`; P4-C sole open; P4-C2/Child 05/target locked |
| `E4-P4C-IMPACT1` | exact fact-to-graph/persistence blast radius and actual affected-consumer inventory | recorded — four editable owners; HIGH/CRITICAL warnings and exact tuples independently confirmed |
| `E4-P4C-SRC1` | smallest production projection/compatibility change with no blanket adapter expansion | recorded — four production owners plus focused/owner tests and goldens; no extraction/terminal/target paths |
| `E4-P4C-BUILD1` | full build | recorded — canonical `npm run full-build` exit `0`; `1,855/735/0`, graph `113,496/156,003` |
| `E4-P4C-TEST1` | graph fact conservation, direct count, access separation, compatibility derivation, negative controls, and zero terminal state | recorded — `414` Export nodes, `240` direct runtime records, `239` runtime local records, compatibility drift `0`, negative controls `0`, forbidden state `0` |
| `E4-P4C-BOUNDARY1` | real graph/affected persistence output with field differences and orphan counts | recorded — Graph JSON/Ladybug `414/414`, `11,592` normalized field comparisons, differences `0`, IDs missing both ways `0`, orphan references `0` |
| `E4-P4C-REVIEW1` | independent Supervisor PASS | recorded — `PASS`; report `reports/Supervisor/rp_supervisor_260821_083556_by_gpt-5_child04_p4c_graph_persistence_review1.md`, `15,261` bytes / `101` LF / SHA-256 `AA94C417A02BE371168D8ADCD5B3F4FECF1941F382518EFF9214EC0FABB93CFF` |
| `E4-P4C-DETECT1` | pre-commit Anvien detect-changes result | recorded — exit `0`; `180` changed semantic units, `16` changed files, `14` affected files, HIGH risk; resolution health `0/0/0`, semantic fields complete |
| `E4-P4C-COMMIT1` | isolated accepted slice commit | pending |

#### `E4-P4C-REVIEW1` — independent Supervisor acceptance

- Verdict: `PASS`.
- Report: `reports/Supervisor/rp_supervisor_260821_083556_by_gpt-5_child04_p4c_graph_persistence_review1.md`.
- Identity: `15,261` bytes / `101` LF / SHA-256 `AA94C417A02BE371168D8ADCD5B3F4FECF1941F382518EFF9214EC0FABB93CFF`; the self-reference-safe canonical hash was independently recomputed and matched.
- Fresh review evidence closes the invariant at `414` Graph Export nodes, `414` File→Export relations, `0` duplicate IDs, `0` orphan local-definition references, `0` compatibility drift, and `11,592` normalized Graph JSON↔Ladybug field comparisons with `0` differences and `0` missing IDs in either direction. Direct runtime records are `240`, runtime local records `239`, and forbidden terminal/resolved-target/public-API keys are `0`.
- Source-level clearance covers `emit.go`, `csv.go`, `schema.go`, and `context.go`; focused owner tests and the canonical full build are PASS. High/critical impacts remain warnings only. Review-induced empty `.tmp\\p4c-tests` and the prior read-only `/root/authority_scan` deviation are recorded as provenance; neither changed repository state or invalidated independence. No target access occurred.

#### `E4-P4C-DETECT1` — Main-owned fresh change detection

- Fresh graph basis after planner refresh: `anvien analyze --force` exit `0`; scanned/parsed/failed `1,857/735/0`; graph `113,523/156,030` nodes/relationships.
- `anvien detect-changes --repo E:\\Anvien --scope all --json` exit `0`; parsed summary: `180` changed semantic units, `16` changed files, `14` affected files, `HIGH` changed-file/overall risk, and `10` affected process entries. Changed app layers are backend `143`, backend-test `13`, docs `24`; changed functional areas are resolution `130`, storage `21`, documentation `24`, unknown `5`.
- Resolution-gap inventory reports `86` changed gap entities/occurrences but persisted health remains clean: `totalResolutionGapCount=0`, `nodesWithGaps=0`, `degradedNodes=0`, and changed source nodes with gaps `0`. `appLayer` and `functionalArea` semantic fields are complete for all `113,523` nodes.
- The exact semantic changed-path set contains the five Child 04 living ledgers plus the tracked P4-C production/test/golden paths. The four new focused tests and untracked Coder/Supervisor/Main provenance reports are outside the semantic changed-file set and are retained only in the explicit staging manifest; no unrelated path is authorized.

### P4-C2 — real-target acceptance

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4C2-ORACLE1` | independent 21-entry direct-export oracle plus negative controls; no copied target source | pending |
| `E4-P4C2-TARGET1` | real target analyze result: direct exports `21/21`, exact fields, access separation, negative controls, and zero Child 05-derived claims | pending |
| `E4-P4C2-BOUNDARY1` | target pre/post source boundary, operational-output record, and contamination check | pending |
| `E4-P4C2-REVIEW1` | independent Supervisor PASS | pending |
| `E4-P4C2-DETECT1` | Anvien-side change/boundary check before evidence commit | pending |
| `E4-P4C2-COMMIT1` | isolated Anvien-side evidence/ledger commit; no target artifact committed | pending |

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-PNA-REVIEW1` | final Supervisor review across all five slices and retained evidence | pending |
| `E4-PNB-CLEAN1` | dead child-created artifacts removed; valid evidence retained | pending |
| `E4-PNB-REVIEW1` | Supervisor PASS for cleanup | pending |
| `E4-PNC-DETECT1` | final detect-changes result for implementation scope | pending |
| `E4-PNC-COMMIT1` | final known commit/worktree state | pending |
| `E4-PNC-HANDOFF1` | Child 05 actual-status refreshed from accepted immutable syntax/direct-export results and exact opening condition recorded | pending |

P4-B detect/commit closure is recorded at `11a37aa8ec0320dd93258c058b088d1070aa778d`. P4-B1 source/build/boundary/Supervisor acceptance, Main-owned detect, and isolated commit `42d167aaf28446ac0b3de479a8afefabb8d06736` are recorded and closed. `E4-P4C-AUTH1` opens only P4-C; P4-C2/Child 05 remain locked.

### `E4-P4B-SOURCE1` — source and cleanup clearance

- Fresh excluded graph refresh before Main-owned change detection: `anvien analyze . --force --exclude 'internal/aicontext/skills/**' --exclude '.claude/skills/**'` exit `0`; scanned/parsed/failed `1,136/626/0`; graph `81,772/121,285` nodes/relationships; forbidden skill trees excluded.
- Exact candidate remains `internal/providers/tsjs/imports.go`, `internal/providers/tsjs/extract.go`, and `internal/providers/tsjs/extract_test.go`; current diff is `1,186` insertions / `10` deletions and no other production/test path.
- `DefinitionFact.Visibility` is unchanged; source-bearing `ImportReexport`/`ImportWildcard` compatibility remains intact; no P4-B1/P4-C/P4-C2/Child 05 state was added.
- Structured unsupported/malformed diagnostics, direct/default/local alias/type-only facts, cardinality, meaning lanes, negative controls, and forward local-definition evidence are covered by the focused matrix and nearest boundary.
- Cleanup resubmission proves `.tmp/p4b_ast_probe/` and all matching probe directories are absent; the prior Supervisor `REJECT` is retained as historical evidence and the resubmission `PASS` closes the artifact lifecycle invariant.

### `E4-P4B-REVIEW1` — independent Supervisor acceptance

- Prior report `reports/Supervisor/rp_supervisor_260821_001125_by_gpt-5_child04_p4b_export_facts.md` recorded the sole cleanup blocker.
- Resubmission report `reports/Supervisor/rp_supervisor_260821_001554_by_gpt-5_child04_p4b_export_facts_resubmission.md` is `PASS`; residual same-invariant surfaces are none.
- `E4-P4B-DETECT1` and `E4-P4B-COMMIT1` are closed at `11a37aa8ec0320dd93258c058b088d1070aa778d`; P4-B was not reopened.

### `E4-P4B1-SRC1` — source, boundary, and Supervisor clearance

- Fresh excluded graph after the REVIEW3 report was present: `1,144/626/0` scanned/parsed/failed and `82,059/121,760` nodes/relationships; exact exclusions were `internal/aicontext/skills/**` and `.claude/skills/**`.
- Candidate production/test paths are only `internal/providers/tsjs/imports.go` and `internal/providers/tsjs/extract_test.go`; current identities, diff, index, formatting, forbidden-field, visibility, and `.tmp` gates are recorded in the Coder and Supervisor reports.
- `recoveredReexportSiblingAfterMalformedAlias` accepts one valid name, anonymous `as`, comma-only error, and valid alias in source order while allowing only comment/parser-extra trivia between semantic tokens; malformed names/errors fail closed. `addSourceExportFact` remains the sole compatibility derivation path.
- Independent Supervisor REVIEW3 is `PASS`; residual same-invariant surfaces are none within P4-B1. Successor verification found no invalidation; `E4-P4C-AUTH1` opens only P4-C, while P4-C2/Child 05 and target access remain locked.

### `E4-P4B1-DETECT1` — Main-owned change detection

- Command: `anvien detect-changes --repo E:\\Anvien --scope all` after the planner refresh and fresh excluded graph.
- Result: final exit `0`; `305` changed semantic units; `7` changed files and `7` affected files; one affected provider process; overall risk `MEDIUM`; changed app layers `backend=219`, `backend_test=61`, `docs=25`; changed functional areas `providers=280`, `documentation=25`.
- Resolution inventory reports changed gap entities but persisted resolution health remains clean: `totalResolutionGapCount=0`, `nodesWithGaps=0`, `degradedNodes=0`; semantic app/functional fields are complete.
- Final excluded graph basis is `1,146/626/0` scanned/parsed/failed with `82,079/121,780` nodes/relationships. Changed path set is the exact two production/test candidate files plus the five Child 04 living plan documents. Untracked reports/provenance remain outside the semantic changed-file set and are staged only in the explicit report/provenance boundary.

### `E4-P4B1-COMMIT1` — isolated slice closure

- Isolated implementation commit: `42d167aaf28446ac0b3de479a8afefabb8d06736` (`feat(tsjs): recover comment-bearing re-export siblings`), with exact P4-B1 production/test boundary and the accepted ledger/report updates.
- Final P4-B1 documentation closure commit: `871189b8c6a4e4bb9ff538407232c913b8cf4db6`; successor verification found a clean worktree/index, branch ahead of origin with no push, both external docs-only commits preserved, and exact accepted source/test hashes unchanged. `E4-P4C-AUTH1` opens only P4-C; P4-C2/Child 05/target access remain locked.

### `E4-P4C-AUTH1` — successor opening authority

- The explicit successor delegation requires campaign continuation after verifying P4-B1; questions/reminders are not pauses, and only an explicit pause/stop halts work.
- Current Git basis is clean HEAD `871189b8c6a4e4bb9ff538407232c913b8cf4db6`. P4-B1 implementation commit `42d167aaf28446ac0b3de479a8afefabb8d06736` and external docs-only commits `84a354940aea8240c99bf4868e721209e7248830` / `ce0e200c55bd96c4374cc6e84bd99a3c82bef641` are ancestors; no reset, checkout, push, or target access occurred.
- `imports.go`, `extract_test.go`, and Supervisor REVIEW3 retain the accepted SHA-256 identities `52AFBA2FC9A7ACEA314D5B39043A054F2671B2EF987161D8F1C3D641B2382749`, `07CF7D49715CA0398DA9485086A37E395FCB8E2E695AF2B649B6BC05074C604D`, and `07DD5BB92F169C5923C0DBCB597F914A28E594496ACDADB55D41F13DE364421C`; `git diff --quiet 42d167aaf28446ac0b3de479a8afefabb8d06736 HEAD` over the two source/test paths exits `0`.
- Rotation handoff `reports/Investigation/rp_main_260821_0333_orchestration_rotation_handoff.md` retains `4,991` bytes / `49` LF / SHA-256 `C02DE5EB447AAADE9F39CC381F7CAA13C1A8B659435EE68F3B6EA7ABAF061226`.
- Fresh excluded self-graph at the successor basis exits `0`: scanned/parsed/failed `1,147/626/0`, graph `82,087/121,788`. Roadmap is LOW with `28` outbound plan links; each changed Child 04 ledger is LOW with one inbound roadmap link; upstream affected files/flows are `0/0`.
- Only P4-C is open. Its first executable gate is `E4-P4C-IMPACT1`: refresh exact production file-detail and upstream symbol/file impact before any code edit. P4-C2, Child 05, target access, terminal resolution, barrels, cycles, ambiguity, and public-API work remain locked.
