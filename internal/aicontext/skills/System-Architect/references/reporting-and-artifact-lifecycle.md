# Lane Reporting And Artifact Lifecycle

> This file is part of System-Architect Skill v2. Read when: generating lane reports, applying report naming conventions, or managing architecture artifact commits.

---

## Lane Report

Used by both Mode 1 and Mode 2. A report is required when work is completed. Each report must contain:
- Scope — what was done in this session
- Output files created — list of files created (SPEC, ADR, phase, job, `AGENTS.md`...)
- Decisions made — summary of architectural decisions
- Residual open questions — unanswered technical questions
- Commit reference

---

## Report Naming Rules

- Report folder: `reports/system-architect/`
- File name: `reports/system-architect/rp_system-architect_<YYMMDD>_<HHMMSS>_by_<model_slug>_<scope>.md`
- Use `system-architect` to distinguish from the `architect-review` lane (review lane uses `reports/architect-review/`).
- `model_slug`: lowercase ASCII, use `-` if needed, no underscores.
- `scope`: lowercase `snake_case` summarizing the content.
- Must commit report before finishing.
- Old reports must not be overwritten — create a new report with a fresh timestamp.

---

## Artifact Commit Rules

When this role writes repo artifacts such as:
- ADRs
- Architecture notes
- Design proposals
- Boundary or ownership documents
- SPECs
- Execution plans

It must stage and commit those artifacts before finishing.

### Commit Rules:
- Commit only the files created or updated by this architecture lane:
  - `reports/system-architect/*`
  - `Docs/SPEC/*` (SPEC files created or updated)
  - `Docs/execution/*` (execution plan files — Mode 2 only)
  - `AGENTS.md` (hard rules — Mode 2 only)
  - matching shared blocker handoff files in `reports/problem/*` when created by this lane
- Do not overwrite an older architecture report just because there is a later follow-up.
- A new architecture step must produce a new timestamped report artifact; old reports stay as historical record unless they were improperly overwritten and need restoration.
- Do not leave architecture docs untracked or half-written in the worktree.
- Do not commit code, screenshots, test artifacts, `.tmp/`, or unrelated files unless the user explicitly asks for them.
- All communication between lanes must go through durable report files. No communication via chat.
