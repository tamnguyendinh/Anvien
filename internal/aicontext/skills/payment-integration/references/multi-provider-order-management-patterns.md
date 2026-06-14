# Multi-Provider Order Management Patterns

Provider-neutral order, payment, refund, entitlement, revenue, and reconciliation patterns for apps that use more than one payment provider.

Use this reference when an app mixes providers such as SePay, Polar, Stripe, Paddle, or Creem.io. Always load the provider-specific references too; this file defines local data ownership and cross-provider patterns, not provider API authority.

## When to Use

Load this file when the app needs any of these:

- One local order table across multiple providers.
- Cross-provider revenue reporting or currency normalization.
- Unified refunds, chargebacks, disputes, and commission reversal.
- Separate checkout state from paid order state.
- Entitlement sync across subscriptions, license keys, benefits, or manual bank payments.
- Reconciliation after webhook outages or manual support actions.

Do not use this file as a replacement for provider docs:

- SePay: load `references/sepay/overview.md`, then `api.md`, `webhooks.md`, `payment-gateway.md`, or `qr-codes.md` as needed.
- Polar: load `references/polar/overview.md`, then `orders-refunds-discounts.md`, `customer-state.md`, `customer-portal.md`, `usage-based-billing.md`, and related refs as needed.
- Stripe: load `references/stripe/stripe-best-practices.md` first.
- Paddle: load `references/paddle/overview.md`, `api.md`, `paddle-js.md`, `subscriptions.md`, `webhooks.md`, and `sdk.md` as needed.
- Creem.io: load `references/creem/overview.md`, `api.md`, `checkouts.md`, `subscriptions.md`, `licensing.md`, `webhooks.md`, and `sdk.md` as needed.

## Core Design Rule

Keep provider facts and local business facts separate.

Provider facts are the raw external records: checkout sessions, orders, transactions, subscriptions, refunds, events, customer IDs, and license IDs. Local business facts are your app decisions: access granted, commission approved, refund policy, reporting currency, and support notes.

Do not collapse every provider state into one vague `completed` value. Store:

- local normalized status for app workflows;
- provider raw status for audit and debugging;
- provider event history for replay and reconciliation;
- separate entitlement state for access decisions.

## Local Status Model

Use a small normalized status set locally and map provider events into it.

| Local Status | Meaning |
|--------------|---------|
| `draft` | Local order exists before checkout or payment instruction is created. |
| `checkout_created` | Provider checkout/session/transaction/payment instruction exists. |
| `awaiting_payment` | Customer has not produced a confirmed paid signal yet. |
| `paid` | Provider paid signal has been verified. Fulfillment may proceed. |
| `partially_refunded` | Some money has been refunded; access policy is app-specific. |
| `refunded` | Fully refunded; access and commission reversal depend on policy/provider. |
| `failed` | Payment failed or provider rejected the attempt. |
| `expired` | Checkout or bank-payment instruction expired. |
| `cancelled` | Customer or merchant cancelled before payment or ended a subscription. |
| `disputed` | Chargeback/dispute opened; fulfillment and payout risk need review. |

Keep subscription lifecycle separate from one-time order state. A subscription can be `active` while an individual renewal order is `paid`, `failed`, or `refunded`.

## Provider State Mapping

Use provider-specific paid signals as the authority for money received.

| Provider | Local Paid Signal | Notes |
|----------|-------------------|-------|
| SePay bank webhook | Verified bank-account webhook transaction matched to the local order and amount policy. | HMAC/API-key/OAuth rules depend on the SePay surface. Load `sepay/webhooks.md`. |
| SePay Payment Gateway | Verified IPN or gateway order state such as `ORDER_PAID`. | Load `sepay/payment-gateway.md`; gateway IPN is not the same surface as bank-account webhooks. |
| Polar | `order.paid` or verified API lookup showing paid order state. | `order.created` is not enough for fulfillment. Customer State is best for current access, not revenue recognition. |
| Stripe | Checkout Session or PaymentIntent/Invoice paid state according to the chosen Stripe integration. | Use Checkout Sessions/Billing/PaymentIntents guidance; do not use Charges as the default. |
| Paddle | `transaction.completed` for successful payment; subscription events for recurring lifecycle. | Load `paddle/webhooks.md` and `paddle/subscriptions.md`. |
| Creem.io | `checkout.completed` or `payment.succeeded`; subscription/license events for lifecycle and access. | Load `creem/webhooks.md`, `creem/subscriptions.md`, and `creem/licensing.md`. |

Refund and dispute events must map separately from paid events:

- Polar: `order.refunded` and refund API state.
- Creem.io: `refund.created`, `chargeback.created`.
- Paddle and Stripe: use their webhook/API refund and dispute surfaces.
- SePay bank transfers: usually require manual refund/reconciliation unless the gateway surface provides a refund operation.

## Recommended Tables

### Orders

Orders represent the local commercial intent. They should survive provider retries, checkout recreation, or provider migration.

```typescript
type PaymentProvider = "sepay" | "polar" | "stripe" | "paddle" | "creem";
type ProviderEnvironment = "sandbox" | "test" | "production";

type LocalOrderStatus =
  | "draft"
  | "checkout_created"
  | "awaiting_payment"
  | "paid"
  | "partially_refunded"
  | "refunded"
  | "failed"
  | "expired"
  | "cancelled"
  | "disputed";

interface LocalOrder {
  id: string;
  userId?: string;
  buyerEmail?: string;

  provider: PaymentProvider;
  providerEnvironment: ProviderEnvironment;

  localStatus: LocalOrderStatus;
  providerStatus?: string;

  amountMinor: number;
  currency: string;
  taxAmountMinor?: number;
  discountAmountMinor?: number;

  reportingAmountUsdMinor?: number;
  reportingFxRate?: string;
  reportingFxSource?: "native" | "provider" | "api" | "cached" | "manual";
  reportingFxCapturedAt?: string;

  productKey?: string;
  quantity?: number;

  createdAt: string;
  updatedAt: string;
  paidAt?: string;
  expiresAt?: string;
}
```

### Provider Payment References

Store provider IDs in a separate table or structured object. Different providers name the same concept differently.

```typescript
interface ProviderPaymentReference {
  orderId: string;
  provider: PaymentProvider;
  providerEnvironment: ProviderEnvironment;

  providerCustomerId?: string;
  providerCheckoutId?: string;
  providerCheckoutUrl?: string;
  providerOrderId?: string;
  providerTransactionId?: string;
  providerPaymentIntentId?: string;
  providerSubscriptionId?: string;
  providerProductId?: string;
  providerPriceId?: string;
  providerRefundId?: string;
  providerLicenseId?: string;

  providerRawStatus?: string;
  metadata?: Record<string, unknown>;
}
```

Examples:

- Polar checkout session ID is not the same as Polar order ID.
- Paddle transaction ID is not the same as Paddle subscription ID.
- Stripe Checkout Session ID is not the same as PaymentIntent ID.
- SePay bank transaction ID is not the same as a local order ID.
- Creem checkout, payment, subscription, and license IDs should be tracked separately.

### Provider Events

Every webhook/IPN event should be recorded before business processing.

```typescript
interface ProviderEvent {
  id: string;
  provider: PaymentProvider;
  providerEnvironment: ProviderEnvironment;

  providerEventId: string;
  eventType: string;
  eventCreatedAt?: string;

  orderId?: string;
  providerObjectId?: string;

  rawPayload: string;
  signatureVerified: boolean;
  processingStatus: "received" | "queued" | "processed" | "ignored" | "failed";
  processingError?: string;

  receivedAt: string;
  processedAt?: string;
}
```

Use a unique key such as `(provider, providerEnvironment, providerEventId)`. If a provider surface does not supply a stable event ID, derive a deterministic fingerprint from provider, environment, event type, object ID, event timestamp, and raw payload hash.

## Webhook Processing Pattern

Do not "always return 200" before authentication. First verify the webhook/IPN signature or configured auth. Then record or queue the event. Then acknowledge quickly.

```typescript
async function handleProviderWebhook(input: {
  provider: PaymentProvider;
  environment: ProviderEnvironment;
  rawBody: string;
  headers: Headers;
}) {
  const verification = await verifyProviderWebhook(input);

  if (!verification.ok) {
    return { status: 400, body: "invalid webhook" };
  }

  const event = await parseProviderEvent(input.provider, input.rawBody);

  await insertProviderEventOnce({
    provider: input.provider,
    providerEnvironment: input.environment,
    providerEventId: event.id,
    eventType: event.type,
    rawPayload: input.rawBody,
    signatureVerified: true,
    processingStatus: "received",
  });

  await enqueueProviderEvent(event.id);

  return { status: 200, body: "ok" };
}
```

Processing should be idempotent:

- Ignore duplicate event IDs after confirming the original event was recorded.
- Handle out-of-order events by fetching the provider object before making final access or revenue decisions.
- Persist failures and retry from your local event table.
- Keep raw payloads for support and audit, with PII retention rules appropriate for the app.

## Checkout and Payment Attempt Pattern

Create a local order before redirecting to a provider. Store local order ID in provider metadata when the provider supports metadata.

```typescript
async function startCheckout(params: {
  provider: PaymentProvider;
  userId: string;
  productKey: string;
  amountMinor: number;
  currency: string;
}) {
  const order = await createLocalOrder({
    provider: params.provider,
    userId: params.userId,
    productKey: params.productKey,
    amountMinor: params.amountMinor,
    currency: params.currency,
    localStatus: "draft",
  });

  const checkout = await createProviderCheckout({
    provider: params.provider,
    localOrderId: order.id,
    productKey: params.productKey,
    amountMinor: params.amountMinor,
    currency: params.currency,
  });

  await storeProviderPaymentReference(order.id, checkout);
  await markOrderStatus(order.id, "checkout_created");

  return checkout.redirectUrl;
}
```

Provider notes:

- Polar: create checkouts with product IDs through `products`; do not use `product_price_id` as the default create field.
- SePay: if using bank transfer/VietQR, payment instructions may be local plus bank webhook reconciliation; if using Payment Gateway, store gateway order/IPN IDs.
- Stripe: Checkout Session is usually the simplest web checkout surface.
- Paddle: transaction checkout state and subscription lifecycle must be tracked separately.
- Creem.io: checkout, subscription, license, and revenue-split IDs can all matter.

## Entitlement Pattern

Access state is not the same thing as order state.

Create a local entitlement table or access snapshot:

```typescript
interface LocalEntitlement {
  userId: string;
  orderId?: string;
  provider: PaymentProvider;
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

## Refunds, Disputes, and Access

Model refunds separately from orders.

```typescript
interface LocalRefund {
  id: string;
  orderId: string;
  provider: PaymentProvider;
  providerRefundId?: string;
  amountMinor: number;
  currency: string;
  reason?: string;
  status: "requested" | "succeeded" | "failed" | "cancelled";
  accessAction: "keep_access" | "revoke_access" | "manual_review";
  createdAt: string;
  completedAt?: string;
}
```

Rules:

- Do not treat a refund as subscription cancellation unless the provider or app policy explicitly cancels/revokes the subscription too.
- Decide whether full or partial refunds revoke access per product type.
- Reverse or hold commissions when refund risk remains.
- Store provider refund status and local access policy separately.
- For MoR providers, tax and fee behavior can differ by plan/provider; record provider-reported amounts instead of recomputing from hard-coded formulas.

## Discounts and Cross-Provider Campaigns

Prefer a local campaign/redemption ledger when the same promotion can be used across providers.

```typescript
interface CampaignRedemption {
  id: string;
  campaignId: string;
  provider: PaymentProvider;
  orderId: string;
  providerDiscountId?: string;
  code?: string;
  amountMinor: number;
  currency: string;
  status: "reserved" | "redeemed" | "released" | "failed";
  redeemedAt?: string;
}
```

Avoid decrementing or deleting a provider-native discount from another provider's checkout unless:

- the app owns that cross-provider campaign;
- the provider API supports the intended operation;
- local idempotency and rollback behavior are defined;
- reconciliation can recover from partial sync failures.

For Polar discounts, load `references/polar/orders-refunds-discounts.md` and `references/polar/checkouts.md`. Use provider-native discounts for provider-native checkout behavior, and use the local campaign ledger for cross-provider limits.

## Currency and Revenue Reporting

Store money in the provider's currency and minor unit. Normalize only for reporting.

Recommended fields:

- `amount_minor`
- `currency`
- `tax_amount_minor`
- `discount_amount_minor`
- `provider_fee_minor`
- `net_amount_minor`
- `reporting_currency`
- `reporting_amount_minor`
- `fx_rate`
- `fx_source`
- `fx_captured_at`

Rules:

- Do not hard-code provider fees as universal truth. Fee schedules and plans change.
- Use provider-reported fees/net amounts when available.
- If provider fee data is not available, mark the value as estimated and keep the formula/version used.
- For MoR providers, separate tax, fee, gross revenue, net payout, and refund amounts.
- For bank transfer providers, record bank amount and any manual operational fee separately.

## Commissions and Revenue Splits

Commission systems should depend on stable local facts, not transient checkout state.

Recommended pattern:

1. Create commission in `pending` state after local order becomes `paid`.
2. Approve commission after refund/dispute risk window or provider payout rules allow it.
3. Reverse or reduce commission on full/partial refunds.
4. Recalculate tier/revenue metrics from immutable paid/refund facts, not by incrementing counters only.
5. Keep provider-native revenue splits separate from local affiliate commissions.

For Creem.io revenue splits, load `references/creem/overview.md` and provider-specific API docs before modeling payout ownership.

## Reconciliation Jobs

Run scheduled reconciliation. Webhooks are necessary but not sufficient.

Useful checks:

- Local `checkout_created` orders older than the expected checkout/session lifetime.
- Local `awaiting_payment` orders with a provider paid object.
- Provider paid orders missing local fulfillment.
- Refunds or disputes missing local access/commission updates.
- Subscription states that changed while webhook delivery was down.
- Unprocessed provider events older than the normal queue latency.
- Amount/currency mismatches between local order and provider object.

Reconciliation should fetch provider object state and update local records through the same idempotent handlers used by webhooks.

## Admin and Support API Pattern

Expose provider details to admins without forcing every caller to understand every provider.

```typescript
interface AdminOrderView {
  id: string;
  provider: PaymentProvider;
  localStatus: LocalOrderStatus;
  providerStatus?: string;
  amountMinor: number;
  currency: string;
  reportingAmountUsdMinor?: number;
  providerReferenceSummary: {
    customerId?: string;
    checkoutId?: string;
    orderId?: string;
    transactionId?: string;
    subscriptionId?: string;
    refundId?: string;
  };
  lastProviderEvent?: {
    type: string;
    receivedAt: string;
    processingStatus: string;
  };
}
```

Support tools should allow:

- filter by provider, environment, local status, provider status, user, email, or provider object ID;
- view raw provider IDs and event history;
- replay failed local processing after fixing app errors;
- trigger read-only provider reconciliation;
- record manual refund or access decisions with audit notes.

## Best Practices Summary

1. Create local orders before provider checkout.
2. Store provider IDs separately and precisely.
3. Verify webhooks/IPNs before acknowledging them.
4. Record every provider event before business processing.
5. Fulfill from provider paid signals, not customer redirects.
6. Keep order, subscription, refund, dispute, and entitlement state separate.
7. Normalize currency for reporting only; keep original provider amounts.
8. Use provider-reported fees/net amounts when possible.
9. Use a local campaign ledger for cross-provider discounts.
10. Reconcile provider state on a schedule and after webhook outages.
