# Notes / Decisions Log - 2026-06-06

## Package Runtime Skill Subtree Staging
- Report: `reports/coder/rp_coder_260606_021558_by_gpt-5-codex_package-runtime-skill-subtree-staging.md`.
- Plan: `docs/plans/2026-06-06-anvien-package-runtime-skill-subtree-staging/2026-06-06-anvien-package-runtime-skill-subtree-staging-plan.md`.
- Commit: `fc6f0c2 fix(package): stage skill subtree in runtime source`.
- Decision: package runtime staging must copy `internal/aicontext/skills` as a subtree artifact, not as extension-filtered `.md` files inside Go-source traversal.
- Evidence: full build sequence passed on fail-fast rerun, `go test ./internal/cli -count=1` passed, repo-local package prepare smoke staged 638 skill files including 295 non-`.md` files, and `anvien detect-changes --repo Anvien --scope all` passed.
- Residual unverified surfaces: none.
