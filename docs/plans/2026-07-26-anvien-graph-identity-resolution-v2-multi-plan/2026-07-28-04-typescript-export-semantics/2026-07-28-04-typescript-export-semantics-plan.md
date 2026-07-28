# Anvien First-Class TypeScript Export Semantics Plan

## Metadata

- Date: `2026-07-28`
- Status: `draft / P0 complete / implementation not yet authorized`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md`
- Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Authoring plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Predecessor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
- Planned successor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`

## Goal

represent module exports explicitly and consistently, separate from access visibility.

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
- When this child ends, update the next child plan's actual-status and latest evidence before handoff.

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

- Every surface `S0`–`S11` exposes the same canonical node/relationship records for the active generation; native Cypher and fallback Cypher are measured separately, and field-level differences plus orphan references are both zero.
- All 21 bounded direct exports are correct (`21/21`) in ScopeIR and the applicable `S0`–`S11` projections; access visibility remains separate.
- The target's known scanner baseline remains quarantined at the same exact eight omissions unless a separate scanner plan is authorized; this plan creates `0` new File-node omissions.
- No source, fixture, report, probe, or temporary investigation artifact is written into or copied from `E:\cheapapp.org`; operational graph/index output remains under `E:\cheapapp.org\.anvien`.
- Every touched/new source and test file passes the one-file/one-responsibility review; `0` new catch-all files and `0` unrelated responsibility additions.
- Every slice has a separate commit after full build, boundary validation, ledger update, Supervisor PASS, and detect-changes.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.

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

- [ ] Pn-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, changed files, generated output, and validation results before closure.
  - Work Steps:
    1. Call the supervisor skill to review the full completed plan work.
    2. If supervisor fails the work, return to the responsible implementation workflow/skill for the failed scope only.
    3. Re-run supervisor review after the fix.
    4. Repeat until supervisor passes or records a blocker.
  - Implementation Gate: all planned implementation phases must be completed or explicitly blocked before this review.
  - Acceptance: supervisor review passes, or the plan records a blocker with evidence and no closure is performed.
- [ ] Pn-B: Remove dead work created during this plan.
  - Goal: ensure the final diff contains only artifacts that still serve the accepted plan.
  - Work Steps:
    1. Review files, sections, generated output, tests, temp files, and plan artifacts created or modified during this plan.
    2. Remove or rewrite any artifact made obsolete by actual-status findings, user corrections, failed approaches, or phase status updates.
    3. Verify no rejected approach, stale placeholder, unused generated output, or dead helper artifact remains in the final diff.
    4. Call supervisor to review the dead-work cleanup.
    5. If supervisor fails the cleanup, return to the responsible implementation workflow/skill for the failed cleanup scope only, then re-run supervisor review.
  - Implementation Gate: only remove artifacts created by this plan unless the user explicitly approves broader cleanup.
  - Acceptance: final `git diff/status` contains no dead plan-created artifacts, supervisor passes the cleanup, and evidence records what was removed or preserved.
- [ ] Pn-C: Close the plan.
  - Goal: finish validation, evidence, benchmark, detect-changes, commit, and final status.
  - Work Steps:
    1. Run the required final validation for the accepted scope, including full build before final runtime validation. For app/runtime scopes, full build must include Docker image/container build.
    2. Start the real built Docker/container runtime for app/runtime validation. If Docker cannot be built or started, record the blocker and do not substitute a host dev server.
    3. Validate public runtime or UI-facing changes with browser or Playwright evidence against the real built Docker/container runtime. Playwright evidence must include Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
    4. Regenerate generated outputs if source-of-truth changes require it.
    5. Run Anvien detect-changes before commit when implementation work was performed.
    6. Record final validation, detect-changes, benchmark, and commit evidence.
    7. Commit the completed scope and verify the worktree state.
  - Implementation Gate: Pn-A and Pn-B must pass or record blockers.
  - Acceptance: final evidence is recorded, required commits exist, and the worktree state is known.

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
