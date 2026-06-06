# Planner Four-Template Standard Plan

## Goal

Upgrade the planner skill so it creates a four-file planning set from bundled templates, including a required `actual-status.md` P0 reality check before implementation work.

## Files

- Plan: `docs/plans/2026-06-06-planner-four-template-standard/2026-06-06-planner-four-template-standard-plan.md`
- Evidence: `docs/plans/2026-06-06-planner-four-template-standard/2026-06-06-planner-four-template-standard-evidence.md`
- Benchmark: `docs/plans/2026-06-06-planner-four-template-standard/2026-06-06-planner-four-template-standard-benchmark.md`
- Actual status: `docs/plans/2026-06-06-planner-four-template-standard/2026-06-06-planner-four-template-standard-actual-status.md`

## Checklist

- [x] P0-A: Complete `actual-status.md` before implementation work.
  - Confirm the true current planner source.
  - Identify which surfaces are correct, partial, missing, fake/demo, or blocked.
  - Rewrite the implementation phases from that evidence.
- [x] P1-A: Update planner skill instructions and add four bundled template files.
  - Keep the planner workflow concise and mandatory.
  - Make `actual-status.md` the P0 gate before implementation work.
  - Tell agents to use templates instead of inferring from older plans.
- [x] P2-A: Validate with existing tests without locking planner content.
  - Do not add assertion tests that pin wording in planner skill/template content.
  - Rely on existing nested payload discovery coverage for packaging behavior.
  - Run focused existing tests after implementation.
- [x] P3-A: Regenerate and inspect generated agent skill mirrors.
  - Regenerate from the source generator instead of editing generated content.
  - Verify `.agents` and `.claude` receive the planner templates.
- [x] P4-A: Run full validation.
  - Run the repo full build through `npm run full-build`.
  - Run focused existing tests after full build.
- [x] P5-A: Run Anvien change detection and commit the completed slice.
  - Record detect-change evidence.
  - Commit scoped implementation changes.

## Phase Notes

P0 determines whether P1 must change only the planner skill, whether package machinery needs edits, or whether generation wiring needs edits. If P0 finds existing package machinery already copies nested skill files, preserve it. Do not add brittle content-locking tests for planner text; planner wording is allowed to evolve.
