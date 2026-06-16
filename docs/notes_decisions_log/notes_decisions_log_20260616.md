# Notes / Decisions Log - 2026-06-16

## File Detail Absolute Path Resolution

- Report: `reports/coder/rp_coder_260616_140522_by_gpt-5-codex_file-detail-absolute-path-resolution.md`.
- Supervisor: `reports/Supervisor/rp_supervisor_260616_140659_by_gpt-5-codex_file-detail-absolute-path-resolution.md` (`PASS`).
- Plan: `docs/plans/2026-06-16-file-detail-absolute-path-resolution/2026-06-16-file-detail-absolute-path-resolution-plan.md`.
- Commits: `bfbc3d6`, `9a27930`, `8499041`, `783c809`, `09aa58c`.
- Decision: graph file paths stay repo-relative; CLI, HTTP, and MCP boundaries normalize user file input against the resolved repo root before file-context lookup.
- Evidence: helper/CLI/HTTP/MCP focused tests passed, applicable `go build ./cmd/... ./internal/...` and `go test ./cmd/... ./internal/... -count=1` passed, manual CLI relative/absolute lookup returned the same normalized path and symbol count, outside-repo absolute lookup exited with a clear error, and final `anvien detect-changes --repo Anvien --scope all` reported no changes.
- Caveat: repo-wide `go build ./...` still fails on known intentionally non-buildable analyzer fixtures under `anvien/test/fixtures`; not caused by this scope.
- Residual unverified surfaces: none for the file target lookup invariant.
