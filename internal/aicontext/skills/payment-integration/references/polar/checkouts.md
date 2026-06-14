# Polar Checkouts

Checkout guidance for Polar checkout links, API-created sessions, embedded checkout, customer identity, ad-hoc prices, discounts, and fulfillment boundaries.

## Source Links

- Checkout API guide: https://polar.sh/docs/features/checkout/session
- Create checkout session API: https://polar.sh/docs/api-reference/checkouts/create-session
- Embedded checkout: https://polar.sh/docs/features/checkout/embedded-checkout
- Checkout links: https://polar.sh/docs/features/checkout/links

## Key Rule

For new checkout sessions, use product IDs through `products`.

Do not use `product_price_id` / `productPriceId` as the default create field. `product_price_id` can still appear in legacy objects or deprecated response fields.

## Checkout Options

| Option | Use When | Reference |
|--------|----------|-----------|
| Checkout Link | Static or campaign link, low backend control | Dashboard or Checkout Links API |
| Checkout Session API | App creates sessions after auth, validation, metadata, or pricing decisions | `POST /v1/checkouts` |
| Embedded Checkout | Checkout should open inside the app while Polar still hosts payment collection | Checkout session with `embed_origin` |
| Framework Adapter | Next.js/Express/etc. should expose checkout route quickly | `sdk.md` |

## Checkout Links

Checkout Links are long-lived URLs configured in Polar. Each visit creates a new short-lived Checkout Session.

Rules:

- Share the Checkout Link URL, not a generated Checkout Session URL.
- Query parameters can override/prefill fields per visit.
- Multiple products on a Checkout Link give the customer a choice; they do not create one bundled multi-product order.
- Metadata on the link is copied to generated sessions and resulting orders/subscriptions.

Useful query params:

- `product_id`: preselect one configured product.
- `customer_email`
- `customer_name`
- `discount_code`: prefill a visible discount code field.
- `amount`: prefill pay-what-you-want amount.
- `locale`: force checkout language when localization is enabled.
- `custom_field_data.{slug}`: prefill custom field values.
- `theme`: `light` or `dark`.
- `reference_id`, `utm_source`, `utm_medium`, `utm_campaign`, `utm_content`, `utm_term`: attribution metadata.

## TypeScript SDK Session

The API docs use `snake_case`; the TypeScript SDK uses `camelCase`.

```typescript
import { Polar } from "@polar-sh/sdk";

const polar = new Polar({
  accessToken: process.env.POLAR_ACCESS_TOKEN,
  server: "sandbox",
});

const checkout = await polar.checkouts.create({
  products: [process.env.POLAR_PRODUCT_ID!],
  successUrl: "https://example.com/success?checkout_id={CHECKOUT_ID}",
  returnUrl: "https://example.com/pricing",
  externalCustomerId: user.id,
  customerEmail: user.email,
  customerName: user.name,
  customerIpAddress: requestIp,
  allowDiscountCodes: true,
  metadata: {
    userId: user.id,
    source: "pricing_page",
  },
});

return checkout.url;
```

## REST Session

```bash
curl -X POST https://api.polar.sh/v1/checkouts \
  -H "Authorization: Bearer $POLAR_ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "products": ["PRODUCT_ID"],
    "success_url": "https://example.com/success?checkout_id={CHECKOUT_ID}",
    "return_url": "https://example.com/pricing",
    "external_customer_id": "user_123",
    "customer_email": "buyer@example.com",
    "allow_discount_codes": true,
    "metadata": {
      "user_id": "user_123"
    }
  }'
```

## Multiple Products

Pass multiple product IDs when the customer should choose between variants such as monthly, yearly, and lifetime plans.

```typescript
const checkout = await polar.checkouts.create({
  products: [
    process.env.POLAR_MONTHLY_PRODUCT_ID!,
    process.env.POLAR_YEARLY_PRODUCT_ID!,
    process.env.POLAR_LIFETIME_PRODUCT_ID!,
  ],
  successUrl: "https://example.com/success?checkout_id={CHECKOUT_ID}",
  externalCustomerId: user.id,
});
```

Products are displayed in the order passed in `products`.

## Ad-Hoc Prices

Ad-hoc prices are temporary prices attached only to the checkout session. Use them for dynamic or negotiated pricing without changing the product catalog.

```typescript
const checkout = await polar.checkouts.create({
  products: [productId],
  prices: {
    [productId]: [
      {
        amountType: "fixed",
        priceAmount: 10000,
        priceCurrency: "usd",
      },
    ],
  },
  successUrl: "https://example.com/success?checkout_id={CHECKOUT_ID}",
});
```

Supported ad-hoc price types mirror catalog price types: fixed, custom/pay-what-you-want, free, seat-based, and metered.

## Customer Identity

Use `externalCustomerId` / `external_customer_id` for application users.

Effects:

- Links the order/subscription/customer to your own user ID.
- Creates a customer if none exists.
- Pre-fills and locks customer email if the customer is known.
- Provides `customer.external_id` in webhooks.

Do not rely on email as the durable key. Email can change.

## Customer IP Address

When your server creates checkout sessions, Polar sees your server IP. Forward the actual customer IP as `customerIpAddress` / `customer_ip_address` so currency detection, billing country, and tax calculation are more accurate.

```typescript
await polar.checkouts.create({
  products: [productId],
  customerIpAddress: ipFromRequest,
  successUrl: "https://example.com/success?checkout_id={CHECKOUT_ID}",
});
```

Checkout links created directly by Polar do not need this workaround because the customer opens Polar directly.

## Discounts

```typescript
await polar.checkouts.create({
  products: [productId],
  discountId: discountId,
  allowDiscountCodes: false,
  successUrl: "https://example.com/success?checkout_id={CHECKOUT_ID}",
});
```

Rules:

- Use `discountId` to pre-apply a discount.
- Set `allowDiscountCodes: false` if pre-applied discounts must not stack with customer-entered codes.
- Free and custom prices may not be discountable in every scenario; check `is_discount_applicable` on checkout responses if needed.

## Embedded Checkout

Set `embedOrigin` / `embed_origin` when embedding checkout in your app.

```typescript
const checkout = await polar.checkouts.create({
  products: [productId],
  embedOrigin: "https://app.example.com",
  externalCustomerId: user.id,
  successUrl: "https://app.example.com/success?checkout_id={CHECKOUT_ID}",
});
```

The response includes `id`, `clientSecret`, `url`, `status`, `expiresAt`, and selected product/price details. Use official embedded checkout docs or framework adapters for client code.

### Embedded Checkout Client

For simple sites, use a checkout link with `data-polar-checkout` and the Polar checkout embed script.

For JS/React apps, use `@polar-sh/checkout`:

```bash
npm install @polar-sh/checkout
```

```typescript
import { PolarEmbedCheckout } from "@polar-sh/checkout/embed";

PolarEmbedCheckout.init();
```

If using an API-created Checkout Session URL instead of a Checkout Link, set `embedOrigin` when creating the session. Embedded checkout events such as `confirmed` or `success` are UX events; they do not replace `order.paid` webhook/API verification.

### Embedded Payment Method

Embedded payment-method flow lets an existing customer add a payment method without leaving your site.

Server:

```typescript
const session = await polar.customerSessions.create({
  customerId: customerId,
});
```

Client:

```typescript
import { PolarEmbedPaymentMethod } from "@polar-sh/checkout/payment-method";

const embed = await PolarEmbedPaymentMethod.create({
  sessionToken: session.token,
  setAsDefault: true,
  returnUrl: "https://app.example.com/account/payment-methods",
});

embed.addEventListener("success", (event) => {
  console.log(event.detail);
});
```

Rules:

- Create the customer session server-side; never expose the Organization Access Token.
- Customer session tokens are short-lived and scoped to one customer.
- Support redirect-based payment methods by setting `returnUrl`.
- Use Customer Portal if you do not need a custom payment-method UI.

## Localization

Checkout localization is beta. When enabled, Polar can auto-detect the customer's browser language.

Override language:

- API: pass `locale`.
- Checkout Link: append `?locale=fr` or another supported BCP 47 tag.

Known limitations during beta:

- Some error messages remain English.
- Transactional emails can remain English.
- There is no checkout-page language selector.

## Checkout Status

| Status | Meaning |
|--------|---------|
| `open` | Customer can use the session. |
| `expired` | Session is no longer usable. |
| `confirmed` | Customer clicked Pay; not proof of successful payment. |
| `succeeded` | Payment completed successfully. |
| `failed` | Checkout failed and cannot be retried. |

Never fulfill solely from the success redirect or `confirmed`. Use `order.paid` or a verified API lookup.

## Webhook Events to Handle

- `checkout.created`: session created.
- `checkout.updated`: session state changed.
- `checkout.expired`: checkout link/session expired before completion.
- `order.created`: order or invoice object exists; can be pending.
- `order.updated`: order changed.
- `order.paid`: payment has been received; fulfill here.
- `order.refunded`: refund completed.

Load `webhooks.md` for verification and idempotency.

## Framework Shortcuts

Next.js and Express adapters expose checkout routes that accept query params such as:

- `?products=PRODUCT_ID`
- `customerId`
- `customerExternalId`
- `customerEmail`
- `customerName`
- URL-encoded `metadata`

Load `sdk.md` for adapter details.

## Best Practices

- Create sessions server-side.
- Use `products` and product IDs for new work.
- Use `success_url` / `successUrl` with `checkout_id={CHECKOUT_ID}` for post-redirect lookup.
- Store local pending order state before creating checkout if your app needs reconciliation.
- Fulfill on verified `order.paid`, not redirect.
- Include `external_customer_id` and stable metadata for reconciliation.
- Forward customer IP when creating sessions from backend/proxy/edge code.
- Do not expose Organization Access Tokens client-side.
