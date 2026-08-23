# System Architect Report: Kiến trúc tối ưu tuần tự `analyze` theo cost center

- Trạng thái: **HOÀN TẤT — durable architecture handoff; không phải plan, implementation hay verdict nghiệm thu**
- Thời điểm: `2026-08-23 11:02:18 +07:00`
- Repository: `E:\Anvien`
- Vai trò: System Architect, chỉ thiết kế solution từ causal evidence đã được Root Cause xác minh
- Next owner: Main Orchestration
- Boundary: không chạy lại build/analyze/profile/test/graph command; không sửa code/test/plan/ledger/AGENTS/skill/Git/target; không dùng phase label làm hệ quy chiếu

## 1. Authority và evidence đã đọc

1. Root Cause report: `reports/Investigation/rp_investigation_260823_103434_by_gpt-5_analyze_performance_root_cause.md`
   - `32,620` bytes / `341 LF` / `0 CR`
   - SHA-256 `100998373E396A7E66B887357214B5952C4F9D90BC4D26558F291F9F60B950B6`
2. Narrow Main verification: `reports/Supervisor/rp_supervisor_260823_104751_by_gpt-5_root_cause_handoff.md`
   - `6,307` bytes / `60 LF` / `0 CR`
   - SHA-256 `F960F54E18C80D517C5899E75D32E512D6F3F346056A9191D4E6986D89CE26CE`
3. Preserved implementation anchor: `reports/Supervisor/rp_supervisor_260823_064551_by_gpt-5_p6d_graph_health_projection_resubmission.md`
   - current amended bytes `33,046` / `239 LF` / `0 CR`
   - SHA-256 `3837059F852CF44F20732F2BB0F7D88234BC6612AAB9C7CD439AEC3288A768DA`
4. Current source owners được đọc trực tiếp tại `internal/resolution`, `internal/graphhealth`, `internal/analyze`, cùng graph mutation và Ladybug CSV/load boundaries được các owner trên gọi trực tiếp.
5. Raw `AGENTS.md`, `working-rules`, và `System-Architect` skill được áp dụng. Điều khoản chung của skill về SPEC/plan/commit bị Owner authority cụ thể thu hẹp: output duy nhất của lane này là report chưa commit.

## 2. Kết luận kiến trúc

Không có một “resolution optimization” nguyên khối. Kiến trúc tốt nhất là chuỗi unit nhỏ, mỗi unit loại đúng một loại công việc đã đo, A/B và rollback độc lập:

1. **Canonical-path reuse:** bỏ normalize lặp trên import path đã canonicalize, nhưng chưa đổi thuật toán scan.
2. **All-import claim index:** thay whole-import scan bằng immutable run-scoped index chứa cả resolved lẫn unresolved imports.
3. **Run-scoped diagnostic accumulator:** normalize/decode mỗi diagnostic một lần trong run, dùng bucket index nhưng vẫn materialize graph property và giữ nguyên thứ tự/failure behavior hiện tại.
4. **Presence-aware single-pass structured decoder:** bỏ lần parse toàn bộ JSON thứ hai mà không nới fail-closed policy.
5. Sau bốn unit trên, **đo lại Pareto trên baseline mới** rồi chỉ chọn cost center tuyệt đối lớn nhất tiếp theo. DB load, snapshot, parse và post-run không được xếp thứ tự dựa trên phần trăm cũ sau khi resolution thay đổi.

Lựa chọn này không dùng cache xuyên run và không thêm concurrency vào primary/secondary path. Vì vậy freshness, invalidation và deterministic merge không bị đưa vào critical path trước khi cần thiết.

Canonical-path reuse không bị gộp vào claim index vì nó có `O(1)` memory, giữ nguyên scan semantics và loại trực tiếp allocation đã đo. Nếu keyed index không đạt parity hoặc resource gate, unit đầu vẫn là một tối ưu độc lập có thể giữ lại. Ngược lại, nếu gộp hai thay đổi, một lỗi index sẽ buộc rollback cả phần path reuse đã chứng minh an toàn và làm mất attribution A/B.

## 3. Causal input: verified và measurement gaps

### Verified current run

| Cost center | Evidence |
|---|---:|
| Retained process wall | `599,915.2275 ms` |
| Benchmark total | `596,568.8634 ms` |
| Resolution | `533,746.9551 ms` (`89.469%`) |
| DB load | `25,403.8452 ms` (`4.258%`) |
| Parse | `14,739.7751 ms` (`2.471%`) |
| In-run unphased | `14,911.2894 ms` (`2.500%`) |
| Wrapper-minus-benchmark | `3,346.3641 ms` |
| Workload | `2,075 / 763 / 0`; `120,853` nodes; `167,428` relationships |

Primary verified mechanism:

- `explicitImportNameClaimed`: `39.43 s` flat / `319.95 s` cumulative.
- `explicitImportCallState`: `3.90 s` flat / `47.51 s` cumulative.
- `cleanPath`: `0.88 s` flat / `324.75 s` cumulative; approximately `149,590.52 MB` allocation and `3.057B` objects.
- Source proves import paths are normalized at workspace construction, then normalized again for every visited item in whole-import scans.

Secondary verified mechanism:

- `AppendDiagnosticToNode`: `113.20 s` cumulative.
- `decodeStructuredResolutionOutcome`: `112.93 s` cumulative.
- Decoder/append allocation approximately `20–22 GB`, on `97,806` successfully attached diagnostics.
- Source proves each append reloads and re-normalizes the existing diagnostic slice; structured notes are decoded once as an envelope and again as the complete outcome.

Các cumulative samples trên **overlap**. Chúng không phải wall time và không được cộng thành speedup.

### Measurement gaps giữ nguyên

- Chưa có exact `Q`, `I`, loop-iteration count, hit/miss distribution hoặc claim-bucket distribution.
- Chưa có diagnostic-list histogram theo node hoặc exact new-versus-existing normalize/decode multiplicity.
- Chưa có exact wall split cho CSV export/native COPY, snapshot encode/flush/rename, parser theo language/file, lock/I/O/block/mutex, hoặc post-run substeps.
- Run hiện tại có profiler overhead; nó là attribution authority, không phải unprofiled throughput baseline.
- Không có comparable historical baseline; report này không gán regression cho commit/slice nào.

## 4. Architecture contract bất biến

Mọi unit phải thỏa đồng thời:

1. **Output equivalence:** cùng corpus/runtime/options phải tạo cùng node, relationship, ID, property, evidence, source-site outcome, diagnostic, count và metadata. `graph.json` phải byte-identical, không chỉ parse-equivalent.
2. **Semantic completeness:** không bỏ file, parser, fact family, lookup stage, edge, proof, diagnostic, resolution outcome, community/process/semantic step hoặc reader.
3. **Determinism:** không iterate map để phát output; mọi bucket giữ original input order; mọi merge/sort dùng đúng comparator và stable behavior hiện tại.
4. **Freshness:** index/accumulator chỉ sống trong một `analyze` run, dựng từ current ScopeIR/graph, không persist và không reuse qua source/config/runtime change.
5. **Fail closed:** malformed structured marker vẫn thành `unclassified/review`; missing source node vẫn trả false; invalid index/accumulator state phải abort, không silent fallback sang kết quả ít evidence hơn.
6. **Persistence parity:** in-memory graph, Graph JSON và Ladybug/native readers phải có cùng canonical entity/relationship/evidence digest và workload counts.
7. **Publication semantics:** giữ temp-write, flush/close, atomic rename và native transaction/rollback contract; failed run không được quảng bá artifact như thành công.
8. **Concurrency independence:** bốn unit đầu không thêm concurrency. Mọi concurrency tương lai phải compute vào indexed result slots rồi reduce theo canonical input order.
9. **Resource bounds:** memory phụ chỉ được `O(I)` cho import index và `O(D)` cho diagnostic state; không có unbounded global cache. Peak Sys/alloc/object counts là acceptance metrics.
10. **Independent rollback:** một unit không được phụ thuộc vào output relaxation của unit khác. Bất kỳ mismatch, nondeterminism, freshness/parity failure hoặc resource regression vượt budget đều kích hoạt rollback unit đó.

## 5. So sánh options theo verified cost center

| Cost center | Feasible options | Quyết định kiến trúc |
|---|---|---|
| Repeated path normalization | (A) reuse stored canonical path; (B) memoize `cleanPath`; (C) persistent path cache | Chọn A. B thêm state không cần thiết; C tạo invalidation/freshness risk. |
| Whole-import scans | (A) scan nhanh hơn; (B) keyed all-import claim index; (C) parallel resolver scan | Chọn B sau canonical-path unit. C không loại wasted work và mở nondeterministic graph merge. |
| Repeated diagnostic normalization | (A) skip validation khi classification đã có; (B) global memo; (C) run-scoped node accumulator | Chọn C. A bypass P6-D fail-closed invariant; B có stale/lifecycle risk. |
| Double whole-note JSON decode | (A) giữ envelope map + full decode; (B) presence-aware one-pass wire decode; (C) trust status marker only | Chọn B sau accumulator. C nới validation và bị cấm. |
| Allocation-driven GC | (A) tune/disable GC; (B) loại upstream allocations | Chọn B. GC samples overlap primary/secondary paths và không phải root độc lập. |
| DB load | (A) tránh aggregate relationship CSV không được loader dùng; (B) parallel native COPY; (C) giảm rows/columns/evidence | A là candidate sau substep timing/contract check; B chờ transaction/order proof; C bị cấm. |
| Parse | (A) bounded ordered file workers; (B) persistent incremental cache; (C) skip slow languages/files | A chỉ khi parse trở thành Pareto và parser thread-safety được chứng minh; B không dùng cho `--force`; C bị cấm. |
| Graph snapshot | (A) reusable byte-compatible encoding scratch; (B) compact/compress/change schema; (C) skip snapshot | A chỉ sau named wall boundary; B/C bị cấm trong contract hiện hành. |
| Unphased/post-run | (A) parallelize blindly; (B) add exact boundaries/counters first | Chọn B; chưa có causal owner đủ hẹp để đổi behavior. |

## 6. Optimization unit 1 — canonical-path reuse

### Source owner

- Workspace import normalization: [`internal/resolution/indexes.go`](../../internal/resolution/indexes.go), `resolveImports` near line 287.
- Hot consumers: [`internal/resolution/resolve.go`](../../internal/resolution/resolve.go), `explicitImportNameClaimed` near line 900; [`internal/resolution/export_resolution.go`](../../internal/resolution/export_resolution.go), `explicitImportCallState` near line 306.

### Algorithm/data structure

`resolvedImport.Fact.FilePath` is already canonicalized before append to `w.imports`. `scopeFilePath` also returns a canonical path. Compare these stored strings directly inside the two hot loops; do not call `cleanPath` for every visited item.

This unit deliberately keeps the scan. Nó isolates path-allocation removal from later indexing so A/B can attribute each improvement.

### Complexity và memory

- Before: `O(Q × I)` comparisons plus `O(Q × I)` path transformations/allocations.
- After: `O(Q × I)` plain string comparisons; path transformations for visited items become zero.
- Additional retained memory: `O(1)`.

### Determinism/concurrency

Loop order and match rule remain unchanged. Không có goroutine, map iteration hoặc reordered work.

### Failure/rollback

Path equivalence is valid only because workspace construction canonicalizes every stored import path. A test/debug differential oracle must compare direct comparison with the legacy `cleanPath` comparison on Windows separators, dot segments, empty paths and unresolved imports. Any answer drift rolls back this unit; production must not retain the legacy scan as a hidden fallback.

### Semantics giữ nguyên

Resolved/unresolved import ownership, shadowing, TypeScript external blocking, target selection, graph counts, evidence, diagnostic carriers and output ordering must be byte/semantic equivalent.

### Proof status và expected impact

- **Can be proved now:** redundant transformation and idempotent canonical storage are source-confirmed.
- **Needs A/B:** exact wall saving and remaining `cleanPath` calls by caller.
- **Confidence:** high directional allocation reduction.
- **Bound:** saving cannot exceed the `533.747 s` resolution envelope; `324.75 s` cumulative CPU samples are not a wall-time promise.

## 7. Optimization unit 2 — immutable all-import claim index

### Source owner

- Workspace/index construction and existing `importsByReceiver`: [`internal/resolution/indexes.go`](../../internal/resolution/indexes.go).
- Claim consumers: [`internal/resolution/resolve.go`](../../internal/resolution/resolve.go) and [`internal/resolution/export_resolution.go`](../../internal/resolution/export_resolution.go).

### Algorithm/data structure

Build a second, purpose-specific index during `resolveImports`:

`{canonical source file path, exact local name} -> []original w.imports index`

The bucket must include **every** matching import, including unresolved imports, and append indices in original `w.imports` order.

Không được repurpose `importsByReceiver`: current source only inserts entries whose `LinkStatus != "unresolved"`, while both hot claim scans currently consider unresolved imports. Reusing it would silently allow forbidden external/global fallback.

- `explicitImportNameClaimed`: membership/len lookup only.
- `explicitImportCallState`: iterate only the keyed bucket, then apply the exact current semantic-import, target, allowed-label and shadow rules in original order.

### Complexity và memory

Let `I` be all imports, `Qn` name-claim queries, `Qc` call-state queries, and `k(file,name)` the matching bucket size.

- Before: `O((Qn + Qc) × I)`.
- After build: `O(I)` once.
- Name query: expected `O(1)`.
- Call-state query: `O(k(file,name))`.
- Memory: `O(I + K)` integers/map buckets, run-scoped only.

### Determinism/concurrency

Map is used only for direct lookup, never iteration. Bucket order is original import order. Workspace remains single-thread constructed and immutable during resolution.

### Failure/rollback

Debug/correctness mode must compute legacy and indexed answers for the full current workload and abort on the first mismatch, recording key/caller without publishing graph output. Production timing mode uses only the index. Không fallback sang a partial/resolved-only index.

### Semantics giữ nguyên

Unresolved imports still claim names; semantic vs non-semantic import distinction remains at the consumer; local shadow handling and `target == nil` behavior remain exact. No resolution/export result is precomputed or cached in this unit.

### Proof status và expected impact

- **Can be proved now:** whole-slice scan and exact safe key are source-confirmed; existing receiver index's unresolved exclusion is also source-confirmed.
- **Needs counters:** `Qn`, `Qc`, `I`, hits/misses and bucket-size distribution for sizing and A/B attribution.
- **Confidence:** high structural improvement; exact seconds unknown.
- **Bound:** shares the same resolution envelope as unit 1 and is not additive with its old profile samples.

## 8. Optimization unit 3 — run-scoped diagnostic accumulator

### Source owner

- Diagnostic policy/aggregation: [`internal/graphhealth/diagnostics.go`](../../internal/graphhealth/diagnostics.go).
- Resolution emitters: [`internal/resolution/emit.go`](../../internal/resolution/emit.go) and [`internal/resolution/outcome.go`](../../internal/resolution/outcome.go).
- Resolution lifecycle: [`internal/resolution/resolve.go`](../../internal/resolution/resolve.go), `ResolveBoundInto`.

### Algorithm/data structure

Create one emitter-owned diagnostic appender per resolution run. For each node it maintains:

- the current normalized diagnostic slice;
- an exact bucket-key to first matching slice index map;
- stable ordering state matching the current comparator;
- the current graph/node identity.

Rules:

1. Validate graph, node ID, kind and node existence at each append exactly as today.
2. Normalize pre-existing node diagnostics once on first access to that node.
3. Normalize each incoming diagnostic exactly once, including the complete structured policy and six-field nested-authority checks.
4. Preserve current duplicate behavior: if pre-existing properties already contain duplicate buckets, index only the first bucket for merge and leave later duplicates untouched.
5. Preserve current merge rules for count, target text and minimum start line.
6. On a new bucket, run the exact current full stable ordering over current values, then rebuild the bucket index so it still points to the first matching value. A prior merge can mutate a sort key without re-sorting; the next new bucket must reproduce today's full-sort behavior, not assume the slice remained sorted.
7. Fetch the latest node before materialization, replace only its diagnostic property, and write it back immediately. Other node properties must not be overwritten. This preserves current partial-result/failure visibility; đây không phải deferred publish cache.
8. Resolution owns the diagnostic property exclusively while this run-scoped appender is active. Any second writer or observed property drift is an invariant failure and aborts fail-closed.

Generic `AppendDiagnosticToNode` remains the compatibility API. Resolution obtains the persistent run-scoped appender so repeated calls share normalized state; no package-global cache is allowed.

### Complexity và memory

Let `d_n` be appends for node `n`, `u_n,j` its existing diagnostic buckets before append `j`, `D` total diagnostics, and `B` total structured-note bytes.

- Current normalization/decode work: proportional to `Σ_n Σ_j u_n,j` existing diagnostics, with whole-note JSON work repeated; bucket search is also linear.
- After: each existing/incoming diagnostic is normalized once, `O(B)` decode work plus expected `O(D)` bucket lookups.
- Stable-order maintenance may retain the current sort asymptotic initially to guarantee exact bytes; it operates on already-normalized state and must be measured separately before any later ordering optimization.
- Additional memory: `O(D)` run-scoped keys/indices. The graph already owns diagnostic values; no duplicate payload cache across runs.

### Determinism/concurrency

Single-threaded appender, canonical input sequence, exact stable comparator, no map iteration for output.

### Failure/rollback

Missing node returns false at the same append point. Any corrupt cached state or graph identity drift aborts analysis fail-closed. A differential harness must feed identical sequences to legacy and new appenders and compare return values plus node properties after every append, including duplicate buckets and mutated sort keys.

### Semantics giữ nguyên

Diagnostic order/count, classification/actionability, SourceSite fields, Note bytes, ResolutionGap inputs, Graph JSON and summary remain exact. Prefilled classification can never bypass structured revalidation.

### Proof status và expected impact

- **Can be proved now:** repeated existing-slice normalization/decode and linear bucket scan are source/profile-confirmed.
- **Needs counters:** list-length histogram, first-duplicate cases, normalize/decode calls and bytes before/after.
- **Confidence:** high for allocation/decode reduction; medium for exact wall magnitude.
- **Bound:** `113.20 s` append and `112.93 s` decoder samples overlap; neither is a wall forecast and they must not be summed.

## 9. Optimization unit 4 — presence-aware single-pass structured decoder

### Source owner

- [`internal/graphhealth/diagnostics.go`](../../internal/graphhealth/diagnostics.go), structured outcome decode and validation boundary.

### Algorithm/data structure

Replace the current whole-note envelope-map unmarshal followed by whole-note typed unmarshal with one presence-aware wire decode that produces typed fields and records whether `schemaVersion` and `status` keys were present.

The decoder contract must match current `encoding/json` behavior for:

- missing versus explicit `null` keys;
- wrong field types;
- malformed JSON;
- duplicate keys/last-value behavior;
- unknown fields;
- nested `target` and `authority` raw values.

Nested authority validation, including all six equality groups, remains unchanged. If single-pass decoding cannot preserve the exact three-result contract `(outcome, structured, valid)`, this unit is rejected.

### Complexity và memory

- Before: two full traversals of note bytes plus nested validation.
- After: one full traversal plus the same nested validation.
- Complexity remains `O(B)` but with a lower constant and without the envelope map/full-note duplicate allocation.
- Memory is bounded by one wire outcome and nested raw fields per call; no memoization.

### Determinism/concurrency

Pure local decode, no concurrency or ordering change.

### Failure/rollback

Marker-present invalid input remains `unclassified/review`. Legacy fallback is allowed only when no structured status marker and no structured keys exist, exactly as today. Differential adversarial tests must compare old/new decode triples and normalized diagnostics byte-for-byte.

### Proof status và expected impact

- **Can be proved now:** current code parses the whole note twice.
- **Ordering condition:** execute only after unit 3 is accepted and rebaseline shows decoder remains a Pareto cost; unit 3 may already remove most repeated calls.
- **Confidence:** high for per-call work removal, medium/low for whole-analyze wall after unit 3.
- **Bound:** no “50% faster” claim; nested authority decode/validation and non-JSON work remain.

## 10. Next Pareto candidates — không được kích hoạt từ old percentages

### Allocation/GC

Không có GC tuning unit độc lập. `gcBgMarkWorker`, `gcDrain` và `scanObject` overlap application work. Acceptance theo dõi GC/alloc/object/peak Sys giảm sau units 1–4; không đổi `GOGC`, không disable GC và không đổi memory pressure thành OOM risk.

### DB load

Current exact envelope: `25.404 s`.

Source proves `LoadCSVExport` consumes per-node and per-endpoint-pair CSVs, not aggregate `relations.csv`; aggregate relationship rows are nevertheless written in the current export path. Candidate an toàn nhất là một analyze-loader export mode that:

- keeps public/default aggregate CSV contract for existing consumers/tests;
- for the analyze native-load path, writes the exact sorted pair rows needed by COPY and computes the same row/skipped metrics without a second aggregate row stream;
- preserves pair-file order, transaction, COPY count, zero-fallback/zero-skip fail-closed gates and canonical native-reader digest.

This candidate removes source-confirmed transient I/O but **must wait** for separate export/native/commit wall and byte counters. Parallel COPY is not selected: the current single transaction/runner and reader ordering require proof first. Reducing rows, properties or evidence is forbidden.

### Graph snapshot

Current evidence: `10.76 s` cumulative CPU, `2,024.95 MB` allocation, `666,772,547` output bytes, but exact wall is only bounded inside `14.911 s` unphased work.

Candidate: reusable per-item encoding scratch that emits byte-for-byte the current `MarshalIndent` representation and preserves node/relationship insertion order, metadata key ordering, temp cleanup, flush/close and atomic rename. Compact JSON, compression, schema change or snapshot omission are rejected. Named encode/flush/close/rename walls are required before activation.

### Parse

Current exact envelope: `14.740 s` for `763` files and `6,974,573` parser input bytes.

Candidate only if it becomes Pareto: bounded file-level workers, each with safe parser ownership; results stored by original scan index and reduced in index order. The failure contract must return the earliest input-order error and the same prefix metrics/IRs, regardless of completion order. Memory is bounded by worker count times in-flight source/tree size. Parser-pool thread safety, per-language timings and file-size distribution must be proved first. Persistent/incremental cache is disallowed for current `--force` semantics.

### Unphased/post-run

Snapshot, DB close, benchmark write, registry, AI context and file projection need named monotonic boundaries and I/O/block/mutex attribution before a solution is selected. Blind overlap/parallel execution is rejected because it can change publication/failure ordering and create resource contention.

### Measured non-primary owners

Workspace construction/cross-file binding, TypeScript catalog lookup, external materialization, outcome finalization và outcome projection đều đã được measurement đặt dưới primary/secondary owners; riêng projection là `3.84 s` cumulative, finalization `0.29 s`, catalog construction khoảng `0.15 s`, external materialization khoảng `0.01 s`. Kiến trúc chọn **không sửa** các owner này trước rebaseline. Precompute/caching/concurrency ở đây sẽ mở invalidation surface lớn hơn bounded saving hiện có. Nếu một owner trở thành Pareto sau units 1–4, nó phải nhận một fresh exclusive/counter boundary trước khi có candidate riêng.

## 11. Sequential A/B protocol per unit

Protocol này là architecture acceptance contract, không phải execution plan hoặc commit topology.

### 11.1 Frozen pair identity

- Same source corpus manifest: exact relative paths, content hashes and scanner exclusions; không tạo report/artifact mới trong scanned corpus giữa A và B.
- Same direct E-only runtime build contract: Go version, build tags, options, native library, environment and executable manifest; A/B differ only by the one unit under evaluation.
- Same `--force` storage state and a declared identical cold/warm OS-cache policy. Không so bare/global wrapper với direct runtime.
- Same benchmark/instrumentation schema in A and B. CPU/heap/profile runs are separate from unprofiled wall trials.

### 11.2 Cost-center metrics

For units 1–2 record:

- claim calls/items/hits/misses by caller;
- call-state calls/items/hits/misses;
- `cleanPath` calls/input-output bytes by caller;
- index entries, bucket count, average/max bucket and build bytes/time.

For units 3–4 record:

- append calls and existing-list length histogram;
- new/existing normalization counts;
- structured decoder calls, full-note passes and decoded bytes;
- accumulator nodes/buckets/bytes and stable-sort work.

For later candidates record exact named substep wall, CPU/allocation and I/O bytes/ops. Mỗi unit phải giảm cost center trực tiếp của nó; whole-analyze delta một mình không đủ attribution.

### 11.3 Whole-analyze performance

- Use alternating unprofiled trials under the same cache policy, not one profiled A versus one unprofiled B.
- Report median, dispersion and confidence interval for cost-center wall and whole wall.
- Acceptance requires cost-center reduction outside measured noise and no unexplained regression in sibling phases/resources.
- Cumulative gain is measured by a final end-to-end run against the original comparable baseline; do not sum per-unit inclusive CPU samples or percentages.

### 11.4 Workload and output equivalence

A and B must match exactly on:

- scanned/parsed/failed and full fact-family denominators;
- graph node/relationship/dependency-edge counts and every resolution metric;
- ordered resolution outcomes, diagnostics, ResolutionGap inventory, classifications/actionability and evidence;
- `graph.json` bytes and SHA-256;
- canonical ordered node/relationship/evidence digests from in-memory, Graph JSON and Ladybug/native readers;
- built CLI/MCP/HTTP reader results where applicable.

Timing fields and newly added diagnostic counters are excluded from graph-output hash comparison, not from benchmark review.

### 11.5 Deterministic replay và freshness

- Run B from two clean forced storage states; both runs must emit identical graph hash and canonical native-reader digest.
- Run-scoped index/accumulator counts must reset to zero at run start and dispose at run end.
- A controlled source-hash change must invalidate all derived state and appear in the next output; restoring source and forcing analysis must restore the original hash. No stale graph/cache reuse is accepted.

### 11.6 Failure equivalence

Exercise missing file/read failure, context cancellation, invalid structured diagnostic, missing source node, DB COPY/transaction failure, snapshot encode/write/rename failure and malformed cache/index state relevant to the unit. Error class, fail-closed behavior, rollback/temp cleanup and publication boundary must match current contract.

### 11.7 Rollback triggers

Rollback the unit if any of these occurs:

- one graph/output/native-reader mismatch;
- one nondeterministic replay;
- stale result or invalidation gap;
- weaker malformed-marker or missing-node behavior;
- skipped/failed relationship, fallback or reduced workload;
- peak memory/resource bound exceeded;
- direct cost center does not improve outside noise, or whole wall regresses without an accepted causal explanation.

## 12. Preserved P6-D implementation anchor

### Exact current anchor

- Seven-path candidate: `377,945` bytes; canonical SHA-256 `E4FF42DE8AA1941010AFFF5DD38B6789AE71168E3EBF9F59D8EECF735D5BD044`.
- `internal/graphhealth/diagnostics.go`: `26,884` bytes; SHA-256 `518BD06FF5F859BD42AD7B3B9FF8E3F9F4E7235720BC5E568E2C4B2DA72BD2AE`.
- `internal/graphhealth/p6d_outcome_projection_test.go`: `19,336` bytes; SHA-256 `6E90E655FFCCF5A5BD3D398D95EB5B6B7F77754CAAF35DAAF4D9609248A4E713`.
- Local repair enforces exact equality for file path, file hash, all four range coordinates, site kind, requested name and requested meaning.

Exact seven-path manifest giữ làm input anchor:

| Path | Bytes | SHA-256 |
|---|---:|---|
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md` | 65,194 | `A30538D7D13BE6E3AED4CFA8965F3BA22CB716543CA35670267FC745FE23390B` |
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md` | 22,207 | `34405EAE2861D56ED3B304475BBF40D53BA2C831D3E2F43AF7F775A9DCED2A53` |
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md` | 138,757 | `07997D8345D79847DC1C08CDBAE7A237E84A8EB1133FCDA4EF8EAFABB5553936` |
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md` | 85,442 | `B125AAE5CB7C163DAEAD6805E2278EBC29D679390F0185BBB99E4934701DD492` |
| `internal/graphhealth/diagnostics.go` | 26,884 | `518BD06FF5F859BD42AD7B3B9FF8E3F9F4E7235720BC5E568E2C4B2DA72BD2AE` |
| `internal/graphhealth/p6d_outcome_projection_test.go` | 19,336 | `6E90E655FFCCF5A5BD3D398D95EB5B6B7F77754CAAF35DAAF4D9609248A4E713` |
| `reports/coder/rp_coder_260823_032449_by_gpt-5_p6d_graph_health_projection_blocked.md` | 20,125 | `6E8AB2DBC96B43C16FDC7102FCA89D5096CD76D6B597D162C732FED533F72420` |

Report này không sửa anchor.

### Gates có thể reuse làm contract

- Existing `3/3` policy rows, `2/2` catalog-proof states, `26/26` hostile rows, `1/1` legacy fallback and `1/1` Graph JSON/ResolutionGap/summary parity remain the required semantic matrix.
- Six-field comparison logic and malformed-present fail-closed behavior remain authoritative expected behavior.
- Units 1–2 do not touch the two graph-health anchor files; their exact local source/test identity can remain the implementation anchor.

### Evidence bị invalidated bởi performance changes

- Units 3–4 touch `diagnostics.go`; therefore exact seven-path candidate identity and prior local source clearance are no longer current evidence. The six-field code must remain semantically unchanged, and the complete matrix above must rerun fresh.
- Units 1–2 still create a different product runtime even if graph-health bytes stay exact. Old full-runtime, graph hash, persistence/native-reader, target/oracle/pre-post evidence cannot be reused as acceptance for that runtime.
- DB/snapshot/parse changes invalidate their respective runtime/persistence/output gates even when P6-D source bytes are untouched.
- The amended Supervisor report remains `REJECT`; it cannot be promoted to PASS by this architecture report. Canonical build/current runtime/native readers/target/oracle/pre-post and clean independent review remain fresh gates under later authority.

## 13. Forbidden shortcuts

- Bỏ file/parser/fact/edge/proof/diagnostic/relationship hoặc giảm output payload.
- Dùng stale graph, global/shared cache, cross-run memo hoặc cache key thiếu source/config/catalog/parser/schema/runtime identity.
- Dùng `--skip-compatibility-cross-file` hoặc mode giảm workload làm production speedup.
- Reuse `importsByReceiver` như all-import claim index khi nó đang loại unresolved imports.
- Trust prefilled classification/actionability để skip structured validation.
- Relax six-field nested-authority equality, marker-present fail-closed behavior hoặc native load completeness.
- Iterate map để emit graph/diagnostics, hoặc parallel workers ghi trực tiếp vào shared graph.
- Compact/compress/change Graph JSON, skip snapshot/native persistence, drop aggregate semantics chưa được chứng minh là transient, hoặc giảm validation.
- Tune GC như substitute cho việc loại allocations.
- Báo speedup bằng cách cộng inclusive CPU samples.

## 14. Bounded impact và confidence

| Unit/candidate | Expected direction | Hard current-run envelope, không phải forecast | Confidence |
|---|---|---|---|
| Canonical-path reuse | rất lớn trên alloc/object count ở import scan | nằm trong resolution `533.747 s`; không dùng `324.75 s` như wall saving | High |
| All-import claim index | loại `O(Q×I)` scan còn lại | cùng resolution envelope; không cộng với unit 1 profile | High structural, wall unknown |
| Diagnostic accumulator | loại repeated existing-note normalize/decode | append/decoder samples overlap trong resolution | High structural, medium wall |
| Single-pass decoder | giảm constant của mỗi decode còn lại | chỉ định lượng sau accumulator rebaseline | High per-call, medium/low whole wall |
| DB candidate | giảm một transient relationship-row stream nếu contract cho phép | tối đa trong current DB `25.404 s` | Medium |
| Snapshot candidate | giảm encoding allocation nhưng giữ `666+ MB` bytes | current wall chỉ biết nằm trong `14.911 s` unphased | Medium |
| Ordered parse concurrency | giảm wall khi parser work parallelizable | current parse `14.740 s` | Low/medium trước thread-safety/timing |
| Post-run | chưa chọn solution | wrapper residual `3.346 s` | Low trước boundaries |

Không unit nào có promised seconds hoặc percentage. Sau mỗi acceptance, baseline kế tiếp thay đổi; old percentages không còn là priority authority.

## 15. Residual unknowns và handoff

Residual unknowns:

1. Exact import query/iteration/bucket multiplicity và diagnostic node-size distribution.
2. Resource size thực của new indexes/accumulator trên full workload.
3. Exact DB export/native/commit, snapshot encode/flush, parse-by-language và post-run wall/I/O splits.
4. Parser pool/thread safety và deterministic earliest-error behavior dưới concurrency.
5. Whether analyze-specific omission of transient aggregate relationship CSV is formally outside the failure/output contract.
6. Noise floor và acceptance threshold của future unprofiled A/B environment.

Các unknown này không ngăn việc chọn units 1–3 về mặt architecture; chúng ngăn speedup promise và ngăn kích hoạt later Pareto candidates thiếu attribution.

## 16. Output và commit reference

- Output file duy nhất: report này.
- Code/test/plan/ledger/AGENTS/skill/target/artifact hiện hữu: không sửa.
- Commit reference: **none — Owner authority cấm stage/commit trong discussion**.
- Handoff: Main Orchestration đọc và zero-trust verify report; Architect lane STOP ngay sau seal.
