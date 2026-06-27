---
name: spec-to-svg-flow-map
description: Convert product, feature, UI, backend, auth, sync, lifecycle, external-contract, or multi-branch specs into detailed semantic SVG flow maps with machine-readable metadata, source-union coverage, flow-by-flow rendering, no-collapse checks, gap detection, verification reports, and BLOCKED/READY_FOR_OWNER_REVIEW status. Use when the user asks to turn a spec into a visual flow map, audit spec completeness before coding, compare against reference diagrams, expose undefined behavior, or identify owner decisions before implementation.
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
- legend

The SVG must help humans and coding agents decide whether the spec is clear enough for safe implementation.

The SVG must render the source as written, not the implementation that the agent thinks the system should have.

## Detail Completeness Objective

Generate a flow map that is more detailed than any existing source diagram or reference artifact provided for the same scope.

Do not create an overview-only map when the source contains implementation roads, command families, lifecycle states, storage checkpoints, guards, cursors, locks, receipts, versions, hash chains, terminal states, or recovery paths.

The output must preserve the union of:

- source spec details
- related authority docs
- existing flow maps or SVGs
- verification reports
- UI/state/API/contract maps
- explicitly named invariants and negative rules

If a prior/reference SVG contains a named flow, branch, state, store, cursor, command, invariant, or terminal state that is not represented in the new SVG, the run fails.

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
3. Existing/reference SVGs and flow maps, if present.
4. Existing verification reports, if present.
5. UI prototype or UI slot map, if present.
6. State/source map, if present.
7. Backend/API/contract map, if present.
8. Actual wiring status, if present.
9. Project rules.

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
12. Build source-union coverage before drawing implementation logic.
13. Render implementation logic flow-by-flow, not from one bulk summary pass.
14. Do not collapse named implementation details into generic nodes.
15. Do not improve, repair, normalize, or connect flows beyond what the source actually states.
16. If the map cannot be completed safely, stop and report why.

If the map contains any unresolved node with type `SPEC_GAP`, `UNDEFINED_BEHAVIOR`, or `OWNER_DECISION_REQUIRED`, set implementation status to `BLOCKED`.

If `missing_source_items` is not empty, set implementation status to `BLOCKED`.

## As-Described Fidelity Rule

Draw the spec exactly as the source artifacts describe it, even when the description is incomplete, wrong, disconnected, inconsistent, or unsafe.

Do not draw the ideal lifecycle, expected architecture, likely implementation order, or missing runtime handoff unless a source artifact explicitly states it.

If the source describes disconnected flows, draw disconnected flows. A connected lifecycle graph is required only when the source describes those connections.

If the source says flow A leads to flow B, draw that relationship with the source reference.

If the source does not say flow A leads to flow B, do not connect them. If that missing relation blocks safe implementation, create a visible gap node such as `UNDEFINED_HANDOFF`, `SPEC_GAP`, or `OWNER_DECISION_REQUIRED`.

If the source describes a wrong order, unsafe branch, missing denial path, or broken recovery path, preserve that wrong or broken shape in the SVG and mark the defect as a gap or risk. Do not silently correct it.

Every implementation-logic node, edge, junction, terminal, and gap must trace to a source reference or to an explicit missing-source/gap record. Unreferenced implementation logic is treated as speculation and fails the run.

Layout groups, legends, lane labels, and navigation anchors may be derived, but they must not add implementation meaning.

## Pre-Render Source-Described Flow Manifest Gate

After building the source union inventory and before drawing any SVG implementation logic, create a source-described flow manifest from the inventory.

The manifest is not a design plan. It is a list of flows, subflows, and roads explicitly described by the source artifacts.

For each manifest item, record:

```text
flow_id
source_name
category: flow | subflow | road
source_refs
described_entry
described_exit_or_terminal
described_handoffs
missing_handoffs_or_unknown_relations
mapping_status: PENDING | MAPPED | MAPPED_AS_GAP | INTENTIONALLY_OUT_OF_SCOPE | MISSING
planned_or_actual_svg_group_id
```

Rules:

1. Every source-described flow, subflow, and road must appear in the manifest before it is drawn.
2. Do not add expected flows that are not described by the source. If an expected flow is necessary but absent, add a gap item, not a real flow.
3. Do not draw a flow that is not in the manifest.
4. If a new source-described flow is discovered while drawing, stop drawing, update the source union inventory and manifest, then resume flow-by-flow rendering.
5. A manifest item is complete only when it maps to visible SVG ids, a visible gap, or an intentional out-of-scope record.
6. Fail the run if a source-described flow is absent from the manifest.
7. Fail the run if a visible implementation flow exists without a manifest item.

## Source Union Inventory Gate

Before drawing, build a source union inventory from all input artifacts.

Do not treat document titles, section titles, or broad flow IDs as sufficient inventory. Each implementation-relevant command, guard, store, cursor, terminal state, recovery path, invariant, negative rule, and described relation must be inventoried as its own source item.

Inventory these categories:

```text
flows
subflows
roads
actors
roles
scopes
screens
UI commands
runtime commands
backend/API commands
IPC commands
permissions
guards
branch conditions
deny conditions
stores
caches
files
DBs
table families
queues
outboxes
projections
snapshots
cursors
checkpoints
receipts
versions
locks
hash chains
coverage markers
lifecycle states
retry rules
rollback rules
terminal states
recovery paths
invariants
negative rules
out-of-scope boundaries
```

For each inventory item, assign one mapping status:

```text
MAPPED_AS_NODE
MAPPED_AS_EDGE
MAPPED_AS_JUNCTION
MAPPED_AS_TERMINAL
MAPPED_AS_GAP
INTENTIONALLY_OUT_OF_SCOPE
MISSING
```

Fail the run if any implementation-relevant item is `MISSING`.

Record the mapped SVG ids for every item that is not intentionally out of scope.

## Flow-By-Flow Rendering Rule

Do not draw all flows in one batch.

Use this loop:

1. Select exactly one named flow or subflow from the source union inventory.
2. Re-read the source spec sections and reference artifacts for that flow.
3. Extract that flow's actors, entry points, preconditions, commands, guards, branches, state/data reads, writes, side effects, terminal states, recovery paths, handoffs, and invariants.
4. Draw only that flow.
5. Add that flow's visible SVG nodes, edges, junctions, terminals, and gaps.
6. Add matching metadata for that flow.
7. Verify that flow against the source inventory before moving to the next flow.
8. Mark each covered source item with its mapped node/edge/junction/terminal/gap id.
9. Only then continue to the next flow.

Do not rely on memory from earlier reading when rendering a later flow. Re-open or re-read the relevant source sections before each flow.

The final SVG may contain all flows, but the construction process must be flow-by-flow.

## No Bulk Drawing Rule

Never render multiple independent pipelines from a single high-level summary pass.

A bulk drawing pass is allowed only for:

- the legend
- global lane layout
- cross-flow index
- high-level navigation anchors
- final consistency cleanup

It is not allowed for implementation logic.

Implementation logic must be added through the Flow-By-Flow Rendering Rule.

## No Collapse Rule

Do not collapse named details into generic nodes.

A node such as `BUSINESS_OPS`, `SYNC_ENGINE`, `AUTH_FLOW`, `REPORTS`, `SETTINGS`, or `LOCAL_COMMAND` is allowed only as an index, grouping label, or junction. It is not sufficient coverage for detailed source behavior.

If the source names roads such as POS order, pay/refund, move/merge/split, shift/cash, inventory, owner setup, report coverage, local print, snapshot bootstrap, manual sync, lifecycle reconnect, or recovery, each road must be represented separately.

Every generic/index node must fan out to detailed nodes or detail sheets.

Fail the run if a generic node replaces a named implementation road.

## Minimum Detail Per Flow

For every named flow or subflow, include at least:

1. Entry trigger.
2. Actor, role, device, and scope preconditions.
3. UI, external, runtime, IPC, or backend command source.
4. Authority or permission gate.
5. State/data reads.
6. Decision branches.
7. Write/apply target.
8. Side effects.
9. Async, outbox, background, or transport behavior.
10. Success terminal state.
11. Failure, denied, blocked, pending, or rollback terminal state.
12. Recovery or retry path.
13. Pipeline handoff junctions.
14. Source-of-truth boundary.
15. Explicit invariants and "must not" rules.

If any required item is not defined by the source, create a visible gap node.

## Required Domain Detail Checklist

When relevant, explicitly model these domains instead of grouping them:

- Auth/session/device/entitlement restore, login handoff, cache, keyring, expiry, logout, denied states.
- Restaurant/scope selection, visible/bound scopes, DB mount, hydrated/not-hydrated states.
- Permission/local command gate, renderer boundary, Go/backend authority, command guards.
- Owner/app setup, setup outbox, setup version, setup receipt, convergence state.
- Business runtime: POS order, pay/refund, move/merge/split, shift/cash, inventory/stocktake.
- Sync transport: LAN/WSS/VPS, relay, delta, dedupe, ack, cursor, hash verify, gap, repair.
- Snapshot/bootstrap/manual sync: baselines, manifest, anchors, allowlist apply, rollback, catchup.
- Reports/coverage: aggregate/detail split, retention, source coverage, export/print gating.
- Local print/export: printer config, preview, spooler result, local-only boundary.
- Lifecycle/reconnect: active/idle/sleep/offline/resume, missed relay, auth refresh, cursor catchup.
- Local settings/device-only behavior.
- External/backend contracts and denied/error responses.

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

## SVG Metadata And Source Coverage Metadata

Every SVG must include one machine-readable metadata block.

The metadata must include graph content and source coverage:

```xml
<metadata id="spec-flow-map">
{
  "feature": "<feature-name>",
  "source_spec": "<source-spec-path>",
  "version": "draft",
  "reference_artifacts": [
    {
      "path": "docs/flow-maps/reference.flow.svg",
      "type": "reference_svg",
      "scope": "<scope>"
    }
  ],
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
  "flow_manifest": [
    {
      "flow_id": "AUTH_FLOW",
      "source_name": "Auth login",
      "category": "flow",
      "source_refs": ["spec.md#auth-login"],
      "described_entry": "User submits login",
      "described_exit_or_terminal": "AUTH_SUCCESS or AUTH_DENIED",
      "described_handoffs": ["AUTH_SUCCESS -> LICENSE_CHECK"],
      "missing_handoffs_or_unknown_relations": [],
      "mapping_status": "MAPPED",
      "planned_or_actual_svg_group_id": "flow-auth-flow"
    }
  ],
  "source_inventory": [
    {
      "source_item": "sync_cursor.last_pulled_relay_id",
      "category": "cursor",
      "source_ref": "spec.md#sync-reconnect",
      "mapping_status": "MAPPED_AS_NODE",
      "mapped_ids": ["node-sync-cursor-last-pulled-relay-id"]
    }
  ],
  "coverage_summary": {
    "total_source_items": 0,
    "mapped_items": 0,
    "missing_items": 0,
    "collapse_violations": 0
  },
  "missing_source_items": [],
  "collapse_violations": [],
  "gaps": [],
  "terminal_states": [],
  "junctions": [],
  "legend": {
    "visible_legend_id": "legend-end",
    "shape_meaning": {},
    "color_meaning": {},
    "arrow_meaning": {}
  }
}
</metadata>
```

Metadata must match the visible SVG nodes and edges.

If `missing_source_items` is not empty, set implementation status to `BLOCKED`.

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

## Visible End Legend Requirement

Every generated SVG must end with a visible, human-readable legend section.

Place the legend as the last visible SVG group before `</svg>`:

```xml
<g
  id="legend-end"
  data-type="legend"
  data-position="end"
>
  <title>Legend</title>
  <desc>Explains node shapes, colors, line styles, and arrow types used in this SVG.</desc>
  ...
</g>
```

The legend must be readable without opening the markdown report. It must explain:

1. What each node shape/box means.
2. What each color means.
3. What each line style means.
4. What each arrow type means.
5. What error, recovery, gap, owner-decision, out-of-scope, terminal, junction, async, sync, data-flow, and control-flow markers mean.

The legend is required even when the SVG is large or split into multiple visual regions.

Do not put the legend only in metadata. The metadata must reference the visible legend id, but the SVG canvas itself must contain the readable legend.

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
13. A source union inventory item is implementation-relevant but lacks a mapped SVG id.
14. A reference artifact contains a named detail that is absent from the new map.
15. A flow lacks a required minimum detail item and the source does not define it.

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

## Updated Workflow

1. Read project rules and requested scope.
2. Read source spec and related authority docs.
3. Read existing/reference SVGs, flow maps, verification reports, UI maps, state maps, API maps, and contract maps when present.
4. Build the source union inventory.
5. Build the source-described flow manifest from inventory items that are explicitly present in the sources.
6. Identify named flows and subflows from the manifest.
7. Create a global lane/layout plan only; do not draw implementation logic yet.
8. For each named flow, repeat:

```text
re-read relevant source sections
extract flow-local inventory
confirm the flow is already present in the source-described flow manifest
draw that flow only
write visible nodes/edges/junctions/terminals/gaps
write matching metadata
verify source coverage for that flow
mark mapped source inventory items
mark the manifest item with mapped SVG ids or gap ids
```

9. After all flows are drawn, run global consistency verification.
10. Check for missing source items.
11. Check for flow manifest omissions or unmanifested visible flows.
12. Check for collapse violations.
13. Check reference SVG delta.
14. Parse SVG as XML.
15. Verify metadata/visible node round-trip.
16. Set implementation status.

Create `docs/flow-maps/<feature-name>.flow.svg` as semantic XML.

Create `docs/flow-maps/<feature-name>.flow-map.md` with:

```text
Feature Name
Source Spec
Reference Artifacts
Source Union Inventory
Source-Described Flow Manifest
Flow List
Lane List
Node Table
Edge Table
Decision Table
Junction Table
Terminal State Table
Gap Table
Risk Table
Out-of-Scope Table
Reference SVG Delta
Collapse Violation Table
Generic Node Fan-Out Audit
```

Create `docs/flow-maps/<feature-name>.flow-verification.md`.

Before closing the SVG, add the visible end legend as the last visible group and ensure its id matches `metadata.legend.visible_legend_id`.

Set final implementation status:

```text
BLOCKED
```

when unresolved gaps, undefined behavior, owner decisions, missing source items, collapse violations, or failed reference deltas remain.

```text
READY_FOR_OWNER_REVIEW
```

only when no unresolved `SPEC_GAP`, `UNDEFINED_BEHAVIOR`, `OWNER_DECISION_REQUIRED`, source coverage, collapse, or reference-delta blockers remain.

Only the Owner can approve implementation after reviewing the generated flow map.

## Round-Trip Verification

Verify:

- Every metadata node exists as visible SVG content.
- Every visible node exists in metadata.
- Every edge has `from`, `to`, `condition`, `data-edge-type`, `flow`, and `data-source-ref`.
- Every implementation-logic node, edge, junction, terminal, and gap has a source reference or an explicit gap source record.
- Every source-described flow appears in the flow manifest before it appears in visible SVG content.
- Every visible implementation flow has a matching flow manifest item.
- Missing source-described relations are represented as gaps, not invented edges.
- Every decision node has explicit branches.
- Every junction connects at least two pipeline segments.
- Every gap node appears in the verification report.
- Every terminal state is explicit.
- Every source inventory item has a final mapping status.
- Every implementation-relevant source inventory item has mapped SVG ids or a visible gap.
- Every reference SVG/detail delta is represented as a mapped id or gap.
- Every generic/index node has allowed purpose and fan-out.
- Every flow records which source sections were re-read before rendering.
- The visible end legend exists as the last visible SVG group, is readable, and explains shapes, colors, line styles, and arrow types.
- The metadata `legend.visible_legend_id` matches the visible end legend group id.
- No important logic exists only in geometry.
- SVG parses as XML when a parser is available.

## Flow Verification Report Format

Use this structure in `docs/flow-maps/<feature-name>.flow-verification.md`:

```md
# Flow Verification Report: <feature-name>

## Source
- Spec:
- Related docs:
- Reference artifacts:
- Generated SVG:

## Summary
- Total flows:
- Flow manifest items:
- Total nodes:
- Total edges:
- Source inventory items:
- Mapped source items:
- Missing source items:
- Collapse violations:
- Decision nodes:
- Junction nodes:
- Terminal states:
- Spec gaps:
- Owner decisions required:
- Out-of-scope items:

## Implementation Status
BLOCKED or READY_FOR_OWNER_REVIEW

## Source Union Inventory Coverage
| Source Item | Category | Source Ref | Mapping Status | Mapped SVG IDs |
|---|---|---|---|---|

## Source-Described Flow Manifest
| Flow ID | Source Name | Category | Source Refs | Described Entry | Described Exit/Terminal | Described Handoffs | Missing Handoffs/Unknown Relations | Mapping Status | SVG Group/Gap IDs |
|---|---|---|---|---|---|---|---|---|---|

## Flow-By-Flow Coverage
| Flow | Source Sections Re-read | Nodes | Edges | Terminals | Gaps | Status |
|---|---|---:|---:|---:|---:|---|

## Missing Source Items
| Source Item | Category | Source Ref | Why It Matters | Required Fix |
|---|---|---|---|---|

## Collapse Violations
| Generic Node | Missing Detail | Source Ref | Required Split |
|---|---|---|---|

## Reference SVG Delta
| Reference Artifact | Item Present In Reference | Present In New SVG | Mapped ID / Gap |
|---|---|---|---|

## Generic Node Fan-Out Audit
| Generic Node | Allowed Purpose | Fan-Out Nodes | Status |
|---|---|---|---|

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
3. The SVG metadata includes source coverage fields.
4. The flow-map markdown matches the SVG.
5. The verification report is complete.
6. Every decision has explicit branches.
7. Every pipeline handoff uses a junction node.
8. Every terminal state is explicit.
9. Every unresolved ambiguity is marked as a gap.
10. Every implementation-relevant source inventory item is mapped or represented as a gap.
11. Every flow is rendered through the Flow-By-Flow Rendering Rule.
12. Every source-described flow appears in the pre-render flow manifest before visible SVG implementation logic is drawn.
13. Every visible implementation flow has a matching manifest item.
14. Every implementation-logic node, edge, junction, terminal, and gap is source-referenced or gap-recorded.
15. The SVG preserves source-described wrong, missing, disconnected, or contradictory flow shapes instead of correcting them by inference.
16. Every generic/index node has a legitimate purpose and fans out to details.
17. The SVG contains a visible end legend that explains node boxes/shapes, colors, line styles, arrow types, and special markers.
18. The visible end legend is the last visible SVG group and is referenced from metadata.
19. No production code was modified.
20. Final status is `BLOCKED` or `READY_FOR_OWNER_REVIEW`.

Fail the run if any of those conditions are false.

## Acceptance Criteria Additions

Fail the run if:

1. A reference/source SVG contains a named flow or detail missing from the new SVG.
2. A generic node replaces a named implementation road.
3. A source inventory item is marked `MISSING`.
4. A decision branch, denial state, recovery path, cursor/checkpoint, or source-of-truth boundary is only described in prose and not represented in SVG metadata.
5. The new SVG is less detailed than any existing diagram for the same scope.
6. Multiple independent implementation pipelines were drawn in one batch without flow-by-flow re-reading and verification.
7. A flow was rendered without recording which source sections were re-read for that flow.
8. A flow was marked complete before its source inventory items were mapped.
9. An agent hides a spec gap or claims the spec is clear while unresolved gap nodes remain.
10. The SVG lacks a readable visible legend at the end, or the legend does not explain the meaning of node boxes, colors, line styles, and arrow types.
11. A source-described flow, subflow, or road is drawn before it appears in the flow manifest.
12. A visible implementation flow exists without a flow manifest item.
13. An edge, junction, or lifecycle handoff is added because it is expected or architecturally desirable but not source-described.
14. A missing relation is drawn as a real edge instead of a visible gap.
15. The agent corrects, normalizes, or reconnects a flawed source flow instead of rendering the flaw and marking it as a gap or risk.

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
- If BLOCKED: list the exact spec sections, source inventory items, reference deltas, or collapse violations that need clarification.
- If READY_FOR_OWNER_REVIEW: ask Owner to review and approve the SVG flow map before implementation.
```
