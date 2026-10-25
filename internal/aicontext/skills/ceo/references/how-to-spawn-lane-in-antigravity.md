# Antigravity Environment — Lane Spawning and Liveness Setup

> Reference for spawning and managing subagent execution lanes within the Google Antigravity environment.

---

## 1. Lane Initialization in Antigravity Environment

* **Tool to Use:** `invoke_subagent` (Native Agent Spawning tool).
* **Call Parameters (Arguments JSON):**
```json
{
  "Model": "inherit",
  "Role": "<Target Role, e.g., System Architect / Coder / QA>",
  "TypeName": "self",
  "Workspace": "inherit",
  "Prompt": "<100% Verbatim Raw Text of Contract Payload>"
}
```

* **Core Prompt Structure (Sourced from `Template-Prompt-for-Opening-a-Session.md`):**
  1. **Supreme Executive Authority Clause:** Strictly enforce orchestration authority from the CEO; forbid the subagent from initiating interactive inquiries to the Owner via `ask_question` / `request_user_input`.
  2. **Role & Mandatory Skills:** Assign exact lane role and mandatory skill paths to read before starting.
  3. **Universal Lifecycle & 6 Absolute Prohibitions:** Enforce strict adherence (no post-creation self-audits, no post-completion loops, no simulated clicks/virtual UI, only Architect/Planner may use anvien).
  4. **Workspace & Execution Discipline:** Store temporary files in `.tmp`, protect git worktree, forbid wildcard cleanups, and mandate reporting 4 transparency items before heavy commands.
  5. **Goal, Scope, Non-goals:** Explicitly define the target task, lock Touch-Map / boundaries, forbid unassigned file modifications, and prohibit reporting to CEO without explicit trigger.
  6. **CEO Session Thread ID (Return Address):** Provide the exact Thread ID of the orchestrating CEO.
  7. **Mandatory Two-Message Protocol:**
     * **Message 1 (Immediate ACK):** Send `UNDERSTOOD` immediately upon creation with goal summary, boundary, and first action.
     * **Message 2 (Completion Handoff):** Send `[VERDICT: PASS / READY_FOR_...]` with commit SHA, report path, and evidence IDs, followed by immediate HARD STOP.

* **Return Result & Acknowledgment:**
  The tool returns a spawned subagent with a dedicated conversation/thread ID. The subagent immediately delivers its mandatory initial `UNDERSTOOD` message back to the CEO Thread ID.

---

## 2. Setting Up Liveness Patrol via `schedule` Tool

* **Tool to Use:** `schedule` (Asynchronous One-Shot Timer / Heartbeat).
* **Call Parameters (Arguments JSON):**
```json
{
  "DurationSeconds": 600,
  "TimerCondition": "any",
  "Prompt": "<Verbatim MANDATORY REMINDER FOR CEO from standby-and-liveness-patrol.md>"
}
```

* **Operational Flow:**
  - Arming the timer transitions the CEO safely into asynchronous `STANDBY`.
  - Setting `TimerCondition: "any"` ensures that if the subagent completes early and delivers a return message before 600 seconds, the timer is automatically cancelled early and wakes up the CEO immediately.
