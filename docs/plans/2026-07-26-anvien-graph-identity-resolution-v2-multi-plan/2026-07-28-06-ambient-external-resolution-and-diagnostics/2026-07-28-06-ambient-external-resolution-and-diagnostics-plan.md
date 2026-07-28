# Anvien Ambient/External Declaration Universe And Truthful Diagnostics Plan

## Metadata

- Date: `2026-07-28`
- Status: `draft / P0 complete / implementation not yet authorized`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`
- Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Authoring plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Predecessor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
- Planned successor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`

## Goal

resolve configured TypeScript standard-library and local package declarations without polluting the repository graph or guessing diagnostics.

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

- Every surface `S0`–`S11` exposes the same canonical node/relationship records for the active generation; native Cypher and fallback Cypher are measured separately, and field-level differences plus orphan references are both zero.
- `Promise`, `Math.max`, and `Math.min` are resolved external/intrinsic outcomes or explicit external-capability outcomes, never false `in_repo_unresolved` gaps.
- Explicit missing imports, type/value mismatches, ambiguous star exports, pure cycles, invalid config, absent package declarations, and security-rejected paths produce their exact structured status with no global-name fallback.
- The target's known scanner baseline remains quarantined at the same exact eight omissions unless a separate scanner plan is authorized; this plan creates `0` new File-node omissions.
- No source, fixture, report, probe, or temporary investigation artifact is written into or copied from `E:\cheapapp.org`; operational graph/index output remains under `E:\cheapapp.org\.anvien`.
- Every touched/new source and test file passes the one-file/one-responsibility review; `0` new catch-all files and `0` unrelated responsibility additions.
- Every slice has a separate commit after full build, boundary validation, ledger update, Supervisor PASS, and detect-changes.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.

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
