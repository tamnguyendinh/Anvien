# SePay Best Practices

Repo-agnostic production patterns for SePay integrations. Choose the SePay surface first, then apply the matching security, idempotency, reconciliation, and rollout rules.

Official docs checked: 2026-06-14.

## Surface Selection

| Scenario | Recommended Surface |
|----------|---------------------|
| Need real-time direct bank-transfer detection | Bank-account Webhooks plus API v2 reconciliation |
| Need hosted checkout, cards, NAPAS, gateway QR, gateway order lifecycle | Payment Gateway plus IPN |
| Need periodic transaction lookup, admin history, backfill, or Order VA management | API v2 |
| Need to render a VietQR image for manual transfer instructions | QR utility, then confirm through Webhooks/API v2 |
| Need one VA per order with bank constraints | API v2 Order VAs |

Avoid hybrid flows that treat gateway IPN, bank Webhooks, API v2, and QR image generation as one interchangeable API.

## Credentials And Environment

Use separate environment variables per surface and environment:

```bash
# API v2
SEPAY_API_BASE_URL=https://userapi.sepay.vn/v2
SEPAY_API_TOKEN=...

# Bank-account webhook
SEPAY_WEBHOOK_SECRET=...
SEPAY_WEBHOOK_API_KEY=...

# Payment Gateway
SEPAY_PG_BASE_URL=https://pgapi.sepay.vn
SEPAY_PG_MERCHANT_ID=...
SEPAY_PG_SECRET_KEY=...
SEPAY_PG_IPN_SECRET_KEY=...

# QR display
SEPAY_QR_BANK=...
SEPAY_QR_ACCOUNT=...
```

Rules:

- Never expose bearer tokens or secret keys in browser or mobile clients.
- Do not reuse Live credentials against Test endpoints or Test credentials against Live endpoints.
- Test mode acceptance is not Production readiness; Live requires HTTPS and may have feature differences.

## Webhook Security

For bank-account Webhooks:

1. Prefer HMAC-SHA256.
2. Verify the raw body before JSON parsing.
3. Require `X-SePay-Signature` and `X-SePay-Timestamp`.
4. Sign `{timestamp}.{raw_body}`.
5. Reject stale timestamps, normally older than five minutes.
6. Compare signatures with timing-safe comparison.
7. Treat API Key as acceptable only when HMAC is not selected.
8. Use no-auth only for short-lived tests.

For Payment Gateway IPN:

1. Use HTTPS.
2. Verify `X-Secret-Key` when the merchant config uses `SECRET_KEY`.
3. Return HTTP 200 after storing the event.
4. Do not expect bank-webhook HMAC headers unless the gateway docs explicitly add them.

## Idempotency

Always record an idempotency row before applying payment effects.

Suggested keys:

| Surface | Idempotency Key |
|---------|-----------------|
| Bank-account Webhook | webhook `id` |
| Bank transaction fallback | `referenceCode` + `transferType` + amount + account when no `id` is available |
| Payment Gateway IPN | `notification_type` + `order.order_invoice_number` + `transaction.transaction_id` |
| Gateway order API action | local order ID + gateway invoice/order ID + action |

Duplicate deliveries should return success after confirming the previous processing result. Never apply payment completion, refund, void, license grant, commission, or fulfillment twice.

## Payment Matching

Prefer deterministic identifiers over memo parsing:

1. Use SePay payment-code recognition and the webhook `code` field when configured.
2. Use official Order VA identifiers when the bank/account supports them.
3. Use exact invoice/order IDs in gateway `order_invoice_number`.
4. Parse `content` only as a controlled fallback.
5. Use amount/time matching only for manual review or low-risk back-office workflows, not as the sole automatic confirmation path.

Memo rules:

- Keep payment codes short and bank-safe.
- Avoid long UUID-only instructions when a configured payment code works.
- For memo-based VA, include `TKP` plus the VA code.
- For VietinBank personal and household-business accounts, include `SEVQR`.
- Document each bank-specific transformation in tests or runbooks when the app relies on memo parsing.

## Amount Policy

Define amount behavior before implementation:

- exact amount required
- overpayment accepted and flagged
- underpayment rejected or moved to manual review
- partial payment allowed or forbidden
- fee handling
- currency always VND unless gateway docs for the selected method say otherwise

Order VAs and Payment Gateway methods can have stricter amount rules than simple bank-transfer memo matching. Do not copy one surface's amount policy to another.

## Fast Acknowledgement

Webhook/IPN handlers should:

1. Verify authentication.
2. Store the raw event and idempotency key.
3. Return the required success response quickly.
4. Process heavy side effects through a queue or background job.
5. Reconcile later if side effects fail after acknowledgement.

Do not block SePay callbacks on email, license generation, CRM sync, notification posts, analytics, or other non-payment side effects.

## Reconciliation

Use reconciliation for:

- missed webhooks
- failed or delayed IPN handling
- deployment incidents
- manual replays
- unmatched transactions
- amount mismatches
- status drift between local orders and SePay

Bank-account Webhooks reconcile through API v2 transaction lookup. Payment Gateway IPN reconciles through Payment Gateway order detail APIs. Keep those code paths separate.

## Monitoring

Production readiness should include:

- webhook/IPN auth failure metrics
- callback latency
- duplicate delivery count
- unmatched payment-code count
- amount mismatch count
- retry and manual replay count
- SePay delivery log review process
- alert channels for repeated failures
- incident workflow for stuck or missed payments
- periodic reconciliation job

Use SePay dashboard delivery logs and metrics as external evidence during incidents.

## IP Allowlisting

Use SePay's current IP address page for firewall allowlists:

https://developer.sepay.vn/en/dia-chi-ip

Rules:

- Allowlist both IPv4 and IPv6 values shown by SePay when firewalling callbacks.
- Re-check the official page periodically.
- Do not treat IP filtering as authentication.
- Keep HMAC/API Key/OAuth2 or `X-Secret-Key` verification enabled.

## Test Mode And Release Gates

Test mode constraints:

- 500 simulated transactions per day, reset at 00:00 Vietnam time.
- 50 bank accounts.
- 100 VAs per bank account.
- 50 webhooks.
- 50 API Access tokens.
- Test mode webhooks accept HTTP and self-signed certificates.
- Live webhooks require valid HTTPS.
- Test mode account selections are isolated from Live accounts.

Release gate before Live:

- Live credentials stored server-side.
- HTTPS endpoint deployed.
- HMAC/API Key/OAuth2 or IPN secret verification tested.
- Idempotency tested with duplicate callback.
- Manual replay tested without duplicate effects.
- Reconciliation job tested.
- Underpayment/overpayment policy tested.
- Bank-specific QR/VA memo rules tested for the chosen bank/account type.

## Common Pitfalls

1. Starting new API work on legacy `userapi/*`.
2. Treating bank Webhooks and Payment Gateway IPN as the same callback.
3. Verifying HMAC after parsing and re-stringifying JSON.
4. Returning non-2xx for already-processed duplicate callbacks.
5. Completing orders from redirect URLs instead of callback/IPN evidence.
6. Assuming QR image generation confirms payment.
7. Using amount-only matching for automatic fulfillment.
8. Ignoring Test mode vs Live credential separation.
9. Hardcoding copied SePay IPs without checking the current official page.
10. Applying one bank's VA or memo rules to every bank.
