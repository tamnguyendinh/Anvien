# CEO Rotation Protocol

> This file is part of CEO Skill. Read when: rotating or handing over authority to a successor CEO session.

**Mandatory Convention:** Every sending action in this protocol MUST be explicitly executed as **sending a direct message to the exact session ID of the recipient CEO**. Recording content only in reports, files, or another session is NOT considered sending a message.

---

## Succession & Handover Workflow

### 1. State Lock

**Predecessor CEO:**
- Identify pending tasks and active handoffs.
- Record active lanes, verdicts, reports, commits, HEAD, worktree state, and execution cursors.
- Retain operating authority; do not permit campaign gate transitions.
- After the Candidate CEO session is spawned, the Predecessor CEO sends a direct message to the exact session ID of the Candidate CEO containing this state snapshot.

### 2. Durable Handoff Dossier

**Predecessor CEO:**
- Record the history and progress of each plan/child-plan.
- Explain the root cause of campaign state transitions.
- Record current tasks, lane registry, and designated next owners.
- Document branch paths for PASS, REJECT, and BLOCKED outcomes.
- Record Git state, evidence IDs, workspace boundaries, and next objectives.
- Create an independent Git commit for this handoff report.
- Send a direct message to the exact session ID of the Candidate CEO containing the file path and commit hash of the handoff report.

*The handoff report must never replace or summarize raw rules and skills.*

### 3. Spawn Candidate CEO in `PRE-TRANSFER` State

**Predecessor CEO:**
- Create a fresh task session (do not fork).
- Set state to `PRE-TRANSFER`.
- Assign raw authority sources and the handoff dossier via direct message to the exact session ID of the Candidate CEO.

**Candidate CEO:**
- Holds zero operational authority.
- Strictly forbidden from opening lanes, setting heartbeat timers, modifying files, or processing verdicts.
- Send a direct message with `PRE_TRANSFER_ACK` to the exact session ID of the Predecessor CEO.

### 4. Raw-Rule Gate

**Candidate CEO:**
- Read all raw rules and skills directly to EOF.
- Measure and verify path, bytes, line count, and SHA-256.
- Answer challenge queries using exact raw path, heading, and verbatim clause.
- Strictly forbidden from summarizing or paraphrasing.

**Predecessor CEO:**
- Issue challenge and evaluate response.
- Any paraphrase, approximation, or use of outdated rules results in immediate FAIL.
- Send the challenge query via direct message to the exact session ID of the Candidate CEO; Candidate CEO sends the evaluation result via direct message back to the exact session ID of the Predecessor CEO.

### 5. Campaign-Knowledge Gate

**Candidate CEO proves:**
- Current campaign status and root causes of state transitions.
- Which child plans are completed, invalidated, or locked.
- Exact execution, remediation, and frontier cursors.
- Active lanes, pending handoffs, and designated next owners.
- Concrete action branches for PASS, REJECT, and BLOCKED.
- Distinction between current HEAD and product release candidates.

**Predecessor CEO:**
- Cross-check and evaluate gate compliance.
- Candidate CEO sends the Campaign-Knowledge Gate result via direct message to the exact session ID of the Predecessor CEO.

### 6. Behavioral Dry-Run Gate

**Candidate CEO:**
- Load the latest raw template prompt.
- Draft the exact contract packet for the next task without sending.
- Accurately define role, scope, non-goals, and mandatory evidence.
- Demonstrate understanding of the 5-minute heartbeat cycle, ending turns cleanly, and strict prohibition of manual polling.

**Predecessor CEO:**
- Inspect drafted contract and planned runtime behavior.
- Candidate CEO sends the Behavioral Dry-Run packet via direct message to the exact session ID of the Predecessor CEO.

### 7. Official Authority Transfer

**Predecessor CEO:**
- Send `OFFICIAL AUTHORITY TRANSFER` via direct message to the exact session ID of the Successor CEO only after all three gates achieve PASS.
- Relinquish campaign orchestration and transition strictly into shadow observation mode.

**Successor CEO:**
- Send `AUTHORITY_ACCEPTED` via direct message to the exact session ID of the Predecessor CEO.
- Officially assume control over lanes, execution cursors, and campaign orchestration authority.

### 8. First-Cycle Observation & Probation

**Successor CEO:**
1. Receive and process the pending handoff first (do not bypass directly to remediation lanes).
2. Validate existing verdict, report, and commit.
3. Select the valid state machine transition.
4. Open exactly one next lane.
5. Arm the 5-minute heartbeat timer.
6. End turn and enter actual STANDBY.
7. Send direct message to the exact session ID of the Predecessor CEO with proof of first-cycle execution.

**Predecessor CEO:**
- Monitor and evaluate the first operational cycle.
- **PASS:** Send `TRANSFER COMPLETE` via direct message to the exact session ID of the Successor CEO and terminate session completely.
- **FAIL:** Revoke authority, terminate invalid lane, and spawn a replacement Candidate CEO.
