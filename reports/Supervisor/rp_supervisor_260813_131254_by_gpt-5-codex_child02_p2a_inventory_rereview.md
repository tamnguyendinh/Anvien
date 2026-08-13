# Supervisor Report: Child 02 P2-A corrected inventory bounded re-review

Verdict: PASS

Slice disposition: ACCEPT_CORRECTED_P2A_INVENTORY_AND_HAND_BACK

## Metadata

- Report file: reports/Supervisor/rp_supervisor_260813_131254_by_gpt-5-codex_child02_p2a_inventory_rereview.md
- Review time: 260813 131254 +07:00 (Asia/Bangkok)
- Reviewer: gpt-5-codex
- Repo/project: E:\Anvien
- Plan: docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md
- Slice reviewed: Child 02 P2-A only, bounded inventory correction after E2-P2A-REJECT1
- Production HEAD: f8b0717752c3d98e55556219567e21685c648207
- Branch: master
- Claim reviewed: the corrected inventory closes only the prior A02, A12, and DEFINES classification findings and refreezes an exact, internally consistent 19-row owner matrix.
- Authority used: latest Owner delegation; E:\Anvien\AGENTS.md; both required session-rule attachments; Supervisor skill; roadmap and all four Child 02 ledgers; graph-accuracy contract; current production source; first QA report; prior Supervisor report; current five-document diff and Git boundary.
- Related artifacts:
  - reports/QA/rp_qa_260813_112829_by_gpt-5-codex_child02_p2a_inventory.md, verified SHA-256 6C8E9AC4DB9B252389C1D441F51A0F02015A172965AED2DEB86D26E92DB39EFA
  - reports/Supervisor/rp_supervisor_260813_123039_by_gpt-5-codex_child02_p2a_inventory.md, verified SHA-256 F02B0B894C6BFEECD379E7BF265721C5753FB5F74259FCE600B393418E026636
- Open blockers for this bounded inventory acceptance: none.
- Remaining process gate: the isolated documentation commit E2-P2A-COMMIT1 is still pending and must remain orchestration-owned before P2-B opens.
- Next owner: orchestration/main.

## Executive Summary

The corrected Child 02 P2-A inventory passes the bounded re-review. Each of the three prior blocking inventory defects is closed against current source at the unchanged production HEAD:

1. C02 now names the complete Definition node-CSV header, dispatch, and value-branch set. Its one-row grouping is valid because the mutually exclusive table branches are selected and written by one indivisible node-to-CSV projection contract.
2. The former mixed A12 row is correctly separated into embedding-label persistence C05/C06 for P2-B and semantic-search interpretation C16 for P2-C. Current source carries an explicit label into EmbeddableNode, drops it before CodeEmbedding persistence, and actively reconstructs it from opaque NodeID during hydration.
3. DEFINES endpoints are explicitly assigned to aligned, validate-only relationship projection C07 and relation schema C08. Downstream COPY/fallback load forwards the declared endpoint IDs and remains transparent.
4. Independent parsing of C01-C19 confirms the stated arithmetic and unique-key result.

This PASS accepts only the corrected P2-A documentation/source inventory. It does not accept or claim corrected production behavior, field parity, reader runtime behavior, repeated analyze behavior, or any P2-B/P2-C/P2-D implementation.

## Finding-Level Disposition

### A. CLOSE_PRIOR_FINDING_A02 — PASS

Question: Does C02 name every exact Definition node CSV header/value/dispatch branch, and is one-row semantic grouping valid?

Decision: yes.

Evidence:

- internal/lbugload/csv.go:25-33 defines symbolNodeColumns, methodNodeColumns, and defaultNodeColumns.
- internal/lbugload/csv.go:34-48 maps Function/Class/Interface/CodeElement to symbolNodeColumns and Method to methodNodeColumns.
- internal/lbugload/csv.go:108-126 selects the table, calls nodeColumns, creates the writer with that header, then writes nodeCSVRow for the same table.
- internal/lbugload/csv.go:300-355 contains the matching Function/Class/Interface/CodeElement, Method, and default value branches.
- internal/lbugload/csv.go:428-435 routes every remaining valid node table to defaultNodeColumns.
- internal/lbugload/csv.go:597-601 derives valid-node membership from lbugschema.NodeTables.
- internal/lbugschema/schema.go:18-51 enumerates the valid Ladybug node tables, including the accepted Definition labels that use the default branch.
- E2-P2A-INVENTORY2 at the evidence ledger line 126 lists ExportGraphCSVs, symbolNodeColumns, methodNodeColumns, defaultNodeColumns, nodeColumnLookup, nodeColumns, and nodeCSVRow.

Grouping decision:

- The normalized owning path is internal/lbugload/csv.go.
- The semantic role is one Definition node-to-CSV projection.
- Header selection and row construction are paired once per node by ExportGraphCSVs.
- Function/Class/Interface/CodeElement, Method, and default branches are mutually exclusive implementations of that same projection contract.
- C07 remains separate because relationship endpoint projection is a different semantic role even though ExportGraphCSVs participates in both.

No exact branch required by the prior finding remains absent. One grouped C02 row is therefore valid under the declared grouping key.

### B. CLOSE_PRIOR_FINDING_A12 — PASS

Question: Are persistence C05/C06 and reader C16 split, routed, and anchored correctly?

Decision: yes.

Persistence evidence:

- internal/embeddings/text.go:19-35 declares EmbeddableNode.Label.
- internal/embeddings/text.go:37-64 copies graph node.Label into EmbeddableNode.
- internal/embeddings/pipeline.go:61-68 declares EmbeddingUpdate without a label field.
- internal/embeddings/pipeline.go:177-192 prepares EmbeddingUpdate rows without propagating Label.
- internal/embeddings/pipeline.go:208-219 creates CodeEmbedding rows without a label property.
- internal/lbugschema/schema.go:12 names the CodeEmbedding table.
- internal/lbugschema/schema.go:444-448 defines EmbeddingSchema without a label column.

Reader evidence:

- internal/embeddings/search.go:46-60 shows ChunkSearchRow and BestChunkMatch retain NodeID but no label.
- internal/embeddings/search.go:62-94 sends vector matches into hydrateSearchResults.
- internal/embeddings/search.go:132-173 carries NodeID through vectorSearchQuery, chunkRows, and DedupBestChunks, then groups hydration by labelFromNodeID.
- internal/embeddings/search.go:212-233 returns emb.nodeId from CodeEmbedding and uses the reconstructed label as the metadata table in metadataQuery.
- internal/embeddings/search.go:236-247 copies vector rows without an explicit label.
- internal/embeddings/search.go:272-277 derives the label from the opaque NodeID prefix.
- internal/embeddings/search_test.go:14-47 proves the path is active by expecting Function and Method metadata-table queries from prefixed IDs.

Routing decision:

- C05 is the first persistence loss/write role and correctly routes to P2-B.
- C06 is the storage-schema role and correctly routes to P2-B.
- C16 is the first semantic reader/interpreter role and correctly routes to P2-C.
- vectorSearchQuery, chunkRows, DedupBestChunks, ChunkSearchRow, and BestChunkMatch are correctly retained as adjacent label-propagation anchors inside C16, not promoted into independent reader rows.
- The rows are linked because hydration needs the explicit label state, but they are not duplicate roles: persistence owns preserving the field; the reader owns consuming it instead of parsing opaque identity.

The corrected split satisfies the contract rule that readers use explicit source-contract fields and do not reconstruct meaning from opaque identifiers when the corrected field is available.

### C. CLOSE_PRIOR_FINDING_DEFINES — PASS

Question: Do C07/C08 explicitly classify the mandatory DEFINES endpoint persistence leads, with downstream load remaining transparent?

Decision: yes.

Evidence:

- internal/resolution/emit.go:209-239 emits one Definition node and a DEFINES relationship with File SourceID and Definition TargetID.
- internal/lbugload/csv.go:20 declares the relationship endpoint columns from and to.
- internal/lbugload/csv.go:154-198 writes every relationship row to the aggregate CSV and to the endpoint-label pair CSV.
- internal/lbugload/csv.go:372-403 writes rel.SourceID and rel.TargetID unchanged at the first relationship CSV projection.
- internal/lbugschema/schema.go:85-113 declares File-to-Definition-table relation pairs, including Function, Class, Interface, Method, CodeElement, and the remaining valid Definition tables.
- internal/lbugschema/schema.go:418-441 builds RelationSchema from RelationPairs and the relationship property columns.
- internal/lbugload/queries.go:54-61 issues relationship COPY using the already classified from/to table pair.
- internal/lbugload/queries.go:68-102 uses the same FromID/ToID in the fallback insert.
- internal/lbugload/queries.go:143-169 reads the CSV endpoint columns into FromID/ToID without reinterpretation.
- internal/lbugload/load.go:71-107 chooses COPY or the same-ID fallback path and does not invent endpoint meaning.

Classification decision:

- C07 is the first endpoint CSV projection and is correctly validate-only because current source preserves both accepted IDs.
- C08 is the relation-pair/schema owner and is correctly validate-only because the required File-to-Definition endpoint pairs are declared.
- C07 and C08 remain separate rows because projection and schema are different semantic roles in different normalized source paths.
- RelationshipCopyQuery, CSV reading, and fallback load are transparent downstream mechanics. They were checked but do not add another owner row.

The previously unclassified mandatory endpoint leads are now explicitly owned and routed to P2-B endpoint parity.

### D. ACCEPT_REFROZEN_MATRIX — PASS

Question: Are the grouping key, counts, and zero-duplicate/unassigned/unclassified claims correct?

Decision: yes for the bounded C01-C19 universe.

Independent matrix parse:

- Total parsed rows: 19.
- Unique keys: 19.
- Duplicate keys: 0.
- Persistence C01-C08: 8 rows.
  - Future edit: C01/C02/C03/C05/C06 = 5.
  - Validate-only: C04/C07/C08 = 3.
- Readers C09-C16: 8 rows.
  - Future edit: C09/C10/C11/C16 = 4.
  - Validate-only: C12/C13/C14/C15 = 4.
- P2-D-only: C17 = 1 row.
- Out-of-campaign: C18/C19 = 2 rows.
- Total unique rows: 8 + 8 + 1 + 2 = 19.

P2-D boundary count:

- C04 is the normal analyze/Graph JSON persistence boundary and is counted once in persistence.
- C17 is Server.runCypherRead and is the one P2-D-only row.
- internal/analyze/analyze.go:404-451 and :620-672 prove the normal DB-load/snapshot boundary.
- internal/mcp/tools.go:570-590 proves the Ladybug-primary/Graph-JSON-unavailable-fallback read boundary.
- Therefore P2-D has two boundaries without double-counting C04 as a twentieth row.

Grouping-key decision:

- The declared key is normalized source path + semantic role + grouped first projection/interpreter owner.
- C02 validly groups mutually exclusive branches of one projection contract.
- C05/C06/C16 are distinct because write, schema, and interpretation are distinct semantic roles.
- C07 is distinct from C02 even inside csv.go because relationship endpoint projection is not node-field projection.
- C04 has dual slice membership but one physical persistence boundary and retains the already accepted single-row treatment.

Prior mandatory leads are all disposed:

- A02 exact branches -> C02.
- A12 persistence loss -> C05/C06.
- A12 opaque-ID reader -> C16.
- DEFINES projection/schema -> C07/C08.
- Already-cleared readers, P2-D boundaries, and completed-plan exclusions remain represented once.

Accordingly, duplicate affected rows = 0, unassigned affected rows = 0, and unclassified mandatory leads = 0 for this bounded correction.

## Source-Level Clearance Notes

- Definition node CSV projection: clear at internal/lbugload/csv.go:25-48, :108-126, :300-355, and :428-435.
- Embedding-label persistence: clear as a future P2-B correction map at internal/embeddings/text.go:19-64, internal/embeddings/pipeline.go:61-68 and :177-219, and internal/lbugschema/schema.go:444-448.
- Semantic-search reader: clear as a future P2-C correction map at internal/embeddings/search.go:46-60, :62-94, :132-173, :212-247, and :272-277.
- DEFINES endpoint projection/schema: clear as validate-only at internal/resolution/emit.go:209-239, internal/lbugload/csv.go:154-198 and :372-403, and internal/lbugschema/schema.go:85-113 and :418-441.
- Downstream relationship COPY/load: clear as transparent at internal/lbugload/queries.go:54-102 and :143-169 and internal/lbugload/load.go:71-107.
- Refrozen matrix and routes: clear at the evidence ledger lines 121-155; consistent with the plan line 161, benchmark lines 34-38, and actual-status lines 83-84, 100, 108-117, and 215-218.

## Diff and Git Boundary

Independent pre-report Git check:

- HEAD: f8b0717752c3d98e55556219567e21685c648207.
- Branch: master.
- Staged paths: 0.
- Modified tracked paths: exactly 5.
  - campaign roadmap;
  - Child 02 plan;
  - Child 02 evidence ledger;
  - Child 02 benchmark ledger;
  - Child 02 actual-status ledger.
- Pre-existing untracked paths: exactly the first QA report and prior Supervisor report.
- Production, test, fixture, runtime, target, and external-target diffs: 0.
- Complete five-document diff inspected.
- git diff --check: exit 0.

The five-document diff:

- records the rejected first attempt as immutable history;
- records the corrected C01-C19 candidate without marking E2-P2A-REVIEW1 or E2-P2A-COMMIT1 complete;
- leaves the P2-A checklist item unchecked;
- keeps P2-B/P2-C/P2-D closed;
- does not claim build, runtime, browser, or production acceptance.

This report is the sole write authorized by this Supervisor session. It is not staged or committed.

## Evidence Checked

Passed/current:

- All required authority documents were read in the required order and in full.
- Both supplied report hashes exactly matched their expected SHA-256 values.
- Current HEAD and source are unchanged from the reviewed production boundary; the Git diff contains documentation only.
- Current source anchors for A-D were inspected independently rather than accepted from ledger prose.
- All five modified document diffs were inspected independently.
- The C01-C19 table was parsed directly and counted independently.
- git diff --check exited 0.

Not rerun:

- Anvien analyze, file-detail, and impact: intentionally not rerun. The prior rejection did not invalidate those completed gates, production HEAD/source did not change, and the Owner expressly bounded this review to inventory correction.
- Full build, tests, runtime, browser, Playwright, and screenshots: not applicable to this documentation/source-inventory correction; no product behavior changed and no runtime PASS is claimed.
- detect-changes: not run because no implementation commit is being made in this Supervisor session.
- Commit/stage: not performed; E2-P2A-COMMIT1 remains with orchestration/main.
- P0-A and the broad P2-A discovery sweep: not reviewed or repeated.

## Invariant Closure

- Affected invariant: exact corrected-field persistence/reader inventory ownership and routing before production slices open.
- Prior blocking surfaces checked: A02/C02 Definition CSV projection; A12/C05/C06/C16 embedding label flow; DEFINES C07/C08 endpoint projection/schema; refrozen key/count arithmetic.
- Old failure modes:
  - Method/default Definition branches omitted from the handoff: closed.
  - Semantic-search flow mislabeled validate-only: closed.
  - Persistence and reader ownership mixed: closed.
  - DEFINES projection/schema absent from the owner matrix: closed.
  - Counts retained without a corrected key calculation: closed.
- Residual unverified same-invariant surfaces for questions A-D: none.
- Later implementation gaps intentionally remain open in their routed slices. They are not residual P2-A inventory blockers.

## Overall Evaluation

The correction is acceptable because it closes exactly the three inventory defects returned by the prior Supervisor, preserves all already-cleared findings, and produces an internally consistent work map that a later P2-B or P2-C owner can follow without missing a Definition CSV branch, conflating label persistence with label interpretation, or overlooking DEFINES endpoint persistence.

Final slice verdict: PASS for the corrected Child 02 P2-A inventory only.

Handoff: orchestration/main should record this report as the bounded E2-P2A-REVIEW1 evidence, update checklist/status only under its authority, perform the isolated documentation commit, and keep P2-B closed until that commit boundary is complete.
