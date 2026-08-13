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
| P2-A | `E2-P2A-IMPACT1` | fresh file-detail/impact for every candidate persistence/reader owner | pending |
| P2-A | `E2-P2A-SOURCE1` | corrected-field source traces through real persistence/read flows | pending |
| P2-A | `E2-P2A-INVENTORY1` | exact symbol/path, field, impact, touch mode, and evidence per affected row | pending |
| P2-A | `E2-P2A-MATRIX1` | unique affected denominator with zero duplicate/unassigned rows and explicit unaffected exclusions | pending |
| P2-A | `E2-P2A-REVIEW1` | unconditional Supervisor PASS for inventory | pending |
| P2-A | `E2-P2A-COMMIT1` | isolated documentation/source-audit commit | pending |
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

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E3-PNA-SUPERVISOR1` | independent acceptance of complete Child 02 source/diff/persistence/read/runtime evidence | pending |
| `E3-PNB-CLEANUP1` | dead-work inventory, removal, and Supervisor cleanup PASS | pending |
| `E3-PNC-DETECT1` | final detect-changes result | pending |
| `E3-PNC-HANDOFF1` | exact accepted persistence fields, affected-reader guarantees, evidence, and Child 03 refresh | pending |
| `E3-PNC-COMMIT1` | final closure commit and known worktree state | pending |
