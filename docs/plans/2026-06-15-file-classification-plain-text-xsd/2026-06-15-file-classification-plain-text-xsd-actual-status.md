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
- Next action: commit P1-A, then add `.xsd` mapping and tests.

## P1-B XML Schema

- Status: pending.
- Current truth: `.xsd` is still unknown after P1-A.
- Next action: add `.xsd` mapping and tests after P1-A.

## Final

- Status: pending.
