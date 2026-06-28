---
name: spec-to-svg-flow-map
description: Chuyển spec sản phẩm, tính năng, UI, backend, auth, sync, lifecycle, external contract, hoặc spec đa nhánh thành SVG flow map ngữ nghĩa có metadata máy đọc được, source coverage, render từng flow, gap detection, verification report, và trạng thái BLOCKED hoặc READY_FOR_OWNER_REVIEW. Dùng khi cần biến spec thành flow map, audit độ đầy đủ của spec trước khi code, so sánh với diagram tham chiếu, phơi bày hành vi chưa định nghĩa, hoặc xác định quyết định Owner trước implementation.
---

# Spec Thành SVG Flow Map

## Mục Đích

Chuyển spec thành bộ artifact flow map SVG có thể đọc được bởi người và tool/code.

Skill này không tạo sơ đồ trang trí. Skill này tạo bản đồ ngữ nghĩa để phơi bày:

- luồng người dùng và UI
- luồng dữ liệu, state, storage, source-of-truth
- luồng auth, session, thiết bị, quyền, ranh giới bảo mật
- luồng sync, mạng, lifecycle, reconnect, recovery
- contract backend, API, hệ thống ngoài
- lỗi, recovery, background job, runtime command
- decision, branch, handoff pipeline, terminal state, out-of-scope reference
- gap spec, undefined behavior, owner decision, missing data, bottleneck, risk
- legend giải thích ở cuối SVG

Mục tiêu là vẽ đúng SPEC đang mô tả, không vẽ hệ thống mà agent nghĩ là đúng.

## Đọc Tài Liệu Tham Chiếu Bắt Buộc

Trước khi vẽ bất kỳ logic triển khai SVG nào, đọc các file reference theo thứ tự:

1. `references/source-fidelity.md`
2. `references/source-inventory-and-manifest.md`
3. `references/rendering-rules.md`
4. `references/svg-contract.md`
5. `references/verification-and-acceptance.md`

Đọc thêm `references/domain-detail-checklist.md` khi scope có auth, session, scope, runtime command, business flow, sync, snapshot, report, print, lifecycle, settings, backend/API, hoặc contract ngoài.

Nếu reference bắt buộc không đọc được, dừng lại và báo `BLOCKED`.

## Quy Tắc Cứng

1. Không implement app/source code.
2. Không sửa production source file.
3. Không đoán hành vi thiếu.
4. Không che giấu ambiguity.
5. Không đặt logic quan trọng chỉ trong geometry hoặc vị trí visual.
6. Mọi quan hệ flow quan trọng phải có text visible và XML metadata.
7. Mọi decision node phải có đủ outgoing branches.
8. Mọi edge phải có condition rõ ràng, dùng `condition="always"` cho unconditional flow.
9. Mọi pipeline handoff phải đi qua junction node thật.
10. Mọi terminal state phải có node riêng.
11. Mọi ambiguity chưa giải quyết phải thành gap node visible.
12. Phải build source-union inventory trước khi vẽ logic triển khai.
13. Phải build source-described flow manifest trước khi vẽ logic triển khai.
14. Phải render logic triển khai từng flow, không render từ một bulk summary pass.
15. Không collapse detail có tên thành generic node.
16. Không improve, repair, normalize, hoặc tự nối flow vượt quá điều source thật sự nói.
17. Nếu không thể hoàn thành map một cách an toàn, dừng lại và nói rõ lý do.

Nếu map có node unresolved với type `SPEC_GAP`, `UNDEFINED_BEHAVIOR`, hoặc `OWNER_DECISION_REQUIRED`, đặt implementation status là `BLOCKED`.

Nếu `missing_source_items` không rỗng, đặt implementation status là `BLOCKED`.

## Đầu Vào Bắt Buộc

Đọc authority có sẵn cho spec slice được yêu cầu:

1. Source spec.
2. Related authority docs.
3. Existing/reference SVGs và flow maps, nếu có.
4. Existing verification reports, nếu có.
5. UI prototype hoặc UI slot map, nếu có.
6. State/source map, nếu có.
7. Backend/API/contract map, nếu có.
8. Actual wiring status, nếu có.
9. Project rules.

Nếu input bắt buộc bị thiếu, không bịa hành vi. Đánh dấu nó là gap trong SVG và verification report.

## Đầu Ra Bắt Buộc

Tạo hoặc cập nhật các file:

```text
docs/flow-maps/<feature-name>.flow.svg
docs/flow-maps/<feature-name>.flow-map.md
docs/flow-maps/<feature-name>.flow-verification.md
```

Dùng lowercase kebab-case cho `<feature-name>` trừ khi project có quy tắc tên chặt hơn.

## Quy Trình Bắt Buộc

1. Đọc project rules và requested scope.
2. Đọc source spec và related authority docs.
3. Đọc existing/reference SVGs, flow maps, verification reports, UI maps, state maps, API maps, và contract maps khi có.
4. Build source union inventory theo `references/source-inventory-and-manifest.md`.
5. Build source-described flow manifest từ inventory item được source mô tả trực tiếp.
6. Tạo global lane/layout plan. Chưa vẽ logic triển khai.
7. Với từng flow trong manifest, lặp:

```text
đọc lại source sections liên quan
extract flow-local inventory
xác nhận flow đã có trong source-described flow manifest
vẽ chỉ flow đó
ghi visible nodes/edges/junctions/terminals/gaps
ghi metadata tương ứng
verify source coverage cho flow đó
mark mapped source inventory items
mark manifest item bằng mapped SVG ids hoặc gap ids
```

8. Sau khi vẽ tất cả flow, chạy global consistency verification.
9. Check missing source items.
10. Check flow manifest omissions và unmanifested visible flows.
11. Check collapse violations.
12. Check reference SVG delta.
13. Parse SVG như XML.
14. Verify metadata/visible node round-trip.
15. Thêm visible legend ở cuối SVG, là visible group cuối cùng trước `</svg>`, và id phải khớp `metadata.legend.visible_legend_id`.
16. Đặt implementation status.

## Trạng Thái Kết Thúc

Đặt:

```text
BLOCKED
```

khi còn unresolved gaps, undefined behavior, owner decisions, missing source items, collapse violations, hoặc failed reference deltas.

Đặt:

```text
READY_FOR_OWNER_REVIEW
```

chỉ khi không còn blocker về `SPEC_GAP`, `UNDEFINED_BEHAVIOR`, `OWNER_DECISION_REQUIRED`, source coverage, collapse, hoặc reference delta.

Chỉ Owner mới được approve implementation sau khi review generated flow map.

## Định Dạng Bàn Giao Cuối

Trả về đúng format:

```md
# Bàn Giao Spec Flow Map

## File Đã Tạo / Cập Nhật
- docs/flow-maps/<feature-name>.flow.svg
- docs/flow-maps/<feature-name>.flow-map.md
- docs/flow-maps/<feature-name>.flow-verification.md

## Trạng Thái
BLOCKED or READY_FOR_OWNER_REVIEW

## Phát Hiện Chính
1.
2.
3.

## Hành Động Owner Cần Làm
- If BLOCKED: liệt kê chính xác spec sections, source inventory items, reference deltas, hoặc collapse violations cần làm rõ.
- If READY_FOR_OWNER_REVIEW: yêu cầu Owner review và approve SVG flow map trước implementation.
```
