# Phương án đổi benchmark sang hai target ổn định

## Phương án chốt

Plan Child 06A không đổi. A001/D001 không đổi. Giải pháp code và cách triển khai không đổi.

Thay đổi duy nhất là target dùng để Anvien sinh graph và đo tốc độ:

```text
Trước đây:
Anvien sinh graph cho E:\Anvien

Thay bằng:
Anvien sinh graph cho E:\cheapapp.org
Anvien sinh graph cho E:\Restaurant_manager
```

E:\Anvien vẫn là engine được tối ưu. Cheapapp và Restaurant Manager chỉ là hai repo đầu vào để Anvien sinh hai graph riêng. Không có thay đổi code nào trong hai repo target.

## Lý do thay đổi

Khi Anvien sinh graph cho chính E:\Anvien, code engine và source đầu vào cùng thay đổi trong quá trình tối ưu. Vì vậy số đo trước và sau không còn dựa trên cùng một target ổn định.

Cheapapp và Restaurant Manager hiện không biến động, nên phù hợp làm hai target cố định. Hai repo cũng có cấu trúc ngôn ngữ khác nhau:

- Cheapapp chủ yếu là TypeScript/TSX.
- Restaurant Manager có khối Go lớn cùng TypeScript/TSX.

Đo trên hai cấu trúc source khác nhau giúp kết quả khách quan hơn việc chỉ đo trên một repo.

## Quy trình đo

### Bước 1 — Đo baseline trước tối ưu

Dùng cùng một bản Anvien trước tối ưu để:

1. Sinh graph cho E:\cheapapp.org và ghi lại số liệu.
2. Sinh graph cho E:\Restaurant_manager và ghi lại số liệu.

Với mỗi target, số liệu gồm:

- thời gian D001 resolve_calls;
- thời gian parent resolution;
- thời gian toàn bộ graph generation/analyze;
- graph và các kết quả đầu ra cần giữ nguyên.

### Bước 2 — Niêm phong baseline

Sau khi hai graph baseline và các số liệu baseline đã được tạo xong, niêm phong toàn bộ kết quả đó làm mốc so sánh bất biến.

Phần được niêm phong là:

- graph baseline sinh từ Cheapapp;
- graph baseline sinh từ Restaurant Manager;
- bộ số liệu D001, resolution và end-to-end của từng target;
- các work counts và output-equivalence facts đi kèm bộ số liệu đó.

Sau khi niêm phong, không ai được sửa hoặc diễn giải lại các số liệu baseline. Chúng là mốc gốc để so sánh với kết quả sau tối ưu.

### Bước 3 — Tối ưu Anvien

Triển khai A001/D001 trong E:\Anvien theo đúng phương án code đã chốt.

Hai repo target không tham gia vào phần triển khai. Anvien chỉ đọc source của chúng để sinh graph.

### Bước 4 — Đo lại sau tối ưu

Dùng Anvien sau tối ưu để:

1. Sinh lại graph cho E:\cheapapp.org và lấy bộ số liệu mới.
2. Sinh lại graph cho E:\Restaurant_manager và lấy bộ số liệu mới.

### Bước 5 — So sánh kết quả

Với từng target, so sánh bộ số liệu mới với đúng baseline đã niêm phong để xác định:

1. D001 resolve_calls nhanh hơn bao nhiêu.
2. Parent resolution nhanh hơn bao nhiêu.
3. Toàn bộ quá trình Anvien sinh graph nhanh hơn bao nhiêu.
4. Graph và kết quả đầu ra có giữ nguyên hay không.

Graph Cheapapp sau tối ưu chỉ so với graph Cheapapp baseline. Graph Restaurant Manager sau tối ưu chỉ so với graph Restaurant Manager baseline. Kết quả của hai target được trình bày riêng để không che trường hợp một target nhanh hơn nhưng target còn lại chậm đi.

## Quan hệ với phương án D001 đã chốt

Việc đổi target benchmark không thay đổi giải pháp D001:

```text
{canonical source path, exact local import name}
    → []original w.imports indices in original order
```

Giải pháp vẫn loại bỏ repeated whole-import scans và repeated path normalization trong call resolution, đồng thời giữ nguyên resolution result, import order, first-match, duplicate behavior, graph, output, persistence, readers và lifecycle semantics.

Khi resolve_calls chạy nhanh hơn, parent resolution và toàn bộ pipeline sinh graph phải nhận được lợi ích tương ứng. Mức tăng tốc thực tế được xác định bằng cách so sánh hai lần đo sau tối ưu với hai baseline đã niêm phong.

## Kết luận

```text
Dùng Anvien trước tối ưu sinh graph cho hai target
→ ghi số liệu baseline
→ niêm phong graph và số liệu baseline
→ tối ưu code Anvien theo A001/D001
→ sinh lại graph cho đúng hai target
→ so sánh với baseline đã niêm phong
→ xác định Anvien sinh graph nhanh hơn bao nhiêu
```

Đây chỉ là thay đổi target benchmark để có số đo ổn định và khách quan hơn. Plan, active child D001, phương án code và cách triển khai vẫn giữ nguyên.
