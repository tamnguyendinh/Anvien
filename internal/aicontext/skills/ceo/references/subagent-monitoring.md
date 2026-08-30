# Subagent Monitoring

> This file is part of Orchestration Skill v2. Read when: monitoring a running lane, detecting loops/stuck, handling transport failures.

## Gate and Verdict Rules

* open separate sessions for lanes requiring user control;
* wait for the verdict of that session: While the subagent session is working, the orchestration agent (CEO agent) must focus on other tasks (without affecting or overwriting the subagent session's work). If there are no tasks because it is forced to wait for the subagent session's report, it must continuously monitor and wait until there is a report/verdict from the subagent session.
* When there is a verdict from the subagent session (subagent lane), the orchestration agent (CEO agent) must update the plan progress, check off the checklist in the plan (if it is the right stage to update) with the exact required skill, then assign work to the next subagent session or close the plan if the plan has ended.
* do not arbitrarily change the Supervisor;
* do not open the next phase when the previous gate has not closed;
* do not resume a session after a pause if the user has not permitted it.

## Asynchronous Orchestration & Reporting Protocol

* **(MUST KNOW)** To prevent context bloat and rule amnesia, CEO operates on an **Asynchronous/Event-Driven** model. CEO DOES NOT monitor the real-time terminal output or step-by-step commands of subagents.
* **(MUST)** After assigning a complete contract to a subagent lane, CEO must immediately transition into a STANDBY state. CEO only wakes up and takes action when it receives a direct message/report from a subagent or a warning from the `governance-rule-guard`.
* **(MUST)** The responsibility of Subagents: When a subagent finishes its task, encounters an out-of-scope blocker, or reaches a completion gate, it MUST actively send a direct message back to the CEO session (e.g., via `send_message` tool). This message MUST contain:
  1. The explicit verdict (PASS / FAIL / BLOCKED).
  2. The exact Evidence IDs or artifact paths.
  3. A concise summary of the outcome.
* **(MUST)** CEO's response: Upon receiving the direct report, CEO wakes up, reads the durable output and evidence[cite: 23, 24], self-verifies the handoff against the acceptance criteria, and then moves the state machine to the next workflow.
* **(MUST)** If a transport failure occurs (the subagent dies or hangs without sending a message after a bounded timeout), CEO must handle it via evidence: revoke the authority and open a replacement session. Do not wait indefinitely[cite: 24].
* Phase Pn-C of a plan is closure/handoff docs-only; It is forbidden to open additional Supervisor loops at this slice.
