# Supervisor Report: Child 04 P4-B1 re-export recovery resubmission REVIEW2

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260821_022347_by_gpt-5_child04_p4b1_reexport_resubmission_review2.md`
- Review time: `2026-08-21 02:23:47 +07:00` (`Asia/Bangkok`)
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien` / `Anvien`
- Scope reviewed: Child 04 `2026-07-28-04-typescript-export-semantics`, slice P4-B1 only; exact candidate worktree relative to authorized baseline `11a37aa8ec0320dd93258c058b088d1070aa778d`.
- Claim reviewed: source-bearing named/default/star/namespace re-export syntax emits one immutable `scopeir.ExportFact` per eligible site, retains valid siblings after malformed syntax, emits structured diagnostics, derives compatibility imports without drift, and contains no terminal/Child 05 state.
- Authority used: latest delegation; full `AGENTS.md`; `working-rules`; `supervisor`; Child 04 plan/evidence/benchmark/actual-status; campaign roadmap and contract; accepted P4-A/P4-B reports/contracts; prior P4-B1 Supervisor REJECT; current source, diff, parser boundary, tests, Anvien evidence, and Git state.
- Related artifacts: prior Supervisor report `reports/Supervisor/rp_supervisor_260821_012743_by_gpt-5_child04_p4b1_reexport_syntax.md`; Coder report `reports/coder/rp_coder_260821_015019_by_gpt-5_child04_p4b1_reexport_resubmission.md`.

## Executive Summary

- Vấn đề: xác định resubmission đã đóng sole blocker của REVIEW1 tại parser → `tsjs.Extract` → `ScopeIR` hay chưa.
- Quyết định: `REJECT`. Case không có comment đã được sửa (`Good, Broken as, AlsoGood` cho `2/1/2`), nhưng cùng invariant vẫn làm mất valid `AlsoGood` khi comment hợp lệ nằm trong recovered malformed specifier.
- Required outcome: sửa đúng comment-bearing recovery trong P4-B1, thêm regression sau production code, rồi gửi lại fresh boundary/build/evidence.
- Campaign boundary: `Anvien Graph Accuracy` vẫn là năm bounded defects qua bảy Child plans / 35 slices; review này không phải campaign closure. P4-C, P4-C2, Child 05 và `E:\cheapapp.org` không được mở.

## Git / Artifact Boundary

- Authorized baseline vẫn là `11a37aa8ec0320dd93258c058b088d1070aa778d`; candidate files không khác nhau giữa baseline và current `HEAD` ngoài unstaged candidate diff (`git diff --quiet 11a37... HEAD -- internal/providers/tsjs/imports.go internal/providers/tsjs/extract_test.go` exit `0`).
- Current `HEAD` đã drift bởi hai docs-only commits ngoài scope, không do review này và không được reset/checkout:
  - `84a354940aea8240c99bf4868e721209e7248830` — `docs(orchestration): mark session rotation mandatory` — chỉ đổi `internal/aicontext/skills/orchestration/SKILL.md`.
  - `ce0e200c55bd96c4374cc6e84bd99a3c82bef641` — `docs(orchestration): enforce session rotation steps` — chỉ đổi `internal/aicontext/skills/orchestration/SKILL.md`.
- Candidate vẫn uncommitted đúng hai tracked paths; index rỗng; không detect-changes, stage, commit, push, plan/ledger/roadmap edit nào được thực hiện trong review.
- Candidate identities hiện tại:
  - `internal/providers/tsjs/imports.go`: `1,297` lines / `36,055` bytes / `1,297` LF / SHA-256 `0B5E7E8F596CAD6B53CEB17D08FB86046BD071FD0D34DC94FD1FF260EF806A44`.
  - `internal/providers/tsjs/extract_test.go`: `3,076` lines / `134,250` bytes / `3,076` LF / SHA-256 `A05951319B835D8231E828FC0E172B40E20D594FF0B096D36643BE93C019A757`.
  - Coder report: `8,742` bytes / `149` LF / SHA-256 `0DB8CF9CF95FAF727BC06C985CDE88F80FEEC52EE65BB6B8F3B8AC1901EF08FF`.
  - Prior Supervisor report: `13,798` bytes / `158` LF / SHA-256 `54D5D6CD7E71DBCF65BD68C0E2BDC0DBF940BE1A48DA51A0B446FAAA0F9652CD`.
- Tracked candidate diff: `imports.go` `419` insertions / `22` deletions; `extract_test.go` `267` insertions / `2` deletions; index manifest is empty. Main-owned untracked provenance reports were preserved, including `reports/Investigation/rp_main_260821_0044_orchestration_rotation_handoff.md`.

## Blocking Findings

### [P1] Comment-bearing recovered malformed alias still drops a valid sibling

File: `internal/providers/tsjs/imports.go:166-191`, helper `internal/providers/tsjs/imports.go:257-286`.

Issue: the new recovery accepts only a four-child shape and requires whitespace-only byte gaps. Tree-sitter represents comments as named, non-error children inside the same recovered `export_specifier`; the valid post-comma alias remains an `alias` node, but the strict shape/gap checks return `false` and the entire recovered node is discarded.

Independent fresh boundary evidence (after canonical build):

```text
Source: export { Good, Broken as, AlsoGood } from "./mixed";
AST recovered node: name=Broken, anonymous as, ERROR text=",", alias=AlsoGood
Extracted: exports=2 [Good, AlsoGood], diagnostics=1 [nodeKind=as, range="as"], imports=2
```

The same valid-sibling invariant with legal comment trivia fails:

```text
Source: export { Good, Broken as, /*keep*/ AlsoGood } from "./mixed";
AST: export_specifier children name=Broken, as, ERROR comma, comment=/*keep*/ (error=false), alias=AlsoGood
Extracted: exports=1 [Good], diagnostics=1 [nodeKind=export_specifier], imports=1

Source: export { Good, Broken as /*bad*/, AlsoGood } from "./mixed";
AST: export_specifier children name=Broken, as, comment=/*bad*/ (error=false), ERROR comma, alias=AlsoGood
Extracted: exports=1 [Good], diagnostics=1 [nodeKind=export_specifier], imports=1
```

The probes used the production `parser` → `tsjs.Extract` → `ScopeIR` boundary and were removed afterward. The `comment` nodes are non-error trivia, while the recovered `alias=AlsoGood` is source-bearing and valid. Therefore this is not an unrelated parser failure: it is a direct P4-B1 cardinality/provenance loss. The implementation preserves the base no-comment case, newline recovery, type-only recovery, string-name recovery, JS recovery, star/namespace, duplicate cardinality, and missing-source diagnostics, but it does not close the comment-bearing sibling family required by “retain valid siblings.”

Why this blocks acceptance: Child 04 P4-B1 requires one immutable fact per eligible specifier and structured diagnostics without dropping valid siblings. For both comment forms, the actual result is `1` fact / `1` diagnostic / `1` compatibility import instead of `2/1/2`, and `AlsoGood` has no `ExportFact` or `ImportReexport`. This occurs at the provider extraction boundary, before P4-C projection or Child 05 resolution.

Fix direction: extend the narrowly owned recovery proof to tolerate comment/trivia children between the malformed alias tokens and the recovered valid alias, while remaining fail-closed for malformed names and emitting exactly one diagnostic for dangling `as`. Do not fabricate `Broken`; do not add terminal resolution, barrel, ambiguity, cycle, persistence, or public-API state.

Re-review evidence required:

1. Fresh production boundary output for both comment placements with `Good` and `AlsoGood` facts, exactly one diagnostic on the malformed site, exactly two derived `ImportReexport` entries, and no `Broken` fact/import.
2. Test-after-code regression asserting exact `Range`, `SelectionRange`, `StatementRange`, `SiteKind`, `TargetRaw`, meaning/type-only state, empty `LocalName`/`LocalDefID`, and zero terminal state for the recovered sibling.
3. Fresh canonical `npm run full-build` before focused/nearest validation, plus the current focused TS/JS matrix, full `internal/providers/tsjs`, nearest `tsjs/scopeir/providers`, and `resolution/analyze` compatibility regression.
4. Fresh candidate/report hashes, current authorized-baseline comparison, exact `.tmp` probe census `0`, `gofmt -d` empty, `git diff --check` exit `0`, and no later-slice or target activity.

## Source-Level Clearance Notes

- `internal/providers/tsjs/imports.go`: blocked only on comment-bearing malformed re-export recovery described above. The normal fact path at lines `192-245` preserves `TargetRaw`, target/exported names, meaning/type-only, exact ranges, statement provenance, and existing compatibility derivation.
- `internal/providers/tsjs/extract_test.go`: blocked as acceptance proof because line `698` covers only the no-comment recovered sibling; it does not cover the two comment-bearing AST shapes. Existing direct/default/type-only/star/namespace/string/JS tests remain preserve evidence.
- `internal/providers/tsjs/extract.go`: clear/preserve-only; `Extract` result wiring was inspected and no candidate diff changes it.
- `internal/providers/tsjs/definitions.go`: clear/preserve-only; no `DefinitionFact.Visibility` write or access-visibility drift was found.
- `internal/scopeir/{facts.go,kinds.go,ir.go,sort_keys.go}`: clear/preserve-only accepted P4-A contract; source facts remain syntax-only and no terminal fields were introduced.
- Later-slice scan: production candidate has zero matches for `TargetFile`, `TargetDefID`, `TargetModuleScope`, `LinkStatus`, `TransitiveVia`, `reachableThroughBarrel`, `publicApi`, `barrel`, `cycle`, `ambiguity`, `persistence`, `projection`, or `terminal`; candidate diff has zero visibility writes.

## Evidence Checked

### Đã xác minh

- Mandatory graph refresh after build, exact exclusions:

  ```text
  anvien analyze --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**" --json
  exit 0; scanned/parsed/failed = 1,141/626/0; graph = 82,020 nodes / 121,721 relationships
  ```

  The graph reports `stale=false`; candidate file-detail is current at indexed/current commit `11a37aa8ec0320dd93258c058b088d1070aa778d` basis and `changedSinceAnalyze=false`.
- Current `file-detail`:
  - `imports.go`: `184` symbols, inbound/outbound/local `7/107/40`, unresolved `501`, linked flows/tests `2/4`, file risk `HIGH`.
  - `extract_test.go`: `709` symbols, inbound/outbound/local `13/366/245`, unresolved `0`, linked flows/tests `0/3`, file risk `LOW`.
- Current upstream impact with `--include-tests`:
  - `imports.go`: `HIGH`, `18` impacted / `18` direct / `1` affected file / `1` affected flow (linked flows `2`).
  - `extract_test.go`: `CRITICAL`, `63` impacted / `60` direct / `4` affected files / `0` affected flows.
  - `emitReexportClauseFacts`: `LOW`, `0` upstream impacted.
  - `recoveredReexportSiblingAfterMalformedAlias`: `LOW`, `0` upstream impacted.
  HIGH/CRITICAL are recorded as blast-radius warnings, not as automatic prohibitions.
- Source and diff were inspected before relying on reports/tests. `emitReexportClauseFacts` now emits recovered facts only when `recoveredReexportSiblingAfterMalformedAlias` succeeds; the helper's `ChildCount()==4` and whitespace-only gap checks explain the comment failure.
- Fresh canonical build after the app restart:

  ```text
  npm run full-build
  exit 0
  runtime 1.2.8; Web 2,943 modules; Vite 22.86s
  ```

  The first build session lost stdout when the app restarted; it was not counted. The second complete run is the fresh build evidence.
- Fresh post-build behavior commands:
  - focused P4-B/P4-B1 matrix: exit `0`, `6/6` top-level tests.
  - `go test ./internal/providers/tsjs -count=1`: exit `0`.
  - `go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers -count=1`: exit `0`, `3/3` packages.
  - `go test ./internal/resolution ./internal/analyze -count=1`: exit `0`, `2/2` packages.
  - independent comment/recovery matrix: base/no-comment, newline, type-only, string, and JS cases retain `2/1/2`; both comment-bearing cases fail `1/1/1` as stated above.
- Artifact checks:
  - `gofmt -d internal/providers/tsjs/imports.go internal/providers/tsjs/extract_test.go`: exit `0`, empty output.
  - `git diff --check`: exit `0`.
  - candidate production forbidden-field scan: `0` matches; visibility-write scan: `0` matches.
  - exact P4-B1/re-export/probe `.tmp` census: `0` after deleting both Supervisor probes; shared/unrelated `.tmp` directories were preserved.
  - no access to `E:\cheapapp.org`.

### Failed

- The sole current blocker is the comment-bearing recovered-sibling invariant in the Blocking Findings section. The green checked-in test matrix is narrower and does not invalidate this fresh boundary failure.

### Not run

- `anvien detect-changes`, stage, commit, push, and plan/ledger/roadmap edits: prohibited/owned by Main for this review.
- P4-C, P4-C2, Child 05, terminal/barrel/cycle/ambiguity/public-API behavior: locked and not opened.
- Browser/Playwright: N/A for this non-UI provider/ScopeIR boundary.

## Invariant Closure

- Affected invariant: malformed source-bearing named re-export recovery must emit one structured diagnostic while retaining every valid sibling, including valid siblings represented as alias children after parser recovery and comment trivia.
- Sibling surfaces checked: exact source loop/helper; parser AST shape; TS/JS named/default/string/type-only/star/namespace cases; newline and duplicate/cardinality cases; missing-source diagnostics; compatibility imports; ScopeIR normalization; `Extract` result wiring; access visibility; P4-A contract; and later-slice field boundary.
- Residual unverified same-invariant surfaces: comment-bearing recovered alias variants remain unclosed (both comment placements above). No residual P4-C/Child 05 surface is being claimed or opened.

## Required Fix List For Resubmission

1. Repair only the P4-B1 comment-bearing recovery invariant so valid `AlsoGood` siblings survive comment trivia around the malformed `as`/comma recovery, with no fabricated `Broken` fact/import.
2. Add production-first regression evidence for both comment placements with exact fact/diagnostic/import cardinality and all syntax/provenance/meaning/zero-terminal fields.
3. Re-run the fresh build, nearest boundary, compatibility regression, graph refresh, and exact artifact checks listed above; route the same P4-B1 lane back to Supervisor.

## Overall Evaluation

Implementation quality is scoped and preserves the accepted P4-A/P4-B contract, direct facts, visibility, compatibility derivation, and later-slice boundary for the cases covered. Evidence quality is independently fresh for the current worktree and includes a parser-to-ScopeIR probe that exposes a remaining same-invariant loss. The external docs-only `HEAD` drift is recorded and does not touch the candidate; it does not excuse or conceal the comment-recovery blocker. P4-B1 is not acceptable until the exact invariant above is closed.

Next responsible party: Main/orchestration task `01a02074-9de3-7072-a2fa-c2ef10db6358`. Supervisor does not self-transition, repair candidate bytes, or open later slices.
