# P1-C Resolution and Module-Binding Comparison

Date: 2026-07-26 15:27 +07
Target: `E:\cheapapp.org`
Target HEAD: `a869876ab6262dacde6cd5d432d099a91852a646`
Graph: `E:\cheapapp.org\.anvien\graph.json`
Graph SHA-256: `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`
Classification: bounded `confirmed wrong`; first-divergence ownership remains P2-B

## Question

For five fixed source cases, does the fresh Anvien graph agree with an independent TypeScript binding oracle about ambient standard-library symbols and a local barrel re-export?

This report is a fresh source-to-graph comparison. It does not accept an earlier report or subagent conclusion, does not claim global resolver accuracy, and does not propose or implement a fix.

## Boundary

- Target source was read in place at `E:\cheapapp.org`; it was not copied.
- The target graph/index remains at `E:\cheapapp.org\.anvien`.
- Oracle scripts, compact captures, and this report are under `E:\Anvien`.
- No target source, test, dependency, or configuration file was edited.
- TypeScript is not installed in the target worktree. The independent oracle used TypeScript `5.9.3` already present under `E:\Anvien\node_modules`; the target declares TypeScript `6.0.3`. This version difference limits any global compiler claim, but all five exact sites are ordinary long-standing ambient/alias semantics and produced concrete declarations with no diagnostic.

## Case A — ambient standard-library bindings

Source sites:

| Source site | Source fact | Fresh Anvien observation | TypeScript oracle |
|---|---|---|---|
| `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts:13` | Return type is `Promise<AdminCommercialConfigReadModel>` | `Promise` is `unresolved_type_reference`, `in_repo_unresolved`, `analyzer_gap`, `unresolved_local_binding` | Symbol `Promise`; declarations in `lib.es5.d.ts`, `lib.es2015.iterable.d.ts`, `lib.es2015.promise.d.ts`, `lib.es2015.symbol.wellknown.d.ts`, and `lib.es2018.promise.d.ts`; no diagnostic on line 13 |
| `modules/email/server/operations/email-operations-observability.ts:191` | `Math.max(1, Math.min(...))` | `Math.max` and `Math.min` are two `unresolved_call` rows, both `in_repo_unresolved`, `analyzer_gap`, `unresolved_local_binding` | `max` and `min` resolve to `MethodSignature` declarations at `lib.es5.d.ts:733` and `lib.es5.d.ts:738`; no diagnostic on line 191 |

Result: Anvien's unresolved/classification surface is wrong for these three bounded standard-library references. Whether the intended graph contract should materialize external library declaration nodes or represent an explicit external/ambient terminal is a P2-B/product-contract question; labeling the sites `in_repo_unresolved` and actionable analyzer gaps is not faithful to compiler binding truth.

## Case B — local barrel re-export binding

Source chain:

1. `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts:10` declares and exports `readAdminCommercialConfig`.
2. `modules/commercial-config/server/admin-commercial-config/index.ts:21` re-exports it from `./read-admin-commercial-config`.
3. `modules/admin-operations/server/commercial-config/save-admin-commercial-config-mutation.ts:4` imports it from the barrel and calls it at line 142.
4. `modules/admin-operations/server/commercial-config/read-admin-commercial-config-route-view.ts:6` imports it from the barrel and calls it at line 32.

Independent TypeScript results:

| Consumer call | Local declaration | Aliased declaration | Diagnostics on call line |
|---|---|---|---|
| `save-admin-commercial-config-mutation.ts:142:30` | import specifier at line 4 | function declaration `read-admin-commercial-config.ts:10:1` | none |
| `read-admin-commercial-config-route-view.ts:32:29` | import specifier at line 6 | function declaration `read-admin-commercial-config.ts:10:1` | none |

Fresh Anvien results:

- The barrel file has `symbolCount: 0`, even though its source contains explicit value/type re-exports.
- Each consumer has an `IMPORTS` relationship to the barrel File node.
- Neither consumer has the expected `USES`/`CALLS` relationship to the `readAdminCommercialConfig` declaration.
- `save-admin-commercial-config-mutation.ts:142` instead contains an `unresolved_type_reference` plus an `unresolved_call` for `readAdminCommercialConfig`; the call is only `global-fallback-low-confidence`.
- `read-admin-commercial-config-route-view.ts:32` contains the same pair of false gaps.

Result: the two fixed barrel-mediated calls are compiler-bound to the unique local function declaration but are exposed by Anvien as unresolved local bindings. This is a bounded module/re-export binding error, not evidence that every barrel export fails.

## Exact counts

- Fixed source sites checked: `5` (`Promise`, `Math.max`, `Math.min`, and two barrel-mediated calls).
- Compiler-bound sites: `5/5`.
- Diagnostics at those sites: `0`.
- Sites Anvien exposes as unresolved: `5/5`.
- Ambient false gaps: `3` graph rows.
- Barrel-call false gaps: `4` graph rows across `2` call sites (`unresolved_type_reference` + `unresolved_call` per site).
- Expected barrel-call `CALLS` edges to the declaration: `2`; observed: `0`.

## Evidence

- `E1-P1C-GRAPH1`: fresh graph identity and up-to-date status at the recorded target HEAD/hash.
- `E1-P1C-SRC1`: exact target source for `Promise`, `Math.max`, and `Math.min`.
- `E1-P1C-ANVIEN1`: fresh `file-detail` rows for the three ambient sites.
- `E1-P1C-ORACLE1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1c-typescript-oracle.mjs` and `p1c-typescript-oracle-output.json`.
- `E1-P1C-SRC2`: exact declaration → barrel re-export → two consumer import/call chains.
- `E1-P1C-ANVIEN2`: fresh barrel/consumer `file-detail` summaries and unresolved rows captured in `p1c-anvien-observations.json`.
- `E1-P1C-ORACLE2`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1c-barrel-oracle.mjs` and `p1c-barrel-oracle-output.json`.
- `E1-P1C-REPORT1`: this report.

## Remaining boundary

P1-C proves the mismatches but does not identify their earliest Anvien owner. P2-B must separately trace:

- ambient/builtin classification and resolution policy;
- extraction and resolution of named barrel re-exports;
- the point where a bound imported alias becomes `unresolved_local_binding` and a low-confidence global fallback.

No remediation is authorized in this investigation plan.
