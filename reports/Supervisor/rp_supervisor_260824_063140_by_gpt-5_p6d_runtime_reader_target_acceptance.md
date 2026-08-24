# Supervisor Report: P6-D Runtime, Reader, and Target Acceptance

## Metadata and verdict

- Review role: independent, visible, review-only Supervisor.
- Reviewed slice: Child 06 P6-D, `Project outcomes and prove the target sites`.
- Verdict: **PASS**.
- Review workspace: sole shared repository `E:\Anvien`; no alternate worktree.
- Review prepared: `2026-08-24T06:31:40.9550331+07:00`.
- Reviewed Coder handoff: `reports/coder/rp_coder_260824_054528_by_gpt-5_p6d_runtime_reader_target_proof_ready_for_main.md`.
- Next responsible: visible Main `01a030b6-9de3-7c72-977b-31d546108138`.
- P6-E disposition: still exactly one U1-U4 slice and **LOCKED**. This report does not open, review, update, or accept P6-E.

The accepted P6-D claim is the current Windows canonical-runtime chain: the already-cleared graph-health projector, exact checked-in Go vendor/native authority, canonical PowerShell 5.1 build, current runtime, package/native/built readers, exact target outcomes, target preservation, and clean procedure. This is acceptance review only. No production, test, script, vendor, third-party, plan, ledger, or existing report was repaired or changed by this Supervisor.

## Authority read and evidence boundary

I read the required raw authority through EOF: `AGENTS.md`, `working-rules`, `supervisor`, the campaign roadmap, all four Child 06 ledgers, both prior P6-D Supervisor reports, the current Coder report, its durable raw proof root, relevant Main rotation handoffs, the P6-A architecture decision, and referenced Child 02 QA/report authority. The Coder report seal independently matches exactly:

- `13,436` bytes;
- `176` LF / `0` CR;
- strict UTF-8 without BOM;
- SHA-256 `E79257D64788BA31DC9B7D9FE010E893A03EDC7CF732137A40EF35089AC50EC3`;
- created and last-written `2026-08-24T05:47:41.1679132+07:00`.

No Anvien semantic-graph command was used in this review. Therefore no ceremonial self-analyze was required. `anvien --help` was the only Anvien CLI invocation; all acceptance conclusions below come from current source/diff inspection, detached filesystem identities, durable artifacts, direct Git reads, and the already-current runtime evidence. No current code/runtime/target identity proved the canonical build, reader matrix, self-analyze, or target analyze invalidated, so none was rerun.

## Executive decision

P6-D now closes the blockers retained by both earlier reviews:

1. the six nested-authority/outer-outcome identity groups are exact and fail closed;
2. the canonical full build produced the current runtime and has not been invalidated;
3. vendor, native, package, built-reader, Graph JSON, and Ladybug identities agree;
4. the reader gate durably precedes target access;
5. Promise, Math.max, and Math.min have exact, source-backed accepted capability outcomes `3/3`, with in-repository analyzer gaps `0/3` and resolved/gap overlap `0`;
6. target source/config/Git state is preserved, with only the allowed analyzer-owned `.anvien` output delta;
7. historical harness failures have demonstrated source/config/index/accepted-output impact `0` and are not used as positive evidence.

No blocking same-scope defect or evidence gap remains. Residual surfaces are recorded below and routed without silently expanding P6-D.

## Prior history and blocker closure

### First P6-D rejection

`rp_supervisor_260823_044455_by_gpt-5_p6d_graph_health_projection_and_target_proof.md` proved six false acceptances when only nested authority drifted from its outer outcome: file path, file hash, full range, site kind, requested name, or requested meaning. Current `internal/graphhealth/diagnostics.go` now requires exact equality for all six groups. `internal/graphhealth/p6d_outcome_projection_test.go` contains six isolated hostile mutations within the complete `26/26` malformed/drifted-carrier matrix. Each returns `unclassified/review`.

### Resubmission rejection

`rp_supervisor_260823_064551_by_gpt-5_p6d_graph_health_projection_resubmission.md` independently cleared the six-field repair but retained absent full-build/current-runtime/native/built-reader/target/oracle/boundary proof and an unclean procedure. Those missing gates now close as follows:

| Historical blocker | Current closure |
|---|---|
| Canonical full build absent | Exact Windows PowerShell 5.1 invocation PASS `1/1`, exit `0`; vendor verification precedes Go; launcher/Vite `2,943` modules; canonical analyze `2,127/765/0` and graph `122,531/169,182` |
| Current runtime absent | Version `1.2.8`; `73,749,504` bytes; SHA-256 `64EECEBAF459F23EA11DA38B29C5E290F6F6B22D5EE4749E068BB927A8EEDEBB` |
| Native provenance absent | Exact durable Windows x86_64 bundle `3/3`, `52,743,720` bytes, no reparse point; canonical DLL byte/hash equals durable DLL |
| Built/native readers absent | Five built readers plus isolated registration PASS; native-only gaps `53,156/53,156`; structured difference `0` |
| Target/oracle absent | Reader marker precedes target; one target analyze; independent compiler oracle plus corrected exact site alignment; accepted sites `3/3` |
| Target boundary absent | Source `203/203`, config `3/3`, HEAD/parent/tree/index/status/diffs preserved; allowed `.anvien` delta only |
| Procedure unclean | No accepted claim depends on `.tmp`; all current actions and historical failures were classified by actual boundary impact; no source/config/index/accepted-runtime corruption |

## Source-level clearance

### Graph-health projection and reader path

- `normalizeDiagnosticMetadata` gives a structured outcome first authority over prefilled or legacy target-text heuristics.
- `decodeStructuredResolutionOutcome` treats status/schema markers as structured and fails malformed carriers closed.
- `validStructuredResolutionDiagnostic` binds the structured payload to the flat diagnostic over source site, status, stage, site kind, file path/hash, full range, requested name, proof kind, and source.
- `validStructuredResolutionAuthority` additionally binds nested authority to the outer outcome over source site/stage plus the exact six previously rejected identity groups. Authority/status lifecycle is constrained for resolved external, unresolved external, and capability-unavailable outcomes.
- `policy.go` already contains the required `standard_library/non_actionable`, `in_repo_unresolved/analyzer_gap`, and fail-closed `unclassified/review` values and required no edit.
- `resolution_gap_inputs.go` copies normalized diagnostic classification, actionability, note, file/hash/range, source-site status, and proof into `ResolutionGap` nodes and `HAS_RESOLUTION_GAP` relationships. The P6-D projection test round-trips Graph JSON, creates the gap graph, and verifies graph-health summary parity.

There is no Promise/Math literal or name-specific reclassification in graph-health production code.

### Vendor, native, package, and publication path

- `materialize-go-vendor.ps1` derives two equal standard baselines, derives the CGO supplemental closure to a fixed point using CFLAGS/CPPFLAGS, complete roots, recursive includes, authenticated module zip/source equality, and inactive conditional handling, then applies one exact tree-sitter-go transform.
- `verify-go-vendor.ps1` reconstructs the standard/supplemental/final partitions, module/checksum/license coverage, deterministic runs, exact patch pre/postimage, grammar/parser/scanner guards, canonical paths, tree digests, and unchanged `go.mod`/`go.sum`.
- `ensure-ladybug-native.ps1` pins `v0.19.1/windows-x86_64`, accepts exactly three regular repository-owned files, checks exact bytes/SHA, and rejects alternate roots, versions, missing, partial, extra, tampered, or reparse input.
- `package_runtime.go` verifies packaged vendor authority before build, uses explicit `-mod=vendor`, `GOPROXY=off`, `GOSUMDB=off`, distinct lane/cache/runtime roots, pinned Windows native input, staged publication, and vendor-manifest-bound runtime metadata. Stale runtime fallback is rejected.
- `full-build.ps1` verifies vendor before the first Go boundary, isolates home/temp/cache/runtime beneath one repo-local lane, builds through `-mod=vendor`, calls the pinned Windows native authority, stages outputs, runs the launcher, publishes canonical outputs, and runs the canonical analyze.
- `anvien-launcher/build.ps1` invokes direct TypeScript/Vite entries instead of package scripts, uses the fixed PowerShell-5.1-safe relative-path implementation, builds with `-mod=vendor`, and publishes only after staged outputs succeed.
- `qa-child02-p2e-validation.ps1` now resolves its native boundary from durable `third_party/ladybugdb`, not `.tmp\ladybug-native`.

`git diff --check` over non-doc implementation changes exits `0`.

## Vendor and native authority verification

The full `third_party/go-vendor/manifest.v1.json` parsed successfully at `952,465` bytes / `25,634` LF / `0` CR / no BOM, SHA-256 `678A598E0D84B4903DB9684B64F20CB1F273A4785FC65CE9B1718DBC67C7ECDC`. It records:

- schema `1`, closure contract `2`, Go `1.26.3`;
- `45` modules;
- `1,798` sealed files / `190,350,404` bytes;
- `1,578` standard files and `163` supplemental files across `20` roots;
- final determinism `2/2`, differences `0`;
- exactly one reviewed patch;
- license disposition `45` copied / `0` exception / `0` incomplete.

The patch was read in full: `378` bytes, SHA-256 `43BB195FFF439C7DCD0D057094EC1169304553E357F6D0B2781DC5713516BEEA`; it removes only the three guarded absent-scanner include lines. Preimage is `333` bytes / `680E50820F2F7429FEAC88B69DDCEE67735E6E692DFD1F98147109D77D296096`; postimage is `245` bytes / `5A456F352CFEF15430F3FCFB155A3E1BDE38F244E5F8537026553B1F21547713`.

The read-only verifier was run against the same current bytes under both shells during this review:

- PowerShell 7: exit `0`, PASS;
- Windows PowerShell `5.1`: exit `0`, PASS;
- both report `45 / 1,798 / 190,350,404`, tree `7C6D6466D73874684F73C1C35D0D928FA0B729B8C8F59D5F9C71079639C6330E`, final vendor tree `B42427FF27ACD6EC1BA92B756AE6BD438DD409A55CAF536010DCA95AA45FC8E2`, and zero missing/extra/mismatch/reparse/license gap.

Current dependency/native identities independently match:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `go.mod` | `2,392` | `C203E25E0A83A583E89A9D8F8E65BC0881052B3075E46915F90A4FB6433BD249` |
| `go.sum` | `16,065` | `5434F39B08F50F1C53A71333E2E971CFCDE0E80BCA7DD028C00E31B9ACCAEC73` |
| durable `lbug.h` | `79,108` | `3D5114D0863B3DAB3B28BD2FEC97A52E6CF669213739921A01814A5BBF5525EB` |
| durable `lbug_shared.lib` | `32,433,956` | `B18AAFC0B712DC1C4CB9DD25F76C3828282D7D460627980E3A4B16EFCD98A955` |
| durable/canonical `lbug_shared.dll` | `20,230,656` | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |

## Canonical build, runtime, and provenance

The exact canonical PowerShell 5.1 build is current and was not invalidated. Durable Main observation records one successful invocation only, inherited/unmodified `PSModulePath`, unset/default `PSModuleAnalysisCachePath`, Rule 13 holder/write-open clearance, Go `1.26.3`, vendor verification before Go, offline build controls, version `1.2.8`, launcher/Vite `2,943`, and final canonical analyze `2,127/765/0` with graph `122,531/169,182`.

Current publication identities are:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `anvien/bin/anvien.exe` | `73,749,504` | `64EECEBAF459F23EA11DA38B29C5E290F6F6B22D5EE4749E068BB927A8EEDEBB` |
| `anvien/bin/lbug_shared.dll` | `20,230,656` | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| `anvien/bin/anvien-runtime.json` | `225` | `C454EBDF4CD28AD61D0D9A903C0C761F2E135316BE2325E3F59C6A3A45B2C78D` |

The metadata binds the Windows/amd64 Ladybug build to vendor manifest SHA-256 `678A598E0D84B4903DB9684B64F20CB1F273A4785FC65CE9B1718DBC67C7ECDC`. Direct execution returns `1.2.8`.

`READER_GATE_PASS.json` records provenance status `complete`, detailed provenance size/hash `170,548 / 149000ABA76ED45F7D6323D46138522E824B9FF8DC9B2CB928D0931EF5336C69`, source identities `10/10`, native files `3/3`, and publication identities `3/3`. The detailed build file beneath `.tmp` is deliberately not used as authority. Its durable summary is corroborated by the Main-observed build, the current ten source-owner identities, the exact vendor/native trees, and the three current publication identities. All ten recorded source-owner files and every tracked production-like file have last-write times before the accepted build/canonical-analyze cutoff; the six untracked implementation/test/verifier owners also predate it. No source/runtime invalidation is present.

## Reader gate and parity

The durable marker is `READER_GATE_PASS.json`, `6,136` bytes, SHA-256 `2F5341598EDAFDDA4EB3A54698368468CDC921823CB8E8EF247A48F84980A282`, embedded at `2026-08-23T22:06:49.9569993Z` and written at `22:07:39.0868152Z`. It authorizes target access only after these gates:

- package-only outer Windows PowerShell `5.1.19041.6456 Desktop`, inherited module path, unset/default module-analysis cache, exit `0`;
- package verifier `45 / 1,798 / 190,350,404`, zero integrity gaps;
- live load of the canonical DLL with module-observation errors `0`;
- P6-D carrier matrix `3+2+26+1+1`, Graph JSON/Ladybug carrier test `1/1`, tagged native carrier `1/1`, all fallback/failure/skip `0`;
- Child 02 C16 command indexes `3,4,5` PASS `3/3`, remote VECTOR attempts `0`;
- isolated registration plus five built reader commands, every exit `0`, stderr empty, canonical DLL observed, module errors `0`, runtime/graph/Ladybug unchanged;
- native-only graph absent and native Ladybug count `53,156`, equal to Graph JSON `53,156`; split `2,730 standard_library/non_actionable + 50,426 unclassified/review`;
- exact built/native structured sample `1/1`, note byte exact, six flat fields plus seventeen authority fields with differences `0`;
- Child 02 Definition parity `46,593/46,593`, DEFINES pair parity `46,593/46,593`, and all missing/extra/field/pair/endpoint/partial mismatch classes `0`.

Independent reconstruction from the durable raw outputs reproduced the `53,156/53,156` count, the `2,730 + 50,426` groups, and the exact structured `1/1` with zero field/authority difference. Current self Graph JSON and Ladybug still match the gate at `672,905,854 / 40B059AD...` and `388,788,224 / 4A66D159...` respectively.

## Target gate and exact three sites

Ordering is valid and independently checked:

`READER_GATE_PASS` -> target pre-state -> independent semantic oracle -> exactly one normal target analyze -> corrected readers/comparison -> target post-state.

The sole durable target analyze command is `E:\Anvien\anvien\bin\anvien.exe analyze E:\cheapapp.org --force`. It started `2026-08-23T22:13:30.7344357Z`, exited `0`, scanned `1,359`, parsed `887`, failed `0`, and wrote graph `95,762/129,716`. Stdout is `1,277` bytes / `BC7D7F8A86F4769768BDCCE43F2ABEB2F5C722C97AE508494367D660AEEBF714`; stderr is empty. It used the accepted runtime and loaded the canonical DLL with module-observation errors `0`.

The target's actual source bytes independently confirm the zero-based/exclusive coordinates:

| Site | Exact SourceSite range | Current source hash | Analyzer outcome |
|---|---|---|---|
| Promise | `13#3#13#10` | `44467B1D6F1778AE4728B4941136321F7831D8FD67EBE2C93A430A3D132F90AC` | `capability_unavailable / config_topology_unsupported` |
| Math.max | `191#9#191#71` | `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C` | `capability_unavailable / config_topology_unsupported` |
| Math.min | `191#21#191#70` | `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C` | `capability_unavailable / config_topology_unsupported` |

For each site, the built Graph JSON diagnostic, native ResolutionGap row, native `HAS_RESOLUTION_GAP` edge, and aligned oracle have denominator `1/1`; flat outcome/proof/policy/note fields and nested authority identity agree with differences `0`. Native external rows are `0`, resolved/gap overlap is `0`, neither-resolved-nor-gap is `0`, and in-repository analyzer gaps are `0/3`.

The independent TypeScript oracle ran before target analyze, did not read analyzer output, and used the target's TypeScript `6.0.3`, current compiler/config/source hashes, and TypeChecker declarations. It found all three capabilities available and owned by compiler library declarations. P6-A intentionally supports only `target`, `lib`, and `noLib`; the target `tsconfig.json` contains unsupported `include` and `exclude` topology. Therefore the analyzer's structured `capability_unavailable/config_topology_unsupported`, mechanically projected as `standard_library/non_actionable`, is the accepted P6-A capability result rather than an in-repository gap.

## Target and artifact preservation

Current read-only target verification matches the durable post-state:

- HEAD `a869876ab6262dacde6cd5d432d099a91852a646`;
- parent `79ab9b101cc21ec8da79dab724d435e87a6ea6f6`;
- tree `dc024c1d92bdee6211f389a97d0431c267fe5324`;
- index `375,683` bytes / `D2417BF08B7128F6069FF117546E78E527514F98354DCADA32CC6164F60B1C71`;
- branch/status, cached diff, and worktree diff match after CRLF/LF normalization;
- source pre/post/current `203/203`, differences `0`;
- config pre/post/current `3/3`, differences `0`;
- `.anvien/settings.json` unchanged; only `graph.json`, `lbug`, and `meta.json` changed under the allowed analyzer-owned output root.

All `476` discovered path/bytes/SHA reference rows across the parseable durable JSON/matrices were checked with hash caching: `473` point to current byte-identical files and the remaining `3` are the expected historical pre-analyze `graph.json`, `lbug`, and `meta.json` identities superseded by the one authorized analyze. There are no missing, forbidden-C data, relative, or unexplained mismatched references. Core JSON documents decode strictly. CLI output artifacts containing a JSON document followed by the product's `--- Next` text are matrix-hashed and were parsed by isolating that documented JSON prefix; they are not mistaken for standalone strict JSON.

## Superseded artifacts and procedure

Two retained target artifacts are explicitly negative/superseded and are not acceptance inputs:

1. `target-exact-reader-matrix.json` initially queried Promise through end column `42`; direct source and corrected matrices prove the exact token range ends at `10`. `target-corrected-site-matrix.json`, built explain, native gap/edge rows, aligned oracle, and final reader parity all use `13#3#13#10`.
2. `target-three-site-parity-matrix.json` attempted to read the `840,614,023`-byte Graph JSON into one Node string and failed with `ERR_STRING_TOO_LONG`; stdout is intentionally zero bytes and exit is `1`. The later streamed/native/built-reader comparison is the accepted chain and independently reproduces `3/3` with zero differences.

The raw semantic oracle initially used a whole type-reference range ending at `42`; its pre-analyze semantic classification remains independent. The later aligned oracle corrects only the source-site token range from current source/AST and does not replace the pre-analyze semantic finding.

Other retained harness incidents include module-path/cache contamination attempts, `$HOME` and wrapper syntax failures, wrong-shell package invocation before native load, two pre-reader registration attempts, capture/alias/suffix parsing failures, and a read-only Git tree quoting error. None mutated target source/config/index, produced an accepted runtime, changed accepted output, enabled network/install/package scripts, or supplied positive evidence. The Coder narrative omitted the `ERR_STRING_TOO_LONG` incident from its prose list, but the failure matrix/stderr are durable; this report dispositions it explicitly rather than hiding it.

Procedure is **CLEAN** on actual boundary impact:

- no C-drive workspace, project data, cache, or write was used;
- `C:\Program Files\nodejs\node.exe` and Windows PowerShell 5.1 are installed OS executables, not C-drive workspaces or evidence roots. Treating executable dispatch as forbidden would contradict the mandated PS5.1 gate itself; all oracle compiler/source/config/output bytes are on E and no C data artifact was accepted;
- no implicit network, install, package script, shared/global/protected/quarantined cache use, alternate worktree, stage, commit, push, reset, checkout, stash, rename, or broad process control occurred in the accepted continuation;
- literal `E:\Anvien\.tmp` is disposable scratch only and supplies no evidence, review input, handoff dependency, source of truth, or procedure input;
- the preserved `Microsoft\Windows\PowerShell\ModuleAnalysisCache` artifact remains at its pre-review `4,641` bytes and `2026-08-23T20:27:28.1179411Z` last-write time;
- this Supervisor ran only the two mandatory read-only vendor verifiers; it did not rerun build, reader, self-analyze, or target analyze.

## Blast radius and residual surfaces

Fresh pre-edit Anvien evidence remains applicable because the relevant source bytes are unchanged:

- `internal/graphhealth/diagnostics.go` containing-file impact: CRITICAL, `107` impacted symbols / `31` files / `34` direct / `1` flow / `19` tests; file detail HIGH `154/48/70/133/186/1/19`.
- `validStructuredResolutionAuthority`: LOW `3/1/1/0`, which does not erase the containing-file blast radius.
- `package_runtime.go`: CRITICAL `19` impacted symbols / `7` files / `14` direct / `1` process / `2` linked tests.
- `buildGoRuntimePackage`, `prepareGoSourcePackage`, and the prior native resolver are CRITICAL at the recorded process surface.
- PowerShell owners report LOW/zero extracted graph impact because the extractor emits no symbols/edges; direct source call/root flow, not the zero graph row, establishes their real ownership.

CRITICAL/HIGH are retained scope warnings, not edit bans. Current source, tests, full build, runtime, native, package, reader, and target boundaries cover the changed invariant.

Residual same-invariant surfaces are explicit:

1. Committed P6-C3 `validateResolutionOutcome` still compares nested authority to outer outcome only over source site and stage, not the six additional identity groups. That known omission is outside P6-D edit authority and remains separately routed to the Owner. It does not invalidate this P6-D verdict because all current target carriers are exact and the P6-D graph-health consumer independently rejects any drift before policy projection.
2. `scripts/ensure-ladybug-native.sh` retains the pre-existing non-Windows auto/download path. It is not used by the accepted Windows `v0.19.1/windows-x86_64` canonical build, was outside the authorized Windows recovery owners, and must not be represented as cleared for Linux/macOS. The Windows `.ps1`, Go runtime Windows branch, launcher, full build, and Child 02 QA consumer all use the exact durable bundle with no `.tmp\ladybug-native` or network fallback.
3. The detailed build-provenance file lived under disposable `.tmp` and is not acceptance authority. Acceptance relies on its durable marker summary plus independent Main observation and current source/native/publication reconstruction. No claim depends on preservation of `.tmp` bytes.
4. Final current-byte Anvien `detect-changes`, planner/ledger transition, exact staging, and isolated commit are Main-owned closure steps after this report. They are not silently counted as completed here.

## Repository Git boundary

Current repository truth at final pre-report capture:

- branch `master`;
- HEAD `741335f7efa5ad215ec28361d88c462f431565ab`;
- parent `838f379ab00c93dfe7507603f6608bb08d61f038`;
- tree `65c4be5d1aeab6cad6faa0f3d9a70b93c76bbda8`;
- staged entries `0`;
- tracked changes `13`;
- untracked rows `1,893` after a concurrent protected Main rotation handoff appeared; that report is unrelated to the P6-D implementation candidate and remains preserved;
- Owner deletion `CONTRIBUTING.md` remains exactly present and untouched.

The implementation boundary remains the graph-health projector/test, exact vendor/manifest/licenses/one patch/verifier/materializer contracts, exact durable Windows native bundle and consumers, package/runtime tests, launcher/full-build/Child 02 QA owners, durable Coder/QA/Supervisor evidence, and Main-owned ledgers. The roadmap/ledgers, Owner deletion, Main handoffs, Microsoft artifact, unrelated reports, and protected work are not normalized into the candidate. Nothing was staged or committed by Coder or Supervisor.

## Handoff

P6-D is accepted by independent Supervisor review. Main `01a030b6-9de3-7c72-977b-31d546108138` is responsible for verifying this report seal, recording the PASS in the four ledgers/roadmap with the planner workflow, running current-byte detect-changes, resolving the exact accepted manifest, and creating the isolated P6-D commit under its own authority. P6-E remains locked until that accepted P6-D boundary is committed; this report grants no P6-E implementation authority.
