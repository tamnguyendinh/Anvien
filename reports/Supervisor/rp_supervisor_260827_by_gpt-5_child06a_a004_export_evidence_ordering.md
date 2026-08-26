# Supervisor Report: Child 06A A004 Export-Binding Evidence Ordering

Verdict: `SUPERVISOR_A004_PASS`

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260827_by_gpt-5_child06a_a004_export_evidence_ordering.md`
- Review time: `2026-08-27 00:50:35 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`
- Authoritative docs HEAD: `5c9817effbda337a07e936e109195c95d181fb83`
- Scope reviewed: `P2-A / A004 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
- Claim reviewed: the exact pinned A004 two-file candidate preserves correctness, exact ordered output, downstream graph/persistence/reader/public behavior, determinism, failure and lifecycle boundaries; implements exact-tuple dedupe followed by at most one canonical final-Note key extraction per unique projected record and exactly one cached-key stable final sort; and is supported by credible production-first validation plus two valid independent same-work target packets.
- Authority used: current delegation; `AGENTS.md`; `working-rules`; `supervisor`; `plan-rules.md`; all four Child 06A standard ledgers through EOF; A004 Architect, Planner, Coder, frozen-build, Cheapapp, and Restaurant Manager reports; current source/diff; one fresh Anvien graph; frozen provenance; raw target comparison/validation packets.
- Role boundary: correctness/equivalence/output/lifecycle acceptance only. Performance disposition, `KEEP`/no-`KEEP`, tolerance/noise, rollback, streak, baseline promotion, ledger updates, detect, staging, and commit remain exclusively with Main Orchestration.

## Executive Summary

- Problem: A004 removes a redundant projected-evidence pre-sort and comparator-amplified JSON decoding from the shared export-binding evidence merge path without changing its observable result.
- Decision: PASS. Current source proves one private, merge-local decorate/stable-sort/undecorate implementation. It preserves the old comparison tuple and fallback behavior, has no alternate authority or retained lifecycle owner, and is byte/order-equivalent to the pre-A004 oracle and both accepted A003 target outputs.
- Required outcome: the exact candidate is accepted for correctness/equivalence/output/lifecycle. Main Orchestration now owns the measured disposition; this verdict does not imply performance `KEEP`.

## Candidate Identity And Worktree Boundary

- Current HEAD is exactly `5c9817effbda337a07e936e109195c95d181fb83`.
- Dirty tracked paths are exactly:
  - `internal/resolution/export_binding_proof.go`
  - `internal/resolution/export_binding_proof_test.go`
- Staged set is empty.
- Production diff: `+18/-4`; SHA-256 `36D41BC7336A5A04B4C77D1EE9DEB45DE4E5E2BA120EEFBF98419BD2D846C4D2`.
- Test diff: `+477/-0`; SHA-256 `99C6F6FC5FD0AE4D1BFDFE547D5F67C958544AA61FFD89895DC0BAC79C839BBC`.
- Relevant committed paths are unchanged from Coder base `4695f7c706c5c627d06c60e12d1b3c2217738599` and frozen-build HEAD `6e961f5a4b8bb5fce0f09c722619519b47503e61` through current HEAD. Later commits are documentation/report checkpoints only for this boundary.
- Coder report SHA-256 matches ledger authority: `52C62354D4F3D90BB4C6C922C77E1CA51E3E92169B86F580938625E00301A620`.
- Frozen executable/DLL/provenance identities are exact:
  - executable: `6D319467D198B8BBA2375339CDF6BD7634FA97E7C503D3DFC0C8C315D965352C`, `73,825,280` bytes, version `1.2.8`;
  - DLL: `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`;
  - provenance: `4EBF871D0008B6BB1A7FDB433DE07CE437F859DFB34221C94D360E866711A20D`.
- The current canonical build runtime still matches recorded `A004BUILD1`: `7C077F8555ABDD9D1A76ED19931B448B208622A3B13F170BB71131A463E826F6`.

## Fresh Anvien Evidence And Blast Radius

- `anvien --help`: PASS.
- Exactly one fresh `anvien analyze --force` ran at current HEAD before graph evidence: exit `0`; `2248` scanned / `766` parsed / `0` failed; graph `124309` nodes / `171237` relationships.
- File-detail for `internal/resolution/export_binding_proof.go`: current/non-stale, `changedSinceAnalyze=false`, risk `HIGH`, `105` symbols, inbound `85`, outbound `49`, local relationships `68`, unresolved source sites `120`, linked flows `1`, linked tests `23`.
- Exact upstream impacts, including tests:

| Helper | Risk | Impacted symbols | Direct | Files | Modules | Processes |
|---|---:|---:|---:|---:|---:|---:|
| `appendExportBindingEvidence` | `CRITICAL` | `12` | `7` | `5` | `2` | `30` |
| `mergeExportBindingEvidence` | `CRITICAL` | `25` | `8` | `10` | `2` | `34` |
| `sortExportBindingEvidence` | `CRITICAL` | `16` | `1` | `6` | `1` | `19` |
| `exportBindingEvidenceOrderFor` | `LOW` | `10` | `1` | `4` | `1` | `1` |

The file-level `HIGH` and three current helper-level `CRITICAL` results are scope warnings, not blockers. The fresh extractor result is `LOW`, superseding the older report's all-CRITICAL count for current graph reporting; it does not narrow the one-file preservation boundary. Affected files/processes still include call/access resolution, relationship coalescing, outcome/reference projection, analyzer carriage, and tests, all handled as run-only or equivalence surfaces.

## Source-Level Clearance Notes

- `internal/resolution/export_binding_proof.go`: clear.
  - `appendExportBindingEvidence` now directly hands the unsorted projected records to `mergeExportBindingEvidence` at line `169`; the prior projection pre-sort is the only removed behavior.
  - `mergeExportBindingEvidence` lines `203-221` retains the exact `(Kind, Weight, Note)` map key, existing-first/incoming-second traversal, duplicate removal, generic first-seen partition, and projected partition. Only after that dedupe does line `222` call the sole production `sortExportBindingEvidence(projected)`.
  - `sortExportBindingEvidence` lines `235-259` allocates one private `orderedExportBindingEvidence` slice, calls `exportBindingEvidenceOrderFor` exactly once for each projected record in a linear decoration loop, runs one `sort.SliceStable`, compares only cached `kindRank/proofOrdinal/hopOrdinal/note`, and copies unchanged `graph.Evidence` values back.
  - `exportBindingEvidenceOrderFor` lines `262-290` is unchanged. It remains the sole production key authority, decodes final serialized `Evidence.Note`, keeps terminal/hop/failure/default rank, sets terminal/failure hop ordinal `-1`, uses maximum ordinals on malformed notes, and retains lexical Note as the final tie break.
  - Repo-wide static call inspection finds one production sort call, one production extractor call site in the decoration loop, and no JSON operation inside the comparator. The other comparator/unmarshal path is confined to the test-local pre-A004 oracle.
  - No exported symbol, function signature, public/persisted `graph.Evidence` shape, producer-carried key, overload, second key authority, second final sort, run/global cache, goroutine, lock, I/O, flush, or finalizer was added.
  - The transient slice is local `O(P)` state for `P` unique projected records. Copying `graph.Evidence` copies string headers, not retained Note bytes; no state survives return.
- `internal/resolution/export_binding_proof_test.go`: clear.
  - A004 additions start after the pre-existing tests. File timestamps corroborate production-first ordering: production `2026-08-26 16:37:09Z`, tests `16:44:24Z`, Coder report `16:53:24Z`.
  - The test-only oracle recreates the complete pre-A004 projection pre-sort, merge/dedupe, comparator decoding, rank, ordinal, malformed fallback, lexical tie break, and stable sort; it is not referenced by production.
  - Differential assertions compare full ordered `[]graph.Evidence`, not sets. Coverage includes terminal/hop/failure across proofs/hops, generic/projected mixing, permutations, repeated coalescing, exact duplicates, equal ordinals/different Notes, distinct SourceSiteIDs, malformed JSON for all three kinds, unknown/non-export kinds, stable equal keys with different weights, empty/no-proof paths, caller-input non-mutation/non-aliasing, result non-mutation, and idempotency.

## Invariant Closure

- Exact tuple dedupe: preserved by the unchanged `(Kind, Weight, Note)` key and existing-first/incoming-second traversal; oracle and target graphs agree.
- Generic order: preserved first-seen and stable; explicitly asserted.
- Projected order: terminal/hop/failure/default rank, proof ordinal, hop ordinal, lexical Note, malformed fallback, and equal-key stability use the unchanged key semantics; explicitly oracle-checked.
- SourceSiteIDs and per-site failures: serialized Note construction is unchanged; multi-site oracle assertions and byte-identical target graphs prove carriage.
- Idempotency/non-mutation/aliasing: direct adversarial assertions pass; source creates new merge/output storage and only local transient decoration.
- Marshal/parse failure: marshal helpers and failure returns are untouched; malformed existing evidence is retained and fallback-ordered by the unchanged extractor.
- Public/persisted/lifecycle contracts: no public type/signature/owner changed; state is merge-local and unreachable on return.
- Call/access/relationship/outcome/reference/diagnostic/Graph JSON/public boundaries: focused call/access/outcome, A003 graphhealth/Graph JSON, and analyzer carriage suites are recorded PASS on the exact candidate. Both target canonical graphs and public stdout/stderr are byte-identical to their accepted A003 packets.
- Ladybug/native reader boundary: persistence/reader code is untouched; each target feeds an identical canonical graph and records identical graph/DB row counts and semantic counters. Raw Ladybug/meta byte identity is correctly non-governing by the plan; no reader or logical persistence mismatch is present.
- Determinism/freshness/failure/transaction/temp/publication: ordering is a stable total-key transformation with unchanged inputs/output; target HEAD/status, candidate hashes, graph/public output, and exit boundaries are unchanged; no lifecycle owner was edited.
- Residual unverified same-invariant surfaces: none for the exact A004 acceptance claim.

## Build And Test Evidence

- The exact hash-bound Coder evidence records canonical `scripts/full-build.ps1` PASS, exit `0`, before test execution, with vendor `1798/0/0/0` and maintenance analyze `2243/766/0`.
- A004 differential/export-binding suite: PASS.
- Focused call/access/outcome regressions: PASS.
- A003 diagnostic and Graph JSON/health parity: PASS.
- Analyzer outcome/graph carriage: PASS.
- Full `internal/resolution`: truthfully exit `1` only at the unchanged preserve-only `TestProofBasedCallAccessGoldenCorpus` mismatch (`unclassified/review`). The package is not called PASS; the golden owner has no candidate diff; no new or changed failure exists.
- Production/test/build identity has not changed since this evidence, so the direct prohibition on repeating passed build/test gates applies.

## Independent Target Packet Clearance

### Cheapapp

- Frozen identity: exact A004 executable/DLL/provenance; one launch; exit `0`.
- Raw validation: `25/25` checks pass, no blocker.
- Inventory: `30/30` unique operations and `17/17` unique children; all operation/child denominators equal; D001 `calls=27890; files=887`.
- Conservation: `2675` intervals, overlap `0`; candidate child sum `13.251820500 s` plus residual `0.014178700 s` equals parent `13.265999200 s`; every child duration equals its interval sum.
- Equivalence: workload, parser, graph/DB counts, dependency/projection fields, resolution semantics, diagnostics/outcomes, full ordered Evidence, canonical Graph JSON, stdout, and stderr all pass exactly.
- Raw packet identities: comparison `EE5AA56C278513045A30DE2E69F0B95B5B7B9677D654B13B6B558CC138DC7FC0`; validation `974E37335F816498EFC56643B1E3ADD0B7A8E0FEB608283A4E24D28622FB727C`.

### Restaurant Manager

- Frozen identity: the same exact A004 executable/DLL/provenance; one launch; exit `0`; exact one `electron/renderer/src/api/userApi.ts` exclusion.
- Raw validation: `32/32` checks pass, `failureCount=0`.
- Inventory: before/after `30/30` unique operations and `17/17` unique mapped children; operation order/boundary/denominators and child order/denominators match; D001 `calls=86030; files=1234`.
- Conservation: `3716` intervals, overlap `0`; interval shapes/sums exact; child sum `19.396949800 s` plus residual `0.019149700 s` equals parent `19.416099500 s`.
- Equivalence: full non-timing semantics, files/parser/dependency/projection/diagnostic/outcome fields, graph/DB counts, full ordered Evidence, canonical Graph JSON, stdout, and stderr pass exactly; target/candidate/staged boundaries remain unchanged.
- Raw packet identities: comparison `980980D5467E8E6BD944E4DD30FE59D722E3D19323FDEC292D89EA5D9438DB54`; validation `20EFA2589FEB61920FAC9EBB93BF1721E8940DA8424A722E267C1F7B49E20E41`.

The packets are independent, target-separated, complete, and bound to the same frozen candidate. No canonical graph was re-hash-scanned in this review; the already accepted packet hashes and raw validation records were consumed as directed.

## Objective Measured Facts — No Supervisor Disposition

| Target | Boundary | A003 | A004 | Delta direction |
|---|---|---:|---:|---|
| Cheapapp | D001 | `3.447846300 s` | `2.074182500 s` | lower |
| Cheapapp | parent | `20.472602300 s` | `13.265999200 s` | lower |
| Cheapapp | analyzer | `93.531974900 s` | `107.287054400 s` | higher |
| Cheapapp | process | `95.630648200 s` | `144.975972400 s` | higher |
| Restaurant Manager | D001 | `9.401585300 s` | `8.975767700 s` | lower |
| Restaurant Manager | parent | `20.850792800 s` | `19.416099500 s` | lower |
| Restaurant Manager | analyzer | `98.020546700 s` | `101.406172300 s` | higher |
| Restaurant Manager | process | `101.096911900 s` | `135.569489100 s` | higher |

These facts are preserved without a tolerance/noise label or causal disposition. This report does not decide `KEEP`, no-`KEEP`, rollback, baseline promotion, or streak.

## Evidence Checked

Passed:

- Full required authority chain read through EOF in mandated order.
- Current HEAD/path/hash/staged boundary.
- One fresh Anvien analyze, current file-detail evidence, and four exact upstream impacts.
- Complete live two-file source/test content and diff.
- Static one-extraction/one-final-sort/no-comparator-decode/no-alternate-authority proof.
- Frozen provenance and output identities.
- Raw Cheapapp and Restaurant comparison/validation packets.
- Hash-bound recorded full-build/focused/package evidence.
- Report-only `git diff --no-index --check` emitted no whitespace diagnostics; the staged set remained empty.

Failed:

- None within the correctness/equivalence/output/lifecycle claim.

Not run:

- Full build, tests, target analyze, benchmark, profile, candidate build, and passed measurements were not rerun, by explicit review instruction and unchanged candidate identity.
- `anvien detect-changes`, staging, commit, ledger update, and performance disposition were not run; they are outside this lane.
- Canonical graph files were not re-hash-scanned; accepted packet identity already proves byte equality.

## Overall Evaluation

The exact A004 candidate is a scoped private algorithmic optimization with a mechanically equivalent comparison authority. The source proof, adversarial pre-A004 oracle, unchanged public/lifecycle surface, exact frozen identity, focused validation, and two complete byte-equivalent target packets jointly prove the full requested correctness/equivalence/output/lifecycle claim. No same-scope blocker remains.

Exact verdict: `SUPERVISOR_A004_PASS`.

Next owner: Main Orchestration for measured disposition only. PASS does not imply performance `KEEP`.
