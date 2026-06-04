# SaaS Multi-Tenant Patterns Reference

> Đọc file này khi dự án là SaaS platform với multiple accounts/tenants.

## Tenant Isolation Models

| Model | Mô tả | Phù hợp khi |
|-------|-------|-------------|
| Row-level | Mọi table có `tenant_id` | Nhỏ, đơn giản |
| Schema-per-tenant | Mỗi tenant 1 DB schema | Compliance, isolation cao |
| DB-per-tenant | Mỗi tenant 1 DB | Enterprise, full isolation |

### Row-level (recommended cho MVP)

```sql
-- Mọi table có tenant_id
ALTER TABLE users ADD COLUMN tenant_id UUID NOT NULL;

-- RLS (Row Level Security) nếu dùng PostgreSQL
ALTER TABLE users ENABLE ROW LEVEL SECURITY;
CREATE POLICY tenant_isolation ON users
  USING (tenant_id = current_setting('app.tenant_id')::uuid);
```

## Data Tuple Chuẩn (SaaS)

```
owner_id + app_type + app_scope_id
```

- `owner_id`: account/tenant sở hữu resource
- `app_type`: loại application (web, mobile, api, ...)
- `app_scope_id`: scope trong app đó (workspace, project, ...)

## Entitlement & Billing

```markdown
### Plan → Feature Mapping
| Feature Flag | Free | Pro | Enterprise |
|-------------|------|-----|------------|
| max_users   | 3    | 25  | unlimited  |
| api_calls   | 1000/mo | 50000/mo | custom |
| custom_domain | No | Yes | Yes |

### Enforcement Points
- API gateway: rate limit theo plan
- Feature gates: check entitlement trước khi expose feature
- Billing webhooks: Stripe → update subscription status

### Subscription Lifecycle
active → past_due (payment fail) → canceled
active → trialing → active (trial convert)
```

## Admin Governance

```markdown
### Admin Roles
| Role | Quyền |
|------|-------|
| super_admin | Toàn bộ |
| tenant_admin | Trong tenant |
| billing_admin | Billing only |
| read_only | View only |

### Audit Log Requirements
- Mọi write action phải có audit trail
- Fields: actor_id, tenant_id, action, resource_type, resource_id, before, after, ip, timestamp
- Retention: 1 year minimum
- Immutable: không cho phép delete/update audit logs
```