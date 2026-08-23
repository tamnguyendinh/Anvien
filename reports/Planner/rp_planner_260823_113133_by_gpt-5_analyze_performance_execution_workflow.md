# Planner Report: Workflow thực thi cost map và tối ưu tuần tự `analyze`

- Trạng thái: **HOÀN TẤT — durable workflow proposal; chưa cho phép implementation**
- Thời điểm lập: `2026-08-23 11:31:33 +07:00`
- Repository duy nhất: `E:\Anvien`
- Vai trò: Planner; chuyển causal evidence đã được chấp nhận hẹp và solution contract đã chọn thành dependency/order, evidence gates, benchmark protocol, rollback/commit boundaries và topology đề xuất
- Next owner: Main Orchestration
- Boundary của report: không chạy code/test/build/analyze/profile/graph/Git/package/target; không sửa code, test, active Child 06 plan/ledgers hoặc artifact hiện hữu; report này là file mới duy nhất

## 1. Kết luận workflow

Khuyến nghị chính xác là **không nhét performance vào candidate graph-health đang REJECT và cũng không coi một nhãn phase kế tiếp là quyết định ban đầu**. Sau khi Main/Supervisor chấp nhận proposal, hãy tạo một campaign performance độc lập, cross-cutting, bắt đầu bằng cost-map instrumentation và comparable baseline, rồi thực thi từng optimization unit nhỏ theo chuỗi:

```text
authority + plan thật + cost-map instrumentation được PASS/commit
  -> comparable unprofiled baseline được seal
  -> Unit 1 canonical-path reuse được PASS/commit
  -> Unit 2 all-import claim index được PASS/commit
  -> fresh absolute rebaseline được seal/commit
  -> hard gate giải quyết ownership của dirty diagnostics anchor
  -> diagnostic instrumentation baseline được PASS/commit
  -> Unit 3 run-scoped diagnostic accumulator được PASS/commit
  -> fresh absolute rebaseline được seal/commit
  -> nếu decoder vẫn là measured Pareto center:
       Unit 4 presence-aware single-pass decoder được PASS/commit
     nếu không:
       ghi N/A có evidence, không implement Unit 4
  -> fresh Pareto rerank cho DB/snapshot/parse/post-run
```

Không optimization unit nào được mở khi unit trước chưa đồng thời có: exact candidate identity, canonical full build hợp lệ, nearest-boundary validation, comparable A/B, accuracy/failure/publication equivalence, independent Supervisor PASS, fresh detect-changes, cleanup disposition và isolated commit. Một improvement nhỏ chỉ được giữ nếu cost center trực tiếp giảm ngoài noise, whole-run không regress, resource budget không regress và mọi invariant giữ nguyên. Không cộng inclusive CPU samples hoặc phần trăm cũ để hứa speedup.

Units 1–2 có thể thực hiện trên cùng dirty seven-path anchor vì production/test surfaces không giao nhau, với điều kiện anchor được seal trước/sau và tuyệt đối không stage. Units 3–4 chạm `internal/graphhealth/diagnostics.go`; chúng **không được mở** cho tới khi lifecycle sở hữu graph-health giải quyết candidate đang REJECT thành một committed, clean baseline theo authority mới. Không dùng stash, reset, checkout hay alternate worktree để né gate này.

## 2. Authority và phân loại evidence đầu vào

### 2.1 Durable inputs đã kiểm chứng identity

| Input | Identity hiện hành | Quyền sử dụng trong workflow |
|---|---|---|
| Root Cause | `reports/Investigation/rp_investigation_260823_103434_by_gpt-5_analyze_performance_root_cause.md`; `32,620` bytes; SHA-256 `100998373E396A7E66B887357214B5952C4F9D90BC4D26558F291F9F60B950B6` | causal authority cho current profiled run; không phải historical regression hoặc unprofiled baseline |
| Narrow causal handoff | `reports/Supervisor/rp_supervisor_260823_104751_by_gpt-5_root_cause_handoff.md`; `6,307` bytes; SHA-256 `F960F54E18C80D517C5899E75D32E512D6F3F346056A9191D4E6986D89CE26CE` | PASS hẹp cho causal handoff; không PASS solution/implementation/phase |
| Selected architecture | `reports/system-architect/rp_system-architect_260823_110218_by_gpt-5_analyze_performance_solution_architecture.md`; `36,360` bytes; SHA-256 `4CD4DB7EE6D195ADDED4CB0E0879EA1C2C288B81596F56248B6624485C2957BC` | solution contract cho Units 1–4 và Pareto loop |
| Current amended graph-health anchor | `reports/Supervisor/rp_supervisor_260823_064551_by_gpt-5_p6d_graph_health_projection_resubmission.md`; current `33,046` bytes; SHA-256 `3837059F852CF44F20732F2BB0F7D88234BC6612AAB9C7CD439AEC3288A768DA` | input anchor có verdict REJECT; local six-field repair cleared nhưng không được promote thành PASS |
| Active Child 06 plan/ledgers | bốn file đã đọc qua EOF | current dirty history/ownership anchor; performance campaign không được sửa hoặc dùng chúng làm ledger riêng |

### 2.2 Measured, estimated và không comparable

| Claim | Phân loại | Cách dùng |
|---|---|---|
| Controlled direct-E profiled process elapsed `599,915.2275 ms`; benchmark `596,568.8634 ms`; resolution `533,746.9551 ms`; workload `2,075/763/0`, graph `120,853/167,428` | **measured**, exact start/end và benchmark/profile artifacts; có profiler overhead | authority cho attribution/cost-center selection hiện tại; không dùng làm unprofiled throughput baseline |
| Successful retained `scripts\full-build.ps1` elapsed `678,182 ms` (`11m18.182s`) với output `2,074/763/0`, graph `120,843/167,418` | **measured single successful command**, nhưng runtime đi qua global wrapper ngoài E | historical workload pointer; không phải E-only runtime hoặc A/B authority |
| Turn elapsed `1,042,071 ms` (xấp xỉ `17m20s`) | **measured turn**, gồm một full-build mất terminal handle rồi rerun | tuyệt đối không gọi là single-command timing |
| “analyze khoảng 8 phút” | **estimated từ poll timing**; không có exact analyze start/end, benchmark JSON hoặc phase timestamps | triệu chứng lịch sử; không phải baseline so sánh |
| Regression từ khoảng 2 phút lên 5–8 phút hoặc commit/slice gây regression | **không có comparable evidence** | không dùng trong threshold, attribution hoặc acceptance |

Artifact/timestamp tối thiểu còn thiếu để một future timing trở thành comparable gồm: exact command argv; wall-clock start/end có timezone; monotonic process start/end và elapsed; per-full-build-step start/end; riêng analyze start/end; exit; runtime/source/dependency/corpus/cache/instrumentation identities; benchmark JSON; direct cost-center counters; stdout/stderr seal; graph/output/native-reader digests; trial order và regime. Thiếu bất kỳ identity quyết định nào thì timing chỉ là historical context.

## 3. Invariant và role boundary

Mọi step giữ tuyệt đối:

1. Accuracy và semantic completeness: không giảm file, parser, fact family, lookup stage, edge, proof, diagnostic, outcome, relationship, metadata, reader hoặc workload.
2. Determinism/order: không emit bằng map iteration; bucket giữ original order; comparator và stable-sort behavior giữ đúng bytes.
3. Freshness: index/accumulator chỉ run-scoped, dựng từ current inputs, reset đầu run, dispose cuối run; không cross-run/shared/stale cache.
4. Persistence/native-reader parity: in-memory, Graph JSON và Ladybug/native reader phải cho canonical ordered entity/relationship/evidence digest giống nhau.
5. Failure/publication semantics: fail closed; error class và point-of-failure giữ nguyên; failed run không publish graph/native/registry artifact như success; temp-write/flush/close/rename và transaction rollback giữ nguyên.
6. Resource bounds: Unit 1 thêm `O(1)` retained memory; Unit 2 tối đa `O(I)`; Unit 3 tối đa `O(D)`; Unit 4 local bounded decode state. Không đổi allocation thành unbounded resident cache.
7. Target `E:\cheapapp.org` giữ locked cho tới authority riêng sau này; không dùng target để chứng minh performance hoặc parity trong campaign này.

Planner không xác minh lại root cause, không thiết kế lại technical solution, không code và không tự PASS implementation. Nếu Coder cần đổi data structure, ownership, failure/publication contract hoặc allowed surface ngoài architecture report, step phải STOP và quay lại Main/Architect. Independent Supervisor là owner duy nhất của acceptance review; Main là owner của stage/commit/transition.

## 4. Living granular cost map contract

Cost map là artifact có version, không phải bảng percentage tĩnh. Schema đề xuất `analyze-cost-map/v1` có bốn lớp record.

### 4.1 Run/trial identity record

| Field | Requirement |
|---|---|
| `campaign_id`, `run_id`, `trial_id`, `pair_id`, `sequence_index` | unique, immutable; liên kết đúng A/B pair và trial order |
| `variant` | `A`, `B`, hoặc `A/A-noise`; không suy từ filename |
| `profile_mode` | `unprofiled`, `counter-only`, `cpu-profile`, `heap-profile`, `trace`; không trộn khi tính wall decision |
| `cache_regime` | `cold`, `warm`, hoặc truthful `fresh-storage/ambient-os-cache`; không gọi cold nếu OS cache không được kiểm soát |
| `started_at`, `ended_at`, `monotonic_start_ns`, `monotonic_end_ns`, `wall_ns` | timestamp ISO-8601 có `+07:00` và monotonic elapsed |
| `repo_head`, `repo_dirty_manifest_sha256`, `candidate_delta_manifest_sha256` | A/B chỉ khác exact allowed unit delta; unrelated và seven-path anchor bytes giữ nguyên |
| `corpus_manifest_sha256`, `scanner_exclusion_sha256`, `options_sha256` | exact input workload identity |
| `runtime_exe_sha256`, `runtime_manifest_sha256`, `go_version`, `build_tags` | direct E-only runtime identity |
| `native_library_sha256`, `dependency_manifest_sha256`, `environment_whitelist_sha256` | provenance và native identity |
| `storage_prestate_sha256`, `cache_policy_id`, `instrumentation_schema_sha256` | cùng forced-storage/cache/instrumentation contract |
| `command_argv`, `cwd`, `exit_code`, `stdout_sha256`, `stderr_sha256` | reproducible invocation; `cwd` chỉ E repo |

### 4.2 Cost-center record

Mỗi row có:

- `cost_center_id`, `parent_cost_center_id`, `phase`, `substep`, `owner_symbol`;
- `caller_symbol`, `callee_symbol`, `call_path`, `source_identity`;
- `wall_inclusive_ns`, `wall_exclusive_ns`, `cpu_sample_ns` hoặc exact CPU time nếu có;
- `alloc_bytes`, `alloc_objects`, `live_bytes`, `peak_rss_bytes`, `peak_sys_bytes`, GC counts/pause nếu có;
- `io_read_bytes`, `io_write_bytes`, `io_read_ops`, `io_write_ops`;
- `wait_ns`, `wait_reason`, `blocked_ns`, `mutex_ns` khi được instrument;
- `work_unit_name`, `work_unit_count`, `call_count`, `bytes_processed`;
- histogram `{count,sum,min,p50,p95,max}` với unit rõ ràng;
- `measurement_kind` (`direct-counter`, `monotonic-span`, `profile-inclusive`, `derived-exclusive`, `unmeasured-residual`);
- `artifact_pointer`, `artifact_sha256`, `measurement_status` (`measured`, `estimated`, `not-instrumented`, `not-applicable`).

Required function/call-path counters cho Units 1–2:

- import name-claim calls/items visited/hits/misses theo direct caller;
- call-state calls/items visited/hits/misses;
- `cleanPath` calls và input/output bytes theo caller;
- `I`, index entries, bucket count/size histogram, index build wall/alloc/bytes.

Required counters cho Units 3–4:

- append calls; existing-list length histogram;
- new-vs-existing normalization counts;
- structured decoder calls, whole-note passes, decoded bytes;
- accumulator node/bucket count, retained bytes, first-duplicate cases, stable-sort calls/items;
- malformed/present-marker/fail-closed outcome counts.

Toàn analyze phải có named spans hoặc một explicit residual row cho: pre-timer lock/force prep; scan; parse; cross-file binding; resolution; semantic/process/community/tool/route/document owners; DB export/native COPY/commit/close; graph snapshot encode/flush/close/rename; benchmark write; heap-profile/GC step; registration; AI context; file projection; final output. Không được để thời gian biến mất vì chưa instrument.

### 4.3 Workload/output equivalence record

Mỗi trial ghi exact scanned/parsed/failed, parser bytes và fact-family denominators; mọi resolution outcome counter; nodes, relationships, dependency edges, unresolved projections; diagnostics và ResolutionGap counts; graph bytes/SHA-256; canonical in-memory/Graph JSON/native digests; reader parity rows; publication pre/post identity; failure injection outcome nếu applicable.

### 4.4 No-double-count và reconciliation rules

1. `exclusive wall = inclusive wall - union(child spans)`; không trừ tổng mù khi child spans overlap.
2. Chỉ cộng sibling exclusive leaves hoặc disjoint monotonic spans. Không cộng caller/callee inclusive samples, GC concurrent samples hoặc two profile rows của cùng call path.
3. CPU/heap profile dùng để attribute, không thay wall time. Profiled và unprofiled trial không nằm cùng speed estimator.
4. Mỗi run phải reconcile `process wall = pre-timer + benchmark named leaves + in-run residual + wrapper/post-run leaves + unreconciled residual`. Residual khác zero phải được nêu, không phân bổ tùy ý.
5. Sau mỗi accepted commit, rerank bằng **absolute current durations** và current direct counters. Percentage/cumulative sample từ baseline cũ hết quyền chọn next unit.
6. Cost map giữ lineage: mỗi accepted candidate trở thành baseline kế tiếp, nhưng raw rows cũ immutable; không overwrite để tạo cumulative story.

## 5. Controlled benchmark và A/B protocol

### 5.1 Frozen identities

- A và B dùng cùng corpus/input bytes, scanner exclusions, command options, environment whitelist, dependency/native identity và cache regime. Repository source chỉ được khác exact unit delta đã seal; mọi unrelated bytes, đặc biệt dirty seven-path anchor ở Units 1–2, phải byte-identical.
- Cả hai runtime dùng cùng build contract và instrumentation schema. Không so direct E runtime với bare/global wrapper.
- Không tạo report/ledger/artifact mới trong scanned corpus giữa hai trial của một pair. Raw evidence ghi vào approved scanner-excluded run root; durable docs được tạo sau khi pair set đóng.
- Baseline và candidate đều chạy `--force`; storage prestate/reset contract phải giống nhau.

### 5.2 Cache regimes

1. `cold`: chỉ dùng tên này nếu có Owner-authorized, repeatable OS-cache reset và exact storage reset trong E. Nếu không có, không fabricate cold result.
2. `warm`: mỗi variant có cùng preconditioning rule; warm-up bị loại khỏi estimator; measured pairs dùng balanced alternating order.
3. Nếu OS cache không controllable, dùng nhãn `fresh-storage/ambient-os-cache`, alternating `AB/BA` để cân order effect, và không trộn với warm/cold trong cùng estimator.

### 5.3 Repetitions và threshold

- Trước khi có candidate result, chạy baseline-vs-baseline (`A/A`) để đo within-host paired noise cho từng regime.
- Owner khóa trước một stopping/repetition rule dựa trên precision cần thiết của confidence interval và run budget. `n` phải cố định trước khi xem B, bằng nhau cho A/B, và không early-stop vì result thuận lợi.
- Whole wall và direct cost-center wall báo median paired delta, dispersion và confidence interval. Raw trials không bị bỏ trừ khi có predeclared invalidation reason (sai identity, holder/lock, nonzero exit, cache-prestate mismatch); mọi exclusion phải retained.
- Owner đặt minimum material improvement là `max(noise-derived detectable delta, product-value floor do Owner chọn)` sau A/A nhưng trước B. Report này không đặt số giây/phần trăm tùy ý.
- Owner cũng đặt resource budget từ baseline distribution cho peak memory, allocation, objects và I/O. Một regression resource không được đổi lấy speed nếu chưa có authority mới.

### 5.4 Profiler và unprofiled separation

- Alternating unprofiled trials quyết định performance acceptance.
- Counter-only paired runs xác minh direct mechanism/work units khi counters tạo overhead không đáng kể và giống A/B.
- CPU/heap/trace runs tách riêng, dùng để attribute và cập nhật call paths; không tham gia whole-wall estimator.
- Final cumulative gain là một current end-to-end A/B so với original comparable baseline đã freeze. Không cộng per-unit gains, inclusive CPU samples hoặc old percentages.

### 5.5 Accuracy, determinism, freshness và failure gates

A/B phải exact-match:

- workload/fact-family denominators và mọi output/resolution counter;
- ordered resolution outcomes, diagnostic bytes/order/count, classification/actionability, SourceSite, evidence và ResolutionGap inventory;
- `graph.json` bytes và SHA-256;
- canonical ordered node/relationship/evidence digest từ in-memory, Graph JSON và Ladybug/native readers;
- built CLI/MCP/HTTP reader outputs khi surface đó bị ảnh hưởng;
- two fresh-forced deterministic B replays;
- controlled source-hash change xuất hiện ở next output; restore + force trả original hash;
- missing/read error, cancellation, invalid structured marker, missing source node, DB failure, snapshot failure và malformed run-state behaviors liên quan unit;
- failed-run pre/post publication identity, transaction rollback và run-owned temp cleanup.

Timing/counter fields chỉ được loại khỏi graph-output hash nếu chúng không phải graph contract; chúng vẫn bắt buộc trong benchmark review.

## 6. Global prerequisite gates

### `G0-PLAN` — durable plan authority

- Owner: Main giao Planner tạo plan thật sau khi report này được chấp nhận.
- Allowed: một new performance plan riêng và đủ bốn ledgers theo planner skill.
- Cấm: sửa bốn active Child 06 ledgers hoặc đổi trạng thái candidate graph-health.
- Evidence ID: `E-PERF-G0-PLAN-001` — exact four-ledger identity, dependencies và Owner approval.
- Stop: không có plan thật/Owner approval thì không code.
- Transition: plan-only review/commit hoàn tất mới mở measurement implementation.

### `G0-AUTH` — build/runtime/dependency authority

- Owner: Owner/Main.
- Required evidence:
  - `E-PERF-G0-BUILD-AUTH-001`: explicit authority cho repository canonical `scripts/full-build.ps1` **hoặc** exact Owner-approved hermetic equivalent giữ cùng product steps/contract;
  - `E-PERF-G0-DEPS-001`: sealed dependency provenance, content hashes, source và new permitted non-protected E-only cache/root;
  - `E-PERF-G0-RUNTIME-001`: exact E-only executable/native manifest; không bare/global wrapper/C authority;
  - `E-PERF-G0-CACHE-001`: storage/cache regimes và cleanup authority;
  - `E-PERF-G0-LOCK-001`: holder preflight protocol chỉ name/PID, không đọc/emit external command metadata.
- Full-build rule: trước mỗi validation set, terminate verified build-related holders theo exact PID trong authority, rồi chạy full build. Internal analyze của full-build là build-step evidence; dedicated validation/A/B vẫn phải chạy sau build.
- Stop: canonical script vẫn đòi network/global install/C discovery mà chưa được authorize/contain; dependencies chưa có provenance; runtime không direct E; holder state không clean.

### `G0-IDENTITY` — corpus và dirty-anchor freeze

- Owner: Main/measurement Coder; Supervisor zero-trust verify.
- Evidence:
  - `E-PERF-G0-CORPUS-001`: exact relative path/content/scanner-exclusion manifest;
  - `E-PERF-G0-REPO-001`: HEAD plus full dirty manifest, with current authority HEAD `c28660118fc606ea3e19ad2f6cfb206a768f46f1`;
  - `E-PERF-G0-ANCHOR-001`: seven-path total `377,945` bytes / canonical SHA-256 `E4FF42DE8AA1941010AFFF5DD38B6789AE71168E3EBF9F59D8EECF735D5BD044`.
- Stop: any unrelated/anchor byte changes during A/B or pre/post unit.

## 7. Standard execution template cho mọi code-changing step

Áp dụng cho measurement instrumentation và Units 1–4:

1. **Main opens exact step** với allowed manifest, non-goals, baseline identity và Owner threshold/budget đã seal.
2. **Coder pre-edit evidence:** current direct-E graph refresh theo repository rule; `file-detail` + upstream impact cho mỗi file/symbol owner; report HIGH/CRITICAL blast radius. Nếu graph/runtime authority chưa hợp lệ, STOP.
3. **Production code first.** Chỉ sửa production behavior/instrumentation trong allowed surfaces. Nếu cần surface/contract ngoài architecture, STOP về Main/Architect.
4. **Tests second.** Chỉ sau production behavior hoàn tất mới thêm/update differential, regression và failure tests.
5. **Canonical full build before validation.** Retain per-step timestamps, runtime/native identity, exact analyze substep start/end/exit; full build phải exit `0`.
6. **Nearest real boundary validation** cho changed contract, rồi affected/full regressions; mỗi command ghi nó chứng minh gì.
7. **A/B + equivalence:** chạy frozen protocol, direct counters, output/native parity, deterministic/freshness/failure/publication/resource gates.
8. **Worker cleanup:** chỉ exact run-owned disposable paths theo authority; không cleanup raw retained benchmark bundle, dirty anchor, shared/protected/quarantined roots.
9. **Worker detect:** fresh graph trên final candidate, `detect-changes --scope all`, giữ cả full-worktree noise và isolated unit attribution; không normalize HIGH/CRITICAL.
10. **Independent Supervisor:** zero-trust source/diff/evidence/identity review trong fresh visible lane; timing decision dựa sealed pair set, không dựa một ad-hoc rerun. REJECT giữ step open và cấm stage.
11. **Main final seal/detect:** cập nhật performance ledgers/reports sau benchmark; nếu docs làm graph stale thì refresh/detect lại; seal exact accepted manifest và exclusions.
12. **Isolated commit:** stage exact path list, verify cached diff/index, assert seven anchor paths absent khi applicable, commit một step. Post-commit verify unit clean và ambient anchor unchanged.

Rollback không dùng `git reset`, `checkout`, stash hoặc alternate worktree. Khi REJECT, không stage/commit; designated Coder áp explicit inverse patch chỉ trên unit manifest về sealed baseline, rồi re-hash baseline và cleanup run-owned outputs. Nếu mismatch là invariant/security/data publication, giữ artifacts để điều tra và STOP trước inverse patch cho tới Main authority.

## 8. Sequential implementation workflow

### `M0` — granular cost-map instrumentation foundation

- Goal: có measurement schema/reconciliation cho toàn analyze và direct resolution counters cần cho Units 1–2, không đổi graph/output semantics.
- Owner: measurement Coder; independent Supervisor review; Main commit.
- Allowed production surfaces: current benchmark/timing owners trong `internal/analyze/analyze.go`, CLI wrapper owners trong `internal/cli/command.go`, exact benchmark schema owners được impact xác nhận, và resolution counter hook owners `internal/resolution/indexes.go`, `internal/resolution/resolve.go`, `internal/resolution/export_resolution.go`. Tests chỉ cạnh các owner này. Exact manifest phải được Main seal sau fresh impact.
- Explicitly forbidden: `internal/graphhealth/diagnostics.go`, `internal/graphhealth/p6d_outcome_projection_test.go`, bốn Child 06 ledgers, existing reports, graph schema/output payload, target.
- Work order: production instrumentation first; schema/reconciliation/disabled-mode tests second; canonical full build; instrumentation-off/on output equivalence; overhead A/A; nearest benchmark JSON boundary; Supervisor; detect; isolated commit.
- Required evidence:
  - `E-PERF-M0-IMPACT-001`, `E-PERF-M0-IMPL-001`, `E-PERF-M0-BUILD-001`;
  - `B-PERF-M0-AA-001` for instrumentation overhead/noise;
  - `E-PERF-M0-RECON-001` for 100% process-wall reconciliation including explicit residual;
  - `E-PERF-M0-EQ-001`, `E-PERF-M0-SUP-001`, `E-PERF-M0-DETECT-001`, `E-PERF-M0-COMMIT-001`.
- Stop/rollback: instrumentation changes graph/native output, publication/failure order, workload, or overhead/resource budget; timer/counter cannot be disabled/separated from unprofiled decision; implementation needs package-global/cross-run state or diagnostics anchor.
- Completion/transition: Supervisor PASS + isolated commit `perf(map): add granular analyze cost accounting`; only then establish `B0`.

Diagnostic interior fields chưa thể instrument mà không chạm anchor phải mang trạng thái `not-instrumented`, không zero. Root Cause profile vẫn là current attribution pointer. `M3` sẽ đóng fields này sau hard anchor gate.

### `B0` — comparable baseline and Owner decision rule

- Goal: tạo baseline distribution cho exact accepted M0 runtime/corpus/anchor, không implement optimization.
- Owner: benchmark Coder; Main/Owner pre-register decision rule; independent Supervisor verifies bundle.
- Required bundle:
  - `B-PERF-B0-AA-WARM-001` và, chỉ nếu có valid reset authority, `B-PERF-B0-AA-COLD-001`;
  - `B-PERF-B0-UNPROFILED-001` alternating baseline trials;
  - `B-PERF-B0-PROFILE-001` separate attribution refresh;
  - `E-PERF-B0-COSTMAP-001`, `E-PERF-B0-THRESHOLD-001`, `E-PERF-B0-SEAL-001`.
- Stop: no comparable distribution, trial identity drift, unknown cache regime, or Owner threshold/resource rule not frozen before B candidate.
- Completion: sealed evidence-only checkpoint/commit `perf(bench): establish analyze cost-map baseline`. B0 is the only performance baseline for U1; historical 17m20/~8m/full-build values remain excluded.

### `U1` — canonical-path reuse

- Goal: loại repeated normalization/allocation trên already-canonical import paths nhưng giữ nguyên whole-import scan/order.
- Owner: Coder; Supervisor; Main.
- Allowed production files: `internal/resolution/indexes.go`, `internal/resolution/resolve.go`, `internal/resolution/export_resolution.go`; focused resolution tests/benchmark hooks only. Exact post-impact manifest may be narrower, không rộng hơn nếu thiếu Architect addendum.
- Non-goals: claim index, concurrency, cache, diagnostics, graph schema, output ordering, P6-D/target.
- Work: direct compare stored canonical paths per architecture; production first; then differential tests cho Windows separators, dot segments, empty paths, unresolved imports và caller outcomes; full build; nearest resolver boundary; full analyze A/B/parity.
- Evidence IDs:
  - `E-PERF-U1-IMPACT-001`, `E-PERF-U1-IMPL-001`, `E-PERF-U1-BUILD-001`, `E-PERF-U1-NEAR-001`;
  - `B-PERF-U1-AB-WARM-001` and conditional `B-PERF-U1-AB-COLD-001`;
  - `E-PERF-U1-COST-001` (`cleanPath` caller counts/bytes and alloc/object delta);
  - `E-PERF-U1-EQ-001`, `E-PERF-U1-DETERMINISM-001`, `E-PERF-U1-FRESHNESS-001`, `E-PERF-U1-FAILURE-001`, `E-PERF-U1-RESOURCE-001`;
  - `E-PERF-U1-SUP-001`, `E-PERF-U1-DETECT-001`, `E-PERF-U1-COMMIT-001`.
- Stop/rollback: any legacy-vs-direct answer drift; one graph/native mismatch; direct center noise-only; whole wall/resource regression; anchor identity drift.
- Completion/transition: Supervisor PASS + isolated commit `perf(resolution): reuse canonical import paths`; commit excludes all seven anchor paths. Only committed U1 becomes baseline A for U2.

### `U2` — immutable all-import claim index

- Goal: thay whole-import scans bằng run-scoped index `{canonical source path, exact local name} -> original w.imports indices`, gồm cả unresolved imports.
- Owner: Coder; Supervisor; Main.
- Allowed production files: current workspace/index owners trong `internal/resolution/indexes.go`, claim consumers `internal/resolution/resolve.go` và `internal/resolution/export_resolution.go`; focused resolution tests/counters. Không repurpose `importsByReceiver` nếu nó còn loại unresolved imports.
- Non-goals: precompute final resolution, cross-run cache, map iteration emission, concurrency, diagnostics/graph-health, target.
- Work: build immutable original-order buckets; production first; then full legacy-vs-index differential oracle on resolved/unresolved, hit/miss, duplicate/local-shadow, semantic/non-semantic, nil target and label cases; full build; nearest boundary; full A/B/equivalence.
- Evidence IDs:
  - `E-PERF-U2-IMPACT-001`, `E-PERF-U2-IMPL-001`, `E-PERF-U2-BUILD-001`, `E-PERF-U2-NEAR-001`;
  - `B-PERF-U2-AB-WARM-001` and conditional `B-PERF-U2-AB-COLD-001`;
  - `E-PERF-U2-COST-001` for Q/I/items/hits/misses and bucket histogram/build cost;
  - `E-PERF-U2-EQ-001`, `E-PERF-U2-DETERMINISM-001`, `E-PERF-U2-FRESHNESS-001`, `E-PERF-U2-FAILURE-001`, `E-PERF-U2-RESOURCE-001`;
  - `E-PERF-U2-SUP-001`, `E-PERF-U2-DETECT-001`, `E-PERF-U2-COMMIT-001`.
- Stop/rollback: unresolved import bị bỏ; bucket/original-order drift; legacy/index mismatch; memory không còn `O(I)`; direct center noise-only; output/resource/anchor mismatch.
- Completion/transition: Supervisor PASS + isolated commit `perf(resolution): index all import claims`; không mở U3 trực tiếp. Bắt buộc vào `R12`.

### `R12` — fresh rebaseline after Units 1–2

- Goal: đo current absolute Pareto; không reuse `89.469%`, `324.75s`, `113.20s` hoặc old percentages như priority.
- Owner: benchmark Coder; Supervisor verifies; Main seals/commits evidence.
- Required evidence: `B-PERF-R12-UNPROFILED-001`, `B-PERF-R12-COUNTERS-001`, `B-PERF-R12-PROFILE-001`, `E-PERF-R12-COSTMAP-001`, `E-PERF-R12-PARETO-001`, `E-PERF-R12-SEAL-001`.
- Decision gate:
  - nếu diagnostic append/decode vẫn là current measured Pareto center ngoài noise và Owner materiality floor, tiếp tục tới `A3`;
  - nếu không, STOP planned U3 và quay về Main/Pareto loop; không implement chỉ vì old profile.
- Commit boundary: evidence-only `perf(bench): rebaseline after import optimizations` trước mọi diagnostic edit.

### `A3` — dirty-anchor ownership resolution, hard prerequisite

- Owner: Main/Owner và graph-health lifecycle owner; Planner không quyết định acceptance.
- Current truth: seven-path anchor REJECT, unstaged, uncommitted; local six-field repair cleared không đồng nghĩa full acceptance.
- Satisfactory transition duy nhất: một authority mới giải quyết graph-health candidate qua full required lifecycle và tạo clean committed baseline, hoặc Owner đưa ra một explicit supersession contract được independently reviewed/committed. Không được stage exact rejected candidate, stash/reset/checkout, copy sang alternate worktree, hoặc lặng lẽ absorb rejected bytes vào performance commit.
- Evidence: `E-PERF-A3-DISPOSITION-001`, `E-PERF-A3-BASELINE-001`, `E-PERF-A3-SUP-001`, `E-PERF-A3-COMMIT-001`.
- Stop: `diagnostics.go` hoặc its test vẫn là rejected uncommitted ownership; target/runtime gates của graph-health vẫn blocked; không có clean Supervisor PASS/commit.
- Transition: chỉ clean committed source/test baseline mới cho phép `M3`.

### `M3` — diagnostic cost-map extension

- Goal: thêm direct diagnostic list/decode/sort counters với exact same instrumentation ở A/B trước U3.
- Owner: measurement Coder; Supervisor; Main.
- Allowed: `internal/graphhealth/diagnostics.go` và exact benchmark instrumentation/test owners được impact xác nhận; resolution emit call sites chỉ nếu architecture contract cần. Không thay diagnostic algorithm trong step này.
- Evidence: `E-PERF-M3-IMPACT-001`, `E-PERF-M3-BUILD-001`, `B-PERF-M3-AA-001`, `E-PERF-M3-EQ-001`, `E-PERF-M3-SUP-001`, `E-PERF-M3-DETECT-001`, `E-PERF-M3-COMMIT-001`, `B-PERF-B3-BASELINE-001`.
- Stop: instrumentation itself thay output/sort/failure, cannot isolate overhead, hoặc mở state ngoài run.
- Completion: isolated commit `perf(map): add diagnostic cost accounting` và comparable B3 baseline seal.

### `U3` — run-scoped diagnostic accumulator

- Goal: normalize/decode each diagnostic once per run, direct bucket lookup, giữ exact first-duplicate/full-stable-sort/immediate-materialization/six-field fail-closed behavior.
- Owner: Coder; independent Supervisor; Main.
- Allowed production surfaces: `internal/graphhealth/diagnostics.go`; resolution lifetime/injection owners `internal/resolution/emit.go`, `internal/resolution/outcome.go`, `internal/resolution/resolve.go` chỉ trong exact Architect contract; focused graphhealth/resolution tests.
- Non-goals: deferred graph publish, global/cross-run memo, relaxed validation/classification, changed comparator, concurrent writers, Unit 4 decoder rewrite.
- Work: run-scoped appender production first; then differential sequence oracle compares return value và node property after every append, including pre-existing duplicate buckets, later sort-key mutation, new bucket full sort, missing node, graph identity drift, concurrent/second writer rejection; rerun complete graph-health matrix including `3/3`, `2/2`, `26/26`, `1/1`, `1/1` and six equality groups; full build; full A/B/native parity.
- Evidence IDs:
  - `E-PERF-U3-IMPACT-001`, `E-PERF-U3-IMPL-001`, `E-PERF-U3-BUILD-001`, `E-PERF-U3-NEAR-001`;
  - `B-PERF-U3-AB-WARM-001` and conditional `B-PERF-U3-AB-COLD-001`;
  - `E-PERF-U3-COST-001`, `E-PERF-U3-MATRIX-001`, `E-PERF-U3-EQ-001`, `E-PERF-U3-DETERMINISM-001`, `E-PERF-U3-FRESHNESS-001`, `E-PERF-U3-FAILURE-001`, `E-PERF-U3-PUBLICATION-001`, `E-PERF-U3-RESOURCE-001`;
  - `E-PERF-U3-SUP-001`, `E-PERF-U3-DETECT-001`, `E-PERF-U3-COMMIT-001`.
- Stop/rollback: first-duplicate, full-sort, immediate materialization, six-field equality, missing-node or fail-closed behavior drift; second writer accepted; memory vượt `O(D)`; any graph/native mismatch/noise-only/resource regression.
- Completion/transition: Supervisor PASS + isolated commit `perf(graphhealth): accumulate diagnostics per run`; then mandatory `R3`.

### `R3` — rebaseline and Unit 4 decision

- Evidence: `B-PERF-R3-UNPROFILED-001`, `B-PERF-R3-COUNTERS-001`, `B-PERF-R3-PROFILE-001`, `E-PERF-R3-COSTMAP-001`, `E-PERF-R3-PARETO-001`, `E-PERF-R3-SEAL-001`.
- Decision:
  - decoder full-note second pass remains measured Pareto ngoài noise/materiality -> open U4;
  - accumulator đã loại phần lớn calls và decoder không còn material -> record `E-PERF-U4-NOT-ACTIVATED-001`, không code U4.
- Commit: evidence-only `perf(bench): rebaseline after diagnostic accumulator`.

### `U4` — conditional presence-aware single-pass decoder

- Goal: one full-note traversal while preserving exact `(outcome, structured, valid)` contract and JSON presence semantics.
- Owner: Coder; Supervisor; Main.
- Allowed: `internal/graphhealth/diagnostics.go` và focused tests only, trừ khi Architect addendum authorizes more.
- Non-goals: trust status marker, relax six-field nested validation, skip malformed-present handling, change unknown/duplicate-key/null semantics, memoize across calls/runs.
- Work: production decoder first; then adversarial old/new differential for missing vs null, wrong types, malformed JSON, duplicate keys/last-value behavior, unknown fields, nested raw target/authority, no-marker legacy; complete graph-health matrix; full build; A/B/native/output/failure parity.
- Evidence IDs: `E-PERF-U4-IMPACT-001`, `E-PERF-U4-IMPL-001`, `E-PERF-U4-BUILD-001`, `E-PERF-U4-NEAR-001`, `B-PERF-U4-AB-WARM-001`, conditional `B-PERF-U4-AB-COLD-001`, `E-PERF-U4-COST-001`, `E-PERF-U4-MATRIX-001`, `E-PERF-U4-EQ-001`, `E-PERF-U4-FAILURE-001`, `E-PERF-U4-RESOURCE-001`, `E-PERF-U4-SUP-001`, `E-PERF-U4-DETECT-001`, `E-PERF-U4-COMMIT-001`.
- Stop/rollback: one decode triple/normalized diagnostic/output mismatch; invalid marker no longer fail closed; direct cost noise-only; resource/whole-wall regression.
- Completion: Supervisor PASS + isolated commit `perf(graphhealth): decode structured outcomes once`.

### `RP` — fresh Pareto loop

Sau U4 commit hoặc U4 N/A, chạy `B-PERF-RP-ABSOLUTE-001` và `E-PERF-RP-PARETO-001`. DB, snapshot, parse hoặc post-run chỉ trở thành next unit khi current absolute named wall/direct counters đặt nó ở Pareto frontier và Architect cung cấp exact solution contract cho owner đó. Không tự kích hoạt candidate từ old report.

Mỗi later unit lặp template: plan delta -> impact -> code first -> tests -> full build -> nearest boundary -> A/B/equivalence -> Supervisor -> detect/cleanup -> isolated commit -> rebaseline. Candidate DB không được giảm rows/columns/evidence; snapshot phải byte-identical và giữ atomic publication; parse concurrency phải prove parser thread safety, canonical reduce order và earliest-input-order failure; post-run phải có named boundaries trước solution.

## 9. Dirty seven-path anchor và commit isolation

Current input anchor:

| Path | Bytes | SHA-256 |
|---|---:|---|
| Child 06 actual-status ledger | `65,194` | `A30538D7D13BE6E3AED4CFA8965F3BA22CB716543CA35670267FC745FE23390B` |
| Child 06 benchmark ledger | `22,207` | `34405EAE2861D56ED3B304475BBF40D53BA2C831D3E2F43AF7F775A9DCED2A53` |
| Child 06 evidence ledger | `138,757` | `07997D8345D79847DC1C08CDBAE7A237E84A8EB1133FCDA4EF8EAFABB5553936` |
| Child 06 plan | `85,442` | `B125AAE5CB7C163DAEAD6805E2278EBC29D679390F0185BBB99E4934701DD492` |
| `internal/graphhealth/diagnostics.go` | `26,884` | `518BD06FF5F859BD42AD7B3B9FF8E3F9F4E7235720BC5E568E2C4B2DA72BD2AE` |
| `internal/graphhealth/p6d_outcome_projection_test.go` | `19,336` | `6E90E655FFCCF5A5BD3D398D95EB5B6B7F77754CAAF35DAAF4D9609248A4E713` |
| `reports/coder/rp_coder_260823_032449_by_gpt-5_p6d_graph_health_projection_blocked.md` | `20,125` | `6E8AB2DBC96B43C16FDC7102FCA89D5096CD76D6B597D162C732FED533F72420` |

Canonical aggregate: `377,945` bytes / SHA-256 `E4FF42DE8AA1941010AFFF5DD38B6789AE71168E3EBF9F59D8EECF735D5BD044`.

Isolation protocol cho M0/U1/U2:

1. Seal exact anchor aggregate và từng path trước edit, trước A/B, sau validation, trước stage và sau commit.
2. A/B baseline cùng dirty anchor; candidate delta chỉ allowed performance files. Nếu một tool/report làm anchor đổi, invalidate pair set.
3. Performance dùng new four-ledger plan; không ghi benchmark/evidence vào Child 06 ledgers.
4. Detect-changes lưu hai views: truthful full worktree (có ambient anchor) và exact unit manifest attribution. Không gọi ambient changes là performance changes.
5. Stage bằng explicit paths, không glob; cached manifest phải exact; assert 0/7 anchor paths staged.
6. Post-commit, performance files clean; seven anchor paths vẫn dirty/unstaged và exact identity. Preserving này không đổi verdict REJECT.

Units 3–4 không có isolation-by-exclusion vì trực tiếp sửa anchor owner. `A3` là hard stop; không có commit topology hợp lệ nào vừa giữ rejected anchor unstaged vừa commit một edit mới của cùng `diagnostics.go` mà không có explicit ownership disposition.

## 10. Evidence reuse và invalidation matrix

| Evidence | M0/U1 | U2 | R12 | U3 | U4/later |
|---|---|---|---|---|---|
| Root Cause current causal report | reuse as historical/current-run attribution | call-path family reusable; magnitudes stale after U1 | must remeasure absolute | secondary source hypothesis reusable only | magnitudes stale after each commit |
| Architecture report | invariant/solution contract reusable | reusable | rerank rule reusable | reusable if implementation stays exact | U4 conditional; later candidates need fresh contract |
| Narrow Root Cause PASS | reusable only as causal handoff | same | same | same | never implementation PASS |
| Amended graph-health REJECT | current truth; anchor identity preserved | current truth; preserved | still REJECT | source/test clearance invalid after edit; fresh full matrix/review required | invalidated again by decoder edit |
| Six-field semantic contract | expected behavior reusable | reusable | reusable | must remain and rerun fresh | must remain and rerun fresh |
| Old local graph-health command outputs | not performance acceptance; runtime changes demand fresh full parity | invalid for U2 runtime | invalid | invalid after source change | invalid after each source change |
| Full build/runtime identity | exact candidate only | cannot reuse U1 build | current committed runtime only | cannot reuse M3/U3 build | per-candidate only |
| A/B benchmark | exact pair/source/cache only | U1 B becomes U2 A if all identities match | new baseline authority | R12 only prioritizes; B3 required | new pair every unit |
| Graph/output/native hash | exact corpus/runtime only | fresh required | fresh required | fresh required | fresh required |
| Detect-changes/Supervisor report | invalid after any source/report/ledger delta | fresh per unit | fresh evidence review | fresh | fresh |
| Target/oracle/pre-post | none; target locked | none | none | none in performance campaign | only later explicit authority; never inferred |

P6-D anchor preservation là non-interference evidence, không acceptance. M0/U1/U2 tạo product runtime mới nên mọi future graph-health current-runtime/native/reader/target gate vẫn phải chạy fresh ngay cả khi two graph-health source files giữ nguyên.

## 11. Full-build, dependency và runtime authority boundary

Canonical repository source hiện định nghĩa `scripts/full-build.ps1` với package install/build, global install/discovery, launcher build và final `anvien analyze . --force`. Historical successful task còn trỏ global wrapper ngoài E. Vì campaign yêu cầu exact E-only authority, implementation không được mở cho tới một trong hai điều kiện:

1. Owner explicit authorize và contain mọi effect của exact script, kể cả dependency source/global discovery, trong boundary được phép; hoặc
2. Owner approve một repository-owned hermetic equivalent chứng minh same product contract/steps, dependency set, launcher/runtime/native outputs và final analyze, nhưng không dựa C/global wrapper.

Mỗi build bundle tối thiểu phải có per-step start/end/elapsed/exit, command argv, lock/holder preflight, dependency manifest, executable/native hashes, version/build tags/VCS dirty manifest, final analyze exact start/end/benchmark/workload và output seal. A build tổng thời gian nhưng thiếu analyze substep timestamp không dùng làm analyze baseline.

Không dùng protected/shared/quarantined cache hoặc zero-byte lock-only cache làm dependency provenance. Không install/network cho tới explicit authority. Không dùng stale packaged runtime để fill gate. `E:\cheapapp.org` không nằm trong build/corpus/reader protocol này.

## 12. Rollback, publication và commit boundaries

### Rollback triggers chung

Rollback/REJECT ngay khi có một trong các điều kiện:

- graph/output/native-reader mismatch, reduced workload hoặc skipped/fallback relationship;
- nondeterministic replay, stale result/invalidation gap;
- fail-closed, missing-node, malformed-marker, transaction/temp publication behavior yếu hơn;
- direct target cost center chỉ thay đổi trong noise hoặc không giảm;
- whole wall regress không có accepted causal disposition;
- peak memory/allocation/object/I/O vượt Owner budget;
- unrelated/dirty-anchor identity drift;
- dependency/runtime/cache/instrumentation comparability mất;
- required full build, Supervisor hoặc detect gate không PASS.

### Exact commit boundaries

| Boundary | Proposed commit subject | Must include | Must exclude |
|---|---|---|---|
| Plan | `docs(plan): add analyze cost-center optimization campaign` | new performance four-ledger set only | Child 06 ledgers, code, anchor |
| M0 | `perf(map): add granular analyze cost accounting` | exact accepted instrumentation code/tests + own ledger/evidence/report manifest | anchor, unrelated reports/artifacts |
| B0 | `perf(bench): establish analyze cost-map baseline` | sealed baseline/cost-map ledger artifacts only | source, anchor, raw disposable temp |
| U1 | `perf(resolution): reuse canonical import paths` | exact U1 code/tests + own evidence ledgers/reports | U2 code, anchor |
| U2 | `perf(resolution): index all import claims` | exact U2 code/tests + own evidence ledgers/reports | diagnostics, anchor |
| R12 | `perf(bench): rebaseline after import optimizations` | rebaseline/cost-map evidence only | source, anchor |
| A3 | graph-health owner names its own accepted commit | only independently accepted disposition | rejected exact candidate |
| M3 | `perf(map): add diagnostic cost accounting` | exact instrumentation-only delta after clean anchor baseline | U3 algorithm change |
| U3 | `perf(graphhealth): accumulate diagnostics per run` | exact U3 code/tests/evidence | U4 decoder change, unrelated Child 06 bytes |
| R3 | `perf(bench): rebaseline after diagnostic accumulator` | evidence only | source |
| U4 | `perf(graphhealth): decode structured outcomes once` | exact U4 code/tests/evidence | later Pareto changes |
| RP | `perf(bench): rerank analyze cost centers` | fresh absolute cost map/evidence | speculative implementation |

Không stage trước independent Supervisor PASS. Main stages explicit path list, verifies cached bytes/hashes and commit manifest, then post-commit verifies no leftover unit delta. Một commit không được chứa hai optimization units; như vậy mỗi unit rollback/attribution độc lập.

## 13. Possible topology và names — proposal mở, không phải choice set

Chỉ sau khi workflow trên được suy ra, topology phù hợp nhất là một standalone campaign thay vì mở successor phase của Child 06:

- Proposed plan root: `2026-08-23-analyze-granular-cost-map-and-sequential-optimization`
- Possible slice IDs/names:
  - `AP-M0` — Measurement Authority and Living Cost Map
  - `AP-B0` — Comparable Baseline and Noise Contract
  - `AP-U1` — Canonical Import Path Reuse
  - `AP-U2` — All-Import Claim Index
  - `AP-R12` — Post-Import Absolute Rebaseline
  - `AP-A3` — Diagnostic Anchor Ownership Gate
  - `AP-M3` — Diagnostic Cost Instrumentation
  - `AP-U3` — Run-Scoped Diagnostic Accumulator
  - `AP-R3` — Post-Accumulator Absolute Rebaseline
  - `AP-U4` — Conditional Single-Pass Structured Decoder
  - `AP-RP` — Fresh Pareto Selection Loop

Đây là naming proposal để Main/Owner chấp nhận hoặc đổi. Nó không đóng tập lựa chọn, không rename/reopen Child 06 và không cho phép code trước khi plan thật được tạo/review/commit. Dependency graph cross-plan ghi rõ: `AP-U1/U2` non-overlap nhưng chạy với frozen Child 06 dirty anchor; `AP-A3` là hard dependency trước mọi `diagnostics.go` edit.

## 14. Residual unknowns và stop conditions còn lại

1. Chưa có comparable unprofiled baseline/noise distribution; không có promised seconds/percentage.
2. Owner chưa chọn numeric materiality/resource threshold hoặc repetition rule; phải chọn sau A/A, trước candidate.
3. Canonical full-build E-only/dependency authority chưa được report này cấp.
4. Exact M0 instrumentation manifest/overhead chỉ biết sau fresh impact và implementation design check; nếu cần contract mới, quay lại Architect.
5. OS-cache reset có thể không có authority; khi đó không báo cold result.
6. Exact Q/I/bucket và diagnostic list distributions chưa instrument; field phải `not-instrumented` cho tới M0/M3.
7. Dirty graph-health anchor chưa có clean committed disposition; U3/U4 hiện bị hard-block tại `A3` dù U1/U2 có thể độc lập.
8. DB export/native/commit, snapshot encode/flush, parse-by-language và post-run splits chưa đủ để chọn next solution.
9. Historical regression commit/slice vẫn unverified; workflow chỉ tối ưu current measured costs.
10. Target/native/oracle gates của graph-health lifecycle không được performance evidence thay thế; target vẫn locked.

## 15. Handoff

Main cần zero-trust verify report này, quyết định có chấp nhận standalone campaign/topology hay không, rồi kích hoạt independent Supervisor theo Owner sequence. Report không tự mở plan, không cho phép implementation, không sửa Child 06 và không thay verdict REJECT của current graph-health anchor.

Planner lane dừng sau khi handoff exact path/seal cho Main.
