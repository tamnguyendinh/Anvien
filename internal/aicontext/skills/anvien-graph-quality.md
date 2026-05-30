---
name: anvien-graph-quality
description: "Use when the user needs graph-health, query-health, resolution inventory, source-site accuracy, or benchmark comparison evidence."
---

# Graph Quality With Anvien

Use this skill when the question is whether the graph itself is healthy, complete, fresh, explainable, or reliable enough for code navigation and impact work.

## Command Choices

| Need | Use |
|---|---|
| Topology and diagnostics | `anvien graph-health summary --repo <repo> --json` |
| Prioritized health candidates | `anvien graph-health report --repo <repo> --limit <n> --json` |
| File-level hotspots | `anvien graph-health files --repo <repo> --json` or `anvien file-hotspots --repo <repo> --sort unresolved --json` |
| Component inventory | `anvien graph-health components --repo <repo> --json` |
| Node/component explanation | `anvien graph-health explain <node-or-name> --repo <repo> --json` |
| Query retrieval benchmark | `anvien query-health --repo <repo> --suite <file>` |
| ResolutionGap inventory | `anvien resolution-inventory --graph .anvien/graph.json` |
| Source-site proof accuracy | `anvien source-site-accuracy --graph .anvien/graph.json` |
| Analyze benchmark comparison | `anvien benchmark-compare <before> <after>` |

## Workflow

1. Run `anvien analyze --force` first when graph freshness matters.
2. Choose the quality command by failure type: topology, file hotspots, query retrieval, unresolved references, source-site proof, or performance/capacity.
3. Use file-layer output to locate the concrete file, then open it with `context file` to inspect symbol tree, derived relationships, unresolved source sites, and linked flows/tests.
4. Keep threshold and exact query-health results separate. Threshold pass means usable navigation; exact pass means all expected targets were found.
5. Preserve counts, samples, missed targets, file groups, and noise reasons in evidence. Do not compress away traceability.
6. If quality output affects an implementation decision, run impact before editing source.

## Query Reliability Guidance

Broad `query` is a candidate discovery command with multiple lanes. It is not a proof that the first ranked result is the owner. For broad intent cases, verify candidates with `context`, exact file inspection, and query-health evidence. When an expected target is missed, record the miss even if threshold pass succeeds.

For the Anvien skill-system upgrade suite, use the dedicated query-health suite under `docs/query-health/2026-05-23-anvien-skill-system-upgrade-suite.json` when validating AI-context, command-surface, graph-quality, API-surface, and cross-repo discovery.

## Evidence To Record

- Fresh analyze counts before quality commands.
- Query-health case id, expected targets, matched targets, missed targets, hit@5, hit@10, threshold pass, exact pass, and noise reason.
- Graph-health component/candidate counts, file hotspot counts, and representative explanations.
- ResolutionGap classification/actionability counts, file groups, nearest source symbols, and top target samples.
- Source-site inventory counts, false-resolved candidates, missing proof counts, file groups, trace samples, and golden fixture results when used.

## Current Limitations

Graph-quality commands report evidence; they do not decide the product fix. A ResolutionGap can be useful diagnostic evidence but is not a resolved in-repo symbol. If quality output is stale or inconsistent, refresh and rerun before making a code decision.
