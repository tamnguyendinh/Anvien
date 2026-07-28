# Anvien Graph Identity Resolution v2 Multi-Plan Authoring Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-plan.md`
- Evidence: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-evidence.md`
- Benchmark: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-benchmark.md`
- Actual status: `docs/plans/2026-07-28-00-multi-plan-authoring/2026-07-28-00-multi-plan-authoring-actual-status.md`

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

### P1-A - Frozen source snapshot and transformation contract

- `E1-P1A-AUTH1`: The user explicitly ordered execution of the accepted authoring plan on 2026-07-28. This opens P1-A while preserving every docs-only and target do-not-touch boundary.
- `E1-P1A-GIT1`: Execution began from clean commit `7b6c8e575a8f4fced05900cc8d42faebab234987` after the guide plan and its initial Supervisor PASS report were committed.
- `E1-P1A-SNAPSHOT1`: The legacy plan still has SHA-256 `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB`, with 98 unique implementation slices distributed `11/42/17/15/4/6/3`; no source drift occurred.
- `E1-P1A-GRAPH1`: Fresh `anvien analyze E:\\Anvien --force --json` completed at commit `7b6c8e57` with 1,511 files, 676 parsed code files, 0 failed/unsupported/unknown code files, 84,617 nodes, and 123,457 relationships.
- `E1-P1A-FD1`: Fresh `file-detail` for this authoring plan reports a parsed Markdown docs file, low risk, zero related files, zero relationships, and no unresolved references.
- `E1-P1A-FD2`: Fresh `file-detail` still cannot locate the 567,584-byte legacy plan. The size-specific documentation-index limitation is recorded; source hash, line structure, and exact Markdown parsing remain the valid transformation boundary.

#### E1-P1A-MAP1 - Exact source-to-child crosswalk

| Source slice | Child slug | Local slice | Source line | Source title |
|--------------|------------|-------------|------------:|--------------|
| `P1-A` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-A` | 334 | Ratify graph identity and ownership contract |
| `P1-B` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-B` | 389 | Introduce range, DeclarationID, SymbolID, and SymbolRef types |
| `P1-C0` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-C0` | 444 | Preserve lossless declaration occurrences |
| `P1-C0A` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-C0A` | 499 | Define RelationshipID and lossless source-site aggregation |
| `P1-C0B` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-C0B` | 544 | Validate lossless graph decode and closure |
| `P1-C` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-C` | 589 | Build declaration-to-symbol identity mapping |
| `P1-D` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-D` | 644 | Introduce strict graph mutation operations and validation |
| `P1-D1` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-D1` | 699 | Migrate core graph producers to explicit operations |
| `P1-D2` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-D2` | 744 | Migrate resolution/projection producers to explicit operations |
| `P1-D3` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-D3` | 789 | Migrate ancillary/document/semantic producers to explicit operations |
| `P1-E` | `2026-07-28-01-graph-identity-contract-and-strict-construction` | `P1-E` | 834 | Emit and validate shadow identity v2 |
| `P2-A` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A` | 950 | Define the index compatibility manifest and failure contract |
| `P2-A1` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A1` | 1005 | Freeze the source-derived reader inventory and owner assignments |
| `P2-A2` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A2` | 1060 | Guard Graph JSON and repository-metadata readers |
| `P2-A3` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A3` | 1105 | Guard native Ladybug readers |
| `P2-A4` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A4` | 1150 | Guard Go/fallback Cypher readers |
| `P2-A5` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A5` | 1195 | Guard CLI readers and dispatch boundaries |
| `P2-A6` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A6` | 1240 | Guard MCP resources and tools |
| `P2-A7` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A7` | 1285 | Guard HTTP handlers and streams |
| `P2-A8` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A8` | 1330 | Guard Web readers, streams, and lifecycle clients |
| `P2-A9` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A9` | 1375 | Guard file-context cache readers |
| `P2-A10` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A10` | 1420 | Guard HTTP/MCP resource-cache readers |
| `P2-A11` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A11` | 1465 | Guard embedding readers and jobs |
| `P2-A12` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A12` | 1510 | Guard global repository-registry readers |
| `P2-A13` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A13` | 1555 | Guard group registry and contract readers |
| `P2-A14` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A14` | 1600 | Guard process projection readers |
| `P2-A15` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-A15` | 1645 | Guard community and cluster projection readers |
| `P2-B` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-B` | 1690 | Make Graph JSON v2 codec and decode closure-safe |
| `P2-B1` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-B1` | 1735 | Write the Ladybug v2 schema and CSV export deterministically |
| `P2-B2` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-B2` | 1780 | Load Ladybug v2 transactionally and fail closed |
| `P2-B3` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-B3` | 1825 | Project canonical v2 records through native Ladybug queries |
| `P2-B4` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-B4` | 1870 | Project canonical v2 records through the Go fallback query path |
| `P2-C` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-C` | 1915 | Remove semantic ID parsing from CLI readers |
| `P2-C1` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-C1` | 1960 | Remove semantic ID parsing from MCP resources and tools |
| `P2-C2` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-C2` | 2005 | Make file-context projections use explicit canonical fields |
| `P2-C3` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-C3` | 2050 | Make file-context cache records generation/config/catalog-bound |
| `P2-C4` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-C4` | 2095 | Make rename use source anchors instead of parsed IDs |
| `P2-C5` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-C5` | 2140 | Make the shared HTTP/MCP resource cache preserve canonical records |
| `P2-C6` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-C6` | 2185 | Make embedding references generation-qualified and ID-opaque |
| `P2-D` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-D` | 2230 | Make group contracts use generation-qualified opaque references |
| `P2-D1` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-D1` | 2275 | Make process projections source-anchored and ID-opaque |
| `P2-D2` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-D2` | 2320 | Make community projections source-anchored and ID-opaque |
| `P2-E` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-E` | 2365 | Expose version/generation and canonical fields through HTTP |
| `P2-E1` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-E1` | 2410 | Negotiate and render version/generation truthfully in Web |
| `P2-E2` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-E2` | 2455 | Freeze the pre-cutover S0-S11 canonical baseline |
| `P2-F` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-F` | 2500 | Stage immutable repo-local generation artifacts |
| `P2-F1` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-F1` | 2545 | Publish the repo-local active generation atomically |
| `P2-F2` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-F2` | 2590 | Publish cache and embedding namespaces by generation |
| `P2-F3` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-F3` | 2635 | Publish the global repository registry atomically |
| `P2-F4` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-F4` | 2680 | Publish group snapshots and member-generation vectors atomically |
| `P2-F5` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-F5` | 2725 | Enforce reader leases and lease-safe generation garbage collection |
| `P2-F6` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-F6` | 2770 | Run the complete publication failure-atomicity matrix |
| `P2-G` | `2026-07-28-02-versioned-persistence-and-v2-cutover` | `P1-G` | 2815 | Cut over to identity v2 and enforce legacy ambiguity |
| `P3-A` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-A` | 2918 | Add recursive binding-pattern facts and walker |
| `P3-B` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-B` | 2973 | Integrate variable-declaration binding contexts |
| `P3-B1` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-B1` | 3028 | Integrate parameter binding contexts |
| `P3-B2` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-B2` | 3073 | Integrate catch binding contexts |
| `P3-B2A` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-B2A` | 3118 | Integrate for-of/for-in binding contexts |
| `P3-C` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C` | 3163 | Project binding occurrences into the graph |
| `P3-C1` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1` | 3208 | Project binding JSON/Ladybug persistence adapters |
| `P3-C1A` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1A` | 3253 | Project binding CLI adapters |
| `P3-C1B` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1B` | 3298 | Project binding MCP adapters |
| `P3-C1C` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1C` | 3343 | Project binding file-context cache records |
| `P3-C1D` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1D` | 3388 | Project binding HTTP adapters |
| `P3-C1E` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1E` | 3433 | Project binding HTTP/MCP resource-cache records |
| `P3-C1F` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1F` | 3478 | Project binding Web adapters |
| `P3-C1G` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1G` | 3523 | Project binding embedding references |
| `P3-C1H` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1H` | 3568 | Project binding registry/group references |
| `P3-C1I` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C1I` | 3613 | Project binding process/community references |
| `P3-C2` | `2026-07-28-03-typescript-binding-pattern-extraction` | `P1-C2` | 3658 | Validate bindings against the real target |
| `P4-A` | `2026-07-28-04-typescript-export-semantics` | `P1-A` | 3728 | Add ExportFact and meaning contracts |
| `P4-B` | `2026-07-28-04-typescript-export-semantics` | `P1-B` | 3783 | Extract direct/default/alias/type-only export facts |
| `P4-B1` | `2026-07-28-04-typescript-export-semantics` | `P1-B1` | 3838 | Extract star/namespace/re-export syntax facts |
| `P4-C` | `2026-07-28-04-typescript-export-semantics` | `P1-C` | 3883 | Project export edge/schema records |
| `P4-C1` | `2026-07-28-04-typescript-export-semantics` | `P1-C1` | 3938 | Project export persistence/read adapters |
| `P4-C1A` | `2026-07-28-04-typescript-export-semantics` | `P1-C1A` | 3983 | Project export CLI adapters |
| `P4-C1B` | `2026-07-28-04-typescript-export-semantics` | `P1-C1B` | 4028 | Project export MCP adapters |
| `P4-C1C` | `2026-07-28-04-typescript-export-semantics` | `P1-C1C` | 4073 | Project export file-context cache records |
| `P4-C1D` | `2026-07-28-04-typescript-export-semantics` | `P1-C1D` | 4118 | Project export HTTP adapters |
| `P4-C1E` | `2026-07-28-04-typescript-export-semantics` | `P1-C1E` | 4163 | Project export HTTP/MCP resource-cache records |
| `P4-C1F` | `2026-07-28-04-typescript-export-semantics` | `P1-C1F` | 4208 | Project export Web adapters |
| `P4-C1G` | `2026-07-28-04-typescript-export-semantics` | `P1-C1G` | 4253 | Project export embedding references |
| `P4-C1H` | `2026-07-28-04-typescript-export-semantics` | `P1-C1H` | 4298 | Project export registry/group references |
| `P4-C1I` | `2026-07-28-04-typescript-export-semantics` | `P1-C1I` | 4343 | Project export process/community references |
| `P4-C2` | `2026-07-28-04-typescript-export-semantics` | `P1-C2` | 4388 | Validate exports against the real target |
| `P5-A` | `2026-07-28-05-module-export-and-reexport-resolution` | `P1-A` | 4469 | Build hash-bound TypeScript project/module request inputs |
| `P5-B` | `2026-07-28-05-module-export-and-reexport-resolution` | `P1-B` | 4524 | Build deterministic per-module export tables |
| `P5-C` | `2026-07-28-05-module-export-and-reexport-resolution` | `P1-C` | 4579 | Resolve re-exports, cycles, ambiguity, aliases, and meanings |
| `P5-D` | `2026-07-28-05-module-export-and-reexport-resolution` | `P1-D` | 4634 | Emit terminal edges/proofs and validate barrel consumers |
| `P6-A` | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | `P1-A` | 4705 | Add declaration-universe and project-profile boundary |
| `P6-B` | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | `P1-B` | 4760 | Build and verify the embedded TypeScript stdlib catalog |
| `P6-C1` | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | `P1-C1` | 4815 | Resolve declaration entrypoints into immutable candidates |
| `P6-C2` | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | `P1-C2` | 4860 | Authorize and materialize referenced external Symbols |
| `P6-C3` | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | `P1-C3` | 4905 | Finalize exhaustive resolution outcomes |
| `P6-D` | `2026-07-28-06-ambient-external-resolution-and-diagnostics` | `P1-D` | 4950 | Project resolver outcomes into graph-health diagnostics |
| `P7-A` | `2026-07-28-07-cross-surface-acceptance-and-target-validation` | `P1-A` | 5044 | Run determinism, closure, version, and failure-atomicity gates |
| `P7-B` | `2026-07-28-07-cross-surface-acceptance-and-target-validation` | `P1-B` | 5099 | Run bounded `cheapapp.org` in-place acceptance |
| `P7-C` | `2026-07-28-07-cross-surface-acceptance-and-target-validation` | `P1-C` | 5154 | Run full runtime/projection/performance acceptance |

- `E1-P1A-MAP2`: Crosswalk validation yields 7 source phases, 7 unique child owners, 98 source IDs, 98 unique local destinations, and zero duplicates. Legacy `P8-A/P8-B/P8-C` map by role to every child's `Pn-A/Pn-B/Pn-C`; they are not implementation rows and do not create child 08. The legacy plan remains active until the candidate campaign passes P3-B Supervisor review.

### P1-B - Roadmap authoring

- `E1-P1B-ROADMAP1`: The roadmap now resides at `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`; its creation snapshot was SHA-256 `889F6231B327A30D3FCD8B891CDA0F6CFE2F70209FE26C23139383AB5908AA23`, 16,923 bytes, and 180 lines before later status/path updates. It owns campaign coordination only: authority, invariants, inventory, order, handoffs, shared ownership, target boundary, status protocol, and acceptance gates.
- `E1-P1B-INVENTORY1`: The roadmap declares exactly seven implementation children, counts `11/42/17/15/4/6/3 = 98`, seven future P0/local-P1/Pn lifecycles, and 28 unique standard child filenames. No child 08 is inventoried.
- `E1-P1B-LINK1`: Existing source-plan, authoring-plan, and exact-crosswalk paths resolve on disk. Future child paths are deliberately code-form planned paths while their status is `not authored`, so the candidate roadmap contains no false live child link. Roadmap authority remains candidate and the legacy plan remains active.
- `E1-P1B-REDTEAM1`: Three bounded read-only red-team audits independently confirmed P1/P2, P3/P4, and P5-P7/P8 counts, order, ledger ownership, and handoffs. The roadmap records their migration hazards, including P3-B2A gate normalization, the P3-C1 title alias, P4's cross-child P1-A authority, nonnumeric S0-S11 adapter order, P5/P6 out-of-slice manifests, and validation-only P7.
- `E1-P1B-SCOPE1`: P1-B created only the campaign-root roadmap and updated the authoring ledgers. It created no child folder, did not edit the legacy plan, production/tests/runtime/graph output, and did not access `E:\cheapapp.org`.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`, `P2-B`, `P2-C`, `P2-D`, `P2-E`, `P2-F`, `P2-G`

Status: P2-A complete; P2-A1 accepted / commit pending; P2-B ready-after-commit. Each later child authoring slice must replace its reserved entries with exact files, hashes, structural results, mapping results, qualified dependencies, and roadmap/actual-status updates.

- `E2-P2A-FILES1`: Child 01 contains exactly four standard files. Accepted SHA-256 values after relationship/dependency refresh and final whitespace hygiene: plan `2804538FE1412C91D92891CF469AD4FB9A39EECDA4E92AED7CDF3FA6B878B343`; evidence `08EB245ECA3D7BA968FFD6F31F6E831600E0BED08422DC209493B844F28B626B`; benchmark `9889791A0F47E25BBD0E1F8B8CB3C8FC322E00CE82BD9BF1D5C429A60794842A`; actual status `C079F5CFE3702D16C502E79E0741184CE22C09E642939C2B4FF0EDEA0333E925`.
- `E2-P2A-STRUCT1`: Child 01 has P0 complete, one local P1, 11 implementation blocks with all required planner fields, 11 source-slice provenance fields, tailored Pn-A/Pn-B/Pn-C, a preserved architecture/decision/one-file owner annex, zero placeholders, and zero trailing whitespace.
- `E2-P2A-MAP1`: Exact mapping is 11 source IDs -> 11 unique local IDs in source order. All 78 source-required P1 evidence IDs exist in the child evidence ledger, appear in local slice acceptance, and have 11 preserved traceability rows. No production, legacy, matrix, or target content changed.
- `E2-P2A-FD1`: After a fresh Anvien analyze, child-01 plan file-detail reports parsed Markdown/docs, low risk, one related file through the roadmap's inbound `IMPORTS` relationship, zero unresolved, and a current index at `c444e8c4`.
- `E2-P2A-SUP1`: `reports/Supervisor/rp_supervisor_260728_131205_by_gpt-5-codex_multi_plan_p2a_child_01.md` records unconditional PASS after deterministic checks and a red-team rejection/resubmission cycle closed the false-upstream dependency invariant.
- `E2-P2A1-USER1`: The user corrected the filesystem contract during P2-B authoring: the legacy plan, split-authoring plan, and resulting multi-plan must be three independent sibling roots directly under `docs/plans/`; child plans belong only under the separate multi-plan root. This correction pauses P2-B and opens P2-A1.
- `E2-P2A1-INVENTORY1`: The verified pre-move legacy root is inside `E:\Anvien\docs\plans\` and has nine immediate entries: the four legacy standard files plus `index-reader-matrix.md` that must remain, and four misplaced campaign entries that must leave the root (one roadmap plus the nested authoring, child-01, and child-02 directories). The move set contains 13 files: one roadmap, four authoring files, four accepted child-01 files, and four unaccepted child-02 draft files.
- `E2-P2A1-MOVE1`: All 13 files were moved once, not copied. The old nested authoring/roadmap/child paths no longer exist. The legacy root retains exactly five immediate files and zero directories; `docs/plans/2026-07-28-00-multi-plan-authoring/` owns the four authoring files; `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/` owns one roadmap plus child-01 and child-02 directories. Child 02 remains an unaccepted draft and is excluded from the P2-A1 commit boundary.
- `E2-P2A1-GRAPH1`: Final fresh `anvien analyze E:\Anvien --force --json` after the move and roadmap draft-link correction completed with 1,523 scanned files, 676 parsed code files, 724 documents, zero failed/unsupported/unknown files, 84,763 graph nodes, and 123,607 relationships at indexed/current commit `b760d156`.
- `E2-P2A1-FD1`: Fresh file-detail at analyzed time `2026-07-28T06:53:43Z` reports all four structural targets parsed as low-risk Markdown with zero unresolved references: authoring plan has zero relations; roadmap has four outbound `IMPORTS` relationships to the accepted child-01 four-file set; child 01 has one inbound roadmap relationship; the unaccepted child-02 draft has zero relationships and therefore is not exposed as a live committed roadmap link before P2-B.
- `E2-P2A1-STRUCT1`: Deterministic post-move validation returned: three required sibling roots exist; legacy root `5 files / 0 dirs`; authoring root `4 files / 0 dirs`; multi-plan root `1 roadmap / 2 child dirs`; child 01 `4 files`; child 02 draft `4 files`; frozen legacy plan hash remains `365E17A7F7CD539426568A2874A1AD1231D120C5391BFE3D8B875D5F760A45FB`; matrix hash remains `A6FE5148341E048425E6240B403D21F37941299319FB194A24B5CBE5E4F97409`; child-01 content is path-only equivalent `4/4` against commit `b760d156` after normalizing the approved prefixes.
- `E2-P2A1-LINK1`: Active authoring/multi-plan artifacts contain zero stale old-path references. Five live Markdown links and all standard-file metadata paths were checked from their real directories with zero broken destinations. Child-02 paths remain code-form draft inventory, so the P2-A1 commit will not claim links to files excluded from that commit.
- `E2-P2A1-DIFF1`: Pre-staging Git scope contains exactly nine tracked old-path deletions and 13 destination files under the two approved new roots; four destination files are the deliberately unaccepted child-02 draft. Thirteen active Markdown files have zero trailing-whitespace findings and zero template/TODO/TBD placeholders; `git diff --check` reports no tracked whitespace error; scoped-status validation reports zero path outside the three approved plan roots. No production, test, runtime, graph-output, repository-root, or target path appears.
- `E2-P2A1-REDTEAM1`: A bounded read-only red-team initially rejected stale pre-move wording in the authoring Actual Status/Evidence ledgers. The resubmission corrected every same-invariant current-state surface, deterministic stale-wording validation returned zero matches, and the same red-team re-review returned unconditional PASS.
- `E2-P2A1-SUP1`: `reports/Supervisor/rp_supervisor_260728_140744_by_gpt-5-6-sol_multi_plan_p2a1_three_root_correction.md` records unconditional PASS for the three-root correction. P2-A1 is accepted; its isolated structural commit remains pending, and P2-B may open only after that commit.
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
