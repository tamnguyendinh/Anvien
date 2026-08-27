# Child 06A A007 D004 Guarded Diagnostic Typed Read — Coder Handoff

Terminal verdict: `CODER_A007_READY_FOR_MEASUREMENT_HANDOFF`

## Role And Boundary

- Role: Coder only for Child 06A `P2-A`, parent `B1-P1A-OP001 resolution`, child `B2-P2A-A001-D004 project_resolution_outcomes`, attempt `A007`.
- Authorized production owner: only the body of `internal/resolution/outcome.go::resolutionOutcomeDiagnosticSites`, plus one private same-file read-only helper and one minimal import.
- Authorized test owner after production: only new `internal/resolution/outcome_diagnostic_sites_test.go`.
- Authorized durable output: this report only.
- Not performed: target measurement, Supervisor, disposition, ledger edit, detect-changes, staging, commit, push, network, install, branch/worktree/fork/alternate checkout, stash, broad cleanup, broad reset, or checkout.
- Next owner: Main Orchestration task `01a0421e-a38e-7192-8e98-8a09fa72f04d`.

## Start Checkpoint

- Expected and actual HEAD: `e61e55677f8a311913077b1a5d676ce3d4b1777c` (`e61e5567`), subject checkpoint `docs(plan): prepare final D004 attempt`.
- Initial worktree: clean (`git status --porcelain=v2 --untracked-files=all` returned no rows).
- Initial staged set: empty.
- Initial production SHA-256: `internal/resolution/outcome.go` = `02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E`.
- New test and this report did not exist at start.
- The same HEAD, clean status, empty staged set, accepted production hash, and absent new test were rechecked immediately before the first mutation. No other writer owned either allowed source/test path.

## Full Raw Authority Read Through EOF

The following were read in the required rule-before-skill order, in full through EOF. Hashes are the exact inputs consumed by this lane.

| Authority | Lines | SHA-256 |
|---|---:|---|
| `E:\Anvien\AGENTS.md` | 192 | `2B3701EC625D1080A0232228EF7FF6821822173F9006B4AAD0A93B8DE3936DA0` |
| `E:\Anvien\.agents\skills\working-rules\SKILL.md` | 167 | `03DCBDF3F7D2853D19B84B6F4D56EA589A795F73562CF4DD7EADDA0C296C4058` |
| `E:\Anvien\.agents\skills\coder\SKILL.md` | 243 | `F209C2EC606D63C73277A3D29C38212251548A518325361C3CC7564C8075C65A` |
| `E:\Anvien\.agents\skills\backend-development\SKILL.md` | 95 | `11F4E820C3DE5E2090DCA5A73F113E6F24803ABF2ECBFCDDF025BF7A9791DA93` |
| `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\plan-rules.md` | 77 | `5B9C1CC7E32E39AF08D31B4634456182AE5B5D6631AE05713131E670FE0166FA` |
| `2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md` | 673 | `762A907A9F3DC0FCB7FD03007E51FA88D6FCBC43737F3EB43034534DA1949074` |
| `2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md` | 1572 | `65DCCF3C7369DD0CC70BF67701972C8BD5266C3E504A1AB72D15435B0A8D36E3` |
| `2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md` | 762 | `513BA10CED5E946A7EBC157DFBE589E4E4ED7AAA8AE253FB03991982A460C524` |
| `2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md` | 578 | `11CB1B6E3F61D34B2B05431446088CE876FC7B2B2C225AEB07CEE992C1F65603` |
| `E:\Anvien\reports\system-architect\rp_system-architect_260827_by_gpt-5_child06a_d004_diagnostic_roundtrip.md` | 261 | `834D4A0C9297602C12B9EF9C55C4149382743B0103A0E592CA830A58C5AA8D21` |

The Architect report hash matched the delegated authority and its exact verdict was `ARCHITECT_D004_READY_FOR_PLANNER`.

## Invariant Family Map

- Family: D004 final-resolution diagnostic carrier validation.
- Source of truth: the actual `node.Properties[graphhealth.DiagnosticPropertyKey]` value scanned from `g.Nodes` in existing order.
- Primary path: canonical write-through `[]graphhealth.Diagnostic` produced by the current run-scoped appender, then read by `resolutionOutcomeDiagnosticSites`.
- Compatibility path: every noncanonical property type and every exact typed slice whose consumed `SourceSiteID` or `Note` is JSON-unstable.
- Required conversion semantics: compatibility path remains the original whole-property `json.Marshal(value)` followed by `json.Unmarshal(..., []graphhealth.Diagnostic)` and the same wrapped errors.
- Required parity: exact node order, diagnostic order, tracked-site lookup, status rejection, payload-byte equality, returned site membership, first-error timing, non-mutation, Graph JSON, Ladybug/native readers, and public/persisted shapes.
- Forbidden alternate path: sidecar, cache, registry, collector reuse, A005 canonical-outcome-byte lifecycle, graphhealth normalization change, caller/signature change, or graph/slice mutation.
- Resource boundary: canonical path retains only borrowed slice/loop locals and `O(1)` extra retained state; fallback retains its prior allocations.
- Sibling surfaces checked/run-only: graphhealth appender, final outcome carriage, source-backed unresolved diagnostics, analyzer result/capability carriage, Graph JSON/Ladybug persistence, and native Ladybug readback.
- Residual unverified surfaces inside the assigned Coder source/build/test boundary: none. Target measurements and later Supervisor/Main disposition are deliberately outside this lane.

## Fresh Pre-Edit Anvien Gate

1. `anvien --help` — exit `0`, command surface printed.
2. Exactly one pre-edit `anvien analyze --force` from `E:\Anvien` — exit `0`; `2284` scanned, `766` parsed code, `0` failed; graph `124656` nodes / `171522` relationships.
3. `anvien file-detail internal/resolution/outcome.go --repo E:\Anvien --json` — exit `0`. The first shell capture emitted no stdout; the immediate read-only repeat, without another graph refresh, returned current compact JSON at indexed/current HEAD `e61e5567`, stale `false`.
4. `anvien impact symbol resolutionOutcomeDiagnosticSites --repo E:\Anvien --direction upstream --include-tests --json` — exit `0`.

File-detail scope:

- file risk `HIGH`;
- `119` symbols, `0` exported;
- inbound `87`, outbound `122`, local relationships `117`;
- unresolved `148`;
- linked flows `11`, linked tests `23`;
- stale `false`, changed-since-analyze `false`.

Exact edited-symbol blast radius:

- risk `CRITICAL` (scope warning, not an edit ban);
- impacted symbols `7`, direct `1`;
- affected files `6`;
- modules `2`;
- processes `23`;
- app-layer counts `backend=4`, `backend_test=3`;
- functional-area counts `resolution=6`, `analyzer=1`.

The change remained inside the one-body boundary despite this HIGH/CRITICAL scope.

## Production-First Implementation

Production was completed and inspected before the new test file was created.

- Exact `[]graphhealth.Diagnostic` receives a narrow direct-read candidate.
- One private helper, `resolutionOutcomeDiagnosticStringsJSONStable`, pre-scans only the two consumed strings, `SourceSiteID` and `Note`, with `utf8.ValidString`.
- If every consumed string is JSON-stable, projection borrows and iterates the graph-owned typed slice directly.
- If the property type is not exact canonical `[]graphhealth.Diagnostic`, or either consumed string is invalid UTF-8, execution enters the unchanged whole-property marshal/unmarshal compatibility block.
- The existing diagnostic loop, status lookup, tracked/untracked behavior, exact payload comparison, error text/class/order, and site insertion are unchanged.
- No diagnostic object or retained Note is copied; no graph property/slice is mutated, appended, sorted, or replaced.
- No signature, caller, `projectResolutionOutcomes` algorithm, collector, graphhealth owner, persistence shape, reader, I/O, goroutine, lock, cache, sidecar, registry, or cross-run state changed.
- `gofmt -d` was empty for production and test; tracked `git diff --check` exited `0`; new-test trailing-whitespace count was `0`.

## Exact Changed Source/Test Identity

| Path | Start state | Final SHA-256 | Numstat |
|---|---|---|---:|
| `internal/resolution/outcome.go` | tracked; SHA-256 `02092F9F...C1F7E` | `6DFCA0CB912B00FC984DD7B25DB469DC74B87512DC8EC48DFF1ACE0438BF0BB2` | `+20/-6` |
| `internal/resolution/outcome_diagnostic_sites_test.go` | absent | `34ABD93FDAA444899E63DBED6DA8647ADE7819FBF9DBB41DA03DB8DE9D94F25F` | `+349/-0` |

Aggregate production/test numstat: `+369/-6` across exactly two paths. Every existing test file remained byte-identical and run-only. This report is the third and only other authorized path; its self-hash is intentionally not embedded recursively.

## New Test Coverage

All D004 tests are in the sole new file and were created after production inspection:

- line 15, `TestResolutionOutcomeDiagnosticSitesCanonicalTypedReadParityAndNonMutation`: exact canonical membership for unresolved and capability-unavailable statuses, payload parity, ignored untracked diagnostic, graph/property/slice non-mutation, and borrowed-slice identity.
- line 87, `TestResolutionOutcomeDiagnosticSitesFailClosed`: resolved-plus-diagnostic rejection, Note drift rejection, multiple nodes/diagnostics, ignored untracked prefix, and exact first-error traversal.
- line 153, `TestResolutionOutcomeDiagnosticSitesLegacyFallbackParity`: legacy `[]any` / `map[string]any` property matches a test-local copy of the pre-A007 round-trip algorithm and remains unmodified.
- line 202, `TestResolutionOutcomeDiagnosticSitesUnsupportedPropertiesPreserveErrors`: unsupported marshal and unmarshalable property errors match the legacy oracle exactly.
- line 237, `TestResolutionOutcomeDiagnosticSitesJSONUnstableTypedStringsUseFallback`: invalid UTF-8 in typed `SourceSiteID` and `Note` is rejected by the guard, follows round-trip semantics, matches the legacy oracle, and does not mutate/replace the graph-owned slice.

## Holder Handling

Commands before the canonical build:

- `anvien doctor locks --repo E:\Anvien --json` — exit `0`; analyze lock `free`, lock file absent.
- `anvien doctor processes --json` — exit `0`; only editor-owned installed npm/Codex `anvien mcp` processes were reported.

No process was proven to hold the repository's build output, so no process was terminated.

## Canonical Full Build

Exact command:

```text
pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1
```

Result: exit `0`, PASS before every test execution.

- vendor verification: `1798` files, `0` missing, `0` extra, `0` mismatched;
- canonical runtime version `1.2.8`;
- canonical runtime SHA-256 `01A238FF23EA871E97C1BD7F3FF1278574F0DBBF3D79C45CB39390D7DF4DBE94`;
- launcher/Vite build PASS; only the printed existing dynamic-import and chunk-size warnings;
- maintenance analyze exit `0`: `2285` scanned, `767` parsed, `0` failed; graph `124787` nodes / `171689` relationships.

No build/test duration is used as product-performance evidence.

## Post-Build Focused And Named Regressions

New focused command:

```text
go test ./internal/resolution -run '^TestResolutionOutcomeDiagnosticSites' -count=1 -v
```

Result: exit `0`, every new test/subtest PASS.

Named run-only commands and results:

| Exact command | Exit | Result |
|---|---:|---|
| `go test ./internal/graphhealth -run '^(TestDiagnosticAppenderMatchesLegacySemantics|TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity)$' -count=1 -v` | 0 | both PASS |
| `go test ./internal/resolution -run '^(TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix|TestResolveAttachesSourceBackedUnresolvedDiagnostics)$' -count=1 -v` | 0 | both PASS |
| `go test ./internal/analyze -run '^(TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage|TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus)$' -count=1 -v` | 0 | both PASS |
| `go test ./internal/lbugload -run '^TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity$' -count=1 -v` | 0 | PASS |
| untagged `go test ./internal/lbugnative -run '^TestP6C3NativeLadybugResolutionOutcomeReadback$' -count=1 -v` | 0 | `[no tests to run]`; explicitly rejected as evidence |
| tagged/native command below | 0 | `TestP6C3NativeLadybugResolutionOutcomeReadback` actually ran and PASS (`8.16s`) |

The real native invocation used the canonical repo-owned Ladybug root, `CGO_ENABLED=1`, `CGO_CFLAGS=-I.../windows-x86_64`, `CGO_LDFLAGS=-L.../windows-x86_64 -llbug_shared`, native PATH prefix, and:

```text
go test -tags ladybugdb ./internal/lbugnative -run '^TestP6C3NativeLadybugResolutionOutcomeReadback$' -count=1 -v
```

## Package Regressions

| Exact command | Exit | Truthful result |
|---|---:|---|
| `go test ./internal/resolution -count=1` | 1 | only known preserve-only `TestProofBasedCallAccessGoldenCorpus` failure; package is not called PASS |
| `go test ./internal/graphhealth -count=1` | 0 | PASS |
| `go test ./internal/analyze -count=1` | 0 | PASS |
| `go test ./internal/lbugload -count=1` | 0 | PASS |
| `go test ./internal/lbugnative -count=1` | 0 | PASS |

The sole resolution failure is the exact historical callback source site `SourceSite:src/golden.ts#call#callback#14#2#14#12`, with classification/actionability `unclassified/review` and the same structured outcome Note already recorded by the accepted A001–A005 authority. The preserve-only golden file was not edited: worktree and HEAD Git blob are both `00d0ae042f109cd7d3c8aeadf855d192a5aa8172`. No new or changed failure exists.

## STOP And Preserve Checks

- Additional production owner or symbol required: no.
- Existing test edit required: no.
- Signature/caller/coordinator/projector change: no.
- Graphhealth/appender/policy/collector/resolve/emit change: no.
- Compatibility fallback removed, loosened, coerced, or silently dropped: no.
- JSON-unstable typed data consumed unchecked: no; test-proven fallback.
- Graph/slice mutation, sort, copy of retained diagnostic objects, or output/order/error-first drift: no.
- Sidecar/cache/registry/cross-run state/goroutine/lock/I/O/lifecycle added: no.
- Public/persisted/Graph JSON/Ladybug/native/reader shape changed: no.
- A005 canonical outcome-byte lifecycle reused: no.
- Canonical build failure or new/changed test failure: no.
- Target measurement or performance claim made: no.
- Protected A003, D001/D002 history, cost-floor rows, A001–A006/WAL bytes, ledgers, scripts, targets, and existing tests changed: no.

`anvien detect-changes` was intentionally not run: `E2-P2A-A007ARCH1/A007PLAN1` locks detect, stage, and commit until after target measurements, the per-attempt Supervisor, and Main disposition. No stage or commit occurred in this lane.

## Final Handoff State

- HEAD remains `e61e55677f8a311913077b1a5d676ce3d4b1777c`.
- Candidate source/test scope is exact and validation-complete for Coder handoff.
- Staged set is empty.
- Target measurement, Supervisor, Main disposition, Owner-scope closure, detect, stage, and commit remain pending and owned by later lanes.
- Next owner: Main Orchestration task `01a0421e-a38e-7192-8e98-8a09fa72f04d`.

Terminal verdict: `CODER_A007_READY_FOR_MEASUREMENT_HANDOFF`
