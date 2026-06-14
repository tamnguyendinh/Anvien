# Polar Usage-Based Billing

Current guidance for Polar usage events, meters, metered prices, credits, customer meter balances, and usage enforcement.

## Source Links

- Usage billing introduction: https://polar.sh/docs/features/usage-based-billing/introduction
- Event ingestion: https://polar.sh/docs/features/usage-based-billing/event-ingestion
- Meters: https://polar.sh/docs/features/usage-based-billing/meters
- Credits: https://polar.sh/docs/features/usage-based-billing/credits
- Customer State: https://polar.sh/docs/integrate/customer-state

## Model

Usage-based billing is:

```text
Application usage -> Events -> Meters -> Metered prices -> Invoice/order
```

Core objects:

- Events: immutable usage records sent from your application.
- Meters: filters and aggregations over events.
- Metered prices: recurring product prices that bill for meter usage.
- Credits: prepaid meter balance issued through Credits benefits.
- Customer meters: per-customer meter balance and usage state.

## Event Shape

Events have:

- `name`: event type, such as `ai_usage` or `video_streamed`.
- `customer_id` or `external_customer_id`: Polar customer ID or your app user ID.
- `metadata`: numeric/string fields used for filtering and aggregation.
- Optional `timestamp`: historical time for dashboards; billing period attribution is based on when Polar receives the event.

REST-style event:

```json
{
  "name": "ai_usage",
  "external_customer_id": "user_123",
  "metadata": {
    "model": "gpt-4.1-nano",
    "requests": 1,
    "total_tokens": 77
  }
}
```

## Ingest Events with TypeScript SDK

Use `events.ingest`, not stale `events.create` examples.

```typescript
await polar.events.ingest({
  events: [
    {
      name: "ai_usage",
      externalCustomerId: user.id,
      metadata: {
        model: "gpt-4.1-nano",
        total_tokens: 77,
        requests: 1,
      },
    },
  ],
});
```

Server-side only:

- Never expose Organization Access Tokens to the browser.
- Ingest only usage your backend has verified.
- Deduplicate before ingestion if user actions can retry.
- Keep a local audit log of usage events and Polar ingestion result.

## Event Immutability and Backdating

Events are immutable after ingestion.

Backdated `timestamp` is useful for replay or batched ingestion, but late events are billed in the cycle when Polar receives them. Do not expect retroactive invoice changes for already closed billing periods.

## Ingestion Strategies

Official docs describe strategy patterns for common event sources. Treat them as implementation choices, not requirements for every app.

- LLM usage: record request count, prompt tokens, completion tokens, model, and cost metadata when useful.
- S3/object operations: batch object-storage events into trusted backend ingestion.
- Streams: aggregate high-volume streaming usage before ingesting if event volume is high.
- Delta-time usage: record time windows for usage that accrues over duration.

Keep local idempotency keys even when ingestion is batched.

## Meters

Meters filter and aggregate events.

Examples:

- Count requests where `name = ai_usage`.
- Sum `metadata.total_tokens`.
- Sum streamed seconds or uploaded bytes.

Design meters before production:

- Use stable event names.
- Use numeric metadata for sums.
- Avoid changing metadata semantics after launch.
- Test with sandbox events and expected customer meter balances.

## Metered Prices

Metered prices are attached to recurring products. They can be paired with a fixed base price and multiple meters.

Examples:

- $20/month base plus tokens.
- Free subscription plus metered API calls.
- Base plan plus separate storage and compute meters.

Load `products.md` before creating product pricing.

## Credits

Credits are prepaid meter balance issued through a Credits benefit.

Behavior:

- Subscription products credit the customer at the start of each subscription cycle.
- One-time products credit once at purchase.
- Credits are consumed before overage charges.
- For credits-only spending, do not add a metered price for the meter.
- Polar does not block usage when balance is exceeded; your app must enforce limits.

## Customer Meter Balance

Use Customer State or Customer Meters API to check balance.

Customer State includes:

- Customer data.
- Active subscriptions.
- Granted benefits.
- Active meters with balances.

Listen to `customer.state_changed` to keep local access and balance state current.

## Usage Enforcement Pattern

```typescript
async function recordUsage(user: User, tokens: number) {
  const state = await getLocalCustomerState(user.id);

  if (state.remainingTokens < tokens) {
    throw new Error("usage_limit_exceeded");
  }

  await appendLocalUsageLedger({
    userId: user.id,
    eventName: "ai_usage",
    units: tokens,
    status: "pending",
  });

  await polar.events.ingest({
    events: [
      {
        name: "ai_usage",
        externalCustomerId: user.id,
        metadata: { total_tokens: tokens },
      },
    ],
  });

  await markLocalUsageLedgerIngested(user.id);
}
```

Do not call Polar synchronously for every request-time authorization check if latency matters. Cache/sync Customer State locally, then reconcile.

## Cost Insights

Polar Cost Insights can use `_cost` metadata to estimate profitability for usage events. Treat `_cost` as optional analytics input, not as the billing source of truth.

## Best Practices

- Use `events.ingest` with arrays.
- Use `externalCustomerId` when app user ID is your reconciliation key.
- Keep event names stable and low-cardinality.
- Store detailed request IDs in local logs, not necessarily in Polar metadata.
- Ingest from backend workers or trusted server routes only.
- Reconcile usage ledger vs Polar ingestion on retries and outages.
- Make overage behavior explicit in product copy and customer portal UX.
