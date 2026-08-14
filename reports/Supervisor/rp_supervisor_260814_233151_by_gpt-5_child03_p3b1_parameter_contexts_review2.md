# Supervisor Report: Child 03 P3-B1 Parameter Contexts REVIEW2

Verdict: PASS

## Metadata

- Review ID: `E3-P3B1-REVIEW2`.
- Review time: `2026-08-14 23:31:51 +07:00` (`SE Asia Standard Time`).
- Reviewer: `gpt-5`, continuing the same visible independent Supervisor identity as REVIEW1.
- Repository: `E:\Anvien`.
- Branch / HEAD: `master` / `06229bea5735e75c8d8a476c738f627bf93def8d`.
- Upstream: `origin/master` at `ea1e27710966eb54a4533255dc80a2ac0c2120d8`; the Owner-owned orchestration-skill commit in HEAD remains outside P3-B1 semantics.
- Scope: focused re-review of stable blocker `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1` only, on the exact two-file repair boundary.
- Immutable REVIEW1: `reports/Supervisor/rp_supervisor_260814_222415_by_gpt-5_child03_p3b1_parameter_contexts.md`, `19,388` bytes, SHA-256 `93A9427D84405ADD763140EEC7241FB1A795E796B41DDCCD4F786523F2E40184`.
- Repair handoff: `reports/coder/rp_coder_260814_230743_by_gpt-5_child03_p3b1_parameter_contexts_repair.md`, `15,424` bytes, SHA-256 `E2F58507B75A951628FE381D5BCAACB67AEBA2AAC9F1F3EB70A48F059D4DF186`.
- Original coder handoff: `reports/coder/rp_coder_260814_215327_by_gpt-5_child03_p3b1_parameter_contexts.md`, `16,257` bytes, SHA-256 `BF6CA0B6F68AE2ECFC23F3F1CF27F9E8EF251961BEC78AE14FE9AC3F7B73E65F`.
- Reserved and not performed: `E3-P3B1-DETECT1`, staging, commit, push, cleanup deletion, or any P3-B2/P3-B2A/P3-C/P3-C2 action.

## Executive Decision

The repair closes the sole REVIEW1 blocker. A `required_parameter` or `optional_parameter` can now reach parameter extraction only when it is a direct child of `formal_parameters`, those formal parameters are directly owned by a recognized runtime callable, and that callable's named `parameters` field resolves to the same Tree-sitter node ID. Labeled tuple members fail the direct-parent gate. Function-, constructor-, call-, and construct-type signature parameters fail the runtime-callable gate. No arbitrary callable-ancestor fallback remains.

The real-`Extract` former-failure test covers required, optional, and rest labeled tuple members in an actual parameter annotation, return type, and body-local type, plus required, optional, and rest parameters in nested function and constructor types. It proves the real runtime parameter still emits exactly once while all 15 type-only names emit zero parameter leaves, Variable definitions, owned definition occurrences, local bindings, and diagnostics. The final repair bytes independently pass the holder-clean full build and every required focused, preservation, downstream, package, and repo-native gate.

Verdict is therefore `PASS` for P3-B1 only. This authorizes Orchestration to perform the remaining planner/`E3-P3B1-DETECT1`/isolated-commit closure. It does not accept or open P3-B2, P3-B2A, P3-C, P3-C2, target validation, or any later slice.

## Rejected Invariant Reconstruction

REVIEW1 rejected only `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1`:

- the preorder collector visits every named TypeScript AST node;
- the prior candidate globally dispatched `required_parameter` and `optional_parameter`;
- pinned Tree-sitter TypeScript aliases labeled tuple members to those same node kinds;
- the prior nearest-function-ancestor lookup therefore attributed legal type-only labels and nested type-signature parameters to an enclosing runtime callable;
- those nodes could emit false `BindingLeaf`, `NodeVariable` Definition, callable `OwnedDefID`, and `BindingLocal` facts.

All other REVIEW1-cleared invariants remained locked. Because the repair changed `definitions.go` and `extract_test.go`, current source plus build/behavior evidence was refreshed once on the final bytes. The six dependent files, prior JSON multiset audit, historical TypeScript baseline, and accepted impact boundary were not reopened without invalidation.

## Source and Grammar Review

### Runtime formal-parameter ownership

Current `internal/providers/tsjs/definitions.go` establishes the following gate before extraction:

1. `formalParameterCallable` requires the candidate node's direct parent to be `formal_parameters`.
2. It requires the direct parent of that node to satisfy the same `isFunctionScopeNode` predicate used by scope construction.
3. It resolves the callable's named `parameters` field and requires its Tree-sitter node ID to equal the candidate's direct `formal_parameters` parent ID.
4. Only that proven callable is passed to `parameterCallableScopeID`; the previous ancestor search is absent.
5. `parameterCallableScopeID` derives the exact existing `ScopeFunction` ID from that callable's range and returns no scope when it is not present.

This is syntax-role and ownership validation, not a name special case.

### Type-only exclusions

Pinned grammar inspection at `github.com/tree-sitter/tree-sitter-typescript@v0.23.2/common/define-grammar.js` independently confirms:

- tuple members alias `tuple_parameter` and `optional_tuple_parameter` to `required_parameter` and `optional_parameter`, but their direct parent is `tuple_type`, so they fail gate 1;
- `function_type` and `constructor_type` own a named `parameters: formal_parameters` field, but neither is a runtime function-scope kind, so they fail gate 2;
- the same gate also excludes formal parameters owned by type-only call/construct signatures;
- no type-only node can fall through to diagnostics because scope ownership is rejected before `extractParameterBindingPattern` is invoked.

### Preserved legal paths

- Runtime function declarations/signatures, methods, constructors, function expressions, and parenthesized arrows use their own `formal_parameters` identity and the matching existing function scope.
- TypeScript unparenthesized arrows retain the separate `arrow_function` + `parameter` field path and pass the arrow callable directly.
- The explicit TypeScript `this` pseudo-parameter still exits before extraction and emits no declaration fact or diagnostic.
- Plain JavaScript unparenthesized-arrow behavior remains behind `scanner.TypeScript` and is unchanged.
- Parameter fact emission still creates one accepted parameter `BindingLeaf`, one `NodeVariable` Definition, one callable-scope `OwnedDefID`, and one `BindingLocal` per legal leaf, preserving ranges, selections, path, modifiers, provenance, and scope.
- No declared/return type is invented. The separate TypeBinding/type-inference path is unchanged.
- The repair changes no P3-A walker, ScopeIR contract, collector traversal/result path, scope builder, type/reference/import owner, Vue/resolution production source, graph owner, catch/loop/assignment behavior, or later context.

## Former-failure Oracle Review

`TestExtractParameterBindingPatternsExcludeTypeOnlyParameterSyntax` is a real `Extract` test, not a helper-only probe. Its fixture contains:

- one actual runtime parameter `arg`;
- annotation tuple labels: required, optional, rest;
- nested function-type parameters: required, optional, rest;
- nested constructor-type parameters: required, optional, rest;
- return tuple labels: required, optional, rest;
- body-local tuple labels: required, optional, rest.

The assertions independently establish:

- global extraction diagnostics are zero;
- `arg` has exactly one parameter-context leaf and exactly one Variable definition with identical range/selection;
- exactly one runtime function scope exists, so type-only callable syntax creates no runtime scope;
- `arg` has exactly one ownership occurrence and one local binding in that scope;
- each of the 15 distinct type-only names has zero leaves, zero Variable definitions, zero ownership occurrences, and zero local bindings;
- body-local variable traversal remains live and emits exactly one `bodyLocal` Variable definition.

The existing focused matrix separately retains actual function, method, constructor, parenthesized-arrow, unparenthesized-arrow, optional, explicit-`this`, JavaScript preservation, range/path/modifier/provenance, lexical ownership, and nested shadowing coverage.

## Git and Byte Boundary

- Initial and pre-report branch/HEAD remained `master` / `06229bea5735e75c8d8a476c738f627bf93def8d`.
- `origin/master` remained `ea1e27710966eb54a4533255dc80a2ac0c2120d8`.
- Staged paths before this report: `0`.
- Tracked diff before this report: exactly `12` paths — four current governance documents plus the eight inherited P3-B1 candidate/oracle paths.
- Repair delta: exactly `internal/providers/tsjs/definitions.go` and `internal/providers/tsjs/extract_test.go` relative to the immutable REVIEW1 checkpoint.
- Untracked paths before this report: exactly the immutable REVIEW1, original coder report, and repair coder report.
- `git diff --check`: exit `0` after all independent build/tests.
- Repair-named temp/debug/probe paths under `.tmp`: `0`.
- No stage, detect, commit, push, reset, checkout, stash, rebase, amend, cleanup deletion, or hidden lane occurred.

### Governance hashes

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| Roadmap | 15,347 | `660F0773B272496EB5FCE6E4AB1D2AB0C1679B280DD443C2EF659044E6EA9D29` |
| Active plan | 55,317 | `24BE498630BB906C8F7A31640DFE75FF53E94477540512376967E6A441D87819` |
| Evidence ledger | 53,198 | `B0BE2A4A34C0BEFD754A950CE4ABBE3CCDF43FEC4631CF71B93B6C7B18285E6F` |
| Actual-status ledger | 40,308 | `67DA28BEDABC9CADD3F73951CDBB9FEF9DA130579D7D7558DE6051BBD271FF12` |
| Benchmark ledger | 7,693 | `B5B01879C59E0556A2A8E04D0C54693F8EF41DBB574D64F621BDDB7A4EA77F1E` |

### Repair-owner hashes

| File | Bytes | SHA-256 |
|---|---:|---|
| `internal/providers/tsjs/definitions.go` | 13,167 | `F810A749F4FEDA8660ACF7C09B13518F78A86F1FBCC3E8471D4197E01A0E7B15` |
| `internal/providers/tsjs/extract_test.go` | 58,183 | `6D63A5AC2F3781A0088581C59DE507F8245E4566A544906FBDBBF5C6503773C6` |

### Byte-locked dependent hashes

| File | Bytes | SHA-256 |
|---|---:|---|
| `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json` | 1,788 | `AD09CD8B8F93B3A4A4D38F6939A05424BD09C89C35E76600152057BFD5884398` |
| `internal/providers/vue/extract_test.go` | 11,895 | `B2F9C261F68DD369FAFE9B4178AAD342896A1A5683325EBAD5F697D983EBC04D` |
| `internal/providers/vue/testdata/vue_scopeir_signature.golden.json` | 1,757 | `3317D30FFC8E640B12EE45C3B95B6DDEAA12E7C66D7BE477556D2978D5A6A03C` |
| `internal/resolution/graph_parity_test.go` | 4,322 | `0607074F015487D0D836C24C55A379A0CAD76292C0AF63049700CF5C8223430A` |
| `internal/resolution/testdata/typescript_graph_baseline_counts.golden.json` | 2,964 | `5DC4FDC0439CC8D6D25C5E72CE2DD5744C094DCDDB5C0B7C46CA9FCD0D304A61` |
| `internal/resolution/testdata/typescript_graph_signature.golden.json` | 8,093 | `FF9396290CB2678743B70A51ABF649B13CC0D29B17CDB54CBE42EE9B1DB3CB67` |

All supplied hashes matched before the report. No dependent byte or observed output drifted, so REVIEW1's direct JSON multiset audit and historical TypeScript baseline truth remain locked and valid. A new JSON audit was correctly not triggered.

## Anvien and Build Evidence

- `anvien --help`: exit `0`.
- Fresh `anvien analyze --force`: exit `0`; `1,646` scanned / `689` parsed code / `0` failed; graph `100,738` nodes / `140,534` relationships.
- `anvien status`: indexed/current commit both `06229be`; status up-to-date.
- Locked file/symbol impact evidence was not rerun because its bytes/evidence were not invalidated and focused REVIEW2 authority forbids reopening it merely to challenge REVIEW1.
- Pre-build `anvien doctor locks --repo E:\Anvien --json`: lock `free`, no lock file.
- Process inspection found no Go/npm/build/Anvien holder. VS Code-owned Playwright test-server PID `9352` was unrelated and preserved.
- Exclusive-open probe succeeded for all `6/6` replacement artifacts: repo CLI/DLL, launcher/server, and global CLI/DLL.
- Independent final-byte `npm run full-build`: exit `0` in `108.660s`; CLI `1.2.8`; Web/Vite `2,943` modules; final analyze `1,646/689/0`; graph `100,738/140,534`.
- npm allow-scripts and Vite dynamic-import/chunk-size output were non-failing warnings.

## Independent Behavior Gates

| Command / gate | Result | What it proves |
|---|---|---|
| `go test ./internal/providers/tsjs -run '^TestExtractParameterBindingPatternsExcludeTypeOnlyParameterSyntax$' -count=1 -v` | PASS `1/1`, `3.043s` | exact former failure: actual `arg` emits once; 15 type-only tuple/function/constructor names emit zero false facts/diagnostics |
| `go test ./internal/providers/tsjs -run '^(TestExtractParameterBindingPatterns.*|TestExtractTypeScriptScopeIRParityFixture)$' -count=1 -v` | PASS `5/5`, `2.710s` | complete parameter matrix, unparenthesized arrow, optional/`this`, JavaScript control, shadowing, TSJS parity |
| `go test ./internal/providers/vue ./internal/resolution -run '^(TestExtractVueScopeIRParityFixture|TestResolveVueGraphParityCounts|TestResolveTypeScriptGraphBaselineCountsAreReconciled|TestResolveTypeScriptGraphSignatureFixture)$' -count=1 -v` | PASS `4/4`, `4.393s` | byte-locked dependent outputs remain truthful with no Vue/resolution drift |
| `go test ./internal/providers/tsjs -run '^(TestExtractVariableBindingPattern.*|TestExtractBindingPattern.*|TestP1BPreservesBindingPatternBoundary)$' -count=1 -v` | PASS `10/10`, `2.673s` | accepted P3-A walker and P3-B variable invariants are preserved |
| `go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers/vue ./internal/resolution -count=1` | PASS `4/4` packages, `4.219s` | nearest complete TSJS/ScopeIR/Vue/resolution package boundary |
| `go test ./cmd/... ./internal/... -count=1` | PASS, `151.608s` | repository-native command/internal product regression |

No gate was retried. Build/test timings are validation evidence, not benchmark claims.

## Invariant Closure

- `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1`: CLOSED.
- Required/optional parameter dispatch is now bound to the owning runtime callable's own formal-parameter surface.
- Labeled tuple members and nested function/constructor/type-signature formal parameters emit zero parameter BindingLeaves, Variable Definitions, OwnedDefIDs, local Bindings, and diagnostics.
- Legal actual parameters continue to emit exactly once with correct callable scope.
- TypeScript unparenthesized arrow and explicit `this` behavior remain correct; plain JavaScript arrow behavior remains unchanged.
- No identifier-name special case, fake type, TypeBinding change, graph behavior change, or later-context change exists.
- Every REVIEW1-cleared invariant remains byte-locked or has fresh final-byte build/behavior proof as required.
- Residual unverified or broken same-invariant surface: none found.

## Overall Evaluation and Handoff

`E3-P3B1-REVIEW2` is `PASS`. Only Child 03 slice P3-B1 is accepted for Orchestration's planner refresh, final `E3-P3B1-DETECT1`, and isolated `E3-P3B1-COMMIT1` closure. No final detect was run by this Supervisor.

Recipient: Orchestration main in task `Điều hành toàn bộ công việc`.

Explicit non-acceptance: P3-B2, P3-B2A, P3-C, P3-C2, graph-projection implementation, real-target validation, and every later slice remain locked and unaccepted.

The report's final byte length and SHA-256 are intentionally recorded in the external handoff after creation; embedding a file's own cryptographic hash would change that hash.
