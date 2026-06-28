# Trung Thành Với Source

## Nguyên Tắc Trung Thành Với Source

Vẽ SPEC đúng như source artifacts mô tả, kể cả khi mô tả đó thiếu, sai, rời rạc, mâu thuẫn, không an toàn, hoặc không đủ để code.

Không vẽ ideal lifecycle, expected architecture, implementation order hợp lý, hoặc runtime handoff bị thiếu nếu source artifact không nói rõ.

Nếu source mô tả các flow rời nhau, vẽ các flow rời nhau. Connected lifecycle graph chỉ bắt buộc khi source mô tả các connection đó.

Nếu source nói flow A dẫn đến flow B, vẽ relationship đó với source reference.

Nếu source không nói flow A dẫn đến flow B, không được nối chúng. Nếu relation bị thiếu đó chặn implementation an toàn, tạo visible gap node như `UNDEFINED_HANDOFF`, `SPEC_GAP`, hoặc `OWNER_DECISION_REQUIRED`.

Nếu source mô tả sai thứ tự, unsafe branch, missing denial path, hoặc broken recovery path, giữ nguyên hình dạng sai/broken đó trong SVG và đánh dấu defect là gap hoặc risk. Không được âm thầm sửa đúng.

Mọi implementation-logic node, edge, junction, terminal, và gap phải trace về source reference hoặc explicit missing-source/gap record. Implementation logic không có reference bị tính là speculation và fail run.

Layout groups, legends, lane labels, và navigation anchors có thể derived, nhưng không được thêm ý nghĩa implementation.

## Mục Tiêu Độ Chi Tiết

Tạo flow map chi tiết hơn bất kỳ source diagram hoặc reference artifact nào đã cung cấp cho cùng scope.

Không tạo overview-only map khi source có implementation roads, command families, lifecycle states, storage checkpoints, guards, cursors, locks, receipts, versions, hash chains, terminal states, hoặc recovery paths.

Output phải giữ union của:

- source spec details
- related authority docs
- existing flow maps hoặc SVGs
- verification reports
- UI/state/API/contract maps
- invariants và negative rules được đặt tên rõ

Nếu prior/reference SVG có named flow, branch, state, store, cursor, command, invariant, hoặc terminal state mà SVG mới không represent, run fail.

## Không Lý Tưởng Hóa

Fail run nếu:

1. Edge, junction, hoặc lifecycle handoff được thêm chỉ vì nó expected hoặc architecturally desirable, nhưng không source-described.
2. Missing relation được vẽ thành real edge thay vì visible gap.
3. Agent correct, normalize, hoặc reconnect flawed source flow thay vì render flaw và mark gap/risk.
4. Agent claim spec rõ trong khi còn unresolved gap node.

## Yêu Cầu Source Ref

Mọi node, edge, junction, terminal, và gap có ý nghĩa implementation phải có một trong hai loại bằng chứng:

- `data-source-ref` / `source_ref` trỏ tới source section cụ thể.
- Gap record nói rõ source thiếu gì và vì sao nó chặn implementation.

Không được dùng document title, section title, hoặc broad flow ID làm bằng chứng duy nhất cho detail implementation.
