# System Architect Report: P6-E Plan Architecture Conformance

- Thời điểm: `2026-08-23 15:14:07 +07:00`
- Vai trò: System Architect, read-only architecture-conformance reviewer
- Claim được review: “Current P6-E plan có chuyển accepted analyze-performance solution architecture thành một slice đúng, đủ và triển khai được mà không đổi design intent hoặc làm yếu invariant hay không?”
- Verdict: **REJECT**
- Boundary: **architecture conformance only; not Supervisor acceptance or implementation authorization**.

## 1. Authority và phạm vi

Review chỉ dùng:

1. Accepted architecture: `reports/system-architect/rp_system-architect_260823_110218_by_gpt-5_analyze_performance_solution_architecture.md`.
2. Current P6-E section: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md:518-604`.

Không audit metadata, hash/seal của plan, wording identity giữa nhiều file, ledger completeness, implementation evidence hay Supervisor acceptance. Không sửa plan.

## 2. Kết luận

P6-E giữ đúng product goal và gần như toàn bộ topology của accepted architecture: baseline/cost map, U1, U2, fresh rebaseline, clean diagnostic ownership, U3, fresh rebaseline, conditional U4, fresh Pareto rerank và final end-to-end acceptance. Exact output/reader parity, determinism, freshness, fail-closed behavior, publication/failure parity, alternating unprofiled A/B, no-double-count và rollback từng step đều được giữ.

REJECT chỉ dựa trên ba khoảng trống kiến trúc thực sự còn cho phép implementation lệch solution đã accepted:

1. U2 chưa khóa exact key và purpose-specific index ownership.
2. U3 chưa khóa normalize-once algorithm/state transition là cơ chế loại secondary measured work.
3. Resource gate mới là numeric/measured budget; chưa giữ hard asymptotic bounds của từng unit.

Không có architecture input bị thiếu, nên verdict không phải BLOCKED.

## 3. Mapping accepted architecture sang P6-E

| Accepted architecture | Current P6-E mapping | Finding |
|---|---|---|
| §2 — exact sequential solution | Plan `:541-546` baseline/cost map; `:547-552` U1; `:553-558` U2; `:559-564` rebaseline/ownership; `:565-570` U3; `:571-576` rebaseline/conditional U4; `:577-582` Pareto; `:583-588` final acceptance | Conformant |
| §4.1 — byte/output equivalence | `:520`, `:543`, `:585`, `:597-598` | Conformant; graph bytes/SHA, ordered semantic outputs and canonical reader digests are exact gates |
| §4.2 — semantic completeness | `:520`, `:524-525`, `:579`, `:597-598` | Conformant; workload/output/persistence reduction is forbidden |
| §4.3 — determinism/order | `:524`, `:549`, `:553-555`, `:565`, `:583-598` | Conformant; scan/bucket order, stable sort, map lookup-only and replay are preserved |
| §4.4 — run freshness/disposal | `:520`, `:553-555`, `:565`, `:583-598` | Conformant; U2/U3 are run-scoped, U2 disposal is explicit, and final invalidation/restore is required |
| §4.5 — fail closed | `:524`, `:565`, `:569`, `:571-573`, `:597-598` | Conformant; missing node, malformed marker, six-field validation and property drift remain fail closed |
| §4.6 — persistence/native-reader parity | `:523`, `:530-531`, `:536`, `:543`, `:567`, `:573`, `:585`, `:597` | Conformant |
| §4.7 — transaction/temp/publication/failure parity | `:531`, `:536`, `:569`, `:583`, `:588`, `:597-598` | Conformant |
| §4.8 — no initial concurrency | `:525`, `:595` | Conformant; unmeasured concurrency and hidden concurrency are forbidden |
| §4.9 — hard resource bounds | `:541`, `:547`, `:553`, `:565`, `:571`, `:592-598` | **Gap A3**; measured budgets exist, but `O(1)`, `O(I)`, `O(D)` and bounded local U4 state do not |
| §4.10 — independent rollback | `:539`, `:547`, `:553`, `:565`, `:571`, `:593-604` | Conformant; each step is sealed, gated and rolled back without weakening prior steps |
| §6 — U1 canonical-path reuse | `:547-552`, `:595` | Semantics conform; scan/order/early exit and differential path cases remain exact. Its missing `O(1)` bound is tracked under Gap A3 |
| §7 — U2 immutable all-import claim index | `:553-558`, `:595` | **Partial — Gap A1**; immutable/run-scoped, unresolved coverage, original order, lookup-only use, shadow/target/label/nil rules and disposal are present, but exact key and separate-index ownership are absent |
| §8 — U3 diagnostic accumulator | `:565-570`, `:595` | **Partial — Gap A2/A3**; first duplicate, full stable sort including mutation/new bucket, immediate materialization, missing node, second writer/property drift and six-field fail-closed behavior are present; normalize-once and `O(D)` are absent |
| §9 — conditional one-pass decoder | `:571-576`, `:595` | Semantic and conditional gates conform; missing/null/type/malformed/unknown/duplicate-last-value/nested-raw and exact triple are explicit. Bounded local/no-memo state is missing under Gap A3 |
| §10 — later Pareto candidates | `:559-564`, `:577-582`, `:595` | Conformant; no DB/snapshot/parse/post-run edit opens without fresh absolute evidence and exact owner/contract gate |
| §11 — sequential A/B protocol | `:528`, `:541-546`, `:547-588`, `:592-598` | Conformant; same identity, A/A noise, alternating unprofiled A/B, profile separation, direct attribution, exact outputs and final end-to-end proof are required |
| §12 — P6-D ownership anchor | `:519`, `:524`, `:526`, `:538-539`, `:559-564`, `:591-592` | Conformant; P6-E cannot edit diagnostics before accepted/committed clean ownership |
| §13 — forbidden shortcuts | `:525-526`, `:538`, `:577-579`, `:595-598` | Conformant; stale cache, reduced workload/output, GC substitute, arbitrary thresholds, unmeasured concurrency and target access are forbidden |
| §14–15 — bounded claims, residual unknowns and fresh rerank | `:526`, `:541`, `:559-582`, `:583`, `:596` | Conformant; no promised seconds/percentage and old percentages lose priority authority |

## 4. Architecture gaps và exact Planner corrections

### A1 — U2 exact data structure/ownership chưa đủ

- Plan location: current plan `:553-558`.
- Violated architecture: §7, especially “Algorithm/data structure” and its prohibition on repurposing the resolved-only receiver index.
- Why this weakens the design: “all-import claim index” alone does not lock the selected safe key. An implementation could use a broader/name-only bucket, or mutate/repurpose an index with different unresolved-import semantics. That changes the accepted purpose-specific ownership, can retain avoidable scan work, and can collide with existing consumers even if basic parity fixtures pass.
- Exact Planner correction: amend Step 3 to require a **second purpose-specific** immutable index with exact shape `{canonical source file path, exact local name} -> []original w.imports index`; append every resolved and unresolved import in original order; do not repurpose `importsByReceiver`; use the map only for direct lookup and never for output iteration; keep final resolution answers uncached.

### A2 — U3 chưa yêu cầu normalize existing/incoming đúng một lần

- Plan location: current plan `:565-570`.
- Violated architecture: §8 “Algorithm/data structure”, rules 2–3 and the retained normalized-state design.
- Why this weakens the design: Step 5 names an accumulator and behavioral parity, but does not require the causal work-removal mechanism. A bucket accumulator could still re-normalize the existing slice on later appends or normalize/decode an incoming diagnostic twice, preserving graph bytes while failing the accepted architecture and direct attribution.
- Exact Planner correction: amend Step 5 to require: normalize a node's pre-existing diagnostic slice exactly once on first access; normalize/decode every incoming diagnostic exactly once, including structured policy and all six nested-authority checks; retain the current normalized slice plus first-matching bucket index; after every new-bucket full stable sort rebuild that index to the first matching value; prove new-vs-existing normalization/decode counters in same-schema A/B evidence.

### A3 — hard asymptotic resource contract bị thay bằng corpus-specific budget

- Plan locations: current plan `:541`, `:547`, `:553`, `:565`, `:571`, `:592-598`.
- Violated architecture: §4 invariant 9; §6/§7/§8/§9 complexity and memory contracts.
- Why this weakens the design: Owner-approved peak/resource budgets are necessary but not equivalent to structural bounds. A superlinear or memoizing implementation can pass on the current corpus while violating the accepted scalability/freshness design.
- Exact Planner correction: state hard per-unit bounds alongside the numeric gates:
  - U1: additional retained memory `O(1)`;
  - U2: index construction/state `O(I)` and run-scoped only;
  - U3: accumulator state `O(D)` and run-scoped only;
  - U4: one bounded local wire/nested-raw state per call, with no memoization;
  - no unbounded package-global, global, shared or cross-run cache.

## 5. Extra/conflicting behavior

- The single final P6-E commit at `:539` and `:604` is not itself an architecture conflict. Steps are sequentially sealed, each failed step is rolled back before continuation (`:593`), and earlier accepted units remain operationally retainable. Accepted architecture requires independent rollback, not a specific commit count.
- Build, impact, cleanup, ledger and Owner-threshold governance add implementation controls but do not alter the solution design.
- No conflicting performance shortcut or P6-D ownership collision was found.

## 6. Handoff

Handoff to Main: return P6-E to Planner for only corrections A1–A3. Do not redesign the accepted sequence, split the Owner-selected checklist slice merely for architecture reasons, or weaken any already-conformant gate. After the exact corrections, route a fresh architecture-conformance review; Supervisor acceptance remains closed until Architect PASS.

Detached file seal is emitted externally after final byte materialization because a SHA-256 of the complete report cannot be embedded inside the bytes it seals.

STOPPED
