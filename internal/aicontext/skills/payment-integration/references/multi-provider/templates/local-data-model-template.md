# Local Data Model Template

Use this template when designing local tables for more than one payment provider. Rename fields to match the target repo's ORM, but keep the ownership boundaries.

## Core Types

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
```

## Orders

Orders represent local commercial intent. They should survive provider retries, checkout recreation, provider migration, or support actions.

```typescript
interface LocalOrder {
  id: string;
  userId?: string;
  buyerEmail?: string;
  provider: PaymentProvider;
  providerEnvironment: ProviderEnvironment;
  localStatus: LocalOrderStatus;
  providerStatus?: string;
  productKey?: string;
  quantity?: number;
  amountMinor: number;
  currency: string;
  taxAmountMinor?: number;
  discountAmountMinor?: number;
  providerFeeMinor?: number;
  providerNetAmountMinor?: number;
  reportingCurrency?: string;
  reportingAmountMinor?: number;
  reportingFxRate?: string;
  reportingFxSource?: "provider" | "api" | "cached" | "manual";
  reportingFxCapturedAt?: string;
  paidAt?: string;
  expiresAt?: string;
  metadata?: Record<string, unknown>;
}
```

Drizzle-style starter:

```typescript
export const orders = pgTable("orders", {
  id: uuid("id").primaryKey().defaultRandom(),
  userId: uuid("user_id"),
  buyerEmail: text("buyer_email"),
  provider: paymentProvider("provider").notNull(),
  providerEnvironment: providerEnvironment("provider_environment").notNull(),
  localStatus: localOrderStatus("local_status").notNull().default("draft"),
  providerStatus: text("provider_status"),
  productKey: text("product_key"),
  quantity: integer("quantity").notNull().default(1),
  amountMinor: integer("amount_minor").notNull(),
  currency: text("currency").notNull(),
  taxAmountMinor: integer("tax_amount_minor"),
  discountAmountMinor: integer("discount_amount_minor"),
  providerFeeMinor: integer("provider_fee_minor"),
  providerNetAmountMinor: integer("provider_net_amount_minor"),
  reportingCurrency: text("reporting_currency"),
  reportingAmountMinor: integer("reporting_amount_minor"),
  reportingFxRate: numeric("reporting_fx_rate", { precision: 18, scale: 8 }),
  reportingFxSource: text("reporting_fx_source"),
  reportingFxCapturedAt: timestamp("reporting_fx_captured_at"),
  metadata: jsonb("metadata").$type<Record<string, unknown>>().default({}),
  createdAt: timestamp("created_at").notNull().defaultNow(),
  updatedAt: timestamp("updated_at").notNull().defaultNow(),
  paidAt: timestamp("paid_at"),
  expiresAt: timestamp("expires_at"),
});
```

## Provider Payment References

Store provider IDs separately. Different providers name the same commercial object differently.

```typescript
export const providerPaymentReferences = pgTable(
  "provider_payment_references",
  {
    id: uuid("id").primaryKey().defaultRandom(),
    orderId: uuid("order_id").notNull(),
    provider: paymentProvider("provider").notNull(),
    providerEnvironment: providerEnvironment("provider_environment").notNull(),
    providerCustomerId: text("provider_customer_id"),
    providerCheckoutId: text("provider_checkout_id"),
    providerCheckoutUrl: text("provider_checkout_url"),
    providerOrderId: text("provider_order_id"),
    providerTransactionId: text("provider_transaction_id"),
    providerPaymentIntentId: text("provider_payment_intent_id"),
    providerSubscriptionId: text("provider_subscription_id"),
    providerProductId: text("provider_product_id"),
    providerPriceId: text("provider_price_id"),
    providerRefundId: text("provider_refund_id"),
    providerLicenseId: text("provider_license_id"),
    providerRawStatus: text("provider_raw_status"),
    metadata: jsonb("metadata").$type<Record<string, unknown>>().default({}),
    createdAt: timestamp("created_at").notNull().defaultNow(),
    updatedAt: timestamp("updated_at").notNull().defaultNow(),
  },
  (table) => ({
    providerObjectUnique: uniqueIndex("provider_payment_refs_provider_object_uq").on(
      table.provider,
      table.providerEnvironment,
      table.providerCheckoutId,
      table.providerOrderId,
      table.providerTransactionId,
      table.providerPaymentIntentId,
      table.providerSubscriptionId,
    ),
  }),
);
```

Provider ID examples:

- Polar checkout ID is not the same as Polar order ID.
- Paddle transaction ID is not the same as Paddle subscription ID.
- Stripe Checkout Session ID is not the same as PaymentIntent ID.
- SePay bank transaction ID is not the same as a local order ID.
- Creem checkout, payment, subscription, and license IDs can all matter.

## Provider Events

Record every verified webhook/IPN event before business processing.

```typescript
export const providerEvents = pgTable(
  "provider_events",
  {
    id: uuid("id").primaryKey().defaultRandom(),
    provider: paymentProvider("provider").notNull(),
    providerEnvironment: providerEnvironment("provider_environment").notNull(),
    providerEventId: text("provider_event_id").notNull(),
    eventType: text("event_type").notNull(),
    eventCreatedAt: timestamp("event_created_at"),
    orderId: uuid("order_id"),
    providerObjectId: text("provider_object_id"),
    rawPayload: text("raw_payload").notNull(),
    signatureVerified: boolean("signature_verified").notNull().default(false),
    processingStatus: text("processing_status").notNull().default("received"),
    processingError: text("processing_error"),
    receivedAt: timestamp("received_at").notNull().defaultNow(),
    processedAt: timestamp("processed_at"),
  },
  (table) => ({
    eventOnce: uniqueIndex("provider_events_once_uq").on(
      table.provider,
      table.providerEnvironment,
      table.providerEventId,
    ),
  }),
);
```

If a provider surface does not supply a stable event ID, derive a deterministic fingerprint from provider, environment, event type, object ID, event timestamp, and raw payload hash.

## Refunds, Entitlements, Campaigns

```typescript
export const refunds = pgTable("refunds", {
  id: uuid("id").primaryKey().defaultRandom(),
  orderId: uuid("order_id").notNull(),
  provider: paymentProvider("provider").notNull(),
  providerEnvironment: providerEnvironment("provider_environment").notNull(),
  providerRefundId: text("provider_refund_id"),
  providerRawStatus: text("provider_raw_status"),
  amountMinor: integer("amount_minor").notNull(),
  currency: text("currency").notNull(),
  reason: text("reason"),
  localStatus: text("local_status").notNull().default("requested"),
  accessAction: text("access_action").notNull().default("manual_review"),
  commissionAction: text("commission_action").notNull().default("manual_review"),
  metadata: jsonb("metadata").$type<Record<string, unknown>>().default({}),
  createdAt: timestamp("created_at").notNull().defaultNow(),
  completedAt: timestamp("completed_at"),
});
```

```typescript
export const entitlements = pgTable("entitlements", {
  id: uuid("id").primaryKey().defaultRandom(),
  userId: uuid("user_id").notNull(),
  orderId: uuid("order_id"),
  provider: paymentProvider("provider").notNull(),
  providerEnvironment: providerEnvironment("provider_environment").notNull(),
  sourceObjectType: text("source_object_type").notNull(),
  sourceObjectId: text("source_object_id").notNull(),
  entitlementKey: text("entitlement_key").notNull(),
  status: text("status").notNull().default("active"),
  quantity: integer("quantity"),
  startsAt: timestamp("starts_at"),
  endsAt: timestamp("ends_at"),
  revokedAt: timestamp("revoked_at"),
  metadata: jsonb("metadata").$type<Record<string, unknown>>().default({}),
  updatedAt: timestamp("updated_at").notNull().defaultNow(),
});
```

```typescript
export const campaignRedemptions = pgTable("campaign_redemptions", {
  id: uuid("id").primaryKey().defaultRandom(),
  campaignId: text("campaign_id").notNull(),
  orderId: uuid("order_id").notNull(),
  provider: paymentProvider("provider").notNull(),
  providerEnvironment: providerEnvironment("provider_environment").notNull(),
  code: text("code"),
  providerDiscountId: text("provider_discount_id"),
  amountMinor: integer("amount_minor").notNull().default(0),
  currency: text("currency").notNull(),
  status: text("status").notNull().default("reserved"),
  reservedAt: timestamp("reserved_at").notNull().defaultNow(),
  redeemedAt: timestamp("redeemed_at"),
  releasedAt: timestamp("released_at"),
});
```
