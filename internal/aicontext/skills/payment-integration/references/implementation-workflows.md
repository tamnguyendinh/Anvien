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

## Creem.io Implementation
1. Load `references/creem/overview.md` for auth and MoR concepts
2. Load `references/creem/api.md` for products and checkout sessions
3. Load `references/creem/checkouts.md` for payment flow options
4. Load `references/creem/webhooks.md` for event handling
5. Load `references/creem/subscriptions.md` if implementing recurring billing
6. Load `references/creem/licensing.md` if implementing device activation
7. Load `references/creem/sdk.md` for framework-specific adapters

## General Workflow
1. Identify platform (Vietnamese → SePay, global SaaS → Polar/Stripe/Creem.io)
2. Load relevant references progressively
3. Implement: auth → products → checkout → webhooks → events
4. Test in sandbox, then production
5. Load only needed references to maintain context efficiency
