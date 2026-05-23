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
| API route/tool/shape/impact commands | pending | pending | pending | Confirm exact MCP tool names and do not invent CLI names where no CLI command exists. |
| Query-health / graph-quality commands | pending | pending | pending | Confirm from source or freshly built binary; PATH binary may be stale. |
| Resolution/source-site/accuracy commands | pending | pending | pending | Confirm from source or freshly built binary; PATH binary may be stale. |
| `benchmark-compare` | pending | pending | pending | Confirm exact command name and scope. |
| group/cross-repo commands | pending | pending | pending | Confirm exact command names. |
| `serve` / Web runtime | pending | pending | pending | Local server/Web UI behavior. |
| `mcp` | pending | pending | pending | MCP server behavior. |
| `setup` | pending | pending | pending | Editor/MCP setup and skill installation. |
| `version` | pending | pending | pending | Version output. |
| package/runtime commands | pending | pending | pending | Build/package runtime assets. |
| wiki/hook commands | pending | pending | pending | Confirm current command names and behavior. |
| MCP resources | pending | pending | pending | `avmatrix://...` resources. |

## B1A - Binary Command Surface Mismatch Inventory

Status: preliminary baseline recorded

| Surface | Command | Observed result | Interpretation |
|---|---|---|---|
| PATH binary | `avmatrix --help` | Did not list `query-health`, `resolution-inventory`, or `source-site-accuracy`. | PATH binary is stale for this plan's command inventory. |
| PATH binary | `avmatrix query-health --help` | `unknown command "query-health" for "avmatrix"` | Do not use PATH binary as source of truth. |
| PATH binary | `avmatrix resolution-inventory --help` | `unknown command "resolution-inventory" for "avmatrix"` | Do not use PATH binary as source of truth. |
| PATH binary | `avmatrix source-site-accuracy --help` | `unknown command "source-site-accuracy" for "avmatrix"` | Do not use PATH binary as source of truth. |
| Current source | `go run .\cmd\avmatrix --help` | Listed `query-health`, `resolution-inventory`, and `source-site-accuracy`. | Current source contains the newer command surface. |
| Current source | `go run .\cmd\avmatrix query-health --help` | Help output exists with `--fail-on-exact`, `--fail-on-threshold`, `--json`, `--limit`, `--out`, `--repo`, and `--suite`. | Skill content can document this after build/source verification. |
| Current source | `go run .\cmd\avmatrix resolution-inventory --help` | Help output exists with `--graph`, `--json`, and `--out`. | Skill content can document this after build/source verification. |
| Current source | `go run .\cmd\avmatrix source-site-accuracy --help` | Help output exists with `--golden`, `--graph`, `--json`, `--max-examples`, and `--out`. | Skill content can document this after build/source verification. |

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
| Package-root `skills/` source exists | not observed in repo root during preliminary review | pending | Setup currently reads package-root `skills/`; reconcile with embedded source. |
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
| MCP setup/resource guidance smoke | pending | pending | Required if MCP resource guidance is touched. |
| `detect-changes` | pending | pending | Scope should match AI context/skills/docs changes. |
