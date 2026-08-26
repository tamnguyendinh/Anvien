# Anvien Ambient And External Resolution Plan

## Metadata

- Date: `2026-07-28`
- Status: `CLOSED at accepted P6-D commit 81163e39718b94a509e41114cada224e8f269e36 / P6-A through P6-D retain accepted isolated boundaries / performance work transferred to Child 06A`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Persistence dependency: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-02-current-graph-persistence-and-reader-consistency/2026-07-28-02-current-graph-persistence-and-reader-consistency-plan.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/2026-07-28-05-module-export-and-reexport-resolution-plan.md`
- Successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`

## Goal

Make TypeScript standard-library and other evidenced external declarations resolve truthfully on the current production resolver, so `Promise`, `Math.max`, and `Math.min` are external results or explicit external-capability outcomes rather than false in-repository analyzer gaps. Child 06 closes at the accepted P6-D correctness/runtime/reader boundary; analyze-performance work is owned by direct successor Child 06A.

## Rules

- `.tmp` là vùng có thể bị Owner xóa toàn bộ bất kỳ lúc nào, không cần lý do. Vì vậy mọi thứ cần tái sử dụng, tái chạy, nghiệm thu, handoff hoặc làm source of truth đều cấm nằm trong `.tmp`.
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
- P6-D historically closed the current-runtime recovery correction by making `third_party/ladybugdb/v0.19.1/windows-x86_64` the repository-owned LadybugDB native-input authority. Commit `81163e39718b94a509e41114cada224e8f269e36` and its per-run `.tmp` lane/cache/runtime/staging/provenance design remain provenance only; the Owner-approved future canonical build contract below supersedes that machinery without reopening P6-D.
- Child 05 results are immutable predecessor inputs. Child 06 does not repeat module/export traversal.
- Child 02 owns persistence/reader preservation. Child 06 changes only consumers identified by its current affected-surface inventory and fresh impact evidence.
- Before graph-based work, run `anvien analyze --force`. Before editing a resolver, graph-health function, provider, shared contract, or projection, run file-detail and upstream impact and report the full blast radius.
- HIGH or CRITICAL impact is a warning requiring narrow edits and affected-boundary regression, not an edit prohibition.
- Implement production behavior before adding or updating tests. Run the full repository build before boundary validation.
- Every completed production slice requires focused behavior evidence, affected regression, ledger refresh, Supervisor PASS, `anvien detect-changes --repo E:\Anvien --scope all`, and an isolated commit. P6-A is a decision/document slice and records its exact worktree/commit evidence without claiming implementation change detection.
- Child 06 closes at P6-D commit `81163e39718b94a509e41114cada224e8f269e36`. Its former P6-E and performance-dependent aggregate/cleanup/handoff responsibilities are transferred to Child 06A and create no remaining Child 06 Pn gate.
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
- repository-owned LadybugDB `v0.19.1` Windows x86_64 native inputs and the historically accepted P6-D runtime recovery/build evidence, with the future build mechanism governed by the later Owner supersession below;
- direct handoff of accepted P6-D correctness/runtime/reader authority to Child 06A.

Out of scope:

- a preselected declaration storage/authority mechanism;
- a mandatory inventory of declaration-source categories not established by P6-A;
- a new public external-traversal command/API without consumer and product evidence;
- command/persistence protocol work unrelated to the corrected facts;
- hardcoded target names, target-text reclassification, or graph-health-only repair;
- graphing all installed declarations or all dependency files without evidence;
- module/export traversal owned by Child 05;
- scanner behavior, target-source edits, and unrelated language semantics.
- a separate build-recovery slice, alternate full-build workflow, or shared `.tmp\ladybug-native` authority;
- analyze-performance detailed measurement/ranking, dynamic measured-bottleneck attempts, per-attempt architecture/planning/accuracy gates, final performance review, cleanup, commit, and Child 07 handoff now owned by Child 06A.

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

- P6-D materializes exactly `{lbug.h, lbug_shared.lib, lbug_shared.dll}` under `third_party/ladybugdb/v0.19.1/windows-x86_64` as durable repository-owned native inputs with exact provenance/hash records; `.tmp\ladybug-native` and every shared/protected/quarantined cache remain forbidden as build inputs or recovery sources.
- Future canonical build flow is exactly `source -> E:\Anvien\anvien\bin\anvien.exe -> launcher -> analyze`; no per-run lane/cache/runtime/staging/provenance system is part of that contract.
- `scripts/full-build.ps1` stops only the exact processes holding canonical output binaries, runs `go run -mod=vendor ..\cmd\anvien package build-runtime` from `E:\Anvien\anvien`, invokes the launcher build with `pwsh`, validates canonical version plus SHA-256/Git HEAD terminal trace, then runs `E:\Anvien\anvien\bin\anvien.exe analyze . --force`.
- `internal/cli/package_runtime.go` builds directly to `E:\Anvien\anvien\bin\anvien.exe` with checked-in vendor, durable Ladybug native authority, offline guards, and the existing `E:\Anvien\anvien\bin\anvien-runtime.json` sidecar. `anvien-launcher/build.ps1` builds launcher/server/web directly into canonical outputs and must not compile the canonical CLI a second time.
- Future builds reuse the existing default Go/TypeScript/Vite caches. They must not create per-run lane/cache/staging/provenance infrastructure, custom `HOME`/`TEMP`, or override `GOCACHE`, `GOMODCACHE`, `GOPATH`, `GOTMPDIR`, or `NPM_CONFIG_CACHE`. `anvien clean` is graph lifecycle only, never build-cache preparation; do not add new cache infrastructure.
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
- Historical P6-D acceptance proves that the repository-owned native bundle contains the exact three LadybugDB `v0.19.1` Windows x86_64 inputs and that the accepted path contains no `.tmp\ladybug-native` authority. Its then-used unique lane-owned cache/runtime root is historical evidence, not a future build requirement.
- Full build, nearest real resolver/graph-health boundary, target boundary, regression, Supervisor, detect-changes, and per-slice commit evidence pass.
- P6-D is accepted and committed at `81163e39718b94a509e41114cada224e8f269e36`; Child 06 has no remaining aggregate, cleanup, or handoff blocker before Child 06A.

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

- Phase Goal: resolve evidenced external declarations and project truthful outcomes into the graph and graph-health through the accepted P6-D boundary.
- Phase Boundary:
  - In scope: P6-A through P6-D only.
  - Out of scope: Child 05 module exports, broad public API changes, unevidenced declaration sources, and Child 06A analyze-performance work.
  - Dependencies: Child 05 closure; Child 02 affected-reader inventory when persisted fields change.
- Phase Implementation Rule: execute, validate, review, refresh, and commit one slice before opening the next.
- Ordered Slice List:
  - P6-A: Establish the declaration-universe behavior and design from evidence.
  - P6-B: Implement the accepted TypeScript standard-library authority.
  - P6-C1: Implement evidenced project/package declaration lookup.
  - P6-C2: Add the necessary external target representation.
  - P6-C3: Produce structured resolution outcomes.
  - P6-D: Project outcomes and prove the target sites.

- [x] P6-A: Establish the declaration-universe behavior and design from evidence.
  - State: `SUPERVISOR PASS / DECISION COMMITTED / P6-B OPEN`. Child 05 is closed at exact handoff commit `ec765debff335540c77d409ebb2c9f45e4a0a77d`. Initial review rejected only the out-of-boundary contract diff; reject-only repair restored that file to HEAD, final review accepted the unchanged four-ledger/report candidate, and Main committed the exact seven-path decision manifest as `b98131e44932a7bcac17b487ecb2914535927d01`.
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
    - Docker runtime: N/A — no runtime change; use the fresh real CLI boundary. A baseline full build is not run in P6-A because the current hard boundary forbids package install/package-script execution.
    - Playwright target: N/A — no UI behavior.
    - Behavior test: exact local TypeScript `5.9.3` differential vectors for default/ES2022/ES5/ES2015/ES5+Promise/noLib/invalid-lib plus type-value meaning separation; production fixtures remain P6-B work after code.
    - Cleanup/quarantine: keep only reusable independently authored oracle fixtures/evidence.
    - External side effects: none; all mechanism side effects are evaluated, not performed.
    - N/A notes: P6-A is the mandatory evidence/decision slice before production implementation.
  - Work Steps:
    1. **Complete:** inventory current inputs/stages/consumers and run fresh file-detail/impact plus the bounded/general TypeScript differential; record supported and unknown declaration cases.
       - UI flow check: N/A — non-UI decision slice.
       - DB/data flow check: trace each target site from source fact to current gap/health output.
       - Render location check: evidence and actual-status ledgers.
       - Mini QA: exercise the built current resolver/graph-health boundary and record actual output.
       - Evidence target: `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-ORACLE1`, `E6-P6A-CONSUMER1`.
    2. **Complete through Supervisor PASS; exact commit follows immediately:** compare feasible authority mechanisms; record the selected behavior, mechanism, owner map, P6-C1/P6-C2 scope, measurements, and updated later-slice steps. Record the full-build N/A boundary, close the sole contract-boundary rejection without changing the candidate, obtain reject-only Supervisor PASS, then commit the exact decision manifest before opening P6-B.
       - UI flow check: N/A — no UI change.
       - DB/data flow check: required facts/outcomes and affected persistence fields are explicit.
       - Render location check: contract/plan/evidence ledgers.
       - Mini QA: confirm the decision matches observed current/runtime constraints.
       - Evidence target: `E6-P6A-DECISION1`, `E6-P6A-BUILD1`, `E6-P6A-REVIEW1`, `E6-P6A-COMMIT1`.
  - Implementation Gate: satisfied and committed. Exact Child 05 handoff is accepted; graph/source/runtime evidence is fresh; no production edit occurred; final reject-only Supervisor PASS and exact decision commit `b98131e44932a7bcac17b487ecb2914535927d01` are recorded. P6-B is open subject to its own fresh owner preflight.
  - Acceptance:
    - Source: current declaration inputs, stages, consumers, and missing authority are fully traced.
    - Runtime/UI: the fresh real `anvien analyze --force` boundary and independent TypeScript compiler oracle are recorded; the baseline full build is explicitly N/A/not-run under the P6-A package-script boundary; UI is N/A.
    - DB/data: necessary graph/persistence outcome fields and affected consumers are identified.
    - Behavior test: standard-library/unavailable/config/language-isolation oracle cases are specified.
    - Cleanup/quarantine: no rejected mechanism artifact or copied target fixture remains.
    - Evidence IDs: `E6-P6A-IMPACT1`, `E6-P6A-SRC1`, `E6-P6A-ORACLE1`, `E6-P6A-CONSUMER1`, `E6-P6A-DECISION1`, `E6-P6A-BUILD1`, `E6-P6A-REVIEW1`, `E6-P6A-COMMIT1`.
    - Actual-status rows refreshed: authority mechanism, supported config, P6-C1 scope, P6-C2 representation, and affected-reader rows.
  - Evidence Targets: source/impact/oracle/consumer inventory, decision comparison, updated plan, build, Supervisor, commit.
  - Actual-status Update: declaration-authority design `blocked -> decision-recorded -> Supervisor PASS`; implementation rows remain pending and commit-gated.
  - Commit Boundary: commit exactly the four living Child 06 ledgers, immutable Architect report, initial Supervisor REJECT, and final reject-only Supervisor PASS. The resulting hash is the P6-B handoff anchor and is supplied directly when that lane opens; contract, production, tests, target, and protected Main handoffs remain outside the commit.
  - Completion Record: `E6-P6A-REVIEW1` records initial report `5679FB24894F5C51E7AF4EB46FC2D8B534F93A9EB838C68CB5F3CF5FCFD66290`, exact contract-boundary repair, and final PASS `39CE249E9D13F1C77FE6F61DD6E9B1D2E4000B004CE3EBBF461C762F3FA28384`; `E6-P6A-COMMIT1` records exact seven-path decision commit `b98131e44932a7bcac17b487ecb2914535927d01` with parent `ec765debff335540c77d409ebb2c9f45e4a0a77d`.

#### P6-A selected decision

- Authority: offline TypeScript `5.9.3` compiler-API generation from official `lib*.d.ts` -> checked-in versioned compact catalog -> `go:embed` immutable runtime lookup. Runtime Node/`tsc`, network, package install, package-script execution, raw declaration parsing, and target-name maps are forbidden.
- Config: TypeScript only; zero or one root JSONC `tsconfig.json`; support only `compilerOptions.target`, `lib`, and `noLib`. Absent config uses the generated compiler-default profile. Invalid/unreadable values or topology (`extends`, `references`, nested/multiple configs, `files/include/exclude`, `jsconfig.json`, ambiguous ownership) yield `capability_unavailable` and do not change scanner input.
- Precedence: immutable repository/P5 result first; explicit-import failure is terminal; only eligible TypeScript global/type/member misses reach the catalog; external member lookup requires a resolved external receiver; no synthetic `IMPORTS`.
- Outcomes: `resolved_internal`, `resolved_external`, `unresolved`, or `capability_unavailable`, with stage/name/meaning/language/target-or-reason/proof and authority/catalog/config hashes; one final outcome per source site.
- P6-C1: `PRESERVE_ONLY`; no project/package `.d.ts`, ambient-module, augmentation, or `node_modules` lookup is proved in this campaign.
- P6-C2: `ACTIVE`; referenced-only deterministic `ExternalSymbol` nodes with TypeScript/catalog/library/range/meaning provenance are required for graph/Ladybug/reader truth.
- Exact architecture/owner/reader/benchmark contract: `reports/system-architect/rp_system-architect_260822_033607_by_gpt-5_p6a_declaration_universe_decision.md`.

- [x] P6-B: Implement the accepted TypeScript standard-library authority.
  - State: `INDEPENDENT SUPERVISOR PASS / CURRENT DETECT PASS / ISOLATED COMMIT 5c1584fef7153a7a331c8fedd1ce64176ddc873d / CLOSED`; implementation base remains `b98131e44932a7bcac17b487ecb2914535927d01`. The accepted commit's parent is external/user-owned HEAD `fa351c60617212635ef57a43b85d7449ef1eea1c`; both earlier orchestration-only commits and their path remain outside the exact P6-B manifest.
  - Goal: provide deterministic standard-library declaration lookup through the mechanism and config behavior accepted in P6-A.
  - Scope Boundary:
    - Editable after fresh impact: new isolated `internal/tsstdlib` catalog/profile owners, offline generator under `anvien-web/scripts`, minimum analyze/resolution wiring in `internal/analyze/analyze.go`, `internal/resolution/types.go`, `internal/resolution/indexes.go`, and bounded hooks in `internal/resolution/resolve.go`; package manifest/lock only when exact generator pinning requires it.
    - Inspect-only: resolver finalization, external representation, graph-health, and persistence.
    - Preserve-only: other languages and Child 05 module/export resolution.
    - Out of scope: project/package declarations unless P6-A explicitly assigns them here.
  - Non-Goals: no `Promise`/`Math` special case and no design substitution after P6-A without a plan update.
  - Pre-flight Questions:
    - Data source: official TypeScript `5.9.3` `lib*.d.ts` through the exact locked compiler API, generated provenance manifest, and the supported root profile only.
    - Display permission: N/A — resolver data.
    - DB read flow: read only declaration/config inputs authorized by P6-A.
    - DB write flow: write only authority indexes/cache/artifacts required by the accepted design.
    - Render location: resolver trace and benchmark/evidence.
    - UI behavior flow: N/A — no UI behavior.
    - Docker runtime: conditional only if packaged/container runtime is affected; otherwise built CLI/resolver boundary.
    - Playwright target: N/A unless P6-A identifies an affected UI surface.
    - Behavior test: general standard-library types/namespaces/members, `Promise`, `Math.max`, `Math.min`, config variations, unavailable/mismatch, determinism, and language isolation.
    - Cleanup/quarantine: accepted reusable fixtures/assets only; remove rejected output.
    - External side effects: offline deterministic catalog generation only; runtime has no Node/`tsc`, `node_modules` discovery, network, package installation, or package-script execution.
    - N/A notes: project/package lookup and final outcomes are later slices unless P6-A explicitly reassigns them.
  - Work Steps:
    1. **Complete:** refresh exact owner impact; generate and verify the versioned catalog, embed it, parse the supported declaration profile data-only, and implement lookup after immutable repository results without target-name branches.
       - UI flow check: N/A unless affected UI was proved.
       - DB/data flow check: declaration provenance/config and lookup meaning are retained as required by P6-A.
       - Render location check: resolver trace/evidence.
       - Mini QA: exercise the real built authority lookup boundary.
       - Evidence target: `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-PROVENANCE1`.
    2. **Complete through independent Supervisor PASS:** the first repair closed lossless valid-catalog site retention, canonical identity, receiver terminality, and the exact ten-row matrix. The second repair closed typed catalog-validation reason/provenance carriage. The third repair closed the sole remaining contradictory `ready` proof with complete hashes plus a catalog-failure reason. Independent report `reports/Supervisor/rp_supervisor_260822_124138_by_gpt-5_p6b_typescript_stdlib_authority.md` verifies the exact ready/missing/rejected state/status/reason/hash matrix, all `25` direct rows, all retained prior clearances, and no residual same-invariant surface. Current detect and the isolated commit remain Main-owned closure steps, not a second acceptance gate.
       - UI flow check: conditional on P6-A affected surfaces.
       - DB/data flow check: exact general fixture targets and explicit unavailable behavior.
       - Render location check: evidence/benchmark ledgers.
       - Mini QA: inspect real built lookup results rather than helper output.
       - Evidence target: `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-REVIEW1`, `E6-P6B-DETECT1`, `E6-P6B-COMMIT1`.
  - Implementation Gate: satisfied and committed. Fresh pre-edit graph/file-detail/file/symbol impact covered `resolve.go`, both validators, and the exact test owner; the edit remained limited to the single proof-state/status/reason/hash contradiction in the third `E6-P6B-REVIEW1` report. Independent Supervisor PASS is sealed with no residual same-invariant surface. Current post-ledger graph refresh and explicit-path detect both PASS. Main staged exact set `35/35`, preserved all `31` excluded paths, and created isolated commit `5c1584fef7153a7a331c8fedd1ce64176ddc873d`; P6-C1 preserve-only is now open.
  - Invariant Family Map:
    - Family: deterministic TypeScript standard-library authority, root declaration profile, and repository-first resolver eligibility.
    - Authority / SSOT: immutable P6-A report `reports/system-architect/rp_system-architect_260822_033607_by_gpt-5_p6a_declaration_universe_decision.md`, accepted by both P6-A Supervisor reports and committed at `b98131e44932a7bcac17b487ecb2914535927d01`.
    - Primary path: offline exact TypeScript `5.9.3` compiler API -> deterministic compact catalog/provenance -> checked-in `go:embed` asset -> root profile selection -> `analyze.Run` injection -> repository workspace -> eligible TypeScript call/access/type miss -> external declaration lookup.
    - Sibling surfaces: global value, global type, namespace/member, external receiver member, no-config default, explicit target, explicit lib closure, `noLib`, invalid/unreadable/unsupported topology, catalog validation, deterministic regeneration, and non-TypeScript isolation.
    - Preserve-only / forbidden fallback: repository/P5 lookup remains first; explicit-import failure remains terminal; no project/package `.d.ts`, ambient module, augmentation, or `node_modules` scan; no runtime Node/`tsc`, network, install, package script, raw declaration parse, target-name branch, synthetic `IMPORTS`, graph target materialization, shared outcome DTO, graph-health projection, or target access.
    - Stale artifacts to inspect after production code: resolver/analyze tests that assume every TypeScript standard-library miss becomes a generic unresolved diagnostic; package range remains unchanged because the generator itself enforces exact version/integrity.
    - Verify matrix: primary default-profile type/value/member lookups; alternate target/lib profiles; failure `noLib` plus invalid/unreadable/topology/catalog cases; explicit-import terminal and repository-shadow precedence; non-TypeScript isolation; two-run byte equality; embedded load/hash; full build; real built CLI against a repo-local fixture; affected resolver/analyze regression; runtime-without-Node check.
    - Observable success: general catalog-backed lookups increment external-authority resolution without generic repository gaps; authority/profile unavailable cases are explicit authority metrics and never generic repository analyzer gaps; later P6-C2/C3/D surfaces remain unchanged.
  - Acceptance:
    - Source: production authority matches P6-A design and contains no target-name allowlist.
    - Runtime/UI: built lookup behaves correctly for general fixtures; UI is conditional.
    - DB/data: declaration target/provenance/config facts required by the decision are exact.
    - Behavior test: standard-library, config, unavailable, determinism, and language-isolation cases pass.
    - Cleanup/quarantine: rejected outputs and debug artifacts removed; accepted assets are reproducible if applicable.
    - Evidence IDs: `E6-P6B-IMPACT1`, `E6-P6B-SRC1`, `E6-P6B-PROVENANCE1`, `E6-P6B-BUILD1`, `E6-P6B-TEST1`, `E6-P6B-BENCH1`, `E6-P6B-REVIEW1`, `E6-P6B-DETECT1`, `E6-P6B-COMMIT1`.
    - Actual-status rows refreshed: TypeScript standard-library authority and config behavior.
  - Evidence Targets: impact/source/provenance, tests/build/runtime, benchmarks, regression, Supervisor, detect, commit.
  - Actual-status Update: standard-library authority `missing -> implemented / coder-validated -> independent Supervisor REJECT -> four-blocker repair coder-validated -> second independent REJECT on catalog-failure carriage only -> catalog-failure repair coder-validated -> third independent REJECT on the ready proof matrix only -> proof-matrix repair coder-validated -> independent Supervisor PASS -> current detect PASS -> isolated commit 5c1584fef7153a7a331c8fedd1ce64176ddc873d`; P6-C1 preserve-only opens automatically.
  - Commit Boundary: after the post-ledger fresh graph and required explicit-path detect, commit exactly the four living Child 06 ledgers, four tracked P6-B production owners, 24 P6-B source/test/fixture assets, active coder reports `rp_coder_260822_101757_by_gpt-5_p6b_typescript_stdlib_authority_catalog_failure_resubmission.md` and `rp_coder_260822_115544_by_gpt-5_p6b_typescript_stdlib_authority_proof_matrix_resubmission.md`, and final independent PASS report `rp_supervisor_260822_124138_by_gpt-5_p6b_typescript_stdlib_authority.md`. Exclude every Main handoff, root graph/index artifact, prior coder/Supervisor history, and both external orchestration-only commits/path. P6-C1 opens only after this isolated commit.
  - Initial Coder Validation Record: embedded catalog `1,388,709` bytes / SHA-256 `78C9F6F49896D2F11E7E45B58C9D7CB90E1CA20413F06FFE9FAC95310C1AF054`; logical catalog hash `22ab8f6986bfa8efd01addffac64b4265fff09244391ecb5d385b80dfa187119`; focused P6-B packages PASS; production packages and both shipping Go modules build; packaged fixture resolves `6` external declarations with all four external failure counters `0`; full regression has only the pre-existing C#/Dart parity failures; final graph refresh PASS at `2,006/750/0`, `111,524` nodes / `156,259` relationships; explicit-path detect-changes PASS with CRITICAL blast-radius warning. Initial `E6-P6B-REVIEW1` is `REJECT`; a later PASS and `E6-P6B-COMMIT1` remain pending, so the checkbox stays open.
  - Initial Independent Review Record: `E6-P6B-REVIEW1` is `REJECT` in `reports/Supervisor/rp_supervisor_260822_062917_by_gpt-5_p6b_typescript_stdlib_authority.md` (`26,041` bytes / `176` LF / strict UTF-8 without BOM / SHA-256 `26FCA7B7678980F2B129DCF8EA3DB6345FDD269C260FAF8C5326F1B203416FB9`). Blocking scope is exactly four invariants: lossless site-level lookup-result retention, canonical unique semantic IDs, repository/local/import receiver terminality before external member lookup, and exact ten-vector fixture proof. P6-B remains unchecked/uncommitted; P6-C1/C2/C3/D remain locked.
  - Reject-Only Repair Record: lossless immutable `TypeScriptAuthorityResult` records now retain stable source site, request/status/reason, canonical target/owner, declaration ranges, and authority/catalog/profile/config proofs; exact final metrics equal unique site records. Canonical two-phase identity yields `14,587/14,587` unique semantic IDs and zero duplicates/format mismatches across `2,030` symbols and `11,802` direct member rows. Local/repository/import receiver claims are terminal in call and access paths, while genuine external receivers resolve. The exact ten named compiler vectors PASS. Final catalog is `2,003,050` bytes / SHA-256 `F188D15A5D91925DF3E724CBAB97964813E3F6DFD9DF7408FDC7B92EA4CEA487` / logical hash `dca7af22ff26510cf9075fcb587ab76649bf169edcb0001e50c234d89c9dbb0b`; final packaged binary is `73,605,632` bytes / SHA-256 `5BCE56084C58510FE120A8014587EA70E49B5A8C296BC5D038976664C9CEE9AA`. Final self-analyze PASS is `2,013/752/0`, `115,877` nodes / `160,913` relationships; explicit-path detect PASS is CRITICAL `35/8/288/8`. Exact cleanup checked `17` paths with zero remaining, protected Main handoffs are exactly `21`, and P6-C1/C2/C3/D remain locked.
  - Second Independent Review Record: `E6-P6B-REVIEW1` remains `REJECT` in immutable `reports/Supervisor/rp_supervisor_260822_090533_by_gpt-5_p6b_typescript_stdlib_authority.md` (`20,729` bytes / `168` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `C373D2413F1D60082904729F5A536B09D588A5D1DABE12181E454BCE2AD3209A`). It independently cleared canonical identity, receiver terminality, the exact ten-vector matrix, and valid-catalog site carriage. Its sole residual finding is exact typed catalog-validation reason plus lossless per-site provenance/absence carriage.
  - Catalog-Failure Residual Repair Record: `CatalogProofState` makes `ready`, `missing`, and `rejected` semantics explicit. Ready results require authority/logical/artifact hashes; missing results forbid all three; rejected results retain only the attempted-artifact hash and forbid fabricated authority/logical identity. Authority kind, exact TypeScript contract, profile hash, and configuration/inventory proof remain mandatory. Empty/schema/version/input-manifest/logical-hash/trailing-artifact failures preserve exact reasons through value/type/member lookups; direct resolver deduplicates duplicate facts to exactly three unique records and built `analyze.Run` emits exactly seven named stable sites per class with counter equality, zero generic gaps, identical replay, and conflict rejection. Final binary is `73,610,752` bytes / SHA-256 `E7F851FCEC43F7CB740707FF8BBE8CB89E8B741A49FB37E72ADFBF2332182D42`; final self-analyze is `2,023/752/0`, `116,099` nodes / `161,226` relationships; detect is CRITICAL `35/8/307/8`; exact cleanup is `20/20` absent; `23` Main handoffs are pathname-only preserved. P6-C1/C2/C3/D remain locked.
  - Third Independent Review Record: `E6-P6B-REVIEW1` remains `REJECT` in immutable `reports/Supervisor/rp_supervisor_260822_104600_by_gpt-5_p6b_typescript_stdlib_authority.md` (`18,998` bytes / `170` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `C41ACF28BC65020CC75925238B02655D56F27363E952449B9CB3C9CE29D2A422`). It retained every earlier clearance and found exactly one same-invariant contradiction: `CatalogProofReady` with all ready hashes could still pair with `capability_unavailable` plus a catalog rejection reason.
  - Proof-Matrix Residual Repair Record: `validateTypeScriptCatalogProof` now accepts ready proof only for `resolved` with no reason, `profile_excluded/profile_excludes_declaration`, `meaning_mismatch/meaning_mismatch`, or `capability_unavailable` with exactly `disabled_by_no_lib`, `config_invalid`, `config_topology_unsupported`, or `config_unreadable`. Missing/rejected absence semantics are unchanged. Direct verbose validation identifies and passes `25` subrows: `7` legitimate ready positives, `6` invalid cross-status negatives, `6` positive missing/rejected failure classes, and `6` ready-plus-catalog-failure negatives with otherwise complete hashes. Affected authority/resolver/analyze six-class carriage, all P6-B resolver regressions, production and launcher builds, destination package build, ensure/help, and packaged fixture PASS. Final binary is `73,613,824` bytes / SHA-256 `87E8B2696C3851F58BD389B8E0C2A6EFF0D87E41A3A1399D59A6781361CB9BF6`; final post-report self-analyze is `2,028/752/0`, `116,193` nodes / `161,365` relationships; detect is CRITICAL `35/8/328/8`; turn-created cleanup is `9/9` absent; `25` Main handoffs are pathname-only preserved after the concurrent `11:50` rotation. P6-B remains unchecked, unstaged, and uncommitted; P6-C1/C2/C3/D remain locked.
  - Final Independent Review Record: `E6-P6B-REVIEW1` is `PASS` in `reports/Supervisor/rp_supervisor_260822_124138_by_gpt-5_p6b_typescript_stdlib_authority.md` (`16,955` bytes / `143` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `F66E6B8337EC462AE6A7CF03C0C7D5A57176FAF22F7F430E52BF745749FF4C6E`, createdAt/lastWrite `2026-08-22T12:43:32.9400916+07:00`). It independently verifies the exact proof-state/status/reason/hash matrix, preserves every earlier clearance, finds no residual same-invariant surface, and authorizes Main's planner/detect/stage/commit closure. Current source/catalog/binary identities match the report. P6-C1/C2/C3/D remain ordered and locked until the isolated P6-B commit exists.
  - Post-PASS Main Detect Record: after planner finalization, one fresh explicit-path analyze PASS at indexed/current HEAD `fa351c60617212635ef57a43b85d7449ef1eea1c`: `2,030` scanned / `752` parsed code / `0` failed / `116,218` nodes / `161,390` relationships / `19,426` dependency edges. Required explicit-path detect exited `0`, risk CRITICAL: `35` affected symbols / `8` files and `328` changed symbols / `8` files; ResolutionGap delta `198` entities / `201` occurrences; Resolution Health total `0`. CRITICAL is preserved as a blast-radius warning, not a blocker. Exact staging and isolated commit subsequently completed in `E6-P6B-COMMIT1`.
  - Commit Record: Main staged the exact accepted `35`-path manifest with `0` missing and `0` extra, while `31` protected paths remained unstaged (`26` Main handoffs, `3` prior Supervisor reports, `2` older coder reports). Cached diff-check passed. Isolated commit `5c1584fef7153a7a331c8fedd1ce64176ddc873d` (`feat(resolution): add TypeScript standard library authority`) has parent `fa351c60617212635ef57a43b85d7449ef1eea1c`, exactly `35` paths, `4,877` insertions / `57` deletions, and post-commit index empty. P6-B is closed; no second review/detect gate is created by this ledger record.

- [x] P6-C1: Implement evidenced project/package declaration lookup.
  - State: `PRESERVE_ONLY BY P6-A DECISION / INDEPENDENT SUPERVISOR PASS / ISOLATED COMMIT 8055f0a6860721e26462572e34469e0d708d4a52 / CLOSED`; P6-C2 is now open.
  - Goal: close preserve-only with proof that this campaign does not require project/package declaration lookup; activation requires a new planner decision before code.
  - Scope Boundary:
    - Editable: Child 06 plan/evidence/benchmark/actual-status closure only; no production or test owner is authorized.
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
    - Behavior test: prove project/package `.d.ts`, ambient modules/augmentations, and `node_modules` lookup remain unavailable and do not duplicate Child 05 resolution; no production fixture is added.
    - Cleanup/quarantine: independent fixtures; no copied dependency tree.
    - External side effects: none; no dependency-tree scan, install, network, or package execution.
    - N/A notes: P6-A proved this behavior is not required for the campaign, so P6-C1 closes without production/test diff.
  - Work Steps:
    1. **Worker complete:** reconfirm the accepted P6-A decision and current boundary; record the evidence-backed preserve-only result and no-production-owner proof.
       - UI flow check: N/A — non-UI lookup.
       - DB/data flow check: module result and declaration lookup remain separate.
       - Render location check: evidence/actual-status.
       - Mini QA: exercise the real current lookup boundary for required cases.
       - Evidence target: `E6-P6C1-IMPACT1`, `E6-P6C1-SCOPE1`.
    2. **Complete through independent Supervisor PASS; Main commit pending:** validate the preserved unavailable behavior after P6-B, record full build and implementation detect as N/A for the zero-production/test diff, seal one durable worker report, obtain independent Supervisor PASS, then let Main commit the exact accepted doc/report-only closure.
       - UI flow check: conditional on affected UI evidence.
       - DB/data flow check: exact declaration target or explicit external-capability outcome.
       - Render location check: resolver trace/evidence.
       - Mini QA: inspect built required-case results.
       - Evidence target: `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-TEST1`, `E6-P6C1-REVIEW1`, `E6-P6C1-DETECT1`, `E6-P6C1-COMMIT1`.
  - Implementation Gate: satisfied for acceptance. P6-A and P6-B are accepted; fresh current evidence reconfirms preserve-only status, zero active cases, zero production/test owners, and no implementation diff. Independent report `reports/Supervisor/rp_supervisor_260822_142120_by_gpt-5_p6c1_project_package_preserve_only.md` returns PASS with no residual same-invariant surface. Only exact staging and the isolated doc/report commit remain before P6-C2 opens.
  - Acceptance:
    - Source: no project/package declaration behavior and no production/test diff.
    - Runtime/UI: required real lookup cases pass; UI is conditional.
    - DB/data: project/package result remains distinct from module/export selection and retains required proof.
    - Behavior test: accepted present/missing/config/security cases or preserve-only proof passes.
    - Cleanup/quarantine: no copied dependency tree or obsolete fixture remains.
    - Evidence IDs: `E6-P6C1-IMPACT1`, `E6-P6C1-SCOPE1`, `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-TEST1`, `E6-P6C1-REVIEW1`, `E6-P6C1-DETECT1`, `E6-P6C1-COMMIT1`.
    - Actual-status rows refreshed: project/package declaration scope and implementation state.
  - Evidence Targets: scope decision, impact, source/preserve proof, tests/build/runtime, Supervisor, detect, commit.
  - Actual-status Update: conditional lookup `blocked -> preserve-only -> worker-validated -> independent Supervisor PASS -> isolated commit 8055f0a6860721e26462572e34469e0d708d4a52`; P6-C2 opens automatically.
  - Commit Boundary: commit exactly the four Child 06 ledgers, worker report `rp_coder_260822_135011_by_gpt-5_p6c1_project_package_preserve_only.md`, and PASS report `rp_supervisor_260822_142120_by_gpt-5_p6c1_project_package_preserve_only.md`; exclude every protected Main/prior report and all production/test/package/generated paths.
  - Worker Validation Record: one fresh explicit-path self-analyze PASS records `2,030/752/0`, `116,218` nodes / `161,390` relationships / `19,426` dependency edges. Accepted runtime/test owner searches are `0/0`; current production/test/package diff is `0`; the two tracked frontend/tooling `.d.ts` paths and all `node_modules` paths are scanner-excluded and fresh graph counts are `0/0`. Root `tsconfig.json` is absent and nested/multiple topology remains fail-closed. Full build and implementation detect are N/A because no production/test/fixture/package/runtime artifact changed. Evidence: `E6-P6C1-IMPACT1`, `E6-P6C1-SCOPE1`, `E6-P6C1-SRC1`, `E6-P6C1-BUILD1`, `E6-P6C1-TEST1`. Report: `reports/coder/rp_coder_260822_135011_by_gpt-5_p6c1_project_package_preserve_only.md`.
  - Independent Review Record: `E6-P6C1-REVIEW1` is PASS in `reports/Supervisor/rp_supervisor_260822_142120_by_gpt-5_p6c1_project_package_preserve_only.md` (`15,401` bytes / `179` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `C1C5D0DEDC2F910D29FB32C9666E192A6221041DB578306F107B504FAF357E9B`, createdAt/lastWrite `2026-08-22T14:23:40.8224653+07:00`). Its fresh graph is `2,032/752/0`, `116,251` nodes / `161,423` relationships / `19,426` dependency edges; exact `.d.ts` and `node_modules` node counts remain `0/0`. It accepts the exact five-path pre-report candidate, reports no residual same-invariant surface, and authorizes only Main's planner/stage/commit closure.
  - Commit Record: Main staged exact set `6/6` with zero missing/extra and cached diff-check PASS. Isolated commit `8055f0a6860721e26462572e34469e0d708d4a52` (`docs(resolution): close P6-C1 preserve-only`) has parent `5c1584fef7153a7a331c8fedd1ce64176ddc873d`, exactly the four Child 06 ledgers plus worker/PASS reports, `454` insertions / `31` deletions, post-commit index empty, and exactly `32` protected untracked history paths. P6-C1 is closed without a second acceptance/detect gate.

- [x] P6-C2: Add the necessary external target representation.
  - State: `INDEPENDENT SUPERVISOR PASS / PLANNER FINALIZED / CLEAN CURRENT ANALYZE-DETECT REUSED ON UNCHANGED IDENTITIES / ISOLATED COMMIT 2acd357db32ac0c2bd3c3968a950c773acf3000d / CLOSED`; two procedural REJECT reports remain immutable protected history and no coder repair exists.
  - Goal: materialize referenced resolved declarations as deterministic `ExternalSymbol` facts and preserve explicit external-capability gaps without repository ownership pollution.
  - Scope Boundary:
    - Edited after fresh impact: `internal/scopeir/kinds.go`, isolated materializer `internal/resolution/external_symbol.go`, the minimum resolver/emitter hook, `internal/lbugload/csv.go`, `internal/lbugschema/schema.go`, `internal/processes/processes.go`, `internal/mcp/context.go`, and `internal/mcp/rename.go`. `internal/mcp/impact.go` was validated but not edited because it already consumes the shared context projection.
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
    1. **Worker complete:** record graph/reader impact and select the minimum representation; current consumers match P6-A and require no planner expansion.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: representation preserves external origin and explicit unresolved capability.
       - Render location check: graph/affected-reader evidence.
       - Mini QA: inspect real graph and nearest affected reader.
       - Evidence target: `E6-P6C2-IMPACT1`, `E6-P6C2-DESIGN1`.
    2. **Complete through independent Supervisor PASS and isolated commit:** implement production representation/materialization, then tests, build, affected parity, regression, review, detect, and commit.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: exact target/provenance and no repository File/Definition misclassification.
       - Render location check: graph and affected readers only.
       - Mini QA: exercise real built graph-bound consumer(s).
       - Evidence target: `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-PARITY1`, `E6-P6C2-REVIEW1`, `E6-P6C2-DETECT1`, `E6-P6C2-COMMIT1`.
  - Implementation Gate: satisfied and closed. P6-B and P6-C1 are committed; every edited existing owner has fresh file-detail/file+symbol impact; two independent technical reviews establish no residual code/test/ledger/runtime invariant; the final commandless procedural-only review closes the sole no-C execution-boundary blocker; unchanged accepted identities allowed reuse of the clean current E-rooted analyze/detect; isolated commit `2acd357db32ac0c2bd3c3968a950c773acf3000d` closes P6-C2 and opens P6-C3. P6-D/target remain locked.
  - Acceptance:
    - Source: minimum accepted representation distinguishes internal, external, intrinsic when applicable, and explicit external-capability gaps.
    - Runtime/UI: affected built consumers show the same truthful result; UI is conditional.
    - DB/data: external origin/provenance survives affected persistence with zero repository ownership pollution.
    - Behavior test: resolved/unavailable/separation/duplicate/parity cases pass.
    - Cleanup/quarantine: no incomplete or superseded representation artifact remains.
    - Evidence IDs: `E6-P6C2-IMPACT1`, `E6-P6C2-DESIGN1`, `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-PARITY1`, `E6-P6C2-REVIEW1`, `E6-P6C2-DETECT1`, `E6-P6C2-COMMIT1`.
    - Actual-status rows refreshed: external target representation and affected readers.
  - Evidence Targets: impact/design/source, tests/build/parity/runtime, Supervisor, detect, commit.
  - Actual-status Update: external representation `missing -> correct / worker-validated -> two procedural REJECTs with no technical blocker -> independent Supervisor PASS / planner-finalized -> isolated commit 2acd357db32ac0c2bd3c3968a950c773acf3000d / closed`; affected P6-C2 readers are independently parity-cleared while graph-health remains locked to P6-D.
  - Commit Boundary: after one current explicit-path analyze/detect, commit exactly `23` accepted paths: `9` production + `8` tests + `4` Child 06 ledgers + coder report `rp_coder_260822_162635_by_gpt-5_p6c2_external_target_representation.md` + final PASS report `rp_supervisor_260822_175500_by_gpt-5_p6c2_external_target_representation_procedural_only.md`. Exclude both procedural REJECT reports and every other protected history path.
  - Worker Validation Record: final P6-B rows drive deterministic `ExternalSymbol` materialization only for `resolved` sites; unavailable/excluded/mismatch rows create no fake target. The packaged two-file fixture emits `3` unique external nodes, `4` coalesced external edges, and `6` immutable per-site authority proof rows with `0` repository ownership edges. Focused and full affected packages pass; native Ladybug COPY/readback passes with zero fallback/skips. Broad `go build ./...` remains false because repository fixtures are intentionally non-buildable; broad `go test ./internal/...` preserves seven repo-local-temp Git-environment failures plus the existing C#/Dart ACCESS parity failures. Production `cmd/anvien`, local packaged destination build, packaged runtime/context/impact/rename, and two-run graph byte equality pass. Final packaged root analyze is `2,039/756/0`, `116,657` nodes / `162,181` relationships / `19,658` dependency edges; explicit detect exits `0` at CRITICAL `44` affected symbols / `17` files and `202` changed symbols / `17` files, with Resolution Health `0`. Evidence: `E6-P6C2-IMPACT1`, `E6-P6C2-DESIGN1`, `E6-P6C2-SRC1`, `E6-P6C2-BUILD1`, `E6-P6C2-TEST1`, `E6-P6C2-PARITY1`, `E6-P6C2-DETECT1`.
  - Independent Review Record: first report `reports/Supervisor/rp_supervisor_260822_171803_by_gpt-5_p6c2_external_target_representation.md` (`18,318` bytes / `229` LF / SHA-256 `C1F7C1959C3BB36CA1F91817FC5FDCF8DCFD193388760394BEB59325D597FA88`) and clean E-rooted report `reports/Supervisor/rp_supervisor_260822_174938_by_gpt-5_p6c2_external_target_representation_clean_rereview.md` (`17,759` bytes / `224` LF / SHA-256 `537F59AA8A350349E0ADC1858C74A4DDA3B7B3102292D9A1E69AFF942F682B2C`) each return procedural-only `REJECT`, explicitly retain all technical clearances, establish no residual code invariant, and request no coder repair. Final commandless report `reports/Supervisor/rp_supervisor_260822_175500_by_gpt-5_p6c2_external_target_representation_procedural_only.md` returns `PASS` (`5,583` bytes / `75` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `D55D2F33BD311E4266D537E17F6A5EAC369FAFACA2650BD41C40140BC1F9EA0C`, createdAt/lastWrite `2026-08-22T17:56:10.5347129+07:00`). It consumes the externally rehashed unchanged `22/22` candidate, preserves both REJECT reports as immutable history, closes only their tool-substrate no-C blocker, and reports residual same-invariant surfaces `none`.
  - Commit Record: successor re-anchor at `2026-08-22T18:25:06.7772180+07:00` verified HEAD `8055f0a6860721e26462572e34469e0d708d4a52`, exact accepted manifest `23/23`, missing/extra `0/0`, empty index, `git diff --check` PASS, unchanged report/ledger seals, and protected outside boundary `39 = 32 Main + 5 Supervisor + 2 coder`. The clean E-rooted post-finalization analyze (`2,045/756/0`, `116,733` nodes / `162,257` relationships / `19,658` dependency edges) and explicit detect (CRITICAL `44/17/202/17`, ResolutionGap `33/33`, Resolution Health `0`, semantic completeness `116,733/116,733`) were reused because accepted identities and boundary were unchanged. Exact staging was `23/23`, cached diff-check PASS, both REJECT reports remained unstaged, and isolated commit `2acd357db32ac0c2bd3c3968a950c773acf3000d` (`feat(resolution): add external target representation`) has parent `8055f0a6860721e26462572e34469e0d708d4a52`, exactly `23` paths, `1,491` insertions / `137` deletions, empty post-commit index, and exactly `39` protected untracked history paths.

- [x] P6-C3: Produce structured resolution outcomes.
  - State: `INDEPENDENT SUPERVISOR PASS / CURRENT DETECT PASS / ISOLATED COMMIT 0d9803777880d8ecca09cae90592d4728a231160 / CLOSED`. The exact 26-path commit contains the accepted worker candidate, four finalized ledgers, and the single PASS report. P6-D is now open; graph-health/target actions belong only to that slice.
  - Goal: finalize exactly one immutable resolver outcome per affected source site from repository, external, intrinsic, or explicit external-capability results.
  - Scope Boundary:
    - Editable: exact outcome/status owner and minimum finalization adapters selected by impact; the implemented validation-unblocker uses exactly six existing test files—`resolution_test.go`, `type_alias_test.go`, `binding_accumulator_test.go`, `graph_parity_test.go`, `identity_occurrence_test.go`, and `p3c_binding_occurrence_test.go`—plus the sole new owner `parser_integration_test.go`. Owner expansion additionally permits only `legacy_scope_symbol_semantics_conversion_test.go` and the directly implicated synthetic facts in `resolution_test.go` for the exact ten current failures, after fresh current-byte impact.
    - Inspect-only: P5/P6 lookup results, graph representation, graph-health, parser-dependent legacy behavior, and every helper/dependency edge outside the exact six-file closure.
    - Preserve-only: declaration lookup decisions, production `internal/resolution` parser-free imports/semantics, parser-dependent assertions and coverage, unrelated diagnostics, and all already-cleared P6-C3 carriers.
    - Out of scope: any owner beyond the exact seven existing test files now named, a second new integration owner, declaration loading, parser refactors, shared contracts/statuses, graph-health rendering, P6-D, and target access. Production resolution may change only if expanded fresh evidence proves a candidate regression cannot be corrected truthfully at the implicated fixtures; current attribution selects no production edit. Any parser/shared-contract/second-new-owner need stops this lane.
  - Non-Goals: no exhaustive status set unrelated to accepted current cases; no target-text inference; no deleting, weakening, skipping, pass-by-default conversion, or silent loss of parser-dependent legacy coverage; no production-semantic edit merely to make tests compile.
  - Pre-flight Questions:
    - Data source: immutable P5 repository results and P6-B/C results.
    - Display permission: N/A — resolver outcome.
    - DB read flow: read accepted stage results only.
    - DB write flow: write one final outcome and its required persisted fields.
    - Render location: resolver trace and affected persistence/readers.
    - UI behavior flow: N/A unless affected-reader inventory says otherwise.
    - Docker runtime: conditional on affected app/UI surface.
    - Playwright target: conditional on affected browser surface.
    - Behavior test: every accepted current status, stage precedence, one-result exclusivity, proof/candidates, and unavailable authority; default core `go test ./internal/resolution` must compile/run from the E-only offline cache, while parser-dependent legacy behavior retains a dedicated explicit integration command whose PASS or environment BLOCKED result is recorded truthfully.
    - Cleanup/quarantine: isolated independently authored outcome fixtures.
    - External side effects: none.
    - N/A notes: graph-health mapping is P6-D.
  - Work Steps:
    1. **Initial worker candidate independently rejected; exact invariant repair locally validated:** retain the sealed REJECT as immutable history, repair repository-before-intrinsic terminality and carrier-independent call/heritage finalization, and add the three hostile durable fixtures without broadening the status/shared contract.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: exactly one outcome per source site; no later stage overwrites an authoritative failure.
       - Render location check: resolver trace/affected persistence.
       - Mini QA: inspect real built outcome output for representative cases.
       - Evidence target: `E6-P6C3-IMPACT1`, `E6-P6C3-SRC1`, `E6-P6C3-STATUS1`.
    2. **Worker and independent review complete with truthful cleanup disposition:** direct parser imports were exactly `resolution_test.go` and `type_alias_test.go`; shared helpers expanded the closure to the four authorized indirect files. The minimum separation uses one new `parser_integration_test.go`, preserves every assertion under explicit tag, and makes the default full resolution package parser-free and PASS. Dedicated integration remains truthfully offline BLOCKED. The coder-sealed `E:\Anvien\.tmp\p6c3-repair` identity was `1,672 / 76,977,924 / 17`; review-induced `registry.json` drift makes the preserved live root `1,673 / 76,978,267 / 17`. It must not be written, used, cleaned, or worked around. The distinct review-owned root `E:\Anvien\.tmp\p6c3-supervisor-rereview-01a02a84` is `1,662 / 93,162,861 / 17`; Main's exact-contained recursive removal was rejected by policy before process creation. Both roots therefore remain explicitly PARTIAL/BLOCKED, never clean, and no workaround is authorized.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: status/stage/target/reason/proof survive affected boundaries.
       - Render location check: evidence/benchmark ledgers.
       - Mini QA: exercise nearest real resolver/graph-bound boundary.
       - Evidence target: `E6-P6C3-BUILD1`, `E6-P6C3-TEST1`, `E6-P6C3-PARITY1`, `E6-P6C3-REVIEW1`, `E6-P6C3-DETECT1`, `E6-P6C3-COMMIT1`.
    3. **Worker complete through final graph/detect; report seal pending:** default package execution exposed exactly ten legacy fixture candidate regressions. All ten were attributed before edit; four ledgers, fresh graph, both owner impacts, and all ten exact symbol impacts preceded the range-only repair. Exactly fifteen synthetic facts and one expected SourceSiteID were corrected. Exact ten tests, default full package, P6-C3, persistence, and native boundaries pass without production edit.
       - UI flow check: N/A — test fixtures only.
       - DB/data flow check: each synthetic fact becomes a stable positive-line source site without changing target/reason/proof semantics.
       - Render location check: outcome evidence/diagnostics remain unchanged except their now-complete site range.
       - Mini QA: exact ten tests, then default full package.
       - Evidence target: `E6-P6C3-IMPACT1`, `E6-P6C3-TEST1`, `E6-P6C3-DETECT1`.
  - Implementation Gate: satisfied and committed. Independent Supervisor report `DBA2111E7739AEB95E3E74CC29A5CA79F122908E34A8D5B75953D55CADA72E0F` is PASS with residual same-invariant surfaces none. Main current explicit-path analyze exits `0` at `2,064/762/0`, `117,576` nodes / `163,875` relationships / `20,221` dependency edges; detect exits `0`, CRITICAL `36` affected symbols / `13` files and `106` changed symbols / `16` files, ResolutionGap `27/27`, Resolution Health `0`, semantic completeness `117,576/117,576`. Exact staging was `26/26`, cached diff-check PASS, and isolated commit `0d9803777880d8ecca09cae90592d4728a231160` has parent `2acd357db32ac0c2bd3c3968a950c773acf3000d`, missing/extra `0/0`, empty post-commit index, and exactly `49` protected reports outside the commit. Cleanup remains explicitly policy-blocked and is not called clean.
  - Acceptance:
    - Source: one structured immutable outcome per affected site; no target-text guess or later override.
    - Runtime/UI: built resolver outputs exact accepted outcomes; UI is conditional.
    - DB/data: required fields survive affected persistence/readers with zero resolved/unresolved overlap.
    - Behavior test: all accepted status and precedence cases pass; the default full `internal/resolution` package test no longer imports parser-only integration dependencies; the same parser-dependent assertions remain reachable through a dedicated explicit command with truthful PASS/BLOCKED evidence.
    - Cleanup/quarantine: no unused status or obsolete fixture remains.
    - Evidence IDs: `E6-P6C3-IMPACT1`, `E6-P6C3-SRC1`, `E6-P6C3-STATUS1`, `E6-P6C3-BUILD1`, `E6-P6C3-TEST1`, `E6-P6C3-PARITY1`, `E6-P6C3-REVIEW1`, `E6-P6C3-DETECT1`, `E6-P6C3-COMMIT1`.
    - Actual-status rows refreshed: structured outcome and affected persistence rows.
  - Evidence Targets: impact/source/status, tests/build/parity/regression, Supervisor, detect, commit.
  - Actual-status Update: initial structured-outcome candidate `worker-validated -> independent REJECT on two owning invariants -> focused repair and validation-unblocker worker-validated -> independent Supervisor PASS -> current detect PASS -> isolated commit 0d9803777880d8ecca09cae90592d4728a231160 / closed`; P6-D opens automatically.
  - Commit Boundary: after one current accepted-byte explicit-path analyze/detect, stage exactly `26` accepted paths: the sealed 25-path corrected candidate plus `reports/Supervisor/rp_supervisor_260823_020452_by_gpt-5_p6c3_structured_resolution_outcomes_rereview.md`. Exclude the old coder report, sealed REJECT, all other protected history, both `.tmp` roots, P6-D/target, shared graph/Ladybug schema, and unrelated paths. Pre-planner post-report Git is `75 = 25` candidate + `1` PASS report + `49` protected, with index empty.
  - Worker Validation Record: sealed independent report `reports/Supervisor/rp_supervisor_260822_215636_by_gpt-5_p6c3_structured_resolution_outcomes.md` rejects repository/predeclared-name precedence and two carrier-coupled zero-outcome exits while retaining all other clearances. Same-worker production repair closes those invariants. Minimum six-file separation plus one tagged owner removes parser imports from the default resolution package while preserving tagged coverage. Owner-expanded fixture-only repair closes the ten newly executable candidate regressions `10/10`; full default resolution, all P6-C3, CSV parity, and P6-C3 native readback pass. Full build, tagged integration, and `internal/analyze` package execution remain truthfully offline/baseline blocked. Corrected candidate manifest is `25` paths / canonical `3,480` bytes / SHA-256 `DFCDCBBF17CBB3F48921E64E7DCA9087C2CEFF8A2926E8E023C2F477EB593285`.
  - Independent Review Record: `E6-P6C3-REVIEW1` is PASS in `reports/Supervisor/rp_supervisor_260823_020452_by_gpt-5_p6c3_structured_resolution_outcomes_rereview.md` (`26,590` bytes / `222` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `DBA2111E7739AEB95E3E74CC29A5CA79F122908E34A8D5B75953D55CADA72E0F`, created/last-write `2026-08-23T02:07:53.5247293+07:00`). It finds residual P6-C3 same-invariant surfaces none, truthfully preserves broad-build/tagged/analyze/package blockers, classifies the `registry.json` drift as review-induced non-candidate contamination, and requires Main to record cleanup disposition before current analyze/detect/stage/commit. Main's exact review-root cleanup attempt was policy-rejected before process creation; no mutation or workaround occurred.
  - Main Closure Record: the first Main analyze setup used the reserved PowerShell `$HOME` variable name; its assignment error was non-terminating, so PID `7788` briefly started with inherited home before Main detected and interrupted it. That invocation is excluded from evidence; no C path was inspected or cleaned, and a later correct E-rooted `--force` analyze fully superseded the repository graph. Exact E-rooted Main home/cache/temp used `E:\Anvien\.tmp\p6c3-main-closure-01a02abb`; it now contains only one `343`-byte registry file. One explicit non-recursive file/empty-directory cleanup attempt was policy-rejected before process creation; no mutation or workaround occurred. Correct analyze/detect results are the current numbers in Implementation Gate.
  - Commit Record: Main staged exact set `26/26` with zero missing/extra, zero unstaged tracked paths, and cached/worktree diff-check PASS. Isolated commit `0d9803777880d8ecca09cae90592d4728a231160` (`feat(resolution): add structured resolution outcomes`) has parent `2acd357db32ac0c2bd3c3968a950c773acf3000d`, exactly `26` paths, `3,337` insertions / `557` deletions, post-commit index empty, and remaining status exactly `49 = 40` Main + `6` prior Supervisor + `3` old coder reports. P6-C3 is closed without a second review/detect loop.

- [x] P6-D: Project outcomes and prove the target sites.
  - State: `INDEPENDENT SUPERVISOR PASS / CANONICAL FULL BUILD AND CURRENT RUNTIME PASS / READER GATE PASS / TARGET CAPABILITY OUTCOMES 3/3 / IN-REPOSITORY ANALYZER GAPS 0/3 / TARGET BOUNDARY PASS / PROCEDURE CLEAN / MAIN FRESH CURRENT-BYTE DETECT PASS / ISOLATED COMMIT 81163e39718b94a509e41114cada224e8f269e36 / CHILD 06 CLOSED / CHILD 06A SUCCESSOR`.
  - Owner-Superseded Future Build Record: P6-D stays closed and commit `81163e39718b94a509e41114cada224e8f269e36` stays accepted provenance. Every remaining P6-D reference to per-run `.tmp` lane/cache/runtime/staging/provenance, custom home/temp roots, or PowerShell 5.1 describes historical execution only and must not drive a future build. The binding future flow and command/cache constraints are the `Persistence and target validation` contract above.
  - Goal: preserve the cleared graph-health projection, materialize the selected exact repository-owned `vendor/` dependency authority alongside the durable LadybugDB inputs, recover the canonical current-runtime path with unique lane-owned derived caches/runtime roots, then validate the three bounded target sites on that built current runtime.
  - Scope Boundary:
    - Editable for the approved correction: `scripts/materialize-go-vendor.ps1`, `scripts/verify-go-vendor.ps1`, `scripts/test-go-vendor-contract.ps1`, `vendor/**`, `third_party/go-vendor/manifest.v1.json`, `third_party/go-vendor/licenses/**`, and exactly `third_party/go-vendor/patches/tree-sitter-go-v0.25.0-remove-absent-scanner.patch`; original P6-D scope also retains the durable Ladybug bundle, `scripts/full-build.ps1`, `scripts/ensure-ladybug-native.ps1`, `internal/cli/package_runtime.go`, `internal/cli/package_command_test.go`, `internal/cli/package_runtime_p6d_test.go`, `anvien-launcher/build.ps1`, and `scripts/qa-child02-p2e-validation.ps1`. Any further owner requires fresh impact and another four-ledger refresh. Existing graph-health production/test bytes remain preserve-only unless a fresh same-invariant rejection requires repair.
    - Inspect-only: unchanged `go.mod`/`go.sum`, resolver outcomes, graph/persistence, `anvien/package.json` lifecycle wiring, canonical publication destinations, target source/worktree, and independent oracle.
    - Preserve-only: the independently cleared six-field projector/test bytes, committed P6-C3, unrelated health policies, target source, and all protected/shared/quarantined caches.
    - Out of scope: resolver/declaration behavior accepted in earlier slices, an alternate full-build workflow, a separate recovery slice/commit, and every Child 06A performance optimization.
  - Non-Goals: no `Promise`/`Math` branch, no reclassification without resolver evidence, no `.tmp\ladybug-native` build authority, no hand-pruned partial vendor tree, and no silent network/download/shared-cache/module-cache fallback.
  - Pre-flight Questions:
    - Data source: final structured outcomes, exact target source sites, sealed blocker report `E6-P6D-RECOVERY1`, Owner correction, and the exact durable LadybugDB bundle.
    - Display permission: preserve current permissions; include UI only if impact proves it affected.
    - DB read flow: read graph/Ladybug and only affected graph-health/readers.
    - DB write flow: canonical build stages current runtime output under the unique lane runtime root; normal analyze writes current graph output; final publication preserves existing explicit product destinations and failure behavior.
    - Render location: graph-health CLI/MCP/HTTP/Web only where affected.
    - UI behavior flow: conditional on affected UI evidence.
    - Docker runtime: conditional on affected app/UI; otherwise the corrected canonical `scripts/full-build.ps1` plus the real lane-built graph-health commands.
    - Playwright target: conditional on affected browser surface.
    - Behavior test: durable-bundle selection, zero shared-root fallback, unique-root propagation, current-runtime provenance, mechanical status mapping, exact three target outcomes, zero in-repository analyzer gaps, parity, and target boundary.
    - Cleanup/quarantine: remove only the active lane's unique cache/runtime artifacts when authorized; preserve the durable third-party bundle, protected/shared/quarantined roots, and target source/worktree.
    - External side effects: canonical product publication and normal target analyze only after their gates; target source remains unchanged.
    - N/A notes: build recovery and target proof remain internal work steps of P6-D, not a new slice/gate/commit; performance work follows only in independent Child 06A.
  - Work Steps:
    1. Preserve the independently cleared six-field projector and refresh file-detail/impact only for the exact native/build recovery owners before edit; do not re-audit or rewrite accepted P6-C3/P6-D semantics.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: accepted structured outcomes and projection carriers remain byte/behavior preserve-only.
       - Render location check: recovery scope is native/build ownership only.
       - Mini QA: verify the exact editable/preserve-only map before source action.
       - Evidence target: `E6-P6D-IMPACT1`, `E6-P6D-RECOVERY1`.
    2. **Completed and independently accepted:** put the exact LadybugDB `v0.19.1` Windows x86_64 header/import-library/DLL into `third_party/ladybugdb/v0.19.1/windows-x86_64`; update `scripts/full-build.ps1` first, then the named helper/package/launcher consumers, so the durable bundle and distinct cache/runtime roots under one unique repository-local lane root flow end to end and `.tmp\ladybug-native` is never an authority or fallback. After those production/build owners are correct, update the exact native-reader QA consumer so its `Get-NativeEnvironment` uses the same durable authority rather than the shared temporary path. Add/update tests only after the production/script contract is correct.
       - UI flow check: N/A unless fresh impact proves an affected UI build surface.
       - DB/data flow check: native inputs are immutable build inputs; cache/runtime output is derived, lane-owned, and never promoted as authority.
       - Render location check: canonical publication destinations remain explicit; validation uses the current lane-built runtime.
       - Mini QA: missing/partial/extra bundle fails visibly; two lanes cannot share cache/runtime roots; package, launcher, and the exact native-reader QA consumer resolve the same durable input identity.
       - Evidence target: `E6-P6D-SRC1`, `E6-P6D-BUILD1`.
    3. **Historical standard-baseline rejection corrected; final closure independently accepted:** retain deterministic standard `go mod vendor` equality `2/2` as the baseline, then machine-derive a module-local C/C++ include fixed-point from exact checksum-verified zip/source authority, including `CFLAGS`, `CPPFLAGS`, recursive headers and conditional references. Copy complete parent source roots into their module-relative `vendor/` locations, then apply exactly one guarded byte-exact patch removing the absent `tree-sitter-go@v0.25.0/src/scanner.c` include only when binding/grammar/parser/zip predicates all match. Manifest v1 gains `closureContractVersion=2` partitions for standard baseline, supplemental closure, reviewed patch and final tree; verifier must fail closed on drift and pass under both `pwsh` and Windows PowerShell 5.1 before the first Go process. Preserve unchanged `go.mod`/`go.sum`, explicit `-mod=vendor`, `GOPROXY=off`/`GOSUMDB=off`, and all lane-local derived-cache controls.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: Graph JSON/Ladybug/native-reader outcomes and proofs remain exact with zero fallback, skip, or shared-root use.
       - Render location check: real built CLI/readers only.
       - Mini QA: verify runtime VCS/source/native-bundle/cache/runtime-root provenance before accepting reader output.
       - Evidence target: `E6-P6D-VENDOR1`, `E6-P6D-VENDOR2`, `E6-P6D-BUILD1`, `E6-P6D-PARITY1`.
    4. **Completed through independent Supervisor PASS, Main fresh current-byte detect, and isolated commit:** captured target pre-state; analyzed `E:\cheapapp.org` with the current lane-built runtime; compared the independent TypeScript oracle; proved `Promise`, `Math.max`, and `Math.min` correct external/capability outcomes with no in-repository analyzer gap; verified target boundary; obtained review; and closed the exact accepted boundary at commit `81163e39718b94a509e41114cada224e8f269e36`.
       - UI flow check: conditional on affected UI.
       - DB/data flow check: exact three outcome/proof records are the acceptance source.
       - Render location check: official evidence under Anvien.
       - Mini QA: run real target query/graph-health/file-detail and visually inspect any affected UI result.
       - Evidence target: `E6-P6D-TARGET1`, `E6-P6D-ORACLE1`, `E6-P6D-BOUNDARY1`, `E6-P6D-REVIEW1`, `E6-P6D-DETECT1`, `E6-P6D-COMMIT1`.
  - Implementation Gate:
    - P6-A/B/C remain accepted; the independently cleared P6-D projector/test bytes and protected history remain preserve-only unless a fresh same-invariant rejection says otherwise.
    - Before editing the named durable-native/build surfaces, run their fresh file-detail and upstream impact; any additional same-invariant owner requires a four-ledger refresh before edit.
    - Recovery graph `2,100/765/0`, `122,147/168,747/20,335`, with `479` unresolved projections, plus fresh file-detail/impact authorized exactly `scripts/qa-child02-p2e-validation.ps1` as the additional QA owner. Its LOW `0` graph impact is an extractor limitation (`0` symbols/edges), not a no-consumer claim. Current source now binds that QA owner to the durable repository bundle; no further owner is authorized.
    - Historical `E6-P6D-RECOVERY2` blocker, now superseded by the current accepted closure: the exact three durable inputs and source contract were correct, while the then-current canonical build exited before runtime creation because its lane module cache lacked module bytes.
    - Historical vendor correction, now satisfied: `E6-P6D-VENDOR1` selected checked-in vendor authority and `E6-P6D-VENDOR2` proved the standard tree CGO-incomplete, requiring the accepted complete-root fixed-point overlay and exactly one guarded patch.
    - The historical pre-build gate is satisfied. `.tmp\ladybug-native`, stale/global runtime, protected/shared/quarantined cache, implicit network/package acquisition, and hidden fallback remain forbidden authorities.
    - The corrected canonical full build now produces the accepted current runtime; native/built-reader parity and the bounded target gate pass. Independent Supervisor report `E71777663D76C4989E744C544F27830E7CB535B3031E98E2450EF8975317F6A3` returns `PASS` with procedure `CLEAN`.
    - P6-D is checked and closed. Fresh current-byte detect passed on the planner-finalized candidate; exact staged equality was `1,883/1,883` with missing/extra `0/0`; isolated commit `81163e39718b94a509e41114cada224e8f269e36` exists; post-commit index is empty; Child 06A is the direct successor.
  - Acceptance:
    - Source: graph-health maps resolver outcomes with no same-invariant target-text guess; durable native inputs and all affected build call sites contain no `.tmp\ladybug-native` authority or silent fallback.
    - Runtime/UI: corrected canonical full build produces a current lane-built runtime; all three target sites show correct external or accepted capability outcomes; UI is conditional.
    - DB/data: exact target outcomes/proofs survive Graph JSON, native Ladybug, and every affected built reader; no site is both resolved and a gap.
    - Behavior test: durable-input/root propagation, missing-input fail-closed behavior, mechanical mapping, no-heuristic, three target sites, and affected regression pass.
    - Cleanup/quarantine: only the active lane's authorized cache/runtime artifacts are removed; repository-owned third-party inputs and target source/worktree are preserved; protected/shared/quarantined roots remain untouched.
    - Evidence IDs: `E6-P6D-IMPACT1`, `E6-P6D-RECOVERY1`, `E6-P6D-RECOVERY2`, `E6-P6D-VENDOR1`, `E6-P6D-VENDOR2`, `E6-P6D-SRC1`, `E6-P6D-BUILD1`, `E6-P6D-TEST1`, `E6-P6D-PARITY1`, `E6-P6D-TARGET1`, `E6-P6D-ORACLE1`, `E6-P6D-BOUNDARY1`, `E6-P6D-REVIEW1`, `E6-P6D-DETECT1`, `E6-P6D-COMMIT1`.
    - Actual-status rows refreshed: graph-health projection, durable native/build authority, current-runtime provenance, target acceptance, and affected readers.
  - Evidence Targets: impact/recovery authority/source/build/tests/parity, target oracle/boundary, Supervisor, detect, commit.
  - Actual-status Update: graph-health `wrong -> local candidate -> independent REJECT -> six-field repair independently cleared -> independent full-slice PASS`; native/build recovery `shared .tmp authority + no durable bundle -> source/native contract complete -> standard vendor incomplete -> complete-root overlay plus one guarded patch -> canonical full build/current runtime/readers accepted`; target `0/3 unproved -> accepted capability outcomes 3/3 with analyzer gaps 0/3`; procedure `not clean -> CLEAN`; Main fresh current-byte detect `pending -> PASS`; isolated commit `pending -> 81163e39718b94a509e41114cada224e8f269e36`; P6-D and Child 06 `open -> closed`; Child 06A becomes the direct successor.
  - Worker Pre-edit Record: `E6-P6D-IMPACT1` selects only `internal/graphhealth/diagnostics.go` for production edit. `internal/graphhealth/policy.go` already owns sufficient classification/actionability values and `internal/graphhealth/resolution_gap_inputs.go` copies normalized fields losslessly, so both are validate/preserve-only. The new direct test owner may be `internal/graphhealth/p6d_outcome_projection_test.go`. UI/Web remain preserve-only because fresh upstream impact contains no frontend file.
  - Worker Source/Build Record: current production bytes decode and validate the committed schema-v1 outcome plus the nested full P6-B authority proof matrix, mechanically override flat/heuristic policy, and fail closed on malformed, incomplete, identity/status/proof drift. No bounded target name is present in production. Tests were added only after production in the sole planned owner, including positive missing/rejected catalog-proof capability states. Exact repo lock was free before both builds. E-only/offline broad `go build -buildvcs=false ./...` exits `1` after `7,653 ms` on the preserved missing-module/mixed-fixture baseline; affected `go build -buildvcs=false ./internal/graphhealth` exits `0` after `11,799 ms`. `E6-P6D-SRC1` and `E6-P6D-BUILD1` are recorded; test/parity/target/oracle/boundary/detect/review/commit remain pending.
  - Worker Test/Parity Record: repaired-byte focused `go test ./internal/graphhealth -run '^TestP6D' -count=1 -v` and full `go test ./internal/graphhealth -count=1` both exit `0`. Exact direct rows are policy `3/3`, valid missing/rejected catalog-proof capability `2/2`, hostile fail-closed `26/26`, legacy `1/1`, and Graph JSON/gap/summary parity `1/1`. The six added hostile rows each drift only one nested/outer identity group and return `unclassified/review`. The three generic external/capability rows remain `standard_library/non_actionable=3/3` and `in_repo_unresolved/analyzer_gap=0/0` across Diagnostics and ResolutionGap inventory. `E6-P6D-TEST1` is recorded; `E6-P6D-PARITY1` is partial until real Ladybug/built readers and target are proven. P6-D remains unchecked.
  - Worker Blocker Record: Main clarified that network/install/package scripts are forbidden and only modules pre-existing in the exact non-blocked E-only lane cache may support a build with `GOPROXY=off`/`GOSUMDB=off`; all three P6-C3 blocker roots, shared Ladybug cache/helper output, stale binaries, and workarounds are forbidden. A network-enabled direct CLI build had already begun before correction; it was interrupted immediately, produced no binary, and all output was discarded. Downloaded cache is quarantined at `5,687 / 469,406,586` plus build cache `3,397 / 281,902,402` and may not be used. One exact lane-only cleanup attempt was policy-rejected before process creation; no retry/workaround occurred. Built-current-runtime is therefore BLOCKED, target was never accessed, and `TARGET1`/`ORACLE1`/target `BOUNDARY1` are explicitly unproved. Disposition must be BLOCKED, never READY or accepted.
  - Worker Detect Record: after the sole BLOCKED coder report was sealed, forced E-only analyze is `2,067/763/0`, `117,852/164,411/20,310`, unresolved file projections `469`. Full detect exits `0`; file layer HIGH is `5` tracked paths / `198` changed symbols (`179` backend graph-health, `19` docs), while summary independently reports affected_count `0` and risk `low`. ResolutionGap change inventory is `95` entities / `99` occurrences, runtime health impact is `0`, and semantic completeness is `117,852/117,852`. The new test/report are untracked and therefore absent from detect's tracked-file inventory but remain in the explicit seven-path manifest. `E6-P6D-DETECT1` is recorded without any commit authority.
  - Independent Supervisor Reject Record: `E6-P6D-REVIEW1` is REJECT in `reports/Supervisor/rp_supervisor_260823_044455_by_gpt-5_p6d_graph_health_projection_and_target_proof.md` (`19,281` bytes / `196` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `D015CBD4847261DD50C9033DFA487657629910B296041744EF0BCAF5E108F2AE`). A hostile probe proves `6/6` false acceptances when nested authority alone drifts from the outer outcome on `filePath`, `fileHash`, complete range, `siteKind`, `requestedName`, or `requestedMeaning`; required result is `unclassified/review`. Repair ownership remains inside P6-D `diagnostics.go` plus six durable rows in the existing P6-D test owner. The same omission in committed P6-C3 is explicitly out of current repair authority and must not be silently edited. Built-current-runtime, native Ladybug/built-reader parity, target `0/3`, oracle, and target pre/post boundary remain independently blocking.
  - Reject-Only Repair Record: fresh forced graph before edit is `2,070/763/0`, `117,891` nodes / `164,450` relationships / `20,310` dependency edges. `validStructuredResolutionAuthority` is LOW `3/1/1/0`, but `diagnostics.go` containing-file impact remains CRITICAL `107` impacted symbols / `31` files / `34` direct / `1` flow / `19` tests and file detail remains HIGH `154/48/70/133/186/1/19`. Production now enforces exact nested/outer equality for file path, file hash, all four range coordinates, site kind, requested name, and requested meaning. Exactly six durable generic hostile rows each mutate one otherwise valid group and PASS `6/6` as `unclassified/review`; the complete hostile matrix is `26/26`. Holder checks are clean; repaired-byte broad offline build remains truthfully BLOCKED/FAIL after `7,156 ms`; affected build PASSes after `12,374 ms`; focused and full graph-health regressions PASS. Repaired-code analyze is `2,070/763/0`, `117,892/164,467/20,310`; detect exits `0`, file layer HIGH `5` tracked files / `200` changed symbols, summary risk `low` / affected_count `0`, ResolutionGap `96/100`, health `0`, semantic completeness `117,892/117,892`. Exact lane dead-cache cleanup was policy-rejected before process creation and remains BLOCKED at `1,504` files / `64,115,238` bytes / `18` zero-byte files; no retry/workaround occurred. The committed P6-C3 omission remains out-of-authority follow-up only. Independent re-review, current final record-byte graph/detect, built runtime, Ladybug, target, oracle, boundary, and commit remain pending/blocked; P6-D stays unchecked and disposition stays BLOCKED/REJECT-READY.
  - Independent Resubmission Review Record: `E6-P6D-REVIEW1` remains REJECT in `reports/Supervisor/rp_supervisor_260823_064551_by_gpt-5_p6d_graph_health_projection_resubmission.md` (`23,068` bytes / `199` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `94925BB382FD31556D740F0A052E6CA1C09365960BBEEC338AE2493707E1861D`). The review independently clears the exact six equality groups, six isolated hostile rows, complete `3+2+26+1+1` matrix, affected build/regression, and generic Graph JSON/ResolutionGap/summary parity. Full build and current CLI build remain offline FAIL/BLOCKED; no runtime binary, native Ladybug, built readers, target `3/3`, oracle, or target pre/post proof exists. Three pre-guard internal lanes were interrupted without outputs being read, but their negative activity cannot be fully audited; analyze-lock inspection also exposed forbidden launcher metadata without following it. The review is therefore not procedurally clean. The committed P6-C3 omission remains routed out of authority. P6-D stays unchecked, unstaged, and uncommitted.
  - Owner Recovery Correction Record: `E6-P6D-RECOVERY1` records sealed blocker report `reports/coder/rp_coder_260823_171110_by_gpt-5_p6d_real_recovery_blocked.md` at SHA-256 `D335ADD9EF4B87928A5F5C9ED769A9FBEE91AAC8B7CB8424AA75C32BB93BBC1A` and supersedes only its proposed smallest-next-action boundary. P6-D now owns materializing the durable repository bundle and correcting the canonical/helper/package/launcher root contract inside this same slice. It does not reopen prior semantic clearances, check P6-D, authorize target access before a current runtime, or create a new gate/slice/commit.
  - Recovery Owner Extension Record: fresh current graph and source evidence proves `scripts/qa-child02-p2e-validation.ps1` is a same-invariant native-reader QA consumer because it still binds `$LadybugRuntime` to `.tmp\ladybug-native\v0.19.1\windows-x86_64`. Its PowerShell graph row has `0` symbols/edges and LOW `0` impact because no ScopeIR call graph is extracted; that limitation does not erase the direct source dependency. This four-ledger refresh authorizes only that exact additional owner after the production/build contract is correct.
  - Recovery Coder Blocked Record: `E6-P6D-RECOVERY2` is the Main-verified Coder report `reports/coder/rp_coder_260823_201156_by_gpt-5_p6d_native_canonical_build_recovery_blocked.md`, sealed at `27,364` bytes / `315` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `7019234AE8BBE030CADEA73B91760001EC9E9089C6B7A775D23CBE12BEF58A22`. Exact durable bundle and official release/asset/license provenance are verified; source/contract/native-smoke gates pass. Corrected canonical build exits `1` after `21,577 ms` with `17` zero-byte module lookup locks, `0` module bytes, and no package runtime. Final current-source graph is `2,101/765/0`, `122,163/168,763/20,335`, unresolved projections `479`; detect exits `0`, overall MEDIUM/file HIGH, affected `3` symbols / `8` files / `3` processes, changed `407` symbols / `13` files, ResolutionGap `206/206`, health `0`, semantic completeness `122,163/122,163`. Two excluded reserved-variable incidents prevent a clean procedure claim. Target remains forbidden/unaccessed. The report's embedded old Main task ID is a routing typo superseded by current orchestration authority and does not alter the technical blocker.
  - Vendor Architecture Decision Record: `E6-P6D-VENDOR1` is `reports/system-architect/rp_system-architect_260823_210034_by_gpt-5_p6d_clean_e_only_go_vendor_closure.md`, Main-verified at `20,868` bytes / `300` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `C2A20A72F94A971317E54A4DF77973E5CA4DC8516221FB72D43DD589BC75C776`. It selects one final mechanism: exact checked-in `vendor/`, explicit `-mod=vendor`, immutable `third_party/go-vendor` integrity/license authority, and no cache/network fallback. Residual architecture questions are none; current allowed E-only surfaces contain no module source bytes, so materialization remains physically blocked before Coder edit.
  - Vendor CGO Closure Correction Record: `E6-P6D-VENDOR2` is `reports/system-architect/rp_system-architect_260823_235012_by_gpt-5_p6d_standard_vendor_cgo_closure_correction.md`, Main-verified at `32,260` bytes / `424` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `62ADCD71F170E19636A47E370A17298312C3D63E7C6ADDE0E77E3B9FAED3B7DE`. Standard vendor `2/2` remains the baseline, but verifier PASS followed by offline compile FAIL proves it omits the CGO source closure. The approved correction is checksum-derived complete-root fixed-point overlay plus exactly one guarded `tree-sitter-go` patch, unchanged `go.mod`/`go.sum`, manifest closure contract v2, dual-shell fail-closed verification, and offline compile as the separate closure proof. Fresh authoring graph is `2,121/765/0`, `122,371/168,985/20,335`, unresolved `479`; roadmap and all four ledgers are LOW with `0` impacted files/flows/tests.
  - Current Runtime/Reader/Target Acceptance Record: Coder report `reports/coder/rp_coder_260824_054528_by_gpt-5_p6d_runtime_reader_target_proof_ready_for_main.md` is sealed at `13,436` bytes / `176` LF / SHA-256 `E79257D64788BA31DC9B7D9FE010E893A03EDC7CF732137A40EF35089AC50EC3`. It proves the canonical PowerShell 5.1 full build, runtime `1.2.8` (`73,749,504` bytes / SHA-256 `64EECEBAF459F23EA11DA38B29C5E290F6F6B22D5EE4749E068BB927A8EEDEBB`), reader gate, exact target outcomes, and preservation boundary.
  - Independent Acceptance Record: `reports/Supervisor/rp_supervisor_260824_063140_by_gpt-5_p6d_runtime_reader_target_acceptance.md` returns `PASS`; identity is `24,954` bytes / `240` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `E71777663D76C4989E744C544F27830E7CB535B3031E98E2450EF8975317F6A3`. Accepted results are reader parity `53,156/53,156` with zero structured differences; target analyze `1,359/887/0`, graph `95,762/129,716`; `Promise`, `Math.max`, and `Math.min` accepted capability outcomes `3/3`; in-repository analyzer gaps `0/3`; target source/config preservation `203/203` and `3/3`; procedure `CLEAN`.
  - Main Current-Byte Detect Record: repo-owned runtime `1.2.8` (`73,749,504` bytes / SHA-256 `64EECEBAF459F23EA11DA38B29C5E290F6F6B22D5EE4749E068BB927A8EEDEBB`) completes one fresh forced self-analyze at `2,170/765/0`, graph `122,649/169,300`, file projection `20,351` edges / `479` unresolved. Full-scope detect exits `0`: summary LOW with `affected_count=0`; file layer HIGH with `13` changed / `8` affected tracked files; `501` changed semantic units split backend `436`, backend-test `16`, docs `49`; ResolutionGap `245/245`, all `unclassified/review`; health `0/0/0`; app-layer and functional-area completeness `122,649/122,649`. The first wrapper attempt stopped before directory creation or process start because PowerShell rejected assignment to reserved `$HOME`; it is excluded and produced no graph state. This ledger record does not trigger a self-referential rerun.
  - Binding `.tmp` Disposition: every path below `E:\Anvien\.tmp` is disposable reproducible scratch and is never evidence, review input, handoff authority, source of truth, acceptance dependency, or procedure blocker. Historical scratch incidents remain history only and do not qualify the current clean verdict.
  - Residual Routing: committed P6-C3 `validateResolutionOutcome` six-field parity and non-Windows `scripts/ensure-ladybug-native.sh` remain explicit out-of-scope residuals; neither invalidates the accepted Windows P6-D boundary or authorizes expansion here.
  - Commit Record: exact staged equality was `1,883/1,883`, missing/extra `0/0`, with `1,871` additions and `12` modifications. The protected outside boundary remained `25/25` and included Owner deletion `CONTRIBUTING.md`. Isolated commit `81163e39718b94a509e41114cada224e8f269e36` (`feat: close P6-D runtime and reader target proof`) has parent `741335f7efa5ad215ec28361d88c462f431565ab`, tree `b6bc216172886d5de9ef80f82c0da3ed28880f13`, exactly `1,883` paths, `5,866,720` insertions / `314` deletions, empty post-commit index, and no push. Fresh post-commit graph is current/up-to-date at `2,170/765/0`, `122,649/169,300`, file projection `20,351/479`.

## Child 06 Closure And Child 06A Handoff

- Child 06 is `CLOSED` at accepted P6-D commit `81163e39718b94a509e41114cada224e8f269e36`.
- Former P6-E goal, scope, requirements, acceptance, checklist, evidence/benchmark execution, aggregate performance review, cleanup, commit, and final Child 07 handoff responsibilities are transferred to the independent Child 06A standard plan set.
- Child 06 has no pending `Pn-A`, `Pn-B`, or `Pn-C` gate. Those former performance-dependent gates cannot block Child 06A.
- Direct successor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`.
- Child 07 follows only the Child 06A closure commit and no longer opens directly from Child 06.

## Risk Notes

- `buildWorkspace`, `resolveCall`, and `resolveTypeAnnotation` have CRITICAL impact; `classifyDiagnostic` has HIGH impact. Current evidence spans resolver/analyze flows, multiple languages, graph-health, and many tests.
- A declaration-authority mechanism can affect packaging, runtime dependencies, licenses, determinism, performance, and configuration semantics. P6-A must measure and decide these before code.
- Hardcoding the three names can pass target checks while leaving the defect class intact; general fixtures and source inspection must reject that shortcut.
- Loading too broad a declaration set can inflate graph/index size or blur repository ownership. Representation and materialization remain conditional on actual consumers.
- A graph-health-only change can hide a resolver gap. P6-D cannot close unless the underlying P6-C3 outcome is present and persisted.
- Project/package declaration lookup is preserve-only under the P6-A evidence. Any activation requires a planner refresh, new current evidence, and review before code.
- Any newly discovered affected consumer expands only the owning slice after plan/actual-status refresh; it does not authorize unrelated command/API changes.
