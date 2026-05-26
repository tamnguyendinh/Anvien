# AVmatrix Skill System Upgrade Benchmark Ledger

Date: 2026-05-23

Status: active

Companion files:

- Plan: [2026-05-23-avmatrix-skill-system-upgrade-plan.md](2026-05-23-avmatrix-skill-system-upgrade-plan.md)
- Evidence: [2026-05-23-avmatrix-skill-system-upgrade-evidence.md](2026-05-23-avmatrix-skill-system-upgrade-evidence.md)

## Benchmark Rules

This file records only quantitative benchmark data: inventory counts, before/after counts, rates, sizes, latency, throughput, graph counts, hit/miss counts, coverage percentages, and measured UI geometry.

Command output, validation logs, impact details, screenshots, and implementation notes belong in the evidence ledger. A benchmark row must be measurable and must use baseline/final or before/after values wherever a comparison exists.

Use `pending` only when a future phase has not measured that value yet.

## B0 - Skill Source And Generated Inventory

Status: baseline recorded; P1.5 guidance size update recorded; final pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Embedded source skill Markdown files | files | 6 | 6 | 0 | 11 |
| Registered base skills | skills | 6 | 6 | 0 | 11 |
| Generated `.claude/skills/avmatrix/**/SKILL.md` files | files | 6 | 6 | 0 | 11 |
| Root generated Skills table rows | rows | 6 | 6 | 0 | 11 |
| Source skill total size | bytes | 17,499 | 18,843 | +1,344 | record |
| Generated skill total size | bytes | 17,499 | 18,843 | +1,344 | equal source total |
| Source/generated matching hashes | pairs | 6 | 6 | 0 | final generated skill count |
| Source skills with `name` frontmatter | files | 6 | 6 | 0 | final source skill count |
| Source skills with `description` frontmatter | files | 6 | 6 | 0 | final source skill count |

## B1 - Command Surface Inventory

Status: baseline recorded; Phase 1 query-lane metrics and P1.5 graph-health command metrics recorded; final pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Visible built CLI top-level commands, including Cobra built-ins | commands | 24 | 28 | +4 | >= 28 |
| Explicit visible source CLI top-level commands | commands | 24 | 26 | +2 | >= 26 |
| Hidden lifecycle command families from source | families | 2 | 2 | 0 | record |
| MCP tools from source | tools | 15 | 16 | +1 | record |
| MCP fixed resources from source | resources | 2 | 2 | 0 | record |
| MCP resource templates from source | templates | 6 | 6 | 0 | record |
| MCP prompts from source | prompts | 2 | 2 | 0 | 2 |
| Discoverable query capability lanes | lanes | 0 | 8 | +8 | >= 8 |
| Query explain metadata fields validated | fields | 0 | 6 | +6 | >= 6 |
| Query-health cases with lane evidence fields | cases | 0 | 9 | +9 | >= 9 |
| Normal query compatibility failures | failures | 0 | 0 | 0 | 0 |
| First-class graph-health CLI command families | families | 0 | 1 | +1 | >= 1 |
| MCP prompt templates reviewed | prompts | 0 | 2 | +2 | 2 |
| Accepted CLI parity gaps implemented | gaps | 0 | 5 | +5 | 5 |

P1.5 graph-health CLI smoke inventory:

| Metric | Unit | Result |
|---|---:|---:|
| Analyze scanned files before smoke | files | 769 |
| Analyze parsed files before smoke | files | 572 |
| Analyze graph nodes before smoke | nodes | 87,382 |
| Analyze graph relationships before smoke | relationships | 120,022 |
| Graph-health counted relationships | relationships | 25,951 |
| Graph-health components | components | 79,089 |
| Graph-health detached components | components | 62 |
| Graph-health root nodes | nodes | 852 |
| Graph-health unresolved references | references | 63,576 |
| Graph-health report total candidates | candidates | 47,316 |
| Graph-health report returned candidates at `--limit 20` | candidates | 20 |
| Graph-health component list total components | components | 79,089 |
| Graph-health component list returned at `--limit 20` | components | 20 |
| First report candidate counted incoming | relationships | 0 |
| First report candidate counted outgoing | relationships | 6 |
| First explained component sample nodes | nodes | 20 |

P1.6 CLI parity inventory and smoke:

| Metric | Unit | Result |
|---|---:|---:|
| Built CLI visible top-level commands, including Cobra built-ins | commands | 28 |
| Built CLI new P1.6 top-level commands | commands | 2 |
| Explicit source visible top-level commands | commands | 26 |
| Explicit source hidden lifecycle families | families | 2 |
| Explicit source hidden lifecycle subcommands | commands | 5 |
| MCP tools inventoried from source | tools | 16 |
| MCP fixed resources inventoried from source | resources | 2 |
| MCP resource templates inventoried from source | templates | 6 |
| MCP prompt templates inventoried from source | prompts | 2 |
| HTTP/Web endpoint handlers inventoried | endpoints | 24 |
| P1.6 accepted parity commands implemented | commands | 5 |
| P1.6 CLI parity focused test package failures | failures | 0 |
| P1.6 fixture analyze scanned files | files | 3 |
| P1.6 fixture analyze parsed files | files | 2 |
| P1.6 fixture analyze graph nodes | nodes | 22 |
| P1.6 fixture analyze graph relationships | relationships | 20 |
| P1.6 fixture `api route-map` total routes | routes | 1 |
| P1.6 fixture `api route-map` consumers for `/api/users` | consumers | 1 |
| P1.6 fixture `api tool-map` total tools | tools | 1 |
| P1.6 fixture `api shape-check` routes checked | routes | 1 |
| P1.6 fixture `api shape-check` mismatches | mismatches | 1 |
| P1.6 fixture `api impact` direct consumers | consumers | 1 |
| P1.6 fixture `api impact` affected flows | flows | 0 |
| P1.6 fixture `rename` dry-run files affected | files | 1 |
| P1.6 fixture `rename` dry-run total edits | edits | 1 |
| P1.6 AVmatrix `rename NewRootCommand` dry-run files affected | files | 4 |
| P1.6 AVmatrix `rename NewRootCommand` dry-run total edits | edits | 4 |
| AVmatrix self graph Route nodes available for API smoke | nodes | 0 |
| AVmatrix self graph Tool nodes available for API smoke | nodes | 0 |

P1.7 MCP prompt inventory and smoke:

| Metric | Unit | Result |
|---|---:|---:|
| MCP prompt templates reviewed | prompts | 2 |
| MCP prompts returned by runtime `prompts/list` smoke | prompts | 2 |
| MCP prompt argument schemas validated | arguments | 3 |
| Runtime MCP JSON-RPC smoke requests | requests | 4 |
| Runtime MCP JSON-RPC smoke responses | responses | 4 |
| Runtime MCP smoke process exit code | code | 0 |
| `generate_map` actionable `{name}` placeholders with repo argument | placeholders | 0 |
| `generate_map` actionable `{name}` placeholders without repo argument | placeholders | 0 |
| `generate_map` concrete repo resource URIs validated | URIs | 3 |
| `generate_map` repo-discovery resource references validated | references | 1 |
| `generate_map` freshness command references validated | references | 1 |
| `generate_map` deterministic selection criteria groups documented | groups | 5 |
| `detect_impact` change-detection surfaces validated | surfaces | 2 |
| Focused P1.7 test package failures | failures | 0 |
| P1.7 final analyze scanned files | files | 772 |
| P1.7 final analyze parsed files | files | 575 |
| P1.7 final analyze graph nodes | nodes | 87,883 |
| P1.7 final analyze graph relationships | relationships | 120,603 |
| P1.7 detect-changes changed files | files | 11 |
| P1.7 detect-changes changed symbols | symbols | 49 |
| P1.7 detect-changes affected processes | processes | 9 |

## B2 - Query Reliability Metrics

Status: Phase 1 and P1-L recorded.

Suite-level before/after:

| Metric | Unit | Baseline | After Phase 1 | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Suite cases | cases | 8 | 8 | 0 | 8 |
| Expected targets | targets | 54 | 54 | 0 | 54 |
| Matched targets | targets | 33 | 46 | +13 | increase |
| Matched target rate | percent | 61.11 | 85.19 | +24.08 pp | increase |
| Missed targets | targets | 21 | 8 | -13 | decrease |
| Missed target rate | percent | 38.89 | 14.81 | -24.08 pp | decrease |
| Threshold passed | cases | 5 | 8 | +3 | 8 |
| Threshold pass rate | percent | 62.50 | 100.00 | +37.50 pp | 100.00 |
| Threshold failed | cases | 3 | 0 | -3 | 0 |
| Exact passed | cases | 1 | 3 | +2 | record |
| Exact pass rate | percent | 12.50 | 37.50 | +25.00 pp | record |
| Exact failed | cases | 7 | 5 | -2 | record |

Per-case before/after:

| Case | Expected | Threshold baseline | Threshold after | Exact baseline | Exact after | Hit@5 baseline | Hit@5 after | Hit@10 baseline | Hit@10 after | Matched baseline | Matched after | Missed baseline | Missed after |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| `ai-context-generated-skills-owner-discovery` | 6 | 0 | 1 | 0 | 0 | 0 | 5 | 1 | 5 | 1 | 5 | 5 | 1 |
| `setup-editor-skill-installation-owner-discovery` | 4 | 0 | 1 | 0 | 1 | 3 | 4 | 3 | 4 | 3 | 4 | 1 | 0 |
| `package-skill-distribution-owner-discovery` | 6 | 1 | 1 | 1 | 0 | 5 | 2 | 6 | 4 | 6 | 4 | 0 | 2 |
| `mcp-setup-resource-prompt-guidance-owner-discovery` | 7 | 1 | 1 | 0 | 0 | 3 | 3 | 5 | 6 | 5 | 6 | 2 | 1 |
| `query-command-surface-owner-discovery` | 8 | 1 | 1 | 0 | 0 | 6 | 4 | 6 | 6 | 6 | 6 | 2 | 2 |
| `graph-quality-command-surface-owner-discovery` | 8 | 0 | 1 | 0 | 0 | 2 | 6 | 2 | 6 | 2 | 6 | 6 | 2 |
| `api-surface-tool-discovery` | 7 | 1 | 1 | 0 | 1 | 5 | 6 | 6 | 7 | 6 | 7 | 1 | 0 |
| `cross-repo-command-surface-discovery` | 8 | 1 | 1 | 0 | 1 | 4 | 8 | 4 | 8 | 4 | 8 | 4 | 0 |

P1-L expanded-suite regression:

| Metric | Unit | After P1-L | Target |
|---|---:|---:|---:|
| Suite cases | cases | 9 | >= 9 |
| Expected targets | targets | 61 | >= 61 |
| Expected process targets | targets | 2 | >= 2 |
| Matched targets | targets | 53 | increase |
| Missed targets | targets | 8 | record |
| Threshold passed | cases | 9 | 9 |
| Threshold failed | cases | 0 | 0 |
| Exact passed | cases | 4 | record |
| Exact failed | cases | 5 | record |

P1-L process case:

| Case | Expected | Threshold | Exact | Hit@5 | Hit@10 | Matched | Missed | Matched process labels |
|---|---:|---:|---:|---:|---:|---:|---:|---|
| `cross-repo-execution-flow-process-discovery` | 7 | 1 | 1 | 7 | 7 | 7 | 0 | `Sync -> NormalizeHTTPPath`, `Query -> GroupProcess` |

## B3 - Canonical Executable And Package Artifact Metrics

Status: P4-H recorded.

| Metric | Unit | Baseline | After P4-H | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Canonical CLI size | bytes | 50,344,448 | 50,395,136 | +50,688 | record |
| Launcher executable size | bytes | 6,993,408 | 6,993,408 | 0 | record |
| Launcher support wrapper size | bytes | 2,053,632 | 2,053,632 | 0 | record |
| Production AVmatrix CLI executable authorities | count | 2 | 1 | -1 | 1 |
| Independent `server-bundle/avmatrix.exe` executables | count | 1 | 0 | -1 | 0 |
| Launcher support wrapper executables | count | 1 | 1 | 0 | 1 |
| Native DLL content mismatches after build | count | pending | 0 | pending | 0 |

## B4 - Graph Labeling And Visual Layout Metrics

Status: baseline problem identified; measured browser baseline and final pending.

| Metric | Unit | Baseline | Final | Delta | Target |
|---|---:|---:|---:|---:|---:|
| On-canvas macro ring labels | labels | 0 | pending | pending | visible macro ring count |
| On-canvas island labels | labels | 0 | pending | pending | visible major island count |
| Macro ring label coverage | percent | 0.00 | pending | pending | 100.00 |
| Island label coverage | percent | 0.00 | pending | pending | 100.00 |
| Label overlap violations | violations | pending | pending | pending | 0 |
| Runtime-visible macro rings | rings | pending | pending | pending | record |
| Runtime-visible islands | islands | pending | pending | pending | record |
| Minimum node spacing inside islands | px | pending | pending | pending | >= 2x rendered node diameter |
| Island-to-island gutter distance | px | pending | pending | pending | >= 1x largest visible island radius |
| Macro ring gutter distance | px | pending | pending | pending | > 0 |
| Rail-like island shape violations | islands | pending | pending | pending | 0 |
| Auto optimizer runs before manual click | runs | 0 | pending | pending | 0 |
| Manual optimizer runs after one click | runs | pending | pending | pending | 1 |

## B5 - Setup And Package Skill Distribution Metrics

Status: final pending.

| Metric | Unit | Baseline | Final | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Editor skill install targets checked | targets | 0 | pending | pending | supported target count |
| Skills installed per checked editor target | skills | pending | pending | pending | 11 |
| Package-root skill source files | files | 0 | pending | pending | 0 or 11, depending on selected package design |
| Packaged embedded skill files | files | pending | pending | pending | 11 |
| Setup/generated skill inventory mismatches | mismatches | pending | pending | pending | 0 |
| Package/generated skill inventory mismatches | mismatches | pending | pending | pending | 0 |
| Setup/package skill content hash mismatches | mismatches | pending | pending | pending | 0 |

## B6 - Generated AI Context Final Metrics

Status: final pending.

| Metric | Unit | Baseline | Final | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Generated root managed blocks | files | 2 | pending | pending | 2 |
| Generated root stats lines | lines | 2 | pending | pending | 2 |
| Generated root command selection table rows | rows | pending | pending | pending | record |
| Generated root Skills table rows | rows | 6 | pending | pending | 11 |
| Generated skill files | files | 6 | pending | pending | 11 |
| Generated skill total size | bytes | 17,499 | pending | pending | record |
| Generated skill/source hash matches | pairs | 6 | pending | pending | 11 |
| Generated files edited directly | files | 0 | pending | pending | 0 |
