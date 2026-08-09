# Anvien Graph Accuracy Multi-Plan Documentation Correction Benchmark Ledger

## Metadata

- Date: `2026-08-09`
- Last revised: `2026-08-10`
- Plan: `docs/plans/2026-08-09-anvien-graph-accuracy-multi-plan-correction/2026-08-09-anvien-graph-accuracy-multi-plan-correction-plan.md`
- Evidence: `docs/plans/2026-08-09-anvien-graph-accuracy-multi-plan-correction/2026-08-09-anvien-graph-accuracy-multi-plan-correction-evidence.md`
- Benchmark: `docs/plans/2026-08-09-anvien-graph-accuracy-multi-plan-correction/2026-08-09-anvien-graph-accuracy-multi-plan-correction-benchmark.md`
- Actual status: `docs/plans/2026-08-09-anvien-graph-accuracy-multi-plan-correction/2026-08-09-anvien-graph-accuracy-multi-plan-correction-actual-status.md`

## Benchmark Rules

- Record measured documentation inventory and consistency counts only.
- Product accuracy, graph size, analyze time, query latency, memory, and child behavior measurements remain in the owning child ledgers.
- Build, link-check, diff-check, review, and commit PASS/FAIL results belong in the evidence ledger.
- A final count is valid only after the final-path and full 34-file audit.

## B0 - P0 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P0 | active contract files | files | 1 | 1 | 1 | 1 | 0 | `E0-P0A-INVENTORY1`, `E3-P3A-INVENTORY1` |
| P0 | active roadmap files | files | 1 | 1 | 1 | 1 | 0 | `E0-P0A-INVENTORY1`, `E3-P3A-INVENTORY1` |
| P0 | child plan directories | directories | 7 | 7 | 7 | 7 | 0 | `E0-P0A-INVENTORY1`, `E3-P3A-INVENTORY1` |
| P0 | child standard ledgers | files | 28 | 28 | 28 | 28 | 0 | `E0-P0A-INVENTORY1`, `E3-P3A-INVENTORY1` |
| P0 | correction standard ledgers | files | 4 | 4 | 4 | 4 | 0 | `E0-P0A-INVENTORY1`, `E3-P3A-INVENTORY1` |
| P0 | total active authority files | files | 34 | 34 | 34 | 34 | 0 | `E0-P0A-INVENTORY1`, `E3-P3A-INVENTORY1` |
| P0 | Anvien scanned files | files | 1,556 | 1,557 | 1,557 | record final with explanation | +1 Supervisor report | `E0-P0A-GRAPH1`, `E3-P3C-ANALYZE1` |
| P0 | Anvien parsed code files | files | 676 | 676 | 676 | no unexplained code loss | 0 | `E0-P0A-GRAPH1`, `E3-P3C-ANALYZE1` |
| P0 | Anvien failed files | files | 0 | 0 | 0 | 0 | 0 | `E0-P0A-GRAPH1`, `E3-P3C-ANALYZE1` |

## B1 - P1 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P1 | root authority files corrected | files / files | 0/2 | 2/2 | 2/2 | 2/2 | +2/2 | `E1-P1A-CONTRACT1`, `E1-P1B-ROADMAP1` |
| P1 | root path entries resolved | resolved / active root references | 0/unknown | 100% | 100% | 100% | complete | `E1-P1B-PATH1`, `E3-P3A-LINK1` |
| P1 | roadmap implementation slices | slices | unreconciled | 35 | 35 | 35 | complete | `E1-P1B-SLICE1`, `E3-P3B-SLICE1` |

## B2 - P2 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P2 | child four-ledger sets corrected | sets / sets | 0/7 | 7/7 | 7/7 | 7/7 | +7/7 | `E2-P2A-CHILD1` through `E2-P2G-CHILD7` |
| P2 | Child 01 implementation slices | slices | unreconciled | 5 | 5 | 5 | complete | `E2-P2A-LEDGER1` |
| P2 | Child 02 implementation slices | slices | unreconciled | 5 | 5 | 5 | complete | `E2-P2B-LEDGER1` |
| P2 | Child 03 implementation slices | slices | unreconciled | 7 | 7 | 7 | complete | `E2-P2C-LEDGER1` |
| P2 | Child 04 implementation slices | slices | unreconciled | 5 | 5 | 5 | complete | `E2-P2D-LEDGER1` |
| P2 | Child 05 implementation slices | slices | unreconciled | 4 | 4 | 4 | complete | `E2-P2E-LEDGER1` |
| P2 | Child 06 implementation slices | slices | unreconciled | 6 | 6 | 6 | complete | `E2-P2F-LEDGER1` |
| P2 | Child 07 implementation slices | slices | unreconciled | 3 | 3 | 3 | complete | `E2-P2G-LEDGER1` |
| P2 | total implementation slices | slices | unreconciled | 35 | 35 | 35 | complete | `E3-P3B-SLICE1` |

## B3 - P3 Benchmarks

| Phase | Metric | Unit | Baseline | Latest | Final | Target | Delta | Evidence |
|-------|--------|------|----------|--------|-------|--------|-------|----------|
| P3 | active files reviewed | files / files | 0/34 | 34/34 | 34/34 | 34/34 | complete | `E3-P3A-INVENTORY1` |
| P3 | active local links unresolved | links | unknown | 0 | 0 | 0 | complete | `E3-P3A-LINK1` |
| P3 | slug/companion mismatches | mismatches | unknown | 0 | 0 | 0 | complete | `E3-P3A-SLUG1` |
| P3 | duplicate or orphan implementation slice IDs | IDs | unknown | 0 | 0 | 0 | complete | `E3-P3B-SLICE1` |
| P3 | missing, duplicate, or orphan evidence IDs | IDs | unknown | 0 | 0 | 0 | complete | `E3-P3B-EVIDENCE1` |
| P3 | contradictory status decisions | rows | unknown | 0 | 0 | 0 | complete | `E3-P3B-STATUS1` |
| P3 | cross-child ownership conflicts | conflicts | unknown | 0 | 0 | 0 | complete | `E3-P3B-SCOPE1` |
| P3 | benchmark rows outside owning scope | rows | unknown | 0 | 0 | 0 | complete | `E3-P3B-BENCH1` |

## Non-Benchmarkable Notes

- The five graph-accuracy target ratios belong to their owning child and terminal Child 07 ledgers, not this documentation-correction benchmark.
- Supervisor review, `git diff --check`, local-link validation result, cleanup result, commit hash, and worktree state are evidence gates rather than measurements.
