# Anvien Versioned Persistence, Opaque Consumers, Atomic Generation, And V2 Cutover Actual Status

Title: Anvien Versioned Persistence, Opaque Consumers, Atomic Generation, And V2 Cutover
Date: 2026-07-28
Status: P0 Complete / implementation blocked pending predecessor handoff and owner authority
Companion plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-02-versioned-persistence-and-v2-cutover/2026-07-28-02-versioned-persistence-and-v2-cutover-benchmark.md`
Source actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-actual-status.md`
Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
Predecessor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
Successor child: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`

## Purpose

This file records the real current state before implementation.

Implementation must not start until the target scope has a completed status row, evidence IDs, and a downstream plan decision.

This file does not replace `evidence.md`. It classifies current state from evidence.

Use exact evidence IDs from `evidence.md`, such as `E0-P0A-SRC1`, not broad section IDs such as `E0` or `E1`.

## Freshness / Refresh Rules

This actual-status file is a living current-state record, not a one-time P0 snapshot.

P0 records the baseline before implementation. After implementation begins, keep the Current Status Matrix updated so the next agent can trust it as the latest repo reality.

Update this file:

- after each completed implementation slice;
- before starting the next phase if repo state changed;
- whenever evidence changes a current-state classification;
- whenever the next phase's status assumptions, next action, or work steps need updating because reality differs from the previous status.

When refreshing status:

- update only the rows affected by the completed work or new evidence;
- use explicit transitions such as `missing -> correct`, `partial -> correct`, `fake-or-stub -> removed`, or `unbound -> bound-correct`;
- append a Status Refresh Log row instead of deleting history;
- keep detailed proof in `evidence.md`; store only classifications, evidence IDs, touch mode, and plan consequences here.

## Scope

Target scope:

- graph Declaration/Symbol identity, range, meaning, strict mutation, persistence version, active generation, and opaque consumer contracts in `E:\Anvien`;
- TypeScript binding patterns, export facts, module/re-export resolution, ambient/external declarations, and structured diagnostic outcomes;
- in-place acceptance against `E:\cheapapp.org` with operational output under `E:\cheapapp.org\.anvien`;
- cross-surface JSON, Ladybug, CLI, MCP, HTTP, Web, cache, group, process, community, and embedding consistency where identity/generation affects them.

Out of scope:

- scanner pruning for nested directories named `env`, `target`, or `logs`;
- target source edits, target copies, target fixtures, or investigation artifacts in `E:\cheapapp.org`;
- full TypeScript compiler conformance, product process semantics, and unrelated refactors;
- runtime network/package install/scripts or graphing all declarations from `node_modules`.

## Relationship / Impact Evidence

For each target file, prefer:

```text
anvien file-detail <path> --repo <repo> --json
```

Record how many files the target is related to before deciding touch mode. A file with many relationships may still be editable, but the plan must narrow the exact phase, touch mode, and validation needed.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| `internal/graph/types.go` | `E0-P0A-FD1` | 238 | central node/relationship storage and indexing | high file scope; `Graph.AddNode`/`Graph.init` CRITICAL |
| `internal/scopeir/facts.go` | `E0-P0A-FD2` | 231 | mixed ScopeIR fact contracts used across providers/resolver | medium file scope; dedicated owner extraction required |
| `internal/scopeir/range.go` | `E0-P0A-FD3` | 227 | shared position/range contract | medium file scope; encoding contract affects all providers |
| `internal/scopeir/definition_index.go` | `E0-P0A-FD4` | 225 | definition identity index | medium file scope; currently hides duplicates |
| `internal/resolution/indexes.go` | `E0-P0A-FD5` | 46 | workspace, import binding, graph ID construction | high file scope; several CRITICAL symbols |
| `internal/resolution/resolve.go` | `E0-P0A-FD6` | 50 | call/type/member resolution | high file scope; CRITICAL call/type paths |
| `internal/resolution/emit.go` | `E0-P0A-FD7` | 42 | graph projection | high file scope; identity/export parity |
| `internal/providers/tsjs/definitions.go` | `E0-P0A-FD8` | 16 | TS definition extraction | high file scope; array binding/export loss |
| `internal/providers/tsjs/imports.go` | `E0-P0A-FD9` | 15 | TS import/re-export syntax facts | high file scope |
| `internal/resolution/import_resolution.go` | `E0-P0A-FD10` | 33 | language/path import strategies | high file scope |
| `internal/graphhealth/diagnostics.go` | `E0-P0A-FD11` | 29 | diagnostic classification | high file scope; classifier HIGH impact |
| `internal/analyze/analyze.go` | `E0-P0A-FD12` | 182 | analyze/build/load/snapshot orchestration | high file scope; publication CRITICAL |
| `internal/repo/types.go` | `E0-P0A-FD13` | 72 | persisted repository metadata | high file scope; version fields missing |
| `internal/lbugload/csv.go` | `E0-P0A-FD14` | 19 | graph-to-Ladybug CSV projection | high file scope; duplicate/export parity risk |
| `internal/httpapi/graph.go` | `E0-P0A-FD15` | 22 | graph HTTP response | high file scope; metadata omitted |
| `anvien-web/src/services/backend-client.ts` | `E0-P0A-FD16` | 24 | Web graph stream parsing | high file scope; version/generation absent |
| `internal/repo/registry.go` / `internal/group/sync.go` / `internal/group/storage.go` | source inventory required in P2-A1 | bounded | global registry and group contract publication | high generation scope; outside repo-local graph |
| `internal/mcp/resource_cache.go` / `internal/filecontext/*` / embedding stores | source inventory required in P2-A1 | bounded | cache/embedding readers and invalidation | high generation scope; stale refs possible |
| native and fallback Cypher readers | source inventory required in P2-A1 | bounded | alternate persistence/query backends | high parity scope; separate gates required |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Already behaves as required. | Preserve. Add evidence or tests only if needed. |
| `partial` | Some required behavior exists, but gaps remain. | Change only the missing parts. Preserve correct parts. |
| `wrong` | Current behavior, source, or contract is incorrect. | Replace with required behavior. Record the exact reason. |
| `missing` | Required behavior, source, or contract does not exist. | Implement the missing piece only. |
| `unbound` | Surface exists but is not wired to the real source, flow, or contract. | Bind to the real source only. Preserve approved surface. |
| `fake-or-stub` | Prototype, demo, mock, fallback, or placeholder data is being used as real behavior. | Remove fake behavior or replace it with an approved truthful state. |
| `blocked` | Source, authority, contract, or required evidence is unclear. | Stop. Do not implement until resolved. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| Architecture authority | no relevant approved SPEC/ADR; accepted reports diagnose but do not authorize remediation | owner-accepted identity/export/external/generation contract | blocked | N/A | `E0-P0A-SPEC1`, `E0-P0A-REPORT2` | open P1-A only; no production code |
| Plan slice template structure | first frozen snapshot had compact Scope Boundary/Acceptance blocks; B1 remediation expanded all eligible slices into auditable fields without changing semantics | every eligible slice has machine-readable ownership and seven-part acceptance gates | correct | 99 eligible slices | `E0-P0A-TEMPLATE4` | request Supervisor re-review; keep implementation blocked until owner authority |
| One-file responsibility | several current owner files mix multiple fact/index/orchestration concerns | each touched/new file owns one semantic responsibility; mixed owners are split before additions | partial | 15-238 per file | `E0-P0A-FD1..E0-P0A-FD16`, `E0-P0A-USER2` | enforce ownership map in every slice |
| Declaration/Symbol identity | one graph ID derived from file/name/arity; scope/range omitted; duplicate nodes overwrite | separate DeclarationID/SymbolID, exact ranges, deterministic collision-safe mapping | wrong | 46/225/238 | `E0-P0A-SRC1`, `E0-P0A-SRC2`, `E0-P0A-IMPACT1..E0-P0A-IMPACT3` | P1 then P2 |
| Strict graph mutation/decode | implicit replacement and silent duplicate suppression exist before and during graph/persistence load | explicit insert/enrich/replace plus fail-closed validation | wrong | 19-238 | `E0-P0A-SRC1..E0-P0A-SRC3` | P1-D |
| Index version/generation | metadata lacks graph/identity/ScopeIR/generation contract; readers do not fail closed | versioned metadata checked by every reader | missing | 22-182 | `E0-P0A-SRC5`, `E0-P0A-FD12..E0-P0A-FD16` | P2-A |
| Reader matrix | no committed denominator covering JSON loaders, startup/status/index, native/fallback Cypher, caches, embeddings, registries, groups, process/community, and audit probes | one source-derived row per reader/backend with guard owner, fixture, expected error, and zero-unlisted scan | missing | bounded until matrix | `E0-P0A-SRC5..E0-P0A-SRC7`, `E0-P0A-FD12..E0-P0A-FD16` | P2-A1 inventory, then P2-A2..P2-A15 guards, before any cutover |
| Old-reader protocol handshake | v1 binaries/clients can ignore unknown fields and may open a v2-compatible path | versioned non-overlapping layout, `readerProtocolVersion`/`minReaderProtocol` negotiation before body open, explicit old-client fixture | missing | bounded until matrix | `E0-P0A-SRC5..E0-P0A-SRC7` | P2-A contract, P2-A2..P2-A15 wiring, and P2-E2 row acceptance before P2-G |
| Atomic generation publication | force analyze removes live artifacts and publishes graph/DB/meta separately | independently stage, publish, generation-bind caches/embeddings, publish repo/group registries, lease/GC, then validate rollback | wrong | 182 | `E0-P0A-SRC6`, `E0-P0A-IMPACT9` | P2-F..P2-F6 before cutover |
| Opaque ID consumers | multiple backend/Web consumers parse labels, paths, lines, route/community values from IDs or order by IDs | explicit fields and generation-qualified references | wrong | 22-72 per owner | `E0-P0A-SRC7` | P2-C..P2-C6, P2-D..P2-D2, and P2-E/P2-E1 |
| Registry/group/cache publication | global registry, group contracts, embeddings, and in-memory caches are outside the current repo-local atomic graph lifecycle | immutable generation-qualified snapshots, CAS active epoch, reader leases, and lease-safe GC across `S7`–`S10` | missing | bounded | `E0-P0A-SRC5..E0-P0A-SRC7` | P2-F2..P2-F6 before cutover |
| Group member-generation vector | group contracts have no atomic vector tying each member repo to one committed generation | `{groupEpoch, memberRepoGenerationVector}` with CAS conflict rejection, concurrent interleaving tests, and lease-safe prior-snapshot retention | missing | bounded | `E0-P0A-SRC5..E0-P0A-SRC7` | P2-F4/P2-F5/P2-F6 before cutover |
| Surface parity denominator | prior work used a smaller projection set and did not separate native/fallback Cypher or cache/registry/derived surfaces | fixed `S0`–`S11` matrix with canonical field-level and orphan-reference checks | missing | bounded | `E0-P0A-SRC5..E0-P0A-SRC7` | P4/P7 use the exact matrix |
| Target boundary | target is analyzed in place; graph stays in target `.anvien`; reports stay in Anvien; supported analyze may timestamp ignored generated guidance files | preserve exact ownership with no contamination/copy and record tool-boundary side effects | correct | 114,125 target relationships | `E0-P0A-BOUNDARY1`, `E0-P0A-BOUNDARY2` | preserve-only except operational analyze and documented ignored-file timestamp caveat |
| Scanner omission | exact eight TS/JS paths omitted because of ignored segments | no additional omissions from this plan; scanner repair separate | wrong but out of scope | 8 missing / 0 extra | `E0-P0A-REPORT1`, `E0-P0A-BOUNDARY2` | do-not-touch; quarantine baseline |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-07-26 | `master` at `1932359bee5d78fd8e6167ef94e7974f33e85bd0`; fresh Anvien self-graph | full plan scope | initial classifications recorded; P0 complete; architecture authority remains blocked | `E0-P0A-GIT1`, `E0-P0A-GRAPH1`, `E0-P0A-STATUS1` | only P1-A may open after owner direction; no production code |
| R1 | 2026-07-26 | same HEAD after plan artifacts were added; fresh `anvien analyze E:\Anvien --force --json` | P0 graph/document baseline and reader-matrix planning artifact | self graph refreshed to 1,505 scanned Files / 676 parsed code / 84,558 nodes / 123,398 relationships; no production source changed; architecture authority remains blocked | `E0-P0A-GRAPH1`, `E0-P0A-STATUS1` | preserve the refreshed baseline; only P1-A may open after owner direction |
| R2 | 2026-07-26 | same HEAD after matrix/ledger/slice revisions; fresh `anvien analyze E:\Anvien --force --json` | reader seed, surface unions, P3/P4 adapter slices, P6 outcome split | self graph reproduced 1,505 scanned Files / 676 parsed code / 84,558 nodes / 123,398 relationships; 153 matrix anchors exist; no production source changed; architecture authority remains blocked | `E0-P0A-GRAPH2`, `E0-P0A-STATUS1` | preserve current plan snapshot; only P1-A may open after owner direction |
| R3 | 2026-07-26 | same HEAD after reader-row range synchronization; fresh `anvien analyze E:\Anvien --force --json` | final P0 graph/document baseline and matrix continuity | self graph reproduced 1,505 scanned Files / 676 parsed code / 84,558 nodes / 123,398 relationships; matrix rows R01-R153 are contiguous; no production source changed; architecture authority remains blocked | `E0-P0A-GRAPH3`, `E0-P0A-STATUS1` | preserve current plan snapshot; only P1-A may open after owner direction |
| R4 | 2026-07-26 | same HEAD after Web/dispatcher/backend matrix corrections and per-slice Mini-QA normalization; fresh `anvien analyze E:\Anvien --force --json` | final P0 graph/document baseline, reader denominator, and template structure | self graph reproduced 1,505 scanned Files / 676 parsed code / 84,558 nodes / 123,398 relationships; matrix rows R01-R195 are contiguous with all current anchors present; 64 implementation/closure slices carry explicit per-work-step Mini-QA contracts; no production source changed; architecture authority remains blocked | `E0-P0A-GRAPH4`, `E0-P0A-MATRIX3`, `E0-P0A-TEMPLATE2`, `E0-P0A-STATUS1` | preserve current plan snapshot; only P1-A may open after owner direction |
| R5 | 2026-07-26 | same HEAD after the expanded P2 slice set, template corrections, and final plan-text cleanup; fresh `anvien analyze E:\Anvien --force --json` | current five-artifact plan set, template structure, reader matrix, and self-graph baseline | current self graph is 1,504 scanned Files / 676 parsed code / 84,532 nodes / 123,372 relationships; this differs from and does not claim reproduction of the historical 1,505 / 84,558 / 123,398 snapshot. The current plan has 102 checklist items, 99 eligible slices, 132 eligible numbered work steps, 101 unique trace rows, and 796 non-meta evidence references with 0 missing declarations; matrix rows R01-R195 and all 195 source anchors pass; no production source changed; architecture authority remains blocked | `E0-P0A-GRAPH5`, `E0-P0A-TEMPLATE3`, `E0-P0A-MATRIX4`, `E0-P0A-STATUS1` | freeze the current plan set for independent Supervisor review; only P1-A may open after owner direction |
| R6 | 2026-07-27 | same HEAD; plan-only B1 template remediation and fresh structural audit | Scope Boundary and Acceptance structure | 99/99 eligible slices now have all required Scope Boundary and Acceptance fields with reasoned N/A values; compact eligible blocks and field-indentation failures are 0; no production source or target artifact changed; architecture authority remains blocked | `E0-P0A-TEMPLATE4` | request independent Supervisor re-review; only P1-A may open after owner direction |
| R7 | 2026-07-27 | same HEAD after B1 ledger refresh and first Supervisor report; fresh `anvien analyze E:\Anvien --force --json` | current five-artifact plan set plus review evidence | current self graph is 1,505 scanned Files / 676 parsed code / 84,540 nodes / 123,380 relationships with 0 failed/unsupported/unknown; graph SHA-256 is recorded; `internal/` remains unchanged and target remains untouched; architecture authority remains blocked | `E0-P0A-GRAPH6`, `E0-P0A-TEMPLATE4` | freeze the corrected five-artifact set for Supervisor re-review; only P1-A may open after owner direction |

## Phase Touch Map

Use this map to prevent accidental edits. A related file is not automatically editable.

`Plan-Relevant Relationship File` lists only a relationship file that can directly affect or be affected by the planned phase or slice. Do not copy the full `file-detail` relationship inventory into this map. Include only files whose relationship can affect the phase/slice decision, touch mode, or validation.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| contract/owner map | plan plus future ADR/SPEC | architecture authority | P1-A | edit docs only | `E0-P0A-SPEC1` | no production edit until accepted |
| `internal/scopeir/facts.go` | proposed dedicated fact files | source-of-truth contract | P1-B/P3-A/P4-A | split then thin adapter | `E0-P0A-FD2` | do not append another fact responsibility |
| `internal/scopeir/range.go` | graphidentity declaration owner | shared range input | P1-B | edit/split | `E0-P0A-FD3` | explicit byte/line/column encoding |
| `internal/scopeir/definition_index.go` | graph identity/validation | duplicate suppression | P1-C/P1-D | edit | `E0-P0A-FD4`, `E0-P0A-SRC1` | no silent duplicate |
| `internal/graph/types.go` | graph mutation/validation owners | storage/index adapter | P1-D | split then thin adapter | `E0-P0A-FD1`, `E0-P0A-IMPACT2..E0-P0A-IMPACT3` | CRITICAL; explicit producer migration |
| `internal/resolution/indexes.go` | identity, export table, re-export owners | current mixed coordinator | P1-C/P5-B/P5-C | split then thin adapter | `E0-P0A-FD5` | no new semantic algorithm in coordinator |
| `internal/resolution/resolve.go` | outcome/external owners | call/type/member coordinator | P5-D/P6-C1/P6-C2 | thin adapter | `E0-P0A-FD6` | CRITICAL; no fallback |
| `internal/resolution/emit.go` | narrow projection owners | graph projection | P1-E/P3-C/P4-C/P5-D/P6-D | thin adapter/edit | `E0-P0A-FD7` | one outcome; property parity |
| `internal/analyze/analyze.go` | generation stage/publish owners | orchestration | P2-F/P2-F1 | split then thin adapter | `E0-P0A-FD12`, `E0-P0A-SRC6` | never delete live generation first |
| `internal/repo/types.go` | index version owner | persisted metadata | P2-A | split/edit | `E0-P0A-FD13` | every reader validates |
| `internal/repo/registry.go` / `internal/group/sync.go` / `internal/group/storage.go` | generation publication/reader matrix | global registry/group contracts | P2-A12/P2-A13/P2-F3/P2-F4 | inspect then edit by exact owner | `E0-P0A-SRC5..E0-P0A-SRC7` | active epoch and leases; no raw SymbolUID |
| `internal/mcp/resource_cache.go` / `internal/filecontext/*` / embedding stores | generation publication/reader matrix | caches/embedding refs | P2-A9..P2-A11/P2-C3/P2-C5/P2-C6/P2-F2 | inspect then edit by exact owner | `E0-P0A-SRC7` | keys include generation/config/catalog hashes |
| `index-reader-matrix.md` | every graph/index reader and startup path | compatibility denominator and old-reader handshake | P2-A1 | create/edit plan artifact | `E0-P0A-SRC5..E0-P0A-SRC7` | exact path/function, backend, later guard owner, fixture, zero-unlisted scan |
| global registry/group/cache/embedding stores | generation stage/publish owners | external publication set | P2-F2..P2-F6/P7-C | inspect then edit by store owner | `E0-P0A-SRC5..E0-P0A-SRC7` | immutable snapshots, CAS epoch, leases, GC after release |
| native/fallback Cypher paths | persistence reader matrix | alternate backend | P2-A3/P2-A4/P2-B3/P2-B4/P7-C | inspect/edit separately | `E0-P0A-FD14` | separate parity rows |
| `internal/lbugload/csv.go` | persistence projection owners | Ladybug CSV/load | P2-B1/P2-B2/P4-C | split/edit by write vs load responsibility | `E0-P0A-FD14`, `E0-P0A-SRC3..E0-P0A-SRC4` | zero skipped duplicates |
| HTTP/Web graph adapters | graph metadata/version owner | public consumer | P2-E/P2-E1 | edit separately | `E0-P0A-FD15..E0-P0A-FD16` | Docker/browser evidence required for Web |
| `E:\cheapapp.org` source/worktree | `E:\cheapapp.org\.anvien` | integration source / operational output | P3-C2/P4-C2/P5-D/P6-D/P7-B | source preserve; analyze/read graph | `E0-P0A-BOUNDARY1` | no copy/report/fixture/temp artifact |
| scanner owners | target File-node inventory | excluded defect | all phases | do-not-touch | `E0-P0A-BOUNDARY2` | exact eight omissions quarantined |

## Detailed Findings

### Version, generation, and consumer migration

Current state:

Index metadata is insufficient, publication is not atomic, readers can ignore unknown fields, and multiple consumers parse meaning/order from IDs.

Required state:

```text
Every artifact and request is bound to one supported version/generation.
Complete generation publication is atomic and rollback-safe.
All IDs are opaque to consumers.
```

Evidence:

- `E0-P0A-SRC5..E0-P0A-SRC7`: metadata/publication/consumer facts.
- `E0-P0A-FD12..E0-P0A-FD16`: high-scope persistence/API/Web owners.

Relationship and impact:

- Related file count: 22-182 for primary owner files, with wider transitive impact.
- Relationship summary: analyze -> graph/Ladybug/meta/embeddings/cache/group -> CLI/MCP/HTTP/Web.
- Impact note: CRITICAL migration boundary.

Classification:

`missing` for compatibility/generation contract; `wrong` for publication and opaque-consumer behavior.

Allowed next action:

P2 only after accepted P1 shadow output.

Forbidden next action:

Metadata-only cutover, dual-active versions, or partial consumer migration.

### Target and scanner boundary

Current state:

`E:\cheapapp.org` is the accepted real target and owns its `.anvien`. Eight File omissions caused by scanner segment names are confirmed but explicitly excluded by the user.

Required state:

```text
Analyze the target in place.
Keep operational index output under target .anvien.
Keep all plan/report/probe/test artifacts in Anvien.
Preserve exactly the excluded scanner baseline unless a separate plan opens it.
```

Evidence:

- `E0-P0A-BOUNDARY1`: exact target/graph identity.
- `E0-P0A-BOUNDARY2`: user ownership and scanner exclusion.

Relationship and impact:

- Related file count: target graph has 114,125 relationships at baseline.
- Relationship summary: external integration boundary only.
- Impact note: contamination invalidates acceptance.

Classification:

`correct` boundary; scanner behavior `wrong but out of scope`.

Allowed next action:

Read/analyze target in place only during named validation slices.

Forbidden next action:

Copying target, creating a root fixture, writing reports/probes there, or silently adding scanner remediation.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P2-A..P2-F6 and P2-G | reader/version/storage/consumer/publication migration is incomplete | preserve the expanded ordered slice list; each owner boundary commits independently; P2-E2 validates all surfaces; cutover stays last |
| P2-G staged candidate | no staged-v2 release measurements exist before active-manifest CAS | require `E2-P2G-PREBASE1` and all six candidate metrics in `E2-P2G-CANDIDATE1`; block CAS on any absent or over-budget metric without an explicit Owner exception |
| P8 | no implementation exists to close | remain pending |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status.
- [x] Each status has evidence IDs.
- [x] Each target file has relationship count evidence from `file-detail` when applicable.
- [x] Phase Touch Map lists plan-relevant relationship files that can affect the current phase/slice.
- [x] Phase Touch Map defines touch mode for every plan-relevant relationship unit that may be affected.
- [x] Correct parts are marked preserve-only.
- [x] Partial, missing, wrong, unbound, and fake-or-stub parts have exact next actions.
- [x] Blockers are recorded: owner acceptance of the proposed architecture contract is required before production code.
- [x] Next phase status assumptions, next action, and work steps have been updated from this status file.
- [x] Status Refresh Log has an R0 baseline row.
- [x] Implementation has not started; post-slice refresh conditions are not yet applicable.
- [x] Scanner exclusion and target-boundary constraints are reflected in every downstream phase.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [x] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

P0 is complete for planning. The plan is ready for owner review, but no production slice is authorized. The only next permitted slice is P1-A, which must record owner acceptance or changes to the proposed identity/export/external/generation decisions. Scanner remediation remains excluded.
