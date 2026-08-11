# Notes / Decisions Log - 2026-08-11

## Child 01 P1-C - Identity and Occurrence Conservation

- Scope: implemented only `P1-C` from `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction`.
- Decision: consume the current source-backed provider occurrence ID, lexical `scopeByDef` membership, label/meaning, owner ID, semantic name, optional arity, and normalized repository-relative file at the existing `graphIDForDef`/`buildWorkspace` owner. Tuple fields are length-prefixed; no hash, new topology/package, random/absolute/order input, or provider rewrite was introduced.
- Preservation: providers, TSJS binding patterns, projection source, `Graph.AddNode`, relationship merge, persistence/readers, later Children, and `E:\cheapapp.org` were not changed.
- Validation: focused identity matrix PASS `4/4`; resolution and provider packages PASS; full build PASS; built CLI fixture oracle PASS `10/10` occurrence conservation with `time 2/2`, `now 2/2`, separate `Shared` meaning lanes, and `0` missing endpoints; product regression PASS `61/61` packages after stale literal-ID QA alignment.
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md` (`E1-P1C-IMPACT1`, `E1-P1C-SOURCE1`, `E1-P1C-BUILD1`, `E1-P1C-TEST1`, `E1-P1C-ORACLE1`).
- Coder report: `reports/coder/rp_coder_260811_083419_by_gpt-5-codex_child01_p1c_identity_occurrence.md`.
- Review state: first zero-trust review cleared the technical boundary but REJECTed stale actual-status prose; history-closure re-review PASSed with unchanged source/test hashes and `50/50` consistency assertions. Nine exact current-slice `.tmp` artifacts were removed while historical work was preserved.
- Detect state: fresh staged all-scope detect-changes PASSed at risk `medium` with `223` changed symbols / `19` accepted paths / one affected process, and zero persisted ResolutionGap, degraded node, or node-with-gap result. The isolated P1-C commit follows this ledger state.
- Later gates: P1-D collision handling, P1-E target validation, Child 02 persistence/readers, and target source remain closed.
