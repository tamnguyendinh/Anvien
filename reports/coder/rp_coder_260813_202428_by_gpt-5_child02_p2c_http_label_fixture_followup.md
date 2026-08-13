# Child 02 P2-C coder follow-up — HTTP explicit-label fixtures

- Status: `READY_FOR_SUPERVISOR_REREVIEW`
- Role: coder follow-up after Supervisor REJECT; this is not a new slice and not an acceptance verdict.
- Repository: `E:\Anvien`
- Date: 2026-08-13 (Asia/Bangkok)
- Next owner: Orchestration main

## Exact returned blocker

Supervisor rejected the current P2-C acceptance boundary for one sibling-regression defect only: the semantic and hybrid vector-row fixtures in `internal/httpapi/search_test.go` represented persisted `CodeEmbedding` rows without the explicit `label` column now required by C16. Production correctly failed closed on the missing label, so the semantic test errored and the hybrid test exercised its intentional BM25 fallback rather than semantic merging.

The official Ladybug VECTOR extension outage is explicitly non-blocking under the authoritative Supervisor disposition. This follow-up did not install, retry, download, or substitute any VECTOR extension.

## Authority consumed

Read in full before the edit:

- `AGENTS.md`
- `.agents/skills/working-rules/SKILL.md`
- `.agents/skills/coder/SKILL.md`
- `.agents/skills/backend-development/SKILL.md`
- the Child 02 P2-C section and full Child 02 plan
- all four Child 02 ledgers: plan, evidence, benchmark, and actual status
- `docs/contracts/graph-accuracy-contract.md`
- current `internal/embeddings/search.go`
- current `internal/httpapi/search.go`
- all of `internal/httpapi/search_test.go`
- relevant current diffs
- Supervisor report `reports/Supervisor/rp_supervisor_260813_200158_by_gpt-5_child02_p2c_affected_readers.md`

Supervisor report SHA-256 consumed and independently verified:

```text
9829F9D2FA483C1DDE97E981CA3AC1B93A4379F5DBEF409506CC66E3F76A60C7
```

This exactly matches the delegated expected hash.

## Invariant family map

- Family: persisted `CodeEmbedding` vector-row contract at the public HTTP semantic/hybrid consumer.
- SSOT: graph-accuracy contract reader rule, accepted P2-B explicit `CodeEmbedding.label` persistence, current production `vectorSearchQuery`, and the authoritative Supervisor return condition.
- Primary path: default `semanticSearchRunner` vector row -> `SearchService.searchSemantic` -> production `embeddings.SemanticSearch` -> explicit-label metadata hydration.
- Sibling path: explicit hybrid `vectorRows` -> semantic merge with BM25.
- Preserved fallback: hybrid embedder/semantic failure still returns BM25 results; no fallback semantic contract changed.
- Forbidden fallback: parsing or reconstructing table label from opaque NodeID.
- Failure path retained: production missing-label rows still return `semantic search match ... has no persisted label`.
- Stale artifact corrected: only the three simulated persisted vector rows in `internal/httpapi/search_test.go`.

Success criteria used:

1. Reproduce both Supervisor failures before editing.
2. Every HTTP test row that represents persisted vector data supplies explicit `label: Function`.
3. No production source changes, opaque-ID parsing, missing-label weakening, or hybrid fallback change.
4. Root full build and every delegated regression command complete successfully.
5. Exact containment, staged-zero, and shared-artifact preservation hold at handoff.

## Git boundary

### Before follow-up edit

- HEAD: `4d456446fcc49aed0c6d489aa9c63e00d030b53c`
- Branch: `master`
- Staged paths: `0`
- Existing dirty boundary: retained P2-C production/tests, native test, reusable Playwright/QA evidence, prior coder report, QA report, and Supervisor REJECT report.
- `internal/httpapi/search_test.go` was clean before this follow-up.

### After follow-up edit and before this report

- Follow-up edit and validation began at HEAD `4d456446fcc49aed0c6d489aa9c63e00d030b53c`.
- During final report hashing, an independent concurrent owner committed only the previously excluded orchestration-skill change as `43aff01e882b787c91da84676e23f3ae28c05720` (`docs(skills): update orchestration rules`).
- Final handoff HEAD: `43aff01e882b787c91da84676e23f3ae28c05720`.
- Branch: `master`
- Staged paths: `0`
- Owned follow-up code change: exactly `internal/httpapi/search_test.go`.
- New durable artifact: exactly this follow-up coder report.
- No stage, commit, main final detect-changes, cleanup, or deletion was performed.

An unrelated concurrent/shared-worktree change appeared during validation at `internal/aicontext/skills/orchestration/SKILL.md`. Orchestration main independently confirmed it was absent at 20:04, had `LastWriteTime` 20:18:53, and changed only `7. Quy tắc sau auto-compact` to `## 7. Quy tắc sau auto-compact`. It does not overlap P2-C. This coder preserved it byte-for-byte and never staged or committed it. A concurrent owner subsequently committed that single path at `43aff01e...`; Git shows the commit contains one insertion/one deletion in that file only. It remains explicitly excluded from P2-C ownership.

## Reproduction before edit

Command:

```text
go test ./internal/httpapi -count=1
```

Result: exit 1, package FAIL. Exact blocking evidence:

```text
--- FAIL: TestSearchServiceUsesSemanticRuntime (0.00s)
    search_test.go:108: Search() error = semantic search match "Function:alpha" has no persisted label
--- FAIL: TestSearchServiceHybridMergesBM25AndSemanticAndFallsBack (0.00s)
    search_test.go:354: results len = 2, want 3: []httpapi.SearchResult{... Sources:[]string{"bm25"} ...}
FAIL
FAIL github.com/tamnguyendinh/anvien/internal/httpapi 2.981s
```

This reproduced the Supervisor blocker without changing production behavior.

## Anvien and blast-radius gate

- `anvien --help`: completed successfully.
- `anvien analyze --force`: terminal wrapper timed out while the real analyze continued; PID 21016 was allowed to finish. `anvien status` then proved a fresh current index at 20:09:47, indexed/current commit `4d45644`, status up-to-date.
- `anvien file-detail internal/httpapi/search_test.go --repo Anvien --json`: current/non-stale, `changedSinceAnalyze=false`, parser `parsed`, file risk LOW, 93 symbols, 19 inbound references, 61 outbound references, 60 local relationships, zero unresolved, zero linked flows, three linked tests.
- `anvien impact file internal/httpapi/search_test.go --repo Anvien --direction upstream --json`: LOW, zero affected files and zero affected processes.
- Exact semantic and hybrid test function impact queries resolved LOW with zero direct/module/process impact. The nested method name did not resolve through symbol dispatch, so the successful current file-level impact is the governing pre-edit gate.

No HIGH or CRITICAL follow-up blast warning was found. The change was kept to three row values in one test file.

## Exact fixture correction

Only `internal/httpapi/search_test.go` changed:

1. Hybrid fixture row for `Function:shared` now contains `"label": "Function"`.
2. Hybrid fixture row for `Function:semantic-only` now contains `"label": "Function"`.
3. `semanticSearchRunner.QueryRows` default vector fixture for `Function:alpha` now contains `"label": "Function"`.

These are the only rows in the file that simulate persisted `CodeEmbedding` vector-index results. They now match production `vectorSearchQuery` column semantics:

```text
emb.nodeId AS nodeId, emb.label AS label, emb.chunkIndex AS chunkIndex,
emb.startLine AS startLine, emb.endLine AS endLine, distance
```

No blanket replacement was used. The fallback runner has no vector-row fixture and remains unchanged. There was no intentional missing-label fixture in `internal/httpapi/search_test.go`; C16's intentional missing-label negative remains in `internal/embeddings/search_test.go`.

## Proof production behavior did not change

- Follow-up edit operations touched only `internal/httpapi/search_test.go`.
- `internal/httpapi/search.go` remains unchanged.
- Current `internal/embeddings/search.go` still selects `emb.label AS label`.
- Missing label still returns `semantic search match %q has no persisted label`.
- `labelFromNodeID` remains absent.
- Hybrid behavior still falls back to BM25 for non-context semantic failures; its existing offline-embedder assertion remained unchanged and passed.
- C09-C15 production/tests/evidence and the existing native test were not edited in this follow-up.

## Mandatory clean-holder build evidence

Source inspection of `scripts/full-build.ps1`, `anvien/package.json`, and `anvien-launcher/build.ps1` established six concrete binary artifacts:

1. `E:\Anvien\anvien\bin\anvien.exe`
2. `E:\Anvien\anvien\bin\lbug_shared.dll`
3. global npm `anvien\bin\anvien.exe`
4. global npm `anvien\bin\lbug_shared.dll`
5. `E:\Anvien\anvien-launcher\AnvienLauncher.exe`
6. `E:\Anvien\anvien-launcher\server-bundle\anvien-server.exe`

The first Restart Manager query found four unique verified holders on the CLI/DLL artifacts:

- `anvien.exe` PIDs `14132`, `4544`, `13156`, `14356`
- verified `cmd.exe` launcher parents `21140`, `24524`, `23576`, `8884`
- holder commands were global `anvien.exe mcp`; parent commands were global `anvien.cmd mcp`
- launcher and server binaries had zero holders

Only those eight verified holder/launcher PIDs were stopped. Remaining verified PID count was zero. A subsequent Restart Manager pass returned zero holders for all six artifacts, and each artifact opened successfully with read/write access and `FileShare.None`.

An initial terminal invocation used a five-second wrapper and was terminated before producing a build verdict. Before retry, the entire rule was repeated: Restart Manager returned zero holders and every artifact again passed exclusive-open. The retry then completed successfully in 108.2 seconds.

When the unrelated concurrent skill edit appeared, main confirmed its ownership/exclusion. To ensure built evidence represented the final shared source boundary, the full clean-holder precheck was run again: Restart Manager returned zero holders and all six artifacts passed exclusive-open. No process needed to be killed on this final invocation. The later `43aff01e...` commit only recorded that already-built file content and introduced no post-build source change.

Final root build:

```text
npm run full-build
```

Result: exit 0 in 121.4 seconds. Packaged runtime, global CLI 1.2.8, launcher/server wrapper, TypeScript build, Vite production build, and final `anvien analyze . --force` completed. Final analyze reported 1,601 scanned files, 686 parsed code files, zero failed files, 96,652 nodes, and 135,814 relationships. Existing Vite chunk/dynamic-import warnings remained non-blocking.

## Final validation matrix

All commands below were run after the final successful full build.

### Exact HTTP blocker

```text
go test ./internal/httpapi -count=1
```

Result: exit 0; `ok github.com/tamnguyendinh/anvien/internal/httpapi 3.126s`. A final post-`gofmt` rerun also completed successfully in 2.992 seconds.

This proves both the semantic public consumer and hybrid merge fixtures now carry the explicit persisted label while the existing hybrid fallback test remains valid.

### C16 and C12-C15 regressions

```text
go test ./internal/embeddings ./internal/filecontext ./internal/mcp -count=1
```

Result: exit 0.

```text
ok github.com/tamnguyendinh/anvien/internal/embeddings 3.261s
ok github.com/tamnguyendinh/anvien/internal/filecontext 0.676s
ok github.com/tamnguyendinh/anvien/internal/mcp 5.476s
```

### Retained C09-C11 focused frontend regressions

```text
cd E:\Anvien\anvien-web
npm test -- --run test/unit/useAppState.local-runtime.test.tsx test/unit/ChatPanel.grounding-links.test.tsx test/unit/CodeReferencesPanel.graph-health.test.tsx
```

Result: exit 0; 3 files and 11/11 tests completed successfully in 6.74 seconds.

### Existing tagged native C16 seam

Repo-bundled native prerequisites were locally available under `.tmp/ladybug-native/v0.19.1/windows-x86_64`, so the accepted seam test was run without any remote extension action:

```text
go test -tags ladybugdb ./internal/embeddings -run '^TestNativeLadybugSemanticSearchHydratesPersistedExplicitLabels$' -count=1 -v
```

Result: exit 0; test completed successfully in 1.04 seconds, package 3.201 seconds.

No VECTOR endpoint or `INSTALL VECTOR` action was invoked. The external outage remains non-blocking exactly as Supervisor decided.

## Containment and artifact lifecycle

- `git diff --check`: clean.
- `git diff --cached --check`: clean.
- Staged paths: zero.
- Owned follow-up implementation diff: only three added `label` values in `internal/httpapi/search_test.go`.
- C12-C15 source files remain unchanged:
  - `internal/filecontext/context.go`
  - `internal/mcp/context.go`
  - `internal/mcp/detect_changes.go`
  - `internal/mcp/rename.go`
- No production source, plan, roadmap, ledger, contract, rule, skill, prior report, QA/Playwright artifact, native test, or C09-C15 test was edited by this follow-up.
- Shared/ownership-unknown `.tmp/ladybug-home`, `.tmp/ladybug-native`, and `.tmp/runtime-p2c` were inventoried and preserved; no delete or cleanup was attempted.
- No follow-up temp/log/dead official artifact was created. The five-second wrapper interruption produced no durable evidence file.
- Concurrent `internal/aicontext/skills/orchestration/SKILL.md` is excluded user/shared work and was committed independently as `43aff01e...`; it is not part of the P2-C diff or ownership.
- No `anvien detect-changes`, stage, commit, P2-D, or P2-E action occurred.

## Sibling surfaces checked

- Semantic `SearchService` path: corrected fixture and passing package test.
- Hybrid semantic merge: corrected explicit rows and passing three-result merge assertion.
- Hybrid BM25 fallback: unchanged and passing.
- Production missing-label failure: unchanged and passing in embeddings package.
- Production metadata-query failure: unchanged and passing in embeddings package.
- Native persistence/readback/hydration seam: unchanged and passing.
- C12-C15 validate-only packages: unchanged and passing.
- C09-C11 frontend tests: unchanged and passing.

## Legacy fallback status

No opaque-ID label parsing was added or restored. Production remains fail-closed when explicit persisted label is absent. Hybrid's documented BM25 fallback remains intact only for semantic runtime failure and was not weakened or reinterpreted.

## Residual unverified surfaces

None for the exact Supervisor return condition. Fully native remote VECTOR extension availability remains an explicitly non-blocking external evidence limitation and was not retried.

## Handoff

The exact HTTP fixture blocker is corrected, the public HTTP sibling regression and all required retained regressions are green on the final clean-built boundary, and no in-scope residual remains. Status is `READY_FOR_SUPERVISOR_REREVIEW`.

This report does not call P2-C accepted. Orchestration main should submit the current P2-C dirty boundary at HEAD `43aff01e882b787c91da84676e23f3ae28c05720` for a fresh independent Supervisor review. The already-committed orchestration-skill path is unrelated concurrent history and must remain excluded from P2-C ownership.
