# Source Inventory Và Flow Manifest

## Cổng Source Union Inventory

Trước khi vẽ, build source union inventory từ tất cả input artifacts.

Không tính document title, section title, hoặc broad flow ID là inventory đầy đủ. Mỗi command, guard, store, cursor, terminal state, recovery path, invariant, negative rule, và described relation có liên quan implementation phải là source item riêng.

Inventory các category:

```text
flows
subflows
roads
actors
roles
scopes
screens
UI commands
runtime commands
backend/API commands
IPC commands
permissions
guards
branch conditions
deny conditions
stores
caches
files
DBs
table families
queues
outboxes
projections
snapshots
cursors
checkpoints
receipts
versions
locks
hash chains
coverage markers
lifecycle states
retry rules
rollback rules
terminal states
recovery paths
invariants
negative rules
out-of-scope boundaries
```

Gán một mapping status cho mỗi inventory item:

```text
MAPPED_AS_NODE
MAPPED_AS_EDGE
MAPPED_AS_JUNCTION
MAPPED_AS_TERMINAL
MAPPED_AS_GAP
INTENTIONALLY_OUT_OF_SCOPE
MISSING
```

Fail run nếu bất kỳ implementation-relevant item nào là `MISSING`.

Ghi mapped SVG ids cho mọi item không phải intentionally out of scope.

## Cổng Flow Manifest Trước Khi Render

Sau khi build source union inventory và trước khi vẽ bất kỳ SVG implementation logic nào, tạo source-described flow manifest từ inventory.

Manifest không phải design plan. Nó là danh sách flows, subflows, và roads được source artifacts mô tả trực tiếp.

Với mỗi manifest item, ghi:

```text
flow_id
source_name
category: flow | subflow | road
source_refs
described_entry
described_exit_or_terminal
described_handoffs
missing_handoffs_or_unknown_relations
mapping_status: PENDING | MAPPED | MAPPED_AS_GAP | INTENTIONALLY_OUT_OF_SCOPE | MISSING
planned_or_actual_svg_group_id
```

Rules:

1. Mọi source-described flow, subflow, và road phải có trong manifest trước khi vẽ.
2. Không add expected flows nếu source không mô tả. Nếu expected flow cần thiết nhưng absent, add gap item, không add real flow.
3. Không vẽ flow không có trong manifest.
4. Nếu phát hiện source-described flow mới trong lúc vẽ, dừng vẽ, update source union inventory và manifest, rồi mới tiếp tục flow-by-flow rendering.
5. Manifest item chỉ complete khi map tới visible SVG ids, visible gap, hoặc intentional out-of-scope record.
6. Fail run nếu source-described flow vắng mặt khỏi manifest.
7. Fail run nếu visible implementation flow tồn tại mà không có manifest item.

## Trường Metadata

SVG metadata phải có `flow_manifest` và `source_inventory`.

Ví dụ:

```json
{
  "flow_manifest": [
    {
      "flow_id": "AUTH_FLOW",
      "source_name": "Auth login",
      "category": "flow",
      "source_refs": ["spec.md#auth-login"],
      "described_entry": "User submits login",
      "described_exit_or_terminal": "AUTH_SUCCESS or AUTH_DENIED",
      "described_handoffs": ["AUTH_SUCCESS -> LICENSE_CHECK"],
      "missing_handoffs_or_unknown_relations": [],
      "mapping_status": "MAPPED",
      "planned_or_actual_svg_group_id": "flow-auth-flow"
    }
  ],
  "source_inventory": [
    {
      "source_item": "sync_cursor.last_pulled_relay_id",
      "category": "cursor",
      "source_ref": "spec.md#sync-reconnect",
      "mapping_status": "MAPPED_AS_NODE",
      "mapped_ids": ["node-sync-cursor-last-pulled-relay-id"]
    }
  ]
}
```
