# Child 04 P4-C2 real-target QA retry

Status: **READY_FOR_SUPERVISOR**

Đây là QA finding/handoff, không phải verdict PASS/REJECT của Supervisor. Active slice là Child 04 P4-C2; P4-C implementation vẫn đóng tại `c99c4070b66e7a96be8c9fa2721a0335a1f94877`. Child 05, terminal resolution, alias/barrel traversal và package public API không được mở.

## Kết luận

Operational retry đã loại bỏ blocker Git trust bằng process-local environment, chạy đúng một target analyze thành công, và hoàn tất toàn bộ comparison `21 + 11` cùng persistence/read boundary.

Kết quả semantic:

- Direct Export facts: hiện diện chính xác `21/21`; kind/name/meaning/typeOnly/range/access/relation đều đúng.
- Positive rows tổng thể: **`6/21` PASS**.
- Negative controls: **`11/11` PASS**.
- Finding duy nhất trên positive rows: **15 exported type aliases có `isExported=false` thay vì `true`**, và affected `FileContext` reader cũng trả `exported=false` cho đúng 15 definitions đó.
- Sáu exported functions PASS đầy đủ.
- Graph JSON ↔ Ladybug persistence parity PASS hoàn toàn; Ladybug bảo toàn chính xác cả giá trị compatibility sai, vì vậy parity không chứng minh semantic correctness.

Rows bị ảnh hưởng: `P001`–`P014` và `P018`.

Mỗi row trên chỉ fail hai field:

```text
compatibility.isExported: expected=true, actual=false
reader.exported:          expected=true, actual=false
```

Không có defect khác trong bounded oracle. Evidence đã đầy đủ để một independent P4-C2 Supervisor lane review finding này.

## Authority và closed-gate references

- Sealed oracle: `E:\Anvien\reports\QA\child04-p4c2\oracle\p4c2-oracle-v1-a869876ab626-260821_110849+0700`.
- Oracle ID: `p4c2-oracle-v1-a869876ab626-260821_110849+0700`.
- Bundle digest: `7749AB14E02FBF61CF7F81B3A7638888F8AD414064F9F969CFED06EB11355439`.
- Counts: `21` positive / `11` negative.
- Prior durable BLOCKED run: `E:\Anvien\reports\QA\child04-p4c2\runs\p4c2-target-validation-260821_113405+0700`.
- Prior full build/runtime evidence remains closed and referenced, not copied: canonical build exit `0`, Anvien `1.2.8`, invoked binary SHA-256 `23D647FC27BDDE841EC3F94316A6BFA7ADA2F24E87BF15D236FB86743BD6E8C6`.
- No authority, seal, full build, or broad manifest gate was rerun in this retry.

## Process-local Git trust correction — PASS

One PowerShell process set exactly:

```text
GIT_CONFIG_COUNT=1
GIT_CONFIG_KEY_0=safe.directory
GIT_CONFIG_VALUE_0=E:/cheapapp.org
```

Within that same process, preflight and the single analyzer invocation ran. The three variables were removed before process exit and are absent in the final QA process.

Git configuration preservation:

- System, global, XDG-global and repository config candidate identities were captured before and after.
- Files checked: `4`.
- Byte/hash differences: `0`.
- No global, system or repository config write occurred.

## Retry preflight — PASS

- Target: `E:\cheapapp.org`.
- HEAD: `a869876ab6262dacde6cd5d432d099a91852a646`.
- Branch: `master`.
- Tracked status: exact seven pre-existing modifications.
- Canonical tracked-status SHA-256: `FCB5AD9155C029FFA6B3D80B8AADD1EE40EDC3B408EA401992E8FEA048E4C5E1`.
- Required source hashes: `3/3` exact.
- Preflight matches both sealed basis and prior run post-state.

## Exact analyzer retry — PASS

- Invocation count in retry: exactly `1`.
- Command: `anvien analyze E:\cheapapp.org --force`.
- Runtime: prior freshly built Anvien `1.2.8`, exact identity referenced above.
- Exit: `0`.
- Duration: `77.37s`.
- Scanned/parsed/failed: `1,359/887/0`.
- Fresh graph: `94,422` nodes / `125,299` relationships.
- Fresh output: `E:\cheapapp.org\.anvien\graph.json` and read-only Ladybug database `E:\cheapapp.org\.anvien\lbug`.
- No retry, `--skip-git`, custom analyzer, stale output substitution, or subsequent Anvien analyzer/query command occurred.

Fresh artifact identities:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `.anvien/graph.json` | 432,028,037 | `BAF4A44E9212BA17D4A07290D7CC6EBF6817B9EEBCC96644D21347BA535F141F` |
| `.anvien/lbug` | 150,351,872 | `B388FCB33B22C1C4307D8BADB542041C53483CB702121B381572AAE7F4CE844D` |
| `.anvien/meta.json` | 259 | `D0AC8C46DB03363E7908198B2336B67B04C72822BEB1A827D6DBDA70DBFFC951` |
| `.anvien/settings.json` | 31 | `6C831C0AC49625D09EE54E76042612B92F47ABC47E4A90D1310F0EA45E824B3C` |

## Complete 21 positive-row result

The durable `comparison-result.json` contains every expected/actual field for every row. The compact result is:

| Rows | Declaration | Result | Exact failure |
|---|---|---:|---|
| `P001`–`P014` | `TypeAlias` | FAIL (`0/14`) | `compatibility.isExported=false`; `FileContext.reader.exported=false` |
| `P015`–`P017` | `Function` | PASS (`3/3`) | none |
| `P018` | `TypeAlias` | FAIL (`0/1`) | `compatibility.isExported=false`; `FileContext.reader.exported=false` |
| `P019`–`P021` | `Function` | PASS (`3/3`) | none |

For every `P001`–`P021` row, these checks PASS independently of the compatibility finding:

- definition occurrence identity, declaration kind and selection range;
- one Export node and one File→Export relation in both Graph JSON and Ladybug;
- local/exported name exact;
- semantic export kind `direct_declaration` exact. Concrete Graph projection is `kind=direct` plus `siteKind=export_declaration`; both raw values and normalized semantic value are recorded per row;
- meaning (`type` or `value`) and `typeOnly` exact;
- source file hash, statement range and selection range exact;
- `access` absent on definition and Export node, kept separate from compatibility;
- `localDefinitionNodeId` exact and endpoint existing;
- export diagnostic count `0`;
- Child 05 derived state empty.

Thus the finding is not a missing Export fact. It is a compatibility projection/consumer defect limited to direct exported `TypeAlias` definitions.

## Complete 11 negative-control result — PASS

`N001`–`N011` all PASS:

- exact definition occurrence found once;
- Export fact count `0`;
- File→Export relation count `0` in Graph JSON and Ladybug;
- `access` absent;
- compatibility `isExported=false`;
- FileContext reader `exported=false`;
- Child 05 derived state empty.

The two same-name `time` controls, two same-name `now` controls and all bounded local bindings do not false-positive.

## Graph JSON / Ladybug persistence parity — PASS

- Target Export records: Graph JSON `21`, Ladybug `21`.
- Export fields compared per record: `28`.
- Total Export field comparisons: `588`.
- Field differences: `0`.
- Export IDs missing Graph→Ladybug / Ladybug→Graph: `0/0`.
- Definition compatibility comparisons: `32`; differences `0`.
- File→Export relationship records: `21`; endpoint/type differences `0`.
- Ladybug was opened through the production native read runner in read-only mode.

Important separation: all 15 incorrect type-alias `isExported=false` values are identical in Graph JSON and Ladybug. Persistence parity passes while semantic compatibility correctness fails.

## Affected reader boundary

The production `FileContext` builder consumed the fresh Graph JSON records directly:

| File | Reader opened | Actual exported-symbol count | Oracle-required direct exports |
|---|---:|---:|---:|
| `modules/email/server/operations/email-operations-observability.ts` | yes | 3 | 17 |
| `modules/release-distribution/server/release-distribution-publication-state.ts` | yes | 2 | 3 |
| `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts` | yes | 1 | 1 |

All six functions appear exported. All 15 type aliases appear non-exported. This exactly mirrors the `isExported` compatibility property and establishes a real affected-reader failure rather than a report-only discrepancy.

## Integrity and scope-boundary checks — PASS

- Graph duplicate node IDs: `0`.
- Missing relationship endpoints: `0`.
- Orphan `localDefinitionNodeId` references: `0`.
- Export diagnostics: `0`.
- Forbidden terminal/resolved-target/public-API property state: `0`.
- Global Export inventory: `3,539`; bounded target direct exports: exactly `21`.
- Scanner omissions, general unresolved source sites and Child 05 resolution are outside this comparison and were not repaired or claimed.

## Target and config post-state — PASS

- HEAD/branch/tracked-status identity remain equal to preflight.
- All seven pre-existing tracked modifications remain preserved.
- Required source files remain `3/3` byte/hash exact.
- Four Git config identities remain unchanged.
- All four fresh analyzer artifacts are unchanged after the read-only comparison.
- Process-local trust variables are absent after the analyzer process.
- No target source/config/report/fixture/probe/debug write occurred; only normal analyzer-owned `.anvien` output changed during the successful analyze.

## Validation asset and execution provenance

- Validation source: `validate_p4c2.go`, born directly in this retry run.
- It streams the 432 MB Graph JSON, matches exact oracle source sites, opens Ladybug read-only, compares 28 Export fields per record, validates relationship endpoints, and runs the production `FileContext` affected reader over the bounded subgraph.
- No Anvien analyzer/query command is invoked by the validator.
- Early validator attempts were non-evidence-producing compile/binder failures: one missing Go brace, `lbugruntime.Row` concrete type assertions, and one generic `CONTAINS` relation-table assumption. The final bounded query uses an unlabeled relationship with persisted `r.type='CONTAINS'`; final `comparison-result.json` is `FAIL_COMPLETE` and comparison-complete.
- `go run` reports a wrapper exit `1` when the validator intentionally exits `2` for a complete semantic mismatch; the structured result's `status=FAIL_COMPLETE`, `comparisonComplete=true`, and per-row data are authoritative.
- The lane-created `.go-tmp` compiler directory was empty and removed by exact literal, non-recursive delete after resolved-boundary verification. No evidence artifact was removed.

## Anvien repository boundary

- Current HEAD/branch at handoff preparation: `e32a412b289453a530bc71b93320ef2b97b3a97a` / `master`.
- Tracked boundary remains exactly five unstaged Child 04 living documents.
- Index remains empty.
- No production/test/golden/plan/ledger edit, stage, commit, push, reset, checkout or detect-changes occurred.
- Prior run, sealed bundle and every pre-existing untracked provenance artifact remain preserved.

## Handoff

- State: **READY_FOR_SUPERVISOR**.
- Evidence targets: `E4-P4C2-TARGET1` and `E4-P4C2-BOUNDARY1` are complete with a bounded semantic finding.
- Active lane: P4-C2 QA is complete; no Supervisor task has been opened by QA.
- Next owner: **Main** for self-verification, then Main alone may open exactly one independent P4-C2 Supervisor lane.
- Completion condition for Supervisor review: independently verify the immutable oracle, exact 15-row type-alias compatibility/reader failure, clean 21-record persistence parity, negative controls and target/config preservation.
- Exact next action: Main reads this run's manifest/report and `comparison-result.json`, self-verifies the finding, then decides whether to open the single permitted P4-C2 Supervisor task. Do not open Child 05.
