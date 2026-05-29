# Anvien Skill System Upgrade Benchmark Ledger

Date: 2026-05-23

Status: active

Companion files:

- Plan: [2026-05-23-anvien-skill-system-upgrade-plan.md](2026-05-23-anvien-skill-system-upgrade-plan.md)
- Evidence: [2026-05-23-anvien-skill-system-upgrade-evidence.md](2026-05-23-anvien-skill-system-upgrade-evidence.md)

## Benchmark Rules

This file records only quantitative benchmark data: inventory counts, before/after counts, rates, sizes, latency, throughput, graph counts, hit/miss counts, coverage percentages, and measured UI geometry.

Command output, validation logs, impact details, screenshots, and implementation notes belong in the evidence ledger. A benchmark row must be measurable and must use baseline/final or before/after values wherever a comparison exists.

Use `pending` only when a future phase has not measured that value yet.

## B0 - Skill Source And Generated Inventory

Status: baseline recorded; P1.5 guidance size update recorded; Phase 3 source/registry update recorded; final regeneration pending.

| Metric | Unit | Baseline | Latest | Delta | Final target |
|---|---:|---:|---:|---:|---:|
| Embedded source skill Markdown files | files | 6 | 11 | +5 | 11 |
| Registered base skills | skills | 6 | 11 | +5 | 11 |
| Generated `.claude/skills/anvien/**/SKILL.md` files | files | 6 | 11 | +5 | 11 |
| Root generated Skills table rows | rows | 6 | 11 | +5 | 11 |
| Source skill total size | bytes | 17,499 | 30,963 | +13,464 | record |
| Generated skill total size | bytes | 17,499 | 30,963 | +13,464 | equal source total |
| Source/generated matching hashes | pairs | 6 | 11 | +5 | final generated skill count |
| Source skills with `name` frontmatter | files | 6 | 11 | +5 | final source skill count |
| Source skills with `description` frontmatter | files | 6 | 11 | +5 | final source skill count |

P3 embedded source skill inventory:

| Skill source file | Bytes |
|---|---:|
| `anvien-ai-context.md` | 2,883 |
| `anvien-api-surface.md` | 2,278 |
| `anvien-cli.md` | 4,259 |
| `anvien-cross-repo.md` | 2,072 |
| `anvien-debugging.md` | 2,454 |
| `anvien-exploring.md` | 3,045 |
| `anvien-graph-quality.md` | 3,086 |
| `anvien-guide.md` | 3,842 |
| `anvien-impact-analysis.md` | 2,657 |
| `anvien-refactoring.md` | 2,170 |
| `anvien-runtime-packaging.md` | 2,217 |

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
| P1.6 Anvien `rename NewRootCommand` dry-run files affected | files | 4 |
| P1.6 Anvien `rename NewRootCommand` dry-run total edits | edits | 4 |
| Anvien self graph Route nodes available for API smoke | nodes | 0 |
| Anvien self graph Tool nodes available for API smoke | nodes | 0 |

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
| Production Anvien CLI executable authorities | count | 2 | 1 | -1 | 1 |
| Independent `server-bundle/anvien.exe` executables | count | 1 | 0 | -1 | 0 |
| Launcher support wrapper executables | count | 1 | 1 | 0 | 1 |
| Native DLL content mismatches after build | count | pending | 0 | pending | 0 |

## B4 - Graph Labeling And Visual Layout Metrics

Status: baseline problem identified; measured browser baseline and final pending.

P2-A screenshot evidence asset:

| Metric | Unit | Baseline |
|---|---:|---:|
| Problem screenshot width | px | 1,314 |
| Problem screenshot height | px | 826 |
| Problem screenshot file size | bytes | 341,738 |

P2-B trace inventory:

| Metric | Unit | Result |
|---|---:|---:|
| Fresh graph nodes before Web trace | nodes | 87,885 |
| Fresh graph relationships before Web trace | relationships | 120,605 |
| Top query owner definitions recorded | definitions | 10 |
| Source owner files traced | files | 7 |
| Source owner symbols traced | symbols | 8 |
| `GraphCanvas` direct upstream dependents | dependents | 1 |
| `GraphCanvas` total upstream impacted nodes | nodes | 2 |

P2-C through P2-H implementation and runtime metrics:

| Metric | Unit | Result |
|---|---:|---:|
| Orientation label contract fields | fields | 11 |
| Orientation viewport label fields | fields | 15 |
| E2E orientation fixture nodes | nodes | 36 |
| E2E orientation fixture relationships | relationships | 35 |
| Desktop viewport width | px | 1,280 |
| Desktop viewport height | px | 800 |
| Desktop e2e screenshot size | bytes | 237,028 |
| Desktop visible ring labels | labels | 3 |
| Desktop visible island labels | labels | 5 |
| Desktop label overlap violations | violations | 0 |
| Smaller viewport width | px | 480 |
| Smaller viewport height | px | 720 |
| Smaller e2e screenshot size | bytes | 92,005 |
| Smaller visible ring labels before filter | labels | 3 |
| Smaller visible island labels before filter | labels | 2 |
| Smaller label overlap violations before filter | violations | 0 |
| Smaller visible ring labels after Method filter | labels | 3 |
| Smaller visible island labels after Method filter | labels | 2 |
| Smaller Method island labels after Method filter | labels | 0 |
| Smaller label overlap violations after Method filter | violations | 0 |
| Narrow-viewport clamped ring overlap violations in unit test | violations | 0 |
| Existing geometry unit tests validating P2-F2 layout behavior | tests | 22 |
| Small island gutter minimum validated | px | 100 |
| Imbalanced island gutter minimum validated | px | 500 |
| Dense large-repo island gutter minimum validated | px | 900 |
| App Layer ring gap minimum validated | px | 200 |
| Rail-like island aspect-ratio violations validated | violations | 0 |

| Metric | Unit | Baseline | Final | Delta | Target |
|---|---:|---:|---:|---:|---:|
| On-canvas macro ring labels | labels | 0 | 3 | +3 | visible macro ring count |
| On-canvas island labels | labels | 0 | 5 | +5 | visible major island count |
| Macro ring label coverage | percent | 0.00 | 100.00 | +100.00 pp | 100.00 |
| Island label coverage | percent | 0.00 | 100.00 | +100.00 pp | 100.00 |
| Label overlap violations | violations | pending | 0 | pending | 0 |
| Runtime-visible macro rings | rings | pending | 3 | pending | record |
| Runtime-visible islands | islands | pending | 5 | pending | record |
| Minimum node spacing inside islands | px | pending | 30 | pending | >= 2x rendered node diameter |
| Island-to-island gutter distance | px | pending | 900 | pending | >= 1x largest visible island radius |
| Macro ring gutter distance | px | pending | 200 | pending | > 0 |
| Rail-like island shape violations | islands | pending | 0 | pending | 0 |
| Auto optimizer runs before manual click | runs | 0 | 0 | 0 | 0 |
| Manual optimizer runs after one click | runs | pending | pending | pending | 1 |

## B5 - Setup And Package Skill Distribution Metrics

Status: P5 setup/package validation recorded after full build and regeneration.

| Metric | Unit | Baseline | Final | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Editor skill install targets checked | targets | 0 | 4 | +4 | 4 supported setup targets |
| Skills installed per checked editor target | skills | 7 | 11 | +4 | 11 |
| Installed embedded skill bytes per checked editor target | bytes | pending | 30,984 | pending | record |
| Package-root skill source files | files | 7 | 0 | -7 | 0 |
| Retired package-root `anvien-pr-review` installed | files | 1 | 0 | -1 | 0 |
| Package fallback copied files | files | pending | 293 | pending | record |
| Package fallback inventory files | files | pending | 294 | pending | record |
| Package fallback copied Go files | files | pending | 278 | pending | record |
| Packaged embedded skill Markdown files | files | pending | 11 | pending | 11 |
| Packaged embedded skill Markdown bytes | bytes | pending | 30,984 | pending | record |
| Package fallback dry-run tarball files | files | pending | 301 | pending | record |
| Package dry-run package-root skill files | files | 7 | 0 | -7 | 0 |
| Setup/generated skill inventory mismatches | mismatches | pending | 0 | pending | 0 |
| Package/generated skill inventory mismatches | mismatches | pending | 0 | pending | 0 |
| Package source/prepared skill SHA-256 mismatches | mismatches | pending | 0 | pending | 0 |
| Setup/package skill content hash mismatches | mismatches | pending | 0 | pending | 0 |

## B6 - Generated AI Context Final Metrics

Status: final inventory captured after P5-E regeneration.

| Metric | Unit | Baseline | Final | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Generated root managed blocks | files | 2 | 2 | 0 | 2 |
| Generated root stats lines | lines | 2 | 2 | 0 | 2 |
| Generated root command selection table rows | rows/root | pending | 28 | pending | record |
| Generated root Skills table rows | rows/root | 6 | 11 | +5 | 11 |
| Generated skill files | files | 6 | 11 | +5 | 11 |
| Generated skill total size | bytes | 17,499 | 30,984 | +13,485 | record |
| Generated skill/source hash matches | pairs | 6 | 11 | +5 | 11 |
| Generated skill/source hash mismatches | mismatches | pending | 0 | pending | 0 |
| Generated skill total lines | lines | pending | 556 | pending | record |
| Generated files edited directly | files | 0 | 0 | 0 | 0 |

## B7 - Phase 5 Command Output Inventory Metrics

Status: complete after Phase 5 command smoke validation.

| Metric | Unit | Baseline | Latest | Delta | Target |
|---|---:|---:|---:|---:|---:|
| MCP `anvien://setup` repo sections | sections | pending | 4 | pending | record |
| MCP `anvien://setup` code-table rows | rows | pending | 92 | pending | record |
| MCP `anvien://setup` CLI equivalent rows | rows | pending | 20 | pending | record |
| MCP `anvien://setup` prompt rows | rows | pending | 8 | pending | record |
| MCP `anvien://setup` AI Context sections | sections | pending | 4 | pending | >=1 |
| MCP `anvien://setup` repo resource links | links | pending | 16 | pending | record |
| Query-health cases | cases | pending | 9 | pending | record |
| Query-health threshold passed | cases | pending | 9 | pending | 9 |
| Query-health threshold failed | cases | pending | 0 | pending | 0 |
| Query-health exact passed | cases | pending | 4 | pending | record |
| Query-health exact failed | cases | pending | 5 | pending | record |
| Query-health expected targets | targets | pending | 63 | pending | record |
| Query-health matched targets | targets | pending | 54 | pending | record |
| Query-health missed targets | targets | pending | 9 | pending | record |
| Query-health unique matched lanes | lanes | pending | 8 | pending | record |
| Query-health top-result source categories | categories | pending | 3 | pending | record |
| Query-health matched target global rank max | rank | pending | 41 | pending | record |
| Query-health matched target source rank max | rank | pending | 9 | pending | record |
| Broad-discovery regression definitions | results | pending | 10 | pending | record |
| Broad-discovery regression process candidates | results | pending | 10 | pending | >0 |
| Broad-discovery definitions with match reasons | results | pending | 10 | pending | 10 |
| Broad-discovery definitions with execution-flow lane | results | pending | 10 | pending | >0 |
| Broad-discovery distinct definition lanes | lanes | pending | 7 | pending | record |
| Graph-health summary nodes | nodes | pending | 88,583 | pending | record |
| Graph-health summary relationships | relationships | pending | 121,474 | pending | record |
| Graph-health component count | components | pending | 80,216 | pending | record |
| Graph-health detached component count | components | pending | 64 | pending | record |
| Graph-health report candidates with `--limit 5` | candidates | pending | 5 | pending | 5 |
| Graph-health components with `--limit 5` | components | pending | 5 | pending | 5 |
| Graph-health explain missing-node exit code | exit_code | pending | 1 | pending | nonzero |
| CLI parity accepted API subcommands | commands | pending | 4 | pending | 4 |
| CLI parity fixture route-map total | routes | pending | 1 | pending | 1 |
| CLI parity fixture tool-map total | tools | pending | 1 | pending | 1 |
| CLI parity fixture shape mismatches | mismatches | pending | 1 | pending | 1 |
| CLI parity fixture API impact direct consumers | consumers | pending | 1 | pending | 1 |
| CLI parity fixture rename files affected | files | pending | 1 | pending | 1 |
| CLI parity fixture rename total edits | edits | pending | 1 | pending | 1 |
| Hidden lifecycle source inventory lines | lines | pending | 9 | pending | record |
| Root help hidden lifecycle commands visible | commands | pending | 0 | pending | 0 |
| Root help intentionally MCP/API/Web-only grep/process/cluster commands visible | commands | pending | 0 | pending | 0 |
| MCP prompt list count | prompts | pending | 2 | pending | 2 |
| MCP `generate_map` argument count | arguments | pending | 1 | pending | record |
| MCP `detect_impact` argument count | arguments | pending | 2 | pending | record |
| MCP prompt get messages checked | messages | pending | 3 | pending | 3 |
| `generate_map` no-arg required guidance checks passed | checks | pending | 8 | pending | 8 |
| `detect_impact` required guidance checks passed | checks | pending | 4 | pending | 4 |

## B8 - Phase 6 Closure Review Inventory Metrics

Status: final closure-review inventory captured after P6-B validation.

| Metric | Unit | Baseline | Latest | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Final analyze scanned files | files | pending | 774 | pending | record |
| Final analyze parsed files | files | pending | 578 | pending | record |
| Final analyze unsupported files | files | pending | 196 | pending | record |
| Final analyze failed files | files | pending | 0 | pending | 0 |
| Final graph nodes | nodes | pending | 88,584 | pending | record |
| Final graph relationships | relationships | pending | 121,475 | pending | record |
| Source embedded skill files | files | 6 | 11 | +5 | 11 |
| Generated embedded skill files | files | 6 | 11 | +5 | 11 |
| Generated/source skill name diffs | diffs | pending | 0 | pending | 0 |
| Generated/source skill SHA-256 mismatches | mismatches | pending | 0 | pending | 0 |
| Source embedded skill bytes | bytes | 17,499 | 30,984 | +13,485 | record |
| Generated embedded skill bytes | bytes | 17,499 | 30,984 | +13,485 | record |
| Source embedded skill lines | lines | pending | 556 | pending | record |
| Generated embedded skill lines | lines | pending | 556 | pending | record |
| Generated root skill rows per root file | rows/root | 6 | 11 | +5 | 11 |
| Generated root command-selection rows per root file | rows/root | pending | 44 | pending | record |
| Package dry-run tarball files after README update | files | pending | 7 | pending | record |
| Package dry-run README files | files | pending | 1 | pending | 1 |
| Package dry-run README size | bytes | pending | 18,635 | pending | record |
| Package dry-run package-root skill files | files | 7 | 0 | -7 | 0 |
| Final desktop graph-label screenshot size | bytes | pending | 237,028 | pending | record |
| Final small filtered graph-label screenshot size | bytes | pending | 92,012 | pending | record |

## B9 - Phase 7 Command-Surface Inventory Metrics

Status: P7-A inventory captured before documentation/generator edits.

| Metric | Unit | Baseline | Latest | Delta | Target |
|---|---:|---:|---:|---:|---:|
| Fresh graph nodes before P7 inventory | nodes | pending | 88,585 | pending | record |
| Fresh graph relationships before P7 inventory | relationships | pending | 121,476 | pending | record |
| Visible root CLI commands in binary help | commands | pending | 28 | pending | record |
| Visible API parity subcommands | commands | pending | 4 | pending | 4 |
| Visible graph-health subcommands | commands | pending | 4 | pending | 4 |
| Visible group subcommands | commands | pending | 7 | pending | record |
| MCP runtime tools | tools | pending | 16 | pending | 16 |
| MCP resources | resources | pending | 2 | pending | 2 |
| MCP resource templates | templates | pending | 6 | pending | 6 |
| MCP prompts | prompts | pending | 2 | pending | 2 |
| MCP setup resource text size | chars | pending | 12,950 | pending | record |
| Local HTTP/API registered endpoints | endpoints | pending | 24 | pending | record |
| Root README visible-command omissions found | commands | pending | 5 | pending | 0 after P7-B |
| Packaged README visible-command omissions found | commands | pending | 2 | pending | 0 after P7-B |
| Generated AI-context visible-command omissions found | commands | pending | 2 | pending | 0 after P7-C |
| Generated AI-context command rows after P7 regeneration | rows/root | 44 | 46 | +2 | record |
| Generated/source skill bytes after P7 regeneration | bytes | 30,984 | 31,693 | +709 | record |
| Generated/source skill lines after P7 regeneration | lines | 556 | 563 | +7 | record |
| Generated/source skill SHA-256 mismatches after rebuilt regeneration | mismatches | 2 | 0 | -2 | 0 |
| Generated root `doctor` fragments after P7 regeneration | fragments/root | 0 | 2 | +2 | 2 |
| Generated root `completion` fragments after P7 regeneration | fragments/root | 0 | 1 | +1 | 1 |
| Focused P7 Go test packages passed | packages | pending | 3 | pending | 3 |
| Root help required command omissions after P7 | commands | 5 | 0 | -5 | 0 |
| Doctor lock smoke status count | statuses | pending | 1 | pending | record |
| PowerShell completion non-empty lines | lines | pending | 222 | pending | record |
| MCP runtime tools after P7 rebuild | tools | 16 | 16 | 0 | 16 |
| MCP runtime prompts after P7 rebuild | prompts | 2 | 2 | 0 | 2 |
| Packaged README size after P7 | bytes | 18,635 | 18,788 | +153 | record |
| Packaged README package-root skill files after P7 | files | 0 | 0 | 0 | 0 |
| P7 final detect changed files | files | pending | 9 | pending | record |
| P7 final detect changed symbols/items | items | pending | 33 | pending | record |
| P7 final detect affected processes | processes | pending | 1 | pending | record |
| P7 final detect resolution gap changed entities | entities | pending | 6 | pending | record |
| P7 final detect degraded resolution-health nodes | nodes | pending | 0 | pending | 0 |
