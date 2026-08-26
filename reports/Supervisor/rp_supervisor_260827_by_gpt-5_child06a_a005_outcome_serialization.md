# Supervisor Report: Child 06A A005 Outcome Serialization

Verdict: `SUPERVISOR_A005_PASS`

## Metadata

- Review time: `2026-08-27 05:25:34 +07:00`
- Reviewer: `gpt-5`
- Repository: `E:\Anvien`, local shared checkout
- Slice: `P2-A / A005 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
- Claim reviewed: the exact uncommitted A005 four-path candidate and both corrected target-separated packets satisfy correctness, equivalence, output, lifecycle, and resource acceptance.
- Authority: current delegation; `AGENTS.md`; `working-rules`; `supervisor`; binding `plan-rules.md`; all four Child 06A ledgers through EOF; A005 Architect, Planner, Coder, Coder-handoff, frozen-build, and target reports; accepted A003 target reports as immutable before/output authority.
- Disposition boundary: this is not a `KEEP`, `NO_KEEP`, streak, rollback, restore, detect, stage, commit, or timing-promotion decision. Main Orchestration owns the next disposition.

## Executive Summary

The exact candidate passes. Source inspection confirms one unchanged record-time marshal for each unique retained `SourceSiteID`, one private run-scoped canonical-string sidecar, equal-duplicate byte reuse, strict bijective finalization, projection consumption without encoding or fallback, and unchanged public result unwrapping. The focused production-first tests and recorded build/package evidence close duplicate, conflict, first-error, clone, coverage, ordering, determinism, carrier, and byte-parity behavior without hiding the known preserve-only golden failure.

Both corrected target packets independently preserve the full same-work and output boundary. Each exits `0`, has one corrected launch, `30/30` operations, `17/17` children, exact denominators, interval conservation, zero positive-width overlap, timing-independent semantic equality to its own A003 packet, byte-identical canonical Graph JSON, and byte-identical stdout/stderr. Aggregate resource fields are exposed. A distinct runtime canonical-payload/sidecar counter is truthfully `NOT EXPOSED`; the `O(U+B)` one-payload lifecycle is instead closed structurally by current source and focused tests.

No concrete correctness, equivalence, output, lifecycle, or resource defect is present.

## Acceptance Questions

### 1. Locked single-encoding/private-sidecar design

**PASS.**

- `internal/resolution/outcome.go:74-105` adds one private `encodedBySourceSite` map, initializes it per collector, writes it only after the unchanged unique record-time marshal succeeds, and returns the stored string for equal/conflicting duplicates without a second marshal.
- `internal/resolution/outcome.go:115-145` returns the private `finalizedResolutionOutcomes` bundle and rejects missing or extra semantic/byte coverage before projection.
- `internal/resolution/outcome.go:416-473` validates and consumes supplied canonical strings. It contains no call to `marshalResolutionOutcome`, repair path, missing-byte tolerance, fallback encoder, or alternate serialized shape.
- `internal/resolution/resolve.go:110-126` passes the private bundle to projection and exposes only `outcomes.values` through the unchanged public `Result.ResolutionOutcomes` shape.
- A scoped production search finds the sole production `marshalResolutionOutcome` call at the unique record boundary. Other matches are the unchanged helper definition or test-local expected-byte construction.
- No public/persisted type, JSON tag, emitter, diagnostic policy/decoder, graph type, persistence reader, CLI, instrumentation, A00x script, A001-A004 owner, or global/cross-run state changed.

### 2. Production-first tests and recorded validation

**PASS.**

- New A005 tests cover all seven status families; first-add/equal-duplicate reuse; conflicts and first-error precedence; input/returned/finalized deep-clone isolation; immediate diagnostic parity; relationship and both reference-index carriers; missing/extra coverage; record-time non-finite marshal failure; duplicate finalized sites; resolved/unresolved overlap; diagnostic/reference drift; SourceSiteID ordering; complete evidence order; and deterministic replay.
- The sentinel duplicate test proves an equal duplicate consumes the previously stored sidecar value rather than re-marshaling.
- The existing P6C3 test changes only three finalized-result accesses and one direct private projection construction; its substantive assertions remain intact.
- Recorded canonical full build exited `0` before tests. The exact A005-focused plus nine named regressions exited `0`; `internal/graphhealth`, `internal/analyze`, `internal/lbugload`, and `internal/lbugnative` exited `0`.
- `internal/resolution` truthfully exited `1` only at unchanged `TestProofBasedCallAccessGoldenCorpus`. It is not called PASS; the golden blob remains exactly `00d0ae042f109cd7d3c8aeadf855d192a5aa8172`. No new or changed failure is recorded.
- These passed gates were not rerun because the exact source/test hashes remain unchanged and no source mutation invalidated them.

### 3. Independent two-target same-work/output proof

**PASS.**

| Target | Corrected identity | D001 / parent / analyzer / process | Structure and parity |
|---|---|---|---|
| Cheapapp | launch `1`, exit `0`, `2026-08-26T21:41:20.947306Z..21:44:44.5461379Z` | `3.036901000 / 18.160962900 / 95.376559900 / 203.598831900 s` | `calls=27890; files=887`; `30/30`; `17/17`; `2675` interval pairs; child sum/residual `18.135869200 / 0.025093700 s`; zero positive-width overlap |
| Restaurant Manager | launch `1`, exit `0`, `2026-08-26T21:41:56.7304748Z..21:44:39.0222374Z` | `9.142619400 / 19.678482100 / 122.900035100 / 162.290202200 s` | `calls=86030; files=1234`; `30/30`; `17/17`; `3716` interval pairs; child sum/residual `19.618545400 / 0.059936700 s`; zero positive-width overlap |

- A timing-independent, leaf-for-leaf comparison excludes only labels, elapsed fields, interval timestamps, and separately reported memory. It is exact for each target. Reviewer hashes are `61B0EBD29C407AFEAA4E688CC42D101FF4A4E0D0B168B449F8F742373F3D1910` for both Cheapapp packets and `05BB2D963078FA9ADDBC03CE6EA6C2A3B1E6C6502E29F81AE3F19F9423BCF5A6` for both Restaurant packets.
- Cheapapp canonical Graph JSON is byte-identical to A003: `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920`. Restaurant is byte-identical: `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`. This preserves complete serialized Evidence content and order.
- Public stdout hashes remain `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` and `CD0028F3B3939451FF4D3623DF7DD7482A993E0934D2DFA773682D37D4FB0E94`; stderr remains `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` on both.
- File/parser, graph/DB, dependency/projection, resolution counter, diagnostic/outcome, and all other non-timing benchmark fields equal each target's A003 authority.
- The original launches produced no benchmark and exited `1` only because the profile environment input was omitted. The failed-root benchmark artifacts remain absent. They are measurement-support failures, not production attempts or streak events.

### 4. `O(U+B)` one-payload lifecycle

**PASS, structurally proven; runtime counter not exposed.**

- A unique retained site has one existing semantic-map entry and one canonical-string sidecar entry. The only sidecar assignment follows the sole unique record-time marshal.
- Equal duplicates return the stored string and do not allocate another encoded payload. Conflicts retain the prior tuple and first error. Marshal failures retain no sidecar payload and fail before projection.
- Finalization returns the existing sidecar map with cloned/sorted semantic values; projection performs lookups only. Relationship, reference, and diagnostic assignments copy Go string headers rather than retaining another canonical encoder or payload owner.
- The collector is created with the emitter, used by one sequential `ResolveBoundInto`, sealed by that production control flow at finalization, and becomes unreachable after return except for strings retained by returned graph carriers. There is no global/cache/I/O/goroutine/lock/flush/finalizer/TTL owner.
- Aggregate resource observations remain bounded and target-separated: Cheapapp start/end/max-Sys `1,315,192 -> 1,317,600`, `886,583,856 -> 908,981,224`, `1,995,295,224 -> 1,820,223,992` bytes; Restaurant `1,474,592 -> 1,476,832`, `1,306,049,888 -> 785,046,824`, `2,634,607,224 -> 2,651,937,400` bytes.
- Neither corrected runtime packet contains a distinct canonical-payload, sidecar, or encoding-count field. That limitation is recorded as `NOT EXPOSED`; no runtime lifecycle counter is invented.

### 5. Concrete defect check

**PASS.** No correctness, equivalence, output, lifecycle, or resource defect was found in the exact candidate or either corrected packet.

## Exact Candidate And Frozen Identity

| Path | SHA-256 |
|---|---|
| `internal/resolution/outcome.go` | `18203DFAB9A227B526F8F7478B516AE6673F635BABC02D9463975E428A3983AF` |
| `internal/resolution/resolve.go` | `76B7B62A060B36EE2438E76689E858544358AD681DB42F6A4FC47D271F1749A1` |
| `internal/resolution/p6c3_structured_outcome_test.go` | `E561280B3F8420D2288001179431A37FC39E6E2B8564863C9AD644EEE005A2E6` |
| `internal/resolution/outcome_serialization_test.go` | `89B168B2764B9A9B8EACDAB37A071BE0E1C51A8F791FE158B111651E9640C957` |

- Tracked production/test numstat is `outcome.go +43/-24`, `resolve.go +2/-2`, and P6C3 `+8/-5`; the new focused test contains `652` lines. Total is `+705/-31`.
- Frozen executable: `73,825,792` bytes, SHA-256 `0178DAC396E0DD4D0BECCE7B803378A3400013F0B508EC95F9221A0020CF96F3`.
- Frozen DLL: `20,230,656` bytes, SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`.
- Provenance: `13,322` bytes, SHA-256 `E4B9C7E9F9850A58356D41131C7CC177CFC6207D7D8CD3E158B633524B56244E`; schema/attempt/exits `1/A005/0/0`; mappings/candidate/native counts `2/4/3`; every expected/actual hash pair matches.

## Fresh Anvien Blast Radius

- Exactly one fresh `anvien analyze --force` completed before graph verification: exit `0`, `2,260` scanned, `767` parsed, `0` failed, graph `124,555 / 171,598`; indexed/current commit `3f517ac65bc0ef197a7f5a8e3cd216e669669619`, stale `false`.
- `internal/resolution/outcome.go` and `internal/resolution/resolve.go` are both HIGH-risk, current, and non-stale. The bounded summaries report `124` and `200` symbols, `11` and `26` linked flows, and `24` and `41` linked tests.
- Every edited production symbol is CRITICAL: collector `33` impacted / `49` processes; constructor `4 / 23`; exact UID-bound `record` `4 / 10`; finalized bundle `5 / 23`; exact UID-bound `finalize` `6 / 32`; projection `6 / 32`; `ResolveBoundInto` `7 / 31`.
- These are broad blast-radius warnings. The private two-file implementation, focused carrier tests, sibling package evidence, and exact two-target output parity close the warned scope for this acceptance decision.

## Factual Limitations And Not Run

- The corrected runs overlap for `162.292 s`; the Restaurant interval is entirely contained within the Cheapapp interval. This was permitted by the target-separated parallel measurement contract. It limits isolated-machine interpretation of analyzer/process wall timing but does not weaken source correctness or the exact per-target output/equivalence packets. Main alone applies timing promotion rules.
- The distinct canonical-payload/sidecar lifecycle counter is `NOT EXPOSED` at runtime; acceptance relies on direct source/test lifecycle proof plus bounded aggregate resource observations.
- No canonical full build, tests, target analyze, benchmark, profile, prior Supervisor gate, detect-changes, staging, commit, restore, or cleanup was rerun by this review.

## Invariant Closure

- Affected invariant: one canonical record-time JSON payload per retained source site, shared by immediate and final consumers without changed semantic/error timing, ordering, public/persisted shape, or expanded lifecycle owner.
- Same-invariant surfaces checked: collector/constructor/record/finalize/bundle; projection; public result wiring; all status families; duplicates/conflicts/errors/clones; diagnostic, relationship, and both reference-index carriers; graph/Graph JSON; graphhealth/analyze/Ladybug/native regression evidence; both target packets; aggregate resources.
- Residual unverified same-invariant surfaces: none for correctness/equivalence/output/lifecycle/resource acceptance. Timing promotion is deliberately outside this verdict.

## Overall Evaluation

The exact A005 candidate and both corrected packets satisfy the locked acceptance boundary. The result is `SUPERVISOR_A005_PASS`. Next owner is Main Orchestration for disposition only.
