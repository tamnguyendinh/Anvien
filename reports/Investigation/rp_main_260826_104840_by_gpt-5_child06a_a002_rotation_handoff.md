# Main Rotation Handoff — Child 06A A002 Corrected Measurement

## Metadata

- Written: `2026-08-26 10:48:40 +07:00`
- Workspace: `E:\Anvien`
- Outgoing Main: `01a03941-882b-7e70-93cf-2fdf6ee549c9`
- Successor Main: `01a03c31-2456-79f2-bab6-e71d42337052`
- Independent Governance Rule Guard: `01a03943-55dd-7dd2-bde2-98fc11b9c763`
- Rotation reason: mandatory 120-minute Main-only rotation was overdue; Guard classified the lapse `CRITICAL — VERIFIED VIOLATION`
- Current slice: `P2-A / A002 / B1-P1A-OP001 resolution / B2-P2A-A001-D001 resolve_calls`
- Git HEAD: `cc420e7ad719d90dc4b2d9991be0249e8d648daa`
- Staged path count: `0`
- Campaign state at transfer: `A00X_SCRIPT_BUILD_COMPLETE / CHEAPAPP_CORRECTED_MEASUREMENT_ACTIVE / RESTAURANT_QUEUED / SUPERVISOR_LOCKED`

## Mandatory Successor Rule Seal

Before any campaign action, the successor Main MUST read fully through EOF and enforce the raw, unmodified contents of:

1. `E:\Anvien\AGENTS.md`
2. `E:\Anvien\.agents\skills\working-rules\SKILL.md`
3. `E:\Anvien\.agents\skills\orchestration\SKILL.md`
4. `E:\Anvien\.agents\skills\governance-rule-guard\SKILL.md`
5. `E:\Anvien\.agents\skills\planner\SKILL.md`
6. `E:\Anvien\.agents\skills\supervisor\SKILL.md`

The successor MUST NOT use this report, a summary, memory, or compacted context as a substitute for any of the three iron rule files `AGENTS.md`, `working-rules/SKILL.md`, and `orchestration/SKILL.md`. Absolute full-raw adherence is mandatory throughout the successor session and in every later Main handoff.

After the rule seal, read fully through EOF:

- `E:\Anvien\docs\plans\2026-07-26-anvien-graph-accuracy-multi-plan\2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy\plan-rules.md`
- all four standard Child 06A ledgers in that directory;
- `E:\Anvien\reports\system-architect\rp_system-architect_260825_184751_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`;
- `E:\Anvien\reports\coder\rp_coder_260825_193200_by_gpt-5_child06a_a002_paused_handoff.md`;
- this handoff.

Re-anchoring restores context only. Do not rerun A001, A002 source/build/test, overlay verification, reusable-script validation, or another gate already recorded PASS unless its source/runtime/boundary is invalidated.

The Governance Guard is independent. Do not command, retarget, request ACK from, wait for, or functionally use it. This official record supplies the successor ID after creation so the Guard can independently change its monitoring target.

## Latest Owner Authority

- Owner explicitly ordered `tiếp tục`; status questions or answers are not a pause.
- Owner requires future A00x benchmark candidates to use reusable `scripts/build-a00x-benchmark.ps1` with explicit overlay/output/hash inputs, canonical vendor/native/CGO/offline flags, and binary/DLL/provenance under repository-local `.tmp`.
- When the script/native contract is unchanged, later A00x attempts call it without repeating general build-interface discovery or audit.
- If a genuinely different A00x feature/input makes the script fail, classify that exact failure, refresh the current attempt, and change only the narrow evidenced script contract. Do not restart a broad audit of the canonical scripts.
- Owner rejected unnecessary forks. The required Main rotation uses a new visible local Main task, not `fork_thread`, and the outgoing Main terminates after authority transfer.

## Verified Campaign Cursor

- `P0-A`: complete.
- `P1-A`: complete.
- `P2-A`: open and unchecked.
- Active parent: unchecked `B1-P1A-OP001 resolution`.
- Active child: unchecked `B2-P2A-A001-D001 resolve_calls`.
- D001 no-KEEP streak: `0`.
- D002-D017: queued, unchecked, unopened, and forbidden in A002.
- A001: `KEEP`, accepted, promoted, committed, and preserved.
- A002 production/source/build/test: valid and uncommitted.
- Reusable A00x build script and frozen A002 packet: validated once under `E2-P2A-A002SCRIPTBUILD1`.
- No valid A002 target performance packet, per-attempt Supervisor result, disposition, detect, stage, or commit exists.
- Restaurant Manager, Supervisor, P3, and Child 07 remain locked by sequence.

## Accepted A001 Baselines — Keep Separate

| Target | D001 `resolve_calls` | Parent `resolution` | Analyzer | Process | Denominator |
|---|---:|---:|---:|---:|---|
| `E:\cheapapp.org` | `25.045225300 s` | `184.481061700 s` | `274.474620900 s` | `279.105934600 s` | `calls=27890; files=887` |
| `E:\Restaurant_manager` | `40.769294200 s` | `136.436879300 s` | `215.972455200 s` | `218.680628900 s` | `calls=86030; files=1234` |

Never average or combine the repositories. Restaurant retains exactly one exclusion: `electron/renderer/src/api/userApi.ts`.

## A002 Candidate And Validation

- Exact production delta: `+56/-4` across only:
  - `internal/graphhealth/diagnostics.go`
  - `internal/resolution/emit.go`
  - `internal/resolution/outcome.go`
- New test only: `internal/graphhealth/diagnostics_test.go`.
- Canonical full build: PASS before tests.
- Five focused tests and `internal/graphhealth`: PASS.
- `internal/resolution`: truthfully FAIL only on the identical current-HEAD-overlay preserve-only `TestProofBasedCallAccessGoldenCorpus`; package is not labeled PASS.
- Coder report status: `CODER_A002_A00X_SCRIPT_BUILD_READY_FOR_MAIN`.
- Candidate remains uncommitted; staged set is empty.

Current scoped Git status at handoff:

```text
 M internal/graphhealth/diagnostics.go
 M internal/resolution/emit.go
 M internal/resolution/outcome.go
 M reports/coder/rp_coder_260825_193200_by_gpt-5_child06a_a002_paused_handoff.md
?? internal/graphhealth/diagnostics_test.go
?? scripts/build-a00x-benchmark.ps1
```

All other dirty/user/campaign work remains protected. No broad reset, checkout, stash, cleanup, or casual staging is allowed.

## Reusable Script And Frozen Packet

- Script: `E:\Anvien\scripts\build-a00x-benchmark.ps1`
- Script SHA-256: `ADA407C7496FCEA988276F03BAD5001ED139A4AEC9A16B9C32947DA440814EC5`
- Packet root: `E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build`
- Executable: `anvien-a002-benchmark.exe`, version `1.2.8`, `73,816,576` bytes, SHA-256 `0F8B8244ABC80339A73E3A29F38D32F55A7A6A65281A64C291785ACF8F4A241E`
- DLL SHA-256: `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`
- Provenance SHA-256: `673FCCFEA17C48640C0AC62243479FE5EFD71E172CB1D4136D046C806E5855F1`
- Provenance verifies schema `1`, attempt `A002`, HEAD, `2/2` overlay mappings, `3/3` A002 sources, `3/3` native inputs, exact flags/environment/output, and exit `0`.
- Validation root: `E:\Anvien\.tmp\child06a_a002_a00x_script_validation`.

Do not build again for the active Cheapapp measurement. Do not invoke or audit either canonical build script.

## Active And Queued Visible Lanes

### Active Cheapapp measurement

- Task: `01a033a5-4700-7272-a28d-de3a71f58135`
- Current turn: `01a03c2e-64d2-7f31-ad22-5f4488977209`
- State at handoff: `inProgress`; no assistant/tool marker had materialized in the first immediate snapshot.
- Latest packet explicitly resumes the old STOP and supersedes the obsolete direct-build packet.
- Exact goal: one corrected Cheapapp analyze using the frozen executable; no build and no retry.
- New raw root: `E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826`.
- New report: `E:\Anvien\reports\Investigation\rp_child06a_a002_cheapapp_benchmark_frozen_17child.md`.
- Both paths were absent at dispatch.
- Required result: exit `0`, `30/30` operations, ordered `17/17` child rows, D001 denominator `27890/887`, `2675` exclusive intervals, overlap `0`, conservation, nonempty profile, exact workload/graph/DB/output/semantic equivalence, resource fields, and D001/parent/analyzer/process deltas.
- Stop on any identity drift, competitor, schema/denominator/conservation/profile/equivalence failure, or source/build need. Do not retry.

Successor action: monitor actual commands/artifacts continuously. Do not resend while the turn is active. If it returns `MEASUREMENT_INCOMPLETE`, preserve the packet and classify the exact failure; do not launch a second Cheapapp run automatically. If it returns `MEASUREMENT_COMPLETE`, independently verify the report/raw packet before accepting it.

### Queued Restaurant Manager measurement

- Existing task: `01a033a5-4ec3-7e42-8946-0ab9172f6088`.
- Do not dispatch until the corrected Cheapapp packet is valid and recorded in the four ledgers.
- When opened, use the same frozen executable and its own new repo-local raw root; keep `--force --json --progress` and exactly one `--exclude electron/renderer/src/api/userApi.ts`; no build or overlap.

### Idle Coder

- Valid Coder task: `01a039b4-f7b2-75b1-b405-1bf6b42a281a` per latest campaign checkpoint; do not reopen for measurement.
- Historical A002 Coder IDs/reports remain evidence only.
- Wrong fork `01a039b2-...` was archived and must never be used.

### Supervisor

- No A002 Supervisor is open.
- Open one fresh visible Supervisor only after both target packets are valid, independently recorded, and still bind the same candidate/source/runtime identity.

## Exact Transition After Cheapapp

1. Consume and zero-trust verify the Cheapapp completion packet. A lane report is evidence input, not a verdict.
2. Use Planner skill to update benchmark/evidence/actual-status/plan immediately with the valid Cheapapp measurement. Do not check parent/child or change streak/baseline/disposition.
3. Only then dispatch the existing Restaurant Manager task with a complete frozen-binary measurement packet. No build and no concurrent benchmark.
4. Verify and record Restaurant independently; never aggregate targets.
5. After both packets exist, open a fresh visible Supervisor for the exact A002 candidate and affected correctness/equivalence/resource/lifecycle boundary.
6. Main may decide `KEEP` only if D001, parent, and process all decrease on both repositories independently, correctness/resource equivalence passes, and Supervisor returns PASS.
7. On no retainable gain or Supervisor REJECT, follow the exact accepted-state restoration/new-Architect rules; do not send correction directly to Coder.
8. Do not detect, stage, commit, open P3, or open Child 07 early.

## Official Authority Transfer

The successor visible task `01a03c31-2456-79f2-bab6-e71d42337052` is the sole Main from this official transfer. The outgoing Main performs no further functional orchestration and terminates. The successor must continue the active Cheapapp gate without restarting it and must enforce the full raw rule seal above.
