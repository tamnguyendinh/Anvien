---
name: avmatrix-graph-quality
description: "Use when the user needs graph-health, query-health, resolution inventory, source-site accuracy, or benchmark comparison evidence."
---

# Graph Quality With AVmatrix

Use this skill when the question is whether the graph itself is healthy, complete, fresh, explainable, or reliable enough for code navigation and impact work.

## Command Choices

| Need | Use |
|---|---|
| Topology and diagnostics | `avmatrix graph-health summary --repo <repo> --json` |
| Prioritized health candidates | `avmatrix graph-health report --repo <repo> --limit <n> --json` |
| Component inventory | `avmatrix graph-health components --repo <repo> --json` |
| Node/component explanation | `avmatrix graph-health explain <node-or-name> --repo <repo> --json` |
| Query retrieval benchmark | `avmatrix query-health --repo <repo> --suite <file>` |
| ResolutionGap inventory | `avmatrix resolution-inventory --graph .avmatrix/graph.json` |
| Source-site proof accuracy | `avmatrix source-site-accuracy --graph .avmatrix/graph.json` |
| Analyze benchmark comparison | `avmatrix benchmark-compare <before> <after>` |

## Workflow

1. Run `avmatrix analyze --force` first when graph freshness matters.
2. Choose the quality command by failure type: topology, query retrieval, unresolved references, source-site proof, or performance/capacity.
3. Keep threshold and exact query-health results separate. Threshold pass means usable navigation; exact pass means all expected targets were found.
4. Preserve counts, samples, missed targets, and noise reasons in evidence. Do not compress away traceability.
5. If quality output affects an implementation decision, run impact before editing source.

## Query Reliability Guidance

Broad `query` is a candidate discovery command with multiple lanes. It is not a proof that the first ranked result is the owner. For broad intent cases, verify candidates with `context`, exact file inspection, and query-health evidence. When an expected target is missed, record the miss even if threshold pass succeeds.

For the AVmatrix skill-system upgrade suite, use the dedicated query-health suite under `docs/query-health/2026-05-23-avmatrix-skill-system-upgrade-suite.json` when validating AI-context, command-surface, graph-quality, API-surface, and cross-repo discovery.

## Evidence To Record

- Fresh analyze counts before quality commands.
- Query-health case id, expected targets, matched targets, missed targets, hit@5, hit@10, threshold pass, exact pass, and noise reason.
- Graph-health component/candidate counts and representative explanations.
- ResolutionGap classification/actionability counts and top target samples.
- Source-site inventory counts, false-resolved candidates, missing proof counts, and golden fixture results when used.

## Current Limitations

Graph-quality commands report evidence; they do not decide the product fix. A ResolutionGap can be useful diagnostic evidence but is not a resolved in-repo symbol. If quality output is stale or inconsistent, refresh and rerun before making a code decision.
