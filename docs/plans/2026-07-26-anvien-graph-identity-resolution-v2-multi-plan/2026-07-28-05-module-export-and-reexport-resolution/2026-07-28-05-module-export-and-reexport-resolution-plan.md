# Anvien Module Export Tables And Barrel/Re-Export Resolution Plan

## Metadata

- Date: `2026-07-28`
- Status: `draft / P0 complete / implementation not yet authorized`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md`
- Source plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
- Authoring plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Predecessor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
- Planned successor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`

## Goal

resolve repository-backed imports through module export surfaces to an `InternalSymbolRef`, `ModuleNamespaceRef`, or `GapRef` with complete proof chains, and hand an immutable `ExternalResolutionCandidate` (`ModuleRef`, exported name, meaning, and proof) to P6 when declaration-universe work is required. P5 never creates `ExternalSymbolRef`.

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

- Both bounded barrel calls resolve through their proof chains to the terminal function (`2/2`); direct `IMPORTS` edges still describe syntactic module dependencies.
- Explicit missing imports, type/value mismatches, ambiguous star exports, pure cycles, invalid config, absent package declarations, and security-rejected paths produce their exact structured status with no global-name fallback.
- The target's known scanner baseline remains quarantined at the same exact eight omissions unless a separate scanner plan is authorized; this plan creates `0` new File-node omissions.
- No source, fixture, report, probe, or temporary investigation artifact is written into or copied from `E:\cheapapp.org`; operational graph/index output remains under `E:\cheapapp.org\.anvien`.
- Every touched/new source and test file passes the one-file/one-responsibility review; `0` new catch-all files and `0` unrelated responsibility additions.
- Every slice has a separate commit after full build, boundary validation, ledger update, Supervisor PASS, and detect-changes.

## Semantic Remediation Overlay (mandatory)

The copied P5 phase block above is historical provenance and remains unchanged. This overlay is the stronger execution contract; it controls whenever a copied gate is weaker or contradictory.

All copied manifest/handshake field lists in this child are inspect-only historical context; the qualified Child 02 manifest contract (`2026-07-28-02-versioned-persistence-and-v2-cutover::E2-P2A-MANIFEST1`, `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-P2A-HANDSHAKE1`, `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-P2A-METADATA1`) is authoritative and is consumed without reinterpretation.

- Local closure is a checkpoint, not campaign completion. Pending P7 rows do not block Child 05 local closure; campaign/release closure remains a Child 07 responsibility.
- Before `P5-A` opens, the exact qualified predecessor gate remains `2026-07-28-04-typescript-export-semantics::E2-PNC-HANDOFF1`; the P4 fact manifest is consumed inspect-only and cannot be regenerated by P5.
- Every implementation slice Acceptance is conjunctive with the complete matching evidence-ledger row (`IMPACT`, `SRC`, `BUILD`, behavior `TEST`/oracle, `REVIEW`, `DETECT`, and `COMMIT`). A `REVIEW1` record alone cannot close a slice.
- Before each job opens, record an exact ownership table with `File`, one unique responsibility, allowed links, and prohibited contents for every production, test, generated, and fixture file. Wildcards, `TBD`, catch-all owners, or mixed responsibilities fail the implementation gate.
- `Pn-C` must create qualified `E2-PNC-NEXTSTATUS1` and `E2-PNC-HANDOFF1`; refresh Child 06 actual-status/refresh-log/next-action/work-step rows from the accepted Child 05 evidence before handoff, and bind the handoff to this child’s own `NEXTSTATUS1`, never Child 06’s future record.
- The qualified names are `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-NEXTSTATUS1` and `2026-07-28-05-module-export-and-reexport-resolution::E2-PNC-HANDOFF1`; Child 06 must consume these namespaced records.
- **Module/export correction:** P4 is the producer of `ModuleRequestFact` and `ImportBindingFact`; P5 consumes them and owns path/specifier resolution, export tables, re-export traversal, and derived `resolvedExportEntryCount`/`publicApiSymbolCount`. P5 must prove a barrel with zero physical declarations can expose a non-empty syntax-derived export surface, and must report unchanged physical path-resolution/`IMPORTS` counts before versus after export traversal.
- P5 pre-flight ownership is exact: P5-A `internal/resolution/module_reference.go` owns module/path conditions only; P5-B `internal/resolution/module_export_table.go` owns export-table construction only; P5-C `internal/resolution/reexport_resolution.go` owns traversal/proof only; P5-D keeps `internal/resolution/emit.go` as a thin adapter and uses a dedicated re-export behavior test. None may load declarations, invent physical definitions, or mutate P4 facts.

### P5 derived-state, handoff, and oracle contract (mandatory)

P5-A must name and validate the exact dependency `2026-07-28-04-typescript-export-semantics::E2-PNC-HANDOFF1` (source `P8-C`, local `Pn-C`) before opening. “Child 04 qualified handoff” without the child slug and exact ID is invalid. The handoff must expose immutable P4 `ModuleRequestFact`, `ImportBindingFact`, syntactic `ExportFact`, and `directExportedDefinitionCount` manifests; P5 may not regenerate or mutate them.

Derived ownership is closed as follows: P5-B owns only the per-module export table; P5-C owns traversal/proof plus `resolvedExportEntryCount` and `reachableThroughBarrel`; P5-D's `internal/resolution/emit.go` owns serialization only, while the dedicated `internal/resolution/public_api_projection.go` owns the single `projectPublicApiSymbolCount` function that computes `publicApiSymbolCount` from the P5-C terminal export-entry set. No other job may write any of the three derived fields. The child ledger must report three separate counts and three separate owner evidence IDs: `E2-PNC-MODULE1A` (direct count consumed unchanged), `E2-PNC-MODULE1B` (resolved entries/reachability), and `E2-PNC-MODULE1C` (public API count).

The resolver oracle is explicit, not aggregate: `E2-PNC-MODULE1D` records the zero-physical-declaration barrel fixture with `physicalDeclarationCount=0`, `resolvedExportEntryCount>0`, terminal proof hops, and exact public API entries; `E2-PNC-MODULE1E` records pre/post physical path-resolution counts and syntactic `IMPORTS` counts with both absolute values and zero deltas. A `2/2` call count alone cannot close either requirement.

| Job | File | Unique responsibility | Allowed links | Prohibited contents |
|---|---|---|---|---|
| P5-A | `internal/resolution/typescript_project.go` (NEW) | hash-bound TypeScript project-profile parsing and immutable profile publication | tsconfig/package metadata | module path resolution, export traversal, declaration loading |
| P5-A | `internal/resolution/module_reference.go` (NEW) | module/path conditions and physical path accounting | project profile, ModuleRequestFact | export traversal or declaration loading |
| P5-A | `internal/resolution/module_reference_test.go` (NEW) | module/path matrix and before-count oracle | module reference owner | export-table mutation |
| P5-A | `internal/resolution/testdata/modules/path-counts.json` (NEW fixture) | pre/post path-resolution expected counts | path tests | target source |
| P5-A | `internal/testdata/generated/p5-a-module-paths.json` (NEW generated evidence) | observed ModuleRef and physical path-count records | module-reference owner and tests | export-table results |
| P5-B | `internal/resolution/module_export_table.go` (NEW) | per-module syntactic export table | immutable P4 facts | traversal, external declarations, public API |
| P5-B | `internal/resolution/module_export_table_test.go` (NEW) | table construction/meaning precedence matrix | table owner | terminal traversal |
| P5-B | `internal/resolution/testdata/modules/zero-physical-barrel.ts` (NEW fixture) | barrel syntax with zero physical declaration files | table/traversal tests | cheapapp source |
| P5-B | `internal/testdata/generated/p5-b-export-table.json` (NEW generated evidence) | observed syntactic export-table entries | table owner and tests | traversal decisions |
| P5-C | `internal/resolution/reexport_resolution.go` (NEW) | alias/star/namespace traversal, proof hops, resolved count/reachability | P4 facts, P5 tables | declaration loader or P4 mutation |
| P5-C | `internal/resolution/reexport_resolution_test.go` (NEW) | cycle/ambiguity/budget/terminal oracle | traversal owner | adapter serialization |
| P5-C | `internal/resolution/testdata/modules/reexport-vectors.json` (NEW fixture) | named semantic-vector expected outcomes | traversal tests | target graph/index |
| P5-C | `internal/testdata/generated/p5-c-reexport-resolution.json` (NEW generated evidence) | observed terminal entries, reachability, and proof hops | traversal owner and tests | public API serialization |
| P5-D | `internal/resolution/emit.go` | thin serialization adapter for immutable P5 derived values | immutable P5 derived values | path resolution, traversal, or count computation |
| P5-D | `internal/resolution/emit_reexport_test.go` (NEW) | adapter output and count preservation matrix | emit owner | new resolution decisions |
| P5-D | `internal/resolution/testdata/modules/emit-counts.json` (NEW fixture) | expected direct/resolved/public count separation | emit tests | target runtime output |
| P5-D | `internal/testdata/generated/p5-d-export-output.json` (NEW generated evidence) | emitted derived counts and path/IMPORTS preservation record | emit owner and tests | resolver decisions |
| P5-D | `internal/resolution/public_api_projection.go` (NEW) | sole `projectPublicApiSymbolCount` derived-count function | immutable P5 terminal export-entry set | serialization or path resolution |
| P5-D | `internal/resolution/public_api_projection_test.go` (NEW) | public API count projection behavior | public-api count owner | traversal decisions |
| P5-D | `internal/resolution/testdata/modules/public-api-counts.json` (NEW fixture) | expected terminal public API sets/counts | public-api tests | target source |
| P5-D | `internal/testdata/generated/p5-d-public-api-counts.json` (NEW generated evidence) | observed public API count records | public-api owner and tests | P4 syntax facts |
| P5-A | `internal/resolution/import_resolution.go` | thin physical module-path adapter | ModuleRequestFact, module profile | export traversal or declaration loading |
| P5-B | `internal/resolution/indexes.go` | sole write-owner for resolution-index registration coordination | P5 table owner | new traversal algorithm or P4 fact mutation |
| P5-C | `internal/resolution/indexes.go` | preserve-only link to the P5-B registration coordinator | P5 traversal owner | edits, new traversal algorithm in coordinator, or P4 fact mutation |
| P5-D | `internal/resolution/resolve.go` | resolver-stage orchestration/delegation link only | immutable P5 derived values | path resolution, traversal, or fallback guessing |
| P5-A | `internal/providers/tsjs/imports.go` | sole write-owner for TS module/import syntax-registration coordination | P4 ModuleRequestFact and ImportBindingFact | terminal declaration lookup |
| P5-B | `internal/providers/tsjs/imports.go` | preserve-only link to the P5-A syntax-registration coordinator | P4 ImportBindingFact | edits or terminal declaration lookup |

Repeated existing coordinator paths in this table do not create multiple business owners. P5-A is the sole write-owner for `internal/providers/tsjs/imports.go`; P5-B consumes that link preserve-only. P5-B is the sole write-owner for `internal/resolution/indexes.go`; P5-C consumes that link preserve-only. `internal/resolution/resolve.go` remains solely owned by P5-D for resolver-stage orchestration. No second slice may edit a shared coordinator; all new semantic logic remains in the dedicated single-responsibility owner file for that job.

The P5 benchmark and evidence ledgers must keep `directExportedDefinitionCount`, `resolvedExportEntryCount`, and `publicApiSymbolCount` as three rows/IDs; a single conflated count row is non-accepting.

For every P5 job, the evidence-ledger row whose first column equals that exact job ID is normatively incorporated as the job's `Acceptance / Evidence IDs` extension. All listed IDs are conjunctive; any plan/ledger mismatch or missing detect/commit row fails acceptance.

#### P5 resolver acceptance override

The copied P5-D `2/2` acceptance is non-terminal by itself. P5 closes only when the zero-physical-declaration barrel oracle reports `physicalDeclarationCount=0`, a proof-backed non-empty `resolvedExportEntryCount`, exact terminal hops, and the expected `publicApiSymbolCount`, while the absolute physical path-resolution and syntactic `IMPORTS` counts are equal before and after traversal. `directExportedDefinitionCount`, `resolvedExportEntryCount`, and `publicApiSymbolCount` remain three separately owned metrics; no P5 job may mutate the P4 fact manifest or substitute a physical declaration for a zero-physical barrel.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.

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
     8. Refresh Child 06 actual-status from the latest accepted Child 05 evidence, append its refresh-log row, and update its next action/work steps before handoff.
     9. Record qualified `E2-PNC-NEXTSTATUS1` and `E2-PNC-HANDOFF1`; bind every required evidence-ledger row and the exact per-job ownership table.
   - Implementation Gate: Pn-A/Pn-B, all local implementation gates including zero-physical-declaration barrel and unchanged path-resolution metrics, full evidence rows, ownership tables, and the successor refresh must pass; pending P7 is not a local-child blocker.
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
