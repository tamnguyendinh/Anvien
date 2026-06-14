# Polar SDK and Framework Adapters

Current SDK and adapter guidance for Polar TypeScript, Python, PHP, Go, Next.js, Express, Laravel, BetterAuth, and webhook handling.

## Source Links

- TypeScript SDK: https://polar.sh/docs/integrate/sdk/typescript
- Next.js adapter: https://polar.sh/docs/integrate/sdk/adapters/nextjs
- Express adapter: https://polar.sh/docs/integrate/sdk/adapters/express
- BetterAuth adapter: https://polar.sh/docs/integrate/sdk/adapters/better-auth
- Laravel adapter: https://polar.sh/docs/integrate/sdk/adapters/laravel

## TypeScript SDK

Install:

```bash
npm install @polar-sh/sdk
```

Configure:

```typescript
import { Polar } from "@polar-sh/sdk";

export const polar = new Polar({
  accessToken: process.env.POLAR_ACCESS_TOKEN,
  server: process.env.POLAR_SERVER === "production" ? "production" : "sandbox",
});
```

The API reference uses `snake_case`; the TypeScript SDK uses `camelCase`.

Examples:

```typescript
await polar.checkouts.create({
  products: [productId],
  successUrl: "https://example.com/success?checkout_id={CHECKOUT_ID}",
  externalCustomerId: user.id,
});

await polar.events.ingest({
  events: [
    {
      name: "ai_usage",
      externalCustomerId: user.id,
      metadata: { total_tokens: 1000 },
    },
  ],
});

await polar.customerSessions.create({
  externalCustomerId: user.id,
  returnUrl: "https://example.com/account",
});
```

Do not copy REST `snake_case` names directly into TypeScript examples unless the SDK explicitly supports them.

## Pagination

Polar APIs use `page` and `limit`; `limit` max is 100. SDKs provide pagination helpers where available.

```typescript
const result = await polar.products.list({ limit: 100 });

for await (const page of result) {
  for (const product of page.result.items) {
    console.log(product.id, product.name);
  }
}
```

Check the generated SDK types for exact pagination shape because SDKs can change.

## Python, PHP, and Go

Official SDKs exist for:

- TypeScript/JavaScript: `@polar-sh/sdk`
- Python: official Polar SDK package
- PHP: official Polar SDK package
- Go: official Polar Go SDK

Use official SDK docs for exact package names and generated method signatures. Keep this skill focused on integration shape and common pitfalls rather than copying large generated SDK examples.

## Next.js Adapter

Install:

```bash
npm install zod @polar-sh/nextjs
```

Checkout route:

```typescript
import { Checkout } from "@polar-sh/nextjs";

export const GET = Checkout({
  accessToken: process.env.POLAR_ACCESS_TOKEN,
  successUrl: "https://app.example.com/success?checkout_id={CHECKOUT_ID}",
  returnUrl: "https://app.example.com/pricing",
  server: "sandbox",
});
```

Use query params:

- `?products=PRODUCT_ID`
- `customerId`
- `customerExternalId`
- `customerEmail`
- `customerName`
- URL-encoded `metadata`

Customer Portal:

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

Webhooks:

```typescript
import { Webhooks } from "@polar-sh/nextjs";

export const POST = Webhooks({
  webhookSecret: process.env.POLAR_WEBHOOK_SECRET!,
  onOrderPaid: async (payload) => {
    await fulfillPaidOrder(payload.data);
  },
  onCustomerStateChanged: async (payload) => {
    await syncCustomerAccess(payload.data);
  },
  onPayload: async (payload) => {
    await storeWebhook(payload);
  },
});
```

## Express Adapter

Install:

```bash
npm install zod @polar-sh/express
```

Checkout:

```typescript
import express from "express";
import { Checkout } from "@polar-sh/express";

const app = express();

app.get(
  "/checkout",
  Checkout({
    accessToken: process.env.POLAR_ACCESS_TOKEN,
    successUrl: "https://app.example.com/success?checkout_id={CHECKOUT_ID}",
    returnUrl: "https://app.example.com/pricing",
    server: "sandbox",
  }),
);
```

Customer Portal:

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

Webhooks:

```typescript
import { Webhooks } from "@polar-sh/express";

app.use(express.json()).post(
  "/webhooks/polar",
  Webhooks({
    webhookSecret: process.env.POLAR_WEBHOOK_SECRET!,
    onPayload: async (payload) => {
      await storeWebhook(payload);
    },
  }),
);
```

If you do not use the adapter, use `@polar-sh/sdk/webhooks` with raw body handling.

## BetterAuth Adapter

Install:

```bash
npm install better-auth @polar-sh/better-auth @polar-sh/sdk
```

BetterAuth plugin features include:

- Automatic Polar customer creation on signup.
- Customer deletion sync.
- Reference system for organization/team purchases.
- Checkout plugin.
- Customer Portal plugin.
- Usage plugin for customer meters and event ingestion.
- Webhook plugin with signature verification.

Configuration shape:

```typescript
import { betterAuth } from "better-auth";
import { polar, checkout, portal, usage, webhooks } from "@polar-sh/better-auth";
import { Polar } from "@polar-sh/sdk";

const polarClient = new Polar({
  accessToken: process.env.POLAR_ACCESS_TOKEN,
  server: "sandbox",
});

export const auth = betterAuth({
  plugins: [
    polar({
      client: polarClient,
      createCustomerOnSignUp: true,
      use: [
        checkout({
          products: [
            {
              productId: process.env.POLAR_PRODUCT_ID!,
              slug: "pro",
            },
          ],
          successUrl: "/success?checkout_id={CHECKOUT_ID}",
          authenticatedUsersOnly: true,
        }),
        portal(),
        usage(),
        webhooks({
          secret: process.env.POLAR_WEBHOOK_SECRET,
          onCustomerStateChanged: async (payload) => {},
          onOrderPaid: async (payload) => {},
          onPayload: async (payload) => {},
        }),
      ],
    }),
  ],
});
```

Use BetterAuth only when the app already uses Better Auth or wants auth-integrated billing.

## Laravel Adapter

The Laravel page currently documents `danestves/laravel-polar` and states that the provider is not maintained or officially supported by Polar. Use it at your own discretion.

Install:

```bash
composer require danestves/laravel-polar
php artisan polar:install
```

Notable patterns:

- Add the Billable trait to the model that owns billing.
- Exclude webhook routes from CSRF.
- Checkout accepts product IDs, including multiple products.
- `redirectToCustomerPortal()` and `customerPortalUrl()` expose the portal.

Example checkout:

```php
Route::get('/subscribe', function (Request $request) {
    return $request->user()->checkout(['product_id_123']);
});
```

Because the adapter is not officially maintained by Polar, inspect its current repository, version, and issue tracker before production use.

## Webhook Validation Choice

Preferred order:

1. Framework adapter `Webhooks` handler.
2. `@polar-sh/sdk/webhooks` `validateEvent` with raw body.
3. Custom Standard Webhooks validation only when no SDK/adapter is viable.

Do not use parsed/re-serialized JSON for custom validation.

## Error and Retry Handling

- Handle HTTP `429` using `Retry-After`.
- Log status codes and request IDs where available.
- Retry transient network/server errors with bounded backoff.
- Do not retry non-idempotent operations without local idempotency keys.
- For checkouts, prefer creating one checkout per explicit user action and storing the resulting checkout ID.

## Adapter Selection

| Stack | Preferred Surface |
|-------|-------------------|
| Next.js | `@polar-sh/nextjs` checkout, portal, webhooks |
| Express | `@polar-sh/express` checkout, portal, webhooks |
| Better Auth app | `@polar-sh/better-auth` plugin |
| Laravel | `danestves/laravel-polar` with caution; not officially maintained by Polar |
| Other JS frameworks | Check official framework adapter docs first |
| Custom backend | `@polar-sh/sdk` or official language SDK |

## Checkout Embed Package

Use `@polar-sh/checkout` for embedded checkout and embedded payment-method flows.

```bash
npm install @polar-sh/checkout
```

Embedded checkout:

```typescript
import { PolarEmbedCheckout } from "@polar-sh/checkout/embed";

PolarEmbedCheckout.init();
```

Embedded payment method:

```typescript
import { PolarEmbedPaymentMethod } from "@polar-sh/checkout/payment-method";

await PolarEmbedPaymentMethod.create({
  sessionToken,
  setAsDefault: true,
  returnUrl: "https://app.example.com/account/payment-methods",
});
```

Use the checkout/customer portal references for server-side session creation and security boundaries.
