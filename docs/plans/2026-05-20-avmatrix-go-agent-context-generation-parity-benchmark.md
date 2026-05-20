# AVmatrix GO Agent Context Generation Parity Benchmark Ledger

Date: 2026-05-20

Status: active

Companion files:

- Plan: [2026-05-20-avmatrix-go-agent-context-generation-parity-plan.md](2026-05-20-avmatrix-go-agent-context-generation-parity-plan.md)
- Evidence ledger: [2026-05-20-avmatrix-go-agent-context-generation-parity-evidence.md](2026-05-20-avmatrix-go-agent-context-generation-parity-evidence.md)

## Benchmark Rules

Record benchmarkable results when they measure generated artifact size, graph inventory counts, package/runtime size, or product/runtime performance. Build and test timings are validation evidence unless the slice intentionally changes those systems.

For this plan, benchmarkable measurements include:

- generated `AGENTS.md` byte length;
- generated `CLAUDE.md` byte length;
- installed base skill byte lengths;
- generated community skill count;
- generated community skill total byte length;
- graph node and relationship counts from smoke analyze;
- generated root block line count if content accuracy changes materially;
- package size delta if rich skills are embedded into the Go binary;
- analyze runtime only if generation changes measurably affect analyze performance.

Do not record inferred or estimated sizes. Every benchmark row must name the command, repo path, and interpretation.

## B0 - Initial Root Context Baseline

Date: 2026-05-20

Status: initial observation recorded; final accepted baseline pending

Command:

```powershell
Get-Item -LiteralPath AGENTS.md,CLAUDE.md -ErrorAction SilentlyContinue | Select-Object Name,Length,LastWriteTime
```

Initial observed local file sizes after preliminary smoke:

| File | Length | Interpretation |
|---|---:|---|
| `AGENTS.md` | `2441` | non-empty after preliminary local generation |
| `CLAUDE.md` | `2441` | non-empty after preliminary local generation |

Required final benchmark:

| Metric | Before fix | After fix | Target |
|---|---:|---:|---|
| Default analyze `AGENTS.md` length | `0` reported/observed pre-fix in this repo | pending final validation | greater than `0`, one managed block |
| Default analyze `CLAUDE.md` length | `0` reported/observed pre-fix in this repo | pending final validation | greater than `0`, one managed block |
| Root block line count | pending | pending | enough to include accurate tools, resources, skills, and CLI fallback without bloating agent context |

Interpretation:

The root files being non-empty is necessary but not sufficient. Final acceptance depends on accurate content and flag behavior as well as byte length.

## B1 - Base Skill Size Baseline

Date: 2026-05-20

Status: recorded

Current Go generated base skill sizes after preliminary smoke:

| Skill | Current Go generated length |
|---|---:|
| `avmatrix-cli` | `432` |
| `avmatrix-debugging` | `290` |
| `avmatrix-exploring` | `398` |
| `avmatrix-guide` | `338` |
| `avmatrix-impact-analysis` | `368` |
| `avmatrix-refactoring` | `290` |

Original TypeScript package skill sizes:

| Skill | Original package length |
|---|---:|
| `avmatrix-cli` | `3223` |
| `avmatrix-debugging` | `3120` |
| `avmatrix-exploring` | `3022` |
| `avmatrix-guide` | `3476` |
| `avmatrix-impact-analysis` | `2931` |
| `avmatrix-refactoring` | `4007` |

Required final benchmark:

| Skill | Baseline Go length | Final Go generated length | Interpretation |
|---|---:|---:|---|
| `avmatrix-cli` | `432` | pending | should be rich package content or documented replacement |
| `avmatrix-debugging` | `290` | pending | should be rich package content or documented replacement |
| `avmatrix-exploring` | `398` | pending | should be rich package content or documented replacement |
| `avmatrix-guide` | `338` | pending | should be rich package content or documented replacement |
| `avmatrix-impact-analysis` | `368` | pending | should be rich package content or documented replacement |
| `avmatrix-refactoring` | `290` | pending | should be rich package content or documented replacement |

Interpretation:

The current Go base skill output is placeholder-level. A final fix should materially increase useful content or explicitly justify a different skill packaging strategy.

## B2 - Analyze Smoke Graph Baseline

Date: 2026-05-20

Status: preliminary observation recorded; final rerun pending

Command:

```powershell
go run .\cmd\avmatrix analyze --force --no-stats
```

Preliminary smoke result:

| Metric | Value |
|---|---:|
| Files scanned | `694` |
| Files parsed | `530` |
| Files unsupported | `164` |
| Files failed | `0` |
| Graph nodes | `20,989` |
| Graph relationships | `52,322` |

Required final smoke benchmark:

| Metric | Preliminary | Final | Interpretation |
|---|---:|---:|---|
| Files scanned | `694` | pending | should remain explainable by repo state |
| Files parsed | `530` | pending | should remain explainable by repo state |
| Files failed | `0` | pending | must remain `0` |
| Graph nodes | `20,989` | pending | generation changes should not unexpectedly perturb graph facts |
| Graph relationships | `52,322` | pending | generation changes should not unexpectedly perturb graph facts |

Interpretation:

AI context generation should not change graph semantics. Graph inventory differences after the fix must be explained by source changes or analyzer behavior, not by root context file generation.

## B3 - Generated Community Skill Benchmark

Status: pending

Required command:

```powershell
go run .\cmd\avmatrix analyze --force --skills --no-stats
Get-ChildItem -LiteralPath .claude\skills\generated -Recurse -Filter SKILL.md -ErrorAction SilentlyContinue | Measure-Object
Get-ChildItem -LiteralPath .claude\skills\generated -Recurse -Filter SKILL.md -ErrorAction SilentlyContinue | Measure-Object -Property Length -Sum
```

Required measurements:

| Metric | Default analyze | Analyze with `--skills` | Interpretation |
|---|---:|---:|---|
| Generated community skill count | pending | pending | default should not create generated community skills |
| Generated community total bytes | pending | pending | only meaningful when `--skills` creates files |
| Generated root skill rows | pending | pending | rows should match generated skill files |

## B4 - Generated Content Accuracy Inventory

Status: pending

Required final inventory:

| Item | Expected final state | Evidence |
|---|---|---|
| Skills section heading | `## Skills` or equivalent non-CLI heading | pending |
| CLI fallback section | real `avmatrix` commands if included | pending |
| MCP tool names | match Go MCP surface | pending |
| MCP resource URIs | match Go MCP surface or documented omissions | pending |
| Stale index wording | internally consistent | pending |
| Generated skill paths | installed files exist | pending |
| Base skill paths | installed files exist | pending |

Interpretation:

This benchmark is an inventory of generated content shape rather than runtime performance. It is benchmarkable because generated content size and surface coverage are product outputs for this plan.

## B5 - Final Benchmark

Status: pending

Final closure must record:

- final root `AGENTS.md` and `CLAUDE.md` byte lengths;
- final base skill byte lengths;
- final generated community skill count and total bytes for `--skills`;
- final analyze smoke graph counts;
- any package or binary size delta if rich skills are embedded;
- final interpretation of what changed and why it is acceptable.
