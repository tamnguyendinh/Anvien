# Checklist Chi Tiết Theo Domain

Khi relevant, model rõ các domain này thay vì group chung:

- Auth/session/device/entitlement restore, login handoff, cache, keyring, expiry, logout, denied states.
- Restaurant/scope selection, visible/bound scopes, DB mount, hydrated/not-hydrated states.
- Permission/local command gate, renderer boundary, Go/backend authority, command guards.
- Owner/app setup, setup outbox, setup version, setup receipt, convergence state.
- Business runtime: POS order, pay/refund, move/merge/split, shift/cash, inventory/stocktake.
- Sync transport: LAN/WSS/VPS, relay, delta, dedupe, ack, cursor, hash verify, gap, repair.
- Snapshot/bootstrap/manual sync: baselines, manifest, anchors, allowlist apply, rollback, catchup.
- Reports/coverage: aggregate/detail split, retention, source coverage, export/print gating.
- Local print/export: printer config, preview, spooler result, local-only boundary.
- Lifecycle/reconnect: active/idle/sleep/offline/resume, missed relay, auth refresh, cursor catchup.
- Local settings/device-only behavior.
- External/backend contracts và denied/error responses.

Nếu source có detail thuộc các domain trên, mọi detail implementation-relevant phải vào source inventory và phải map thành node, edge, junction, terminal, hoặc gap.

Nếu source không mô tả detail mà detail đó cần để implementation an toàn, tạo gap. Không được tự điền hành vi.
