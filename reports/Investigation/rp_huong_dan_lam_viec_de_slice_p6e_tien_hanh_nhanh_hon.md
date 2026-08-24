# Guide: Hướng dẫn làm việc để plan Child 06A tiến hành nhanh hơn

Status: `BINDING_MANDATORY_MAIN_GUIDE`

Scope: `Child 06A Plan: Accelerate Analyze Without Sacrificing Accuracy`

Owner authority date: `2026-08-24`

Last revised: `2026-08-24` — execution-control layer synchronized to the active Child 06A four-file plan set

Current execution state: `READY_FOR_EXECUTION_AT_M0`. Child 06A chưa có implementation candidate, A/A run hoặc baseline; executor bắt đầu trực tiếp tại M0 run-contract/harness seal.

## 1. Authority và quy tắc bắt buộc

Tài liệu này là hướng dẫn điều hành bắt buộc đối với mọi Visible Main Orchestration đang nắm authority của campaign Anvien Graph Accuracy / Analyze Performance.

Mọi Main hiện tại và mọi successor Main sau handoff phải:

1. Đọc FULL RAW tài liệu này trước khi giao lệnh hoặc chuyển bước trong plan Child 06A.
2. Áp dụng tài liệu này cùng FULL RAW `AGENTS.md`, `working-rules`, `orchestration`, active Child 06A plan và bốn Child 06A ledgers. Bốn Child 06 ledgers là predecessor provenance/preserve-only, không phải performance execution authority.
3. Ghi exact path của guide này vào mọi Main handoff:
   `E:\Anvien\reports\Investigation\rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md`.
4. Yêu cầu successor Main ACK rằng đã đọc FULL RAW guide và xác định đúng first uncompleted Child 06A work item.
5. Không được dùng handoff, summary, memory hoặc auto-compact thay cho guide này.
6. Không được restart một measurement, optimization unit hoặc validation đã hoàn thành nếu source/runtime/boundary tương ứng không bị invalidated.

Executor dùng guide này cùng active Child 06A four-file set và bắt đầu tại M0. Một explicit PAUSE/STOP ban hành sau đó có hiệu lực tức thời và được ghi như operational state, không phải phase mới.

## 2. Mục tiêu duy nhất

Làm `anviens analyze` chạy nhanh hơn trên end-to-end runtime thật trong khi giữ nguyên:

- workload và corpus;
- accuracy và semantic completeness;
- graph correctness;
- ordered outcomes, diagnostics và evidence;
- deterministic output;
- freshness invalidation/restore;
- failure, transaction, temporary-file và publication behavior;
- Graph JSON, Ladybug và native-reader parity.

Tài liệu, report và evidence chỉ ghi lại kết quả của công việc. Chúng không phải sản phẩm chính và không được biến thành audit gate.

## 3. Topology bắt buộc của plan Child 06A

Plan Child 06A sở hữu đúng một implementation slice `P1-A` và một final implementation commit boundary:

```text
M0 run-contract / harness seal
→ M1 current cost map
→ M2 A/A pair 01..10
→ M3 Architect numeric seal
→ M4 frozen baseline-A
→ U1 canonical-path reuse
→ U2 all-import claim index
→ rebaseline / diagnostic ownership
→ U3 diagnostic accumulator
→ rebaseline
→ conditional U4 one-pass decoder
→ dynamic Pareto optimization
→ final equivalence
→ one final Supervisor
→ exact cleanup
→ one post-cleanup detect
→ one implementation commit
```

M0-M4, U1, U2, U3 và U4 là ordered internal checkpoints/units bên trong `P1-A`, không phải các phase, slice, Supervisor boundary hoặc commit riêng.

State machine bắt buộc:

```text
READY_FOR_EXECUTION_AT_M0
→ RUN_CONTRACT_SEALED
→ COSTMAP_READY
→ AA01_ACCEPTED ... AA10_ACCEPTED
→ ARCH_RULES_SEALED
→ BASELINE_A_FROZEN
→ U1_IMPACT_READY
→ U1_AUTHORIZED
```

Mỗi transition ghi exact evidence/benchmark paths, candidate/capture identity, invalidated gates và next exact owner/action, rồi chuyển thẳng sang technical gate kế tiếp.

## 4. Chia lane theo thực thể code thật

Không được đặt trước số lượng lane theo ví dụ như 8, 10 hoặc 20.

Một measurement/optimization work item được định nghĩa bởi:

```text
cost-center ID
+ owner symbol
+ caller và complete call path
+ work-unit type
+ metric cần đo
+ output/equivalence boundary
```

Quy tắc chia:

- Một cost center đi qua nhiều file vẫn có thể thuộc một lane.
- Một file chứa nhiều cost center độc lập có thể tạo nhiều lane.
- Số lane `N` bằng số cost center thật được active plan và current cost map chứng minh.
- Không mở lane cho một filename hoặc phần trăm lịch sử nếu chưa có owner/caller/call-path thật.
- Không biến DB, snapshot, parse hoặc post-run thành lane bắt buộc; chúng chỉ mở khi fresh Pareto evidence chứng minh material.

## 5. Các logical unit lấy trực tiếp từ plan

### U1 — Canonical-path reuse

Initial owner boundary:

- `internal/resolution/indexes.go`;
- `internal/resolution/resolve.go`;
- `internal/resolution/export_resolution.go`;
- exact callers/symbols được current source/cost map chứng minh.

Measurement work items có thể bao gồm canonicalization khi xây workspace, call/access/type resolution và export traversal. Main chỉ tạo lane cho cost center thật.

### U2 — Run-scoped all-import claim index

Contract:

```text
{canonical source file path, exact local name}
    -> []original w.imports index
```

Measurement work items lấy từ exact index construction và exact consumers của resolved/unresolved import claims. Bucket phải giữ original order; index không được cache final resolution answers.

### U3 — Diagnostic accumulator

Owner chính là `internal/graphhealth/diagnostics.go` cùng exact lifecycle callers được current evidence chứng minh.

Cost centers có thể gồm:

- normalization của existing diagnostic list;
- normalization/decode của incoming diagnostic;
- duplicate-bucket lookup;
- full stable sort;
- first-match index rebuild;
- immediate node-property materialization.

### U4 — Conditional one-pass decoder

Chỉ mở khi post-U3 rebaseline chứng minh decoder cost vẫn material theo numeric rule đã được seal.

Cost centers có thể gồm wire decode, nested authority decode, six-field validation và malformed/absent paths. Nếu không material, ghi `U4 N/A` và không code.

### Dynamic Pareto work

DB, snapshot, parse và post-run chỉ là candidate families. Fresh rerank quyết định có bao nhiêu actual cost centers và có mở lane hay không.

## 6. Measurement workflow

### 6.1 Dùng aggregate evidence hiện có

Không restart từ đo tổng nếu current aggregate/profile evidence vẫn hợp lệ. Dùng nó để drill down trực tiếp vào logical units của plan.

Historical/profiling numbers chỉ dùng để chọn nơi đo sâu; chúng không tự trở thành baseline hoặc speedup claim.

### 6.2 Seal run contract trước current cost-center inventory

First executable item là `E1-P1A-AUTH1` cùng `E1-P1A-HARNESS1`. Trước run đầu tiên, Main phải ghi một run contract đầy đủ với:

- exact working directory và executable/runtime/native/dependency identity;
- exact command, corpus, options và environment whitelist;
- cache/storage/instrumentation regime;
- warm-up, pair ordering, timeout, stopping, retry và invalid-run rules;
- output/digest commands và workload/output validity predicates;
- raw JSON/concise Markdown paths và cleanup/preservation rule.

Durable result root là `E:\Anvien\reports\benchmark\child-06a\`. Directory name bằng sealed checkpoint ID, file stem bằng sealed run ID; ví dụ AA01 run 1 dùng `reports/benchmark/child-06a/aa01/aa01-a1.json` và cùng stem `.md`. Ephemeral cache/debug data chỉ được nằm dưới `E:\Anvien\.tmp\child-06a\` và không được dùng làm durable evidence.

Current command/corpus/cache/protocol value chưa được seal. Mọi field chưa biết phải ghi literal `none` hoặc `not-instrumented`; không đoán từ historical/protected report.

### 6.3 Lập current cost-center inventory

Main duy trì một bảng sống:

| ID | Unit | Owner symbol/file | Complete caller/call path | Work unit | Metric set | Equivalence boundary | Capture ID | Absolute cost/rank | Lane owner/state | Evidence/benchmark | Next action |
|----|------|-------------------|---------------------------|-----------|------------|----------------------|------------|--------------------|------------------|--------------------|-------------|

Chỉ những hàng có owner/call path và phép đo cụ thể mới được dispatch.

Không có active cost-center row trước M1. Initial filenames/U1-U4 families chỉ là seed candidates, không phải dispatchable cost centers; M0 và M1 được mở tuần tự ngay khi executor nhận plan.

### 6.4 Concurrency

- Tối đa `3` measurement lanes RUNNING cùng lúc theo current Owner direction.
- Dùng rolling queue: lane nào hoàn thành và ghi checkpoint thì Main cấp slot cho work item tiếp theo ngay.
- Không chờ toàn batch mới ghi report/evidence.
- Implementation trên shared worktree chạy từng optimization unit một để attribution rõ và tránh conflict.
- Nếu nhiều lane dùng chung một accepted raw runtime/profile capture, mỗi lane đọc đúng cost center của mình; không rerun cùng capture chỉ để tạo report riêng.

Lane state dùng đúng vocabulary:

```text
QUEUED → RUNNING → RESULT_READY → CHECKPOINTED → CLOSED
             ↘ HELD
```

Slot chỉ được release sau `CHECKPOINTED`. Tối đa ba measurement rows ở `RUNNING`; chỉ một owner được sửa production/test bytes trên shared worktree.

### 6.5 Metrics bắt buộc theo khả năng instrumentation

Mỗi cost center ghi:

- calls và work-unit count;
- items/bytes visited;
- hits, misses và early-exit position khi phù hợp;
- process/span wall và CPU;
- allocation bytes/objects;
- live/peak/retained memory khi phù hợp;
- read/write bytes/operations và wait/block khi phù hợp;
- histogram schema và buckets;
- tỷ lệ đóng góp vào current end-to-end runtime;
- exact workload/output identities cần cho equivalence.

Field chưa instrument được phải ghi `not-instrumented`, không ước lượng.

### 6.6 A/A protocol và baseline semantics trong active plan

Nếu active plan vẫn yêu cầu exact A/A repetitions, đó là baseline/noise protocol, không phải số code segments hoặc số lanes.

- Protocol phải được seal trong `E1-P1A-HARNESS1` trước run 1, gồm warm-up, order, cache, timeout, stopping/retry, invalid-run rules, histogram buckets, dispersion/CI estimator và exact run-validity predicate.
- Mỗi pair hoàn thành phải có checkpoint ngay bằng `E1-P1A-AA1` đến `E1-P1A-AA10`; benchmark pair rows là `B1-P1A-AA1-P01` đến `B1-P1A-AA1-P10`.
- Invalid attempts dùng monotonically increasing `E1-P1A-AARUN1` suffix family, ghi exact invalid reason và không tăng valid run count.
- Không chạy hết rồi mới tổng hợp.
- Một accepted run/capture nên cung cấp counters cho nhiều cost centers khi instrumentation cho phép.
- Một run chỉ valid khi exact identity/command/exit khớp, required metrics có giá trị hoặc literal `not-instrumented`, workload/output parity khớp và raw/Markdown/evidence/benchmark/actual-status checkpoint đã được ghi.
- Sau `20/20`, một Architect chỉ nhận runtime numeric table để ghi `E1-P1A-ARCHSEAL1` và seal materiality/resource/A-B repetition/stopping/cache rules; Architect không audit report chain và đóng lane ngay sau quyết định.

Baseline order không được đảo:

```text
A/A 20/20 trên runtime A
→ Architect numeric seal
→ freeze baseline-A identity bằng E1-P1A-BASELINE1 / B1-P1A-BASELINE1
→ production edit tạo candidate-B
→ alternating/interleaved frozen-A vs candidate-B trong benchmark của unit
```

`B1-P1A-BASELINE1` là frozen pre-edit A reference, không phải A/B result. A/B chỉ tồn tại sau candidate-B và thuộc `B1-P1A-U1AB1`, `U2AB1`, `U3AB1`, conditional `U4AB1`, Pareto rows hoặc `FINAL1`.

## 7. Report và evidence ngay sau mỗi bước

Sau mỗi measurement hoặc implementation result:

```text
command hoàn thành
→ có result
→ Coder ghi raw JSON + concise Markdown
→ hand exact paths và numbers cho Main
→ Main dùng planner ghi ngay evidence/benchmark/actual status cần thiết
→ Main quyết định transition
→ mới cấp work item tiếp theo
```

Không được chờ tới cuối turn hoặc cuối batch.

Mỗi checkpoint sử dụng schema append-only:

| Field group | Required record |
|-------------|-----------------|
| identity | evidence ID, recorded-at, stage/unit, cost-center ID, candidate checkpoint, capture/reused-capture ID |
| execution | exact invocation/exit, start/end, harness/runtime/input/output identities |
| artifacts | raw JSON, concise Markdown, exact benchmark row IDs |
| result | metrics, workload/output parity, validity/invalid reason, before/after/delta |
| transition | decision, terminal disposition when applicable, actual-status refresh ID, next exact owner/action |

Checkpoint chưa ghi đủ cả ba ledger không được xem là `CHECKPOINTED` và không release rolling-queue slot.

Minimum result record:

- cost-center/unit ID;
- source/runtime/corpus/cache/instrumentation identity;
- exact invocation và exit code;
- start/end và metrics;
- workload/output parity;
- validity state;
- before/after và delta nếu đã code;
- `KEEP`, `REWORK`, `ROLLBACK`, `NOT_MATERIAL` hoặc `BLOCKED`;
- exact next owner/action.

Không audit byte count, LF count, hash hoặc wording của report. Chỉ output/runtime artifact hash cần cho product equivalence mới là technical evidence.

## 8. Optimization loop

Sau khi có đủ current cost-center results cho unit đang mở:

1. Main rank theo absolute current cost, không dùng old percentages.
2. Chọn cost center lớn nhất nằm trong ordered unit hiện tại.
3. Giao lại đúng owner Coder.
4. Chạy fresh graph/file-detail/impact đúng một lần ngay trước edit nếu repo rule yêu cầu; đây là pre-edit command trong coding assignment, không phải một audit-only lane hoặc STOP boundary.
5. Sửa production code trước.
6. Sau khi production đúng mới sửa/thêm tests.
7. Chạy full build trước validation.
8. Đo lại exact cost center.
9. Đo lại end-to-end runtime và equivalence boundary bị ảnh hưởng.
10. Ghi report và evidence ngay.
11. Main quyết định:
    - `KEEP`: nhanh hơn và equivalence/resource gates pass;
    - `REWORK`: invariant cụ thể còn sửa được;
    - `ROLLBACK`: regression, nondeterminism hoặc không có gain;
    - `NOT_MATERIAL`: không còn đáng code.
12. Re-rank current costs rồi chuyển cost center tiếp theo.

Không mở unit sau trước khi ordered unit hiện tại đạt result rõ ràng.

Decision và terminal disposition phải map chính xác:

| Decision | Required transition | Terminal disposition |
|----------|---------------------|----------------------|
| `KEEP` | giữ candidate chỉ sau direct + end-to-end A/B và equivalence/resource gates | `OPTIMIZED` |
| `REWORK` | quay lại đúng owning cost center; không mở unit sau | nonterminal |
| `ROLLBACK` | restore accepted candidate và chứng minh candidate/source identity | `REJECTED` |
| `NOT_MATERIAL` | không production edit hoặc rollback thử nghiệm | `NOT_MATERIAL` |
| `BLOCKED` | HOLD, ghi exact blocker và next authority | `BLOCKED` |

Truth status, gate state, lane state, decision và terminal disposition là năm field riêng; không dùng compound value như `blocked/conditional`.

## 9. Definition of exhausted optimization work

Một cost center chỉ kết thúc khi thuộc một trạng thái:

- `OPTIMIZED`: change được giữ và measured speedup pass;
- `REJECTED`: đã thử nhưng không thể giữ vì regression hoặc không nhanh hơn;
- `NOT_MATERIAL`: current measurement không biện minh cho code;
- `BLOCKED`: exact blocker và next authority được ghi rõ.

Dynamic Pareto loop dừng khi không còn measured material cost center trong Child 06A plan boundary hoặc mọi cost center đã có một trong các disposition trên.

## 10. Final closure

Chỉ sau khi optimization work exhausted:

1. Full build cuối.
2. Comparable end-to-end baseline/final measurement.
3. Exact graph/output/native-reader equivalence.
4. Determinism replay.
5. Freshness invalidation/restore.
6. Failure/transaction/temp/publication parity.
7. Resource-budget check.
8. Một independent Supervisor review toàn bộ stable Child 06A `P1-A` candidate.
9. REJECT chỉ quay lại exact owning cost center/invariant.
10. Main dùng planner cập nhật ledgers.
11. Exact Child 06A `P2-B` artifact cleanup disposition; cleanup không được mở review mới.
12. Fresh post-cleanup `detect-changes` trong Child 06A `P2-C` trước commit.
13. Một isolated Child 06A implementation commit.

Không có intermediate Supervisor cho từng U1/U2/U3/U4.

## 11. Trách nhiệm của Main

Main phải:

- hiểu unified state của plan Child 06A;
- duy trì cost-center inventory, rolling queue và ranking;
- giữ tối đa ba measurement lanes active;
- monitor actual commands, writes, code diff và runtime results;
- chặn lane đọc/audit tài liệu ngoài nhu cầu thực thi;
- ghi plan/evidence/benchmark ngay sau mỗi result;
- không passive-wait khi còn work item dispatchable;
- quyết định transition, KEEP/REWORK/ROLLBACK/NOT_MATERIAL;
- route đúng một final Supervisor;
- bảo vệ one-commit boundary;
- chuẩn bị rotation đúng hạn mà không làm gián đoạn functional lanes.

Main không làm Coder, QA, Architect hoặc Supervisor work.

| Role | Write/decision authority | Forbidden ownership |
|------|--------------------------|---------------------|
| Owner | campaign scope change và explicit later PAUSE/STOP ngoài repo authority | routine technical transition đã được plan/seal quyết định |
| Main | cursor, dispatch, monitoring, ledger refresh, transition, one-commit boundary | Coder, QA, Architect, Supervisor work |
| Measurement worker | exact capture/cost-center run và raw JSON/Markdown | production/test edits hoặc ledger transition decision |
| Coder | một active internal unit/cost center; production trước tests | overlapping implementation hoặc final acceptance |
| Architect | one numeric seal after A/A `20/20` | report-chain audit, implementation acceptance |
| Supervisor | sole final effective `P2-A` review | intermediate unit review hoặc cleanup review mới |

## 12. Handoff contract bắt buộc

Mỗi outgoing Main handoff phải ghi:

- exact guide path và yêu cầu successor đọc FULL RAW;
- current Child 06A state trong topology;
- current cost-center inventory và ranking;
- completed measurement/implementation checkpoint paths;
- exact evidence/benchmark rows đã ghi;
- active/held lanes và cost-center ownership;
- current candidate code identity và invalidated gates;
- first uncompleted work item;
- next exact command/lane transition;
- warm/hard rotation times.

Interim Main handoffs dùng append-only evidence family `E1-P1A-MAINHO1` trở lên. `E2-P2C-HANDOFF1` được giữ riêng cho final Child 07 opening; nó không được dùng cho rotation nội bộ.

Successor Main phải:

1. ACK đã đọc FULL RAW guide.
2. Preserve completed checkpoints.
3. Tiếp tục first uncompleted work item.
4. Không rerun measurement, audit hoặc review đã pass nếu code/runtime/boundary không đổi.
5. Không biến re-anchor thành restart campaign.

## 13. Anti-loop — tuyệt đối cấm

- Cấm assignment `impact-only → report → STOP` khi impact nằm trong một authorized coding unit.
- Cấm audit report chain, handoff chain hoặc incident history không thay đổi.
- Cấm evidence-about-evidence.
- Cấm kiểm report hash/bytes/LF/wording như functional progress.
- Cấm reread toàn bộ plan/report chain sau mỗi result.
- Cấm tạo lane chỉ để chứng minh ledger vừa được cập nhật đúng.
- Cấm rerun PASSed measurement/build/review vì auto-compact hoặc Main rotation.
- Cấm mở fixed lane count từ một con số ví dụ.
- Cấm biến optional Pareto family thành mandatory code unit khi chưa measured material.
- Cấm batch measurement xong hết mới ghi report/evidence.
- Cấm nhiều Coder sửa overlap trong shared worktree.
- Cấm Supervisor lặp lại sau từng internal unit.
- Cấm coi documentation là deliverable chính của plan Child 06A.

## 14. Main operational checklist

- [x] Child 06A execution cursor bắt đầu trực tiếp tại M0.
- [ ] `E1-P1A-HARNESS1` đã seal exact run contract trước run 1.
- [ ] Main và successor đã đọc FULL RAW guide.
- [ ] Current plan topology và first uncompleted unit đã xác định.
- [ ] Current cost-center inventory lấy từ actual plan/code/cost map, không từ số lượng giả định.
- [ ] Rolling queue có tối đa ba measurement lanes.
- [ ] Mỗi completed result có report và ledger update ngay.
- [ ] Current ranking dùng absolute fresh costs.
- [ ] Production code đi trước tests.
- [ ] Full build đi trước validation.
- [ ] Mỗi kept unit có segment và end-to-end before/after.
- [ ] Mỗi rejected/not-material unit có disposition rõ.
- [ ] Rebaseline xảy ra đúng sau U1/U2 và U3.
- [ ] U4 chỉ mở khi material.
- [ ] Dynamic Pareto chỉ mở measured owners.
- [ ] Final equivalence và resource gates pass.
- [ ] Một final Supervisor PASS.
- [ ] Exact cleanup, post-cleanup detect và một Child 06A implementation commit hoàn tất.

## 15. Completion statement

Plan Child 06A không hoàn thành vì có nhiều report. Plan Child 06A hoàn thành khi current product runtime đo được là nhanh hơn, correctness/equivalence được giữ nguyên, mọi measured material cost center đã có disposition, independent final review PASS và exact one-commit boundary tồn tại.
