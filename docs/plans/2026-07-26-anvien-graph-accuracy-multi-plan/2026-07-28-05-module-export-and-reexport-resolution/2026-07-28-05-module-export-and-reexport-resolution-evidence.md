# Anvien Module Export And Re-Export Resolution Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md`

## Evidence Rules

- The originating problem report is the problem origin. Its proposed architecture is not implementation authority.
- The causal synthesis and final Supervisor report are accepted bounded verification. They do not establish global prevalence or remediation design.
- Current source, fresh Anvien graph evidence, runtime behavior, and accepted predecessor outputs determine implementation scope.
- A pending evidence ID is a required target, not proof.
- Every slice requires impact, production source, build, behavior, boundary, Supervisor, detect-changes, and commit evidence appropriate to its scope.
- Record exact commands, commit, input corpus/config, observed result, affected counts, and artifact path. Do not replace exact target-site proof with aggregate graph counts.
- Long measurements belong in the benchmark ledger.

### Evidence ID Naming

Evidence IDs use `E<phase>-<item>-<kind><n>` and remain stable across all four Child 05 files. `E0` maps to P0; `E5` maps to P5 and closure.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-RULE1`: read `E:\Anvien\AGENTS.md` in full. It requires planner use, graph refresh, file-detail/impact before implementation edits, code before tests, full build, truthful boundary validation, Supervisor acceptance, detect-changes, and per-slice commits.
- `E0-P0A-SKILL1`: read `.agents/skills/planner/SKILL.md` and all four planner templates in full before rewriting this four-file set.
- `E0-P0A-ORIGIN1`: read `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` in full. It is the problem origin and records the target barrel acceptance (`2/2` terminal calls), but labels its architecture as DRAFT; only its findings and acceptance targets are used here.
- `E0-P0A-VERIFY1`: read `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md` in full. C6 confirms that both target module paths resolve, while `resolveImportedDef` stops at physical definitions in the barrel; the direct-import control changes imports `10 -> 10`, import uses `1 -> 3`, resolved calls `37 -> 39`, and unresolved references `542 -> 540`.
- `E0-P0A-VERIFY2`: read `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md` in full. Verdict PASS applies only to the bounded investigation; broader module semantics and remediation remain unapproved.
- `E0-P0A-GRAPH1`: after graph refresh in the shared workspace, `anvien status` reported indexed/current commit `238aec06d28286acc0fca0f4e6a69f9eb4ff6a49` and up-to-date status. Current Child 05 file-detail checks report graph `stale=false`, `changedSinceAnalyze=false`, analyzed at `2026-08-09T19:19:54Z`.
- `E0-P0A-SRC1`: full source read of `internal/scopeir/facts.go` and `internal/providers/tsjs/imports.go`. Current `ImportFact` carries source file, kind, local/imported/alias names, raw/resolved target fields, target definition, transitive path, and link status. TS re-export syntax is currently emitted through import-shaped facts; the accepted Child 04 output contract is still a dependency.
- `E0-P0A-SRC2`: full source read of `internal/resolution/import_resolution.go` and the module-file lookup section of `internal/resolution/indexes.go`. Current code already has relative/index TypeScript candidates and many language-specific strategies; the bounded target reaches the barrel. No source evidence supports replacing all package/path resolution before P5-A inventory.
- `E0-P0A-SRC3`: full source read of `internal/resolution/indexes.go`. `resolveImports` calls `resolveImportedDef`, which searches only `defsByFile[targetFile]` by name and has no export table or re-export traversal. A missing target definition prevents the import binding.
- `E0-P0A-SRC4`: full source read of `internal/resolution/resolve.go`. `resolveCall` tries scoped/same-file/same-package paths and then a repository-global name lookup; low-confidence global matches are emitted as unresolved, so an explicit import failure is not retained as an export-boundary result.
- `E0-P0A-FD1`: `anvien file-detail internal/resolution/indexes.go --repo E:\Anvien --json` reported 46 related files, 192 symbols, 164 inbound, 93 outbound, 369 local relationships, 29 linked flows, 23 linked tests, and high file risk.
- `E0-P0A-FD2`: `anvien file-detail internal/providers/tsjs/imports.go --repo E:\Anvien --json` reported 15 related files, 21 symbols, 5 inbound, 26 outbound, 1 local relationship, 1 linked flow, 3 linked tests, and high file risk.
- `E0-P0A-FD3`: `anvien file-detail internal/resolution/import_resolution.go --repo E:\Anvien --json` reported 33 related files, 68 symbols, 21 inbound, 59 outbound, 7 local relationships, 1 linked flow, 17 linked tests, and high file risk.
- `E0-P0A-FD4`: `anvien file-detail internal/resolution/resolve.go --repo E:\Anvien --json` reported 50 related files, 40 symbols, 77 inbound, 118 outbound, 18 local relationships, 21 linked flows, 26 linked tests, and high file risk.
- `E0-P0A-IMPACT1`: upstream impact for `workspace.resolveImportedDef` is CRITICAL: 4 impacted symbols, 3 affected files, 1 module, 16 processes; selected graph evidence was not stale.
- `E0-P0A-IMPACT2`: upstream impact for `workspace.resolveImports` is CRITICAL: 6 impacted symbols, 5 affected files, 3 modules, 18 processes.
- `E0-P0A-IMPACT3`: upstream impact for `buildWorkspace` is CRITICAL: 8 impacted symbols, 6 affected files, 5 modules, 28 processes.
- `E0-P0A-IMPACT4`: upstream impact for `resolveCall` is CRITICAL: 6 impacted symbols, 4 affected files, 4 modules, 35 processes.
- `E0-P0A-DEPEND1`: Child 05 consumes the accepted Child 04 module/import/export facts. P5-A remains blocked until the predecessor plan, evidence, benchmark, actual-status, Supervisor result, and slice commits are complete.
- `E0-P0A-SCOPE1`: the campaign correction requires Child 05 to use only persistence/readers identified by the current Child 02 impact inventory and fresh P5-D evidence; no fixed all-reader denominator is authorized.
- `E0-P0A-BOUNDARY1`: accepted target boundary is `E:\cheapapp.org`, HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, graph at `E:\cheapapp.org\.anvien\graph.json`, graph hash `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`, 84,807 nodes, and 114,125 relationships. Target source remains preserve-only.
- `E0-P0A-STATUS1`: actual status classifies bounded module path as correct, module-input contract as partial, export table as missing, physical-definition lookup/global-name rescue as wrong, terminal binding as unbound, and affected readers as blocked pending exact evidence.

## E5 - P5 Evidence

### P5-A - Current module-request and path inputs

Fresh inventory artifact: `reports/coder/rp_coder_260821_161136_by_gpt-5_child05_p5a_current_input_inventory_blocked_for_plan_refresh.md`, `21,009` bytes / `295` LF / SHA-256 `82D9F651A0BF6CF13CD66F0EEF6DC310F9DAA69A4782E3769383A3294F8672DE`. The artifact records an inventory blocker, not completion or acceptance; Main resolved only that blocker through the bounded R2 planner refresh.

Coder candidate report: `reports/coder/rp_coder_260821_170956_by_gpt-5_child05_p5a_requested_meanings_ready_for_supervisor.md`, `15,237` bytes / `283` LF / SHA-256 `C2599B890BB75D0783DFAB6F643F42522C0D1E3C163E490BAE5A287B6F5C7968`. Status is `READY_FOR_SUPERVISOR`, not acceptance.

- `E5-P5A-IMPACT1` — recorded: fresh graph at HEAD `0aa49c87628c9e8b2041754515d6ebf0a930d55b`, `114,738` nodes / `157,553` relationships, with every selected file `stale=false` and `changedSinceAnalyze=false`. File-detail covers `internal/scopeir/facts.go` (`247` related), `internal/providers/tsjs/imports.go` (`17`), `internal/scopeir/ir.go` (`245`), `internal/scopeir/sort_keys.go` (`242`), `internal/resolution/indexes.go` (`51`), and `internal/resolution/import_resolution.go` (`34`). Exact-symbol impact covers fourteen candidates; the primary shared-contract result is `scopeir.ImportFact` CRITICAL with `624` impacted symbols, `73` files, `25` modules, and `67` processes. The full HIGH/CRITICAL file/symbol blast radius is retained in the inventory artifact. The mandatory R2 pre-edit refresh rebuilt the graph to `114,757 / 157,572` after the known additional report, then reconfirmed the four editable owners at `247 / 17 / 245 / 242` related files and the same complete HIGH/CRITICAL results: `facts.go` CRITICAL `867` symbols / `123` files, `imports.go` HIGH `18 / 1`, `ir.go` CRITICAL `610 / 76`, `sort_keys.go` CRITICAL `32 / 5`, `ImportFact` CRITICAL `624 / 73 / 25 modules / 67 processes`, and `ScopeIR.Normalized` CRITICAL `21 / 6 / 4 / 15`; all exact rows were `stale=false`, `changedSinceAnalyze=false`.
- `E5-P5A-INPUT1` — recorded: before implementation, source-written default/named/alias imported names and raw module text were present; namespace stored a module basename rather than an exported-name request; statement-level and inline type-only syntax plus requested meaning were absent. `resolvedImport` already separated a valid module/file result from optional physical-definition lookup. The R2 candidate now records `RequestedMeanings []ExportMeaning` as a canonical allowed-set and `TypeOnly bool` exactly for source-written TS/JS imports: normal default/named/alias `{namespace,type,value}` after canonical ordering, statement-level/inline type-only `{type}`, plain namespace no exported name plus `{namespace}`, and type-only namespace no exported name plus `{type}`. Accepted `ScopeIR.ExportFact` remains the sole re-export semantic source; compatibility re-export imports and non-TS/JS writers retain empty requested fields; side-effect-only imports remain unchanged.
- `E5-P5A-COUNT1` — recorded before and after on the same `736` parsed-code corpus. Baseline and post-change candidate both report `resolution.ImportsResolved=5,072` physical target-file resolutions, `resolution.FinalizedImportsEmitted=5,072` resolver-emitted syntactic `IMPORTS`, and final persisted graph-wide `IMPORTS=5,088` from `MATCH ()-[r]->() WHERE r.type = 'IMPORTS' RETURN count(r)`. Delta is `0 / 0 / 0`. Post-change built-CLI analyze exited `0` with `1,945 / 736 / 0` scanned/parsed/failed and graph `114,788 / 157,690`; the additional scanned non-code report is the Main-classified R2 artifact and does not change the parsed-code corpus. Debug benchmark `.tmp/p5a-postchange-analyze-r2.json` is `9,378` bytes / `365` LF with SHA-256 `4B66AF8DE0197357D5DBCAC0BDB109D56CD4130E32BD92C6E4CC1FEF2B89518D`; current graph SHA-256 is `D0E49082B0CB016E93E5A77818641A995060789E5D76AE40B198D1E44614F790`.
- `E5-P5A-SRC1` — recorded: production diff is limited to `internal/scopeir/facts.go`, `internal/providers/tsjs/imports.go`, `internal/scopeir/ir.go`, and `internal/scopeir/sort_keys.go`; focused tests are limited to `internal/providers/tsjs/extract_test.go` and `internal/scopeir/scopeir_test.go`. Production implements the two-field contract, exact TS/JS syntax mapping, compatibility `nil,false`, namespace empty exported-name request, owned deep clone, sort/dedupe, and deterministic compare. Candidate SHA-256 values are respectively `4548FF2C312F329B0672A5CF5999F112F647CBAC6C060E489B08D6B7C1063646`, `538FD9723AB77C672DD35F7B5D17D01C565167757FA73FFCF23DE05E0A1640D4`, `4BCE48A4E490810707708C0D199C3B837899347F19DD084DB6142F59DE7D6ED6`, `741F83729AB1CED8E09945A74779C10E18BD4BF526EA0AA03A228236FFEB06B2`, `CCA89257D3E428D39DB8BDB3E61136ABACE135A8F20D3B800FC5563B01CE6D59`, and `AC757B1FB75E529CAF29299620E80C0CB46144B2327C574D5B62AFD378249964`; preserve-only resolver hashes remain `AA19B9D543012309A90974089BACBD0A122594C7481FFB0790DE5C01F3D3D76B` and `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413`.
- `E5-P5A-BUILD1` — canonical repository full build PASS on the authoritative checkout. Command: `powershell -ExecutionPolicy Bypass -File .\scripts\full-build.ps1`, cwd `E:\Anvien`, exit `0`. The script built the packaged Go runtime, canonical CLI `E:\Anvien\anvien\bin\anvien.exe`, native runtime `E:\Anvien\anvien\bin\lbug_shared.dll`, launcher/server/web bundle, global npm MCP install, version check, and final `anvien analyze . --force` with graph `114,842` nodes / `157,744` relationships at `E:\Anvien\.anvien\graph.json`. No `X:`/`subst`/alternate output was used and no build/config source entered the diff.
- `E5-P5A-TEST1` — recorded: exact new tests `TestExtractImportRequestedMeaningsAndTypeOnly` and `TestImportRequestedMeaningsCanonicalizeCloneAndOrder` passed; complete `go test ./internal/providers/tsjs ./internal/scopeir -count=1` passed; nearest non-UI `go test ./internal/resolution ./internal/analyze ./internal/cli -count=1` passed, including CLI in `119.619s`. `go test ./internal/providers/... -count=1` passed every provider except the already-recorded out-of-slice C#/Dart `ACCESSES` parity baselines (`2 -> 0`, `2 -> 1`) documented by Child 04 Supervisor evidence; the focused unaffected extractors `TestExtractCSharpScopeIRParityFixture` and `TestExtractDartScopeIRParityFixture` both passed. No other-language production writer sets the two new fields.

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E5-P5A-IMPACT1` | fresh file-detail/impact for every exact editable owner after Child 04 handoff | recorded — six files, fourteen exact symbols, complete HIGH/CRITICAL inventory |
| `E5-P5A-INPUT1` | current input/result manifest from source fact through module/file lookup, including requested name/meaning | recorded — R2 contract authorized before code |
| `E5-P5A-COUNT1` | absolute physical path-resolution and syntactic `IMPORTS` counts before/after change | recorded — `5,072 / 5,072 / 5,088`, delta `0 / 0 / 0` on `736` parsed-code files |
| `E5-P5A-SRC1` | production diff limited to proved module-input/path gaps | recorded — exact four production owners plus two focused test owners; resolver/path owners preserved |
| `E5-P5A-BUILD1` | full repository build after production code | recorded — canonical `scripts/full-build.ps1` exit `0` from `E:\Anvien`, canonical CLI/runtime outputs in `E:\Anvien\anvien\bin`, final analyze completed |
| `E5-P5A-TEST1` | focused current-path and unaffected-language regression after code | recorded — focused, owner, resolver/analyze/CLI, and unaffected extractor boundaries pass; two aggregate provider failures match accepted out-of-slice baseline |
| `E5-P5A-REVIEW1` | Supervisor PASS for P5-A | recorded — `E:\Anvien\reports\Supervisor\rp_supervisor_260821_190000_by_gpt-5_child05_p5a_requested_meanings.md`, 9,451 bytes / 86 LF / SHA-256 `33FCE155065D285C694C8BAC330E2ED5025BDB6AB59B0A54E294C0ACA7068781` |
| `E5-P5A-DETECT1` | detect-changes result before commit | recorded — `anvien detect-changes --repo E:\Anvien --scope all` exit `0`; exact changed manifest `10` files (four ledgers + six authorized source/test owners), affected app layers `backend=3,mixed=1`, functional area `providers=4`, changed count `134`; capture SHA-256 `CCBF6248C6F4A73744F76664F033F7CCDF9EC0CE07D82974B7EAF922D35FF1F0` |
| `E5-P5A-COMMIT1` | isolated P5-A commit hash | pending |

### P5-B - Export tables

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E5-P5B-IMPACT1` | fresh impact for table owner and affected consumers | pending |
| `E5-P5B-SRC1` | production table derived only from accepted Child 04 facts | pending |
| `E5-P5B-ZEROBARREL1` | zero physical declarations and a non-empty syntax-derived export surface | pending |
| `E5-P5B-BUILD1` | full repository build | pending |
| `E5-P5B-TEST1` | named/default/alias/type-only/star/namespace table cases | pending |
| `E5-P5B-COUNT1` | unchanged absolute path-resolution and `IMPORTS` counts | pending |
| `E5-P5B-REVIEW1` | Supervisor PASS for P5-B | pending |
| `E5-P5B-DETECT1` | detect-changes result before commit | pending |
| `E5-P5B-COMMIT1` | isolated P5-B commit hash | pending |

### P5-C - Re-export traversal

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E5-P5C-IMPACT1` | fresh traversal/orchestration impact and blast radius | pending |
| `E5-P5C-SRC1` | production traversal diff with cycle, ambiguity, meaning, and proof behavior | pending |
| `E5-P5C-PROOF1` | exact terminal/cycle/ambiguity proof for every named fixture | pending |
| `E5-P5C-BUILD1` | full repository build | pending |
| `E5-P5C-TEST1` | alias/star/namespace/meaning/cycle/ambiguity behavior tests | pending |
| `E5-P5C-NOGLOBAL1` | explicit import miss cannot bind a repository-global same-name Symbol | pending |
| `E5-P5C-REVIEW1` | Supervisor PASS for P5-C | pending |
| `E5-P5C-DETECT1` | detect-changes result before commit | pending |
| `E5-P5C-COMMIT1` | isolated P5-C commit hash | pending |

### P5-D - Terminal binding and target proof

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E5-P5D-IMPACT1` | fresh emission/projection/affected-reader impact | pending |
| `E5-P5D-SRC1` | production terminal binding/proof projection diff | pending |
| `E5-P5D-BUILD1` | full repository build | pending |
| `E5-P5D-TEST1` | direct/barrel equality and terminal graph behavior tests | pending |
| `E5-P5D-PARITY1` | graph and only affected persistence/readers retain equal terminal/proof fields | pending |
| `E5-P5D-TARGET1` | exact two target call sites resolve to expected terminal Symbols (`2/2`) with zero matching false gaps | pending |
| `E5-P5D-ORACLE1` | independent source/TypeScript oracle for both target sites | pending |
| `E5-P5D-BOUNDARY1` | target pre/post HEAD, worktree, source hashes, graph path, and artifact inventory | pending |
| `E5-P5D-COUNT1` | absolute before/after physical path-resolution and syntactic `IMPORTS` counts; both deltas `0` | pending |
| `E5-P5D-REVIEW1` | Supervisor PASS for P5-D | pending |
| `E5-P5D-DETECT1` | detect-changes result before commit | pending |
| `E5-P5D-COMMIT1` | isolated P5-D commit hash | pending |

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E5-PNA-REVIEW1` | independent Supervisor acceptance of all Child 05 source, target, ledger, benchmark, and commit claims | pending |
| `E5-PNB-CLEAN1` | dead-work inventory, removal result, final diff, and Supervisor confirmation | pending |
| `E5-PNC-DETECT1` | final detect-changes evidence after accepted cleanup | pending |
| `E5-PNC-COMMITS1` | ordered P5-A/P5-B/P5-C/P5-D commit hashes and worktree ownership | pending |
| `E5-PNC-HANDOFF1` | exact accepted facts/results supplied to Child 06 and its refreshed opening status | pending |
