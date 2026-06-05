# Backend Contract Map

Map backend endpoints to approved UI slots.

Do not render backend fields that are not mapped to approved slots.

---

## Endpoint: GET /api/pricing/current

Purpose:

```text
Provides current public pricing values.
```

Allowed UI slots:

| Backend Field | UI Slot | Fallback | Notes |
|---|---|---|---|
| monthlyPriceLabel | pricing.priceLabel | $3/month | Replace label only |
| trialLabel | pricing.trialLabel | 30-day free trial | Replace label only |
| paymentUrl | pricing.paymentUrl | /pricing | Used for href only |

Forbidden fields to render:

- internalPlanId
- debugMode
- createdAt
- updatedAt
- apiStatus
- adminNote

---

## Endpoint: GET /api/apps/public

Purpose:

```text
Provides public app listing data.
```

Allowed UI slots:

| Backend Field | UI Slot | Fallback | Notes |
|---|---|---|---|
| title | appList.items.title | Approved prototype title | Text only |
| shortDescription | appList.items.shortDescription | Approved prototype description | Text only |
| downloadUrl | appList.items.downloadUrl | # | href only |

Forbidden fields to render:

- id
- tenantId
- ownerId
- debugLabel
- isMock
- isBeta
- internalStatus
- adminNotes
