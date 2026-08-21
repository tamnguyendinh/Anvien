# QA Report — Child 04 P4-C2 Oracle Gate

Scope: P4-C2 real-target validation preflight, oracle-before-target gate only

Role: QA / real-target validation-only; no code fix, no acceptance verdict

Created: 2026-08-21 09:12:13 +07:00

Anvien repo basis: `e32a412b289453a530bc71b93320ef2b97b3a97a`, branch `master`

Target: `E:\cheapapp.org` — **not accessed** because the oracle gate did not pass

No-fix boundary: no production/test/plan/ledger/target edit; no detect/stage/commit/push/reset/checkout

## Final State

`BLOCKED`

The mandatory independent oracle cannot be locked to the required 21 exact entries from the immutable sources currently retained under `E:\Anvien`. The accepted evidence fixes the three files, their source hashes, and all 21 declaration start lines, but it does not retain the exported name and complete semantic fields for every entry. The accepted raw TypeScript oracle referenced by the reports was under `.tmp\cheapapp-graph-root-cause-restart`; that debug tree is no longer present. Per the oracle-before-target stop condition, QA did not read, stat, hash, run Git in, analyze, or write `E:\cheapapp.org`.

This is a validation blocker, not a product defect verdict and not a Supervisor acceptance decision.

## Authority And Integrity Basis

The following authority was read in full before the gate decision:

- `E:\Anvien\AGENTS.md`.
- `.agents\skills\working-rules\SKILL.md`.
- `.agents\skills\qa\SKILL.md`.
- `.agents\skills\Data-Integrity\SKILL.md`, used only for conservation/parity/fail-closed framing; its generic DB/sync/commit instructions do not override the explicit P4-C2 boundary.
- Graph Accuracy roadmap and all four Child 04 ledgers.
- `docs\contracts\graph-accuracy-contract.md`.
- problem-origin report, bounded causal synthesis, final bounded Supervisor report, P1-B/P2-A investigation reports, and their independent Supervisor PASS reports.
- P4-C Supervisor report and P4-C implementation/commit authority.

Key immutable identities:

| Artifact | Bytes / LF | SHA-256 |
| --- | ---: | --- |
| Child 04 roadmap | `27,109 / 191` | `FFD76CF50666FB24300FADAF24401F15C1D04D45154799D0FDFE69B48DC4CFE4` |
| Child 04 plan | `39,001 / 384` | `CD0E928F7258FBDDAA477B4F7E8CDB1991BFDD7AA6E0B486448707B652CA9523` |
| Child 04 evidence | `38,056 / 290` | `AACD9C12832C8966536CFB061742EE9714E9E80A83EAB9C4ACBCAFF98D3684BE` |
| Child 04 benchmark | `9,658 / 74` | `B2A6274E8C219CF096520193382367A4322A610EC7B8312C42A82FD4785D8592` |
| Child 04 actual status | `27,130 / 194` | `AE831B58B49EE9A5E5410296F67014905092B2502EA55379667C8FE56B1FB89F` |
| Graph Accuracy contract | `9,104 / 100` | `68CB65EF964E6D3D7BB8697BD786AE1451DADB1B36D10CC38B5F9CA3839F2592` |
| problem-origin report | `20,904 / 572` | `AE3AB5AF0BBD19084EA717AA4432CF1EFA0D3D786E62B7102D3001F57A8D54BC` |
| P1-B investigation | `4,862 / 56` | `E7EF2FC4E7881FB791E4E6CD2A73204FCCDA88A3F975EF0EC7061F694A77C170` |
| P2-A first-divergence investigation | `8,252 / 128` | `C47EDCDA4B00481CF44593141C939D0D53693922D498B37664E3229D8E4333B2` |
| P2-A exact root-cause trace (21 start lines/source hashes) | `14,474 / 200` | `9F10136F0163EB9AAA286AFC8B292860A89BDD1152B685B069D51821402767E1` |
| P1-B Supervisor PASS | `5,006 / 59` | `C034D15205F2D660A7A3B8D184D7436B5443D0C1943EBB3D889A5379503AF0D0` |
| P2-A Supervisor PASS | `6,942 / 64` | `341E834CA54970130E7F7A8B3F3965006C4478DD820577499CB6C91DA21056BF` |
| causal synthesis | `13,971 / 133` | `063A7AE0920F61A30E2BF92FB915D82C0CDFC5BE807B3646A52D5D8406303405` |
| final bounded Supervisor PASS | `7,973 / 65` | `C7543DE0947B8C7F87DAC72978AA5EF11AB2DF73B2B9B6D8613D3C721E914022` |
| P4-C Coder report | `16,037 / 223` | canonical `1944C72E51FCCFBAF5BB77BDEC319EDEE11716B0292A2790AF623187EEAC1B3F`; raw `668DF0ECF76F1EA0287659C89C767CDDA2BC15ECEF9FBE3E3C66697F5F23A7B0` |
| P4-C Supervisor PASS | `15,261 / 101` | canonical `AA94C417A02BE371168D8ADCD5B3F4FECF1941F382518EFF9214EC0FABB93CFF`; raw `A2253137537509ECBC1A1100B1A62FE7A53ADCBCBCE3500D48FDF525CA50235F` |

The P4-C Supervisor canonical identity was independently recomputed by replacing exactly one declared 64-hex value in the identity line with 64 ASCII zeroes. It is not byte drift.

## Recoverable 21-Slot Provenance

Accepted reports pin the target basis at historical HEAD `a869876ab6262dacde6cd5d432d099a91852a646` and the following source hashes:

| File | Accepted source SHA-256 | Positive slots |
| --- | --- | ---: |
| `modules/email/server/operations/email-operations-observability.ts` | `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C` | `17` |
| `modules/release-distribution/server/release-distribution-publication-state.ts` | `B53F966F82AEDAC58EC15A892B9B71BA0A5716638703105BB35B126CB398D474` | `3` |
| `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts` | `44467B1D6F1778AE4728B4941136321F7831D8FD67EBE2C93A430A3D132F90AC` | `1` |

The exact accepted start-line denominator is:

| Slot | File | Start line | Name evidence retained in accepted E:\Anvien sources | Gate state |
| ---: | --- | ---: | --- | --- |
| 1 | email observability | 35 | not retained | incomplete |
| 2 | email observability | 48 | not retained | incomplete |
| 3 | email observability | 55 | not retained | incomplete |
| 4 | email observability | 61 | not retained | incomplete |
| 5 | email observability | 68 | not retained | incomplete |
| 6 | email observability | 76 | not retained | incomplete |
| 7 | email observability | 82 | not retained | incomplete |
| 8 | email observability | 91 | not retained | incomplete |
| 9 | email observability | 99 | not retained | incomplete |
| 10 | email observability | 101 | not retained | incomplete |
| 11 | email observability | 109 | not retained | incomplete |
| 12 | email observability | 118 | not retained | incomplete |
| 13 | email observability | 150 | not retained | incomplete |
| 14 | email observability | 158 | not retained | incomplete |
| 15 | email observability | 254 | `buildEmailOperationsReport` | partial only |
| 16 | email observability | 438 | not retained | incomplete |
| 17 | email observability | 497 | `readEmailOperationsReport` | partial only |
| 18 | release publication state | 1 | type-alias name not retained | incomplete |
| 19 | release publication state | 7 | `canPublishRelease` | partial only |
| 20 | release publication state | 13 | `canRollbackRelease` | partial only |
| 21 | admin commercial-config read | 10 | `readAdminCommercialConfig` | partial only |

This table is deliberately **not** promoted to an oracle. It lacks 16 exported names and lacks a complete authoritative tuple of `kind/name/meaning/typeOnly/access/compatibility` for all 21 slots. Filling those values from a later target read would make the oracle circular and violate the gate.

## Missing Data Required To Unblock

One immutable Anvien-side source is required that carries all 21 rows before target access. Each row must contain:

- repository-relative file and stable source identity/hash;
- source site/range or accepted line identity;
- export `kind`;
- exact exported `name` and local name;
- exact `meaning`;
- exact `typeOnly`;
- expected access value, independently separated from module export;
- expected compatibility value (`isExported` or the bounded reader equivalent).

An acceptable recovery would be the hash-pinned machine-readable output of the previously accepted independent TypeScript oracle, restored under `E:\Anvien`, or another immutable accepted Anvien-side artifact with equivalent row-level data. Reading/copying current target source first and then calling the result an oracle is not acceptable.

## Negative-Control Definition Available From Authority

The following controls are independently named by accepted Anvien-side evidence and are absent from the accepted 21 export-line denominator. If the positive oracle is restored, these controls can be locked with expected `Export` cardinality `0` and definition compatibility `isExported=false`:

| File | Control | Site / owner evidence | Expected direct-export result |
| --- | --- | --- | --- |
| email observability | `boundedReadLimit` | function at line `190` | no Export fact; compatibility false |
| email observability | `time` | line `207`, owner `parseTime` | no Export fact; compatibility false |
| email observability | `time` | line `214`, reducer callback in `latestIso` | no Export fact; compatibility false |
| email observability | `now` | line `262`, owner `buildEmailOperationsReport` | no Export fact; compatibility false |
| email observability | `now` | line `501`, owner `readEmailOperationsReport` | no Export fact; compatibility false |
| email observability | `messageRows` | line `504`, owner `readEmailOperationsReport` | no Export fact; compatibility false |
| email observability | `attemptRows` | line `505`, owner `readEmailOperationsReport` | no Export fact; compatibility false |
| email observability | `eventRows` | line `506`, owner `readEmailOperationsReport` | no Export fact; compatibility false |
| email observability | `providerEventRows` | line `507`, owner `readEmailOperationsReport` | no Export fact; compatibility false |
| email observability | `readinessRows` | line `508`, owner `readEmailOperationsReport` | no Export fact; compatibility false |
| email observability | `suppressionRows` | line `509`, owner `readEmailOperationsReport` | no Export fact; compatibility false |

Control semantics: matching is by exact file + source occurrence, never name alone. A same-name local must not be conflated with a module-level export. These controls also require zero terminal/module-resolution/public-API state, but that absence cannot be tested until the positive oracle gate is restored and the normal target run is authorized.

## Exhaustion Evidence

QA searched only inside `E:\Anvien` and did not use target source:

- all living problem-origin, causal, Supervisor, Investigation, QA, roadmap, contract, and Child 04 ledger artifacts;
- retained report/QA JSON captures and repository files;
- current `.tmp` presence (`.tmp\cheapapp-graph-root-cause-restart` is absent);
- Git path history and 768 unreachable blobs, read-only, for the referenced oracle names/output and known target symbols;
- retained graph copies under `E:\Anvien` (only the Anvien self-graph exists; no accepted target graph copy is retained).

The search recovered the 21 line identities and source hashes but no complete row-level oracle. No file was restored, generated, staged, or committed during this search.

## Target And Runtime Gates

| Gate | Result | Evidence |
| --- | --- | --- |
| Oracle-before-target | `BLOCKED` | 21 slots known; complete row tuples unavailable |
| Target pre-state | `not captured by rule` | target access prohibited until oracle PASS |
| Normal built analyzer command/runtime identity | `not run` | downstream of blocked oracle gate |
| Fresh target graph counts | `not run` | downstream of blocked oracle gate |
| 21/21 exact comparison | `blocked` | complete oracle absent |
| Graph JSON/Ladybug/reader parity | `blocked` | target run prohibited |
| orphan/forbidden-state counts | `blocked` | target run prohibited |
| Target post-state/contamination | `not captured by rule` | no target access occurred; no target-side artifact was created |

No scanner remediation was attempted. The known eight omitted paths remain out of scope. No terminal resolution, export-table/barrel traversal, alias/cycle/ambiguity, package public API, ambient/external, UI/Playwright, transport/cache, or Child 05 claim was opened.

## Anvien Worktree Boundary At Report Creation

- HEAD: `e32a412b289453a530bc71b93320ef2b97b3a97a`.
- Branch: `master`, ahead of `origin/master` by `30`.
- Tracked worktree diff: empty.
- Index/staged diff: empty.
- Protected untracked provenance preserved unchanged:
  - `reports/Investigation/rp_main_260821_0631_orchestration_rotation_handoff.md`;
  - `reports/Investigation/rp_main_260821_0721_orchestration_rotation_handoff.md`.
- Concurrent Main provenance observed and not modified/staged:
  - `reports/Investigation/rp_main_260821_0902_orchestration_rotation_handoff.md`.
- QA-created artifact: this report only.
- Debug artifact created by this lane: none.

## Coverage Verdict

- Pass: authority read, repository boundary, canonical/raw P4-C report identity, target non-access, negative-control candidate definition.
- Fail: none classified as a product/runtime defect.
- Blocked: complete 21-entry oracle; therefore every target/runtime/persistence/read check.
- 100% of declared P4-C2 validation scope: no; blocked at the mandatory first data gate.

## Handoff

Next owner: Main for self-verification of this blocker and artifact/worktree reality. If Main can supply or restore an immutable accepted 21-row Anvien-side oracle, open a fresh validation continuation; do not open Supervisor P4-C2 acceptance or Child 05 from this blocked handoff.

## Final Decision

`BLOCKED` for declared P4-C2 scope only.

Handoff to Main - oracle-before-target gate is blocked by the missing immutable 21-row semantic oracle; target was not accessed.
