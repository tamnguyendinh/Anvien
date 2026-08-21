# Child 04 P4-C2 post-repair real-target QA

Status: **READY_FOR_SUPERVISOR**

Đây là QA validation/handoff, không phải verdict nghiệm thu. Active slice là Child 04 P4-C2; Child 05 và mọi terminal/module-resolution/public-API claim vẫn khóa.

## Kết luận

Fresh post-repair validation PASS toàn bộ bounded contract:

- positive oracle: `21/21`;
- negative controls: `11/11`;
- 15 exported `TypeAlias` và sáu exported `Function` đều có compatibility `isExported=true`;
- affected `FileContext` reader trả đúng `17/3/1` exported symbols trên ba bounded files;
- Graph JSON ↔ Ladybug record parity: `588` Export-field comparisons, `0` differences;
- duplicate node IDs, orphan endpoints, orphan local-definition references, export diagnostics và forbidden Child 05 state đều `0`;
- target source/worktree, seven pre-existing modifications và Git config được bảo toàn; chỉ normal analyzer-owned `.anvien` output thay đổi bởi fresh analyze.

QA không tự gán PASS/REJECT acceptance. Evidence đã đủ để Main self-verify rồi mở đúng một independent P4-C2 Supervisor lane.

## Sealed oracle và repair authority

- Oracle bundle: `E:\Anvien\reports\QA\child04-p4c2\oracle\p4c2-oracle-v1-a869876ab626-260821_110849+0700`.
- Oracle ID: `p4c2-oracle-v1-a869876ab626-260821_110849+0700`.
- Bundle digest recomputed/declared: `7749AB14E02FBF61CF7F81B3A7638888F8AD414064F9F969CFED06EB11355439` / same.
- `seal.json`: SHA-256 `00FFA78CAB1B584FB9290EEF8578CBF07B52A86351E9327DA60C1BE39956FE4F`.
- Counts: `21` positive / `11` negative.
- Oracle and source-basis files were read-only; no expected value was changed.

Coder handoff:

- `reports/Coder/rp_coder_260821_131129_by_gpt-5_child04_p4c2_typealias_compatibility_repair_ready_for_qa.md`.
- Identity: `14,218` bytes / `215` LF.
- Raw SHA-256: `34E3B872403621D644CC6C1B1F3756D2365B7BCCCFFDF917F957288C3E355A57`.
- Canonical self-reference-safe SHA-256: `2B3A9B04801DE8F76D79017172EA8E251B4A2975F1772E0A7BD1E278DE1997F6`.

Exact repair candidate verified before target access:

| Path | Bytes | LF | SHA-256 |
|---|---:|---:|---|
| `internal/resolution/emit.go` | 26,772 | 815 | `B1B8F22AEFE94F72483CD284B7E267FE3897412C04D581AFAB7F9FE30F0B2867` |
| `internal/resolution/p4c_export_projection_test.go` | 9,247 | 203 | `F575F16D20D8F95C347E142BCBBEE96E18DA84D23E1433E00B6C2545132E0162` |

Pre-target Anvien boundary:

- HEAD/branch: `310502a88849fe75f86a45a987ba21490d19dbe2` / `master`.
- Tracked unstaged paths: exact five living Child 04 documents plus the two repair paths.
- Index: empty.

## Reused canonical build identity

Coder's canonical full build gate remained closed and was not repeated. Reuse was allowed only after exact candidate/runtime verification:

- CLI version: `1.2.8`.
- Repo-local and globally invoked runtime: `71,372,288` bytes each.
- Both runtime SHA-256: `C9BE636BA375F77B77168666FE904914D91BB0BC57A723C4557B5277B3C146E4`.
- Local/global identities are exact.
- Both runtime timestamps are `817.926s` newer than the latest repair candidate timestamp.

The first freshness expression returned false because PowerShell converted an ISO `Z` string to local wall-clock time before comparing it with a UTC `DateTime`. `runtime-freshness-correction.json` records the fail-closed correction using explicit-zero-offset `DateTimeOffset.UtcTicks`; the underlying captured timestamps show no drift. No target access occurred before this corrected gate passed.

Referenced Coder build result: canonical `npm run full-build` exit `0`, CLI `1.2.8`, final self analyze `1,917/736/0`, graph `114,546/157,361`. QA did not reopen Coder tests or rebuild merely to repeat this passed gate.

## Process-local Git trust and target pre-state

One PowerShell process set exactly:

```text
GIT_CONFIG_COUNT=1
GIT_CONFIG_KEY_0=safe.directory
GIT_CONFIG_VALUE_0=E:/cheapapp.org
```

The same process performed Git preflight, the single target analyze, and immediate post-analyze capture. These variables were cleared before process exit.

Pre-state PASS:

- target: `E:\cheapapp.org`;
- HEAD/branch: `a869876ab6262dacde6cd5d432d099a91852a646` / `master`;
- tracked status: exact seven pre-existing modifications;
- canonical tracked-status SHA-256: `FCB5AD9155C029FFA6B3D80B8AADD1EE40EDC3B408EA401992E8FEA048E4C5E1`;
- index: empty;
- sealed source identities: `3/3` exact;
- all seven modified-file byte/hash identities captured;
- system/global/XDG/repository Git config candidates: four identities captured;
- `.anvien` pre-state identities captured.

No global, system or repository Git config write occurred.

## Exactly one fresh target analyze — PASS

- Command: `anvien analyze E:\cheapapp.org --force`.
- Invocation count: exactly `1`.
- Runtime: built Anvien `1.2.8`, SHA-256 `C9BE636BA375F77B77168666FE904914D91BB0BC57A723C4557B5277B3C146E4`.
- Exit: `0`.
- Duration: `89.123s`.
- Scanned/parsed/failed: `1,359/887/0`.
- Fresh graph: `94,422` nodes / `125,299` relationships.
- No second retry, custom analyzer, `--skip-git`, stale substitution, detect-changes or later Anvien graph/query command occurred.

Fresh analyzer artifacts:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `.anvien/graph.json` | 432,026,884 | `688216B304A3F832BC76AFDA0E8C9D14E50CB643E26FECC5B3AA04C26A4910D7` |
| `.anvien/lbug` | 150,355,968 | `F3E8C833D4F0262052DEB3F9D042E964A09B80A15D4BC68E1B496BD7735BCF69` |
| `.anvien/meta.json` | 259 | `68E034285AE87246A7AED23488A5211C778530D35EB848CC79F8E2244B0C0358` |
| `.anvien/settings.json` | 31 | `6C831C0AC49625D09EE54E76042612B92F47ABC47E4A90D1310F0EA45E824B3C` |

## Positive oracle — 21/21 PASS

Complete per-row expected/actual records are in `comparison-result.json`; compact rows are in `positive-row-summary.tsv`.

| Rows | Declaration kind | Meaning / typeOnly | Result |
|---|---|---|---:|
| `P001`–`P014` | `TypeAlias` | `type` / `true` | `14/14` PASS |
| `P015`–`P017` | `Function` | `value` / `false` | `3/3` PASS |
| `P018` | `TypeAlias` | `type` / `true` | `1/1` PASS |
| `P019`–`P021` | `Function` | `value` / `false` | `3/3` PASS |

Every positive row independently proves:

- exact definition occurrence/declaration kind and selection range;
- exact exported/local name;
- semantic kind `direct_declaration` from the concrete `kind=direct` plus `siteKind=export_declaration` projection;
- exact meaning and `typeOnly`;
- one Export fact and one File→Export relationship in both Graph JSON and Ladybug;
- exact file hash, source range and selection range;
- `access` absent on definition and Export node;
- compatibility `isExported=true`, evaluated separately from access;
- affected-reader symbol found and `exported=true`;
- exact `localDefinitionNodeId` with an existing endpoint;
- export diagnostic count `0`;
- Child 05-derived state empty.

The rejected old-byte TypeAlias state is repaired: all 15 TypeAlias rows now have compatibility and reader values `true` without adding runtime-value or terminal-resolution claims.

## Negative controls — 11/11 PASS

For `N001`–`N011`:

- exact bounded definition occurrence found once;
- Export fact count `0`;
- File→Export relationship count `0` in Graph JSON and Ladybug;
- compatibility `isExported=false`;
- access absent;
- FileContext `exported=false`;
- Child 05 state empty.

The two same-name `time` controls, two same-name `now` controls and bounded local binding leaves do not false-positive.

## Persistence, reader and integrity boundary — PASS

Graph JSON ↔ Ladybug:

- bounded Export records: `21/21`;
- Export fields compared per record: `28`;
- total field comparisons: `588`;
- field differences: `0`;
- IDs missing either direction: `0/0`;
- definition compatibility comparisons: `32`, differences `0`;
- File→Export relationship rows: `21`, differences `0`;
- Ladybug opened through the production native reader in read-only mode.

Affected production FileContext reader:

| File | Reader opened | Exported symbols |
|---|---:|---:|
| `modules/email/server/operations/email-operations-observability.ts` | yes | 17 |
| `modules/release-distribution/server/release-distribution-publication-state.ts` | yes | 3 |
| `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts` | yes | 1 |

Integrity/scope counts:

- global Export inventory: `3,539`;
- bounded target direct Export count: `21`;
- duplicate node IDs: `0`;
- missing relationship endpoints: `0`;
- orphan local-definition references: `0`;
- export diagnostics: `0`;
- forbidden terminal/resolved-target/public-API property state: `0`.

Scanner omissions, general unresolved source sites and Child 05 resolution remain outside scope and were not repaired or claimed.

## Post-state and cleanup — PASS

After analyze and after read-only comparison:

- target HEAD/branch/status/index remained exact;
- all three sealed source identities remained exact;
- all seven pre-existing modified-file identities remained exact;
- all four Git config identities remained exact;
- the four fresh `.anvien` artifacts remained unchanged by comparison;
- process-local Git trust variables were absent;
- no target source/config/report/fixture/probe/debug file was written;
- only the normal successful analyze changed target-local `.anvien` output.

The reusable non-production validator was read-only and identity-locked in `comparison-validator-provenance.json`; no historical comparison output was read or reused. All new raw output/result/provenance was born in this run. The compiler dependency path under repo-local `.tmp` was tooling only, never an evidence input or output. The lane-created `.go-tmp` under this durable run was empty and removed by exact literal, non-recursive delete; no evidence artifact was removed.

## Handoff

- State: **READY_FOR_SUPERVISOR**.
- Evidence targets: `E4-P4C2-TARGET1` and `E4-P4C2-BOUNDARY1` are complete on post-repair bytes.
- Active lane: P4-C2 QA complete; QA opened no Supervisor task.
- Next owner: **Main** for self-verification, then exactly one independent P4-C2 Supervisor lane.
- Completion condition for Supervisor: independently verify candidate/runtime identity, immutable `21+11`, TypeAlias compatibility repair, FileContext values, record-level persistence parity, integrity zeros and target/config preservation.
- Exact next action: Main reads this report, `comparison-result.json`, pre/post captures and manifest; then Main alone decides whether to open the permitted P4-C2 Supervisor. Do not open Child 05.
