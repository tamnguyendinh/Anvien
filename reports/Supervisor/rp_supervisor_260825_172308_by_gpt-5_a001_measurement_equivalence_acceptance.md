# Supervisor Report: A001 Measurement And Equivalence Acceptance

Verdict: `SUPERVISOR_A001_PASS`

## Metadata

- Review time: `2026-08-25 17:23:08 +07:00` (Asia/Bangkok)
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`, Child 06A slice `A001`
- Claim reviewed: the two independent A001 target measurements and their equivalence evidence are valid and eligible inputs for Main's later `KEEP` decision.
- Authority: the official Child 06A A001 Supervisor assignment, root `AGENTS.md`, the current four Child 06A ledgers, the two target reports, and their raw artifacts.
- Role boundary: this verdict accepts the evidence only. Main alone owns `KEEP`, disposition, checklist transition, staging, and commit decisions.

## Executive Summary

The A001 evidence passes. Cheapapp and Restaurant Manager were measured as distinct repositories with one baseline/current pair per target. Within each pair, the target and workload flags are identical except for the expected benchmark output path and label; Restaurant Manager uses the exact exclusion `electron/renderer/src/api/userApi.ts` in both runs. All four process exits are `0`, executable hashes resolve to the recorded binaries, every raw benchmark and comparison has `30` top-level operations plus `17` resolution children, and the required D001, parent, and process-wall values match the raw JSON exactly.

Same-work evidence is closed: all 30 operation denominators and all 17 child denominators match within each target, resolution semantic metrics match, workload/output counts match, and canonical Graph JSON plus public stdout/stderr evidence is equivalent. Cheapapp's Ladybug file and `meta.json` are not byte-identical and were not represented as such; this is non-blocking because the canonical graph hash is exact, DB node/relationship readback counts match, public output matches, and A001 does not touch persistence or metadata code.

## Decisive Evidence

1. **Independent targets and identical pair configuration**
   - Cheapapp raw `process.json`: target `E:\cheapapp.org` and graph root `E:\cheapapp.org\.anvien` in both runs. Arguments differ only at the baseline/optimized benchmark artifact path and label. Target tracked status is identical after both captures.
   - Restaurant raw `pair_process.json`: target `E:\Restaurant_manager`, HEAD `605c0bda99491789e7f07628ec4b3d39a3ae1c67`, and worktree-diff hash `7EEEE1FBFD5DF429D00147AF4AEAAA6C0FBB63C396B86FFDFECDF266E81158C6` remain identical before and after both runs. Arguments differ only at benchmark artifact path and label; both contain the exact exclusion once.

2. **Exit and executable/build identity**
   - Cheapapp baseline/current exits: `0 / 0`; executable SHA-256: `5C0F4C6BC13204ABFCBBD05C53499F7C141C819875BB29D006EE130A2A04C53F / DA34A01158811833F3DE780E8308407FEC3B97CD903E446FD8BA2ED9D8164F51`.
   - Restaurant baseline/current exits: `0 / 0`; executable SHA-256: `F25C9C17E232834F0ED5FB8E7F3935BE61F5F457E181F105DB76EFA90CA20642 / DA34A01158811833F3DE780E8308407FEC3B97CD903E446FD8BA2ED9D8164F51`.
   - Direct file hashing returned the same values and `73,814,016` bytes for all four executables. Cheapapp `preflight.json` and Restaurant `build_identity.json`/`prepared_identity.json` record the build/overlay/Go identities.

3. **Cardinality and required timing values**
   - Both baseline/current benchmark JSON pairs contain `30 / 30` operations and `17 / 17` resolution children; Cheapapp `comparison.json` has `30` `topLevelRows` and `17` `childRows`; Restaurant `comparison.json` has `30` `operations` and `17` `children`.
   - Each report contains one product boundary, 30 unique `B1-P1A-OP001..OP030` rows, and 17 unique `B2-P2A-A001-D001..D017` rows.
   - Cheapapp raw values: D001 `209.823705100 -> 25.045225300 s`; parent `733.384225100 -> 184.481061700 s`; process wall `890.314783200 -> 279.105934600 s`.
   - Restaurant raw values: D001 `501.638742800 -> 40.769294200 s`; parent `1090.085959900 -> 136.436879300 s`; process wall `1178.391336900 -> 218.680628900 s`.

4. **Same-work and graph/output equivalence**
   - Independent raw comparison found zero operation-denominator mismatches and zero child-denominator mismatches for both targets. D001 denominators remain Cheapapp `calls=27890, files=887` and Restaurant `calls=86030, files=1234`.
   - Cheapapp: canonical graph `840614023` bytes and SHA-256 `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920` in both runs; workload/graph/DB row counts, resolution semantic metrics, stdout SHA-256, and stderr SHA-256 match.
   - Restaurant: canonical graph `900212685` bytes and SHA-256 `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C` in both runs; workload/output counts and stdout/stderr hashes match; `denominatorMismatchRows` is empty.

5. **Current candidate and recorded validation consistency**
   - Current candidate remains exactly five files and `+174/-10`, with scoped `git diff --check` clear and staged set empty. Its five SHA-256 values exactly match `E2-P2A-A001SRC1` and the measurement preflight identities.
   - Source inspection clears the exact approved owner boundary: `indexes.go` adds and populates the private original-index claim bucket; `resolve.go` replaces only `explicitImportNameClaimed`'s whole-import scan; `export_resolution.go` replaces only `explicitImportCallState` candidate acquisition while retaining semantic filters.
   - Current HEAD `72f7da1dc41d4d0b65e6a5477fa12640b51a7454` descends from recorded HEAD `e087f77032ff871b766c6af769eaa7c1aece9c73`. The intervening commit adds only one architecture report and changes none of the five candidate paths, so recorded build/test evidence remains attached to identical source bytes.
   - `proof_accuracy_golden_test.go` has no staged or unstaged change. The ledger's exact-HEAD overlay proof therefore remains consistent: the known golden failure is pre-existing/preserve-only, not a newly evidenced A001 regression.
   - Fresh required Anvien graph refresh exited `0` with `2214` files scanned, `765` parsed code files, and `0` failures. It was used only as current source/index freshness evidence, not as benchmark timing or a reopened gate.

6. **Ledger state**
   - The benchmark ledger holds separate Cheapapp and Restaurant blocks, each with 30 operation rows and 17 child rows; no averaging or combined result exists.
   - The plan ledger contains exactly `30/30` unchecked parent checklist items and `17/17` unchecked A001 child checklist items, with no checked item in either authoritative queue.
   - Actual-status and evidence ledgers retain the targets separately, record the selected-child no-`KEEP` streak as `0`, keep the candidate unpromoted, and state that no Supervisor result, `KEEP`, or disposition existed before this review. The attempt-history and terminal-disposition tables contain no fabricated row.

## Source-Level Clearance Notes

- `internal/resolution/indexes.go`: clear; private/run-scoped index only, original import indices and order preserved.
- `internal/resolution/resolve.go`: clear; exact approved helper body only.
- `internal/resolution/export_resolution.go`: clear; acquisition narrowed to the claim bucket while semantic filtering remains.
- Authorized tests: hashes and diff scope match recorded evidence; no test or build rerun was needed to establish byte consistency.

## Not Run

- Neither benchmark capture was rerun, as explicitly prohibited.
- Full build and tests were not rerun; acceptance used the already-recorded build/test evidence only after proving the current five candidate files are byte-identical to that evidence.
- Architect, Planner, Coder, impact, disposition, staging, and commit gates were not reopened.

## Invariant Closure

- Affected invariant: same-target/same-work performance comparison with unchanged resolution semantics and canonical graph/public output.
- Checked sibling surfaces: both target pairs, process/executable identities, all operation and child denominators, canonical graph/output counts/hashes, current production diff, validation chronology, and all four current ledgers.
- Residual unverified same-invariant surfaces: none required for this acceptance boundary.

## Overall Evaluation

The measurement and equivalence evidence is valid and eligible for Main's `KEEP` decision. This report makes no `KEEP`, promotion, rollback, disposition, staging, or commit decision.
