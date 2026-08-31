# Subagent Reporting, Handoff & Emergency Blocker

> This file is part of CEO Skill. Read when: receiving completion handoffs (READY_FOR_QA/PASS), performing management-level verification, or handling in-flight emergency blockers (FAST BLOCK / FAST REJECT).

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
  3. Actively send a direct handoff message back to the CEO session (e.g., via `send_message` tool) containing:
     - The explicit completion verdict: `PASS`, `READY_FOR_<SUBAGENT>`, or `READY_FOR_REVIEW`.
     - The exact Report file path and Evidence IDs recorded on disk.
     - A concise summary of the technical outcome achieved.

---

## 3. FAST BLOCK / FAST REJECT Protocol (Emergency Blocker & Reinforcement Request)

* **Nature:** This is **NOT a task completion handoff**, but an **in-flight emergency alert** indicating that a subagent is blocked by unexpected external/precondition defects and requires CEO to **dispatch an emergency reinforcement lane** to resolve the impediment.
* (MUST KNOW) If a FAIL/UNAVAILABLE/NOT_RUN result does not cripple the subagent lane's ability to execute other rows, it is strictly classified as a "finding". The subagent lane must record this finding, step over it, and continue execution. A FAST BLOCK must ONLY be used when a true blocker prevents the entire lane from continuing safely, or completely invalidates the candidate/evidence.

### Trigger Conditions:
A specialist lane (such as QA, Supervisor, or other functional lane) activates FAST BLOCK / FAST REJECT immediately upon encountering external defects:
1. **Precondition Failure:** Missing environment setup commands, unbuilt dependencies, or missing execution prerequisites.
2. **Stale / Unstable Build Inputs (Artifact Drift):** Discrepancies between logs/manifests and actual disk files (e.g., binding log/manifest records stale hashes `F0D9.../FE74...`, but actual disk files have changed to `F308.../8B7D...`).
3. **Broken Upstream Deliverables:** Missing required output files, corrupted contract formats, or incomplete checklist items from prior lanes.

### Strict Subagent Mandate Upon Blocker:
* **(STRICT PROHIBITION):** Subagents are STRICTLY FORBIDDEN from attempting to fix code outside their assigned role authority. Subagents MUST NOT hang, wait silently, or enter retry loops. **ABSOLUTELY FORBIDDEN FROM WRITING FULL REPORT FILES** (since work is not finished, generating lengthy reports wastes critical time and stalls execution).
* **(MANDATORY - Zero-Lag):** Subagents MUST immediately issue an emergency direct message back to the CEO session with this exact 3-field structure:

```text
[VERDICT: BLOCKED or REJECT] — In-Flight Blocker / Emergency Reinforcement Request
- Blocker Type: <BLOCK_TYPE> or <REJECT_TYPE>
- Exact Evidence: <blocker_or_rejection_details_and_evidence>
- For a FAST BLOCK / FAST REJECT: record the evidence only, do not write a report.
- Remedy Target & Action: Request CEO to dispatch subagent reinforcement to resolve the issue so execution can resume.
```

---

## 4. CEO Fast Repair Rerouting Protocol

When CEO receives a `FAST BLOCK / FAST REJECT` message from a specialist lane:

1. **Acknowledge & Pause:** CEO acknowledges the blocker and pauses the reporting lane (e.g., QA) with status `BLOCKED` (preserving its working state).
2. **Package Fast Repair Packet:** CEO extracts the exact *Blocker Evidence* and *Remedy Target & Action* provided by the specialist lane without second-guessing or manual debugging.
3. **Reroute to Remediation Owner:** CEO immediately opens or messages the responsible lane (typically Coder or DevOps):
   - *Command Pattern:* "`<reporting_lane>` reported BLOCKED due to `<blocker_reason_and_evidence>`. `<remediation_owner>` must resolve the root cause, update evidence/artifacts, and re-issue `READY_FOR_<reporting_lane>` upon completion."
   - *Example:* "QA reported BLOCKED due to stale manifest hash (F0D9... vs F308...). Coder must rebuild bindings, synchronize manifest hashes, update evidence.md, and re-issue `READY_FOR_QA` upon completion."
4. **Resume Operational Workflow:** As soon as the recovery lane reports success (PASS), the CEO must immediately reactivate the original specialized lane (the session that issued the FAST BLOCK / FAST REJECT) so the Subagent session can seamlessly resume its interrupted work.

---

## 5. Gate and Verdict Rules

* Open separate sessions for lanes requiring user control or intervention.
* Do not open the next phase when the previous gate has not achieved a full PASS.
* If a transport failure occurs (the subagent dies or hangs without sending a message after a bounded timeout), CEO must handle it via evidence: revoke authority and open a replacement session. Do not wait indefinitely.
* Phase Pn-C of a plan is closure/handoff docs-only; it is forbidden to open additional Supervisor loops at this slice.
