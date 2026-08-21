# Architecture Decision — Child 04 P4-C2 Non-Circular Evidence Lifecycle

Created: `2026-08-21 10:38:12 +07:00`

Role: Evidence Architect / Data-Integrity advisor

Scope: Child 04 P4-C2 only

Implementation basis: P4-C closed at `c99c4070b66e7a96be8c9fa2721a0335a1f94877`

Target-access statement: this architecture lane did not read, stat, hash, analyze, execute Git in, or write `E:\cheapapp.org`.

Plan-edit statement: this lane did not edit the roadmap, any Child 04 ledger, the contract, production, tests, golden files, QA assets, or historical reports.

Final verdict: `READY_FOR_PLANNER`

## 1. Decisive answer

**YES — the corrected lifecycle is already authorized by the open P4-C2 scope and the latest Owner invariant. No new Owner decision is required.**

P4-C2 already makes the target source `inspect-only`, the target worktree `preserve-only`, the independent 21-entry oracle an input, and the normal target analyze output the later actual-value input. The latest Owner invariant resolves the only ambiguity: source-only inspection is the authorized way to derive ground truth, while observing analyzer/QA output before sealing the oracle is forbidden. The plan phrase “before any target read” must therefore be replaced with “before any analyzer/QA-output observation”; it cannot be read as a ban on the source-only read needed to author the oracle.

The historical `.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle-output.json` is conclusively reclassified as an **invalid-lifecycle debug capture**. It is not an accepted oracle, not a lost acceptance artifact, and not recoverable/promotable authority. The historical QA and Recovery reports remain immutable historical records, but their “accepted raw oracle” and “restore/recover `.tmp` to unblock” method is rejected and must not be copied into the corrected plan.

## 2. Authority read in full

| Authority | Bytes / LF | SHA-256 |
| --- | ---: | --- |
| `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md` | `27,109 / 191` | `FFD76CF50666FB24300FADAF24401F15C1D04D45154799D0FDFE69B48DC4CFE4` |
| Child 04 plan | `39,001 / 384` | `CD0E928F7258FBDDAA477B4F7E8CDB1991BFDD7AA6E0B486448707B652CA9523` |
| Child 04 evidence ledger | `38,056 / 290` | `AACD9C12832C8966536CFB061742EE9714E9E80A83EAB9C4ACBCAFF98D3684BE` |
| Child 04 benchmark ledger | `9,658 / 74` | `B2A6274E8C219CF096520193382367A4322A610EC7B8312C42A82FD4785D8592` |
| Child 04 actual-status ledger | `27,130 / 194` | `AE831B58B49EE9A5E5410296F67014905092B2502EA55379667C8FE56B1FB89F` |
| `docs/contracts/graph-accuracy-contract.md` | `9,104 / 100` | `68CB65EF964E6D3D7BB8697BD786AE1451DADB1B36D10CC38B5F9CA3839F2592` |
| `reports/QA/rp_qa_260821_091213_by_gpt-5_child04_p4c2_oracle_gate.md` | `13,142 / 188` | `CAD7F4300A8B5B2EB0FD9A59808030DCC1ACCDB8B1DB6C5B25784B1D513B175D` |
| `reports/Investigation/rp_investigation_260821_0932_by_gpt-5_p4c2_oracle_recovery.md` | `15,552 / 204` | `AC5B5A6DF78B329E3DE89C1C50F2B215CFA438A298D4F1F7B3C6E3227BE4797F` |

The full `E:\Anvien\AGENTS.md`, `working-rules`, `System-Architect`, and `Data-Integrity` skill instructions were also read before this decision. The latest Owner invariant delivered to this lane is later and more specific than the two historical P4-C2 reports, so it controls their lifecycle interpretation.

## 3. Governing invariant

Let:

- `S` be the exact hash-pinned target source basis;
- `E = F_source(S)` be the 21 positive expected rows plus source-derived negative controls authored without Anvien output;
- `H(E)` be the sealed oracle bundle identity;
- `A = F_analyzer(B, S)` be actual values produced later by the normally built Anvien runtime `B`;
- `C = Compare(E, A)` be the QA comparison.

The mandatory invariant is:

```text
source-only authorship completes
  -> E is written directly to a durable approved path
  -> H(E) is sealed and routed
  -> QA verifies H(E) and the unchanged S basis
  -> only then may B run and produce A
  -> QA computes C without editing E
  -> Supervisor independently accepts or rejects C
```

At no point may `A`, current implementation behavior, current tests/goldens, a prior QA interpretation, or any `.tmp` artifact contribute a value to `E`. Expected and actual values have different owners, different allowed inputs, different durable bundles, and a one-way dependency: QA consumes the sealed oracle; Oracle Authoring never consumes QA/analyzer output.

Chronology is a gate, not a narrative claim:

```text
oracle bundle SEALED
  < Main routes bundle digest to paused QA task
  < QA verifies seal and source basis
  < full build/analyzer command starts
  < actual capture/comparison is created
```

The pre-existing historical QA report does not contaminate a new Oracle Authoring lane only if Main creates that lane with clean context and does not expose the report or its contents to the author. This architecture report does not author any expected row.

## 4. Why target source is valid ground truth

Target source is the valid derivation authority because the graph-accuracy contract defines canonical graph facts as distinctions measured from source occurrences, source ranges, language meanings, export entries, and provenance. Child 04 owns “what the source exports”; Child 05 separately owns terminal module/re-export resolution. TypeScript syntax in hash-pinned source therefore determines the expected direct-export fact independently of Anvien’s current behavior.

Source inspection remains independent when all of these are true:

1. the source bytes are pinned before derivation;
2. only language/source semantics are used to fill expected values;
3. no target source bytes or excerpts are copied into Anvien—only fact tuples, ranges, names, and hashes are retained;
4. the author cannot see Anvien graph output, `.anvien`, analyzer stdout/stderr, QA comparisons, current implementation code, tests, or golden files;
5. a later lane compares, rather than authors, actual values.

The current analyzer cannot author expected values because it is the system under test. Using its Graph JSON, Ladybug rows, command output, implementation code, tests, or golden files to fill expected tuples would reduce validation to `actual == actual` or to two representations derived from the same possibly defective logic. QA output is also downstream interpretation and cannot backfill expected values. Historical line/name summaries can identify the bounded denominator, but they do not become the semantic oracle; the clean author must re-derive complete tuples from the pinned source.

## 5. Valid lifecycle and exact lane sequence

### Stage 0 — Planner synchronization, no target access

Planner patches only the roadmap and four Child 04 ledgers with the wording in section 13. The slice remains P4-C2; no production, test, golden, contract, target, QA asset, Supervisor, Child 05, or analyzer action opens. All five documents must state that evidence-bearing artifacts are born durable and that `.tmp` can never be promoted.

### Stage 1 — Main routes a clean Oracle Authoring lane

Main pauses the existing QA continuation and opens a separate clean-context Oracle Authoring lane. The prompt may contain the corrected P4-C2 authority, schema, three accepted file/hash/slot anchors, and this lifecycle decision. It must not contain the historical QA/Recovery report text, current analyzer output, or current implementation/test/golden content.

Main must use a no-history/clean-context handoff for the Oracle lane. If the platform cannot prevent inherited QA/analyzer context, the lane is ineligible and must stop before target access.

### Stage 2 — Oracle lane captures the source-only basis

The Oracle lane may perform read-only source/worktree metadata operations needed to establish `S`:

- target Git HEAD and branch identity;
- tracked-worktree status without copying diff content;
- SHA-256 and byte count of the three whitelisted source files;
- direct reading of only those source files;
- exact one-based ranges/selection ranges and lexical owners for the 21 positives and 11 negative controls.

It must not list, stat, hash, or read target `.anvien`; run Anvien; run build/test/QA; inspect Anvien production/tests/goldens; inspect prior QA/analyzer captures; or write anything in the target. The pre-existing target worktree is preserved even if it is not clean.

The accepted source basis is:

| Repository-relative source file | Required SHA-256 | Positive anchors (one-based start lines) |
| --- | --- | --- |
| `modules/email/server/operations/email-operations-observability.ts` | `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C` | `35, 48, 55, 61, 68, 76, 82, 91, 99, 101, 109, 118, 150, 158, 254, 438, 497` |
| `modules/release-distribution/server/release-distribution-publication-state.ts` | `B53F966F82AEDAC58EC15A892B9B71BA0A5716638703105BB35B126CB398D474` | `1, 7, 13` |
| `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts` | `44467B1D6F1778AE4728B4941136321F7831D8FD67EBE2C93A430A3D132F90AC` | `10` |

Any hash mismatch, missing/duplicate anchor, or anchor that no longer resolves uniquely to the expected source occurrence is `ORACLE_SOURCE_BASIS_MISMATCH`; no oracle may be sealed and no analyzer may run. The lane must report the mismatch to Main without asking Owner and without inventing a new denominator.

### Stage 3 — Oracle lane authors durable expected values

All oracle-bearing files are created directly under the durable path in section 6. No intermediate input/output, raw capture, manifest, provenance, expected values, or reproducibility material may ever be created in `.tmp`, even briefly. The lane records complete positive and negative tuples but stores no source excerpt or copied target file.

The author verifies exact counts, schema completeness, unique source occurrence identity, and negative-control cardinality before sealing. A partially written durable bundle is `UNSEALED` and cannot be routed to QA.

### Stage 4 — Oracle lane seals and closes

`seal.json` is written last. After sealing, no file in the oracle bundle may be appended, reformatted, renamed, or overwritten. The Oracle lane returns only the durable bundle path, `oracleId`, `bundleDigest`, `seal.json` SHA-256, source-basis HEAD/hashes, row counts, and `SEALED` status to Main. It does not run validation.

### Stage 5 — Main resumes the existing QA lane

Main sends the existing QA continuation only the sealed path and identities. QA independently recomputes every file SHA-256, bundle digest, schema/count gate, and source-basis hash before reading expected semantics. A mismatch stops the lane as `ORACLE_SEAL_INVALID`.

QA must never edit or regenerate expected values. If the oracle is incomplete or wrong, QA rejects it and returns to Main; it does not fix the oracle from observed actuals.

### Stage 6 — QA captures pre-run boundary, builds, and runs normal analyze

Only after seal verification may QA:

1. recapture target HEAD, full tracked/untracked status metadata, and the three source hashes;
2. capture a pre-run path/size/SHA-256 manifest of target-local `.anvien` without treating old output as actual evidence;
3. clear build-related holders as required by `AGENTS.md`, run the full canonical build, and record the built runtime identity;
4. invoke the normal command grammar `anvien analyze E:\cheapapp.org --force` through the normally built runtime;
5. capture stdout, stderr, exit state, command/config/runtime identity, and affected Graph JSON/Ladybug/reader rows directly into the durable QA run path;
6. compare actual rows against the sealed expected rows;
7. capture target post-state and prove that only analyzer-owned target-local `.anvien` operational output changed.

An old target graph, self-graph, substituted artifact, failed command output, or pass-by-default test can never replace the fresh normal run.

### Stage 7 — Main routes later Supervisor

After QA writes a complete JSON and Markdown result bundle, Main routes the sealed oracle bundle plus QA run bundle to a fresh Supervisor lane. Supervisor verifies source basis, chronology, hashes, row comparison, target pre/post boundary, output ownership, and absence of Child 05-derived claims. Only Supervisor `PASS` allows Planner/Main to close evidence IDs, detect/stage/commit the authorized Anvien-side evidence, and later open Child 05. This architecture verdict is not that acceptance.

## 6. Exact durable artifact locations and formats

### Oracle bundle

Canonical directory:

```text
reports/QA/child04-p4c2/oracle/<oracle_id>/
```

`<oracle_id>` must be immutable and collision-free, using `p4c2-oracle-v1-<target_head_12>-<YYMMDD_HHMMSS+0700>`.

Required files, all created directly in that directory:

| File | Format / owner | Required content |
| --- | --- | --- |
| `oracle.schema.json` | UTF-8 JSON Schema, LF | exact v1 row and bundle validation contract |
| `source-basis.json` | UTF-8 JSON, LF | target HEAD/branch/status identity; three relative paths, byte counts, SHA-256 values; 21 positive and 11 control anchors; no source bytes |
| `expected-values.json` | UTF-8 JSON, LF | exactly 21 positive rows and 11 negative rows in deterministic order |
| `provenance.json` | UTF-8 JSON, LF | author task/role/skills, authority identities, semantic input allowlist, denylist, exact read/hash commands, zero target writes, zero forbidden observations, creation timestamps |
| `authoring-report.md` | UTF-8 Markdown, LF | human-readable derivation/count review, source-only attestation, fail-closed result, no copied source |
| `seal.json` | UTF-8 JSON, LF; written last | file byte counts/SHA-256 values, row counts, source-basis identity, bundle digest, seal timestamp, status `SEALED` |

No symlink, junction, target-side mirror, `.tmp` staging area, or later promotion is permitted.

### QA run bundle

Canonical directory:

```text
reports/QA/child04-p4c2/runs/<run_id>/
```

`<run_id>` must use `p4c2-run-v1-<target_head_12>-<YYMMDD_HHMMSS+0700>`.

Required files, created directly in that directory:

| File | Required content |
| --- | --- |
| `oracle-seal-verification.json` | supplied/recomputed oracle identities and `PASS`/failure reason |
| `target-pre.json` | HEAD/branch/status metadata, three source hashes, and pre-run `.anvien` manifest |
| `build.json` | canonical full-build command, exit, runtime/binary identity, and warnings classification |
| `analyze-command.json` | exact argv, cwd, environment/config identity, start/end time, exit code |
| `analyze.stdout.txt` / `analyze.stderr.txt` | direct durable raw command streams; never staged through `.tmp` |
| `actual-values.json` | only the affected normalized actual records needed for the 21 positives and 11 controls, with source-occurrence mapping and no expected-value mutation |
| `comparison.json` | row-by-row expected/actual result, aggregate `21/21`, `11/11`, drift/orphan/forbidden-state counts |
| `target-post.json` | post-run HEAD/status/source hashes, post-run `.anvien` manifest, and contamination delta |
| `run-report.md` | human-readable QA result, commands, what each proves, failures, and evidence IDs |
| `evidence-manifest.json` | byte count/SHA-256 for every run artifact and deterministic run-bundle digest; written last |

The final repo-native root QA handoff report may additionally follow the existing `reports/QA/rp_qa_<YYMMDD>_<HHMMSS>_by_<model_slug>_child04_p4c2_*.md` convention, but it must cite the immutable oracle/run bundle identities rather than duplicate or replace them.

## 7. Provenance and sealing identity

The oracle seal must be non-circular. `seal.json` does not hash itself. It lists the other five bundle files in ordinal filename order. For each file, construct this exact UTF-8 line from the actual bytes:

```text
<filename>\0<byte_count_decimal>\0<SHA256_UPPER_HEX>\n
```

Then compute:

```text
bundleDigest = SHA256(
  UTF8("p4c2-oracle-v1\n")
  + all five sorted file-identity lines
)
```

`bundleDigest` is the oracle sealing identity. Main also records the ordinary SHA-256 and byte count of `seal.json` at handoff. QA recomputes both before any build/analyzer action. Filesystem timestamps are provenance only, never the identity.

`provenance.json` must include at least:

- `oracleId`, `evidenceId: E4-P4C2-ORACLE1`, author task/thread ID, author role, and skills;
- corrected roadmap/four-ledger/contract identities and P4-C commit;
- target HEAD/branch and exact source-file identities;
- governance inputs separately from semantic inputs;
- semantic-input allowlist containing only the three target source files;
- explicit forbidden-input list: target `.anvien`, any analyzer/QA output or report, current Anvien implementation/tests/goldens, old `.tmp` captures, self-graph, Child 05 output;
- exact commands/resources read and their purpose;
- `forbiddenInputsObserved: []`;
- `analyzerOrQaOutputObservedBeforeSeal: false`;
- `targetWrites: []` and `evidenceArtifactsCreatedInTmp: []`;
- creation and seal timestamps in UTC+7.

An attestation without the allowlist/command inventory/file hashes is insufficient provenance.

## 8. Exact row schema

`expected-values.json` must have top-level fields:

```text
schemaVersion = "p4c2.oracle.v1"
oracleId
evidenceId = "E4-P4C2-ORACLE1"
targetBasisRef = "source-basis.json"
positiveCount = 21
negativeCount = 11
positiveRows[]
negativeControls[]
```

Every positive row must contain all of these non-optional fields:

| Field | Contract |
| --- | --- |
| `rowId` | unique `P001`…`P021`; stable deterministic ordering by relative path then start position |
| `polarity` | literal `positive_direct_export` |
| `source.path` | normalized repository-relative path, `/` separators |
| `source.fileSha256` | uppercase 64-hex source identity |
| `source.range` | one-based start/end line and column of the declared construct |
| `source.selectionRange` | one-based start/end line and column of the declaring token |
| `source.lexicalOwner` | explicit module/top-level owner identity; never inferred by name alone |
| `declarationKind` | source declaration kind independently read from TypeScript source |
| `localName` | exact local binding name; explicit `null` only when source has no local binding |
| `expected.exportedName` | exact exported name |
| `expected.exportKind` | exact Child 04 syntax kind, derived from source form |
| `expected.meanings` | non-empty deterministic array drawn from `value`, `type`, `namespace`; dual meanings preserved |
| `expected.typeOnly` | explicit boolean |
| `expected.exportFactCount` | exactly `1` for each positive row |
| `expected.fileToExportRelationCount` | exactly `1` for each positive row |
| `expected.access.state` | `absent` or `present`, derived independently from source access semantics |
| `expected.access.value` | explicit value or `null`; export must not overwrite it |
| `expected.compatibility.isExported` | exactly `true` for the bounded direct definition |
| `expected.exportDiagnosticCount` | exactly `0` for a supported positive row |
| `expected.child05DerivedState` | empty object/list; no terminal target, traversal, ambiguity, cycle, link status, transitive path, or public-API claim |
| `derivation` | concise source-semantic rationale without copying a source excerpt and without citing analyzer behavior |

Every negative-control row uses the same source identity/range/owner fields and contains:

```text
rowId = N001…N011
polarity = "negative_not_direct_export"
expected.exportFactCount = 0
expected.fileToExportRelationCount = 0
expected.compatibility.isExported = false
expected.access = exact source-derived preserved state
expected.child05DerivedState = empty
```

Rows match by `(source.path, fileSha256, exact occurrence range, lexicalOwner, localName)`, never by name alone. Implementation-generated graph IDs are forbidden oracle fields because the source-only lane cannot independently know them; QA maps actual graph records back to source occurrence before comparison.

## 9. Exact negative controls

All 11 controls are in `modules/email/server/operations/email-operations-observability.ts` at the accepted file hash above. The Oracle lane must independently confirm their exact range/owner from source and seal all 11:

| Row | Accepted start line | Local name | Lexical owner | Expected direct-export result |
| --- | ---: | --- | --- | --- |
| `N001` | `190` | `boundedReadLimit` | containing function/source owner confirmed from source | no Export fact; `isExported=false` |
| `N002` | `207` | `time` | `parseTime` | no Export fact; `isExported=false` |
| `N003` | `214` | `time` | reducer callback in `latestIso` | no Export fact; `isExported=false` |
| `N004` | `262` | `now` | `buildEmailOperationsReport` | no Export fact; `isExported=false` |
| `N005` | `501` | `now` | `readEmailOperationsReport` | no Export fact; `isExported=false` |
| `N006` | `504` | `messageRows` | `readEmailOperationsReport` | no Export fact; `isExported=false` |
| `N007` | `505` | `attemptRows` | `readEmailOperationsReport` | no Export fact; `isExported=false` |
| `N008` | `506` | `eventRows` | `readEmailOperationsReport` | no Export fact; `isExported=false` |
| `N009` | `507` | `providerEventRows` | `readEmailOperationsReport` | no Export fact; `isExported=false` |
| `N010` | `508` | `readinessRows` | `readEmailOperationsReport` | no Export fact; `isExported=false` |
| `N011` | `509` | `suppressionRows` | `readEmailOperationsReport` | no Export fact; `isExported=false` |

The same-name `time` and `now` controls are specifically required to prove occurrence/owner isolation. If any line no longer identifies the exact occurrence under the accepted hash, the source basis is inconsistent and the lane fails closed; it must not remap by name.

## 10. Exact target pre/post boundary

The target boundary has two deliberately separate captures.

### Source-only authoring boundary

Before reading source semantics, Oracle Authoring records target HEAD/branch, tracked status metadata, and the three source hashes in `source-basis.json`. It does not inspect `.anvien`. After authoring and immediately before seal, it recomputes those three hashes and records equality. Any source/worktree write or hash drift invalidates the bundle.

### QA runtime boundary

After the oracle is sealed, QA records in `target-pre.json`:

- HEAD/branch;
- raw porcelain status metadata plus a normalized status that excludes only allowed `.anvien/**` operational paths;
- SHA-256 for all three oracle source files;
- path, byte count, SHA-256, and existence state for pre-run `.anvien` files;
- no source/diff content.

After the normal analyzer run, `target-post.json` repeats the same capture and computes the delta. PASS requires:

- HEAD unchanged;
- normalized non-`.anvien` worktree status byte-for-byte unchanged from pre-state;
- all three source hashes unchanged from both sealed source basis and QA pre-state;
- every new/changed target path owned by normal analyzer output and under `.anvien/**`;
- no target-side report, fixture, probe, copied source, script, or debug artifact;
- the pre-existing target worktree state preserved rather than cleaned or rewritten.

A target that is dirty before the run is not automatically invalid; loss or alteration of that pre-existing state is invalid.

## 11. Independence controls

| Boundary | Required control |
| --- | --- |
| Main → Oracle | clean/no-history context; do not forward QA/Recovery report text or analyzer output |
| Oracle semantic input | only the three hash-pinned target source files |
| Oracle governance input | corrected P4-C2 authority/schema only; it may constrain fields but not supply actual values from Anvien |
| Oracle forbidden surfaces | `.anvien`, Anvien production/tests/goldens, build output, analyzer/QA reports/results, `.tmp` oracle material, Child 05 state |
| Oracle writes | only its new durable Anvien-side bundle; zero target writes |
| Main handoff | route path and seal identities only; Main must not edit rows |
| QA expected input | sealed bundle, read-only after independent digest verification |
| QA actual input | one fresh normal built-analyzer run on the same source basis |
| QA mutation | may write only a new durable QA run bundle and normal target `.anvien`; never expected values |
| Supervisor | fresh independent review; does not repair oracle or QA output |

No person/task/lane may serve as both expected-value author and actual-value validator for this slice.

## 12. Fail-closed conditions

### Oracle Authoring must stop without seal if any condition occurs

- semantic-input context contains or reveals analyzer/QA output, current implementation/tests/goldens, target `.anvien`, or historical `.tmp` oracle contents;
- any evidence-bearing artifact is created in `.tmp`, on the target, or outside the approved durable bundle path;
- one of the three source hashes differs from the accepted basis;
- positive anchors are not exactly 21 unique source occurrences (`17 + 3 + 1`);
- negative controls are not exactly the 11 unique owner-qualified occurrences;
- any required row field is missing, guessed from line number/name alone, or derived from Anvien behavior;
- the author copies source bytes/excerpts into Anvien;
- target source/worktree changes during authoring;
- seal inputs, counts, or hashes do not verify.

### QA must stop without comparison PASS if any condition occurs

- the oracle path is not durable, mentions `.tmp` as an input, is unsealed, or has a digest/byte/count/schema mismatch;
- the source basis at QA pre-run differs from the sealed basis;
- the full build fails or the runtime/command is not the normally built analyzer path;
- analyze fails, is skipped, uses old output, or actual records cannot be mapped uniquely to source occurrences;
- any positive row has zero, duplicate, or differing fact/field/edge values;
- the aggregate is not exactly `21/21`;
- any of the 11 negative controls produces an Export fact, `isExported=true`, or name-only collision;
- access state changes because of export extraction, compatibility differs, a supported positive emits a diagnostic, an orphan exists, or any Child 05-derived state appears;
- non-`.anvien` target state changes or a target-side QA/report/fixture/probe/debug artifact appears;
- any QA artifact was staged through `.tmp` or expected values changed after seal.

### Supervisor must reject closure if any condition occurs

- chronology cannot prove seal-before-QA/analyzer;
- Oracle and QA lanes are not independent;
- manifest/provenance identities are missing or inconsistent;
- only counts pass while row-level fields/parity/boundary are unproven;
- historical `.tmp` material is treated as accepted or promoted evidence;
- any P4-C2 evidence ID is closed by a pending, blocked, historical, or pass-by-default artifact.

Failure never authorizes QA to patch expectations. A new expected bundle, if needed, must come from a new clean source-only authoring lane and receive a new `oracle_id`; the old sealed bundle is never overwritten.

## 13. Exact lane ownership and skill packages

| Lane | Required skills | Owns | Must not own |
| --- | --- | --- | --- |
| Planner synchronization | `working-rules`, `planner` | patch the roadmap + four Child 04 ledgers only; record lifecycle and paths | target access, oracle rows, analyzer, acceptance |
| Main routing | `working-rules`, `orchestration` | enforce task isolation/order, freeze/resume lanes, route seal identities, later route Supervisor | expected-value authorship, actual comparison, acceptance |
| Oracle Authoring | `working-rules`, `Data-Integrity` | source basis, 21+11 expected tuples, provenance, seal | Anvien implementation/tests/goldens, `.anvien`, analyzer/QA output, validation verdict |
| Existing QA continuation | `working-rules`, `qa`, `Data-Integrity` | verify seal, full build, normal target run, actual extraction/comparison, pre/post boundary, JSON+MD evidence | author/edit expected values, production repair, Supervisor verdict |
| Later Supervisor | `working-rules`, `supervisor` | independent acceptance/rejection of source, seal, chronology, QA, target boundary, evidence integrity | invent missing authority, repair oracle/QA, Child 05 work |

This Evidence Architect lane used `working-rules`, `System-Architect`, and `Data-Integrity` to define the one-way pipeline, artifact ownership, integrity identity, and fail-closed gates. It did not invoke Planner or Supervisor because this delegation explicitly forbids ledger edits and acceptance work.

## 14. Literal wording Planner should apply to the five living documents

Planner should preserve historical records and patch current authority. The following blocks are normative; equivalent paraphrase is not preferred because the prior ambiguity was caused by broad wording.

### 14.1 Roadmap — append a P4-C2 oracle lifecycle checkpoint

```markdown
## Child 04 P4-C2 Oracle Lifecycle Checkpoint

- P4-C2 remains the sole open slice. Source-only Oracle Authoring is authorized to inspect the three hash-pinned target source files while the target worktree remains preserve-only. This source read is the ground-truth derivation step; it is not analyzer validation.
- The Oracle Authoring lane must be clean-context and must not observe target `.anvien`, current Anvien implementation/tests/goldens, analyzer output, QA output/reports, Child 05 state, or historical `.tmp` oracle material. It must seal exactly 21 positive rows plus the 11 owner-qualified negative controls before the existing QA lane resumes.
- Every oracle, raw capture, manifest, provenance record, expected-values input/output, benchmark, and reproducibility artifact must be created directly under `reports/QA/child04-p4c2/...` from its first byte. `.tmp` is debug-only and can never close an evidence ID or be promoted/restored as accepted evidence.
- Historical `.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle-output.json` is an invalid-lifecycle debug capture, not an accepted or lost oracle. The QA oracle-gate and Recovery reports remain historical; their `.tmp` acceptance/recovery method is non-authoritative.
- Main routes only the sealed oracle path and digest to the existing QA continuation. QA verifies the seal and unchanged source basis, then performs the full build and one normal `anvien analyze E:\cheapapp.org --force` run, compares actual output without editing expected values, and records the target pre/post boundary. Later Supervisor remains the only acceptance owner; Child 05 stays locked until PASS.
```

### 14.2 Child 04 plan — add to Rules

```markdown
- Evidence-bearing artifacts are born durable: oracle rows, raw captures, command streams, manifests, provenance, expected-values inputs/outputs, benchmark material, and reproducibility files must be written directly under an approved `reports/QA/child04-p4c2/...` path. `.tmp` is debug-only; anything evidence-bearing created there is invalid and cannot be promoted, restored, or used to close an evidence ID.
- P4-C2 separates two one-way lanes. Oracle Authoring derives expected values only from hash-pinned target source and seals them before observing target `.anvien`, current Anvien implementation/tests/goldens, analyzer output, or QA output. The existing QA lane later consumes the sealed oracle read-only and obtains actual values only from the normally built analyzer.
```

Replace the P4-C2 `Scope Boundary`, `Pre-flight Questions`, `Work Steps`, `Implementation Gate`, `Acceptance`, and `Commit Boundary` with this wording:

```markdown
  - Scope Boundary:
    - Oracle Authoring inspect-only: exactly the three accepted hash-pinned target source files plus read-only target HEAD/tracked-status metadata; no target `.anvien` observation.
    - QA inspect-only after seal: target source hashes, target-local analyzer output, and only affected persisted records.
    - Preserve-only: all target source and the complete pre-existing target worktree.
    - Anvien-side writable evidence: new durable oracle and QA bundles under `reports/QA/child04-p4c2/...`, plus this child’s ledgers after valid evidence.
    - Out of scope: production/test/golden repair, target fixtures/reports/probes, copied target source, terminal resolution, other Child defects, and every `.tmp` evidence path.
  - Pre-flight Questions:
    - Expected-value source: direct TypeScript semantics read from the exact hash-pinned target source by a clean-context Oracle Authoring lane; current implementation/analyzer/tests/goldens/QA output are forbidden inputs.
    - Actual-value source: one fresh normal built-analyzer run on the same sealed source basis, performed only by the existing QA continuation after seal verification.
    - DB read flow: Oracle Authoring reads no `.anvien`; QA reads target-local graph/affected persistence only after seal.
    - DB write flow: normal target-local `.anvien` operational output only during QA; every Anvien-side evidence artifact is written directly to its durable approved path.
    - Behavior test: exact row-level comparison for 21 positives and 11 owner-qualified negative controls, access/export separation, compatibility parity, zero diagnostics for supported positives, zero orphans, zero Child 05-derived state, and target pre/post preservation.
    - Cleanup/quarantine: `.tmp` may contain only disposable debug data that is unnecessary for any gate or rerun; no evidence-bearing artifact may originate there.
  - Work Steps:
    1. Main opens a clean Oracle Authoring lane and keeps QA paused. The lane verifies target HEAD/status plus the three accepted source hashes, reads only those source files, authors exactly 21 positive and 11 negative rows directly in `reports/QA/child04-p4c2/oracle/<oracle_id>/`, writes provenance, and writes `seal.json` last. Any source mismatch, forbidden input observation, target write, `.tmp` evidence creation, incomplete row, or seal mismatch stops before analyzer validation.
       - Evidence target: `E4-P4C2-ORACLE1`.
    2. Main routes only the sealed path/digest to the existing QA continuation. QA independently verifies the seal and unchanged source basis, captures target pre-state, performs the canonical full build, runs the normal built analyzer, and writes raw/normalized actual output directly under `reports/QA/child04-p4c2/runs/<run_id>/`.
       - Evidence target: `E4-P4C2-TARGET1`.
    3. QA compares without editing expected values, records `21/21`, `11/11`, exact field parity, access separation, zero forbidden state, and target post-state/contamination. Main then routes later Supervisor; only after Supervisor PASS may ledgers/detect/commit close the slice.
       - Evidence target: `E4-P4C2-BOUNDARY1`, `E4-P4C2-REVIEW1`, `E4-P4C2-DETECT1`, `E4-P4C2-COMMIT1`.
  - Implementation Gate: P4-C is accepted and committed at `c99c4070b66e7a96be8c9fa2721a0335a1f94877`. Source-only target inspection for Oracle Authoring is authorized now; analyzer/QA-output observation is forbidden until the durable 21+11 oracle bundle is sealed. Target source/worktree remains preserve-only and Child 05 remains locked.
  - Acceptance:
    - Oracle: durable-from-creation source-only bundle, complete schema/counts/provenance, valid seal, and no forbidden observation or `.tmp` lifecycle violation.
    - Source: `21/21` exact target definitions have one correct direct-export fact; 11 owner-qualified controls remain non-exported; access state is unchanged.
    - Runtime/data: fresh normally built analyzer matches every sealed field, compatibility parity has zero drift, supported positives have zero diagnostics, orphan count is zero, and Child 05-derived state is absent.
    - Boundary: target HEAD, pre-existing non-`.anvien` worktree state, and three source hashes are unchanged; only normal analyzer-owned `.anvien` operational output changes.
    - Evidence: JSON+Markdown oracle/QA bundles are complete; later Supervisor PASS, detect, and isolated commit are recorded.
  - Commit Boundary: after Supervisor PASS, commit only valid Anvien-side durable oracle/QA evidence and the authorized five living-document refresh; never commit or depend on target artifacts or `.tmp` material.
```

### 14.3 Evidence ledger — add rule, historical disposition, and replace P4-C2 rows

Add this evidence rule:

```markdown
- Every evidence-bearing artifact must be created directly in its durable repo-approved path. `.tmp` is disposable debug-only; an oracle, raw capture, manifest, provenance record, expected-values input/output, benchmark, or reproducibility artifact created there is invalid and cannot be promoted/restored or close an evidence ID.
```

Add this historical disposition before the P4-C2 table:

```markdown
The historical QA oracle-gate and Recovery reports are retained without rewrite. Their target non-access/blocker chronology remains historical fact, but their classification of `.tmp\...\p1b-identity-oracle-output.json` as an accepted raw oracle and their recovery/promotion path are superseded by the Owner invariant. That JSON is invalid-lifecycle debug capture and supplies no expected value. P4-C2 proceeds through a new clean source-only authoring lane and durable-from-creation seal.
```

Replace the six P4-C2 evidence rows with:

```markdown
| `E4-P4C2-ORACLE1` | clean-context source-only bundle created directly at `reports/QA/child04-p4c2/oracle/<oracle_id>/`: accepted source basis, exactly 21 positive + 11 negative rows, complete schema/provenance, valid bundle digest/seal, zero target writes, zero forbidden observations, zero `.tmp` evidence lifecycle violations | pending |
| `E4-P4C2-TARGET1` | existing QA continuation verifies the oracle seal/source basis, completes the canonical full build, runs one fresh normal built analyzer, and records row-level `21/21`, negative `11/11`, exact fields, access/compatibility separation, zero diagnostics/orphans/Child 05-derived claims in a durable run bundle | pending |
| `E4-P4C2-BOUNDARY1` | durable target pre/post manifests prove unchanged HEAD, source hashes, and pre-existing non-`.anvien` worktree state; only normal analyzer-owned `.anvien` output changes; no target report/fixture/probe/debug artifact | pending |
| `E4-P4C2-REVIEW1` | independent later Supervisor PASS over source basis, lane independence, seal chronology/identity, row comparison, and target boundary | pending |
| `E4-P4C2-DETECT1` | Anvien-side change/boundary check after Supervisor PASS and before evidence commit; exact durable evidence + five-ledger manifest only | pending |
| `E4-P4C2-COMMIT1` | isolated Anvien-side durable oracle/QA evidence and ledger commit; no target or `.tmp` artifact committed or required for reproducibility | pending |
```

### 14.4 Benchmark ledger — replace the three P4-C2 rows

```markdown
| P4-C2 | bounded target direct exports against sealed source-only oracle | correct / expected | 0/21 accepted baseline | pending sealed oracle + fresh target run | pending | 21/21 | pending | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1` |
| P4-C2 | owner-qualified target negative controls remaining non-exported | correct / expected; false positives | pending source-only oracle | pending | pending | 11/11; false positives 0 | pending | `E4-P4C2-ORACLE1`, `E4-P4C2-TARGET1` |
| P4-C2 | target access/export conflations and Child 05-derived claims | records | pending source-only oracle | pending | pending | 0 / 0 | pending | `E4-P4C2-TARGET1`, `E4-P4C2-BOUNDARY1` |
```

Add this non-benchmarkable note:

```markdown
- Oracle row count, provenance completeness, seal identity, durable-path compliance, and `.tmp` absence are evidence-integrity gates rather than product benchmarks; they close only through `E4-P4C2-ORACLE1` and later Supervisor review.
```

### 14.5 Actual-status ledger — replace target touch mode and P4-C2 next action, then append refresh row

Replace the target `Relationship / Impact Evidence` / `Phase Touch Map` meaning with:

```markdown
| `E:\cheapapp.org` source/worktree | real bounded P4-C2 source input and later analyzer target | Oracle Authoring: inspect-only for the three hash-pinned source files plus read-only HEAD/tracked-status metadata, with no `.anvien` observation; QA after seal: inspect source hashes and target-local normal analyzer output | preserve-only for all target writes except normal analyzer-owned `.anvien` operational output during QA | no copied source, target report, fixture, probe, debug artifact, or evidence-bearing `.tmp` path |
```

Replace the P4-C2 Next Phase decision with:

```markdown
| P4-C2 | target boundary is known; historical `.tmp` oracle/recovery method is invalid; no durable sealed row oracle exists yet | keep P4-C2 open; Main first routes a clean `working-rules + Data-Integrity` Oracle Authoring lane to derive and seal 21+11 rows directly under `reports/QA/child04-p4c2/oracle/...`; only then resume existing QA; Child 05 remains locked |
```

Append this refresh-log row (using Planner’s actual repo basis at edit time):

```markdown
| R11 | 2026-08-21 | P4-C commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877`; Owner evidence-lifecycle clarification; architecture verdict `READY_FOR_PLANNER` | P4-C2 oracle authorship/validation separation and durable artifact lifecycle | P4-C2 remains `open`; historical `.tmp` capture `accepted/lost -> invalid-lifecycle debug`; QA/Recovery reports retained historical; source-only Oracle Authoring authorized before analyzer/QA output | `reports/Architect/rp_architect_260821_103812_by_gpt-5_p4c2_evidence_lifecycle.md` | route clean Oracle Authoring; do not resume analyzer QA until seal; do not open Supervisor/Child 05 |
```

Add these P4-C2 implementation-gate rows without marking them complete:

```markdown
- [ ] `E4-P4C2-ORACLE1` durable source-only 21+11 bundle and seal are verified; no forbidden input, target write, or `.tmp` evidence lifecycle violation occurred.
- [ ] Existing QA resumed only after seal identity routing and completed `E4-P4C2-TARGET1` plus `E4-P4C2-BOUNDARY1` against the same source basis.
```

## 15. What remains deliberately unclaimed

- No 21 expected semantic tuples were authored by this architecture lane.
- No target basis was verified by this lane because target access was explicitly forbidden.
- No build, analyzer, QA comparison, Supervisor review, detect, stage, commit, or Child 05 action occurred.
- This report authorizes and constrains the next lanes; it does not close any P4-C2 evidence ID.

## 16. Handoff

Send to: Main, then Planner synchronization lane.

Planner action: patch exactly the roadmap and four Child 04 ledgers using section 14; do not edit this report or the two historical reports. After the five-document authority is synchronized, Main routes the clean Oracle Authoring lane described above.

Residual Owner question: none.

Final verdict: `READY_FOR_PLANNER`
