---
name: working-rules
description: This skill should be used when a work session begins, before the session selects skills, uses tools, edits files, validates results, or commits work.
---
Quy tắc làm việc
1. Rule là tầng điều phối cao nhất
•	Trước khi làm bất kỳ việc gì, phải đọc đầy đủ AGENTS.md.
•	AGENTS.md điều phối cách tìm rule, chọn skill, dùng công cụ, triển khai, kiểm tra và nghiệm thu.
•	Thứ tự bắt buộc là: rule trước, skill sau.
•	Không được dùng bản tóm tắt, trí nhớ phiên trước hoặc context đã rút gọn thay cho nội dung rule gốc.
•	Nếu context bị rút gọn hoặc task thay đổi, phải đọc lại nguồn gốc cần thiết trước khi hành động.
2. Mọi quyết định phải dựa trên bằng chứng
•	Không suy diễn từ lời nói.
•	Không suy diễn từ từ khóa.
•	Không suy diễn từ tên file, tên module hoặc sự giống nhau về thuật ngữ.
•	Không mang pattern từ hệ thống, dự án hoặc ứng dụng khác gán vào app này.
•	Không xem lời của user là chân lý tự động.
•	Ý kiến, phản biện hoặc giả thuyết của user là một góc nhìn cần được kiểm tra thêm.
•	Kết luận phải được chứng minh bằng rule, tài liệu nguồn, source code, runtime hoặc evidence tương ứng.
•	Không được hỏi user về điều đã được ghi lại trong repository; phải tự truy tìm bằng chứng trước.
3. Phải hiểu toàn bộ vấn đề trước khi xử lý
Khi user đưa ra một vấn đề:
•	Phải xác định toàn bộ phạm vi của vấn đề.
•	Phải đọc đầy đủ các rule và tài liệu liên quan.
•	Không được chỉ tìm một đoạn có từ khóa giống lời user nói rồi kết luận.
•	Không được chuyển trọng tâm sang chi tiết dễ sửa nhất.
•	Phải phân biệt nguyên nhân, biểu hiện, boundary và tác động liên quan.
Khi user lật lại một slice nhỏ:
•	Slice đó chỉ là một góc nhìn bổ sung.
•	Không được biến slice nhỏ thành toàn bộ vấn đề.
•	Không được bỏ mất mục tiêu và phạm vi ban đầu.
4. Phân loại pipeline theo liên kết thật
•	Mỗi pipeline phải được xác định theo command, state, ownership, boundary và kết quả của nó.
•	Không gom hai pipeline chỉ vì chúng dùng từ giống nhau.
•	Hai pipeline chỉ được coi là liên kết khi command của pipeline này thực sự cần state của pipeline kia để quyết định.
•	Không tự tạo quan hệ nghiệp vụ chỉ để làm kiến trúc trông đầy đủ hơn.
•	Không được làm phức tạp một luồng vốn đơn giản nếu repository không có bằng chứng cho sự phức tạp đó.
5. Tự tìm rule cho từng hành động
Trước khi sửa code, tài liệu, QA, artifact hoặc commit:
•	Phải xác định hành động đó chịu sự điều chỉnh của rule nào.
•	Phải tìm đủ các file chứa rule liên quan; rule có thể nằm ở nhiều nơi.
•	Không được cho rằng chỉ cần đọc một file là đủ.
•	Không được dựa vào context hội thoại để thay thế việc truy tìm rule.
•	Nếu không tìm thấy bằng chứng, phải tiếp tục điều tra trong repository trước khi hỏi user.
6. Chọn skill sau khi đã hiểu rule và nhiệm vụ
•	Không được chọn skill trước rồi ép vấn đề đi theo skill.
•	Phải xác định loại công việc thực tế trước, sau đó mới chọn skill.
•	Phải đọc đầy đủ SKILL.md của từng skill được sử dụng.
•	Không được rút gọn context hoặc chỉ lấy một số rule thuận tiện trong skill.
•	Rule riêng của skill chỉ áp dụng trong phạm vi task mà skill đó đảm nhiệm.
•	Trong lúc làm, nếu xuất hiện một loại công việc mới, phải tìm thêm skill phù hợp ngay lúc đó.
•	Không được giới hạn vào một bộ skill đã chọn từ đầu phiên.
•	Không gọi skill máy móc chỉ vì tên skill có từ giống vấn đề.
7. Dùng subagent có phạm vi
•	Chỉ gọi subagent khi có nhiệm vụ độc lập, cụ thể và có boundary rõ.
•	Không giao một vấn đề mơ hồ hoặc toàn bộ task cho subagent rồi chờ kết quả.
•	Subagent không được tự ý mở rộng phạm vi.
•	Kết quả subagent chỉ là đầu vào kiểm tra, không tự động trở thành kết luận.
•	Agent chính phải kiểm tra lại kết quả bằng bằng chứng và quy trình Supervisor.
8. Plan phải có trước code
•	Khi cần tạo, viết hoặc sửa plan, phải dùng planner.
•	Plan phải là tài liệu thật trong DOCS/plans.
•	Phải đọc đầy đủ toàn bộ các ledger liên quan của plan.
•	Chỉ triển khai slice đang được mở.
•	Không được tự chuyển phase hoặc slice.
•	Khi phát hiện vấn đề làm thay đổi phạm vi slice, phải cập nhật plan trước khi tiếp tục code.
•	Checklist, actual status, evidence và benchmark phải được cập nhật ngay khi trạng thái tương ứng thay đổi.
•	Chỉ chuyển sang slice tiếp theo sau khi slice hiện tại đã hoàn tất toàn bộ quy trình nghiệm thu và commit.
9. Áp dụng khi làm việc với Prototype: 
a. Prototype sẽ dẫn đường 
•	Prototype UI/UX là đối tượng đang được tinh chỉnh.
•	Không dùng SPEC để áp đặt ngược lên prototype.
•	Không đọc hoặc sử dụng product SPEC trong vòng tinh chỉnh hiện tại.
•	Không tự cập nhật SPEC.
•	Chỉ được cập nhật SPEC khi user ra đúng lệnh “cập nhật spec”.
•	Những câu tương đương không được tự suy diễn thành lệnh cập nhật SPEC.
•	ui-driven-spec vẫn phải được dùng để giữ đúng phương pháp UI-first, boundary và khả năng handoff; không được dùng nó để đảo chiều và ép SPEC lên prototype.
b. Prototype phải thể hiện luồng rõ ràng
Prototype không phải BE hoặc DB thật, nhưng phải thể hiện rõ:
•	Người dùng phát command nào.
•	State nào sở hữu dữ liệu.
•	State thay đổi ở đâu.
•	Presenter lấy dữ liệu từ đâu.
•	Module nào chịu trách nhiệm.
•	Boundary giữa các pipeline nằm ở đâu.
•	Luồng nào chỉ là dữ liệu demo.
•	Code thật sau này cần hiện thực điều gì.
Không được:
•	Giả lập kiến trúc backend không tồn tại.
•	Tạo database binding giả.
•	Thêm nghiệp vụ chưa được yêu cầu.
•	Dùng snapshot hoặc lớp trung gian không có bằng chứng (snapshot là nghiệp vụ thuộc duy nhất nghiệp vụ SYNC).
•	Biến app quản lý nhà hàng/khách sạn/quán bar… thành hệ thống kế toán hoặc một loại ứng dụng khác.
10. Bảo vệ phạm vi và trách nhiệm module
•	Chỉ sửa file/module thuộc phạm vi slice.
•	Không được đụng vào chức năng của tab/module khác nếu slice không sở hữu nó.
•	Các sibling module chỉ được kiểm tra bảo toàn.
•	Mỗi file chỉ được sở hữu một trách nhiệm nghiệp vụ chính .
•	Khi xuất hiện trách nhiệm mới, phải tìm hoặc tạo file/module đúng chủ sở hữu.
•	Một file có thể gọi nhiều module khác nhưng không được ôm hơn 1 nghiệp vụ không liên quan.
•	Không được tiện tay refactor hoặc sửa khu vực ngoài phạm vi.
11. Dùng Anvien theo đúng rule
•	Phải xem hướng dẫn Anvien trước khi sử dụng.
•	Phải làm mới graph trước mọi công việc dựa trên graph.
•	Trước khi sửa function, class, method, exported symbol, shared contract hoặc các đối tượng được rule liệt kê, phải chạy file-detail và impact analysis.
•	Phải báo cáo blast radius.
•	HIGH hoặc CRITICAL là cảnh báo phải làm cẩn thận, không phải lệnh cấm sửa.
•	Không được giảm hoặc che bớt bằng chứng graph để đầu ra trông đơn giản hơn.
•	Trước khi commit implementation phải chạy detect-changes.
12. Code đúng trước, QA sau
Thứ tự bắt buộc:
1.	Hiểu rule và hành vi cần có.
2.	Sửa production code.
3.	Sau khi hành vi đúng mới cập nhật QA.
4.	Chạy full build.
5.	Mở runtime thật mà người dùng nhìn thấy.
6.	Kiểm tra hành vi tại boundary thực.
7.	Chạy QA và thu evidence.
8.	Kiểm tra trực quan kết quả.
9.	Chạy regression cho các boundary cần bảo toàn.
Không được sửa test trước để ép test PASS trong khi production code chưa đúng.
13. QA phải chứng minh hành vi thật
•	Playwright phải là script tái sử dụng trong playwright/.
•	Không dùng script tạm làm evidence chính thức.
•	.tmp chỉ dùng để debug và phải nằm trong repository.
•	Evidence chính thức phải được lưu trong Reports/qa/playwright/....
•	Evidence phải có cả JSON và Markdown.
•	UI phải được kiểm tra trên runtime thật.
•	Phải kiểm tra screenshot bằng mắt, không chỉ tin assertion.
•	Không dùng test không liên quan, test cũ, test pass mặc định hoặc evidence đã bị thay thế để tuyên bố hoàn thành.
•	Regression phải kiểm tra cả chức năng đang sửa và các sibling boundary có nguy cơ bị ảnh hưởng.
14. Artifact phải được quản lý theo vòng đời
•	Artifact FAIL, retry, duplicate hoặc đã bị evidence mới thay thế là dead work.
•	Dead work phải được xóa khi kết thúc phase/slice tương ứng.
•	Không được để artifact trung gian tích tụ.
•	Trước khi xóa phải xác định chính xác artifact nào đã bị thay thế.
•	Evidence còn hiệu lực phải được giữ lại.
•	Nếu user yêu cầu giữ hoặc commit artifact cụ thể, phải xử lý đúng phạm vi đó.
15. Nghiệm thu theo zero-trust
•	Không được tự tuyên bố hoàn thành dựa trên việc code đã chạy.
•	Mọi completion claim phải qua Supervisor.
•	Supervisor phải kiểm tra độc lập source, diff, runtime, report, screenshot và evidence cần thiết.
•	Kết quả từ coder, QA hoặc subagent không tự động được chấp nhận.
•	Chỉ đánh dấu slice hoàn thành khi Supervisor PASS.
•	Nếu Supervisor REJECT, chỉ sửa invariant bị từ chối và thực hiện lại vòng kiểm tra.
16. Commit theo từng slice
Trước commit implementation phải có:
•	Production code hoàn thành.
•	Full build PASS.
•	Runtime validation PASS.
•	QA evidence hợp lệ.
•	Regression cần thiết PASS.
•	Artifact trung gian được xử lý.
•	Ledger được cập nhật.
•	Supervisor PASS.
•	Anvien detect-changes hoàn tất.
Mỗi slice phải có commit độc lập. Không gộp slice tiếp theo vào commit hiện tại.
17. Cách trao đổi trong quá trình làm việc
•	Phải báo đang làm gì, dựa trên rule hoặc bằng chứng nào.
•	Không được im lặng quá lâu khi đang build, QA hoặc chờ subagent.
•	Không báo cáo phỏng đoán như sự thật.
•	Phải phân biệt rõ: đã xác minh, đang kiểm tra và chưa có bằng chứng.
•	Khi gặp blocker, phải nêu chính xác blocker nằm ở đâu.
•	Không đẩy quyết định ngược lại cho user nếu repository đã có luật.
•	Khi user sửa hoặc phản biện, phải kiểm chứng lại toàn bộ vấn đề liên quan.
•	Khi user yêu cầu dừng, phải dừng ngay; không tiếp tục code, QA, xóa, cập nhật tài liệu, commit hoặc điều khiển subagent.
Đây mới là bộ quy tắc làm việc, không chứa tiến độ, trạng thái bàn giao hay kết quả của plan.
