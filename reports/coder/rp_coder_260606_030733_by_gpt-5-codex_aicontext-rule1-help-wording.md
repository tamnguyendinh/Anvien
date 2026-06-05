# Coder Report: AI Context Rule 1 Help Wording

Status: READY FOR SUPERVISOR REVIEW

## Metadata
- Report file: `reports/coder/rp_coder_260606_030733_by_gpt-5-codex_aicontext-rule1-help-wording.md`
- Time: 260606 030733 Asia/Bangkok
- Coder: gpt-5-codex
- Scope: implement `docs/plans/2026-06-06-aicontext-rule1-help-wording/2026-06-06-aicontext-rule1-help-wording-plan.md`
- Authority: `AGENTS.md`, `.agents/skills/coder/SKILL.md`, user-provided rule-1 replacement wording
- Implementation checkpoint: `8aa2493 fix(aicontext): update generated rule help wording`.

## Invariant Family Map
- Family name: generated AI context master rule wording.
- SSOT / authority source: `internal/aicontext/aicontext.go` generated rule block.
- Sibling runtime surfaces checked: `renderMasterRulesBlock`, `renderAnvienBlock`, generated `AGENTS.md`, generated `CLAUDE.md`, aicontext tests, CLI generation path.
- Forbidden fallback / alternate path: editing only generated `AGENTS.md` / `CLAUDE.md` without changing generator source.
- Stale tests/helpers/plans updated: `internal/aicontext/aicontext_test.go`, plan/evidence ledger.
- Verify matrix: full build sequence, generated output smoke, targeted Go tests, Anvien detect-changes.

## Files Changed
- `internal/aicontext/aicontext.go`: changed generated rule 1 to `1. How to use anvien: run command "anvien --help".`
- `internal/aicontext/aicontext_test.go`: updated generated rule assertion.
- `docs/plans/2026-06-06-aicontext-rule1-help-wording/*`: plan and evidence.
- `reports/coder/rp_coder_260606_030733_by_gpt-5-codex_aicontext-rule1-help-wording.md`: handoff report.
- `docs/notes_decisions_log/notes_decisions_log_20260606.md`: report link and verification summary.

## Verify Outputs
| Command | Result |
|---|---|
| Full build sequence from repo root | Pass after stopping the single MCP process that locked `anvien\bin\anvien.exe`. Version checks returned `1.2.5`; final analyze reported 1410 files, 84500 nodes, 122986 relationships, 16570 dependency edges, 430 unresolved. |
| `go test ./internal/aicontext ./internal/cli -count=1` | Pass. `internal/aicontext` 33.584s; `internal/cli` 97.688s. |
| Generated output smoke | Pass. `AGENTS.md` and `CLAUDE.md` both contain `1. How to use anvien: run command "anvien --help".` |
| Final pre-commit `anvien analyze . --force` | Pass. Final graph: 1411 files, 84509 nodes, 122995 relationships, 16570 dependency edges, 430 unresolved. |
| `anvien detect-changes --repo Anvien --scope all` | Pass. Summary risk `low`; affected processes: none reported. |

## E2E Flow
- Trigger: `anvien analyze . --force`.
- Process: analyze post-run calls AI context generation, which calls `GenerateAIContextFiles`, then `renderAnvienBlock`, then `renderMasterRulesBlock`.
- Observable result: generated `AGENTS.md` and `CLAUDE.md` contain the new rule-1 wording.

## Closure Notes
- Only rule 1 wording changed.
- Residual unverified surfaces: none.
- Risks/open points: direct Supervisor acceptance is still required before this scope is DONE.
