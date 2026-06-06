# Planner Four-Template Standard Benchmark

## Benchmark Ledger

| Time | Phase | Metric | Before | Target / After |
|------|-------|--------|--------|----------------|
| 2026-06-06 | P0 | Planner standard artifact count | 3 files | 4 files |
| 2026-06-06 | P1 | Bundled planner template files | 0 | 4 |
| 2026-06-06 | P0 | Graph inventory before docs/template changes | `files=1409`, `nodes=84018`, `relationships=122400` | Baseline |
| 2026-06-06 | P4 | Graph inventory after full build analyze | `files=1417`, `documents=494`, `nodes=84089`, `relationships=122471` | Expected increase from new docs/template/plan files |
| 2026-06-06 | P4 | Graph inventory after final template full build analyze | `files=1417`, `documents=494`, `nodes=84083`, `relationships=122465` | Final inventory after stronger actual-status template |
| 2026-06-06 | P3 | Generated planner template files in `.agents` | 0 | 4 |
| 2026-06-06 | P3 | Generated planner template files in `.claude` | 0 | 4 |
| 2026-06-06 | P1 | Actual-status relationship fields | 0 dedicated relationship-count sections | 2 sections: Relationship / Impact Evidence, Phase Touch Map |
| 2026-06-06 | P4 | Graph inventory after relationship-count full build analyze | `files=1417`, `documents=494`, `nodes=84085`, `relationships=122467` | Final inventory after adding relationship-count guidance |

## Notes

This change is mostly documentation and packaging surface. Build/test timings are validation evidence, not product performance benchmarks.
