# Báo cáo thảo luận: Anvien — Môi trường làm việc đa provider và Coordinator giao tiếp

- Ngày: 2026-09-11
- Repo: `E:\Anvien`
- Loại: Báo cáo tổng hợp thảo luận
- Trạng thái: Đề xuất để Owner xem xét; chưa phải đặc tả được phê duyệt hoặc xác nhận triển khai.
- Nguồn: Nguyên văn phản hồi tổng hợp trong phiên `01a0911d-65f9-73f3-9156-070f6da14fd0`.
- Phạm vi biên tập: Chỉ bổ sung tiêu đề và metadata; giữ nguyên nội dung phản hồi.

---

# 1. Anvien sẽ trở thành gì?

**Anvien là dev tool cung cấp một không gian làm việc chung cho nhiều agent thuộc các provider khác nhau.**

Owner sử dụng một web UI để:

- Chọn provider/model và tạo các agent.
- Gán vai trò cho từng agent.
- Giao mục tiêu.
- Theo dõi các agent trao đổi và thực hiện công việc.
- Tham gia hội thoại, trả lời hoặc thay đổi chỉ dẫn.
- Xem kết quả, báo cáo và thay đổi code.

Đây **không phải xây lại VS Code**, cũng không phải chỉ thêm dashboard xem log.

**Hội thoại chung là nơi công việc diễn ra.** Các agent gửi nội dung vào đó, nhận tin, được kích hoạt để xử lý và tiếp tục trao đổi.

Phần dưới là bản thiết kế tổng thể để bạn xem lại; chưa khẳng định các module này đã tồn tại trong source Anvien.

# 2. Vai trò và agent là hai lớp khác nhau

**CEO, coder, reviewer, QA… là tên vai trò.** Bên dưới mỗi vai trò là một agent cụ thể, được cấu hình bằng provider/runtime và model.

Ví dụ:

| Vai trò | Agent | Provider/runtime | Model |
|---|---|---|---|
| CEO | Agent A | Antigravity | Gemini do Owner chọn |
| Coder | Agent B | Codex | GPT do Owner chọn |
| Reviewer | Agent C | Codex hoặc provider khác | Model do Owner chọn |
| QA | Agent D | Provider được chọn | Model được chọn |

Không hard-code:

- CEO phải là Antigravity.
- Coder phải là Codex.
- Một provider chỉ được giữ một vai trò.
- Một model chỉ được có một agent.

Hai agent có thể dùng cùng model nhưng có **phiên, vai trò, chỉ dẫn và lịch sử riêng**.

**Owner là người sử dụng**, có quyền chỉ định và quản lý các agent.

# 3. Ba trách nhiệm cốt lõi

| Thành phần | Làm gì? | Không làm gì? |
|---|---|---|
| **Owner** | Chọn agent, phân vai, giao mục tiêu, cấp quyền | Không phải tự chuyển tin giữa các môi trường |
| **Agent AI** | Đọc, suy luận, phản hồi, làm việc hoặc phân công theo vai trò | Không phải tự xây đường liên lạc |
| **Coordinator** | Nhận, đóng gói, lưu và chuyển tin đúng đích | Không suy luận, chia việc hay quyết định nội dung |

**Coordinator là tool/module phần mềm, không phải AI.**

AI giữ vai trò CEO quyết định giao việc cho ai. Coordinator chỉ chuyển yêu cầu đó đến đúng agent.

# 4. Không gian làm việc chung

Đề xuất tổ chức thành các cấp:

| Cấp | Ý nghĩa |
|---|---|
| **Workspace** | Không gian làm việc gắn với repo hoặc phạm vi phát triển |
| **Agent** | Một thành viên AI được cấu hình provider/model |
| **Conversation** | Cuộc hội thoại chung hoặc trao đổi có nhóm người nhận |
| **Lane** | Nhánh công việc có mục tiêu và agent phụ trách |
| **Message** | Nội dung trao đổi |
| **Artifact** | Báo cáo, log, diff, commit hoặc đầu ra liên quan |

Lane là cách tổ chức công việc, **không phải một provider hay model mới**.

Một agent có thể tiếp tục lane cũ. Nếu được cấp quyền, CEO có thể yêu cầu tạo lane mới và chỉ định agent hiện có hoặc tạo agent mới.

Các quyết định này do CEO hoặc Owner đưa ra, không do Coordinator suy đoán.

# 5. Luồng hoạt động chính

Ví dụ Owner chọn:

- CEO: Gemini qua Antigravity.
- Coder: GPT qua Codex.
- Reviewer: một agent khác.

Luồng làm việc:

1. **Owner gửi mục tiêu vào hội thoại chung.**
2. Coordinator lưu tin, hiển thị trên UI và chuyển đến CEO.
3. Hệ thống kích hoạt lượt xử lý của agent CEO.
4. CEO đọc và quyết định chia công việc.
5. CEO gửi chỉ dẫn cho coder qua Coordinator.
6. Tin xuất hiện trên UI, được chuyển vào đúng phiên coder.
7. Coder xử lý, dùng công cụ để làm việc và gửi tiến độ/kết quả.
8. Tin của coder được chuyển đến CEO, kích hoạt CEO xử lý.
9. CEO quyết định yêu cầu sửa, giao reviewer, mở lane khác hoặc hỏi Owner.
10. Các phản hồi tiếp tục xuất hiện trong không gian chung.

```text
Owner gửi mục tiêu
        ↓
Coordinator → UI chung → Agent CEO
                            ↓
                      CEO quyết định
                            ↓
Coordinator → UI chung → Agent Coder
                            ↓
                    Coder làm và trả lời
                            ↓
Coordinator → UI chung → Agent CEO
                            ↓
                Giao Reviewer / trả lời Owner
```

**UI không cần luôn mở để các agent liên lạc.** Backend thực hiện giao nhận và kích hoạt; UI hiển thị cùng dữ liệu đó.

# 6. Coordinator hoạt động thế nào?

## Nhận tin

Người gửi hoặc agent chỉ định:

- Gửi trong cuộc hội thoại/lane nào.
- Gửi đến agent nào.
- Nội dung.
- Đang trả lời tin nào, nếu có.

Nếu gửi theo vai trò như “CEO”, hệ thống tra cấu hình vai trò để lấy agent đích. Đây là định tuyến theo dữ liệu, không phải AI hiểu nội dung rồi chọn người nhận.

## Đóng gói

Coordinator tạo envelope gồm metadata:

| Metadata | Mục đích |
|---|---|
| Message ID | Nhận diện tin và chống xử lý trùng |
| Workspace/conversation/lane | Xác định nơi tin thuộc về |
| Sender | Ai gửi |
| Recipient | Agent hoặc người nhận |
| Reply-to | Liên kết phản hồi |
| Timestamp | Thời điểm |
| Payload type | Loại nội dung |
| Delivery metadata | Theo dõi vận chuyển |

Nội dung chính nằm trong payload.

## Chuyển tin

Coordinator:

1. Kiểm tra quyền gửi.
2. Lưu tin.
3. Phát sự kiện để UI cập nhật.
4. Định tuyến đến đầu nối của agent nhận.
5. Theo dõi trạng thái vận chuyển.

**Lưu tin và ghi nhận việc cần chuyển phải nhất quán**, tránh tình huống UI đã hiện tin nhưng hệ thống quên chuyển.

## Không diễn giải nội dung

Coordinator không:

- Tóm tắt hay viết lại tin.
- Tự chọn model.
- Tự chia task.
- Tự tạo lane dựa trên câu chữ.
- Đánh giá kết quả.
- Tự phản hồi thay agent.

# 7. Cơ chế kích hoạt agent khi có tin mới

Đây là phần bắt buộc của hệ thống.

**Tin đến đúng agent phải dẫn tới một lượt xử lý phù hợp**, không chỉ nằm trong hộp thư chờ AI tình cờ đọc.

### Agent đang idle

Hệ thống đưa tin vào phiên và yêu cầu runtime bắt đầu lượt xử lý.

### Agent đang làm việc

Xử lý theo khả năng runtime:

- Đưa chỉ dẫn vào lượt hiện tại nếu hỗ trợ.
- Hoặc giữ tin chờ để xử lý ở lượt tiếp theo.

UI phải thể hiện cách đã áp dụng. Không tự ngắt công việc chỉ để ép tin vào.

### Agent mất kết nối

Tin được giữ lại và hiển thị trạng thái chờ/chưa chuyển được.

Không tự tạo một agent thay thế rồi giao lại toàn bộ công việc.

### Nhiều tin đến cùng lúc

Hệ thống cần quản lý thứ tự và tránh mở nhiều lượt xử lý xung đột trên cùng phiên.

Cơ chế này là xử lý hàng đợi và vòng đời runtime, **không phải quyết định nghiệp vụ**.

Việc khảo sát giao thức từng provider phục vụ xây đầu nối này; không thay thế mục tiêu xây cơ chế liên lạc.

# 8. Giao nhận và phản hồi là hai chuyện khác nhau

| Mốc | Ý nghĩa |
|---|---|
| Đã lưu | Anvien đã lưu tin |
| Đang chờ chuyển | Tin nằm trong hàng đợi |
| Đã gửi sang runtime | Adapter đã chuyển yêu cầu |
| Runtime đã tiếp nhận | Có xác nhận kỹ thuật nếu được hỗ trợ |
| Agent bắt đầu xử lý | Có sự kiện tương ứng từ runtime |
| Agent phản hồi | Có tin trả lời thực tế |

Không coi “đã gửi” là “đã hiểu” hoặc “đã hoàn thành”.

Phản hồi như “đã nhận việc” phải do agent tạo. Coordinator không viết thay.

# 9. Provider adapter và quản lý phiên

Mỗi provider/runtime có adapter riêng.

Adapter phụ trách:

- Tạo/kết nối phiên.
- Chuyển tin đến phiên.
- Kích hoạt lượt xử lý.
- Nhận tin phản hồi và sự kiện.
- Thực hiện thao tác điều khiển được hỗ trợ.
- Báo lỗi hoặc khả năng chưa hỗ trợ.

Session management phụ trách:

- Ánh xạ agent trong Anvien với session runtime.
- Theo dõi vòng đời.
- Bảo đảm một đầu mối điều khiển mỗi session.
- Giữ liên kết khi reconnect.
- Phân biệt phiên đang hoạt động và phiên đã kết thúc.

**Provider không đồng nghĩa model.** Owner cần chọn runtime/provider trước hoặc cùng với model, vì runtime quyết định công cụ và cơ chế thực thi.

Đối với phiên có sẵn ngoài Anvien, chỉ kết nối qua cơ chế được hỗ trợ; không cưỡng ép chiếm phiên hoặc sửa lock.

# 10. Context của mỗi provider được giữ riêng

Anvien không ép mọi provider về một giới hạn context/token.

- Context window theo model.
- Compact theo runtime.
- Reasoning setting theo provider.
- Token budget chỉ dùng khi được cấu hình rõ và provider hỗ trợ.
- Usage chỉ hiển thị dữ liệu có thật.
- Coordinator không tự cắt hoặc tóm tắt payload.

**Lịch sử hội thoại chung trong Anvien không đồng nghĩa toàn bộ lịch sử luôn được nhét vào context của mọi agent.**

Agent nhận nội dung được gửi cho mình. Nếu cần đọc thêm lịch sử, có thể dùng công cụ đọc hội thoại theo quyền. Cách runtime duy trì context vẫn thuộc provider.

Nếu nội dung không thể chuyển nguyên vẹn, phải báo rõ thay vì âm thầm làm mất dữ liệu.

# 11. Công cụ dành cho agent

Cần tách công cụ giao tiếp khỏi công cụ quản lý đội.

## Công cụ giao tiếp

- Gửi tin.
- Trả lời tin.
- Đọc hội thoại.
- Xem trạng thái giao nhận.
- Gửi tham chiếu artifact.

Đây là nhóm Coordinator.

## Công cụ quản lý, nếu được cấp quyền

- Tạo lane.
- Gán agent vào lane.
- Tạo agent từ cấu hình được phép.
- Xem các agent hiện có.
- Yêu cầu đóng hoặc thay đổi lane.

CEO suy luận rồi gọi công cụ tương ứng. Backend kiểm tra quyền và thực hiện.

**Coordinator không tự biến câu “hãy tạo reviewer” thành một agent mới.** Agent CEO phải gọi thao tác có cấu trúc hoặc Owner thực hiện trên UI.

# 12. Web UI — nơi vận hành chính

## A. Workspace overview

Hiển thị:

- Các agent và vai trò.
- Provider/model của từng agent.
- Các lane.
- Hội thoại gần đây.
- Nội dung đang chờ Owner.
- Trạng thái kết nối và hoạt động.

## B. Khu vực tổ chức agent

Owner có thể:

- Tạo agent.
- Chọn provider/model.
- Đặt tên và vai trò.
- Cấu hình chỉ dẫn.
- Cấp quyền.
- Gán workspace và lane.

Vai trò có thể đổi nhưng việc chuyển phiên/model phải được xử lý rõ, không hứa giữ nguyên context nếu runtime không hỗ trợ.

## C. Hội thoại chung

Mỗi tin hiển thị:

- Người/agent gửi.
- Vai trò và agent thực tế.
- Đích nhận.
- Nội dung.
- Lane liên quan.
- Trạng thái giao nhận.
- Phản hồi và artifact.

Owner có thể gửi cho CEO, một agent khác hoặc nhóm đích được chọn.

**Hiển thị chung không có nghĩa gửi cho tất cả agent.** Phải phân biệt người được xem và agent được kích hoạt để trả lời.

## D. Lane detail

Hiển thị:

- Mục tiêu lane.
- Agent tham gia.
- Hội thoại.
- Hoạt động được runtime cung cấp.
- Báo cáo và kết quả.

## E. Hoạt động kỹ thuật

Xem công cụ, lệnh, trạng thái thực thi và lỗi nếu provider cung cấp.

Không yêu cầu mọi runtime xuất cùng mức chi tiết. Không hiển thị suy đoán như sự thật.

## F. Artifact và thay đổi code

Xem báo cáo, diff, log và commit liên quan.

Đây là phần hỗ trợ vận hành công việc, không phải xây một editor mới.

# 13. Trạng thái quan sát

Cần hiển thị riêng:

| Trục | Ví dụ |
|---|---|
| Kết nối | Kết nối, mất kết nối, chưa xác định |
| Lượt xử lý | Idle, đang xử lý, kết thúc, lỗi |
| Hoạt động | Đang chạy công cụ, đã có kết quả — nếu có dữ liệu |
| Giao nhận | Đã lưu, chờ chuyển, đã tiếp nhận |
| Nghiệp vụ | Nội dung/status do agent hoặc Owner công bố |

Ví dụ agent gửi “BLOCKED”:

- Coordinator chuyển nguyên nội dung.
- UI hiển thị tin.
- CEO được kích hoạt để đọc và quyết định.
- Coordinator không tự phân tích chữ BLOCKED để tạo kế hoạch khắc phục.

Nếu muốn có trạng thái nghiệp vụ có cấu trúc, agent gửi qua thao tác tương ứng; hệ thống không cần đoán từ văn bản.

# 14. Security & Observability

## Security

- Xác thực Owner và adapter.
- Bảo vệ thông tin xác thực provider.
- Kiểm soát ai được gửi/đọc trong workspace.
- Kiểm soát quyền tạo agent, mở lane và điều khiển phiên.
- Ghi audit.
- Phân biệt rõ chỉ dẫn Owner với nội dung từ AI.
- Không nâng quyền chỉ vì một agent yêu cầu.

## Observability

- Theo dõi sự kiện thật.
- Hiển thị thời điểm cập nhật cuối.
- Báo lỗi vận chuyển.
- Phân biệt mất kết nối với kết thúc công việc.
- Ghi nhận thông tin chưa xác định.

**Không có giới hạn context chung. Không tự kill/restart vì im lặng. Không tự kết luận agent treo chỉ từ CPU hoặc thời gian.**

Chống tin trùng và vòng kích hoạt lặp là bảo vệ cơ chế giao tiếp, không phải hạn chế context của model.

# 15. Telegram

Telegram là cửa truy cập bổ sung cho Owner:

- Gửi tin đến agent.
- Nhận phản hồi.
- Theo dõi hội thoại cần chú ý.
- Trả lời câu hỏi.
- Xem kết quả.

Tin Telegram đi qua cùng Coordinator và xuất hiện trên web.

Không cần Telegram để các agent trong Anvien liên lạc với nhau. Không để Telegram trở thành nguồn dữ liệu chính hoặc một hệ thống trạng thái song song.

# 16. Anvien code intelligence

Anvien tiếp tục cung cấp công cụ code intelligence cho các agent:

- Tìm context.
- Xem quan hệ code.
- Đánh giá impact.
- Phân tích thay đổi.
- Lấy bằng chứng phục vụ triển khai và review.

**Giao tiếp và phân tích code là hai phân hệ riêng.**

Coordinator không cần graph để chuyển tin. Các agent dùng graph khi công việc cần.

Bằng chứng phải gắn đúng repo/revision. Các agent ở worktree khác nhau không được mặc nhiên coi là đang nhìn cùng trạng thái code.

# 17. Các module cần triển khai

| Module | Trách nhiệm |
|---|---|
| **Workspace & Agent Registry** | Workspace, agent, vai trò, provider/model và địa chỉ phiên |
| **Conversation & Lane Store** | Hội thoại, thành viên, lane và liên kết công việc |
| **Communication Core** | Tin, envelope, hàng đợi, định tuyến và giao nhận |
| **Coordinator Tools/API** | Công cụ gửi/đọc/trả lời và xem trạng thái |
| **Session & Activation Runtime** | Vòng đời phiên, nhận tin và kích hoạt lượt xử lý |
| **Provider Adapters** | Kết nối Codex, Antigravity và các runtime khác |
| **Agent/Lane Management Tools** | Thao tác tạo/gán/quản lý theo quyền |
| **Web Workspace UI** | Giao diện làm việc chung |
| **Security & Observability** | Quyền, audit, thông tin xác thực và quan sát |
| **Code Intelligence Integration** | Công cụ Anvien cho các agent |
| **Telegram Gateway** | Kênh Owner từ xa |

Các module này không bắt buộc là các service độc lập.

# 18. Thứ tự triển khai

### Giai đoạn 1 — Chốt cấu trúc và trải nghiệm

- Agent, role, session, conversation, lane.
- Cách Owner chọn provider/model.
- Địa chỉ gửi–nhận.
- Trạng thái giao nhận.
- Luồng UI.
- Ranh giới quyền.

Khảo sát source Anvien và runtime để xác định cách tích hợp cụ thể.

### Giai đoạn 2 — Xây lõi hội thoại và Coordinator

- Lưu hội thoại/tin.
- Đóng gói.
- Định tuyến.
- Hàng đợi.
- Chống trùng.
- Đọc lại sau reconnect.

Kiểm chứng bằng đầu nối thử nghiệm, không cần AI suy luận để chứng minh vận chuyển đúng.

### Giai đoạn 3 — Session, kích hoạt và hai provider đầu tiên

- Tạo/kết nối agent.
- Đưa tin vào đúng phiên.
- Kích hoạt lượt xử lý.
- Thu phản hồi.
- Chuyển phản hồi về đúng đích.

Kết quả cần đạt: **agent A và B trao đổi qua Coordinator mà Owner không phải chuyển tin thủ công.**

### Giai đoạn 4 — Web UI vận hành thật

- Owner tạo agent và phân vai.
- Gửi mục tiêu.
- Xem trao đổi.
- Gửi chỉ dẫn trực tiếp.
- Quan sát giao nhận và lỗi.

Thiết kế UI có thể thực hiện sớm; giai đoạn này nối vào backend thật.

### Giai đoạn 5 — Lane và quyền điều phối

Cho agent CEO được cấp quyền:

- Giao lane.
- Dùng agent hiện có.
- Yêu cầu tạo agent mới.
- Nhận phản hồi và phân công tiếp.

Backend chỉ thực hiện thao tác được yêu cầu, không thêm AI điều phối ngầm.

### Giai đoạn 6 — Công cụ Anvien và artifact

Nối code intelligence, báo cáo, diff và kết quả vào môi trường làm việc.

### Giai đoạn 7 — Telegram

Thêm kênh Owner từ xa trên nền giao tiếp đã hoạt động.

### Giai đoạn 8 — Kiểm chứng toàn hệ thống

- Nhiều agent cùng model.
- Nhiều provider.
- Tin đến khi agent đang làm.
- Mất kết nối và restart.
- Tin trùng hoặc đến muộn.
- Agent kết thúc giữa lúc giao tin.
- Thay agent được gán vai trò.
- Nhiều lane thao tác cùng repo.

Các bảo đảm này được thiết kế từ đầu và kiểm thử dần, không đợi giai đoạn cuối mới xử lý.

# 19. Hình ảnh cuối cùng

Bạn mở Anvien, chọn **Gemini/Antigravity làm CEO**, **GPT/Codex làm coder**, một agent khác làm reviewer, rồi gửi mục tiêu.

Các agent trình bày nội dung trong cùng không gian làm việc. Coordinator chuyển tin đến đúng agent. Agent nhận được kích hoạt, đọc và phản hồi hoặc gọi công cụ để hành động. CEO có thể tiếp tục phân lane theo quyền bạn cấp.

Bạn nhìn thấy toàn bộ quá trình và tham gia bất cứ lúc nào.

**Anvien là môi trường. Vai trò là cách Owner tổ chức đội. Agent/provider/model là bên thực hiện. Coordinator là đường liên lạc. UI chung là nơi tất cả cùng làm việc.**

