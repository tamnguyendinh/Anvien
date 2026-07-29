# Anvien TypeScript Binding-Pattern Extraction Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md`
- Source evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-evidence.md`
- Multi-plan roadmap: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`

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

### P3 semantic correction binding

The execution rows below are conjunctive extensions to the base P3 evidence rows. Every listed correction/dependency ID remains pending until independently proved.

| Plan job | Required correction evidence IDs | Status |
|---|---|---|
| P3-A | `E2-PNC-BINDING1A` | pending |
| P3-B | `E2-PNC-BINDING1B` | pending |
| P3-B2A | `E3-P3C-B2A-GATE1` | pending |
| P3-C | `E2-PNC-BINDING1A`, `E2-PNC-BINDING1B`, `E3-P3C-B2A-GATE1` | pending |
| P3-C1 | `E2-PNC-BINDING1E` | pending |
| P3-C2 | `E3-P3C2-IMPACT1`, `E3-P3C2-SRC1`, `E3-P3C2-BUILD1`, `E2-PNC-BINDING1C`, `E2-PNC-BINDING1D`, `E2-PNC-BINDING1E`, `E3-P3C-B2A-GATE1` | pending |

## Closure Evidence

Closure is pending; these are mandatory declarations, not proof. The full evidence row for every implementation slice remains conjunctive with its plan Acceptance.

| Evidence ID | Source slice | Local slice | Required proof | Commit | Owner decision | Successor opening condition | Freshness evidence | Status |
|-------------|--------------|-------------|---------------|--------|---------------|----------------------------|--------------------|--------|
| `E2-PNC-OVERLAY1` | P8-C | Pn-C | semantic overlay precedence and local-child-versus-campaign closure decision | pending | pending | local checkpoint only | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-OWNERSHIP1` | P8-C | Pn-C | complete 17-job ownership table with exact production/test/generated/fixture paths and no wildcard/TBD | pending | pending | all 17 jobs have rows | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-LEDGER1` | P8-C | Pn-C | every slice Acceptance bound to complete IMPACT/SRC/BUILD/TEST-oracle/REVIEW/DETECT/COMMIT row | pending | pending | no review-only closure | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-BINDING1` | P3-A..P3-C2 | Pn-C | binding correction index | pending | pending | all five subrows accepted | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-BINDING1A` | P3-A | P3-A | structured unsupported-pattern diagnostic schema and countable metric | pending | pending | diagnostic emitted independently of type inference | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-BINDING1B` | P3-B | P3-B | declaration emission remains present when type inference fails | pending | pending | no inference-dependent drop | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-BINDING1C` | P3-C/P3-C2 | P3-C2 | six named `.map()` sites at `6/6` with `ResolutionGap=0` | pending | pending | all six target rows | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-BINDING1D` | P3-C/P3-C2 | P3-C2 | nested same-name shadowing selects the expected SymbolID | pending | pending | source/symbol oracle agrees | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-BINDING1E` | P3-C1* | P3-C1 | imports counted exactly once with zero delta | pending | pending | all applicable S0-S11 rows | `E2-PNC-REFRESH1` | pending |
| `E3-P3C-B2A-GATE1` | P3-B2A | P3-C | accepted/committed P3-B2A for-of/for-in row union and zero pending loop contexts before P3-C/P3-C2 | pending | pending | hard predecessor satisfied | `E2-PNC-REFRESH1` | pending |
| `E3-P3-OWNERSHIP1` | P3-A..P3-C2 | P3 | exact P3 job production/test/generated/fixture ownership manifest | pending | pending | all named jobs have rows | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-REFRESH1` | P8-C | Pn-C | current HEAD/graph/file-detail snapshot and successor-status refresh-log/next-action/work-step update | pending | pending | Child 04 actual-status fresh | `anvien analyze E:\Anvien --force --json`; `git rev-parse HEAD` | pending |
| `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-P2A-MANIFEST1` | Child 02 P2-A | P0 | inspect-only authoritative manifest proof | n/a | pending | Child 03 consumes without reinterpretation | Child 02 qualified handoff | pending |
| `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-P2A-HANDSHAKE1` | Child 02 P2-A | P0 | inspect-only authoritative handshake proof | n/a | pending | Child 03 consumes without reinterpretation | Child 02 qualified handoff | pending |
| `2026-07-28-02-versioned-persistence-and-v2-cutover::E2-P2A-METADATA1` | Child 02 P2-A | P0 | inspect-only authoritative metadata inventory | n/a | pending | Child 03 consumes without reinterpretation | Child 02 qualified handoff | pending |
| `E2-PNC-NEXTSTATUS1` | P8-C | Pn-C | successor Child 04 actual-status refreshed from latest accepted evidence; refresh-log, next-action, and work-step rows updated | pending | pending | Child 04 may open only after exact qualified handoff | `E2-PNC-REFRESH1` | pending |
| `E2-PNC-HANDOFF1` | P8-C | Pn-C | qualified Child 03 local evidence, commit, predecessor gate, owner decision, and Child 04 opening condition | pending | pending | `2026-07-28-03-typescript-binding-pattern-extraction::E2-PNC-NEXTSTATUS1` accepted after refreshing Child 04 | `E2-PNC-REFRESH1` | pending |
