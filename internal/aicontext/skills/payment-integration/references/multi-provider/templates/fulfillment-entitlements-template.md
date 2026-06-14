# Fulfillment and Entitlements Template

Use this template when a provider paid signal must update local order state and grant access. The paid handler should be shared; provider-specific code should only extract trusted facts and resolve the local order.

## Paid Facts

```typescript
interface ProviderPaidFacts {
  provider: PaymentProvider;
  providerEnvironment: ProviderEnvironment;
  localOrderId?: string;
  providerOrderId?: string;
  providerTransactionId?: string;
  providerPaymentIntentId?: string;
  providerSubscriptionId?: string;
  amountMinor: number;
  currency: string;
  taxAmountMinor?: number;
  discountAmountMinor?: number;
  providerFeeMinor?: number;
  providerNetAmountMinor?: number;
  paidAt: string;
  raw: unknown;
}
```

## Paid Handler

```typescript
async function handlePaidProviderEvent(event: ProviderEvent) {
  const facts = await extractPaidFacts(event);
  const order = await findLocalOrderForPaidFacts(facts);

  if (!order) {
    await flagProviderEventForManualReview(event.id, "no_local_order_match");
    return;
  }

  if (!amountMatchesPolicy(order, facts)) {
    await flagProviderEventForManualReview(event.id, "amount_or_currency_mismatch");
    return;
  }

  const firstPaidTransition = await markOrderPaid({
    orderId: order.id,
    providerStatus: getProviderPaidStatus(event.provider, event.eventType),
    paidAt: facts.paidAt,
    taxAmountMinor: facts.taxAmountMinor,
    discountAmountMinor: facts.discountAmountMinor,
    providerFeeMinor: facts.providerFeeMinor,
    providerNetAmountMinor: facts.providerNetAmountMinor,
  });

  if (!firstPaidTransition) return;

  await grantEntitlementsForPaidOrder(order.id);
  await createPendingCommissionsForPaidOrder(order.id);
}
```

Important rules:

- `markOrderPaid` must be idempotent.
- If a later event provides better provider fee/net amount data, update financial fields without regranting entitlements.
- Trust provider paid signals over customer success redirects.
- If a success redirect arrives before paid webhook, keep local state as `checkout_created` or `awaiting_payment`.

## Entitlement Snapshot

Access state is not order state. Store access separately so subscriptions, licenses, benefits, manual grants, and one-time purchases can be reconciled.

```typescript
interface LocalEntitlement {
  userId: string;
  orderId?: string;
  provider: PaymentProvider;
  providerEnvironment: ProviderEnvironment;
  sourceObjectType: "order" | "subscription" | "benefit" | "license" | "manual";
  sourceObjectId: string;
  entitlementKey: string;
  status: "active" | "inactive" | "revoked" | "expired";
  quantity?: number;
  startsAt?: string;
  endsAt?: string;
  updatedAt: string;
}
```

Provider notes:

- Polar: use Customer State for current access and `order.paid` for paid-order accounting.
- Creem.io: license activation/deactivation events can be an access source.
- Stripe/Paddle: subscription active/cancelled/past-due state should drive subscription entitlements.
- SePay: access is usually app-owned after verified bank/gateway payment.

## Entitlement Grant Example

```typescript
interface EntitlementPlan {
  entitlementKey: string;
  sourceObjectType: "order" | "subscription" | "benefit" | "license" | "manual";
  quantity?: number;
  durationDays?: number;
}

async function grantEntitlementsForPaidOrder(orderId: string) {
  const order = await getOrder(orderId);
  const plans = await resolveEntitlementPlans(order.productKey);

  for (const plan of plans) {
    await upsertEntitlement({
      userId: order.userId,
      orderId: order.id,
      provider: order.provider,
      providerEnvironment: order.providerEnvironment,
      sourceObjectType: plan.sourceObjectType,
      sourceObjectId: order.id,
      entitlementKey: plan.entitlementKey,
      status: "active",
      quantity: plan.quantity,
      startsAt: order.paidAt,
      endsAt: plan.durationDays ? addDays(order.paidAt, plan.durationDays) : undefined,
    });
  }
}
```

Example product-to-entitlement mapping:

```typescript
const entitlementCatalog: Record<string, EntitlementPlan[]> = {
  "one_time_download": [
    { entitlementKey: "download:product_files", sourceObjectType: "order" },
  ],
  "pro_monthly": [
    { entitlementKey: "app:pro_features", sourceObjectType: "subscription" },
    { entitlementKey: "usage:monthly_credits", sourceObjectType: "benefit", quantity: 1000 },
  ],
  "team_seat": [
    { entitlementKey: "team:seat", sourceObjectType: "subscription", quantity: 1 },
  ],
};
```

Those keys are examples. A course platform might use `course:access`; a desktop app might use `license:activation`; an API product might use `api:quota`; an internal tool might use `workspace:seat`.

## Subscription Sync Shape

```typescript
async function syncSubscriptionFromProvider(event: ProviderEvent) {
  const subscription = await fetchProviderSubscription(event);
  const localSubscription = await upsertLocalSubscriptionFromProvider(subscription);
  const entitlementPlans = await resolveEntitlementPlans(localSubscription.productKey);

  for (const plan of entitlementPlans) {
    await upsertEntitlement({
      userId: localSubscription.userId,
      provider: event.provider,
      providerEnvironment: event.providerEnvironment,
      sourceObjectType: "subscription",
      sourceObjectId: localSubscription.providerSubscriptionId,
      entitlementKey: plan.entitlementKey,
      status: localSubscription.status === "active" ? "active" : "inactive",
      quantity: plan.quantity,
      startsAt: localSubscription.currentPeriodStart,
      endsAt: localSubscription.currentPeriodEnd,
    });
  }
}
```
