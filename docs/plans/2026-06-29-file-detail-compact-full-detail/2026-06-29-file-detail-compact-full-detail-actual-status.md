# File Detail Compact Full Detail Actual Status

Title: File Detail Compact Full Detail
Date: 2026-06-29
Status: P0 Complete
Companion plan: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-plan.md`
Companion evidence: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-evidence.md`
Companion benchmark: `docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-benchmark.md`

## Purpose

This file records the real current state before implementation.

Implementation must not start until the target scope has a completed status row, evidence IDs, and a downstream plan decision.

This file does not replace `evidence.md`. It classifies current state from evidence.

Use exact evidence IDs from `evidence.md`, such as `E0-P0A-SRC1`, not broad section IDs such as `E0` or `E1`.

## Freshness / Refresh Rules

This actual-status file is a living current-state record, not a one-time P0 snapshot.

P0 records the baseline before implementation. After implementation begins, keep the Current Status Matrix updated so the next agent can trust it as the latest repo reality.

Update this file:

- after each completed implementation slice;
- before starting the next phase if repo state changed;
- whenever evidence changes a current-state classification;
- whenever the next phase's status assumptions, next action, or work steps need updating because reality differs from the previous status.

When refreshing status:

- update only the rows affected by the completed work or new evidence;
- use explicit transitions such as `missing -> correct`, `partial -> correct`, `fake-or-stub -> removed`, or `unbound -> bound-correct`;
- append a Status Refresh Log row instead of deleting the history;
- keep detailed proof in `evidence.md`; store only classifications, evidence IDs, touch mode, and plan consequences here.

## Scope

Target scope:

- `file-detail` compact full-detail representation, related-file metadata, CLI/API format selection, Web contract/types, and Web panel consumption.

Out of scope:

- Graph analysis semantics, file lookup normalization behavior, unrelated Web graph UI redesign, and manual edits to generated contract artifacts as source of truth.

## Relationship / Impact Evidence

For each target file, prefer:

```text
anvien file-detail <path> --repo <repo> --json
```

Record how many files the target is related to before deciding touch mode. A file with many relationships may still be editable, but the plan must narrow the exact phase, touch mode, and validation needed.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------|----------------------|-------------|
| `internal/filecontext/context.go` | `E0-P0A-FD1` | 42 | local 565, outbound 88, inbound 252, unresolved 542, linked flows 53, linked tests 8 | CRITICAL/HIGH scope warning from `E0-P0A-IMPACT1`; editable only in narrow P1 slices |
| `internal/cli/file_detail_command.go` | `E0-P0A-FD2` | 19 | local 54, outbound 67, inbound 9, unresolved 198 | CRITICAL scope warning from `E0-P0A-IMPACT2`; editable in P2-A only |
| `internal/httpapi/file_context.go` | `E0-P0A-FD3` | 20 | local 55, outbound 38, inbound 10, unresolved 175 | HIGH/CRITICAL scope warning from `E0-P0A-IMPACT3`; editable in P2-B only |
| `anvien-web/src/components/FileDetailPanel.tsx` | `E0-P0A-FD4` | 4 | local 20, outbound 5, inbound 4, unresolved 298 | lower Web blast radius from `E0-P0A-IMPACT4`; editable in P3-B only |
| `anvien-web/src/services/backend-client.ts` | `E0-P0A-FD5` | 22 | local 132, outbound 9, inbound 177, unresolved 392 | high file risk; edit only in P3-A after generated types exist |
| `internal/contracts/web_ui.go` | `E0-P0A-FD6` | 31 | local 160, outbound 59, inbound 38, unresolved 636 | high contract risk; edit only in P2-C and regenerate outputs |
| `contracts/web-ui/anvien-web-contract.schema.json` | `E0-P0A-FD7` | 0 | generated contract artifact | generated output, regenerate-only |

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Already behaves as required. | Preserve. Add evidence or tests only if needed. |
| `partial` | Some required behavior exists, but gaps remain. | Change only the missing parts. Preserve correct parts. |
| `wrong` | Current behavior, source, or contract is incorrect. | Replace with required behavior. Record the exact reason. |
| `missing` | Required behavior, source, or contract does not exist. | Implement the missing piece only. |
| `unbound` | Surface exists but is not wired to the real source, flow, or contract. | Bind to the real source only. Preserve approved surface. |
| `fake-or-stub` | Prototype, demo, mock, fallback, or placeholder data is being used as real behavior. | Remove fake behavior or replace it with an approved truthful state. |
| `blocked` | Source, authority, contract, or required evidence is unclear. | Stop. Do not implement until resolved. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| Builder expanded file detail | `BuildFileContext` still returns expanded `FileContext`; tests prove `Options{}` keeps the existing default sample limit while explicit negative full-detail limits are honored | Preserve expanded facts while adding compact full-detail representation | correct | 42 baseline related files for builder source | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E1-P1C-SRC1`, `E1-P1C-TEST1` | preserve expanded fields while wiring P2 surfaces |
| Compact representation | Compact DTO/schema/row tables, related-file metadata, full-default behavior, limited-output metadata, and real graph size/count measurements now exist | Compact full-detail DTO with schema, dictionaries, tables, row references, tuple ranges, and explicit version/format | correct | 43 current compact related-file rows after P1 graph changes | `E0-P0A-SRC1`, `E0-P0A-FD1`, `E1-P1A-SRC1`, `E1-P1A-TEST1`, `E1-P1B-SRC1`, `E1-P1B-TEST1`, `E1-P1C-TEST1`, `E1-P1C-MEASURE1` | start P2-A surface wiring |
| Compact limit/default contract | Builder-level compact output defaults unspecified relationship, unresolved, and linked limits to full-detail `-1`; explicit limits expose total, returned, and omitted counts | Default compact machine output is full-row unless explicit limits are supplied; limited output exposes total, returned, and omitted counts | correct | N/A | `E0-P0A-DISCUSS1`, `E0-P0A-REVIEW1`, `E1-P1C-SRC1`, `E1-P1C-LIMIT1` | preserve in P2 surface help/contracts |
| Related-file metadata | Compact builder-aware output exposes `relatedFiles` rows with graph-derived file metadata, direction flags, totals, and relationship-kind counts; current graph measurement reports 43 rows | `relatedFiles` table with file metadata, direction flags, counts, kind counts, and stable refs | correct | 43 current compact related-file rows after P1 graph changes | `E0-P0A-DISCUSS1`, `E0-P0A-SRC1`, `E0-P0A-FD1`, `E1-P1B-SRC1`, `E1-P1B-TEST1`, `E1-P1C-MEASURE1` | preserve in P2 surface wiring |
| CLI JSON surface | `--json` now writes compact `file-detail.compact` by default, `--format expanded` preserves expanded `FileContext`, compact output is full-row unless explicit limit flags are supplied, and unsupported/irrelevant format usage fails visibly | CLI supports compact full-detail and explicit expanded access with documented format behavior | correct | 19 baseline related files for CLI source | `E0-P0A-SRC2`, `E0-P0A-FD2`, `E0-P0A-IMPACT2`, `E2-P2A-IMPACT1`, `E2-P2A-SRC1`, `E2-P2A-TEST1`, `E2-P2A-SMOKE1` | preserve CLI behavior; start P2-B |
| HTTP `/api/file-detail` surface | Endpoint defaults to compact `file-detail.compact`, supports `format=expanded`, preserves full-row compact defaults unless explicit limit query params are present, and rejects unsupported format with HTTP 400 | Endpoint supports compact/expanded format selection and validates invalid format visibly | correct | 20 baseline related files for HTTP source | `E0-P0A-SRC3`, `E0-P0A-FD3`, `E0-P0A-IMPACT3`, `E2-P2B-IMPACT1`, `E2-P2B-SRC1`, `E2-P2B-TEST1`, `E2-P2B-SMOKE1` | preserve HTTP behavior; start P2-C contracts |
| Web/API contract source | Contract source now declares `/api/file-detail` query param `format`, response type `FileDetailResponse`, compact response shape, and expanded `FileContextResponse` compatibility | Contract declares compact response shape and format query behavior while preserving expanded compatibility | correct | 31 baseline related files for contract source | `E0-P0A-SRC4`, `E0-P0A-FD6`, `E2-P2C-IMPACT1`, `E2-P2C-SRC1`, `E2-P2C-TEST1` | preserve contract; continue P2-D/P2-E |
| Generated contract artifacts | Generated TS/schema now reflect `format`, `FileDetailResponse`, `CompactFileContextResponse`, and compact table/row types | Generated artifacts reflect updated source contract | correct | 0 related files for schema artifact | `E0-P0A-SRC4`, `E0-P0A-FD7`, `E2-P2C-GENERATED1`, `E2-P2C-TEST1` | preserve generated artifacts; do not edit manually |
| MCP/agent file-context surface | MCP `context file`, file layer, and file impact flows intentionally preserve expanded `FileContext`/`fileLayer` compatibility with tests guarding against accidental compact payloads | Preserve expanded MCP compatibility or add an explicit compact/expanded MCP contract with tests/smoke evidence | correct | 21/28/40 baseline related files for target dispatch/context/impact sources | `E0-P0A-SRC6`, `E0-P0A-FD8`, `E0-P0A-FD9`, `E0-P0A-FD10`, `E0-P0A-IMPACT5`, `E2-P2D-IMPACT1`, `E2-P2D-SRC1`, `E2-P2D-TEST1` | preserve expanded MCP compatibility; continue P2-E docs |
| User-facing command docs | README/RUNBOOK now document compact default machine output, expanded compatibility, full-row compact defaults, explicit limit metadata, and MCP expanded compatibility | Docs reflect final CLI/API/MCP behavior for compact/expanded/default and limit semantics | correct | N/A | `E0-P0A-SRC7`, `E0-P0A-REVIEW1`, `E2-P2E-DOC1`, `E2-P2E-DOC2`, `E2-P2E-SMOKE1` | P2 docs complete; start P3-A Web client |
| Web client | `fetchFileContext` now requests `format=compact`, adapts compact row tables into an expanded-compatible `FileDetailContext`, preserves compact `relatedFiles`, and fails visibly on malformed compact rows | Client requests compact format and adapts compact payload without hidden fallback | correct | 22 related files for client source | `E0-P0A-SRC5`, `E0-P0A-FD5`, `E3-P3A-IMPACT1`, `E3-P3A-SRC1`, `E3-P3A-TEST1`, `E3-P3A-BUILD1` | preserve adapter; edit FileDetailPanel in P3-B |
| Web panel render | Panel now renders compact-backed related-file metadata and preserves summary, quality, symbol tree, relationships, unresolved, and linked sections | Panel renders compact-backed sections and related-file metadata while preserving existing section semantics | correct | 4 related files for panel source | `E0-P0A-SRC5`, `E0-P0A-FD4`, `E0-P0A-IMPACT4`, `E3-P3B-IMPACT1`, `E3-P3B-SRC1`, `E3-P3B-TEST1`, `E3-P3B-BUILD1`, `E3-P3B-DETECT1` | preserve render; execute P3-C runtime validation |
| Runtime/browser validation | Docker server and web images now build successfully for built-runtime validation; runtime endpoint/UI validation is not yet run | Built runtime validates compact endpoint and visible file detail UI | partial | N/A | `E0-P0A-SRC5`, `E3-P3C-IMPACT1`, `E3-P3C-SRC1`, `E3-P3C-TEST1`, `E3-P3C-BUILD1`, `E3-P3C-DETECT1` | start built containers and run endpoint/browser validation |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|------------|----------------|----------------|----------|-------------------|
| R0 | 2026-06-29 | baseline before implementation | file-detail compact full-detail scope | initial classification complete | `E0-P0A-DISCUSS1`, `E0-P0A-GRAPH1`, `E0-P0A-SRC1..E0-P0A-SRC5`, `E0-P0A-FD1..E0-P0A-FD7`, `E0-P0A-IMPACT1..E0-P0A-IMPACT4` | P1/P2/P3 split into narrow slices; implementation may begin at P1-A |
| R1 | 2026-06-29 | readiness review supplement before implementation | MCP/agent surface, limit contract, docs surface | added missing surfaces and clarified compact full-detail default/limit contract | `E0-P0A-REVIEW1`, `E0-P0A-GRAPH2`, `E0-P0A-SRC6`, `E0-P0A-SRC7`, `E0-P0A-FD8..E0-P0A-FD10`, `E0-P0A-IMPACT5` | P2 expanded with P2-D/P2-E; P1-C must include limit semantics tests before surface wiring |
| R2 | 2026-06-29 | P1-A implementation | compact DTO and converter | compact representation `missing -> partial` | `E1-P1A-IMPACT1`, `E1-P1A-SRC1`, `E1-P1A-TEST1`, `E1-P1A-BUILD1` | continue P1-B to add related-file metadata table |
| R3 | 2026-06-29 | P1-B implementation | related-file metadata inventory | related-file metadata `partial -> correct` for builder-aware compact output | `E1-P1B-IMPACT1`, `E1-P1B-SRC1`, `E1-P1B-TEST1`, `E1-P1B-BUILD1` | continue P1-C to prove full parity and explicit limit semantics |
| R4 | 2026-06-29 | P1-C implementation | compact fact parity, full-default limits, limited-output metadata, and real graph measurement | builder expanded file detail `partial -> correct`; compact representation `partial -> correct`; compact limit/default contract `missing -> correct` | `E1-P1C-IMPACT1`, `E1-P1C-SRC1`, `E1-P1C-TEST1`, `E1-P1C-LIMIT1`, `E1-P1C-BUILD1`, `E1-P1C-MEASURE1`, `E1-P1C-DETECT1` | P1 complete; start P2-A CLI format selection |
| R5 | 2026-06-29 | P2-A implementation | CLI compact/expanded format selection and compact JSON encoding | CLI JSON surface `partial -> correct` | `E2-P2A-IMPACT1`, `E2-P2A-SRC1`, `E2-P2A-TEST1`, `E2-P2A-BUILD1`, `E2-P2A-SMOKE1`, `E2-P2A-DETECT1` | start P2-B HTTP format selection |
| R6 | 2026-06-29 | P2-B implementation | HTTP compact/expanded format selection and limit query semantics | HTTP `/api/file-detail` surface `partial -> correct` | `E2-P2B-IMPACT1`, `E2-P2B-SRC1`, `E2-P2B-TEST1`, `E2-P2B-BUILD1`, `E2-P2B-SMOKE1`, `E2-P2B-DETECT1` | start P2-C contract source and generated outputs |
| R7 | 2026-06-29 | P2-C implementation | Web/API contract source and generated artifacts | Web/API contract source `partial -> correct`; generated contract artifacts `partial -> correct` | `E2-P2C-IMPACT1`, `E2-P2C-SRC1`, `E2-P2C-GENERATED1`, `E2-P2C-TEST1`, `E2-P2C-BUILD1`, `E2-P2C-DETECT1` | start P2-D MCP/agent surface decision |
| R8 | 2026-06-29 | P2-D implementation | MCP/agent file-context compatibility | MCP/agent file-context surface `partial -> correct` with expanded compatibility preserved | `E2-P2D-IMPACT1`, `E2-P2D-SRC1`, `E2-P2D-TEST1`, `E2-P2D-BUILD1`, `E2-P2D-DETECT1` | start P2-E README/RUNBOOK docs |
| R9 | 2026-06-29 | P2-E docs update | README/RUNBOOK command and operator examples | User-facing command docs `partial -> correct` | `E2-P2E-DOC1`, `E2-P2E-DOC2`, `E2-P2E-SMOKE1`, `E2-P2E-DETECT1` | P2 complete; start P3-A Web client adapter |
| R10 | 2026-06-29 | P3-A implementation | Web client compact adapter | Web client `partial -> correct` for compact request and adapter conversion | `E3-P3A-IMPACT1`, `E3-P3A-SRC1`, `E3-P3A-TEST1`, `E3-P3A-BUILD1`, `E3-P3A-DETECT1` | start P3-B related-file render |
| R11 | 2026-06-29 | P3-B implementation | Web panel related-file render | Web panel render `partial -> correct` for related-file metadata while preserving existing sections | `E3-P3B-IMPACT1`, `E3-P3B-SRC1`, `E3-P3B-TEST1`, `E3-P3B-BUILD1`, `E3-P3B-DETECT1` | start P3-C built runtime validation |
| R12 | 2026-06-29 | P3-C build-blocker fix | Docker server image Linux build and Web image build | Runtime/browser validation `missing -> partial`; Docker images build, runtime endpoint/browser validation still pending | `E3-P3C-IMPACT1`, `E3-P3C-SRC1`, `E3-P3C-TEST1`, `E3-P3C-BUILD1`, `E3-P3C-DETECT1` | start built containers and run endpoint/browser validation |

## Phase Touch Map

Use this map to prevent accidental edits. A related file is not automatically editable.

`Plan-Relevant Relationship File` lists only a relationship file that can directly affect or be affected by the planned phase or slice. Do not copy the full `file-detail` relationship inventory into this map. Include only files whose relationship can affect the phase/slice decision, touch mode, or validation.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| `internal/filecontext/context.go` | `internal/cli/file_detail_command.go` | consumer | P1-A/P1-B/P1-C | inspect-only during P1; edit later in P2-A | `E0-P0A-FD1`, `E0-P0A-FD2` | keep expanded behavior stable until CLI slice |
| `internal/filecontext/context.go` | `internal/httpapi/file_context.go` | consumer | P1-A/P1-B/P1-C | inspect-only during P1; edit later in P2-B | `E0-P0A-FD1`, `E0-P0A-FD3` | keep endpoint stable until HTTP slice |
| `internal/filecontext/context.go` | `internal/mcp/target_dispatch.go` | MCP/agent consumer | P1/P2-D | inspect-only during P1; edit only in P2-D if contract decision requires it | `E0-P0A-SRC6`, `E0-P0A-FD8`, `E0-P0A-IMPACT5` | preserve MCP compatibility or add explicit tested MCP contract |
| `internal/filecontext/context.go` | `internal/mcp/context.go` | MCP context payload consumer | P1/P2-D | inspect-only during P1; edit only in P2-D if contract decision requires it | `E0-P0A-SRC6`, `E0-P0A-FD9` | `context file` payload must remain tested |
| `internal/filecontext/context.go` | `internal/mcp/impact.go` | MCP file impact consumer | P1/P2-D | inspect-only during P1; edit only in P2-D if contract decision requires it | `E0-P0A-SRC6`, `E0-P0A-FD10` | file-impact flow must remain tested |
| `internal/filecontext/context.go` | `internal/contracts/web_ui.go` | contract consumer | P1/P2-C | inspect-only until P2-C | `E0-P0A-FD1`, `E0-P0A-FD6` | update source contract before generated output |
| `internal/contracts/web_ui.go` | `contracts/web-ui/anvien-web-contract.schema.json` | generated output | P2-C | regenerate | `E0-P0A-FD6`, `E0-P0A-FD7` | never edit generated schema as source of truth |
| `internal/contracts/web_ui.go` | `anvien-web/src/generated/anvien-contracts.ts` | generated output | P2-C | regenerate | `E0-P0A-SRC4` | never edit generated TS as source of truth |
| final CLI/API/MCP behavior | `README.md` | user-facing docs | P2-E | edit if final defaults or examples require it | `E0-P0A-SRC7` | keep docs aligned with compact/expanded and limit behavior |
| final CLI/API/MCP behavior | `RUNBOOK.md` | operator docs | P2-E | edit if final defaults or examples require it | `E0-P0A-SRC7` | keep smoke examples aligned with final command/API behavior |
| `anvien-web/src/services/backend-client.ts` | `anvien-web/src/components/FileDetailPanel.tsx` | consumer/render | P3-A/P3-B | edit in separate slices | `E0-P0A-FD4`, `E0-P0A-FD5` | adapter first, render second |
| `anvien-web/src/components/FileDetailPanel.tsx` | `anvien-web/src/components/CodeReferencesPanel.tsx` | caller | P3-B | inspect-only unless prop contract changes | `E0-P0A-IMPACT4` | preserve caller behavior |
| worktree | `internal/aicontext/skills/Spec-to-SVG-Flow-Map/spec-to-svg-flow-map.vi.md` | unrelated untracked file | all phases | do-not-touch | `E0-P0A-WT1` | ignore unrelated user/generated work |

## Detailed Findings

### Builder Shape

Current state:

`internal/filecontext/context.go` owns `FileContext` and `BuildFileContext`. The current response is expanded and nested. It already includes full major sections, but it repeats IDs, paths, field names, and range object keys.

Required state:

```text
Keep expanded facts available, add compact full-detail representation with schemas, dictionaries, row tables, tuple ranges, and fact parity tests.
```

Evidence:

- `E0-P0A-SRC1`: current builder model and expanded fields.
- `E0-P0A-FD1`: large output and relationship counts for builder source.
- `E0-P0A-IMPACT1`: wide blast radius and CRITICAL struct/field impact.

Relationship and impact:

- Related file count: 42
- Relationship summary: local 565, outbound 88, inbound 252, unresolved 542
- Impact note: high/critical scope warning; implement in narrow slices and preserve expanded shape until consumers are ready.

Classification:

partial

Allowed next action:

Implement P1-A, P1-B, and P1-C with impact evidence before editing symbols.

Forbidden next action:

Do not remove expanded fields or collapse output to summary-only in P1.

### Compact Limit Contract

Current state:

Current CLI/API sample flags limit returned samples, but there is no compact schema that distinguishes complete output from intentionally limited output.

Required state:

```text
Default compact machine JSON returns full rows for the requested file unless the caller explicitly supplies limits; limited output exposes total, returned, and omitted counts.
```

Evidence:

- `E0-P0A-DISCUSS1`: user rejected cutting data and accepted lossless compaction.
- `E0-P0A-REVIEW1`: supervisor rejected readiness until this contract was explicit.

Relationship and impact:

- Relationship count: N/A; this is a contract invariant across builder, CLI, HTTP, MCP, and Web consumers.
- Impact note: tests must prove both default full-row behavior and explicit limited behavior before CLI/API/MCP wiring is accepted.

Classification:

missing

Allowed next action:

Add P1-C tests for full default and limited output metadata, then surface the behavior in P2-A through P2-E.

Forbidden next action:

Do not make compact output shorter by silently preserving sampled-only data as if it were complete file detail.

### CLI Surface

Current state:

`internal/cli/file_detail_command.go` writes expanded `FileContext` when `--json` is provided. Existing flags limit samples per group but do not compact the structure.

Required state:

```text
CLI supports compact full-detail output and explicit expanded output, with help and tests documenting behavior.
```

Evidence:

- `E0-P0A-SRC2`: CLI command owner and flags.
- `E0-P0A-FD2`: CLI owner relationship/risk counts.
- `E0-P0A-IMPACT2`: CLI blast-radius evidence.

Relationship and impact:

- Related file count: 19
- Relationship summary: local 54, outbound 67, inbound 9, unresolved 198
- Impact note: critical scope warning for command helpers.

Classification:

partial

Allowed next action:

Edit in P2-A after P1 compact model exists.

Forbidden next action:

Do not change CLI default/format behavior without updating help and tests in the same slice.

### HTTP and Contract Surface

Current state:

`/api/file-detail` returns expanded `FileContextResponse`; contract source and generated TypeScript describe the expanded-only shape.

Required state:

```text
HTTP supports compact and expanded format selection; contract source and generated outputs document both safely.
```

Evidence:

- `E0-P0A-SRC3`: HTTP endpoint owner.
- `E0-P0A-SRC4`: contract and generated type owner.
- `E0-P0A-FD3`, `E0-P0A-FD6`, `E0-P0A-FD7`: relationship/risk counts.
- `E0-P0A-IMPACT3`: HTTP blast-radius evidence.

Relationship and impact:

- HTTP related file count: 20
- Contract source related file count: 31
- Generated schema related file count: 0
- Impact note: endpoint and contract changes must be split; generated output is regenerate-only.

Classification:

partial

Allowed next action:

Edit HTTP in P2-B and contract source/generation in P2-C.

Forbidden next action:

Do not manually edit generated schema/TS as source of truth.

### MCP and Agent Surface

Current state:

`internal/mcp/target_dispatch.go` builds file context through `mcpBuildRepoFileContext`; `internal/mcp/context.go` returns `fileContext` in context file payloads; `internal/mcp/impact.go` uses the same path for file impact flows. These surfaces currently expose expanded file-context data or depend on expanded file-context helpers.

Required state:

```text
MCP/agent file-context consumers either preserve expanded payload compatibility with tests or expose an explicit compact/expanded contract with tests and smoke evidence.
```

Evidence:

- `E0-P0A-SRC6`: MCP source ownership and file-context consumers.
- `E0-P0A-FD8`, `E0-P0A-FD9`, `E0-P0A-FD10`: relationship and unresolved counts for MCP target files.
- `E0-P0A-IMPACT5`: target dispatch blast radius across MCP files and `ContextToolInternal -> ContextCandidate`.

Relationship and impact:

- `target_dispatch.go`: 21 unique related files, local 8, outbound 79, inbound 51, unresolved 105.
- `context.go`: 28 unique related files, local 50, outbound 111, inbound 48, unresolved 261.
- `impact.go`: 40 unique related files, local 115, outbound 128, inbound 47, unresolved 497.
- Impact note: CRITICAL blast-radius warning on MCP dispatch helpers; proceed with a narrow P2-D slice.

Classification:

partial

Allowed next action:

Implement P2-D after compact model and surface decisions exist; choose preserve-expanded or explicit compact/expanded MCP behavior, then test `context file` and file-impact flows.

Forbidden next action:

Do not let shared file-context model changes alter MCP payload shape accidentally or without focused tests.

### User-Facing Docs

Current state:

README/RUNBOOK document `anvien file-detail`, `anvien context file`, and `/api/file-detail` examples without compact/expanded and limit semantics.

Required state:

```text
Docs either reflect the final compact/expanded/default behavior or record that no change is required because defaults remain compatible.
```

Evidence:

- `E0-P0A-SRC7`: README/RUNBOOK command and API example locations.
- `E0-P0A-REVIEW1`: supervisor finding that docs scope must be included or explicitly excluded.

Relationship and impact:

- Relationship count: N/A.
- Impact note: generated `AGENTS.md`/`CLAUDE.md` must not be edited as source of truth.

Classification:

partial

Allowed next action:

Implement P2-E after CLI/API/MCP behavior is final for this plan.

Forbidden next action:

Do not leave stale docs if defaults, format parameters, or limit semantics changed.

### Web Consumer

Current state:

Web client and `FileDetailPanel` consume expanded `FileContextResponse`. The panel already renders major detail sections, but it does not have a rich related-file metadata section because the backend does not publish one.

Required state:

```text
Web client fetches compact file detail, adapts it without hidden fallback, and panel renders existing detail sections plus related-file metadata.
```

Evidence:

- `E0-P0A-SRC5`: Web client and panel ownership.
- `E0-P0A-FD4`, `E0-P0A-FD5`: relationship/risk counts for Web surfaces.
- `E0-P0A-IMPACT4`: Web component blast radius.

Relationship and impact:

- Panel related file count: 4
- Client related file count: 22
- Impact note: Web changes are narrower but must follow generated contract changes.

Classification:

partial

Allowed next action:

Edit Web client in P3-A and panel in P3-B after P2-C.

Forbidden next action:

Do not add hidden expanded fallback that masks compact parse failures.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | Compact representation now has DTO/schema/conversion, but later facts are still pending | P1-A complete; preserve converter while continuing P1-B/P1-C |
| P1-B | Builder-aware compact output now includes related-file metadata table | P1-B complete; preserve `relatedFiles` in P1-C/P2 |
| P1-C | Builder-level compact parity, full-default semantics, explicit limit metadata, build, and real graph measurements now exist | P1-C complete; preserve builder behavior while wiring command/API/MCP surfaces |
| P2-A | CLI compact default, expanded mode, explicit limit behavior, help, smoke size, tests, and build evidence now exist | P2-A complete; preserve CLI behavior while wiring HTTP |
| P2-B | HTTP compact default, expanded mode, explicit query limit behavior, invalid format 400, tests, build, and local runtime smoke evidence now exist | P2-B complete; preserve HTTP behavior while updating contracts |
| P2-C | Contract source and generated artifacts now expose `FileDetailResponse`, compact shape, expanded compatibility, and `format` query param | P2-C complete; preserve generated artifacts from source only |
| P2-D | MCP/agent file-context consumers intentionally preserve expanded payload compatibility with tests guarding `context file` and file impact shapes | P2-D complete; preserve expanded MCP behavior |
| P2-E | README/RUNBOOK now document compact default, expanded compatibility, full-row compact defaults, explicit limit metadata, and MCP expanded compatibility | P2-E complete; preserve docs during P3 |
| P3-A | Web client now requests compact file-detail and adapts compact rows into `FileDetailContext` | P3-A complete; preserve adapter and start P3-B render work |
| P3-B | Panel renders related-file metadata from compact adapter and preserves existing sections | P3-B complete; start P3-C runtime validation |
| P3-C | Docker images build; runtime endpoint/browser validation still pending | keep P3-C as final Web/runtime validation |

## Implementation Gate

- [x] Target scope is listed in Current Status Matrix.
- [x] Each target unit has a status.
- [x] Each status has evidence IDs.
- [x] Each target file has relationship count evidence from `file-detail` when applicable.
- [x] Phase Touch Map lists plan-relevant relationship files that can affect the current phase/slice.
- [x] Phase Touch Map defines touch mode for every plan-relevant relationship unit that may be affected.
- [x] Correct parts are marked preserve-only.
- [x] Partial, missing, wrong, unbound, and fake-or-stub parts have exact next actions.
- [x] Blockers are recorded, if any.
- [x] Next phase status assumptions, next action, and work steps have been updated from this status file when needed.
- [x] Status Refresh Log has R0 baseline and R1 readiness-supplement rows.
- [x] If implementation has started, affected Current Status Matrix rows have been refreshed from latest evidence.
- [x] If refreshed statuses changed next work, only the stale next-phase status assumptions, next action, or work steps have been updated before the next phase.

## Final P0 Decision

Choose one:

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [x] P0 complete. Next phase can proceed unchanged.
- [ ] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

P0 is complete after readiness supplementation. Implementation may start at P1-A using the updated plan. The plan now preserves expanded facts, adds compact full-detail output, adds related-file metadata, defines full-default versus limited-output semantics, splits CLI/API/contract/MCP/docs/Web slices, and requires focused impact evidence before each code edit.
