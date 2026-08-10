# Coder Report - Child 01 P1-B Range, Selection Range, and Identity Inputs

## Review Status

- Status: `ACCEPTED / DETECT PASS / ISOLATED COMMIT BOUNDARY`
- Scope: `Child 01 / P1-B` only.
- Git boundary: base commit `8cbf6006d53836494d5cc862baea77939af9f938`; the accepted candidate is committed as one isolated P1-B boundary immediately after the final ledger update, with the resulting hash reported directly from Git.
- Authority: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md` and `docs/contracts/graph-accuracy-contract.md`.

## Invariant Family Map

- Family: provider-backed definition position and existing lexical-owner/meaning inputs.
- Source truth / owners: `scopeir.DefinitionFact`, TSJS `nodeRange`, TSJS `collector.addDefinition`, and the already-existing `ScopeFact.OwnedDefIDs`/bindings, `DefinitionFact.Label`, and `DefinitionFact.OwnerID` inputs.
- Primary path: UTF-8 source bytes -> tree-sitter TSJS AST -> `collector.addDefinition` -> normalized ScopeIR -> deterministic JSON round-trip.
- Sibling surfaces checked: TSJS scope construction/containment, ScopeIR normalization/sorting, parser UTF-8 input, other provider facts, the Child 03 binding-pattern branch, identity construction, graph mutation, projection, persistence, and readers.
- Forbidden legacy/alternate path: no range-only identity, no guessed UTF-16 conversion, no shared inclusive/exclusive rewrite, no redundant lexical/meaning topology, no binding-pattern extraction, and no P1-C/P1-D/persistence change.
- Verify matrix: construct/selection range; LF/CRLF; Unicode byte columns; tab; exclusive end point; nested scopes; exact scope binding; same lexeme across provider labels; member owner ID; JSON round-trip; binding-pattern preservation; product-package regression.

## Implementation

- `internal/scopeir/facts.go`: added additive optional `SelectionRange *Range` with `json:"selectionRange,omitempty"` to the shared definition fact.
- `internal/providers/tsjs/nodes.go`: documented the source-proven TSJS contract: one-based lines, zero-based UTF-8 byte columns, exclusive end point.
- `internal/providers/tsjs/definitions.go`: retained the existing construct range and added the declaring `nameNode` range to `SelectionRange` in the exact `addDefinition` path.
- `internal/providers/tsjs/definition_position_inputs_test.go`: added reusable provider/ScopeIR boundary coverage.
- Child ledgers: recorded impact, source, build, runtime, QA, status refresh, and the non-benchmarkable classification.

No production edit was made to `Range`, `PositionIndex`, TSJS scope containment, other providers, identity construction, collision handling, projection, persistence, readers, or the target repository.

## Blast Radius

- `DefinitionFact`: CRITICAL, `586` symbols / `80` files / `23` modules / `63` processes.
- `Range`: CRITICAL, `854 / 143 / 22 / 75`; it was inspected and left unchanged.
- TSJS `nodeRange`: MEDIUM, `8 / 4 / 1 / 0`.
- TSJS `collector.addDefinition`: LOW, `0` upstream symbols / `1` selected file / `0` modules / `0` processes.
- Final test file: file risk LOW; edited test functions are LOW with zero upstream impact, and the lexical-owner helper is LOW with one caller in one file and zero processes.

CRITICAL/HIGH values were treated as scope warnings. The change remained within the P1-B owner set.

## Tests and Runtime Evidence

- `internal/providers/tsjs/definition_position_inputs_test.go:12` — `TestExtractDefinitionPositionInputs`: PASS; proves full versus selection ranges, LF/CRLF stability, Unicode byte columns, tab handling, exclusive endpoint slicing, and deterministic JSON round-trip.
- `internal/providers/tsjs/definition_position_inputs_test.go:70` — `TestExtractDefinitionOwnerMeaningAndLexicalInputs`: PASS; proves exact-one nested lexical owners, parent/binding linkage, member `OwnerID`, same `Shared` lexeme in interface/value labels, selection tokens, and JSON round-trip.
- `internal/providers/tsjs/definition_position_inputs_test.go:165` — `TestP1BPreservesBindingPatternBoundary`: PASS; proves P1-B did not take Child 03 binding-pattern ownership.

Commands and results:

- Full build: `powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\full-build.ps1` — PASS after `npm ci` restored the committed `anvien-web` dependencies. It rebuilt the packaged Go runtime, web, launcher/server, reported CLI `1.2.8`, and completed built-CLI analyze with `1,560` scanned, `677` parsed-code, `0` failed, `85,193` nodes, and `124,090` relationships.
- Build transparency: the first attempt failed because `anvien-web/node_modules` was absent. A later retry after test-only assertion strengthening was blocked before compile by three editor-owned `anvien mcp` processes locking the global-junction runtime binary; `anvien doctor processes --json` says those integrations are expected to stay running, so none was terminated. Production/build inputs did not change after the successful full build.
- Built boundary: Go 1.26.3 compiled `.tmp/p1b-tsjs-boundary.test.exe`; after correcting one PowerShell dotted-flag invocation, the final executable run of all three tests PASSed.
- Focused QA: `go test ./internal/providers/tsjs -run '^Test(ExtractDefinitionPositionInputs|ExtractDefinitionOwnerMeaningAndLexicalInputs|P1BPreservesBindingPatternBoundary)$' -count=1 -v` — PASS.
- Owner regression: `go test ./internal/providers/tsjs ./internal/scopeir ./internal/parser -count=1` — PASS.
- Raw repository regression: `go test ./... -count=1` — exit `1` only for five intentionally non-standalone/compile-negative corpus packages under `anvien/test/fixtures`; all reported `internal/...` packages PASSed.
- Product regression: `go list -e ./...` excluding only `*/anvien/test/fixtures/*`, then `go test` — PASS, `61/61` packages.

## E2E Verification

```text
E2E Verification:
  [PASS] Compiled: repository full-build completed against the P1-B production diff
  [PASS] Runtime: compiled TSJS boundary executable exercised parser -> provider -> ScopeIR/JSON
  [PASS] Happy path: construct/name nodes -> full/selection ranges with lexical/meaning inputs preserved
  [PASS] Edge case: CRLF, Unicode, tab, nested duplicate name, same lexeme across labels, and binding-pattern preservation
```

## Sibling and Legacy Status

- Sibling surfaces checked: `Range`, `PositionIndex`, TSJS scopes, normalization/sort, other providers, identity, collision, projection, persistence/readers, and the target boundary.
- Legacy fallback status: none added; existing PositionIndex and TSJS scope containment semantics are unchanged.
- Stale tests/helpers/plans updated: new P1-B boundary test and Child 01 live ledgers only.
- Debug artifact lifecycle: `.tmp/p1b-tsjs-boundary.test.exe` and superseded `.tmp/p1b-anvien` traces were removed after Supervisor PASS; unrelated pre-existing `.tmp` content was preserved.
- Residual unverified surfaces: `none` inside P1-B. Identity consumption remains intentionally closed in P1-C; collision remains P1-D; target `4/4` remains P1-E.

## Risks / Open Points

- `SelectionRange` is optional so unrelated providers preserve their existing output until their own source-backed scope opens.
- The raw `go test ./...` command includes deliberately invalid fixture packages; the exact product-package regression passes and the raw failure is not hidden.
- Editor-owned MCP processes can lock the globally-junctioned runtime binary on subsequent build retries; no editor runtime was killed for this slice.

## Supervisor History

- First review: `reports/Supervisor/rp_supervisor_260810_115437_by_gpt-5-codex_child01_p1b.md` — `REJECT`; technical boundary cleared, but two Detailed Findings paragraphs retained the pre-P1-B state.
- Correction: actual-status narrative only; no production/test change and no later slice opened.
- History-closure review: `reports/Supervisor/rp_supervisor_260810_120042_by_gpt-5-codex_child01_p1b_resubmission.md` — `PASS`; detect-changes and the isolated commit were intentionally pending at review time and are recorded by the final closure rows below.

## Change Detection

- Mandatory fresh graph: `anvien analyze --force --json` — PASS, `1,564` scanned / `677` parsed-code / `0` failed / `85,247` nodes / `124,149` relationships.
- Mandatory pre-commit check: `anvien detect-changes --repo E:\Anvien --scope all --json` — PASS (exit `0`), overall risk `low`, `50` changed symbols in `7` changed files, `0` affected symbols, `6` affected-file summaries, and no affected process.
- File-layer risk remains `high` at the already-reviewed TSJS provider owner. Resolution health remains non-degraded: `0` changed source nodes with gaps, `0` total resolution-gap count, `0` degraded nodes, and `0` nodes with gaps.
- The repo-local raw detect trace was used only to verify and record these exact counts, then removed as debug evidence after ledger recording.

## Evidence

- `E1-P1B-IMPACT1`
- `E1-P1B-SOURCE1`
- `E1-P1B-BUILD1`
- `E1-P1B-TEST1`
- `E1-P1B-RUNTIME1`
- `E1-P1B-REVIEW1`
- `E1-P1B-DETECT1`
- `E1-P1B-COMMIT1`
