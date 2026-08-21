# Child 05 Pn-A Child-Wide Acceptance Review

Date: 2026-08-22 02:12:18 +07:00
Role: independent Supervisor, existing E-only lane
Authority checkout: `E:\Anvien`
Verdict: REJECT

## Claim reviewed

This review covers the complete Child 05 P5-A through P5-D chain: requested import meanings and type-only state; syntax-derived export tables; deterministic alias/re-export/star/cycle/ambiguity/meaning/member traversal with fail-closed explicit-import behavior; retained terminal/hop/failure proof projected through CALLS/ACCESSES and the four accepted readers; the prior rejection/repair history; committed manifests; sealed runtime and target evidence; four living ledgers; benchmark/evidence closure; and the current Git/worktree boundary.

The production implementation, tests, sealed runtime/target evidence, and commit chain clear independently. Pn-A nevertheless cannot be accepted because the current living-ledger/evidence-closure invariant is internally contradictory and incomplete. This is the only blocking invariant found.

## Authority and review boundary

Read in full before deciding:

- `E:\Anvien\AGENTS.md`.
- `.agents/skills/working-rules/SKILL.md`, `.agents/skills/supervisor/SKILL.md`, `.agents/skills/Data-Integrity/SKILL.md`, and `.agents/skills/Edge-Case/SKILL.md`.
- The complete four-file Child 05 plan set under `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-05-module-export-and-reexport-resolution/`.
- The P5-A, P5-B, and P5-C inventory/candidate/Supervisor reports, including both earlier rejection and exact resubmission histories.
- P5-D Work Step 1 report `97235D73735497328DD7DE41DB0B5023FC6A7ACC05CC46F362A965D0BDB0FB18`, Work Step 2 report `2740EDB71ADFFE1242E192DAEC86C70805F6B656568B49C95C089B981EFECDF2`, and resealed Supervisor report `2A9EC861D443539749C8478D307416EDC257C9178C29049BC5C76309C420471A`.
- Current committed production owners, focused test owners, per-slice diffs, cumulative internal manifest, and relevant preserve-only owners.

Data-Integrity was used as a lens for owned proof projection/persistence parity; Edge-Case was used for ambiguity, cycle, fail-closed behavior, and fallback boundaries. Neither skill created another lane or artifact.

No build, test, analyze, detect-changes, target command, or graph gate was rerun. No access to `E:\cheapapp.org` or any non-E checkout occurred. The only write is this report.

## Git and commit-chain clearance

The current authoritative boundary was independently checked:

- cwd/repo/branch: `E:\Anvien` / `master`.
- HEAD/parent: `831f4d73e27405835c01980859cae5ebd3c9e62b` / `bb4cf46509716259c3bf24a1ca041a6e763d5419`.
- `origin/master...HEAD` is `0 / 50` (zero behind, fifty ahead).
- Index entries: `0`; tracked worktree entries: `0`; `git diff --check`: exit `0`.
- Pre-report status: exactly fourteen protected untracked Main rotation handoffs; none was read, edited, staged, or removed.

The first-parent chain is linear and source-backed:

```text
0aa49c87 Child 04 closure
  -> 2560f914 P5-A implementation (14 paths)
  -> 40ea0095 / 17f3dad2 P5-B transitions
  -> c1559df9 P5-B implementation (11 paths)
  -> cd35b48f / 861000cb P5-C transitions
  -> 76899d45 P5-C implementation (12 paths)
  -> fd6cb52f / 26cb03ee P5-D transitions
  -> bb4cf465 P5-D implementation (15 paths)
  -> 831f4d73 four-ledger transition (4 paths)
```

Every listed implementation commit is an ancestor of the next. The cumulative Child 05 internal diff contains exactly eighteen expected production/test paths. The commit-range diff check from Child 04 closure through current HEAD is clean. No hidden production, test, schema, reader, or target path was found.

## Source-level clearance by slice

### P5-A — requested meanings and type-only input

Cleared. `ImportFact` adds only `RequestedMeanings` and `TypeOnly`; TS/JS default/named/alias/namespace plus statement/inline type-only forms map to the accepted lanes; compatibility re-export facts retain empty requested fields; side-effect-only handling is unchanged. `Normalized`, `NormalizeOwned`, `NormalizeInPlace`, and `compareImport` own clone/canonicalization/order deterministically. All four production and two focused-test files remain byte-identical to commit `2560f914`; their current hashes match the accepted P5-A report.

### P5-B — syntax-derived export tables

Cleared. `export_tables.go` remains byte-identical to commit `c1559df9` and derives explicit entries/star adjacency only from accepted `ExportFact` plus already-resolved target-file candidates. It neither scans physical definitions nor performs terminal traversal; star does not synthesize default. The later orchestration changes preserve this owner and build the table once after all file candidates exist.

The first P5-B rejection was evidence-only: a relative build command and mismatched mutable report identity. The immutable resubmission matches `7,385` bytes / `106` LF / SHA-256 `ADA5725E6D68680519AC72FEC0D3A28D7FB6A2E811F612E62C300A49BA27961A`, records the exact absolute E build command and exit `0`, and closes both blockers without source drift.

### P5-C — deterministic export resolution and fail-closed lookup

Cleared. Current `resolveImports` has one path-candidate phase, one `buildExportTables()` call, then one terminal-binding phase; it does not repeat path resolution. `export_resolution.go` owns deterministic terminal/ambiguity/cycle/missing/meaning-mismatch results, explicit-over-star, no star default, same-terminal proof retention, distinct-terminal no-selection, meaning lanes, namespace/member composition, and owned cloning.

The earlier rejection correctly found that an ambiguous owner could collapse to the sole surviving member branch. The two-file repair retains owner-level ambiguity, gives each owner a terminal or owned missing-member proof, makes `definition()` and `resolveImportedMember` fail closed, and emits neither CALLS nor ACCESSES to that branch. Current repaired hashes match `566A69B9...6248F` and `97FF4990...1620`; `indexes.go` and `resolve.go` preserved the previously cleared candidate bytes at the repair gate. Generic global/name helpers remain unchanged, while the explicit-import no-global guard remains confined to `resolveCall`.

### P5-D — retained proof and generic-first projection

Cleared. `resolvedImport` retains the result from the existing semantic lookup; proof-returning wrappers expose that result without another traversal. `export_binding_proof.go` uses concrete versioned terminal/hop/failure Notes with exact source-site ownership. `resolveCall` and `resolveAccess` preserve generic Evidence as item zero and attach proof only when the retained result identifies the selected endpoint. `mergeRelationship` performs deterministic stable union and exact `(Kind, Weight, Note)` dedupe without changing edge identity or syntactic IMPORTS.

All four production and four test owner hashes match the immutable Work Step 1/P5-D acceptance manifests and are unchanged since commit `bb4cf465`. Current canonical `anvien.exe`, `lbug_shared.dll`, and the fixed-736 corpus artifact still match the sealed hashes `BC060FDD...43532`, `20CBD878...CDB7`, and `7687C519...961A`. The current graph differs from the Work Step 1 build artifact only because the separately recorded pre-commit analyze/detect gate produced its later graph; that chronology is explicit in the ledgers and is not an implementation invalidation.

The sealed Work Step 2 evidence records the bounded target result without any target access in this review: two exact CALLS to one terminal, zero matching gaps, two complete ordered proof chains, four-reader Evidence-hash parity, consumer-to-implementation IMPORTS `0`, and physical/emitted/persisted deltas `0 / 0 / 0`.

## Cross-slice invariant closure

The committed source closes the intended chain:

```text
source-written requested meaning/type-only
  -> accepted syntax-derived export table
  -> deterministic export traversal and fail-closed terminal selection
  -> one retained semantic result
  -> generic-first terminal/hop/failure Evidence
  -> Graph JSON / Ladybug / MCP context / MCP impact parity
```

P5-A production owners, the P5-B table owner, the repaired P5-C traversal owner/test, and all P5-D production owners have no post-acceptance byte drift. Syntactic IMPORTS and physical path behavior remain separate from terminal export lookup. The accepted fixed corpus remains `5,072 / 5,072 / 5,088`, delta `0 / 0 / 0`. No source-level residual same-invariant surface was found.

## Blocking invariant — living-ledger and evidence closure

Pn-A explicitly claims four current living ledgers plus benchmark/evidence closure. The current committed ledger bytes do not satisfy that claim:

1. `...-evidence.md:77` still says `E5-P5A-COMMIT1 | isolated P5-A commit hash | pending`. The same ledger states that a pending evidence ID is a required target rather than proof. Git independently proves commit `2560f914334e65961f755febdda6585840a4260e` with parent `0aa49c87` and exactly fourteen paths, but that does not make the living evidence ledger closed.
2. `...-actual-status.md` has current header/matrix/R15-R17/Next Phase entries showing P5-D accepted, committed, target-validated, and Pn-A open, while its current-state/detail and final-decision text still states the opposite. Concrete contradictions remain at lines `166`, `174`, `176`, `213`, `222`, `224`, `248`, `250`, and `300`: proof is described as lost or `partial/unbound`, target as pending/locked, P5-D as the sole open slice, and target access/evidence as nonexistent.
3. `...-benchmark.md:50-51` leaves the P5-D final values for repository-wide CALLS and ACCESSES generic-Evidence conservation as `pending final`. Work Step 1 proves generic-first behavior on affected proof edges, but none of the three immutable P5-D reports records final absolute `11,467 / 11,467` and `5,974 / 5,974` conservation. Those pending cells therefore cannot be promoted to closed benchmark evidence by inference.

This is one blocker family, not three unrelated source defects: the authoritative living-ledger state does not consistently identify the already-committed implementation and does not close every required evidence/benchmark target. A child-wide Supervisor acceptance would certify mutually incompatible current truths.

## Re-review evidence required

Main must make a ledger-only correction under its own planner authority, without reopening P5 source/test/runtime/target work:

1. Record `E5-P5A-COMMIT1` with commit `2560f914334e65961f755febdda6585840a4260e` and its exact fourteen-path manifest.
2. Reconcile the stale P5-D current-state/detail/final-decision text with R15-R17 and the current plan, or mark genuinely historical text unambiguously so it cannot assert current state.
3. Replace the two `pending final` generic-Evidence benchmark cells with exact final values backed by an already-sealed E artifact or a bounded read-only measurement; do not rerun build/analyze/target analyze merely to update prose.
4. Provide the exact four-ledger diff/hashes and a clean E boundary for a reject-only Pn-A re-review. No source, test, existing report, target, Pn-B, Pn-C, or Child 06 change is required or authorized by this report.

## Evidence checked and intentionally not rerun

Checked: full authorities and reports, raw report identities, current source hashes, per-slice and cumulative diffs, first-parent ancestry, exact commit manifests, current binary/DLL/fixed-corpus identities, four ledger bytes, and start boundary.

Not rerun: build, tests, E analyze, target analyze, target reads, accepted file-detail/impact, detect-changes, stage/commit, cleanup, planner edit, or any command outside `E:\Anvien`. Existing final-byte and target evidence remained sealed and coherent; the blocker is documentary/evidence-state truth, not code invalidation.

State: `IDLE/REJECT`.
