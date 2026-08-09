# Anvien Graph Identity Contract and Strict Graph Construction Plan

## Metadata

- Date: `2026-07-28`
- Status: `active / P1-A complete and committed 2fec691a / P1-B open`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md`
- Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Predecessor child: `none — first child`
- Successor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md`

## Goal

ratify and implement the Declaration/Symbol identity foundation without activating v2 for readers.

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
- Both `time` declarations and both `now` declarations in the bounded target remain distinct and traceable; no source occurrence is overwritten.
- The target's known scanner baseline remains quarantined at the same exact eight omissions unless a separate scanner plan is authorized; this plan creates `0` new File-node omissions.
- No source, fixture, report, probe, or temporary investigation artifact is written into or copied from `E:\cheapapp.org`; operational graph/index output remains under `E:\cheapapp.org\.anvien`.
- Before identity cutover, at least five measured v1 runs on the same commit-bound corpus/config/build/machine/cache policy establish analyze median, Ladybug-load median, native-query p95, fallback-query p95, graph size, and peak RSS. Final medians and p95 values regress by no more than `10%`, peak RSS by no more than `15%`, unless the owner accepts a measured exception with both baseline and final values before cutover.
- Every touched/new source and test file passes the one-file/one-responsibility review; `0` new catch-all files and `0` unrelated responsibility additions.
- Every slice has a separate commit after full build, boundary validation, ledger update, Supervisor PASS, and detect-changes.

## Semantic Remediation Overlay (mandatory)

The copied P1 phase block above is historical provenance and remains unchanged. This overlay is the stronger execution contract; it controls whenever a copied gate is weaker or contradictory.

- Local closure is a checkpoint, not campaign completion. Pending P7 rows do not block Child 01 local closure; campaign/release closure remains a Child 07 responsibility.
- Every implementation slice Acceptance is conjunctive with the complete matching evidence-ledger row (`IMPACT`, `SRC`, `BUILD`, behavior `TEST`/oracle, `REVIEW`, `DETECT`, and `COMMIT`). A `REVIEW1` record alone cannot close a slice.
- Before each job opens, record an exact ownership table with `File`, one unique responsibility, allowed links, and prohibited contents for every production, test, generated, and fixture file. Wildcards, `TBD`, catch-all owners, or mixed responsibilities fail the implementation gate.
- `Pn-C` must create qualified `E2-PNC-NEXTSTATUS1` and `E2-PNC-HANDOFF1`. The former proves the next child actual-status/refresh-log/next-action/work-step refresh; the latter proves the local evidence, commit, predecessor gate, and successor opening condition using this child’s own `NEXTSTATUS1` after that refresh, never a future successor-owned record. Missing or stale records block closure.
- The qualified names are `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-NEXTSTATUS1` and `2026-07-28-01-graph-identity-contract-and-strict-construction::E2-PNC-HANDOFF1`; downstream plans must consume these namespaced records.
- Identity contract acceptance must prove the Declaration/Symbol/range/scope tuple, collision and merge behavior, and opaque downstream references without silently changing the active v1 path. Child 02 owns the expanded manifest/handshake metadata correction; Child 01 supplies the identity fields it consumes.

### Identity job ownership/evidence binding (mandatory)

Before each P1-A–P1-E/P1-C0* job opens, its closed file set must be recorded in `E2-PNC-OWNERSHIP1`. Implementation jobs require exact production, test, generated, and fixture paths. Contract-only P1-A instead requires one exact authority-document path and explicit reasoned `N/A` entries for production code, test code, generated runtime evidence, and fixtures. Each source file owns one identity responsibility only; graph adapters may link to many consumers but may not construct a second identity scheme. Every copied Acceptance line is conjunctive with its exact `IMPACT1`, `SRC1`, `BUILD1`, behavior `TEST1`/`ORACLE1`, `REVIEW1`, `DETECT1`, and `COMMIT1` ledger IDs. A review-only or parity-only record cannot close an identity slice.

The binding is literal and follows each job type. The contract-only P1-A requires `E1-P1A-CONTRACT1`, `E1-P1A-REVIEW1`, and `E1-P1A-COMMIT1`. Implementation jobs such as P1-B and P1-C0 require their exact `IMPACT1`, `SRC1`, `BUILD1`, behavior `TEST1`, `REVIEW1`, `DETECT1`, and `COMMIT1` ledger IDs. Every remaining P1 implementation slice must bind the complete exact row declared for that slice; no nonexistent generic ID is invented and no review-only record closes implementation.

| Job | Primary owner file | Test owner file | Generated evidence file | Fixture file | Unique responsibility | Allowed links | Prohibited contents |
|---|---|---|---|---|---|---|---|
| P1-A | `docs/contracts/graph-identity-generation-and-migration-contract.md` (NEW) | `N/A` — documentation-only authority review | `N/A` — evidence is recorded in the child ledger and Supervisor report | `N/A` — accepted problem/decision records are inspect-only inputs | accepted identity/ownership authority only | accepted problem report, decision table, identity fields | production/test code, graph mutation, or reader policy |
| P1-B | `internal/graphidentity/declaration_id.go` (NEW) | `internal/graphidentity/declaration_id_test.go` (NEW) | `internal/testdata/generated/p1-b-identity.json` | `internal/testdata/p1/b-ranges.json` | range/DeclarationID/SymbolID/SymbolRef types | canonical position/owner inputs | merge or graph mutation |
| P1-C0 | `internal/graph/declaration_occurrences.go` (NEW) | `internal/graph/declaration_occurrences_test.go` (NEW) | `internal/testdata/generated/p1-c0-occurrences.json` | `internal/testdata/p1/c0-collisions.json` | lossless declaration occurrence index | declaration facts | relationship identity |
| P1-C0A | `internal/graph/relationship_identity.go` (NEW) | `internal/graph/relationship_identity_test.go` (NEW) | `internal/testdata/generated/p1-c0a-relationships.json` | `internal/testdata/p1/c0a-source-sites.json` | RelationshipID/source-site aggregation | immutable occurrence refs | node identity |
| P1-C0B | `internal/graph/validation.go` (NEW) | `internal/graph/validation_test.go` (NEW) | `internal/testdata/generated/p1-c0b-closure.json` | `internal/testdata/p1/c0b-invalid.json` | decode/closure validation | immutable graph view | mutation or repair |
| P1-C | `internal/graphidentity/symbol_mapping.go` (NEW) | `internal/graphidentity/symbol_mapping_test.go` (NEW) | `internal/testdata/generated/p1-c-symbols.json` | `internal/testdata/p1/c-shadowing.json` | declaration-to-symbol mapping | declaration facts, owner scopes | graph mutation |
| P1-D | `internal/graph/mutation.go` (NEW) | `internal/graph/mutation_test.go` (NEW) | `internal/testdata/generated/p1-d-mutation.json` | `internal/testdata/p1/d-duplicates.json` | strict graph mutation API | graph types/validation | producer semantics |
| P1-D1 | `internal/graph/core_producer_adapter.go` (NEW) | `internal/graph/core_producer_adapter_test.go` (NEW) | `internal/testdata/generated/p1-d1-core.json` | `internal/testdata/p1/d1-core.json` | core producer migration adapter | strict mutation API | resolver/document producers |
| P1-D2 | `internal/resolution/graph_producer_adapter.go` (NEW) | `internal/resolution/graph_producer_adapter_test.go` (NEW) | `internal/testdata/generated/p1-d2-resolution.json` | `internal/testdata/p1/d2-resolution.json` | resolution/projection producer migration | strict mutation API | ancillary/document producers |
| P1-D3 | `internal/graph/ancillary_producer_adapter.go` (NEW) | `internal/graph/ancillary_producer_adapter_test.go` (NEW) | `internal/testdata/generated/p1-d3-ancillary.json` | `internal/testdata/p1/d3-ancillary.json` | thin explicit-operation orchestration for the inventory-defined ancillary producer set | strict mutation API plus document/COBOL/semantic/diagnostic producer links | domain logic, identity policy, or core/resolution producer ownership |
| P1-E | `internal/analyze/shadow_identity_v2.go` (NEW) | `internal/analyze/shadow_identity_v2_test.go` (NEW) | `internal/testdata/generated/p1-e-shadow.json` | `internal/testdata/p1/e-shadow.json` | shadow-v2 emission/comparison | accepted identity/mutation contracts | active-v1 cutover |

This 11-job table is the closed Child-01 ownership set. Any path change must be made in the affected row before editing and cannot merge responsibilities.

For every P1 job, the evidence-ledger row whose first column equals that exact job ID is normatively incorporated as the job's `Acceptance / Evidence IDs` extension. All listed IDs are conjunctive; any plan/ledger mismatch or missing detect/commit row fails acceptance.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md` has a final P0 decision.
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

- [x] P1-A: Ratify graph identity and ownership contract.
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
     8. Refresh Child 02 actual-status from the latest accepted Child 01 evidence, append its refresh-log row, and update its next action/work steps before handoff.
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
