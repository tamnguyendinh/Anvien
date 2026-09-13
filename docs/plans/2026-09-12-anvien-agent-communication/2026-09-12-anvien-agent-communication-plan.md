# Anvien Agent Communication Plan

## Metadata

- Date: `2026-09-12`
- Status: `draft`
- Plan: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-plan.md`
- Evidence: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-evidence.md`
- Benchmark: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-benchmark.md`
- Actual status: `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-actual-status.md`

## Goal

Biến báo cáo thảo luận 2026-09-11 thành kiến trúc triển khai có căn cứ: Coordinator thuần vận chuyển, UI hội thoại chung, kích hoạt đúng agent. Trạng thái draft/P0 incomplete; chưa cho phép implementation.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
- Anvien only helps localize candidate boundaries; relationships and blast radius must be cross-checked against imports, call paths, and actual code before deciding scope.
- Run Anvien detect-changes before every implementation-slice commit when implementation work was performed.
- For public runtime or UI-facing changes, validate the real user-visible runtime with browser or Playwright evidence.
- For app/runtime validation, full build must include Docker image/container build. If Docker is missing or not run, full build is incomplete.
- Any Playwright validation must target the real built Docker/container runtime. Running Playwright against a host dev server, framework dev mode, mocked server, or source-run shortcut is not valid runtime evidence.
- If the Docker runtime cannot be built or started, the slice/plan is blocked; do not replace it with dev-server Playwright evidence.
- Playwright evidence must record the Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
- Keep the standard planner structure. These detail rules only make phase checklist items concrete enough to implement safely.
- Every implementation phase must be decomposed into multiple implementation slices that are as small as practical. A phase is a grouping and ordering container; a slice is the executable implementation unit.
- Do not implement a phase directly. Work starts from a slice ID such as `P1-A`, `P1-B`, or `P2-C`.
- Prefer many narrow slices over one broad slice. A single-slice implementation phase is allowed only when the plan explicitly states why the phase cannot be split further without creating empty or non-executable slices.
- Each implementation slice must include:
  + Goal
  + Scope Boundary
  + Non-Goals when useful
  + Pre-flight Questions
  + Work Steps (must act as a strict "control plane" following the 6-step execution lifecycle: Code → Inspect → Update tests → Build → QA → Acceptance/Commit. Step 1 (Code) must be decomposed into leaf atomic tasks, where each leaf task defines: Action, Inputs, Allowed Edit Scope, Outputs, and Verification Condition).
  + Implementation Gate
  + Acceptance
  + Evidence Targets
  + Actual-status Update
  + Commit Boundary
- Every slice's Work Steps must strictly follow the 6-step execution lifecycle:
  1. Code: Execute leaf atomic tasks in core source files only. Each leaf atomic task must explicitly specify Action, Inputs, Allowed Edit Scope, Outputs, local Verification Condition, and a Task Checkpoint with explicit prefix (e.g. wip(<slice_id>): task 1.x - <verified task outcome>; micro WIP commit for granular task progress, NOT a slice completion commit). Do not touch test files during this step.
  2. Code Inspection & Lint: Review diff, syntax, linter, typecheck, and ensure edits remain strictly within assigned boundary.
  3. Update Tests: Following Master Rule 6 (Code first, test second), add or update unit/integration tests to verify the newly implemented code behavior.
  4. Build: Run the repository's full build to ensure type contracts, package bundles, and binaries compile cleanly.
  5. QA (Mini QA): Verify user-visible runtime behavior (browser, desktop app, API, CLI, or Playwright evidence).
  6. Slice Acceptance & Official Milestone Commit: Verify slice acceptance criteria, record evidence in evidence.md, refresh actual-status.md, and create the official slice completion commit (e.g. feat(<slice_id>): <slice outcome>).
- Split planned work into separate slices when it contains more than one primary user-visible behavior, user trigger, render location, permission or visibility rule, DB write target, DB state transition, API/CLI/MCP contract, async/event/webhook flow, external side effect, cleanup/quarantine domain, behavior test target, independent acceptance gate, or independent commit boundary.
-  Hidden fallback is forbidden. Prefer a visible failure over a fallback that hides a broken primary path.
- When touching DB-backed content, verify the full loop when applicable: UI input -> submit action -> DB write -> DB read after reload/new request -> correct UI render or omission. If there is no UI, replace UI steps with the real caller/consumer flow.
- Tests must prove product behavior. Delete or replace tests that only assert implementation details, helper output, static DOM existence, or mocked plumbing without proving trigger -> process -> observable result.
- If a planned item uses wording such as `and`, `also`, `then wire`, `plus update`, `both`, or `handle all`, check whether it is actually multiple slices.
- Do not write broad actionable items such as `Implement checkout, webhook, entitlement update, and billing UI`; split them into narrow slices such as `Create checkout session request`, `Persist checkout session state`, `Handle provider webhook`, `Update entitlement from webhook event`, and `Render billing status from entitlement`.
- Each slice work step must include UI flow, DB/data flow, render location, and evidence target checks. Use `N/A` with a reason when a check does not apply.
- If tests write DB rows, app state, files, queues, provider state, or other persistent data, the slice must define cleanup or quarantine before implementation.

## Problem

Source hiện tại là single-Codex chat theo repo, chưa đáp ứng nhiều agent liên lạc và kích hoạt qua shared conversation.

## Scope

Khảo sát internal/session, internal/httpapi, internal/mcp, contracts và anvien-web chat. Chỉ viết bộ plan; production inspect-only.

## Non-Goals

Không sửa code, không start agent/provider, không xây editor kiểu VS Code, không AI Coordinator, không giới hạn context chung, không thay provider bằng API khác.

## Requirements

Giữ nguyên report gốc. Owner phân vai. Payload không bị diễn giải. UI và activation là hai consumer độc lập. Rules nằm trong plan theo chỉ dẫn Owner; không tạo rules.md.

## Acceptance Criteria

P0 đủ source/graph/protocol evidence trước implementation; mỗi slice có boundary và chứng minh trigger -> delivery -> activation -> reply -> UI. Không tuyên bố hệ thống đã hoạt động từ source inspection.

## Checklist

- [ ] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-09-12-anvien-agent-communication/2026-09-12-anvien-agent-communication-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.

## Hướng triển khai đề xuất — chưa mở implementation
### P1 — Contract và identity
- Agent ID độc lập role/model; role binding được resolve và đóng dấu agent đích tại lúc nhận tin.
- Chốt envelope, reply routing, publication vs activation, terminal message vs stream delta.
- Owner chọn runtime/model; không thay Gemini qua Antigravity bằng Gemini API hoặc CLI khác một cách ngầm định.
### P2 — Communication core
- Lưu message và outbox trong cùng transaction; delivery riêng cho từng recipient.
- Một nguồn dữ liệu cho UI và dispatcher, không dùng UI làm đường truyền giữa agent.
- De-duplicate ingress; uncertain external delivery không retry mù; acknowledgement có ngữ nghĩa rõ.
### P3 — Session activation và adapters
- Mailbox theo agent/session, không khóa toàn repo.
- Codex adapter giữ provider session ID, kích hoạt lượt mới/đưa tin giữa lượt theo capability.
- Antigravity adapter là deliverable tích hợp cần xây; giao thức cụ thể phải được chứng minh, không thay provider ngoài ý Owner.
- Nội dung final chỉ publish một lần; telemetry/tool progress không tự kích hoạt CEO.
### P4 — API/MCP và web hội thoại chung
- MCP gửi tin và HTTP gửi tin đi vào cùng service; mỗi caller có identity, không tin sender do client tự khai.
- Nếu MCP stdio chạy process riêng, forward về serve đang sở hữu workspace; không tạo memory store thứ hai.
- UI chọn agent/provider/model/role; conversation transcript hiển thị source/recipient/delivery.
- Subscribe event có cursor, reload nhận lại lịch sử; UI disconnect không cancel agent.
### P5 — Lane và quyền quản lý đội
- Tạo lane/gán agent/tạo agent là tool riêng; Coordinator không phân tích văn bản để làm thay CEO.
- Role reassignment không âm thầm chuyển tin đang queued sang agent khác.
### P6 — Code intelligence, artifact và Telegram
- Graph là công cụ agent dùng, không điều kiện bắt buộc của vận chuyển.
- Artifact references gắn workspace/revision; Telegram dùng chung ingress/store.
### P7 — Acceptance toàn vòng
- Owner -> Gemini/Antigravity CEO -> GPT/Codex coder -> CEO -> reviewer -> Owner.
- Đóng UI vẫn tiếp tục; reload có đủ hội thoại; restart không mất tin đã accepted.
- Thử hai agent cùng model, tin đến khi busy, lost ack, retry, permission bypass.
- Không nghiệm thu multi-provider chỉ bằng mock adapter.

## Điều kiện trước khi phân rã implementation slices
Các P1..P7 phía trên là dependency roadmap, KHÔNG phải slice được phép code.
P0 phải chốt protocol, persistence, schema contract, UI states và impact từng điểm tích hợp.
Sau đó thay từng phase bằng các mini-plan đúng template: Goal, Scope, Pre-flight,
atomic Code tasks, Inspect, Tests, Full Build, QA, Acceptance/Commit.
Không tạo task leaf giả khi chưa biết protocol hoặc file boundary.
Chưa có thay đổi production, chưa có acceptance hoặc ủy quyền tự chạy provider.

## Risk Notes
- Controller hiện tại CRITICAL (8 affected files); graph chỉ là chỉ dẫn phạm vi, không lệnh cấm.
- Runtime hiện tại auto-cancel theo repo và disconnect browser không phù hợp shared conversation.
- Chưa có bảo đảm exactly-once end-to-end nếu provider không có idempotent submit/reconciliation.
- Context/provider-native compact không thuộc Coordinator.
- Antigravity runtime protocol, DB implementation và migration chưa được chốt.
- Preserve code-intelligence routes và chat cũ cho tới khi có migration authority.
