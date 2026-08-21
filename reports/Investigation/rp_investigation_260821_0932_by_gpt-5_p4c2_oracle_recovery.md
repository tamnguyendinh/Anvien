# Child 04 P4-C2 Immutable Oracle Recovery Investigation

Created: 2026-08-21 09:32 +07:00

Role: Evidence Recovery / Data-Integrity investigator

Scope: Anvien-side forensic recovery only

Final state: `BLOCKED_NO_ORACLE`

Next owner: Main, for independent verification of this blocker and artifact boundary

## Decision

No immutable Anvien-side artifact recoverable from the declared surfaces contains the 21 complete P4-C2 rows. The surviving accepted evidence fixes three repository-relative files, three SHA-256 source identities, and 21 unique accepted declaration start lines (`17 + 3 + 1`). It retains only five of the 21 exported names and retains zero complete row tuples with the required `export kind / exported name / local name / meaning / typeOnly / expected access / expected compatibility` fields.

The deleted machine-readable capture `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle-output.json` cannot be recovered from the living repository, refs, reflogs, stash, commits, trees, blobs, unreachable object inventory, repo-local caches, or retained graph files. No target source was accessed, so the missing fields were not inferred or reconstructed circularly.

This is a recovery blocker only. It is not a product verdict, QA target validation, or Supervisor acceptance.

## Hard boundary observed

- All command working directories and explicit filesystem roots were under `E:\Anvien`.
- `E:\cheapapp.org` was not read, resolved, tested for presence, statted, hashed, used as a Git worktree, analyzed, or written.
- No build, fresh graph analyze, Anvien query, target validation, detect-changes, stage, commit, push, reset, or checkout occurred.
- No production, test, golden, plan, ledger, contract, or existing report was edited.
- `anvien --help` was the only Anvien CLI invocation; it was used to comply with the repository rule for command-surface orientation and did not read a graph.
- No internal subagent was used.
- No debug directory was created. `.tmp\p4c2-oracle-recovery` remained absent.

## Authority and immutable basis

The recovery used the repository rule/skill authority, the Graph Accuracy roadmap, all four Child 04 ledgers, the graph-accuracy contract, the P4-C2 blocker report, and the accepted P1-B/P2-A investigations and Supervisor artifacts identified by that blocker.

Key checked identities:

| Artifact | Bytes / LF | SHA-256 |
| --- | ---: | --- |
| P4-C2 oracle-gate blocker | `13,142 / 188` | `CAD7F4300A8B5B2EB0FD9A59808030DCC1ACCDB8B1DB6C5B25784B1D513B175D` |
| P1-B investigation | `4,862 / 56` | `E7EF2FC4E7881FB791E4E6CD2A73204FCCDA88A3F975EF0EC7061F694A77C170` |
| P2-A first-divergence investigation | `8,252 / 128` | `C47EDCDA4B00481CF44593141C939D0D53693922D498B37664E3229D8E4333B2` |
| P2-A exact root-cause trace | `14,474 / 200` | `9F10136F0163EB9AAA286AFC8B292860A89BDD1152B685B069D51821402767E1` |
| P1-B Supervisor PASS | `5,006 / 59` | `C034D15205F2D660A7A3B8D184D7436B5443D0C1943EBB3D889A5379503AF0D0` |
| P2-A Supervisor PASS | `6,942 / 64` | `341E834CA54970130E7F7A8B3F3965006C4478DD820577499CB6C91DA21056BF` |

The Child 04 ledgers were checked at their current identities recorded by the blocker: plan `CD0E928F...CA9523`, evidence `AACD9C12...3684BE`, benchmark `B2A6274E...5D8592`, and actual status `AE831B58...FB89F`. Their P4-C2 oracle rows remain pending and contain no hidden row table.

## Surviving 21-slot basis

| File | Accepted source SHA-256 | Accepted start lines | Slots |
| --- | --- | --- | ---: |
| `modules/email/server/operations/email-operations-observability.ts` | `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C` | `35, 48, 55, 61, 68, 76, 82, 91, 99, 101, 109, 118, 150, 158, 254, 438, 497` | 17 |
| `modules/release-distribution/server/release-distribution-publication-state.ts` | `B53F966F82AEDAC58EC15A892B9B71BA0A5716638703105BB35B126CB398D474` | `1, 7, 13` | 3 |
| `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts` | `44467B1D6F1778AE4728B4941136321F7831D8FD67EBE2C93A430A3D132F90AC` | `10` | 1 |

The 21 file/start-line pairs are unique; duplicate slot count is `0` and missing slot count relative to the accepted denominator is `0`.

The only exact exported names retained in accepted Anvien-side text are:

| File / line | Retained name |
| --- | --- |
| email / 254 | `buildEmailOperationsReport` |
| email / 497 | `readEmailOperationsReport` |
| release / 7 | `canPublishRelease` |
| release / 13 | `canRollbackRelease` |
| admin commercial-config / 10 | `readAdminCommercialConfig` |

Therefore exact exported-name completeness is `5/21`; 16 names are missing. The accepted reports mention that release line 1 is a type alias, but do not retain its name. Converting that prose, declaration labels, or current source into the new P4 export schema would be inference and was rejected.

## Required-field completeness gate

| Required field | Retained immutable coverage | Missing | Gate |
| --- | ---: | ---: | --- |
| repository-relative file identity | `21/21` | `0` | complete |
| accepted source SHA-256 | `21/21` by file | `0` | complete |
| accepted source site (start line) | `21/21` | `0` | complete |
| exact exported name | `5/21` | `16` | incomplete |
| exact local name as an explicit row field | `0/21` | `21` | incomplete |
| exact P4 export kind | `0/21` | `21` | incomplete |
| exact P4 meaning lane | `0/21` | `21` | incomplete |
| exact `typeOnly` value | `0/21` | `21` | incomplete |
| expected access value, separate from export | `0/21` | `21` | incomplete |
| expected compatibility value | `0/21` | `21` | incomplete |
| complete required tuple | `0/21` | `21` | blocked |

The historical observation that all selected graph definitions lacked export/visibility metadata is a baseline defect observation, not the expected P4-C2 access/compatibility oracle. It cannot fill those expected fields.

## Full recovery-surface inventory

### Living E:\Anvien files

- `33,733` living files outside `.git` were inventoried with `rg --files --hidden -uuu -g '!.git/**'`.
- Binary-safe content search used `rg -a -l --hidden -uuu -g '!.git/**'` for the deleted filename marker, the two rarer fixed source hashes, and fixed source identities.
- The living hits were only the current blocker, the current/forced Main handoff, three accepted P1-B/P2-A reports, and the historical investigation evidence ledger. No machine-readable oracle or oracle script survived.
- A recursive filename inventory found `0` living files whose names match `p1b.*oracle` or `oracle.*output`.
- `.tmp` contains `738` files. The historical `.tmp\cheapapp-graph-root-cause-restart` tree and the declared JSON are absent. Searches of `.tmp` produced no deleted-oracle marker or fixed three-file hash pair.
- `.anvien` contains exactly four Anvien self-index files: `graph.json`, `lbug`, `meta.json`, and `settings.json`. Binary-safe searches produced no deleted-oracle marker or fixed three-file hash pair. This is the Anvien self-graph, not a retained accepted target graph.
- Git LFS is absent (`.git/lfs` does not exist).
- One non-Git archive exists: `.tmp\ladybug-native\downloads\liblbug-windows-x86_64-0.19.1.zip`, `8,292,678` bytes, SHA-256 `865E2C8765064BE76D41E4D786DFB0CD3AD0C258DDAF7B522FA3DA7159ECD3EF`. Its four entries have zero oracle/P1-B/target-source identity hits.

### Git refs, history, reflogs, stash, and path history

- `git for-each-ref`: `23` refs.
- `git reflog show --all`: `496` reflog entries.
- `git stash list`: one stash, `stash@{0}: On master: wip p7-g repo-ignore-settings-tests`.
- `git rev-list --objects --all --reflog`: `23,137` reachable object/path rows.
- `git log --all --reflog -- <oracle paths>` and `git log --all --reflog -S<marker/hash>` found only the committed reports/ledger that reference the deleted capture. No oracle path ever appears in reachable path history.

### Complete Git object database

The object database was enumerated with:

```text
git cat-file --batch-all-objects --batch-check='%(objectname) %(objecttype) %(objectsize)'
git cat-file --batch
git fsck --full --unreachable --no-reflogs
```

Every blob payload was streamed and searched, including objects reachable only through refs/reflogs/stash and every unreachable blob:

| Object type | Count | Uncompressed bytes read |
| --- | ---: | ---: |
| blob | `13,313` | `255,103,550` |
| tree | `10,230` | `6,984,864` |
| commit | `958` | `294,486` metadata census |

`git fsck` reported exactly `768` unreachable blobs, `743` unreachable trees, and `107` unreachable commits. Candidate selection was fail-closed around the recovery contract: any qualifying 21-row artifact must carry the fixed file/source identity, so blobs were retained when they contained the exact deleted filename marker, at least two fixed source basenames, at least two fixed source hashes, or the email basename plus P4 semantic-field markers. Seventeen blobs qualified; 12 are reachable Markdown reports/plans and five are unreachable Markdown report versions. No JSON or oracle script qualified.

All tree payloads were also searched for `p1b-identity-oracle`, `oracle-output.json`, and the historical root-directory name. There were `121` historical directory-name hits for `cheapapp-graph-root-cause-restart`, but `0` tree-entry hits for `p1b-identity-oracle` and `0` for `oracle-output.json`. Thus neither the JSON nor its script is present as a named tree entry in stored Git history.

`git count-objects -v` reported `35` garbage entries (`1,303` KiB), all surfaced as orphan `.idx` files with no corresponding `.pack`. Pack indexes do not contain recoverable blob payloads or tree filenames; without the matching pack bytes they cannot supply an oracle and were classified as insufficient metadata, not a recovery source.

## Candidate objects and why they fail

### Closest unreachable marker candidate

| Property | Value |
| --- | --- |
| Git object | `a2d2cbd9dd36a8ccad7554ac1dbf95658496f423` |
| Git type / reachability | blob / unreachable, no path from refs or reflogs |
| Bytes / LF | `4,872 / 56` |
| SHA-256 | `39C62EC90AEA47871BDD85A08C9EC843BE28C2F4295C7BB1AC609C1DD002BE20` |
| Content identity | older Markdown version of the P1-B investigation |
| Oracle marker | one sentence pointing to the deleted JSON path |
| Row payload | aggregate `17 + 3 + 1`; no 21-row names/semantics |

This blob is not the oracle. It is a report referencing the oracle. Its current reachable successor is blob `76f5105f2ab503e80bb8386fe560dfea73f5c251` at `reports/Investigation/20260726_155807_p1b_ts_identity.md`, `4,862` bytes / 56 LF / SHA-256 `E7EF2FC4...7C170`, and it has the same aggregate-only limitation.

### Closest accepted row-slot candidate

Reachable blob `13f746143495065d2929bd0edbb5e0df44d04e17` is `reports/Investigation/20260726_161952_p2a_identity_root_cause.md`, `14,474` bytes / 200 LF / SHA-256 `9F10136F...67E1`. It carries all three source hashes and all 21 start-line slots, but only five exact export names and no P4 row tuples. It is a summary/causal trace, not a machine-readable oracle.

### Remaining four unreachable candidates

| OID | Bytes / LF | SHA-256 | Content | Result |
| --- | ---: | --- | --- | --- |
| `0a002f4c0a4d48ca63ecb0d946ccb9ab4a51d087` | `6,456 / 90` | `C4BB054EE9A1C543BC10170A3F08404FFFCDD5DAAC763BE3E91CCD47C778BC51` | P1-C resolution/module-binding report | unrelated; no export-row oracle |
| `74b4830b82061a543a2eb817bc2d1b4434957541` | `5,194 / 81` | `43D7035AC0BC1381831E42F6AA4751AF26688DBF10E3EA978A57FB8E0FD1B8CE` | P1-E command-projection report | unrelated; no export-row oracle |
| `75c6e4b0f4ccfefb881886e9a7b4bf83f3238059` | `6,943 / 65` | `0A834FAF8A3FC2A8C7BF795E072D8A2B0B32A00EA1A5A0FDE9C0A40922930933` | older P2-A Supervisor report | aggregate 21 only; no row tuples |
| `a573e9a3d5e88c361731fd94137e266772fc09a4` | `14,260 / 183` | `C82220B06BB8E3616353C24A9BEC9AC5F89BE2AC326FEE2EE6BD0D8F61B81760` | P1-D projection-parity report | unrelated; no export-row oracle |

The other 11 reachable candidate blobs map to the accepted P1-B/P2-A/Supervisor reports, historical investigation plan/evidence, and unrelated P1-C/P1-D/P1-E/P2-B reports. Each is Markdown and each is a reference, aggregate, causal summary, or unrelated source-site report rather than the deleted row-level JSON.

## Circular candidates rejected

- The accepted source hashes and start lines are sufficient to identify slots, but not to derive the missing 16 names or semantic tuple values without source bytes.
- Current target source was not opened. Reading it and then labeling extracted values an independent oracle would violate oracle-before-target.
- No Anvien-side copied target source was found as an accepted oracle. Even if a copied source blob had appeared, source-to-schema reconstruction would still be a new inference and would have been rejected unless the artifact itself were previously accepted as the row oracle.
- Current P4 production behavior, tests, golden files, self-graph Export nodes, or generic TypeScript rules cannot supply target-specific expected rows independently. Using them would make validation self-referential.
- Negative controls retained by the blocker do not repair the missing positive denominator.

## Minimum input needed to unblock

Main/Owner must supply one of the following before the existing QA lane can continue:

1. a provenance-bound backup of the original `p1b-identity-oracle-output.json` (and preferably its generating script/capture metadata), plus immutable evidence tying those exact bytes to the previously accepted Anvien-side investigation; if that historical schema lacks any P4 fields, an independently accepted immutable augmentation must supply them; or
2. another immutable pre-target Anvien-side artifact containing all 21 rows with file/source identity, accepted site/range, exact export and local names, export kind, meaning, `typeOnly`, expected access, and expected compatibility.

A newly reconstructed table from the current target, current analyzer output, line-number guesses, language conventions, or current P4 implementation is not sufficient.

## Git and artifact boundary

Initial baseline at 09:23:26 +07:00:

- HEAD `e32a412b289453a530bc71b93320ef2b97b3a97a`, branch `master`.
- tracked worktree diff empty.
- index/staged diff empty.
- four untracked files: the `0631`, `0721`, and `0902` protected Main handoffs plus the P4-C2 blocker report.

During the read-only investigation, Main concurrently created `reports/Investigation/rp_main_260821_0925_orchestration_forced_handoff.md`, `6,325` bytes / 77 LF / SHA-256 `1846F95038E6E6668588C4114790B60665CE59EEDCC40D3CDF5F9747A81CE750`. It was observed and preserved unchanged. This report is the only artifact created by the recovery lane.

Expected final state before handoff:

- HEAD unchanged at `e32a412b289453a530bc71b93320ef2b97b3a97a`.
- tracked diff empty; index/staged diff empty.
- protected untracked set preserved, plus the concurrent `0925` Main handoff and this Investigation report.
- no JSON oracle created because recovery failed.
- no `.tmp\p4c2-oracle-recovery` directory created.

## Handoff

`BLOCKED_NO_ORACLE`

Main must independently verify this report identity and the final Git boundary. Do not call any candidate accepted, do not resume target access, do not open Supervisor or Child 05, and do not duplicate the completed blocked QA lane. If a qualifying immutable oracle is later supplied, Main may route continuation to the existing P4-C2 QA validation task.
