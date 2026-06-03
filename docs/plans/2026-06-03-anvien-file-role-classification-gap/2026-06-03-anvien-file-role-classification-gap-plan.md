# Anvien File Group Classification Plan

Date: 2026-06-03

Status: Reopened - file group classification required

Companion files:

- Evidence ledger: [2026-06-03-anvien-file-role-classification-gap-evidence.md](2026-06-03-anvien-file-role-classification-gap-evidence.md)
- Benchmark ledger: [2026-06-03-anvien-file-role-classification-gap-benchmark.md](2026-06-03-anvien-file-role-classification-gap-benchmark.md)

## Checklist

- [ ] [G0] Planner contract and corrected target.
  - This checklist is the source of truth for the corrected plan.
  - Every checklist item must be a mini-plan with enough information to execute that item without reading hidden context from another section.
  - Do not move naming rules, membership rules, output rules, validation gates, or acceptance criteria outside checklist items.
  - This corrected plan replaces the earlier narrow "add a fileRole label" plan.
  - The previous implementation commit `444dcdd feat: add file role classification` added `fileRole`, but that was only role metadata.
  - The product bug is:
    ```text
    Files Anvien can identify by real file type are not placed into a concrete file group that users can read directly.
    ```
  - The implementation target is:
    ```text
    fileGroup: backend_support_model_helper
    label: Backend support/model/helper files
    ```
  - The group must be visible directly in graph/file projection/API/CLI/Web outputs.
  - The group must not require a user to infer meaning from `rawUnresolvedFiles - unresolvedFiles`.
  - Acceptance:
    - This plan has no required implementation rules outside checklist items.
    - Future implementation work updates checklist items as work completes.
    - Reading only checklist items gives the full task, naming rule, target fields, target output, validation, and commit gate.

- [ ] [G1] Audit current Anvien classification axes before naming any new group.
  - Run `.\anvien\bin\anvien.exe analyze --force`.
  - Record current file classification axes from file projection:
    - `kind`: examples are `source`, `test`, `docs`, `config`, `generated`.
    - `appLayer`: examples are `backend`, `backend_test`, `frontend`, `frontend_test`, `api`, `docs`, `config`, `generated_contract`.
    - `functionalArea`: examples are `providers`, `storage`, `analyzer`, `cli`, `query`, `resolution`, `session`, `mcp`, `web_graph_ui`.
    - `fileRole`: examples are `helper`, `config`, `parser_model`, `runtime_model`, `adapter`, `fallback_adapter`, `test_helper`, `analyzer_helper`, `contract_model`, `storage_helper`, `model`, `unknown`.
  - Record current graph node classification axes only as naming reference, not as the target group:
    - `label`: examples are `File`, `Function`, `Method`, `Struct`, `ResolutionGap`, `Process`, `Community`.
    - node `appLayer`.
    - node `functionalArea`.
  - Record current ResolutionGap axes only as reference for source-site diagnostics, not as the file group:
    - `classification`: examples are `builtin`, `standard_library`, `test_framework`, `in_repo_unresolved`, `external_library`.
    - `actionability`: examples are `analyzer_gap`, `non_actionable`, `review`.
  - Acceptance:
    - Evidence shows the current naming style uses lower snake case for machine keys.
    - Evidence shows file grouping must start from file axes, not from ResolutionGap axes.
    - Evidence includes the current 17 target files with `kind=source`, `appLayer=backend`, their `functionalArea`, their `fileRole`, default unresolved count, and raw count.

- [ ] [G2] Define the file group naming procedure and apply it to the target group.
  - Naming procedure must be followed in this order:
    1. Identify the object being grouped.
       - For this plan the object is `File`.
       - Therefore the group field is `fileGroup`.
    2. Identify the application side from file `appLayer`.
       - For the target sample the side is `backend`.
       - Do not start the name from unresolved/source-site words because the group is a file group.
    3. Identify the file purpose family from `fileRole`.
       - For the target sample the purpose family is `support_model_helper`.
       - This family covers support files, helper files, config files, parser/runtime model files, adapter files, and contract model files.
    4. Compose the machine key as:
       ```text
       <app_side>_<purpose_family>
       ```
    5. Compose the display label as:
       ```text
       <App Side> <purpose family words> files
       ```
  - Required key:
    ```text
    backend_support_model_helper
    ```
  - Required display label:
    ```text
    Backend support/model/helper files
    ```
  - Required metadata contract:
    ```text
    fileGroup = "backend_support_model_helper"
    fileGroupLabel = "Backend support/model/helper files"
    ```
  - Acceptance:
    - Contract tests assert the exact key `backend_support_model_helper`.
    - Contract tests assert the exact label `Backend support/model/helper files`.
    - The plan evidence records why the name starts with `backend`: all 17 current target files are backend files.
    - The plan evidence records why the purpose family is `support_model_helper`: the current roles are support/helper/config/model/adapter roles.

- [ ] [G3] Define membership rules for `backend_support_model_helper`.
  - A file enters this group only when all required conditions are true:
    ```text
    kind == source
    appLayer == backend
    fileRole is in the allowed role set
    ```
  - Allowed role set:
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
  - `functionalArea` is supporting evidence but not part of the group key for this slice.
    - Example: `storage`, `providers`, `analyzer`, `session`, `query`, `cli`, and `resolution` can all still be part of the same backend support/model/helper group if the file role matches.
  - Boundary rules:
    - `fileRole=unknown` does not enter this group.
    - `appLayer=frontend` does not enter this group.
    - `appLayer=frontend_test` does not enter this group.
    - `kind=docs`, `kind=config`, and `kind=generated` do not enter this group in this slice.
    - A backend source file with `fileRole=test_helper` can enter this group when it is backend support code such as `internal/testutil/path.go`.
  - Required current sample membership:
    - `internal/frameworks/frameworks.go`
    - `internal/scopeir/sort_keys.go`
    - `internal/group/types.go`
    - `internal/repo/paths.go`
    - `internal/testutil/path.go`
    - `internal/repo/settings.go`
    - `internal/repo/runtime_config.go`
    - `internal/cobol/copy_expander.go`
    - `internal/parser/metrics.go`
    - `internal/session/error.go`
    - `internal/resolution/source_site.go`
    - `internal/scopeir/facts.go`
    - `internal/scopeir/range.go`
    - `internal/session/types.go`
    - `internal/cli/exit_error.go`
    - `internal/lbugnative/runner.go`
    - `internal/lbugnative/runner_default.go`
  - Acceptance:
    - Unit tests assert all 17 files enter `backend_support_model_helper`.
    - Unit tests assert frontend/source files do not enter it.
    - Unit tests assert unknown-role backend source files do not enter it.
    - Unit tests assert docs/config/generated files do not enter it for this slice.

- [ ] [G4] Put `fileGroup` into the graph/file classification layer, not only into display labels.
  - Use Anvien impact before editing the owner functions.
  - Inspect and identify the exact source locations where `File` metadata is assembled:
    - where `kind` is set;
    - where `appLayer` is set;
    - where `functionalArea` is set;
    - where the current `fileRole` is calculated or copied into `FileSummary`;
    - where `FileSummary` is built for file-context/file-hotspots.
  - Add group classification as a real file classification step:
    ```text
    File metadata -> kind/appLayer/functionalArea/fileRole -> fileGroup
    ```
  - Required storage behavior:
    - `FileSummary.fileGroup` must exist.
    - `FileSummary.fileGroupLabel` may exist if labels are not supplied through metadata tables.
    - If Anvien writes semantic properties to `File` graph nodes during analyze, add `fileGroup` there too.
    - If File node storage cannot be updated in this slice, record that as a blocker before implementation; do not silently keep the group only as Web/UI display text.
  - Acceptance:
    - The implementation has one shared classifier for file groups.
    - CLI/API/Web do not reimplement path checks.
    - `fileGroup` is produced before output rendering.

- [ ] [G5] Define group metadata in contracts and semantic definitions.
  - Add a semantic contract for file groups:
    ```text
    FILE_GROUPS = ["backend_support_model_helper", ...]
    FILE_GROUP_LABELS includes:
      key: backend_support_model_helper
      displayLabel: Backend support/model/helper files
      cliLabel: backend-support-model-helper
      webLabel: Backend support/model/helper files
    ```
  - Add generated Web type support:
    ```text
    FileGroup
    FileGroupLabel
    FileSummary.fileGroup
    ```
  - Keep `fileRole` as a subcategory inside the group.
  - Acceptance:
    - Go contract tests assert `FILE_GROUPS`.
    - Generated TypeScript includes `FileSummary.fileGroup`.
    - No component invents group names locally.

- [ ] [G6] Add group aggregation to file projection.
  - Add a file projection group summary keyed by `fileGroup`.
  - Required aggregate for each group:
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
  - Required current target aggregate:
    ```text
    key="backend_support_model_helper"
    label="Backend support/model/helper files"
    files=17
    defaultUnresolved=0
    rawUnresolved=376
    ```
  - Role breakdown for current target aggregate must include:
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
  - Acceptance:
    - Group summary is computed directly from file group membership.
    - The 17-file group is not discovered by subtracting `unresolvedFiles` from `rawUnresolvedFiles`.

- [ ] [G7] Make analyze and CLI output show the group directly.
  - Analyze output must include a direct group line:
    ```text
    fileProjection.group key="backend_support_model_helper" label="Backend support/model/helper files" files=17 defaultUnresolved=0 rawUnresolved=376
    ```
  - File hotspot human output should include `Group`:
    ```text
    Path    Group    Role    Layer    Area    ...
    ```
  - Graph-health file output should include `Group` where file rows are printed.
  - JSON output should include `fileGroup` on each file summary and group summaries where file projection summaries are returned.
  - Acceptance:
    - CLI tests assert the group line.
    - CLI tests assert the `Group` column.
    - JSON tests assert `backend_support_model_helper` without requiring subtraction math.

- [ ] [G8] Make Web UI show the file group as the main file identity group.
  - File Map:
    - show `Backend support/model/helper files` as group metadata or group/filter surface;
    - keep `fileRole` visible as a subcategory;
    - do not replace the group with only `Role`.
  - File Detail:
    - show:
      ```text
      Group: Backend support/model/helper files
      Role: <role label>
      Layer: Backend
      Area: <functional area label>
      ```
  - File Tree:
    - inspect whether it renders file summary or semantic file filters;
    - if it does, include the file group in the same file metadata/filter layer.
  - Acceptance:
    - Web unit tests assert the group label.
    - Web/e2e tests assert File Map or File Detail renders `Backend support/model/helper files`.
    - No Web code derives group from path.

- [ ] [G9] Validate that the corrected model solves the actual display problem.
  - Run full build before tests:
    ```powershell
    powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1
    ```
  - Run focused Go tests for semantic, filecontext, contracts, CLI, HTTP/API, and MCP consumers.
  - Run Web unit tests and e2e tests if Web display changes.
  - Run:
    ```powershell
    .\anvien\bin\anvien.exe analyze --force
    .\anvien\bin\anvien.exe file-hotspots --repo Anvien --json --sort path --limit 0
    .\anvien\bin\anvien.exe graph-health summary --repo Anvien --json
    ```
  - Required validation facts:
    - `backend_support_model_helper.files=17`.
    - All 17 required sample files have `fileGroup=backend_support_model_helper`.
    - `backend_support_model_helper.defaultUnresolved=0`.
    - `backend_support_model_helper.rawUnresolved=376`.
    - Analyze output prints the direct `fileProjection.group` line.
    - UI shows the group label directly.
  - Acceptance:
    - Evidence and benchmark ledgers record group count, role breakdown, and output proof.
    - A user can see the group by name without computing any delta.

- [ ] [G10] Run impact, detect changes, update ledgers, and commit.
  - Before implementation edits, run impact for:
    - file group classifier;
    - `FileSummary`;
    - file projection aggregation;
    - Web contract generator;
    - Web File Map/File Detail if edited.
  - Before commit, run:
    ```powershell
    .\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all
    ```
  - Update:
    - evidence ledger with commands, impact, validation, failures, and handling;
    - benchmark ledger with group counts and role breakdown;
    - this checklist item when complete.
  - Commit after evidence is updated.
  - Acceptance:
    - Implementation commit exists.
    - Plan artifacts record the commit hash.
    - No required work remains for this corrected plan.
