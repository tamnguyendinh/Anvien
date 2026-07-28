# Anvien Graph Identity Resolution v2 Multi-Plan Authoring Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-actual-status.md`

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
- Reference exact evidence IDs from `actual-status.md` and `benchmark.md`; avoid referencing only broad section IDs such as `E1`.
- Use ranges such as `E0-P0A-FD1..E0-P0A-FD17` only for compact inventory summaries; use exact IDs when a specific status decision depends on a specific fact.
- If nearby plans already use a clear local evidence naming style, follow that style instead of inventing a new one.

Examples only:

- `E0-P0A-SRC1`
- `E0-P0A-GRAPH1`
- `E1-P1A-ROADMAP1`
- `E2-P2B-MAP1`
- `E3-P3B-DETECT1`

Evidence sections must follow the plan phases:

- `E0` corresponds to `P0`.
- `E1` corresponds to `P1`.
- `E2` corresponds to `P2`.
- `E3` corresponds to `P3`.
- Use exact evidence IDs inside each section, not broad section IDs as proof.
- Each evidence section must name the plan phase or checklist item it supports.
- Do not invent fixed evidence categories; record the evidence required by the matching plan phase.
- For cross-child evidence after the split, qualify every reference with the child slug and exact child-local evidence ID.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

### User and rule authority

- `E0-P0A-USER1`: The user explicitly requested one plan that will guide a later conversion of the existing large plan into a multi-plan campaign. The current request does not authorize executing that conversion.
- `E0-P0A-USER2`: The user fixed the decomposition rule: every existing implementation phase becomes one child plan; the number of children follows this plan's phase structure rather than the sample campaign's count.
- `E0-P0A-USER3`: The user required every child plan to be a complete plan with its own P0 and Pn lifecycle; a child cannot be treated as a fragment that omits those sections.
- `E0-P0A-USER4`: The user required one primary responsibility per file while allowing one file to link to multiple modules or files when those links serve that responsibility.
- `E0-P0A-BOUNDARY1`: The active repository rules and user direction keep all reports/plans in `E:\Anvien`, forbid copying the target repository into Anvien, and forbid writing plan/investigation artifacts into `E:\cheapapp.org`.

### Repository and Anvien basis

- `E0-P0A-GIT1`: `git status --short; git rev-parse HEAD; git branch --show-current` from `E:\Anvien` returned an empty short status, HEAD `68811c1643b604573e70551c7d4becb46e6ebbd8`, and branch `master` before this authoring plan set was created.
- `E0-P0A-GRAPH1`: `anvien analyze E:\Anvien --force --json` completed before graph-based inspection with 1,506 files, 676 parsed code files, 0 failed/unsupported/unknown code files, 84,548 nodes, and 123,388 relationships. Graph path: `E:\Anvien\.anvien\graph.json`.
- `E0-P0A-STATUS1`: `anvien status` reported repository `E:\Anvien`, indexed commit `68811c1`, current commit `68811c1`, status up to date.
- `E0-P0A-REPOS1`: `anvien list` identifies registry name `Anvien` at `E:\Anvien` and separately identifies `cheapapp-accuracy-direct` at `E:\cheapapp.org`; the repositories are not merged or copied.
- `E0-P0A-FD1`: `anvien file-detail "docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md" --repo E:\Anvien --json` failed truthfully with `file ... not found in repo Anvien`. The docs plan has no graph relationship count to report.
- `E0-P0A-QUERY1`: `anvien query files "graph identity resolution plan" --repo Anvien` returned source-code matches rather than the docs plan, corroborating that semantic graph queries do not represent this documentation target. Disk structure, hashes, links, and Git scope are therefore the valid nearest boundary for this docs-only plan.
- `E0-P0A-TEMPLATE1`: The complete current planner `SKILL.md` and all four bundled templates were read before authoring. This plan set retains the required Plan, Evidence Ledger, Benchmark Ledger, and Actual Status structures.

### Legacy source snapshot

- `E0-P0A-SRC1`: The active source plan is `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md` with SHA-256 `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB`, 567,584 bytes, 5,467 total lines, and 5,291 nonblank lines.
- `E0-P0A-SRC2`: The legacy campaign folder initially contains exactly five artifacts: the plan (567,584 bytes), evidence ledger (49,106 bytes), benchmark ledger (27,929 bytes), actual-status ledger (35,530 bytes), and `index-reader-matrix.md` (36,657 bytes). It contains no roadmap and no numbered implementation child folders.
- `E0-P0A-STRUCT1`: The source implementation headings begin at P1 line 313, P2 line 899, P3 line 2,891, P4 line 3,703, P5 line 4,433, P6 line 4,689, and P7 line 5,031. Legacy P8 closure begins at line 5,209.
- `E0-P0A-SLICE1`: Legacy P1 contains 11 slices: `P1-A`, `P1-B`, `P1-C0`, `P1-C0A`, `P1-C0B`, `P1-C`, `P1-D`, `P1-D1`, `P1-D2`, `P1-D3`, `P1-E`.
- `E0-P0A-SLICE2`: Legacy P2 contains 42 slices: `P2-A`, `P2-A1`-`P2-A15`, `P2-B`, `P2-B1`-`P2-B4`, `P2-C`, `P2-C1`-`P2-C6`, `P2-D`, `P2-D1`, `P2-D2`, `P2-E`, `P2-E1`, `P2-E2`, `P2-F`, `P2-F1`-`P2-F6`, and `P2-G`.
- `E0-P0A-SLICE3`: Legacy P3 contains 17 slices: `P3-A`, `P3-B`, `P3-B1`, `P3-B2`, `P3-B2A`, `P3-C`, `P3-C1`, `P3-C1A`-`P3-C1I`, and `P3-C2`.
- `E0-P0A-SLICE4`: Legacy P4 contains 15 slices: `P4-A`, `P4-B`, `P4-B1`, `P4-C`, `P4-C1`, `P4-C1A`-`P4-C1I`, and `P4-C2`.
- `E0-P0A-SLICE5`: Legacy P5 contains four slices: `P5-A`, `P5-B`, `P5-C`, and `P5-D`.
- `E0-P0A-SLICE6`: Legacy P6 contains six slices: `P6-A`, `P6-B`, `P6-C1`, `P6-C2`, `P6-C3`, and `P6-D`.
- `E0-P0A-SLICE7`: Legacy P7 contains three slices: `P7-A`, `P7-B`, and `P7-C`.
- `E0-P0A-TOTAL1`: The exact implementation total is `11 + 42 + 17 + 15 + 4 + 6 + 3 = 98`. Legacy P8 contains `P8-A`, `P8-B`, and `P8-C` closure items and is excluded from the implementation-child count.
- `E0-P0A-AUTH1`: The legacy plan metadata states `draft / P0 complete / implementation not yet authorized`; it remains the active planning authority until a future verified cutover.

### Multi-plan precedent and derived decision

- `E0-P0A-SAMPLE1`: The inspected sample at `G:\Restaurant_manager\DOCS\plans\2026-07-02-prototype-ux-overhaul` contains one roadmap, one standard `00-plan-set-authoring` plan set, and numbered implementation child plan sets. The sample count is campaign-specific.
- `E0-P0A-SAMPLE2`: Sample implementation children each have a complete lifecycle: P0, one local P1 decomposed into slices, and Pn-A/Pn-B/Pn-C. This supports structure, not a fixed child count.
- `E0-P0A-DECISION1`: Applying the user's phase-to-child rule to the measured source yields exactly seven implementation children. P8 becomes each child's Pn closure and does not create child 08.
- `E0-P0A-DECISION2`: The authoring plan set is control material and is not counted among the seven implementation children. The future roadmap is coordination material and cannot replace a child standard plan set.
- `E0-P0A-OWNER1`: `index-reader-matrix.md` concerns the versioned-reader/cutover phase and therefore has child 02 as its sole mutation owner; other children may reference it inspect-only.
- `E0-P0A-SCOPE1`: No command used for P0 wrote to or inspected source content in `E:\cheapapp.org`; this plan's artifacts are confined to the approved Anvien plan directory.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`, `P1-B`

Status: not started. These IDs are reserved for evidence recorded during authorized execution; they do not assert that the roadmap or transformation snapshot has been produced.

- `E1-P1A-SNAPSHOT1`: future source hash, path, line/count, and five-artifact snapshot verification.
- `E1-P1A-GIT1`: future execution-start Git basis and scope state.
- `E1-P1A-MAP1`: future seven-phase and 98-source-ID transformation manifest.
- `E1-P1A-MAP2`: future proof of local-ID remap and P8-to-Pn distribution rules.
- `E1-P1B-ROADMAP1`: future roadmap path/hash and required-section evidence.
- `E1-P1B-INVENTORY1`: future seven-child/28-file inventory contract evidence.
- `E1-P1B-LINK1`: future roadmap link and slug validation evidence.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`, `P2-B`, `P2-C`, `P2-D`, `P2-E`, `P2-F`, `P2-G`

Status: not started. Each child authoring slice must replace its reserved entries with exact files, hashes, structural results, mapping results, qualified dependencies, and roadmap/actual-status updates.

- `E2-P2A-FILES1`, `E2-P2A-STRUCT1`, `E2-P2A-MAP1`: future child-01 four-file, completeness, and 11-slice mapping evidence.
- `E2-P2B-FILES1`, `E2-P2B-STRUCT1`, `E2-P2B-MAP1`, `E2-P2B-MATRIX1`: future child-02 four-file, completeness, 42-slice mapping, and single matrix-owner evidence.
- `E2-P2C-FILES1`, `E2-P2C-STRUCT1`, `E2-P2C-MAP1`: future child-03 four-file, completeness, and 17-slice mapping evidence.
- `E2-P2D-FILES1`, `E2-P2D-STRUCT1`, `E2-P2D-MAP1`: future child-04 four-file, completeness, and 15-slice mapping evidence.
- `E2-P2E-FILES1`, `E2-P2E-STRUCT1`, `E2-P2E-VECTOR1`, `E2-P2E-MAP1`: future child-05 four-file, completeness, semantic-vector ownership, and four-slice mapping evidence.
- `E2-P2F-FILES1`, `E2-P2F-STRUCT1`, `E2-P2F-STATUS1`, `E2-P2F-MAP1`: future child-06 four-file, completeness, status-contract ownership, and six-slice mapping evidence.
- `E2-P2G-FILES1`, `E2-P2G-STRUCT1`, `E2-P2G-DEPENDENCY1`, `E2-P2G-MAP1`, `E2-P2G-CLOSURE1`: future child-07 four-file, completeness, six-upstream-dependency, three-slice mapping, and distributed-closure evidence.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`, `P3-B`

Status: not started. Authority remains with the legacy plan until every required P3-B evidence item exists and records PASS.

- `E3-P3A-STRUCT1`: future 1-roadmap/7-folder/28-file/P0/P1/Pn/required-field validation.
- `E3-P3A-LINK1`: future local and cross-plan link validation.
- `E3-P3A-OWNER1`: future one-primary-responsibility and single mutable-owner validation.
- `E3-P3A-MAP1`: future exact `98 -> 98` bijection proof with zero missing/duplicate/extra mappings.
- `E3-P3A-FIELDS1`: future source/destination field and semantic-preservation proof.
- `E3-P3A-DIFF1`: future Git scope proof showing docs-only candidate output.
- `E3-P3B-SUPERVISOR1`: future unconditional Supervisor verdict against frozen candidate hashes.
- `E3-P3B-CUTOVER1`: future one-active-authority pre/post assertion and exact legacy metadata/pointer diff.
- `E3-P3B-DIFF1`: future final scoped Git diff evidence.
- `E3-P3B-DETECT1`: future Anvien detect-changes result before an authorized cutover commit.

## Closure Evidence

Use this section for final detect-changes, commit hash, and closure evidence when the plan reaches completion.

Status: not started. No multi-plan conversion, authority cutover, implementation commit, build, Docker run, browser session, Playwright run, or target-repository action is claimed by the creation of this authoring plan.

Required closure evidence during authorized execution:

- final Supervisor PASS after Pn-A;
- dead-work review and cleanup result after Pn-B;
- final roadmap/child/legacy hashes and exact file inventory;
- final `98 -> 98` mapping and zero-error structural/link/ownership results;
- docs-only Git diff/status and Anvien detect-changes result;
- authorized commit hash or an explicit record that commit authorization was not granted;
- confirmation that `E:\cheapapp.org`, production source, tests, runtime files, graph output, and repository root contain no plan-created changes.
