# QA Report — Child 02 P2-C affected readers

Verdict: **BLOCKED** for the exact `8`-row P2-C QA denominator.

Disposition: **7 PASS / 8 total; C16 is BLOCKED/UNVERIFIED only at the fully native vector-index boundary because the required official Ladybug vector extension is unavailable.** This is an independent QA verdict, not Supervisor acceptance. No production defect was reproduced.

## Scope and authority

- Role: independent QA for Child 02 P2-C only.
- Affected-reader denominator: exactly C09-C16 = `8`.
- Production/test authority: read-only. No production source or existing test was edited by QA.
- QA writes: one reusable script under `playwright/`, paired official evidence plus screenshots under the repository-canonical `reports/QA/playwright/child02-p2c/` path, and this sole durable QA report.
- Explicitly not run: Anvien `detect-changes`, stage, commit, P2-D, P2-E, closure gates, or accepted P2-A/P2-B gates without invalidation.
- No acceptance claim: QA evidence is handed back for Orchestration/Supervisor decision.

## Git boundary

### Before QA

- Repository: `E:\Anvien`
- Branch: `master`
- HEAD/baseline: `4d456446fcc49aed0c6d489aa9c63e00d030b53c`
- Staged paths: `0`
- Dirty boundary: exactly the nine delegated P2-C paths:
  1. `anvien-web/src/hooks/useAppState.local-runtime.tsx`
  2. `anvien-web/src/components/CodeReferencesPanel.tsx`
  3. `internal/embeddings/search.go`
  4. `anvien-web/test/unit/useAppState.local-runtime.test.tsx`
  5. `anvien-web/test/unit/ChatPanel.grounding-links.test.tsx`
  6. `anvien-web/test/unit/CodeReferencesPanel.graph-health.test.tsx`
  7. `internal/embeddings/search_test.go`
  8. `internal/embeddings/search_ladybugdb_test.go`
  9. `reports/coder/rp_coder_260813_185027_by_gpt-5_child02_p2c_affected_readers.md`
- Coder-report SHA-256 independently matched the handoff: `DC36748AE81F4209F65D03E4F5E3FBE69099E2B4EC51E9E28CD8CE5C7C6B1F4A`.

### After QA

- Branch/HEAD unchanged: `master` / `4d456446fcc49aed0c6d489aa9c63e00d030b53c`.
- Staged paths: `0`.
- Original nine-path dirty boundary preserved.
- QA-added paths only:
  - `playwright/child02-p2c-affected-readers.mjs`
  - `reports/QA/playwright/child02-p2c/child02-p2c-affected-readers.json`
  - `reports/QA/playwright/child02-p2c/child02-p2c-affected-readers.md`
  - six screenshots under `reports/QA/playwright/child02-p2c/screenshots/`
  - this report.
- Validate-only source paths C12-C15 have zero worktree diff.
- `git diff --check`: PASS.
- No stage, commit, or detect-changes operation was performed.

## Authority and source clearance

Read in full before conclusions:

- `AGENTS.md`;
- mandatory skills: `working-rules`, `qa`, `Edge-Case`, and `control-in-app-browser`;
- graph-accuracy roadmap and all four Child 02 ledgers;
- `docs/contracts/graph-accuracy-contract.md`;
- accepted P2-A and P2-B Supervisor reports;
- current P2-C coder report;
- complete current production/test diff and new native test;
- exact graph/file API/session stream/runtime routes and C09-C16 owners.

Source/diff clearance:

- Production diff is confined to the exact edit rows C09/C10, C11, and C16 in three files.
- Focused test diff is confined to four existing test files plus one new build-tagged native test.
- C12-C15 files are unchanged.
- No plan, roadmap, ledger, accepted report, skill/rule, target source, Git index, or P2-D/P2-E owner changed.
- The approved C11 layout, copy, component hierarchy, styles, and interaction structure remain unchanged; the source diff is coordinate handling only.

## Anvien graph and blast radius

`anvien --help` was run first. `anvien analyze --force` then refreshed graph evidence before graph reads:

- exit `0`, `59.7s`;
- scanned `1,596`, parsed code `685`, failed `0`;
- graph `96,284` nodes / `135,403` relationships;
- all seven affected/validate-only file-detail targets reported indexed/current commit `4d456446...`, `stale=false`, `changedSinceAnalyze=false`.

Current exact-symbol impact evidence:

| Reader owner | Exact risk / impacted | Host/file warning | QA handling |
| --- | ---: | ---: | --- |
| C09 `resolveNodeIds` | LOW / `0` | HIGH host file | mounted tool-result regression and exact/suffix/near-match visual proof |
| C10 nested `handleNodeGroundingReference` | exact symbol unresolved by current graph | HIGH host file | full source trace plus mounted ambiguity/unique flow |
| C11 `CodeReferencesPanel` | LOW / `3`, direct `2` | HIGH host file | built UI, official Playwright, screenshot and file API proof |
| C16 `SemanticSearch` | MEDIUM / `12`, direct `6`, modules `2` | HIGH host file | package, native persistence/hydration, failure paths, extension boundary |
| C16 `hydrateSearchResults` / query helpers | LOW / `8-12` | HIGH host file | focused and native tests |
| C16 `ChunkSearchRow` / `BestChunkMatch` | MEDIUM / `18` and `17` | HIGH host file | explicit-label propagation/dedup coverage |
| C12 nested `nodeRange` | exact symbol unresolved by current graph | HIGH host file | unchanged-source proof plus filecontext package |
| C13 payload builders | CRITICAL / `4` and `1`; processes `34` and `7` | HIGH host file | unchanged-source proof plus MCP context regressions |
| C14 `detectChangedSymbols` | HIGH / `1`; processes `4` | HIGH host file | unchanged-source proof plus MCP detect regressions |
| C15 `collectRenameChanges` | CRITICAL / `1`; processes `8` | HIGH host file | unchanged-source proof plus MCP rename regressions |

No returned impact contained a degraded node or active ResolutionGap. HIGH/CRITICAL values are reported as blast-radius warnings and were covered by unchanged-source checks plus nearest package regressions; they are not edit prohibitions.

## Clean-build holder chronology

The root full build was run only after exact build artifacts were inspected with Windows Restart Manager.

Artifacts queried:

- `E:\Anvien\anvien\bin\anvien.exe`
- `C:\Users\TAM NGUYEN\AppData\Roaming\npm\node_modules\anvien\bin\anvien.exe`
- `E:\Anvien\anvien-launcher\AnvienLauncher.exe`
- `E:\Anvien\anvien-launcher\server-bundle\anvien-server.exe`

Initial holder proof:

- Restart Manager identified PID `20732`, `anvien.exe`, command `...\npm\node_modules\anvien\bin\anvien.exe mcp`, parent launcher PID `20744`, `cmd.exe ... anvien.cmd mcp`.
- Restart Manager identified PID `16696`, same global MCP binary, parent launcher PID `17936`, `cmd.exe ... anvien.cmd mcp`.
- Both holders were registered against the canonical and global CLI artifacts.
- Launcher and server-wrapper artifacts had zero holders.
- Only the two proven holders and their proven launchers were stopped. The first stop command ended after parent `20744` had already exited; follow-up verification/cleanup confirmed all four PIDs dead.
- Unrelated MCP pairs, VS Code Playwright test-server, Codex/browser runtime, and processes not holding an artifact were not stopped.

Pre-build cleanliness:

- Restart Manager: zero holders for all four artifacts.
- Exclusive-open probe: PASS for all four artifacts.

Fresh full build:

```text
npm run full-build
```

- exit `0` in `208.146s`;
- built packaged Go/Ladybug runtime, canonical/global CLI `1.2.8`, launcher, server wrapper, TypeScript, and Vite production bundle;
- built `anvien analyze . --force` PASS: `1,596/685/0`, graph `96,284/135,403`;
- post-build Restart Manager: zero holders for all four artifacts;
- only existing npm/Vite chunk/dynamic-import warnings remained.

This proves the current production/test worktree builds from a clean artifact state and produces the built runtime used below.

## Command-by-command validation

| Command/boundary | Result | What it proves |
| --- | --- | --- |
| focused frontend Vitest, three named files | PASS, `3` files / `11/11` tests, `33.237s` | C09 exact/suffix, C10 ambiguous/unique coordinates, C11 range/UI unit support |
| `go test ./internal/embeddings -count=1 -v` | PASS, `6.463s` | C16 label selection/propagation/dedup/hydration; missing label and metadata failure return errors |
| tagged native C16 test with repo v0.19.1 Ladybug runtime | PASS, test `0.53s`, command `8.839s` | real CodeEmbedding label persistence/readback and native metadata hydration for opaque same-name rows |
| `go test ./internal/filecontext -count=1 -v` | PASS, `2.052s` | nearest C12 file-context/range behavior and determinism |
| `go test ./internal/mcp -count=1 -v` | PASS, `7.447s` | nearest C13 context, C14 detect, C15 rename regressions |
| built backend `/api/info` | HTTP `200`, version `1.2.8`, Go `1.26.3` | built backend listener health |
| built Vite preview `/` | HTTP `200` | freshly built frontend served by production preview, not a dev server |
| BrowserPanel visible start/repository screens | PASS | real user entry, `Start Anvien`, current Anvien card `1,596 / 96,284 / 662`; zero console warning/error |
| reusable official Playwright script | final PASS | mounted production C09-C11 flows with controlled external fixture/event inputs |
| one official native `INSTALL VECTOR` boundary attempt | BLOCKED | exact v0.19.0 extension endpoint unavailable; no retry/substitution |
| validate-only source diff check | PASS | C12-C15 source remained unchanged |
| final JSON parse, hashes, `git diff --check`, staged/status/ports/temp checks | PASS | evidence integrity and final containment |

## Visible runtime and official Playwright evidence

Mounted runtime:

- backend: `E:\Anvien\anvien\bin\anvien.exe serve --host 127.0.0.1 --port 4848`
- frontend: `npm run preview` from `anvien-web`, strict port `5228`
- browser-first gate: Codex in-app BrowserPanel opened the built start screen and repository-selection screen before standalone Playwright automation.

Official reusable script:

- `playwright/child02-p2c-affected-readers.mjs`
- SHA-256: `276B4D6EA54E97ED6E99A945BB13368685FE4FDC3CB4C5A6540BB82F8F3DA058`

Paired evidence:

- `reports/QA/playwright/child02-p2c/child02-p2c-affected-readers.json`
  - SHA-256: `1D46E4DA2689BD91DC228E40636ABEAA4ECF56C867EA288779E4F6FE1348BF69`
- `reports/QA/playwright/child02-p2c/child02-p2c-affected-readers.md`
  - SHA-256: `0B917711B905F15CAB07712ECC8F2BE43F34C0874725C1C717E4B152F67E3FD3`

Controlled-boundary truth statement:

- Only external HTTP graph/session/file responses were deterministic fixtures.
- The built production bundle executed unchanged `AppStateProvider`, `ChatRuntimeContext`, `handleToolResult`, `resolveNodeIds`, `handleContentGrounding`, `handleNodeGroundingReference`, graph rendering, citation creation, selection, `CodeReferencesPanel`, and `readFile` client path.
- There was no DOM injection, test-only route, component mock, or replacement reader implementation.

Final evidence summary:

- C09-C11: `3/3 PASS`.
- screenshots: `6`.
- console warnings/errors: `0`.
- page errors: `0`.
- failed HTTP responses: `0`.
- request failures: `0`.

Screenshots visually inspected at original resolution:

1. `screenshots/01-mounted-production-fixture.png`
2. `screenshots/02-before-c09-tool-result.png`
3. `screenshots/03-c09-exact-id-only.png`
4. `screenshots/04-c10-ambiguous-fails-closed.png`
5. `screenshots/05-c10-unique-persisted-reference.png`
6. `screenshots/06-c11-lines-10-11-highlighted-line-12-excluded.png`

Visual result:

- C09: only `uniqueTarget` is cyan-active; suffix/near-match nodes remain dim; no red blast-radius node.
- C10 ambiguous: no citation and no Code Inspector opens.
- C10 unique: `FUNCTION`, `uniqueTarget`, `src/unique.ts`, and `L10–11` are visible.
- C11: viewer scroll target is line `10`; lines `10-11` are cyan highlighted; line `12` is unhighlighted; approved layout/copy/hierarchy remains intact with no observed overlap, overflow, missing font, or missing asset.

## C09-C16 verdict matrix

| Row | Verdict | Independent evidence | Residual |
| --- | --- | --- | --- |
| C09 `resolveNodeIds` | PASS | mounted production tool-result with exact `opaque-exact-alpha`, suffix `exact-alpha`, and near-match impact `suffix-alpha`; only exact node visually active | none in declared exact/suffix/fragment boundary |
| C10 grounding | PASS | duplicate same-label/same-name IDs produced citation count `0` and panel count `0`; unique persisted node produced exact path and `{10,0,12,0}` range | none in declared ambiguity/unique boundary |
| C11 panel range | PASS | visible `L10–11`; `/api/file` request `startLine=0,endLine=60`; line 10/11 cyan styles, line 12 transparent; recorded native scroll target `10` | none in declared range/UI boundary |
| C12 `nodeRange` | PASS | source unchanged; direct four-coordinate copy; filecontext package PASS | exact graph impact symbol resolution remained unavailable, handled as HIGH host warning |
| C13 context payloads | PASS | source unchanged; UID/path/line and ambiguity payload behavior; MCP context regressions PASS | none in declared validate-only boundary |
| C14 changed-symbol detection | PASS | source unchanged; direct persisted range/hunk intersection; MCP detect regressions PASS | none in declared validate-only boundary |
| C15 rename collection | PASS | source unchanged; UID/path/range and one-based-to-array conversion; MCP rename target/reference-line regressions PASS | none in declared validate-only boundary |
| C16 semantic search | **BLOCKED/UNVERIFIED** | package and native persistence/readback/metadata hydration PASS; explicit label flows vector row → chunk → dedup → hydration → metadata table; missing label/query failure fail clearly | full native vector-index creation/query unavailable because official extension install failed |

Denominator disposition: `7 PASS + 1 BLOCKED/UNVERIFIED = 8`. Therefore QA cannot report `8/8 PASS`.

## C16 native vector-extension boundary

Evidence that passed:

- `emb.label AS label` is selected by the production vector query.
- Label reaches `ChunkSearchRow`, nearest-chunk `BestChunkMatch`, hydration grouping, and metadata table selection.
- Opaque NodeID is not parsed for label meaning.
- Two same-name Function declarations with distinct opaque IDs/paths/ranges remain distinct, and a same-name Method hydrates through its own label/table.
- Missing label returns `semantic search match ... has no persisted label` before metadata query.
- Metadata query errors return a contextual error instead of silent skip.
- Native Ladybug writes/reads the real persisted `CodeEmbedding.label`, then production `SemanticSearch` performs native metadata hydration.

Exact availability/load attempt, once only:

```text
INSTALL VECTOR
```

- runtime: repo-bundled Ladybug native `v0.19.1`;
- requested official URL: `https://extension.ladybugdb.com/v0.19.0/win_amd64/vector/libvector.lbug_extension`;
- result: `IO exception: Failed to download extension: vector ... (ERROR: Failed to read connection)`;
- install called: `true`;
- load called: `false` because install failed;
- elapsed: `15.938s`;
- no retry, no loop, and no incompatible 0.15/0.16 substitution.

The coder handoff separately recorded HTTP `522` from the same endpoint. This QA run independently confirms the endpoint remains unavailable but did not receive a distinct HTTP status from the native client. Full native vector-index creation/query is therefore BLOCKED/UNVERIFIED, and this prevents `8/8` QA.

## Failures, corrections, and artifact lifecycle

- An initial combined skill read was display-truncated; QA stopped and reread the affected skill files sequentially through EOF before continuing.
- Initial exact-impact batch timed out; smaller complete batches replaced it without reducing the owner set.
- The first official Playwright run reached C09/C10 but timed out waiting for C11 lines because the harness had not performed the user-visible `Focus in graph` action. Production was not changed; the reusable script was corrected to click the real control.
- The next run passed all three reader assertions but reported one unlocated transient `404` console error. The script was strengthened to record every HTTP `>=400` response and request failure rather than filtering the error. The final run passed with zero console errors, zero failed responses, and zero request failures.
- Deterministic final evidence paths replaced the two superseded failed attempts; no duplicate/superseded official artifact remains.
- Session-created debug-only artifacts were exactly `.tmp/qa-child02-p2c/` (runtime logs and one-shot vector probe source/DB/WAL). A targeted `git clean -n -d -f -x -- .tmp/qa-child02-p2c` listed that directory alone; exact cleanup removed it and final existence check is false.
- Valid shared `.tmp/ladybug-native/v0.19.1` runtime and all final official evidence were preserved.
- Session-owned runtime listeners/processes were identified exactly: backend PID `6724`; frontend node `11132`, launcher cmd `21852`, npm node `15208`. Only those were stopped; final listener count on ports 4848/5228 is `0`.

## Final decision

**BLOCKED** for Child 02 P2-C QA as a whole. C09-C15 pass their declared runtime/nearest-real boundaries. C16 implementation-level and native persistence/metadata behavior passes, but the required fully native vector-index boundary remains unavailable at the official extension endpoint, so the exact affected-reader denominator is not `8/8 PASS`.

Handoff: Orchestration main — retain C09-C15 PASS, record C16 full native vector-index as BLOCKED/UNVERIFIED due official extension unavailability, and route the evidence to independent Supervisor without claiming acceptance.
