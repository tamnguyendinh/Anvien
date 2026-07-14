---
name: ui-animation-design
description: Điều phối bộ skill UI animation/design gồm tra thuật ngữ animation, tư vấn Apple-style fluid interaction, tư vấn Emil-style design engineering, review animation code/diff, và audit/cải thiện animation toàn codebase. Dùng khi user hỏi về UI animation, motion design, gesture, transition, animation review, polish UI, hoặc cải thiện cảm giác chuyển động của app.
---

# UI Animation Design

## Mục Đích

Skill này là router điều phối cho bộ skill UI animation/design.

Không dùng file này để thay thế các skill con. Dùng nó để chọn đúng skill con, đọc `SKILL.md` của skill con đó, rồi làm theo workflow của skill con.

## Skill Con

| Nhu cầu của user | Skill cần dùng | Ghi chú |
|---|---|---|
| Hỏi hiệu ứng này gọi là gì, hoặc mô tả motion nhưng không biết tên | `animation-vocabulary` | Chỉ trả thuật ngữ/khái niệm, không thiết kế hoặc implement |
| Tư vấn gesture, spring, velocity, momentum, sheet, drag, swipe, material, Apple-like fluid UI | `apple-design` | Dùng cho motion vật lý, trực tiếp, interruptible |
| Tư vấn polish component, easing, duration, component feel, design engineering craft | `emil-design-eng` | Dùng cho taste/craft, decision framework, component-level motion |
| Audit animation toàn app, “improve animations”, “make app feel better”, roadmap animation fixes | `improve-animations` | Read-only trên source code; tạo audit/plan, không sửa code |
| Review diff/file/code animation cụ thể | `review-animations` | Dùng standards nghiêm ngặt, cite `file:line`, trả findings + verdict |

## Quy Tắc Đọc

1. Không đọc tất cả skill con mặc định.
2. Chọn skill con theo intent của user.
3. Đọc toàn bộ `SKILL.md` của skill con đã chọn trước khi làm.
4. Nếu skill con trỏ tới file phụ như `AUDIT.md`, `PLAN-TEMPLATE.md`, hoặc `STANDARDS.md`, đọc file đó khi workflow yêu cầu.
5. Khi nhiều skill con cùng liên quan, chọn một skill chính và chỉ dùng skill phụ làm context hỗ trợ.
6. Nếu intent của user chưa rõ, hỏi lại ngắn gọn trước khi chọn skill con.

## Routing

### Naming / Vocabulary

Dùng `animation-vocabulary` khi user hỏi:

- “Hiệu ứng này gọi là gì?”
- “Cái kiểu popover mọc ra từ nút gọi là gì?”
- “Motion này tên gì để prompt AI/designer?”

Output phải ngắn, ưu tiên thuật ngữ chính xác.

### Review Code Hoặc Diff

Dùng `review-animations` khi user đưa:

- diff
- file code
- PR/change cần review
- câu như “review animation này”, “xem motion code này đúng chưa”

Nếu cần giá trị chính xác về easing, duration, performance, accessibility, đọc thêm `review-animations/STANDARDS.md`.

### Audit / Improve Toàn Codebase

Dùng `improve-animations` khi user hỏi:

- “improve animations”
- “audit motion”
- “make this app feel better”
- “tạo roadmap animation fixes”

Skill này read-only trên source code. Không sửa code, không format, không build có side effect, không commit.

Khi viết plan, đọc thêm:

- `improve-animations/AUDIT.md`
- `improve-animations/PLAN-TEMPLATE.md`

### Apple-Style Fluid Interaction

Dùng `apple-design` khi vấn đề xoay quanh:

- gesture
- direct manipulation
- spring
- velocity handoff
- momentum projection
- rubber-banding
- draggable sheet/drawer
- swipe interactions
- material, translucency, depth
- reduced motion theo kiểu Apple-style

### Emil-Style Design Engineering

Dùng `emil-design-eng` khi vấn đề xoay quanh:

- component polish
- motion có nên tồn tại không
- easing/duration cho UI thường ngày
- press feedback
- popover/dropdown origin
- animation feel
- performance và craft detail

## Conflict Rules

1. `review-animations` thắng khi user yêu cầu review code/diff cụ thể.
2. `improve-animations` thắng khi user yêu cầu audit hoặc cải thiện animation toàn app.
3. `animation-vocabulary` thắng khi user chỉ hỏi tên hiệu ứng.
4. `apple-design` thắng khi trọng tâm là gesture, spring, momentum, velocity, sheet, drag, swipe, hoặc material.
5. `emil-design-eng` thắng khi trọng tâm là polish, taste, component feel, easing, duration, hoặc perceived quality.
6. Nếu user vừa cần naming vừa cần thiết kế, dùng `animation-vocabulary` trước để đặt tên, rồi chuyển sang skill thiết kế phù hợp.
7. Nếu user vừa cần audit toàn app vừa hỏi một diff cụ thể, dùng `review-animations` cho diff cụ thể và `improve-animations` cho roadmap toàn app.

## Không Được Trộn Mode

Không dùng `improve-animations` để review một diff nhỏ.

Không dùng `review-animations` để audit toàn codebase.

Không dùng `animation-vocabulary` để đề xuất implementation.

Không dùng `apple-design` hoặc `emil-design-eng` để bỏ qua standards khi đang review code.

Không dùng root skill này để tự tạo tiêu chuẩn mới thay cho skill con.

## Scope Và Safety

1. Không sửa code nếu skill con là read-only.
2. Không tạo plan triển khai trừ khi skill con yêu cầu.
3. Không dùng animation chỉ vì “đẹp”; motion phải có purpose.
4. Luôn tôn trọng `prefers-reduced-motion`.
5. Với review, cite `file:line`.
6. Với audit, phân biệt finding đã verify với cảm giác cần feel-check.
7. Với gesture/touch interaction, yêu cầu kiểm tra trên runtime hoặc thiết bị thật khi kết luận phụ thuộc cảm giác.

## Output Contract

Theo skill con đã chọn:

- `animation-vocabulary`: trả thuật ngữ ngắn gọn, có 1-2 alternate nếu cần.
- `review-animations`: trả markdown table `Before | After | Why`, rồi verdict.
- `improve-animations`: trả audit table, missed opportunities, rồi chờ user chọn plan.
- `apple-design`: trả khuyến nghị interaction cụ thể, nêu rõ spring/velocity/material/reduced-motion khi liên quan.
- `emil-design-eng`: trả khuyến nghị design engineering cụ thể, có easing/duration/implementation guidance khi cần.

Nếu không chắc nên dùng skill con nào, trả lời bằng một câu hỏi làm rõ thay vì đoán.
