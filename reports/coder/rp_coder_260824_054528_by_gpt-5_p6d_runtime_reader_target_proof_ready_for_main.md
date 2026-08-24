# P6-D R43 — Current Runtime, Reader, and Conditional Target Proof

- Status: `READY_FOR_MAIN`
- Procedure verdict: `CLEAN`
- Coder lane: continuation of the visible P6-D implementation Coder
- Next responsible: Visible Main task `01a030b6-9de3-7c72-977b-31d546108138`
- Repository authority: `E:\Anvien`
- Durable raw proof root: `E:\Anvien\reports\coder\p6d_runtime_reader_target_proof_260824_045429_raw`
- Report prepared: `2026-08-24T05:45:28+07:00`

## Verdict

The current canonical runtime identity, source/native/publication provenance, package/native/built-reader parity, independent target oracle, exact three-site target outcomes, and target pre/post preservation boundary all pass. The target was not accessed until the durable `READER_GATE_PASS` marker existed. There is zero reader fallback, zero skip, zero structured difference, zero resolved/unresolved overlap, and zero in-repository analyzer gap across the three target sites.

The P6-D candidate is ready for zero-trust verification by Visible Main. This report is a Coder handoff, not self-acceptance. No plan/ledger update, Supervisor run, Git stage, commit, or push was performed.

## Fixed inputs and current runtime

The dependency inputs remain byte-identical:

| Authority | Bytes | SHA-256 |
|---|---:|---|
| `E:\Anvien\go.mod` | 2,392 | `C203E25E0A83A583E89A9D8F8E65BC0881052B3075E46915F90A4FB6433BD249` |
| `E:\Anvien\go.sum` | 16,065 | `5434F39B08F50F1C53A71333E2E971CFCDE0E80BCA7DD028C00E31B9ACCAEC73` |

The accepted exact canonical Windows PowerShell 5.1 full-build invocation passed `1/1`, exit `0`, with inherited/unmodified `PSModulePath`, unset/default `PSModuleAnalysisCachePath`, Rule 13 holder/write-open gate, vendor verification before the first Go boundary, local Go `1.26.3`, and the offline vendor contract. It published version `1.2.8`; launcher/Vite transformed `2,943` modules; the canonical final analyze reported `2,127` scanned, `765` parsed, `0` failed, graph `122,531` nodes / `169,182` relationships, and `479` unresolved. This gate was not rerun during the reader/target continuation.

Current identities were remeasured read-only after the target proof:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `E:\Anvien\anvien\bin\anvien.exe` | 73,749,504 | `64EECEBAF459F23EA11DA38B29C5E290F6F6B22D5EE4749E068BB927A8EEDEBB` |
| `E:\Anvien\anvien\bin\lbug_shared.dll` | 20,230,656 | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| `E:\Anvien\anvien\bin\anvien-runtime.json` | 225 | `C454EBDF4CD28AD61D0D9A903C0C761F2E135316BE2325E3F59C6A3A45B2C78D` |

The canonical runtime provenance is `complete`, 170,548 bytes, SHA-256 `149000ABA76ED45F7D6323D46138522E824B9FF8DC9B2CB928D0931EF5336C69`, with source identities `10/10`, native files `3/3`, and publication identities `3/3`.

## Reader-before-target gate

The durable authorization marker is:

- `E:\Anvien\reports\coder\p6d_runtime_reader_target_proof_260824_045429_raw\READER_GATE_PASS.json`
- 6,136 bytes
- SHA-256 `2F5341598EDAFDDA4EB3A54698368468CDC921823CB8E8EF247A48F84980A282`
- LF `148`, CR `0`, UTF-8 without BOM

### Package and native-through-runtime

The corrected package-only boundary used outer Windows PowerShell `5.1.19041.6456 Desktop`, inherited `PSModulePath` unchanged (SHA-256 `7FCFB77D5B356E43406539C62F0C8F228B2E6C1AE53754175A25A77FD52690AE`), left `PSModuleAnalysisCachePath` unset/default, resolved `Get-FileHash`, and invoked the current canonical runtime's `package ensure-runtime`. It exited `0`.

The package verifier reported `45` modules, `1,798` files, `190,350,404` bytes, vendor tree SHA-256 `7C6D6466D73874684F73C1C35D0D928FA0B729B8C8F59D5F9C71079639C6330E`, and `0/0/0/0` missing/extra/mismatch/license gaps. The live process loaded exactly `E:\Anvien\anvien\bin\lbug_shared.dll` with SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`; module observation errors were `0`.

### Carrier and Child 02 readers

The nearest existing carrier tests passed after the production/full-build gate:

| Boundary | Result | Denominator |
|---|---|---:|
| `internal/graphhealth` P6-D carrier | PASS | `3+2+26+1+1` |
| Graph JSON -> Ladybug round-trip | PASS | `1/1` |
| Native Ladybug outcome readback | PASS | `1/1` |
| Child 02 C16 native-reader commands 3, 4, 5 | PASS | `3/3` |

All carrier fallback/failure/skip counts are zero. Child 02 remote VECTOR attempts are zero.

Child 02 Definition/DEFINES parity also passed:

- Definition rows: Graph JSON `46,593`, Ladybug `46,593`.
- DEFINES pairs: Graph JSON `46,593`, Ladybug `46,593`.
- Missing/extra/field/pair/endpoint/partial mismatch: `0/0/0/0/0/0`.
- Durable QA JSON: `E:\Anvien\reports\QA\child02-p2e-parity\p6d-r43-260824_050457\qa_child02_p2e_parity_260824_050601.json`, 3,169 bytes, SHA-256 `B18E607FCD4AFBD8FFA02326FC7653134162A45C7387744472F0288C31193A91`.
- Durable QA Markdown: 631 bytes, SHA-256 `D56D328B83127357FB6862D62C012C5D733E145CB855658E981292BBFDE4E8AA`.

### Built and native-only readers

The isolated registration preparation and all five required built-reader commands passed with exit `0`, empty stderr, canonical DLL observation, zero module-observation errors, and unchanged runtime/Graph JSON/Ladybug identities:

1. `resolution-inventory`
2. `graph-health summary`
3. `graph-health report`
4. `graph-health components`
5. `graph-health files`

The command matrix is 7,426 bytes, SHA-256 `405A751C9CFB84087F5EA9798A71632E8640BE468930319E99BB311250EBA275`.

The native-only proof ran with no `graph.json` available and reconstructed the same `53,156` ResolutionGap rows from the Ladybug authority: `2,730 standard_library/non_actionable` plus `50,426 unclassified/review`; count difference `0`, fallback possible `false`. Its matrix SHA-256 is `560DC0312176F5E2BD66051C76C43D7AC235576EDAED9B2CB647AF1C571E9D96`.

The byte-exact structured sample compared one built row and one native row over outcome, proof, classification, actionability, resolution source, note, and authority fields. Differences were `0`; `catalogProofState=ready`. Durable proof SHA-256: `331D127F0503CB7883AE3318CA77590F67F91D9982D7C337090DCD3FD24EA823`.

Reader-gate verdict: `READER_GATE_PASS`; fallback `0`, skip `0`, difference `0`.

## Conditional target proof

Target access began only after the durable reader marker above.

### Pre-state and independent oracle

The target pre-state was captured before first access:

- HEAD `a869876ab6262dacde6cd5d432d099a91852a646`
- Parent `79ab9b101cc21ec8da79dab724d435e87a6ea6f6`
- Tree `dc024c1d92bdee6211f389a97d0431c267fe5324`
- Index delta `0`
- Pre-existing target worktree changes were recorded and preserve-only.

The independent TypeScript oracle used the target's TypeScript `6.0.3` and ES2022 semantics, derived from source/config/compiler behavior rather than analyzer output. It independently classified Promise, Math.max, and Math.min as external standard-library capabilities `3/3`; exact aligned source-site IDs passed `3/3`. The aligned oracle is 5,927 bytes, SHA-256 `4FBB9D5958CFCF7346710A738225100BA3AD0F1E76E8F0A883EDE7C51ED3B32B`.

### Exact normal analyze

Exactly one normal target analyze was run with the current canonical runtime:

`E:\Anvien\anvien\bin\anvien.exe analyze E:\cheapapp.org --force`

It exited `0`, scanned `1,359`, parsed `887`, failed `0`, and produced graph `95,762` nodes / `129,716` relationships. Stdout is 1,277 bytes, SHA-256 `BC7D7F8A86F4769768BDCCE43F2ABEB2F5C722C97AE508494367D660AEEBF714`; stderr is empty. The process loaded the canonical DLL and reported zero module-observation errors.

### Three required source sites

| Site | Analyzer outcome | Reason | Proof/classification/actionability | Acceptance |
|---|---|---|---|---:|
| Promise at `read-admin-commercial-config.ts:13` | `capability_unavailable` | `config_topology_unsupported` | `typescript-standard-library-authority` / `standard_library` / `non_actionable` | `1/1` |
| Math.max at `email-operations-observability.ts:191` | `capability_unavailable` | `config_topology_unsupported` | `typescript-standard-library-authority` / `standard_library` / `non_actionable` | `1/1` |
| Math.min at `email-operations-observability.ts:191` | `capability_unavailable` | `config_topology_unsupported` | `typescript-standard-library-authority` / `standard_library` / `non_actionable` | `1/1` |

The independent oracle says all three names are external/capability-available under TypeScript semantics. The analyzer correctly reports an accepted capability-unavailable outcome because the target configuration topology contains unsupported `include` and `exclude` keys. This is not an in-repository unresolved reference and is not an analyzer gap.

Aggregate acceptance is `3/3`; in-repository analyzer gaps are `0/3`; resolved/unresolved overlap is `0`. Graph JSON reader rows, native gap rows, and native gap edges are each `3/3`; external rows are `0`, as expected for this capability-unavailable outcome. All structured authority, outcome, proof, classification, actionability, note, source-hash, and config-hash comparisons agree.

Durable target proofs:

| Proof | Bytes | SHA-256 |
|---|---:|---|
| `target-three-site-reader-parity.json` | 5,413 | `D161D81DA0E07D8C9590567F1FCA93E82FC84E15C6C54573A9B33E4BAD746C75` |
| `target-oracle-analyzer-comparison.json` | 2,285 | `4FE5CD5C70FABEFFA4CB893D3B3AC4DB81A8A254FD75C31BCF8ABAECB0780CC2` |
| `target-boundary-comparison.json` | 1,764 | `FD53FB49A76DA81E75FC045A2164B0CE2DD47F7248C7DB228EA8BEBB8D3255EE` |

### Target post-state and allowed delta

The post-state retained byte-identical HEAD, parent, tree, status, cached diff, worktree diff, and Git index identity. Candidate source files matched `203/203` with differences `0`; compiler/config files matched `3/3` with differences `0`; reparse points were `0`. The only target delta was normal analyzer-owned output under `.anvien`: `graph.json`, `lbug`, and `meta.json`; `.anvien/settings.json` remained unchanged. Runtime identity remained unchanged.

Target-gate verdict: `PASS`.

## Procedure and artifact lifecycle

Procedure verdict is `CLEAN`, based only on durable artifacts and actual boundary impact.

Literal `E:\Anvien\.tmp` is disposable reproducible scratch. It is not evidence, review input, handoff authority, recovery authority, source of truth, acceptance dependency, or procedure-verdict input. Owner may delete it wholesale at any time. Its existence or absence is therefore not a blocker, and no acceptance statement in this report depends on it.

The durable lifecycle statement is `E:\Anvien\reports\coder\p6d_runtime_reader_target_proof_260824_045429_raw\tmp-disposition.json`, 482 bytes, SHA-256 `9E16ECABD18404EFCE31A32CF192DE9278830058C8961D84693DE18021AE9060`, LF `14`, CR `0`, strict UTF-8 without BOM. Its action is `RETAINED_DISPOSABLE_EXCLUDED` and every authority/evidence/review/handoff/acceptance/procedure flag is false.

Historical harness incidents remain recorded rather than erased: outer-pwsh `PSModulePath` contamination, shared module-analysis-cache contamination, the `$HOME` constant-variable wrapper failure, the earlier `New-Item -LiteralPath` wrapper error, a wrong-shell package invocation that failed before native load, two registry-preparation attempts before `--allow-non-git`, a missing-whitespace capture syntax error, a `Get-History` alias collision, an initial comparator suffix parse error, and one final read-only `HEAD^{tree}` quoting error during report preparation. These incidents caused no unauthorized network access, no source/config/index mutation, no accepted runtime or target-boundary drift, and no durable output corruption. They are excluded historical harness failures, not product/reader mismatches.

The unexpected repo-local `Microsoft\Windows\PowerShell\ModuleAnalysisCache` artifact is preserved unchanged and excluded as directed. Owner deletion `CONTRIBUTING.md`, unrelated/protected work, `vendor/`, and `third_party/` were preserved. No C-drive workspace/read/write, alternate worktree, shared dependency cache, network, install, package script, stage, commit, or push occurred in the reader/target continuation.

## Git preservation at seal preparation

Read-only final capture before report seal:

- HEAD `741335f7efa5ad215ec28361d88c462f431565ab`
- Parent `838f379ab00c93dfe7507603f6608bb08d61f038`
- Tree `65c4be5d1aeab6cad6faa0f3d9a70b93c76bbda8`
- Git index delta `0`

The worktree remains intentionally dirty with the open P6-D implementation, durable vendor/provenance/report authorities, Main-owned ledger/roadmap updates, Owner deletion `CONTRIBUTING.md`, the excluded Microsoft artifact, and unrelated pre-existing work. Nothing was staged. A post-seal read-only check must confirm the same HEAD and empty index; the only additional expected path is this untracked Coder report.

## Handoff

`READY_FOR_MAIN` means:

- current runtime/provenance identity PASS;
- package/native/built-reader gate PASS with zero fallback/skip/difference;
- target Promise/Math.max/Math.min acceptance `3/3` and in-repository analyzer gaps `0/3`;
- target source/config/Git preservation PASS with only normal `.anvien` output delta;
- procedure `CLEAN` under the Owner's definitive `.tmp` semantics;
- Git index empty and no commit performed.

Visible Main task `01a030b6-9de3-7c72-977b-31d546108138` must verify this sealed report and durable raw proofs. It alone decides planner refresh and any later Supervisor routing.
