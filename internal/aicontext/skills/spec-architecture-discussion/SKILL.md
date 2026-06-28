---
name: spec-architecture-discussion
description: Thảo luận kiến trúc dựa trên SPEC bằng cách đọc tuần tự và đọc hết từng file SPEC trước khi kết luận; trả lại bản nắm bắt kiến trúc gồm invariant, boundary, pipeline, luồng dữ liệu, liên kết thật giữa pipeline, điểm chưa rõ, và câu hỏi thảo luận. Dùng khi user yêu cầu thảo luận kiến trúc, thảo luận về mức độ kiến trúc, bóc tách kiến trúc từ SPEC, hoặc tranh luận một vấn đề kiến trúc dựa trên bộ SPEC. Không dùng để tự cập nhật SPEC, viết code, hoặc review acceptance.
---

# Thảo Luận Kiến Trúc Từ SPEC

## Mục Đích

Skill này dùng để thảo luận kiến trúc dựa trên bộ SPEC đặc biệt của dự án.

Nhiệm vụ không phải là áp pattern kiến trúc có sẵn vào app. Nhiệm vụ là đọc hiểu SPEC, nắm invariant, boundary, pipeline, luồng dữ liệu, điểm chưa rõ, rồi tạo nền thảo luận chính xác.

Bộ SPEC này không được xem như tài liệu thông thường và không được xử lý bằng keyword search. Phải đọc hiểu toàn bộ file SPEC liên quan trước khi thảo luận.

## Nguyên Tắc Cứng

1. Phải đọc tuần tự từng file SPEC liên quan.
2. Mỗi file SPEC đã chọn phải được đọc hết toàn bộ trước khi dùng nó làm căn cứ.
3. Cấm đọc theo từ khóa.
4. Cấm làm việc theo từ khóa.
5. Cấm dùng keyword giống nhau để gom cụm pipeline nếu chúng không thật sự liên quan.
6. Cấm mang pattern học được từ hệ thống khác để gán vào app này.
7. Cấm tự ý cập nhật tài liệu SPEC.
8. Chỉ được cập nhật SPEC khi user ra đúng lệnh: `cập nhật spec`.
9. Những từ tương đương như "ghi lại", "sửa lại", "bổ sung", "làm rõ", "update docs", "chỉnh tài liệu" không được xem là lệnh cập nhật SPEC.
10. Nếu chưa đọc đủ SPEC liên quan, không được kết luận kiến trúc.
11. Nếu không xác định được bộ SPEC cần đọc, hỏi lại user thay vì đoán.

## Xác Định Bộ SPEC Cần Đọc

Khi user đưa file SPEC cụ thể, đọc hết từng file đó.

Khi user đưa thư mục SPEC, liệt kê các file SPEC trong thư mục theo thứ tự ổn định rồi đọc hết từng file.

Khi user nêu vấn đề nhưng không chỉ rõ file, xác định bộ SPEC liên quan bằng cấu trúc tài liệu, chỉ mục, tên file, hoặc ngữ cảnh đã được user cung cấp. Không được dùng keyword search trong nội dung để thay thế việc đọc.

Nếu chỉ có thể tìm được một slice nhỏ nhưng vấn đề rõ ràng thuộc tổng thể lớn hơn, nói rõ rằng phạm vi SPEC chưa đủ và hỏi user cần đọc thêm file nào.

## Ledger Đọc SPEC

Mỗi lần thảo luận phải có ledger thể hiện đã đọc SPEC nào.

Ledger tối thiểu:

```md
## SPEC Đã Đọc
| Thứ tự | File SPEC | Đã đọc hết toàn bộ | Vai trò trong vấn đề |
|---:|---|---|---|
```

Nếu một file chưa được đọc hết vì lỗi tool, thiếu quyền, context quá lớn, hoặc file không tồn tại, không được dùng file đó làm căn cứ kết luận. Ghi rõ `chưa đủ điều kiện kết luận`.

## Nguyên Tắc Thảo Luận

Khi user đưa ra một vấn đề, đọc lại tất cả SPEC liên quan để hiểu sâu và hiểu cốt lõi của vấn đề cần thảo luận.

Khi user lật lại một slice nhỏ trong một vấn đề, chỉ xem slice đó là một góc nhìn trong tổng thể của vấn đề. Không được chuyển vấn đề gốc sang slice nhỏ đó.

Không xem lời phản biện, giả thuyết, hoặc cách user lật ngược vấn đề là chân lý. Ý kiến của user là thêm góc nhìn để tìm hiểu nguyên nhân, không phải authority thay thế SPEC.

Nếu SPEC và nhận định của user khác nhau, nêu rõ:

- SPEC đang nói gì
- user đang đặt nghi vấn gì
- điểm nào cần kiểm chứng thêm
- kết luận nào chưa được phép rút ra

## Phân Loại Pipeline

Phải phân loại vấn đề theo pipeline rõ ràng.

Không gom pipeline vì:

- có từ khóa giống nhau
- cùng domain noun
- cùng xuất hiện trong một file
- cùng nằm trên một UI screen
- cùng liên quan đến một actor

Liên kết thật giữa hai pipeline chỉ tồn tại khi:

1. một command trong pipeline này cần state từ pipeline khác để quyết định;
2. một invariant của pipeline này phụ thuộc state hoặc terminal state của pipeline khác;
3. source SPEC mô tả trực tiếp handoff, dependency, hoặc read/write boundary giữa chúng.

Nếu chỉ có khả năng liên quan nhưng SPEC chưa mô tả, ghi là điểm chưa rõ hoặc câu hỏi thảo luận. Không tự nối pipeline.

## Cấm Gán Pattern Ngoài

Không tự gán các pattern như SaaS, POS, DDD, CQRS, event sourcing, sync engine, offline-first, state machine, event bus, command bus, clean architecture, hoặc microservice nếu SPEC không mô tả.

Chỉ được dùng pattern ngoài trong hai trường hợp:

1. SPEC tự dùng khái niệm đó.
2. User yêu cầu so sánh với pattern đó.

Ngay cả khi được so sánh, phải nói rõ phần nào là SPEC, phần nào là so sánh ngoài SPEC.

## Quy Trình Bắt Buộc

1. Xác định vấn đề kiến trúc user muốn thảo luận.
2. Xác định bộ SPEC cần đọc.
3. Đọc tuần tự từng file SPEC và đọc hết toàn bộ mỗi file.
4. Lập ledger SPEC đã đọc.
5. Tách invariant, boundary, pipeline, command, state, luồng dữ liệu, terminal state, và dependency thật.
6. Phân biệt rõ điều SPEC nói, điều user đang nghi vấn, và điều chưa đủ căn cứ.
7. Trả lại bản nắm bắt kiến trúc để thảo luận.
8. Không cập nhật SPEC trừ khi user ra đúng lệnh `cập nhật spec`.

## Đầu Ra Bắt Buộc

Trả lời theo cấu trúc này, có thể rút gọn phần không liên quan nhưng không được bỏ ledger:

```md
# Nắm Bắt Kiến Trúc

## SPEC Đã Đọc
| Thứ tự | File SPEC | Đã đọc hết toàn bộ | Vai trò trong vấn đề |
|---:|---|---|---|

## Vấn Đề Đang Thảo Luận

## Cốt Lõi Kiến Trúc Theo SPEC

## Invariant

## Boundary

## Pipeline
| Pipeline | Căn cứ SPEC | Command / Trigger | State cần đọc | Quyết định / Terminal | Ghi chú |
|---|---|---|---|---|---|

## Luồng Dữ Liệu

## Liên Kết Thật Giữa Pipeline
| Từ pipeline | Sang pipeline | Điều kiện liên kết thật | Căn cứ SPEC |
|---|---|---|---|

## Điểm Chưa Rõ

## Câu Hỏi Thảo Luận

## Những Điều Không Được Kết Luận

## Trạng Thái SPEC
Không cập nhật SPEC trong lượt này.
```

## Khi Không Đủ Điều Kiện

Nếu chưa đọc đủ SPEC liên quan, trả lời ngắn gọn:

```md
# Chưa Đủ Điều Kiện Thảo Luận Kiến Trúc

## Lý Do
- Chưa đọc đủ SPEC liên quan: ...

## Cần Bổ Sung
- File/spec/scope cần đọc tiếp: ...

## Điều Không Được Kết Luận
- ...
```

Không được lấp khoảng trống bằng phỏng đoán.
