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

Status: baseline recorded; final pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Embedded source skill Markdown files | files | 6 | 6 | 0 | 11 |
| Registered base skills | skills | 6 | 6 | 0 | 11 |
| Generated `.claude/skills/avmatrix/**/SKILL.md` files | files | 6 | 6 | 0 | 11 |
| Root generated Skills table rows | rows | 6 | 6 | 0 | 11 |
| Source skill total size | bytes | 17,499 | 17,499 | 0 | record |
| Generated skill total size | bytes | 17,499 | 17,499 | 0 | equal source total |
| Source/generated matching hashes | pairs | 6 | 6 | 0 | final generated skill count |
| Source skills with `name` frontmatter | files | 6 | 6 | 0 | final source skill count |
| Source skills with `description` frontmatter | files | 6 | 6 | 0 | final source skill count |

## B1 - Command Surface Inventory

Status: baseline recorded; Phase 1 query-lane metrics recorded; final pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Visible built CLI top-level commands | commands | 24 | 24 | 0 | >= 24 |
| Visible source CLI top-level commands | commands | 24 | 24 | 0 | >= 24 |
| Hidden lifecycle command families from source | families | 2 | 2 | 0 | record |
| MCP tools from source | tools | 15 | 15 | 0 | record |
| MCP fixed resources from source | resources | 2 | 2 | 0 | record |
| MCP resource templates from source | templates | 6 | 6 | 0 | record |
| MCP prompts from source | prompts | 2 | 2 | 0 | 2 |
| Discoverable query capability lanes | lanes | 0 | 8 | +8 | >= 8 |
| Query explain metadata fields validated | fields | 0 | 6 | +6 | >= 6 |
| Query-health cases with lane evidence fields | cases | 0 | 8 | +8 | 8 |
| Normal query compatibility failures | failures | 0 | 0 | 0 | 0 |
| First-class graph-health CLI command families | families | 0 | 0 | 0 | >= 1 |
| MCP prompt templates reviewed | prompts | 0 | 0 | 0 | 2 |
| Accepted CLI parity gaps implemented | gaps | 0 | pending | pending | pending |

## B2 - Query Reliability Metrics

Status: Phase 1 recorded.

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
