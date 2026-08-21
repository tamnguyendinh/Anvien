# Planner Handoff — Child 04 P4-C2 Durable Oracle Correction

Created: `2026-08-21 10:40:59 +07:00`

Role: Planner — P4-C2 only

Verdict: `READY_FOR_MAIN`

## Outcome

P4-C2 remains the sole open slice and can proceed immediately. No slice, phase, evidence ID, recovery task, audit loop, Main verification loop, or intermediate commit was added.

The existing pipeline is now unambiguous:

1. A clean-context Oracle Authoring lane reads only the three hash-pinned target source files, writes the `21+11` oracle bundle directly to its durable QA path, and seals it before any analyzer/QA-output observation.
2. Existing QA consumes the sealed bundle read-only, runs the canonical full build and one normal target analyze, compares actual values, and continues through the existing Supervisor/detect/commit/closure evidence IDs.

`E4-P4C2-ORACLE1` remains pending. This Planner lane did not author an oracle, access the target, perform QA, or claim `21/21`.

## Exact Rule Contradiction Corrected

| Authority / report | Exact statement or method | Correct disposition |
| --- | --- | --- |
| `AGENTS.md:20`; `working-rules/SKILL.md:124-126` | `.tmp` is debug-only; temporary scripts/output are not official evidence | Governing invariant |
| QA blocker `rp_qa_260821_091213...md:19,110,182` | calls the deleted `.tmp` capture an accepted raw oracle and permits restore/supply as unblock | Historical report retained; method superseded |
| Recovery `rp_investigation_260821_0932...md:172-176,204` | makes a backup/recovery of the `.tmp` capture a minimum input/continuation route | Historical report retained; route removed |

`.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle-output.json` is an `invalid-lifecycle debug capture`, not an accepted or lost oracle. It cannot close an evidence ID and cannot become evidence by restore, promotion, copy, rename, or later hashing.

The correct gate is:

```text
source-only Oracle Authoring
  -> durable bundle created at approved path and sealed
  -> existing QA may observe analyzer output and compare read-only expectations
```

It is not “oracle before any target read.” P4-C2 already authorizes the target source as inspect-only/preserve-only; the source-only read is the independent ground-truth derivation step.

## Durable Authority

Architect candidate:

- Path: `reports/Architect/rp_architect_260821_103812_by_gpt-5_p4c2_evidence_lifecycle.md`
- Verified identity: `42,936` bytes / `553` LF / SHA-256 `1E7EEB6DD83F05384BF43ED9216C042C535DAC5E776B56032A17E3D4288BDEEA`
- Verdict: `READY_FOR_PLANNER`, not Supervisor acceptance

It fixes the exact schema, paths, non-circular digest, target boundary, lane ownership, and stop conditions used by the five living-document correction.

## Exact Durable Bundle Contract

Oracle directory:

```text
reports/QA/child04-p4c2/oracle/<oracle_id>/
```

Required files, created directly there from their first byte:

- `oracle.schema.json`
- `source-basis.json`
- `expected-values.json`
- `provenance.json`
- `authoring-report.md`
- `seal.json` written last

QA run directory after seal:

```text
reports/QA/child04-p4c2/runs/<run_id>/
```

All oracle, raw stream, normalized output, manifest, provenance, benchmark, expected-value input/output, and reproducibility material is born durable. No evidence-bearing artifact may be staged through `.tmp`.

## Exact `21+11` Schema Boundary

`expected-values.json` requires:

```text
schemaVersion = "p4c2.oracle.v1"
oracleId
evidenceId = "E4-P4C2-ORACLE1"
targetBasisRef = "source-basis.json"
positiveCount = 21
negativeCount = 11
positiveRows[21]
negativeControls[11]
```

Each positive row `P001..P021` must include normalized source path, uppercase source SHA-256, exact one-based range and selection range, lexical owner, declaration kind, exact local name, exported name, export kind, deterministic meanings from `value|type|namespace`, explicit `typeOnly`, export-fact count `1`, File→Export relation count `1`, independent access state/value, compatibility `isExported=true`, export diagnostic count `0`, empty Child 05-derived state, and a source-semantic derivation. Implementation-generated graph IDs are forbidden oracle fields.

Accepted positive anchors:

| Rows | Source path | SHA-256 | Start lines |
| --- | --- | --- | --- |
| `P001..P017` | `modules/email/server/operations/email-operations-observability.ts` | `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C` | `35,48,55,61,68,76,82,91,99,101,109,118,150,158,254,438,497` |
| `P018..P020` | `modules/release-distribution/server/release-distribution-publication-state.ts` | `B53F966F82AEDAC58EC15A892B9B71BA0A5716638703105BB35B126CB398D474` | `1,7,13` |
| `P021` | `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts` | `44467B1D6F1778AE4728B4941136321F7831D8FD67EBE2C93A430A3D132F90AC` | `10` |

Negative rows `N001..N011` use the same exact occurrence identity and require Export fact/relation counts `0`, compatibility `isExported=false`, preserved access state, and empty Child 05-derived state:

| Row | Line | Name | Owner |
| --- | ---: | --- | --- |
| `N001` | 190 | `boundedReadLimit` | source-confirmed containing owner |
| `N002` | 207 | `time` | `parseTime` |
| `N003` | 214 | `time` | reducer callback in `latestIso` |
| `N004` | 262 | `now` | `buildEmailOperationsReport` |
| `N005` | 501 | `now` | `readEmailOperationsReport` |
| `N006..N011` | 504–509 | `messageRows`, `attemptRows`, `eventRows`, `providerEventRows`, `readinessRows`, `suppressionRows` | `readEmailOperationsReport` |

Rows match by `(source.path, fileSha256, exact occurrence range, lexicalOwner, localName)`, never name alone.

## Seal Identity

`seal.json` excludes itself from the bundle digest. For the other five files in ordinal filename order:

```text
<filename>\0<byte_count_decimal>\0<SHA256_UPPER_HEX>\n
```

Then:

```text
bundleDigest = SHA256(
  UTF8("p4c2-oracle-v1\n")
  + all five sorted file-identity lines
)
```

QA recomputes the bundle digest plus ordinary `seal.json` SHA-256/byte count before any build/analyzer action. QA never edits expected values.

## Next Visible Lane

Owner: clean-context source-only Oracle Authoring

Skills: `working-rules`, `Data-Integrity`

Allowed target access: read-only target HEAD/branch/tracked-status metadata and exactly the three hash-pinned source files. The pre-existing target worktree is preserve-only.

Forbidden: target `.anvien`; Anvien implementation/tests/goldens; analyzer, build, QA, or historical `.tmp` oracle output; target writes; copied source; evidence-bearing `.tmp` artifacts; Child 05 state.

Stop before seal on source hash/anchor mismatch, incomplete or guessed row, forbidden observation, copied source, target write, evidence outside the durable bundle, or seal failure. Return the mismatch to Main; do not run analyzer and do not create another plan/audit loop.

Deliverable back to Main: sealed oracle path, `oracleId`, `bundleDigest`, `seal.json` identity, target HEAD/source hashes, counts `21+11`, and `SEALED`. Main then routes existing QA.

## Exact Edit Manifest

The five living-document correction boundary is exactly:

1. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
2. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-plan.md`
3. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-evidence.md`
4. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-benchmark.md`
5. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-04-typescript-export-semantics/2026-07-28-04-typescript-export-semantics-actual-status.md`

The only new Planner artifact is:

- `reports/Planner/rp_planner_260821_104059_by_gpt-5_p4c2_durable_oracle_gate_correction.md`

No production, test, golden, contract, existing report, QA asset, target, or historical handoff was modified.

## Planner Boundary

- No target access, QA, oracle construction, build, validation, Supervisor acceptance, detect, stage, commit, push, reset, or checkout.
- Passed P4-A/P4-B/P4-B1/P4-C gates remain unchanged.
- `E4-P4C2-ORACLE1` is pending; no `21/21` claim.
- Next action is executable now: Main routes Oracle Authoring.

Final verdict: `READY_FOR_MAIN`
