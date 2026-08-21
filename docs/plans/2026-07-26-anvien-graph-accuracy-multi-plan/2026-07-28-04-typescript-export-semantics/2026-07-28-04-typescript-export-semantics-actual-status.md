# Anvien First-Class TypeScript Export Semantics Actual Status

Title: Anvien First-Class TypeScript Export Semantics
Date: 2026-07-28
Status: P4-A/P4-B/P4-B1/P4-C/P4-C2 are committed at their recorded isolated boundaries; aggregate `E4-PNA-REVIEW1` is `PASS`; Pn-B cleanup is the sole open gate; Pn-C and Child 05 remain locked
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
Predecessor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
Successor child: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`

## Purpose

This file classifies the current TypeScript export-syntax and direct-projection scope before implementation. Historical investigation proves a bounded 21-export defect but does not substitute for fresh current-source, graph, file-detail, impact, and affected-consumer evidence. `E3-PNC-HANDOFF1` now records the accepted Child 03 predecessor closure; it supplies no export implementation evidence and does not waive Child 04 P0.

Detailed proof belongs in the evidence ledger. This file stores classifications, touch modes, plan consequences, and status transitions.

## Freshness / Refresh Rules

- Complete a fresh P0 against the current HEAD before P4-A.
- Refresh affected rows after every accepted slice and before the next slice opens.
- Append refresh-log rows; do not erase the bounded baseline or earlier transitions.
- Use exact evidence IDs. A reserved pending ID is not proof.
- If fresh source evidence changes an owner or affected surface, update the next slice before production code is edited.

## Scope

Target scope:

- TypeScript direct/default/alias/type-only export facts and star/namespace/re-export syntax facts;
- value/type/namespace meaning and separation of module export from access visibility;
- direct-export compatibility derivation, graph projection, and only persistence/read owners proven affected;
- unsupported-syntax diagnostics, negative controls, and bounded `21/21` target validation.

Out of scope:

- terminal export resolution, alias traversal, cycles, ambiguity, and package public API;
- binding and ambient/external semantics;
- unrelated artifact behavior or broad reader policy;
- scanner remediation, blanket transport/cache changes, unrelated refactors, and target-source edits.

## Relationship / Impact Evidence

P0-A resolves the current owners, related-file counts, and blast radii. CRITICAL/HIGH results are scope warnings; each implementation slice must refresh its own exact edit boundary before production changes.

| Unit / File / Surface | Current Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|------------------|-------------------:|----------------------|-------------|
| Current TS export extraction owner | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | `imports.go=17`; `definitions.go=24`; `extract.go=25` | `emitExportStatement` is the syntax owner; `extract.go` wires collection; `definitions.go` preserves named declarations | P4-B/B1 may edit only `imports.go`/`extract.go` after slice-local refresh; `definitions.go` is inspect/preserve-only |
| Current fact/meaning owner | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | `facts.go=245`; `kinds.go=239`; `ir.go=243`; `sort_keys.go=239` | four-file deterministic module-export fact/meaning/collection/order contract | exact P4-A production boundary; CRITICAL/HIGH shared-contract scope |
| Current graph export projection owner | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E4-P4C-IMPACT1`, `E4-P4C-REVIEW1` | `resolution/emit.go=43` | `emitDefinitionNodes` now projects accepted `ScopeIR.Exports` into Export nodes and derives runtime `isExported` | `correct`; exact impact CRITICAL `26/20/5/1`, no terminal/public-API state |
| Current persistence mapping | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1` | `lbugload/csv.go=21`; `lbugschema/schema.go=21` | Ladybug Export CSV/schema/loader and File→Export pair preserve the accepted fact fields; Graph JSON and Ladybug match after lossless representation normalization | `correct`; `11,592` comparisons, differences `0`, orphan references `0` |
| Existing CLI/MCP/HTTP/Web/cache/derived surfaces | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E4-P4C-REVIEW1` | `filecontext/context.go=44`; `mcp/context.go=28` | file-context canonical reader consumes `isExported`; CLI/MCP/HTTP expose derived results; Web/generated contracts remain carriers | `correct` within proven semantic reader scope; transports/carriers remain preserve-only |
| `E:\cheapapp.org` source/worktree | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1` | N/A | sealed source basis and retry pre/post state prove HEAD/branch/seven tracked modifications/three source hashes and four Git-config identities preserved; only normal `.anvien` output changed | `correct` boundary; preserve while Supervisor reviews the semantic finding |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current evidence proves the required behavior. | Preserve; add regression evidence only when needed. |
| `partial` | Some required behavior exists, but the contract is incomplete. | Change only the missing behavior after impact evidence. |
| `wrong` | Accepted evidence proves behavior conflicts with the requirement. | Repair the exact responsible boundary. |
| `missing` | A required behavior or contract does not exist. | Implement the missing responsibility only. |
| `unbound` | A fact exists but is not connected to the real downstream flow. | Bind it at the proven owner. |
| `fake-or-stub` | Placeholder behavior is being presented as real. | Remove or replace it with truthful behavior. |
| `blocked` | Current authority, ownership, or evidence is incomplete. | Do not implement until P0 resolves it. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|-------------------:|----------|--------------------|
| Problem authority | Original report is DRAFT; causal synthesis and Supervisor PASS verify only the bounded defect | findings retained; proposed architecture treated as non-authoritative | `correct` | N/A | `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve evidence hierarchy |
| Child 03 predecessor handoff | Child 03 binding-pattern invariant is accepted and closed at isolated Pn-C commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; non-claims explicitly exclude export semantics | consume only accepted predecessor facts and independently resolve Child 04 owners | `correct` | `1` inbound handoff/roadmap relationship | `E3-PNC-HANDOFF1`, `E3-PNC-COMMIT1`, `E3-PNA-REVIEW1` | predecessor gate closed; no export behavior inferred from Child 03 |
| Current production owner inventory | Four ScopeIR contract owners, TSJS extraction owners, graph/persistence owners, semantic consumers, and preserve-only carriers are current and impact-classified | exact editable/inspect-only/preserve-only owners with current counts | `correct` | `12` P0 file-detail rows; `4` P4-A final-byte rows | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E4-P4A-IMPACT1` | preserve accepted P4-A owner boundary; P4-B requires its own extraction-owner refresh after commit |
| First-class export fact contract | Exact six-file candidate adds one source-site `ExportFact`, seven source-form kinds, three meaning lanes, explicit type-only/provenance/diagnostic state, owning deep copies, and deterministic JSON ordering | one canonical deterministic in-memory export fact/diagnostic contract with no later-slice state | `correct` | `4` production owners + `2` test-after-code files | `E4-P4A-SRC1`, `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1`, `E4-P4A-REVIEW1` | preserve after isolated commit; P4-B may consume but not redesign the contract |
| Bounded direct-export definitions | real-target QA finds all `21/21` definitions and exactly one direct Export fact/relation per row | preserve definitions and attach truthful direct export facts | `correct` fact presence | `21` target Export records / `21` File→Export relations | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1` | preserve fact extraction/projection; do not conflate this PASS with compatibility correctness |
| Bounded export metadata | fresh post-repair QA proves all `21/21` rows, including 15 TypeAliases and six Functions, have correct definition compatibility and FileContext output while TypeAliases remain type-only/type-meaning | 21 of 21 correct direct export facts and compatibility/reader projections on the real target | `correct` | `21/21` overall PASS; `0/15` TypeAlias mismatches | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW2`, `E4-P4C2-DETECT1` | preserve accepted bytes; stage exact boundary and create isolated commit |
| Access visibility versus module export | P4-A provides an independent export fact contract; P4-C derives canonical `isExported` from the same runtime fact without rewriting `visibility` | independent concepts through graph/persistence/reader boundaries | `correct` | `4` production owners + preserve-only siblings | `E4-P4A-SRC1`, `E4-P4A-TEST1`, `E4-P4A-REVIEW1`, `E4-P4C-TEST1`, `E4-P4C-REVIEW1` | preserve separation in P4-C2; no target claim yet |
| General direct/default/alias/type-only syntax | source-less export statements now queue and emit direct/default/local alias/type-only facts with structured diagnostics | one truthful fact per supported exported binding/specifier | `correct` | `4/4` first-class syntax classes | `E0-P0A-SRC1`, `E4-P4B-SRC1`, `E4-P4B-TEST1`, `E4-P4B-REVIEW1`, `E4-P4B-COMMIT1` | preserve committed P4-B boundary; do not reopen |
| Star/namespace/re-export syntax | named/star/namespace/type-only source-bearing forms emit immutable facts; comment-bearing recovered siblings retain `AlsoGood` with one dangling-as diagnostic and two compatibility imports | immutable export syntax facts without terminal resolution | `correct` | `17` related files at imports.go; exact test owner preserved | `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-TEST1`, `E4-P4B1-BOUNDARY1`, `E4-P4B1-REVIEW1`, `E4-P4B1-DETECT1`, `E4-P4B1-COMMIT1` | preserve committed P4-B1; explicit successor authority opens only P4-C under `E4-P4C-AUTH1` |
| Export graph projection | Historical emitter omitted empty provider metadata, so bounded exports appeared unexported; P4-C now emits one Export node per accepted fact and derives runtime compatibility | corrected fact kind/names/meaning/provenance and access separation | `correct` | `414` Export nodes / `414` containment relations | `E0-P0A-VERIFY1`, `E4-P4C-TEST1`, `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1`, `E4-P4C-COMMIT1` | preserve committed projection; P4-C2 is now open |
| Persistence field parity | Graph JSON and Ladybug now preserve the same accepted Export fields; nullable/boolean/list representation is normalized without semantic loss | same affected export facts in every proven persistence owner | `correct` | `2` persistence owners plus generic Graph JSON | `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1`, `E4-P4C-COMMIT1` | preserve committed parity; Child 05 remains locked |
| Terminal module/re-export resolution | Separate bounded barrel defect exists | Child 05 resolves syntax facts to terminal outcomes | `wrong` but out of scope | pending | `E0-P0A-VERIFY1` | preserve boundary; no Child 05 logic here |
| Target boundary | immutable oracle was born durable before analyzer observation; retry pre/post evidence preserves target source/worktree/config, and only normal analyzer-owned `.anvien` output changed | source-only authorship before analyzer observation; durable evidence under `reports/QA/child04-p4c2/...`; normal operational output target-local during QA | `correct` | `19` retry-manifest files; `4` preserved config identities | `E4-P4C2-ORACLE1`, `E4-P4C2-BOUNDARY1` | preserve exact target/Anvien boundaries through Supervisor and repair-or-closure handoff |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | documentation correction against full problem-origin and bounded-verification reports | Child 04 plan authority and scope | removed unrelated campaign assumptions; separated syntax/direct projection from Child 05 resolution; P0 reset to incomplete pending current owner/impact proof | `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | run fresh P0; do not open P4-A yet |
| R1 | 2026-08-20 | Child 03 Pn-B commit `0231d7f9cc7cfd05d8d5da4787161a8ec2cd1550`; excluded graph `1,126/626/0`, `80,908/120,167`; roadmap/Child 03/Child 04 actual/handoff governance rows LOW/non-stale with zero upstream affected files/processes/flows/tests | predecessor handoff and successor opening condition | Child 03 predecessor `dependency-blocked -> accepted handoff`; Child 04 P0 remains incomplete, current owner/file-detail/impact/syntax/consumer evidence remains pending | `E3-PNC-HANDOFF1`, `E3-PNB-COMMIT1` | after Pn-C docs-only commit, open only Child 04 P0-A; no P4 implementation or redundant Supervisor loop |
| R2 | 2026-08-20 | Child 03 Pn-C commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; fresh excluded graph `1,126/626/0`, `80,908/120,167`; full current source plus 12 file-detail and impact rows | Child 04 current owner, syntax, projection, persistence, and consumer boundary | owner inventory `blocked -> correct`; direct syntax `blocked -> missing`; re-export syntax `blocked -> partial`; persistence `blocked -> wrong`; P0-A complete | `E0-P0A-GRAPH1`, `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E0-P0A-STATUS1` | commit exact five-document boundary, then open only P4-A with its updated four-file contract scope |
| R3 | 2026-08-20 | P0-A commit `ff2467bb92f94a9c53c4de030685686700051a98`; exact P4-A candidate; canonical full build and nearest boundary PASS; independent Supervisor report `1B8DEB2F...175573B2`; fresh excluded closure graph `1,130/626/0`, `81,132/120,514`; detect exit `0` | first-class export fact/meaning/diagnostic/deterministic ScopeIR boundary | export fact contract `missing -> correct`; access/export contract separation `wrong -> correct`; direct/re-export extraction and projection classifications unchanged | `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-BUILD1`, `E4-P4A-TEST1`, `E4-P4A-BOUNDARY1`, `E4-P4A-REVIEW1`, `E4-P4A-DETECT1` | create isolated P4-A commit; keep P4-B locked until commit success |
| R4 | 2026-08-21 | P4-B candidate at HEAD `479e8ac229a17f2f6f94be9a4d04e07d74ac4d43`; fresh excluded graph `1,136/626/0`, `81,772/121,285`; Supervisor resubmission PASS | direct/default/alias/type-only extraction plus artifact cleanup | direct syntax `missing -> correct`; bounded metadata `wrong -> partial` because graph projection remains P4-C; P4-B closure is later proven at `11a37aa...` | `E4-P4B-IMPACT1`, `E4-P4B-SRC1`, `E4-P4B-TEST1`, `E4-P4B-BOUNDARY1`, `E4-P4B-REVIEW1`, `E4-P4B-DETECT1`, `E4-P4B-COMMIT1` | preserve committed P4-B; open P4-B1 |
| R5 | 2026-08-21 | P4-B1 candidate at HEAD `ce0e200c55bd96c4374cc6e84bd99a3c82bef641`; final excluded graph `1,146/626/0`, `82,079/121,780`; independent Supervisor REVIEW3 PASS | comment-bearing recovered re-export sibling invariant and full P4-B1 syntax boundary | re-export syntax `partial -> correct`; comment recovery `wrong -> correct`; no terminal state added | `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-BUILD1`, `E4-P4B1-TEST1`, `E4-P4B1-BOUNDARY1`, `E4-P4B1-REVIEW1` | final detect exit `0` with `305` changed units / `7` files / `1` affected process / MEDIUM; implementation commit `42d167aaf28446ac0b3de479a8afefabb8d06736` closes P4-B1; keep P4-C/P4-C2/Child 05 locked |
| R6 | 2026-08-21 | Accepted P4-B1 implementation commit `42d167aaf28446ac0b3de479a8afefabb8d06736`, handoff HEAD `a12e0ccb77bda7da8aad2ec29b8050b55f81bc08`, clean worktree/index | Main-owned P4-B1 closure and authority transfer | `E4-P4B1-DETECT1` and `E4-P4B1-COMMIT1` transition `pending -> recorded/closed`; P4-C/P4-C2/Child 05 remain locked | `E4-P4B1-DETECT1`, `E4-P4B1-COMMIT1` | preserve the accepted syntax boundary; do not open P4-C without explicit authority |
| R7 | 2026-08-21 | Successor-verified clean HEAD `871189b8c6a4e4bb9ff538407232c913b8cf4db6`; P4-B1 source/test unchanged; fresh excluded graph `1,147/626/0`, `82,087/121,788` | P4-C opening authority and planner synchronization | P4-C `locked -> open`; P4-C2/Child 05/target remain locked | `E4-P4C-AUTH1` | open one visible Coder lane for P4-C; require fresh exact production file-detail/impact before edits |
| R8 | 2026-08-21 | P4-C candidate at HEAD `871189b8c6a4e4bb9ff538407232c913b8cf4db6`; fresh graph `1,855/735/0`, `113,496/156,003`; Supervisor REVIEW1 report `15,261` bytes / `101` LF / `AA94C417A02BE371168D8ADCD5B3F4FECF1941F382518EFF9214EC0FABB93CFF` | independent P4-C source/build/boundary/review closure | graph projection and persistence parity `wrong -> correct`; Supervisor `READY_FOR_SUPERVISOR -> PASS`; P4-C2/Child 05/target remain locked; detect/commit pending | `E4-P4C-IMPACT1`, `E4-P4C-SRC1`, `E4-P4C-BUILD1`, `E4-P4C-TEST1`, `E4-P4C-BOUNDARY1`, `E4-P4C-REVIEW1` | run fresh `analyze --force`, then `detect-changes`, stage exact P4-C boundary and commit; record the review-induced empty `.tmp\\p4c-tests` and Supervisor `/root/authority_scan` deviation as provenance only |
| R9 | 2026-08-21 | planner-refresh basis; fresh graph `1,857/735/0`, `113,523/156,030`; detect exit `0` with `180` changed units / `16` files / `14` affected files / HIGH | Main-owned change detection after Supervisor PASS | `E4-P4C-DETECT1` transitions `pending -> recorded`; P4-C remains uncommitted; P4-C2/Child 05/target remain locked | `E4-P4C-DETECT1` | stage only the exact candidate + five ledgers + verified reports/provenance, then commit P4-C |
| R10 | 2026-08-21 | isolated P4-C commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877`; post-commit HEAD matches; `git diff --check` PASS; two older handoffs preserved untracked | P4-C closure and successor opening | `E4-P4C-COMMIT1` transitions `pending -> recorded/closed`; P4-C `open -> committed`; P4-C2 becomes sole open slice; Child 05 remains locked | `E4-P4C-COMMIT1` | create one visible P4-C2 validation lane with independent 21-entry oracle; do not access target before lane pre-state gate |
| R11 | 2026-08-21 | P4-C commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877`; Owner lifecycle correction; Architect `READY_FOR_PLANNER` at `1E7EEB6D...BDEEA` | P4-C2 oracle lifecycle only | `.tmp` capture `accepted/lost -> invalid-lifecycle debug`; source-only Oracle Authoring authorized before analyzer/QA output; existing evidence IDs remain pending | Architect report; Planner handoff | route Oracle Authoring now, then existing QA; no recovery/audit loop or new gate |
| R12 | 2026-08-21 | HEAD `e32a412b289453a530bc71b93320ef2b97b3a97a`; sealed oracle digest `7749AB14...5439`; retry report `C831004F...FDB4`; run digest `9F414A2C...ABC4` | P4-C2 durable oracle, one target retry, complete comparison, and target/config boundary | `E4-P4C2-ORACLE1 pending -> SEALED`; `E4-P4C2-TARGET1/BOUNDARY1 pending -> evidence-complete with finding`; bounded fact presence `partial -> correct`; bounded compatibility/reader state `partial -> wrong` for 15 TypeAlias rows; target boundary `partial -> correct` | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1` | exactly one independent Supervisor reviews the bounded finding; P4-C2 remains open and Child 05 locked |
| R13 | 2026-08-21 | Supervisor report `8E37F4B1...81F7`; current HEAD `6b93e80601b0549f1b3e56bd6b68b07b9bc9680a`; unrelated concurrent orchestration-skill commit/edit does not touch P4-C2 | independent P4-C2 acceptance and rejection routing | `E4-P4C2-REVIEW1 pending -> REJECT`; accepted oracle/fact-presence/negatives/parity/boundary remain closed; 15-TypeAlias compatibility/reader state remains `wrong` | `E4-P4C2-REVIEW1` | authorize one Coder lane for only the rejected `emit.go` compatibility invariant after fresh impact; preserve Child 05 lock |
| R14 | 2026-08-21 | HEAD `310502a88849fe75f86a45a987ba21490d19dbe2`; exact two-file repair; fresh QA run digest `5CA04508...33E`; REVIEW2 canonical SHA `5B99A74B...8DC0B` | P4-C2 rejection repair, fresh target validation and independent re-review | 15-TypeAlias compatibility/reader `wrong -> correct`; `E4-P4C2-REVIEW2 -> PASS`; Oracle/Functions/negatives/parity/boundary remain closed | `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW2` | keep P4-C2 open only for fresh self graph, detect-changes and isolated commit; Child 05 locked |
| R15 | 2026-08-21 | HEAD `310502a88849fe75f86a45a987ba21490d19dbe2`; fresh self graph `1,939/736/0`, `114,628/157,443`; detect exit `0` | Main-owned P4-C2 pre-commit change detection | `E4-P4C2-DETECT1 pending -> recorded/PASS`; exact seven tracked changed/affected paths; persisted health remains `0/0/0` | `E4-P4C2-DETECT1` | stage only the exact accepted P4-C2 manifest and create `E4-P4C2-COMMIT1`; Child 05 remains locked |
| R16 | 2026-08-21 | isolated P4-C2 commit `03f09b43f652b9a14b3e49774dc805c0dfd24a27`; parent `310502a88849fe75f86a45a987ba21490d19dbe2`; exact 89 paths; post-commit index/tracked worktree clean | P4-C2 closure and aggregate governance opening | `E4-P4C2-COMMIT1 pending -> recorded/closed`; P4-C2 `open -> committed`; aggregate Pn-A `locked -> open`; Pn-B/Pn-C/Child 05 remain locked | `E4-P4C2-COMMIT1` | open exactly one visible aggregate Supervisor lane for Pn-A; do not reopen slice gates or enter cleanup/Child 05 |
| R17 | 2026-08-21 | HEAD `a2e0c4ab7654f42c6a3c69402cf4c6b63bbb0bdd`; independent aggregate report canonical SHA `7EBFD508...DC5A8`; fresh review graph `1,939/736/0`, `114,630/157,445` | complete Child 04 aggregate acceptance | `E4-PNA-REVIEW1 pending -> PASS`; all slice/source/artifact/target invariants closed; residual same-invariant surfaces none; Pn-B `locked -> open` | `E4-PNA-REVIEW1` | open one visible cleanup executor for Pn-B; preserve all accepted artifacts and keep Pn-C/Child 05 locked |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|----------------------------|-----------|------------|----------|------------|
| `internal/scopeir/{facts.go,kinds.go,ir.go,sort_keys.go}` | source of deterministic module-export syntax truth | P4-A | accepted production boundary; preserve after isolated commit | `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-REVIEW1` | access visibility remains independent; later slices consume without redesign |
| `internal/providers/tsjs/{imports.go,extract.go}` | produces direct/default/alias/type-only facts | P4-B | accepted edit boundary; preserve after isolated commit | `E4-P4B-IMPACT1`, `E4-P4B-SRC1`, `E4-P4B-REVIEW1` | preserve definitions, visibility, and non-export behavior; no P4-B1/P4-C state |
| `internal/providers/tsjs/{imports.go,extract.go}` | produces named/star/namespace/type-only re-export syntax facts | P4-B1 | accepted edit boundary; preserve after isolated commit | `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-REVIEW1` | only imports.go changed in production; extract.go preserve-only; no terminal resolution |
| `internal/resolution/emit.go`, `internal/lbugload/csv.go`, `internal/lbugschema/schema.go`, proven semantic consumers | project/persist/consume exact export facts | P4-C | inspect now; edit only accepted changed-field consumers after fresh impact | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1` | one source fact; only named affected consumers |
| CLI/MCP/HTTP/Web/generated carriers and unaffected caches | transport or non-consuming surfaces | all P4 | preserve-only unless slice evidence proves direct semantic consumption | `E0-P0A-SRC1` | plan update required before expansion |
| Child 05 resolver owners | consume immutable syntax facts | all P4 | preserve-only | successor boundary | no traversal/cycle/ambiguity work here |
| target source/worktree | real integration input | P4-C2 | Oracle Authoring and QA gates complete; preserve-only while Supervisor consumes durable evidence read-only | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1` | no copied source, target report/fixture/probe, `.tmp` evidence, or write outside the already-recorded normal analyzer-owned `.anvien` output |

## Detailed Findings

### Bounded export-metadata omission

Current evidence:

- The causal synthesis records 21 selected top-level exported TypeScript declarations that survived as definitions.
- At the accepted baseline, none carried the export/visibility metadata expected by graph consumers.
- The first confirmed divergence was provider fact creation; downstream projection serialized only what it received.

Required state:

```text
export syntax site
  -> one immutable export fact with meaning and provenance
  -> direct export projection derived from that fact
  -> affected persistence readers observe the same value

access visibility remains independent
```

Classification: `wrong` for the bounded baseline and current source; P0-A resolves the ownership boundary.

Allowed next action: commit the exact accepted P4-A boundary; only after commit success may P4-B open.

Forbidden next action: fill one field solely to improve counts, treat re-export reachability as direct export, or add transport adapters without affected-field evidence.

### Child 04 versus Child 05 boundary

Current evidence:

- The same causal synthesis identifies barrel lookup as a separate first divergence from export metadata.
- Syntax collection and terminal export resolution have different commands/state/ownership: Child 04 records what source exports; Child 05 determines what terminal symbol a consumer import reaches.

Classification: boundary `correct` after this documentation correction.

Allowed next action: make Child 04 facts complete and immutable enough for Child 05.

Forbidden next action: recurse through barrels, select candidates, or claim public API during Child 04.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Next-Action Update |
|-----------|-----------------------|-----------------------------|
| Predecessor handoff | Child 03 final evidence/benchmark/target boundary is accepted at `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`; no export claim is transferred | `correct`; predecessor gate closed |
| P4-A | first-class contract, tests, build, boundary, independent Supervisor acceptance, closure detect, and isolated commit `479e8ac229a17f2f6f94be9a4d04e07d74ac4d43` are complete | `correct`; preserve accepted contract |
| P4-B | direct/default/local alias/type-only facts, diagnostics, visibility separation, cleanup, detect, and isolated commit are closed at `11a37aa8ec0320dd93258c058b088d1070aa778d` | `correct`; preserve and do not reopen |
| P4-B1 | named/star/namespace/type-only facts and comment-bearing recovered sibling behavior pass independent Supervisor REVIEW3; terminal state remains absent | preserve the closed commit `42d167aaf28446ac0b3de479a8afefabb8d06736`; successor verification found no invalidation |
| P4-C | graph projection, Ladybug persistence, and semantic consumer dialects are accepted and committed | `correct`; preserve P4-C at `c99c4070b66e7a96be8c9fa2721a0335a1f94877` |
| P4-C2 | durable oracle remains sealed; exact repair, fresh `21/21 + 11/11` target QA, independent REVIEW2, Main-owned detect and isolated commit are closed | `correct`; preserve commit `03f09b43f652b9a14b3e49774dc805c0dfd24a27`; open only aggregate Pn-A |
| Pn-A | aggregate Supervisor independently clears all five Child 04 slices, current production blobs, evidence lifecycle, target boundary and rejection history | `correct`; preserve report `7EBFD508...DC5A8`; open only Pn-B cleanup |
| Pn-B | no cleanup executor or cleanup Supervisor verdict yet | inventory only Child 04-created artifacts; remove only proven dead work; preserve current evidence |

## Implementation Gate

- [x] Current graph and HEAD basis are recorded.
- [x] Current production source and existing tests for export facts/extraction/projection/persistence are read in full.
- [x] Every candidate editable file has fresh `file-detail` related-file count evidence.
- [x] Every candidate function/method/exported symbol has fresh upstream impact evidence and reported blast radius.
- [x] Current syntax coverage and unsupported paths are inventoried without inferring from report proposals.
- [x] Every actual consumer of a changed export field is classified; unaffected consumers are preserve-only.
- [x] Child 04 syntax/direct-state ownership is separated from Child 05 terminal-resolution state.
- [x] Exact production, test, fixture, generated-output, and validation ownership is recorded for P4-A.
- [x] P4-A work steps match the discovered current owners.
- [x] Status Refresh Log contains the completed P0 refresh.
- [x] Child 03 predecessor handoff `E3-PNC-HANDOFF1` is recorded without selecting a Child 04 implementation owner.
- [x] P4-B affected status rows and next-phase decisions are refreshed from current source, cleanup, detect, commit, and Supervisor PASS evidence.
- [x] P4-B1 source, build, boundary, cleanup, Supervisor PASS, Main-owned detect, and isolated commit evidence is recorded and preserved.
- [x] P4-C source, build, boundary, independent Supervisor REVIEW1, Main-owned detect, and isolated commit evidence is recorded and closed.
- [x] Main-owned `E4-P4C-DETECT1` is recorded from a fresh graph; exact semantic changed-path set and explicit report staging boundary are known.
- [x] `E4-P4C-COMMIT1` is recorded at isolated commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877`; post-commit HEAD/worktree boundary is verified; P4-C2 is the sole open slice and Child 05 remains locked.
- [x] `E4-P4C2-ORACLE1` records the durable-from-creation sealed source-only 21+11 bundle with no forbidden input, target write, or evidence-bearing `.tmp` artifact.
- [x] Existing QA resumed only after seal and recorded `E4-P4C2-TARGET1` plus `E4-P4C2-BOUNDARY1` against the same source basis; the evidence-complete result contains a bounded semantic finding and is not acceptance.
- [x] Main-owned `E4-P4C2-DETECT1` is recorded from one fresh self graph with exact seven-path semantics, HIGH changed-file scope warning, LOW overall risk, zero affected flows, and zero persisted health regression.
- [x] `E4-P4C2-COMMIT1` is recorded at isolated commit `03f09b43f652b9a14b3e49774dc805c0dfd24a27`; exact 89-path boundary and clean post-commit tracked/index state are verified.

## Final P0 Decision

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The bounded `0/21` baseline remains historical authority. P4-A/P4-B/P4-B1/P4-C/P4-C2 retain their accepted commits and aggregate `E4-PNA-REVIEW1` is `PASS`. The first uncompleted gate is `E4-PNB-CLEAN1`; open exactly one visible cleanup executor, then a distinct cleanup Supervisor. Pn-C, Child 05 and later slices remain locked.
