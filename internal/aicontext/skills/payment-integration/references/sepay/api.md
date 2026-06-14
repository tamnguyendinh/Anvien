# SePay API v2 Reference

Use API v2 by default for new SePay API integrations. The old `userapi/*` endpoints are legacy/migration-only.

Official docs checked: 2026-06-14.

## Base URLs

| Environment | Base URL | Notes |
|-------------|----------|-------|
| Production | `https://userapi.sepay.vn/v2` | Live bank data and Live API tokens |
| Sandbox/Test mode | `https://userapi-sandbox.sepay.vn/v2` | Isolated Test mode data and Test mode API tokens |

Order VA APIs are Production-only.

## Authentication

```http
Authorization: Bearer <API_TOKEN>
Content-Type: application/json
```

Rules:

- Create Live tokens from API Access while in Live mode.
- Create Test tokens from API Access while in Test mode.
- Do not reuse tokens across environments.
- Keep tokens server-side only.

## Rate Limit

- Maximum: 3 requests per second per IP.
- On excess: HTTP 429.
- Retry based on `Retry-After`.
- Track `X-RateLimit-Limit`, `X-RateLimit-Remaining`, and `X-RateLimit-Reset`.

```javascript
if (response.status === 429) {
  const retryAfter = Number(response.headers.get('Retry-After') || '1');
  await sleep(retryAfter * 1000);
}
```

## Response Shape

List endpoints use a `data` array and `meta.pagination` object.

```json
{
  "status": "success",
  "data": [],
  "meta": {
    "pagination": {
      "total": 0,
      "per_page": 20,
      "current_page": 1,
      "last_page": 1,
      "has_more": false
    }
  }
}
```

Detail endpoints return `data` as an object. Identifiers are UUIDs. Money fields are integer VND unless a specific endpoint documents another format.

## Main API Groups

### Transactions

Use transactions for lookup, backfill, customer-facing history, admin dashboards, and reconciliation.

Common paths:

```http
GET /transactions
GET /transactions/{transaction_uuid}
```

Common filters:

- bank account
- date range
- amount range
- transaction content search
- transfer type
- `since_id` polling cursor

Webhook payloads and transaction API responses are related but not interchangeable. Webhooks are real-time delivery; API v2 is the source to query when a delivery is missed, delayed, or ambiguous.

### Bank Accounts

Use bank-account endpoints to list and inspect linked SePay bank accounts.

```http
GET /bank-accounts
GET /bank-accounts/{bank_account_uuid}
```

Cache account lists carefully and refresh when SePay dashboard configuration changes.

### Virtual Accounts

Use VA endpoints to look up existing virtual accounts across supported banks.

```http
GET /virtual-accounts
GET /virtual-accounts/{va_uuid}
```

Do not assume every bank supports the same VA type or memo rule. Coordinate display instructions with `qr-codes.md`.

## Order VAs

Use Order VAs when each order should receive a bank-specific VA or bank-specific VA/memo rule instead of a generic transfer memo.

Supported official Order VA families:

| Bank / Account Type | Key Rules |
|---------------------|-----------|
| BIDV enterprise | Supports Order VA flow; amount can be optional where the account capability allows it; custom VA holder name only when enabled. |
| Sacombank personal / household business | Requires `va_prefix`; exact amount required; partial payment is not supported. |
| Vietcombank enterprise / household business | Requires raw `tid` from terminals; exact amount required; partial payment is not supported. |

Implementation rules:

- Treat bank-specific request fields as required for that bank even when another bank does not need them.
- Validate `order_code` length and character constraints against the selected bank's endpoint docs.
- Keep amount policy explicit: exact-only, optional, overpayment, and partial-payment behavior vary by bank/account product.
- Show generated VA/QR instructions only after the API response confirms the order/VA state.
- Use API v2 reconciliation and bank Webhooks to confirm actual incoming transactions.

## Error Handling

Use HTTP status codes, not the legacy v1 `status` field, as the first failure signal.

| Status | Meaning |
|--------|---------|
| 200 | Success |
| 201 | Created |
| 400 | Bad request |
| 401 | Missing or invalid token |
| 403 | Forbidden |
| 404 | Resource not found |
| 422 | Validation failed |
| 429 | Rate limited |
| 500 / 503 | SePay or upstream service issue |

Log:

- request path and environment
- correlation/order ID
- status code
- SePay error body
- retry headers when present

## Legacy v1 Migration Notes

Legacy v1 endpoints such as `https://my.sepay.vn/userapi/transactions/list` may exist in older integrations. Do not use them as the default for new work.

| Legacy Habit | Current Guidance |
|--------------|------------------|
| Numeric transaction IDs | Use API v2 UUID identifiers. |
| `limit` / `since_id` v1 pagination | Use API v2 `page` / `per_page` and documented filters. |
| 2 calls/second guidance | Use API v2 3 requests/second per IP and standard rate-limit headers. |
| `x-sepay-userapi-retry-after` only | Prefer standard `Retry-After` for API v2. |
| v1 response fields as canonical | Map to API v2 `data` / `meta` shape and current field names. |

When migrating, keep the old endpoint in a clearly isolated adapter until all callers are moved, then remove it. Do not mix v1 and v2 response shapes in the same downstream contract.

## Cross-Surface Links

- Use `webhooks.md` for real-time bank-account transaction delivery.
- Use `payment-gateway.md` for hosted checkout and Payment Gateway IPN.
- Use `qr-codes.md` for QR display parameters and bank-specific memo/VA display rules.
- Use `best-practices.md` for idempotency, reconciliation, and production rollout.
