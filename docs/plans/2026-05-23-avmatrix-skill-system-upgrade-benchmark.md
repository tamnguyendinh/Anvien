# AVmatrix Skill System Upgrade Benchmark Ledger

Date: 2026-05-23

Status: Planned

Companion files:

- Plan: [2026-05-23-avmatrix-skill-system-upgrade-plan.md](2026-05-23-avmatrix-skill-system-upgrade-plan.md)
- Evidence ledger: [2026-05-23-avmatrix-skill-system-upgrade-evidence.md](2026-05-23-avmatrix-skill-system-upgrade-evidence.md)

## Benchmark Rules

Record benchmarkable results as each benchmarkable task is completed. Benchmarkable means measured product/runtime performance, capacity, package/startup size, graph/DB throughput, query hit rate, command output inventory, graph inventory counts, source-site inventory counts, generated skill inventory counts, setup/package file inventories, or resolved-edge accuracy.

Build/test/e2e timings are validation evidence unless the implementation changes those systems.

Do not count generated output as correct because it exists. Record generated file count, paths, sizes, representative content checks, and source-to-generated agreement.

## B0 - Source Skill Inventory

Status: pending

Record before and after implementation:

| Metric | Baseline | Final | Notes |
|---|---:|---:|---|
| Embedded source skill Markdown files under `internal/aicontext/skills` | pending | pending | Count source files, not generated files. |
| Registered base skills in `internal/aicontext/aicontext.go` | pending | pending | Count `baseSkills` entries. |
| Generated `.claude/skills/avmatrix/**/SKILL.md` files | pending | pending | Measured after normal generation. |
| Root generated Skills table rows | pending | pending | Measured in generated `AGENTS.md` or `CLAUDE.md`. |
| Source skill total bytes | pending | pending | Sum source Markdown bytes. |
| Generated skill total bytes | pending | pending | Sum generated local skill bytes after generation. |

## B1 - Command Surface Coverage Inventory

Status: pending

Record the current command/tool/resource inventory and final skill coverage:

| Command/tool/resource family | Implemented in current code? | Covered by final skill | Primary skill | Notes |
|---|---|---|---|---|
| `analyze` / graph refresh | pending | pending | pending | Base graph generation. |
| `status` | pending | pending | pending | Freshness and repo index state. |
| `query` | pending | pending | pending | Concept/process retrieval. |
| `context` | pending | pending | pending | Symbol/process neighborhood. |
| `impact` | pending | pending | pending | Blast radius. |
| `detect-changes` | pending | pending | pending | Git-diff graph impact. |
| `cypher` | pending | pending | pending | Read-only graph queries. |
| `rename` | pending | pending | pending | Graph-guided rename. |
| `augment` | pending | pending | pending | Confirm exact behavior from source. |
| API route/tool/shape/impact commands | pending | pending | pending | Confirm exact command and MCP tool names. |
| Query-health / graph-quality commands | pending | pending | pending | Confirm exact command names and output shape. |
| Resolution/source-site/accuracy commands | pending | pending | pending | Confirm exact command names and output shape. |
| `benchmark-compare` | pending | pending | pending | Confirm exact command name and scope. |
| group/cross-repo commands | pending | pending | pending | Confirm exact command names. |
| `serve` / Web runtime | pending | pending | pending | Local server/Web UI behavior. |
| `mcp` | pending | pending | pending | MCP server behavior. |
| `setup` | pending | pending | pending | Editor/MCP setup and skill installation. |
| `version` | pending | pending | pending | Version output. |
| package/runtime commands | pending | pending | pending | Build/package runtime assets. |
| wiki/hook commands | pending | pending | pending | Confirm current command names and behavior. |
| MCP resources | pending | pending | pending | `avmatrix://...` resources. |

## B2 - Generated Output Inventory

Status: pending

Record after normal generation through `avmatrix analyze --force`:

| Generated artifact | Expected final count/content | Observed final result | Pass/fail |
|---|---|---|---|
| `AGENTS.md` managed AVmatrix block | Final skill table, repo-agnostic guidance, current stats line | pending | pending |
| `CLAUDE.md` managed AVmatrix block | Same final skill table and guidance | pending | pending |
| `.claude/skills/avmatrix/<skill>/SKILL.md` | One file for each final base skill | pending | pending |
| Skill table/source registry agreement | Generated table matches `baseSkills` final set | pending | pending |
| Generated skill/source content agreement | Generated files match embedded source content except expected path wrapping | pending | pending |

## B3 - Setup And Package Inventory

Status: pending

Record if setup/package behavior is touched or tested:

| Surface | Baseline | Final | Notes |
|---|---:|---:|---|
| Editor skill install targets checked | pending | pending | Cursor, Claude Code, OpenCode, Codex if available in setup path. |
| Skills copied per editor target | pending | pending | Must match final base skill count. |
| Packaged embedded skills | pending | pending | Count packaged source skill files if package flow materializes them. |
| Package/runtime size delta | pending | pending | Only record if package build is run. |

## B4 - Validation Command Results

Status: pending

Record final validation:

| Command | Result | Pass/fail count | Notes |
|---|---|---:|---|
| `powershell -ExecutionPolicy Bypass -File avmatrix-launcher\build.ps1` | pending | pending | Full build gate before tests. |
| `go test .\internal\aicontext .\internal\cli -count=1` | pending | pending | Minimum focused test scope. |
| `avmatrix analyze --force` | pending | pending | Must not use `--skip-agents-md`. |
| Generated skill inventory check | pending | pending | Count and content fragments. |
| Setup/package smoke commands | pending | pending | Required if touched. |
| `detect-changes` | pending | pending | Scope should match AI context/skills/docs changes. |

