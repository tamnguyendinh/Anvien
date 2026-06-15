# File Classification Plain Text XSD Actual Status

## P0-A Actual Status

- Status: complete.
- Baseline command: `anvien analyze --force`.
- Baseline result: `unknown=116`.
- Unknown extension inventory:
  - `.txt`: 38 files.
  - `.xsd`: 78 files.
- Decision: implement generic extension classification, not path-specific rules.
- Blast radius:
  - `internal/documents/documents.go`: file risk `high`.
  - `documents.Kind`: symbol impact risk `CRITICAL`.
  - Proceeding because change is a narrow extension mapping with behavior tests.

## P1-A Plain Text

- Status: complete.
- Current truth: `.txt` maps to `plain_text`.
- Evidence: `go test ./internal/documents ./internal/analyze` passed.
- Analyze evidence: source-current `go run ./cmd/anvien analyze --force` reports `unknown=78`, leaving only `.xsd` unknown files from the original inventory.
- Commit: `8dd1179 classify plain text documents`.

## P1-B XML Schema

- Status: complete.
- Current truth: `.xsd` maps to `xml_schema`.
- Evidence: `go test ./internal/documents ./internal/analyze` passed.
- Analyze evidence: source-current analyze and post-build binary analyze both report `unknown=0`.
- Next action: commit P1-B.

## Final

- Status: complete.
- Final build: `npm run full-build` passed after clearing stale Anvien binary lock processes.
- Final count: `unknown=0`.
- UI/browser validation: N/A; no UI changed.
- Docker validation: N/A; no Docker/container source changed.
