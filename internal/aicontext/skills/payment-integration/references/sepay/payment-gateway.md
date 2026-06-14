# SePay Payment Gateway And IPN

Use this reference for hosted checkout, gateway payment methods, Payment Gateway API, and gateway IPN. Do not treat this as the same surface as bank-account Webhooks.

Official docs checked: 2026-06-14.

## When To Use Payment Gateway

Use Payment Gateway when the product needs:

- hosted checkout
- card payments
- bank-transfer QR through the gateway
- NAPAS bank transfer
- gateway order lifecycle APIs
- gateway IPN for order or transaction outcomes

Use bank-account Webhooks instead when the product receives direct bank transfers into a linked account and needs real-time transaction events.

## Base URLs And Auth

| Environment | Base URL |
|-------------|----------|
| Production | `https://pgapi.sepay.vn` |
| Sandbox | `https://pgapi-sandbox.sepay.vn` |

Payment Gateway API uses Basic Auth:

```http
Authorization: Basic base64(merchant_id:secret_key)
Content-Type: application/json
```

Keep `merchant_id` and `secret_key` server-side. Never send them to browser or mobile clients.

## Checkout Init

Checkout form submission uses:

```http
POST /v1/checkout/init
```

This is a payment form submission flow, not the same as the API v2 transaction lookup surface.

Core fields:

- `merchant`
- `operation` (`PURCHASE` for one-time payment)
- `payment_method` (`CARD`, `BANK_TRANSFER`, or `NAPAS_BANK_TRANSFER`)
- `order_amount`
- `currency` (`VND`)
- `order_invoice_number` (unique)
- `order_description`
- `customer_id`
- `success_url`
- `error_url`
- `cancel_url`
- `signature`

Signature rules:

- Sign only the documented allowlist of fields.
- Preserve SePay's required field order when building custom forms/signatures.
- Use HMAC-SHA256 and base64 output as documented by SePay.
- Duplicate invoice numbers or reordered signing fields can make the checkout fail.

Prefer the official SDK when possible because it creates the checkout fields and signature in the expected order.

## Gateway IPN Contract

Configure IPN in SePay Payment Gateway configuration.

Rules:

- IPN URL must use HTTPS.
- Return HTTP 200 to acknowledge receipt.
- If the merchant chooses auth type `SECRET_KEY`, verify `X-Secret-Key`.
- Do not assume bank-webhook HMAC headers are present for Payment Gateway IPN.

Payload shape:

```json
{
  "timestamp": 1757058220,
  "notification_type": "ORDER_PAID",
  "order": {},
  "transaction": {},
  "customer": {}
}
```

Known notification types include:

- `ORDER_PAID`
- `TRANSACTION_VOID`

Idempotency keys:

- `order.id`
- `order.order_id`
- `order.order_invoice_number`
- `transaction.id`
- `transaction.transaction_id`
- `notification_type`

Store the IPN event before applying payment effects. Replays, retries, and manual operator actions must not double-complete an order or double-void a transaction.

## Order Lifecycle APIs

Use Payment Gateway order APIs for gateway order state, not bank-account transaction lookup.

Common operations:

```http
GET /v1/order/detail/{order_id}
POST /v1/order/cancel
```

Order detail responses include:

- order status
- order amount and currency
- order invoice number
- authentication status
- transaction history

Important statuses:

- `CAPTURED`: paid
- `CANCELLED`: cancelled
- `AUTHENTICATION_NOT_NEEDED`: awaiting payment

Cancel rules:

- Applies to `BANK_TRANSFER` and `NAPAS_BANK_TRANSFER`.
- Only cancel orders that are not already captured or cancelled.
- Use the invoice number expected by the endpoint/SDK.

Card void rules:

- Card transaction void belongs to the gateway order/transaction lifecycle.
- Do not model it as a bank-account transaction reversal.

## Reconciliation

IPN is the real-time gateway notification, but the merchant system should still reconcile:

1. Receive IPN.
2. Verify `X-Secret-Key` when configured.
3. Store idempotency record.
4. Load local order by `order_invoice_number` or the merchant's mapped order ID.
5. Validate amount, currency, and expected status transition.
6. Apply payment or void effect exactly once.
7. Return HTTP 200 quickly.
8. Backfill ambiguous or missed events with the Payment Gateway order detail API.

If the same business also receives direct bank transfers, keep gateway reconciliation and bank-account webhook reconciliation in separate code paths with separate idempotency namespaces.

## IP Allowlisting

SePay publishes current outbound IPs for Webhooks, Payment Gateway IPN, Bank Hub IPN, and other callbacks at:

https://developer.sepay.vn/en/dia-chi-ip

Use that page for firewall rules. Do not hardcode copied IPs as permanent guidance. IP allowlisting is an additional control; it does not replace `X-Secret-Key` or other documented authentication.

## Related References

- `sdk.md`: official Node.js and PHP SDK guidance for gateway form generation and order APIs.
- `webhooks.md`: bank-account transaction webhooks, separate from gateway IPN.
- `api.md`: API v2 for proactive SePay bank-account lookup and Order VAs, separate from gateway orders.
