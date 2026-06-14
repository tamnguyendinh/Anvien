# Polar Subscriptions

Current guidance for subscription lifecycle, plan changes, proration, trials, cancellation, revoke, renewals, and Customer Portal boundaries.

## Source Links

- Manage subscriptions: https://polar.sh/docs/features/subscriptions/manage
- Update subscription API: https://polar.sh/docs/api-reference/subscriptions/update
- Revoke subscription API: https://polar.sh/docs/api-reference/subscriptions/revoke
- Webhook events: https://polar.sh/docs/integrate/webhooks/events
- Customer Portal: https://polar.sh/docs/features/customer-portal/introduction

## Current Status Values

Subscription objects can use these statuses:

- `incomplete`
- `incomplete_expired`
- `trialing`
- `active`
- `past_due`
- `canceled`
- `unpaid`

Also watch these fields:

- `cancel_at_period_end`
- `canceled_at`
- `started_at`
- `ends_at`
- `ended_at`
- `current_period_start`
- `current_period_end`
- `trial_start`
- `trial_end`
- `product_id`
- `discount_id`
- `seats`
- `meters`
- `pending_update`

## Listing and Reading

```typescript
const subscriptions = await polar.subscriptions.list({
  customerId: "customer_id",
  productId: "product_id",
  status: "active",
});

const subscription = await polar.subscriptions.get({
  id: subscriptionId,
});
```

Use `external_customer_id`/`externalCustomerId` where supported by the endpoint, or resolve the Polar customer from your own user ID first.

## Changing Plans

Update to another product with `product_id` / `productId`.

```typescript
await polar.subscriptions.update({
  id: subscriptionId,
  subscriptionUpdate: {
    productId: newProductId,
    prorationBehavior: "prorate",
  },
});
```

REST shape:

```json
{
  "product_id": "new_product_id",
  "proration_behavior": "prorate"
}
```

Do not use `product_price_id` for current plan changes.

## Proration Behavior

Current options:

| Value | Meaning |
|-------|---------|
| `invoice` | Invoice immediately for the change. |
| `prorate` | Apply prorated credit/charge according to Polar behavior. |
| `next_period` | Schedule the change for the next period. |
| `reset` | Reset billing period where applicable. |

If omitted, Polar uses the organization default.

## Discounts

```typescript
await polar.subscriptions.update({
  id: subscriptionId,
  subscriptionUpdate: {
    discountId: discountId,
  },
});

await polar.subscriptions.update({
  id: subscriptionId,
  subscriptionUpdate: {
    discountId: null,
  },
});
```

Discount removal or change may apply on the next billing cycle. Verify exact product behavior in sandbox.

## Trials

Products can define trials. Checkout can also override or disable trials.

Subscription update can set or extend `trial_end`; using `now` ends the trial immediately.

```json
{
  "trial_end": "now"
}
```

Handle:

- `subscription.created`: subscription object created.
- `subscription.active`: subscription becomes active.
- `subscription.past_due`: payment failed.
- `order.paid`: paid invoice/order.

## Cancellation vs Revoke

### Cancel at Period End

Customer portal cancellation and merchant scheduled cancellation preserve access until the current period ends.

Expected event sequence:

1. `subscription.updated`
2. `subscription.canceled`
3. At period end: `subscription.updated`
4. `subscription.revoked`

During the waiting period, `cancel_at_period_end` is true and status can still be active.

### Immediate Revoke

Immediate revoke cancels now and stops access now.

REST:

```http
DELETE /v1/subscriptions/{id}
```

SDKs expose this as a revoke operation, for example `polar.subscriptions.revoke(...)`.

Expected event sequence:

1. `subscription.updated`
2. `subscription.canceled`
3. `subscription.revoked`

## Renewal Handling

Renewal sequence:

1. `subscription.updated`
2. `order.created` with `billing_reason = subscription_cycle`
3. `order.updated`
4. `order.paid`

`order.created` can be pending. Record it for reconciliation, but grant paid-period entitlements on `order.paid` or from Customer State after payment.

## Failed Payments

Handle `subscription.past_due`.

Recommended app behavior:

- Keep a local grace policy explicit.
- Notify the user.
- Send the user to Customer Portal to update payment method.
- Sync final access from `customer.state_changed` or `subscription.revoked`.

## Customer Portal Boundary

Use Customer Portal for customer self-service:

- Cancel subscription.
- Update payment method.
- View invoices and receipts.
- Manage benefits.
- Manage seats when enabled.
- View metered usage if enabled.

Do not build privileged Core API endpoints for ordinary user self-service when a Customer Portal session is sufficient. Load `customer-portal.md`.

## Local Data Model

Store:

- Polar subscription ID.
- Polar customer ID.
- external customer ID.
- product ID.
- status.
- current period start/end.
- trial start/end.
- cancel-at-period-end flag.
- discount ID.
- seats.
- latest order ID for renewal if needed.
- last synced event ID.

For entitlement decisions, prefer Customer State or benefit grants over locally inferred subscription-only logic when benefits are attached.

## Best Practices

- Use product IDs for plan changes.
- Use `proration_behavior`, not stale proration names.
- Distinguish scheduled cancellation from immediate revoke.
- Treat `order.created` as not-yet-paid unless status says otherwise.
- Fulfill renewal/payment access on `order.paid` or Customer State.
- Let Customer Portal handle user payment method recovery and subscription self-service.
- Reconcile periodically against API for critical access or revenue state.
