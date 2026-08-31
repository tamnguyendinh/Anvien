# Báo Cáo Thí Nghiệm: Kết Hợp Graph Anvien Với Language Tooling Để Nâng Độ Chính Xác `file-detail`

## Metadata

- Thời gian báo cáo: `2026-08-31 03:32:29 +07:00`
- Repository: `E:\Anvien`
- Lệnh Anvien được thí nghiệm: `file-detail`
- File mục tiêu: `internal/aicontext/aicontext.go`
- Ngôn ngữ: Go
- Language tool dùng để đối chứng: `gopls v0.22.0`
- Trạng thái: thí nghiệm hoàn tất, chưa sửa production code
- Worktree sau thí nghiệm: sạch
- Benchmark update: `2026-08-31 04:14:54 +07:00` (không tính thời gian `analyze`)

## 1. Vấn Đề

Anvien có graph toàn repository. Graph giúp AI nhanh chóng tìm được file, symbol, flow và vùng code có khả năng liên quan mà không phải tự đọc hàng chục hoặc hàng trăm file.

Tuy nhiên, `file-detail` hiện lấy graph projection của file làm kết quả cuối. Cách này đưa vào output:

- local variable và intermediate node không cần thiết để hiểu cấu trúc file;
- quan hệ được gắn sai loại;
- quan hệ bị nối nhầm sang symbol cùng tên ở file hoặc package khác;
- import cấp package bị diễn giải thành quan hệ với chính file dù symbol thật nằm ở sibling file;
- hàng loạt source site bị báo `unresolved` dù compiler/language server giải được chính xác;
- metadata và graph fact làm output lớn nhưng không giúp AI quyết định công việc nhanh hơn.

Ở chiều ngược lại, nếu không có Anvien, AI vẫn có thể tìm quan hệ chính xác bằng các lệnh của ngôn ngữ như `gopls symbols`, `gopls call_hierarchy`, `gopls references`, `gopls definition` và `gopls check`. Nhưng AI phải tự tìm đúng file, đúng symbol, đúng vị trí và có thể phải chạy nhiều lệnh trên phạm vi lớn trước khi biết cần đọc gì.

Hai phía đang có hai năng lực bổ sung cho nhau:

```text
Anvien có graph để khoanh vùng nhanh.
Language tool có compiler knowledge để xác minh chính xác.
```

Phần còn thiếu là kết hợp trực tiếp hai năng lực này bên trong một lệnh Anvien.

## 2. Nguyên Tắc Cốt Lõi

Anvien là tool, không phải AI.

Mục tiêu không phải đưa AI, agent, planner hoặc mô hình suy luận vào Anvien. Mục tiêu là xác định những câu lệnh cụ thể mà AI vốn phải tự chạy để kiểm tra source, rồi cho `file-detail` gọi chính các câu lệnh hoặc capability tương đương đó.

Ví dụ đơn giản:

```text
Muốn xem thư mục → chạy dir.
Muốn biết symbol thật của file Go → chạy gopls symbols.
Muốn biết caller/callee thật → chạy gopls call_hierarchy.
Muốn biết reference thật → chạy gopls references.
Muốn biết target thật của một use site → chạy gopls definition.
Muốn biết unresolved thật → chạy gopls check.
```

Anvien không cần tự “suy nghĩ” như AI. Anvien chỉ cần:

1. dùng graph để chọn đúng phạm vi cần kiểm tra;
2. gọi đúng language command trên phạm vi đó;
3. hợp nhất kết quả đã được language tool xác minh;
4. trả output cuối cho AI.

## 3. Lý Do Chọn Thí Nghiệm Này

Thay vì tiếp tục tranh luận kiến trúc ở mức trừu tượng, thí nghiệm chọn một file thật và chạy hai đường độc lập:

### Đường A — Anvien hiện tại

```powershell
anvien file-detail internal/aicontext/aicontext.go \
  --repo E:\Anvien \
  --json \
  --format expanded \
  --relationships -1 \
  --unresolved -1 \
  --linked -1
```

Đường này ghi nhận toàn bộ những gì `file-detail` hiện trả về.

### Đường B — Cách AI kiểm tra nếu không dùng Anvien

Đường này không dùng thêm command Anvien để tìm quan hệ. Nó dùng source và các lệnh Go thông thường:

```powershell
Get-Content internal/aicontext/aicontext.go
go list -json ./internal/aicontext
gopls symbols internal/aicontext/aicontext.go
gopls call_hierarchy <file:line:column>
gopls references <file:line:column>
gopls definition <file:line:column>
gopls check internal/aicontext/aicontext.go
go test -run '^$' ./internal/aicontext
```

Thí nghiệm này được chọn vì nó trả lời trực tiếp ba câu hỏi:

1. `file-detail` hiện đúng ở đâu?
2. `file-detail` hiện thừa, thiếu hoặc sai ở đâu?
3. Câu lệnh ngôn ngữ cụ thể nào tạo ra kết quả chính xác hơn?

Nhờ vậy, hướng nâng cấp có thể xuất phát từ bằng chứng runtime thay vì từ giả thuyết về compact, pagination hoặc một ontology mới chưa được kiểm chứng.

## 4. Mong Muốn Đạt Được Sau Thí Nghiệm

Thí nghiệm nhằm xác minh phương pháp kết hợp mới:

```text
Graph Anvien khoanh vùng
        +
Language command xác minh trong phạm vi đã khoanh
        =
Output Anvien chính xác, ngắn và có proof
```

Trạng thái mong muốn của `file-detail` sau khi được nâng cấp:

- trả đúng những declaration thật của file;
- trả đúng caller, callee và reference;
- phân biệt relation do compiler/language server xác nhận với candidate chỉ đến từ graph;
- không nối nhầm symbol cùng tên;
- không coi import package là quan hệ với mọi sibling file trong package;
- không báo unresolved khi language tool giải được;
- mỗi relation mang file, source range và proof đủ để truy vết;
- primary output chỉ chứa relation đã được xác minh;
- graph fact chưa xác minh, nếu cần giữ để chẩn đoán, phải nằm ngoài primary answer.

Mục tiêu cuối cùng là output đạt độ chính xác tuyệt đối trong phạm vi mà language authority của ngôn ngữ xác minh được. AI nhận output này chỉ việc đọc đúng file và tiến hành chỉnh sửa hoặc thực hiện yêu cầu của user; AI không cần chạy lại cùng chuỗi `symbols`, `references`, `definition`, `call_hierarchy` và `check`.

Lợi ích trực tiếp:

- giảm số tool call của AI;
- giảm số file AI phải mở;
- giảm token dùng để đọc output graph noise;
- giảm thời gian điều tra codebase;
- tránh AI phải tái xác minh những điều Anvien có thể xác minh sẵn;
- làm Anvien trở thành một command system trả lời hoàn chỉnh hơn, thay vì chỉ trả candidate graph để AI tự điều tra lại.

## 5. Kết Quả Thí Nghiệm Tóm Tắt

### 5.1 Cấu trúc file

`file-detail` trả:

- 79 symbols;
- `exportedSymbols = 0`.

`gopls symbols` trả:

- 18 declaration cấp file;
- 10 struct fields;
- tổng cộng 28 thành phần có ý nghĩa để hiểu cấu trúc file.

`file-detail` đã đưa package node và hàng chục local variable vào cùng tập “symbols”, trong khi `gopls` trả declaration structure phù hợp hơn với câu hỏi “file chứa gì”.

`file-detail` cũng báo sai `exportedSymbols = 0` dù file có nhiều exported declarations như `Options`, `Stats`, `Result`, `Generate`, `GenerateAIContextFiles`, `RenderSkillSelectionGuide`, `RenderSkillSelectionGuideForRepo` và `FormatCrossRepoGroupsSection`.

### 5.2 Caller/callee trong repository

Sau khi chạy `gopls call_hierarchy` cho 11 function cấp file và so sánh logical call relations:

| Tập kết quả | Số relation |
|---|---:|
| `file-detail` | 27 |
| `gopls` | 25 |
| Hai bên cùng tìm thấy | 24 |
| Chỉ `file-detail` có | 3 |
| Chỉ `gopls` có | 1 |

Ba relation chỉ `file-detail` có:

```text
Generate → Result
statsFromRun → Stats
installBaseSkills → baseSkillInstallResult
```

Đây là type/composite-literal construction nhưng bị Anvien gắn loại `CALLS`.

Relation `gopls` tìm được nhưng Anvien thiếu:

```text
installBaseSkills → SkillInstallResult.Summary
```

Kết quả cho thấy graph hiện có recall tốt để tìm candidate, còn `gopls` cung cấp semantic authority để xác nhận loại và target thật.

### 5.3 Quan hệ nối nhầm

`file-detail` trả:

```text
renderAnvienBlock → USES → internal/filecontext.Builder
RenderSkillSelectionGuide → USES → internal/filecontext.Builder
```

Nhưng:

```powershell
gopls definition E:\Anvien\internal\aicontext\aicontext.go:95:22
```

trả:

```text
C:\Program Files\Go\src\strings\builder.go
type strings.Builder
```

Anvien đã nối `strings.Builder` sang một `Builder` cùng tên trong `internal/filecontext`.

`file-detail` cũng đưa `internal/cli/setup_command.go` vào inbound của `aicontext.go`. `gopls definition` chứng minh các symbol được file đó dùng là `SkillInstallResult` và `InstallSkillPackagesTo`, đều nằm trong `skill_packages.go`, không nằm trong `aicontext.go`.

### 5.4 Unresolved

`file-detail` báo 186 unresolved:

| Target text | Số lượng |
|---|---:|
| `writeCommandRow` | 52 |
| `builder.WriteString` | 48 |
| `strings.TrimSpace` | 8 |
| `filepath.Join` | 7 |
| `os.WriteFile` | 5 |
| các target khác | 66 |

Nhưng:

- `gopls references` tìm chính xác toàn bộ 52 reference của `writeCommandRow`;
- `gopls definition` giải đúng `builder.WriteString` về standard library;
- các command tương tự giải được `os.ReadFile`, `os.WriteFile`, `regexp.FindStringIndex`, `filepath.Join`, `codexResult.Summary` và chuỗi field access `run.Metrics.Files.Scanned`;
- `gopls check` không trả diagnostic;
- `go test -run '^$' ./internal/aicontext` compile package thành công.

Do đó 186 ở đây không phải 186 unresolved thật của source. Chúng chủ yếu là analyzer gaps mà language tool đã có thể xác minh trực tiếp.

### 5.5 Test boundary

`go test ./internal/aicontext` có một failure vì test còn chờ chuỗi generated rule cũ. Failure này không liên quan tới identifier resolution của `aicontext.go`.

Compile-only boundary:

```powershell
go test -run '^$' ./internal/aicontext
```

PASS.

## 6. Phần Còn Thiếu Cụ Thể Của `file-detail`

Phần còn thiếu không phải một lớp AI mới. Với file Go, đó là bước language verification:

```text
1. Graph tìm file và candidate symbols/relations.
2. gopls symbols xác nhận declaration structure.
3. gopls call_hierarchy xác nhận caller/callee.
4. gopls references xác nhận usage sites.
5. gopls definition xác nhận exact target của candidate edge.
6. gopls check và Go compiler xác nhận unresolved thật.
7. file-detail chỉ xuất primary relations đã được xác nhận.
```

Pipeline hiện tại:

```text
graph projection → file-detail output
```

Pipeline được thí nghiệm chứng minh là khả thi:

```text
graph candidates
    → language-native verification commands
    → verified declarations and relations
    → file-detail output
```

## 7. Hướng Áp Dụng Cho Anvien

### Trong phạm vi hiện tại

Chỉ dùng `file-detail` làm command mẫu. Chưa mở rộng implementation sang các command khác trong báo cáo này.

Đối với Go, adapter tối thiểu cần cung cấp capability tương đương:

| Câu hỏi của `file-detail` | Language command |
|---|---|
| File chứa declaration nào? | `gopls symbols <file>` |
| Function gọi hoặc được gọi bởi ai? | `gopls call_hierarchy <position>` |
| Declaration được dùng ở đâu? | `gopls references <position>` |
| Use site trỏ đến symbol/file nào? | `gopls definition <position>` |
| Source còn unresolved thật không? | `gopls check <file>` |
| Package có type-check được không? | `go test -run '^$' <package>` hoặc compiler-equivalent check |

Anvien có thể dùng graph để tránh chạy language command trên toàn repository. Nó chỉ chạy trên:

- file đã được target;
- declaration thật trong file;
- candidate use sites do graph chỉ ra;
- candidate caller/callee cần xác minh;
- unresolved sites cần kiểm tra lại.

Đây là điểm kết hợp tạo hiệu quả: graph giảm phạm vi, language tool tăng độ chính xác.

### Sau khi phương pháp được xác nhận cho `file-detail`

Các command khác của Anvien có thể áp dụng cùng nguyên tắc:

```text
Anvien graph khoanh vùng candidate
    +
tool chuyên môn xác minh candidate
    =
output cuối chính xác hơn
```

Tuy nhiên, việc mở rộng sang `impact`, `context`, `query`, `detect-changes` hoặc command khác không nằm trong phạm vi implementation của thí nghiệm này.

## 8. Output Mong Muốn Cho AI

Một `file-detail` đã nâng cấp phải đủ để AI chuyển thẳng sang làm việc:

```text
File: internal/aicontext/aicontext.go

Declarations:
- GenerateAIContextFiles, Function, exported, lines 61-92
- RenderSkillSelectionGuideForRepo, Function, exported, lines 197-203
- ...

Verified outbound calls:
- GenerateAIContextFiles → SkillPackagesForRepo
  proof: gopls call hierarchy, aicontext.go:62
- installBaseSkills → SkillInstallResult.Summary
  proof: gopls definition/call hierarchy, aicontext.go:350
- ...

Verified inbound calls:
- generateAnalyzeAIContext → GenerateAIContextFiles
  proof: analyze_postrun.go:81
- writeSkillHelp → RenderSkillSelectionGuideForRepo
  proof: skill_command.go:35
- ...

Unresolved:
- none according to gopls check and package type-check

Rejected graph candidates:
- strings.Builder ↛ internal/filecontext.Builder
- setup_command.go ↛ aicontext.go declarations
```

AI nhận output này không cần chạy lại:

- `rg` để tìm caller;
- `gopls symbols` để lọc local variables;
- `gopls references` để kiểm tra usage;
- `gopls definition` để sửa target nối nhầm;
- `gopls check` để xác nhận unresolved.

AI chỉ cần mở các file mà output đã chứng minh là liên quan và tiến hành chỉnh sửa.

## 9. Benchmark Thời Gian (Không Tính `analyze`)

### 9.1 Phạm vi đo

Tất cả phép đo dưới đây chạy trên cùng file `internal/aicontext/aicontext.go`, cùng repository và cùng máy. `analyze --force` được chạy trước để làm mới graph theo rule freshness, nhưng stopwatch chỉ bắt đầu sau khi lệnh đó hoàn tất; **0 giây của `analyze` được cộng vào các con số**. Mỗi phép đo được chạy 3 lần khi phù hợp; báo cáo dùng median cho lệnh `file-detail`. Bộ language command được chạy tuần tự đúng như quy trình điều tra thủ công để có tổng thời gian bảo thủ, tái lập được.

Đường Anvien được đo bằng:

```powershell
anvien file-detail internal/aicontext/aicontext.go --repo E:\Anvien --json
```

Đường language tooling gồm 53 command/process lần lượt:

- đọc source bằng `Get-Content`;
- `go list -json ./internal/aicontext`;
- `gopls symbols` một lần;
- `gopls call_hierarchy` cho 11 declaration/function;
- `gopls references` cho 18 declaration/candidate;
- `gopls definition` cho 19 candidate use site;
- `gopls check` một lần;
- `go test -run '^$' ./internal/aicontext` để compile-only.

Trong 53 command tổng cộng có 50 language-tool query/check và 3 command đọc/compile hỗ trợ. Không có command Anvien nào nằm trong bộ 53 command này.

### 9.2 Kết quả đo

| Đường chạy | Lần đo / số command | Tổng thời gian | Median / ghi chú |
|---|---:|---:|---:|
| `file-detail` human mặc định | 3 | 10.9987–11.5658 s | **11.0789 s** |
| `file-detail --json` (mốc chính cho AI) | 3 | 10.8928–11.1729 s | **11.0197 s** |
| `file-detail --json --format expanded --... -1` | 3 trước đó | 10.5463–10.5829 s | **10.5745 s** |
| Chuỗi không dùng Anvien, tuần tự | 53 command | **491.2436 s** | ≈ **8 phút 11.24 s** |

Phân rã của một lần chạy chuỗi 53 command:

| Nhóm lệnh | Số command | Tổng thời gian |
|---|---:|---:|
| `gopls call_hierarchy` | 11 | 113.7728 s |
| `gopls references` | 18 | 172.4282 s |
| `gopls definition` | 19 | 184.8975 s |
| `gopls symbols` | 1 | 7.9229 s |
| `gopls check` | 1 | 9.4539 s |
| `go list` | 1 | 1.0947 s |
| đọc source | 1 | 0.0612 s |
| compile-only | 1 | 1.5757 s |
| **Tổng wall-clock** | **53** | **491.2436 s** |

### 9.3 Tổng thời gian nếu kết hợp

Dùng machine-output mốc chính:

```text
file-detail --json       11.0197 s
+ language verification 491.2436 s
---------------------------------
= combined               502.2633 s
```

Tổng thời gian đo được là **502.2633 giây**, tương đương khoảng **8 phút 22.26 giây**.

Nếu dùng expanded full-detail làm mốc `file-detail`, tổng là **501.8181 giây**, tương đương khoảng **8 phút 21.82 giây**.

Đây là tổng theo cách chạy tuần tự và spawn riêng từng CLI query. Nó trả lời đúng câu hỏi “nếu Anvien thực hiện toàn bộ chuỗi thao tác mà AI đã chạy thì mất bao lâu”, nhưng chưa phải latency tối ưu của một implementation tích hợp.

### 9.4 Diễn giải benchmark

Phần lớn thời gian của đường B nằm ở việc mỗi truy vấn `gopls` tự khởi tạo hoặc nạp lại workspace. Một thử nghiệm batch song song đã tạo hơn 20 process `gopls` và bị nghẽn tài nguyên, nên không dùng kết quả đó làm số liệu chính.

Vì vậy cần tách hai kết luận:

1. **Độ chính xác:** đã được chứng minh bằng các kết quả `gopls` và compile/check ở phần trước.
2. **Thời gian tích hợp:** con số hiện đo được là khoảng 8 phút 22 giây nếu gọi nguyên chuỗi CLI tuần tự; để đưa vào `file-detail` production cần tái sử dụng một language-server session hoặc capability tương đương. Chưa có benchmark session persistent thành công trong thí nghiệm này, nên không đưa ra con số ước lượng thay thế.

Anvien vẫn giúp tiết kiệm thời gian điều tra của AI ở tầng workflow: AI chỉ gọi một `file-detail` và nhận kết quả đã xác minh, thay vì tự phát hiện file, chọn vị trí, chạy 50 query, đọc lại output và tự đối chiếu. Benchmark trên đo CPU/wall-clock mà Anvien phải thực hiện; nó không phủ nhận phần tiết kiệm token, số lượt tương tác và thời gian chờ của AI.

## 10. Kết Luận

Thí nghiệm trên `aicontext.go` xác nhận phương pháp kết hợp là đúng và thực dụng:

1. Graph Anvien rất hữu ích để khoanh vùng nhanh.
2. Graph projection chưa đủ làm kết quả cuối của `file-detail`.
3. Language tooling giải được chính xác nhiều quan hệ mà graph hiện gắn sai, bỏ sót hoặc đánh dấu unresolved.
4. Không cần biến Anvien thành AI.
5. Cần đưa đúng câu lệnh/capability của ngôn ngữ vào pipeline `file-detail` sau bước graph narrowing.
6. Kết quả cuối có thể giúp AI bắt đầu chỉnh sửa ngay mà không phải lặp lại toàn bộ chuỗi điều tra.

Tuyên bố kết quả của thí nghiệm:

> `file-detail` nên chuyển từ “graph projection output” thành “graph-guided, language-verified output”. Graph chịu trách nhiệm khoanh vùng; language tool chịu trách nhiệm xác minh; Anvien chịu trách nhiệm hợp nhất và trả một câu trả lời cuối đủ chính xác cho AI hành động.

---

# Phụ Lục A — Toàn Bộ Phản Hồi Kết Quả Thí Nghiệm Đã Gửi Trong Cuộc Trao Đổi

Thí nghiệm xác nhận đúng hướng bạn nói: với Go, phần còn thiếu của `file-detail` là gọi các lệnh `gopls` để kiểm chứng source sau khi graph đã khoanh vùng file.

Không cần AI bên trong Anvien. Chỉ cần Anvien gọi đúng tool ngôn ngữ.

## Đường A — `file-detail`

Với [aicontext.go](E:/Anvien/internal/aicontext/aicontext.go), `file-detail` trả:

- 79 symbols.
- `exportedSymbols = 0`.
- 23 local relations.
- 15 outbound relations.
- 16 inbound relations.
- 186 unresolved.
- 248.019 ký tự JSON expanded.

Nó khoanh được những vùng gần đúng như:

- `skill_packages.go`
- `analyze.go`
- `analyze_postrun.go`
- `skill_command.go`
- các file test

Đây là giá trị của graph: tìm vùng liên quan rất nhanh.

## Đường B — Không dùng Anvien

### 1. Lấy thành phần thật của file

```powershell
gopls symbols E:\Anvien\internal\aicontext\aicontext.go
```

Kết quả:

- 18 declaration cấp file.
- 10 struct fields.
- Tổng cộng 28 thành phần có ý nghĩa.
- Không đưa 51 local variables và package node vào danh sách chính.

Ngoài ra, `file-detail` báo không có exported symbol, nhưng source thực tế có nhiều export như:

- `Options`
- `Stats`
- `Result`
- `Generate`
- `GenerateAIContextFiles`
- `RenderSkillSelectionGuide`
- `RenderSkillSelectionGuideForRepo`
- `FormatCrossRepoGroupsSection`

Như vậy `gopls symbols` đã sửa ngay phần “file chứa gì”.

### 2. Lấy caller và callee thật

Ví dụ:

```powershell
gopls call_hierarchy E:\Anvien\internal\aicontext\aicontext.go:61:6
```

Nó trả chính xác:

```text
GenerateAIContextFiles
├─ Caller:
│  ├─ Generate
│  ├─ generateAnalyzeAIContext
│  └─ 4 test functions
└─ Callee:
   ├─ SkillPackagesForRepo
   ├─ renderAnvienBlock
   ├─ upsertSection
   ├─ installBaseSkills
   ├─ removeGeneratedSkills
   └─ filepath.Join
```

Mỗi relation có file và call range chính xác.

Mình chạy lệnh này cho 11 function cấp file và đối chiếu:

| Kết quả call trong repo | Số lượng |
|---|---:|
| `file-detail` | 27 |
| `gopls` | 25 |
| Hai bên cùng tìm thấy | 24 |
| Chỉ `file-detail` có | 3 |
| Chỉ `gopls` có | 1 |

Ba relation chỉ Anvien có là:

```text
Generate → Result
statsFromRun → Stats
installBaseSkills → baseSkillInstallResult
```

Đây là composite literal/type construction, không phải function call. Anvien gắn sai loại `CALLS`.

Relation Anvien thiếu là:

```text
installBaseSkills → SkillInstallResult.Summary
```

`gopls` xác định method nằm tại [skill_packages.go](E:/Anvien/internal/aicontext/skill_packages.go:196).

Như vậy, trên ví dụ này:

- Graph tìm được 24/25 call relation thật trong repo.
- `gopls` loại ba relation gắn sai nghĩa.
- `gopls` bổ sung một relation graph bỏ sót.

Đây chính xác là sự kết hợp mà bạn đang nói.

### 3. Kiểm tra các quan hệ bị nối nhầm file

`file-detail` trả:

```text
renderAnvienBlock → USES → internal/filecontext.Builder
RenderSkillSelectionGuide → USES → internal/filecontext.Builder
```

Nhưng chạy:

```powershell
gopls definition E:\Anvien\internal\aicontext\aicontext.go:95:22
```

Kết quả thật:

```text
C:\Program Files\Go\src\strings\builder.go
type strings.Builder
```

Tức là Anvien nối nhầm `strings.Builder` với một `Builder` cùng tên trong `filecontext`.

`file-detail` cũng nói `setup_command.go` có inbound relation tới `aicontext.go`. Nhưng `gopls definition` chứng minh các symbol mà file đó dùng là:

```text
SkillInstallResult       → skill_packages.go
InstallSkillPackagesTo   → skill_packages.go
```

Do đó `setup_command.go` không có relation thật với `aicontext.go`. Nó chỉ import cùng package.

### 4. Xác minh “unresolved”

`file-detail` báo 186 unresolved, trong đó:

```text
52 × writeCommandRow
48 × builder.WriteString
8  × strings.TrimSpace
7  × filepath.Join
5  × os.WriteFile
...
```

Nhưng:

```powershell
gopls references E:\Anvien\internal\aicontext\aicontext.go:119:2
```

tìm chính xác cả 52 lần dùng `writeCommandRow`.

```powershell
gopls definition E:\Anvien\internal\aicontext\aicontext.go:96:15
```

xác định `builder.WriteString` là method chuẩn của `strings.Builder`.

Các lệnh tương tự giải được:

- `os.ReadFile`
- `os.WriteFile`
- `regexp.FindStringIndex`
- `filepath.Join`
- `codexResult.Summary`
- `run.Metrics.Files.Scanned`

Cuối cùng:

```powershell
gopls check E:\Anvien\internal\aicontext\aicontext.go
go test -run '^$' ./internal/aicontext
```

đều xác nhận file/package compile không có identifier unresolved.

Full test có một failure về expected generated rule text, không liên quan tới resolution.

## Phần thiếu của `file-detail` đã được khoanh vùng

Với Go, công thức rất đơn giản:

```text
Graph của Anvien
    ↓
Khoanh file và candidate relations
    ↓
gopls symbols <file>
gopls call_hierarchy <symbol-position>
gopls references <symbol-position>
gopls definition <candidate-source-site>
gopls check <file>
    ↓
Chỉ trả relation đã được language tool xác nhận
```

Vai trò từng lệnh:

- `symbols`: file thật sự chứa gì.
- `call_hierarchy`: caller/callee thật.
- `references`: ai thực sự dùng declaration đó.
- `definition`: candidate edge trỏ chính xác tới symbol/file nào.
- `check`: unresolved nào thực sự tồn tại.

Kết luận của thí nghiệm:

> Graph của Anvien đang có recall khá tốt để khoanh vùng. Phần thiếu không phải AI hay một hệ suy luận phức tạp; đó là bước kiểm chứng bằng language tooling trước khi `file-detail` công bố relation.

Nói ngắn gọn hơn:

```text
Hiện tại:
graph projection → file-detail output

Cần có:
graph candidates → gopls verification → file-detail output
```

Đây là kết luận mình nghĩ đã bám đúng phương pháp bạn muốn tìm. Worktree không có thay đổi.
