# P6-A Declaration-Universe Architect/Planner Handoff

Created: `2026-08-22 03:36:07 +07:00`

Role: visible Child 06, P6-A Architect/Planner lane

Send to: Main task `01a025d5-811e-7d33-a528-10aa6513b06c`, then independent Supervisor

Verdict: `READY_FOR_SUPERVISOR`

Acceptance state: candidate decision only. This lane does not self-accept, stage, commit, or open P6-B/C/D.

## 1. Outcome

P6-A selects a deterministic, declaration-backed TypeScript standard-library authority that is generated offline from the official TypeScript `5.9.3` `lib*.d.ts` corpus, checked in as a versioned compact catalog, and embedded in the Go runtime. Repository resolution remains authoritative and first. The external catalog is consulted only for eligible TypeScript global/type/member sites after repository resolution has produced a non-authoritative miss; it never rescues an explicit-import failure.

The supported configuration boundary is deliberately narrow and fail-closed. P6-B may support only zero or one repository-root `tsconfig.json` in JSONC form and only the declaration-profile inputs `compilerOptions.target`, `compilerOptions.lib`, and `compilerOptions.noLib`. Absent config uses the compiler-default profile recorded by the generated manifest. Invalid values, conflicting inputs, `extends`, `references`, nested or multiple project configs, `files`, `include`, `exclude`, `jsconfig.json`, or any source-ownership ambiguity produce a structured `capability_unavailable` outcome for the affected TypeScript declaration lookup. They are not guessed, executed, or silently ignored, and P6 does not change the scanner's source corpus.

P6-C1 is `PRESERVE_ONLY`: current evidence proves no campaign requirement for project/package `.d.ts`, ambient-module, augmentation, or `node_modules` lookup. P6-C2 is `ACTIVE`: a resolved declaration needs a dedicated referenced-only `ExternalSymbol` identity because the current Ladybug adapter drops unknown node labels and relationships whose endpoints are unsupported, while current context/impact/rename/process readers otherwise misinterpret or omit the fact.

## 2. Authority and boundary basis

- Authoritative checkout: `E:\Anvien` only.
- Baseline/HEAD: exact Child 05 closure `ec765debff335540c77d409ebb2c9f45e4a0a77d`.
- Accepted predecessors: Pn-B `fcc44334c0f75b3b19046dc8f9f4de40eb459fa9`, Pn-A `b68e738d64eebea65a045afbf0b12d94dd43cbf4`, and P5-A/B/C/D commits named in the Child 05 handoff.
- Child 05 invariants preserved: repository module/file lookup stays separate from syntax-derived export tables and deterministic export traversal; explicit-import failure is fail-closed; terminal/hop/failure proof remains generic-first; fixed-corpus physical/resolver-emitted/persisted `IMPORTS` stays `5,072 / 5,072 / 5,088`.
- The fifteen untracked Main orchestration handoffs were pre-existing protected work and were not read or modified.
- `E:\cheapapp.org` remained locked and was not read, analyzed, or written.
- No production, test, fixture, target, package, lock, process configuration, or Child 05 file was edited.

The requested first durable handoff deadline was `2026-08-22 03:30:00 +07:00`. At checkpoint continuation the clock was already `03:34:32 +07:00` and no report existed. This report therefore records a deadline miss instead of claiming timely delivery. There is no dependency mismatch or unexpected tracked/index drift; Supervisor/Main must decide whether the late handoff is acceptable.

## 3. Fresh graph and current runtime evidence

Before any graph query, the lane ran `anvien analyze --force` at the authoritative root. It scanned `1,985` files, parsed `743`, failed `0`, produced `116,467` nodes and `160,591` relationships, indexed commit `ec765debff335540c77d409ebb2c9f45e4a0a77d`, and reported `stale=false`.

Current source proves:

- `internal/analyze/analyze.go` passes repository ScopeIR into resolution and has no TypeScript declaration-profile reader.
- `internal/resolution/types.go` exposes only two compatibility booleans; it has no catalog/config/authority outcome input.
- `internal/resolution/indexes.go:buildWorkspace` accepts only `[]scopeir.ScopeIR` and indexes repository facts.
- `internal/resolution/resolve.go` resolves repository scope/import/member facts before emitting a generic unresolved diagnostic.
- `internal/resolution/emit.go:emitUnresolvedReference` and `internal/graphhealth` carry no authority/catalog/config/stage/reason/candidate contract.
- Graph-health still classifies target text using Go-oriented builtin/standard-library/external qualifier tables instead of consuming resolver truth.
- `anvien/package.json` ships only `bin` and `go-src`; packaged runtime does not ship Node, TypeScript, or `node_modules` and has lifecycle package scripts.
- `anvien-web/package-lock.json` resolves TypeScript `5.9.3`, integrity `sha512-jl1vZzPDinLr9eUt3J/t7V6FgNEw9QjvBPdysz9KfQDD41fQrC2Y4vKQdiaUpFT4bXlb1RHhLpp8wtm6M5TgSw==`, license Apache-2.0. `package.json` currently declares a range, so P6-B must make the generator's exact-version gate explicit instead of treating the range as provenance.

## 4. TypeScript corpus and oracle evidence

The locally installed TypeScript `5.9.3` corpus contains `100` `lib*.d.ts` files totaling `3,141,835` bytes. Observed profiles were:

| Profile | Declaration files | Bytes |
| --- | ---: | ---: |
| all local `lib*.d.ts` inputs | 100 | 3,141,835 |
| `target ES2022` default profile | 63 | 2,389,540 |
| explicit `lib ES2015` | 13 | 310,653 |
| explicit `lib ES5` | 3 | 232,957 |

Meaning lanes are required, not optional metadata: `Math.max`/`Math.min` are declared in `lib.es5.d.ts`; `Promise` type declarations are visible in the ES5 lane while the runtime `Promise` value comes from `lib.es2015.promise.d.ts`.

The bounded/general compiler differential used the exact local `5.9.3` compiler with no network, install, or package script:

| Vector | Compiler result |
| --- | --- |
| compiler default; Promise value/Map/Set/Symbol value | fail |
| `target ES2022`; Promise/Map/Set/Symbol value | pass |
| explicit `lib ES5`; ES2015 values | fail |
| explicit `lib ES2015`; ES2015 values | pass |
| `ES5 + ES2015.Promise`; Promise vs Map/Set/Symbol values | Promise pass; Map/Set/Symbol fail |
| `noLib`; global types | fail |
| invalid `lib` | compiler-option error |
| compiler default; type-only Promise and Math members | pass |
| explicit `lib ES5`; type-only Promise and Math members | pass |
| `noLib`; type-only Promise and Math members | fail |

This is `10/10` decision vectors matching the selected semantics. It is design evidence, not an Anvien implementation pass.

## 5. Mechanism comparison

| Mechanism | Determinism/provenance | Runtime/package/security | Performance/maintenance | Decision |
| --- | --- | --- | --- | --- |
| Hardcoded builtin/name map | cannot retain declaration source, profile, meaning, or member lineage | small but would encode target names | drifts manually and violates generality | reject |
| Invoke Node/`tsc` or compiler API during analyze | TypeScript-authoritative per run | Node/TypeScript absent from shipped runtime; process execution and package surface expand | startup/process cost and operational dependency | reject |
| Discover/install declarations or execute package scripts at runtime | environment-dependent and non-reproducible | network/package execution violates P6-A boundary | unbounded I/O and attack surface | reject |
| Embed raw `lib*.d.ts` and parse them on every analyze | source provenance possible | no Node needed if custom Go parsing exists | repeats ~3.14 MB parse, duplicates TypeScript semantic merging/profile logic, higher maintenance | reject |
| Reuse repository ScopeIR as if libs were repository files | blurs external/repository ownership and imports | packageable | pollutes graph/indexes and does not preserve compiler profile/meaning faithfully | reject |
| Offline TypeScript-API generation -> checked-in versioned compact catalog -> `go:embed` runtime lookup | exact compiler version, ordered inputs/hashes, schema and output hash can be verified | no runtime Node/network/install/script; catalog ships inside current Go artifact | one deterministic generation cost; bounded startup/lookup; explicit regeneration on TS upgrade | select |

## 6. Selected catalog contract

P6-B owns a new isolated `internal/tsstdlib` package and a developer-only generator. The generator may use the already locked TypeScript compiler API only as an offline build-time authority. Runtime code must not discover `node_modules`, invoke Node/`tsc`, install packages, access the network, or execute package scripts.

The generated DTO is versioned and contains only facts required by supported lookup:

- canonical symbol identity and display name;
- `value`, `type`, and `namespace` meanings;
- owner/member identity and member meaning;
- source `lib*.d.ts`, declaration range, and profile membership;
- default profiles by supported `target`, explicit `lib` dependency closure, and `noLib` behavior;
- TypeScript version, package-lock integrity, ordered input paths and SHA-256 values, generator schema version, generation command identity, and catalog SHA-256.

The checked-in catalog is immutable at runtime. Loader validation is fail-closed: missing asset, unsupported schema, version mismatch, input-manifest mismatch, or catalog-hash mismatch makes the authority unavailable and yields a reasoned capability outcome. It never falls back to a hand-authored map.

Identity must be deterministic from authority kind, TypeScript version, catalog hash, canonical declaration library/source range, semantic owner path, name, and meaning. The same catalog/profile/input must generate byte-identical catalog output and identical external IDs across at least two clean generation runs.

## 7. Supported configuration contract

P6-B supports exactly:

1. no `tsconfig.json`: use the generated TypeScript `5.9.3` compiler-default declaration profile;
2. one repository-root `tsconfig.json` parsed as data-only JSONC;
3. within that root file, only `compilerOptions.target`, `compilerOptions.lib`, and `compilerOptions.noLib` affect the catalog profile.

Rules:

- Explicit `lib` replaces the target-derived default library set according to the generated manifest.
- Absent `lib` selects the default library closure for the accepted target/default.
- `noLib: true` disables the declaration authority for global lookup.
- Unknown/invalid target or lib, invalid JSONC, conflicting compiler options, unreadable config, or unsupported catalog mapping is `capability_unavailable`, never a generic repository gap.
- `extends`, `references`, nested/multiple `tsconfig.json`, `files`, `include`, `exclude`, `jsconfig.json`, and any source-to-project ownership ambiguity are unsupported topology. Presence produces `capability_unavailable` for affected TypeScript declaration lookup; P6 does not infer inheritance, execute config code, select scanner files, or merge projects.
- Non-TypeScript sites never consult this authority. Explicit import failures never consult it. Project/package declaration lookup remains outside this contract.

The narrow topology is a deliberate minimum based on current evidence. Expanding it requires a later planner update, new oracle vectors, owner/impact refresh, and Supervisor acceptance before code.

## 8. Lookup precedence and outcome contract

The immutable stage order is:

```text
repository lexical/binding/import/export traversal
  -> authoritative repository resolution or fail-closed explicit-import failure
  -> eligible TypeScript global/type/member declaration lookup
  -> resolved external target or structured unavailable/unresolved outcome
```

External lookup eligibility requires TypeScript language plus a global/type/member source-site meaning. Member lookup occurs only after an eligible external receiver has resolved. A plain name match across the wrong meaning lane is `meaning_mismatch`, not success. External resolution does not emit a synthetic `IMPORTS` edge.

The minimum final status set is:

- `resolved_internal`
- `resolved_external`
- `unresolved`
- `capability_unavailable`

Each affected site carries schema version, final status, resolution stage, requested name, requested meaning, language, target identity when resolved, reason when not resolved, proof/candidates when applicable, and authority/catalog/config hashes. Minimum reason codes are `disabled_by_no_lib`, `profile_excludes_declaration`, `meaning_mismatch`, `config_invalid`, `config_topology_unsupported`, `config_unreadable`, `catalog_missing`, `catalog_schema_unsupported`, `catalog_version_mismatch`, and `catalog_hash_mismatch`.

P6-C3 should first use a versioned concrete JSON DTO in the existing `graph.Evidence.Note`/diagnostic `Note` carriage where that preserves the contract, rather than expanding shared graph structs speculatively. If fresh implementation impact proves typed fields necessary for lossless projection, the plan must be refreshed before that shared edit. One source site has one final outcome; an earlier authoritative failure cannot be overwritten by a later name guess.

## 9. External representation and reader contract

P6-C2 is active. A successfully referenced external declaration is a dedicated `ExternalSymbol` node, not a repository `File`, Definition-like node, `Export`, `Package`, or `ResolutionGap`. It is referenced-only: materialize only targets reached by accepted source sites, never all catalog declarations.

Required properties are deterministic ID, name/qualified name, meanings, authority kind, TypeScript version, catalog hash, declaration library, declaration source range, and external origin. Resolved CALLS/USES/ACCESSES/type-reference relationships retain source-site and outcome proof. Capability-unavailable cases remain outcome/gap facts and do not create fake external targets. No synthetic File/DEFINES/IMPORTS or repository ownership edge is allowed.

Affected consumer decisions:

| Consumer | Required P6 behavior |
| --- | --- |
| graph JSON | generic node/relationship transport should preserve the new label/properties; validate before editing `internal/graph/types.go` |
| Ladybug CSV/schema | add explicit `ExternalSymbol` table/columns and allowed relationship endpoint pairs; otherwise current exporter skips the node/pair |
| `internal/processes` | exclude external endpoints from process call adjacency/terminal discovery so external calls do not become product process steps or terminals |
| MCP context | project external identity/provenance explicitly; do not present a fake repository path/range |
| MCP impact | expose external identity/proof without treating it as an editable upstream repository symbol |
| MCP rename | reject `ExternalSymbol` before edit collection, fail-closed with a stable non-editable reason |
| graph-health / ResolutionGap | consume the structured outcome in P6-D; never infer authority from target text |
| HTTP/Web generic transport | preserve generic unknown-label fallback; no P6-A evidence justifies UI redesign |
| file-detail | preserve unless later impact proves semantic interpretation rather than generic graph transport |

## 10. Exact owner map for later slices

No file below is authorized for edit until its owning slice runs a fresh graph refresh, full `file-detail`, and upstream symbol impact. New-file names are architecture ownership boundaries; P6-B may adjust a path only through a planner refresh before code.

### P6-B — active standard-library authority

- `anvien-web/scripts/generate-tsstdlib-catalog.mjs` — new offline TypeScript-API generator; no package script entry or runtime use.
- `internal/tsstdlib/catalog.v1.json` — new checked-in deterministic generated DTO.
- `internal/tsstdlib/catalog.go` — new `go:embed`, hash/schema validation, immutable lookup.
- `internal/tsstdlib/profile.go` — new root JSONC declaration-profile reader and supported profile selection.
- `internal/analyze/analyze.go` — minimum repo-root/profile construction and resolver wiring.
- `internal/resolution/types.go` — typed authority/profile input only; existing compatibility options preserved.
- `internal/resolution/indexes.go` — keep repository workspace separate; attach external lookup index without importing it as repository ScopeIR.
- `internal/resolution/resolve.go` — bounded global/type/member hooks after immutable repository results and before generic unresolved emission.
- `anvien-web/package.json` and `anvien-web/package-lock.json` — validate and, only if required for reproducible generation, pin the generator's exact TypeScript version/integrity without adding a production runtime dependency.

### P6-C1 — preserve-only

No production owner. Close with evidence that project/package `.d.ts`, ambient modules/augmentations, and `node_modules` scanning remain unimplemented and unavailable; do not repeat Child 05 module/export traversal.

### P6-C2 — active external representation

- `internal/scopeir/kinds.go` — `ExternalSymbol` label.
- new isolated external-symbol materializer under `internal/resolution/` — referenced-only node identity/properties.
- `internal/resolution/emit.go` — emit truthful external relationships/proof and no synthetic imports.
- `internal/lbugload/csv.go` — explicit table columns and lossless CSV projection.
- `internal/lbugschema/schema.go` — table and relationship endpoint pairs.
- `internal/processes/processes.go` — prevent external endpoints from entering process topology.
- `internal/mcp/context.go` — external projection.
- `internal/mcp/impact.go` — non-editable external impact projection.
- `internal/mcp/rename.go` — stable fail-closed rejection before file edit collection.

### P6-C3 — active structured outcomes

- `internal/resolution/types.go` — versioned outcome DTO/status/reason contract.
- `internal/resolution/resolve.go` and `internal/resolution/emit.go` — one final outcome per affected source site and precedence enforcement.
- existing graph evidence/diagnostic note carriers and only lossless affected persistence readers proved by refreshed impact.

### P6-D — active projection/target validation

- `internal/graphhealth/policy.go`, `internal/graphhealth/diagnostics.go`, and the source-backed `ResolutionGap` projection owner — mechanical mapping from structured outcomes after fresh file-detail/impact.
- Exact target access remains locked until P6-D.

Validate/preserve-only unless later impact invalidates the decision: `internal/graph/types.go`, graph JSON writer, generic HTTP/Web transport/fallback display, and file-detail.

## 11. Blast-radius warnings

Fresh P6-A impact evidence is a warning requiring narrow slicing and affected-boundary regression, not an edit prohibition:

- `analyze.Run`: CRITICAL — `24` symbols / `9` files / `8` modules / `23` processes.
- `buildWorkspace`: CRITICAL — `59 / 24 / 8 / 23`.
- `resolveCall`, `resolveAccess`, `resolveTypeAnnotation`: CRITICAL — each `28 / 11 / 7 / 32`.
- `emitUnresolvedReference`: CRITICAL — `9 / 4 / 2 / 34`.
- `graph.Node`: CRITICAL — `1,717 / 273 / 48 / 428`.
- `scopeir.NodeLabel`: CRITICAL — `1,846 / 291 / 47 / 422`.
- `graphhealth.Diagnostic`: CRITICAL — `70 / 17 / 5 / 77`.
- `nodeCSVRow`: CRITICAL — `25 / 11 / 1 / 12`.
- `NodeSchema`: MEDIUM — `8 / 6 / 2 / 0`.
- `classifyDiagnostic`: HIGH — `10 / 4 / 3 / 0`.
- `SourceBackedResolutionGapInputs`: CRITICAL — `13 / 7 / 1 / 15`.
- `processes.Apply`: CRITICAL — `29 / 8 / 7 / 31`; `buildCallsGraph`: CRITICAL — `28 / 8 / 6 / 25`.
- `contextSymbolPayload`: CRITICAL — `1 / 1 / 1 / 7`.
- `impactItemPayload`: CRITICAL — `11 / 5 / 2 / 17`.
- `collectRenameChanges`: CRITICAL — `1 / 1 / 1 / 8`.

Related-file counts recorded by fresh `file-detail` are: analyze `187`; resolution types `34`; indexes `58`; resolve `60`; emit `50`; graph types `256`; scopeir kinds `250`; Ladybug CSV `23`; schema `22`; graph-health policy `31`; diagnostics `30`; ResolutionGap inputs `41`; processes `14`; MCP context `28`; MCP impact `41`; MCP rename `9`; HTTP graph `22`; filecontext `46`.

## 12. Validation and measurement gates for later slices

P6-B must record:

- exact generator inputs, TypeScript version/integrity, ordered hashes, catalog byte size/counts/hash;
- two clean generation outputs with byte equality;
- embedded runtime catalog load/validation behavior;
- authority initialization time, lookup latency/throughput, peak memory, and packaged binary size delta using a recorded baseline protocol;
- all ten P6-A compiler vectors as independently authored general fixtures, including meaning-lane and language-isolation cases;
- full build before boundary validation, package-runtime verification without runtime Node, affected regressions, Supervisor PASS, change detection, and isolated commit.

No arbitrary performance ceiling is accepted in P6-A. P6-B records measured baseline/final values; any threshold requires Owner decision.

P6-C2 must record referenced target count, graph/Ladybug node and relationship parity, skipped-node/skipped-relationship delta `0` for external facts, external/repository ownership pollution `0`, process pollution `0`, and context/impact/rename behavior. P6-C3/P6-D must prove outcome-field parity, resolved/unresolved overlap `0`, and mechanical health mapping.

## 13. Build and boundary disposition

A full build was not run in P6-A. The repository package lifecycle includes install/postinstall/postpack scripts, while the current lane explicitly forbids package installation and package-script execution. Running a partial command and calling it a full build would be false evidence. `E6-P6A-BUILD1` is therefore `N/A for this decision slice / not run by boundary`, with the fresh real CLI `anvien analyze --force` recorded as current non-UI runtime evidence. Every production slice still requires exact lock/process preflight, full build, and nearest real affected boundary.

The two local `.tmp\p6a-oracle` source probes were debug-only and are removed before handoff. Their results are retained as summarized decision evidence in the living evidence ledger; no `.tmp` artifact is an official evidence dependency.

## 14. Open risks and stop conditions

- Exact TypeScript upgrade policy beyond `5.9.3` is not authorized. A version change requires regenerated provenance, oracle rerun, benchmark comparison, plan refresh, and review.
- Config topology beyond the supported root subset remains unavailable; do not silently broaden it for target convenience.
- `ExternalSymbol` is high/critical shared-graph scope. C2 must remain referenced-only and must validate every named reader before edit.
- Versioned JSON in `Evidence.Note` is the minimum carrier decision, not permission to hide data from structured readers. If lossless consumption cannot be achieved, refresh the plan before editing a shared graph struct.
- P6-C1 remains preserve-only unless new current evidence materially invalidates P6-A; activation requires a planner update before code.
- Stop on predecessor invalidation, new tracked/index drift outside the listed P6-A docs/report, catalog/oracle mismatch, unsupported config ambiguity, or any need to read `E:\cheapapp.org` before P6-D.

## 15. Main boundary intervention and disclosed deviation

At `2026-08-22 03:42:26 +07:00`, Main verified that this lane had modified `docs/contracts/graph-accuracy-contract.md` by `18` added lines / `0` deleted lines. That tracked file is outside the sealed P6-A transfer boundary, which permits only the four Child 06 living ledgers and this one report. The contract change is therefore a disclosed scope deviation, not an accepted P6-A deliverable.

Per Main's intervention, this lane stopped every further contract write immediately and did not revert, delete, overwrite, stage, or commit the deviation. Supervisor/Main owns disposition of that existing diff. No further graph, build, test, target, network, install, package-script, or subagent action occurred after the intervention.

The final transfer must inventory the contract path separately from the allowed four ledgers/report. The report deadline miss remains as recorded in section 2. This report remains a non-self-accepted decision candidate.

## 16. Handoff

Independent Supervisor should verify:

1. exact repo/HEAD/predecessor boundary and no protected-handoff mutation;
2. fresh Anvien graph/file-detail/impact evidence and the reported CRITICAL/HIGH counts;
3. source/config/package/consumer claims against current files;
4. local TypeScript `5.9.3` corpus counts/hashes and the `10/10` compiler differential;
5. selected offline catalog mechanism, fail-closed config topology, outcome fields/reasons, P6-C1 preserve-only, P6-C2 active, and exact later-slice owner/reader map;
6. no production/test/fixture/target edit, no build misclaim, `.tmp` cleanup, and the exact four-ledger/report authorized transfer;
7. the recorded `03:30 +07` deadline miss;
8. the existing out-of-boundary `docs/contracts/graph-accuracy-contract.md` `18+/0-` deviation, which this lane neither reverted nor touched again after Main's intervention.

On Supervisor PASS, Main may authorize the P6-A decision commit and only then open P6-B. Without Supervisor PASS, P6-B/C/D remain locked.

Final verdict: `READY_FOR_SUPERVISOR`
