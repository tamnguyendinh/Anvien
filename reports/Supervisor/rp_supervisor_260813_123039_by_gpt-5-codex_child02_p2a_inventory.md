# Supervisor Report: Child 02 P2-A exact inventory gate

Verdict: REJECT

Slice disposition: RETURN_P2A_FOR_INVENTORY_CORRECTION

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260813_123039_by_gpt-5-codex_child02_p2a_inventory.md`
- Review time: `260813 123039 +07:00` (`Asia/Bangkok`)
- Reviewer: `gpt-5-codex`
- Repo/project: `E:\Anvien`
- Scope reviewed: Child 02 P2-A inventory candidate at HEAD `f8b0717752c3d98e55556219567e21685c648207`
- Claim reviewed: before any production edit, the candidate has exactly classified the corrected-field persistence/read flow as four persistence owners (`3` future edit, `1` validate-only), eight Child 02 affected readers (`3` future edit, `5` validate-only), two P2-D boundaries with `analyze` shared with persistence, and 15 unique frozen rows including two audit/probe exclusions, with zero duplicate affected rows, zero unassigned affected rows, and zero unclassified mandatory leads; A14/A15 have the corrected completed-plan ownership; P2-B/P2-C/P2-D remain closed.
- Authority used: latest Owner/main acceptance-boundary clarification; `AGENTS.md`; the two session-rule attachments; Supervisor skill; current roadmap and all four Child 02 ledgers; accepted Child 01 identity/range contract and handoff; `docs/contracts/graph-accuracy-contract.md`; current source; completed 2026-05-19 and 2026-05-16 owner plans; current Git diff/worktree.
- Related artifact: `reports/QA/rp_qa_260813_112829_by_gpt-5-codex_child02_p2a_inventory.md`, SHA-256 `6C8E9AC4DB9B252389C1D441F51A0F02015A172965AED2DEB86D26E92DB39EFA`.

## Two-Axis Decision

### A. Review-subject verdict

`REJECT` applies only to the current **P2-A inventory deliverable and its exactness claim**. It does not reject the existence of known production gaps, and it does not mean those later implementation gaps had to be fixed in this audit-only slice. The rejected claim is that the current frozen matrix already provides exact symbol/path membership, correct touch mode and route, correct denominator arithmetic, and zero unclassified mandatory leads.

Current production behavior is not the review subject for acceptance in P2-A. Where current source proves a future gap and the matrix has inventoried and routed it correctly—for example A01, A03, A05, A06, and A07—that discovery is successful P2-A output and is classified `ACCEPT_P2A_AND_ROUTE_NEXT` at the finding level.

### B. Slice disposition

`RETURN_P2A_FOR_INVENTORY_CORRECTION` — do not close P2-A yet. Correct only the inventory artifact: A02 exact symbol membership, A12 touch mode/route and exact anchors, and the explicit assignment/exclusion of accepted `DEFINES` endpoint projection/schema. Then refreeze the dedupe/arithmetic and resubmit this bounded gate. P2-B/P2-C/P2-D remain closed until the corrected P2-A deliverable is accepted.

Finding-level dispositions:

- A01/A03 persistence gaps and A05/A06/A07 reader gaps: `ACCEPT_P2A_AND_ROUTE_NEXT`. They are correctly inventoried as future work for P2-B/P2-C and do not hold P2-A merely because production is unchanged.
- A02 exact header-symbol membership: `RETURN_P2A_FOR_INVENTORY_CORRECTION`. It violates the P2-A acceptance clauses requiring exact symbol/path evidence and zero unclassified mandatory leads.
- A12 validate-only classification: `RETURN_P2A_FOR_INVENTORY_CORRECTION`. It violates the P2-A classification/touch-mode and routing clauses and would give P2-C the wrong work map.
- `DEFINES` relationship projection/schema: `RETURN_P2A_FOR_INVENTORY_CORRECTION`. It violates the exact affected persistence-path and zero-unclassified-mandatory-lead clauses unless the corrected matrix explicitly groups or excludes these anchors under its dedupe rule.
- A14/A15: `OUT_OF_SCOPE_HANDOFF`. Their completed-plan ownership is proven, they remain outside this campaign, and they do not block P2-A.

## Executive Summary

- Acceptance boundary: P2-A is an inventory/classification/routing slice. Existing production gaps are valid P2-A output when they are inventoried exactly and routed correctly; they are not blockers merely because P2-B/P2-C/P2-D have not implemented them.
- Decision basis: the current inventory artifact itself is incomplete or incorrectly classified in three source-proven places. A02 omits mandatory Ladybug projection header symbols used by accepted Definition labels; A12 is marked `validate-only` even though its current flow drops an explicit label and reconstructs meaning from opaque NodeID; and accepted `DEFINES` endpoint persistence lacks an explicit first-projection/schema assignment or exclusion under the frozen matrix's own key.
- Consequence: the handoff `3` future-edit readers / `5` validate-only readers is false, and the claims of exact four-owner coverage, 15 fully classified keys, and zero unclassified mandatory leads are not established. No replacement denominator is invented by this review.
- Required outcome: correct and refreeze the matrix, regenerate its arithmetic, update the P2-A evidence/ledgers through orchestration/main, and submit a bounded re-review. Do not open P2-B/P2-C/P2-D meanwhile.

## Blocking Findings

### [HIGH] A02 does not name the exact Ladybug node projection symbols required by the accepted Definition set

Finding disposition: `RETURN_P2A_FOR_INVENTORY_CORRECTION`

File: `internal/lbugload/csv.go:25`

Issue: A02 is frozen as `internal/lbugload/csv.go::{symbolNodeColumns,nodeCSVRow}` and the later handoff says to edit that exact projection. That set is not exact. `symbolNodeColumns` is selected only for `Function`, `Class`, `Interface`, and `CodeElement`; `Method` uses `methodNodeColumns`, while other valid Definition tables use `defaultNodeColumns` through `nodeColumns`. `nodeCSVRow` has matching Function/Class/etc., Method, and default value branches. Accepted TS/JS source supplies the same construct/selection position contract across Definition labels, including Method and Variable evidence.

Evidence:

- `internal/lbugload/csv.go:25-33` defines three separate relevant header arrays: `symbolNodeColumns`, `methodNodeColumns`, and `defaultNodeColumns`.
- `internal/lbugload/csv.go:37-41` dispatches Function/Class/Interface/CodeElement to `symbolNodeColumns` and Method to `methodNodeColumns`.
- `internal/lbugload/csv.go:312-315` and `:354-355` implement the corresponding value branches in `nodeCSVRow`.
- `internal/lbugload/csv.go:428-434` dispatches remaining valid node tables through `defaultNodeColumns`.
- `internal/providers/tsjs/definitions.go:95-109` constructs every TS/JS `DefinitionFact` with construct `Range` and declaring-token `SelectionRange`; `internal/providers/tsjs/definition_position_inputs_test.go:97-155` records Method/Variable/Interface examples in the accepted position matrix.
- Current ledger evidence `E2-P2A-SOURCE1` and `E2-P2A-INVENTORY1` names only `{symbolNodeColumns,nodeCSVRow}`; the current actual-status handoff constrains P2-B to that incomplete exact set.

Why this blocks this gate: this is not a complaint that the columns are not implemented. It is a P2-A exact-inventory defect: a later P2-B agent following the frozen symbol set can update the Function/Class/etc. header while still omitting Method and default Definition-table headers. It violates plan lines 71 and 163 (`exact source symbol/path` / every affected path has exact evidence) and lines 158 and 166 (zero unassigned affected rows / source-anchor ownership audit), while the matrix's own dedupe key includes the first projection symbol.

Fix direction: A02 may remain one grouped semantic owner row if the grouping is made explicit and justified, but its exact symbol set must at least account for `symbolNodeColumns`, `methodNodeColumns`, `defaultNodeColumns`, `nodeCSVRow`, and the `nodeColumns` dispatch boundary. If the frozen dedupe rule instead makes them distinct keys, classify them distinctly and regenerate the arithmetic.

Re-review evidence required: corrected A02 row(s), exact Definition-label/header coverage, an updated P2-B handoff, and refreshed unique-key/dedup/unassigned arithmetic.

### [HIGH] A12 is wrongly routed as validate-only and omits the exact metadata persistence/interpreter anchors

Finding disposition: `RETURN_P2A_FOR_INVENTORY_CORRECTION`

File: `internal/embeddings/search.go:158`

Issue: A12 says the embedding/search boundary preserves and returns accepted metadata and therefore needs validation only. Current source proves the opposite classification. `NodesFromGraph` has explicit `Label`; the derived embedding update and `CodeEmbedding` schema omit it; hydration then derives the Ladybug table label by splitting opaque `NodeID` at the first colon and uses that reconstructed value in `MATCH (n:<label>)`.

Evidence:

- `internal/embeddings/text.go:19-26` declares `EmbeddableNode.Label`; `:37-54` copies `node.Label` into it.
- `internal/embeddings/pipeline.go:61-68` defines `EmbeddingUpdate` without label; `:177-192` drops label while preparing rows; `:208-219` writes no label.
- `internal/lbugschema/schema.go:444-448` defines `CodeEmbedding` without a label column.
- `internal/embeddings/search.go:158-173` groups matches by `labelFromNodeID`; `:272-277` derives the label from the opaque ID prefix; `:224-233` uses the derived label as the queried table.
- `internal/embeddings/search_test.go:14-47` deliberately encodes IDs such as `Function:alpha` and expects `MATCH (n:Function)`, confirming that the behavior is active rather than dead code.
- `docs/contracts/graph-accuracy-contract.md:63-64` requires current-source reader inventory and forbids reconstructing meaning from an opaque identifier when an explicit corrected field is available.
- Current `E2-P2A-INVENTORY1`, actual-status touch map, and P2-C handoff classify A12 as `validate-only`, leaving no later edit owner for this proven explicit-field violation.

Why this blocks this gate: the production defect alone is not the reason. The P2-A defect is the wrong touch mode and incomplete handoff: P2-C would be told to validate a flow that source proves must stop reconstructing meaning from NodeID. It violates plan step 154 (classify every row as edit/validate-only/preserve-only/out-of-scope), requirement 71 (record exact symbol/path, touch mode and evidence), and acceptance lines 163-166 (exact affected path and exact denominator/ownership). The claimed reader split is therefore already false even if A05-A12 remain eight unique reader-flow rows.

Fix direction: reclassify A12 as future edit and enumerate the exact first persistence/interpretation anchors. Route each correction to P2-B or P2-C according to the plan's real ownership. The row may remain one semantic search flow if that is consistent with the documented dedupe rule; otherwise split the persistence and reader roles explicitly. This review does not prescribe a replacement count.

Re-review evidence required: corrected A12 touch mode/route, exact `EmbeddingUpdate`/write/schema/hydration/label-interpreter anchors, refreshed P2-B/P2-C handoff, and regenerated arithmetic.

### [HIGH] Accepted `DEFINES` endpoint persistence is absent from the exact frozen owner/exclusion matrix

Finding disposition: `RETURN_P2A_FOR_INVENTORY_CORRECTION`

File: `internal/lbugload/csv.go:154`

Issue: accepted authority explicitly requires Graph JSON and Ladybug to preserve endpoints, and A01 emits `DEFINES`. The current persistence rows name only the node projection (`A02`) and node schema (`A03`). The matrix does not explicitly assign or exclude the actual Ladybug relationship first projection and schema. Calling downstream COPY/load transparent does not classify these upstream endpoint owners.

Evidence:

- `docs/contracts/graph-accuracy-contract.md:60-62` requires record-level preservation of corrected facts and endpoints with no silently discarded corrected record.
- Child 02 plan `:74-76` requires Graph JSON/Ladybug preservation of corrected facts and endpoints and no silent skip.
- `internal/resolution/emit.go:209-231` emits each Definition node and its `DEFINES` relationship.
- `internal/lbugload/csv.go:154-198` is the relationship projection loop; `relationshipCSVRow` at `:372-403` writes `SourceID` and `TargetID` at `:382-383`.
- `internal/lbugschema/schema.go:418-441` owns the relationship schema and its FROM/TO pairs.
- `internal/lbugload/queries.go:54-61` and `internal/lbugload/load.go:71-105` are downstream copy/load boundaries. They are legitimately transparent only after the projection/schema contract has been classified.
- The current mandatory matrix assigns `internal/lbugload/csv.go` only through A02's missing node columns and labels downstream COPY/load transparent. It assigns `NodeSchema` only through A03 and contains no `relationshipCSVRow` or `RelationSchema` row/exclusion. Yet `E2-P2A-MATRIX1` claims zero unclassified mandatory leads under `normalized path + first projection/interpretation symbol + semantic role`.

Why this blocks this gate: current source appears aligned for endpoints; no endpoint implementation change is demanded here. The blocker is solely inventory completeness. The matrix claims exact first-projection coverage and zero unclassified leads while omitting two accepted-field persistence anchors. That violates plan requirement 71 and acceptance line 163 (every affected persistence path has exact symbol/path evidence), plus the line 158/166 zero-unassigned ownership audit. If the corrected matrix proves these anchors are already covered by a stated semantic grouping, it may retain that grouping; absent such an assignment/exclusion, the current zero-unclassified claim is unsupported.

Fix direction: explicitly classify `relationshipCSVRow`/relationship export and `RelationSchema` as grouped validate-only endpoint persistence or as separate rows, with a source-backed grouping/dedupe explanation. Keep downstream `RelationshipCopyQuery`/load classified as transparent if that remains the source-proven conclusion.

Re-review evidence required: explicit endpoint projection/schema assignment or exclusion and refreshed frozen-row arithmetic. Four persistence rows and 15 total rows may be retained only if the corrected grouping and dedupe key prove those values; they cannot be assumed.

## Source-Level Clearance Notes

### Accepted field and persistence flow

- Accepted Child 01 authority: clear. `internal/scopeir/facts.go:3` carries construct `Range`, optional `SelectionRange`, `QualifiedName`, label/meaning and identity inputs. `internal/providers/tsjs/nodes.go:18-27` defines one-based lines, zero-based UTF-8 byte columns, and exclusive end. `internal/providers/tsjs/definitions.go:95-109` makes SelectionRange the declaring-token range. `internal/resolution/indexes.go:814` is the accepted canonical GraphID input boundary.
- A01 `emitDefinitionNodes`: correctly inventoried as future edit. `internal/resolution/emit.go:210-231` preserves identity, label, name, path, qualified name, construct lines and `DEFINES`, while omitting construct columns and SelectionRange. This production gap is properly routed to P2-B and is not itself a review blocker.
- A02 Ladybug node CSV projection: blocked as an exact symbol inventory, for the reasons above. One semantic grouped row remains possible after its full symbol membership is explicit.
- A03 `NodeSchema`: correctly identifies the node schema gap as future edit. Its row does not cover the independently missing relationship schema classification.
- A04 Graph JSON/normal analyze: correctly inventoried as validate-only. `internal/analyze/analyze.go:404-450` writes Ladybug and graph JSON; `writeGraphSnapshotJSON` at `:620-671` generically marshals the graph without field interpretation.
- `DEFINES` endpoint persistence: current source is aligned at the inspected projection/schema/copy path, but its exact owner/exclusion classification is missing from the frozen matrix.

### Eight Child 02 reader rows

- A05 `resolveNodeIds`: correct future-edit row. `anvien-web/src/hooks/useAppState.local-runtime.tsx:656-673` uses suffix plus first-match fallback to guess an opaque ID.
- A06 `handleNodeGroundingReference`: correct, distinct future-edit row. `useAppState.local-runtime.tsx:578-611` selects the first label/name match and then consumes its identity/path/range.
- A07 `CodeReferencesPanel`: correct future-edit row. `anvien-web/src/components/CodeReferencesPanel.tsx:403-475`, `:811-818`, and `:884-890` interpret one-based graph lines inconsistently with the zero-based file API.
- A08 `nodeRange`: clear as one affected validate-only first interpreter. `internal/filecontext/context.go:1279-1285` maps all four construct coordinates directly; wrappers delegate to it rather than adding reader rows.
- A09 Definition MCP context payload: clear as one affected validate-only row for its declared line payload. `internal/mcp/context.go:272-308` copies Definition identity/path/lines; the four-coordinate helper at `:529-562` is strictly `ResolutionGap`-gated and is correctly excluded from the Definition denominator.
- A10 `detectChangedSymbols`: its unique reader-row identity is clear at `internal/mcp/detect_changes.go:373-406`. The adjacent suffix-path and interval behavior is not used as an independent blocker in this report; it remains an exact P2-C validation concern within A10, not a ninth row.
- A11 `collectRenameChanges`: its unique reader-row identity is clear at `internal/mcp/rename.go:91-127`. The adjacent relationship-ID line fallback at `:177-186` is not used as an independent blocker in this report; it remains within A11 rather than creating another row.
- A12 embedding/search: row identity as one semantic search flow is supportable, but its validate-only classification and exact write/read anchor set are blocked as described above. `internal/httpapi/search.go` is transparent downstream and does not add a ninth reader.
- Reader-row uniqueness: the bounded current-source sweep supports A05-A12 as eight distinct reader flows, including intentionally separate A05/A06. It found no ninth same-invariant Child 02 reader. That does not rescue the false `3 edit + 5 validate-only` split.

### P2-D and exclusions

- P2-D boundaries: clear. A04 normal analyze plus A13 `Server.runCypherRead` are the two boundaries. `internal/mcp/tools.go:570-590` uses Ladybug first and falls back to the same repository's graph JSON only on native unavailability; it does not interpret corrected fields. Analyze is one physical A04 row with dual membership and should not be double-counted.
- A14 `buildPropertyAccessAudit`: corrected current routing is clear. `cmd/property-access-audit/main.go:11-35` is the standalone wrapper; `internal/graphaccuracy/property_access.go:89`/`:157` owns the audit. The completed 2026-05-19 property/access plan explicitly owns `internal/graphaccuracy` plus `cmd/property-access-audit` at line 150. No normal Anvien CLI/MCP/HTTP/Web/analyze caller was found. A14 is out of this campaign and is not Child 03.
- A15 `nodeCanonicalKey`/`idName`: corrected current routing is clear. `cmd/graph-accuracy-probe/main.go:17-87` is the standalone Node-vs-Go wrapper; `internal/graphaccuracy/graphaccuracy.go:350-385` owns the key helpers. The completed 2026-05-16 accuracy plan explicitly owns `internal/graphaccuracy` plus `cmd/graph-accuracy-probe` at line 136. Any call direction is probe to analyze, never analyze to probe; no normal CLI/MCP/HTTP/Web/analyze caller reaches the probe. A15 is out of this campaign and is not Child 07.
- The original QA report's A14/A15 later-Child wording is superseded by the current five-ledger correction and was not edited by this review.

## Denominator and Invariant Closure

- Source supports eight unique Child 02 reader-flow rows A05-A12 and two P2-D boundaries with analyze counted once.
- Source supports the corrected completed-plan exclusion ownership for A14/A15.
- Source disproves the claimed reader touch split: A12 is a future-edit flow, so `3 edit + 5 validate-only` is false. At minimum the existing eight-row grouping would become `4 edit + 4 validate-only`; this report does not freeze that as the final value because the corrected matrix must decide the exact A12 persistence/read grouping.
- The current A02 exact membership is incomplete, and accepted endpoint projection/schema leads have no exact assignment/exclusion. Therefore the claimed `4 + 8 + 2 + 1 = 15` unique-key proof and zero-unclassified-mandatory-lead result cannot be accepted from the current matrix.
- No duplicate is proven among the 15 currently written rows, and no ninth reader was found. However overall `0 duplicate / 0 unassigned / 0 unclassified` cannot be certified until the omitted mandatory symbols are classified and the dedupe calculation is rerun.
- No replacement persistence count, total frozen-row count, or dedupe result is asserted here. Grouping may preserve the existing counts only if the corrected exact symbols, semantic roles, and frozen key prove it.

## Exact Files and Artifacts Inspected

Authority/artifacts read in full in the required order before judgment:

1. `E:\Anvien\AGENTS.md`
2. `C:\Users\TAM NGUYEN\.codex\attachments\81b5dc31-d0ac-4dff-a249-98a6e63f736b\pasted-text.txt`
3. `C:\Users\TAM NGUYEN\.codex\attachments\266e114d-04d9-4621-903a-3ffd0c4b6e14\pasted-text.txt`
4. `E:\Anvien\.agents\skills\supervisor\SKILL.md`
5. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
6. All four Child 02 ledgers: plan, evidence, benchmark, and actual-status
7. `docs/contracts/graph-accuracy-contract.md`
8. `reports/QA/rp_qa_260813_112829_by_gpt-5-codex_child02_p2a_inventory.md`
9. Completed 2026-05-19 property/access plan
10. Completed 2026-05-16 graph-accuracy plan
11. Current complete five-document diff, source anchors/call paths, Git status, branch, HEAD, staged state, and QA hash

Relevant current source inspected directly:

- Accepted facts/positions/identity: `internal/scopeir/facts.go`, `internal/providers/tsjs/nodes.go`, `internal/providers/tsjs/definitions.go`, `internal/providers/tsjs/definition_position_inputs_test.go`, `internal/resolution/indexes.go`.
- Projection/persistence: `internal/resolution/emit.go`, `internal/lbugload/csv.go`, `internal/lbugschema/schema.go`, `internal/lbugload/queries.go`, `internal/lbugload/load.go`, `internal/analyze/analyze.go`.
- Readers: `anvien-web/src/hooks/useAppState.local-runtime.tsx`, `anvien-web/src/components/CodeReferencesPanel.tsx`, `anvien-web/src/services/backend-client.ts`, `internal/httpapi/file.go`, `internal/filecontext/context.go`, `internal/mcp/context.go`, `internal/mcp/detect_changes.go`, `internal/mcp/rename.go`, `internal/embeddings/text.go`, `internal/embeddings/pipeline.go`, `internal/embeddings/search.go`, `internal/embeddings/search_test.go`, `internal/httpapi/search.go`, `internal/mcp/tools.go`.
- Exclusions/ownership: `cmd/property-access-audit/main.go`, `internal/graphaccuracy/property_access.go`, `cmd/graph-accuracy-probe/main.go`, `internal/graphaccuracy/graphaccuracy.go`, and both completed owner plans.
- Transparent sibling paths inspected: HTTP graph transport, Web graph transport/conversion, Ladybug copy/load/query-row paths, ResolutionGap-gated context path, and embedding HTTP result mapping.

## Git and Worktree Boundary

- HEAD: `f8b0717752c3d98e55556219567e21685c648207`
- Branch: `master`
- Staged paths before this report: `0`
- Pre-report worktree: exactly five modified documentation files (roadmap plus four Child 02 ledgers) and one untracked QA report.
- Complete tracked diff inspected; `git diff --check` exited `0`.
- P2-A checkbox intentionally remains unchecked; `E2-P2A-REVIEW1`/`E2-P2A-COMMIT1` remain pending; P2-B/P2-C/P2-D remain closed.
- No production, test, fixture, runtime, target, or external-target mutation was present in the reviewed candidate.
- This report is the only Supervisor write. Nothing was staged, committed, cleaned, or repaired.

## Evidence Checked

Succeeded/current:

- Exactly one fresh `anvien analyze --force` was run before graph-based review work: `1,583` scanned, `680` parsed-code, `0` failed, `95,867` nodes, `134,798` relationships.
- Index/current commit both resolved to `f8b0717`; inspected file-detail results reported `stale=false` and `changedSinceAnalyze=false`.
- Fresh file-detail: `internal/lbugload/csv.go` = `184` symbols / `49` inbound / `49` outbound; `internal/embeddings/search.go` = `92` symbols / `24` inbound / `18` outbound.
- Fresh upstream impact: `emitDefinitionNodes` = CRITICAL, `6` impacted symbols / `4` modules / `33` processes; `SemanticSearch` = LOW, `2` impacted symbols / one module, with callers in `internal/httpapi/search.go`. CRITICAL is retained as a blast-radius warning, not an edit prohibition.
- Current full diff, Git boundary, QA hash, exact source anchors, first-interpreter paths, endpoint projection/schema, and A14/A15 completed-plan ownership were checked independently of the QA claim.
- A broad concept query was noisy and was not used as denominator evidence.

Not run:

- Full build, unit/package tests, runtime, browser, Playwright, screenshots, and UI validation: N/A for this documentation/source-inventory slice, which changes no product behavior. No behavioral success is claimed from their absence.
- `detect-changes`: not required because this Supervisor did not edit implementation or commit; the P2-A commit gate remains pending.
- P0-A QA/review: deliberately not rerun; P0-A remains accepted and outside this slice.

## Required Fix List for Resubmission

1. Correct A02's exact symbol grouping so every accepted Definition-table header/value dispatch is named and the P2-B handoff cannot miss Method/default Definition labels.
2. Reclassify A12 as future edit; name its exact metadata persistence and first-interpreter anchors; route them to P2-B/P2-C by real ownership.
3. Explicitly classify the `DEFINES` relationship projection and relation schema as grouped validate-only endpoint persistence or separate frozen rows, and retain downstream COPY/load as transparent only with that upstream classification present.
4. Refreeze the matrix and recompute persistence count, reader touch split, total unique keys, duplicate count, unassigned count, and unclassified mandatory leads from the corrected rows. Do not preserve `4` or `15` by assumption and do not add a row merely because a filename was mentioned.
5. Have orchestration/main update the owned roadmap/Child 02 ledgers and create or refresh QA evidence. Keep the current QA report immutable if required by artifact policy.
6. Submit only the corrected P2-A inventory for a fresh bounded Supervisor review. Do not implement P2-B/P2-C/P2-D during this correction.

## Overall Evaluation

The current candidate correctly identifies the eight reader-flow identities, the two P2-D boundaries, the no-code Git boundary, and the completed-plan routing of A14/A15. It also correctly routes several real production gaps forward. Acceptance fails only because the inventory artifact does not yet meet its own exactness invariant: it omits mandatory first-projection symbols, assigns A12 the wrong touch mode, and leaves accepted endpoint persistence without an explicit classification. Those defects would hand later slices an incomplete or incorrect work map. The next owner is orchestration/main for documentation/evidence correction and bounded resubmission; implementation slices remain closed.
