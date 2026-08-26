# Planner Report — Child 06A A005 Outcome Serialization

Date: `2026-08-27`
Role: visible Planner for Child 06A A005
Scope: `P2-A / A005 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
Verdict: `PLANNER_A005_READY_FOR_MAIN_VERIFY`
State transition: `A005_ARCHITECT_COMPLETE / PLANNER_PENDING / CODER_LOCKED -> A005_PLAN_READY_FOR_MAIN_VERIFY / CODER_LOCKED`
Decision boundary: exact architecture-to-plan translation only; no architecture redesign, source/test/script/target edit, build, test, target analyze, profile, measurement, Supervisor, disposition, detect, stage, commit, P3, Child 07, or Coder authorization
Repository basis: HEAD `e6ab44705b9abb4d1a55094cfb1a2bce1f06ea8d` (`docs(plan): record Child 06A A005 architecture`)
Accepted performance checkpoint: A003 `b6bf45bce95323aa6b53b182edfea8628bd8b463`
Accepted storage checkpoint: `0f3a572331dd23d17688886fcbfebeb7d37ee35d`

## 1. Authority consumed

The Planner read the full binding sources through EOF in the required order:

1. `E:\Anvien\AGENTS.md`;
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`;
3. `E:\Anvien\.agents\skills\planner\SKILL.md`;
4. binding `plan-rules.md`;
5. all four standard Child 06A ledgers;
6. `E:\Anvien\reports\system-architect\rp_system-architect_260827_by_gpt-5_child06a_a005_outcome_serialization.md`; and
7. `E:\Anvien\reports\Supervisor\rp_supervisor_260827_021244_by_gpt-5_child06a_a005_architect_handoff.md`.

Architect verdict is `ARCHITECT_A005_READY_FOR_PLANNER`; Main handoff verification is `PASS`. There is no residual architecture ambiguity inside A005. Planner consumed the fresh graph/file-detail/impact packet already recorded by `E2-P2A-A005ARCH1` and did not rerun attribution, architecture, profiles, target analysis, broad graph work, A001-A004, WAL, build, tests, or measurement.

Planner workflow evidence:

- `anvien --help`: PASS;
- initial HEAD: `e6ab44705b9abb4d1a55094cfb1a2bce1f06ea8d`;
- initial worktree: clean;
- initial staged set: empty; and
- temporary directory creation: none.

This evidence proves only the plan-authoring basis. It is not product, graph, source-impact, validation, performance, or acceptance evidence.

## 2. Current campaign state preserved

- P2-A, parent `B1-P1A-OP001`, and child `B2-P2A-A001-D001` remain active and unchecked.
- D001 unsuccessful-attempt streak remains `1`; A004 `NO_KEEP` created that exact count.
- D002-D017 remain queued, unchecked, and unopened.
- Accepted A003 target bases remain separate and unchanged:
  - Cheapapp D001/parent/analyzer/process `3.447846300 / 20.472602300 / 93.531974900 / 95.630648200 s`; `calls=27890; files=887`.
  - Restaurant Manager D001/parent/analyzer/process `9.401585300 / 20.850792800 / 98.020546700 / 101.096911900 s`; `calls=86030; files=1234`.
- A004 remains `SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE`. Its source/test identities are restored, and its candidate measurements remain rejected history only.
- A005 planning creates no candidate, benchmark value, allocation result, speed claim, Supervisor result, disposition, baseline promotion, rollback, or streak transition.
- Coder remains locked. Planner does not open, command, or release Coder.

## 3. Locked A005 architecture translated

The living plan now binds one semantic authority and one byte authority per retained SourceSiteID:

```text
clone incoming ResolutionOutcome
-> validate at the existing record boundary
-> existing conflict/equal-duplicate decision
-> retain unique valid semantic clone
-> unchanged marshalResolutionOutcome exactly once
-> private run-scoped SourceSiteID -> canonical JSON string sidecar
-> one private finalized values-plus-sidecar bundle
   -> immediate diagnostic bytes when required
   -> final relationship/reference/diagnostic byte-parity carriers
-> public Result unwraps only sorted semantic values
```

`projectResolutionOutcomes` remains in A005 ownership only as strict consumer and validator. It must not call the marshaler, regenerate missing bytes, tolerate missing/extra coverage, introduce a fallback encoder, or create a second serialized shape.

## 4. Exact future production boundary

Only these production surfaces may change after separate Main release.

### `internal/resolution/outcome.go`

- Add only one private canonical encoded sidecar to `resolutionOutcomeCollector`.
- Initialize only that sidecar in `newResolutionOutcomeCollector`.
- Preserve the `record` signature and all existing return semantics while storing/reusing one canonical string.
- Add one private `finalizedResolutionOutcomes` representation containing SourceSiteID-sorted cloned values and the sealed sidecar.
- Make `finalize` return that bundle while preserving first-error precedence, semantic revalidation, cloning, SourceSiteID sorting, and fail-closed one-to-one coverage validation.
- Make `projectResolutionOutcomes` consume the private bundle, remove only its marshal/local-encoded-map construction, and preserve every carrier/overlap/duplicate/parity validation.

### `internal/resolution/resolve.go`

- Change only the existing `finalize -> project -> error/result ResolutionOutcomes` wiring block in `ResolveBoundInto`.
- Pass the complete private bundle to projection and unwrap only its values for `Result.ResolutionOutcomes`.
- Preserve the exported signature, loop order, branch order, metrics, graph/reference ownership, error return shapes, and every other line.

Explicitly preserve `marshalResolutionOutcome` body/signature, `validateResolutionOutcome`, `cloneResolutionOutcome`, all outcome construction methods, `recordTypeScriptLookup`, `emitTypeScriptOutcomeDiagnostic`, `emitUnresolvedReference`, `projectReferenceIndexOutcomes`, and `resolutionOutcomeDiagnosticSites`. All graph/public/persisted types, readers, CLI, instrumentation, scripts, A001-A004 owners, and D002-D017 owners are preserve-only.

## 5. Exact record/finalize/projection semantics

1. Clone before validation/retention exactly as today.
2. Validate before duplicate lookup. Invalid input remains unretained, returns `added=false`, and records the first error.
3. Compare duplicate semantic clones exactly as today.
4. Conflicting duplicates retain the existing first error and return the prior clone/prior canonical bytes—or empty bytes if its initial marshal failed—with `added=false`.
5. Equal duplicates return the prior immutable clone/string with `added=false`, emit no duplicate diagnostic, and never re-marshal.
6. A unique valid outcome is retained before its marshal, preserving lifecycle and initial error timing. The unchanged marshaler runs once at that same record boundary; canonical bytes are stored only on success; `added=true` remains success-only.
7. A record-time marshal failure returns empty bytes and `added=false`, suppresses the immediate diagnostic, retains the contextual first error, and causes `finalize` to fail before projection. Duplicates never retry or hide it.
8. `finalize` preserves revalidation, deep cloning, and SourceSiteID-only sorting, then fails closed unless semantic and encoded inventories are exactly one-to-one.
9. Projection consumes supplied bytes for relationship/reference evidence and diagnostic parity validation; it does not encode, repair, or fall back.

## 6. Tests only after production is correct

Allowed test owners are exactly:

- new `internal/resolution/outcome_serialization_test.go`; and
- existing `internal/resolution/p6c3_structured_outcome_test.go` only for the smallest mechanical adaptation of three direct private `finalize` calls and one direct private `projectResolutionOutcomes` call.

The A005 focused matrix must prove:

- one canonical byte string for repository-resolved, intrinsic-resolved, TypeScript resolved-external, repository-unresolved, TypeScript capability-unavailable, profile-excluded, and meaning-mismatch cases;
- pre-A005 `json.Marshal` byte identity for immediate unresolved/non-resolved TypeScript `Diagnostic.Note` and equality with final projection parity bytes;
- resolved relationship and both reference-index carrier bytes equal the same canonical string;
- first-add/equal-duplicate `added` behavior, no duplicate diagnostic, and reuse of the prior tuple;
- conflicting SourceSiteID fail-closed behavior with unchanged first-error precedence;
- mutation isolation for input, returned, finalized, target, proof evidence, authority, and declaration-range objects;
- record-time marshal failure using a naturally unsupported non-finite `graph.Evidence.Weight`, including empty bytes, `added=false`, no diagnostic, contextual error, and `finalize` failure before projection;
- missing/extra encoding coverage failure with no re-encoding fallback;
- rejection of duplicate finalized source sites, resolved-plus-unresolved overlap, diagnostic/outcome payload drift, and non-resolved reference-carrier drift;
- SourceSiteID-sorted outcomes, unchanged graph/reference traversal, complete ordered evidence, and repeated-run determinism.

Test-local expected-byte/oracle logic cannot become a production fallback or second production encoder. Every existing P6C3 assertion outside the four mechanical call adaptations remains unchanged.

## 7. Future Coder pre-edit gate

Only after Main independently verifies the five Planner-owned artifacts and separately releases one Coder may that Coder run:

1. `anvien --help`;
2. exactly one fresh `anvien analyze --force`;
3. file-detail for `internal/resolution/outcome.go` and `internal/resolution/resolve.go`; and
4. exact upstream impact for every symbol it will edit.

The Architect packet records both files HIGH and `record`, `marshalResolutionOutcome`, `finalize`, `projectResolutionOutcomes`, and `ResolveBoundInto` as CRITICAL boundaries. These are scope warnings, not edit bans. Failure to prove current graph/source identity or the exact owner list is a STOP before editing. If production requires any additional owner, Coder stops before tests and returns the blocker to Main.

## 8. Holder, full-build, focused-test, and package sequence

After production is correct and authorized tests are complete:

1. Run `anvien doctor locks --repo E:\Anvien --json` and `anvien doctor processes --json`.
2. Prove and terminate only actual build-output holders; start no build while any such holder remains.
3. Run canonical `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1` and require exit `0`.
4. Only after full-build PASS, run the new A005 focused tests and these exact existing regressions:
   - `TestP6C3P5ProofNestingConflictAndImmutableReplay`;
   - `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`;
   - `TestResolveAttachesSourceBackedUnresolvedDiagnostics`;
   - `TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite`;
   - `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`;
   - `TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage`;
   - `TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus`;
   - `TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity`;
   - `TestP6C3NativeLadybugResolutionOutcomeReadback`.
5. Run package regressions for `internal/resolution`, `internal/graphhealth`, `internal/analyze`, `internal/lbugload`, and `internal/lbugnative`.
6. Run scoped diff-check on the exact candidate boundary.

The known `TestProofBasedCallAccessGoldenCorpus` mismatch may be recorded only when its exact preserve-only payload/baseline remains unchanged. The package is never called PASS. Any new or changed failure blocks A005. Detect, stage, and commit remain locked until later measurements, Supervisor, and Main disposition.

## 9. Unchanged A00x build and target-separated measurement contract

- Reuse `scripts/build-a00x-benchmark.ps1` unchanged with explicit A005 attempt, overlay, output, mapped-source, candidate-source, native, and hash inputs. Do not repeat build-interface discovery and do not edit either canonical build script.
- Because A005 changes overlay-mapped `resolve.go`, mechanically regenerate/verify that mapping from exact A005 production source plus accepted 17-child instrumentation. A stale A003/A004 overlay that erases A005 wiring is a STOP. Keep the `types.go` mapping unchanged unless exact compilation evidence proves otherwise.
- Frozen identity is accepted A003/WAL state plus only A005-authorized production/test bytes and the accepted measurement overlay. No rejected A004 byte may enter it.
- Run Cheapapp once with its accepted options and Restaurant Manager once with its accepted options plus exactly one `electron/renderer/src/api/userApi.ts` exclusion. Each target uses its own accepted A003 packet as `before`; never rerun, average, or combine those bases.
- Each packet records D001, parent, analyzer, process wall, `30/30` operations, `17/17` children, calls/files, interval conservation/zero overlap, workload, diagnostics/outcomes, full ordered evidence, Graph JSON, graph/DB readback, stdout/stderr, semantic counters, `startAllocBytes`, `endAllocBytes`, `maxObservedSys`, and private `O(U+B)` one-payload lifecycle proof.

## 10. Supervisor and disposition boundary

Only after both complete independent packets exist may one fresh visible A005 Supervisor review the exact candidate and affected correctness/equivalence/output/lifecycle/resource boundary.

Main alone decides `KEEP`, `REWORK`, `ROLLBACK`, or streak effect. Planner makes no performance claim and changes no benchmark value, baseline, streak, checkbox, or queue. Current streak is `1`; one unsuccessful A005 can move it only to `2`, never terminal `3`.

## 11. Preserved invariants and private resource bound

Preserve exactly:

- ResolutionOutcome schema, JSON tags/field order, status/stage/site/target/reason/proof/authority shapes, and contextual marshal errors;
- SourceSiteID identity, duplicate/conflict detection, first-error retention, deep clone semantics, and `added` behavior;
- immediate `Diagnostic.Note` and relationship/reference `Evidence.Note` byte parity;
- A002 diagnostic appender write-through/bucket/count/order/`[]Diagnostic` behavior and `O(T)` lifecycle;
- A003 canonical decoder tuple/policy/fail-closed behavior;
- graph nodes, relationships, IDs, labels, properties, counts, reference indexes, outcome inventory, and complete evidence order;
- analyzer result carriage, Graph JSON, Ladybug/native persistence/readback, affected readers, stdout/stderr/public output, and metrics;
- determinism, freshness/invalidation, failure propagation, transaction rollback, temporary-file flush/close/rename, publication visibility, and known golden truth; and
- A001 import-claim work, A002 appender, A003 decoder, WAL cleanup, P1 timing instrumentation, accepted target workloads/options/denominators, and restored A004 identities.

Let `U` be unique valid retained outcomes and `B` the total canonical JSON byte length. Retained A005 state is `O(U+B)`: one map entry and at most one encoded backing payload per retained SourceSiteID, no duplicated outcome object in the sidecar, only string-header copies in carriers, run-scoped lifetime, and no global/cross-run cache, I/O, goroutine, lock, concurrency, flush, finalizer, or TTL.

## 12. Exact rollback

Rollback only A005-owned bytes:

- collector canonical sidecar and initialization;
- `record` reuse hunk;
- private finalized bundle and `finalize` return;
- strict projection-consumer change, restoring projection-time marshal/local map only as the exact A005 reverse;
- narrow `ResolveBoundInto` wiring;
- new focused test file;
- four mechanical existing-test call adaptations; and
- rejected A005 frozen/overlay packet.

Do not reset, checkout, stash, overwrite whole files, or disturb A001-A003, WAL, P1 instrumentation, reports, ledgers, target work, user/protected changes, or restored A004 code/test identities.

## 13. Mandatory STOP conditions

STOP and return the exact blocker to Main if:

- any production/test file or production symbol outside the exact allowed list is required;
- `marshalResolutionOutcome`, immediate diagnostic callers, schema/type/JSON tags, diagnostic policy/appender/decoder, graph/reference/public types, persistence/readers, CLI, instrumentation, or the A00x script must change;
- record-time validation or initial marshal timing moves later;
- conflict detection, first-error retention, `added`, clone immutability, or SourceSiteID behavior changes;
- immediate unresolved/non-resolved TypeScript diagnostic bytes are not identical;
- projection gains a fallback/second encoder, missing-byte tolerance, or altered carrier validation;
- outcome/evidence/diagnostic order or restored A004 behavior changes;
- retained resources exceed `O(U+B)`, duplicate payload bytes per carrier/site, survive the run, or require global state/concurrency/I/O/locks/flush/finalizers;
- full build fails, any new/changed test failure appears, or the known golden changes;
- the A005 overlay lacks exact production wiring, the unchanged script contract cannot build it, or candidate identity is mixed;
- either target packet is missing/incomparable, changes workload/denominator, lacks exact equivalence, or shows an unexplained resource/sibling regression;
- D002-D017, another parent, P3, or Child 07 opens; or
- one future unsuccessful A005 is incorrectly treated as terminal streak `3` rather than streak `2`.

## 14. Owned output, verification, and next owner

Planner owns only:

1. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`;
2. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-evidence.md`;
3. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-benchmark.md`;
4. `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-actual-status.md`; and
5. this report.

Final verification:

- HEAD remains `e6ab44705b9abb4d1a55094cfb1a2bce1f06ea8d`.
- Changed-path count is exactly `5`, equal to the four standard Child 06A ledgers plus this report; no source/test/script/target/Architect/Main report or other path changed.
- Scoped `git diff --check` over the exact five allowed paths returned exit `0`; the untracked Planner report has `0` trailing-whitespace lines.
- All `53` plan checkbox lines match HEAD exactly.
- All `1,470` benchmark elapsed-value tokens match HEAD exactly, and the Attempt Numeric History block matches HEAD exactly.
- `git diff --cached --name-only` is empty; nothing is staged.
- Planner performed no detect, stage, or commit.

Next owner: Main Orchestration for independent verification of exactly these five artifacts and a separate future Coder-release decision.

`PLANNER_A005_READY_FOR_MAIN_VERIFY`
