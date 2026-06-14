# Implementation Workflows

## SePay Implementation
1. Load `references/sepay/overview.md` first and choose the SePay surface.
2. For proactive lookup, reconciliation, bank-account lookup, or Order VAs, load `references/sepay/api.md`.
3. For real-time bank-account transaction events, load `references/sepay/webhooks.md`.
4. For hosted checkout, cards/NAPAS, gateway order lifecycle, or Payment Gateway IPN, load `references/sepay/payment-gateway.md`; then load `references/sepay/sdk.md` if using the official SDK.
5. For manually rendered VietQR transfer instructions, load `references/sepay/qr-codes.md`.
6. Use `scripts/sepay-webhook-verify.js` only as an inspectable helper for its currently supported auth modes; prefer the HMAC guidance in `references/sepay/webhooks.md` for new production webhook work.
7. Load `references/sepay/best-practices.md` for production readiness.

## Polar Implementation
1. Load `references/polar/overview.md` first and choose the Polar surface.
2. For catalog setup, pricing, currencies, tax behavior, or custom fields, load `references/polar/products.md`.
3. For checkout sessions, checkout links, embedded checkout, discounts, or checkout identity, load `references/polar/checkouts.md`.
4. For recurring billing, trials, plan changes, proration, cancellation, or revoke, load `references/polar/subscriptions.md`.
5. For webhook delivery, event taxonomy, idempotency, fulfillment, or entitlement sync, load `references/polar/webhooks.md`; prefer official SDK/adapters over `scripts/polar-webhook-verify.js` for new production webhook work.
6. For automated entitlements, license keys, feature flags, credits, GitHub/Discord, files, or custom benefits, load `references/polar/benefits.md`.
7. For event ingestion, meters, metered prices, credits, customer meter balances, or usage gates, load `references/polar/usage-based-billing.md`.
8. For self-service billing, customer sessions, portal links, payment method updates, invoices, or portal API boundaries, load `references/polar/customer-portal.md`.
9. For entitlement state, active subscriptions, granted benefits, meter balances, or one-call access sync, load `references/polar/customer-state.md`.
10. For paid order state, invoices, receipts, refunds, benefit revocation, or discounts, load `references/polar/orders-refunds-discounts.md`.
11. For TypeScript SDK, framework adapters, Next.js, Express, BetterAuth, Laravel, or embedded checkout/payment-method packages, load `references/polar/sdk.md`.
12. Load `references/polar/best-practices.md` for production readiness, fees, retries, reconciliation, and data-model patterns.

## Stripe Implementation
1. Load `references/stripe/stripe-best-practices.md` for integration design
2. Load `references/stripe/stripe-sdks.md` for server-side SDK setup
3. Load `references/stripe/stripe-js.md` for client-side Elements/Checkout
4. Use `stripe listen` via CLI for local webhook testing (`references/stripe/stripe-cli.md`)
5. Choose integration: Checkout (hosted/embedded) or Payment Element
6. Use CheckoutSessions API for most payment flows
7. Use Billing APIs for subscriptions (combine with Checkout)
8. Load `references/stripe/stripe-upgrade.md` when upgrading API versions

## Paddle Implementation
1. Load `references/paddle/overview.md` first for MoR concepts, environments, entity IDs, and SDK installation.
2. Load `references/paddle/api.md` for products, prices, transactions, customers, subscriptions, response format, and errors.
3. Load `references/paddle/paddle-js.md` for checkout overlay, inline checkout, and client integration.
4. Load `references/paddle/subscriptions.md` for trials, upgrades, pause/resume, cancellation, and subscription lifecycle.
5. Load `references/paddle/webhooks.md` for `Paddle-Signature` verification, event taxonomy, and idempotency.
6. Load `references/paddle/sdk.md` for Node.js, Python, PHP, and Go SDK usage.
7. Load `references/paddle/best-practices.md` for production readiness.

## Creem.io Implementation
1. Load `references/creem/overview.md` for auth and MoR concepts
2. Load `references/creem/api.md` for products and checkout sessions
3. Load `references/creem/checkouts.md` for payment flow options
4. Load `references/creem/webhooks.md` for event handling
5. Load `references/creem/subscriptions.md` if implementing recurring billing
6. Load `references/creem/licensing.md` if implementing device activation
7. Load `references/creem/sdk.md` for framework-specific adapters

## Multi-Provider Order Management
1. Load `references/multi-provider-order-management-patterns.md` when the target app uses more than one payment provider or needs unified orders, refunds, revenue, commissions, entitlements, or reconciliation.
2. Load each provider-specific reference before implementing provider API calls. The multi-provider reference defines local app data ownership; provider references define API authority.
3. Load the focused template under `references/multi-provider/templates/` for the slice being implemented: data model, webhook processing, checkout factory, fulfillment/entitlements, refunds/campaigns, revenue/reconciliation, or admin/support.
4. Model local orders, provider payment references, provider events, refunds, disputes, and entitlements separately.
5. Fulfill from provider paid signals, not customer redirects.
6. Keep provider raw statuses and local normalized statuses side by side for support and reconciliation.

## General Workflow
1. Identify platform: Vietnamese bank payments -> SePay; SaaS/MoR -> Polar, Paddle, or Creem.io; custom global payment infrastructure -> Stripe.
2. Load relevant references progressively
3. If using multiple providers, load `references/multi-provider-order-management-patterns.md`, then the needed template file before designing shared order/refund/reporting tables.
4. Implement: auth -> products/prices -> checkout/payment instruction -> webhooks/IPN -> orders/refunds/entitlements -> reconciliation.
5. Test in sandbox/test mode, then production.
6. Load only needed references to maintain context efficiency.
