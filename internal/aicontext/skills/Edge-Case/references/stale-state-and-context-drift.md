# Attack Guide: Stale State & Context Drift

> This file is part of Edge-Case Review Skill v2. Use when: testing stale client caches, store drift after context/branch switching, or cross-tenant/cross-scope data leaks.

---

## Attack Surface Overview

Stateful frontends and caching proxies often hold onto data across tenant switches, account logouts, or project switching. This leads to severe security risks like cross-tenant data contamination, ghost UI elements, and submitting mutations against the wrong target.

---

## Targeted Invariants

- **Complete State Purge on Context Switch**: Switching active tenant, project, branch, or account must completely purge all local memory, stores, and caches.
- **Tenant / Owner Isolation**: Every query and mutation must enforce `owner_id` / `scope_id` scoping at the storage query boundary.
- **Zero Cross-Scope Contamination**: Cached entities from Tenant A must never be visible to Tenant B under any circumstance.
- **Cache Invalidation on Mutation**: Writing updates to an entity must immediately invalidate or update cached read models.

---

## Hostile Perturbations & Attack Scenarios

### 1. Rapid Context / Tenant Switch
- Login to `Tenant A`, navigate through data tables, populate local cache.
- Switch directly to `Tenant B` without refreshing the page.
- *Check:* Does `Tenant B` display any flash of unpurged data from `Tenant A`? Does any API request from `Tenant B` send `Tenant A` IDs?

### 2. Fast Navigation Back/Forward Cache (BFCache)
- View sensitive data in Account A.
- Log out.
- Hit the browser "Back" button.
- *Check:* Does the UI render sensitive user data from the browser cache? Are mutation buttons functional?

### 3. Cross-Scope ID Injection
- Authenticate as User X with access only to `Scope 1`.
- Submit a request to update an entity belonging to `Scope 2` by manually overriding the `id` or `scope_id` parameter in the payload.
- *Check:* Does the backend verify ownership at the query level (`WHERE id = ? AND scope_id = ?`) or allow cross-scope modification?

### 4. Mutation Followed by Stale Cache Read
- Update entity `E1` from value `A` to value `B`.
- Immediately query `E1` with high concurrency.
- *Check:* Does any replica or cache return stale value `A` past the allowed SLA?

---

## Pass / Fail Checklist

- [ ] Complete store reset occurs on logout or tenant/scope switch.
- [ ] Cross-scope ID injection fails with 403 Forbidden or 404 Not Found.
- [ ] Browser history navigation does not expose authenticated data post-logout.
- [ ] Cache invalidation prevents ghost mutations against stale snapshots.
