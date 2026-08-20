# Supervisor Report: Child 04 P4-B1 re-export syntax facts

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260821_012743_by_gpt-5_child04_p4b1_reexport_syntax.md`
- Review time: `2026-08-21 01:27:43 +07:00` (`SE Asia Standard Time` / `Asia/Bangkok`)
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien` / `Anvien`
- Baseline HEAD: `11a37aa8ec0320dd93258c058b088d1070aa778d`
- Scope reviewed: Child 04 `P4-B1` candidate, exact unstaged production/test diff tại `internal/providers/tsjs/imports.go` và `internal/providers/tsjs/extract_test.go`; P4-C/P4-C2/Child 05 giữ locked.
- Claim reviewed: mỗi eligible source-bearing named/default re-export specifier hoặc star/namespace site trên TS/JS tạo đúng một immutable `scopeir.ExportFact`, giữ syntax-only names/ranges/provenance/meaning/type-only state, derive compatibility facts, phát structured diagnostics cho malformed site nhưng vẫn giữ valid siblings, bảo toàn P4-B direct facts và access visibility, không thêm later-slice state.
- Authority used: latest delegation; full `AGENTS.md`; `working-rules`; `supervisor`; internal orchestration authority; Child 04 plan/evidence/benchmark/actual-status; campaign roadmap; accepted P4-A/P4-B source/contracts/reports; Main handoff; current source/diff/Git/runtime evidence.
- Related artifacts: `reports/coder/rp_coder_260821_010023_by_gpt-5_child04_p4b1_reexport_syntax.md`; `reports/Investigation/rp_main_260821_0044_orchestration_rotation_handoff.md`; accepted P4-A/P4-B Supervisor reports.

## Executive Summary

- Vấn đề: quyết định candidate P4-B1 có đóng đầy đủ source-bearing re-export syntax invariant hay không, không biến review này thành campaign closure.
- Quyết định: `REJECT`. Candidate làm mất một valid re-export sibling khi sibling đó đứng sau một malformed alias trong cùng clause. Green build và các test hiện tại không đo failure shape này.
- Required outcome: sửa production recovery để giữ cả valid sibling phía trước và phía sau malformed specifier, sau đó thêm regression test sau code và gửi lại đúng fresh evidence nêu dưới đây.
- Boundary: campaign `Anvien Graph Accuracy` vẫn phải đóng năm bounded defects qua bảy Child plans / 35 slices. P4-C, P4-C2, Child 05 và `E:\cheapapp.org` không được mở bởi verdict này.

## Blocking Findings

### [P1] Malformed specifier recovery làm mất valid sibling phía sau

File: `internal/providers/tsjs/imports.go:166`

Issue: với source hợp lệ-phục-hồi theo clause `export { Good, Broken as, AlsoGood } from "./mixed";`, tree-sitter phục hồi phần sau thành một `export_specifier` có `name=Broken`, một `ERROR`, và `alias=AlsoGood`. `emitReexportClauseFacts` gọi `nodeHasMalformedSyntax(specifier)` rồi tại lines `166-174` phát một diagnostic và bỏ toàn bộ recovered node. Vì vậy `AlsoGood`, dù nằm sau dấu phẩy như một sibling source site hợp lệ, không tạo `ExportFact` hoặc derived `ImportReexport`.

Evidence:

```text
go run ./.tmp/supervisor_p4b1_valid_sibling_probe.go
exit 1
root=(program (export_statement (export_clause (export_specifier name: (identifier)) (export_specifier name: (identifier) (ERROR) alias: (identifier))) source: (string (string_fragment))))
actual exports=1 names=[Good] diagnostics=1 imports=1
required exports=2 names=[AlsoGood Good] diagnostics=1 imports=2
exit status 1
```

Probe dùng production boundary `parser -> tsjs.Extract -> ScopeIR`, được tạo dưới repo-local `.tmp` sau canonical full build và đã bị xóa chính xác sau output. Final census xác nhận probe không còn.

Current test gap: `internal/providers/tsjs/extract_test.go:698` chỉ dùng `export { Good, Broken as } from './mixed';`, tức malformed alias nằm cuối clause. Test đó chứng minh giữ sibling phía trước nhưng không chứng minh invariant “retaining valid siblings” khi valid sibling nằm phía sau error recovery.

Why this blocks acceptance: P4-B1 claim và plan yêu cầu structured diagnostic cho malformed source-bearing site đồng thời giữ valid siblings. Actual cardinality là `1` fact / `1` compatibility thay vì `2` / `2`; đây là mất source fact tại chính extraction boundary, không phải P4-C projection hoặc Child 05 resolution.

Fix direction: production-first, thu hẹp trong re-export clause recovery owner; tách malformed `Broken as` khỏi valid `AlsoGood` theo source/AST evidence, phát đúng một malformed diagnostic, không tạo fact cho `Broken`, và vẫn tạo đúng một fact cùng một compatibility import cho mỗi `Good` và `AlsoGood`. Không mở graph/persistence/terminal resolution.

Re-review evidence required:

1. Exact production diff đóng case `export { Good, Broken as, AlsoGood } from "./mixed";` với `Exports=2`, names `Good` + `AlsoGood`, `ExportDiagnostics=1`, `Imports=2`, không có fact/import cho `Broken`.
2. Test-after-code regression trong existing focused test owner, kiểm tra exact range/selection/provenance, empty `LocalName`/`LocalDefID`, syntax-only `TargetRaw`, value meaning và compatibility cardinality cho cả hai valid siblings.
3. Fresh canonical `npm run full-build` trước fresh focused matrix; sau đó full `tsjs`, nearest `tsjs/scopeir/providers`, và `resolution/analyze` compatibility regression.
4. Fresh exact candidate/report hashes, clean index, exact `.tmp` P4-B1 census `0`, `gofmt -d` rỗng và `git diff --check` exit `0`; không detect-changes, stage, commit, push, target access, ledger/roadmap edit hay later-slice work trong Coder resubmission.

## Source-Level Clearance Notes

- `internal/providers/tsjs/imports.go`: blocked — lines `166-174` discard the recovered specifier wholesale and lose `AlsoGood`; surrounding named/star/namespace emission remains inside the authorized syntax owner.
- `internal/providers/tsjs/extract_test.go`: blocked as acceptance proof — the current malformed mixed case at line `698` has no post-error valid sibling, nên matrix không đóng exact recovery invariant.
- `internal/providers/tsjs/extract.go`: clear/preserve-only — full source read, zero candidate diff; existing P4-B `Exports`/`ExportDiagnostics`/`Imports` result wiring is sufficient.
- `internal/providers/tsjs/definitions.go`: clear/preserve-only — full source read, zero candidate diff; `DefinitionFact.Visibility` không bị ghi hoặc đổi.
- `internal/scopeir/{facts.go,kinds.go,ir.go,sort_keys.go}`: clear/preserve-only — full accepted P4-A contract read, zero diff; `ExportFact` remains syntax/local-only and owned normalization preserves immutable nested state.
- Later-slice/source scan: candidate production diff không thêm `TargetFile`, `TargetDefID`, `TargetModuleScope`, `LinkStatus`, `TransitiveVia`, barrel, cycle, ambiguity, persistence/projection hoặc public-API state.

## Evidence Checked

### Passed

- Fresh excluded graph trước graph-based review:

  ```text
  anvien analyze --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**" --json
  exit 0
  scanned/parsed/failed = 1,138/626/0
  graph = 81,978 nodes / 121,633 relationships
  ```

  Current count khác Coder pointer đúng `+1` scanned file và `+12/+12` graph entities; current manifest có durable Coder report được tạo sau Coder final graph. Candidate bytes và HEAD không đổi. Build-internal unexcluded analyze `1,839/731/0`, `112,689/155,039` không được dùng làm semantic evidence.

- Current file-detail, `stale=false`, `changedSinceAnalyze=false`:
  - `imports.go`: `173` symbols; inbound/outbound/local `7/106/38`; unresolved `472`; flows/tests `2/4`; risk HIGH.
  - `extract_test.go`: `703`; `13/364/245`; unresolved `0`; flows/tests `0/3`; file-detail risk LOW.
  - `extract.go`: `39`; `24/35/32`; unresolved `71`; flows/tests `3/6`; risk HIGH.
  - `definitions.go`: `129`; `8/107/14`; unresolved `308`; flows/tests `0/4`; risk HIGH.
  - `scopeir/facts.go`: `192`; `1200/27/174`; unresolved `5`; tests `98`; risk MEDIUM.
- Current upstream impact with `--include-tests`:
  - `imports.go`: HIGH, `17` impacted / `17` direct / `1` file / `1` flow.
  - `extract_test.go`: CRITICAL, `63` / `60` / `4` / `0`.
  - preserve-only `extract.go`: CRITICAL, `131` / `20` / `29` / `1`.
  - accepted shared contract `scopeir/facts.go`: CRITICAL, `861` / `243` / `122` / `1`.
  - chín changed/new P4-B1 syntax methods/helpers được kiểm tra riêng; mỗi symbol có risk LOW và upstream impacted count `0`.
  HIGH/CRITICAL là blast-radius warnings, không phải edit prohibition; failure probe ở trên là behavior evidence độc lập.
- Git/provenance final gate:
  - HEAD `11a37aa8ec0320dd93258c058b088d1070aa778d`; index empty.
  - Unstaged tracked paths đúng `imports.go` `369/22` và `extract_test.go` `248/2`; tổng `617` insertions / `24` deletions.
  - Untracked trước report đúng Main handoff và Coder report; Main handoff được attribution, không phải Coder drift.
  - `imports.go`: `34,219` bytes / `1,247` LF / SHA-256 `4D6A796F305D4CB9812B6600385E0215DEFF8A27F557B41559A7BF634A95C850`.
  - `extract_test.go`: `133,138` bytes / `3,057` LF / SHA-256 `7BC5C215414DFECC23F5E6B26EDC9F90BAE2826F8F8AC03A9E1BAD34A5CC9AEA`.
  - Coder report: `14,885` bytes / `229` LF / SHA-256 `1830942DB180C4833750C17A5872F244511E1D6184B1C33B172546CE2EEE22FA`.
  - Main handoff preserved: `8,715` bytes / `85` LF / SHA-256 `E3AC50340AAD0A27CE07DB0D8235AE47EDFD4556DDE5C5915647BDA11F5D4DA9`.
  - `gofmt -d internal/providers/tsjs/imports.go internal/providers/tsjs/extract_test.go` exit `0`, output empty; `git diff --check` exit `0`.
  - `.tmp/p4b1_ast_probe.go`, `.tmp/p4b1_extract_probe.go`, và hai Supervisor probes đều absent; exact P4-B1/re-export/probe census `0`. Không xóa shared/unrelated `.tmp` artifacts.
- Pre-build holder gate: `anvien doctor locks --repo E:\Anvien --json` báo `free`, lock absent. Chỉ ba global `anvien.exe mcp` đang giữ package được `npm install -g .` thay thế; exact child PIDs `2932`, `10472`, `7672` và parent PIDs `19548`, `17632`, `16448` được dừng, recount build holders `0`. Hai Codex `node.exe` runtime không liên quan build được giữ nguyên.
- Fresh canonical build:

  ```text
  npm run full-build
  exit 0
  ```

  Packaged/global CLI `1.2.8`; Web `2,943` modules; Vite `22.37s`. Allow-scripts, mixed static/dynamic import và chunk-size messages là non-failing warnings.
- Fresh validation sau full build:

  ```text
  go test -v ./internal/providers/tsjs -run '^(TestExtractTypeScriptDirectAndDefaultExportFacts|TestExtractTypeScriptLocalAliasAndTypeOnlyExportFacts|TestExtractJavaScriptDirectDefaultAndLocalExportFacts|TestExtractTypeScriptReexportSyntaxFacts|TestExtractJavaScriptReexportSyntaxFacts|TestExtractExportDiagnosticsAndLaterSliceBoundaries)$' -count=1
  exit 0; 6/6 top-level tests PASS

  go test ./internal/providers/tsjs -count=1
  exit 0

  go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers -count=1
  exit 0; 3/3 packages PASS

  go test ./internal/resolution ./internal/analyze -count=1
  exit 0; 2/2 packages PASS
  ```

### Failed

- Fresh post-build nearest-boundary valid-sibling probe: exit `1`; actual `exports=1 names=[Good] diagnostics=1 imports=1`, required `exports=2 names=[AlsoGood Good] diagnostics=1 imports=2`. Đây là blocking evidence; green checked-in matrix không bao phủ shape này.

### Not run

- `anvien detect-changes`, stage, commit, push: cấm theo delegation; Main owns only after Supervisor PASS.
- Plan/ledger/roadmap edits: cấm trong review này.
- P4-C, P4-C2, Child 05, terminal/barrel/cycle/ambiguity/public API: locked và không mở.
- `E:\cheapapp.org`: không access.
- Browser/Playwright: N/A cho non-UI provider/ScopeIR boundary.

## Invariant Closure

- Affected invariant: malformed source-bearing named re-export recovery phải phát structured diagnostic mà không làm mất bất kỳ valid sibling site nào trong cùng clause.
- Sibling surfaces checked for this finding: exact source loop, parser recovery shape, fact/diagnostic/import cardinality, current malformed-clause test, ScopeIR normalization, extract/result wiring và compatibility derivation.
- Residual unverified same-invariant surfaces: acceptance dừng ở confirmed failure này; không có cơ sở để tuyên bố residual `none` trước khi production fix và full fresh resubmission matrix hoàn tất.

## Required Fix List For Resubmission

1. Sửa đúng production recovery invariant tại named re-export clause để case `Good, Broken as, AlsoGood` giữ `Good` và `AlsoGood`, diagnostic đúng malformed site, không fabricate `Broken`.
2. Sau production fix, thêm regression test exact-field/cardinality/compatibility cho valid sibling phía sau malformed specifier.
3. Nộp fresh source/diff/hash/.tmp evidence, canonical build trước validations, focused/full/nearest/consumer regression, và durable Coder resubmission report. Không mở hoặc sửa bất kỳ later slice nào.

## Overall Evaluation

Candidate có source boundary gọn, identities/provenance sạch, canonical build và declared matrices đều PASS. Tuy nhiên zero-trust source/runtime review chứng minh claim “retaining valid siblings” là false cho một recovery shape trực tiếp trong P4-B1. Vì fact và compatibility entry bị mất ngay tại provider extraction, P4-B1 chưa thể được acceptance dù các test hiện tại xanh.

Next responsible party: Main/orchestration task `01a02046-f3e4-7bc0-ba36-d10141a55e25`. Supervisor không tự transition, không sửa candidate và không mở slice kế tiếp.
