# Polar Webhooks

Production guidance for Polar webhook setup, Standard Webhooks verification, delivery behavior, event taxonomy, idempotency, and fulfillment.

## Source Links

- Setup webhooks: https://polar.sh/docs/integrate/webhooks/endpoints
- Handle deliveries: https://polar.sh/docs/integrate/webhooks/delivery
- Webhook events: https://polar.sh/docs/integrate/webhooks/events
- TypeScript SDK: https://polar.sh/docs/integrate/sdk/typescript
- Next.js adapter: https://polar.sh/docs/integrate/sdk/adapters/nextjs
- Express adapter: https://polar.sh/docs/integrate/sdk/adapters/express

## Setup

1. In Polar dashboard, open organization webhook settings.
2. Add a public endpoint URL.
3. Use Raw delivery format for custom app integrations.
4. Set a webhook secret.
5. Subscribe to the event types your app handles.
6. Use sandbox and `polar listen` for local development when appropriate.

Requirements:

- Publicly reachable route.
- No auth middleware blocking Polar delivery.
- No redirect URL. Polar treats redirects as failures.
- Return 2xx after signature validation and durable enqueue/logging.

## Verification

Polar follows the Standard Webhooks specification. Prefer official SDK/adapters over custom signature code.

Headers:

```text
webhook-id: msg_...
webhook-timestamp: 1710000000
webhook-signature: v1,...
```

### TypeScript SDK

Use the raw request body.

```typescript
import express from "express";
import { validateEvent } from "@polar-sh/sdk/webhooks";

const app = express();

app.post(
  "/webhooks/polar",
  express.raw({ type: "application/json" }),
  async (req, res) => {
    let event;

    try {
      event = validateEvent(
        req.body,
        req.headers,
        process.env.POLAR_WEBHOOK_SECRET!,
      );
    } catch {
      return res.status(400).json({ error: "invalid_signature" });
    }

    await persistWebhookForAsyncProcessing(event);
    return res.status(202).json({ received: true });
  },
);
```

### Next.js Adapter

```typescript
import { Webhooks } from "@polar-sh/nextjs";

export const POST = Webhooks({
  webhookSecret: process.env.POLAR_WEBHOOK_SECRET!,
  onOrderPaid: async (payload) => {
    await fulfillPaidOrder(payload.data);
  },
  onCustomerStateChanged: async (payload) => {
    await syncCustomerAccess(payload.data);
  },
});
```

### Express Adapter

```typescript
import express from "express";
import { Webhooks } from "@polar-sh/express";

const app = express();

app.use(express.json()).post(
  "/webhooks/polar",
  Webhooks({
    webhookSecret: process.env.POLAR_WEBHOOK_SECRET!,
    onPayload: async (payload) => {
      await persistWebhookForAsyncProcessing(payload);
    },
  }),
);
```

When using framework adapters, follow their middleware order exactly. If you roll your own validation, preserve raw bytes and verify against the exact body received.

## Delivery Behavior

Current official behavior:

- Polar times out delivery requests after 10 seconds.
- Polar recommends handlers respond within 2 seconds.
- Failed delivery is retried up to 10 times with exponential backoff.
- Endpoint is disabled after 10 consecutive failed deliveries.
- Redirect responses are failures; Polar does not follow them.

Implementation pattern:

1. Verify signature.
2. Insert event into a durable table/queue using event ID or webhook ID as idempotency key.
3. Return 2xx quickly.
4. Process fulfillment asynchronously.
5. Reconcile missed/failed processing through API lookup.

## Event Taxonomy

### Checkout

- `checkout.created`
- `checkout.updated`
- `checkout.expired`

### Customer

- `customer.created`
- `customer.updated`
- `customer.deleted`
- `customer.state_changed`

`customer.state_changed` includes active subscriptions, granted benefits, and active meter balances. It is often the cleanest entitlement sync event.

### Subscription

- `subscription.created`
- `subscription.active`
- `subscription.uncanceled`
- `subscription.canceled`
- `subscription.past_due`
- `subscription.updated`
- `subscription.revoked`

Use `subscription.updated` as a catch-all, but handle specific events where the business behavior differs.

### Order

- `order.created`
- `order.updated`
- `order.paid`
- `order.refunded`

Important: `order.created` can be pending. Fulfillment that requires money received should happen on `order.paid`.

`billing_reason` values include:

- `purchase`
- `subscription_create`
- `subscription_cycle`
- `subscription_update`

### Refund

- `refund.created`
- `refund.updated`

### Benefit and Benefit Grant

- `benefit.created`
- `benefit.updated`
- `benefit_grant.created`
- `benefit_grant.updated`
- `benefit_grant.revoked`

### Product and Organization

- `product.created`
- `product.updated`
- `organization.updated`

## Subscription Sequences

End-of-period cancellation:

1. `subscription.updated`
2. `subscription.canceled`
3. At period end: `subscription.updated`
4. `subscription.revoked`

Immediate revoke:

1. `subscription.updated`
2. `subscription.canceled`
3. `subscription.revoked`

Renewal:

1. `subscription.updated`
2. `order.created` with `billing_reason = subscription_cycle`
3. `order.updated`
4. `order.paid`

For access control, revoke benefits when the subscription is revoked or when customer state no longer includes the entitlement.

## Handler Pattern

```typescript
async function handlePolarEvent(event: { id?: string; type: string; data: any }) {
  const key = event.id ?? `${event.type}:${event.data?.id}`;
  const firstSeen = await insertWebhookEventOnce(key, event.type, event.data);
  if (!firstSeen) return;

  switch (event.type) {
    case "order.paid":
      await fulfillPaidOrder(event.data);
      break;
    case "customer.state_changed":
      await syncAccessFromCustomerState(event.data);
      break;
    case "subscription.revoked":
      await revokeSubscriptionAccess(event.data);
      break;
    case "refund.updated":
      await syncRefundState(event.data);
      break;
    default:
      await markUnhandledButStored(key);
  }
}
```

## Idempotency Storage

Store:

- webhook ID/event ID
- event type
- raw payload or normalized payload
- provider (`polar`)
- received timestamp
- processing status
- attempts
- last error
- related Polar object ID
- related internal user/order/subscription ID if known

Use a unique constraint on the event ID or webhook ID.

## Troubleshooting

- 404: route mismatch, trailing slash, or deployment path mismatch.
- 3xx: configured URL redirects; use the final URL.
- 403: auth middleware or WAF blocked Polar; exempt webhook route and review firewall settings.
- Invalid signature: wrong secret, parsed/re-serialized body, wrong base64 handling if custom verifier is used.
- Slow endpoint: respond within 2 seconds and process asynchronously.

## IP Allowlist

IP allowlisting is defense in depth. It does not replace signature verification. Check official docs before hard-coding ranges because they can change.

## Script Helper Note

`scripts/polar-webhook-verify.js` exists as an inspectable local helper. For new production integrations, prefer `@polar-sh/sdk/webhooks` or framework adapters. If changing the helper, treat it as source code: run impact analysis, update tests, and validate separately.
