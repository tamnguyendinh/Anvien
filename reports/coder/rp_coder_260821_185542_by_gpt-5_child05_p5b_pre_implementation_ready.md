# Child 05 P5-B Export Tables — Pre-Implementation Inventory

## Status

- `PRE_IMPLEMENTATION_READY / READY_FOR_MAIN_AUTHORIZATION`.
- Đây là durable inventory handoff, không phải production candidate, Supervisor verdict, hay self-acceptance.
- Main có thể dùng report này để resume đúng Supervisor lane đã park; Coder lane này chưa gọi Supervisor.
- Sole scope: P5-B xây export tables từ accepted Child 04 `ScopeIR.ExportFact` và accepted P5-A module/file results.
- P5-C/P5-D, terminal traversal/binding, global-name rescue, `E:\cheapapp.org`, build, test, detect-changes, commit, push/reset/checkout đều chưa mở hoặc chưa chạy.

## Repository Boundary

- Cwd và repo root đã xác minh: `E:\Anvien` / `E:/Anvien`.
- HEAD: `40ea0095a79084a3c6805cf5d5f46108926d1dca` (`docs(plan): open P5-B after P5-A`).
- Accepted P5-A source commit: `2560f914334e65961f755febdda6585840a4260e`.
- Sáu untracked Main orchestration handoff reports trong `reports/Investigation/` được giữ nguyên; lane này không đọc, sửa, stage, hay xóa chúng.
- Không tạo output/temp ngoài `E:\Anvien`.

## Authority Read

Đã đọc đầy đủ trước inventory:

- `E:\Anvien\AGENTS.md`;
- `working-rules`, `coder`, `backend-development`, và `debugging` skills;
- toàn bộ bốn current P5 files: plan, actual-status, evidence, benchmark;
- accepted P5-A Coder report, canonical E full-build addendum, và final Supervisor PASS;
- current accepted fact owners `internal/scopeir/facts.go`, `internal/scopeir/kinds.go`, `internal/scopeir/ir.go`;
- current owner source `internal/resolution/indexes.go`.

Accepted fact invariant:

- `ExportFact` là sole re-export semantic SSOT; seven source-form kinds, canonical meaning lanes, type-only state, local/source names, optional `LocalDefID`, source-written `TargetRaw`, source-side `TargetExportedName`, ranges, và provenance.
- `ExportFact` không chứa terminal resolution/public-API state.
- Compatibility re-export `ImportFact` chỉ cung cấp path compatibility; P5-A requested fields của các fact này vẫn empty.
- Physical definitions không tự động trở thành export entries.

## Fresh Graph Gate

Một graph refresh duy nhất được chạy từ E trước graph work:

```text
E:\Anvien\anvien\bin\anvien.exe analyze E:\Anvien --force
PASS: scanned=1,951 parsed_code=736 failed=0
graph: 114,852 nodes / 157,754 relationships
path: E:\Anvien\.anvien\graph.json
indexed/current commit: 40ea0095a79084a3c6805cf5d5f46108926d1dca
```

Sau refresh, `anvien status` báo `up-to-date`. Không lặp analyze.

## Exact Owner Decision

Existing storage/integration owner:

- `internal/resolution/indexes.go::workspace` — sở hữu toàn bộ in-memory workspace indices.
- `internal/resolution/indexes.go::buildWorkspace` — điểm duy nhất normalize toàn bộ `[]ScopeIR`, giữ `w.files`, index module/file/scopes/definitions/import results, rồi gọi `resolveImports`.

Source evidence:

- `buildWorkspace` đã nhìn thấy mọi normalized `ir.Exports` nhưng hiện không index chúng.
- Không có first-class export table trong package `internal/resolution`.
- `resolveImports` đã tạo `resolvedImport` tách module/file result khỏi optional physical `TargetDef`; source-bearing compatibility imports giữ exact `Fact`, `TargetFile`, và `TargetFiles` để P5-B có thể reuse path results mà không gọi path resolution lần hai.
- `resolveImportedDef` vẫn chỉ tìm physical definitions. Đây là preserve-only behavior trong P5-B; terminal export traversal thuộc P5-C.

Implementation owner sau khi Main authorize:

- Semantic table types/construction: file chuyên trách mới `internal/resolution/export_tables.go` để giữ một trách nhiệm duy nhất.
- Minimal integration-only edit: thêm storage field vào `workspace` và một wiring call trong `buildWorkspace` tại `internal/resolution/indexes.go`.
- Không thêm table-construction logic vào provider/ScopeIR contract, `resolveImports`, `resolveImportedDef`, `import_resolution.go`, graph emission, persistence, hay readers trong P5-B.

## File Detail And Upstream Impact

Command basis: fresh graph ở HEAD `40ea0095`; mọi selected row `stale=false`, `changedSinceAnalyze=false`.

`anvien file-detail internal/resolution/indexes.go --repo E:\Anvien --json`:

| Evidence | Result |
|---|---:|
| related files | 51 |
| symbols | 319 |
| inbound / outbound / local relationships | 207 / 95 / 371 |
| linked flows / tests | 13 / 28 |
| unresolved | 439 (`170` calls, `269` refs, `0` imports) |
| file risk | HIGH |
| current SHA-256 | `AA19B9D543012309A90974089BACBD0A122594C7481FFB0790DE5C01F3D3D76B` |

Exact upstream impacts:

| Exact target | Risk | Impacted symbols | Files | Modules | Processes | Depth 1/2/3 |
|---|---|---:|---:|---:|---:|---:|
| `workspace` struct UID at `indexes.go:45` | CRITICAL | 70 | 9 | 4 | 49 | 63 / 4 / 3 |
| `buildWorkspace` at `indexes.go:72` | CRITICAL | 8 | 6 | 5 | 25 | 2 / 3 / 3 |

`workspace` affected files: `internal/analyze/analyze.go`, `internal/cli/command.go`, `internal/graphaccuracy/access_candidate.go`, `internal/resolution/access_audit.go`, `emit.go`, `import_resolution.go`, `indexes.go`, `resolve.go`, và `types.go`.

`buildWorkspace` affected files: `cmd/access-candidate-audit/main.go`, `internal/analyze/analyze.go`, `internal/cli/command.go`, `internal/graphaccuracy/access_candidate.go`, `internal/resolution/access_audit.go`, và `internal/resolution/resolve.go`.

Blast-radius conclusion: CRITICAL/HIGH là scope warning. P5-B phải giữ edit ở dedicated table owner cộng minimal `workspace/buildWorkspace` wiring, và validation sau authorization phải bao phủ resolver, analyze/CLI, access-audit, graphaccuracy, cùng preservation counts.

## Invariant Family Map

| Surface | P5-B disposition |
|---|---|
| accepted `ScopeIR.Exports` | immutable input / inspect-only |
| accepted P5-A `resolvedImport` module/file results | immutable path result / consume only |
| per-file explicit entries | build deterministically from non-star `ExportFact` records keyed by source `FilePath` and `ExportedName` |
| star adjacency | build deterministically from `ExportStar` facts plus matching accepted module/file result; retain the source fact/provenance |
| local/default/alias/re-export/namespace/type-only metadata | preserve exact accepted fact bytes/meaning; do not reinterpret syntax |
| physical definitions | never synthesize implicit exports |
| zero-physical barrel | table must remain non-empty when accepted export facts exist |
| terminal traversal, precedence resolution, cycle/ambiguity outcome | P5-C locked |
| syntactic `IMPORTS` and physical path lookup | preserve counts and behavior |
| graph/persistence/readers/target | no P5-B edit |

Forbidden fallbacks: no physical-definition export inference, no repository-global same-name rescue, no target special case, no second re-export semantic source, no side-effect-only expansion, and no terminal result/proof field in the table slice.

## Authorized-Next-Action Request

Main authorization is the only remaining gate before production code. If authorized, the bounded implementation tranche is:

1. Add dedicated `internal/resolution/export_tables.go` with deterministic explicit-entry and star-adjacency construction from immutable accepted facts plus existing module/file results.
2. Add only the required `workspace` storage and `buildWorkspace` wiring in `internal/resolution/indexes.go`.
3. Implement production code first; add focused tests only afterward.
4. Run only the canonical full build entrypoint required by Main, then nearest real resolver/CLI boundary, P5-A three-count preservation, Supervisor, detect-changes, and isolated P5-B commit.

No source edit, build, test, detect-changes, or commit was performed during this inventory tranche.

## Handoff

`PRE_IMPLEMENTATION_READY`

Main should either authorize the exact two-owner implementation boundary above or return a precise owner/invariant correction. P5-C/P5-D and target remain locked regardless of this inventory readiness.
