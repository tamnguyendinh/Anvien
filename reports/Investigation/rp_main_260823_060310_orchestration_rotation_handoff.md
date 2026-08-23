# Anvien Graph Accuracy — Main Orchestration Rotation Handoff

Status: `SEALED / OFFICIAL_TRANSFER_READY / P6-D SIX-FIELD REPAIR CODER-VALIDATED / FULL P6-D BLOCKED`

Internal createdAt: `2026-08-23 06:03:10 +07:00`

Outgoing Main task: `01a02b71-95fd-7520-a207-65fa67c4e3d7`

Warmed successor task: `01a02ba8-0f4d-71c0-8808-ce154facc73f`, host `local`, saved project/workspace `E:\Anvien`

## Authority and boundary

- Only workspace: `E:\Anvien`. Direct `C:\`, alternate worktree, network/install/package scripts, undelegated paths, stale-runtime substitution, and cache workaround remain forbidden.
- `E:\cheapapp.org` remains locked and unaccessed.
- P6-D is the sole open slice. Pn-A/Pn-B/Pn-C and Child 07 remain locked.
- Coder/Supervisor lanes have no deadlines; only Main rotates.
- P6-D remains unchecked, unstaged, and uncommitted. Built-current-runtime, native Ladybug/built-reader, target `0/3`, oracle, and target pre/post boundary remain absent.

## Accepted anchor

- P6-A `b98131e44932a7bcac17b487ecb2914535927d01`.
- P6-B `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- P6-C1 `8055f0a6860721e26462572e34469e0d708d4a52`.
- P6-C2 `2acd357db32ac0c2bd3c3968a950c773acf3000d`.
- P6-C3 PASS report SHA-256 `DBA2111E7739AEB95E3E74CC29A5CA79F122908E34A8D5B75953D55CADA72E0F`; isolated commit `0d9803777880d8ecca09cae90592d4728a231160`, parent `2acd357db32ac0c2bd3c3968a950c773acf3000d`, exact `26` paths / `3,337` insertions / `557` deletions.

## Independent P6-D REJECT anchor

- Report `reports/Supervisor/rp_supervisor_260823_044455_by_gpt-5_p6d_graph_health_projection_and_target_proof.md`.
- `19,281` bytes / `196` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `D015CBD4847261DD50C9033DFA487657629910B296041744EF0BCAF5E108F2AE`.
- Blocking defect was missing nested-authority-to-outer-outcome equality for six groups: `filePath`, `fileHash`, complete range, `siteKind`, `requestedName`, `requestedMeaning`; hostile probe falsely accepted all `6/6`.
- Same omission in committed P6-C3 remains routed/out-of-authority. Do not edit P6-C3 without explicit scope authority.

## Coder repair handoff

- Existing coder task `01a02af5-88ef-7561-8c2f-35f195a6ef3f`, host `local`.
- Repair turn `01a02b76-7d9f-79f0-aaa0-0152deaf07f7` completed at compact cursor `1cd55979-c3fa-4a26-98ec-b42057bbdc6d:145` with no delivered assistant/tool marker, but live worktree measurement proved repaired bytes and a revised durable report.
- Coder report `reports/coder/rp_coder_260823_032449_by_gpt-5_p6d_graph_health_projection_blocked.md`: `20,125` bytes / `262` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `6E8AB2DBC96B43C16FDC7102FCA89D5096CD76D6B597D162C732FED533F72420`.
- Disposition: `BLOCKED_BUILT_CURRENT_RUNTIME / SIX-FIELD_REPAIR_CODER-VALIDATED / REJECT-READY`, never P6-D PASS.
- Source repair adds exact equality for the six groups in `internal/graphhealth/diagnostics.go`; the sole test owner adds exactly six isolated hostile mutations, now within a `26/26` hostile matrix.
- No P6-C3 source/test edit, target access, stage, commit, network retry, cache reuse, cleanup workaround, or stale runtime evidence occurred.

## Exact seven-path repaired candidate

Canonical construction: ordinal/case-sensitive relative-path sort; each row `relative-path|bytes|SHA256`; terminal LF.

Total `371,376` bytes. Canonical SHA-256 `758066DE6FDF37DCCE1EEB22DD4B2B8F577CE8A58CA739445F9AFE6F72B0D3BE`.

| Path | Bytes | SHA-256 |
|---|---:|---|
| `internal/graphhealth/diagnostics.go` | 26,884 | `518BD06FF5F859BD42AD7B3B9FF8E3F9F4E7235720BC5E568E2C4B2DA72BD2AE` |
| `internal/graphhealth/p6d_outcome_projection_test.go` | 19,336 | `6E90E655FFCCF5A5BD3D398D95EB5B6B7F77754CAAF35DAAF4D9609248A4E713` |
| plan ledger | 84,324 | `3493034B66BC3DC7B59C839C4CFD3C174CD8AB3432EAEAD5F509D335C5DC6A29` |
| actual-status ledger | 63,410 | `611DCB131F65E4D124256397D9A971BCA6663516C51B485BB44B8097E181B4FD` |
| evidence ledger | 135,411 | `EAB886BFA5A05E8E3E6FAC557705CD33D9070246040D39B69CB4BBDE982B2C0A` |
| benchmark ledger | 21,886 | `B2C2C4E5A604A17F0AE41DC5440C8426036188D7FA83FA487835B414DF18E2F4` |
| coder report | 20,125 | `6E8AB2DBC96B43C16FDC7102FCA89D5096CD76D6B597D162C732FED533F72420` |

## Main external verification on repaired bytes

- Source/diff inspection confirms all six exact equality groups and six distinct hostile rows; no target literal, P6-C3 edit, or scope expansion.
- Fresh Main E-only root: `.tmp\p6d-main-verification-01a02b71-02`; it is non-candidate and must not be reused as product-runtime evidence.
- Pre-build analyze lock absent/free and relevant build-process count `0`.
- Broad offline `go build -buildvcs=false ./...`: exit `1`, preserved missing-module/mixed-package/C-fixture blockers; no full-build PASS claimed.
- Affected `go build -buildvcs=false ./internal/graphhealth`: exit `0`.
- Focused `go test ./internal/graphhealth -run '^TestP6D' -count=1 -v`: exit `0`, exact groups `3 + 2 + 26 + 1 + 1`.
- Full `go test ./internal/graphhealth -count=1`: exit `0`.
- Forced-current Anvien analyze: `2,070` scanned / `763` parsed / `0` failed / `117,894` nodes / `164,469` relationships / `20,310` dependency edges / `469` unresolved projections.
- Current `diagnostics.go` file-detail: HIGH, `154` symbols / `48` inbound / `70` outbound / `149` local / `187` unresolved / `1` flow / `19` tests.
- Exact repair function impact remains LOW `3` impacted symbols / `1` file / `1` module / `0` processes, while containing-file impact remains CRITICAL `107` impacted symbols / `31` files / `34` direct / `1` flow / `19` tests. Preserve both layers.
- Current detect exits `0`: file-layer HIGH `5` tracked files / `201` changed symbols; summary risk `low`, affected_count `0`; ResolutionGap `96` entities / `100` occurrences (`38` analyzer-gap / `58` non-actionable; `14` builtin / `38` in-repository / `44` standard-library); Resolution Health `0`; semantic completeness `117,894/117,894` for both fields.
- These local gates prove the six-field repair boundary only. They do not prove full P6-D acceptance.

## Git and report boundary

- Snapshot before adding this handoff: branch `master`; HEAD `0d9803777880d8ecca09cae90592d4728a231160`; parent `2acd357db32ac0c2bd3c3968a950c773acf3000d`.
- Status `60 = 5` tracked + `55` untracked; index `0`; `git diff --check` exit `0`.
- Reports `54 = 43` Investigation + `7` Supervisor + `4` coder. This handoff adds one Investigation report, so immediate expected post-seal status is `61` and reports `55 = 44 + 7 + 4`; remeasure rather than infer.
- Candidate remains the exact seven paths above; this Main handoff is protected history outside the candidate.

## Quarantine and temp truth

- Preserved read-only roots at the pre-verification snapshot: `.tmp\p6c3-repair` `1,673 / 76,978,267 / 19`; `.tmp\p6c3-supervisor-rereview-01a02a84` `1,662 / 93,162,861 / 18`; `.tmp\p6c3-main-closure-01a02abb` `1 / 343 / 0`; `.tmp\p6d-coder-01a02abb` `10,053 / 971,554,982 / 84`; `.tmp\p6d-main-verification-01a02b25` `1,512 / 64,175,973 / 19`; `.tmp\p6d-supervisor-review-01a02b4f-260823-01` `1,522 / 66,564,280 / 19` (`files / bytes / zero-byte files`).
- New repair root `.tmp\p6d-identity-repair-01a02b25-01` and new Main root `.tmp\p6d-main-verification-01a02b71-02` are non-candidate, E-only validation roots; cleanup attempts already policy-blocked where reported. Do not reuse, clean, or treat any temp executable/cache as runtime evidence. Successor must remeasure live identities.

## Warmed successor and rotation

- Successor task `01a02ba8-0f4d-71c0-8808-ce154facc73f`, host `local`, same saved project.
- ACK cursor `21a58cdd-d3bb-426a-8c20-23881c04f3d8:1`; correct `UNDERSTOOD / WAITING_FOR_OFFICIAL_TRANSFER`; no tool marker.
- Next successor warmup target: `2026-08-23 06:49:30 +07:00`.
- Next successor absolute transfer deadline: `2026-08-23 07:04:30 +07:00` if campaign remains active.

## Exact next ownership

1. Zero-trust verify this report identity, reread full rules/skills/four ledgers/coder report/REJECT, and remeasure Git/candidate/reports/temp roots.
2. Open exactly one new independent visible Supervisor P6-D resubmission. Do not reuse the prior Supervisor task. Review the exact seven-path repaired candidate and retain full-P6-D blockers.
3. Supervisor must independently inspect source/diff/six hostile rows, rerun fresh graph/file-detail/impact/detect and local build/tests/parity, verify Git/index/quarantine/report seals and target-never-accessed provenance.
4. Expected truthful verdict cannot be full P6-D PASS while built-current-runtime/native Ladybug/built-reader/target/oracle/pre-post gates are absent. Keep P6-D unchecked/unstaged/uncommitted on REJECT/BLOCKED evidence.
5. Do not edit committed P6-C3; route its same-invariant omission separately for Owner authority. Do not access target or use stale/quarantined runtime/cache.
6. Only Main may planner-finalize/stage/commit after a complete independent PASS. Pn-A/Pn-B/Pn-C and Child 07 remain locked.
7. Warm next Main by `06:49:30` and officially transfer by `07:04:30 +07:00` if campaign remains active.

