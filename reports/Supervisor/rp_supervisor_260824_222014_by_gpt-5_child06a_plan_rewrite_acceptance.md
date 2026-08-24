# Supervisor Report: Child 06A Plan Rewrite Acceptance

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260824_222014_by_gpt-5_child06a_plan_rewrite_acceptance.md`
- Review time: `2026-08-24 22:20:14 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`
- Scope reviewed: current Child 06A four-ledger plan set plus auxiliary `plan-rules.md`, binding optimization-method report, roadmap, direct Child 06 predecessor surfaces, and relevant Child 07 successor surfaces
- Claim reviewed: the completed rewrite is a correct, execution-ready realization of the user/Main discussion for exhaustive elapsed-time optimization with one P2-A implementation slice, per-attempt Architect/Planner/Coder/measurement/Supervisor control, one final Supervisor, and one implementation commit
- Authority used: latest direct user decisions; `AGENTS.md`; planner and supervisor skills; `reports/Investigation/rp_main_260824_204615_child06a_optimization_method_discussion.md`; current roadmap and Child 06/06A/07 ledgers; repository command surface
- Related artifacts: historical `reports/Supervisor/rp_supervisor_260824_205346_by_gpt-5_child06a_execution_ready_plan_upgrade.md` was read only as prior-scope history; it reviewed the now-superseded M0/A/A/U1-U4 plan and is not current acceptance evidence

## Executive Summary

- Problem: determine whether Planner's final Child 06A rewrite faithfully turns the agreed measure -> rank -> drill down -> optimize -> remeasure -> accuracy-check method into an executable plan.
- Decision: REJECT. The rewritten control model is substantially aligned with the discussion, but the current bytes name the wrong CLI command and leave the conditional P1 timing-instrumentation write path internally under-specified against the read-only measurement pool and repository pre-edit/build rules.
- Required outcome: correct the exact command identity and close the minimal-instrumentation ownership/build/comparability state machine, then resubmit the same plan set for review. No implementation or measurement should start from the current bytes.

## Blocking Findings

### [HIGH] The plan names a non-existent command as the real product boundary

File: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md:21`

Issue: the Goal, Scope, and P1 data-source text call the executable command ``anviens analyze``. The benchmark's product-boundary row repeats the same command. The binding method report also carries this typo. The repository command surface is ``anvien analyze``.

Evidence:

- Current plan lines 21, 38, and 209 contain ``anviens analyze``; current benchmark line 77 contains the same token.
- Binding method report lines 5, 37, 104, 226, and 478 also contain ``anviens analyze``.
- `AGENTS.md:35`, `AGENTS.md:39`, and `AGENTS.md:61-64` define the real graph command and benchmark form as ``anvien analyze`` / ``anvien analyze --benchmark-json <file>``.
- A repository-source/documentation scan excluding plans and reports found the canonical ``anvien analyze`` command throughout README, runbook, CLI source/tests, and generated command guidance, with no ``anviens analyze`` command surface.
- Independent structural assertion counted three wrong-command occurrences and zero correct-command occurrences in `plan.md`.

Why this blocks acceptance: P1's first required action is to run and bind the real command. An execution plan cannot be accepted while its controlling product boundary names the wrong executable, especially when the plan repeatedly calls that spelling exact/real.

Fix direction: replace command-form occurrences with ``anvien analyze`` across the current plan/benchmark and correct the binding method record or add an explicit authoritative correction there. Keep exact options/workload/cache identity dynamic until P1-A records the accepted invocation.

Re-review evidence required: fresh exact-token scan proving no executable ``anviens analyze`` residue remains in active authority, all current command references agree, and companion links/structure remain unchanged.

### [HIGH] Conditional P1 instrumentation has no closed writer/build/use/restore path

File: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md:203`

Issue: P1-A allows edits for minimum temporary timing instrumentation and instructs the executor to add it when existing timing is insufficient, but `plan-rules.md` defines the measurement lanes as read-only. The P1 evidence contract does not identify a single writer, the pre-edit impact proof, the build that makes the instrumented runtime valid, or the exact rebuild/remeasure transition when instrumentation is removed. This leaves two unsafe interpretations: a read-only measurement lane writes source, or timing is accepted from source changes without a defined canonical-build boundary.

Evidence:

- Plan lines 203 and 228 authorize conditional measurement-instrumentation source changes; line 242 says they must be removed or carried into the first refreshed P2-A attempt.
- `plan-rules.md:44` makes all active measurement-pool lanes read-only while only stating generally that production/test writing is sequential.
- `plan-rules.md:48` binds the canonical full build to production attempts, while P1 explicitly excludes Architect/Coder/Supervisor attempt work at plan line 194.
- Evidence line 93 defines `E1-P1A-INSTR1` only as smallest instrumentation, overhead/comparability boundary, and cleanup disposition. It does not require writer identity, exact edited owner, graph/file-detail/impact, canonical build, accepted instrumented-runtime identity, or post-removal rebuild/remeasure proof.
- `AGENTS.md:21` requires a full build before validation; `AGENTS.md:39-40` requires fresh graph plus file-detail/impact before editing covered owners.

Why this blocks acceptance: the user explicitly asked whether deeper timing may require code and agreed that measurement should drill down inside the selected bottleneck. That conditional branch must be executable without inventing a harness project, violating the read-only measurement-lane contract, or using an unbuilt/mismatched instrumentation state as the numeric baseline.

Fix direction: keep existing timing first. If one timing value is genuinely missing, make P1 explicitly assign exactly one visible sequential instrumentation writer outside the read-only measurement analysts (or explicitly transition one lane into the sole writer role); require fresh graph/file-detail/impact before the edit; require canonical full build before any resulting timing is ranked; record output equivalence and like-for-like instrumentation state/overhead; and define one exact disposition: retain it under named ownership into the refreshed P2-A attempt, or remove it, rebuild, and re-establish the accepted timing basis. Extend `E1-P1A-INSTR1` and actual-status with those pointers. This is a minimal conditional branch, not a new phase, harness, or report system.

Re-review evidence required: synchronized plan/rules/evidence/status text proving one writer, read-only measurement analysts, pre-edit impact, build-before-use, comparable timing state, and exact removal/carry transition without adding another implementation slice or commit.

## Source-Level Clearance Notes

- Child 06A `plan.md`: blocked only by the two findings above; core goal, exhaustive hierarchy, attempt loop, checklist semantics, and P3 closure topology are otherwise internally coherent.
- Child 06A `plan-rules.md`: blocked by the unresolved read-only-lane versus conditional-writer branch; the `10`-lane pool and `3`-ACTIVE cap otherwise match the direct decision.
- Child 06A `evidence.md`, `benchmark.md`, and `actual-status.md`: structurally clear and conservative; benchmark additionally repeats the wrong command token, and the E1 instrumentation contract needs the missing proof fields.
- Roadmap / Child 06 / Child 07 ledgers: clear for topology. Current bytes consistently encode `Child 06 -> Child 06A -> Child 07`, P6-D closure at `81163e39718b94a509e41114cada224e8f269e36`, and Child 07 blocked until Child 06A closure.
- Production source/tests: not applicable; this review concerns plan authority and no functional implementation was claimed.

## Evidence Checked

Passed:

- The Child 06A directory contains exactly four standard planner files plus one auxiliary `plan-rules.md`; all five expected files exist.
- All four standard ledgers link to `plan-rules.md`; ten parsed local Markdown links resolve with zero missing targets.
- P0-A is checked; P1-A, P2-A, P3-A, P3-B, and P3-C are present once in their expected incomplete state.
- Goal and execution ordering correctly make elapsed wall-clock time controlling; secondary metrics are diagnostic/comparability-only.
- P1 creates the full top-level list and matching parent checklist; P2 creates complete parent-local child lists and matching nested checklist items; smaller rows cannot be omitted.
- The per-attempt chain is present: current child/parent/total -> fresh Architect -> Planner refresh -> Coder -> full build -> child/parent/end-to-end remeasurement -> fresh Supervisor -> disposition.
- `KEEP`, `REJECT`, restoration, same-child retry, three consecutive no-KEEP attempts, child-only `SYSTEM_CHARACTERISTIC`, parent completion, `BLOCKED` behavior, final equivalence, one final Supervisor, cleanup, one detect, one implementation commit, and Child 07 handoff are synchronized across the four ledgers and rules.
- Measurement capacity is recorded as ten pre-opened visible lanes, at most three ACTIVE, one problem per lane, shared accepted capture preferred, no invented work, and sequential production/test writing.
- Forbidden active-control residue scan found no M0, fixed A/A 20/20, numeric seal, `NOT_MATERIAL`, or executable fixed U1-U4 sequence in the current five-file set. U1-U4 and cost-map mentions are provenance or explicit prohibitions only.
- Markdown placeholder, code-fence, and table-column structural checks reported zero issues.
- Scoped tracked-file `git diff --check` exited `0`; all five new Child 06A files have zero trailing-whitespace matches.
- Commit `81163e39718b94a509e41114cada224e8f269e36` exists and is the recorded P6-D closure commit.
- Verification freshness: current for the exact reviewed file hashes recorded during this review.

Failed:

- Exact command identity: failed because active authority uses ``anviens analyze`` instead of ``anvien analyze``.
- Conditional P1 instrumentation state machine: failed because the current ledgers do not close writer, impact, build-before-use, and removal/carry ownership across the read-only measurement contract.

Not run:

- No `anvien analyze`, functional benchmark, build, test, target access, source edit, lane dispatch, staging, commit, cleanup, or push. The review was plan-only; Coder and measurement execution remain held.
- No plan repair was performed. Supervisor role is verdict-only for this turn.

## Invariant Closure

- Affected invariant: Child 06A must be directly executable as an empirical elapsed-time optimization loop without a guessed command, guessed writer, hidden mutation, harness project, accuracy waiver, skipped bottleneck, extra implementation slice, or extra commit.
- Sibling surfaces checked: method report, all five Child 06A files, roadmap, four Child 06 ledgers at the transfer/closure boundary, four Child 07 ledgers at the predecessor/opening boundary, prior same-scope Supervisor history, repo command guidance, and current Git documentation diff shape.
- Residual unverified same-invariant surfaces: exact active CLI token and conditional P1 instrumentation execution path remain open as described above.

## Required Fix List For Resubmission

1. Normalize the active command boundary to ``anvien analyze`` in the method authority and all current Child 06A ledger references without hardcoding an unmeasured workload/options tuple.
2. Add one minimal conditional P1 instrumentation branch that names the sole visible writer, keeps measurement analysts read-only, requires pre-edit graph/file-detail/impact and canonical build before measurement use, proves like-for-like timing/output, and records exact removal/rebuild/remeasure or carry-forward ownership.
3. Synchronize the corresponding plan, rules, `E1-P1A-INSTR1`, benchmark comparability wording, and actual-status transition; preserve four ledgers plus one auxiliary file, one P2-A implementation slice, ten/three measurement capacity, one final Supervisor, and one implementation commit.
4. Re-run the exact structural, link, forbidden-residue, command-token, and whitespace assertions and submit the current bytes for independent re-review.

## Overall Evaluation

The rewrite successfully replaces the harness-heavy fixed-candidate workflow with the measurement-driven process the user requested. Its elapsed-time goal, exhaustive parent/child bottleneck checklists, dynamic per-attempt architecture and correctness gates, bounded measurement parallelism, and terminal rules are materially improved and mostly correct. Acceptance is nevertheless blocked because the first command is misspelled and the only conditional source-writing branch before P2 is not closed well enough for another executor to follow safely. These are narrow fixes; the empirical method and phase topology should be preserved.
