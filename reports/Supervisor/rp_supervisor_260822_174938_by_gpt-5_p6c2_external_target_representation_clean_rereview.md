# Supervisor Clean Re-review: Child 06 P6-C2 External Target Representation

Verdict: **REJECT**

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260822_174938_by_gpt-5_p6c2_external_target_representation_clean_rereview.md`.
- Review date: `2026-08-22 +07:00`; clean re-review began before the current runtime evidence and report creation was anchored at `17:49:38 +07:00`.
- Reviewer: `gpt-5`, the existing independent Child 06 Supervisor lane.
- Repository: `E:\Anvien`, branch `master`.
- HEAD: `8055f0a6860721e26462572e34469e0d708d4a52`.
- Parent / accepted P6-B commit: `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- Open slice reviewed: P6-C2 only. P6-C3, P6-D, graph-health policy, target validation, project/package declaration lookup, and `E:\cheapapp.org` remained locked.
- Next owner: Main task `01a028ce-9460-7282-ad6e-5a3dc75f972a`, or its officially transferred successor.

## Claim And Authority

Claim reviewed: the unchanged 22-path P6-C2 candidate materializes only referenced final resolved P6-B authority rows as deterministic `ExternalSymbol` facts, emits no fake target for unavailable/excluded/mismatch rows, preserves complete per-site provenance through the graph and affected readers, and never pollutes repository ownership.

Authority read in full before judgment:

1. Root `AGENTS.md`.
2. Repo-local `working-rules`, `supervisor`, `Data-Integrity`, and `Edge-Case` skills.
3. All four current Child 06 ledgers.
4. Current coder report `reports/coder/rp_coder_260822_162635_by_gpt-5_p6c2_external_target_representation.md`.
5. Immutable prior report `reports/Supervisor/rp_supervisor_260822_171803_by_gpt-5_p6c2_external_target_representation.md`.
6. Current source bytes, Git boundary, packaged artifact, fixture graph, native Ladybug, MCP readers, root graph, and detect reality.

Reports and prior test/build claims were treated as evidence inputs, not acceptance. The Data-Integrity and Edge-Case lenses were applied inside this Supervisor review: owner separation, referenced-only scope, duplicate/replay determinism, conflicting identity fail-closed behavior, persistence losslessness, process isolation, and non-editable external semantics remained the exact invariant family.

## Immutable Input Seals

The coder report was independently reproduced:

- `12,480` bytes / `198` LF / `0` CR;
- strict UTF-8, no BOM;
- SHA-256 `08BCEA399B3ECEED4721858176E3845B19A67DB376FD21138E8B9A017AEB05B0`;
- createdAt = lastWrite `2026-08-22T16:28:50.6827037+07:00`.

The prior P6-C2 report was independently reproduced:

- `18,318` bytes / `229` LF / `0` CR;
- strict UTF-8, no BOM;
- SHA-256 `C1F7C1959C3BB36CA1F91817FC5FDCF8DCFD193388760394BEB59325D597FA88`;
- createdAt = lastWrite `2026-08-22T17:20:18.7964790+07:00`.

## Exact Candidate Identity And Reuse Gate

The initial clean snapshot independently matched the assignment exactly:

- status `58` = `17` tracked + `41` untracked;
- candidate `22/22` = `9` production + `8` tests + `4` ledgers + coder report;
- candidate storage `17` tracked + `5` untracked;
- outside candidate `36` = `30` pathname-only Main handoffs + `3` prior P6-B Supervisor reports + `2` older coder reports + the immutable first P6-C2 report;
- index `0`;
- `git diff --check` exit `0` with empty output;
- missing/extra candidate paths `0/0` and unexpected outside paths `0`.

All nine production identities exactly match the immutable prior report:

| Production owner | Bytes | SHA-256 |
| --- | ---: | --- |
| `internal/scopeir/kinds.go` | 5,253 | `CAC3AD018B82E7B2950D0F13E24793D3CD84F708CF8E9533198C74C749B93B6E` |
| `internal/resolution/external_symbol.go` | 11,845 | `275683BD680107A19DCB40B782566C087A793A64CAD0341E040A43DBEBBE623C` |
| `internal/resolution/emit.go` | 26,822 | `DE88FA7717B3E1BBE9AB099616B5D43095A91C86C22706F17F0F5A5CD793C078` |
| `internal/resolution/resolve.go` | 38,706 | `7D1394E661E78679602429212E97F29A605147AFD2DD894E6864A3D15A808A69` |
| `internal/lbugload/csv.go` | 24,986 | `1B7D38FDE913A6A0F05033947B069650C01B8D2F55F64FDDD1D9DC406D894BCA` |
| `internal/lbugschema/schema.go` | 20,174 | `AF0B5727C9261FDB542B94A7DC1383022EA4F569F85FDEC8C816B1CC9B4F23FA` |
| `internal/processes/processes.go` | 14,942 | `F2DE1EB156D118075A3E5D6F62DE926F18E4F13B59000807ACF49F1E85DBD1B9` |
| `internal/mcp/context.go` | 22,267 | `3A0397ACF3A9E45771E4E49A1588AE634E331ADAC2DA464E295D8C6424A50528` |
| `internal/mcp/rename.go` | 8,945 | `B4EE673CFC1B01427B7D309BC1173DEB98D32FC7CFEB652B81E285E5822A2B14` |

The eight current test identities were independently recorded:

- `internal/resolution/p6b_tsstdlib_test.go`: `DFBEDC54BC2B8325ADDF459ECD2ED094A2E4E80A516774207C731BDC6745DB9F`;
- `internal/resolution/p6c2_external_symbol_test.go`: `9CE625B08E19B7499D102CB9DCF33A00E731EAE92145DAC77B71E8D0D7B1E11E`;
- `internal/analyze/p6b_tsstdlib_test.go`: `9A16673AC92B9F4929F2357841641282408A7EBC615A6D1FF65D58BC846BA864`;
- `internal/lbugload/csv_test.go`: `4829048837DE2FA9FAFBA8A66DF5C8DEF50A3085FCB6C1DC2CF3E44E0486A759`;
- `internal/lbugschema/schema_test.go`: `910181E1958A700B94C3A46737E1E578ABA3A5C54219ACE993A603D37AAB66DA`;
- `internal/processes/processes_test.go`: `7855D6FC7265E5B0B1FF42B3B4B496E2E039D10BEDB686C4300D2A582FD61683`;
- `internal/mcp/p6c2_external_symbol_test.go`: `5EE1AF26B0FAFFE756C64CA913927FC69624CA62288889B1BD254F13814AF96C`;
- `internal/lbugnative/p6c2_external_symbol_integration_test.go`: `2C1631DD0A6EE65679A954DAC049E0641AB712C05C137D7EA4BF203CE582DDD8`.

The canonical current 22-path manifest fingerprint, formed from sorted relative path, byte count, and SHA-256 rows, is `C1F469EDA7DF75B6A8ECCE6A086E585074EC1A3C96535758308E74F532C53937`. Every candidate last-write timestamp is at or before the prior immutable report seal. The four ledger hashes are `24F318B0...99D6`, `0B474821...BD9`, `CDC61C7D...D5BA`, and `EA77CFDA...34EF` for plan/evidence/benchmark/actual-status respectively. No package manifest, lockfile, `go.mod`, `go.sum`, generated catalog, P6-C3/P6-D owner, or target path appears in the candidate.

This proves the source/test/coder/ledger candidate bytes and exact path set used by the prior technical review have not changed. Reuse of its same-invariant technical clearances is therefore authorized.

## Preserved Technical Clearances

Residual P6-C2 code invariant blockers established by either review: none.

The following prior source clearances remain exact on current hashes:

- `internal/scopeir/kinds.go` adds only the dedicated external label.
- `internal/resolution/external_symbol.go` performs referenced-only final-result materialization, stable semantic-ID grouping, endpoint/site validation, deterministic sorting, identical replay coalescing, and conflict rejection.
- `internal/resolution/emit.go` carries transient endpoint inventory and preserves coalesced relationship proof rows.
- `internal/resolution/resolve.go` records eligible endpoint facts and invokes materialization only after final P6-B validation; unavailable, excluded, and mismatch rows never produce a target.
- Ladybug CSV/schema owners preserve fixed columns, table, endpoint pairs, and relationship proof fields without fallback loss.
- Process extraction excludes a CALLS relation when either endpoint is external.
- MCP context/impact preserve exact provenance with empty repository path and explicit `external=true`, `editable=false`, `repositoryOwned=false`.
- Rename rejects the external symbol before old-name comparison or edit collection.
- Graph JSON remains a generic transport. `impact.go`, HTTP/Web, file-detail, graph-health, the shared P6-C3 DTO, target behavior, public traversal API, project/package lookup, synthetic `IMPORTS`, runtime Node/`tsc`, network, install, and package scripts remain untouched.

All eight changed/new test owners retain direct coverage for resolved/unavailable separation, deterministic identity, identical replay, site-proof union, conflict rejection, CSV/schema parity, native readback, process exclusion, MCP provenance, and pre-edit rename rejection.

## Clean E-rooted Evidence Chain

Before Git, Go, Anvien, or packaged runtime helpers, this review created one tree under `E:\Anvien\.tmp\supervisor-p6c2-clean-rereview` and assigned these process variables without reading their defaults: `ANVIEN_HOME`, `HOME`, `USERPROFILE`, `GOPATH`, `GOMODCACHE`, `GOCACHE`, `GOTMPDIR`, `TMP`, `TEMP`, and `XDG_CONFIG_HOME`. Git global/system configuration was disabled or redirected to the same E tree. `GOPROXY=off` and `GOSUMDB=off` were set. All assigned paths were independently normalized and verified under `E:\Anvien` before helper use.

The new E-only module cache contained `0` entries. Root `go.mod` requires multiple external modules, including Cobra, tree-sitter packages, and Hugot. A fresh compile/test could therefore not be supported without an external module source. Per the clean re-review authority, Go compile/tests were not run and no fallback to an external cache, install, proxy, network, or package script occurred. This does not narrow the unchanged technical clearances because the production/test/coder manifest is byte-identical, the current packaged artifact is exact, and the affected packaged boundaries below were executed independently.

The packaged binary identity is unchanged:

- path `E:\Anvien\anvien\bin\anvien.exe`;
- `73,653,760` bytes;
- SHA-256 `7B14655DBB65472CE83FBAF0677294CCB235EC09A34DB3AB877F5C5C1B12F568`;
- lastWrite `2026-08-22T16:04:06.1123976+07:00`.

An E-only `PATH` diagnostic reached loader exit `-1073741515` before help output, matching the already disclosed native dependency-chain baseline. It is not counted as a runtime success. Packaged help then exited `0` with all application home/module/cache/temp values still explicitly E-rooted; the default PATH value was neither read nor printed.

## Packaged Runtime, Persistence, And Reader Evidence

The resolved TypeScript fixture was analyzed twice from clean force runs. Both graph artifacts were byte-identical:

- `11` nodes / `10` relationships / `32,666` bytes;
- SHA-256 `CB308956ECD89DDB801BC37ED72D01AD73656D2D01EDC9C49C3A8ADB0CC7920A`;
- `3` referenced external nodes;
- `4` coalesced external edges;
- `6` exact `typescript-authority-result-v1` proof rows;
- node source-site counts `[2,1,3]`;
- external File/DEFINES/IMPORTS/CONTAINS ownership pollution `0`.

The noLib fixture exited `0` with `5` nodes / `3` relationships, `0` external nodes, and `0` external edges. Its graph was `4,729` bytes with SHA-256 `6904164C79643B61CD53696CA8034217FAB12F81AC82C7AFCF11A0EA41CBA38D`.

The compiler-default fixture exited `0` with `6` nodes / `5` relationships and materialized only one external `Promise` type node and one edge. Its graph was `12,935` bytes with SHA-256 `F7346C71EF87BA12592590A828B43DE687CFE25142ED31498EE3AD0B0DA16254`. Excluded/mismatch value paths created no fake target.

After a fresh resolved-fixture force analyze, native Ladybug Cypher returned exactly four `CodeRelation` rows targeting `ExternalSymbol`: `Math.max` CALLS/USES and `Promise` CALLS/USES. Every row retained `typescript_standard_library`, TypeScript `5.9.3`, `ready`, `typescript-standard-library-authority`, exact edge site count, `external=True`, `editable=False`, and `repositoryOwned=False`. This is native fixed-schema/readback evidence, not graph-JSON substitution.

Packaged context after its own fresh force analyze selected `Math.max` as `ExternalSymbol`, returned empty `filePath`, complete semantic owner/catalog/profile/config/declaration/source-site facts, two incoming relations with lossless authority proofs, and `processes=[]`.

Packaged upstream impact after its own fresh force analyze returned one direct repository caller, `affected_modules=[]`, `affected_processes=[]`, exact external provenance, and zero resolution-health gaps. This independently closes process-membership exclusion at the reader boundary.

Packaged dry-run rename after its own fresh force analyze returned `status=rejected`, reason `external_symbol_non_editable`, and no edit collection.

## Current Root Graph And Detect Evidence

Repo-local packaged help and important analyze/detect flags were read before graph use. Exactly one fresh explicit-path root `analyze E:\Anvien --force --json` was run for current root evidence. It exited `0`:

- `2,042` scanned / `756` parsed code / `0` failed;
- `116,698` nodes / `162,222` relationships;
- `19,658` dependency edges;
- graph path `E:\Anvien\.anvien\graph.json`.

The immutable prior review recorded `2,041 / 756 / 0 / 116,684 / 162,208`. Current `+1 file / +14 nodes / +14 relationships` is the later immutable prior P6-C2 report becoming graph input; it is time-bounded evidence and was not normalized away.

Current explicit-path `detect-changes --repo E:\Anvien --scope all --json` exited `0` with the full CRITICAL warning preserved:

- affected `44` symbols / `17` files;
- changed `202` symbols / `17` files;
- affected layers `api=16`, `backend=28`;
- affected areas `mcp=16`, `mixed=7`, `resolution=19`, `storage=2`;
- changed layers `api=11`, `backend=119`, `backend_test=53`, `docs=19`;
- changed areas `analyzer=32`, `documentation=19`, `mcp=11`, `providers=38`, `resolution=55`, `storage=47`;
- ResolutionGap delta `33` entities / `33` occurrences;
- actionability `analyzer_gap=20`, `non_actionable=13`;
- classifications `builtin=12`, `in_repo_unresolved=20`, `standard_library=1`;
- fact families `access=14`, `call=14`, `type-reference=5`;
- current Resolution Health total `0`;
- semantic app-layer and functional-area fields/sources complete for all `116,698` nodes.

CRITICAL is a blast-radius warning. It is neither hidden nor used as an automatic defect or acceptance substitute.

## Build, Regression, And Benchmark Disposition

- Fresh compile/tests: not run because the clean E-only module cache was empty and the repository requires external modules; no forbidden fallback occurred.
- The exact packaged artifact independently executed every affected runtime/persistence/reader boundary listed above.
- Known broad repository facts remain unchanged and are not relabeled: broad `go build ./...` is unsuccessful on intentional language fixtures; broad `go test ./internal/...` retains seven repo-local Git/temp environment failures plus the pre-existing C# and Dart ACCESS parity baselines.
- Official lifecycle/package scripts, network, install, target, stage, commit, push, branch, UI/Playwright, and subagent work were not run.
- Packaged size delta over accepted P6-B remains `+39,936` bytes (`+0.05425%`) with no Owner threshold.
- Current product/runtime observations are the `3 / 6 / 4 / 6` node/site/edge/proof inventory, zero ownership/process pollution, byte-identical replay, exact native rows, and exact binary identity.

## Sole Blocking Invariant

The candidate implementation remains technically clear. Acceptance is blocked solely because the mandated no-C execution boundary was invalidated by the Supervisor tool substrate during cleanup.

Concrete event:

1. After all technical/runtime/root/detect evidence had completed, the review issued a cleanup command whose operands and validation were exclusively the four named E-side review artifact roots.
2. The command was rejected before execution by tool policy.
3. The tool error itself disclosed and attempted to launch its shell as `C:\Program Files\PowerShell\7\pwsh.exe`.
4. The assignment's stop condition is unconditional when any command exposes or attempts a C path. The review therefore stopped acceptance work immediately, did not inspect or repair that path, and selected this procedural verdict.

This is not a code, test, ledger, packaged artifact, graph, persistence, or reader failure. No coder repair is requested. The issue is the review execution substrate: even an E-only command payload can have its launcher path exposed by the tool layer when policy rejects the call.

Required re-review direction:

- use a Supervisor execution substrate whose command launcher and error surface satisfy the literal no-C rule; or obtain an Owner clarification explicitly excluding the host shell/OS runtime launcher from the repository-data boundary;
- if candidate bytes/path set still match the fingerprint above, retain every technical clearance and review only the procedural boundary;
- repeat only the proportionate E-rooted packaged/runtime/root/detect/final-boundary chain needed to prove the clarified rule;
- do not route to coder unless a new code invariant is independently established.

## Cleanup And Final Pre-report Boundary

All review-created E-side artifacts are absent:

- `E:\Anvien\.tmp\supervisor-p6c2-clean-rereview`;
- resolved fixture `.anvien`;
- noLib fixture `.anvien`;
- compiler-default fixture `.anvien`.

The root `E:\Anvien\.anvien` remains intentionally present as the authorized current graph output. No protected history was cleaned.

At the final pre-report snapshot, one concurrent protected Main handoff had appeared since the initial snapshot. Only its pathname class was observed. It is not candidate drift:

- branch `master`, HEAD `8055f0a...`, parent `5c1584f...`;
- status `59` = `17` tracked + `42` untracked;
- unchanged implementation candidate `22/22`, missing/extra `0/0`;
- protected/current history outside candidate `37` = `31` pathname-only Main handoffs + `3` prior P6-B Supervisor reports + `2` older coder reports + the immutable first P6-C2 report;
- index `0`;
- `git diff --check` exit `0`;
- P6-C2 remains `[ ]`, unstaged, and uncommitted;
- P6-C3/P6-D and target remain locked.

This report is the only authorized durable file created by the clean re-review. Its external byte/LF/CR/UTF-8/BOM/SHA/timestamp seal is measured after write and is not embedded self-referentially.
