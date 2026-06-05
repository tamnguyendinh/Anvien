# Evidence Ledger

Title: AI Context Master Rules Generation
Date: 2026-06-06
Status: Completed
Plan: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-plan.md
Evidence: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-evidence.md
Benchmark: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-benchmark.md

## E0 - Scope Evidence

- `anvien analyze . --force`
  Result: completed before plan creation. Graph reported 1404 files, 84448 nodes, 122937 relationships, 16570 file projection dependency edges, and 430 unresolved items.

- `anvien query files "AGENTS.md CLAUDE.md aicontext generate master iron rules" --repo Anvien`
  Result: broad discovery surfaced CLI command files but confirmed `aicontext` as the correct concept family.

- `anvien query files "Skill Selection Guide AI agent chooses skill generator" --repo Anvien`
  Result: identified `internal/aicontext/aicontext_test.go` and `internal/aicontext/skill_packages.go`; test surface includes `TestGenerateAIContextFilesCreatesManagedContextAndSkillPackages`.

- `anvien file-context internal/aicontext/aicontext.go --repo Anvien --json`
  Result: `internal/aicontext/aicontext.go` is parsed, backend, high risk, 12 inbound refs, 12 outbound refs, 2 linked flows, and 2 linked tests.

- `anvien context symbol "renderAnvienBlock" --repo Anvien`
  Result: symbol found in `internal/aicontext/aicontext.go`, range 94-207, return type `string`. Linked tests include `internal/aicontext/aicontext_test.go` and `internal/cli/command_test.go`.

- `anvien impact symbol "renderAnvienBlock" --repo Anvien --direction upstream`
  Result: CRITICAL blast radius. Direct chain reaches `GenerateAIContextFiles`, then `generateAnalyzeAIContext`, then `newAnalyzeCommand`. Affected files: `internal/aicontext/aicontext.go`, `internal/cli/analyze_postrun.go`, and `internal/cli/command.go`. Affected process examples include `NewAnalyzeCommand -> UpsertSection`, `NewAnalyzeCommand -> RemoveGeneratedSkills`, and `NewAnalyzeCommand -> AnalyzeAIContextResult`.

- `anvien context symbol "GenerateAIContextFiles" --repo Anvien`
  Result: symbol found in `internal/aicontext/aicontext.go`, range 61-92. It calls `renderAnvienBlock` for both `.agents/skills/` and `.claude/skills/`, confirming one generator change covers both target files.

## E1 - Implementation Evidence

- P1-A source change:
  `internal/aicontext/aicontext.go` now calls `renderMasterRulesBlock()` at the start of `renderAnvienBlock`, after `<!-- anvien:start -->` and before the existing `# Anvien - Code Intelligence` content. The helper contains the full user-provided `# Master iron rules` and `# AGENTS Rules` sections, including rules 0 through 10. Because `GenerateAIContextFiles` calls `renderAnvienBlock` for both `.agents/skills/` and `.claude/skills/`, the rule block is generated into both `AGENTS.md` and `CLAUDE.md`.

- P2-A test change:
  `internal/aicontext/aicontext_test.go` now has `requireGeneratedMasterRules`, which asserts the generated master-rule block and AGENTS rules 0 through 10 are present. `TestGenerateAIContextFilesCreatesManagedContextAndSkillPackages` calls that helper for both generated `AGENTS.md` and generated `CLAUDE.md`.

## E2 - Validation Evidence

- `gofmt -w internal\aicontext\aicontext.go internal\aicontext\aicontext_test.go`
  Result: completed successfully.

- Full build sequence from repository root:
  Result: completed successfully with fail-fast PowerShell wrapper around the exact required command sequence.
  Commands executed in order:
  `cd .\anvien`; `npm install`; `npm run build`; `npm install -g .`; `Get-Command anvien`; `anvien version`; `cd ..`; `powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1`; `anvien version`; `anvien analyze . --force`.
  Version checks returned `1.2.5`.
  Vite emitted existing non-failing warnings about a static/dynamic import mix and large chunks.
  Final analyze in the full build reported 1407 files, 84478 nodes, 122964 relationships, 16570 file projection dependency edges, and 430 unresolved items.

- `go test ./internal/aicontext ./internal/cli -count=1`
  Result: passed. Package results: `ok github.com/tamnguyendinh/anvien/internal/aicontext 35.845s`; `ok github.com/tamnguyendinh/anvien/internal/cli 85.416s`.

- Generated output smoke after full build/analyze:
  `AGENTS.md` and `CLAUDE.md` both start their managed section with `# Master iron rules`, include `# AGENTS Rules`, and include rules 0 through 10 before the existing `# Anvien - Code Intelligence` content.

## E3 - Detect Changes Evidence

- Final pre-commit graph refresh:
  `anvien analyze . --force` completed successfully after report/notes docs were added. Graph reported 1408 files, 84487 nodes, 122973 relationships, 16570 file projection dependency edges, and 430 unresolved items.

- `anvien detect-changes --repo Anvien --scope all`
  Result: completed successfully.
  Summary risk: low.
  Changed app layers: backend 3, backend_test 2, docs 3.
  Changed functional areas: documentation 3, unknown 5.
  Changed implementation/test files reported by detect-changes: `internal/aicontext/aicontext.go`, `internal/aicontext/aicontext_test.go`, and tracked `docs/notes_decisions_log/notes_decisions_log_20260606.md`.
  Affected processes: none reported in final detect output.

## E4 - Commit Evidence

- Implementation checkpoint: `31c1e5b fix(aicontext): generate master rules`.
