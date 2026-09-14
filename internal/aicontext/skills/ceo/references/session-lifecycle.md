# Subagent Lane Lifecycle

> This file is part of CEO Skill. Read when: spawning a new subagent lane, closing a lane, or performing handoff between lanes.

## Lane Classification

### Spawned Agents (Subagent Lanes), Visible & Interactable to Owner

**Mandatory for:**

* Coder
* Architect
* Supervisor review
* Planner
* Long QA gates
* Tasks requiring 1 or more specialized skills
* Tasks with a risk of modifying production
* Tasks with external targets/repositories
* Tasks requiring 1 or more specialized skills that have long execution phases or times
* Tasks where the user might need to stop or change direction midway.

> **(IRON RULE - ANTI-SESSION EXPLOSION):** All functional worker lanes MUST be initialized via agent spawning (`spawn agent / create agent`) within the active campaign environment, visible and directly interactable to the Owner. DO NOT open detached top-level sessions/windows for worker subagents to avoid flooding the workspace with hundreds of redundant sessions.

### Internal Subagents

**Only used for:**

* discovery read-only;
* small inventory;
* independent checks with clear boundaries;
* tasks not requiring direct user intervention;
* tasks not allowed to self-commit or expand the scope.

Do not use internal subagents to hold a long Supervisor gate and then require all other agents to wait in an unobservable state.

## Conditions Prior to Spawning a Lane (Agent)

A newly spawned agent lane must fully receive:

* (MUST) Each lane must receive a complete, explicit contract from CEO: exact goal, slice, ownership, authority, exact input/path/hash, permitted commands, expected output, stop conditions, verdict criteria, and designated recipient (Thread ID of the CEO Session). Do not leave the lane to invent its own workflow or expand the task into a general audit.
* exact goal;
* currently open plan and slice;
* scope and non-goals;
* applied authority;
* files/modules allowed to be touched;
* evidence to be collected;
* stop conditions;
* completion conditions;
* designated recipient (Thread ID of the CEO Session).

The spawned agent must not deduce/assume a new architecture from audits, file names, or keywords in the problem report on its own.

## Mandatory Acknowledgment When a Lane Starts

In the first response, the spawned agent must clearly answer with one of two states:

UNDERSTOOD or NOT UNDERSTOOD

and send a direct message to the Thread ID of the CEO Session.

Then it must briefly state:

* the understood goal;
* the currently open slice;
* the boundary;
* the first action.

If answering NOT UNDERSTOOD, the agent must stop and accurately state the unclear point. It is not allowed to run commands, modify code, QA, cleanup, or commit before being explained.

## Handoff Between Lanes

Each handoff must point to:

* plan/slice;
* report;
* evidence IDs;
* commit or HEAD;
* current worktree;
* open blockers;
* next to Orchestration agent (Thread ID of the CEO Session).

The subagent's result is not automatically a conclusion. The Orchestration agent (CEO Session) must read the durable output and verify it according to the Supervisor protocol.

## Conditions for Closing a Lane

An agent lane may only be closed when:

* the slice's goal has been evaluated;
* the report is recorded;
* evidence IDs are updated;
* open blockers are categorized;
* the verdict is clear;
* the next handoff is determined.

Do not declare completion just because code/build/test can run.

## Template Prompt for Spawning a Lane:

- `internal/aicontext/skills/ceo/references/Template-Prompt-for-Opening-a-Session.md`
