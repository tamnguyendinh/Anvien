# Supervisor Report: Child 02 P2-E persistence and affected-reader closure

Verdict: PASS

Disposition: `ACCEPT_P2E_AND_HAND_BACK`

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260813_235237_by_gpt-5_child02_p2e_validation.md`
- Review time: `2026-08-13 23:52:37 +07:00` (`Asia/Bangkok`)
- Reviewer: `gpt-5`
- Repository: `E:\Anvien`
- Branch: `master`
- Production HEAD: `35939e7e6a621593d3d3065b9493a97c2c9a4f25`
- Scope reviewed: Child 02 P2-E validation-only, and no later slice.
- Claim reviewed: current production HEAD preserves corrected Definition facts through Graph JSON and Ladybug, exposes them truthfully through all eight source-proven affected readers, remains semantically deterministic across repeated normal analyze/failure/recovery, and is ready for the isolated P2-E evidence/ledger commit.
- Authority: Owner delegation; `AGENTS.md`; full `working-rules`, `supervisor`, `Data-Integrity`, and `Edge-Case` skills; graph-accuracy contract; roadmap; all four Child 02 ledgers; accepted P2-A/P2-B/P2-C/P2-D reports and commits; current source/tests; all P2-E harnesses and evidence.
- Review mode: independent zero-trust, review-only, evidence review.
- Next owner: orchestration/main.

## Executive Decision

P2-E receives an unconditional `PASS`. The exact C01-C17 matrix is closed at the current production HEAD with `17/17` validation rows, `8/8` persistence rows, `8/8` affected readers, and `7/7` repeated-analyze scenarios. The build, raw record parity, reader boundaries, visible UI states, failure paths, recovery, hashes, and containment all agree with current source and focused tests. No P2-B, P2-C, or P2-D invariant must be reopened.

The text `Code not available in memory for src/unique.ts` visible inside the citation card in screenshots 05 and 06 is explicitly classified as non-blocking and outside the owned C10/C11 invariant. It is not evidence of a backend/file-reader failure: the citation-card renderer intentionally has no snippet hydration, while focusing the same reference selects the exact persisted node, performs a successful `/api/file?path=src%2Funique.ts&repo=qa-p2e&startLine=0&endLine=60` request, renders the source in the selected Code Inspector, targets line 10, highlights lines 10-11, and excludes line 12. No follow-up on this presentation limitation is required before P2-E acceptance.

Blocking findings: none.

Non-blocking product findings: none.

Residual unverified same-invariant surfaces: none inside P2-E.

## Git and Worktree Boundary

The final pre-report check established:

- branch `master` and exact HEAD `35939e7e6a621593d3d3065b9493a97c2c9a4f25`;
- staged paths `0`;
- exactly five tracked dirty paths, all main-owned plan documents;
- exactly 22 untracked P2-E-owned paths before this report;
- no production or existing product-test diff;
- no target-source, scanner, generated rule, contract, accepted-report, Pn-A/Pn-B/Pn-C, Child 03, or unrelated dirty path;
- `git diff --check` and `git diff --cached --check` exit `0`;
- this Supervisor report is the only path added by this review lane.

The five protected document hashes remained byte-identical from review entry through the final pre-report check:

| Protected document | SHA-256 |
|---|---|
| roadmap | `972A581FAD301922659CC8DE2A92469BBC2142EEE037033B556E3839568347DC` |
| Child 02 actual status | `B142DF0153CF631C1EC077486A0D2AF3A3F48FCA6C1EBDD1B4A9300D768886D3` |
| Child 02 benchmark | `1D822F6F6AD691C59387C5B06F7D49DD1E172F941F4713FAD03018FA4BC2BB99` |
| Child 02 evidence | `4899F326D16C36A7E13982275011982341D593936CBB05F5DB6E60BD8BBC3A89` |
| Child 02 plan | `49BF9DB9AA528A8BA25710263F950476557C61C0190644B2A0FC932C6526916F` |

## Predecessor and Denominator Authority

The predecessor chain and durable reports were independently verified rather than inferred from the QA matrix:

| Slice | Isolated commit/report basis | Verified report SHA-256 | Decision retained |
|---|---|---|---|
| P2-A | `c3821b32a65016ee6eb9f1e56ca1fd769bab1aed` | `988C137D5C864192C38180A47812BAB53FE315204DA4907F6FFCC15EEE659369` | exact C01-C19 inventory accepted |
| P2-B | `4d456446fcc49aed0c6d489aa9c63e00d030b53c` | `24F403FC7A058AD36943A4664B0EFF8ABE2784644BA08BBC79D94B16C8CC750E` | C01-C08 persistence accepted |
| P2-C | `927a676653963e8001d7789291010d5b819bac83` | `F54C4225520FFC57A6438CE3F4C3B0816FE8095B31AFB55FD004E91A89808698` | C09-C16 affected readers accepted `8/8` |
| P2-D | current HEAD `35939e7e6a621593d3d3065b9493a97c2c9a4f25` | `C5582385945607FDB730B918EAA758AB2C7A3D37D3772B884D667B0E6133204D` | C04/C17 repeat/failure accepted `7/7` |

The accepted accounting remains:

- C01-C08: eight persistence rows;
- C09-C16: eight affected-reader rows;
- C17: one P2-D-only read/failure boundary;
- C04 participates in repeat validation without becoming an extra physical row;
- C18/C19 remain explicit out-of-campaign exclusions and were not opened;
- total P2-E rows: 17 unique rows, with no duplicate or missing accepted owner.

## Durable P2-E Artifact and Hash Inspection

All 22 P2-E-owned paths were inspected. The three delegated top-level hashes match exactly: QA report `41509F4379B3B5A4976BC7469831EF8E57EA902FABBAE4BD02184EEFB14D8D16`, consolidated JSON `A2A80193A48B0583B1562FD6D0D8CD8DC2870B6C9B56547C05AE4DE293BA0B02`, and consolidated Markdown `FD598EBF962FE10ECB88D71EEF310C20A28B2EFA0A88DB6655DA31E458E84AD0`.

| Artifact | Verified SHA-256 |
|---|---|
| `scripts/qa-child02-p2e-validation.ps1` | `72390FE4043AA5E7560436BE0E3ADD7C7A226A025415791043FECEF55ADA128A` |
| `scripts/qa-child02-p2e-parity.go` | `4476AB18FA31F69127C5F93216AB5633B472CC7E6A6A3F1ABBAB1C16B2601989` |
| `playwright/child02-p2e-affected-readers.mjs` | `4DD174E53A2305F4F105394820967D66E80A536462E2665E6EFFE02940225488` |
| build JSON | `FE907F5EB2E36DBD46D50561872A932E3588F553462FE5211AE6075D638E9B8B` |
| build Markdown | `08801E0A6CB2BD8359714C82EBDA50E5134EC709F8A707FA8BE114AE9FD60D6F` |
| parity JSON | `CD712668A0D851D28AA93AFD18C7B257CCA45856AB291A2D9942421BC7452A91` |
| parity Markdown | `82625D6FCEA82A6129CF6E23629B5A4A20B6D5F97002B759643EB03701C6FF0F` |
| readers JSON | `4F4DFA62CF080D806074A6BF6D233C699FA5F24060E83A5C8F7B4823AA592288` |
| readers Markdown | `7F86AE5BF18230F4D5DC8DB039F32AC0AD4B906239541C9EFBA60ACAAB8948B5` |
| repeat JSON | `95CBF8314585BE6C0AF754F346F73733AE84CE7CAD7AD4AA78A6B8E12E87856F` |
| repeat Markdown | `E202F5E712912B0CF228C745B058542FF8871D24C5B4A4C76E80DC3C6B2BA7EE` |
| consolidated JSON | `A2A80193A48B0583B1562FD6D0D8CD8DC2870B6C9B56547C05AE4DE293BA0B02` |
| consolidated Markdown | `FD598EBF962FE10ECB88D71EEF310C20A28B2EFA0A88DB6655DA31E458E84AD0` |
| Playwright JSON | `4F5AB1B618D5A056D95F808DBCFFF0F0A855CB7C7AC20BE61850CC1659E927D6` |
| Playwright Markdown | `CFCADC359EEF9AB3CEC2B264B812BFD68807848E79F2502C0EE8BD1D074DB102` |
| screenshot 01 | `931EB2E373A2BC07A45F2D376D32B367369E655FB2BB7AD94CEBB829F79E5228` |
| screenshot 02 | `2D666AC3F2BA23A4DFEB9807D359876E348333FD0E64C9DB4184541864402EA9` |
| screenshot 03 | `18E3DFE8582B09873A5ED7E054B0DF7C069E314F4BA34274475957DAC351BA6D` |
| screenshot 04 | `9AFC2268128F9EE43A056782D9AC6B2C04C1025E897AFDC917B818E06267B64B` |
| screenshot 05 | `949B933547AD4C288C29EA5C616FF621AD7F52A42180526AD4FCC63D6149DA7B` |
| screenshot 06 | `3EB4B78ACE93DF63BBA6FC1E16783C09C33E1A5EF9E1C1FACB9E16F3B1274569` |
| QA report | `41509F4379B3B5A4976BC7469831EF8E57EA902FABBAE4BD02184EEFB14D8D16` |

The consolidated matrix itself is correctly non-authoritative: `acceptanceVerdictIssued=false` and `handoffState=READY_FOR_SUPERVISOR`. Its 16 referenced source-evidence/screenshot hashes all match the files present.

## E2-P2E-BUILD1: Clean-Holder Full Build

Decision: PASS.

- Six build artifacts were registered with Windows Restart Manager.
- PIDs `15412`, `18544`, and `24380` were verified as global `anvien.exe mcp` processes from the global npm Anvien binary.
- PID `13688` was verified as `cmd.exe /c ... anvien.cmd mcp`, the related launcher parent of PID `24380`.
- The three live verified holders were killed. PID `13688` exited with its child and therefore was not separately killed. No unrelated process was killed.
- Final Restart Manager holder count was zero for each of six artifacts.
- Exclusive read/write open with no sharing succeeded for all six artifacts before build.
- `npm run full-build` ran once, with zero retry, exit `0`, from `2026-08-13T15:51:36.6717111Z` to `2026-08-13T15:53:33.3910896Z`, duration `116.702s`.
- The fresh CLI is `67,040,768` bytes, version `1.2.8`, SHA-256 `14DEB1820B58E4BBE68E5C8B542D09231CFAB49FA73521BC0163DA754588606B`.
- The CLI file still has that exact hash and version during Supervisor review.

The build output contains warnings about Vite chunk sizing/dynamic import composition, but no build failure. Those existing warnings do not contradict a P2-E behavior invariant.

## E2-P2E-PARITY1: C01-C08

Decision: PASS, persistence denominator `8/8`.

The parity harness was read in full. It streams current `.anvien/graph.json`, queries current `.anvien/lbug` through the tagged native read runner, keys records by opaque Definition ID, and compares these 13 values exactly:

`id`, `label`, `name`, `filePath`, `qualifiedName`, construct `startLine/startCol/endLine/endCol`, and optional `selectionStartLine/selectionStartCol/selectionEndLine/selectionEndCol`.

Independent raw-evidence reconciliation:

- Graph JSON/Ladybug Definition records: `36,611/36,611`;
- missing/extra/field mismatch/drop: `0/0/0/0`;
- SelectionRange present/absent/partial: `4,941/31,670/0`;
- missing construct coordinates: `0`;
- real `startCol=0`: `5,650` records;
- absence is retained as NULL rather than coerced to zero;
- exact Graph JSON/Ladybug `DEFINES` pairs: `36,611/36,611`;
- missing/extra/pair-count mismatch: `0/0/0`;
- Definitions without exactly one `DEFINES`: `0`;
- missing graph endpoints and missing `DEFINES` endpoints: `0/0`.

Current source corroborates the data evidence: C01 emits full construct and optional all-or-none selection coordinates; C02/C03 preserve ordered fields across symbol, Method, and every valid default Definition table; optional CSV values preserve zero and encode absence as blank/NULL input; partial selection fails closed; C04's generic snapshot remains field-transparent; C05/C06 persist typed explicit `CodeEmbedding.label`; C07/C08 preserve exact relationship endpoints.

The normal artifact has `0` CodeEmbedding rows. This is not represented as normal-artifact label parity. C05/C06/C16 are instead dynamically proven by the fresh repository-local tagged persistence/readback and production hydration seams. Both tagged commands exited `0`, and no remote VECTOR install/query/substitution was attempted.

## E2-P2E-READERS1: C09-C16

Decision: PASS, affected-reader denominator `8/8`.

The raw reader command evidence proves:

- focused frontend: three files, `11/11` tests, exit `0`;
- C12 `internal/filecontext`: exit `0`;
- C13-C15 `internal/mcp`: exit `0`;
- C16 `internal/embeddings ./internal/httpapi`: exit `0`;
- tagged `TestNativeLadybugSemanticSearchHydratesPersistedExplicitLabels`: exit `0`;
- tagged `TestNativeLadybugPersistenceReadbackAndStream`: exit `0`;
- remote VECTOR attempts: `0`.

Current source/test clearance:

- C09 `resolveNodeIds`: exact set membership only; suffix and near-match inputs are ignored.
- C10 `handleNodeGroundingReference`: exactly one matching persisted label/name identity is required; ambiguity returns without citation or panel.
- C11 `CodeReferencesPanel`: one-based graph lines and exclusive ends are converted once at the zero-based file-read boundary; zero-based UTF-8 byte columns are retained; selected source highlighting follows the exclusive-end rule.
- C12 `nodeRange`: copies all four corrected coordinates directly.
- C13 context payloads: use persisted UID/path/line values without opaque-ID reconstruction.
- C14 changed-symbol membership: compares persisted one-based node ranges to Git hunks without a second conversion.
- C15 rename: uses persisted UID/path/start line and converts once to the file array index.
- C16 semantic search: vector rows select `emb.label`; label propagates through chunk/dedup matching; missing label fails closed; metadata hydration selects the explicit label table; no active `labelFromNodeID` reconstruction exists.

### Browser and headed Playwright classification

The visible in-app Browser proved the freshly built application mounted and began the approximately 315 MB current-graph load. Two Browser MCP JavaScript waits timed out/reset during that large load. Under the Owner's explicit direction, the same wait was not attempted again. Because the product health/mount was already visible and no product error was observed, this is classified as a Browser MCP/large-load limitation, not a product failure.

The bounded headed Chromium run then exercised unchanged mounted production `AppStateProvider`, chat bridge, C09/C10 handlers, and C11 panel against external HTTP fixture responses on the health-proven built runtime. C09-C11 passed `3/3`. Product console warnings/errors, page errors, failed HTTP responses, and request failures were `0/0/0/0`.

The exact Chromium SwiftShader message `GPU stall due to ReadPixels` remains preserved under `nonBlockingConsoleDiagnostics`. The wrapper excludes only this exact screenshot/WebGL diagnostic from the product-console gate; all other warning/error classes still fail the run. It is a non-blocking browser-renderer diagnostic, not an Anvien reader failure.

## Independent Original-Resolution Screenshot Review

All six screenshots were opened directly at original resolution:

1. `01-mounted-production-fixture.png`: READY state, `qa-p2e`, one `src/unique.ts` file, five nodes/four edges, and no broken layout/asset.
2. `02-before-c09-tool-result.png`: mounted My AI panel, suggestions, composer, graph, and controls are coherent without clipping or overlap.
3. `03-c09-exact-id-only.png`: only `uniqueTarget` is cyan-active; suffix/near-match candidates remain dim and no red blast-radius node is present.
4. `04-c10-ambiguous-fails-closed.png`: the ambiguous `Function:sharedName` response produces no AI citation and no Code Inspector.
5. `05-c10-unique-persisted-reference.png`: exactly one FUNCTION citation is present for `uniqueTarget`, `src/unique.ts`, `L10-11`.
6. `06-c11-lines-10-11-highlighted-line-12-excluded.png`: the selected Code Inspector renders source, highlights lines 10 and 11, and leaves line 12 unhighlighted.

### Explicit `Code not available in memory` decision

Decision: non-blocking; outside the P2-E-owned C10/C11 invariant; no acceptance follow-up required.

Evidence and reasoning:

- The phrase is visible only in the AI citation card in screenshots 05/06.
- `CodeReferencesPanel.refsWithSnippets` deliberately assigns every citation-card entry `content: null`; the card therefore renders the phrase by design. It is not a failed backend read status.
- C10 owns fail-closed uniqueness and preservation of the exact persisted ID/path/range. The card shows exactly one correct FUNCTION reference with `src/unique.ts` and `L10-11`; ambiguity generated none.
- Clicking `Focus in graph` selects the exact persisted node. The selected-viewer path separately calls `readFile`.
- Raw network evidence records the successful exact request `/api/file?path=src%2Funique.ts&repo=qa-p2e&startLine=0&endLine=60`, with no failed response or request failure.
- Screenshot 06 proves that the selected viewer received and rendered the file. Its line styles are cyan for 10/11 and transparent for 12, and recorded scrolling targeted one-based line 10.
- Therefore the citation phrase means only that the citation card does not own hydrated snippets. It neither contradicts the C10 persisted-reference contract nor the C11 file-read/range/highlight contract.

## E2-P2E-REPEAT1: C04/C17 and M1-M7

Decision: PASS, repeated-analyze matrix `7/7`.

The P2-E wrapper verifies the accepted P2-D harness hash before transforming only lane names/paths. The full accepted harness and raw P2-E repeat evidence were read. The harness predicates do not accept aggregate counts alone: they compare input manifests, normalized Definition records including IDs and all construct/selection coordinates, exact `DEFINES` records, endpoint closure, current native reads, owned failure responses, and recovery facts.

| Matrix row | Decision | Independent reconciliation |
|---|---|---|
| M1 unchanged x2 | PASS | same fixture path and input signature; exits `0/0`; `10/10` Definitions and `10/10` exact pairs; zero missing endpoints; normalized fact signatures equal |
| M2 changed input | PASS | source signature changes; `later` appears; `now` reduces to one; the old second-`now` opaque ID is absent; native current read returns `later, now, time, time` |
| M3 analyze failure | PASS | making the owned `.anvien` path a regular file returns exit `1`, concrete mkdir error, and no graph at the expected success path |
| M4 native current read | PASS | `graph.json` is absent for the complete MCP call, Ladybug is present, and current changed facts return successfully |
| M5 no readable backend | PASS | both same-repository backends are absent; MCP host exits normally but returns JSON-RPC `-32602`, no `result`, and no stale rows |
| M6 recovery | PASS | source/artifacts restored; normal analyze exits `0`; baseline exact signature and `now, now, time, time` native rows return |
| M7 exact facts | PASS | exact IDs/ranges/selection ranges, `10/10` pairs, and zero missing endpoints close the accepted semantic unit |

The two unchanged Graph JSON hashes are identical. Whole Ladybug hashes differ, but whole-container byte identity is explicitly not the determinism contract. Exact accepted semantic identities, values, ranges, and endpoint pairs are stable. This classification matches the contract and accepted P2-D decision.

`Server.runCypherRead` is Ladybug-primary and falls back to same-repository Graph JSON only on `errors.Is(openErr, lbugnative.ErrUnavailable)`. The normal packaged runtime is built with `-tags ladybugdb`; deleting/corrupting a runtime artifact yields a non-sentinel operational error, not the compile-time capability sentinel defined by `!ladybugdb`. The run truthfully proves native success, non-sentinel fail-closed behavior, no stale result, and recovery. It does not manufacture an untagged/dev binary or fake sentinel. This is sufficient and non-blocking, consistently with the accepted P2-D authority.

## Data-Integrity and Edge-Case Closure

The Data-Integrity lens is clear:

- record identity and every compared field are conserved by opaque ID;
- absent optional coordinates remain NULL and real zero remains data;
- partial optional ranges are rejected/absent, never silently normalized;
- exact `DEFINES` pair multiplicity and endpoint existence are conserved;
- explicit embedding label persistence/readback/hydration is dynamically proven without a false normal-artifact claim;
- no silent record drop, alternate artifact, or cross-repository substitution exists.

The Edge-Case lens is clear:

- exact identity accepts the complete ID and rejects suffix/near-match values;
- same-name ambiguity fails closed;
- missing embedding label fails closed rather than parsing opaque identity;
- unchanged/change ordering exposes current facts and removes stale identity;
- analyze-storage and no-backend faults return clear non-success;
- recovery is exact and leaves no held fault artifact;
- no partial success remains armed for a later invocation.

## Containment and Artifact Lifecycle

Final verified containment before report creation:

- fault artifacts restored and held paths remaining `0`;
- runtime stopped through manifest-owned process identities only;
- listeners on ports `4848` and `5228`: `0/0`;
- `.tmp/qa-child02-p2e` absent;
- shared `.tmp/ladybug-home`, `.tmp/ladybug-native`, and `.tmp/runtime-p2c` still present;
- five protected plan documents unchanged by hash;
- staged paths `0`;
- no production, existing product-test, target, scanner, contract, or generated-file edit;
- no failed, retry, duplicate, or superseded P2-E evidence artifact remains;
- official evidence consists of one build pair, one parity pair, one reader-command pair, one repeat pair, one consolidated pair, one Playwright pair, six screenshots, three reusable harnesses, and one QA report.

## Not Run

- No full build, parity, reader suite, Playwright run, repeat matrix, or accepted predecessor gate was rerun. Current HEAD/source, raw evidence, current matching binary hash, and exact artifact hashes supplied sufficient evidence; no invalidation was found.
- No Anvien graph command was needed. Scope and exact source owners were already established by accepted inventory/current source, and the assigned review expressly prohibited `detect-changes`.
- No QA skill was used to recreate the matrix.
- No private orchestration attachment was requested or read.
- No cleanup, stage, commit, push, branch switch, Pn-A/Pn-B/Pn-C opening, or Child 03 opening was performed.

## Exact Handoff to Orchestration/Main

1. Accept this unconditional P2-E `PASS` as `E2-P2E-REVIEW1` with disposition `ACCEPT_P2E_AND_HAND_BACK`.
2. Record the exact accepted P2-E facts in only the roadmap and four Child 02 ledgers: build `1/1`, persistence `8/8`, readers `8/8`, repeat `7/7`, C01-C17 `17/17`, citation-card decision, browser/SwiftShader classification, hashes, and containment.
3. Preserve all 22 P2-E artifacts and this Supervisor report without rewriting their evidence content.
4. Under orchestration/main authority, run the required final graph refresh and `anvien detect-changes --repo Anvien --scope all` before the implementation/validation commit, record `E2-P2E-DETECT1`, stage only the exact P2-E closure boundary, and create the isolated P2-E commit as `E2-P2E-COMMIT1`.
5. Keep Pn-A, Pn-B, Pn-C, Child 03, and every later Child closed until Git confirms the isolated P2-E commit. The next allowed slice transition is an orchestration decision after that commit, not an action authorized by this report.

## Overall Evaluation

P2-E closes the stated validation-only contract at current production HEAD. The evidence is current, internally correlated, field-level rather than count-only, source-backed, visually inspected, failure-aware, and contained. No acceptance proof is missing or contradicted. The visible citation-card phrase is an intentional no-snippet state that does not impair exact grounding or selected file reading and is therefore not a reader defect within this slice.

Final verdict: PASS for Child 02 P2-E only.
