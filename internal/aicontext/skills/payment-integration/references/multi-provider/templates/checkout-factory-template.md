# Checkout Factory Template

Use this template when one app can start checkout/payment attempts through multiple providers. Keep the shared command neutral and put provider-specific API details behind builders.

## Shared Command

```typescript
interface CheckoutCommand {
  provider: PaymentProvider;
  environment: ProviderEnvironment;
  localOrderId: string;
  userId?: string;
  buyerEmail?: string;
  productKey: string;
  quantity: number;
  amountMinor: number;
  currency: string;
  successUrl: string;
  cancelUrl?: string;
  metadata?: Record<string, string>;
}

interface CheckoutResult {
  provider: PaymentProvider;
  providerCheckoutId?: string;
  providerCheckoutUrl?: string;
  providerOrderId?: string;
  providerTransactionId?: string;
  providerPaymentIntentId?: string;
  providerCustomerId?: string;
  redirectUrl?: string;
  expiresAt?: string;
  raw: unknown;
}

async function createProviderCheckout(command: CheckoutCommand): Promise<CheckoutResult> {
  switch (command.provider) {
    case "polar":
      return createPolarCheckout(command);
    case "stripe":
      return createStripeCheckout(command);
    case "paddle":
      return createPaddleTransactionCheckout(command);
    case "creem":
      return createCreemCheckout(command);
    case "sepay":
      return createSepayPaymentInstruction(command);
  }
}
```

## Start Checkout Flow

```typescript
async function startCheckout(params: {
  provider: PaymentProvider;
  environment: ProviderEnvironment;
  userId?: string;
  buyerEmail?: string;
  productKey: string;
  quantity?: number;
  amountMinor: number;
  currency: string;
  successUrl: string;
  cancelUrl?: string;
}) {
  const order = await createLocalOrder({
    provider: params.provider,
    providerEnvironment: params.environment,
    userId: params.userId,
    buyerEmail: params.buyerEmail,
    productKey: params.productKey,
    quantity: params.quantity ?? 1,
    amountMinor: params.amountMinor,
    currency: params.currency,
    localStatus: "draft",
  });

  const checkout = await createProviderCheckout({
    provider: params.provider,
    environment: params.environment,
    localOrderId: order.id,
    userId: params.userId,
    buyerEmail: params.buyerEmail,
    productKey: params.productKey,
    quantity: params.quantity ?? 1,
    amountMinor: params.amountMinor,
    currency: params.currency,
    successUrl: params.successUrl,
    cancelUrl: params.cancelUrl,
    metadata: { localOrderId: order.id },
  });

  await storeProviderPaymentReference(order.id, checkout);
  await markOrderStatus(order.id, "checkout_created");

  return checkout.redirectUrl;
}
```

## Provider Builders

Polar:

```typescript
async function createPolarCheckout(command: CheckoutCommand): Promise<CheckoutResult> {
  const productId = await resolvePolarProductId(command.productKey);

  const checkout = await polar.checkouts.create({
    products: [productId],
    successUrl: `${command.successUrl}?checkout_id={CHECKOUT_ID}`,
    customerEmail: command.buyerEmail,
    externalCustomerId: command.userId,
    metadata: {
      localOrderId: command.localOrderId,
      productKey: command.productKey,
      ...command.metadata,
    },
  });

  return {
    provider: "polar",
    providerCheckoutId: checkout.id,
    providerCheckoutUrl: checkout.url,
    redirectUrl: checkout.url,
    expiresAt: checkout.expiresAt,
    raw: checkout,
  };
}
```

Stripe:

```typescript
async function createStripeCheckout(command: CheckoutCommand): Promise<CheckoutResult> {
  const priceId = await resolveStripePriceId(command.productKey, command.currency);

  const session = await stripe.checkout.sessions.create({
    mode: "payment",
    line_items: [{ price: priceId, quantity: command.quantity }],
    success_url: `${command.successUrl}?session_id={CHECKOUT_SESSION_ID}`,
    cancel_url: command.cancelUrl,
    client_reference_id: command.localOrderId,
    customer_email: command.buyerEmail,
    metadata: { localOrderId: command.localOrderId, productKey: command.productKey },
  });

  return {
    provider: "stripe",
    providerCheckoutId: session.id,
    providerPaymentIntentId:
      typeof session.payment_intent === "string" ? session.payment_intent : undefined,
    providerCustomerId: typeof session.customer === "string" ? session.customer : undefined,
    redirectUrl: session.url ?? undefined,
    expiresAt: session.expires_at ? new Date(session.expires_at * 1000).toISOString() : undefined,
    raw: session,
  };
}
```

Paddle:

```typescript
async function createPaddleTransactionCheckout(command: CheckoutCommand): Promise<CheckoutResult> {
  const priceId = await resolvePaddlePriceId(command.productKey, command.currency);

  const transaction = await paddle.transactions.create({
    items: [{ priceId, quantity: command.quantity }],
    customerId: await maybeResolvePaddleCustomerId(command.userId),
    customData: { localOrderId: command.localOrderId, productKey: command.productKey },
  });

  return {
    provider: "paddle",
    providerTransactionId: transaction.id,
    providerCustomerId: transaction.customerId,
    providerCheckoutUrl: transaction.checkout?.url,
    redirectUrl: transaction.checkout?.url,
    raw: transaction,
  };
}
```

Creem.io:

```typescript
async function createCreemCheckout(command: CheckoutCommand): Promise<CheckoutResult> {
  const productId = await resolveCreemProductId(command.productKey);

  const checkout = await creem.checkouts.create({
    product_id: productId,
    success_url: command.successUrl,
    customer: command.buyerEmail ? { email: command.buyerEmail } : undefined,
    metadata: { localOrderId: command.localOrderId, productKey: command.productKey },
  });

  return {
    provider: "creem",
    providerCheckoutId: checkout.id,
    providerCustomerId: checkout.customer_id,
    redirectUrl: checkout.url,
    raw: checkout,
  };
}
```

SePay bank transfer/VietQR:

```typescript
async function createSepayPaymentInstruction(command: CheckoutCommand): Promise<CheckoutResult> {
  const transferCode = buildTransferCode(command.localOrderId);
  const account = await resolveSepayReceivingAccount(command.currency);

  const instruction = {
    bankAccount: account,
    amountMinor: command.amountMinor,
    currency: command.currency,
    transferContent: transferCode,
    qrUrl: buildVietQrUrl({ amountMinor: command.amountMinor, content: transferCode, account }),
  };

  return {
    provider: "sepay",
    providerCheckoutId: transferCode,
    providerCheckoutUrl: instruction.qrUrl,
    redirectUrl: instruction.qrUrl,
    raw: instruction,
  };
}
```

Provider notes:

- Polar: create checkouts with product IDs through `products`; do not use `product_price_id` as the default create field.
- SePay: bank transfer/VietQR can be local payment instructions plus bank webhook reconciliation; Payment Gateway should store gateway order/IPN IDs.
- Stripe: Checkout Session is usually the simplest web checkout surface.
- Paddle: transaction checkout state and subscription lifecycle must be tracked separately.
- Creem.io: checkout, subscription, license, and revenue-split IDs can all matter.
