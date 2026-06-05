# UI Slot Map

This file defines exactly where backend data may appear in the approved UI.

If a backend field is not mapped here, it must not be rendered.

---

## Page: Home

### Slot: hero.brandName

Approved text:

```text
AV Cheap Apps
```

Backend binding:

```text
Not allowed
```

Fallback:

```text
AV Cheap Apps
```

Rules:

- Do not change brand name.
- Do not bind this from backend.

---

### Slot: hero.slogan

Approved text:

```text
Always have what you need
```

Backend binding:

```text
Not allowed
```

Fallback:

```text
Always have what you need
```

Rules:

- Do not rewrite marketing copy.

---

### Slot: pricing.priceLabel

Approved text:

```text
$3/month
```

Backend binding:

```text
Allowed
```

Backend source:

```text
GET /api/pricing/current
field: monthlyPriceLabel
```

Fallback:

```text
$3/month
```

Rules:

- Replace only the price label.
- Do not add discount badges.
- Do not add billing explanation text.
- Do not add API/backend status text.

---

### Slot: pricing.trialLabel

Approved text:

```text
30-day free trial
```

Backend binding:

```text
Allowed
```

Backend source:

```text
GET /api/pricing/current
field: trialLabel
```

Fallback:

```text
30-day free trial
```

Rules:

- Replace only the trial label.
- Do not add trial explanation panels.

---

### Slot: appList.items

Approved static behavior:

```text
Render approved app card layout.
```

Backend binding:

```text
Allowed
```

Backend source:

```text
GET /api/apps/public
fields: title, shortDescription, downloadUrl
```

Fallback:

```text
Preserve approved prototype items if backend data is unavailable and no empty state is approved.
```

Rules:

- Use existing card layout only.
- Do not add badges.
- Do not add helper descriptions.
- Do not show internal IDs.
- Do not show API status.
- Do not show “No data available” unless approved in state-map.md.

---

## Page: Pricing

Add page-specific slots here.

---

## Page: Login

Add page-specific slots here.

---

## Page: Payment

Add page-specific slots here.
