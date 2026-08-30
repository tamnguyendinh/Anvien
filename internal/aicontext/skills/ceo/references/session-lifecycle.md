# Session Lifecycle

> This file is part of Orchestration Skill v2. Read when: opening a new session, closing a session, performing handoff between sessions.

## 2. Session Classification

### Separate sessions subagents, visible to the user

**Mandatory for:**

* orchestration (main agent)
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

## 3. Conditions Prior to Opening a Session (Lane)

A new session (lane) must fully receive:

* (MUST) Each lane must receive a complete, explicit contract from Main: exact goal, slice, ownership, authority, exact input/path/hash, permitted commands, expected output, stop conditions, verdict criteria, and next owner. Do not leave the lane to invent its own workflow or expand the task into a general audit.
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

## 4. Mandatory Acknowledgment When a Session (Lane) Starts

In the first response, the session must clearly answer with one of two states:

UNDERSTOOD or NOT UNDERSTOOD

Then it must briefly state:

* the understood goal;
* the currently open slice;
* the boundary;
* the first action.

If answering NOT UNDERSTOOD, the session must stop and accurately state the unclear point. It is not allowed to run commands, modify code, QA, cleanup, or commit before being explained.

## 9. Handoff Between Sessions

Each handoff must point to:

* plan/slice;
* report;
* evidence IDs;
* commit or HEAD;
* current worktree;
* open blockers;
* next to Orchestration agent (main agent).

The subagent's result is not automatically a conclusion. The Orchestration agent (main agent) must read the durable output and verify it according to the Supervisor protocol.

## 10. Verification Policy

A session must not evaluate itself as completed. Completion evaluation belongs to:

* Supervisor;
* QA;
* Architect;
* Main session.

Do not continue just because the subagent "seems to be done" or has successfully run tests.

## 11. Conditions for Closing a Session (Lane)

A session may only be closed when:

* the slice's goal has been evaluated;
* the report is recorded;
* evidence IDs are updated;
* open blockers are categorized;
* the verdict is clear;
* the next handoff is determined.

Do not declare completion just because code/build/test can run.

## 13. Template Prompt for Opening a Session:

- for codex: .agents/skills/ceo/references/Template-Prompt-for-Opening-a-Session.md
- for Claude code: .claude/skills/ceo/references/Template-Prompt-for-Opening-a-Session-claude.md
- for other models: .agents/skills/ceo/references/Template-Prompt-for-Opening-a-Session.md
