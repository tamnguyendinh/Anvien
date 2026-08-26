# Supervisor Report — Child 06A A003 Canonical Decoder

Verdict: `SUPERVISOR_A003_PASS`

## Metadata

- Review time: `2026-08-26 17:40:13 +07:00`
- Reviewer: fresh visible per-attempt Supervisor (`gpt-5`)
- Repo / HEAD: `E:\Anvien` / `90edf7fe99cd9600b99c1947c2483f8c5fe2c67c`
- Accepted source checkpoint: A002 `ecf825d709b761390a5df4a2147b6ed6eec04499`
- Scope: `P2-A / A003 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
- Authority: `E2-P2A-A003ARCH1/PLAN2`, `AGENTS.md`, `plan-rules.md`, current source, and the two frozen target packets
- Next owner: Main Orchestration

## Exact Candidate Identity and Scope

- `internal/graphhealth/diagnostics.go`: `+109/-7`; SHA-256 `6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30`.
- `internal/graphhealth/diagnostics_test.go`: `+182/-0`; SHA-256 `06677DBA4FE9EDA4FE8651E08B93ABC8659129A61355E34BFFA10380AE429371`.
- The staged set is empty. Concurrent ledger/report changes were not reviewed and are outside this mutation gate.
- Frozen executable: version `1.2.8`, `73,823,232` bytes, SHA-256 `2DBBD7AC70C04FB62C3D8AB4A90F50E7F90C51251E82B67883A58B61D42B426B`.
- Frozen DLL SHA-256: `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`.
- Provenance SHA-256: `CB91D5F7FC2EB2810E6A02AED36B5C7E0791F3973487356336B67A1BC8A66F76`; schema `1`, attempt `A003`, HEAD `90edf7fe`, overlay `2/2`, candidate sources `3/3`, native inputs `3/3`, build exit `0`.

## Fresh Graph Blast Radius

- Required single fresh `anvien analyze --force`: PASS; `2235` scanned, `766` parsed code, `0` failed; graph `123974` nodes / `170800` relationships, non-stale.
- `file-detail` for `internal/graphhealth/diagnostics.go`: containing-file risk `HIGH`; `171` symbols, `58` inbound, `81` outbound, `176` local relationships, `2` linked flows, `20` linked tests.
- Exact upstream impact for `decodeStructuredResolutionOutcome`: risk `LOW`; `7` impacted symbols in one backend/`graph_health` file, `1` direct caller, `2` affected communities, and `0` affected processes.
- Verdict: the file-level HIGH result is a handled scope warning. The candidate remains confined to the exact decoder owner and its authorized test.

## Source Invariant Verdict

PASS.

- `decodeStructuredResolutionOutcome` is the sole production interpreter of the outer `Diagnostic.Note`: one `json.Decoder` object traversal records exact case-sensitive marker presence, applies the existing case-insensitive typed-field behavior, retains type errors as structured-invalid, rejects malformed/non-object/trailing data, and returns the existing `(outcome, structured, valid)` tuple.
- The unchanged policy maps structured-invalid to fail-closed `unclassified/review`; unstructured and structured-valid behavior remains unchanged.
- There is no second outer-note decoder, recovery, legacy, retry, fallback, cache, shared/public decode contract, or additional production owner. The nested `Authority` value remains a preserved `json.RawMessage` and is decoded only by the unchanged nested-authority validator.
- The A002 appender, shared-slice lifetime, write-through graph update, graph serialization, resolution emitters, persistence/readback, and public output owners are unchanged.
- The appended 25-case test-local pre-A003 oracle covers full decode tuple plus classification/actionability parity, including exact/case-variant markers, status evidence, malformed/non-object/null input, typed errors, duplicates, unknown fields, conflicting evidence, invalid proof/authority, and whitespace.

## Consumed Build and Test Truth

- Canonical full build: PASS before tests.
- Focused decoder/appender/P6D graph parity: PASS.
- Focused resolution regressions: PASS.
- Full `internal/graphhealth`: PASS.
- Full `internal/resolution`: truthfully FAIL only at unchanged preserve-only `TestProofBasedCallAccessGoldenCorpus` with `unclassified/review`; the package is not called PASS, the owner was not edited, and this is not an A003 regression.
- These passed/known gates were consumed and were not rerun by Supervisor.

## Cheapapp Equivalence and Resources

Verdict: PASS for accuracy/equivalence/output/lifecycle/resource acceptance.

- Separate A002 -> A003 timings: D001 `3.090914200 -> 3.447846300 s`; parent `19.040468000 -> 20.472602300 s`; analyzer `100.843249000 -> 93.531974900 s`; process `136.729876000 -> 95.630648200 s`.
- Same work: `calls=27890`, `files=887`, operations/children `30/17`, intervals/overlaps `2675/0`, exact conservation PASS.
- Semantic counters, graph/DB counts, source/target identity, stdout, stderr, and canonical graph are equivalent. Current canonical graph is `840,614,023` bytes with SHA-256 `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920`, exactly matching A002.
- Resources: start allocation `-136` bytes; end allocation `-168,537,408` bytes; max observed system memory `+23,179,264` bytes. No retained cache/map/goroutine/lock/I/O/lifecycle owner was introduced.
- The worse Cheapapp D001/parent elapsed values are Main's disposition criterion, not an accuracy rejection condition.

## Restaurant Manager Equivalence and Resources

Verdict: PASS for accuracy/equivalence/output/lifecycle/resource acceptance.

- Separate A002 -> A003 timings: D001 `9.909636600 -> 9.401585300 s`; parent `21.242055400 -> 20.850792800 s`; analyzer `109.339859600 -> 98.020546700 s`; process `145.066210900 -> 101.096911900 s`.
- Same work: `calls=86030`, `files=1234`, operations/children `30/17`, intervals/overlaps `3716/0`, exact conservation PASS.
- Semantic counters, graph/DB counts, source/target identity, stdout, stderr, and canonical graph are equivalent. Current canonical graph is `900,212,685` bytes with SHA-256 `09708B038AD63B659A95BAC34304A752F2C6350952B91CF435EBD5C5F9D94B7C`, exactly matching A002.
- Resources: start allocation `-600` bytes; end allocation `+186,997,152` bytes; max observed system memory `-276,983,808` bytes. The mixed one-run counters do not expose retained state or lifecycle drift; the assigned resource boundary passes.

## Residual Blockers

- Candidate/invariant/raw-evidence blockers: none.
- Residual same-invariant unverified surfaces: none within the assigned A003 acceptance boundary.
- KEEP/no-KEEP, rollback, streak, ledger/checklist, commit, and next-attempt decisions remain exclusively with Main.

## Final Verdict

`SUPERVISOR_A003_PASS`

Next owner: Main Orchestration.
