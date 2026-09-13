# Anvien Agent Communication Evidence Ledger

## Metadata

- Date: `2026-09-12`
- Plan: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-plan.md`
- Rules: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-plan.md#rules`
- Evidence: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-evidence.md`
- Benchmark: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-benchmark.md`
- Actual status: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-actual-status.md`

## Evidence Rules

The evidence file explains why the work is known to be correct.

It should contain:

- metadata and companion files;
- evidence rules or evidence template;
- evidence sections such as `E0`, `E1`, or sections by phase/task;
- user report or problem evidence;
- source inspection, codebase facts, and document facts;
- commands run and pass/fail result;
- impact or blast-radius evidence when code/graph behavior changes;
- implementation evidence: files changed and behavior changed;
- validation evidence: build, tests, e2e, screenshots, or traces;
- failures encountered and how they were handled;
- detect-changes before commit;
- commit hash and closure evidence.

Evidence can reference short metric traces, but long metric tables belong in the benchmark file.

### Evidence ID Naming

Use stable, phase-scoped evidence IDs so `plan.md`, `actual-status.md`, `benchmark.md`, and later agents can reference exact proof without ambiguity.

Format:

```text
E<phase>-<item>-<kind><n>
```

Rules:

- `E<phase>` matches the plan phase number: `E0` for `P0`, `E1` for `P1`, `E2` for `P2`, and so on.
- `<item>` matches the checklist item without the dash: `P0A`, `P1A`, `P2B`.
- `<kind>` is plan-local. Choose a short uppercase token that is meaningful for this repo and this plan.
- `<n>` is a 1-based sequence number within that phase item and kind.
- Keep the same `<kind>` meaning stable inside one plan.
- Do not reuse an evidence ID for different facts.
- Reference exact evidence IDs from `actual-status.md` and `benchmark.md`; avoid referencing only broad section IDs such as `E1`.
- Use ranges such as `E0-P0A-FD1..E0-P0A-FD17` only for compact inventory summaries; use exact IDs when a specific status decision depends on a specific fact.
- If nearby plans already use a clear local evidence naming style, follow that style instead of inventing a new one.

Examples only:

- `E0-P0A-SRC1`
- `E0-P0A-GRAPH1`
- `E1-P1A-ROUTE1`
- `E2-P2B-KEYBOARD1`
- `E2-P2B-DETECT1`

Evidence sections must follow the plan phases:

- `E0` corresponds to `P0`.
- `E1` corresponds to `P1`.
- `E2` corresponds to `P2`.
- Use exact evidence IDs inside each section, not broad section IDs as proof.
- Each evidence section must name the plan phase or checklist item it supports.
- Do not invent fixed evidence categories; record the evidence required by the matching plan phase.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

### E0-P0A-SRC1
Đọc báo cáo reports/problem/2026-09-11-anvien-multi-provider-workspace-coordinator-discussion.md.
Owner xác định Coordinator là tool không AI; role gắn agent/provider/model, UI là không gian chung.

### E0-P0A-SRC2
internal/session/controller.go:15 Adapter hiện có RunChat; :24 controller giữ một adapter; :117-124 chat mới cancel chat đang chạy cùng repo.
internal/session/types.go:85 ChatRequest chỉ có repoName/repoPath/message; không có model/agent/session target.
internal/session/codex.go:35 cmd.Run thu stdout rồi :232 mới parse. :198 chạy codex exec; chưa dùng durable conversational session tại đường code này.
internal/httpapi/session.go:104 disconnect browser gọi CancelSession.
internal/session/job.go:34 history là memory slice.
Đây là quan sát source, không phải kết quả runtime test.

### E0-P0A-SRC3
anvien-web/src/hooks/chat-runtime/ChatRuntimeContext.tsx:43 chatMessages trong React state; :83 ép activeProvider codex.
ChatPanel.tsx dùng ChatTranscript/ChatComposer; transcript có MarkdownRenderer/ToolCallCard.
Có ứng viên tái sử dụng phần render, không đồng nghĩa lifecycle dùng được nguyên trạng.

### E0-P0A-GRAPH1
anvien/bin/anvien.exe analyze --force hoàn tất exit 0: 2465 scanned, 897 parsed_code, failed=0;
graph 170912 nodes / 222544 relationships. Command chạy sau CLI PATH failure; dùng binary repo, không cài lại global.
### E0-P0A-IMPACT1
impact file internal/session/controller.go --repo E:\Anvien --direction upstream --json: CRITICAL;
45 contained symbols, 10 direct, 8 affected files, 1 affected flow, 4 linked flows, 2 linked tests.
Affected paths: cmd/anvien/main.go; internal/cli/command.go; internal/httpapi/listen.go;
internal/httpapi/server.go; internal/httpapi/session.go; internal/session/codex.go;
internal/session/controller.go; reports/Investigation/rp_investigation_260823_103434_by_gpt-5_analyze_performance_root_cause.md.
Report path is graph evidence, not editable code scope.
### E0-P0A-FD1
file-detail internal/session/controller.go --json completed. Full direct tool output exceeded display budget.
Observed relationship counts local=49 outbound=37 inbound=15; unresolved=81; linked flows=4 tests=2.
Do not claim full relationship review or finalized P0; recover complete graph evidence before implementation.
### E0-P0A-RULE1
Owner explicitly removes separate rules.template.md/rules.md requirement; rules are embedded in plan.
No production code changes, no full build, no runtime QA, no provider jobs dispatched.


## E1 - P1 Evidence

Matching plan item(s): `P1-A`

Chưa triển khai.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`

Chưa triển khai.

## Closure Evidence

Use this section for final detect-changes, commit hash, and closure evidence when the plan reaches completion.

Draft, chưa closure/commit. Không có claim runtime PASS.
