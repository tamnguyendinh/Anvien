# Anvien vs GitNexus Deep Comparison Report

Date: 2026-06-01

Status: regenerated after Anvien update

Evidence:

- Plan: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-plan.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-plan.md)
- Evidence ledger: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-evidence.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-evidence.md)
- Benchmark ledger: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-benchmark.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-benchmark.md)

## Executive Conclusion

On this Windows machine and these pinned commits, Anvien still wins cold full-repo graph generation across all three benchmark targets. The gap is largest on the Anvien target, still clear on GitNexus, and still clear but smaller on Restaurant_manager.

GitNexus remains strong and mature as an independent code-intelligence system. It is slower at full indexing in this rerun, but it remains faster for sampled direct lookup commands on the Anvien target and has concise runtime output, npm packaging, and explicit vector/FTS/native-runtime degradation messages.

The practical split is unchanged:

- Anvien wins on cold analyze speed, graph diagnostic depth, source-site proof, ResolutionGap inventory, benchmark JSON, and graph-health workflows.
- GitNexus wins on sampled query/context/impact lookup latency, compact `meta.json` status, npm distribution ergonomics, and explicit runtime degradation reporting.
- Both passed the earlier reduced deterministic accuracy audit. Anvien still has higher audit confidence because it exposes exhaustive source-site/false-positive checks; GitNexus does not expose an equivalent unresolved/false-positive inventory command.

## Methodology

The stale benchmark report was deleted before rerunning because Anvien changed after the previous benchmark. The rerun used clean target clones outside `E:\Anvien` under `E:\avgn-rerun`.

Pinned commits:

| Repo | Commit |
|---|---|
| Anvien | `97a45525820c609410796b1f11fa38239e31cbfa` |
| GitNexus | `ce7f45e18d8dceedbcecffad83e5ae23ca105149` |
| Restaurant_manager | `fdfacba78e5445522dd09cca98fa27d39e0e22c8` |

Anvien used the freshly built local binary at `E:\Anvien\anvien\bin\anvien.exe`. GitNexus used the built CLI at `E:\avgn-rerun\tools\GitNexus\gitnexus\dist\cli\index.js`.

## Performance

Cold full analysis rerun:

| Tool | Target | Wall seconds | Tool-reported seconds | Files | Nodes | Relationships/edges | Output/index bytes |
|---|---|---:|---:|---:|---:|---:|---:|
| Anvien | Anvien | 37.985012 | 37.238526 | 816 | 96211 | 131684 | 327984313 |
| GitNexus | Anvien | 117.262267 | 113.8 | 815 | 23264 | 60687 | 240978915 |
| Anvien | GitNexus | 100.934158 | 99.7532442 | 1339 | 225455 | 245957 | 823484355 |
| GitNexus | GitNexus | 215.510712 | 210.3 | 1339 | 31622 | 50171 | 339209809 |
| Anvien | Restaurant_manager | 93.133052 | 91.4707017 | 6198 | 202810 | 253342 | 657873587 |
| GitNexus | Restaurant_manager | 173.754737 | 171.6 | 6198 | 72792 | 143910 | 641866383 |

Wall-clock speed ratio:

| Target | Anvien advantage |
|---|---:|
| Anvien | 3.09x faster |
| GitNexus | 2.13x faster |
| Restaurant_manager | 1.87x faster |

The large target confirms Anvien keeps a meaningful speed advantage at 6,198 files while producing more graph facts. The ratio is not purely file-count-driven: the Anvien target produced the largest relative gap, while Restaurant_manager produced the smallest of the three rerun gaps.

Compared with the previous run, Anvien improved on the Anvien self target despite scanning more files. The GitNexus and Restaurant_manager Anvien runs were slower than before while graph counts stayed stable. Treat that as a rerun measurement, not standalone proof of algorithmic regression, because workspace path and cold disk/cache state changed.

Sample lookup latency on the Anvien target still favored GitNexus:

| Operation | Anvien seconds | GitNexus seconds |
|---|---:|---:|
| Concept query | 7.5428505 | 4.73722 |
| Symbol context | 7.8183634 | 2.9433826 |
| Impact analysis | 10.8241073 | 2.3872532 |

## Graph and Accuracy

Graph inventory rerun:

| Tool | Target | Relationship types | Flows/processes | Unresolved/gaps |
|---|---|---:|---:|---:|
| Anvien | Anvien | 14 | 700 | 69988 |
| Anvien | GitNexus | 21 | 381 | 191224 |
| Anvien | Restaurant_manager | 15 | 508 | 129135 |
| GitNexus | Anvien | 10 | 300 | not exposed |
| GitNexus | GitNexus | 13 | 300 | not exposed |
| GitNexus | Restaurant_manager | 9 | 300 | not exposed |

Anvien source-site accuracy reported 0 false-resolved edge candidates and 0 resolved edges missing source-site proof on all three Anvien-generated graphs. GitNexus passed the earlier reduced deterministic sample audit, but no equivalent source-site accuracy or ResolutionGap inventory command was exposed.

Anvien's visible weakness remains unresolved volume:

| Target | Unresolved references | In-repo analyzer gaps |
|---|---:|---:|
| Anvien | 69988 | 41200 |
| GitNexus | 191224 | 178924 |
| Restaurant_manager | 129135 | 80306 |

## Functionality

Both repos are broad code-intelligence systems rather than simple graph visualizers.

Anvien strengths:

- graph-health, query-health, source-site-accuracy, resolution-inventory, benchmark-compare;
- rich impact/detect-changes output with app layers, functional areas, files, flows, and tests;
- MCP stdio, MCP-over-HTTP, resources, prompts, Web/API, groups, and packaging/runtime helpers;
- first-class benchmark JSON from analyze.

GitNexus strengths:

- worker/cache-oriented TypeScript ingestion internals;
- compact `meta.json` with stats and capability status;
- direct Node CLI context/impact latency remains lower in the sample;
- strong runtime recovery messages around native bindings, OOM, WAL/FTS, and vector degradation;
- npm package distribution plus Docker/Web bundling.

## Maturity

Both are mature enough to build and run locally on this machine.

| Dimension | Anvien | GitNexus |
|---|---|---|
| Build/install | Passed in 42.996s | Passed, but cold install/build took 182.624685s across shared/core |
| Tests inventory | 174 Go test files, 65 Web test/e2e/spec files | 429 core TS test files, 34 Web test/e2e/spec files |
| CI workflows | 14 | 22 |
| Packaging | Windows launcher/native runtime helpers | npm bin package, Dockerfiles, bundled Web |
| Diagnostics | Strong graph/source-site diagnostics | Strong runtime/storage diagnostics |

GitNexus has more raw core test files and CI workflows. Anvien has stronger graph-quality observability and faster local product rebuild/analyze loops.

## Risks and Limitations

- Benchmark runs were cold full rebuilds only.
- Accuracy scoring was a deterministic reduced audit, not a complete golden corpus over every supported language.
- GitNexus vector search was unavailable on this Windows platform in all rerun `meta.json` files; embeddings were 0.
- Anvien graph size is not automatically better by itself; the stronger claim comes from source-site proof and diagnostics, not only count volume.
- Rerun timing can be affected by disk/cache/workspace changes, so stable graph-count deltas are stronger evidence than single-run timing deltas.

## Actionable Lessons for Anvien

1. Add a compact `meta.json`-style index summary with stats and capability statuses.
2. Make degraded capability states as explicit as GitNexus does for vector/FTS/native runtime paths.
3. Investigate `impact` latency specifically; the rerun impact sample moved to 10.8241073s while GitNexus completed in 2.3872532s.
4. Borrow the concept of parse-cache/worker-pool ergonomics where it fits Anvien's Go pipeline.
5. Keep reducing in-repo analyzer gaps; Anvien's diagnostic surface is strong, and the next maturity gain is lowering the unresolved count.
