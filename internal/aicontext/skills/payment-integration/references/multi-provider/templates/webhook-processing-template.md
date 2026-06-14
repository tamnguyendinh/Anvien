# Webhook Processing Template

Use this template for provider-neutral webhook/IPN intake. Load provider-specific docs for the actual signature algorithm, raw-body requirements, event names, and retry behavior.

## Shared Intake Rule

Do not acknowledge a webhook/IPN before authentication. Verify first, record or queue the event, then ACK quickly.

```typescript
interface ParsedProviderEvent {
  id: string;
  type: string;
  providerObjectId?: string;
  createdAt?: string;
  data: unknown;
}

interface ProviderWebhookAdapter {
  provider: PaymentProvider;
  verify(input: {
    rawBody: string;
    headers: Headers;
    environment: ProviderEnvironment;
  }): Promise<{ ok: true } | { ok: false; status: number; message: string }>;
  parse(rawBody: string): Promise<ParsedProviderEvent>;
}

const webhookAdapters: Record<PaymentProvider, ProviderWebhookAdapter> = {
  sepay: sepayWebhookAdapter,
  polar: polarWebhookAdapter,
  stripe: stripeWebhookAdapter,
  paddle: paddleWebhookAdapter,
  creem: creemWebhookAdapter,
};
```

## HTTP Handler

```typescript
export async function POST(request: Request, params: { provider: PaymentProvider }) {
  const provider = params.provider;
  const adapter = webhookAdapters[provider];
  const rawBody = await request.text();
  const environment = resolveProviderEnvironment(provider);

  const verified = await adapter.verify({
    rawBody,
    headers: request.headers,
    environment,
  });

  if (!verified.ok) {
    return new Response(verified.message, { status: verified.status });
  }

  const parsed = await adapter.parse(rawBody);

  await recordProviderEventOnce({
    provider,
    providerEnvironment: environment,
    providerEventId: parsed.id,
    eventType: parsed.type,
    eventCreatedAt: parsed.createdAt,
    providerObjectId: parsed.providerObjectId,
    rawPayload: rawBody,
    signatureVerified: true,
    processingStatus: "received",
  });

  await enqueueProviderEvent({
    provider,
    providerEnvironment: environment,
    providerEventId: parsed.id,
  });

  return new Response("ok", { status: 200 });
}
```

## Dispatcher

Keep provider-specific event names at the edge. Convert them into local actions.

```typescript
type LocalEventAction =
  | "mark_paid"
  | "mark_payment_failed"
  | "mark_refunded"
  | "mark_disputed"
  | "sync_subscription"
  | "sync_entitlement"
  | "ignore";

function mapProviderEventToAction(provider: PaymentProvider, eventType: string): LocalEventAction {
  if (provider === "polar") {
    if (eventType === "order.paid") return "mark_paid";
    if (eventType === "order.refunded") return "mark_refunded";
    if (eventType === "customer.state_changed") return "sync_entitlement";
    if (eventType.startsWith("subscription.")) return "sync_subscription";
  }

  if (provider === "paddle") {
    if (eventType === "transaction.completed") return "mark_paid";
    if (eventType === "transaction.payment_failed") return "mark_payment_failed";
    if (eventType.startsWith("subscription.")) return "sync_subscription";
  }

  if (provider === "creem") {
    if (eventType === "checkout.completed" || eventType === "payment.succeeded") return "mark_paid";
    if (eventType === "refund.created") return "mark_refunded";
    if (eventType === "chargeback.created") return "mark_disputed";
    if (eventType.startsWith("license.")) return "sync_entitlement";
    if (eventType.startsWith("subscription.")) return "sync_subscription";
  }

  if (provider === "stripe") {
    if (eventType === "checkout.session.completed") return "mark_paid";
    if (eventType === "invoice.paid") return "mark_paid";
    if (eventType === "charge.refunded") return "mark_refunded";
    if (eventType.includes("dispute")) return "mark_disputed";
  }

  if (provider === "sepay") {
    if (eventType === "bank_transaction_matched") return "mark_paid";
    if (eventType === "gateway_order_paid") return "mark_paid";
  }

  return "ignore";
}
```

## Idempotent Processing

```typescript
async function processRecordedProviderEvent(eventId: string) {
  const event = await getProviderEventForUpdate(eventId);
  if (!event || event.processingStatus === "processed") return;

  const action = mapProviderEventToAction(event.provider, event.eventType);

  try {
    switch (action) {
      case "mark_paid":
        await handlePaidProviderEvent(event);
        break;
      case "mark_payment_failed":
        await handlePaymentFailedProviderEvent(event);
        break;
      case "mark_refunded":
        await handleRefundProviderEvent(event);
        break;
      case "mark_disputed":
        await handleDisputeProviderEvent(event);
        break;
      case "sync_subscription":
        await syncSubscriptionFromProvider(event);
        break;
      case "sync_entitlement":
        await syncEntitlementsFromProvider(event);
        break;
      case "ignore":
        await markProviderEventIgnored(event.id);
        return;
    }

    await markProviderEventProcessed(event.id);
  } catch (error) {
    await markProviderEventFailed(event.id, serializeError(error));
    throw error;
  }
}
```

Rules:

- Ignore duplicate event IDs only after confirming the original event was recorded.
- Handle out-of-order events by fetching the provider object before final access or revenue decisions.
- Persist failures and retry from the local event table.
- Keep raw payloads for support and audit, with app-appropriate PII retention.
