# Coder Report: File Group Post-Review Fix

Status: READY FOR SUPERVISOR REVIEW

## Metadata
- Report file: `reports/coder/rp_coder_260605_234312_by_gpt-5-codex_file_group_post_review_fix.md`
- Time: 260605 234312 Asia/Bangkok
- Coder: gpt-5-codex
- Scope: close Supervisor report `reports/Supervisor/rp_supervisor_260605_232104_by_gpt-5-codex_file-role-classification-gap-review.md`
- Authority: `AGENTS.md`, `.agents/skills/coder/SKILL.md`, `docs/plans/2026-06-03-anvien-file-role-classification-gap/*`, current repo evidence
- Implementation checkpoint: to be created after this report; hash will be recorded in a follow-up ledger entry

## Invariant Family Map
- Family name: first-class backend file identity group `backend_support_model_helper`.
- SSOT / authority source: backend semantic identity in `internal/semantic`, file projection in `internal/filecontext`, plan/evidence/benchmark ledgers under `docs/plans/2026-06-03-anvien-file-role-classification-gap`.
- Sibling runtime surfaces checked: graph File-node enrichment, file projection/file-context output, CLI analyze output, Web generated contract consumers, FileMap/FileDetail UI tests, Playwright e2e entrypoint, plan/evidence/benchmark ledgers.
- Forbidden fallback / alternate path: no Web-only path-pattern inference and no display-only group classifier.
- Stale tests/helpers/plans updated: `anvien-web/playwright.config.ts` now makes the recorded e2e command self-contained; evidence and benchmark wording now separates historical E11 counts from current E12 counts.
- Verify matrix: analyze graph enrichment, focused Go tests, Web unit tests, Playwright e2e, `file-context`, `file-hotspots`, direct `.anvien/graph.json` inspection, `detect-changes`.

## Files Changed
- `internal/semantic/app_layer.go`: `semantic.Apply` now persists `fileRole` and `fileGroup` on graph `File` nodes using the backend file identity classifier.
- `internal/semantic/app_layer_test.go`: added `TestApplyPersistsFileGroupOnFileNodes`.
- `anvien-web/playwright.config.ts`: added Playwright `webServer` for the existing Vite dev server on `127.0.0.1:5228`.
- `docs/plans/2026-06-03-anvien-file-role-classification-gap/*`: refreshed E12 evidence, benchmark counts, and closure notes.
- `reports/Supervisor/rp_supervisor_260605_232104_by_gpt-5-codex_file-role-classification-gap-review.md`: recorded the reject that drove this fix.

## Verify Outputs
| Command | Result |
|---|---|
| Full build sequence from repo root | Pass. Version remained `1.2.5`; final analyze passed with `files.scanned=1383`, `parsed_code=682`, `failed=0`, `nodes=84192`, `relationships=122663`, `dependencyEdges=16569`. |
| `go test ./internal/semantic ./internal/filecontext ./internal/contracts ./internal/cli ./internal/httpapi ./internal/mcp -count=1` | Pass. |
| `npm test -- FileMapPanel.test.tsx FileDetailPanel.test.tsx` | Pass. 2 files, 8 tests. |
| `npm run test:e2e -- file-map-test-unresolved.spec.ts` | Pass. 1 Chromium test; Playwright starts the frontend server through `webServer`. |
| `anvien file-context internal/repo/runtime_config.go --repo Anvien --json` | Pass. Reports `fileRole=config` and `fileGroup=backend_support_model_helper`. |
| Direct `.anvien/graph.json` inspection | Pass. `File:internal/repo/runtime_config.go` has `fileGroup=backend_support_model_helper` and `fileRole=config`. |
| `anvien file-hotspots --repo Anvien --json --sort path --limit 0` | Pass. `backend_support_model_helper` has `files=47`, `unresolved=2531`, and current role breakdown recorded in E12/B6. |
| `anvien detect-changes --repo Anvien --scope all` | Pass. Summary risk `low`; changed files limited to implementation, e2e config, and plan ledgers. |

## E2E Flow
- Trigger: `npm run test:e2e -- file-map-test-unresolved.spec.ts`.
- Process: Playwright starts `npm run dev` on `127.0.0.1:5228`, opens the File Map scenario, and reads backend-provided group labels.
- Observable result: Chromium e2e passes with the recorded command and no pre-existing frontend server.

## Closure Notes
- Supervisor finding 1 is closed by self-contained Playwright server startup.
- Supervisor finding 2 is closed by graph File-node `fileGroup` persistence in semantic enrichment.
- Supervisor finding 3 is closed by refreshing current benchmark evidence from `42` historical closure files to `47` current files.
- Residual unverified surfaces: none.
- Risks/open points: direct Supervisor acceptance is still required before this scope is DONE.
