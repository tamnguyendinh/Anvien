# Anvien Module Export And Re-Export Resolution Plan

## Metadata

- Date: `2026-07-28`
- Status: `CLOSED / P5-A 2560f914 / P5-B c1559df9 / P5-C 76899d45 / P5-D bb4cf465 / Pn-A Supervisor PASS / Pn-B committed at fcc44334 / Pn-C CLOSED`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-actual-status.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Persistence dependency: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
- Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`

## Goal

Make repository-backed TypeScript imports resolve through module export surfaces to the same terminal Symbol as direct imports, while preserving the syntactic module dependency and a source-backed re-export proof.

## Rules

- Complete and refresh P0 actual status before implementation work.
- The originating problem report supplies defect findings and target acceptance; its architecture proposals remain proposals unless current source evidence establishes that they are necessary.
- The causal synthesis and final Supervisor report verify the bounded defect only; they do not prescribe remediation design.
- Work directly on the current production source, build, command, and graph path.
- Existing resolver and analyze behavior is baseline evidence, not an immutable design. Change only the part whose necessity is established by the active slice.
- Child 05 owns module-request/path input assessment, export-table construction, re-export traversal, terminal binding, and proof for repository-backed imports only.
- Child 04 owns export syntax facts. Child 05 consumes its accepted facts and does not regenerate or reinterpret them.
- Child 02 owns preservation of corrected fields across persistence/readers. Child 05 changes only consumers shown by fresh impact evidence to depend on its new facts.
- Child 06 owns ambient and external declaration authority and resolution outcomes beyond repository-backed module exports.
- Before graph-based work, run `anvien analyze --force`. Before editing any resolver, provider, emitter, or shared contract, run `file-detail` and upstream `impact`, record the complete blast radius, and update the touch map.
- HIGH or CRITICAL impact is a scope warning. Keep the edit within the active slice and validate every affected boundary identified by evidence.
- Implement production behavior before adding or changing tests. Run the full repository build before boundary validation.
- Every completed slice requires focused behavior evidence, affected-boundary regression, ledger refresh, Supervisor PASS, `anvien detect-changes --repo E:\Anvien --scope all`, and an isolated commit.
- Each touched production or test file owns one primary semantic responsibility. Do not add unrelated logic or catch-all files.
- Analyze `E:\cheapapp.org` in place only for the named target-validation slice. Keep its source unchanged and keep plan, QA, and report artifacts in `E:\Anvien`.
- Update checklist, evidence, benchmark, and actual-status rows immediately when their state changes.

## Problem

The originating report identifies a bounded module-resolution failure: two imports reach a barrel file, but the resolver does not follow the barrel's re-export binding to the terminal function. The accepted causal synthesis locates the first confirmed divergence in `resolveImportedDef`: it searches physical definitions in the resolved file. A barrel with no matching physical declaration therefore produces no import binding, and the two consumer calls produce no terminal `CALLS` relationship.

Current source confirms that module-path resolution and symbol resolution are coupled after `ImportFact` creation. `resolveImportFiles` can find a target file, but `resolveImportedDef` searches only `defsByFile[targetFile]`. Current call resolution later reaches a global-name branch when scoped/import resolution is absent. Fixing only the target symptom would leave alias, star, cycle, ambiguity, and meaning errors unresolved or incorrectly guessed.

The defect does not prove that all package or path resolution must be redesigned. P5-A must first identify the actual current inputs and the minimum missing behavior.

## Scope

In scope:

- current `ImportFact`, module request, resolved path, module scope, and Child 04 export-fact inputs used by repository-backed TypeScript imports;
- separation of module/file lookup from exported-name lookup where current evidence requires it;
- deterministic per-module export surfaces built from accepted export facts;
- named/default/alias/type-only/star/namespace re-export traversal, cycle termination, ambiguity handling, meaning lanes, and proof retention to the extent established by Child 04 facts and slice fixtures;
- terminal call/use binding for the two bounded target sites;
- preservation of syntactic `IMPORTS` and physical path-resolution accounting;
- affected persistence/reader projection only when Child 02 inventory and fresh impact identify an actual consumer.

Out of scope:

- export syntax extraction or direct-export classification owned by Child 04;
- standard-library, ambient, intrinsic, or installed declaration authority owned by Child 06;
- a general package-resolution rewrite unless P5-A proves it is required and the plan is updated before code;
- unrelated graph-output, compatibility, or reader work;
- scanner exclusions and unrelated language-provider behavior;
- target-source edits or copied target fixtures.

## Non-Goals

- Do not expose all physical definitions as exports.
- Do not choose the first candidate to hide ambiguity.
- Do not use a repository-global same-name Symbol to rescue a failed explicit import.
- Do not create a consumer-to-implementation `IMPORTS` relationship when the source imports only the barrel.
- Do not preselect new owner files or project-profile infrastructure before P5-A impact/source evidence identifies the necessary owners.
- Do not claim complete TypeScript module-system conformance from the bounded target acceptance.

## Requirements

### Module and export boundaries

- P5-A records the current input contract before editing: syntax fact, source module/file, requested module text, resolved file candidates, imported/local/exported names, meaning/type-only information supplied by Child 04, and existing path/accounting outputs.
- Finding a module/file and finding an exported Symbol are separate results. A valid module result may still have a missing, ambiguous, cyclic, or meaning-incompatible export result.
- The export table is derived only from accepted Child 04 facts. Physical definitions are not implicit export entries.
- A barrel with zero physical declarations can expose a non-empty export surface when its syntax facts re-export terminal entries.
- Direct import and re-export lookup for the same declaration return the same terminal Symbol identity; only the proof path differs.
- Explicit entries take precedence over star-derived entries. Star traversal does not synthesize a default export.
- Multiple paths to the same terminal Symbol are deduplicated. Distinct terminal Symbols remain ambiguous; no ordering rule selects one.
- A cycle terminates. A cycle with a terminal branch can resolve that branch; a cycle without a terminal returns an explicit unresolved result.
- Requested value/type/namespace meaning is preserved through every hop. A meaning mismatch is not repaired by another name-based lookup.
- Every terminal result retains the consumer source site, imported name, barrel/re-export hops, target module/file, and terminal Symbol.
- An explicit import failure cannot fall through to a repository-global same-name target.

### Preservation and validation boundaries

- Syntactic `IMPORTS` continues to describe the source-written module dependency. Re-export traversal must not change the physical path-resolution count or syntactic `IMPORTS` count for the same input.
- Graph and persisted proof fields are added only at owners/consumers identified by Child 02 and fresh impact evidence.
- Synthetic fixtures live under the owning package's `testdata` and are independently authored. They cover zero-physical barrels, aliases, star precedence, cycles, ambiguity, meaning mismatch, and direct-versus-barrel equivalence.
- Target validation uses the same source/config and a freshly built current runtime. It proves exactly the two bounded calls and does not use aggregate graph counts as a substitute.

## Acceptance Criteria

- Direct and barrel imports of each bounded function resolve to the same terminal Symbol.
- The two bounded barrel call sites have the correct terminal `CALLS` relationships (`2/2`) and no corresponding false gap.
- A zero-physical-declaration barrel fixture has a non-empty syntax-derived export surface and complete terminal proofs.
- Named/default/alias/type-only/star/namespace fixtures pass the behavior defined by accepted Child 04 facts.
- Cycle fixtures terminate; ambiguity fixtures retain all distinct candidates and select none.
- Explicit import failure never binds to a repository-global same-name Symbol.
- Physical module-path resolution and syntactic `IMPORTS` absolute counts are recorded before and after traversal, with a delta of `0`.
- Only affected persistence/readers identified by Child 02 and fresh impact are changed, and their terminal Symbol/proof fields match the graph source.
- Full build, nearest real resolver/CLI boundary, target boundary, regression, Supervisor, detect-changes, and per-slice commit evidence all pass.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the current module fact, path lookup, physical-definition lookup, global-name rescue risk, target baseline, and blast radius.
  - Work Steps:
    1. Read the problem origin, bounded verification reports, current resolver/provider source, and all four Child 05 ledgers.
    2. Refresh graph evidence; run file-detail and impact for the current module/re-export owners; record current classifications and the predecessor block.
    3. Rewrite the next-slice assumptions from those facts without choosing implementation files beyond current evidence.
  - Implementation Gate: no production edit starts until the actual-status file has a final P0 decision and the Child 04 handoff is accepted.
  - Acceptance: actual status distinguishes existing path inputs, missing export surface, wrong physical-definition lookup, global-name rescue risk, affected-file counts, and the exact next action.
  - Evidence Targets: `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2`, `E0-P0A-GRAPH1`, `E0-P0A-SRC1..E0-P0A-SRC4`, `E0-P0A-FD1..E0-P0A-FD4`, `E0-P0A-IMPACT1..E0-P0A-IMPACT4`, `E0-P0A-DEPEND1`, `E0-P0A-SCOPE1`, `E0-P0A-BOUNDARY1`, and `E0-P0A-STATUS1`.

### P5: Module export and re-export resolution

- Phase Goal: resolve repository-backed imports through accepted export facts to terminal Symbols with accurate proofs.
- Phase Boundary:
  - In scope: P5-A through P5-D only.
  - Out of scope: Child 04 extraction, Child 06 declaration authority, and broad persistence work.
  - Dependencies: Child 04 closure; Child 02 affected-reader inventory when persisted fields change.
- Phase Implementation Rule: execute, validate, review, refresh, and commit one slice before opening the next.
- Ordered Slice List:
  - P5-A: Establish current module-request and path inputs.
  - P5-B: Build export tables from accepted export facts.
  - P5-C: Resolve aliases, re-exports, stars, cycles, ambiguity, and meanings.
  - P5-D: Emit terminal bindings and prove the target calls.

- [x] P5-A: Establish current module-request and path inputs.
  - Goal: identify and, only where evidence requires, correct the inputs that connect a source import to a resolved repository module before export lookup.
  - Scope Boundary:
    - Editable: `internal/scopeir/facts.go` for the requested-meaning contract; `internal/providers/tsjs/imports.go` for source-written TS/JS import population; `internal/scopeir/ir.go` and `internal/scopeir/sort_keys.go` for owning clone, canonicalization, and deterministic ordering of that contract.
    - Inspect-only: focused contract/provider tests selected after production code and the current resolver consumers needed to prove the new fields reach `buildWorkspace` unchanged.
    - Preserve-only: `internal/resolution/indexes.go` module/file candidate selection and `resolvedImport` separation; `internal/resolution/import_resolution.go` and every unaffected language strategy; accepted Child 04 `ExportFact` semantics; compatibility re-export `ImportFact` records as path compatibility only.
    - Out of scope: export traversal, declaration authority, and package-resolution expansion not proven necessary.
  - Non-Goals: no assumed project-profile layer, module-system rewrite, side-effect-only import fact expansion, dormant `ImportFact.Target*` activation/removal, or re-export semantic duplication.
  - Pre-flight Questions:
    - Data source: source-written TS/JS `ImportFact` records for module requests plus immutable accepted Child 04 `ExportFact` records for re-export semantics.
    - Display permission: N/A — resolver data has no display permission decision.
    - DB read flow: N/A — this slice inspects in-memory facts and repository files; persisted validation is later.
    - DB write flow: N/A unless fresh impact proves a persisted input is required.
    - Render location: resolver trace and evidence ledger.
    - UI behavior flow: N/A — no browser-visible behavior is owned.
    - Docker runtime: N/A — validate the built CLI/resolver boundary.
    - Playwright target: N/A — no UI surface is owned.
    - Behavior test: default/named/alias/namespace imports, statement-level and inline type-only imports, canonical requested-meaning sets, relative/index paths, unaffected-language regressions, and all three absolute count denominators.
    - Cleanup/quarantine: reusable fixtures only under package `testdata`.
    - External side effects: read-only repository/config inspection.
    - N/A notes: side-effect-only imports have no exported-name request and remain unchanged in P5-A; module export-table construction/traversal starts in P5-B/P5-C.
  - Work Steps:
    1. Use the accepted fresh inventory at HEAD `0aa49c87628c9e8b2041754515d6ebf0a930d55b`: preserve current module/file candidates and `resolvedImport`, retain accepted `ExportFact` as the sole re-export semantic source, and use the recorded six-file/fourteen-symbol impact plus the `5,072 / 5,072 / 5,088` count baseline.
       - UI flow check: N/A — non-UI resolver boundary.
       - DB/data flow check: preserve the exact source fact through resolved module/file result and keep the optional physical-definition result separate.
       - Render location check: evidence and benchmark ledgers.
       - Mini QA: inventory gate is complete; do not repeat graph analysis unless source changes invalidate it.
       - Evidence target: `E5-P5A-IMPACT1`, `E5-P5A-INPUT1`, `E5-P5A-COUNT1`.
    2. Add `RequestedMeanings []ExportMeaning` and `TypeOnly bool` to `ImportFact`. Treat `RequestedMeanings` as a canonical allowed-set, not a claim that the target supplies every lane: plain default/named/alias imports request `{value,type,namespace}`; statement-level or inline type-only forms request `{type}` and set `TypeOnly=true`; plain namespace imports request `{namespace}` with no exported-name request; type-only namespace imports request `{type}` and set `TypeOnly=true`. Leave both fields empty on non-TS/JS facts and compatibility re-export imports so unaffected providers and the accepted `ExportFact` authority remain unchanged.
       - UI flow check: N/A — non-UI resolver boundary.
       - DB/data flow check: source-written import syntax gains only requested semantic inputs; dormant output-looking `Target*` fields remain unused.
       - Render location check: `ImportFact` JSON plus focused ScopeIR/provider tests.
       - Mini QA: verify normal versus type-only import facts from the production extractor before changing tests.
       - Evidence target: `E5-P5A-SRC1`, `E5-P5A-TEST1`.
    3. Deep-clone, sort, deduplicate, and compare `RequestedMeanings` deterministically in `ScopeIR` normalization; include `TypeOnly` in import ordering. Add focused tests only after production code, run the full build and the nearest real non-UI resolver/CLI boundary, validate unaffected import strategies, prove zero delta for physical target-file resolutions (`5,072`), resolver-emitted `IMPORTS` (`5,072`), and final persisted graph-wide `IMPORTS` (`5,088`), then obtain Supervisor review, detect changes, and commit.
       - UI flow check: N/A — non-UI resolver boundary.
       - DB/data flow check: requested meanings are canonical and owned; module lookup remains deterministic and distinct from export lookup.
       - Render location check: resolver trace and test evidence.
       - Mini QA: exercise direct and barrel module lookup on the built runtime.
       - Evidence target: `E5-P5A-SRC1`, `E5-P5A-BUILD1`, `E5-P5A-TEST1`, `E5-P5A-REVIEW1`, `E5-P5A-DETECT1`, `E5-P5A-COMMIT1`.
  - Implementation Gate: satisfied for production implementation only after this planner refresh: Child 04 closure is accepted; exact editable/preserve-only owners are named; the requested-meaning/type-only representation and side-effect-import disposition are authorized; physical target-file resolutions (`5,072`), resolver-emitted `IMPORTS` (`5,072`), and final persisted graph-wide `IMPORTS` (`5,088`) are recorded.
  - Acceptance:
    - Source: current module-request/path inputs are explicit; TS/JS requested meanings/type-only state are represented exactly as authorized; compatibility re-export imports do not become a second semantic source; only the four proved production owners change.
    - Runtime/UI: built resolver/CLI lookup is correct; UI is N/A.
    - DB/data: module result is preserved separately from exported-name resolution.
    - Behavior test: default/named/alias/namespace, statement-level/inline type-only, canonicalization/round-trip, current import-path, target module lookup, and unaffected-language regressions pass.
    - Cleanup/quarantine: no copied target source or obsolete fixture remains.
    - Evidence IDs: `E5-P5A-IMPACT1`, `E5-P5A-INPUT1`, `E5-P5A-COUNT1`, `E5-P5A-SRC1`, `E5-P5A-BUILD1`, `E5-P5A-TEST1`, `E5-P5A-REVIEW1`, `E5-P5A-DETECT1`, `E5-P5A-COMMIT1`.
    - Actual-status rows refreshed: module-request/path input rows.
  - Evidence Targets: input manifest, counts, source diff, behavior tests, build, boundary, Supervisor, detect, commit.
  - Actual-status Update: requested name/meaning/type-only `partial|missing -> correct` only after evidence; preserve current module/file results and path strategies.
  - Commit Boundary: commit P5-A alone after acceptance.
  - Completion Record: Supervisor PASS `E5-P5A-REVIEW1`; detect-changes PASS `E5-P5A-DETECT1`; isolated commit `2560f914334e65961f755febdda6585840a4260e`.

- [x] P5-B: Build export tables from accepted export facts.
  - State: `SUPERVISOR PASS / DETECT RECORDED / COMMITTED` at isolated commit `c1559df953a277b099009f8489576d00ed25aa58`; P5-C is now the sole open slice; P5-D and target remain locked.
  - Goal: represent each repository module's syntax-derived export surface without treating physical definitions as implicit exports.
  - Scope Boundary:
    - Editable: new dedicated semantic owner `internal/resolution/export_tables.go`, plus minimal storage and wiring in `internal/resolution/indexes.go` (`workspace` and `buildWorkspace`) selected by fresh graph/source evidence.
    - Inspect-only: accepted Child 04 fact owners and P5-A module results.
    - Preserve-only: physical path lookup, `resolvedImport` module/file results, `resolveImports`, `resolveImportedDef`, and syntactic `IMPORTS` emission; terminal traversal remains P5-C.
    - Out of scope: terminal traversal and ambient declarations.
  - Non-Goals: no consumer call binding and no public-package API metric unrelated to the bounded defect.
  - Pre-flight Questions:
    - Data source: immutable Child 04 export facts plus P5-A module results.
    - Display permission: N/A — internal resolver structure.
    - DB read flow: N/A — construct from in-memory accepted facts.
    - DB write flow: only if affected persistence is proven by Child 02/fresh impact.
    - Render location: export-table trace and evidence.
    - UI behavior flow: N/A — no UI behavior.
    - Docker runtime: N/A — use full build and resolver boundary.
    - Playwright target: N/A — no UI surface.
    - Behavior test: local/default/alias/re-export/star/namespace/type-only facts, deterministic ordering, and zero-physical barrel.
    - Cleanup/quarantine: synthetic module fixtures under package `testdata`.
    - External side effects: none.
    - N/A notes: terminal resolution is P5-C.
  - Work Steps:
    1. Use the durable P5-B inventory `E5-P5B-IMPACT1` to implement production export-table construction in `internal/resolution/export_tables.go` from accepted facts, including explicit entries and star adjacency; add only the required `workspace` storage and `buildWorkspace` wiring in `internal/resolution/indexes.go`.
       - UI flow check: N/A — non-UI structure.
       - DB/data flow check: physical definitions do not become implicit entries.
       - Render location check: resolver trace/evidence.
       - Mini QA: exercise the built table boundary on the zero-physical barrel fixture.
       - Evidence target: `E5-P5B-IMPACT1`, `E5-P5B-SRC1`, `E5-P5B-ZEROBARREL1`.
    2. Add focused behavior tests after production code, run the full build, validate deterministic entries and unchanged P5-A counts, then review/detect/commit.
       - UI flow check: N/A — non-UI boundary.
       - DB/data flow check: entry identity, meaning, and source provenance are exact.
       - Render location check: evidence/benchmark ledgers.
       - Mini QA: inspect actual table output from the built resolver path.
       - Evidence target: `E5-P5B-BUILD1`, `E5-P5B-TEST1`, `E5-P5B-COUNT1`, `E5-P5B-REVIEW1`, `E5-P5B-DETECT1`, `E5-P5B-COMMIT1`.
  - Implementation Gate: satisfied. P5-A and Child 04 facts are accepted; the durable inventory report `E:\Anvien\reports\coder\rp_coder_260821_185542_by_gpt-5_child05_p5b_pre_implementation_ready.md` recorded the exact owner, CRITICAL blast radius, and E-only boundary before code; the authorized implementation has Supervisor PASS, while detect/commit remain Main-owned pending gates.
  - Acceptance:
    - Source: export tables are deterministic and derived only from accepted syntax facts.
    - Runtime/UI: built resolver table boundary passes; UI is N/A.
    - DB/data: zero-physical barrel exposes the expected non-empty surface; physical path and `IMPORTS` counts remain unchanged.
    - Behavior test: direct/default/alias/type-only/star/namespace table cases pass.
    - Cleanup/quarantine: only reusable accepted fixtures remain.
    - Evidence IDs: `E5-P5B-IMPACT1`, `E5-P5B-SRC1`, `E5-P5B-ZEROBARREL1`, `E5-P5B-BUILD1`, `E5-P5B-TEST1`, `E5-P5B-COUNT1`, `E5-P5B-REVIEW1`, `E5-P5B-DETECT1`, `E5-P5B-COMMIT1`.
    - Actual-status rows refreshed: export-table row.
  - Evidence Targets: source/table trace, fixture result, counts, tests, build, Supervisor, detect, commit.
  - Actual-status Update: export surface `missing -> correct`.
  - Commit Boundary: commit P5-B alone after acceptance.
  - Completion Record: Supervisor PASS `E5-P5B-REVIEW1`; `anvien detect-changes --repo E:\Anvien --scope all` exit `0` recorded as `E5-P5B-DETECT1`; exact 11-path manifest was staged and committed as `c1559df953a277b099009f8489576d00ed25aa58`, recorded as `E5-P5B-COMMIT1`.

- [x] P5-C: Resolve aliases, re-exports, stars, cycles, ambiguity, and meanings.
  - State: `SUPERVISOR PASS / DETECT RECORDED / COMMITTED` at isolated commit `76899d45a21fce55f6328b4cb30a6a5cb8719a81` after one source-backed REJECT and exact reject-only repair; P5-D is now the sole open slice, while target action remains gated by P5-D inventory and work-step authority.
  - Goal: resolve one requested export to a terminal repository Symbol or an explicit unresolved result with complete proof and no global-name rescue.
  - Scope Boundary:
    - Editable: new dedicated semantic owner `internal/resolution/export_resolution.go` for deterministic proof-bearing export lookup; `internal/resolution/indexes.go` only at `resolveImports`, `resolveImportedDef`, and `resolveImportedMember` plus removal of the now-redundant standalone `buildExportTables` call; `internal/resolution/resolve.go` only at `resolveCall` for an explicit-import-failure guard. Focused test owners are selected after production code and may include new `internal/resolution/export_resolution_test.go` plus bounded additions to existing resolver tests.
    - Inspect-only: accepted P5-A requested meanings, P5-B tables/tests, `bindingRef`/`resolvedImport`, current call emission, and affected tests named by `E5-P5C-IMPACT1`.
    - Preserve-only: Child 04 facts; P5-B `export_tables.go` data shape/construction; `resolveImportFiles`, `resolveImportFile`, `import_resolution.go`, `resolvedImport.TargetFiles`; generic `resolveName`, `resolveGlobalName`, and `resolveGlobalCallName`; `emitImportEdges`; graph/persistence/readers; syntactic dependency edges and all three accepted count denominators.
    - Out of scope: ambient/external declaration lookup and health projection.
  - Non-Goals: no first-candidate selection and no arbitrary topology ceiling without measured evidence.
  - Pre-flight Questions:
    - Data source: P5-B tables and requested name/meaning from accepted import facts.
    - Display permission: N/A — resolver behavior.
    - DB read flow: N/A — in-memory table traversal.
    - DB write flow: immutable resolution/proof result only; persisted adapters require proven impact.
    - Render location: resolver trace and evidence.
    - UI behavior flow: N/A — no UI behavior.
    - Docker runtime: N/A — built resolver/CLI boundary.
    - Playwright target: N/A — no UI surface.
    - Behavior test: alias chains, same-terminal dedupe, explicit-over-star, default exclusion, meaning mismatch, namespace, terminal/pure cycles, ambiguity, and missing export.
    - Cleanup/quarantine: bounded synthetic topology fixtures.
    - External side effects: none.
    - N/A notes: terminal graph binding is P5-D.
  - Work Steps:
    1. Use the accepted inventory `E5-P5C-IMPACT1`. Implement immutable/deterministic terminal, ambiguity, cycle, missing, and meaning-mismatch outcomes plus proof hops in dedicated `export_resolution.go`, consuming only P5-B tables and current repository definitions. Refactor `resolveImports` into two phases: first resolve and retain all module/file candidates, then build the accepted P5-B tables once, then resolve export terminals and create import bindings. Remove the separate post-`resolveImports` table call from `buildWorkspace`; do not add a corrective second pass or duplicate path resolution. Make `resolveImportedMember` consume the same export lookup for namespace/member requests.
       - UI flow check: N/A — non-UI resolver.
       - DB/data flow check: terminal identity and every hop are retained; distinct candidates remain distinct.
       - Render location check: resolver trace.
       - Mini QA: exercise the built resolver against every named topology fixture.
       - Evidence target: `E5-P5C-IMPACT1`, `E5-P5C-SRC1`, `E5-P5C-PROOF1`.
    2. At `resolveCall`, gate only the repository-global fallback when current import state proves the called identifier is an explicit import whose export lookup failed; keep generic global helpers unchanged for non-import calls. Add tests only after production behavior: alias/re-export chains, explicit-over-star precedence, star default exclusion, namespace/member traversal, same-terminal dedupe, distinct-terminal ambiguity, terminal and pure cycles, meaning mismatch, explicit-import no-global rescue, direct/path/`IMPORTS` preservation, and unaffected-language regression. Then run the full build and nearest real resolver/CLI boundary before Supervisor/detect/commit.
       - UI flow check: N/A — non-UI resolver.
       - DB/data flow check: one exact terminal, explicit ambiguity, or explicit unresolved result per lookup.
       - Render location check: evidence/benchmark ledgers.
       - Mini QA: inspect actual result/proof from the built runtime.
       - Evidence target: `E5-P5C-BUILD1`, `E5-P5C-TEST1`, `E5-P5C-NOGLOBAL1`, `E5-P5C-REVIEW1`, `E5-P5C-DETECT1`, `E5-P5C-COMMIT1`.
  - Implementation Gate: satisfied after this planner refresh is committed. P5-B is Supervisor-accepted and committed at `c1559df953a277b099009f8489576d00ed25aa58`; requested meaning is available; fresh graph/file-detail/upstream impact and exact owner selection are recorded in `E5-P5C-IMPACT1`. Production edits remain limited to the three owners above; P5-D/target stay locked.
  - Review Record: the initial Supervisor REJECT `1B8E32DA...` identified only distinct-owner member composition; the exact two-file repair is recorded as `E5-P5C-REPAIR1`; final Supervisor report `AC600E17...` records PASS as `E5-P5C-REVIEW1`. Final candidate identities are `export_resolution.go=566A69B9...`, `indexes.go=A44C94FA...`, `resolve.go=047CCDAE...`, and `export_resolution_test.go=97FF4990...` at unchanged HEAD `861000cb6b6e36ce105623f0dc8c093b089f61fa`. P5-C remains unchecked while detect-changes and isolated commit remain pending.
  - Acceptance:
    - Source: traversal follows export facts, terminates cycles, retains ambiguity, and never uses global-name rescue for explicit imports.
    - Runtime/UI: built resolver cases pass; UI is N/A.
    - DB/data: proof hops and terminal identity are deterministic.
    - Behavior test: all named topology/meaning cases pass.
    - Cleanup/quarantine: no abandoned topology or generated debug artifact remains.
    - Evidence IDs: `E5-P5C-IMPACT1`, `E5-P5C-SRC1`, `E5-P5C-PROOF1`, `E5-P5C-BUILD1`, `E5-P5C-TEST1`, `E5-P5C-NOGLOBAL1`, `E5-P5C-REPAIR1`, `E5-P5C-REVIEW1`, `E5-P5C-DETECT1`, `E5-P5C-COMMIT1`.
    - Actual-status rows refreshed: re-export traversal and explicit-import global-name-rescue rows.
  - Evidence Targets: impact, source diff, proof vectors, behavior tests, build, regression, Supervisor, detect, commit.
  - Actual-status Update: re-export lookup `wrong -> correct`; terminal graph binding remains pending.
  - Commit Boundary: commit P5-C alone after acceptance.
  - Completion Record: Supervisor PASS is recorded as `E5-P5C-REVIEW1`; exactly one fresh analyze and full detect result are recorded as `E5-P5C-DETECT1`; exact 12-path manifest was staged and committed as `76899d45a21fce55f6328b4cb30a6a5cb8719a81`, recorded as `E5-P5C-COMMIT1`.

- [x] P5-D: Emit terminal bindings and prove the target calls.
  - State: `SUPERVISOR PASS / DETECT RECORDED / COMMITTED` at isolated commit `bb4cf46509716259c3bf24a1ca041a6e763d5419`. Work Step 1 report `reports/coder/rp_coder_260822_003300_by_gpt-5_child05_p5d_proof_projection_ready_for_supervisor.md` is `18,358` bytes / `285` LF / SHA-256 `97235D73735497328DD7DE41DB0B5023FC6A7ACC05CC46F362A965D0BDB0FB18`; target report `reports/coder/rp_coder_260822_011119_by_gpt-5_child05_p5d_target_validation_target_ready.md` is `22,185` bytes / `453` LF / SHA-256 `2740EDB71ADFFE1242E192DAEC86C70805F6B656568B49C95C089B981EFECDF2`; Supervisor PASS report `reports/Supervisor/rp_supervisor_260822_013836_by_gpt-5_child05_p5d_full_acceptance.md` is `16,643` bytes / `198` LF / SHA-256 `2A9EC861D443539749C8478D307416EDC257C9178C29049BC5C76309C420471A`.
  - Goal: bind source calls/uses to P5-C terminal Symbols while preserving source-written module dependencies and proof.
  - Scope Boundary:
    - Editable: new `internal/resolution/export_binding_proof.go` as the sole P5-D proof-to-`graph.Evidence` adapter; `internal/resolution/indexes.go` only for owned retention of the already-computed semantic export result and a proof-returning namespace/member seam; `internal/resolution/resolve.go` only at `resolveCall`/`resolveAccess` for proof attachment; `internal/resolution/emit.go` only at `mergeRelationship` for deterministic Evidence union/dedupe.
    - Inspect/test/validate-only: accepted P5-C result/proof types; `internal/graph/types.go`; Graph JSON; Ladybug CSV/schema/COPY/fallback; MCP context and impact; HTTP graph transport only if selected as the nearest public boundary.
    - Preserve-only: accepted P5-B/P5-C source and tests, `emitter.emitReference`, `graph.Evidence` field shape, Graph JSON/Ladybug/MCP/HTTP/Web production transport, file-detail/UI, Child 02 C09-C16, syntactic `IMPORTS`, path resolution, target source/worktree, and unrelated relationships.
    - Out of scope: typed Evidence fields, relationship schema/column changes, decoded UI models, ambient/external outcomes, unrelated reader changes, and P5-C semantic traversal changes.
  - Non-Goals: no synthetic dependency from consumer to implementation module.
  - Pre-flight Questions:
    - Data source: the exact `exportResolutionResult`, candidates, failures, ordered proofs, and hops already produced once by accepted P5-C, plus the call/access source-site ID.
    - Display permission: existing machine-readable graph/CLI/MCP visibility only; no UI consumer interprets `Relationship.Evidence`, so UI/Playwright is N/A.
    - DB read flow: Graph JSON, Ladybug relationship evidence, MCP context, and MCP impact are the exact four affected-reader denominator.
    - DB write flow: normal analyze writes Graph JSON and Ladybug through the existing generic Evidence channel; no schema migration.
    - Render location: CALLS/ACCESSES `Relationship.Evidence` plus those four readers.
    - UI behavior flow: N/A because fresh file-detail/impact and source inspection found no semantic UI proof consumer.
    - Docker runtime: N/A; this is the non-UI canonical full-build/analyze/CLI/MCP/Ladybug boundary.
    - Playwright target: N/A.
    - Behavior test: direct/barrel terminal equality, aliases/stars/namespace/member branches, cycle/ambiguity/failure provenance, coalesced source-site conservation, `2/2` target terminal calls, zero matching false gaps, complete proofs, and unchanged path/`IMPORTS` counts.
    - Cleanup/quarantine: reports stay in Anvien; no target-source artifact.
    - External side effects: target `.anvien` may be regenerated by normal analyze; source remains unchanged.
    - N/A notes: Child 02 C09-C16 are Definition identity/selection/embedding readers, not P5-D export-proof readers; preserve their accepted behavior.
  - Work Steps:
    1. After this docs authorization commit, resume the existing E-only Coder and implement production first in exactly the four owners above. Retain one owned P5-C semantic result per resolved import without a second traversal/path pass; expose the same result from namespace/member resolution while preserving the existing wrapper for unaffected consumers. Convert proof data to the existing `graph.Evidence{Kind,Weight,Note}` shape only: keep the generic call/access evidence entry first; append `export-terminal-v1`, `export-hop-v1`, and `export-failure-v1` records with `Weight=1`; marshal `Note` from concrete versioned structs, never maps. Every Note includes `sourceSiteId`; terminal Notes include outcome/terminal kind/requested name/canonical meanings/canonical target files/terminal definition and graph identity/file/namespace file; hop Notes include deterministic proof and hop ordinals plus the complete accepted `ExportFact`, target file, owner/member, and terminal IDs; failure Notes include deterministic proof ordinal, outcome, failure file/name, namespace file, and canonical meanings. Dedupe by exact `(Kind, Weight, Note)`. At coalescing, preserve the existing generic first item so `relationshipCallName()` and semantic edge keys stay unchanged, then stable-union remaining generic records and sort only the P5-D records by kind rank, proof ordinal, hop ordinal, and canonical Note.
       - UI flow check: N/A — no semantic UI consumer.
       - DB/data flow check: P5-C terminal/proof is retained once, attached to CALLS/ACCESSES, and survives coalescing without changing endpoints, `ProofKind`, source-site IDs, or the Evidence schema.
       - Render location check: `Relationship.Evidence`; Graph JSON/Ladybug/MCP are validate-only downstream readers.
       - Mini QA: after production code, add only the authorized focused tests, run focused/full resolution, then the canonical full build before Graph JSON/Ladybug/MCP parity. HTTP is validate-only only if needed as the nearest public boundary.
       - Evidence target: `E5-P5D-IMPACT1`, `E5-P5D-SRC1`, `E5-P5D-BUILD1`, `E5-P5D-TEST1`, `E5-P5D-PARITY1`.
    2. Only after Main verifies the local Coder-ready handoff, final source/test bytes, canonical full build, and four-reader parity may Main explicitly open the target work step. Then capture target pre-state; invoke the freshly built canonical `E:\Anvien` runtime once from `E:\cheapapp.org`; prove direct/barrel equality, `2/2` terminal calls, zero matching false gaps, complete ordered proofs, `0 / 0 / 0` path/emitted-`IMPORTS`/persisted-`IMPORTS` delta, and unchanged target source/worktree; then hand the durable candidate to the existing Supervisor before detect/commit.
       - UI flow check: N/A unless a new invalidating impact proves a semantic UI consumer and the plan is refreshed first.
       - DB/data flow check: fresh target graph is the sole acceptance source for the two sites.
       - Render location check: official evidence under Anvien.
       - Mini QA: inspect the two source-site relationships through Graph JSON, Ladybug, MCP context, and MCP impact; no manufactured UI scope.
       - Evidence target: `E5-P5D-TARGET1`, `E5-P5D-ORACLE1`, `E5-P5D-BOUNDARY1`, `E5-P5D-COUNT1`, `E5-P5D-REVIEW1`, `E5-P5D-DETECT1`, `E5-P5D-COMMIT1`.
  - Implementation Gate: satisfied for production only after this exact five-path docs authorization commit (four living ledgers plus the immutable inventory report). `E5-P5D-IMPACT1` records the fresh graph `115,947 / 159,626`, all sixteen file-detail rows, all nineteen HIGH/CRITICAL/LOW impact tuples, the four-reader denominator, the existing-shape encoding, and the exact production/test manifest. Any typed-field/schema/UI expansion invalidates this authorization and requires a new planner/impact gate. Target remains locked until Work Step 1 passes and Main explicitly opens Work Step 2.
  - Acceptance:
    - Source: call/use emission consumes the retained P5-C terminal result/proofs exactly once; generic evidence and semantic edge identity remain first/unchanged; coalescing retains every distinct source-site proof.
    - Runtime/UI: target calls resolve `2/2`; the four affected machine readers show the same terminal/proof; UI/Playwright is N/A unless new invalidating impact forces a plan refresh.
    - DB/data: Graph JSON, Ladybug, MCP context, and MCP impact retain identical terminal/proof records with zero field loss; graph schema/reader production bytes remain unchanged; path/`IMPORTS` deltas are `0 / 0 / 0`.
    - Behavior test: direct/barrel equality, exact encoding/order/dedupe, namespace/member and failure branches, coalesced source sites, two target calls, zero matching gaps, unchanged P5-C matrix, and unaffected-language regression pass.
    - Cleanup/quarantine: target source/worktree unchanged; obsolete debug evidence removed.
    - Evidence IDs: `E5-P5D-IMPACT1`, `E5-P5D-SRC1`, `E5-P5D-BUILD1`, `E5-P5D-TEST1`, `E5-P5D-PARITY1`, `E5-P5D-TARGET1`, `E5-P5D-ORACLE1`, `E5-P5D-BOUNDARY1`, `E5-P5D-COUNT1`, `E5-P5D-REVIEW1`, `E5-P5D-DETECT1`, `E5-P5D-COMMIT1`.
    - Actual-status rows refreshed: terminal binding, target acceptance, and affected readers.
  - Evidence Targets: impact, source/build/tests, affected parity, target oracle/boundary/counts, Supervisor, detect, commit.
  - Actual-status Update: terminal binding `wrong -> correct`; bounded target `0/2 -> 2/2`.
  - Commit Boundary: commit P5-D alone after acceptance.
  - Completion Record: Supervisor PASS `E5-P5D-REVIEW1`; fresh detect `E5-P5D-DETECT1`; exact 15-path isolated commit `bb4cf46509716259c3bf24a1ca041a6e763d5419` recorded as `E5-P5D-COMMIT1`. Pn-A is now the sole open item.

- [x] Pn-A: Call Supervisor for Child 05 acceptance.
  - State: `INITIAL SUPERVISOR REJECT / LEDGER-ONLY REPAIR / REJECT-ONLY SUPERVISOR PASS`. Pn-A is closed; Pn-B and Pn-C have since closed. No Child 05 item remains open.
  - Goal: independently verify all four slices, source diff, runtime evidence, target boundary, ledgers, and commits.
  - Work Steps:
    1. Run the Supervisor review over the complete Child 05 scope.
    2. Return only rejected invariants to the owning slice; repeat validation and review until PASS or a recorded blocker.
  - Implementation Gate: P5-A through P5-D are accepted locally or explicitly blocked with evidence.
  - Acceptance: `E5-PNA-REVIEW1` records the precise ledger-closure blocker; `E5-PNA-REVIEW2` must record reject-only Supervisor PASS before Pn-A can be checked.
  - Review Record: initial child-wide report `reports/Supervisor/rp_supervisor_260822_021218_by_gpt-5_child05_pna_child_wide_acceptance_reject.md`, `11,899` bytes / `121` LF / SHA-256 `E3E28E79D5BC929E7F8FAA29F67D6C26E881A4E1AEEACA8CC95FF61207D0CB28`, verdict `REJECT`; production, tests, commits, and sealed target evidence are cleared, while only living-ledger/evidence closure is rejected.
  - Reject-Only Resubmission Gate: `E5-PNA-LEDGER1` reconciles the four living ledgers and `E5-PNA-MEASURE1` closes the two pending generic-Evidence benchmark cells without build/analyze/target reruns; resume only the existing Supervisor lane for `E5-PNA-REVIEW2`.
  - Completion Record: `E5-PNA-REVIEW2` records reject-only Supervisor PASS at `reports/Supervisor/rp_supervisor_260822_023256_by_gpt-5_child05_pna_ledger_closure_resubmission_pass.md`, `8,542` bytes / `137` LF / SHA-256 `1A11CCF1AA5279E03F0FF06B0E057EB2BFB4267F661496788828CB6BF46E3C68`; residual same-invariant surfaces are none.

- [x] Pn-B: Remove dead work created by Child 05.
  - State: `SUPERVISOR PASS / COMMITTED`. Cleanup removed exactly one failed, unreferenced Child 05 debug capture; all accepted and traceability-bearing artifacts were preserved. The exact six-path cleanup/docs/report slice is committed at `fcc44334c0f75b3b19046dc8f9f4de40eb459fa9`.
  - Goal: retain only accepted production, test, fixture, evidence, and ledger artifacts.
  - Work Steps:
    1. Identify failed, duplicate, superseded, or unused Child 05 artifacts.
    2. Remove only those artifacts, verify the final diff, and obtain Supervisor confirmation.
  - Implementation Gate: do not remove pre-existing user work or artifacts owned by another child.
  - Acceptance: `E5-PNB-CLEAN1` records the complete inventory, exact deletion, seven retained P5 temp artifacts, clean tracked/index boundary, and independent Supervisor PASS.
  - Completion Record: Coder handoff `reports/coder/rp_coder_260822_024606_by_gpt-5_child05_pnb_cleanup_ready_for_supervisor.md`, SHA-256 `15DA59857DE9C401C67D5CD9C65726C13554F33C40EF13BD3C2BDAD6C4DEBCB1`; Supervisor PASS `reports/Supervisor/rp_supervisor_260822_025711_by_gpt-5_child05_pnb_cleanup_acceptance_pass.md`, SHA-256 `E98B199011AAF2795F2C115296B7318230C52D1E7D67E6451CE01FB6B2889B6D`; residual same-invariant surfaces none; exact six-path commit `fcc44334c0f75b3b19046dc8f9f4de40eb459fa9`.

- [x] Pn-C: Close and hand off Child 05.
  - State: `CLOSED / DOCS-ONLY CLOSURE`. The Child 05 plan is declared closed by this exact living-plan state; the closure commit and direct Child 06 handoff follow immediately without another gate.
  - Goal: close the current Child plan, commit that exact living-plan closure, and hand off the accepted boundary to Child 06.
  - Work Steps:
    1. Declare the current Child 05 plan `CLOSED` in the living plan set.
    2. Commit the exact living-plan closure immediately.
    3. Hand off the accepted Child 05 facts/results and opening boundary to Child 06.
  - Implementation Gate: satisfied — Pn-A and Pn-B have Supervisor PASS and the Pn-B cleanup/docs/report slice is committed at `fcc44334c0f75b3b19046dc8f9f4de40eb459fa9`.
  - Acceptance: `E5-PNC-CLOSE1` is this exact closed four-ledger state; `E5-PNC-COMMIT1` is the immediate exact living-plan closure commit; `E5-PNC-HANDOFF1` is the immediate post-commit Child 06 transfer recorded by the successor opening.
  - Hard Boundary: no audit/build/test/analyze/detect/file-detail/impact/QA/target action, dedicated closure report, worker/Supervisor loop, or successor-plan work may be mixed into Pn-C.
  - Completion Record: Child 05 is `CLOSED`; accepted implementation commits are `2560f914`, `c1559df9`, `76899d45`, and `bb4cf465`; Pn-A closure is `b68e738d`; Pn-B closure is `fcc44334`. The resulting exact four-ledger closure commit is the sole handoff anchor for Child 06.

## Risk Notes

- `resolveImportedDef`, `resolveImports`, and `buildWorkspace` have CRITICAL upstream impact; current evidence spans resolver, analyze, access-audit, multiple languages, 23 linked tests, and up to 29 linked flows.
- Re-export traversal can produce cycles and multiple candidates. Tests must prove termination and ambiguity without arbitrary selection.
- Export meaning depends on the accepted Child 04 fact contract. Starting P5 before that handoff risks rebuilding syntax semantics in the wrong owner.
- Current path strategies cover many languages. P5-A must preserve unaffected strategies and avoid a TypeScript fix that changes other providers.
- Target aggregate call/gap counts can change because of other children. Acceptance uses the exact two source sites and their terminal proofs, not an aggregate coincidence.
- Any newly discovered affected reader expands only the owning slice after the plan and actual-status touch map are updated; it does not authorize broad adapter work.
