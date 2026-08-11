# Supervisor Report: Child 01 P1-C Identity and Occurrence Conservation

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260811_090631_by_gpt-5-codex_child01_p1c.md`
- Review time: `2026-08-11 09:06:31 +07:00` (Asia/Bangkok)
- Reviewer: `gpt-5-codex`
- Repo/project: `E:\Anvien`
- Scope reviewed: Child 01 `P1-C` candidate only, at base HEAD `6b3a3fc4480c7f4c2291d8004207f50b52d91d21` plus the current worktree. P1-D collision policy, P1-E target acceptance, `E:\cheapapp.org`, persistence/readers, relationship merge policy, and later Children are out of scope except preserve-boundary checks.
- Claim reviewed: the candidate implements deterministic source-occurrence graph identity and 100% occurrence conservation at the owned `defsByFile -> graph identity -> definition node/endpoint` boundary, has completed build/runtime/QA/regression evidence, and is ready for the P1-C Supervisor gate without opening later slices.
- Authority used: latest user work rules, `AGENTS.md`, the complete Supervisor skill, active roadmap, `docs/contracts/graph-accuracy-contract.md`, all four Child 01 ledgers, current source/diff/runtime, prior same-scope investigations and Child 01 review history.
- Related artifacts: `reports/coder/rp_coder_260811_083419_by_gpt-5-codex_child01_p1c_identity_occurrence.md`, `.tmp/p1c-runtime-260811`, Child 01 P0/P1-A/P1-B Supervisor reports, and the bounded identity root-cause reports.

## Executive Summary

- Problem: determine whether the P1-C implementation and its live authority close only declaration occurrence identity/conservation, without promoting the old DRAFT architecture or moving into collision policy, target validation, persistence, or relationship redesign.
- Decision: the production implementation and current technical evidence clear the P1-C code boundary, but the live actual-status authority still describes mutually exclusive P1-C states. Its current matrix/R10-R14/next-decision row say full build, built runtime, and the 61-package regression are complete, while its Detailed Findings say built-runtime acceptance is pending, direct the workflow to rerun already-completed build/runtime/regression steps, and allow P1-D classification before the P1-C review/commit gate.
- Required outcome: update only the stale P1-C Detailed Findings in the actual-status ledger, preserve the unchanged production/QA candidate and all closed later-slice boundaries, then resubmit for history-closure review.

## Blocking Findings

### [HIGH] The live actual-status authority gives mutually exclusive P1-C state and next-action instructions

File: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-actual-status.md:141`

Issue: the current-state sections do not describe one reviewable P1-C state. R11 records the built runtime oracle as completed at line 110, R14 records the exact `61/61` product regression at line 113, and the P1-C next-decision row at line 241 says full build, built-runtime fixture, regression, and 100% owned-boundary conservation are complete with only Supervisor/detect/commit pending. In conflict, line 141 says built-runtime acceptance remains pending; lines 153 and 155 say build/runtime/review still must pass and instruct the workflow to run build/runtime/regression; line 196 again says built-runtime acceptance is pending; line 198 advances to P1-D classification; and line 231 says normal runtime may be used only after P1-D, even though P1-C itself requires and already records normal built-runtime validation.

Evidence: the contradiction is direct in the live ledger at lines `110`, `113`, `141`, `153`, `155`, `196`, `198`, `231`, and `241`. Current source and independent runtime evidence confirm the later state: `internal/resolution/indexes.go:143,814-834` consumes lexical membership and encodes the length-prefixed occurrence tuple; the freshly compiled current-source binary reproduced `20` nodes / `19` relationships, `10/10` unique definitions, `10/10` `DEFINES`, `time` at lines `5/10`, `now` at `15/20`, distinct Function/Interface `Shared`, and zero missing endpoint references. Fresh current-worktree QA also completed `4/4` focused identity cases and `61/61` product packages. The plan requires live ledger state to change when evidence state changes and keeps P1-D closed until P1-C review, detect-changes, and isolated commit.

Why this blocks acceptance: actual-status is active execution authority, not commentary. The contradictory instructions can send the workflow backward into completed P1-C validation or forward into P1-D before the mandatory acceptance/commit boundary. A completion gate cannot be accepted while its authority states both “runtime complete” and “runtime pending,” and while a later slice is presented as an allowed next action.

Fix direction: change only the stale P1-C Detailed Findings. Record that build, built-runtime fixture validation, endpoint checks, and product regression are complete; keep P1-C overall incomplete only because Supervisor, detect-changes, artifact cleanup, and isolated commit remain gated; make zero-trust resubmission the sole next action; keep P1-D/P1-E/target/persistence/readers/relationship policy closed. Do not change production code, QA, contract topology, or later-slice ownership.

Re-review evidence required: an exact ledger-only diff; current hashes for `internal/resolution/indexes.go`, `internal/resolution/identity_occurrence_test.go`, and `internal/providers/provider_parity_test.go` unchanged; `git diff --check`; and consistency assertions covering the header, Current Status Matrix, R10-R14, Detailed Findings, Next Phase Status Decisions, unchecked P1-C checkbox, pending review/detect/commit rows, and closed P1-D/P1-E rows.

## Source-Level Clearance Notes

- `internal/resolution/indexes.go`: clear for the technical P1-C boundary. `buildWorkspace` at line 143 passes the existing `scopeByDef[def.ID]`; `graphIDForDef` at lines 814-830 uses normalized repository-relative file, semantic/qualified name, optional arity, provider occurrence ID, lexical scope ID, owner ID, and the provider label as meaning prefix; `graphIdentityPart` at lines 832-834 makes tuple boundaries unambiguous. Nil arity and zero arity remain distinct.
- Provider/ScopeIR input group: clear for the current owner assumption. All 13 production provider `defID` implementations use repository file plus declaration start line/column, label, and name; `ScopeFact.OwnedDefIDs` supplies lexical membership; normal analyze passes scanner-relative `file.Path` to providers. No provider production file changed in P1-C.
- `internal/resolution/emit.go` and `internal/graph/types.go`: clear as preserved siblings. `emitDefinitionNodes` still emits one attempted node and `DEFINES` relationship per `defsByFile` ref; relationship merge remains separate; `Graph.AddNode` collision policy is unchanged and remains P1-D-owned.
- P1-C QA group: clear. The new identity test directly checks occurrence denominator, same-name scopes, provider meaning lanes, lexical/owner/arity sensitivity, nil-versus-zero arity, path separator normalization, input reordering, unsupported merge, and all relationship endpoints. Legacy relationship tests now resolve semantic nodes rather than accepting old/new ID alternatives.
- Child 01 roadmap/plan/evidence/benchmark ledgers and coder/decision reports: clear for scope and pending review/detect/commit gates. The actual-status ledger alone is blocked by the current-state contradiction above.

## Evidence Checked

Passed:

- Full authority/history read: original problem report was treated only as bounded finding/oracle input; its two-tier/hash/migration proposal remained DRAFT and was not imported into the implementation verdict.
- Current source/diff: only `internal/resolution/indexes.go` is a touched production file. No diff exists in `internal/graph`, `internal/resolution/emit.go`, `internal/scopeir`, or `internal/providers/tsjs`; no `internal/graphidentity` package or `internal/graph/declaration_occurrences.go` was added.
- Fresh self-graph refresh before all graph queries: `1,568` scanned, `679` parsed-code, `0` failed, `95,556` nodes, and `134,423` relationships; `anvien status` reports indexed/current commit `6b3a3fc` and up-to-date.
- Fresh `file-detail` for `internal/resolution/indexes.go`: parsed backend/resolution source, `319` symbols, `47` files in the full file dictionary (`46` related files excluding the target), `28` linked flows, `24` linked tests, file risk high, `stale=false`, `changedSinceAnalyze=false`.
- Fresh current-worktree blast radius with tests: `graphIDForDef` is CRITICAL with `27` impacted symbols across `14` files, `3` modules, and `19` processes; new `graphIdentityPart` is CRITICAL with `18` symbols / `10` files / `1` module / `5` processes; `buildWorkspace` is CRITICAL with `44` symbols / `18` files / `8` modules / `25` processes. Affected app layers include backend, backend tests, and for `buildWorkspace` the CLI launcher. These are preserved as scope warnings, not edit bans.
- Fresh focused QA: all four `TestP1CIdentity*` cases succeeded; complete `internal/resolution` succeeded; the 12-language provider heritage matrix and complete `internal/providers` succeeded.
- Fresh product regression: exactly `61/61` product packages succeeded. Exactly `11` packages under `*/anvien/test/fixtures/*` were excluded as measured non-standalone corpus packages; no additional package was excluded.
- Existing artifact inspection and two independent reruns: tracked/runtime fixture SHA-256 values both equal `327F6BD02FCB341078902FE4302CAB03A2077DDFE98B379164C17E96B547A078`; packaged runtime and a separately compiled current-source runtime each reproduced graph SHA-256 `C6C942B666931CE1AD505E3FE44B52C3887A6C16FBE0B2DAF94C25FE0AF8C1E9` and the `10/10` identity/endpoint oracle. The Supervisor-only alternate build artifact was deleted after verification.
- Current-source native compile: `go build` with production `ladybugdb` tag, trimpath, linker flags, and the repository's v0.19.1 native runtime succeeded to an alternate repo-local output; the binary reported version `1.2.8` and SHA-256 `83D321B5579A471A82DD488BA4DDECFCED6EC83B099955144BE63BC99DE3DB7B` before cleanup.
- Worktree/document integrity: base HEAD is `6b3a3fc4480c7f4c2291d8004207f50b52d91d21`; no staged path; `git diff --check` succeeds; all 17 current modified/new text files are valid UTF-8 without BOM; preserve-only source groups have no diff.
- Verification freshness: current for the reviewed worktree before this report was written.

Failed:

- A fresh canonical `scripts/full-build.ps1` rerun stopped during its first `npm install`/package-runtime step because Windows could not replace `anvien\bin\anvien.exe`. `anvien doctor processes --json` identifies four editor-owned MCP runtimes whose global npm package is a junction to `E:\Anvien\anvien`; they are expected to stay running, so they were not killed. This did not expose a compile or assertion failure: current-source alternate native compilation, version execution, runtime fixture analysis, and the exact product regression succeeded. The prior packaged binary is timestamped after the production and focused identity-test edits; only the later provider QA assertion changed afterward and it compiled/executed in fresh tests. This operational failure does not replace the ledger blocker or authorize a success claim for the fresh canonical rerun.

Not run:

- `anvien detect-changes` and isolated P1-C commit: intentionally remain post-Supervisor gates.
- P1-D conflict/mutation policy, P1-E target `4/4`, `E:\cheapapp.org` analyze, persistence/readers, relationship redesign, and later-Child behavior: out of scope and preserved closed.
- UI/Playwright: not applicable to this non-UI analyzer identity slice.

## Invariant Closure

- affected invariant: every provider/ScopeIR definition occurrence at the production `defsByFile` boundary must map to a deterministic, source-backed, meaning/owner-aware graph identity and emitted endpoint without silent same-name loss; the active ledger must truthfully record the achieved boundary and ordered next action.
- sibling surfaces checked: all provider occurrence-ID constructors, scanner-relative provider input, ScopeIR ownership/normalization, workspace maps, definition emission/endpoints, graph node mutation as preserve-only, relationship merge as preserve-only, legacy index tests, provider heritage parity, all 61 product packages, runtime artifact, all four Child 01 ledgers, roadmap/contract, and prior same-scope review history.
- residual unverified same-invariant surfaces: none in the technical P1-C identity/occurrence boundary. The process/authority invariant remains open because the Detailed Findings contradict the current matrix/log/decision rows. Collision failure policy, final target integration, persistence/readers, and relationship policy are separate closed slices and are not residual P1-C work.

## Required Fix List For Resubmission

1. Reconcile only the stale P1-C Detailed Findings and next-action prose named in the blocking finding, keeping P1-C pending review/detect/cleanup/commit and keeping P1-D/P1-E/target/later Children closed.
2. Resubmit the unchanged production/QA candidate with the ledger-only diff, unchanged source/test hashes, consistency assertions, and `git diff --check` evidence.

## Overall Evaluation

The implementation itself is bounded and technically well-supported: it consumes existing source-proven identity inputs at the existing owner, mathematically separates occurrences through an unambiguous tuple, preserves later pipelines, and reproduces the owned runtime oracle with current source. Acceptance is withheld solely because the active actual-status authority still directs two incompatible P1-C states and can violate the ordered slice gate. The required correction is ledger-only; this review does not reopen or redesign production identity.
