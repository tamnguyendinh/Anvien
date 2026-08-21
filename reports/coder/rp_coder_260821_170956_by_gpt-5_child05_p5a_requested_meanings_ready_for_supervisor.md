# Child 05 / P5-A Requested Import Meanings Coder Report

## Candidate State

- Trạng thái: `READY_FOR_SUPERVISOR`.
- Đây là durable Coder handoff, không phải Supervisor verdict và không phải self-accept.
- Sole open slice: `P5-A — establish current module-request and path inputs`.
- Git basis: HEAD `0aa49c87628c9e8b2041754515d6ebf0a930d55b`, parent `d1d8eb9002ce9c449c3713de0837ac8216d17a8d`.
- Candidate chưa commit; `E5-P5A-REVIEW1`, `E5-P5A-DETECT1`, và `E5-P5A-COMMIT1` còn pending.
- P5-B/P5-C/P5-D, target access, export-table/traversal/global-rescue work, detect/commit, push/reset/checkout đều chưa mở hoặc không thực hiện.
- Bốn Main-owned R2 ledgers và bốn handoff/inventory reports có sẵn được bảo toàn; chỉ các P5-A rows của evidence/benchmark/actual-status được refresh.

## Authority And Gate

Đã đọc đầy đủ trước code:

- `AGENTS.md` và `.agents/skills/working-rules/SKILL.md`;
- coder, backend-development, planner skill và bốn planner templates;
- graph-accuracy roadmap và graph-accuracy contract;
- toàn bộ bốn Child 05 ledgers sau khi SHA-256 khớp chính xác R2;
- durable inventory report SHA-256 `82D9F651A0BF6CF13CD66F0EEF6DC310F9DAA69A4782E3769383A3294F8672DE`;
- fresh production source của bốn editable owners và các preserve-only resolver boundaries.

R2 ledger hashes tại gate trước code:

| Ledger | SHA-256 |
|---|---|
| plan | `719DD0CF1CA5442CA206409BBC6352C812EF524A006968A890648751907E2A1E` |
| evidence | `595AE33C01BAF9614DECAD0077DC27755D143DD0D834A385358501FDCA69A57B` |
| benchmark | `67E2D9E108C2FB00EE8415874A76BF242439ADFB54F2489DFA7355987A17C322` |
| actual-status | `F9460ECF1995FBECB3E0D1F4F714D34D2873BB47C1B7FB65F3E162B496DB4D2F` |

`anvien --help` được đọc; mandatory pre-edit `anvien analyze --force` PASS tại C worktree với `1,945 / 736 / 0` scanned/parsed/failed và graph `114,757 / 157,572`. Main xác nhận delta so với inventory là một report R2 đã biết, không thay đổi boundary hoặc mở planner lần nữa.

## Invariant Family Map

| Surface | Authority / owner | Candidate result | Boundary |
|---|---|---|---|
| Import request contract | `internal/scopeir/facts.go::ImportFact` | thêm `RequestedMeanings []ExportMeaning` và `TypeOnly bool` | chỉ requested semantic input; dormant `Target*` giữ nguyên |
| TS/JS source import syntax | `internal/providers/tsjs/imports.go::emitImportStatement` | default/named/alias/namespace và statement/inline type-only map đúng R2 | side-effect-only import không mở rộng |
| Re-export semantic SSOT | accepted `ScopeIR.ExportFact` | không đổi | compatibility `ImportFact` chỉ path compatibility, fields mới empty |
| Canonical ownership | `ScopeIR.Normalized`, `NormalizeInPlace`, `NormalizeOwned` | deep clone, sort, dedupe requested meanings | accepted export normalization không đổi |
| Deterministic ordering | `compareImport` | compare requested meanings và `TypeOnly` | existing keys giữ nguyên |
| Module/file result | `internal/resolution/indexes.go::resolvedImport` | hash unchanged, behavior preserved | P5-B/P5-C export lookup chưa mở |
| Language path strategies | `internal/resolution/import_resolution.go` | hash unchanged | other-language strategy preserve-only |
| Syntactic dependency output | resolver metrics và persisted `IMPORTS` | `5,072 / 5,072 / 5,088`, delta `0 / 0 / 0` | không tạo traversal edge mới |

Forbidden fallback status:

- không biến physical definitions thành implicit exports;
- không thêm global-name rescue;
- không kích hoạt/xóa dormant `ImportFact.Target*`;
- không đưa re-export meaning sang compatibility imports;
- không thêm side-effect-only facts;
- không sửa target hoặc truy cập `E:\cheapapp.org`.

## Fresh Relationship And Impact Evidence

File-detail tại pre-edit graph, mọi row `stale=false`, `changedSinceAnalyze=false`:

| Editable owner | Related files | Symbols | In / out / local | Linked flows / tests | File risk |
|---|---:|---:|---:|---:|---|
| `internal/scopeir/facts.go` | 247 | 192 | 1,232 / 27 / 174 | 0 / 100 | MEDIUM |
| `internal/providers/tsjs/imports.go` | 17 | 185 | 7 / 107 / 40 | 1 / 4 | HIGH |
| `internal/scopeir/ir.go` | 245 | 57 | 1,147 / 38 / 89 | 3 / 98 | HIGH |
| `internal/scopeir/sort_keys.go` | 242 | 122 | 254 / 128 / 54 | 2 / 97 | HIGH |

Complete HIGH/CRITICAL upstream blast radius:

| Target | Risk | Impacted | Direct | Files | Modules | Processes / flows |
|---|---|---:|---:|---:|---:|---:|
| `internal/scopeir/facts.go` | CRITICAL | 867 | 244 | 123 | N/A | 1 affected flow |
| `internal/providers/tsjs/imports.go` | HIGH | 18 | 18 | 1 | N/A | 0 affected flows |
| `internal/scopeir/ir.go` | CRITICAL | 610 | 442 | 76 | N/A | 1 affected flow |
| `internal/scopeir/sort_keys.go` | CRITICAL | 32 | 24 | 5 | N/A | 1 affected flow |
| `scopeir.ImportFact` | CRITICAL | 624 | 18 | 73 | 25 | 67 processes |
| `ScopeIR.Normalized` | CRITICAL | 21 | 11 | 6 | 4 | 15 processes |

Exact editable methods `collector.emitImportStatement`, `collector.addImport`, `collector.addSourceExportFact`, `ScopeIR.NormalizeInPlace`, `ScopeIR.NormalizeOwned`, và `compareImport` đều LOW. Focused test owners were separately gated: `extract_test.go` CRITICAL `63` impacted / `4` files / `0` flows; `scopeir_test.go` MEDIUM `9` / `1` / `0`.

## Production Implementation

Production được sửa trước tests, chỉ ở bốn owner R2:

1. `internal/scopeir/facts.go:197`
   - thêm `RequestedMeanings []ExportMeaning` với JSON `requestedMeanings,omitempty`;
   - thêm `TypeOnly bool` với JSON `typeOnly,omitempty`;
   - không thay đổi dormant resolution-output fields.
2. `internal/providers/tsjs/imports.go:29`
   - plain default/named/alias request canonical allowed-set `{namespace,type,value}` sau normalization;
   - statement-level hoặc inline `type`/`typeof` request `{type}` và `TypeOnly=true`;
   - plain namespace có `ImportedName=""`, `{namespace}`;
   - type-only namespace có `ImportedName=""`, `{type}`, `TypeOnly=true`;
   - compatibility re-export calls pass `nil,false`;
   - side-effect-only early return giữ nguyên.
3. `internal/scopeir/ir.go:70`
   - `Normalized` và `NormalizeOwned` deep-clone `RequestedMeanings`;
   - `NormalizeInPlace` sort và dedupe từng requested allowed-set.
4. `internal/scopeir/sort_keys.go:39`
   - `compareImport` gồm requested meanings và `TypeOnly` trước target/ID tie-breakers.

Production diff: `119` insertions / `26` deletions. Test diff: `154` insertions / `0` deletions.

Candidate SHA-256:

| File | SHA-256 |
|---|---|
| `internal/scopeir/facts.go` | `4548FF2C312F329B0672A5CF5999F112F647CBAC6C060E489B08D6B7C1063646` |
| `internal/providers/tsjs/imports.go` | `538FD9723AB77C672DD35F7B5D17D01C565167757FA73FFCF23DE05E0A1640D4` |
| `internal/scopeir/ir.go` | `4BCE48A4E490810707708C0D199C3B837899347F19DD084DB6142F59DE7D6ED6` |
| `internal/scopeir/sort_keys.go` | `741F83729AB1CED8E09945A74779C10E18BD4BF526EA0AA03A228236FFEB06B2` |
| `internal/providers/tsjs/extract_test.go` | `CCA89257D3E428D39DB8BDB3E61136ABACE135A8F20D3B800FC5563B01CE6D59` |
| `internal/scopeir/scopeir_test.go` | `AC757B1FB75E529CAF29299620E80C0CB46144B2327C574D5B62AFD378249964` |

Preserve-only hashes remain exact:

- `internal/resolution/indexes.go`: `AA19B9D543012309A90974089BACBD0A122594C7481FFB0790DE5C01F3D3D76B`;
- `internal/resolution/import_resolution.go`: `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413`.

## Focused Test Evidence

New tests were added only after production behavior:

- `internal/providers/tsjs/extract_test.go:92` — `TestExtractImportRequestedMeaningsAndTypeOnly`:
  - normal default, named, alias;
  - plain namespace with no exported-name request;
  - statement-level type-only default/named/alias;
  - inline type-only named/alias;
  - type-only namespace;
  - side-effect-only preservation;
  - compatibility re-export requested fields empty.
- `internal/scopeir/scopeir_test.go:88` — `TestImportRequestedMeaningsCanonicalizeCloneAndOrder`:
  - deep clone for `Normalized` and `NormalizeOwned`;
  - canonical sort/dedupe;
  - deterministic order distinguishes meanings and `TypeOnly`;
  - JSON round-trip preservation.

Commands:

```text
go test ./internal/providers/tsjs ./internal/scopeir -run '^(TestExtractImportRequestedMeaningsAndTypeOnly|TestImportRequestedMeaningsCanonicalizeCloneAndOrder)$' -count=1
PASS: tsjs 0.260s; scopeir 0.578s

go test ./internal/providers/tsjs ./internal/scopeir -count=1
PASS: tsjs 0.194s; scopeir 0.540s
```

## Full Build Evidence

Build-holder gate:

- `anvien doctor locks --repo <C-worktree> --json`: `status=free`, lock absent;
- `anvien doctor processes --json`: only editor-owned MCP processes, explicitly expected to remain running; no build/analyze holder.

Canonical full build initially exposed two environment prerequisites, both recorded and resolved without source/config edits:

1. Direct C-path run failed because package runtime wrote unquoted `CGO_LDFLAGS=-L<C:\Users\TAM NGUYEN\...>` and cgo split the path.
2. The exact worktree was exposed temporarily as unused drive `X:`; the same `npm run full-build` passed Go packaging but first revealed missing `anvien-web/node_modules`.
3. `npm ci` under `anvien-web` used the existing lockfile and exited `0` (`536` packages). Its audit reported the existing dependency inventory (`1` low, `6` moderate, `9` high, `2` critical); no audit-fix or dependency source change was performed.
4. The same canonical `npm run full-build` through `X:` exited `0`: Go runtime, package build/global install, web build, launcher/server, CLI version, and final analyze passed.
5. `X:` was removed in `finally`. The global npm junction was restored from `X:\anvien` to the real C package with `npm install -g . --ignore-scripts`; `anvien version` returned `1.2.8`.

Successful built-CLI analyze inside full build:

```text
scanned=1,945
parsed_code=736
failed=0
nodes=114,788
relationships=157,690
```

No build, package, web, launcher, dependency, or config source entered the tracked diff.

## Nearest Real Non-UI Boundary

```text
go test ./internal/resolution ./internal/analyze ./internal/cli -count=1
PASS
resolution: 0.294s
analyze: 3.930s
cli: 119.619s
```

This proves the unchanged module/file resolver consumes the extended facts, the analyze pipeline completes, and the real CLI command boundary remains operational after the built runtime change.

## Unaffected-Language Regression

```text
go test ./internal/providers/... -count=1
exit 1
```

All provider packages passed except two previously accepted out-of-slice parity baselines:

- C#: `TestResolveCSharpGraphParityCounts`, expected `ACCESSES:2`, got none;
- Dart: `TestResolveDartGraphParityCounts`, expected `ACCESSES:2`, got `1`.

The exact same failures are recorded in `reports/Supervisor/rp_supervisor_260820_221058_by_gpt-5_child04_p4a_export_fact_boundary.md` and the Child 04 Coder report. P5-A changes no access extraction/emission. Focused unaffected extractors both pass:

```text
go test ./internal/providers/csharp ./internal/providers/dart -run '^TestExtract(CSharp|Dart)ScopeIRParityFixture$' -count=1
PASS: csharp 0.318s; dart 0.313s
```

All other provider packages passed in the aggregate run, and repository source inspection confirms no non-TS/JS writer assigns `RequestedMeanings` or `TypeOnly`.

## Post-Change Three-Count Proof

Built CLI command on the real C path:

```text
anvien analyze --force --benchmark-json .tmp\p5a-postchange-analyze-r2.json --benchmark-label p5a-postchange-r2
exit 0
files: 1,945 / 736 / 0
graph: 114,788 / 157,690
```

The parsed-code corpus is unchanged at `736`; the one extra scanned file is the Main-classified non-code R2 report.

| Metric | Baseline | Post-change | Delta |
|---|---:|---:|---:|
| physical target-file resolutions (`resolution.ImportsResolved`) | 5,072 | 5,072 | 0 |
| resolver-emitted syntactic `IMPORTS` (`resolution.FinalizedImportsEmitted`) | 5,072 | 5,072 | 0 |
| persisted graph-wide `IMPORTS` | 5,088 | 5,088 | 0 |

Persisted query:

```text
MATCH ()-[r]->() WHERE r.type = 'IMPORTS' RETURN count(r) AS imports
row_count=1
imports=5,088
```

Artifacts:

- `.tmp/p5a-postchange-analyze-r2.json`: `9,378` bytes / `365` LF / SHA-256 `4B66AF8DE0197357D5DBCAC0BDB109D56CD4130E32BD92C6E4CC1FEF2B89518D` (debug-only benchmark capture; metrics copied into durable ledgers);
- `.anvien/graph.json`: SHA-256 `D0E49082B0CB016E93E5A77818641A995060789E5D76AE40B198D1E44614F790`.

## Ledger Refresh

- Evidence: `E5-P5A-IMPACT1`, `E5-P5A-INPUT1`, `E5-P5A-COUNT1`, `E5-P5A-SRC1`, `E5-P5A-BUILD1`, `E5-P5A-TEST1` recorded.
- Benchmark: parsed-code corpus and all three P5-A denominators updated with post-change latest values and delta `0`.
- Actual status: R3 appended; requested namespace name `partial -> correct`; requested meaning/type-only `missing -> correct`; candidate state set to `READY_FOR_SUPERVISOR`.
- `E5-P5A-REVIEW1`, `E5-P5A-DETECT1`, `E5-P5A-COMMIT1` remain pending.
- Plan checklist was not self-ticked; P5-B+ remains locked.

## E2E Verification

```text
E2E Verification:
  [PASS] Compiled: npm run full-build -> exit 0 after documented path/dependency prerequisites
  [PASS] Runtime: built CLI analyze on real C path -> 1,945/736/0 and 114,788/157,690
  [PASS] Happy path: TS/JS source import -> ImportFact -> canonical ScopeIR -> resolver/analyze/CLI -> exact requested meanings and unchanged module/file result
  [PASS] Edge case: statement/inline type-only, namespace without exported name, compatibility re-export empty fields, side-effect-only preservation, clone/dedupe/order/round-trip
```

Invariant family: P5-A source-written module request semantics before export lookup.

Authority / SSOT: R2 Child 05 plan set plus accepted Child 04 `ScopeIR.ExportFact` contract.

Sibling surfaces checked: default/named/alias/namespace; statement/inline type-only; compatibility re-export; side-effect-only; all normalization paths; deterministic compare; resolution/analyze/CLI; non-TS/JS providers; three count denominators.

Legacy fallback status: no new fallback; physical definition lookup, export traversal, and global-name rescue remain untouched and deferred to their locked slices.

Stale tests/helpers/plans updated: two focused tests added; only P5-A evidence/benchmark/actual-status rows refreshed; no stale test or helper was removed.

Residual unverified surfaces: none inside the authorized P5-A Coder scope. Independent Supervisor review, detect-changes, and commit are workflow gates still pending, not Coder acceptance claims.

## Handoff

`READY_FOR_SUPERVISOR`

Supervisor should independently verify:

1. exact four-production/two-test diff and hashes;
2. source form matrix and compatibility/non-TS empty-field boundary;
3. deep clone/canonicalization/comparator behavior;
4. full-build provenance through the temporary drive alias and its cleanup;
5. nearest resolver/analyze/CLI boundary;
6. unaffected-language baseline classification;
7. `5,072 / 5,072 / 5,088` post-change counts and deltas `0 / 0 / 0`;
8. P5-B+, target, detect, commit, and unrelated Main-owned paths remain unopened.
