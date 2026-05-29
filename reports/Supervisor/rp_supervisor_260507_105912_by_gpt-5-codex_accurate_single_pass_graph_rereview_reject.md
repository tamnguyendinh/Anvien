# Supervisor Report: Accurate Single-Pass Graph Re-review

Verdict: REJECT

Scope reviewed:
- Plan: `docs/plans/2026-05-06-accurate-single-pass-graph-plan.md`
- Code window: `ba2f0e4..8536695`
- Current head: `8536695 Implement accurate single-pass graph optimization`
- Authority used: the specified plan plus current code/tests/benchmark artifacts in this repo only. No SPEC, architecture doc, or other markdown authority was used for the verdict.

## Critical Issues

None found.

## High Issues

None found. The previous runtime blockers are closed:
- Same-file same-name TypeScript methods now emit owner-qualified defs and map to real graph nodes. Repro now emits `CALLS` to `Method:src/app.ts:A.save` and `Method:src/app.ts:B.save`, not `def:*`.
- Duplicate legacy/scope edges now merge scope audit metadata. Repro keeps one `CALLS` edge with `resolutionSource`, `evidence`, and `fileHash`.
- Full `npm test` now passes.

## Medium Issues

### [MEDIUM] Final benchmark artifacts still do not satisfy the plan's reproducibility/language-coverage protocol

File: `anvien/src/core/analyze/analyze-benchmark-snapshot.ts:60`

Issue: the plan's benchmark protocol requires commit hashes and language coverage for every benchmarked repo. Current benchmark schema/key metrics only expose aggregate parse/scope counters (`parseableFiles`, `scopeParsedFiles`, `scopeExtractionAstReusedFiles`, etc.) and do not define any per-language coverage field (`analyze-benchmark-snapshot.ts:60`, `analyze-benchmark-snapshot.ts:73`, `analyze-benchmark-snapshot.ts:77`, `analyze-benchmark-snapshot.ts:82`). The source counters have the same aggregate-only shape (`analyze-metrics.ts:25`, `analyze-metrics.ts:41`, `analyze-metrics.ts:46`).

The environment schema has optional `repoGitCommit` and `repoGitDirty` fields (`analyze-benchmark-snapshot.ts:31`, `analyze-benchmark-snapshot.ts:36`), and `collectBenchmarkEnvironment` only writes them when `git -C <repoPath>` succeeds (`analyze.ts:435`, `analyze.ts:443`). The final repeated benchmark artifacts in `reports/benchmark/2026-05-07-anvien-final-equivalent-accuracy-run*.json` contain `environment`, but no `repoGitCommit`, no `repoGitDirty`, and no explicit unavailable/non-git marker.

Observed final artifact evidence:

```text
reports/benchmark/2026-05-07-anvien-final-equivalent-accuracy-run1-gitnexus-main.json
environment: anvienVersion, nodeVersion, platform, arch
missing: repoGitCommit, repoGitDirty, language coverage
key metrics: parseableFiles=750, scopeParsedFiles=750, scopeExtractionAstReusedFiles=750
```

Why this blocks plan approval: the plan's final acceptance is not only "tests pass"; it also requires benchmark results to follow the protocol and report language coverage. Aggregate `750/750` scope coverage proves global coverage for that run, but it does not publish which languages were covered vs legacy shares. Missing git commit/dirty metadata also makes the final benchmark point less reproducible.

Fix:
- Add a benchmark field such as `languageCoverageByLanguage` with parseable files, AST-reused scope files, compatibility/no-hook/failed counts, resolved/unresolved reference counts where available, and covered-vs-legacy share per language.
- Make benchmark git metadata explicit: either require `repoGitCommit`/`repoGitDirty` for benchmark artifacts, or write `repoGitUnavailable: true` plus the reason so the artifact is auditable.
- Regenerate the final repeated benchmark artifacts with the new fields and update the benchmark comparison output/tests to include them.

## Suggestions

- Keep `--skip-legacy-cross-file` documented as diagnostic. The automatic skip is guarded by complete AST-reused scope coverage (`cross-file.ts:80`, `cross-file.ts:122`), and current benchmark artifacts show parity for the measured repo, but language-coverage reporting should make that guard inspectable per language before this is treated as fully complete.
- Consider adding a test that asserts every relationship emitted by `emitReferencesToGraph` has existing source/target graph nodes. The code now fail-closes unresolved mappings (`emit-references.ts:127`, `emit-references.ts:133`, `emit-references.ts:421`), but a broad invariant test would protect future emitters.

## Source-Level Clearance Notes

- Graph emission/runtime mapping: cleared. `emitReferencesToGraph` now skips missing caller/target mappings (`emit-references.ts:127`, `emit-references.ts:133`) and `createGraphNodeResolver` returns `undefined` instead of falling back to `def.nodeId` (`emit-references.ts:415`, `emit-references.ts:421`). Duplicate edge handling now merges audit metadata into the existing edge (`emit-references.ts:139`, `emit-references.ts:141`, `emit-references.ts:577`).
- TypeScript/JavaScript scope capture: cleared for the reviewed blocker family. Methods/properties now emit owner-qualified names (`typescript-javascript.ts:145`, `typescript-javascript.ts:154`, `typescript-javascript.ts:166`), and `emitNamedDeclaration` forwards `@declaration.qualified_name` (`typescript-javascript.ts:750`, `typescript-javascript.ts:766`).
- Python scope capture: cleared for the added provider slice. Python uses AST-reused capture emission (`python.ts:33`, `python.ts:36`), emits qualified method/property names (`python.ts:118`, `python.ts:120`, `python.ts:141`, `python.ts:149`), and emits member call/access facts from tree nodes (`python.ts:298`, `python.ts:305`, `python.ts:319`).
- Resolver/worker path: cleared. Default resolution calls `resolveScopeReferenceSitesInWorkers` (`resolution.ts:105`); auto mode falls back to serial below threshold (`scope-reference-resolver.ts:156`) and forced worker mode serializes indexes once via workerData (`scope-reference-resolver.ts:161`, `scope-reference-resolver.ts:164`), then merges chunks deterministically (`scope-reference-resolver.ts:300`).
- Cross-file overlap: conditionally cleared for measured covered repos. Legacy cross-file is skipped only when every parseable file has AST-reused scope facts (`cross-file.ts:80`, `cross-file.ts:129`) or by explicit diagnostic option (`cross-file.ts:76`). The remaining approval blocker is benchmark evidence granularity, not this guard's basic wiring.
- Benchmark/metrics: blocked by the medium finding. Semantic relationship metrics were added (`analyze-benchmark-snapshot.ts:197`, `analyze-benchmark-snapshot.ts:315`), but benchmark output still lacks required language coverage and final artifacts lack explicit target repo git state.

## Validation Run

Passed:
- `git diff --check ba2f0e4..HEAD`
- `cd anvien && npm run build`
- `cd anvien && npx tsc --noEmit`
- `cd anvien && npx vitest run test/unit/scope-resolution/typescript-single-pass-parity.test.ts test/unit/scope-resolution/python-single-pass-parity.test.ts test/unit/scope-resolution/scope-reference-resolver.test.ts test/unit/scope-resolution/emit-references.test.ts test/unit/scope-resolution/finalize-orchestrator.test.ts test/unit/cross-file.test.ts test/unit/analyze-benchmark-snapshot.test.ts test/unit/benchmark-compare-command.test.ts --pool=vmForks` (`8` files, `70` tests)
- `cd anvien && npm test` (exit code `0`)

Additional repro checks:
- Same-file duplicate method mapping now resolves to real graph node ids.
- Duplicate legacy edge audit metadata is now merged into the persisted in-memory relationship.

## Required Fix List For Resubmission

1. Add per-language benchmark coverage reporting to code and artifacts.
2. Make target repo git state explicit in benchmark artifacts, including a clear unavailable marker if commit/dirty cannot be read.
3. Regenerate final benchmark artifacts and re-run benchmark comparison/tests.

## Overall Coder Evaluation

The coder closed the prior runtime blockers with focused changes and materially improved the implementation: graph emission now fail-closes, duplicate audit metadata is preserved, TypeScript/Python scope coverage is broader, cross-file narrowing is guarded, workerized resolution is parity-tested, and the full test suite passes. The remaining issue is evidence completeness against the plan's benchmark protocol, not core runtime correctness.
