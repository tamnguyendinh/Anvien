# Technical Standards Reference

> Chi tiết đầy đủ cho Bước 6 — Technical Standards.
> SKILL.md chứa phần tóm tắt; đọc file này khi cần output hoàn chỉnh.

## 6.1 Git Workflow

```markdown
### Branching Strategy
- main: production-ready, protected, require PR
- develop: integration branch (gitflow) — tuỳ chọn
- feature/{ticket-id}-{short-desc}
- fix/{ticket-id}-{short-desc}
- hotfix/{desc}: vá nhanh production

### Commit Convention (Conventional Commits)
<type>(<scope>): <subject>

Types: feat | fix | chore | docs | refactor | test | perf | ci | revert
Body: giải thích WHY, không phải WHAT
Footer: BREAKING CHANGE: ... hoặc Closes #123

Examples:
  feat(auth): add JWT refresh token rotation
  fix(billing): correct proration on plan downgrade
  refactor(db): extract query builder to separate module
  chore(deps): upgrade prisma to 5.14

### PR Rules
- Squash merge vào main/develop
- PR title = commit message format
- Min 1 reviewer approve
- CI xanh trước khi merge
- Không merge WIP PR — dùng Draft PR thay thế
- Link issue trong PR description

### Branch Protection Rules (GitHub)
- Require PR review: 1
- Require status checks: lint, test, build
- No force push to main
- Auto-delete head branch after merge
```

## 6.2 Error Handling

```markdown
### Phân loại lỗi
1. **Operational Error**: expected, predictable
   - Validation failed, not found, unauthorized, rate limited
   - Xử lý: catch → log warn → trả error response chuẩn
   
2. **Programmer Error**: bug, không nên xảy ra
   - Null pointer, wrong type, unhandled promise rejection
   - Xử lý: log error + full stack trace → alert → (optional) restart process

### Rules
- KHÔNG swallow errors: `catch (e) {}` là forbidden
- KHÔNG expose stack trace ra client (production)
- KHÔNG throw raw DB errors ra ngoài service layer
- Centralized error handler tại top-level (Express: `app.use(errorHandler)`)
- Async: wrap tất cả async route handlers, không để unhandled rejection

### Error Mapping
| Error Type | HTTP Status | Error Code |
|-----------|-------------|------------|
| Validation | 400 | VALIDATION_ERROR |
| Auth required | 401 | UNAUTHORIZED |
| Permission denied | 403 | FORBIDDEN |
| Not found | 404 | NOT_FOUND |
| Business rule violation | 422 | BUSINESS_ERROR |
| Duplicate / state conflict | 409 | CONFLICT |
| Rate limited | 429 | RATE_LIMITED |
| Internal | 500 | INTERNAL_ERROR |
```

## 6.3 Logging

```markdown
### Log Format (Structured JSON — bắt buộc)
{
  "timestamp": "2024-01-15T10:30:00.000Z",
  "level": "info",
  "service": "api",
  "version": "1.2.3",
  "trace_id": "01HX...",
  "span_id": "abc123",
  "user_id": "uuid | null",
  "tenant_id": "uuid | null",
  "method": "POST",
  "path": "/api/v1/orders",
  "status": 201,
  "duration_ms": 45,
  "message": "Order created",
  "context": { "order_id": "..." }
}

### Log Levels
- error: điều không nên xảy ra, cần điều tra ngay
- warn: unexpected nhưng hệ thống vẫn chạy (deprecated API, retry thành công)
- info: business events quan trọng (request in/out, job start/complete, auth events)
- debug: chi tiết internal — TẮT production

### Trace ID / Correlation ID
- Sinh tại entry point (API gateway hoặc first middleware)
- Format: ULID hoặc UUID v7 (sortable)
- Truyền qua header: `X-Trace-ID`
- Đưa vào mọi log entry trong scope request đó
- Đưa vào error response về client (để support tra cứu)

### Không log
- Passwords, tokens, credit card numbers
- PII đầy đủ (chỉ log id, không log full name + email cùng lúc)
- Request body đầy đủ (chỉ log relevant fields)
```

## 6.4 Caching Strategy

```markdown
### Cache Layers
| Layer | Tool | TTL Default | Invalidation |
|-------|------|-------------|--------------|
| HTTP Response | CDN (Cloudflare) | 5 min | purge by tag |
| Application | Redis | 60s – 1h | explicit on write |
| In-process | LRU (node-lru-cache) | 30s | TTL only |

### Cache Key Format
`{service}:{version}:{entity}:{id}:{variant}`
Examples:
  api:v1:user:uuid123:profile
  api:v1:product-list:page2:sort-price

### Rules
- Cache READ-heavy, stable data. Không cache mutation results.
- Invalidate explicitly khi data thay đổi — đừng chỉ dựa TTL cho mutable data
- Ghi rõ TTL và LÝ DO trong code comment
- Cache miss phải safe (gracefully fetch từ DB)
- Thundering herd: dùng lock hoặc stale-while-revalidate cho cache phổ biến

### Redis Key Expiry Policy
- maxmemory-policy: allkeys-lru (nếu Redis chỉ làm cache)
- Không dùng PERSIST cho cache keys
```

## 6.5 Conflict Resolution

```markdown
### Optimistic Locking
- Thêm cột `version INTEGER NOT NULL DEFAULT 0`
- Mỗi UPDATE: WHERE version = :expected AND SET version = version + 1
- Nếu 0 rows affected → trả 409 CONFLICT

### Idempotency
- Mọi mutation endpoint nhận header: `Idempotency-Key: <client-uuid>`
- Backend: lưu (idempotency_key, response) trong Redis/DB, TTL 24h
- Nếu duplicate key trong TTL → trả lại response đã lưu, không execute lại

### Distributed Lock (Redis)
```javascript
// SET key value NX PX 30000  — atomic lock
const lock = await redis.set(`lock:${resourceId}`, requestId, 'NX', 'PX', 30000);
if (!lock) throw new ConflictError('Resource locked');
try {
  // critical section
} finally {
  // release chỉ nếu chính mình giữ lock
  if (await redis.get(`lock:${resourceId}`) === requestId) {
    await redis.del(`lock:${resourceId}`);
  }
}
```
```

## 6.6 Background Jobs

```markdown
### Job Naming
{module}.{entity}.{verb}
Examples: email.welcome.send | billing.invoice.generate | report.monthly.compile

### Job Contract
- Input: plain serializable JSON — KHÔNG truyền class instance
- Idempotent: chạy lại N lần = kết quả giống chạy 1 lần
- Timeout: define per job (default 60s)
- Retry: 3 lần, exponential backoff (delay: 1s, 5s, 25s)
- Dead Letter Queue: sau max retry → DLQ + Slack/PagerDuty alert

### Job Logging
- Log khi: start, complete, fail (với error + attempt number)
- Include: job_id, queue, attempt, duration_ms, payload summary (không log PII)

### Queue Monitoring
- Queue depth alert: > 1000 pending jobs
- Worker heartbeat: mỗi 30s
- Failed job rate alert: > 5% trong 5 phút
```

## 6.7 Blockchain / Web3 (nếu trong scope)

```markdown
### Smart Contract Standards
- ERC-20 / ERC-721 / ERC-1155 — chọn standard phù hợp
- Audit: bắt buộc trước production (Slither tự động + manual review)
- Upgradability: Proxy pattern (OpenZeppelin) nếu cần

### On-chain vs Off-chain Decision
- On-chain: ownership, transfers, immutable records
- Off-chain: metadata, search, UX data (lưu DB + IPFS hash on-chain)

### Gas Optimization
- Batch operations thay vì multiple calls
- Events thay vì storage cho historical data
- Mapping > array cho lookups

### Wallet Integration
- EIP-1193 provider interface
- Multi-wallet: RainbowKit / ConnectKit
- Sign messages: EIP-712 (typed structured data)

### Testing
- Hardhat / Foundry — local network
- Fork mainnet để test với real contracts
- Coverage: 100% cho contract logic
```