# Anvien Ambient And External Resolution Actual Status

Title: Anvien Ambient And External Resolution
Date: 2026-07-28
Status: P0 Complete / P6-A Committed / P6-B Committed At 5c1584fef7153a7a331c8fedd1ce64176ddc873d / P6-C1 Committed At 8055f0a6860721e26462572e34469e0d708d4a52 / P6-C2 Independent Supervisor PASS And Planner-finalized / Isolated Commit Pending
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`

## Purpose

This file records the current declaration inputs, resolver workspace, standard-library/external lookup, outcome, graph-health, target, and affected-reader state before Child 06 work.

P6-A must turn unknown design questions into a source-backed decision before production implementation. Detailed proof belongs in the evidence ledger.

## Freshness / Refresh Rules

Refresh this file:

- after Child 05 changes the repository resolution result consumed by Child 06;
- during P6-A when source/config/runtime evidence selects the declaration-authority design;
- before every production slice after fresh graph, file-detail, and impact evidence;
- after each accepted slice and whenever affected readers or owner files change.

Use explicit transitions and append refresh-log rows. Do not preserve superseded assumptions as active authority.

## Scope

Target scope:

- source-backed declaration-universe behavior and implementation decision;
- TypeScript standard-library authority selected by P6-A;
- project/package declaration lookup only when P6-A proves it required;
- necessary external target representation and structured outcomes;
- graph-health projection from resolver outcomes;
- exact target validation for `Promise`, `Math.max`, and `Math.min`;
- only persistence/readers proven affected by Child 02 inventory and fresh impact.

Out of scope:

- preselected authority/storage/runtime architecture;
- unevidenced declaration-source coverage or public traversal options;
- Child 05 module/export resolution;
- graph-output behavior, scanner behavior, target-source edits, and unrelated language semantics.

## Relationship / Impact Evidence

The graph used for this P0 inspection reported current commit parity and `stale=false` for every listed file-detail result.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|-------------------:|----------------------|-------------|
| `internal/resolution/indexes.go` | `E0-P0A-FD1` | 46 | workspace built from repository ScopeIR, name/type/member/import indexes | high file risk; `buildWorkspace` CRITICAL |
| `internal/resolution/resolve.go` | `E0-P0A-FD2` | 50 | call/member/type resolution and unresolved emission | high file risk; call/type symbols CRITICAL |
| `internal/graphhealth/diagnostics.go` | `E0-P0A-FD3` | 29 | diagnostic normalization and target-text classification | high file risk; classifier HIGH |

| Symbol | Impact Evidence | Risk | Impacted Symbols | Affected Files | Modules | Processes | Linked Flows / Tests |
|--------|-----------------|------|-----------------:|---------------:|--------:|----------:|----------------------:|
| `buildWorkspace` | `E0-P0A-IMPACT1` | CRITICAL | 8 | 6 | 5 | 28 | 29 / 23 from containing file |
| `resolveCall` | `E0-P0A-IMPACT2` | CRITICAL | 6 | 4 | 4 | 35 | 21 / 26 from containing file |
| `resolveTypeAnnotation` | `E0-P0A-IMPACT3` | CRITICAL | 6 | 4 | 4 | 35 | 21 / 26 from containing file |
| `classifyDiagnostic` | `E0-P0A-IMPACT4` | HIGH | 7 | 2 | 3 | 3 | 8 / 14 from containing file |

Fresh P6-A graph basis is HEAD `ec765debff335540c77d409ebb2c9f45e4a0a77d`, analyzed with `anvien analyze --force`: `1,985` scanned, `743` parsed, `0` failed, `116,467` nodes, `160,591` relationships, `stale=false`.

| P6-A owner / consumer | Related files | Fresh impact warning |
|-----------------------|--------------:|----------------------|
| `internal/analyze/analyze.go` / `analyze.Run` | 187 | CRITICAL `24` symbols / `9` files / `8` modules / `23` processes |
| `internal/resolution/types.go` | 34 | options/outcome shared owner; exact edited symbol must be refreshed in P6-B/C3 |
| `internal/resolution/indexes.go` / `buildWorkspace` | 58 | CRITICAL `59 / 24 / 8 / 23` |
| `internal/resolution/resolve.go` / call-access-type owners | 60 | each CRITICAL `28 / 11 / 7 / 32` |
| `internal/resolution/emit.go` / `emitUnresolvedReference` | 50 | CRITICAL `9 / 4 / 2 / 34` |
| `internal/graph/types.go` / `graph.Node` | 256 | CRITICAL `1,717 / 273 / 48 / 428`; validate/preserve-only unless later evidence requires edit |
| `internal/scopeir/kinds.go` / `scopeir.NodeLabel` | 250 | CRITICAL `1,846 / 291 / 47 / 422` |
| Ladybug CSV / schema | 23 / 22 | `nodeCSVRow` CRITICAL `25 / 11 / 1 / 12`; `NodeSchema` MEDIUM `8 / 6 / 2 / 0` |
| graph-health policy / diagnostics / ResolutionGap inputs | 31 / 30 / 41 | Diagnostic CRITICAL `70 / 17 / 5 / 77`; classifier HIGH `10 / 4 / 3 / 0`; inputs CRITICAL `13 / 7 / 1 / 15` |
| processes | 14 | `Apply` CRITICAL `29 / 8 / 7 / 31`; call graph CRITICAL `28 / 8 / 6 / 25` |
| MCP context / impact / rename | 28 / 41 / 9 | payload/collector CRITICAL `1 / 1 / 1 / 7`, `11 / 5 / 2 / 17`, `1 / 1 / 1 / 8` |
| HTTP graph / filecontext | 22 / 46 | generic transport/fallback; validate/preserve-only pending later impact |

HIGH/CRITICAL is a blast-radius warning requiring narrow edits and affected-boundary regression, not an edit ban.

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current behavior satisfies this child's bounded requirement. | Preserve and regress. |
| `partial` | A relevant input/behavior exists but the accepted contract is incomplete. | Change only the proved gap. |
| `wrong` | Current behavior conflicts with the accepted requirement. | Replace at the proved owner. |
| `missing` | Required behavior does not exist. | Add it after owner/design evidence. |
| `unbound` | A result exists but is not connected to the truthful downstream consumer. | Bind only that boundary. |
| `blocked` | A predecessor, design decision, or impact inventory is missing. | Do not implement until refreshed. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|-------------------:|----------|--------------------|
| Problem authority | originating report records the false external gaps but labels its architecture DRAFT; bounded reports verify C5 only | findings/acceptance separated from design; current source decides implementation | correct | N/A | `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2` | preserve evidence hierarchy |
| Child 05 handoff | exact four-ledger closure commit `ec765debff335540c77d409ebb2c9f45e4a0a77d` remains an ancestor of current external/user-owned HEAD `fa351c60617212635ef57a43b85d7449ef1eea1c`; P5-A/P5-B/P5-C/P5-D and Pn-A/Pn-B are accepted and committed | immutable repository module/file, syntax-derived export-table, deterministic traversal, fail-closed explicit-import, and generic-first proof contract available before external lookup | correct | exact closure commit; four living ledgers | `E0-P0A-DEPEND1`, `E6-P6A-DEPEND1` | consume as immutable predecessor; preserve both external orchestration-only commits outside P6-B |
| P6-A declaration-universe decision | offline generated/check-in/embedded TypeScript `5.9.3` catalog, narrow root config profile, final outcome contract, owner/reader map, P6-C1 preserve-only, and P6-C2 active are recorded in an immutable report; final review passed and Main committed the exact seven-path manifest as `b98131e44932a7bcac17b487ecb2914535927d01` | accepted and committed decision with no residual same-invariant surface | correct / committed | report `25,326` bytes / `289` LF / SHA-256 `77D5E9AC8D76D98C76D1816C8D6E69265D4AFB30367E3DA50DF3EAA3445D7BA2`; final review `6,011` bytes / `83` LF / SHA-256 `39CE249E9D13F1C77FE6F61DD6E9B1D2E4000B004CE3EBBF461C762F3FA28384`; commit `b98131e44932a7bcac17b487ecb2914535927d01` | `E6-P6A-DECISION1`, `E6-P6A-REVIEW1`, `E6-P6A-COMMIT1` | preserve decision; execute only P6-B after fresh exact owner preflight |
| Resolver workspace | repository ScopeIR/P5 indexes remain first; explicit-import and local/repository/import receiver claims are terminal before external fallback; handled catalog results survive as immutable site records; final validation enforces the exact proof-state/status/reason/hash matrix | repository/local/import receiver claims are terminal and every handled external lookup remains losslessly observable without contradictory ready/rejected proof | correct / committed | validator impact is CRITICAL; direct matrix passes `25` named subrows and six failure classes preserve exact site/counter equality, identical replay, and conflict rejection | `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-TEST1`, `E6-P6B-REVIEW1`, `E6-P6B-COMMIT1` | preserve accepted P6-B behavior; P6-C1 preserve-only is open and P6-C2/C3/D remain locked |
| Supported TypeScript config inputs | zero/one root JSONC `tsconfig.json`; only `target`, `lib`, `noLib` affect the profile; absent config uses manifest default; invalid/unreadable/unsupported topology fails closed | exact selected profile without scanner-corpus change or ownership guess | correct / independent Supervisor PASS | profile owner current impact CRITICAL `36` symbols / `11` files; direct fixtures cover ready/unavailable paths | `E6-P6B-SRC1`, `E6-P6B-PROVENANCE1`, `E6-P6B-TEST1`, `E6-P6B-REVIEW1` | preserve accepted config boundary; do not broaden scope |
| TypeScript standard-library authority | offline exact TS `5.9.3` generator -> checked-in catalog -> `go:embed`; two-phase identity commits every P6-A semantic component and keeps artifact hash outside the identity payload; loader reason remains typed through lookup; ready proof accepts only legitimate ready-catalog outcomes | general declaration-backed lookup with unique canonical semantic identity and lossless fail-closed validation selected by P6-A | correct / committed | catalog `2,030` symbols / `11,802` direct member rows / `14,587` semantic IDs; `14,587` unique, `0` duplicate/format mismatch; artifact `2,003,050` bytes; direct matrix `25/25` subrows PASS | `E6-P6B-SRC1`, `E6-P6B-PROVENANCE1`, `E6-P6B-TEST1`, `E6-P6B-REVIEW1`, `E6-P6B-COMMIT1` | preserve accepted provenance/profile/package behavior |
| P6-B site-level authority results | every handled global/type/member lookup is retained with stable site, request/status/reason, target/owner, declaration, and explicit `ready`/`missing`/`rejected` authority/catalog/profile/config proof semantics; ready identity cannot pair with catalog missing/rejection reasons | one lossless per-source-site record for resolved and unavailable lookup results, including every catalog-validation failure class and a non-contradictory proof matrix | correct / committed | direct six-class matrix emits exactly `3` unique sites/class; built `analyze.Run` emits exactly `7` named stable sites/class; direct validator adds `7` ready positives, `6` cross-status negatives, `6` missing/rejected positives, and `6` ready-failure negatives; counters equal records, replay identical, conflicts fail closed | `E6-P6B-SRC1`, `E6-P6B-TEST1`, `E6-P6B-REVIEW1`, `E6-P6B-COMMIT1` | preserve P6-B result carriage; do not open C2 materialization or C3 final shared DTO before ordered gates |
| External member receiver eligibility | call/access external lookup runs only for an unclaimed receiver that resolves through the external authority; local/repository/import claims are terminal even when the member is missing | only a resolved external receiver may reach external member lookup; repository/local/import claims remain terminal | correct / independent Supervisor PASS | `resolveCall` and `resolveAccess` remain CRITICAL; all local/repository/import/external call/access regressions PASS | `E6-P6B-SRC1`, `E6-P6B-TEST1`, `E6-P6B-REVIEW1` | preserve accepted behavior |
| Exact compiler-vector proof | table-driven durable matrix names and checks exactly all ten P6-A compiler vectors, including every value/type/member expectation | one exact durable fixture/result row per P6-A vector with value/type/member outcomes | correct / independent Supervisor PASS | `10/10` exact named rows PASS | `E6-P6A-ORACLE1`, `E6-P6B-TEST1`, `E6-P6B-REVIEW1` | preserve exact denominator/result; no category-overlap substitution |
| Project/package declaration lookup | current product has no project/package `.d.ts`, ambient-module, augmentation, or `node_modules` authority; the accepted runtime has zero matching implementation/test owners, scanner policy excludes `.d.ts` and `node_modules`, and fresh graph output contains neither file-path class | preserve unavailable behavior; no production/test owner and no invented lookup | correct / preserve-only / committed | `0` production owners / `0` exact P6-C1 test owners / `2` tracked frontend-tooling `.d.ts` files but `0` indexed `.d.ts` nodes / `0` indexed `node_modules` nodes; commit `8055f0a6860721e26462572e34469e0d708d4a52` | `E6-P6C1-IMPACT1`, `E6-P6C1-SCOPE1`, `E6-P6C1-SRC1`, `E6-P6C1-TEST1`, `E6-P6C1-REVIEW1`, `E6-P6C1-COMMIT1` | preserve closed result; P6-C2 is open; any future P6-C1 activation requires planner refresh |
| External target representation | final resolved P6-B sites materialize deterministic referenced-only `ExternalSymbol` nodes/edges with semantic owner, TS/catalog/profile/config/library/range/site proof; unavailable/excluded/mismatch rows create no node | referenced-only deterministic `ExternalSymbol`; unavailable outcomes stay gaps; affected readers adapt without repository ownership pollution | correct / independent Supervisor PASS / planner-finalized | packaged fixture `3` external nodes / `4` coalesced edges / `6` site proofs; `0` File/DEFINES/IMPORTS/CONTAINS ownership; two runs byte-equal; residual same-invariant blockers none | `E6-P6C2-IMPACT1`, `E6-P6C2-DESIGN1`, `E6-P6C2-SRC1`, `E6-P6C2-TEST1`, `E6-P6C2-PARITY1`, `E6-P6C2-REVIEW1` | run one current Main analyze/detect and commit exact accepted 23-path manifest; keep P6-C3/D/target locked until commit |
| Structured resolution outcomes | resolver emits resolved references or generic unresolved diagnostics; no accepted external capability result exists | one structured target/reason/stage/proof per affected source site | missing | 50 related files at resolver owner | `E0-P0A-SRC2`, `E0-P0A-FD2` | P6-C3 after lookup/representation |
| Graph-health classification | `classifyDiagnostic` infers from target text using Go-oriented builtin/stdlib/external qualifier tables | mechanical projection from resolver outcome | wrong | 29 related files; exact symbol affects 2 files | `E0-P0A-SRC4`, `E0-P0A-FD3`, `E0-P0A-IMPACT4` | P6-D after P6-C3 |
| Bounded target sites | TypeScript oracle resolves all three; Anvien records gaps and later labels them as in-repository/analyzer issues | correct external or accepted external-capability outcome for all three; no in-repository analyzer gap | wrong | 3 exact sites | `E0-P0A-VERIFY1`, `E0-P0A-TARGET1` | terminal P6-D acceptance |
| Affected persistence/readers | graph JSON remains generic; Ladybug has fixed ExternalSymbol table/pairs and native readback; processes exclude external CALLS endpoints; context/impact expose provenance/non-editable semantics; rename rejects before edit collection; HTTP/Web/file-detail remain unchanged | edit only the named semantic consumers after fresh slice impact; preserve generic transports unless invalidated | P6-C2 active rows correct / independently accepted; P6-D health row still pending | native Ladybug `2` node COPY / `1` relation COPY / `0` fallback/skips; built context/impact/rename PASS; external process members `0` | `E6-P6C2-IMPACT1`, `E6-P6C2-SRC1`, `E6-P6C2-TEST1`, `E6-P6C2-PARITY1`, `E6-P6C2-REVIEW1` | commit P6-C2; keep graph-health and shared final DTO locked |
| Target boundary | target analyzed in place; source/worktree are not implementation surfaces | preserve target source; regenerate only normal target index during P6-D | correct | accepted target graph: 114,125 relationships | `E0-P0A-BOUNDARY1` | preserve until P6-D pre/post evidence |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | HEAD `238aec06d28286acc0fca0f4e6a69f9eb4ff6a49`; graph file-detail `stale=false`, analyzed `2026-08-09T19:19:54Z` | Child 06 source/report/ledger reset | removed preselected authority, source-category, public-option, and adapter assumptions; P6-A becomes mandatory design slice; P6-C1/C2 conditional on evidence | `E0-P0A-GRAPH1`, `E0-P0A-SRC1..E0-P0A-SRC4`, `E0-P0A-FD1..E0-P0A-FD3`, `E0-P0A-IMPACT1..E0-P0A-IMPACT4` | open P6-A only after Child 05 handoff; no production authority mechanism selected yet |
| R1 | 2026-08-22 | HEAD and exact Child 05 closure commit `ec765debff335540c77d409ebb2c9f45e4a0a77d`; parent/Pn-B `fcc44334c0f75b3b19046dc8f9f4de40eb459fa9`; Pn-A `b68e738d64eebea65a045afbf0b12d94dd43cbf4` | predecessor closure and Child 06 opening boundary | Child 05 dependency `blocked -> correct`; P6-A `locked -> open`; all production/test/target slices remain locked | `E6-P6A-DEPEND1` | execute only source/config/runtime/consumer/oracle inventory and P6-A decision documents; no production/test/target action |
| R2 | 2026-08-22 | same HEAD; fresh graph `1,985` scanned / `743` parsed / `0` failed / `116,467` nodes / `160,591` relationships / `stale=false`; immutable decision report SHA-256 `77D5E9AC8D76D98C76D1816C8D6E69265D4AFB30367E3DA50DF3EAA3445D7BA2` | P6-A source/config/package/consumer/oracle decision only | authority/config/consumer design `blocked -> decision-recorded`; P6-C1 `conditional -> preserve-only`; P6-C2 `conditional -> active`; P6-A review/commit and every production/test/target slice remain locked; `03:30 +07` report deadline missed; contract `18+/0-` scope deviation disclosed and frozen | `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-ORACLE1`, `E6-P6A-CONSUMER1`, `E6-P6A-DECISION1`, `E6-P6A-BUILD1` | route immutable candidate to Supervisor; Main decides late-handoff/deviation acceptance; no production action |
| R3 | 2026-08-22 | initial Supervisor REJECT `5679FB24...6290`; contract restored byte-for-byte to HEAD blob `2020b479...`; reject-only Supervisor PASS `39CE249E...8384`; fresh Main planner graph `116,521 / 160,645` | sole sealed-boundary blocker repair and P6-A acceptance/commit transition | contract `18+/0- -> zero diff`; unchanged four-ledger/report technical candidate `REJECT -> PASS`; deadline miss remains disclosed/non-blocking; P6-A `decision-recorded -> Supervisor PASS`; P6-B remains production-locked until exact commit | `E6-P6A-REVIEW1`, `E6-P6A-COMMIT1` | commit exact seven-path P6-A decision manifest immediately, then hand the resulting commit anchor directly to P6-B |
| R4 | 2026-08-22 | HEAD `b98131e44932a7bcac17b487ecb2914535927d01`, parent `ec765debff335540c77d409ebb2c9f45e4a0a77d`; fresh P6-B graph `1,990` scanned / `743` parsed / `0` failed / `116,535` nodes / `160,659` relationships; official corrected protected Main handoff inventory `18` | exact P6-A commit closure and P6-B opening preflight | P6-A `Supervisor PASS -> committed`; P6-B `locked -> implementation-ready`; production/test/config/catalog remain unchanged; exact owner file-detail/symbol impact and invariant/verify matrix recorded | `E6-P6A-COMMIT1`, `E6-P6B-IMPACT1` | implement only the isolated P6-B catalog/profile/lookup family; keep P6-C1/C2/C3/D locked |
| R5 | 2026-08-22 | dirty candidate on anchor HEAD `b98131e44932a7bcac17b487ecb2914535927d01`; final source/cleanup graph `2,006` scanned / `750` parsed code / `0` failed / `111,524` nodes / `156,259` relationships; official protected Main handoff inventory `18` after the `05:24` transfer | P6-B catalog/profile/resolver/analyze behavior, direct tests/fixtures, build/package/runtime/benchmark/detect evidence, and exact dead-debug cleanup | standard-library authority/config `missing -> implemented + coder-validated`; P6-B `implementation-ready -> READY_FOR_INDEPENDENT_SUPERVISOR`; P6-C1/C2/C3/D remain locked; no stage/commit/target access; all 18 Main handoffs remain pathname-only protected state | `E6-P6B-SRC1`, `E6-P6B-PROVENANCE1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-DETECT1` | independent Supervisor reviews exact diff/evidence; Main alone may commit after PASS |
| R6 | 2026-08-22 | same anchor HEAD; Supervisor graph `2,007/750/0`, `111,540` nodes / `156,275` relationships; current detect CRITICAL `23/8/148/8`; report SHA-256 `26FCA7B7678980F2B129DCF8EA3DB6345FDD269C260FAF8C5326F1B203416FB9`; protected handoffs `19` | independent P6-B source/invariant review and report-only seal | P6-B `READY_FOR_INDEPENDENT_SUPERVISOR -> REJECT`; profile/provenance/package/cleanup cleared; site retention `missing`, identity/member precedence `wrong`, exact ten-vector proof `partial`; P6-C1/C2/C3/D remain locked | `E6-P6B-REVIEW1` | reuse only the existing P6-B coder for the exact four blockers; refresh impact before edits and resubmit same slice |
| R7 | 2026-08-22 | implementation base `b98131e44932a7bcac17b487ecb2914535927d01`; authoritative HEAD `5bfdfb3ea66f4a51c3efd44fc325abc80a317077` adds only the preserved external orchestration-skill commit; final graph `2,013/752/0`, `115,877` nodes / `160,913` relationships; detect CRITICAL `35/8/288/8`; protected handoffs `21` | exact four-invariant P6-B reject-only repair, real self-analyze duplicate-site hardening, invalidated build/runtime/regression/benchmark gates, final cleanup/detect, and re-anchor inventory | site retention `missing -> correct`; canonical identity/member precedence `wrong -> correct`; exact matrix `partial -> correct`; P6-B `REJECT -> REPAIR CODER-VALIDATED / READY FOR INDEPENDENT RE-REVIEW`; P6-C1/C2/C3/D remain locked; no stage/commit/target access | `E6-P6B-SRC1`, `E6-P6B-PROVENANCE1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-DETECT1`, `E6-P6B-REVIEW1` | seal one immutable coder resubmission report; independent Supervisor re-reviews exact P6-B candidate; Main alone may act after PASS |
| R8 | 2026-08-22 | authoritative HEAD `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`; second Supervisor report SHA-256 `C373D2413F1D60082904729F5A536B09D588A5D1DABE12181E454BCE2AD3209A`; review graph `2,015/752/0`, `115,911` nodes / `160,947` relationships; protected handoffs `22` | independent re-review of the exact four-blocker repair and report-only seal | canonical identity, receiver terminality, exact ten-vector matrix, and valid-catalog site carriage remain cleared; P6-B `READY FOR RE-REVIEW -> REJECT` only for typed catalog-validation reason/provenance carriage; P6-C1/C2/C3/D remain locked | `E6-P6B-REVIEW1` | repair only the residual catalog-failure carriage path; no generator/catalog identity rerun or later-slice expansion |
| R9 | 2026-08-22 | same authoritative HEAD and P6-B base; final graph `2,023/752/0`, `116,099` nodes / `161,226` relationships; explicit detect CRITICAL `35/8/307/8`; protected handoffs `23` after concurrent `09:49` Main transfer | exact catalog-failure reason/provenance residual repair, six-class authority/resolver/analyze matrix, patch-invalidated build/package/runtime/regression/benchmark gates, and `20/20` cleanup | catalog-failure carriage `wrong/error -> correct coder-validated`; P6-B `REJECT -> RESIDUAL REPAIR CODER-VALIDATED / READY FOR INDEPENDENT RE-REVIEW`; P6-B remains unchecked, unstaged, uncommitted; P6-C1/C2/C3/D and target remain locked | `E6-P6B-SRC1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-DETECT1`, `E6-P6B-REVIEW1` | seal exactly one new immutable coder resubmission report; independent Supervisor reviews same P6-B slice; Main alone may act after PASS |
| R10 | 2026-08-22 | authoritative HEAD `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`, P6-B base `b98131e44932a7bcac17b487ecb2914535927d01`; final post-report graph `2,028/752/0`, `116,193` nodes / `161,365` relationships; explicit detect CRITICAL `35/8/328/8`; protected handoffs `25` after the concurrent `11:50` rotation | third independent review, exact ready proof-state/status/reason/hash repair, direct accepted/forbidden matrix, patch-invalidated build/package/runtime/affected regression, artifact measurement, holder cleanup, and `9/9` turn cleanup | P6-B `READY FOR RE-REVIEW -> REJECT` on one ready-proof contradiction -> `PROOF-MATRIX REPAIR CODER-VALIDATED / READY FOR INDEPENDENT RE-REVIEW`; all earlier clearances retained; P6-B remains unchecked, unstaged, uncommitted; P6-C1/C2/C3/D and target remain locked | `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-DETECT1`, `E6-P6B-REVIEW1` | seal exactly one new immutable coder report; independent Supervisor reviews the same P6-B candidate; Main alone may act after PASS |
| R11 | 2026-08-22 | current external/user-owned HEAD `fa351c60617212635ef57a43b85d7449ef1eea1c`, parent `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`; final PASS report SHA-256 `F66E6B8337EC462AE6A7CF03C0C7D5A57176FAF22F7F430E52BF745749FF4C6E`; accepted source/catalog/binary hashes unchanged; post-ledger graph `2,030/752/0`, `116,218` nodes / `161,390` relationships / `19,426` dependency edges; explicit detect CRITICAL `35/8/328/8`; index empty | final independent P6-B review, Main seal/source/Git re-anchor, post-PASS four-ledger finalization, one current graph refresh, and required explicit detect | P6-B `READY FOR INDEPENDENT RE-REVIEW -> PASS -> CURRENT DETECT PASS`; all earlier clearances retained and residual same-invariant surfaces `none`; P6-B stays unstaged/uncommitted only for exact staging and one isolated commit; P6-C1/C2/C3/D and target stay locked until commit | `E6-P6B-REVIEW1`, `E6-P6B-SRC1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-DETECT1` | stage and commit exact accepted 35-path manifest; then open P6-C1 preserve-only |
| R12 | 2026-08-22 | isolated HEAD `5c1584fef7153a7a331c8fedd1ce64176ddc873d`, parent `fa351c60617212635ef57a43b85d7449ef1eea1c`; exact commit manifest `35` paths; post-commit index empty; remaining status `31` protected paths; diff-check PASS | exact accepted-manifest staging, isolated P6-B commit, and commit-boundary verification | P6-B `CURRENT DETECT PASS -> COMMITTED/CLOSED`; expected/staged/commit paths `35/35/35`, missing/extra `0/0`; P6-C1 `LOCKED -> PRESERVE-ONLY OPEN`; P6-C2/C3/D and target remain locked | `E6-P6B-COMMIT1` | execute only P6-C1 preserve-only closure; no production/test/target edit |
| R13 | 2026-08-22 | same HEAD/parent; fresh self-analyze `2,030/752/0`, `116,218` nodes / `161,390` relationships / `19,426` dependency edges; graph indexed/current and stale=false; pre-report status `35` paths with empty index | P6-C1 authority/current-source/config/package/commit inventory and preserve-only nearest-boundary validation; no production/test/fixture/package edit | P6-C1 `PRESERVE-ONLY OPEN -> WORKER-VALIDATED / READY_FOR_INDEPENDENT_SUPERVISOR`; active cases and production/test owners remain `0`; build/detect are N/A for zero implementation diff; P6-C1 stays unchecked and P6-C2/C3/D/target stay locked | `E6-P6C1-IMPACT1`, `E6-P6C1-SCOPE1`, `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-TEST1` | independent Supervisor reviews `reports/coder/rp_coder_260822_135011_by_gpt-5_p6c1_project_package_preserve_only.md`; Main alone may commit closure after PASS |
| R14 | 2026-08-22 | same HEAD/parent; independent PASS report SHA-256 `C1C5D0DEDC2F910D29FB32C9666E192A6221041DB578306F107B504FAF357E9B`; Supervisor graph `2,032/752/0`, `116,251` nodes / `161,423` relationships / `19,426` dependency edges; pre-report candidate five paths, post-report status `38`, index empty, diff-check PASS | independent P6-C1 preserve-only authority/source/config/package/graph/Git review and report-only seal | P6-C1 `READY_FOR_INDEPENDENT_SUPERVISOR -> PASS / PLANNER FINALIZED`; active cases/production owners/test owners/indexed `.d.ts`/indexed `node_modules` remain `0/0/0/0`; residual same-invariant surfaces none; build and implementation detect remain N/A; P6-C2/C3/D/target stay locked until exact commit | `E6-P6C1-REVIEW1` | stage and commit exactly four ledgers + worker report + PASS report; then open P6-C2 without a second acceptance loop |
| R15 | 2026-08-22 | HEAD `8055f0a6860721e26462572e34469e0d708d4a52`, parent `5c1584fef7153a7a331c8fedd1ce64176ddc873d`; exact commit manifest six paths / `454` insertions / `31` deletions; post-commit index empty; remaining status `32` protected paths; diff-check PASS | exact accepted-manifest staging, isolated P6-C1 doc/report commit, and commit-boundary verification | P6-C1 `PASS / PLANNER FINALIZED -> COMMITTED/CLOSED`; expected/staged/commit paths `6/6/6`, missing/extra `0/0`; P6-C2 `LOCKED -> OPEN`; P6-C3/D/target remain locked | `E6-P6C1-COMMIT1` | execute only P6-C2 fresh owner/consumer preflight; do not access target or open P6-C3/D |
| R16 | 2026-08-22 | same HEAD/parent; final packaged root graph `2,039/756/0`, `116,657` nodes / `162,181` relationships / `19,658` dependency edges; explicit detect CRITICAL `44/17/202/17`; current protected history observed pathname-only as `34` (`29` Main / `3` Supervisor / `2` old coder) after two concurrent Main rotations | P6-C2 exact owner/consumer impact, referenced-only materializer, Ladybug/process/MCP adapters, production/package/runtime/native/affected/broad regression, benchmark, cleanup, final graph and detect | external representation `missing -> correct worker-validated`; P6-C2 `OPEN -> CURRENT DETECT PASS / READY_FOR_INDEPENDENT_SUPERVISOR`; no File/DEFINES/IMPORTS/CONTAINS pollution; P6-C2 remains unchecked/unstaged/uncommitted; P6-C3/D/target remain locked | `E6-P6C2-IMPACT1`, `E6-P6C2-DESIGN1`, `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-PARITY1`, `E6-P6C2-DETECT1` | seal exactly one coder handoff; independent Supervisor reviews exact candidate; Main alone may stage/commit after PASS |
| R17 | 2026-08-22 | same HEAD/parent; exact candidate remains `22/22`, every production/test/ledger/coder identity unchanged; two immutable procedural REJECTs `C1F7C195...7FA88` and `537F59AA...82B2C`; final commandless PASS `D55D2F33...EA0C`; post-PASS status before ledger finalization `61` with empty index | independent P6-C2 technical reviews, clean E-rooted runtime/persistence/reader proof, procedural no-C isolation, external candidate rehash, and one planner finalization | P6-C2 `READY_FOR_INDEPENDENT_SUPERVISOR -> procedural REJECT with no code blocker -> clean procedural REJECT with no code blocker -> commandless PASS / PLANNER FINALIZED`; residual same-invariant surfaces none; no coder repair; P6-C3/D/target remain locked until commit | `E6-P6C2-REVIEW1` | run one fresh explicit-path Main analyze/detect, stage exact accepted `23` paths excluding both REJECT reports/protected history, and create one isolated P6-C2 commit |

## Phase Touch Map

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| repository workspace | `internal/resolution/indexes.go` | builds all current resolution indexes from repository facts | P6-A/P6-B | inspect-only in P6-A; later minimum edit only after decision/impact | `E0-P0A-FD1`, `E0-P0A-IMPACT1` | CRITICAL; preserve unrelated languages/lookup behavior |
| call/type/member resolution | `internal/resolution/resolve.go` | consumes workspace and emits resolved/gap evidence | P6-A/P6-C3 | inspect-only until exact finalizer impact | `E0-P0A-FD2`, `E0-P0A-IMPACT2`, `E0-P0A-IMPACT3` | no target-name branch or stage overwrite |
| health classification | `internal/graphhealth/diagnostics.go` | currently guesses classification from target text | P6-A/P6-D | inspect in P6-A; edit same-invariant projection in P6-D | `E0-P0A-FD3`, `E0-P0A-IMPACT4` | health cannot invent resolver truth |
| offline catalog/config owners | new `anvien-web/scripts/generate-tsstdlib-catalog.mjs`, new `internal/tsstdlib/catalog.v1.json`, `catalog.go`, `profile.go`, plus minimum analyze/resolution wiring | exact authority/config inputs selected by P6-A | P6-B | locked until P6-A PASS/commit; then fresh impact before edit | `E6-P6A-SRC1`, `E6-P6A-DECISION1` | no runtime Node/network/install/script; no name map |
| Child 05 result | predecessor four-file set | repository terminal/unresolved input | P6-A/P6-C3 | dependency / inspect-only | `E0-P0A-DEPEND1` | no duplicate module/export traversal |
| affected readers | Ladybug CSV/schema, processes, MCP context/impact/rename, graph-health/ResolutionGap; graph JSON/HTTP/Web/file-detail validate/preserve-only | external target/outcome preservation | P6-C2/P6-C3/P6-D | edit only named active semantic rows after fresh impact | `E6-P6A-CONSUMER1`, `E6-P6A-IMPACT1` | no broad all-reader/UI expansion |
| `E:\cheapapp.org` source | target `.anvien` output | real integration subject | P6-D | source preserve; normal analyze/read only | `E0-P0A-BOUNDARY1` | no copy, fixture, report, or source edit in target |

## Detailed Findings

### Missing declaration authority precedes graph-health

Current state:

- `buildWorkspace` accepts repository ScopeIR and indexes its definitions, scopes, bindings, imports, members, and types.
- No current input supplies TypeScript standard-library declarations.
- `resolveTypeAnnotation` exempts a small primitive list, then queries the repository workspace.
- Member calls require receiver/type/member owners in the same workspace.
- `classifyDiagnostic` later infers categories from target strings and Go-focused tables.

Required state:

```text
source/config -> accepted declaration authority -> resolver target or explicit capability result
resolver outcome -> graph/persistence -> graph-health projection
```

Classification: declaration authority and structured outcomes are missing; graph-health classification is wrong for this invariant.

Allowed next action: P6-A chooses the authority and required target/outcome facts from source/config/runtime/consumer evidence; later slices implement only that accepted decision.

Forbidden next action: changing only graph-health labels or hardcoding the three target names.

### Broader declaration scope is not yet proved

Current state:

- The originating report proposes standard-library, project, package, ambient, augmentation, and intrinsic sources.
- The accepted bounded investigation proves the target workspace lacks TypeScript ambient/lib authority.
- It explicitly leaves broader declaration semantics and policy unresolved.

Required state:

```text
P6-A evidence decides which sources/config cases this campaign owns.
P6-B implements standard-library authority.
P6-C1/P6-C2 activate only for proved lookup/representation needs.
```

Classification after P6-A: standard-library authority is missing implementation; project/package scope is preserve-only; referenced external representation is required and missing implementation.

Allowed next action: independent Supervisor reviews the recorded decision. Only after PASS and commit may P6-B implement the exact embedded-catalog authority.

Forbidden next action: importing a prewritten file/source/transport matrix into the implementation plan.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P6-A | authority/config/outcome/owner/consumer decision has final reject-only Supervisor PASS and exact seven-path commit `b98131e44932a7bcac17b487ecb2914535927d01` | preserve accepted decision; no further P6-A work |
| P6-B | final independent report accepts the exact proof-state/status/reason/hash matrix, all `25` named direct rows, affected six-class carriage, and every retained prior clearance; current graph/detect PASS and isolated commit `5c1584fef7153a7a331c8fedd1ce64176ddc873d` contains exact 35-path manifest | closed; preserve accepted implementation and evidence; no second acceptance/detect loop |
| P6-C1 | accepted P6-A authority and current source/config/package/graph evidence require zero project/package declaration cases; worker validation, independent Supervisor PASS, and exact commit `8055f0a6860721e26462572e34469e0d708d4a52` close the preserve-only slice | closed; preserve result and require planner refresh before any future activation |
| P6-C2 | referenced-only `ExternalSymbol`, Ladybug/process/MCP adapters, real packaged runtime, and exact provenance/ownership tests are independently accepted; both procedural REJECTs remain immutable history and final commandless report closes their sole no-C blocker | PASS / PLANNER FINALIZED; run current analyze/detect and commit exact accepted 23-path manifest before P6-C3 opens |
| P6-C3 | current generic gap is insufficient; minimum status/reason/provenance contract is selected | implement one final versioned outcome per site after C2; refresh shared-struct plan only if Note carriage is not lossless |
| P6-D | health guesses from text and target is `0/3` | map outcomes mechanically and prove the three exact sites |

## Implementation Gate

- [x] Target scope is limited to Child 06 responsibilities.
- [x] Each current unit has a status and exact P0 evidence IDs.
- [x] Current target files have file-detail related-file counts.
- [x] HIGH/CRITICAL symbol impacts and blast-radius counts are recorded.
- [x] Problem origin, bounded verification, and proposed design are distinguished.
- [x] Missing, wrong, and blocked behaviors have one owning slice.
- [x] Target/scanner and predecessor boundaries are explicit.
- [x] R0 records the current repo/graph basis.
- [x] Child 05 handoff is accepted at exact closure commit `ec765debff335540c77d409ebb2c9f45e4a0a77d` and reflected in refresh row R1.
- [x] P6-A records the selected authority/config/consumer decision and updates later slice steps before production code.
- [x] P6-A immutable report records the missed `03:30 +07` deadline instead of claiming timely delivery.
- [x] Independent Supervisor accepts the P6-A decision after exact contract-boundary repair; Main is authorized to create the isolated decision commit before P6-B.
- [x] P6-B production behavior, tests, full-build attempt/production builds, packaged runtime boundary, mechanism benchmarks, cleanup, fresh graph, and detect-changes are recorded on final coder bytes.
- [x] Initial independent P6-B review is sealed as `REJECT` with exactly four blocking invariants and explicit clearances; repair remains inside P6-B.
- [x] Reject-only P6-B repair closes all four blockers locally, passes real self-analyze after identical-site canonicalization, records final build/runtime/regression/benchmark/detect/cleanup evidence, and preserves all 21 Main handoffs.
- [x] Catalog-failure residual repair preserves exact loader reason and explicit ready/missing/rejected proof semantics through direct resolver and built `analyze.Run`, reruns only invalidated gates, cleans all 20 exact paths, and pathname-only preserves all 23 observed Main handoffs.
- [x] Third reject-only repair enforces the exact accepted proof-state/status/reason/hash matrix, passes all `25` direct named rows and affected six-class authority/resolver/analyze gates, rebuilds the real destination after exact holder cleanup, cleans `9/9` turn paths, and pathname-only preserves all 25 observed Main handoffs.
- [x] Independent Supervisor accepts P6-B with no residual same-invariant surface.
- [x] Current post-ledger graph refresh and explicit detect PASS with full CRITICAL warning/counts preserved.
- [x] Exact accepted-manifest staging and isolated P6-B commit `5c1584fef7153a7a331c8fedd1ce64176ddc873d` are complete and independently rechecked against the 35-path manifest.
- [x] P6-C1 worker validation records fresh graph/source/config/package/commit evidence, zero production/test owners, two tracked but scanner-excluded frontend `.d.ts` files, zero indexed `.d.ts`/`node_modules` nodes, and exact build/detect N/A dispositions.
- [x] Independent Supervisor accepts P6-C1 preserve-only with no residual same-invariant surface; exact report seal, fresh graph inventory, five-path pre-report candidate, protected history, and zero implementation drift are recorded. Only the isolated six-path doc/report commit remains before P6-C2 opens.
- [x] Exact accepted six-path P6-C1 staging and isolated commit `8055f0a6860721e26462572e34469e0d708d4a52` are complete; post-commit index/diff/protected-history boundary is verified and P6-C2 is open.
- [x] P6-C2 production-first implementation, focused/affected/native persistence tests, packaged runtime/context/impact/rename boundary, deterministic graph and package/graph measurements are worker-validated with full baseline failure classification; independent review, detect seal, and commit remain ordered gates.
- [x] Independent P6-C2 review preserves both procedural REJECTs, accepts the unchanged technical candidate through a separate commandless PASS, records no coder repair or residual same-invariant surface, and planner-finalizes the exact 23-path accepted commit boundary. Current Main analyze/detect and the isolated commit remain ordered closure gates.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing predecessor evidence.

Decision note:

P0 proves the missing workspace declaration authority and downstream text inference with current source and blast-radius evidence. Child 05 is closed at `ec765debff335540c77d409ebb2c9f45e4a0a77d`; P6-A has reject-only Supervisor PASS and exact decision commit `b98131e44932a7bcac17b487ecb2914535927d01`. All three immutable P6-B `REJECT` reports remain history: the first identified four blockers, the second retained only catalog-validation failure reason/provenance carriage, and the third retained only the ready-proof/status/reason contradiction. Final independent report `F66E6B8337EC462AE6A7CF03C0C7D5A57176FAF22F7F430E52BF745749FF4C6E` accepts the repaired invariant with no residual same-invariant surface. Isolated commit `5c1584fef7153a7a331c8fedd1ce64176ddc873d` closes P6-B. P6-C1 preserve-only has independent PASS `C1C5D0DEDC2F910D29FB32C9666E192A6221041DB578306F107B504FAF357E9B` and exact six-path commit `8055f0a6860721e26462572e34469e0d708d4a52`. P6-C2 is independently accepted: the unchanged candidate has deterministic referenced-only external facts, zero repository ownership pollution, native persistence/readback, process exclusion, and explicit MCP non-editable semantics; two procedural REJECT reports remain immutable history, and commandless PASS `D55D2F33BD311E4266D537E17F6A5EAC369FAFACA2650BD41C40140BC1F9EA0C` closes their sole execution-boundary blocker with no coder repair or residual surface. P6-C2 is planner-finalized and remains unstaged/uncommitted only for the required current Main analyze/detect and exact isolated commit. P6-C3/D and target access remain locked until that commit.
