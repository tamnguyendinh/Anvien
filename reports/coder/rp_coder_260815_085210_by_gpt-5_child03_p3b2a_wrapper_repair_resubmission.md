# Child 03 P3-B2A wrapper-repair resubmission report

Date: 2026-08-15  
Role: visible Owner-controlled Coder lane  
State: `READY_FOR_SUPERVISOR_REVIEW2`  
Slice: Child 03 / P3-B2A / `P3B2A-ASSIGNMENT-TARGET-WRAPPER-CLOSURE-1` only  
HEAD and `origin/master`: `19247b4eb58a4e01a6256f3d63bbb59839644d64`

This is a resubmission candidate, not an acceptance claim. Independent `E3-P3B2A-REVIEW2` remains required.

## Authority and rejected invariant

The immutable first review is `reports/Supervisor/rp_supervisor_260815_082939_by_gpt-5_child03_p3b2a_loop_contexts.md` (`15,125` bytes; SHA-256 `6D9CD9B381D851762A65A42B0DE2170DB5F242A8F1555F54E91C967DF494A628`). It rejected only `P3B2A-ASSIGNMENT-TARGET-WRAPPER-CLOSURE-1`.

The rejected runtime path was narrow and reproducible: legal `parenthesized_expression` wrappers in TS/JS and TypeScript `non_null_expression` wrappers around representable identifier/member assignment targets bypassed the existing write walker and exact-member suppression. Parenthesized identifiers lost their write; parenthesized members retained a false read; TypeScript non-null identifiers lost their write. `subscript_expression`/bracket targets are explicitly out of the current dot-member `AccessFact` contract and are not converted or invented.

The updated P3-B2A repair checkpoint authorizes exactly `references.go` and `extract_test.go`. `definitions.go`, `scopes.go`, `nodes.go`, all other production/test/fixture/golden/contract files, governance, and all existing reports are locked. P3-A/P3-B/P3-B1/P3-B2 behavior, graph/resolution/persistence, types/imports/calls, `await`/`using`, control flow, P3-C+, target work, detect, staging, commit, and push remain out of scope.

## Invariant Family Map

- Family: no-kind TS/JS `for-in`/`for-of` assignment-target extraction and exact member-read suppression.
- Authority: active P3-B2A repair checkpoint, immutable REVIEW1 rejection, accepted P3-A BindingLeaf/ScopeIR contracts, current parser source, and the existing direct-target implementation.
- Primary runtime surface: `internal/providers/tsjs/references.go` (`emitReferenceKind`, final line 13; wrapper helper at line 96; target recursion at line 67; member containment at line 125).
- Oracle surface: `internal/providers/tsjs/extract_test.go`, `TestExtractLoopAssignmentTargetWrappersAndBracketControl` at final line 1155, plus the locked original loop tests.
- Sibling surfaces checked: direct identifier, array/object/pair/rest/default recursion; direct/nested member targets; computed-key/default-initializer reads; TypeScript and JavaScript parser parity; bracket target contract classification; declaration/scope ownership; prior P3 preservation tests.
- Forbidden fallback: no target-name special cases, no computed-access or new ScopeIR contract, no property-key/default-initializer/nested-receiver write promotion, no broad `as`/general reference rewrite.
- Stale artifact handling: the first implementation report and REVIEW1 report remain immutable; the disposable probe was removed; this is a new resubmission report.
- Verify matrix: real parser probe, focused original loop suite plus wrapper/bracket controls, P3-A/B/B1 preservation, P3-B2 catch preservation, TSJS scope/parity, ScopeIR tree/normalization, four-package boundary, repo-native command/internal regression, final hashes/manifest/lock cleanup.

## Production repair (`references.go` only)

The repair is surgical:

1. `emitAssignmentTargetWrites` now passes its target through `unwrapLoopAssignmentTarget` before the existing switch.
2. `unwrapLoopAssignmentTarget` recurses only through `parenthesized_expression` and `non_null_expression`, following the first named child. It does not add a bracket/subscript case and does not unwrap unrelated expression forms.
3. `assignmentTargetContainsMember` uses the same wrapper recursion before its existing target-position walk, so exact suppression recognizes a member below either legal wrapper.
4. Existing pair/value, default/left, rest, array/object, direct identifier, and direct member behavior is unchanged.
5. Because containment stops at a member leaf, a nested receiver member remains a read. Because pair recursion follows `value` and default recursion follows `left`, property keys and initializers remain outside write suppression.

Final production repair facts from a disposable real-parser/real-`Extract` probe (created under `.tmp`, then deleted before handoff):

| Runtime input | Leaves | Definitions | ScopeBlocks | Access facts |
|---|---:|---:|---:|---|
| TS `for ((plain) of rows) {}` | 0 | 0 | 0 | exactly `plain/write` |
| JS `for ((plain) of rows) {}` | 0 | 0 | 0 | exactly `plain/write` |
| TS `for ((target.value) of rows) {}` | 0 | 0 | 0 | exactly `value/write`, receiver `target`; no `value/read` |
| JS `for ((target.value) of rows) {}` | 0 | 0 | 0 | exactly `value/write`, receiver `target`; no `value/read` |
| TS `for (target! of rows) {}` | 0 | 0 | 0 | exactly `target/write` |
| TS `for (target.value! of rows) {}` | 0 | 0 | 0 | exactly `value/write`, receiver `target`; no `value/read` |
| TS `for ((target.nested.value) of rows) {}` | 0 | 0 | 0 | `nested/read` receiver `target` plus `value/write` receiver `target.nested` |
| TS/JS `for (target[index] of rows) {}` | 0 | 0 | 0 | zero AccessFacts; no invented dot-member fact |

Probe command: `go run ./.tmp/p3b2a-wrapper-repair-probe.go`. It exited `0`; `.tmp/p3b2a-wrapper-repair-probe.go` was removed immediately afterward and final residue search found no matching path.

## Focused test repair (`extract_test.go` only)

`TestExtractLoopAssignmentTargetWrappersAndBracketControl` adds nine real-parser subcases:

- TS and JS parenthesized identifiers;
- TS and JS parenthesized members;
- TS non-null identifier and member;
- TS nested receiver preservation;
- TS and JS explicit bracket/subscript out-of-contract controls.

Every subcase asserts zero BindingLeaves, zero Definitions, zero diagnostics, zero loop `ScopeBlock`s, exact access count, exact access kind/name/range/receiver, no duplicate fact, and zero `value/read`. The existing direct loop matrix remains unchanged and is rerun below.

## E3-P3B2A-SRC1 / BOUNDARY1 result

The production probe and focused oracle now close the rejected wrapper surface:

- parenthesized identifier: one unqualified write in both languages;
- parenthesized member: one property write and no false property read in both languages;
- TypeScript non-null identifier/member: one underlying write and no false read;
- nested member receiver: legitimate receiver read retained while only the outer member is written;
- bracket target: zero invented access/declaration facts, with no contract expansion;
- all wrapper forms: zero declaration facts, zero BindingLeaves, zero loop scopes, and zero diagnostics.

The locked direct-target recursion and all declaration/scope semantics remain untouched.

## E3-P3B2A-BUILD1

Before build:

- `anvien doctor locks --repo E:\Anvien --json`: exact `E:\Anvien\.anvien\analyze.lock` absent; `status=free`; `alive=false`, `stale=false`, `recoverable=false`.
- Repository-bound build-holder inventory: `0`. The editor-owned Playwright test server was not a build holder and was left untouched.

Final-byte command:

`npm run full-build`

Result: PASS, exit `0`. The packaged Go runtime, npm package, launcher, web TypeScript/Vite bundle, CLI version (`1.2.8`), and authoritative `anvien analyze . --force` all completed. Analyze reported `scanned=1653`, `parsed_code=689`, `failed=0`; graph `101,476` nodes / `141,508` relationships; file projection `1,653` files / `17,525` dependency edges / `427` unresolved.

Non-blocking warnings were the existing Vite dynamic/static import and large-chunk warnings plus npm's global install `allowScripts` warning. The command exit was `0`.

## E3-P3B2A-TEST1

All commands below ran after the final repair build and no source bytes changed afterward.

### Focused original loop matrix plus wrapper oracle

`go test ./internal/providers/tsjs -run '^TestExtractLoop' -count=1 -v`

PASS, exit `0`: original declaration/assignment/scope/parity tests plus `TestExtractLoopAssignmentTargetWrappersAndBracketControl` passed (`5/5` top-level tests; `9/9` wrapper subcases).

### P3-A/P3-B/P3-B1 preservation

`go test ./internal/providers/tsjs -run '^(TestExtractVariableBindingPattern.*|TestExtractParameterBindingPattern.*|TestExtractBindingPattern.*|TestP1BPreservesBindingPatternBoundary)$' -count=1 -v`

PASS, exit `0`, `14/14`.

### P3-B2 catch preservation

`go test ./internal/providers/tsjs -run '^(TestExtractCatchBindingPatternsEmitScopeIRFacts|TestExtractCatchBindingPatternsOptionalAndJavaScriptControls|TestExtractCatchBindingPatternsPreserveShadowingAndSiblingContexts|TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts)$' -count=1 -v`

PASS, exit `0`, `4/4`.

### TSJS scope/parity and ScopeIR tree/normalization

`go test ./internal/providers/tsjs -run '^(TestScopeIDParity.*|TestExtractTypeScriptScopeIR|TestExtractTypeScriptScopeIRParityFixture|TestExtractJavaScriptScopeIR)$' -count=1 -v`

PASS, exit `0`, `6/6` selected TSJS scope/parity tests.

`go test ./internal/scopeir -run '^(TestBuildScopeTreeLegacyParity|TestLegacyP7ScopeExtractorConversionCoversScopeTreeInvariants|TestUnmarshalNormalizesScopeIR)$' -count=1 -v`

PASS, exit `0`, every selected ScopeIR test passed.

### Four-package boundary

`go test ./internal/providers/tsjs ./internal/scopeir ./internal/providers/vue ./internal/resolution -count=1`

PASS, exit `0`; all four packages passed.

### Repository-native regression

`go test ./cmd/... ./internal/... -count=1`

PASS, exit `0`; all command/internal packages passed or reported no test files. This includes TSJS, ScopeIR, Vue, and resolution. The command completed with no failure.

## Superseded history and failure disclosure

- REVIEW1's direct-target candidate remains immutable and rejected only for the wrapper invariant; no broader defect was found or reopened.
- The wrapper probe was run before the test edit to prove production behavior independently; it passed after the production-only patch.
- No final repair test/build attempt failed. A pre-validation source review caught a test-only field name that would have been invalid Go (`range`); it was corrected before formatting/build, and no failed command or extra artifact resulted.
- The prior implementation report's earlier superseded direct-matrix attempt remains immutable and is not reused as final evidence.

## Final hashes and locked bytes

Final repair files:

| File | Bytes | SHA-256 |
|---|---:|---|
| `internal/providers/tsjs/references.go` | 7,021 | `9F7C61DA0CD8B2F9EDFE0D1300740E6EB45149BC0D134132520E1CBD6961B019` |
| `internal/providers/tsjs/extract_test.go` | 96,540 | `9BB57405BC90A830B354D624A93F7B7B24D57071EEFFC864A803649A1E193293` |

Locked production files remain exact:

| File | Bytes | SHA-256 |
|---|---:|---|
| `internal/providers/tsjs/definitions.go` | 17,871 | `4936F9DD0012F787E0BF007DA8070527C945E28D0DC0B1A0BA083A49D45350D7` |
| `internal/providers/tsjs/scopes.go` | 4,229 | `54132D72927F721E32336C075A553DCF8947CC05180240CF0B2AF6F8A7345509` |

Protected governance bytes remain exact:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| roadmap | 15,659 | `22BF726A2A2010E43046A4B66C5551E9D630F8806A4614299AA05834F3E1F1F1` |
| Child 03 plan | 69,549 | `6E3BF2BE978A2E87640BCDFC782E18F654BE9D63679A3358AA2B06776AEB4382` |
| Child 03 evidence | 69,328 | `8A1856BB2F69E532992AF139676B622D1AC4262457F59355752237368F4CF58A` |
| Child 03 actual status | 49,358 | `84B60688430C315932BE5BD10C5A1600136A51D9EF7399376F08804B6351D3E1` |
| Child 03 benchmark | 8,207 | `DAAEBB824082DF5B4F1634916631FBBEE9B03846AD5CCB6054827E721870D1A2` |

Immutable reports remain exact:

- impact report: `18,519` / `5D29AFBABB96A0549705A445D1E2262154F37A442947B01F102A5E75DC673612`;
- first implementation report: `16,298` / `F0877B7B6965D0587174E6F1EC7FF9C4BA017E1125D19974C51E026708785648`;
- REVIEW1 report: `15,125` / `6D9CD9B381D851762A65A42B0DE2170DB5F242A8F1555F54E91C967DF494A628`.

## Cleanup and Git boundary

- `git diff --check`: PASS, exit `0`.
- Staged count: `0`.
- Final lock status: free; final build-holder inventory: `0`.
- Probe path `.tmp/p3b2a-wrapper-repair-probe.go`: absent.
- `.tmp` search for P3B2A/wrapper-repair/loop-context/Child03 task residue: `0` matches.
- No governance or existing report was edited.
- No `detect-changes`, stage, commit, push, Supervisor action, P3-C+, target work, or contract expansion was performed.

Final expected manifest is exactly twelve paths: four main-owned governance files, the four candidate code/test files, the immutable impact report, the immutable first implementation report, the immutable REVIEW1 report, and this new resubmission report. HEAD and origin remain equal at `19247b4eb58a4e01a6256f3d63bbb59839644d64`.

This report's own final byte size and SHA-256 are computed after file creation and supplied in the visible handoff; embedding a whole-file self-hash would change the hash.

## Handoff

State: `READY_FOR_SUPERVISOR_REVIEW2`.

Exact next Orchestration action: open one independent visible Supervisor task for `E3-P3B2A-REVIEW2` against the final two-file hashes and this report. Do not run detect-changes, stage, commit, push, accept P3-B2A, or open P3-C until the independent review returns `PASS`.
