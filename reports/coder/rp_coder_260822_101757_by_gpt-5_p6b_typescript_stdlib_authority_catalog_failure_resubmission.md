# Coder Resubmission: P6-B Catalog-Validation Failure Carriage Residual Repair

Status: READY_FOR_INDEPENDENT_SUPERVISOR

This is coder evidence, not acceptance. P6-B remains unchecked, unstaged, uncommitted, and blocked from P6-C1/P6-C2/P6-C3/P6-D until an independent Supervisor decision and Main action.

## Metadata

- Created: `2026-08-22T10:17:57.0423971+07:00`
- Coder lane: Child 06 P6-B only
- Checkout: `E:\Anvien`
- Branch: `master`
- Authoritative HEAD: `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`
- P6-B implementation base / HEAD parent: `b98131e44932a7bcac17b487ecb2914535927d01`
- Residual-review authority: `reports/Supervisor/rp_supervisor_260822_090533_by_gpt-5_p6b_typescript_stdlib_authority.md`
- Residual-review verdict: `REJECT`
- Residual-review seal: `20,729` bytes / `168` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `C373D2413F1D60082904729F5A536B09D588A5D1DABE12181E454BCE2AD3209A`
- Report identity protocol: this file is written once and externally sealed after creation. Its final bytes, LF/CR counts, UTF-8/BOM state, SHA-256, and filesystem createdAt are supplied in the terminal handoff; no self-hash is embedded.

## Scope And Verdict Boundary

The repair changes only the residual P6-B invariant: exact typed catalog-validation failure reason plus lossless per-source-site capability carriage. The second independent review already cleared canonical identity, receiver terminality, the exact ten-vector compiler matrix, and valid-catalog site carriage; those paths were preserved and only regression-tested.

No CLI projection, graph persistence, graph-health, `ExternalSymbol`, shared P6-C3 final DTO, project/package lookup, synthetic `IMPORTS`, target access, or P6-C/P6-D code was opened. No network, install, package-script execution, stage, commit, push, or subagent occurred. No alternate checkout or target repository was accessed. Required doctor/process diagnostics named installed processes, but no filesystem operation read or wrote outside `E:\Anvien`.

## Exact Defect And Repair

The rejected path had four linked defects:

1. `NewAuthority` retained the typed loader reason inside an unavailable profile.
2. `Authority.availability` collapsed every nil catalog to `catalog_missing`.
3. `baseResult` did not carry explicit rejected/missing provenance semantics.
4. `validateTypeScriptAuthorityResult` required ready-catalog authority/logical/artifact/config hashes for every status and converted a handled catalog failure into an analysis error.

The repair introduces `CatalogProofState` with a fail-closed contract:

- `ready`: validated authority hash, logical catalog hash, and artifact hash are all mandatory.
- `missing`: all three hashes are forbidden; no missing asset identity is fabricated.
- `rejected`: authority and logical catalog hashes are forbidden; only SHA-256 of the attempted non-empty artifact bytes is retained.

Every state retains the expected authority kind, exact TypeScript `5.9.3` runtime contract, profile hash, and configuration/inventory proof hash. `Authority.availability` returns the exact typed profile reason before the nil-catalog fallback. The source-site DTO carries the proof state unchanged, and the resolver validator enforces every state/reason/hash combination before accepting the record.

`NewAuthorityFromCatalog` validates only the same compact offline catalog DTO and supplies the existing resolver/analyzer injection boundary for exact failure tests. Production `NewAuthority` still consumes only the checked-in `go:embed` artifact. No raw declaration parsing or alternate runtime discovery path was added.

## Exact Failure Matrix

The durable table names exactly six classes:

| Row | Accepted reason | Proof state | Attempted artifact hash |
| --- | --- | --- | --- |
| `empty-missing` | `catalog_missing` | `missing` | absent |
| `schema` | `catalog_schema_unsupported` | `rejected` | present |
| `version` | `catalog_version_mismatch` | `rejected` | present |
| `input-manifest` | `catalog_input_manifest_mismatch` | `rejected` | present |
| `logical-hash` | `catalog_hash_mismatch` | `rejected` | present |
| `trailing-artifact-integrity` | `catalog_input_manifest_mismatch` | `rejected` | present |

Authority tests assert every row through value `parseInt`, type `Promise`, and member `Math.max` lookups. Direct resolver tests duplicate all three facts and produce exactly three unique records/counters per row, zero generic gaps, identical replay, and conflict rejection.

Built `analyze.Run` tests identify exactly these seven stable source sites for every row:

- `SourceSite:main.ts#call#Math.max#2#18#2#44`
- `SourceSite:main.ts#call#Promise#3#9#3#59`
- `SourceSite:main.ts#call#resolve#3#42#3#58`
- `SourceSite:main.ts#type-reference#Math.max#2#8#2#44`
- `SourceSite:main.ts#type-reference#Promise#1#27#1#34`
- `SourceSite:main.ts#type-reference#Promise#1#45#1#52`
- `SourceSite:main.ts#type-reference#Promise<number>#1#25#1#42`

Each built run has seven capability records, seven unavailable counters, zero generic unresolved/gap counters, no duplicate site ID, and a deeply identical second-run record slice.

## Fresh Anvien Preflight And Blast Radius

Repo-local packaged help and important impact flags were read before graph work. Fresh pre-edit analyze PASS:

- `2,017` scanned
- `752` parsed code
- `0` failed
- `115,942` nodes
- `160,978` relationships

File-level upstream blast radius:

- `internal/tsstdlib/catalog.go`: CRITICAL, `176` impacted, `24` affected files, `43` direct, `328` contained symbols.
- `internal/resolution/resolve.go`: CRITICAL, `29` impacted, `5` affected files, `24` direct, `180` contained symbols.
- `internal/resolution/types.go`: CRITICAL, `116` impacted, `13` affected files, `15` direct, `92` contained symbols.

Symbol-level upstream blast radius:

- `NewAuthority`: CRITICAL `5 impacted / 1 direct / 4 modules / 31 processes`.
- `Authority`: CRITICAL `100 / 11 / 5 / 70`.
- `LookupResult`: CRITICAL `15 / 10 / 3 / 35`.
- `TypeScriptAuthorityResult`: CRITICAL `95 / 6 / 5 / 70`.
- `recordTypeScriptLookup`: CRITICAL `6 / 3 / 2 / 35`.
- `validateTypeScriptAuthorityResult`: CRITICAL `4 / 1 / 2 / 23`.
- `Authority.availability`: LOW, `5` impacted / `1` direct.
- `Authority.baseResult`: LOW, `5` impacted / `2` direct.

Test-owner file impacts with tests were MEDIUM `7/7` for catalog, MEDIUM `8/8` for resolution, and LOW `2/2` for analyze. Counter-equality helper impacts were MEDIUM `8/7` and LOW `2/2`. HIGH/CRITICAL values were treated as blast-radius warnings and preserved in full; they were not edit bans.

## Build, Package, And Runtime Evidence

Rule 13 preflight reported the analyze lock free/nonexistent. Restart Manager identified the real packaged destination holders:

| PID | Process | Role | Ownership | Final state |
| ---: | --- | --- | --- | --- |
| `16968` | `anvien.exe` | `mcp` | editor-owned | terminated, verified gone |
| `8540` | `anvien.exe` | `mcp` | editor-owned | terminated, verified gone |
| `5204` | `anvien.exe` | `mcp` | editor-owned | terminated, verified gone |
| `18124` | `anvien.exe` | `mcp` | editor-owned | terminated, verified gone |
| `19056` | `anvien.exe` | `mcp` | editor-owned | terminated, verified gone |

After exact termination, Restart Manager returned no holder and exclusive-open on `E:\Anvien\anvien\bin\anvien.exe` succeeded. No broad process kill was used.

Build results:

- `go build ./cmd/... ./internal/...`: PASS.
- `go build ./...` in `anvien-launcher/src`: PASS; exact emitted wrapper later cleaned.
- `go build ./...` in `anvien-launcher/server-wrapper`: PASS; exact emitted wrapper later cleaned.
- Broad `go build ./...`: exit `1` only on intentional non-buildable fixtures: unresolved fixture imports, mixed packages, and standalone C source without cgo. It is not labeled PASS.
- Official lifecycle script: not run because it performs prohibited install/package scripts.
- A repo-local wrapper discovery probe exited `1` before any build output and is not counted as evidence.
- Direct source invocation of hidden `package build-runtime` with cached Ladybug `0.19.1`: PASS; wrote the real destination.
- `package ensure-runtime`: PASS, Windows/amd64.
- Packaged `--help`: PASS.
- Packaged real fixture analyze: PASS, `2` scanned / `1` parsed code / `0` failed / `8` nodes / `6` relationships. Its `.anvien` was deleted by exact `anvien clean --force`.

Final packaged artifact:

- Path: `E:\Anvien\anvien\bin\anvien.exe`
- Bytes: `73,610,752`
- SHA-256: `E7F851FCEC43F7CB740707FF8BBE8CB89E8B741A49FB37E72ADFBF2332182D42`
- Go: `1.26.3`
- Build: `ladybugdb`, trimpath, `CGO_ENABLED=1`
- VCS: `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`, modified=true
- Delta from accepted P5-D `71,509,504`: `+2,101,248` bytes (`+2.9384%`)
- Delta from prior rejected repair `73,605,632`: `+5,120` bytes (`+0.0070%`)

## Tests And Regression

New exact focused gates:

- `go test ./internal/tsstdlib -run '^TestCatalogFailuresPreserveExactLookupReasonAndProofAbsence$' -count=1`: PASS (`1.435s`).
- `go test ./internal/resolution -run '^TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite$' -count=1`: PASS (`0.197s`).
- `go test ./internal/analyze -run '^TestP6BRunCarriesCatalogValidationFailuresPerSourceSite$' -count=1`: PASS (`1.546s`).

Full affected packages:

- `internal/tsstdlib`: PASS (`2.307s`).
- `internal/resolution`: PASS (`0.579s`).
- `internal/analyze`: PASS (`4.766s`).

These runs preserve the cleared exact ten-vector compiler matrix, canonical ID inventory/component rejection, valid-catalog site/counter equality, receiver call/access terminality, repository/import precedence, config/profile behavior, and TypeScript-only isolation.

Fresh broad `go test ./cmd/... ./internal/... -count=1` ran to completion. P6-B packages and `internal/cli` PASS. Only two preserved provider parity baselines remain:

- C#: expected `ACCESSES=2`, got none.
- Dart: expected `ACCESSES=2`, got `1`.

Three earlier environment-limited CLI failures did not reproduce in this fresh run; the prior results remain historical and are not rewritten as successes. No baseline test was edited.

## Benchmark And Artifact Evidence

Fresh Windows/amd64, i7-3770, three-run measurements:

| Benchmark | Time | Bytes/op | Allocs/op |
| --- | --- | --- | --- |
| embedded catalog load | `156.659-164.050 ms/op` | `36,350,514-36,453,763` | `291,808-291,820` |
| global lookup | `247.4-255.0 ns/op` | `192` | `1` |
| inherited-member lookup | `621.3-708.0 ns/op` | `216` | `4` |

No Owner threshold exists, so these are measurements rather than invented acceptance gates. Build/test durations remain validation evidence.

Generator and checked-in catalog bytes did not change in this residual repair. Therefore the accepted two clean byte-equal generations, `2,003,050`-byte artifact, logical hash `dca7af22ff26510cf9075fcb587ab76649bf169edcb0001e50c234d89c9dbb0b`, artifact SHA-256 `F188D15A5D91925DF3E724CBAB97964813E3F6DFD9DF7408FDC7B92EA4CEA487`, `100` inputs / `3,141,835` bytes, and `14,587/14,587` unique semantic IDs remain authoritative. No unnecessary generator rerun occurred.

## Final Graph And Detect Evidence

Final packaged self-analyze PASS:

- `2,023` scanned
- `752` parsed code
- `0` failed
- `116,099` nodes
- `161,226` relationships

Required explicit-path `anvien detect-changes --repo E:\Anvien --scope all --json` exited `0` with CRITICAL risk:

- Affected: `35` symbols / `8` files.
- Changed: `307` symbols / `8` files.
- Affected layers: `backend=31`, `mixed=4`.
- Affected areas: `analyzer=8`, `mixed=9`, `resolution=18`.
- Changed layers: `backend=288`, `docs=19`.
- Changed areas: `analyzer=28`, `documentation=19`, `resolution=260`.
- Resolution-gap delta: `177` entities / `180` occurrences.
- Actionability: `analyzer_gap=125`, `non_actionable=52`.
- Classifications: `builtin=30`, `in_repo_unresolved=125`, `standard_library=22`.
- Families: `access=94`, `call=71`, `type-reference=12`.
- Functional areas: `analyzer=14`, `resolution=163`.
- Semantic app-layer and functional-area fields: complete for all `116,099` nodes, no missing fields/sources.
- Current health total: `0` ResolutionGaps.

No count or warning was normalized or reduced.

## Cleanup

Twenty exact paths were checked after cleanup; zero remain:

- `.tmp/p6b-repair-generation`
- `.tmp/p6b-repair-packaged-benchmark.json`
- `.tmp/p6b-runtime`
- `.tmp/p6b-held-runtime`
- `.tmp/p6b-locker.ps1`
- `.tmp/ps-compile`
- `.tmp/go-build`
- `.tmp/go-cache`
- `.tmp/p6b-temp`
- `anvien/.tmp/go-build`
- `internal/analyze/testdata/p6b-tsstdlib-runtime/.anvien`
- `internal/analyze/testdata/p6b-tsstdlib-runtime-no-lib/.anvien`
- `internal/analyze/testdata/p6b-tsstdlib-runtime-default/.anvien`
- `anvien-launcher/server-wrapper/anvien-server-wrapper.exe`
- `anvien-launcher/src/anvien-launcher.exe`
- `.tmp/p6b-graph-evidence`
- `.tmp/p6b-analyze-benchmark.json`
- `.tmp/p6b-fixture-benchmark.json`
- `.tmp/p6b_catalog_carriage_builder.exe`
- `.tmp/p6b_restart_manager_probe.ps1`

Authoritative root `.anvien`, repo-local Ladybug cache, accepted fixtures/assets, external HEAD commit, protected reports, and user work were preserved.

## Exact Candidate Inventory

Tracked P6-B modifications (`8`):

- `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`
- `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
- `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
- `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
- `internal/analyze/analyze.go`
- `internal/resolution/indexes.go`
- `internal/resolution/resolve.go`
- `internal/resolution/types.go`

New P6-B source/test/fixture assets (`24`):

- `anvien-web/scripts/generate-tsstdlib-catalog.mjs`
- `internal/analyze/p6b_tsstdlib_test.go`
- `internal/analyze/testdata/p6b-tsstdlib-runtime-default/main.ts`
- `internal/analyze/testdata/p6b-tsstdlib-runtime-no-lib/main.ts`
- `internal/analyze/testdata/p6b-tsstdlib-runtime-no-lib/tsconfig.json`
- `internal/analyze/testdata/p6b-tsstdlib-runtime/main.ts`
- `internal/analyze/testdata/p6b-tsstdlib-runtime/tsconfig.json`
- `internal/resolution/p6b_tsstdlib_test.go`
- `internal/tsstdlib/catalog.go`
- `internal/tsstdlib/catalog.v1.json`
- `internal/tsstdlib/catalog_test.go`
- `internal/tsstdlib/profile.go`
- `internal/tsstdlib/testdata/catalog-failures/input-manifest.json`
- `internal/tsstdlib/testdata/catalog-failures/logical-hash.json`
- `internal/tsstdlib/testdata/catalog-failures/schema.json`
- `internal/tsstdlib/testdata/catalog-failures/trailing-artifact-integrity.json`
- `internal/tsstdlib/testdata/catalog-failures/version.json`
- `internal/tsstdlib/testdata/profiles/es2022/tsconfig.json`
- `internal/tsstdlib/testdata/profiles/es5-promise/tsconfig.json`
- `internal/tsstdlib/testdata/profiles/es5/tsconfig.json`
- `internal/tsstdlib/testdata/profiles/invalid-lib/tsconfig.json`
- `internal/tsstdlib/testdata/profiles/jsonc/tsconfig.json`
- `internal/tsstdlib/testdata/profiles/no-lib/tsconfig.json`
- `internal/tsstdlib/testdata/profiles/topology/tsconfig.json`

This report is the only new coder report. Candidate after this file is exactly `33` paths: `8` tracked modifications + `24` new source/test/fixture assets + `1` report.

Immutable history outside the candidate:

- `reports/coder/rp_coder_260822_055455_by_gpt-5_p6b_typescript_stdlib_authority.md`
- `reports/coder/rp_coder_260822_082304_by_gpt-5_p6b_typescript_stdlib_authority_resubmission.md`
- `reports/Supervisor/rp_supervisor_260822_062917_by_gpt-5_p6b_typescript_stdlib_authority.md`
- `reports/Supervisor/rp_supervisor_260822_090533_by_gpt-5_p6b_typescript_stdlib_authority.md`

Exactly `23` `reports/Investigation/rp_main_*` handoffs were observed pathname-only. The concurrent newest path is `reports/Investigation/rp_main_260822_094916_orchestration_rotation_handoff.md`. None was read, modified, removed, staged, or committed.

Final pre-report boundary: HEAD/base correct, index empty, `git diff --check` PASS, `59` status paths = `32` candidate-before-report + `23` protected handoffs + `4` immutable history reports. Expected post-report status is `60` paths with the exact `33`-path candidate.

## Remaining Blockers And Handoff

No residual same-scope blocker is known locally. The following are disclosed boundaries, not hidden PASS claims:

- Independent Supervisor acceptance remains pending.
- Broad `go build ./...` remains blocked by intentional non-buildable fixtures; the production build boundary passes.
- Broad tests retain the two pre-existing C#/Dart parity failures.
- Official lifecycle script remains prohibited because it performs install/package scripts.
- Target validation and all later P6 slices remain locked.

Terminal coder status: READY_FOR_INDEPENDENT_SUPERVISOR.
