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

Status: recorded; pre-fix command evidence unavailable in this ledger

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

| Metric | Pre-fix evidence | After fix | Target |
|---|---|---:|---|
| Default analyze `AGENTS.md` length | reported as `0`; direct command output was not captured in this ledger before preliminary generation | `3676` | greater than `0`, one managed block |
| Default analyze `CLAUDE.md` length | reported as `0`; direct command output was not captured in this ledger before preliminary generation | `3676` | greater than `0`, one managed block |
| Root block line count after `--skills` smoke | pending | `55` | enough to include accurate tools, resources, skills, and CLI fallback without bloating agent context |

Interpretation:

The root files being non-empty is necessary but not sufficient. Final acceptance depends on accurate content and flag behavior as well as byte length.

## B1 - Base Skill Size Baseline

Date: 2026-05-20

Status: final sizes recorded

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
| `avmatrix-cli` | `432` | `2669` | embedded rich package content |
| `avmatrix-debugging` | `290` | `2168` | embedded rich package content |
| `avmatrix-exploring` | `398` | `2274` | embedded rich package content |
| `avmatrix-guide` | `338` | `2374` | embedded rich package content |
| `avmatrix-impact-analysis` | `368` | `2283` | embedded rich package content |
| `avmatrix-refactoring` | `290` | `1934` | embedded rich package content |

Interpretation:

The current Go base skill output is placeholder-level. A final fix should materially increase useful content or explicitly justify a different skill packaging strategy.

## B2 - Analyze Smoke Graph Baseline

Date: 2026-05-20

Status: final rerun recorded

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
| Files scanned | `694` | `703` | changed because generated/ignored artifacts and repo state changed during validation |
| Files parsed | `530` | `530` | unchanged parsed source count |
| Files unsupported | `164` | `173` | changed with ignored/generated artifact inventory |
| Files failed | `0` | `0` | remains clean |
| Graph nodes | `20,989` | `21,090` | explainable by current repo state after implementation files were added |
| Graph relationships | `52,322` | `52,444` | explainable by current repo state after implementation files were added |

Interpretation:

AI context generation should not change graph semantics. Graph inventory differences after the fix must be explained by source changes or analyzer behavior, not by root context file generation.

## B3 - Generated Community Skill Benchmark

Status: recorded

Required command:

```powershell
go run .\cmd\avmatrix analyze --force --skills --no-stats
Get-ChildItem -LiteralPath .claude\skills\generated -Recurse -Filter SKILL.md -ErrorAction SilentlyContinue | Measure-Object
Get-ChildItem -LiteralPath .claude\skills\generated -Recurse -Filter SKILL.md -ErrorAction SilentlyContinue | Measure-Object -Property Length -Sum
```

Required measurements:

| Metric | Default analyze | Analyze with `--skills` | Interpretation |
|---|---:|---:|---|
| Generated community skill count | `0` when directory absent before default smoke | `20` | default does not create generated community skills; `--skills` does |
| Generated community total bytes | `0` | `59,690` | only `--skills` creates generated files |
| Generated root file length | `3676` | `5376` | `--skills` adds generated skill rows |

## B4 - Generated Content Accuracy Inventory

Status: recorded

Required final inventory:

| Item | Expected final state | Evidence |
|---|---|---|
| Skills section heading | `## Skills` or equivalent non-CLI heading | `AGENTS.md` smoke output and `aicontext` test assertions |
| CLI fallback section | real `avmatrix` commands if included | `avmatrix analyze --force`, `status`, `query`, `context`, `impact`, and `detect-changes` present |
| MCP tool names | match Go MCP surface | `go test ./internal/mcp` passed |
| MCP resource URIs | match Go MCP surface or documented omissions | resources include `repos`, `setup`, `context`, `clusters`, `processes`, `schema`, `cluster/{name}`, and `process/{name}` |
| Stale index wording | internally consistent | stale warning points to `avmatrix analyze --force`, matching Always Do refresh rule |
| Generated skill paths | installed files exist | `--skills` produced `20` files under `.claude/skills/generated` |
| Base skill paths | installed files exist | six installed files under `.claude/skills/avmatrix` |

Interpretation:

This benchmark is an inventory of generated content shape rather than runtime performance. It is benchmarkable because generated content size and surface coverage are product outputs for this plan.

## B5 - Final Benchmark

Status: recorded

Final closure must record:

- final root `AGENTS.md` and `CLAUDE.md` byte lengths: `3676` default, `5376` with `--skills`;
- final base skill byte lengths: `1934` to `2669`;
- final generated community skill count and total bytes for `--skills`: `20` files, `59690` bytes;
- final analyze smoke graph counts: `21090` nodes and `52444` relationships;
- package size delta: not measured because source markdown assets were embedded and binary packaging size was not part of this slice's acceptance;
- final interpretation: generated root context is now non-empty by default, content headings match actual surfaces, base skills are no longer placeholders, and `--skills` behavior is separated from default root context generation.
