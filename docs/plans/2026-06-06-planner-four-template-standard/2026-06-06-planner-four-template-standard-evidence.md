# Planner Four-Template Standard Evidence

## Evidence Ledger

| Time | Phase | Evidence | Result |
|------|-------|----------|--------|
| 2026-06-06 | P0 | Read `internal/aicontext/skills/planner/SKILL.md`. | Current planner standard requires three files: plan, evidence, benchmark. No `actual-status.md` gate and no template usage. |
| 2026-06-06 | P0 | Read `internal/aicontext/skills/ui-be-binding-skill/docs/actual-wiring-status.template.md`. | Existing UI/BE actual-wiring template proves the useful pattern, but it is domain-specific and should be generalized for planner. |
| 2026-06-06 | P0 | Ran `anvien analyze --force`. | Graph refreshed: `files scanned=1409`, `documents=486`, `nodes=84018`, `relationships=122400`, stale false for later file-detail evidence. |
| 2026-06-06 | P0 | Ran `anvien file-detail internal/aicontext/skills/planner/SKILL.md --repo Anvien --json`. | Planner source is documentation-only, low risk, no code symbols or relationships reported. |
| 2026-06-06 | P0 | Searched planner/package references with `rg`. | `internal/aicontext/aicontext_test.go` protects planner guidance fragments; `internal/aicontext/skill_packages.go` already supports nested skill payloads; package command tests already use nested planner assets. |
| 2026-06-06 | P1 | User correction on test strategy. | Do not add content-locking tests for planner wording/templates. Planner content may evolve, and brittle text locks can break the system for the wrong reason. |
| 2026-06-06 | P1 | Updated `internal/aicontext/skills/planner/SKILL.md`. | Planner now defines a four-file standard plan set, requires bundled templates, and makes `actual-status.md` the P0 gate before implementation work. |
| 2026-06-06 | P1 | Added planner template assets under `internal/aicontext/skills/planner/templates/`. | Added `plan.template.md`, `evidence.template.md`, `benchmark.template.md`, and `actual-status.template.md`. |
| 2026-06-06 | P1 | User supplied a stronger `actual-status` structure. | Replaced `actual-status.template.md` with the stronger structure: metadata, allowed next action per status, per-unit current/required/evidence/decision fields, rewrite decision table, and explicit final P0 choices. |
| 2026-06-06 | P4 | First `npm run full-build`. | Failed because two running `anvien mcp` processes held `anvien.exe` locked. This was an environment lock, not a planner/template failure. |
| 2026-06-06 | P4 | Stopped the two `anvien mcp` processes that held the runtime binary. | Released the binary lock so the official full-build script could rebuild runtime. |
| 2026-06-06 | P4 | Reran `npm run full-build`. | Passed. Runtime rebuilt, web build passed, `anvien version` printed `1.2.6`, and final `anvien analyze . --force` completed. |
| 2026-06-06 | P3 | Inspected `.agents/skills/planner` and `.claude/skills/planner`. | Both generated mirrors contain updated `SKILL.md` plus all four templates. |
| 2026-06-06 | P2 | Ran `go test ./internal/aicontext`. | Passed: `ok github.com/tamnguyendinh/anvien/internal/aicontext 16.544s`. |
| 2026-06-06 | P3 | Reran `anvien analyze . --force` after strengthening the actual-status template. | Passed and regenerated generated context from source: `files=1417`, `documents=494`, `nodes=84083`, `relationships=122465`. |
| 2026-06-06 | P3 | Verified `.agents` and `.claude` actual-status template mirrors. | Both mirrors include `Allowed next action`, `Current Status Matrix`, `Final P0 Decision`, and the note that actual-status does not replace `evidence.md`. |
| 2026-06-06 | P4 | Reran `npm run full-build` after final template update. | Passed. Runtime rebuilt, web build passed, `anvien version` printed `1.2.6`, and final analyze completed. |
| 2026-06-06 | P2 | Reran `go test ./internal/aicontext` after final template update. | Passed: `ok github.com/tamnguyendinh/anvien/internal/aicontext 16.913s`. |
| 2026-06-06 | P5 | Ran `anvien detect-changes --repo Anvien --scope all`. | Risk LOW. Changed app layer `docs`, functional area `documentation`, affected processes none. Detect-changes reported touched planner `SKILL.md` sections; `git status` also shows new planner templates and plan artifacts to stage. |
| 2026-06-06 | P1 | User requested relationship count and impact clarity in actual-status. | Added `Relationship / Impact Evidence` and `Phase Touch Map` to `actual-status.template.md`; updated planner skill guidance to require `file-detail` relationship counts for file targets. |
| 2026-06-06 | P3 | Regenerated context after relationship-count template update with `anvien analyze . --force`. | Passed: `files=1417`, `documents=494`, `nodes=84085`, `relationships=122467`. |
| 2026-06-06 | P3 | Verified `.agents` and `.claude` mirrors after relationship-count update. | Both mirrors include `Relationship / Impact Evidence`, `Related File Count`, `Phase Touch Map`, and `file-detail` guidance. |
| 2026-06-06 | P4 | Reran `npm run full-build` after relationship-count update. | Passed. Runtime build, web build, install, `anvien version`, and final analyze all completed. |
| 2026-06-06 | P2 | Reran `go test ./internal/aicontext` after relationship-count update. | Passed: `ok github.com/tamnguyendinh/anvien/internal/aicontext 16.254s`. |
| 2026-06-06 | P5 | Reran `anvien detect-changes --repo Anvien --scope all` after final validation. | Risk LOW. Changed app layer `docs`, functional area `documentation`, affected processes none, resolution gap delta 0. |

## Decisions From Evidence

- Add templates under `internal/aicontext/skills/planner/templates/`.
- Update planner `SKILL.md` to require four files and to use bundled templates.
- Run existing focused tests after changing behavior; do not add brittle tests that lock planner wording.
- Do not edit generated root `AGENTS.md` / `CLAUDE.md` as source of truth.
