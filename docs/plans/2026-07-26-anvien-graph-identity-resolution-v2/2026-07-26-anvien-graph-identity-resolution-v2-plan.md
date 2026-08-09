# Anvien Graph Identity and TypeScript Resolution Correctness v2 Plan

## Metadata

- Date: `2026-07-26`
- Status: `superseded / reference-only / authority moved to the seven-child roadmap on 2026-08-09`
- Active roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-actual-status.md`

## Goal

Correct the five accepted, bounded graph defects in dependency order: graph identity, TypeScript binding-pattern extraction, export semantics, barrel/re-export resolution, and ambient/external resolution with truthful diagnostics. The implementation must introduce production-grade, versioned contracts; preserve all source occurrences; fail closed on identity or index inconsistency; prove canonical field-level parity across exact surfaces `S0`–`S11`; and validate the real target `E:\cheapapp.org` in place without copying or contaminating it.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
- Run Anvien detect-changes before every implementation-slice commit when implementation work was performed.
- For public runtime or UI-facing changes, validate the real user-visible runtime with browser or Playwright evidence.
- For app/runtime validation, full build must include Docker image/container build. If Docker is missing or not run, full build is incomplete.
- Any Playwright validation must target the real built Docker/container runtime. Running Playwright against a host dev server, framework dev mode, mocked server, or source-run shortcut is not valid runtime evidence.
- If the Docker runtime cannot be built or started, the slice/plan is blocked; do not replace it with dev-server Playwright evidence.
- Playwright evidence must record the Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
- Keep the standard planner structure. These detail rules only make phase checklist items concrete enough to implement safely.
- Every implementation phase must be decomposed into multiple implementation slices that are as small as practical. A phase is a grouping and ordering container; a slice is the executable implementation unit.
- Do not implement a phase directly. Work starts from a slice ID such as `P1-A`, `P1-B`, or `P2-C`.
- Prefer many narrow slices over one broad slice. A single-slice implementation phase is allowed only when the plan explicitly states why the phase cannot be split further without creating empty or non-executable slices.
- Each implementation slice must include:
  + Goal
  + Scope Boundary
  + Non-Goals when useful
  + Pre-flight Questions
  + Work Steps
  + Implementation Gate
  + Acceptance
  + Evidence Targets
  + Actual-status Update
  + Commit Boundary
- Split planned work into separate slices when it contains more than one primary user-visible behavior, user trigger, render location, permission or visibility rule, DB write target, DB state transition, API/CLI/MCP contract, async/event/webhook flow, external side effect, cleanup/quarantine domain, behavior test target, independent acceptance gate, or independent commit boundary.
- Hidden fallback is forbidden. Prefer a visible failure over a fallback that hides a broken primary path.
- When touching DB-backed content, verify the full loop when applicable: UI input -> submit action -> DB write -> DB read after reload/new request -> correct UI render or omission. If there is no UI, replace UI steps with the real caller/consumer flow.
- Tests must prove product behavior. Delete or replace tests that only assert implementation details, helper output, static DOM existence, or mocked plumbing without proving trigger -> process -> observable result.
- If a planned item uses wording such as `and`, `also`, `then wire`, `plus update`, `both`, or `handle all`, check whether it is actually multiple slices.
- Do not write broad actionable items such as `Implement checkout, webhook, entitlement update, and billing UI`; split them into narrow slices such as `Create checkout session request`, `Persist checkout session state`, `Handle provider webhook`, `Update entitlement from webhook event`, and `Render billing status from entitlement`.
- Each slice work step must include UI flow, DB/data flow, render location, and evidence target checks. Use `N/A` with a reason when a check does not apply.
- Mini QA is mandatory for every completed implementation slice and is inherited by every work step in this plan. For Codex, use the applicable Browser, Chrome, Computer Use, or Playwright plugin against the real mounted runtime when the slice has a browser-visible or installed-app boundary; for a non-UI library/storage slice, record those controls as `N/A` with the reason and exercise the nearest real CLI, MCP, loader, persistence, or resolver boundary after the full build. Other agents must use the equivalent browser/session/computer-control or Playwright-like capability exposed by their runtime. Mini-QA commands, target runtime, observed result, and evidence ID must be recorded before Supervisor review.
- If tests write DB rows, app state, files, queues, provider state, or other persistent data, the slice must define cleanup or quarantine before implementation.
- Follow this fixed major order: identity contract and cutover -> binding patterns -> export facts -> module export/re-export resolution -> ambient/external resolution and diagnostic projection.
- Do not start a later slice automatically. The owner must explicitly open each slice; every completed implementation slice is independently built, validated, Supervisor-reviewed, ledger-refreshed, and committed before the next slice opens.
- Before graph-based work, run `anvien analyze E:\Anvien --force`; before editing a target file or symbol, record fresh `file-detail` and upstream `impact`; HIGH/CRITICAL is a blast-radius warning, not an edit ban.
- Production code is implemented before tests are added or updated. After production code is complete, add behavior tests, run the full build, then validate the nearest real boundary.
- Every new or touched production file and test file must own exactly one primary business or semantic responsibility. A file may link to many modules or files only when every link serves that one responsibility.
- Do not add catch-all `utils`, `common`, `helpers`, or `misc` files. If an existing file mixes responsibilities and must be touched, first isolate the relevant responsibility into a dedicated owner file in the same slice or a preceding refactor-only slice; do not append another responsibility.
- Tests follow the same ownership rule: one behavior contract per test file. Shared test construction belongs in a narrowly named package-local test helper only when it has one reusable construction responsibility.
- No unrelated cleanup. Files and sibling modules outside the active slice are inspect-only or preserve-only.
- Scanner pruning for nested directories named `env`, `target`, or `logs` is explicitly out of scope. Its measured `887/895` target File-node baseline is quarantined evidence, not a gate requiring `895/895`; this plan must introduce no additional File-node omissions.
- `E:\cheapapp.org` is analyzed in place. Its graph/index and any operational staging remain inside `E:\cheapapp.org\.anvien`; no target source, fixture, report, probe, or temporary investigation artifact may be copied into `E:\Anvien` or written into the target.
- Anvien's supported target analyze may regenerate or timestamp ignored guidance files such as `AGENTS.md`/`CLAUDE.md`; record that tool-boundary side effect against a pre/post baseline and never manually revert it as part of this plan.
- All plan, investigation, QA, and Supervisor artifacts belong in `E:\Anvien`. Debug-only temporary material belongs under `E:\Anvien\.tmp\`; reusable synthetic fixtures belong under the owning package's `testdata\`.
- Synthetic fixtures must be independently authored from the language contract. Do not copy `cheapapp.org` source into Anvien testdata and do not create fixtures at repository root.
- No runtime network access, package installation, package scripts, Node/TypeScript execution, or declaration borrowing from another repository may be used by the production analyzer. A pinned TypeScript compiler may be used only as a test oracle.
- Unsupported, missing, stale, or mixed index versions must fail visibly with `INDEX_VERSION_MISMATCH`; no empty-success, implicit downgrade, or version inference from ID strings.
- Until P2-G, the active v1 path is protected by an explicit compatibility adapter and byte/hash preservation gate. Strict/lossless occurrence, RelationshipID, and v2 producer behavior is exercised on the shadow-v2 path; it must not silently alter the active v1 graph before cutover.
- Before each implementation commit: update actual-status/evidence/benchmark, remove superseded plan-created artifacts, obtain Supervisor PASS, run `anvien detect-changes --repo E:\Anvien --scope all`, and commit only the active slice.

## Problem

The accepted bounded investigation proves five independent defects after file discovery:

1. TypeScript array binding patterns are rejected before a `DefinitionFact` is created.
2. graph definition identity omits scope/range, so distinct same-name locals collide and `Graph.AddNode` overwrites one occurrence;
3. TypeScript module export semantics are not represented in `DefinitionFact`, and persistence consumers disagree between `visibility` and `isExported`;
4. import resolution stops at physical definitions in a barrel and does not traverse re-export bindings to the terminal declaration;
5. the resolver workspace has no TypeScript ambient/lib declaration universe, while graph-health guesses classifications from unresolved target strings instead of consuming structured resolver outcomes.

These defects cannot be safely fixed as five local patches. Identity changes propagate into relationship endpoints, process references, resolution gaps, embeddings, groups, rename, MCP, HTTP, Web, persistence order, and derived IDs. The current analyze lifecycle deletes live artifacts before rebuild and does not atomically publish graph, Ladybug, metadata, embeddings, and cache generation. A partial Node-ID-v2 cutover would therefore create silent mixed-version or mixed-generation results.

No relevant `SPEC-MAP.md` or approved `Docs/SPEC` authority exists in the current repository. This plan records the recommended contract as a proposed owner decision. Production implementation remains blocked until the owner accepts this plan or the equivalent contract is recorded in an approved ADR/SPEC; no implementation agent may silently choose a different identity/export/external contract.

## Scope

In scope:

- versioned `DeclarationID`, `SymbolID`, source range, selection range, meaning, and graph-generation contracts;
- explicit graph insert, enrich/update, replace, decode, collision, and closure validation behavior;
- full v2 reindex, version-aware readers, opaque-ID consumers, legacy ambiguity behavior, generation staging, atomic publication, and rollback;
- recursive TypeScript binding-pattern extraction for declarations, parameters, catch bindings, and `for-of` declarations;
- first-class TypeScript module export facts, direct/default/alias/type-only/re-export semantics, and projection compatibility;
- TypeScript module request resolution, per-module export tables, barrel/re-export traversal, alias proofs, cycles, ambiguity, and meaning lanes;
- a versioned declaration universe for TypeScript standard libraries and locally installed package declarations, with structured resolution outcomes;
- graph-health projection from resolver outcomes;
- JSON/Ladybug/CLI/MCP/HTTP/Web parity, deterministic output, target integration, performance/capacity measurement, and fault-injection rollback proof.

Boundary:

- Anvien source and tests are changed only in `E:\Anvien`.
- The target `E:\cheapapp.org` is a read-only integration subject except for Anvien's normal repo-local `.anvien` operational output.
- Each phase changes only its named contract and adapters. Language providers other than TS/JS are regression surfaces unless an identity API migration requires a narrow adapter.

## Non-Goals

- No fix for scanner treatment of nested `env`, `target`, or `logs` directories.
- No claim of complete TypeScript compiler conformance.
- No materialization of all `node_modules` or all `.d.ts` declarations as repository File/Definition nodes.
- No network download, runtime `tsc`/Node invocation, package-script execution, or dependency installation by analyze.
- No redesign of product Features, Architectural Boundaries, Processes, communities, or app-layer semantics beyond keeping them generation-consistent and ID-opaque.
- No target-source edit, target copy, target fixture, or report stored in `E:\cheapapp.org`.
- No rewrite of v1 graph data in place; lost declarations must be recovered by reparsing source.
- No dual-active v1/v2 graph that lets readers silently combine versions.
- No unrelated refactor of scanner, UI design, database scalar/nullability contracts, or other language semantics.

## Requirements

### Architecture / contract rules

The production graph model is:

```text
File --DECLARES--> Declaration --DECLARES_SYMBOL--> Symbol
Module --EXPORTS/REEXPORTS-----------------------> Symbol
SourceSite --ResolutionOutcome--> InternalSymbol | ExternalSymbol | Gap
```

- `DeclarationID` identifies one declaration occurrence; `SymbolID` identifies one logical language symbol. They are separate types and cannot be substituted.
- Function/Class/Variable-style semantic nodes represent logical Symbols in v2. Declaration nodes retain every source occurrence, including overload signatures, implementations, merged declarations, invalid duplicates, and destructured leaves.
- A durable reference is `{repoKey, graphGeneration, graphSchemaVersion, identitySchemaVersion, symbolID}`. A raw graph ID alone is not a cross-generation or cross-repository contract.
- IDs are deterministic, versioned, opaque, and derived from canonical tuples. Consumers must use explicit fields for file path, label, line, route, community, and provenance instead of parsing IDs.
- `range` covers the syntactic construct; `selectionRange` covers the identifying token. Canonical offsets are zero-based half-open bytes; line/column base and encoding are explicit, with UTF-16 conversion only at editor/API boundaries.
- Same-name locals in different lexical owners remain distinct. Body/comment edits do not change `SymbolID`; rename, move, owner, language meaning, or module identity changes may do so. Declaration occurrence IDs may change with their canonical source anchor.
- Overloads and declaration merging produce multiple Declaration nodes linked to one Symbol only when the provider has language-semantic merge evidence. Heuristic candidates remain separate and explicitly unverified.
- Duplicate canonical IDs with conflicting payload/provenance fail the generation. Byte-identical idempotent reinsertion is allowed only through an explicit operation. No first-wins, last-wins, generic upsert, panic-only, or warning-only behavior.
- Module export is first-class and separate from access visibility. `isExported` may exist only as a versioned derived compatibility field; it is not resolution authority.
- Type, value, and namespace meanings are separate lanes. Explicit import failure never falls back to a global same-name definition.
- Export resolution preserves every alias/re-export hop, excludes `default` from star exports, gives explicit exports precedence, deduplicates paths to the same terminal `ResolutionTarget`, returns ambiguity for different terminals, and handles cycles without silent truncation.
- Ambient/external resolution uses a hash-bound project profile and versioned declaration universe. External Symbols materialize lazily when referenced.
- `TargetDef == nil` is not a diagnostic contract. Every source site has one structured resolution outcome with code, stage, candidates, proof, config/catalog hash, and optional target.
- Graph-health projects resolver outcomes; it does not reclassify by target-name heuristics.
- `scopeIrVersion`, `graphSchemaVersion`, `identitySchemaVersion`, `positionEncoding`, analyzer build, config/catalog hash, and `graphGeneration` are persisted and checked at every reader boundary.
- A full generation is staged, validated, and published as one unit. Faults retain the prior queryable generation. Root compatibility paths remain under the analyzed repository's `.anvien`.
- The active-generation manifest is the compatibility authority and contains `protocolVersion`, `minReaderProtocol`, `minReaderBuild`, `graphSchemaVersion`, `identitySchemaVersion`, `scopeIrVersion`, `generation`, `configHash`, and `catalogHash`. A reader must validate all of them before opening JSON, Ladybug, native Cypher, fallback Cypher, a cache, a group contract, an embedding row, or an HTTP/Web stream.
- v2 storage uses a non-overlapping generation/protocol path. The legacy root compatibility path either remains an older supported generation or returns `INDEX_VERSION_MISMATCH`; it must never expose v2 records to a v1 binary/client that ignores unknown fields. CLI startup/status/index/file-detail/graph-health, MCP resources/tools/rename/detect-changes, HTTP, Web, native/fallback Cypher, caches, groups, embeddings, and global registry all participate in the old-reader matrix.
- Cutover is a protocol negotiation, not a tolerant JSON-field rollout: v2 is published only beneath a versioned media/layout namespace, every reader sends `readerProtocolVersion`/`readerBuild`, the manifest returns `minReaderProtocol`/`minReaderBuild`, and the body is unopened until the handshake passes. An old binary/client fixture must fail closed with `INDEX_VERSION_MISMATCH`; no unaware v1 path may open, cache, stream, or register v2 data.
- Until P2-G, the active v1 adapter remains byte/hash-preserving and is never fed strict v2 mutation results. Strict/lossless occurrence and RelationshipID behavior runs on the isolated in-memory/shadow-v2 adapter; a compatibility adapter may reproduce the historical v1 projection only behind an explicit version flag. P2-G is the first slice allowed to select v2 for a newly published generation.
- The canonical parity record for a Node includes `id`, label, repo-relative file path, name, DeclarationID, SymbolID, range, selectionRange, position encoding, meaning mask, access visibility, direct export fact, generation, and provenance. The canonical parity record for a Relationship includes `id`, type, source/target refs, step, confidence, meaning, resolution status, proof hops, generation, and provenance. Parity counts field-level differences, not only IDs/endpoints.
- The exact public parity surface matrix is fixed and reused by P4/P6/P7 ledgers: `S0` Graph JSON; `S1` Ladybug native Cypher; `S2` Go/fallback Cypher; `S3` the exact union of every `index-reader-matrix.md` row tagged `S3`; `S4` the exact union of every row tagged `S4`; `S5` the exact union of every row tagged `S5`; `S6` the exact union of every row tagged `S6`; `S7` the exact union of every row tagged `S7`; `S8` the exact union of every row tagged `S8`; `S9` the exact union of every row tagged `S9`; `S10` the exact union of every row tagged `S10`; and `S11` the exact union of every row tagged `S11`. The matrix row tags, not a prose category list, are authoritative for every reader/cache/registry/derived-surface denominator. A surface is PASS only when its canonical records and orphan/reference checks are zero-difference; no plan slice may invent a smaller “six surface” denominator.
- The handshake wire contract is authoritative and reused verbatim by the reader matrix: request `{readerProtocolVersion, readerBuild, supportedGraphSchemaVersions[], supportedIdentitySchemaVersions[], supportedScopeIrVersions[]}`; manifest `{protocolVersion, minReaderProtocol, minReaderBuild, graphSchemaVersion, identitySchemaVersion, scopeIrVersion, generation, configHash, catalogHash}`; failure envelope `{code:"INDEX_VERSION_MISMATCH", expected:{...}, actual:{...}, retryable:false}`. S0/S1/S2 return the typed loader/driver error before body/query rows; S3 returns the same code on stderr with nonzero exit; S4 returns a JSON-RPC error whose `data.code` is the same string; S5 returns HTTP `409` with the envelope before a stream body; S6 renders the mismatch state without parsing records; S7/S8 treat it as a cache miss/error and never return stale rows; S9/S10/S11 return the typed error and no projection. No reader may downgrade or infer a version from an opaque ID.
- A resolution target is a tagged union: `InternalSymbolRef`, `ExternalSymbolRef`, `ModuleNamespaceRef`, `IntrinsicSymbolRef`, or `GapRef`. `ModuleNamespaceRef` is a synthetic namespace Symbol with module identity and export-surface provenance; it is not a physical definition named `ns`. `IntrinsicSymbolRef` is a language-owned primitive/builtin identity with language, meaning, version, and provenance; it is never represented as a repository `InternalSymbolRef`.
- `ExternalSymbolRef` contains an opaque external ID, name, meaning mask, origin (`stdlib_catalog`, `package_declaration`, or `ambient_module`), module/package, declaration locator when safe, catalog/config hashes, generation, and provenance. External nodes are materialized only by the dedicated external-symbol owner when referenced; catalog declarations never become repository File nodes.
- A `ResolutionOutcome` contains exactly one top-level source-site result plus nested module/export/meaning/member stages, target union, candidates, proof hops, severity, actionability, retryability, config/catalog hashes, and budget/security metadata. A module failure cannot be overwritten by a later export/member guess.
- Resolution outcomes use this closed status set: `resolved_local`, `resolved_external`, `resolved_intrinsic` (language primitives only; not `Promise` or `Math` catalog declarations), `module_not_found`, `module_blocked_by_exports`, `export_not_found`, `meaning_mismatch`, `member_not_found`, `ambiguous_export`, `cycle_no_terminal`, `config_invalid`, `external_declaration_unavailable`, `catalog_version_mismatch`, `declaration_parse_failed`, `unsupported_syntax`, `budget_exceeded`, and `security_path_rejected`. Every status has a defined stage, allowed target union, severity, actionability, and retry policy in the P6 status matrix.
- Re-export resolution uses a resolution set: a cycle branch contributes no candidate; a cycle with another terminal branch resolves the terminal; a cycle with no terminal returns `cycle_no_terminal`; paths to the same terminal are deduplicated; distinct terminals return `ambiguous_export`; no branch is silently truncated.
- The declaration/package boundary is split: P5 resolves a source specifier, path mapping, and package `exports` conditions to a hash-bound `ModuleRef` plus immutable package metadata; the P6 declaration-entrypoint owner consumes that `ModuleRef` to select `types`/`typesVersions` or catalog declaration entrypoints and must not reimplement source-specifier, path-mapping, or package-`exports` condition selection.
- Package ownership is exact: P5 owns only source-specifier, path mapping, and package `exports` condition selection and records the selected module plus immutable package metadata; P6 owns `types`/`typesVersions` declaration-entrypoint selection from that metadata and must not re-evaluate `exports` conditions. A `ModuleRef` without a declaration entrypoint is still a valid module result, not a symbol result.
- Publication includes repo-local graph/Ladybug/meta/embedding/cache namespaces plus the global repo registry and every group contract registry. Immutable generation-qualified copies are staged, validated, and fsynced; a compare-and-swap active epoch/manifest is the only publication pointer. Readers pin a generation lease, and GC cannot remove an old generation until all leases (including registry/group/cache readers) are released. A crash at any artifact write, registry/group write, cache publication, active-pointer switch, restart, or GC boundary must leave the previous epoch queryable.
- Every S10 group snapshot is `{groupEpoch, memberRepoGenerationVector:[{repoKey, graphGeneration, graphSchemaVersion, identitySchemaVersion}]}`. Group sync pins all member generations, CASes that vector with the active group pointer, and returns `INDEX_VERSION_MISMATCH` on any member-generation conflict; it retains the previous group snapshot. Global-repo publication and group-contract publication are separate fault/parity rows, while both obey the same lease/GC rules. Every external pointer and S7–S10 cache key carries the selected repo generation and, for groups, the member vector hash.

### Proposed decisions requiring owner acceptance before P1-B

| Decision | Recommended choice | Rejected shortcut |
|----------|--------------------|-------------------|
| Node meaning | Semantic nodes are Symbols; add Declaration nodes | put range into the old single node and call it solved |
| Durable identity | generation/version-qualified `SymbolRef` | persist raw UID indefinitely |
| Export meaning | direct module export is authoritative; package public API is separate derived data | one transitive `isExported` boolean |
| Legacy collision | unique mappings may redirect with warning; collided mappings return all candidates as ambiguous | auto-pick one legacy `time`/`now` |
| External declarations | embedded versioned stdlib catalog plus locally resolved package declarations | hardcode `Promise`/`Math`, scan everything, or download |
| Publication | generation directory plus atomic active-generation manifest; compatibility files follow the active generation | delete live data first and publish artifacts independently |
| Canonical identity tuple | repo-relative canonical module/file, lexical owner chain, normalized semantic name, declaration role, meaning namespace, provider merge key, binding path, and explicit stability tier; path separators/case and Unicode normalization are fixed | file/name/arity or raw range alone |
| Meaning lanes | one Symbol may carry a `MeaningMask` for a TS declaration that intentionally declares both type and value (class/enum); distinct type-only and value-only declarations with the same name receive distinct SymbolIDs | one global name bucket or one ID per requested operation |
| Relationship identity | each source-site edge has a `RelationshipID` from type, source-site ID, source/target refs, meaning, ordinal, and provenance; aggregation is an explicit lossless operation retaining all source-site IDs | merge same endpoints and discard occurrence provenance |
| Declaration stability matrix | body/comment/blank-line edit: SymbolID same, DeclarationID follows selection anchor; local rename/owner/module move: new IDs plus evidence-backed legacy candidate; added overload/merge: existing SymbolID same, new DeclarationID; anonymous/default expression: snapshot stability tier only; unverified merge: separate Symbols | silently rebind an occurrence after a structural change |
| Export authority | direct export facts, default/alias/type-only/star/namespace forms, anonymous/default expressions, package public API, and transitive barrel reachability each have separate fields and expected values | infer all forms from one `isExported` boolean |

### One-file / one-responsibility ownership rule

The following names are implementation suggestions, not architecture authority. Pre-flight may choose an equivalent existing owner, but the responsibility boundary is mandatory and the plan must be updated before editing if the file map changes.

| Responsibility | Planned owner file | May depend on | Must not own |
|----------------|--------------------|---------------|--------------|
| identity schema/version values | `internal/graphidentity/version.go` | primitive types only | ID construction, graph mutation |
| DeclarationID construction | `internal/graphidentity/declaration_id.go` | canonical position/owner inputs | Symbol merge policy |
| SymbolID construction | `internal/graphidentity/symbol_id.go` | semantic owner/meaning inputs | declaration range conversion |
| legacy identity mapping | `internal/graphidentity/legacy_mapping.go` | v1/v2 references | reader compatibility policy |
| declaration fact contract | `internal/scopeir/declaration_fact.go` | range, labels, meanings | imports, calls, exports |
| binding-pattern fact contract | `internal/scopeir/binding_pattern_fact.go` | declaration reference/range | AST traversal |
| export fact contract | `internal/scopeir/export_fact.go` | module request, meanings | resolution algorithm |
| module request fact contract | `internal/scopeir/module_request_fact.go` | range and language | file-system resolution |
| strict graph mutation operations | `internal/graph/mutation.go` | graph node/edge types | snapshot/version publishing |
| graph invariant validation | `internal/graph/validation.go` | immutable graph view | mutation or repair |
| relationship identity/aggregation | `internal/graph/relationship_identity.go` | source-site/provenance refs | node identity or resolver decisions |
| index compatibility policy | `internal/repo/index_version.go` | persisted metadata | analyze staging |
| generation staging | `internal/analyze/generation_stage.go` | repository storage | active-generation switch |
| generation publication/rollback | `internal/analyze/generation_publish.go` | validated staged generation | graph construction |
| recursive TS binding traversal | `internal/providers/tsjs/binding_patterns.go` | Tree-sitter + ScopeIR facts | export or import extraction |
| TS export extraction | `internal/providers/tsjs/exports.go` | Tree-sitter + ExportFact | re-export resolution |
| TypeScript project profile | `internal/resolution/typescript_project.go` | tsconfig/package metadata | export closure |
| module reference/condition resolution | `internal/resolution/module_reference.go` | project profile, package conditions | declaration `types`/`typesVersions` lookup |
| module export table construction | `internal/resolution/module_export_table.go` | ScopeIR export facts | call/type resolution |
| barrel/re-export traversal | `internal/resolution/reexport_resolution.go` | module export tables | package declaration loading |
| structured resolution outcome | `internal/resolution/outcome.go` | finalized target/proof/status | graph-health classification or external materialization |
| resolution status matrix | `internal/resolution/outcome_status.go` | stage/target/severity policy | entrypoint lookup, materialization, or graph-health inference |
| declaration-universe interface | `internal/resolution/declaration_universe.go` | project profile | concrete catalog storage |
| declaration entrypoint lookup | `internal/resolution/declaration_entrypoint.go` | `ModuleRef`, declaration universe | module specifier/condition resolution |
| external Symbol materialization | `internal/resolution/external_symbol.go` | ExternalSymbolRef/catalog | repository declaration extraction |
| declaration path security/budgets | `internal/resolution/declaration_security.go` | canonical roots and limits | network/package execution |
| embedded TS stdlib catalog | `internal/resolution/typescript_stdlib_catalog.go` | generated immutable data | network/package execution |
| catalog manifest/provenance | `internal/resolution/typescript_catalog_manifest.go` | version/hash/license/NOTICE | resolution traversal |
| outcome-to-health projection | `internal/graphhealth/resolution_outcome_projection.go` | structured outcomes | name-based inference |

Every corresponding test file owns one behavior matrix. Existing broad files such as `internal/scopeir/facts.go`, `internal/resolution/indexes.go`, `internal/resolution/resolve.go`, and `internal/graph/types.go` may retain only their current coordinating responsibility; new semantic logic must be extracted into the dedicated owner above. Adapters left in those files must be thin delegation only.

### Binding and export fact schemas

`BindingPatternFact` and `BindingFact` are source-of-truth contracts, not test-only shapes:

```text
BindingPatternFact:
  patternID, kind(identifier|array|object|rest|assignment), scopeID,
  range, selectionRange, declarationKind, origin

BindingFact:
  bindingID, declarationID, symbolID, patternID, bindingPath, localName,
  propertyKey, arrayIndex, isRest, hasDefault, scopeID,
  range, selectionRange, meaningMask, origin
```

`bindingPath` is a typed sequence: `array(index)`, `object(key)`, `computed(keyExpression)`, and `rest`. A default initializer is a modifier (`hasDefault=true`, with its own range), not an identity path segment. Array holes have no BindingFact but retain source indexes: in `[a,,b]`, `a` is `array(0)` and `b` is `array(2)`. Assignment destructuring emits a write to an existing binding, not a new DefinitionFact. Imports are either represented once by the shared walker or explicitly excluded with a test.

`ExportFact` is one fact per exported binding/specifier, not one fact per declaration statement. It contains `exportID`, source-site `range`/`selectionRange`, `moduleScopeID`, `kind(local|default|alias|reexport|star|namespace)`, `exportedName`, `localName`, `localDeclarationID`/`localSymbolID` when present, `targetModuleRef`, `targetExportedName`, `meaningMask`, `typeOnly`, `direct`, `reachableThroughBarrel`, `publicApi`, and immutable syntax/link status (`syntactic_valid`, `target_module_known`, `target_name_known`, or `syntax_unsupported`). It does not contain a mutable P5 `ResolutionOutcome`; P5 stores that outcome separately by export/source-site ID. A multi-specifier list yields one fact per specifier; a multi-declarator statement yields one local export fact per binding. Anonymous/default expressions use an explicit expression anchor and stability tier, never “first definition in file”.

The derived compatibility field `isExported` means only `directExportCount > 0` for the defining Symbol in its own module. Alias/re-export entries use `exportEntry`/`reachableThroughBarrel` fields; type-only exports retain `typeOnly=true`; package public API uses `publicApi=true`. No consumer may use `isExported` to resolve a transitive alias.

The P1-A/P4-A semantic vector is authoritative:

| Syntax case | Expected fact/target |
|-------------|----------------------|
| named local/default class, function, interface, or expression | one local/default ExportFact; anonymous expression uses expression anchor; class/enum may carry type+value mask |
| `export {x as default}` | alias ExportFact to local `x`, default name, proof to local Declaration/Symbol |
| `export {default as X} from "m"` | re-export ExportFact whose target name is `default` in module `m` |
| inline type specifier / `export type {X}` / `export type *` | type-only meaning; value use is `meaning_mismatch` |
| `export * from "m"` | star adjacency only; `default` excluded |
| `export * as ns from "m"` | `ModuleNamespaceRef` for `ns`; member use resolves through the target export surface |
| class/enum dual lane | one Symbol with type+value mask when provider evidence says one declaration supplies both |
| type-only declaration used as value | no fallback; `meaning_mismatch` |
| overloads/declaration merging | one logical Symbol with multiple DeclarationRefs only when merge evidence exists |
| explicit export vs star export | explicit entry wins; different terminal candidates are `ambiguous_export` |

### Declaration security, profile, and budget contract

The following defaults are mandatory starting ceilings; P1-A/P5-A may tighten them, but may not remove a limit or leave it qualitative. Exceeding a limit returns `budget_exceeded` and emits no partial/truncated resolution.

| Boundary | Required contract |
|----------|-------------------|
| Allowed roots | The repository root, configured workspace roots, and the pinned embedded catalog root only; package declaration roots are accepted only after `ModuleRef` conditions and package policy validate them. |
| Canonical paths | Resolve absolute/relative paths, reject `..` escape, normalize case according to the repository platform policy, and reject ambiguous aliases before opening a file. |
| Symlink/junction/pnpm store | Resolve links before authorization; reject a symlink/junction or pnpm-store path that escapes an allowed root unless an explicit, hash-bound package root is in the profile. |
| Triple-slash/reference closure | `/// <reference path>` and lib references must remain inside the authorized declaration universe/root and obey the same depth/byte limits. |
| Network/runtime | No network, Node, package scripts, package installation, or arbitrary executable; catalog is embedded at build time and package declarations are read-only local inputs. |
| Declaration files | Maximum `50,000` files, `256 MiB` cumulative bytes, `8 MiB` per file, and reference depth `64` per resolution request. |
| Export traversal | Maximum `64` re-export hops, `2,000` star edges per module, `10,000` candidates per lookup, and `2,000` modules per SCC. |
| Catalog | Uncompressed catalog `<=32 MiB`, load `<=250 ms` cold on the benchmark machine, and package declaration cache `<=256 MiB`; values are measured in B6 and cannot be waived without measured owner approval. |
| Caches | Cache keys include repository, generation, protocol, project-config hash, catalog hash, and package-root hash; stale entries are rejected and debug traces redact absolute paths outside approved evidence. |
| Manifest/provenance | Catalog manifest records TypeScript/compiler version, target/lib closure, source hashes, license/NOTICE, generation, and regeneration command; a hash mismatch returns `catalog_version_mismatch`. |

The TypeScript profile must record compiler version, `module`, `moduleResolution`, `target`, explicit `lib` closure, `noLib`, `types`, `typeRoots`, `baseUrl`, `paths`, package conditions, `skipLibCheck`, and triple-slash references. `Promise` and `Math` are catalog-backed external/ambient declarations, not intrinsic primitive types.

### Slice execution protocol

Every implementation slice follows this order:

1. refresh the Anvien self-graph and actual-status row;
2. record file-detail and symbol/file impact for the exact editable owner;
3. implement production behavior only;
4. after production behavior is complete, add focused behavior tests and reusable package `testdata`;
5. run the repository full build before validation;
6. validate the nearest real CLI/MCP/HTTP/Web/Ladybug boundary and record benchmarks;
7. refresh all four ledgers, remove superseded slice artifacts, and obtain Supervisor PASS;
8. run detect-changes, commit only that slice, and verify the worktree before another slice opens.

## Acceptance Criteria

- Graph JSON and every persisted/read projection contain `0` duplicate Node IDs, `0` duplicate Relationship IDs, `0` missing endpoints, `0` orphan Process/ResolutionGap/embedding references, and no mixed generation.
- Five force-analyzes of identical source/config/analyzer inputs produce identical canonical node-ID and relationship-ID sets; worker/scan order does not alter the result.
- Conflicting identity injection fails closed with structured evidence. Identical idempotent insertion works only through its explicit operation.
- Unsupported, absent, or mixed index versions fail every CLI/MCP/HTTP/Cypher reader with `INDEX_VERSION_MISMATCH`.
- Fault injection at graph build, Ladybug load, metadata, embedding, active-manifest, and compatibility-publication boundaries leaves the prior generation intact and queryable.
- Every surface `S0`–`S11` exposes the same canonical node/relationship records for the active generation; native Cypher and fallback Cypher are measured separately, and field-level differences plus orphan references are both zero.
- Both `time` declarations and both `now` declarations in the bounded target remain distinct and traceable; no source occurrence is overwritten.
- All six bounded array bindings are represented (`6/6`) with correct declaration, binding path, range, scope, and downstream reference behavior.
- All 21 bounded direct exports are correct (`21/21`) in ScopeIR and the applicable `S0`–`S11` projections; access visibility remains separate.
- Both bounded barrel calls resolve through their proof chains to the terminal function (`2/2`); direct `IMPORTS` edges still describe syntactic module dependencies.
- `Promise`, `Math.max`, and `Math.min` are resolved external/intrinsic outcomes or explicit external-capability outcomes, never false `in_repo_unresolved` gaps.
- Explicit missing imports, type/value mismatches, ambiguous star exports, pure cycles, invalid config, absent package declarations, and security-rejected paths produce their exact structured status with no global-name fallback.
- The target's known scanner baseline remains quarantined at the same exact eight omissions unless a separate scanner plan is authorized; this plan creates `0` new File-node omissions.
- No source, fixture, report, probe, or temporary investigation artifact is written into or copied from `E:\cheapapp.org`; operational graph/index output remains under `E:\cheapapp.org\.anvien`.
- Before identity cutover, at least five measured v1 runs on the same commit-bound corpus/config/build/machine/cache policy establish analyze median, Ladybug-load median, native-query p95, fallback-query p95, graph size, and peak RSS. Final medians and p95 values regress by no more than `10%`, peak RSS by no more than `15%`, unless the owner accepts a measured exception with both baseline and final values before cutover.
- Every touched/new source and test file passes the one-file/one-responsibility review; `0` new catch-all files and `0` unrelated responsibility additions.
- Every slice has a separate commit after full build, boundary validation, ledger update, Supervisor PASS, and detect-changes.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.

### P1: Graph identity contract and strict graph construction

- Phase Goal: ratify and implement the Declaration/Symbol identity foundation without activating v2 for readers.
- Phase Boundary:
  - In scope: contract authority, responsibility extraction, ranges, meanings, canonical IDs, strict mutation operations, graph validation, and shadow v2 output.
  - Out of scope: active index cutover, binding-pattern behavior, export resolution, ambient declarations.
  - Dependencies: P0 complete; owner accepts the proposed contract before P1-B.
- Phase Implementation Rule: do not implement `P1` directly. Complete, review, and commit each ordered slice before opening the next.
- Ordered Slice List:
  - P1-A: Ratify graph identity and ownership contract.
  - P1-B: Introduce range, DeclarationID, SymbolID, and SymbolRef types.
  - P1-C0: Preserve lossless declaration occurrences.
  - P1-C0A: Define RelationshipID and lossless source-site aggregation.
  - P1-C0B: Validate lossless graph decode and closure.
  - P1-C: Build declaration-to-symbol identity mapping.
  - P1-D: Introduce strict graph mutation operations and validation.
  - P1-D1: Migrate core graph producers to explicit operations.
  - P1-D2: Migrate resolution/projection producers to explicit operations.
  - P1-D3: Migrate ancillary/document/semantic producers to explicit operations.
  - P1-E: Emit and validate shadow identity v2.

- [ ] P1-A: Ratify graph identity and ownership contract.
  - Goal: convert the proposed decisions in this plan into explicit owner-approved authority before production code.
  - Scope Boundary:
    - Editable: plan/ADR or approved contract documentation under `E:\Anvien` only.
    - Inspect-only: accepted causal reports and current source owners.
    - Preserve-only: all production code and target files.
    - Out of scope: implementation.
  - Non-Goals: no exact helper/function prescription beyond ownership boundaries.
  - Pre-flight Questions:
    - Data source: accepted investigation, P0 evidence, and the explicit Owner response.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read accepted reports and the current proposed contract; no runtime store is opened.
    - DB write flow: N/A — this documentation-only authority slice writes no graph/index state.
    - Render location: the canonical Anvien contract/ADR document and Supervisor report.
    - UI behavior flow: N/A — architecture decision and review workflow only.
    - Docker runtime: N/A — documentation-only authority work.
    - Playwright target: N/A — nearest observable boundaries are document cross-reference and Supervisor review.
    - Behavior test: decision-table completeness, contradiction review, version/rollback coverage, and one-file ownership coverage.
    - Cleanup/quarantine: retain one accepted authority document; remove draft duplicates.
    - External side effects: None.
    - N/A notes: Documentation-only slice; production source and target remain preserve-only.
  - Work Steps:
    1. Resolve every proposed decision row and record accepted/rejected alternatives, stability guarantees, version behavior, and one-file ownership.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: confirm persisted reference and generation contract.
       - Render location check: canonical Anvien documentation only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1A-CONTRACT1`.
    2. Run zero-trust architecture/Supervisor review and update P1-B only if authority changes.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: verify no unresolved storage/rollback decision.
       - Render location check: review report under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1A-REVIEW1`.
  - Implementation Gate: every proposed decision is accepted or explicitly changed; no architecture question remains implicit.
  - Acceptance:
    - Source: no production source changed.
    - Runtime/UI: N/A — this slice changes no UI; the named real non-UI boundary passes after the full build.
    - DB/data: identity, persistence, and rollback contracts are explicit.
    - Behavior test: authority cross-reference has no contradiction.
    - Cleanup/quarantine: no draft duplicate authority remains.
    - Evidence IDs: `E1-P1A-CONTRACT1`, `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1`.
    - Actual-status rows refreshed: architecture authority and P1-B assumptions.
  - Evidence Targets: accepted decision table, ownership map, Supervisor verdict.
  - Actual-status Update: `architecture authority missing -> correct` or record a blocker.
  - Commit Boundary: commit the contract-only slice and record `E1-P1A-COMMIT1`; no Anvien detect-changes is required for a documentation-only commit.

- [ ] P1-B: Introduce range, DeclarationID, SymbolID, and SymbolRef types.
  - Goal: add versioned, non-interchangeable identity and position types in single-responsibility owners without changing active graph IDs.
  - Scope Boundary:
    - Editable: graph-identity and declaration/range owner files plus narrow serialization adapters.
    - Inspect-only: provider, resolver, graph emitter, API consumers.
    - Preserve-only: active v1 output and all target files.
    - Out of scope: symbol merge, graph mutation, cutover.
  - Non-Goals: no range-only ID shortcut and no consumer parsing changes.
  - Pre-flight Questions:
    - Data source: the approved P1-A contract and current ScopeIR range inputs.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: deserialize and round-trip isolated current-v1 and proposed-v2 position/identity fixtures.
    - DB write flow: test-only serialization fixtures; no live database/index mutation.
    - Render location: compiler, serialization, and evidence output only.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real serialization/type round-trip boundary after the full build.
    - Behavior test: Unicode, CRLF/LF, tabs, emoji, combining characters, base/encoding, and type-safety round trips.
    - Cleanup/quarantine: independently authored package `testdata`.
    - External side effects: None.
    - N/A notes: No public runtime or live persistence is changed.
  - Work Steps:
    1. Refresh graph; record file-detail/impact; extract mixed declaration/range responsibilities before adding the production types.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: preserve v1 serialization byte-for-byte.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1B-IMPACT1`, `E1-P1B-SRC1`.
    2. Implement production types, then focused tests; run full build before round-trip validation.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: v1 output unchanged; v2 types deterministic.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1B-BUILD1`, `E1-P1B-TEST1`.
  - Implementation Gate: P1-A PASS; ownership files each have one responsibility; fresh HIGH/CRITICAL warnings recorded.
  - Acceptance:
    - Source: distinct types and explicit encoding/base contract compile.
    - Runtime/UI: N/A — this slice changes no UI; the named real non-UI boundary passes after the full build.
    - DB/data: deterministic serialization and v1 preservation pass.
    - Behavior test: full position matrix passes.
    - Cleanup/quarantine: no root fixture or catch-all helper.
    - Evidence IDs: `E1-P1B-IMPACT1`, `E1-P1B-SRC1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-REVIEW1`.
    - Actual-status rows refreshed: range/identity types.
  - Evidence Targets: source diff, full build, range/type tests, Supervisor, detect-changes, commit.
  - Actual-status Update: `identity types missing -> correct`; keep active graph identity wrong until P2-G.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-C0: Preserve lossless declaration occurrences.
  - Goal: remove first-wins/last-wins loss of declaration occurrences before Declaration-to-Symbol mapping.
  - Scope Boundary:
    - Editable: dedicated occurrence-index owner and shadow-v2 compatibility adapter.
    - Inspect-only: RelationshipID and decode owners.
    - Preserve-only: active v1 output.
    - Out of scope: relationship aggregation, graph decode, logical symbol merge, and resolver semantics.
  - Non-Goals: no active v2 cutover and no relationship or endpoint policy changes.
  - Pre-flight Questions:
    - Data source: the complete ScopeIR declaration-producer inventory.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read in-memory declaration occurrences and active-v1 compatibility bytes.
    - DB write flow: write only the isolated in-memory shadow-v2 occurrence index.
    - Render location: structured validation error and evidence ledger.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real occurrence round-trip boundary after the full build.
    - Behavior test: duplicate, overload, merged, invalid-duplicate, empty, and subset declaration occurrences.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores and remove superseded debug artifacts before review.
    - External side effects: None.
    - N/A notes: The shadow path is deliberately non-active and unavailable to readers.
  - Work Steps:
    1. Implement only the lossless Declaration occurrence index and shadow-v2 adapter; keep the active v1 adapter byte/hash-preserving until P2-G.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: occurrence counts survive the v2 shadow path while v1 output remains byte/hash-identical.
       - Render location check: structured validation error only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1C0-IMPACT1`, `E1-P1C0-SRC1`.
    2. Add occurrence-only tests after production code, run full build, and compare declaration occurrence conservation on synthetic collision fixtures; defer RelationshipID and decode checks to P1-C0A/P1-C0B.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no first/last-wins declaration loss in v2; v1 compatibility behavior is explicitly isolated.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1C0-BUILD1`, `E1-P1C0-TEST1`.
  - Implementation Gate: P1-B PASS; every current declaration duplicate-suppression site is inventoried; v1 compatibility byte/hash gate is defined.
  - Acceptance:
    - Source: every v2-shadow input occurrence remains addressable; active v1 output is unchanged.
    - Runtime/UI: N/A — this slice changes no UI; the named real non-UI boundary passes after the full build.
    - DB/data: `0` silently dropped v2 declarations; conflicting payloads fail closed; v1 hash remains unchanged.
    - Behavior test: declaration collision, overload, merge-evidence, and empty/subset occurrence matrix passes; relationship/decode failures are tested only in their named slices.
    - Cleanup/quarantine: old first/last-wins occurrence helpers are removed or isolated behind explicit compatibility code.
    - Evidence IDs: `E1-P1C0-IMPACT1`, `E1-P1C0-SRC1`, `E1-P1C0-BUILD1`, `E1-P1C0-TEST1`, `E1-P1C0-REVIEW1`.
    - Actual-status rows refreshed: occurrence index and v1 compatibility adapter.
  - Evidence Targets: producer inventory, source diff, v1 hash gate, full build, conservation/fault tests, Supervisor, detect-changes, commit.
  - Actual-status Update: declaration occurrence preservation `missing/wrong -> correct`; unlock P1-C0A and P1-C0B.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-C0A: Define RelationshipID and lossless source-site aggregation.
  - Goal: give each relationship occurrence a deterministic identity and make aggregation explicitly retain all source-site IDs.
  - Scope Boundary:
    - Editable: `internal/graph/relationship_identity.go` and one aggregation adapter.
    - Inspect-only: declaration index/decode.
    - Preserve-only: v1 output.
    - Out of scope: resolver semantics.
  - Non-Goals: no endpoint-only dedupe.
  - Pre-flight Questions:
    - Data source: relationship producers and source-site facts.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read in-memory relationship/source-site inputs.
    - DB write flow: write only in-memory RelationshipID and lossless aggregation results.
    - Render location: Anvien evidence ledger and conservation manifest.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real relationship aggregation/conservation boundary after the full build.
    - Behavior test: same endpoints with distinct source sites, meanings, ordinals, and provenance.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores and remove superseded debug artifacts before review.
    - External side effects: None.
    - N/A notes: In-memory semantic owner; no live reader or storage.
  - Work Steps:
    1. record impacts; implement RelationshipID/aggregation; add tests; run full build; validate conservation; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: every source-site ID survives aggregation.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1C0A-IMPACT1`, `E1-P1C0A-SRC1`, `E1-P1C0A-BUILD1`, `E1-P1C0A-TEST1`.
  - Implementation Gate: P1-C0 PASS and tuple/aggregation authority accepted.
  - Acceptance:
    - Source: same-endpoint occurrences remain distinct or losslessly aggregated.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: same-endpoint occurrences remain distinct or losslessly aggregated.
    - Behavior test: The behavior gate is: same-endpoint occurrences remain distinct or losslessly aggregated.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E1-P1C0A-IMPACT1`, `E1-P1C0A-SRC1`, `E1-P1C0A-BUILD1`, `E1-P1C0A-TEST1`, `E1-P1C0A-REVIEW1`
    - Actual-status rows refreshed: RelationshipID/aggregation `missing -> correct`.
  - Evidence Targets: tuple/conservation source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: RelationshipID/aggregation `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-C0B: Validate lossless graph decode and closure.
  - Goal: reject duplicate/conflicting IDs, missing endpoints, and empty/subset graph snapshots before any v2 shadow publication.
  - Scope Boundary:
    - Editable: graph decode/validation owner.
    - Inspect-only: occurrence/relationship builders.
    - Preserve-only: active v1 reader.
    - Out of scope: producer migration.
  - Non-Goals: no repair or first/last-wins fallback.
  - Pre-flight Questions:
    - Data source: isolated Graph JSON/Ladybug decode and closure fixtures.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: decode isolated snapshots only.
    - DB write flow: N/A — invalid snapshots are rejected before any publication.
    - Render location: structured duplicate/conflict/endpoint/closure error.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real decoder fault matrix boundary after the full build.
    - Behavior test: every duplicate ID, conflicting payload, missing endpoint, empty/subset, and generation mismatch case.
    - Cleanup/quarantine: isolated repo-local stores.
    - External side effects: None.
    - N/A notes: No repair, live reader, or publication path is owned.
  - Work Steps:
    1. record impacts; implement validation; add fault tests; run full build; validate fail-closed decode; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no invalid snapshot reaches a reader.
       - Render location check: structured error.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1C0B-IMPACT1`, `E1-P1C0B-SRC1`, `E1-P1C0B-BUILD1`, `E1-P1C0B-TEST1`.
  - Implementation Gate: P1-C0/C0A PASS.
  - Acceptance:
    - Source: duplicate/conflict/endpoint/subset matrix fails closed.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: duplicate/conflict/endpoint/subset matrix fails closed.
    - Behavior test: The behavior gate is: duplicate/conflict/endpoint/subset matrix fails closed.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E1-P1C0B-IMPACT1`, `E1-P1C0B-SRC1`, `E1-P1C0B-BUILD1`, `E1-P1C0B-TEST1`, `E1-P1C0B-REVIEW1`
    - Actual-status rows refreshed: decode closure `wrong -> correct`.
  - Evidence Targets: decode/closure source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: decode closure `wrong -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-C: Build declaration-to-symbol identity mapping.
  - Goal: deterministically map every declaration occurrence to one DeclarationID and an evidence-backed logical SymbolID.
  - Scope Boundary:
    - Editable: declaration/symbol identity owners and narrow ScopeIR/provider adapters.
    - Inspect-only: graph emission and resolution.
    - Preserve-only: active v1 graph and non-TS language behavior.
    - Out of scope: strict graph mutation and export resolution.
  - Non-Goals: no heuristic declaration merging presented as authoritative.
  - Pre-flight Questions:
    - Data source: ScopeIR declarations, lexical owner chain, meaning, binding path, and canonical module.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read in-memory declaration facts.
    - DB write flow: write only in-memory DeclarationID/SymbolID mapping results.
    - Render location: identity-set comparison and evidence output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real deterministic identity mapping boundary after the full build.
    - Behavior test: same-name locals, overloads, evidence-backed merging, anonymous/default forms, body edits, rename, owner change, and move.
    - Cleanup/quarantine: independently authored package `testdata`.
    - External side effects: None.
    - N/A notes: In-memory mapping only; active graph output is preserved.
  - Work Steps:
    1. Record impacts and implement production canonical tuples, stability tiers, merge evidence, and collision diagnostics.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: mapping is deterministic independent of input order.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1C-IMPACT1`, `E1-P1C-SRC1`.
    2. Add tests after production code; run full build and five order/worker permutations at the identity boundary.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: exact DeclarationID/SymbolID sets compared.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1C-BUILD1`, `E1-P1C-TEST1`.
  - Implementation Gate: P1-B, P1-C0, P1-C0A, and P1-C0B are correct; provider merge evidence policy accepted.
  - Acceptance:
    - Source: every declaration has one occurrence identity; same-name scopes do not collide.
    - Runtime/UI: N/A — this slice changes no UI; the named real non-UI boundary passes after the full build.
    - DB/data: deterministic sets and explicit ambiguous/unverified merges.
    - Behavior test: complete identity matrix passes.
    - Cleanup/quarantine: no target-derived fixture.
    - Evidence IDs: `E1-P1C-IMPACT1`, `E1-P1C-SRC1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-REVIEW1`.
    - Actual-status rows refreshed: declaration/symbol mapping.
  - Evidence Targets: canonical tuple diff, determinism runs, full build, Supervisor, detect-changes, commit.
  - Actual-status Update: mapping `missing -> correct`; graph projection remains unbound.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-D: Introduce strict graph mutation operations and validation.
  - Goal: make insert, idempotent enrich/update, replace, and decode behavior explicit and fail closed on conflicting identities before producer migration.
  - Scope Boundary:
    - Editable: graph mutation/validation owners only.
    - Inspect-only: all current node/relationship producers.
    - Preserve-only: producer semantics not required by explicit operation selection.
    - Out of scope: producer migration and ID v2 activation.
  - Non-Goals: no generic upsert and no ignored error return.
  - Pre-flight Questions:
    - Data source: the complete graph producer inventory and provenance.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read in-memory graph/decode fixtures and current producer intent.
    - DB write flow: write only through the strict in-memory builder; caller migration is excluded.
    - Render location: structured mutation/decode errors and evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real graph mutation/decode boundary after the full build.
    - Behavior test: insert, identical idempotence, conflict, enrich, replace, missing endpoint, duplicate relationship, and ignored-error prevention.
    - Cleanup/quarantine: graph-package fault fixtures.
    - External side effects: None.
    - N/A notes: No producer or UI migration occurs.
  - Work Steps:
    1. Freeze a producer inventory with exact path, symbol, current operation, allowed operation, provenance, and planned migration slice; implement strict builder/validator without migrating callers.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: build and decode both reject conflicts.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1D-IMPACT1`, `E1-P1D-MAP1`, `E1-P1D-SRC1`.
    2. Add negative tests after code, run full build, and inject each conflict/closure failure.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no partial graph escapes.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1D-BUILD1`, `E1-P1D-TEST1`.
  - Implementation Gate: all producers are classified and assigned to P1-D1/D2/D3; build path cannot ignore mutation failures.
  - Acceptance:
    - Source: strict insert/enrich/replace/decode APIs and a committed producer inventory exist; no producer migration is implied.
    - Runtime/UI: N/A — this slice changes no UI; the named real non-UI boundary passes after the full build.
    - DB/data: `0` silent duplicate/endpoint failures; structured build failure.
    - Behavior test: full negative matrix passes.
    - Cleanup/quarantine: no obsolete compatibility helper remains.
    - Evidence IDs: `E1-P1D-IMPACT1`, `E1-P1D-MAP1`, `E1-P1D-SRC1`, `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-REVIEW1`.
    - Actual-status rows refreshed: mutation/decode strictness.
  - Evidence Targets: producer map, source diff, full build, fault tests, Supervisor, detect-changes, commit.
  - Actual-status Update: duplicate behavior `wrong -> correct`; active IDs remain v1 and the strict API is reachable only through the shadow-v2 adapter.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-D1: Migrate core graph producers to explicit operations.
  - Goal: migrate only the core node/relationship producers named in the P1-D inventory.
  - Scope Boundary:
    - Editable: exact core producer allowlist.
    - Inspect-only: all other families.
    - Preserve-only: business payloads.
    - Out of scope: identity cutover.
  - Non-Goals: no bulk search/replace and no active-v1 byte/hash change; route every migrated caller through the explicit shadow-v2 adapter, never directly into the active-v1 compatibility path.
  - Pre-flight Questions:
    - Data source: the exact P1-D core-producer allowlist.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read core producer inputs and the active-v1 byte/hash baseline.
    - DB write flow: write only shadow-v2 explicit-operation results; active v1 remains unchanged.
    - Render location: core producer conservation evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real core producer graph boundary after the full build.
    - Behavior test: every allowlisted core producer, payload/provenance conservation, counts, closure, and ignored-error prevention.
    - Cleanup/quarantine: remove temporary migration adapters.
    - External side effects: None.
    - N/A notes: Public behavior and active-v1 bytes remain preserve-only.
  - Work Steps:
    1. record file/symbol impacts; migrate production callers; add focused tests after code; run full build; validate counts/closure; update ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: payload/provenance unchanged.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`.
  - Implementation Gate: P1-D PASS and exact core allowlist attached.
  - Acceptance:
    - Source: all core producers use explicit operations, no unrelated caller changes, evidence `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`, `E1-P1D1-REVIEW1`, and actual-status refresh.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: all core producers use explicit operations, no unrelated caller changes, evidence `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`, `E1-P1D1-REVIEW1`, and actual-status refresh.
    - Behavior test: The behavior gate is: all core producers use explicit operations, no unrelated caller changes, evidence `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`, `E1-P1D1-REVIEW1`, and actual-status refresh.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`, `E1-P1D1-REVIEW1`
    - Actual-status rows refreshed: core producer migration `unbound -> bound-correct`.
  - Evidence Targets: allowlist, source/build/tests, Supervisor, detect-changes, commit.
  - Actual-status Update: core producer migration `unbound -> bound-correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-D2: Migrate resolution/projection producers to explicit operations.
  - Goal: migrate only resolver, emitter, health, and persistence producers named in the P1-D inventory.
  - Scope Boundary:
    - Editable: exact resolution/projection allowlist.
    - Inspect-only: other families.
    - Preserve-only: outcome payloads.
    - Out of scope: TS semantic fixes.
  - Non-Goals: no resolver behavior change and no active-v1 byte/hash change; route every migrated caller through the explicit shadow-v2 adapter, never directly into the active-v1 compatibility path.
  - Pre-flight Questions:
    - Data source: the exact P1-D resolution/projection producer allowlist.
    - Display permission: Preserve current public visibility; no new display permission is introduced.
    - DB read flow: read resolver facts plus canonical Graph JSON/Ladybug parity inputs.
    - DB write flow: write only shadow-v2 projection operations; active v1 remains unchanged.
    - Render location: projection/parity evidence; existing public surface only if fields actually change.
    - UI behavior flow: Conditional — if a public field changes, exercise its real built surface; otherwise no browser-visible behavior is owned.
    - Docker runtime: Conditional — mandatory only if this slice changes a public runtime field; otherwise validate JSON/Ladybug after the full build.
    - Playwright target: Conditional — use the real built runtime when public fields change; otherwise exercise JSON/Ladybug parity.
    - Behavior test: every allowlisted producer, canonical field/provenance conservation, duplicate/closure failure, and v1 hash preservation.
    - Cleanup/quarantine: remove temporary migration adapters.
    - External side effects: None.
    - N/A notes: Public adapters are preserve-only unless the actual-status refresh explicitly opens them.
  - Work Steps:
    1. record impacts; migrate production callers; add focused tests after code; run full build; validate canonical projection/closure; update ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no dropped source-site/provenance fields.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`.
  - Implementation Gate: P1-D1 PASS and exact resolution/projection allowlist attached.
  - Acceptance:
    - Source: allowlisted producers use explicit operations and parity is unchanged, evidence `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`, and actual-status refresh.
    - Runtime/UI: Conditional — if a public field changes, exercise its real built surface; otherwise no browser-visible behavior is owned.
    - DB/data: The data/contract condition is: allowlisted producers use explicit operations and parity is unchanged, evidence `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`, and actual-status refresh.
    - Behavior test: The behavior gate is: allowlisted producers use explicit operations and parity is unchanged, evidence `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`, and actual-status refresh.
    - Cleanup/quarantine: The boundary/cleanup condition is: allowlisted producers use explicit operations and parity is unchanged, evidence `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`, and actual-status refresh.
    - Evidence IDs: `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`
    - Actual-status rows refreshed: resolution/projection migration `unbound -> bound-correct`.
  - Evidence Targets: allowlist, source/build/tests/parity, Supervisor, detect-changes, commit.
  - Actual-status Update: resolution/projection migration `unbound -> bound-correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-D3: Migrate ancillary/document/semantic producers to explicit operations.
  - Goal: migrate only remaining document, COBOL, semantic, and diagnostic producers named in the inventory.
  - Scope Boundary:
    - Editable: exact ancillary allowlist.
    - Inspect-only: prior families.
    - Preserve-only: labels/payloads.
    - Out of scope: new semantics.
  - Non-Goals: no unrelated cleanup and no active-v1 byte/hash change; route every migrated caller through the explicit shadow-v2 adapter, never directly into the active-v1 compatibility path.
  - Pre-flight Questions:
    - Data source: the exact P1-D ancillary/document/semantic producer allowlist.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read document, COBOL, semantic, and diagnostic facts.
    - DB write flow: write only shadow-v2 explicit operations for the allowlisted producers.
    - Render location: ancillary producer evidence ledger.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real ancillary producer graph boundary after the full build.
    - Behavior test: every allowlisted family, payload/provenance conservation, duplicate/closure regression, and no skipped row.
    - Cleanup/quarantine: remove temporary migration adapters.
    - External side effects: None.
    - N/A notes: No public surface or new semantics are owned.
  - Work Steps:
    1. record impacts; migrate production callers; add focused tests after code; run full build; validate no duplicate/closure regressions; update ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: ancillary rows retained.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1D3-IMPACT1`, `E1-P1D3-SRC1`, `E1-P1D3-BUILD1`, `E1-P1D3-TEST1`.
  - Implementation Gate: P1-D2 PASS and exact ancillary allowlist attached.
  - Acceptance:
    - Source: all remaining producers use explicit operations, no new responsibility or skipped row, evidence `E1-P1D3-IMPACT1`, `E1-P1D3-SRC1`, `E1-P1D3-BUILD1`, `E1-P1D3-TEST1`, `E1-P1D3-REVIEW1`, and actual-status refresh.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: all remaining producers use explicit operations, no new responsibility or skipped row, evidence `E1-P1D3-IMPACT1`, `E1-P1D3-SRC1`, `E1-P1D3-BUILD1`, `E1-P1D3-TEST1`, `E1-P1D3-REVIEW1`, and actual-status refresh.
    - Behavior test: The behavior gate is: all remaining producers use explicit operations, no new responsibility or skipped row, evidence `E1-P1D3-IMPACT1`, `E1-P1D3-SRC1`, `E1-P1D3-BUILD1`, `E1-P1D3-TEST1`, `E1-P1D3-REVIEW1`, and actual-status refresh.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E1-P1D3-IMPACT1`, `E1-P1D3-SRC1`, `E1-P1D3-BUILD1`, `E1-P1D3-TEST1`, `E1-P1D3-REVIEW1`
    - Actual-status rows refreshed: ancillary migration `unbound -> bound-correct`.
  - Evidence Targets: allowlist, source/build/tests, Supervisor, detect-changes, commit.
  - Actual-status Update: ancillary migration `unbound -> bound-correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P1-E: Emit and validate shadow identity v2.
  - Goal: produce a non-active v2 graph signature beside v1 so identity correctness and size can be measured before reader migration.
  - Scope Boundary:
    - Editable: v2 emitter/shadow comparison owner and narrow analyze hook.
    - Inspect-only: active snapshot, Ladybug, consumers.
    - Preserve-only: active v1 files and target graph contract.
    - Out of scope: cutover or dual-active reads.
  - Non-Goals: shadow output is evidence only, never a fallback.
  - Pre-flight Questions:
    - Data source: P1-C0 lossless occurrences, P1-C identities, and P1-D/D1/D2/D3 strict producers.
    - Display permission: Existing graph permission only if an optional parity read is used; shadow data is never publicly exposed.
    - DB read flow: read the active-v1 baseline and target source/ScopeIR for an in-memory oracle comparison.
    - DB write flow: write self-repo shadow debug output under `E:\Anvien\.tmp` only; never publish target v2/Ladybug artifacts.
    - Render location: benchmark, evidence, boundary, and oracle artifacts under Anvien.
    - UI behavior flow: N/A for shadow runs — the graph is deliberately non-active; an optional existing graph parity read does not expose shadow data.
    - Docker runtime: N/A unless the existing graph UI is explicitly used for a parity read.
    - Playwright target: N/A unless the existing graph UI is explicitly used; otherwise exercise the shadow comparator and source oracle.
    - Behavior test: five deterministic signatures, occurrence/relationship conservation, empty/subset/conflict failures, node expansion, and target `4/4` source-oracle agreement.
    - Cleanup/quarantine: remove `E:\Anvien\.tmp` shadow/debug output after ledger capture.
    - External side effects: Read/analyze target source in place for the in-memory oracle only; active target `.anvien` and source/worktree remain unchanged.
    - N/A notes: Shadow is intentionally non-active; no target v2 publication or reader consumption.
  - Work Steps:
    1. Implement shadow v2 emission and canonical signature comparison without altering active output.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: shadow validates before it is written.
       - Render location check: benchmark/evidence ledgers only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1E-IMPACT1`, `E1-P1E-SRC1`, `E1-P1E-SHADOW1`.
    2. Add tests, run full build, five deterministic self-repo shadow runs, and measure node/edge/size/RSS plus ScopeIR->Declaration, Declaration->Symbol, source-site aggregation, endpoint, and provenance conservation.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: exact set comparison; empty/subset shadow graphs fail even when signatures repeat.
       - Render location check: benchmark ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1E-BUILD1`, `E1-P1E-BENCH1`.
    3. Before P2-G is eligible, capture the target's active-v1 pre-state/hash and run an in-memory v2 shadow from the real `E:\cheapapp.org` source/ScopeIR without writing v2 artifacts to the target. The `4/4` gate is an independent source/TS-oracle comparison of the v2 shadow (two `time` and two `now` occurrences, ranges/scopes, and relationship provenance); it is not an assertion that the known v1 graph already contains `4/4`. Capture target post-state, graph hash/generation, ignored-guidance timestamp caveat, and contamination manifest.
       - UI flow check: N/A unless the existing graph UI is used as a parity read.
       - DB/data flow check: active v1 graph remains unchanged under `E:\cheapapp.org\.anvien`; v2 shadow exists only in memory/Anvien evidence and is never published to the target.
       - Render location check: oracle/report artifacts remain under `E:\Anvien`.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E1-P1E-TARGET1`, `E1-P1E-ORACLE1`, `E1-P1E-BOUNDARY1`.
  - Implementation Gate: P1-C0/C0A/C0B/P1-C/P1-D/D1/D2/D3 correct and committed; active-v1 byte/hash gate passes; shadow path cannot be consumed as active; target `4/4` is proven in memory only.
  - Acceptance:
    - Source: v2 shadow uses only approved identities.
    - Runtime/UI: N/A — this slice changes no UI; the named real non-UI boundary passes after the full build.
    - DB/data: five identical signatures, no duplicate/orphan, active-v1 hash unchanged, and every shadow source occurrence/source-site relationship is conserved with endpoint/provenance.
    - Behavior test: shadow comparison, conservation, subset/empty, and failure cases pass.
    - Cleanup/quarantine: debug shadows removed after ledger capture.
    - Evidence IDs: `E1-P1E-IMPACT1`, `E1-P1E-SRC1`, `E1-P1E-SHADOW1`, `E1-P1E-TARGET1`, `E1-P1E-ORACLE1`, `E1-P1E-BOUNDARY1`, `E1-P1E-BUILD1`, `E1-P1E-BENCH1`, `E1-P1E-REVIEW1`.
    - Actual-status rows refreshed: v2 projection `unbound -> partial`.
  - Evidence Targets: shadow signatures, graph counts/size/RSS, build/tests, Supervisor, detect-changes, commit.
  - Actual-status Update: record measured migration expansion and update P2 budgets.
  - Commit Boundary: commit after this slice when acceptance passes.

### P2: Versioned persistence, opaque consumers, atomic generation, and v2 cutover

- Phase Goal: migrate each persistence, reader, consumer, and publication owner independently, then activate identity v2 without mixed versions or generations.
- Phase Boundary:
  - In scope: compatibility contract and inventory; per-surface guards; Graph JSON/Ladybug/native/fallback persistence; opaque consumer families; HTTP/Web; immutable generation publication, registries/groups/caches/embeddings, leases/GC, failure atomicity, and cutover.
  - Out of scope: TypeScript binding, export, barrel, and ambient/external semantic fixes.
  - Dependencies: P1 shadow accepted; Owner explicitly opens each ordered P2 slice after its predecessor is accepted and committed.
- Phase Implementation Rule: do not implement `P2` directly; every capability/runtime/store boundary below is an independent slice, Supervisor gate, and commit. Validation-only slices never repair code.
- Ordered Slice List:
  - P2-A: Define the index compatibility manifest and failure contract.
  - P2-A1: Freeze the source-derived reader inventory and owner assignments.
  - P2-A2: Guard Graph JSON and repository-metadata readers.
  - P2-A3: Guard native Ladybug readers.
  - P2-A4: Guard Go/fallback Cypher readers.
  - P2-A5: Guard CLI readers and dispatch boundaries.
  - P2-A6: Guard MCP resources and tools.
  - P2-A7: Guard HTTP handlers and streams.
  - P2-A8: Guard Web readers, streams, and lifecycle clients.
  - P2-A9: Guard file-context cache readers.
  - P2-A10: Guard HTTP/MCP resource-cache readers.
  - P2-A11: Guard embedding readers and jobs.
  - P2-A12: Guard global repository-registry readers.
  - P2-A13: Guard group registry and contract readers.
  - P2-A14: Guard process projection readers.
  - P2-A15: Guard community and cluster projection readers.
  - P2-B: Make Graph JSON v2 codec and decode closure-safe.
  - P2-B1: Write the Ladybug v2 schema and CSV export deterministically.
  - P2-B2: Load Ladybug v2 transactionally and fail closed.
  - P2-B3: Project canonical v2 records through native Ladybug queries.
  - P2-B4: Project canonical v2 records through the Go fallback query path.
  - P2-C: Remove semantic ID parsing from CLI readers.
  - P2-C1: Remove semantic ID parsing from MCP resources and tools.
  - P2-C2: Make file-context projections use explicit canonical fields.
  - P2-C3: Make file-context cache records generation/config/catalog-bound.
  - P2-C4: Make rename use source anchors instead of parsed IDs.
  - P2-C5: Make the shared HTTP/MCP resource cache preserve canonical records.
  - P2-C6: Make embedding references generation-qualified and ID-opaque.
  - P2-D: Make group contracts use generation-qualified opaque references.
  - P2-D1: Make process projections source-anchored and ID-opaque.
  - P2-D2: Make community projections source-anchored and ID-opaque.
  - P2-E: Expose version/generation and canonical fields through HTTP.
  - P2-E1: Negotiate and render version/generation truthfully in Web.
  - P2-E2: Freeze the pre-cutover S0-S11 canonical baseline.
  - P2-F: Stage immutable repo-local generation artifacts.
  - P2-F1: Publish the repo-local active generation atomically.
  - P2-F2: Publish cache and embedding namespaces by generation.
  - P2-F3: Publish the global repository registry atomically.
  - P2-F4: Publish group snapshots and member-generation vectors atomically.
  - P2-F5: Enforce reader leases and lease-safe generation garbage collection.
  - P2-F6: Run the complete publication failure-atomicity matrix.
  - P2-G: Cut over to identity v2 and enforce legacy ambiguity.
- [ ] P2-A: Define the index compatibility manifest and failure contract.
  - Goal: create one versioned compatibility decision that every reader can call without wiring any reader family yet.
  - Scope Boundary:
    - Editable: the dedicated index-version/compatibility contract owner and narrow manifest serialization types.
    - Inspect-only: current repository metadata, graph/Ladybug headers, and the reader-matrix seed.
    - Preserve-only: active v1 bytes, reader behavior, and every consumer adapter.
    - Out of scope: reader wiring, identity-field migration, publication, and cutover.
  - Non-Goals: no version inference from opaque IDs, warning-only compatibility, or reader-family edits.
  - Pre-flight Questions:
    - Data source: the accepted P1 identity contract, current repository metadata, and the handshake fields fixed in this plan.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: inspect existing meta/manifest shapes in isolated fixtures; do not open active graph bodies through the new policy yet.
    - DB write flow: serialize only isolated manifest fixtures; do not mutate the active repository index.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: absent, supported, unsupported, mixed, stale-generation, old-reader, and malformed handshake decisions with the exact failure envelope.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Refresh the self graph, record file-detail/upstream impact for the compatibility and metadata owners, then implement the production manifest, request, decision, and `INDEX_VERSION_MISMATCH` envelope without wiring readers.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: manifest serialization is deterministic and preserves active-v1 bytes.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real compatibility decision/manifest serialization boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real compatibility decision/manifest serialization probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A-IMPACT1`, `E2-P2A-SRC1`.
    2. After production code is complete, add the decision-table tests, run the full build, and exercise manifest encode/decode at the real repository metadata boundary.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: every invalid combination fails before a graph/database/cache body can be opened.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real compatibility decision/manifest serialization boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real compatibility decision/manifest serialization probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A-BUILD1`, `E2-P2A-TEST1`.
  - Implementation Gate: P1-E is accepted and committed; P1-A authority fixes the exact request, manifest, minimum-reader, generation, and error-envelope fields.
  - Acceptance:
    - Source: one compatibility authority owns the typed request/manifest/decision; no reader adapter is mixed into the owner.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: all decision-table rows are deterministic, fail closed, and active-v1 serialization is unchanged.
    - Behavior test: positive and negative compatibility vectors pass with no downgrade or ID inference.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A-IMPACT1`, `E2-P2A-SRC1`, `E2-P2A-BUILD1`, `E2-P2A-TEST1`, `E2-P2A-REVIEW1`.
    - Actual-status rows refreshed: index compatibility contract.
  - Evidence Targets: source diff, decision-table tests, full build, repository-metadata probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: compatibility authority `missing -> correct`; every reader surface remains unbound.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A1: Freeze the source-derived reader inventory and owner assignments.
  - Goal: turn the checked-in seed into a complete, source-backed reader/non-reader matrix before any family is wired.
  - Scope Boundary:
    - Editable: `index-reader-matrix.md` and the four plan ledgers only.
    - Inspect-only: all exact CLI, MCP, HTTP, Web, loader, query, cache, embedding, registry, group, process, and community entrypoints.
    - Preserve-only: all production source and the target repository.
    - Out of scope: compatibility implementation or reader wiring.
  - Non-Goals: no catch-all row, parent dispatcher substitution, or completeness claim based only on the seed.
  - Pre-flight Questions:
    - Data source: a fresh Anvien self graph plus a fresh source audit of every graph/index open, query, stream, cache, registry, projection, and dispatcher path.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read source and current graph topology only; classify the real backend/layout of every entrypoint.
    - DB write flow: N/A — this documentation-only slice writes no graph/index state.
    - Render location: the committed matrix and evidence ledger under `E:\Anvien`.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: row-ID continuity, unique anchor/owner assignment, anchor existence, truthful backend classification, and `unlisted_readers == 0`.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Re-run the source scan, verify every exact path/function anchor and backend, classify dispatchers/non-readers, assign each reader to exactly one later guard-owner slice, and update the matrix without changing production code.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: every backend/layout and later guard owner is explicit; native and fallback paths remain separate.
       - Render location check: the committed matrix and evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real source/matrix audit boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real source/matrix audit probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A1-MATRIX1`, `E2-P2A1-R01..E2-P2A1-R195`.
    2. Run independent matrix review for continuity, duplicates, missing anchors, unassigned rows, and unlisted readers; update row IDs if fresh source adds readers, then commit only the frozen matrix/ledger slice.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: `rows_classified == rows_total`, `unassigned_rows == 0`, and `unlisted_readers == 0`.
       - Render location check: the matrix review record under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real source/matrix audit boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real source/matrix audit probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A1-MATRIXREVIEW1`, `E2-P2A1-REVIEW1`.
  - Implementation Gate: P2-A contract is accepted; the self graph is fresh; no production file is edited in this slice.
  - Acceptance:
    - Source: every current reader, backend, dispatcher, and explicit non-reader has one truthful exact row and one later owner slice.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: row IDs are contiguous, anchors exist, backend semantics match source, and no reader is unlisted or multiply owned.
    - Behavior test: the structural/source audit passes every row; `pending` still means guard runtime evidence is not yet produced.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A1-MATRIX1`, `E2-P2A1-MATRIXREVIEW1`, `E2-P2A1-R01..E2-P2A1-R195`, `E2-P2A1-REVIEW1`, `E2-P2A1-COMMIT1`.
    - Actual-status rows refreshed: reader inventory and owner assignment.
  - Evidence Targets: source-scan output, exact row proofs, matrix hash, independent review, and documentation-only commit.
  - Actual-status Update: reader matrix `missing -> correct inventory`; runtime guards remain unbound.
  - Commit Boundary: commit only the frozen matrix and ledgers after Supervisor PASS; no implementation detect-changes is required for this documentation-only slice.

- [ ] P2-A2: Guard Graph JSON and repository-metadata readers.
  - Goal: wire the common compatibility decision into every exact S0 row owned by repository metadata and Graph JSON loader owners, without changing other surfaces.
  - Scope Boundary:
    - Editable: only repository metadata and Graph JSON loader owners compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to repository metadata and Graph JSON loader owners and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: repo-local manifest and Graph JSON body; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S0 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S0 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real Graph JSON/repository loader boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before repo-local manifest and Graph JSON body; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S0 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real Graph JSON/repository loader boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real Graph JSON/repository loader probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A2-IMPACT1`, `E2-P2A2-SRC1`, `E2-P2A2-BUILD1`, `E2-P2A2-TEST1`, `E2-P2A2-S0GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S0 row for repository metadata and Graph JSON loader owners is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S0 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S0 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A2-IMPACT1`, `E2-P2A2-SRC1`, `E2-P2A2-BUILD1`, `E2-P2A2-TEST1`, `E2-P2A2-S0GUARD1`, `E2-P2A2-REVIEW1`.
    - Actual-status rows refreshed: S0 compatibility guard rows owned by repository metadata and Graph JSON loader owners.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real Graph JSON/repository loader probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S0 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A3: Guard native Ladybug readers.
  - Goal: wire the common compatibility decision into every exact S1 row owned by native Ladybug driver/query owners, without changing other surfaces.
  - Scope Boundary:
    - Editable: only native Ladybug driver/query owners compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to native Ladybug driver/query owners and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: native Ladybug database; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S1 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S1 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real native Ladybug driver/query boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before native Ladybug database; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S1 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real native Ladybug driver/query boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real native Ladybug driver/query probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A3-IMPACT1`, `E2-P2A3-SRC1`, `E2-P2A3-BUILD1`, `E2-P2A3-TEST1`, `E2-P2A3-S1GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S1 row for native Ladybug driver/query owners is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S1 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S1 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A3-IMPACT1`, `E2-P2A3-SRC1`, `E2-P2A3-BUILD1`, `E2-P2A3-TEST1`, `E2-P2A3-S1GUARD1`, `E2-P2A3-REVIEW1`.
    - Actual-status rows refreshed: S1 compatibility guard rows owned by native Ladybug driver/query owners.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real native Ladybug driver/query probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S1 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A4: Guard Go/fallback Cypher readers.
  - Goal: wire the common compatibility decision into every exact S2 row owned by Go fallback selector/query owners, without changing other surfaces.
  - Scope Boundary:
    - Editable: only Go fallback selector/query owners compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to Go fallback selector/query owners and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: Graph JSON plus the explicit in-memory fallback backend; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S2 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S2 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real fallback query boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before Graph JSON plus the explicit in-memory fallback backend; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S2 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real fallback query boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real fallback query probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A4-IMPACT1`, `E2-P2A4-SRC1`, `E2-P2A4-BUILD1`, `E2-P2A4-TEST1`, `E2-P2A4-S2GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S2 row for Go fallback selector/query owners is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S2 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S2 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A4-IMPACT1`, `E2-P2A4-SRC1`, `E2-P2A4-BUILD1`, `E2-P2A4-TEST1`, `E2-P2A4-S2GUARD1`, `E2-P2A4-REVIEW1`.
    - Actual-status rows refreshed: S2 compatibility guard rows owned by Go fallback selector/query owners.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real fallback query probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S2 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A5: Guard CLI readers and dispatch boundaries.
  - Goal: wire the common compatibility decision into every exact S3 row owned by CLI command and exact child-entrypoint adapters, without changing other surfaces.
  - Scope Boundary:
    - Editable: only CLI command and exact child-entrypoint adapters compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to CLI command and exact child-entrypoint adapters and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: the row-specific graph, query, registry, or metadata backend; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S3 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S3 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real packaged CLI boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before the row-specific graph, query, registry, or metadata backend; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S3 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real packaged CLI boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real packaged CLI probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A5-IMPACT1`, `E2-P2A5-SRC1`, `E2-P2A5-BUILD1`, `E2-P2A5-TEST1`, `E2-P2A5-S3GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S3 row for CLI command and exact child-entrypoint adapters is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S3 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S3 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A5-IMPACT1`, `E2-P2A5-SRC1`, `E2-P2A5-BUILD1`, `E2-P2A5-TEST1`, `E2-P2A5-S3GUARD1`, `E2-P2A5-REVIEW1`.
    - Actual-status rows refreshed: S3 compatibility guard rows owned by CLI command and exact child-entrypoint adapters.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real packaged CLI probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S3 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A6: Guard MCP resources and tools.
  - Goal: wire the common compatibility decision into every exact S4 row owned by MCP server, resource, tool, and exact dispatcher adapters, without changing other surfaces.
  - Scope Boundary:
    - Editable: only MCP server, resource, tool, and exact dispatcher adapters compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to MCP server, resource, tool, and exact dispatcher adapters and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: the row-specific graph, query, cache, registry, or group backend; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S4 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S4 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real MCP JSON-RPC boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before the row-specific graph, query, cache, registry, or group backend; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S4 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real MCP JSON-RPC boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real MCP JSON-RPC probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A6-IMPACT1`, `E2-P2A6-SRC1`, `E2-P2A6-BUILD1`, `E2-P2A6-TEST1`, `E2-P2A6-S4GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S4 row for MCP server, resource, tool, and exact dispatcher adapters is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S4 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S4 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A6-IMPACT1`, `E2-P2A6-SRC1`, `E2-P2A6-BUILD1`, `E2-P2A6-TEST1`, `E2-P2A6-S4GUARD1`, `E2-P2A6-REVIEW1`.
    - Actual-status rows refreshed: S4 compatibility guard rows owned by MCP server, resource, tool, and exact dispatcher adapters.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real MCP JSON-RPC probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S4 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A7: Guard HTTP handlers and streams.
  - Goal: wire the common compatibility decision into every exact S5 row owned by HTTP middleware, exact handlers, and stream adapters, without changing other surfaces.
  - Scope Boundary:
    - Editable: only HTTP middleware, exact handlers, and stream adapters compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to HTTP middleware, exact handlers, and stream adapters and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: the row-specific graph, fallback query, cache, registry, job, or session backend; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S5 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S5 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real HTTP handler/stream boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before the row-specific graph, fallback query, cache, registry, job, or session backend; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S5 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real HTTP handler/stream boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real HTTP handler/stream probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A7-IMPACT1`, `E2-P2A7-SRC1`, `E2-P2A7-BUILD1`, `E2-P2A7-TEST1`, `E2-P2A7-S5GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S5 row for HTTP middleware, exact handlers, and stream adapters is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S5 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S5 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A7-IMPACT1`, `E2-P2A7-SRC1`, `E2-P2A7-BUILD1`, `E2-P2A7-TEST1`, `E2-P2A7-S5GUARD1`, `E2-P2A7-REVIEW1`.
    - Actual-status rows refreshed: S5 compatibility guard rows owned by HTTP middleware, exact handlers, and stream adapters.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real HTTP handler/stream probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S5 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A8: Guard Web readers, streams, and lifecycle clients.
  - Goal: wire the common compatibility decision into every exact S6 row owned by Web backend-client handshake, parser, and lifecycle adapters, without changing other surfaces.
  - Scope Boundary:
    - Editable: only Web backend-client handshake, parser, and lifecycle adapters compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to Web backend-client handshake, parser, and lifecycle adapters and the P2-A compatibility decision.
    - Display permission: Preserve existing graph-view permission and visibility behavior; compatibility does not grant access.
    - DB read flow: the negotiated HTTP response or stream; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the existing Web mismatch/error state and supported graph surface.
    - UI behavior flow: Supported data proceeds only after handshake; mismatch renders a blocking truthful state and discards partial records.
    - Docker runtime: Mandatory; run the full build, build/start the real Docker runtime, and pin one generation.
    - Playwright target: Mandatory; exercise every assigned Web row against the built Docker URL and visually inspect mismatch/supported states.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Keep reusable Playwright under `playwright/` and official JSON+MD evidence under `Reports/qa/playwright/...`; remove debug-only artifacts.
    - External side effects: None.
    - N/A notes: DB writes are N/A; Docker/Playwright are mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S6 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real built Web runtime boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: exercise supported, absent, stale, mixed, and unsupported states without consuming incompatible records.
       - DB/data flow check: all assigned rows validate one pinned generation before the negotiated HTTP response or stream; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the real Docker-served Web surface and official QA artifacts.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser or Chrome against the real built Docker runtime at real built Web runtime; inspect the visible supported and mismatch states.
         - Playwright: use the reusable repository Playwright suite against the built Docker runtime; retain JSON, Markdown, screenshot, and trace evidence.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A8-IMPACT1`, `E2-P2A8-SRC1`, `E2-P2A8-BUILD1`, `E2-P2A8-PLAY1`, `E2-P2A8-S6GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S6 row for Web backend-client handshake, parser, and lifecycle adapters is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S6 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned Web row passes on the built Docker runtime; incompatible data never renders or enters client state.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only reusable Playwright and official JSON+MD/visual evidence remain; no dev-server evidence is accepted.
    - Evidence IDs: `E2-P2A8-IMPACT1`, `E2-P2A8-SRC1`, `E2-P2A8-BUILD1`, `E2-P2A8-PLAY1`, `E2-P2A8-S6GUARD1`, `E2-P2A8-REVIEW1`.
    - Actual-status rows refreshed: S6 compatibility guard rows owned by Web backend-client handshake, parser, and lifecycle adapters.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real built Web runtime probe, Docker/Playwright and visual inspection, Supervisor, detect-changes, and commit.
  - Actual-status Update: S6 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A9: Guard file-context cache readers.
  - Goal: wire the common compatibility decision into every exact S7 row owned by file-context cache key/read owner, without changing other surfaces.
  - Scope Boundary:
    - Editable: only file-context cache key/read owner compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to file-context cache key/read owner and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: generation/config/catalog-qualified file-context cache; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S7 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S7 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real file-context cache boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before generation/config/catalog-qualified file-context cache; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S7 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real file-context cache boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real file-context cache probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A9-IMPACT1`, `E2-P2A9-SRC1`, `E2-P2A9-BUILD1`, `E2-P2A9-TEST1`, `E2-P2A9-S7GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S7 row for file-context cache key/read owner is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S7 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S7 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A9-IMPACT1`, `E2-P2A9-SRC1`, `E2-P2A9-BUILD1`, `E2-P2A9-TEST1`, `E2-P2A9-S7GUARD1`, `E2-P2A9-REVIEW1`.
    - Actual-status rows refreshed: S7 compatibility guard rows owned by file-context cache key/read owner.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real file-context cache probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S7 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A10: Guard HTTP/MCP resource-cache readers.
  - Goal: wire the common compatibility decision into every exact S8 row owned by shared resource-cache key/read owner, without changing other surfaces.
  - Scope Boundary:
    - Editable: only shared resource-cache key/read owner compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to shared resource-cache key/read owner and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: generation/config/catalog-qualified resource cache; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S8 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S8 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real resource cache boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before generation/config/catalog-qualified resource cache; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S8 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real resource cache boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real resource cache probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A10-IMPACT1`, `E2-P2A10-SRC1`, `E2-P2A10-BUILD1`, `E2-P2A10-TEST1`, `E2-P2A10-S8GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S8 row for shared resource-cache key/read owner is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S8 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S8 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A10-IMPACT1`, `E2-P2A10-SRC1`, `E2-P2A10-BUILD1`, `E2-P2A10-TEST1`, `E2-P2A10-S8GUARD1`, `E2-P2A10-REVIEW1`.
    - Actual-status rows refreshed: S8 compatibility guard rows owned by shared resource-cache key/read owner.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real resource cache probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S8 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A11: Guard embedding readers and jobs.
  - Goal: wire the common compatibility decision into every exact S9 row owned by embedding row/search/job guard owners, without changing other surfaces.
  - Scope Boundary:
    - Editable: only embedding row/search/job guard owners compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to embedding row/search/job guard owners and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: generation/catalog/dimension/hash-qualified embedding store; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S9 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S9 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real embedding store/search boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before generation/catalog/dimension/hash-qualified embedding store; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S9 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real embedding store/search boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real embedding store/search probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A11-IMPACT1`, `E2-P2A11-SRC1`, `E2-P2A11-BUILD1`, `E2-P2A11-TEST1`, `E2-P2A11-S9GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S9 row for embedding row/search/job guard owners is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S9 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S9 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A11-IMPACT1`, `E2-P2A11-SRC1`, `E2-P2A11-BUILD1`, `E2-P2A11-TEST1`, `E2-P2A11-S9GUARD1`, `E2-P2A11-REVIEW1`.
    - Actual-status rows refreshed: S9 compatibility guard rows owned by embedding row/search/job guard owners.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real embedding store/search probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S9 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A12: Guard global repository-registry readers.
  - Goal: wire the common compatibility decision into every exact S10 row owned by global repository-registry read adapters, without changing other surfaces.
  - Scope Boundary:
    - Editable: only global repository-registry read adapters compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to global repository-registry read adapters and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: immutable registry snapshot and active epoch; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S10 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S10 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real repository registry boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before immutable registry snapshot and active epoch; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S10 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real repository registry boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real repository registry probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A12-IMPACT1`, `E2-P2A12-SRC1`, `E2-P2A12-BUILD1`, `E2-P2A12-TEST1`, `E2-P2A12-S10GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S10 row for global repository-registry read adapters is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S10 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S10 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A12-IMPACT1`, `E2-P2A12-SRC1`, `E2-P2A12-BUILD1`, `E2-P2A12-TEST1`, `E2-P2A12-S10GUARD1`, `E2-P2A12-REVIEW1`.
    - Actual-status rows refreshed: S10 compatibility guard rows owned by global repository-registry read adapters.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real repository registry probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S10 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A13: Guard group registry and contract readers.
  - Goal: wire the common compatibility decision into every exact S10 row owned by group registry, contract, status, query, and sync read adapters, without changing other surfaces.
  - Scope Boundary:
    - Editable: only group registry, contract, status, query, and sync read adapters compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to group registry, contract, status, query, and sync read adapters and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: group snapshot and pinned member-generation vector; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S10 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S10 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real group registry/contract boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before group snapshot and pinned member-generation vector; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S10 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real group registry/contract boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real group registry/contract probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A13-IMPACT1`, `E2-P2A13-SRC1`, `E2-P2A13-BUILD1`, `E2-P2A13-TEST1`, `E2-P2A13-S10GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S10 row for group registry, contract, status, query, and sync read adapters is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S10 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S10 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A13-IMPACT1`, `E2-P2A13-SRC1`, `E2-P2A13-BUILD1`, `E2-P2A13-TEST1`, `E2-P2A13-S10GUARD1`, `E2-P2A13-REVIEW1`.
    - Actual-status rows refreshed: S10 compatibility guard rows owned by group registry, contract, status, query, and sync read adapters.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real group registry/contract probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S10 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A14: Guard process projection readers.
  - Goal: wire the common compatibility decision into every exact S11 row owned by process projection/resource/handler read adapters, without changing other surfaces.
  - Scope Boundary:
    - Editable: only process projection/resource/handler read adapters compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to process projection/resource/handler read adapters and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: generation/provenance-bound process projection; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S11 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S11 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real process projection boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before generation/provenance-bound process projection; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S11 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real process projection boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real process projection probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A14-IMPACT1`, `E2-P2A14-SRC1`, `E2-P2A14-BUILD1`, `E2-P2A14-TEST1`, `E2-P2A14-S11GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S11 row for process projection/resource/handler read adapters is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S11 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S11 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A14-IMPACT1`, `E2-P2A14-SRC1`, `E2-P2A14-BUILD1`, `E2-P2A14-TEST1`, `E2-P2A14-S11GUARD1`, `E2-P2A14-REVIEW1`.
    - Actual-status rows refreshed: S11 compatibility guard rows owned by process projection/resource/handler read adapters.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real process projection probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S11 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-A15: Guard community and cluster projection readers.
  - Goal: wire the common compatibility decision into every exact S11 row owned by community/cluster projection/resource/handler read adapters, without changing other surfaces.
  - Scope Boundary:
    - Editable: only community/cluster projection/resource/handler read adapters compatibility adapters and narrowly named behavior tests.
    - Inspect-only: P2-A contract, P2-A1 assigned matrix rows, and the concrete backend implementation.
    - Preserve-only: identity payload semantics, unassigned reader families, and active-v1 compatibility output.
    - Out of scope: other reader surfaces, identity-field migration, atomic publication, and cutover.
  - Non-Goals: no fallback after version mismatch, no parent-row substitution, and no generic cross-surface adapter.
  - Pre-flight Questions:
    - Data source: the frozen P2-A1 rows assigned to community/cluster projection/resource/handler read adapters and the P2-A compatibility decision.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: generation/provenance-bound community projection; validate protocol, schemas, generation, config hash, and catalog hash before open/read/query.
    - DB write flow: N/A — guard wiring does not publish or mutate index data.
    - Render location: the native S11 typed error/output boundary.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every assigned row receives supported, absent, stale, mixed, and unsupported fixtures; expected failure is the exact transport-specific `INDEX_VERSION_MISMATCH`.
    - Cleanup/quarantine: Use isolated package fixtures or stores and remove stale row-result drafts.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact file/symbol/API impacts; wire only the assigned S11 rows to the common guard; after production code is complete add per-row behavior tests; run the full build; exercise the real community/cluster projection boundary; emit a row-result manifest; refresh ledgers, obtain Supervisor PASS, run detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: all assigned rows validate one pinned generation before generation/provenance-bound community projection; mismatch opens zero bodies/rows and cannot fall back.
       - Render location check: the native S11 error/result capture under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real community/cluster projection boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real community/cluster projection probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2A15-IMPACT1`, `E2-P2A15-SRC1`, `E2-P2A15-BUILD1`, `E2-P2A15-TEST1`, `E2-P2A15-S11GUARD1`.
  - Implementation Gate: P2-A and P2-A1 PASS; every S11 row for community/cluster projection/resource/handler read adapters is assigned exactly once and its real backend is known.
  - Acceptance:
    - Source: all assigned S11 readers call the common guard before backend access; no other owner changes.
    - Runtime/UI: Every assigned S11 boundary returns its exact native mismatch contract and nonzero/no-data behavior.
    - DB/data: assigned rows open zero incompatible bodies/rows, never downgrade, and preserve one generation.
    - Behavior test: 100% of assigned rows pass supported and negative fixtures; skipped or unassigned rows are zero.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2A15-IMPACT1`, `E2-P2A15-SRC1`, `E2-P2A15-BUILD1`, `E2-P2A15-TEST1`, `E2-P2A15-S11GUARD1`, `E2-P2A15-REVIEW1`.
    - Actual-status rows refreshed: S11 compatibility guard rows owned by community/cluster projection/resource/handler read adapters.
  - Evidence Targets: owner source diff, per-row result manifest, full build, real community/cluster projection probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: S11 guard wiring `unbound -> correct`; identity payload migration remains pending.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-B: Make Graph JSON v2 codec and decode closure-safe.
  - Goal: round-trip every v2 identity/range/meaning/generation field through Graph JSON and reject invalid snapshots before exposure.
  - Scope Boundary:
    - Editable: Graph JSON codec and decode-validation owner only.
    - Inspect-only: P2-A compatibility types, the preceding storage format, and downstream readers.
    - Preserve-only: other persistence backends, public consumers, and active-v1 compatibility bytes until cutover.
    - Out of scope: reader-family DTO migration, publication, and active cutover.
  - Non-Goals: no silent deduplication, partial load, cross-backend fallback, or unrelated scalar/nullability rewrite.
  - Pre-flight Questions:
    - Data source: the accepted shadow-v2 canonical node/relationship manifest and isolated storage fixtures.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: decode a version-qualified isolated Graph JSON generation.
    - DB write flow: write one deterministic version-qualified Graph JSON fixture/generation.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: field-level round trip, duplicate IDs, conflicting payloads, missing endpoints, subset/empty snapshots, and manifest mismatch.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol impacts; implement only this production storage responsibility; after code add focused round-trip/fault tests; run the full build; exercise the real Graph JSON codec/loader boundary; compare canonical fields and closure; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: Graph JSON canonical fields match the source manifest exactly; invalid snapshots fail before returning a graph.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real Graph JSON codec/loader boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real Graph JSON codec/loader probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2B-IMPACT1`, `E2-P2B-SRC1`, `E2-P2B-BUILD1`, `E2-P2B-TEST1`.
  - Implementation Gate: P2-A2 PASS and the canonical shadow-v2 manifest is fixed.
  - Acceptance:
    - Source: the touched owner contains only this storage responsibility and thin adapters; no other backend is changed.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: Graph JSON canonical fields match the source manifest exactly; invalid snapshots fail before returning a graph.
    - Behavior test: positive round trips and duplicate/conflict/version/orphan negative cases pass with no skipped records.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2B-IMPACT1`, `E2-P2B-SRC1`, `E2-P2B-BUILD1`, `E2-P2B-TEST1`, `E2-P2B-REVIEW1`.
    - Actual-status rows refreshed: Make Graph JSON v2 codec and decode closure-safe.
  - Evidence Targets: source diff, full build, canonical manifest comparison, real Graph JSON codec/loader probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: make graph json v2 codec and decode closure-safe. `wrong/unbound -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-B1: Write the Ladybug v2 schema and CSV export deterministically.
  - Goal: encode canonical v2 node/relationship fields into one Ladybug schema/CSV responsibility without loading the database.
  - Scope Boundary:
    - Editable: Ladybug schema and CSV export owner only.
    - Inspect-only: P2-A compatibility types, the preceding storage format, and downstream readers.
    - Preserve-only: other persistence backends, public consumers, and active-v1 compatibility bytes until cutover.
    - Out of scope: reader-family DTO migration, publication, and active cutover.
  - Non-Goals: no silent deduplication, partial load, cross-backend fallback, or unrelated scalar/nullability rewrite.
  - Pre-flight Questions:
    - Data source: the accepted shadow-v2 canonical node/relationship manifest and isolated storage fixtures.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read the canonical in-memory v2 graph.
    - DB write flow: write deterministic staged CSV/schema artifacts in an isolated repo-local store.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: schema columns/types, escaping, ordering, duplicate detection, ranges, export/access separation, generation, and provenance.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol impacts; implement only this production storage responsibility; after code add focused round-trip/fault tests; run the full build; exercise the real Ladybug CSV/schema writer boundary; compare canonical fields and closure; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: CSV/schema rows are deterministic, lossless, and contain zero silently skipped duplicate or required fields.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real Ladybug CSV/schema writer boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real Ladybug CSV/schema writer probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2B1-IMPACT1`, `E2-P2B1-SRC1`, `E2-P2B1-BUILD1`, `E2-P2B1-TEST1`.
  - Implementation Gate: P2-B PASS; active Ladybug is untouched.
  - Acceptance:
    - Source: the touched owner contains only this storage responsibility and thin adapters; no other backend is changed.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: CSV/schema rows are deterministic, lossless, and contain zero silently skipped duplicate or required fields.
    - Behavior test: positive round trips and duplicate/conflict/version/orphan negative cases pass with no skipped records.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2B1-IMPACT1`, `E2-P2B1-SRC1`, `E2-P2B1-BUILD1`, `E2-P2B1-TEST1`, `E2-P2B1-REVIEW1`.
    - Actual-status rows refreshed: Write the Ladybug v2 schema and CSV export deterministically.
  - Evidence Targets: source diff, full build, canonical manifest comparison, real Ladybug CSV/schema writer probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: write the ladybug v2 schema and csv export deterministically. `wrong/unbound -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-B2: Load Ladybug v2 transactionally and fail closed.
  - Goal: load the staged v2 CSV/schema as one transaction and reject partial, duplicate, orphan, or mixed-generation state.
  - Scope Boundary:
    - Editable: Ladybug loader/transaction owner only.
    - Inspect-only: P2-A compatibility types, the preceding storage format, and downstream readers.
    - Preserve-only: other persistence backends, public consumers, and active-v1 compatibility bytes until cutover.
    - Out of scope: reader-family DTO migration, publication, and active cutover.
  - Non-Goals: no silent deduplication, partial load, cross-backend fallback, or unrelated scalar/nullability rewrite.
  - Pre-flight Questions:
    - Data source: the accepted shadow-v2 canonical node/relationship manifest and isolated storage fixtures.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read P2-B1 staged CSV/schema plus the expected manifest.
    - DB write flow: write only an isolated Ladybug staging database; active database remains untouched.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: transaction rollback at every table/load boundary, duplicate/orphan/version failure, and zero partial visibility.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol impacts; implement only this production storage responsibility; after code add focused round-trip/fault tests; run the full build; exercise the real Ladybug transactional loader boundary; compare canonical fields and closure; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: either every canonical row is queryable in one generation or no new generation is visible; skipped/partial rows are zero.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real Ladybug transactional loader boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real Ladybug transactional loader probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2B2-IMPACT1`, `E2-P2B2-SRC1`, `E2-P2B2-BUILD1`, `E2-P2B2-TEST1`.
  - Implementation Gate: P2-B1 PASS and the staging/rollback contract is explicit.
  - Acceptance:
    - Source: the touched owner contains only this storage responsibility and thin adapters; no other backend is changed.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: either every canonical row is queryable in one generation or no new generation is visible; skipped/partial rows are zero.
    - Behavior test: positive round trips and duplicate/conflict/version/orphan negative cases pass with no skipped records.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2B2-IMPACT1`, `E2-P2B2-SRC1`, `E2-P2B2-BUILD1`, `E2-P2B2-TEST1`, `E2-P2B2-REVIEW1`.
    - Actual-status rows refreshed: Load Ladybug v2 transactionally and fail closed.
  - Evidence Targets: source diff, full build, canonical manifest comparison, real Ladybug transactional loader probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: load ladybug v2 transactionally and fail closed. `wrong/unbound -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-B3: Project canonical v2 records through native Ladybug queries.
  - Goal: make the native query driver return the exact canonical node/relationship record without ID parsing or fallback.
  - Scope Boundary:
    - Editable: native Ladybug read/query projection owner only.
    - Inspect-only: P2-A compatibility types, the preceding storage format, and downstream readers.
    - Preserve-only: other persistence backends, public consumers, and active-v1 compatibility bytes until cutover.
    - Out of scope: reader-family DTO migration, publication, and active cutover.
  - Non-Goals: no silent deduplication, partial load, cross-backend fallback, or unrelated scalar/nullability rewrite.
  - Pre-flight Questions:
    - Data source: the accepted shadow-v2 canonical node/relationship manifest and isolated storage fixtures.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: query the isolated P2-B2 database through the native driver.
    - DB write flow: N/A — this read-only slice writes no database state.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: native query result fields, ordering, empty/no-data, mismatch, and orphan detection measured separately from fallback.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol impacts; implement only this production storage responsibility; after code add focused round-trip/fault tests; run the full build; exercise the real native Ladybug query boundary; compare canonical fields and closure; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: native query records equal the Graph JSON canonical manifest field-for-field with `differing_records == 0` and `orphan_refs == 0`.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real native Ladybug query boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real native Ladybug query probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2B3-IMPACT1`, `E2-P2B3-SRC1`, `E2-P2B3-BUILD1`, `E2-P2B3-TEST1`.
  - Implementation Gate: P2-B2 PASS and native-driver generation pinning is correct.
  - Acceptance:
    - Source: the touched owner contains only this storage responsibility and thin adapters; no other backend is changed.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: native query records equal the Graph JSON canonical manifest field-for-field with `differing_records == 0` and `orphan_refs == 0`.
    - Behavior test: positive round trips and duplicate/conflict/version/orphan negative cases pass with no skipped records.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2B3-IMPACT1`, `E2-P2B3-SRC1`, `E2-P2B3-BUILD1`, `E2-P2B3-TEST1`, `E2-P2B3-REVIEW1`.
    - Actual-status rows refreshed: Project canonical v2 records through native Ladybug queries.
  - Evidence Targets: source diff, full build, canonical manifest comparison, real native Ladybug query probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: project canonical v2 records through native ladybug queries. `wrong/unbound -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-B4: Project canonical v2 records through the Go fallback query path.
  - Goal: make the explicit fallback backend return the same canonical records while remaining observably distinct from native execution.
  - Scope Boundary:
    - Editable: Go/in-memory fallback query projection owner only.
    - Inspect-only: P2-A compatibility types, the preceding storage format, and downstream readers.
    - Preserve-only: other persistence backends, public consumers, and active-v1 compatibility bytes until cutover.
    - Out of scope: reader-family DTO migration, publication, and active cutover.
  - Non-Goals: no silent deduplication, partial load, cross-backend fallback, or unrelated scalar/nullability rewrite.
  - Pre-flight Questions:
    - Data source: the accepted shadow-v2 canonical node/relationship manifest and isolated storage fixtures.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read the version-qualified Graph JSON through the explicit fallback adapter.
    - DB write flow: N/A — this read-only slice writes no database state.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: supported fallback query shapes, explicit unsupported query, mismatch, ordering, and parity measured separately from native.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol impacts; implement only this production storage responsibility; after code add focused round-trip/fault tests; run the full build; exercise the real fallback query boundary; compare canonical fields and closure; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: fallback records equal the canonical manifest with zero field/orphan differences; unsupported queries fail explicitly and never masquerade as native.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real fallback query boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real fallback query probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2B4-IMPACT1`, `E2-P2B4-SRC1`, `E2-P2B4-BUILD1`, `E2-P2B4-TEST1`.
  - Implementation Gate: P2-B3 PASS and fallback selection cannot hide native compatibility errors.
  - Acceptance:
    - Source: the touched owner contains only this storage responsibility and thin adapters; no other backend is changed.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: fallback records equal the canonical manifest with zero field/orphan differences; unsupported queries fail explicitly and never masquerade as native.
    - Behavior test: positive round trips and duplicate/conflict/version/orphan negative cases pass with no skipped records.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2B4-IMPACT1`, `E2-P2B4-SRC1`, `E2-P2B4-BUILD1`, `E2-P2B4-TEST1`, `E2-P2B4-REVIEW1`.
    - Actual-status rows refreshed: Project canonical v2 records through the Go fallback query path.
  - Evidence Targets: source diff, full build, canonical manifest comparison, real fallback query probe, Supervisor, detect-changes, and commit.
  - Actual-status Update: project canonical v2 records through the go fallback query path. `wrong/unbound -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-C: Remove semantic ID parsing from CLI readers.
  - Goal: make CLI JSON/text DTO and render adapters consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only CLI JSON/text DTO and render adapters production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S3 rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read guarded Graph JSON/native/fallback results for one generation.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: CLI JSON/text outputs only.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every exact CLI consumer row plus randomized IDs, missing explicit fields, and stale generation.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only CLI JSON/text DTO and render adapters; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real packaged CLI boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real packaged CLI boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real packaged CLI probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2C-IMPACT1`, `E2-P2C-SRC1`, `E2-P2C-BUILD1`, `E2-P2C-TEST1`.
  - Implementation Gate: P2-B4 and P2-A5 PASS.
  - Acceptance:
    - Source: CLI JSON/text DTO and render adapters has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real packaged CLI output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2C-IMPACT1`, `E2-P2C-SRC1`, `E2-P2C-BUILD1`, `E2-P2C-TEST1`, `E2-P2C-REVIEW1`.
    - Actual-status rows refreshed: Remove semantic ID parsing from CLI readers.
  - Evidence Targets: parsing inventory delta, source diff, full build, real packaged CLI captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S3 opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-C1: Remove semantic ID parsing from MCP resources and tools.
  - Goal: make MCP resource/tool DTO and response adapters consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only MCP resource/tool DTO and response adapters production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S4 rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read guarded graph/query/resource results for one generation.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: MCP JSON-RPC/resource payloads only.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every exact MCP resource/tool row, JSON-RPC mismatch data, randomized IDs, and stale generation.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only MCP resource/tool DTO and response adapters; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real MCP JSON-RPC boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real MCP JSON-RPC boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real MCP JSON-RPC probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2C1-IMPACT1`, `E2-P2C1-SRC1`, `E2-P2C1-BUILD1`, `E2-P2C1-TEST1`.
  - Implementation Gate: P2-C and P2-A6 PASS.
  - Acceptance:
    - Source: MCP resource/tool DTO and response adapters has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real MCP JSON-RPC output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2C1-IMPACT1`, `E2-P2C1-SRC1`, `E2-P2C1-BUILD1`, `E2-P2C1-TEST1`, `E2-P2C1-REVIEW1`.
    - Actual-status rows refreshed: Remove semantic ID parsing from MCP resources and tools.
  - Evidence Targets: parsing inventory delta, source diff, full build, real MCP JSON-RPC captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S4 opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-C2: Make file-context projections use explicit canonical fields.
  - Goal: make file-context projection owner and transport-neutral record consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only file-context projection owner and transport-neutral record production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact file-context rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read canonical graph records for the requested file and pinned generation.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: file-detail/file-context record consumed by existing transports.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: paths/ranges/labels/references from explicit fields, missing file, stale generation, and no ID parsing.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only file-context projection owner and transport-neutral record; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real file-context projection boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real file-context projection boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real file-context projection probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2C2-IMPACT1`, `E2-P2C2-SRC1`, `E2-P2C2-BUILD1`, `E2-P2C2-TEST1`.
  - Implementation Gate: P2-C1 PASS; cache behavior remains out of scope until P2-C3.
  - Acceptance:
    - Source: file-context projection owner and transport-neutral record has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real file-context projection output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2C2-IMPACT1`, `E2-P2C2-SRC1`, `E2-P2C2-BUILD1`, `E2-P2C2-TEST1`, `E2-P2C2-REVIEW1`.
    - Actual-status rows refreshed: Make file-context projections use explicit canonical fields.
  - Evidence Targets: parsing inventory delta, source diff, full build, real file-context projection captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: file-context opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-C3: Make file-context cache records generation/config/catalog-bound.
  - Goal: make file-context cache record/key owner consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only file-context cache record/key owner production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S7 rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read and hydrate canonical file-context records only after exact key match.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: cache hit/miss/stale evidence only.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: hit, miss, stale generation, config/catalog mismatch, eviction, and orphan hydration.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only file-context cache record/key owner; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real file-context cache boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real file-context cache boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real file-context cache probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2C3-IMPACT1`, `E2-P2C3-SRC1`, `E2-P2C3-BUILD1`, `E2-P2C3-TEST1`.
  - Implementation Gate: P2-C2 and P2-A9 PASS.
  - Acceptance:
    - Source: file-context cache record/key owner has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real file-context cache output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2C3-IMPACT1`, `E2-P2C3-SRC1`, `E2-P2C3-BUILD1`, `E2-P2C3-TEST1`, `E2-P2C3-REVIEW1`.
    - Actual-status rows refreshed: Make file-context cache records generation/config/catalog-bound.
  - Evidence Targets: parsing inventory delta, source diff, full build, real file-context cache captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S7 opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-C4: Make rename use source anchors instead of parsed IDs.
  - Goal: make rename source-anchor and legacy-reference adapter consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only rename source-anchor and legacy-reference adapter production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact rename rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read a generation-qualified SymbolRef plus explicit declaration/file/range fields.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: CLI/MCP rename preview/result contracts.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: unique/ambiguous/missing legacy mapping, stale source commit, collision, and opaque random IDs.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only rename source-anchor and legacy-reference adapter; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real rename preview boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real rename preview boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real rename preview probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2C4-IMPACT1`, `E2-P2C4-SRC1`, `E2-P2C4-BUILD1`, `E2-P2C4-TEST1`.
  - Implementation Gate: P2-C3 PASS; graph-guided rename authority is preserved.
  - Acceptance:
    - Source: rename source-anchor and legacy-reference adapter has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real rename preview output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2C4-IMPACT1`, `E2-P2C4-SRC1`, `E2-P2C4-BUILD1`, `E2-P2C4-TEST1`, `E2-P2C4-REVIEW1`.
    - Actual-status rows refreshed: Make rename use source anchors instead of parsed IDs.
  - Evidence Targets: parsing inventory delta, source diff, full build, real rename preview captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: rename opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-C5: Make the shared HTTP/MCP resource cache preserve canonical records.
  - Goal: make resource-cache record/key owner consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only resource-cache record/key owner production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S8 rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read canonical resource rows only after generation/config/catalog key match.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: cache hit/miss/stale evidence only.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: both HTTP and MCP callers, cross-caller mismatch, stale generation, eviction, and orphan rows.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only resource-cache record/key owner; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real shared resource cache boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real shared resource cache boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real shared resource cache probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2C5-IMPACT1`, `E2-P2C5-SRC1`, `E2-P2C5-BUILD1`, `E2-P2C5-TEST1`.
  - Implementation Gate: P2-C4 and P2-A10 PASS.
  - Acceptance:
    - Source: resource-cache record/key owner has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real shared resource cache output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2C5-IMPACT1`, `E2-P2C5-SRC1`, `E2-P2C5-BUILD1`, `E2-P2C5-TEST1`, `E2-P2C5-REVIEW1`.
    - Actual-status rows refreshed: Make the shared HTTP/MCP resource cache preserve canonical records.
  - Evidence Targets: parsing inventory delta, source diff, full build, real shared resource cache captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S8 opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-C6: Make embedding references generation-qualified and ID-opaque.
  - Goal: make embedding projection, lookup-reference, and cache-key owner consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only embedding projection, lookup-reference, and cache-key owner production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S9 rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read canonical SymbolRefs and existing embedding rows for one generation/catalog/dimension/hash.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: embedding/search/job evidence only.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: node/generation/catalog/dimension/hash mismatch, stale reuse, orphan hydration, and randomized IDs.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only embedding projection, lookup-reference, and cache-key owner; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real embedding store/search boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real embedding store/search boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real embedding store/search probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2C6-IMPACT1`, `E2-P2C6-SRC1`, `E2-P2C6-BUILD1`, `E2-P2C6-TEST1`.
  - Implementation Gate: P2-C5 and P2-A11 PASS.
  - Acceptance:
    - Source: embedding projection, lookup-reference, and cache-key owner has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real embedding store/search output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2C6-IMPACT1`, `E2-P2C6-SRC1`, `E2-P2C6-BUILD1`, `E2-P2C6-TEST1`, `E2-P2C6-REVIEW1`.
    - Actual-status rows refreshed: Make embedding references generation-qualified and ID-opaque.
  - Evidence Targets: parsing inventory delta, source diff, full build, real embedding store/search captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S9 opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-D: Make group contracts use generation-qualified opaque references.
  - Goal: make group contract/reference adapter consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only group contract/reference adapter production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S10 group rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read one pinned member-repo generation vector and canonical contract records.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: group list/status/query/contracts outputs.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: same-HEAD new generation, vector conflict, raw-ID persistence, stale member, and deterministic source-key ordering.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only group contract/reference adapter; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real group CLI/MCP/query boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real group CLI/MCP/query boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real group CLI/MCP/query probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2D-IMPACT1`, `E2-P2D-SRC1`, `E2-P2D-BUILD1`, `E2-P2D-TEST1`.
  - Implementation Gate: P2-C6 and P2-A13 PASS.
  - Acceptance:
    - Source: group contract/reference adapter has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real group CLI/MCP/query output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2D-IMPACT1`, `E2-P2D-SRC1`, `E2-P2D-BUILD1`, `E2-P2D-TEST1`, `E2-P2D-REVIEW1`.
    - Actual-status rows refreshed: Make group contracts use generation-qualified opaque references.
  - Evidence Targets: parsing inventory delta, source diff, full build, real group CLI/MCP/query captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S10 group opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-D1: Make process projections source-anchored and ID-opaque.
  - Goal: make process membership/reference/order adapter consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only process membership/reference/order adapter production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S11 process rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read one active generation and explicit source/provenance fields.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: process CLI/MCP/HTTP projection outputs.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: randomized opaque IDs, source-anchor membership conservation, stable tie-break fields, and caps preserved.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only process membership/reference/order adapter; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real process projection boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real process projection boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real process projection probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2D1-IMPACT1`, `E2-P2D1-SRC1`, `E2-P2D1-BUILD1`, `E2-P2D1-TEST1`.
  - Implementation Gate: P2-D and P2-A14 PASS.
  - Acceptance:
    - Source: process membership/reference/order adapter has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real process projection output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2D1-IMPACT1`, `E2-P2D1-SRC1`, `E2-P2D1-BUILD1`, `E2-P2D1-TEST1`, `E2-P2D1-REVIEW1`.
    - Actual-status rows refreshed: Make process projections source-anchored and ID-opaque.
  - Evidence Targets: parsing inventory delta, source diff, full build, real process projection captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S11 process opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-D2: Make community projections source-anchored and ID-opaque.
  - Goal: make community/cluster membership/reference/order adapter consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only community/cluster membership/reference/order adapter production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S11 community rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read one active generation and explicit source/provenance fields.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: community/cluster CLI/MCP/HTTP projection outputs.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: randomized opaque IDs, source-anchor membership conservation, stable tie-break fields, and caps preserved.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only community/cluster membership/reference/order adapter; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real community/cluster projection boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real community/cluster projection boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real community/cluster projection probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2D2-IMPACT1`, `E2-P2D2-SRC1`, `E2-P2D2-BUILD1`, `E2-P2D2-TEST1`.
  - Implementation Gate: P2-D1 and P2-A15 PASS.
  - Acceptance:
    - Source: community/cluster membership/reference/order adapter has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real community/cluster projection output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2D2-IMPACT1`, `E2-P2D2-SRC1`, `E2-P2D2-BUILD1`, `E2-P2D2-TEST1`, `E2-P2D2-REVIEW1`.
    - Actual-status rows refreshed: Make community projections source-anchored and ID-opaque.
  - Evidence Targets: parsing inventory delta, source diff, full build, real community/cluster projection captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S11 community opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-E: Expose version/generation and canonical fields through HTTP.
  - Goal: make HTTP graph/query/file/panel response and stream adapters consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only HTTP graph/query/file/panel response and stream adapters production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S5 rows owned by this consumer.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read a guarded active generation and emit its negotiated manifest before any body/stream rows.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: HTTP status/error envelope, JSON response, NDJSON/SSE header and records.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: supported/old/unsupported/mixed clients, pre-body `409`, stream abort without partial rows, and canonical field parity.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only HTTP graph/query/file/panel response and stream adapters; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real HTTP server boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real HTTP server boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real HTTP server probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2E-IMPACT1`, `E2-P2E-SRC1`, `E2-P2E-BUILD1`, `E2-P2E-TEST1`.
  - Implementation Gate: P2-D2 and P2-A7 PASS.
  - Acceptance:
    - Source: HTTP graph/query/file/panel response and stream adapters has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The real HTTP server output remains behaviorally correct with randomized opaque IDs.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2E-IMPACT1`, `E2-P2E-SRC1`, `E2-P2E-BUILD1`, `E2-P2E-TEST1`, `E2-P2E-REVIEW1`.
    - Actual-status rows refreshed: Expose version/generation and canonical fields through HTTP.
  - Evidence Targets: parsing inventory delta, source diff, full build, real HTTP server captures, Supervisor, detect-changes, and commit.
  - Actual-status Update: S5 opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-E1: Negotiate and render version/generation truthfully in Web.
  - Goal: make Web handshake, parser, lifecycle, and mismatch-state adapters consume explicit canonical fields and generation-qualified references instead of parsing or persisting raw ID meaning.
  - Scope Boundary:
    - Editable: only Web handshake, parser, lifecycle, and mismatch-state adapters production adapters and one-responsibility tests.
    - Inspect-only: canonical Graph JSON/native/fallback records, P2-A guards, and sibling consumers.
    - Preserve-only: query/business semantics, unrelated consumers, and approved UI design.
    - Out of scope: other consumer families, atomic publication, and cutover.
  - Non-Goals: no labels, paths, lines, ordering, routes, export state, or provenance recovered from opaque ID text.
  - Pre-flight Questions:
    - Data source: the canonical v2 records and exact S6 rows owned by this consumer.
    - Display permission: Preserve existing graph-view permission and visibility behavior.
    - DB read flow: read only negotiated HTTP responses/streams from one pinned generation.
    - DB write flow: N/A — consumer migration does not publish graph/index state; isolated caches are covered only when named.
    - Render location: existing Web graph and truthful mismatch surfaces; no redesign.
    - UI behavior flow: Supported canonical fields render normally; stale/mismatched generation is blocked before client state mutates.
    - Docker runtime: Mandatory after full build; run the real Docker/container runtime.
    - Playwright target: Mandatory against the built Docker URL with reusable scripts and visual inspection.
    - Behavior test: all exact Web rows, supported/mismatch/stream/lifecycle states, partial-array discard, and canonical field rendering.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: Only non-UI fields explicitly marked N/A above are excluded; Docker and Playwright remain mandatory.
  - Work Steps:
    1. Record exact impacts and the current ID-parsing inventory; implement only Web handshake, parser, lifecycle, and mismatch-state adapters; after production code add opaque-random-ID and stale-generation tests; run the full build; exercise the real built Web runtime boundary; prove the parsing-site delta; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: exercise supported and mismatch states on the real built runtime.
       - DB/data flow check: request and response/cache references use explicit fields and one pinned generation; stale or orphan references fail truthfully.
       - Render location check: the existing real Docker-served surface and official QA artifacts.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser or Chrome against the real built Docker runtime at real built Web runtime; inspect the visible supported and mismatch states.
         - Playwright: use the reusable repository Playwright suite against the built Docker runtime; retain JSON, Markdown, screenshot, and trace evidence.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2E1-IMPACT1`, `E2-P2E1-SRC1`, `E2-P2E1-BUILD1`, `E2-P2E1-PLAY1`.
  - Implementation Gate: P2-E and P2-A8 PASS; Docker and reusable Playwright are available.
  - Acceptance:
    - Source: Web handshake, parser, lifecycle, and mismatch-state adapters has zero semantic ID parsing sites and no sibling consumer changes.
    - Runtime/UI: The built Web runtime renders explicit fields and truthful mismatch state; no partial/stale record appears.
    - DB/data: all durable references are version/generation-qualified; stale or orphan refs return typed errors and never hit.
    - Behavior test: randomized opaque-ID, stale-generation, missing-field, and legacy-ID negative vectors pass.
    - Cleanup/quarantine: Reusable Playwright and official JSON+MD/visual evidence only; no dev-server artifact.
    - Evidence IDs: `E2-P2E1-IMPACT1`, `E2-P2E1-SRC1`, `E2-P2E1-BUILD1`, `E2-P2E1-PLAY1`, `E2-P2E1-REVIEW1`.
    - Actual-status rows refreshed: Negotiate and render version/generation truthfully in Web.
  - Evidence Targets: parsing inventory delta, source diff, full build, real built Web runtime captures, Docker/Playwright/visual evidence, Supervisor, detect-changes, and commit.
  - Actual-status Update: S6 opaque-consumer behavior `wrong/partial -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-E2: Freeze the pre-cutover S0-S11 canonical baseline.
  - Goal: prove the already-committed storage and consumer slices agree on one canonical generation before publication/cutover work.
  - Scope Boundary:
    - Editable: Anvien evidence, benchmark, actual-status, and validation manifests only.
    - Inspect-only: all accepted P2-A through P2-E1 production slices and the frozen reader matrix.
    - Preserve-only: all production source, built artifacts, and target source/worktree.
    - Out of scope: repairs, publication changes, and active cutover.
  - Non-Goals: no code patch inside validation; any failure reopens only its responsible earlier slice.
  - Pre-flight Questions:
    - Data source: one canonical source manifest and one built, isolated v2 generation produced by accepted predecessors.
    - Display permission: Preserve the existing graph-view permission contract; this slice adds no new visibility rule.
    - DB read flow: exercise S0 Graph JSON, S1 native, S2 fallback, exact S3-S6 row unions, S7-S8 caches, S9 embeddings, S10 registries/groups, and S11 derived projections.
    - DB write flow: write only Anvien-side evidence/benchmark manifests; no target artifact and no production state mutation.
    - Render location: built Web runtime only for S6; all reports remain under Anvien.
    - UI behavior flow: Use the real built Web runtime for S6 supported/mismatch states; all other surfaces use their nearest real non-UI boundary.
    - Docker runtime: Mandatory for the S6 portion after the full build.
    - Playwright target: Mandatory for S6 against the real built Docker URL; other surfaces record N/A with their named boundary.
    - Behavior test: every exact matrix row runs; native/fallback stay separate; each S0-S11 result reports canonical field differences, orphan refs, rows passed/total, and unlisted readers.
    - Cleanup/quarantine: remove superseded validation drafts; retain only final manifests and official Web JSON+MD/visual evidence.
    - External side effects: Built runtime and repo-local isolated generation only; no target analyze in this slice.
    - N/A notes: DB write is limited to isolated validation generation/evidence; no production source edit is allowed.
  - Work Steps:
    1. Run the full build, start the real built runtime, execute every frozen matrix row and every S0-S11 canonical comparison, run reusable Playwright for S6, inspect visuals, and record separate native/fallback results.
       - UI flow check: exercise supported and mismatch Web states; confirm incompatible/partial records never render.
       - DB/data flow check: one manifest/generation feeds all surfaces; each reports `differing_records`, `orphan_refs`, and row coverage independently.
       - Render location check: the real Docker URL plus Anvien evidence/benchmark and official QA directories.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser or Chrome against the real built Docker runtime at cross-surface built-runtime acceptance; inspect the visible supported and mismatch states.
         - Playwright: use the reusable repository Playwright suite against the built Docker runtime; retain JSON, Markdown, screenshot, and trace evidence.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2E2-BUILD1`, `E2-P2E2-S0BASE1`, `E2-P2E2-S1BASE1`, `E2-P2E2-S2BASE1`, `E2-P2E2-S3BASE1`, `E2-P2E2-S4BASE1`, `E2-P2E2-S5BASE1`, `E2-P2E2-S6BASE1`, `E2-P2E2-S7BASE1`, `E2-P2E2-S8BASE1`, `E2-P2E2-S9BASE1`, `E2-P2E2-S10BASE1`, `E2-P2E2-S11BASE1`, `E2-P2E2-MATRIX1`, `E2-P2E2-PLAY1`.
  - Implementation Gate: P2-A through P2-E1 are independently accepted and committed; the matrix hash and canonical source manifest are frozen.
  - Acceptance:
    - Source: no production source changes in this validation slice; failures identify and reopen an exact earlier owner slice.
    - Runtime/UI: the real built runtime and every non-UI boundary pass independently; S6 visual states are inspected.
    - DB/data: every S0-S11 surface has `differing_records == 0`, `orphan_refs == 0`, one generation, and native/fallback results remain separate.
    - Behavior test: `rows_passed == rows_total`, `unlisted_readers == 0`, and no parent dispatcher substitutes for an exact child row.
    - Cleanup/quarantine: only final manifests and official reusable QA evidence remain; no target or debug artifact remains.
    - Evidence IDs: `E2-P2E2-BUILD1`, `E2-P2E2-S0BASE1`, `E2-P2E2-S1BASE1`, `E2-P2E2-S2BASE1`, `E2-P2E2-S3BASE1`, `E2-P2E2-S4BASE1`, `E2-P2E2-S5BASE1`, `E2-P2E2-S6BASE1`, `E2-P2E2-S7BASE1`, `E2-P2E2-S8BASE1`, `E2-P2E2-S9BASE1`, `E2-P2E2-S10BASE1`, `E2-P2E2-S11BASE1`, `E2-P2E2-MATRIX1`, `E2-P2E2-PLAY1`, `E2-P2E2-REVIEW1`, `E2-P2E2-DETECT1`, `E2-P2E2-COMMIT1`.
    - Actual-status rows refreshed: all S0-S11 pre-cutover canonical and reader-coverage rows.
  - Evidence Targets: full build, exact matrix-run manifest, twelve canonical baseline records, Docker/Playwright/visual evidence, Supervisor, detect-changes, and validation-artifact commit.
  - Actual-status Update: cross-surface pre-cutover baseline `missing -> correct`; publication remains wrong.
  - Commit Boundary: after Supervisor PASS and Anvien-side change/boundary detection, commit only the final validation ledgers/manifests; no implementation or target artifact.

- [ ] P2-F: Stage immutable repo-local generation artifacts.
  - Goal: create a complete, hash-verified generation directory without changing any active pointer.
  - Scope Boundary:
    - Editable: generation staging owner for graph, Ladybug, meta, and manifest artifacts only.
    - Inspect-only: preceding immutable generation artifacts, compatibility manifest, and downstream publication readers.
    - Preserve-only: artifact payload semantics, prior active generation, and unrelated global stores.
    - Out of scope: other publication stores, identity cutover, and semantic TS fixes.
  - Non-Goals: no best-effort multi-file publication, early live-data deletion, hidden fallback, or cross-owner write.
  - Pre-flight Questions:
    - Data source: the accepted P2-E2 generation manifest and the exact store/artifact class named by this slice.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read accepted canonical graph/Ladybug/meta inputs and current active manifest.
    - DB write flow: write immutable staged graph/Ladybug/meta/manifest files under the repo-local generation namespace, then fsync files/directories.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: failure after every artifact write/hash/fsync, incomplete stage rejection, restart discovery, and no active-state change.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol/store impacts; implement only this production publication responsibility; after code add crash/interleaving tests; run the full build; exercise the real repo-local generation staging boundary; prove prior-state queryability; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: a stage is either complete and hash-valid or discarded; the prior active generation/pointer is byte-identical and queryable.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real repo-local generation staging boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real repo-local generation staging probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2F-IMPACT1`, `E2-P2F-SRC1`, `E2-P2F-BUILD1`, `E2-P2F-TEST1`.
  - Implementation Gate: P2-E2 PASS; staging layout, hashes, fsync, and cleanup policy are approved.
  - Acceptance:
    - Source: one dedicated owner controls this artifact/store transition; adapters cannot invent or switch generations.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: a stage is either complete and hash-valid or discarded; the prior active generation/pointer is byte-identical and queryable.
    - Behavior test: success, every named failure point, restart, and concurrent conflict preserve one truthful active state.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2F-IMPACT1`, `E2-P2F-SRC1`, `E2-P2F-BUILD1`, `E2-P2F-TEST1`, `E2-P2F-REVIEW1`.
    - Actual-status rows refreshed: Stage immutable repo-local generation artifacts.
  - Evidence Targets: source/store diff, full build, fault/interleaving evidence, real repo-local generation staging probe, prior-state hashes/queryability, Supervisor, detect-changes, and commit.
  - Actual-status Update: stage immutable repo-local generation artifacts. `missing/wrong -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-F1: Publish the repo-local active generation atomically.
  - Goal: switch one repo-local active manifest only after a complete P2-F stage validates and preserve rollback.
  - Scope Boundary:
    - Editable: repo-local active-manifest CAS and rollback owner only.
    - Inspect-only: preceding immutable generation artifacts, compatibility manifest, and downstream publication readers.
    - Preserve-only: artifact payload semantics, prior active generation, and unrelated global stores.
    - Out of scope: other publication stores, identity cutover, and semantic TS fixes.
  - Non-Goals: no best-effort multi-file publication, early live-data deletion, hidden fallback, or cross-owner write.
  - Pre-flight Questions:
    - Data source: the accepted P2-E2 generation manifest and the exact store/artifact class named by this slice.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read the current active manifest and one complete staged generation.
    - DB write flow: CAS one active pointer/manifest; never rename or delete payloads as the authority.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: CAS conflict, crash before/after switch, restart, compatibility-path update, and explicit rollback.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol/store impacts; implement only this production publication responsibility; after code add crash/interleaving tests; run the full build; exercise the real repo-local active-manifest publication boundary; prove prior-state queryability; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: readers observe either the complete old or complete new generation; CAS conflict/fault preserves the old queryable generation.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real repo-local active-manifest publication boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real repo-local active-manifest publication probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2F1-IMPACT1`, `E2-P2F1-SRC1`, `E2-P2F1-BUILD1`, `E2-P2F1-TEST1`.
  - Implementation Gate: P2-F PASS and reader pinning contract is defined.
  - Acceptance:
    - Source: one dedicated owner controls this artifact/store transition; adapters cannot invent or switch generations.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: readers observe either the complete old or complete new generation; CAS conflict/fault preserves the old queryable generation.
    - Behavior test: success, every named failure point, restart, and concurrent conflict preserve one truthful active state.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2F1-IMPACT1`, `E2-P2F1-SRC1`, `E2-P2F1-BUILD1`, `E2-P2F1-TEST1`, `E2-P2F1-REVIEW1`.
    - Actual-status rows refreshed: Publish the repo-local active generation atomically.
  - Evidence Targets: source/store diff, full build, fault/interleaving evidence, real repo-local active-manifest publication probe, prior-state hashes/queryability, Supervisor, detect-changes, and commit.
  - Actual-status Update: publish the repo-local active generation atomically. `missing/wrong -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-F2: Publish cache and embedding namespaces by generation.
  - Goal: make file-context/resource caches and embeddings immutable generation-qualified publications without switching repo/global pointers.
  - Scope Boundary:
    - Editable: cache/embedding namespace publication and invalidation owners only.
    - Inspect-only: preceding immutable generation artifacts, compatibility manifest, and downstream publication readers.
    - Preserve-only: artifact payload semantics, prior active generation, and unrelated global stores.
    - Out of scope: other publication stores, identity cutover, and semantic TS fixes.
  - Non-Goals: no best-effort multi-file publication, early live-data deletion, hidden fallback, or cross-owner write.
  - Pre-flight Questions:
    - Data source: the accepted P2-E2 generation manifest and the exact store/artifact class named by this slice.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read the selected repo generation, config/catalog hashes, and existing immutable namespaces.
    - DB write flow: write new generation-qualified cache/embedding namespaces and mark them eligible only for the matching active generation.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: stale-hit, partial namespace, invalidation, restart, orphan, and publish-after-repo-CAS interleavings.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol/store impacts; implement only this production publication responsibility; after code add crash/interleaving tests; run the full build; exercise the real cache/embedding publication boundary; prove prior-state queryability; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: no cache/embedding row can cross generation/config/catalog/dimension/hash; partial namespaces are never eligible.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real cache/embedding publication boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real cache/embedding publication probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2F2-IMPACT1`, `E2-P2F2-SRC1`, `E2-P2F2-BUILD1`, `E2-P2F2-TEST1`.
  - Implementation Gate: P2-F1 PASS; P2-C3/P2-C5/P2-C6 key contracts are correct.
  - Acceptance:
    - Source: one dedicated owner controls this artifact/store transition; adapters cannot invent or switch generations.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: no cache/embedding row can cross generation/config/catalog/dimension/hash; partial namespaces are never eligible.
    - Behavior test: success, every named failure point, restart, and concurrent conflict preserve one truthful active state.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2F2-IMPACT1`, `E2-P2F2-SRC1`, `E2-P2F2-BUILD1`, `E2-P2F2-TEST1`, `E2-P2F2-REVIEW1`.
    - Actual-status rows refreshed: Publish cache and embedding namespaces by generation.
  - Evidence Targets: source/store diff, full build, fault/interleaving evidence, real cache/embedding publication probe, prior-state hashes/queryability, Supervisor, detect-changes, and commit.
  - Actual-status Update: publish cache and embedding namespaces by generation. `missing/wrong -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-F3: Publish the global repository registry atomically.
  - Goal: store immutable registry snapshots and CAS one registry epoch without coupling it to group publication.
  - Scope Boundary:
    - Editable: global repository-registry snapshot and active-epoch owner only.
    - Inspect-only: preceding immutable generation artifacts, compatibility manifest, and downstream publication readers.
    - Preserve-only: artifact payload semantics, prior active generation, and unrelated global stores.
    - Out of scope: other publication stores, identity cutover, and semantic TS fixes.
  - Non-Goals: no best-effort multi-file publication, early live-data deletion, hidden fallback, or cross-owner write.
  - Pre-flight Questions:
    - Data source: the accepted P2-E2 generation manifest and the exact store/artifact class named by this slice.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read the current registry epoch and the newly active repo generation.
    - DB write flow: write an immutable registry snapshot then CAS the registry active pointer.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: write/fsync/CAS conflict, concurrent register/delete, restart, and stale reader interleavings.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol/store impacts; implement only this production publication responsibility; after code add crash/interleaving tests; run the full build; exercise the real global repository registry publication boundary; prove prior-state queryability; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: registry readers observe one complete epoch whose repo entry names one valid active generation; conflicts retain the prior epoch.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real global repository registry publication boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real global repository registry publication probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2F3-IMPACT1`, `E2-P2F3-SRC1`, `E2-P2F3-BUILD1`, `E2-P2F3-TEST1`.
  - Implementation Gate: P2-F1 PASS and P2-A12 guard is correct.
  - Acceptance:
    - Source: one dedicated owner controls this artifact/store transition; adapters cannot invent or switch generations.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: registry readers observe one complete epoch whose repo entry names one valid active generation; conflicts retain the prior epoch.
    - Behavior test: success, every named failure point, restart, and concurrent conflict preserve one truthful active state.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2F3-IMPACT1`, `E2-P2F3-SRC1`, `E2-P2F3-BUILD1`, `E2-P2F3-TEST1`, `E2-P2F3-REVIEW1`.
    - Actual-status rows refreshed: Publish the global repository registry atomically.
  - Evidence Targets: source/store diff, full build, fault/interleaving evidence, real global repository registry publication probe, prior-state hashes/queryability, Supervisor, detect-changes, and commit.
  - Actual-status Update: publish the global repository registry atomically. `missing/wrong -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-F4: Publish group snapshots and member-generation vectors atomically.
  - Goal: pin every member repo generation and CAS one group snapshot/vector without coupling it to repo-registry writes.
  - Scope Boundary:
    - Editable: group snapshot/member-vector and group active-pointer owner only.
    - Inspect-only: preceding immutable generation artifacts, compatibility manifest, and downstream publication readers.
    - Preserve-only: artifact payload semantics, prior active generation, and unrelated global stores.
    - Out of scope: other publication stores, identity cutover, and semantic TS fixes.
  - Non-Goals: no best-effort multi-file publication, early live-data deletion, hidden fallback, or cross-owner write.
  - Pre-flight Questions:
    - Data source: the accepted P2-E2 generation manifest and the exact store/artifact class named by this slice.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read the current group epoch, group config, and pinned active generation for every member repo.
    - DB write flow: write an immutable group snapshot/vector then CAS the group active pointer.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: two-repo interleaving, member change during sync, write/fsync/CAS conflict, restart, and stale group reader.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol/store impacts; implement only this production publication responsibility; after code add crash/interleaving tests; run the full build; exercise the real group snapshot publication boundary; prove prior-state queryability; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: every group snapshot matches its committed member vector; conflict returns `INDEX_VERSION_MISMATCH` and the prior snapshot remains queryable.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real group snapshot publication boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real group snapshot publication probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2F4-IMPACT1`, `E2-P2F4-SRC1`, `E2-P2F4-BUILD1`, `E2-P2F4-TEST1`.
  - Implementation Gate: P2-F3 PASS and P2-D/P2-A13 group contracts are correct.
  - Acceptance:
    - Source: one dedicated owner controls this artifact/store transition; adapters cannot invent or switch generations.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: every group snapshot matches its committed member vector; conflict returns `INDEX_VERSION_MISMATCH` and the prior snapshot remains queryable.
    - Behavior test: success, every named failure point, restart, and concurrent conflict preserve one truthful active state.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2F4-IMPACT1`, `E2-P2F4-SRC1`, `E2-P2F4-BUILD1`, `E2-P2F4-TEST1`, `E2-P2F4-REVIEW1`.
    - Actual-status rows refreshed: Publish group snapshots and member-generation vectors atomically.
  - Evidence Targets: source/store diff, full build, fault/interleaving evidence, real group snapshot publication probe, prior-state hashes/queryability, Supervisor, detect-changes, and commit.
  - Actual-status Update: publish group snapshots and member-generation vectors atomically. `missing/wrong -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-F5: Enforce reader leases and lease-safe generation garbage collection.
  - Goal: prevent deletion of any repo, registry, group, cache, or embedding generation while a reader still pins it.
  - Scope Boundary:
    - Editable: generation lease registry and garbage-collection owner only.
    - Inspect-only: preceding immutable generation artifacts, compatibility manifest, and downstream publication readers.
    - Preserve-only: artifact payload semantics, prior active generation, and unrelated global stores.
    - Out of scope: other publication stores, identity cutover, and semantic TS fixes.
  - Non-Goals: no best-effort multi-file publication, early live-data deletion, hidden fallback, or cross-owner write.
  - Pre-flight Questions:
    - Data source: the accepted P2-E2 generation manifest and the exact store/artifact class named by this slice.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: read active pointers, retained immutable generations, and all lease references.
    - DB write flow: write lease lifecycle/retention metadata and delete only generations proven unleased and outside retention policy.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: lease acquisition/release crash, long reader, restart, GC before/after pointer switch, and cross-store retained references.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Record file/symbol/store impacts; implement only this production publication responsibility; after code add crash/interleaving tests; run the full build; exercise the real lease and garbage-collection lifecycle boundary; prove prior-state queryability; refresh ledgers, Supervisor, detect-changes, and commit.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: GC removes only unleased obsolete generations; every pinned old repo/group/cache/embedding state remains queryable until release.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real lease and garbage-collection lifecycle boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real lease and garbage-collection lifecycle probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2F5-IMPACT1`, `E2-P2F5-SRC1`, `E2-P2F5-BUILD1`, `E2-P2F5-TEST1`.
  - Implementation Gate: P2-F2/P2-F3/P2-F4 PASS; every reader/store participates in the lease inventory.
  - Acceptance:
    - Source: one dedicated owner controls this artifact/store transition; adapters cannot invent or switch generations.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: GC removes only unleased obsolete generations; every pinned old repo/group/cache/embedding state remains queryable until release.
    - Behavior test: success, every named failure point, restart, and concurrent conflict preserve one truthful active state.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2F5-IMPACT1`, `E2-P2F5-SRC1`, `E2-P2F5-BUILD1`, `E2-P2F5-TEST1`, `E2-P2F5-REVIEW1`.
    - Actual-status rows refreshed: Enforce reader leases and lease-safe generation garbage collection.
  - Evidence Targets: source/store diff, full build, fault/interleaving evidence, real lease and garbage-collection lifecycle probe, prior-state hashes/queryability, Supervisor, detect-changes, and commit.
  - Actual-status Update: enforce reader leases and lease-safe generation garbage collection. `missing/wrong -> correct`.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.

- [ ] P2-F6: Run the complete publication failure-atomicity matrix.
  - Goal: validate already-committed staging/publication/lease owners under every fault and interleaving without patching code in the acceptance slice.
  - Scope Boundary:
    - Editable: Anvien evidence, benchmark, actual-status, and fault manifests only.
    - Inspect-only: P2-F through P2-F5 source, built artifacts, stores, and retained generations.
    - Preserve-only: all production source and active data.
    - Out of scope: repairs and cutover.
  - Non-Goals: no local fix inside validation; failure reopens exactly one responsible publication slice.
  - Pre-flight Questions:
    - Data source: one old queryable generation, one complete candidate, registry/group snapshots, caches/embeddings, and the fixed fault inventory.
    - Display permission: N/A — this slice has no user-visible display or permission decision.
    - DB read flow: query the old/new repo generation, registry epoch, group vector, cache/embedding namespaces, and lease state after every injected fault.
    - DB write flow: only isolated repo-local test stores and Anvien evidence manifests; never target state.
    - Render location: N/A — this slice has no user-visible render output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction.
    - Docker runtime: N/A — no app/UI runtime is changed; validate the named non-UI boundary after the full build.
    - Playwright target: N/A — no browser surface is owned; use the named CLI/MCP/loader/store boundary.
    - Behavior test: every artifact write, fsync, CAS, restart, concurrent group/repo interleaving, lease acquisition/release, compatibility publication, rollback, and GC boundary.
    - Cleanup/quarantine: Use package-local `testdata` or isolated repo-local stores; remove superseded debug output before review.
    - External side effects: None.
    - N/A notes: UI/display/Docker/Playwright are N/A because this is a non-UI contract; full build and nearest real boundary validation remain mandatory.
  - Work Steps:
    1. Run the full build, execute the complete fault/interleaving matrix across the committed owners, query the prior state after each failure, verify hashes/pointers/leases, remove superseded failed-stage artifacts, and obtain Supervisor review.
       - UI flow check: N/A — no browser-visible or installed-app behavior is owned by this work step.
       - DB/data flow check: every failed transition retains the exact prior queryable repo/registry/group/cache/embedding state; no mixed epoch/generation is observable.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: N/A — this step has no browser-visible or installed-app boundary; exercise the real publication fault/recovery runtime boundary instead.
         - Playwright: N/A — no browser UI is owned by this step; record the full-build command and the real publication fault/recovery runtime probe.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2F6-BUILD1`, `E2-P2F6-FAULT1`, `E2-P2F6-RECOVERY1`.
  - Implementation Gate: P2-F through P2-F5 are independently accepted and committed; fault injection can target every named boundary.
  - Acceptance:
    - Source: no production source edit occurs; each failure maps to one responsible owner slice.
    - Runtime/UI: N/A — no UI is changed; the named real non-UI boundary passes.
    - DB/data: 100% of named failures and interleavings retain a queryable prior state with zero mixed generation/epoch and correct lease retention.
    - Behavior test: the matrix has no skipped boundary; restart and post-switch GC cases pass.
    - Cleanup/quarantine: Only current evidence and reusable fixtures remain; no root or target artifact exists.
    - Evidence IDs: `E2-P2F6-BUILD1`, `E2-P2F6-FAULT1`, `E2-P2F6-RECOVERY1`, `E2-P2F6-REVIEW1`, `E2-P2F6-DETECT1`, `E2-P2F6-COMMIT1`.
    - Actual-status rows refreshed: publication failure atomicity and recovery.
  - Evidence Targets: full build, complete fault manifest, hashes/pointers/query probes, lease/GC evidence, Supervisor, detect-changes, and validation-artifact commit.
  - Actual-status Update: publication invariant `partial -> correct` or reopen the exact failed predecessor.
  - Commit Boundary: after Supervisor PASS and Anvien-side change/boundary detection, commit only final fault/recovery ledgers/manifests.

- [ ] P2-G: Cut over to identity v2 and enforce legacy ambiguity.
  - Goal: activate v2 only after every reader, consumer, storage, publication, and performance gate is independently accepted.
  - Scope Boundary:
    - Editable: active emitter selection, version constants, legacy mapping/response owner, and the narrow cutover orchestrator.
    - Inspect-only: all accepted P1/P2 slices, active readers, publication stores, and rollback artifacts.
    - Preserve-only: target source, prior generation, and all later TypeScript semantic behavior.
    - Out of scope: binding/export/resolver fixes and in-place v1 rewriting.
  - Non-Goals: no ambiguous auto-redirect, dual-active read, unmeasured cutover, or old-client access to v2 bytes.
  - Pre-flight Questions:
    - Data source: commit-bound five-run v1 baseline, five-run staged-v2 candidate, complete P2-E2 baseline, and P2-F6 rollback evidence.
    - Display permission: Preserve existing graph permissions; supported v2 and legacy/mismatch states use approved surfaces only.
    - DB read flow: read one pinned v1 baseline, one staged candidate, and all active reader/publication manifests.
    - DB write flow: CAS one newly built complete v2 generation through the accepted P2-F1/P2-F3/P2-F4 owners; preserve rollback.
    - Render location: packaged CLI/MCP/HTTP/Web supported and mismatch/legacy outputs.
    - UI behavior flow: Supported v2 renders; old/unsupported clients fail before data; ambiguous legacy references show all candidates without selecting one.
    - Docker runtime: Mandatory for final public runtime validation.
    - Playwright target: Mandatory against the real built Docker runtime for Web supported/mismatch/legacy states.
    - Behavior test: five-run baseline/candidate metrics, unique/ambiguous/missing legacy refs, old binary/client rejection, deterministic active-v2 sets, rollback, and no body before negotiation.
    - Cleanup/quarantine: remove retired shadow/debug generations only after lease-safe rollback evidence; retain required benchmark and QA artifacts.
    - External side effects: one Anvien repo-local active generation plus normal registry/group publication; no target investigation artifact.
    - N/A notes: No pre-cutover metric, reader, or rollback gate is optional.
  - Work Steps:
    1. Run at least five identical commit-bound v1 baseline measurements for analyze median, Ladybug-load median, native-query p95, fallback-query p95, graph size, and peak RSS; block the slice if any metric is absent.
       - UI flow check: N/A — the baseline step is non-UI; no client sees the candidate.
       - DB/data flow check: all runs use one corpus/config/build/machine/cache policy and one reproducible query set.
       - Render location check: Anvien benchmark ledger only; no user-visible render.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser or Chrome against the real built Docker runtime at benchmark runtime; inspect the visible supported and mismatch states.
         - Playwright: use the reusable repository Playwright suite against the built Docker runtime; retain JSON, Markdown, screenshot, and trace evidence.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2G-PREBASE1`.
    2. Build and measure at least five staged, non-active v2 candidates with the identical methodology; compare every metric and refuse the active CAS on any absent or over-budget result without explicit measured Owner approval.
       - UI flow check: N/A — the candidate remains unreachable to readers.
       - DB/data flow check: candidate artifacts stay generation-qualified and v1 readers cannot open them.
       - Render location check: Anvien benchmark ledger only; no user-visible render.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser or Chrome against the real built Docker runtime at staged candidate benchmark; inspect the visible supported and mismatch states.
         - Playwright: use the reusable repository Playwright suite against the built Docker runtime; retain JSON, Markdown, screenshot, and trace evidence.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2G-CANDIDATE1`.
    3. Record fresh impacts, activate one complete v2 generation through the accepted publication owners, and expose unique/ambiguous/missing legacy mapping plus minimum-reader protocol behavior.
       - UI flow check: exercise supported v2, old-reader mismatch, and ambiguous legacy states on approved surfaces.
       - DB/data flow check: all active pointers/registries/groups select one v2 generation/vector; no v1/v2 mixture or reused v1 embedding.
       - Render location check: the existing real Docker-served surface and official QA artifacts.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser or Chrome against the real built Docker runtime at active-v2 cutover; inspect the visible supported and mismatch states.
         - Playwright: use the reusable repository Playwright suite against the built Docker runtime; retain JSON, Markdown, screenshot, and trace evidence.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2G-IMPACT1`, `E2-P2G-SRC1`, `E2-P2G-CUTOVER1`.
    4. Run the full build, five deterministic active-v2 self-analyzes, complete old-reader matrix, packaged CLI/MCP/HTTP/Web runtime, rollback drill, Supervisor review, detect-changes, and isolated commit.
       - UI flow check: run reusable Playwright on the real built Docker runtime and visually inspect supported/mismatch/legacy states.
       - DB/data flow check: canonical set, generation/epoch/vector, cache/embedding, and rollback parity are exact.
       - Render location check: the existing real Docker-served surface and official QA artifacts.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser or Chrome against the real built Docker runtime at active-v2 built runtime; inspect the visible supported and mismatch states.
         - Playwright: use the reusable repository Playwright suite against the built Docker runtime; retain JSON, Markdown, screenshot, and trace evidence.
         - Other agents: use the equivalent real-runtime control when applicable; otherwise exercise the same non-UI boundary and record why UI control is N/A.
         - Evidence: record command, runtime or boundary, observed result, visual inspection when applicable, and every exact evidence ID named below.
       - Evidence target: `E2-P2G-BUILD1`, `E2-P2G-RUNTIME1`, `E2-P2G-PLAY1`, `E2-P2G-ROLLBACK1`.
  - Implementation Gate: every ordered predecessor P2-A through P2-F6 PASS and is committed; P1-E proves the target `4/4` only in memory; P2-E2 has zero-difference S0-S11 baselines; both five-run metric sets are complete and within budget; rollback evidence is ready.
  - Acceptance:
    - Source: v2 selection and legacy mapping are explicit, single-responsibility owners; no old path can open v2 bytes.
    - Runtime/UI: supported v2 works on packaged CLI/MCP/HTTP/Web; old/unsupported clients fail with `INDEX_VERSION_MISMATCH`; ambiguous legacy references never auto-select.
    - DB/data: one complete active generation/registry/group vector is visible, deterministic sets are `5/5`, and rollback restores the prior queryable state.
    - Behavior test: baseline-before-candidate-before-CAS, old-reader no-body, legacy matrix, deterministic runs, and rollback all pass.
    - Cleanup/quarantine: only lease-safe retired shadow/debug data is removed; required benchmark/QA evidence remains and no target artifact exists.
    - Evidence IDs: `E2-P2G-PREBASE1`, `E2-P2G-CANDIDATE1`, `E2-P2G-IMPACT1`, `E2-P2G-SRC1`, `E2-P2G-CUTOVER1`, `E2-P2G-BUILD1`, `E2-P2G-RUNTIME1`, `E2-P2G-PLAY1`, `E2-P2G-ROLLBACK1`, `E2-P2G-REVIEW1`.
    - Actual-status rows refreshed: active identity, reader protocol, publication epochs, performance, and legacy mapping.
  - Evidence Targets: two five-run metric sets, cutover manifests, complete old-reader row results, deterministic signatures, built Docker/Playwright evidence, rollback, Supervisor, detect-changes, and commit.
  - Actual-status Update: active identity `wrong -> correct`; P3 remains closed until Owner explicitly opens it.
  - Commit Boundary: commit only this slice after acceptance, Supervisor PASS, ledger refresh, detect-changes, and known worktree state.


### P3: TypeScript binding-pattern extraction

- Phase Goal: preserve every legal bound identifier as a declaration/binding with correct range, path, scope, and meaning.
- Phase Boundary:
  - In scope: recursive pattern facts, TS AST traversal, declaration contexts, graph projection, six target bindings.
  - Out of scope: assignment destructuring as declaration, export/re-export semantics.
  - Dependencies: identity v2 active.
- Phase Implementation Rule: model/walker, context integration, and projection/target validation are separate commits.
- Ordered Slice List:
  - P3-A: Add recursive binding-pattern facts and walker.
  - P3-B: Integrate variable-declaration binding contexts.
  - P3-B1: Integrate parameter binding contexts.
  - P3-B2: Integrate catch binding contexts.
  - P3-B2A: Integrate for-of/for-in binding contexts.
  - P3-C: Project binding occurrences into the graph.
  - P3-C1: Project binding persistence/read adapters.
  - P3-C1A: Project binding CLI adapters.
  - P3-C1B: Project binding MCP adapters.
  - P3-C1C: Project binding file-context cache records.
  - P3-C1D: Project binding HTTP adapters.
  - P3-C1E: Project binding HTTP/MCP resource-cache records.
  - P3-C1F: Project binding Web adapters.
  - P3-C1G: Project binding embedding references.
  - P3-C1H: Project binding registry/group references.
  - P3-C1I: Project binding process/community references.
  - P3-C2: Validate bindings against the real target.

- [ ] P3-A: Add recursive binding-pattern facts and walker.
  - Goal: recursively enumerate identifier leaves, binding paths, property keys, holes, defaults, and rest without creating graph nodes yet.
  - Scope Boundary:
    - Editable: binding fact and TS binding walker owners.
    - Inspect-only: existing collector.
    - Preserve-only: exports/imports.
    - Out of scope: graph emission.
  - Non-Goals: destructuring assignment creates writes, not definitions.
  - Pre-flight Questions:
    - Data source: Tree-sitter binding-pattern AST nodes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read AST and current ScopeIR in memory.
    - DB write flow: write only in-memory BindingPatternFact/BindingFact results.
    - Render location: ScopeIR golden and evidence output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real ScopeIR binding golden boundary after the full build.
    - Behavior test: array, object, nested, default, rest, holes, aliases, computed keys, assignment-vs-declaration, and deterministic leaf order.
    - Cleanup/quarantine: independently authored package `testdata`.
    - External side effects: None.
    - N/A notes: No graph, persistence, or public adapter exists in this slice.
  - Work Steps:
    1. Implement production fact/walker after impact and responsibility extraction.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: deterministic ScopeIR leaf order.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3A-IMPACT1`, `E3-P3A-SRC1`.
    2. Add matrix tests, run full build, and validate ScopeIR golden stability.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: exactly one fact per bound leaf.
       - Render location check: golden/evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3A-BUILD1`, `E3-P3A-TEST1`.
  - Implementation Gate: identity/range types correct; imports cannot be double-counted.
  - Acceptance:
    - Source: complete binding matrix and no assignment-created definition.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: complete binding matrix and no assignment-created definition.
    - Behavior test: The behavior gate is: complete binding matrix and no assignment-created definition.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3A-IMPACT1`, `E3-P3A-SRC1`, `E3-P3A-BUILD1`, `E3-P3A-TEST1`, `E3-P3A-REVIEW1`
    - Actual-status rows refreshed: binding model `missing -> correct`; contexts remain partial.
  - Evidence Targets: fact/walker diff, full build, matrix/golden tests, Supervisor, detect-changes, commit.
  - Actual-status Update: binding model `missing -> correct`; contexts remain partial.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-B: Integrate variable-declaration binding contexts.
  - Goal: use the walker for variable declarations while preserving identifier behavior.
  - Scope Boundary:
    - Editable: narrow TS context adapters.
    - Inspect-only: reference/type inference.
    - Preserve-only: import binding and assignment writes.
    - Out of scope: graph emission.
  - Non-Goals: no fake per-leaf type inference without evidence.
  - Pre-flight Questions:
    - Data source: variable-declaration AST nodes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read variable AST and current ScopeIR.
    - DB write flow: write only in-memory variable binding facts.
    - Render location: focused ScopeIR/evidence output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real variable ScopeIR boundary after the full build.
    - Behavior test: variable patterns, exact ranges/scopes/paths, identifier regression, and no import/definition duplication.
    - Cleanup/quarantine: package `testdata`.
    - External side effects: None.
    - N/A notes: In-memory extractor only.
  - Work Steps:
    1. Implement the variable-declaration production adapter after impact.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: ScopeIR definition/binding/source ranges.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3B-IMPACT1`, `E3-P3B-SRC1`.
    2. Add focused variable-pattern tests, run full build, and compare identifier-only regression signatures.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no duplicate import/definition facts.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3B-BUILD1`, `E3-P3B-TEST1`.
  - Implementation Gate: P3-A PASS; variable adapter owns only variable declarations.
  - Acceptance:
    - Source: variable declaration patterns correct and existing identifiers unchanged.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: variable declaration patterns correct and existing identifiers unchanged.
    - Behavior test: The behavior gate is: variable declaration patterns correct and existing identifiers unchanged.
    - Cleanup/quarantine: The boundary/cleanup condition is: variable declaration patterns correct and existing identifiers unchanged.
    - Evidence IDs: `E3-P3B-IMPACT1`, `E3-P3B-SRC1`, `E3-P3B-BUILD1`, `E3-P3B-TEST1`, `E3-P3B-REVIEW1`
    - Actual-status rows refreshed: variable binding context `partial -> correct`; parameter/catch/loop contexts remain partial until their named slices.
  - Evidence Targets: variable context diff/tests/signatures, build, Supervisor, detect-changes, commit.
  - Actual-status Update: variable binding context `partial -> correct`; parameter/catch/loop contexts remain partial until their named slices.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-B1: Integrate parameter binding contexts.
  - Goal: apply the walker to function/method/arrow parameter patterns only.
  - Scope Boundary:
    - Editable: parameter adapter and one test owner.
    - Inspect-only: variable/catch/for-of.
    - Preserve-only: call/type inference.
    - Out of scope: graph emission.
  - Non-Goals: no shared context switch or broad collector rewrite.
  - Pre-flight Questions:
    - Data source: parameter AST nodes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read parameter AST and ScopeIR.
    - DB write flow: write only in-memory parameter binding facts.
    - Render location: parameter ScopeIR/evidence output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real parameter ScopeIR boundary after the full build.
    - Behavior test: function, method, arrow, default/rest/nested parameter patterns and sibling-context preservation.
    - Cleanup/quarantine: package `testdata`.
    - External side effects: None.
    - N/A notes: In-memory extractor only.
  - Work Steps:
    1. record impact; implement production adapter; add tests; run full build; validate ScopeIR counts/ranges; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: parameter binding paths and scopes exact.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3B1-IMPACT1`, `E3-P3B1-SRC1`, `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`.
  - Implementation Gate: P3-B PASS.
  - Acceptance:
    - Source: parameter patterns correct with no variable/catch regression.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: parameter patterns correct with no variable/catch regression.
    - Behavior test: The behavior gate is: parameter patterns correct with no variable/catch regression.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3B1-IMPACT1`, `E3-P3B1-SRC1`, `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, `E3-P3B1-REVIEW1`
    - Actual-status rows refreshed: parameter extraction `partial -> correct`.
  - Evidence Targets: parameter diff/tests/build/Supervisor/detect/commit.
  - Actual-status Update: parameter extraction `partial -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-B2: Integrate catch binding contexts.
  - Goal: apply the walker to catch clauses only.
  - Scope Boundary:
    - Editable: catch adapter and one test owner.
    - Inspect-only: other contexts.
    - Preserve-only: assignment writes.
    - Out of scope: graph emission.
  - Non-Goals: no assignment destructuring definitions.
  - Pre-flight Questions:
    - Data source: catch-clause AST nodes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read catch AST and ScopeIR.
    - DB write flow: write only in-memory catch binding facts.
    - Render location: catch ScopeIR/evidence output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real catch ScopeIR boundary after the full build.
    - Behavior test: identifier/nested catch patterns, exact catch scope, assignment/import preservation, and no duplicate definitions.
    - Cleanup/quarantine: package `testdata`.
    - External side effects: None.
    - N/A notes: In-memory extractor only.
  - Work Steps:
    1. record impact; implement catch adapter; add tests; run full build; validate ScopeIR counts/ranges; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: catch scope and binding paths exact.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3B2-IMPACT1`, `E3-P3B2-SRC1`, `E3-P3B2-BUILD1`, `E3-P3B2-TEST1`.
  - Implementation Gate: P3-B1 PASS.
  - Acceptance:
    - Source: catch patterns correct and assignment/import behavior preserved.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: catch patterns correct and assignment/import behavior preserved.
    - Behavior test: The behavior gate is: catch patterns correct and assignment/import behavior preserved.
    - Cleanup/quarantine: The boundary/cleanup condition is: catch patterns correct and assignment/import behavior preserved.
    - Evidence IDs: `E3-P3B2-IMPACT1`, `E3-P3B2-SRC1`, `E3-P3B2-BUILD1`, `E3-P3B2-TEST1`, `E3-P3B2-REVIEW1`
    - Actual-status rows refreshed: catch extraction `partial -> correct`.
  - Evidence Targets: catch diff/tests/build/Supervisor/detect/commit.
  - Actual-status Update: catch extraction `partial -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-B2A: Integrate for-of/for-in binding contexts.
  - Goal: apply the walker to `for-of`/`for-in` declarations only.
  - Scope Boundary:
    - Editable: loop adapter and one test owner.
    - Inspect-only: other contexts.
    - Preserve-only: assignment writes.
    - Out of scope: graph emission.
  - Non-Goals: no catch/variable context changes.
  - Pre-flight Questions:
    - Data source: `for-of` and `for-in` declaration AST nodes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read loop AST and ScopeIR.
    - DB write flow: write only in-memory loop binding facts.
    - Render location: loop ScopeIR/evidence output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real loop ScopeIR boundary after the full build.
    - Behavior test: loop declaration patterns, holes/index paths, exact loop scope, and catch/assignment/import preservation.
    - Cleanup/quarantine: package `testdata`.
    - External side effects: None.
    - N/A notes: In-memory extractor only.
  - Work Steps:
    1. record impact; implement loop adapter; add tests; run full build; validate loop scope/indexes; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: loop binding paths and scopes exact.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3B2A-IMPACT1`, `E3-P3B2A-SRC1`, `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`.
  - Implementation Gate: P3-B2 PASS.
  - Acceptance:
    - Source: loop patterns correct and catch/assignment/import behavior preserved.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: loop patterns correct and catch/assignment/import behavior preserved.
    - Behavior test: The behavior gate is: loop patterns correct and catch/assignment/import behavior preserved.
    - Cleanup/quarantine: The boundary/cleanup condition is: loop patterns correct and catch/assignment/import behavior preserved.
    - Evidence IDs: `E3-P3B2A-IMPACT1`, `E3-P3B2A-SRC1`, `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`, `E3-P3B2A-REVIEW1`
    - Actual-status rows refreshed: loop extraction `partial -> correct`.
  - Evidence Targets: loop diff/tests/build/Supervisor/detect/commit.
  - Actual-status Update: loop extraction `partial -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C: Project binding occurrences into the graph.
  - Goal: emit distinct Declaration/Symbol nodes and source-site references for binding leaves without touching persistence readers or the target.
  - Scope Boundary:
    - Editable: binding graph projection owner only.
    - Inspect-only: resolver/persistence/commands.
    - Preserve-only: exports.
    - Out of scope: target validation.
  - Non-Goals: no target-derived fixture.
  - Pre-flight Questions:
    - Data source: accepted P3 ScopeIR binding facts.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read in-memory binding/declaration/source-site facts.
    - DB write flow: write only the in-memory graph projection.
    - Render location: canonical graph evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real in-memory binding graph boundary after the full build.
    - Behavior test: Declaration/Symbol/source-site endpoint and provenance conservation, duplicate/conflict rejection, and no persistence edit.
    - Cleanup/quarantine: remove Anvien debug output.
    - External side effects: None.
    - N/A notes: Persistence readers and target remain untouched.
  - Work Steps:
    1. Implement graph projection, then add graph behavior tests and run full build.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: in-memory canonical records and source-site endpoints.
       - Render location check: graph evidence only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C-IMPACT1`, `E3-P3C-SRC1`, `E3-P3C-BUILD1`, `E3-P3C-TEST1`.
  - Implementation Gate: P3-A/B/B1/B2 PASS; persistence readers remain untouched.
  - Acceptance:
    - Source: graph contains all binding occurrences with exact endpoints/provenance and no persistence/target edits.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: graph contains all binding occurrences with exact endpoints/provenance and no persistence/target edits.
    - Behavior test: The behavior gate is: graph contains all binding occurrences with exact endpoints/provenance and no persistence/target edits.
    - Cleanup/quarantine: The boundary/cleanup condition is: graph contains all binding occurrences with exact endpoints/provenance and no persistence/target edits.
    - Evidence IDs: `E3-P3C-IMPACT1`, `E3-P3C-SRC1`, `E3-P3C-BUILD1`, `E3-P3C-TEST1`, `E3-P3C-REVIEW1`
    - Actual-status rows refreshed: graph binding projection `unbound -> partial`.
  - Evidence Targets: projection diff, build/tests, Supervisor, detect-changes, commit.
  - Actual-status Update: graph binding projection `unbound -> partial`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1: Project binding JSON/Ladybug persistence adapters.
  - Goal: carry binding Declaration/Symbol/range/path/meaning fields through Graph JSON, Ladybug native Cypher, and fallback Cypher only.
  - Scope Boundary:
    - Editable: named JSON/Ladybug persistence adapters.
    - Inspect-only: graph projection.
    - Preserve-only: unrelated DTO fields.
    - Out of scope: CLI/MCP/file-context/HTTP and target validation.
  - Non-Goals: no resolver changes.
  - Pre-flight Questions:
    - Data source: P3-C canonical graph records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read Graph JSON plus native and fallback Ladybug results.
    - DB write flow: write isolated persistence test stores only.
    - Render location: S0/S1/S2 parity manifest.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real binding persistence parity boundary after the full build.
    - Behavior test: field-level binding parity, native/fallback separation, generation mismatch, duplicate/orphan/closure cases.
    - Cleanup/quarantine: isolated repo-local stores.
    - External side effects: None.
    - N/A notes: Public adapters are excluded.
  - Work Steps:
    1. record impacts; implement only `S0` Graph JSON, `S1` Ladybug native Cypher, and `S2` fallback Cypher adapters; add field-level parity tests with native/fallback measured separately; run full build; validate canonical fields; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: exact field-level parity.
       - Render location check: `S0`/`S1`/`S2` parity manifest only; CLI/MCP/file-context are out of scope.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1-IMPACT1`, `E3-P3C1-SRC1`, `E3-P3C1-BUILD1`, `E3-P3C1-TEST1`.
  - Implementation Gate: P3-C PASS and persistence owner matrix complete.
  - Acceptance:
    - Source: `S0`, `S1`, and `S2` preserve binding fields with field-level difference `0`, with native and fallback results reported as separate rows.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S0`, `S1`, and `S2` preserve binding fields with field-level difference `0`, with native and fallback results reported as separate rows.
    - Behavior test: The behavior gate is: `S0`, `S1`, and `S2` preserve binding fields with field-level difference `0`, with native and fallback results reported as separate rows.
    - Cleanup/quarantine: The boundary/cleanup condition is: `S0`, `S1`, and `S2` preserve binding fields with field-level difference `0`, with native and fallback results reported as separate rows.
    - Evidence IDs: `E3-P3C1-IMPACT1`, `E3-P3C1-SRC1`, `E3-P3C1-BUILD1`, `E3-P3C1-TEST1`, `E3-P3C1-REVIEW1`
    - Actual-status rows refreshed: binding persistence `unbound -> correct`.
  - Evidence Targets: parity manifest, build/tests/Supervisor/detect/commit.
  - Actual-status Update: binding persistence `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1A: Project binding CLI adapters.
  - Goal: carry canonical binding fields through only the exact `S3` CLI rows in the reader matrix.
  - Scope Boundary:
    - Editable: CLI DTO/render adapters.
    - Inspect-only: `S0`–`S2`.
    - Preserve-only: MCP/HTTP/Web/caches.
    - Out of scope: target validation.
  - Non-Goals: no MCP/API/cache behavior and no graph mutation.
  - Pre-flight Questions:
    - Data source: P3-C1 canonical records and exact S3 row union.
    - Display permission: Existing CLI access only; no new permission rule.
    - DB read flow: read one guarded active generation.
    - DB write flow: N/A — this slice performs no persistent write.
    - Render location: CLI JSON/text captures.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real packaged CLI boundary after the full build.
    - Behavior test: every S3 row, canonical fields, mismatch metadata, randomized IDs, and orphan checks.
    - Cleanup/quarantine: no persistent test state.
    - External side effects: None.
    - N/A notes: CLI output is the native user-visible boundary; graphical UI is not applicable.
  - Work Steps:
    1. record file/symbol impacts; implement production CLI adapters; add CLI contract tests after code; run full build; compare canonical fields and mismatch metadata for the complete `S3` row union; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: active generation and binding fields match `S0`.
       - Render location check: CLI JSON/text captures.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1A-IMPACT1`, `E3-P3C1A-SRC1`, `E3-P3C1A-BUILD1`, `E3-P3C1A-TEST1`.
  - Implementation Gate: P3-C1 PASS and every `S3` row is assigned.
  - Acceptance:
    - Source: complete `S3` union has field difference/orphan count `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: complete `S3` union has field difference/orphan count `0`.
    - Behavior test: The behavior gate is: complete `S3` union has field difference/orphan count `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1A-REVIEW1`
    - Actual-status rows refreshed: binding CLI projection `unbound -> correct`.
  - Evidence Targets: CLI source/contracts/build/Supervisor/detect/commit.
  - Actual-status Update: binding CLI projection `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1B: Project binding MCP adapters.
  - Goal: carry canonical binding fields through only the exact `S4` MCP resource/tool rows.
  - Scope Boundary:
    - Editable: MCP DTO/resource/tool adapters.
    - Inspect-only: CLI/persistence.
    - Preserve-only: HTTP/Web/caches.
    - Out of scope: target validation.
  - Non-Goals: no cache implementation and no HTTP transport change.
  - Pre-flight Questions:
    - Data source: P3-C1 canonical records and exact S4 row union.
    - Display permission: Existing MCP tool/resource access only; no new permission rule.
    - DB read flow: read one guarded active generation.
    - DB write flow: N/A — this slice performs no persistent write.
    - Render location: MCP JSON-RPC/resource captures.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real MCP JSON-RPC boundary after the full build.
    - Behavior test: every S4 row, canonical fields, mismatch `data.code`, randomized IDs, and orphan checks.
    - Cleanup/quarantine: no persistent test state.
    - External side effects: None.
    - N/A notes: MCP output is the native boundary; graphical UI is not applicable.
  - Work Steps:
    1. record impacts; implement MCP adapters; add JSON-RPC/resource/tool contract tests; run full build; compare complete `S4` canonical fields/errors; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no field or generation drift.
       - Render location check: MCP JSON-RPC/resource captures.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1B-IMPACT1`, `E3-P3C1B-SRC1`, `E3-P3C1B-BUILD1`, `E3-P3C1B-TEST1`.
  - Implementation Gate: P3-C1A PASS and every `S4` row is assigned.
  - Acceptance:
    - Source: complete `S4` union has field difference/orphan count `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: complete `S4` union has field difference/orphan count `0`.
    - Behavior test: The behavior gate is: complete `S4` union has field difference/orphan count `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1B-REVIEW1`
    - Actual-status rows refreshed: binding MCP projection `unbound -> correct`.
  - Evidence Targets: MCP source/contracts/build/Supervisor/detect/commit.
  - Actual-status Update: binding MCP projection `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1C: Project binding file-context cache records.
  - Goal: make `S7` file-context cache keys and records generation/config/catalog-bound while preserving binding fields.
  - Scope Boundary:
    - Editable: file-context cache owner only.
    - Inspect-only: CLI/MCP/HTTP callers.
    - Preserve-only: other caches.
    - Out of scope: target validation.
  - Non-Goals: no caller DTO changes or generic cache framework.
  - Pre-flight Questions:
    - Data source: P3-C1 canonical records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read existing S7 file-context cache records.
    - DB write flow: write isolated generation/config/catalog-qualified S7 cache rows.
    - Render location: cache hit/miss/stale evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real file-context cache boundary after the full build.
    - Behavior test: hit, miss, stale generation, mismatch, eviction, canonical fields, and orphan hydration.
    - Cleanup/quarantine: invalidate isolated cache rows.
    - External side effects: None.
    - N/A notes: No caller DTO or UI change.
  - Work Steps:
    1. record impacts; implement cache record/key change; add cache behavior tests; run full build; prove `S7` difference/stale-hit counts `0`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: stale generation cannot hit.
       - Render location check: cache evidence only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1C-IMPACT1`, `E3-P3C1C-SRC1`, `E3-P3C1C-BUILD1`, `E3-P3C1C-TEST1`.
  - Implementation Gate: P3-C1B PASS and `S7` owner is exact.
  - Acceptance:
    - Source: `S7` canonical/cache-key parity is exact.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S7` canonical/cache-key parity is exact.
    - Behavior test: The behavior gate is: `S7` canonical/cache-key parity is exact.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1C-REVIEW1`
    - Actual-status rows refreshed: binding file-context cache `unbound -> correct`.
  - Evidence Targets: cache source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: binding file-context cache `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1D: Project binding HTTP adapters.
  - Goal: carry binding fields through only the exact `S5` HTTP handlers, excluding Web and caches.
  - Scope Boundary:
    - Editable: HTTP request/response/stream adapters.
    - Inspect-only: prior surfaces.
    - Preserve-only: Web/caches.
    - Out of scope: target validation.
  - Non-Goals: no visual change or cache mutation.
  - Pre-flight Questions:
    - Data source: P3-C1 canonical records and exact S5 handler rows.
    - Display permission: Existing HTTP access contract only; no visual permission change.
    - DB read flow: read one guarded active generation.
    - DB write flow: N/A — this slice performs no persistent write.
    - Render location: HTTP response/stream captures.
    - UI behavior flow: N/A — this is a non-visual handler/stream contract; Web rendering belongs to P3-C1F.
    - Docker runtime: N/A — Web runtime is owned by P3-C1F; run the full build and real HTTP handler tests.
    - Playwright target: N/A — exercise every S5 handler/stream at the HTTP boundary; P3-C1F owns browser evidence.
    - Behavior test: every S5 row, pre-body mismatch, stream generation, canonical fields, and orphan checks.
    - Cleanup/quarantine: no persistent state.
    - External side effects: Local HTTP test runtime only.
    - N/A notes: Web UI is explicitly out of scope and owned by P3-C1F.
  - Work Steps:
    1. record API/file impacts; implement HTTP adapters; add handler/stream tests; run full build; compare complete `S5` union; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: response/stream generation and fields match `S0`.
       - Render location check: HTTP captures.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1D-IMPACT1`, `E3-P3C1D-SRC1`, `E3-P3C1D-BUILD1`, `E3-P3C1D-TEST1`.
  - Implementation Gate: P3-C1C PASS and every `S5` row is assigned.
  - Acceptance:
    - Source: complete `S5` union has field difference/orphan count `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: complete `S5` union has field difference/orphan count `0`.
    - Behavior test: The behavior gate is: complete `S5` union has field difference/orphan count `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1D-REVIEW1`
    - Actual-status rows refreshed: binding HTTP projection `unbound -> correct`.
  - Evidence Targets: HTTP source/contracts/build/Supervisor/detect/commit.
  - Actual-status Update: binding HTTP projection `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1E: Project binding HTTP/MCP resource-cache records.
  - Goal: preserve canonical binding records and generation keys in the single `S8` resource-cache responsibility shared by HTTP/MCP callers.
  - Scope Boundary:
    - Editable: resource-cache owner only.
    - Inspect-only: transports.
    - Preserve-only: `S7`.
    - Out of scope: Web/target.
  - Non-Goals: no transport DTO changes.
  - Pre-flight Questions:
    - Data source: P3-C1 canonical records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read existing S8 shared resource-cache records.
    - DB write flow: write isolated generation/config/catalog-qualified S8 rows.
    - Render location: resource-cache evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real shared resource cache boundary after the full build.
    - Behavior test: both HTTP and MCP callers, hit/miss/stale/cross-caller mismatch, eviction, fields, and orphan rows.
    - Cleanup/quarantine: invalidate isolated rows.
    - External side effects: None.
    - N/A notes: Transport DTOs and UI are excluded.
  - Work Steps:
    1. record impacts; implement `S8` cache records/keys; add stale/cross-caller tests; run full build; prove difference/stale-hit counts `0`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: HTTP and MCP never share a mismatched generation.
       - Render location check: cache evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1E-IMPACT1`, `E3-P3C1E-SRC1`, `E3-P3C1E-BUILD1`, `E3-P3C1E-TEST1`.
  - Implementation Gate: P3-C1D PASS and one resource-cache owner is identified.
  - Acceptance:
    - Source: `S8` canonical/cache-key parity is exact.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S8` canonical/cache-key parity is exact.
    - Behavior test: The behavior gate is: `S8` canonical/cache-key parity is exact.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1E-REVIEW1`
    - Actual-status rows refreshed: binding resource caches `unbound -> correct`.
  - Evidence Targets: resource-cache source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: binding resource caches `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1F: Project binding Web adapters.
  - Goal: carry and render canonical binding fields plus version mismatch state through the complete exact `S6` Web row union.
  - Scope Boundary:
    - Editable: Web parsing/render adapters.
    - Inspect-only: HTTP.
    - Preserve-only: approved layout.
    - Out of scope: caches/target.
  - Non-Goals: no redesign or backend contract change.
  - Pre-flight Questions:
    - Data source: P3-C1D HTTP contract and the complete exact S6 row union.
    - Display permission: Preserve the existing approved graph UI permission/visibility contract.
    - DB read flow: read the negotiated HTTP manifest/graph/stream for one generation.
    - DB write flow: N/A — Web projection performs no persistent write.
    - Render location: real built Docker URL and official Playwright artifacts.
    - UI behavior flow: Binding fields render truthfully; supported, mismatch, stream, and lifecycle states discard stale/partial records.
    - Docker runtime: Mandatory; build and start the real Docker/container runtime after the full build.
    - Playwright target: Mandatory for every exact S6 adapter against the real built Docker URL, with visual inspection and JSON+MD evidence.
    - Behavior test: complete S6 union, supported/mismatch/stream/lifecycle states, canonical fields, partial-record discard, and explicit non-reader classifications.
    - Cleanup/quarantine: retain reusable `playwright/` plus official JSON+MD/visual evidence; remove debug-only artifacts.
    - External side effects: Local Docker/browser runtime only.
    - N/A notes: DB write is N/A; Docker and Playwright are mandatory.
  - Work Steps:
    1. record impacts; implement the exact `S6` Web row-union adapters without changing approved layout; add contract tests for every row after code; run full build including Docker; exercise supported/mismatch/stream/lifecycle states; visually inspect; compare the complete row union; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: correct binding fields render and mismatch blocks stale records.
       - DB/data flow check: Web record generation matches HTTP manifest.
       - Render location check: real Docker URL and `Reports/qa/playwright/...`.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1F-IMPACT1`, `E3-P3C1F-SRC1`, `E3-P3C1F-BUILD1`, `E3-P3C1F-PLAY1`.
  - Implementation Gate: P3-C1E PASS, Docker runtime available, and every current `S6` row is assigned without a parent/transport row substituting for an exact child.
  - Acceptance:
    - Source: the complete `S6` row union has field difference/orphan count `0`, all explicit non-reader classifications remain truthful, and supported/mismatch visual states pass.
    - Runtime/UI: The runtime/UI condition is: the complete `S6` row union has field difference/orphan count `0`, all explicit non-reader classifications remain truthful, and supported/mismatch visual states pass.
    - DB/data: The data/contract condition is: the complete `S6` row union has field difference/orphan count `0`, all explicit non-reader classifications remain truthful, and supported/mismatch visual states pass.
    - Behavior test: The behavior gate is: the complete `S6` row union has field difference/orphan count `0`, all explicit non-reader classifications remain truthful, and supported/mismatch visual states pass.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1F-REVIEW1`
    - Actual-status rows refreshed: binding Web projection `unbound -> correct`.
  - Evidence Targets: Web source/Docker/Playwright/Supervisor/detect/commit.
  - Actual-status Update: binding Web projection `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1G: Project binding embedding references.
  - Goal: make `S9` embedding rows reference generation-qualified binding Symbols without parsing IDs.
  - Scope Boundary:
    - Editable: embedding projection/reference adapter only.
    - Inspect-only: caches/graph.
    - Preserve-only: ranking semantics.
    - Out of scope: target.
  - Non-Goals: no embedding algorithm change.
  - Pre-flight Questions:
    - Data source: P3-C canonical graph records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read graph and embedding rows for one generation.
    - DB write flow: write isolated generation/catalog/dimension/hash-qualified embedding rows.
    - Render location: embedding parity/orphan evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real embedding store boundary after the full build.
    - Behavior test: node/generation/dimension/hash parity, stale reuse, orphan hydration, and randomized opaque IDs.
    - Cleanup/quarantine: remove isolated embedding rows.
    - External side effects: None.
    - N/A notes: Ranking semantics and UI are preserved.
  - Work Steps:
    1. record impacts; implement reference adapter; add behavior tests; run full build; compare `S9` rows/orphans; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no stale or orphan binding reference.
       - Render location check: embedding evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1G-IMPACT1`, `E3-P3C1G-SRC1`, `E3-P3C1G-BUILD1`, `E3-P3C1G-TEST1`.
  - Implementation Gate: P3-C1F PASS.
  - Acceptance:
    - Source: `S9` field/orphan differences are `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S9` field/orphan differences are `0`.
    - Behavior test: The behavior gate is: `S9` field/orphan differences are `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1G-REVIEW1`
    - Actual-status rows refreshed: binding embedding references `unbound -> correct`.
  - Evidence Targets: embedding source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: binding embedding references `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1H: Project binding registry/group references.
  - Goal: preserve generation-qualified binding Symbol references in only `S10` registry/group contracts.
  - Scope Boundary:
    - Editable: registry/group reference adapter.
    - Inspect-only: group algorithm.
    - Preserve-only: product contracts.
    - Out of scope: target.
  - Non-Goals: no group semantics redesign.
  - Pre-flight Questions:
    - Data source: P3-C canonical graph records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read registry/group epoch and member-vector records.
    - DB write flow: write isolated registry/group fixtures.
    - Render location: registry/group parity evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real registry/group store boundary after the full build.
    - Behavior test: epoch/vector, stale generation, orphan SymbolRef, and opaque-reference cases.
    - Cleanup/quarantine: isolated registries.
    - External side effects: None.
    - N/A notes: Product group semantics and UI are preserved.
  - Work Steps:
    1. record impacts; implement reference adapter; add group/registry tests; run full build; compare `S10`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no raw binding ID is durable across generations.
       - Render location check: registry/group evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1H-IMPACT1`, `E3-P3C1H-SRC1`, `E3-P3C1H-BUILD1`, `E3-P3C1H-TEST1`.
  - Implementation Gate: P3-C1G PASS.
  - Acceptance:
    - Source: `S10` field/orphan differences are `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S10` field/orphan differences are `0`.
    - Behavior test: The behavior gate is: `S10` field/orphan differences are `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1H-REVIEW1`
    - Actual-status rows refreshed: binding registry/group references `unbound -> correct`.
  - Evidence Targets: registry/group source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: binding registry/group references `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C1I: Project binding process/community references.
  - Goal: preserve source-anchored binding membership and deterministic ordering in only `S11` process/community projections.
  - Scope Boundary:
    - Editable: process/community reference/order adapters.
    - Inspect-only: derivation algorithms.
    - Preserve-only: product semantics/caps.
    - Out of scope: target.
  - Non-Goals: no completeness claim or process redesign.
  - Pre-flight Questions:
    - Data source: P3-C canonical graph records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read graph records feeding derived projections.
    - DB write flow: write in-memory process/community projections only.
    - Render location: process/community parity evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real derived projection boundary after the full build.
    - Behavior test: randomized opaque IDs, source-anchor membership/order conservation, caps, and orphan refs.
    - Cleanup/quarantine: no persistent state.
    - External side effects: None.
    - N/A notes: Product algorithms/caps and UI are preserved.
  - Work Steps:
    1. record impacts; implement reference/order adapter; add conservation tests; run full build; compare `S11`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: membership/order cannot depend on parsed/lexicographic IDs.
       - Render location check: derived evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C1I-IMPACT1`, `E3-P3C1I-SRC1`, `E3-P3C1I-BUILD1`, `E3-P3C1I-TEST1`.
  - Implementation Gate: P3-C1H PASS.
  - Acceptance:
    - Source: `S11` membership/order/orphan differences are `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S11` membership/order/orphan differences are `0`.
    - Behavior test: The behavior gate is: `S11` membership/order/orphan differences are `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E3-P3C1I-REVIEW1`
    - Actual-status rows refreshed: binding process/community references `unbound -> correct`.
  - Evidence Targets: derived source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: binding process/community references `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C2: Validate bindings against the real target.
  - Goal: prove the six independent target bindings and boundary safety after projection/read adapters are committed.
  - Scope Boundary:
    - Editable: Anvien evidence/benchmark/status only.
    - Inspect-only: inspect/read/analyze target in place.
    - Preserve-only: target source/worktree.
    - Out of scope: scanner.
  - Non-Goals: no target fixture or target repair.
  - Pre-flight Questions:
    - Data source: an independently authored six-binding oracle manifest.
    - Display permission: Use the existing graph permission contract only if S6 parity is applicable.
    - DB read flow: read the target active `.anvien` generation and every applicable S0-S11 row.
    - DB write flow: normal supported analyze output under target `.anvien` only.
    - Render location: all reports/manifests under Anvien; existing graph UI only when S6 applies.
    - UI behavior flow: Conditional — exercise the existing graph UI only for applicable S6 parity; otherwise compare source/ScopeIR/graph records.
    - Docker runtime: Conditional — mandatory when S6 is checked; otherwise validate the named non-UI target boundaries.
    - Playwright target: Conditional — use the real built Docker target only when S6 applies; otherwise record N/A with source/graph reason.
    - Behavior test: exact `6/6`, declaration/path/range/scope/meaning/reference fields, row parity, no new File omission, and contamination boundary.
    - Cleanup/quarantine: remove Anvien debug artifacts; never write target reports/probes/fixtures.
    - External side effects: Supported analyze may update target `.anvien` and ignored generated-guidance timestamps only; target source/worktree otherwise unchanged.
    - N/A notes: UI/Docker/Playwright apply only to an applicable S6 row.
  - Work Steps:
    1. capture target pre-state/hash/graph path; analyze in place; compare AST/source oracle -> ScopeIR -> the exact `S0`–`S11` matrix, recording an explicit reason for any non-applicable surface; capture post-state, ignored-guidance timestamp caveat, and contamination manifest; run Supervisor.
       - UI flow check: N/A unless existing graph UI is part of parity.
       - DB/data flow check: active generation only, exact `6/6` and field-level parity.
       - Render location check: all reports/manifests under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E3-P3C2-ORACLE1`, `E3-P3C2-TARGET1`, `E3-P3C2-BOUNDARY1`.
  - Implementation Gate: P3-C1 through P3-C1I PASS and oracle has six file/name/range/selection/scope/path/meaning/reference entries without copied target source.
  - Acceptance:
    - Source: `6/6` target bindings and same-name IDs are correct across each applicable `S0`–`S11` row (non-applicable rows explicitly justified), with `0` new File omissions and unchanged target source/worktree boundary.
    - Runtime/UI: Conditional — exercise the existing graph UI only for applicable S6 parity; otherwise compare source/ScopeIR/graph records.
    - DB/data: The data/contract condition is: `6/6` target bindings and same-name IDs are correct across each applicable `S0`–`S11` row (non-applicable rows explicitly justified), with `0` new File omissions and unchanged target source/worktree boundary.
    - Behavior test: The behavior gate is: `6/6` target bindings and same-name IDs are correct across each applicable `S0`–`S11` row (non-applicable rows explicitly justified), with `0` new File omissions and unchanged target source/worktree boundary.
    - Cleanup/quarantine: The boundary/cleanup condition is: `6/6` target bindings and same-name IDs are correct across each applicable `S0`–`S11` row (non-applicable rows explicitly justified), with `0` new File omissions and unchanged target source/worktree boundary.
    - Evidence IDs: `E3-P3C2-ORACLE1`, `E3-P3C2-TARGET1`, `E3-P3C2-BOUNDARY1`, `E3-P3C2-REVIEW1`, `E3-P3C2-DETECT1`, `E3-P3C2-COMMIT1`
    - Actual-status rows refreshed: binding extraction/projection `wrong -> correct`.
  - Evidence Targets: oracle/parity/boundary manifests, Supervisor, detect-changes/document closure, commit.
  - Actual-status Update: binding extraction/projection `wrong -> correct`.
  - Commit Boundary: commit Anvien-side validation artifact only after `E3-P3C2-DETECT1`; never target artifacts.

### P4: First-class TypeScript export semantics

- Phase Goal: represent module exports explicitly and consistently, separate from access visibility.
- Phase Boundary:
  - In scope: ExportFact/meaning model, extraction of all export forms, derived compatibility properties, projection parity.
  - Out of scope: traversing export tables to consumers.
  - Dependencies: P3 complete; identity v2 active.
- Phase Implementation Rule: fact contract, extraction, and projection are separate commits.
- Ordered Slice List:
  - P4-A: Add ExportFact and meaning contracts.
  - P4-B: Extract direct/default/alias/type-only export facts.
  - P4-B1: Extract star/namespace/re-export syntax facts.
  - P4-C: Project export edge/schema records.
  - P4-C1: Project export persistence/read adapters.
  - P4-C1A: Project export CLI adapters.
  - P4-C1B: Project export MCP adapters.
  - P4-C1C: Project export file-context cache records.
  - P4-C1D: Project export HTTP adapters.
  - P4-C1E: Project export HTTP/MCP resource-cache records.
  - P4-C1F: Project export Web adapters.
  - P4-C1G: Project export embedding references.
  - P4-C1H: Project export registry/group references.
  - P4-C1I: Project export process/community references.
  - P4-C2: Validate exports against the real target.

- [ ] P4-A: Add ExportFact and meaning contracts.
  - Goal: model ExportFact, meaning lanes, direct-export authority, and compatibility field values.
  - Scope Boundary:
    - Editable: export fact/meaning owner files.
    - Inspect-only: module-request/provider/resolver.
    - Preserve-only: access visibility.
    - Out of scope: AST extraction.
  - Non-Goals: no boolean export authority.
  - Pre-flight Questions:
    - Data source: the ECMAScript/TypeScript export syntax and meaning contract.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: round-trip the in-memory contract serialization.
    - DB write flow: write only in-memory ExportFact/meaning values and isolated fixtures.
    - Render location: contract/evidence output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real ExportFact serialization boundary after the full build.
    - Behavior test: all fields/statuses, one fact per binding/specifier, anonymous/default anchors, meaning masks, and access/export separation.
    - Cleanup/quarantine: synthetic package fixtures.
    - External side effects: None.
    - N/A notes: No AST extraction or public adapter is owned.
  - Work Steps:
    1. Implement production contracts/serialization after impacts.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: deterministic ScopeIR.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4A-IMPACT1`, `E4-P4A-SRC1`.
    2. Add contract tests and run full build/round trips.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no access-visibility conflation.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4A-BUILD1`, `E4-P4A-TEST1`.
  - Implementation Gate: P1-A authority includes all export decision rows and the canonical field/value matrix.
  - Acceptance:
    - Source: complete deterministic ExportFact/meaning/compatibility contract with one fact per binding/specifier and explicit anonymous/default/alias/type-only values.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: complete deterministic ExportFact/meaning/compatibility contract with one fact per binding/specifier and explicit anonymous/default/alias/type-only values.
    - Behavior test: The behavior gate is: complete deterministic ExportFact/meaning/compatibility contract with one fact per binding/specifier and explicit anonymous/default/alias/type-only values.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-BUILD1`, `E4-P4A-TEST1`, `E4-P4A-REVIEW1`
    - Actual-status rows refreshed: export contract `missing -> correct`.
  - Evidence Targets: source/build/tests/Supervisor/detect/commit.
  - Actual-status Update: export contract `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-B: Extract direct/default/alias/type-only export facts.
  - Goal: emit truthful ExportFacts for every supported TypeScript export syntax.
  - Scope Boundary:
    - Editable: TS export extraction owner and thin collector dispatch.
    - Inspect-only: definitions/imports.
    - Preserve-only: access modifiers.
    - Out of scope: export closure.
  - Non-Goals: no choice of terminal re-export symbol.
  - Pre-flight Questions:
    - Data source: TypeScript export-declaration AST nodes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read export AST and current declaration facts.
    - DB write flow: write only in-memory direct/default/alias/type-only ExportFacts.
    - Render location: extractor golden/evidence output.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real TS export extraction boundary after the full build.
    - Behavior test: named/anonymous default, `export {x as default}`, aliases, inline type, `export type`, multi-specifier, and multi-declarator forms.
    - Cleanup/quarantine: package `testdata` with synthetic fixtures.
    - External side effects: None.
    - N/A notes: No terminal re-export resolution.
  - Work Steps:
    1. Implement production extraction after impact, keeping export logic out of `definitions.go` and `imports.go`.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: one fact per exported binding/specifier, with explicit join to the declaration statement.
       - Render location check: N/A — this work step produces no user-visible render; proof is stored in the named Anvien ledger, manifest, or trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4B-IMPACT1`, `E4-P4B-SRC1`.
    2. Add focused syntax tests, run full build, and verify direct-definition regressions.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: deterministic facts and exact ranges.
       - Render location check: evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4B-BUILD1`, `E4-P4B-TEST1`.
  - Implementation Gate: P4-A PASS.
  - Acceptance:
    - Source: direct/default/alias/type-only syntax matrix complete; property keys/local names/meanings and `isExported` derivation correct.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: direct/default/alias/type-only syntax matrix complete; property keys/local names/meanings and `isExported` derivation correct.
    - Behavior test: The behavior gate is: direct/default/alias/type-only syntax matrix complete; property keys/local names/meanings and `isExported` derivation correct.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4B-IMPACT1`, `E4-P4B-SRC1`, `E4-P4B-BUILD1`, `E4-P4B-TEST1`, `E4-P4B-REVIEW1`
    - Actual-status rows refreshed: export extraction `missing -> correct`.
  - Evidence Targets: extractor diff, build/tests, Supervisor, detect-changes, commit.
  - Actual-status Update: export extraction `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-B1: Extract star/namespace/re-export syntax facts.
  - Goal: emit syntax facts for `export *`, `export * as ns`, named re-exports, `export {default as X} from`, `export type *`, and cycles without resolving terminal symbols.
  - Scope Boundary:
    - Editable: dedicated TS re-export syntax owner.
    - Inspect-only: direct export extractor.
    - Preserve-only: access visibility.
    - Out of scope: export-table traversal.
  - Non-Goals: no terminal candidate selection.
  - Pre-flight Questions:
    - Data source: TypeScript re-export AST nodes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read star/namespace/named re-export syntax.
    - DB write flow: write only in-memory star/namespace/re-export facts.
    - Render location: re-export syntax evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real TS re-export syntax boundary after the full build.
    - Behavior test: `export *`, namespace, named/default re-export, `export type *`, cycles, exact ranges, and one fact per specifier.
    - Cleanup/quarantine: package `testdata`.
    - External side effects: None.
    - N/A notes: No terminal traversal or public adapter.
  - Work Steps:
    1. record impact; implement syntax extraction; add tests; run full build; validate one fact per specifier; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: target module/name/typeOnly/meaning fields exact.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-BUILD1`, `E4-P4B1-TEST1`.
  - Implementation Gate: P4-A/B PASS.
  - Acceptance:
    - Source: all named re-export syntax facts correct and no terminal resolution is attempted.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: all named re-export syntax facts correct and no terminal resolution is attempted.
    - Behavior test: The behavior gate is: all named re-export syntax facts correct and no terminal resolution is attempted.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-BUILD1`, `E4-P4B1-TEST1`, `E4-P4B1-REVIEW1`
    - Actual-status rows refreshed: re-export syntax `missing -> correct`.
  - Evidence Targets: syntax diff/tests/build/Supervisor/detect/commit.
  - Actual-status Update: re-export syntax `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C: Project export edge/schema records.
  - Goal: expose first-class export relationships and canonical export fields without implementing persistence/read adapters or target validation.
  - Scope Boundary:
    - Editable: export edge/schema projection owner.
    - Inspect-only: persistence/read adapters.
    - Preserve-only: module dependency edges.
    - Out of scope: re-export traversal.
  - Non-Goals: reachable-through-barrel is not direct export or public package API.
  - Pre-flight Questions:
    - Data source: accepted P4 ExportFacts.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read in-memory export facts.
    - DB write flow: write only in-memory graph export edges/schema records.
    - Render location: canonical graph-schema evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real in-memory export graph boundary after the full build.
    - Behavior test: direct/export-entry/reachable/publicApi values, access visibility separation, RelationshipIDs, and no reader mutation.
    - Cleanup/quarantine: remove debug output; no target fixture.
    - External side effects: None.
    - N/A notes: Persistence/readers and target are untouched.
  - Work Steps:
    1. Implement production projection and derived compatibility semantics after impacts.
       - UI flow check: N/A; no public adapter is touched.
       - DB/data flow check: in-memory canonical export records and RelationshipIDs only.
       - Render location check: graph evidence only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C-IMPACT1`, `E4-P4C-SRC1`.
    2. Add graph-schema tests and run full build; do not analyze target in this slice.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: canonical export records and RelationshipIDs exact.
       - Render location check: evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C-BUILD1`, `E4-P4C-TEST1`.
  - Implementation Gate: P4-A/B/B1 PASS and response contract settled.
  - Acceptance:
    - Source: graph export edges/schema values correct, access visibility unchanged, and no persistence reader is silently updated.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: graph export edges/schema values correct, access visibility unchanged, and no persistence reader is silently updated.
    - Behavior test: The behavior gate is: graph export edges/schema values correct, access visibility unchanged, and no persistence reader is silently updated.
    - Cleanup/quarantine: The boundary/cleanup condition is: graph export edges/schema values correct, access visibility unchanged, and no persistence reader is silently updated.
    - Evidence IDs: `E4-P4C-IMPACT1`, `E4-P4C-SRC1`, `E4-P4C-BUILD1`, `E4-P4C-TEST1`, `E4-P4C-REVIEW1`
    - Actual-status rows refreshed: export graph schema `missing -> partial`.
  - Evidence Targets: graph schema/source/build/tests, Supervisor, detect-changes, commit.
  - Actual-status Update: export graph schema `missing -> partial`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1: Project export persistence/read adapters.
  - Goal: carry canonical export fields through Graph JSON, Ladybug native Cypher, and fallback Cypher adapters only.
  - Scope Boundary:
    - Editable: named JSON/Ladybug persistence adapters.
    - Inspect-only: graph export schema.
    - Preserve-only: unrelated DTO fields.
    - Out of scope: CLI/MCP/HTTP/Web and target validation.
  - Non-Goals: no terminal re-export resolution.
  - Pre-flight Questions:
    - Data source: P4-C canonical graph records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read Graph JSON plus native and fallback Ladybug records.
    - DB write flow: write isolated persistence test stores.
    - Render location: S0/S1/S2 export parity manifest.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real export persistence parity boundary after the full build.
    - Behavior test: field-level parity, native/fallback separation, derived `isExported` matrix, access/export separation, duplicate/orphan/version cases.
    - Cleanup/quarantine: isolated repo-local stores.
    - External side effects: None.
    - N/A notes: Public adapters are excluded.
  - Work Steps:
    1. record impacts; implement only `S0` Graph JSON, `S1` Ladybug native Cypher, and `S2` fallback Cypher adapters; add field-level parity tests with native/fallback measured separately; run full build; compare canonical record fields; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A; public adapters are split across P4-C1A through P4-C1F.
       - DB/data flow check: zero field-level drift.
       - Render location check: `S0`/`S1`/`S2` parity manifest only; public readers are out of scope.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1-IMPACT1`, `E4-P4C1-SRC1`, `E4-P4C1-BUILD1`, `E4-P4C1-TEST1`.
  - Implementation Gate: P4-C PASS and reader matrix rows assigned.
  - Acceptance:
    - Source: `S0`, `S1`, and `S2` preserve export fields and the `isExported` value matrix with field-level difference `0`, with native/fallback rows separate.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S0`, `S1`, and `S2` preserve export fields and the `isExported` value matrix with field-level difference `0`, with native/fallback rows separate.
    - Behavior test: The behavior gate is: `S0`, `S1`, and `S2` preserve export fields and the `isExported` value matrix with field-level difference `0`, with native/fallback rows separate.
    - Cleanup/quarantine: The boundary/cleanup condition is: `S0`, `S1`, and `S2` preserve export fields and the `isExported` value matrix with field-level difference `0`, with native/fallback rows separate.
    - Evidence IDs: `E4-P4C1-IMPACT1`, `E4-P4C1-SRC1`, `E4-P4C1-BUILD1`, `E4-P4C1-TEST1`, `E4-P4C1-REVIEW1`
    - Actual-status rows refreshed: export persistence `unbound -> correct`.
  - Evidence Targets: field-parity manifest, build/tests/Supervisor/detect/commit.
  - Actual-status Update: export persistence `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1A: Project export CLI adapters.
  - Goal: carry canonical export/access fields through only the exact `S3` CLI row union.
  - Scope Boundary:
    - Editable: CLI DTO/render adapters.
    - Inspect-only: `S0`–`S2`.
    - Preserve-only: MCP/HTTP/Web/caches.
    - Out of scope: target validation.
  - Non-Goals: no MCP/API/cache behavior and no graph mutation.
  - Pre-flight Questions:
    - Data source: P4-C1 records and exact S3 row union.
    - Display permission: Existing CLI access only; no new permission rule.
    - DB read flow: read one guarded active generation.
    - DB write flow: N/A — this slice performs no persistent write.
    - Render location: CLI JSON/text captures.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real packaged CLI boundary after the full build.
    - Behavior test: every S3 row, export/access/meaning fields, mismatch metadata, randomized IDs, and orphans.
    - Cleanup/quarantine: no persistent state.
    - External side effects: None.
    - N/A notes: CLI output is the native boundary; graphical UI is not applicable.
  - Work Steps:
    1. record impacts; implement CLI adapters; add contract tests; run full build; compare direct/type-only/alias/export/access fields and mismatch metadata across complete `S3`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: generation/export/access fields match `S0`.
       - Render location check: CLI captures.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1A-IMPACT1`, `E4-P4C1A-SRC1`, `E4-P4C1A-BUILD1`, `E4-P4C1A-TEST1`.
  - Implementation Gate: P4-C1 PASS and every `S3` row is assigned.
  - Acceptance:
    - Source: complete `S3` union has field difference/orphan count `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: complete `S3` union has field difference/orphan count `0`.
    - Behavior test: The behavior gate is: complete `S3` union has field difference/orphan count `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1A-REVIEW1`
    - Actual-status rows refreshed: export CLI projection `unbound -> correct`.
  - Evidence Targets: CLI source/contracts/build/Supervisor/detect/commit.
  - Actual-status Update: export CLI projection `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1B: Project export MCP adapters.
  - Goal: carry canonical export/access fields through only the exact `S4` MCP rows.
  - Scope Boundary:
    - Editable: MCP DTO/resource/tool adapters.
    - Inspect-only: CLI/persistence.
    - Preserve-only: HTTP/Web/caches.
    - Out of scope: target validation.
  - Non-Goals: no cache implementation or HTTP transport change.
  - Pre-flight Questions:
    - Data source: P4-C1 records and exact S4 row union.
    - Display permission: Existing MCP access only; no new permission rule.
    - DB read flow: read one guarded active generation.
    - DB write flow: N/A — this slice performs no persistent write.
    - Render location: MCP JSON-RPC/resource captures.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real MCP JSON-RPC boundary after the full build.
    - Behavior test: every S4 row, export/access/meaning fields, mismatch data, randomized IDs, and orphans.
    - Cleanup/quarantine: no persistent state.
    - External side effects: None.
    - N/A notes: MCP output is the native boundary; graphical UI is not applicable.
  - Work Steps:
    1. record impacts; implement MCP adapters; add JSON-RPC/resource/tool tests; run full build; compare complete `S4`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no export/access or generation drift.
       - Render location check: MCP captures.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1B-IMPACT1`, `E4-P4C1B-SRC1`, `E4-P4C1B-BUILD1`, `E4-P4C1B-TEST1`.
  - Implementation Gate: P4-C1A PASS and every `S4` row is assigned.
  - Acceptance:
    - Source: complete `S4` union has field difference/orphan count `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: complete `S4` union has field difference/orphan count `0`.
    - Behavior test: The behavior gate is: complete `S4` union has field difference/orphan count `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1B-REVIEW1`
    - Actual-status rows refreshed: export MCP projection `unbound -> correct`.
  - Evidence Targets: MCP source/contracts/build/Supervisor/detect/commit.
  - Actual-status Update: export MCP projection `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1C: Project export file-context cache records.
  - Goal: preserve export/access fields in generation/config/catalog-bound `S7` file-context cache records.
  - Scope Boundary:
    - Editable: file-context cache owner only.
    - Inspect-only: callers.
    - Preserve-only: other caches.
    - Out of scope: target.
  - Non-Goals: no caller DTO or generic cache change.
  - Pre-flight Questions:
    - Data source: P4-C1 records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read existing S7 cache rows.
    - DB write flow: write isolated generation/config/catalog-qualified S7 rows.
    - Render location: cache parity evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real file-context cache boundary after the full build.
    - Behavior test: hit/miss/stale/mismatch/eviction, export/access fields, and orphan hydration.
    - Cleanup/quarantine: invalidate isolated rows.
    - External side effects: None.
    - N/A notes: No caller DTO or UI change.
  - Work Steps:
    1. record impacts; implement cache keys/records; add tests; run full build; prove `S7` difference/stale-hit counts `0`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: stale generation cannot expose export data.
       - Render location check: cache evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1C-IMPACT1`, `E4-P4C1C-SRC1`, `E4-P4C1C-BUILD1`, `E4-P4C1C-TEST1`.
  - Implementation Gate: P4-C1B PASS.
  - Acceptance:
    - Source: `S7` field/key parity is exact.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S7` field/key parity is exact.
    - Behavior test: The behavior gate is: `S7` field/key parity is exact.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1C-REVIEW1`
    - Actual-status rows refreshed: export file-context cache `unbound -> correct`.
  - Evidence Targets: cache source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: export file-context cache `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1D: Project export HTTP adapters.
  - Goal: carry canonical export/access fields through only the exact `S5` HTTP handlers.
  - Scope Boundary:
    - Editable: HTTP response/stream adapters.
    - Inspect-only: prior surfaces.
    - Preserve-only: Web/caches.
    - Out of scope: target.
  - Non-Goals: no visual or cache behavior.
  - Pre-flight Questions:
    - Data source: P4-C1 records and exact S5 handler rows.
    - Display permission: Existing HTTP contract only; no visual permission change.
    - DB read flow: read one guarded active generation.
    - DB write flow: N/A — this slice performs no persistent write.
    - Render location: HTTP response/stream captures.
    - UI behavior flow: N/A — non-visual HTTP contract; Web belongs to P4-C1F.
    - Docker runtime: N/A — Web runtime belongs to P4-C1F; run full build and real HTTP tests.
    - Playwright target: N/A — exercise every S5 handler/stream; P4-C1F owns browser evidence.
    - Behavior test: every S5 row, pre-body mismatch, stream generation, export/access fields, and orphans.
    - Cleanup/quarantine: no persistent state.
    - External side effects: Local HTTP test runtime only.
    - N/A notes: Web UI is explicitly excluded and owned by P4-C1F.
  - Work Steps:
    1. record API/file impacts; implement HTTP adapters; add handler/stream tests; run full build; compare complete `S5`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: response/stream generation/export fields match `S0`.
       - Render location check: HTTP captures.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1D-IMPACT1`, `E4-P4C1D-SRC1`, `E4-P4C1D-BUILD1`, `E4-P4C1D-TEST1`.
  - Implementation Gate: P4-C1C PASS and every `S5` row is assigned.
  - Acceptance:
    - Source: complete `S5` union has field difference/orphan count `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: complete `S5` union has field difference/orphan count `0`.
    - Behavior test: The behavior gate is: complete `S5` union has field difference/orphan count `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1D-REVIEW1`
    - Actual-status rows refreshed: export HTTP projection `unbound -> correct`.
  - Evidence Targets: HTTP source/contracts/build/Supervisor/detect/commit.
  - Actual-status Update: export HTTP projection `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1E: Project export HTTP/MCP resource-cache records.
  - Goal: preserve export/access fields and generation keys in the single `S8` resource-cache responsibility shared by HTTP/MCP.
  - Scope Boundary:
    - Editable: resource-cache owner only.
    - Inspect-only: transports.
    - Preserve-only: `S7`.
    - Out of scope: Web/target.
  - Non-Goals: no transport DTO changes.
  - Pre-flight Questions:
    - Data source: P4-C1 records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read existing S8 shared resource-cache records.
    - DB write flow: write isolated generation/config/catalog-qualified S8 rows.
    - Render location: resource-cache evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real shared resource cache boundary after the full build.
    - Behavior test: both callers, stale/cross-caller mismatch, export/access fields, eviction, and orphans.
    - Cleanup/quarantine: invalidate isolated rows.
    - External side effects: None.
    - N/A notes: Transport DTOs and UI are excluded.
  - Work Steps:
    1. record impacts; implement `S8` records/keys; add stale/cross-caller tests; run full build; prove difference/stale-hit counts `0`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: HTTP and MCP cannot share a mismatched generation.
       - Render location check: cache evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1E-IMPACT1`, `E4-P4C1E-SRC1`, `E4-P4C1E-BUILD1`, `E4-P4C1E-TEST1`.
  - Implementation Gate: P4-C1D PASS.
  - Acceptance:
    - Source: `S8` field/key parity is exact.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S8` field/key parity is exact.
    - Behavior test: The behavior gate is: `S8` field/key parity is exact.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1E-REVIEW1`
    - Actual-status rows refreshed: export resource caches `unbound -> correct`.
  - Evidence Targets: resource-cache source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: export resource caches `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1F: Project export Web adapters.
  - Goal: carry and render canonical export/access fields plus mismatch state through the complete exact `S6` Web row union.
  - Scope Boundary:
    - Editable: Web parsing/render adapters.
    - Inspect-only: HTTP.
    - Preserve-only: approved layout.
    - Out of scope: caches/target.
  - Non-Goals: no redesign or backend change.
  - Pre-flight Questions:
    - Data source: P4-C1D HTTP contract and complete exact S6 row union.
    - Display permission: Preserve existing approved graph UI permission/visibility.
    - DB read flow: read negotiated HTTP manifest/export records for one generation.
    - DB write flow: N/A — Web projection performs no persistent write.
    - Render location: real built Docker URL and official Playwright artifacts.
    - UI behavior flow: Export/access/type-only fields render truthfully; mismatch/stream/lifecycle states block stale/partial records.
    - Docker runtime: Mandatory; build and start the real Docker/container runtime.
    - Playwright target: Mandatory for complete S6 union against the real built Docker URL, with visual inspection and JSON+MD evidence.
    - Behavior test: complete S6 union, canonical export/access fields, supported/mismatch/stream/lifecycle states, and explicit non-readers.
    - Cleanup/quarantine: retain reusable Playwright and official JSON+MD/visual evidence; remove debug artifacts.
    - External side effects: Local Docker/browser runtime only.
    - N/A notes: DB write is N/A; Docker and Playwright are mandatory.
  - Work Steps:
    1. record impacts; implement the exact `S6` Web row-union adapters without redesign; add tests for every row after code; run full build including Docker; exercise supported/mismatch/stream/lifecycle states, compare the complete row union, and inspect visuals; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: direct/type-only/export/access fields render truthfully and mismatch blocks stale records.
       - DB/data flow check: Web generation matches HTTP manifest.
       - Render location check: real Docker URL and official Playwright artifacts.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1F-IMPACT1`, `E4-P4C1F-SRC1`, `E4-P4C1F-BUILD1`, `E4-P4C1F-PLAY1`.
  - Implementation Gate: P4-C1E PASS, Docker available, and every current `S6` row is assigned without a parent/transport row substituting for an exact child.
  - Acceptance:
    - Source: the complete `S6` row union has field/orphan difference `0`, explicit non-reader classifications remain truthful, and visual states pass.
    - Runtime/UI: The runtime/UI condition is: the complete `S6` row union has field/orphan difference `0`, explicit non-reader classifications remain truthful, and visual states pass.
    - DB/data: The data/contract condition is: the complete `S6` row union has field/orphan difference `0`, explicit non-reader classifications remain truthful, and visual states pass.
    - Behavior test: The behavior gate is: the complete `S6` row union has field/orphan difference `0`, explicit non-reader classifications remain truthful, and visual states pass.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1F-REVIEW1`
    - Actual-status rows refreshed: export Web projection `unbound -> correct`.
  - Evidence Targets: Web source/Docker/Playwright/Supervisor/detect/commit.
  - Actual-status Update: export Web projection `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1G: Project export embedding references.
  - Goal: make `S9` embedding rows reference generation-qualified exported Symbols without deriving export semantics from IDs.
  - Scope Boundary:
    - Editable: embedding projection/reference adapter only.
    - Inspect-only: graph/caches.
    - Preserve-only: ranking semantics.
    - Out of scope: target.
  - Non-Goals: no embedding algorithm change.
  - Pre-flight Questions:
    - Data source: P4-C canonical graph records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read graph and embedding records.
    - DB write flow: write isolated generation-qualified embedding rows.
    - Render location: embedding export-parity evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real embedding store boundary after the full build.
    - Behavior test: node/export/generation/dimension/hash parity, stale reuse, orphans, and opaque IDs.
    - Cleanup/quarantine: remove isolated rows.
    - External side effects: None.
    - N/A notes: Ranking semantics and UI are preserved.
  - Work Steps:
    1. record impacts; implement adapter; add tests; run full build; compare `S9`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no stale/orphan export reference.
       - Render location check: embedding evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1G-IMPACT1`, `E4-P4C1G-SRC1`, `E4-P4C1G-BUILD1`, `E4-P4C1G-TEST1`.
  - Implementation Gate: P4-C1F PASS.
  - Acceptance:
    - Source: `S9` field/orphan differences are `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S9` field/orphan differences are `0`.
    - Behavior test: The behavior gate is: `S9` field/orphan differences are `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1G-REVIEW1`
    - Actual-status rows refreshed: export embedding references `unbound -> correct`.
  - Evidence Targets: embedding source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: export embedding references `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1H: Project export registry/group references.
  - Goal: preserve generation-qualified export/public-API references in only `S10` registry/group contracts.
  - Scope Boundary:
    - Editable: registry/group reference adapter.
    - Inspect-only: group algorithms.
    - Preserve-only: product contract policy.
    - Out of scope: target.
  - Non-Goals: no group/public-API policy redesign.
  - Pre-flight Questions:
    - Data source: P4-C canonical graph records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read registry/group contracts and epochs/vectors.
    - DB write flow: write isolated registry/group rows.
    - Render location: registry/group export-parity evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real registry/group store boundary after the full build.
    - Behavior test: epoch/vector/stale/orphan cases and separation of direct/reachable/publicApi fields.
    - Cleanup/quarantine: isolated registries.
    - External side effects: None.
    - N/A notes: Product/public-API policy and UI are preserved.
  - Work Steps:
    1. record impacts; implement adapter; add group/registry tests; run full build; compare `S10`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: direct/reachable/publicApi fields remain distinct and generation-bound.
       - Render location check: registry/group evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1H-IMPACT1`, `E4-P4C1H-SRC1`, `E4-P4C1H-BUILD1`, `E4-P4C1H-TEST1`.
  - Implementation Gate: P4-C1G PASS.
  - Acceptance:
    - Source: `S10` field/orphan differences are `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S10` field/orphan differences are `0`.
    - Behavior test: The behavior gate is: `S10` field/orphan differences are `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1H-REVIEW1`
    - Actual-status rows refreshed: export registry/group references `unbound -> correct`.
  - Evidence Targets: registry/group source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: export registry/group references `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C1I: Project export process/community references.
  - Goal: preserve source-anchored export references and deterministic ordering in only `S11` process/community projections.
  - Scope Boundary:
    - Editable: process/community reference/order adapters.
    - Inspect-only: algorithms.
    - Preserve-only: product semantics/caps.
    - Out of scope: target.
  - Non-Goals: no completeness or public-API inference claim.
  - Pre-flight Questions:
    - Data source: P4-C canonical graph records.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read graph records feeding derived projections.
    - DB write flow: write in-memory derived projections only.
    - Render location: process/community export-parity evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real derived projection boundary after the full build.
    - Behavior test: randomized opaque IDs, membership/order conservation, export-field non-inference, and orphans.
    - Cleanup/quarantine: no persistent state.
    - External side effects: None.
    - N/A notes: No completeness or public-API inference; UI is preserved.
  - Work Steps:
    1. record impacts; implement adapter; add conservation tests; run full build; compare `S11`; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: membership/order cannot infer export from ID text.
       - Render location check: derived evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C1I-IMPACT1`, `E4-P4C1I-SRC1`, `E4-P4C1I-BUILD1`, `E4-P4C1I-TEST1`.
  - Implementation Gate: P4-C1H PASS.
  - Acceptance:
    - Source: `S11` membership/order/orphan differences are `0`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `S11` membership/order/orphan differences are `0`.
    - Behavior test: The behavior gate is: `S11` membership/order/orphan differences are `0`.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E4-P4C1I-REVIEW1`
    - Actual-status rows refreshed: export process/community references `unbound -> correct`.
  - Evidence Targets: derived source/tests/build/Supervisor/detect/commit.
  - Actual-status Update: export process/community references `unbound -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P4-C2: Validate exports against the real target.
  - Goal: compare an independent phase-local 21-entry source/TS-oracle manifest and negative unexported controls against the target graph.
  - Scope Boundary:
    - Editable: Anvien oracle/evidence/benchmark/status only.
    - Inspect-only: inspect/read/analyze target in place.
    - Preserve-only: target source/worktree.
    - Out of scope: scanner and resolver.
  - Non-Goals: no copied target source and no self-referential graph-only oracle.
  - Pre-flight Questions:
    - Data source: an independent 21-entry export oracle plus negative controls.
    - Display permission: Use existing graph permissions only if S6 parity is applicable.
    - DB read flow: read target active `.anvien` and every applicable S0-S11 row.
    - DB write flow: normal supported target analyze output under `.anvien` only.
    - Render location: reports/manifests under Anvien; existing graph UI only when S6 applies.
    - UI behavior flow: Conditional — exercise existing graph UI only for applicable S6 parity.
    - Docker runtime: Conditional — mandatory when S6 is checked; otherwise validate non-UI target boundaries.
    - Playwright target: Conditional — real built Docker target only when S6 applies; otherwise record N/A with source/graph reason.
    - Behavior test: `21/21`, negative unexported controls, exact export/access/meaning fields, every applicable row, and contamination boundary.
    - Cleanup/quarantine: remove Anvien debug artifacts; never target reports/probes/fixtures.
    - External side effects: Supported analyze may update target `.anvien` and ignored guidance timestamps only.
    - N/A notes: UI/Docker/Playwright apply only to an applicable S6 row.
  - Work Steps:
    1. capture target pre-state/hash/graph path; run independent source/TS oracle; analyze in place; compare ScopeIR -> graph -> S0-S11 surface matrix; capture post-state and contamination manifest; run Supervisor.
       - UI flow check: N/A unless existing graph UI is part of parity.
       - DB/data flow check: `21/21` and exact field/value matrix across all named surfaces.
       - Render location check: reports/manifests under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`.
  - Implementation Gate: P4-C1 through P4-C1I PASS; phase-local oracle is independent and contains no copied target source.
  - Acceptance:
    - Source: `21/21` direct exports and negative controls correct across exact S0-S11 surfaces; access visibility separate; target boundary unchanged.
    - Runtime/UI: Conditional — exercise existing graph UI only for applicable S6 parity.
    - DB/data: The data/contract condition is: `21/21` direct exports and negative controls correct across exact S0-S11 surfaces; access visibility separate; target boundary unchanged.
    - Behavior test: The behavior gate is: `21/21` direct exports and negative controls correct across exact S0-S11 surfaces; access visibility separate; target boundary unchanged.
    - Cleanup/quarantine: The boundary/cleanup condition is: `21/21` direct exports and negative controls correct across exact S0-S11 surfaces; access visibility separate; target boundary unchanged.
    - Evidence IDs: `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW1`, `E4-P4C2-DETECT1`, `E4-P4C2-COMMIT1`
    - Actual-status rows refreshed: export metadata/projection `wrong -> correct`.
  - Evidence Targets: oracle/parity/boundary manifests, Supervisor, detect-changes/document closure, commit.
  - Actual-status Update: export metadata/projection `wrong -> correct`.
  - Commit Boundary: commit Anvien-side validation artifact only after `E4-P4C2-DETECT1`; never target artifacts.

### P5: Module export tables and barrel/re-export resolution

- Phase Goal: resolve repository-backed imports through module export surfaces to an `InternalSymbolRef`, `ModuleNamespaceRef`, or `GapRef` with complete proof chains, and hand an immutable `ExternalResolutionCandidate` (`ModuleRef`, exported name, meaning, and proof) to P6 when declaration-universe work is required. P5 never creates `ExternalSymbolRef`.
- Phase Boundary:
  - In scope: project module settings, module requests, export tables, re-export algorithm, emission, target barrel calls.
  - Out of scope: ambient/external declaration parsing and materialization; only the immutable deferred `ExternalResolutionCandidate` handoff is in scope.
  - Dependencies: P4 complete.
- Phase Implementation Rule: path/config, table, traversal, and graph emission are separate commits.

#### P5 semantic vector manifest (authoritative for P5-C)

P5-C must check each named vector below against the expected tagged repository `ResolutionTarget` or deferred `ExternalResolutionCandidate`, status/provisional stage, and proof shape. The vector manifest is phase-local and independently authored; a broad “full matrix” statement is not sufficient.

| Vector | Expected target/status | Required proof or negative assertion |
|--------|------------------------|--------------------------------------|
| named default class/function/interface/expression | `InternalSymbolRef`, `resolved_local` | direct declaration anchor; anonymous expression uses expression anchor, never first definition in file |
| `export {x as default}` | `InternalSymbolRef`, `resolved_local` | alias hop to local `x` and default exported name |
| `export {default as X} from "m"` | terminal `ResolutionTarget` from `m` | default-name hop is retained; no physical barrel-name guess |
| inline type specifier / `export type {X}` / `export type *` used as value | `GapRef`, `meaning_mismatch` | type lane is preserved; no value fallback |
| `export * from "m"` | table adjacency only; terminal decided by traversal | `default` is excluded from star adjacency |
| `export * as ns from "m"` then member use | `ModuleNamespaceRef` then member target | namespace is synthetic and module/provenance-bound, not a physical `ns` definition |
| `import type {X} from "m"` | type-lane `InternalSymbolRef` for the in-repo fixture | value use returns `meaning_mismatch`; no value/global fallback |
| `import {type X} from "m"` | type-lane `InternalSymbolRef` for the in-repo fixture | inline type modifier survives every alias hop |
| `import * as ns from "m"` then member use | `ModuleNamespaceRef` then the selected member target | namespace import is module-bound; missing member returns `member_not_found` after P6 finalization |
| class/enum dual lane | one `InternalSymbolRef` with type+value mask | type/value meaning is not split or collapsed by consumer request |
| declaration overloads/merging with provider evidence | one logical Symbol with multiple DeclarationRefs | unverified merge remains multiple Symbols |
| explicit export vs star export | explicit terminal wins | distinct terminals are `ambiguous_export`; same terminal paths deduplicate |
| cycle with another terminal branch | terminal branch resolves | cyclic branch contributes no candidate; proof records cycle |
| pure re-export cycle | `GapRef`, `cycle_no_terminal` | no silent truncation or first-candidate choice |
| missing export/module/invalid config | `GapRef` with exact status | module failure cannot be overwritten by export/member guess |
- Ordered Slice List:
  - P5-A: Build hash-bound TypeScript project/module request inputs.
  - P5-B: Build deterministic per-module export tables.
  - P5-C: Resolve re-exports, cycles, ambiguity, aliases, and meanings.
  - P5-D: Emit terminal edges/proofs and validate barrel consumers.

- [ ] P5-A: Build hash-bound TypeScript project/module request inputs.
  - Goal: model project configuration and resolve a source specifier, path mapping, and package `exports` conditions to an immutable `ModuleRef` plus package metadata, separately from exported symbols and declaration files.
  - Scope Boundary:
    - Editable: TS project profile/module-reference/conditions owners.
    - Inspect-only: current import resolver.
    - Preserve-only: non-TS strategies.
    - Out of scope: declaration entrypoint lookup and export lookup.
  - Non-Goals: module-found does not mean symbol-found.
  - Pre-flight Questions:
    - Data source: tsconfig/package metadata and accepted P4 module-request facts.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read repository config/package metadata and authorized package `exports` inputs.
    - DB write flow: write only in-memory hash-bound `ModuleRef` and immutable package metadata.
    - Render location: structured module-resolution trace and evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real module-reference resolver boundary after the full build.
    - Behavior test: module/moduleResolution/target/lib/baseUrl/paths, package `exports` conditions, canonical roots, invalid config, and path-security vectors.
    - Cleanup/quarantine: independently authored package testdata.
    - External side effects: Read-only local metadata; no network, Node, package scripts, install, or target borrowing.
    - N/A notes: No UI, live DB, or declaration-entrypoint lookup.
  - Work Steps:
    1. Implement production profile/ModuleRef resolution with config hash, package condition result, canonical root, and structured module outcomes.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no absolute repo root persisted in ScopeIR/graph.
       - Render location check: diagnostics/trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E5-P5A-IMPACT1`, `E5-P5A-SRC1`.
    2. Add tests, run full build, compare pinned TypeScript oracle only in tests.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: exact module targets and failure codes.
       - Render location check: evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E5-P5A-BUILD1`, `E5-P5A-ORACLE1`.
  - Implementation Gate: P4 module request facts available; no runtime Node/network.
  - Acceptance:
    - Source: `ModuleRef` outcomes are deterministic and config-bound; package `exports` condition selection has exactly one owner; declaration `types`/`typesVersions` selection is explicitly delegated to P6; all non-TS regression passes.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: `ModuleRef` outcomes are deterministic and config-bound; package `exports` condition selection has exactly one owner; declaration `types`/`typesVersions` selection is explicitly delegated to P6; all non-TS regression passes.
    - Behavior test: The behavior gate is: `ModuleRef` outcomes are deterministic and config-bound; package `exports` condition selection has exactly one owner; declaration `types`/`typesVersions` selection is explicitly delegated to P6; all non-TS regression passes.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E5-P5A-IMPACT1`, `E5-P5A-SRC1`, `E5-P5A-BUILD1`, `E5-P5A-ORACLE1`, `E5-P5A-REVIEW1`
    - Actual-status rows refreshed: module resolution input `partial -> correct`.
  - Evidence Targets: config/path matrix, oracle, build, Supervisor, detect, commit.
  - Actual-status Update: module resolution input `partial -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P5-B: Build deterministic per-module export tables.
  - Goal: index explicit/default/alias/star/namespace exports without resolving consumer use sites.
  - Scope Boundary:
    - Editable: module export table owner.
    - Inspect-only: resolver/emitter.
    - Preserve-only: import dependency edges.
    - Out of scope: traversal to terminal `ResolutionTarget`.
  - Non-Goals: no wildcard exposure of unexported physical definitions.
  - Pre-flight Questions:
    - Data source: accepted ExportFacts plus immutable ModuleRef outcomes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read in-memory export facts and module adjacency inputs.
    - DB write flow: write only in-memory deterministic per-module export tables.
    - Render location: table signature/topology benchmark evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real module export table boundary after the full build.
    - Behavior test: explicit/default/direct/alias entries, default exclusion, star adjacency, namespace entries, deterministic order, and raw topology ceilings.
    - Cleanup/quarantine: synthetic module graph testdata.
    - External side effects: None.
    - N/A notes: No terminal resolution, UI, or persistent store.
  - Work Steps:
    1. Implement production explicit table and star adjacency with deterministic ordering/budget status.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: table is in-memory/hash-bound.
       - Render location check: debug trace only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E5-P5B-IMPACT1`, `E5-P5B-SRC1`.
    2. Add table tests, run full build, and measure only raw star-edge/fan-out/depth topology against the numeric ceilings; do not resolve terminal candidates here.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: no silent truncation.
       - Render location check: benchmark ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E5-P5B-BUILD1`, `E5-P5B-TEST1`, `E5-P5B-BENCH1`.
  - Implementation Gate: P5-A/P4 complete.
  - Acceptance:
    - Source: tables deterministic, only exports visible, raw topology limits explicit, and no terminal ambiguity/cycle decision is made here.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: tables deterministic, only exports visible, raw topology limits explicit, and no terminal ambiguity/cycle decision is made here.
    - Behavior test: The behavior gate is: tables deterministic, only exports visible, raw topology limits explicit, and no terminal ambiguity/cycle decision is made here.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E5-P5B-IMPACT1`, `E5-P5B-SRC1`, `E5-P5B-BUILD1`, `E5-P5B-TEST1`, `E5-P5B-BENCH1`, `E5-P5B-REVIEW1`
    - Actual-status rows refreshed: export table `missing -> correct`.
  - Evidence Targets: table signatures, benchmarks, build/tests, Supervisor, detect, commit.
  - Actual-status Update: export table `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P5-C: Resolve re-exports, cycles, ambiguity, aliases, and meanings.
  - Goal: return one repository-backed `InternalSymbolRef`/`ModuleNamespaceRef`, a structured `GapRef`, or an immutable `ExternalResolutionCandidate` for P6, with a complete hop proof for each export lookup.
  - Scope Boundary:
    - Editable: re-export traversal owner and outcome integration.
    - Inspect-only: call/type resolvers.
    - Preserve-only: graph emission.
    - Out of scope: ambient lookup.
  - Non-Goals: no first-candidate selection or global fallback.
  - Pre-flight Questions:
    - Data source: P5-B tables and the authoritative semantic vector manifest.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read in-memory table/adjacency graph and immutable external candidates.
    - DB write flow: write only memoized resolution sets and structured proof results in memory.
    - Render location: structured resolution trace plus evidence/benchmark.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real re-export resolver boundary after the full build.
    - Behavior test: all named vectors, aliases, meanings, namespace/member use, same-terminal dedupe, ambiguity, cycles, merge/overload, in-budget exactness, and explicit over-limit status.
    - Cleanup/quarantine: synthetic resolver fixtures and bounded adversarial graphs.
    - External side effects: None.
    - N/A notes: No UI, persistent DB, external materialization, or global fallback.
  - Work Steps:
    1. Implement memoized DFS or SCC-safe production resolution and proof preservation.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: exact terminal identity and hop list.
       - Render location check: structured resolution trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E5-P5C-IMPACT1`, `E5-P5C-SRC1`.
    2. Add behavior/oracle tests from the authoritative vector manifest, run full build, and measure hop/fan-out/candidate/SCC budgets against numeric ceilings; in-budget vectors must resolve exactly and only explicitly over-limit adversarial vectors may return `budget_exceeded`.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: ambiguous/cycle statuses deterministic.
       - Render location check: evidence/benchmark.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E5-P5C-BUILD1`, `E5-P5C-TEST1`, `E5-P5C-VECTOR1`, `E5-P5C-BENCH1`.
  - Implementation Gate: table contract correct; meaning lanes available.
  - Acceptance:
    - Source: full vector/status matrix passes with no silent fallback/truncation; cycle resolution-set rules and namespace/merge targets are exact; external declarations remain immutable candidates for P6; in-budget correctness is `100%`.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: full vector/status matrix passes with no silent fallback/truncation; cycle resolution-set rules and namespace/merge targets are exact; external declarations remain immutable candidates for P6; in-budget correctness is `100%`.
    - Behavior test: The behavior gate is: full vector/status matrix passes with no silent fallback/truncation; cycle resolution-set rules and namespace/merge targets are exact; external declarations remain immutable candidates for P6; in-budget correctness is `100%`.
    - Cleanup/quarantine: The boundary/cleanup condition is: full vector/status matrix passes with no silent fallback/truncation; cycle resolution-set rules and namespace/merge targets are exact; external declarations remain immutable candidates for P6; in-budget correctness is `100%`.
    - Evidence IDs: `E5-P5C-IMPACT1`, `E5-P5C-SRC1`, `E5-P5C-BUILD1`, `E5-P5C-TEST1`, `E5-P5C-VECTOR1`, `E5-P5C-BENCH1`, `E5-P5C-REVIEW1`
    - Actual-status rows refreshed: re-export lookup `wrong -> correct`; graph emission remains unbound.
  - Evidence Targets: algorithm diff, vector manifest, proofs, tests/benchmarks, Supervisor, detect, commit.
  - Actual-status Update: re-export lookup `wrong -> correct`; graph emission remains unbound.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P5-D: Emit terminal edges/proofs and validate barrel consumers.
  - Goal: preserve syntactic module dependency edges while binding calls/uses to terminal `ResolutionTarget`s with proof chains.
  - Scope Boundary:
    - Editable: import/call/use outcome emission adapters.
    - Inspect-only: table/resolver; target read-only.
    - Preserve-only: direct IMPORTS semantics.
    - Out of scope: ambient.
  - Non-Goals: no fake consumer-to-declaration-module IMPORTS edge.
  - Pre-flight Questions:
    - Data source: structured P5 outcomes plus an independent target barrel oracle.
    - Display permission: Preserve existing graph permissions; use existing UI only if this slice exposes fields there.
    - DB read flow: read canonical resolver-edge records on S0-S5 and target `.anvien` after supported in-place analyze.
    - DB write flow: write production graph outcome edges in Anvien; target writes are limited to normal `.anvien` analyze output.
    - Render location: CLI/MCP/HTTP and existing Web only if fields are exposed; all reports remain under Anvien.
    - UI behavior flow: Conditional — validate existing UI only when its proof fields change; otherwise use CLI/MCP/HTTP and graph records.
    - Docker runtime: Conditional — mandatory if Web fields change; otherwise full build plus real non-UI boundaries.
    - Playwright target: Conditional — real built Docker runtime only if Web is part of the slice.
    - Behavior test: `2/2` terminal barrel calls, syntactic IMPORTS preservation, proof hops, no false gaps, S0-S5 parity, and target boundary.
    - Cleanup/quarantine: remove Anvien debug artifacts; never write target reports/probes/fixtures.
    - External side effects: Supported in-place target analyze may update target `.anvien` and ignored guidance timestamps only.
    - N/A notes: S6-S11 are not claimed here and must be classified by P7.
  - Work Steps:
    1. Implement emission/provenance and focused tests; run full build.
       - UI flow check: proof data renders only in approved existing surfaces.
       - DB/data flow check: source site has one outcome and terminal endpoint.
       - Render location check: CLI/MCP/HTTP/Web.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E5-P5D-IMPACT1`, `E5-P5D-SRC1`, `E5-P5D-BUILD1`.
    2. Capture target pre-state/worktree/hash, graph path, and artifact boundary; analyze target in place; compare an independent barrel source/TS-oracle manifest; verify both barrel calls (`2/2`), dependency edges, no false gaps, and projection parity; capture post-state and ignored-guidance timestamp caveat.
       - UI flow check: built runtime if Web changed.
       - DB/data flow check: exact hop chain through barrel to terminal.
       - Render location check: Anvien reports/QA only; no target report/probe.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E5-P5D-TARGET1`, `E5-P5D-PARITY1`, `E5-P5D-BOUNDARY1`, `E5-P5D-ORACLE1`.
  - Implementation Gate: P5-A/B/C PASS; no global fallback remains for explicit imports.
  - Acceptance:
    - Source: `2/2` calls resolve the expected terminal target, imports remain syntactic, and proofs survive each applicable `S0`–`S5` row named by this slice; `S6`–`S11` are not claimed here and must be classified explicitly as preserved or applicable by P7 before closure. The independent oracle agrees and the target source/worktree/artifact boundary is unchanged.
    - Runtime/UI: Conditional — validate existing UI only when its proof fields change; otherwise use CLI/MCP/HTTP and graph records.
    - DB/data: The data/contract condition is: `2/2` calls resolve the expected terminal target, imports remain syntactic, and proofs survive each applicable `S0`–`S5` row named by this slice; `S6`–`S11` are not claimed here and must be classified explicitly as preserved or applicable by P7 before closure. The independent oracle agrees and the target source/worktree/artifact boundary is unchanged.
    - Behavior test: The behavior gate is: `2/2` calls resolve the expected terminal target, imports remain syntactic, and proofs survive each applicable `S0`–`S5` row named by this slice; `S6`–`S11` are not claimed here and must be classified explicitly as preserved or applicable by P7 before closure. The independent oracle agrees and the target source/worktree/artifact boundary is unchanged.
    - Cleanup/quarantine: The boundary/cleanup condition is: `2/2` calls resolve the expected terminal target, imports remain syntactic, and proofs survive each applicable `S0`–`S5` row named by this slice; `S6`–`S11` are not claimed here and must be classified explicitly as preserved or applicable by P7 before closure. The independent oracle agrees and the target source/worktree/artifact boundary is unchanged.
    - Evidence IDs: `E5-P5D-IMPACT1`, `E5-P5D-SRC1`, `E5-P5D-BUILD1`, `E5-P5D-TARGET1`, `E5-P5D-PARITY1`, `E5-P5D-BOUNDARY1`, `E5-P5D-ORACLE1`, `E5-P5D-REVIEW1`
    - Actual-status rows refreshed: barrel resolution `wrong -> correct`.
  - Evidence Targets: target/projection/oracle/boundary manifest, build/runtime, Supervisor, detect, commit.
  - Actual-status Update: barrel resolution `wrong -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

### P6: Ambient/external declaration universe and truthful diagnostics

- Phase Goal: resolve configured TypeScript standard-library and local package declarations without polluting the repository graph or guessing diagnostics.
- Phase Boundary:
  - In scope: declaration-universe interface, project profile, embedded catalog, local package declarations, lazy external Symbols, structured outcomes, graph-health projection.
  - Out of scope: network/package execution and full node_modules graphing.
  - Dependencies: P5 complete.
- Phase Implementation Rule: interface/profile, catalog, declaration-entrypoint candidate lookup, external authorization/materialization, final outcome assembly, and diagnostic projection are separate commits.
- Ordered Slice List:
  - P6-A: Add declaration-universe and project-profile boundary.
  - P6-B: Build and verify the embedded TypeScript stdlib catalog.
  - P6-C1: Resolve declaration entrypoints into immutable candidates.
  - P6-C2: Authorize and materialize referenced external Symbols.
  - P6-C3: Finalize exhaustive resolution outcomes.
  - P6-D: Project resolver outcomes into graph-health diagnostics.

- [ ] P6-A: Add declaration-universe and project-profile boundary.
  - Goal: let resolution request configured intrinsic/ambient/package symbols through a language-scoped interface.
  - Scope Boundary:
    - Editable: universe interface and TS profile owners.
    - Inspect-only: resolver/catalog.
    - Preserve-only: other languages.
    - Out of scope: catalog contents.
  - Non-Goals: no global `defsByName` shared across languages.
  - Pre-flight Questions:
    - Data source: approved P1/P5 contracts, immutable ModuleRef metadata, and TypeScript project configuration.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read project profile fields and declaration-universe capability metadata.
    - DB write flow: write only in-memory/hash-bound profile and universe requests.
    - Render location: profile/capability diagnostic trace.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real declaration-universe/profile boundary after the full build.
    - Behavior test: compiler version, module/moduleResolution, target/lib closure, noLib, types/typeRoots, baseUrl/paths, package-condition handoff, skipLibCheck, triple-slash, and language isolation.
    - Cleanup/quarantine: synthetic profile fixtures.
    - External side effects: Read-only local config; no runtime network/package execution.
    - N/A notes: No catalog contents, UI, or live persistence.
  - Work Steps:
    1. Implement production interface/profile selection and capability outcomes with a hash-bound profile manifest; P6-A consumes P5 `ModuleRef` and does not resolve package conditions.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: profile/config hash included in outcome.
       - Render location check: diagnostic trace.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6A-IMPACT1`, `E6-P6A-SRC1`.
    2. Add isolation/profile tests and run full build.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: Go/Python cannot see TS catalog.
       - Render location check: evidence.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6A-BUILD1`, `E6-P6A-TEST1`.
  - Implementation Gate: P5 project profile stable.
  - Acceptance:
    - Source: version/hash-bound language-scoped interface, complete profile manifest, and explicit unavailable outcomes.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: version/hash-bound language-scoped interface, complete profile manifest, and explicit unavailable outcomes.
    - Behavior test: The behavior gate is: version/hash-bound language-scoped interface, complete profile manifest, and explicit unavailable outcomes.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-BUILD1`, `E6-P6A-TEST1`, `E6-P6A-REVIEW1`
    - Actual-status rows refreshed: declaration universe boundary `missing -> correct`.
  - Evidence Targets: source/build/isolation tests/Supervisor/detect/commit.
  - Actual-status Update: declaration universe boundary `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P6-B: Build and verify the embedded TypeScript stdlib catalog.
  - Goal: ship a compact, pinned, integrity-checked catalog with lib-reference closure and member meanings/signatures.
  - Scope Boundary:
    - Editable: catalog generation source, immutable generated asset, manifest, loader owner.
    - Inspect-only: resolver.
    - Preserve-only: runtime network boundary.
    - Out of scope: package declarations.
  - Non-Goals: no hand-coded `Promise`/`Math` allowlist.
  - Pre-flight Questions:
    - Data source: pinned licensed TypeScript declaration inputs at build time.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read the pinned build-time declaration source and generated catalog manifest.
    - DB write flow: generate the immutable catalog asset/manifest under Anvien source control only.
    - Render location: version/status output and benchmark ledger.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real catalog generator/loader boundary after the full build.
    - Behavior test: source hash, compiler/catalog version, lib closure, Promise/Math members, license/NOTICE, corruption, regeneration determinism, size/load/RSS ceilings.
    - Cleanup/quarantine: remove superseded generated/debug output; retain one authoritative generated asset and manifest.
    - External side effects: Build-time pinned source only; production runtime has no network, Node, package install, or scripts.
    - N/A notes: No UI; generated asset is a build artifact, not a live DB.
  - Work Steps:
    1. Implement production generator/loader and manifest integrity/version checks; regenerate the asset and record provenance/license/NOTICE.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: catalog immutable and hash verified.
       - Render location check: version/status command.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6B-IMPACT1`, `E6-P6B-SRC1`.
    2. Add catalog closure/member tests, run full build, measure binary/catalog/load/RSS.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: profile selects exact lib closure.
       - Render location check: benchmark ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`.
  - Implementation Gate: provenance/license/integrity, regeneration authority, and numeric size/load ceilings are explicit.
  - Acceptance:
    - Source: pinned catalog passes integrity/closure/member tests within `32 MiB`/`250 ms` ceilings; no runtime network.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: pinned catalog passes integrity/closure/member tests within `32 MiB`/`250 ms` ceilings; no runtime network.
    - Behavior test: The behavior gate is: pinned catalog passes integrity/closure/member tests within `32 MiB`/`250 ms` ceilings; no runtime network.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-REVIEW1`
    - Actual-status rows refreshed: stdlib catalog `missing -> correct`.
  - Evidence Targets: generator/manifest/asset diff, build/tests/benchmarks, Supervisor, detect, commit.
  - Actual-status Update: stdlib catalog `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P6-C1: Resolve declaration entrypoints into immutable candidates.
  - Goal: consume immutable P5 `ModuleRef` metadata to select `types`/`typesVersions` or catalog entrypoints and return an immutable `ExternalResolutionCandidate` plus nested stage proof, without publishing a final top-level outcome.
  - Scope Boundary:
    - Editable: declaration-entrypoint/candidate owner.
    - Inspect-only: P5 module-condition resolver, security/materializer, outcome/status owner, call/type/member adapters, and graph-health.
    - Preserve-only: repo definitions.
    - Out of scope: external authorization/materialization, final outcome assembly, and diagnostic projection.
  - Non-Goals: no re-evaluation of package `exports` conditions, no scan of all `.d.ts`, no package install, and no global fallback.
  - Pre-flight Questions:
    - Data source: declaration universe, project profile, and immutable P5 ModuleRef.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read catalog/package metadata to select one authorized declaration entrypoint candidate.
    - DB write flow: store only immutable in-memory candidate/proof results; no external Symbol or final outcome.
    - Render location: entrypoint trace/candidate evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real declaration-entrypoint lookup boundary after the full build.
    - Behavior test: Promise/Math catalog lookup, package present/absent, types/typesVersions, nested proof, determinism, and no package-exports re-evaluation.
    - Cleanup/quarantine: isolated candidate fixtures.
    - External side effects: Read-only authorized local declarations/catalog; no network or package execution.
    - N/A notes: Security/materialization/final outcome/UI are explicitly deferred.
  - Work Steps:
    1. record impacts; implement declaration-entrypoint lookup from immutable `ModuleRef` metadata; add positive/negative candidate tests; run full build; compare pinned entrypoint oracle results; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A; health projection is P6-D.
       - DB/data flow check: one immutable candidate/proof per eligible source site; no final status is emitted.
       - Render location check: entrypoint trace/candidate fields only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6C1-IMPACT1`, `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-ORACLE1`.
  - Implementation Gate: P6-A/B PASS; P5 `ModuleRef` handoff is immutable; P5 owns `exports` conditions and P6-C1 owns only `types`/`typesVersions` or catalog entrypoint selection.
  - Acceptance:
    - Source: candidate selection and nested proof are deterministic; Promise/Math and package fixtures produce the expected candidates without claiming final `resolved_external`; no false local gap or silent fallback.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: candidate selection and nested proof are deterministic; Promise/Math and package fixtures produce the expected candidates without claiming final `resolved_external`; no false local gap or silent fallback.
    - Behavior test: The behavior gate is: candidate selection and nested proof are deterministic; Promise/Math and package fixtures produce the expected candidates without claiming final `resolved_external`; no false local gap or silent fallback.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E6-P6C1-IMPACT1`, `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-ORACLE1`, `E6-P6C1-REVIEW1`
    - Actual-status rows refreshed: declaration-entrypoint candidate lookup `missing -> correct`; authorization/materialization/outcome remain missing; health projection remains wrong.
  - Evidence Targets: entrypoint/candidate source/tests/oracle, Supervisor, detect-changes, commit.
  - Actual-status Update: declaration-entrypoint candidate lookup `missing -> correct`; authorization/materialization/outcome remain missing; health projection remains wrong.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P6-C2: Authorize and materialize referenced external Symbols.
  - Goal: authorize immutable P6-C1 candidates, enforce declaration/security/cache budgets, and materialize generation-bound `ExternalSymbolRef` nodes lazily without creating catalog File nodes.
  - Scope Boundary:
    - Editable: dedicated declaration-security owner, external-symbol materializer, and catalog/package cache owners.
    - Inspect-only: P6-C1 candidate construction, P6-C3 outcome finalizer, and P5 module resolver.
    - Preserve-only: repository definitions.
    - Out of scope: final top-level status precedence and graph-health classification.
  - Non-Goals: no full `.d.ts` graph, no package execution/network, no package-condition lookup, and no diagnostic reclassification.
  - Pre-flight Questions:
    - Data source: P6-C1 candidates plus authorized catalog/package declaration bytes.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read only authorized declaration roots and generation/catalog/config-qualified cache entries.
    - DB write flow: write referenced ExternalSymbols and isolated generation-qualified cache rows only.
    - Render location: materializer/security trace and benchmark evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real external-symbol materializer boundary after the full build.
    - Behavior test: exact node/descriptor count, hashes, invalidation, zero catalog File nodes, allowed-root/symlink rejection, file/byte/depth/cache ceilings, and typed failures.
    - Cleanup/quarantine: remove staged materialization stores.
    - External side effects: Read-only authorized local files/catalog; no network, package install, scripts, or arbitrary execution.
    - N/A notes: No UI or final ResolutionOutcome mutation.
  - Work Steps:
    1. record impacts; implement the dedicated descriptor materializer and cache invalidation; add tests; run full build; measure catalog/package file-byte/cache limits and external-node counts on in-budget and explicit over-limit fixtures; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A; public projection is P7 parity.
       - DB/data flow check: referenced node count equals referenced descriptors, no duplicate external IDs, and no catalog declaration becomes a repository File node.
       - Render location check: materializer trace/benchmark only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6C2-IMPACT1`, `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-BENCH1`.
  - Implementation Gate: P6-C1 PASS; candidate handoff and security/budget ceilings are fixed.
  - Acceptance:
    - Source: lazy external-node count and hashes are exact, invalidation rejects stale catalog/config/package roots, zero catalog File nodes exist, in-budget fixtures are `100%` correct, and named security/over-limit fixtures return typed materialization failures for P6-C3; no final `ResolutionOutcome` is mutated here.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: lazy external-node count and hashes are exact, invalidation rejects stale catalog/config/package roots, zero catalog File nodes exist, in-budget fixtures are `100%` correct, and named security/over-limit fixtures return typed materialization failures for P6-C3; no final `ResolutionOutcome` is mutated here.
    - Behavior test: The behavior gate is: lazy external-node count and hashes are exact, invalidation rejects stale catalog/config/package roots, zero catalog File nodes exist, in-budget fixtures are `100%` correct, and named security/over-limit fixtures return typed materialization failures for P6-C3; no final `ResolutionOutcome` is mutated here.
    - Cleanup/quarantine: The boundary/cleanup condition is: lazy external-node count and hashes are exact, invalidation rejects stale catalog/config/package roots, zero catalog File nodes exist, in-budget fixtures are `100%` correct, and named security/over-limit fixtures return typed materialization failures for P6-C3; no final `ResolutionOutcome` is mutated here.
    - Evidence IDs: `E6-P6C2-IMPACT1`, `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-BENCH1`, `E6-P6C2-REVIEW1`
    - Actual-status rows refreshed: external authorization/materialization/security `missing -> correct`; final outcome and health projection remain missing/wrong.
  - Evidence Targets: materializer source/tests/benchmarks, Supervisor, detect-changes, commit.
  - Actual-status Update: external authorization/materialization/security `missing -> correct`; final outcome and health projection remain missing/wrong.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P6-C3: Finalize exhaustive resolution outcomes.
  - Goal: combine repository targets, P6-C1 candidates, and P6-C2 authorization/materialization results into exactly one immutable top-level `ResolutionOutcome` per source site.
  - Scope Boundary:
    - Editable: structured outcome and status-matrix owners plus thin call/type/member finalization adapters.
    - Inspect-only: entrypoint/materializer internals and graph-health.
    - Preserve-only: graph projection.
    - Out of scope: diagnostic rendering.
  - Non-Goals: no entrypoint lookup, file authorization, node materialization, name heuristic, or later mutation of a finalized outcome.
  - Pre-flight Questions:
    - Data source: immutable P5 repository results, P6-C1 candidates, and P6-C2 materialization results.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read the immutable stage results and authoritative status matrix.
    - DB write flow: write exactly one final in-memory immutable ResolutionOutcome per source site.
    - Render location: finalized resolver trace/outcome evidence.
    - UI behavior flow: N/A — this slice has no browser-visible interaction or installed-app behavior.
    - Docker runtime: N/A — no app/UI runtime is changed; run the full build and validate the named non-UI boundary.
    - Playwright target: N/A — no browser surface is owned; exercise the real outcome finalizer boundary after the full build.
    - Behavior test: positive/negative case for every status, stage precedence, member lookup, target union, severity/actionability/retryability, and resolved/gap exclusivity.
    - Cleanup/quarantine: isolated outcome fixtures.
    - External side effects: None.
    - N/A notes: No entrypoint lookup, materialization, graph-health projection, or UI.
  - Work Steps:
    1. record impacts; implement final precedence and the versioned status matrix; add positive/negative tests for every code after production code; run full build; compare pinned oracle outcomes; refresh ledgers/Supervisor/detect/commit.
       - UI flow check: N/A; graph-health projection is P6-D.
       - DB/data flow check: exactly one top-level immutable outcome per source site; nested module/export/meaning/member proof is retained and no later stage overwrites an earlier terminal failure.
       - Render location check: finalized resolver trace/outcome fields only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6C3-IMPACT1`, `E6-P6C3-SRC1`, `E6-P6C3-BUILD1`, `E6-P6C3-TEST1`, `E6-P6C3-STATUS1`.
  - Implementation Gate: P6-C1/C2 PASS; every candidate/materialization result maps to exactly one status row; severity/actionability/retryability are deterministic.
  - Acceptance:
    - Source: every status code has positive/negative evidence; Promise/Math finalize as `resolved_external` (never `resolved_intrinsic`); missing members use `member_not_found`; no source site is both resolved and a gap; no final outcome can be mutated.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: every status code has positive/negative evidence; Promise/Math finalize as `resolved_external` (never `resolved_intrinsic`); missing members use `member_not_found`; no source site is both resolved and a gap; no final outcome can be mutated.
    - Behavior test: The behavior gate is: every status code has positive/negative evidence; Promise/Math finalize as `resolved_external` (never `resolved_intrinsic`); missing members use `member_not_found`; no source site is both resolved and a gap; no final outcome can be mutated.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E6-P6C3-IMPACT1`, `E6-P6C3-SRC1`, `E6-P6C3-BUILD1`, `E6-P6C3-TEST1`, `E6-P6C3-STATUS1`, `E6-P6C3-REVIEW1`
    - Actual-status rows refreshed: structured outcome/status `missing/wrong -> correct`; health projection remains wrong.
  - Evidence Targets: outcome/status source/tests/oracle, Supervisor, detect-changes, commit.
  - Actual-status Update: structured outcome/status `missing/wrong -> correct`; health projection remains wrong.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P6-D: Project resolver outcomes into graph-health diagnostics.
  - Goal: make diagnostics a faithful projection of structured outcomes and remove target-name inference.
  - Scope Boundary:
    - Editable: outcome-to-health projection owner and narrow graph-health adapter.
    - Inspect-only: resolver.
    - Preserve-only: unrelated health policies.
    - Out of scope: resolver behavior.
  - Non-Goals: no special cases for Promise/Math.
  - Pre-flight Questions:
    - Data source: final immutable ResolutionOutcomes plus an independent Promise/Math oracle.
    - Display permission: Preserve existing diagnostic permissions; use public UI only when diagnostics are exposed there.
    - DB read flow: read outcome records through Graph JSON/Ladybug and applicable S0-S11 diagnostic surfaces.
    - DB write flow: write production health projection in Anvien; target writes limited to supported `.anvien` analyze output.
    - Render location: graph-health CLI/MCP/HTTP and existing Web when applicable; reports under Anvien.
    - UI behavior flow: Conditional — if diagnostics are Web-visible, exercise truthful resolved/gap states; otherwise use real graph-health boundaries.
    - Docker runtime: Conditional — mandatory if Web diagnostics are in scope.
    - Playwright target: Conditional — real built Docker runtime only when Web diagnostics are exposed.
    - Behavior test: every status maps mechanically, no heuristic site remains, exact target Promise/Math `3/3`, no resolved+gap overlap, parity, and target boundary.
    - Cleanup/quarantine: remove Anvien debug artifacts; never target reports/probes/fixtures.
    - External side effects: Supported target analyze may update target `.anvien` and ignored guidance timestamps only.
    - N/A notes: UI/Docker/Playwright are conditional on actual public diagnostic exposure.
  - Work Steps:
    1. Implement production projection and retire same-invariant heuristic classification.
       - UI flow check: diagnostics remain truthful in existing UI.
       - DB/data flow check: status/stage/proof survive persistence.
       - Render location check: graph-health CLI/MCP/HTTP/Web.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6D-IMPACT1`, `E6-P6D-SRC1`.
    2. Add exhaustive status tests for every status code, run full build and real graph-health boundaries, then capture target pre-state/hash/graph path, independent Promise/Math oracle, and target artifact manifest before analyzing `E:\cheapapp.org`; capture post-state and ignored-guidance timestamp caveat.
       - UI flow check: built Docker/Playwright if Web-visible.
       - DB/data flow check: no source site both resolved and unresolved.
       - Render location check: official evidence under Anvien; no target report/probe.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E6-P6D-BUILD1`, `E6-P6D-TARGET1`, `E6-P6D-BOUNDARY1`, `E6-P6D-ORACLE1`.
  - Implementation Gate: P6-C1/C2/C3 candidates, materializer, final outcomes/status matrix, and security budgets are correct; every old heuristic site is inventoried; target oracle and boundary procedure are ready.
  - Acceptance:
    - Source: exact diagnostic status across each applicable `S0`–`S11` surface; the bounded target Promise/Math sites are exactly `resolved_external` (`3/3`), while `external_declaration_unavailable` is accepted only in explicit negative fixtures; no false local gap exists, the independent oracle agrees, and the target source/worktree/artifact boundary is unchanged.
    - Runtime/UI: Conditional — if diagnostics are Web-visible, exercise truthful resolved/gap states; otherwise use real graph-health boundaries.
    - DB/data: The data/contract condition is: exact diagnostic status across each applicable `S0`–`S11` surface; the bounded target Promise/Math sites are exactly `resolved_external` (`3/3`), while `external_declaration_unavailable` is accepted only in explicit negative fixtures; no false local gap exists, the independent oracle agrees, and the target source/worktree/artifact boundary is unchanged.
    - Behavior test: The behavior gate is: exact diagnostic status across each applicable `S0`–`S11` surface; the bounded target Promise/Math sites are exactly `resolved_external` (`3/3`), while `external_declaration_unavailable` is accepted only in explicit negative fixtures; no false local gap exists, the independent oracle agrees, and the target source/worktree/artifact boundary is unchanged.
    - Cleanup/quarantine: The boundary/cleanup condition is: exact diagnostic status across each applicable `S0`–`S11` surface; the bounded target Promise/Math sites are exactly `resolved_external` (`3/3`), while `external_declaration_unavailable` is accepted only in explicit negative fixtures; no false local gap exists, the independent oracle agrees, and the target source/worktree/artifact boundary is unchanged.
    - Evidence IDs: `E6-P6D-IMPACT1`, `E6-P6D-SRC1`, `E6-P6D-BUILD1`, `E6-P6D-TARGET1`, `E6-P6D-BOUNDARY1`, `E6-P6D-ORACLE1`, `E6-P6D-REVIEW1`
    - Actual-status rows refreshed: diagnostic classification `wrong -> correct`.
  - Evidence Targets: heuristic-inventory delta, status matrix, build/tests/runtime/target/oracle/boundary, Supervisor, detect, commit.
  - Actual-status Update: diagnostic classification `wrong -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

### P6 status matrix (authoritative contract)

P6-C3 records this matrix in the versioned outcome-status owner and every code gets a negative/positive test. `Target` is the tagged union defined above; `GapRef` preserves candidates and nested stage proof.

| Code | Stage | Allowed target | Severity | Actionability | Retry policy | Graph-health projection |
|------|-------|----------------|----------|---------------|--------------|-------------------------|
| `resolved_local` | symbol/namespace | `InternalSymbolRef` or `ModuleNamespaceRef` | info | none | no retry | resolved/local |
| `resolved_external` | declaration | `ExternalSymbolRef` | info | none | no retry | resolved/external |
| `resolved_intrinsic` | language primitive | `IntrinsicSymbolRef` | info | none | no retry | resolved/intrinsic |
| `module_not_found` | module | `GapRef` | error | source/config | retry after source/config change | resolution_gap/module |
| `module_blocked_by_exports` | module | `GapRef` | error | package-exports/config | retry after package/config change | resolution_gap/exports |
| `export_not_found` | export | `GapRef` | error | source/export surface | retry after source change | resolution_gap/export |
| `meaning_mismatch` | meaning | `GapRef` plus candidates | error | source/type use | retry after source/use change | meaning_gap |
| `member_not_found` | member | `GapRef` plus namespace/type candidate | error | source/declaration member | retry after source/declaration change | member_gap |
| `ambiguous_export` | export | `GapRef` plus all candidates | error | source/export authority | no automatic retry | ambiguity_gap |
| `cycle_no_terminal` | export | `GapRef` plus cycle proof | error | source/module graph | retry after source change | resolution_gap/cycle |
| `config_invalid` | module/profile | `GapRef` | error | project config | retry after config change | config_gap |
| `external_declaration_unavailable` | declaration | `GapRef` | error | install/catalog/profile | retry after declaration availability change | external_gap |
| `catalog_version_mismatch` | declaration/catalog | `GapRef` | error | analyzer/catalog build | retry after rebuild | catalog_gap |
| `declaration_parse_failed` | declaration | `GapRef` | error | declaration input | retry after input change | external_gap/parse |
| `unsupported_syntax` | syntax | `GapRef` | error | analyzer capability | retry after analyzer upgrade | capability_gap |
| `budget_exceeded` | traversal/catalog/security | `GapRef` plus measured limit | error | scope/config | retry only with explicit budget change | budget_gap |
| `security_path_rejected` | path | `GapRef` | error | allowed-root/config | no automatic retry | security_gap |

The top-level code is immutable for a source site; nested module/export/meaning/member failures are retained as proof and cannot overwrite the top-level stage. Graph-health maps this table mechanically and never infers a category from target text.

### P7: Cross-surface acceptance, target validation, and performance

- Phase Goal: prove the complete invariant on Anvien and the real target before closure.
- Phase Boundary:
  - In scope: determinism, closure, version/fault matrix, target oracles, all command surfaces, Docker/Web, performance/capacity.
  - Out of scope: new product behavior or unrelated fixes.
  - Dependencies: P1-P6 complete.
- Phase Implementation Rule: validation findings reopen only the responsible prior slice; P7 does not patch code.
- Ordered Slice List:
  - P7-A: Run determinism, closure, version, and failure-atomicity gates.
  - P7-B: Run bounded `cheapapp.org` in-place acceptance.
  - P7-C: Run full runtime/projection/performance acceptance.

- [ ] P7-A: Run determinism, closure, version, and failure-atomicity gates.
  - Goal: prove graph/index invariants across repeated builds and hostile failures.
  - Scope Boundary:
    - Editable: evidence/benchmark/status only.
    - Inspect-only: all implementation.
    - Preserve-only: source.
    - Out of scope: fixes.
  - Non-Goals: a failed gate is not repaired in this slice.
  - Pre-flight Questions:
    - Data source: all accepted and committed implementation slices plus the final fault/version inventories.
    - Display permission: N/A — this slice introduces no user-visible display or permission decision.
    - DB read flow: read completed generation artifacts, reader results, and retained prior generations.
    - DB write flow: write isolated validation generations and Anvien evidence only.
    - Render location: structured CLI/API error, evidence, and benchmark captures.
    - UI behavior flow: N/A — public visual runtime acceptance belongs to P7-C; API errors are validated at native boundaries.
    - Docker runtime: N/A — P7-C owns the final built Docker runtime; P7-A validates non-UI determinism/version/fault boundaries.
    - Playwright target: N/A — exercise real loaders/CLI/MCP/HTTP fault boundaries; P7-C owns browser evidence.
    - Behavior test: five order/worker analyzes, exact canonical sets, duplicates/endpoints/orphans, every version row, every fault, and prior-generation queryability.
    - Cleanup/quarantine: remove temporary validation generations after evidence and lease-safe checks.
    - External side effects: Local Anvien runtime and isolated stores only.
    - N/A notes: No repair is allowed; failure reopens the responsible prior slice.
  - Work Steps:
    1. Full build, five analyzes with varied worker/order conditions, canonical set/closure/parity comparison.
       - UI flow check: N/A — this work step has no browser-visible or installed-app behavior; Mini-QA exercises the named nearest non-UI boundary.
       - DB/data flow check: exact generation-consistent nodes/edges/refs.
       - Render location check: evidence/benchmark.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E7-P7A-BUILD1`, `E7-P7A-DETERMINISM1`.
    2. Execute version and fault matrix; query retained generation after each failure.
       - UI flow check: API/CLI errors visible where applicable.
       - DB/data flow check: prior hashes/queryability unchanged.
       - Render location check: structured captures.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E7-P7A-VERSION1`, `E7-P7A-FAULT1`.
  - Implementation Gate: all implementation slices PASS and committed.
  - Acceptance:
    - Source: all zero-count, determinism, mismatch, and rollback gates pass.
    - Runtime/UI: N/A — no separate runtime/UI condition is stated in the compact acceptance; the nearest real non-UI boundary is covered by the work-step checks.
    - DB/data: The data/contract condition is: all zero-count, determinism, mismatch, and rollback gates pass.
    - Behavior test: The behavior gate is: all zero-count, determinism, mismatch, and rollback gates pass.
    - Cleanup/quarantine: N/A — no separate cleanup/quarantine condition is stated in the compact acceptance; the existing Scope Boundary, Non-Goals, and work-step cleanup checks remain authoritative.
    - Evidence IDs: `E7-P7A-BUILD1`, `E7-P7A-DETERMINISM1`, `E7-P7A-VERSION1`, `E7-P7A-FAULT1`, `E7-P7A-REVIEW1`, `E7-P7A-DETECT1`, `E7-P7A-COMMIT1`
    - Actual-status rows refreshed: global identity/generation invariant confirmed or responsible slice reopened.
  - Evidence Targets: manifests, fault captures, Supervisor verdict.
  - Actual-status Update: global identity/generation invariant confirmed or responsible slice reopened.
  - Commit Boundary: run Anvien-side change/boundary detection as `E7-P7A-DETECT1`, then commit validation ledgers/report only as `E7-P7A-COMMIT1`; no implementation is mixed in.

- [ ] P7-B: Run bounded `cheapapp.org` in-place acceptance.
  - Goal: prove the five repaired defects at the original real-source sites without contaminating the target.
  - Scope Boundary:
    - Editable: Anvien evidence/reports only.
    - Inspect-only: inspect/read-analyze target in place.
    - Preserve-only: target source/worktree.
    - Out of scope: scanner fix/global accuracy.
  - Non-Goals: `887/895` scanner baseline is not promoted to failure.
  - Pre-flight Questions:
    - Data source: recorded target HEAD/worktree plus independent bounded source/TypeScript oracles.
    - Display permission: Use existing graph permissions only when Web parity is part of the accepted matrix.
    - DB read flow: read target active `.anvien` and exact applicable S0-S11 rows.
    - DB write flow: normal supported in-place target analyze output under `.anvien` only.
    - Render location: all evidence/reports under Anvien; existing graph UI only when applicable.
    - UI behavior flow: Conditional — exercise existing graph UI only for applicable Web parity.
    - Docker runtime: Conditional — mandatory when Web is checked.
    - Playwright target: Conditional — real built Docker target only when S6 applies.
    - Behavior test: five defect families, `6/6`, `21/21`, two time, two now, `2/2` barrels, Promise/Math `3/3`, exact scanner quarantine, and contamination boundary.
    - Cleanup/quarantine: remove Anvien debug artifacts; never target reports/probes/fixtures.
    - External side effects: Supported target analyze may update `.anvien` and ignored guidance timestamps only.
    - N/A notes: Scanner repair remains excluded; UI/Docker/Playwright apply only to applicable Web rows.
  - Work Steps:
    1. Capture target pre-state, analyze in place, verify generation/version/hash and exact `6/6`, `21/21`, two `time`, two `now`, `2/2` barrel calls, and Promise/Math outcomes.
       - UI flow check: existing graph UI only if part of parity.
       - DB/data flow check: raw graph/Ladybug/CLI/MCP/HTTP agree.
       - Render location check: evidence stays in Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E7-P7B-TARGET1`, `E7-P7B-ORACLE1`.
    2. Compare post-state, exact File-node omissions, and artifact manifest; run independent source oracle and Supervisor.
       - UI flow check: built runtime if used.
       - DB/data flow check: no new omission/orphan/mixed generation.
       - Render location check: Anvien report only.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E7-P7B-BOUNDARY1`, `E7-P7B-REVIEW1`.
  - Implementation Gate: target owner boundary and pre-existing changes captured; operational staging remains under target `.anvien`.
  - Acceptance:
    - Source: every bounded semantic gate passes, exact scanner omissions unchanged, no target contamination.
    - Runtime/UI: Conditional — exercise existing graph UI only for applicable Web parity.
    - DB/data: The data/contract condition is: every bounded semantic gate passes, exact scanner omissions unchanged, no target contamination.
    - Behavior test: The behavior gate is: every bounded semantic gate passes, exact scanner omissions unchanged, no target contamination.
    - Cleanup/quarantine: The boundary/cleanup condition is: every bounded semantic gate passes, exact scanner omissions unchanged, no target contamination.
    - Evidence IDs: `E7-P7B-TARGET1`, `E7-P7B-ORACLE1`, `E7-P7B-BOUNDARY1`, `E7-P7B-REVIEW1`, `E7-P7B-DETECT1`, `E7-P7B-COMMIT1`
    - Actual-status rows refreshed: five defect rows become correct or responsible phase reopens.
  - Evidence Targets: target manifests/oracles/parity/boundary/Supervisor.
  - Actual-status Update: five defect rows become correct or responsible phase reopens.
  - Commit Boundary: commit Anvien-side evidence/report only after `E7-P7B-DETECT1`; never target artifacts.

- [ ] P7-C: Run full runtime/projection/performance acceptance.
  - Goal: prove packaged CLI/MCP/HTTP/Web/Ladybug parity and final capacity budgets on built artifacts.
  - Scope Boundary:
    - Editable: reusable QA/evidence/benchmark/status only.
    - Inspect-only: source and built artifacts.
    - Preserve-only: behavior.
    - Out of scope: fixes.
  - Non-Goals: no host dev-server substitution.
  - Pre-flight Questions:
    - Data source: the final active generation, frozen reader matrix, canonical manifest, and P2-G five-run baseline.
    - Display permission: Preserve existing permissions across the real packaged CLI/MCP/HTTP/Web runtime.
    - DB read flow: read all S0-S11 surfaces, caches, embeddings, registries/groups, and derived projections for one generation/vector.
    - DB write flow: write only normal built-runtime validation state and Anvien evidence/benchmarks.
    - Render location: real Docker URL, native command outputs, and official QA artifacts.
    - UI behavior flow: Exercise supported graph, mismatch, stream, lifecycle, and diagnostic states on the real built runtime.
    - Docker runtime: Mandatory; full package/build includes Docker image/container build and the real container must start.
    - Playwright target: Mandatory against the real built Docker URL with reusable scripts, JSON+MD, screenshots/traces, and visual inspection.
    - Behavior test: every exact matrix row, all canonical fields/orphans, native/fallback separation, five-run performance/capacity, and zero unlisted readers.
    - Cleanup/quarantine: remove obsolete/debug QA; retain official JSON+MD/visual evidence.
    - External side effects: Local built Docker/runtime and repo-local operational index only.
    - N/A notes: No dev-server substitution and no repair inside acceptance.
  - Work Steps:
    1. Run full package/build including Docker, start the real container/runtime, and exercise `S0` Graph JSON, `S1` Ladybug native Cypher, `S2` Go/fallback Cypher, the exact current reader-matrix row unions for `S3` CLI/`S4` MCP/`S5` HTTP/`S6` Web, `S7` file-context caches, `S8` HTTP/MCP resource caches, `S9` embeddings, `S10` repo/group registries and contracts, and `S11` process/community projections. A parent router/command does not substitute for any child row; the run records `rows_passed`, `rows_total`, and `unlisted_readers`.
       - UI flow check: supported graph and error states.
       - DB/data flow check: compare canonical node/edge records including IDs, labels, endpoints, ranges, selection ranges, meanings, export/access fields, generation/version, provenance, external refs, cache keys, embedding dimension/hash, group records, and source-anchored process/community membership/order; run native and fallback Cypher separately.
       - Render location check: real exposed URL and official QA directory.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E7-P7C-BUILD1`, `E7-P7C-RUNTIME1`, `E7-P7C-PLAY1`, and one parity record each for `E7-P7C-S0-PARITY1` through `E7-P7C-S11-PARITY1`.
    2. Measure analyze/DB-load/native-query p95/fallback-query p95/RSS/graph/catalog/binary metrics over at least five runs and compare the exact P2-G baseline methodology; publish separate native and fallback query regressions and do not average them.
       - UI flow check: N/A for performance.
       - DB/data flow check: same input/generation for comparisons.
       - Render location check: benchmark ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E7-P7C-BENCH1`, `E7-P7C-NATIVEBENCH1`, `E7-P7C-FALLBACKBENCH1`.
  - Implementation Gate: Docker and packaged runtime available; baseline/final methodology identical.
  - Acceptance:
    - Source: final reader/cache/embedding/group/process/community matrix is complete with zero unlisted surface.
    - Runtime/UI: built Docker runtime and official Playwright supported/error states pass.
    - DB/data: every canonical field and external reference has `differing_records == 0` and `orphan_refs == 0` on each independent `S0`–`S11` row; native and fallback Cypher are reported separately.
    - Behavior test: native/fallback Cypher and all CLI/MCP/HTTP/Web/cache/embedding/group/derived probes pass independently.
    - Cleanup/quarantine: debug/obsolete QA artifacts removed; official JSON+MD evidence retained.
    - Evidence IDs: `E7-P7C-BUILD1`, `E7-P7C-RUNTIME1`, `E7-P7C-PLAY1`, `E7-P7C-BENCH1`, `E7-P7C-NATIVEBENCH1`, `E7-P7C-FALLBACKBENCH1`, `E7-P7C-S0-PARITY1` through `E7-P7C-S11-PARITY1`, `E7-P7C-REVIEW1`, `E7-P7C-DETECT1`, `E7-P7C-COMMIT1`.
    - Actual-status rows refreshed: every cross-surface row and performance metric.
  - Evidence Targets: complete parity matrix, build/container/runtime/Playwright/parity/benchmark/Supervisor.
  - Actual-status Update: cross-surface/runtime/performance rows final.
  - Commit Boundary: run Anvien-side change/boundary detection as `E7-P7C-DETECT1`, then commit validation artifacts as `E7-P7C-COMMIT1` only after acceptance.

### P8: Independent acceptance, dead-work cleanup, and closure

- Phase Goal: obtain a zero-trust verdict, remove only superseded plan-created work, and close the ledgers/commits without hiding any blocker.
- Phase Boundary:
  - In scope: full-plan Supervisor review, dead-work lifecycle, final build/runtime/evidence/detect/commit state.
  - Out of scope: new implementation behavior or unrelated cleanup.
  - Dependencies: P1-P7 complete or explicitly blocked with evidence.
- Phase Implementation Rule: P8 reviews and closes; it does not repair a rejected invariant. A rejection reopens only the responsible earlier slice.
- Ordered Slice List:
  - P8-A: Call Supervisor for the implemented-plan acceptance loop.
  - P8-B: Remove dead work created during this plan.
  - P8-C: Close the plan.

- [ ] P8-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, changed files, generated output, and validation results before closure.
  - Work Steps:
    1. Call the supervisor skill to review the full completed plan work.
    2. If supervisor fails the work, return to the responsible implementation workflow/skill for the failed scope only.
    3. Re-run supervisor review after the fix.
    4. Repeat until supervisor passes or records a blocker.
  - Implementation Gate: all planned implementation phases must be completed or explicitly blocked before this review.
  - Acceptance: supervisor review passes with `E8-P8A-REVIEW1`, or the plan records a blocker with that exact evidence and no closure is performed.
- [ ] P8-B: Remove dead work created during this plan.
  - Goal: ensure the final diff contains only artifacts that still serve the accepted plan.
  - Work Steps:
    1. Review files, sections, generated output, tests, temp files, and plan artifacts created or modified during this plan.
    2. Remove or rewrite any artifact made obsolete by actual-status findings, user corrections, failed approaches, or phase status updates.
    3. Verify no rejected approach, stale placeholder, unused generated output, or dead helper artifact remains in the final diff.
    4. Call supervisor to review the dead-work cleanup.
    5. If supervisor fails the cleanup, return to the responsible implementation workflow/skill for the failed cleanup scope only, then re-run supervisor review.
  - Implementation Gate: only remove artifacts created by this plan unless the user explicitly approves broader cleanup.
  - Acceptance: final `git diff/status` contains no dead plan-created artifacts, Supervisor passes the cleanup in `E8-P8B-REVIEW1`, and `E8-P8B-CLEAN1` records what was removed or preserved.
- [ ] P8-C: Close the plan.
  - Goal: finish validation, evidence, benchmark, detect-changes, commit, and final status.
  - Scope Boundary:
    - Editable: final Anvien plan/evidence/benchmark/actual-status and approved review artifacts only.
    - Inspect-only: production source and target.
    - Preserve-only: target source/worktree.
    - Out of scope: new implementation behavior or unrelated cleanup.
  - Non-Goals: no new production behavior, no repair inside closure, no target mutation beyond already-approved validation, and no unrelated cleanup.
  - Pre-flight Questions:
    - Data source: the completed accepted plan, all four ledgers, final commits, generated outputs, and Supervisor reports.
    - Display permission: Preserve the approved public permission/visibility behavior verified by the completed slices.
    - DB read flow: read final production artifacts and the active generation/runtime state needed by closure evidence.
    - DB write flow: write only final Anvien ledgers/review artifacts and any source-of-truth-required regenerated output.
    - Render location: real built runtime for public changes and final Anvien closure report.
    - UI behavior flow: Re-run the already-approved public supported/error states only when final closure requires freshness; do not add behavior.
    - Docker runtime: Mandatory when any completed slice changed app/runtime behavior; otherwise record the exact non-app reason from actual-status.
    - Playwright target: Mandatory against the real built Docker URL when public runtime is in scope; otherwise record the exact N/A reason.
    - Behavior test: final build/runtime freshness, generated-output consistency, detect-changes, commit hashes, worktree state, and no dead work.
    - Cleanup/quarantine: remove only superseded plan-created debug/dead artifacts; retain accepted evidence.
    - External side effects: Limited to the final Anvien documentation/review commit and already-approved local runtime validation.
    - N/A notes: No new implementation, repair, target mutation, or unrelated cleanup is permitted.
  - Work Steps:
    1. Run the required final validation for the accepted scope, including full build before final runtime validation. For app/runtime scopes, full build must include Docker image/container build.
       - UI flow check: exercise only the already-approved supported/error states; add no behavior.
       - DB/data flow check: the full build and final validation use one accepted generation/epoch/vector and the final committed source.
       - Render location check: real built runtime when public changes exist; otherwise the exact native non-UI boundary recorded by actual-status.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E8-P8C-BUILD1`.
    2. Start the real built Docker/container runtime for app/runtime validation. If Docker cannot be built or started, record the blocker and do not substitute a host dev server.
       - UI flow check: open the real built runtime and confirm the previously accepted startup/supported/error states mount.
       - DB/data flow check: the container pins the same final manifest/generation used by closure validation.
       - Render location check: the real exposed Docker/container URL.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E8-P8C-RUNTIME1`.
    3. Validate public runtime or UI-facing changes with browser or Playwright evidence against the real built Docker/container runtime. Playwright evidence must include Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
       - UI flow check: re-run the accepted public triggers and supported/mismatch/error states; visually inspect every retained screenshot.
       - DB/data flow check: UI-visible records and errors match the final active generation and canonical contracts.
       - Render location check: the real Docker URL and official `Reports/qa/playwright/...` JSON/Markdown/visual artifacts.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E8-P8C-PLAY1`.
    4. Regenerate generated outputs if source-of-truth changes require it.
       - UI flow check: N/A — regeneration owns generated contracts/assets, not a new interaction; any affected runtime is rechecked by step 3.
       - DB/data flow check: generated output hashes and source-of-truth inputs are current and no stale generated contract remains.
       - Render location check: generated output in its repository-native owner plus the Anvien regeneration manifest.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E8-P8C-REGEN1`.
    5. Run Anvien detect-changes before commit when implementation work was performed.
       - UI flow check: N/A — change detection is a non-UI graph/impact boundary.
       - DB/data flow check: the refreshed Anvien graph and final worktree diff are the exact inputs; affected files/symbols/flows are retained without truncation.
       - Render location check: Anvien detect-changes capture in the evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E8-P8C-DETECT1`.
    6. Record final validation, detect-changes, benchmark, and commit evidence.
       - UI flow check: N/A — ledger reconciliation records already-observed behavior and creates no new UI state.
       - DB/data flow check: every plan/actual-status/evidence/benchmark reference resolves to an exact current evidence ID and measurement.
       - Render location check: the four standard ledgers and final review report under Anvien.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E8-P8C-LEDGER1`.
    7. Commit the completed scope and verify the worktree state.
       - UI flow check: N/A — commit/worktree verification is a Git boundary.
       - DB/data flow check: the commit contains only accepted final Anvien artifacts and no target/dead/debug file.
       - Render location check: Git commit hash and post-commit worktree capture in the evidence ledger.
       - Mini QA for each completed implementation slice (MUST) — this work step:
         - Codex browser/control: use Browser, Chrome, or Computer Use against the real mounted runtime when applicable; otherwise record `N/A` with the non-UI reason.
         - Playwright: use reusable Playwright against the built Docker runtime for UI/runtime work; otherwise record `N/A` and name the nearest CLI/MCP/loader/resolver boundary.
         - Other agents: use the equivalent browser/session/computer-control or Playwright-like capability exposed by that runtime.
         - Evidence: record the command, target runtime or boundary, observed result, visual inspection when applicable, and exact evidence ID.
       - Evidence target: `E8-P8C-COMMIT1`.
  - Implementation Gate: P8-A and P8-B must pass or record blockers.
  - Acceptance:
    - Source: final build/runtime/UI/regeneration/ledger evidence is `E8-P8C-BUILD1`, `E8-P8C-RUNTIME1`, `E8-P8C-PLAY1`, `E8-P8C-REGEN1`, and `E8-P8C-LEDGER1`; final change detection is `E8-P8C-DETECT1`; final closure commit/worktree evidence is `E8-P8C-COMMIT1`; required commits exist and the worktree state is known.
    - Runtime/UI: The runtime/UI condition is: final build/runtime/UI/regeneration/ledger evidence is `E8-P8C-BUILD1`, `E8-P8C-RUNTIME1`, `E8-P8C-PLAY1`, `E8-P8C-REGEN1`, and `E8-P8C-LEDGER1`; final change detection is `E8-P8C-DETECT1`; final closure commit/worktree evidence is `E8-P8C-COMMIT1`; required commits exist and the worktree state is known.
    - DB/data: The data/contract condition is: final build/runtime/UI/regeneration/ledger evidence is `E8-P8C-BUILD1`, `E8-P8C-RUNTIME1`, `E8-P8C-PLAY1`, `E8-P8C-REGEN1`, and `E8-P8C-LEDGER1`; final change detection is `E8-P8C-DETECT1`; final closure commit/worktree evidence is `E8-P8C-COMMIT1`; required commits exist and the worktree state is known.
    - Behavior test: The behavior gate is: final build/runtime/UI/regeneration/ledger evidence is `E8-P8C-BUILD1`, `E8-P8C-RUNTIME1`, `E8-P8C-PLAY1`, `E8-P8C-REGEN1`, and `E8-P8C-LEDGER1`; final change detection is `E8-P8C-DETECT1`; final closure commit/worktree evidence is `E8-P8C-COMMIT1`; required commits exist and the worktree state is known.
    - Cleanup/quarantine: The boundary/cleanup condition is: final build/runtime/UI/regeneration/ledger evidence is `E8-P8C-BUILD1`, `E8-P8C-RUNTIME1`, `E8-P8C-PLAY1`, `E8-P8C-REGEN1`, and `E8-P8C-LEDGER1`; final change detection is `E8-P8C-DETECT1`; final closure commit/worktree evidence is `E8-P8C-COMMIT1`; required commits exist and the worktree state is known.
    - Evidence IDs: `E8-P8C-BUILD1`, `E8-P8C-RUNTIME1`, `E8-P8C-PLAY1`, `E8-P8C-REGEN1`, `E8-P8C-LEDGER1`, `E8-P8C-DETECT1`, `E8-P8C-COMMIT1`
    - Actual-status rows refreshed: plan status `draft / P0 complete / implementation blocked` becomes `closed` only after all required authority and evidence gates pass; otherwise retain the exact blocker.
  - Evidence Targets: final build/runtime, detect-changes, closure report, commit, and worktree state.
  - Actual-status Update: plan status `draft / P0 complete / implementation blocked` becomes `closed` only after all required authority and evidence gates pass; otherwise retain the exact blocker.
  - Commit Boundary: commit only the final Anvien documentation/review artifacts after closure evidence is recorded; never include target artifacts or production code.

## Slice-to-ledger traceability binding

The following exact IDs are mandatory for each named slice. A prose phrase such as “Supervisor, detect-changes, commit” never substitutes for exact `DETECT1` and `COMMIT1` IDs. `REVIEW1` is the independent Supervisor gate; implementation slices cannot advance while any ID in their row is pending.

| Slice | Required exact evidence IDs |
|-------|------------------------------|
| P1-A | `E1-P1A-CONTRACT1`, `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1` |
| P1-B | `E1-P1B-IMPACT1`, `E1-P1B-SRC1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-REVIEW1`, `E1-P1B-DETECT1`, `E1-P1B-COMMIT1` |
| P1-C0 | `E1-P1C0-IMPACT1`, `E1-P1C0-SRC1`, `E1-P1C0-BUILD1`, `E1-P1C0-TEST1`, `E1-P1C0-REVIEW1`, `E1-P1C0-DETECT1`, `E1-P1C0-COMMIT1` |
| P1-C0A | `E1-P1C0A-IMPACT1`, `E1-P1C0A-SRC1`, `E1-P1C0A-BUILD1`, `E1-P1C0A-TEST1`, `E1-P1C0A-REVIEW1`, `E1-P1C0A-DETECT1`, `E1-P1C0A-COMMIT1` |
| P1-C0B | `E1-P1C0B-IMPACT1`, `E1-P1C0B-SRC1`, `E1-P1C0B-BUILD1`, `E1-P1C0B-TEST1`, `E1-P1C0B-REVIEW1`, `E1-P1C0B-DETECT1`, `E1-P1C0B-COMMIT1` |
| P1-C | `E1-P1C-IMPACT1`, `E1-P1C-SRC1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` |
| P1-D | `E1-P1D-IMPACT1`, `E1-P1D-MAP1`, `E1-P1D-SRC1`, `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-REVIEW1`, `E1-P1D-DETECT1`, `E1-P1D-COMMIT1` |
| P1-D1 | `E1-P1D1-IMPACT1`, `E1-P1D1-SRC1`, `E1-P1D1-BUILD1`, `E1-P1D1-TEST1`, `E1-P1D1-REVIEW1`, `E1-P1D1-DETECT1`, `E1-P1D1-COMMIT1` |
| P1-D2 | `E1-P1D2-IMPACT1`, `E1-P1D2-SRC1`, `E1-P1D2-BUILD1`, `E1-P1D2-TEST1`, `E1-P1D2-REVIEW1`, `E1-P1D2-DETECT1`, `E1-P1D2-COMMIT1` |
| P1-D3 | `E1-P1D3-IMPACT1`, `E1-P1D3-SRC1`, `E1-P1D3-BUILD1`, `E1-P1D3-TEST1`, `E1-P1D3-REVIEW1`, `E1-P1D3-DETECT1`, `E1-P1D3-COMMIT1` |
| P1-E | `E1-P1E-IMPACT1`, `E1-P1E-SRC1`, `E1-P1E-SHADOW1`, `E1-P1E-TARGET1`, `E1-P1E-ORACLE1`, `E1-P1E-BOUNDARY1`, `E1-P1E-BUILD1`, `E1-P1E-BENCH1`, `E1-P1E-REVIEW1`, `E1-P1E-DETECT1`, `E1-P1E-COMMIT1` |
| P2-A | `E2-P2A-IMPACT1`, `E2-P2A-SRC1`, `E2-P2A-BUILD1`, `E2-P2A-TEST1`, `E2-P2A-REVIEW1`, `E2-P2A-DETECT1`, `E2-P2A-COMMIT1` |
| P2-A1 | `E2-P2A1-MATRIX1`, `E2-P2A1-MATRIXREVIEW1`, `E2-P2A1-R01..E2-P2A1-R195`, `E2-P2A1-REVIEW1`, `E2-P2A1-COMMIT1` |
| P2-A2 | `E2-P2A2-IMPACT1`, `E2-P2A2-SRC1`, `E2-P2A2-BUILD1`, `E2-P2A2-TEST1`, `E2-P2A2-S0GUARD1`, `E2-P2A2-REVIEW1`, `E2-P2A2-DETECT1`, `E2-P2A2-COMMIT1` |
| P2-A3 | `E2-P2A3-IMPACT1`, `E2-P2A3-SRC1`, `E2-P2A3-BUILD1`, `E2-P2A3-TEST1`, `E2-P2A3-S1GUARD1`, `E2-P2A3-REVIEW1`, `E2-P2A3-DETECT1`, `E2-P2A3-COMMIT1` |
| P2-A4 | `E2-P2A4-IMPACT1`, `E2-P2A4-SRC1`, `E2-P2A4-BUILD1`, `E2-P2A4-TEST1`, `E2-P2A4-S2GUARD1`, `E2-P2A4-REVIEW1`, `E2-P2A4-DETECT1`, `E2-P2A4-COMMIT1` |
| P2-A5 | `E2-P2A5-IMPACT1`, `E2-P2A5-SRC1`, `E2-P2A5-BUILD1`, `E2-P2A5-TEST1`, `E2-P2A5-S3GUARD1`, `E2-P2A5-REVIEW1`, `E2-P2A5-DETECT1`, `E2-P2A5-COMMIT1` |
| P2-A6 | `E2-P2A6-IMPACT1`, `E2-P2A6-SRC1`, `E2-P2A6-BUILD1`, `E2-P2A6-TEST1`, `E2-P2A6-S4GUARD1`, `E2-P2A6-REVIEW1`, `E2-P2A6-DETECT1`, `E2-P2A6-COMMIT1` |
| P2-A7 | `E2-P2A7-IMPACT1`, `E2-P2A7-SRC1`, `E2-P2A7-BUILD1`, `E2-P2A7-TEST1`, `E2-P2A7-S5GUARD1`, `E2-P2A7-REVIEW1`, `E2-P2A7-DETECT1`, `E2-P2A7-COMMIT1` |
| P2-A8 | `E2-P2A8-IMPACT1`, `E2-P2A8-SRC1`, `E2-P2A8-BUILD1`, `E2-P2A8-PLAY1`, `E2-P2A8-S6GUARD1`, `E2-P2A8-REVIEW1`, `E2-P2A8-DETECT1`, `E2-P2A8-COMMIT1` |
| P2-A9 | `E2-P2A9-IMPACT1`, `E2-P2A9-SRC1`, `E2-P2A9-BUILD1`, `E2-P2A9-TEST1`, `E2-P2A9-S7GUARD1`, `E2-P2A9-REVIEW1`, `E2-P2A9-DETECT1`, `E2-P2A9-COMMIT1` |
| P2-A10 | `E2-P2A10-IMPACT1`, `E2-P2A10-SRC1`, `E2-P2A10-BUILD1`, `E2-P2A10-TEST1`, `E2-P2A10-S8GUARD1`, `E2-P2A10-REVIEW1`, `E2-P2A10-DETECT1`, `E2-P2A10-COMMIT1` |
| P2-A11 | `E2-P2A11-IMPACT1`, `E2-P2A11-SRC1`, `E2-P2A11-BUILD1`, `E2-P2A11-TEST1`, `E2-P2A11-S9GUARD1`, `E2-P2A11-REVIEW1`, `E2-P2A11-DETECT1`, `E2-P2A11-COMMIT1` |
| P2-A12 | `E2-P2A12-IMPACT1`, `E2-P2A12-SRC1`, `E2-P2A12-BUILD1`, `E2-P2A12-TEST1`, `E2-P2A12-S10GUARD1`, `E2-P2A12-REVIEW1`, `E2-P2A12-DETECT1`, `E2-P2A12-COMMIT1` |
| P2-A13 | `E2-P2A13-IMPACT1`, `E2-P2A13-SRC1`, `E2-P2A13-BUILD1`, `E2-P2A13-TEST1`, `E2-P2A13-S10GUARD1`, `E2-P2A13-REVIEW1`, `E2-P2A13-DETECT1`, `E2-P2A13-COMMIT1` |
| P2-A14 | `E2-P2A14-IMPACT1`, `E2-P2A14-SRC1`, `E2-P2A14-BUILD1`, `E2-P2A14-TEST1`, `E2-P2A14-S11GUARD1`, `E2-P2A14-REVIEW1`, `E2-P2A14-DETECT1`, `E2-P2A14-COMMIT1` |
| P2-A15 | `E2-P2A15-IMPACT1`, `E2-P2A15-SRC1`, `E2-P2A15-BUILD1`, `E2-P2A15-TEST1`, `E2-P2A15-S11GUARD1`, `E2-P2A15-REVIEW1`, `E2-P2A15-DETECT1`, `E2-P2A15-COMMIT1` |
| P2-B | `E2-P2B-IMPACT1`, `E2-P2B-SRC1`, `E2-P2B-BUILD1`, `E2-P2B-TEST1`, `E2-P2B-REVIEW1`, `E2-P2B-DETECT1`, `E2-P2B-COMMIT1` |
| P2-B1 | `E2-P2B1-IMPACT1`, `E2-P2B1-SRC1`, `E2-P2B1-BUILD1`, `E2-P2B1-TEST1`, `E2-P2B1-REVIEW1`, `E2-P2B1-DETECT1`, `E2-P2B1-COMMIT1` |
| P2-B2 | `E2-P2B2-IMPACT1`, `E2-P2B2-SRC1`, `E2-P2B2-BUILD1`, `E2-P2B2-TEST1`, `E2-P2B2-REVIEW1`, `E2-P2B2-DETECT1`, `E2-P2B2-COMMIT1` |
| P2-B3 | `E2-P2B3-IMPACT1`, `E2-P2B3-SRC1`, `E2-P2B3-BUILD1`, `E2-P2B3-TEST1`, `E2-P2B3-REVIEW1`, `E2-P2B3-DETECT1`, `E2-P2B3-COMMIT1` |
| P2-B4 | `E2-P2B4-IMPACT1`, `E2-P2B4-SRC1`, `E2-P2B4-BUILD1`, `E2-P2B4-TEST1`, `E2-P2B4-REVIEW1`, `E2-P2B4-DETECT1`, `E2-P2B4-COMMIT1` |
| P2-C | `E2-P2C-IMPACT1`, `E2-P2C-SRC1`, `E2-P2C-BUILD1`, `E2-P2C-TEST1`, `E2-P2C-REVIEW1`, `E2-P2C-DETECT1`, `E2-P2C-COMMIT1` |
| P2-C1 | `E2-P2C1-IMPACT1`, `E2-P2C1-SRC1`, `E2-P2C1-BUILD1`, `E2-P2C1-TEST1`, `E2-P2C1-REVIEW1`, `E2-P2C1-DETECT1`, `E2-P2C1-COMMIT1` |
| P2-C2 | `E2-P2C2-IMPACT1`, `E2-P2C2-SRC1`, `E2-P2C2-BUILD1`, `E2-P2C2-TEST1`, `E2-P2C2-REVIEW1`, `E2-P2C2-DETECT1`, `E2-P2C2-COMMIT1` |
| P2-C3 | `E2-P2C3-IMPACT1`, `E2-P2C3-SRC1`, `E2-P2C3-BUILD1`, `E2-P2C3-TEST1`, `E2-P2C3-REVIEW1`, `E2-P2C3-DETECT1`, `E2-P2C3-COMMIT1` |
| P2-C4 | `E2-P2C4-IMPACT1`, `E2-P2C4-SRC1`, `E2-P2C4-BUILD1`, `E2-P2C4-TEST1`, `E2-P2C4-REVIEW1`, `E2-P2C4-DETECT1`, `E2-P2C4-COMMIT1` |
| P2-C5 | `E2-P2C5-IMPACT1`, `E2-P2C5-SRC1`, `E2-P2C5-BUILD1`, `E2-P2C5-TEST1`, `E2-P2C5-REVIEW1`, `E2-P2C5-DETECT1`, `E2-P2C5-COMMIT1` |
| P2-C6 | `E2-P2C6-IMPACT1`, `E2-P2C6-SRC1`, `E2-P2C6-BUILD1`, `E2-P2C6-TEST1`, `E2-P2C6-REVIEW1`, `E2-P2C6-DETECT1`, `E2-P2C6-COMMIT1` |
| P2-D | `E2-P2D-IMPACT1`, `E2-P2D-SRC1`, `E2-P2D-BUILD1`, `E2-P2D-TEST1`, `E2-P2D-REVIEW1`, `E2-P2D-DETECT1`, `E2-P2D-COMMIT1` |
| P2-D1 | `E2-P2D1-IMPACT1`, `E2-P2D1-SRC1`, `E2-P2D1-BUILD1`, `E2-P2D1-TEST1`, `E2-P2D1-REVIEW1`, `E2-P2D1-DETECT1`, `E2-P2D1-COMMIT1` |
| P2-D2 | `E2-P2D2-IMPACT1`, `E2-P2D2-SRC1`, `E2-P2D2-BUILD1`, `E2-P2D2-TEST1`, `E2-P2D2-REVIEW1`, `E2-P2D2-DETECT1`, `E2-P2D2-COMMIT1` |
| P2-E | `E2-P2E-IMPACT1`, `E2-P2E-SRC1`, `E2-P2E-BUILD1`, `E2-P2E-TEST1`, `E2-P2E-REVIEW1`, `E2-P2E-DETECT1`, `E2-P2E-COMMIT1` |
| P2-E1 | `E2-P2E1-IMPACT1`, `E2-P2E1-SRC1`, `E2-P2E1-BUILD1`, `E2-P2E1-PLAY1`, `E2-P2E1-REVIEW1`, `E2-P2E1-DETECT1`, `E2-P2E1-COMMIT1` |
| P2-E2 | `E2-P2E2-BUILD1`, `E2-P2E2-S0BASE1`, `E2-P2E2-S1BASE1`, `E2-P2E2-S2BASE1`, `E2-P2E2-S3BASE1`, `E2-P2E2-S4BASE1`, `E2-P2E2-S5BASE1`, `E2-P2E2-S6BASE1`, `E2-P2E2-S7BASE1`, `E2-P2E2-S8BASE1`, `E2-P2E2-S9BASE1`, `E2-P2E2-S10BASE1`, `E2-P2E2-S11BASE1`, `E2-P2E2-MATRIX1`, `E2-P2E2-PLAY1`, `E2-P2E2-REVIEW1`, `E2-P2E2-DETECT1`, `E2-P2E2-COMMIT1` |
| P2-F | `E2-P2F-IMPACT1`, `E2-P2F-SRC1`, `E2-P2F-BUILD1`, `E2-P2F-TEST1`, `E2-P2F-REVIEW1`, `E2-P2F-DETECT1`, `E2-P2F-COMMIT1` |
| P2-F1 | `E2-P2F1-IMPACT1`, `E2-P2F1-SRC1`, `E2-P2F1-BUILD1`, `E2-P2F1-TEST1`, `E2-P2F1-REVIEW1`, `E2-P2F1-DETECT1`, `E2-P2F1-COMMIT1` |
| P2-F2 | `E2-P2F2-IMPACT1`, `E2-P2F2-SRC1`, `E2-P2F2-BUILD1`, `E2-P2F2-TEST1`, `E2-P2F2-REVIEW1`, `E2-P2F2-DETECT1`, `E2-P2F2-COMMIT1` |
| P2-F3 | `E2-P2F3-IMPACT1`, `E2-P2F3-SRC1`, `E2-P2F3-BUILD1`, `E2-P2F3-TEST1`, `E2-P2F3-REVIEW1`, `E2-P2F3-DETECT1`, `E2-P2F3-COMMIT1` |
| P2-F4 | `E2-P2F4-IMPACT1`, `E2-P2F4-SRC1`, `E2-P2F4-BUILD1`, `E2-P2F4-TEST1`, `E2-P2F4-REVIEW1`, `E2-P2F4-DETECT1`, `E2-P2F4-COMMIT1` |
| P2-F5 | `E2-P2F5-IMPACT1`, `E2-P2F5-SRC1`, `E2-P2F5-BUILD1`, `E2-P2F5-TEST1`, `E2-P2F5-REVIEW1`, `E2-P2F5-DETECT1`, `E2-P2F5-COMMIT1` |
| P2-F6 | `E2-P2F6-BUILD1`, `E2-P2F6-FAULT1`, `E2-P2F6-RECOVERY1`, `E2-P2F6-REVIEW1`, `E2-P2F6-DETECT1`, `E2-P2F6-COMMIT1` |
| P2-G | `E2-P2G-PREBASE1`, `E2-P2G-CANDIDATE1`, `E2-P2G-IMPACT1`, `E2-P2G-SRC1`, `E2-P2G-CUTOVER1`, `E2-P2G-BUILD1`, `E2-P2G-RUNTIME1`, `E2-P2G-PLAY1`, `E2-P2G-ROLLBACK1`, `E2-P2G-REVIEW1`, `E2-P2G-DETECT1`, `E2-P2G-COMMIT1` |
| P3-A | `E3-P3A-IMPACT1`, `E3-P3A-SRC1`, `E3-P3A-BUILD1`, `E3-P3A-TEST1`, `E3-P3A-REVIEW1`, `E3-P3A-DETECT1`, `E3-P3A-COMMIT1` |
| P3-B | `E3-P3B-IMPACT1`, `E3-P3B-SRC1`, `E3-P3B-BUILD1`, `E3-P3B-TEST1`, `E3-P3B-REVIEW1`, `E3-P3B-DETECT1`, `E3-P3B-COMMIT1` |
| P3-B1 | `E3-P3B1-IMPACT1`, `E3-P3B1-SRC1`, `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, `E3-P3B1-REVIEW1`, `E3-P3B1-DETECT1`, `E3-P3B1-COMMIT1` |
| P3-B2 | `E3-P3B2-IMPACT1`, `E3-P3B2-SRC1`, `E3-P3B2-BUILD1`, `E3-P3B2-TEST1`, `E3-P3B2-REVIEW1`, `E3-P3B2-DETECT1`, `E3-P3B2-COMMIT1` |
| P3-B2A | `E3-P3B2A-IMPACT1`, `E3-P3B2A-SRC1`, `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`, `E3-P3B2A-REVIEW1`, `E3-P3B2A-DETECT1`, `E3-P3B2A-COMMIT1` |
| P3-C | `E3-P3C-IMPACT1`, `E3-P3C-SRC1`, `E3-P3C-BUILD1`, `E3-P3C-TEST1`, `E3-P3C-REVIEW1`, `E3-P3C-DETECT1`, `E3-P3C-COMMIT1` |
| P3-C1 | `E3-P3C1-IMPACT1`, `E3-P3C1-SRC1`, `E3-P3C1-BUILD1`, `E3-P3C1-TEST1`, `E3-P3C1-REVIEW1`, `E3-P3C1-DETECT1`, `E3-P3C1-COMMIT1` |
| P3-C1A | `E3-P3C1A-IMPACT1`, `E3-P3C1A-SRC1`, `E3-P3C1A-BUILD1`, `E3-P3C1A-TEST1`, `E3-P3C1A-REVIEW1`, `E3-P3C1A-DETECT1`, `E3-P3C1A-COMMIT1` |
| P3-C1B | `E3-P3C1B-IMPACT1`, `E3-P3C1B-SRC1`, `E3-P3C1B-BUILD1`, `E3-P3C1B-TEST1`, `E3-P3C1B-REVIEW1`, `E3-P3C1B-DETECT1`, `E3-P3C1B-COMMIT1` |
| P3-C1C | `E3-P3C1C-IMPACT1`, `E3-P3C1C-SRC1`, `E3-P3C1C-BUILD1`, `E3-P3C1C-TEST1`, `E3-P3C1C-REVIEW1`, `E3-P3C1C-DETECT1`, `E3-P3C1C-COMMIT1` |
| P3-C1D | `E3-P3C1D-IMPACT1`, `E3-P3C1D-SRC1`, `E3-P3C1D-BUILD1`, `E3-P3C1D-TEST1`, `E3-P3C1D-REVIEW1`, `E3-P3C1D-DETECT1`, `E3-P3C1D-COMMIT1` |
| P3-C1E | `E3-P3C1E-IMPACT1`, `E3-P3C1E-SRC1`, `E3-P3C1E-BUILD1`, `E3-P3C1E-TEST1`, `E3-P3C1E-REVIEW1`, `E3-P3C1E-DETECT1`, `E3-P3C1E-COMMIT1` |
| P3-C1F | `E3-P3C1F-IMPACT1`, `E3-P3C1F-SRC1`, `E3-P3C1F-BUILD1`, `E3-P3C1F-PLAY1`, `E3-P3C1F-REVIEW1`, `E3-P3C1F-DETECT1`, `E3-P3C1F-COMMIT1` |
| P3-C1G | `E3-P3C1G-IMPACT1`, `E3-P3C1G-SRC1`, `E3-P3C1G-BUILD1`, `E3-P3C1G-TEST1`, `E3-P3C1G-REVIEW1`, `E3-P3C1G-DETECT1`, `E3-P3C1G-COMMIT1` |
| P3-C1H | `E3-P3C1H-IMPACT1`, `E3-P3C1H-SRC1`, `E3-P3C1H-BUILD1`, `E3-P3C1H-TEST1`, `E3-P3C1H-REVIEW1`, `E3-P3C1H-DETECT1`, `E3-P3C1H-COMMIT1` |
| P3-C1I | `E3-P3C1I-IMPACT1`, `E3-P3C1I-SRC1`, `E3-P3C1I-BUILD1`, `E3-P3C1I-TEST1`, `E3-P3C1I-REVIEW1`, `E3-P3C1I-DETECT1`, `E3-P3C1I-COMMIT1` |
| P3-C2 | `E3-P3C2-ORACLE1`, `E3-P3C2-TARGET1`, `E3-P3C2-BOUNDARY1`, `E3-P3C2-REVIEW1`, `E3-P3C2-DETECT1`, `E3-P3C2-COMMIT1` |
| P4-A | `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-BUILD1`, `E4-P4A-TEST1`, `E4-P4A-REVIEW1`, `E4-P4A-DETECT1`, `E4-P4A-COMMIT1` |
| P4-B | `E4-P4B-IMPACT1`, `E4-P4B-SRC1`, `E4-P4B-BUILD1`, `E4-P4B-TEST1`, `E4-P4B-REVIEW1`, `E4-P4B-DETECT1`, `E4-P4B-COMMIT1` |
| P4-B1 | `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-BUILD1`, `E4-P4B1-TEST1`, `E4-P4B1-REVIEW1`, `E4-P4B1-DETECT1`, `E4-P4B1-COMMIT1` |
| P4-C | `E4-P4C-IMPACT1`, `E4-P4C-SRC1`, `E4-P4C-BUILD1`, `E4-P4C-TEST1`, `E4-P4C-REVIEW1`, `E4-P4C-DETECT1`, `E4-P4C-COMMIT1` |
| P4-C1 | `E4-P4C1-IMPACT1`, `E4-P4C1-SRC1`, `E4-P4C1-BUILD1`, `E4-P4C1-TEST1`, `E4-P4C1-REVIEW1`, `E4-P4C1-DETECT1`, `E4-P4C1-COMMIT1` |
| P4-C1A | `E4-P4C1A-IMPACT1`, `E4-P4C1A-SRC1`, `E4-P4C1A-BUILD1`, `E4-P4C1A-TEST1`, `E4-P4C1A-REVIEW1`, `E4-P4C1A-DETECT1`, `E4-P4C1A-COMMIT1` |
| P4-C1B | `E4-P4C1B-IMPACT1`, `E4-P4C1B-SRC1`, `E4-P4C1B-BUILD1`, `E4-P4C1B-TEST1`, `E4-P4C1B-REVIEW1`, `E4-P4C1B-DETECT1`, `E4-P4C1B-COMMIT1` |
| P4-C1C | `E4-P4C1C-IMPACT1`, `E4-P4C1C-SRC1`, `E4-P4C1C-BUILD1`, `E4-P4C1C-TEST1`, `E4-P4C1C-REVIEW1`, `E4-P4C1C-DETECT1`, `E4-P4C1C-COMMIT1` |
| P4-C1D | `E4-P4C1D-IMPACT1`, `E4-P4C1D-SRC1`, `E4-P4C1D-BUILD1`, `E4-P4C1D-TEST1`, `E4-P4C1D-REVIEW1`, `E4-P4C1D-DETECT1`, `E4-P4C1D-COMMIT1` |
| P4-C1E | `E4-P4C1E-IMPACT1`, `E4-P4C1E-SRC1`, `E4-P4C1E-BUILD1`, `E4-P4C1E-TEST1`, `E4-P4C1E-REVIEW1`, `E4-P4C1E-DETECT1`, `E4-P4C1E-COMMIT1` |
| P4-C1F | `E4-P4C1F-IMPACT1`, `E4-P4C1F-SRC1`, `E4-P4C1F-BUILD1`, `E4-P4C1F-PLAY1`, `E4-P4C1F-REVIEW1`, `E4-P4C1F-DETECT1`, `E4-P4C1F-COMMIT1` |
| P4-C1G | `E4-P4C1G-IMPACT1`, `E4-P4C1G-SRC1`, `E4-P4C1G-BUILD1`, `E4-P4C1G-TEST1`, `E4-P4C1G-REVIEW1`, `E4-P4C1G-DETECT1`, `E4-P4C1G-COMMIT1` |
| P4-C1H | `E4-P4C1H-IMPACT1`, `E4-P4C1H-SRC1`, `E4-P4C1H-BUILD1`, `E4-P4C1H-TEST1`, `E4-P4C1H-REVIEW1`, `E4-P4C1H-DETECT1`, `E4-P4C1H-COMMIT1` |
| P4-C1I | `E4-P4C1I-IMPACT1`, `E4-P4C1I-SRC1`, `E4-P4C1I-BUILD1`, `E4-P4C1I-TEST1`, `E4-P4C1I-REVIEW1`, `E4-P4C1I-DETECT1`, `E4-P4C1I-COMMIT1` |
| P4-C2 | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW1`, `E4-P4C2-DETECT1`, `E4-P4C2-COMMIT1` |
| P5-A | `E5-P5A-IMPACT1`, `E5-P5A-SRC1`, `E5-P5A-BUILD1`, `E5-P5A-ORACLE1`, `E5-P5A-REVIEW1`, `E5-P5A-DETECT1`, `E5-P5A-COMMIT1` |
| P5-B | `E5-P5B-IMPACT1`, `E5-P5B-SRC1`, `E5-P5B-BUILD1`, `E5-P5B-TEST1`, `E5-P5B-BENCH1`, `E5-P5B-REVIEW1`, `E5-P5B-DETECT1`, `E5-P5B-COMMIT1` |
| P5-C | `E5-P5C-IMPACT1`, `E5-P5C-SRC1`, `E5-P5C-BUILD1`, `E5-P5C-TEST1`, `E5-P5C-VECTOR1`, `E5-P5C-BENCH1`, `E5-P5C-REVIEW1`, `E5-P5C-DETECT1`, `E5-P5C-COMMIT1` |
| P5-D | `E5-P5D-IMPACT1`, `E5-P5D-SRC1`, `E5-P5D-BUILD1`, `E5-P5D-TARGET1`, `E5-P5D-PARITY1`, `E5-P5D-BOUNDARY1`, `E5-P5D-ORACLE1`, `E5-P5D-REVIEW1`, `E5-P5D-DETECT1`, `E5-P5D-COMMIT1` |
| P6-A | `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-BUILD1`, `E6-P6A-TEST1`, `E6-P6A-REVIEW1`, `E6-P6A-DETECT1`, `E6-P6A-COMMIT1` |
| P6-B | `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-REVIEW1`, `E6-P6B-DETECT1`, `E6-P6B-COMMIT1` |
| P6-C1 | `E6-P6C1-IMPACT1`, `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-ORACLE1`, `E6-P6C1-REVIEW1`, `E6-P6C1-DETECT1`, `E6-P6C1-COMMIT1` |
| P6-C2 | `E6-P6C2-IMPACT1`, `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-BENCH1`, `E6-P6C2-REVIEW1`, `E6-P6C2-DETECT1`, `E6-P6C2-COMMIT1` |
| P6-C3 | `E6-P6C3-IMPACT1`, `E6-P6C3-SRC1`, `E6-P6C3-BUILD1`, `E6-P6C3-TEST1`, `E6-P6C3-STATUS1`, `E6-P6C3-REVIEW1`, `E6-P6C3-DETECT1`, `E6-P6C3-COMMIT1` |
| P6-D | `E6-P6D-IMPACT1`, `E6-P6D-SRC1`, `E6-P6D-BUILD1`, `E6-P6D-TARGET1`, `E6-P6D-BOUNDARY1`, `E6-P6D-ORACLE1`, `E6-P6D-REVIEW1`, `E6-P6D-DETECT1`, `E6-P6D-COMMIT1` |
| P7-A | `E7-P7A-BUILD1`, `E7-P7A-DETERMINISM1`, `E7-P7A-VERSION1`, `E7-P7A-FAULT1`, `E7-P7A-REVIEW1`, `E7-P7A-DETECT1`, `E7-P7A-COMMIT1` |
| P7-B | `E7-P7B-TARGET1`, `E7-P7B-ORACLE1`, `E7-P7B-BOUNDARY1`, `E7-P7B-REVIEW1`, `E7-P7B-DETECT1`, `E7-P7B-COMMIT1` |
| P7-C | `E7-P7C-BUILD1`, `E7-P7C-RUNTIME1`, `E7-P7C-PLAY1`, `E7-P7C-BENCH1`, `E7-P7C-NATIVEBENCH1`, `E7-P7C-FALLBACKBENCH1`, `E7-P7C-S0-PARITY1`, `E7-P7C-S1-PARITY1`, `E7-P7C-S2-PARITY1`, `E7-P7C-S3-PARITY1`, `E7-P7C-S4-PARITY1`, `E7-P7C-S5-PARITY1`, `E7-P7C-S6-PARITY1`, `E7-P7C-S7-PARITY1`, `E7-P7C-S8-PARITY1`, `E7-P7C-S9-PARITY1`, `E7-P7C-S10-PARITY1`, `E7-P7C-S11-PARITY1`, `E7-P7C-REVIEW1`, `E7-P7C-DETECT1`, `E7-P7C-COMMIT1` |
| P8-A | `E8-P8A-REVIEW1` |
| P8-B | `E8-P8B-CLEAN1`, `E8-P8B-REVIEW1` |
| P8-C | `E8-P8C-BUILD1`, `E8-P8C-RUNTIME1`, `E8-P8C-PLAY1`, `E8-P8C-REGEN1`, `E8-P8C-DETECT1`, `E8-P8C-LEDGER1`, `E8-P8C-COMMIT1` |

## Risk Notes

- Identity v2 is CRITICAL blast radius: relationship endpoints, ResolutionGap, Process, embeddings, group registry, rename, MCP, HTTP/Web, sorting, communities, and caches all depend on current IDs.
- Adding range to old IDs alone prevents some collisions but creates widespread churn after line edits and still fails logical symbol/overload semantics.
- A strict `AddNode` change without explicit producer classification can break legitimate enrichment and still allow ignored errors.
- Declaration nodes may materially increase graph size, Ladybug load time, RSS, and query cost; P1 shadow and P7 benchmarks are release gates.
- Case normalization and path canonicalization can drift between Windows and Linux; determinism tests must cover both packaged platform policies where supported.
- TypeScript declaration merging cannot be inferred safely from name equality or Tree-sitter shape alone; ambiguous/unverified facts must remain separate.
- Star exports and cycles can cause combinatorial work; bounded memoization must return `budget_exceeded`, never silently truncate or choose a candidate.
- Embedded declaration data has provenance, license, version, integrity, binary-size, and regeneration risks; these must be explicit before catalog acceptance.
- Generation atomicity cannot be achieved by renaming graph.json alone; all readers must bind to the active manifest and one generation.
- Old binaries ignore unknown JSON fields. The cutover must prevent unsupported binaries/clients from reading v2 rather than assuming metadata alone protects them.
- The target has pre-existing changes. Every target run must capture and preserve the boundary; do not reset, clean, copy, or repair target state.
- The missing approved architecture SPEC means this draft is not implementation authority until the owner accepts its decision table or an equivalent approved contract is recorded.
