# Codex Environment — Lane Spawning and Dispatch

> Reference for spawning and managing subagent execution lanes within the Codex App environment.

---

## 1. Mandatory Tool Selection

* **Tool Name:** `mcp__codex_app__create_thread`
* **Supreme Rule:** Strictly FORBIDDEN to use `spawn_agent` (never spawn hidden, invisible agents). CEO MUST use `mcp__codex_app__create_thread` so that the newly created lane appears publicly on the Codex App Sidebar. This ensures the Owner can directly click into the lane to inspect, question, or intervene at any time.

---

## 2. Tool Parameter Structure (Arguments JSON)

The tool invocation must supply all 5 required fields with standardized structure:

```json
{
  "title": "<Role — Specific Task, e.g., Coder — UI Implementation>",
  "model": "<model_name_required_for_the_lane>",
  "thinking": "high",
  "target": {
    "type": "project",
    "projectId": "<repo_project_id>",
    "environment": {
      "type": "local"
    }
  },
  "prompt": "<100% VERBATIM RAW TEXT OF CONTRACT PAYLOAD>"
}
```

> [!IMPORTANT]
> **Key Requirement for the `target` Parameter:** Must provide all three fields: `type: "project"`, the repository's `projectId`, and `environment: { "type": "local" }`. Omitting any of these will cause Codex to fail project resolution with an `invalid arguments` error.

---

## 3. Prompt Composition Specifications (Verbatim Contract)

The prompt content must be copied 100% verbatim from `references/Template-Prompt-for-Opening-a-Session.md` without summarizing, shortening, or omitting any section:

1. **Executive Authority & Legal Preamble:** Confirm the session is initiated by the CEO under supreme authorization from the Owner; strictly forbid the lane from invoking `request_user_input`.
2. **Role & Mandatory Skills:** Designate exact operational role and mandatory skills to read before starting.
3. **6 Absolute Prohibitions:** Strictly forbid document audit loops, unauthorized subagent spawning, simulated UI actions, unapproved anvien usage, etc.
4. **Workspace & Execution Discipline:** Mandate `.tmp` usage for scratch artifacts, protect git worktree, forbid wildcard cleanups, and require explicit declaration of 4 transparency items prior to long-running commands.
5. **Goal & CEO Thread ID:** Specific objective of the slice and return address (CEO Thread ID).
6. **Scope (Touch-Map):** Hard-lock the list of allowed files to modify.
7. **Non-goals:** Explicitly list forbidden out-of-scope tasks.
8. **Mandatory Evidence:** Specific tangible evidence required for verification.
9. **Mandatory Two-Message Protocol:**
   * **Message 1:** Immediately return ACK `UNDERSTOOD` with goal summary, boundary, and first action.
   * **Message 2:** Upon completion, testing, and clean git commit, deliver `[VERDICT: PASS / READY_FOR_...]` and HARD STOP immediately.

---

## 4. CEO Actions Immediately Following Successful Thread Creation

1. **Extract `threadId`** from the tool's return payload.
2. **Emit UI Directive** on its own distinct line in the message to the Owner:
   ```text
   ::created-thread{threadId="<THREAD_ID>"}
   ```
3. **Send Brief Status Update** to the Owner detailing the opened lane and its assigned scope.
4. **Transition Immediately to Asynchronous STANDBY:** Strictly forbid polling terminal/logs or running exploratory commands to prevent context bloat; patiently wait for the lane's return handoff message.
