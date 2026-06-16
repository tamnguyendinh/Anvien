# Community Singleton MEMBER_OF Plan

## Metadata

- Date: `2026-06-16`
- Status: `complete`
- Plan: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-plan.md`
- Evidence: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-evidence.md`
- Benchmark: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-benchmark.md`
- Actual status: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-actual-status.md`

## Goal

Fix the community detection bug that emits `MEMBER_OF` relationships pointing to skipped singleton community IDs, causing `anvien analyze --force` to fail during LadybugDB load.

## Scope

In scope:

- `internal/communities/communities.go`
- focused community regression tests
- analyze/load validation for current repo

Out of scope:

- skill directory classification
- LadybugDB schema changes
- broad analyzer refactors

## Requirements

- Preserve the existing intent that singleton communities are skipped.
- Preserve graph integrity: every emitted relationship must point to existing source and target nodes.
- Keep the fix local to community emission unless validation proves another owner.

## Checklist

- [x] P0-A: Confirm actual status and root cause before editing.
  - Goal: classify the failing state and identify the owner.
  - Work Steps: reproduce analyze failure, compare with the last known good commit, isolate the cyber skill trigger, and identify the dangling relationship pattern.
  - Implementation Gate: source owner and required behavior are known.
  - Acceptance: actual status records `MEMBER_OF` edges to missing `comm_N` targets as the current bug.

- [x] P1-A: Stop emitting membership edges for skipped singleton communities.
  - Goal: align membership edge creation with community node creation.
  - Work Steps: run pre-edit impact if available, move the singleton guard before `MEMBER_OF` emission, and keep non-singleton behavior unchanged.
  - Implementation Gate: impact evidence is recorded, or the current broken index blocker is recorded.
  - Acceptance: singleton partitions do not create `MEMBER_OF` edges; non-singleton communities still create nodes and memberships.

- [x] P1-B: Add focused regression coverage.
  - Goal: prove singleton community skipping leaves no dangling relationship.
  - Work Steps: update the existing singleton community test or add a sibling assertion after code behavior is corrected.
  - Implementation Gate: code behavior is fixed first.
  - Acceptance: test fails on the old behavior and passes with the fix.

- [x] P2-A: Validate and close.
  - Goal: prove the repo can analyze cleanly again and record final evidence.
  - Work Steps: run full build before validation, run targeted tests, run `anvien analyze --force`, run change detection before commit, and record results.
  - Implementation Gate: P1-A and P1-B are complete.
  - Acceptance: `anvien analyze --force` completes with zero skipped DB relationships or another explicit blocker is recorded.

## Risk Notes

- `communities.Apply` is shared graph-builder behavior; keep the change surgical and verify current community tests.
- Current repo index is unavailable because the bug blocks analyze; impact commands may fail until the fix lands.
