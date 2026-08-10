# Supervisor Report: Child 01 P1-B Range and Identity Inputs

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260810_115437_by_gpt-5-codex_child01_p1b.md`
- Review time: `2026-08-10 11:54:37 +07:00` (Asia/Bangkok)
- Reviewer: `gpt-5-codex`
- Repo/project: `E:\\Anvien`
- Scope reviewed: Child 01 / P1-B candidate at base HEAD `8cbf6006d53836494d5cc862baea77939af9f938` plus the current P1-B worktree; no P1-C identity, P1-D collision, persistence, target `4/4`, or later-Child behavior is reviewed or accepted.
- Claim reviewed: P1-B makes source-backed full/selection ranges and the existing lexical-owner/meaning inputs available at the TSJS provider/ScopeIR boundary, preserves later-slice boundaries, and has current, non-contradictory ledgers sufficient for Supervisor acceptance.
- Authority used: latest user direction to follow the active plan without inventing architecture; `AGENTS.md`; Supervisor skill; Child 01 plan, evidence, benchmark, and actual-status ledgers; graph-accuracy contract; current source/diff/runtime evidence.
- Related artifacts: `reports/coder/rp_coder_260810_113052_by_gpt-5-codex_child01_p1b.md`, `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1`, prior Child 01 P0/P1-A Supervisor reports, and the bounded earlier P1-B investigation report.

## Executive Summary

- Problem: decide whether the current P1-B implementation, validation, and live ledgers close only the provider/ScopeIR range and identity-input slice.
- Decision: REJECT. The production diff and fresh behavior evidence clear the P1-B code boundary, and the editor-owned executable lock is not a source compile failure. Acceptance is nevertheless blocked because the live actual-status ledger contradicts itself: its current matrix and R5 row say P1-B range/selection inputs are correct, while its Detailed Findings still say `DefinitionFact` has no selection range, classify the boundary as partial, and direct P1-B to begin as future work.
- Required outcome: update only the stale P1-B Detailed Findings in the actual-status ledger so they match current source/evidence and preserve P1-C/P1-D/P1-E boundaries, then resubmit the unchanged slice for Supervisor review. No production or test repair is required by this verdict.

## Blocking Findings

### [HIGH] The live actual-status authority contains mutually exclusive P1-B states

File: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md:141`

Issue: the status matrix and refresh log record the implemented P1-B boundary, but the Detailed Findings retain the superseded P0 state and future-action instructions. At lines 82-83 and 102, the ledger says `SelectionRange` exists, the TSJS coordinate contract is explicit, and these inputs moved to `correct`. At lines 141 and 143, it still classifies the shared IR input as `partial` and says P1-B must select owners later. At lines 151, 163, and 165, it says `DefinitionFact` has no selection-range field, the range boundary is `partial`, and P1-B still needs to establish the provider contract. Those statements cannot all describe the current candidate.

Evidence: current source has `SelectionRange *Range` at `internal/scopeir/facts.go:10-11`; TSJS documents its one-based-line, zero-based UTF-8 byte-column, exclusive-end contract at `internal/providers/tsjs/nodes.go:18-28`; `collector.addDefinition` retains construct range and assigns the declaring-token range at `internal/providers/tsjs/definitions.go:95-113`. Fresh focused tests and the 61-package product regression pass. The Child plan requires matching checklist/actual-status/evidence/benchmark updates when state changes at line 33, makes ledger refresh an implementation-slice gate at line 92, and explicitly requires P1-B range/identity-input actual-status rows to be refreshed at line 223.

Why this blocks acceptance: the plan ledger is active authority for which slice is open and what remains partial. A contradictory current-state section leaves a required acceptance action outstanding and can send the next implementation step back into P1-B or cause P1-C to use a false premise. Supervisor PASS requires the full claim and its live authority to be current and non-contradictory, with no required follow-up remaining.

Fix direction: refresh only the stale `Identity, lexical owner, and meaning` and `Ranges and selection range` Detailed Findings. Record that the provider/ScopeIR input boundary is correct for P1-B; retain graph consumption/identity as partial or wrong only at the already-owned P1-C/P1-E boundaries; retain shared `Range`, `PositionIndex`, scope containment, collision, persistence, and target behavior unchanged. Do not add architecture, change production/test code, open P1-C, or claim target `4/4`.

Re-review evidence required: the exact actual-status diff, `git diff --check`, and a consistency check showing the header, Current Status Matrix, R5, Detailed Findings, and Next Phase Status Decisions describe one P1-B state while `E1-P1B-REVIEW1`, detect-changes, and commit remain pending.

## Source-Level Clearance Notes

- `internal/scopeir/facts.go`: clear for P1-B source behavior. Lines 3-30 add only the optional declaring-token `SelectionRange`; existing owner/label fields and `ScopeFact` ownership/bindings remain in their established owners. Fresh post-diff impact is CRITICAL (`592` impacted symbols, `81` affected-file rows, `23` modules, `63` processes; app layers `backend 181`, `backend_test 411`), which is a blast-radius warning and was contained by tests.
- `internal/providers/tsjs/nodes.go`: clear. Lines 18-28 document the existing source-backed TSJS point conversion without changing coordinates. Fresh impact for `nodeRange` is MEDIUM (`8` symbols, `4` files, `1` module, `0` processes).
- `internal/providers/tsjs/definitions.go`: clear. Lines 79-141 change only `collector.addDefinition`; lines 95-113 preserve the construct range and add the `nameNode` selection range. The non-identifier variable-declarator branch at lines 64-68 is unchanged. Fresh upstream impact for `collector.addDefinition` is LOW (`0` impacted symbols/files/modules/processes).
- `internal/providers/tsjs/definition_position_inputs_test.go`: clear. Lines 12-176 cover construct/selection positions, LF/CRLF, Unicode byte columns, tabs, exclusive endpoints, exact nested lexical ownership/bindings, same-lexeme provider labels, member owner ID, JSON round-trip, and preservation of the Child 03 binding-pattern boundary.
- Child 01 plan ledgers: blocked only by the actual-status contradiction described above. Evidence and benchmark rows otherwise keep P1-C, P1-D, persistence, and target validation pending.

## Evidence Checked

Passed:

- Fresh `anvien analyze --force --json` on the unchanged candidate: `1,562` scanned, `677` parsed/parsed-code, `0` failed, `85,227` nodes, and `124,129` relationships; graph commit/current commit are both `8cbf6006d53836494d5cc862baea77939af9f938`, `stale=false`.
- Fresh file-detail after that refresh: `facts.go` `120` symbols/risk medium; `nodes.go` `42`/high; `definitions.go` `31`/high; the new test `46`/low. All report `stale=false` and `changedSinceAnalyze=false`.
- Fresh impact: `DefinitionFact` CRITICAL `592` impacted symbols / `81` affected-file rows / `23` modules / `63` processes; `nodeRange` MEDIUM `8/4/1/0`; `collector.addDefinition` LOW `0/0/0/0`. The pre-edit ledger values remain valid preflight evidence; the larger current shared-contract count includes post-diff test consumers.
- Direct source inspection confirms parser input is UTF-8 bytes (`internal/parser/parse.go:78-84`), tree-sitter points are zero-based, lexer columns advance by UTF-8 byte width, and `Utf8Text` slices `[StartByte:EndByte]`; the TSJS one-based-line/byte-column/exclusive-end comment is therefore source-backed.
- Fresh focused matrix: `go test ./internal/providers/tsjs -run '^Test(ExtractDefinitionPositionInputs|ExtractDefinitionOwnerMeaningAndLexicalInputs|P1BPreservesBindingPatternBoundary)$' -count=1 -v` PASS.
- Fresh owner regression: `go test ./internal/providers/tsjs ./internal/scopeir ./internal/parser -count=1` PASS.
- Fresh product regression: `go list -e ./...` excluding only `*/anvien/test/fixtures/*`, followed by `go test`, PASS for `61/61` product packages.
- Independent current-source native compile to a repo-local alternate output using the production `ladybugdb` tag and native flags PASS; the resulting CLI reports `1.2.8`. Supervisor-created binaries were removed after validation.
- Full-build applicability check: the successful canonical full build recorded under `E1-P1B-BUILD1` occurred after all three production files were last modified; `anvien/bin/anvien.exe` is timestamped `11:12:18 +07`, while production files are timestamped `11:00:41-11:00:47 +07`. Only the focused test file changed later (`11:22:53 +07`), and fresh Go compilation/tests cover that final test state. Current `anvien version` is `1.2.8`.
- `git diff --check` passes. The production diff is limited to the three P1-B owners, with one package-owned test; no P1-C/P1-D/persistence/target source was changed.
- Verification freshness: current for the reviewed worktree and base HEAD; the report write is the only Supervisor mutation.

Failed:

- A later canonical full-build retry recorded in `E1-P1B-BUILD1` could not overwrite `anvien/bin/anvien.exe` because editor-owned `anvien mcp` processes execute through the global npm junction. Current `anvien doctor processes --json` independently shows four such executable processes and marks them expected to stay running. This is not treated as a source/build-input failure because the complete build already passed after the production diff, production inputs did not change afterward, and the final current-source native compile plus product regression pass. It does not remove the ledger blocker above.

Not run:

- `anvien detect-changes` and the isolated P1-B commit: intentionally pending until an unconditional Supervisor PASS, per the plan.
- Target `E:\\cheapapp.org`, bounded `4/4`, graph identity consumption, collision behavior, projection/persistence/readers, and P1-C/P1-D/P1-E validation: outside P1-B and deliberately not used for this verdict.
- UI/Playwright: not applicable to this non-UI provider/ScopeIR slice.

## Invariant Closure

- affected invariant: provider-supplied definition facts must retain construct and declaring-token ranges with explicit source-backed TSJS coordinates, while existing lexical-owner/meaning inputs survive normalization/JSON and later-slice ownership remains intact; the live plan authority must accurately record that boundary.
- sibling surfaces checked: TSJS scope ownership and bindings, `Range`/PositionIndex and containment preservation, ScopeIR normalization/JSON, other-provider optional omission, non-identifier binding-pattern preservation, identity/collision/projection/persistence boundaries, plan/evidence/benchmark/actual-status ledgers, and current Git diff.
- residual unverified same-invariant surfaces: the technical P1-B source/runtime invariant has no uncovered surface in this scope. The live actual-status narrative remains contradictory and therefore prevents process-invariant closure. P1-C identity consumption, P1-D collision handling, and P1-E target/integration remain intentionally pending, not residual P1-B gaps.

## Required Fix List For Resubmission

1. Reconcile the stale actual-status Detailed Findings at lines 125-165 with the already-current matrix/R5/source evidence, without changing production/test code or later-slice ownership.
2. Provide a clean ledger-only diff and consistency proof; keep `E1-P1B-REVIEW1`, detect-changes, commit, P1-C, P1-D, and P1-E pending until re-review.

## Overall Evaluation

The implementation is technically scoped and independently validated: it adds the optional selection range at the existing fact owner, records the actual TSJS coordinate contract, preserves lexical/meaning inputs and later-Child boundaries, and does not invent a new graph architecture. The editor-owned executable lock does not invalidate the applicable successful full build or the fresh current-source compile/regression evidence. The candidate still cannot be accepted because its active actual-status authority makes mutually exclusive claims about whether P1-B has happened. The required correction is ledger-only and narrowly defined; this REJECT does not reopen production design.
