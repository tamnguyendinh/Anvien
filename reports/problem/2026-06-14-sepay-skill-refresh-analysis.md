# SePay Skill Refresh Analysis

Date: 2026-06-14
Scope: `internal/aicontext/skills/payment-integration/references/sepay`
Status: discussion report only; no skill/code changes made in this report.

## Problem Summary

The current SePay references in the `payment-integration` skill are partially stale and mix several SePay product surfaces into one mental model:

- SePay API v1 / legacy `userapi`
- SePay API v2
- SePay Webhooks for bank-account transaction events
- SePay Payment Gateway checkout and IPN
- VietQR image generation
- Order Virtual Accounts
- Bank Hub / OAuth2 / optional products

This can make agents implement the wrong integration path: for example, using legacy API endpoints for a new reconciliation integration, treating bank transaction webhooks as Payment Gateway IPN, or verifying webhooks with API key only when HMAC-SHA256 is the stronger current production option.

## Local Files Reviewed

- `references/sepay/overview.md`
- `references/sepay/api.md`
- `references/sepay/webhooks.md`
- `references/sepay/sdk.md`
- `references/sepay/qr-codes.md`
- `references/sepay/best-practices.md`
- `references/implementation-workflows.md`
- `scripts/sepay-webhook-verify.js`

## Current-State Findings

### 1. SePay API v2 should be the default for new API integrations

Current local docs mostly describe legacy `https://my.sepay.vn/userapi/` endpoints.

Current official docs describe API v2 as the REST API for proactive transaction lookup, bank-account lookup, virtual-account lookup, and order VA management. API v2 uses:

- Production base URL: `https://userapi.sepay.vn/v2`
- Sandbox base URL: `https://userapi-sandbox.sepay.vn/v2`
- Bearer token auth
- UUID identifiers
- integer money fields
- `page` / `per_page` pagination
- proper HTTP status codes
- unified `data` / `meta` response shape

Legacy `userapi/*` remains available, but should be documented as legacy or migration-only, not the default for new work.

### 2. Rate-limit guidance is stale

Local docs say 2 calls/second.

Current official API v2 docs say maximum 3 requests per second per IP, returning HTTP 429 with `Retry-After` plus `X-RateLimit-Limit`, `X-RateLimit-Remaining`, and `X-RateLimit-Reset` headers.

The skill should update rate-limit handling and avoid hard-coding v1-only retry header behavior for all SePay APIs.

### 3. Webhook authentication should prioritize HMAC-SHA256

Local docs and `scripts/sepay-webhook-verify.js` focus on `No_Authen`, API Key, and OAuth2.

Current official webhook docs list four methods:

- None: testing only
- API Key: medium security
- HMAC-SHA256: recommended, detects payload tampering
- OAuth2: high security when the receiver already has an OAuth server

The skill should add HMAC-SHA256 verification as the production default or recommended option. Verification needs:

- raw request body, not parsed/re-serialized JSON
- `X-SePay-Signature` in `sha256={hex_hash}` format
- `X-SePay-Timestamp`
- signed string format `{timestamp}.{raw_body}`
- replay window, e.g. reject timestamps older than 5 minutes
- constant-time comparison

The helper script should eventually support this path, because current API-key-only verification can teach weaker production behavior.

### 4. Webhook response and retry behavior need correction

Local docs say respond within 5 seconds and retries continue over about 5 hours.

Current official webhook error docs say:

- Success is HTTP 200/201 with body `{"success": true}`.
- A timeout is more than 30 seconds with no response.
- Auto retry is 8 total attempts: initial delivery plus 7 retries.
- Retry spacing follows Fibonacci intervals and normally ends around 33 minutes.
- Queue items older than 5 hours are not picked up by cron, but the practical retry cap is about 33 minutes.

The skill should keep the fast-response recommendation, but make it precise: return 200 quickly, enqueue heavy work, and reconcile if missed beyond retry window.

### 5. Webhooks and Payment Gateway IPN must be separated

Local content can lead agents to treat SePay Webhooks and Payment Gateway IPN as the same surface.

They are different:

- SePay Webhooks notify bank-account transaction events.
- Payment Gateway IPN notifies hosted checkout/payment gateway order result events.

The skill should make agents choose the correct surface first:

- Need real-time bank transfer detection or reconciliation? Use Webhooks plus API v2.
- Need hosted checkout with QR/cards/NAPAS? Use Payment Gateway checkout and IPN.
- Need proactive lookup/backfill? Use API v2.

### 6. Payment Gateway API endpoint/auth are stale in local docs

Local docs mention old gateway endpoints such as `sandbox.pay.sepay.vn/v1/init` and `pay.sepay.vn/v1/init`.

Current official Payment Gateway API docs show:

- Production base URL: `https://pgapi.sepay.vn`
- Sandbox base URL: `https://pgapi-sandbox.sepay.vn`
- Basic Authentication: `merchant_id:secret_key`
- Create payment order endpoint under Payment Gateway API, including `/v1/checkout/init`

The SDK and overview references should align with current Payment Gateway base URLs and auth.

### 7. Test Mode should become explicit

Current skill references sandbox in several places but does not clearly separate Live/Test behavior across API v2, webhooks, QR, and gateway flows.

Current official docs indicate separate environments/tokens for API v2. API v2 Order VAs are Production-only. Test Mode supports development/simulation flows but does not imply every feature is available in sandbox.

The skill should add a "Test Mode vs Live" section to avoid agents using Live tokens on sandbox endpoints or assuming Order VAs can be tested in sandbox.

### 8. QR generation docs are incomplete

Local QR docs cover:

- `acc`
- `bank`
- `amount`
- `des`
- `template`
- `download`

Current official QR docs also include:

- `showinfo`
- `fullacc`
- `holder`
- `store`
- `template=standee`
- `bank` may use `short_name`, `alias`, `code`, or `bin`
- bank-specific VA / memo rules

Important bank rules missing or under-emphasized locally:

- Some banks require VA for automatic matching.
- VietinBank personal and household-business accounts require `SEVQR` in memo.
- Memo-based VA requires `TKP` plus VA code in `des`.
- QR `acc` and `des` depend on VA type.

### 9. Order VA support is broader and bank-specific

Local API docs mention BIDV and "others".

Current API v2 docs say Order VAs support:

- BIDV enterprise
- Sacombank personal / household business
- Vietcombank enterprise / household business

Bank-specific request rules matter:

- Vietcombank requires raw `tid`, fetched from terminals.
- Sacombank requires `va_prefix`.
- Sacombank and Vietcombank require exact amount and do not support partial payment.
- `order_code` length differs by bank.
- BIDV has optional amount and optional custom VA holder name only when enabled.

The skill should not present VA creation as one generic flow.

### 10. SDK guidance should be refreshed

Local Node SDK install uses:

```bash
npm install github:sepay/sepay-pg-node
```

Current official docs use npm package installation:

```bash
npm i sepay-pg-node
```

PHP SDK `composer require sepay/sepay-pg` remains aligned with official Payment Gateway SDK docs.

Laravel package guidance should be reviewed and possibly labeled as Laravel/webhook/community package unless it is confirmed as the official Payment Gateway SDK path.

### 11. Payment code recognition should be first-class in webhook guidance

Official webhook docs now describe configurable payment-code recognition instead of leaving each integration to parse transfer content ad hoc.

The configuration includes:

- company-level enable/disable for payment-code recognition
- one or more code templates
- prefix length rules
- min/max suffix length rules
- suffix character type: numeric or alphanumeric
- active/inactive template status
- matching order across active templates

When a transaction description matches a configured template, SePay attaches the matched value to the webhook payload `code` field. Webhook filters such as "only send when a payment code is present" and "filter by payment code" depend on this field.

The skill should teach agents to prefer SePay's `code` field when available, and only fall back to custom memo parsing when the integration deliberately does not use SePay payment-code recognition.

### 12. Webhook operations need monitoring, replay, and incident coverage

Official webhook docs include operational controls that are missing from the local skill:

- delivery logs with request/response details
- dashboard metrics such as success rate, failures, timeouts, average response, and P95 response
- alert channels for Telegram, Slack, and Discord
- incident records for repeated failures
- manual replay for failed or successful deliveries
- replay limits and replay labels

Manual replay can deliver the same payload again, so the skill should make idempotency and deduplication mandatory for replay-safe endpoints. This is separate from automatic retries: replay is an operator action and must not create duplicate order/payment effects.

### 13. IP allowlisting should reference SePay's official current IP page

Official docs list public SePay outbound IPs for Webhooks, Payment Gateway IPN, Bank Hub IPN, and other callbacks. They also warn that the list may be updated.

The skill should:

- link to the official IP address page instead of treating any copied list as permanent
- explain that allowlisting is optional infrastructure hardening, not a replacement for webhook/IPN authentication
- mention that HTTPS is required for production callback URLs
- avoid baking current IP values into reusable skill guidance unless the source and access date are recorded

### 14. Test Mode needs quotas and live-difference details

Current local guidance says to separate sandbox/live, but it misses concrete Test Mode constraints.

Official Test Mode quota docs list hard per-company limits:

- 500 simulated transactions per day, reset at 00:00 Vietnam time
- 50 bank accounts
- 100 virtual accounts per bank account
- 50 webhooks
- 50 API Access tokens

Official Test Mode webhook docs also say Test Mode accepts HTTP and self-signed certificates, while Live requires a valid HTTPS certificate. Test Mode webhook creation omits the Live alert step and only allows Test Mode accounts.

The skill should make these differences explicit so agents do not overinterpret a Test Mode success as production readiness.

### 15. Payment Gateway IPN contract needs explicit handling

Payment Gateway IPN is not the same as SePay bank-transaction Webhooks.

Current official IPN docs show:

- merchant-configured IPN URL under Payment Gateway configuration
- HTTPS URL requirement
- endpoint must return HTTP 200 to acknowledge receipt
- optional `X-Secret-Key` header when merchant auth type is `SECRET_KEY`
- JSON payload with `timestamp`, `notification_type`, `order`, `transaction`, and `customer`
- notification types including `ORDER_PAID` and `TRANSACTION_VOID`

The skill should add a clear IPN contract section for hosted checkout integrations, including idempotency by order/invoice/transaction identifiers and a warning not to apply bank-webhook HMAC assumptions to Payment Gateway IPN unless that surface explicitly supports them.

### 16. Payment Gateway SDK and order lifecycle rules need coverage

Current SDK guidance only updates the package install command. The official Node SDK docs also include operational details that affect implementation correctness:

- supported `payment_method` values: `CARD`, `BANK_TRANSFER`, `NAPAS_BANK_TRANSFER`
- the SDK builds one-time payment form fields and checkout URL
- custom HTML form/signature generation must preserve SePay's required field order
- SDK methods cover order list, order detail, card transaction void, and QR order cancel
- order-detail responses include order status, authentication status, and transaction history
- documented order statuses include `CAPTURED`, `CANCELLED`, and `AUTHENTICATION_NOT_NEEDED`
- cancel order applies to `BANK_TRANSFER` or `NAPAS_BANK_TRANSFER` orders that are not already captured/canceled

The skill should distinguish gateway order lifecycle behavior from API v2 transaction lookup and from bank-account webhook reconciliation.

## Recommended Skill Update Shape

### `overview.md`

Restructure around product surfaces:

- Webhooks: real-time bank transaction events
- API v2: lookup, reconciliation, account/VA/order management
- Payment Gateway: hosted checkout, cards/NAPAS/QR, IPN
- QR utility: static/dynamic VietQR image generation
- Order VAs: bank-specific VA order flow
- Bank Hub / OAuth2 / eInvoice / SoundBox: optional, load only for specific use cases

Also update rate limits, Test Mode quotas/live differences, IP allowlisting, payment-code recognition, and surface-selection rules.

### `api.md`

Make API v2 the default:

- base URLs
- auth
- response envelope
- pagination
- status/error codes
- transaction/account/VA/order endpoints
- migration note from v1

Keep v1 as "legacy API" with a warning and mapping table.

### `webhooks.md`

Update with:

- HMAC-SHA256 as recommended production auth
- API Key and OAuth2 as supported alternatives
- raw-body verification
- replay protection
- valid response contract
- retry schedule
- diagnostics
- monitoring/replay/reconciliation guidance
- payment-code recognition and the webhook `code` field
- delivery logs, dashboard metrics, alerts, incidents, and manual replay
- replay-safe idempotency/deduplication rules
- IP allowlist guidance pointing to SePay's official current IP page

Also explicitly separate Webhooks from Payment Gateway IPN.

### `qr-codes.md`

Add:

- `showinfo`, `fullacc`, `holder`, `store`, `standee`
- supported bank identifier types
- bank-specific VA and memo rules
- examples for official VA, memo-based VA, no VA, VietinBank `SEVQR`

### `sdk.md`

Update:

- Node install command
- Payment Gateway base URLs
- Basic Auth
- checkout init endpoint
- current Node/PHP SDK split
- supported gateway `payment_method` values
- one-time payment form/signature field-order warning
- order list/detail, card transaction void, and QR order cancel APIs
- order status/authentication status handling

Review Laravel package accuracy before presenting it as official.

### Payment Gateway IPN reference

Either add a dedicated `payment-gateway.md` / `ipn.md` reference or add a clearly separated IPN section to the current SePay references.

Cover:

- HTTPS IPN URL requirement
- HTTP 200 acknowledgement
- `X-Secret-Key` validation when configured
- payload shape and notification types
- order/transaction/customer mapping
- idempotency and reconciliation after IPN failures

### `best-practices.md`

Replace project-specific examples with more general SePay patterns:

- choose integration surface first
- HMAC first for webhook security
- use SePay payment-code recognition and `code` field when available
- idempotency by transaction ID / IPN event
- replay-safe processing for manual webhook replay
- enqueue heavy work after fast 200
- API v2 reconciliation backup
- exact amount/content/VA matching
- bank memo transformations
- Test Mode/Live token separation
- Test Mode quota and Live HTTPS differences
- IP allowlisting as defense-in-depth
- underpayment/overpayment policy per product

### `scripts/sepay-webhook-verify.js`

Future code update should support:

- HMAC-SHA256 verification
- raw body input
- timestamp replay window
- timing-safe comparison
- API Key fallback
- clear errors for missing signature/timestamp

## Suggested Update Priority

1. `webhooks.md` and `scripts/sepay-webhook-verify.js`: highest production-risk area because webhook spoofing, weak payment-code matching, replay duplication, or missed failures directly affects money movement.
2. `api.md`: API v2 should become the default for new reconciliation/order-VA work.
3. Payment Gateway IPN / SDK coverage: hosted checkout integrations need a separate IPN contract and gateway order lifecycle rules.
4. `overview.md`: prevent agents from choosing the wrong SePay surface.
5. `qr-codes.md`: bank-specific QR/VA rules affect whether transactions are detected correctly.
6. `best-practices.md`: generalize and de-stale implementation patterns.

## Sources Checked

- SePay API v2 overview: https://developer.sepay.vn/en/sepay-api/v2/gioi-thieu
- Upgrade from v1: https://developer.sepay.vn/en/sepay-api/v2/nang-cap-tu-v1
- API v2 quick start / rate limits: https://developer.sepay.vn/en/sepay-api/v2/bat-dau-nhanh
- Create Order VA API: https://developer.sepay.vn/en/sepay-api/v2/don-hang/tao-don-hang
- Webhook authentication: https://developer.sepay.vn/en/sepay-webhooks/xac-thuc
- Webhook error handling / retries: https://developer.sepay.vn/en/sepay-webhooks/xu-ly-loi
- Payment Gateway overview: https://developer.sepay.vn/en/cong-thanh-toan/gioi-thieu
- Payment Gateway API overview: https://developer.sepay.vn/en/cong-thanh-toan/API/tong-quan
- Create payment order API: https://developer.sepay.vn/en/cong-thanh-toan/API/don-hang/form-thanh-toan
- Payment Gateway IPN: https://developer.sepay.vn/en/cong-thanh-toan/IPN
- Payment Gateway Node.js SDK: https://developer.sepay.vn/en/cong-thanh-toan/sdk/nodejs
- Payment Gateway order details: https://developer.sepay.vn/en/cong-thanh-toan/API/don-hang/chi-tiet-don-hang
- Payment Gateway cancel order: https://developer.sepay.vn/en/cong-thanh-toan/API/don-hang/huy-don-hang
- VietQR generation: https://developer.sepay.vn/en/tien-ich-khac/tao-qr-code
- Webhook payment code structure: https://developer.sepay.vn/en/sepay-webhooks/cau-hinh-ma-thanh-toan
- Webhook monitoring: https://developer.sepay.vn/en/sepay-webhooks/giam-sat
- SePay public IP addresses: https://developer.sepay.vn/en/dia-chi-ip
- Test Mode quotas: https://developer.sepay.vn/en/tien-ich-khac/test-mode/han-muc
- Test Mode webhook creation: https://developer.sepay.vn/en/tien-ich-khac/test-mode/tao-webhook
- PHP SDK package: https://packagist.org/packages/sepay/sepay-pg
