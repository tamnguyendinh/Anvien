# Anvien vs GitNexus Deep Comparison Report

Date: 2026-06-01

Evidence:

- Plan: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-plan.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-plan.md)
- Evidence ledger: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-evidence.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-evidence.md)
- Benchmark ledger: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-benchmark.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-benchmark.md)

## Executive Conclusion

On this Windows machine and these pinned commits, Anvien is faster at cold full-repo analysis and exposes a richer graph-correctness surface. GitNexus is also a mature, independent code-intelligence system with a strong TypeScript implementation, worker/cache-oriented analyzer internals, npm packaging, substantial test inventory, and concise runtime outputs.

The practical split is:

- Anvien wins on cold analyze speed, graph diagnostic depth, source-site proof, ResolutionGap inventory, benchmark output, and graph-health workflows.
- GitNexus wins or is stronger on concise lookup latency for sampled `query`/`context`/`impact`, compact `meta.json` capabilities, npm distribution ergonomics, and explicit runtime degradation messages.
- Both passed the reduced deterministic accuracy audit on the Anvien target. Anvien has higher confidence for relationship accuracy because it exposes exhaustive source-site/false-positive audit commands; GitNexus did not expose an equivalent unresolved/false-positive inventory command.

## Methodology

The comparison used code/source reads, build/runtime commands, generated graph/index output, and deterministic source-backed samples. It did not use GitHub popularity signals, stars, README claims as proof, or fork assumptions.

Pinned commits:

| Repo | Commit |
|---|---|
| Anvien | `7b4d48d9bf44b5aa0c6f394861a7d356929521cb` |
| GitNexus | `ce7f45e18d8dceedbcecffad83e5ae23ca105149` |

All GitNexus work and clean benchmark targets were kept under `%TEMP%\anvien-gitnexus-comparison`, outside `E:\Anvien`.

## Performance

Cold full analysis:

| Tool | Target | Elapsed seconds | Files | Nodes | Relationships | Output/index bytes |
|---|---|---:|---:|---:|---:|---:|
| Anvien | Anvien | 41.5269639 | 810 | 95845 | 131188 | 326917002 |
| GitNexus | Anvien | 69.3610866 | 809 | 23121 | 60428 | 240083880 |
| Anvien | GitNexus | 85.377302 | 1339 | 225455 | 245957 | 823484355 |
| GitNexus | GitNexus | 227.4896233 | 1339 | 31622 | 50171 | 339238569 |

Anvien was about 1.67x faster on the Anvien target and about 2.66x faster on the GitNexus target. Warm/incremental runs were not mixed into this table because the two tools expose different no-op/incremental semantics.

Sample query/runtime latency on the Anvien target favored GitNexus for direct context and impact:

| Operation | Anvien seconds | GitNexus seconds |
|---|---:|---:|
| Concept query | 7.6748876 | 4.1021882 |
| Symbol context | 7.8181696 | 2.6881255 |
| Impact analysis | 7.5394923 | 2.9984849 |

## Graph and Accuracy

Anvien produced larger graphs and more relationship categories:

| Tool | Target | Relationship types | Flows/processes | Unresolved/gaps |
|---|---:|---:|---:|---:|
| Anvien | Anvien | 14 | 700 | 69807 |
| Anvien | GitNexus | 21 | 381 | 191224 |
| GitNexus | Anvien | 10 | 300 | not exposed |
| GitNexus | GitNexus | 13 | 300 | not exposed |

The reduced deterministic audit used shared Anvien source facts for file nodes, symbol declarations, methods/functions, imports/dependencies, and calls/references. Both tools scored 100 percent on the sampled facts. The important difference is auditability:

- Anvien reported 0 false-resolved edge candidates and 0 resolved edges missing source-site proof in source-site accuracy output.
- GitNexus passed the sampled facts, but no equivalent source-site accuracy or ResolutionGap inventory command was exposed.

Anvien's weakness is visible too: unresolved volume is high. On clean Anvien, Anvien reported 69,807 unresolved references, including 41,091 in-repo analyzer gaps. On GitNexus, it reported 191,224 unresolved references, including 178,924 in-repo analyzer gaps.

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
- direct Node CLI context/impact latency was lower in the sample;
- strong runtime recovery messages around native bindings, OOM, WAL/FTS, and vector degradation;
- npm package distribution plus Docker/Web bundling.

## Maturity

Both are mature enough to build and run locally on this machine.

| Dimension | Anvien | GitNexus |
|---|---|---|
| Build/install | Passed in about 49.5s | Passed, but cold install/build took 211.2s |
| Tests inventory | 174 Go test files, 65 Web test/e2e/spec files | 429 core TS test files, 34 Web test/e2e/spec files |
| CI workflows | 14 | 22 |
| Packaging | Windows launcher/native runtime helpers | npm bin package, Dockerfiles, bundled Web |
| Diagnostics | Strong graph/source-site diagnostics | Strong runtime/storage diagnostics |

GitNexus has more raw core test files and workflows. Anvien has stronger graph-quality observability and faster local product rebuild/analyze loops.

## Risks and Limitations

- Benchmark runs were cold full rebuilds only.
- Accuracy scoring was a deterministic reduced audit, not a complete golden corpus over every supported language.
- GitNexus vector search was unavailable on this Windows platform in both `meta.json` files; embeddings were 0.
- Anvien graph size is not automatically better by itself; the stronger claim comes from source-site proof and diagnostics, not only count volume.
- GitNexus source-line output appears 0-based in sampled Cypher/context results; the audit treated it as correct with a convention note.

## Actionable Lessons for Anvien

1. Add a compact `meta.json`-style index summary with stats and capability statuses.
2. Make degraded capability states as explicit as GitNexus does for vector/FTS/native runtime paths.
3. Investigate context/impact CLI latency; GitNexus was faster on the sampled lookup commands.
4. Borrow the concept of parse-cache/worker-pool ergonomics where it fits Anvien's Go pipeline.
5. Keep strengthening unresolved-reference reduction; Anvien's diagnostic surface is strong, and the next maturity gain is reducing the in-repo analyzer-gap count.
