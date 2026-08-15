# Child 03 P3-B2A Loop Contexts — Discovery and Impact Handoff

Date: 2026-08-15  
Scope: `child03_p3b2a_loop_contexts`  
Evidence ID: `E3-P3B2A-IMPACT1` candidate  
Status: `BLOCKED_ON_ORCHESTRATION_PLAN_UPDATE`  
Role boundary: discovery/impact only; no production, test, fixture, golden, plan, ledger, roadmap, build, behavior-test, detect, stage, commit, push, or Supervisor action

## Executive determination

P3-B2A is not a test/evidence-only slice on current source.

The accepted P3-B `variable_declarator` adapter is never reached by `for-in` or `for-of` controls. Both pinned grammars and the real parser place the loop binding directly in `for_in_statement.left`, with no `variable_declarator` wrapper. Current production has no `for_in_statement` definition dispatch, so declaration-form identifier and pattern controls emit no loop declaration/binding facts.

The assignment-form non-declaration invariant is already correct: forms without a loop `kind` create zero declarations. Their write/reference behavior is not correct or preserve-only: plain identifiers and destructured identifier leaves emit no write fact, while a member target is currently classified as a read. The current plan keeps scope construction and assignment/reference extraction inspect-only, so implementation cannot begin until Orchestration changes that exact boundary.

## Authority and opening boundary

- Branch: `master`.
- Entry `HEAD`: `19247b4eb58a4e01a6256f3d63bbb59839644d64`.
- Entry `origin/master`: `19247b4eb58a4e01a6256f3d63bbb59839644d64`.
- Ahead/behind: `+0/-0`.
- Staged paths at entry: `0`.
- Untracked paths before this task's report: `0`.
- P3-B2 is accepted, committed, and pushed at the entry commit.
- P3-B2A is the sole open slice. P3-C, P3-C2, Pn, Child 04+, graph projection/resolution, target work, export/module/ambient/scanner work, and unrelated refactors remain locked.
- Fresh `anvien analyze --force` completed once, exit `0`: `1,650` scanned, `689` parsed code, `0` failed, `101,001` nodes, `140,908` relationships, graph `E:\Anvien\.anvien\graph.json`, indexed/current commit equal to entry `HEAD`.

Four main-owned governance files were byte-verified before discovery and remained preserve-only:

| File | Bytes | SHA-256 |
|---|---:|---|
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md` | 15,519 | `A5B8AA9E93FFC0F38D5F86207DFF937252849D008E923F57796E60A5A4140EBD` |
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md` | 63,357 | `B6AE2F44A1E4232157F73E6EF28C01D43B1BF70612DADD4A20BA0F8C60725E3C` |
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md` | 65,559 | `0E8FC63858CE0463CD6005C3EF7667B0C9A01AEB928D8CC74D4B873335C72F55` |
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md` | 46,311 | `803F7FF8D74EA6E415627928A090DAC1B8BC771E49C35190A407E9391C8EDCFA` |

The clean/preserve-only benchmark ledger was `8,207` bytes, SHA-256 `DAAEBB824082DF5B4F1634916631FBBEE9B03846AD5CCB6054827E721870D1A2`.

## Grammar and real-parser evidence

Pinned authority:

- TypeScript: `github.com/tree-sitter/tree-sitter-typescript v0.23.2`, `typescript/src/node-types.json`.
- JavaScript: `github.com/tree-sitter/tree-sitter-javascript v0.25.0`, `src/node-types.json`.
- Parser registry: `internal/parser/registry.go` selects `LanguageTypescript()` for `.ts` and JavaScript `Language()` for `.js`.

Both node-type contracts use one named node, `for_in_statement`, for both operators. Its required fields are `left`, `operator`, `right`, and `body`; `operator` is `in|of`; optional `kind` distinguishes a declaration form. `left` directly accepts identifier, array pattern, object pattern, and assignment targets. Neither grammar inserts `variable_declarator` below loop `left`.

A repo-local temporary Go parser probe used the production `internal/parser.Pool`, printed root S-expressions plus `for_in_statement` fields, and was removed before handoff. All ten language/case parses had `hasError=false`.

| Source form | TS v0.23.2 actual fields | JS v0.25.0 actual fields | Classification |
|---|---|---|---|
| `for (const [a, ...rest] of xs)` | node `for_in_statement`; `kind=const`; `operator=of`; `left=array_pattern` | same | declaration; accepted walker required with `BindingContextForOf` |
| `for (let {key: local = fallback} in obj)` | `kind=let`; `operator=in`; `left=object_pattern`; pair value is `assignment_pattern` | same | declaration; `local` is leaf, `.key` path, default modifier true |
| `for (var v in obj)` | `kind=var`; direct `left=identifier` | same | declaration; function/module ownership required |
| `for (let l of xs)` | `kind=let`; direct `left=identifier` | same | lexical loop declaration |
| `for (const c of xs)` | `kind=const`; direct `left=identifier` | same | lexical loop declaration |
| `for ([a] of xs)` | empty `kind`; `left=array_pattern` | same | assignment; zero declaration, write missing |
| `for ({x} in obj)` | empty `kind`; `left=object_pattern` | same | assignment; zero declaration, write missing |
| `for (plain of xs)` | empty `kind`; `left=identifier` | same | assignment; zero declaration, write missing |
| `for (target.value of xs)` | empty `kind`; `left=member_expression` | same | assignment; currently emitted as read, not write |
| module/function same-name shadowing | loop shape identical; TypeScript wraps function parameter in `required_parameter` | JavaScript formal parameter is direct `identifier` | loop shape parity; parameter wrapper is an unrelated grammar difference |

JavaScript's node-type contract additionally permits `await` and `using` tokens in its multi-valued `kind` field; TypeScript's pinned node type lists only `var|let|const`. No `await`/`using` semantic expansion is authorized by P3-B2A discovery.

## Current production truth

### Traversal and declaration routing

`internal/providers/tsjs/extract.go` collects and builds scopes before one complete named-node walk. For every node the walk invokes definition, import, type, and reference dispatch in that order.

`internal/providers/tsjs/definitions.go:14` owns `collector.emitDefinitionKind`. It handles variables only at `variable_declarator` (`definitions.go:76`) and has no `for_in_statement` branch. Therefore:

- declaration-form loop identifier controls currently emit `0` loop Definitions, `0` loop local Bindings, and `0` loop BindingLeaves;
- declaration-form loop patterns do not reach `emitVariableBindingPattern` or the accepted P3-B adapter;
- assignment forms also emit `0` declarations, which is correct for that invariant.

The existing test `TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts` explicitly treats any `BindingContextForIn`/`BindingContextForOf` leaf as a locked-sibling failure. That exact assertion must transition after production code is correct.

### Scope ownership

`internal/providers/tsjs/scopes.go:33` owns `collector.collectScopeCandidateForKind`. Current explicit scopes are module, class/interface, function, and catch `ScopeBlock`; loops are absent.

Required loop ownership is distinct by declaration kind:

- `let`/`const`: own each loop declaration in a loop lexical `ScopeBlock` whose range includes the loop body, so nested same-name declarations stay distinct and body facts are inside the loop environment;
- `var`: own the declaration in the nearest function or module scope, not the loop block;
- assignment form: create no declaration scope.

Because current `variableBindingScopeID` selects the smallest containing scope, implementation must use a loop-specific owner decision for `var` rather than silently routing it through a new lexical loop block. Existing variable, parameter, and catch scope selection must remain unchanged.

### Assignment writes/references

`internal/providers/tsjs/references.go:13` owns `collector.emitReferenceKind`. It emits calls, member accesses, constructor calls, and heritage facts. It does not emit unqualified identifier reference facts and has no loop-target dispatch.

| Assignment target | Declaration facts now | Access/write facts now | Truth classification |
|---|---:|---:|---|
| plain identifier `plain` | 0 | 0 | non-declaration correct; write missing |
| `[a]` | 0 | 0 for `a` | non-declaration correct; write missing |
| `{x}` | 0 | 0 for `x` | non-declaration correct; write missing |
| member `target.value` | 0 | one member AccessFact classified `read` | wrong access kind |
| subscript target allowed by grammar | 0 | no dedicated subscript access path | missing; exact support must follow existing representable AccessFact contract only |

This slice can emit assignment writes through existing `scopeir.AccessFact{Kind: AccessWrite}` without changing ScopeIR contracts. It must not invent a new general reference type, graph edge, exception/control-flow model, or P3-C resolution behavior.

## Invariant Family Map

- Family: TS/JS `for-in`/`for-of` binding-context discrimination.
- Authority: active Child 03 P3-B2A plan, accepted P3-A BindingLeaf contract, accepted P3-B variable semantics, pinned TS/JS grammars, current parser and provider source.
- Primary paths: declaration loop -> loop adapter -> accepted walker -> one Definition/OwnedDefID/local Binding per leaf; assignment loop -> recursive assignment target write extraction -> AccessWrite facts and zero declarations.
- Sibling surfaces to preserve: identifier variables, destructured variables, parameters, catches, imports, type inference, calls/member accesses outside loop targets, module/class/function/catch scopes, ScopeIR normalization, JS/TS parity.
- Forbidden fallback: no name-based declaration rescue; no declaration from assignment-form target; no use of `variable_declarator` assumptions for loop AST; no invented type/reference/graph/control-flow facts.
- Later boundary: graph projection and lexical resolution remain P3-C; this slice proves ScopeIR ownership and AccessWrite facts only.

## Exact proposed editable boundary

The following is a plan-update proposal, not current edit authorization.

| File | Touch mode after planner authorization | Existing symbol allowed to change | New private symbols allowed |
|---|---|---|---|
| `internal/providers/tsjs/definitions.go` | editable only for loop declaration discrimination/emission and loop-specific ownership | `(*collector).emitDefinitionKind` | `emitLoopBindingPattern`, `loopBindingContext`, `loopBindingScopeID`, `addLoopBindingLeafDefinition` (names may be normalized to local style, responsibility may not expand) |
| `internal/providers/tsjs/scopes.go` | editable only for lexical `let`/`const` loop `ScopeBlock` candidate | `(*collector).collectScopeCandidateForKind` | `isLexicalLoopDeclaration` if needed |
| `internal/providers/tsjs/references.go` | editable only for assignment-form loop target writes and suppression of the current member-target false read | `(*collector).emitReferenceKind` | `emitLoopAssignmentWrites`, `emitAssignmentTargetWrites` if needed |
| `internal/providers/tsjs/extract_test.go` | editable after production code for focused real-parser/real-`Extract` loop tests and one stale sibling transition | `TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts` | focused declaration, assignment-write, scope/shadowing, and TS/JS parity test functions |

Preserve-/inspect-only: `binding_patterns.go`, `extract.go`, `nodes.go`, `types.go`, `imports.go`, every `internal/scopeir` contract, all accepted golden artifacts, Vue/resolution code/tests, graphaccuracy, graph projection/resolution/persistence, target source, and every governance file owned by main.

## File-detail and impact evidence

All four file-detail results were current: indexed/current commit matched entry `HEAD`, `stale=false`, `changedSinceAnalyze=false`.

| Candidate | Symbols | Related files | In / out / local | Unresolved | File-detail relationship risk | Linked tests |
|---|---:|---:|---:|---:|---|---:|
| `definitions.go` | 106 | 22 | 6 / 91 / 13 | 250 | HIGH | 3 |
| `scopes.go` | 35 | 18 | 12 / 44 / 11 | 83 | HIGH | 4 |
| `references.go` | 33 | 15 | 5 / 35 / 3 | 70 | HIGH | 3 |
| `extract_test.go` | 455 | 24 | 13 / 210 / 135 | 0 | LOW | 3 |

Upstream file impact:

| File | Severity | Direct impacted symbols | Affected files | Affected file split | Processes / flows |
|---|---|---:|---:|---|---:|
| `definitions.go` | MEDIUM | 9 | 2 | `definitions.go:8`, `types.go:1` | 0 / 0 |
| `scopes.go` | MEDIUM | 9 | 2 | `scopes.go:6`, `definitions.go:3` | 0 / 0 |
| `references.go` | LOW | 1 | 1 | `references.go:1` | 0 / 0 |
| `extract_test.go` | LOW | 0 | 0 | none | 0 / 0 |

Exact canonical-UID upstream impact:

| Existing symbol | Canonical UID basis | Risk | Impacted / modules / processes | File-layer affected files |
|---|---|---|---:|---:|
| `collector.emitDefinitionKind` | file-detail UID at `definitions.go#14:0-92:1` | LOW | 0 / 0 / 0 | 0 |
| `collector.collectScopeCandidateForKind` | file-detail UID at `scopes.go#33:0-44:1` | LOW | 0 / 0 / 0 | 0 |
| `collector.emitReferenceKind` | file-detail UID at `references.go#13:0-56:1` | LOW | 0 / 0 / 0 | 0 |
| `TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts` | file-detail UID at `extract_test.go#854:0-948:1` | LOW | 0 / 0 / 0 | 0 |

Three production owners carry HIGH relationship risk at the file-detail layer. This is a blast-radius warning requiring the narrow preservation gates below, not an edit prohibition. No candidate has HIGH or CRITICAL upstream file-impact severity, and every exact existing-symbol impact is LOW `0/0/0`.

Four exploratory calls incorrectly passed the UID as the positional target and returned `UNKNOWN/not found`; they are rejected command history, not evidence. The accepted commands used the documented `--uid` flag and `dispatchMode=explicit`.

## Required plan-boundary change

Yes. Orchestration must update P3-B2A before implementation:

1. Replace generic “exact loop-declaration owner and focused test owner” with the four exact files above.
2. Move only `collector.collectScopeCandidateForKind` loop handling from inspect-only to editable for lexical `let`/`const` loop ownership.
3. Move only `collector.emitReferenceKind` loop handling from inspect-only to editable for assignment-form AccessWrite truth and member false-read suppression.
4. State explicitly that declaration form is detected by non-empty `for_in_statement.kind`, while assignment form has no `kind` and creates zero declarations.
5. State `var -> nearest function/module`; `let|const -> exact loop ScopeBlock`; assignment -> no declaration scope.
6. Keep graph/reference resolution, new ScopeIR contracts, `await`/`using` semantic expansion, general block-scope repair, control-flow/process modeling, and all later slices out of scope.

Until that planner update and explicit resume authorization exist, no production or test file is editable.

## Mandatory future implementation and validation gates

After Orchestration authorization, the implementation slice must use this order:

1. Production code first within the exact four-file boundary; tests remain unchanged until production behavior is correct.
2. Declaration matrix: TS and JS array/object patterns, aliases/default/rest/holes, and `var|let|const` identifiers; one accepted leaf -> one Definition, one OwnedDefID, one local Binding, correct `BindingContextForIn|ForOf`, exact ranges/path/modifiers.
3. Scope matrix: module and nested function cases; `var` function/module ownership; `let|const` exact loop ScopeBlock; same-name outer/loop/inner identities distinct; no leakage to siblings.
4. Assignment matrix: `[a]`, `{x}`, plain identifier, member target, and grammar-supported representable targets; exactly zero declarations/BindingLeaves and correct AccessWrite facts with no duplicate/false read.
5. TS/JS parity plus pinned grammar difference controls; no `await`/`using` claim without separate authority.
6. Preservation: accepted P3-A walker, P3-B variables/identifier/type-inference miss/assignment-zero/import delta, P3-B1 parameters/type-only exclusion, P3-B2 catches/optional catch/exact catch scope, existing calls/accesses/types/imports, ScopeIR normalization, Vue/resolution goldens.
7. Before the full build, terminate every build-artifact holder and prove the build lock free. Run the repository full build on final production/test bytes before validation.
8. Validate the nearest real ScopeIR boundary, then focused tests, sibling preservation, TSJS/ScopeIR packages, four-package provider/resolution boundary, and repo-native `./cmd/... ./internal/...` gate. Build/test timings are validation evidence, not benchmarks unless product performance changes.
9. Remove task temp/debug output. Obtain independent Supervisor PASS, run final `anvien detect-changes --repo E:\Anvien --scope all`, refresh main-owned ledgers, create one isolated P3-B2A commit, and push only when Orchestration authorizes those later actions.

No build, behavior test, detect-changes, Supervisor review, stage, commit, or push was performed in this discovery turn.

## Cleanup and final handoff state

- The only temporary tree created was `E:\Anvien\.tmp\p3b2a-parser-evidence`; it was verified inside the repository and removed completely before this report was written.
- No target repository or target-local output was touched.
- No production, test, fixture, golden, roadmap, plan, evidence, benchmark, or actual-status byte was edited.
- This report is the sole durable file created by the coder lane.
- Final staged state must remain `0`; final Git manifest and this report's byte size/SHA-256 are captured externally after write so the artifact can be hashed without self-reference.

## Exact next action for Orchestration main

Invoke the planner skill and update only the P3-B2A scope/evidence/status rows to authorize the exact four-file/four-existing-symbol boundary above. Then send an explicit resume command to this same visible coder task:

`RESUME P3-B2A IMPLEMENTATION: definitions.go + scopes.go + references.go + extract_test.go only; preserve P3-A/P3-B/P3-B1/P3-B2 and keep P3-C+ locked.`

Do not authorize implementation before the planner update is durable.
