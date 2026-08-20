# QA Report — Child 03 P3-C2 real-target binding validation

Final decision: `READY_FOR_SUPERVISOR`

## Metadata

- Created: `2026-08-20 14:37:39 +07:00` (`Asia/Bangkok`)
- Role: QA + data-integrity validation-only; no repair and no self-acceptance
- Repository: `E:\Anvien`
- Target: `E:\cheapapp.org`
- Open slice: Child 03 `P3-C2` only
- P3-C commit: `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`
- P3-C parent: `a569b8674fefdaa757cf7fdf63f454caf7925215`
- Next responsible person: Orchestration Main
- Boundary-recovery closure: `2026-08-20`; pre-closure report SHA-256 `A7CB04165C2CB064CCDEF75991FA75B2BF332B73AC95724FBC5FECA1FB49A41E`
- No-fix boundary: this report is the sole durable QA artifact edited by this lane. QA made no production, test, fixture, plan/ledger, target-source, or validation-asset edit. The separately authorized probe candidate and its coder report were consumed read-only at their exact Main-verified hashes.

## Executive result

`E3-P3C2-ORACLE1` passes. The authorized metadata-only probe closes the former observability gap directly on the real target: six variable-context BindingLeaf facts with ordered typed array paths `0..5`, six exact matching Variable definitions, one exact lexical owner, each DefID once in `OwnedDefIDs`, and one exact owner-local `BindingLocal` row per name. Those direct rows match the retained six distinct persisted Variables/`DEFINES`, six exact lexical `ACCESSES` endpoints, zero bounded binding gaps, and the accepted assignment/import/shadowing controls. `E3-P3C2-TARGET1` therefore passes.

`E3-P3C2-BOUNDARY1` now passes. The earlier final wrapper observed every constituent as exact but could not reproduce the retained status snapshot SHA-256 and correctly left P3-C2 `BLOCKED`. A separate bounded recovery lane then recovered the original porcelain-v2 recipe from the raw QA rollout, reproduced the exact `1,833`-byte payload and `FE3573...` hash on current target state, proved all five pre/post equality fields, and reconciled `6 tracked` versus `7 tracked` as a non-byte-backed handoff miscount. Orchestration Main independently verified and accepted that recovery handoff. With `ORACLE1`, `TARGET1`, and `BOUNDARY1` all PASS, this QA lane is `READY_FOR_SUPERVISOR`; it does not issue the Supervisor verdict itself.

## Authority and locked boundaries

Read completely before validation and verdict:

- `AGENTS.md`
- `.agents/skills/working-rules/SKILL.md`
- `.agents/skills/qa/SKILL.md`
- `.agents/skills/Data-Integrity/SKILL.md`
- `.agents/skills/supervisor/SKILL.md` as the zero-trust guard for subagent inputs; no Supervisor verdict is issued here
- the complete graph-accuracy roadmap and all four Child 03 plan/evidence/benchmark/actual-status ledgers
- `docs/contracts/graph-accuracy-contract.md`
- `reports/Supervisor/rp_supervisor_260820_131548_by_gpt-5_e3_p3c_review2.md`
- `reports/coder/rp_coder_260820_123548_by_gpt-5_p3c_owner_scope_repair.md`
- `reports/Investigation/rp_main_260820_161744_p3c2_status_boundary_recovery.md` — `16,512` bytes / `327` lines / SHA-256 `F44293C45E142742D63D71A400D975D7E48101D7C1A91F69D18739FDE1A27700` / verdict `RECOVERED`
- `reports/Investigation/rp_main_260820_162535_p3c2_boundary_recovery_verification.md` — `6,637` bytes / `93` lines / SHA-256 `24EA2A47420DC1C65D9DA3B67339AFE622B83A5F2EDCB9C09F36F3B37F8A9B5B` / Main verdict `PASS` for recovery handoff only
- the complete bounded problem origin, causal synthesis, and final bounded-investigation PASS referenced by the active ledgers

The current instruction overrides the broader active-plan closure workflow for this lane: no ledger edit, detect-changes, stage, commit, push, Supervisor verdict, or later-slice action is permitted.

## Accepted build/runtime and probe reuse gate

The P3-C full build was not rerun because every invalidation input remains exact.

| Artifact | Bytes | Current SHA-256 | Accepted SHA-256 | Result |
| --- | ---: | --- | --- | --- |
| `internal/resolution/resolve.go` | 20,103 | `C1FF5C515D401ECAD4FBF93C271DF4AC19101B2F5410D4174F7F598502BBC96A` | same | exact |
| `internal/resolution/p3c_binding_occurrence_test.go` | 15,596 | `9BAE2F63575C313B5F5F8EF4C265360BCB38F8D2F616CDCDC8942C57A080CE7F` | same | exact |
| `internal/lbugload/p3c_binding_occurrence_persistence_test.go` | 10,020 | `C704B45FD350F2A1B064D79E78B4DC99F6378D44358B4E86149A09FA38D4A850` | same | exact |
| repo-built `anvien\bin\anvien.exe` | 71,281,152 | `65D04A5B77FF8BFD47ADEEB564E4A830F7FE848366403DC5D87B3126D479DA74` | same | exact |
| installed npm runtime `anvien\bin\anvien.exe` | 71,281,152 | `65D04A5B77FF8BFD47ADEEB564E4A830F7FE848366403DC5D87B3126D479DA74` | same | exact |

`anvien version` remains `1.2.8`. Current Anvien HEAD and parent are the exact P3-C commit boundary above. The unchanged P3-C focused gate accepted by `E3-P3C-REVIEW2` remains the control authority for assignment/import/shadowing behavior; this P3-C2 lane does not claim a fresh rerun.

The observability asset was implemented and verified in the separately authorized coder/Main lanes. QA read both command files and both handoff reports completely before use, then rechecked the exact current bytes:

| Artifact | Bytes | SHA-256 | Main-authorized result |
| --- | ---: | --- | --- |
| `cmd/binding-contract-probe/main.go` | 19,607 | `C68E520AF4F220AD2E9E4A9B71941EB9B816C3EAFC15D6F7B44212305BB7BBA7` | exact |
| `cmd/binding-contract-probe/main_test.go` | 10,809 | `DA3B500FB161411A4D134DDF11935A83ECB1955F08C3087B5C78D4B4073C4F9A` | exact |
| `reports/coder/rp_coder_260820_152054_by_gpt-5_p3c2_binding_contract_probe.md` | 16,942 | `FC25935961463336A554503F9D5045B3F117BE5DC7ED7E61FEBF88F062927A5F` | exact |
| `reports/Investigation/rp_main_260820_152947_p3c2_probe_candidate_verification.md` | current | `1B743382BF1E6F2074B74CA0DF052D75A7FA7DF623F15168F7DA197B4C96E95F` | exact |

Main's zero-trust verification accepted, on those exact cmd hashes, the holder-clean `npm run full-build`, focused probe tests, uncached parser/TSJS/ScopeIR regression, `go vet`, and synthetic real-CLI happy/negative gates. QA reused those gates as explicitly authorized and did not rerun the full build, focused tests, vet, or synthetic CLI. The existing QA report's pre-update hash was `2109A1C9451A4D910D9BDFEFD3991F8C9FB061621784D275468DA38E15C19487`.

## E3-P3C2-ORACLE1 — PASS

Oracle source:

- target file: `modules/email/server/operations/email-operations-observability.ts`
- SHA-256: `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C`
- coordinates: one-based line, zero-based UTF-8 byte column, half-open ranges
- lexical owner: `readEmailOperationsReport`, selection `[497:22,497:47)`, scope `[497:7,634:1)`
- root array binding pattern: `[503:8,510:3)`
- declaration context: `const`, `BindingContextVariable`, TypeScript array binding pattern

No target source body is copied into this report.

| Name | Binding path | Declaration range = selection | Downstream name / `.map` / call span | Expected lexical target |
| --- | --- | --- | --- | --- |
| `messageRows` | `array-index(0)` | `[504:4,504:15)` | `[597:16,597:27)` / `[597:16,597:31)` / `[597:16,608:9)` | local Variable at `[504:4,504:15)` |
| `attemptRows` | `array-index(1)` | `[505:4,505:15)` | `[586:16,586:27)` / `[586:16,586:31)` / `[586:16,591:9)` | local Variable at `[505:4,505:15)` |
| `eventRows` | `array-index(2)` | `[506:4,506:13)` | `[592:14,592:23)` / `[592:14,592:27)` / `[592:14,596:9)` | local Variable at `[506:4,506:13)` |
| `providerEventRows` | `array-index(3)` | `[507:4,507:21)` | `[609:22,609:39)` / `[609:22,609:43)` / `[609:22,614:9)` | local Variable at `[507:4,507:21)` |
| `readinessRows` | `array-index(4)` | `[508:4,508:17)` | `[615:17,615:30)` / `[615:17,615:34)` / `[615:17,621:9)` | local Variable at `[508:4,508:17)` |
| `suppressionRows` | `array-index(5)` | `[509:4,509:19)` | `[622:20,622:35)` / `[622:20,622:39)` / `[622:20,626:9)` | local Variable at `[509:4,509:19)` |

Independent controls:

- each selected name has exactly one declaration and one downstream read in the oracle file;
- no selected name is an import or assignment target inside the lexical owner;
- `providerEventRows` also exists in `modules/admin-operations/server/monitoring/admin-monitoring-read-model.ts`, so a global-name or same-name rescue cannot satisfy the expected file/scope/range endpoint;
- the oracle file is tracked, absent from target status, and byte-identical before/after inspection.

## E3-P3C2-TARGET1 — PASS

### Authorized direct ScopeIR probe on the real target

QA ran the probe exactly once, read-only and stdout-only, against the one authorized target file. All Go/system temp and cache roots (`TEMP`, `TMP`, `TMPDIR`, `GOTMPDIR`, `GOCACHE`, `GOMODCACHE`, `GOPATH`, and `XDG_CACHE_HOME`) were set beneath `E:\Anvien\.tmp\p3c2-binding-probe-qa` before execution:

```powershell
go run ./cmd/binding-contract-probe `
  -file E:\cheapapp.org\modules\email\server\operations\email-operations-observability.ts `
  -name messageRows `
  -name attemptRows `
  -name eventRows `
  -name providerEventRows `
  -name readinessRows `
  -name suppressionRows
```

Result: exit `0`; stdout `26,555` bytes; stdout SHA-256 `8E382A3ADDEE3542505A0B98A8A9A57EC41606133D72E86455A1CF090F0D8F21`; schema `binding-contract-probe.v1`; caller order exact; target hash `a9247972d44d0a74c411fb7cf83d681feba371c7a85eeb83f8986afd6446e40c` (case-normalized exact oracle hash). Stderr contained dependency-download progress and the existing non-fatal Swift scanner warning only. Output was captured/parsed in memory, contained no source body/snippet, and created no output file or target artifact.

Common direct facts for all six rows:

- `BindingContextVariable` (serialized `variable`), `patternKind=array_pattern`;
- provenance pattern `[503:8,510:3)` inside construct `[503:8,582:4)`;
- `rest=false`, `default=false`;
- one matching `NodeVariable` Definition with leaf/definition full and selection ranges identical and a nonempty unique DefID;
- one `Function` owner, ID `scope:E:/cheapapp.org/modules/email/server/operations/email-operations-observability.ts#497:7-634:1:Function`, range `[497:7,634:1)`;
- owner `OwnedDefIDs` includes `readEmailOperationsReport` and every selected DefID, with the selected DefID occurring exactly once;
- exactly one owner-local `(BindingLocal,name,DefID)` (`origin=local`) for each row.

| Caller index / name | Typed path and source range | Leaf = Definition range = both selections | Exact nonempty DefID | Owner/local counts |
| --- | --- | --- | --- | --- |
| `0` / `messageRows` | `array-index(0)` at `[504:4,504:15)` | `[504:4,504:15)` | `def:E:/cheapapp.org/modules/email/server/operations/email-operations-observability.ts#504:4:Variable:messageRows` | `1 / 1` |
| `1` / `attemptRows` | `array-index(1)` at `[505:4,505:15)` | `[505:4,505:15)` | `def:E:/cheapapp.org/modules/email/server/operations/email-operations-observability.ts#505:4:Variable:attemptRows` | `1 / 1` |
| `2` / `eventRows` | `array-index(2)` at `[506:4,506:13)` | `[506:4,506:13)` | `def:E:/cheapapp.org/modules/email/server/operations/email-operations-observability.ts#506:4:Variable:eventRows` | `1 / 1` |
| `3` / `providerEventRows` | `array-index(3)` at `[507:4,507:21)` | `[507:4,507:21)` | `def:E:/cheapapp.org/modules/email/server/operations/email-operations-observability.ts#507:4:Variable:providerEventRows` | `1 / 1` |
| `4` / `readinessRows` | `array-index(4)` at `[508:4,508:17)` | `[508:4,508:17)` | `def:E:/cheapapp.org/modules/email/server/operations/email-operations-observability.ts#508:4:Variable:readinessRows` | `1 / 1` |
| `5` / `suppressionRows` | `array-index(5)` at `[509:4,509:19)` | `[509:4,509:19)` | `def:E:/cheapapp.org/modules/email/server/operations/email-operations-observability.ts#509:4:Variable:suppressionRows` | `1 / 1` |

The direct DefIDs use the absolute caller-provided path, while the retained graph node IDs use the repository-relative path. QA did not claim raw ID-string equality. The comparison key was exact file hash plus normalized repository-relative file, name, `Variable` label, range, selection, and owner scope. On that semantic identity, direct facts -> distinct Variables/`DEFINES` are `6/6`; each then reaches the already retained exact downstream endpoint, and no bounded binding gap exists.

### Normal built target analyze

The first normal command:

```powershell
anvien analyze E:\cheapapp.org --force
```

returned exit `1` because Git rejected the repository as dubious ownership. Boundary comparison showed no target write from that failed attempt.

The retry preserved normal Anvien grammar and used process-local Git configuration only:

```powershell
$env:GIT_CONFIG_COUNT='1'
$env:GIT_CONFIG_KEY_0='safe.directory'
$env:GIT_CONFIG_VALUE_0='E:/cheapapp.org'
anvien analyze E:\cheapapp.org --force
```

Result: exit `0`; scanned/parsed/failed `1359/887/0`; graph `90,899` nodes / `121,868` relationships. The process-local environment did not write global or repository Git configuration.

### Persisted graph observations

The normal Cypher surface returns six distinct Variable IDs. Each has the exact oracle file, full range, selection range, and lexical owner encoded in its ID: `readEmailOperationsReport` scope `[497:7,634:1)`.

| Name | Variable/selection range | `DEFINES` | Exact downstream `ACCESSES` span |
| --- | --- | ---: | --- |
| `messageRows` | `[504:4,504:15)` | 1 | `[597:16,608:9)` |
| `attemptRows` | `[505:4,505:15)` | 1 | `[586:16,591:9)` |
| `eventRows` | `[506:4,506:13)` | 1 | `[592:14,596:9)` |
| `providerEventRows` | `[507:4,507:21)` | 1 | `[609:22,614:9)` |
| `readinessRows` | `[508:4,508:17)` | 1 | `[615:17,621:9)` |
| `suppressionRows` | `[509:4,509:19)` | 1 | `[622:20,626:9)` |

All six `DEFINES` rows originate from `File:modules/email/server/operations/email-operations-observability.ts` and terminate at the matching distinct Variable ID. The native schema represents these as `CodeRelation {type:'DEFINES'}`.

All six downstream rows originate from `readEmailOperationsReport`, terminate at the exact expected Variable ID, and have identical contract metadata:

- `type=ACCESSES`
- `reason=read`
- `step=1`
- `confidence=1`
- `resolutionSource=scope-resolution`
- `proofKind=scope-binding`
- `sourceSiteStatus=resolved`
- `sourceSiteCount=1`
- `targetRole=binding`

This is exact endpoint evidence; `.map` call presence is not used as a substitute.

Two independent ResolutionGap queries return `row_count=0`:

1. exact target texts for all six names plus each `<name>.map` form in the oracle file;
2. exact downstream start lines `586,592,597,609,615,622` constrained to `targetRole='binding'`.

Therefore persisted binding occurrences are `6/6`, downstream exact lexical endpoints are `6/6`, and binding-caused gaps are `0` for the bounded sites.

### Assignment, import, and shadowing controls

The accepted P3-C focused boundary remains byte-valid and proves:

- assignment produces no eighth definition;
- import conservation remains `1/1/1`;
- outer/inner same-name shadowing resolves `2/2` to the nearest distinct lexical declaration;
- Graph JSON/Ladybug parity and zero skipped/fallback/warning outcomes remain accepted.

These are preserved P3-C controls under exact source/test/runtime hashes. They are not presented as new target-source cases.

### Prior observability blocker — closed by the authorized probe

At the first QA checkpoint, P3-C2 required direct proof of source oracle -> extracted binding facts -> graph occurrences -> exact endpoints -> gap outcome, while only the graph-and-later portion was observable through the normal built CLI.

Current source proves why:

- `internal/providers/tsjs/extract.go:20` returns `scopeir.ScopeIR`; line `92` places `collector.bindingLeaves` into `ScopeIR.BindingLeaves`.
- `internal/cli/command.go:203-204` sets `WriteGraphSnapshot=true` and `ReleaseScopeIRsAfterResolution=true` for normal analyze.
- `internal/analyze/analyze.go:120` marks `Result.ScopeIRs` as `json:"-"`.
- `internal/analyze/analyze.go:351-353` sets `result.ScopeIRs=nil` and releases parsed IRs after resolution.
- `cmd/graph-accuracy-probe/main.go:17-89` compares two graph snapshots through `graphaccuracy.Run`; it does not emit ScopeIR binding facts.

Consequently, at that checkpoint, `.anvien/graph.json`, Ladybug/Cypher, `file-detail`, source-site accuracy, and resolution inventory could not directly show:

- six `BindingLeafFact` rows;
- paths `array-index(0)` through `array-index(5)` and their provenance/context;
- six exact matching Variable `DefinitionFact` rows at the fact layer;
- the lexical owner's six `OwnedDefIDs`;
- exactly one matching `BindingLocal` for each name/DefID inside that owner scope.

The observed graph could only have been produced through the accepted binding-occurrence pipeline, but using that as proof would have been an inference. QA therefore correctly reported `BLOCKED`, requested a bounded metadata-only asset, and did not self-accept. The exact direct probe evidence above now closes this former `TARGET1` blocker.

## Retained owner/file-detail/impact history for the reusable asset

Before this gate, the self-graph was refreshed with both forbidden-tree exclusions:

```powershell
anvien analyze . --force `
  --exclude "internal/aicontext/skills/**" `
  --exclude ".claude/skills/**"
```

Exit `0`; scanned/parsed/failed `1109/624/0`; graph `80,177` nodes / `119,120` relationships; indexed/current commit both `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`; stale `false`.

| Existing owner | Current file-detail | Upstream impact | Ownership conclusion |
| --- | --- | --- | --- |
| `internal/providers/tsjs/extract.go` | 36 symbols; 22 inbound; 32 outbound; 29 local; 1 flow; 6 tests; risk HIGH; current/unchanged | file CRITICAL: 21 symbols / 10 files; exact `Extract` CRITICAL: 8 symbols / 5 files / 20 processes; health `0/0` | production fact source; inspect/use only, do not edit for validation |
| `cmd/graph-accuracy-probe/main.go` | 40 symbols; 0 inbound; 8 outbound; 4 local; 4 flows; 0 tests; risk HIGH; current/unchanged | file LOW: 2 symbols / 1 file; `main` LOW; `runFreshAnalyze` HIGH: 1 symbol / 1 file / 4 processes | existing owner compares/copies graph snapshots; adding ScopeIR inspection would mix responsibilities |
| `internal/graphaccuracy/graphaccuracy.go` | 300 symbols; 142 inbound; 3 outbound; 184 local; 11 flows; 11 tests; risk HIGH; current/unchanged | file CRITICAL: 107 symbols / 32 files / 63 direct; `Run` HIGH: 1 symbol / 1 file / 4 processes | graph comparison owner, not a binding-fact probe owner |

Exact canonical UIDs were obtained from current `file-detail` and passed with `impact symbol <name> --uid <uid>`. HIGH/CRITICAL are blast-radius warnings and justify keeping any future validation asset isolated; they are not edit prohibitions.

At that earlier checkpoint, no asset was created or edited by QA.

## Prior authorization request and closure

QA requested, and Orchestration Main later authorized in a separate visible coder lane, exactly this one-responsibility validation boundary:

- asset owner/path: `cmd/binding-contract-probe/main.go`
- focused test owner/path: `cmd/binding-contract-probe/main_test.go`
- purpose: read one explicitly named TypeScript file, invoke the production parser/`tsjs.Extract`, filter an explicit list of local names, and emit metadata-only JSON to stdout; never copy source bodies and never write into the target
- minimal code boundary: the two new files only; no edit to `internal/providers/tsjs/**`, `internal/scopeir/**`, `internal/analyze/**`, `internal/cli/**`, `internal/resolution/**`, `cmd/graph-accuracy-probe/**`, or `internal/graphaccuracy/**`
- minimal test boundary: independent synthetic in-memory TypeScript, not copied target source; assert deterministic metadata-only output and failure on missing/ambiguous facts
- write boundary: no `--out` option and no target artifact; retained command output/report remains under `E:\Anvien`

The authorized asset must directly observe and report, for the six selected names:

1. one `BindingLeafFact` each with `BindingContextVariable`, exact range/selection, path segments `[0]` through `[5]`, and pattern provenance;
2. one matching Variable `DefinitionFact` each with the exact DefID/range/selection;
3. lexical owner scope `[497:7,634:1)` with each DefID present exactly once in `OwnedDefIDs`;
4. exactly one `(BindingLocal,name,DefID)` inside that same owner;
5. no source body and no write outside Anvien.

The coder lane implemented only those two cmd files, ran the holder-clean full build and focused gates, and Main independently verified the exact final hashes. This resumed QA lane consumed the asset read-only, compared its direct rows to the retained graph/endpoints/gap evidence, and did not open a Supervisor gate because the final boundary wrapper remains blocked below.

## E3-P3C2-BOUNDARY1 — PASS

### Retained pre-recovery `BLOCKED` checkpoint

Pre-run target boundary:

- HEAD: `a869876ab6262dacde6cd5d432d099a91852a646`
- status snapshot SHA-256: `FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E`
- tracked diff Git object, produced by `git diff --no-ext-diff --binary --full-index --no-color | git hash-object --stdin`: `941c3e00be7357de8393d959b91ca93be72e64fb`
- untracked manifest SHA-256: `886DE642875AA7D58F5F8357054927AADA49CAE65604120AB28AA788F0E81B97`
- untracked inventory: 6 files / 573,963 bytes
- oracle source SHA-256: `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C`

Immediate post-analyze comparison returned `true` for HEAD, status, tracked diff, untracked manifest, source hash, and targeted report/fixture/probe/debug path sets. Those retained facts authorized reuse after the read-only probe.

The final permitted post-probe byte-stream invocation completed without target mutation and observed:

- HEAD exact: `a869876ab6262dacde6cd5d432d099a91852a646`;
- status count exact: 13 entries (seven unstaged tracked modifications and the same six untracked paths);
- tracked diff object exact: `941c3e00be7357de8393d959b91ca93be72e64fb`;
- untracked manifest exact: six files / 573,963 bytes / SHA-256 `886DE642875AA7D58F5F8357054927AADA49CAE65604120AB28AA788F0E81B97`;
- oracle source exact: 19,328 bytes / SHA-256 `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C`;
- all four existing analyzer-owned `.anvien` files exact as shown below;
- target artifact-name scan outside `.git`, `.anvien`, and `node_modules`: zero match for P3-C2/binding-contract-probe/QA-report patterns.

The six current untracked content rows also match the retained manifest exactly:

| Path | Bytes | SHA-256 |
| --- | ---: | --- |
| `reports/QA/rp_qa_260619_145255_by_gpt5_user_admin_lifecycle.md` | 12,910 | `354A19EC02BF7989F5CC41AFF6E817918AC025A92DB6F0949F81C821980E1725` |
| `reports/QA/rp_qa_260619_152926_by_gpt5_release_catalog_mutation.md` | 5,442 | `D781C90EA939C71348ED3A3CF0DF6FCD7E5916E0829AC81722BEAD66E6FCFBD3` |
| `reports/problem/images/login-idea-2.jpg` | 150,157 | `06FD3EA35F17BFB94CA9BF1A3AD1297DC0EBB6BF9E6F247AC2DA9E0634C49368` |
| `reports/problem/images/login-idea-3.jpg` | 260,724 | `24B8E928F2D2EBE517010F0A535D3AA0703C62D9F89F931CADDB477A61F8192D` |
| `reports/problem/images/login-idea.jpg` | 75,770 | `71AF93B0D5A516F7B22CBA75FCF4A479797BAF04001E3E395FBBCAEBE2007944` |
| `reports/problem/images/wallet-idea.jpg` | 68,960 | `30EECD84F8218248B27B9616A5580D5C5F115132FFEEC87E5CFC3E1BE6813B38` |

The only analyzer-owned target changes are permitted operational output:

| Target path | Bytes | Current SHA-256 | Classification |
| --- | ---: | --- | --- |
| `.anvien/graph.json` | 423,734,194 | `5D82C464C291ED7734E9BF32C3F79CF2350008C62F83F06B67517FD52EF97E05` | normal analyze output |
| `.anvien/lbug` | 144,744,448 | `B3173C66D0273149B3AD63D11D529A5B15B9FDA098408922704A32E938550EE8` | normal analyze output |
| `.anvien/meta.json` | 259 | `226BB1FDC696C15153B650EDE607DF067C1C78EA218BA7C5FF72CB8B939ECFEC` | normal analyze output; records `1359` files / `90899` nodes / `121868` edges |
| `.anvien/settings.json` | 31 | `6C831C0AC49625D09EE54E76042612B92F47ABC47E4A90D1310F0EA45E824B3C` | unchanged |

No target-side P3-C2 report, fixture, copied source, probe, or debug artifact was observed. The probe was read-only. QA removed only `E:\Anvien\.tmp\p3c2-binding-probe-qa` after verification; final existence was `false`. The full-build Ladybug cache and every unrelated temp path were preserved.

The sole unresolved boundary field is the retained status snapshot SHA-256. The final wrapper tested the current 13-line status under git-order/sorted ordering; LF/CRLF/NUL/JSON serialization; and UTF-8/UTF-16 LE/BE with and without BOM, but returned `hashMatches=[]` for expected `FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E`. Because HEAD, exact status entries/count, tracked diff object, untracked manifest and content hashes, oracle, `.anvien`, and contamination scan all agree, this does not identify a target delta. It identifies that the original status-snapshot serialization recipe was not retained/reproduced by the authorized final wrapper.

Orchestration Main explicitly directed: if that final invocation still failed, do not run a fourth wrapper/command and update this report `BLOCKED` with the exact wrapper blocker. QA complied. At that checkpoint, `BOUNDARY1` could not be promoted to PASS even though no contamination was observed. This historical classification and its failure ledger are retained rather than erased.

### Recovered status provenance and invariant closure

The recovery lane subsequently recovered the exact original status-capture recipe from the completed QA task's raw rollout and ran exactly one bounded current-target reproduction batch. This resumed QA turn did not access `E:\cheapapp.org` and did not rerun that batch.

Exact effective command:

```text
git -c safe.directory=E:/cheapapp.org -c core.quotepath=false status --porcelain=v2 --branch --untracked-files=all
```

Exact serialization contract:

- preserve Git output order as PowerShell materializes the ordered `[string[]]`;
- serialize with `[string]::Join([char]10, $statusLines)`;
- LF separators only, with no trailing LF and no NUL;
- encode through `System.Text.Encoding.UTF8.GetBytes`, without BOM;
- payload `1,833` bytes / `15` lines: two branch headers plus `13` status entries;
- entries `7 tracked + 6 untracked`;
- SHA-256 `FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E`.

The original transcript rows and bounded current reproduction are byte-identical. Both yield:

| Boundary field | Recovered/current result |
| --- | --- |
| target HEAD | `a869876ab6262dacde6cd5d432d099a91852a646` — exact |
| status payload | `1,833` bytes / `15` lines / `13` entries (`7 tracked + 6 untracked`) — exact |
| status SHA-256 | `FE3573ADB361095CF36BAD90D56912FB2EC1F92FFF0628BE19DB12E4F27E4E2E` — exact |
| tracked diff object | `941c3e00be7357de8393d959b91ca93be72e64fb` — exact |
| untracked manifest | `6` files / `573,963` bytes / SHA-256 `886DE642875AA7D58F5F8357054927AADA49CAE65604120AB28AA788F0E81B97` — exact |
| pre/post equality | HEAD, status bytes, status entries, tracked diff, and untracked manifest all `true` |

The phrase `6 tracked` in `reports/Investigation/rp_main_260820_141241_orchestration_rotation_handoff.md:48` is a non-byte-backed handoff miscount. No direct evidence supports a 12-entry state: the retained hash, exact original rows, and exact current reproduction all prove seven tracked rows plus six untracked rows. No omitted path was inferred.

Recovery authority and independent verification:

- `reports/Investigation/rp_main_260820_161744_p3c2_status_boundary_recovery.md`: `16,512` bytes / `327` lines / SHA-256 `F44293C45E142742D63D71A400D975D7E48101D7C1A91F69D18739FDE1A27700` / `RECOVERED`;
- `reports/Investigation/rp_main_260820_162535_p3c2_boundary_recovery_verification.md`: `6,637` bytes / `93` lines / SHA-256 `24EA2A47420DC1C65D9DA3B67339AFE622B83A5F2EDCB9C09F36F3B37F8A9B5B` / Main `PASS` for the recovery handoff only, explicitly not `E3-P3C2-REVIEW1`.

Main independently traced the original and recovery rollouts, verified that the recovery task used exactly one target-workdir command, and confirmed `StatusHashMatch`, `HeadMatch`, `TrackedDiffMatch`, `UntrackedManifestMatch`, `CountMatch`, and all five pre/post equality fields as `true`. The recovered evidence closes the sole former wrapper-provenance blocker without changing or rerunning accepted product behavior. `E3-P3C2-BOUNDARY1` is therefore `PASS`.

## Command and failure ledger

Principal successful commands and results:

```powershell
anvien --help
anvien analyze E:\cheapapp.org --force
anvien cypher "<six exact Variable rows>" --repo E:\cheapapp.org
anvien cypher "<six CodeRelation/DEFINES rows>" --repo E:\cheapapp.org
anvien cypher "<six exact binding ACCESSES rows>" --repo E:\cheapapp.org
anvien cypher "<gap query by exact names and .map forms>" --repo E:\cheapapp.org
anvien cypher "<gap query by exact lines and targetRole=binding>" --repo E:\cheapapp.org
anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
anvien file-detail internal/providers/tsjs/extract.go --repo E:\Anvien --json
anvien file-detail cmd/graph-accuracy-probe/main.go --repo E:\Anvien --json
anvien file-detail internal/graphaccuracy/graphaccuracy.go --repo E:\Anvien --json
anvien impact file <exact-file> --repo E:\Anvien --direction upstream --json
anvien impact symbol <name> --uid <canonical-uid> --repo E:\Anvien --direction upstream --json
go run ./cmd/binding-contract-probe -file E:\cheapapp.org\modules\email\server\operations\email-operations-observability.ts -name messageRows -name attemptRows -name eventRows -name providerEventRows -name readinessRows -name suppressionRows
git -c safe.directory=E:/cheapapp.org rev-parse HEAD
git -c safe.directory=E:/cheapapp.org status --short --untracked-files=all
git -c safe.directory=E:/cheapapp.org diff --no-ext-diff --binary --full-index --no-color
git -c safe.directory=E:/cheapapp.org ls-files --others --exclude-standard
Get-FileHash -Algorithm SHA256 <oracle-or-existing-.anvien-path>
rg --files -uu -g '!.git/**' -g '!.anvien/**' -g '!node_modules/**'
```

The final wrapper kept the `git diff` byte stream in memory and computed its Git blob object as SHA-1 of `blob <byte-length>\0<exact-bytes>`; it kept status/manifest candidate serializations and all file hashes in memory. It created no target or report-side capture file.

Failures and classification:

- target analyze attempt 1: exit `1`, Git dubious-ownership gate; environment boundary failure, zero target write; retry with process-local safe-directory configuration passed.
- `MATCH ...-[:DEFINES]->...`: exit `1`; query-schema mismatch because the persisted native schema uses `CodeRelation {type:'DEFINES'}`; corrected query returned six rows. This is not a missing-edge failure.
- first exact-UID impact invocation passed the UID positionally and returned `UNKNOWN/not found`; `anvien impact symbol --help` showed that exact UID requires `--uid`; corrected commands passed. This is an invocation classification, not missing graph ownership.
- one combined parallel `file-detail` capture was truncated by the output layer; it was discarded. Compact summaries and canonical UIDs were re-extracted directly from fresh JSON.
- several read-only PowerShell formatting/hash experiments failed before mutation (parser/escaping or alias conflicts). None is validation evidence, none wrote a file, and later direct Git/hash commands produced the recorded boundary values.
- post-probe boundary wrapper attempt 1 stopped after read-only HEAD/status at `Argument types do not match` while returning a generic in-memory list; no boundary verdict and no target write.
- post-probe boundary wrapper attempt 2 stopped after read-only HEAD/status because `cmd.exe` passed the escaped pipe to Git, which returned `ambiguous argument '|'`; no boundary verdict and no target write.
- the Main-authorized final byte-stream wrapper completed all constituent checks, but `status.hashMatches=[]` for the retained SHA-256. Every other constituent was exact. Main explicitly prohibited a fourth wrapper/command, so this was the precise `BOUNDARY1` blocker at that checkpoint; the later raw-rollout recovery above closes it without obscuring this history.
- repo-temp cleanup: two `Remove-Item` invocations were rejected before process start; recursive `.NET Directory.Delete` then stopped on cached read-only `.editorconfig`; QA normalized attributes only inside the already verified literal `E:\Anvien\.tmp\p3c2-binding-probe-qa` and deleted it successfully (`existsAfter=false`). No target path or unrelated temp/cache was touched.
- recovery provenance: the original recipe was not one of the prior candidate recipes because it uses porcelain-v2 with two branch headers. Direct raw-rollout reconstruction and one bounded current reproduction both produced the retained `FE3573...` hash; Main independently verified both rollouts. This is closure evidence, not a reclassification of the historical wrapper attempts as successful.

Not run by design:

- no P3-C2 production/test repair and no QA edit to the separately authorized reusable asset;
- no full-build/focused-test/vet/synthetic-CLI repeat because exact cmd hashes and Main-accepted gates remained valid;
- no second target analyze, probe rerun, Cypher/query expansion, or fourth boundary wrapper;
- no `E:\cheapapp.org` access and no product/target command in this recovery-consumption turn;
- no detect-changes, Git add/commit/push/merge/reset/checkout/stash/amend;
- no plan/ledger edit, Supervisor verdict, or later-slice action;
- no scanner/export/module/re-export/ambient/external/reader/transport/cache/MCP/HTTP/Web/UI expansion.

## Evidence verdict matrix

| Evidence ID | Verdict | Reason |
| --- | --- | --- |
| `E3-P3C2-ORACLE1` | PASS | six independently identified source sites, typed paths, exact ranges/selections, lexical owner, and downstream targets; no copied source body |
| `E3-P3C2-TARGET1` | PASS | direct BindingLeaf/path/Definition/owner/OwnedDefID/BindingLocal rows `6/6` match retained distinct Variables/`DEFINES` `6/6`, exact endpoints `6/6`, bounded binding gaps `0`, and accepted assignment/import/shadowing controls |
| `E3-P3C2-BOUNDARY1` | PASS | exact original porcelain-v2 recipe and one bounded current reproduction both yield `FE3573...`; HEAD/status/diff/untracked pre/post equality is exact, `7 tracked + 6 untracked` is byte-backed, and Main independently verified the recovery |

P3-C2 is PASS-ready for a separate visible Supervisor gate. This QA conclusion is evidence readiness only and is not self-acceptance or `E3-P3C2-REVIEW1`.

Handoff to Orchestration Main - independently verify this final QA report and, only then, open one separate visible `E3-P3C2-REVIEW1` Supervisor gate.

READY_FOR_SUPERVISOR
