# Supervisor Review: Child 06 P6-C3 Structured Resolution Outcomes

Verdict: **REJECT**

## Metadata

- Review role: independent Supervisor, zero-trust, review-only.
- Scope: exact current P6-C3 candidate, "Structured resolution outcomes".
- Report path: `E:\Anvien\reports\Supervisor\rp_supervisor_260822_215636_by_gpt-5_p6c3_structured_resolution_outcomes.md`.
- Workspace: the sole reviewed workspace was `E:\Anvien`; branch `master`.
- Git anchor: HEAD `2acd357db32ac0c2bd3c3968a950c773acf3000d`; parent `8055f0a6860721e26462572e34469e0d708d4a52`.
- Review authority read in full before disposition: `E:\Anvien\AGENTS.md`; working-rules, supervisor, backend-development, Data-Integrity, and Edge-Case skill instructions; the four current Child 06 ledgers; the active P6-C3 plan slice; and the sealed coder report.
- The full candidate diff, every new candidate source/test file, and each of the five production owners were inspected before runtime disposition.
- P6-D, graph-health policy/rendering, the bounded target proof, `E:\cheapapp.org`, shared graph structs, Ladybug production schema/CSV, P6-C2 representation, declaration loading/catalog/config, generic transports, and UI remained locked or preserve-only. They were not reviewed as implementation scope and were not changed by this lane.
- No implementation, test, ledger, coder-report, or existing report byte was edited. No stage, commit, branch, stash, push, reset, checkout, package script, network operation, install, target access, alternate worktree, or packaged-binary reuse occurred.

## Sealed Coder Report And Exact Candidate Identity

The coder report independently measures exactly as delegated:

- Path: `E:\Anvien\reports\coder\rp_coder_260822_205432_by_gpt-5_p6c3_structured_resolution_outcomes.md`.
- Size/newlines: `18,206` bytes / `254` LF / `0` CR.
- Encoding: strict UTF-8, no BOM.
- SHA-256: `9398C3B5B662EB9B7EA1141DA12BAF93D74C80975C91C36AD01DD2B0B5920FC8`.
- Created: `2026-08-22T21:08:14.5271883+07:00`.
- Last write: `2026-08-22T21:08:14.5277119+07:00`.

The exact ordered 18-path candidate manifest below was measured after all review commands. Every file is strict UTF-8 without BOM and has `0` CR. The manifest digest rule is `relative-path|bytes|SHA256` in the displayed order, joined by LF with no terminal LF. The canonical manifest is `2,630` bytes and SHA-256 `E9DA6DE84ED25AB9677333F20E1D7E3127E57482C19BF870F6EC89D36E280D7F`.

| # | Candidate path | Bytes | LF | SHA-256 |
|---:|---|---:|---:|---|
| 1 | `internal/resolution/outcome.go` | 19,944 | 508 | `DF2787622F0E49DC75A0BBA47F79CB101F1C5D03FA3CB674B8F6F8139DE1E5C9` |
| 2 | `internal/resolution/types.go` | 4,933 | 132 | `7C5E113F5E50584665D6D9AED0BCB2B3C6F6085219A62CD5AE74DD7654CD5DC3` |
| 3 | `internal/resolution/emit.go` | 27,319 | 823 | `13A31C8EAE9F0665BBBA9552B3D75D1B2CA663DE3FA8594780ED0BCEACCDCB6E` |
| 4 | `internal/resolution/resolve.go` | 39,912 | 1,158 | `D01322440AC0A2CF6D5F39402791F61B23AA6894707C1E350E5B56687080E919` |
| 5 | `internal/analyze/analyze.go` | 33,734 | 1,078 | `CC9329FB8AD3510DFC5F643F39EE3CFBB749ABBB3F3CB130E519CE3DA2BBE608` |
| 6 | `internal/resolution/p6c3_structured_outcome_test.go` | 20,659 | 530 | `7D470306A42C062D30C0CEDCA11F68438D6FCC451E5AE0E5B5CBE337481AAE37` |
| 7 | `internal/analyze/p6c3_structured_outcome_test.go` | 5,101 | 131 | `5269BDC95FCDA93BD512D52F681CA0CF1B2F7C5685BB6D3BF1D8F78B36255D56` |
| 8 | `internal/lbugload/p6c3_resolution_outcome_persistence_test.go` | 9,885 | 270 | `03911F25CFF4EDD642DC1E4F044A61D1859033E1B2216CC58EB46F00267696AC` |
| 9 | `internal/lbugnative/p6c3_resolution_outcome_integration_test.go` | 9,067 | 238 | `95BF253EA9750F66069B1BDD0752951EFC2092E85D2DBC5058AAE9360572CEA4` |
| 10 | `internal/resolution/p6b_tsstdlib_test.go` | 29,684 | 735 | `F6AA3C4A3E39838565C7FC02B197897FA3DC304FB1AAEF1F632100E2D82B5BB8` |
| 11 | `internal/analyze/p6b_tsstdlib_test.go` | 14,261 | 360 | `AD902C97D3F2750955335D858CEA9D4588E0DC961A5E2A9439EA9D28DEAC0E5F` |
| 12 | `internal/resolution/resolution_test.go` | 61,139 | 1,425 | `F46DD3595831A64495E7765F012C83CF2F2A777588B0950C9CDAB0DB60B98EC4` |
| 13 | `internal/resolution/export_resolution_test.go` | 29,500 | 612 | `134035FF5303D6F0CA44AAE5BA10B422285D0FB5D1C73467ADA5AC5778149F44` |
| 14 | `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md` | 48,480 | 254 | `C3C18303AEE745762FA2E6DAA58C5D1323E37C5D5069CB6BB299F21F5F8AB0F3` |
| 15 | `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md` | 18,755 | 96 | `5E4790C5D6540C3865681FA8C3D67E21FFB23937C963E8F77C53EF302D8D0D12` |
| 16 | `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md` | 89,963 | 457 | `FB0184A8ED154201B4E50FCE5699043454577562065BA4BFE9DD78CB7A378C77` |
| 17 | `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md` | 70,897 | 519 | `FBAB751794976B12A74D9E9939FF5A6596FEDF5A68C0C2776752B54A5FEFDF49` |
| 18 | `reports/coder/rp_coder_260822_205432_by_gpt-5_p6c3_structured_resolution_outcomes.md` | 18,206 | 254 | `9398C3B5B662EB9B7EA1141DA12BAF93D74C80975C91C36AD01DD2B0B5920FC8` |

At the delegated Main check, status was exactly `60 = 18 candidate + 42 protected`, with missing/unexpected `0/0`. Before this Supervisor report was created, current status was exactly `61 = 18 candidate + 43 protected`: the same candidate plus `36` protected Main Investigation reports, `5` prior untracked Supervisor reports, and `2` old coder reports. The one-report increase from the delegated check is the concurrent protected Main handoff `reports/Investigation/rp_main_260822_214330_orchestration_rotation_handoff.md`; it is not part of the candidate and was not used as technical proof. The index remained empty.

## Fresh Anvien Evidence And Blast Radius

`anvien --help` was run first. A fresh explicit `anvien analyze E:\Anvien --force` then completed before any graph read. The current graph is fresh/stale=false and reports:

- `2,055` files scanned; `761` code files parsed; `0` failed.
- `117,379` nodes; `163,621` relationships; `20,145` dependency edges.
- Semantic completeness `117,379 / 117,379`.

The coder's sealed observation was `2,054 / 761 / 0 / 117,364 / 163,606 / 20,145`. The current `+1` scanned file, `+15` nodes, and `+15` relationships are explained by indexing the subsequently sealed coder report. The production/test candidate bytes did not drift.

Fresh file-detail summaries are reported as symbols / inbound / outbound / local relationships / unresolved / linked flows / linked tests:

- `outcome.go`: HIGH `119 / 84 / 124 / 117 / 147 / 11 / 23`.
- `types.go`: HIGH `94 / 155 / 26 / 90 / 35 / 1 / 24`.
- `emit.go`: HIGH `153 / 79 / 230 / 79 / 389 / 13 / 24`.
- `resolve.go`: HIGH `200 / 128 / 289 / 87 / 400 / 26 / 41`.
- `analyze.go`: HIGH `334 / 92 / 320 / 175 / 524 / 20 / 8`.

Fresh exact upstream impacts are impacted symbols / files / modules / processes:

- `ResolutionOutcome`: CRITICAL `44 / 8 / 4 / 58`.
- `resolutionOutcomeCollector`: CRITICAL `33 / 5 / 2 / 49`.
- `projectResolutionOutcomes`: CRITICAL `6 / 4 / 4 / 32`.
- `ResolveBoundInto`: CRITICAL `7 / 5 / 5 / 31`.
- `resolveCall`, `resolveAccess`, and `resolveTypeAnnotation`: each CRITICAL `6 / 4 / 4 / 32`.
- `recordTypeScriptLookup`: CRITICAL `6 / 2 / 2 / 36`.
- `emitReference`: CRITICAL `8 / 3 / 2 / 36`.
- `emitUnresolvedReference`: CRITICAL `7 / 2 / 2 / 36`.
- `analyze.Result`: CRITICAL `5 / 4 / 4 / 31`; `analyze.Run`: CRITICAL `5 / 4 / 3 / 23`.

Fresh `anvien detect-changes --repo E:\Anvien --scope all` exited `0` but correctly retained CRITICAL risk:

- Affected: `36` symbols / `12` files; layers `backend=32`, `mixed=4`; areas `analyzer=8`, `mixed=8`, `resolution=20`.
- Changed: `76` symbols / `12` files; layers `backend=50`, `backend_test=6`, `docs=20`; areas `analyzer=6`, `documentation=20`, `resolution=50`.
- ResolutionGap delta: `19` entities / `19` occurrences.
- Resolution Health: `0`.
- Semantic completeness: `117,379 / 117,379`.

The sealed coder observation was `72` changed symbols and `docs=16`; the current four additional documentation symbols reflect the final living-ledger state present in the exact manifest. Detect's Git denominator is the `12` tracked modified files and, by Git design, excludes the five untracked candidate source/test owners. Those five were independently covered by the fresh graph, file-detail, source inspection, targeted builds, and runnable tests. CRITICAL/HIGH are blast-radius warnings, not edit prohibitions; their full counts are preserved here and make the two behavioral failures release-blocking.

## Blocking Finding 1: Intrinsic Classification Bypasses Repository Terminality

`internal/resolution/resolve.go:735-749` computes `targetName`, then at lines `740-742` calls `isBuiltinType(targetName)`, records an intrinsic final outcome, and returns. Repository lookup through `w.resolveName(...)` does not occur until line `749`.

This ordering violates the required authority chain: repository/P5 resolution must be terminal before intrinsic or TypeScript fallback is considered. A language token that is normally predeclared is still a valid repository declaration name in languages that allow shadowing; target text alone cannot preempt repository ownership.

An E-only hostile resolver probe used a valid Go repository type named `string`. The candidate finalized its type-reference site as:

- status `resolved_internal`;
- stage `intrinsic`;
- target `Intrinsic:go#type#string`.

The repository declaration was present but never queried. A control repository type named `CustomString` resolved at repository stage. This differential result isolates the early intrinsic branch, rather than fixture or workspace construction, as the cause.

Required fix direction: run repository/P5 resolution first and honor any repository success or authoritative terminal failure. Only an unclaimed/non-terminal site may proceed to language-intrinsic classification, then to the accepted TypeScript authority fallback. Add a durable collision fixture using a legal repository declaration that shadows a predeclared identifier; assert exact repository status, stage, target, proof, and one final outcome.

## Blocking Finding 2: Stable Sites Can Exit Without A Final Outcome

The collector enforces uniqueness only for outcomes that reach it. Two production control paths bypass the collector entirely.

### Local receiver missing-member call

In `internal/resolution/resolve.go:536-545`, the unresolved call branch constructs site identity and P5 evidence. When `bindingReceiverResolved && w.typeScriptStandardLibrary == nil`, lines `542-543` return before `emitUnresolvedReference`, which is currently also the repository-unresolved outcome recorder.

An E-only no-authority probe with a local receiver and `Math.missing` produced an outcome for the receiver access `Math`, but no outcome for the stable call site `SourceSite:src/member.ts#call#Math.missing#3#1#3#13`. The site therefore has neither target nor explicit reason/proof, even though the resolver visited it and the local receiver/P5 path was known.

### Resolved heritage with compatibility reference disabled

`internal/resolution/resolve.go:85-89` processes every resolved heritage item and increments `ResolvedInheritance`. `internal/resolution/emit.go:564-603` emits the primary EXTENDS/IMPLEMENTS relationship, but lines `582-584` return when `DisableScopeInheritsCompatibility=true`. The sole `emitReference` call that supplies stable site identity and records the repository-resolved outcome is after that return at lines `585-603`.

An E-only resolved-heritage probe with compatibility disabled emitted the primary resolved relationship but zero outcomes for the heritage source site. A final resolver decision is therefore incorrectly coupled to optional compatibility/reference materialization.

Required fix direction: finalize every visited call/access/type/heritage site independently of optional graph, diagnostic, or ReferenceIndex emission. A silent-return policy may suppress a compatibility carrier, but it must not suppress the immutable result. Add durable fixtures for (1) a local resolved receiver whose member is missing with no external authority and (2) a resolved heritage target with compatibility emission disabled. Each fixture must assert exactly one outcome with exact status/stage/target-or-reason/proof and no resolved/unresolved overlap.

## Invariant Disposition

| Required invariant | Supervisor result | Evidence |
|---|---|---|
| One immutable final outcome per stable source site | Blocked | The `Math.missing` call and compatibility-disabled resolved heritage site each produce zero outcomes. |
| Exact status/stage/target-or-reason/proof | Locally clear for produced outcomes; globally blocked | Validation at `outcome.go:144-217` enforces the four accepted shapes, but missing sites have no shape and the shadowed repository type has the wrong stage/target. |
| Repository/P5 terminal before intrinsic/TypeScript fallback | Blocked | `resolve.go:740-749` finalizes intrinsic before `w.resolveName`; hostile shadowing probe confirms wrong authority. |
| Immutable snapshot and duplicate behavior | Clear | Collector clones input/nested evidence, canonicalizes identical duplicates, and rejects conflicting duplicate final decisions at `outcome.go:83-128`; focused mutation/replay/conflict tests pass. |
| Conflict fail-closed | Clear | Conflicting source-site results set collector error and finalization returns it; constructed conflict test passes. |
| Zero resolved/unresolved overlap | Clear for represented carriers | Projection rejects missing outcomes for represented resolved carriers, rejects resolved/non-resolved mismatches, and rejects relationship/diagnostic overlap at `outcome.go:399-507`; constructed overlap test passes. This does not detect stable sites that emit no carrier at all. |
| P5 proof nesting | Clear | Resolved terminal plus two hops and unresolved missing-member proof with one namespace hop plus one failure record survive inside `ResolutionProof`; focused deep-equality evidence passes. |
| Graph JSON and Ladybug CSV carriage | Clear for exercised carriers | Resolved relationship Evidence and unresolved Diagnostic/ResolutionGap Note round-trip to equal outcome structs; zero export skips/fallback failures/warnings. |
| Native Ladybug readback | Clear for exercised carriers | Real repo-local Ladybug v0.19.1 readback passes with `2` node COPY / `2` relationship-pair COPY / `0` fallback / `0` fallback failures / `0` skips. |
| `analyze.Result` is carriage only | Source-level clear; runtime not established | `analyze.go:120-124` exposes the slice and `analyze.go:356-359` copies resolution output without reclassification or target-text inference. The package/runtime could not execute under the offline dependency boundary. |
| No target-text inference in P6-C3 | Clear | No candidate production branch derives graph-health classification/actionability from target text. The intrinsic precedence bug is resolver authority ordering, not P6-D classification code. |
| P6-D and graph-health policy untouched | Clear | No graph-health production byte is in the diff. Existing policy remains owned by P6-D and is neither accepted nor modified here. |

The zero-overlap and conflict guards are meaningful, but they are downstream guards over the collector/carrier population. They cannot establish completeness when a visited source site returns before recording anything.

## Production Source-Level Clearance

| Production owner | Source-level disposition |
|---|---|
| `internal/resolution/outcome.go` | The versioned schema, exact four statuses, target/reason shape validation, nested cloning, deterministic sorting, duplicate/conflict handling, outcome encoding, projection parity, and represented-carrier overlap checks are locally sound. It lacks an independent visited-site completeness inventory, so producer omissions remain invisible. |
| `internal/resolution/types.go` | `resolution.Result` carries `[]ResolutionOutcome`; accepted P5 authority remains nested and ownership-separated. No shared graph contract or P6-D policy was added. Source-level carriage is clear. |
| `internal/resolution/emit.go` | Normal resolved/unresolved reference paths attach exact outcome JSON and preserve proof. The heritage compatibility early return at `emit.go:582-584` is a blocking completeness defect because outcome creation is coupled to the optional reference carrier. |
| `internal/resolution/resolve.go` | Final collection, projection, and returned result are fail-closed for recorded outcomes. The early intrinsic branch at `resolve.go:740-742` violates authority order, and the call early return at `resolve.go:542-543` drops a stable site. Source-level acceptance is blocked. |
| `internal/analyze/analyze.go` | `analyze.Result` is carriage-only: it receives the resolver slice and performs no inference or mutation. Source is clear; `internal/analyze` package execution and packaged candidate runtime remain unproven, not successful. |

## Build, Test, Persistence, And Blocker Truth

Before building, only E-rooted lock/holder evidence was considered. No proven active E-rooted build holder required termination, no broad process inventory was invoked, and no process was killed.

Independent build observations:

- `go build -buildvcs=false ./internal/resolution` passed and proves the primary resolution production package compiles.
- `go build -buildvcs=false ./internal/lbugload ./internal/graphhealth ./internal/semantic` passed and proves the affected persistence/projection dependencies compile.
- The required `go build -buildvcs=false ./...` exited `1`; it is blocked, never a success. The E-local offline module cache lacks Cobra, Hugot, and tree-sitter modules while `GOPROXY=off` forbids retrieval. Baseline fixture-wide mixed-package and standalone-C failures also remain.
- `internal/analyze` package execution did not pass because its external dependency chain is unavailable offline.
- Normal package-mode resolver/native test attempts that require the parser dependency chain are likewise environment-blocked before candidate execution; they are not candidate successes or behavioral failures.
- The invalidated packaged candidate runtime was not built. No stale packaged binary was used as candidate proof, so packaged runtime status is not established.

Independent runnable evidence that did pass:

- Focused P6-C3 file-list tests cover the four status shapes, duplicate canonicalization, immutability, deterministic replay, conflict fail-closed, represented-carrier overlap rejection, P5 proof nesting, and Diagnostic.Note carriage.
- The broad runnable parser-independent P6-B/P6-C2/P6-C3 selection covers repository terminality for its existing non-collision fixtures, external materialization/coalesced proof, definition conflicts, access audit, export-table wiring, suffix resolution, and shadow parity.
- `TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity` passes: Graph JSON round-trip is deep-equal for resolved relationship Evidence and non-resolved ResolutionGap Note; Ladybug CSV records zero skipped nodes/relationships and zero load fallback failures/warnings.
- `TestP6C3NativeLadybugResolutionOutcomeReadback` passes against the repo-local Ladybug v0.19.1 assets: resolved Evidence and unresolved Note decode to the exact input structs, with the COPY/fallback/skip counts recorded above.
- Two additional hostile probes expose the release blockers: repository type `string` loses to intrinsic finalization, and the two early-return families produce no final outcome. These negative probes are stronger than the existing green fixture set because they exercise authority collision and carrier-disabled paths omitted by current tests.

Formatting and repository hygiene remained clean: `gofmt -d` over all 13 candidate Go files emitted zero lines; `git diff --check` exited `0`; the index is empty. Review-only temp roots `E:\Anvien\.tmp\p6c3-supervisor-01a029d2`, `E:\Anvien\.tmp\p6c3-lbugload-tests`, and `E:\Anvien\.tmp\p6c3-lbugnative-tests` were verified inside the repo, removed, and verified absent. Exactly `17` zero-byte module-download lock markers created by the offline attempt were removed; final matching lock count is `0`. No unrelated temp content was removed.

## Residual Surfaces And Required Re-review

The same coder must repair only the exact P6-C3 producer/finalization defects:

1. Move repository/P5 authority ahead of intrinsic finalization without weakening terminal success or terminal failure semantics.
2. Decouple final outcome recording from optional graph/reference/diagnostic emission so every visited call, access, type-reference, and heritage site finalizes once.
3. Add durable regression fixtures for the repository/predeclared-name collision, no-authority local-receiver missing member, and compatibility-disabled resolved heritage case.
4. Preserve the already-clear conflict, overlap, P5 nesting, lossless Note, Graph JSON, Ladybug CSV, native readback, and carriage-only behavior.

Re-review must use the same exact candidate lineage and independently show:

- one and only one exact outcome for each new hostile fixture;
- repository stage/target/proof for the shadowing declaration;
- explicit unresolved reason/proof for the local missing member;
- repository-resolved target/proof for compatibility-disabled heritage;
- no resolved/unresolved overlap;
- existing status, conflict, immutability, P5 proof, Graph JSON/CSV/native tests still passing;
- fresh analyze, file-detail, impact, and detect evidence with the complete CRITICAL/HIGH counts;
- truthful full-build, `internal/analyze`, and packaged-runtime status under the actual offline boundary.

P6-C3 remains unchecked and uncommitted. P6-D, graph-health policy/rendering, and the target proof remain locked. Main must return the two concrete blockers to the same coder, must not stage or commit this candidate, and must open a new independent Supervisor review only after the corrected sealed coder report and exact boundary are available.
