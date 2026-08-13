# Child 02 P2-C coder report — affected readers

- Status: `READY_FOR_QA`
- Role: coder implementation, Child 02 P2-C only
- Date: 2026-08-13 (Asia/Bangkok)
- Model slug: `gpt-5`
- Acceptance authority: independent Supervisor; this report records coder implementation and validation evidence only.

## Scope and authority observed

This slice corrected the source-proven affected readers C09, C10, C11, and C16. C12, C13, C14, and C15 were inspected and regression-tested as validate-only rows and were not edited. P2-D, P2-E, closure work, persistence redesign, scanner/target work, relationship identity, repeated-analyze behavior, permissions, approved UI layout/copy/component hierarchy/interaction states, plans, roadmap, ledgers, P2-B reports, skills/rules, and the Git index remained outside this coder lane.

No `anvien detect-changes`, stage, or commit operation was performed because the delegation reserves those operations for Orchestration main after independent Supervisor review.

## Git boundary

### Before implementation

- Repository: `E:\Anvien`
- Branch: `master`
- HEAD: `4d456446fcc49aed0c6d489aa9c63e00d030b53c`
- Accepted baseline message: `fix(persistence): preserve corrected graph facts`
- Staged paths: 0
- Worktree: clean

### At report creation

- Branch: `master`
- HEAD: `4d456446fcc49aed0c6d489aa9c63e00d030b53c` (unchanged)
- Staged paths: 0
- Dirty code/test paths before this report: 8
- This report is the ninth and only report artifact added by the slice.
- No stage or commit was made.

## Authority and source verification

The following were read in full before implementation conclusions were made:

- `AGENTS.md`
- required skills, in delegated order: working-rules, coder, frontend-development, backend-development, ui-be-binding-skill
- debugging skill after the build-artifact lock occurred
- graph-accuracy roadmap and the Child 02 plan, evidence, benchmark, and actual-status files
- `docs/contracts/graph-accuracy-contract.md`
- P2-A inventory Supervisor report and P2-B persistence Supervisor report
- all seven relevant production files and the focused existing tests/data shapes/dispatch paths

Report claims were not used as a substitute for current source or graph verification.

## Current graph, file-detail, and impact evidence

`anvien --help` was run first. `anvien analyze --force` completed before graph-based work. The initial current index represented accepted HEAD `4d45644` with 1,594 files, 96,139 nodes, and 135,260 edges. The root full build later analyzed the dirty production source successfully with 1,594 files scanned, 684 code files parsed, 0 parse failures, 96,236 graph nodes, and 135,314 relationships.

Full `file-detail --json` evidence was collected for:

- `anvien-web/src/hooks/useAppState.local-runtime.tsx`
- `anvien-web/src/components/CodeReferencesPanel.tsx`
- `internal/embeddings/search.go`
- `internal/filecontext/context.go`
- `internal/mcp/context.go`
- `internal/mcp/detect_changes.go`
- `internal/mcp/rename.go`

All were current/non-stale at inspection time. Pre-edit upstream impact evidence was collected for every edited or validate-only owner:

| Row / owner | Symbol impact | File / surrounding blast warning |
|---|---:|---:|
| C09 `resolveNodeIds` | LOW, 0 impacted | runtime hook file CRITICAL: 59 symbols / 28 files / 10 direct |
| C10 `handleNodeGroundingReference` | LOW, 0 impacted | same CRITICAL host-file warning |
| C11 `CodeReferencesPanel` | LOW, 3 impacted / 2 direct | HIGH: 5 symbols / 4 files / 4 direct |
| C16 `SemanticSearch` | LOW, 9 / 3 direct / 2 modules | embeddings file CRITICAL: 53 symbols / 23 files / 27 direct |
| C16 `hydrateSearchResults` | LOW, 5 / 1 / 2 | same CRITICAL host-file warning |
| C16 former `labelFromNodeID` | LOW, 5 / 1 / 2 | removed only after proving no explicit-field consumer needed it |
| C16 `metadataQuery` | LOW, 5 / 1 / 2 | same CRITICAL host-file warning |
| C16 `vectorSearchQuery` | LOW, 9 / 1 / 2 | same CRITICAL host-file warning |
| C16 `chunkRows` | LOW, 9 / 1 / 2 | same CRITICAL host-file warning |
| C16 `DedupBestChunks` | LOW, 10 / 2 / 2 | same CRITICAL host-file warning |
| C16 `ChunkSearchRow` | MEDIUM, 15 / 5 / 2 | same CRITICAL host-file warning |
| C16 `BestChunkMatch` | MEDIUM, 14 / 5 / 2 | same CRITICAL host-file warning |
| C12 `nodeRange` | CRITICAL: 8 / 4 direct / 1 module / 16 processes / 51 linked flows | validate-only; unchanged |
| C13 `contextCandidatePayloads` | CRITICAL: 4 / 4 / 1 / 34 processes | validate-only; unchanged |
| C13 `contextSymbolPayload` | CRITICAL: 1 / 1 / 1 / 7 processes | validate-only; unchanged |
| C14 `detectChangedSymbols` | HIGH: 1 / 1 / 1 / 4 processes | validate-only; unchanged |
| C15 `collectRenameChanges` | CRITICAL: 1 / 1 / 1 / 8 processes | validate-only; unchanged |

HIGH/CRITICAL results were treated as blast-radius warnings. Changes were confined to the exact reader owners and their focused tests.

## Source-level reader mapping

| Row | P2-B persisted fact consumed | P2-C reader behavior |
|---|---|---|
| C09 | exact persisted graph node ID | `resolveNodeIds` now accepts exact graph-set membership only. Graph absence or a suffix/fragment returns no match; the old first suffix match is gone. |
| C10 | persisted label, name, path, one-based/exclusive range, and byte columns | grounding filters all nodes by label+name and proceeds only for exactly one match. Ambiguous same-name declarations fail closed. Numeric fields are narrowed before the one-based/exclusive range and zero-based UTF-8 byte columns are stored. |
| C11 | one-based lines, zero-based UTF-8 byte columns, exclusive ends | the panel preserves graph coordinates, converts once at the zero-based-inclusive `/api/file` read boundary, displays one-based lines, scrolls to the same one-based line, and excludes an `endLine` whose `endCol` is zero from highlight/display. Approved JSX hierarchy, copy, styles, and interaction states were not redesigned. |
| C12 | full persisted range | `nodeRange` already copies all four coordinates directly. Inspected, unchanged, and covered by `internal/filecontext` regression. |
| C13 | UID/path/range and ambiguity facts | payload builders already expose UID, path, one-based lines, and ambiguous candidates. Inspected, unchanged, and covered by `internal/mcp` regression. |
| C14 | persisted start/end lines | changed-symbol detection already intersects the persisted line range directly. Inspected, unchanged, and covered by `internal/mcp` regression. |
| C15 | persisted UID and range | rename collection already resolves by UID and converts a one-based line only at the file-array edit boundary. Inspected, unchanged, and covered by `internal/mcp` regression. |
| C16 | explicit persisted `CodeEmbedding.label` | label is selected by the vector query, read into `ChunkSearchRow`, preserved by nearest-chunk dedup into `BestChunkMatch`, and used to select the metadata table. Missing label and metadata query failures now return clear errors. `labelFromNodeID` was removed from active code. |

Coder evidence denominator: **8/8 affected-reader rows implemented or validate-only evidenced**. This is a coverage count, not an acceptance verdict.

## Changed files and containment rationale

Production:

1. `anvien-web/src/hooks/useAppState.local-runtime.tsx` — C09 exact-ID/fail-closed resolution and C10 unambiguous grounding with corrected persisted coordinates.
2. `anvien-web/src/components/CodeReferencesPanel.tsx` — C11 consistent one-based/exclusive display, highlight, scroll, and API slicing conversion.
3. `internal/embeddings/search.go` — C16 explicit-label propagation and fail-clear hydration.

Focused tests:

4. `anvien-web/test/unit/useAppState.local-runtime.test.tsx` — exact ID, suffix-negative, ambiguous same-name, and corrected coordinate coverage.
5. `anvien-web/test/unit/ChatPanel.grounding-links.test.tsx` — dispatch-level one-based grounding expectations.
6. `anvien-web/test/unit/CodeReferencesPanel.graph-health.test.tsx` — read/display/highlight/scroll proof for one-based lines and exclusive terminal line.
7. `internal/embeddings/search_test.go` — vector row label propagation, opaque IDs, same-name cross-label hydration, dedup, missing-label negative, and metadata-failure negative.
8. `internal/embeddings/search_ladybugdb_test.go` — build-tagged native Ladybug persistence/readback plus public `SemanticSearch` hydration for opaque identities and same-name declarations.

No C12-C15 source file or forbidden plan/roadmap/ledger/report/skill/rule file changed. `git diff --check` was clean.

## Production-first chronology

1. Read authority, skills, source, tests, contracts, and real dispatch/data flows.
2. Refreshed graph and collected file-detail/upstream-impact evidence.
3. Implemented C09/C10/C11/C16 production behavior.
4. Ran the root full build. The first meaningful retry revealed a binary lock; after resource-holder cleanup the next build exposed TypeScript narrowing errors. Production code was corrected to narrow unknown graph properties to numbers.
5. Repeated verified-holder/exclusive-open checks and ran the root full build successfully.
6. Only after production behavior was correct, updated/added focused frontend and backend tests.
7. Ran focused package and native-boundary validations, validate-only regressions, built runtime mount, and containment audit.

Test-only assertion corrections during validation were limited to harness accuracy: split citation text, JSDOM color serialization, and a double-`requestAnimationFrame` mock. They did not change production behavior.

## Clean-build holder handling and full build

Owner's latest clean-build authority was applied. When `anvien/bin/anvien.exe` was locked, Windows Restart Manager identified the exact artifact holders as `anvien.exe` PIDs 16284, 19416, 15648, and 13880, with launcher `cmd.exe` parents 18320, 13908, 16996, and 21328. These verified holders/launchers only were stopped. Exclusive-open/resource cleanliness was then confirmed before retry. No unrelated process was stopped.

Final root gate:

```text
cd E:\Anvien
npm run full-build
```

Result: exit 0 in 129.1 seconds. It completed packaged runtime, global CLI 1.2.8, launcher, `tsc -b`, Vite production build, and built-source analyze. Only pre-existing Vite chunk/dynamic-import warnings remained.

No subsequent full-build retry was necessary after test-only additions. If main or QA rebuilds, it must repeat exact holder identification, verified-holder/launcher cleanup, exclusive-open confirmation, then build.

## Focused validation results

### Frontend

```text
cd E:\Anvien\anvien-web
npm test -- --run test/unit/useAppState.local-runtime.test.tsx test/unit/ChatPanel.grounding-links.test.tsx test/unit/CodeReferencesPanel.graph-health.test.tsx
```

Result: 3 test files, 11/11 tests completed successfully; final run duration 11.80 seconds.

The focused evidence covers:

- exact opaque ID accepted and suffix fragment rejected;
- ambiguous same-label/same-name grounding adds no arbitrary citation and does not open the code panel;
- file and node grounding preserve one-based ranges;
- graph range 10:0 to 12:0 reads through the API conversion, displays `L10–11`, highlights lines 10–11, excludes line 12, and scrolls to the corresponding one-based line.

### Backend C16 package boundary

```text
go test ./internal/embeddings -count=1
```

Final result: exit 0, package completed successfully in 3.708 seconds.

Evidence includes a production `SemanticSearch` call with opaque IDs, nearest-chunk dedup, explicit `emb.label AS label`, distinct `Function`/`Method` hydration despite the same name, missing-label rejection without an ID-derived metadata query, metadata-query failure propagation, dimension mismatch, and fetch-window regressions.

### Native Ladybug boundary

The tagged test used the packaged Ladybug native library under repo-local `.tmp/ladybug-native/v0.19.1/windows-x86_64`:

```text
go test -tags ladybugdb ./internal/embeddings -run '^TestNativeLadybugSemanticSearchHydratesPersistedExplicitLabels$' -count=1 -v
```

Result: exit 0; test completed successfully in 0.53 seconds (package 2.703 seconds).

The test persists and reads real `CodeEmbedding.label` values through native Ladybug, then invokes public production `SemanticSearch`. It verifies three opaque identities: two `Function` declarations with the same `sharedName` in the same table and different paths/ranges, plus a `Method` with the same name. All three hydrate to their own ID, label, path, and range. The test seams only the vector-index row return from the rows just read from native Ladybug; production still issues the vector query, and all row propagation/dedup/native metadata hydration is production code.

The initially attempted fully native vector-index test could not install/load the required vector extension because Ladybug requested `https://extension.ladybugdb.com/v0.19.0/win_amd64/vector/libvector.lbug_extension` and the official endpoint returned HTTP 522. A direct download attempt returned the same 522. No incompatible cached 0.15/0.16 extension was substituted. This is a residual environment surface for QA/Supervisor, not counted as native vector-index evidence.

### Validate-only C12-C15 regressions

```text
go test ./internal/filecontext -count=1
```

Result: exit 0, completed successfully (final recorded focused run 1.894 seconds).

```text
go test ./internal/mcp -count=1
```

Result: exit 0, completed successfully (final recorded focused run 13.224 seconds).

Git path checks confirmed all four validate-only source files unchanged:

- `internal/filecontext/context.go`
- `internal/mcp/context.go`
- `internal/mcp/detect_changes.go`
- `internal/mcp/rename.go`

## Built user-visible runtime mount and QA handoff

The already-built production surfaces were mounted, not a development server:

- backend: `E:\Anvien\anvien\bin\anvien.exe serve --host 127.0.0.1 --port 4848`
- frontend: `npm run preview` in `E:\Anvien\anvien-web`, strict port 5228

BrowserPanel opened `http://127.0.0.1:5228/`. The first visible DOM contained the real `Anvien` start screen and `Start Anvien`; console warning/error logs were empty. Clicking `Start Anvien` completed the backend/frontend handshake and mounted the real repository-selection screen. The Anvien repository card displayed the current built graph inventory: 1,594 files, 96,236 symbols, and 662 flows. Console warning/error logs remained empty.

The repository card was not clicked because that visible action states it will rebuild the graph; another graph mutation was unnecessary for coder mount proof. No official Playwright verdict/artifact was manufactured. Both runtime listeners were verified as this session's processes before cleanup (`anvien.exe` PID 11740 and the Anvien Vite preview `node.exe` PID 16776), then stopped; final listener counts on ports 4848 and 5228 were both zero.

Exact independent QA target and steps:

1. Repeat verified-holder/exclusive-open handling if a build is required; mount built backend on 4848 and built Vite preview on 5228.
2. Open `http://127.0.0.1:5228/`, start Anvien, and select the already-indexed Anvien repository using the QA lane's controlled fixture/event setup.
3. C09: send a tool result containing one complete opaque node ID and one suffix-only fragment. Verify only the complete ID focuses/highlights; the suffix selects nothing.
4. C10: exercise a grounding link whose label+name maps to two distinct same-name nodes. Verify no arbitrary citation/panel opening occurs; exercise a unique label+name and verify its persisted path/range/columns are used.
5. C11: select a graph symbol with one-based range 10:0 to 12:0. Verify citation copy shows lines 10–11, line 10 scroll is targeted, highlights cover 10–11 only, line 12 is excluded, and the file request converts exactly once at the API boundary.
6. Preserve existing layout, copy, hierarchy, and interaction states. Any official script must be reusable under `playwright/`; official evidence must be paired JSON and Markdown under `Reports/qa/playwright/...`.

## Failures, corrections, and artifact lifecycle

- An initial graph-analyze command outlived the terminal wrapper timeout; the process completed and current index status/counts were verified before using its evidence.
- A full-build retry exposed a verified `anvien.exe` artifact lock. Exact holders and launchers were identified, stopped under Owner authority, and exclusive-open was confirmed before retry.
- The clean build then exposed real TypeScript `unknown`-to-number errors; production numeric narrowing corrected them, and the clean full build completed.
- The initial native vector-index integration attempt failed while downloading Ladybug vector extension; direct endpoint retry returned HTTP 522. The test was narrowed truthfully to native persistence/readback plus public `SemanticSearch` and native metadata hydration, with the vector row seam explicitly recorded above.
- Runtime/debug logs and the failed extension-cache directory were created only under repository-local `.tmp`. Shell policy rejected automated removal commands; they are ignored debug artifacts, do not appear in the Git boundary, and contain no official QA evidence. Runtime processes were stopped and ports were clean.

## Residual unverified surfaces and handoff decision

- Independent browser QA of the exact C09-C11 visible flows remains required; coder evidence proves the built runtime mounts and provides exact targets/steps.
- A fully native Ladybug vector-index invocation remains unverified because the required official 0.19.0 vector extension endpoint returned HTTP 522. Package-level `SemanticSearch` plus native persistence/readback/metadata hydration covers the P2-C label contract without substituting an incompatible extension.
- No official QA JSON/Markdown artifact was produced in the coder lane.
- No acceptance verdict is asserted here.

Within the exact P2-C boundary, production implementation and coder-level validation are complete enough to hand off as `READY_FOR_QA` to Orchestration main for an independent QA lane and subsequent Supervisor review.
