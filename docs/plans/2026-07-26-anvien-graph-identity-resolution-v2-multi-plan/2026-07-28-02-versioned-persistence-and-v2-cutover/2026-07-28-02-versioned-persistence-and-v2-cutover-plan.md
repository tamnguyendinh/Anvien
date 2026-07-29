# Anvien Versioned Persistence, Opaque Consumers, Atomic Generation, And V2 Cutover Plan

## Metadata

- Date: `2026-07-28`
- Status: `draft / P0 complete / implementation not yet authorized`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-actual-status.md`
- Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Authoring plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Predecessor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
- Planned successor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`

## Goal

migrate each persistence, reader, consumer, and publication owner independently, then activate identity v2 without mixed versions or generations.

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
- The target's known scanner baseline remains quarantined at the same exact eight omissions unless a separate scanner plan is authorized; this plan creates `0` new File-node omissions.
- No source, fixture, report, probe, or temporary investigation artifact is written into or copied from `E:\cheapapp.org`; operational graph/index output remains under `E:\cheapapp.org\.anvien`.
- Before identity cutover, at least five measured v1 runs on the same commit-bound corpus/config/build/machine/cache policy establish analyze median, Ladybug-load median, native-query p95, fallback-query p95, graph size, and peak RSS. Final medians and p95 values regress by no more than `10%`, peak RSS by no more than `15%`, unless the owner accepts a measured exception with both baseline and final values before cutover.
- Every touched/new source and test file passes the one-file/one-responsibility review; `0` new catch-all files and `0` unrelated responsibility additions.
- Every slice has a separate commit after full build, boundary validation, ledger update, Supervisor PASS, and detect-changes.

## Semantic Remediation Overlay (mandatory)

The copied P2 phase block above is historical provenance and remains unchanged. This overlay is the stronger execution contract; it controls whenever a copied gate is weaker or contradictory.

- Local closure is a checkpoint, not campaign completion. Pending P7 rows do not block Child 02 local closure; campaign/release closure remains a Child 07 responsibility.
- Before `P2-A` opens, the exact qualified predecessor gate is `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1`; a generic “Child 01 handoff” phrase is non-accepting. `P2-A` must consume that record inspect-only and then publish its own qualified successor records.
- Every implementation slice Acceptance is conjunctive with the complete matching evidence-ledger row (`IMPACT`, `SRC`, `BUILD`, behavior `TEST`/oracle, `REVIEW`, `DETECT`, and `COMMIT`). A `REVIEW1` record alone cannot close a slice.
- Before each job opens, record an exact ownership table with `File`, one unique responsibility, allowed links, and prohibited contents for every production, test, generated, and fixture file. Wildcards, `TBD`, catch-all owners, or mixed responsibilities fail the implementation gate.
- `Pn-C` must create qualified `E2-PNC-NEXTSTATUS1` and `E2-PNC-HANDOFF1`; refresh the next child actual-status, refresh log, next action, and work steps from the latest accepted handoff before Child 03 can open, then bind the handoff to this child’s own `NEXTSTATUS1`, never Child 03’s future record.
- The qualified names are `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-NEXTSTATUS1` and `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-PNC-HANDOFF1`; Child 03 must consume these namespaced records.
- The manifest/handshake is not complete without `graphSchemaVersion`, `identitySchemaVersion`, `scopeIRSchemaVersion`, `graphGeneration`, `analyzerVersion`, `columnEncoding`/`positionEncoding`, `sourceFingerprint`, and `configFingerprint` in both persisted manifest and request/response validation. Missing metadata fails closed before any body opens.
- Corrected wire shape: request `{readerProtocolVersion, readerBuild, supportedGraphSchemaVersions[], supportedIdentitySchemaVersions[], supportedScopeIrVersions[], supportedAnalyzerVersions[], supportedColumnEncodings[], supportedPositionEncodings[], sourceFingerprint, configFingerprint}`; manifest `{protocolVersion, minReaderProtocol, minReaderBuild, graphSchemaVersion, identitySchemaVersion, scopeIRSchemaVersion, generation, graphGeneration, analyzerVersion, columnEncoding, positionEncoding, sourceFingerprint, configFingerprint, configHash, catalogHash}`; mismatch envelope `{code:"INDEX_VERSION_MISMATCH", expected:{...}, actual:{...}, retryable:false}`. Readers validate this before opening any body, cache, stream, registry, group, or embedding.
- Child 02 is the sole owner of the reader matrix and generation manifest; later children consume its qualified handoff inspect-only. The local closure gate must not wait for P7 performance rows.

### Manifest/handshake correction contract (mandatory)

The copied P2-A field lists and wire-shape paragraphs above are historical provenance. The following contract supersedes them wherever they disagree and is the only contract that may be implemented or accepted.

| Contract field | Required type/meaning | Persisted manifest | Reader request/response check |
|---|---|---|---|
| `graphSchemaVersion` | non-empty opaque schema version | required | supported-version intersection |
| `identitySchemaVersion` | non-empty identity tuple version | required | supported-version intersection |
| `scopeIRSchemaVersion` | exact ScopeIR schema version (spelling is case-sensitive) | required | supported-version intersection |
| `graphGeneration` | immutable generation identifier for the graph body | required and equal to the opened body generation | equality before body open |
| `analyzerVersion` | analyzer build/version plus reproducible build identifier | required | supported analyzer set and exact expected build policy |
| `columnEncoding` | enum for persisted tabular columns | required | supported column-encoding set |
| `positionEncoding` | enum for source ranges/selection ranges | required | supported position-encoding set |
| `sourceFingerprint` | SHA-256 of canonical eligible-source manifest | required | exact equality; mismatch is fail-closed |
| `configFingerprint` | SHA-256 of canonical analyzer/config/profile inputs | required | exact equality; mismatch is fail-closed |

`generation` and `graphGeneration` are both retained: the former identifies the publication epoch, the latter binds graph records to that epoch. `configHash` and `catalogHash` remain supplementary provenance and cannot replace either fingerprint. The validator must reject absent, empty, type-invalid, or mismatched fields before opening JSON, Ladybug, Cypher, cache, registry, group, stream, or embedding bodies.

Required correction evidence is split, not aggregate: `E2-P2A-MANIFEST1` (persisted manifest round-trip), `E2-P2A-HANDSHAKE1` (request/response and mismatch matrix), and `E2-P2A-METADATA1` (the nine semantic-correction fields plus the complete 15-field manifest/10-field request inventory, including presence, type, spelling, and fingerprint checks). “Nine-field” names the semantic subset; it is not a reduced wire contract. These IDs stay pending until their owning slice produces the evidence.

The complete 42-job manifest below is the pre-implementation closed set. Any evidence-backed owner-path change must update the exact affected row and `E2-PNC-OWNERSHIP1` before editing; a wildcard, `TBD`, or post-hoc file choice fails the job gate.

For every P2-A2–P2-A15 reader adapter and every P2-B–P2-G consumer slice, the copied Acceptance text is conjunctive with the complete matching ledger row: exact `IMPACT1`, `SRC1`, `BUILD1`, behavior `TEST1`/`PLAY1`, `REVIEW1`, `DETECT1`, and `COMMIT1` (or the slice-specific equivalent explicitly listed in the ledger). A parent matrix row or review-only record cannot close an individual reader guard.

### Complete P2 job ownership manifest (mandatory, closed before implementation)

The following is the complete 42-job manifest, not a future placeholder. Each implementation/validation row has one semantic owner, one behavior-test owner, one generated evidence file, and one independent fixture file. P2-A1 is the sole documentation-only exception: it owns only the exact reader-matrix document and records reasoned `N/A` cells for code/test/generated/fixture ownership, matching its preserve-only source boundary and documentation-only commit. P2-E2 and P2-F6 own reusable validation harnesses only; those exact listed validation files are overlay extensions to their editable validation scope and may not contain implementation repair. A path absent at the current HEAD is a planned new path and must be created only by the named job; a path present at HEAD remains subject to its current one-responsibility owner and impact gate. For P2-A2–P2-A15, the exact existing adapter files remain enumerated by P2-A1 reader-matrix row IDs and are incorporated into the owning row without a wildcard; the listed production file owns only the shared guard. Any newly discovered adapter must first receive an exact reader-matrix row and an exact ownership row before editing.

| Job | Production owner file | Test owner file | Generated evidence file | Fixture file | Unique responsibility | Allowed links | Prohibited contents |
|---|---|---|---|---|---|---|---|
| P2-A | `internal/repo/index_version.go` | `internal/repo/index_version_test.go` | `internal/testdata/generated/p2-a-manifest.json` | `internal/testdata/p2/a-mismatch.json` | manifest/handshake contract | repo metadata types | reader adapters and staging |
| P2-A1 | `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/index-reader-matrix.md` | `N/A` — documentation/source audit only | `N/A` — results are recorded in the child evidence ledger | `N/A` — source anchors are inspected in place | sole reader inventory and ownership map | exact source reader anchors and later guard-owner rows | production/test code or runtime guard implementation |
| P2-A2 | `internal/repo/graph_loader_guard.go` | `internal/repo/graph_loader_guard_test.go` | `internal/testdata/generated/p2-a2-s0.json` | `internal/testdata/p2/a2-mismatch.json` | S0 graph/metadata guard | index_version contract | other reader surfaces |
| P2-A3 | `internal/lbugload/reader_guard.go` | `internal/lbugload/reader_guard_test.go` | `internal/testdata/generated/p2-a3-s1.json` | `internal/testdata/p2/a3-mismatch.json` | S1 native Ladybug guard | index_version contract | Graph JSON or Cypher adapters |
| P2-A4 | `internal/lbugruntime/query_guard.go` | `internal/lbugruntime/query_guard_test.go` | `internal/testdata/generated/p2-a4-s2.json` | `internal/testdata/p2/a4-mismatch.json` | S2 fallback Cypher guard | index_version contract | CLI/MCP routing |
| P2-A5 | `internal/cli/index_guard.go` | `internal/cli/index_guard_test.go` | `internal/testdata/generated/p2-a5-s3.json` | `internal/testdata/p2/a5-mismatch.json` | S3 CLI guard | index_version contract | MCP/HTTP adapters |
| P2-A6 | `internal/mcp/index_guard.go` | `internal/mcp/index_guard_test.go` | `internal/testdata/generated/p2-a6-s4.json` | `internal/testdata/p2/a6-mismatch.json` | S4 MCP guard | index_version contract | HTTP/Web adapters |
| P2-A7 | `internal/httpapi/index_guard.go` | `internal/httpapi/index_guard_test.go` | `internal/testdata/generated/p2-a7-s5.json` | `internal/testdata/p2/a7-mismatch.json` | S5 HTTP/stream guard | index_version contract | Web client rendering |
| P2-A8 | `anvien-web/src/index-guard.ts` | `anvien-web/src/index-guard.test.ts` | `anvien-web/testdata/generated/p2-a8-s6.json` | `anvien-web/testdata/p2/a8-mismatch.json` | S6 Web lifecycle guard | HTTP mismatch envelope | server-side readers |
| P2-A9 | `internal/filecontext/index_guard.go` | `internal/filecontext/index_guard_test.go` | `internal/testdata/generated/p2-a9-s7.json` | `internal/testdata/p2/a9-mismatch.json` | S7 file-context cache guard | index_version contract | resource-cache policy |
| P2-A10 | `internal/mcp/resource_cache_guard.go` | `internal/mcp/resource_cache_guard_test.go` | `internal/testdata/generated/p2-a10-s8.json` | `internal/testdata/p2/a10-mismatch.json` | S8 resource-cache guard | index_version contract | file-context implementation |
| P2-A11 | `internal/embeddings/index_guard.go` | `internal/embeddings/index_guard_test.go` | `internal/testdata/generated/p2-a11-s9.json` | `internal/testdata/p2/a11-mismatch.json` | S9 embedding reader guard | index_version contract | graph generation writer |
| P2-A12 | `internal/repo/registry_guard.go` | `internal/repo/registry_guard_test.go` | `internal/testdata/generated/p2-a12-s10-repo.json` | `internal/testdata/p2/a12-mismatch.json` | S10 repository registry guard | index_version contract | group contract publication |
| P2-A13 | `internal/group/registry_guard.go` | `internal/group/registry_guard_test.go` | `internal/testdata/generated/p2-a13-s10-group.json` | `internal/testdata/p2/a13-mismatch.json` | S10 group registry guard | index_version contract | global registry implementation |
| P2-A14 | `internal/processes/index_guard.go` | `internal/processes/index_guard_test.go` | `internal/testdata/generated/p2-a14-s11-process.json` | `internal/testdata/p2/a14-mismatch.json` | S11 process projection guard | index_version contract | community projection |
| P2-A15 | `internal/communities/index_guard.go` | `internal/communities/index_guard_test.go` | `internal/testdata/generated/p2-a15-s11-community.json` | `internal/testdata/p2/a15-mismatch.json` | S11 community projection guard | index_version contract | process projection |
| P2-B | `internal/repo/graph_v2_codec.go` | `internal/repo/graph_v2_codec_test.go` | `internal/testdata/generated/p2-b-graph.json` | `internal/testdata/p2/b-invalid.json` | Graph JSON v2 codec | canonical graph contract | reader negotiation |
| P2-B1 | `internal/lbugload/csv_v2.go` | `internal/lbugload/csv_v2_test.go` | `internal/testdata/generated/p2-b1-csv.json` | `internal/testdata/p2/b1-csv-fixture.json` | Ladybug v2 schema/CSV export | canonical graph fields | transactional load |
| P2-B2 | `internal/lbugload/transactional_v2.go` | `internal/lbugload/transactional_v2_test.go` | `internal/testdata/generated/p2-b2-load.json` | `internal/testdata/p2/b2-fault.json` | transactional Ladybug load | manifest and CSV contract | query projection |
| P2-B3 | `internal/lbugnative/query_v2.go` | `internal/lbugnative/query_v2_test.go` | `internal/testdata/generated/p2-b3-native.json` | `internal/testdata/p2/b3-query.json` | native Ladybug canonical projection | v2 loaded graph | fallback query |
| P2-B4 | `internal/lbugruntime/fallback_v2.go` | `internal/lbugruntime/fallback_v2_test.go` | `internal/testdata/generated/p2-b4-fallback.json` | `internal/testdata/p2/b4-query.json` | fallback Cypher canonical projection | v2 loaded graph | native driver |
| P2-C | `internal/cli/opaque_id_reader.go` | `internal/cli/opaque_id_reader_test.go` | `internal/testdata/generated/p2-c-cli.json` | `internal/testdata/p2/c-opaque.json` | CLI opaque-ID migration | canonical fields | ID parsing fallback |
| P2-C1 | `internal/mcp/opaque_id_reader.go` | `internal/mcp/opaque_id_reader_test.go` | `internal/testdata/generated/p2-c1-mcp.json` | `internal/testdata/p2/c1-opaque.json` | MCP opaque-ID migration | canonical fields | CLI implementation |
| P2-C2 | `internal/filecontext/canonical_projection.go` | `internal/filecontext/canonical_projection_test.go` | `internal/testdata/generated/p2-c2-context.json` | `internal/testdata/p2/c2-canonical.json` | file-context canonical projection | canonical graph fields | cache persistence |
| P2-C3 | `internal/filecontext/cache_generation.go` | `internal/filecontext/cache_generation_test.go` | `internal/testdata/generated/p2-c3-cache.json` | `internal/testdata/p2/c3-stale.json` | generation/config/catalog cache binding | manifest fingerprints | graph writer |
| P2-C4 | `internal/mcp/rename_anchor.go` | `internal/mcp/rename_anchor_test.go` | `internal/testdata/generated/p2-c4-rename.json` | `internal/testdata/p2/c4-anchor.json` | source-anchor rename | canonical source anchors | ID parsing or graph mutation |
| P2-C5 | `internal/mcp/resource_cache_record.go` | `internal/mcp/resource_cache_record_test.go` | `internal/testdata/generated/p2-c5-resource-cache.json` | `internal/testdata/p2/c5-stale.json` | generation-qualified resource cache | manifest/generation | resolver semantics |
| P2-C6 | `internal/embeddings/opaque_reference.go` | `internal/embeddings/opaque_reference_test.go` | `internal/testdata/generated/p2-c6-embedding.json` | `internal/testdata/p2/c6-opaque.json` | opaque embedding references | canonical IDs/generation | graph construction |
| P2-D | `internal/group/contract_generation.go` | `internal/group/contract_generation_test.go` | `internal/testdata/generated/p2-d-group.json` | `internal/testdata/p2/d-contract.json` | generation-qualified group contracts | registry generation | process projection |
| P2-D1 | `internal/processes/source_anchor.go` | `internal/processes/source_anchor_test.go` | `internal/testdata/generated/p2-d1-process.json` | `internal/testdata/p2/d1-anchor.json` | source-anchored process projection | canonical source anchors | group registry |
| P2-D2 | `internal/communities/source_anchor.go` | `internal/communities/source_anchor_test.go` | `internal/testdata/generated/p2-d2-community.json` | `internal/testdata/p2/d2-anchor.json` | source-anchored community projection | canonical source anchors | process projection |
| P2-E | `internal/httpapi/canonical_fields.go` | `internal/httpapi/canonical_fields_test.go` | `internal/testdata/generated/p2-e-http.json` | `internal/testdata/p2/e-http.json` | HTTP canonical/version fields | manifest and graph records | Web rendering |
| P2-E1 | `anvien-web/src/generation-negotiation.ts` | `anvien-web/src/generation-negotiation.test.ts` | `anvien-web/testdata/generated/p2-e1-web.json` | `anvien-web/testdata/p2/e1-mismatch.json` | Web generation negotiation | HTTP contract | server-side manifest |
| P2-E2 | `internal/validation/s0_s11_baseline.go` | `internal/validation/s0_s11_baseline_test.go` | `internal/testdata/generated/p2-e2-s0-s11.json` | `internal/testdata/p2/e2-baseline.json` | pre-cutover S0-S11 baseline | reader matrix | v2 cutover |
| P2-F | `internal/analyze/generation_stage.go` | `internal/analyze/generation_stage_test.go` | `internal/testdata/generated/p2-f-stage.json` | `internal/testdata/p2/f-stage-fault.json` | immutable generation staging | validated graph artifacts | active pointer |
| P2-F1 | `internal/analyze/generation_publish.go` | `internal/analyze/generation_publish_test.go` | `internal/testdata/generated/p2-f1-publish.json` | `internal/testdata/p2/f1-cas.json` | repo-local active-generation CAS | staged generation manifest | graph construction |
| P2-F2 | `internal/analyze/generation_cache_publish.go` | `internal/analyze/generation_cache_publish_test.go` | `internal/testdata/generated/p2-f2-cache.json` | `internal/testdata/p2/f2-cache-fault.json` | cache/embedding generation namespaces | generation manifest | reader adapters |
| P2-F3 | `internal/repo/registry_publish.go` | `internal/repo/registry_publish_test.go` | `internal/testdata/generated/p2-f3-registry.json` | `internal/testdata/p2/f3-registry-fault.json` | global registry CAS publication | generation manifest | group snapshots |
| P2-F4 | `internal/group/snapshot_publish.go` | `internal/group/snapshot_publish_test.go` | `internal/testdata/generated/p2-f4-group.json` | `internal/testdata/p2/f4-group-fault.json` | group snapshot/member vector CAS | generation manifest | global registry |
| P2-F5 | `internal/analyze/generation_lease.go` | `internal/analyze/generation_lease_test.go` | `internal/testdata/generated/p2-f5-lease.json` | `internal/testdata/p2/f5-lease-fault.json` | reader leases and GC safety | active generation pointer | publication logic |
| P2-F6 | `internal/analyze/publication_fault_matrix.go` | `internal/analyze/publication_fault_matrix_test.go` | `internal/testdata/generated/p2-f6-fault-matrix.json` | `internal/testdata/p2/f6-faults.json` | publication failure-atomicity matrix | all P2-F owners | target runtime |
| P2-G | `internal/analyze/v2_cutover.go` | `internal/analyze/v2_cutover_test.go` | `internal/testdata/generated/p2-g-cutover.json` | `internal/testdata/p2/g-rollback.json` | identity-v2 cutover/legacy ambiguity | accepted P2-A..P2-F evidence | new semantic algorithms |

The `E2 - P2 Evidence` table row whose first column equals the job ID is normatively incorporated as that job's `Acceptance / Evidence IDs` extension. Its comma-separated exact IDs are conjunctive. This is an exact row binding, not a broad `E2` reference: if the plan job and evidence row differ, either omits `DETECT1`/`COMMIT1`, or contains an undeclared ID, the job fails acceptance.

#### P2-A direct acceptance override

For `P2-A`, the copied Implementation Gate/Acceptance is superseded by this explicit binding: the gate requires `E2-P2A-IMPACT1`, `E2-P2A-SRC1`, `E2-P2A-BUILD1`, `E2-P2A-TEST1`, `E2-P2A-MANIFEST1`, `E2-P2A-HANDSHAKE1`, `E2-P2A-METADATA1`, `E2-P2A-REVIEW1`, `E2-P2A-DETECT1`, and `E2-P2A-COMMIT1`; the behavior test must cover absent, empty, type-invalid, stale, fingerprint-mismatch, and supported metadata rows for every required field. No P2-A acceptance may be recorded from `E2-P2A-REVIEW1` alone.

The copied phrase “P1-A authority fixes the manifest fields” is narrowed here: Child 01 supplies the accepted identity/schema version values only; Child 02/P2-A solely owns the persisted manifest shape, handshake shape, field validation, and reader failure contract. P1-A cannot add, omit, or rename P2 metadata fields.

The same direct binding is required for every later P2 slice: its Acceptance must name the exact ledger IDs for impact, source, full build, behavior oracle, Supervisor, detect-changes, and commit. The child evidence ledger is the authoritative enumeration; a missing ID in either the plan overlay or ledger is a failed traceability gate.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.

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
   - Overlay Gate: apply the mandatory Semantic Remediation Overlay; this is local-child closure, not campaign/release completion.
  - Work Steps:
    1. Run the required final validation for the accepted scope, including full build before final runtime validation. For app/runtime scopes, full build must include Docker image/container build.
    2. Start the real built Docker/container runtime for app/runtime validation. If Docker cannot be built or started, record the blocker and do not substitute a host dev server.
    3. Validate public runtime or UI-facing changes with browser or Playwright evidence against the real built Docker/container runtime. Playwright evidence must include Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
    4. Regenerate generated outputs if source-of-truth changes require it.
    5. Run Anvien detect-changes before commit when implementation work was performed.
    6. Record final validation, detect-changes, benchmark, and commit evidence.
    7. Commit the completed scope and verify the worktree state.
     8. Refresh Child 03 actual-status from the latest accepted Child 02 evidence, append its refresh-log row, and update its next action/work steps before handoff.
     9. Record qualified `E2-PNC-NEXTSTATUS1` and `E2-PNC-HANDOFF1`; bind every required evidence-ledger row and the exact per-job ownership table.
   - Implementation Gate: Pn-A/Pn-B, all local implementation gates, full evidence rows, ownership tables, and the successor refresh must pass; pending P7 is not a local-child blocker.
   - Acceptance: local closure evidence, commit, `E2-PNC-NEXTSTATUS1`, and `E2-PNC-HANDOFF1` are recorded and the worktree is known; campaign/release closure remains separate.

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
