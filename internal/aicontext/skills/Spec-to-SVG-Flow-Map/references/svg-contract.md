# Contract SVG

## Contract SVG Ngữ Nghĩa

SVG phải là valid XML và vẫn đọc được như source code.

Mỗi node phải theo shape:

```xml
<g
  id="node-..."
  data-type="..."
  data-flow="..."
  data-status="..."
  data-lane="..."
  data-source-ref="..."
>
  <title>...</title>
  <desc>...</desc>
  <text>...</text>
</g>
```

Mỗi edge phải theo shape:

```xml
<g
  id="edge-..."
  data-type="edge"
  data-edge-type="..."
  data-from="..."
  data-to="..."
  data-condition="..."
  data-flow="..."
  data-source-ref="..."
>
  <title>...</title>
  <desc>...</desc>
  <path ... />
  <text>...</text>
</g>
```

Không tạo anonymous paths cho important logic.

Bad:

```xml
<path d="M120 80 L300 80" />
```

Good:

```xml
<g
  id="edge-auth-success-to-license-check"
  data-type="edge"
  data-edge-type="CONTROL_FLOW"
  data-from="AUTH_LOGIN"
  data-to="LICENSE_CHECK"
  data-condition="auth_success"
  data-flow="AUTH_FLOW"
  data-source-ref="spec.md#auth-login"
>
  <title>AUTH_LOGIN -> LICENSE_CHECK</title>
  <desc>User credentials are valid. Continue to license validation.</desc>
  <path d="M120 80 L300 80" />
  <text>if auth_success</text>
</g>
```

## Metadata SVG

Mỗi SVG phải có một machine-readable metadata block.

Metadata phải gồm graph content và source coverage:

```xml
<metadata id="spec-flow-map">
{
  "feature": "<feature-name>",
  "source_spec": "<source-spec-path>",
  "version": "draft",
  "reference_artifacts": [
    {
      "path": "docs/flow-maps/reference.flow.svg",
      "type": "reference_svg",
      "scope": "<scope>"
    }
  ],
  "nodes": [
    {
      "id": "AUTH_LOGIN",
      "type": "SYSTEM_ACTION",
      "lane": "AUTH_SECURITY",
      "flow": "AUTH_FLOW",
      "status": "defined",
      "source_ref": "spec.md#auth-login"
    }
  ],
  "edges": [
    {
      "id": "edge-auth-success-to-license-check",
      "type": "CONTROL_FLOW",
      "from": "AUTH_LOGIN",
      "to": "LICENSE_CHECK",
      "condition": "auth_success",
      "flow": "AUTH_FLOW",
      "source_ref": "spec.md#auth-login"
    }
  ],
  "flow_manifest": [],
  "source_inventory": [],
  "coverage_summary": {
    "total_source_items": 0,
    "mapped_items": 0,
    "missing_items": 0,
    "collapse_violations": 0
  },
  "missing_source_items": [],
  "collapse_violations": [],
  "gaps": [],
  "terminal_states": [],
  "junctions": [],
  "legend": {
    "visible_legend_id": "legend-end",
    "shape_meaning": {},
    "color_meaning": {},
    "arrow_meaning": {}
  }
}
</metadata>
```

Metadata phải khớp visible SVG nodes và edges.

Nếu `missing_source_items` không rỗng, đặt implementation status là `BLOCKED`.

## Lanes, Nodes, Edges

Chỉ dùng lanes liên quan đến spec slice hiện tại.

Suggested lanes:

```text
USER
UI_APP
LOCAL_STATE
LOCAL_DB
SYNC_ENGINE
BACKEND_API
AUTH_SECURITY
PAYMENT_LICENSE
EMAIL_NOTIFICATION
ERROR_RECOVERY
EXTERNAL_SYSTEM
OWNER_DECISION
```

Mỗi node phải thuộc đúng một lane.

Node types:

```text
START
END
EVENT
SCREEN
USER_ACTION
SYSTEM_ACTION
BACKGROUND_JOB
DECISION
STATE
DATA_STORE
DOCUMENT
EXTERNAL_SYSTEM
JUNCTION
ERROR
RECOVERY
SPEC_GAP
UNDEFINED_BEHAVIOR
OWNER_DECISION_REQUIRED
OUT_OF_SCOPE
```

Edge types:

```text
CONTROL_FLOW
DATA_FLOW
ASYNC_EVENT
SYNC_FLOW
ERROR_FLOW
RECOVERY_FLOW
HANDOFF
BLOCKED_BY
OUT_OF_SCOPE_REFERENCE
```

Mỗi edge phải có `from`, `to`, `condition`, `data-edge-type`, `flow`, và `data-source-ref`.

## Contract Visual

Dùng meaning visual ổn định:

```text
START / END                  = circle
SCREEN                       = rounded rectangle
USER_ACTION                  = rounded rectangle
SYSTEM_ACTION                = rectangle
BACKGROUND_JOB               = dashed rectangle
DECISION                     = diamond
STATE                        = pill / capsule
DATA_STORE                   = cylinder
DOCUMENT                     = document icon/shape
EXTERNAL_SYSTEM              = dashed rectangle
JUNCTION                     = small circle
ERROR                        = red rounded rectangle
RECOVERY                     = orange rounded rectangle
SPEC_GAP                     = thick red border rectangle
UNDEFINED_BEHAVIOR           = thick red border diamond
OWNER_DECISION_REQUIRED      = thick orange border rectangle
OUT_OF_SCOPE                 = gray dashed rectangle
```

Dùng meaning màu ổn định:

```text
Green       = normal / automated / value-producing flow
Blue        = data / storage / source of truth
Purple      = auth / security / permission boundary
Yellow      = delay / pending / waiting
Red         = bottleneck / error / danger / spec break
Orange      = owner decision / unclear business rule
Gray        = manual / external / out of scope
Black       = default control flow
Dashed line = async / background / indirect dependency
Solid line  = direct flow
Double line = sync / bidirectional data exchange
```

Không đổi meaning màu theo từng diagram.

## Yêu Cầu Legend Visible Ở Cuối SVG

Mỗi generated SVG phải kết thúc bằng visible, human-readable legend section.

Đặt legend là visible SVG group cuối cùng trước `</svg>`:

```xml
<g
  id="legend-end"
  data-type="legend"
  data-position="end"
>
  <title>Legend</title>
  <desc>Explains node shapes, colors, line styles, and arrow types used in this SVG.</desc>
  ...
</g>
```

Legend phải đọc được mà không cần mở markdown report. Legend phải giải thích:

1. Mỗi node shape/box có nghĩa gì.
2. Mỗi màu có nghĩa gì.
3. Mỗi line style có nghĩa gì.
4. Mỗi arrow type có nghĩa gì.
5. Error, recovery, gap, owner-decision, out-of-scope, terminal, junction, async, sync, data-flow, và control-flow markers có nghĩa gì.

Legend bắt buộc cả khi SVG lớn hoặc split thành nhiều visual regions.

Không đặt legend chỉ trong metadata. Metadata phải reference visible legend id, nhưng SVG canvas phải có legend visible đọc được.
