# Invariant And Resubmission

> This file is part of Supervisor Skill v2. Read when: verifying invariant closure, sweeping sibling surfaces, closing prior issue history, or evaluating resubmissions after fix attempts.

## Invariant Closure

Do not approve a local symptom fix when the same invariant may span other surfaces.

Identify the affected invariant: runtime contract, data integrity rule, owner boundary, permission rule, isolation rule, API shape, tool contract, state transition, generated artifact contract, or process rule.

Start from the provided artifact or diff, then sweep only the relevant same-invariant surfaces, such as:

- route or entrypoint;
- alternate trigger;
- UI panel, dialog, or state path;
- store, service, API, tool handler, repository, schema, job, worker, or generated contract;
- stale helper, fallback path, fixture, test, or doc contract when it can preserve the old behavior.

Do not expand into unrelated domains. Do not approve until the affected invariant is closed for the reviewed scope.

## History Closure

If prior reports, review comments, blocker notes, QA findings, bug reports, or resubmissions exist in the same scope, consume them as evidence pointers.

Do not read only the latest artifact when unresolved earlier evidence can still affect acceptance.

A prior issue is closed only when current evidence proves it is closed. If closure cannot be proven, REJECT.

## Resubmission Review Guard

When reviewing a fix after prior rejection, QA finding, review comment, blocker note, or failed claim:

1. Start from the previous blocking finding.
2. Verify the claimed fix in source/project reality.
3. Check the same invariant surfaces named in the prior review.
4. Confirm the old failure mode cannot still occur in the reviewed scope.
5. Require fresh evidence for closure.

A resubmission is not accepted because it addresses the latest visible symptom. It is accepted only when current evidence closes the prior blocker and the affected invariant.
