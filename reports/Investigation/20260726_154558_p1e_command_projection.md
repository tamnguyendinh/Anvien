# P1-E Context, Impact, and Bounded Derived Projection

Date: 2026-07-26 15:45 +07
Target: `E:\cheapapp.org`
Graph: `E:\cheapapp.org\.anvien\graph.json`
Graph SHA-256: `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`
Classification: bounded downstream behavior; no global command/process claim

## Question

Do `context` and `impact` add, remove, or distort facts for the fixed P1-C cases, or do they faithfully expose the upstream graph—including upstream omissions?

## Independent source baseline

- The wrapper `readAdminCommercialConfig` is declared at `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts:10`.
- It is re-exported at `modules/commercial-config/server/admin-commercial-config/index.ts:21`.
- Independent source search finds exactly two wrapper call sites: `save-admin-commercial-config-mutation.ts:142` and `read-admin-commercial-config-route-view.ts:32`.
- `boundedReadLimit` is declared at `modules/email/server/operations/email-operations-observability.ts:190` and has one source call at line 502.

## `context` for the wrapper

Command: `anvien context symbol readAdminCommercialConfig --file modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts --repo E:\cheapapp.org --json`

Observed:

- The selected function identity and range are correct (`10..23`), including return type `Promise<AdminCommercialConfigReadModel>`.
- The file summary reports fan-in `2`, fan-out `11`, and one source-resolution gap for `Promise` at line 13.
- The only incoming symbol-level edge is the barrel File's `USES`/re-export edge. The two consumer files are not incoming `CALLS` edges.
- Outgoing calls/uses inside the wrapper are represented: `readAdminCommercialConfigReadModel`, `AdminCommercialConfigActor`, `AdminCommercialConfigReadOptions`, and `AdminCommercialConfigReadModel`.
- No process is returned for this bounded symbol. This is recorded as an observed cap/result, not a claim that the product has no process.

Comparison: the command does not invent the missing consumers; it exposes the incomplete upstream binding already present in the graph. The `Promise` gap is also carried through from the source diagnostic.

## `impact` for the wrapper

Command: `anvien impact symbol --uid Function:modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts:readAdminCommercialConfig --repo E:\cheapapp.org --direction upstream --json`

Observed:

- Risk: `LOW`; impacted count `7`; direct count `1`; affected modules `0`; affected processes `0`.
- Depth 1 reaches only the barrel File through `scope-finalize import-use reexport readAdminCommercialConfig`.
- Depth 2 reaches consumer and route files through file-level `IMPORTS` projections, including both consumer files, but still has no symbol-level call edge to the wrapper.

Comparison: impact's file reach is consistent with the available barrel/import graph. It cannot recover the two missing `CALLS` relationships. The result is therefore a downstream projection of upstream graph loss, not an independent third root cause.

## `context` for the bounded email helper

Command: `anvien context symbol boundedReadLimit --file modules/email/server/operations/email-operations-observability.ts --repo E:\cheapapp.org --json`

Observed:

- Function range `190..192` and file fan-in `1` are correct.
- One incoming `CALLS` edge from `readEmailOperationsReport` is present, matching the independent source call at line 502.
- The two ambient calls on line 191 remain source-resolution gaps (`Math.max` and `Math.min`), inherited from the upstream graph.
- No process is returned for this bounded symbol; no completeness claim is made.

This is a bounded control: where an ordinary in-repo call edge exists, `context` exposes it; where the resolver emitted an ambient gap, `context` exposes the gap rather than manufacturing a target.

## Classification

| Surface | Result |
|---|---|
| Symbol identity/ranges | bounded valid |
| Wrapper consumer `CALLS` edges | upstream missing; command faithfully exposes absence |
| Ambient gap propagation | bounded valid as projection of the upstream wrong diagnostic |
| Impact file traversal | bounded valid for observed graph; cannot repair missing symbol edges |
| Process completeness | unresolved/not assessed; empty bounded output is not proof of no process |

## Evidence

- `E1-P1E-SRC1`: independent source reference counts and exact source lines.
- `E1-P1E-CMD1`: context output for `readAdminCommercialConfig`.
- `E1-P1E-CMD2`: upstream impact output for the wrapper UID.
- `E1-P1E-CMD3`: context output for `boundedReadLimit`.
- `E1-P1E-CAP1`: observed empty process arrays recorded as bounded command output, not completeness evidence.
- `E1-P1E-CMP1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1e-command-captures.json` reconciles commands with source counts.
- `E1-P1E-REPORT1`: this report.

## Remaining boundary

P1-E does not decide whether process/community/semantic derivation is globally complete. P2-C remains open for any downstream first-divergence case beyond the fixed command projections. No remediation is authorized.
