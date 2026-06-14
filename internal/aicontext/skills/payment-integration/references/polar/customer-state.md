# Polar Customer State

Customer State is the compact entitlement and access-sync surface for Polar customers.

## Source Links

- Customer State: https://polar.sh/docs/integrate/customer-state
- Webhook events: https://polar.sh/docs/integrate/webhooks/events
- Customer Portal: https://polar.sh/docs/features/customer-portal/introduction

## What Customer State Contains

Customer State includes:

- Customer data.
- Active subscriptions.
- Granted benefits.
- Active meters and current balances.

Use it when your app needs to answer: "What should this customer have access to right now?"

## Access Patterns

Use Customer State by:

- Polar customer ID.
- External customer ID, which should be your stable app user ID.
- `customer.state_changed` webhook for push-based sync.

Recommended model:

1. Set `externalCustomerId` / `external_customer_id` during checkout or customer creation.
2. On `customer.state_changed`, upsert a local entitlement snapshot.
3. Gate request-time access from local state.
4. Reconcile from Customer State API for critical access and support flows.

## Why Prefer Customer State for Entitlements

Subscription-only logic can miss:

- One-time purchase benefits.
- License keys or files attached through benefit grants.
- Feature Flag benefits.
- Credits and active meter balances.
- Product benefit changes affecting existing customers.

Customer State gives one normalized view across those surfaces.

## Webhook Handler

```typescript
async function handleCustomerStateChanged(customerState: CustomerState) {
  await db.customerEntitlements.replaceForCustomer({
    externalCustomerId: customerState.customer.externalId,
    polarCustomerId: customerState.customer.id,
    activeSubscriptions: customerState.activeSubscriptions,
    grantedBenefits: customerState.grantedBenefits,
    activeMeters: customerState.activeMeters,
    syncedAt: new Date(),
  });
}
```

If your SDK uses different property names, follow generated types. The important contract is the state shape: customer, active subscriptions, granted benefits, and meter balances.

## Local Entitlement Table

Store normalized rows:

- app user ID / external customer ID.
- Polar customer ID.
- benefit ID.
- benefit type.
- grant ID if present.
- subscription ID or order ID when relevant.
- active flag.
- meter ID and balance for usage gates.
- source event ID.
- synced timestamp.

## Access Checks

Feature flag:

```typescript
function hasFeature(state: LocalCustomerState, benefitId: string) {
  return state.benefits.some((benefit) => benefit.id === benefitId && benefit.active);
}
```

Usage gate:

```typescript
function canConsume(state: LocalCustomerState, meterId: string, units: number) {
  const meter = state.meters.find((item) => item.id === meterId);
  return Boolean(meter && meter.balance >= units);
}
```

## Boundaries

Customer State is not a replacement for every event:

- Use `order.paid` for paid-order fulfillment and revenue records.
- Use refund events for refund accounting and support workflows.
- Use subscription events for lifecycle analytics and dunning behavior.
- Use Customer State for "current access" decisions.

## Best Practices

- Treat Customer State as the source of truth for entitlements.
- Keep local access state synchronized from `customer.state_changed`.
- Reconcile state after webhook outages or support interventions.
- Use benefit IDs rather than names for access checks.
- Keep a degraded access policy for temporary Polar outages.
