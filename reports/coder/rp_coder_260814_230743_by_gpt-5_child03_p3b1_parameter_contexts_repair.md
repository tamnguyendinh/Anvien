# Child 03 / P3-B1 Parameter Context Disambiguation Repair — Coder Handoff

## Candidate State

- State: `READY_FOR_SUPERVISOR`.
- This is a repaired coder candidate for the same P3-B1 slice. It is not an acceptance verdict.
- Handoff recipient: Orchestration main in task `Điều hành toàn bộ công việc`.
- Stable prior verdict: `E3-P3B1-REVIEW1 = REJECT`.
- Sole repaired invariant: `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1`.
- Candidate evidence refreshed here: `E3-P3B1-SRC1`, `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, and `E3-P3B1-BOUNDARY1`. The accepted owner/impact basis for `E3-P3B1-IMPACT1` was re-anchored and not redundantly rerun.
- `E3-P3B1-REVIEW2` is reserved for an independent Supervisor. `E3-P3B1-DETECT1`, `E3-P3B1-COMMIT1`, staging, commit, and push were not run.
- P3-B2, P3-B2A, P3-C, P3-C2, and all later slices remain locked.

## Authority and Re-anchor

The repair was performed under:

1. `E:\Anvien\AGENTS.md`.
2. Full `.agents/skills/working-rules/SKILL.md` and `.agents/skills/coder/SKILL.md`.
3. Current roadmap and Child 03 P3-B1, E3-P3B1, R19, and touch-map authority.
4. Immutable Supervisor report `reports/Supervisor/rp_supervisor_260814_222415_by_gpt-5_child03_p3b1_parameter_contexts.md`, 19,388 bytes, SHA-256 `93A9427D84405ADD763140EEC7241FB1A795E796B41DDCCD4F786523F2E40184`.
5. Stable repair handoff limiting production/test edits to `internal/providers/tsjs/definitions.go` and `internal/providers/tsjs/extract_test.go`.

No hidden lane or subagent was created. Accepted discovery, graph owner, file-detail, and impact gates were not restarted. The full build's repository-native analyze step still ran as part of the mandated build workflow.

## Git and Main-owned Boundary

- Repository: `E:\Anvien`.
- Branch: `master`.
- HEAD: `06229bea5735e75c8d8a476c738f627bf93def8d`.
- Index/staged paths: `0` before repair, through validation, and before report creation.
- Current worktree also contains the pre-existing main-owned ledgers, accepted P3-B1 candidate/oracle bytes, immutable prior coder report, and immutable Supervisor report. None was reverted, overwritten, staged, or claimed as a repair edit.

Main-owned hashes remained exact through final validation:

| Main-owned file | SHA-256 |
|---|---|
| roadmap | `660F0773B272496EB5FCE6E4AB1D2AB0C1679B280DD443C2EF659044E6EA9D29` |
| Child 03 plan | `24BE498630BB906C8F7A31640DFE75FF53E94477540512376967E6A441D87819` |
| Child 03 evidence | `B0BE2A4A34C0BEFD754A950CE4ABBE3CCDF43FEC4631CF71B93B6C7B18285E6F` |
| Child 03 actual-status | `67DA28BEDABC9CADD3F73951CDBB9FEF9DA130579D7D7558DE6051BBD271FF12` |
| Child 03 benchmark | `B5B01879C59E0556A2A8E04D0C54693F8EF41DBB574D64F621BDDB7A4EA77F1E` |

## Rejected Invariant and Source Reasoning

Supervisor proved that Tree-sitter TypeScript aliases labeled tuple members to `required_parameter` / `optional_parameter`, the same kinds used by runtime formal parameters. The accepted preorder walker visits every named node. Therefore the rejected implementation's global kind dispatch plus nearest-function-ancestor search allowed tuple labels and nested function/constructor type-signature parameters to emit false runtime parameter facts.

The repair in `internal/providers/tsjs/definitions.go` establishes an exact ownership proof before extraction:

1. The candidate node's direct parent must be `formal_parameters`.
2. The direct parent of those formal parameters must be a recognized runtime function-scope node.
3. The callable's named `parameters` field must be the same Tree-sitter node, verified by node ID equality.
4. Only that proven callable is passed to scope lookup; arbitrary ancestor climbing was removed.

Consequences:

- A tuple label fails at step 1 because its parent is `tuple_type`.
- A nested function-type or constructor-type parameter passes the `formal_parameters` shape but fails at step 2 because `function_type` and `constructor_type` are not runtime function-scope nodes.
- A real function, method, constructor, parenthesized arrow, or function-expression formal parameter passes all three checks and emits into that callable's existing function scope.
- The TypeScript-only unparenthesized-arrow path still passes its arrow callable directly.
- The explicit `this` pseudo-parameter early return remains unchanged.
- Plain JavaScript arrow behavior remains guarded by `scanner.TypeScript`.
- No identifier-name special case, fake type, type-binding-path change, walker change, scope-construction change, or graph projection change was introduced.

## Exact Two-file Repair

Only these two existing files changed during the repair:

| Repair owner | Pre-repair SHA-256 | Final SHA-256 | Repair delta |
|---|---|---|---|
| `internal/providers/tsjs/definitions.go` | `416C997DF55D885A5F92353441D6AB943927C2DC2C49B2060983ACE21D920DC3` | `F810A749F4FEDA8660ACF7C09B13518F78A86F1FBCC3E8471D4197E01A0E7B15` | added direct formal-parameter ownership proof; passed the proven callable into emission/scope lookup; removed nearest-ancestor acceptance |
| `internal/providers/tsjs/extract_test.go` | `F096D08EC1482666A58E79E52DE4D21DC9E6397CF939E36AD2CC201E9B34A158` | `6D63A5AC2F3781A0088581C59DE507F8245E4566A544906FBDBBF5C6503773C6` | added one real-`Extract` former-failure test and its zero-fact assertions |

The current Git numstat relative to HEAD is `definitions.go +138/-0` and `extract_test.go +525/-5`; those totals include the earlier uncommitted P3-B1 candidate. The table above and the immutable prior-candidate hashes distinguish the exact repair checkpoint from the whole slice diff.

No third production/test owner was needed.

## Former-failure Oracle

`TestExtractParameterBindingPatternsExcludeTypeOnlyParameterSyntax` uses one real runtime parameter, `arg`, whose annotation contains:

- required, optional, and rest labeled tuple members;
- a nested function type with required, optional, and rest parameters;
- a nested constructor type with required, optional, and rest parameters.

The same runtime callable also contains required/optional/rest tuple labels in its return type and in a body-local variable type.

The test asserts:

- `arg` emits exactly one parameter `BindingLeaf`, one `NodeVariable` Definition, one `OwnedDefID`, and one `BindingLocal` in the sole runtime function scope;
- all 15 distinct type-only names emit zero BindingLeaves, zero Variable Definitions, zero OwnedDefIDs, and zero local Bindings;
- extraction diagnostics are globally zero;
- the body-local declaration is still traversed and emitted;
- type-only callable syntax does not create a runtime function scope.

No assertion depends on special identifier spellings in production code. With the rejected nearest-function-ancestor implementation, the legal type-only nodes would have reached parameter emission and failed these zero-fact assertions.

## Holder Gate and Full Build — `E3-P3B1-BUILD1`

Before both repair builds:

- `anvien doctor locks --repo . --json` reported `status: free`; `.anvien/analyze.lock` did not exist.
- `anvien doctor processes --json` showed no Go/npm/build holder.
- The only persistent relevant-looking process was VS Code-owned Playwright test-server PID `9352`; it was unrelated to this non-UI build and was preserved.
- No process termination was necessary during this repair.

Build history:

1. `npm run full-build` initially passed in `122.651s`, with `1,645` files scanned, `689` parsed code files, `0` failures, and graph `100,725` nodes / `140,521` relationships. This build was superseded after the first test invocation exposed a test-helper assertion issue and the test bytes changed.
2. After the test-only correction and a fresh clean holder/lock gate, final-byte `npm run full-build` passed in `107.415s`, exit `0`.
   - Runtime/package build: PASS; CLI version `1.2.8`.
   - Web TypeScript/Vite build: PASS; `2,943` modules transformed.
   - Analyze: `1,645` scanned, `689` parsed code, `0` failed.
   - Final graph: `100,724` nodes / `140,520` relationships.
   - npm allow-scripts and Vite dynamic-import/chunk-size messages were non-failing warnings.

Build durations are validation evidence, not product-performance benchmarks.

## Final Validation — `E3-P3B1-TEST1` / `E3-P3B1-BOUNDARY1`

All commands below ran after the final-byte full build:

| Command | Exact result | What it proves |
|---|---|---|
| `go test ./internal/providers/tsjs -run '^TestExtractParameterBindingPatternsExcludeTypeOnlyParameterSyntax$' -count=1 -v` | exit `0`, `3.310s`, 1/1 PASS | former failure: real `arg` emits once; all tuple/function-type/constructor-type labels emit zero facts and diagnostics |
| `go test ./internal/providers/tsjs -run '^(TestExtractParameterBindingPatterns.*\|TestExtractTypeScriptScopeIRParityFixture)$' -count=1 -v` | exit `0`, `2.806s`, 5/5 PASS | full parameter matrix, optional/`this`, TypeScript/JavaScript arrows, shadowing/siblings, and TSJS parity |
| `go test ./internal/providers/vue ./internal/resolution -run '^(TestExtractVueScopeIRParityFixture\|TestResolveVueGraphParityCounts\|TestResolveTypeScriptGraphBaselineCountsAreReconciled\|TestResolveTypeScriptGraphSignatureFixture)$' -count=1 -v` | exit `0`, `4.216s`, 4/4 PASS | all six byte-locked downstream oracle bytes remain truthful; no Vue/resolution drift |
| `go test ./internal/providers/tsjs -run '^(TestExtractVariableBindingPattern.*\|TestExtractBindingPattern.*\|TestP1BPreservesBindingPatternBoundary)$' -count=1 -v` | exit `0`, `2.749s`, 10/10 PASS | accepted P3-A walker and P3-B variable behavior remain preserved |
| `go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers/vue ./internal/resolution -count=1` | exit `0`, `4.447s`, 4/4 packages PASS | nearest full package boundary |
| `go test ./cmd/... ./internal/... -count=1` | exit `0`, `159.422s` | repository-native Go command/internal regression |

Because the exact downstream tests passed and all six oracle hashes remained byte-identical, no legitimate output delta existed. Per the repair authority, the direct JSON multiset audit was not triggered.

## Failures, Retries, and Superseded Evidence

Repair-local retained history:

1. The first final-looking full build passed in `122.651s`; it became superseded when test bytes changed.
2. The first new-test invocation exited `1` in `3.470s`. Production output already showed the repaired behavior (`arg` owned/bound once and type-only syntax remaining only on the preserve-only TypeBinding path), but the test reused a one-line source-range prefix helper against a multiline fixture and could not locate the function scope. Only `extract_test.go` changed: the test now identifies the sole real `ScopeFunction` directly and does not assert the unrelated pre-existing scope owner of `bodyLocal`. Production did not change.
3. The holder gate and full build were rerun on the corrected final test bytes; all final evidence is from the later `107.415s` build and commands listed above.
4. The first post-report read-only closeout wrapper had a PowerShell parser error (`An empty pipe element is not allowed`) in its JSON-formatting branch. It did not edit files or run product code. The corrected wrapper then completed: all 13 locked/governance hashes matched, tracked `git diff --check` exited `0`, and staged paths remained `0`.

Earlier P3-B1 history remains durably recorded in `reports/coder/rp_coder_260814_215327_by_gpt-5_child03_p3b1_parameter_contexts.md` (SHA-256 `BF6CA0B6F68AE2ECFC23F3F1CF27F9E8EF251961BEC78AE14FE9AC3F7B73E65F`). In particular, its first repo-native regression exited `1` in `209.3s` only on four stale Vue/resolution oracles; Orchestration later authorized and reconciled the exact downstream test artifacts. All builds in that report, including its final `120.6s` build, are superseded for the repaired source/test bytes. The immutable Supervisor report then rejected only the syntax-role invariant repaired here.

## Locked-file Verification

All six downstream files stayed byte-identical to the repair opening boundary:

| Byte-locked file | Final SHA-256 |
|---|---|
| `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json` | `AD09CD8B8F93B3A4A4D38F6939A05424BD09C89C35E76600152057BFD5884398` |
| `internal/providers/vue/extract_test.go` | `B2F9C261F68DD369FAFE9B4178AAD342896A1A5683325EBAD5F697D983EBC04D` |
| `internal/providers/vue/testdata/vue_scopeir_signature.golden.json` | `3317D30FFC8E640B12EE45C3B95B6DDEAA12E7C66D7BE477556D2978D5A6A03C` |
| `internal/resolution/graph_parity_test.go` | `0607074F015487D0D836C24C55A379A0CAD76292C0AF63049700CF5C8223430A` |
| `internal/resolution/testdata/typescript_graph_baseline_counts.golden.json` | `5DC4FDC0439CC8D6D25C5E72CE2DD5744C094DCDDB5C0B7C46CA9FCD0D304A61` |
| `internal/resolution/testdata/typescript_graph_signature.golden.json` | `FF9396290CB2678743B70A51ABF649B13CC0D29B17CDB54CBE42EE9B1DB3CB67` |

The prior coder report and Supervisor report were not modified.

## Cleanup and Exact Manifest

- `gofmt` ran only on the two authorized Go owners.
- `git diff --check`: exit `0` before report creation and again during post-report closeout.
- Post-report boundary verification checked 13 main/oracle/immutable-report hashes with `AllMatch: true`; staged paths remained `0`.
- No repair-local temp/debug/probe file was created. Full build reused repo-local `.tmp\ladybug-native` dependency state.
- No browser or Playwright action was run for this non-UI slice.
- No reset, checkout, stash, rebase, amend, final detect, stage, commit, or push was performed.

Exact current P3-B1 coder lane manifest:

1. `internal/providers/tsjs/definitions.go` — repaired production owner.
2. `internal/providers/tsjs/extract_test.go` — repaired focused test owner.
3. `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json` — inherited byte-locked P3-B1 oracle.
4. `internal/providers/vue/extract_test.go` — inherited byte-locked dependent test.
5. `internal/providers/vue/testdata/vue_scopeir_signature.golden.json` — inherited byte-locked dependent oracle.
6. `internal/resolution/graph_parity_test.go` — inherited byte-locked dependent test.
7. `internal/resolution/testdata/typescript_graph_baseline_counts.golden.json` — inherited byte-locked dependent oracle.
8. `internal/resolution/testdata/typescript_graph_signature.golden.json` — inherited byte-locked dependent oracle.
9. `reports/coder/rp_coder_260814_215327_by_gpt-5_child03_p3b1_parameter_contexts.md` — immutable original candidate report.
10. `reports/coder/rp_coder_260814_230743_by_gpt-5_child03_p3b1_parameter_contexts_repair.md` — this new repair report.

The five main-owned plan/benchmark files and the immutable Supervisor report are boundary/governance files, not coder manifest items.

## Handoff

- Coder-side blocker: none known within the exact repair boundary.
- Repaired invariant: runtime formal-parameter ownership is now proven syntactically before parameter fact emission.
- Preserved controls: explicit `this`, TypeScript unparenthesized arrow, plain JavaScript arrow, type-binding path, variables, shadowing, sibling contexts, TSJS/Vue/resolution oracles, and repo-native regression all pass.
- Required next action: Orchestration main may open `E3-P3B1-REVIEW2` with an independent Supervisor against this exact worktree and report.
- The coder makes no acceptance claim and grants no authority for final detect, commit, push, P3-B2, or P3-C.
