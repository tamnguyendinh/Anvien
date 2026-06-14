# Polar Best Practices

Production patterns for Polar integrations across checkout, products, webhooks, subscriptions, benefits, usage billing, customer portal, rate limits, fees, and reconciliation.

## Source Links

- API overview: https://polar.sh/docs/api-reference/introduction
- Fees: https://polar.sh/docs/merchant-of-record/fees
- Checkout API: https://polar.sh/docs/features/checkout/session
- Webhook delivery: https://polar.sh/docs/integrate/webhooks/delivery
- Customer State: https://polar.sh/docs/integrate/customer-state

## Production Defaults

- Use server-side Organization Access Tokens only.
- Use sandbox for all first-pass checkout, webhook, subscription, refund, usage, and portal tests.
- Use product IDs with `products` for checkout creation.
- Use `external_customer_id` / `externalCustomerId` for app user reconciliation.
- Fulfill paid access from `order.paid` or Customer State, not success redirects.
- Prefer Customer Portal for customer self-service.
- Prefer SDK/adapters for webhook validation.
- Record all webhook events idempotently before processing.
- Re-check official fee/rate docs before implementing public calculators or quota assumptions.
- Check supported countries, acceptable use, and account-review requirements before launch.

## Environment Variables

```bash
POLAR_ACCESS_TOKEN=polar_at_xxx
POLAR_WEBHOOK_SECRET=whsec_xxx
POLAR_SERVER=sandbox
POLAR_ORGANIZATION_ID=org_xxx

POLAR_PRODUCT_PRO_MONTHLY_ID=prod_xxx
POLAR_PRODUCT_PRO_YEARLY_ID=prod_xxx
```

Do not use one token across sandbox and production. Products, customers, discounts, webhooks, and orders are environment-specific.

Sandbox limitation: customer-facing emails are only delivered to organization members. Do not use arbitrary customer email addresses to validate sandbox email flows.

## Lazy Client Initialization

```typescript
import { Polar } from "@polar-sh/sdk";
import { z } from "zod";

const envSchema = z.object({
  POLAR_ACCESS_TOKEN: z.string().min(1),
  POLAR_SERVER: z.enum(["sandbox", "production"]).default("sandbox"),
});

let client: Polar | null = null;

export function getPolar() {
  if (!client) {
    const env = envSchema.parse(process.env);
    client = new Polar({
      accessToken: env.POLAR_ACCESS_TOKEN,
      server: env.POLAR_SERVER,
    });
  }
  return client;
}
```

Lazy initialization lets modules import during build without requiring secrets immediately.

## Checkout Pattern

```typescript
export async function createCheckoutForUser(user: User, productId: string) {
  const checkout = await getPolar().checkouts.create({
    products: [productId],
    successUrl: `${process.env.APP_URL}/billing/success?checkout_id={CHECKOUT_ID}`,
    returnUrl: `${process.env.APP_URL}/pricing`,
    externalCustomerId: user.id,
    customerEmail: user.email,
    customerName: user.name,
    metadata: {
      userId: user.id,
      source: "pricing",
    },
  });

  await db.checkoutAttempts.insert({
    checkoutId: checkout.id,
    userId: user.id,
    productId,
    status: checkout.status,
  });

  return checkout.url;
}
```

Rules:

- Validate product ID server-side.
- Store pending checkout/order context before redirect if reconciliation matters.
- Include `checkout_id={CHECKOUT_ID}` in success URL for lookup.
- Do not grant access on redirect alone.

## Webhook Processing Pattern

```typescript
async function persistThenProcess(event: PolarWebhookEvent) {
  const inserted = await db.webhookEvents.insertOnce({
    provider: "polar",
    eventId: event.id ?? `${event.type}:${event.data?.id}`,
    eventType: event.type,
    payload: event,
    processed: false,
  });

  if (!inserted) return;

  queue.add("polar-webhook", { eventId: inserted.eventId });
}
```

Processing worker:

```typescript
switch (event.type) {
  case "order.paid":
    await fulfillOrder(event.data);
    break;
  case "customer.state_changed":
    await syncCustomerEntitlements(event.data);
    break;
  case "subscription.revoked":
    await revokeSubscriptionAccess(event.data);
    break;
  default:
    await markUnhandled(event);
}
```

Rules:

- Verify signature before storing.
- Use raw body or framework adapter.
- Respond within 2 seconds where possible.
- Return non-2xx only for invalid signature or route-level failures.
- Store and process idempotently.

## Customer State as Entitlement Source

Customer State includes:

- Customer data.
- Active subscriptions.
- Granted benefits.
- Active meter balances.

Recommended pattern:

1. Set `externalCustomerId` at checkout/customer creation.
2. Subscribe to `customer.state_changed`.
3. Store a normalized entitlement snapshot locally.
4. Gate app features from local state.
5. Periodically reconcile critical customers from Polar.

This avoids fragile logic that infers access from only subscription status.

## Usage Billing Pattern

- Ingest usage server-side using `events.ingest`.
- Keep a local usage ledger for dedupe and audit.
- Design stable event names and metadata keys.
- Define meters around numeric metadata.
- Use Credits benefits for prepaid usage.
- Enforce usage limits in your app; Polar will not block usage when meter balance is exceeded.
- Use Customer State or Customer Meters API to sync remaining balance.

## Customer Portal Pattern

Use Customer Portal for:

- Payment method updates.
- Failed payment recovery.
- Subscription cancellation.
- Invoice/receipt access.
- Benefits and license keys.
- Seat management.
- Metered usage display.

Create customer sessions server-side and redirect to `customer_portal_url`.

## Rate Limits

Current official limits:

- Production: 500 requests per minute per organization/customer or OAuth2 client.
- Sandbox: 100 requests per minute per organization/customer or OAuth2 client.
- Unauthenticated license validate/activate/deactivate: 3 requests per second.

Handle `429` with `Retry-After`.

Pagination: list endpoints use `page` and `limit`; use `limit <= 100` and iterate until the response pagination says no more pages.

```typescript
async function callWithRetry<T>(fn: () => Promise<T>, attempts = 3): Promise<T> {
  let lastError: unknown;

  for (let i = 0; i < attempts; i++) {
    try {
      return await fn();
    } catch (error: any) {
      lastError = error;
      if (error.statusCode !== 429 && error.status !== 429) throw error;
      const retryAfter = Number(error.headers?.["retry-after"] ?? 1);
      await new Promise((resolve) => setTimeout(resolve, retryAfter * 1000));
    }
  }

  throw lastError;
}
```

## Fees

Do not hard-code Early Member fees as the default.

Current fee categories:

- Starter: current default free plan for new organizations.
- Pro/Growth/Scale: monthly plans with lower per-transaction rates.
- Early Member: only for organizations created before May 27, 2026 that stay on that plan.
- International card fee can apply.
- Subscription fee is separate only for Early Member according to current docs.

For public calculators:

- Link to official fee docs.
- Store fee config as data, not inline constants.
- Include plan selection.
- Treat tax separately from fees.

## Data Tables

Recommended tables:

- `polar_customers`: app user ID, Polar customer ID, external ID, email, environment.
- `polar_checkouts`: checkout ID, user ID, product IDs, status, metadata.
- `polar_orders`: order ID, billing reason, status, amount, currency, product ID, subscription ID.
- `polar_refunds`: refund ID, order ID, amount, tax amount, status, reason, benefit revocation state.
- `polar_discounts`: discount ID, code, type, redemption limits, campaign metadata.
- `polar_subscriptions`: subscription ID, customer ID, product ID, status, current period, cancel flag.
- `polar_entitlements`: user ID, benefit ID, grant ID, active flag, source event ID.
- `polar_webhook_events`: event ID/webhook ID, type, payload, processed state, attempts, error.
- `usage_ledger`: local usage ID, user ID, event name, units, status, Polar ingestion status.

## Reconciliation

Run reconciliation for:

- Webhook outages or disabled endpoints.
- Failed worker jobs.
- Payment disputes/refunds.
- Customer support corrections.
- Usage ingestion retry gaps.

At minimum, reconcile:

- Orders paid vs local fulfilled state.
- Refunds vs local access and support state.
- Discount usage vs local campaign/referral state.
- Active subscriptions vs local subscription table.
- Customer State vs local entitlements.
- Usage ledger vs events successfully ingested.

## Common Pitfalls

- Using `product_price_id` as the default checkout create field.
- Trusting success redirect as payment proof.
- Parsing JSON before custom webhook signature validation.
- Treating `order.created` as paid.
- Ignoring `customer.state_changed` when benefits are the real entitlement source.
- Letting customers call Core API routes with merchant privileges.
- Hard-coding old rate limits or old fee schedules.
- Assuming sandbox tokens/products/webhooks work in production.
- Building custom payment-method update flows instead of sending users to Customer Portal.
- Refunding a subscription order and assuming that cancels/revokes the subscription.
- Forgetting that checkout links create short-lived sessions and that generated session URLs expire.
