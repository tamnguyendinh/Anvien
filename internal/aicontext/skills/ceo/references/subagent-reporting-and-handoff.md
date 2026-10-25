# Subagent Reporting, Handoff & Verdict Rules

> This file is part of CEO Skill. Read when: receiving completion handoffs (READY_FOR_QA/PASS), performing management-level verification, or enforcing gate and verdict rules.

---

## 1. Lightweight Handoff Protocol of CEO (Management-Level Verification)

1. **CEO Objective Upon Receiving Completion Reports (e.g., `READY_FOR_CODER/QA/SUPERVISOR/PLANNER/ARCHITECT/<OTHER SUBAGENT>`):**
   - Verify strictly at the **Management Level**: Confirm that the subagent's message contains an explicit verdict (`READY_FOR_<SUBAGENT>`, `READY_FOR_REVIEW`, `PASS`), and that the declared report/artifact files exist on disk or in `git status`.
   - **Do not perform the specialist's job**: CEO does not need to run deep technical verification commands (such as importing code modules into the node runtime to inspect exported keys, parsing ASTs, executing full test runners, or deep-scanning internal scenario rows).

2. **Immediate Delegation to Dedicated Specialists:**
   - Upon confirming a valid handoff package, CEO immediately opens the designated specialist lane (e.g., **QA** for live runtime testing, **Supervisor** for deep zero-trust review, or **Mechanical Planner** for ticking plan items) and delegates all technical verification within that domain directly to that lane.
   - CEO focuses purely on managing progress, coordinating workflows, and advancing the state machine to the next step.

---

## 2. Task Completion & Milestone Handoff Protocol

* **(MUST - Upon Task Completion):** When a subagent (Coder, QA, Architect) **fully completes its assigned milestone task** (reaching the slice completion gate), it MUST:
  1. Record its own raw execution logs, test results, benchmarks, and Evidence IDs (`E1-P1A-...`) directly into the target `evidence.md` and `benchmark.md` files using its own write tools BEFORE sending a message.
  2. **MANDATORY: Generate an official report file** (`Reports/...`).
  3. Actively send a direct handoff message back to the Thread ID of the CEO Session (via `send_message` tool with Recipient: "<CEO_THREAD_ID>") containing:
     - The explicit completion verdict: `PASS`, `READY_FOR_<SUBAGENT>`, or `READY_FOR_REVIEW`.
     - The exact Report file path and Evidence IDs recorded on disk.
     - A concise summary of the technical outcome achieved.

---

## 3. Gate and Verdict Rules

* Spawn distinct agents (`spawn agent / create agent`) for lanes requiring user control or intervention; do not open detached top-level sessions.
* Do not open the next phase when the previous gate has not achieved a full PASS.
* If a transport failure occurs (the subagent dies or hangs without sending a message after a bounded timeout), CEO must handle it via evidence: revoke authority and spawn a replacement agent lane. Do not wait indefinitely.
* Phase Pn-C of a plan is closure/handoff docs-only; it is forbidden to open additional Supervisor loops at this slice.
