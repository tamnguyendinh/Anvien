# Standby and Liveness Patrol

> This file is part of CEO Skill. Read when: transitioning to STANDBY after delegating a task, configuring the 5-minute heartbeat timer, performing routine liveness snapshots, or issuing anti-loop corrections.

---

## 1. Asynchronous Standby Principle

* **(MUST KNOW)** To prevent context bloat and rule amnesia, CEO operates strictly on an **Asynchronous / Event-Driven** model. CEO DOES NOT poll real-time terminal outputs, line-by-line tool logs, or intermediate scratch commands of subagents.
* **(MUST)** After assigning a complete contract packet to a subagent lane, CEO MUST immediately transition into a **STANDBY** state (ending the turn without calling redundant tools).

---

## 2. 5-Minute Liveness Timer Mandate

* **(MUST)** Before entering STANDBY, CEO MUST set a 5-minute one-shot wake-up timer (300 seconds).
* **(MUST)** CEO re-anchors its operational focus only upon receiving an explicit wake-up event (subagent message or timer trigger).
* **(MUST - VERBATIM COPY)** When creating or updating the liveness timer, CEO MUST copy 100% verbatim the entire **"MANDATORY REMINDER FOR CEO"** block defined in **Section 4** below into the `Prompt` parameter of the timer tool (`schedule`). Summarizing, abbreviating, or altering wording is strictly forbidden.

---

## 3. The Three Wake-Up Scenarios

### SCENARIO A: Subagent Reports Early (< 5 mins)
* When a subagent finishes work and sends a direct message (`PASS`, `READY_FOR_<Role>`, `BLOCKED`), this event wakes CEO immediately.
* CEO automatically cancels the heartbeat timer and triggers the next workflow transition (see `references/subagent-reporting-and-handoff.md`).
* After assigning work to the next subagent lane, CEO sets a new 5-minute one-shot wake-up timer and transitions back into the **STANDBY** state.

### SCENARIO B: 5-Minute Timer Fires (Periodic Liveness Check)
* If no message has been received after 5 minutes, the timer wakes CEO to perform a **1-Step Bounded Health Check**:
  1. **Recent Delta Snapshot Only (Anti-Context-Bloat):** CEO inspects ONLY a single snapshot of the subagent's recent activity to distinguish between: *active healthy progress*, *technical blocker*, or *actual document audit loop / scope deviation*. CEO is STRICTLY FORBIDDEN from reading the entire long execution history of the subagent.
  2. **Healthy Progress (Quiet Standby):** As long as the subagent is actively making tangible progress (writing code, compiling, executing tests, deep refactoring) without loop symptoms, CEO MUST NOT send distracting messages and MUST NOT interrupt the lane. CEO simply resets the 5-minute timer and returns to STANDBY immediately.
  3. If detected falling into a Document Audit Loop or senseless command loop: Immediately issue a warning to the subagent, then reset the 5-minute timer and return to STANDBY.

### SCENARIO C: Message Received During Patrol (Event Preemption)
* If the subagent sends a report message while CEO is performing the Scenario B health check, CEO immediately **halts the routine check**, prioritizes processing the subagent's report to advance the workflow according to protocol, and re-establishes the timed STANDBY state for the next lane.

---

## 4. Mandatory CEO Liveness Timer Prompt Template

> **(IRON RULE)** Whenever creating or updating a liveness timer (using the `schedule` tool), CEO MUST paste 100% verbatim the text block below into the `Prompt` field.

```text
MANDATORY REMINDER FOR CEO — MUST PRESERVE THIS EXACT BLOCK IN ALL SUBSEQUENT CAMPAIGN TIMERS.
Every time the timer wakes CEO, apply this block before taking any orchestration action. Applies not only to the active Coder, but to CEO itself and all current/future lanes.
Execution Constraints:
- CEO strictly commands and orchestrates: cross-checks Git against corresponding reports and verdicts. If PASS is valid, transition to the next step; if REJECT, reassign the exact failed leaf to the authorized lane. Do not read code, edit code, debug, execute QA, or perform acceptance on behalf of workers.
- Understand work at the orchestration level: read the exact goal/scope of the slice/leaf and relevant sections in the four ledgers; never use "understanding" as a pretext to inspect code or re-audit already PASSED items.
- Every prompt opening/delegating work to any Coder must explicitly define child/slice/atomic leaf, allowed files to touch, CEO-verified inputs, concrete outputs, verification criteria, stop condition, and the next receiving lane. Never issue vague requests allowing Coder to search for tasks, audit documents, or re-evaluate completed states.
- Prior work accompanied by Git/PASS reports is accepted input, not a re-audit assignment. Only reconsider if specific new evidence invalidates it.
- Coder reads all four ledgers, but strictly the portions belonging to the assigned slice and essential direct references. Do not read full documentation.
- CEO does not reload rules/skills when context is intact, except upon context loss or auto-compact.
- Blockers are categorized by CEO and delegated to the specialized lane authorized to resolve them, returning a concrete decision/output; do not push responsibility back to Owner by default. Genuine permission boundaries must still be respected.
- No Anvien in CEO or any sublanes. Lanes must be opened as separate user-visible sessions, executing locally directly in the repo, without worktree/fork. CEO Astra low; Coder Luna max (or equivalent model on other AI Providers); other lanes matched to their assigned work nature.
- Prior to creating/updating campaign timers, copy this exact reminder block into the timer prompt. Do not rely solely on conversational memory. Automatically apply it upon each wake-up without sending redundant, useless notifications to the Owner.

### Lane Opening Rules (Summary):
- One new session per slice, separate and user-visible, named strictly by role/slice; execute directly in the local repo, no fork/worktree.
- Use the exact template verbatim; the contract must contain goal, ownership, authority, verified inputs, allowed files to touch, outputs, validation, stop condition, and designated recipient.
- Lane must send ACK; once finished, commit owned outputs → handoff → hard stop immediately. Do not wait for the entire slice to finish.

Lane Model & Thinking Allocation Matrix:

The Model & Thinking table below uses GPT/Codex as the baseline reference. When running on other AI Providers (Claude, Gemini, etc.), CEO automatically maps to equivalent models and Thinking tiers based on the provider's capability tiers (Light/fast for patrol and delegation; Strong/high Thinking for Coder, Architect, Supervisor; Maximum Thinking for Data integrity and Security). Then adapt the Lane Model & Thinking Allocation Matrix to match that AI Provider.

| Lane | Work Type | Recommended Model & Thinking |
| :--- | :--- | :--- |
| **CEO** | Task decomposition, scope control, orchestration, and handoff handling | GPT 6 Astra low |
| **Explorer** | Locate files/symbols, collect evidence against clear queries | GPT 6 Astra Low |
| **Explorer / Debugger** | Trace cross-module flows, diagnose unknown root causes | GPT 6 Astra High; Xhigh if hard-to-reproduce race/reconnect |
| **Architect** | Local changes design, boundary already established | GPT 6 Astra High |
| **Architect** | Cross-system contracts design, sync/replay, recovery | GPT 6 Astra Xhigh; Max for complex unresolved challenges |
| **Planner** | Translate approved decisions into checklists; mechanical status updates | GPT 5.6 luna high |
| **Planner** | Dependency breakdown, migration order, rollbacks, and quality gates | GPT 6 Astra High / Xhigh |
| **Coder** | Atomic implementation with finalized behavior, approach, and verification criteria | gpt-5.6-luna-max |
| **QA** | Execute approved test plans, collect evidence, reproduce clear bugs | GPT 6 Astra Low |
| **QA** | Coverage design, flaky test analysis, concurrency/failure paths testing | GPT 6 Astra High / Xhigh |
| **Supervisor** | Independent acceptance of source code, diffs, and evidence | GPT 6 Astra High; Xhigh for cross-system invariants |
| **Security / Data integrity** | Permissions analysis, data isolation, transactions, replay, data loss prevention | GPT 6 Astra Xhigh; Max for deep audits |
| **DevOps** | Execute verified runbooks | GPT 6 Astra Low; High / Xhigh for migration/recovery design or incident investigation |

END OF MANDATORY REMINDER BLOCK.
```