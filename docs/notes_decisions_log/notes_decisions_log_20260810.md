# Notes / Decisions Log - 2026-08-10

## Child 01 P1-B - Range, Selection Range, and Identity Inputs

- Scope: implemented only `P1-B` from `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction`.
- Decision: retain TSJS tree-sitter coordinates as one-based lines, zero-based UTF-8 byte columns, and exclusive endpoints; store the existing construct range plus an optional declaring-token selection range. Existing `ScopeFact` lexical membership, `Label`, and `OwnerID` remain the owner/meaning inputs; no new identity topology or later-slice behavior was introduced.
- Preservation: `Range`, `PositionIndex`, TSJS scope containment, other providers, Child 03 binding patterns, P1-C identity, P1-D collision handling, projection, persistence/readers, and `E:\cheapapp.org` were not changed.
- Validation: full build PASS; compiled TSJS provider/ScopeIR boundary PASS; focused matrix and three owner packages PASS; product regression PASS `61/61` packages. Raw `go test ./...` retains the known compile-negative/non-standalone fixture failures.
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-evidence.md` (`E1-P1B-IMPACT1`, `E1-P1B-SOURCE1`, `E1-P1B-BUILD1`, `E1-P1B-TEST1`, `E1-P1B-RUNTIME1`, `E1-P1B-REVIEW1`, `E1-P1B-DETECT1`, `E1-P1B-COMMIT1`).
- Coder report: `reports/coder/rp_coder_260810_113052_by_gpt-5-codex_child01_p1b.md`.
- Supervisor reports: `reports/Supervisor/rp_supervisor_260810_115437_by_gpt-5-codex_child01_p1b.md` (`REJECT`, stale ledger narrative only), followed by `reports/Supervisor/rp_supervisor_260810_120042_by_gpt-5-codex_child01_p1b_resubmission.md` (`PASS`).
- Change detection: fresh final analyze PASS (`1,564/677/0`, `85,247` nodes, `124,149` relationships); detect-changes PASS at overall risk `low` with `50` changed symbols in `7` files, `0` affected symbols, no affected process, and zero resolution-health degradation.
- Implementation commit: accepted P1-B work is committed immediately after this note as one isolated boundary with message `fix(scopeir): preserve TSJS definition selection ranges`; final hash is reported from Git because a commit cannot contain its own hash.
