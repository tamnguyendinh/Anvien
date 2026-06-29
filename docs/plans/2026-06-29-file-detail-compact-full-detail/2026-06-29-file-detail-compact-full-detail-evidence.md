# File Detail Compact Full Detail Evidence Ledger

## Metadata

- Date: `2026-06-29`
- Plan: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-plan.md`
- Evidence: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-evidence.md`
- Benchmark: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-benchmark.md`
- Actual status: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-actual-status.md`

## Evidence Rules

The evidence file explains why the work is known to be correct.

It should contain:

- metadata and companion files;
- evidence rules or evidence template;
- evidence sections such as `E0`, `E1`, or sections by phase/task;
- user report or problem evidence;
- source inspection, codebase facts, and document facts;
- commands run and pass/fail result;
- impact or blast-radius evidence when code/graph behavior changes;
- implementation evidence: files changed and behavior changed;
- validation evidence: build, tests, e2e, screenshots, or traces;
- failures encountered and how they were handled;
- detect-changes before commit;
- commit hash and closure evidence.

Evidence can reference short metric traces, but long metric tables belong in the benchmark file.

### Evidence ID Naming

Use stable, phase-scoped evidence IDs so `plan.md`, `actual-status.md`, `benchmark.md`, and later agents can reference exact proof without ambiguity.

Format:

```text
E<phase>-<item>-<kind><n>
```

Rules:

- `E<phase>` matches the plan phase number: `E0` for `P0`, `E1` for `P1`, `E2` for `P2`, and so on.
- `<item>` matches the checklist item without the dash: `P0A`, `P1A`, `P2B`.
- `<kind>` is plan-local. Choose a short uppercase token that is meaningful for this repo and this plan.
- `<n>` is a 1-based sequence number within that phase item and kind.
- Keep the same `<kind>` meaning stable inside one plan.
- Do not reuse an evidence ID for different facts.
- Reference exact evidence IDs from `actual-status.md` and `benchmark.md`; avoid referencing only broad section IDs such as `E1`.
- Use ranges such as `E0-P0A-FD1..E0-P0A-FD17` only for compact inventory summaries; use exact IDs when a specific status decision depends on a specific fact.
- If nearby plans already use a clear local evidence naming style, follow that style instead of inventing a new one.

Examples only:

- `E0-P0A-SRC1`
- `E0-P0A-GRAPH1`
- `E1-P1A-ROUTE1`
- `E2-P2B-KEYBOARD1`
- `E2-P2B-DETECT1`

Evidence sections must follow the plan phases:

- `E0` corresponds to `P0`.
- `E1` corresponds to `P1`.
- `E2` corresponds to `P2`.
- Use exact evidence IDs inside each section, not broad section IDs as proof.
- Each evidence section must name the plan phase or checklist item it supports.
- Do not invent fixed evidence categories; record the evidence required by the matching plan phase.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-DISCUSS1`: User clarified and accepted that `file-detail` must stay a full-detail surface, not a summary or `impact` replacement. Accepted direction: compact the representation losslessly, do not cut data, and add richer related-file metadata.
- `E0-P0A-GRAPH1`: `anvien analyze --force` completed for `E:\Anvien`; output reported `files: scanned=1447 parsed_code=674 failed=0`, `nodes=83186`, `relationships=121455`, graph path `E:\Anvien\.anvien\graph.json`, and fileProjection built with `files=1447 dependencyEdges=16510 unresolved=420 hotspots=5`.
- `E0-P0A-WT1`: `git status --short` showed one unrelated untracked file: `internal/aicontext/skills/Spec-to-SVG-Flow-Map/spec-to-svg-flow-map.vi.md`. This plan must not touch it.
- `E0-P0A-QUERY1`: `anvien query "file-detail command file context compact normalized related files" --repo Anvien` identified `internal/filecontext/context.go` as the file-detail builder owner and `anvien-web/src/components/FileDetailPanel.tsx` as the Web display owner. Query output also showed relationship hints for `internal/filecontext/context.go`: local `565`, inbound `252`, outbound `88`, linked flows `53`, linked tests `8`, unresolved `542`, risk `high`.
- `E0-P0A-SRC1`: Source search found the current top-level file-detail model in `internal/filecontext/context.go`: `FileContext` contains `summary`, `symbolTree`, `relationships`, `unresolved`, `linked`, `quality`, and `limits`; `RelationshipSample` is an expanded object; `BuildFileContext` builds expanded sections directly.
- `E0-P0A-SRC2`: Source search found CLI ownership in `internal/cli/file_detail_command.go`: `newFileDetailCommand` currently has `--json`, `--relationships`, `--unresolved`, and `--linked` flags, and writes `FileContext` through `writeJSON`.
- `E0-P0A-SRC3`: Source search found HTTP ownership in `internal/httpapi/file_context.go`: `/api/file-detail` calls `BuildFileContext` and writes the expanded context as JSON with bounded sample query params.
- `E0-P0A-SRC4`: Source search found contract ownership in `internal/contracts/web_ui.go`: route `/api/file-detail` currently declares response type `FileContextResponse`; generated TypeScript currently exposes `FileContextResponse` with `symbolTree`, `relationships`, `unresolved`, `linked`, and `limits`.
- `E0-P0A-SRC5`: Source search found Web consumer ownership in `anvien-web/src/services/backend-client.ts` and `anvien-web/src/components/FileDetailPanel.tsx`; the panel currently fetches file detail with sample limits and renders major sections using expanded generated types.
- `E0-P0A-FD1`: `anvien file-detail internal/filecontext/context.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `253168` compacted-output characters from the current expanded shape, `429` symbols, `565` local relationships, `88` outbound, `252` inbound, `42` unique related files, `542` unresolved, `88` unresolved groups, `53` linked flows, `8` linked tests, risk `high`, stale `false`.
- `E0-P0A-FD2`: `anvien file-detail internal/cli/file_detail_command.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `55635` characters, `78` symbols, `54` local, `67` outbound, `9` inbound, `19` unique related files, `198` unresolved, risk `high`, stale `false`.
- `E0-P0A-FD3`: `anvien file-detail internal/httpapi/file_context.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `54237` characters, `71` symbols, `55` local, `38` outbound, `10` inbound, `20` unique related files, `175` unresolved, risk `high`, stale `false`.
- `E0-P0A-FD4`: `anvien file-detail anvien-web/src/components/FileDetailPanel.tsx --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `34089` characters, `70` symbols, `20` local, `5` outbound, `4` inbound, `4` unique related files, `298` unresolved, risk `high`, stale `false`.
- `E0-P0A-FD5`: `anvien file-detail anvien-web/src/services/backend-client.ts --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `98945` characters, `197` symbols, `132` local, `9` outbound, `177` inbound, `22` unique related files, `392` unresolved, risk `high`, stale `false`.
- `E0-P0A-FD6`: `anvien file-detail internal/contracts/web_ui.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `104671` characters, `162` symbols, `160` local, `59` outbound, `38` inbound, `31` unique related files, `636` unresolved, risk `high`, stale `false`.
- `E0-P0A-FD7`: `anvien file-detail contracts/web-ui/anvien-web-contract.schema.json --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `1908` characters, no symbols or relationships, risk `low`, stale `false`; this is generated output and must not be edited as source of truth.
- `E0-P0A-IMPACT1`: `anvien impact file internal/filecontext/context.go --repo Anvien --direction upstream` showed wide blast radius: 30 affected files including CLI, HTTP/API, MCP, contracts, Web file detail, and graph-health consumers; output included CRITICAL symbol-risk entries in file-detail structs and fields.
- `E0-P0A-IMPACT2`: `anvien impact file internal/cli/file_detail_command.go --repo Anvien --direction upstream` showed 5 affected files and one CLI process (`NewFileDetailCommand -> FileProjectionGraphInfo`); output included CRITICAL symbol-risk entries for CLI helpers including JSON rendering.
- `E0-P0A-IMPACT3`: `anvien impact file internal/httpapi/file_context.go --repo Anvien --direction upstream` showed 3 affected files and one API process (`HandleFileHotspots -> CollectGitPathList`); output included HIGH/CRITICAL symbol-risk entries for endpoint helpers.
- `E0-P0A-IMPACT4`: `anvien impact file anvien-web/src/components/FileDetailPanel.tsx --repo Anvien --direction upstream` showed 4 affected files (`App.tsx`, `CodeReferencesPanel.tsx`, `FileDetailPanel.tsx`, `main.tsx`) and no affected processes; Web component impact is narrower than backend shape impact.
- `E0-P0A-REVIEW1`: Supervisor report `reports/Supervisor/rp_supervisor_260629_154304_by_gpt-5-codex_file-detail-plan-readiness.md` rejected implementation readiness until the plan added MCP/agent surface coverage, explicit compact full-detail limit/default semantics, and README/RUNBOOK docs scope.
- `E0-P0A-GRAPH2`: `anvien analyze --force` was rerun before updating this plan; output reported `files: scanned=1452 parsed_code=674 failed=0`, `nodes=83252`, `relationships=121521`, graph path `E:\Anvien\.anvien\graph.json`, and fileProjection built with `files=1452 dependencyEdges=16510 unresolved=420 hotspots=5`.
- `E0-P0A-SRC6`: Source inspection found MCP file-context consumers: `internal/mcp/target_dispatch.go` calls `BuildFileContext` through `mcpBuildRepoFileContext`, `internal/mcp/context.go` returns `fileContext` for context file payloads, and `internal/mcp/impact.go` uses the same file-context path for file impact flows.
- `E0-P0A-SRC7`: Source inspection found user-facing command docs for this scope in `README.md` and `RUNBOOK.md`: examples document `anvien context file`, `anvien file-detail`, and `/api/file-detail` requests.
- `E0-P0A-FD8`: `anvien file-detail internal/mcp/target_dispatch.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `40602` compacted-output characters from the current expanded shape, `8` local relationships, `79` outbound, `51` inbound, `21` unique related files, and `105` unresolved sites.
- `E0-P0A-FD9`: `anvien file-detail internal/mcp/context.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `55736` compacted-output characters from the current expanded shape, `50` local relationships, `111` outbound, `48` inbound, `28` unique related files, and `261` unresolved sites.
- `E0-P0A-FD10`: `anvien file-detail internal/mcp/impact.go --repo Anvien --json --relationships 1 --unresolved 1 --linked 1` reported `83000` compacted-output characters from the current expanded shape, `115` local relationships, `128` outbound, `47` inbound, `40` unique related files, and `497` unresolved sites.
- `E0-P0A-IMPACT5`: `anvien impact file internal/mcp/target_dispatch.go --repo Anvien --direction upstream` showed `9` affected files, `1` affected process (`ContextToolInternal -> ContextCandidate`), and CRITICAL symbol-risk entries across the MCP file-context/dispatch helpers.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`, `P1-B`, `P1-C`

- `E1-P1A-IMPACT1`: Before editing, `anvien analyze --force` was rerun and `anvien impact file internal/filecontext/context.go --repo Anvien --direction upstream` plus symbol impacts for `FileContext`, `BuildFileContext`, and `RelationshipSample` were checked. Evidence showed high/critical blast radius across CLI, HTTP/API, MCP, contracts, Web, and filecontext relationship processes. P1-A stayed scoped to a new compact model/converter and did not change existing `FileContext` JSON fields.
- `E1-P1A-SRC1`: Added `internal/filecontext/compact.go` with `CompactFileContext`, explicit `format`/`version`, schema row definitions, dictionaries for repeated files/symbols/source-sites/kinds/statuses, tuple ranges, compact row tables for symbols/relationships/unresolved/linked data, and `CompactFileContextFromExpanded` conversion from existing expanded `FileContext`.
- `E1-P1A-TEST1`: `go test ./internal/filecontext -count=1` passed after adding `TestCompactFileContextFromExpandedInternsRowsAndPreservesCounts`. The test proves compact conversion preserves top-level metadata, summary, quality, limits, relationship counts, linked counts, symbol rows, relationship rows, unresolved rows, dictionary references, source-site interning, JSON marshalability, and omitted-row metadata for sample-limited expanded input.
- `E1-P1A-BUILD1`: `go build ./...` was attempted and failed only because repository fixtures under `anvien/test/fixtures` include intentionally non-buildable sample packages/C files. The usable repo build boundary `go build ./cmd/... ./internal/...` passed, then `go test ./internal/filecontext -count=1` passed again.
- `E1-P1B-IMPACT1`: Before P1-B edits, `anvien analyze --force` was rerun after P1-A commit. `anvien impact file internal/filecontext/compact.go --repo Anvien --direction upstream` showed affected CLI/HTTP/MCP files through package-level relationships and one compact process (`CompactFileContextFromExpanded -> CompactContextBuilder`); `anvien impact symbol CompactFileContextFromExpanded --repo Anvien --direction upstream` showed LOW direct symbol blast radius. P1-B stayed within `internal/filecontext/compact.go` and tests.
- `E1-P1B-SRC1`: Added builder-aware compact entrypoints `BuildCompactFileContext` and `(*Builder).CompactFileContextFromExpanded`. Compact tables now include `relatedFiles` rows with file dictionary ref, language, kind, fileRole, fileGroup, appLayer, functionalArea, parseStatus, symbolCount, unresolved count, risk, outbound/inbound/local flags, relationship total, and relationship-kind counts.
- `E1-P1B-TEST1`: `go test ./internal/filecontext -count=1` passed after adding `TestBuildCompactFileContextAddsRelatedFileMetadata`. The test proves related-file metadata is present for outbound `src/store.go` and inbound `src/app_test.go`, with graph-derived kind/app layer, symbol counts, risk, direction flags, totals, and `CALLS` counts.
- `E1-P1B-BUILD1`: `go build ./cmd/... ./internal/...` passed after P1-B changes.
- `E1-P1C-IMPACT1`: Before P1-C edits, `anvien analyze --force` was rerun. `anvien impact file internal/filecontext/context.go --repo Anvien --direction upstream` showed high/critical blast radius across CLI, HTTP/API, MCP, contracts, Web, and filecontext consumers. `anvien impact symbol normalizeLimits --repo Anvien --direction upstream` showed affected files `internal/cli/resolution_inventory_command.go`, `internal/filecontext/compact.go`, and `internal/filecontext/context.go`; `anvien impact file internal/filecontext/compact.go --repo Anvien --direction upstream` showed package-level CLI/HTTP/MCP consumers; `anvien impact symbol BuildCompactFileContext --repo Anvien --direction upstream` showed LOW direct symbol blast radius. P1-C stayed scoped to compact/default limit semantics and tests.
- `E1-P1C-SRC1`: Added `FullDetailSampleLimit = -1`, preserved expanded `Options{}` defaults by making `normalizeLimits` substitute `defaultSampleLimit` only for zero values, and made `BuildCompactFileContext` map unspecified compact options to the full-detail sentinel before building expanded facts for compact conversion.
- `E1-P1C-TEST1`: `go test ./internal/filecontext -count=1` passed after adding `TestBuildCompactFileContextDefaultsToFullRows` and `TestBuildCompactFileContextExplicitLimitsExposeOmittedRows`. The tests prove compact default output returns all available relationship and unresolved rows in fixtures while expanded `BuildFileContext(..., Options{})` keeps the existing `defaultSampleLimit`.
- `E1-P1C-LIMIT1`: Limited compact test proves explicit limit metadata is visible: unresolved rows report total/returned/omitted `3/1/2`, linked flow rows report `2/1/1`, and an unspecified relationship limit remains the full-detail sentinel instead of silently inheriting the expanded default.
- `E1-P1C-BUILD1`: `go build ./cmd/... ./internal/...` passed after P1-C changes.
- `E1-P1C-MEASURE1`: A repo-local temporary helper under `.tmp/` decoded `.anvien/graph.json`, built expanded full-detail and compact full-detail payloads for `internal/filecontext/context.go`, recorded the P1 benchmark measurements, and was removed before commit. The measurement reported expanded full JSON `849155` characters, compact full JSON `281704` characters, `429` compact symbol rows, `43` related-file rows, `986` compact relationship rows, `542` unresolved rows, `56` linked flow rows, and compact default limits `-1/-1/-1`.
- `E1-P1C-DETECT1`: `anvien detect-changes --repo Anvien --scope all` ran before the P1-C commit. It reported `7` changed files, changed-file risk `high`, overall risk `medium`, affected process `BuildFileContext -> Limits`, changed symbols including `FullDetailSampleLimit`, `normalizeLimits`, `BuildCompactFileContext`, `compactDefaultOptions`, `TestBuildCompactFileContextDefaultsToFullRows`, and `TestBuildCompactFileContextExplicitLimitsExposeOmittedRows`, plus analyzer-only resolution-gap changes in `compactDefaultOptions`.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`, `P2-B`, `P2-C`, `P2-D`, `P2-E`

- `E2-P2A-IMPACT1`: Before P2-A edits, `anvien analyze --force` was rerun. `anvien impact file internal/cli/file_detail_command.go --repo Anvien --direction upstream` showed 5 affected files and process `NewFileDetailCommand -> FileProjectionGraphInfo`; `anvien impact symbol newFileDetailCommand --repo Anvien --direction upstream` and `anvien impact symbol renderFileContext --repo Anvien --direction upstream` showed CLI launcher/main command blast radius. `anvien impact symbol writeJSON --repo Anvien --direction upstream` was ambiguous across CLI and HTTP helpers, so P2-A did not edit the existing helper and added a compact-only CLI JSON writer instead.
- `E2-P2A-SRC1`: `anvien file-detail` now accepts `--format compact|expanded` for JSON output. `--json` defaults to compact full-detail output through `BuildCompactFileContext`; `--json --format expanded` preserves the legacy expanded `FileContext` shape. Compact JSON uses compact encoding rather than pretty JSON, while expanded JSON still uses the existing pretty writer. `--format` without `--json` fails visibly, and explicit `--relationships`, `--unresolved`, and `--linked` flags are passed to compact output only when the caller supplies that flag.
- `E2-P2A-TEST1`: `go test ./internal/cli -run 'TestFileDetailCommand|TestDirectToolHelpShowsCompatibilityFlags' -count=1` passed, then `go test ./internal/cli -count=1` passed. Tests cover compact default JSON, absolute path compact JSON, explicit expanded JSON, explicit compact relationship limit with total/returned/omitted `2/1/1`, unsupported format, `--format` without `--json`, and CLI help exposing `--format`.
- `E2-P2A-BUILD1`: `go build ./cmd/... ./internal/...` passed after P2-A changes.
- `E2-P2A-SMOKE1`: `go run .\cmd\anvien file-detail internal/filecontext/context.go --repo Anvien --json` returned compact JSON with `format=file-detail.compact`, `relationshipSamplesPerGroup=-1`, `relatedFiles=43`, and `282588` characters. The expanded equivalent smoke `--format expanded --relationships -1 --unresolved -1 --linked -1` returned `1282361` characters.
- `E2-P2A-DETECT1`: `anvien detect-changes --repo Anvien --scope all` ran before the P2-A commit. It reported `7` changed files, changed-file risk `high`, overall risk `high`, and affected processes rooted at `NewFileDetailCommand`, including `NewFileDetailCommand -> FileProjectionGraphInfo`. The high risk is a CLI-surface blast-radius warning; focused CLI tests, full CLI package tests, smoke command, and build evidence passed.
- `E2-P2B-IMPACT1`: Before P2-B edits, `anvien analyze --force` was rerun. `anvien impact file internal/httpapi/file_context.go --repo Anvien --direction upstream` showed 3 affected files (`file_context.go`, `listen.go`, `server.go`) and an API process blast-radius warning. `anvien impact symbol "Server.handleFileContext" --repo Anvien --direction upstream` showed LOW direct symbol blast radius; `anvien impact symbol boundedNonNegativeQueryInt --repo Anvien --direction upstream` showed API processes including `HandleFileContext -> WriteJSON`; HTTP `writeJSON` impact by UID showed broad API blast radius, so P2-B did not edit the shared response helper.
- `E2-P2B-SRC1`: `/api/file-detail` now accepts `format=compact|expanded`. Missing/empty `format` defaults to compact full-detail output through `BuildCompactFileContext`; `format=expanded` preserves expanded `FileContext`; unsupported format returns HTTP 400. Compact query options only pass `relationships`, `unresolved`, or `linked` limits when the query parameter is present, so omitted compact limits remain full-row.
- `E2-P2B-TEST1`: `go test ./internal/httpapi -run 'TestFileDetailEndpoint|TestFileContextEndpoint' -count=1` passed, then `go test ./internal/httpapi -count=1` passed. Tests cover compact default response, absolute path compact response, expanded response, explicit compact relationship limit with total/returned/omitted `2/1/1`, invalid format 400, missing file, outside repo path, and legacy `/api/file-context` absence.
- `E2-P2B-BUILD1`: `go build ./cmd/... ./internal/...` passed after P2-B changes.
- `E2-P2B-SMOKE1`: A local HTTP runtime was started with `go run .\cmd\anvien serve --host 127.0.0.1 --port 18765`, verified through `/api/info`, then stopped after smoke. `/api/file-detail?repo=Anvien&path=internal%2Ffilecontext%2Fcontext.go` returned compact JSON with `format=file-detail.compact`, `relationshipSamplesPerGroup=-1`, `relatedFiles=43`, and `282949` characters. The max-limited expanded request `format=expanded&relationships=100&unresolved=100&linked=100` returned `615355` characters.
- `E2-P2B-DETECT1`: `anvien detect-changes --repo Anvien --scope all` ran before the P2-B commit. It reported `6` changed files, changed-file risk `high`, overall risk `medium`, and affected processes under `HandleFileContext`, including `HandleFileContext -> BoundedNonNegativeQueryInt`, `HandleFileContext -> HasWindowsDrivePrefix`, and `HandleFileContext -> NormalizePath`. Focused HTTP tests, full HTTP tests, smoke endpoint, and build evidence passed.
- `E2-P2C-IMPACT1`: Before P2-C edits, `anvien analyze --force` was rerun. `anvien impact file internal/contracts/web_ui.go --repo Anvien --direction upstream` showed 2 affected files (`cmd/generate-web-contracts/main.go`, `internal/contracts/web_ui.go`) and a contracts process blast-radius warning; `anvien impact symbol FileContextResponse --repo Anvien --direction upstream` showed LOW direct symbol blast radius. Generator symbol lookup for `GenerateWebUIContracts` did not resolve because the generator entrypoint is the command package main.
- `E2-P2C-SRC1`: Updated the Go-owned Web UI contract source so `/api/file-detail` declares `format` in query params, route response type `FileDetailResponse`, and compact-default/expanded description. Generated TypeScript source now defines `CompactFileContextResponse`, compact schema/dict/table row types, and `FileDetailResponse = CompactFileContextResponse | FileContextResponse` while preserving expanded `FileContextResponse`.
- `E2-P2C-GENERATED1`: Ran `go run ./cmd/generate-web-contracts`, updating `contracts/web-ui/anvien-web-contract.schema.json` and `anvien-web/src/generated/anvien-contracts.ts` from `internal/contracts/web_ui.go`; no generated artifact was manually edited.
- `E2-P2C-TEST1`: `go test ./internal/contracts -count=1` passed and `go run ./cmd/generate-web-contracts --check` passed after regeneration. Contract tests assert the route uses `FileDetailResponse`, includes `format`, and generated TS contains compact response/union/related-file row types.
- `E2-P2C-BUILD1`: `go build ./cmd/... ./internal/...` passed after P2-C contract and generated-artifact changes.
- `E2-P2C-DETECT1`: `anvien detect-changes --repo Anvien --scope all` ran before the P2-C commit. It reported `8` changed files, `6` affected files, changed-file risk `high`, overall risk `medium`, and affected contract generator processes rooted at `Main`, including `Main -> CopyStringMap`, `Main -> GraphHealthResolutionHealthBucketStrings`, `Main -> LabelStrings`, `Main -> RelationshipDisplayPolicy`, and `Main -> TitleWords`.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`, `P3-B`, `P3-C`

- Pending implementation evidence.

## Closure Evidence

Use this section for final detect-changes, commit hash, and closure evidence when the plan reaches completion.

- Pending closure evidence.
