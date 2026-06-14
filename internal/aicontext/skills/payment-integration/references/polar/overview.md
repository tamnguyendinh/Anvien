# Polar Overview

Polar is a Merchant of Record payment and billing platform for software, SaaS, subscriptions, usage-based billing, digital products, license keys, and automated benefits.

Use this file first to choose the right Polar surface. Load only the detailed references needed for the task.

## Source Links

- Docs index: https://polar.sh/docs/llms.txt
- API overview: https://polar.sh/docs/api-reference/introduction
- Products: https://polar.sh/docs/features/products
- Checkout API: https://polar.sh/docs/features/checkout/session
- Webhooks: https://polar.sh/docs/integrate/webhooks/endpoints
- Customer Portal: https://polar.sh/docs/features/customer-portal/introduction
- Fees: https://polar.sh/docs/merchant-of-record/fees

## Surface Selection

| Need | Load |
|------|------|
| Product catalog, pricing model, currencies, tax behavior, custom fields | `products.md` |
| Hosted checkout, embedded checkout, checkout sessions, checkout links | `checkouts.md` |
| Recurring plan changes, trials, cancellation, revoke, renewal handling | `subscriptions.md` |
| Event delivery, signature verification, fulfillment, idempotency | `webhooks.md` |
| License keys, feature flags, credits, Discord/GitHub/file benefits | `benefits.md` |
| Usage events, meters, metered prices, credits, customer meter balance | `usage-based-billing.md` |
| Customer self-service portal, customer sessions, portal API boundaries | `customer-portal.md` |
| Fast entitlement sync and access decisions from one customer state object | `customer-state.md` |
| Orders, paid signals, invoices, receipts, refunds, discount behavior | `orders-refunds-discounts.md` |
| TypeScript SDK, framework adapters, BetterAuth, Laravel, Express | `sdk.md` |
| Production readiness, fees, retries, reconciliation, storage patterns | `best-practices.md` |

## Core Capabilities

- Merchant of Record: Polar handles checkout, tax calculation/remittance, invoices, receipts, refunds, and payouts.
- Products: one-time and recurring products share one product model.
- Checkout: dashboard checkout links, API checkout sessions, and embedded checkout.
- Subscriptions: trials, plan changes, proration, cancellation, revoke, dunning, and renewals.
- Usage billing: ingest immutable events, define meters, attach metered prices, and optionally issue Credits benefits.
- Benefits: license keys, credits, feature flags, file downloads, GitHub access, Discord access, and custom benefits.
- Customer Portal: hosted self-service for subscriptions, receipts, invoices, benefits, payment methods, seats, and metered usage.
- Orders/refunds/discounts: paid transaction records, invoices/receipts, refund operations, and promotional discounts.
- Webhooks: Standard Webhooks-compatible event delivery with SDK/adapters for validation.

## Authentication

### Organization Access Tokens

Use Organization Access Tokens for server-side Core API calls. Never expose them in browser code.

```http
Authorization: Bearer $POLAR_ACCESS_TOKEN
Accept: application/json
```

Create tokens separately in each environment. Production tokens do not work in sandbox, and sandbox tokens do not work in production.

### OAuth 2.0

Use OAuth 2.0 for third-party apps that act on behalf of Polar organizations.

- Authorization URL: `https://polar.sh/oauth2/authorize`
- Token URL: `https://api.polar.sh/v1/oauth2/token`
- Common scopes: `products:read`, `products:write`, `checkouts:write`, `orders:read`, `subscriptions:read`, `subscriptions:write`, `benefits:read`, `benefits:write`, `customers:read`, `customers:write`, `refunds:write`

### Customer Sessions and Customer Portal API

Use customer sessions to create authenticated portal links. Customer Portal API tokens are scoped to one customer and must not be treated like organization access tokens.

Load `customer-portal.md` for portal session fields and customer-facing API boundaries.

## Base URLs

| Environment | Dashboard | Core API |
|-------------|-----------|----------|
| Production | `https://polar.sh` | `https://api.polar.sh/v1` |
| Sandbox | `https://sandbox.polar.sh` | `https://sandbox-api.polar.sh/v1` |

TypeScript SDK:

```typescript
import { Polar } from "@polar-sh/sdk";

export const polar = new Polar({
  accessToken: process.env.POLAR_ACCESS_TOKEN,
  server: process.env.POLAR_SERVER === "production" ? "production" : "sandbox",
});
```

## Rate Limits

Official API limits can change. Verify the API overview before hard-coding a policy.

Current documented limits:

- Production Core API: 500 requests per minute per organization/customer or OAuth2 client.
- Sandbox Core API: 100 requests per minute per organization/customer or OAuth2 client.
- Unauthenticated license validation, activation, and deactivation: 3 requests per second.
- Exceeded limits return HTTP `429` with `Retry-After`.

```typescript
async function sleepForRateLimit(response: Response) {
  if (response.status !== 429) return;
  const retryAfter = Number(response.headers.get("Retry-After") ?? "1");
  await new Promise((resolve) => setTimeout(resolve, retryAfter * 1000));
}
```

## Fees

Do not hard-code one fee model as universal. Current official fee docs distinguish:

- Starter: free monthly plan with a higher per-transaction fee.
- Pro, Growth, Scale: monthly plans with lower per-transaction fees.
- Early Member: applies only to organizations created before May 27, 2026 that remain on that plan.
- Additional international card fee may apply.
- Separate subscription fee applies to Early Member only, not Starter/Pro/Growth/Scale.

Load `best-practices.md` before implementing fee estimates, and link users to the official fee page when showing calculators.

## Key Data Concepts

### Product IDs, Not Price IDs for New Checkout Creation

Current checkout session creation uses product IDs:

```typescript
await polar.checkouts.create({
  products: [process.env.POLAR_PRODUCT_ID!],
  successUrl: "https://example.com/success?checkout_id={CHECKOUT_ID}",
});
```

`product_price_id` may still appear in legacy objects or deprecated response fields. Do not use it as the default create shape for new work.

### External Customer ID

Use `external_customer_id` in REST or `externalCustomerId` in the TypeScript SDK to map Polar customers to application users.

- Set it at checkout or customer creation.
- Use stable internal user IDs.
- Treat it as a reconciliation key.
- Do not use mutable emails as primary identity.

### Metadata

Metadata is for audit and reconciliation fields. Store small scalar values only. Do not put secrets, full payload archives, or PII-heavy objects in metadata.

### Customer State

Customer State is often the fastest entitlement source. It includes customer data, active subscriptions, granted benefits, and active meter balances. Subscribe to `customer.state_changed` to keep local access state current.

## Environments

Sandbox is isolated:

- Separate organization and dashboard.
- Separate products, prices, discounts, webhooks, and tokens.
- Stripe test cards work in sandbox.
- Verify webhook events, checkout flows, subscriptions, refunds, customer portal, and usage billing in sandbox before production.
- Sandbox customer-facing emails are limited; use organization-member test addresses when validating emails.

## Merchant of Record Launch Checks

Before production launch, verify official docs for:

- Supported payout country and Stripe Connect Express onboarding.
- Acceptable use policy for the product category.
- Account review requirements.
- Current fee plan.
- Tax behavior and customer-facing price display.
- Refund policy and support workflow.

## Next Steps

- Products and pricing: `products.md`
- Checkout sessions/links/embedded checkout: `checkouts.md`
- Subscriptions and plan changes: `subscriptions.md`
- Webhooks and event handling: `webhooks.md`
- Benefits and entitlements: `benefits.md`
- Usage-based billing: `usage-based-billing.md`
- Customer Portal: `customer-portal.md`
- Customer State: `customer-state.md`
- Orders/refunds/discounts: `orders-refunds-discounts.md`
- SDKs/adapters: `sdk.md`
- Production readiness: `best-practices.md`
