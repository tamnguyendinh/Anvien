# Anvien Graph Accuracy — Main Orchestration Rotation Handoff

Status: `SEALED / OFFICIAL_TRANSFER_READY / P6-D SUPERVISOR-REJECT REPAIR ACTIVE`

Internal createdAt: `2026-08-23 04:57:36 +07:00`

Outgoing Main task: `01a02b25-9dbc-78a1-8bd3-bbba39c10a31`

Warmed successor task: `01a02b71-95fd-7520-a207-65fa67c4e3d7`, host `local`, exact saved project/workspace `E:\Anvien`

## Authority and hard boundary

- This is the durable zero-trust handoff for the next Visible Main. Authority remains with outgoing Main until the exact official follow-up is delivered to the warmed successor.
- The only workspace boundary is `E:\Anvien`. Direct access to `C:\`, any alternate worktree, network/install/package scripts, and every undelegated path is forbidden.
- `E:\cheapapp.org` remains locked and unaccessed. Do not open it until an explicitly authorized, policy-compliant built-current-runtime gate exists.
- P6-D is the sole open slice. Pn-A, Pn-B, Pn-C, and Child 07 remain locked.
- Coder and Supervisor lanes have no deadlines. Only Main rotates.

## Incoming handoff and accepted campaign anchor

Outgoing Main zero-trust read the predecessor handoff, full `AGENTS.md`, `working-rules`, `orchestration`, `planner`, `supervisor`, all four Child 06 ledgers, the coder report, and the independent Supervisor report before governing this transition.

- Predecessor handoff: `reports/Investigation/rp_main_260823_034308_orchestration_rotation_handoff.md`, `11,569` bytes / `174` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `A8953B70F5A6866FEFBB96D972E25662CC069ECDDFBAB2B4B85F2F13CB5D84BC`.
- P6-A commit: `b98131e44932a7bcac17b487ecb2914535927d01`.
- P6-B commit: `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- P6-C1 commit: `8055f0a6860721e26462572e34469e0d708d4a52`.
- P6-C2 commit: `2acd357db32ac0c2bd3c3968a950c773acf3000d`.
- P6-C3 independent PASS report SHA-256: `DBA2111E7739AEB95E3E74CC29A5CA79F122908E34A8D5B75953D55CADA72E0F`.
- P6-C3 isolated commit: `0d9803777880d8ecca09cae90592d4728a231160`, parent `2acd357db32ac0c2bd3c3968a950c773acf3000d`, exact `26` paths / `3,337` insertions / `557` deletions.

## Independent P6-D Supervisor verdict

Visible Supervisor task `01a02b4f-20b7-7643-bb72-bc2b038cc6b6` completed with exact verdict `REJECT`.

Report:

- path `reports/Supervisor/rp_supervisor_260823_044455_by_gpt-5_p6d_graph_health_projection_and_target_proof.md`;
- `19,281` bytes / `196` LF / `0` CR / strict UTF-8 without BOM;
- SHA-256 `D015CBD4847261DD50C9033DFA487657629910B296041744EF0BCAF5E108F2AE`.

Blocking local invariant:

- `validStructuredResolutionAuthority` validates the nested authority internally but compares it with the outer outcome only by `sourceSiteId`, stage, and status family.
- It does not enforce cross-object equality for six groups: `filePath`, `fileHash`, complete range, `siteKind`, `requestedName`, and `requestedMeaning`.
- The independent hostile probe mutated one group at a time and reproduced false `standard_library/non_actionable` acceptance on all `6/6`. Required fail-closed result is `unclassified/review`.
- The probe source was deleted; the then-current candidate remained unchanged.
- Minimum P6-D repair is limited to the six equality groups in `internal/graphhealth/diagnostics.go` plus six durable hostile rows in `internal/graphhealth/p6d_outcome_projection_test.go`.
- Supervisor also observed the same omission in committed P6-C3 `validateResolutionOutcome`. That concern is explicitly routed but is outside current P6-D repair authority. Do not edit committed P6-C3 without new explicit scope authority.

Independent remaining blockers are unchanged: no policy-compliant built current runtime; no native Ladybug/built-reader parity; target remains `0/3` unproved; oracle and target pre/post boundary are unrun. A truthful BLOCKED handoff is not P6-D acceptance.

## Main-owned ledger transition

Using the `planner` skill, Main updated all four current Child 06 ledgers once to reflect the Supervisor REJECT. No checklist was checked, no stage/commit occurred, and no self-referential ledger-only analyze/detect loop was started.

Current ledger state records:

- P6-D `INDEPENDENT SUPERVISOR REJECT / SIX-FIELD NESTED-AUTHORITY IDENTITY REPAIR OPEN`;
- `E6-P6D-REVIEW1 = REJECT`;
- actual-status refresh row `R35`;
- benchmark hostile equality denominator `0/6`, target `6/6` fail-closed;
- built-current-runtime/Ladybug/target/oracle/boundary blockers unchanged;
- P6-D remains unchecked and uncommitted.

## Active P6-D coder lane

- task `01a02af5-88ef-7561-8c2f-35f195a6ef3f`, host `local`, same saved project/worktree;
- current repair turn `01a02b76-7d9f-79f0-aaa0-0152deaf07f7`;
- latest compact cursor at snapshot: `1cd55979-c3fa-4a26-98ec-b42057bbdc6d:144`;
- status active/inProgress with no task error or attention flag;
- no tool marker had appeared for the repair turn at the `04:56:45 +07:00` snapshot.

Main reused the existing coder exactly once and routed only the six-field P6-D equality repair plus six durable hostile rows. A coordination correction reserved the initial four-ledger REJECT sync for Main; Main then confirmed that sync complete and allowed the coder to proceed from current bytes.

Do not STOP, restart, duplicate-prompt, impose a deadline, or open another coder. Resume compact `wait_threads` from cursor `:144`. Monitor actual commands and enforce these boundaries:

- allowed edits: the two P6-D graph-health source/test paths, four current Child 06 ledgers after Main's R35, and the sole coder report;
- no P6-C3 source/test edit;
- `GOPROXY=off`, `GOSUMDB=off`, no network/install/package script;
- no target access;
- no forbidden cache/root reuse, cleanup retry, stale binary, or workaround;
- no stage/commit;
- final disposition remains BLOCKED/REJECT-READY even if the six-field local repair passes.

Coder report before repair turn:

- `reports/coder/rp_coder_260823_032449_by_gpt-5_p6d_graph_health_projection_blocked.md`;
- `15,101` bytes / `243` LF / `0` CR / strict UTF-8 without BOM;
- SHA-256 `2F2E09ED46DE482696949C0BBC4FE71DDD28BFD26AF829135515084083AB2DAA`.

Remeasure report and every candidate identity after the coder declares final bytes.

## Snapshot-bound candidate identity

At `2026-08-23 04:56:45.2276596 +07:00`, before any repair-turn tool marker, the seven-path candidate totals `354,936` bytes. Canonical manifest SHA-256, using ordinal/case-sensitive relative-path sorting and `relative-path|bytes|SHA256` plus a terminal LF, is `F16BE003A262502C46BE8CDAA0375A343911F8DC0C309C68380B7F656B748E60`.

| Path | Bytes | LF | CR | SHA-256 |
|---|---:|---:|---:|---|
| `internal/graphhealth/diagnostics.go` | 26,410 | 822 | 0 | `796ADF6B3D7B0D643F34E114408EAFE9C003B3B2551AAD328287204C649CBA03` |
| `internal/graphhealth/p6d_outcome_projection_test.go` | 18,197 | 405 | 0 | `2222BFC2F769C394A47CA840BBF5B27038C27A746FB908A259F19D88D4CC3CBB` |
| plan ledger | 82,481 | 535 | 0 | `DE68A590A27C4231B1EF89207BACEF6B422A319BD5CE93F10CB862BC2FEF665B` |
| actual-status ledger | 61,842 | 270 | 0 | `D3FDB1AD08513EC222480AC6F977D14B879B83FC34E9BC753CE6E73C8328C94D` |
| evidence ledger | 129,604 | 594 | 0 | `2983B6958B520897BC71F9E1D635ADD90FF3923AB3122D9EA3E86E0F00FFC3F4` |
| benchmark ledger | 21,301 | 102 | 0 | `38A77D8F06AF570D34E628D3A5E781E59B3A866E41293A948A725BB13C155AF5` |
| coder report | 15,101 | 243 | 0 | `2F2E09ED46DE482696949C0BBC4FE71DDD28BFD26AF829135515084083AB2DAA` |

All eight measured candidate/review files were strict UTF-8 without BOM. This identity is a timestamped handoff snapshot, not permission to infer current bytes after the active coder starts editing.

## External validation truth preserved

Before Supervisor review, Main independently verified the then-current candidate:

- forced graph `2,068` scanned / `763` parsed / `0` failed, `117,865` nodes / `164,424` relationships / `20,310` dependency edges;
- broad build exit `1` on preserved offline dependencies, mixed-package fixture, and C fixture;
- affected graph-health build PASS;
- focused P6-D matrix PASS `3 + 2 + 20 + 1 + 1` and full `internal/graphhealth` PASS;
- generic Graph JSON/ResolutionGap/summary parity PASS;
- detect exit `0`, file-layer HIGH `5` tracked files / `198` symbols, summary low/affected `0`, ResolutionGap `95/99`, semantic completeness `117,865/117,865`;
- coder command history contained `153` executions and zero command containing the target name/path.

These green gates are narrower than the hostile six-field defect and are invalidated where the repair changes source/test bytes. Do not reuse them as post-repair acceptance evidence.

## Network deviation and quarantine

The earlier coder began one network-enabled pinned-module build contrary to exact authority. Main STOPped it. It exited `1`, produced no runtime binary, and left invalid downloaded caches quarantined inside the exact coder root. One exact lane-only cleanup attempt was policy-rejected before process creation; no retry, workaround, broad cleanup, or reuse occurred.

At the `04:56:45 +07:00` read-only snapshot:

| Repo-local root | Files | Bytes | Zero-byte files |
|---|---:|---:|---:|
| `.tmp\p6c3-repair` | 1,673 | 76,978,267 | 19 |
| `.tmp\p6c3-supervisor-rereview-01a02a84` | 1,662 | 93,162,861 | 18 |
| `.tmp\p6c3-main-closure-01a02abb` | 1 | 343 | 0 |
| `.tmp\p6d-coder-01a02abb` | 10,053 | 971,554,982 | 84 |
| `.tmp\p6d-main-verification-01a02b25` | 1,512 | 64,175,973 | 19 |
| `.tmp\p6d-supervisor-review-01a02b4f-260823-01` | 1,522 | 66,564,280 | 19 |

Do not use, write, clean, or work around any predecessor blocker root or quarantined downloaded cache. Debug roots are not runtime evidence. Remeasure rather than infer.

## Git boundary

At `2026-08-23 04:56:45.2276596 +07:00`:

- branch `master`;
- HEAD `0d9803777880d8ecca09cae90592d4728a231160`;
- parent `2acd357db32ac0c2bd3c3968a950c773acf3000d`;
- status `59 = 5` tracked + `54` untracked;
- reports `53 = 42` Investigation + `7` Supervisor + `4` coder;
- the remaining untracked path is the P6-D test;
- index empty;
- `git diff --check` exit `0`.

This handoff report adds one protected Investigation report, so the immediate post-seal expectation is `60` status paths and reports `54 = 43 + 7 + 4`, unless the active coder concurrently changes path membership. Do not stage or commit without complete evidence and independent PASS.

## Warmed successor and next rotation

- successor task `01a02b71-95fd-7520-a207-65fa67c4e3d7`, host `local`, same saved project;
- warmed around `2026-08-23 04:47 +07:00` before the required `04:49:30` target;
- compact ACK cursor `ffb5de63-d767-4bb0-958b-c7c5074d39c3:1`;
- correct `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER` response;
- no tool marker before transfer.

After official transfer, the successor's next rotation target is warmup by `2026-08-23 05:49:30 +07:00` and absolute transfer by `2026-08-23 06:04:30 +07:00` if the campaign remains active.

## Exact next ownership

1. Verify this handoff byte-for-byte; reread full `AGENTS.md`, `working-rules`, `orchestration`, `planner`, `supervisor`, all four current Child 06 ledgers, and the independent REJECT report.
2. Re-snapshot Git, candidate/report identities, report split, active coder turn, and repo-local temp roots. Never infer live bytes from this snapshot.
3. Resume compact monitoring of the same coder from cursor `:144`. Do not duplicate prompt, restart, STOP, or assign a deadline.
4. Enforce the exact six-field P6-D repair boundary and six durable hostile rows. Block any P6-C3 edit, target access, network/cache reuse, cleanup workaround, or commit attempt immediately.
5. After coder completion, externally verify source/diff, all six hostile rows, broad/affected build truth, focused/full tests, parity, forced-current graph, file-detail/impact, detect, Git/index, quarantine, report seal, and target-never-accessed provenance.
6. Open exactly one new independent visible Supervisor resubmission only after the repair handoff is sealed. Do not reuse the prior Supervisor task for acceptance.
7. P6-D cannot PASS while built-current-runtime/native Ladybug/built-reader/target/oracle/pre-post evidence is missing. Keep P6-D unchecked, unstaged, and uncommitted on REJECT/BLOCKED evidence.
8. Only Main may planner-finalize/stage/commit after a complete independent PASS. Pn-A/Pn-B/Pn-C and Child 07 remain locked.
9. Warm the next Main by `05:49:30` and officially transfer by `06:04:30 +07:00` if the campaign remains active.

Owner correction remains binding: coder/Supervisor lanes have no deadlines; only Main rotates.
