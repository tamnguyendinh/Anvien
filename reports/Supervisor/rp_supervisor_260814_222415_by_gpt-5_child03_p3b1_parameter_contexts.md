# Supervisor Report: Child 03 P3-B1 TypeScript Parameter Contexts

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260814_222415_by_gpt-5_child03_p3b1_parameter_contexts.md`
- Review time: `2026-08-14 22:24:15 +07:00` (`SE Asia Standard Time`)
- Reviewer: `gpt-5`
- Repository: `E:\Anvien`
- Branch / HEAD: `master` / `06229bea5735e75c8d8a476c738f627bf93def8d`
- Upstream: `origin/master` at `ea1e27710966eb54a4533255dc80a2ac0c2120d8`; HEAD is one Owner-owned orchestration-skill commit ahead and that commit is outside P3-B1 semantics.
- Scope reviewed: Child 03 slice P3-B1 only — TypeScript parameter-context binding extraction and its exact dependent test/oracle reconciliation.
- Claim reviewed: every legal function, method, constructor, parenthesized-arrow, and unparenthesized-arrow parameter leaf emits exactly one parameter BindingLeaf, Variable Definition, callable-scope OwnedDefID, and local Binding while non-parameter syntax and all locked sibling boundaries remain unchanged.
- Authority used: Owner delegation; `AGENTS.md`; full `working-rules` and `supervisor` skills; active roadmap; complete Child 03 plan, evidence, benchmark, and actual-status ledgers; current source/diff; pinned Tree-sitter JavaScript/TypeScript grammars; fresh Anvien evidence; independent build/tests; direct JSON multiset audit.
- Related artifact: `reports/coder/rp_coder_260814_215327_by_gpt-5_child03_p3b1_parameter_contexts.md`, `16,257` bytes, SHA-256 `BF6CA0B6F68AE2ECFC23F3F1CF27F9E8EF251961BEC78AE14FE9AC3F7B73E65F`.
- Reserved and not performed: `E3-P3B1-DETECT1`, `E3-P3B1-COMMIT1`, staging, commit, push, cleanup deletion, and every P3-B2/P3-C action.

## Executive Summary

- Problem: the candidate correctly emits parameter facts for its focused callable matrix and passes all declared build/regression gates, but its dispatch does not prove that a Tree-sitter `required_parameter` or `optional_parameter` node is actually a callable parameter.
- Decision: REJECT. TypeScript deliberately aliases labeled tuple members to the same `required_parameter` / `optional_parameter` node kinds. The candidate walks every named node, accepts either kind globally, falls back from field `pattern` to field `name`, and accepts any nearest function ancestor. Consequently, legal type-only tuple labels nested under a runtime callable are projected as false parameter declarations in that callable scope.
- Required outcome: repair only stable invariant `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1`, add a focused former-failure oracle, rerun only evidence invalidated by the final repair bytes, and return the same P3-B1 slice for independent re-review. P3-B1 is not accepted for planner/detect/commit closure. P3-B2, P3-B2A, P3-C, P3-C2, and all later slices remain locked and unaccepted.

## Blocking Findings

### [HIGH] P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1 — type-only labeled tuple members become false runtime parameter declarations

File: `internal/providers/tsjs/definitions.go:26`

Issue: `emitDefinitionKind` dispatches every `required_parameter` and `optional_parameter` node to parameter extraction without checking its syntactic role. This node-kind test is insufficient for the pinned TypeScript grammar.

Evidence:

1. `internal/providers/tsjs/extract.go:31-35` invokes `emitDefinitionKind` for every named node in the preorder `walkKind`; there is no context filter before dispatch.
2. `internal/providers/tsjs/definitions.go:26-27` accepts both parameter-looking node kinds globally. `parameterPattern` at lines `92-98` first reads field `pattern`, then falls back to field `name`. `parameterCallableScopeID` at lines `166-178` accepts the nearest ancestor recognized as a function scope; it does not verify membership in that callable's own `formal_parameters` field.
3. The pinned grammar `github.com/tree-sitter/tree-sitter-typescript@v0.23.2/common/define-grammar.js:748-749` aliases `tuple_parameter` to `required_parameter` and `optional_tuple_parameter` to `optional_parameter`. Those tuple nodes use field `name`, so they are explicitly accepted by the candidate fallback.
4. The pinned grammar corpus `github.com/tree-sitter/tree-sitter-typescript@v0.23.2/test/corpus/types.txt:1512` records legal labeled tuple syntax `type t = [a: A, b?: B, ...c: C[]]`; its expected AST contains `required_parameter`, `optional_parameter`, and a rest `required_parameter` under `tuple_type`.
5. Deterministic reachable example: in `function f(arg: [tupleLabel: string]): [resultLabel: boolean] { const local: [bodyLabel: number] = [1] }`, the outer real parameter `arg` is valid, but each labeled tuple member is also a named `required_parameter`. For `tupleLabel`, `resultLabel`, and `bodyLabel`, `parameterPattern` returns field `name`, and `parameterCallableScopeID` climbs to `f`. Each therefore reaches `extractParameterBindingPattern` and `addParameterBindingLeafDefinition`, producing a false parameter BindingLeaf, Variable Definition, OwnedDefID, and local Binding. The same flaw applies to type-level function/constructor parameter nodes nested beneath a runtime callable.
6. Candidate tests cover actual parameters, explicit `this`, and one plain JavaScript arrow, but contain no labeled-tuple or nested function-type disambiguation oracle. All green tests therefore exercise a narrower syntax-role set than the claimed invariant.

Why this blocks acceptance: P3-B1 requires one fact set per legal callable parameter leaf and locked type/context behavior. A labeled tuple member is legal TypeScript but is not a runtime callable parameter declaration. Emitting it violates exact context classification, no-false-declaration conservation, lexical ownership, and the preserve-only type path. Generic graph consumers would then expose invented Variable nodes and `DEFINES` edges, so the defect is not confined to an unused helper.

Fix direction:

- Restrict `required_parameter` / `optional_parameter` dispatch to nodes proven to belong to the owning callable's own formal-parameter surface. Do not accept a node merely because some function is an ancestor.
- Preserve the TypeScript-only unparenthesized-arrow path and the current explicit-`this` behavior.
- Explicitly exclude labeled tuple members and nested function/constructor type-signature parameters from an outer runtime callable's parameter facts. Do not solve this by special-casing names or by changing the accepted P3-A walker/ScopeIR contract.
- Add one focused real-`Extract` former-failure test containing actual parameters whose annotations/return/body types include labeled tuples and nested function types. Assert actual parameters emit once, while every type-only label emits zero BindingLeaves, Variable Definitions, OwnedDefIDs, local Bindings, and extraction diagnostics.

Exact repair boundary:

- Production: `internal/providers/tsjs/definitions.go` only.
- Focused test: `internal/providers/tsjs/extract_test.go` only.
- Keep the other six candidate/oracle files byte-locked unless the new focused test proves a legitimate output delta and Orchestration first reopens the boundary with fresh plan/impact authority. No production Vue/resolution edit is authorized.

Re-review evidence required:

- Fresh source/diff proof closing this one role-disambiguation invariant.
- The new exact former-failure test, plus the currently declared focused, former-failure, P3-A/P3-B preservation, four-package, and repo-native product gates on final repair bytes.
- Holder-clean full build on final repair bytes.
- Direct JSON multiset audit only if any oracle byte changes or a final test exposes output drift.
- No orchestration-owned final detect before Supervisor PASS.

## Source-Level Clearance Notes

- `internal/providers/tsjs/definitions.go`: BLOCKED at lines `26-32`, `92-115`, and `166-206`. Legal actual parameters in the focused matrix are emitted with the claimed ranges, selections, modifiers, provenance, and callable ownership, but syntax-role discrimination is incomplete and produces false facts for type-only aliases.
- `internal/providers/tsjs/extract_test.go`: BLOCKED for acceptance coverage. The three new tests are internally coherent and independently passed, and the two variable assertions were narrowed only to `BindingContextVariable`; however, no test distinguishes callable parameters from same-kind type-only nodes.
- `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json`: CLEAR for the reviewed fixture; direct multiset delta is exactly `+1 Variable:repo:::` and `+1 Variable:user:::` with no removals or other top-level change.
- `internal/providers/vue/extract_test.go` and `internal/providers/vue/testdata/vue_scopeir_signature.golden.json`: CLEAR for the reviewed fixture; the only assertion change is `DEFINES 10 -> 13`, and the golden delta is exactly `repo`, `user`, `value` Definitions.
- `internal/resolution/graph_parity_test.go`: CLEAR for the authorized test-only reconciliation. Historical TypeScript counts remain compared independently from current Go counts, and the new node reconciliation validates `30 - 24 = 6` with a non-empty classification.
- `internal/resolution/testdata/typescript_graph_baseline_counts.golden.json`: CLEAR. `source`, `generatedAt`, `resolutionOnly`, `fullGraph`, and both `mustMatch*` arrays are JSON-identical to HEAD. Historical TypeScript remains `24` nodes / `54` relationships / `18` `DEFINES`; full graph remains `28/67/18`. Go reconciliation is separate at nodes `30`, delta `+6`, and `DEFINES 25`, while TypeScript values remain `18/18`.
- `internal/resolution/testdata/typescript_graph_signature.golden.json`: CLEAR. Exact multiset delta is six Variable nodes and six `DEFINES` relationships, zero removals, and zero other relationship/source-owner delta.
- Production Vue, resolution, graph projection, ScopeIR contracts, P3-A walker, P3-B collector plumbing, type inference, references, imports, catch, and loop owners: CLEAR as unchanged in this candidate; no P3-C acceptance is implied.

## Git Boundary and Byte Integrity

- Initial and final branch/commit: `master` / `06229bea5735e75c8d8a476c738f627bf93def8d`.
- `origin/master`: `ea1e27710966eb54a4533255dc80a2ac0c2120d8`; `git diff origin/master..HEAD` contains only `internal/aicontext/skills/orchestration/SKILL.md`.
- Staged files before report: `0`.
- Tracked worktree boundary before report: exactly four governance files plus eight candidate/oracle files.
- Untracked before report: only the coder report.
- `git diff --check`: exit `0` before and after build/tests.
- No stage, commit, push, reset, checkout, stash, rebase, amend, cleanup deletion, or final detect occurred.

### Governance hashes preserved

| File | SHA-256 |
|---|---|
| roadmap | `F643F8E2DD69198E45BDC4AE417F14F7FC90DC1151D59A0319B6380B560A0BAB` |
| active plan | `51633CE900EF79589D0A6B30E57DB8BBC565256CFD2E8201D0B4823C4EC25EB8` |
| evidence ledger | `8920E32F8A4347F6003C8AAD3D852EACD3007F30F2A5583216410E059B8DC2B9` |
| actual-status ledger | `FDFFAE7852273CD29DF4E8BABBAE1080C56BC14C7A7EDF772DEBF656F957C166` |
| benchmark ledger | `B5B01879C59E0556A2A8E04D0C54693F8EF41DBB574D64F621BDDB7A4EA77F1E` |

### Candidate/oracle hashes preserved

| File | Bytes | SHA-256 |
|---|---:|---|
| `internal/providers/tsjs/definitions.go` | 12,774 | `416C997DF55D885A5F92353441D6AB943927C2DC2C49B2060983ACE21D920DC3` |
| `internal/providers/tsjs/extract_test.go` | 53,736 | `F096D08EC1482666A58E79E52DE4D21DC9E6397CF939E36AD2CC201E9B34A158` |
| `internal/providers/tsjs/testdata/typescript_scopeir_signature.golden.json` | 1,788 | `AD09CD8B8F93B3A4A4D38F6939A05424BD09C89C35E76600152057BFD5884398` |
| `internal/providers/vue/extract_test.go` | 11,895 | `B2F9C261F68DD369FAFE9B4178AAD342896A1A5683325EBAD5F697D983EBC04D` |
| `internal/providers/vue/testdata/vue_scopeir_signature.golden.json` | 1,757 | `3317D30FFC8E640B12EE45C3B95B6DDEAA12E7C66D7BE477556D2978D5A6A03C` |
| `internal/resolution/graph_parity_test.go` | 4,322 | `0607074F015487D0D836C24C55A379A0CAD76292C0AF63049700CF5C8223430A` |
| `internal/resolution/testdata/typescript_graph_baseline_counts.golden.json` | 2,964 | `5DC4FDC0439CC8D6D25C5E72CE2DD5744C094DCDDB5C0B7C46CA9FCD0D304A61` |
| `internal/resolution/testdata/typescript_graph_signature.golden.json` | 8,093 | `FF9396290CB2678743B70A51ABF649B13CC0D29B17CDB54CBE42EE9B1DB3CB67` |

## Evidence Checked

Passed:

- Full-read authority: roadmap, complete active plan, complete evidence/benchmark/actual-status ledgers, coder report, source, and every touched diff.
- Fresh Git/hash reconstruction: exact HEAD/upstream, staged `0`, candidate manifest and all supplied SHA-256 values matched.
- `anvien --help`: exit `0`.
- Fresh `anvien analyze --force`: exit `0`; `1,644` scanned / `689` parsed code / `0` failed; `100,694` nodes / `140,486` relationships; indexed/current commit both `06229bea...`; stale `false`.
- Fresh Anvien file/symbol evidence: `definitions.go` relationship risk HIGH; upstream file impact MEDIUM with `8` impacted symbols across `definitions.go` and preserve-only `types.go`, no process/flow; exact current UID `collector.emitDefinitionKind` LOW `0/0/0`. Every opened test/golden owner is LOW with zero affected files/processes/flows. HIGH/MEDIUM were treated as scope warnings, not rejection reasons.
- Holder gate: analyze lock free; three exact global Anvien MCP child/parent chains (`15480/15072`, `14620/14012`, `9992/14372`) were terminated; unrelated Playwright PID `9352` was preserved; recheck found zero global Anvien holder and lock free.
- Independent `npm run full-build`: exit `0` in `110.7s`; CLI `1.2.8`; Web build `2,943` modules; final analyze `1,644/689/0`, graph `100,694/140,486`. npm allow-scripts and Vite import/chunk messages were non-failing warnings.
- Focused parameter/TSJS parity command: exit `0`; `4/4` tests PASS.
- Exact four former-failure Vue/resolution command: exit `0`; `4/4` tests PASS.
- P3-A/P3-B preservation command: exit `0`; `10/10` tests PASS.
- Nearest package command for TSJS/ScopeIR/Vue/resolution: exit `0`; `4/4` packages PASS.
- Repo-native `go test ./cmd/... ./internal/... -count=1`: exit `0` in `159.8s`; all listed product packages PASS.
- Independent HEAD-versus-worktree JSON audit: exact TSJS `+2`, Vue `+3`, resolution `+6 Variable/+6 DEFINES`; zero removal/other relationship drift; historical TypeScript baseline fields unchanged.
- Historical source truth: the original conversion benchmark records `npx tsx .tmp\resolution_baseline_benchmark.ts`, `5` files, `24` nodes, `54` edges, and `DEFINES=18`; Git history shows the baseline file has no intervening tracked revision before this candidate.
- Cleanup/boundary: before this report the only untracked path was the coder report; no lane-named temp/debug path existed; existing `.tmp` entries predated this review; final candidate/governance hashes remained exact.

Failed:

- `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1`: source and pinned grammar prove same-kind type-only nodes reach parameter declaration emission.

Not run:

- No new executable tuple-label probe was written because Supervisor authority permits exactly one durable report and forbids any production/test/temp/probe edit. The rejection is based on deterministic pinned-grammar AST plus current source control-flow proof; the repair must add the executable former-failure test.
- No `E3-P3B1-DETECT1`, staging, commit, push, or later-slice validation was run because those actions are reserved or out of scope.
- Browser/Playwright is N/A for this non-UI ScopeIR slice.

## Locked Passing Invariants for Repair

The following are cleared for this exact candidate and must not be rerun merely to challenge this verdict while their bytes/evidence remain unchanged:

- Git boundary, governance hashes, candidate/oracle hashes, cleanup state, and `git diff --check`.
- Fresh Anvien impact boundary for all already opened files/symbols; no Vue/resolution production owner is opened.
- Exact JSON multiset deltas and historical baseline truth.
- Current explicit-`this` zero-declaration/zero-diagnostic control, TypeScript unparenthesized-arrow control, plain JavaScript arrow preservation, ranges/selections/modifiers/provenance, callable-scope ownership, and nested callable shadowing for the tested actual-parameter forms.
- Current full build, focused `4/4`, former failures `4/4`, P3-A/P3-B preservation `10/10`, four-package PASS, and repo-native product PASS.

Repair necessarily changes `definitions.go` and `extract_test.go`; that byte change invalidates build and behavior evidence, so those build/test gates must be rerun once on final repair bytes. The six cleared dependent test/oracle files and their JSON audit remain locked and must not be edited or re-audited unless final repair output actually changes them. Do not rerun or reopen P3-B2/P3-C work.

## Invariant Closure

- Affected invariant: parameter syntax-role discrimination plus one-source-leaf-to-one-declaration conservation with zero false declarations.
- Sibling surfaces checked: P3-A walker, P3-B variable facts, collector traversal, scope construction, TypeBinding path, explicit `this`, TypeScript/JavaScript arrow grammar, TypeScript tuple grammar aliases/corpus, Vue generic adapter, generic resolution projection, goldens, historical baseline, catch/loop/assignment/import/call/reference preservation.
- Residual unverified/broken same-invariant surface: labeled tuple members and nested type-signature parameters are not excluded from an outer runtime callable. This is a confirmed source blocker.
- Closure status: open; P3-B1 cannot be accepted.

## Required Fix List For Resubmission

1. Close only `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1` inside `definitions.go` without changing the P3-A walker, ScopeIR contract, type inference, traversal, scope construction, Vue/resolution production, or later contexts.
2. Add the focused legal tuple/function-type former-failure test in `extract_test.go` with exact zero-fact assertions for type-only labels.
3. Preserve the other six candidate/oracle bytes unless actual final output proves a delta and Orchestration explicitly reopens the boundary.
4. On final repair bytes, rerun holder-clean full build, the new former-failure test, current focused/former-failure/preservation/package gates, and repo-native product regression once; record what each proves.
5. Return the same P3-B1 slice for independent Supervisor re-review. Do not run final detect or begin P3-B2/P3-C before PASS.

## Overall Evaluation

The candidate is well-scoped, its declared fixtures and downstream reconciliation are exact, and its build/regression evidence independently passes. Acceptance nevertheless fails because parameter-context dispatch is keyed to a Tree-sitter node name that is not unique to callable parameters. The missing syntax-role gate creates false declaration facts on legal TypeScript and violates the core P3-B1 conservation invariant. Verdict is therefore REJECT for P3-B1 only; no conclusion is made for later slices.

## Handoff

- Recipient: Orchestration main in task `Điều hành toàn bộ công việc`.
- Action: route only stable rejected invariant `P3B1-PARAMETER-CONTEXT-DISAMBIGUATION-1` back to the P3-B1 repair lane with the exact two-file repair boundary above.
- Acceptance state: P3-B1 REJECTED; no planner/detect/commit closure; P3-B2 and later remain locked.
