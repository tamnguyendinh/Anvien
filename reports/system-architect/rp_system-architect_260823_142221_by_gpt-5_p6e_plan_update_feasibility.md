# System Architect Review — P6-E Plan Update Feasibility

- Created: `2026-08-23 14:22:21 +07:00`
- Exact product title: `P6-E: Accelerate Analyze Without Sacrificing Accuracy`
- Verdict: `BLOCKED`
- Send to: `Main`
- Acceptance authority: this is architecture feasibility/reasonableness review only; it is not Supervisor acceptance and authorizes no implementation.
- Next lane: `Main`; Supervisor must not open on the current five-file bytes. Main must route the two minimal Planner corrections below first.

## Claim, authority, and scope

This review answers one bounded question: whether the actual corrected five-file plan update is coherent and executable as one P6-E checklist slice that improves end-to-end `anviens analyze` performance without sacrificing accuracy, semantic completeness, graph correctness, deterministic output, freshness, failure/publication behavior, or persistence/reader parity.

Fresh raw reads through EOF were completed for:

1. `E:\Anvien\AGENTS.md`
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`
3. `E:\Anvien\.agents\skills\System-Architect\SKILL.md`
4. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`
5. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
6. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
7. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
8. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`

No source, plan, ledger, Git state, graph, runtime, target, benchmark, test, build, cache, or existing report was mutated. The only created artifact is this report. No runtime measurement or acceptance verdict is inferred from document wording.

## Executive verdict

`BLOCKED` applies to the current plan-document boundary, not to the underlying performance strategy. The product goal, optimization order, correctness contracts, measurement protocol, and P6-D → P6-E → Pn-A placement are architecturally reasonable. However, two exact cross-file contradictions make the current five-file authority unsafe to execute or send to Supervisor:

1. two internal rebaseline gates require intermediate U2/U3 commits even though P6-E is explicitly one checklist slice with one final isolated commit;
2. the active P6-D prerequisite identity is the amended report, while active plan/status wording still presents the pre-amendment seal as current.

Both are narrow Planner corrections. No architecture redesign, new method, new slice, or implementation is required to clear this review.

## VERIFIED clearances

### 1. Product outcome is separated from methods

- The roadmap defines Child 06's required result as faster end-to-end `anviens analyze` with correctness, semantic, graph, determinism, freshness, failure/publication, and persistence/reader parity: roadmap lines `45`, `82`, and `167`.
- The plan states the same product outcome at lines `19`, `66`, `520`, and `596-598`.
- The evidence ledger explicitly classifies cost map, A/A+A/B, canonical-path reuse, all-import index, accumulator, decoder, and Pareto rerank as methods/work steps rather than the product goal: evidence line `612`.
- The benchmark ledger repeats that the granular cost map is attribution for the end-to-end goal: benchmark line `13`.

Clearance: the corrected wording no longer treats the cost map as the goal.

### 2. Placement and lock topology are coherent

- Roadmap line `58` places P6-E after open `REJECT/BLOCKED` P6-D and locks it behind P6-D PASS, detect, cleanup disposition, and commit.
- Plan lines `127`, `516`, `519`, `591`, `604`, and `611` make P6-D the sole active slice, keep P6-E implementation unauthorized, and keep Pn-A closed until P6-E's final commit.
- Evidence lines `612`, `627`, and `648` preserve the same ordering.
- Actual-status lines `5`, `115`, `162`, `236-238`, and `269` agree that P6-D is sole active `REJECT/BLOCKED`, P6-E is declared/locked, and Pn-A cannot open.

Clearance: topology is `P6-D -> P6-E -> Pn-A`; placement itself is not reopened by this review.

### 3. Intended slice and rollback architecture is sound

- Plan lines `129-131` define an ordered sequence within one slice, exact per-unit equivalence/resource gates, and rollback on any mismatch.
- Plan line `539` says explicitly that method steps are rollback gates, not separate slices or commit boundaries.
- Plan lines `583` and `604` define one final isolated P6-E commit after all required and evidence-backed conditional work.
- Evidence line `642` declares exactly one isolated slice commit, and actual-status lines `237` and `269` call P6-E one checklist slice with no commit yet.

Clearance: the intended architecture is one slice/one final commit. Blocker B1 below is a ledger wording contradiction against this otherwise coherent contract.

### 4. Optimization contracts preserve correctness

- Canonical-path reuse must preserve scan order, early exit, import claims, ordered outcomes, graph, persistence, and path edge cases: plan lines `547-552`.
- The all-import index is immutable, run-scoped, includes resolved/unresolved imports and original order, is lookup-only, and is never iterated for output: plan lines `553-558`.
- Diagnostic aggregation preserves first duplicate, the full stable sort, immediate materialization, missing-node behavior, six-field fail-closed validation, second-writer/property-drift abort, reader parity, failure, and publication behavior: plan lines `565-570`.
- The decoder is conditional on fresh materiality and must preserve missing/null/type/malformed/unknown/duplicate-key/nested-raw behavior plus the exact `(outcome, structured, valid)` triple: plan lines `571-576`.
- Later DB/snapshot/parse/post-run owners may open only from fresh absolute evidence and cannot reduce rows, properties, proofs, fallback, transaction, snapshot, workload, output, schema, or format: plan lines `577-581`, `595-599`.
- Workload reduction, skipped validation/persistence, stale or cross-run caches, unmeasured concurrency, GC substitution, target access, output/schema changes, and unsupported speed promises are forbidden at plan lines `524-526`.

Clearance: each proposed method has an accuracy/parity/failure boundary and a unit rollback trigger.

### 5. Benchmark protocol is comparable and fail-closed

- Profile attribution is separated from comparable wall timing; only alternating same-condition unprofiled A/B after A/A noise can support acceptance: benchmark lines `16-21`.
- Pair identity freezes repository/dirty/corpus/runtime/native/dependency/options/environment/cache/storage/instrumentation state, and identity drift invalidates the pair: benchmark line `71`.
- Time accounting keeps residual explicit and forbids double counting; work denominators and output/hash/digests must match; Owner decides materiality/resource rules after baseline: benchmark lines `72-78`.
- A/A, comparable baseline, every unit A/B, fresh rebaselines, Pareto rerank, and final end-to-end result have distinct locked rows: benchmark lines `84-95`.
- Profile samples are not a speedup promise, and no A/A, comparable baseline, A/B result, resource/materiality rule, equivalence result, or achieved speedup exists yet: benchmark lines `144-145`; actual-status lines `115` and `269` agree.

Clearance: historical/profile-only data cannot pass P6-E. Current result state is correctly `NO EVIDENCE`, not failure or success.

### 6. Five-file counts are consistent

- Roadmap lines `45` and `47` record Child 06 `7` slices and campaign total `36`.
- Roadmap lines `63` and `121-123` retain `34` active documents and the `36`-slice integrity gate.
- Actual-status line `162` records the transition `6 -> 7`, `35 -> 36`, with active documents still `34`.

Clearance: the required `7 / 36 / 34` counts agree.

## BLOCKED findings and exact minimal corrections

### B1 — Intermediate-commit wording creates an impossible internal gate

Evidence:

- The governing plan says each method is an internal rollback gate, not a commit boundary: plan line `539`.
- It defines only one final P6-E commit after all method gates: plan lines `583` and `604`.
- Evidence line `642` also defines one isolated slice commit.
- Contradicting those rules, evidence line `633` locks `E6-P6E-REBASE1` behind a `U2 commit`, and evidence line `636` locks `E6-P6E-REBASE2` behind a `U3 commit`.

Why blocking:

- If interpreted literally, Work Step 4 cannot start without a U2 commit and Work Step 6 cannot start without a U3 commit, but those commits are forbidden by the one-slice topology until final Work Step 8.
- Creating the intermediate commits would violate the declared slice/commit topology; not creating them deadlocks the ordered work.

Exact minimal Planner correction:

1. At evidence line `633`, replace `pending/locked behind U2 commit` with `pending/locked behind accepted and externally sealed U2 internal equivalence/rollback gate`.
2. At evidence line `636`, replace `pending/locked behind U3 commit` with `pending/locked behind accepted and externally sealed U3 internal equivalence/rollback gate`.
3. Do not add U1/U2/U3/U4 commit IDs, roadmap slices, or intermediate Git boundaries. Retain only `E6-P6E-COMMIT1` at the final slice boundary.

### B2 — Active P6-D prerequisite identity still points to the superseded seal

Evidence:

- Actual-status line `28` forbids preserving superseded assumptions as active authority.
- Evidence line `620` names the current amended P6-D report identity: `33,046` bytes / `239` LF / SHA-256 `3837059F852CF44F20732F2BB0F7D88234BC6612AAB9C7CD439AEC3288A768DA`; it also states that the old seven-path aggregate is no longer current.
- Plan line `515` still says the same report path `remains REJECT` using the pre-amendment identity `23,068` bytes / `199` LF / SHA-256 `94925BB382FD31556D740F0A052E6CA1C09365960BBEEC338AE2493707E1861D` without labeling that identity historical.
- Actual-status line `283` is the active decision note but still cites `94925...` as the report returning the current full-slice REJECT and calls the pre-plan-update graph/detect figures current.
- Actual-status line `161` and evidence line `601` may retain `94925...` only as the dated historical pre-amendment R37/review record.

Why blocking:

- `E6-P6E-ANCHOR1` is a zero-trust prerequisite gate. Active plan/status authority cannot identify its prerequisite by two different raw report identities.
- The semantic disposition remains REJECT/BLOCKED, but exact current bytes determine which review, detect, cleanup, and later commit evidence may be reused or must be refreshed.

Exact minimal Planner correction:

1. Mark the `23,068` / `199` LF / `94925...` references at plan line `515`, evidence line `601`, and actual-status line `161` explicitly as `historical pre-amendment seal`.
2. Update the active actual-status decision note at line `283` and any active P6-D anchor wording to the amended identity already recorded at evidence line `620`.
3. Relabel `2,071/763/0`, `117,906/164,481/20,310`, and its detect figures as the last pre-plan-update review measurement, not current-byte proof. Preserve the existing requirement for a fresh current-byte seal/detect/cleanup/PASS/commit before P6-E opens.
4. Do not change P6-D's substantive disposition: it remains unchecked, unstaged, uncommitted, and `REJECT/BLOCKED`.

## Five-file consistency summary

| File | Clearance | Blocking inconsistency |
|---|---|---|
| Roadmap | Correct outcome, `7/36/34`, P6-D -> P6-E ordering, and campaign completion contract | None found |
| Child 06 plan | Correct product goal, one-slice intent, ordered optimization contracts, rollback, final commit, and Pn-A lock | B2: pre-amendment P6-D seal is presented without historical qualification at line `515` |
| Child 06 evidence | Correct outcome-vs-method wording, bounded causal claims, pending evidence IDs, final commit, and current amended anchor | B1 at lines `633`, `636`; historical seal at line `601` needs explicit qualification because line `620` is current |
| Child 06 benchmark | Correct attribution/wall separation, identity protocol, A/A+A/B rows, output/resource contracts, and no speedup claim | None found |
| Child 06 actual status | Correct locked state, topology, `7/36/34`, implementation prohibition, and NO-EVIDENCE state | B2: active decision note line `283` uses the superseded report identity/current-measurement wording |

## Residual unknowns — NO EVIDENCE, not additional plan blockers

- No comparable P6-E A/A distribution, alternating unprofiled baseline, per-unit A/B result, Owner-approved materiality/resource rule, deterministic replay, freshness invalidation/restore, failure/publication equivalence result, or achieved speedup exists: benchmark lines `86-95`, `145`.
- Exact P6-E file/symbol impact and later Pareto owners remain pending until P6-D is accepted and committed: evidence lines `626-630`, `639`.
- Root Cause exact invocation multiplicities and historical regression attribution remain unverified: evidence line `616`.
- P6-D current runtime, native/built-reader, target `3/3`, oracle, boundary, clean review, detect, cleanup, and commit gates remain blocked/pending: evidence lines `627` and plan line `591`.

These unknowns are correctly fail-closed by the plan. They do not justify a speedup promise or implementation authorization, and they do not require a new architecture choice before the two documentation contradictions are corrected.

## Required routing

1. Main receives this `BLOCKED` report and exact external seal.
2. Planner applies only B1 and B2 to the existing five-file authority; no code, benchmark execution, target access, topology change, or new performance method is needed.
3. Architect performs a bounded current-byte correction verification.
4. Supervisor is the next lane only if that verification returns `FEASIBLE`.

## Durable artifact identity protocol

- Report path: `E:\Anvien\reports\system-architect\rp_system-architect_260823_142221_by_gpt-5_p6e_plan_update_feasibility.md`
- Encoding contract: strict UTF-8, no BOM, LF-only.
- The exact full-file bytes/LF/CR/SHA-256 identity is measured after this report is frozen and supplied as the detached Main handoff seal. It is intentionally not embedded into the hashed file because doing so would make the raw SHA-256 self-referential.
- Commit: `N/A / prohibited by latest Owner authority`; the report remains uncommitted.
