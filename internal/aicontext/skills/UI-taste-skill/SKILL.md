---
name: ui-taste-skill
description: Điều phối bộ UI taste cho frontend premium, redesign, image-to-code, imagegen web/mobile, brandkit, Stitch DESIGN.md, style variants và full-output.
version: 1.0.0
---

# UI Taste Skill Router

Skill này chỉ điều phối bộ `UI-taste-skill`. Không dùng file gốc này như một bộ luật thiết kế đầy đủ; luôn đọc skill con phù hợp trước khi làm.

## Khi dùng

Dùng khi yêu cầu liên quan đến:
- frontend visual taste, landing page, portfolio, marketing site, product page, hero section
- nâng cấp UI hiện có, làm UI bớt generic, bớt AI slop, premium hơn
- image-to-code cho website có yêu cầu thị giác cao
- sinh ảnh concept cho web, mobile app, hoặc brand kit
- tạo `DESIGN.md` cho Google Stitch
- áp một hướng thẩm mỹ cụ thể như minimalist, brutalist, high-end/soft
- ép output đầy đủ, không placeholder, không rút gọn

Không dùng làm skill chính cho backend, database, API, infra, bug fix logic, dashboard nghiệp vụ phức tạp, hoặc product workflow nếu yêu cầu chính không phải visual quality.

## Nguyên tắc điều phối

1. Đọc yêu cầu và xác định deliverable thật: code frontend, redesign dự án có sẵn, ảnh tham chiếu, mobile screen, brand kit, `DESIGN.md`, hay chỉ ép output.
2. Chọn đúng một skill con làm primary route.
3. Chỉ thêm skill phụ khi nó thật sự bổ sung route chính.
4. Xác nhận file route tồn tại trước khi dùng; nếu thiếu, báo thiếu file và không giả lập nội dung skill.
5. Mở và đọc toàn bộ `SKILL.md` của mọi skill con đã chọn trước khi hành động.
6. Nếu skill con trỏ tới file phụ như `DESIGN.md`, asset, hoặc reference bắt buộc, đọc file đó trước khi dùng.
7. Không nạp toàn bộ bộ skill cho mọi tác vụ; dùng progressive disclosure theo route.
8. Khi user gọi rõ tên skill con, ưu tiên đúng skill đó trừ khi nó mâu thuẫn trực tiếp với deliverable user yêu cầu.

## Route Chính

| Trường hợp | Primary route |
|---|---|
| Landing page, portfolio, marketing site, hero, frontend visual implementation mới | `skills/UI-taste-skill/skills/taste-skill/SKILL.md` |
| Cần hành vi v1 cũ, user gọi v1, hoặc workflow đang pin v1 | `skills/UI-taste-skill/skills/taste-skill-v1/SKILL.md` |
| Muốn taste rất gắt kiểu Awwwards/Codex/GPT, motion mạnh, GSAP/anti-slop nghiêm | `skills/UI-taste-skill/skills/gpt-tasteskill/SKILL.md` |
| Website cần image-first: sinh ảnh trước, phân tích ảnh sâu, rồi code bám ảnh | `skills/UI-taste-skill/skills/image-to-code-skill/SKILL.md` |
| Dự án/site/app đã tồn tại cần redesign hoặc polish mà không phá chức năng | `skills/UI-taste-skill/skills/redesign-skill/SKILL.md` |
| Sinh ảnh tham chiếu website/landing/product comp, không code | `skills/UI-taste-skill/skills/imagegen-frontend-web/SKILL.md` |
| Sinh ảnh mobile app screens/flows, không code | `skills/UI-taste-skill/skills/imagegen-frontend-mobile/SKILL.md` |
| Sinh brand guidelines board, logo system, identity deck, visual world | `skills/UI-taste-skill/skills/brandkit/SKILL.md` |
| Tạo design-system `DESIGN.md` cho Google Stitch | `skills/UI-taste-skill/skills/stitch-skill/SKILL.md` |
| Output bắt buộc đầy đủ, không `...`, không TODO, không placeholder | `skills/UI-taste-skill/skills/output-skill/SKILL.md` |

## Route Phong Cách

Các skill dưới đây thường là support route, không thay thế primary route nếu user còn yêu cầu code, redesign, image-to-code, hoặc image generation.

| Hướng thẩm mỹ | Skill |
|---|---|
| High-end agency, mềm, cinematic, whitespace lớn, micro-interaction tinh | `skills/UI-taste-skill/skills/soft-skill/SKILL.md` |
| Minimal editorial, warm monochrome, bento phẳng, ít shadow, không gradient | `skills/UI-taste-skill/skills/minimalist-skill/SKILL.md` |
| Industrial brutalist, Swiss print, tactical terminal, grid cứng, data dày | `skills/UI-taste-skill/skills/brutalist-skill/SKILL.md` |

Chỉ kết hợp nhiều style route khi user yêu cầu hybrid rõ ràng. Nếu không, chọn một hướng và giữ nhất quán.

## Luật Kết Hợp

- Existing project redesign: dùng `redesign-skill` làm primary. Có thể thêm một style route nếu user chỉ rõ hướng thẩm mỹ.
- Visual website code mới: dùng `taste-skill` mặc định. Nếu visual fidelity là trung tâm và có image generation, dùng `image-to-code-skill`.
- Image-only web comp: dùng `imagegen-frontend-web`; không viết code.
- Image-only mobile app: dùng `imagegen-frontend-mobile`; không viết code.
- Brand identity board: dùng `brandkit`; không biến nó thành UI implementation.
- Google Stitch: dùng `stitch-skill`; nếu cần template chuẩn, đọc thêm `skills/UI-taste-skill/skills/stitch-skill/DESIGN.md`.
- Full output: dùng `output-skill` như add-on cho route chính, không dùng nó làm luật thiết kế.
- Dashboard/admin/data table/product workflow: không tự ép `taste-skill`; chỉ dùng style/taste khi user đang yêu cầu visual polish, còn workflow phải theo skill frontend/product phù hợp.

## Luật Không Được Làm

- Không code trước khi đọc skill con đã chọn.
- Không dùng image-only skill rồi tự ý chuyển sang code.
- Không gom web imagegen và mobile imagegen vào cùng một route nếu user không yêu cầu cả hai.
- Không dùng `taste-skill-v1` chỉ vì nó tồn tại; v2 là mặc định.
- Không áp nhiều aesthetic trái nhau làm output mất identity.
- Không biến router này thành bản sao dài của các skill con.

## Cách Báo Cáo Khi Bắt Đầu

Nói ngắn gọn:

```text
Đang dùng UI Taste router: primary = <skill con>, support = <skill phụ nếu có>, lý do = <deliverable user yêu cầu>.
```

Sau đó đọc file skill con tương ứng và làm theo workflow của file đó.

## Kiểm Tra Trước Khi Kết Thúc

- Đã dùng đúng primary route theo deliverable.
- Đã đọc toàn bộ skill con đã chọn.
- Nếu có support route, support route không mâu thuẫn primary route.
- Nếu route là image-only, không sinh code.
- Nếu route là redesign, không rewrite dự án từ đầu.
- Nếu route là image-to-code, ảnh tham chiếu là source thị giác chính trước khi code.
- Nếu route là `output-skill`, không còn placeholder hoặc phần bị rút gọn.
