# AVmatrix Deadcode Audit Report

Date: 2026-05-26  
Repository: `E:\AVmatrix-GO`  
Indexed commit: `e504399`  
Scope: identify and report deadcode candidates. The audit phase did not delete source code.

Status update:

- 2026-05-26: the high-confidence Web source deadcode group listed in this report was removed after review and validation.
- 2026-05-27: the product-dead but test-referenced Web candidates listed in this report were removed after review and validation.
- Remaining unresolved sections are still review material for later decisions.

## Summary

This audit used AVmatrix as the primary codebase graph source, then cross-checked candidates with source search to reduce false positives.

Original findings:

- High-confidence Web source deadcode candidates: 7 files, about 1,162 lines. Status: removed on 2026-05-26.
- Product-dead but test-referenced Web candidates: 2 files, about 714 lines. Status: removed on 2026-05-27.
- Unused exported Web API client functions inside an otherwise live file: 9 functions.
- Standalone Go audit binaries needing owner decision: 3 `cmd/*` programs.
- No production Go internal package was flagged as package-level deadcode in this pass. `internal/testutil` and root `internal/providers` are test-only/test-container packages and were excluded.

Important: this report is not a removal plan. It is a decision list for review.

## Method

AVmatrix refresh:

```powershell
.\avmatrix\bin\avmatrix.exe analyze --force
```

Result:

```text
files: scanned=777 parsed=578 unsupported=199 failed=0
graph: nodes=89034 relationships=122030 path=E:\AVmatrix-GO\.avmatrix\graph.json
```

Graph schema checks:

```powershell
.\avmatrix\bin\avmatrix.exe cypher "MATCH (n) RETURN labels(n) AS labels, count(n) AS count ORDER BY count DESC LIMIT 50" --repo AVmatrix
.\avmatrix\bin\avmatrix.exe cypher "MATCH ()-[r]->() RETURN r.type AS relType, count(r) AS count ORDER BY count DESC LIMIT 80" --repo AVmatrix
```

Relevant relationship inventory:

| Relationship | Count |
|---|---:|
| `HAS_RESOLUTION_GAP` | 64,014 |
| `DEFINES` | 20,765 |
| `CALLS` | 8,437 |
| `USES` | 6,018 |
| `IMPORTS` | 4,368 |
| `ACCESSES` | 3,974 |
| `ENTRY_POINT_OF` | 700 |

Primary deadcode query for Web source files:

```cypher
MATCH (f)
WHERE labels(f) = 'File'
  AND f.filePath STARTS WITH 'avmatrix-web/src/'
  AND (f.filePath ENDS WITH '.ts' OR f.filePath ENDS WITH '.tsx')
OPTIONAL MATCH (src)-[r]->(f)
WHERE r.type = 'IMPORTS'
WITH f, count(r) AS imports
WHERE imports = 0
RETURN f.filePath, f.appLayer, f.functionalArea
ORDER BY f.filePath
```

Symbol-level query used on candidate files:

```cypher
MATCH (n)
WHERE labels(n) IN ['Function','Method','Class','Interface','TypeAlias','Const']
  AND n.filePath IN [candidate files]
OPTIONAL MATCH (src)-[r]->(n)
WHERE r.type IN ['CALLS','USES','ACCESSES','EXTENDS','INHERITS']
WITH n, count(r) AS refs
WHERE refs = 0
RETURN labels(n), n.filePath, n.name, n.startLine, n.endLine
ORDER BY n.filePath, n.startLine
```

Source cross-checks used `rg` for candidate names/imports across `avmatrix-web/src`, `avmatrix-web/test`, `cmd`, `internal`, and `avmatrix-launcher`.

## Classification

| Level | Meaning |
|---|---|
| High | AVmatrix reports no incoming imports/refs and `rg` found no external source/test usage. |
| Medium | No product source usage, but tests or non-product references still mention it. |
| Review | Not wired into primary runtime/build flow, but may be an intentional manual tool. |
| Excluded | Looks unused from one graph angle, but source review shows it is an entrypoint, alias import, test helper, or generated path. |

## High-Confidence Candidates

These files had no incoming `IMPORTS` in the AVmatrix graph and source search did not find product/test usage outside their own file.

| Candidate | Lines | Evidence | Notes |
|---|---:|---|---|
| `avmatrix-web/src/components/HelpPanel.the-press.tsx` | 395 | AVmatrix: no incoming imports. `context HelpPanelThePress`: `incoming: {}`. `rg HelpPanelThePress` only matches this file. | Looked like an older/alternate Help panel implementation. The separate `HelpPanel.tsx` candidate was also removed later as product-dead code. |
| `avmatrix-web/src/components/RightPanel.tsx` | 66 | AVmatrix: no incoming imports. `context RightPanel`: `incoming: {}`. `rg RightPanel` shows product code uses `RightPanelResizable`, not `RightPanel`. | Likely superseded by `RightPanel.resizable.tsx`. |
| `avmatrix-web/src/components/settings/ProviderConfigCard.tsx` | 108 | AVmatrix: no incoming imports. `rg ProviderConfigCard` only matches this file. | Settings UI appears to use local-runtime settings paths directly. |
| `avmatrix-web/src/config/ignore-service.ts` | 312 | AVmatrix: no incoming imports. `context shouldIgnorePath`: `incoming: {}`. `rg shouldIgnorePath` only matches this file. | Frontend ignore matcher appears unused; backend has current ignore handling under `internal/ignore`. |
| `avmatrix-web/src/core/ingestion/cluster-enricher.ts` | 243 | AVmatrix: no incoming imports. `context enrichClusters`: `incoming: {}`. `rg enrichClusters` only matches this file. | Old browser-side LLM cluster enrichment path; current runtime graph/search is backend/local-runtime driven. |
| `avmatrix-web/src/core/llm/index.ts` | 28 | AVmatrix: no incoming imports. `rg` found no imports from bare `core/llm` or `core/llm/index`. | Barrel file is unused; consumers import `session-client`, `types.local-runtime`, or settings modules directly. |
| `avmatrix-web/src/hooks/useSettings.ts` | 10 | AVmatrix: no incoming imports. `rg useSettings` only matches this file. | Wrapper around `useChatRuntime`; no current caller. |

Detailed symbols reported by AVmatrix in these files:

| File | Symbols |
|---|---|
| `HelpPanel.the-press.tsx` | `HelpPanelThePress` at line 323, plus local `HelpPanelProps`, `Tab`, `Kbd`, `InfoCard`, `TabContent`. |
| `RightPanel.tsx` | `RightPanel` at line 12. |
| `ProviderConfigCard.tsx` | `ProviderConfigCard` at line 31, `ProviderConfigCardProps`, `ApiKeyField`, `ModelField`. |
| `ignore-service.ts` | `shouldIgnorePath` at line 268, backed by local ignore constants. |
| `cluster-enricher.ts` | `CommunityNode`, `ClusterEnrichment`, `EnrichmentResult`, `LLMClient`, `ClusterMemberInfo`, `enrichClusters`, `enrichClustersBatch`. |
| `core/llm/index.ts` | Re-exports from `types.local-runtime`, `settings-service-local-runtime`, and `session-client`. |
| `useSettings.ts` | `useSettings` at line 3. |

## Product-Dead But Test-Referenced Candidates

These are not used by product source files, but still have tests or test-only references. Removing them requires updating or deleting the tests intentionally.

Status: removed on 2026-05-27. The test-only references were removed with the dead source.

| Candidate | Lines | Evidence | Notes |
|---|---:|---|---|
| `avmatrix-web/src/components/HelpPanel.tsx` | 711 | AVmatrix source-only import query reports no incoming imports from `avmatrix-web/src`. `rg HelpPanel` shows only `avmatrix-web/test/unit/Branding.local-only.test.tsx` imports/renders it. | Removed on 2026-05-27 with the test-only branding assertion. |
| `avmatrix-web/src/lib/utils.ts` | 3 | AVmatrix source-only import query reports no incoming imports from `avmatrix-web/src`. `rg generateId` shows test-only usage through `avmatrix-web/test/unit/utils.test.ts`. | Removed on 2026-05-27 with `utils.test.ts`. |

## Unused Exports In Live Files

`avmatrix-web/src/services/backend-client.ts` is a live file, but these exported functions have no incoming AVmatrix `CALLS/USES/ACCESSES` and `rg` found no call sites outside their own declarations.

| Function | Location | Endpoint/path touched | Evidence |
|---|---|---|---|
| `fetchServerInfo` | `backend-client.ts:308` | `/api/info` | No call sites found by AVmatrix or `rg fetchServerInfo(`. |
| `grep` | `backend-client.ts:560` | `/api/grep` | No call sites found by AVmatrix or `rg grep(`. The word `grep` appears in comments/UI labels only. |
| `fetchProcesses` | `backend-client.ts:605` | `/api/processes` | No call sites found. Current UI appears to load processes through other state/runtime paths. |
| `fetchProcessDetail` | `backend-client.ts:614` | `/api/process` | No call sites found. |
| `fetchClusters` | `backend-client.ts:623` | `/api/clusters` | No call sites found. |
| `fetchClusterDetail` | `backend-client.ts:632` | `/api/cluster` | No call sites found. |
| `getAnalyzeStatus` | `backend-client.ts:679` | `/api/analyze/{jobId}` | No call sites found. Current analyze flow appears to use SSE progress. |
| `getEmbedStatus` | `backend-client.ts:730` | `/api/embed/{jobId}` | No call sites found. |
| `cancelEmbeddings` | `backend-client.ts:737` | `/api/embed/{jobId}` `DELETE` | No call sites found. |

Decision note: these may be preserved deliberately as public browser-client API helpers. If they are not intended public API, they are good cleanup candidates.

## Go Standalone Command Review Candidates

AVmatrix found all `cmd/*/main.go` files have zero incoming file imports, which is expected for standalone binaries. Source/workflow search separates true entrypoints from review candidates.

Not candidates:

- `cmd/avmatrix/main.go`: primary CLI/runtime entrypoint.
- `cmd/generate-web-contracts/main.go`: documented contract generator; referenced by README, architecture docs, generated contract header, and contract tests.

Review candidates:

| Candidate | Evidence | Notes |
|---|---|---|
| `cmd/access-candidate-audit/main.go` | AVmatrix: standalone main with no incoming imports. `rg` found references only in semantic classification code, not package scripts/workflows. | Manual graph accuracy audit tool around `graphaccuracy.RunAccessCandidateAudit`. Keep if still part of manual QA process; otherwise document or remove later. |
| `cmd/property-access-audit/main.go` | Same pattern: standalone main, no workflow/package wiring found. | Manual property access audit tool around `graphaccuracy.RunPropertyAccessAudit`. |
| `cmd/graph-accuracy-probe/main.go` | Same pattern: standalone main, no workflow/package wiring found. | Manual Node-vs-Go graph accuracy probe. Could be historical migration tooling. |

These are not marked high-confidence deadcode because they are valid Go command entrypoints and may be intended manual tools.

## Go Package Reachability Check

Supplemental package reachability:

```powershell
go list ./cmd/... ./internal/...
go list -deps ./cmd/avmatrix
go list -deps ./cmd/generate-web-contracts
```

Comparing `./internal/...` to production dependencies from `cmd/avmatrix` plus `cmd/generate-web-contracts` left only:

- `github.com/tamnguyendinh/avmatrix-go/internal/providers`
- `github.com/tamnguyendinh/avmatrix-go/internal/testutil`

Both were excluded:

- `internal/providers` root contains `provider_parity_test.go` and is a test package container; provider subpackages are imported by `internal/analyze`.
- `internal/testutil` is intentionally test-only.

## Excluded False Positives

| Item | Why excluded |
|---|---|
| `avmatrix-web/src/main.tsx` | Web entrypoint; no incoming imports is expected. |
| `avmatrix-web/src/lib/lucide-icons.tsx` | AVmatrix source-only import query missed alias imports, but `rg lucide-icons` shows many product imports via `@/lib/lucide-icons`. |
| Local React callbacks such as `handleClick`, `onMove`, `onUp`, `tick` | Many have no incoming graph refs because they are local closures used by JSX/event handlers/effects; not actionable deadcode from graph-only evidence. |
| Go test functions and benchmarks | No incoming call refs is normal; Go test runner discovers them by name. |
| Cobra command constructors / CLI command functions | Some graph refs are indirect through command registration and closures. Need targeted impact review before any cleanup. |
| Generated files under `avmatrix-web/src/generated/` | Generated contract artifacts; update generator/source contracts first. |

## Recommended Review Order

1. High-confidence Web files: resolved, removed on 2026-05-26.
2. Product-dead but test-referenced Web files: resolved, removed on 2026-05-27.
3. For `backend-client.ts` unused exports, decide whether they are public API helpers or stale client methods.
4. For the three standalone Go audit commands, decide whether to keep them as documented manual tools, wire them into a QA workflow, or remove them.
5. Only after decisions, create a cleanup plan with build/test gates. Do not delete directly from this report.
