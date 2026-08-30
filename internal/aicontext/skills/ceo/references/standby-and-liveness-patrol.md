# Standby and Liveness Patrol

> This file is part of CEO Skill v2. Read when: transitioning to STANDBY after delegating a task, configuring the 10-minute heartbeat timer, performing routine liveness snapshots, or issuing anti-loop corrections.

---

## 1. Asynchronous Standby Principle

* **(MUST KNOW)** To prevent context bloat and rule amnesia, CEO operates strictly on an **Asynchronous / Event-Driven** model. CEO DOES NOT poll real-time terminal outputs, line-by-line tool logs, or intermediate scratch commands of subagents.
* **(MUST)** After assigning a complete contract packet to a subagent lane, CEO MUST immediately transition into a **STANDBY** state.

---

## 2. 10-Minute Liveness Timer Mandate

* **(MUST)** Before entering STANDBY, CEO MUST set a 10-minute one-shot wake-up timer (via the `schedule` tool with `DurationSeconds: 600`).
* CEO re-anchors its operational focus only upon receiving an explicit wake-up event (subagent message or timer trigger).

---

## 3. The Three Wake-Up Scenarios

### SCENARIO A: Subagent Reports Early (< 10 mins)
* When a subagent finishes and sends a direct message (`PASS`, `READY_FOR_QA`, `BLOCKED`), this event wakes CEO immediately.
* CEO automatically cancels the heartbeat timer and triggers the next workflow transition (see `references/subagent-reporting-and-handoff.md`).

### SCENARIO B: 10-Minute Timer Fires (Periodic Liveness Check)
* If no message has been received after 10 minutes, the timer wakes CEO to perform a **1-Step Bounded Health Check**:
  1. **Recent Delta Snapshot Only (Anti-Context-Bloat):** CEO inspects ONLY a single snapshot of the subagent's recent activity to distinguish between: *active healthy progress*, *technical blocker*, or *actual document audit loop / scope deviation*. CEO is STRICTLY FORBIDDEN from reading the entire long execution history of the subagent.
  2. **Healthy Progress (Quiet Standby):** As long as the subagent is actively making tangible progress (writing code, compiling, executing tests, deep refactoring) without loop symptoms, CEO MUST NOT send distracting messages and MUST NOT interrupt the task. CEO simply resets the 10-minute timer and returns to STANDBY immediately.
  3. If detected falling into a Document Audit Loop / command loop, immediately issue a warning to the subagent, after which CEO resets the 10-minute timer and returns to STANDBY.

### SCENARIO C: Message Received During Patrol (Event Preemption)
* If the subagent sends a report message while CEO is performing the Scenario B health check, CEO immediately **halts the routine check**, prioritizes processing the subagent's report to advance the workflow according to protocol, and re-establishes the timed STANDBY state for the next lane.
