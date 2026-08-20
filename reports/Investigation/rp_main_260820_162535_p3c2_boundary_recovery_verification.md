# Main Verification — P3-C2 status-boundary recovery

Verdict: `PASS` for the `RECOVERED` boundary-evidence handoff only. This is not `E3-P3C2-REVIEW1`, does not accept P3-C2, and does not open Pn-A.

## Metadata

- Review time: `2026-08-20 16:25:35 +07:00` (`Asia/Bangkok`)
- Reviewer role: Orchestration Main using the zero-trust Supervisor guard
- Repository / HEAD: `E:\Anvien` / `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`
- Sole open slice: Child 03 `P3-C2`
- Claim reviewed: the retained target status serialization is reproducible, its exact bytes encode `7 tracked + 6 untracked`, and the earlier `6 tracked` wording is a non-byte-backed handoff miscount rather than a target delta
- Recovery report: `reports/Investigation/rp_main_260820_161744_p3c2_status_boundary_recovery.md`
- Main did not access `E:\cheapapp.org`; verification used the durable report, both raw task rollouts, current Anvien artifacts, and current Anvien Git state.

## Artifact integrity

| Artifact | Bytes | Lines | SHA-256 | Result |
| --- | ---: | ---: | --- | --- |
| recovery report | `16,512` | `327` | `F44293C45E142742D63D71A400D975D7E48101D7C1A91F69D18739FDE1A27700` | exact |
| current QA report | `30,745` | `378` | `A7CB04165C2CB064CCDEF75991FA75B2BF332B73AC95724FBC5FECA1FB49A41E` | preserved |
| rotation handoff | `10,873` | `203` | `B84C02D7769B6C821266ADD078461EB44FC1A37C3ABBDD3A3F171FE479812C70` | exact |
| `cmd/binding-contract-probe/main.go` | `19,607` | `617` | `C68E520AF4F220AD2E9E4A9B71941EB9B816C3EAFC15D6F7B44212305BB7BBA7` | preserved |
| `cmd/binding-contract-probe/main_test.go` | `10,809` | `306` | `DA3B500FB161411A4D134DDF11935A83ECB1955F08C3087B5C78D4B4073C4F9A` | preserved |
| probe Coder report | `16,942` | `317` | `FC25935961463336A554503F9D5045B3F117BE5DC7ED7E61FEBF88F062927A5F` | preserved |

No accepted source, test, probe, QA, plan, ledger, contract, or prior report byte changed during the recovery lane.

## Independent original-transcript verification

Main read the original QA rollout directly at:

`C:\Users\TAM NGUYEN\.codex\sessions\2026\08\20\rollout-2026-08-20T13-51-46-01a01df0-aaec-76e0-8cdc-6108f793fd7f.jsonl`

At verification it remained `4,157,794` bytes / `1,489` JSONL records / SHA-256 `5C64C5D78B7F6D2BCA3BDC0367BAF05827E80AD5D4B7030C8C5D56A17BABFA2F`.

The corrected capture is call line `243`, output line `244`, turn `01a01df0-ad63-75a0-a3a7-65823754c0af`, and `call_id=call_OiCZ9HIuUHvK1SgIoatknXIh`. Main independently parsed the embedded `statusLines`, preserved their order, serialized them with `[string]::Join([char]10, ...)`, encoded UTF-8 without BOM, and recomputed:

- SHA-256 `FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E`;
- `1,833` bytes;
- `15` lines = `2` branch headers + `13` entries;
- `7` tracked porcelain-v2 `1` rows;
- `6` untracked `?` rows;
- no BOM, CR, or NUL;
- terminal byte `0x67`.

The seven exact tracked paths and six exact untracked paths parsed by Main equal those recorded in the recovery report.

## Independent recovery-rollout verification

Recovery task rollout:

`C:\Users\TAM NGUYEN\.codex\sessions\2026\08\20\rollout-2026-08-20T16-07-22-01a01e6c-ccaf-7b90-827b-15a580e6f323.jsonl`

At verification it was `1,485,211` bytes / `343` records / SHA-256 `640F51D00FB901AEA53404DE739908A0045D5EF16F6078691CADC4B788E1B5AA`.

Main enumerated every tool call in that rollout. Exactly one `exec` call used target workdir `E:\cheapapp.org`: line `247`, timestamp `2026-08-20T09:16:45.895Z`, `call_id=call_JoNLDNxyHNVm3ILj2R8DotHd`. Its output is line `248`, exit `0`. Main independently decoded that output and verified:

- observed pre/post HEAD `a869876ab6262dacde6cd5d432d099a91852a646`;
- pre/post status SHA-256 `FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E`;
- `1,833` bytes, `15` serialized lines, `13` entries, `7` tracked, `6` untracked;
- tracked-diff object `941c3e00be7357de8393d959b91ca93be72e64fb`;
- untracked manifest `886DE642875AA7D58F5F8357054927AADA49CAE65604120AB28AA788F0E81B97`, `6` files / `573,963` bytes;
- `StatusHashMatch`, `HeadMatch`, `TrackedDiffMatch`, `UntrackedManifestMatch`, `CountMatch`, and all five pre/post equality fields are `true`.

No second target-workdir command exists in the recovery task. The lane did not rerun analyze, probe, build, tests, vet, Cypher/query, or any product gate.

## Count discrepancy closure

The sole direct source for `6 tracked` is line `48` of `reports/Investigation/rp_main_260820_141241_orchestration_rotation_handoff.md` (`8,526` bytes; SHA-256 `F2B9275653ECE4992C0D2FB32FA7C3E4C8826052D4CFB668B0194B094F803130`). That handoff was created at `14:12:41 +07:00`, after the corrected original capture at `14:01:51–14:01:53 +07:00`, and contains no six-row status artifact.

The retained hash, original raw rows, and current bounded reproduction all prove `7 tracked + 6 untracked`. Therefore the earlier `6 tracked` phrase is a summary miscount. No evidence supports a different twelve-entry target state, and no omitted path is invented.

## Current Anvien Git boundary

- HEAD / parent remain `656a0445ff3e25b6225b994cdaf7cf1b35eb665c` / `a569b8674fefdaa757cf7fdf63f454caf7925215`.
- Tracked modifications: `0`.
- Staged paths: `0`.
- Recovery added only `reports/Investigation/rp_main_260820_161744_p3c2_status_boundary_recovery.md` to the prior eight-path untracked boundary.
- This Main verification report is the only additional path created by Main after accepting the recovery evidence.
- No target, source, test, probe, QA, plan, ledger, contract, or forbidden-tree path was edited or staged.

## Invariant closure and route

- Affected invariant: reproducible provenance and equality of the target pre/post status boundary.
- Original serialization provenance: closed by direct raw-rollout reconstruction.
- Current equality: closed by the one bounded read-only reproduction recorded in the recovery rollout.
- `6` versus `7` count discrepancy: closed as a handoff miscount by exact original/current rows.
- Residual same-invariant surface: none for boundary recovery.
- This review does not accept P3-C2. The existing QA task must consume the exact recovered recipe/report, update the same QA report to `READY_FOR_SUPERVISOR`, and stop.

## Overall evaluation

The boundary-recovery handoff is accepted as `RECOVERED`. The evidence is current, exact, independently traceable to both raw rollouts, and closes the sole prior `E3-P3C2-BOUNDARY1` provenance blocker without rerunning or changing any accepted product behavior. P3-C2 remains behind the existing QA handoff and a separate visible `E3-P3C2-REVIEW1` Supervisor gate.
