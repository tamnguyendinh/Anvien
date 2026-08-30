# Recovery and Documentation

> This file is part of CEO Skill. Read when: auto-compact occurs or CEO falls into documentation audit loop.

## Documentation Principles for Orchestration:

1. Documentation is merely a ledger reflecting actual progress; it is not a standalone phase, slice, gate, or work product.
2. When documentation is missing or its state is out of sync, correct it in the exact place once and proceed with the work. Prohibited loop: read → audit → fix → re-verify → write report → re-audit just to prove the documentation is correct.

 > **Keep it simple:** Update the documentation/plan as quickly as possible so that the work progress always moves forward.

3. Documentation updates (`plan.md`/`actual-status.md`) must belong to the currently open slice and be executed by delegating a strict mechanical contract to a short-lived Planner Lane after Supervisor PASS; CEO does not directly edit markdown files to preserve its context window.
4. Do not create durable reports, Supervisor loops, or evidence gates solely to prove that a few lines of documentation have been updated.
5. Once a slice has achieved a Supervisor PASS, is committed, the next slice in the plan opens automatically. Do not re-audit to check if it "is allowed to be opened yet."
6. Evidence is a step within a Phase/slice; it must not be turned into an intermediate gate.
7. Do not halt implementation to wait for "planner authorization" after an impact if the scope remains within the opened slice. Only pause when evidence proves there is an actual change to the boundary/contract.
8. Do not re-run a PASSed gate when the associated source, evidence, and boundary have not been invalidated. Re-anchoring after compaction is strictly for context recovery, not for restarting an audit on work that has already passed.
9. Do not continuously cross-check hashes, HEAD, and wording just because the ledger was recently updated. Only check the Git boundary when it genuinely serves handoff, Supervisor, or commit operations.
10. The responsibility of CEO is to understand the plan, delegate tasks, monitor commands/diffs/scopes, prevent deviations, and drive user or plan requirements toward results—do not turn yourself or the workers into documentation auditors (except for phases/slices dedicated specifically to documentation auditing).
11. (MUST) Simple tasks must remain simple: Assign Planner Lane to apply the exact mechanical change at the specified location, perform minimal checks, and immediately close. It is strictly forbidden for Planner Lane to turn a simple checklist tick into an exhaustive full-file audit, proofreading session, or wording debate.
12. (MUST) Data Entry Delegation: Functional subagents (Coder, QA, Architect) are strictly responsible for using their own tools to record raw execution logs, test outputs, and evidence IDs (`E1-P1A-...`) directly into `evidence.md` and `benchmark.md` BEFORE reporting PASS to CEO. CEO MUST NOT delegate Planner Lane to act as a log-copying secretary for functional workers.

> **In short:** Documentation must follow actual progress; progress must not get stuck chasing documentation.

## Rules After Auto-Compact

After an auto-compact or loss of context, the CEO session must: 

* Re-read raw AGENTS.md, the CEO skill, and the skill Working-rules.md.
* Re-read the active `<plan>.md` (or `<child-plan>.md` + `<campaign-roadmap>.md` if operating within a multi-plan campaign) solely to extract the active slice objectives and done criteria for subagent task contracts.

**The CEO session must not**:

* MUST NOT restart reviews or rerun gates that have already achieved a durable PASS.
* MUST NOT turn re-anchoring into a new documentation audit loop.
* restart the entire review.
* rerun a PASSED gate without a reason that the evidence was invalidated.
* turn re-anchoring into a new audit loop; Re-anchoring is for restoring context, not for resetting progress.
* forget gates that have been recorded in durable evidence.

**Execution Continuity:**
* Check the current Git state / latest durable report to identify the active slice.
* Continue execution immediately from the first uncompleted gate.
