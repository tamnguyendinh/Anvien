# Polar Skill Refresh Analysis

Date: 2026-06-14
Scope: `internal/aicontext/skills/payment-integration/references/polar`
Status: discussion report only; no Polar skill/code changes made in this report.

## Problem Summary

The current Polar references in the `payment-integration` skill are directionally correct, but several details are now stale against official Polar documentation. The most important risk is that agents may implement older API shapes, especially around checkout creation, subscription updates, webhook validation, usage-based billing, and fee calculation.

The skill does not need a full rewrite, but it should be refreshed before being used as implementation guidance for new Polar integrations.

## Local Files Reviewed

- `references/polar/overview.md`
- `references/polar/checkouts.md`
- `references/polar/webhooks.md`
- `references/polar/sdk.md`
- `references/polar/products.md`
- `references/polar/subscriptions.md`
- `references/polar/benefits.md`
- `references/polar/best-practices.md`

## Official Sources Reviewed

- Polar documentation index: https://polar.sh/docs/llms.txt
- API overview: https://polar.sh/docs/api-reference/introduction
- Create checkout session: https://polar.sh/docs/api-reference/checkouts/create-session
- Checkout API: https://polar.sh/docs/features/checkout/session
- Webhook setup and delivery: https://polar.sh/docs/integrate/webhooks/endpoints and https://polar.sh/docs/integrate/webhooks/delivery
- Webhook events: https://polar.sh/docs/integrate/webhooks/events
- TypeScript SDK: https://polar.sh/docs/integrate/sdk/typescript
- Next.js adapter: https://polar.sh/docs/integrate/sdk/adapters/nextjs
- Express adapter: https://polar.sh/docs/integrate/sdk/adapters/express
- BetterAuth adapter: https://polar.sh/docs/integrate/sdk/adapters/better-auth
- Products: https://polar.sh/docs/features/products
- Create product API: https://polar.sh/docs/api-reference/products/create
- Subscriptions: https://polar.sh/docs/features/subscriptions/manage
- Update subscription API: https://polar.sh/docs/api-reference/subscriptions/update
- Proration: https://polar.sh/docs/features/subscriptions/proration
- Usage-based billing: https://polar.sh/docs/features/usage-based-billing/introduction
- Event ingestion: https://polar.sh/docs/features/usage-based-billing/event-ingestion
- Credits: https://polar.sh/docs/features/usage-based-billing/credits
- Benefits: https://polar.sh/docs/features/benefits/introduction
- Feature flags: https://polar.sh/docs/features/benefits/feature-flags
- License keys: https://polar.sh/docs/features/benefits/license-keys
- Customer sessions: https://polar.sh/docs/api-reference/customer-portal/sessions/create
- Fees: https://polar.sh/docs/merchant-of-record/fees

## Current-State Findings

### 1. Checkout API examples should move from price IDs to product IDs

Current local docs still use `product_price_id` / `productPriceId` in several examples.

Official Polar docs now describe checkout creation as `products: [productId]`, where `products` is a required array of product IDs. The checkout response still contains `product_price_id`, but it is marked deprecated.

Recommended update:

- Change checkout creation examples to use `products`.
- Document multiple-product checkout sessions.
- Keep `product_price_id` only as legacy/deprecated response context.
- Add ad-hoc price examples through the checkout `prices` mapping.

Affected local files:

- `checkouts.md`
- `sdk.md`
- `best-practices.md`
- `subscriptions.md`

### 2. API rate-limit guidance is stale

Local docs say production rate limits are `300/min`.

Official API overview now states:

- Production: `500 requests/minute` per organization/customer or OAuth2 client.
- Sandbox: `100 requests/minute`.
- Unauthenticated license validation, activation, and deactivation endpoints: `3 requests/second`.
- `429` responses include `Retry-After`.

Recommended update:

- Update `overview.md` and retry guidance.
- Split production and sandbox limits explicitly.
- Preserve the unauthenticated license endpoint limit.

### 3. Fee calculation changed after May 27, 2026

Local `best-practices.md` uses the older Early Member rate as the generic fee model: `4% + $0.40`, plus `+0.5%` for subscriptions.

Official Polar fees now distinguish current plans from Early Member:

- Starter: `5% + $0.50`.
- Pro/Growth/Scale: monthly plans with lower variable rates.
- Early Member: `4% + $0.40 + 0.5% subscription fee`, only for organizations created before May 27, 2026 and still on that plan.
- Additional `+1.5%` international card fee still applies.

Recommended update:

- Replace the hard-coded fee helper with plan-aware guidance.
- Label the old rate as Early Member only.
- Avoid treating December 2025 fees as current default behavior.

### 4. Webhook docs need stronger raw-body and Standard Webhooks guidance

Local docs have several correct concepts, but examples are inconsistent. Some examples use parsed bodies, and manual HMAC guidance is not aligned with the current official recommendation.

Official docs say Polar webhooks follow the Standard Webhooks specification. The TypeScript example uses `express.raw({ type: 'application/json' })` and `validateEvent(req.body, req.headers, secret)`.

Official delivery behavior:

- Endpoint timeout is currently 10 seconds.
- Polar recommends responding within 2 seconds.
- Failed deliveries are retried up to 10 times with exponential backoff.
- Endpoints are disabled after 10 consecutive failed deliveries.
- SDK validation avoids manual base64-secret pitfalls.

Recommended update:

- Prefer SDK/adapters for validation.
- For Express, use raw body middleware for the webhook route.
- Remove or clearly mark hand-rolled HMAC examples as advanced fallback.
- Update timeout guidance from 20 seconds to current 10 seconds / 2-second recommendation.
- Add retry and endpoint-disable behavior.

Affected local files:

- `webhooks.md`
- `subscriptions.md`
- `benefits.md`
- `best-practices.md`

### 5. Webhook event set is broader than local examples

Local docs list many important events, but miss or underplay some current official events.

Official current event guidance includes:

- `checkout.created`
- `checkout.updated`
- `checkout.expired`
- `customer.created`
- `customer.updated`
- `customer.deleted`
- `customer.state_changed`
- `subscription.created`
- `subscription.active`
- `subscription.uncanceled`
- `subscription.canceled`
- `subscription.past_due`
- `subscription.updated`
- `subscription.revoked`
- `order.created`
- `order.paid`
- `order.updated`
- `order.refunded`
- `refund.created`
- `refund.updated`
- `benefit_grant.created`
- `benefit_grant.updated`
- `benefit_grant.revoked`
- `benefit.created`
- `benefit.updated`
- `product.created`
- `product.updated`
- `organization.updated`

Recommended update:

- Add `checkout.expired`, customer events, `subscription.uncanceled`, `subscription.past_due`, `order.updated`, and `order.paid`.
- Clarify that `order.created` can be pending, while `order.paid` means payment has been received.
- Add cancellation and renewal event sequences from official docs.

### 6. Subscription update examples are stale

Local subscription examples use `product_price_id` and proration values such as `invoice_immediately` / `next_invoice`.

Official docs now use:

- `product_id` for plan changes.
- `proration_behavior` with values such as `invoice`, `prorate`, and `next_period`.
- `discount_id` for subscription discounts.
- `trial_end` for extending or ending trials.
- DELETE `/v1/subscriptions/{id}` to revoke immediately.

Official subscription statuses include `incomplete`, `incomplete_expired`, `trialing`, `active`, `past_due`, `canceled`, and `unpaid`.

Recommended update:

- Replace subscription update examples with `product_id`.
- Replace proration names with `invoice`, `prorate`, and `next_period`.
- Document cancel-at-period-end vs immediate revoke.
- Update status mapping.

### 7. Usage-based billing needs its own refreshed reference

Local `products.md` mentions usage-based billing, but the API shape appears stale.

Official usage-based billing is based on:

- Events ingested through `polar.events.ingest`.
- Event fields: `name`, `customerId` or `externalCustomerId`, and `metadata`.
- Meters that filter and aggregate ingested events.
- Metered prices on subscription products.
- Credits benefit for prepaid usage.
- Customer State / Customer Meters API for balance.

Local examples use `events.create`, `event_name`, and `properties`, which do not match the current official examples.

Recommended update:

- Add `usage-based-billing.md`.
- Use `events.ingest({ events: [...] })` examples.
- Emphasize server-side ingestion for billable events.
- Document that Polar does not block usage when a balance is exceeded; the app must enforce its own usage gate.
- Add credits-only spending and customer meter balance guidance.

### 8. Benefits docs are missing current benefit types and details

Local `benefits.md` covers license keys, GitHub, Discord, files, credits, and custom benefits. It does not cover Feature Flag benefits, and credits are described like a custom benefit instead of the official Credits benefit.

Official benefit types include:

- Credits
- License Keys
- Feature Flags
- File Downloads
- GitHub Repository Access
- Discord Invite / roles
- Custom benefits

License key official guidance also says to validate `benefit_id` when an organization has multiple license-key benefits.

Recommended update:

- Add Feature Flag benefits.
- Update Credits to use the official benefit model.
- Add license-key `benefit_id` scoping guidance.
- Keep the existing benefit-grant lifecycle guidance; it remains useful.

### 9. Customer Portal and Customer Sessions deserve a dedicated reference

Local docs mention customer sessions and portal access, but the current official response shape should be documented more precisely.

Official customer session API returns:

- `token`
- `expires_at`
- `return_url`
- `customer_portal_url`
- `customer_id`
- `customer`

The API supports creation by `customer_id` or `external_customer_id`, and newer member-model fields are present: `member_id` and `external_member_id`.

Recommended update:

- Add `customer-portal.md`.
- Use `customer_portal_url`, not a generic `session.url`, unless a specific SDK adapter returns a different alias.
- Document Customer Portal API vs Core API security boundaries.

### 10. Framework adapter examples should be refreshed

Official docs now show framework-specific adapters:

- Next.js: `@polar-sh/nextjs` with `Checkout`, `CustomerPortal`, and `Webhooks`.
- Express: `@polar-sh/express`.
- Laravel: `checkout([product_id])`, `redirectToCustomerPortal`, and CSRF exclusion for webhook routes.
- BetterAuth: `@polar-sh/better-auth` with `checkout`, `portal`, `usage`, and `webhooks` plugins.

Recommended update:

- Refresh adapter examples in `sdk.md`.
- Prefer adapter-native checkout and webhook handlers where the framework is supported.
- Use `products` query params for adapter checkout routes.

### 11. Product model docs need current pricing and product fields

Local product docs are close conceptually, but should reflect current official product capabilities:

- Product billing cycles include one-time and recurring day/week/month/year intervals.
- `recurring_interval_count` supports intervals such as every 2 weeks or every 3 months.
- Pricing models include fixed, pay-what-you-want/custom, free, metered, and seat-based.
- Metered prices can stack with fixed base pricing.
- Multiple payment currencies are supported.
- Tax behavior and customer IP forwarding can affect checkout currency/tax behavior.
- Product visibility, media, custom fields, and attached benefits are official product concerns.

Recommended update:

- Refresh product creation examples.
- Avoid legacy recurring price fields where official docs mark them deprecated.
- Add multiple-currency and customer IP guidance.
- Add seat-based limitations and beta/feature-flag note where appropriate.

## Content That Remains Useful

These local concepts still align with official Polar docs:

- Polar as Merchant of Record.
- Production and sandbox are isolated environments.
- Organization Access Tokens should stay server-side.
- SDK `server: "sandbox"` usage.
- `external_customer_id` for reconciliation.
- Do not trust checkout success redirects as final payment proof; verify through webhooks/API.
- Idempotent webhook processing.
- Benefit grants as the entitlement lifecycle.
- Customer portal as the preferred self-service surface.

## Suggested Update Priority

### P0: Correct stale facts that can cause broken integrations

- Checkout creation: `products` array instead of `product_price_id`.
- Subscription update: `product_id` and `proration_behavior`.
- Rate limits.
- Fees and Early Member distinction.
- Webhook raw body, timeout, retry, and endpoint-disable behavior.

### P1: Add missing current feature references

- Usage-based billing reference.
- Customer portal/customer sessions reference.
- Feature Flag benefits.
- Cost insights and `_cost` metadata if the skill should cover profitability tracking.

### P2: Refresh SDK and framework examples

- TypeScript SDK camelCase note.
- Next.js, Express, Laravel, BetterAuth adapters.
- Adapter-specific checkout query params and webhook handlers.

## Evidence Notes

- Local files were read from `internal/aicontext/skills/payment-integration/references/polar`.
- Official Polar docs were checked on 2026-06-14.
- This report intentionally does not modify the Polar skill references.
