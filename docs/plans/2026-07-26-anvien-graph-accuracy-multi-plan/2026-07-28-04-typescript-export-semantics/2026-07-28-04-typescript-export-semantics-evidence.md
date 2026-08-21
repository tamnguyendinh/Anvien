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
- Target source remains inspect-only and must never be copied into an Anvien plan/report artifact. P4-C2 source-only Oracle Authoring may derive expected facts from the authorized target source before analyzer observation; no target-side report may be used as a plan artifact.
- Every evidence-bearing artifact must be created directly in its durable repo-approved `reports/QA/child04-p4c2/...` path. `.tmp` is disposable debug-only; an oracle, raw capture, command stream, manifest, provenance record, expected-values input/output, benchmark, or reproducibility artifact created there is invalid and cannot be promoted/restored/copied/renamed or close an evidence ID.
- Expected-value authorship and analyzer/QA validation are separate lanes. The source-only author may not run/read analyzer output or use current implementation/tests/goldens/QA output to fill expected fields. QA receives a durable, hash-sealed, committed bundle and may compare actual output without revising it.

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

P4-A/P4-B/P4-B1/P4-C/P4-C2 retain their accepted isolated boundaries. Aggregate `E4-PNA-REVIEW1` and cleanup `E4-PNB-REVIEW1` are independently `PASS`; `E4-PNB-DETECT1` and `E4-PNB-COMMIT1` are closed at `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`. Pn-C closes the Child 04 plan by the current exact five-document commit boundary.

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
| `E4-P4C-COMMIT1` | isolated accepted slice commit | recorded — `c99c4070b66e7a96be8c9fa2721a0335a1f94877`; post-commit HEAD matches and `git diff --check` passes |

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

#### `E4-P4C-COMMIT1` — isolated slice closure

- Commit: `c99c4070b66e7a96be8c9fa2721a0335a1f94877` (`feat(graph): project TypeScript export facts`). The staged 23-path manifest contained the five refreshed Child 04 living ledgers, four production owners, four focused tests, seven owner test/golden updates, Coder/Supervisor reports, and the current `0819` Main handoff.
- Post-commit verification: `git rev-parse HEAD` matches the commit; `git diff --check` passes; no push, reset, checkout, target access, or Child 05 action occurred. Only the two preserved older Main handoffs (`0631`, `0721`) remain untracked because their historical blank EOF lines were intentionally not rewritten.
- P4-C is complete. P4-C2 may now open with its independent 21-entry oracle and target pre-state gate; Child 05 and later slices remain locked.

### P4-C2 — real-target acceptance

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4C2-ORACLE1` | clean-context source-only bundle created directly at `reports/QA/child04-p4c2/oracle/<oracle_id>/`: accepted source basis, exactly 21 positive + 11 negative rows, complete schema/provenance, valid bundle digest/seal, zero target writes, zero forbidden observations, zero `.tmp` evidence lifecycle violations | recorded — `SEALED`; oracle ID `p4c2-oracle-v1-a869876ab626-260821_110849+0700`; bundle digest `7749AB14E02FBF61CF7F81B3A7638888F8AD414064F9F969CFED06EB11355439`; `21+11`; zero forbidden lifecycle events |
| `E4-P4C2-TARGET1` | existing QA continuation verifies oracle seal/source basis, reuses the valid identical-byte build, runs one fresh normal built analyzer, and records row-level `21/21`, negative `11/11`, exact fields, access/compatibility separation, zero diagnostics/orphans/Child 05-derived claims in a durable run bundle | recorded — post-repair QA `READY_FOR_SUPERVISOR`; positives `21/21`; negatives `11/11`; FileContext `17/3/1`; parity `588/0`; zero diagnostics/orphans/Child 05 state |
| `E4-P4C2-BOUNDARY1` | durable target pre/post manifests prove unchanged HEAD, source hashes, and pre-existing non-`.anvien` worktree state; only normal analyzer-owned `.anvien` output changes; no target report/fixture/probe/debug artifact | recorded — post-repair target HEAD/branch/seven tracked modifications/three source hashes and four Git-config identities preserved; only normal `.anvien` output changed; process-local trust cleared |
| `E4-P4C2-REVIEW1` | independent later Supervisor PASS over source basis, lane independence, seal chronology/identity, row comparison, and target boundary | recorded — `REJECT`; exact 15-TypeAlias compatibility/reader invariant confirmed; report `8E37F4B1...81F7`; repair and re-review required |
| `E4-P4C2-REVIEW2` | independent resubmission review proves the rejected invariant and same-invariant siblings close on post-repair bytes | recorded — `PASS`; report `5B99A74B...8DC0B`; residual same-invariant surfaces none |
| `E4-P4C2-DETECT1` | Anvien-side change/boundary check after Supervisor PASS and before evidence commit; exact durable evidence + five-ledger manifest only | recorded — exit `0`; fresh graph `1,939/736/0`, `114,628/157,443`; `50` changed units; exact `7/7` changed/affected tracked files; changed-file HIGH/overall LOW; zero affected processes/flows and persisted health `0/0/0` |
| `E4-P4C2-COMMIT1` | isolated Anvien-side durable oracle/QA evidence and ledger commit; no target or `.tmp` artifact committed or required for reproducibility | recorded — `03f09b43f652b9a14b3e49774dc805c0dfd24a27`; parent `310502a88849fe75f86a45a987ba21490d19dbe2`; exact 89 paths; post-commit index/tracked worktree clean; no push |

#### P4-C2 durable oracle lifecycle correction

- Broken invariant: AGENTS says `.tmp` is debug-only and working-rules says temporary scripts/output are not official evidence. Owner authority further makes this apply from artifact creation time to every oracle, raw capture, manifest, provenance record, benchmark, expected-value input/output, and reproducibility artifact.
- Exact contradiction retained as historical record: `reports/QA/rp_qa_260821_091213_by_gpt-5_child04_p4c2_oracle_gate.md:19,110,182` calls the deleted `.tmp` capture an accepted raw oracle and accepts restore/supply as the unblock route; `reports/Investigation/rp_investigation_260821_0932_by_gpt-5_p4c2_oracle_recovery.md:172-176,204` makes a provenance-bound backup/recovery of that capture a minimum input and continuation condition. Those methods conflict with `AGENTS.md:20` and `.agents/skills/working-rules/SKILL.md:124-126`.
- Correct classification: `.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle-output.json` was an `invalid-lifecycle debug capture`, never an accepted oracle. Its absence is not a recovery blocker, and restore/promote/copy/rename cannot make its bytes authoritative. Both durable blocker reports remain unchanged as historical records only.
- Correct pipeline: the already-authorized source-only Oracle Author writes and seals exactly `21+11` expected rows directly under `reports/QA/child04-p4c2/oracle/<oracle_id>/`; existing QA then consumes that bundle read-only for full-build/analyze/compare and writes its run bundle under `reports/QA/child04-p4c2/runs/<run_id>/`; existing Supervisor/detect/commit closure follows. Analyzer/QA output cannot author or amend expected values.
- Exact schema, paths, digest, and lane boundary are fixed by Architect candidate `reports/Architect/rp_architect_260821_103812_by_gpt-5_p4c2_evidence_lifecycle.md`, verified at `42,936` bytes / `553` LF / SHA-256 `1E7EEB6DD83F05384BF43ED9216C042C535DAC5E776B56032A17E3D4288BDEEA`; this is `READY_FOR_PLANNER`, not Supervisor acceptance.
- Planner handoff is `reports/Planner/rp_planner_260821_104059_by_gpt-5_p4c2_durable_oracle_gate_correction.md`; next owner is source-only Oracle Authoring with `working-rules` and `Data-Integrity`.
- The correction itself did not access/stat/hash/read/write `E:\cheapapp.org`; subsequent clean-context Oracle Authoring closed `E4-P4C2-ORACLE1` as `SEALED` without changing that historical correction record.

#### `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, and `E4-P4C2-BOUNDARY1` — durable target validation handoff

- Sealed oracle bundle: `reports/QA/child04-p4c2/oracle/p4c2-oracle-v1-a869876ab626-260821_110849+0700/`; digest `7749AB14E02FBF61CF7F81B3A7638888F8AD414064F9F969CFED06EB11355439`; exactly `21` positive and `11` negative rows; target basis HEAD `a869876ab6262dacde6cd5d432d099a91852a646`; zero target writes, forbidden observations, or evidence-bearing `.tmp` artifacts.
- Historical first QA run remains a durable Git-trust `BLOCKED` run. Its canonical full build PASS and target preservation evidence were referenced, not repeated. The retry used process-local `safe.directory` only, changed no Git config file, and executed exactly one target analyze: exit `0`, `77.37s`, `1,359/887/0`, graph `94,422/125,299`.
- Durable QA report: `reports/QA/child04-p4c2/runs/p4c2-target-validation-retry-260821_115050+0700/p4c2-qa-retry-validation-report.md`; `11,342` bytes / `200` LF / SHA-256 `C831004F049A563A2387B599BE01C943F5B9416C72B1C45E50A8C1F9D2CEFDB4`. Comparison SHA-256 is `2C78AB3BDF67D857E5C2A1B75B0F1FDFBFEBE2B70D92E7EBC8EB45A0AC5A3F27`; 19-file run digest is `9F414A2C54C42F4E39AD8ED03DC042CCC3E1FB5993DA842B22F64851D16AABC4`.
- QA finding: all `21/21` direct Export facts are present with exact semantic fields and separated access, but `P001`–`P014` and `P018` TypeAlias definitions retain `isExported=false` and FileContext `exported=false`; only six Function rows pass overall. Negative controls pass `11/11`; Graph JSON↔Ladybug compares `588` fields with `0` differences; duplicate IDs, relationship/local-definition orphans, diagnostics, and forbidden Child 05 state are `0`.
- `E4-P4C2-TARGET1` and `E4-P4C2-BOUNDARY1` are evidence-complete with a semantic finding, not accepted. Target source/worktree/config are preserved. Exactly one independent Supervisor lane now owns `E4-P4C2-REVIEW1`; P4-C2 and Child 05 remain open/locked respectively until its verdict and the correct repair-or-closure handoff.

#### `E4-P4C2-REVIEW1` — independent rejection

- Verdict: `REJECT`. Report: `reports/Supervisor/rp_supervisor_260821_123030_by_gpt-5_child04_p4c2_durable_retry_review1.md`; `15,671` bytes / `126` LF / self-reference-safe canonical SHA-256 `8E37F4B126ABB38A5DCE071C35937F59D9152EFA135B9245408CA138BEC781F7`.
- Supervisor independently recomputed the oracle and 19-file run digests, aggregated the exact failed row set, inspected validator/source logic, and confirmed the 15 target occurrences are direct exported TypeAlias declarations. No closed gate was rerun.
- Exact source cause: `internal/resolution/emit.go` uses runtime eligibility to populate `directExportDefIDs`; `exportFactIsRuntime` excludes TypeOnly/type-meaning facts, so definition `isExported` becomes false and `internal/filecontext/context.go` deterministically propagates false. `internal/resolution/p4c_export_projection_test.go` codifies the rejected expectation.
- Accepted evidence remains closed: oracle seal; all 21 Export facts/relations and semantic fields; six Function positives; `11/11` negatives; `588/588` parity; zero duplicate/orphan/diagnostic/Child 05 state; target/config preservation. Only the definition compatibility and affected-reader invariant is rejected.
- Next owner is one bounded Coder repair lane. It must change production first, preserve type-only/meaning/access semantics and negatives, then return through build, fresh target QA, re-review, detect, and commit. P4-C2 stays open; Child 05 stays locked.

#### Post-repair Coder, QA, and `E4-P4C2-REVIEW2`

- Exact repair candidate is `internal/resolution/emit.go` (`+1/-15`, SHA-256 `B1B8F22A...B2867`) plus test-after-code `internal/resolution/p4c_export_projection_test.go` (`+23/-6`, SHA-256 `F575F16D...E0162`). Canonical full build, focused rejection test, nearest resolution→FileContext→Ladybug boundary and all four affected packages pass. Coder report: `reports/Coder/rp_coder_260821_131129_by_gpt-5_child04_p4c2_typealias_compatibility_repair_ready_for_qa.md`, canonical SHA-256 `2B3A9B04...997F6`.
- Fresh durable QA run: `reports/QA/child04-p4c2/runs/p4c2-post-repair-validation-260821_131750+0700/`; exactly one target analyze passes `1,359/887/0`, graph `94,422/125,299`; comparison passes `21/21`, `11/11`, FileContext `17/3/1`, parity `588/0`, integrity/Child 05 zeros and preservation. QA report SHA-256 `607A5EC0...248F`; comparison SHA-256 `7FA58C69...CFD5`; 18-file run digest `5CA04508...33E`.
- Independent `E4-P4C2-REVIEW2` verdict is `PASS`: `reports/Supervisor/rp_supervisor_260821_133927_by_gpt-5_child04_p4c2_typealias_compatibility_review2.md`, `13,650` bytes / `120` LF / canonical SHA-256 `5B99A74B1A8D91D48F5E62F0BA1FFCB26317BF818AC6AE044E6CD650B208DC0B`. REVIEW1 remains historical rejection; only detect/commit remain open.

#### `E4-P4C2-DETECT1` — Main-owned fresh graph and change detection

- Preflight found the repo analyze lock free and only editor-owned MCP processes; no process was terminated. `anvien analyze --force --json` then exited `0` exactly once with `1,939/736/0` scanned/parsed/failed and `114,628/157,443` nodes/relationships.
- `anvien detect-changes --repo E:\Anvien --scope all` exited `0`: `50` changed semantic units, exactly `7` changed files and `7` affected files, changed-file risk `HIGH`, overall risk `LOW`, and no affected process entries. Changed layers are `backend=3`, `backend_test=13`, `docs=34`; changed functional areas are `resolution=16`, `documentation=34`.
- The semantic changed/affected path set is exactly the five living documents plus `internal/resolution/emit.go` and `internal/resolution/p4c_export_projection_test.go`. The two production source-site gap changes are `fact.LocalDefID` and `struct{}`; persisted resolution health remains `totalResolutionGapCount=0`, `nodesWithGaps=0`, `degradedNodes=0`, and both required semantic fields cover `114,628/114,628` nodes.
- HIGH is a blast-radius warning for the accepted `emit.go` boundary, not an edit or commit prohibition. Exact report/evidence provenance remains outside the semantic changed-file set and is admitted only by the explicit staging manifest.
- Exact staging is `89/89`: seven tracked repair/living-document paths plus 82 P4-C2 lifecycle/evidence paths; tracked unstaged count is `0`. The only remaining untracked paths are preserved P4-C handoffs `0631` and `0721`. Default staged diff-check over source/test/living documents exits `0`; accepted hash-sealed QA TSV/build-log artifacts retain intentional trailing whitespace, so the precedent-aligned full-manifest command `git -c core.whitespace=-trailing-space diff --cached --check` exits `0` without rewriting evidence bytes or changing Git config.

#### `E4-P4C2-COMMIT1` — isolated slice closure

- Commit `03f09b43f652b9a14b3e49774dc805c0dfd24a27` (`fix(graph): preserve type-only export compatibility`) has parent `310502a88849fe75f86a45a987ba21490d19dbe2` and exactly `89` paths: the accepted two-file repair, five living documents, 82 durable P4-C2 evidence/lifecycle/provenance paths.
- Post-commit verification proves HEAD match, empty index, zero tracked unstaged paths, worktree/index diff-check exit `0`, unchanged repair identities, and exactly two protected untracked P4-C handoffs (`0631`, `0721`). Neither excluded path is in the commit; no push/reset/checkout occurred.
- P4-C2 is complete. Aggregate Child 04 Pn-A opens automatically; Pn-B/Pn-C and Child 05 remain locked.

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-PNA-REVIEW1` | final Supervisor review across all five slices and retained evidence | recorded — independent `PASS`; report `23,563` bytes / `178` LF / canonical SHA-256 `7EBFD508...DC5A8`; residual same-invariant surfaces none |
| `E4-PNB-CLEAN1` | dead child-created artifacts removed; valid evidence retained | recorded — exactly `.tmp/p4c-tests` removed (`0` files / `0` bytes); executor report canonical SHA-256 `5BD0338A...B260` |
| `E4-PNB-REVIEW1` | Supervisor PASS for cleanup | recorded — independent `PASS`; report canonical SHA-256 `EDFCF6CA...1E1F`; residual same-invariant surfaces none |
| `E4-PNB-DETECT1` | proportionate Main-owned change detection for the cleanup/report/ledger boundary | `PASS`: `19/5/5`, LOW/LOW, zero processes/flows/gap/health impact |
| `E4-PNB-COMMIT1` | isolated Pn-B cleanup/report/ledger commit | accepted at `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`, exact eight-path manifest, clean tracked/index post-state |
| `E4-PNC-CLOSE1` | Child 04 plan closure declaration | recorded — all P4, aggregate, cleanup, and Pn-B gates closed; no Child 04 gate remains open |
| `E4-PNC-COMMIT1` | final known commit/worktree state | exact isolated five-document plan-closure boundary; final hash and clean post-state reported externally after Git success |

P4-B detect/commit closure is recorded at `11a37aa8ec0320dd93258c058b088d1070aa778d`. P4-B1 source/build/boundary/Supervisor acceptance, Main-owned detect, and isolated commit `42d167aaf28446ac0b3de479a8afefabb8d06736` are recorded and closed. P4-C closes at `c99c4070b66e7a96be8c9fa2721a0335a1f94877`; P4-C2 closes at `03f09b43f652b9a14b3e49774dc805c0dfd24a27`; aggregate Pn-A and cleanup review are closed; Pn-B detect/commit closes at `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`. Pn-C closes the plan by the exact five-document boundary.

### `E4-PNA-REVIEW1` — independent aggregate acceptance

- Verdict: `PASS`. Report: `reports/Supervisor/rp_supervisor_260821_142429_by_gpt-5_child04_pna_aggregate_review1.md`; identity `23,563` bytes / `178` LF / raw SHA-256 `634555CB9B6917A37EA2826302BC4C8949920B754307EF661C1D4DA3EBC277BC` / canonical SHA-256 `7EBFD5087F8593660A94E70B0816A7FC98944FDE7B9D1F1BC9388CEB9F6DC5A8`.
- Fresh review self graph is `1,939/736/0`, `114,630/157,445`. Current production blobs match accepted slice boundaries; P4-C2 is the exact intentional successor delta. All required commits are ancestors of current HEAD.
- Oracle and QA manifests/digests recompute exactly; aggregate target results are `21/21`, `11/11`, TypeAliases `15/15`, parity `588/0`, and zero duplicate/orphan/diagnostic/Child 05 state with target source/worktree/config preserved.
- No closed build/QA/target/review gate was rerun because source and artifact identities show no invalidation. Residual unverified same-invariant surfaces are none. Pn-B is the next ordered gate; Pn-C and Child 05 remain locked.

### `E4-PNB-CLEAN1` and `E4-PNB-REVIEW1` — cleanup acceptance

- Executor candidate: `reports/coder/rp_coder_260821_144325_by_gpt-5_child04_pnb_cleanup_ready_for_supervisor.md`; `24,399` bytes / `472` LF / raw SHA-256 `0209C39BE833312100DFA3948B9676A8AC091A40286F94F9E2E1220B3278839C` / canonical SHA-256 `5BD0338A8949B58933988FAA6DF448EFBF5B4F4D506C91D6DD7B5B44B8F7B260`.
- Exact cleanup: remove only the review-induced empty `.tmp/p4c-tests` parent (`0` files / `0` bytes). Current source plus four provenance locations prove ownership and deadness; the path was never tracked or committed and supplied no acceptance evidence.
- Retention proof: all `136` tracked Child 04 paths remain; P4-C2 is `89/89` with its 84 non-living paths unchanged; Oracle and three QA generations have exact file sets, zero byte/LF/SHA mismatches and exact digests; the no-missed report sweep is `129` paths with `0/108` known paths missed.
- Independent verdict: `E4-PNB-REVIEW1 PASS` at `reports/Supervisor/rp_supervisor_260821_150355_by_gpt-5_child04_pnb_cleanup_review1.md`; `14,130` bytes / `160` LF / raw SHA-256 `F114AF17513B56952B81B71110BA1C4D838AD40CE3D4E6049D4AD7C2D0ABD18F` / canonical SHA-256 `EDFCF6CACA23DE0F8F38BA4376A25009B33B525BDAF70D267D062155AEC91E1F`. Residual same-invariant surfaces: none.
- No build, QA, target, or closed acceptance gate was rerun because cleanup changed no tracked/code/evidence bytes.

### `E4-PNB-DETECT1` and `E4-PNB-COMMIT1` — isolated cleanup closure

- Fresh final candidate graph was `1,943/736/0`, `114,723/157,538`. `anvien detect-changes --repo E:\Anvien --scope all` exited `0`: `19` changed documentation sections, exactly `5` changed files and `5` affected files, LOW changed-file and overall risk, `0` affected processes/flows, ResolutionGap delta `0/0`, persisted health `0/0/0`, and complete semantic fields over all `114,723` nodes.
- Commit `d1d8eb9002ce9c449c3713de0837ac8216d17a8d` (`docs(plan): close Child 04 cleanup`) has parent `c7997886a0faeb32b7cfe05b4f7d08e38fc57228` and exactly eight paths: five Child 04 living documents, cleanup Coder report, cleanup Supervisor report, and `reports/Investigation/rp_main_260821_1436_orchestration_rotation_handoff.md`.
- Post-commit verification found an empty index, zero tracked unstaged paths, `git diff --check` PASS, and only the protected `0631`/`0721` handoffs untracked before the current rotation handoff was created. No push/reset/checkout occurred. Pn-C opened only after this boundary succeeded.

### `E4-PNC-CLOSE1` — Child 04 plan closure

- All five P4 slices retain accepted isolated commits; aggregate and cleanup reviews are `PASS`; Pn-B is committed at `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`; no Child 04 gate remains open.
- Pn-C changes only the roadmap and four Child 04 ledgers. `E4-PNC-COMMIT1` is this exact five-document plan-closure commit, with final hash and clean post-state reported externally after Git success.

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
