# Coder Report — Child 06A A003 Canonical Decoder

Date: `2026-08-26`
Lane: fresh visible Coder task for `P2-A / A003 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
Status: `CODER_A003_SOURCE_BUILD_READY_FOR_MAIN`
Next owner: Main Orchestration

## Authority and checkpoint

- Clean source HEAD before A003: `90edf7fe99cd9600b99c1947c2483f8c5fe2c67c`.
- Accepted A002 checkpoint ancestor: `ecf825d709b761390a5df4a2147b6ed6eec04499`.
- Authoritative PLAN2 commit ancestor: `45885c95102738b187f643f477d60c669df7f616`.
- Additional preserved ancestors: P1 instrumentation `3436024038cff20ddaa8c4f72f45dbc2e7769806`; package/launcher `56863d1fa1472cde744693a45245a6c0a2865dc6`.
- Execution authority: `E2-P2A-A003CURRENT1/ATTRIB1/ARCH1/PLAN2`. `PLAN1` was not used.
- Pre-edit worktree and staged set were empty. The stale archived Coder task was not inspected or reused.

## Invariant family

- Family: one outer `Diagnostic.Note` interpretation producing the exact `(outcome, structured, valid)` tuple, followed by one existing policy decision and one graph write.
- Production owner: `internal/graphhealth/diagnostics.go::decodeStructuredResolutionOutcome` only.
- Preserved sibling surfaces: `structuredResolutionDiagnosticPolicy`, `validStructuredResolutionDiagnostic`, the nested Authority decoder, A002 appender lifecycle, resolution emitters, Graph JSON/projection, persistence/readers, and public contracts.
- Forbidden routes: no second full-note decoder, recovery, legacy, retry, fallback, cache, shared/public wire contract, or additional semantic owner.
- Residual unverified surfaces inside the assigned source/build/test boundary: `none`.

## Fresh Anvien pre-edit evidence

The pre-clean graph was not reused.

1. `anvien --help`: PASS, exit `0`.
2. Fresh `anvien analyze --force`: PASS, exit `0`; `2231` scanned, `766` parsed code, `0` failed; graph `123887` nodes / `170645` relationships.
3. `anvien file-detail internal/graphhealth/diagnostics.go --repo E:\Anvien --json`: PASS at indexed/current commit `90edf7fe99cd9600b99c1947c2483f8c5fe2c67c`; graph non-stale and `changedSinceAnalyze=false`. File risk is `HIGH`: `164` symbols, `51` inbound, `81` outbound, `161` local relationships, `2` linked flows, `20` linked tests.
4. `anvien impact symbol "decodeStructuredResolutionOutcome" --repo E:\Anvien --direction upstream`: PASS. Exact-symbol risk is `LOW`, while the containing file remains `HIGH`. Full fresh blast radius: `7` impacted symbols, `1` direct caller, `1` module, `0` processes, all in `internal/graphhealth/diagnostics.go`:
   - depth 1: `structuredResolutionDiagnosticPolicy`, one resolved same-file call site at line `354`;
   - depth 2: `normalizeDiagnosticMetadata`, one resolved same-file call site at line `275`;
   - depth 3: `normalizeDiagnostics`, three resolved same-file call sites at lines `208`, `213`, and `219`;
   - depth 3: `AppendDiagnosticToNode`, one resolved same-file call site at line `39`;
   - depth 3: `normalizeDiagnosticSlice`, one resolved same-file call site at line `269`;
   - depth 3: `appendDiagnosticsFromProperties`, one resolved same-file call site at line `140`;
   - depth 3: `diagnosticAppender.appendToNode`, one resolved same-file call site at line `77`.

The containing-file `HIGH` result was treated as a scope warning, not an edit prohibition. No supplemental graph audit or post-gate owner expansion was used.

## Production-first implementation

Production was edited before test bytes.

- `internal/graphhealth/diagnostics.go`: `+109/-7`.
  - Adds decoder-owned `io` import for an exact EOF/trailing-data check.
  - Replaces the envelope-map full-note decode plus typed full-note decode with one `json.Decoder` traversal of the outer object.
  - The same traversal records exact case-sensitive top-level `schemaVersion`/`status` presence, assigns typed outcome fields with existing case-insensitive JSON field semantics, retains type-error state, rejects syntax/non-object/trailing-data states with the pre-A003 tuple, and preserves duplicate/unknown-field behavior.
  - Structured type errors remain fail-closed as zero outcome / `structured=true` / `valid=false`.
  - No private representation was required. No validator, policy, nested Authority decoder, appender, emitter, schema, persistence, reader, contract, or other production owner changed.
- `internal/graphhealth/diagnostics_test.go`: `+182/-0`.
  - Preserves all A002 test content.
  - Adds only `strings` import and `TestDecodeStructuredResolutionOutcomeCanonicalSingleInterpretationParity` at lines `304-483` after production review.
  - Uses a test-local exact pre-A003 two-unmarshal oracle; it is not a production fallback.
  - Compares full decode tuple and policy tuple across 25 cases: status evidence, exact marker evidence, both evidence sources, valid markerless notes, unstructured/empty notes, malformed JSON, non-object JSON, `null`, top-level/nested typed errors, missing/exact/case-variant markers, duplicate keys including persistent type errors, unknown fields, conflicting evidence, invalid proof, invalid authority, and whitespace.

`gofmt` ran on exactly the two owned source/test paths. Scoped `git diff --check -- internal/graphhealth/diagnostics.go internal/graphhealth/diagnostics_test.go` passed with exit `0`.

## Holder preflight and canonical full build

- `anvien doctor locks --repo E:\Anvien --json`: exit `0`; analyze lock absent/free.
- `anvien doctor processes --json`: exit `0`; only editor-owned MCP services using the installed npm runtime were reported. No process was proven to hold a canonical repository build output or repo-scoped build lock, so nothing was terminated.
- Canonical command: `pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`.
- Result: PASS, exit `0`, before every test command.
- Vendor verification: `1798` files, `0` missing, `0` extra, `0` mismatched, `0` reparse points, `0` incomplete license dispositions.
- Canonical runtime: version `1.2.8`, SHA-256 `B98CCF6DF86A106F2D924F33E0E5C7EBCBFF61AA75F2BDD724526DE062B9DD2C`.
- Launcher/Vite build passed with the existing dynamic-import and chunk-size warnings only.
- Maintenance analyze: exit `0`; `2231` scanned, `766` parsed, `0` failed; graph `123937 / 170763`. This is validation, not A003 performance measurement.

## Post-build tests

Tests ran only after the canonical full-build PASS and in the required order.

1. `go test ./internal/graphhealth -run '^(TestDecodeStructuredResolutionOutcomeCanonicalSingleInterpretationParity|TestDiagnosticAppenderMatchesLegacySemantics|TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity)$' -count=1 -v`
   - PASS, exit `0`.
   - New parity test: all 25 subtests PASS.
   - A002 appender equivalence and P6D Graph JSON/health parity PASS.
2. `go test ./internal/resolution -run '^(TestResolveAttachesSourceBackedUnresolvedDiagnostics|TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix|TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite)$' -count=1 -v`
   - PASS, exit `0`; all named tests and subtests PASS.
3. `go test ./internal/graphhealth -count=1`
   - PASS, exit `0`.
4. `go test ./internal/resolution -count=1`
   - Truthful FAIL, exit `1`, solely `TestProofBasedCallAccessGoldenCorpus` at `internal/resolution/proof_accuracy_golden_test.go:166`.
   - The complete actual Diagnostic remains the documented preserve-only payload with `Classification:"unclassified"` and `Actionability:"review"`, including unchanged source/range/site/proof/outcome-note fields.
   - No other test failed. The package is not called PASS. The golden was not edited, and this failure is neither A003 regression nor A003 validation proof.

## Final owned source/test identity

Before the report was written, the only modified paths were the two owned source/test files and the staged set was empty.

| Path | HEAD blob / bytes | A003 blob / bytes | A003 SHA-256 | Diff |
|---|---|---|---|---|
| `internal/graphhealth/diagnostics.go` | `0db79f881b19cfe8cd2f919a5b3ae745ef5b6bbb` / `28325` | `cacf81e15961d1426685d90612f8b2e5a265f7b3` / `31391` | `6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30` | `+109/-7` |
| `internal/graphhealth/diagnostics_test.go` | `d8cd94b4fb1f835150f25d259b8095995fb6a87d` / `10030` | `6b4731f97513fd335bde8541f737fedcd611c46a` / `18730` | `06677DBA4FE9EDA4FE8651E08B93ABC8659129A61355E34BFFA10380AE429371` | `+182/-0` |

No detect-changes, measurement, Supervisor, staging, commit, ledger edit, plan edit, target-repository action, or unrelated cleanup was performed.

## Exact rollback ownership

Rollback only these A003 bytes by applying a scoped reverse patch, never a broad checkout/reset/stash:

1. In `internal/graphhealth/diagnostics.go`, remove the A003 `io` import and replace only the current canonical decoder body at lines `378-502` with the HEAD/A002 body from blob `0db79f881b19cfe8cd2f919a5b3ae745ef5b6bbb`. Preserve every other A002 byte in the file.
2. In `internal/graphhealth/diagnostics_test.go`, remove only the A003 `strings` import and appended test function at lines `304-483`, restoring blob `d8cd94b4fb1f835150f25d259b8095995fb6a87d` without deleting or changing the A002 test body.

No blocker exists inside the authorized source/build/test boundary. Candidate measurement and independent Supervisor acceptance remain intentionally unperformed and belong to Main/later visible lanes.

CODER_A003_SOURCE_BUILD_READY_FOR_MAIN

## A003 Frozen Candidate Packet

- Invocation: exactly one call to `scripts/build-a00x-benchmark.ps1` with `AttemptId=A003`, accepted overlay `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\optimized-overlay.json` / SHA-256 `7B138FBA06B41CBD4C6709E6DF00C80C03E4015B615D1E42F787DE2A65D378F7`, output `E:\Anvien\.tmp\child06a_a003_a00x_benchmark_build\anvien-a003-benchmark.exe`, the Owner-supplied mapped-source/candidate/native hash arrays, and Go `C:\Program Files\Go\bin\go.exe`. No retry and no `analyze` launch occurred.
- Result: PASS, exit `0`; script status `A00X_BENCHMARK_BUILD_COMPLETE`. Exactly three top-level artifact files exist beside the packet-local `.a00x-A003-work` directory:
  - executable `anvien-a003-benchmark.exe`: version `1.2.8`, `73,823,232` bytes, SHA-256 `2DBBD7AC70C04FB62C3D8AB4A90F50E7F90C51251E82B67883A58B61D42B426B`;
  - pinned `lbug_shared.dll`: `20,230,656` bytes, SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`;
  - provenance `anvien-a003-benchmark.provenance.json`: `13,012` bytes, SHA-256 `CB91D5F7FC2EB2810E6A02AED36B5C7E0791F3973487356336B67A1BC8A66F76`.
- Provenance validation: schema `1`; attempt `A003`; HEAD `90edf7fe99cd9600b99c1947c2483f8c5fe2c67c`; build exit `0`; exact overlay mappings `2/2`; candidate sources `3/3`; native inputs `3/3`; every expected/actual hash pair and both output hashes match.
- Preserve boundary: before this section was appended, the report SHA-256 remained `065F9798B8B777E1588F408D96792C9E3DEF3352D2D5E5B6A85C5E38EFB7A5FB`. Source/test hashes remained `diagnostics.go=6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30`, `diagnostics_test.go=06677DBA4FE9EDA4FE8651E08B93ABC8659129A61355E34BFFA10380AE429371`, `emit.go=73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060`, and `outcome.go=02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E`. Staged set was empty and tracked diff remained exactly the two A003 graphhealth paths. A concurrent Main-owned untracked Investigation report present at script start/end was untouched.
- Next owner: Main Orchestration for the separately governed target-measurement lanes. This packet build is not measurement, acceptance, detect, stage, commit, ledger, plan, script, source, test, target, or canonical-output work.

CODER_A003_CANDIDATE_PACKET_READY_FOR_MAIN

## A003 Rollback Result

- Owner/Main disposition: `A003 = NO_KEEP / ROLLBACK` after `SUPERVISOR_A003_PASS`; Cheapapp D001 worsened `3.090914200 -> 3.447846300 s` and parent worsened `19.040468000 -> 20.472602300 s`, so the accepted A002 checkpoint remains authoritative and Main records D001 no-KEEP streak `1`.
- Applied only a surgical reverse patch to the A003 canonical-decoder/import hunk in `internal/graphhealth/diagnostics.go` and the A003-appended parity oracle/import in `internal/graphhealth/diagnostics_test.go`. No whole-file overwrite, checkout, reset, restore, stash, or other source/test edit was used.
- Post-rollback SHA-256: `diagnostics.go=AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8`; `diagnostics_test.go=58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA`.
- Validation: `git diff --numstat ecf825d709b761390a5df4a2147b6ed6eec04499 --` is empty for both paths; scoped `git diff --check` PASS; staged set empty; no source/test diff remains. Main-owned ledgers and measurement/Supervisor reports were preserved untouched.
- No build, test, benchmark, profile, analyze, detect, stage, commit, or candidate-packet cleanup was performed. Next owner: Main Orchestration.

CODER_A003_ROLLBACK_COMPLETE

## A003 Owner KEEP Exact Restore

- Starting SHA-256 matched accepted A002 exactly: `diagnostics.go=AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8`; `diagnostics_test.go=58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA`.
- Reapplied the original `apply_patch` payloads from rollout lines `433`, `461`, and `489`, in that exact order; no code was reconstructed or redesigned.
- Final SHA-256 matches the already-built/measured/`SUPERVISOR_A003_PASS` candidate exactly: `diagnostics.go=6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30`; `diagnostics_test.go=06677DBA4FE9EDA4FE8651E08B93ABC8659129A61355E34BFFA10380AE429371`.
- Scoped `git diff --check` PASS; numstat is `diagnostics.go +109/-7`, `diagnostics_test.go +182/-0`; staged set empty.
- No passed gate was rerun: no Anvien, build, test, benchmark, profile, Supervisor, target, detect, stage, commit, cleanup, ledger, or plan action. Next owner: Main Orchestration.

CODER_A003_OWNER_KEEP_RESTORE_COMPLETE
