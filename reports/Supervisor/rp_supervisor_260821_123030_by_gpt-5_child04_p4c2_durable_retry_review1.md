# Supervisor Report: Child 04 P4-C2 durable QA retry REVIEW1

Verdict: REJECT

## Metadata

- Report ID: `E4-P4C2-REVIEW1-260821-durable-retry`.
- Report file: `E:\Anvien\reports\Supervisor\rp_supervisor_260821_123030_by_gpt-5_child04_p4c2_durable_retry_review1.md`.
- Review time: `2026-08-21 12:30:30 +07:00` (`Asia/Bangkok`).
- Reviewer: `gpt-5` / independent P4-C2 Supervisor.
- Repo/project: `E:\Anvien` / Anvien.
- Scope reviewed: Child 04 P4-C2 only, against accepted P4-C implementation commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877`, sealed oracle `p4c2-oracle-v1-a869876ab626-260821_110849+0700`, and retry run `p4c2-target-validation-retry-260821_115050+0700`.
- Claim reviewed: the retry handoff reports complete comparison with 21/21 Export facts, but 15 exported TypeAlias rows (`P001`-`P014`, `P018`) have compatibility `isExported=false` and FileContext `exported=false`; six Function positives and 11/11 negatives pass; Graph JSON/Ladybug persistence parity is clean.
- Authority used: latest Owner/Main delegation; `E:\Anvien\AGENTS.md`; mandatory `working-rules`, `supervisor`, and `data-integrity-review` skills; roadmap and all four Child 04 ledgers; `docs/contracts/graph-accuracy-contract.md`; accepted P4-C source/diff/report; sealed oracle; retry run artifacts; exact target source anchors.
- Review boundary: review-only. No production, test, golden, plan, ledger, oracle, QA artifact, target, config, stage, commit, push, reset, checkout, analyze, query, comparison rerun, Child 05, or repair action occurred. This file is the only artifact created by this review.

## Executive Summary

- Problem: decide whether P4-C2 satisfies the bounded 21-positive/11-negative direct-export contract, while distinguishing persistence parity from semantic correctness.
- Decision: `REJECT`. The sealed oracle requires `compatibility.isExported=true` for every positive direct export. Fresh target evidence contains all 21 Export facts and relations, but the exact 15 TypeAlias definitions are projected as `isExported=false`; production FileContext consumes that canonical false and reports those same 15 symbols as non-exported. Only the six Function positives pass, so the positive contract is `6/21`, not `21/21`.
- Required outcome: Main owns the rejection handoff. P4-C2 remains open; Child 05 and later slices remain locked. A repair lane must close only the direct-export compatibility/reader invariant and return fresh bounded evidence for re-review.

## Blocking Finding

### [HIGH] Direct exported TypeAlias definitions are suppressed by the runtime-only compatibility projection

- Broken invariant: one accepted direct export fact must derive truthful compatibility for its local definition and every affected reader, independently from access visibility. The sealed positive schema requires `isExported=true` for all 21 direct exports; `typeOnly` and meaning remain separate Export-fact semantics and must not turn a source-exported definition into an unexported definition.
- Expected rule: `P001`-`P021` each have one direct Export fact, one File-to-Export relation, absent access state, compatibility `isExported=true`, and FileContext `exported=true`. Type aliases retain `meanings=[type]` and `typeOnly=true`; they are not runtime values, but they are still directly exported source definitions.
- Actual behavior: `P001`-`P014` and `P018` each have the correct Export fact, relation, type meaning, type-only state, source identity, range, provenance, zero diagnostics, and empty Child 05 state, but compatibility and FileContext are both false. Their only failed fields are exactly `compatibility.isExported` and `reader.exported`.
- Source cause: `internal/resolution/emit.go:374-375` adds a local definition to `directExportDefIDs` only when `exportFactIsRuntime` returns true. `internal/resolution/emit.go:455-466` returns false for `TypeOnly` facts and facts without value/namespace meaning. `internal/resolution/emit.go:280-281` therefore writes `isExported=false` on all 15 direct exported TypeAlias definitions. `internal/filecontext/context.go:1334-1340` treats the canonical boolean as authoritative, so the affected reader faithfully propagates the wrong compatibility value. `internal/resolution/p4c_export_projection_test.go:146-147` currently codifies the rejected `type-only definition isExported=false` expectation.
- Affected flow: source `export type` declaration -> `ScopeIR.ExportFact` -> Graph Export node/File containment -> local definition compatibility `isExported` -> Ladybug definition row -> FileContext exported-symbol result.
- Impact: 15/21 bounded positives fail; the exported-symbol counts are `3/17`, `2/3`, and `1/1` for the three target files. Six Function rows (`P015`-`P017`, `P019`-`P021`) pass. All 11 negative controls remain false and do not false-positive.
- Why this blocks acceptance: the roadmap, plan, contract, and sealed oracle require `21/21` correct direct exports plus correct affected readers. Record presence and persistence equality cannot substitute for semantic truth.
- Fix direction: preserve the first-class Export fact's `typeOnly`/meaning semantics, but derive definition-level direct-export compatibility from direct source export membership rather than runtime-value eligibility. Keep access visibility independent and preserve all negative controls and Child 05 exclusions. Update tests only after production behavior is corrected.
- Re-review evidence required: fresh bounded `21/21` positives and `11/11` negatives; exact TypeAlias compatibility and FileContext results; Graph JSON/Ladybug record-level parity; zero duplicate/orphan/diagnostic/forbidden state; target/config boundary preservation; required build/analyze/change-detection evidence under Main's authorized workflow.

## Exact Rejected and Accepted Invariants

Rejected:

- `E4-P4C2-TARGET1` semantic contract: positive rows are `6/21`; `P001`-`P014` and `P018` fail compatibility and reader output.
- Direct-definition compatibility: 15 exported TypeAlias definitions are `isExported=false`, expected `true`.
- Affected FileContext reader: the same 15 definitions are `exported=false`, expected `true`.
- P4-C2 overall acceptance: the required `21/21` direct-export result is not met.

Accepted as evidence or preserved boundaries:

- `E4-P4C2-ORACLE1`: SEALED. Five non-seal files match declared bytes/SHA-256 and independently recompute bundle digest `7749AB14E02FBF61CF7F81B3A7638888F8AD414064F9F969CFED06EB11355439`; counts are exactly 21 positive and 11 negative.
- Target source basis: the three live read-only source hashes match the sealed identities. Direct source inspection confirms 15 `export type` declarations and six `export function` declarations at the sealed anchors.
- Export-fact conservation: all 21 positives have one Export node and one File-to-Export relation in Graph JSON and Ladybug; exact kind/name/meaning/typeOnly/source/range/access/provenance checks pass.
- Function positives: `P015`-`P017`, `P019`-`P021` pass completely (`6/6`).
- Negative controls: `N001`-`N011` pass completely (`11/11`), including owner-qualified same-name `time` and `now` occurrences and array-binding leaves.
- Integrity boundary: duplicate node IDs `0`; missing relationship endpoints `0`; orphan local-definition references `0`; export diagnostics `0`; forbidden Child 05 terminal/resolved-target/public-API state `0`.
- `E4-P4C2-BOUNDARY1`: target HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, branch `master`, tracked-status SHA-256 `FCB5AD9155C029FFA6B3D80B8AADD1EE40EDC3B408EA401992E8FEA048E4C5E1`, three source identities, four Git-config identities, and four analyzer-artifact identities remain equal across sealed/pre/post/final evidence; config differences are `0`, process-local trust was cleared, and no target source/config/report/fixture/probe/debug write is evidenced.

## Persistence Parity Is Not Semantic Correctness

- Graph JSON and Ladybug contain the same 21 bounded Export rows.
- Export comparison is `28 x 21 = 588` fields with `0` differences and `0` IDs missing in either direction.
- Definition compatibility comparison is `32` fields with `0` persistence differences; File-to-Export comparison is 21 relations with `0` differences.
- This parity is lossless persistence of the current value, including the wrong `false` compatibility on 15 TypeAlias definitions. It proves storage consistency, not that the stored value satisfies the source oracle.

## Source-Level Clearance Notes

- `internal/resolution/emit.go`: blocked for the bounded compatibility invariant. The first-class Export nodes are conserved, but runtime eligibility is incorrectly reused as definition export membership.
- `internal/filecontext/context.go`: reader behavior is internally deterministic and fail-closed, but it exposes the upstream false value; the bounded reader result is therefore rejected.
- `internal/lbugload/csv.go` and `internal/lbugschema/schema.go`: clear for persistence parity. They preserve the definition compatibility and 28 Export fields without loss; they do not repair semantic values.
- P4-C focused tests: blocked where `p4c_export_projection_test.go` expects a type-only direct definition to be unexported. Other conservation, persistence, reader precedence, orphan, and forbidden-state tests remain useful preserve-only coverage.
- Accepted implementation commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877` remains an ancestor and its four production owners have no drift through current HEAD.

## Evidence Checked

Passed:

- Oracle file identity and non-circular digest recomputation: all five non-seal files exact; seal status `SEALED`; 21/11 unique rows; positive schema fixes `compatibility.isExported=true`, negative schema fixes it false.
- Retry manifest identity: all 19 listed files match exact bytes, LF counts, and SHA-256; actual file set and ordinal order match; run digest independently recomputes to `9F414A2C54C42F4E39AD8ED03DC042CCC3E1FB5993DA842B22F64851D16AABC4`.
- QA report identity: `11,342` bytes / `200` LF / SHA-256 `C831004F049A563A2387B599BE01C943F5B9416C72B1C45E50A8C1F9D2CEFDB4`.
- Comparison identity: `245,764` bytes / `4,899` LF / SHA-256 `2C78AB3BDF67D857E5C2A1B75B0F1FDFBFEBE2B70D92E7EBC8EB45A0AC5A3F27`; structured status `FAIL_COMPLETE`, `comparisonComplete=true`, `contractPass=false`.
- Independent row aggregation over `comparison-result.json`: exact failed set `P001`-`P014`,`P018`; exact two failed fields per row; every other positive field passes; six Function positives pass; all 11 negatives pass.
- Validator source inspection: graph values, Ladybug read-only values, relation endpoints, and production FileContext outputs are compared directly; no analyzer/query command is invoked by the validator.
- Target source read-only check: byte/SHA identities match the seal and source text confirms the bounded direct exports.
- Target/config preservation cross-check: sealed basis, retry preflight, retry post-state, and final post-state have identical HEAD/branch/status/source identities; Git config difference count `0`; analyzer artifacts are unchanged after comparison.
- Current Anvien boundary before this report: five unstaged Child 04 living documents, canonical tracked-status SHA-256 `70C58FE493E8B1B9ABEFA91473F1F35F1DB6A33F31E75E46F4B1CD75AD1B5E6C`, empty index, and `git diff --check` exit `0`.

Failed:

- 15 TypeAlias compatibility checks: expected true, actual false.
- 15 corresponding FileContext exported checks: expected true, actual false.
- Positive acceptance count: `6/21`; P4-C2 contract is false.

Not run:

- No oracle authoring, full build, target preflight, target analyze, Anvien graph refresh/query, comparison gate, Ladybug query, detect-changes, test, commit, or cleanup rerun. Closed gates were not repeated because their identities were not invalidated.
- No Child 05, terminal/barrel/alias-chain/cycle/ambiguity/public-API work.

## Repo and Concurrent Boundary

- Handoff HEAD was `e32a412b289453a530bc71b93320ef2b97b3a97a`.
- During review, current HEAD advanced to `6b93e80601b0549f1b3e56bd6b68b07b9bc9680a`, direct parent `e32a412b289453a530bc71b93320ef2b97b3a97a`, commit `docs(orchestration): clarify main coordinator boundaries` at `2026-08-21 12:28:40 +07:00`.
- Exact delta is only `internal/aicontext/skills/orchestration/SKILL.md` (`2` insertions). It touches no P4-C/P4-C2 production, test, golden, oracle, QA evidence, target, or living ledger; Main independently confirmed this concurrent provenance-only boundary. It does not invalidate the review target.
- At report creation, the five Child 04 living documents remain the only tracked unstaged paths and the index is empty. This Supervisor report is intentionally untracked and is the only review-created artifact.
- Final verification at `2026-08-21 12:33:12 +07:00` found a later concurrent unstaged edit on that same out-of-scope orchestration skill (`2` insertions / `9` deletions relative to HEAD). The final tracked worktree set is the five living documents plus this orchestration skill (`6` paths), canonical tracked-status SHA-256 `CE7A69FE8599746512193E5ED2C47CA63B0685DDC96B9310B1CE6AFF2898364D`; the index remains empty and P4-C production drift remains empty. This concurrent edit was not created or modified by Supervisor and does not invalidate the bounded verdict.

## Evidence-ID Disposition and Handoff

- `E4-P4C2-ORACLE1`: accepted as sealed source authority, digest `7749AB14E02FBF61CF7F81B3A7638888F8AD414064F9F969CFED06EB11355439`.
- `E4-P4C2-TARGET1`: comparison complete, but semantic result rejected (`FAIL_COMPLETE`, positives `6/21`, negatives `11/11`).
- `E4-P4C2-BOUNDARY1`: accepted; target/config/source/analyzer boundary preservation is proven by the durable bundle.
- `E4-P4C2-REVIEW1`: `REJECT` by this report.
- Next owner: `Main`. Main must keep P4-C2 open, keep Child 05 locked, and route only the rejected compatibility/reader invariant to the authorized repair workflow. Supervisor does not close the slice or open Coder/Child 05.

## Residual Uncertainty

- No residual uncertainty remains about the bounded defect or affected row set: oracle authority, source syntax, production cause, persisted values, reader values, and exact failures agree.
- Supervisor did not rerun the large closed analyzer/Ladybug gates or live target Git/config preflight. Their preservation is established by identity-checked pre/post/final artifacts; the three semantic source files were independently hash-checked live. This deliberate non-rerun does not weaken the REJECT because the existing immutable evidence already proves the failure.
- Out-of-scope scanner omissions, general unresolved sites, and Child 05 resolution remain unreviewed non-claims.

## Overall Evaluation

The retry handoff is evidence-complete and internally trustworthy, but it proves a live semantic defect rather than acceptance. Export facts and persistence are correct for the bounded records; compatibility projection conflates runtime eligibility with direct source-export membership, and FileContext propagates that false value. Therefore P4-C2 cannot pass until all 21 positives, including the 15 TypeAlias definitions, satisfy the sealed compatibility and reader contract without regressing negatives, persistence, access separation, or Child 05 boundaries.

## Report Identity

- Encoding/line ending: UTF-8 / LF.
- Canonical SHA-256 basis: this report with the 64-character value after `canonical SHA-256:` replaced by 64 ASCII zeroes (self-reference-safe).
- Bytes / LF / canonical SHA-256: `15671` bytes / `126` LF / canonical SHA-256: `8E37F4B126ABB38A5DCE071C35937F59D9152EFA135B9245408CA138BEC781F7`.
