---
name: working-rules
description: This skill should be used when a work session begins, before the session selects skills, uses tools, edits files, validates results, or commits work.
---

# Working Rules

## 1. Rules Are the Highest Layer of Coordination
- Before doing anything, you must read `AGENTS.md` completely.
- `AGENTS.md` coordinates how to find rules, select skills, use tools, implement, verify, and accept work.
- Mandatory sequence: rules first, skills second.
- Never use summaries, prior session memory, or compacted context in place of the original raw rules.
- If context is compacted or the task changes, re-read the required original sources before taking action.

## 2. Every Decision Must Be Evidence-Based
- Do not deduce from verbal statements.
- Do not deduce from keywords.
- Do not deduce from filenames, module names, or terminology similarities.
- Do not project patterns from other systems, projects, or applications onto this app.
- Do not treat user statements as automatic ground truth.
- User opinions, rebuttals, or hypotheses are perspectives that require additional verification.
- Conclusions must be proven by rules, source documents, source code, runtime, or corresponding evidence.
- Do not ask the user about things already documented in the repository; independently trace the evidence first.

## 3. Must Understand the Entire Problem Before Acting
When the user presents an issue:
- Must identify the entire scope of the issue.
- Must read all relevant rules and documentation completely.
- Never simply find a snippet with keywords matching the user's words and jump to conclusions.
- Never shift focus to the easiest detail to fix.
- Must distinguish between root cause, symptoms, boundaries, and associated impacts.

When the user revisits a small slice:
- That slice is only a supplementary perspective.
- Never turn a small slice into the entire problem.
- Never lose sight of the original goal and scope.

## 4. Classify Pipelines by Real Linkages
- Each pipeline must be identified by its command, state, ownership, boundary, and outcome.
- Do not group two pipelines simply because they use similar words.
- Two pipelines are only linked when a command in one pipeline genuinely requires state from the other pipeline to make decisions.
- Do not invent business relationships just to make the architecture look more complete.
- Do not complicate an inherently simple flow unless the repository contains evidence for such complexity.

## 5. Independently Find Rules for Every Action
Before modifying code, documentation, QA, artifacts, or committing:
- Must identify which rules govern that action.
- Must find all relevant rule files; rules may reside in multiple locations.
- Never assume reading a single file is sufficient.
- Never rely on conversational context as a substitute for rule tracing.
- If evidence is not found, continue investigating in the repository before asking the user.

## 6. Select Skills Only After Understanding Rules and Tasks
- Do not select a skill first and force the problem to conform to the skill.
- Must determine the actual nature of the work first, then select skills.
- Must read the full `SKILL.md` of each skill being used.
- Never compact context or pick convenient subsets of rules within a skill.
- Skill-specific rules apply only within the scope of tasks governed by that skill.
- While working, if a new type of task emerges, find additional suitable skills immediately.
- Do not restrict yourself to an initial set of skills selected at the beginning of the session.
- Do not invoke skills mechanically just because the skill name shares words with the issue.

## 7. Scoped Use of Subagents
- Only invoke subagents for independent, specific tasks with clear boundaries.
- Never assign a vague problem or an entire task to a subagent and passively wait for results.
- Subagents must not autonomously expand scope.
- Subagent output is only verification input; it never automatically becomes a conclusion.
- Main agent must independently verify results using evidence and the Supervisor protocol.

## 8. Plan Proportionally Before Code
- Before coding, use `planner` and create a real `docs/plans` plan when the work is multi-step, affects multiple files or modules, changes behavior, contracts, or architecture, carries meaningful risk, or when the user explicitly requests a plan.
- Do not require a plan for a trivial, atomic, low-risk edit with obvious scope, such as correcting a few words, deleting one redundant line, or making a self-contained one-line change that does not alter behavior, contracts, architecture, or cross-module boundaries.
- Validate a trivial planless edit at the nearest relevant boundary.
- When uncertain whether a change is trivial, create a plan.
- When creating, writing, or editing a plan, use `planner`; the plan must be a real document in `docs/plans`.
- When a plan governs the work, read all relevant ledgers completely and only implement the currently open slice.
- Never autonomously advance phases or slices in planned work.
- When discovering issues that alter a planned slice's scope, update the plan before continuing to code.
- Update plan checklists, actual status, evidence, and benchmarks immediately as corresponding states change.
- Only transition to the next planned slice after the current slice completes the full acceptance and commit workflow.

## 9. Rules When Working with Prototypes
### a. Prototype leads the way:
- UI/UX Prototype is the active refinement target.
- Do not use SPEC to impose backward constraints onto the prototype.
- Do not read or use product SPEC during the active refinement loop.
- Do not autonomously update SPEC.
- Only update SPEC when the user issues the exact command "update spec".
- Equivalent phrasing must not be inferred as a spec update command.
- `ui-driven-spec` must still be used to maintain the UI-first methodology, boundaries, and handoff readiness; do not use it to reverse direction and force SPEC onto the prototype.

### b. Prototype must represent flows clearly:
The prototype is not real backend or database, but must clearly represent:
- What command the user issues.
- Which state owns the data.
- Where state transitions occur.
- Where the presenter retrieves data.
- Which module is responsible.
- Where boundaries between pipelines lie.
- Which flows are demo data only.
- What real future code needs to implement.

Must NOT:
- Simulate non-existent backend architectures.
- Create fake database bindings.
- Add unrequested business logic.
- Use snapshots or intermediary layers without evidence (snapshot is exclusive to SYNC domain).
- Turn restaurant/hotel/bar management apps into accounting systems or other application types.

## 10. Protect Scope and Module Responsibilities
- Only edit files/modules within the slice scope.
- Never touch functionalities of other tabs/modules if the slice does not own them.
- Sibling modules are only verified for preservation/non-regression.
- Each file must own a single primary business responsibility.
- When a new responsibility emerges, find or create the rightfully owned file/module.
- A file may call multiple other modules but must not encompass more than one unrelated business domain.
- Never opportunistically refactor or touch areas outside scope.

## 11. Use Anvien Strictly According to Rules
- Must review Anvien instructions before use.
- Must refresh graph (`anvien analyze --force`) before any graph-based work.
- Before modifying functions, classes, methods, exported symbols, shared contracts, or objects listed in rules, run `file-detail` and `impact` analysis.
- Must report blast radius.
- HIGH or CRITICAL blast radius is a warning for extra caution, not a prohibition against editing.
- Never reduce or conceal graph evidence to make output appear simpler.
- Must run `detect-changes` before committing implementation work.

## 12. Correct Code First, QA Second
Mandatory sequence:
1. Understand rules and required behavior.
2. Modify production code.
3. Update QA only after behavior is correct in code.
4. Run full build.
5. Open the real user-visible runtime.
6. Verify behavior at the nearest real boundary.
7. Execute QA and collect evidence.
8. Visually inspect results.
9. Run regression for preserved boundaries.

Never modify tests beforehand to force a test PASS while production code is not yet correct.

## 13. QA Must Prove Real Runtime Behavior
- Playwright scripts must be reusable under `playwright/`.
- Do not use temporary scripts as official evidence.
- `.tmp` is debug-only and must stay inside the repository.
- Official evidence must be stored in `Reports/qa/playwright/...`.
- Evidence must include both JSON and Markdown formats.
- UI must be verified on real runtime.
- Must visually inspect screenshots with human eyes, never rely solely on assertions.
- Never use unrelated tests, stale tests, pass-by-default tests, or superseded evidence to claim completion.
- Regression must verify both the changed feature and sibling boundaries at risk.

## 14. Artifacts Must Follow Lifecycle Management
- Failed, retried, duplicated, or superseded artifacts are dead work.
- Dead work must be deleted at the close of the corresponding phase/slice.
- Never allow intermediate artifacts to accumulate.
- Explicitly identify which artifacts are superseded before deletion.
- Active, valid evidence must be retained.
- If the user requests retaining or committing specific artifacts, handle strictly within that scope.

## 15. Acceptance Under Zero-Trust
- Never self-declare completion simply because code runs.
- All completion claims must pass through Supervisor.
- Supervisor must independently verify source, diff, runtime, reports, screenshots, and required evidence.
- Output from Coder, QA, or subagents is never automatically accepted.
- Only mark a slice completed upon Supervisor PASS.
- If Supervisor REJECTs, fix only the rejected invariant and repeat the verification cycle.

## 16. Commit Slice-by-Slice
Before committing implementation work:
- Production code completed.
- Full build PASS.
- Runtime validation PASS.
- Valid QA evidence gathered.
- Necessary regression tests PASS.
- Intermediate artifacts cleaned up.
- Plan ledgers updated.
- Supervisor PASS.
- Anvien `detect-changes` completed.

Each slice must have an independent commit. Never bundle the next slice into the current commit.

## 17. Communication Discipline During Work
- Must report current actions and the governing rules or evidence.
- Do not remain silent for extended periods during build, QA, or subagent waiting.
- Do not report assumptions as facts.
- Clearly distinguish: Verified, Checking, and No evidence yet.
- When encountering blockers, state exactly where the blocker lies.
- Do not push decisions back to the user if the repository already defines the rules.
- When the user provides corrections or rebuttals, re-verify all associated aspects.
- When the user orders a stop, stop immediately: do not continue coding, QA, deleting, updating documents, committing, or directing subagents.

---
*This document defines working rules only; it does not contain progress tracking, handoff statuses, or plan results.*
