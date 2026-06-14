# SePay VietQR Generation

Use this reference to render VietQR images for bank-transfer instructions. QR generation does not confirm payment; pair it with Webhooks and API v2 reconciliation, or use Payment Gateway if you need hosted checkout.

Official docs checked: 2026-06-14.

## Endpoint

```text
https://qr.sepay.vn/img
```

Example:

```text
https://qr.sepay.vn/img?acc=0123456789&bank=Vietcombank&amount=100000&des=ORDER123
```

## Parameters

Required:

| Parameter | Meaning |
|-----------|---------|
| `acc` | Account number or VA value, depending on the bank/VA model |
| `bank` | Bank identifier |

Common optional parameters:

| Parameter | Meaning |
|-----------|---------|
| `amount` | VND amount; omit only when the customer may choose the amount |
| `des` | Transfer memo/content |
| `template` | Image template such as default, compact, qronly, or standee where supported |
| `download` | `true` to download the image |
| `showinfo` | Controls whether extra transfer information is shown |
| `fullacc` | Shows full account information where supported |
| `holder` | Account holder display value |
| `store` | Store/merchant display value |

`bank` can use supported short name, alias, code, or BIN. For dynamic apps, fetch and cache the bank list rather than hardcoding a stale list:

```text
GET https://qr.sepay.vn/banks.json
```

## QR Construction

```javascript
export function buildSePayQrUrl({ account, bank, amount, description, template }) {
  const url = new URL('https://qr.sepay.vn/img');
  url.searchParams.set('acc', account);
  url.searchParams.set('bank', bank);
  if (amount != null) url.searchParams.set('amount', String(amount));
  if (description) url.searchParams.set('des', description);
  if (template) url.searchParams.set('template', template);
  return url.toString();
}
```

Rules:

- URL-encode descriptions through `URLSearchParams` or a framework equivalent.
- Use integer VND amounts.
- Do not trust a displayed QR as payment evidence.
- Show manual transfer details next to the QR so customers can recover if the QR fails.
- Add cache-control around generated URLs only if the account, amount, and memo are immutable for that order.

## Bank And VA Rules

### Official VA Required

Some banks require an official VA for automatic matching. In that case:

- `acc` should be the official VA value from SePay/bank.
- `des` should include the order code or required memo.
- Do not substitute the source bank account number if the bank requires a VA.

Example:

```text
https://qr.sepay.vn/img?acc=VQRQ12345678&bank=OCB&amount=100000&des=DH001%20thanh%20toan
```

### Memo-Based VA

Memo-based VAs require the transfer memo to contain `TKP` plus the VA code.

Example for VA `001`:

```text
https://qr.sepay.vn/img?acc=0987654321&bank=TPBank&amount=200000&des=TKP001%20DH001
```

### No VA / Memo-Only Matching

If the integration uses a normal bank account and payment-code recognition:

- `acc` is the linked account number.
- `des` should contain the configured payment code.
- Prefer a concise, bank-safe code over long natural-language memos.
- Use Webhooks `code` when SePay extracts the payment code.

Example:

```text
https://qr.sepay.vn/img?acc=0123456789&bank=Vietcombank&amount=150000&des=ORD12345
```

### VietinBank `SEVQR`

For VietinBank personal and household-business accounts, the transfer memo must contain `SEVQR` or transaction notifications may not be pushed to SePay.

Example:

```text
https://qr.sepay.vn/img?acc=0123456789&bank=VietinBank&amount=100000&des=SEVQR%20DH001
```

## Templates

Common template choices:

- default/full QR with bank and account information
- compact
- `qronly`
- `standee`

Choose based on display surface. For invoices, emails, and receipts, prefer stable dimensions and alt text. For kiosk or print flows, test the actual scanner path.

## Error And Fallback Handling

- Keep bank transfer text instructions visible when the image fails.
- Provide copy buttons for account, bank, amount, and memo.
- Treat missing QR image as a UI/display failure, not as a payment failure.
- Monitor QR image errors separately from payment confirmation errors.
- Do not retry payment confirmation from the QR endpoint; confirm through Webhooks/API v2 or Payment Gateway IPN/order APIs.

## Cross-Surface Links

- `webhooks.md`: confirm bank-account transfers in real time.
- `api.md`: reconcile missed or ambiguous bank transactions and manage Order VAs.
- `payment-gateway.md`: use hosted checkout instead of manually displaying bank-transfer QR.
