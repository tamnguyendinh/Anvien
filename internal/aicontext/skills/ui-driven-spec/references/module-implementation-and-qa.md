# Module Implementation And Runtime QA

> This file is part of UI-Driven Spec Workflow. Read when: drafting single-module implementation plans or conducting runtime QA and sign-off per module.

---

## 1. Single-Module Implementation Planning

**⛔ STRICTLY PROHIBITED to plan multiple modules simultaneously — even if requested by the owner.**
If pushed → explain: batch planning leads to dependency hell and boundless, untraceable QA scope.

**Prerequisite:** Only begin after receiving explicit approval from the Owner Review Gate.

**Default Implementation Order:** Auth first → module with highest FE dependencies → remaining modules by priority.

### Per-Module Plan Template

```markdown
## Implementation Plan: Module [Name]

### Scope (derived from backend-contract-map.md)
Endpoints: [list]
DB tables: [list]
Business rules: [list — cited from Spec §X]

### Tasks
1. DB migration
2. Model / schema
3. Service layer (business rules)
4. Controller / handler (endpoints)
5. Middleware (auth, validation)
6. Unit tests
7. Integration tests
8. Wire FE: replace mock handlers with real API integration

### Definition of Done
- [ ] Endpoints match path/method/shape defined in api-contract.md
- [ ] Response matches TypeScript types of FE
- [ ] Business rules strictly conform to Spec
- [ ] Unit + integration tests pass
- [ ] FE renders correctly on the real UI with live API integration
- [ ] Runtime QA passes
```

---

## 2. Runtime QA After Each Module

**Failing QA = Strictly prohibited from proceeding to the next module.**

```markdown
## Runtime QA: Module [Name]

### Happy path
- [ ] Endpoints return correct data → UI renders into correct slot (verify with slot-map.md)
- [ ] Actions trigger correct state transitions (verify with state-map.md)

### Error handling
- [ ] Network error → error banner renders (application does not crash)
- [ ] Validation error → field-level message displays
- [ ] 401 Unauthorized → redirect to login
- [ ] 403 Forbidden → appropriate access-denied message
- [ ] 404 Not Found → empty state renders

### Edge cases
- [ ] Empty list → empty state slot renders
- [ ] Pagination boundary (1 item, max items)
- [ ] Concurrent actions produce no race conditions

### Performance
- [ ] Loading state visible during data fetching
- [ ] No flash of unstyled or empty content

### Sign-off
- [ ] Pass → Log to docs/qa-log.md → notify owner → receive approval → proceed to next module plan
- [ ] Fail → Log bug, implement fix, re-run QA — do NOT proceed to next module
```
