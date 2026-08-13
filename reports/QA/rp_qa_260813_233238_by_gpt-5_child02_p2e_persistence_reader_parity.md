# QA Report — Child 02 P2-E persistence/reader parity

Scope: independent validation-only closure of the P2-E matrix on production HEAD `35939e7e6a621593d3d3065b9493a97c2c9a4f25`.

Runtime: freshly built Anvien `1.2.8`; CLI `E:\Anvien\anvien\bin\anvien.exe` SHA-256 `14DEB1820B58E4BBE68E5C8B542D09231CFAB49FA73521BC0163DA754588606B`; production preview at `127.0.0.1:5228` backed by `anvien serve` at `127.0.0.1:4848`.

Browser: visible in-app Browser for real built-runtime mount, followed by headed Chromium Playwright after Browser MCP large-graph waits were classified as tool timeouts. No equivalent Browser MCP wait was repeated after Owner direction.

No-fix boundary: no production source, existing product test, contract, roadmap, Child 02 ledger, accepted report/evidence, P2-D harness/evidence, target source, scanner, generated AGENTS/CLAUDE content, branch, index registration or remote VECTOR state was modified. No detect-changes, stage, commit, push or branch switch was performed.

This report records QA validation results only. It does not issue an acceptance verdict.

## Outcome summary

| Evidence ID | Validation result | Exact disposition |
|---|---|---|
| E2-P2E-BUILD1 | PASS | clean-holder full build exit `0`; one attempt, zero retry |
| E2-P2E-PARITY1 | PASS | field-level Graph JSON/Ladybug parity with zero missing/extra/mismatch/drop |
| E2-P2E-READERS1 | PASS | C09-C16 `8/8`, including headed mounted C09-C11 and local-native C16 |
| E2-P2E-REPEAT1 | PASS | M1-M7 `7/7`, covering unchanged/change/failure/current-read/no-backend/recovery |

Consolidated matrix:

- JSON: `reports/QA/child02-p2e-matrix/qa_child02_p2e_matrix_260813_232913.json`
- JSON SHA-256: `A2A80193A48B0583B1562FD6D0D8CD8DC2870B6C9B56547C05AE4DE293BA0B02`
- Markdown: `reports/QA/child02-p2e-matrix/qa_child02_p2e_matrix_260813_232913.md`
- Markdown SHA-256: `FD598EBF962FE10ECB88D71EEF310C20A28B2EFA0A88DB6655DA31E458E84AD0`

The paired JSON contains all 17 C01-C17 rows with scenario, command/boundary, expected, observed and QA validation verdict. C18/C19 remain explicit out-of-campaign exclusions and were not opened.

## Build evidence

Before the only full-build attempt, the gate inspected all six build artifacts. Restart Manager identified global `anvien.exe mcp` holders PIDs `15412`, `18544`, `24380`; PID `13688` was verified as the related `cmd.exe` launcher for `anvien.cmd mcp`. The three live holders were stopped; the launcher exited after its child holder; no unrelated process was killed. Restart Manager then reported zero holders for every artifact and exclusive-open succeeded for all six artifacts.

`npm run full-build` exited `0` in `116.702s`. The duration is validation evidence, not a product benchmark. The build JSON retains full stdout/stderr, command, UTC timestamps, artifact paths, PIDs, classifications, killed-process rows and post-build hashes.

## Persistence parity — C01-C08

Persistence denominator is `8/8`: five accepted edit rows C01/C02/C03/C05/C06 and three validate-only rows C04/C07/C08.

- Definition records: Graph JSON/Ladybug `36,611/36,611`.
- Compared fields: `id`, `label`, `name`, `filePath`, `qualifiedName`, construct `startLine/startCol/endLine/endCol`, and the four optional SelectionRange coordinates.
- Missing/extra/field mismatch/drop: `0/0/0/0`.
- SelectionRange present/absent/partial: `4,941/31,670/0`.
- Real zero `startCol`: `5,650`; zero is retained as data and absence is NULL.
- Exact DEFINES pairs: `36,611/36,611`; missing/extra/pair-count mismatch `0/0/0`.
- Definitions without exactly one DEFINES: `0`; missing graph/DEFINES endpoints `0/0`.
- Current normal artifact contains `0` CodeEmbedding rows, so it cannot supply a record-level label comparison. C05/C06/C16 are instead dynamically closed by fresh repository-local tagged native explicit-label persistence/readback/hydration.

## Affected readers — C09-C16

Reader denominator is `8/8`.

- C09: exact opaque-ID resolution passed focused unit coverage and headed mounted runtime; only the complete ID was active.
- C10: ambiguity failed closed with citation/panel counts `0/0`; unique grounding produced exactly one FUNCTION citation for `src/unique.ts`, `L10–11`.
- C11: graph range `10:0–12:0` converted once to `/api/file startLine=0`; scroll targeted line 10; lines 10–11 were highlighted and line 12 excluded.
- C12: nearest `internal/filecontext` regression exited `0`.
- C13-C15: nearest `internal/mcp` regressions exited `0`.
- C16: `internal/embeddings`, `internal/httpapi`, tagged native semantic hydration and tagged native persistence/readback commands all exited `0`. Remote VECTOR attempts: `0`.

The visible in-app Browser proved the real built UI mounted and began loading the current approximately 315 MB graph. Two Browser MCP JavaScript waits timed out/reset during that large load. Per Owner direction this is classified as a browser-tool/large-load limitation, not a product failure. The bounded alternative was headed Playwright on the already health-proven built runtime.

Headed Playwright passed C09-C11 `3/3`; product console warnings/errors, page errors, failed HTTP responses and request failures were all zero. Chromium emitted one exact SwiftShader `GPU stall due to ReadPixels` screenshot/WebGL diagnostic. It remains present under `nonBlockingConsoleDiagnostics`; the harness excludes only that exact browser diagnostic from the product error gate.

## Screenshot inspection

All six final screenshots were opened and inspected at original resolution:

1. Mounted runtime was READY with 5 nodes, 4 edges and correct `src/unique.ts` tree.
2. My AI panel, suggestions and composer mounted without clipping/overlap.
3. C09 showed only exact `uniqueTarget` cyan-active; no red node.
4. C10 ambiguous state showed no citation or Code Inspector.
5. C10 unique state showed exactly one FUNCTION citation at `src/unique.ts`, `L10–11`.
6. C11 selected viewer highlighted lines 10–11 and excluded line 12.

No broken font/asset, disappearing control, overlap, clipping or layout shift was observed. Individual screenshot hashes and inspection results are in the consolidated JSON.

## Repeated analyze/current read/failure/recovery

M1-M7 passed `7/7` on the newly built binary.

- unchanged input: exact semantic identities/facts/ranges/endpoints remained stable;
- changed input: the new `later` fact became current, `now` reduced to one and the removed second identity was absent;
- analyze failure: returned nonzero and created no graph at the expected normal path;
- current read: Ladybug succeeded while Graph JSON was absent for the complete MCP invocation;
- no readable backend: returned JSON-RPC non-success and no stale result;
- recovery: restored baseline source produced the original fact signature and readable current rows;
- exact fact proof: compared IDs, construct/selection ranges, `10/10` DEFINES and zero missing endpoints.

Semantic determinism does not claim whole-Ladybug byte identity; the two Ladybug container hashes differed while exact semantic facts remained stable.

The fresh tagged production binary dynamically validates native success, non-sentinel native failure and recovery. `Server.runCypherRead` permits Graph JSON fallback only for `errors.Is(openErr, ErrUnavailable)`. That exact sentinel represents compile-time native capability absence and is unreachable by deleting/corrupting a native artifact in the normal tagged binary; no dev/untagged binary or fake sentinel was manufactured.

## Integrity and edge-case lenses

The Data-Integrity lens required record-by-record identity/value/NULL/zero/endpoint conservation and explicit label readback. These checks found zero cross-artifact record mismatch, zero partial optional ranges, zero endpoint loss and no label reconstruction at the executed C16 seam.

The Edge-Case lens required unchanged/change ordering, stale-current failure prevention, no-readable-backend fail-closed behavior and recovery. The executed M1-M6 transitions returned no stale success and restored the baseline after owned failures.

No P2-B, P2-C or P2-D routing finding was produced. No outside-Child-02 scope was opened.

## Artifact lifecycle and containment

- Fault artifacts were restored; repeat evidence reports no held path remaining.
- Runtime was stopped using only manifest-owned PIDs `14132` and `12256`; ports `4848` and `5228` have zero listeners.
- P2-E temp root `E:\Anvien\.tmp\qa-child02-p2e` was removed through its owner-checked cleanup mode.
- Shared temp roots `.tmp/ladybug-home`, `.tmp/ladybug-native` and `.tmp/runtime-p2c` remain.
- Five main-owned plan documents retain their exact entry SHA-256 values.
- Final permitted boundary is the five pre-existing main-owned plan documents plus P2-E-owned harnesses, JSON/Markdown evidence, screenshots, consolidated matrix and this report.
- Staged paths remain `0`.

## Durable evidence hashes

| Artifact | SHA-256 |
|---|---|
| `reports/QA/child02-p2e-build/qa_child02_p2e_build_260813_225335.json` | `FE907F5EB2E36DBD46D50561872A932E3588F553462FE5211AE6075D638E9B8B` |
| `reports/QA/child02-p2e-build/qa_child02_p2e_build_260813_225335.md` | `08801E0A6CB2BD8359714C82EBDA50E5134EC709F8A707FA8BE114AE9FD60D6F` |
| `reports/QA/child02-p2e-parity/qa_child02_p2e_parity_260813_225713.json` | `CD712668A0D851D28AA93AFD18C7B257CCA45856AB291A2D9942421BC7452A91` |
| `reports/QA/child02-p2e-parity/qa_child02_p2e_parity_260813_225713.md` | `82625D6FCEA82A6129CF6E23629B5A4A20B6D5F97002B759643EB03701C6FF0F` |
| `reports/QA/child02-p2e-readers/qa_child02_p2e_readers_260813_225909.json` | `4F4DFA62CF080D806074A6BF6D233C699FA5F24060E83A5C8F7B4823AA592288` |
| `reports/QA/child02-p2e-readers/qa_child02_p2e_readers_260813_225909.md` | `7F86AE5BF18230F4D5DC8DB039F32AC0AD4B906239541C9EFBA60ACAAB8948B5` |
| `reports/QA/playwright/child02-p2e-affected-readers/child02-p2e-affected-readers-affected-readers.json` | `4F5AB1B618D5A056D95F808DBCFFF0F0A855CB7C7AC20BE61850CC1659E927D6` |
| `reports/QA/playwright/child02-p2e-affected-readers/child02-p2e-affected-readers-affected-readers.md` | `CFCADC359EEF9AB3CEC2B264B812BFD68807848E79F2502C0EE8BD1D074DB102` |
| `reports/QA/child02-p2e-repeated-analyze/qa_child02_p2d_repeated_analyze_260813_232124.json` | `95CBF8314585BE6C0AF754F346F73733AE84CE7CAD7AD4AA78A6B8E12E87856F` |
| `reports/QA/child02-p2e-repeated-analyze/qa_child02_p2d_repeated_analyze_260813_232124.md` | `E202F5E712912B0CF228C745B058542FF8871D24C5B4A4C76E80DC3C6B2BA7EE` |
| `scripts/qa-child02-p2e-validation.ps1` | `72390FE4043AA5E7560436BE0E3ADD7C7A226A025415791043FECEF55ADA128A` |
| `scripts/qa-child02-p2e-parity.go` | `4476AB18FA31F69127C5F93216AB5633B472CC7E6A6A3F1ABBAB1C16B2601989` |
| `playwright/child02-p2e-affected-readers.mjs` | `4DD174E53A2305F4F105394820967D66E80A536462E2665E6EFFE02940225488` |

Next action: orchestration/main should hand these exact artifacts and hashes to an independent Supervisor for the acceptance decision; QA must not open Pn-A/Pn-B/Pn-C or Child 03.

READY_FOR_SUPERVISOR
