---
name: payment-integration
description: Use for payments, subscriptions, webhooks, refunds, entitlements, or multi-provider orders with SePay, Polar, Stripe, Paddle, or Creem.
version: 2.2.1
license: MIT
---

# Payment Integration

Repo-agnostic payment integration guidance for target apps using SePay, Polar, Stripe, Paddle, Creem.io, or multiple providers.

Do not assume this skill is for integrating payments into the current repository. Apply provider and template guidance to the target app's domain model.

## When to Use

- Payment gateway integration (checkout, processing)
- Subscription management (trials, upgrades, billing)
- Webhook handling (notifications, idempotency)
- QR code payments (VietQR, NAPAS)
- Software licensing (device activation)
- Multi-provider order management
- Revenue splits and commissions

## Platform Selection

| Platform | Best For |
|----------|----------|
| **SePay** | Vietnamese market, VND, bank transfers, VietQR |
| **Polar** | Global SaaS, subscriptions, automated benefits (GitHub/Discord) |
| **Stripe** | Enterprise payments, Connect platforms, custom checkout |
| **Paddle** | MoR subscriptions, global tax compliance, churn prevention |
| **Creem.io** | MoR + licensing, revenue splits, no-code checkout |

## Quick Reference

### SePay
- `references/sepay/overview.md` - Surface selection, environments, safety rules
- `references/sepay/api.md` - API v2, transactions, reconciliation, Order VAs
- `references/sepay/webhooks.md` - Bank-account webhook setup, HMAC, replay safety
- `references/sepay/payment-gateway.md` - Hosted checkout, gateway API, IPN
- `references/sepay/sdk.md` - Official Payment Gateway Node.js/PHP SDKs
- `references/sepay/qr-codes.md` - VietQR generation and bank-specific rules
- `references/sepay/best-practices.md` - Repo-agnostic production patterns

### Polar
- `references/polar/overview.md` - Auth, MoR concept
- `references/polar/products.md` - Product model, pricing, currencies, tax
- `references/polar/checkouts.md` - Checkout sessions, links, embedded checkout
- `references/polar/subscriptions.md` - Lifecycle management
- `references/polar/webhooks.md` - Event handling
- `references/polar/benefits.md` - Automated delivery
- `references/polar/usage-based-billing.md` - Events, meters, credits
- `references/polar/customer-portal.md` - Customer sessions and self-service
- `references/polar/customer-state.md` - Entitlement state and access sync
- `references/polar/orders-refunds-discounts.md` - Orders, refunds, discounts
- `references/polar/sdk.md` - Multi-language SDKs
- `references/polar/best-practices.md` - Production patterns

### Stripe
- `references/stripe/stripe-best-practices.md` - Integration design
- `references/stripe/stripe-sdks.md` - Server SDKs
- `references/stripe/stripe-js.md` - Payment Element
- `references/stripe/stripe-cli.md` - Local testing
- `references/stripe/stripe-upgrade.md` - Version upgrades
- External: https://docs.stripe.com/llms.txt

### Paddle
- `references/paddle/overview.md` - MoR, auth, entity IDs
- `references/paddle/api.md` - Products, prices, transactions
- `references/paddle/paddle-js.md` - Checkout overlay/inline
- `references/paddle/subscriptions.md` - Trials, upgrades, pause
- `references/paddle/webhooks.md` - SHA256 verification
- `references/paddle/sdk.md` - Node, Python, PHP, Go
- `references/paddle/best-practices.md` - Production patterns
- External: https://developer.paddle.com/llms.txt

### Creem.io
- `references/creem/overview.md` - MoR, auth, global support
- `references/creem/api.md` - Products, checkout sessions
- `references/creem/checkouts.md` - No-code links, storefronts
- `references/creem/subscriptions.md` - Trials, seat-based
- `references/creem/licensing.md` - Device activation
- `references/creem/webhooks.md` - Signature verification
- `references/creem/sdk.md` - Next.js, Better Auth
- External: https://docs.creem.io/llms.txt

### Multi-Provider
- `references/multi-provider-order-management-patterns.md` - Provider-neutral design rules and template routing
- `references/multi-provider/templates/local-data-model-template.md` - Orders, provider IDs, events, refunds, entitlements, campaigns
- `references/multi-provider/templates/webhook-processing-template.md` - Verification, event ledger, dispatcher, retries
- `references/multi-provider/templates/checkout-factory-template.md` - Provider-neutral checkout command and provider builders
- `references/multi-provider/templates/fulfillment-entitlements-template.md` - Paid handlers and entitlement grants
- `references/multi-provider/templates/refunds-disputes-campaigns-template.md` - Refunds, disputes, access consequences, campaigns
- `references/multi-provider/templates/revenue-commissions-reconciliation-template.md` - FX, reporting, commissions, reconciliation
- `references/multi-provider/templates/admin-support-app-variants-template.md` - Admin views, support actions, app variants

### Scripts
- `scripts/sepay-webhook-verify.js` - SePay webhook verification
- `scripts/polar-webhook-verify.js` - Polar webhook verification
- `scripts/checkout-helper.js` - Legacy checkout helper examples; prefer provider references for new API work

## Key Capabilities

| Platform | Highlights |
|----------|------------|
| **SePay** | API v2, QR/bank/cards, VN banks, webhooks, gateway IPN |
| **Polar** | MoR, subscriptions, usage billing, benefits, customer portal |
| **Stripe** | CheckoutSessions, Billing, Connect, Payment Element |
| **Paddle** | MoR, overlay/inline checkout, Retain (churn prevention), tax |
| **Creem.io** | MoR, licensing, revenue splits, no-code checkout |
| **Multi-Provider** | Unified order state, provider event ledgers, refunds, entitlements, reporting, reusable templates |

## Implementation

See `references/implementation-workflows.md` for step-by-step guides per platform.

For multi-provider work, load `references/multi-provider-order-management-patterns.md` first, then load only the focused template needed for the current slice.

**General flow:** auth -> products/prices -> checkout/payment instruction -> webhooks/IPN -> orders/refunds/entitlements -> reconciliation
