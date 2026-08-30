# Lane and Skill Coordination

> This file is part of Orchestration Skill v2. Read when: designing a new lane, deciding skill package, share/separate lane decisions, acceptance flow.

## Principles of Lane and Skill Coordination

### The Nature of Lanes and Skills

* A lane is the unit responsible for creating a specific work outcome.
* A skill is the capability granted to a lane to achieve that outcome.
* A lane can use multiple skills.
* A skill can be used in multiple different lanes.
* Skills do not self-determine the authority or scope of a lane.

Each lane must be clearly defined by four elements:

* Ownership: The outcome the lane is responsible for.
* Capability: The skills the lane needs to use.
* Authority: What the lane is allowed to modify, check, or give a verdict on.
* Boundary: The scope the lane is permitted to touch and the point where it must stop.

Example: A Supervisor can use backend, frontend, or data-integrity skills to review, but is still not allowed to modify code because the lane's authority is review-only.

### How to Select Skills

CEO must:

* Understand the goal, pipeline, state, invariants, and acceptance of the slice before selecting skills.
* Select skills based on the guidelines in AGENTS.md and the nature of the work, not by keywords.
* Not need to read every SKILL.md to route; the skill table in AGENTS.md is used for this.
* Whichever session uses a skill, that session must fully read the corresponding SKILL.md.
* CEO only reads SKILL.md when CEO itself directly uses that skill.
* Grant the lane the full necessary skills, not limited by the lane's role name.

Examples:

* Implementation can use coder along with frontend, backend, database, design, or debugging.
* Review can use supervisor along with backend, frontend, data-integrity, edge-case, or design.
* Runtime/build errors can be supplemented with debugging.
* Real UI/browser QA uses qa.
* Mechanical status updates are delegated to a short-lived Mechanical Planner Lane.

The examples above are routing guidelines, not fixed formulas.

### When to Share or Separate Lanes

Keep work within the same lane when the tasks share:

* goal;
* ownership;
* authority;
* boundary;
* deliverables;
* completion conditions.

Only separate into a dedicated lane when there is a practical reason:

* conflicting authorities, such as simultaneously modifying and self-accepting;
* independent deliverables or boundaries;
* requiring independent zero-trust review;
* requiring the Owner to monitor or intervene separately;
* can be run independently and in parallel;
* ownership has transferred to another work unit.

Do not separate lanes just because the work requires multiple skills.

### Adjusting Lanes During Work

CEO must continuously monitor to determine:

* which skills the lane is lacking or has in excess;
* whether new work still belongs to the current lane or has separate ownership;
* whether the lane lacks evidence, authority, time, or tools;
* whether the lane deviates from scope, loops gates, or performs unnecessary work.

If ownership and boundary remain unchanged, CEO can add or remove skills directly within the current lane.

Adding skills must not automatically expand the slice. Each skill only operates within the assigned authority and boundary.

### Operating Responsibilities of CEO

CEO must:

1. Read the plan and relevant SPEC family on demand to understand the active slice and next objectives.
2. Understand the function of each phase/slice and maintain a unified progress state.
3. Only open the current slice.
4. Distinguish:
* work belonging to the current slice;
* findings that need to be moved to another slice;
* issues outside the campaign.

5. Design the session with complete:
* goal;
* ownership;
* skill package;
* authority;
* scope and non-goals;
* files/modules allowed to be touched;
* mandatory evidence;
* timeout;
* stop conditions;
* completion conditions;
* the next person to receive the handoff.

6. Monitor the actual behavior of the lane: commands, modified files, completed gates, scope, and loops.
7. Proactively handle coordination:
* if lacking a skill, add it;
* for long commands, use an appropriate timeout and wait for the exact invocation;
* for simple blockers, assign specific actions;
* for findings outside the slice, record and transfer to the correct owner;
* if a lane deviates, block it immediately.

8. Upon receiving a handoff, self-verify the report, source, diff, Git boundary, and evidence before deciding the next step.

### Acceptance and Transitioning Slices

* Only the Supervisor is allowed to give an acceptance verdict.
* QA is only used when the nature of the work truly requires QA; QA is not a default gate for all code changes.
* **(MUST)** A Commit is a true rollback anchor, not a ritual. CEO must only stage owned paths, commit exactly at valid checkpoints, check the manifest, and ensure a clean boundary. Broad resets/stashes/cleanups in a shared checkout are strictly prohibited.
* After Supervisor PASS, CEO:
1. delegates a short-lived `Mechanical Planner Lane` to tick the checklist and update actual status in `plan.md`/`actual-status.md`;
2. organizes change detection (`anvien detect-changes`);
3. commits the independent slice;
4. reads the updated section and next slice scope, packages the contract, and opens the next slice.

### Special Lane Archetype: Mechanical Planner Lane

* **Purpose:** Strictly for applying mechanical status updates to `plan.md` and `actual-status.md` upon receiving an official Supervisor PASS for an entire Slice/Phase.
* **Characteristics:** Short-lived (executes in seconds to modify the exact specified item and self-terminates immediately).
* **Iron Prohibition:** STRICTLY FORBIDDEN from auditing documents, proofreading text, or debating wording. It applies only the specified mechanical tick/state transition.
