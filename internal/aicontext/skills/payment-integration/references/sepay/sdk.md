# SePay SDK Integration

Use this file for official Payment Gateway SDKs. For API v2 bank-account lookup, use `api.md`; for bank-account webhooks, use `webhooks.md`.

Official docs checked: 2026-06-14.

## Node.js SDK

Official package:

```bash
npm i sepay-pg-node
```

Requirements:

- Node.js 16+
- gateway merchant ID
- gateway secret key

Initialize:

```javascript
import { SePayPgClient } from 'sepay-pg-node';

const client = new SePayPgClient({
  env: 'sandbox', // or 'production'
  merchant_id: process.env.SEPAY_MERCHANT_ID,
  secret_key: process.env.SEPAY_MERCHANT_SECRET_KEY,
});
```

## One-Time Payment Form

```javascript
const checkoutUrl = client.checkout.initCheckoutUrl();
const fields = client.checkout.initOneTimePaymentFields({
  operation: 'PURCHASE',
  payment_method: 'CARD', // CARD, BANK_TRANSFER, NAPAS_BANK_TRANSFER
  order_invoice_number: invoiceNumber,
  order_amount: amountVnd,
  currency: 'VND',
  order_description: description,
  customer_id: customerId,
  success_url: successUrl,
  error_url: errorUrl,
  cancel_url: cancelUrl,
  custom_data: JSON.stringify({ orderId }),
});
```

Render the returned fields as hidden form inputs and submit to `checkoutUrl`.

Do not hand-roll the signature unless needed. If you do, preserve SePay's documented field order exactly:

```text
merchant, operation, payment_method, order_amount, currency,
order_invoice_number, order_description, customer_id,
success_url, error_url, cancel_url
```

Changing field order can invalidate the signature.

## Payment Methods

Gateway checkout supports:

- `CARD`
- `BANK_TRANSFER`
- `NAPAS_BANK_TRANSFER`

Use `payment-gateway.md` for choosing the gateway surface and IPN contract. Use `qr-codes.md` only when rendering QR instructions yourself outside the gateway checkout flow.

## Gateway Order APIs

The SDK exposes Payment Gateway order APIs.

```javascript
await client.order.all({
  per_page: 20,
  q: 'search-keyword',
  order_status: 'CAPTURED',
  from_created_at: '2026-01-01',
  to_created_at: '2026-01-31',
});

await client.order.retrieve(orderInvoiceNumber);
await client.order.voidTransaction(orderInvoiceNumber); // cards
await client.order.cancel(orderInvoiceNumber); // unpaid bank/NAPAS transfer orders
```

Order detail responses can include:

- `order_status`
- `authentication_status`
- transaction history
- amount and currency
- invoice number

Important statuses:

- `CAPTURED`: paid
- `CANCELLED`: cancelled
- `AUTHENTICATION_NOT_NEEDED`: awaiting payment

Cancel applies to `BANK_TRANSFER` or `NAPAS_BANK_TRANSFER` orders that are not captured or cancelled.

## PHP SDK

Official package:

```bash
composer require sepay/sepay-pg
```

Requirements:

- PHP 7.4+
- `ext-json`
- `ext-curl`
- Guzzle

Initialize:

```php
use SePay\SePayClient;
use SePay\Builders\CheckoutBuilder;

$sepay = new SePayClient(
    getenv('SEPAY_MERCHANT_ID'),
    getenv('SEPAY_MERCHANT_SECRET_KEY'),
    SePayClient::ENVIRONMENT_SANDBOX
);
```

Build checkout fields:

```php
$checkoutData = CheckoutBuilder::make()
    ->currency('VND')
    ->orderAmount(100000)
    ->operation('PURCHASE')
    ->orderDescription('Payment for invoice INV_001')
    ->orderInvoiceNumber('INV_001')
    ->successUrl('https://example.com/payment/success')
    ->errorUrl('https://example.com/payment/error')
    ->cancelUrl('https://example.com/payment/cancel')
    ->build();

$formFields = $sepay->checkout()->generateFormFields($checkoutData);
```

Use the SDK-generated fields when possible. Custom forms must keep the documented field order.

## Laravel Packages

Do not present `sepayvn/laravel-sepay` as the official Payment Gateway SDK unless current official SePay docs explicitly identify it for Payment Gateway.

Current guidance:

- Official gateway SDK path: Node.js `sepay-pg-node` and PHP `sepay/sepay-pg`.
- Laravel-specific packages may be wrappers or webhook helpers; inspect the package, maintainer, and current docs before using.
- For Laravel apps, prefer the official PHP SDK plus explicit webhook/IPN handlers unless the user asks for a specific package.

## Error Handling

Handle SDK exceptions by category:

- authentication/signature failure
- validation failure
- not found
- rate limit
- server error

Log:

- environment
- merchant/order invoice number
- gateway method called
- status/error body
- retry-after value if provided

## Production Checklist

- Store gateway credentials in server-side env vars.
- Keep gateway order IDs separate from API v2 bank transaction IDs.
- Use unique `order_invoice_number`.
- Preserve field order for custom signatures.
- Verify IPN separately from checkout redirects.
- Make IPN idempotent by invoice/order/transaction identifiers.
- Backfill unclear states with gateway order detail API.
- Keep card void and QR/bank-transfer cancel rules separate.
