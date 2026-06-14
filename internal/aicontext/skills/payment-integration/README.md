# Payment Integration Skill

Repo-agnostic skill for implementing payments, subscriptions, webhooks, refunds, entitlements, reporting, and multi-provider order management in any target app.

This skill is not a payment implementation for this repository. Use it as reusable guidance for target apps that need SePay, Polar, Stripe, Paddle, Creem.io, or a combination of providers.

## Supported Providers

### SePay
- Vietnamese bank payments, VietQR, NAPAS/cards, bank-account webhooks, Payment Gateway, Order VAs, and reconciliation.
- Start with `references/sepay/overview.md`, then load API, webhook, payment gateway, QR, SDK, or best-practice references as needed.

### Polar
- Global SaaS monetization, Merchant of Record, subscriptions, usage-based billing, benefits, customer portal, customer state, orders, refunds, and discounts.
- Start with `references/polar/overview.md`, then load products, checkouts, subscriptions, webhooks, benefits, usage billing, customer portal/state, orders/refunds/discounts, SDK, or best-practice references as needed.

### Stripe
- Global payment infrastructure, Checkout Sessions, Payment Element, Billing, Connect, CLI testing, SDKs, and API version upgrades.
- Start with `references/stripe/stripe-best-practices.md`, then load SDK, Stripe.js, CLI, or upgrade references as needed.

### Paddle
- Merchant of Record subscriptions and transactions, Paddle.js checkout, lifecycle changes, webhook verification, SDK usage, and production readiness.
- Start with `references/paddle/overview.md`, then load API, Paddle.js, subscriptions, webhooks, SDK, or best-practice references as needed.

### Creem.io
- Merchant of Record checkout, subscriptions, licensing, revenue splits, no-code checkout, framework adapters, and webhook handling.
- Start with `references/creem/overview.md`, then load API, checkouts, subscriptions, licensing, webhooks, or SDK references as needed.

## Multi-Provider Templates

For apps that combine providers, load `references/multi-provider-order-management-patterns.md` first, then load the focused template for the slice being implemented:

| Need | Template |
|------|----------|
| Local schemas and ownership boundaries | `references/multi-provider/templates/local-data-model-template.md` |
| Webhook verification, event ledger, dispatcher, retries | `references/multi-provider/templates/webhook-processing-template.md` |
| Provider-neutral checkout command and provider builders | `references/multi-provider/templates/checkout-factory-template.md` |
| Paid handlers, local fulfillment, entitlement grants | `references/multi-provider/templates/fulfillment-entitlements-template.md` |
| Refunds, disputes, access consequences, campaigns | `references/multi-provider/templates/refunds-disputes-campaigns-template.md` |
| FX, reporting, commissions, reconciliation | `references/multi-provider/templates/revenue-commissions-reconciliation-template.md` |
| Admin views, support actions, app variants | `references/multi-provider/templates/admin-support-app-variants-template.md` |

Template placeholders such as `productKey`, `entitlementKey`, `campaignId`, and `providerObjectId` are intentionally generic. Replace them with the target app's domain terms.

## Structure

```text
payment-integration/
├── SKILL.md
├── README.md
├── references/
│   ├── implementation-workflows.md
│   ├── multi-provider-order-management-patterns.md
│   ├── multi-provider/templates/
│   ├── sepay/
│   ├── polar/
│   ├── stripe/
│   ├── paddle/
│   └── creem/
└── scripts/
    ├── sepay-webhook-verify.js
    ├── polar-webhook-verify.js
    ├── checkout-helper.js
    ├── test-scripts.js
    ├── package.json
    └── .env.example
```

## Usage

1. Read `SKILL.md` for activation and quick routing.
2. Read `references/implementation-workflows.md` for the provider workflow.
3. Load only the provider references needed for the target app.
4. For multi-provider work, load the pattern file and the focused template.
5. Treat scripts as helper examples, not as provider API authority.

## Script Helpers

Scripts are optional helpers for inspection and local validation:

- `scripts/sepay-webhook-verify.js` verifies supported SePay webhook auth modes.
- `scripts/polar-webhook-verify.js` verifies Polar webhook signatures.
- `scripts/checkout-helper.js` contains legacy helper examples; for new Polar checkout work, prefer `references/polar/checkouts.md` and the official SDK/API shape with `products: [productId]`.

Run script tests from the skill directory only when changing scripts:

```bash
cd scripts
npm test
```

## Selection Guide

- Use SePay for Vietnamese bank payments, VietQR, local transfer flows, or VN payment gateway needs.
- Use Polar for SaaS monetization, subscriptions, benefits, customer portal, and Merchant of Record flows.
- Use Stripe for custom global payment infrastructure, Checkout Sessions, Payment Element, Billing, or Connect.
- Use Paddle for Merchant of Record subscription businesses with Paddle checkout and lifecycle workflows.
- Use Creem.io for Merchant of Record software products, licensing, checkout, subscriptions, and revenue splits.
- Use multi-provider patterns when one app needs unified orders, provider events, refunds, entitlements, campaigns, reporting, commissions, or reconciliation across providers.

## Version

2.2.1
