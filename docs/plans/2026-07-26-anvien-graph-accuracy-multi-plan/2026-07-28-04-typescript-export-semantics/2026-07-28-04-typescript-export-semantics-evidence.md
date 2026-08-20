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

All rows below are reserved implementation evidence targets. Status remains pending until the exact command, file/artifact, observed result, Supervisor verdict, and commit are recorded.

### P4-A — export fact, meaning, and visibility boundary

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4A-IMPACT1` | fresh file-detail/impact and one-responsibility ownership for exact editable owners | pending |
| `E4-P4A-SRC1` | smallest production export fact/meaning contract and structured unsupported diagnostic | pending |
| `E4-P4A-BUILD1` | full build after production and focused tests | pending |
| `E4-P4A-TEST1` | deterministic serialization, one-fact-per-specifier, meaning, access separation, and zero-derived-state matrices | pending |
| `E4-P4A-BOUNDARY1` | real production fact round trip and unsupported diagnostic count | pending |
| `E4-P4A-REVIEW1` | independent Supervisor PASS | pending |
| `E4-P4A-DETECT1` | pre-commit Anvien detect-changes result | pending |
| `E4-P4A-COMMIT1` | isolated accepted slice commit | pending |

### P4-B — direct/default/alias/type-only extraction

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4B-IMPACT1` | exact direct-export extraction file/symbol impact | pending |
| `E4-P4B-SRC1` | production direct/default/alias/type-only extraction | pending |
| `E4-P4B-BUILD1` | full build | pending |
| `E4-P4B-TEST1` | direct/default/anonymous/alias/type-only/multi-specifier matrix and negative controls | pending |
| `E4-P4B-BOUNDARY1` | real provider/ScopeIR direct export facts and exact counts | pending |
| `E4-P4B-REVIEW1` | independent Supervisor PASS | pending |
| `E4-P4B-DETECT1` | pre-commit Anvien detect-changes result | pending |
| `E4-P4B-COMMIT1` | isolated accepted slice commit | pending |

### P4-B1 — star/namespace/re-export syntax

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4B1-IMPACT1` | exact re-export syntax owner impact | pending |
| `E4-P4B1-SRC1` | production named/default/star/namespace/type-only re-export syntax extraction | pending |
| `E4-P4B1-BUILD1` | full build | pending |
| `E4-P4B1-TEST1` | syntax matrix with one fact per specifier and no terminal resolution state | pending |
| `E4-P4B1-BOUNDARY1` | real provider/ScopeIR re-export syntax output | pending |
| `E4-P4B1-REVIEW1` | independent Supervisor PASS | pending |
| `E4-P4B1-DETECT1` | pre-commit Anvien detect-changes result | pending |
| `E4-P4B1-COMMIT1` | isolated accepted slice commit | pending |

### P4-C — graph and affected persistence projection

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E4-P4C-IMPACT1` | exact fact-to-graph/persistence blast radius and actual affected-consumer inventory | pending |
| `E4-P4C-SRC1` | smallest production projection/compatibility change with no blanket adapter expansion | pending |
| `E4-P4C-BUILD1` | full build | pending |
| `E4-P4C-TEST1` | graph fact conservation, direct count, access separation, compatibility derivation, negative controls, and zero terminal state | pending |
| `E4-P4C-BOUNDARY1` | real graph/affected persistence output with field differences and orphan counts | pending |
| `E4-P4C-REVIEW1` | independent Supervisor PASS | pending |
| `E4-P4C-DETECT1` | pre-commit Anvien detect-changes result | pending |
| `E4-P4C-COMMIT1` | isolated accepted slice commit | pending |

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

No implementation or closure claim is made in this ledger yet.
