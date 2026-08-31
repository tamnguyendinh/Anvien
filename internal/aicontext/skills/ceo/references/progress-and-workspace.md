# Progress and Workspace

> This file is part of CEO Skill. Read when: reporting progress, managing artifacts, workspace rules.

## Rules for Reporting Progress

The session must clearly distinguish:

* Verified;
* Checking;
* No evidence yet;
* Blocked.

Before long commands or long QA, the session must report:

* what it is doing;
* what that command proves;
* where the artifact/output is located;
* conditions to continue.

Do not report assumptions as facts. Do not remain silent for prolonged periods while running a gate.

## Workspace and Artifact Rules

The session must:

* keep temporary artifacts in the repo-local `.tmp`;
* protect the user worktree;
* only delete artifacts strictly identified as dead work;
* not use broad wildcard cleanups when there is a risk of touching other artifacts;
* not commit when lacking sufficient build, runtime, evidence, Supervisor, and detect-changes;
* not modify files outside the scope.
