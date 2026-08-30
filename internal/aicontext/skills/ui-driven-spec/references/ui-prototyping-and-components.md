# UI Prototyping And Frontend Components

> This file is part of UI-Driven Spec Workflow. Read when: building the static HTML prototype or developing frontend components with mock data.

---

## 1. Prototype UI (HTML)

**Purpose:** Create the user interaction map. Static HTML, no logic.

**Input questions to ask the user:**
- List of screens / flows to cover
- Priority order (which flow is most critical)
- Existing design system (colors, fonts, component library)

**Output structure:**
```text
prototype/
├── index.html          ← navigation between screens
├── screen-{name}.html  ← 1 file per screen
└── assets/             ← static css, images
```

**Checklist:**
- [ ] Every user action (click, submit, navigate) is visible
- [ ] Empty states, loading states, and error states are mocked
- [ ] Happy path + at least 1 error path per flow
- [ ] No real logic — HTML/CSS only

---

## 2. Frontend Components

**Purpose:** Convert the prototype into real components with state management.
Data is hardcoded / mocked — **no real API calls yet**.

**Output structure:**
```text
src/
├── components/
├── pages/
├── hooks/              ← mock data fetching
├── types/              ← TypeScript types defining data shapes
└── mocks/              ← mock data reflecting expected API response
```

**`types/` and `mocks/` are the source of truth for contract extraction:**

```typescript
// types/order.ts — what FE expects BE to return
export interface Order {
  id: string
  status: 'pending' | 'confirmed' | 'cancelled'
  items: OrderItem[]
  total: number
  createdAt: string
}
```

**Checklist:**
- [ ] Every component renders correctly from mock data
- [ ] TypeScript types are complete for every entity
- [ ] Mock data matches the exact shape FE needs from BE
- [ ] No `any` types at data boundaries
