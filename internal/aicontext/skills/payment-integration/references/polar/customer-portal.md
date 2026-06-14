# Polar Customer Portal

Current guidance for Customer Portal, customer sessions, portal API boundaries, self-service billing, and integration patterns.

## Source Links

- Customer Portal feature: https://polar.sh/docs/features/customer-portal/introduction
- Create customer session API: https://polar.sh/docs/api-reference/customer-portal/sessions/create
- Customer State: https://polar.sh/docs/integrate/customer-state
- Next.js adapter: https://polar.sh/docs/integrate/sdk/adapters/nextjs
- Express adapter: https://polar.sh/docs/integrate/sdk/adapters/express

## What Customers Can Do

The hosted Customer Portal lets customers:

- View active subscriptions and purchase history.
- Download receipts and invoices.
- Edit invoice details such as company name, tax ID, or billing address.
- Access benefits, including license keys and file downloads.
- Cancel active subscriptions.
- Update default payment method.
- Optionally change email, switch plans, manage seats, and view metered usage depending on portal settings.

The portal is always available and cannot be disabled. Use it as the default self-service billing surface unless you have a strong reason to build a custom portal.

## Customer Sessions

Create a customer session server-side with Organization Access Token.

REST by external customer ID:

```json
{
  "external_customer_id": "user_123",
  "return_url": "https://app.example.com/account"
}
```

REST by Polar customer ID:

```json
{
  "customer_id": "customer_id",
  "return_url": "https://app.example.com/account"
}
```

Response includes:

- `id`
- `token`
- `expires_at`
- `return_url`
- `customer_portal_url`
- `customer_id`
- `customer`

Redirect users to `customer_portal_url`, not a guessed `session.url`.

Generate pre-authenticated portal links at click time. Do not store portal URLs as permanent account links.

TypeScript:

```typescript
const session = await polar.customerSessions.create({
  externalCustomerId: user.id,
  returnUrl: "https://app.example.com/account",
});

return redirect(session.customerPortalUrl);
```

## Member Model Fields

For organizations with member model enabled, customer sessions can include:

- `member_id`
- `external_member_id`

If omitted, Polar can create a member session for the owner member of the customer where applicable.

## Core API vs Customer Portal API

Core API:

- Uses Organization Access Token.
- Merchant/server privileged.
- Can create products, subscriptions, checkouts, refunds, benefits, customer sessions.

Customer Portal API:

- Uses customer-scoped session token.
- Customer-specific.
- Can read/update only allowed customer-facing surfaces.

Never expose an Organization Access Token to build a custom customer portal. Use customer sessions or framework adapters.

## Framework Adapters

### Next.js

```typescript
import { CustomerPortal } from "@polar-sh/nextjs";

export const GET = CustomerPortal({
  accessToken: process.env.POLAR_ACCESS_TOKEN,
  getCustomerId: async (req) => {
    const user = await requireUser(req);
    return user.polarCustomerId;
  },
  returnUrl: "https://app.example.com/account",
  server: "sandbox",
});
```

### Express

```typescript
import { CustomerPortal } from "@polar-sh/express";

app.get(
  "/portal",
  CustomerPortal({
    accessToken: process.env.POLAR_ACCESS_TOKEN,
    getCustomerId: async (event) => event.user.polarCustomerId,
    returnUrl: "https://app.example.com/account",
    server: "sandbox",
  }),
);
```

## When to Use Portal Instead of Custom API

Prefer Customer Portal for:

- Payment method update.
- Failed payment recovery.
- Subscription cancellation.
- Invoice and receipt download.
- License key/file access.
- Seat management when enabled.
- Metered usage display when enabled.

Use your own Core API routes only for merchant-side actions that customers should not perform directly, such as curated plan migrations, internal refunds, or admin support operations.

## Embedded Payment Method

If you need a custom "add payment method" UI, use embedded payment method with a short-lived customer session token.

Server:

```typescript
const session = await polar.customerSessions.create({
  customerId: customerId,
});
```

Client:

```typescript
import { PolarEmbedPaymentMethod } from "@polar-sh/checkout/payment-method";

await PolarEmbedPaymentMethod.create({
  sessionToken: session.token,
  setAsDefault: true,
  returnUrl: "https://app.example.com/account/payment-methods",
});
```

Options include modal, inline, and React component flows. Redirect-based payment methods such as Amazon Pay or Klarna need a `returnUrl`.

Use Customer Portal instead if the customer only needs standard failed-payment recovery or default payment method management.

## Portal and Webhooks

Portal actions still produce webhooks:

- Cancellation: `subscription.updated`, `subscription.canceled`, and later `subscription.revoked`.
- Payment method recovery can lead to renewed successful orders.
- Entitlement changes can trigger `customer.state_changed`.

Your local database should treat the portal as another source of billing changes and sync through webhooks/API reconciliation.

## Security

- Create sessions server-side after authenticating your user.
- Use `externalCustomerId` only after verifying the user owns that ID.
- Keep portal URLs short-lived and do not store them as permanent links.
- Keep `return_url` on trusted app domains.
- Do not use portal session tokens for merchant operations.

## Best Practices

- Add "manage billing" to account settings and redirect to Customer Portal.
- Use portal for failed payment recovery emails.
- Store Polar customer ID after first checkout/customer creation, but keep external customer ID as the durable app mapping.
- Keep entitlement state synchronized from `customer.state_changed`.
- Link users to portal for invoice/tax ID changes instead of collecting card/tax details yourself.
