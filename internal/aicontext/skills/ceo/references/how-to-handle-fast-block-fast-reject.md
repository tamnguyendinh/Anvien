# How to Handle FAST BLOCK & FAST REJECT

> This file is part of CEO Skill. Read when: handling in-flight emergency blockers (FAST BLOCK / FAST REJECT), dispatching emergency reinforcement, or rerouting technical repairs.

## CEO Fast Repair Rerouting Protocol

When CEO receives a `FAST BLOCK / FAST REJECT` message from a specialist lane:

1. **Acknowledge & Pause:** CEO acknowledges the blocker and pauses the reporting lane (e.g., QA) with status `BLOCKED` (preserving its working state).
2. **Package Fast Repair Packet:** CEO extracts the exact *Blocker Evidence* and *Remedy Target & Action* provided by the specialist lane without second-guessing or manual debugging.
3. **Reroute to Remediation Owner (Strict Separation of Concerns):** CEO immediately delegates the blocker to the specialized lane authorized to resolve it.
   - **(MUST - FAST BLOCK Technical Delegation Iron Rule):** When a FAST BLOCK involves code, architecture, data flow, schema, or Touch-Map conflicts (e.g., file remapping or scope modifications needed): CEO is STRICTLY FORBIDDEN from making technical decisions autonomously, deducing design solutions, or expanding the Touch-Map. CEO MUST immediately spawn an Architect Lane via `invoke_subagent` (or `mcp__codex_app__create_thread` in Codex) for the Architect to conduct professional evaluation and issue an official architecture decision / Touch-Map Addendum.
   - When a FAST BLOCK involves implementation bugs or test failures: route to Coder.
   - When a FAST BLOCK involves infrastructure or build failures: route to DevOps.
   - *Command Pattern:* "`<reporting_lane>` reported BLOCKED due to `<blocker_reason_and_evidence>`. `<remediation_owner>` must resolve the root cause, update evidence/artifacts, and re-issue `READY_FOR_<reporting_lane>` upon completion."
4. **Resume Operational Workflow:** As soon as the recovery lane reports success (PASS), the CEO must immediately reactivate the original specialized lane (the agent lane that issued the FAST BLOCK / FAST REJECT) so the subagent can seamlessly resume its interrupted work.
