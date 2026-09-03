# Giao thức QA Backend (API & Database)

> Áp dụng khi QA Backend, API, Database, non-UI service. Không mở trình duyệt, không chụp ảnh màn hình.

---

### 1. Khi thực thi kiểm thử
* **Trước khi gọi API:** Query DB lấy snapshot trạng thái ban đầu của thực thể mục tiêu.
* **Khi gọi API:** Gửi request với đủ các biến thể payload (hợp lệ, sai schema, thiếu auth, boundary limits).
* **Sau khi gọi API:** Query lại DB để đối chiếu State Diff thực tế (xác nhận đúng schema, không ghi đè dữ liệu, bảo toàn foreign keys).

### 2. Khi ghi nhận bằng chứng (`evidence.md`)
* Lưu mã HTTP status và Response payload.
* Lưu log câu lệnh SQL/DB query và thời gian thực thi (phát hiện N+1 queries, unindexed sequential scans).
* Lưu log xác nhận transaction kết thúc bằng `COMMIT` hoặc `ROLLBACK`.

### 3. Khi đo hiệu năng (`benchmark.md`)
* Ghi lại độ trễ phản hồi: Latency p50, p95 (ms).
* Ghi lại mức tiêu thụ RAM (Heap/RSS) và số lượng connection active trong Connection Pool.
* Ghi lại tỷ lệ test pass (unit test, integration test, contract test).

### 4. Khi rà soát bẫy lỗi kỹ thuật
* **Kiểm tra Rollback:** Cố tình kích hoạt failure ở bước cuối của transaction $\rightarrow$ Query DB chứng minh không có partial write hoặc dữ liệu mồ côi (orphaned records).
* **Kiểm tra Race Condition:** Bắn concurrent requests vào cùng một resource ID $\rightarrow$ Kiểm tra có xảy ra Lost Update hoặc vi phạm ràng buộc dữ liệu (constraint violation) không.
* **Kiểm tra Rò rỉ tài nguyên:** Bắn dồn dập request lỗi $\rightarrow$ Kiểm tra có bị cạn kiệt connection pool (pool starvation/leak) hoặc RAM không được giải phóng không.

### 5. Khi nào chuyển giao sang `Edge-Case`
* Khi phạm vi chứa các mutation tranh chấp cao (high-contention resources), atomic counters, distributed lock, hoặc multi-table ACID transactions:
  * QA xác nhận baseline functional đúng theo hợp đồng.
  * Ghi chú đề xuất CEO kích hoạt lane `Edge-Case` để thực hiện chaos attacks (hostile timing, out-of-order execution, process crash mid-flight).
