# Child 03 P0-A current-source inventory

Status: `READY_FOR_SUPERVISOR` — QA candidate only, not an acceptance verdict.

Date/time: `2026-08-14 03:27:55 +07:00`
Role: QA/source-of-truth inventory and planner-ledger update, no fix
Repository: `E:\Anvien`
Entry branch/HEAD: `master` / `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0`

## Authority and reading boundary

The following assigned authority was read in full, in order: `AGENTS.md`; `.agents/skills/working-rules/SKILL.md`; `.agents/skills/qa/SKILL.md`; `.agents/skills/planner/SKILL.md`; the graph-accuracy roadmap; all four Child 03 ledgers; `docs/contracts/graph-accuracy-contract.md`; `reports/Investigation/rp_main_260814_020654_child02_pnc_closure_handoff.md`; `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien`; and both assigned Supervisor reports. Current Git, graph evidence, production source, tests and fixture artifacts were then read from the assigned repo only. Historical names and summaries were treated as leads, never current ownership proof.

The working-rules, QA and planner skills controlled the no-fix source-of-truth inventory and the four-ledger update. This lane did not act as coder, Supervisor, target validator, acceptance authority or main orchestrator.

## Scope and non-goals

The only goal was P0-A: replace historical assumptions with current-HEAD evidence for the TypeScript/JavaScript binding pipeline and record `E0-P0A-GRAPH1`, `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, and `E0-P0A-STATUS1`.

No production, test, fixture, target or runtime source was changed. No build, tests, runtime/serve gate, target access, target analyze, detect-changes, staging, commit, branch operation, reset, checkout, stash or push occurred. P3-A and later slices remain closed. Export semantics route to Child 04; module/re-export resolution to Child 05; ambient/external behavior to Child 06; scanner behavior remains outside Child 03.

## Git boundary

Entry verification for the correction found the expected boundary: branch `master`, HEAD `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0`, staged paths `0`, and exactly six dirty paths (the four Child 03 ledgers, this candidate report, and the protected Supervisor REJECT report). Child 02 Pn-C is the entry commit and was not rerun. The Supervisor report was not edited. Staging and commit remain forbidden.

## Graph invocation sequence — E0-P0A-GRAPH1

1. `anvien --help` ran first and exited `0`.
2. The first `anvien analyze --force` wrapper timed out after `5.025s`, wrapper exit `124`, with no stdout. This is not called PASS. Process/lock inspection identified PID `23376` as the live analyze process; the other Anvien processes were MCP. PID `23376` exited on its own, its valid lock released, and meta publication/up-to-date state was observed.
3. The sole complete invocation capture, and the only evidence used for `E0-P0A-GRAPH1`, was `anvien analyze --force`: exit `0`; scanned `1,629`; parsed code `688`; failed `0`; nodes `97,340`; relationships `136,614`; artifact `E:\Anvien\.anvien\graph.json`; indexed/current source commit `181b8cb800f5fe34fa6fe85ddd359f514ead9fb0`.
4. One planned `anvien status` ran after capture: indexed `2026-08-14 02:34:41 +07`, current/up to date. No additional analyze was run.

Correction temporal boundary: the authorized repair pass then ran exactly one fresh `anvien analyze --force` after `anvien --help`; it exited `0` with scanned `1,631`, parsed code `688`, failed `0`, nodes `97,369`, relationships `136,643`, and published the same HEAD `181b8cb8` artifact. This graph includes the already-written documentation/report paths. It is recorded for manifest-repair freshness only; it does not replace the original `E0-P0A-GRAPH1` capture and does not rerun or alter the accepted `27/27` file-detail or `15/15` file-impact gates.

## Frozen universe — E0-P0A-SRC1

The owner universe was frozen before the detail/impact batch and never expanded:

- Set A, full-read source/evidence paths: `62`.
- Set B, unique behavior-owner rows: `27` = `15` production + `12` exact test owners.
- Set C, transparent/excluded evidence paths: `35`.
- Arithmetic: `62 = 27 + 35`; duplicates `0`; assigned mandatory leads `27/27`; unassigned mandatory leads `0`.

A provisional denominator `40` was rejected before the first file-detail because it incorrectly promoted generic contracts/helpers and fixture artifacts into independent owners. It is not used as an evidence denominator.

## Durable Set A/B/C manifest and partition checks

This is the exact sorted path manifest used for the authorized correction pass. Set A is defined mechanically as `B ∪ C`; Set B retains the previously cleared 27 behavior-owner rows verbatim; Set C contains the 35 transparent/excluded/deferred paths that were actually read in full to close the mandatory leads. Path kind, responsibility, mode and route are recorded per row so an independent reviewer can recompute membership without relying on category prose.

| # | Set | Exact repo-relative path | Kind | Responsibility | Mode / route | Source-backed rationale |
|---:|:---:|---|---|---|---|---|
| 1 | C | `internal/graph/snapshot_test.go` | test | generic graph snapshot reader | excluded / Child 02 | Tests graph snapshot persistence, not binding extraction. |
| 2 | C | `internal/graph/types.go` | production model | graph node/relationship types | inspect-only / Child 02 | Generic graph model consumes accepted facts; no binding-specific field. |
| 3 | C | `internal/graphaccuracy/access_candidate.go` | production audit | access-audit candidate selection | inspect-only / deferred | Generic audit surface; does not extract omitted leaves. |
| 4 | C | `internal/graphaccuracy/graphaccuracy.go` | production audit | Go/general graph-accuracy command | excluded / unrelated | Go/general audit path, not TS binding extraction. |
| 5 | C | `internal/graphaccuracy/graphaccuracy_test.go` | test | Go/general graph-accuracy tests | excluded / unrelated | Tests unrelated general graph-accuracy behavior. |
| 6 | C | `internal/graphhealth/compute.go` | production reader | graph-health computation | inspect-only / Child 02 | Consumes generic graph diagnostics, not binding facts directly. |
| 7 | C | `internal/graphhealth/diagnostics.go` | production reader | diagnostic classification/counting | inspect-only / Child 02 | Reads existing diagnostics; no extraction-diagnostic contract. |
| 8 | C | `internal/graphhealth/report.go` | production reader | graph-health reporting | inspect-only / Child 02 | Reports generic health data already emitted downstream. |
| 9 | C | `internal/graphhealth/resolution_gap_aggregation.go` | production reader | ResolutionGap aggregation | inspect-only / Child 02 | Aggregates existing gaps; cannot recover missing binding leaves. |
| 10 | C | `internal/graphhealth/resolution_gap_aggregation_test.go` | test | ResolutionGap aggregation validation | excluded / Child 02 | Validates generic gap aggregation, not extraction ownership. |
| 11 | C | `internal/graphhealth/resolution_gap_inputs.go` | production reader | ResolutionGap input projection | inspect-only / Child 02 | Consumes generic unresolved-reference inputs. |
| 12 | C | `internal/graphhealth/resolution_gap_inputs_test.go` | test | ResolutionGap input validation | excluded / Child 02 | Tests generic gap input behavior. |
| 13 | C | `internal/graphhealth/resolution_gap_validation.go` | production reader | ResolutionGap policy validation | inspect-only / Child 02 | Validates existing gaps, not binding extraction. |
| 14 | C | `internal/graphhealth/resolution_gap_validation_test.go` | test | ResolutionGap policy tests | excluded / Child 02 | Tests generic gap policy. |
| 15 | C | `internal/providers/astro/extract.go` | bridge | Astro SFC extraction bridge | preserve-only / shared SFC | Thin pass-through into shared SFC/TSJS extraction. |
| 16 | C | `internal/providers/provider_parity_test.go` | test | cross-provider parity | excluded / preserve | Provider parity control; not a TS binding owner. |
| 17 | C | `internal/providers/sfc/extract.go` | bridge | shared SFC extraction | preserve-only / shared SFC | Delegates embedded script extraction; no independent pattern walker. |
| 18 | C | `internal/providers/svelte/extract.go` | bridge | Svelte SFC extraction bridge | preserve-only / shared SFC | Thin pass-through into shared extraction. |
| 19 | B | `internal/providers/tsjs/definition_position_inputs_test.go` | test owner | Definition position contract | future validation / P3-B | Exact test boundary for declaration position inputs. |
| 20 | B | `internal/providers/tsjs/definitions.go` | production owner | Definition/local-binding emission | future edit / P3-B | Identifier gate and accepted Definition/BindingFact emission. |
| 21 | B | `internal/providers/tsjs/extract_test.go` | test owner | TSJS extraction contract | future validation / P3-A/B | Exact extraction behavior test boundary. |
| 22 | B | `internal/providers/tsjs/extract.go` | production owner | collector entry and AST traversal | inspect-only / P3-A | Walks named AST nodes; no recursive pattern contract. |
| 23 | B | `internal/providers/tsjs/imports.go` | production owner | import-binding pipeline | preserve-only / all P3 | Separate import emission must remain exactly once. |
| 24 | B | `internal/providers/tsjs/legacy_p7_type_env_conversion_test.go` | test owner | type-environment conversion control | future validation / P3-B | Exact type-inference independence regression boundary. |
| 25 | B | `internal/providers/tsjs/legacy_scope_captures_test.go` | test owner | scope capture control | future validation / P3-B1/B2/B2A | Exact lexical-scope regression boundary. |
| 26 | B | `internal/providers/tsjs/nodes.go` | production owner | AST node helpers | inspect-only / P3-A | Identifier-like helpers currently stop short of recursive patterns. |
| 27 | B | `internal/providers/tsjs/references.go` | production owner | references/calls/accesses | inspect-only / P3-C | Missing leaves cannot receive lexical reference targets. |
| 28 | B | `internal/providers/tsjs/scope_id_test.go` | test owner | scope identity contract | future validation / P3-B1/B2A | Exact scope-ID behavior boundary. |
| 29 | B | `internal/providers/tsjs/scopes.go` | production owner | lexical scope facts | future edit per context / P3-B1/B2/B2A | No catch scope; loop declaration scope is partial. |
| 30 | B | `internal/providers/tsjs/types.go` | production owner | type-binding inference | inspect-only / P3-B/B1 | Separate inference branch independently rejects patterns. |
| 31 | C | `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json` | fixture | TSJS ScopeIR signature golden | fixture artifact / test owner | Golden belongs to `extract_test.go`; not an owner row. |
| 32 | C | `internal/providers/vue/extract.go` | bridge | Vue SFC extraction bridge | preserve-only / shared SFC | Thin pass-through into shared extraction. |
| 33 | C | `internal/resolution/access_audit.go` | production reader | resolution access audit | inspect-only / deferred | Reads resolved access facts; no binding extraction. |
| 34 | C | `internal/resolution/binding_accumulator.go` | production helper | type-binding accumulator | excluded / type path | Type-binding accumulator, not declaration-pattern extraction. |
| 35 | C | `internal/resolution/definition_collision_test.go` | test | Definition collision control | excluded / resolution | Generic collision behavior; no binding owner. |
| 36 | B | `internal/resolution/emit.go` | production owner | graph occurrence/`DEFINES` projection | inspect-only / P3-C | Projects accepted Definition facts one-for-one. |
| 37 | C | `internal/resolution/import_resolution.go` | production reader | module/import resolution | deferred / Child 05 | Module/re-export semantics belong to Child 05. |
| 38 | B | `internal/resolution/identity_occurrence_test.go` | test owner | occurrence identity contract | future validation / P3-C | Exact graph occurrence validation boundary. |
| 39 | B | `internal/resolution/indexes.go` | production owner | Definition/scope indexes and lookup | inspect-only / P3-C | Indexes accepted facts; missing DefID bindings are skipped. |
| 40 | B | `internal/resolution/p2b_persistence_test.go` | test owner | persistence conservation control | future validation / P3-C | Exact persistence boundary for changed binding facts. |
| 41 | B | `internal/resolution/proof_accuracy_golden_test.go` | test owner | source-site proof control | future validation / P3-C | Exact source-site/resolution proof test boundary. |
| 42 | B | `internal/resolution/resolution_test.go` | test owner | lexical resolution contract | future validation / P3-C | Exact lexical resolution and gap regression boundary. |
| 43 | C | `internal/resolution/source_site.go` | production reader | source-site metadata helpers | inspect-only / P3-C | Generic source-site support; no extraction ownership. |
| 44 | C | `internal/resolution/testdata/p1c_identity_repo/src/identity.ts` | fixture | identity occurrence fixture | fixture artifact / test owner | Fixture belongs to `identity_occurrence_test.go`. |
| 45 | C | `internal/resolution/testdata/typescript_graph_signature.golden.json` | fixture | graph signature golden | fixture artifact / test owner | Golden belongs to `resolution_test.go`. |
| 46 | C | `internal/resolution/type_alias_test.go` | test | type-alias resolution control | excluded / resolution | Generic type-alias behavior, not binding extraction. |
| 47 | C | `internal/resolution/types.go` | production helper | resolution data types | inspect-only / deferred | Generic resolver types; no changed binding contract proven. |
| 48 | C | `internal/scopeir/definition_index.go` | production helper | Definition index helper | inspect-only / ScopeIR | Indexes existing definitions; no recursive extraction. |
| 49 | C | `internal/scopeir/kinds.go` | production contract | node/fact kind constants | preserve-only / ScopeIR | Existing labels consumed by accepted facts; no pattern data. |
| 50 | C | `internal/scopeir/position_index.go` | production helper | position lookup index | inspect-only / ScopeIR | Supports ranges/positions after extraction; not owner itself. |
| 51 | C | `internal/scopeir/range.go` | production contract | source range type | preserve-only / ScopeIR | Existing range primitive reused by future facts. |
| 52 | C | `internal/scopeir/scope_tree.go` | production helper | lexical scope tree | inspect-only / P3-B/C | Builds scope relationships; catch/loop behavior remains future context work. |
| 53 | B | `internal/scopeir/facts.go` | production owner | Definition/Binding/Scope contracts | future edit / P3-A | BindingFact exists but lacks pattern metadata/diagnostics. |
| 54 | B | `internal/scopeir/ir.go` | production owner | ScopeIR storage/normalization | future edit / P3-A | No pattern/diagnostic collection currently exists. |
| 55 | B | `internal/scopeir/scopeir_test.go` | test owner | ScopeIR deterministic contract | future validation / P3-A | Exact ScopeIR normalization/golden boundary. |
| 56 | B | `internal/scopeir/sort_keys.go` | production owner | deterministic fact ordering | future edit / P3-A | Pattern/diagnostic facts need deterministic sort keys. |
| 57 | C | `internal/scopeir/testdata/scopeir.golden.json` | fixture | ScopeIR golden | fixture artifact / test owner | Golden belongs to `scopeir_test.go`. |
| 58 | B | `internal/graphaccuracy/property_access.go` | production owner | Property audit | inspect-only / P3-C audit | Reads emitted Property nodes; cannot recover omitted leaves. |
| 59 | B | `internal/graphaccuracy/property_access_test.go` | test owner | Property audit contract | future validation / P3-C | Exact audit reader boundary. |
| 60 | B | `internal/graphaccuracy/source_site_accuracy.go` | production owner | source-site audit | inspect-only / P3-C audit | Reads generic unresolved diagnostics only. |
| 61 | B | `internal/graphaccuracy/source_site_accuracy_test.go` | test owner | source-site audit contract | future validation / P3-C | Exact source-site audit test boundary. |
| 62 | B | `internal/resolution/resolve.go` | production owner | resolution orchestration and gaps | inspect-only / P3-C | Resolves present facts; cannot reconstruct omitted leaves. |

Partition checks were run against the exact sorted lists above (PowerShell set operations, path strings only): `|A|=62`, `|B|=27`, `|C|=35`, duplicate path count `0`, `|B∩C|=0`, `|A\(B∪C)|=0`, `|(B∪C)\A|=0`, mandatory Child 03 leads assigned `27/27`, unassigned `0`. The observed machine-readable result was `A=62; B=27; C=35; Overlap=0; Missing=0`.

Canonical sorted path manifest (the per-row table above supplies the metadata for each entry):

```text
internal/graph/snapshot_test.go
internal/graph/types.go
internal/graphaccuracy/access_candidate.go
internal/graphaccuracy/graphaccuracy.go
internal/graphaccuracy/graphaccuracy_test.go
internal/graphaccuracy/property_access.go
internal/graphaccuracy/property_access_test.go
internal/graphaccuracy/source_site_accuracy.go
internal/graphaccuracy/source_site_accuracy_test.go
internal/graphhealth/compute.go
internal/graphhealth/diagnostics.go
internal/graphhealth/report.go
internal/graphhealth/resolution_gap_aggregation.go
internal/graphhealth/resolution_gap_aggregation_test.go
internal/graphhealth/resolution_gap_inputs.go
internal/graphhealth/resolution_gap_inputs_test.go
internal/graphhealth/resolution_gap_validation.go
internal/graphhealth/resolution_gap_validation_test.go
internal/providers/astro/extract.go
internal/providers/provider_parity_test.go
internal/providers/sfc/extract.go
internal/providers/svelte/extract.go
internal/providers/tsjs/definition_position_inputs_test.go
internal/providers/tsjs/definitions.go
internal/providers/tsjs/extract.go
internal/providers/tsjs/extract_test.go
internal/providers/tsjs/imports.go
internal/providers/tsjs/legacy_p7_type_env_conversion_test.go
internal/providers/tsjs/legacy_scope_captures_test.go
internal/providers/tsjs/nodes.go
internal/providers/tsjs/references.go
internal/providers/tsjs/scope_id_test.go
internal/providers/tsjs/scopes.go
internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json
internal/providers/tsjs/types.go
internal/providers/vue/extract.go
internal/resolution/access_audit.go
internal/resolution/binding_accumulator.go
internal/resolution/definition_collision_test.go
internal/resolution/emit.go
internal/resolution/identity_occurrence_test.go
internal/resolution/import_resolution.go
internal/resolution/indexes.go
internal/resolution/p2b_persistence_test.go
internal/resolution/proof_accuracy_golden_test.go
internal/resolution/resolution_test.go
internal/resolution/resolve.go
internal/resolution/source_site.go
internal/resolution/testdata/p1c_identity_repo/src/identity.ts
internal/resolution/testdata/typescript_graph_signature.golden.json
internal/resolution/type_alias_test.go
internal/resolution/types.go
internal/scopeir/definition_index.go
internal/scopeir/facts.go
internal/scopeir/ir.go
internal/scopeir/kinds.go
internal/scopeir/position_index.go
internal/scopeir/range.go
internal/scopeir/scope_tree.go
internal/scopeir/scopeir_test.go
internal/scopeir/sort_keys.go
internal/scopeir/testdata/scopeir.golden.json
```

The rejected provisional denominator `40` consisted of the retained 27 B owners plus these 13 rows that were removed from the owner denominator and finally classified C: `internal/providers/sfc/extract.go`, `internal/providers/vue/extract.go`, `internal/providers/svelte/extract.go`, `internal/providers/astro/extract.go`, `internal/providers/provider_parity_test.go`, `internal/scopeir/definition_index.go`, `internal/scopeir/kinds.go`, `internal/scopeir/position_index.go`, `internal/scopeir/range.go`, `internal/scopeir/scope_tree.go`, `internal/resolution/binding_accumulator.go`, `internal/resolution/import_resolution.go`, and `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json`. Each is present once in Set C with the path-level rationale above; the remaining 22 C rows were transparent/excluded/deferred sibling paths read during the authorized reconstruction pass.

Reproduction recipe (run from `E:\Anvien`, no graph mutation): place the exact `B` and `C` path columns from this table into sorted PowerShell arrays, then evaluate `$A=@($B+$C|Sort-Object -Unique); $A.Count; $B.Count; $C.Count; @($B|Where-Object{$C -contains $_}).Count; @($A|Where-Object{(-not($B -contains $_))-and(-not($C -contains $_))}).Count`. The observed output is `62`, `27`, `35`, `0`, `0`; checking `@($B+$C).Count-$A.Count` yields duplicate count `0`.

### Set B production owners (15)

| Path | Responsibility | Current behavior/classification | Mode | Future slice |
|------|----------------|---------------------------------|------|--------------|
| `internal/providers/tsjs/extract.go` | collector entry and AST traversal | visits nodes by kind; no pattern contract/walker | inspect-only for P3-A | P3-A |
| `internal/providers/tsjs/nodes.go` | AST node helpers | identifier-like helpers do not recursively enumerate patterns | inspect-only; edit only if re-proved | P3-A |
| `internal/providers/tsjs/definitions.go` | Definition/local-binding emission | identifier declarators emit; non-identifier declarators silently return | future edit | P3-B |
| `internal/providers/tsjs/scopes.go` | lexical scope facts | function/module candidates; no catch scope; loop identifier can land in enclosing scope | future edit per context | P3-B1/B2/B2A |
| `internal/providers/tsjs/references.go` | references/calls/accesses | references traverse independently; missing leaves supply no target binding | inspect/future narrow edit | P3-C or context slice only if proven |
| `internal/providers/tsjs/types.go` | type-binding inference | annotated identifier parameters and identifier variables can bind types; patterns independently early-return | inspect/future narrow edit | P3-B/B1 |
| `internal/scopeir/facts.go` | DefinitionFact/BindingFact/ScopeFact contracts | `BindingFact` exists, but pattern path/rest/default/provenance and extraction diagnostic contract are absent | future edit | P3-A |
| `internal/scopeir/ir.go` | ScopeIR storage/normalization | current IR has no pattern/diagnostic collection | future edit | P3-A |
| `internal/scopeir/sort_keys.go` | deterministic fact ordering | no pattern/diagnostic sort key | future edit | P3-A |
| `internal/providers/tsjs/imports.go` | separate import-binding pipeline | current imports emit separately | preserve-only | all P3 |
| `internal/resolution/indexes.go` | Definition/scope binding indexes and lexical lookup | indexes accepted Definition/Binding facts; missing DefID binding is silently skipped; inner-to-parent lookup fails ambiguity | inspect-only | P3-C |
| `internal/resolution/emit.go` | graph occurrence/`DEFINES` and unresolved diagnostics | one accepted Definition projects; no separate legal-fact drop observed | inspect-only | P3-C |
| `internal/resolution/resolve.go` | resolution orchestration and gaps | resolves present facts; cannot reconstruct omitted leaves | inspect-only | P3-C |
| `internal/graphaccuracy/property_access.go` | existing Property audit | heuristic-classifies emitted Property nodes; cannot see omitted leaves | inspect-only | P3-C audit |
| `internal/graphaccuracy/source_site_accuracy.go` | existing resolved-edge/unresolved diagnostic audit | reads existing generic `unresolved_reference` diagnostics; no extraction-diagnostic reader | inspect-only | P3-C audit |

### Set B exact test owners (12)

`internal/providers/tsjs/extract_test.go`; `definition_position_inputs_test.go`; `legacy_scope_captures_test.go`; `legacy_p7_type_env_conversion_test.go`; `scope_id_test.go`; `internal/scopeir/scopeir_test.go`; `internal/resolution/identity_occurrence_test.go`; `p2b_persistence_test.go`; `proof_accuracy_golden_test.go`; `resolution_test.go`; `internal/graphaccuracy/property_access_test.go`; and `source_site_accuracy_test.go`.

These are owner rows only because each is an exact validation boundary for a Child 03 contract. Their fixture/golden inputs are artifacts of those owners, not separate owner rows: `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json`, `internal/scopeir/testdata/scopeir.golden.json`, `internal/resolution/testdata/typescript_graph_signature.golden.json`, and `internal/resolution/testdata/p1c_identity_repo/src/identity.ts`.

### Set C explicit exclusions/deferred evidence paths (35)

Set C includes the four fixture artifacts above plus current-source bridge/copier/generic-consumer paths read to close mandatory leads. The classification is by responsibility:

- Vue, Svelte and Astro thin extractors plus the shared SFC bridge: transparent/preserve-only pass-through into `tsjs.Extract`, not separate Child 03 binding owners.
- Generic graph types/builders, graphhealth, ResolutionGap inventory, Ladybug/JSON persistence and generic MCP/CLI/UI readers: Child 02 or generic consumers; they consume already-existing nodes/diagnostics, not a binding-specific changed field.
- `internal/providers/tsjs/binding_accumulator.go`: type-binding accumulator, not a declaration-pattern owner.
- Go-only graphaccuracy paths and historical baseline-count golden inputs: unrelated or preserve-only.
- Export owners: deferred to Child 04. Module/re-export resolution: Child 05. Ambient/external: Child 06. Scanner paths: excluded.
- `E:\cheapapp.org`: forbidden in P0-A and untouched.

No generic helper, bridge, persistence copier, fixture data or reader was promoted to the owner denominator. If a later implementation changes a concrete field they directly consume, that later slice must re-prove ownership before expansion.

## Current behavior classifications

- Variable declarations: plain identifiers are correct and preserve-only; all object/array pattern names are rejected before a DefinitionFact or binding-pattern-specific local-binding record with path/provenance can be created. The existing `BindingFact` type itself is not claimed absent.
- Nested object/array, aliases, rest, defaults, computed keys and holes: no recursive leaf enumeration exists. Computed/unsupported syntax also has no structured extraction diagnostic.
- Parameters: annotated identifiers may produce a separate TypeBindingFact, but function/method/constructor/arrow parameters do not produce a parameter-specific DefinitionFact/local-binding record with binding-pattern metadata, even for a plain identifier.
- Catch: no binding handler and no catch scope.
- `for-of`/`for-in`: declaration-form plain identifiers can be seen as nested variable declarators but use the enclosing function/module scope; declaration patterns reject; assignment forms do not create declarations. The assignment-form non-declaration invariant is established, but plain-identifier writes and assignment-destructuring write/reference behavior are not implemented or fully proven.
- Type inference: Definition emission is not gated by inference success; the branches are separate. Both independently return for non-identifier variable patterns.
- Imports: separate pipeline; preserve exactly once. Export behavior is Child 04.
- Projection: every accepted Definition is indexed and emitted as a graph node with `DEFINES`; no additional omission of legal accepted Definition facts was found.
- Resolution: lexical lookup proceeds from the inner scope through parents; more than one same-name candidate in a scope is ambiguous/fails. A missing extracted binding therefore remains a missing Definition/binding and causes downstream unresolved truth/gaps.
- Diagnostics/audits: unsupported extraction paths are silent. Current graphaccuracy audits see generic graph nodes/unresolved diagnostics only and cannot detect/recover omitted binding leaves.

## File-detail table — E0-P0A-FD1

Every command was exactly `anvien file-detail <path> --repo E:\Anvien --json`. All `27/27` exited `0`, stale `false`, changed-since-analyze `false`, and reported indexed/current HEAD `181b8cb8`. Columns are symbols / inbound / outbound / local / unresolved / related files / risk.

| Path | Metrics |
|------|---------|
| `tsjs/extract.go` | 34 / 19 / 30 / 27 / 66 / 23 / high |
| `tsjs/nodes.go` | 52 / 73 / 26 / 5 / 68 / 20 / high |
| `tsjs/definitions.go` | 51 / 11 / 42 / 7 / 93 / 16 / high |
| `tsjs/scopes.go` | 35 / 8 / 44 / 11 / 81 / 17 / high |
| `tsjs/references.go` | 33 / 5 / 35 / 3 / 70 / 15 / high |
| `tsjs/types.go` | 74 / 5 / 61 / 5 / 131 / 17 / high |
| `scopeir/facts.go` | 120 / 874 / 18 / 112 / 4 / 239 / medium |
| `scopeir/ir.go` | 36 / 942 / 28 / 60 / 118 / 237 / high |
| `scopeir/sort_keys.go` | 64 / 243 / 76 / 23 / 63 / 234 / high |
| `tsjs/imports.go` | 29 / 5 / 26 / 1 / 52 / 15 / high |
| `resolution/indexes.go` | 319 / 180 / 95 / 371 / 439 / 49 / high |
| `resolution/emit.go` | 108 / 39 / 166 / 63 / 310 / 42 / high |
| `resolution/resolve.go` | 61 / 83 / 118 / 18 / 128 / 53 / high |
| `graphaccuracy/property_access.go` | 177 / 10 / 40 / 162 / 214 / 6 / high |
| `graphaccuracy/source_site_accuracy.go` | 267 / 10 / 45 / 380 / 237 / 6 / high |
| `tsjs/extract_test.go` | 109 / 13 / 80 / 47 / 0 / 21 / low |
| `tsjs/definition_position_inputs_test.go` | 60 / 0 / 50 / 7 / 0 / 13 / low |
| `tsjs/legacy_scope_captures_test.go` | 33 / 1 / 40 / 7 / 0 / 14 / low |
| `tsjs/legacy_p7_type_env_conversion_test.go` | 13 / 0 / 26 / 1 / 0 / 20 / low |
| `tsjs/scope_id_test.go` | 18 / 0 / 15 / 0 / 0 / 9 / low |
| `scopeir/scopeir_test.go` | 51 / 0 / 32 / 7 / 0 / 5 / low |
| `resolution/identity_occurrence_test.go` | 67 / 0 / 41 / 5 / 0 / 16 / low |
| `resolution/p2b_persistence_test.go` | 23 / 0 / 24 / 1 / 0 / 15 / low |
| `resolution/proof_accuracy_golden_test.go` | 91 / 0 / 60 / 12 / 0 / 24 / low |
| `resolution/resolution_test.go` | 304 / 31 / 192 / 56 / 0 / 45 / low |
| `graphaccuracy/property_access_test.go` | 21 / 0 / 11 / 6 / 0 / 3 / low |
| `graphaccuracy/source_site_accuracy_test.go` | 27 / 0 / 9 / 3 / 0 / 3 / low |

## File-impact table — E0-P0A-IMPACT1

Every final row used `anvien impact file <path> --repo E:\Anvien --direction upstream --json`. All `15/15` were found, stale `false`. Metrics are impact count / affected files / direct / linked flows / linked tests / severity.

| Path | Metrics |
|------|---------|
| `tsjs/extract.go` | 21 / 10 / 10 / 1 / 5 / CRITICAL |
| `tsjs/nodes.go` | 41 / 6 / 40 / 2 / 3 / CRITICAL |
| `tsjs/definitions.go` | 22 / 19 / 7 / 0 / 3 / CRITICAL |
| `tsjs/scopes.go` | 6 / 1 / 6 / 0 / 4 / MEDIUM |
| `tsjs/references.go` | 1 / 1 / 1 / 1 / 3 / LOW |
| `tsjs/types.go` | 3 / 1 / 3 / 0 / 3 / LOW |
| `scopeir/facts.go` | 286 / 74 / 98 / 0 / 94 / CRITICAL |
| `scopeir/ir.go` | 113 / 29 / 48 / 2 / 93 / CRITICAL |
| `scopeir/sort_keys.go` | 17 / 3 / 14 / 2 / 92 / HIGH |
| `tsjs/imports.go` | 1 / 1 / 1 / 1 / 3 / LOW |
| `resolution/indexes.go` | 108 / 11 / 51 / 28 / 26 / CRITICAL |
| `resolution/emit.go` | 25 / 5 / 19 / 12 / 16 / CRITICAL |
| `resolution/resolve.go` | 14 / 5 / 9 / 22 / 29 / CRITICAL |
| `graphaccuracy/property_access.go` | 29 / 2 / 28 / 12 / 1 / HIGH |
| `graphaccuracy/source_site_accuracy.go` | 37 / 4 / 33 / 1 / 1 / CRITICAL |

Distribution: CRITICAL `9`, HIGH `2`, MEDIUM `1`, LOW `3`; missing `0`; stale `0`. Three multi-path wrappers timed out technically; their completed rows were retained only when already fully returned, and missing paths were subsequently run to a final successful result. `indexes.go` had one single-command timeout before a successful result. No path remained after two timeouts without a result.

## Canonical symbol impact

The exact frozen list is `39` symbols, reconciled as TS/JS `15` + ScopeIR `11` + import `2` + resolution/projection `7` + audit readers `4` = `39`. The provisional count `38` stated before canonical execution was an arithmetic error; it did not change the owner denominator `27`.

Canonical targets:

- TS/JS: `extract.go::{Request,Extract,walkKind}`; `nodes.go::{firstIdentifierLikeChild,isIdentifierLike,stripTypeAnnotation}`; `definitions.go::{emitDefinitionKind,addDefinition}`; `scopes.go::{collectScopeCandidateForKind,buildScopes,innermostScopeID}`; `references.go::emitReferenceKind`; `types.go::{emitTypeBindingKind,emitVariableTypeBinding,addTypeBinding}`.
- ScopeIR: `facts.go::{DefinitionFact,BindingFact,ScopeFact}`; `ir.go::{ScopeIR,Normalized,NormalizeInPlace,NormalizeOwned,MarshalDeterministic,Unmarshal}`; `sort_keys.go::{compareDefinition,compareBinding}`.
- Import: `imports.go::{emitImportKind,addImport}`.
- Resolution/projection: `indexes.go::{buildWorkspace,resolveName,resolveScopedName}`; `emit.go::{emitDefinitionNodes,emitUnresolvedReference}`; `resolve.go::{Resolve,ResolveInto}`.
- Audit readers: `property_access.go::{buildPropertyAccessAudit,classifyTSJSProperty}`; `source_site_accuracy.go::{buildSourceSiteAccuracy,sourceSiteDiagnosticsFromNode}`.

For every final row, the command was `anvien impact symbol <exact UID> --uid <exact UID> --repo E:\Anvien --direction upstream --json`; the UID came from file-detail for the exact owning path. Results: completed/found `39/39`; exit nonzero `0`; stale `0`; ambiguous `0`; not-found `0`; unresolved `0`; CRITICAL `14`, LOW `25`, HIGH/MEDIUM `0`.

Rejected attempts are disclosed but are not `E0-P0A-IMPACT1` evidence: exploratory bare names produced ambiguity (`Extract`, `addDefinition`, `emitDefinitionKind`); exploratory `collector.*` names returned not-found; one empty-UID invocation caused `accepts at most 1 arg(s), received 2`; and two exploratory wrapper batches timed out after partial output. No first match was selected. Only the final exact-UID `39/39` batch is symbol proof.

The largest symbol blast radii were `DefinitionFact` CRITICAL `181`, `BindingFact` CRITICAL `171`, `ScopeFact` CRITICAL `173`, and `ScopeIR` CRITICAL `112`. `Extract` and `walkKind` were CRITICAL `8` and `10`; `buildWorkspace` `8`; `resolveName` `17`; `resolveScopedName` `14`; `emitDefinitionNodes` `6`; and `emitUnresolvedReference` `7`. These are scope warnings that justify separating the future walker from context wiring and keeping generic consumers inspect-only.

## P3-A boundary decision — E0-P0A-STATUS1

The old P3-A text was not exact because it referred generically to “fact and TS walker owners” even though no current recursive walker or pattern/diagnostic contract exists. The plan now names the candidate boundary:

- future edit: one focused `internal/providers/tsjs/binding_patterns.go` recursive walker;
- future edit only as required: `internal/scopeir/facts.go`, `ir.go`, `sort_keys.go` for deterministic pattern path/rest/default/computed-key/hole/provenance and structured unsupported diagnostics;
- inspect-only in P3-A: current collector/traversal/declaration/scope/reference/type owners;
- preserve-only: established identifier-only behavior, the separate import pipeline, and the assignment-form non-declaration invariant; assignment write/reference behavior remains partial/missing;
- deferred: variable, parameter, catch and loop wiring to P3-B/B1/B2/B2A; graph projection/resolution to P3-C; target acceptance to P3-C2.

Final P0 candidate (non-acceptance note): the owner boundary is resolved and all five evidence IDs are recorded. Established preservation is limited to identifier/import behavior and the assignment-form non-declaration invariant; assignment write/reference behavior remains partial/missing and routes to P3-B2A/P3-C. This lane does not self-PASS. P3-A remains closed until main orchestration obtains independent Supervisor PASS and records an isolated P0-A commit.

## Changed-path manifest and artifacts

Permitted intended changes only:

1. `docs/plans/.../2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
2. `docs/plans/.../2026-07-28-03-typescript-binding-pattern-extraction-evidence.md`
3. `docs/plans/.../2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md`
4. `docs/plans/.../2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md`
5. `reports/QA/rp_qa_260814_032755_by_gpt-5_child03_p0a_inventory.md`

Analyzer-owned operational artifact: `E:\Anvien\.anvien\graph.json`, refreshed at the correction temporal boundary and still indexed at entry HEAD; it is normal repo-local output and not a source edit. No `.tmp` directory or retained debug artifact was created.

## Blockers and handoff

Unresolved owner boundary: none. Build/runtime/target blocker: none encountered because those gates were forbidden. Acceptance remains intentionally unperformed.

`READY_FOR_SUPERVISOR`. Next owner: main orchestration for a fresh independent Supervisor re-review of this evidence repair. If accepted, main orchestration may later record the isolated P0-A commit. Until independent PASS and commit exist, P0-A remains unchecked and P3-A remains closed.
