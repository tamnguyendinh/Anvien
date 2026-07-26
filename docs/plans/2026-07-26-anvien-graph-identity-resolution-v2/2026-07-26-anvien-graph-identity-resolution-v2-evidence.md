# Anvien Graph Identity and TypeScript Resolution Correctness v2 Evidence Ledger

## Metadata

- Date: `2026-07-26`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-actual-status.md`

## Evidence Rules

The evidence file explains why the work is known to be correct.

It should contain:

- metadata and companion files;
- evidence rules or evidence template;
- evidence sections such as `E0`, `E1`, or sections by phase/task;
- user report or problem evidence;
- source inspection, codebase facts, and document facts;
- commands run and pass/fail result;
- impact or blast-radius evidence when code/graph behavior changes;
- implementation evidence: files changed and behavior changed;
- validation evidence: build, tests, e2e, screenshots, or traces;
- failures encountered and how they were handled;
- detect-changes before commit;
- commit hash and closure evidence.

Evidence can reference short metric traces, but long metric tables belong in the benchmark file.

### Evidence ID Naming

Use stable, phase-scoped evidence IDs so `plan.md`, `actual-status.md`, `benchmark.md`, and later agents can reference exact proof without ambiguity.

Format:

```text
E<phase>-<item>-<kind><n>
```

Rules:

- `E<phase>` matches the plan phase number: `E0` for `P0`, `E1` for `P1`, `E2` for `P2`, and so on.
- `<item>` matches the checklist item without the dash: `P0A`, `P1A`, `P2B`.
- `<kind>` is plan-local. Choose a short uppercase token that is meaningful for this repo and this plan.
- `<n>` is a 1-based sequence number within that phase item and kind.
- Keep the same `<kind>` meaning stable inside one plan.
- Do not reuse an evidence ID for different facts.
- Reference exact evidence IDs from `actual-status.md` and `benchmark.md`; avoid referencing only broad section IDs such as `E0` or `E1`.
- Use ranges such as `E0-P0A-FD1..E0-P0A-FD17` only for compact inventory summaries; use exact IDs when a specific status decision depends on a specific fact.
- If nearby plans already use a clear local evidence naming style, follow that style instead of inventing a new one.

Evidence sections follow plan phases. A future evidence ID listed as `pending` is a reserved target, not proof. It becomes valid only when its exact command/artifact/result is recorded.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-RULE1`: read `E:\Anvien\AGENTS.md` and the user-attached master working rules. They require planner use, graph refresh, file-detail/impact before code, code before tests, full build before validation, Supervisor acceptance, per-slice commit, target-boundary protection, and one file per primary business responsibility.
- `E0-P0A-SKILL1`: read the complete planner skill and all four current templates before creating this four-file set.
- `E0-P0A-SKILL2`: read the System-architect guidance. No `SPEC-MAP.md` or approved `Docs/SPEC` family exists for this scope, so the architecture decision table is explicitly proposed and implementation remains unauthorized until owner acceptance; the plan is not presented as pre-existing SPEC authority.
- `E0-P0A-SKILL3`: read the Supervisor skill for the zero-trust plan review gate.
- `E0-P0A-SPEC1`: `rg --files -g 'SPEC-MAP.md' -g 'Docs/SPEC/**' ...` returned no relevant architecture SPEC files.
- `E0-P0A-HELP1`: `anvien --help` was run before relying on Anvien commands.
- `E0-P0A-LOCK1`: `anvien doctor locks --repo E:\Anvien --json` checked current local runtime locks before the forced analyze.
- `E0-P0A-GRAPH1`: `anvien analyze E:\Anvien --force --json` completed at repo HEAD `1932359bee5d78fd8e6167ef94e7974f33e85bd0` after the plan set was created; current graph: `1,505` scanned Files, `676` parsed code files, `0` failed/unsupported/unknown files, `84,558` nodes, and `123,398` relationships. The five-file increase is plan/evidence/benchmark/actual-status/matrix documentation indexed as repository documents, not production code.
- `E0-P0A-GRAPH2`: after the matrix/ledger/slice revisions, a fresh `anvien analyze E:\Anvien --force --json` at the same HEAD reproduced `1,505` scanned Files, `676` parsed code files, `0` failed/unsupported/unknown files, `84,558` nodes, and `123,398` relationships; no production source changed.
- `E0-P0A-GRAPH3`: after synchronizing the reader-row range to `R01..R153`, the latest fresh `anvien analyze E:\Anvien --force --json` completed with the same `1,505` scanned Files, `676` parsed code files, `0` failed/unsupported/unknown files, `84,558` nodes, and `123,398` relationships; no production source changed and the matrix has 153 contiguous anchors.
- `E0-P0A-MATRIX3`: fresh source audit after the red-team correction pass expanded the seed to `195` contiguous rows (`R01..R195`) with no duplicate IDs or anchors and all path/function anchors present. Current tag counts are `S0=4`, `S1=5`, `S2=3`, `S3=56`, `S4=37`, `S5=40`, `S6=30`, `S7=4`, `S8=2`, `S9=17`, `S10=18`, and `S11=14`; native/fallback selectors, concrete fallback backends, Web readers, dispatchers, and explicit non-readers are separate rows. This remains a checked-in seed: P2-A1 must prove complete classification/assignment with zero unlisted readers, and P2-E2 later must prove `rows_passed == rows_total`.
- `E0-P0A-TEMPLATE2`: historical structural audit of the pre-P2-split snapshot found `67` checklist items, `66` implementation/closure trace rows, and `64` P1-P7/P8-C slices with per-work-step Mini-QA. It was accurate for that snapshot but is superseded by the later template/slice audit; it must not be used as proof for the current expanded plan.
- `E0-P0A-GRAPH4`: final refresh command `anvien analyze E:\Anvien --force --json` completed after the Web/dispatcher/backend matrix and Mini-QA edits with `1,505` scanned Files, `676` parsed code, `0` failed/unsupported/unknown, `84,558` nodes, and `123,398` relationships; no production source changed. The graph path is `E:\Anvien\.anvien\graph.json`.
- `E0-P0A-TEMPLATE3`: fresh current-plan structural audit found `102` unique checklist items, `99` eligible implementation/closure slices, `132` numbered work steps inside those slices, and `101` unique traceability rows with no missing or duplicate checklist binding. Every eligible slice contains all `12` explicit Pre-flight fields; every numbered work step contains UI-flow, DB/data-flow, render-location, Mini-QA, and exact evidence-target fields. The evidence-reference audit contains `796` non-meta references (`747` unique implementation traceability references plus `49` unique P0 cross-ledger references), and all `796` are declared in this ledger. Bare `UI flow check: N/A.` and bare `Render location check: N/A.` counts are both `0`; unresolved template or work-marker placeholders are `0`. `E0-P0A-TEMPLATE2` remains historical and superseded.
- `E0-P0A-TEMPLATE4`: post-B1 structural audit of the corrected plan found `99/99` eligible slices with explicit four-field Scope Boundary blocks (`Editable`, `Inspect-only`, `Preserve-only`, `Out of scope`) and `99/99` eligible slices with explicit seven-field Acceptance blocks (`Source`, `Runtime/UI`, `DB/data`, `Behavior test`, `Cleanup/quarantine`, `Evidence IDs`, `Actual-status rows refreshed`). All `N/A` values carry a reason; compact eligible Scope Boundary/Acceptance lines are `0`; field indentation is machine-readable; the existing `102` checklist, `132` work steps, `101` trace rows, pre-existing `796` non-meta evidence references, target boundary, and P2 split are unchanged. `TEMPLATE4` and later refresh/review entries are new P0 meta-evidence outside that frozen non-meta denominator. This closes the structural issue identified by the first Supervisor review; implementation remains blocked by owner authority.
- `E0-P0A-MATRIX4`: fresh source audit of the current matrix found `195` contiguous rows (`R01..R195`), `195` unique IDs, and `195` unique path/function anchors; all `195` files and normalized function/entrypoint anchors exist in current source. Tag counts are `S0=4`, `S1=5`, `S2=3`, `S3=56`, `S4=37`, `S5=40`, `S6=30`, `S7=4`, `S8=2`, `S9=17`, `S10=18`, and `S11=14`. This remains a planned seed: P2-A1 must still prove complete source classification and `unlisted_readers == 0`, and P2-E2 must still prove runtime `rows_passed == rows_total`.
- `E0-P0A-GRAPH5`: latest fresh `anvien analyze E:\Anvien --force --json` at unchanged HEAD `1932359bee5d78fd8e6167ef94e7974f33e85bd0` completed after the current plan read/audit with `1,504` scanned Files, `676` parsed code, `0` failed/unsupported/unknown, `84,532` nodes, and `123,372` relationships; graph path is `E:\Anvien\.anvien\graph.json`. These are the current counts and are not described as reproducing the historical `1,505`/`84,558`/`123,398` snapshot. `git diff --exit-code -- internal` remained clean and `git status --short --untracked-files=all` listed only the five plan artifacts.
- `E0-P0A-GRAPH6`: fresh `anvien analyze E:\Anvien --force --json` at unchanged HEAD `1932359bee5d78fd8e6167ef94e7974f33e85bd0` completed after the B1 plan/ledger correction and first Supervisor report were present. Current graph: `1,505` scanned Files, `676` parsed code, `0` failed/unsupported/unknown, `84,540` nodes, and `123,380` relationships; graph path is `E:\Anvien\.anvien\graph.json`, SHA-256 `069558C97E37E7B904C69101513EA7A58E295250A1A38E559409A68ED380735E`. The extra File versus GRAPH5 is the repo-local Supervisor report; no production source changed, `git diff --exit-code -- internal` remained clean, and the target was not written.
- `E0-P0A-GIT1`: before creation of this plan set, branch was `master`, HEAD was `1932359bee5d78fd8e6167ef94e7974f33e85bd0`, and `git status --short` was empty.
- `E0-P0A-REPORT1`: current accepted causal source is `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md`; it separates the five in-scope defects from the scanner defect and downstream bounded observations.
- `E0-P0A-REPORT2`: `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md` is PASS for the bounded investigation record only; it does not approve remediation.
- `E0-P0A-HISTORY1`: the two earlier report paths originally named by the user are absent at current HEAD. They have not been silently treated as current authority; the committed final causal synthesis and final Supervisor PASS above are the current replacements.
- `E0-P0A-HISTORY2`: the accepted investigation's single target analyze duration is historical baseline context only; it is not proof of the P2-G five-run baseline. P2-G must capture fresh v1 analyze, Ladybug-load, native-query, fallback-query, graph-size, and RSS medians under one reproducible protocol.
- `E0-P0A-BOUNDARY1`: target baseline from the accepted record is `E:\cheapapp.org`, HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, graph `E:\cheapapp.org\.anvien\graph.json`, SHA-256 `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`, `1,359` Files, `84,807` nodes, and `114,125` relationships. The target is analyzed in place and not copied into Anvien.
- `E0-P0A-BOUNDARY2`: the user explicitly excluded the scanner `env/target/logs` defect, prohibited target contamination/copying/root fixtures, and required all reports/plans under Anvien. The target's `.anvien` remains its operational graph location.
- `E0-P0A-USER1`: the user required the repair order: identity -> binding patterns -> export metadata -> module/re-export resolution -> ambient/external resolution and diagnostics.
- `E0-P0A-USER2`: the user required that every file own one primary business responsibility while allowing links to many files/modules when they serve that responsibility.

Fresh file-detail evidence:

| Evidence ID | File | Related files | Risk |
|-------------|------|--------------:|------|
| `E0-P0A-FD1` | `internal/graph/types.go` | 238 | high |
| `E0-P0A-FD2` | `internal/scopeir/facts.go` | 231 | medium |
| `E0-P0A-FD3` | `internal/scopeir/range.go` | 227 | medium |
| `E0-P0A-FD4` | `internal/scopeir/definition_index.go` | 225 | medium |
| `E0-P0A-FD5` | `internal/resolution/indexes.go` | 46 | high |
| `E0-P0A-FD6` | `internal/resolution/resolve.go` | 50 | high |
| `E0-P0A-FD7` | `internal/resolution/emit.go` | 42 | high |
| `E0-P0A-FD8` | `internal/providers/tsjs/definitions.go` | 16 | high |
| `E0-P0A-FD9` | `internal/providers/tsjs/imports.go` | 15 | high |
| `E0-P0A-FD10` | `internal/resolution/import_resolution.go` | 33 | high |
| `E0-P0A-FD11` | `internal/graphhealth/diagnostics.go` | 29 | high |
| `E0-P0A-FD12` | `internal/analyze/analyze.go` | 182 | high |
| `E0-P0A-FD13` | `internal/repo/types.go` | 72 | high |
| `E0-P0A-FD14` | `internal/lbugload/csv.go` | 19 | high |
| `E0-P0A-FD15` | `internal/httpapi/graph.go` | 22 | high |
| `E0-P0A-FD16` | `anvien-web/src/services/backend-client.ts` | 24 | high |

Fresh upstream impact evidence:

| Evidence ID | Symbol | Risk | Impacted symbols | Files | Modules | Processes |
|-------------|--------|------|-----------------:|------:|--------:|----------:|
| `E0-P0A-IMPACT1` | `graphIDForDef` | CRITICAL | 6 | 5 | 3 | 18 |
| `E0-P0A-IMPACT2` | `Graph.AddNode` | CRITICAL | 36 | 15 | 6 | 82 |
| `E0-P0A-IMPACT3` | `Graph.init` | CRITICAL | 68 | 18 | 9 | 82 |
| `E0-P0A-IMPACT4` | `buildWorkspace` | CRITICAL | 8 | 6 | 5 | 28 |
| `E0-P0A-IMPACT5` | `resolveImportedDef` | CRITICAL | 4 | 3 | 1 | 16 |
| `E0-P0A-IMPACT6` | `resolveImports` | CRITICAL | 6 | 5 | 3 | 18 |
| `E0-P0A-IMPACT7` | `resolveTypeAnnotation` | CRITICAL | 6 | 4 | 4 | 35 |
| `E0-P0A-IMPACT8` | `resolveCall` | CRITICAL | 6 | 4 | 4 | 35 |
| `E0-P0A-IMPACT9` | `writeGraphSnapshot` | CRITICAL | 5 | 4 | 4 | 39 |
| `E0-P0A-IMPACT10` | `classifyDiagnostic` | HIGH | 7 | 2 | 3 | 3 |

Current source/migration facts:

- `E0-P0A-SRC1`: `internal/scopeir/definition_index.go` silently ignores duplicate Definition IDs.
- `E0-P0A-SRC2`: `internal/graph/types.go` uses replacement semantics for duplicate node IDs during mutation and last-wins indexing during init/decode.
- `E0-P0A-SRC3`: `internal/lbugload/csv.go` silently first-wins duplicate node IDs.
- `E0-P0A-SRC4`: graph emission writes `visibility`, while Ladybug semantic CSV reads `isExported`.
- `E0-P0A-SRC5`: `internal/repo/types.go` metadata has no graph schema, identity schema, ScopeIR, position encoding, or generation contract.
- `E0-P0A-SRC6`: `internal/analyze/analyze.go` deletes live graph/Ladybug before forced rebuild, loads DB before graph snapshot publication, and does not publish a complete multi-artifact generation atomically.
- `E0-P0A-SRC7`: current consumers derive data from ID text in embeddings, rename, groups, and Web graph components; process/community outputs also depend on lexicographic IDs.
- `E0-P0A-SRC8`: current resolver indexes target ScopeIR definitions only, resolves imported definitions physically in the target module, and does not have a first-class export surface or TypeScript declaration universe.
- `E0-P0A-REDTEAM1`: identity subagent advised separate Declaration/Symbol identities, explicit ranges/encoding, merge evidence, versioned migration, and collision failure. Advisory only; accepted points are independently represented in source evidence and the proposed decision table.
- `E0-P0A-REDTEAM2`: projection subagent identified reader/version/atomicity and ID-parsing blast radius. Advisory only; current source facts `E0-P0A-SRC1..E0-P0A-SRC7` are the P0 authority.
- `E0-P0A-REDTEAM3`: resolver subagent advised separate module/export/outcome/declaration-universe contracts. Advisory only; no subagent verdict is treated as acceptance.

P0 decision evidence:

- `E0-P0A-STATUS1`: current actual-status matrix is populated with exact classifications, relationship counts, touch modes, phase consequences, and an R0 refresh.
- `E0-P0A-PLAN1`: plan contains the required fixed phase order, multiple bounded slices per implementation phase, code-first/full-build/Supervisor/detect/commit gates, target boundary, scanner exclusion, and the one-file/one-responsibility rule.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`, `P1-B`, `P1-C`, `P1-D`, `P1-E`

| Reserved evidence | Required proof | Status |
|-------------------|----------------|--------|
| `E1-P1A-CONTRACT1` | owner-accepted identity/ownership decision record | pending |
| `E1-P1A-REVIEW1` | independent architecture/Supervisor verdict | pending |
| `E1-P1A-COMMIT1` | isolated contract-only commit and worktree state | pending |
| `E1-P1B-IMPACT1` | fresh file-detail/impact for identity/range owners | pending |
| `E1-P1B-SRC1` | production type/ownership diff | pending |
| `E1-P1B-BUILD1` | full build after production code and tests | pending |
| `E1-P1B-TEST1` | position/serialization/type-safety behavior tests | pending |
| `E1-P1B-REVIEW1` | Supervisor PASS | pending |
| `E1-P1B-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1B-COMMIT1` | isolated slice commit and worktree state | pending |
| `E1-P1C-IMPACT1` | fresh identity mapping impact | pending |
| `E1-P1C-SRC1` | canonical tuple/merge/stability implementation | pending |
| `E1-P1C-BUILD1` | full build | pending |
| `E1-P1C-TEST1` | identity, stability, overload, collision, order tests | pending |
| `E1-P1C-REVIEW1` | Supervisor PASS | pending |
| `E1-P1C-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1C-COMMIT1` | isolated slice commit and worktree state | pending |
| `E1-P1D-IMPACT1` | graph mutation/producer impact | pending |
| `E1-P1D-MAP1` | complete producer-to-operation inventory | pending |
| `E1-P1D-SRC1` | strict mutation/decode/validation diff | pending |
| `E1-P1D-BUILD1` | full build | pending |
| `E1-P1D-TEST1` | duplicate, closure, enrich, replace failure matrix | pending |
| `E1-P1D-REVIEW1` | Supervisor PASS | pending |
| `E1-P1D-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1D-COMMIT1` | isolated slice commit and worktree state | pending |
| `E1-P1E-IMPACT1` | fresh emitter/analyze-hook file-detail and upstream impact | pending |
| `E1-P1E-SRC1` | shadow v2 emitter/comparator implementation | pending |
| `E1-P1E-SHADOW1` | canonical v1/v2 shadow signature comparison | pending |
| `E1-P1E-BUILD1` | full build | pending |
| `E1-P1E-BENCH1` | five-run graph expansion/size/RSS measurements | pending |
| `E1-P1E-TARGET1` | in-memory target v2 shadow target-site run before cutover | pending |
| `E1-P1E-ORACLE1` | independent `time`/`now` 4/4 source/TS-oracle manifest | pending |
| `E1-P1E-BOUNDARY1` | target pre/post hash/worktree/graph-path/artifact boundary | pending |
| `E1-P1E-REVIEW1` | Supervisor PASS | pending |
| `E1-P1E-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1E-COMMIT1` | isolated slice commit and worktree state | pending |
| `E1-P1C0-IMPACT1` | declaration occurrence-index impact | pending |
| `E1-P1C0-SRC1` | lossless v2 occurrence index and v1 compatibility adapter | pending |
| `E1-P1C0-BUILD1` | full build | pending |
| `E1-P1C0-TEST1` | declaration conservation/collision tests | pending |
| `E1-P1C0-REVIEW1` | Supervisor PASS | pending |
| `E1-P1C0-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1C0-COMMIT1` | isolated slice commit | pending |
| `E1-P1C0A-IMPACT1` | RelationshipID/aggregation impact | pending |
| `E1-P1C0A-SRC1` | RelationshipID and source-site aggregation implementation | pending |
| `E1-P1C0A-BUILD1` | full build | pending |
| `E1-P1C0A-TEST1` | same-endpoint/multi-source-site conservation tests | pending |
| `E1-P1C0A-REVIEW1` | Supervisor PASS | pending |
| `E1-P1C0A-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1C0A-COMMIT1` | isolated slice commit | pending |
| `E1-P1C0B-IMPACT1` | decode/closure impact | pending |
| `E1-P1C0B-SRC1` | fail-closed graph decode/validation implementation | pending |
| `E1-P1C0B-BUILD1` | full build | pending |
| `E1-P1C0B-TEST1` | duplicate/conflict/endpoint/subset failure tests | pending |
| `E1-P1C0B-REVIEW1` | Supervisor PASS | pending |
| `E1-P1C0B-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1C0B-COMMIT1` | isolated slice commit | pending |
| `E1-P1D1-IMPACT1` | core producer migration impact | pending |
| `E1-P1D1-SRC1` | core producer explicit-operation diff | pending |
| `E1-P1D1-BUILD1` | full build | pending |
| `E1-P1D1-TEST1` | core producer behavior tests | pending |
| `E1-P1D1-REVIEW1` | Supervisor PASS | pending |
| `E1-P1D1-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1D1-COMMIT1` | isolated slice commit | pending |
| `E1-P1D2-IMPACT1` | resolution/projection producer migration impact | pending |
| `E1-P1D2-SRC1` | resolution/projection explicit-operation diff | pending |
| `E1-P1D2-BUILD1` | full build | pending |
| `E1-P1D2-TEST1` | resolution/projection behavior tests | pending |
| `E1-P1D2-REVIEW1` | Supervisor PASS | pending |
| `E1-P1D2-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1D2-COMMIT1` | isolated slice commit | pending |
| `E1-P1D3-IMPACT1` | ancillary producer migration impact | pending |
| `E1-P1D3-SRC1` | ancillary explicit-operation diff | pending |
| `E1-P1D3-BUILD1` | full build | pending |
| `E1-P1D3-TEST1` | ancillary behavior tests | pending |
| `E1-P1D3-REVIEW1` | Supervisor PASS | pending |
| `E1-P1D3-DETECT1` | Anvien detect-changes before commit | pending |
| `E1-P1D3-COMMIT1` | isolated slice commit | pending |

## E2 - P2 Evidence

Matching plan item(s): `P2-A` through `P2-G`, including every ordered sub-slice in the P2 checklist.

Every row below is reserved evidence only until its exact command, artifact, result, Supervisor verdict, detect-changes result when applicable, and commit are recorded. P2 slices never advance automatically.

| Slice | Required evidence set | Status |
|-------|-----------------------|--------|
| P2-A | `E2-P2A-IMPACT1`, `E2-P2A-SRC1`, `E2-P2A-BUILD1`, `E2-P2A-TEST1`, `E2-P2A-REVIEW1`, `E2-P2A-DETECT1`, `E2-P2A-COMMIT1` | pending |
| P2-A1 | `E2-P2A1-MATRIX1`, `E2-P2A1-MATRIXREVIEW1`, `E2-P2A1-R01..E2-P2A1-R195`, `E2-P2A1-REVIEW1`, `E2-P2A1-COMMIT1` | pending |
| P2-A2 | `E2-P2A2-IMPACT1`, `E2-P2A2-SRC1`, `E2-P2A2-BUILD1`, `E2-P2A2-TEST1`, `E2-P2A2-S0GUARD1`, `E2-P2A2-REVIEW1`, `E2-P2A2-DETECT1`, `E2-P2A2-COMMIT1` | pending |
| P2-A3 | `E2-P2A3-IMPACT1`, `E2-P2A3-SRC1`, `E2-P2A3-BUILD1`, `E2-P2A3-TEST1`, `E2-P2A3-S1GUARD1`, `E2-P2A3-REVIEW1`, `E2-P2A3-DETECT1`, `E2-P2A3-COMMIT1` | pending |
| P2-A4 | `E2-P2A4-IMPACT1`, `E2-P2A4-SRC1`, `E2-P2A4-BUILD1`, `E2-P2A4-TEST1`, `E2-P2A4-S2GUARD1`, `E2-P2A4-REVIEW1`, `E2-P2A4-DETECT1`, `E2-P2A4-COMMIT1` | pending |
| P2-A5 | `E2-P2A5-IMPACT1`, `E2-P2A5-SRC1`, `E2-P2A5-BUILD1`, `E2-P2A5-TEST1`, `E2-P2A5-S3GUARD1`, `E2-P2A5-REVIEW1`, `E2-P2A5-DETECT1`, `E2-P2A5-COMMIT1` | pending |
| P2-A6 | `E2-P2A6-IMPACT1`, `E2-P2A6-SRC1`, `E2-P2A6-BUILD1`, `E2-P2A6-TEST1`, `E2-P2A6-S4GUARD1`, `E2-P2A6-REVIEW1`, `E2-P2A6-DETECT1`, `E2-P2A6-COMMIT1` | pending |
| P2-A7 | `E2-P2A7-IMPACT1`, `E2-P2A7-SRC1`, `E2-P2A7-BUILD1`, `E2-P2A7-TEST1`, `E2-P2A7-S5GUARD1`, `E2-P2A7-REVIEW1`, `E2-P2A7-DETECT1`, `E2-P2A7-COMMIT1` | pending |
| P2-A8 | `E2-P2A8-IMPACT1`, `E2-P2A8-SRC1`, `E2-P2A8-BUILD1`, `E2-P2A8-PLAY1`, `E2-P2A8-S6GUARD1`, `E2-P2A8-REVIEW1`, `E2-P2A8-DETECT1`, `E2-P2A8-COMMIT1` | pending |
| P2-A9 | `E2-P2A9-IMPACT1`, `E2-P2A9-SRC1`, `E2-P2A9-BUILD1`, `E2-P2A9-TEST1`, `E2-P2A9-S7GUARD1`, `E2-P2A9-REVIEW1`, `E2-P2A9-DETECT1`, `E2-P2A9-COMMIT1` | pending |
| P2-A10 | `E2-P2A10-IMPACT1`, `E2-P2A10-SRC1`, `E2-P2A10-BUILD1`, `E2-P2A10-TEST1`, `E2-P2A10-S8GUARD1`, `E2-P2A10-REVIEW1`, `E2-P2A10-DETECT1`, `E2-P2A10-COMMIT1` | pending |
| P2-A11 | `E2-P2A11-IMPACT1`, `E2-P2A11-SRC1`, `E2-P2A11-BUILD1`, `E2-P2A11-TEST1`, `E2-P2A11-S9GUARD1`, `E2-P2A11-REVIEW1`, `E2-P2A11-DETECT1`, `E2-P2A11-COMMIT1` | pending |
| P2-A12 | `E2-P2A12-IMPACT1`, `E2-P2A12-SRC1`, `E2-P2A12-BUILD1`, `E2-P2A12-TEST1`, `E2-P2A12-S10GUARD1`, `E2-P2A12-REVIEW1`, `E2-P2A12-DETECT1`, `E2-P2A12-COMMIT1` | pending |
| P2-A13 | `E2-P2A13-IMPACT1`, `E2-P2A13-SRC1`, `E2-P2A13-BUILD1`, `E2-P2A13-TEST1`, `E2-P2A13-S10GUARD1`, `E2-P2A13-REVIEW1`, `E2-P2A13-DETECT1`, `E2-P2A13-COMMIT1` | pending |
| P2-A14 | `E2-P2A14-IMPACT1`, `E2-P2A14-SRC1`, `E2-P2A14-BUILD1`, `E2-P2A14-TEST1`, `E2-P2A14-S11GUARD1`, `E2-P2A14-REVIEW1`, `E2-P2A14-DETECT1`, `E2-P2A14-COMMIT1` | pending |
| P2-A15 | `E2-P2A15-IMPACT1`, `E2-P2A15-SRC1`, `E2-P2A15-BUILD1`, `E2-P2A15-TEST1`, `E2-P2A15-S11GUARD1`, `E2-P2A15-REVIEW1`, `E2-P2A15-DETECT1`, `E2-P2A15-COMMIT1` | pending |
| P2-B | `E2-P2B-IMPACT1`, `E2-P2B-SRC1`, `E2-P2B-BUILD1`, `E2-P2B-TEST1`, `E2-P2B-REVIEW1`, `E2-P2B-DETECT1`, `E2-P2B-COMMIT1` | pending |
| P2-B1 | `E2-P2B1-IMPACT1`, `E2-P2B1-SRC1`, `E2-P2B1-BUILD1`, `E2-P2B1-TEST1`, `E2-P2B1-REVIEW1`, `E2-P2B1-DETECT1`, `E2-P2B1-COMMIT1` | pending |
| P2-B2 | `E2-P2B2-IMPACT1`, `E2-P2B2-SRC1`, `E2-P2B2-BUILD1`, `E2-P2B2-TEST1`, `E2-P2B2-REVIEW1`, `E2-P2B2-DETECT1`, `E2-P2B2-COMMIT1` | pending |
| P2-B3 | `E2-P2B3-IMPACT1`, `E2-P2B3-SRC1`, `E2-P2B3-BUILD1`, `E2-P2B3-TEST1`, `E2-P2B3-REVIEW1`, `E2-P2B3-DETECT1`, `E2-P2B3-COMMIT1` | pending |
| P2-B4 | `E2-P2B4-IMPACT1`, `E2-P2B4-SRC1`, `E2-P2B4-BUILD1`, `E2-P2B4-TEST1`, `E2-P2B4-REVIEW1`, `E2-P2B4-DETECT1`, `E2-P2B4-COMMIT1` | pending |
| P2-C | `E2-P2C-IMPACT1`, `E2-P2C-SRC1`, `E2-P2C-BUILD1`, `E2-P2C-TEST1`, `E2-P2C-REVIEW1`, `E2-P2C-DETECT1`, `E2-P2C-COMMIT1` | pending |
| P2-C1 | `E2-P2C1-IMPACT1`, `E2-P2C1-SRC1`, `E2-P2C1-BUILD1`, `E2-P2C1-TEST1`, `E2-P2C1-REVIEW1`, `E2-P2C1-DETECT1`, `E2-P2C1-COMMIT1` | pending |
| P2-C2 | `E2-P2C2-IMPACT1`, `E2-P2C2-SRC1`, `E2-P2C2-BUILD1`, `E2-P2C2-TEST1`, `E2-P2C2-REVIEW1`, `E2-P2C2-DETECT1`, `E2-P2C2-COMMIT1` | pending |
| P2-C3 | `E2-P2C3-IMPACT1`, `E2-P2C3-SRC1`, `E2-P2C3-BUILD1`, `E2-P2C3-TEST1`, `E2-P2C3-REVIEW1`, `E2-P2C3-DETECT1`, `E2-P2C3-COMMIT1` | pending |
| P2-C4 | `E2-P2C4-IMPACT1`, `E2-P2C4-SRC1`, `E2-P2C4-BUILD1`, `E2-P2C4-TEST1`, `E2-P2C4-REVIEW1`, `E2-P2C4-DETECT1`, `E2-P2C4-COMMIT1` | pending |
| P2-C5 | `E2-P2C5-IMPACT1`, `E2-P2C5-SRC1`, `E2-P2C5-BUILD1`, `E2-P2C5-TEST1`, `E2-P2C5-REVIEW1`, `E2-P2C5-DETECT1`, `E2-P2C5-COMMIT1` | pending |
| P2-C6 | `E2-P2C6-IMPACT1`, `E2-P2C6-SRC1`, `E2-P2C6-BUILD1`, `E2-P2C6-TEST1`, `E2-P2C6-REVIEW1`, `E2-P2C6-DETECT1`, `E2-P2C6-COMMIT1` | pending |
| P2-D | `E2-P2D-IMPACT1`, `E2-P2D-SRC1`, `E2-P2D-BUILD1`, `E2-P2D-TEST1`, `E2-P2D-REVIEW1`, `E2-P2D-DETECT1`, `E2-P2D-COMMIT1` | pending |
| P2-D1 | `E2-P2D1-IMPACT1`, `E2-P2D1-SRC1`, `E2-P2D1-BUILD1`, `E2-P2D1-TEST1`, `E2-P2D1-REVIEW1`, `E2-P2D1-DETECT1`, `E2-P2D1-COMMIT1` | pending |
| P2-D2 | `E2-P2D2-IMPACT1`, `E2-P2D2-SRC1`, `E2-P2D2-BUILD1`, `E2-P2D2-TEST1`, `E2-P2D2-REVIEW1`, `E2-P2D2-DETECT1`, `E2-P2D2-COMMIT1` | pending |
| P2-E | `E2-P2E-IMPACT1`, `E2-P2E-SRC1`, `E2-P2E-BUILD1`, `E2-P2E-TEST1`, `E2-P2E-REVIEW1`, `E2-P2E-DETECT1`, `E2-P2E-COMMIT1` | pending |
| P2-E1 | `E2-P2E1-IMPACT1`, `E2-P2E1-SRC1`, `E2-P2E1-BUILD1`, `E2-P2E1-PLAY1`, `E2-P2E1-REVIEW1`, `E2-P2E1-DETECT1`, `E2-P2E1-COMMIT1` | pending |
| P2-E2 | `E2-P2E2-BUILD1`, `E2-P2E2-S0BASE1`, `E2-P2E2-S1BASE1`, `E2-P2E2-S2BASE1`, `E2-P2E2-S3BASE1`, `E2-P2E2-S4BASE1`, `E2-P2E2-S5BASE1`, `E2-P2E2-S6BASE1`, `E2-P2E2-S7BASE1`, `E2-P2E2-S8BASE1`, `E2-P2E2-S9BASE1`, `E2-P2E2-S10BASE1`, `E2-P2E2-S11BASE1`, `E2-P2E2-MATRIX1`, `E2-P2E2-PLAY1`, `E2-P2E2-REVIEW1`, `E2-P2E2-DETECT1`, `E2-P2E2-COMMIT1` | pending |
| P2-F | `E2-P2F-IMPACT1`, `E2-P2F-SRC1`, `E2-P2F-BUILD1`, `E2-P2F-TEST1`, `E2-P2F-REVIEW1`, `E2-P2F-DETECT1`, `E2-P2F-COMMIT1` | pending |
| P2-F1 | `E2-P2F1-IMPACT1`, `E2-P2F1-SRC1`, `E2-P2F1-BUILD1`, `E2-P2F1-TEST1`, `E2-P2F1-REVIEW1`, `E2-P2F1-DETECT1`, `E2-P2F1-COMMIT1` | pending |
| P2-F2 | `E2-P2F2-IMPACT1`, `E2-P2F2-SRC1`, `E2-P2F2-BUILD1`, `E2-P2F2-TEST1`, `E2-P2F2-REVIEW1`, `E2-P2F2-DETECT1`, `E2-P2F2-COMMIT1` | pending |
| P2-F3 | `E2-P2F3-IMPACT1`, `E2-P2F3-SRC1`, `E2-P2F3-BUILD1`, `E2-P2F3-TEST1`, `E2-P2F3-REVIEW1`, `E2-P2F3-DETECT1`, `E2-P2F3-COMMIT1` | pending |
| P2-F4 | `E2-P2F4-IMPACT1`, `E2-P2F4-SRC1`, `E2-P2F4-BUILD1`, `E2-P2F4-TEST1`, `E2-P2F4-REVIEW1`, `E2-P2F4-DETECT1`, `E2-P2F4-COMMIT1` | pending |
| P2-F5 | `E2-P2F5-IMPACT1`, `E2-P2F5-SRC1`, `E2-P2F5-BUILD1`, `E2-P2F5-TEST1`, `E2-P2F5-REVIEW1`, `E2-P2F5-DETECT1`, `E2-P2F5-COMMIT1` | pending |
| P2-F6 | `E2-P2F6-BUILD1`, `E2-P2F6-FAULT1`, `E2-P2F6-RECOVERY1`, `E2-P2F6-REVIEW1`, `E2-P2F6-DETECT1`, `E2-P2F6-COMMIT1` | pending |
| P2-G | `E2-P2G-PREBASE1`, `E2-P2G-CANDIDATE1`, `E2-P2G-IMPACT1`, `E2-P2G-SRC1`, `E2-P2G-CUTOVER1`, `E2-P2G-BUILD1`, `E2-P2G-RUNTIME1`, `E2-P2G-PLAY1`, `E2-P2G-ROLLBACK1`, `E2-P2G-REVIEW1`, `E2-P2G-DETECT1`, `E2-P2G-COMMIT1` | pending |

### P2-A1 reader inventory and guard ownership

- `E2-P2A1-MATRIX1`: frozen source-derived matrix with exact path/function, truthful backend/layout, dispatcher/non-reader classification, later guard-owner slice, fixture, and expected transport failure.
- `E2-P2A1-MATRIXREVIEW1`: fresh source scan, matrix SHA-256, contiguous/unique row check, anchor existence, `rows_classified == rows_total`, `unassigned_rows == 0`, and `unlisted_readers == 0`.
- `E2-P2A1-R01..E2-P2A1-R195`: one exact classification proof per current seed row; if the fresh scan adds rows, continue the row/evidence numbering. These proofs establish source anchor, backend, surface tags, and assigned guard owner. They do not claim runtime guard PASS.

| Surface | Guard slice | Owner boundary | Runtime row-result evidence |
|---------|-------------|----------------|-----------------------------|
| `S0` | `P2-A2` | Graph JSON/repository metadata | `E2-P2A2-S0GUARD1` |
| `S1` | `P2-A3` | native Ladybug | `E2-P2A3-S1GUARD1` |
| `S2` | `P2-A4` | Go/fallback Cypher | `E2-P2A4-S2GUARD1` |
| `S3` | `P2-A5` | CLI | `E2-P2A5-S3GUARD1` |
| `S4` | `P2-A6` | MCP | `E2-P2A6-S4GUARD1` |
| `S5` | `P2-A7` | HTTP | `E2-P2A7-S5GUARD1` |
| `S6` | `P2-A8` | Web | `E2-P2A8-S6GUARD1` |
| `S7` | `P2-A9` | file-context cache | `E2-P2A9-S7GUARD1` |
| `S8` | `P2-A10` | HTTP/MCP resource cache | `E2-P2A10-S8GUARD1` |
| `S9` | `P2-A11` | embeddings | `E2-P2A11-S9GUARD1` |
| `S10` | `P2-A12` | global repository registry | `E2-P2A12-S10GUARD1` |
| `S10` | `P2-A13` | group registry/contracts | `E2-P2A13-S10GUARD1` |
| `S11` | `P2-A14` | process projections | `E2-P2A14-S11GUARD1` |
| `S11` | `P2-A15` | community/cluster projections | `E2-P2A15-S11GUARD1` |

Each guard slice records the exact matrix row IDs it executed. A parent command, router, selector, transport, or dispatcher never substitutes for an exact child row. Native and fallback query paths remain separate.

### P2-E2 pre-cutover canonical baselines

| Surface | Baseline evidence | Required comparison |
|---------|-------------------|---------------------|
| `S0` | `E2-P2E2-S0BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S1` | `E2-P2E2-S1BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S2` | `E2-P2E2-S2BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S3` | `E2-P2E2-S3BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S4` | `E2-P2E2-S4BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S5` | `E2-P2E2-S5BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S6` | `E2-P2E2-S6BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S7` | `E2-P2E2-S7BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S8` | `E2-P2E2-S8BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S9` | `E2-P2E2-S9BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S10` | `E2-P2E2-S10BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |
| `S11` | `E2-P2E2-S11BASE1` | canonical fields, one generation/epoch, `differing_records == 0`, and `orphan_refs == 0`; exact matrix-row union where applicable |

- `E2-P2E2-MATRIX1`: complete built-runtime execution result with `rows_passed == rows_total`, `unlisted_readers == 0`, no skipped row, and no parent-row substitution.
- `E2-P2E2-PLAY1`: reusable Playwright JSON+Markdown, screenshots/traces, and visual inspection for the exact S6 row union against the real built Docker runtime.
- Any S0-S11 failure reopens only the responsible already-committed storage/consumer/guard slice; P2-E2 does not patch code.

### P2 publication and cutover proofs

- `E2-P2F6-FAULT1` and `E2-P2F6-RECOVERY1`: complete staging/write/fsync/CAS/restart/concurrency/lease/GC fault matrix with exact prior-state hashes, pointers, generation vectors, and queryability after every failure.
- `E2-P2G-PREBASE1`: at least five identical pre-cutover v1 runs recording analyze median, Ladybug-load median, native-Cypher p95, fallback-Cypher p95, graph size, peak RSS, and the bound corpus/config/build/machine/cache policy.
- `E2-P2G-CANDIDATE1`: at least five staged non-active v2 runs with the identical methodology and per-metric delta before active CAS.
- `E2-P2G-PLAY1`: built Docker/Web supported, mismatch, and legacy-ambiguity evidence; an old client consumes zero v2 records.
- `E2-P2G-ROLLBACK1`: active-v2 rollback restores the prior queryable generation/registry/group vector without mixed cache or embedding state.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`, `P3-B`, `P3-C`

| Slice | Required evidence set | Status |
|-------|-----------------------|--------|
| P3-A | `E3-P3A-IMPACT1`, `E3-P3A-SRC1`, `E3-P3A-BUILD1`, `E3-P3A-TEST1`, `E3-P3A-REVIEW1`, `E3-P3A-DETECT1`, `E3-P3A-COMMIT1` | pending |
| P3-B | `E3-P3B-IMPACT1`, `E3-P3B-SRC1`, `E3-P3B-BUILD1`, `E3-P3B-TEST1`, `E3-P3B-REVIEW1`, `E3-P3B-DETECT1`, `E3-P3B-COMMIT1` | pending |
| P3-C | `E3-P3C-IMPACT1`, `E3-P3C-SRC1`, `E3-P3C-BUILD1`, `E3-P3C-TEST1`, `E3-P3C-REVIEW1`, `E3-P3C-DETECT1`, `E3-P3C-COMMIT1` | pending |

Expanded P3 slices:

| Slice | Required evidence set | Status |
|-------|-----------------------|--------|
| P3-B1 | `E3-P3B1-IMPACT1`, `E3-P3B1-SRC1`, `E3-P3B1-BUILD1`, `E3-P3B1-TEST1`, `E3-P3B1-REVIEW1`, `E3-P3B1-DETECT1`, `E3-P3B1-COMMIT1` | pending |
| P3-B2 | `E3-P3B2-IMPACT1`, `E3-P3B2-SRC1`, `E3-P3B2-BUILD1`, `E3-P3B2-TEST1`, `E3-P3B2-REVIEW1`, `E3-P3B2-DETECT1`, `E3-P3B2-COMMIT1` | pending |
| P3-B2A | `E3-P3B2A-IMPACT1`, `E3-P3B2A-SRC1`, `E3-P3B2A-BUILD1`, `E3-P3B2A-TEST1`, `E3-P3B2A-REVIEW1`, `E3-P3B2A-DETECT1`, `E3-P3B2A-COMMIT1` | pending |
| P3-C1 | `E3-P3C1-IMPACT1`, `E3-P3C1-SRC1`, `E3-P3C1-BUILD1`, `E3-P3C1-TEST1`, `E3-P3C1-REVIEW1`, `E3-P3C1-DETECT1`, `E3-P3C1-COMMIT1` | pending |
| P3-C1A | `E3-P3C1A-IMPACT1`, `E3-P3C1A-SRC1`, `E3-P3C1A-BUILD1`, `E3-P3C1A-TEST1`, `E3-P3C1A-REVIEW1`, `E3-P3C1A-DETECT1`, `E3-P3C1A-COMMIT1` | pending |
| P3-C1B | `E3-P3C1B-IMPACT1`, `E3-P3C1B-SRC1`, `E3-P3C1B-BUILD1`, `E3-P3C1B-TEST1`, `E3-P3C1B-REVIEW1`, `E3-P3C1B-DETECT1`, `E3-P3C1B-COMMIT1` | pending |
| P3-C1C | `E3-P3C1C-IMPACT1`, `E3-P3C1C-SRC1`, `E3-P3C1C-BUILD1`, `E3-P3C1C-TEST1`, `E3-P3C1C-REVIEW1`, `E3-P3C1C-DETECT1`, `E3-P3C1C-COMMIT1` | pending |
| P3-C1D | `E3-P3C1D-IMPACT1`, `E3-P3C1D-SRC1`, `E3-P3C1D-BUILD1`, `E3-P3C1D-TEST1`, `E3-P3C1D-REVIEW1`, `E3-P3C1D-DETECT1`, `E3-P3C1D-COMMIT1` | pending |
| P3-C1E | `E3-P3C1E-IMPACT1`, `E3-P3C1E-SRC1`, `E3-P3C1E-BUILD1`, `E3-P3C1E-TEST1`, `E3-P3C1E-REVIEW1`, `E3-P3C1E-DETECT1`, `E3-P3C1E-COMMIT1` | pending |
| P3-C1F | `E3-P3C1F-IMPACT1`, `E3-P3C1F-SRC1`, `E3-P3C1F-BUILD1`, `E3-P3C1F-PLAY1`, `E3-P3C1F-REVIEW1`, `E3-P3C1F-DETECT1`, `E3-P3C1F-COMMIT1` | pending |
| P3-C1G | `E3-P3C1G-IMPACT1`, `E3-P3C1G-SRC1`, `E3-P3C1G-BUILD1`, `E3-P3C1G-TEST1`, `E3-P3C1G-REVIEW1`, `E3-P3C1G-DETECT1`, `E3-P3C1G-COMMIT1` | pending |
| P3-C1H | `E3-P3C1H-IMPACT1`, `E3-P3C1H-SRC1`, `E3-P3C1H-BUILD1`, `E3-P3C1H-TEST1`, `E3-P3C1H-REVIEW1`, `E3-P3C1H-DETECT1`, `E3-P3C1H-COMMIT1` | pending |
| P3-C1I | `E3-P3C1I-IMPACT1`, `E3-P3C1I-SRC1`, `E3-P3C1I-BUILD1`, `E3-P3C1I-TEST1`, `E3-P3C1I-REVIEW1`, `E3-P3C1I-DETECT1`, `E3-P3C1I-COMMIT1` | pending |
| P3-C2 | `E3-P3C2-ORACLE1`, `E3-P3C2-TARGET1`, `E3-P3C2-BOUNDARY1`, `E3-P3C2-REVIEW1`, `E3-P3C2-DETECT1`, `E3-P3C2-COMMIT1` | pending |

P3 target evidence must name all six bounded bindings and preserve exact declaration/binding/range/scope/reference counts without copying target source.

`E3-P3C2-DETECT1` records the Anvien-side change/boundary check before the validation-artifact commit; `E3-P3C2-COMMIT1` records that isolated commit and confirms no target artifact was committed.

## E4 - P4 Evidence

Matching plan item(s): `P4-A`, `P4-B`, `P4-C`

| Slice | Required evidence set | Status |
|-------|-----------------------|--------|
| P4-A | `E4-P4A-IMPACT1`, `E4-P4A-SRC1`, `E4-P4A-BUILD1`, `E4-P4A-TEST1`, `E4-P4A-REVIEW1`, `E4-P4A-DETECT1`, `E4-P4A-COMMIT1` | pending |
| P4-B | `E4-P4B-IMPACT1`, `E4-P4B-SRC1`, `E4-P4B-BUILD1`, `E4-P4B-TEST1`, `E4-P4B-REVIEW1`, `E4-P4B-DETECT1`, `E4-P4B-COMMIT1` | pending |
| P4-C | `E4-P4C-IMPACT1`, `E4-P4C-SRC1`, `E4-P4C-BUILD1`, `E4-P4C-TEST1`, `E4-P4C-REVIEW1`, `E4-P4C-DETECT1`, `E4-P4C-COMMIT1` | pending |

Expanded P4 slices:

| Slice | Required evidence set | Status |
|-------|-----------------------|--------|
| P4-B1 | `E4-P4B1-IMPACT1`, `E4-P4B1-SRC1`, `E4-P4B1-BUILD1`, `E4-P4B1-TEST1`, `E4-P4B1-REVIEW1`, `E4-P4B1-DETECT1`, `E4-P4B1-COMMIT1` | pending |
| P4-C1 | `E4-P4C1-IMPACT1`, `E4-P4C1-SRC1`, `E4-P4C1-BUILD1`, `E4-P4C1-TEST1`, `E4-P4C1-REVIEW1`, `E4-P4C1-DETECT1`, `E4-P4C1-COMMIT1` | pending |
| P4-C1A | `E4-P4C1A-IMPACT1`, `E4-P4C1A-SRC1`, `E4-P4C1A-BUILD1`, `E4-P4C1A-TEST1`, `E4-P4C1A-REVIEW1`, `E4-P4C1A-DETECT1`, `E4-P4C1A-COMMIT1` | pending |
| P4-C1B | `E4-P4C1B-IMPACT1`, `E4-P4C1B-SRC1`, `E4-P4C1B-BUILD1`, `E4-P4C1B-TEST1`, `E4-P4C1B-REVIEW1`, `E4-P4C1B-DETECT1`, `E4-P4C1B-COMMIT1` | pending |
| P4-C1C | `E4-P4C1C-IMPACT1`, `E4-P4C1C-SRC1`, `E4-P4C1C-BUILD1`, `E4-P4C1C-TEST1`, `E4-P4C1C-REVIEW1`, `E4-P4C1C-DETECT1`, `E4-P4C1C-COMMIT1` | pending |
| P4-C1D | `E4-P4C1D-IMPACT1`, `E4-P4C1D-SRC1`, `E4-P4C1D-BUILD1`, `E4-P4C1D-TEST1`, `E4-P4C1D-REVIEW1`, `E4-P4C1D-DETECT1`, `E4-P4C1D-COMMIT1` | pending |
| P4-C1E | `E4-P4C1E-IMPACT1`, `E4-P4C1E-SRC1`, `E4-P4C1E-BUILD1`, `E4-P4C1E-TEST1`, `E4-P4C1E-REVIEW1`, `E4-P4C1E-DETECT1`, `E4-P4C1E-COMMIT1` | pending |
| P4-C1F | `E4-P4C1F-IMPACT1`, `E4-P4C1F-SRC1`, `E4-P4C1F-BUILD1`, `E4-P4C1F-PLAY1`, `E4-P4C1F-REVIEW1`, `E4-P4C1F-DETECT1`, `E4-P4C1F-COMMIT1` | pending |
| P4-C1G | `E4-P4C1G-IMPACT1`, `E4-P4C1G-SRC1`, `E4-P4C1G-BUILD1`, `E4-P4C1G-TEST1`, `E4-P4C1G-REVIEW1`, `E4-P4C1G-DETECT1`, `E4-P4C1G-COMMIT1` | pending |
| P4-C1H | `E4-P4C1H-IMPACT1`, `E4-P4C1H-SRC1`, `E4-P4C1H-BUILD1`, `E4-P4C1H-TEST1`, `E4-P4C1H-REVIEW1`, `E4-P4C1H-DETECT1`, `E4-P4C1H-COMMIT1` | pending |
| P4-C1I | `E4-P4C1I-IMPACT1`, `E4-P4C1I-SRC1`, `E4-P4C1I-BUILD1`, `E4-P4C1I-TEST1`, `E4-P4C1I-REVIEW1`, `E4-P4C1I-DETECT1`, `E4-P4C1I-COMMIT1` | pending |
| P4-C2 | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW1`, `E4-P4C2-DETECT1`, `E4-P4C2-COMMIT1` | pending |

P4 projection evidence must prove `21/21` bounded direct exports against the exact `S0`–`S11` matrix: ScopeIR/source oracle feeds `S0` Graph JSON, `S1` native Ladybug Cypher, `S2` fallback Cypher, `S3` CLI, `S4` MCP, `S5` HTTP, `S6` Web, `S7`/`S8` caches, `S9` embeddings, `S10` registries/groups, and `S11` process/community projections. Every applicable row compares the canonical fields from the plan and reports both `differing_records == 0` and `orphan_refs == 0`; access visibility remains separate. P2 must bind each row to the same canonical source manifest before P4 can claim parity.

`E4-P4C2-DETECT1` records the Anvien-side change/boundary check before the validation-artifact commit; `E4-P4C2-COMMIT1` records that isolated commit and confirms no target artifact was committed.

## E5 - P5 Evidence

Matching plan item(s): `P5-A`, `P5-B`, `P5-C`, `P5-D`

| Slice | Required evidence set | Status |
|-------|-----------------------|--------|
| P5-A | `E5-P5A-IMPACT1`, `E5-P5A-SRC1`, `E5-P5A-BUILD1`, `E5-P5A-ORACLE1`, `E5-P5A-REVIEW1`, `E5-P5A-DETECT1`, `E5-P5A-COMMIT1` | pending |
| P5-B | `E5-P5B-IMPACT1`, `E5-P5B-SRC1`, `E5-P5B-BUILD1`, `E5-P5B-TEST1`, `E5-P5B-BENCH1`, `E5-P5B-REVIEW1`, `E5-P5B-DETECT1`, `E5-P5B-COMMIT1` | pending |
| P5-C | `E5-P5C-IMPACT1`, `E5-P5C-SRC1`, `E5-P5C-BUILD1`, `E5-P5C-TEST1`, `E5-P5C-VECTOR1`, `E5-P5C-BENCH1`, `E5-P5C-REVIEW1`, `E5-P5C-DETECT1`, `E5-P5C-COMMIT1` | pending |
| P5-D | `E5-P5D-IMPACT1`, `E5-P5D-SRC1`, `E5-P5D-BUILD1`, `E5-P5D-TARGET1`, `E5-P5D-PARITY1`, `E5-P5D-BOUNDARY1`, `E5-P5D-ORACLE1`, `E5-P5D-REVIEW1`, `E5-P5D-DETECT1`, `E5-P5D-COMMIT1` | pending |

P5 vector/benchmark proofs:

- `E5-P5C-VECTOR1`: named/anonymous default, alias, type-only, namespace, merge/overload, precedence, cycle, ambiguity, and status expected-vector manifest.
- `E5-P5C-BENCH1`: hop/star-edge/candidate/SCC numeric-budget measurements with in-budget `100%` correctness and explicit over-limit `budget_exceeded`.
- `E5-P5D-IMPACT1`: fresh P5-D emission/consumer impact.
- `E5-P5D-BOUNDARY1`: target pre/post source/worktree/hash/graph-path/artifact boundary.
- `E5-P5D-ORACLE1`: independent target barrel source/TypeScript oracle.
- `E5-P5D-DETECT1`: Anvien detect-changes before commit.
- `E5-P5D-COMMIT1`: isolated P5-D commit.

P5 target evidence must prove two syntactic consumer-to-barrel dependencies, two terminal call bindings, complete hop proofs, and zero corresponding false gaps.

## E6 - P6 Evidence

Matching plan item(s): `P6-A`, `P6-B`, `P6-C1`, `P6-C2`, `P6-C3`, `P6-D`

| Slice | Required evidence set | Status |
|-------|-----------------------|--------|
| P6-A | `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-BUILD1`, `E6-P6A-TEST1`, `E6-P6A-REVIEW1`, `E6-P6A-DETECT1`, `E6-P6A-COMMIT1` | pending |
| P6-B | `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-REVIEW1`, `E6-P6B-DETECT1`, `E6-P6B-COMMIT1` | pending |
| P6-C1 | `E6-P6C1-IMPACT1`, `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-ORACLE1`, `E6-P6C1-REVIEW1`, `E6-P6C1-DETECT1`, `E6-P6C1-COMMIT1` | pending |
| P6-C2 | `E6-P6C2-IMPACT1`, `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-BENCH1`, `E6-P6C2-REVIEW1`, `E6-P6C2-DETECT1`, `E6-P6C2-COMMIT1` | pending |
| P6-C3 | `E6-P6C3-IMPACT1`, `E6-P6C3-SRC1`, `E6-P6C3-BUILD1`, `E6-P6C3-TEST1`, `E6-P6C3-STATUS1`, `E6-P6C3-REVIEW1`, `E6-P6C3-DETECT1`, `E6-P6C3-COMMIT1` | pending |
| P6-D | `E6-P6D-IMPACT1`, `E6-P6D-SRC1`, `E6-P6D-BUILD1`, `E6-P6D-TARGET1`, `E6-P6D-BOUNDARY1`, `E6-P6D-ORACLE1`, `E6-P6D-REVIEW1`, `E6-P6D-DETECT1`, `E6-P6D-COMMIT1` | pending |

P6 status/security proofs:

- `E6-P6C3-STATUS1`: every code in the P6 status matrix has positive/negative behavior evidence and mechanical graph-health mapping.
- `E6-P6C2-BENCH1`: catalog/package file-byte/cache/external-node/zero-catalog-File budget measurements with explicit in-budget and over-limit fixtures.
- `E6-P6D-BOUNDARY1`: target pre/post source/worktree/hash/graph-path/artifact boundary.
- `E6-P6D-ORACLE1`: independent Promise/Math source/TS oracle.

P6 evidence must include catalog provenance/integrity/version, project profile selection, language isolation, local package present/absent cases, P6-C1 immutable candidates, P6-C2 authorization/materialization results, the exact P6-C3 status matrix, lazy ExternalSymbol materialization, and exact target Promise/Math `resolved_external` outcomes (`3/3`).

## E7 - P7 Evidence

Matching plan item(s): `P7-A`, `P7-B`, `P7-C`

| Reserved evidence | Required proof | Status |
|-------------------|----------------|--------|
| `E7-P7A-BUILD1` | final full build before validation | pending |
| `E7-P7A-DETERMINISM1` | five canonical set-equality runs | pending |
| `E7-P7A-VERSION1` | every reader mismatch matrix | pending |
| `E7-P7A-FAULT1` | failure atomicity and old-generation query proof | pending |
| `E7-P7A-REVIEW1` | Supervisor PASS | pending |
| `E7-P7A-DETECT1` | Anvien-side change/boundary detection before validation-artifact commit | pending |
| `E7-P7A-COMMIT1` | isolated validation-artifact commit | pending |
| `E7-P7B-TARGET1` | target generation and bounded fact inventory | pending |
| `E7-P7B-ORACLE1` | independent source/TypeScript oracle comparison | pending |
| `E7-P7B-BOUNDARY1` | target pre/post contamination and scanner-quarantine check | pending |
| `E7-P7B-REVIEW1` | Supervisor PASS | pending |
| `E7-P7B-DETECT1` | Anvien-side change/boundary check before validation-artifact commit | pending |
| `E7-P7B-COMMIT1` | isolated Anvien evidence commit with no target artifact | pending |
| `E7-P7C-BUILD1` | packaged/Docker full build | pending |
| `E7-P7C-RUNTIME1` | CLI/MCP/HTTP/Web/Ladybug parity | pending |
| `E7-P7C-PLAY1` | reusable Playwright JSON+MD and visual inspection | pending |
| `E7-P7C-BENCH1` | five-run final performance/capacity measurements | pending |
| `E7-P7C-NATIVEBENCH1` | five-run final native-Cypher query p95 and regression | pending |
| `E7-P7C-FALLBACKBENCH1` | five-run final fallback-Cypher query p95 and regression | pending |
| `E7-P7C-REVIEW1` | Supervisor PASS | pending |
| `E7-P7C-DETECT1` | Anvien-side change/boundary detection before validation-artifact commit | pending |
| `E7-P7C-COMMIT1` | isolated runtime/parity/benchmark artifact commit | pending |

P7 parity evidence uses one row per exact surface, not a six-surface aggregate:

| Surface | Canonical fields checked | Evidence |
|---------|--------------------------|----------|
| `S0` Graph JSON | IDs, labels, paths, ranges, meanings, export/access, generation, provenance, external refs | `E7-P7C-S0-PARITY1` |
| `S1` Ladybug native Cypher | same node/relationship fields and closure | `E7-P7C-S1-PARITY1` |
| `S2` fallback Cypher | same node/relationship fields and closure | `E7-P7C-S2-PARITY1` |
| `S3` CLI | exact union of every current reader-matrix row tagged `S3`; no prose subset | `E7-P7C-S3-PARITY1` |
| `S4` MCP | exact union of every current reader-matrix row tagged `S4`; no prose subset | `E7-P7C-S4-PARITY1` |
| `S5` HTTP | exact union of every current reader-matrix row tagged `S5`; no router/catch-all substitution | `E7-P7C-S5-PARITY1` |
| `S6` Web | exact union of every current reader-matrix row tagged `S6`; built-runtime graph fields, lifecycle/stream contracts, and mismatch state | `E7-P7C-S6-PARITY1` |
| `S7` file-context caches | generation/config/catalog keys and canonical records | `E7-P7C-S7-PARITY1` |
| `S8` HTTP/MCP resource caches | generation/config/catalog keys and canonical records | `E7-P7C-S8-PARITY1` |
| `S9` embeddings | node/generation/dimension/hash refs | `E7-P7C-S9-PARITY1` |
| `S10` registries/groups | repo/group contract records, active epoch, staleness | `E7-P7C-S10-PARITY1` |
| `S11` process/community | source-anchored memberships and ordering | `E7-P7C-S11-PARITY1` |

Each row must report `differing_records == 0` and `orphan_refs == 0`; native and fallback Cypher are never merged into one result.

## E8 - P8 Evidence

Matching plan item(s): `P8-A`, `P8-B`, `P8-C`

- `E8-P8A-REVIEW1`: pending final Supervisor verdict on the full implemented plan.
- `E8-P8B-CLEAN1`: pending dead-work inventory and cleanup verdict.
- `E8-P8B-REVIEW1`: pending Supervisor verdict on dead-work cleanup.
- `E8-P8C-BUILD1`: pending final full build/runtime validation.
- `E8-P8C-RUNTIME1`: pending real built runtime startup and mounted-state validation.
- `E8-P8C-PLAY1`: pending final reusable Playwright JSON+Markdown/visual evidence when public runtime is in scope.
- `E8-P8C-REGEN1`: pending generated-output/source-of-truth hash and regeneration evidence.
- `E8-P8C-DETECT1`: pending final Anvien detect-changes evidence.
- `E8-P8C-LEDGER1`: pending final plan/evidence/benchmark/actual-status reconciliation evidence.
- `E8-P8C-COMMIT1`: pending final closure commit/worktree evidence.

## Closure Evidence

The plan set itself is created, but implementation closure is intentionally pending. No source fix, build, QA, detect-changes, or implementation commit is claimed by this ledger.
