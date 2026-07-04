---
name: spec-to-svg-flow-map
description: Chuyển đổi spec sản phẩm, tính năng, UI, backend, auth, sync, lifecycle, contract ngoài, hoặc spec đa nhánh thành SVG flow map có metadata machine-readable, kiểm tra source-union, render từng flow, phát hiện gap, báo cáo xác minh, và trạng thái BLOCKED/READY_FOR_OWNER_REVIEW.
---

# Spec → SVG Flow Map (Tiếng Việt)

## Mục Đích

Chuyển spec thành file `.svg` có thể đọc được bởi cả người và tool/code.

Không phải sơ đồ trang trí — đây là **bản đồ ngữ nghĩa** phơi bày:

- Flow người dùng / UI
- Flow dữ liệu / state / storage / source-of-truth
- Flow auth / session / thiết bị / quyền / ranh giới bảo mật
- Flow sync / mạng / lifecycle / reconnect / recovery
- Flow backend / API / contract hệ thống ngoài
- Flow lỗi / recovery / background job / lệnh runtime
- Quyết định, nhánh, handoff pipeline, terminal state, và tham chiếu ngoài phạm vi
- Gap spec, hành vi chưa định nghĩa, quyết định chủ sở hữu, dữ liệu thiếu, bottleneck, rủi ro
- Legend

SVG phải giúp người và coding agent quyết định spec có đủ rõ để triển khai an toàn không.

SVG phải render **đúng như nguồn mô tả**, không phải implementation mà agent cho là đúng.

---

## Mục Tiêu Độ Chi Tiết

Tạo flow map **chi tiết hơn** bất kỳ sơ đồ tham chiếu nào đã cung cấp cho cùng phạm vi.

Đầu ra phải bảo tồn hợp nhất của:
- Chi tiết spec nguồn
- Tài liệu có thẩm quyền liên quan
- Flow map / SVG hiện có
- Báo cáo xác minh
- Bản đồ UI / state / API / contract
- Bất biến và quy tắc phủ định đã đặt tên

Nếu SVG tham chiếu có flow, nhánh, state, store, cursor, lệnh, bất biến, hoặc terminal state mà SVG mới không có → **run thất bại**.

---

## Phạm Vi & Không-Mục-Tiêu

- Chỉ dùng skill này để tạo hoặc cập nhật flow map từ spec.
- Không viết code app trong lúc dùng skill này.
- Không biến sự mơ hồ chưa giải quyết thành hành vi đoán mò.
- Không tạo sơ đồ trang trí không có metadata ngữ nghĩa.
- Không bắt đầu triển khai từ đầu ra này cho đến khi Owner xem xét và phê duyệt.

---

## Đầu Vào Bắt Buộc

Đọc toàn bộ nguồn có thẩm quyền cho phần spec được yêu cầu:

1. Spec nguồn
2. Tài liệu có thẩm quyền liên quan
3. SVG / flow map tham chiếu hiện có (nếu có)
4. Báo cáo xác minh hiện có (nếu có)
5. Prototype UI hoặc slot map UI (nếu có)
6. Bản đồ state / source (nếu có)
7. Bản đồ backend / API / contract (nếu có)
8. Trạng thái kết nối thực tế (nếu có)
9. Quy tắc dự án

Nếu đầu vào bắt buộc thiếu → không bịa hành vi. Đánh dấu là gap trong SVG và báo cáo xác minh.

---

## Đầu Ra Bắt Buộc

Tạo hoặc cập nhật:

```
docs/flow-maps/<feature-name>.flow.svg
docs/flow-maps/<feature-name>.flow-map.md
docs/flow-maps/<feature-name>.flow-verification.md
```

Dùng kebab-case viết thường cho `<feature-name>` trừ khi dự án có quy tắc đặt tên chặt hơn.

---

## Quy Tắc Cứng

1. Không viết code app / source.
2. Không sửa file production.
3. Không suy ra hành vi còn thiếu.
4. Không che giấu sự mơ hồ.
5. Không để logic quan trọng chỉ tồn tại trong hình học hoặc vị trí visual.
6. Mọi quan hệ flow quan trọng phải xuất hiện trong **text hiển thị và metadata XML**.
7. Mọi decision node phải có đủ nhánh đi ra.
8. Mọi edge phải có điều kiện rõ ràng; dùng `condition="always"` cho flow vô điều kiện.
9. Mọi handoff pipeline phải dùng junction node thực sự.
10. Mọi terminal state phải có node tường minh.
11. Mọi sự mơ hồ chưa giải quyết → đánh dấu là gap node.
12. Xây dựng source-union coverage trước khi vẽ logic implementation.
13. Render logic implementation từng flow, không phải một lần tổng quan.
14. Không gộp chi tiết đã đặt tên vào node generic.
15. Không cải thiện, sửa, chuẩn hóa, hoặc kết nối flow vượt quá những gì nguồn thực sự mô tả.
16. Nếu map không thể hoàn thành an toàn → dừng và báo cáo lý do.

**Nếu map có node chưa giải quyết kiểu `SPEC_GAP`, `UNDEFINED_BEHAVIOR`, hoặc `OWNER_DECISION_REQUIRED` → trạng thái triển khai là `BLOCKED`.**

**Nếu `missing_source_items` không rỗng → `BLOCKED`.**

---

## Quy Tắc Trung Thực Theo Mô Tả

Vẽ spec **đúng như các nguồn mô tả**, kể cả khi mô tả không đầy đủ, sai, ngắt kết nối, mâu thuẫn, hoặc không an toàn.

- Không vẽ lifecycle lý tưởng, kiến trúc kỳ vọng, thứ tự implementation có thể xảy ra, hoặc handoff runtime còn thiếu trừ khi nguồn nói tường minh.
- Nguồn mô tả flow ngắt kết nối → vẽ flow ngắt kết nối.
- Nguồn nói flow A dẫn đến flow B → vẽ quan hệ đó với tham chiếu nguồn.
- Nguồn không nói A dẫn đến B → không kết nối. Nếu quan hệ còn thiếu đó chặn triển khai an toàn → tạo gap node hiển thị.
- Nguồn mô tả thứ tự sai, nhánh không an toàn, đường denial thiếu, hoặc đường recovery hỏng → bảo tồn hình dạng sai/hỏng đó và đánh dấu là gap/rủi ro. Không tự ý sửa.

Mọi node logic implementation, edge, junction, terminal, và gap phải truy được về tham chiếu nguồn hoặc bản ghi gap tường minh. Logic implementation không có tham chiếu → coi là suy đoán và **run thất bại**.

---

## Cổng Manifest Flow Theo Mô Tả Nguồn (Trước Khi Render)

Sau khi xây source union inventory và **trước khi vẽ** bất kỳ logic SVG nào, tạo **manifest flow được mô tả bởi nguồn**.

Manifest không phải kế hoạch thiết kế — đó là danh sách flow, subflow, và road được nguồn mô tả tường minh.

Với mỗi mục manifest, ghi lại:

```
flow_id
source_name
category: flow | subflow | road
source_refs
described_entry
described_exit_or_terminal
described_handoffs
missing_handoffs_or_unknown_relations
mapping_status: PENDING | MAPPED | MAPPED_AS_GAP | INTENTIONALLY_OUT_OF_SCOPE | MISSING
planned_or_actual_svg_group_id
```

**Quy tắc:**
1. Mọi flow, subflow, road được nguồn mô tả phải xuất hiện trong manifest trước khi vẽ.
2. Không thêm flow kỳ vọng không có trong nguồn — nếu cần nhưng vắng mặt → thêm mục gap.
3. Không vẽ flow không có trong manifest.
4. Nếu phát hiện flow mới trong khi vẽ → dừng vẽ, cập nhật inventory và manifest, rồi tiếp tục.
5. Mục manifest hoàn chỉnh khi và chỉ khi ánh xạ được vào SVG id hiển thị, gap hiển thị, hoặc bản ghi ngoài phạm vi tường minh.
6. Flow được nguồn mô tả vắng mặt trong manifest → **run thất bại**.
7. Flow implementation hiển thị không có mục manifest → **run thất bại**.

---

## Cổng Source Union Inventory

Trước khi vẽ, xây **source union inventory** từ tất cả artifact đầu vào.

Không đủ khi chỉ dùng tiêu đề tài liệu hay flow ID. Mỗi lệnh, guard, store, cursor, terminal state, đường recovery, bất biến, quy tắc phủ định, và quan hệ đã mô tả liên quan đến implementation phải được inventoried riêng lẻ.

**Các loại cần inventory:**

```
flows, subflows, roads, actors, roles, scopes, screens
UI commands, runtime commands, backend/API commands, IPC commands
permissions, guards, branch conditions, deny conditions
stores, caches, files, DBs, table families, queues, outboxes
projections, snapshots, cursors, checkpoints, receipts, versions
locks, hash chains, coverage markers
lifecycle states, retry rules, rollback rules, terminal states
recovery paths, invariants, negative rules, out-of-scope boundaries
```

**Trạng thái ánh xạ cho mỗi mục:**

```
MAPPED_AS_NODE | MAPPED_AS_EDGE | MAPPED_AS_JUNCTION | MAPPED_AS_TERMINAL
MAPPED_AS_GAP | INTENTIONALLY_OUT_OF_SCOPE | MISSING
```

Mục liên quan đến implementation có trạng thái `MISSING` → **run thất bại**.

---

## Quy Tắc Render Từng Flow

Không vẽ tất cả flow trong một lần.

**Vòng lặp:**

1. Chọn đúng một flow hoặc subflow đã đặt tên từ source union inventory.
2. Đọc lại các phần spec nguồn và artifact tham chiếu cho flow đó.
3. Trích xuất: actor, entry point, tiền điều kiện, lệnh, guard, nhánh, đọc/ghi state/data, side effect, terminal state, đường recovery, handoff, bất biến.
4. Vẽ chỉ flow đó.
5. Thêm node / edge / junction / terminal / gap SVG hiển thị.
6. Thêm metadata tương ứng.
7. Xác minh flow đó với source inventory trước khi chuyển sang flow tiếp.
8. Đánh dấu mục source inventory đã bao phủ với id node/edge/junction/terminal/gap được ánh xạ.
9. Chỉ sau đó mới tiếp tục flow tiếp theo.

Không dựa vào ký ức từ lần đọc trước khi render flow sau. Đọc lại phần nguồn liên quan trước mỗi flow.

---

## Quy Tắc Không Vẽ Hàng Loạt

Không bao giờ render nhiều pipeline độc lập từ một lần tóm tắt tổng quan.

Vẽ hàng loạt chỉ được phép cho: legend, layout lane toàn cục, cross-flow index, navigation anchor cấp cao, và cleanup nhất quán cuối cùng.

**Không được dùng cho logic implementation** — phải qua Quy Tắc Render Từng Flow.

---

## Quy Tắc Không Gộp

Không gộp chi tiết đã đặt tên vào node generic.

Node như `BUSINESS_OPS`, `SYNC_ENGINE`, `AUTH_FLOW` chỉ được dùng làm index, nhãn nhóm, hoặc junction — không đủ bao phủ hành vi nguồn chi tiết.

Nếu nguồn đặt tên các road (POS order, pay/refund, move/merge/split, shift/cash, inventory, owner setup, report coverage, local print, snapshot bootstrap, manual sync, lifecycle reconnect, recovery) → mỗi road phải được đại diện riêng.

Mọi node generic/index phải fan-out đến node chi tiết hoặc detail sheet.

Node generic thay thế một road implementation đã đặt tên → **run thất bại**.

---

## Chi Tiết Tối Thiểu Mỗi Flow

Với mỗi flow / subflow đã đặt tên, bao gồm ít nhất:

1. Trigger entry
2. Actor, role, thiết bị, tiền điều kiện phạm vi
3. Nguồn lệnh (UI, external, runtime, IPC, backend)
4. Cổng quyền / permission
5. Đọc state/data
6. Các nhánh quyết định
7. Target ghi/áp dụng
8. Side effect
9. Hành vi async, outbox, background, hoặc transport
10. Terminal state thành công
11. Terminal state lỗi / bị từ chối / bị chặn / pending / rollback
12. Đường recovery hoặc retry
13. Junction handoff pipeline
14. Ranh giới source-of-truth
15. Bất biến tường minh và quy tắc "không được"

Nếu bất kỳ mục nào không được nguồn định nghĩa → tạo gap node hiển thị.

---

## Checklist Domain Chi Tiết Bắt Buộc

Khi liên quan, mô hình hóa tường minh các domain sau thay vì nhóm chúng:

- **Auth/session/thiết bị:** restore, login handoff, cache, keyring, expiry, logout, denied states
- **Restaurant/phạm vi:** chọn scope, scopes visible/bound, DB mount, trạng thái hydrated/chưa hydrated
- **Permission/command gate:** ranh giới renderer, thẩm quyền Go/backend, command guard
- **Owner/app setup:** setup outbox, setup version, setup receipt, convergence state
- **Business runtime:** POS order, pay/refund, move/merge/split, shift/cash, inventory/stocktake
- **Sync transport:** LAN/WSS/VPS, relay, delta, dedupe, ack, cursor, hash verify, gap, repair
- **Snapshot/bootstrap/manual sync:** baselines, manifest, anchors, áp dụng allowlist, rollback, catchup
- **Reports/coverage:** tách aggregate/detail, retention, source coverage, gating export/print
- **Local print/export:** cấu hình printer, preview, kết quả spooler, ranh giới local-only
- **Lifecycle/reconnect:** active/idle/sleep/offline/resume, relay bị bỏ lỡ, auth refresh, cursor catchup
- **Local settings/hành vi chỉ thiết bị**
- **External/backend contracts:** denied/error responses

---

## Hợp Đồng SVG Ngữ Nghĩa

SVG phải là XML hợp lệ và đọc được như source code.

**Mỗi node:**

```xml
<g
  id="node-..."
  data-type="..."
  data-flow="..."
  data-status="..."
  data-lane="..."
  data-source-ref="..."
>
  <title>...</title>
  <desc>...</desc>
  <text>...</text>
</g>
```

**Mỗi edge:**

```xml
<g
  id="edge-..."
  data-type="edge"
  data-edge-type="..."
  data-from="..."
  data-to="..."
  data-condition="..."
  data-flow="..."
  data-source-ref="..."
>
  <title>...</title>
  <desc>...</desc>
  <path ... />
  <text>...</text>
</g>
```

Không bao giờ tạo path ẩn danh cho logic quan trọng.

---

## Metadata SVG & Bao Phủ Nguồn

Mỗi SVG phải có một block metadata machine-readable:

```xml
<metadata id="spec-flow-map">
{
  "feature": "<feature-name>",
  "source_spec": "<source-spec-path>",
  "version": "draft",
  "reference_artifacts": [...],
  "nodes": [...],
  "edges": [...],
  "flow_manifest": [...],
  "source_inventory": [...],
  "coverage_summary": {
    "total_source_items": 0,
    "mapped_items": 0,
    "missing_items": 0,
    "collapse_violations": 0
  },
  "missing_source_items": [],
  "collapse_violations": [],
  "gaps": [],
  "terminal_states": [],
  "junctions": [],
  "legend": {
    "visible_legend_id": "legend-end",
    "shape_meaning": {},
    "color_meaning": {},
    "arrow_meaning": {}
  }
}
</metadata>
```

Metadata phải khớp với node và edge SVG hiển thị. `missing_source_items` không rỗng → `BLOCKED`.

---

## Lanes, Nodes, Edges

**Lanes được đề xuất:**

```
USER, UI_APP, LOCAL_STATE, LOCAL_DB, SYNC_ENGINE
BACKEND_API, AUTH_SECURITY, PAYMENT_LICENSE
EMAIL_NOTIFICATION, ERROR_RECOVERY, EXTERNAL_SYSTEM, OWNER_DECISION
```

Mỗi node thuộc đúng một lane.

**Loại Node:**

```
START, END, EVENT, SCREEN, USER_ACTION, SYSTEM_ACTION
BACKGROUND_JOB, DECISION, STATE, DATA_STORE, DOCUMENT
EXTERNAL_SYSTEM, JUNCTION, ERROR, RECOVERY
SPEC_GAP, UNDEFINED_BEHAVIOR, OWNER_DECISION_REQUIRED, OUT_OF_SCOPE
```

**Loại Edge:**

```
CONTROL_FLOW, DATA_FLOW, ASYNC_EVENT, SYNC_FLOW
ERROR_FLOW, RECOVERY_FLOW, HANDOFF, BLOCKED_BY, OUT_OF_SCOPE_REFERENCE
```

Mỗi edge phải có: `from`, `to`, `condition`, `data-edge-type`, `flow`, `data-source-ref`.

---

## Hợp Đồng Visual

**Hình dạng:**

| Loại | Hình dạng |
|---|---|
| START / END | Hình tròn |
| SCREEN / USER_ACTION | Hình chữ nhật bo góc |
| SYSTEM_ACTION | Hình chữ nhật |
| BACKGROUND_JOB | Hình chữ nhật nét đứt |
| DECISION | Hình thoi |
| STATE | Viên nang / capsule |
| DATA_STORE | Hình trụ |
| DOCUMENT | Icon tài liệu |
| EXTERNAL_SYSTEM | Hình chữ nhật nét đứt |
| JUNCTION | Hình tròn nhỏ |
| ERROR | Hình chữ nhật bo góc viền đỏ |
| RECOVERY | Hình chữ nhật bo góc viền cam |
| SPEC_GAP | Hình chữ nhật viền đỏ đậm |
| UNDEFINED_BEHAVIOR | Hình thoi viền đỏ đậm |
| OWNER_DECISION_REQUIRED | Hình chữ nhật viền cam đậm |
| OUT_OF_SCOPE | Hình chữ nhật xám nét đứt |

**Màu sắc:**

| Màu | Ý nghĩa |
|---|---|
| Xanh lá | Flow bình thường / tự động / tạo giá trị |
| Xanh dương | Dữ liệu / storage / source of truth |
| Tím | Auth / bảo mật / ranh giới quyền |
| Vàng | Delay / pending / chờ |
| Đỏ | Bottleneck / lỗi / nguy hiểm / vi phạm spec |
| Cam | Quyết định chủ sở hữu / quy tắc nghiệp vụ chưa rõ |
| Xám | Thủ công / external / ngoài phạm vi |
| Đen | Control flow mặc định |
| Nét đứt | Async / background / phụ thuộc gián tiếp |
| Nét liền | Flow trực tiếp |
| Nét đôi | Sync / trao đổi dữ liệu hai chiều |

Không thay đổi ý nghĩa màu giữa các sơ đồ.

---

## Yêu Cầu Legend Cuối (Hiển Thị)

Mọi SVG phải kết thúc bằng một phần legend hiển thị, đọc được bởi người.

Đặt legend là group SVG hiển thị cuối cùng trước `</svg>`:

```xml
<g
  id="legend-end"
  data-type="legend"
  data-position="end"
>
  <title>Legend</title>
  <desc>Giải thích hình dạng node, màu sắc, kiểu đường, và loại mũi tên.</desc>
  ...
</g>
```

Legend phải đọc được mà không cần mở báo cáo markdown. Phải giải thích:

1. Ý nghĩa mỗi hình dạng / hộp node
2. Ý nghĩa mỗi màu sắc
3. Ý nghĩa mỗi kiểu đường
4. Ý nghĩa mỗi loại mũi tên
5. Ý nghĩa các marker: error, recovery, gap, owner-decision, out-of-scope, terminal, junction, async, sync, data-flow, control-flow

Legend bắt buộc kể cả khi SVG lớn hoặc chia thành nhiều vùng visual. Không đặt legend chỉ trong metadata.

---

## Phát Hiện Gap

Tạo gap node khi:

1. Decision node thiếu nhánh đầy đủ
2. State có thể đạt được nhưng không có bước tiếp theo
3. Lỗi có thể xảy ra nhưng không có đường recovery
4. Xung đột sync không có quy tắc giải quyết
5. Backend contract được đề cập nhưng không được định nghĩa
6. UI state cần thiết nhưng không được định nghĩa
7. Payment, license, auth, hoặc quyền ảnh hưởng đến flow nhưng chưa rõ
8. Source of truth chưa rõ
9. Bước thủ công thiếu actor hoặc chủ sở hữu
10. Terminal state mơ hồ
11. Flow chuyển sang pipeline khác nhưng không có junction
12. Background job có thể thất bại nhưng không có quy tắc retry/failure
13. Mục source union inventory liên quan đến implementation nhưng thiếu SVG id đã ánh xạ
14. Artifact tham chiếu có chi tiết đặt tên vắng mặt trong map mới
15. Flow thiếu mục chi tiết tối thiểu bắt buộc và nguồn không định nghĩa nó

Mọi gap node phải hiển thị trong SVG và được liệt kê trong báo cáo xác minh.

---

## Quy Tắc Handoff Pipeline

Khi một pipeline nói chuyện với pipeline khác → tạo junction node.

Ví dụ: `AUTH_FLOW → LICENSE_FLOW` phải được biểu diễn là:
```
AUTH_LOGIN → AUTH_SUCCESS_JUNCTION → LICENSE_CHECK
```

Không biểu diễn handoff pipeline chỉ bằng mũi tên giao nhau.

---

## Quy Trình Cập Nhật

1. Đọc quy tắc dự án và phạm vi được yêu cầu
2. Đọc spec nguồn và tài liệu có thẩm quyền liên quan
3. Đọc SVG / flow map / báo cáo xác minh / bản đồ UI / bản đồ state / bản đồ API / bản đồ contract tham chiếu (nếu có)
4. Xây source union inventory
5. Xây source-described flow manifest từ inventory
6. Xác định flow và subflow đã đặt tên từ manifest
7. Tạo kế hoạch lane/layout toàn cục; chưa vẽ logic implementation
8. Với mỗi flow đã đặt tên, lặp lại:
   ```
   đọc lại phần spec nguồn liên quan
   trích xuất inventory cục bộ của flow
   xác nhận flow đã có trong manifest
   vẽ chỉ flow đó
   viết node/edge/junction/terminal/gap SVG hiển thị
   viết metadata tương ứng
   xác minh bao phủ nguồn cho flow đó
   đánh dấu mục source inventory đã bao phủ
   đánh dấu mục manifest với SVG id đã ánh xạ hoặc gap id
   ```
9. Sau khi vẽ xong tất cả flow, chạy xác minh nhất quán toàn cục
10. Kiểm tra mục nguồn còn thiếu
11. Kiểm tra flow manifest bị bỏ sót hoặc flow hiển thị chưa có manifest
12. Kiểm tra vi phạm collapse
13. Kiểm tra delta SVG tham chiếu
14. Parse SVG như XML
15. Xác minh round-trip metadata / node hiển thị
16. Đặt trạng thái implementation

**Tạo 3 file đầu ra** (mô tả ở phần Đầu Ra Bắt Buộc).

**Đặt trạng thái implementation cuối:**
- `BLOCKED`: khi còn gap chưa giải quyết, hành vi chưa định nghĩa, quyết định chủ sở hữu, mục nguồn còn thiếu, vi phạm collapse, hoặc delta tham chiếu thất bại
- `READY_FOR_OWNER_REVIEW`: chỉ khi không còn blocker nào

Chỉ Owner mới có thể phê duyệt triển khai sau khi xem xét flow map được tạo ra.

---

## Xác Minh Round-Trip

Xác minh:

- Mọi metadata node tồn tại như SVG content hiển thị và ngược lại
- Mọi edge có đủ: `from`, `to`, `condition`, `data-edge-type`, `flow`, `data-source-ref`
- Mọi node / edge / junction / terminal / gap logic implementation có tham chiếu nguồn hoặc bản ghi gap tường minh
- Mọi flow được nguồn mô tả xuất hiện trong flow manifest trước khi xuất hiện trong SVG content hiển thị
- Mọi flow implementation hiển thị có mục manifest tương ứng
- Quan hệ còn thiếu được biểu diễn là gap, không phải edge bịa
- Mọi decision node có nhánh tường minh
- Mọi junction kết nối ít nhất hai đoạn pipeline
- Mọi gap node xuất hiện trong báo cáo xác minh
- Mọi terminal state là tường minh
- Mọi mục source inventory có trạng thái ánh xạ cuối
- Legend cuối hiển thị tồn tại là group SVG hiển thị cuối cùng, đọc được, và giải thích hình dạng, màu sắc, kiểu đường, loại mũi tên
- `metadata.legend.visible_legend_id` khớp với id group legend cuối hiển thị
- Không có logic quan trọng chỉ tồn tại trong hình học
- SVG parse được như XML khi có parser

---

## Định Dạng Báo Cáo Xác Minh Flow

```md
# Báo Cáo Xác Minh Flow: <feature-name>

## Nguồn
- Spec:
- Tài liệu liên quan:
- Artifact tham chiếu:
- SVG được tạo:

## Tóm Tắt
- Tổng số flow: | Mục flow manifest: | Tổng số node: | Tổng số edge:
- Mục source inventory: | Đã ánh xạ: | Còn thiếu: | Vi phạm collapse:
- Decision node: | Junction node: | Terminal state: | Gap spec:
- Quyết định chủ sở hữu cần: | Mục ngoài phạm vi:

## Trạng Thái Implementation
BLOCKED hoặc READY_FOR_OWNER_REVIEW

## Bao Phủ Source Union Inventory
| Mục Nguồn | Loại | Tham Chiếu Nguồn | Trạng Thái Ánh Xạ | SVG IDs Đã Ánh Xạ |

## Manifest Flow Được Mô Tả Bởi Nguồn
| Flow ID | Tên Nguồn | Loại | Tham Chiếu | Entry | Exit/Terminal | Handoff | Thiếu | Trạng Thái | Group/Gap IDs |

## Bao Phủ Từng Flow
| Flow | Phần Nguồn Đã Đọc Lại | Node | Edge | Terminal | Gap | Trạng Thái |

## Mục Nguồn Còn Thiếu
| Mục Nguồn | Loại | Tham Chiếu | Tại Sao Quan Trọng | Cần Sửa |

## Vi Phạm Collapse
| Node Generic | Chi Tiết Thiếu | Tham Chiếu | Cần Tách |

## Delta SVG Tham Chiếu
| Artifact Tham Chiếu | Mục Trong Tham Chiếu | Có Trong SVG Mới | ID Đã Ánh Xạ / Gap |

## Kiểm Tra Fan-Out Node Generic
| Node Generic | Mục Đích Cho Phép | Node Fan-Out | Trạng Thái |

## Gap Nghiêm Trọng
| Gap ID | Loại | Vị Trí | Tại Sao Chặn Coding | Cần Sửa Spec |

## Bao Phủ Quyết Định
| Decision Node | Nhánh Tìm Thấy | Nhánh Còn Thiếu | Trạng Thái |

## Handoff Pipeline
| Junction | Từ Flow | Đến Flow | Điều Kiện | Trạng Thái |

## Terminal State
| Terminal State | Đến Từ | Điều Kiện | Trạng Thái |

## Ghi Chú Rủi Ro
| Rủi Ro | Node/Edge Liên Quan | Mức Độ | Khuyến Nghị |

## Kết Luận Cuối
Nêu rõ BLOCKED hoặc READY_FOR_OWNER_REVIEW.
```

---

## Tiêu Chí Chấp Nhận

Run thành công chỉ khi **tất cả** điều kiện sau đúng:

1. SVG là XML hợp lệ
2. SVG có metadata machine-readable cho node và edge
3. Metadata SVG có trường bao phủ nguồn
4. Markdown flow-map khớp với SVG
5. Báo cáo xác minh hoàn chỉnh
6. Mọi decision có nhánh tường minh
7. Mọi handoff pipeline dùng junction node
8. Mọi terminal state là tường minh
9. Mọi sự mơ hồ chưa giải quyết được đánh dấu là gap
10. Mọi mục source inventory liên quan đến implementation được ánh xạ hoặc biểu diễn là gap
11. Mọi flow được render qua Quy Tắc Render Từng Flow
12. Mọi flow được nguồn mô tả xuất hiện trong manifest trước logic SVG hiển thị
13. Mọi flow implementation hiển thị có mục manifest tương ứng
14. Mọi node / edge / junction / terminal / gap logic implementation có tham chiếu nguồn hoặc bản ghi gap
15. SVG bảo tồn hình dạng flow sai, thiếu, ngắt kết nối, hoặc mâu thuẫn theo nguồn mô tả thay vì tự sửa
16. Mọi node generic/index có mục đích hợp lệ và fan-out đến chi tiết
17. SVG có legend cuối hiển thị giải thích hình dạng, màu sắc, kiểu đường, loại mũi tên, và marker đặc biệt
18. Legend cuối hiển thị là group SVG hiển thị cuối cùng và được tham chiếu từ metadata
19. Không có code production nào bị sửa
20. Trạng thái cuối là `BLOCKED` hoặc `READY_FOR_OWNER_REVIEW`

**Thêm tiêu chí thất bại:**

- SVG tham chiếu có flow hoặc chi tiết đặt tên vắng mặt trong SVG mới
- Node generic thay thế road implementation đã đặt tên
- Mục source inventory có trạng thái `MISSING`
- Nhánh decision, denied state, đường recovery, cursor/checkpoint, hoặc ranh giới source-of-truth chỉ có trong prose không có trong SVG metadata
- SVG mới kém chi tiết hơn sơ đồ hiện có cho cùng phạm vi
- Nhiều pipeline implementation độc lập được vẽ hàng loạt mà không có re-read và xác minh từng flow
- Flow được render mà không ghi lại phần nguồn nào đã đọc lại
- Flow được đánh dấu hoàn thành trước khi mục source inventory được ánh xạ
- Agent che giấu gap spec hoặc tuyên bố spec rõ ràng trong khi còn gap node chưa giải quyết
- SVG thiếu legend cuối hiển thị, hoặc legend không giải thích ý nghĩa hộp node, màu sắc, kiểu đường, loại mũi tên
- Flow được mô tả bởi nguồn được vẽ trước khi xuất hiện trong manifest
- Flow implementation hiển thị tồn tại mà không có mục manifest
- Edge, junction, hoặc handoff lifecycle được thêm vì kỳ vọng hoặc kiến trúc mong muốn nhưng không có trong nguồn
- Quan hệ còn thiếu được vẽ như edge thực thay vì gap hiển thị
- Agent sửa, chuẩn hóa, hoặc kết nối lại flow nguồn bị lỗi thay vì render lỗi đó và đánh dấu là gap/rủi ro

---

## Định Dạng Handoff Cuối

```md
# Handoff Flow Map Spec

## File Đã Tạo / Cập Nhật
- docs/flow-maps/<feature-name>.flow.svg
- docs/flow-maps/<feature-name>.flow-map.md
- docs/flow-maps/<feature-name>.flow-verification.md

## Trạng Thái
BLOCKED hoặc READY_FOR_OWNER_REVIEW

## Phát Hiện Hàng Đầu
1.
2.
3.

## Hành Động Chủ Sở Hữu Cần
- Nếu BLOCKED: liệt kê chính xác các phần spec, mục source inventory, delta tham chiếu, hoặc vi phạm collapse cần làm rõ.
- Nếu READY_FOR_OWNER_REVIEW: yêu cầu Owner xem xét và phê duyệt SVG flow map trước khi triển khai.
```
