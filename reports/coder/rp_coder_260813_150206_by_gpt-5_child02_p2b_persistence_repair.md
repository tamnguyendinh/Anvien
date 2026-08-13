# Coder Report: Child 02 P2-B corrected-fact persistence repair

Verdict: **READY_FOR_QA**

Disposition: `HANDOFF_TO_ORCHESTRATION_MAIN_FOR_QA`

## Metadata

- Repository: `E:\Anvien`
- Slice: Child 02 P2-B only — Preserve corrected facts through Graph JSON and Ladybug
- HEAD: `c3821b32a65016ee6eb9f1e56ca1fd769bab1aed`
- Branch: `master`
- Staged paths: `0`
- Commit created: no
- Report created: `2026-08-13 15:02:06 +07:00`
- Report finalized after Owner-authorized lock release: `2026-08-13 15:15:39 +07:00`
- Candidate evidence disposition:
  - `E2-P2B-IMPACT1`: retained current pre-edit evidence; no new production owner added
  - `E2-P2B-SOURCE1`: implementation independently audited; scoped production correction complete
  - `E2-P2B-BUILD1`: PASS at the root repo-native build entrypoint
  - `E2-P2B-TEST1`: PASS in all affected packages after the full build
  - `E2-P2B-PARITY1`: PASS at Graph JSON, CSV/schema/load, and tagged native Ladybug readback boundaries

## Process containment

- The prior coder lane's conclusion and prior PASS claims were not trusted. This lane independently read the current source and complete dirty diff, then ran fresh commands from the current worktree.
- Only the closed authority set named by the Owner was used, plus permitted production/test source, build manifests/scripts, runtime state, and command output. The forbidden orchestration-only attachment was not opened or used.
- P2-C, P2-D, P2-E, C09-C16 reader implementation, semantic-search hydration, opaque-ID parsing, repeated-analyze behavior, identity design, relationship identity, scanner, UI, target source, plans, ledgers, benchmarks, and checklist state were not edited.
- No subagent, Supervisor, stage, or commit was used.
- A parity-clone direction was explicitly stopped and disqualified by Owner directive. No clone result is evidence in this report. The repo-local debug clone was deleted; `E:\Anvien\.tmp\p2b-full-build` no longer exists.
- The first build-lock investigation did not terminate editor-owned processes. A later explicit Owner override authorized terminating only processes proven by Windows Restart Manager to hold `E:\Anvien\anvien\bin\anvien.exe`. The exact current lockers, PIDs `7024` and `14348`, were terminated; no other process was terminated. The same OS query then returned zero remaining lockers.

## Invariant Family Map

Family: Child 02 P2-B corrected-fact persistence.

Authority / source facts:

- A Definition retains accepted identity, label, name, path, qualified name, construct range, optional declaring-token SelectionRange, meaning/occurrence facts, and its `DEFINES` endpoint.
- Coordinates are one-based lines, zero-based UTF-8 byte columns, exclusive ends. Coordinate `0` is data, not absence.
- SelectionRange is optional as one all-or-none range. Absent input remains absent/NULL; a partial range must not silently persist.
- Every CodeEmbedding chunk retains the explicit label supplied by `EmbeddableNode`; the label is not inferred from opaque node identity in P2-B.
- Every `DEFINES` relationship retains the exact File source ID and Definition target ID, and all endpoints must exist.

Sibling runtime surfaces:

1. C01 `internal/resolution/emit.go::emitDefinitionNodes` projects Definition facts into `graph.Node` and emits `DEFINES`.
2. C04 `internal/analyze/analyze.go::{Run,writeGraphSnapshotJSON}` generically serializes graph records to Graph JSON and is validate-only.
3. C02 `internal/lbugload/csv.go` selects exact Definition headers, constructs rows, exports aggregate/pair relationship CSVs, and feeds load.
4. C03 `internal/lbugschema/schema.go::NodeSchema` must exactly match every C02 Definition column and order.
5. C05 `internal/embeddings/pipeline.go::{EmbeddingUpdate,prepareBatch,CreateEmbeddingQuery}` carries explicit label into every persisted chunk.
6. C06 `internal/lbugschema/schema.go::EmbeddingSchema` declares the CodeEmbedding label column.
7. C07/C08 relationship CSV/schema owners preserve exact `DEFINES` endpoints; downstream COPY/fallback is transparent and validate-only.

Forbidden fallback / alternate behavior:

- no half-range persistence;
- no interpretation of coordinate `0` as absence;
- no validation applied to unrelated node-table projections;
- no reader/C16 change or opaque-ID label reconstruction in P2-B;
- no endpoint rewriting, compatibility adapter, identity/occurrence change, repeated-analyze change, or unrelated refactor.

Verification matrix:

| Surface | Positive case | Absent/subset/fault case | Required observable result | Current evidence |
|---|---|---|---|---|
| C01 graph projection | construct + selection with column `0` | absent selection | exact properties; absent keys remain absent; exact `DEFINES` endpoints | focused PASS |
| C04 Graph JSON | encode/decode corrected Function and Variable | absent selection + endpoint existence | field-level round-trip and zero record loss | focused PASS, pre-build |
| C02/C03 CSV/schema | Function/Class/Interface/CodeElement, Method, and all default Definition labels | alternating selection present/absent; partial range; unrelated File property | exact ordered header/row/schema; blank->NULL input; partial Definition range rejected; unrelated table unaffected | focused PASS, pre-build |
| C05/C06 embedding | multi-chunk explicit label | quotes, backslash, newline escaping | every update/query persists the exact explicit label | focused PASS, pre-build |
| C07/C08 endpoints/load | File -> each of 25 Definition labels | empty graph, missing endpoint, COPY/fallback fault regressions | zero corrected drops, zero missing endpoints, exact load pair | focused PASS, pre-build |
| Native Ladybug | reopen read-only DB and query all corrected fields, label, endpoints | present/absent selection, zero columns, missing row, malformed/read-only query | actual stored values/NULLs and 25/25 endpoint closure | tagged native test PASS after full build |

Residual same-invariant surfaces: none.

## Production diff independently kept

The four pre-existing production diffs were retained after full source audit; no production line was accepted merely because it came from the prior lane.

- `internal/resolution/emit.go`
  - keeps Definition identity, label, meaning, occurrence, qualified name, and `DEFINES` behavior unchanged;
  - adds construct `startCol`/`endCol`;
  - projects all four SelectionRange coordinates only when `SelectionRange != nil`.
- `internal/lbugload/csv.go`
  - adds `qualifiedName`, construct columns, and optional selection columns in exact matching order for the symbol, Method, and default Definition families;
  - uses blank CSV values for absent optional selection fields, preserving `0` as `"0"`;
  - rejects a partial selection range only on tables whose projection owns selection columns; unrelated File/Folder/Community/etc. projections are not rejected;
  - leaves relationship endpoint projection unchanged.
- `internal/lbugschema/schema.go`
  - gives every Definition dispatch family exact C02 column/type/order parity;
  - adds `label STRING` to CodeEmbedding;
  - leaves RelationPairs/RelationSchema behavior unchanged.
- `internal/embeddings/pipeline.go`
  - adds typed explicit `Label` to `EmbeddingUpdate`;
  - copies `EmbeddableNode.Label` to every chunk;
  - persists the escaped label property in every CodeEmbedding CREATE query.

No C16 reader behavior was implemented.

## Test diff kept or changed

Kept and aligned from the dirty worktree:

- `internal/embeddings/pipeline_test.go`: explicit label query and multi-chunk propagation; this lane strengthened label escaping.
- `internal/lbugload/csv_test.go`: existing Function CSV contract updated to the corrected exact header and value position.
- `internal/lbugload/load_test.go`: existing COPY contract expectations updated for symbol, Method, and default Definition families.
- `internal/lbugschema/schema_test.go`: existing CodeEmbedding schema assertion includes label.
- `internal/lbugnative/integration_ladybugdb_test.go`: expanded by this lane from a narrow sample to all 25 Definition labels, alternating present/absent SelectionRange, zero-valued columns, explicit embedding label, and 25 exact `DEFINES` endpoint queries after read-only reopen.

New focused P2-B tests retained:

- `internal/resolution/p2b_persistence_test.go:11` — `TestDefinitionGraphProjectionPreservesCorrectedFactsAndDefinesEndpoints`.
- `internal/analyze/p2b_snapshot_persistence_test.go:13` — `TestGraphSnapshotPreservesCorrectedDefinitionFactsAndDefinesEndpoints`.
- `internal/lbugload/p2b_definition_persistence_test.go:14` — exact 25-label CSV/load projection and zero-drop/endpoint closure.
- `internal/lbugload/p2b_definition_persistence_test.go:149` — partial selection rejection in symbol, Method, and default families.
- `internal/lbugload/p2b_definition_persistence_test.go:167` — validation does not widen to unrelated File projection.
- `internal/lbugload/p2b_definition_persistence_test.go:188` — every valid default Definition table has corrected columns.
- `internal/lbugschema/p2b_persistence_schema_test.go:9` — full ordered schema parity for all 25 Definition labels.
- `internal/lbugschema/p2b_persistence_schema_test.go:26` — explicit embedding label schema.
- `internal/lbugschema/p2b_persistence_schema_test.go:41` — every File-to-Definition endpoint pair.

One additional directly affected stale fixture was corrected:

- `internal/resolution/definition_collision_test.go::p1dDefinitionNode` now constructs the exact projected Definition payload with `startCol` and `endCol`. This preserves the existing collision guard and exact-reinsertion semantics; no compatibility fallback was added.

Removed:

- no production or test behavior was removed;
- the disqualified repo-local debug clone was removed and is not evidence.

## Blast radius

Pre-edit graph evidence was already completed and was not invalidated or repeated. No new production owner was added.

- `emitDefinitionNodes`: **CRITICAL**.
- `internal/lbugload/csv.go` file risk: **HIGH**.
- `nodeCSVRow`: **CRITICAL**, retained evidence `19` impacted symbols / `12` processes / `16` test symbols.
- File-detail and upstream impacts were already complete for all four editable files and all accepted C01/C02/C03/C05/C06 symbols named by the delegation.

These are blast-radius warnings, not edit bans. The candidate remains surgical to four accepted production owners.

## Command chronology and results

1. Closed authority, Git boundary, and report hash
   - Read the permitted authority set in full.
   - Supervisor report SHA-256: `988C137D5C864192C38180A47812BAB53FE315204DA4907F6FFCC15EEE659369` — matched exactly.
   - Initial Git boundary: HEAD `c3821b32a65016ee6eb9f1e56ca1fd769bab1aed`, branch `master`, exactly the declared dirty files, staged `0`.

2. Source/diff audit and formatting
   - Read the complete production/test diff and every allowed production/inspect-only surface.
   - `gofmt` on all dirty Go files — PASS.
   - `git diff --check` — PASS.

3. First fresh focused invocation
   - Command: `go test ./internal/resolution ./internal/analyze ./internal/lbugload ./internal/lbugschema ./internal/embeddings ./internal/lbugnative -count=1`
   - Result: FAIL only in `TestP1DDefinitionExactReinsertionIsIdempotentAndPreservesEnrichment`; the pre-existing test helper still built the old Definition projection without `startCol`/`endCol`.
   - Other five packages passed in that same invocation.

4. Correction and focused rerun
   - Corrected only the stale exact-projection fixture; production collision behavior remained fail-closed.
   - Command: `gofmt -w internal/resolution/definition_collision_test.go; go test ./internal/resolution ./internal/analyze ./internal/lbugload ./internal/lbugschema ./internal/embeddings ./internal/lbugnative -count=1`
   - Result: PASS for all six packages:
     - `internal/resolution` PASS
     - `internal/analyze` PASS
     - `internal/lbugload` PASS
     - `internal/lbugschema` PASS
     - `internal/embeddings` PASS
     - default-build `internal/lbugnative` PASS
   - Important limitation: the `ladybugdb`-tagged native integration test is excluded from this default invocation.

5. Root repo-native full build
   - Entrypoint established from root `package.json`: `npm run full-build` -> `scripts/full-build.ps1`.
   - First command: `npm run full-build` from `E:\Anvien`, with a 30-minute timeout.
   - First result: FAIL at the first `npm install` lifecycle build. Go could not write `E:\Anvien\anvien\bin\anvien.exe` because another process held the file.

6. File-lock debugging
   - `anvien doctor processes --json` identified the long-lived MCP/editor-owned process family.
   - Windows Restart Manager queried the exact file `E:\Anvien\anvien\bin\anvien.exe` and returned exactly two lockers: PID `14348` and PID `7024`.
   - Both exact PIDs are `anvien.exe mcp`; their parent `cmd.exe` processes are owned by Codex PID `12192`.
   - `C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien` is a junction to `E:\Anvien\anvien`, proving that the global MCP executable and the root build output are the same physical file.
   - Owner then explicitly authorized terminating only processes currently proven to hold that file.
   - Restart Manager re-confirmed the current exact lockers as `7024,14348`; only those two PIDs were terminated.
   - The same Restart Manager query returned an empty locker set immediately afterward.

7. Root repo-native full-build retry after exact lock release
   - Command: `npm run full-build` from `E:\Anvien`, with a 30-minute timeout.
   - Result: PASS in `116.7s`.
   - Package runtime was built and written at the canonical root path; npm install/build/global install completed; `anvien version` returned `1.2.8`.
   - Launcher/Web production build passed (`2943` modules transformed).
   - The same entrypoint's root runtime analyze passed: `1590` scanned / `684` parsed-code / `0` failed; graph `96,078` nodes / `135,199` relationships.
   - Vite chunk-size and mixed static/dynamic-import messages, plus npm's global install-script allowlist message, were warnings only; the entrypoint exited `0`.

8. Tagged native Ladybug post-build validation
   - Command: `go test -tags ladybugdb ./internal/lbugnative -run '^TestNativeLadybugPersistenceReadbackAndStream$' -count=1 -v`, with CGO and PATH pointed to the root full-build Ladybug runtime.
   - Result: PASS; test body `7.03s`, package `8.995s`, command invocation `25.7s`.
   - Observable proof: writable load then read-only reopen/query preserved corrected fields for all `25` Definition labels, alternating present/absent SelectionRange including zero columns, explicit CodeEmbedding label, and all `25` exact File-to-Definition `DEFINES` pairs. Missing-row, malformed-query, unknown-label, and read-only-write fault assertions passed.

9. Final post-build focused regression
   - Command: `go test ./internal/resolution ./internal/analyze ./internal/lbugload ./internal/lbugschema ./internal/embeddings ./internal/lbugnative -count=1`.
   - Result: PASS for all six affected packages in `9.6s`.
   - This rerun occurred after the successful full build and tagged native boundary; no code or test changed afterward.

10. Final containment audit
   - `git diff --check` — PASS.
   - HEAD/branch unchanged.
   - Dirty boundary: exactly `4` allowed production owners, `10` focused tests, and this one coder report; unexpected paths `0`.
   - Staged paths: `0`.
   - No commit.
   - No clone artifact remains.

## Full-build incident: confirmed root cause and owner

Symptom:

```text
go build ... -o E:\Anvien\anvien\bin\anvien.exe ./cmd/anvien
open E:\Anvien\anvien\bin\anvien.exe: The process cannot access the file because it is being used by another process.
```

Confirmed root cause:

```text
Codex-owned MCP PIDs 14348 and 7024
-> execute global npm Anvien package
-> global package is a junction to E:\Anvien\anvien
-> both processes hold E:\Anvien\anvien\bin\anvien.exe
-> package build-runtime and launcher build both overwrite that same canonical path
-> Windows rejects the write while the editor-owned MCP processes remain live
```

Resolution: an explicit Owner override authorized terminating only the processes currently proven to hold the canonical binary. Restart Manager re-confirmed PIDs `7024` and `14348`; only those PIDs were terminated, and the same OS probe then confirmed zero remaining lockers. The unchanged root full-build entrypoint passed on the next invocation. No build/package source or script was edited.

## Real-boundary evidence

Proven at the required build and post-build chronology:

- Graph projection keeps construct columns, optional selection, semantic fields, and exact `DEFINES` endpoints.
- Graph JSON encode/decode retains corrected fields, absent selection, record counts, and endpoint existence.
- CSV header/value/load-query and schema order match for all 25 Definition labels.
- Selection column `0` remains present; absent selection becomes blank CSV input; partial selection fails closed in all three dispatch families; unrelated File projection remains unaffected.
- Recording-runner load closes all 25 File-to-Definition pairs with zero skipped relationships and zero fallback failures.
- Every embedding chunk carries explicit label, and query escaping covers quotes, backslashes, and newlines.
- Empty graph, missing endpoint, COPY/fallback failure, rollback, and related existing focused regressions passed in the affected packages.
- Root repo-native full build and its built `anvien analyze . --force` runtime boundary passed.
- Tagged native Ladybug readback passed for all 25 Definition labels, present/absent SelectionRange, explicit embedding label, and all exact `DEFINES` endpoints.
- Native load/readback observed zero corrected record drops and zero missing endpoints.
- All affected focused packages passed again after full build and native validation.

## Failures and corrections

- Focused FAIL: exact-reinsertion test fixture encoded the old Definition projection. Corrected the fixture to include the two newly mandatory construct columns; focused suite then passed. Production collision rules were not weakened.
- Initial full-build FAIL: exact root binary lock confirmed by OS/runtime evidence. After explicit Owner authorization, only the two current lockers were terminated; the handle-release probe and unchanged root full-build retry both passed.
- Clone path: stopped and disqualified by Owner; all results discarded, artifact removed, and no clone evidence appears in acceptance claims.

## Residual unverified surfaces

- Same-invariant P2-B surfaces: none.
- `E2-P2B-DETECT1` remains an orchestration/main-owned process gate under the latest Owner directive and was intentionally not run by this lane.
- QA/Supervisor review, staging, and commit remain orchestration-owned; they are not residual persistence surfaces.

Residual same-invariant surfaces: **none**.

## Final Git boundary

- HEAD: `c3821b32a65016ee6eb9f1e56ca1fd769bab1aed`
- Branch: `master`
- Production paths modified: exactly the four allowed owners.
- Test paths modified/new: focused P2-B tests plus the directly affected exact-projection regression fixture described above.
- Report: this single coder report.
- Staged paths: `0`.
- Commit: none.

## Handoff

Status is `READY_FOR_QA` for Child 02 P2-B.

Orchestration/main receives the current root worktree, this durable report, full-build/runtime evidence, tagged native Ladybug readback, post-build focused regression, and exact Git boundary. Main retains ownership of detect-changes, QA/Supervisor routing, staging, commit, and any plan/ledger/checklist updates.

No P2-C/P2-D/P2-E implementation was performed by this lane.
