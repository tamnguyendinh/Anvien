# Planner Report: Nâng cấp Child 06A thành execution-ready plan

- Trạng thái: **ĐÃ ÁP DỤNG — Child 06A ready for execution tại M0; chưa có functional result**
- Thời điểm lập: `2026-08-24 20:21:15 +07:00`
- Cập nhật gần nhất: `2026-08-24 20:53:46 +07:00` — execution-control upgrade applied and scoped verification completed
- Repository duy nhất: `E:\Anvien`
- Vai trò: Planner; chuyển kết quả rà soát Child 06A thành plan/ledger có thể triển khai trực tiếp
- Next owner: assigned Child 06A executor at M0 run-contract/harness seal
- Nguồn đầu vào: nội dung phân tích do user cung cấp, mandatory execution guide và current Child 06A four-file plan set được dẫn chiếu trong báo cáo
- Boundary của report: không chạy analyze, benchmark, build, test, profile, functional measurement, graph command, lane, Git stage/commit hoặc target access; documentation-only upgrade đã sửa mandatory guide và bốn Child 06A ledgers

## 1. Kết luận điều hành

Child 06A giữ đúng topology/invariant và hiện đã có execution cursor cụ thể tại M0. Bản nâng cấp giữ nguyên:

- đúng một implementation slice `P1-A`;
- đúng một final Supervisor boundary;
- đúng một implementation commit;
- U1, U2, U3 và conditional U4 là ordered internal units;
- tối đa ba measurement lanes theo actual cost centers, không predeclare lane count.

Execution-control layer đã được thêm ngay trong mandatory guide và bốn ledger Child 06A hiện có. Lớp này làm rõ current cursor, run contract, cost-center inventory, rolling queue, checkpoint schema, candidate identity, invalidation rules, decision vocabulary và closure handoff.

Không đề xuất tạo thêm Child plan, implementation slice, fixed lane count hoặc ledger thứ năm.

## 2. Luồng vận hành hiện hành

```text
M0 RUN-CONTRACT / HARNESS SEAL
→ M1 CURRENT COST MAP
→ M2 A/A PAIR 01…10
→ M3 ARCHITECT NUMERIC SEAL
→ M4 FROZEN BASELINE-A
→ U1
→ U2
→ REBASELINE-1 / OWNERSHIP
→ U3
→ REBASELINE-2 / U4 GATE
→ U4 hoặc U4-N/A
→ DYNAMIC PARETO
→ FINAL EQUIVALENCE
→ P2-A FINAL SUPERVISOR
→ P2-B CLEANUP
→ P2-C DETECT → ONE COMMIT → CHILD 07
```

`M0` đến `M4` chỉ là checkpoint nội bộ của `P1-A`. Chúng không phải phase, implementation slice, Supervisor review hoặc commit boundary mới.

## 3. Các nâng cấp bắt buộc

### 3.1 Duy trì mandatory guide như authority của plan Child 06A

[Mandatory execution guide](E:/Anvien/reports/Investigation/rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md) xác định plan Child 06A là authority thực thi và bind active Child 06A four-file set; bốn Child 06 ledgers chỉ còn là predecessor provenance/preserve-only.

Terminology và ownership hiện hành:

| Authority | Child 06A owner | Rule |
|-----------|-----------------|------|
| implementation work | Child 06A `P1-A` | đúng một implementation slice; U1-U4 là internal units |
| final review | Child 06A `P2-A` | sole final effective Supervisor |
| cleanup | Child 06A `P2-B` | exact cleanup, không mở review mới |
| detect/commit | Child 06A `P2-C` | sole post-cleanup detect và one implementation commit |
| predecessor history | bốn Child 06 ledgers | provenance/preserve-only, không phải performance execution authority |

Exact guide filename/path được giữ nguyên để không phá handoff link. Nội dung và authority terminology bên trong guide đã được đồng bộ thành plan Child 06A trong documentation update này.

### 3.2 Cursor bắt đầu trực tiếp tại M0

Không thêm `START/HOLD` evidence, approval report hoặc approval-only lane. Active plan và assignment là đủ để executor bắt đầu M0. Current status được chốt là `READY_FOR_EXECUTION_AT_M0`; first uncompleted item là seal literal run contract/harness bằng `E1-P1A-AUTH1` và `E1-P1A-HARNESS1`.

Các gate còn lại đều là technical/data gates: cost map, A/A, Architect numeric calculation, frozen baseline-A, fresh impact, A/B, equivalence và resource proof. Chúng không phải vòng duyệt hành chính.

### 3.3 Tách measurement preflight thành các checkpoint cụ thể

Trước nâng cấp, Work Step 1 gộp harness, cost map, 20 A/A runs, baseline và Architect seal. Current plan đã tách thành M0-M4 với evidence riêng.

Evidence đã được tách thành:

- `E1-P1A-AUTH1`;
- `E1-P1A-HARNESS1`;
- `E1-P1A-COSTMAP1`;
- `E1-P1A-AA1` đến `E1-P1A-AA10`;
- `E1-P1A-ARCHSEAL1`;
- `E1-P1A-BASELINE1`.

Numeric output của Architect được ghi riêng tại `B1-P1A-ARCHRULE1`; baseline-A vẫn không chứa A/B result.

`E1-P1A-BASELINE1` giờ chỉ giữ frozen pre-edit baseline-A. A/A pair, Architect seal và baseline-A có ID riêng nên mỗi kết quả được checkpoint ngay.

### 3.4 Sửa semantics của baseline

Trước nâng cấp, wording yêu cầu “alternating unprofiled baseline” trước khi có candidate B. Current plan đã loại bỏ mâu thuẫn này.

Semantics hiện được định nghĩa chính xác:

1. A/A `20/20` đo noise trên runtime A.
2. Architect seal materiality/resource rules từ numeric table.
3. Freeze `baseline-A` gồm executable/runtime/input/output identity.
4. U1 tạo candidate B.
5. `B1-P1A-U1AB1` mới chạy interleaved/alternating frozen-A và candidate-B.
6. Lặp lại cùng contract cho U2, U3, conditional U4, Pareto candidate và final candidate.

Theo đó:

- `B1-P1A-BASELINE1` là frozen A reference trước edit;
- alternating A/B comparison thuộc từng `B1-P1A-*-AB1`;
- baseline-A phải có exact artifact/run identity và không được biến thành stale/global/cross-run cache.

### 3.5 Thêm current control snapshot và rolling queue

Actual status đã có ba bảng sống dưới đây.

#### Current Control Snapshot

| Field | Required value |
|-------|----------------|
| Current cursor | exact internal gate/work item |
| First uncompleted item | one actionable item, không phải broad phase name |
| Candidate identity | current accepted/uncommitted candidate checkpoint |
| Harness / capture / baseline | exact IDs hoặc `none` |
| A/A progress | valid runs/pairs; invalid attempts không tăng count |
| Current ranking snapshot | exact cost-map/rebaseline ID |
| Invalidated gates | exact E/B IDs và lý do |
| Next owner/action | exact owner và next command/transition |

#### Current Cost-Center Inventory

| Required field |
|----------------|
| cost-center ID |
| internal unit |
| owner symbol/file |
| complete caller/call path |
| work-unit type |
| metric set |
| equivalence boundary |
| capture ID |
| absolute current cost/rank |
| lane/owner/state |
| evidence/benchmark IDs |
| next action |

#### Rolling Queue

| Lane owner | Cost center | Queue state | Start | Checkpoint paths | Release condition |
|------------|-------------|-------------|-------|------------------|-------------------|

Allowed queue states:

```text
QUEUED → RUNNING → RESULT_READY → CHECKPOINTED → CLOSED
                         ↘ HELD
```

Chỉ cost-center row đủ owner, complete call path, metric và equivalence boundary mới được dispatch. Tối đa ba measurement rows ở trạng thái `RUNNING`; chỉ một owner được sửa production/test bytes trong shared worktree.

### 3.6 Chuẩn hóa decision và terminal disposition

Guide dùng hai vocabulary khác nhau:

- decision: `KEEP`, `REWORK`, `ROLLBACK`, `NOT_MATERIAL`, `BLOCKED`;
- terminal disposition: `OPTIMIZED`, `REJECTED`, `NOT_MATERIAL`, `BLOCKED`.

Mapping bắt buộc đã được ghi:

| Decision | Transition | Terminal disposition |
|----------|------------|----------------------|
| `KEEP` | giữ candidate sau direct + end-to-end + equivalence/resource gates | `OPTIMIZED` |
| `REWORK` | quay lại đúng owning cost center; không cấp unit sau | nonterminal |
| `ROLLBACK` | restore accepted candidate và chứng minh byte/candidate identity | `REJECTED` |
| `NOT_MATERIAL` | không production edit hoặc rollback thử nghiệm | `NOT_MATERIAL` |
| `BLOCKED` | HOLD, ghi exact blocker/next authority | `BLOCKED` |

Actual status đã tách riêng:

- Truth Status;
- Gate State;
- Lane State;
- Decision;
- Terminal Disposition.

Không dùng giá trị ghép như `blocked/conditional` trong một cột status.

### 3.7 Chốt run contract và artifact convention

Plan hiện có `P1-A Run Contract`; executor seal literal current binary, command, corpus, cache policy và output path tại M0 thay vì suy đoán từ historical report. Contract gồm:

- exact command, cwd và executable identity;
- corpus, options, environment whitelist, cache/storage và instrumentation regime;
- warm-up, timeout, stopping, retry và invalid-run rules;
- output/digest commands;
- raw JSON và concise Markdown paths;
- accepted repo-local durable artifact root `reports/benchmark/child-06a/` và ephemeral root `.tmp/child-06a/`;
- `run_id`, `capture_id`, `candidate_id`, `harness_id`, `cost_map_version`;
- exit and validity predicates;
- artifact cleanup/preservation rule.

Current plan còn có executable PowerShell command card cho M0 identity/help checks, M1 profiled cost-map capture và canonical full build; các lệnh chưa được chạy trong documentation turn này.

Một accepted capture có thể phục vụ nhiều cost centers. Không rerun capture chỉ để tạo report riêng.

### 3.8 Bổ sung ledger schema có thể điền trực tiếp

| File | Đã bổ sung |
|------|-------------|
| Plan | internal state machine, role/write-authority matrix, operational gate matrix, invalidation/reuse matrix |
| Evidence | append-only result checkpoint, `AA1..AA10`, invalidation log, exact full evidence IDs, handoff checkpoint |
| Benchmark | A/A pair table, per-cost-center metric rows, rebaseline ranking rows, resource-budget table, final A/B table |
| Actual status | control snapshot, cost-center inventory, rolling queue, candidate/invalidation state |

Benchmark hiện giữ aggregate A/A row và mười pair rows `P01..P10`, đáp ứng immediate checkpoint sau từng pair.

#### Append-only result checkpoint schema

Mỗi measurement hoặc implementation result phải ghi:

- evidence ID và recorded-at;
- stage/internal unit/cost-center ID;
- candidate checkpoint và capture/reused-capture ID;
- run/harness/input/output identities;
- exact invocation, exit, start/end;
- raw JSON và concise Markdown paths;
- exact benchmark rows;
- workload/output parity;
- validity hoặc invalid reason;
- decision và terminal disposition;
- actual-status refresh row;
- next exact owner/action.

#### A/A pair schema

Mỗi pair phải có:

- pair ID và accepted run IDs;
- candidate/capture identity;
- metric group IDs;
- pair delta/noise;
- output identity match;
- validity;
- cumulative valid run count;
- exact evidence và actual-status refresh.

Invalid attempts phải có evidence riêng nhưng không tăng valid run count.

#### Rebaseline ranking schema

Mỗi rebaseline phải có one-row-per-cost-center:

- candidate/capture identity;
- rank và previous rank;
- owner/call path;
- absolute wall/CPU/allocation/I/O/wait/work units;
- explicit process residual;
- end-to-end contribution;
- materiality decision;
- disposition/next unit;
- exact evidence.

### 3.9 Chuyển impact evidence về đúng thời điểm

Trước upgrade, `E1-P1A-IMPACT1` nằm trong preflight target. Current plan giữ nó làm legacy-compatible summary và dùng fresh per-owner impact ngay trước actual edit.

Impact evidence đã chuyển thành per-owner IDs:

- `E1-P1A-U1IMPACT1`;
- `E1-P1A-U2IMPACT1`;
- `E1-P1A-U3IMPACT1`;
- conditional `E1-P1A-U4IMPACT1` và `E1-P1A-PARETOIMPACT1` suffix family.

Impact phải nằm trong cùng authorized coding assignment, chạy đúng trước edit khi repo rule yêu cầu. Không tạo lane `impact-only → report → STOP`.

### 3.10 Hoàn thiện closure và handoff

Closure evidence đã bổ sung:

- `E1-P1A-FINALBUILD1`;
- final equivalence manifest theo từng surface;
- failure/transaction/temp/publication case matrix;
- resource-budget result matrix;
- candidate/staging manifest;
- invalidation log;
- append-only Main handoff record.

Evidence ledger hiện có handoff rules và append-only handoff row schema.

Mỗi handoff record phải có:

- outgoing Main và successor/ACK;
- guide FULL RAW ACK;
- topology position và first uncompleted item;
- inventory/ranking snapshot;
- active, queued và held lanes cùng owner;
- completed raw/Markdown paths;
- exact evidence/benchmark rows;
- candidate identity và invalidated gates;
- next exact command/transition;
- warm/hard rotation times.

Closure order phải được chốt nhất quán:

```text
P2-A sole final effective Supervisor
→ P2-B exact cleanup, no review
→ P2-C sole post-cleanup detect
→ one implementation commit
→ Child 07 handoff
```

Nếu P2-B đổi production/test bytes, final equivalence và P2-A bị invalidated; work quay lại exact owning unit. Không tạo Supervisor boundary hoặc commit boundary thứ hai.

## 4. Role và write-authority matrix hiện hành

| Role | Owns | Must not own |
|------|------|--------------|
| Owner | campaign scope change và explicit later PAUSE/STOP ngoài repo authority | intermediate technical choices đã được plan/Architect seal |
| Main | cursor, dispatch, monitoring, ledger updates, transition, one-commit boundary | Coder, QA, Architect hoặc Supervisor work |
| Measurement worker | exact capture/cost-center measurement và raw JSON/Markdown checkpoint | production/test edits hoặc ledger transition decision |
| Coder | một active internal unit/cost center, production-first rồi tests | parallel overlapping implementation hoặc final acceptance |
| Architect | numeric materiality/resource seal sau A/A `20/20` | report-chain audit hoặc implementation acceptance |
| Supervisor | sole final effective P2-A review | intermediate unit reviews hoặc cleanup review mới |

Routine measurement/result updates nên sửa evidence, benchmark và actual status ngay. Plan chỉ đổi khi fresh evidence thay đổi scope, owner, contract, major work steps hoặc transition authority.

## 5. Thứ tự thực hiện hiện hành

### P0 — execution entry đã hoàn tất

1. Mandatory guide đã đồng bộ với plan Child 06A.
2. Current control snapshot và cursor M0 đã được ghi.
3. Run-contract/artifact convention đã có schema và exact repo-local roots; literal runtime values được M0 seal trước run 1.
4. Cost-center inventory và rolling queue đã có schema điền trực tiếp.
5. A/A pair evidence đã tách thành AA01..AA10; baseline semantics đã sửa.
6. Decision/gate/lane/terminal state model đã chuẩn hóa.

### P1 — trước U1 production edit

1. Seal A/A `20/20` và Architect numeric rules.
2. Freeze baseline-A reference.
3. Populate exact U1 cost-center owners/call paths.
4. Run fresh per-owner file-detail/impact ngay trước edit.
5. Record candidate/invalidation matrix và first exact coding assignment.

### P2 — trước final closure

1. Hoàn thiện final build/equivalence/failure/resource manifests.
2. Chứng minh Pareto exhausted bằng terminal disposition cho mọi measured material cost center.
3. Chạy sole P2-A Supervisor boundary.
4. P2-B cleanup without production/test mutation.
5. P2-C post-cleanup detect, exact staging, one commit và Child 07 handoff.

## 6. Authority files đã được nâng cấp

1. `reports/Investigation/rp_huong_dan_lam_viec_de_slice_p6e_tien_hanh_nhanh_hon.md` — Child 06A terminology, M0 entry, run contract, checkpoint, state, decision và closure order đã đồng bộ
2. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`
3. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md`
4. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md`
5. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md`

Roadmap, Child 06 và Child 07 chưa cần thay đổi cho execution-control enhancement này, trừ khi một future Owner decision thay đổi topology hoặc successor contract.

## 7. Non-goals và acceptance của bản nâng cấp

Bản nâng cấp đã áp dụng không:

- chạy measurement hoặc production implementation trong documentation turn này;
- thay đổi topology `Child 06 → Child 06A → Child 07`;
- chia P1-A thành nhiều implementation slices;
- tạo intermediate Supervisor hoặc per-unit commit;
- predeclare số cost centers hoặc lane count;
- coi historical/profile values là comparable baseline;
- bịa current runtime identity, baseline number, threshold, PASS hoặc speedup; executable command card được lấy từ current repo command surface và vẫn phải seal ở M0;
- rerun aggregate evidence, report-chain audit hoặc evidence-about-evidence;
- sửa roadmap, Child 06, Child 07, production code, tests hoặc runtime artifact.

Bản nâng cấp đạt mục tiêu tài liệu khi:

- mọi gap đã được đóng tại exact guide/plan/ledger surface;
- one-slice/one-Supervisor/one-commit topology được giữ nguyên;
- first executable cursor được xác định trực tiếp là M0 run-contract/harness seal, không phải U1 và không có approval gate bổ sung;
- không có functional completion claim.

## 8. Current next action

Bản execution-control upgrade đã được áp dụng trên đúng năm authority files ở mục 6. Verification cần xác nhận:

- guide terminology và Child 06A mapping nhất quán;
- internal cursor/gate order nhất quán giữa plan, evidence, benchmark và actual status;
- A/A `AA1..AA10`, Architect seal, baseline-A và A/B IDs liên kết chính xác;
- no template placeholders;
- all cross-links resolve;
- `git diff --check` PASS trên exact edited paths.

Executor tiếp theo đọc active plan và bắt đầu trực tiếp tại M0: seal exact run contract/harness, sau đó M1 current cost map. U1 chỉ mở sau A/A `20/20`, Architect numeric seal, frozen baseline-A và fresh per-owner impact; không có bước duyệt hành chính xen giữa.
