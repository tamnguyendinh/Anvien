# Admin, Support, and App Variants Template

Use this template when building admin views, manual support actions, or adapting the multi-provider model to a specific type of app.

## Admin Order View

Expose provider details to admins without forcing every caller to understand every provider.

```typescript
interface AdminOrderView {
  id: string;
  provider: PaymentProvider;
  localStatus: LocalOrderStatus;
  providerStatus?: string;
  amountMinor: number;
  currency: string;
  reportingAmountMinor?: number;
  reportingCurrency?: string;
  providerReferenceSummary: {
    customerId?: string;
    checkoutId?: string;
    orderId?: string;
    transactionId?: string;
    paymentIntentId?: string;
    subscriptionId?: string;
    refundId?: string;
    licenseId?: string;
  };
  lastProviderEvent?: {
    type: string;
    receivedAt: string;
    processingStatus: string;
  };
}
```

## Admin List Endpoint

```typescript
export async function listAdminOrders(query: {
  provider?: PaymentProvider;
  environment?: ProviderEnvironment;
  localStatus?: LocalOrderStatus;
  providerStatus?: string;
  email?: string;
  providerObjectId?: string;
  limit?: number;
  cursor?: string;
}): Promise<{ orders: AdminOrderView[]; nextCursor?: string }> {
  const rows = await queryOrdersWithProviderReferences(query);

  return {
    orders: rows.map((row) => ({
      id: row.order.id,
      provider: row.order.provider,
      localStatus: row.order.localStatus,
      providerStatus: row.order.providerStatus,
      amountMinor: row.order.amountMinor,
      currency: row.order.currency,
      reportingAmountMinor: row.order.reportingAmountMinor,
      reportingCurrency: row.order.reportingCurrency,
      providerReferenceSummary: {
        customerId: row.refs.providerCustomerId,
        checkoutId: row.refs.providerCheckoutId,
        orderId: row.refs.providerOrderId,
        transactionId: row.refs.providerTransactionId,
        paymentIntentId: row.refs.providerPaymentIntentId,
        subscriptionId: row.refs.providerSubscriptionId,
        refundId: row.refs.providerRefundId,
        licenseId: row.refs.providerLicenseId,
      },
      lastProviderEvent: row.lastEvent
        ? {
            type: row.lastEvent.eventType,
            receivedAt: row.lastEvent.receivedAt,
            processingStatus: row.lastEvent.processingStatus,
          }
        : undefined,
    })),
    nextCursor: rows.at(-1)?.cursor,
  };
}
```

## Support Actions

```typescript
type SupportAction =
  | "replay_event"
  | "reconcile_order"
  | "mark_manual_refund"
  | "grant_manual_entitlement"
  | "revoke_entitlement"
  | "release_campaign_reservation";

async function runSupportAction(params: {
  actorId: string;
  action: SupportAction;
  orderId?: string;
  providerEventId?: string;
  reason: string;
}) {
  await insertAuditLog({
    actorId: params.actorId,
    action: params.action,
    targetOrderId: params.orderId,
    targetProviderEventId: params.providerEventId,
    reason: params.reason,
  });

  switch (params.action) {
    case "replay_event":
      return replayProviderEvent(params.providerEventId!);
    case "reconcile_order":
      return reconcileOneOrder(params.orderId!);
    case "mark_manual_refund":
      return markManualRefund(params.orderId!, params.reason);
    case "grant_manual_entitlement":
      return grantManualEntitlement(params.orderId!, params.reason);
    case "revoke_entitlement":
      return revokeEntitlementForOrder(params.orderId!, params.reason);
    case "release_campaign_reservation":
      return releaseCampaignReservation(params.orderId!);
  }
}
```

Support tools should allow:

- filter by provider, environment, local status, provider status, user, email, or provider object ID;
- view raw provider IDs and event history;
- replay failed local processing after fixing app errors;
- trigger read-only provider reconciliation;
- record manual refund or access decisions with audit notes.

## Example App Variants

Use these as patterns, not as fixed product assumptions.

### SaaS With Global MoR and Local Bank Transfer

- Providers: Polar or Paddle for cards/subscriptions, SePay for Vietnamese bank transfers.
- Local order status: `awaiting_payment` until bank webhook or MoR paid event arrives.
- Entitlements: subscription feature flags, usage credits, or team seats.
- Reconciliation: compare MoR paid orders and SePay bank transactions against local orders.

### Digital Downloads

- Providers: Stripe Checkout, Polar checkout, Creem.io checkout, or Paddle transaction checkout.
- Local order status: `paid` grants file/download entitlement.
- Refund policy: full refund may revoke file access; partial refund may keep access.
- Admin support: resend receipt, regenerate download link, reconcile provider order.

### Licensed Desktop or API Product

- Providers: Creem.io licensing, Polar License Keys, Stripe/Paddle payments with app-owned licenses.
- Local entitlement: `license:activation`, `api:quota`, or `app:pro_features`.
- Paid event creates or activates license; refund/dispute can revoke or flag for review.
- Reconciliation: ensure provider subscription/license state matches local license state.

### Marketplace or Revenue Split App

- Providers: Stripe Connect or Creem.io revenue splits, plus local affiliate commissions.
- Keep provider-native splits separate from app-owned commission tables.
- Commission approval should wait for refund/dispute windows where required.
- Revenue reports should separate gross, provider fees, taxes, split payouts, local commissions, and net app revenue.
