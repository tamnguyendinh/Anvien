# Web UI Full Analyze Contract Re-review Approval

Timestamp: 2026-05-08 18:33:25 UTC+7
Lane: supervisor
Scope: `docs/plans/2026-05-08-web-ui-full-analyze-contract-plan.md`
Reviewed commit: `6b32755` (`fix: giu dung repo path cho web ui analyze`)
Basis: plan + current source only. No SPEC, `ARCHITECTURE.md`, or other architecture docs used as authority.
Verdict: PASS

## Summary

The three blockers from `reports/Supervisor/rp_supervisor_260508_181008_by_gpt-5_web_ui_full_analyze_contract_rereview_reject.md` are closed in current source.

The selected-path chain in the plan is now preserved for the reviewed paths:

```text
selected repoPath -> analyze repoPath -> complete repoPath -> load graph repoPath -> render graph same repoPath
```

No CRITICAL/HIGH/MEDIUM blocking issue found in this re-review.

## Prior Blocker Closure

### Closed: absolute-path resolver retry no longer falls back to basename/name

File: `anvien/src/server/api.ts:91`

Clearance: `findRepoAfterRegistryRefresh` now resolves an absolute `repoName` by `samePath(path.resolve(repo.repoPath), requestedPath)` before name/id matching at `anvien/src/server/api.ts:101` and `anvien/src/server/api.ts:103`. Runtime refresh paths use this helper after active analyze completion and after emergency backend refresh at `anvien/src/server/api.ts:361` and `anvien/src/server/api.ts:378`. This closes the same-name/same-basename redirect bug identified in the prior report.

### Closed: graph load now uses canonical `repoPath` from repo hydration

File: `anvien-web/src/services/backend-client.ts:811`

Clearance: `connectToServer` hydrates repo info, derives `graphRepo = repoInfo.repoPath ?? repoInfo.path ?? repoName` at `anvien-web/src/services/backend-client.ts:812`, then passes `graphRepo` to `fetchGraph` at `anvien-web/src/services/backend-client.ts:815`. This means a display name or ambiguous original input cannot decide the post-hydration graph target when the backend returns a canonical path.

### Closed: header re-analyze failures are visible and path-specific

File: `anvien-web/src/components/Header.tsx:71`

Clearance: header state now tracks `reanalyzeError` with `repoName`, `repoPath`, and `message` at `anvien-web/src/components/Header.tsx:71`. SSE failure logs `{ repoPath, error }` and stores a visible error at `anvien-web/src/components/Header.tsx:327` and `anvien-web/src/components/Header.tsx:331`. Start failure does the same at `anvien-web/src/components/Header.tsx:346` and `anvien-web/src/components/Header.tsx:350`. The dropdown renders the path-specific `role="alert"` at `anvien-web/src/components/Header.tsx:429` through `anvien-web/src/components/Header.tsx:440`.

## Additional Source Clearance

- Backend full-analyze contract remains intact: Web/API analyze worker options force full analyze at `anvien/src/server/api.ts:67` and worker start sends those options at `anvien/src/server/api.ts:1206`.
- Analyze completion payload still carries `repoPath`: `buildAnalyzeCompleteEventPayload` includes `repoPath` at `anvien/src/server/api.ts:83`, and SSE terminal events use it at `anvien/src/server/api.ts:177` and `anvien/src/server/api.ts:202`.
- Header re-analyze success path still prefers completed path: `onAnalyzeComplete?.(data.repoPath ?? repoPath)` at `anvien-web/src/components/Header.tsx:323`.
- Regression coverage was added for the reviewed fixes: backend duplicate-name refresh at `anvien/test/unit/analyze-api.test.ts:69`, canonical graph load at `anvien-web/test/unit/server-connection.test.ts:78`, SSE failure UI/logging at `anvien-web/test/unit/Header.reanalyze-flow.test.tsx:113`, and start failure UI/logging at `anvien-web/test/unit/Header.reanalyze-flow.test.tsx:156`.

## Verification Run

Passed:

```powershell
cd anvien
npx vitest run test/unit/analyze-api.test.ts test/unit/repo-resolver.test.ts
```

Result: 2 files passed, 15 tests passed.

Passed:

```powershell
cd anvien-web
npm test -- test/unit/DropZone.full-analyze-flow.test.tsx test/unit/Header.reanalyze-flow.test.tsx test/unit/RepoAnalyzer.local-only.test.tsx test/unit/repo-list.test.ts test/unit/server-connection.test.ts
```

Result: 5 files passed, 33 tests passed.

Passed:

```powershell
cd anvien-web
npx tsc --noEmit
```

Passed:

```powershell
cd anvien
npx tsc --noEmit
```

Passed:

```powershell
powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1
```

Build completed. Vite emitted only bundle-size/dynamic-import warnings; no build failure.

## Notes

I did not rerun the browser full-analyze smoke in this turn because it would trigger a full analyze. This report approves closure of the previously rejected source/test/build blockers against the plan.
