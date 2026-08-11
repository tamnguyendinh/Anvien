# Anvien Graph Identity Contract and Strict Construction Actual Status

Title: Anvien Graph Identity Contract and Strict Construction
Date: 2026-08-11
Status: P1-A/P1-B/P1-C/P1-D/P1-E Accepted at Isolated Commit Boundaries / P1-E Accepted by Explicit Owner Decision, Not a Fabricated Supervisor PASS / Pn-A Open / Pn-B, Pn-C, and Child 02 Closed
Companion plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
Companion evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
Companion benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`

## Purpose

This file records the current source-backed identity/range/occurrence/collision boundary for Child 01. It separates the measured target finding from the DRAFT architecture proposal. P1-B is accepted in isolated commit `6b3a3fc4480c7f4c2291d8004207f50b52d91d21`; P1-C is accepted in isolated commit `df0d7b7b753620f56c932f0e13391c1c016a8f72`; P1-D is accepted in isolated commit `10f66ffdd084a4a8710fd347836a1a083d4021bc`. P1-D's complete caller inventory keeps global mutation preserve-only. P1-E passed its full build, five-run local determinism, owned-boundary integrity, bounded target `4/4`, exact target preservation, focused/sibling/product regression, durable reporting, and exact `25/25` current-slice dead-artifact cleanup without a production/test edit. On 2026-08-11 the Owner explicitly accepted P1-E and ordered the slice transition. The terminated hidden P1-E Supervisor lane produced no durable verdict, so this acceptance is recorded truthfully as an Owner decision, not a Supervisor PASS. The isolated P1-E detect/commit boundary is completed before Pn-A opens as a new visible Child-level Supervisor task; Pn-B/Pn-C and Child 02 remain closed.

## Freshness / Refresh Rules

- The Anvien graph was refreshed with `anvien analyze --force --json` at P1-C implementation HEAD `6b3a3fc4480c7f4c2291d8004207f50b52d91d21` before the current file-detail and impact queries; `anvien status` and `.anvien/meta.json` report the same indexed/current commit and `stale=false`.
- After the production/test edits, the required QA-impact refresh on 2026-08-11 reports `1,566` scanned, `679` parsed-code, `0` failed files, `85,295` nodes, and `124,278` relationships; current test file-detail rows are `stale=false/changedSinceAnalyze=false`.
- The required full build then completed at `214.1s` with built-CLI self-analyze `1,566/679/0` and `95,539` nodes / `134,405` relationships; this inventory is validation evidence, not an occurrence-conservation denominator.
- After P1-C commit `df0d7b7b`, the final post-plan-refresh P1-D analyze reports `1,570` scanned, `679` parsed-code, `0` failed files, `95,577` nodes, and `134,444` relationships. Exact UID impact counts remained stable across the plan refresh; all final P1-D file-detail rows are current for the production files.
- After the first candidate was rejected and the exact source/QA invariant was corrected, the latest planner-boundary refresh reports `1,571` scanned, `680` parsed-code, `0` failed files, `95,698` nodes, and `134,629` relationships. Current Child ledger/roadmap file-detail rows are low risk and `stale=false/changedSinceAnalyze=false`.
- The exact repo-native full build then exited `0` in `140.4s`; built self-analyze reported `1,571/680/0`, `95,699` nodes, and `134,630` relationships. The generated old binary moved to `.tmp\p1d-held-anvien-before-full-build.exe` was deleted exactly after the old handles released.
- After runtime/QA/cleanup and coder-report completion, the pre-Supervisor refresh reports `1,572` scanned, `680` parsed-code, `0` failed, `95,713` nodes, and `134,644` relationships; the added document inventory includes the coder report and final pre-review ledgers.
- The first formal P1-D Supervisor review reran source, a `106.6s` full build, built/runtime boundaries, focused/package/11-family/`61/61` QA, and returns `REJECT` only for living-ledger truthfulness and current-date `.tmp` lifecycle. A main-agent fresh graph before reject-correction plan edits reports `1,573` scanned, `680` parsed-code, `0` failed, `95,725` nodes, and `134,656` relationships. The correction removes 11 exact FAIL/retry/superseded artifacts, retains nine final/unique valid captures, and changes no production/test file.
- After the reject correction, consistency assertions PASS `19/19`, `git diff --check` exits `0`, production/test hashes remain unchanged, the current-date inventory is exactly nine retained artifacts, and a fresh graph reports `1,573` scanned, `680` parsed-code, `0` failed, `95,726` nodes, and `134,657` relationships.
- History-closure Supervisor report `rp_supervisor_260811_122327_by_gpt-5-codex_child01_p1d_history_closure.md` returns `PASS` after fresh direct built runtime, focused/package/11-family/`61/61` QA, graph/file-detail/exact-UID CRITICAL impact, nine-artifact content validation, and `25/25` living-authority assertions. Its optional full-build retry encountered the known editor-owned output lock before compile-validation and is not used as PASS evidence.
- Post-PASS fresh analyze reports `1,574/680/0`, `95,739/134,670`; staged all-scope detect-changes exits `0` at overall/file risk `high`, covers all 12 P1-D paths, reports `168` changed symbols, `13` affected symbols, and 13 affected processes, with every changed gap classified and zero persisted gap/degraded-node/node-with-gap result.
- At clean P1-D HEAD `10f66ffdd084a4a8710fd347836a1a083d4021bc`, the P1-E pre-flight refresh again reports `1,574/680/0`, `95,739/134,670`. Fresh file-detail reports roadmap `28`, contract `0`, each Child 01 ledger `1`, and Child 02 actual-status `1` related file; all are low risk and `stale=false/changedSinceAnalyze=false`.
- The P1-E target pre-state is freshly recaptured before any target analyze: HEAD/branch `a869876ab6262dacde6cd5d432d099a91852a646`/`master`; `13` entries (`7` tracked modifications, `6` untracked, `0` staged); status hash `D39E4DF20380C0F45F8200CEECD6F862C15FAABDE86ED1B0E5EB386422899351`; oracle source hash `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C`; graph hash/size `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`/`315,569,880` bytes; meta inventory `1,359/84,807/114,125`. Per-path dirty-file hashes are retained in the P1-E validation record. No target write has occurred.
- The first exact P1-E full-build attempt stopped before compile-validation because editor-owned MCP processes hold the generated packaged binary. Doctor evidence classifies the processes as expected editor integrations and the analyze lock as free; no process was killed. `E1-P1E-BUILD1` remains pending until the same full-build command passes after a reversible repo-local generated-binary relocation.
- After reversible relocation of only the held generated binary into repo-local `.tmp`, the exact full-build command exits `0` in `114.4s`; packaged Go/Ladybug, Web production bundle, global CLI install, launcher/server, CLI `1.2.8`, and built self-analyze all complete. The self-analyze reports `1,574/680/0`, `95,739/134,670`; it is build evidence, not determinism/integrity proof. The held old binary remains explicitly tracked for cleanup after its handle releases.
- The first matched-run harness is invalidated because its own empty-stderr byte-count step threw after one successful analyze but before the canonical result row. It is not counted toward `5/5`; exact failed stdout/stderr/runtime graph artifacts are removed before retry, and no product conclusion is taken from the partial run.
- The clean matched-run retry PASSes `5/5`: identical `20/19` graphs, `10/10` unique Definition targets and `DEFINES`, zero missing endpoints, identical complete canonical hash `09DA4C6741D8D35DEA7B8680121F1B449C90E145D03B287E7917F67DB6ADA6FA`, identical graph SHA-256/size `C6C942B666931CE1AD505E3FE44B52C3887A6C16FBE0B2DAF94C25FE0AF8C1E9`/`44,199`, and stable `time`/`now` graph-visible lines. Median duration is `3.220s`; max RSS is `456,040,448` bytes. Full owned-boundary integrity remains partial until range/selection and compiled collision/fault regression pass.
- Post-build owned-boundary integrity now PASSes: P1-B provider/ScopeIR range-selection-owner matrix `3/3`; P1-C identity family `4/4`; P1-D focused matrix `6/6` plus three collision subcases; resolution/graph packages; and the independently compiled public resolution fault boundary at SHA-256 `233EAEE1C1510627AEEFCA4C852FC20AC4A5121C7098159A1D31D5767C468F2E`. No production file changed.
- The final fresh normal target analyze using packaged CLI `1.2.8` exits `0` in `41.350s` at peak process RSS `1,263,955,968` bytes, reports `1,359/887/0` files and graph `86,164/115,888`, and has empty stderr. Exact Ladybug queries prove the bounded `Variable` oracle `4/4`, four unique IDs/scopes, four `DEFINES`, and one valid File endpoint. Two consecutive target runs produce byte-identical Graph JSON at `415,998,487` bytes / SHA-256 `CAEAE4DB27B94B78204DB9EF8660DE1B8D275B637A7B16E4F349E33BC8954F09`.
- Target post-state preservation PASSes: same HEAD/branch, `0` staged, same 13-entry status and status hash, all `13/13` current-file hashes unchanged, oracle source hash unchanged, and settings hash unchanged. Only normal `.anvien` graph/Ladybug/meta output changed; exact final artifact hashes are recorded by `E1-P1E-BOUNDARY1`.
- Fresh post-report cleanup verified all six residual paths inside `E:\Anvien\.tmp`, with zero reparse points in the four directory trees. A fixed allowlist removed the three empty directories, matched-runtime tree, and compiled resolution executable; all five report absent. The final pre-build binary initially failed one exclusive-open probe, but OS-level evidence then found zero exact-file users and showed that the old MCP PID mapped a different `anvien.exe~` file object. Independent exclusive-open and exact one-file Git dry-run gates passed; the matching exact cleanup command removed the binary. The current-slice pattern inventory is empty (`25/25` removed), no process was killed, and no unrelated `.tmp` path was touched. The fresh graph refresh after cleanup and living-ledger reconciliation exits `0` at `1,577/680/0`, `95,768/134,699`.
- After the compact-safe rule/plan re-anchor, a new mandatory fresh analyze again exits `0` at `1,577/680/0`, `95,768/134,699`. The unchanged inventory validates graph freshness only and does not rerun or replace completed P1-E build/runtime/QA gates.
- The Owner then explicitly accepted P1-E and ordered the slice transition. No durable P1-E Supervisor verdict exists; the decision is an Owner acceptance/override limited to this slice. Pn-A remains the independent zero-trust Supervisor review of the complete Child 01 result.
- All file-detail rows below report `stale=false` and `changedSinceAnalyze=false`.
- Target validation is not run in P0; target source, graph, and pre-existing worktree state are preserve-only.
- Future slices must refresh only rows affected by accepted evidence or a completed slice and append a status-refresh row.

## Scope

Target scope:

- declaration/symbol identity and lexical-owner/meaning inputs;
- full range and selection-range facts;
- declaration occurrence conservation through the production ScopeIR-to-graph path;
- the proven graph node collision/replacement boundary;
- deterministic projection and endpoint integrity;
- bounded `time`/`now` target oracle at P1-E.

Out of scope:

- Graph JSON/Ladybug persistence and readers;
- TypeScript binding-pattern extraction, export semantics, module/re-export, ambient/external, and graph-health semantics;
- relationship identity/merge policy unless a later Child 01 impact proves it is required;
- scanner repair, target-source edits, and target analyze in P0.

## Relationship / Impact Evidence

P0 rows retain their recorded baseline evidence at `7042dda8bfc02ee42b40c3a1c0ede89138439481`; the P1-C candidate rows for `internal/resolution/indexes.go` and inspect-only `internal/resolution/emit.go` were refreshed at `6b3a3fc4480c7f4c2291d8004207f50b52d91d21`. Related-file counts are file counts, not symbol counts.

| Unit / File / Surface | File Detail Evidence | Related File Count | Relationship Summary | Impact Note |
|-----------------------|----------------------|--------------------:|----------------------|-------------|
| `internal/resolution/indexes.go` | `E0-P0A-FD1`, `E1-P1C-IMPACT1` | 46 | production workspace/index and graph-ID construction | high file risk; current `graphIDForDef`/`buildWorkspace` impacts CRITICAL |
| `internal/scopeir/range.go` | `E0-P0A-FD2` | 227 | shared four-coordinate range shape | medium file risk; `Range` CRITICAL |
| `internal/scopeir/definition_index.go` | `E0-P0A-FD3` | 225 | standalone definition-index utility | medium file risk; `BuildDefinitionIndex` LOW, test-only |
| `internal/graph/types.go` | `E0-P0A-FD4`, `E1-P1D-IMPACT1` | 239 | graph storage, mixed duplicate-ID mutation, index rebuild | high file risk; current `Graph.AddNode` and `Graph.init` impacts CRITICAL; inspect/preserve-only for P1-D after caller classification |
| `internal/resolution/emit.go` | `E0-P0A-FD5`, `E1-P1C-IMPACT1`, `E1-P1D-IMPACT1` | 42 | mixed File enrichment and canonical Definition projection, endpoints, separate relationship merge | high file risk; `emitNode`/`emitDefinitionNodes` CRITICAL; exact P1-D editable owner |
| `internal/scopeir/facts.go` | `E0-P0A-FD6` | 231 | `DefinitionFact`/`ScopeFact` shared contracts | medium file risk; `DefinitionFact` CRITICAL |
| `internal/scopeir/ir.go` | `E0-P0A-FD7` | 229 | owned copies and deterministic normalization | high file risk; inspect-only |
| `internal/scopeir/sort_keys.go` | `E0-P0A-FD8` | 226 | definition/range sort keys | high file risk; inspect-only |
| `internal/providers/tsjs/definitions.go` | `E0-P0A-FD9` | 16 | TS definition construct range, scope binding, visibility input | high file risk; exact functions conditional P1-B/P1-C |
| `internal/providers/tsjs/nodes.go` | `E0-P0A-FD10` | 20 | tree-sitter position conversion | high file risk; `nodeRange` MEDIUM |
| `internal/providers/tsjs/scopes.go` | `E0-P0A-FD11` | 17 | lexical scope candidates and containment | high file risk; inspect-only unless P1-B proves need |
| `internal/providers/tsjs/extract.go` | `E0-P0A-FD12` | 23 | provider orchestration into ScopeIR | high file risk; preserve orchestration |
| `internal/resolution/resolve.go` | `E0-P0A-FD13`, `E1-P1D-IMPACT1` | 51 | binding/result orchestration and fail-clear command boundary | high file risk; `ResolveBoundInto` CRITICAL; minimum P1-D error propagation only |
| `internal/resolution/types.go` | `E0-P0A-FD14` | 31 | resolution result/metrics contracts | high file risk; inspect-only |
| `internal/analyze/analyze.go` | `E0-P0A-FD15` | 182 | normal analyze phase orchestration | high file risk; preserve; no P0 edit |

Exact symbol impacts (`--direction upstream --include-tests`) are recorded in `E0-P0A-IMPACT1..E0-P0A-IMPACT2`. CRITICAL/HIGH is a blast-radius warning, not a prohibition.

## Status Rules

| Status | Meaning | Allowed next action |
|--------|---------|---------------------|
| `correct` | Current behavior satisfies the scoped requirement or a preserve-only sibling contract. | Preserve and validate. |
| `partial` | A required fact exists but is incomplete or not projected. | Change only the proven gap after the slice gate. |
| `wrong` | Current behavior contradicts a measured requirement. | Correct only at the evidence-backed owner. |
| `missing` | Required behavior or fact does not exist. | Implement only after owner evidence. |
| `unbound` | A fact exists but is not connected to its real flow or contract. | Bind only at the proven boundary. |
| `fake-or-stub` | Placeholder or fallback behavior is presented as real. | Remove or replace truthfully. |
| `blocked` | Required authority or acceptance evidence is not yet closed. | Do not implement. |

## Current Status Matrix

| Unit | Current State | Required State | Status | Relationship Count | Evidence | Next Plan Decision |
|------|---------------|----------------|--------|--------------------|----------|--------------------|
| problem/acceptance authority | bounded reports prove four ScopeIR occurrences become two graph nodes; report repair design remains DRAFT | use only measured `2/4 -> 4/4` finding and source-backed invariants | correct | N/A | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | preserve; ratify mechanics only in P1-A |
| graph-accuracy contract authority | active contract records the production occurrence denominator, current identity inputs/collision owner, endpoint traceability, and separate relationship boundary without prescribing a repair topology | source-backed Child 01 invariants are authoritative while implementation shape remains slice-evidence-driven | correct | 0 related files | `E1-P1A-SOURCE1`, `E1-P1A-CONTRACT1`, `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1` | preserve contract wording; P1-C uses only its recorded source/impact owner boundary |
| Anvien graph baseline/current candidate | P0 baseline at `7042dda8` was `1,557/676/0`, `85,114` nodes, `123,982` relationships after its ledger refresh; P1-B/P1-C/P1-D inventories remain in their evidence rows; the clean P1-D commit and P1-E local refresh are `1,574/680/0`, `95,739/134,670`; latest P1-E closure refresh is `1,577/680/0`, `95,768/134,699`; final target analyze is `1,359/887/0`, `86,164/115,888` | graph commands use a fresh index and inventory deltas are not treated as behavior proof | correct | N/A | `E0-P0A-GRAPH1`, `E1-P1B-DETECT1`, `E1-P1C-IMPACT1`, `E1-P1C-TEST1`, `E1-P1C-BUILD1`, `E1-P1D-IMPACT1`, `E1-P1D-GRAPH1`, `E1-P1D-BUILD1`, `E1-P1D-GRAPH2`, `E1-P1D-DETECT1`, `E1-P1E-PREFLIGHT1`, `E1-P1E-TARGET1`, `E1-P1E-REPORT1` | preserve the accepted P1-E evidence; Pn-A must refresh before its own graph-based review |
| declaration contract (`DefinitionFact`) | full construct range plus optional declaring-token selection range, label, owner ID, and visibility inputs exist; TSJS supplies selection range from `nameNode`, while unrelated providers omit the optional field | source-backed full/selection position plus identity inputs without inventing a topology | correct | 231 | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1` | preserve the P1-B fact shape; graph consumption remains P1-C/P1-E |
| range semantics | TSJS explicitly retains tree-sitter's one-based lines, zero-based UTF-8 byte columns, and exclusive end point; construct/selection ranges survive LF/CRLF, Unicode, tab, normalization, and JSON round-trip; shared `Range`/legacy `PositionIndex` semantics are unchanged | explicit source-backed provider coordinate/encoding/interval semantics and a provider-supported selection range | correct | 227 | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1` | preserve TSJS contract; no guessed cross-provider conversion |
| lexical owner / meaning | `buildWorkspace` supplies `scopeByDef` lexical membership to graph identity; provider `Label` remains the meaning prefix and `OwnerID` is part of the occurrence tuple; package, built-runtime, product-regression, review, and detect evidence distinguish scope and `Shared` meaning lanes | distinct lexical owners and provider-distinguished meanings remain distinguishable through the graph identity boundary | correct | 46 identity / 17 TS scope | `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1` | preserve at the isolated P1-C commit boundary; P1-D must run its own fresh pre-flight |
| graph identity | `graphIDForDef` maps the accepted repo-relative source inputs through an unambiguous deterministic tuple; focused/full package tests, the normal built CLI fixture, the 61-package product regression, history-closure review, and detect-changes pass | injective deterministic identity for distinct source occurrences and supported semantic lanes | correct | 46 | `E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-BUILD1`, `E1-P1C-ORACLE1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` | preserve the accepted P1-C boundary; do not reopen it without a failed later gate |
| production occurrence path | all 13 production providers retain occurrence IDs; `defsByFile` remains the denominator; package fixture conservation is `100%`; built runtime emits `10/10` unique definitions and `10/10` `DEFINES` endpoints; product regression, review, and detect-changes pass | every required input occurrence reaches a distinct accepted graph fact | correct | 46 / 42 | `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` | preserve the accepted occurrence denominator and identity output |
| graph collision/mutation | generic `Graph.AddNode` still has mixed update/insert semantics and remains unchanged. Canonical Definition emission tracks payloads in the current invocation, requires exact same-batch equality in both property-presence orders, and separately permits only compatible pre-existing base enrichment. Full build, positive/fault runtime, focused/package/caller-family/product QA, exact cleanup, history-closure Supervisor, all-scope detect-changes, and isolated commit PASS. | preserve source-proven generic enrichment while a non-exact canonical Definition identity conflict fails clearly before replacement | correct | 239 | `E0-P0A-SOURCE4`, `E0-P0A-FD4`, `E0-P0A-IMPACT2`, `E1-P1D-IMPACT1`, `E1-P1D-IMPACT2`, `E1-P1D-FAIL1`, `E1-P1D-SOURCE1`, `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-COLLISION1`, `E1-P1D-REVIEW1`, `E1-P1D-DETECT1`, `E1-P1D-COMMIT1` | preserve the accepted P1-D boundary while P1-E reruns the conflict/integrity matrix |
| definition projection/endpoints | Definition collision error propagates through `ResolveBoundInto` before the conflicting Definition's relationships are accepted; conflict/enrichment tests use the public resolution boundary and expected `DEFINES` cardinality. P1-E independently revalidates provider/ScopeIR construct-selection positions, stable graph-visible start/end lines, occurrence/scope inputs in GraphID, exact conflict behavior, and valid local/target `DEFINES` endpoints. Current graph Definition nodes do not project columns or selection range, and P1-E adds no persistence field semantics. | File enrichment remains explicit; exact canonical reinsertion is idempotent; non-exact Definition conflict fails clear; source-backed positions round-trip at their nearest owner; affected graph endpoints exist | correct for Child 01 owned boundaries | 43 | `E0-P0A-SOURCE8`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1D-IMPACT1`, `E1-P1D-SOURCE1`, `E1-P1D-TEST1`, `E1-P1D-COLLISION1`, `E1-P1E-INTEGRITY1`, `E1-P1E-TARGET1` | preserve; Child 02 may inventory actual persistence/reader consumers only after accepted handoff |
| standalone definition index | first-writer duplicate behavior is covered by legacy tests but has zero production callers | retain its existing test contract; do not mislabel it as runtime loss owner | correct | 225 | `E0-P0A-SOURCE3`, `E0-P0A-FD3`, `E0-P0A-IMPACT2` | preserve-only; do not touch in Child 01 |
| analyze/resolution orchestration | direct provider → ScopeIR → workspace → bound resolution → definition emission path is source-proven; no hidden alternate runtime found | use normal production path; change orchestration only with new evidence | correct | 182 / 50 | `E0-P0A-SOURCE5`, `E0-P0A-FD13`, `E0-P0A-FD15` | preserve-only; no P0 edit; inspect each slice |
| target boundary | final normal target analyze/query returns the exact `4/4` Variable oracle; HEAD/branch, staged state, 13-entry status/hash, all 13 current-file hashes, oracle source, and settings remain identical to pre-state; only `.anvien` normal output changed | preserve all pre-existing target source/worktree state while validating `4/4`; only normal target `.anvien` output may change | correct | external repository | `E0-P0A-BOUNDARY1`, `E0-P0A-BOUNDARY2`, `E1-P1E-PREFLIGHT1`, `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1` | preserve target; proceed only with Anvien regression/report/Supervisor gates |

## Status Refresh Log

| Refresh | Date | Repo Basis | Changed Scope | Status Changes | Evidence | Next Phase Update |
|---------|------|-----------|---------------|----------------|----------|-------------------|
| R0 | 2026-08-10 | correction-closed HEAD `7042dda8`; prior bounded reports | Child 01 identity/range candidates | broad retained candidates pending refresh | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | refresh all candidate owners before implementation |
| R1 | 2026-08-10 | fresh Anvien analyze + source/impact/target audits at `7042dda8` | Child 01 P0 source boundary | `BuildDefinitionIndex` causal candidate -> preserve-only; graph identity/collision -> confirmed wrong; range/IR inputs -> partial; orchestration -> preserve-only | `E0-P0A-GRAPH1`, `E0-P0A-SOURCE1..E0-P0A-SOURCE9`, `E0-P0A-FD1..E0-P0A-FD15`, `E0-P0A-IMPACT1..E0-P0A-IMPACT2`, `E0-P0A-BOUNDARY2` | P1-A remains gated by Supervisor; P1-B/P1-C/P1-D work steps narrowed to exact owners |
| R2 | 2026-08-10 | independent Supervisor report at `7042dda8` | P0-A acceptance and next-slice opening | P0 source/impact/status boundary -> accepted; P1-A documentation contract slice -> open; production implementation remains closed until P1-A and later implementation gates | `E0-P0A-REVIEW1` | P1-A may reconcile the source-backed contract; P1-B/P1-C/P1-D/P1-E remain ordered and closed |
| R3 | 2026-08-10 | P1-A contract reconciliation at `5dc9855b` with fresh contract file-detail/impact | graph-accuracy contract authority and P1-B assumptions | contract `partial -> correct`; occurrence denominator/collision/relationship boundaries ratified; implementation topology remains unprescribed | `E1-P1A-SOURCE1`, `E1-P1A-CONTRACT1` | P1-A awaits Supervisor/commit; P1-B remains closed and must use its own exact owner/impact gate |
| R4 | 2026-08-10 | P1-A Supervisor PASS and isolated documentation commit boundary | contract authority and P1-B opening | P1-A `review pending -> accepted`; production remains unchanged; P1-B becomes the only open implementation slice | `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1` | open P1-B pre-flight only; P1-C/P1-D/P1-E remain closed |
| R5 | 2026-08-10 | P1-B production diff at `8cbf6006` plus full-build, compiled provider boundary, focused matrix, and 61-package product regression | range/selection and identity-input boundary | `DefinitionFact` selection input `partial -> correct`; TSJS coordinate/interval contract `partial -> correct`; lexical/meaning inputs verified correct at ScopeIR while graph consumption remains `partial` | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1` | P1-B awaits Supervisor, detect-changes, and isolated commit; P1-C/P1-D/P1-E stay closed |
| R6 | 2026-08-10 | first P1-B Supervisor REJECT, ledger-only correction, and history-closure resubmission PASS | actual-status consistency and review gate | stale Detailed Findings reconciled with matrix/R5; P1-B `review pending -> Supervisor accepted`; production/test behavior unchanged | `E1-P1B-REVIEW1` | run fresh detect-changes and isolated P1-B commit only; P1-C/P1-D/P1-E stay closed |
| R7 | 2026-08-10 | final P1-B candidate after fresh `1,564/677/0` graph refresh and detect-changes | P1-B acceptance and Git boundary | detect-changes PASSes at overall risk `low` with `50` changed symbols / `7` changed files / `0` affected symbols / no affected process; P1-B `Supervisor accepted -> accepted at isolated commit boundary` | `E1-P1B-DETECT1`, `E1-P1B-COMMIT1` | create the isolated P1-B commit; P1-C opens only after that Git boundary and its own pre-flight refresh |
| R8 | 2026-08-10 | clean isolated P1-B commit `6b3a3fc4`; fresh P1-C analyze/file-detail/impact and source-owner audit | P1-C identity/occurrence pre-flight | P1-B Git gate confirmed; `graphIDForDef`/`buildWorkspace` remain the only editable production owner; projection and mutation remain inspect-only; absent historical `graphidentity`/declaration-occurrence modules are not production candidates | `E1-P1B-COMMIT1`, `E1-P1C-IMPACT1` | open P1-C only; production implementation may begin in `internal/resolution/indexes.go`; P1-D/P1-E and target remain closed |
| R9 | 2026-08-11 | P1-C scoped production diff, production-first QA, fresh QA-impact refresh, focused matrix, and resolution-package regression | identity, lexical/meaning consumption, occurrence conservation, endpoint QA | graph identity/occurrence `wrong -> partial` with package boundary PASS; full build/runtime/Supervisor/detect/commit remain pending | `E1-P1C-SOURCE1`, `E1-P1C-TEST1` | continue P1-C validation only; P1-D/P1-E and target remain closed |
| R10 | 2026-08-11 | repo-native full build and built-CLI self-analyze | P1-C build gate | build `pending -> PASS`; production/runtime behavior still requires the fixture boundary | `E1-P1C-BUILD1` | run built CLI fixture oracle; do not open P1-D/P1-E |
| R11 | 2026-08-11 | packaged CLI against repo-local copy of the stable P1-C fixture plus direct graph inspection | P1-C nearest normal runtime/oracle boundary | runtime oracle `pending -> PASS`: `10/10` unique definitions, `time 2/2`, `now 2/2`, separate `Shared` lanes, `0` missing endpoints | `E1-P1C-ORACLE1` | run wider regression and prepare zero-trust review; target/P1-D/P1-E remain closed |
| R12 | 2026-08-11 | first post-build 61-package product regression | P1-C sibling QA | all reported product packages pass except `internal/providers`; its representative heritage relationships exist but the test still expects legacy literal IDs | `E1-P1C-TEST1` | refresh graph and impact the exact provider parity test before QA-only correction; no production scope expansion |
| R13 | 2026-08-11 | fresh provider-parity QA impact, semantic endpoint assertion, 12-language focused matrix, and provider package | P1-C sibling QA repair | stale heritage assertion `FAIL -> PASS`; production/provider behavior unchanged | `E1-P1C-TEST1` | rerun the exact 61-package product regression |
| R14 | 2026-08-11 | exact post-fix 61-package product regression | P1-C full product regression | regression `FAIL -> 61/61 PASS`; only 11 measured non-standalone fixture corpus packages excluded | `E1-P1C-TEST1` | prepare coder report and zero-trust Supervisor review; P1-D/P1-E remain closed |
| R15 | 2026-08-11 | first zero-trust P1-C review plus ledger-only history correction; production/QA hashes unchanged | P1-C actual-status consistency | technical P1-C boundary cleared; review `REJECT` because Detailed Findings still described completed build/runtime/regression gates as pending; stale narrative reconciled without changing implementation or acceptance state | `reports/Supervisor/rp_supervisor_260811_090631_by_gpt-5-codex_child01_p1c.md` | resubmit the ledger-only correction; P1-D/P1-E/target remain closed and `E1-P1C-REVIEW1` remains pending |
| R16 | 2026-08-11 | history-closure Supervisor PASS plus exact current-slice artifact inventory | P1-C review and artifact lifecycle | review `REJECT -> PASS`; production/direct-test/provider-test hashes unchanged; `50/50` consistency assertions PASS; six pre-flight JSON files, one runtime debug tree, and two lane-input Markdown files removed from repo-local `.tmp`, while historical `.tmp/p1c0*` work was preserved | `E1-P1C-REVIEW1` | refresh graph, run detect-changes, record exact evidence, then create the isolated P1-C commit; P1-D/P1-E/target remain closed |
| R17 | 2026-08-11 | mandatory fresh `1,570/679/0` graph and staged all-scope detect-changes across the accepted 19-path boundary | P1-C detect and Git closure | detect-changes exits `0`, risk `medium`, `223` changed symbols / `19` changed files, one affected process (`Main -> CleanPath`, changed step `buildWorkspace` 4/5), zero persisted gaps/degraded nodes/nodes-with-gaps; P1-C `partial -> correct` at its isolated commit boundary | `E1-P1C-DETECT1`, `E1-P1C-COMMIT1` | create the isolated commit immediately; P1-D remains unopened until Git confirms the boundary |
| R18 | 2026-08-11 | clean P1-C commit `df0d7b7b`; fresh P1-D analyze/file-detail/exact-UID impact and complete AddNode inbound inventory | P1-D collision owner and caller policy | P1-C Git boundary confirmed; global AddNode candidate narrowed to inspect/preserve-only because 11 families contain mixed enrichment/update and insert commands; canonical Definition emission plus minimum error propagation becomes the exact editable owner | `E1-P1C-COMMIT1`, `E1-P1D-IMPACT1` | P1-D implementation opens only in `internal/resolution/emit.go` and minimum `resolve.go`; P1-E/target/later children remain closed |
| R19 | 2026-08-11 | first P1-D candidate, zero-trust source/QA rejection, exact impact, rejected-invariant correction, final planner-boundary graph refresh, and failed full-build attempt | P1-D canonical conflict implementation and validation boundary | first candidate `rejected -> corrected source/QA candidate`; same-batch non-exact property presence now conflicts in both orders; public resolution-boundary tests replace helper-only assertions; expected `DEFINES` cardinality is explicit; built CLI negative fixture wording is replaced by compiled malformed-ScopeIR boundary. Full build remains `pending -> operationally blocked` before compile because four editor-owned MCP executables hold the packaged binary | `E1-P1D-IMPACT2`, `E1-P1D-FAIL1`, `E1-P1D-BLOCK1`, `E1-P1D-GRAPH1`, `E1-P1D-SOURCE1` | do not run QA/runtime or claim source acceptance until full build passes; do not kill editor-owned processes; P1-E/target/later children remain closed |
| R20 | 2026-08-11 | exact repo-native full build after reversible relocation of the locked generated binary | P1-D full-build gate and artifact lifecycle | full build `operationally blocked -> PASS` in `140.4s`; built self-analyze `1,571/680/0`, `95,699/134,630`; no editor process killed. The held pre-build binary remains locked by four prior processes and is the only current-slice cleanup blocker | `E1-P1D-BUILD1` | run P1-D built positive CLI, compiled malformed-ScopeIR boundary, focused QA, and regressions; retry exact artifact cleanup when old handles release; P1-E/target/later children remain closed |
| R21 | 2026-08-11 | direct built CLI fixture plus compiled real resolution-package fault executable | P1-D runtime/collision boundary | runtime boundary `pending -> PASS`: positive fixture `1/1`, `0` failed, graph `20/19`, exact hash parity; fault probe PASSes range mismatch and optional canonical property removal/addition through `ResolveBoundInto` | `E1-P1D-COLLISION1` | run focused QA and source-proven regressions; retain exact runtime artifacts until their evidence is recorded and then remove them; P1-E/target/later children remain closed |
| R22 | 2026-08-11 | post-build focused six-test resolution matrix | P1-D conflict/enrichment/endpoint QA | focused matrix `pending -> PASS`: six top-level tests plus three collision subcases pass, including File enrichment and P1-C occurrence endpoints; `E1-P1D-TEST1` remains partial until package and product regressions pass | `E1-P1D-TEST1` | run full resolution/graph, 11 production caller-family, and exact product regressions; P1-E/target/later children remain closed |
| R23 | 2026-08-11 | full resolution/graph packages, exact 11 caller-family regression, and product package inventory/run | P1-D regression closure | `E1-P1D-TEST1` `partial -> recorded`: resolution/graph PASS, caller families `11/11` PASS, product packages `61/61` PASS after excluding only the measured 11 non-standalone corpus fixtures | `E1-P1D-TEST1` | clean exact current-slice artifacts, write coder report, and start zero-trust Supervisor; P1-E/target/later children remain closed |
| R24 | 2026-08-11 | exact current-slice artifact inventory after evidence capture | P1-D artifact lifecycle | positive runtime repo/home, compiled resolution executable, and held old binary `present -> removed`; no unrelated `.tmp` path touched and no process killed | `E1-P1D-BUILD1`, `E1-P1D-COLLISION1` | write coder report and start zero-trust Supervisor; P1-E/target/later children remain closed |
| R25 | 2026-08-11 | coder report against current P1-D source/diff/evidence | P1-D handoff package | invariant map, changed files, source lines, build/runtime/QA outputs, sibling checks, failed-attempt closure, and residual surface declaration recorded in `rp_coder_260811_111652_by_gpt-5-codex_child01_p1d_definition_collision.md` | `E1-P1D-SOURCE1`, `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-COLLISION1` | begin zero-trust Supervisor review; do not run detect/commit or open P1-E before verdict |
| R26 | 2026-08-11 | mandatory fresh graph after final pre-review source/test/ledger/report state | P1-D pre-Supervisor graph inventory | analyze PASSes `1,572/680/0`, `95,713/134,644`; no behavior or acceptance claim is inferred from counts | `E1-P1D-GRAPH2` | begin zero-trust Supervisor review; Supervisor must refresh again before its own graph-based evidence |
| R27 | 2026-08-11 | first formal P1-D Supervisor REJECT, main-agent source/report verification, fresh `1,573/680/0` graph, and exact current-date artifact classification | P1-D living-ledger truthfulness and artifact lifecycle only | production/build/runtime/QA remain cleared and all three production/test hashes remain unchanged; current-date top-level `p1d-*` inventory `20 -> 9` after deleting 3 FAIL outputs, 5 superseded exact-UID retries, and 3 superseded file-detail captures; nine final/unique valid impact/file-detail captures remain; stale matrix/Detailed Findings/next-action prose is reconciled while R19-R26 remain historical | `E1-P1D-REVIEW1` | resubmit only the ledger/artifact correction for unconditional zero-trust review; detect/commit and P1-E remain closed |
| R28 | 2026-08-11 | post-correction consistency/hash/artifact verification plus fresh graph | P1-D history-closure resubmission package | consistency `19/19` PASS; `git diff --check` PASS; exact current-date retained inventory `9`; production/test hashes unchanged; analyze PASSes `1,573/680/0`, `95,726/134,657` | `E1-P1D-REVIEW1` | launch a fresh zero-trust history-closure review; do not run detect/commit or open P1-E before unconditional PASS |
| R29 | 2026-08-11 | fresh history-closure Supervisor review and durable PASS report | P1-D Supervisor acceptance | source/diff, byte-identical production/test hashes, direct built runtime, focused/package/11-family/`61/61` QA, current graph/file-detail/CRITICAL exact-UID impacts, nine retained artifact contents, deleted-path absence, and living authority all independently PASS; residual same-invariant surfaces `none` | `E1-P1D-REVIEW1` | run mandatory fresh analyze and `detect-changes --scope all`; commit P1-D only after detect evidence and final ledger/worktree checks; P1-E remains closed until commit |
| R30 | 2026-08-11 | mandatory post-PASS fresh graph, exact 12-path staging, and all-scope change detection | P1-D pre-commit impact closure | analyze PASSes `1,574/680/0`, `95,739/134,670`; detect exits `0`, overall/file risk `high`, `168` changed symbols, `13` affected symbols, `12` changed/affected files, `13` affected processes; all 24 changed gap occurrences/entities are classified while persisted gaps/degraded nodes/nodes-with-gaps remain zero and semantic status is complete | `E1-P1D-DETECT1` | run final staged diff/ledger/hash/artifact checks and create the isolated P1-D commit; P1-E remains closed until Git confirms it |
| R31 | 2026-08-11 | final staged diff/ledger/hash/artifact closure after detect | P1-D isolated Git boundary | P1-D checklist and collision/endpoint rows move `partial -> correct`; `E1-P1D-COMMIT1` records the immediate isolated commit convention; production/test hashes, nine-artifact inventory, Supervisor PASS, detect evidence, and closed P1-E boundary remain intact | `E1-P1D-COMMIT1` | execute the isolated commit immediately; report the final hash from Git; do not start P1-E unless commit succeeds |
| R32 | 2026-08-11 | clean P1-D commit `10f66ffd`; fresh local graph/file-detail; fresh read-only target Git/source/graph pre-state | P1-E integration pre-flight | P1-D Git gate `pending -> confirmed`; P1-E `closed -> open`; target preservation is `pre-state recorded / post-state pending`; production, persistence/readers, later semantics, closure slices, and Child 02 remain closed | `E1-P1D-COMMIT1`, `E1-P1E-PREFLIGHT1`, partial `E1-P1E-BOUNDARY1` | run full repository build; then matched local validation; do not touch target until those gates pass |
| R33 | 2026-08-11 | first exact P1-E full-build attempt plus process/lock diagnosis | P1-E build gate | full build `pending -> operationally blocked before compile-validation`; source behavior remains unclassified by this attempt; editor-owned MCP processes retained; analyze lock free | `E1-P1E-BLOCK1` | relocate only the generated binary to repo-local `.tmp`, rerun the exact full build, and keep all validation/target operations closed until PASS |
| R34 | 2026-08-11 | exact full-build retry after verified repo-local generated-binary relocation | P1-E build gate | full build `operationally blocked -> PASS` in `114.4s`; packaged runtime, production Web, launcher/server, CLI `1.2.8`, and self-analyze complete; determinism/integrity/target remain pending | `E1-P1E-BUILD1` | run matched local analyze and owned-boundary integrity/fault validation; target remains closed until those results pass |
| R35 | 2026-08-11 | first matched-run harness attempt and immediate harness-error classification | P1-E determinism harness | one analyze completed but harness result `pending -> invalidated` because empty stderr was read as null before the canonical row; no product PASS/FAIL inferred; failed artifacts classified for exact removal | `E1-P1E-FAIL1` | remove only the invalidated run outputs/runtime graph, fix the null-safe measurement step, and run a clean five-run matrix |
| R36 | 2026-08-11 | five clean normal built-analyze runs with full canonical graph extraction and process measurement | P1-E determinism and local graph integrity | determinism `pending -> 5/5 PASS`; local graph occurrence/endpoints `pending -> partial PASS`; all canonical and whole-graph hashes/counts/ranges match; full integrity remains pending owned range/selection and collision/fault checks | `E1-P1E-DETERMINISM1`, partial `E1-P1E-INTEGRITY1` | run provider/ScopeIR range-selection matrix and compiled resolution collision/integrity fault boundary; target remains closed |
| R37 | 2026-08-11 | post-build P1-B/P1-C/P1-D focused/package matrix plus compiled resolution fault executable | P1-E owned-boundary integrity | integrity `partial -> PASS`: provider/ScopeIR ranges and selection, identity reorder/ownership/occurrence, exact conflict/enrichment/endpoints, resolution/graph packages, and compiled range/property fault cases all pass; production unchanged | `E1-P1E-INTEGRITY1` | local gates permit the bounded target analyze; first verify target pre-state still matches, then run only normal in-place output |
| R38 | 2026-08-11 | final packaged-CLI target analyze, exact label/relationship queries, two-run Graph JSON comparison, and complete post-state hash matrix | P1-E bounded oracle and target boundary | target oracle `pending -> 4/4 PASS`; target preservation `partial -> PASS`; four unique Variable IDs/scopes and four valid `DEFINES` are present; all 13 user worktree hashes and oracle source remain unchanged; no production/test edit | `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1` | record durable coder/QA reports, run the exact product/sibling regression, classify current-slice artifacts, then start zero-trust Supervisor; P1-E remains open and Child 02 closed |
| R39 | 2026-08-11 | final post-target focused, package, caller-family, and exact product regression | P1-E regression closure | TSJS `3/3`, identity/collision `9/9` plus three subcases, resolution/graph `2/2`, caller families `11/11`, and exact product packages `61/61` PASS with `0` failures | `E1-P1E-INTEGRITY1` | write durable coder/QA reports, classify/remove exact current-slice dead artifacts, then submit P1-E to zero-trust Supervisor; detect/commit and Child 02 remain closed |
| R40 | 2026-08-11 | durable coder/QA Markdown+JSON plus exact current-slice artifact inventory and allowed cleanup attempts | P1-E reporting and artifact lifecycle | durable reports `missing -> recorded`; QA JSON parse and final handoff PASS; `.tmp` inventory `25 -> 6` top-level paths after safe text/log/cache/runtime-graph cleanup; raw deletion policy blocks the six-path attempt and the old binary fails an exclusive-open probe, but exact handle ownership is not yet proven | `E1-P1E-REPORT1` | submit the complete package and explicit blocker to zero-trust Supervisor; do not accept, detect, commit, or open Child 02 while cleanup remains blocked |
| R41 | 2026-08-11 | compact-safe rule/plan re-anchor, exact six-path containment/reparse/lock verification, fixed-path cleanup, process ownership audit, and fresh graph | P1-E artifact lifecycle only | `.tmp` inventory `6 -> 1`: three empty directories, the matched runtime tree, and compiled resolution executable are removed with post-delete absence proof; only `p1e-held-anvien-before-full-build.exe` remains, and its exclusive-open probe fails while exact handle ownership remains unproven; no production/test or target source changed | `E1-P1E-REPORT1` | keep P1-E open; obtain exact-file lock evidence and complete cleanup before acceptance, detect, commit, or Child 02 opening |
| R42 | 2026-08-11 | independent exact-file Restart Manager/module/file-object audit plus main-agent exclusive-open and one-file Git dry-run verification | P1-E artifact lifecycle closure | the provisional MCP-lock inference is disproved; the last binary is a separate file object with zero mapped users, exact cleanup removes it, and current-slice `.tmp` inventory moves `1 -> 0` (`25/25` removed) without stopping any process or touching unrelated work | `E1-P1E-REPORT1` | artifact gate is complete; re-anchor zero-trust Supervisor to the current diff, then run detect/commit only after unconditional PASS |
| R43 | 2026-08-11 | mandatory post-cleanup and post-reconciliation Anvien graph refresh | P1-E pre-review graph evidence | analyze exits `0` at `1,577` scanned, `680` parsed-code, `0` failed, `95,768` nodes, and `134,699` relationships; counts are freshness/inventory evidence, not acceptance by themselves | `E1-P1E-REPORT1` | await an acceptance decision; detect/commit remain closed before acceptance |
| R44 | 2026-08-11 | explicit Owner decision after termination of the non-responsive hidden P1-E Supervisor lane | P1-E acceptance and slice transition authority | P1-E `review pending -> Owner accepted`; no durable Supervisor verdict is asserted or fabricated. The acceptance is limited to P1-E and does not satisfy the separate Pn-A Child-level Supervisor gate | `E1-P1E-REVIEW1` | run mandatory staged detect-changes and create the isolated P1-E commit; open Pn-A only after Git confirms that boundary |
| R45 | 2026-08-11 | exact nine-path staging, mandatory fresh `1,577/680/0` graph, and all-scope change detection | P1-E detect and isolated Git boundary | detect exits `0` at overall/file risk `low`: `50` changed symbols / `9` changed files, `0` affected symbols/processes, `0` gap/health degradation, and complete semantic fields across `95,768` nodes. P1-E `Owner accepted -> complete at isolated commit boundary`; the final hash is reported by Git after the self-referential ledger commit | `E1-P1E-DETECT1`, `E1-P1E-COMMIT1` | create the isolated P1-E commit immediately; after Git confirms it, open Pn-A as a new visible Supervisor task and keep Pn-B/Pn-C/Child 02 closed |

## P1-D Production AddNode Caller Classification

The fresh inbound graph contains `16` production calls in exactly `11` production files/families. Classification follows each caller's actual command and state transition; shared method name alone does not make the operations equivalent.

| Caller family | Command/state result | Same-ID contract proven by source | P1-D touch mode |
|---------------|----------------------|-----------------------------------|-----------------|
| COBOL/JCL | File metadata enrichment plus program/section/paragraph/job/step/dataset projection | File is enrichment/update; structural facts are insert-only or conflicting when non-exact | preserve; non-resolution identity defects are outside this slice |
| communities | deterministic synthetic community projection | insert-only; unchanged replay may be exact-idempotent | preserve |
| documents | File metadata enrichment plus occurrence-anchored Markdown sections | File is enrichment/update; sections are insert-only/exact replay | preserve |
| graph-health diagnostics | append/aggregate diagnostic evidence onto an existing real node | enrichment/update is required | preserve |
| ORM | create model node only after per-run/existing-node checks | insert-only; unexpected non-exact same ID conflicts | preserve |
| processes | normalized/deduplicated trace projection to synthetic Process nodes | insert-only; no incremental enrichment is source-proven | preserve |
| resolution | wrapper mixes existing File enrichment with canonical Definition insertion | File enrichment required; Definition exact reinsertion idempotent; non-exact Definition conflict must fail clear | edit Definition path only |
| routes | URL facts merge before one Route node is emitted | insert-only after producer normalization | preserve |
| semantic gaps | source-site-backed ResolutionGap projection | exact-idempotent allowed; non-exact same ID conflicts | preserve; separate semantic owner |
| structure | repeated folder/file traversal enriches existing structure nodes | enrichment/update and exact replay required | preserve |
| tools | definitions dedupe by name before one Tool node is emitted | insert-only after producer normalization | preserve |

## Phase Touch Map

Related files are not automatically editable. P0 touch mode is inspect-only or preserve-only for every production surface.

| Unit / File / Surface | Plan-Relevant Relationship File | Relationship to Target | Plan Item | Touch Mode | Evidence | Constraint |
|-----------------------|---------------------------------|------------------------|-----------|------------|----------|------------|
| graph-accuracy contract | Child 01 plan and ledgers | authority | P1-A | edit documentation | `E0-P0A-REPORT1`, `E0-P0A-VERIFY1` | exclude DRAFT implementation mechanics |
| `DefinitionFact`/`ScopeFact` | `internal/scopeir/facts.go` | shared IR contract | P1-B/P1-C | P1-B edited only optional `SelectionRange`; `ScopeFact` preserved | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1` | no redundant lexical/meaning topology; later consumption belongs to P1-C |
| range shape/semantics | `internal/scopeir/range.go`, `internal/providers/tsjs/nodes.go` | shared/provider position input | P1-B | `nodes.go` documents existing TSJS contract; `range.go` preserved | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1` | no guessed conversion or shared interval rewrite |
| TS definition constructor | `internal/providers/tsjs/definitions.go` | construct range/scope/visibility input | P1-B/P1-C | P1-B edited exact `addDefinition` path to retain `nameNode` range | `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1` | binding-pattern branch remains unchanged for Child 03 |
| TS scope builder | `internal/providers/tsjs/scopes.go` | lexical owner source | P1-B/P1-C | inspect-only | `E0-P0A-FD11`, `E0-P0A-SOURCE7` | preserve containment semantics unless P1 proves gap |
| normalization/sort | `internal/scopeir/ir.go`, `internal/scopeir/sort_keys.go` | deterministic occurrence ordering | P1-C/P1-E | inspect-only | `E0-P0A-FD7`, `E0-P0A-FD8` | no deduplication by name or ID without contract |
| graph identity/index | `internal/resolution/indexes.go` | `DefinitionFact` + existing lexical membership -> `defRef.GraphID` | P1-C | edit only `graphIDForDef`/`buildWorkspace` | `E0-P0A-FD1`, `E0-P0A-IMPACT1`, `E1-P1C-IMPACT1` | CRITICAL warning; no broad resolver rewrite, provider edit, or new identity package |
| identity behavior QA | `internal/resolution/identity_occurrence_test.go`, `internal/resolution/testdata/p1c_identity_repo/src/identity.ts`, affected legacy resolution tests | identity/relationship behavior and stale literal-ID consumers | P1-C | edit tests/fixture only after production behavior | `E1-P1C-TEST1` | semantic relationship tests resolve actual definition nodes; no compatibility fallback or target-derived fixture |
| graph mutation | `internal/graph/types.go` | mixed same-ID enrichment/update and insert boundary across 11 caller families | P1-D | inspect/preserve-only after complete caller classification | `E0-P0A-FD4`, `E0-P0A-IMPACT2`, `E1-P1D-IMPACT1` | no global policy rewrite; generic enrichment and unrelated producers remain intact |
| node projection | `internal/resolution/emit.go` | File enrichment plus canonical Definition nodes/endpoints | P1-C/P1-D/P1-E | edit only canonical Definition conflict path in P1-D | `E0-P0A-FD5`, `E0-P0A-SOURCE8`, `E1-P1D-IMPACT1` | preserve File enrichment and `emitRelationship`; range projection remains P1-E |
| standalone `DefinitionIndex` | `internal/scopeir/definition_index.go` | legacy test utility | P1-C/P1-D | preserve-only / validate-only | `E0-P0A-FD3`, `E0-P0A-IMPACT2` | do not claim it is production loss owner |
| resolution/analyze orchestration | `internal/resolution/resolve.go`, `internal/resolution/types.go`, `internal/analyze/analyze.go` | normal command fail-clear flow | all P1 slices | minimum `ResolveBoundInto` error propagation in P1-D; other orchestration preserve-only | `E0-P0A-FD13..E0-P0A-FD15`, `E0-P0A-SOURCE5`, `E1-P1D-IMPACT1` | no phase reorder, result-contract expansion, or unrelated orchestration change |
| target source/graph/worktree | `E:\cheapapp.org` | bounded oracle and external boundary | P1-E | preserve-only / validate-only | `E0-P0A-BOUNDARY1`, `E0-P0A-BOUNDARY2` | no source copy, manual graph edit, or P0 analyze |

## Detailed Findings

### Identity, lexical owner, and meaning

Current state:

The source-backed ScopeIR path retains distinct `DefinitionFact.ID` values, exact lexical scope ownership/bindings, provider `Label`, and member `OwnerID`. P1-C consumes provider occurrence ID, lexical scope, owner, label/meaning, semantic name, arity, and cleaned repository-relative file at the existing graph-ID owner. Package tests prove the accepted inputs are distinguishable and invariant to fact ordering; the normal built-runtime fixture and 61-package product regression pass at the owned boundary.

Required state:

```text
source occurrence + lexical owner + provider-supported meaning
-> deterministic injective graph identity
-> no name-only/range-only collapse
```

Evidence: `E0-P0A-SOURCE1`, `E0-P0A-SOURCE7`, `E1-P1B-SOURCE1`, `E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`.

Classification: `correct` at the provider/ScopeIR, package, built-runtime, product-regression, Supervisor, and detect/commit boundaries.

Allowed next action: preserve the accepted P1-C identity behavior and the completed P1-E matched-run/target oracle evidence; Pn-A reviews the complete Child result without altering production code. Reopen P1-C only if a Pn-A finding rejects that exact invariant after a plan refresh.

Forbidden next action: import the problem report's DRAFT two-tier topology, add only line/range to the old ID, use random/absolute-path identity, or fall back to a global name.

### Ranges and selection range

Current state:

`DefinitionFact` carries the existing construct `Range` plus an optional declaring-token `SelectionRange`, and TSJS `addDefinition` supplies both from the construct and `nameNode`. TSJS `nodeRange` explicitly records its source-backed one-based lines, zero-based UTF-8 byte columns, and exclusive end point. Shared `Range`, legacy `PositionIndex`, and TSJS scope containment remain unchanged. P1-E re-runs the provider/ScopeIR round-trip matrix and verifies stable graph-visible lines plus occurrence/scope identity on the real target. Current graph Definition nodes expose only start/end lines; no source evidence makes column/selection projection a Child 01 requirement.

Required state:

```text
provider position facts -> explicit coordinate/interval semantics
                       -> full declaration range + identifier selection range
                       -> stable identity/projection inputs
```

Evidence: `E0-P0A-SOURCE2`, `E0-P0A-FD2`, `E0-P0A-FD10`, `E1-P1B-SOURCE1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1`.

Classification: `correct` for the Child 01 provider/ScopeIR and graph-visible identity/line boundaries.

Allowed next action: preserve the accepted P1-B range/position inputs and P1-C identity consumption; hand the accepted fields to Child 02 only after P1-E/Child 01 closure so Child 02 can inventory actual persistence/read consumers.

Forbidden next action: silently convert byte columns to UTF-16, assume half-open/inclusive semantics, make range alone the identity key, or add projection/persistence semantics in this validation slice.

### Occurrence conservation and collision

Current state:

The production `defsByFile` slice preserves each incoming definition occurrence. The identity mapping gives each independently evidenced occurrence a unique GraphID at the P1-C package boundary; package QA proves `len(defsByFile[file]) == len(ScopeIR.Definitions)`, and the normal built-runtime fixture proves `10/10` unique definition IDs, `10/10` `DEFINES` endpoints, and zero missing relationship endpoints. The 61-package product regression passes. P1-D's complete caller inventory proves generic `Graph.AddNode` has mixed enrichment/update and insert semantics, while resolution Definition emission is the exact remaining non-exact collision path. `BuildDefinitionIndex` remains preserve-only.

Required state:

```text
input occurrence denominator -> accepted graph occurrences
                         100% with zero unexplained drops
```

Evidence: `E0-P0A-SOURCE3`, `E0-P0A-SOURCE4`, `E0-P0A-SOURCE5`, `E0-P0A-SOURCE8`, `E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1D-IMPACT1`.

Classification: P1-C occurrence identity and P1-D canonical Definition conflict handling are `correct` at their accepted isolated commit boundaries. Generic graph mutation is `partial but preserve-only` because source proves legitimate update callers that this slice must not rewrite. The standalone definition index remains `correct/preserve-only` relative to its tests.

Allowed next action: preserve the accepted P1-C/P1-D production/tests and the completed P1-E occurrence/collision/integrity evidence. Pn-A must not rerun a completed gate unless repo/evidence state invalidates it; any exact rejection returns to its owning earlier slice after a plan refresh.

Forbidden next action: blanket-change every `AddNode` caller, merge relationship deduplication into node identity, or use output node counts alone as conservation proof.

### Projection and endpoint integrity

Current state:

`emitDefinitionNodes` emits a definition node and `DEFINES` edge per `defsByFile` entry. The corrected P1-D implementation distinguishes pre-existing graph enrichment from canonical nodes already emitted in the current invocation, fails non-exact same-batch payloads in either order, and returns the error through `ResolveBoundInto` before accepting the conflicting Definition's relationships. P1-E independently re-runs the compiled fault/local endpoint boundary and verifies four unique target Variable IDs/scopes with four valid `DEFINES`. Graph-visible line projection is stable; occurrence and lexical scope remain traceable in GraphID. Columns and selection range are not current graph node fields, so any persistence/read consumption question remains Child 02 inventory work rather than an implicit P1-E production task.

Required state:

```text
corrected identity/range facts -> graph node properties
                             -> existing endpoints with no silent loss
```

Evidence: `E0-P0A-SOURCE8`, `E0-P0A-FD5`, `E0-P0A-IMPACT2`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1D-IMPACT1`.

Classification: `correct` for every Child 01 owned projection/endpoint and nearest range/selection boundary.

Allowed next action: keep the cleared canonical Definition branch, minimum error propagation, File enrichment, relationships, and tests unchanged; complete only P1-E detect/commit transition mechanics, then let Pn-A review the complete Child boundary.

Forbidden next action: treat `emitRelationship`'s semantic merge as declaration conservation or broaden Child 01 into persistence/readers.

### Orchestration and target boundary

Current state:

The normal production flow is source-proven and has no hidden definition-index stage. The target has a dirty pre-existing worktree. Its Git/source/graph pre-state was captured before target execution; the final in-place run and exact post-state comparison now prove `4/4` with all 13 pre-existing worktree entries and hashes preserved.

Classification: orchestration `correct/preserve-only`; target oracle and boundary `correct` for P1-E.

Allowed next action: preserve target state, record the Owner acceptance, complete P1-E detect/commit, and submit the complete Child 01 package to a new visible Pn-A zero-trust Supervisor task.

Forbidden next action: copy target source into Anvien, manually edit target graph, or treat pre-existing target files as this campaign's artifacts.

## Next Phase Status Decisions

| Plan Item | Actual Status Finding | Required Status / Next-Action Update |
|-----------|-----------------------|--------------------------------------|
| P1-A | source-backed contract reconciliation is Supervisor-accepted and committed; DRAFT architecture remains non-authoritative | complete; preserve accepted contract wording |
| P1-B | range/selection and owner/meaning input behavior is implemented, verified, Supervisor-accepted, detect-clean, and committed as `6b3a3fc4480c7f4c2291d8004207f50b52d91d21`; shared `Range` remains unchanged despite CRITICAL blast radius | complete and preserve |
| P1-C | scoped identity mapping, production-first package QA, full build, built-runtime fixture oracle, 61-package product regression, history-closure Supervisor PASS, exact artifact cleanup, and detect-changes are complete; `defsByFile`/runtime conservation is `100%` | complete at the isolated P1-C commit boundary; preserve until a later owning gate proves a regression |
| P1-D | implementation distinguishes exact same-batch canonical reinsertion from compatible pre-existing enrichment; production/build/runtime/QA/cleanup, history-closure Supervisor, all-scope detect-changes, and isolated commit `10f66ffd` PASS | complete and preserve; reopen only if a P1-E gate proves an exact P1-D invariant failure |
| P1-E | full build, matched local determinism/integrity, bounded target `4/4`, exact target preservation, focused/sibling regression, exact product `61/61`, durable reports, and `25/25` exact dead-artifact cleanup PASS with no production/test edit; Owner explicitly accepts the slice and orders transition | complete at the isolated P1-E commit boundary; preserve all accepted facts and do not relabel the Owner decision as Supervisor PASS |
| Pn-A | P1-A through P1-E are accepted at isolated boundaries; complete Child 01 acceptance has not yet received its own Supervisor verdict | open a new visible Supervisor task after the P1-E commit; review the complete Child result and record `E2-PNA-SUPERVISOR1` as PASS or REJECT |

## Implementation Gate

- [x] Target scope is listed in the Current Status Matrix.
- [x] Each target unit has a current implementation-HEAD status.
- [x] Each status has exact evidence IDs.
- [x] Each candidate target file has fresh file-detail evidence and a file count.
- [x] Phase Touch Map lists plan-relevant relationship files and touch modes.
- [x] Correct and preserve-only surfaces are explicit.
- [x] Partial/wrong surfaces have exact next actions and forbidden fallbacks.
- [x] Target/pre-existing boundary is recorded.
- [x] Next-phase assumptions, next actions, and work steps are refreshed from P0 evidence.
- [x] Status Refresh Log has R0 and fresh R1 rows.
- [x] P0 Supervisor review passes.

## Final P0 Decision

- [ ] P0 actual-status incomplete. Implementation is blocked.
- [ ] P0 complete. Next phase can proceed unchanged.
- [x] P0 complete. Next phase status, next action, or work steps must be updated before implementation.
- [ ] P0 complete. Target scope is preserve-only.
- [ ] P0 complete. Implementation is blocked by missing authority or evidence.

Decision note:

The source, fresh graph, file-detail, impact, source-flow, and target-boundary evidence closes the P0 classification. P1-A/P1-B/P1-C/P1-D are accepted and committed, with P1-C at `df0d7b7b753620f56c932f0e13391c1c016a8f72` and P1-D at `10f66ffdd084a4a8710fd347836a1a083d4021bc`. P1-E full build, five-run local determinism, owned-boundary integrity, target `4/4`, exact target preservation, focused/sibling regression, exact product `61/61`, durable reports, and `25/25` exact dead-artifact cleanup PASS without production/test edits. The Owner explicitly accepts P1-E and orders the slice transition; because the terminated hidden lane produced no durable verdict, this is recorded as Owner acceptance rather than Supervisor PASS. P1-E closes at its isolated detect/commit boundary, after which Pn-A opens as a new visible Supervisor task. Pn-B/Pn-C and Child 02 remain closed.
