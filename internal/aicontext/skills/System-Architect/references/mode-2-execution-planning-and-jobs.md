# Mode 2 — Execution Planning And Job Breakdown

> This file is part of System-Architect Skill v2. Read when: all required SPECs are APPROVED (Mode 2) to synthesize hard rules into `AGENTS.md`, break SPECs into `Docs/execution/` phases and jobs, and maintain tracking ledgers.

---

## Required Inputs — Read Before Planning

- `Docs/execution/README.md` (if already exists, read it; if not, create it in Mode 2)
- `Docs/execution/progress.md` (if already exists, read it; if not, create it in Mode 2)
- `Docs/notes_decisions_log/notes_decisions_log_YYYYMMDD.md` (if today's file already exists, read it; if not, create it in Mode 2)
- All `Docs/SPEC/*` family related (passed SPEC Readiness Check)
- Additional reference: `.agent/skills/execution_planner.md` for format details and process

---

## Core Workflow — Read SPEC -> Extract Core Architecture -> Create `AGENTS.md`

Mode 2 must run in this exact order:
1. Read all approved SPECs related to the scope
2. Extract the core architecture from those SPECs
3. Create `AGENTS.md`
4. Put that core architecture into `AGENTS.md` as hard rules
5. Create `Docs/execution/README.md`
6. Create `Docs/execution/progress.md` using the mandatory base content defined below
7. Create `Docs/notes_decisions_log/notes_decisions_log_YYYYMMDD.md`
8. Use approved SPECs + hard rules to create phase/job execution plan files in `Docs/execution/*`

### Hard Rules For Mode 2:
- `AGENTS.md` is created strictly from approved SPECs.
- `Docs/execution/README.md` is created in Mode 2.
- `Docs/execution/progress.md` is created in Mode 2.
- `Docs/notes_decisions_log/notes_decisions_log_YYYYMMDD.md` is created in Mode 2.
- Content inside `AGENTS.md` = hard rules (forbidden to violate; violation = architecture breakage).
- If a previous planning cycle already exists, update existing Mode 2 artifacts from current approved SPECs; do not treat old planning files as the source of truth.

---

## Planning Modes — Select Exactly 1 Mode Before Planning

| Mode | When to Use | Rules |
|------|-------------|---------|
| **Append** (default) | Adding new scope on top of existing execution plan | Keep existing phases/jobs intact. Add new phase after the last phase. |
| **Patch** | Fix / clarify a specific phase or job | Only modify affected files. Do not renumber unrelated phases/jobs. |
| **Reset** | Owner requests a complete plan rewrite | Clearly state in today's `notes_decisions_log_YYYYMMDD.md` that this is a reset. Rebuild from SPEC, not from old plan. |

---

## Output (Mode 2)

Mode 2 produces 3 output types:

### 1. Hard Rules — `AGENTS.md` (Root Project)

Read all approved SPECs, extract the core architecture, and put it into `AGENTS.md` as hard rules.
- Content inside `AGENTS.md` = hard rules.
- Hard rule = forbidden to violate.
- Violation = architecture breakage.
- Clearly categorize: ownership rules, data flow rules, security rules, isolation rules.
- `AGENTS.md` is the highest authority after SPEC — all lanes (coder, supervisor, QA...) must comply.
- DO NOT invent new rules — only synthesize from existing SPECs.
- Clearly state which SPEC each rule is sourced from for traceability.

---

### 2. Phase / Job Architecture — `Docs/execution/`

Split SPECs into an actionable execution plan for Coder to implement:

**Phase** = 1 folder, grouping related jobs:
```text
Docs/execution/
├── phase-1-<scope>/
│   ├── _overview.md          # Phase overview: objective, dependency, order
│   ├── job-1.1-<scope>.md
│   ├── job-1.2-<scope>.md
│   └── job-1.3-<scope>.md
├── phase-2-<scope>/
│   ├── _overview.md
│   ├── job-2.1-<scope>.md
│   └── job-2.2-<scope>.md
└── ...
```

#### Principles Applied to Every Phase / Job:

- Standardize every job to a pragmatic, contract-level template:
  - Context
  - Authority
  - Dependencies
  - Exact write scope
  - Exact read dependencies
  - Implementation tasks
  - Must preserve
  - Must reject / fail closed
  - Required tests
  - Operational evidence if applicable
  - Done criteria
- Eliminate vague language (e.g. "preserve", "aligned with", "where required", "remain") unless accompanied by exact, testable conditions.
- Each job must explicitly specify:
  - Where Coder is allowed to create or modify code.
  - Where Coder is strictly forbidden to touch.
  - What the exact input contract is.
  - What the exact output contract is.
  - Which failure paths must be handled.
  - Which unit/integration tests are mandatory.
- Hardening, logging, infra, and readiness jobs must not just state "implement X"; they must contain explicit evidence checklists and negative case criteria.
- UI jobs must pull both UI contracts and visual-reference expectations down into execution, rather than generically citing UI SPEC.
- Wiring jobs must explicitly state which route files, module entrypoints, and data boundaries are invoked.
- Boundary jobs must define DTO / read-model / write-result shapes at the contract level (no internal function names required).

#### Principles for Splitting Phases:
- Split by capability boundary and runtime boundary — not by arbitrary file batches.
- Good examples: `"identity core"`, `"public catalog"`, `"billing and entitlements"`, `"admin governance"`.
- Bad examples: `"misc fixes"`, `"remaining files"`, `"cleanup"`.
- Each phase must have: clear owner boundary, clear dependency entry point, and clear exit condition.

#### Structure of `_overview.md`:
- Phase objective
- List of jobs in the phase
- Dependencies between phases (which phase must complete first)
- Related SPECs
- Exit criteria — when is the phase complete

#### Structure of `job-*.md`:
- `## Context` — background, why this job is needed
- `## Rules` — related hard rules (reference AGENTS.md)
- `## Input` — what is needed before starting (which phase/job must complete, which SPEC to read)
- `## Scope` — what exactly to do, what NOT to do
- `## Tasks` — list of specific implementation tasks
- `## Output Files` — which file/runtime surfaces must exist after the job
- `## Verify` — exact commands to verify the job is completed correctly
- `## Done Criteria` — conditions that must be true for the job to be marked complete

#### Principles for Splitting Jobs:
- Each job must be small enough for Coder to complete and commit in 1 batch.
- Jobs must not have implicit dependencies — all dependencies must be explicitly stated in `Input`.
- Jobs must not contain architecture decisions — architecture is decided in SPEC.
- Avoid jobs that only write governance without creating specific artifacts.

---

### 3. Execution Docs

Create these tracking files in Mode 2:

#### `Docs/execution/README.md`:
- Explain the execution plan structure
- Phase ordering and dependency sequence
- How to read and navigate execution docs

#### `Docs/execution/progress.md`:
MUST be initialized with this mandatory base content:

```markdown
# Progress Tracking

Use this file as the single source of truth for execution status.
Rule: mark `x` only after both `Verify` and `Integration Gate` pass.

## Current Status
- **Mode:** Bug-fix / stabilization
- **Phase:** Implementation table complete through phase `76` (all listed jobs approved)
- **Jobs:** 824/824 approved
- **Overall:** 100.00% approved

## Approval Policy
- Each job is complete only when **both** checks are marked:
  - `Coder`: implementation + verify + integration gate done.
  - `Supervisor`: reviewed evidence and confirmed real completion.
- Do not mark phase completed if any job is missing either check.

## Integration Checklist (apply for every job)
- [ ] `wire ./...` passes
- [ ] compile/test passes
- [ ] runtime path is wired (not orphan code)
- [ ] E2E smoke works for job scope
- [ ] no TODO/FIXME/stub/dead path in touched files
- [ ] `AGENTS.md` hard rules pass
- [ ] data values/contract match original SPEC (no hardcoded/stub/raw ID) — Gate 5

## Phase Checklist and E2E Verification Log
Legend: `x` = done, `-` = pending/not verified yet.
Status values: `-`, `READY REVIEW`, `APPROVED`, `REJECTED`, `REJECTED (TECH DEBT)`, `READY REVIEW (RESUBMIT)`.
| Job | Content | Coder | Supervisor | Wire | Compile/Test | Runtime Wired | E2E Smoke | Debt-Free | Status |
|-----|---------|-------|------------|------|--------------|---------------|-----------|-----------|--------|
```

- This base content is mandatory for consistency across the workspace.
- Extend the table with real phase/job rows for the current execution plan.
- Keep summary counts accurate when adding or changing jobs.

#### `Docs/notes_decisions_log/notes_decisions_log_YYYYMMDD.md`:
- Daily log file under `Docs/notes_decisions_log/`
- One file per day for notes and decisions arising during implementation
- Each note records: timestamp (`UTC+7`), authority used, planning mode (append/patch/reset), phases/jobs changed, blockers
- If today's file does not exist yet, create it with header first

---

## Consistency Checklist — Verify Before Finishing

- [ ] Phase numbering is continuous, no duplicates
- [ ] Job numbering is continuous within each phase
- [ ] `progress.md` references match actually existing job files
- [ ] `_overview.md` job checklist matches actual `job-*.md` files
- [ ] Verify commands are viable for the target repo's actual stack, not copied from another stack
- [ ] Scope language matches `AGENTS.md` and SPEC authority
- [ ] Summary counts in `progress.md` are accurate

---

## Output Rules (Execution Planning)

- DO NOT write new SPECs — only read SPECs and generate execution plans.
- DO NOT invent hard rules — only synthesize from SPECs.
- Phases/jobs MUST NOT contain architecture decisions — only contain implementation instructions.
- Each job must be traceable back to the source SPEC.
- If a missing SPEC is discovered -> stop, go back to Mode 1 to supplement the SPEC first.
