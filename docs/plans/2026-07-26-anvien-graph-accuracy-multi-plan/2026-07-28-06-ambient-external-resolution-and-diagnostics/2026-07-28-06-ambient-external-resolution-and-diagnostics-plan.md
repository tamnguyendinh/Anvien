# Anvien Ambient And External Resolution Plan

## Metadata

- Date: `2026-07-28`
- Status: `P0 complete / dependency-blocked`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Persistence dependency: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
- Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`

## Goal

Make TypeScript standard-library and other evidenced external declarations resolve truthfully on the current production resolver, so `Promise`, `Math.max`, and `Math.min` are external results or explicit external-capability outcomes rather than false in-repository analyzer gaps.

## Rules

- Complete and refresh P0 actual status before implementation work.
- The originating problem report supplies defect findings and target acceptance. Its proposed declaration architecture is not implementation authority.
- The causal synthesis and final Supervisor report verify the bounded missing-declaration-universe cause only; broader TypeScript declaration semantics remain unresolved until P6-A evidence closes them.
- Work directly on the current production source, build, command, and graph path.
- Existing resolver/analyze behavior is baseline evidence, not an immutable design. Change only the part whose necessity is established by the active slice.
- P6-A must inspect current source/config/runtime and choose the minimum declaration-authority behavior before P6-B selects an implementation mechanism.
- P6-B implements TypeScript standard-library authority according to the accepted P6-A decision; its mechanism is selected only by that decision and current source/runtime evidence.
- P6-C1 handles project/package declaration lookup only if P6-A evidence establishes that it belongs to this campaign's accepted contract.
- P6-C2 adds external target representation or materialization only to the extent required by accepted resolver/persistence consumers.
- P6-C3 owns structured resolver outcomes. P6-D only projects those outcomes into graph-health; it must not infer resolution from target text.
- Child 05 results are immutable predecessor inputs. Child 06 does not repeat module/export traversal.
- Child 02 owns persistence/reader preservation. Child 06 changes only consumers identified by its current affected-surface inventory and fresh impact evidence.
- Before graph-based work, run `anvien analyze --force`. Before editing a resolver, graph-health function, provider, shared contract, or projection, run file-detail and upstream impact and report the full blast radius.
- HIGH or CRITICAL impact is a warning requiring narrow edits and affected-boundary regression, not an edit prohibition.
- Implement production behavior before adding or updating tests. Run the full repository build before boundary validation.
- Every completed production slice requires focused behavior evidence, affected regression, ledger refresh, Supervisor PASS, `anvien detect-changes --repo E:\Anvien --scope all`, and an isolated commit. P6-A is a decision/document slice and records its exact worktree/commit evidence without claiming implementation change detection.
- Each touched production/test file owns one primary semantic responsibility. Do not add catch-all files or unrelated logic.
- Do not introduce network access, package installation, package-script execution, or a new production runtime dependency implicitly. If P6-A evidence makes one necessary, update scope, risks, validation, and authority before implementation.
- Analyze `E:\cheapapp.org` in place only in P6-D target validation. Keep target source unchanged and all plan/QA/report artifacts in `E:\Anvien`.
- Update checklists, evidence, benchmark, and actual status immediately when state changes.

## Problem

The originating report identifies three bounded false gaps: TypeScript resolves `Promise`, `Math.max`, and `Math.min`, while Anvien does not. The accepted causal synthesis locates the first confirmed divergence in the resolver workspace: it indexes only definitions extracted from target repository ScopeIR and loads no TypeScript ambient/standard-library declaration authority.

Current source matches that finding. `buildWorkspace` builds indexes only from the supplied repository files. `resolveTypeAnnotation` ignores a small primitive list and otherwise queries that workspace. Member calls depend on a receiver type and an owner/member found in the same indexes. When resolution fails, graph-health later classifies unresolved diagnostics from target text using Go-oriented builtin, standard-library, and qualifier tables.

The evidence proves the missing authority and incorrect downstream inference. It does not prove that one particular authority mechanism, seven fixed declaration sources, or a new public traversal option is required. P6-A must resolve those design questions from current source/config/runtime evidence before code.

## Scope

In scope:

- a source-backed declaration-universe behavior contract for the current analyzer;
- the minimum TypeScript standard-library authority needed to resolve configured language declarations accurately;
- project-owned or installed-package declaration lookup only when P6-A establishes its requirement and owner boundary;
- external target identity/representation/materialization required by actual graph and reader consumers;
- one structured resolution outcome per affected source site, including explicit capability/unavailable results;
- graph-health projection from the resolver's outcome;
- exact target validation for `Promise`, `Math.max`, and `Math.min`;
- affected persistence/readers identified by Child 02 inventory and fresh P6 impact.

Out of scope:

- a preselected declaration storage/authority mechanism;
- a mandatory inventory of declaration-source categories not established by P6-A;
- a new public external-traversal command/API without consumer and product evidence;
- command/persistence protocol work unrelated to the corrected facts;
- hardcoded target names, target-text reclassification, or graph-health-only repair;
- graphing all installed declarations or all dependency files without evidence;
- module/export traversal owned by Child 05;
- scanner behavior, target-source edits, and unrelated language semantics.

## Non-Goals

- Do not treat `Promise`, `Math`, or the three target sites as a name allowlist.
- Do not label an unresolved site as external merely to improve graph-health output.
- Do not decide external representation before inspecting actual graph/persistence/reader needs.
- Do not claim full TypeScript compiler conformance from three bounded target sites.
- Do not expand project/package lookup beyond the accepted P6-A contract.
- Do not add broad context/impact/rename/process/group behavior unless fresh impact and product evidence require those surfaces.

## Requirements

### P6-A declaration-universe decision

- Record the current declaration inputs, TypeScript configuration inputs actually read today, resolver lookup stages, failure shape, affected graph facts, and affected readers before selecting a design.
- Use the accepted TypeScript differential as the bounded oracle and add independently authored fixtures that test behavior rather than target names.
- Define which declaration authority is required for standard-library lookup and whether project-owned/package declarations are required in this campaign.
- Define config-sensitive behavior only for configuration inputs the product will actually support; unsupported or unavailable authority must produce an explicit outcome.
- Compare feasible mechanisms against current packaging/runtime, determinism, provenance, security, performance, and maintenance evidence. Record why the selected mechanism is necessary and why rejected mechanisms do not satisfy the contract.
- Identify exact production/test owners with file-detail and impact. P6-A may update later slice steps before code; it does not pre-authorize a file map.

### Standard-library and declaration resolution

- P6-B implements the accepted P6-A standard-library authority without hardcoded target names.
- The authority must distinguish type/value/namespace/member meaning needed by accepted resolver inputs.
- `Promise`, `Math`, `max`, and `min` behavior must arise from declaration data/semantics selected by P6-A and general fixtures, not special branches.
- Configuration or authority mismatch/unavailability is visible as an external-capability outcome and never becomes `in_repo_unresolved`/`analyzer_gap`.
- P6-C1 implements project/package declaration lookup only if the P6-A decision marks it required. Otherwise it records a source-backed preserve-only result and does not invent code.
- Any path/security/resource limits introduced by P6-C1 are derived from measured inputs and explicit policy, not arbitrary copied ceilings.

### External targets and structured outcomes

- P6-C2 selects the minimum representation needed to distinguish internal Symbols, external declaration targets, language intrinsics, and unresolved external capability. It records actual consumer requirements before choosing nodes, references, or another representation.
- If external targets are materialized, they retain their external origin/provenance and do not masquerade as repository-owned File/Definition facts.
- A source site has one final structured outcome with requested name/meaning, resolution stage, target when resolved, explicit reason when unresolved, candidates/proof as applicable, and authority/config evidence required by P6-A.
- An earlier authoritative failure is not overwritten by a later name guess.
- P6-C3 defines only statuses required by current P5/P6 cases and accepted fixtures. New cases expand the status contract through a plan update before code.
- P6-D maps graph-health classification/actionability mechanically from the structured outcome. It removes same-invariant target-text inference at every affected site identified by fresh impact.

### Persistence and target validation

- Persist/project only fields proven necessary by affected graph/readers. JSON, Ladybug, and each affected reader must agree on the target/outcome facts with zero loss.
- The three target sites pass only when each is either correctly resolved to an external declaration target or carries the explicitly accepted external-capability outcome for the real configured environment.
- None of the three may remain `in_repo_unresolved` or `analyzer_gap`.
- Target validation records exact source sites, outcomes, proofs, graph records, and pre/post target boundary. Aggregate graph-health counts alone are insufficient.

## Acceptance Criteria

- P6-A records a Supervisor-accepted, source-backed declaration-universe decision before any P6-B production edit.
- P6-B resolves general TypeScript standard-library fixtures according to that decision without target-name hardcoding.
- P6-C1 either implements the exact evidenced project/package declaration behavior or records why it is not required; no extra lookup architecture is added.
- P6-C2 represents only external targets/outcomes required by actual consumers and preserves external provenance.
- P6-C3 produces one structured outcome per affected site with no resolved/unresolved overlap and no later name-based override.
- Graph-health consumes resolver outcomes and does not independently guess the three target classifications.
- `Promise`, `Math.max`, and `Math.min` reach correct external or accepted external-capability outcomes (`3/3`), with `0/3` remaining in-repository analyzer gaps.
- No target-name allowlist or copied target fixture is present.
- Only affected persistence/readers identified by Child 02 and fresh impact are changed; their outcomes match the graph source.
- Full build, nearest real resolver/graph-health boundary, target boundary, regression, Supervisor, detect-changes, and per-slice commit evidence pass.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish current declaration inputs, workspace lookup, failure/classification behavior, target baseline, and blast radius.
  - Work Steps:
    1. Read the problem origin, bounded verification reports, current workspace/resolver/graph-health source, and all four Child 06 ledgers.
    2. Refresh graph evidence; run file-detail and impact for current resolver and diagnostic owners; classify proven facts separately from design proposals.
    3. Rewrite all later slices around evidence-selected decisions and record the Child 05 dependency.
  - Implementation Gate: no production edit starts until actual status has a final P0 decision and Child 05 closure is accepted.
  - Acceptance: actual status distinguishes missing standard-library authority, conditional project/package scope, undecided external representation, missing structured outcomes, wrong graph-health inference, and exact next actions.
  - Evidence Targets: `E0-P0A-RULE1`, `E0-P0A-SKILL1`, `E0-P0A-ORIGIN1`, `E0-P0A-VERIFY1`, `E0-P0A-VERIFY2`, `E0-P0A-GRAPH1`, `E0-P0A-SRC1..E0-P0A-SRC4`, `E0-P0A-FD1..E0-P0A-FD3`, `E0-P0A-IMPACT1..E0-P0A-IMPACT4`, `E0-P0A-DEPEND1`, `E0-P0A-SCOPE1`, `E0-P0A-SCOPE2`, `E0-P0A-TARGET1`, `E0-P0A-BOUNDARY1`, and `E0-P0A-STATUS1`.

### P6: Ambient and external resolution

- Phase Goal: resolve evidenced external declarations and project truthful outcomes into the graph and graph-health.
- Phase Boundary:
  - In scope: P6-A through P6-D only.
  - Out of scope: Child 05 module exports, broad public API changes, and unevidenced declaration sources.
  - Dependencies: Child 05 closure; Child 02 affected-reader inventory when persisted fields change.
- Phase Implementation Rule: execute, validate, review, refresh, and commit one slice before opening the next.
- Ordered Slice List:
  - P6-A: Establish the declaration-universe behavior and design from evidence.
  - P6-B: Implement the accepted TypeScript standard-library authority.
  - P6-C1: Implement evidenced project/package declaration lookup.
  - P6-C2: Add the necessary external target representation.
  - P6-C3: Produce structured resolution outcomes.
  - P6-D: Project outcomes and prove the target sites.

- [ ] P6-A: Establish the declaration-universe behavior and design from evidence.
  - Goal: decide the minimum production declaration authority, supported config behavior, failure contract, owner boundary, and later-slice scope from current evidence.
  - Scope Boundary:
    - Editable: Child 06 plan/contract/ledger decisions only; production source remains inspect-only in this slice.
    - Inspect-only: current workspace, resolver, TypeScript provider/config readers, packaging/runtime, persistence/readers, and Child 05 result contract.
    - Preserve-only: current production behavior until the decision is accepted.
    - Out of scope: implementation of the selected authority.
  - Non-Goals: no mechanism selection from document precedent or terminology alone.
  - Pre-flight Questions:
    - Data source: current source/config/runtime, accepted Child 05 outcomes, target source sites, and TypeScript oracle.
    - Display permission: N/A — contract decision has no display permission.
    - DB read flow: inspect graph/persistence consumers that require external facts.
    - DB write flow: N/A — no production data write.
    - Render location: plan, evidence, actual-status, and decision record.
    - UI behavior flow: N/A — no UI change.
    - Docker runtime: N/A — use full baseline build and current resolver boundary.
    - Playwright target: N/A — no UI behavior.
    - Behavior test: differential fixtures for standard-library, unavailable authority, config variation, and language isolation.
    - Cleanup/quarantine: keep only reusable independently authored oracle fixtures/evidence.
    - External side effects: none; all mechanism side effects are evaluated, not performed.
    - N/A notes: P6-A is the mandatory evidence/decision slice before production implementation.
  - Work Steps:
    1. Inventory current inputs/stages/consumers and run fresh file-detail/impact plus the bounded/general TypeScript differential; record supported and unknown declaration cases.
       - UI flow check: N/A — non-UI decision slice.
       - DB/data flow check: trace each target site from source fact to current gap/health output.
       - Render location check: evidence and actual-status ledgers.
       - Mini QA: exercise the built current resolver/graph-health boundary and record actual output.
       - Evidence target: `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-ORACLE1`, `E6-P6A-CONSUMER1`.
    2. Compare feasible authority mechanisms; record the accepted behavior, chosen mechanism, owner map, conditional P6-C1/P6-C2 scope, measurements, and updated later-slice steps; run baseline full build and Supervisor review, then commit the decision slice.
       - UI flow check: N/A — no UI change.
       - DB/data flow check: required facts/outcomes and affected persistence fields are explicit.
       - Render location check: contract/plan/evidence ledgers.
       - Mini QA: confirm the decision matches observed current/runtime constraints.
       - Evidence target: `E6-P6A-DECISION1`, `E6-P6A-BUILD1`, `E6-P6A-REVIEW1`, `E6-P6A-COMMIT1`.
  - Implementation Gate: exact Child 05 handoff is accepted; graph/source/runtime evidence is fresh; no production edit occurs in P6-A.
  - Acceptance:
    - Source: current declaration inputs, stages, consumers, and missing authority are fully traced.
    - Runtime/UI: baseline built resolver/oracle behavior is recorded; UI is N/A.
    - DB/data: necessary graph/persistence outcome fields and affected consumers are identified.
    - Behavior test: standard-library/unavailable/config/language-isolation oracle cases are specified.
    - Cleanup/quarantine: no rejected mechanism artifact or copied target fixture remains.
    - Evidence IDs: `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-ORACLE1`, `E6-P6A-CONSUMER1`, `E6-P6A-DECISION1`, `E6-P6A-BUILD1`, `E6-P6A-REVIEW1`, `E6-P6A-COMMIT1`.
    - Actual-status rows refreshed: authority mechanism, supported config, P6-C1 scope, P6-C2 representation, and affected-reader rows.
  - Evidence Targets: source/impact/oracle/consumer inventory, decision comparison, updated plan, build, Supervisor, commit.
  - Actual-status Update: declaration-authority design `blocked -> correct`; implementation rows remain pending.
  - Commit Boundary: commit P6-A decision/ledger scope alone after acceptance.

- [ ] P6-B: Implement the accepted TypeScript standard-library authority.
  - Goal: provide deterministic standard-library declaration lookup through the mechanism and config behavior accepted in P6-A.
  - Scope Boundary:
    - Editable: exact authority/config/lookup owners selected by P6-A and refreshed impact.
    - Inspect-only: resolver finalization, external representation, graph-health, and persistence.
    - Preserve-only: other languages and Child 05 module/export resolution.
    - Out of scope: project/package declarations unless P6-A explicitly assigns them here.
  - Non-Goals: no `Promise`/`Math` special case and no design substitution after P6-A without a plan update.
  - Pre-flight Questions:
    - Data source: P6-A accepted authority inputs and supported TypeScript config.
    - Display permission: N/A — resolver data.
    - DB read flow: read only declaration/config inputs authorized by P6-A.
    - DB write flow: write only authority indexes/cache/artifacts required by the accepted design.
    - Render location: resolver trace and benchmark/evidence.
    - UI behavior flow: N/A — no UI behavior.
    - Docker runtime: conditional only if packaged/container runtime is affected; otherwise built CLI/resolver boundary.
    - Playwright target: N/A unless P6-A identifies an affected UI surface.
    - Behavior test: general standard-library types/namespaces/members, `Promise`, `Math.max`, `Math.min`, config variations, unavailable/mismatch, determinism, and language isolation.
    - Cleanup/quarantine: accepted reusable fixtures/assets only; remove rejected output.
    - External side effects: exactly those accepted in P6-A and recorded before implementation.
    - N/A notes: project/package lookup and final outcomes are later slices unless P6-A explicitly reassigns them.
  - Work Steps:
    1. Refresh exact owner impact; implement the P6-A authority and lookup production behavior without target-name branches.
       - UI flow check: N/A unless affected UI was proved.
       - DB/data flow check: declaration provenance/config and lookup meaning are retained as required by P6-A.
       - Render location check: resolver trace/evidence.
       - Mini QA: exercise the real built authority lookup boundary.
       - Evidence target: `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-PROVENANCE1`.
    2. Add behavior tests after code; run full build, packaging/runtime checks required by the design, determinism/performance measurements, regressions, review, detect, and commit.
       - UI flow check: conditional on P6-A affected surfaces.
       - DB/data flow check: exact general fixture targets and explicit unavailable behavior.
       - Render location check: evidence/benchmark ledgers.
       - Mini QA: inspect real built lookup results rather than helper output.
       - Evidence target: `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-REVIEW1`, `E6-P6B-DETECT1`, `E6-P6B-COMMIT1`.
  - Implementation Gate: P6-A decision/Supervisor/commit accepted; exact owners, inputs, side effects, and measurements are fixed.
  - Acceptance:
    - Source: production authority matches P6-A design and contains no target-name allowlist.
    - Runtime/UI: built lookup behaves correctly for general fixtures; UI is conditional.
    - DB/data: declaration target/provenance/config facts required by the decision are exact.
    - Behavior test: standard-library, config, unavailable, determinism, and language-isolation cases pass.
    - Cleanup/quarantine: rejected outputs and debug artifacts removed; accepted assets are reproducible if applicable.
    - Evidence IDs: `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-PROVENANCE1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-REVIEW1`, `E6-P6B-DETECT1`, `E6-P6B-COMMIT1`.
    - Actual-status rows refreshed: TypeScript standard-library authority and config behavior.
  - Evidence Targets: impact/source/provenance, tests/build/runtime, benchmarks, regression, Supervisor, detect, commit.
  - Actual-status Update: standard-library authority `missing -> correct`.
  - Commit Boundary: commit P6-B alone after acceptance.

- [ ] P6-C1: Implement evidenced project/package declaration lookup.
  - Goal: implement only the project/package declaration behavior that P6-A marks required, or close preserve-only with proof that this campaign does not require it.
  - Scope Boundary:
    - Editable: exact declaration-lookup owners selected by P6-A when the slice is active.
    - Inspect-only: P6-B authority, Child 05 module result, repository config, and actual package/project declaration inputs.
    - Preserve-only: package/module selection already owned by Child 05.
    - Out of scope: broad dependency scanning, package execution, and unapproved declaration sources.
  - Non-Goals: no duplicate module/export resolver and no scope expansion from filenames alone.
  - Pre-flight Questions:
    - Data source: P6-A decision, Child 05 module result, and authorized current project/package inputs.
    - Display permission: N/A — resolver input.
    - DB read flow: only sources authorized by P6-A.
    - DB write flow: lookup indexes/cache only if accepted design requires them.
    - Render location: lookup trace and evidence.
    - UI behavior flow: N/A — no UI behavior.
    - Docker runtime: conditional on selected packaging behavior.
    - Playwright target: N/A unless an affected UI was proved.
    - Behavior test: exact required present/missing/config/security cases or preserve-only proof.
    - Cleanup/quarantine: independent fixtures; no copied dependency tree.
    - External side effects: only P6-A-approved local reads/actions.
    - N/A notes: close without production edit when P6-A proves this behavior is not required.
  - Work Steps:
    1. Reconfirm the P6-A decision and fresh impact; record the exact required source/lookup boundary or the evidence-backed preserve-only result.
       - UI flow check: N/A — non-UI lookup.
       - DB/data flow check: module result and declaration lookup remain separate.
       - Render location check: evidence/actual-status.
       - Mini QA: exercise the real current lookup boundary for required cases.
       - Evidence target: `E6-P6C1-IMPACT1`, `E6-P6C1-SCOPE1`.
    2. When active, implement production lookup then tests, full build, real boundary, measurements, review, detect, and commit; otherwise validate preserve-only behavior and commit the ledger decision.
       - UI flow check: conditional on affected UI evidence.
       - DB/data flow check: exact declaration target or explicit external-capability outcome.
       - Render location check: resolver trace/evidence.
       - Mini QA: inspect built required-case results.
       - Evidence target: `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-TEST1`, `E6-P6C1-REVIEW1`, `E6-P6C1-DETECT1`, `E6-P6C1-COMMIT1`.
  - Implementation Gate: P6-A marks the slice active or preserve-only; exact owners/sources/side effects are recorded.
  - Acceptance:
    - Source: only P6-A-required behavior is implemented; otherwise no production diff.
    - Runtime/UI: required real lookup cases pass; UI is conditional.
    - DB/data: project/package result remains distinct from module/export selection and retains required proof.
    - Behavior test: accepted present/missing/config/security cases or preserve-only proof passes.
    - Cleanup/quarantine: no copied dependency tree or obsolete fixture remains.
    - Evidence IDs: `E6-P6C1-IMPACT1`, `E6-P6C1-SCOPE1`, `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-TEST1`, `E6-P6C1-REVIEW1`, `E6-P6C1-DETECT1`, `E6-P6C1-COMMIT1`.
    - Actual-status rows refreshed: project/package declaration scope and implementation state.
  - Evidence Targets: scope decision, impact, source/preserve proof, tests/build/runtime, Supervisor, detect, commit.
  - Actual-status Update: conditional lookup `blocked -> correct` or `blocked -> preserve-only`.
  - Commit Boundary: commit P6-C1 alone after acceptance.

- [ ] P6-C2: Add the necessary external target representation.
  - Goal: represent resolved external declarations and explicit external-capability gaps in the minimum truthful form required by actual consumers.
  - Scope Boundary:
    - Editable: exact representation/materialization and affected persistence owners selected by P6-A/current impact.
    - Inspect-only: P6-B/P6-C1 lookup results, graph types, Child 02 affected readers, and final outcome owner.
    - Preserve-only: repository File/Definition identity and unrelated graph consumers.
    - Out of scope: graph-health policy and broad traversal API changes.
  - Non-Goals: no external target disguised as repository-owned data and no consumer expansion without evidence.
  - Pre-flight Questions:
    - Data source: accepted declaration lookup results and P6-A consumer inventory.
    - Display permission: preserve current permissions; no new display surface unless impact requires it.
    - DB read flow: inspect graph/persistence consumers requiring the representation.
    - DB write flow: only accepted target/provenance facts through affected persistence owners.
    - Render location: graph and affected readers.
    - UI behavior flow: conditional on affected reader inventory.
    - Docker runtime: conditional on an affected app/UI surface.
    - Playwright target: conditional on an affected browser surface.
    - Behavior test: resolved external target, unavailable capability, internal/external separation, duplicate/provenance, persistence parity, and no repository ownership pollution.
    - Cleanup/quarantine: remove failed/incomplete plan-created output.
    - External side effects: none beyond accepted P6-A design.
    - N/A notes: public traversal controls are not part of this slice without separate evidence.
  - Work Steps:
    1. Record graph/reader impact and select the minimum representation; update the plan before code if consumer requirements differ from P6-A.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: representation preserves external origin and explicit unresolved capability.
       - Render location check: graph/affected-reader evidence.
       - Mini QA: inspect real graph and nearest affected reader.
       - Evidence target: `E6-P6C2-IMPACT1`, `E6-P6C2-DESIGN1`.
    2. Implement production representation/materialization, then tests, full build, affected parity, regression, review, detect, and commit.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: exact target/provenance and no repository File/Definition misclassification.
       - Render location check: graph and affected readers only.
       - Mini QA: exercise real built graph-bound consumer(s).
       - Evidence target: `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-PARITY1`, `E6-P6C2-REVIEW1`, `E6-P6C2-DETECT1`, `E6-P6C2-COMMIT1`.
  - Implementation Gate: P6-B and active P6-C1 behavior accepted; exact consumer/representation need and affected readers are proved.
  - Acceptance:
    - Source: minimum accepted representation distinguishes internal, external, intrinsic when applicable, and explicit external-capability gaps.
    - Runtime/UI: affected built consumers show the same truthful result; UI is conditional.
    - DB/data: external origin/provenance survives affected persistence with zero repository ownership pollution.
    - Behavior test: resolved/unavailable/separation/duplicate/parity cases pass.
    - Cleanup/quarantine: no incomplete or superseded representation artifact remains.
    - Evidence IDs: `E6-P6C2-IMPACT1`, `E6-P6C2-DESIGN1`, `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-PARITY1`, `E6-P6C2-REVIEW1`, `E6-P6C2-DETECT1`, `E6-P6C2-COMMIT1`.
    - Actual-status rows refreshed: external target representation and affected readers.
  - Evidence Targets: impact/design/source, tests/build/parity/runtime, Supervisor, detect, commit.
  - Actual-status Update: external representation `blocked/missing -> correct`.
  - Commit Boundary: commit P6-C2 alone after acceptance.

- [ ] P6-C3: Produce structured resolution outcomes.
  - Goal: finalize exactly one immutable resolver outcome per affected source site from repository, external, intrinsic, or explicit external-capability results.
  - Scope Boundary:
    - Editable: exact outcome/status owner and minimum finalization adapters selected by impact.
    - Inspect-only: P5/P6 lookup results, graph representation, and graph-health.
    - Preserve-only: declaration lookup decisions and unrelated diagnostics.
    - Out of scope: declaration loading and graph-health rendering.
  - Non-Goals: no exhaustive status set unrelated to accepted current cases and no target-text inference.
  - Pre-flight Questions:
    - Data source: immutable P5 repository results and P6-B/C results.
    - Display permission: N/A — resolver outcome.
    - DB read flow: read accepted stage results only.
    - DB write flow: write one final outcome and its required persisted fields.
    - Render location: resolver trace and affected persistence/readers.
    - UI behavior flow: N/A unless affected-reader inventory says otherwise.
    - Docker runtime: conditional on affected app/UI surface.
    - Playwright target: conditional on affected browser surface.
    - Behavior test: every accepted current status, stage precedence, one-result exclusivity, proof/candidates, and unavailable authority.
    - Cleanup/quarantine: isolated independently authored outcome fixtures.
    - External side effects: none.
    - N/A notes: graph-health mapping is P6-D.
  - Work Steps:
    1. Record finalizer/consumer impact; define and implement only statuses/fields required by accepted P5/P6 cases and fixtures.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: exactly one outcome per source site; no later stage overwrites an authoritative failure.
       - Render location check: resolver trace/affected persistence.
       - Mini QA: inspect real built outcome output for representative cases.
       - Evidence target: `E6-P6C3-IMPACT1`, `E6-P6C3-SRC1`, `E6-P6C3-STATUS1`.
    2. Add tests after code, run full build and affected parity/regression, prove no resolved/unresolved overlap, then review/detect/commit.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: status/stage/target/reason/proof survive affected boundaries.
       - Render location check: evidence/benchmark ledgers.
       - Mini QA: exercise nearest real resolver/graph-bound boundary.
       - Evidence target: `E6-P6C3-BUILD1`, `E6-P6C3-TEST1`, `E6-P6C3-PARITY1`, `E6-P6C3-REVIEW1`, `E6-P6C3-DETECT1`, `E6-P6C3-COMMIT1`.
  - Implementation Gate: P6-B/C results accepted; exact status cases and affected consumers are inventoried.
  - Acceptance:
    - Source: one structured immutable outcome per affected site; no target-text guess or later override.
    - Runtime/UI: built resolver outputs exact accepted outcomes; UI is conditional.
    - DB/data: required fields survive affected persistence/readers with zero resolved/unresolved overlap.
    - Behavior test: all accepted status and precedence cases pass.
    - Cleanup/quarantine: no unused status or obsolete fixture remains.
    - Evidence IDs: `E6-P6C3-IMPACT1`, `E6-P6C3-SRC1`, `E6-P6C3-STATUS1`, `E6-P6C3-BUILD1`, `E6-P6C3-TEST1`, `E6-P6C3-PARITY1`, `E6-P6C3-REVIEW1`, `E6-P6C3-DETECT1`, `E6-P6C3-COMMIT1`.
    - Actual-status rows refreshed: structured outcome and affected persistence rows.
  - Evidence Targets: impact/source/status, tests/build/parity/regression, Supervisor, detect, commit.
  - Actual-status Update: structured outcomes `missing -> correct`; graph-health remains pending.
  - Commit Boundary: commit P6-C3 alone after acceptance.

- [ ] P6-D: Project outcomes and prove the target sites.
  - Goal: make graph-health a faithful projection of P6-C3 outcomes and validate the three bounded target sites on the built current runtime.
  - Scope Boundary:
    - Editable: exact graph-health projection and affected output adapters identified by fresh impact.
    - Inspect-only: resolver outcomes, graph/persistence, target source/worktree, and independent oracle.
    - Preserve-only: unrelated health policies and target source.
    - Out of scope: resolver/declaration behavior already accepted in earlier slices.
  - Non-Goals: no `Promise`/`Math` branch and no reclassification without resolver evidence.
  - Pre-flight Questions:
    - Data source: final structured outcomes and exact target source sites.
    - Display permission: preserve current permissions; include UI only if impact proves it affected.
    - DB read flow: read graph/Ladybug and only affected graph-health/readers.
    - DB write flow: normal analyze writes current graph output; this slice observes that existing boundary.
    - Render location: graph-health CLI/MCP/HTTP/Web only where affected.
    - UI behavior flow: conditional on affected UI evidence.
    - Docker runtime: conditional on affected app/UI; otherwise full build plus real graph-health command.
    - Playwright target: conditional on affected browser surface.
    - Behavior test: mechanical status mapping, no heuristic override, exact three target outcomes, zero in-repository analyzer gaps, parity, and target boundary.
    - Cleanup/quarantine: reports in Anvien; no target report/fixture/source edit.
    - External side effects: normal target analyze may regenerate target `.anvien`; source remains unchanged.
    - N/A notes: no broad output-surface matrix beyond actual impact.
  - Work Steps:
    1. Inventory every same-invariant heuristic/adapter from fresh impact; implement mechanical outcome projection, then tests, full build, and affected-reader regression.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: graph-health classification/actionability derives from outcome fields.
       - Render location check: graph-health and affected output adapters.
       - Mini QA: exercise the nearest real built graph-health boundary; inspect UI only if affected.
       - Evidence target: `E6-P6D-IMPACT1`, `E6-P6D-SRC1`, `E6-P6D-BUILD1`, `E6-P6D-TEST1`, `E6-P6D-PARITY1`.
    2. Capture target pre-state; analyze `E:\cheapapp.org`; compare the independent TypeScript oracle; prove `Promise`, `Math.max`, and `Math.min` are correct external/capability outcomes with no in-repository analyzer gap; verify target boundary, then review/detect/commit.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: exact three outcome/proof records are the acceptance source.
       - Render location check: official evidence under Anvien.
       - Mini QA: run real target query/graph-health/file-detail and visually inspect any affected UI result.
       - Evidence target: `E6-P6D-TARGET1`, `E6-P6D-ORACLE1`, `E6-P6D-BOUNDARY1`, `E6-P6D-REVIEW1`, `E6-P6D-DETECT1`, `E6-P6D-COMMIT1`.
  - Implementation Gate: P6-A/B/C accepted; affected graph-health/readers and target baseline are recorded; full build passes.
  - Acceptance:
    - Source: graph-health maps resolver outcomes and contains no same-invariant target-text guess.
    - Runtime/UI: all three target sites show correct external or accepted capability outcomes; UI is conditional.
    - DB/data: exact target outcomes/proofs survive graph and affected readers; no site is both resolved and a gap.
    - Behavior test: mechanical mapping, no-heuristic, three target sites, and affected regression pass.
    - Cleanup/quarantine: target source/worktree unchanged; obsolete debug evidence removed.
    - Evidence IDs: `E6-P6D-IMPACT1`, `E6-P6D-SRC1`, `E6-P6D-BUILD1`, `E6-P6D-TEST1`, `E6-P6D-PARITY1`, `E6-P6D-TARGET1`, `E6-P6D-ORACLE1`, `E6-P6D-BOUNDARY1`, `E6-P6D-REVIEW1`, `E6-P6D-DETECT1`, `E6-P6D-COMMIT1`.
    - Actual-status rows refreshed: graph-health projection, target acceptance, and affected readers.
  - Evidence Targets: impact/source/build/tests/parity, target oracle/boundary, Supervisor, detect, commit.
  - Actual-status Update: graph-health `wrong -> correct`; bounded target `0/3 -> 3/3`.
  - Commit Boundary: commit P6-D alone after acceptance.

- [ ] Pn-A: Call Supervisor for Child 06 acceptance.
  - Goal: independently verify the six slices, design decision, source diff, runtime/target evidence, ledgers, and commits.
  - Work Steps:
    1. Run Supervisor review over the complete Child 06 scope.
    2. Return only rejected invariants to their owning slice and repeat until PASS or a precise blocker.
  - Implementation Gate: P6-A through P6-D are accepted locally or explicitly blocked with evidence.
  - Acceptance: `E6-PNA-REVIEW1` records Supervisor PASS or a precise blocker.

- [ ] Pn-B: Remove dead work created by Child 06.
  - Goal: retain only accepted source, tests, fixtures/assets, evidence, and ledgers.
  - Work Steps:
    1. Inventory failed, duplicate, superseded, or unused Child 06 artifacts and decisions.
    2. Remove only those artifacts, verify the final diff, and obtain Supervisor confirmation.
  - Implementation Gate: do not remove pre-existing user work or another child's artifacts.
  - Acceptance: `E6-PNB-CLEAN1` records cleanup and Supervisor result.

- [ ] Pn-C: Close and hand off Child 06.
  - Goal: finish final validation, ledgers, detect-changes, commit evidence, and Child 07 handoff.
  - Work Steps:
    1. Run final full build and all accepted affected runtime boundaries.
    2. Refresh actual status, evidence, and benchmark with final values.
    3. Run detect-changes, record all slice commits, verify worktree ownership, and refresh Child 07 from accepted results.
  - Implementation Gate: Pn-A and Pn-B pass; no Child 06 acceptance row remains pending.
  - Acceptance: `E6-PNC-DETECT1`, `E6-PNC-COMMITS1`, and `E6-PNC-HANDOFF1` record final closure and exact Child 07 opening condition.

## Risk Notes

- `buildWorkspace`, `resolveCall`, and `resolveTypeAnnotation` have CRITICAL impact; `classifyDiagnostic` has HIGH impact. Current evidence spans resolver/analyze flows, multiple languages, graph-health, and many tests.
- A declaration-authority mechanism can affect packaging, runtime dependencies, licenses, determinism, performance, and configuration semantics. P6-A must measure and decide these before code.
- Hardcoding the three names can pass target checks while leaving the defect class intact; general fixtures and source inspection must reject that shortcut.
- Loading too broad a declaration set can inflate graph/index size or blur repository ownership. Representation and materialization remain conditional on actual consumers.
- A graph-health-only change can hide a resolver gap. P6-D cannot close unless the underlying P6-C3 outcome is present and persisted.
- Project/package declaration lookup may be necessary, partially necessary, or outside the accepted campaign. P6-C1 must follow P6-A evidence rather than assume coverage.
- Any newly discovered affected consumer expands only the owning slice after plan/actual-status refresh; it does not authorize unrelated command/API changes.
