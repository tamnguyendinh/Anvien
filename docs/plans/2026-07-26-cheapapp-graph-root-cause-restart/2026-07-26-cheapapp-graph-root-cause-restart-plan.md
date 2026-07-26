# Cheapapp Graph Root-Cause Restart Investigation Plan

## Metadata

- Date: `2026-07-26`
- Status: `accepted-bounded`
- Plan: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-plan.md`
- Evidence: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-evidence.md`
- Benchmark: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-benchmark.md`
- Actual status: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-actual-status.md`

## Goal

Restart the graph-accuracy investigation from a clean evidence boundary. Analyze `E:\cheapapp.org` in place with Anvien's supported repository workflow, keep its graph/index under `E:\cheapapp.org\.anvien`, independently compare source truth with Anvien graph and command results, and trace each confirmed discrepancy back to its first divergence in the Anvien codebase at `E:\Anvien`. Produce an evidence-bounded root-cause synthesis without copying the target into Anvien, without accepting prior agent conclusions, and without designing or implementing remediation during this investigation.

## Rules

- Complete P0 actual status before any investigation slice proceeds.
- `E:\cheapapp.org` is the target source of truth. Apart from Anvien's normal repo-local graph/index under `E:\cheapapp.org\.anvien`, do not write reports, scripts, probes, fixtures, build output, or temporary files there.
- Never copy, clone, checkpoint, or otherwise import `E:\cheapapp.org` into `E:\Anvien`; analyze it in place and leave its graph/index at `E:\cheapapp.org\.anvien`.
- All plans, reports, evidence, benchmarks, probes, and debug artifacts belong under `E:\Anvien`; temporary debug artifacts must remain under `E:\Anvien\.tmp\`.
- Read Anvien command help before relying on command syntax or output behavior. Refresh the relevant graph before every graph-based Anvien query, validation, or report command.
- Prior reports and prior agent/subagent outputs are evidence pointers only. Reproduce every finding independently before classifying it.
- Every slice must compare an Anvien result with an independent source-side check against the real `E:\cheapapp.org` path.
- Do not edit production source, tests, SPEC files, generated graph output, or target source as remediation in this plan.
- Each delegated slice has one owner, one bounded question, one report, and one acceptance gate. Subagent results are untrusted until independently checked by the main agent and Supervisor.
- Preserve raw counts, samples, hashes, paths, ranges, command output, and limitations; do not compress evidence to make it look simpler.
- Classify findings as `confirmed wrong`, `bounded valid`, or `unresolved`; do not make a global accuracy claim from bounded cases.
- Update the checklist, evidence ledger, benchmark ledger, and actual-status file immediately as each state changes.
- Use file-detail and impact evidence before editing any Anvien symbol if a later remediation task is explicitly opened; no edit is authorized by this investigation plan.
- Supervisor must independently review the completed investigation record before any closure claim.

## Problem

## Owner Correction R2 (2026-07-26)

The owner clarified that the intended target repository is `E:\cheapapp.org`, not a nonexistent `E:\cheapapp`. This correction supersedes the initial P0 path interpretation. The graph/index must remain at `E:\cheapapp.org\.anvien` as Anvien's normal repo-local operational output. Investigation reports, ledgers, probes, benchmark files, and causal analysis remain under `E:\Anvien`; the target is not copied into Anvien. P0 must be rerun against the corrected target before P1 slices open.

The earlier P0 blocker report is retained as historical evidence of the path misunderstanding and is not a blocker for the corrected target.

Reports from an earlier, context-contaminated investigation described possible graph inaccuracies such as omitted or extra facts, incorrect identity and visibility, unresolved standard-library references, missing module bindings, projection/cardinality hazards, and downstream command or semantic distortions. Those results are not accepted as established facts for this restart. The current task is to determine, against the actual `E:\cheapapp.org` source and a freshly produced Anvien graph, which discrepancies are real, where each first divergence occurs, and what remains unproven.

## Scope

Target repository (read-only source of truth):

- Path: `E:\cheapapp.org`
- Commit, tree state, and source inventory: to be captured during P0 without changing the target.
- Anvien analysis: direct analysis of this path using the supported external-repository/index workflow discovered from command help and source/config inspection.

Anvien repository (artifact and root-cause workspace):

- Path: `E:\Anvien`
- Plan, evidence, benchmark, actual-status, reports, probes, and debug output are written here.
- Anvien source pipeline to inspect: discovery/scanner, parser/extractor, IR and identity, resolver/module binding, graph emission, database/projection, command consumers, and derived semantic/process surfaces.

Investigation surfaces, opened as narrow slices:

- target file discovery and File-node completeness;
- TypeScript/JavaScript extraction, bindings, visibility, scope, ranges, and identity;
- module, re-export, ambient/builtin, and resolver behavior;
- graph serialization, persistence, nullability, multiplicity, and cardinality;
- context/impact and other command projections;
- process, community, layer, and boundary derivation only where source and graph inputs can be compared;
- first-divergence tracing in the corresponding Anvien source owners.

## Non-Goals

- No copy, clone, checkpoint, or import of `E:\cheapapp.org` into `E:\Anvien`.
- No manual mutation of `E:\cheapapp.org` source or worktree state. Anvien's supported analyze operation may also regenerate ignored guidance files (`AGENTS.md`/`CLAUDE.md`); that side effect is recorded, not used for investigation artifacts, and is never manually reverted here. The graph/index remains under `E:\cheapapp.org\.anvien`.
- No production fix, parser rewrite, resolver redesign, graph optimization, database redesign, or speed diagnosis.
- No acceptance of prior reports as truth and no global graph-accuracy closure claim.
- No broad random audit or agent-selected substitute cases before the fixed slice questions are answered.
- No cleanup of unrelated user changes or prior artifacts outside this plan's own artifacts.
- No UI, Docker, or Playwright claim; these are source/graph investigation slices and those fields are N/A unless a real runtime boundary becomes part of a later explicitly approved slice.

## Requirements

- The plan uses the four standard planner files and keeps them synchronized.
- P0 records the target's exact commit/tree state, Anvien version/help, supported output/index location, graph identity/inventory, and proof that the target remained unchanged.
- Every source-to-graph slice has a direct Anvien command result, an independent source-side result, exact discrepancy samples, and a durable report under `E:\Anvien\reports\Supervisor\`.
- Every root-cause slice identifies the first divergence, owning source file/symbol, data or control boundary, evidence strength, and blast-radius warning where applicable.
- Parallel work is limited to independently bounded slices. The main agent continues a non-overlapping slice while subagents work and rechecks every subagent result.
- Benchmark ledger records graph generation, inventory, command duration, and accuracy/cardinality counts as measurements only; it does not turn the investigation into performance work.
- A slice is complete only when its report is readable, evidence IDs are stable, actual-status rows are refreshed, and Supervisor has accepted the bounded claim or recorded an evidence-backed blocker.

## Acceptance Criteria

- P0 is complete and proves the read-only target boundary and supported direct-analysis mechanism.
- Fresh graph/source comparisons identify confirmed, bounded-valid, and unresolved cases without relying on prior agent output.
- Each confirmed discrepancy has a reproducible source location, graph/command observation, first divergence in Anvien code, and bounded impact statement.
- No investigation artifact created by this plan exists in `E:\cheapapp.org`; the repo-local `.anvien` graph/index remains the only expected Anvien operational output there.
- The synthesis distinguishes structural source-to-graph facts from product-intent interpretations and records missing authority such as absent SPEC material where relevant.
- Supervisor independently reviews the complete plan/evidence/report set and passes the investigation record, or the plan remains explicitly blocked/rejected with required next evidence.

## Checklist

- [x] P0-A: Complete actual status before investigation slices. (owner-corrected target and repo-local graph boundary confirmed)
  - Goal: establish the real target state, supported Anvien external-repository workflow, artifact destination, and read-only guard before any graph analysis.
  - Work Steps: read exact command help; inspect relevant Anvien source/config for path and index output behavior; capture target HEAD/status/tree metadata; create the four plan files; run only the minimum safe baseline command; record post-command target status and all baseline evidence.
  - Implementation Gate: no source-to-graph slice or subagent assignment starts until the target write boundary and graph artifact location are proven.
  - Acceptance: actual status classifies the target, Anvien workspace, prior evidence, and blockers with exact evidence IDs; P0 decision is explicit.

### P1: Fresh Source-to-Graph Comparison Slices

- Phase Goal: reproduce bounded graph discrepancies from the real `E:\cheapapp.org` path and classify them independently.
- Phase Boundary:
  - In scope: fresh graph baseline, file discovery, TS/JS extraction and identity, resolver/module behavior, persistence/projection, and command/derived outputs.
  - Out of scope: fixes, optimization, target mutation, and unsupported product-semantic claims.
  - Dependencies: P0-A complete; each slice uses the exact graph identity recorded by P0.
- Phase Implementation Rule: do not implement `P1` directly. Run `P1-A`, verify and record it, then assign only the next bounded slice. The main agent may run a non-overlapping slice while a subagent works.
- Ordered Slice List:
  - P1-A: Fresh graph baseline and bidirectional file inventory.
  - P1-B: TypeScript/JavaScript extraction, visibility, binding, scope, and identity.
  - P1-C: Resolution, imports, re-exports, ambient/builtin, and module binding.
  - P1-D: Graph serialization/database projection, multiplicity, nullability, and cardinality.
  - P1-E: Context/impact and bounded derived-process/semantic projections.

- [x] P1-A: Fresh graph baseline and bidirectional file inventory. (bounded comparison complete; final bounded-record Supervisor PASS; no global file-coverage claim)
  - Goal: prove which exact target files Anvien sees and which File nodes it emits for the fresh graph.
  - Scope Boundary:
    - Editable: none; report and evidence artifacts in Anvien only.
    - Inspect-only: `E:\cheapapp.org` tree/Git metadata, Anvien scan output, graph metadata and File nodes.
    - Preserve-only: all target source and existing Anvien artifacts outside this plan.
    - Out of scope: symbol, relationship, resolver, semantic, and remediation conclusions.
  - Non-Goals: no scanner fix and no claim that file presence proves downstream correctness.
  - Pre-flight Questions:
    - Data source: direct target file inventory plus fresh Anvien graph File inventory.
    - Display permission: N/A; no UI.
    - DB read flow: read-only graph metadata/query flow only.
    - DB write flow: N/A; any Anvien-managed index output must be proven outside target.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A; source/graph audit only.
    - Playwright target: N/A.
    - Behavior test: bidirectional path-set comparison with exact samples.
    - Cleanup/quarantine: keep debug output under `E:\Anvien\.tmp\` and remove only dead artifacts created by this slice.
    - External side effects: none permitted in target.
    - N/A notes: no production edit and no runtime UI boundary.
  - Work Steps:
    1. Refresh the target graph using the supported direct-path command and save benchmark/summary output under Anvien; verify target status/tree is unchanged.
       - UI flow check: N/A.
       - DB/data flow check: prove graph identity and output ownership before reading graph data.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A because this is a read-only investigation slice with no implementation/UI.
       - Evidence target: `E1-P1A-GRAPH1..E1-P1A-GRAPH3`.
    2. Enumerate source and graph File paths in both directions, classify missing/extra/path-normalization cases, and write the Supervisor report in Anvien.
       - UI flow check: N/A.
       - DB/data flow check: preserve full counts and exact path samples.
       - Render location check: report is readable under `E:\Anvien\reports\Supervisor\`.
       - Mini QA for each completed implementation slice (MUST): N/A; source/graph comparison is the behavior check.
       - Evidence target: `E1-P1A-FILE1..E1-P1A-REPORT1`.
  - Implementation Gate:
    - P0-A proves direct analysis and artifact placement without target writes; graph metadata is tied to the captured target commit; no prior report is used as a verdict.
  - Acceptance:
    - Source: target path set is captured and independently reproducible.
    - Runtime/UI: N/A.
    - DB/data: graph File set is compared in both directions and every discrepancy is classified.
    - Behavior test: post-command target status/tree check passes.
    - Cleanup/quarantine: no artifact exists in target.
    - Evidence IDs: `E1-P1A-GRAPH1..E1-P1A-REPORT1`.
    - Actual-status rows refreshed: `P1-A`.
  - Evidence Targets: command help, target baseline/post-state, graph metadata/hash/inventory, source/graph path ledger, Supervisor report.
  - Actual-status Update: classify file discovery as correct, wrong, partial, or unresolved and update P1-B assumptions from measured paths.
  - Commit Boundary: no implementation commit; report and ledger artifacts only.

- [x] P1-B: TypeScript/JavaScript extraction, visibility, binding, scope, and identity. (fresh three-file oracle and P2-A identity/extractor trace Supervisor PASS; bounded record accepted)
  - Goal: compare independently enumerated declarations and binding facts with graph symbols for fixed reproducible target files selected from P1-A, without assuming prior cases are valid.
  - Scope Boundary:
    - Editable: none; reports/probes in Anvien only.
    - Inspect-only: selected target files, AST/source ranges, extractor IR/graph nodes, and identity facts.
    - Preserve-only: scanner and resolver behavior not needed to establish extraction facts.
    - Out of scope: module resolution and command ranking.
  - Non-Goals: no extractor or identity fix.
  - Pre-flight Questions:
    - Data source: direct source/AST enumeration plus Anvien graph/context output.
    - Display permission: N/A.
    - DB read flow: read-only graph facts.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: source declaration/binding ledger versus graph node/field ledger.
    - Cleanup/quarantine: probes and normalized ledgers under Anvien `.tmp` only.
    - External side effects: none.
    - N/A notes: no target copy or fixture injection.
  - Work Steps:
    1. Select exact files and symbols from fresh P1-A evidence, enumerate bindings/exports/scopes/ranges independently, and run the narrow Anvien context/file-detail queries after graph refresh.
       - UI flow check: N/A.
       - DB/data flow check: compare one physical source site to one or more graph facts without collapsing multiplicity.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A; read-only source/graph evidence.
       - Evidence target: `E1-P1B-SRC1..E1-P1B-GRAPH2`.
    2. Classify omissions, extras, wrong identity, wrong visibility, scope collision, and duplicate facts; write the bounded report.
       - UI flow check: N/A.
       - DB/data flow check: preserve ranges, source text, IDs, and counts.
       - Render location check: Anvien report path only.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E1-P1B-CMP1..E1-P1B-REPORT1`.
  - Implementation Gate:
    - P1-A identifies the exact target files and graph identity; all candidate findings have independent source ranges.
  - Acceptance:
    - Source: every claimed extraction discrepancy has a source-side ledger entry.
    - Runtime/UI: N/A.
    - DB/data: graph field/identity comparison is complete for the bounded files.
    - Behavior test: exact source-site and node comparison is reproducible.
    - Cleanup/quarantine: no target artifacts.
    - Evidence IDs: `E1-P1B-SRC1..E1-P1B-REPORT1`.
    - Actual-status rows refreshed: `P1-B`.
  - Evidence Targets: source/AST ledger, Anvien file-detail/context output, normalized identity comparison, report.
  - Actual-status Update: update P1-B and P2-A assumptions with confirmed versus unproven extraction boundaries.
  - Commit Boundary: no implementation commit; report and ledger artifacts only.

- [x] P1-C: Resolution, imports, re-exports, ambient/builtin, and module binding. (five fixed sites compared; bounded wrong result recorded; P2-B causal trace independently reviewed PASS)
  - Goal: determine whether resolver outputs and relationship edges match the source module/binding facts for exact cases selected from P1-B.
  - Scope Boundary:
    - Editable: none.
    - Inspect-only: import/export declarations, target definitions, resolver facts, and graph edges.
    - Preserve-only: downstream command traversal and semantic derivation.
    - Out of scope: resolver remediation and global TypeScript compiler conformance.
  - Non-Goals: no claim beyond measured module/ambient cases.
  - Pre-flight Questions:
    - Data source: source import/export ledger plus Anvien resolver/edge output.
    - Display permission: N/A.
    - DB read flow: read-only graph relationship inspection.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: source binding paths versus resolved edge paths.
    - Cleanup/quarantine: output only under Anvien.
    - External side effects: none.
    - N/A notes: product intent is not inferred from unresolved status.
  - Work Steps:
    1. Enumerate exact import/re-export/ambient cases and run the matching Anvien resolution, context, or graph queries after refresh.
       - UI flow check: N/A.
       - DB/data flow check: preserve edge direction, type, source site, and unresolved classification.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E1-P1C-SRC1..E1-P1C-GRAPH2`.
    2. Trace any mismatch only to the earliest observable resolver boundary and write the bounded report, leaving downstream interpretation separate.
       - UI flow check: N/A.
       - DB/data flow check: classify missing, extra, wrong target/type, duplicate, or not proven.
       - Render location check: report under Anvien.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E1-P1C-CMP1..E1-P1C-REPORT1`.
  - Implementation Gate:
    - P1-B source sites and graph symbols are stable; no resolver conclusion relies solely on a command summary.
  - Acceptance:
    - Source: module/binding truth is independently enumerated.
    - Runtime/UI: N/A.
    - DB/data: resolved and unresolved edges are compared with exact samples.
    - Behavior test: re-run yields the same bounded classification or records nondeterminism.
    - Cleanup/quarantine: no target artifacts.
    - Evidence IDs: `E1-P1C-SRC1..E1-P1C-REPORT1`.
    - Actual-status rows refreshed: `P1-C`.
  - Evidence Targets: import/export ledger, resolver output, edge comparison, report.
  - Actual-status Update: update P1-C and P2-B assumptions with first resolver divergence candidates.
  - Commit Boundary: no implementation commit; report and ledger artifacts only.

- [x] P1-D: Graph serialization/database projection, multiplicity, nullability, and cardinality. (seven canonical facts parity-checked; Supervisor bounded PASS; global projection audit remains out of scope)
  - Goal: compare raw serialized graph facts with the persisted/projection surface without treating repeated representations as separate physical source sites.
  - Scope Boundary:
    - Editable: none.
    - Inspect-only: graph JSON/meta, supported read-only database queries, property values, counts, source-site keys, and relationship endpoints.
    - Preserve-only: target source and resolver logic.
    - Out of scope: database implementation changes and undocumented engine internals.
  - Non-Goals: no speed interpretation from timings.
  - Pre-flight Questions:
    - Data source: fresh graph artifact and read-only projection tied to the same target analysis.
    - Display permission: N/A.
    - DB read flow: raw JSON versus read-only Cypher/projection comparison.
    - DB write flow: N/A; no database mutation.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: bidirectional identity/property/cardinality comparison.
    - Cleanup/quarantine: retain only evidence needed to reproduce the comparison under Anvien.
    - External side effects: none.
    - N/A notes: any unsupported projection field is marked unresolved, not guessed.
  - Work Steps:
    1. Define exact comparison keys for nodes, relationships, source sites, nulls, and multiplicity; run read-only projection commands against the proven fresh graph.
       - UI flow check: N/A.
       - DB/data flow check: distinguish fact multiplicity from physical source-site multiplicity.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E1-P1D-KEY1..E1-P1D-PROJ2`.
    2. Classify projection discrepancies and write the report with raw samples and unsupported boundaries.
       - UI flow check: N/A.
       - DB/data flow check: record NULL, stale-row, duplicate, missing-property, and count cases separately.
       - Render location check: Anvien report path only.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E1-P1D-CMP1..E1-P1D-REPORT1`.
  - Implementation Gate:
    - Graph/projection identity is tied to the same target analysis; no live or unrelated index may be substituted.
  - Acceptance:
    - Source: serialized graph schema and source-site meaning are documented.
    - Runtime/UI: N/A.
    - DB/data: supported projection parity is proven or explicitly blocked field-by-field.
    - Behavior test: repeated read-only comparison is stable.
    - Cleanup/quarantine: no target artifacts and no mutation.
    - Evidence IDs: `E1-P1D-KEY1..E1-P1D-REPORT1`.
    - Actual-status rows refreshed: `P1-D`.
  - Evidence Targets: comparison-key definition, raw graph/projection samples, count ledger, report.
  - Actual-status Update: update P1-D and P2-C assumptions.
  - Commit Boundary: no implementation commit; report and ledger artifacts only.

- [x] P1-E: Context/impact and bounded derived-process/semantic projections. (fixed command cases reconciled; bounded record accepted with global process/semantic scope unresolved)
  - Goal: verify whether agent-facing commands and derived outputs faithfully expose the bounded source/graph facts already established, without calling heuristic output a complete execution or product model.
  - Scope Boundary:
    - Editable: none.
    - Inspect-only: exact symbols/flows from P1-B through P1-D, command outputs, process traces, and semantic labels.
    - Preserve-only: unrelated command surfaces and target source.
    - Out of scope: command or semantic fixes and unsupported product-intent claims.
  - Non-Goals: no global process completeness or feature-boundary conclusion.
  - Pre-flight Questions:
    - Data source: source ledger, fresh graph, and exact Anvien command outputs.
    - Display permission: N/A.
    - DB read flow: read-only command/projection flow.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: expected bounded command output versus observed output.
    - Cleanup/quarantine: reports and command captures in Anvien only.
    - External side effects: none.
    - N/A notes: missing product authority is recorded as unresolved.
  - Work Steps:
    1. Re-run exact fixed command/process/semantic cases only after the graph refresh and compare outputs with upstream source/graph ledgers.
       - UI flow check: N/A.
       - DB/data flow check: verify traversal boundary, entity kind, and derived-input completeness.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E1-P1E-CMD1..E1-P1E-DERIVED2`.
    2. Separate upstream graph loss from downstream projection behavior, classify product-intent gaps, and write the report.
       - UI flow check: N/A.
       - DB/data flow check: preserve trace samples and caps/limits as contract facts.
       - Render location check: Anvien report path only.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E1-P1E-CMP1..E1-P1E-REPORT1`.
  - Implementation Gate:
    - Upstream source/graph ledgers for the selected cases are accepted or explicitly blocked; no semantic claim is made without an authority source.
  - Acceptance:
    - Source: command/derived expected inputs are enumerated.
    - Runtime/UI: N/A.
    - DB/data: observed output is reconciled with graph facts and bounded limits.
    - Behavior test: exact command rerun and output capture are reproducible.
    - Cleanup/quarantine: no target artifacts.
    - Evidence IDs: `E1-P1E-CMD1..E1-P1E-REPORT1`.
    - Actual-status rows refreshed: `P1-E`.
  - Evidence Targets: command captures, traversal/trace ledger, semantic authority notes, report.
  - Actual-status Update: update P1-E and P2-C assumptions.
  - Commit Boundary: no implementation commit; report and ledger artifacts only.

### P2: Anvien First-Divergence Root-Cause Slices

- Phase Goal: trace each accepted discrepancy from P1 to its earliest responsible Anvien source boundary.
- Phase Boundary:
  - In scope: Anvien source inspection, graph-supported file/symbol context, impact evidence, and bounded causal probes in Anvien.
  - Out of scope: changing the owner code, writing a fix, broad refactors, and claims beyond P1 evidence.
  - Dependencies: corresponding P1 report is accepted or explicitly blocked.
- Phase Implementation Rule: do not implement `P2` directly. Open one causal slice at a time and keep its owner/boundary fixed.
- Ordered Slice List:
  - P2-A: scanner/extractor/identity first divergence.
  - P2-B: resolver/module/ambient first divergence.
  - P2-C: projection/command/derived first divergence.

- [x] P2-A: Trace scanner, extractor, and identity first divergence. (scanner and TypeScript identity/extractor first divergences separately reported; identity report Supervisor PASS; bounded phase accepted)
  - Goal: map confirmed P1-A/P1-B discrepancies to the earliest Anvien source function and data-contract loss.
  - Scope Boundary:
    - Editable: none; source inspection only.
    - Inspect-only: Anvien scanner/parser/extractor/IR/identity code and related callers.
    - Preserve-only: resolver and downstream consumers unless needed to prove the boundary.
    - Out of scope: remediation edits.
  - Non-Goals: no assumption that all extraction defects share one cause.
  - Pre-flight Questions:
    - Data source: accepted P1 evidence plus Anvien source and graph context.
    - Display permission: N/A.
    - DB read flow: N/A; source/graph inspection only.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: bounded source probe or reproducible command, if needed.
    - Cleanup/quarantine: probes only under Anvien `.tmp`.
    - External side effects: none.
    - N/A notes: no edit means impact is recorded as inspection evidence only.
  - Work Steps:
    1. Refresh the Anvien-code graph before graph-based owner/impact queries, then inspect the exact source path and callers/callees for each P1-confirmed discrepancy.
       - UI flow check: N/A.
       - DB/data flow check: trace source AST/IR/identity fields.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E2-P2A-SRC1..E2-P2A-IMPACT2`.
    2. Reproduce the first divergence with a read-only probe and write the causal report with blast-radius warning and unresolved remainder.
       - UI flow check: N/A.
       - DB/data flow check: show input/output at the boundary.
       - Render location check: Anvien report path only.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E2-P2A-CAUSE1..E2-P2A-REPORT1`.
  - Implementation Gate:
    - P1 evidence identifies a reproducible discrepancy and the exact Anvien pipeline family; no edit is authorized.
  - Acceptance:
    - Source: first divergence and owner are line-traceable.
    - Runtime/UI: N/A.
    - DB/data: boundary input/output and graph effect are shown.
    - Behavior test: read-only repro matches P1 observation.
    - Cleanup/quarantine: no target artifacts.
    - Evidence IDs: `E2-P2A-SRC1..E2-P2A-REPORT1`.
    - Actual-status rows refreshed: `P2-A`.
  - Evidence Targets: Anvien source paths, file-detail/impact output, probe result, causal report.
  - Actual-status Update: record root-cause confidence and any affected P2-B/P2-C assumptions.
  - Commit Boundary: no implementation commit; report and probe artifacts only.

- [x] P2-B: Trace resolver, module, and ambient first divergence. (bounded report independently reproduced and Supervisor-reviewed PASS)
  - Goal: identify the earliest resolver/index boundary responsible for accepted P1-C discrepancies.
  - Scope Boundary:
    - Editable: none.
    - Inspect-only: Anvien resolver/index/module source, relevant IR and relationship emission.
    - Preserve-only: unrelated command consumers.
    - Out of scope: resolver fixes or compiler-semantic expansion.
  - Non-Goals: no global TypeScript semantics claim.
  - Pre-flight Questions:
    - Data source: P1-C evidence and Anvien source/graph context.
    - Display permission: N/A.
    - DB read flow: N/A.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: bounded resolver probe.
    - Cleanup/quarantine: Anvien `.tmp` only.
    - External side effects: none.
    - N/A notes: no target copy.
  - Work Steps:
    1. Use fresh Anvien-code graph context and source inspection to trace the exact binding/index lookup and its callers.
       - UI flow check: N/A.
       - DB/data flow check: preserve module identity and edge construction fields.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E2-P2B-SRC1..E2-P2B-IMPACT2`.
    2. Reproduce the first resolver divergence and write the causal report, explicitly separating proven mechanism from missing language authority.
       - UI flow check: N/A.
       - DB/data flow check: show exact source binding and graph result.
       - Render location check: Anvien report path only.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E2-P2B-CAUSE1..E2-P2B-REPORT1`.
  - Implementation Gate:
    - P1-C is accepted or blocked with exact evidence; no source edit is pending.
  - Acceptance:
    - Source: resolver first divergence is line-traceable.
    - Runtime/UI: N/A.
    - DB/data: causal output explains the bounded graph discrepancy.
    - Behavior test: repro is fresh and repeatable.
    - Cleanup/quarantine: no target artifacts.
    - Evidence IDs: `E2-P2B-SRC1..E2-P2B-REPORT1`.
    - Actual-status rows refreshed: `P2-B`.
  - Evidence Targets: resolver source, graph context/impact, probe, report.
  - Actual-status Update: record resolver causal confidence and downstream implications.
  - Commit Boundary: no implementation commit; report and probe artifacts only.

- [x] P2-C: Trace projection, command, and derived first divergence. (bounded report produced; Supervisor PASS; global process/semantic scope unresolved)
  - Goal: distinguish projection/consumer defects from upstream graph loss for accepted P1-D/P1-E cases.
  - Scope Boundary:
    - Editable: none.
    - Inspect-only: graph projection, database adapter, command traversal, and derived process/semantic source.
    - Preserve-only: scanner/extractor/resolver owners already covered by P2-A/P2-B.
    - Out of scope: consumer fixes and semantic redesign.
  - Non-Goals: no claim that a bounded trace is a complete product execution model.
  - Pre-flight Questions:
    - Data source: P1-D/P1-E evidence plus Anvien source and graph context.
    - Display permission: N/A.
    - DB read flow: read-only projection and command path.
    - DB write flow: N/A.
    - Render location: N/A.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: exact command/trace repro.
    - Cleanup/quarantine: Anvien `.tmp` only.
    - External side effects: none.
    - N/A notes: product-intent gaps remain unresolved without authority.
  - Work Steps:
    1. Refresh Anvien-code graph and inspect the projection/consumer source, callers, and impact for each accepted downstream discrepancy.
       - UI flow check: N/A.
       - DB/data flow check: trace nullability, multiplicity, traversal and cap handling.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E2-P2C-SRC1..E2-P2C-IMPACT2`.
    2. Reproduce the boundary behavior and write the causal report with invariant scope and residual unverified surfaces.
       - UI flow check: N/A.
       - DB/data flow check: distinguish upstream missing fact from downstream projection loss.
       - Render location check: Anvien report path only.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E2-P2C-CAUSE1..E2-P2C-REPORT1`.
  - Implementation Gate:
    - P1-D/P1-E evidence is accepted or explicitly blocked; no consumer edit is authorized.
  - Acceptance:
    - Source: first projection/consumer divergence is line-traceable.
    - Runtime/UI: N/A.
    - DB/data: invariant and blast-radius boundary are explicit.
    - Behavior test: command/trace repro is fresh.
    - Cleanup/quarantine: no target artifacts.
    - Evidence IDs: `E2-P2C-SRC1..E2-P2C-REPORT1`.
    - Actual-status rows refreshed: `P2-C`.
  - Evidence Targets: source paths, file-detail/impact, command repro, causal report.
  - Actual-status Update: record projection/consumer confidence and unresolved semantic scope.
  - Commit Boundary: no implementation commit; report and probe artifacts only.

### P3: Synthesis and Independent Acceptance

- Phase Goal: combine only independently verified slices into a causal matrix and obtain a zero-trust Supervisor verdict.
- Phase Boundary:
  - In scope: report reconciliation, evidence/benchmark ledger completeness, artifact cleanup for this plan, and Supervisor review.
  - Out of scope: remediation, optimization, target changes, and unsupported global claims.
  - Dependencies: P1 and P2 slices complete or explicitly blocked.
- Phase Implementation Rule: do not declare closure until the synthesis and Supervisor report are written and the target-boundary check passes.
- Ordered Slice List:
  - P3-A: Causal synthesis and unresolved-boundary ledger.
  - P3-B: Supervisor review of the investigation record.

- [x] P3-A: Causal synthesis and unresolved-boundary ledger. (10-row bounded causal matrix, classification ledger, R10/R11 reconciliation, and target-boundary proof complete; final P3-B review PASS)
  - Goal: produce a single evidence-bounded report separating confirmed wrong, bounded valid, unresolved, and blocked findings.
  - Scope Boundary:
    - Editable: plan/evidence/benchmark/status/report artifacts in Anvien only.
    - Inspect-only: all accepted slice reports and fresh raw evidence.
    - Preserve-only: `E:\cheapapp.org` and Anvien production source.
    - Out of scope: fixes and performance conclusions.
  - Non-Goals: no global accuracy verdict.
  - Pre-flight Questions:
    - Data source: P0-P2 evidence IDs and raw artifacts.
    - Display permission: N/A.
    - DB read flow: N/A; report synthesis only.
    - DB write flow: N/A.
    - Render location: Anvien report directory.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: evidence cross-reference and target status check.
    - Cleanup/quarantine: remove only dead artifacts created by this plan after identifying replacements.
    - External side effects: none.
    - N/A notes: no source changes.
  - Work Steps:
    1. Reconcile all slice reports against raw evidence, preserve disagreements, update ledgers and actual-status rows.
       - UI flow check: N/A.
       - DB/data flow check: retain counts, samples, and cardinality distinctions.
       - Render location check: report/evidence files in Anvien.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E3-P3A-SYN1..E3-P3A-LEDGER2`.
    2. Write the causal synthesis and verify tracked target source/worktree state remains unchanged; separately record any Anvien-generated ignored guidance-file side effect alongside the expected `.anvien` graph refresh.
       - UI flow check: N/A.
       - DB/data flow check: no new graph mutation.
       - Render location check: synthesis readable in Anvien.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E3-P3A-REPORT1..E3-P3A-TARGET1`.
  - Implementation Gate:
    - Every included claim has exact evidence IDs and a bounded scope; unresolved rows remain explicit.
  - Acceptance:
    - Source: causal matrix traces each confirmed finding to source and Anvien owner evidence.
    - Runtime/UI: N/A.
    - DB/data: graph/projection boundaries and counts are preserved.
    - Behavior test: final target read-only check passes.
    - Cleanup/quarantine: dead plan-created artifacts are identified and handled in Anvien only.
    - Evidence IDs: `E3-P3A-SYN1..E3-P3A-TARGET1`.
    - Actual-status rows refreshed: `P3-A`.
  - Evidence Targets: synthesis report, cross-reference table, target status/hash check.
  - Actual-status Update: update P3-B with the exact review scope and remaining blockers.
  - Commit Boundary: no implementation commit; report/ledger artifacts only.

- [x] P3-B: Supervisor review of the investigation record. (final bounded-record Supervisor PASS in `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md`; no global accuracy or remediation approval)
  - Goal: independently decide whether the bounded investigation record is acceptable without accepting a global accuracy or remediation claim.
  - Scope Boundary:
    - Editable: Supervisor report and review metadata in Anvien only.
    - Inspect-only: plan, actual-status, evidence, benchmark, slice reports, raw samples, source and graph identities.
    - Preserve-only: target and Anvien production source.
    - Out of scope: repairing findings during review.
  - Non-Goals: no implementation approval.
  - Pre-flight Questions:
    - Data source: current plan artifacts and fresh evidence.
    - Display permission: N/A.
    - DB read flow: use Anvien where topology/impact matters.
    - DB write flow: N/A.
    - Render location: Supervisor report under Anvien.
    - UI behavior flow: N/A.
    - Docker runtime: N/A.
    - Playwright target: N/A.
    - Behavior test: review evidence freshness and invariant closure.
    - Cleanup/quarantine: preserve valid evidence; identify dead work only.
    - External side effects: none.
    - N/A notes: review does not repair code.
  - Work Steps:
    1. Read the complete current plan/ledgers and verify source, graph, target-boundary, and report claims independently.
       - UI flow check: N/A.
       - DB/data flow check: inspect affected invariant surfaces.
       - Render location check: N/A.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E3-P3B-REVIEW1`.
    2. Write exactly one Supervisor verdict, record required follow-up for REJECT, and update plan status without claiming unsupported closure.
       - UI flow check: N/A.
       - DB/data flow check: residual unverified surfaces explicit.
       - Render location check: report readable in Anvien.
       - Mini QA for each completed implementation slice (MUST): N/A.
       - Evidence target: `E3-P3B-REPORT1`.
  - Implementation Gate:
    - P3-A report and all referenced evidence are current and readable.
  - Acceptance:
    - Source: Supervisor inspected source/graph evidence where required.
    - Runtime/UI: N/A.
    - DB/data: affected invariant closure is evaluated.
    - Behavior test: review freshness is recorded.
    - Cleanup/quarantine: no dead evidence is silently removed.
    - Evidence IDs: `E3-P3B-REVIEW1..E3-P3B-REPORT1`.
    - Actual-status rows refreshed: `P3-B`.
  - Evidence Targets: Supervisor report and final target-boundary check.
  - Actual-status Update: mark plan `accepted-bounded`, `rejected`, or `blocked` only from the Supervisor result.
  - Commit Boundary: no implementation commit; review artifact only.

- [x] Pn-A: Call supervisor for the completed investigation acceptance loop. (final bounded-record Supervisor PASS: `E3-P3B-REPORT1`)
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, reports, target boundary, and unresolved scope before closure.
  - Work Steps:
    1. Call the supervisor skill to review the complete investigation record.
    2. If Supervisor rejects, return only to the rejected evidence slice, update the causal record, and rerun review.
    3. Repeat until Supervisor passes or records a blocker.
  - Implementation Gate: P0-P3 are complete or explicitly blocked.
  - Acceptance: Supervisor PASS for the bounded investigation record, or an evidence-backed blocker is recorded without closure.
- [x] Pn-B: Remove dead work created during this plan. (cleanup audit found no dead plan-created artifact; valid historical rejection/probe artifacts were preserved)
  - Goal: ensure Anvien contains only artifacts that still support the accepted investigation record.
  - Work Steps:
    1. Identify artifacts created by this plan that were superseded or failed.
    2. Remove only those dead artifacts inside Anvien after preserving valid evidence.
    3. Ask Supervisor to review the cleanup.
  - Implementation Gate: no target artifact is selected for cleanup.
  - Acceptance: no dead plan-created artifact remains and cleanup evidence is recorded.
- [x] Pn-C: Close the plan. (closed as `accepted-bounded`; no implementation/detect-changes/commit claim)
  - Goal: finalize ledgers, target-boundary proof, Supervisor result, and plan status.
  - Work Steps:
    1. Record final evidence and benchmark state.
    2. Run `anvien detect-changes` only if implementation work was performed; otherwise record N/A with reason.
    3. Verify worktree and target state, then close as bounded accepted or blocked.
  - Implementation Gate: Pn-A and Pn-B pass or record blockers.
  - Acceptance: final status distinguishes investigation acceptance from graph correctness, remediation, and performance closure.

## Risk Notes

## Live Status R5 (2026-07-26)

- Owner correction applied: target is `E:\cheapapp.org`; graph/index remains at `E:\cheapapp.org\.anvien`; reports/evidence remain in `E:\Anvien`.
- P0 safety/path gate is reopened and completed for the corrected target after a fresh direct analyze.
- P1-A bounded finding is recorded in `reports/Investigation/20260726_144919_p1a_file_inventory.md`: eight tracked non-ignored TS/JS source files are absent from the graph, with zero graph-only TS/JS paths.
- P2-A scanner first-divergence tracing is bounded and reported; P1-C has five fixed resolution/binding mismatches; P1-B, P1-D, P1-E and P2-B/P2-C remain open. No remediation or global claim is authorized.

- The target may have user changes or a moving HEAD; P0 must capture and preserve its exact baseline without writing to it.
- The graph/index is intentionally target-local at `E:\cheapapp.org\.anvien`. Anvien analyze also timestamp-touched ignored generated `AGENTS.md`/`CLAUDE.md`; this side effect is recorded as a boundary caveat, while all investigation artifacts remain Anvien-only.
- The live Anvien worktree already contains untracked prior reports and plan artifacts. They are user-owned context; do not delete, rewrite, or treat them as fresh evidence unless this plan explicitly creates a replacement artifact.
- Missing SPEC authority can block product-intent or boundary claims while leaving structural source-to-graph checks possible.

## Live Status R6 (2026-07-26)

- P2-B bounded resolver/module tracing is reported in `reports/Investigation/20260726_155230_p2b_resolver_root_cause.md`.
- Ambient first divergence: TypeScript lib declarations are absent from the ScopeIR workspace; Promise/Math source facts are extracted, then `resolveTypeAnnotation` / member resolution emit unresolved gaps. Go-oriented graph-health classification subsequently labels them `in_repo_unresolved/analyzer_gap`.
- Barrel first divergence: `resolveImportedDef` searches only physical definitions in the consumer's target file (the barrel), does not follow the barrel's re-export binding, and leaves `TargetDef` nil; direct-import in-memory control resolves both consumer calls.
- Historical R6 snapshot (superseded by R10): P2-B was unchecked until the root agent and Supervisor independently reviewed the report/evidence. No remediation, tests, build, detect-changes, or commit occurred in that snapshot.
- Command duration and graph size are benchmark metadata only; this plan does not diagnose performance.
- A HIGH or CRITICAL blast-radius result is a warning requiring scope control, not an automatic prohibition, if a later implementation task is separately approved.

## Live Status R7 (2026-07-26)

- P2-C bounded tracing is reported in `reports/Investigation/20260726_161241_p2c_projection_command_root_cause.md`.
- `context` and `impact` read exact graph relationships for the fixed cases; they preserve existing calls/gaps and cannot recover the two upstream barrel-mediated `CALLS` edges missing from P2-B.
- The selected symbols/files have zero raw `STEP_IN_PROCESS` edges, so empty command process arrays are not a separate projection loss. Derived process completeness remains unresolved because process generation is heuristic/capped and no product authority was accepted.
- Database projection changes representation without changing selected cardinality: `relationshipCSVRow` collapses absent `Relationship.Step` to `0`, and the native Ladybug read adapter returns all selected scalars as strings.
- Historical R7 snapshot (superseded by R10): P2-C was unchecked until root/Supervisor independently reviewed the report and raw/source/control evidence. No target analyze, target write, production edit, build, detect-changes, or commit occurred in that snapshot.
- Historical R7 continuation (superseded by R13): the first P2-C Supervisor pass rejected stale derived-process wording (`E2-P2C-REJECT1`); the durable report and status ledger were corrected and subsequently re-reviewed PASS.

## Current Status R10 (2026-07-26 16:47 +07)

- P1-B's bounded three-file comparison remains accepted by Supervisor; the P2-A TypeScript identity/extractor trace is now independently reviewed PASS in `reports/Supervisor/rp_supervisor_260726_164714_by_gpt-5-6-sol_p2a_ts_identity.md` (`E2-P2A-REVIEW1`).
- P2-B is Supervisor PASS (`E2-P2B-REVIEW1`); its ambient-workspace and barrel-re-export boundaries remain bounded and do not authorize remediation.
- P2-C is Supervisor PASS after the prior evidence-integrity rejection (`E2-P2C-REVIEW2`); the corrected control remains `3,771/662/2,761/0/0` versus `3,773/662/2,761/0/0`, with no global process or semantic conclusion.
- At the R10 checkpoint the plan/evidence/benchmark/actual-status ledgers were reconciled; the later R11–R13 blocks record P3-A/P3-B completion and final bounded acceptance.
- Target HEAD and graph hash remain `a869876ab6262dacde6cd5d432d099a91852a646` and `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`; the target remains read-only for this investigation and `p0-rc-c-fixture` remains removed.
- R5–R7 above are historical snapshots. Their measured facts and the P2-C rejection history remain preserved, but their pending-review wording is superseded by this R10 current-state block.

## Final Status R13 (2026-07-26 17:00 +07)

- P3-B and the supervisor acceptance loop are complete with `E3-P3B-REPORT1` PASS for the bounded investigation record.
- Cleanup review found no dead plan-created artifact that could be removed without discarding valid historical evidence; no target artifact was selected.
- Plan status is `accepted-bounded`. This closes the evidence workflow only; it does not claim global graph accuracy, product-semantic completeness, remediation, performance, or implementation correctness.
- R10 is now a historical checkpoint superseded by this R13 final-status block; its evidence remains valid provenance.
