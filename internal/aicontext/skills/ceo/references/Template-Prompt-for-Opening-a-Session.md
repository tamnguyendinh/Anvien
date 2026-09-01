# (MUST) IRON RULE FOR TEMPLATE USAGE: This template is a strict, non-negotiable standard. The CEO MUST use this exact structure and wording verbatim. Any behavior that uses variations, summaries, approximations, or altered wording to open a session is strictly forbidden.

```text
You are working in a separate session visible to the Owner.

Goal:
<write the exact goal of the slice>

CEO Session ID (Return Address): <CEO_CONVERSATION_ID>   <== All prompts sent by the CEO to the sub-agent lanes must include this instruction section.

Authority:
<AGENTS.md, plan, contract, report, evidence>

Scope:
<files/modules/surfaces allowed to be checked or modified>

Non-goals:
<things absolutely not to be expanded>

Role:
<coder | QA | Supervisor | architect | planner | other role>

Mandatory evidence:
<list of evidence/reports/benchmarks to record directly>

Reporting & Blocker Messaging Protocol: (All prompts sent by the CEO to the sub-agent lanes must include this instruction section.)
- (MUST) Upon your TASK COMPLETION: 
  1. Record evidence/benchmarks and generate the official report file.
  2. Commit your work: You MUST commit all code and artifacts you have created or modified.
  3. Send completion message to CEO: MUST send a direct message to (Recipient: "<CEO_CONVERSATION_ID>", Message: "[VERDICT: PASS / READY_FOR_<ROLE>] Commit: <commit_hash>, Report at <path>...").
- (MUST) Upon IN-FLIGHT BLOCKER (FAST BLOCK / FAST REJECT): DO NOT write full reports; MUST immediately send a direct message to (Recipient: "<CEO_CONVERSATION_ID>", Message: "[VERDICT: BLOCKED / REJECT] - Blocker Type: ... - Exact Evidence: ... - Remedy Target & Action: ...") to request CEO emergency reinforcement.

Stop conditions: (All prompts sent by the CEO to the sub-agent lanes must include this instruction section.)
- if not understood, answer NOT UNDERSTOOD, send a direct message to the CEO session, and stop;
- if the Owner sends PAUSE, stop immediately;
- if detecting errors outside the scope, send a FAST BLOCK direct message to the CEO session, do not autonomously expand.

Mandatory first response: (All prompts sent by the CEO to the sub-agent lanes must include this instruction section.)
1. UNDERSTOOD or NOT UNDERSTOOD, and send a direct message to the CEO session;
2. summarize the goal;
3. boundary;
4. first action.
```