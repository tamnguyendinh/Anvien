# Multi-Provider Order Management Patterns

Provider-neutral order, payment, refund, entitlement, revenue, and reconciliation patterns for apps that use more than one payment provider.

Use this reference when an app mixes providers such as SePay, Polar, Stripe, Paddle, or Creem.io. Always load the provider-specific references too; this file defines local data ownership and cross-provider patterns, not provider API authority.

Detailed examples are intentionally split into templates. The templates use placeholders such as `productKey`, `entitlementKey`, `campaignId`, and `providerObjectId`; replace them with the target app's domain terms.

## When to Use

Load this file when the app needs any of these:

- One local order table across multiple providers.
- Cross-provider revenue reporting or currency normalization.
- Unified refunds, chargebacks, disputes, and commission reversal.
- Separate checkout state from paid order state.
- Entitlement sync across subscriptions, license keys, benefits, or manual bank payments.
- Reconciliation after webhook outages or manual support actions.

Do not use this file as a replacement for provider docs:

- SePay: load `references/sepay/overview.md`, then `api.md`, `webhooks.md`, `payment-gateway.md`, or `qr-codes.md` as needed.
- Polar: load `references/polar/overview.md`, then `orders-refunds-discounts.md`, `customer-state.md`, `customer-portal.md`, `usage-based-billing.md`, and related refs as needed.
- Stripe: load `references/stripe/stripe-best-practices.md` first.
- Paddle: load `references/paddle/overview.md`, `api.md`, `paddle-js.md`, `subscriptions.md`, `webhooks.md`, and `sdk.md` as needed.
- Creem.io: load `references/creem/overview.md`, `api.md`, `checkouts.md`, `subscriptions.md`, `licensing.md`, `webhooks.md`, and `sdk.md` as needed.

## Template Routing

Load only the template needed for the implementation slice:

| Need | Template |
|------|----------|
| Local order, provider reference, event, refund, entitlement, campaign tables | `references/multi-provider/templates/local-data-model-template.md` |
| Webhook verification, event ledger, dispatcher, idempotent processing | `references/multi-provider/templates/webhook-processing-template.md` |
| Provider-neutral checkout command and provider-specific checkout builders | `references/multi-provider/templates/checkout-factory-template.md` |
| Paid-order handling, local fulfillment, entitlement grants | `references/multi-provider/templates/fulfillment-entitlements-template.md` |
| Refunds, disputes, access consequences, cross-provider campaigns | `references/multi-provider/templates/refunds-disputes-campaigns-template.md` |
| Currency reporting, commissions, revenue splits, reconciliation jobs | `references/multi-provider/templates/revenue-commissions-reconciliation-template.md` |
| Admin/support APIs, manual actions, example app variants | `references/multi-provider/templates/admin-support-app-variants-template.md` |

## Core Design Rule

Keep provider facts and local business facts separate.

Provider facts are the raw external records: checkout sessions, orders, transactions, subscriptions, refunds, events, customer IDs, and license IDs. Local business facts are app decisions: access granted, commission approved, refund policy, reporting currency, and support notes.

Do not collapse every provider state into one vague `completed` value. Store:

- local normalized status for app workflows;
- provider raw status for audit and debugging;
- provider event history for replay and reconciliation;
- separate entitlement state for access decisions.

## Local Status Model

Use a small normalized status set locally and map provider events into it.

| Local Status | Meaning |
|--------------|---------|
| `draft` | Local order exists before checkout or payment instruction is created. |
| `checkout_created` | Provider checkout/session/transaction/payment instruction exists. |
| `awaiting_payment` | Customer has not produced a confirmed paid signal yet. |
| `paid` | Provider paid signal has been verified. Fulfillment may proceed. |
| `partially_refunded` | Some money has been refunded; access policy is app-specific. |
| `refunded` | Fully refunded; access and commission reversal depend on policy/provider. |
| `failed` | Payment failed or provider rejected the attempt. |
| `expired` | Checkout or bank-payment instruction expired. |
| `cancelled` | Customer or merchant cancelled before payment or ended a subscription. |
| `disputed` | Chargeback/dispute opened; fulfillment and payout risk need review. |

Keep subscription lifecycle separate from one-time order state. A subscription can be `active` while an individual renewal order is `paid`, `failed`, or `refunded`.

## Provider State Mapping

Use provider-specific paid signals as the authority for money received.

| Provider | Local Paid Signal | Notes |
|----------|-------------------|-------|
| SePay bank webhook | Verified bank-account webhook transaction matched to the local order and amount policy. | HMAC/API-key/OAuth rules depend on the SePay surface. Load `sepay/webhooks.md`. |
| SePay Payment Gateway | Verified IPN or gateway order state such as `ORDER_PAID`. | Load `sepay/payment-gateway.md`; gateway IPN is not the same surface as bank-account webhooks. |
| Polar | `order.paid` or verified API lookup showing paid order state. | `order.created` is not enough for fulfillment. Customer State is best for current access, not revenue recognition. |
| Stripe | Checkout Session or PaymentIntent/Invoice paid state according to the chosen Stripe integration. | Use Checkout Sessions/Billing/PaymentIntents guidance; do not use Charges as the default. |
| Paddle | `transaction.completed` for successful payment; subscription events for recurring lifecycle. | Load `paddle/webhooks.md` and `paddle/subscriptions.md`. |
| Creem.io | `checkout.completed` or `payment.succeeded`; subscription/license events for lifecycle and access. | Load `creem/webhooks.md`, `creem/subscriptions.md`, and `creem/licensing.md`. |

Refund and dispute events must map separately from paid events:

- Polar: `order.refunded` and refund API state.
- Creem.io: `refund.created`, `chargeback.created`.
- Paddle and Stripe: use their webhook/API refund and dispute surfaces.
- SePay bank transfers: usually require manual refund/reconciliation unless the gateway surface provides a refund operation.

## Data Ownership

Model these local records independently:

- `orders`: local commercial intent, status, amount, currency, product key, buyer.
- `provider_payment_references`: provider IDs such as checkout, order, transaction, subscription, payment intent, refund, license, customer.
- `provider_events`: raw verified webhook/IPN events with idempotency and processing status.
- `refunds`: refund intent, provider refund state, access and commission consequences.
- `entitlements`: current access state sourced from order, subscription, benefit, license, or manual support action.
- `campaign_redemptions`: cross-provider discount/campaign reservations and redemptions.

Use `local-data-model-template.md` for concrete interface and schema examples.

## Processing Rules

1. Create local orders before provider checkout or payment instruction.
2. Store local order ID in provider metadata/custom data when supported.
3. Verify webhooks/IPNs before acknowledging them.
4. Record every provider event before business processing.
5. Fulfill from provider paid signals, not customer redirects.
6. Keep order, subscription, refund, dispute, and entitlement state separate.
7. Normalize currency for reporting only; keep original provider amounts.
8. Use provider-reported fees/net amounts when possible.
9. Use a local campaign ledger for cross-provider discounts.
10. Reconcile provider state on a schedule and after webhook outages.

## Implementation Checklist

Before shipping a multi-provider payment implementation, verify:

- Provider-specific docs were loaded for every provider in use.
- Local order ID is stored in provider metadata where supported.
- Every provider event is signature/auth verified before ACK.
- Event table has a uniqueness rule for idempotency.
- Paid handler is idempotent and does not depend on success redirects.
- Refunds and disputes update local order, entitlement, and commission state.
- Subscription state is separate from one-time order state.
- Entitlements have a source object and can be reconciled.
- Cross-provider campaigns use local reservations and releases.
- Revenue reports preserve provider currency and reporting currency separately.
- Fee/net amount fields are provider-reported or clearly marked estimated.
- Reconciliation jobs cover stale checkout, paid-without-fulfillment, refund, dispute, subscription, and failed-event cases.
- Admin tools can search by provider object IDs.
- Manual support actions write audit logs.
