# Notes / Decisions Log - 2026-07-16

## Tree-sitter Go Module Drift

- Scope: implemented plan `docs/plans/2026-07-16-tree-sitter-go-module-drift`.
- Decision: Tree-sitter drift monitoring must follow Go module source of truth, not the retired npm `tree-sitter` runtime dependency model.
- Evidence: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-evidence.md`.
- Coder report: `reports/coder/rp_coder_260716_144412_by_gpt-5-codex_tree-sitter-go-module-drift.md`.
- Supervisor report: `reports/Supervisor/rp_supervisor_260716_144532_by_gpt-5-codex_tree-sitter-go-module-drift.md`.
- Implementation commit: `22fee7030dae322f0589907000ca45f067383e05`.

## Release Version 1.2.8

- Scope: bumped Anvien CLI/runtime package version from `1.2.7` to `1.2.8` and updated root changelog.
- Decision: current release metadata must stay aligned across `anvien/package.json`, `anvien/package-lock.json`, `internal/version/version.go`, `README.md`, and `RUNBOOK.md`.
- Evidence: `npm run full-build` passed and printed `anvien version` as `1.2.8`; `go test ./internal/version ./internal/cli ./internal/mcp` passed; `anvien detect-changes --repo Anvien --scope all` reported `risk_level=low`.
- Coder report: `reports/coder/rp_coder_260716_151122_by_gpt-5-codex_release-version-1-2-8.md`.
- Supervisor report: `reports/Supervisor/rp_supervisor_260716_151512_by_gpt-5-codex_release-version-1-2-8.md`.
- Implementation commit: `73612979fd9300fbc36ff47731099c8110962dab`.
