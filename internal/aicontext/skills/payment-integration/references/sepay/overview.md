# SePay Overview

SePay is a Vietnamese payment automation platform. Treat it as several related product surfaces, not one generic API.

Official docs checked: 2026-06-14.

## Choose The Surface First

| Need | Use | Load Next |
|------|-----|-----------|
| Proactive transaction lookup, reconciliation, bank-account lookup, VA lookup, or Order VA management | SePay API v2 | `api.md` |
| Real-time bank-account transaction notifications | SePay Webhooks | `webhooks.md` |
| Hosted checkout with QR, cards, NAPAS, gateway order lifecycle, or gateway IPN | SePay Payment Gateway | `payment-gateway.md`, then `sdk.md` if using an SDK |
| Render VietQR images for transfer instructions | QR utility | `qr-codes.md` |
| Give each order a bank-specific virtual account | Order VAs through API v2 | `api.md`, then `qr-codes.md` for display rules |
| Bank Hub, OAuth2 partner flows, eInvoice, SoundBox, or other optional SePay products | Product-specific SePay docs | Load only when the task names that product |

Do not mix bank-account Webhooks with Payment Gateway IPN:

- Webhooks notify bank-account transactions and are paired with API v2 reconciliation.
- Payment Gateway IPN notifies hosted checkout order and transaction results.
- API v2 is for server-side lookup and management, not a real-time delivery channel.

## API v2 Summary

Use API v2 by default for new server-side SePay API work.

| Environment | Base URL | Token |
|-------------|----------|-------|
| Production | `https://userapi.sepay.vn/v2` | Live API token |
| Sandbox/Test mode | `https://userapi-sandbox.sepay.vn/v2` | Test mode API token |

Rules:

- Authenticate with `Authorization: Bearer <API_TOKEN>`.
- Live and Sandbox tokens are not interchangeable.
- API v2 uses UUID identifiers, integer money fields, `data` / `meta` response envelopes, `page` / `per_page` pagination, and normal HTTP status codes.
- Rate limit is 3 requests per second per IP. Handle HTTP 429 with `Retry-After`, `X-RateLimit-Limit`, `X-RateLimit-Remaining`, and `X-RateLimit-Reset`.
- Legacy `https://my.sepay.vn/userapi/*` is migration-only. Do not start new work there.
- API v2 Order VAs are Production-only and bank-specific.

## Webhooks Summary

Use Webhooks for real-time bank-account transaction events.

Production guidance:

- Prefer HMAC-SHA256 authentication.
- Keep API Key or OAuth2 only when that is the chosen security model.
- Use no authentication only for quick tests.
- Verify HMAC over the raw request body with `X-SePay-Signature`, `X-SePay-Timestamp`, signed string `{timestamp}.{raw_body}`, a short replay window, and timing-safe comparison.
- Return HTTP 200 or 201 with `{"success": true}` quickly, then process heavy work asynchronously.
- Automatic retry is initial delivery plus seven retries on Fibonacci spacing, normally ending around 33 minutes.
- Use SePay payment-code recognition and the webhook `code` field when configured.
- Make webhook handlers idempotent because auto-retry and manual replay can deliver the same transaction again.
- IP allowlisting is defense in depth. Link to SePay's current IP page instead of copying permanent IP values into reusable code.

## Payment Gateway Summary

Use Payment Gateway for hosted checkout, gateway payment methods, and gateway IPN.

| Environment | Base URL |
|-------------|----------|
| Production | `https://pgapi.sepay.vn` |
| Sandbox | `https://pgapi-sandbox.sepay.vn` |

Rules:

- Payment Gateway API uses Basic Auth: `merchant_id:secret_key`.
- Checkout form submission uses `/v1/checkout/init`.
- Gateway IPN URL must be HTTPS and return HTTP 200 to acknowledge receipt.
- If merchant auth type is `SECRET_KEY`, verify the `X-Secret-Key` header.
- IPN payloads contain `timestamp`, `notification_type`, `order`, `transaction`, and `customer`.
- Do not apply bank-webhook HMAC assumptions to Payment Gateway IPN unless that specific gateway surface documents them.

## QR Utility Summary

Use `https://qr.sepay.vn/img` only to render VietQR images. QR generation does not confirm payment by itself.

Required parameters:

- `acc`
- `bank`

Common optional parameters:

- `amount`
- `des`
- `template`
- `download`
- `showinfo`
- `fullacc`
- `holder`
- `store`

Bank and VA rules matter:

- `bank` can use SePay-supported short name, alias, code, or BIN.
- Some banks require official VAs for automatic matching.
- Memo-based VAs require `TKP` plus VA code in `des`.
- VietinBank personal and household-business accounts require `SEVQR` in the memo.

## Test Mode Vs Live

- Test mode uses isolated API tokens, accounts, VAs, webhooks, and simulated transactions.
- Test mode quotas: 500 simulated transactions per day reset at 00:00 Vietnam time, 50 bank accounts, 100 VAs per bank account, 50 webhooks, and 50 API Access tokens.
- Test mode webhooks accept HTTP and self-signed certificates. Live requires valid HTTPS.
- Test mode webhook creation has no Live alert step and only allows Test mode accounts.
- Passing in Test mode does not prove every Production-only feature is available; Order VAs are Production-only.

## Production Safety Rules

1. Choose the SePay surface before writing code.
2. Keep Live and Test credentials separate.
3. Never expose API v2 bearer tokens or gateway secret keys to client code.
4. Use HMAC-first webhook verification for bank transactions.
5. Use `code` from SePay payment-code recognition when available.
6. Always store an idempotency key before applying payment effects.
7. Reconcile via API v2 after webhook/IPN misses, outages, retries, or manual replay.
8. Treat IP allowlisting as an extra control, not as authentication.
9. Keep gateway IPN and bank-account webhooks in separate handlers or explicit branches.

## Official Source Pointers

- API v2: https://developer.sepay.vn/en/sepay-api/v2/gioi-thieu
- Webhooks: https://developer.sepay.vn/en/sepay-webhooks/xac-thuc
- Payment Gateway API: https://developer.sepay.vn/en/cong-thanh-toan/API/tong-quan
- Payment Gateway IPN: https://developer.sepay.vn/en/cong-thanh-toan/IPN
- QR generation: https://developer.sepay.vn/en/tien-ich-khac/tao-qr-code
- Test mode quotas: https://developer.sepay.vn/en/tien-ich-khac/test-mode/han-muc
- Current SePay outbound IPs: https://developer.sepay.vn/en/dia-chi-ip
