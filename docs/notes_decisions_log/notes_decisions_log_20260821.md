# Notes / Decisions Log - 2026-08-21

## Child 04 P4-B1 — comment-bearing TypeScript/JavaScript re-export recovery

- Scope: P4-B1 source-bearing named/default/star/namespace/type-only re-export syntax only. P4-C graph/persistence projection, P4-C2 target validation, Child 05 terminal/barrel/cycle/ambiguity/public API, and `E:\cheapapp.org` remain locked.
- REVIEW2 blocker: legal comment trivia around recovered malformed `as`/comma syntax dropped valid `AlsoGood` siblings. REVIEW3 production recovery now preserves `Good` and `AlsoGood` at exactly `2` facts / `1` dangling-as diagnostic / `2` derived imports for both comment placements, with no `Broken` fact/import and zero terminal state.
- Evidence: Coder report `reports/coder/rp_coder_260821_025056_by_gpt-5_child04_p4b1_comment_recovery_review3.md`; Supervisor PASS `reports/Supervisor/rp_supervisor_260821_031004_by_gpt-5_child04_p4b1_comment_recovery_review3.md` (`07DD5BB92F169C5923C0DBCB597F914A28E594496ACDADB55D41F13DE364421C`).
- Validation: fresh canonical `npm run full-build` PASS; focused `6/6`; full `tsjs` PASS; nearest `3/3`; `resolution/analyze 2/2`; fresh excluded graph `1,144/626/0`, `82,059/121,760`; formatting, forbidden-field, visibility, and `.tmp` cleanup gates PASS.
- Main-owned detect: final `anvien detect-changes --repo E:\Anvien --scope all` exits `0` with `305` changed semantic units, `7` changed/affected files, `1` affected process, MEDIUM risk, docs layer `25`, and zero persisted resolution-health degradation; final graph is `1,146/626/0`, `82,079/121,780`. Next gate is staging the accepted exact boundary and creating one isolated P4-B1 commit with no push. Preserve P4-C/P4-C2/Child 05 locks.
