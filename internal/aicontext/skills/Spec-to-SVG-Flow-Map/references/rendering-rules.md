# Quy Tắc Render

## Quy Tắc Render Từng Flow

Không vẽ tất cả flow trong một batch.

Dùng loop này:

1. Chọn đúng một named flow hoặc subflow từ source union inventory và flow manifest.
2. Đọc lại source spec sections và reference artifacts cho flow đó.
3. Extract actors, entry points, preconditions, commands, guards, branches, state/data reads, writes, side effects, terminal states, recovery paths, handoffs, và invariants của flow đó.
4. Chỉ vẽ flow đó.
5. Thêm visible SVG nodes, edges, junctions, terminals, và gaps của flow đó.
6. Thêm metadata tương ứng cho flow đó.
7. Verify flow đó với source inventory trước khi sang flow tiếp theo.
8. Mark mỗi covered source item với mapped node/edge/junction/terminal/gap id.
9. Chỉ sau đó mới tiếp tục flow tiếp theo.

Không dựa vào trí nhớ từ lần đọc trước khi render flow sau. Re-open hoặc re-read relevant source sections trước mỗi flow.

Final SVG có thể chứa tất cả flows, nhưng construction process phải là flow-by-flow.

## Quy Tắc Không Vẽ Bulk

Không bao giờ render nhiều independent pipelines từ một high-level summary pass.

Bulk drawing pass chỉ được phép cho:

- legend
- global lane layout
- cross-flow index
- high-level navigation anchors
- final consistency cleanup

Bulk pass không được dùng cho logic triển khai.

Logic triển khai phải được add qua Flow-By-Flow Rendering Rule.

## Quy Tắc Không Collapse

Không collapse named details thành generic nodes.

Node như `BUSINESS_OPS`, `SYNC_ENGINE`, `AUTH_FLOW`, `REPORTS`, `SETTINGS`, hoặc `LOCAL_COMMAND` chỉ được dùng làm index, grouping label, hoặc junction. Nó không đủ coverage cho detailed source behavior.

Nếu source đặt tên roads như POS order, pay/refund, move/merge/split, shift/cash, inventory, owner setup, report coverage, local print, snapshot bootstrap, manual sync, lifecycle reconnect, hoặc recovery, mỗi road phải được represent riêng.

Mọi generic/index node phải fan out tới detailed nodes hoặc detail sheets.

Fail run nếu generic node thay thế named implementation road.

## Chi Tiết Tối Thiểu Cho Mỗi Flow

Với mỗi named flow hoặc subflow, include ít nhất:

1. Entry trigger.
2. Actor, role, device, và scope preconditions.
3. UI, external, runtime, IPC, hoặc backend command source.
4. Authority hoặc permission gate.
5. State/data reads.
6. Decision branches.
7. Write/apply target.
8. Side effects.
9. Async, outbox, background, hoặc transport behavior.
10. Success terminal state.
11. Failure, denied, blocked, pending, hoặc rollback terminal state.
12. Recovery hoặc retry path.
13. Pipeline handoff junctions.
14. Source-of-truth boundary.
15. Explicit invariants và `must not` rules.

Nếu item bắt buộc không được source define, tạo visible gap node.

## Phát Hiện Gap

Tạo gap node khi:

1. Decision node thiếu complete branches.
2. Reachable state không có next step.
3. Error có thể xảy ra nhưng không có recovery path.
4. Sync conflict không có resolution rule.
5. Backend contract được mention nhưng không define.
6. UI state cần thiết nhưng không define.
7. Payment, license, auth, hoặc permission ảnh hưởng flow nhưng không rõ.
8. Source of truth không rõ.
9. Manual step thiếu actor hoặc owner.
10. Terminal state mơ hồ.
11. Flow cross sang pipeline khác mà không có junction.
12. Background job có thể fail nhưng không có retry/failure rule.
13. Source union inventory item là implementation-relevant nhưng không có mapped SVG id.
14. Reference artifact có named detail vắng mặt trong new map.
15. Flow thiếu required minimum detail item và source không define nó.
16. Relation cần thiết để implementation an toàn nhưng source không mô tả.

Mọi gap node phải visible trong SVG và được liệt kê trong verification report.

## Quy Tắc Pipeline Handoff

Khi một pipeline nói với pipeline khác, tạo junction node nếu source mô tả handoff đó.

Ví dụ source-described:

```text
AUTH_FLOW -> LICENSE_FLOW
```

Represent thành:

```text
AUTH_LOGIN -> AUTH_SUCCESS_JUNCTION
AUTH_SUCCESS_JUNCTION -> LICENSE_CHECK
```

Không represent pipeline handoff chỉ bằng crossing arrows.

Nếu source không mô tả handoff, không được tự tạo junction/edge. Tạo gap nếu handoff thiếu chặn implementation.
