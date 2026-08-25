# Coder Handoff — Child 06A A002 Run-Scoped Diagnostic Appender

Date: `2026-08-25`
Status: `PAUSED_BY_OWNER_DURING_CANONICAL_FULL_BUILD`
Attempt: `A002` on active unchecked `B2-P2A-A001-D001 resolve_calls`
Authority: architecture commit `dafcf7a25e254ef8db09d28eca3575d5225c553e`; plan commit `80eb8c3c9a9cf516dd17f8b3bf15802694bb1277`

## Current result

The authorized source and focused test bytes are written but are **not build/test validated**. The Owner requested an immediate pause while the canonical full build was running. The build was interrupted with Ctrl+C and exited `1`; this is an interrupted validation, not a product/build failure and not a PASS. No test command has run.

## Invariant family

- Family: per-`ResolveBoundInto` resolution diagnostic append semantics.
- Authority: `E2-P2A-A002ARCH1/PLAN1` and `reports/system-architect/rp_system-architect_260825_184751_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`.
- Primary path: `ResolveBoundInto -> newEmitter -> resolveCall -> emitUnresolvedReference|emitTypeScriptOutcomeDiagnostic -> run-scoped write-through appender -> Graph.AddNode`.
- Preserved sibling path: all non-resolution callers continue using `graphhealth.AppendDiagnosticToNode` unchanged at the call surface.
- Preserved rules: false-return/default behavior, supported first-touch representation normalization, diagnostic bucket identity, count merge, target fill, earliest-line behavior, stable ordering, policy fields, `[]Diagnostic` graph representation, and visibility after every append.
- Resource/lifecycle: one emitter-owned map of `nodeID -> []Diagnostic` slice headers; the graph property and cache receive the same merged slice; no second retained diagnostic-object copy, I/O, goroutine, lock, global cache, flush, or finalizer.

## Changed files and exact hunks

1. `internal/graphhealth/diagnostics.go`
   - Adds private `diagnosticAppender` state and exported internal-package factory `NewDiagnosticAppender`, returning the private append closure.
   - First successful node touch normalizes existing supported diagnostics once; every incoming diagnostic is defaulted and normalized once.
   - Every append re-reads the node, preserves sibling properties, writes `DiagnosticPropertyKey` as `[]Diagnostic`, and calls `Graph.AddNode`.
   - Refactors `appendDiagnosticsFromProperties` through private `appendNormalizedDiagnostic`; legacy `AppendDiagnosticToNode` remains available and behavior-compatible.
2. `internal/resolution/emit.go`
   - Adds the emitter-owned appender function field.
   - Initializes it in `newEmitter` from the same graph.
   - Routes `emitUnresolvedReference` through it.
3. `internal/resolution/outcome.go`
   - Routes `emitTypeScriptOutcomeDiagnostic` through the same emitter-owned appender.
4. `internal/graphhealth/diagnostics_test.go` (new)
   - Compares legacy and run-scoped writes after every append for encoded legacy first-touch input, multiple unique diagnostics, duplicate merge/count, target fill, earliest line, stable order, policy/default fields, graph JSON/write-through, and invalid/absent-node false behavior.

No existing test file, ledger, plan-rules file, graph type, policy/contract, `resolve.go`, persistence/reader, schema, public product contract, or timing instrumentation was edited.

## Pre-edit Anvien evidence

- Exact `anvien --help`: PASS.
- Fresh `anvien analyze --force`: exit `0`; scanned `2217`, parsed `765`, failed `0`; graph `123605` nodes / `170322` relationships; index fresh at commit `80eb8c3c9a9cf516dd17f8b3bf15802694bb1277`.
- File risk:
  - `internal/graphhealth/diagnostics.go`: HIGH; `154` symbols; inbound/outbound `50/70`; linked flows/tests `1/19`.
  - `internal/resolution/emit.go`: HIGH; `154`; `81/231`; `13/24`.
  - `internal/resolution/outcome.go`: HIGH; `119`; `87/124`; `11/23`.
- Symbol impact:
  - `AppendDiagnosticToNode`: CRITICAL; impacted/direct `7/2`; modules/processes `1/24`.
  - `appendDiagnosticsFromProperties`: CRITICAL; `7/1`; `2/13`.
  - `emitter`: CRITICAL; `28/25`; `2/49`.
  - `newEmitter`: CRITICAL; `6/1`; `4/32`.
  - `emitUnresolvedReference`: CRITICAL; `7/4`; `2/36`.
  - `emitTypeScriptOutcomeDiagnostic`: LOW `0/0`, while its file remains HIGH.

HIGH/CRITICAL were treated as blast-radius warnings; the candidate remains inside the exact Architect boundary.

## Source inspection completed

- Production-first order: PASS.
- Production diff inspection: PASS; scoped and surgical.
- `gofmt` on the four authorized paths: complete.
- Scoped `git diff --check` before build: PASS.
- At the last boundary, the only A002 paths were the three modified production files plus the new test file. All pre-existing unrelated dirty work was preserved and untouched.

## Build state at pause

Command started exactly:

`pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`

Preflight before launch:

- `anvien doctor locks --repo E:\Anvien --json`: analyze lock free/absent.
- `anvien doctor processes --json`: no process using the repo-local build output; listed Anvien/Go processes were editor-owned MCP services, so none was terminated.

Observed before Owner pause:

- vendor/runtime source verification reported `PASS`;
- canonical runtime build wrote `E:\Anvien\anvien\bin\anvien.exe`;
- launcher asset build began;
- Owner then requested immediate pause;
- build was interrupted with Ctrl+C and exited `1`.

Therefore `E2-P2A-A002BUILD1` is **not satisfied**.

## Tests

No tests have run. The next Coder must not reuse any PASS claim from this handoff.

After a fresh build-holder/lock preflight, rerun the canonical full build from the beginning. Only after full-build PASS, run in this order:

1. `go test ./internal/graphhealth -run '^TestDiagnosticAppenderMatchesLegacySemantics$' -count=1 -v`
2. `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`
3. `TestResolveAttachesSourceBackedUnresolvedDiagnostics`
4. `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`
5. `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`
6. `go test ./internal/graphhealth -count=1`
7. `go test ./internal/resolution -count=1`

The known `TestProofBasedCallAccessGoldenCorpus` failure may be recorded only with exact unchanged-baseline proof. Any new failure blocks A002.

## Residual work and restrictions

- Rerun the complete canonical full build; do not resume or count the interrupted run.
- Run the exact post-build validation above.
- Reinspect scoped diff/semantics and exact changed-file boundary.
- Do not benchmark or capture analyze performance in this Coder slice.
- Do not run detect, stage, commit, push, reset, checkout, stash, cleanup, network/install, or edit ledgers/reports already present.
- Do not open D002-D017.
- If another production owner is required, return `BLOCKED_A002_ARCHITECTURE_BOUNDARY`.

## Handoff status

`CODER_A002_SOURCE_WRITTEN_BUILD_INTERRUPTED_BY_OWNER_READY_FOR_NEXT_CODER`
