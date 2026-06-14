# Polar Products and Pricing

Current guidance for Polar product catalog design, billing cycles, pricing models, currencies, tax behavior, custom fields, media, and lifecycle constraints.

## Source Links

- Products guide: https://polar.sh/docs/features/products
- Create product API: https://polar.sh/docs/api-reference/products/create
- Usage-based billing: https://polar.sh/docs/features/usage-based-billing/introduction
- Seat-based pricing: https://polar.sh/docs/features/seat-based-pricing
- Custom fields: https://polar.sh/docs/features/custom-fields
- Tax inclusive pricing: https://polar.sh/docs/features/tax-inclusive-pricing

## Product Model

Polar uses products for both one-time purchases and subscriptions. Do not model one-time products and subscription plans as unrelated concepts in your app; they are different product configurations in the same Polar catalog.

## Billing Cycle

Products are either:

- One-time: charged once; access can be lifetime or benefit-defined.
- Recurring: billed on an interval.

Recurring intervals:

- `day`
- `week`
- `month`
- `year`

`recurring_interval_count` supports patterns like every 2 weeks or every 3 months.

Billing cycle and recurring interval are locked at creation. Create a new product if they need to change.

## Pricing Models

| Model | Use For | Notes |
|-------|---------|-------|
| Fixed | Standard price | Fixed amount can be changed for new customers. |
| Custom / Pay What You Want | Customer chooses amount | Can set minimum/default amount. |
| Free | Lead magnets, free tiers, sign-up gated benefits | May still grant benefits. |
| Metered | Usage-based billing | Recurring products only; load `usage-based-billing.md`. |
| Seat-based | Team seats with optional volume tiers | Load official seat-based docs before implementation. |

Pricing type is locked at creation. For fixed prices, existing subscribers keep the price they started with; new purchases see the changed price.

## Metered Prices

Metered prices can stack with other recurring price components. This supports plans such as:

- Base fee plus API-call overage.
- Base fee plus token usage.
- Multiple meters on one product, such as API calls and storage.

Usage-based billing requires events, meters, and metered prices. Load `usage-based-billing.md` before implementing.

## Multiple Currencies

Products can have prices in multiple currencies. Polar chooses a checkout currency from customer geolocation and falls back to the organization default if no matching product currency is enabled.

Important:

- The price structure must match across enabled currencies.
- When creating checkout sessions server-side, pass `customer_ip_address` / `customerIpAddress` so Polar can infer customer geolocation instead of your server IP.

## Tax Behavior

Tax behavior controls whether displayed price includes tax or tax is added on top. Polar defaults by customer country conventions, but organization settings can override defaults.

For customer-facing calculators:

- Do not assume gross/net amount until checkout or tax calculation data is available.
- Include customer country and tax ID handling where relevant.
- Treat fee estimates separately from tax estimates.

## Product Creation Shape

REST uses `snake_case`; TypeScript SDK uses `camelCase`.

TypeScript example:

```typescript
await polar.products.create({
  organizationId: process.env.POLAR_ORGANIZATION_ID!,
  name: "Pro Monthly",
  description: "Full product access",
  recurringInterval: "month",
  recurringIntervalCount: 1,
  prices: [
    {
      amountType: "fixed",
      priceAmount: 2000,
      priceCurrency: "usd",
    },
  ],
  metadata: {
    tier: "pro",
  },
});
```

REST example:

```json
{
  "organization_id": "org_id",
  "name": "Pro Monthly",
  "description": "Full product access",
  "recurring_interval": "month",
  "recurring_interval_count": 1,
  "prices": [
    {
      "amount_type": "fixed",
      "price_amount": 2000,
      "price_currency": "usd"
    }
  ]
}
```

## Product Media and Checkout Page

Products can include checkout copy and media. Product media helps customers inspect what they are buying.

Use product media for:

- Digital asset previews.
- Screenshots.
- Course/package thumbnails.
- Product-specific trust material.

Keep checkout description and media product-specific. Put broad marketing pages in the app, not in Polar metadata.

## Custom Fields

Custom fields collect customer-provided data at checkout. Supported field types include text, number, date, checkbox, and select.

Use custom fields for:

- Billing or fulfillment details not otherwise captured.
- Legal checkbox confirmations.
- Project setup data needed after purchase.

Do not collect secrets or unnecessary personal data. Values appear on resulting orders or subscriptions.

## Benefits

Benefits are attached to products and grant what customers receive:

- Credits
- License Keys
- Feature Flags
- File Downloads
- GitHub Repository Access
- Discord Access
- Custom benefits

Load `benefits.md` before using benefits as entitlements.

## Product Updates

Editable after creation:

- Name and description.
- Product media.
- Metadata.
- Benefits.
- Fixed price amount for new customers.
- Archive/unarchive state.

Locked or effectively immutable:

- Billing cycle.
- Recurring interval.
- Pricing type.
- Historical orders and subscriptions.

Products are archived rather than permanently deleted. Existing customers keep their access according to subscription/order/benefit state.

## Design Recommendations

- Use separate products for materially different billing cycles or pricing models.
- Use multiple products in checkout for monthly/yearly/lifetime choice.
- Use metadata for internal identifiers, not customer-visible data.
- Use custom fields for customer-provided checkout data.
- Treat product IDs as current checkout inputs.
- Treat price IDs as legacy/deprecated unless an official current endpoint specifically asks for them.
