# Recovery and Documentation

> This file is part of Orchestration Skill v2. Read when: auto-compact occurs, session rotation (120-min), or CEO falls into documentation audit loop.

## Orchestration session rotation rules:

* (MUST) The 120-minute rotation deadline applies ONLY to the Orchestration agent (CEO agent).
* (MUST) After exactly 120 minutes of operation, the current orchestration session must create a handoff report to transfer authority to a new visible orchestration session.
* (MUST) Initialize the new visible CEO Orchestration successor session 15 minutes prior to the rotation deadline. When the deadline is reached, execute the handoff to this successor session.
* (MUST) Upon handoff, the outgoing orchestration session must clearly document the following information for the new session: the current state of work, overall progress, and any active subagent lanes.
* (MUST) As soon as the new session becomes active, the old session must immediately terminate.
* (MUST) The new session continues to strictly adhere to this 120-minute handoff cycle.

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

## 6. Rules After Auto-Compact

After an auto-compact or loss of context, the session must:

1. Re-read AGENTS.md, the skill Orchestration.md file, and the skill Working-rules.md.
2. Re-read "Rules for opening separate task sessions for subagents".
3. Re-read the currently applied SKILL.md.
4. Re-read the current authority and plan slice.
5. Check the latest durable report, ledger, and checkpoint.
6. Continue from the first uncompleted gate.

The session must not:

* restart the entire review;
* rerun a PASSED gate without a reason that the evidence was invalidated;
* turn re-anchoring into a new audit loop;
* forget gates that have been recorded in durable evidence.

Re-anchoring is for restoring context, not for resetting progress.
