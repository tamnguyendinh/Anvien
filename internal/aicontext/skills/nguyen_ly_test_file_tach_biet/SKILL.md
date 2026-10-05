---
name: nguyen-ly-test-file-tach-biet
description: (Use when) Dùng khi cần tạo, review, hoặc tổ chức cấu trúc test file (unit, component) theo mô hình tách biệt 100% khỏi mã nguồn sản phẩm (Electron/Desktop/Frontend), hoặc theo Go convention (_test.go cùng package). Bao gồm fixtures, path aliases, build isolation, phân định rõ giữa electron/test/ (nội bộ) với playwright/ (E2E), và quy tắc riêng cho Go backend.
---

# Nguyên Lý Thư Mục Kiểm Thử Tách Biệt (electron/test/) cho tầng Desktop/Frontend.

# 1. Tổng Quan Kiến Trúc
•	Nguyên lý cốt lõi: Tách bạch 100% mối quan tâm (Separation of Concerns).
o	Cây thư mục mã nguồn (renderer/, main/, preload/) chỉ chứa mã nguồn phục vụ kinh doanh và vận hành của app.
o	Toàn bộ mã kiểm thử nội bộ, mock data, fixture giả lập được gom gọn vào thư mục electron/test/.
o	Kiểm thử giao diện và hành vi toàn diện người dùng (E2E) giữ nguyên độc lập tại playwright/ ở root.

# 2. Bản Đồ Cây Thư Mục Chuẩn Hóa (Directory Tree)
```
G:\<root_repo>\
├── playwright/                         <-- [TẦNG 1] E2E & Visual Regression (Chạy app Electron thật)
│   ├── app/
│   │   ├── pos-checkout.spec.ts
│   │   └── table-management.spec.ts
│   └── shared/
│
├── electron/                           <-- Thư mục gốc Desktop App
│   ├── main/                           <-- [CODE SẢN PHẨM] Node.js Main Process (Window, IPC, Tray, HW)
│   │   ├── index.ts
│   │   ├── ipc-handlers.ts
│   │   └── printer-manager.ts
│   │
│   ├── preload/                        <-- [CODE SẢN PHẨM] ContextBridge an toàn giữa Main và UI
│   │   └── index.ts
│   │
│   ├── renderer/                       <-- [CODE SẢN PHẨM] React UI (Giao diện người dùng)
│   │   ├── src/
│   │   │   ├── components/
│   │   │   │   ├── TableCard.tsx
│   │   │   │   └── OrderSummary.tsx
│   │   │   ├── stores/
│   │   │   │   └── useCartStore.ts
│   │   │   ├── utils/
│   │   │   │   └── currency-formatter.ts
│   │   │   ├── App.tsx
│   │   │   └── main.tsx
│   │
│   └── test/                           <-- [TẦNG 2] TOÀN BỘ KIỂM THỬ NỘI BỘ (TÁCH BIỆT 100%)
│       ├── fixtures/                   <-- Dữ liệu mẫu (mock menus, orders, tables, settings)
│       │   ├── mock-menu-items.json
│       │   └── mock-table-states.json
│       │
│       ├── setup.ts                    <-- Cấu hình môi trường test (DOM ảo, mock IPC bridge)
│       │
│       ├── unit/                       <-- Test hàm thuần túy, logic tính toán, store state
│       │   ├── currency-formatter.test.ts
│       │   └── useCartStore.test.ts
│       │
│       ├── components/                 <-- Test render React component (soi chiếu với renderer/components)
│       │   ├── TableCard.test.tsx
│       │   └── OrderSummary.test.tsx
│       │
│       └── main/                       <-- Test logic của Node.js Main process
│           ├── ipc-handlers.test.ts
│           └── printer-manager.test.ts
```

# 3. Cơ Chế Đường Dẫn Thông Minh (Path Aliases)
Nhờ có cấu hình TypeScript (tsconfig.json) và Vite, các file test trong electron/test/ không bao giờ phải dùng đường dẫn tương đối dài dòng như ../../../../renderer/src/components/TableCard.
Thay vào đó, file test sẽ import trực tiếp qua các bí danh (Path Aliases) cực kỳ sạch sẽ:
// File: electron/test/components/TableCard.test.tsx

import { render, screen } from '@testing-library/react';
import { TableCard } from '@renderer/components/TableCard';    // <-- Import thẳng qua Alias @renderer
import { mockTables } from '../fixtures/mock-table-states.json'; // <-- Lấy dữ liệu mock tập trung

describe('TableCard Component', () => {
  it('Hiển thị đúng số bàn và trạng thái đang phục vụ', () => {
    render(<TableCard table={mockTables[0]} />);
    expect(screen.getByText('Bàn 01')).toBeInTheDocument();
  });
});

#4. Cơ Chế Đóng Gói (Build Isolation) Tuyệt Đối
1.	Khi Build ra app cài đặt (pnpm build / electron-builder):
o	File cấu hình tsconfig.json và vite.config.ts được thiết lập loại trừ thư mục test/:
{
	  "exclude": ["test/**/*", "**/*.test.*"]
}
o	Quá trình build hoàn toàn lướt qua thư mục test/. File binary của app bán hàng (POS Desktop) tuyệt đối sạch sẽ, không tốn dù chỉ 1 byte cho mã kiểm thử.
2.	Khi Chạy Test (pnpm test):
o	Trình test runner (như Vitest) chỉ quét duy nhất trong phạm vi electron/test/:
	// vitest.config.ts
	export default defineConfig({
	  test: {
	    include: ['test/**/*.{test,spec}.{ts,tsx}'],
	    environment: 'happy-dom'
	  }
  });
o	Tốc độ chạy test cực nhanh do gom gọn vùng quét.

# 5. So Sánh Phân Định Rõ Ràng: electron/test/ vs playwright/
Tiêu Chí	electron/test/ (Tầng 2 - Nội Bộ)	playwright/ (Tầng 1 - Ngoại Vi)
Vị trí	Nằm bên trong electron/test/	Nằm riêng biệt ngoài root playwright/
Bản chất	Unit Test & Component Test	End-to-End (E2E) & System Test
Môi trường chạy	Môi trường Node.js / DOM ảo (chạy mất vài phần nghìn giây)	Khởi động cả phần mềm Electron thật trên màn hình
Mục đích	Kiểm tra logic tính tiền, render đúng màu nút, validate form	Kiểm tra toàn bộ thao tác click chuột, in hóa đơn thật, IPC bridge thật
Báo cáo	In kết quả console	Xuất báo cáo Markdown + JSON vào Reports/qa/playwright/

# 6. Quy Tắc Riêng Cho Go Backend (_test.go)

Nguyên lý "tách biệt 100%" ở các section trên **chỉ áp dụng cho Electron/Desktop/Frontend**.
Go có convention riêng do thiết kế ngôn ngữ — **không được áp mô hình electron/test/ cho Go**.

## 6.1. Test file nằm cùng package với source

Go bắt buộc `_test.go` nằm **cạnh file source** trong cùng thư mục/package.
Đây không phải lộn xộn — đây là convention của ngôn ngữ.

```
internal/analyze/
├── analyze.go            <-- Code sản phẩm
├── analyze_test.go       <-- Test cho analyze.go (cùng package)
├── metrics.go
├── metrics_test.go
└── testdata/             <-- Dữ liệu test (Go convention, tương đương fixtures/)
    ├── sample-repo/
    └── expected-graph.json
```


**Không được** tạo thư mục `test/` riêng rồi dời `_test.go` ra khỏi package.

## 6.2. Hai chế độ test package

| Chế độ | Package declaration | Truy cập |
|--------|-------------------|----------|
| White-box | `package analyze` | Truy cập cả unexported (hàm viết thường) |
| Black-box | `package analyze_test` | Chỉ truy cập exported — test như consumer bên ngoài |

## 6.3. testdata/ thay cho fixtures/

Go toolchain tự nhận thư mục `testdata/` trong mỗi package:
- `go build` tự động bỏ qua `testdata/`
- Truy cập bằng đường dẫn tương đối: `os.ReadFile("testdata/sample.json")`
- Không cần path alias, không cần config exclude

## 6.4. Build isolation là tự động

- Go compiler **tự động loại** mọi file `_test.go` khỏi binary sản phẩm
- Không cần `tsconfig.exclude`, không cần config gì thêm
- `go build ./...` → sạch 100%, `go test ./...` → chạy test

## 6.5. So sánh nhanh

| Tiêu chí | Electron/Frontend | Go Backend |
|----------|-------------------|------------|
| Vị trí test | Tách ra `electron/test/` | Cùng package, cạnh source |
| Dữ liệu mock | `test/fixtures/` | `testdata/` trong package |
| Build isolation | Cần config exclude | Tự động bởi compiler |
| Path alias | `@renderer/...` | Không cần — import trực tiếp |
| Test runner | Vitest / Jest | `go test ./...` (built-in) |

# 7. Tổng Kết Lợi Ích Của phương pháp
1.	Sạch mắt & Ngăn nắp: Mở thư mục code nào là chỉ thấy code tính năng đó, không bị rối mắt bởi hàng chục file .test.tsx xen kẽ.
2.	Quản lý Mock Data chuẩn: Toàn bộ file JSON giả lập menu, bàn ăn, tài khoản mẫu được gom vào fixtures/, không vương vãi trong component.
3.	Tránh test mồ côi: Cây thư mục soi chiếu giúp nhìn vào là biết component nào đã có test, component nào chưa.
4.	Phù hợp tiêu chuẩn doanh nghiệp: Rất nhiều tập đoàn công nghệ áp dụng cách này cho các ứng dụng Frontend/Desktop lớn để kiểm soát chất lượng chặt chẽ.
