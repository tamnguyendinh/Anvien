# Polar Orders, Refunds, and Discounts

Commercial operations guidance for paid order signals, billing reasons, invoices, receipts, refunds, benefit consequences, and discounts.

## Source Links

- Orders: https://polar.sh/docs/features/orders
- Refunds: https://polar.sh/docs/features/refunds
- Discounts: https://polar.sh/docs/features/discounts
- Webhook events: https://polar.sh/docs/integrate/webhooks/events

## Orders

Every paid transaction on Polar is an order. Orders are created for:

- One-time purchases.
- Initial subscription checkout.
- Subscription renewals.
- Subscription changes that immediately invoice a prorated difference.

`billing_reason` values:

- `purchase`
- `subscription_create`
- `subscription_cycle`
- `subscription_update`

## Order Status

| Status | Meaning |
|--------|---------|
| `pending` | Order exists and Polar is attempting payment collection. |
| `paid` | Payment succeeded. |
| `refunded` | Fully refunded. |
| `partially_refunded` | Partially refunded. |
| `void` | Will not be collected. |

Free orders can become `paid` immediately with no payment step.

## Canonical Paid Signal

Use `order.paid` as the canonical paid signal for fulfillment. `order.created` means an order exists, but it can still be pending.

Recommended event handling:

```typescript
switch (event.type) {
  case "order.created":
    await recordOrder(event.data);
    break;
  case "order.paid":
    await fulfillPaidOrder(event.data);
    break;
  case "order.refunded":
    await syncRefundedOrder(event.data);
    break;
}
```

## Invoices and Receipts

Orders include amounts, tax, billing details, customer/product/subscription links, invoice data, and custom field data.

- Polar generates invoices for paid orders.
- Customers can download and edit invoices from Customer Portal.
- Once an invoice is generated, billing details are frozen; use Customer Portal for customer-side invoice correction.
- Receipts are proof of payment and include payment method, refunds, line items, taxes, totals, and linked invoice number.
- Receipt generation can be asynchronous; handle `202 Accepted` by retrying.

## Refunds

Refunds can be full or partial.

Important behavior:

- Maximum refundable amount is based on net amount excluding tax, less previous refunds and customer balance handling.
- Tax is refunded automatically: full on full refunds, prorated on partial refunds.
- Payment fees are not refunded.
- Polar can issue refunds on your behalf in some chargeback-prevention scenarios.

## Refunds and Benefits

One-time purchases:

- Refund flow can revoke product benefits such as files, license keys, Discord, or GitHub access.
- Benefit revocation is selected by default for full refunds in the dashboard.

Subscriptions:

- Refunding a subscription order returns money but does not end the subscription relationship.
- To end access, cancel or revoke the subscription.
- Benefits are revoked when the subscription itself is revoked.

Do not treat refund as subscription cancellation.

## Discount Types

Discounts can be:

- Percentage.
- Fixed amount.
- Recurring for once, several months, or forever.

Restrictions:

- Specific products or all products.
- Start date.
- End date.
- Maximum redemptions.

## Applying Discounts

| Mode | Behavior |
|------|----------|
| Checkout Link preset | Discount is silently applied when customer visits the link. |
| `discount_code` query parameter | Prefills the visible code field; customer can see it. |
| Checkout API `discount_id` | Programmatically applies a discount to a session. |

Discounts without a code can only be auto-applied through Checkout Links or the API.

When creating checkout sessions, set `allowDiscountCodes: false` if you must prevent customer-entered codes from stacking with your pre-applied discount.

## Local Data Recommendations

Orders table:

- Polar order ID.
- billing reason.
- status.
- amount, discount amount, tax amount, total amount, currency.
- product ID.
- subscription ID.
- customer ID and external customer ID.
- checkout ID.
- invoice/receipt status if needed.
- source webhook event ID.

Refunds table:

- Polar refund ID.
- order ID.
- amount and tax amount.
- reason.
- status.
- benefit revocation choice if available.
- source webhook event ID.

Discount redemptions:

- discount ID.
- code.
- checkout ID.
- order ID.
- customer ID.
- local campaign/referral metadata.

## Best Practices

- Record `order.created`, but fulfill only on `order.paid`.
- Sync refunds separately from orders.
- Decide and document whether one-time-purchase refunds revoke access.
- For subscription refunds, explicitly decide whether to also cancel/revoke the subscription.
- Store billing reason for analytics and lifecycle handling.
- Use Customer Portal for customer invoice and receipt access.
- Reconcile local order/refund state from API after webhook outages.
