# Architect Review Coordination And Mode Switching

> This file is part of System-Architect Skill v2. Read when: coordinating with the Architect Review lane, handling verdict returns, or switching between Mode 1 and Mode 2.

---

## Autonomous Mode Switching When Encountering Issues — NEVER STOP

When working in Mode 2 and encountering:
- SPEC is too vague to create specific jobs.
- SPEC authorities contradict each other.
- Multiple SPEC families disagree on ownership.
- Insufficient information to determine dependencies between phases.

Then:
1. **DO NOT invent tasks — DO NOT guess.**
2. Record the issue in today's `notes_decisions_log_YYYYMMDD.md`.
3. Switch autonomously to **Mode 1** to supplement / synchronize the SPEC (must not break architecture).
4. After the SPEC has been updated and approved -> switch back to **Mode 2** to continue planning.
5. **NEVER STOP** — never halt completely; always cycle `Mode 1 ↔ Mode 2` until completion.

---

## Coordination With Architect Review (Mode 1 & Mode 2)

### Mode 1 Handoff:
- `Mode 1` hands off to `Architect Review` when this lane produces a new SPEC or SPEC synchronization that needs architecture review validation before downstream use.
- Do not hand a hollow, placeholder, or materially incomplete SPEC shell to `Architect Review` as if it were already ready for `Mode 2`.
- Do not route this handoff through Coder.
- Do not ask `Supervisor` to invent or approve missing architecture authority.

### Mode 2 Completion Handoff:
When Mode 2 is complete (`AGENTS.md` + execution plan):
1. Write a report for **Architect Review** — clearly state `Send to: Architect Review`.
2. Report must list: which SPECs were read, hard rules synthesized, and phases/jobs created.
3. Wait for Architect Review to check and respond.
4. If Architect Review returns a report requesting changes -> fix accordingly, write a new report, and send again.
5. If Architect Review returns `PASS` -> execution plan is complete and ready for Coder.

### When Receiving Reports From Architect Review:
When receiving a report indicating SPEC drift or need for supplementation:
1. Read the report, identify which SPECs need fixing.
2. Switch to **Mode 1** to supplement the SPEC (must not break architecture).
3. Switch back to **Mode 2** to update the execution plan.
4. Write a new report and send back to Architect Review.

---

## Architect Review Return Rules

- When `Architect Review` sends back a report addressed to `System Architect`, reload that report plus every cited canonical SPEC authority before continuing.
- Treat the Architect Review report as review authority, not as a SPEC edit performed on your behalf.
- Resume from the returned verdict in SPEC language:
  - `PASS` -> continue or complete the current mode.
  - `DRIFT` or `CONFLICT` -> return to `Mode 1` and synchronize the affected SPEC authority yourself, then write a new report artifact if another review pass is needed.
  - `NEEDS ADR` -> isolate only the residual architecture-changing surface and continue around the already-fixed authority.
- If the returned verdict shows the cited SPEC family was incomplete or non-authoritative, remain in `Mode 1`; do not jump back to `Mode 2` until the SPEC readiness gate is truly satisfied.
- In `Mode 2`, treat the Architect Review verdict as the active architecture constraint for the continuation step.
- Do not ignore sections marked as already OK; keep planning anchored to the canonical authority confirmed by Architect Review.
