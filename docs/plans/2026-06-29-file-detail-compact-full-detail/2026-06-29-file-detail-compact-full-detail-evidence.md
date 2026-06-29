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

## E2 - P2 Evidence

Matching plan item(s): `P2-A`, `P2-B`, `P2-C`, `P2-D`, `P2-E`

- Pending implementation evidence.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`, `P3-B`, `P3-C`

- Pending implementation evidence.

## Closure Evidence

Use this section for final detect-changes, commit hash, and closure evidence when the plan reaches completion.

- Pending closure evidence.
