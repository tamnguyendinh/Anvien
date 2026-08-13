# Anvien Current Graph Persistence and Reader Consistency Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`

## Evidence Rules

- The original problem report is the problem origin; use its bounded findings and acceptance targets, not its DRAFT architecture proposals.
- The causal synthesis and bounded Supervisor PASS verify selected persistence/representation observations but do not define the implementation.
- Child 01 handoff and current source/runtime are the authority for corrected fields and their actual consumers.
- A reader enters the affected denominator only through exact source evidence recorded by P2-A.
- A pending evidence ID is a declared target, not proof.
- Every slice requires every exact ID named in its Acceptance; broad phase references are insufficient.
- Measured counts/timings belong in benchmark; source facts, commands, and verdicts belong here.

### Evidence ID Naming

Use `E<phase>-<item>-<kind><n>`. Do not reuse an ID for another fact.

## E0 - P0 Evidence

Matching plan item: `P0-A`

- `E0-P0A-REPORT1` — recorded: `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` is the campaign problem origin, but it does not establish a Child 02 persistence requirement. Child 02 takes its bounded representation evidence from `E0-P0A-VERIFY1` and its corrected-field contract from the accepted Child 01 handoff; the report's proposed storage design is not authority.
- `E0-P0A-VERIFY1` — recorded: `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md` observes one-to-one selected-fact cardinality with bounded representation changes across raw/Cypher/file-detail paths; `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md` passes only that bounded record.
- `E0-P0A-HANDOFF1` — recorded by orchestration/main in the Pn-C handoff candidate: Child 01 accepted identity is the deterministic length-prefixed tuple of normalized repository-relative file, semantic/qualified name, optional arity, provider occurrence ID, lexical scope ID, and owner ID, with provider label/meaning retained; accepted position facts are the construct `Range` plus optional TSJS declaring-token `SelectionRange`, with one-based lines, zero-based UTF-8 byte columns, and exclusive ends. The accepted occurrence denominator is production `defsByFile`; local built evidence proves `10/10` definitions/endpoints and the bounded target proves `4/4` Variable IDs/`DEFINES` with zero missing endpoints. Canonical Definition non-exact same-batch conflicts fail clearly before relationships are accepted; generic `Graph.AddNode` enrichment and relationship merge remain preserve-only. Source-backed evidence/commits: `E1-P1B-SOURCE1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1`, `E1-P1B-COMMIT1`, `E1-P1C-SOURCE1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1C-COMMIT1`, `E1-P1D-SOURCE1`, `E1-P1D-TEST1`, `E1-P1D-COLLISION1`, `E1-P1D-COMMIT1`, `E1-P1E-INTEGRITY1`, `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1`, `E1-P1E-COMMIT1`, `E2-PNA-SUPERVISOR1`, `E2-PNA-COMMIT1`, `E2-PNB-COMMIT1`; durable handoff report: `reports/Investigation/rp_main_260811_210243_child01_pnc_closure_handoff.md`. Explicit non-claims: Child 01 does not prove Graph JSON/Ladybug field parity, affected-reader membership, repeated-analyze parity at Child 02 readers, or projection of columns/SelectionRange; P2-A must source-audit those boundaries before any edit. Handoff is accepted and P0-A has now recorded the current Child 02 graph/source/file-detail/impact baseline; exact persistence/reader completeness remains P2-A-owned.
- `E0-P0A-QA1` — PASS at the bounded P0-A boundary: `reports/QA/rp_qa_260813_090733_by_gpt-5-codex_child02_p0a_inventory.md` captures the current-codebase source/impact candidates, accepted predecessor dependency, no-fix Git boundary, and the pending reader-discovery boundary. Its proposed denominator is candidate input for P2-A, not a P0-A completion requirement or accepted denominator.
- `E0-P0A-GRAPH1` — recorded by `E0-P0A-QA1`: fresh analyze PASS at the current HEAD, `1,581` scanned / `680` parsed-code / `0` failed, graph `95,830` nodes / `134,761` relationships, and file-detail `stale=false`.
- `E0-P0A-SOURCE1` — recorded by `E0-P0A-QA1`: Graph JSON and Ladybug field flow; construct columns and optional `SelectionRange` stop at `emitDefinitionNodes`, and Ladybug symbol projection/schema also omit `qualifiedName`, construct columns, and `SelectionRange`.
- `E0-P0A-SOURCE2` — recorded by `E0-P0A-QA1`: HTTP and Web graph paths are transparent preserve-only candidates; `filecontext.nodeRange`, `resolveNodeIds`, and `runCypherRead` are current-source leads for P2-A classification. P0-A does not establish the final reader denominator, editable subset, or later-slice ownership for any lead.
- `E0-P0A-FD1` — recorded current: `internal/lbugload/csv.go`, `19` related files.
- `E0-P0A-FD2` — recorded current: `internal/analyze/analyze.go`, `182` related files.
- `E0-P0A-FD3` — recorded current: `internal/httpapi/graph.go`, `22` related files.
- `E0-P0A-FD4` — recorded current: `anvien-web/src/services/backend-client.ts`, `24` related files.
- `E0-P0A-FD5` — recorded current: `anvien-web/src/hooks/useAppState.local-runtime.tsx`, `30` related files; exact owner `resolveNodeIds`.
- `E0-P0A-FD6` — recorded current: `internal/resolution/emit.go::emitDefinitionNodes`, `42` related files.
- `E0-P0A-FD7` — recorded current: `internal/lbugschema/schema.go::NodeSchema`, `18` related files.
- `E0-P0A-FD8` — recorded current: `internal/mcp/tools.go::runCypherRead`, `55` related files.
- `E0-P0A-FD9` — recorded current: `internal/filecontext/context.go::nodeRange`, `44` related files.
- `E0-P0A-IMPACT1` — recorded current: `internal/lbugload/csv.go`, `CRITICAL`, `43` impacted symbols / `16` files / `1` flow.
- `E0-P0A-IMPACT2` — recorded current: `internal/analyze/analyze.go`, `CRITICAL`, `51` / `14` / `1`.
- `E0-P0A-IMPACT3` — recorded current: `internal/httpapi/graph.go`, `MEDIUM`, `8` / `1` / `1`.
- `E0-P0A-IMPACT4` — recorded current: `backend-client.ts`, file `CRITICAL`, `82` / `29` / `1`.
- `E0-P0A-IMPACT5` — recorded current: `useAppState.local-runtime.tsx`, file `CRITICAL`, `42` / `17` / `1`; exact `resolveNodeIds` symbol `LOW`, `0 / 0 / 0`.
- `E0-P0A-IMPACT6` — recorded current: `emitDefinitionNodes`, `CRITICAL`, `6` impacted symbols / `4` modules / `33` processes.
- `E0-P0A-IMPACT7` — recorded current: `NodeSchema`, `LOW`, `3` impacted symbols / `2` modules / `0` processes.
- `E0-P0A-IMPACT8` — recorded current: `runCypherRead`, `LOW`, `0 / 0 / 0`.
- `E0-P0A-IMPACT9` — recorded current: `nodeRange`, `CRITICAL`, `8` impacted symbols / `1` module / `16` processes.
- `E0-P0A-BOUNDARY1` — recorded: target source/worktree is preserve-only; validation uses normal target-repository output and stores plan/evidence artifacts in Anvien.
- `E0-P0A-STATUS1` — PASS: current codebase state, accepted predecessor dependency, candidate owner counts/impacts, preserve-only target boundary, and the pending P2-A discovery work are explicit and non-contradictory. Exact affected-reader completeness is explicitly absent from P0-A and assigned to P2-A.
- `E0-P0A-REVIEW1` — PASS at the scope-corrected P0-A boundary: completed zero-trust review independently cleared the accepted predecessor, current graph/source/file-detail/impact basis, three persistence candidate leads, transparent transport boundaries, no-production/test/runtime diff, and long-timeout impact completion. Owner rejected the review's attempt to make the P2-A direct-consumer denominator a P0-A blocker and accepted the completed QA/Supervisor gates for P0-A. No review gate is rerun; the out-of-slice findings are preserved as `E2-P2A-INPUT1`.

## E2 - P2 Evidence

Matching plan items: `P2-A`, `P2-B`, `P2-C`, `P2-D`, `P2-E`

| Slice | Exact evidence ID | Required proof | Status |
|-------|-------------------|----------------|--------|
| P2-A | `E2-P2A-INPUT1` | unverified candidate leads transferred from P0-A/review without treating them as a denominator or conclusion: `internal/graphaccuracy/property_access.go`, `internal/graphaccuracy/graphaccuracy.go`, Web node grounding in `useAppState.local-runtime.tsx`, direct filecontext range consumers, and the MCP `ResolutionGap`-only range consumer; P2-A must verify or exclude each row from current source | recorded input; not acceptance evidence |
| P2-A | `E2-P2A-IMPACT1` | fresh file-detail/impact for every candidate persistence/reader owner | recorded and retained; not invalidated by inventory rejection |
| P2-A | `E2-P2A-SOURCE1` | first corrected-field source trace | recorded history; exactness superseded by `E2-P2A-SOURCE2` |
| P2-A | `E2-P2A-INVENTORY1` | first symbol/path, field, touch-mode, and route inventory | rejected/superseded by `E2-P2A-REJECT1` |
| P2-A | `E2-P2A-MATRIX1` | first affected-denominator and exclusion matrix | rejected/superseded by `E2-P2A-REJECT1` |
| P2-A | `E2-P2A-REJECT1` | bounded zero-trust verdict on the first inventory | recorded `REJECT`; correction limited to A02, A12, and `DEFINES` endpoint classification |
| P2-A | `E2-P2A-SOURCE2` | corrected source trace for the three rejected invariants | recorded; accepted by `E2-P2A-REVIEW1` |
| P2-A | `E2-P2A-INVENTORY2` | exact corrected 19-row inventory and later-slice routes | recorded; accepted by `E2-P2A-REVIEW1` |
| P2-A | `E2-P2A-MATRIX2` | refrozen grouping, arithmetic, and zero duplicate/unassigned/unclassified proof | recorded; accepted by `E2-P2A-REVIEW1` |
| P2-A | `E2-P2A-REVIEW1` | unconditional Supervisor PASS for inventory | recorded: bounded PASS report/hash below |
| P2-A | `E2-P2A-COMMIT1` | isolated documentation/source-audit commit | recorded: exact boundary committed immediately after ledger update; final hash reported by Git |
| P2-B | `E2-P2B-IMPACT1` | fresh impact for exact affected Graph JSON/Ladybug owners | pending |
| P2-B | `E2-P2B-SOURCE1` | production persistence correction, if source proves needed | pending |
| P2-B | `E2-P2B-BUILD1` | full build | pending |
| P2-B | `E2-P2B-TEST1` | corrected-field round-trip/drop/closure tests | pending |
| P2-B | `E2-P2B-PARITY1` | Graph JSON/Ladybug corrected-field parity and zero dropped records | pending |
| P2-B | `E2-P2B-REVIEW1` | unconditional Supervisor PASS | pending |
| P2-B | `E2-P2B-DETECT1` | Anvien detect-changes | pending |
| P2-B | `E2-P2B-COMMIT1` | isolated slice commit | pending |
| P2-C | `E2-P2C-IMPACT1` | exact impact for P2-A affected readers | pending |
| P2-C | `E2-P2C-SOURCE1` | affected-reader production corrections only | pending |
| P2-C | `E2-P2C-BUILD1` | full build | pending |
| P2-C | `E2-P2C-TEST1` | exact affected-reader behavior and sibling regression tests | pending |
| P2-C | `E2-P2C-RUNTIME1` | nearest real boundary per affected reader, including UI evidence only when affected | pending |
| P2-C | `E2-P2C-REVIEW1` | unconditional Supervisor PASS | pending |
| P2-C | `E2-P2C-DETECT1` | Anvien detect-changes | pending |
| P2-C | `E2-P2C-COMMIT1` | isolated slice commit | pending |
| P2-D | `E2-P2D-IMPACT1` | current repeated-analyze/read owner impact before any edit | pending |
| P2-D | `E2-P2D-SOURCE1` | current behavior trace and any evidence-required production correction | pending |
| P2-D | `E2-P2D-BUILD1` | full build | pending |
| P2-D | `E2-P2D-TEST1` | unchanged/changed input plus failure/subsequent-success matrix | pending |
| P2-D | `E2-P2D-REPEAT1` | normal built repeated-analyze results and matching affected reads | pending |
| P2-D | `E2-P2D-REVIEW1` | unconditional Supervisor PASS | pending |
| P2-D | `E2-P2D-DETECT1` | Anvien detect-changes or documented no-production-change result | pending |
| P2-D | `E2-P2D-COMMIT1` | isolated implementation/validation commit | pending |
| P2-E | `E2-P2E-BUILD1` | final full build before acceptance | pending |
| P2-E | `E2-P2E-PARITY1` | independent corrected-field Graph JSON/Ladybug parity | pending |
| P2-E | `E2-P2E-READERS1` | affected readers passed/total equals exact P2-A denominator | pending |
| P2-E | `E2-P2E-REPEAT1` | independent repeated normal analyze and clear-failure acceptance | pending |
| P2-E | `E2-P2E-REVIEW1` | unconditional Supervisor PASS | pending |
| P2-E | `E2-P2E-DETECT1` | final Child 02 detect-changes result | pending |
| P2-E | `E2-P2E-COMMIT1` | isolated validation/ledger commit | pending |

### P2-A inventory evidence detail

- `E2-P2A-IMPACT1` — retained from `reports/QA/rp_qa_260813_112829_by_gpt-5-codex_child02_p2a_inventory.md` (SHA-256 `6C8E9AC4DB9B252389C1D441F51A0F02015A172965AED2DEB86D26E92DB39EFA`) at HEAD `f8b0717752c3d98e55556219567e21685c648207`: the pre-sweep fresh analyze passed with `1,582` scanned / `680` parsed-code / `0` failed and graph `95,848` nodes / `134,779` relationships; the frozen batch completed `22/22` file-detail targets with `stale=false` and `23/23` impact targets with zero command errors in `546.8s`. Numeric tuples lost to tool-output elision remain explicitly unclaimed. Retained blast-radius evidence includes `emitDefinitionNodes` CRITICAL (`6` symbols / `4` modules / `33` processes), `nodeRange` CRITICAL (`8` / `1` / `16`), `NodeSchema` LOW (`3` / `2` / `0`), `resolveNodeIds` LOW (`0/0/0`) within a CRITICAL host file, and `runCypherRead` LOW (`0/0/0`). The bounded review did not invalidate this graph/impact gate, so it is not rerun.
- `E2-P2A-SOURCE1`, `E2-P2A-INVENTORY1`, `E2-P2A-MATRIX1` — retained as the immutable first attempt. They recorded the useful frozen sweep and correctly identified multiple production gaps, but their exactness claim is superseded below.
- `E2-P2A-REJECT1` — recorded from `reports/Supervisor/rp_supervisor_260813_123039_by_gpt-5-codex_child02_p2a_inventory.md` (SHA-256 `F02B0B894C6BFEECD379E7BF265721C5753FB5F74259FCE600B393418E026636`) at the same HEAD. Verdict `REJECT` applies only to the P2-A inventory deliverable: A02 omitted Method/default Definition CSV header/dispatch symbols; A12 was wrongly `validate-only` although its reader reconstructs label from opaque NodeID after embedding persistence drops the explicit label; and `DEFINES` relationship CSV/schema ownership was not explicitly assigned or excluded. Findings already inventoried and routed to P2-B/P2-C remain valid later-slice work; A14/A15 remain out of campaign. Disposition: `RETURN_P2A_FOR_INVENTORY_CORRECTION`.
- `E2-P2A-SOURCE2` — corrected source trace, limited to the rejected invariants:
  - `internal/lbugload/csv.go` selects `symbolNodeColumns` for Function/Class/Interface/CodeElement, `methodNodeColumns` for Method, and `defaultNodeColumns` through `nodeColumns` for every other valid table; `nodeCSVRow` has the matching explicit Method and default value branches. The grouped Definition node-CSV owner therefore includes `symbolNodeColumns`, `methodNodeColumns`, `defaultNodeColumns`, `nodeColumnLookup`, `nodeColumns`, `nodeCSVRow`, and the `ExportGraphCSVs` dispatch that pairs the selected header with the row.
  - `internal/embeddings/text.go::{EmbeddableNode,NodesFromGraph}` carries the explicit node label. `internal/embeddings/pipeline.go::{EmbeddingUpdate,prepareBatch,CreateEmbeddingQuery}` drops it, and `internal/lbugschema/schema.go::EmbeddingSchema` has no label column. `internal/embeddings/search.go::{hydrateSearchResults,labelFromNodeID,metadataQuery}` reconstructs the table label by splitting opaque NodeID and uses that derived value in the metadata query. The persistence loss and reader interpretation are distinct, truly linked roles because semantic-search hydration needs the dropped label state.
  - `internal/lbugload/csv.go::{ExportGraphCSVs,relationshipColumns,relationshipCSVRow}` projects every relationship's `SourceID`/`TargetID`, including `DEFINES`; `internal/lbugschema/schema.go::{RelationPairs,RelationSchema}` declares the endpoint pairs. These anchors are aligned and validate-only, while downstream relationship COPY/load remains transparent.
  - Current source and completed-plan entrypoints continue to prove `buildPropertyAccessAudit` and `nodeCanonicalKey`/`idName` are standalone completed-plan tools outside this campaign.

#### `E2-P2A-INVENTORY2` — corrected frozen owner matrix

Grouping key: normalized source path + semantic role + grouped first projection/interpreter owner. Mutually exclusive dispatch branches remain one row only when they implement one indivisible semantic projection contract and every exact branch symbol is listed. Persistence and reader roles are separate rows even when they form one real linked pipeline.

| Key | Role | Exact owner / symbols | Current source fact | Touch mode | Route |
|-----|------|-----------------------|---------------------|------------|-------|
| C01 | persistence: Definition graph projection | `internal/resolution/emit.go::emitDefinitionNodes` | emits accepted ID/label/name/path/qualifiedName/construct lines and `DEFINES`, but omits construct columns and optional `SelectionRange` | future edit | P2-B |
| C02 | persistence: grouped Definition node CSV projection | `internal/lbugload/csv.go::{ExportGraphCSVs,symbolNodeColumns,methodNodeColumns,defaultNodeColumns,nodeColumnLookup,nodeColumns,nodeCSVRow}` | explicit Function/Class/Interface/CodeElement, Method, and default valid-table branches omit qualifiedName, construct columns, and optional `SelectionRange` | future edit | P2-B |
| C03 | persistence: Definition node schema | `internal/lbugschema/schema.go::NodeSchema` | matching explicit Method and default table branches omit qualifiedName, construct columns, and optional `SelectionRange` | future edit | P2-B |
| C04 | persistence: Graph JSON / normal analyze | `internal/analyze/analyze.go::{writeGraphSnapshotJSON,Run}` normal persistence boundary | generic writer preserves properties already present; also one of the two P2-D repeat/failure boundaries | validate-only | P2-B parity + P2-D |
| C05 | persistence: embedding label write | `internal/embeddings/text.go::{EmbeddableNode,NodesFromGraph}` source anchor; `internal/embeddings/pipeline.go::{EmbeddingUpdate,prepareBatch,CreateEmbeddingQuery}` first loss/write owners | explicit Label reaches `EmbeddableNode` but is dropped before the CodeEmbedding row | future edit | P2-B |
| C06 | persistence: embedding schema | `internal/lbugschema/schema.go::EmbeddingSchema` / `CodeEmbedding` | schema lacks the explicit node label needed by semantic-search hydration | future edit | P2-B |
| C07 | persistence: `DEFINES` relationship CSV endpoints | `internal/lbugload/csv.go::{ExportGraphCSVs,relationshipColumns,relationshipCSVRow}` | relationship export writes source and target IDs and does not drop accepted `DEFINES` endpoints | validate-only | P2-B endpoint parity |
| C08 | persistence: relationship schema endpoints | `internal/lbugschema/schema.go::{RelationPairs,RelationSchema}` | relation schema declares FROM/TO pairs used by accepted `DEFINES` endpoint persistence | validate-only | P2-B endpoint parity |
| C09 | reader: Web opaque-ID resolution | `anvien-web/src/hooks/useAppState.local-runtime.tsx::resolveNodeIds` | suffix/first-match fallback guesses an opaque ID | future edit | P2-C |
| C10 | reader: Web grounding resolution | `anvien-web/src/hooks/useAppState.local-runtime.tsx::handleNodeGroundingReference` | first label/name match loses distinct same-name identity | future edit | P2-C |
| C11 | reader: code-reference range presentation | `anvien-web/src/components/CodeReferencesPanel.tsx::CodeReferencesPanel` | interprets accepted one-based graph lines inconsistently as zero-based for display/highlight/slicing | future edit | P2-C |
| C12 | reader: file-context range | `internal/filecontext/context.go::nodeRange` | directly consumes all four construct coordinates and is semantically aligned | validate-only | P2-C |
| C13 | reader: Definition MCP context | `internal/mcp/context.go::{contextCandidatePayloads,contextSymbolPayload}` | consumes Definition identity/path/line payload without opaque-ID reconstruction; ResolutionGap-only four-coordinate helper stays excluded | validate-only | P2-C |
| C14 | reader: changed-symbol membership | `internal/mcp/detect_changes.go::detectChangedSymbols` | directly consumes corrected path/range membership | validate-only | P2-C |
| C15 | reader: rename location | `internal/mcp/rename.go::collectRenameChanges` | directly consumes corrected identity/path/range | validate-only | P2-C |
| C16 | reader: semantic-search hydration | `internal/embeddings/search.go::{SemanticSearch,hydrateSearchResults,labelFromNodeID,metadataQuery}`; adjacent label propagation path `vectorSearchQuery`, `chunkRows`, `DedupBestChunks`, `ChunkSearchRow`, `BestChunkMatch` | active hydration derives table label from opaque NodeID instead of explicit persisted label | future edit | P2-C |
| C17 | P2-D-only query/failure boundary | `internal/mcp/tools.go::Server.runCypherRead` | Ladybug-primary read with Graph JSON fallback on native unavailability; does not interpret corrected fields | validate-only | P2-D |
| C18 | out-of-campaign standalone audit | `internal/graphaccuracy/property_access.go::buildPropertyAccessAudit` via `cmd/property-access-audit` | owned by completed 2026-05-19 property/access plan; no normal analyze/CLI/MCP/HTTP/Web wiring | preserve/out of campaign | completed-plan owner |
| C19 | out-of-campaign standalone probe | `internal/graphaccuracy/graphaccuracy.go::{nodeCanonicalKey,idName}` via `cmd/graph-accuracy-probe` | owned by completed 2026-05-16 Node-vs-Go accuracy plan; no normal analyze/CLI/MCP/HTTP/Web wiring | preserve/out of campaign | completed-plan owner |

#### `E2-P2A-MATRIX2` — refrozen arithmetic and exclusions

- Persistence rows: C01-C08 = `8` unique rows: `5` future edit (C01/C02/C03/C05/C06) and `3` validate-only (C04/C07/C08).
- Child 02 reader rows: C09-C16 = `8` unique rows: `4` future edit (C09/C10/C11/C16) and `4` validate-only (C12/C13/C14/C15).
- P2-D boundaries: `2` — C04 normal analyze (already counted once in persistence) and C17 `Server.runCypherRead` (P2-D-only).
- Out-of-campaign exclusions: C18-C19 = `2` explicit rows.
- Total unique classified rows: `8 persistence + 8 readers + 1 P2-D-only + 2 exclusions = 19`.
- Duplicate affected rows: `0`; unassigned affected rows: `0`; unclassified mandatory leads: `0`.
- C02 is one row because its exact header/dispatch/value branches are mutually exclusive implementations of the same Definition node-CSV projection contract; C07 is separate because relationship endpoint projection is a different semantic role. C05 and C16 are separate because one persists required label state and the other interprets it; they are linked only because C16 needs C05's label to decide the metadata table.
- Transparent HTTP/Web graph transport, Graphology copying, Ladybug query-row copying, downstream relationship COPY/load, wrapper-only filecontext paths, and the ResolutionGap-gated MCP range helper remain explicit preserve-only exclusions rather than extra rows.
- P2-A mutation boundary remains documentation/report-only: no production, test, fixture, runtime, external target, build, browser, Playwright, or screenshot claim is introduced. `E2-P2A-REVIEW1` is accepted; `E2-P2A-COMMIT1` records the immediate isolated boundary below. P2-C/P2-D remain closed, and P2-B opens only after that commit succeeds.

- `E2-P2A-REVIEW1` — recorded: visible Supervisor task `019ff9b6-370d-75c3-95f3-cab233017c02` produced `reports/Supervisor/rp_supervisor_260813_131254_by_gpt-5-codex_child02_p2a_inventory_rereview.md` at SHA-256 `988C137D5C864192C38180A47812BAB53FE315204DA4907F6FFCC15EEE659369`; verdict `PASS`, disposition `ACCEPT_CORRECTED_P2A_INVENTORY_AND_HAND_BACK`. The review independently verified the complete C02 Definition node-CSV branch grouping, the C05/C06 persistence-to-C16 reader linkage and routes, C07/C08 `DEFINES` endpoint classification with transparent downstream COPY/load, and the `19`-row/`0/0/0` arithmetic. It inspected the full five-document diff and exact source anchors, confirmed HEAD `f8b0717752c3d98e55556219567e21685c648207`, branch `master`, staged paths `0`, and zero production/test/runtime diff, and deliberately did not rerun non-invalidated graph/impact/QA gates. Main read the durable report in full, independently rechecked its hash, Git boundary, `git diff --check`, and exact A-D disposition before accepting the handoff.
- `E2-P2A-COMMIT1` — recorded: the accepted P2-A boundary is exactly the roadmap, four Child 02 ledgers, immutable first QA report, prior REJECT report, and bounded PASS report. It is committed immediately after this row with message `docs(plan): accept child 02 persistence inventory`; Git reports the final hash because the commit cannot contain its own hash. No implementation detect-changes is required for this documentation/report-only commit under the repository doc-commit rule.

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E3-PNA-SUPERVISOR1` | independent acceptance of complete Child 02 source/diff/persistence/read/runtime evidence | pending |
| `E3-PNB-CLEANUP1` | dead-work inventory, removal, and Supervisor cleanup PASS | pending |
| `E3-PNC-DETECT1` | final detect-changes result | pending |
| `E3-PNC-HANDOFF1` | exact accepted persistence fields, affected-reader guarantees, evidence, and Child 03 refresh | pending |
| `E3-PNC-COMMIT1` | final closure commit and known worktree state | pending |
