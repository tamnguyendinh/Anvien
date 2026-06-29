# File Detail Compact Full Detail Plan

## Metadata

- Date: `2026-06-29`
- Status: `P0 complete - supplemented for implementation readiness`
- Plan: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-plan.md`
- Evidence: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-evidence.md`
- Benchmark: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-benchmark.md`
- Actual status: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-actual-status.md`

## Goal

Change `file-detail` so it remains the full detail surface for everything related to one file, but emits a compact lossless representation and includes rich metadata for files related to the target file.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality, while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
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
  + Work Steps
  + Implementation Gate
  + Acceptance
  + Evidence Targets
  + Actual-status Update
  + Commit Boundary
- Split planned work into separate slices when it contains more than one primary user-visible behavior, user trigger, render location, permission or visibility rule, DB write target, DB state transition, API/CLI/MCP contract, async/event/webhook flow, external side effect, cleanup/quarantine domain, behavior test target, independent acceptance gate, or independent commit boundary.
- Hidden fallback is forbidden. Prefer a visible failure over a fallback that hides a broken primary path.
- When touching DB-backed content, verify the full loop when applicable: UI input -> submit action -> DB write -> DB read after reload/new request -> correct UI render or omission. If there is no UI, replace UI steps with the real caller/consumer flow.
- Tests must prove product behavior. Delete or replace tests that only assert implementation details, helper output, static DOM existence, or mocked plumbing without proving trigger -> process -> observable result.
- If a planned item uses wording such as `and`, `also`, `then wire`, `plus update`, `both`, or `handle all`, check whether it is actually multiple slices.
- Do not write broad actionable items such as `Implement checkout, webhook, entitlement update, and billing UI`; split them into narrow slices such as `Create checkout session request`, `Persist checkout session state`, `Handle provider webhook`, `Update entitlement from webhook event`, and `Render billing status from entitlement`.
- Each slice work step must include UI flow, DB/data flow, render location, and evidence target checks. Use `N/A` with a reason when a check does not apply.
- If tests write DB rows, app state, files, queues, provider state, or other persistent data, the slice must define cleanup or quarantine before implementation.
- File-detail-specific invariant: compact output must not be summary-only. It must preserve all available detail facts for the requested file-detail scope by using normalized tables, dictionaries, references, and tuple ranges instead of repeated expanded objects.
- File-detail-specific invariant: `relatedFiles` must include metadata about relationship neighbor files, not just neighbor paths and edge counts.
- File-detail-specific invariant: compact full-detail machine output must not rely on hidden sample truncation. Default compact machine output must represent all rows available for the requested file unless the caller explicitly supplies limit flags. When limits are supplied, totals, returned counts, and omitted counts must make the limitation visible.
- Compatibility invariant: if any existing expanded response shape is changed, an explicit `expanded` mode or documented migration path must remain until Web/API/CLI/MCP consumers are updated and tested.

## Problem

`anvien file-detail <path> --json` currently returns an expanded nested object. It contains useful detail, but the representation repeats file paths, symbol IDs, source site IDs, range field names, and relationship object keys across many records. For large files, even limited samples are very long. The current shape also exposes `relationships.outboundByFile` and `relationships.inboundByFile`, but those groups only identify related file paths, counts, and samples. They do not publish rich metadata about each related file.

The accepted direction from discussion is not to reduce `file-detail` into a summary or `impact` clone. The change must keep `file-detail` as the full detail/inventory surface for one file, while making its representation compact and adding a first-class `relatedFiles` inventory.

## Scope

- Build a compact, lossless `file-detail` representation for machine output.
- Add rich related-file metadata for all files connected to the target through local/inbound/outbound relationship groups.
- Wire compact output through CLI, HTTP/API, and MCP/agent-facing file context surfaces with explicit format or preservation decisions.
- Update Web contract/types and the Web `FileDetailPanel` consumer so the user-visible file detail remains complete and accurate.
- Update README/RUNBOOK command examples or record an explicit no-change decision tied to the final default behavior.
- Preserve or intentionally gate the existing expanded shape while the new compact shape is introduced.
- Add tests and measurements proving compact output is smaller without dropping facts.

## Non-Goals

- Do not turn `file-detail` into a summary-only, top-N-only, or impact-only command.
- Do not remove relationship, unresolved, linked, or symbol detail facts to reduce size.
- Do not rename `file-detail` or resurrect the removed `file-context` command.
- Do not change graph analysis semantics, relationship resolution semantics, or unresolved classification rules.
- Do not rewrite unrelated Web graph UI layout beyond what is needed to consume and display compact file-detail data.
- Do not edit generated contract artifacts directly as source of truth; update the generator and regenerate.

## Requirements

- Compact output must provide a declared schema for row/table fields.
- Repeated file paths, symbol IDs, source site IDs, relationship kinds, proof kinds, statuses, and ranges must be normalized or interned.
- Ranges must use compact tuple form such as `[startLine, startColumn, endLine, endColumn]` with schema documentation.
- Relationships must be represented as rows referencing normalized files, symbols, and source sites.
- Unresolved source sites must be represented as rows referencing source symbols and source sites, with target text, kind, classification, actionability, proof, and status preserved.
- Symbol information must remain complete for the requested scope, but may use rows and parent references rather than expanded nested objects.
- `relatedFiles` must include path, language, kind, file role, file group, app layer, functional area, parse status, symbol count, unresolved count, risk, direction flags, relationship totals, and relationship-kind counts.
- Compact mode default for machine JSON must be full-row for relationships, unresolved sites, linked items, symbols, and related files unless explicit CLI/API/MCP limit parameters are supplied.
- Explicit limits must be visible in the payload through original totals, returned row counts, and omitted row counts; a limited response must not be indistinguishable from a complete response.
- Existing counts must remain accurate and must not be recomputed differently from the expanded builder without test coverage.
- CLI/API/MCP help, contract docs, and user-facing command docs must make compact versus expanded behavior and limit behavior explicit.
- Full build must pass before final validation, and Web-facing changes require real runtime validation per repo rules.

## Acceptance Criteria

- `anvien file-detail internal/filecontext/context.go --repo Anvien --json` returns compact full-detail output by the accepted default, with all available rows for the requested file unless explicit limit flags are supplied. Expanded output remains available through an explicit mode.
- A legacy or expanded mode remains available until all internal consumers are migrated, or the plan records an explicit breaking-change decision with tests and contract updates.
- Limited compact output records total, returned, and omitted counts so limits cannot masquerade as complete file detail.
- The compact response includes a `relatedFiles` table with metadata for every unique inbound/outbound relationship file.
- Tests prove compact output can represent the same symbols, relationship rows, unresolved rows, linked items, counts, source-site data, and graph metadata as the expanded builder output for fixture data.
- CLI tests cover compact mode, expanded mode, help text, and sample/row limit behavior.
- HTTP tests cover `format=compact`, `format=expanded`, invalid format, and related-file metadata.
- MCP/context tests or smoke evidence cover `context file` and file-impact flows after the shared file-context model changes. The plan must either preserve existing expanded MCP payloads with tests or introduce an explicit compact/expanded MCP contract.
- Web contracts and generated TS types are regenerated from source-of-truth contract changes.
- README/RUNBOOK examples are updated for final default/format behavior, or evidence records that examples remain valid because defaults are backwards-compatible.
- `FileDetailPanel` renders the same major sections from compact data and displays related-file metadata.
- Full build passes before final runtime validation.
- Required runtime/browser validation is recorded for the Web-visible file detail flow.
- `anvien detect-changes --repo Anvien --scope all` runs before each implementation-slice commit.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: inspect source-of-truth files, classify each surface, record blocked or missing pieces, and update later phase status assumptions, next actions, and work steps from evidence.
  - Implementation Gate: no implementation or editing starts until `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies correct, partial, missing/unbound, fake-or-stub, and blocked surfaces for this scope.

### P1: Compact Builder Model

- Phase Goal: add the compact full-detail data model and builder output while preserving current expanded data until consumers are migrated.
- Phase Boundary:
  - In scope: `internal/filecontext/context.go`, `internal/filecontext/context_test.go`, model helpers local to `internal/filecontext`.
  - Out of scope: CLI flag wiring, HTTP query params, Web contract generation, Web UI rendering.
  - Dependencies: P0 actual status; fresh Anvien graph; symbol/file impact evidence for edited builder symbols.
- Phase Implementation Rule: do not implement `P1` directly. Implement `P1-A`, verify it, record evidence, refresh actual-status, commit when required, then continue to `P1-B`.
- Ordered Slice List:
  - P1-A: Add compact DTO schema and conversion helpers.
  - P1-B: Add related-file metadata inventory.
  - P1-C: Prove compact facts preserve expanded facts.

- [x] P1-A: Add compact DTO schema and conversion helpers.
  - Goal: define compact file-detail structs and deterministic conversion from existing `FileContext`.
  - Scope Boundary:
    - Editable: `internal/filecontext/context.go`, `internal/filecontext/context_test.go`.
    - Inspect-only: `internal/cli/file_detail_command.go`, `internal/httpapi/file_context.go`, `internal/contracts/web_ui.go`, `anvien-web/src/generated/anvien-contracts.ts`.
    - Preserve-only: existing `FileContext` expanded struct fields and existing `BuildFileContext` behavior.
    - Out of scope: CLI/API format selection and Web rendering.
  - Non-Goals: do not change graph analysis, file lookup normalization, or current expanded output behavior in this slice.
  - Pre-flight Questions:
    - Data source: existing `FileContext` and `Builder` indexes.
    - Display permission: N/A for backend model.
    - DB read flow: N/A, graph JSON is the source.
    - DB write flow: N/A, no persistence writes.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A for this backend model slice.
    - Playwright target: N/A.
    - Behavior test: Go tests comparing compact facts against fixture expanded facts.
    - Cleanup/quarantine: no persistent test data.
    - External side effects: none.
    - N/A notes: this slice is model-only.
  - Work Steps:
    1. Run `anvien impact symbol` for `FileContext`, `BuildFileContext`, `RelationshipSample`, `UnresolvedSample`, and any edited helper before editing.
       - UI flow check: N/A.
       - DB/data flow check: confirms graph-derived facts are the only data source.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E1-P1A-IMPACT1`.
    2. Add compact structs with schema metadata, dictionary/table sections, row formats, tuple ranges, and explicit format/version fields.
       - UI flow check: N/A.
       - DB/data flow check: compact rows derive only from expanded `FileContext`.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E1-P1A-SRC1`, `E1-P1A-TEST1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - Do not modify current expanded fields or their JSON tags unless the slice records a compatibility decision.
  - Acceptance:
    - Source: compact DTOs and conversion helpers exist with explicit schema/version fields.
    - Runtime/UI: N/A.
    - DB/data: conversion derives from existing graph/file detail facts without new persistence.
    - Behavior test: fixture tests prove symbols, relationships, unresolved, linked, graph, target, quality, summary, and limits are all represented.
    - Cleanup/quarantine: no temp or persistent files outside repo.
    - Evidence IDs: `E1-P1A-IMPACT1`, `E1-P1A-SRC1`, `E1-P1A-TEST1`.
    - Actual-status rows refreshed: `internal/filecontext/context.go`.
  - Evidence Targets: impact, source diff summary, focused Go test output, compact size sample.
  - Actual-status Update: mark builder compact model `missing -> partial` or `correct` based on tests.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P1-B: Add related-file metadata inventory.
  - Goal: add a compact `relatedFiles` table with metadata for every unique related file.
  - Scope Boundary:
    - Editable: `internal/filecontext/context.go`, `internal/filecontext/context_test.go`.
    - Inspect-only: file summary classification helpers in the same package.
    - Preserve-only: existing relationship counts and current `outboundByFile`/`inboundByFile` semantics.
    - Out of scope: CLI/API/Web wiring.
  - Non-Goals: do not infer related files from text search or path heuristics when graph facts already provide relationships.
  - Pre-flight Questions:
    - Data source: `filesByPath`, existing relationship groups, and file summary classification.
    - Display permission: N/A.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: Go fixture with inbound, outbound, and bidirectional related files.
    - Cleanup/quarantine: no persistent data.
    - External side effects: none.
    - N/A notes: backend projection only.
  - Work Steps:
    1. Run symbol/file impact for any builder helper or summary helper changed by related-file metadata.
       - UI flow check: N/A.
       - DB/data flow check: confirms graph file nodes and relationships are source.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E1-P1B-IMPACT1`.
    2. Build related-file rows with metadata, direction flags, relation totals, relation-kind counts, and stable file references.
       - UI flow check: N/A.
       - DB/data flow check: row counts equal unique inbound/outbound relationship files.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E1-P1B-TEST1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - Related-file rows must not replace relationship rows; they must complement them.
  - Acceptance:
    - Source: compact output includes `relatedFiles` metadata rows.
    - Runtime/UI: N/A.
    - DB/data: related-file count equals unique inbound/outbound relationship files for fixtures and real smoke files.
    - Behavior test: tests prove metadata fields and direction flags.
    - Cleanup/quarantine: no persistent data.
    - Evidence IDs: `E1-P1B-IMPACT1`, `E1-P1B-TEST1`, `B1-P1B-SIZE1`.
    - Actual-status rows refreshed: `internal/filecontext/context.go`.
  - Evidence Targets: impact, focused Go tests, compact related-file inventory benchmark.
  - Actual-status Update: mark related-file metadata `missing -> correct` when tests and smoke pass.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P1-C: Prove compact facts preserve expanded facts.
  - Goal: add tests that prevent compact output from becoming a lossy summary.
  - Scope Boundary:
    - Editable: `internal/filecontext/context_test.go`.
    - Inspect-only: compact helpers from P1-A/P1-B.
    - Preserve-only: existing builder tests.
    - Out of scope: CLI/API/Web changes.
  - Non-Goals: do not snapshot huge JSON strings as the only proof.
  - Pre-flight Questions:
    - Data source: Go fixture graph and selected real smoke command after implementation.
    - Display permission: N/A.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: round-trip or count parity tests.
    - Cleanup/quarantine: no persistent data.
    - External side effects: none.
    - N/A notes: test-only slice.
  - Work Steps:
    1. Add parity assertions for full-scope counts, row totals, ranges, symbols, source sites, unresolved classifications, linked items, and related-file metadata.
       - UI flow check: N/A.
       - DB/data flow check: fixture graph -> expanded context -> compact context facts.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E1-P1C-TEST1`.
    2. Add explicit limited-output assertions proving total, returned, and omitted counts are present when relationship, unresolved, linked, or row limits are supplied.
       - UI flow check: N/A.
       - DB/data flow check: fixture graph -> limited compact context with visible omission metadata.
       - Render location check: N/A.
       - Evidence target: `E1-P1C-LIMIT1`.
    3. Measure expanded versus compact JSON size for `internal/filecontext/context.go`.
       - UI flow check: N/A.
       - DB/data flow check: same target file and graph basis for before/after.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `B1-P1C-SIZE1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - Tests must assert behavior and fact parity, not only struct existence.
    - Tests must separately prove full default output and explicit limited output semantics.
  - Acceptance:
    - Source: parity tests exist and fail if compact output drops facts.
    - Runtime/UI: N/A.
    - DB/data: no persistence.
    - Behavior test: `go test ./internal/filecontext -count=1` passes with full-scope and limited-output cases.
    - Cleanup/quarantine: no persistent data.
    - Evidence IDs: `E1-P1C-TEST1`, `E1-P1C-LIMIT1`, `B1-P1C-SIZE1`.
    - Actual-status rows refreshed: `internal/filecontext/context.go`.
  - Evidence Targets: focused Go tests and size benchmark.
  - Actual-status Update: mark compact builder model `correct` when parity tests pass.
  - Commit Boundary: commit after this slice when acceptance passes.

### P2: CLI, HTTP, MCP, Contract, and Docs Surface

- Phase Goal: expose compact full-detail output through command/API/MCP surfaces with explicit format controls, regenerated contracts, and user-facing command docs.
- Phase Boundary:
  - In scope: `internal/cli/file_detail_command.go`, `internal/cli/file_detail_command_test.go`, `internal/httpapi/file_context.go`, `internal/httpapi/file_context_test.go`, `internal/contracts/web_ui.go`, generated Web contract artifacts, MCP file-context consumers, README/RUNBOOK command docs.
  - Out of scope: Web visual rendering and browser validation.
  - Dependencies: P1 compact model completed.
- Phase Implementation Rule: do not implement `P2` directly. Implement `P2-A`, verify it, record evidence, refresh actual-status, commit when required, then continue through each P2 slice in order.
- Ordered Slice List:
  - P2-A: Wire CLI format selection.
  - P2-B: Wire HTTP format selection.
  - P2-C: Update Web/API contracts and generated types.
  - P2-D: Close MCP and agent file-context surfaces.
  - P2-E: Update command docs and examples.

- [x] P2-A: Wire CLI format selection.
  - Goal: make `anvien file-detail` emit compact full-detail JSON through a clear CLI contract while preserving expanded output access.
  - Scope Boundary:
    - Editable: `internal/cli/file_detail_command.go`, `internal/cli/file_detail_command_test.go`, `internal/cli/command_test.go`.
    - Inspect-only: `internal/filecontext/context.go`.
    - Preserve-only: path normalization, missing-file errors, outside-repo error behavior.
    - Out of scope: HTTP and Web.
  - Non-Goals: do not add a summary-only mode as the main fix.
  - Pre-flight Questions:
    - Data source: `BuildFileContext` plus compact conversion.
    - Display permission: CLI user has requested file detail.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: terminal output.
    - UI behavior flow: CLI command -> graph load -> file detail -> JSON/human output.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: CLI tests for compact/expanded/help/errors.
    - Cleanup/quarantine: no persistent data.
    - External side effects: reads graph only.
    - N/A notes: CLI-only slice.
  - Work Steps:
    1. Run impact for `newFileDetailCommand`, `renderFileContext`, `writeJSON`, and any edited helper.
       - UI flow check: CLI command flow only.
       - DB/data flow check: graph JSON read only.
       - Render location check: terminal stdout.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E2-P2A-IMPACT1`.
    2. Add `--format compact|expanded` or equivalent explicit option, document default behavior in help, and write compact JSON when selected.
       - UI flow check: CLI command output is observable.
       - DB/data flow check: same graph facts as expanded mode.
       - Render location check: stdout.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E2-P2A-TEST1`, `B2-P2A-SIZE1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - CLI help must be updated in the same slice as the new CLI option.
  - Acceptance:
    - Source: CLI supports compact and expanded output modes.
    - Runtime/UI: CLI smoke shows compact JSON and expanded JSON can both be requested.
    - DB/data: both modes use same graph and target.
    - Behavior test: focused CLI tests pass.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E2-P2A-IMPACT1`, `E2-P2A-TEST1`, `B2-P2A-SIZE1`.
    - Actual-status rows refreshed: CLI surface.
  - Evidence Targets: impact, CLI tests, CLI smoke, size benchmark.
  - Actual-status Update: mark CLI format surface `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P2-B: Wire HTTP format selection.
  - Goal: expose compact full-detail over `/api/file-detail` with an explicit query contract.
  - Scope Boundary:
    - Editable: `internal/httpapi/file_context.go`, `internal/httpapi/file_context_test.go`.
    - Inspect-only: `internal/filecontext/context.go`, `internal/contracts/web_ui.go`.
    - Preserve-only: existing route path, repo resolution, path normalization, stale graph handling.
    - Out of scope: generated contract update and Web renderer.
  - Non-Goals: do not add a second route unless contract evidence shows it is safer than a format param.
  - Pre-flight Questions:
    - Data source: `BuildFileContext` plus compact conversion.
    - Display permission: local runtime request.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: HTTP JSON response.
    - UI behavior flow: Web client -> `/api/file-detail` -> JSON response.
    - Docker runtime: final runtime validation only, not this focused Go test slice.
    - Playwright target: final Web validation after P3.
    - Behavior test: HTTP tests for compact, expanded, invalid format, and errors.
    - Cleanup/quarantine: test server only.
    - External side effects: none beyond local HTTP test.
    - N/A notes: HTTP unit/integration slice.
  - Work Steps:
    1. Run impact for `Server.handleFileContext` and any helper added for format parsing.
       - UI flow check: endpoint response surface.
       - DB/data flow check: graph snapshot read only.
       - Render location check: HTTP JSON body.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E2-P2B-IMPACT1`.
    2. Add HTTP format parsing, bounded validation, response selection, and tests for compact/expanded behavior.
       - UI flow check: endpoint returns expected format.
       - DB/data flow check: both formats share graph facts.
       - Render location check: response body schema.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E2-P2B-TEST1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - Invalid format must fail visibly; no hidden fallback to expanded or compact.
  - Acceptance:
    - Source: `/api/file-detail` supports compact and expanded response selection.
    - Runtime/UI: HTTP smoke can request both formats.
    - DB/data: response facts match target graph.
    - Behavior test: focused HTTP tests pass.
    - Cleanup/quarantine: test server cleaned up.
    - Evidence IDs: `E2-P2B-IMPACT1`, `E2-P2B-TEST1`.
    - Actual-status rows refreshed: HTTP surface.
  - Evidence Targets: impact, HTTP tests, API smoke.
  - Actual-status Update: mark HTTP format surface `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P2-C: Update Web/API contracts and generated types.
  - Goal: publish the compact file-detail contract and regenerate generated Web types.
  - Scope Boundary:
    - Editable: `internal/contracts/web_ui.go`, `internal/contracts/web_ui_test.go`, contract generation source, generated contract artifacts through the repo-native generator.
    - Inspect-only: `anvien-web/src/generated/anvien-contracts.ts`, `contracts/web-ui/anvien-web-contract.schema.json`.
    - Preserve-only: existing unrelated route and graph contracts.
    - Out of scope: Web component rendering changes.
  - Non-Goals: do not edit generated artifacts manually as source of truth.
  - Pre-flight Questions:
    - Data source: Go contract source.
    - Display permission: public local Web/API contract.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: generated TS/schema files.
    - UI behavior flow: Web client consumes generated type later.
    - Docker runtime: final validation only.
    - Playwright target: final validation only.
    - Behavior test: contract tests and generated-output diff review.
    - Cleanup/quarantine: generated output must match source.
    - External side effects: none.
    - N/A notes: contract slice.
  - Work Steps:
    1. Run impact for contract symbols edited and file-detail route metadata.
       - UI flow check: generated Web type consumer surface.
       - DB/data flow check: N/A.
       - Render location check: generated TS/schema paths.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E2-P2C-IMPACT1`.
    2. Update source contract for compact response shape, format query parameter, and regenerate generated artifacts using repo-native commands.
       - UI flow check: generated types are available to Web client.
       - DB/data flow check: N/A.
       - Render location check: `anvien-web/src/generated/anvien-contracts.ts`, `contracts/web-ui/anvien-web-contract.schema.json`.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E2-P2C-CONTRACT1`, `E2-P2C-TEST1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - Generated files must be produced by the generator, not manual patching.
  - Acceptance:
    - Source: contract source documents compact and expanded response choices.
    - Runtime/UI: generated Web type exposes compact response safely.
    - DB/data: N/A.
    - Behavior test: contract tests pass and generated files match source.
    - Cleanup/quarantine: no stale generated contract output.
    - Evidence IDs: `E2-P2C-IMPACT1`, `E2-P2C-CONTRACT1`, `E2-P2C-TEST1`.
    - Actual-status rows refreshed: contract/generation surfaces.
  - Evidence Targets: impact, contract tests, generation command, generated file inventory.
  - Actual-status Update: mark contract surface `partial -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P2-D: Close MCP and agent file-context surfaces.
  - Goal: keep Anvien MCP/agent file-context behavior correct after shared compact file-detail model changes.
  - Scope Boundary:
    - Editable: `internal/mcp/target_dispatch.go`, `internal/mcp/context.go`, `internal/mcp/impact.go`, focused MCP tests only if the chosen contract requires changes.
    - Inspect-only: `internal/filecontext/context.go`, `internal/mcp/tools.go`, `internal/mcp/resources.go`, CLI context tests.
    - Preserve-only: unrelated MCP route/tool-map behavior.
    - Out of scope: Web rendering and HTTP endpoint contract.
  - Non-Goals: do not silently change MCP payload shape without tests and an explicit compatibility decision.
  - Pre-flight Questions:
    - Data source: `mcpBuildRepoFileContext` and the compact/expanded builder output from P1.
    - Display permission: local MCP/agent caller requests file context or file impact.
    - DB read flow: N/A, graph snapshot read only.
    - DB write flow: N/A.
    - Render location: MCP JSON/tool payload and CLI-equivalent context output.
    - UI behavior flow: agent/tool request -> MCP context or impact handler -> file-context payload.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: focused MCP or CLI-equivalent tests for `context file` and file-impact flows.
    - Cleanup/quarantine: no persistent data.
    - External side effects: reads graph only.
    - N/A notes: MCP/API surface slice, not browser UI.
  - Work Steps:
    1. Run Anvien impact/file-detail evidence for `internal/mcp/target_dispatch.go`, `internal/mcp/context.go`, and `internal/mcp/impact.go`, then record the blast radius.
       - UI flow check: MCP tool payload flow.
       - DB/data flow check: graph JSON read only.
       - Render location check: MCP payload and CLI-equivalent JSON.
       - Evidence target: `E2-P2D-IMPACT1`.
    2. Choose and document the MCP contract: preserve expanded payload, add explicit compact/expanded selection, or adapt internal compact data while preserving the public payload.
       - UI flow check: caller-visible MCP payload contract.
       - DB/data flow check: same graph facts as CLI/API compact or expanded mode.
       - Render location check: MCP `fileContext`, file layer, and file-impact payload sections.
       - Evidence target: `E2-P2D-SRC1`.
    3. Add or update focused tests/smoke evidence for `context file` and file impact so the chosen MCP behavior is proven.
       - UI flow check: tool/CLI-equivalent invocation is observable.
       - DB/data flow check: payload facts match target graph.
       - Render location check: JSON payload shape and compatibility fields.
       - Evidence target: `E2-P2D-TEST1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for every MCP file or symbol touched and record HIGH/CRITICAL blast-radius warnings.
    - The chosen MCP behavior must be explicit before any MCP payload shape change is made.
  - Acceptance:
    - Source: MCP file-context consumers either preserve expanded payload compatibility or expose a tested explicit compact/expanded contract.
    - Runtime/UI: MCP or CLI-equivalent smoke proves `context file` and file-impact flows still work.
    - DB/data: same graph/file facts are used as file-detail builder output.
    - Behavior test: focused MCP/context tests pass.
    - Cleanup/quarantine: no temp files outside repo.
    - Evidence IDs: `E2-P2D-IMPACT1`, `E2-P2D-SRC1`, `E2-P2D-TEST1`.
    - Actual-status rows refreshed: MCP/agent file-context surface.
  - Evidence Targets: impact, source contract decision, focused MCP/context tests, smoke command.
  - Actual-status Update: mark MCP/agent file-context surface `partial -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P2-E: Update command docs and examples.
  - Goal: keep user-facing `file-detail`, `context file`, and `/api/file-detail` documentation aligned with the final compact/expanded/default behavior.
  - Scope Boundary:
    - Editable: `README.md`, `RUNBOOK.md`, and source-of-truth docs that own generated agent command guidance if a generated doc update is required.
    - Inspect-only: `AGENTS.md`, generated `CLAUDE.md` or generated command-guide outputs.
    - Preserve-only: generated docs unless their generator/source is updated first.
    - Out of scope: implementation code changes.
  - Non-Goals: do not edit generated `AGENTS.md` or `CLAUDE.md` content as the permanent source of truth.
  - Pre-flight Questions:
    - Data source: final CLI/API/MCP behavior from P2-A through P2-D.
    - Display permission: public repo docs and local operator runbook.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: README/RUNBOOK command examples and any generated command guide source.
    - UI behavior flow: user reads docs -> runs compact or expanded command/API request.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: docs examples match implemented command/API behavior by smoke command or explicit no-change evidence.
    - Cleanup/quarantine: no generated-doc manual edits.
    - External side effects: none.
    - N/A notes: documentation slice.
  - Work Steps:
    1. Inspect README/RUNBOOK and generated-guide source ownership for `file-detail`, `context file`, and `/api/file-detail` examples.
       - UI flow check: docs-to-command path.
       - DB/data flow check: N/A.
       - Render location check: README/RUNBOOK sections and generated-guide source if applicable.
       - Evidence target: `E2-P2E-DOC1`.
    2. Update docs or record a no-change decision if the examples remain valid because defaults are backwards-compatible.
       - UI flow check: examples name compact/expanded and limit semantics clearly.
       - DB/data flow check: N/A.
       - Render location check: docs diff only.
       - Evidence target: `E2-P2E-DOC2`.
  - Implementation Gate:
    - P2-A through P2-D behavior decisions must be known before docs are finalized.
    - Do not edit generated agent docs as source of truth.
  - Acceptance:
    - Source: README/RUNBOOK or owning doc source reflects compact/expanded/default and limit behavior, or evidence records why no change is required.
    - Runtime/UI: N/A.
    - DB/data: N/A.
    - Behavior test: smoke command/API examples are valid or documented no-change evidence is recorded.
    - Cleanup/quarantine: no generated doc drift.
    - Evidence IDs: `E2-P2E-DOC1`, `E2-P2E-DOC2`, `E2-P2E-SMOKE1`, `E2-P2E-DETECT1`.
    - Actual-status rows refreshed: docs surface.
  - Evidence Targets: docs source inspection, docs diff or no-change decision, optional command/API smoke.
  - Actual-status Update: mark docs surface `partial -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

### P3: Web Consumer and Runtime Validation

- Phase Goal: make the Web file detail panel consume compact full-detail data and display related-file metadata without losing existing sections.
- Phase Boundary:
  - In scope: `anvien-web/src/services/backend-client.ts`, `anvien-web/src/components/FileDetailPanel.tsx`, focused Web unit tests, e2e/runtime validation.
  - Out of scope: unrelated graph canvas and file map redesign.
  - Dependencies: P1 and P2 complete.
- Phase Implementation Rule: do not implement `P3` directly. Implement `P3-A`, verify it, record evidence, refresh actual-status, commit when required, then continue to `P3-B`.
- Ordered Slice List:
  - P3-A: Update Web client and adapter.
  - P3-B: Render related-file detail and preserve existing sections.
  - P3-C: Validate built runtime file-detail flow.

- [ ] P3-A: Update Web client and adapter.
  - Goal: fetch compact file-detail data and adapt it into a stable view model for the existing panel.
  - Scope Boundary:
    - Editable: `anvien-web/src/services/backend-client.ts`, new or existing Web adapter/test files.
    - Inspect-only: `anvien-web/src/components/FileDetailPanel.tsx`, generated contract types.
    - Preserve-only: backend URL resolution and existing error handling.
    - Out of scope: UI layout changes.
  - Non-Goals: do not make a hidden fallback to expanded output if compact parsing fails.
  - Pre-flight Questions:
    - Data source: `/api/file-detail?format=compact`.
    - Display permission: local Web runtime.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: File detail panel view model.
    - UI behavior flow: panel selects file -> client requests compact detail -> adapter prepares render model.
    - Docker runtime: final P3-C.
    - Playwright target: final P3-C.
    - Behavior test: Web unit tests for fetch params and adapter conversion.
    - Cleanup/quarantine: no persistent browser data.
    - External side effects: local HTTP request only.
    - N/A notes: client/adapter slice.
  - Work Steps:
    1. Run impact/file-detail for `backend-client.ts` and relevant exported client function before editing.
       - UI flow check: file-detail fetch function.
       - DB/data flow check: HTTP response parse only.
       - Render location check: view model consumed by panel.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E3-P3A-IMPACT1`.
    2. Update fetch options and adapter logic so compact rows become renderable without data loss.
       - UI flow check: unit test confirms fetch includes compact format and adapter keeps all rows.
       - DB/data flow check: compact fixture -> view model.
       - Render location check: FileDetailPanel props/state.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E3-P3A-TEST1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - Adapter must fail visibly on malformed compact payload.
  - Acceptance:
    - Source: client requests compact file-detail and adapter preserves detail.
    - Runtime/UI: unit tests cover fetch and adapter.
    - DB/data: N/A.
    - Behavior test: Web unit tests pass for client/adapter.
    - Cleanup/quarantine: no persistent browser state.
    - Evidence IDs: `E3-P3A-IMPACT1`, `E3-P3A-TEST1`.
    - Actual-status rows refreshed: Web client.
  - Evidence Targets: impact, unit tests.
  - Actual-status Update: mark Web client compact binding `missing -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-B: Render related-file detail and preserve existing sections.
  - Goal: update `FileDetailPanel` so compact data renders summary, quality, symbols, relationships, unresolved, linked, and related-file metadata.
  - Scope Boundary:
    - Editable: `anvien-web/src/components/FileDetailPanel.tsx`, `anvien-web/test/unit/FileDetailPanel.test.tsx`.
    - Inspect-only: `CodeReferencesPanel.tsx` caller.
    - Preserve-only: file focus behavior and existing major section test IDs.
    - Out of scope: broad visual redesign.
  - Non-Goals: do not replace relationship details with only related-file summary.
  - Pre-flight Questions:
    - Data source: compact file-detail view model.
    - Display permission: selected file detail in local Web UI.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: File detail right-panel section.
    - UI behavior flow: file selection -> file detail panel -> compact sections visible.
    - Docker runtime: final P3-C.
    - Playwright target: final P3-C.
    - Behavior test: component test with related-file metadata.
    - Cleanup/quarantine: no persistent browser state.
    - External side effects: none.
    - N/A notes: UI component slice.
  - Work Steps:
    1. Run impact/file-detail for `FileDetailPanel.tsx` and inspect caller props.
       - UI flow check: panel render.
       - DB/data flow check: compact fixture data only in unit tests.
       - Render location check: `file-detail-section-*` test IDs.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E3-P3B-IMPACT1`.
    2. Render related-file metadata in a compact table/list while keeping relationship and unresolved details visible.
       - UI flow check: unit test confirms related file metadata appears with direction and risk.
       - DB/data flow check: compact fixture -> visible rows.
       - Render location check: existing sections and new related-file section if added.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E3-P3B-TEST1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - Existing sections must remain represented unless actual status is updated with an approved reason.
  - Acceptance:
    - Source: panel renders compact detail and related-file metadata.
    - Runtime/UI: component tests pass for major sections and related files.
    - DB/data: N/A.
    - Behavior test: Web unit tests prove compact payload render.
    - Cleanup/quarantine: no persistent state.
    - Evidence IDs: `E3-P3B-IMPACT1`, `E3-P3B-TEST1`.
    - Actual-status rows refreshed: FileDetailPanel.
  - Evidence Targets: impact, unit tests, screenshots if useful.
  - Actual-status Update: mark Web render surface `partial -> correct`.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] P3-C: Validate built runtime file-detail flow.
  - Goal: prove the real built local runtime serves compact file-detail and the Web UI renders it.
  - Scope Boundary:
    - Editable: tests only if missing a real behavior assertion.
    - Inspect-only: Docker/runtime scripts and Playwright config.
    - Preserve-only: unrelated Web flows.
    - Out of scope: feature development.
  - Non-Goals: do not substitute a dev server or mocked backend for required runtime evidence.
  - Pre-flight Questions:
    - Data source: real built runtime graph and `/api/file-detail`.
    - Display permission: local runtime.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: browser-visible Web file detail panel.
    - UI behavior flow: open Web runtime -> select file -> file detail panel renders compact-backed details.
    - Docker runtime: required.
    - Playwright target: built Docker/container runtime URL.
    - Behavior test: Playwright or browser evidence with screenshot/trace/result.
    - Cleanup/quarantine: stop runtime containers/processes after validation.
    - External side effects: local Docker/container runtime only.
    - N/A notes: final runtime validation slice.
  - Work Steps:
    1. Run full build including Docker/container build and record command/result.
       - UI flow check: built app exists.
       - DB/data flow check: graph/runtime can serve file detail.
       - Render location check: Web runtime URL.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E3-P3C-BUILD1`.
    2. Start built runtime, run browser/Playwright validation for file detail, and record screenshot/trace/result.
       - UI flow check: user-visible file detail displays related-file metadata and existing detail sections.
       - DB/data flow check: `/api/file-detail?format=compact` returns compact payload from runtime.
       - Render location check: FileDetailPanel in Web UI.
       - Mini QA for each completed implementation slice (MUST)
          For Codex: When Mini QA must use plugins
            - Browser: Control the in-app browser
            - Chrome: Control the user's real Chrome browser
            - Computer Use: Control Windows apps or installed artifacts when mini QA requires real app interaction outside a browser.
            - Playwright: use as an automation arm for browser actions, control sweeps, screenshots, videos, traces, and reports.
          For Claude or other agents:
            - Use the equivalent browser, Chrome/session, or computer-control capability exposed by that agent/runtime or Playwright-like capability available in that environment.
       - Evidence target: `E3-P3C-RUNTIME1`.
  - Implementation Gate:
    - Before editing target files, run the relevant Anvien impact/file-detail command for files, symbols, routes, tools, or contracts touched by this slice, and record the evidence IDs.
    - Docker/container runtime must be used for final UI validation.
  - Acceptance:
    - Source: no new source changes unless validation reveals a bug.
    - Runtime/UI: built runtime file-detail UI passes browser/Playwright validation.
    - DB/data: compact endpoint returns expected payload.
    - Behavior test: runtime trace/screenshot/result recorded.
    - Cleanup/quarantine: runtime stopped or container state recorded.
    - Evidence IDs: `E3-P3C-BUILD1`, `E3-P3C-RUNTIME1`.
    - Actual-status rows refreshed: final Web/runtime status.
  - Evidence Targets: full build, Docker runtime, browser/Playwright evidence.
  - Actual-status Update: mark runtime UI validation `missing -> correct` or blocked.
  - Commit Boundary: commit after this slice when acceptance passes.

- [ ] Pn-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, changed files, generated output, and validation results before closure.
  - Work Steps:
    1. Call the supervisor skill to review the full completed plan work.
    2. If supervisor fails the work, return to the responsible implementation workflow/skill for the failed scope only.
    3. Re-run supervisor review after the fix.
    4. Repeat until supervisor passes or records a blocker.
  - Implementation Gate: all planned implementation phases must be completed or explicitly blocked before this review.
  - Acceptance: supervisor review passes, or the plan records a blocker with evidence and no closure is performed.
- [ ] Pn-B: Remove dead work created during this plan.
  - Goal: ensure the final diff contains only artifacts that still serve the accepted plan.
  - Work Steps:
    1. Review files, sections, generated output, tests, temp files, and plan artifacts created or modified during this plan.
    2. Remove or rewrite any artifact made obsolete by actual-status findings, user corrections, failed approaches, or phase status updates.
    3. Verify no rejected approach, stale placeholder, unused generated output, or dead helper artifact remains in the final diff.
    4. Call supervisor to review the dead-work cleanup.
    5. If supervisor fails the cleanup, return to the responsible implementation workflow/skill for the failed cleanup scope only, then re-run supervisor review.
  - Implementation Gate: only remove artifacts created by this plan unless the user explicitly approves broader cleanup.
  - Acceptance: final `git diff/status` contains no dead plan-created artifacts, supervisor passes the cleanup, and evidence records what was removed or preserved.
- [ ] Pn-C: Close the plan.
  - Goal: finish validation, evidence, benchmark, detect-changes, commit, and final status.
  - Work Steps:
    1. Run the required final validation for the accepted scope, including full build before final runtime validation. For app/runtime scopes, full build must include Docker image/container build.
    2. Start the real built Docker/container runtime for app/runtime validation. If Docker cannot be built or started, record the blocker and do not substitute a host dev server.
    3. Validate public runtime or UI-facing changes with browser or Playwright evidence against the real built Docker/container runtime. Playwright evidence must include Docker build/run or compose command, container/service name, exposed URL, Playwright command, and screenshot/trace/result.
    4. Regenerate generated outputs if source-of-truth changes require it.
    5. Run Anvien detect-changes before commit when implementation work was performed.
    6. Record final validation, detect-changes, benchmark, and commit evidence.
    7. Commit the completed scope and verify the worktree state.
  - Implementation Gate: Pn-A and Pn-B must pass or record blockers.
  - Acceptance: final evidence is recorded, required commits exist, and the worktree state is known.

## Risk Notes

- `internal/filecontext/context.go` has high/critical blast radius across CLI, HTTP/API, MCP, Web, contracts, and graph-health consumers. Treat blast radius as a scope warning, not a prohibition.
- Changing default JSON response shape can break existing consumers. The plan requires explicit compact/expanded modes and contract updates.
- Compact representation can accidentally become lossy if tests only compare counts. Parity tests must cover rows, ranges, symbol/source-site references, unresolved metadata, linked items, and related-file metadata.
- Related-file metadata must come from graph/file summary facts, not path heuristics.
- The existing Web panel currently slices visible groups for rendering. That UI presentation may stay limited, but compact machine payload must remain a full-detail representation for the requested scope.
