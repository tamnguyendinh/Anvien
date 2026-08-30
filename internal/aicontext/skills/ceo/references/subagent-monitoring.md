# Subagent Monitoring

> This file is part of Orchestration Skill v2. Read when: monitoring a running lane, detecting loops/stuck, handling transport failures.

## Gate and Verdict Rules

* open separate sessions for lanes requiring user control;
* wait for the verdict of that session: While the subagent session is working, the orchestration agent (main agent) must focus on other tasks (without affecting or overwriting the subagent session's work). If there are no tasks because it is forced to wait for the subagent session's report, it must continuously monitor and wait until there is a report/verdict from the subagent session.
* When there is a verdict from the subagent session (subagent lane), the orchestration agent (main agent) must update the plan progress, check off the checklist in the plan (if it is the right stage to update) with the exact required skill, then assign work to the next subagent session or close the plan if the plan has ended.
* do not arbitrarily change the Supervisor;
* do not open the next phase when the previous gate has not closed;
* do not resume a session after a pause if the user has not permitted it.

## Monitoring subagent sessions:

* orchestration agent (main agent) must monitor the actual behavior of subagents until a conclusion or verdict is reached, not just listen to its reports;
* a. The Owner can directly intervene in the Subagents task, but the main session's responsibility remains to stay, continuously monitor, receive durable reports/verdicts, self-verify the handoff, and then continue the process/plan.
* b. When monitoring a subagent session: If the subagent deviates from the goal or falls into an infinite loop, the main agent must issue a reminder into the subagent session so the subagent returns to the exact original goal.
* c. The orchestration agent (main agent) is tasked with updating the status for the plan's next phase/slice or updating the codebase's latest status for the next plan (if it is a multi-plan), then assigning work to the subagent session to execute the next phase/slice.
* d. (MUST) Main must monitor the actual behavior of the lane, not just wait for the final report. Main must actively observe what commands the lane runs, which files it reads, and whether it is expanding the scope or getting caught in a loop.
* e. (MUST NOT) A gate that has already achieved a PASS status MUST NOT be re-run simply because of a change in Main, a compaction event, or a feeling of "uneasiness". For example, if a phase/slice or issue has already generated evidence (e.g., profile, graph refresh, pprof), that evidence must be consumed directly. Rerunning wastes time and creates artificial variance that derails the campaign.
* f. (MUST NOT) Measurement truth must be preserved. Attribution must not be turned into a production attempt, streak, or speedup claim.
* g. (MUST) Transport failures must be handled via evidence. If there is no ACK/tool-response/write after a bounded wait, Main must revoke the authority and open exactly one replacement session. Do not wait indefinitely, and do not allow two functional owners to exist simultaneously.
* Phase Pn-C of a plan is closure/handoff docs-only; It is forbidden to open additional Supervisor loops at this slice.
