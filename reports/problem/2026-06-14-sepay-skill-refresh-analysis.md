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

## Recommended Skill Update Shape

### `overview.md`

Restructure around product surfaces:

- Webhooks: real-time bank transaction events
- API v2: lookup, reconciliation, account/VA/order management
- Payment Gateway: hosted checkout, cards/NAPAS/QR, IPN
- QR utility: static/dynamic VietQR image generation
- Order VAs: bank-specific VA order flow
- Bank Hub / OAuth2 / eInvoice / SoundBox: optional, load only for specific use cases

Also update rate limits, Test Mode, and surface-selection rules.

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

Review Laravel package accuracy before presenting it as official.

### `best-practices.md`

Replace project-specific examples with more general SePay patterns:

- choose integration surface first
- HMAC first for webhook security
- idempotency by transaction ID / IPN event
- enqueue heavy work after fast 200
- API v2 reconciliation backup
- exact amount/content/VA matching
- bank memo transformations
- Test Mode/Live token separation
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

1. `webhooks.md` and `scripts/sepay-webhook-verify.js`: highest production-risk area because webhook spoofing or duplicate processing directly affects money movement.
2. `api.md`: API v2 should become the default for new reconciliation/order-VA work.
3. `overview.md`: prevent agents from choosing the wrong SePay surface.
4. `qr-codes.md`: bank-specific QR/VA rules affect whether transactions are detected correctly.
5. `sdk.md`: align install commands and Payment Gateway base URLs.
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
- VietQR generation: https://developer.sepay.vn/en/tien-ich-khac/tao-qr-code
- PHP SDK package: https://packagist.org/packages/sepay/sepay-pg

