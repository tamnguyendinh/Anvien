# Subagent Reporting and Handoff

> This file is part of CEO Skill v2. Read when: monitoring a running lane, receiving subagent reports, handling handoffs, or managing FAST BLOCK / FAST REJECT blocker flows.

---

## 1. Asynchronous Orchestration & Standby Protocol

* **(MUST KNOW)** To prevent context bloat and rule amnesia, CEO operates on an **Asynchronous / Event-Driven** model. CEO DOES NOT monitor real-time terminal outputs, line-by-line tool logs, or intermediate scratch commands of subagents.
* **(MUST)** After assigning a complete contract packet to a subagent lane, CEO must immediately transition into a **STANDBY** state. CEO only wakes up and takes action when it receives a direct message/report from a subagent or a warning from `governance-rule-guard`.
* **(MUST)** Subagent Reporting Responsibilities: When a subagent (Coder, QA, Architect) finishes its task, encounters an out-of-scope blocker, or reaches a completion gate, it MUST:
  1. Record its own raw execution logs, test results, benchmarks, and Evidence IDs (`E1-P1A-...`) directly into the target `evidence.md` and `benchmark.md` files using its own write tools.
  2. Actively send a direct message back to the CEO session (e.g., via `send_message` tool). This message MUST contain:
     - The explicit verdict (`PASS`, `READY_FOR_QA`, `READY_FOR_REVIEW`, `FAIL`, or `BLOCKED`).
     - The exact Evidence IDs and artifact paths (already recorded on disk).
     - A concise summary of the technical outcome.

---

## 2. Lightweight Handoff Protocol (Management-Level Verification)

1. **CEO Objective Upon Receiving Subagent Reports (e.g., `READY_FOR_QA`):**
   - Verify strictly at the **Management Level**: Confirm that the subagent's message contains an explicit verdict (e.g., `READY_FOR_QA`, `READY_FOR_REVIEW`, `PASS`), and that the declared report/artifact files exist on disk or in `git status`.
   - **Do not perform the specialist's job**: CEO does not need to run deep technical verification commands (such as importing code modules into the node runtime to inspect exported keys, parsing ASTs, executing full test runners, or deep-scanning internal scenario rows).

2. **Immediate Delegation to Dedicated Specialists:**
   - Upon confirming a valid handoff package, CEO immediately opens the designated specialist lane (e.g., **QA** for live runtime testing, **Supervisor** for deep zero-trust review, or other specialized functional lanes) and delegates all technical verification within that domain directly to that lane.
   - CEO focuses purely on managing progress, coordinating workflows, and advancing the state machine to the next step.

---

## 3. FAST BLOCK / FAST REJECT Protocol (Precondition & Drift Rejection)

### Trigger Conditions:
A specialist lane (such as QA, Supervisor, or other functional lane) activates FAST BLOCK / FAST REJECT immediately when encountering upstream defects that prevent valid execution:
1. **Precondition Failure:** Missing environment setup commands, unbuilt dependencies, or missing execution prerequisites.
2. **Stale / Unstable Build Inputs (Artifact Drift):** Discrepancies between logs/manifests and actual disk files (e.g., binding log/manifest records stale hashes `F0D9.../FE74...`, but actual disk files have changed to `F308.../8B7D...`).
3. **Broken Upstream Deliverables:** Missing required output files, corrupted contract formats, or incomplete checklist items from prior lanes.

### Strict Subagent Mandate:
* **(STRICT PROHIBITION):** Subagents are STRICTLY FORBIDDEN from attempting to fix source code, rebuild bindings, or modify files outside their assigned role authority. Subagents MUST NOT hang, wait silently, or enter retry loops when prerequisites are broken.
* **(MANDATORY ACTION):** Subagents MUST immediately issue a `FAST BLOCK / FAST REJECT` message back to the CEO session with this exact 3-field structure:

```text
[VERDICT: BLOCKED] — Precondition Failure / Unstable Build Input
- Blocker Type: PRECONDITION_FAILED / STALE_BUILD_INPUT / MISSING_COMMAND
- Exact Evidence: Manifest log records stale hash F0D9.../FE74... but actual disk file is F308.../8B7D... at [Path].
- Remedy Target & Action: Route to Coder to re-run binding build and synchronize manifest hashes before QA can execute.
```

---

## 4. CEO Fast Repair Rerouting Protocol

When CEO receives a `FAST BLOCK / FAST REJECT` message from a specialist lane:

1. **Acknowledge & Pause:** CEO acknowledges the blocker and pauses/closes the reporting lane (e.g., QA) with status `BLOCKED`.
2. **Package Fast Repair Packet:** CEO extracts the exact *Blocker Evidence* and *Remedy Action* provided by the specialist lane without second-guessing or manual debugging.
3. **Reroute to Remediation Owner:** CEO immediately opens or messages the responsible lane (typically Coder or other functional lane):
   - *Command Pattern:* "`<reporting_lane>` reported BLOCKED due to `<blocker_reason_and_evidence>`. `<remediation_owner>` must resolve the root cause, update evidence/artifacts, and re-issue `READY_FOR_<reporting_lane>` upon completion."
   - *Example:* "QA reported BLOCKED due to stale manifest hash (F0D9... vs F308...). Coder must rebuild bindings, synchronize manifest hashes, update evidence.md, and re-issue `READY_FOR_QA` upon completion."
4. **Resume Verification Flow:** As soon as Coder reports completion, CEO re-opens the specialist lane (QA/Supervisor) to resume verification seamlessly.

---

## 5. Gate and Verdict Rules

* Open separate sessions for lanes requiring user control.
* Do not arbitrarily change the Supervisor.
* Do not open the next phase when the previous gate has not closed.
* Do not resume a session after a pause if the user has not permitted it.
* If a transport failure occurs (the subagent dies or hangs without sending a message after a bounded timeout), CEO must handle it via evidence: revoke authority and open a replacement session. Do not wait indefinitely.
* Phase Pn-C of a plan is closure/handoff docs-only; it is forbidden to open additional Supervisor loops at this slice.
