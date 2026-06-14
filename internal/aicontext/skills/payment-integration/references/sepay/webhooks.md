# SePay Bank-Account Webhooks

Use this reference for real-time bank-account transaction notifications. Payment Gateway IPN is a separate surface; load `payment-gateway.md` for hosted checkout IPN.

Official docs checked: 2026-06-14.

## Setup

Create webhooks from the SePay dashboard.

Configure:

- name
- event type: incoming, outgoing, or all transactions
- payload format: JSON is recommended
- source bank accounts
- payment verification filters
- payment-code filters
- endpoint URL
- retry-on-error behavior
- authentication method
- alert channels in Live mode

Live webhook URLs require valid HTTPS. Test mode accepts HTTP and self-signed certificates, but that does not prove Live readiness.

## Payload

Typical JSON payload:

```json
{
  "id": 92704,
  "gateway": "Vietcombank",
  "transactionDate": "2023-03-25 14:02:37",
  "accountNumber": "0123499999",
  "code": null,
  "content": "payment content",
  "transferType": "in",
  "transferAmount": 2277000,
  "accumulated": 19077000,
  "subAccount": null,
  "referenceCode": "MBVCB.3278907687"
}
```

Use `id` as the primary bank-transaction idempotency key. Include `referenceCode`, `transferType`, and `transferAmount` as supporting audit fields.

## Authentication Methods

| Method | Use |
|--------|-----|
| None | Testing only. Do not use for production. |
| API Key | Medium security. SePay sends `Authorization: Apikey <API_KEY>`. |
| HMAC-SHA256 | Recommended production default. Detects payload tampering. |
| OAuth2 | High security when the receiver already operates an OAuth server. |

### HMAC-SHA256 Verification

Requirements:

- Read the raw request body before JSON parsing.
- Read `X-SePay-Signature`.
- Read `X-SePay-Timestamp`.
- Rebuild the signed string as `{timestamp}.{raw_body}`.
- Compute HMAC-SHA256 with the webhook secret.
- Compare with `sha256=<hex_hash>` using timing-safe comparison.
- Reject stale timestamps, normally older than five minutes.

Node.js example:

```javascript
import crypto from 'node:crypto';

export function verifySePayWebhook({ rawBody, signature, timestamp, secret }) {
  if (!rawBody || !signature || !timestamp || !secret) return false;

  const ts = Number(timestamp);
  if (!Number.isFinite(ts)) return false;

  const nowSeconds = Math.floor(Date.now() / 1000);
  if (Math.abs(nowSeconds - ts) > 300) return false;

  const expected = 'sha256=' + crypto
    .createHmac('sha256', secret)
    .update(`${timestamp}.${rawBody}`)
    .digest('hex');

  const expectedBytes = Buffer.from(expected);
  const providedBytes = Buffer.from(signature);
  if (expectedBytes.length !== providedBytes.length) return false;

  return crypto.timingSafeEqual(expectedBytes, providedBytes);
}
```

Framework warning:

- In Express, mount `express.raw()` or capture the raw body before `express.json()`.
- In Next.js/Fetch handlers, call `request.text()` for verification, then parse that exact string with `JSON.parse()`.
- Do not stringify a parsed object and sign the result; spacing and field order may change.

## API Key Verification

```javascript
function verifyApiKey(authHeader, expectedKey) {
  if (!authHeader?.startsWith('Apikey ')) return false;
  const provided = Buffer.from(authHeader.slice('Apikey '.length));
  const expected = Buffer.from(expectedKey);
  return provided.length === expected.length &&
    crypto.timingSafeEqual(provided, expected);
}
```

## Response Contract

Success:

```json
HTTP 200
{"success": true}
```

Rules:

- HTTP 200 or 201 with `{"success": true}` is the safe acknowledgement.
- Return quickly, then enqueue heavy work.
- SePay treats connection failures, timeouts, and non-2xx responses as delivery failures.
- A timeout is more than 30 seconds with no response.

## Retry And Replay

Automatic retry:

- Initial delivery plus 7 retries.
- Fibonacci spacing.
- Last automatic attempt is normally around 33 minutes after the first attempt.
- Queue items older than 5 hours are not picked up by cron, but do not rely on that as the normal retry window.

Manual replay:

- Operators can replay failed or successful webhook deliveries from the dashboard.
- Replays can duplicate a previously successful delivery.
- Endpoint processing must be idempotent and safe to call multiple times.

## Payment-Code Recognition

Prefer SePay payment-code recognition over custom memo parsing when available.

Configuration concepts:

- company-level enable/disable
- one or more templates
- prefix length
- suffix min/max length
- suffix character type: numeric or alphanumeric
- active/inactive template status
- matching order across active templates

When SePay matches a configured template, it sends the matched value in the payload `code` field. Use `code` as the primary order/payment reference. Fall back to parsing `content` only when the integration intentionally does not use SePay's payment-code recognition or when a legacy bank/account does not provide the field.

Webhook filters can require a payment code or filter by payment-code prefix. If a webhook is used for order confirmation, prefer requiring a code unless the business explicitly supports unmatched manual transfers.

## Processing Pattern

1. Verify authentication before parsing or processing.
2. Parse the payload from the verified raw body.
3. Reject unexpected `transferType` values for the current route.
4. Store an idempotency record keyed by `id`.
5. Resolve the order by `code` first.
6. Validate amount policy: exact, overpayment, underpayment, or manual review.
7. Apply the payment effect exactly once.
8. Record the SePay payload and processing result.
9. Return success quickly, even for duplicates already processed.
10. Reconcile later through API v2 if processing or downstream effects fail.

## Monitoring And Operations

Use SePay dashboard operations as part of production readiness:

- Delivery logs with request and response detail.
- cURL reproduction from delivery logs.
- Dashboard metrics: success rate, failures, timeouts, average response, P95 response.
- Error breakdown: DNS, timeout, 4xx, 5xx.
- Alerts to Telegram, Slack, or Discord.
- Incidents for repeated failures.
- Manual replay after fixes.

Application-side monitoring should track:

- auth failures
- duplicate deliveries
- unmatched payment codes
- underpayments and overpayments
- processing latency
- reconciliation misses
- replay counts

## Reconciliation

Use API v2 to reconcile:

- webhook downtime
- endpoint deployment incidents
- retry window exhaustion
- manual replay verification
- unmatched transactions
- disputed amount or memo cases

Do not rely on webhooks as the only source of truth. Webhooks are the real-time trigger; API v2 is the lookup/backfill path.

## IP Allowlisting

Use SePay's current IP page when configuring firewalls:

https://developer.sepay.vn/en/dia-chi-ip

Rules:

- Allowlisting is optional hardening, not authentication.
- Keep HMAC/API Key/OAuth2 verification even when IP filtering is enabled.
- Do not bake copied IP addresses into reusable skill guidance; the list can change.
- Live callback URLs must use HTTPS.

## Explicit Non-Matches

Do not use this file for:

- hosted checkout order IPN
- `X-Secret-Key` gateway IPN verification
- gateway order detail/cancel/void APIs
- gateway card or NAPAS lifecycle

Load `payment-gateway.md` for those flows.
