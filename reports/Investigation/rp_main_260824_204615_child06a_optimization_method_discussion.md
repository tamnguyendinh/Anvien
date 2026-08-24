# Detailed Child 06A Analyze Optimization Execution Method

Status: `ACTIVE_BINDING_METHOD_RECORD_FOR_CHILD_06A_REWRITE`

Scope: `Child 06A — make the real anvien analyze pipeline faster by measuring real runtime work, optimizing the largest measured cost, and measuring again`

This document records the detailed method concluded in the direct user/Main discussion. The direct rewrite command has now been issued to the visible Planner, so this report is the durable method authority for that rewrite. It is not implementation permission, an acceptance report, or a substitute for the standard four-file plan set. Its purpose is to prevent any later session or auto-compaction from losing the agreed execution method.

## 1. Finding about the current Child 06A plan

The current Child 06A plan is technically detailed but methodologically wrong for performance optimization.

It currently treats the following as the execution spine:

```text
construct and seal a harness/control system
-> construct a broad cost-map schema
-> run a fixed A/A protocol
-> obtain a separate numeric seal
-> freeze a predefined baseline/candidate system
-> execute predetermined U1/U2/U3/U4 work
-> maintain many run/capture/harness/candidate/checkpoint identities
```

That structure makes measurement serve a predetermined plan. Real optimization requires the opposite: measurements determine what work belongs in the plan next.

The current technical constraints can describe how to preserve correctness if a measured owner is edited. They cannot preselect which owner, phase, function, or optimization must be implemented.

Because the goal and execution method are wrong, the plan should be rewritten around the empirical optimization loop rather than incrementally patched around the current fixed-unit and harness-heavy control layer.

## 2. Correct product goal

The goal is not to complete U1, U2, U3, U4, a harness, a cost-map document, or a fixed list of technical ideas.

The goal is:

> Reduce the measured elapsed time of the real `anvien analyze` pipeline on the same workload, while preserving the accepted output, correctness, persistence, reader, deterministic, freshness, failure, transaction, temporary-file, and publication behavior.

Optimization is demonstrated only by measured before/after runtime improvement.

Elapsed wall-clock time is the primary and controlling metric throughout Child 06A. The product goal is always to reduce total graph-generation time as far as the measured system can safely be optimized. CPU, RAM, allocation, GC, I/O, wait, call counts, bytes, and work denominators are secondary diagnostic or comparability evidence. They may explain why elapsed time is high and guide a solution, but they cannot rank a bottleneck by themselves, cannot replace direct elapsed-time measurement, and cannot prove an optimization. If secondary metrics improve but total graph-generation time does not decrease, the candidate is not a retained speed optimization.

Every opened bottleneck implementation attempt must inherit that product goal in concrete form: reduce the measured elapsed time owned by the benchmark-selected bottleneck and reduce total graph-generation elapsed time from the current accepted baseline while preserving accuracy/equivalence. A bottleneck attempt cannot have harness construction, instrumentation, profiling, reporting, audit, cleanup, or architecture documentation as its goal. Those activities are permitted only as the smallest supporting steps needed to measure, decide, implement, or validate the time reduction.

Example:

1. The real graph-insertion operation is measured on the accepted workload.
2. Its elapsed time and work denominator are recorded.
3. The exact graph-insertion cause and owner are investigated.
4. Production code is optimized.
5. The same operation is measured again on the same workload.
6. If graph insertion becomes faster, end-to-end analyze also retains the gain, and output/lifecycle behavior remains equivalent, the change is an optimization.

Writing a harness, report, schema, profile parser, checkpoint table, or audit document is not an optimization because none of those activities reduce the elapsed time of the real analyze pipeline.

## 3. The only valid execution loop

```text
run the real pipeline
-> measure real phase elapsed times
-> record the measurements immediately
-> build the complete top-level bottleneck list ordered by absolute elapsed time
-> process the largest remaining top-level bottleneck first
-> measure deeply inside that parent and build its complete child-bottleneck list ordered by elapsed time
-> process the largest remaining child first, then every smaller child in order; no measured child is skipped because its time is small
-> open a fresh visible Architect invocation for this optimization attempt
-> Architect decides the exact technical direction and preserved boundary
-> Planner updates the current P2-A attempt from that decision
-> Coder implements only the updated attempt
-> build
-> rerun the same workload
-> compare the same phase before/after
-> compare end-to-end before/after
-> visible Supervisor independently verifies accuracy/equivalence and affected lifecycle invariants
-> if faster and Supervisor PASS: KEEP and promote the candidate to the current baseline
-> if Supervisor REJECT or no gain can be retained: restore the last accepted baseline and record one unsuccessful attempt
-> for another attempt on the same bottleneck: return to a fresh Architect decision, Planner refresh, Coder, measurement, and Supervisor
-> after three consecutive attempts without KEEP: record SYSTEM_CHARACTERISTIC for that child's remaining cost
-> continue through every remaining measured child bottleneck
-> when all children of the selected parent are terminal, remeasure the parent and whole pipeline
-> continue through every remaining top-level bottleneck, largest first
```

The loop is dynamic, but it is exhaustive rather than winner-only. No phase, file, function, optimization idea, or owner is active before current measurements select it. Every measured bottleneck is eventually processed; absolute elapsed time controls order only, with the largest processed first and smaller rows following rather than being discarded. The unit that requires architecture and review is each production optimization attempt, not each unique bottleneck. Pure measurement/inventory does not open Architect or Supervisor.

## 4. No predetermined optimization candidates

Before measurement there is only the current runtime and its real work. There is no optimization candidate list.

U1 canonical-path reuse, U2 import indexing, U3 diagnostic accumulation, U4 decoding, resolver work, diagnostics, graph construction/insertion, DB, snapshot, parse, and post-run work are not mandatory steps and are not pre-ranked candidates.

They may be historical hypotheses. A historical hypothesis never becomes work by name; only a real operation measured into the current bottleneck list is processed, when its elapsed-time order is reached.

Consequences:

- If graph insertion is the largest measured phase, graph insertion is investigated first.
- If resolution is small, it is not promoted ahead of larger measured rows merely because U1/U2 were previously written down; it remains in the measured queue and is processed later in elapsed-time order.
- If diagnostic work is small, fixed U3 is still not implemented by name; the real measured diagnostic row remains in the queue and is processed later with the same evidence-driven method as every other row.
- If a measured phase becomes small after one retained optimization, its rank is updated but the row remains mandatory queued work until its internal child list is fully processed and terminal.
- Old percentages, profile samples, reports, or solution architecture cannot override a fresh absolute timing ranking.

## 5. First action: measure the real pipeline immediately

The executor does not start by building a measurement framework. It starts by running the accepted real `anvien analyze` command on the accepted workload. P1-A records the exact options, workload, cache regime, and runtime identity from that real invocation; this method does not invent or freeze an unmeasured command tuple.

Only the minimum comparability identity is recorded before the run:

- exact runtime/executable used;
- exact command and options;
- exact repository/input workload;
- cache regime relevant to the run;
- start time, end time, exit code;
- output/work denominator needed to confirm the same work ran.

This preparation should take minutes. It must not become an independent implementation project.

The first timing map records real pipeline operations such as:

| Operation | Required primary value | Required work denominator |
|-----------|------------------------|---------------------------|
| total analyze | elapsed wall time | files/workload identity |
| scan | elapsed wall time | files/bytes scanned |
| parse | elapsed wall time | files/bytes parsed |
| resolution | elapsed wall time | facts/claims/lookups processed |
| graph construction/insertion | elapsed wall time | nodes/relationships/properties inserted |
| Graph JSON/persistence | elapsed wall time | records/bytes written |
| Ladybug/native persistence | elapsed wall time | records/bytes/transactions written |
| graph-health/diagnostics | elapsed wall time | diagnostics/nodes processed |
| post-run/publication | elapsed wall time | outputs finalized/published |

Use the real phase boundaries that exist in the current pipeline. Do not invent a phase because a plan or report previously named it.

The mandatory primary proof is elapsed wall-clock time for a real operation and elapsed wall-clock time for the complete graph-generation pipeline. Bottleneck rank is determined by current absolute elapsed time. CPU, RAM, allocation, GC, I/O, wait, call-count, byte-count, and denominator measurements are added only to explain why the selected operation is slow or to prove that before/after runs performed comparable work. They never replace time as the optimization result.

The mandatory output of this first measurement is a complete ordered bottleneck optimization list written directly into `benchmark.md`. P1-A cannot close and P2-A cannot open from a prose summary, historical profile, or a single hand-selected target. The benchmark list must contain every measured current operation in descending absolute elapsed-time order. Each row records at least rank, operation/bottleneck, current elapsed time, work denominator, percentage or delta when meaningful, proven owner/call path when available, processing/terminal state, and exact evidence. The first unprocessed row is processed first, but every later smaller row remains mandatory work. A small elapsed time is never a reason to omit or terminalize a measured bottleneck; accumulated improvements across many small rows contribute to the total graph-generation reduction.

## 6. How to measure a selected slow operation

After the first timing map, sort by absolute elapsed time. Select the largest real operation.

Selecting that top-level operation starts the processing of that parent bottleneck. The first mandatory work inside the selected parent is to measure more deeply within its own real boundary and create a complete list of its measured child bottlenecks, ordered by absolute elapsed time. The largest child is processed first, followed sequentially by every remaining smaller child; the process does not optimize only the largest child. This drill-down is part of processing the selected parent bottleneck; it is not a separate phase, a parallel discovery project, or optional work deferred until after architecture. Architect must not choose an optimization direction from only a coarse parent name or from an incomplete child inventory.

For that operation only:

1. Identify its actual start and end boundary in the current execution path.
2. Record the complete call path and actual owner.
3. Record the work denominator so before/after runs perform the same work.
4. Add only the minimum missing instrumentation needed to separate its internal costs.
5. Rerun the real command.
6. Record and rank the complete measured child-cost list of that operation.
7. Select the largest unprocessed child, prove its cause and actual owner/call path, and use that exact child evidence as the Architect input packet.
8. After that child reaches a terminal retained state, continue to the next-largest remaining child until every measured child row is terminal. Never discard a child solely because its elapsed time is small.

Example for graph insertion:

- measure total graph-insertion elapsed time;
- record node, relationship, property, transaction, and byte counts;
- if needed, separate node insertion, relationship insertion, index work, property writes, transaction commit, and serialization;
- list the elapsed time of every real sub-operation found inside graph insertion;
- optimize the largest child first, then node insertion, relationship insertion, index work, property writes, commit, serialization, and every other measured child in descending current-time order until all are terminal;
- accumulate the retained savings from large and small children into the parent and end-to-end result.

Do not collect every possible metric for every subsystem upfront. Instrumentation exists to answer the next concrete optimization question, not to build a permanent observability project.

## 7. RAM, CPU, allocation, GC, I/O, and wait evidence

Supporting resource measurements are used only when they explain the selected elapsed-time cost or when memory itself is an explicit optimization goal.

Useful memory evidence can include:

- peak working set/RSS during the selected operation;
- live heap before and after the operation;
- allocated bytes and object count;
- retained heap growth;
- GC count, pause, and CPU contribution;
- repeated allocation caused by a measured hot loop.

Useful I/O/wait evidence can include:

- bytes and operations read/written;
- transaction or file flush time;
- lock/block/wait time;
- time spent waiting for downstream persistence.

These metrics support diagnosis. They do not replace the required proof that the selected real operation and the end-to-end graph-generation command became faster. A favorable resource or count delta without lower elapsed time is not a speedup and cannot receive KEEP.

## 8. Production optimization procedure

Once current measurement selects the next top-level parent bottleneck:

1. Begin processing the benchmark-selected parent by measuring more deeply inside its real boundary. Create and record its complete child-bottleneck list ordered by absolute elapsed time, including the smaller child costs.
2. Select the largest unprocessed child row, determine its exact cause, source owner, complete call path, and work denominator, then open an optimization-attempt record against the current accepted baseline and that child's consecutive no-KEEP count.
3. Invoke one visible Architect for this child attempt only after the complete current parent-child drill-down exists. Give it the exact parent row, complete child elapsed-time list, selected child row, work denominator, cause, owner/call path, preserved invariants, scope boundaries, and any prior rejection evidence. A prior Architect decision never authorizes another production attempt.
4. Architect decides this attempt's exact cause, technical direction, allowed owners, invariants, expected observable gain, validation boundary, resource constraints, and rollback condition. Architect does not choose bottleneck priority, rerun measurements, audit report chains, or accept implementation.
5. Main routes that decision to Planner. Planner updates the exact current P2-A attempt in the four-file plan set with concrete work steps and gates. Coder remains blocked until this fresh update exists.
6. Run the required fresh Anvien graph, file-detail, and impact immediately before editing the actual owner.
7. Record HIGH/CRITICAL as blast-radius warnings and constrain affected validation accordingly; do not treat them as edit bans.
8. Coder implements only the fresh Architect-approved and Planner-recorded production direction; Coder does not invent, reuse, extend, or substitute architecture.
9. Edit production code first, then add or update tests only after production behavior is correct.
10. Before building, terminate build-related lock holders as required by repository rules, then run the canonical full build.
11. Rerun the same workload under the same relevant conditions.
12. Measure the same operation before/after and measure end-to-end analyze before/after.
13. Open a visible Supervisor review for this attempt. Supervisor independently checks accuracy, correctness, output/persistence/reader equivalence, determinism, freshness, failure, transaction, temporary-file, publication, Graph JSON, Ladybug, native-reader, and the exact affected lifecycle boundary.
14. Only when the selected child operation is faster, its parent and end-to-end gain survive, and Supervisor returns PASS may Main record KEEP and promote the candidate to the current baseline. KEEP resets this child bottleneck's consecutive no-KEEP count to zero.
15. If Supervisor returns REJECT, the candidate cannot become baseline. Restore or retain the last accepted baseline using only the exact owned rollback boundary, record the unsuccessful attempt, send the rejection evidence back to a fresh Architect invocation, route the new decision through Planner, then let Coder start another attempt.
16. If Supervisor passes accuracy but the attempt produces no retainable direct and end-to-end gain, restore or retain the last accepted baseline and record an unsuccessful attempt rather than a speedup claim.
17. Every additional attempt on the same child bottleneck repeats Architect -> Planner -> Coder -> build/remeasure -> Supervisor. Three attempts on one child therefore require three Architect decisions, three Planner refreshes, and three Supervisor reviews. Ten measured child bottlenecks require at least ten of each, plus any repeated attempts.
18. After three consecutive attempts without KEEP from the current accepted baseline, stop optimizing that child in Child 06A and record `SYSTEM_CHARACTERISTIC`: the child's remaining measured cost cannot be optimized further while preserving the current system boundary and is accepted as an inherent system feature/cost.
19. A `SYSTEM_CHARACTERISTIC` child row remains in the numeric inventory with its denominator, retained time, three attempt references, and evidence. It is excluded from further attempts, but it does not close the parent while another measured child remains unprocessed.
20. After KEEP, remeasure the selected child, parent, and whole pipeline; refresh their numeric values and the complete child ordering. Continue processing the same child until it is terminal, then select the largest remaining unprocessed child.
21. Only when every measured child of the selected parent is terminal may the parent be marked processed. Then remeasure the parent and full pipeline, refresh the complete top-level list, and proceed to the largest remaining unprocessed top-level bottleneck.
22. Continue until every measured top-level bottleneck and every measured child beneath each selected parent has been processed to a terminal retained state. Small measured rows are never skipped because of size.

Only one shared-worktree production owner edits at a time so attribution remains clear.

## 9. Before/after comparison semantics

Each retained optimization has two comparisons.

### Direct operation comparison

Compare the same measured operation immediately before and after the change, with the same work denominator.

### End-to-end comparison

Compare the real complete `anvien analyze` command before and after the change. The direct saving must survive at the product boundary rather than disappear into another phase or resource regression.

After a KEEP decision, the retained code becomes the next current baseline.

Example sequence:

```text
current state A0
-> optimize measured cost 1
-> compare A0 versus B1
-> KEEP B1
-> B1 becomes current state A1
-> optimize newly measured cost 2
-> compare A1 versus B2
-> KEEP B2
```

Do not compare every later unit only against the original A0 when deciding its incremental value. That would mix gains from earlier work into later attribution. The original initial baseline remains useful for final initial-versus-final product comparison.

## 10. Decision rules

| Decision | Meaning | Immediate next action |
|----------|---------|-----------------------|
| `KEEP` | selected operation and end-to-end runtime improved and the per-attempt Supervisor passed all required accuracy/equivalence invariants | retain bytes, promote current baseline, reset no-KEEP count to zero, remeasure full pipeline |
| `REWORK` | the attempt is not accepted but current evidence supports another scoped solution attempt | restore/retain last accepted baseline, return rejection evidence to a fresh Architect, refresh the attempt through Planner, then run Coder again |
| `ROLLBACK` | no retained gain, regression, nondeterminism, or failed equivalence | restore only exact owned bytes to the last accepted baseline, count the unsuccessful attempt, and return through Architect and Planner if fewer than three consecutive attempts have failed |
| `SYSTEM_CHARACTERISTIC` | three consecutive attempts from the current accepted baseline produced no KEEP; remaining cost cannot be optimized further under preserved correctness/current system boundary | retain the last accepted correct baseline, keep the numeric row as an inherent system feature/cost, exclude it from further Child 06A target selection, remeasure/rerank |
| `BLOCKED` | exact required runtime, authority, dependency, or evidence is temporarily unavailable | record blocker and next real authority/action; keep the row open and unchecked because BLOCKED cannot terminalize or bypass mandatory bottleneck work |

Rollback must be scoped to the exact unit-owned diff. The process cannot rely on broad reset, checkout, stash, or cleanup. The pre-edit boundary and rollback method are identified before editing. A rejected candidate is never a baseline, and Coder never receives a rejection directly as permission to invent a correction.

## 11. Remeasurement and dynamic ranking

Every retained optimization changes the performance profile. Therefore the plan cannot keep following an old ranking.

After each child KEEP or terminal `SYSTEM_CHARACTERISTIC` disposition:

1. Run the whole real pipeline again.
2. Update each measured phase's current elapsed time.
3. Refresh the selected parent's complete child list and the complete top-level list with current accepted elapsed times.
4. Keep processing the current child until terminal; then select the largest remaining unprocessed child of that parent.
5. When all children of the parent are terminal, select the largest remaining unprocessed top-level bottleneck.
6. Discard previous percentages as priority authority, but never discard a measured row merely because it is small.

Do not rerank from a rejected candidate. Restore or retain the last accepted baseline first. A KEEP resets the current child bottleneck's consecutive no-KEEP count; each later attempt starts from the newly accepted baseline. After three consecutive no-KEEP attempts, `SYSTEM_CHARACTERISTIC` closes only that child row. Processing continues through the remaining child list, then the remaining top-level list.

The optimization loop ends only when every measured top-level bottleneck and every measured child bottleneck has a recorded terminal retained disposition. `BLOCKED` is not terminal and prevents parent/plan completion until resolved. There is no time-size materiality cutoff that permits a measured small bottleneck to be ignored.

## 12. Measurement lanes and implementation lanes

Parallelism is optional, not a goal.

The execution pool contains ten pre-opened visible measurement lanes ready for assignment. They run in bounded waves: at most three measurement lanes may be active concurrently to avoid overloading the machine. Each active lane owns exactly one measured top-level or child bottleneck problem. Remaining lanes stay waiting until a slot is released.

Use parallel read-only measurement analysis only when current evidence already defines independent measured work. The ten-lane pool is capacity, not a requirement to invent work; each wave contains one to three lanes according to actual independent measured rows.

The initial accepted P1-A runtime capture is produced by exactly one visible measurement executor. After its command/workload/runtime identity and acceptance are recorded, up to three ACTIVE read-only measurement analysts may consume that same capture for independent measured problems. The analysts do not edit source and do not launch duplicate competing benchmark processes merely to create concurrency.

Rules:

- One accepted real runtime capture may be read by multiple measurement lanes for different already-defined measured operations.
- Do not rerun the same expensive capture just to create separate reports.
- Prefer one accepted capture shared by the active wave. The three-lane concurrency cap governs active agents; it does not justify simultaneous duplicate benchmark processes whose CPU/RAM/disk/cache contention would invalidate elapsed-time comparability.
- Lane count is determined by actual independent measured work, not by examples or filenames.
- A completed measurement frees one of the three active slots only after its numeric result is recorded in benchmark, its proof is recorded in evidence, the matching plan checklist/status pointers are updated, and Main has the information needed for the next decision. Then the next waiting lane may be dispatched.
- Production/test implementation on the shared worktree is sequential.
- Main owns the current ranking, dispatch, decision, and transition.

## 13. Minimal measurement tooling

Use existing product capabilities first: real elapsed timing, existing benchmark output, and existing profiles.

A harness is permitted only when it is the smallest inexpensive tool required to repeat the real command consistently.

The harness:

- does not count as optimization progress;
- does not become its own phase or product;
- does not require a separate architecture/review/acceptance chain;
- does not create a hierarchy of harness ID, capture ID, run ID, candidate ID, and checkpoint ID unless the real measurement cannot otherwise be distinguished;
- does not generate a durable report tree for every run;
- must not consume hours that should be spent measuring and optimizing the real pipeline.

If existing timing/benchmark capabilities already provide the needed value, do not create a harness.

If existing timing/benchmark/profile capability lacks the one operation timing needed to complete P1-A, Main assigns exactly one visible sequential instrumentation-writer lane. This lane sits outside the read-only measurement analysts and is the sole source writer for that conditional branch. The branch is not a new phase, implementation slice, optimization attempt, harness project, Architect decision, Supervisor gate, or progress claim.

The conditional P1-A instrumentation branch is closed as follows:

1. Before editing, refresh the required graph and record file-detail plus impact for the exact timing owner.
2. Add only the smallest instrumentation required for the missing timing; bind the exact writer, edited owner, owned bytes, and runtime/instrumentation identity.
3. Run the canonical full build and require PASS before any instrumented timing may be used, compared, or ranked.
4. Produce the accepted instrumented capture through one visible measurement executor, then record instrumentation overhead, denominator/comparability, and output equivalence. Before/after timing is valid only when instrumentation state and work are like-for-like.
5. End the branch with exactly one disposition: either carry the instrumentation, with exact ownership, into the first refreshed P2-A attempt; or remove the exact owned bytes, run the canonical full build again, and re-establish and remeasure the accepted timing basis before ranking or transition. No half-removed, unbuilt, mixed-instrumentation state may control benchmark or actual status.

This conditional source-writing branch is recorded through `E1-P1A-INSTR1` and the matching actual-status pointer. It does not change the ten-waiting/three-ACTIVE read-only analyst pool, and it creates no separate report system or commit.

## 14. Exact use of the standard plan ledgers

The existing four-file Child 06A plan set is the recording system. Do not build a second reporting system beside it.

### Benchmark ledger

`benchmark.md` is the primary numeric source of truth and the central control surface for this optimization plan. It owns both the complete top-level bottleneck list and the complete child-bottleneck list for the parent currently being processed. Order comes from elapsed time: largest first, then every remaining smaller row. The next optimization target is derived from these lists, not from plan prose, reports, historical percentages, an assumed solution list, or a size cutoff.

The living benchmark table must make the optimization loop visible in one place:

| Parent | Level | Rank | Operation / bottleneck | Work denominator | Initial baseline | Current before | Share / delta | Processing state | Attempt | Current after | Direct delta | Parent before/after | End-to-end before/after | Cumulative delta from initial | Consecutive no-KEEP | Disposition | Owner / call path | Evidence |
|--------|-------|------|-------------------------|------------------|------------------|----------------|---------------|------------------|---------|---------------|--------------|---------------------|-------------------------|-------------------------------|---------------------|-------------|-------------------|----------|

Rules:

- The first detailed measurement populates initial/current times and produces the first ordered bottleneck list.
- Absolute elapsed wall-clock time is the ranking authority and primary numeric result. Secondary CPU/RAM/allocation/GC/I/O/wait/count metrics may explain a row but cannot outrank or replace its elapsed time.
- The complete top-level ordered bottleneck list in `benchmark.md` is a mandatory P1-A completion gate. A prose-only ranking or a target named only in plan/actual-status is insufficient.
- When a parent row is selected, its complete measured child list becomes a mandatory nested control table before Architect. The child list contains every measured internal bottleneck in descending absolute elapsed-time order.
- The largest unprocessed row is first, but every smaller row remains mandatory queued work. No measured row receives a terminal disposition merely because its time is small.
- Every optimization attempt writes the same bottleneck's attempt number, new before/after values, and current consecutive no-KEEP count immediately.
- Every KEEP requires lower same-bottleneck elapsed time, lower retained end-to-end graph-generation elapsed time, and exact per-attempt Supervisor PASS evidence; it then updates the current baseline and cumulative improvement from the initial graph-generation time and resets the no-KEEP count to zero.
- Every REJECT/REWORK/ROLLBACK records the exact Architect decision, Planner refresh, Coder candidate, measurement, Supervisor verdict, rollback boundary, and next Architect invocation through evidence references; benchmark retains numbers rather than narrative.
- After three consecutive attempts without KEEP, the child row records `SYSTEM_CHARACTERISTIC`, the retained correct time and denominator, and references all three attempts.
- After remeasurement, parent and child ranks are recalculated directly from current absolute times.
- After every KEEP or `SYSTEM_CHARACTERISTIC`, direct child, parent, and full-pipeline measurements refresh the nested and top-level lists. A stale list cannot select the next target.
- The highest current unprocessed child of the active parent becomes the next child action. A new top-level parent is selected only after every measured child of the active parent is terminal.
- `actual-status.md` references the exact benchmark row for current bottleneck and next action; it does not duplicate the numeric table.
- `evidence.md` references the same row and explains why the measurement and disposition are valid.
- No separate cost-map/report tree competes with this ledger.

Record numeric measurements immediately:

- timestamp/checkpoint;
- real operation;
- exact workload denominator;
- current/before elapsed time;
- after elapsed time;
- delta and percentage when meaningful;
- supporting RAM/CPU/I/O/GC values only when used;
- current rank;
- evidence ID.

### Evidence ledger

Record why the result is valid:

- exact command and exit code;
- relevant runtime/input/options identity;
- operation boundary and denominator;
- files/owner changed;
- build/test result;
- affected equivalence result;
- fresh Architect decision and exact Planner refresh for each production attempt;
- per-attempt Supervisor verdict and, for REJECT, violated invariant plus rollback/re-entry evidence;
- KEEP/REWORK/ROLLBACK/SYSTEM_CHARACTERISTIC/BLOCKED decision;
- next action/owner.

### Actual-status ledger

Maintain current reality:

- current full-pipeline timing;
- complete current top-level ranking;
- active parent bottleneck and its complete current child ranking;
- current child row being processed and remaining child queue;
- current accepted baseline and current attempt number;
- consecutive no-KEEP count for the active bottleneck;
- current active owner;
- last disposition;
- exact next action.

### Plan file

Contain only:

- the empirical loop;
- scope and preserved invariants;
- currently open single implementation slice;
- for every current bottleneck attempt, the explicit goal of reducing that row's elapsed time and total graph-generation elapsed time from the current accepted baseline;
- the exhaustive largest-first rule: process every measured child of the active parent, then every remaining top-level bottleneck; never drop a row because its time is small;
- the per-attempt Architect -> Planner -> Coder -> measurement -> Supervisor state machine;
- the three-consecutive-no-KEEP `SYSTEM_CHARACTERISTIC` rule;
- completion criteria;
- final review/cleanup/detect/commit boundary.

The detailed rule body is stored in one auxiliary file named `plan-rules.md` in the Child 06A plan directory. The standard plan keeps its required `## Rules` heading but replaces the long duplicated rule body with an explicit relative link to `plan-rules.md` and states that the linked file is the binding rule authority for this plan.

`plan-rules.md` is an auxiliary readability split, not a fifth standard ledger, phase, gate, report, evidence source, numeric control surface, or permission boundary. It owns the full plan-wide execution rules consolidated from this method. The standard four files retain their normal responsibilities: plan controls goals/checklists/current attempt and links the rules; benchmark owns numbers/order; evidence owns proof; actual-status owns current state.

Do not store command logs, metric dumps, report histories, or fixed hypothetical solutions in the plan or `plan-rules.md`. Do not duplicate the full rule body in both files; update the rule once in `plan-rules.md` and keep the plan pointer stable.

The plan must also contain one living `Measured Bottleneck Checklist` populated only from real benchmark discoveries:

- After P1-A creates the complete top-level benchmark list, add one unchecked top-level checklist item in `plan.md` for every measured top-level bottleneck. Each item references the exact benchmark row ID and real operation name.
- When drill-down of one parent discovers its complete child list, add one nested unchecked child checklist item for every measured child before the first child Architect attempt. If one parent has ten measured children, the plan receives ten nested child checkboxes and `benchmark.md` receives ten corresponding numeric child rows.
- `benchmark.md` remains authoritative for elapsed-time numbers and ordering. The plan checklist mirrors row identity, operation name, parent/child relationship, and completion state so execution can mark exactly which bottlenecks have been processed.
- A child stays unchecked after an individual KEEP because optimization continues from the promoted baseline. Check the child only when its remaining cost reaches a terminal retained disposition. A concrete BLOCKED child remains unchecked.
- Check a parent only after every measured child beneath it is checked and accepted parent/full-pipeline remeasurement is recorded.
- If later accepted measurement exposes a missing child or top-level bottleneck, append its real benchmark row and matching unchecked plan item immediately; do not leave discovered work outside the checklist.
- P2-A and final exhaustion cannot close while any measured parent or child checklist item remains unchecked.

## 15. Purpose and limits of lane reports

A lane report exists to let Main and later sessions trace:

- which lane performed the work;
- what exact work it ran or changed;
- the concise result;
- current checkpoint/state;
- next owner/action.

The report is not the technical source of truth. It does not create another gate.

Forbidden loop:

```text
lane works
-> lane writes report
-> another lane audits report wording/numbers
-> another check compares report to Git/hash
-> another report proves the first report is correct
-> progress stops while documentation audits itself
```

Measurements belong in the benchmark ledger. Command/equivalence/implementation proof belongs in the evidence ledger. The report summarizes those results for orchestration.

Git and hash verification occurs only at a genuine technical identity, stable candidate, acceptance, or commit boundary where it materially matters. Do not hash or re-audit every report.

No report-about-report, evidence-about-evidence, documentation Supervisor, or handoff-chain audit is part of optimization.

## 16. Authority model

Owner is not an execution lane, START/HOLD gate, routine approver, or intermediate technical decision dependency.

The Owner has already granted Main command authority below Owner. Therefore:

- Main starts the authorized work;
- Main selects the next action from current measurements;
- Main dispatches and controls lanes;
- Main decides transitions and dispositions from evidence;
- lanes execute Main commands and report results;
- no lane commands Main or requires Owner approval;
- no plan step sends Main back to Owner for permission to continue.

Owner remains outside the operational workflow and intervenes only when Owner directly chooses to change direction, PAUSE, or STOP.

Architect is mandatory once for every production optimization attempt, not once for every unique bottleneck. Benchmark evidence selects the bottleneck; a fresh Architect invocation selects the technical direction for exactly one attempt; Planner turns that one decision into executable P2-A attempt steps; Coder executes only those steps. If the same bottleneck is attempted three times, Architect is invoked three times. If ten bottlenecks each receive one attempt, Architect is invoked at least ten times. Architect is not a numeric gate for routine measurement, does not choose priority, does not audit reports, and does not accept implementation.

Visible Supervisor is mandatory after remeasurement for every production optimization attempt. Supervisor independently decides whether accuracy/equivalence and preserved lifecycle invariants survived. A final whole-candidate Supervisor still runs at P3-A, but it is additional and never substitutes for per-attempt review. Supervisor REJECT returns evidence to a fresh Architect decision and Planner refresh before Coder can act again.

## 17. Correct replacement plan structure

The replacement standard four-file plan should use this structure.

### P0-A — Current state

- Record that P6-D is closed and Child 06A implementation has not started.
- Record the accepted real runtime/workload command.
- Record that no current phase ranking exists until the first real measurement.
- Preserve correctness/equivalence invariants.
- Open P1-A directly for the first real timing run; no Owner START gate and no harness project gate.

### P1-A — Detailed measurement and bottleneck inventory

This first work slice measures the current real pipeline. It does not start from a solution list and does not optimize production code.

1. Use exactly one visible measurement executor to run the accepted real `anvien analyze` workload and record current total graph-generation time; record the real options/workload/cache/runtime identity at execution rather than predeclaring it.
2. Measure real internal operation times: scan, parse, resolution, graph construction/insertion, persistence, graph-health/diagnostics, and post-run/publication, using the actual boundaries present in current code.
3. For each operation, record its elapsed time and work denominator directly in `benchmark.md`; record command validity and boundary proof in `evidence.md`.
4. Use existing timing/benchmark/profile capability first. Only when one required operation timing remains missing, run the single visible sequential instrumentation-writer branch defined in Section 13: fresh required graph/file-detail/impact on the exact owner, minimum owned edit, canonical full-build PASS before timing use, one accepted capture, like-for-like overhead/comparability and output-equivalence proof, then exactly one carry-forward or remove/rebuild/remeasure disposition. Do not build a general harness or parallel report system.
5. Produce the first evidence-ranked bottleneck list from current absolute times.
6. Write that full ordered bottleneck list directly into `benchmark.md` with rank, operation, current elapsed time, denominator, meaningful share/delta, proven owner/call path when available, processing/terminal state, and evidence. P1-A remains open until this complete list exists.
7. Populate the plan's `Measured Bottleneck Checklist` with one unchecked parent item for every top-level benchmark row. P1-A remains open until the benchmark list and matching plan checklist contain the same complete discovered parent inventory.
8. Record the current largest benchmark row and the exact next action in `actual-status.md`; do not duplicate numeric values there.
9. Before optimization begins, update the P2 optimization work from these real measurements. No hypothetical U1/U2/U3/U4 list is carried forward.

P1-A completes only when current total time, detailed internal timings, denominators, the mandatory evidence-backed ordered bottleneck list in `benchmark.md`, and the matching complete top-level checklist in `plan.md` exist. It is measurement work, not a speedup claim.

### P2 — Sequential measured-bottleneck optimization

P2 owns one implementation slice, `P2-A`, and one eventual implementation commit. The number and identity of bottlenecks are supplied by P1-A measurements and refreshed after every retained optimization.

Each dynamically opened bottleneck attempt inside P2-A must state the same concrete goal: reduce the current elapsed time of the exact benchmark-selected bottleneck and reduce total graph-generation elapsed time relative to the current accepted baseline. Supporting measurement or instrumentation work never replaces that goal.

P2-A is exhaustive. It processes every measured top-level bottleneck, largest first, then all remaining rows. Processing each selected parent begins with drill-down measurement inside that parent and creation of a complete measured child-bottleneck list. The child list is processed largest first and then through every remaining smaller child. The attempt cannot reach Architect until the current parent-child inventory and the selected child's elapsed-time owner/cause are identified with evidence. This drill-down remains part of the same dynamic P2-A bottleneck processing flow.

### P2-A — Optimize measured bottlenecks until exhausted

For the largest top-level bottleneck, then every remaining top-level bottleneck in descending elapsed-time order:

1. Start from its current measured time and denominator.
2. Process that selected parent by measuring more deeply within its boundary, writing its complete measured child-bottleneck list into benchmark, and inserting one nested unchecked plan checklist item for every discovered child before the first child attempt.
3. Select the largest remaining child, record the attempt goal against that exact child and parent row, and prove the child's exact cause, source owner, complete call path, and denominator.
4. Open attempt 1 for that child from the current accepted baseline and invoke a visible Architect with the exact parent/child benchmark and constraint package.
5. Update that exact P2-A attempt in all required ledgers from the fresh Architect decision; only then open Coder.
6. Run fresh file-detail/impact immediately before editing the actual owner.
7. Implement only the Architect-approved, Planner-recorded production optimization, then update tests.
8. Run the canonical full build.
9. Measure that same child, its parent bottleneck, and total graph-generation time again on the same work.
10. Open a visible Supervisor to verify the exact affected accuracy, correctness, output, persistence, reader, determinism, freshness, failure, transaction, temporary-file, publication, and lifecycle boundaries.
11. Record KEEP only if both direct/end-to-end speed improved and Supervisor passed; promote that candidate to current baseline and reset the no-KEEP count.
12. On Supervisor REJECT, restore/retain the last accepted baseline, record the unsuccessful attempt, return exact rejection evidence to a fresh Architect, refresh the attempt through Planner, then let Coder retry. Coder never chooses the corrective direction.
13. When accuracy passes but no gain can be retained, restore/retain the last accepted baseline and record the unsuccessful attempt.
14. Repeat the complete Architect -> Planner -> Coder -> measurement -> Supervisor sequence for every additional attempt on that child, even when its apparent direction looks unchanged.
15. After three consecutive attempts without KEEP from the current accepted baseline, record `SYSTEM_CHARACTERISTIC` for that child's remaining cost and stop targeting that child.
16. Check one child plan item only when that child reaches its terminal retained disposition, then continue to the next-largest unchecked child of the same parent. Small children are mandatory work; do not skip them or terminalize them by a time threshold.
17. When every measured child of the selected parent is checked, record accepted parent/full-pipeline remeasurement and check the parent item. Refresh the complete top-level list and move to the largest remaining unchecked top-level bottleneck.
18. Continue until every measured child and every measured top-level parent in both benchmark and the plan checklist has been optimized to a terminal retained disposition and checked.

Every KEEP becomes the next current baseline. Each child bottleneck's incremental before/after comparison uses the immediately preceding accepted state; the original initial baseline is preserved for final initial-versus-final comparison. Savings from both large and small children accumulate into the parent and final end-to-end reduction.

### P3-A — One final Supervisor

Review the stable complete final candidate once against the measured initial/final graph-generation speedup and all required preserved invariants. This whole-candidate review is additional to the mandatory per-attempt Supervisor checks. A reject returns through the exact affected bottleneck's fresh Architect -> Planner -> Coder -> measurement -> Supervisor loop.

### P3-B — Exact dead-work cleanup

Remove only Child 06A-created failed/superseded/debug artifacts. Do not create another documentation review. If cleanup changes accepted production/test bytes, final acceptance is invalidated and work returns to the exact owner.

### P3-C — Detect and one commit

Refresh final ledgers, run required detect-changes on the accepted implementation boundary, stage only accepted work, create one implementation commit, and hand Child 07 the closure commit.

## 18. Completion definition

Child 06A is not complete because it has a harness, many measurements, many reports, a full cost-map schema, or implemented previously suggested U units.

It is complete only when:

- the initial real full-pipeline timing exists;
- `benchmark.md` contains the mandatory complete current ordered bottleneck lists derived from real absolute timings and refreshes them after every retained or terminal child disposition;
- `plan.md` contains a matching living checklist for every measured top-level parent and every measured child; benchmark row IDs and checklist entries are complete and no discovered row is absent from either file;
- optimization processed every measured top-level bottleneck and every measured child bottleneck, ordered largest first at each level without dropping smaller rows;
- every opened bottleneck implementation attempt explicitly targeted lower elapsed time for that benchmark row and lower total graph-generation elapsed time from the current accepted baseline;
- each retained change has same-operation and end-to-end before/after proof;
- every production attempt has a fresh Architect decision, a corresponding Planner refresh, and a post-remeasurement Supervisor verdict;
- no Supervisor-rejected candidate became a baseline, and every correction returned through Architect and Planner before Coder;
- each child bottleneck with three consecutive no-KEEP attempts is recorded as `SYSTEM_CHARACTERISTIC` with its retained correct timing, denominator, and three attempt references;
- the pipeline was remeasured and reranked after each retained change;
- no measured top-level or child bottleneck remains unprocessed, including the smaller elapsed-time rows;
- no measured parent or child checklist item remains unchecked;
- final end-to-end analyze time is lower on the same workload;
- no candidate was counted as an optimization merely because a secondary CPU/RAM/allocation/GC/I/O/wait/count metric improved while elapsed graph-generation time did not;
- required correctness/output/persistence/reader/lifecycle equivalence passes;
- one final Supervisor passes;
- exact cleanup, detect, one commit, and Child 07 handoff complete.

## 19. Current coordination boundary

Visible Planner `01a033b7-9e01-7151-bdce-5d87173e1ac5` is actively rewriting the Child 06A four-file plan set under a direct Main command. Main has sent the binding clarifications recorded above and is monitoring actual writes without editing the plan in parallel.

At this checkpoint, direct workspace inspection proved the four Child 06A files still contained the superseded M0 harness, fixed A/A `20/20`, Architect numeric seal, fixed U1/U2/U3/U4 sequence, and sole-final-only Supervisor controls because the replacement patch had not yet completed. Main immediately instructed Planner to replace those executable residues rather than treat them as valid provenance. The plan is not accepted until actual file contents match this report and scoped verification passes.

The latest direct clarification expands the bottleneck method from largest-only targeting to exhaustive largest-first processing. P1-A must preserve the complete top-level measured list. When P2-A selects the largest parent bottleneck, processing that parent begins by measuring inside it and writing a complete child-bottleneck list. P2-A optimizes the largest child first and then every smaller measured child; it does not stop after the largest child and does not dismiss small elapsed-time rows. Retained savings across many small children accumulate into the parent and total graph-generation improvement. A parent is complete only after every measured child has a terminal retained disposition, and Child 06A is complete only after every measured top-level bottleneck has been processed in this manner.

To reduce the reading burden of the standard plan, the latest direct structural decision moves the long Child 06A rule body into the auxiliary `plan-rules.md`. The `## Rules` section of the standard plan points to that file instead of repeating the rules. Checklist, phase work, acceptance, the living Measured Bottleneck Checklist, and Current P2-A Attempt remain in the standard plan; no execution authority or ledger responsibility moves out of the four standard files except the detailed rule text itself.

Coder and measurement remain held. No implementation, build, benchmark run, Supervisor execution, commit, cleanup, or lane transition is authorized by this report update. Main rotation/handoff is explicitly paused until a later direct user command; Main must remain on the current task, monitor Planner, and correct deviations.

This document preserves the complete detailed method for Owner/Main discussion. It must not be treated as permission to patch the current plan until a later direct Owner decision explicitly converts the discussion into a Planner assignment.
