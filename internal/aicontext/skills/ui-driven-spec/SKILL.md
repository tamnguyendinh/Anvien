---
name: ui-driven-spec
description: >
  Workflow phát triển phần mềm UI-first cho AI agent. Dùng skill này khi người
  dùng muốn: xây app mới theo hướng UI trước BE, extract spec từ prototype đã
  có, chuẩn bị handoff BE implementation từ FE đã code, hoặc nói các cụm như
  "làm UI trước", "prototype rồi mới spec", "UI-driven", "FE trước BE",
  "extract contract từ UI", "slot map", "state map", "backend contract map",
  "chuyển giao kiến trúc sư duyệt". Trigger khi user có HTML prototype hoặc
  FE components và cần chuyển sang BE. Đây là flow ngược truyền thống —
  UI là bản đồ thao tác, Spec là luật tối cao, human phải duyệt trước khi implement.
---

# UI-Driven Spec — FE-First Development Workflow

## Triết lý cốt lõi

> **UI là bản đồ thao tác người dùng, không phải luật tối cao.**
> **Luật tối cao vẫn là Spec/Authority.**
> **Human duyệt trước khi có 1 dòng BE nào được viết.**

Truyền thống: `Spec → BE → FE` → gap phát hiện muộn, patch rối.

Flow này: `UI → FE → SPEC → Extract → Authority align → review → BE module-by-module`
→ gap phát hiện sớm nhất, fix rẻ nhất, không có module nào chạy khi chưa được duyệt.

---

## Thứ tự thực hiện

```
[1]  Prototype UI (HTML)
[2]  Frontend Components
[3]  SPEC
[4]  Data source map
[5]  API Contract File
[6]  UI Slot Map
[7]  State Map
[8]  Backend Contract Map
[9]  Authority/Spec chỉnh lại [4][5][6][7]
[10]  Chuyển giao kiến trúc sư / owner duyệt  ← GATE
[11] Plan implement từng module (cấm plan hàng loạt)
[12] Runtime QA sau mỗi module
```

> **Quy tắc cứng:**
> - Không skip, không đảo thứ tự
> - Bước [9] là gate bắt buộc — AI không tự proceed sang [10]
> - Bước [10] chỉ plan 1 module tại 1 thời điểm — QA xong mới plan tiếp

---

## Bước 1 — Prototype UI (HTML)

**Mục đích:** Tạo bản đồ thao tác người dùng. HTML tĩnh, không có logic.

**Input cần hỏi user:**
- Danh sách screens / flows cần cover
- Thứ tự ưu tiên (flow nào critical nhất)
- Design system hiện có không (màu, font, component library)

**Output:**
```
prototype/
├── index.html          ← nav giữa các screens
├── screen-{name}.html  ← mỗi screen 1 file
└── assets/             ← css, images tĩnh
```

**Checklist:**
- [ ] Mọi user action (click, submit, navigate) đều visible
- [ ] Empty states, loading states, error states đã mock
- [ ] Happy path + ít nhất 1 error path mỗi flow
- [ ] Không có logic thật — chỉ HTML/CSS

---

## Bước 2 — Frontend Components

**Mục đích:** Chuyển prototype thành components thật với state management.
Data hardcode / mock — **chưa gọi API thật**.

**Output:**
```
src/
├── components/
├── pages/
├── hooks/              ← data fetching mock
├── types/              ← TypeScript types định nghĩa data shape
└── mocks/              ← mock data phản ánh expected API response
```

**`types/` và `mocks/` là nguồn sự thật cho bước 4-5-6-7:**

```typescript
// types/order.ts — FE đang expect BE trả về gì
export interface Order {
  id: string
  status: 'pending' | 'confirmed' | 'cancelled'
  items: OrderItem[]
  total: number
  createdAt: string
}
```

**Checklist:**
- [ ] Mọi component render được từ mock data
- [ ] TypeScript types đủ cho mọi entity
- [ ] Mock data đúng shape FE cần từ BE
- [ ] Không có `any` ở data boundary

---

## Bước 3 — SPEC

**Mục đích:** Viết hoặc review SPEC dựa trên UI/FE đã có.
SPEC này sẽ là **chuẩn để align lại** các artifacts ở bước 8.
- Soạn SPEC theo yêu cầu kiến trúc sư / owner, không tự quyết định.
- Đọc lại, đánh dấu [CONFIRMED] / [NEEDS_REVIEW] / [MISSING] so với UI/FE thực tế
- Không override gì ở bước này — chỉ ghi nhận delta


**Output:** `docs/SPEC.md` (hoặc version mới nếu đã có)

**Checklist:**
- [ ] Mọi flow trong prototype có entry trong SPEC
- [ ] Business rules quan sát từ UI đã ghi vào SPEC
- [ ] Chỗ chưa rõ → [TBD] chứ không tự điền

---

## Bước 4: Data source map

This map describes where production data should come from after converting this static HTML prototype to Electron React + Go local backend or some other stack.

For endpoint, payload, and response shape expected by the frontend, see [backend-source-map.md](backend-source-map.md).

---

## Bước 5 — API Contract File

**Mục đích:** Ghi lại FE expects gì từ mỗi endpoint.
Đây là target BE phải hit — không phải BE tự thiết kế.
Chức năng: Chi tiết payload cho từng nút/form/input theo feature,

**Output file:** `DOCS\SPEC\IMPLEMENTATION-MAPS\API-payload\<feature>-API-payload.md`
Examples:

| Feature | File |
|---|---|
| Orders page | `orders-API-payload.md` |
| Tables/POS runtime | `tables-pos-API-payload.md` |
| Settings | `settings-API-payload.md` |

## Payload Doc Format

Each action should document:

| Section | Meaning |
|---|---|
| UI trigger | Button, row action, dialog submit, dropdown selection |
| Source fields | Fields read from UI state or user input |
| Backend surface | Local Go endpoint/service command |
| Required validation | FE affordance and Go service enforcement |
| Result projection | Which UI projection must refetch/update |
| Sync/audit rule | Event/outbox/ledger behavior when relevant |

Rules:

- Do not put payload details back into `backend-contract-map.md`, `data-source-map.md`, `state-map.md`, or `ui-slot-map.md`.
- Use `snake_case` for documented contract fields.
- FE must not send `owner_id`.
- FE disabled states are UX only. Go service still enforces permissions, active shift, locks, and hash-chain health.
- Prefer field tables over large sample JSON blocks. Add exact JSON only when it is necessary for a migration/test fixture.

```markdown
## API Contract

### Global
- Base URL: `/api/v1/`
- Auth: `Authorization: Bearer <jwt>`
- Response envelope: { "success": true, "data": {}, "error": null }

### Endpoints

#### POST /auth/login
Request:  { email: string, password: string }
Response: { token: string, user: User }
FE uses:  LoginForm.tsx → onSubmit

#### GET /orders
Request:  ?status=pending&page=1&limit=20
Response: { orders: Order[], total: number, page: number }
FE uses:  OrderList.tsx → useOrders hook

[Liệt kê đủ mọi endpoint FE cần]
```

> **Note:** Extract từ `types/` và `hooks/` của Bước 2 — không đoán.

---

## Bước 6 — UI Slot Map

**Mục đích:** Ghi lại cái gì render ở đâu và điều kiện hiển thị.

**Output file:** `docs/slot-map.md`

```markdown
## UI Slot Map

### Screen: OrderDashboard

| Slot | Component | Điều kiện hiển thị | Data source |
|------|-----------|-------------------|-------------|
| header | PageHeader | always | static |
| stats-row | StatsCard × 3 | role === 'admin' | GET /stats |
| order-list | OrderTable | orders.length > 0 | GET /orders |
| empty-state | EmptyOrders | orders.length === 0 | — |
| error-banner | ErrorBanner | fetchError !== null | — |
| pagination | Pagination | total > pageSize | from response |

### Conditional Renders
- `CreateOrderButton`: visible nếu user.permissions.includes('order:create')
- `CancelButton` mỗi row: visible nếu order.status === 'pending'
```

---

## Bước 7 — State Map

**Mục đích:** Ghi lại data flow, loading/error/empty states, transitions.

**Output file:** `docs/state-map.md`

```markdown
## State Map

### Global State
| Key | Type | Source | Persist? |
|-----|------|--------|----------|
| currentUser | User \| null | POST /auth/login | localStorage |
| authToken | string \| null | POST /auth/login | localStorage |

### Page State: OrderDashboard
| State | Type | Initial | Transitions |
|-------|------|---------|-------------|
| orders | Order[] | [] | ← GET /orders success |
| loading | boolean | true | true → false on fetch complete |
| error | string \| null | null | ← fetch error message |
| page | number | 1 | ← pagination click |

### User Action → State Transition
| Action | Trigger | State change | Side effect |
|--------|---------|-------------|-------------|
| Click "Cancel Order" | Button click | order.status = 'cancelling' | POST /orders/:id/cancel |
| Cancel success | API response | remove from list | toast success |
| Cancel fail | API error | revert status | toast error |
```

---

## Bước 8 — Backend Contract Map

**Mục đích:** Tổng hợp endpoint + payload + response shape mà FE cần,
nhóm theo module BE. Nguồn sự thật duy nhất cho BE implementation.

**Output file:** `docs/backend-contract-map.md`

```markdown
## Backend Contract Map

> Không implement thứ gì không có trong document này.

### Module: Auth
Endpoints: POST /auth/login, POST /auth/logout, POST /auth/refresh
DB tables: users, sessions
Business rules:
  - Login fail 5 lần → lock 15 phút
  - Token: access 15m, refresh 7d
FE triggers: LoginForm submit, auto-refresh khi token gần hết hạn
UI slots affected: header (user avatar), all protected routes

### Module: Orders
Endpoints: GET /orders, POST /orders, GET /orders/:id, POST /orders/:id/cancel
DB tables: orders, order_items, order_status_history
Business rules:
  - Chỉ cancel được nếu status === 'pending'
  - Cancel ghi vào order_status_history
FE triggers: OrderList load, CreateOrderForm submit, CancelButton click
UI slots affected: OrderTable, StatsCard, EmptyOrders

[Lặp lại cho mỗi module]

### Module dependency graph
Auth → prerequisite cho mọi module
Orders → Products (price lookup), Inventory (stock check)
```

---

## Bước 9 — Authority/Spec chỉnh lại [4][5][6][7]

**Mục đích:** Dùng SPEC (Bước 3) làm chuẩn để **align lại** toàn bộ 4 artifacts:
`api-contract.md`, `slot-map.md`, `state-map.md`, `backend-contract-map.md`.

Không phải Spec được viết lại từ artifacts — chiều ngược lại:
**Spec phán xét artifacts, artifacts phải conform theo Spec.**

**Quy trình:**

```
Với mỗi artifact trong [4][5][6][7]:

  1. So sánh từng item với SPEC
     → [OK]      — khớp Spec, giữ nguyên
     → [CONFLICT] — mâu thuẫn với Spec → sửa artifact theo Spec
     → [MISSING_IN_SPEC] — artifact có nhưng Spec không đề cập
                         → escalate, không tự quyết

  2. Với [CONFLICT]: ghi rõ "Spec §X.Y nói A, artifact đang ghi B → sửa thành A"

  3. Với [MISSING_IN_SPEC]:
     → Nếu rõ ràng là UI evidence (quan sát trực tiếp từ FE) → flag để bổ sung vào Spec
     → Nếu không chắc → [TBD], đưa vào danh sách câu hỏi cho Bước 9

  4. Update artifacts — ghi version + ngày chỉnh
```

**Output:** Các file [4][5][6][7] đã được align + `docs/alignment-notes.md`
(ghi lại mọi conflict đã resolve và câu hỏi còn [TBD])

**Checklist trước khi sang Bước 9:**
- [ ] Không còn conflict giữa Spec và bất kỳ artifact nào
- [ ] Mọi [MISSING_IN_SPEC] đã được flag rõ ràng
- [ ] Mọi [TBD] đã được list trong alignment-notes.md
- [ ] Không tự quyết định điều gì không có trong Spec

---

## Bước 10 — Chuyển giao kiến trúc sư / owner duyệt

**⛔ GATE BẮT BUỘC — AI agent không tự proceed sang Bước 10.**

**Package chuyển giao:**
```
docs/
├── SPEC.md                  ← Spec hiện hành
├── api-contract.md          ← đã align với Spec
├── slot-map.md              ← đã align với Spec
├── state-map.md             ← đã align với Spec
├── backend-contract-map.md  ← đã align với Spec
└── alignment-notes.md       ← conflicts đã resolve + [TBD] còn lại
```

**AI agent note khi giao:**
- Liệt kê rõ những [TBD] cần người duyệt quyết định
- Liệt kê rủi ro hoặc điểm unclear nếu có
- Đề xuất thứ tự implement modules (từ backend-contract-map dependency graph)

**Sau khi giao → AI agent dừng hoàn toàn.**
- Không tự proceed
- Không diễn giải "silence = approval"
- Nếu reject → quay đúng bước được chỉ định, không rewrite toàn bộ

---

## Bước 11 — Plan implement từng module

**⛔ CẤM plan nhiều module cùng lúc — dù owner có yêu cầu.**
Nếu bị push → giải thích: plan hàng loạt dẫn đến dependency hell và QA không có boundary.

**Chỉ bắt đầu khi có approval rõ ràng từ Bước 9.**

**Thứ tự mặc định:** Auth trước → module nhiều FE dependency nhất → còn lại theo priority.

**Per-module plan template:**

```markdown
## Implementation Plan: Module [tên]

### Scope (từ backend-contract-map.md)
Endpoints: [list]
DB tables: [list]
Business rules: [list — trích từ Spec §X]

### Tasks
1. DB migration
2. Model / schema
3. Service layer (business rules)
4. Controller / handler (endpoints)
5. Middleware (auth, validation)
6. Unit tests
7. Integration tests
8. Wire FE: thay mock bằng real API

### Definition of Done
- [ ] Endpoints đúng path/method/shape theo api-contract.md
- [ ] Response khớp TypeScript types của FE
- [ ] Business rules đúng theo Spec
- [ ] Unit + integration tests pass
- [ ] FE render đúng trên UI thật với real API
- [ ] Bước 11 QA pass
```

---

## Bước 12 — Runtime QA sau mỗi module

**Không pass QA = không sang module tiếp.**

```markdown
## Runtime QA: Module [tên]

### Happy path
- [ ] Endpoints trả đúng data → UI render đúng slot (verify với slot-map.md)
- [ ] Actions trigger đúng state transitions (verify với state-map.md)

### Error handling
- [ ] Network error → error banner (không crash)
- [ ] Validation error → field-level message
- [ ] 401 → redirect login
- [ ] 403 → appropriate message
- [ ] 404 → empty state

### Edge cases
- [ ] Empty list → empty state slot
- [ ] Pagination boundary (1 item, max items)
- [ ] Concurrent actions không có race condition

### Performance
- [ ] Loading state visible trong lúc fetch
- [ ] Không flash of empty content

### Sign-off
- [ ] Pass → log vào docs/qa-log.md → báo owner → nhận approval → Bước 10 module tiếp
- [ ] Fail → ghi bug, fix, re-run — không proceed
```

---

## Ghi chú cho AI agent

1. **Bước 4-7 là extraction, không phải sáng tạo** — đọc FE code, không đoán.
   Thiếu thông tin → báo rõ cái gì thiếu, hỏi — không tự điền.

2. **Bước 8: Spec phán xét artifacts, không phải ngược lại.**
   Artifact mâu thuẫn với Spec → sửa artifact. Không sửa Spec ở bước này.

3. **Bước 9 là gate tuyệt đối.** Không có approval = không có Bước 10.
   Silence không phải approval.

4. **Bước 10: 1 module = 1 plan.** Cấm tuyệt đối plan 2+ module cùng lúc.

5. **Spec là luật tối cao.** UI/FE mâu thuẫn Spec → dừng, escalate,
   không tự quyết định bên nào thắng.

6. **BE không implement gì ngoài backend-contract-map.** Không tự thêm endpoint.

7. **Test trên UI thật là bắt buộc.** "Done" chưa hợp lệ nếu chưa test real API trên FE thật.