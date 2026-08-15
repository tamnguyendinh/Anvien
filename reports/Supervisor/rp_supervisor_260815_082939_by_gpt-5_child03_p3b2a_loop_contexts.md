# Supervisor Report: Child 03 P3-B2A loop contexts

Verdict: REJECT

Verdict ID: `E3-P3B2A-REVIEW1`

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260815_082939_by_gpt-5_child03_p3b2a_loop_contexts.md`
- Review time: 2026-08-15 08:29:39 +07:00 (Asia/Bangkok)
- Reviewer: `gpt-5`
- Repository: `E:\Anvien`, branch `master`
- Reviewed base: `HEAD == origin/master == 19247b4eb58a4e01a6256f3d63bbb59839644d64`
- Scope reviewed: exact P3-B2A four-file candidate in `definitions.go`, `scopes.go`, `references.go`, and `extract_test.go`; exact ten-path opening worktree boundary
- Claim reviewed: declaration-form TS/JS `for-in`/`for-of` binding facts and ownership are exact; assignment-form targets create zero declarations and emit truthful existing-contract writes without duplicate/false reads or lost legitimate reads; all prior contracts are preserved; `E3-P3B2A-SRC1`, `BUILD1`, `TEST1`, and `BOUNDARY1` are satisfied
- Authority used: Owner delegation; replacement `AGENTS.md`; working-rules and supervisor skills; active P3-B2A plan lines 348-395; evidence/actual-status/benchmark P3-B2A rows; accepted `E3-P3B2A-IMPACT1`; pinned TypeScript v0.23.2 and JavaScript v0.25.0 node-type contracts; current production/test source and production parser runtime
- Related artifacts:
  - `reports/coder/rp_coder_260815_021253_by_gpt-5_child03_p3b2a_loop_contexts.md`, 18,519 bytes, SHA-256 `5D29AFBABB96A0549705A445D1E2262154F37A442947B01F102A5E75DC673612`
  - `reports/coder/rp_coder_260815_024936_by_gpt-5_child03_p3b2a_loop_contexts_implementation.md`, 16,298 bytes, SHA-256 `F0877B7B6965D0587174E6F1EC7FF9C4BA017E1125D19974C51E026708785648`

## Executive Summary

- Problem: determine whether the exact four-file candidate closes the complete P3-B2A loop declaration/assignment invariant rather than only the direct-target fixtures.
- Decision: REJECT. Direct declaration ownership and the direct object/array/rest/default/member recursion are source-cleared, but legal assignment-target wrappers remain broken. Production passes raw `for_in_statement.left` into switches that do not handle `parenthesized_expression` or TypeScript `non_null_expression`. Fresh production-parser execution proves parenthesized identifiers emit no write, parenthesized members retain the exact false read P3-B2A must remove, and a TypeScript non-null identifier emits no write.
- Required outcome: repair only `references.go` and `extract_test.go`, preserve all locked declaration/scope bytes and semantics, then rerun the final-byte build and complete bounded validation matrix before requesting `E3-P3B2A-REVIEW2`.

## Blocking Findings

### [HIGH] `P3B2A-ASSIGNMENT-TARGET-WRAPPER-CLOSURE-1` — legal representable wrapper targets lose writes or retain a false read

File: `internal/providers/tsjs/references.go:64`

Issue: `emitLoopAssignmentWrites` forwards the raw `left` node. `emitAssignmentTargetWrites` at lines 67-92 handles direct identifiers, members, patterns, pairs, defaults, and rest, but not legal `parenthesized_expression` or `non_null_expression` wrappers. The exact-member suppression path has the same gap: `isLoopAssignmentMemberTarget` passes raw `left` at line 103, and `assignmentTargetContainsMember` at lines 108-129 has no wrapper recursion. Consequently, a wrapper around an otherwise supported identifier/member is not treated as an assignment leaf.

Evidence:

- Active plan authority requires all no-`kind` assignment targets to route to truthful `AccessWrite` facts, one truthful write per representable target, and elimination of the loop-member false read (`plan.md:370`, `:372`, `:384`).
- The accepted impact report requires the assignment matrix to cover “grammar-supported representable targets” (`rp_coder_260815_021253...md:182`).
- Pinned TS and JS node-type contracts both permit `parenthesized_expression` as `for_in_statement.left`; the TypeScript contract also permits `non_null_expression`. These are wrappers around already-supported identifier/member leaves and require no ScopeIR contract expansion.
- The repository already has the exact inspect-only unwrapping primitive: `unwrapExpression` handles `non_null_expression` and `parenthesized_expression` at `internal/providers/tsjs/nodes.go:141-155`.
- Fresh disposable probe command: `go run ./.tmp/p3b2a-supervisor-probe`. It used production `internal/parser.Pool` and production `tsjs.Extract`; the probe was located under the repository and removed completely after capture.

| Runtime case | Parser result | Current ScopeIR result | Required result |
|---|---|---|---|
| TS `for ((plain) of rows) {}` | `hasError=false`, left `parenthesized_expression` | zero `AccessFact` | one unqualified `plain/write` |
| JS `for ((plain) of rows) {}` | `hasError=false`, left `parenthesized_expression` | zero `AccessFact` | one unqualified `plain/write` |
| TS `for ((target.value) of rows) {}` | `hasError=false`, left `parenthesized_expression` | one `value/read`, receiver `target`; zero write | one `value/write`; zero `value/read` |
| JS `for ((target.value) of rows) {}` | `hasError=false`, left `parenthesized_expression` | one `value/read`, receiver `target`; zero write | one `value/write`; zero `value/read` |
| TS `for (target! of rows) {}` | `hasError=false`, left `non_null_expression` | zero `AccessFact` | one unqualified `target/write` |

- The focused oracle at `internal/providers/tsjs/extract_test.go:1100-1153` covers only direct identifier/array/object/rest/default/member targets. It therefore passes without exercising the broken legal wrappers and is narrower than the acceptance claim.

Why this blocks acceptance: the full claim says representable identifier/member assignment targets are handled and the member false read is suppressed at the exact loop assignment leaf. These legal wrappers do not change the underlying representable target, yet current runtime output contradicts both statements. The same-scope invariant is not closed, so `E3-P3B2A-SRC1`, `E3-P3B2A-TEST1`, and `E3-P3B2A-BOUNDARY1` cannot be accepted.

Fix direction: in `references.go` only, recurse through the grammar-legal assignment-target wrappers before write classification and through the same wrappers in exact-member containment/suppression. Keep recursion restricted to target positions; do not suppress computed property-key/default-initializer/nested-receiver reads and do not change general reference behavior or ScopeIR contracts. Add real-parser/real-`Extract` wrapper controls in `extract_test.go` after production is correct.

Re-review evidence required: exact final hashes for the two-file repair; TS/JS parenthesized identifier and member cases; TypeScript non-null identifier/member cases; zero declarations/scopes/leaves for every assignment form; exactly one write and zero false leaf read; preservation of computed-key/default-initializer/nested-receiver reads; the original direct target matrix; holder-clean full build and the complete original bounded regression matrix on final bytes.

## Bracket-Target Classification

Bracket targets are not the rejection basis. Both pinned grammars accept `subscript_expression`, and the probe confirmed TS/JS `for (target[index] of rows) {}` parses without error and currently emits no access fact. The existing `AccessFact` contract stores only a dot-member-like `Name` plus `ExplicitReceiver`; downstream resolution reconstructs `receiver.name` and has no computed/dynamic-access discriminator. Encoding `target[index]` as that contract would invent false dot-member semantics, while adding computed-access semantics is explicitly outside P3-B2A. This legal target is therefore not truthfully representable by the current contract. The resubmission oracle should record this out-of-contract control explicitly so it is not silently conflated with the supported wrapper/member surface.

## Source-Level Clearance Notes

- `internal/providers/tsjs/definitions.go`: source-clear and byte-locked. Lines 96-209 discriminate only `var|let|const`, map `in|of`, consume the accepted walker once, select `var` function/module ownership versus exact lexical loop scope, and emit one Definition/OwnedDefID/local Binding per returned leaf. No blocker was found in this file. Locked hash: `4936F9DD0012F787E0BF007DA8070527C945E28D0DC0B1A0BA083A49D45350D7`.
- `internal/providers/tsjs/scopes.go`: source-clear and byte-locked. Lines 33-52 add an exact loop `ScopeBlock` only for `let|const`; `var`, assignment form, and unsupported kinds add none. Locked hash: `54132D72927F721E32336C075A553DCF8947CC05180240CF0B2AF6F8A7345509`.
- `internal/providers/tsjs/references.go`: blocked at lines 64, 71-92, 103, and 112-129 by `P3B2A-ASSIGNMENT-TARGET-WRAPPER-CLOSURE-1`. Direct target-position recursion and exact direct-member suppression are otherwise appropriately narrow.
- `internal/providers/tsjs/extract_test.go`: blocked as an incomplete oracle. Lines 1100-1153 prove only unwrapped direct targets and cannot support the broader claim.

## Locked Clearances and Bytes

The repair must preserve:

- exact declaration/assignment discrimination by non-empty `kind`;
- one accepted declaration leaf to one Variable Definition, one owner occurrence, and one local binding;
- `var -> nearest function/module`; `let|const -> exact loop ScopeBlock`;
- assignment form creates zero BindingLeaves, Definitions, declaration ScopeBlocks, and diagnostics;
- target-only object/array/pair/default/rest recursion;
- no property-key/default-initializer write;
- exact direct member leaf suppression while preserving nested receiver, computed-key, and initializer reads;
- all P3-A/P3-B/P3-B1/P3-B2, TS/JS parity, import/type/call/access, ScopeIR, Vue/resolution, graph/persistence, await/using, and later-slice boundaries.

Locked files/artifacts:

- `definitions.go` and `scopes.go` at the hashes above;
- all five governance artifacts at their supplied hashes;
- both existing coder reports, including accepted `E3-P3B2A-IMPACT1`, at their supplied hashes;
- P3-A/P3-B/P3-B1/P3-B2 accepted commits and behavior.

Only this narrower repair boundary is authorized by the rejection: `internal/providers/tsjs/references.go` and `internal/providers/tsjs/extract_test.go`. No other production, test, fixture, golden, contract, governance, or report byte needs repair. Existing reports remain immutable; any resubmission handoff must be a new artifact.

## Evidence Checked

Passed:

- Opening and pre-report integrity: `HEAD == origin/master == 19247b4eb58a4e01a6256f3d63bbb59839644d64`; exact ten-path manifest; staged count `0`; all 11 supplied source/test/governance/report hashes matched; `git diff --check` exit `0`.
- Fresh Anvien analyze: `1,652` scanned, `689` parsed, `0` failed; `101,421` nodes and `141,443` relationships; index current.
- Exact Anvien evidence: the three production files have HIGH relationship risk; exact authorized existing symbols are LOW `0/0/0`; focused test owner is LOW. This is a scoped blast-radius warning.
- Full production diff and all touched production source were read before test reliance.
- Source/runtime narrow preservation: TS and JS computed member key plus member target plus member default initializer produced exactly `keys.name/read`, `target.value/write`, and `defaults.next/read`, demonstrating the intended target-position recursion is narrow for that case.
- Disposable probe residue check: removed; no `.tmp` path matching P3-B2A/loop-context/Child 03 remains.

Failed:

- Fresh production parser/`tsjs.Extract` wrapper matrix listed in the blocking finding.
- Source closure: raw wrapper nodes are absent from both assignment-write recursion and exact-member containment.
- Oracle quality: focused assignment test omits every failing wrapper form.

Not run:

- Independent `npm run full-build`.
- Focused loop `4/4`, prior-slice `14/14`, catch `4/4`, TSJS scope/parity `6/6`, selected ScopeIR tests, four-package boundary, and repo-native `go test ./cmd/... ./internal/... -count=1`.
- `detect-changes`, staging, commit, push, or later-slice work.

Reason: Owner continuity explicitly directs REJECT as soon as a legal active-plan/accepted-impact form fails and forbids unrelated build/test gates merely to delay verdict. Any repair changes final bytes, so the full build and bounded validation matrix must run fresh after repair.

## Invariant Closure

- Affected invariant: no-`kind` loop assignment targets must emit truthful existing-contract writes, suppress only the exact member target leaf, and preserve legitimate sibling reads while creating no declaration facts.
- Sibling surfaces checked: direct identifier; array/object; pair/property key; rest; default left versus initializer; direct and nested member; computed member key; parenthesized identifier/member; TypeScript non-null identifier; bracket target contract classification; declaration scope/ownership paths; focused test oracle.
- Residual same-invariant surface: confirmed live wrapper failure; acceptance is blocked. No broader domain sweep is needed for this REJECT.

## Invalidated Gates and Required Fresh Reruns

- `E3-P3B2A-IMPACT1`: remains accepted and locked.
- `E3-P3B2A-SRC1`: REJECTED by `P3B2A-ASSIGNMENT-TARGET-WRAPPER-CLOSURE-1`.
- `E3-P3B2A-BUILD1`: candidate output is not acceptance evidence and will be stale after repair; rerun once holder-clean on final repair bytes.
- `E3-P3B2A-TEST1`: invalidated by the missing wrapper oracle; rerun the full original matrix after repair.
- `E3-P3B2A-BOUNDARY1`: rejected because fresh real ScopeIR output is wrong for legal representable wrapper targets; rerun after repair.
- `E3-P3B2A-REVIEW1`: this report is the sole verdict and is REJECT.
- `E3-P3B2A-DETECT1` and `E3-P3B2A-COMMIT1`: remain locked and must not run until a later independent Supervisor PASS.

## Required Fix List For Resubmission

1. Repair target-wrapper recursion and exact-member containment only in `references.go`; preserve every locked direct-target behavior and all non-loop reference paths.
2. Add focused real-parser/real-`Extract` TS/JS wrapper controls only in `extract_test.go`, including explicit bracket out-of-contract behavior and no invented computed-access fact.
3. Reverify the two changed hashes plus all locked hashes, terminate build holders, prove the lock free, run one final-byte full build, then rerun every original behavior/preservation command and record exact exit codes/counts/warnings.
4. Create a new immutable coder resubmission report; do not edit either existing coder report or governance artifact.
5. Request a new independent `E3-P3B2A-REVIEW2`. Do not run detect, stage, commit, push, P3-C, or any later slice before PASS.

## Overall Evaluation

The declaration-side implementation and most direct assignment recursion are carefully scoped, and the candidate preserves computed-key/default-initializer reads in the exercised nested case. Acceptance nevertheless fails because source, grammar, and fresh runtime agree on a residual legal wrapper path that returns the old missing-write/false-read behavior. Green direct-target tests cannot substitute for closure of that same invariant.
