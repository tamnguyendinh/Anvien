# Evidence

Title: AI Context Rule 1 Help Wording
Date: 2026-06-06
Status: Completed

## E0 - Scope Evidence

- `anvien analyze . --force`
  Result: completed before plan. Graph reported 1408 files, 84487 nodes, 122973 relationships, 16570 file projection dependency edges, and 430 unresolved items.

## E1 - Impact Evidence

- `anvien context symbol "renderMasterRulesBlock" --repo Anvien`
  Result: symbol found in `internal/aicontext/aicontext.go`, range 210-228, return type `string`.

- `anvien impact symbol "renderMasterRulesBlock" --repo Anvien --direction upstream`
  Result: HIGH blast radius. Direct chain reaches `renderAnvienBlock`, `GenerateAIContextFiles`, and `generateAnalyzeAIContext`. Affected files: `internal/aicontext/aicontext.go` and `internal/cli/analyze_postrun.go`. Affected processes include `NewAnalyzeCommand -> UpsertSection`, `NewAnalyzeCommand -> RemoveGeneratedSkills`, and `NewAnalyzeCommand -> AnalyzeAIContextResult`.

## E2 - Implementation Evidence

- `internal/aicontext/aicontext.go`
  Changed generated rule 1 from the previous codebase-analysis wording to `1. How to use anvien: run command "anvien --help".`

- `internal/aicontext/aicontext_test.go`
  Updated `requireGeneratedMasterRules` to assert the new lowercase `anvien --help` wording.

## E3 - Validation Evidence

- `gofmt -w internal\aicontext\aicontext.go internal\aicontext\aicontext_test.go`
  Result: completed successfully.

- Full build sequence, first execution:
  Result: failed during `npm install` / package runtime build because `E:\Anvien\anvien\bin\anvien.exe` was locked.
  Investigation: a stale aborted `anvien analyze --force` process was stopped first. Restart Manager then identified global Anvien MCP PID `12304` as the process holding the local runtime binary. That single locking process was stopped; the other Anvien MCP process remained running.

- Full build sequence, rerun:
  Result: completed successfully with fail-fast PowerShell wrapper around the exact required command sequence.
  Commands executed in order:
  `cd .\anvien`; `npm install`; `npm run build`; `npm install -g .`; `Get-Command anvien`; `anvien version`; `cd ..`; `powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1`; `anvien version`; `anvien analyze . --force`.
  Version checks returned `1.2.5`.
  Vite emitted existing non-failing warnings about a static/dynamic import mix and large chunks.
  Final analyze in the full build reported 1410 files, 84500 nodes, 122986 relationships, 16570 file projection dependency edges, and 430 unresolved items.

- `go test ./internal/aicontext ./internal/cli -count=1`
  Result: passed. Package results: `ok github.com/tamnguyendinh/anvien/internal/aicontext 33.584s`; `ok github.com/tamnguyendinh/anvien/internal/cli 97.688s`.

- Generated output smoke:
  `AGENTS.md`, `CLAUDE.md`, `internal/aicontext/aicontext.go`, and `internal/aicontext/aicontext_test.go` all contain `1. How to use anvien: run command "anvien --help".`; no old rule-1 codebase-analysis wording was found in those target files.

## E4 - Commit Evidence

- Final pre-commit graph refresh:
  `anvien analyze . --force` completed successfully after report/notes docs were added. Graph reported 1411 files, 84509 nodes, 122995 relationships, 16570 file projection dependency edges, and 430 unresolved items.

- `anvien detect-changes --repo Anvien --scope all`
  Result: completed successfully.
  Summary risk: low.
  Changed app layers: backend 1, backend_test 1, docs 3.
  Changed functional areas: documentation 3, unknown 2.
  Changed implementation/test files reported by detect-changes: `internal/aicontext/aicontext.go`, `internal/aicontext/aicontext_test.go`, and tracked `docs/notes_decisions_log/notes_decisions_log_20260606.md`.
  Affected processes: none reported in final detect output.

- Implementation checkpoint:
  Pending until git checkpoint is created.
