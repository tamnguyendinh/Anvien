# Revenue, Commissions, and Reconciliation Template

Use this template when the app needs cross-provider reporting, fee/net tracking, revenue splits, commissions, or scheduled reconciliation.

## Currency and Reporting

Store money in the provider's currency and minor unit. Normalize only for reporting.

Recommended fields:

- `amount_minor`
- `currency`
- `tax_amount_minor`
- `discount_amount_minor`
- `provider_fee_minor`
- `provider_net_amount_minor`
- `reporting_currency`
- `reporting_amount_minor`
- `fx_rate`
- `fx_source`
- `fx_captured_at`

Rules:

- Do not hard-code provider fees as universal truth.
- Use provider-reported fees/net amounts when available.
- If provider fee data is not available, mark the value as estimated and keep the formula/version used.
- For MoR providers, separate tax, fee, gross revenue, net payout, and refund amounts.
- For bank transfer providers, record bank amount and any manual operational fee separately.

## FX Rate Example

```typescript
interface FxRate {
  base: string;
  quote: string;
  rate: string;
  source: "provider" | "api" | "cached" | "manual";
  capturedAt: string;
}

const DEFAULT_REPORTING_CURRENCY = "USD";

async function normalizeForReporting(params: {
  amountMinor: number;
  currency: string;
  provider?: PaymentProvider;
  providerFxRate?: FxRate;
}) {
  if (params.currency === DEFAULT_REPORTING_CURRENCY) {
    return {
      reportingCurrency: DEFAULT_REPORTING_CURRENCY,
      reportingAmountMinor: params.amountMinor,
      fxRate: "1",
      fxSource: "provider",
      fxCapturedAt: new Date().toISOString(),
    };
  }

  const fx =
    params.providerFxRate ??
    (await getFreshFxRate(params.currency, DEFAULT_REPORTING_CURRENCY)) ??
    (await getCachedFxRate(params.currency, DEFAULT_REPORTING_CURRENCY));

  if (!fx) throw new Error("missing_fx_rate");

  return {
    reportingCurrency: DEFAULT_REPORTING_CURRENCY,
    reportingAmountMinor: convertMinorUnits(params.amountMinor, fx.rate),
    fxRate: fx.rate,
    fxSource: fx.source,
    fxCapturedAt: fx.capturedAt,
  };
}
```

## Revenue Summary Example

```typescript
async function getRevenueSummary(params: {
  from: Date;
  to: Date;
  provider?: PaymentProvider;
}) {
  const rows = await listPaidOrdersAndRefunds(params);

  const summary = {
    grossMinor: 0,
    refundsMinor: 0,
    providerFeesMinor: 0,
    netMinor: 0,
    reportingCurrency: DEFAULT_REPORTING_CURRENCY,
    byProvider: {} as Record<PaymentProvider, { grossMinor: number; netMinor: number }>,
  };

  for (const row of rows) {
    const provider = row.provider;
    summary.byProvider[provider] ??= { grossMinor: 0, netMinor: 0 };
    summary.grossMinor += row.reportingGrossMinor;
    summary.refundsMinor += row.reportingRefundMinor;
    summary.providerFeesMinor += row.reportingProviderFeeMinor;
    summary.netMinor += row.reportingNetMinor;
    summary.byProvider[provider].grossMinor += row.reportingGrossMinor;
    summary.byProvider[provider].netMinor += row.reportingNetMinor;
  }

  return summary;
}
```

## Commission Table

```typescript
export const commissions = pgTable("commissions", {
  id: uuid("id").primaryKey().defaultRandom(),
  orderId: uuid("order_id").notNull(),
  campaignId: text("campaign_id"),
  recipientUserId: uuid("recipient_user_id").notNull(),
  baseAmountMinor: integer("base_amount_minor").notNull(),
  baseCurrency: text("base_currency").notNull(),
  rate: numeric("rate", { precision: 8, scale: 6 }).notNull(),
  commissionAmountMinor: integer("commission_amount_minor").notNull(),
  commissionCurrency: text("commission_currency").notNull(),
  reportingAmountMinor: integer("reporting_amount_minor"),
  reportingCurrency: text("reporting_currency"),
  status: text("status").notNull().default("pending"),
  holdReason: text("hold_reason"),
  approvedAt: timestamp("approved_at"),
  reversedAt: timestamp("reversed_at"),
  paidAt: timestamp("paid_at"),
  createdAt: timestamp("created_at").notNull().defaultNow(),
});
```

## Commission Creation and Reversal

```typescript
async function createPendingCommissionsForPaidOrder(orderId: string) {
  const order = await getOrder(orderId);
  const commissionRules = await resolveCommissionRules(order);

  for (const rule of commissionRules) {
    const baseAmountMinor = chooseCommissionBaseAmount(order, rule);
    const commissionAmountMinor = Math.round(baseAmountMinor * rule.rate);
    const reporting = await normalizeForReporting({
      amountMinor: commissionAmountMinor,
      currency: order.currency,
      provider: order.provider,
    });

    await insertCommissionOnce({
      orderId: order.id,
      campaignId: rule.campaignId,
      recipientUserId: rule.recipientUserId,
      baseAmountMinor,
      baseCurrency: order.currency,
      rate: String(rule.rate),
      commissionAmountMinor,
      commissionCurrency: order.currency,
      reportingAmountMinor: reporting.reportingAmountMinor,
      reportingCurrency: reporting.reportingCurrency,
      status: rule.requiresApproval ? "pending" : "approved",
    });
  }
}

async function reverseCommissionsForOrder(
  orderId: string,
  context: { reason: "refund" | "dispute" | "manual"; refundId?: string },
) {
  const commissions = await listCommissionsForOrder(orderId);

  for (const commission of commissions) {
    if (commission.status === "paid") {
      await createCommissionAdjustment({
        commissionId: commission.id,
        amountMinor: -commission.commissionAmountMinor,
        reason: context.reason,
        metadata: { refundId: context.refundId },
      });
    } else {
      await markCommissionReversed(commission.id, context.reason);
    }
  }
}
```

## Reconciliation Jobs

Webhooks are necessary but not sufficient. Reconcile stale checkouts, paid provider objects without local fulfillment, refunds/disputes without consequences, subscription drift, failed events, and amount/currency mismatches.

```typescript
async function reconcileStaleOrders() {
  const staleOrders = await listOrdersForReconciliation({
    statuses: ["checkout_created", "awaiting_payment"],
    olderThanMinutes: 30,
  });

  for (const order of staleOrders) {
    try {
      const providerState = await fetchProviderOrderState(order);

      if (providerState.paid && providerState.paidFacts) {
        await handlePaidProviderFacts(providerState.paidFacts);
        continue;
      }

      if (providerState.expired) {
        await markOrderExpired(order.id, providerState.providerStatus);
        await releaseCampaignReservation(order.id);
        continue;
      }

      if (providerState.failed) {
        await markOrderFailed(order.id, providerState.providerStatus);
        await releaseCampaignReservation(order.id);
      }
    } catch (error) {
      await recordReconciliationFailure(order.id, serializeError(error));
    }
  }
}
```

```typescript
async function retryFailedProviderEvents() {
  const failedEvents = await listProviderEvents({ statuses: ["failed"], limit: 100 });

  for (const event of failedEvents) {
    await markProviderEventQueued(event.id);
    await enqueueProviderEvent({
      provider: event.provider,
      providerEnvironment: event.providerEnvironment,
      providerEventId: event.providerEventId,
    });
  }
}
```
