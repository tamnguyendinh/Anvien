# Web UI Full Analyze Contract Plan

Date: 2026-05-08

Status: Reopened after real UI review.

## Goal

Make the Web UI behave as a visual control surface for the same full analyze behavior users get from the CLI full analyze path. The Web UI should not introduce a separate analyze meaning, shortcut, mode, or cached-graph interpretation for actions that users understand as analyzing a repository.

The user-facing contract is:

- Clicking a repository card runs a full analyze for that repository, then loads the newly generated graph.
- Analyzing by local path runs a full analyze, then loads the newly generated graph.
- Re-analyze runs a full analyze, then loads the newly generated graph.

The stronger technical contract is:

```text
selected repoPath
  -> analyze repoPath
  -> complete repoPath
  -> load graph repoPath
  -> render graph for the same repoPath
```

`repoName` is display metadata. It must not be the primary identity for post-analyze graph loading, because name-only routing can drift from the selected repository path or become ambiguous when multiple repos share a basename/name.

## Non-Goals

- Do not change CLI analyze semantics.
- Do not change the core analyze pipeline.
- Do not add new analyze modes such as `full` versus `ifChanged`.
- Do not add benchmark work for this fix.
- Do not compare graph-load time with analyze time.
- Do not make cached graph loading the default repo-card behavior.

## Problem

The current Web UI has multiple paths that can be interpreted as analyze:

- Repo-card click can connect to an already indexed repository and load the stored graph.
- Analyze by path can call `/api/analyze` without forcing a rebuild, allowing the up-to-date shortcut.
- Header re-analyze currently forces a rebuild.

That creates contract drift. The CLI full analyze path rebuilds the graph, while some Web UI paths can return quickly because they only load an existing graph or take the non-force shortcut. This makes the UI easy to misread and can show stale graph data after source code changes.

Re-review finding:

- The first implementation fixed only the narrow ordering rule: repo-card click starts analyze before graph loading.
- It did not lock the selected repository identity through the whole flow.
- Landing repo-card analyze starts from `repo.path`, but the completion path can fall back to `repoName`.
- The backend SSE complete payload currently carries `repoName`, not the analyzed `repoPath`.
- Some graph-load calls use `awaitAnalysis: true`; the landing post-analyze graph load does not.
- The existing tests mostly prove call order with mocks. They do not prove that the repo being analyzed and the repo being loaded are the same physical path.
- Therefore the plan was marked complete too early.

## Correct Contract

The Web UI is only the button/progress/graph layer over the analyze runtime.

It should do this:

1. User clicks a Web UI analyze entry point.
2. Web UI starts analyze for the selected repository/path.
3. Backend runs full analyze using the existing analyze implementation.
4. Web UI streams analyze progress.
5. After analyze completes, Web UI loads the graph produced by that analyze.
6. Web UI renders that graph.

It should not do this for any default repo/open/analyze action:

1. User clicks a repository card.
2. Web UI loads the existing graph directly.
3. User sees stale graph data while believing analyze just ran.

It also must not do this:

1. User selects repo B.
2. Web UI starts analyze for repo B by path.
3. Analyze completes.
4. Web UI loads graph by name only, or by a stale/default/current repo binding.
5. User sees repo A, stale repo B, or a LadybugDB missing error for the wrong path/timing.

All analyze entry points must preserve one repository identity object or canonical path through the whole chain.

## Flow Family To Keep In Sync

These flows are one family and must be reviewed together whenever this contract changes:

- Landing repo card selection.
- Landing "Analyze Another Repository" by path.
- Header repository dropdown selection.
- Header re-analyze button.
- Header "Analyze a new repository..." sheet.
- Auto-connect from `?project=...`.
- Backend `/api/analyze` job creation, worker execution, SSE completion.
- Backend `/api/repo` and `/api/graph` repo resolution.
- Frontend repo list refresh after analyze.
- Browser reload/back-to-landing behavior after failures.

## Implementation Plan

- [x] Keep CLI behavior unchanged.
- [x] Keep `runFullAnalysis` and the pipeline behavior unchanged unless a direct Web API wrapper bug requires a minimal fix.
- [x] Make backend `/api/analyze` always start a full analyze for Web/API analyze requests by passing `force: true` to the analyze worker.
- [x] Remove Web UI reliance on client-provided `force` for analyze semantics.
- [x] Keep any `force` request field only as legacy-tolerated input if needed, but do not let it make Web analyze non-full.
- [x] Update local-path analyze flow to call `/api/analyze`, wait for completion, then load the newly generated graph.
- [x] Update header re-analyze flow to call the same analyze path without adding special semantics beyond full analyze.
- [x] Update repo-card click flow so it starts full analyze for `repo.path`, streams progress, and only connects/loads graph after completion.
- [x] Ensure repo-card click no longer calls graph load directly as the primary action.
- [x] Keep graph loading as a post-analyze rendering step, not the semantic action.
- [x] Adjust user-visible progress wording so analyze and graph loading are distinct steps.
- [x] Add or update backend tests proving `/api/analyze` sends worker options with `force: true` even when the request omits `force` or sends `force: false`.
- [x] Add or update Web tests proving repo-card click starts analyze before graph loading.
- [x] Add or update Web tests proving path analyze starts analyze and then loads the completed repo graph.
- [x] Add or update Web tests proving re-analyze uses the same full-analyze flow.
- [x] Run full launcher build before tests.
- [x] Run targeted backend and Web tests for the analyze flows.
- [x] Run broader relevant test suites if targeted tests pass.

## Reopened Implementation Plan

- [ ] Treat `repoPath` as the primary identity for post-analyze loading.
- [ ] Change landing repo-card selection to pass the selected repo object or canonical `repoPath`, not only `repo.name`.
- [ ] Make analyze completion carry the analyzed `repoPath` from backend job to SSE complete payload.
- [ ] Update frontend analyze completion handlers to load graph by the selected/completed `repoPath`, not by `data.repoName ?? repo.name`.
- [ ] Make landing post-analyze `connectToServer` use the same hold-queue behavior as auto-connect/header flows when loading a just-analyzed graph.
- [ ] Verify `connectToServer`, `fetchRepoInfo`, and `fetchGraph` can route by full path without path basename normalization breaking the lookup.
- [ ] Ensure header "Analyze a new repository..." refreshes the repo list and adds/selects the newly analyzed repo by path.
- [ ] Ensure header re-analyze and landing repo-card share the same path-preserving load-after-analyze helper.
- [ ] Keep header repo dropdown switch behavior intentionally separate if it remains a fast graph switch, and label/implement it as such.
- [ ] Add regression tests where current/default repo A exists and selected repo B is analyzed; assert all API calls after completion target repo B's path.
- [ ] Add regression tests for duplicate repo names/basenames so name-only routing cannot pass.
- [ ] Add tests for landing repo-card, local-path analyze, header re-analyze, and header analyze-new using the same identity contract.
- [ ] Add tests proving SSE complete includes `repoPath` and UI prefers it over `repoName`.
- [ ] Run full launcher build before browser validation.
- [ ] Validate manually in a real browser with a clean session/cache:
  - click landing repo card for repo B;
  - confirm POST `/api/analyze` body path is repo B;
  - confirm SSE complete includes repo B path;
  - confirm `/api/repo` and `/api/graph` load repo B;
  - confirm repo B indexed timestamp changes;
  - confirm UI renders repo B graph and no related UI, console, or network error appears.
- [ ] Repeat browser validation for:
  - landing "Analyze Another Repository";
  - header re-analyze;
  - header "Analyze a new repository...";
  - header dropdown repo switch if it is intentionally load-only.
- [ ] Only mark this plan complete after real browser validation and tests both prove the same repoPath contract.

## Acceptance Criteria

- CLI analyze behavior is unchanged.
- Web UI repo-card click cannot show an old graph without first running analyze.
- Web UI path analyze cannot early-return as up-to-date through a non-force request.
- Web UI re-analyze remains a full analyze.
- Graph loading is visibly and structurally a post-analyze step.
- Tests lock the Web/API contract so future edits cannot reintroduce cached-load-as-analyze behavior.

Additional reopened acceptance criteria:

- The physical repository path selected by the user is the same path analyzed by the backend and the same path used to load graph data.
- A repo name returned by analyze completion cannot redirect graph loading away from the selected path.
- Landing repo-card flow no longer has weaker post-analyze loading semantics than header/auto-connect flows.
- Real browser validation confirms the complete flow, not just mocked unit tests.

## Validation Log

- `avmatrix-launcher\build.ps1` passed before test validation.
- `avmatrix/test/unit/analyze-api.test.ts` passed.
- Web targeted analyze-flow tests passed:
  - `analyze-contract.local-only.test.tsx`
  - `RepoAnalyzer.local-only.test.tsx`
  - `DropZone.full-analyze-flow.test.tsx`
  - `Header.reanalyze-flow.test.tsx`
  - `Branding.local-only.test.tsx`
- `avmatrix npm test` passed.
- `avmatrix-web npm test` passed.

## Re-review Validation Gap

The previous validation was insufficient because it did not include browser-level confirmation of repository identity across analyze completion and graph load. The next validation must include real UI/browser execution and network/request inspection for the selected repo path.
