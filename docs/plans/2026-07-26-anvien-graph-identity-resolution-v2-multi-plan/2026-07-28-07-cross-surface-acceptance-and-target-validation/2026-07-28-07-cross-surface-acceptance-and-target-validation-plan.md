# Anvien Cross-Surface Acceptance, Target Validation, And Performance Plan

## Metadata

- Date: `2026-07-28`
- Status: `draft / P0 complete / implementation not yet authorized`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md`
- Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Authoring plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Predecessor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Successor child: `none — terminal child`

## Goal

prove the complete invariant on Anvien and the real target before closure.

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
- When this terminal child ends, update the roadmap's actual status and latest evidence before campaign handoff.

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
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.

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
