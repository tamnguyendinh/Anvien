# Notes / Decisions Log - 2026-08-20

## Child 04 P4-B — TS/JS first-class export facts

- Scope: direct/default/local-alias/type-only TypeScript and JavaScript export syntax only; source-bearing re-export compatibility, graph/projection/persistence, P4-B1/P4-C/P4-C2, Child 05, and target remain locked or preserve-only.
- Evidence: Coder report [`rp_coder_260820_235311_by_gpt-5_child04_p4b_export_facts.md`](../../reports/coder/rp_coder_260820_235311_by_gpt-5_child04_p4b_export_facts.md) records `E4-P4B-IMPACT1`, `E4-P4B-SRC1`, `E4-P4B-BUILD1`, `E4-P4B-TEST1`, and `E4-P4B-BOUNDARY1`.
- Result: provider-focused and nearest boundary tests pass; canonical `npm run full-build` exits `0`; `DefinitionFact.Visibility` remains unchanged; no stage/commit/push/final detect was performed.
- REVIEW1 closure: removed the exact empty repo-local debug directory `.tmp/p4b_ast_probe/`; fresh filesystem/status/diff checks show it absent and candidate production/test paths unchanged. Nearest boundary re-run passes all three packages.
- Handoff: `READY_FOR_SUPERVISOR` resubmission; Main owns independent review and subsequent final gates.
