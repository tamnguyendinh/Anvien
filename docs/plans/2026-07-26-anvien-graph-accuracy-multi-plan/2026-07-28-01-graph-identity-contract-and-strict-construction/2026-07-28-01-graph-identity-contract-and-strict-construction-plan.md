# Anvien Graph Identity Contract and Strict Construction Plan

## Metadata

- Date: `2026-07-28`
- Status: `active / P0, P1-A through P1-E, and Pn-A accepted at isolated commit boundaries / P1-E remains explicit Owner acceptance, not a fabricated Supervisor PASS / Pn-A independent Child-level Supervisor PASS / Pn-B is the only open slice / Pn-C and Child 02 closed`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Predecessor: none
- Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`

## Goal

Make declaration and symbol identity precise enough to conserve distinct source occurrences, lexical ownership, meaning, and ranges; remove the proven silent collision/loss behavior; and raise the bounded `time`/`now` result from `2/4` to `4/4` through the normal production runtime.

## Rules

- Complete P0 actual status before implementation work.
- The original problem report supplies the identity finding and target oracle; its DRAFT repair proposal is not implementation authority.
- Current source, fresh graph evidence, file-detail, impact, and real behavior determine the implementation owner and whether any pipeline component must change.
- The analyze implementation present at P0 is baseline evidence, not a fixed implementation sequence or file-layout constraint.
- Work directly in the repository-root production codebase with the normal command grammar.
- Graph persistence mechanics are outside this Child.
- Change only the identity/range/occurrence/collision owners proven by the active slice. Persistence, readers, binding patterns, exports, module resolution, and external declarations belong to later Children.
- If P1-C or P1-D impact proves that an additional producer or relationship contract must change, update this plan and actual status before editing it.
- Implement production behavior before changing tests. Run the full build before real-boundary validation.
- Before editing a graph builder, resolver, analyzer, shared contract, function, class, or method, refresh the graph and record file-detail plus upstream impact. HIGH/CRITICAL is a scope warning, not an edit ban.
- Every touched production or test file owns one primary semantic responsibility. Do not add catch-all owners or unrelated refactors.
- Update the matching checklist, actual status, evidence, and benchmark rows as soon as their state changes.
- Each accepted implementation slice requires Supervisor PASS, Anvien detect-changes, and an isolated commit before the next slice opens.
- Hidden data substitution is forbidden. A failed command or invalid graph reports the failure at its normal boundary.
- Target validation analyzes `E:\cheapapp.org` in place. Target source and pre-existing worktree state are preserve-only; normal target output remains under the target repository.
- Debug artifacts stay under repository-local `.tmp`; reusable fixtures stay under the owning package's `testdata` and are independently authored.
- Scanner treatment of the eight known omitted paths is outside this Child.

## Problem

The bounded investigation traced four distinct source declarations (`time` twice and `now` twice) through ScopeIR, but the baseline graph retained only two nodes. The first confirmed graph divergence is the identity computed without sufficient scope/range information, followed by replacement of a duplicate node ID.

This is an identity and occurrence-conservation defect. It does not establish that every graph producer, every relationship identity, or every reader must be redesigned. The exact owner set must be proven during P0 and each slice pre-flight.

## Scope

In scope:

- source-backed declaration and symbol identity rules;
- lexical owner and meaning inputs required to distinguish declarations;
- full range and selection-range facts required by identity and source navigation;
- conservation of declaration occurrences from provider/ScopeIR input into graph output;
- removal of silent collision, replacement, or skip behavior at the proven owner;
- deterministic and integrity-valid graph construction for the accepted identity facts;
- bounded target validation for `time` and `now`.

## Non-Goals

- persistence format and reader changes;
- TypeScript binding-pattern extraction;
- export, module/re-export, ambient/external, or graph-health semantics;
- a mandatory relationship-identity redesign or broad producer rewrite without P1-C/P1-D impact evidence;
- a predetermined node topology, hash function, package split, or implementation filename;
- scanner repair, target-source edits, or unrelated legacy refactoring;
- complete TypeScript language conformance.

## Requirements

- A declaration represents one source occurrence; a symbol represents the semantic entity to which the declaration belongs.
- Same-name declarations in different lexical owners must remain distinct.
- Meaning lanes that the provider distinguishes must not collapse into one name-only identity.
- Range covers the construct and selection range identifies the declaring token when available. Encoding, coordinate base, and interval semantics must be explicit and source-backed.
- Identity inputs must be deterministic and repository-relative; traversal order, worker order, random values, and absolute checkout paths must not alter accepted identity sets.
- Declaration merging or overload sharing requires provider evidence; name equality alone is insufficient.
- Every input occurrence required by the oracle must remain represented after identity construction.
- A conflicting identity must not cause silent replacement, first-win, last-win, or skip behavior.
- Relationship endpoints emitted from corrected identities must exist.
- The normal analyze command must expose either a valid graph satisfying these invariants or a clear failure.
- Implementation files are selected only after current file-detail and impact evidence identifies the true owner. The plan is updated before an evidence-backed scope expansion.

## Acceptance Criteria

- The bounded target represents both `time` declarations and both `now` declarations: `4/4`.
- Declaration occurrence conservation for every Child 01 fixture and bounded target site is `100%`, with zero unexplained drops.
- Proven collision groups fall from `2` to `0` without conflating lexical owners or meanings.
- Reordered equivalent identity inputs produce the same canonical identity set; repeated matched analyze runs are deterministic.
- Corrected ranges and selection ranges round-trip at their nearest owned boundary.
- Graph integrity checks report zero missing endpoints for the affected records.
- A deliberately conflicting canonical fact produces a clear failure rather than silent loss.
- No later-Child semantic behavior is implemented in this Child.
- Every implementation slice passes full build, nearest real boundary, ledger refresh, Supervisor, detect-changes, and isolated commit gates.

## Invariant Family Map

| Invariant family | Source truth / current owner | First affected boundary | Owning slice | Sibling surfaces to preserve | Forbidden fallback |
|------------------|-----------------------------|-------------------------|--------------|-----------------------------|--------------------|
| Identity injectivity | `DefinitionFact.ID`, `ScopeFact` ownership, `graphIDForDef` | `internal/resolution/indexes.go:814-823` | P1-C | provider occurrence IDs and language meaning lanes | name-only, range-only, random, absolute-path, or traversal-order identity |
| Range and selection | `Range`, TSJS `nodeRange`, construct/name AST nodes | provider/IR contract and graph node projection | P1-B, then P1-E | `PositionIndex` and reference ranges until semantics are ratified | guessed UTF-16/byte conversion or unannounced inclusive/exclusive change |
| Lexical owner and meaning | `ScopeFact.Parent/OwnedDefIDs`, `scopeByDef`, `OwnerID`, `Label` | graph identity omits scope/meaning inputs | P1-B/P1-C | member ownership, type/value distinctions, later language semantics | global-name fallback or treating name equality as a merge rule |
| Occurrence conservation | provider/ScopeIR definitions and production `defsByFile` | graph emission attempts collide after `graphIDForDef` | P1-C | standalone `BuildDefinitionIndex` legacy tests; later persistence denominator | using `len(defsByID)` or graph node counts as the input denominator |
| Collision and mutation | `Graph.AddNode`, `Graph.init`, `emitter.emitNode` | duplicate ID performs silent last-writer replacement | P1-D | legitimate enrichment/reinsertion callers and separate relationship merge | blanket `AddNode` policy change, warning-only loss, or silent skip |
| Projection and endpoints | `emitDefinitionNodes`, `DEFINES`, owner edges | node properties omit columns/scope/provider ID; colliding endpoints survive | P1-C/P1-E | `emitRelationship` semantic edge merge, persistence/readers in Child 02 | endpoint existence alone as conservation proof or persistence redesign here |

Sibling boundary: `emitRelationship`/`AddRelationship` is a separate relationship pipeline; persistence, readers, bindings, exports, module resolution, ambient outcomes, and scanner behavior remain owned by later Children. The P0 source audit proves no orchestration change is required at this stage.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish current identity/range/occurrence owners and the exact implementation boundary.
  - Work Steps:
    1. Read the problem-origin report and bounded verification reports; separate measured identity facts from DRAFT design.
    2. Refresh the Anvien graph; inspect the normal provider/ScopeIR/workspace/emission path; run file-detail and impact for every candidate identity/range/collision owner and supporting shared contract.
    3. Distinguish the production occurrence denominator from standalone legacy indexes, record actual file counts, blast radius, current behavior classifications, touch modes, and the exact consequence for P1-A through P1-E.
    4. Obtain Supervisor review of the P0 classification before implementation opens.
  - Implementation Gate: no production edit until every target unit has current evidence and the actual-status Final P0 Decision is complete.
  - Acceptance: `E0-P0A-REPORT1`, `E0-P0A-VERIFY1`, `E0-P0A-GRAPH1`, `E0-P0A-SOURCE1..E0-P0A-SOURCE9`, `E0-P0A-FD1..E0-P0A-FD15`, `E0-P0A-IMPACT1`, `E0-P0A-IMPACT2`, `E0-P0A-BOUNDARY1`, `E0-P0A-BOUNDARY2`, `E0-P0A-STATUS1`, `E0-P0A-STATUS2`, and `E0-P0A-REVIEW1` are recorded, current, and non-contradictory.

### P1: Graph identity and strict construction

- Phase Goal: implement only the source-proven identity/range/occurrence/collision corrections and validate them through the normal runtime.
- Phase Boundary:
  - In scope: P1-A through P1-E below.
  - Out of scope: persistence/readers and Children 03-06 semantics.
  - Dependencies: P0 complete; each slice accepted and committed before its successor.
- Phase Implementation Rule: do not implement P1 as one change. Execute the five ordered slices independently.
- Ordered Slice List:
  - P1-A: Ratify the source-backed identity contract.
  - P1-B: Establish range, selection range, and identity inputs.
  - P1-C: Implement declaration/symbol identity and occurrence conservation.
  - P1-D: Remove silent collision or overwrite at the proven owner.
  - P1-E: Integrate and validate identity accuracy.

- [x] P1-A: Ratify the source-backed identity contract.
  - Goal: accept only the identity invariants supported by the report, current source, and P0 evidence.
  - Scope Boundary:
    - Editable: `docs/contracts/graph-accuracy-contract.md` and this Child's four ledgers.
    - Inspect-only: current identity/range/graph-construction source and bounded reports.
    - Preserve-only: production source, tests, runtime, and target repository.
    - Out of scope: implementation design not proven by P0.
  - Non-Goals: no production edit, implementation filename mandate, or later-Child contract.
  - Pre-flight Questions:
    - Data source: P0 source evidence and bounded identity oracle.
    - Display permission: N/A — no UI or permission behavior.
    - DB read flow: N/A — documentation/source review only.
    - DB write flow: N/A — no graph or database write.
    - Render location: contract and Child ledgers.
    - UI behavior flow: N/A — no visible app behavior.
    - Docker runtime: N/A — documentation-only slice.
    - Playwright target: N/A — document/link/Supervisor boundary.
    - Behavior test: contract-to-source contradiction audit.
    - Cleanup/quarantine: remove superseded contract text from active authority.
    - External side effects: none.
    - N/A notes: production/runtime gates do not apply because this slice changes documentation only.
  - Work Steps:
    1. Reconcile declaration/symbol, lexical owner, meaning, range, occurrence, collision, and ownership rules with P0 evidence.
       - UI flow check: N/A — documentation-only.
        - DB/data flow check: verify that persistence remains outside this documentation slice.
       - Render location check: contract and ledgers.
       - Mini QA: validate links, slice/evidence mappings, and forbidden architecture terms.
       - Evidence target: `E1-P1A-SOURCE1`, `E1-P1A-CONTRACT1`.
    2. Obtain zero-trust Supervisor review and commit only the accepted documentation slice.
       - UI flow check: N/A — documentation-only.
       - DB/data flow check: N/A — no runtime data changed.
       - Render location check: Supervisor report and evidence ledger.
       - Mini QA: re-run document integrity checks after review corrections.
       - Evidence target: `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1`.
  - Implementation Gate: P0 is complete and every contract rule has current source or acceptance justification.
  - Acceptance:
    - Source: production source unchanged.
    - Runtime/UI: N/A — no runtime/UI change.
    - DB/data: no unproven persistence or writer rule remains.
    - Behavior test: contract, roadmap, and Child ownership agree.
    - Cleanup/quarantine: superseded active contract text is absent.
    - Evidence IDs: `E1-P1A-SOURCE1`, `E1-P1A-CONTRACT1`, `E1-P1A-REVIEW1`, `E1-P1A-COMMIT1`.
    - Actual-status rows refreshed: contract authority and P1-B assumptions.
  - Evidence Targets: source/contract trace, document audits, Supervisor, commit.
  - Actual-status Update: contract `partial -> correct` or record the exact blocker.
  - Commit Boundary: documentation-only commit after acceptance.

- [x] P1-B: Establish range, selection range, and identity inputs.
  - Goal: make the source-proven position and owner/meaning inputs available without yet changing graph collision behavior.
  - Scope Boundary:
    - Editable: only the exact position/identity-input symbols selected after the P1-A contract gate, currently candidates in `internal/scopeir/facts.go`, `internal/scopeir/range.go`, `internal/providers/tsjs/nodes.go`, and the relevant `addDefinition` path in `internal/providers/tsjs/definitions.go`.
    - Inspect-only: `internal/providers/tsjs/scopes.go`, normalization/sort owners, identity construction, graph mutation, projection, persistence, and readers.
    - Preserve-only: the TSJS non-identifier binding-pattern branch owned by Child 03, unrelated language/provider semantics, and target source.
    - Out of scope: collision handling, broad producer rewrite, and persistence.
  - Non-Goals: no range-only identity shortcut or guessed encoding conversion.
  - Pre-flight Questions:
    - Data source: provider parse facts and current ScopeIR position/owner/meaning data.
    - Display permission: N/A — no UI permission behavior.
    - DB read flow: nearest in-memory/provider fact boundary.
    - DB write flow: N/A — no persisted graph contract in this slice.
    - Render location: source/fixture trace and command evidence.
    - UI behavior flow: N/A — non-UI analyzer contract.
    - Docker runtime: N/A — validate through full build and analyzer boundary.
    - Playwright target: N/A — no browser surface.
    - Behavior test: provider-supported range/selection range, coordinate encoding, interval boundary, line-ending, Unicode, tab, and nested-scope cases; owner/meaning inputs are checked but not redesigned here.
    - Cleanup/quarantine: package-local independently authored fixtures only.
    - External side effects: none.
    - N/A notes: persistence and UI are later boundaries.
  - Work Steps:
    1. Refresh graph, record exact file-detail/impact for the shared contract and provider position symbols, and select the smallest source owner set without opening Child 03 binding work.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: trace provider input to the identity-input record.
       - Render location check: evidence ledger.
       - Mini QA: exercise the nearest provider/ScopeIR boundary.
       - Evidence target: `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`.
    2. Implement production inputs, then focused tests; run the full build and real analyzer-boundary validation.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: every accepted input round-trips without silent normalization drift.
       - Render location check: command/test evidence.
       - Mini QA: run the built runtime against the identity fixture.
       - Evidence target: `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1`.
  - Implementation Gate: P1-A accepted; exact owners and HIGH/CRITICAL blast radius recorded.
  - Acceptance:
    - Source: source-backed full/selection position and owner/meaning inputs exist in their single-responsibility owners with explicit coordinate/interval semantics; no field name or topology is prescribed without proof.
    - Runtime/UI: built analyzer exposes the expected inputs; no UI change.
    - DB/data: no persisted-format claim is made.
    - Behavior test: position/owner/meaning matrix passes.
    - Cleanup/quarantine: no root/target fixture or catch-all helper.
    - Evidence IDs: `E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1`, `E1-P1B-REVIEW1`, `E1-P1B-DETECT1`, `E1-P1B-COMMIT1`.
    - Actual-status rows refreshed: range and identity-input rows.
  - Evidence Targets: impact, production diff, build, behavior/runtime, Supervisor, detect, commit.
  - Actual-status Update: range/identity inputs `partial -> correct`.
  - Commit Boundary: isolated implementation-slice commit.

- [x] P1-C: Implement declaration/symbol identity and occurrence conservation.
  - Goal: map every relevant source declaration to a deterministic occurrence identity and the correct lexical symbol without losing duplicate-name occurrences.
  - Scope Boundary:
    - Editable: only `graphIDForDef`/`buildWorkspace` and an exact projection/input owner proven by P1-C impact; the current candidate is `internal/resolution/indexes.go`.
    - Inspect-only: `emitDefinitionNodes`/`emitter.emitNode`, graph mutation/collision, persistence, readers, and later semantics unless the fresh impact proves a required plumbing edit.
    - Preserve-only: `BuildDefinitionIndex`, the TSJS binding-pattern branch, unrelated providers, relationship identity/merge, and unrelated graph facts.
    - Out of scope: collision API redesign beyond what this identity mapping needs.
  - Non-Goals: no heuristic merge by name and no relationship redesign without new impact evidence.
  - Pre-flight Questions:
    - Data source: accepted P1-B inputs and every current occurrence in `defsByFile`; `defsByID` is an index, not the conservation denominator.
    - Display permission: N/A — no UI permission behavior.
    - DB read flow: in-memory provider/ScopeIR-to-graph identity flow.
    - DB write flow: normal graph construction only; persistence method is not changed here.
    - Render location: normal analyze graph and identity evidence.
    - UI behavior flow: N/A — non-UI analyzer behavior.
    - Docker runtime: N/A — full build plus built analyzer boundary.
    - Playwright target: N/A — no browser surface.
    - Behavior test: same-name nested scopes, provider-supported meaning lanes, reorder/determinism, merge-evidence, occurrence conservation, and endpoint existence.
    - Cleanup/quarantine: package-local fixtures and repository-local evidence.
    - External side effects: none beyond normal analyze output.
    - N/A notes: reader parity belongs to Child 02.
  - Work Steps:
    1. Refresh graph; record impact for `graphIDForDef`/`buildWorkspace` and exact projection symbols; update the plan before touching any new producer family or the standalone definition index.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: enumerate the `defsByFile` input denominator, `DefinitionFact.ID` set, lexical-owner links, accepted graph-ID set, and affected endpoints.
       - Render location check: evidence ledger.
       - Mini QA: nearest identity-construction boundary.
       - Evidence target: `E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`.
    2. Implement identity/occurrence behavior, then tests; run full build and the bounded same-name oracle.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: input/output occurrence counts and identity sets match.
       - Render location check: built analyze output and oracle manifest.
       - Mini QA: normal built analyze on fixtures and bounded target sites.
       - Evidence target: `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`.
  - Implementation Gate: P1-B accepted; exact identity/occurrence owners and input denominator recorded.
  - Acceptance:
    - Source: deterministic declaration/symbol mapping uses source-proven lexical owner and meaning inputs; no name-only or range-only identity shortcut is accepted.
    - Runtime/UI: normal built analyze represents the expected identities; no UI change.
    - DB/data: occurrence conservation is 100% against the recorded `defsByFile`/source denominator at this boundary.
    - Behavior test: same-name, ordering, and merge-evidence matrix passes.
    - Cleanup/quarantine: no target-derived fixture or unused compatibility code.
    - Evidence IDs: `E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1C-REVIEW1`, `E1-P1C-DETECT1`, `E1-P1C-COMMIT1`.
    - Actual-status rows refreshed: declaration/symbol identity and occurrence conservation.
  - Evidence Targets: impact, source, build, conservation/oracle, Supervisor, detect, commit.
  - Actual-status Update: identity and occurrence loss `wrong -> correct` or record exact residual owner.
  - Commit Boundary: isolated implementation-slice commit.

- [x] P1-D: Remove silent collision or overwrite at the proven owner.
  - Goal: ensure a conflicting canonical fact cannot silently replace, skip, or collapse a distinct occurrence.
  - Scope Boundary:
    - Editable: only the canonical Definition emission conflict boundary in `internal/resolution/emit.go` and the minimum fail-clear propagation in `internal/resolution/resolve.go`, as proven by the P1-D caller inventory and exact impact.
    - Inspect-only: generic `Graph.AddNode`, `Graph.init`, all 11 production caller families, and their tests needed to distinguish legitimate enrichment/reinsertion from canonical Definition conflict.
    - Preserve-only: generic same-ID enrichment/update behavior, `AddRelationship`/`emitRelationship`, the ten non-resolution caller families, persistence/readers, and later semantics.
    - Out of scope: changing the global `AddNode` contract, rewriting producer families, fixing non-resolution producer identity, or adding relationship identity.
  - Non-Goals: no bulk API rewrite, warning-only behavior, or broad refactor.
  - Pre-flight Questions:
    - Data source: P1-C canonical identity output and the current `Graph.AddNode` caller/operation inventory.
    - Display permission: N/A — no UI permission behavior.
    - DB read flow: in-memory graph/index construction and decode paths proven affected.
    - DB write flow: normal graph construction only; no storage algorithm change.
    - Render location: clear analyze error and integrity evidence.
    - UI behavior flow: N/A — non-UI analyzer behavior.
    - Docker runtime: N/A — full build plus built CLI boundary.
    - Playwright target: N/A — no browser surface.
    - Behavior test: conflicting canonical identity, idempotent reinsertion if supported, legitimate enrichment, duplicate occurrence, missing endpoint, and all affected caller classes. The positive path runs through the built CLI; the negative collision path uses a compiled `ResolveBoundInto` malformed-ScopeIR probe because normal providers generate occurrence IDs from source position and cannot express that fault without an out-of-scope provider hack.
    - Cleanup/quarantine: package-local fault fixtures.
    - External side effects: none beyond isolated/normal analyze output.
    - N/A notes: implementation shape follows impact evidence.
  - Work Steps:
    1. Refresh graph; inventory the exact collision path and legitimate same-ID operations across all 11 production caller families; keep global mutation and unrelated producers preserve-only when source proves mixed insert/enrichment behavior.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: distinguish conflict from evidence-backed enrichment.
       - Render location check: evidence ledger.
       - Mini QA: graph-construction fault boundary.
       - Evidence target: `E1-P1D-IMPACT1`, `E1-P1D-SOURCE1`.
    2. Implement the smallest fail-clear correction at canonical Definition emission with exact-idempotent reinsertion and non-exact conflict rejection, then tests; run full build, a built-CLI positive fixture, and a compiled `ResolveBoundInto` collision/integrity fault probe.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: zero silent replacements/skips and zero missing affected endpoints.
       - Render location check: command failure/result evidence.
       - Mini QA: built CLI analyze on the positive identity fixture plus a compiled real resolution-package test binary for the malformed duplicate-ScopeIR conflict boundary; do not invent a provider/source collision fixture.
       - Evidence target: `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-COLLISION1`.
  - Implementation Gate: P1-C accepted; every editable owner and caller has fresh impact evidence.
  - Acceptance:
    - Source: silent collision/loss is removed only at canonical Definition emission; global `AddNode` enrichment/update and separate relationship replacement/merge remain unchanged.
    - Runtime/UI: conflicting input fails clearly or is handled by an explicit proven rule; no UI change.
    - DB/data: distinct occurrences remain conserved, conflict outcomes are explicit, and affected endpoints exist.
    - Behavior test: collision/enrichment/integrity matrix passes.
    - Cleanup/quarantine: obsolete defect-specific behavior and debug data removed.
    - Evidence IDs: `E1-P1D-IMPACT1`, `E1-P1D-SOURCE1`, `E1-P1D-BUILD1`, `E1-P1D-TEST1`, `E1-P1D-COLLISION1`, `E1-P1D-REVIEW1`, `E1-P1D-DETECT1`, `E1-P1D-COMMIT1`.
    - Actual-status rows refreshed: collision/mutation owner and integrity result.
  - Evidence Targets: impact/caller inventory, source, build, fault/integrity, Supervisor, detect, commit.
  - Actual-status Update: silent collision behavior `wrong -> correct`.
  - Commit Boundary: isolated implementation-slice commit.

- [x] P1-E: Integrate and validate identity accuracy.
  - Goal: prove the accepted identity correction through the normal built runtime and bounded target without adding new production semantics.
  - Scope Boundary:
    - Editable: validation fixtures/harnesses and only an exact earlier owner reopened by a failed gate after plan refresh.
    - Inspect-only: normal analyze flow, graph output, target boundary, and Child 02 handoff fields.
    - Preserve-only: the captured target source/worktree state, target graph baseline until the authorized in-place run, persistence implementation, readers, and later semantic owners.
    - Out of scope: production repair inside the validation slice without returning to P1-B/P1-C/P1-D.
  - Non-Goals: no persistence/reader change or target fixture copy.
  - Pre-flight Questions:
    - Data source: accepted P1-B/P1-C/P1-D output and bounded target source oracle.
    - Display permission: N/A — no UI permission behavior.
    - DB read flow: normal graph artifact and nearest graph query used for the oracle.
    - DB write flow: normal analyze output only; implementation mechanism remains source-owned.
    - Render location: CLI output, oracle manifest, ledgers, and Supervisor report.
    - UI behavior flow: N/A — CLI/analyzer acceptance.
    - Docker runtime: N/A — full repository build and built command boundary.
    - Playwright target: N/A — no browser surface.
    - Behavior test: `4/4`, occurrence conservation, collision failure, endpoint integrity, and repeated deterministic runs.
    - Cleanup/quarantine: remove debug-only traces; retain accepted evidence.
    - External side effects: normal target `.anvien` output only.
    - N/A notes: failures return to their owning implementation slice.
  - Work Steps:
    1. Run full build, then repeated matched normal analyze runs and compare accepted identity sets, ranges, occurrences, and endpoints.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: compare canonical facts across runs without assuming a write method.
       - Render location check: benchmark/evidence ledgers.
       - Mini QA: normal built CLI analyze/query boundary.
       - Evidence target: `E1-P1E-BUILD1`, `E1-P1E-DETERMINISM1`, `E1-P1E-INTEGRITY1`.
    2. Run the bounded target oracle in place only after the captured pre-state is compared, verify target preservation, obtain Supervisor PASS, detect changes, and commit the accepted slice.
       - UI flow check: N/A — non-UI.
       - DB/data flow check: `time`/`now` `4/4`, no affected occurrence loss.
       - Render location check: target oracle and Supervisor evidence.
       - Mini QA: built command against the real target boundary.
       - Evidence target: `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1`, `E1-P1E-REVIEW1`, `E1-P1E-DETECT1`, `E1-P1E-COMMIT1`.
  - Implementation Gate: P1-B through P1-D accepted and committed; target boundary captured before execution.
  - Acceptance:
    - Source: no new production semantic change is introduced by validation.
    - Runtime/UI: normal built commands pass all identity acceptance rows; no UI change.
    - DB/data: `4/4`, 100% occurrence conservation, zero proven collisions, zero affected missing endpoints, and no target pre-state loss.
    - Behavior test: deterministic matched-run and conflict matrices pass.
    - Cleanup/quarantine: target source/worktree preserved and debug artifacts removed.
    - Evidence IDs: `E1-P1E-BUILD1`, `E1-P1E-DETERMINISM1`, `E1-P1E-INTEGRITY1`, `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1`, `E1-P1E-REPORT1`, `E1-P1E-REVIEW1`, `E1-P1E-DETECT1`, `E1-P1E-COMMIT1`.
    - Actual-status rows refreshed: all Child 01 rows and Child 02 handoff facts.
  - Evidence Targets: build, deterministic/integrity/oracle/target boundary, Supervisor, detect, commit.
  - Actual-status Update: Child 01 identity scope `partial/wrong -> correct`; refresh Child 02 from accepted fields only.
  - Commit Boundary: isolated validation-slice commit.

- [x] Pn-A: Call Supervisor for the implemented-plan acceptance loop.
  - Goal: independently verify source, diff, runtime, reports, evidence, benchmark, and target boundary.
  - Work Steps:
    1. Run the Supervisor skill against the complete Child 01 result.
    2. Return each rejection only to its owning slice and repeat review after correction.
  - Implementation Gate: P1-A through P1-E are complete or explicitly blocked.
  - Acceptance: unconditional PASS recorded as `E2-PNA-SUPERVISOR1`, followed by report-only change detection and the isolated Pn-A report/ledger commit recorded as `E2-PNA-DETECT1` and `E2-PNA-COMMIT1`; or an exact blocker prevents closure.
  - Current State: accepted by the independent visible-task report `reports/Supervisor/rp_supervisor_260811_184507_by_gpt-5-codex_child01_pna.md`; the report returns one `PASS`, residual same-invariant surfaces `none`, and hands control back to the orchestration/main session. `E2-PNA-DETECT1` and `E2-PNA-COMMIT1` close the report-only boundary before Pn-B opens.

- [ ] Pn-B: Remove dead work created during this plan.
  - Goal: leave only accepted source, tests, fixtures, and evidence.
  - Work Steps:
    1. Identify superseded debug, failed, duplicate, and rejected artifacts created by Child 01.
    2. Remove only those artifacts and obtain Supervisor cleanup review.
  - Implementation Gate: Pn-A has reviewed the implementation result.
  - Acceptance: cleanup proof and Supervisor result recorded as `E2-PNB-CLEANUP1`.
  - Current State: open only after Git confirms the isolated Pn-A report/ledger commit; inspect and classify Child 01 artifacts without deleting historical or unrelated work.

- [ ] Pn-C: Close the plan and hand off accepted fields.
  - Goal: finish final validation, ledgers, detect-changes, commits, and the Child 02 status refresh.
  - Work Steps:
    1. Run final required validation and record the accepted benchmark/evidence values.
    2. Run detect-changes, record final commit/worktree state, and refresh Child 02 actual status using exact accepted Child 01 facts.
  - Implementation Gate: Pn-A and Pn-B pass; every slice evidence row is complete.
  - Acceptance: `E2-PNC-DETECT1`, `E2-PNC-HANDOFF1`, and `E2-PNC-COMMIT1` are recorded; Child 02 opens only from that evidence.

## Risk Notes

- Identity construction has HIGH/CRITICAL blast radius; exact impact controls scope but does not forbid the needed edit.
- Fresh P0 source-flow evidence shows `BuildDefinitionIndex` is test-only and must not be edited as a proxy for the production occurrence loss.
- Adding a line or range to the old identity can hide the bounded collision while leaving lexical-owner or meaning collisions intact.
- Tightening a generic mutation method without classifying legitimate enrichment can break unrelated producers.
- `Graph.AddNode` has 11 production caller families and CRITICAL upstream impact; the P1-D conflict policy must be scoped to canonical source nodes.
- Prescribing all producers or relationship identities before impact evidence would expand the defect beyond its proven boundary.
- Range encoding can drift across providers and editor-facing adapters; P1-B must state what current source actually supplies.
- Determinism claims require matched inputs and explicit canonical comparisons, not graph counts alone.
- The target has pre-existing state and a separate scanner defect; neither may be folded into Child 01 repair.
