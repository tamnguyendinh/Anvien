# Tree-sitter Go Module Drift Evidence Ledger

## Metadata

- Date: `2026-07-16`
- Plan: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-plan.md`
- Evidence: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-evidence.md`
- Benchmark: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-benchmark.md`
- Actual status: `docs/plans/2026-07-16-tree-sitter-go-module-drift/2026-07-16-tree-sitter-go-module-drift-actual-status.md`

## Evidence Rules

The evidence file explains why the work is known to be correct.

Evidence ID format:

```text
E<phase>-<item>-<kind><n>
```

- `E0` corresponds to `P0`.
- `E1` corresponds to `P1`.
- `E2` corresponds to `P2`.
- `E3` corresponds to `P3`.
- Use exact evidence IDs inside status and benchmark files.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-GRAPH1`: `anvien analyze --force` completed before graph-based inspection. Output recorded `files: scanned=1464 parsed_code=676 failed=0`, `graph: nodes=84074 relationships=122914`, and graph path `E:\Anvien\.anvien\graph.json`.
- `E0-P0A-FD1`: `anvien file-detail go.mod --repo Anvien --json` succeeded. Target `go.mod` has `parseStatus=parsed`, `risk=low`, `relatedFiles=[]`, and `relationshipTotal=0`.
- `E0-P0A-FD2`: `anvien file-detail .github/scripts/check-tree-sitter-upgrade-readiness.py --repo Anvien --json` failed with `file ".github/scripts/check-tree-sitter-upgrade-readiness.py" not found in repo Anvien`.
- `E0-P0A-FD3`: `anvien file-detail .github/workflows/tree-sitter-upgrade-readiness.yml --repo Anvien --json` failed with `file ".github/workflows/tree-sitter-upgrade-readiness.yml" not found in repo Anvien`.
- `E0-P0A-FD4`: `anvien file-detail .github/dependabot.yml --repo Anvien --json` failed with `file ".github/dependabot.yml" not found in repo Anvien`.
- `E0-P0A-SRC1`: Source inspection of `.github/scripts/check-tree-sitter-upgrade-readiness.py` shows `read_current_runtime()` reads `ANVIEN_DIR / "package.json"` and then `pkg["dependencies"]["tree-sitter"]`.
- `E0-P0A-SRC2`: Source inspection of `.github/scripts/check-tree-sitter-upgrade-readiness.py` shows the report target is hard-coded to `TARGET_RUNTIME = "0.25.0"` and report wording is `Tree-sitter 0.25 upgrade readiness`.
- `E0-P0A-SRC3`: Source inspection of `.github/workflows/tree-sitter-upgrade-readiness.yml` shows workflow name `Tree-sitter Upgrade Readiness`, comments about upgrading `tree-sitter@0.25.0`, and tracking issue title `Tree-sitter 0.25 upgrade readiness`.
- `E0-P0A-SRC4`: Source inspection of `.github/dependabot.yml` shows npm comments and grouping for `/anvien` Tree-sitter grammars, plus comments saying the Tree-sitter runtime is pinned and waiting for the drift check workflow.
- `E0-P0A-SRC5`: Source inspection of `go.mod` shows the current Tree-sitter parser stack is Go modules: `github.com/tree-sitter/go-tree-sitter v0.25.0`, `tree-sitter-go v0.25.0`, `tree-sitter-javascript v0.25.0`, `tree-sitter-typescript v0.23.2`, and additional indirect Tree-sitter grammar modules.
- `E0-P0A-RUN1`: `python .github\scripts\check-tree-sitter-upgrade-readiness.py` failed with `KeyError: 'dependencies'` at line 88 while reading `pkg["dependencies"]["tree-sitter"]`.
- `E0-P0A-RUN2`: `go run ..\cmd\anvien package build-runtime` from `E:\Anvien\anvien` passed and wrote `E:\Anvien\anvien\bin\anvien.exe`.
- `E0-P0A-RUN3`: `.\anvien\bin\anvien.exe analyze --force` passed after package build, proving the current runtime still analyzes the repo with the existing parser stack.
- `E0-P0A-GIT1`: `git status --short` was clean before creating this plan.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`, `P1-B`, `P1-C`

- `E1-P1A-GRAPH1`: Pre-edit `anvien file-detail go.mod --repo Anvien --json` succeeded with `parseStatus=parsed`, `risk=low`, and `relatedFiles=[]`.
- `E1-P1A-IMPACT1`: Pre-edit `anvien impact file go.mod --repo Anvien --direction upstream` returned `risk=LOW`, `impactedCount=0`, `affectedFiles=[]`.
- `E1-P1A-IMPACT2`: Pre-edit `anvien file-detail .github/scripts/check-tree-sitter-upgrade-readiness.py --repo Anvien --json` failed because `.github` files are not indexed as current file-detail targets.
- `E1-P1A-IMPACT3`: Pre-edit `anvien impact file .github/scripts/check-tree-sitter-upgrade-readiness.py --repo Anvien --direction upstream` returned `status=not_found`, `risk=UNKNOWN`, confirming direct source inspection was required for ignored `.github` files.
- `E1-P1A-SRC1`: Replaced the retired npm `package.json["dependencies"]["tree-sitter"]` inventory with `go list -m -u -json all` parsing.
- `E1-P1A-SRC2`: Added Tree-sitter Go module filtering for `github.com/tree-sitter/*`, `github.com/tree-sitter-grammars/*`, and known external Tree-sitter grammar/binding modules.
- `E1-P1B-SRC1`: Added explicit status classification: `UP_TO_DATE`, `GO_MODULE_UPDATE_AVAILABLE`, `UPSTREAM_CORE_AHEAD_GO_BINDING`, `GRAMMAR_UPDATE_AVAILABLE`, and `UNKNOWN_FETCH_FAILED`.
- `E1-P1B-RUN1`: Live checker command produced a Markdown report with `Tree-sitter modules found: **19**`, `Actionable drift rows: **2**`, and `UPSTREAM_CORE_AHEAD_GO_BINDING` for `github.com/tree-sitter/go-tree-sitter`; wrapper validated exit code `1` as expected because two actionable module drifts exist.
- `E1-P1C-TEST1`: Added `.github/scripts/test_check_tree_sitter_upgrade_readiness.py` with fixture coverage proving no npm package dependencies are required.
- `E1-P1C-TEST2`: Added fixture coverage for update parsing, core-ahead informational behavior, actionable grammar updates, and upstream fetch failure classification.
- `E1-P1C-RUN1`: `python .github\scripts\test_check_tree_sitter_upgrade_readiness.py` passed: `Ran 5 tests in 0.001s OK`.
- `E1-P1C-COMPILE1`: `python -m py_compile .github\scripts\check-tree-sitter-upgrade-readiness.py .github\scripts\test_check_tree_sitter_upgrade_readiness.py` passed.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`, `P2-B`

- `E2-P2A-SRC1`: Updated `.github/workflows/tree-sitter-upgrade-readiness.yml` name and issue title from Tree-sitter 0.25 readiness to Tree-sitter Go module drift.
- `E2-P2A-SRC2`: Updated workflow issue parsing to read module rows with uppercase status values from the new Markdown report table.
- `E2-P2A-SRC3`: Added the new checker test file to workflow `pull_request.paths`.
- `E2-P2B-SRC1`: Added root `gomod` Dependabot monitoring for Go dependencies and grouped Tree-sitter Go modules.
- `E2-P2B-SRC2`: Removed stale `/anvien` npm Tree-sitter runtime/grammar comments and ignore rules.
- `E2-P2B-VALIDATE1`: PyYAML parsed `.github/workflows/tree-sitter-upgrade-readiness.yml` and `.github/dependabot.yml` successfully.
- `E2-P2B-VALIDATE2`: Source scan found no remaining stale `tree-sitter@0.25`, `Tree-sitter 0.25`, `peer-dep`, `npm registry`, or `package.json` references in the touched `.github` checker/workflow/dependabot files.
- `E2-P2B-GIT1`: `git status --short --ignored .github ...` shows `.github/` is ignored by `.gitignore`; final commit must use `git add -f` for the touched `.github` files.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`

- `E3-P3A-BUILD1`: `npm run full-build` passed. It built packaged runtime with LadybugDB native runtime `E:\Anvien\.tmp\ladybug-native\v0.18.2\windows-x86_64`, built `anvien-web`, installed global package, ran `anvien version` returning `1.2.7`, and ran `anvien analyze . --force`.
- `E3-P3A-GRAPH1`: Full-build analyze output recorded `files: scanned=1468 parsed_code=676 failed=0`, `graph: nodes=84126 relationships=122966`.
- `E3-P3A-RUN1`: Post-build `python .github\scripts\test_check_tree_sitter_upgrade_readiness.py` passed: `Ran 5 tests in 0.001s OK`.
- `E3-P3A-RUN2`: Post-build live checker wrapper passed: checker exited `1` for expected actionable drift, output contained `Actionable drift rows: **2**`, contained `UPSTREAM_CORE_AHEAD_GO_BINDING`, and did not contain `KeyError`.
- `E3-P3A-YAML1`: Post-build PyYAML parse of workflow and Dependabot config passed.
- `E3-P3A-GRAPH2`: Final pre-detect `anvien analyze --force` passed with `files: scanned=1468 parsed_code=676 failed=0`, `graph: nodes=84126 relationships=122966`.
- `E3-P3A-DETECT1`: `anvien detect-changes --repo Anvien --scope all` passed with `risk_level=low`, `changed_files=1`, and changed file `.gitignore`. Current Anvien graph detection does not report newly untracked `.github` and plan files, so Git status/staging evidence is required for commit closure.
- `E3-P3A-GIT1`: After the user removed `.github/` from `.gitignore`, `.github` files are no longer ignored. `git status --short --ignored` shows many untracked `.github` files; only target plan files and Tree-sitter drift files should be staged for this scope.
- `E3-P3A-CLEAN1`: Removed `.github/scripts/__pycache__/` created by Python tests and added generic `__pycache__/` ignore so Python cache files do not pollute Git after `.github` becomes trackable.

## Closure Evidence

- `ECLOSE-PNA-SUP1`: Supervisor report `reports/Supervisor/rp_supervisor_260716_144532_by_gpt-5-codex_tree-sitter-go-module-drift.md` recorded `Verdict: PASS`.
- `ECLOSE-PNB-CLEAN1`: Dead-work cleanup removed `.github/scripts/__pycache__/`; `.gitignore` now ignores `__pycache__/`.
- `ECLOSE-PNC-PENDING1`: Commit hash and final worktree state will be recorded after commit.
