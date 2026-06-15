# File Classification Plain Text XSD Plan

## Metadata

- Date: `2026-06-15`
- Status: `complete`
- Plan: `docs/plans/2026-06-15-file-classification-plain-text-xsd/2026-06-15-file-classification-plain-text-xsd-plan.md`
- Evidence: `docs/plans/2026-06-15-file-classification-plain-text-xsd/2026-06-15-file-classification-plain-text-xsd-evidence.md`
- Benchmark: `docs/plans/2026-06-15-file-classification-plain-text-xsd/2026-06-15-file-classification-plain-text-xsd-benchmark.md`
- Actual status: `docs/plans/2026-06-15-file-classification-plain-text-xsd/2026-06-15-file-classification-plain-text-xsd-actual-status.md`

## Goal

Classify repo-agnostic `.txt` files as `document/plain_text` and `.xsd` files as `schema/xml_schema` so they no longer appear in `anvien analyze` unknown file gaps.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Run Anvien file-detail and impact before editing target files or symbols.
- Run Anvien detect-changes before every implementation-slice commit.
- Run full build before final validation.
- Keep this scoped to generic file classification, not Anvien-repo-specific path rules.
- Do not add hidden fallback behavior.

## Problem

`anvien analyze --force` reports `unknown=116`. Investigation showed every unknown file is a generic extension category that can be classified safely across repositories:

- `.txt` count: 38
- `.xsd` count: 78

## Scope

- `internal/documents/documents.go`
- `internal/documents/documents_test.go`
- `internal/analyze/file_classification_test.go`
- Plan evidence and actual-status files

## Non-Goals

- No path-specific Anvien repository exceptions.
- No new parser or analyzer for plain text or XML schema content.
- No UI/runtime behavior changes.
- No metadata-only reclassification for `.txt`.

## Requirements

- `.txt` maps to document kind `plain_text`.
- `.xsd` maps to document/schema kind `xml_schema`.
- File classification metrics count these files under `documents`, not `unknown`.
- Repo scan unknown count drops from 116 to 0 when only these two extension groups remain unknown.

## Acceptance Criteria

- `documents.Kind("notes.txt") == "plain_text"`.
- `documents.Kind("schemas/example.xsd") == "xml_schema"`.
- `classifyFileMetrics` counts `.txt` and `.xsd` as `documents`.
- Relevant Go tests pass.
- Full build passes.
- Final `anvien analyze --force` shows `unknown=0`.
- `anvien detect-changes --repo E:\Anvien --scope all` runs before each implementation-slice commit.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state.
  - Work Steps: refresh graph, inspect classifier source/tests, record unknown extension counts, run target file-detail and impact.
  - Implementation Gate: no source edit starts until this row is complete.
  - Acceptance: actual status identifies baseline unknown count and target blast radius.

### P1: Generic Extension Classification

- Phase Goal: remove generic `.txt` and `.xsd` files from unknown classification without path-specific rules.
- Phase Boundary:
  - In scope: extension-to-kind mapping and behavior tests.
  - Out of scope: parser extraction, content indexing, UI, DB, MCP contract changes.
  - Dependencies: fresh Anvien graph and impact evidence.
- Phase Implementation Rule: do not implement `P1` directly. Implement `P1-A`, verify it, record evidence, refresh actual-status, commit, then continue to `P1-B`.
- Ordered Slice List:
  - P1-A: Classify plain text files.
  - P1-B: Classify XML schema files.

- [x] P1-A: Classify plain text files.
  - Goal: `.txt` files classify as `document/plain_text`.
  - Scope Boundary:
    - Editable: `internal/documents/documents.go`, relevant tests, plan evidence.
    - Inspect-only: `internal/analyze/file_classification.go`.
    - Preserve-only: existing markdown/word/pdf/spreadsheet behavior.
    - Out of scope: path-specific rules and text parsing.
  - Non-Goals: no metadata-only treatment for `.txt`.
  - Pre-flight Questions:
    - Data source: scanner file extension.
    - Display permission: N/A; CLI/analyze metric only.
    - DB read flow: N/A; no DB.
    - DB write flow: N/A; no DB.
    - Render location: CLI analyze output and JSON file metrics.
    - UI behavior flow: N/A; no UI.
    - Behavior test: Go tests prove `.txt` is document bucket and document kind `plain_text`.
    - Cleanup/quarantine: tests use in-memory/t.TempDir only.
    - External side effects: N/A.
    - N/A notes: UI/DB/provider flows do not apply to this analyzer classification change.
  - Work Steps:
    1. Add `.txt` to the shared document kind mapping.
       - UI flow check: N/A; no UI.
       - DB/data flow check: scanner path -> `documents.Kind` -> `classifyFile` -> `FileBucketDocuments`.
       - Render location check: `anvien analyze` file metrics.
       - Evidence target: `E-P1A-SOURCE`.
    2. Add behavior tests for plain text classification.
       - UI flow check: N/A; no UI.
       - DB/data flow check: test input file path -> metrics document bucket.
       - Render location check: JSON/CLI metrics contract.
       - Evidence target: `E-P1A-TEST`.
  - Implementation Gate:
    - `E-P0-FD-DOCS`, `E-P0-IMPACT-DOCS`, `E-P0-IMPACT-KIND`, `E-P0-FD-DOCS-TEST`, and `E-P0-FD-CLASS-TEST` are recorded.
  - Acceptance:
    - Source: `.txt` returns `plain_text`.
    - Runtime/UI: N/A; no UI.
    - DB/data: file metrics count `.txt` under documents.
    - Behavior test: relevant Go tests pass.
    - Cleanup/quarantine: no persistent test data.
    - Evidence IDs: `E-P1A-TEST`, `E-P1A-DETECT`.
    - Actual-status rows refreshed: P1-A.
  - Evidence Targets: source diff, Go test output, detect-changes output, commit hash.
  - Actual-status Update: mark `.txt` behavior implemented and verified.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] P1-B: Classify XML schema files.
  - Goal: `.xsd` files classify as `schema/xml_schema`.
  - Scope Boundary:
    - Editable: `internal/documents/documents.go`, relevant tests, plan evidence.
    - Inspect-only: `internal/analyze/file_classification.go`.
    - Preserve-only: existing XML metadata-only behavior for `.xml`.
    - Out of scope: XML schema parsing or validation.
  - Non-Goals: no path-specific OOXML schema rule.
  - Pre-flight Questions:
    - Data source: scanner file extension.
    - Display permission: N/A; CLI/analyze metric only.
    - DB read flow: N/A; no DB.
    - DB write flow: N/A; no DB.
    - Render location: CLI analyze output and JSON file metrics.
    - UI behavior flow: N/A; no UI.
    - Behavior test: Go tests prove `.xsd` is document bucket and document kind `xml_schema`.
    - Cleanup/quarantine: tests use in-memory/t.TempDir only.
    - External side effects: N/A.
    - N/A notes: UI/DB/provider flows do not apply to this analyzer classification change.
  - Work Steps:
    1. Add `.xsd` to the shared document/schema kind mapping.
       - UI flow check: N/A; no UI.
       - DB/data flow check: scanner path -> `documents.Kind` -> `classifyFile` -> `FileBucketDocuments`.
       - Render location check: `anvien analyze` file metrics.
       - Evidence target: `E-P1B-SOURCE`.
    2. Add behavior tests for XML schema classification and run final analyze/build validation.
       - UI flow check: N/A; no UI.
       - DB/data flow check: test input file path and real repo scan -> unknown drops to 0.
       - Render location check: CLI analyze output shows `gaps: unknown=0`.
       - Evidence target: `E-P1B-TEST`, `E-P1B-ANALYZE`, `E-P1B-BUILD`.
  - Implementation Gate:
    - P1-A is committed and actual-status is refreshed.
  - Acceptance:
    - Source: `.xsd` returns `xml_schema`.
    - Runtime/UI: N/A; no UI.
    - DB/data: real scan unknown count is 0.
    - Behavior test: relevant Go tests pass and full build passes.
    - Cleanup/quarantine: no persistent test data.
    - Evidence IDs: `E-P1B-TEST`, `E-P1B-ANALYZE`, `E-P1B-BUILD`, `E-P1B-DETECT`.
    - Actual-status rows refreshed: P1-B.
  - Evidence Targets: source diff, Go test output, full build output, analyze output, detect-changes output, commit hash.
  - Actual-status Update: mark `.xsd` behavior implemented and verified.
  - Commit Boundary: commit after this slice when acceptance passes.

- [x] Pn-A: Close the plan.
  - Goal: final evidence, benchmark statement, detect-changes, commit, and known worktree state.
  - Work Steps:
    1. Verify final status and evidence files.
    2. Confirm Docker is out of scope because no container/runtime packaging files changed.
    3. Confirm UI/browser validation is out of scope because no UI changed.
  - Implementation Gate: P1-A and P1-B complete.
  - Acceptance: final evidence and commits are recorded and worktree state is known.

## Risk Notes

- `documents.Kind` impact is CRITICAL because it feeds analyzer classification and document indexing.
- The edit is intentionally limited to generic extension mapping and behavior tests.
