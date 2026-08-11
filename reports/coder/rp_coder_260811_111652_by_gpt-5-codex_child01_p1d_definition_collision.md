# Coder Report: Child 01 P1-D Canonical Definition Collision

Status: ACCEPTED FOR ISOLATED COMMIT; FINAL HASH REPORTED BY GIT

## Metadata

- Report file: `reports/coder/rp_coder_260811_111652_by_gpt-5-codex_child01_p1d_definition_collision.md`
- Report time: `2026-08-11 11:16:52 +07:00`
- Coder: `gpt-5-codex`
- Repository: `E:\Anvien`
- Git base: `df0d7b7b753620f56c932f0e13391c1c016a8f72`
- Scope: Child 01 `P1-D` only — reject non-exact canonical Definition identity conflicts without changing global graph mutation.
- Authority: `AGENTS.md`, the user work rules, `docs/contracts/graph-accuracy-contract.md`, the graph-accuracy roadmap, and the Child 01 plan/evidence/benchmark/actual-status ledgers.
- Problem origin: `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien`; only the measured collision/loss finding is authority, while its DRAFT two-tier design remains non-authoritative.

## Invariant Family Map

- Family: accepted P1-C Definition GraphID -> canonical Definition projection -> exact-idempotent replay or explicit conflict -> valid relationships/error boundary.
- Authority / source truth: the P1-C `defsByFile` occurrence denominator and canonical GraphID, plus the contract rule that conflicting canonical facts cannot silently replace, skip, or collapse.
- Production owners changed: `internal/resolution/emit.go` for canonical Definition emission only; `internal/resolution/resolve.go` for minimum fail-clear propagation.
- Sibling surfaces checked: generic `Graph.AddNode`/`Graph.init`, File-node enrichment, `emitRelationship`/`AddRelationship`, all 11 production `AddNode` caller families, resolution metrics/disposal, P1-C occurrence/endpoints, and command-side error-before-write behavior.
- Preserve-only siblings: global node mutation, ten non-resolution caller families, relationship identity/merge, provider identity, persistence/readers, `E:\cheapapp.org`, P1-E, and later Children.
- Forbidden fallback: blanket `AddNode` contract rewrite, warning-only loss, same-batch subset acceptance, first/last-writer selection, provider hack, fake source collision fixture, relationship redesign, persistence rollback architecture, or target execution.
- Verify matrix: built positive provider path; compiled malformed-ScopeIR fault path; exact duplicate; property add/remove in both orders; base enrichment; no replacement; relationship cardinality/endpoints; File enrichment; P1-C conservation; 11 caller families; 61 product packages.

## Blast Radius and Scope Decision

- Fresh pre-edit impact at clean P1-C commit is CRITICAL: `Graph.AddNode` `296` symbols / `70` files / `9` modules / `77` processes; `Graph.init` `314/69/11/84`; `emitter.emitNode` `6/5/2/25`; `emitDefinitionNodes` `24/9/7/33`; `ResolveBoundInto` `78/35/24/33`.
- Full inbound inventory contains `142` direct `AddNode` calls across `55` files: `126` test calls and `16` production callsites in exactly 11 families (`cobol`, `communities`, `documents`, `graphhealth`, `orm`, `processes`, `resolution`, `routes`, `semantic`, `structure`, `tools`).
- Source classification shows mixed commands: File metadata, diagnostics, and structure need enrichment/update, while canonical resolution Definition insertion needs exact replay or fail-clear conflict.
- The implementation therefore leaves `internal/graph/types.go` unchanged and edits only the Definition branch plus minimum error propagation. CRITICAL is retained as a blast-radius warning, not an edit ban.
- Post-correction impact remains CRITICAL for `emitter` `21/6/2/37`, `newEmitter` `26/10/7/33`, `emitter.emitDefinitionNode` `8/6/2/25`, and `definitionNodeCompatible` `5/3/2/13`; the four new P1-D test functions are LOW with zero upstream impact.

## Implementation

- `internal/resolution/emit.go:15-32` adds invocation-local `definitionNodesSeen` state to the existing emitter responsibility.
- `internal/resolution/emit.go:40-57` distinguishes canonical nodes emitted in this invocation from pre-existing graph state.
- `internal/resolution/emit.go:59-63` requires exact label and complete property-map equality for same-batch replay.
- `internal/resolution/emit.go:74-85` accepts a pre-existing node only when label and every incoming canonical property match, preserving additional pre-existing enrichment without rewriting the node.
- `internal/resolution/emit.go:187-263` returns a collision before emitting the conflicting Definition's `DEFINES` or owner relationship.
- `internal/resolution/resolve.go:50-70` propagates the error with `Result.Graph == nil`, records current metrics, and preserves deferred binding-accumulator disposal.
- Generic node replacement/enrichment, File enrichment, relationship merge, all non-resolution callers, persistence/readers, and later semantics are unchanged.

## Direct QA

- New direct behavior test: `internal/resolution/definition_collision_test.go`.
- `TestP1DDefinitionConflictFailsClearAtResolveBoundary` at line 13 covers range drift and optional canonical property removal/addition in both orders, clear category/ID, nil result graph, and accumulator disposal.
- `TestP1DDefinitionConflictDoesNotReplaceExistingNode` at line 48 runs through `ResolveInto`, compares the complete pre-existing node, and requires one node with the identity.
- `TestP1DDefinitionExactReinsertionIsIdempotentAndPreservesEnrichment` at line 78 runs through `ResolveInto`, preserves the full enriched node, and requires one identity node.
- `TestP1DDefinitionExactDuplicateKeepsOneNodeAndValidEndpoints` at line 104 requires one Definition node, exactly one expected `DEFINES`, and existing endpoints for every relationship.
- Sibling tests explicitly retained: `TestResolveIntoPreservesExistingFileNodeMetadata` and `TestP1CIdentityConservesSameNameOccurrencesAndEndpoints`.

## Verification Outputs

- `[PASS]` full build: `powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\full-build.ps1` exited `0` in `140.4s`; packaged Go/Ladybug runtime, npm install/build/global install, web/launcher, CLI `1.2.8`, and built self-analyze completed. Self-analyze: `1,571` scanned, `680` parsed-code, `0` failed, `95,699` nodes, `134,630` relationships.
- `[PASS]` built positive boundary: `anvien\bin\anvien.exe analyze <repo-local P1-C fixture copy> --force --json` produced `1/1` parsed-code, `0` failed, `20` nodes, `19` relationships. Source/copy SHA-256 equality: `327F6BD02FCB341078902FE4302CAB03A2077DDFE98B379164C17E96B547A078`.
- `[PASS]` compiled fault boundary: package test executable SHA-256 `233EAEE1C1510627AEEFCA4C852FC20AC4A5121C7098159A1D31D5767C468F2E`; isolated `TestP1DDefinitionConflictFailsClearAtResolveBoundary` PASSed its three subcases.
- `[PASS]` focused six-test matrix: all six top-level tests and three collision subcases PASS.
- `[PASS]` `go test ./internal/resolution -count=1`.
- `[PASS]` `go test ./internal/graph -count=1`.
- `[PASS]` exact 11 production caller-family package regression.
- `[PASS]` product regression: `go list -e ./...` returned `72`; excluding only the measured 11 non-standalone `*/anvien/test/fixtures/*` corpus packages left `61`, and the exact run PASSed `61/61`.
- `[PASS]` `git diff --check` after final pre-review ledger updates.
- `[PASS after first-review correction]` current-slice cleanup: the positive repo/home, compiled test executable, and held old binary remain absent. A complete top-level current-date `p1d-*` inventory found 20 additional artifacts; 11 exact FAIL/retry/superseded paths were removed and nine final/unique valid pre-flight captures were retained. Historical `p1d1*`, `p1d2*`, `p1d3*`, nested target-debug work, and unrelated `.tmp` artifacts were untouched.

## Failed Attempts and Closure

- First source candidate used an asymmetric incoming-subset comparison. Zero-trust review demonstrated that same-batch optional-property removal could be accepted in one order. The correction separates invocation-local canonical payload from base-graph enrichment and requires exact same-batch equality.
- First QA candidate used two helper-only tests and an endpoint loop that could pass with zero relationships. The final tests use `ResolveInto`/`Resolve`, compare complete nodes/cardinality, and require exactly one expected `DEFINES`.
- The first full-build attempt stopped before compile because four editor-owned MCP executables held the generated ignored binary through the global npm junction. No process was killed. The old binary was moved reversibly under repo-local `.tmp`, the exact script then PASSed, and the old binary was deleted after its handles released.
- No provider/source collision fixture was invented: normal providers derive occurrence IDs from source positions, so the negative boundary is the compiled real `ResolveBoundInto` malformed-ScopeIR probe approved by the plan.
- The first formal Supervisor review independently cleared production source, full build, runtime, focused/package/11-family/`61/61` QA, and preserve-only siblings, but REJECTed contradictory living-ledger prose and the incomplete broader `.tmp` inventory. The correction below changes only ledgers/report/artifacts; production and tests remain byte-identical.

## E2E Verification

E2E Verification:
  [PASS] Compiled: exact repo-native full build -> exit `0`.
  [PASS] Runtime happy path: built CLI -> provider/ScopeIR -> resolution -> graph `20/19` with hash-identical fixture.
  [PASS] Runtime edge case: compiled real resolution package -> malformed duplicate ScopeIR -> explicit Definition collision with nil result graph.
  [PASS] Regression: focused matrix, resolution/graph packages, 11 caller families, and product packages `61/61`.

## Sibling Surfaces Checked

- Global `Graph.AddNode`/`Graph.init`: unchanged; mixed enrichment/update semantics preserved.
- File metadata enrichment: unchanged and direct regression PASSed.
- `emitRelationship`/`AddRelationship`: unchanged; exact duplicate still yields one expected `DEFINES` with valid endpoints.
- Ten non-resolution production caller families: unchanged; all 11-family package tests PASS.
- P1-C identity/occurrence boundary: unchanged; conservation/endpoint regression PASSed.
- CLI failure side effects: source checks analyze error before result recording/snapshot persistence; no new transaction or storage behavior was introduced.
- Persistence/readers and `E:\cheapapp.org`: not accessed or changed; Child 02 and P1-E remain closed.

## Legacy Fallback Status

No production or QA compatibility fallback was added. Non-exact canonical facts always return an error; only exact same-batch replay and source-proven pre-existing enrichment are accepted.

## First Supervisor Reject Correction

- Reject report: `reports/Supervisor/rp_supervisor_260811_114011_by_gpt-5-codex_child01_p1d.md` — verdict `REJECT`; production behavior clear; living-ledger truthfulness and artifact lifecycle only.
- Before cleanup: exactly 20 top-level `p1d-*` files dated 2026-08-11 under repo-local `.tmp`.
- Deleted FAIL outputs: `.tmp/p1d-impact-graph-addnode.json`, `.tmp/p1d-impact-graph-init.json`, `.tmp/p1d-impact-emitter-emitnode.json`.
- Deleted exact-UID retries superseded by newer `*-final` captures: `.tmp/p1d-impact-emit-definition-nodes-exact.json`, `.tmp/p1d-impact-emitter-emitnode-exact.json`, `.tmp/p1d-impact-graph-addnode-exact.json`, `.tmp/p1d-impact-graph-init-exact.json`, `.tmp/p1d-impact-resolve-bound-into-exact.json`.
- Deleted earlier file-detail captures superseded by their newer `*-final` captures: `.tmp/p1d-file-detail-graph-types.json`, `.tmp/p1d-file-detail-resolution-emit.json`, `.tmp/p1d-file-detail-resolution-resolve.json`.
- Retained final file-detail captures: `.tmp/p1d-file-detail-graph-types-final.json`, `.tmp/p1d-file-detail-resolution-emit-final.json`, `.tmp/p1d-file-detail-resolution-resolve-final.json`.
- Retained unique/final impact captures: `.tmp/p1d-file-impact-graph-types.json`, `.tmp/p1d-impact-emit-definition-nodes-final.json`, `.tmp/p1d-impact-emitter-emitnode-final.json`, `.tmp/p1d-impact-graph-addnode-final.json`, `.tmp/p1d-impact-graph-init-final.json`, `.tmp/p1d-impact-resolve-bound-into-final.json`.
- Retention reason: these nine files are the latest or only successful pre-flight file-detail/impact captures behind `E1-P1D-IMPACT1`; none is a failed attempt, retry duplicate, or superseded output. They are debug-layer traceability only; durable claims remain in the Child evidence ledger and Supervisor report.
- After cleanup: current-date top-level `p1d-*` inventory is exactly `9`; every deleted path is absent. Historical and unrelated `.tmp` work is preserved.
- Production/test SHA-256 remains unchanged: `internal/resolution/emit.go` `356B667B53E6ACF74D3ACA5686F9A192FF4CF6159CE4F0DE57A8BD4E41081569`; `internal/resolution/resolve.go` `09902FFB236AF124CA3440669671EBDEC93ADE3E61ACC4E3BF18D262F3AC44CE`; `internal/resolution/definition_collision_test.go` `04DC174419393E59EA68126FF7504C404617383FBA46097991E3C11DE0D26D92`.
- Living authority correction: the actual-status matrix, Detailed Findings, allowed-next-action text, next-decision row, plan/roadmap status, evidence row, benchmark inventory, daily notes, and this report now agree that production/build/runtime/QA are clear; the first verdict remains `REJECT`; unconditional re-review/detect/commit remain pending; P1-E and later boundaries remain closed.
- Post-correction verification: consistency assertions PASS `19/19`; `git diff --check` exits `0`; the nine-path retained inventory is exact; production/test hashes remain unchanged; fresh analyze exits `0` with `1,573` scanned, `680` parsed-code, `0` failed, `95,726` nodes, and `134,657` relationships.

## Residual Unverified Surfaces

None inside the P1-D behavior boundary. P1-E target validation, range projection, and Child 02 persistence/readers are explicitly separate closed scopes, not residual P1-D work.

## Review Gate

- History-closure result: `PASS` in `reports/Supervisor/rp_supervisor_260811_122327_by_gpt-5-codex_child01_p1d_history_closure.md`; source, diff, direct runtime, QA, graph/impact, artifact lifecycle, living authority, and closed later boundaries were independently verified with no residual P1-D same-invariant surface.
- Detect result: mandatory fresh analyze reports `1,574/680/0`, `95,739/134,670`; staged all-scope detect-changes exits `0`, covers all 12 paths at risk `high`, records `168` changed symbols / `13` affected symbols / `13` affected processes, fully classifies 24 changed gap occurrences/entities, and reports zero persisted gap/degraded-node/node-with-gap result.
- Final closure: staged diff/ledger/hash/artifact verification PASS; P1-D checklist and `E1-P1D-COMMIT1` use the repository convention that the isolated commit executes immediately after the final ledger state and Git reports the hash externally.
- Required next action: execute `fix(resolution): fail clear on definition collisions` immediately; do not open P1-E if Git fails.
