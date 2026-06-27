---
name: spec-to-svg-flow-map
description: Convert product, feature, UI, backend, auth, sync, lifecycle, external-contract, or multi-branch specs into semantic SVG flow maps with machine-readable metadata, gap detection, verification reports, and BLOCKED/READY_FOR_OWNER_REVIEW status. Use when the user asks to turn a spec into a visual flow map, audit spec completeness before coding, expose undefined behavior, or identify owner decisions before implementation.
---

# Spec to SVG Flow Map

## Purpose

Convert a spec into a structured `.svg` flow map that is readable by both humans and code/text tools.

This skill does not create decorative diagrams. It creates a semantic visual map that exposes:

- user/UI flow
- data/state/storage/source-of-truth flow
- auth/session/device/permission/security boundary flow
- sync/network/lifecycle/reconnect/recovery flow
- backend/API/external system contract flow
- error/recovery/background job/runtime command flow
- decisions, branches, pipeline handoffs, terminal states, and out-of-scope references
- spec gaps, undefined behavior, owner decisions, missing data, bottlenecks, and risk areas

The SVG must help humans and coding agents decide whether the spec is clear enough for safe implementation.

## Scope And Non-Goals

- Use this skill only to create or update flow-map artifacts from a spec.
- Do not implement app/source code while using this skill.
- Do not turn unresolved ambiguity into guessed behavior.
- Do not use this skill to make a decorative diagram without semantic metadata.
- Do not begin implementation from this skill output unless the Owner has reviewed and approved it.

## Required Inputs

Read the available authority for the requested spec slice:

1. Source spec.
2. Related authority docs.
3. UI prototype or UI slot map, if present.
4. State/source map, if present.
5. Backend contract map, if present.
6. Actual wiring status, if present.
7. Project rules.

If required input is missing, do not invent behavior. Mark it as a gap in the SVG and in the verification report.

## Required Outputs

Create or update these files:

```text
docs/flow-maps/<feature-name>.flow.svg
docs/flow-maps/<feature-name>.flow-map.md
docs/flow-maps/<feature-name>.flow-verification.md
```

Use lowercase kebab-case for `<feature-name>` unless the project already has a stricter naming rule.

## Hard Rules

1. Do not implement app/source code.
2. Do not modify production source files.
3. Do not infer missing behavior.
4. Do not hide ambiguity.
5. Do not put important logic only in geometry or visual position.
6. Put every important flow relationship in visible text and XML metadata.
7. Give every decision node complete outgoing branches.
8. Give every edge a clear condition; use `condition="always"` for unconditional flow.
9. Represent every pipeline handoff with a real junction node.
10. Give every terminal state an explicit node.
11. Mark every unresolved ambiguity as a gap node.
12. If the map cannot be completed safely, stop and report why.

If the map contains any unresolved node with type `SPEC_GAP`, `UNDEFINED_BEHAVIOR`, or `OWNER_DECISION_REQUIRED`, set implementation status to `BLOCKED`.

## Semantic SVG Contract

The SVG must be valid XML and still read like source code.

Each node must use this shape:

```xml
<g
  id="node-..."
  data-type="..."
  data-flow="..."
  data-status="..."
  data-lane="..."
  data-source-ref="..."
>
  <title>...</title>
  <desc>...</desc>
  <text>...</text>
</g>
```

Each edge must use this shape:

```xml
<g
  id="edge-..."
  data-type="edge"
  data-edge-type="..."
  data-from="..."
  data-to="..."
  data-condition="..."
  data-flow="..."
  data-source-ref="..."
>
  <title>...</title>
  <desc>...</desc>
  <path ... />
  <text>...</text>
</g>
```

Never create anonymous paths for important logic.

Bad:

```xml
<path d="M120 80 L300 80" />
```

Good:

```xml
<g
  id="edge-auth-success-to-license-check"
  data-type="edge"
  data-edge-type="CONTROL_FLOW"
  data-from="AUTH_LOGIN"
  data-to="LICENSE_CHECK"
  data-condition="auth_success"
  data-flow="AUTH_FLOW"
  data-source-ref="spec.md#auth-login"
>
  <title>AUTH_LOGIN -> LICENSE_CHECK</title>
  <desc>User credentials are valid. Continue to license validation.</desc>
  <path d="M120 80 L300 80" />
  <text>if auth_success</text>
</g>
```

## SVG Metadata

Every SVG must include one machine-readable metadata block:

```xml
<metadata id="spec-flow-map">
{
  "feature": "<feature-name>",
  "source_spec": "<source-spec-path>",
  "version": "draft",
  "nodes": [
    {
      "id": "AUTH_LOGIN",
      "type": "SYSTEM_ACTION",
      "lane": "AUTH_SECURITY",
      "flow": "AUTH_FLOW",
      "status": "defined",
      "source_ref": "spec.md#auth-login"
    }
  ],
  "edges": [
    {
      "id": "edge-auth-success-to-license-check",
      "type": "CONTROL_FLOW",
      "from": "AUTH_LOGIN",
      "to": "LICENSE_CHECK",
      "condition": "auth_success",
      "flow": "AUTH_FLOW",
      "source_ref": "spec.md#auth-login"
    }
  ],
  "gaps": [],
  "terminal_states": [],
  "junctions": [],
  "legend": {}
}
</metadata>
```

Metadata must match the visible SVG nodes and edges.

## Lanes, Nodes, Edges

Use only lanes relevant to the current spec slice.

Suggested lanes:

```text
USER
UI_APP
LOCAL_STATE
LOCAL_DB
SYNC_ENGINE
BACKEND_API
AUTH_SECURITY
PAYMENT_LICENSE
EMAIL_NOTIFICATION
ERROR_RECOVERY
EXTERNAL_SYSTEM
OWNER_DECISION
```

Every node must belong to exactly one lane.

Node types:

```text
START
END
EVENT
SCREEN
USER_ACTION
SYSTEM_ACTION
BACKGROUND_JOB
DECISION
STATE
DATA_STORE
DOCUMENT
EXTERNAL_SYSTEM
JUNCTION
ERROR
RECOVERY
SPEC_GAP
UNDEFINED_BEHAVIOR
OWNER_DECISION_REQUIRED
OUT_OF_SCOPE
```

Edge types:

```text
CONTROL_FLOW
DATA_FLOW
ASYNC_EVENT
SYNC_FLOW
ERROR_FLOW
RECOVERY_FLOW
HANDOFF
BLOCKED_BY
OUT_OF_SCOPE_REFERENCE
```

Each edge must have `from`, `to`, `condition`, `data-edge-type`, `flow`, and `data-source-ref`.

## Visual Contract

Use stable visual meaning:

```text
START / END                  = circle
SCREEN                       = rounded rectangle
USER_ACTION                  = rounded rectangle
SYSTEM_ACTION                = rectangle
BACKGROUND_JOB               = dashed rectangle
DECISION                     = diamond
STATE                        = pill / capsule
DATA_STORE                   = cylinder
DOCUMENT                     = document icon/shape
EXTERNAL_SYSTEM              = dashed rectangle
JUNCTION                     = small circle
ERROR                        = red rounded rectangle
RECOVERY                     = orange rounded rectangle
SPEC_GAP                     = thick red border rectangle
UNDEFINED_BEHAVIOR           = thick red border diamond
OWNER_DECISION_REQUIRED      = thick orange border rectangle
OUT_OF_SCOPE                 = gray dashed rectangle
```

Use stable color meaning:

```text
Green       = normal / automated / value-producing flow
Blue        = data / storage / source of truth
Purple      = auth / security / permission boundary
Yellow      = delay / pending / waiting
Red         = bottleneck / error / danger / spec break
Orange      = owner decision / unclear business rule
Gray        = manual / external / out of scope
Black       = default control flow
Dashed line = async / background / indirect dependency
Solid line  = direct flow
Double line = sync / bidirectional data exchange
```

Do not change color meanings per diagram.

## Gap Detection

Create a gap node when:

1. A decision node lacks complete branches.
2. A reachable state has no next step.
3. An error can occur but has no recovery path.
4. A sync conflict has no resolution rule.
5. A backend contract is mentioned but not defined.
6. A UI state is required but not defined.
7. Payment, license, auth, or permission affects the flow but is unclear.
8. Source of truth is unclear.
9. A manual step lacks actor or owner.
10. A terminal state is vague.
11. A flow crosses into another pipeline without a junction.
12. A background job can fail but has no retry/failure rule.

Every gap node must be visible in the SVG and listed in the verification report.

## Pipeline Handoff Rule

When one pipeline talks to another pipeline, create a junction node.

Example:

```text
AUTH_FLOW -> LICENSE_FLOW
```

Represent it as:

```text
AUTH_LOGIN -> AUTH_SUCCESS_JUNCTION
AUTH_SUCCESS_JUNCTION -> LICENSE_CHECK
```

Do not represent a pipeline handoff only as crossing arrows.

## Workflow

1. Extract from the spec:

```text
actors, lanes, entry points, screens, states, actions, decisions, branch conditions, data stores, external systems, background jobs, terminal states, errors, recovery paths, manual steps, automation points, unknowns
```

2. Build a node inventory:

```text
Flow ID, Node ID, Node Type, Lane, Label, Description, Source Spec Reference, Risk Level
```

3. Build an edge inventory:

```text
Edge ID, From Node, To Node, Edge Type, Condition, Flow ID, Source Spec Reference
```

4. List gaps before writing SVG:

```text
missing branches, undefined conditions, missing states, unclear ownership, missing recovery path, missing data source, unclear terminal state
```

5. Create `docs/DOCS/flow-maps/Spec-to-SVG-Flow-Map.svg` as semantic XML.
6. Create `docs/DOCS/flow-maps/<feature-name>.flow-map.md` with flow, lane, node, edge, decision, junction, terminal, gap, risk, and out-of-scope tables.
7. Create `docs/DOCS/flow-maps/<feature-name>.flow-verification.md`.
8. Parse the generated SVG as XML when tooling is available, then verify round-trip consistency.
9. Set final implementation status:

```text
BLOCKED
```

when unresolved gap/undefined/owner-decision nodes remain.

```text
READY_FOR_OWNER_REVIEW
```

only when no unresolved `SPEC_GAP`, `UNDEFINED_BEHAVIOR`, or `OWNER_DECISION_REQUIRED` nodes remain.

```text
APPROVED
```

## Round-Trip Verification

Verify:

- Every metadata node exists as visible SVG content.
- Every visible node exists in metadata.
- Every edge has `from`, `to`, `condition`, `data-edge-type`, `flow`, and `data-source-ref`.
- Every decision node has explicit branches.
- Every junction connects at least two pipeline segments.
- Every gap node appears in the verification report.
- Every terminal state is explicit.
- No important logic exists only in geometry.
- SVG parses as XML when a parser is available.

## Flow Verification Report Format

Use this structure in `docs/flow-maps/<feature-name>.flow-verification.md`:

```md
# Flow Verification Report: <feature-name>

## Source
- Spec:
- Related docs:
- Generated SVG:

## Summary
- Total flows:
- Total nodes:
- Total edges:
- Decision nodes:
- Junction nodes:
- Terminal states:
- Spec gaps:
- Owner decisions required:
- Out-of-scope items:

## Implementation Status
BLOCKED or READY_FOR_OWNER_REVIEW

## Critical Gaps
| Gap ID | Type | Location | Why It Blocks Coding | Required Spec Fix |
|---|---|---|---|---|

## Decision Coverage
| Decision Node | Branches Found | Missing Branches | Status |
|---|---|---|---|

## Pipeline Handoffs
| Junction | From Flow | To Flow | Condition | Status |
|---|---|---|---|---|

## Terminal States
| Terminal State | Reached From | Condition | Status |
|---|---|---|---|

## Risk Notes
| Risk | Related Node/Edge | Severity | Recommendation |
|---|---|---|---|

## Final Verdict
State clearly whether implementation is BLOCKED or READY_FOR_OWNER_REVIEW.
```

## Acceptance Criteria

The skill run succeeds only when:

1. The SVG is valid XML.
2. The SVG has machine-readable metadata for nodes and edges.
3. The flow-map markdown matches the SVG.
4. The verification report is complete.
5. Every decision has explicit branches.
6. Every pipeline handoff uses a junction node.
7. Every terminal state is explicit.
8. Every unresolved ambiguity is marked as a gap.
9. No production code was modified.
10. Final status is `BLOCKED` or `READY_FOR_OWNER_REVIEW`.

Fail the run if any of those conditions are false, especially if an agent hides a spec gap or claims the spec is clear while unresolved gap nodes remain.

## Final Handoff Format

Return this format:

```md
# Spec Flow Map Handoff

## Files Created / Updated
- docs/flow-maps/<feature-name>.flow.svg
- docs/flow-maps/<feature-name>.flow-map.md
- docs/flow-maps/<feature-name>.flow-verification.md

## Status
BLOCKED or READY_FOR_OWNER_REVIEW

## Top Findings
1.
2.
3.

## Owner Action Required
- If BLOCKED: list the exact spec sections that need clarification.
- If READY_FOR_OWNER_REVIEW: ask Owner to review and approve the SVG flow map before implementation.
```
