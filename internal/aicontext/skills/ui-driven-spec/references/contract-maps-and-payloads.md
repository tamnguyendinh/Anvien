# Contract Maps And API Payloads

> This file is part of UI-Driven Spec Workflow. Read when: generating Data Source Map, API Payload Contracts, UI Slot Map, State Map, or Backend Contract Map.

---

## 1. Data Source Map

**Purpose:** Describe where production data should come from after converting the static HTML prototype to Electron React + Go local backend or another target stack.

For endpoint, payload, and response shape expected by the frontend, see `backend-source-map.md` (or feature API payload docs).

---

## 2. API Payload Contract File

**Purpose:** Record what FE expects from each endpoint.
This is the target BE must satisfy — not what BE designs in isolation.
**Function:** Document detailed payload for each button / form / input by feature.

**Output file path:** `DOCS/SPEC/IMPLEMENTATION-MAPS/API-payload/<feature>-API-payload.md`

**Examples:**

| Feature | File |
|---|---|
| Orders page | `orders-API-payload.md` |
| Tables/POS runtime | `tables-pos-API-payload.md` |
| Settings | `settings-API-payload.md` |

### Payload Doc Format

Each action should document:

| Section | Meaning |
|---|---|
| UI trigger | Button, row action, dialog submit, dropdown selection |
| Source fields | Fields read from UI state or user input |
| Backend surface | Local Go endpoint / service command |
| Required validation | FE affordance and Go service enforcement |
| Result projection | Which UI projection must refetch / update |
| Sync/audit rule | Event / outbox / ledger behavior when relevant |

### Hard Rules For API Payloads:
- Do not put payload details back into `backend-contract-map.md`, `data-source-map.md`, `state-map.md`, or `ui-slot-map.md`.
- Use `snake_case` for documented contract fields.
- FE must not send `owner_id`.
- FE disabled states are UX only. Go service still enforces permissions, active shift, locks, and hash-chain health.
- Prefer field tables over large sample JSON blocks. Add exact JSON only when it is necessary for a migration/test fixture.

### Sample API Contract Format:

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

[List all endpoints FE requires]
```

> **Note:** Extract directly from `types/` and `hooks/` — do not guess.

---

## 3. UI Slot Map

**Purpose:** Record what renders where and its display conditions.

**Output file:** `docs/slot-map.md`

```markdown
## UI Slot Map

### Screen: OrderDashboard

| Slot | Component | Display Condition | Data Source |
|------|-----------|-------------------|-------------|
| header | PageHeader | always | static |
| stats-row | StatsCard × 3 | role === 'admin' | GET /stats |
| order-list | OrderTable | orders.length > 0 | GET /orders |
| empty-state | EmptyOrders | orders.length === 0 | — |
| error-banner | ErrorBanner | fetchError !== null | — |
| pagination | Pagination | total > pageSize | from response |

### Conditional Renders
- `CreateOrderButton`: visible if user.permissions.includes('order:create')
- `CancelButton` per row: visible if order.status === 'pending'
```

---

## 4. State Map

**Purpose:** Record data flow, loading / error / empty states, and state transitions.

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
| Action | Trigger | State Change | Side Effect |
|--------|---------|--------------|-------------|
| Click "Cancel Order" | Button click | order.status = 'cancelling' | POST /orders/:id/cancel |
| Cancel success | API response | remove from list | toast success |
| Cancel fail | API error | revert status | toast error |
```

---

## 5. Backend Contract Map

**Purpose:** Aggregate endpoint + payload + response shape that FE needs, grouped by BE module. Single source of truth for BE implementation.

**Output file:** `docs/backend-contract-map.md`

```markdown
## Backend Contract Map

> Do not implement anything that is not defined in this document.

### Module: Auth
Endpoints: POST /auth/login, POST /auth/logout, POST /auth/refresh
DB tables: users, sessions
Business rules:
  - Login fail 5 times → lock 15 minutes
  - Token: access 15m, refresh 7d
FE triggers: LoginForm submit, auto-refresh when token expires soon
UI slots affected: header (user avatar), all protected routes

### Module: Orders
Endpoints: GET /orders, POST /orders, GET /orders/:id, POST /orders/:id/cancel
DB tables: orders, order_items, order_status_history
Business rules:
  - Can only cancel if status === 'pending'
  - Cancel writes to order_status_history
FE triggers: OrderList load, CreateOrderForm submit, CancelButton click
UI slots affected: OrderTable, StatsCard, EmptyOrders

[Repeat for each module]

### Module Dependency Graph
Auth → prerequisite for all modules
Orders → Products (price lookup), Inventory (stock check)
```
