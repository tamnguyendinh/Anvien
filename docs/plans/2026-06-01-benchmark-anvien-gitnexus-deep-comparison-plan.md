# Anvien GitNexus Deep Comparison Plan

Date: 2026-06-01

Status: rerun completed after Anvien update

Companion files:

- Evidence ledger: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-evidence.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-evidence.md)
- Benchmark ledger: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-benchmark.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-benchmark.md)
- Final report: [2026-06-01-benchmark-anvien-gitnexus-deep-comparison-report.md](2026-06-01-benchmark-anvien-gitnexus-deep-comparison-report.md)

## Master Rules

1. Use Anvien for Anvien codebase analysis and impact checks when graph-based evidence is needed.
2. Run `anvien analyze --force` before any Anvien graph query, context, impact, graph-health, benchmark, or detect-changes command.
3. Do not use GitHub stars, popularity, fork status, README claims, or marketing language as evaluation evidence.
4. Read implementation code deeply before judging architecture, accuracy, performance, or maturity.
5. Compare both tools on the same local machine, same commit inputs, and same benchmark corpus wherever their command surfaces allow it.
6. Keep GitNexus cloned only in a temporary location outside `E:\Anvien`; never create the GitNexus clone inside this repository because other agents may be working in this checkout.
7. Record benchmark numbers as soon as each benchmarkable task completes.
8. Record evidence as soon as each evidenced task completes.
9. If any Web UI behavior is changed in Anvien during a future implementation slice, run a full build and include e2e coverage. This comparison plan itself is doc-only and does not change Web UI behavior.
10. If this work later leads to implementation edits in Anvien, run impact analysis before editing functions, classes, methods, exported symbols, API handlers, graph builders, resolvers, analyzers, or shared contracts.
11. If this work later leads to implementation commits, run `anvien detect-changes --repo Anvien --scope all` before each commit.

## Goal

Produce an evidence-backed technical comparison between Anvien and GitNexus as independent code-graph/code-intelligence systems.

The final report should answer:

- which tool is faster and on which workloads;
- which tool extracts a more complete and accurate graph;
- which tool has broader and more useful functionality;
- which implementation is more mature and operationally usable;
- where each repo has real architectural strengths or weaknesses;
- what Anvien can learn from GitNexus without copying code or assuming fork lineage.

## Repositories

| Repo | Role | Source |
|---|---|---|
| Anvien | Local product under evaluation | `E:\Anvien`, `https://github.com/tamnguyendinh/Anvien` |
| GitNexus | Independent comparison target | Temporary clone from `https://github.com/abhigyanpatwari/GitNexus` |

Temporary clone convention:

```powershell
$root = Join-Path $env:TEMP "anvien-gitnexus-comparison"
$gitnexus = Join-Path $root "GitNexus"
git clone https://github.com/abhigyanpatwari/GitNexus $gitnexus
```

Required path safety check before cloning:

```powershell
$anvienRoot = (Resolve-Path "E:\Anvien").Path
$tempRoot = [System.IO.Path]::GetFullPath($root)
if ($tempRoot.StartsWith($anvienRoot, [System.StringComparison]::OrdinalIgnoreCase)) {
    throw "GitNexus temp clone path must be outside E:\Anvien"
}
```

Cleanup after final evidence capture:

```powershell
Remove-Item -LiteralPath $root -Recurse -Force
```

## Scope

In scope:

- Deep source reading of both repositories.
- Build, install, and command-surface discovery for GitNexus.
- Fresh Anvien graph refresh before Anvien graph-based analysis.
- Benchmarking both tools against identical target repositories where possible.
- Inventory comparison of graph output: files, symbols, relationships, relationship types, unresolved references, execution flows, and metadata.
- Accuracy audit using deterministic samples from the generated graph output and source code.
- Feature comparison across CLI, API, Web UI, MCP/agent integration, graph query, impact analysis, health diagnostics, multi-repo support, packaging, docs, and tests.
- Maturity assessment based on implementation evidence: architecture, tests, error handling, persistence, configuration, release/build path, portability, and maintainability.
- Final comparative report in `docs/plans` or a follow-up report path selected before execution.

Out of scope:

- Using popularity metrics as quality evidence.
- Treating README claims as true without code or runtime verification.
- Copying implementation code between repositories.
- Permanent vendoring or retaining the GitNexus clone in this repo.
- Changing Anvien functionality during the comparison unless a separate implementation task is opened.

## Benchmark Corpus

Use pinned commits for all target repositories and record every SHA in the evidence ledger.

Primary corpus:

| Target | Reason |
|---|---|
| Anvien repo | Large, known local graph system; validates behavior on this project's real workload. |
| GitNexus repo | Lets both tools analyze the comparison target. |
| Restaurant_manager repo | Larger local workload used to stress scaling behavior across a 6,198-file target. |

Optional neutral corpus, only if needed for fairness:

| Target | Reason |
|---|---|
| Small mixed-language fixture repo | Reduces bias from each tool analyzing its own source tree. |
| Medium public repo pinned to SHA | Gives a third workload if command compatibility allows it. |

## Benchmark Dimensions

Performance metrics:

- cold analyze elapsed time;
- warm analyze elapsed time;
- files scanned or accepted;
- files parsed or indexed;
- graph nodes;
- graph relationships;
- unresolved references or equivalent gaps;
- output/index size on disk;
- peak memory if measurable without intrusive instrumentation;
- startup time for CLI/API runtime if supported;
- query latency for comparable graph queries if both tools expose query commands.

Accuracy metrics:

- symbol detection recall on deterministic samples;
- relationship detection recall on deterministic samples;
- source location correctness;
- import/dependency correctness;
- member/method/property ownership correctness where supported;
- call/reference correctness where supported;
- unresolved reference rate and representative causes;
- false positive examples found during manual audit.

Functional metrics:

- language coverage;
- CLI coverage;
- API/Web UI coverage;
- graph query support;
- impact analysis support;
- graph health/diagnostic support;
- benchmark/reporting support;
- multi-repo support;
- agent/MCP integration;
- docs and onboarding depth.

Maturity metrics:

- test coverage and test realism;
- build reliability on Windows;
- configuration and error handling;
- persistence/index format clarity;
- packaging/release path;
- internal architecture boundaries;
- ability to run repeatedly without manual cleanup;
- observability and debuggability.

## Implementation Plan

- [x] [P0] Create a clean comparison workspace and capture host environment details: OS, CPU, RAM, Go/Node/Python versions, Git version, and PATH-relevant tool versions.
- [x] [P1] Refresh Anvien graph with `anvien analyze --force` before graph-based discovery.
- [x] [P2] Deep-read Anvien implementation using Anvien context/query plus source reads for analyzer pipeline, graph schema, resolver behavior, query surfaces, Web/API/MCP support, benchmarks, tests, and packaging.
- [x] [P3] Clone GitNexus into a temporary directory, record commit SHA, dependency manifests, project layout, and build/run instructions found in code and scripts.
- [x] [P4] Deep-read GitNexus implementation directly from source for analyzer pipeline, graph schema, resolver behavior, query surfaces, UI/API support, benchmarks, tests, and packaging.
- [x] [P5] Build or install both tools from the checked-out commits and record any setup failures or patches avoided.
- [x] [P6] Run Anvien on the Anvien repo and record performance plus graph inventory.
- [x] [P7] Run Anvien on the GitNexus repo and record performance plus graph inventory.
- [x] [P8] Run GitNexus on the Anvien repo, if supported, and record performance plus graph inventory.
- [x] [P9] Run GitNexus on the GitNexus repo and record performance plus graph inventory.
- [x] [P10] If GitNexus cannot analyze one of the target repos, document the exact command, error, and functional implication instead of inventing replacement numbers.
- [x] [P11] Normalize benchmark tables so every metric clearly distinguishes unavailable, unsupported, failed, and measured values.
- [x] [P12] Perform deterministic accuracy audits on both tools using source-backed samples from the same target repo where possible.
- [x] [P13] Compare features using verified commands, source files, tests, and runtime surfaces.
- [x] [P14] Compare maturity using build/test evidence, architecture reading, error handling, persistence, packaging, docs, and repeated-run behavior.
- [x] [P15] Clean up the temporary GitNexus clone and record cleanup evidence.
- [x] [P16] Write the final comparison report with conclusions separated by category: performance, accuracy, functionality, maturity, risks, and actionable lessons for Anvien.
- [x] [P17] Delete stale benchmark report output after Anvien code changed.
- [x] [P18] Rerun build, analyze, graph-quality, and lookup-latency benchmarks on Anvien, GitNexus, and Restaurant_manager clean targets outside `E:\Anvien`.
- [x] [P19] Regenerate benchmark, evidence, and final report from the rerun data and clean up temporary benchmark clones.

## Evidence Collection Plan

Record source-backed evidence in the evidence ledger for:

- exact commits and local paths used;
- setup/build commands;
- command availability and command help output summary;
- architecture ownership maps for both repos;
- graph schema or index format facts;
- analyzer and resolver code paths;
- CLI/API/Web/MCP surfaces;
- tests and CI/build evidence;
- accuracy sample methodology;
- failures and unsupported cases;
- temp clone cleanup.

## Benchmark Methodology

Minimum benchmark run shape:

```powershell
# Record command availability first.
Get-Command git, go, node, npm, python -ErrorAction SilentlyContinue

# Use Measure-Command for elapsed time when the tool does not emit benchmark JSON.
$elapsed = Measure-Command {
    # tool analyze command here
}
$elapsed.TotalSeconds
```

Preferred Anvien benchmark commands, adjusted after checking the current CLI help:

```powershell
.\anvien\bin\anvien.exe analyze --force --name Anvien --benchmark-json .\reports\benchmarks\anvien-self.json
.\anvien\bin\anvien.exe graph-health summary --repo Anvien --json
```

GitNexus commands must be discovered from the temporary clone before execution. Do not assume CLI names or flags from README alone.

## Accuracy Audit Method

Use a deterministic sample set so the result is reproducible.

Minimum sample categories:

| Category | Minimum samples | Evidence |
|---|---:|---|
| File nodes | 20 | Source file exists and graph includes/excludes it correctly. |
| Top-level symbols | 30 | Source declaration exists and graph symbol metadata is correct. |
| Methods/functions | 30 | Ownership and source location are correct. |
| Imports/dependencies | 20 | Import edge matches source code. |
| Calls/references | 30 | Edge maps to a real call/reference site if the tool claims support. |
| Unresolved/gap examples | 20 | Cause is categorized from source code. |

Scoring rules:

- `correct`: graph fact matches source.
- `partial`: concept is present but location, ownership, or relationship type is incomplete.
- `missing`: source fact should be represented under the tool's claimed scope but is absent.
- `false_positive`: graph fact cannot be traced to source.
- `unsupported`: tool does not claim this capability or cannot represent the fact.

## Final Report Outline

The final report should include:

1. Executive conclusion.
2. Methodology and limitations.
3. Repository and commit inventory.
4. Architecture comparison.
5. Performance benchmark tables.
6. Graph inventory comparison.
7. Accuracy audit results.
8. Feature matrix.
9. Maturity assessment.
10. Failure and unsupported-case analysis.
11. Anvien improvement opportunities.
12. Raw command appendix or links to evidence/benchmark ledgers.

## Validation Plan

This is a documentation and evaluation task. Validation is the recorded evidence itself.

For any future implementation slice triggered by findings:

```powershell
powershell -ExecutionPolicy Bypass -File anvien-launcher/build.ps1
.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all
```

Add or update e2e tests if Web UI behavior changes.

## Success Criteria

1. The comparison uses code and runtime evidence, not reputation signals.
2. Benchmark numbers are reproducible from recorded commands, commits, and environment details.
3. Unsupported or failed GitNexus/Anvien runs are reported explicitly and not converted into fake metrics.
4. Accuracy conclusions cite source-backed samples and classify partial/missing/false-positive cases.
5. Feature and maturity conclusions point to concrete files, commands, tests, or runtime behavior.
6. Temporary GitNexus clone is removed after evidence is saved.
