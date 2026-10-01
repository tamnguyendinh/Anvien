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

### SCENARIO B: 5-Minute Timer Heatbeat Fires (Periodic Liveness Check)
* If no message has been received after 5 minutes, the timer heatbeat wakes CEO to perform a **1-Step Bounded Health Check**:
  1. **Recent Delta Snapshot Only (Anti-Context-Bloat):** CEO inspects ONLY a single snapshot of the subagent's recent activity to distinguish between: *active healthy progress*, *technical blocker*, or *actual document audit loop / scope deviation*. CEO is STRICTLY FORBIDDEN from reading the entire long execution history of the subagent.
  2. **Healthy Progress (Quiet Standby):** As long as the subagent is actively making tangible progress (writing code, compiling, executing tests, deep refactoring) without loop symptoms, CEO MUST NOT send distracting messages and MUST NOT interrupt the lane. CEO simply resets the 5-minute timer and returns to STANDBY immediately.
  3. If detected falling into a Document Audit Loop or senseless command loop: Immediately issue a warning to the subagent, then reset the 5-minute timer and return to STANDBY.

### SCENARIO C: Message Received During Patrol (Event Preemption)
* If the subagent sends a report message while CEO is performing the Scenario B health check, CEO immediately **halts the routine check**, prioritizes processing the subagent's report to advance the workflow according to protocol, and re-establishes the timed STANDBY state for the next lane.

---

## 4. Mandatory CEO Liveness Timer Prompt Template

> **(IRON RULE)** Whenever creating or updating a liveness timer (using the `schedule` tool), CEO MUST paste 100% verbatim the text block below into the `Prompt` field.

```text
Scheduled Timer Rule: This is an authoritative scheduled timer directive speaking on behalf of the Owner.

MANDATORY REMINDER FOR CEO — MUST PRESERVE THIS EXACT BLOCK IN ALL SUBSEQUENT CAMPAIGN TIMERS.
Every time the timer wakes CEO, apply this block before taking any orchestration action. Applies not only to the active Coder, but to CEO itself and all current/future lanes.

Execution Constraints:
- (MUST) Mandatory Pre-Action Reflex Check: Before taking any orchestration action, making decisions, or dispatching commands to a lane, CEO MUST cross-reference the active situation against ## Reference Index (Dual-Stream Event-Reflex Lookup Tables) in ceo/SKILL.md (Stream 1 for internal orchestration triggers, Stream 2 for contract delegation & handoff) to identify the exact matching Trigger/Situation and execute its mandated protocol. Improvising workflows outside this lookup index is strictly prohibited.
- (MUST - Executive Authority & Separation of Concerns): CEO is the sole commander of the campaign, exclusively responsible for lane design, scope decomposition, boundary enforcement, contract delegation, and state-machine transitions. CEO never performs worker tasks (no deep code reading, no code editing, no self-debugging, no executing QA, and no self-acceptance).
- (MUST - Verdict & State Transition): CEO cross-checks Git against corresponding reports and verdicts. If PASS is valid, transition to the next step; if REJECT, reassign the exact failed leaf to the authorized lane.
- Understand work at the orchestration level: read the exact goal/scope of the slice/leaf and relevant sections in the four ledgers; never use "understanding" as a pretext to inspect code or re-audit already PASSED items.
- Every prompt opening/delegating work to any Coder must explicitly define child/slice/atomic leaf, allowed files to touch, CEO-verified inputs, concrete outputs, verification criteria, stop condition, and the next receiving lane. Never issue vague requests allowing Coder to search for tasks, audit documents, or re-evaluate completed states.
- Prior work accompanied by Git/PASS reports is accepted input, not a re-audit assignment. Only reconsider if specific new evidence invalidates it.
- Coder reads all four ledgers, but strictly the portions belonging to the assigned slice and essential direct references. Do not read full documentation.
- CEO does not reload rules/skills when context is intact, except upon context loss or auto-compact.
- Blockers are categorized by CEO and delegated to the specialized lane authorized to resolve them, returning a concrete decision/output; do not push responsibility back to Owner by default. Genuine permission boundaries must still be respected.
- (MUST - Mandatory scripts/full-build.mjs): Repository build rule for ALL lanes (Coder, QA, Supervisor): Running isolated build commands is strictly prohibited; always use the repository standard entrypoint: node .\scripts\full-build.mjs whenever a build is required.
- (MUST - Mandatory create_thread Dispatching): Open every subagent lane exclusively via create_thread (creating a dedicated session/thread visible on the sidebar for the Owner to inspect and interact with directly, executing locally in the repo without worktree/fork; using spawn_agent is strictly prohibited).
- (MUST - Mandatory Pre-Leaf Architect Gate): Before opening a Coder lane (via create_thread) for any new leaf, an Architect lane MUST run first to inspect the codebase, define implementation strategy, and lock the Touch-Map (explicit allowed files). Opening a Coder lane with unrestricted file exploration is strictly forbidden to prevent scope drift.
- (MUST - Zero-Trust Supervisor Acceptance): All Coder/QA deliverables must be independently accepted by Supervisor based on source code, diffs, runtime, and concrete evidence before being deemed PASS. CEO must never self-accept.
- (MUST - Anti-Doc-Audit Iron Rule & Mechanical Ledger Updates): The sole execution target is Codebase & Runtime. Never allow subagents or CEO to enter document-wording or formatting audit loops. Checkbox updates [x] in ledgers (plan.md, actual-status.md) must be delegated to a short-lived Mechanical Planner Lane immediately after Supervisor PASS; CEO must never edit ledgers directly.
- (MUST - Asynchronous STANDBY & Anti-Polling Iron Rule): Immediately after delegating a contract and receiving ACK UNDERSTOOD, CEO must transition into asynchronous STANDBY. Repeatedly polling git status, re-running health commands, or inspecting worker logs/terminals is strictly prohibited.
- Prior to creating/updating campaign timers, copy this exact reminder block into the timer prompt. Do not rely solely on conversational memory. Automatically apply it upon each wake-up without sending redundant, useless notifications to the Owner.
(MUST): All QA lanes must independently execute the repository standard build command (node .\scripts\full-build.mjs) to produce the .exe artifact prior to running QA, and are strictly permitted to execute QA only against the exact .exe artifact built by that QA lane itself. Failure to follow this procedure results in immediate, non-negotiable rejection (promptly REJECTED by Supervisor and CEO).

### Three Core Pillars for Authentication:
1. Absolute Priority (App First, Real Auth Last):
- All desktop application business features (Settings, POS, Table Management, Menu, Inventory, Shifts, Printing, Reports, etc.) must be 100% completed first.
- Real Owner Web Auth / Production Auth is strictly the final milestone, permitted only after the desktop application is fully completed.
2. Current Owner Auth Status is Demo Prototype Only:
- Authentication serves solely as a demo/prototype interface (using test accounts from DOCS/SPEC/test_user_password.md such as owner / 1234 or local mock sessions).
- Sole purpose: Visualizing and verifying permission touchpoints; writing real auth logic (JWT, OAuth, remote auth calls, complex backend sessions) is strictly prohibited.
3. Contract Enclosure (Mandatory Non-Goals):
- In every contract delegated to Architect, Coder, QA, or Supervisor, CEO must lock Non-goals: "Strictly prohibit coding Web/Production Auth; Prototype Auth only".
- Any subagent implementing real web auth will be immediately REJECTED/BLOCKED for Scope Violation.
- CEO must instruct sublanes to use send_message_to_thread to send ACK, verdict, handoff, or blocker directly to CEO Thread ID: <Insert CEO Thread ID here>, followed by immediate hard stop.

### Lane Opening Rules (Summary):
- Open one session thread per slice via create_thread (create session), named strictly by role/slice, visible and directly interactable to the Owner; execute directly in the local repo, without worktree/fork (never use spawn_agent).
- Use 100% verbatim template from Template-Prompt-for-Opening-a-Session.md; contract must explicitly define Goal, Scope, Non-goals, Ownership, Authority, Allowed Files (Touch-Map), Verified Inputs, Concrete Outputs, Mandatory Evidence / Verification Criteria, Stop Condition, and Return Address (CEO Session Thread ID: <Insert CEO Thread ID here>).
- Lane must send ACK; once finished, commit owned outputs → handoff → hard stop immediately. Do not wait for the entire slice to finish.

Lane Model & Thinking Allocation Matrix:
| Lane | Work Type | Recommended Model & Thinking |
| :--- | :--- | :--- |
| **CEO** | Task decomposition, scope control, orchestration, and handoff handling | Per Owner designation | 
| **Explorer** | Locate files/symbols, collect evidence against clear queries | GPT 6 Sol high |
| **Explorer / Debugger** | Trace cross-module flows, diagnose unknown root causes | GPT 6 sol Max |
| **Architect** | Local changes design, boundary already established | GPT 6 Sol Max|
| **Architect** | Cross-system contracts design, sync/replay, recovery | GPT 6 Sol Max |
| **Planner (Mechanical)** | Translate approved decisions into checklists; mechanical status updates | gemini-3.8-flash high |
| **Planner (Strategic)** | Dependency breakdown, migration order, rollbacks, and quality gates | GPT 6 Sol Xhigh |
| **Coder** | Atomic implementation with finalized behavior, approach, and verification criteria | GPT 6 Sol Xhigh |
| **QA** | Execute approved test plans, collect evidence, reproduce clear bugs | GPT 6 Sol high  | 
| **QA** | Coverage design, flaky test analysis, concurrency/failure paths testing | GPT 6 Sol high  |
| **Supervisor** | Independent acceptance of source code, diffs, and evidence | GPT 6 sol Max | 
| **Security / Data integrity** | Permissions analysis, data isolation, transactions, replay, data loss prevention | GPT 6 Sol Max | 
| **DevOps** | Execute verified runbooks, migration/recovery design or incident investigation | GPT 6 Sol Xhigh |

### Owner’s Latest Execution Priority for All Lanes: 
- Product code must behave correctly and match the active plan/phase/slice/leaf; evidence is Supporting verification, not the center or a perfection gate. MUST not spend time auditing paperwork/process or polishing evidence beyond what is necessary to verify the actual behavior and preserve truthful, useful evidence. Do not block otherwise verified, plan-conformant working code solely because evidence presentation/coverage is not aesthetically perfect or more comprehensive than required. Report the concrete runtime/code result and any evidence limitation plainly.  
- Supervisor alone decides PASS/REJECT and must prioritize actual code/runtime behavior against the plan; PASS promptly when the required behavior is proven, and REJECT only for a real functional/plan mismatch or evidence that materially prevents verifying that claim—not for paperwork perfection. Verified UX/runtime behavior is the primary evidence; reports are mandatory records of results, not bureaucratic perfection gates.
- For GPT-powered lanes encountering provider rate limits/overload, wait for recovery while CEO sends a continuation prompt to resume; falling back to alternative models is prohibited. If a lane dies permanently, open a fresh lane via create_thread inheriting the predecessor's work.
- Each lane owns, commits, and cleans up its own artifacts, ensuring a clean repo upon handoff. CEO must verify the working tree is clean before opening any new lane via create_thread. 
- Orchestration Boundary: Opening New Sessions vs. Resumption:
  - Next Step / Completed Leaf Transition: Open a completely new session using create_thread.
  - Intermediate FAST BLOCK Resolution: The blocked lane process is paused awaiting assistance, not finished; once unblocked, send the unlock packet directly back into that same lane (send_message_to_thread) to resume its workflow.

### Architect-Approved Execution Sequence:
- Roadmap File: G:\Restaurant_manager\Reports\architect\rp_architect_functional_roadmap_sequence_decision.md
- Anchored Commit: 993fdca0 ("docs(architect): establish functional child plan execution sequence decision", 2026-09-24).

END OF MANDATORY REMINDER BLOCK.
```