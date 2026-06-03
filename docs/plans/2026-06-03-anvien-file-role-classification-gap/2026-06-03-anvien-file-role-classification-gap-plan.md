# Anvien File Group Classification Plan

Date: 2026-06-03

Status: Detect-changes complete - implementation commit pending

Companion files:

- Evidence ledger: [2026-06-03-anvien-file-role-classification-gap-evidence.md](2026-06-03-anvien-file-role-classification-gap-evidence.md)
- Benchmark ledger: [2026-06-03-anvien-file-role-classification-gap-benchmark.md](2026-06-03-anvien-file-role-classification-gap-benchmark.md)

## Master Rules

1. Follow active workspace and repository instructions, including generated `AGENTS.md`.
2. Use Anvien for codebase analysis and impact checks before implementation edits.
3. Refresh graph evidence with `.\anvien\bin\anvien.exe analyze --force` before graph-based plan evidence or implementation evidence.
4. Run impact analysis before editing semantic classifiers, file projection summaries, graph-health/file-hotspot output, API contracts, graph builders, or Web file views.
5. Treat HIGH/CRITICAL Anvien impact as blast-radius evidence, not an automatic edit ban.
6. Keep generated `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**` as generated output only.
7. Run the full build before testing: `powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1`.
8. If Web UI behavior changes, include a relevant Web/e2e test; browser or screenshot validation can supplement e2e evidence but cannot replace it.
9. Record evidence as each evidenced task completes.
10. Record benchmarkable inventory/count changes as each benchmarkable task completes.
11. Run `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all` before implementation commits.
12. Update the corresponding checklist item immediately as each implementation slice completes.
13. Commit each completed implementation slice after evidence and benchmark ledgers are updated.

## Planner Rules

1. The plan file must keep the standard planner structure: metadata, rules, goal, problem, scope, requirements, invariants, technical direction, definition of done, checklist, and risk notes.
2. The phase checklist is not a short todo list. Every checklist item must be a complete mini-plan.
3. Each checklist item must include:
   - Goal: what the phase achieves.
   - Work Steps: concrete ordered work, including naming, grouping, implementation sequence, and validation sequence when relevant.
   - Implementation Gate: the condition that must be true before editing or moving forward.
   - Acceptance: the condition that proves the phase is done.
4. Do not hide essential implementation rules only in prose outside the checklist. If a rule is required to implement correctly, repeat it in the checklist item that uses it.
5. Do not rewrite the whole plan structure when a checklist is incomplete. Fix the checklist mini-plans.

## Goal

Create a first-class file group for backend support/model/helper files so Anvien displays the group directly instead of forcing users to infer it from unresolved count differences.

The target group is:

```text
fileGroup: backend_support_model_helper
label: Backend support/model/helper files
```

The group must be visible in graph/file projection/API/CLI/Web outputs as a file identity group.

## Problem

The previous implementation commit `444dcdd feat: add file role classification` added `fileRole`.

That was only role metadata. It did not put the files into an official file group.

The product bug is:

```text
Files Anvien can identify by real file type are not placed into a concrete file group that users can read directly.
```

The current user-visible failure mode is:

```text
unresolvedFiles=337
rawUnresolvedFiles=354
```

Users must infer that the difference is a known backend support/model/helper group. That is wrong. The file group must be named and shown directly:

```text
Backend support/model/helper files: 17
```

## Scope

In scope:

- Define a first-class `fileGroup` classification axis for file summaries and file projection.
- Add the concrete file group `backend_support_model_helper`.
- Attach the group to file metadata after `kind`, `appLayer`, `functionalArea`, and `fileRole` are known.
- Surface the group through `FileSummary`, CLI output, API/generated Web contracts, and Web file views.
- Add file projection group aggregation.
- Prove the current 17-file sample belongs to `backend_support_model_helper`.

Out of scope:

- Renaming `kind`, `appLayer`, `functionalArea`, or `fileRole`.
- Creating vague group names such as `classified`, `known`, `raw_audit`, `other`, or `misc`.
- Treating ResolutionGap/source-site classification as the file group name.
- Resolving Go builtins or standard library calls.
- Deleting source-site facts just to make counts smaller.

## Requirements

1. `fileGroup` must be a backend-owned classification field, not a Web-only label.
2. The new group key must be exactly `backend_support_model_helper`.
3. The display label must be exactly `Backend support/model/helper files`.
4. Group naming must follow this order:
   - object being grouped;
   - app side from `appLayer`;
   - purpose family from `fileRole`;
   - machine key as `<app_side>_<purpose_family>`;
   - display label as `<App Side> <purpose family words> files`.
5. `fileRole` remains a subcategory inside the file group.
6. The current 17 sample files must all have `fileGroup=backend_support_model_helper`.
7. Unknown-role files, frontend files, docs files, config-kind files, and generated files must not enter this group in this slice.
8. Analyze/CLI/Web output must show the group by name directly.
9. The group must not be discovered by subtracting `unresolvedFiles` from `rawUnresolvedFiles`.
10. Generated Web contracts must include the group field and group metadata.

## Invariants

1. File grouping starts from file identity axes: `kind`, `appLayer`, `functionalArea`, `fileRole`.
2. Source-site/ResolutionGap classifications are diagnostics, not file group names.
3. `backend_support_model_helper` is a file group, not a resolution bucket.
4. A file can retain raw unresolved counts while also belonging to a concrete file group.
5. Group membership must be deterministic and implemented once in shared backend logic.
6. CLI/API/Web consumers must read the backend-provided group; they must not reimplement path checks.

## Technical Direction

Likely owner areas:

- `internal/semantic/*` for semantic taxonomy and metadata definitions.
- `internal/filecontext/context.go` for `FileSummary`, file-context aggregation, file-hotspot summaries, and unresolved count fields.
- `internal/cli/command.go`, `internal/cli/file_context_command.go`, and graph-health CLI output for analyze/file projection display.
- `internal/contracts/web_ui.go` and generated Web contract files for API/Web type propagation.
- `anvien-web/src/components/FileMapPanel.tsx` and `anvien-web/src/components/FileDetailPanel.tsx` for Web display.

Before editing any of these areas, run the required Anvien impact checks named in the relevant checklist item.

## Definition Of Done

The corrected plan is complete when:

1. `backend_support_model_helper` exists as a first-class file group.
2. The current 17 sample files have `fileGroup=backend_support_model_helper`.
3. `FileSummary` and generated Web contracts include `fileGroup`.
4. Analyze output emits a direct group line for `backend_support_model_helper`.
5. CLI/Web display the group label directly.
6. Tests prove membership, boundary exclusions, generated contracts, CLI output, and Web rendering.
7. Full build, focused tests, Web/e2e tests when UI changes, analyze validation, graph-health/file-hotspots validation, and detect-changes evidence are recorded.
8. Evidence and benchmark ledgers record group count, role breakdown, validation, and commit hash.

## Phase Checklist

- [x] [G0-A] Audit planner compliance and current plan state.
  - Goal: make sure the plan itself follows the Anvien planner skill before implementation starts.
  - Work Steps:
    1. Read `.claude/skills/anvien/anvien-planner/SKILL.md`.
    2. Confirm the plan keeps metadata, rules, goal, problem, scope, requirements, invariants, technical direction, definition of done, phase checklist, and risk notes.
    3. Confirm every checklist item has Goal, Work Steps, Implementation Gate, and Acceptance.
    4. Confirm essential naming and grouping rules appear in the checklist items that use them, not only in prose above the checklist.
  - Implementation Gate: do not edit product code until this checklist audit passes.
  - Acceptance: evidence records the planner skill path read and confirms checklist items are complete mini-plans.

- [x] [G1-A] Audit current Anvien classification axes and target sample.
  - Goal: ground the file group name and membership rule in existing Anvien classification axes instead of inventing vague buckets.
  - Work Steps:
    1. Run `.\anvien\bin\anvien.exe analyze --force`.
    2. Use file projection output to record current file axes:
       - `kind`: `source`, `test`, `docs`, `config`, `generated`.
       - `appLayer`: `backend`, `backend_test`, `frontend`, `frontend_test`, `api`, `docs`, `config`, `generated_contract`, and other current values.
       - `functionalArea`: `providers`, `storage`, `analyzer`, `cli`, `query`, `resolution`, `session`, `mcp`, `web_graph_ui`, and other current values.
       - `fileRole`: `helper`, `config`, `parser_model`, `runtime_model`, `adapter`, `fallback_adapter`, `test_helper`, `analyzer_helper`, `contract_model`, `storage_helper`, `model`, `unknown`.
    3. Record current graph node axes as context only:
       - node `label`, node `appLayer`, node `functionalArea`.
    4. Record ResolutionGap axes as diagnostic context only:
       - `classification`, `actionability`.
    5. Record the current 17 sample files with `kind`, `appLayer`, `functionalArea`, `fileRole`, default unresolved count, and raw count.
  - Implementation Gate: do not choose or change a file group name until this inventory is recorded.
  - Acceptance: evidence proves all 17 target files are `kind=source`, `appLayer=backend`, and have non-unknown `fileRole` values.

- [x] [G2-A] Define and freeze the file group naming rule.
  - Goal: define an exact naming rule that matches current Anvien naming style and yields a concrete group name.
  - Work Steps:
    1. Identify the object being grouped:
       ```text
       object = File
       field = fileGroup
       ```
    2. Identify the app side from file `appLayer`:
       ```text
       app_side = backend
       ```
    3. Identify the purpose family from file `fileRole` values:
       ```text
       purpose_family = support_model_helper
       ```
       This family covers helper, config, parser/runtime model, adapter, fallback adapter, analyzer helper, test helper, contract model, storage helper, and model files.
    4. Compose the machine key:
       ```text
       <app_side>_<purpose_family>
       backend_support_model_helper
       ```
    5. Compose the display label:
       ```text
       <App Side> <purpose family words> files
       Backend support/model/helper files
       ```
    6. Add contract tests for the exact key and label.
  - Implementation Gate: do not implement group assignment until the key and label are frozen in tests.
  - Acceptance: tests assert `backend_support_model_helper` and `Backend support/model/helper files` exactly.

- [x] [G3-A] Define membership rules and boundary rules.
  - Goal: make group membership deterministic from file axes, not from unresolved count math.
  - Work Steps:
    1. Define required membership conditions:
       ```text
       kind == source
       appLayer == backend
       fileRole in allowed_role_set
       ```
    2. Define allowed role set:
       ```text
       helper
       storage_helper
       config
       adapter
       fallback_adapter
       test_helper
       analyzer_helper
       parser_model
       runtime_model
       contract_model
       model
       ```
    3. Treat `functionalArea` as supporting evidence, not part of the group key for this slice.
    4. Define exclusion rules:
       - `fileRole=unknown` is excluded.
       - frontend and frontend test files are excluded.
       - docs/config/generated file kinds are excluded.
       - non-source files are excluded.
    5. Add tests that all current 17 sample files enter the group.
    6. Add tests that each exclusion rule stays out of the group.
  - Implementation Gate: do not surface `fileGroup` until membership and boundary tests exist.
  - Acceptance: tests prove the 17-file sample enters the group and boundary fixtures do not.

- [x] [G4-A] Implement shared backend file group classification.
  - Goal: produce `fileGroup` once in backend classification logic before CLI/API/Web rendering.
  - Work Steps:
    1. Run impact checks for the exact symbols/files to edit:
       - `FileSummary`;
       - file context/list builder functions;
       - semantic metadata functions;
       - generated Web contract functions.
    2. Locate where `kind`, `appLayer`, `functionalArea`, and `fileRole` are assembled.
    3. Add a shared classifier that computes:
       ```text
       fileGroup = backend_support_model_helper | unknown/empty
       ```
    4. Add `fileGroup` to `FileSummary`.
    5. Inspect the graph File-node enrichment path and record exactly one outcome in evidence:
       - File nodes already receive semantic identity properties during analyze: add `fileGroup` to that same path.
       - File nodes do not yet receive semantic identity properties during analyze: keep `fileGroup` authoritative in shared backend `FileSummary` and file projection for this slice, record the graph-node follow-up explicitly, and do not replace it with Web-only or display-only grouping.
    6. Ensure CLI/API/Web consumers read `FileSummary.fileGroup` and do not reimplement path checks.
  - Implementation Gate: impact output must be recorded before editing shared backend or contract symbols.
  - Acceptance: backend file summaries for the 17 sample files include `fileGroup=backend_support_model_helper`.

- [x] [G5-A] Add file group metadata and generated contract support.
  - Goal: make group names and labels stable API/Web contract data.
  - Work Steps:
    1. Add semantic metadata for file groups:
       ```text
       FILE_GROUPS includes backend_support_model_helper
       FILE_GROUP_LABELS includes Backend support/model/helper files
       ```
    2. Add TypeScript generated contract support:
       ```text
       FileGroup
       FileGroupLabel
       FileSummary.fileGroup
       ```
    3. Regenerate Web contracts.
    4. Add Go contract tests and generated TypeScript assertions.
  - Implementation Gate: backend group classifier tests must pass before generated contract changes are accepted.
  - Acceptance: generated Web contract includes the group field and metadata; no Web component hard-codes the group key or label.

- [x] [G6-A] Add file projection group aggregation.
  - Goal: make Anvien report the file group directly instead of relying on `rawUnresolvedFiles - unresolvedFiles`.
  - Work Steps:
    1. Add a file projection group summary keyed by `fileGroup`.
    2. Include these aggregate fields:
       ```text
       key
       label
       files
       defaultUnresolved
       rawUnresolved
       roles
       appLayers
       functionalAreas
       sampleFiles
       ```
    3. For the current 17-file anchor sample, assert:
       ```text
       sampleFiles=17
       sampleDefaultUnresolved=0
       sampleRawUnresolved=376
       ```
    4. Record the full group totals from backend identity rules after implementation. Do not force the full group total to equal 17 unless the identity rules genuinely return only the current 17-file anchor sample.
    5. Assert sample role breakdown:
       ```text
       analyzer_helper=2
       helper=3
       contract_model=1
       storage_helper=1
       test_helper=1
       config=2
       parser_model=3
       runtime_model=2
       adapter=1
       fallback_adapter=1
       ```
  - Implementation Gate: group membership must be computed before aggregation.
  - Acceptance: group summary is computed from `fileGroup` membership directly, the 17-file sample is fully covered, and no count is derived from `rawUnresolvedFiles - unresolvedFiles`.

- [x] [G7-A] Surface the group in analyze, CLI, API, and graph-health outputs.
  - Goal: make command output name the group directly.
  - Work Steps:
    1. Add analyze output line:
       ```text
       fileProjection.group key="backend_support_model_helper" label="Backend support/model/helper files" files=<measured_total> defaultUnresolved=<measured_total_default> rawUnresolved=<measured_total_raw>
       ```
    2. Add `Group` to file-hotspots human rows:
       ```text
       Path    Group    Role    Layer    Area    ...
       ```
    3. Add `Group` to graph-health file rows where file projection rows are printed.
    4. Ensure JSON output contains `fileGroup` on file summaries and group summaries where applicable.
    5. Add CLI and JSON tests for these outputs.
  - Implementation Gate: file projection group aggregation must exist before output formatting.
  - Acceptance: users can see `Backend support/model/helper files` directly in command/API output.

- [x] [G8-A] Surface the group in Web file views.
  - Goal: make Web UI show the file group as the main file identity group, with role as a subcategory.
  - Work Steps:
    1. Update File Map to display `Backend support/model/helper files` where file identity metadata is shown.
    2. Keep `fileRole` visible as a subcategory, not as the only grouping signal.
    3. Update File Detail to show:
       ```text
       Group: Backend support/model/helper files
       Role: <role label>
       Layer: Backend
       Area: <functional area label>
       ```
    4. Inspect File Tree; if it renders file summary metadata or semantic file filters, add group there too.
    5. Add Web unit tests and e2e coverage for group rendering.
  - Implementation Gate: generated Web contracts must include `FileSummary.fileGroup` before Web component edits.
  - Acceptance: Web renders the backend group label from backend data and does not infer group from path.

- [x] [G9-A] Validate the corrected model.
  - Goal: prove the group exists, is visible, and solves the display problem directly.
  - Work Steps:
    1. Run full build before tests:
       ```powershell
       powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1
       ```
    2. Run focused Go tests for semantic, filecontext, contracts, CLI, HTTP/API, and MCP consumers.
    3. Run Web unit tests and Web/e2e tests if Web changed.
    4. Run:
       ```powershell
       .\anvien\bin\anvien.exe analyze --force
       .\anvien\bin\anvien.exe file-hotspots --repo Anvien --json --sort path --limit 0
       .\anvien\bin\anvien.exe graph-health summary --repo Anvien --json
       ```
    5. Verify:
       - The `backend_support_model_helper` full group total is measured and recorded.
       - All 17 required files have `fileGroup=backend_support_model_helper`.
       - The 17-file sample has `sampleDefaultUnresolved=0`.
       - The 17-file sample has `sampleRawUnresolved=376`.
       - Analyze output prints the direct `fileProjection.group` line.
       - Web shows the group label directly.
  - Implementation Gate: all relevant implementation tests must pass before graph/UI validation is accepted.
  - Acceptance: evidence and benchmark ledgers record group count, role breakdown, command output proof, and UI proof.

- [ ] [G10-A] Run detect-changes, update ledgers, and commit.
  - Goal: close the corrected implementation with traceable evidence.
  - Work Steps:
    1. Run:
       ```powershell
       .\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all
       ```
    2. Review affected files, flows, app layers, functional areas, and risk level.
    3. Update the evidence ledger with:
       - commands run;
       - impact/blast radius;
       - implementation evidence;
       - validation evidence;
       - failures and handling;
       - detect-changes output summary.
    4. Update the benchmark ledger with group counts and role breakdown.
    5. Commit the implementation slice.
    6. Record the commit hash in plan/evidence.
  - Implementation Gate: do not commit before detect-changes and ledger updates.
  - Acceptance: commit exists, plan artifacts record closure, and no required checklist item remains open.

## Risk Notes

- `FileSummary`, file projection, Web contracts, and CLI output have broad consumers; expect HIGH/CRITICAL impact and validate carefully.
- The group must be named from real file identity: backend support/model/helper.
- Do not replace the concrete group with a generic bucket.
