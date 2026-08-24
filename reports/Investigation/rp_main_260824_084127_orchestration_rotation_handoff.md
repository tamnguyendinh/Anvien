# Visible Main Orchestration Rotation Handoff

## Chuyển giao authority

- Campaign: Anvien Graph Accuracy / Analyze Performance.
- Outgoing Visible Main: `01a0311d-ad2f-7840-8beb-2f063e917d34`.
- Incoming visible successor: `01a03162-ccd2-76a1-910b-88ebdb8fec9c`.
- Raw official authority user-message UUIDv7 của outgoing Main: `01a03123-0f6f-72d3-b832-21b2e9b231b6`.
- Transport-source authority start: `2026-08-24T07:19:36.431+07:00`.
- Nominal warm successor target: `2026-08-24T08:04:36.431+07:00`.
- Nominal hard transfer deadline: `2026-08-24T08:19:36.431+07:00`.
- Owner tạm hoãn rotation để thảo luận bằng raw message `01a03149-d08b-7c52-9843-9b136d9ea8b7` lúc `2026-08-24T08:01:56.235+07:00`.
- Owner sau đó thay lệnh bằng raw message `01a03158-5020-7bc1-87f2-fab5c23e61b6` lúc `2026-08-24T08:17:46.400+07:00`: phải hoàn tất toàn bộ quy trình đóng P6-D rồi mới rotation; Main mới bắt đầu P6-E ngay và không cập nhật lại checklist P6-D.
- Successor creation UUIDv7: `01a03162-ccd2-76a1-910b-88ebdb8fec9c`.
- Successor creation: `2026-08-24T08:29:13.682+07:00`.
- Nominal warm lateness at successor creation: `1477.251s`.
- Nominal hard lateness at successor creation: `577.251s`.
- Hai độ trễ ngay trên được ghi theo clock danh nghĩa nhưng xảy ra trong Owner-directed rotation pause/P6-D closure override; không tự gán thành governance violation mới.
- Successor đã trả `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER`, xác nhận boundary/bans, không action trước official transfer, và hiện idle.
- Exact incoming authority time là timestamp mã hóa trong raw UUIDv7 của follow-up có exact heading `OFFICIAL AUTHORITY TRANSFER`. Không backfill thời điểm đó vào report immutable này.
- Outgoing Main phải chấm dứt tuyệt đối ngay sau official dispatch.

## Iron startup bắt buộc cho successor

Trước mọi orchestration action, đọc từng file FULL RAW, tuần tự qua EOF:

1. `E:\Anvien\AGENTS.md`
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`
4. Toàn bộ sealed handoff này, rồi verify detached seal từ official transfer message.

Trước mọi plan/ledger action, đọc tiếp FULL RAW qua EOF:

1. `E:\Anvien\.agents\skills\planner\SKILL.md`
2. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-26-anvien-graph-accuracy-roadmap.md`
3. Đủ bốn Child 06 ledgers dưới `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-07-28-06-ambient-external-resolution-and-diagnostics\`.
4. Các durable report hiện hành được liệt kê bên dưới.

Không dùng summary, compacted context, trí nhớ phiên trước hoặc tên file thay raw. Sau mọi auto-compaction, thực hiện lại raw re-anchor theo orchestration §7 trước action. `governance-rule-guard` là independent Guard hiện hữu; Main không đọc hoặc tự dùng skill đó.

## Current plan state — P6-D đã đóng hoàn toàn

- P6-D là `[x]`, `CLOSED`.
- Independent Supervisor verdict là `PASS`; procedure `CLEAN`.
- P6-D implementation commit: `81163e39718b94a509e41114cada224e8f269e36`.
  - Parent: `741335f7efa5ad215ec28361d88c462f431565ab`.
  - Tree: `b6bc216172886d5de9ef80f82c0da3ed28880f13`.
  - Subject: `feat: close P6-D runtime and reader target proof`.
  - Exact `1,883` paths: `1,871` added / `12` modified.
  - `5,866,720` insertions / `314` deletions.
- Five-document planner closure commit: `599e401548bd69e5214148a9d2ffd515a7e3305d`.
  - Parent: `81163e39718b94a509e41114cada224e8f269e36`.
  - Tree/current HEAD tree: `cf35c431962627167f83619c7512ded78705c202`.
  - Subject: `docs: close P6-D and open P6-E`.
  - Exact boundary: roadmap plus four Child 06 ledgers only.
  - `56` insertions / `43` deletions.
- P6-D checklist, evidence, benchmark, actual status, commit evidence, and P6-E transition đã được cập nhật và commit. Successor không được cập nhật lại checklist P6-D hoặc tạo vòng docs/audit mới.
- Không push.

## Sole open slice — P6-E

- P6-E là sole open Child 06 slice.
- P6-E đúng một checklist slice và một isolated commit boundary chứa ordered internal units U1-U4.
- U1-U4 không phải bốn slices, không phải bốn independent gates, và không có intermediate U2/U3 commit.
- P6-E hiện `OPEN / PREFLIGHT AND IMPLEMENTATION AUTHORIZED UNDER ORDERED GATES / NO SPEEDUP CLAIM YET`.
- Pn-A và các gate sau vẫn locked.
- Successor phải bắt đầu P6-E ngay sau mandatory startup bằng fresh authority/reality preflight, không quay lại P6-D closure.

Ordered start boundary:

1. Chạy thật một fresh `anvien analyze --force` đến exit `0` trước graph-based work. Docs-only closure commit `599e4015...` xảy ra sau graph current tại implementation HEAD, nên successor không được dùng graph cũ như graph current cho P6-E.
2. Đọc command help cần thiết, rồi chạy full current `file-detail` và upstream file/symbol impact cho exact P6-E owners trước edit. HIGH/CRITICAL là blast-radius warning, không phải edit ban.
3. Freeze exact E-only repository bytes, dirty manifest, corpus, direct runtime/native/dependency/cache/instrumentation identity.
4. Xây living granular cost map và A/A noise protocol; sau đó alternating unprofiled baseline với predeclared repetitions/stopping rule/cache regime, median/dispersion/confidence interval, exact workload/output identities, và Owner-approved materiality/resource decisions.
5. Chỉ sau các gate trên mới triển khai U1, rồi U2; seal candidate/equivalence/resource gate giữa units nhưng không commit trung gian.
6. Rebaseline current absolute costs, đóng clean diagnostic ownership, rồi U3; U4 chỉ mở nếu fresh materiality gate chứng minh decoder cost còn material.
7. Fresh Pareto rerank cho DB/snapshot/parse/post-run; không mở owner từ old percentages.
8. Sau toàn bộ required/conditional work: exact equivalence, determinism, freshness, failure/publication/resource gates, một independent Supervisor PASS, fresh detect, exact cleanup disposition, và một isolated P6-E commit theo AGENTS Rule 12. Không xin Owner cấp lại commit authority; không push nếu Owner chưa yêu cầu push.

## Accepted technical state cần preserve

- Canonical Windows PowerShell 5.1 full build PASS `1/1`; không rerun nếu chưa có source/runtime/boundary invalidation.
- Runtime `1.2.8`, `73,749,504` bytes, SHA-256 `64EECEBAF459F23EA11DA38B29C5E290F6F6B22D5EE4749E068BB927A8EEDEBB`.
- Vendor dual-shell verifier PASS: `45` modules / `1,798` files / `190,350,404` bytes; missing/extra/mismatch/license gap đều `0`.
- Reader parity `53,156/53,156`, structured differences `0`.
- Target analyze `1,359/887/0`; graph `95,762/129,716`.
- Promise, Math.max, Math.min accepted capability outcomes `3/3`.
- In-repository analyzer gaps `0/3`; resolved/gap overlap `0`.
- Target source/config preservation `203/203` và `3/3`; procedure `CLEAN`.
- Hai residual được route-only: committed P6-C3 six-field parity omission và non-Windows `scripts/ensure-ladybug-native.sh`. Không residual nào invalidates accepted Windows P6-D hoặc tự mở P6-E scope.

## Fresh post-commit graph evidence và invalidation boundary

- Sau implementation commit `81163e3...`, fresh `anvien analyze --force` thật sự exit `0`, duration `595574ms`.
- Result: `2,170` scanned / `765` parsed / `0` failed.
- Graph: `122,649` nodes / `169,300` relationships.
- File projection: `20,351` dependency edges / `479` unresolved.
- `anvien status` khi HEAD là `81163e3...` báo indexed/current và up-to-date.
- Planner closure commit `599e4015...` là docs-only; không rerun graph sau nó để tránh self-referential docs loop. Vì HEAD đã đổi, P6-E successor phải fresh analyze trước graph work.
- Không rerun canonical build, reader matrix, target analyze hoặc P6-D Supervisor chỉ vì rotation/re-anchor. Chỉ rerun khi evidence chứng minh relevant bytes/runtime/boundary invalidated.

## Durable reports và detached identities

P6-D Coder:

- Path: `E:\Anvien\reports\coder\rp_coder_260824_054528_by_gpt-5_p6d_runtime_reader_target_proof_ready_for_main.md`.
- Identity: `13,436` bytes / `176` LF / `0` CR.
- SHA-256: `E79257D64788BA31DC9B7D9FE010E893A03EDC7CF732137A40EF35089AC50EC3`.

P6-D independent Supervisor PASS:

- Path: `E:\Anvien\reports\Supervisor\rp_supervisor_260824_063140_by_gpt-5_p6d_runtime_reader_target_acceptance.md`.
- Identity: `24,954` bytes / `240` LF / `0` CR.
- SHA-256: `E71777663D76C4989E744C544F27830E7CB535B3031E98E2450EF8975317F6A3`.

P6-E plan-acceptance PASS:

- Path: `E:\Anvien\reports\Supervisor\rp_supervisor_260823_164652_by_gpt-5_p6e_current_plan_acceptance_rereview.md`.
- Identity: `10,635` bytes / `106` LF / `0` CR.
- SHA-256: `A66C2CC27A3CAED5A625C775783C860FBBC49195E76F0A898A8425C3D31B20C0`.
- Verdict chỉ là plan acceptance; implementation/speedup evidence vẫn pending.

P6-E discussion/architecture chain cần đọc trước lane design:

- Root Cause: `reports/Investigation/rp_investigation_260823_103434_by_gpt-5_analyze_performance_root_cause.md`, `32,620` bytes / `341` LF / SHA-256 `100998373E396A7E66B887357214B5952C4F9D90BC4D26558F291F9F60B950B6`.
- Architecture: `reports/system-architect/rp_system-architect_260823_110218_by_gpt-5_analyze_performance_solution_architecture.md`, `36,360` bytes / `485` LF / SHA-256 `4CD4DB7EE6D195ADDED4CB0E0879EA1C2C288B81596F56248B6624485C2957BC`.
- Execution workflow: `reports/planner/rp_planner_260823_113133_by_gpt-5_analyze_performance_execution_workflow.md`, `48,426` bytes / `494` LF / SHA-256 `B6AED404298D5399001EE3D33E2D17590DDC7D52F1025C49397B3E2610061F5B`.
- Feasibility Supervisor PASS: `reports/Supervisor/rp_supervisor_260823_123112_by_gpt-5_analyze_performance_discussion_chain_review.md`, `17,270` bytes / `140` LF / SHA-256 `0458C59504C6463F0E0C3BDAF2BCA7A1B7A4A3E03373C771647B73AF03F3E6D2`.
- Các report này chứng minh strategy/plan feasibility, không phải current cost map, comparable baseline, implementation acceptance hoặc speedup proof.

## Existing visible lanes

- Existing P6-D Coder: `01a02f10-c175-7162-ac32-6df30a4d604a`.
  - State: idle after sealed `READY_FOR_MAIN` handoff.
  - Final known cursor: `2446f274-b195-47c7-92f5-369e5fe49840:567`.
  - Preserve; không stop, pressure, duplicate cho P6-D, replace hoặc archive.
- Existing independent P6-D Supervisor: `01a030d2-3ae2-7b33-a57a-c23503703b1f`.
  - State: idle after PASS.
  - Final known cursor: `2d087646-f6a7-4928-aaa1-84aeaa93d1d3:163`.
  - Preserve; không stop, pressure, duplicate, replace, archive hoặc mở Supervisor thứ hai cho P6-D.
- Existing independent governance Guard: `01a0307e-fe39-73c2-83df-e2bb234143a5`.
  - Preserve; không stop, replace hoặc archive.
  - Guard tiếp tục là lane độc lập gửi verified warnings; Main không đọc/dùng `governance-rule-guard` skill và không tự kiêm Guard.
- P6-E là boundary mới. Sau mandatory fresh preflight, Main phải thiết kế một visible P6-E implementation/performance lane đúng orchestration ownership/boundary, monitor actual commands/files/cost-map/gates, và chỉ mở independent P6-E Supervisor sau complete candidate. Main không tự làm worker implementation.

## Current Git/worktree reality

- Sole workspace: `E:\Anvien`.
- Branch: `master`.
- HEAD: `599e401548bd69e5214148a9d2ffd515a7e3305d`.
- Parent: `81163e39718b94a509e41114cada224e8f269e36`.
- Tree: `cf35c431962627167f83619c7512ded78705c202`.
- Index: empty, `0` staged paths.
- Tracked worktree rows: exactly `1`, Owner deletion `CONTRIBUTING.md`.
- Pre-handoff status: exactly `25` protected outside-boundary rows.
- `git diff --check`: PASS.
- This handoff is uncommitted and becomes the expected 26th protected row after creation.

Exact pre-handoff protected rows:

1. `CONTRIBUTING.md` — Owner deletion, preserve.
2. `Microsoft/Windows/PowerShell/ModuleAnalysisCache` — unrelated/protected artifact.
3. `reports/Investigation/rp_main_260823_164621_orchestration_rotation_handoff.md`
4. `reports/Investigation/rp_main_260823_174805_orchestration_rotation_handoff.md`
5. `reports/Investigation/rp_main_260823_184842_orchestration_rotation_handoff.md`
6. `reports/Investigation/rp_main_260823_194346_orchestration_rotation_handoff.md`
7. `reports/Investigation/rp_main_260823_203755_orchestration_rotation_handoff.md`
8. `reports/Investigation/rp_main_260823_213118_orchestration_rotation_handoff.md`
9. `reports/Investigation/rp_main_260823_222651_orchestration_rotation_handoff.md`
10. `reports/Investigation/rp_main_260823_232232_orchestration_rotation_handoff.md`
11. `reports/Investigation/rp_main_260824_002000_orchestration_rotation_handoff.md`
12. `reports/Investigation/rp_main_260824_011928_orchestration_rotation_handoff.md`
13. `reports/Investigation/rp_main_260824_021430_orchestration_rotation_handoff.md`
14. `reports/Investigation/rp_main_260824_031452_orchestration_rotation_handoff.md`
15. `reports/Investigation/rp_main_260824_041516_orchestration_rotation_handoff.md`
16. `reports/Investigation/rp_main_260824_051851_orchestration_rotation_handoff.md`
17. `reports/Investigation/rp_main_260824_062456_orchestration_rotation_handoff.md`
18. `reports/Investigation/rp_main_260824_072456_orchestration_rotation_handoff.md`
19. `reports/Supervisor/rp_supervisor_260823_164652_by_gpt-5_p6e_current_plan_acceptance_rereview.md`
20. `reports/coder/p6d_runtime_reader_target_proof_260824_045131_raw/graph-health-summary.stderr.txt`
21. `reports/coder/p6d_runtime_reader_target_proof_260824_045131_raw/graph-health-summary.stdout.json`
22. `reports/coder/p6d_runtime_reader_target_proof_260824_045131_raw/resolution-inventory.stderr.txt`
23. `reports/coder/p6d_runtime_reader_target_proof_260824_045131_raw/resolution-inventory.stdout.json`
24. `reports/coder/p6d_runtime_reader_target_proof_260824_045314_raw/registry-index-existing.stderr.txt`
25. `reports/coder/p6d_runtime_reader_target_proof_260824_045314_raw/registry-index-existing.stdout.json`

Không normalize, clean, stash, reset, checkout, broad-stage hoặc broad-commit worktree. `third_party/`, `vendor/`, durable reports, Owner deletion và unrelated/protected work phải được giữ nguyên. P6-E eventual commit phải exact isolated accepted boundary; không push nếu Owner chưa yêu cầu.

## Binding `.tmp` correction

- Mọi path dưới literal `E:\Anvien\.tmp` là disposable reproducible scratch và có thể bị Owner xóa wholesale bất kỳ lúc nào.
- `.tmp` không bao giờ là evidence, review input, handoff/recovery authority, source of truth, acceptance dependency hoặc procedure blocker.
- Không tái tạo false per-root ownership hoặc broad-cleanup prohibition đã được rút.
- P6-E vẫn phải name exact unit-created roots/artifacts để tránh chạm unrelated work, nhưng tồn tại/drift dưới `.tmp` không thể trở thành evidence hoặc acceptance authority.

## Preserved governance findings và corrections

- Prior successor creation warm-late `69.118s` — HIGH; preserve without normalization/downgrade.
- Prior official transfer hard-late `0.518s` — HIGH; preserve without rounding to on-time.
- Previous rotation successor creation warm-late `226.627s` — HIGH.
- Roadmap patch-before-required-fresh-graph — HIGH; later revalidated on fresh graph, historical finding remains.
- Prolonged Main silence interval — MEDIUM RISK; later corrected with explicit verified/checking/blocker status.
- Current Main misread a handoff's “does not grant Git authority” wording as revoking standing AGENTS Rule 12 commit authority, pushed commit choice back to Owner, and final/yielded instead of continuing — HIGH. Corrected: standing slice-commit authority governs; P6-D implementation and planner closure commits now exist; no push occurred.
- Current Main delayed executing AGENTS Rule 12 after Supervisor PASS/detect — HIGH. Corrected by exact isolated commits `81163e3...` and `599e4015...`.
- Current Main staged/committed after an auto-compaction before completing raw re-anchor — HIGH. Correction was to freeze further action, not undo/reset either commit; full raw re-anchor through current rules/authority/planner/roadmap/four ledgers/reports was later completed and post-commit reality verified.
- Guard finding “fresh graph invocation did not exist” was formally withdrawn in full after the completed tool record showed real `anvien analyze --force`, exit `0`, duration `595574ms`. Không tính finding này là violation và không resurrect nó.
- Successor nominal warm/hard lateness at creation is recorded above as Owner-directed rotation pause/P6-D closure override, not a newly invented violation.

Không blame outgoing predecessor hoặc Owner. Các findings là historical governance facts, không phải lý do reopen accepted P6-D technical gates.

## Workspace, authority và stop conditions

- Sole repository workspace là `E:\Anvien`, saved project local, không worktree.
- Cấm C-drive workspace/read/write, alternate worktree, implicit network, install/package-script side effects, shared/global/protected/quarantined cache authority.
- Stage/isolated commit sau completed accepted implementation slice là standing campaign authority theo AGENTS Rule 12; không xin Owner cấp lại. Không push nếu Owner chưa nói push.
- Owner `PAUSE` hoặc `STOP` dừng ngay mọi command/edit/QA/commit/lane control.
- Không accept transfer thiếu exact heading `OFFICIAL AUTHORITY TRANSFER` và detached seal khớp report này.
- Incoming Main phải khởi tạo successor kế tiếp tại exact warm target tính từ raw official transfer UUIDv7 mới (+45 phút), và transfer không muộn hơn hard deadline (+60 phút), trừ khi Owner ra lệnh pause/override mới.

## Immediate successor continuation

Sau mandatory startup và detached-seal verification:

1. Xác minh current HEAD/index/26-row protected boundary một lần phục vụ handoff; không biến thành audit loop.
2. Chạy fresh `anvien analyze --force` thật sự đến exit `0`; báo exact result và điều nó chứng minh.
3. Đọc full P6-E causal/architecture/workflow/feasibility/plan-acceptance reports và current four-ledger P6-E authority.
4. Chạy current P6-E file-detail/impact/cost-map/A/A/baseline preflight.
5. Thiết kế và mở visible P6-E implementation/performance lane với exact goal, authority, editable/inspect/preserve boundaries, evidence, rollback, resource, timeout, stop và completion conditions; monitor actual behavior continuously.
6. Không cập nhật lại P6-D checklist, không rerun P6-D PASSed gates nếu chưa chứng minh invalidation, không mở Pn-A, không push.

Outgoing Main terminates absolutely after official transfer dispatch.
