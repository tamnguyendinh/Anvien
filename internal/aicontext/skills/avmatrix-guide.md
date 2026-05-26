---
name: avmatrix-guide
description: "Use when the user asks about AVmatrix itself, available tools, MCP resources, graph schema, or workflow reference. Examples: \"What AVmatrix tools are available?\", \"How do I use AVmatrix?\""
---

# AVmatrix Tool And Resource Guide

Use this skill as the reference for the local AVmatrix MCP and CLI surfaces.

## MCP Tools

| Tool | Use |
|---|---|
| `list_repos` | List indexed repositories |
| `query` | Process-grouped search for a concept, with App Layer, Functional Area, and ResolutionGap summaries when available |
| `context` | Symbol callers, callees, refs, processes, source-site proof/status, and related ResolutionGap rows |
| `impact` | Blast radius analysis before edits, with affected App Layers, Functional Areas, and resolution-health risks |
| `detect_changes` | Git-diff impact before commit, with changed/affected semantic layers and ResolutionGap impact |
| `rename` | Graph-guided symbol rename |
| `cypher` | Read-only graph query |
| `route_map` | Route surface inspection with semantic route, consumer, and flow fields |
| `tool_map` | Tool surface inspection |
| `shape_check` | Contract and shape checks with semantic route/consumer fields |
| `api_impact` | API route impact analysis with consumer/flow layer summaries and resolution-health impact |

## Resources

| Resource | Use |
|---|---|
| `avmatrix://repos` | All indexed repos |
| `avmatrix://setup` | Setup and surface reference |
| `avmatrix://repo/{name}/context` | Repo overview and freshness |
| `avmatrix://repo/{name}/clusters` | Functional areas |
| `avmatrix://repo/{name}/cluster/{clusterName}` | Functional area detail |
| `avmatrix://repo/{name}/processes` | Execution flow list |
| `avmatrix://repo/{name}/process/{processName}` | Step-by-step trace |
| `avmatrix://repo/{name}/schema` | Graph schema for Cypher |

## MCP Prompts

| Prompt | Use |
|---|---|
| `detect_impact` | Agent template for pre-commit impact analysis using `detect_changes`, `context`, and `impact`, with CLI fallback guidance. HIGH/CRITICAL are blast-radius warnings, not edit bans. |
| `generate_map` | Agent template for architecture documentation from graph evidence. It resolves the repo through `avmatrix://repos` when needed, checks freshness, reads context/clusters/processes/process details, and restricts claims and Mermaid edges to resources/tools/commands actually read. |

Prompts are MCP prompt templates, not CLI commands. They automate a workflow for the agent; they do not replace repository rules, impact checks, or change detection.

## Standard Workflows

### Explore

Use `query` for concepts, then `context` on important symbols, then process resources for traces.

### Edit

Run `impact({target: "<symbol>", direction: "upstream"})` before editing functions, classes, or methods. Warn on HIGH or CRITICAL risk and review direct callers first.

### Commit

Run `detect_changes({scope: "all"})` before committing. Confirm changed symbols and affected flows match the intended scope.

### Rename

Use `rename` instead of text replacement. Start with a dry run, inspect confidence, apply only when the edits make sense, then run `detect_changes`.

## CLI Fallback

When MCP tools are unavailable, use the CLI equivalent with `--repo <name>`.

| Need | CLI |
|---|---|
| Concept search | `avmatrix query "<concept>" --repo <name>` |
| Symbol context | `avmatrix context "<symbol>" --repo <name>` |
| Blast radius | `avmatrix impact "<symbol>" --repo <name> --direction upstream` |
| Changed-scope review | `avmatrix detect-changes --repo <name> --scope all` |
| Graph query | `avmatrix cypher "<query>" --repo <name>` |
| Symbol rename | `avmatrix rename <symbol> <newName> --repo <name>` |
| Route map | `avmatrix api route-map [route] --repo <name>` |
| Tool map | `avmatrix api tool-map [tool] --repo <name>` |
| Shape check | `avmatrix api shape-check [route] --repo <name>` |
| API impact | `avmatrix api impact [route] --repo <name>` |

For graph quality checks, use these CLI-only diagnostic commands:

| Command | Use |
|---|---|
| `avmatrix graph-health` | Audit topology health, diagnostics, component membership, confidence, resolution-health overlays, and prioritized graph-health candidates from the indexed repo graph. |
| `avmatrix query-health` | Benchmark query retrieval against expected files and symbols, with separate threshold pass and exact target-coverage pass results. |
| `avmatrix resolution-inventory` | Report persisted ResolutionGap and Resolution Health counts, including classification/actionability and non-actionable builtin/standard-library/test-framework breakdown. |
| `avmatrix source-site-accuracy` | Report source-site inventory and proof-based resolved-edge accuracy metrics, with optional golden fixture validation. |

Use these commands for different graph-quality questions: `graph-health` answers topology/component/diagnostic triage, `query-health` answers retrieval quality, `resolution-inventory` answers persisted gap inventory, and `source-site-accuracy` answers proof/source-site accuracy.

## Semantic Graph Fields

Fresh `analyze` output can include:

- `appLayer` and `functionalArea` on graph nodes;
- source-site IDs, proof kind, target role, target text, and range metadata on proven relationships;
- persisted `ResolutionGap` entities for unresolved or diagnostic references;
- `resolutionConfidence`, `resolutionGapCount`, and Resolution Health buckets for graph readers.

ResolutionGap entities are diagnostic facts, not proven in-repo symbols. Do not treat them as resolved topology unless the graph also provides a proven relationship.

## Stale Index Policy

If any AVmatrix surface reports stale data, refresh with `avmatrix analyze --force` before relying on graph facts.
