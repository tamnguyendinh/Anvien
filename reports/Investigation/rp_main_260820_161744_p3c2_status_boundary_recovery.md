# P3-C2 status-boundary provenance recovery

Created: `2026-08-20 16:17:44 +07:00` (`Asia/Bangkok`)  
Repository: `E:\Anvien`  
Target: `E:\cheapapp.org`  
Slice: Child 03 `P3-C2` only  
Role: Data-integrity + QA evidence forensics, target read-only  
Next responsible person: Orchestration Main

## Scope and non-acceptance boundary

This report recovers the provenance of target status snapshot SHA-256
`FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E`
and reconciles the `6 tracked` versus `7 tracked` wording discrepancy. It does
not repeat behavior QA, rerun any accepted product gate, review P3-C2 for
acceptance, edit target state, or update the existing QA report/plan/ledgers.

Retained facts such as `E3-P3C2-ORACLE1`, `E3-P3C2-TARGET1`, normal analyze
`1359/887/0`, graph `90,899/121,868`, and probe stdout hash
`8E382A3ADDEE3542505A0B98A8A9A57EC41606133D72E86455A1CF090F0D8F21`
were not rerun.

## Transcript provenance

The exact recipe and its original output were recovered from the completed QA
task rather than reconstructed from the QA report:

- task: `01a01df0-aaec-76e0-8cdc-6108f793fd7f`;
- task title: `Child 03 P3-C2 — Real-target validation`;
- host: `local`;
- turn: `01a01df0-ad63-75a0-a3a7-65823754c0af`;
- raw rollout log:
  `C:\Users\TAM NGUYEN\.codex\sessions\2026\08\20\rollout-2026-08-20T13-51-46-01a01df0-aaec-76e0-8cdc-6108f793fd7f.jsonl`;
- rollout log at forensic read: `4,157,794` bytes / `1,489` JSONL records /
  SHA-256 `5C64C5D78B7F6D2BCA3BDC0367BAF05827E80AD5D4B7030C8C5D56A17BABFA2F`;
- corrected capture call: JSONL line `243`, timestamp
  `2026-08-20T07:01:51.651Z` (`2026-08-20 14:01:51.651 +07:00`),
  `call_id=call_OiCZ9HIuUHvK1SgIoatknXIh`,
  record `ctc_030b5caf7a9a0fe6016a86a65377508191bc87f53e3692f090`;
- exact call-input SHA-256:
  `6F893CD2F17D2510A2CC4F88F053196046F0ED30E40C85F5E169D91FAAB9F644`;
- original output: JSONL line `244`, timestamp
  `2026-08-20T07:01:53.592Z` (`2026-08-20 14:01:53.592 +07:00`),
  record `ctco_01a01df9-ed38-7962-9c6c-1af7ba4b287f`;
- command workdir: `E:\cheapapp.org`.

The app-level `read_thread` summary identified the correct turn but had compacted
away raw shell executions. Lines `243–244` of the source rollout are therefore
the exact output location used for provenance.

### Exact status-capture recipe

The relevant command fragment at line `243` is:

```powershell
$gitArgs = @('-c','safe.directory=E:/cheapapp.org','-c','core.quotepath=false')
function Invoke-GitLines([string[]]$Arguments) {
  $result = @(& git @gitArgs @Arguments)
  if ($LASTEXITCODE -ne 0) {
    throw "git $($Arguments -join ' ') exited $LASTEXITCODE"
  }
  return $result
}
function Get-StringSha([string]$Text) {
  $sha = [System.Security.Cryptography.SHA256]::Create()
  try {
    return [Convert]::ToHexString(
      $sha.ComputeHash([System.Text.Encoding]::UTF8.GetBytes($Text))
    )
  } finally {
    $sha.Dispose()
  }
}
$statusLines = @(
  Invoke-GitLines @('status','--porcelain=v2','--branch','--untracked-files=all')
)
$statusText = [string]::Join([char]10, $statusLines)
$statusHash = Get-StringSha $statusText
```

The effective Git command is:

```text
git -c safe.directory=E:/cheapapp.org -c core.quotepath=false status --porcelain=v2 --branch --untracked-files=all
```

### Exact byte contract

- Byte source: PowerShell-native command output materialized as an ordered
  `[string[]]`, not a raw redirected Git stdout file.
- Ordering: Git output order is preserved; there is no sorting of status rows.
- Serialization: `[string]::Join([char]10, $statusLines)`.
- Newline: exactly one LF byte (`0A`) between adjacent lines.
- Terminal behavior: PowerShell line materialization removes Git's terminal
  record newline; `Join` adds no terminal LF and no NUL.
- Encoding: `System.Text.Encoding.UTF8.GetBytes`; BOM is not included.
- Original serialized length: `1,833` bytes.
- Original line count: `15` = `2` branch headers + `13` status entries.
- CR present: `false`; NUL present: `false`; UTF-8 BOM present: `false`.
- First 16 bytes:
  `23206272616E63682E6F696420613836`.
- Last 16 bytes:
  `2F77616C6C65742D696465612E6A7067`.
- Terminal byte: `67` (`g`, the final byte of `wallet-idea.jpg`).
- SHA-256 recomputed directly from transcript rows:
  `FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E`.

No CRLF, BOM, UTF-16, NUL-delimited, JSON, or sorted-status candidate is part
of the recovered recipe.

## Exact original and current ordered status

The original transcript and the one authorized current reproduction yielded
the following byte-identical ordered lines:

```text
# branch.oid a869876ab6262dacde6cd5d432d099a91852a646
# branch.head master
1 .M N... 100644 100644 100644 39823deb19bfd8feea98ed8b9de2c0e8718aa368 39823deb19bfd8feea98ed8b9de2c0e8718aa368 .dockerignore
1 .M N... 100644 100644 100644 8c9171221716b3d6b933005af172512cb6380f1a 8c9171221716b3d6b933005af172512cb6380f1a reports/QA/2026-06-15-admin-artifact-delete-confirmation/docker/real-chrome-desktop-admin-artifact-delete-confirmation.png
1 .M N... 100644 100644 100644 b9c2d11b08b766a96af2d209d9e84358fc27b402 b9c2d11b08b766a96af2d209d9e84358fc27b402 reports/QA/2026-06-15-admin-artifact-delete-confirmation/docker/real-chrome-desktop-admin-artifact-delete-disabled.png
1 .M N... 100644 100644 100644 fc00c5c45bef149bed36a1b5dfdb19a06cf6acdb fc00c5c45bef149bed36a1b5dfdb19a06cf6acdb reports/QA/2026-06-15-runtime-landing-header-footer-i18n/docker/real-chrome-desktop-apps-empty.png
1 .M N... 100644 100644 100644 8c50fab5da3346d22d20cc007742419681ca8b34 8c50fab5da3346d22d20cc007742419681ca8b34 reports/QA/2026-06-15-runtime-landing-header-footer-i18n/docker/real-chrome-desktop-download-empty.png
1 .M N... 100644 100644 100644 29498fc4c0a48183b16f623f9f53d5487f33d8b8 29498fc4c0a48183b16f623f9f53d5487f33d8b8 reports/QA/2026-06-15-runtime-landing-header-footer-i18n/docker/real-chrome-desktop-landing-en-review.png
1 .M N... 100644 100644 100644 1515c3b7b418fdf8d2108426959cd8c6596298ff 1515c3b7b418fdf8d2108426959cd8c6596298ff reports/QA/2026-06-15-runtime-landing-header-footer-i18n/docker/real-chrome-mobile-landing-en-review.png
? reports/QA/rp_qa_260619_145255_by_gpt5_user_admin_lifecycle.md
? reports/QA/rp_qa_260619_152926_by_gpt5_release_catalog_mutation.md
? reports/problem/images/login-idea-2.jpg
? reports/problem/images/login-idea-3.jpg
? reports/problem/images/login-idea.jpg
? reports/problem/images/wallet-idea.jpg
```

Classification is direct from porcelain-v2 prefixes:

- `7` tracked entries, all `1 .M`: unstaged worktree modifications with no
  staged index modification;
- `6` untracked entries, all `?`;
- `13` status entries total; branch headers are serialized for the hash but are
  not status entries.

## Reconciliation of `6 tracked` versus `7 tracked`

The `6 tracked` statement has one identified source:

- `reports/Investigation/rp_main_260820_141241_orchestration_rotation_handoff.md`
  (`8,526` bytes; SHA-256
  `F2B9275653ECE4992C0D2FB32FA7C3E4C8826052D4CFB668B0194B094F803130`),
  line `48`: “six pre-existing tracked modifications and six pre-existing
  untracked entries”.

That handoff was created at `2026-08-20 14:12:41 +07:00`, after the corrected
raw capture at `14:01:51–14:01:53 +07:00`. The same handoff gives only hash
prefixes and explicitly directs readers to exact lane evidence; it contains no
six-row status artifact.

Direct evidence contradicts only the handoff's tracked-count summary, not the
target boundary:

1. corrected original transcript output contains seven exact tracked rows;
2. its recovered byte recipe yields the retained `FE3573...` hash;
3. the current reproduction yields the same `1,833` bytes and same hash;
4. both original and current states have six identical untracked paths;
5. tracked diff Git object and untracked manifest also remain exact.

Therefore `6 tracked` is a handoff summarization/miscount with no independent
byte provenance. `7 tracked` is the source-of-truth count committed by the
retained status hash. The handoff does not list its six assumed paths, so there
is no direct evidence identifying which of the seven rows its author omitted
mentally; inventing that omission would be unsupported.

## One bounded reproduction batch

Exactly one post-recovery target batch was run at
`2026-08-20 16:16:49 +07:00`, read-only. It implemented the recovered recipe
verbatim and captured target state twice inside that single batch to attest
pre/post equality. No later target command was run.

| Field | Expected | Observed pre | Observed post | Equality |
| --- | --- | --- | --- | --- |
| HEAD | `a869876ab6262dacde6cd5d432d099a91852a646` | exact | exact | `true` |
| status SHA-256 | `FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E` | exact | exact | `true` |
| status bytes | `1,833` | `1,833` | `1,833` | `true` |
| ordered status entries | `13` (`7` tracked + `6` untracked) | exact rows above | exact rows above | `true` |
| tracked diff Git object | `941c3e00be7357de8393d959b91ca93be72e64fb` | exact | exact | `true` |
| untracked manifest SHA-256 | `886DE642875AA7D58F5F8357054927AADA49CAE65604120AB28AA788F0E81B97` | exact | exact | `true` |
| untracked content total | `6` files / `573,963` bytes | exact | exact | `true` |

The tracked-diff object was reproduced with the original command shape:

```powershell
git -c safe.directory=E:/cheapapp.org -c core.quotepath=false diff --no-ext-diff --binary --full-index HEAD -- |
  git hash-object --stdin
```

The untracked manifest was reproduced from
`git ls-files --others --exclude-standard | Sort-Object -Unique` as UTF-8 rows
`<relative-path>|<byte-length>|<SHA-256>\n`. It is `714` bytes, includes a
terminal LF by recipe, and has the expected SHA-256 above.

### Current untracked content rows

| Path | Bytes | SHA-256 |
| --- | ---: | --- |
| `reports/problem/images/login-idea-2.jpg` | `150,157` | `06FD3EA35F17BFB94CA9BF1A3AD1297DC0EBB6BF9E6F247AC2DA9E0634C49368` |
| `reports/problem/images/login-idea-3.jpg` | `260,724` | `24B8E928F2D2EBE517010F0A535D3AA0703C62D9F89F931CADDB477A61F8192D` |
| `reports/problem/images/login-idea.jpg` | `75,770` | `71AF93B0D5A516F7B22CBA75FCF4A479797BAF04001E3E395FBBCAEBE2007944` |
| `reports/problem/images/wallet-idea.jpg` | `68,960` | `30EECD84F8218248B27B9616A5580D5C5F115132FFEEC87E5CFC3E1BE6813B38` |
| `reports/QA/rp_qa_260619_145255_by_gpt5_user_admin_lifecycle.md` | `12,910` | `354A19EC02BF7989F5CC41AFF6E817918AC025A92DB6F0949F81C821980E1725` |
| `reports/QA/rp_qa_260619_152926_by_gpt5_release_catalog_mutation.md` | `5,442` | `D781C90EA939C71348ED3A3CF0DF6FCD7E5916E0829AC81722BEAD66E6FCFBD3` |

## Target pre/post equality and non-actions

Within the sole reproduction batch:

- `PrePostHeadEqual=true`;
- `PrePostStatusBytesEqual=true`;
- `PrePostStatusEntriesEqual=true`;
- `PrePostTrackedDiffEqual=true`;
- `PrePostUntrackedManifestEqual=true`.

No target delta was observed. The batch created no target file, report, temp,
config entry, checkout, index change, or process-local artifact. It did not read
unrelated target content beyond the four required Git/status boundary
constituents and the six already-authorized untracked manifest files.

No `anvien analyze`, probe, build, test, vet, Cypher/query, graph command,
product gate, `git config`, checkout/reset/stash/clean/add/commit/push, or target
repair was run in this recovery lane.

## Anvien Git boundary

### Before report creation

- HEAD: `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`;
- branch: `master`, upstream `origin/master`, ahead/behind `+15/-0`;
- tracked modifications: `0`;
- staged paths: `0`;
- status serialization SHA-256:
  `F5CC9B306B488176EF1E6EACA8C19868BC3B465AC06901D66D7854B7F09BF020`;
- pre-existing untracked paths:
  - `cmd/binding-contract-probe/main.go`;
  - `cmd/binding-contract-probe/main_test.go`;
  - `reports/Investigation/rp_main_260820_141241_orchestration_rotation_handoff.md`;
  - `reports/Investigation/rp_main_260820_144839_p3c2_blocked_handoff_verification.md`;
  - `reports/Investigation/rp_main_260820_152947_p3c2_probe_candidate_verification.md`;
  - `reports/Investigation/rp_main_260820_155906_orchestration_rotation_handoff.md`;
  - `reports/QA/rp_qa_260820_143739_by_gpt-5_p3c2_target_binding_validation.md`;
  - `reports/coder/rp_coder_260820_152054_by_gpt-5_p3c2_binding_contract_probe.md`.

### After report creation

- read-only check time: `2026-08-20 16:20:00 +07:00`;
- HEAD remains `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`;
- branch/upstream/ahead-behind remain `master` / `origin/master` / `+15/-0`;
- tracked modifications remain `0`;
- staged paths remain `0`;
- untracked paths are the same eight pre-existing paths plus exactly:
  `reports/Investigation/rp_main_260820_161744_p3c2_status_boundary_recovery.md`;
- report path occurs exactly once in porcelain-v2 output;
- complete post-status serialization SHA-256:
  `774738567705CC59DD65DED2577790A59FAAA81EFDF96E0F11D73743C7E45D19`;
- after removing exactly the report's one `?` row, the serialization SHA-256 is
  `F5CC9B306B488176EF1E6EACA8C19868BC3B465AC06901D66D7854B7F09BF020`,
  byte-identical to the pre-report boundary;
- direct comparison `OnlyNewEntryIsReport=true`.

Because the report is untracked, closing its final content bytes does not alter
the porcelain status row or the equality above. No pre-existing Anvien path was
edited, staged, removed, or committed by this lane.

## Failures and non-evidence

- The first combined authority read exceeded the display budget; it was not
  treated as a complete read. Every affected file was reread fully in bounded,
  numbered chunks before investigation.
- The first `read_thread` request used `maxOutputCharsPerItem=40000`, above the
  tool maximum; it was rejected without state change. The valid `20000` request
  identified both turns but compacted raw shell output, so it was discovery only.
- Transcript line `234` used the same status recipe and produced the correct
  status hash, but its HEAD expression indexed after a premature string cast and
  returned only `"a"`. It is not HEAD authority. Corrected line `243` first
  materialized `$headLines` and output the full HEAD; line `244` is authoritative.
- Direct `Get-FileHash` on the live rollout log hit a sharing violation. That
  attempt produced no hash and is non-evidence. A read-only `FileStream` opened
  with `FileShare.ReadWrite|Delete` then produced the rollout hash recorded above.
- One local classification expression used PowerShell wildcard
  `-like '? *'`, which matches any one leading character and temporarily labeled
  all `15` lines untracked. That derived count was discarded. Exact prefix checks
  `.StartsWith('1 ')` and `.StartsWith('? ')` over the unchanged transcript output
  yield `7` and `6` respectively.
- Prior wrapper candidate enumeration in the QA task is not the recovered
  recipe and is not used as equality evidence here.
- The sole target reproduction batch completed exit `0`; it had no failed
  target command and required no retry.
- The first final report-verification wrapper had a PowerShell parser error
  (`Missing closing ')' in expression`) before execution because a helper call
  placed a semicolon inside its argument list. It produced no verification
  evidence and no filesystem/process state change. The corrected read-only
  attempt followed.
- The next read-only verification attempt named its local SHA helper `H`, which
  PowerShell resolved to the higher-precedence built-in `h`/`Get-History` alias.
  Report metadata/content still read successfully, but the stripped-status hash
  was null; that attempt is non-evidence for Git equality and changed no state.
  The final wrapper uses the unambiguous function name `Get-TextSha256`.

No temporary directory or debug file was created by this recovery lane.

## Durable artifact integrity

Path:
`E:\Anvien\reports\Investigation\rp_main_260820_161744_p3c2_status_boundary_recovery.md`

The final byte count and SHA-256 are intentionally recorded in the completion
handoff after final bytes are closed; embedding a file's own final hash inside
it would recursively change that hash.

## Verdict

`RECOVERED`

The exact original recipe reproduces `FE3573...` on current target state,
pre/post equality is direct for every required constituent, and the `6` versus
`7` discrepancy is reconciled to a non-byte-backed handoff miscount. This is
boundary-evidence recovery only; it is not a P3-C2 acceptance decision.

Next responsible person: Orchestration Main.
