# Refunds, Disputes, and Campaigns Template

Use this template when refunds, chargebacks, access changes, commissions, or cross-provider discounts must be handled consistently across providers.

## Refund Command

Model refunds separately from orders. A refund has provider state and local business consequences.

```typescript
interface RefundCommand {
  orderId: string;
  amountMinor?: number;
  reason?: string;
  accessAction: "keep_access" | "revoke_access" | "manual_review";
  commissionAction: "keep_commission" | "reverse_commission" | "manual_review";
}

interface RefundResult {
  providerRefundId?: string;
  providerRawStatus?: string;
  amountMinor: number;
  currency: string;
  raw: unknown;
}
```

## Refund Request Flow

```typescript
async function requestRefund(command: RefundCommand) {
  const order = await getOrder(command.orderId);

  if (!canRequestRefund(order)) {
    throw new Error("order_not_refundable");
  }

  const refund = await createLocalRefund({
    orderId: order.id,
    provider: order.provider,
    providerEnvironment: order.providerEnvironment,
    amountMinor: command.amountMinor ?? refundableAmount(order),
    currency: order.currency,
    reason: command.reason,
    accessAction: command.accessAction,
    commissionAction: command.commissionAction,
    localStatus: "requested",
  });

  if (order.provider === "sepay" && usesManualBankRefund(order)) {
    await flagRefundForManualBankTransfer(refund.id);
    return refund;
  }

  const result = await providerRefundAdapters[order.provider].refund({ order, refund });
  await updateRefundFromProvider(refund.id, result);

  if (providerRefundAlreadySucceeded(result)) {
    await applyRefundConsequences(refund.id);
  }

  return refund;
}
```

## Refund Consequences

```typescript
async function applyRefundConsequences(refundId: string) {
  const refund = await getRefund(refundId);

  await updateOrderRefundStatus(refund.orderId);

  if (refund.accessAction === "revoke_access") {
    await revokeEntitlementsForOrder(refund.orderId, {
      reason: "refund",
      refundId,
    });
  }

  if (refund.commissionAction === "reverse_commission") {
    await reverseCommissionsForOrder(refund.orderId, {
      reason: "refund",
      refundId,
    });
  }
}
```

Rules:

- Do not treat a refund as subscription cancellation unless the provider or app policy explicitly cancels/revokes the subscription too.
- Decide whether full or partial refunds revoke access per product type.
- Reverse or hold commissions when refund risk remains.
- Store provider refund status and local access policy separately.
- For MoR providers, record provider-reported tax, fee, gross, net, and refund amounts instead of recomputing from hard-coded formulas.

## Dispute Flow

```typescript
async function handleDisputeProviderEvent(event: ProviderEvent) {
  const facts = await extractDisputeFacts(event);
  const order = await findOrderByProviderObject(facts);

  if (!order) {
    await flagProviderEventForManualReview(event.id, "dispute_without_order");
    return;
  }

  await markOrderDisputed(order.id, {
    providerStatus: facts.providerStatus,
    disputedAt: facts.disputedAt,
    amountMinor: facts.amountMinor,
    currency: facts.currency,
  });

  await holdCommissionsForOrder(order.id, "provider_dispute");

  if (facts.riskRequiresAccessReview) {
    await flagEntitlementsForManualReview(order.id, "provider_dispute");
  }
}
```

## Cross-Provider Campaign Ledger

Prefer a local campaign/redemption ledger when the same promotion can be used across providers.

```typescript
interface CampaignRedemption {
  id: string;
  campaignId: string;
  provider: PaymentProvider;
  providerEnvironment: ProviderEnvironment;
  orderId: string;
  providerDiscountId?: string;
  code?: string;
  amountMinor: number;
  currency: string;
  status: "reserved" | "redeemed" | "released" | "failed";
  redeemedAt?: string;
}
```

## Reserve, Redeem, Release

Use a reservation flow so failed checkouts do not consume cross-provider campaign limits forever.

```typescript
async function reserveCampaignRedemption(params: {
  campaignId: string;
  orderId: string;
  provider: PaymentProvider;
  providerEnvironment: ProviderEnvironment;
  code?: string;
  amountMinor: number;
  currency: string;
}) {
  const campaign = await getCampaignForUpdate(params.campaignId);

  if (!campaign.active) throw new Error("campaign_inactive");
  if (campaign.remainingRedemptions <= 0) throw new Error("campaign_exhausted");

  await insertCampaignRedemption({ ...params, status: "reserved" });
  await decrementCampaignRemaining(params.campaignId);
}

async function markCampaignRedeemed(orderId: string) {
  await updateCampaignRedemptionByOrder(orderId, {
    status: "redeemed",
    redeemedAt: new Date(),
  });
}

async function releaseCampaignReservation(orderId: string) {
  const redemption = await getReservedCampaignRedemption(orderId);
  if (!redemption) return;

  await updateCampaignRedemption(redemption.id, {
    status: "released",
    releasedAt: new Date(),
  });

  await incrementCampaignRemaining(redemption.campaignId);
}
```

Where to call it:

- Reserve before provider checkout if a limited campaign applies.
- Mark redeemed after provider paid signal.
- Release after checkout expiration, failed payment, or support cancellation.
- Reconcile old `reserved` rows linked to stale `checkout_created` or `awaiting_payment` orders.

Avoid decrementing or deleting a provider-native discount from another provider's checkout unless the app owns that cross-provider campaign, the provider API supports the operation, local idempotency is defined, and reconciliation can recover from partial failures.
