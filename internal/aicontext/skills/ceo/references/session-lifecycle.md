# Session Lifecycle

> This file is part of CEO Skill. Read when: opening a new session, closing a session, performing handoff between sessions.

## Session Classification

### Separate sessions subagents, visible to the user

**Mandatory for:**

* CEO agent
* coder
* architect
* Supervisor review;
* Planner
* Long QA gates;
* Tasks requiring 1 or more specialized skills;
* Tasks with a risk of modifying production;
* Tasks with external targets/repositories;
* Tasks requiring 1 or more specialized skills that have long execution phases or times;
* Tasks where the user might need to stop or change direction midway.

### Internal subagents

**Only used for:**

* discovery read-only;
* small inventory;
* independent checks with clear boundaries;
* tasks not requiring direct user intervention;
* tasks not allowed to self-commit or expand the scope.

Do not use internal subagents to hold a long Supervisor gate and then require all other agents to wait in an unobservable state.

## Conditions Prior to Opening a Session (Lane)

A new session (lane) must fully receive:

* (MUST) Each lane must receive a complete, explicit contract from CEO: exact goal, slice, ownership, authority, exact input/path/hash, permitted commands, expected output, stop conditions, verdict criteria, and next owner. Do not leave the lane to invent its own workflow or expand the task into a general audit.
* exact goal;
* currently open plan and slice;
* scope and non-goals;
* applied authority;
* files/modules allowed to be touched;
* evidence to be collected;
* stop conditions;
* completion conditions;
* the next responsible person.

The session must not deduce/assume a new architecture from audits, file names, or keywords in the problem report on its own.

## Mandatory Acknowledgment When a Session (Lane) Starts

In the first response, the session must clearly answer with one of two states:

UNDERSTOOD or NOT UNDERSTOOD

Then it must briefly state:

* the understood goal;
* the currently open slice;
* the boundary;
* the first action.

If answering NOT UNDERSTOOD, the session must stop and accurately state the unclear point. It is not allowed to run commands, modify code, QA, cleanup, or commit before being explained.

## Standard Universal Lane Lifecycle:
**[Receive Contract] -> [Execute Assigned Role] -> [Necessary Validation Only] -> [Record Report/Evidence] -> [(MUST) Commit All Owned Output] -> [Send Direct Message to CEO] -> [HARD STOP IMMEDIATELY]**

**Subsequent Flow:** The CEO forwards the **Exact Handoff Packet** directly to the designated next specialist lane. The incoming lane focuses **100% on its own domain-specific Codebase & Runtime Invariants**, strictly prohibited from auditing the paperwork or wording of the previous lane.

**The 5 Absolute Prohibitions (Applied Universally Across All Lanes):**

1. **NO Post-Creation Self-Audits:** Strictly forbidden from self-auditing Git logs, hashes, or manifest files immediately after creating outputs.
2. **NO Post-Completion Verification Cycles:** Strictly forbidden from entering redundant post-completion re-checking loops once task deliverables are generated.
3. **NO CEO Deep-Technical Inspection:** CEO operates strictly at the management & routing level (checking verdict, report path, commit SHA); CEO never performs deep technical re-verifications of specialist outputs.
4. **NO Goal Inversion (Target is Codebase, Not Paperwork):** Reports and evidence exist solely to record progress; they must never displace the real codebase and runtime behavior as the primary target.
5. **NO Autonomous Subagent Spawning:** Individual subagent lanes are strictly forbidden from independently opening secondary reviewer, auditor, or helper lanes.

## Handoff Between Sessions

Each handoff must point to:

* plan/slice;
* report;
* evidence IDs;
* commit or HEAD;
* current worktree;
* open blockers;
* next to Orchestration agent (CEO agent).

The subagent's result is not automatically a conclusion. The Orchestration agent (CEO agent) must read the durable output and verify it according to the Supervisor protocol.

## Conditions for Closing a Session (Lane)

A session may only be closed when:

* the slice's goal has been evaluated;
* the report is recorded;
* evidence IDs are updated;
* open blockers are categorized;
* the verdict is clear;
* the next handoff is determined.

Do not declare completion just because code/build/test can run.

## Template Prompt for Opening a Session:

- for codex: .agents/skills/ceo/references/Template-Prompt-for-Opening-a-Session.md
- for Claude code: .claude/skills/ceo/references/Template-Prompt-for-Opening-a-Session-claude.md
- for other models: .agents/skills/ceo/references/Template-Prompt-for-Opening-a-Session.md
