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

## B1Q - Query Reliability Bug Inventory

Status: pending

Dedicated suite for this plan:

| Artifact | Baseline | Final | Notes |
|---|---|---|---|
| Query-health suite path | pending | `docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json` or recorded equivalent | Keep this plan's query cases separate from older suites with different product goals. |
| Suite case count | pending | pending | Must include AI-context owner discovery and representative non-AI-context query capability lanes. |
| Suite validates query as umbrella command | pending | pending | Cases must not define `query` as only an AI-context lookup. |

Query behavior contract for this plan:

| Behavior | Baseline measurement | Final requirement |
|---|---|---|
| Broad concept discovery | pending | `query` still returns relevant code/process candidates for broad intent. |
| Exact owner discovery | pending | Exact owner files/symbols outrank weak-overlap process flows for the AI-context case. |
| Execution-flow preservation | pending | Useful process/flow results remain present when they have clear query overlap. |
| Wrong-owner demotion | pending | Plausible but unrelated launcher/resolution/frontend/backend flows do not outrank stronger AI-context owner evidence. |
| Result auditability | pending | Query or query-health output exposes match reason evidence and noise/miss reasons. |
| `query` versus `context` separation | pending | `query` remains broad discovery; `context` remains exact symbol inspection. |
| Query lane evidence | pending | Results expose query lane or equivalent match reason when available. |
| Source/global ranking semantics | pending | Hit@5 and hit@10 are based on documented result ordering. |
| User-facing lane discovery | pending | CLI/MCP help or command output lets users and agents discover query lanes. |
| Explainable query output | pending | CLI/MCP JSON output exposes lane/rank/match evidence without reading internal code. |
| Normal query compatibility | pending | Existing `avmatrix query "<intent>" --repo <repo>` behavior still works. |

Query Capability Taxonomy coverage:

| Query capability | Baseline case | Final case | Result fields to record |
|---|---|---|---|
| owner discovery | pending | pending | expected owners, matched owners, missed owners, rank, reason |
| concept discovery | pending | pending | top code areas, lane/reason evidence, unrelated result count |
| execution-flow discovery | pending | pending | process/flow candidates, matched steps, flow overlap evidence |
| API surface discovery | pending | pending | handlers/contracts/consumers, route/tool evidence |
| graph-quality discovery | pending | pending | query-health/resolution/source-site/graph-health owners |
| docs/setup/AI-context discovery | pending | pending | generator/setup/package/resource guidance owners |
| command-surface discovery | pending | pending | CLI/MCP/resource/Web/API command owners |
| cross-repo discovery | pending | pending | group/cross-repo owners when indexed data supports them |

Record broad-intent query behavior before and after implementation:

| Query intent | Expected owner targets | Baseline result | Final result | Notes |
|---|---|---|---|---|
| AI context generated skills and `AGENTS.md`/`CLAUDE.md` | `internal/aicontext/aicontext.go`, `internal/aicontext/skills/*.md`, `internal/cli/analyze_postrun.go` | pending | pending | Broad query previously returned unrelated launcher/resolution/frontend flows; measure instead of relying on memory. |
| Setup/editor skill installation | `internal/cli/setup_command.go`, `setupInstallSkillsTo`, package-root `skills/` handling | pending | pending | Verify query can surface setup owner or record exact misses. |
| Package skill distribution | `internal/cli/package_command.go`, `internal/cli/package_runtime.go`, package metadata/source-copy behavior | pending | pending | Package commands are hidden today; query-health should still make owner discovery auditable. |
| MCP setup/resource guidance | `internal/mcp/resources.go`, `setupResource`, resource/tool reference tests | pending | pending | Verify broad query can locate MCP-facing guidance owner. |

Root-cause and fix metrics:

| Metric | Baseline | Final |
|---|---:|---:|
| unrelated top results before expected owner | pending | pending |
| expected owner files in top 5 | pending | pending |
| expected owner files in top 10 | pending | pending |
| expected owner symbols in top 10 | pending | pending |
| top result has match reason evidence | pending | pending |
| missed targets reported explicitly | pending | pending |
| noise reason reported explicitly | pending | pending |
| broad-discovery regression check passes | pending | pending |
| useful execution-flow candidates preserved | pending | pending |
| query lane coverage recorded | pending | pending |
| source/global rank behavior recorded | pending | pending |
| CLI lane/explain commands validated | pending | pending |
| MCP query evidence fields validated | pending | pending |
| existing query command compatibility validated | pending | pending |

Record final query-health fields for the AI-context intent case:

| Metric | Final |
|---|---:|
| threshold pass/fail | pending |
| exact pass/fail | pending |
| usable retrieval meaning | pending |
| exact coverage meaning | pending |
| expected targets | pending |
| matched targets | pending |
| missed targets | pending |
| noise reason | pending |
| query lane / match reason | pending |
| source rank and global rank behavior | pending |

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

## B1G - Graph Labeling And Visual Orientation Inventory

Status: pending

Baseline problem artifact:

| Artifact | Baseline | Final | Notes |
|---|---|---|---|
| `reports/problem/screenshot_1779517751.png` | present | pending | Shows visible graph rings/islands without direct on-canvas names for the macro rings or each island. |

Record before and after the graph labeling phase:

| Metric | Baseline | Final | Notes |
|---|---:|---:|---|
| Runtime-visible macro ring count | pending | pending | Count from browser diagnostics or graph label metadata. |
| Macro rings with on-canvas labels | pending | pending | Must match visible macro rings unless a ring is intentionally hidden by filter. |
| Runtime-visible island count | pending | pending | Count per macro ring where possible. |
| Islands with on-canvas labels | pending | pending | Must cover major visible node islands and investigation islands such as ResolutionGap/Unresolved/Unknown where present. |
| Labels sourced from graph metadata | pending | pending | Record whether labels came from app layer, node type/filter, island key, semantic group, or fallback. |
| Labels update after filters/depth changes | pending | pending | Ring/island label count must match the currently visible graph subset, not stale initial conversion state. |
| Runtime diagnostics/test selector label count | pending | pending | Browser validation must have a machine-checkable count or selector in addition to screenshot review. |
| Desktop screenshot label readability | pending | pending | Browser screenshot must show ring and island names without relying on hover or side panel text. |
| Smaller viewport screenshot label readability | pending | pending | Labels may simplify at far zoom, but names must be recoverable by normal zoom/select behavior. |
| Label overlap violations | pending | pending | Count obvious incoherent overlap with dense nodes, edges, panels, controls, or other labels. |
| Graph guidance fallback required | pending | pending | Record only if a short toggle/explanation is added; guidance cannot substitute for graph labels. |

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
| `go test .\internal\mcp .\internal\cli .\internal\aicontext -count=1` | pending | pending | Minimum focused test scope for query, query-health, AI context, setup/package surfaces. |
| `avmatrix analyze --force` | pending | pending | Must not use `--skip-agents-md`. |
| `go run .\cmd\avmatrix query --help` | pending | pending | Must expose normal query usage and any lane/explain surface. |
| `go run .\cmd\avmatrix query lanes` or recorded equivalent | pending | pending | Must let users/agents discover query capability lanes. |
| `go run .\cmd\avmatrix query explain "<intent>" --repo AVmatrix --json` or recorded equivalent | pending | pending | Must expose lane/rank/match evidence without reading code. |
| `go run .\cmd\avmatrix query "<intent>" --repo AVmatrix` | pending | pending | Existing broad query behavior must remain compatible. |
| `go run .\cmd\avmatrix query-health --suite docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json --repo AVmatrix --json` | pending | pending | Dedicated suite for this plan. |
| `cd avmatrix-web; npm run test` | pending | pending | Required after graph labeling work; must include focused label/layout tests. |
| `cd avmatrix-web; npm run test:e2e` or focused Playwright expanded to full e2e before closure | pending | pending | Must prove ring and island labels are visible/readable in browser. |
| Browser screenshot validation for graph labels | pending | pending | Capture desktop and smaller viewport screenshots and record artifact paths. |
| Generated skill inventory check | pending | pending | Count and content fragments. |
| Setup/package smoke commands | pending | pending | Required if touched. |
| MCP setup/resource guidance smoke | pending | pending | Required if MCP resource guidance is touched. |
| MCP `query` smoke through local tool wrapper or focused test | pending | pending | Must verify machine-readable lane/rank/match evidence for agents. |
| `detect-changes` | pending | pending | Scope should match query, Web graph labels, AI context, skills, setup/package, and docs changes. |
