# P1-D Raw Graph ↔ Read-only Projection Parity

Date: 2026-07-26 15:40 +07
Target: `E:\cheapapp.org`
Target HEAD: `a869876ab6262dacde6cd5d432d099a91852a646`
Graph: `E:\cheapapp.org\.anvien\graph.json`
Graph SHA-256: `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`
Classification: bounded projection parity; no root-cause or remediation claim

## Question

For the five requested source occurrences—`Promise` at line 13, `Math.max` and `Math.min` at line 191, and the two `readAdminCommercialConfig` calls at lines 142 and 32—does the persisted raw graph agree with the supported read-only `cypher` and `file-detail` projections? The comparison explicitly separates a physical syntax occurrence from the multiple fact-family records that can be emitted for it.

## Boundary and freshness

- The target was read in place. It was not copied into `E:\Anvien`, and no target source, dependency, configuration, test, `.anvien` file, `AGENTS.md`, or `CLAUDE.md` was edited.
- `anvien analyze --force` was **not** rerun in this slice. The supplied fresh graph identity was independently rechecked: raw graph has `84,807` nodes and `114,125` relationships; graph metadata reports `60,861` source-backed unresolved references; SHA-256 and target HEAD above match the supplied identity; graph last write is `2026-07-26T14:49:43.8894011+07:00`.
- The target had `10` pre-existing Git status lines at the identity check. This slice did not change them.
- Probes, command captures, and this report are under `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart` and `E:\Anvien\reports\Investigation` only.
- This is a data/projection comparison. It does not decide whether any unresolved classification is semantically correct, trace the earliest analyzer owner, or authorize a fix.

## Keys and counting policy

The following keys are used throughout the comparison:

1. **Physical-occurrence grouping key:** `filePath + source line + token`. This groups records that refer to one visible source occurrence (or to multiple nested occurrences on one line) only for display.
2. **Canonical source-site key:** the exact persisted `sourceSiteId`, whose form is `SourceSite:<filePath>#<factFamily>#<targetText>#<startLine>#<startCol>#<endLine>#<endCol>`. This is the identity used for duplicate and missing checks.
3. **Raw gap-node key:** `node.id == ResolutionGap:<sourceSiteId>` (or `node.properties.sourceSiteId == sourceSiteId`).
4. **Raw edge key:** `relationship.id`, joined to the fact by `sourceSiteId`/`sourceSiteIds`.

Repeating one `sourceSiteId` on a gap node, its `HAS_RESOLUTION_GAP` edge, and the `file-detail` sample is a representation of one fact, not three source sites. Conversely, a `call` and a `type-reference` with different canonical IDs are counted as two fact records even when they share a line and token.

## Source facts and same-line companions

The target source text was read directly at these locations:

```text
read-admin-commercial-config.ts:13
): Promise<AdminCommercialConfigReadModel> {

email-operations-observability.ts:191
return Math.max(1, Math.min(limit ?? defaultReadLimit, maxReadLimit));

save-admin-commercial-config-mutation.ts:142
const currentState = await readAdminCommercialConfig(input.actor, {

read-admin-commercial-config-route-view.ts:32
const readModel = await readAdminCommercialConfig(input?.actor, options);
```

The five requested occurrences expand to seven canonical fact records because each `readAdminCommercialConfig` call also has a separate `type-reference` record:

| Requested occurrence | Canonical fact records selected | Canonical range(s) |
|---|---|---|
| `read-admin-commercial-config.ts:13`, `Promise` | `type-reference#Promise` | `13:3–13:10` |
| `email-operations-observability.ts:191`, `Math.max` | `call#Math.max` | `191:9–191:71` |
| `email-operations-observability.ts:191`, `Math.min` | `call#Math.min` | `191:21–191:70` |
| `save-admin-commercial-config-mutation.ts:142`, `readAdminCommercialConfig` | `call#readAdminCommercialConfig`; companion `type-reference#readAdminCommercialConfig` | `142:29–144:4`; `142:8–144:4` |
| `read-admin-commercial-config-route-view.ts:32`, `readAdminCommercialConfig` | `call#readAdminCommercialConfig`; companion `type-reference#readAdminCommercialConfig` | `32:28–32:76`; `32:10–32:76` |

Line-level source-site inventory is intentionally larger than the selected facts:

| Physical file/line | All raw source-site IDs at that line | Interpretation |
|---|---:|---|
| `read-admin-commercial-config.ts:13` | 2 | selected `Promise` type reference; separate resolved `AdminCommercialConfigReadModel` type reference |
| `email-operations-observability.ts:191` | 2 | two distinct nested calls: `Math.max` and `Math.min` |
| `save-admin-commercial-config-mutation.ts:142` | 3 | selected call + type-reference companions; separate `input.actor` access is not the call |
| `read-admin-commercial-config-route-view.ts:32` | 3 | selected call + type-reference companions; separate `input.actor` access is not the call |

Thus, a line/name-only count would overstate “duplicate calls.” The canonical key and fact family are required.

## Cardinality parity

The raw probe, seven-ID Cypher predicates, and expanded `file-detail` projections all found exactly one record for every selected canonical source site:

| Canonical fact | Raw `ResolutionGap` nodes | Raw `HAS_RESOLUTION_GAP` edges | Materialized `SourceSite` nodes | Cypher gap-node rows | Cypher CodeRelation rows | `file-detail` matching samples | Edge `sourceSiteCount` | `sourceSiteIds` length |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| `Promise` type reference | 1 | 1 | 0 | 1 | 1 | 1 | 1 | 1 |
| `Math.max` call | 1 | 1 | 0 | 1 | 1 | 1 | 1 | 1 |
| `Math.min` call | 1 | 1 | 0 | 1 | 1 | 1 | 1 | 1 |
| save call | 1 | 1 | 0 | 1 | 1 | 1 | 1 | 1 |
| save type-reference companion | 1 | 1 | 0 | 1 | 1 | 1 | 1 | 1 |
| route call | 1 | 1 | 0 | 1 | 1 | 1 | 1 | 1 |
| route type-reference companion | 1 | 1 | 0 | 1 | 1 | 1 | 1 | 1 |

The count queries returned `row_count: 7` with count `1` for both gap nodes and CodeRelation edges. The generic source-site-node query returned `row_count: 0`. There is no duplicate exact `sourceSiteId` in the selected raw node or edge sets, and no selected site disappears from either supported projection.

As an additional raw integrity check, every selected `ResolutionGap:<sourceSiteId>` target ID has exactly one incoming raw relationship, and that relationship is the same one found by `sourceSiteId`/`sourceSiteIds`; there are zero target-matched edges lacking source-site linkage.

## Raw fact values

All seven raw nodes have label `ResolutionGap` and the same persisted status shape:

| Fact family / site class | `gapKind` | `classification` | `actionability` | `sourceSiteStatus` | `proofKind` | Raw note / edge evidence note |
|---|---|---|---|---|---|---|
| `Promise` | `unresolved_type_reference` | `in_repo_unresolved` | `analyzer_gap` | `unresolved_local_binding` | `none` | `type target not resolved` |
| `Math.max`, `Math.min` | `unresolved_call` | `in_repo_unresolved` | `analyzer_gap` | `unresolved_local_binding` | `none` | `call target not resolved` |
| save/route `readAdminCommercialConfig` calls | `unresolved_call` | `in_repo_unresolved` | `analyzer_gap` | `unresolved_local_binding` | `global-fallback-low-confidence` | `call target matched low-confidence global fallback only` |
| save/route type-reference companions | `unresolved_type_reference` | `in_repo_unresolved` | `analyzer_gap` | `unresolved_local_binding` | `none` | `type target not resolved` |

For every selected edge, raw `sourceSiteId` equals the target gap node’s `sourceSiteId`; `sourceSiteIds` is a one-element array containing that same ID; `sourceSiteCount` is numeric `1`; and the raw edge has no `step` property. Raw node and edge properties contain no explicit `null` values in this slice. Missing `step` is a missing key, not a JSON `null`.

## Field-by-field projection comparison

### `file-detail` expanded unresolved sample

For all seven facts, the matching sample had these ten fields and no `null` values:

| Raw gap-node field | `file-detail` field | Result |
|---|---|---|
| `startLine` | `line` | equal |
| `startCol` | `column` | equal |
| `name` | `targetText` | equal |
| `sourceNodeId` | `sourceSymbol` | equal |
| `gapKind` | `gapKind` | equal |
| `classification` | `classification` | equal |
| `actionability` | `actionability` | equal |
| `proofKind` | `proofKind` | equal |
| `sourceSiteId` | `sourceSiteId` | equal |
| `sourceSiteStatus` | `sourceSiteStatus` | equal |

`file-detail` does not include `filePath` in the unresolved sample, even though the enclosing file is known by the response `target`/`summary`. It also omits `endLine`, `endCol`, `fileHash`, `targetRole`, `note`, `resolutionSource`, `source`, app/function-area fields, the raw gap-node ID, and all edge fields. Those omissions are projection shape boundaries, not missing raw records. The unresolved cap was set to `10000`; each enclosing file’s total was below that cap (`1`, `379`, `87`, and `19`), so the seven matching samples were not sample-cap misses.

### Cypher gap-node projection

The seven-ID gap-node query returned all 28 requested raw node properties for all seven rows (`row_count: 7`). The CLI’s JSON `rows` representation serializes every scalar—including `startLine`, `startCol`, `endLine`, `endCol`, `count`, and status values—as a string. This is output-type normalization, not a raw graph type change. The values and canonical IDs match the raw node records.

### Cypher CodeRelation projection

The supported relationship table is `CodeRelation` with a `type` property. Each selected raw edge has `21` present top-level keys (no explicit nulls; `step` is absent), while the seven-ID relationship query returned ten columns and seven rows:

`sourceId`, `targetId`, `relationshipType`, `confidence`, `reason`, `step`, `resolutionSource`, `evidence`, `fileHash`, and the joined target-node `sourceSiteId`.

The following field behavior is material:

- Raw numeric `confidence: 1` is rendered as the CLI string `"1.000000"`.
- Raw `evidence` is an array of objects; the CLI row renders it as a JSON string.
- Raw `step` is absent on every `HAS_RESOLUTION_GAP` edge; Cypher returns `step: "0"` (the relationship-table default).
- The relationship query cannot directly return raw edge `id`, `sourceSiteId`, `sourceSiteIds`, `sourceSiteCount`, `sourceSiteStatus`, `targetRole`, `targetText`, file/range fields, or the edge’s own `sourceSite` metadata. The source-site join in this query is `g.sourceSiteId` from the target gap node.
- `file-detail`’s unresolved sample does not expose the edge record; edge parity therefore comes from raw graph + CodeRelation Cypher, not from the file-detail sample.

### Source-site materialization

There are zero `SourceSite`-label nodes for all seven IDs in raw graph and Cypher. A source site is persisted here as a string field (`sourceSiteId`/`sourceSiteIds`) on a `ResolutionGap` node and its edge, plus the derived gap-node ID. Treating those repeated strings as separate graph nodes would manufacture duplicates.

## Unsupported boundaries captured

The following read-only probes were run and captured:

| Probe | Result |
|---|---|
| `MATCH (s)-[r:HAS_RESOLUTION_GAP]->(g) ...` | Binder error: `Table HAS_RESOLUTION_GAP does not exist`; use `CodeRelation` and filter `r.type` instead |
| `MATCH (s)-[r]->(g) ... RETURN r.id` | Binder error: `Cannot find property id for r`; raw relationship IDs are not exposed as Cypher properties |
| `MATCH (s:SourceSite) RETURN count(s)` | Binder error: `Table SourceSite does not exist`; no materialized source-site label/table |
| Generic `MATCH (s) WHERE s.id = '<sourceSiteId>' RETURN count(s)` | Valid; returns `0`, confirming no node with that ID |

The schema resource `anvien://repo/cheapapp-accuracy-direct/schema` documents canonical nodes plus a single `CodeRelation` table, but does not enumerate the `ResolutionGap`/source-site fields used by these extra audit records. The raw JSON remains the authority for those fields; Cypher is a partial projection.

## Exact commands and artifacts

The exact command construction and seven canonical IDs are recorded in:

- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1d-command-manifest.md`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1d-cypher-run.ps1`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1d-raw-graph-probe.mjs`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1d-parity-probe.mjs`

Primary captures:

- `p1d-raw-graph-probe-output.json` — selected raw nodes, edges, related source-site IDs, and graph counts.
- `p1d-parity-probe-output.json` — machine-readable cardinality, null/key inventory, physical-line grouping, field mapping, and Cypher capture summaries.
- `p1d-file-detail-read-admin-full.json`
- `p1d-file-detail-email-observability-full.json`
- `p1d-file-detail-save-admin-full.json`
- `p1d-file-detail-route-view-full.json`
- `p1d-cypher-gap-nodes.txt`, `p1d-cypher-gap-rels.txt`, `p1d-cypher-node-counts.txt`, `p1d-cypher-rel-counts.txt`, `p1d-cypher-sourcesite-counts.txt`
- `p1d-cypher-unsupported-rel-label.txt`, `p1d-cypher-unsupported-rel-id.txt`, `p1d-cypher-unsupported-sourcesite-label.txt`

## Bounded result

For these exact seven canonical source-site facts, raw graph, Cypher gap-node/CodeRelation projections, and expanded file-detail unresolved samples agree on presence and one-to-one cardinality. No projection-layer duplicate or missing record was observed. The projections intentionally differ in shape: file-detail is a compact unresolved-sample view, Cypher is a typed-table view with scalar/string normalization and a default `step`, and raw JSON retains the complete source-site/edge metadata. Any later accuracy or root-cause conclusion must preserve the canonical `sourceSiteId + factFamily` key and must not infer source-site multiplicity from line-only rows or repeated embedded IDs.

This report stops at projection parity. It does not trace why the seven facts received their unresolved classifications, does not accept a root cause, and does not change code or tests.
