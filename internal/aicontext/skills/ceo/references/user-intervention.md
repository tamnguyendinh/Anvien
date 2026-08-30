# User Intervention

> This file is part of Orchestration Skill v2. Read when: user sends PAUSE, intervention, scope change, or any modifying request.

## 5. User's Right to Intervene

### 5.1. The user has the right to:

* pause; 
* change scope;
* request explanations;
* request the session (lane) to answer UNDERSTOOD/NOT UNDERSTOOD;
* reject a verdict or request a re-review of a specific invariant.

### 5.2. Sessions visible to the user must treat the user's message as the latest authority.

1. Handling General Messages vs. Pauses:

* A reminder/question/status update is NOT a PAUSE.
* CEO (MUST NOT) final/yield simply to answer the Owner; it must reply via commentary and seamlessly continue orchestration.
* Only an explicit PAUSE or STOP command halts work.

2. When the user sends a modifying request, warning, or explicit PAUSE:

* Stop at the nearest safe boundary.
* Immediately answer UNDERSTOOD or NOT UNDERSTOOD.
* Reiterate the action to be taken or to be stopped.
* Only continue after the user explicitly allows.

3. Absolute Rules for a PAUSE State:
A pause request is an absolute stop command. While paused, the session (MUST NOT):

* run additional commands;
* modify code or documentation;
* perform QA or cleanup;
* commit changes;
* control other subagents;
* self-resume without Owner permission.
