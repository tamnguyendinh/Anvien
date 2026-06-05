# Notes / Decisions Log - 2026-06-06

## Package Runtime Skill Subtree Staging
- Report: `reports/coder/rp_coder_260606_021558_by_gpt-5-codex_package-runtime-skill-subtree-staging.md`.
- Plan: `docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-plan.md`.
- Commit: `fc6f0c2 fix(package): stage skill subtree in runtime source`.
- Decision: package runtime staging must copy `internal/aicontext/skills` as a subtree artifact, not as extension-filtered `.md` files inside Go-source traversal.
- Evidence: full build sequence passed on fail-fast rerun, `go test ./internal/cli -count=1` passed, repo-local package prepare smoke staged 638 skill files including 295 non-`.md` files, and `anvien detect-changes --repo Anvien --scope all` passed.
- Residual unverified surfaces: none.

## AI Context Master Rules Generation
- Report: `reports/coder/rp_coder_260606_023426_by_gpt-5-codex_aicontext-master-rules-generation.md`.
- Plan: `docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-plan.md`.
- Commit: `31c1e5b fix(aicontext): generate master rules`.
- Decision: the repository master rule block must be emitted by `internal/aicontext/aicontext.go` into generated `AGENTS.md` and `CLAUDE.md`, not maintained only as hand-written repo text.
- Evidence: full build sequence passed, `go test ./internal/aicontext ./internal/cli -count=1` passed, generated output contains `# Master iron rules`, `# AGENTS Rules`, and rules 0 through 10 in both target files, and `anvien detect-changes --repo Anvien --scope all` passed with risk `low`.
- Residual unverified surfaces: none.

## AI Context Rule 1 Help Wording
- Report: `reports/coder/rp_coder_260606_030733_by_gpt-5-codex_aicontext-rule1-help-wording.md`.
- Plan: `docs/plans/2026-06-06-aicontext-rule1-help-wording/2026-06-06-aicontext-rule1-help-wording-plan.md`.
- Commit: `8aa2493 fix(aicontext): update generated rule help wording`.
- Decision: generated rule 1 should read `1. How to use anvien: run command "anvien --help".`
- Evidence: full build sequence passed after stopping the single MCP process that locked `anvien\bin\anvien.exe`, `go test ./internal/aicontext ./internal/cli -count=1` passed, generated output smoke found the new wording in both `AGENTS.md` and `CLAUDE.md`, and `anvien detect-changes --repo Anvien --scope all` passed with risk `low`.
- Residual unverified surfaces: none.
