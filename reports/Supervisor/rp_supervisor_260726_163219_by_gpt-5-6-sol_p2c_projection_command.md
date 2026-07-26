# Supervisor Report: P2-C projection/command causal report

Verdict: REJECT

## Metadata

- Report file: `rp_supervisor_260726_163219_by_gpt-5-6-sol_p2c_projection_command.md`
- Review time: 260726 163219 +07
- Reviewer: `gpt-5-6-sol`
- Repo/project: Anvien investigation against `E:\cheapapp.org`
- Scope reviewed: `reports/Investigation/20260726_161241_p2c_projection_command_root_cause.md` plus current P2-C probes and ledgers
- Claim reviewed: the report accurately separates upstream graph loss from downstream command/projection behavior, including its derived-process differential
- Authority used: current user boundary, `E:\Anvien\AGENTS.md`, current target graph, current Anvien source, current probe code/output, and fresh Supervisor reruns
- Related artifacts: `p2c/process_control_probe.go`, `p2c/process_control_probe_output.json`, five fresh Supervisor process-probe outputs, P2-C evidence/benchmark/actual-status entries

## Executive Summary

- Problem: the durable P2-C report must agree with its current machine-readable evidence before any bounded causal claim can be accepted.
- Decision: REJECT. The report's process-control table and interpretation contradict the current probe output and five fresh independent reruns.
- Required outcome: correct the report and affected status wording to the reproducible no-change control result, then resubmit for source/projection review.

## Blocking Findings

### [HIGH] Durable report contains a false derived-process differential

File: `E:\Anvien\reports\Investigation\20260726_161241_p2c_projection_command_root_cause.md`

Issue: the report states that the +2-call control changed `662/2,761` processes/steps to `661/2,758` and produced three selected consumer memberships. The current artifact reports `662/2,761` and zero selected consumer/wrapper memberships in both actual and control runs.

Evidence:

- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2c\process_control_probe_output.json`: actual `3771 calls, 662 processes, 2761 steps, 0/0 memberships`; control `3773 calls, 662 processes, 2761 steps, 0/0 memberships`.
- Five fresh independent executions retained under `p2c/supervisor/process-run-1.json` through `process-run-5.json` reproduce the same no-change result on every run.
- Target graph hash remained `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090` before and after all five runs.
- The evidence and benchmark ledgers already carry the no-change values, so the durable report is internally inconsistent with its own companion record.

Why this blocks acceptance: the derived-process paragraph is part of the causal claim and currently asserts an effect that cannot be reproduced. A Supervisor PASS cannot rely on a report whose primary table contradicts the current evidence artifact.

Fix direction: replace the stale table and every dependent statement with the exact reproducible values. State only that the control adds two `CALLS` but leaves the selected default process/step set unchanged; this does not prove completeness, correctness, monotonicity, or sensitivity of process generation.

Re-review evidence required: corrected report; corrected actual-status wording where it says the control proves “no local monotonic guarantee”; one referenced current probe output; unchanged target graph hash.

## Source-Level Clearance Notes

- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2c\process_control_probe.go`: clear as a read-only target probe; it mutates only decoded in-memory graph copies and prints JSON.
- `E:\Anvien\internal\processes\processes.go`: not yet accepted as support for the stale differential. Its heuristic/cap behavior may explain why added calls do not change output, but source structure cannot substitute for the measured result.
- Other P2-C source groups (`internal/mcp`, `internal/filecontext`, `internal/lbugload`, `internal/lbugnative`) remain pending full acceptance review after the blocking report inconsistency is repaired.

## Evidence Checked

Passed:

- Five fresh process-control reruns are mutually consistent.
- Target graph hash is unchanged.
- Current benchmark/evidence ledgers use the current `662/2761` no-change result.

Failed:

- Report table and related prose do not match current evidence.
- Actual-status phrase “showing ... no local monotonic guarantee” is not established by an unchanged output; no change only establishes lack of response in this exact control.

Not run:

- Final acceptance sweep of all remaining P2-C source/projection claims; deferred until the blocking durable artifact is internally consistent.
- No production edit, build, detect-changes, or commit.

## Invariant Closure

- affected invariant: every causal statement in the durable report must be reproducible from the referenced current artifact, and measured no-change must not be promoted into an unsupported semantic property.
- sibling surfaces checked: report, raw process artifact, probe source, evidence ledger, benchmark ledger, actual-status entry, and five fresh reruns.
- residual unverified same-invariant surfaces: corrected P2-C report and final source/projection acceptance review.

## Required Fix List For Resubmission

1. Correct the report's process-control table to actual `3771/662/2761/0/0` versus control `3773/662/2761/0/0`.
2. Remove the claim that selected memberships changed or that this control proves sensitivity/non-monotonicity.
3. Align actual-status/detail prose with the bounded conclusion: no output change in this control; global process completeness remains unresolved.
4. Preserve the graph hash and current probe output, then resubmit the report for Supervisor review.

## Overall Evaluation

The command, file-detail, and Cypher representation traces may still be sound, but the submitted report cannot be accepted while its derived-process evidence is factually inconsistent. This is an evidence-integrity rejection, not a rejection of every P2-C subfinding.
