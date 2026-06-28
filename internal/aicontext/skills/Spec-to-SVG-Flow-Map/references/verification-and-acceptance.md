# Verification Và Acceptance

## Verification Round-Trip

Verify:

- Mọi metadata node tồn tại như visible SVG content.
- Mọi visible node tồn tại trong metadata.
- Mọi edge có `from`, `to`, `condition`, `data-edge-type`, `flow`, và `data-source-ref`.
- Mọi implementation-logic node, edge, junction, terminal, và gap có source reference hoặc explicit gap source record.
- Mọi source-described flow xuất hiện trong flow manifest trước khi xuất hiện trong visible SVG content.
- Mọi visible implementation flow có matching flow manifest item.
- Missing source-described relations được represent là gaps, không phải invented edges.
- Mọi decision node có explicit branches.
- Mọi junction connect ít nhất hai pipeline segments.
- Mọi gap node xuất hiện trong verification report.
- Mọi terminal state explicit.
- Mọi source inventory item có final mapping status.
- Mọi implementation-relevant source inventory item có mapped SVG ids hoặc visible gap.
- Mọi reference SVG/detail delta được represent là mapped id hoặc gap.
- Mọi generic/index node có allowed purpose và fan-out.
- Mọi flow ghi source sections nào đã re-read trước khi render.
- Visible end legend tồn tại như visible SVG group cuối cùng, đọc được, và giải thích shapes, colors, line styles, arrow types.
- Metadata `legend.visible_legend_id` khớp visible end legend group id.
- Không có important logic chỉ tồn tại trong geometry.
- SVG parse được như XML khi có parser.

## Định Dạng Flow Verification Report

Dùng structure này trong `docs/flow-maps/<feature-name>.flow-verification.md`:

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

## Nội Dung Flow-Map Markdown

`docs/flow-maps/<feature-name>.flow-map.md` phải có:

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

## Tiêu Chí Chấp Nhận

Skill run chỉ succeed khi:

1. SVG là valid XML.
2. SVG có machine-readable metadata cho nodes và edges.
3. SVG metadata có source coverage fields.
4. Flow-map markdown khớp SVG.
5. Verification report complete.
6. Mọi decision có explicit branches.
7. Mọi pipeline handoff dùng junction node khi source mô tả handoff đó.
8. Mọi terminal state explicit.
9. Mọi unresolved ambiguity được mark là gap.
10. Mọi implementation-relevant source inventory item được mapped hoặc represented as gap.
11. Mọi flow được render theo Flow-By-Flow Rendering Rule.
12. Mọi source-described flow xuất hiện trong pre-render flow manifest trước khi visible SVG implementation logic được vẽ.
13. Mọi visible implementation flow có matching manifest item.
14. Mọi implementation-logic node, edge, junction, terminal, và gap source-referenced hoặc gap-recorded.
15. SVG preserve source-described wrong, missing, disconnected, hoặc contradictory flow shapes thay vì correct bằng inference.
16. Mọi generic/index node có legitimate purpose và fan out tới details.
17. SVG có visible end legend giải thích node boxes/shapes, colors, line styles, arrow types, và special markers.
18. Visible end legend là visible SVG group cuối cùng và được reference từ metadata.
19. Không sửa production code.
20. Final status là `BLOCKED` hoặc `READY_FOR_OWNER_REVIEW`.

Fail run nếu bất kỳ condition nào sai.

## Tiêu Chí Fail Bổ Sung

Fail run nếu:

1. Reference/source SVG có named flow hoặc detail bị thiếu trong new SVG.
2. Generic node thay thế named implementation road.
3. Source inventory item bị mark `MISSING`.
4. Decision branch, denial state, recovery path, cursor/checkpoint, hoặc source-of-truth boundary chỉ được mô tả trong prose và không represented trong SVG metadata.
5. New SVG ít chi tiết hơn bất kỳ existing diagram nào cho cùng scope.
6. Nhiều independent implementation pipelines được vẽ trong một batch mà không flow-by-flow re-reading và verification.
7. Flow được render mà không record source sections đã re-read.
8. Flow được mark complete trước khi source inventory items của nó được mapped.
9. Agent che gap spec hoặc claim spec rõ trong khi unresolved gap nodes vẫn còn.
10. SVG thiếu readable visible legend ở cuối, hoặc legend không giải thích node boxes, colors, line styles, arrow types.
11. Source-described flow, subflow, hoặc road được vẽ trước khi xuất hiện trong flow manifest.
12. Visible implementation flow tồn tại mà không có flow manifest item.
13. Edge, junction, hoặc lifecycle handoff được add vì expected/architecturally desirable nhưng không source-described.
14. Missing relation được vẽ thành real edge thay vì visible gap.
15. Agent correct, normalize, hoặc reconnect flawed source flow thay vì render flaw và mark gap/risk.
