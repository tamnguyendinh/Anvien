# Coder Resubmission: P6-B TypeScript Authority Proof-Matrix Residual Repair

Status: **READY_FOR_INDEPENDENT_SUPERVISOR**

This is a coder handoff, not acceptance. P6-B remains unchecked, unstaged, uncommitted, and blocked from P6-C1/C2/C3/D until an independent Supervisor PASS and Main action.

## Metadata

- Prepared at: `2026-08-22T11:55:44.2354518+07:00`.
- Workspace: only `E:\Anvien`.
- Branch: `master`.
- Authoritative HEAD: `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`.
- Sole parent and P6-B implementation base: `b98131e44932a7bcac17b487ecb2914535927d01`.
- External HEAD-only path remains preserved and outside P6-B.
- Third immutable Supervisor authority: `reports/Supervisor/rp_supervisor_260822_104600_by_gpt-5_p6b_typescript_stdlib_authority.md`, verdict `REJECT`, `18,998` bytes / `170` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `C41ACF28BC65020CC75925238B02655D56F27363E952449B9CB3C9CE29D2A422` / createdAt `2026-08-22T10:48:08.0156563+07:00`.
- Pre-report Git boundary: exactly `63` paths, index empty, `git diff --check` exit `0`.
- Pre-report candidate: exactly `33` paths. This report is the only new report in this turn and makes the candidate exactly `34` paths.
- No stage, commit, push, target access, alternate worktree, network, install, package script, or internal subagent occurred.

## Scope And Verdict Boundary

The sole repair is the same-invariant contradiction identified by the third Supervisor report:

```text
CatalogProofReady
+ non-empty AuthorityHash/CatalogHash/CatalogArtifactHash
+ capability_unavailable
+ catalog_missing or a catalog-rejection reason
```

The previous validator accepted that payload because the ready branch checked only hashes and the outer unavailable branch checked only a non-empty reason plus empty resolved symbol/declarations. This simultaneously claimed validated ready identity and rejected catalog authority.

Preserved clearances were not reworked: exact typed loader reason, missing/rejected absence semantics, six catalog-failure classes through authority/resolver/built `analyze.Run`, `14,587/14,587` canonical semantic IDs, receiver terminality, the exact ten-vector matrix, valid-catalog carriage/dedupe/counter equality, offline TypeScript `5.9.3` provenance, profile/config fail-closed behavior, and package/runtime isolation.

No CLI projection, `ExternalSymbol`, persistence, graph-health, shared final DTO, project/package lookup, synthetic `IMPORTS`, P6-C/P6-D owner, or target repository was opened for edit.

## Invariant Family Map

- Family: lossless TypeScript standard-library authority proof validation at the per-source-site finalizer boundary.
- Authority / SSOT: accepted P6-A decision committed at `b98131e44932a7bcac17b487ecb2914535927d01`; exact third Supervisor report above; current `tsstdlib.Authority` lookup/profile contract.
- Primary path: `LookupResult` -> `recordTypeScriptLookup` -> `TypeScriptAuthorityResult` -> `finalizeTypeScriptAuthorityResults` -> `validateTypeScriptAuthorityResult` -> `validateTypeScriptCatalogProof`.
- Sibling surfaces checked: ready resolved, ready profile-excluded, ready meaning-mismatch, ready noLib/config-unavailable, missing catalog, rejected schema/version/logical-hash/input-manifest catalog, site/counter equality, replay, and conflict rejection.
- Forbidden fallback: no generic non-empty reason acceptance for ready proof; no fabricated authority/logical hashes for invalid catalog; no conversion to a repository gap; no later-slice materialization or projection.
- Observable closure: every accepted proof-state/status/reason/hash row validates, every cross-state contradiction fails, and generated valid/missing/rejected records still traverse direct resolution and built `analyze.Run` without analysis error.
- Residual unverified surfaces: none inside this assigned blocker. Independent acceptance remains pending.

## Production Repair

Production code was changed before tests.

`internal/resolution/resolve.go` now enforces this explicit accepted matrix:

| Proof state | Accepted status | Exact reason | Hash contract |
|-------------|-----------------|--------------|---------------|
| `ready` | `resolved` | empty | authority, logical catalog, and artifact hashes present |
| `ready` | `profile_excluded` | `profile_excludes_declaration` | all three ready hashes present |
| `ready` | `meaning_mismatch` | `meaning_mismatch` | all three ready hashes present |
| `ready` | `capability_unavailable` | exactly `disabled_by_no_lib`, `config_invalid`, `config_topology_unsupported`, or `config_unreadable` | all three ready hashes present |
| `missing` | `capability_unavailable` | `catalog_missing` | authority, logical catalog, and artifact hashes all absent |
| `rejected` | `capability_unavailable` | exactly schema, version, logical-hash, or input-manifest rejection | authority/logical hashes absent; attempted-artifact hash present |

Ready proof rejects `not_found`, empty/unknown unavailable reasons, every catalog-failure reason, and status-specific reason swaps. The outer status validator also keeps profile-excluded and meaning-mismatch reasons explicit while preserving current resolved/unavailable output-shape semantics.

`internal/resolution/p6b_tsstdlib_test.go` adds the direct table-driven accepted/forbidden matrix only after production behavior was correct. No production/test owner outside the existing P6-B resolution family changed in this turn.

## Fresh Anvien Preflight And Blast Radius

- Packaged help and important analyze/file-detail/impact/detect/doctor/clean flag help were read before use.
- Initial fresh analyze: `2,026` scanned / `752` parsed code / `0` failed / `116,147` nodes / `161,274` relationships; indexed/current commit equal `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`; stale=false.
- `internal/resolution/resolve.go`: HIGH file risk; `183` symbols / `110` inbound / `281` outbound / `81` local / `366` unresolved / `23` linked flows / `36` linked tests.
- `resolve.go` upstream file impact with tests: CRITICAL `132` impacted symbols / `44` affected files / `105` direct / `1` flow.
- `validateTypeScriptAuthorityResult`: CRITICAL `8` impacted / `1` direct / `2` modules / `23` processes.
- `validateTypeScriptCatalogProof`: CRITICAL `5` impacted / `1` direct / `1` module / `11` processes.
- Post-production/pre-test fresh analyze: `2,026/752/0`, `116,168` nodes / `161,295` relationships.
- `internal/resolution/p6b_tsstdlib_test.go`: LOW file risk; `111` symbols / `0` inbound / `96` outbound / `33` local / `0` unresolved. File impact was MEDIUM `9` impacted / `1` file.
- Each of the four Child 06 ledgers was LOW with one inbound relationship and zero impacted symbols before planner refresh.

HIGH/CRITICAL values are fully disclosed blast-radius warnings. They were not treated as prohibitions or reduced for presentation.

## Exact Matrix And Affected Regression

Direct verbose command:

```text
go test ./internal/resolution -run '^TestP6BTypeScriptAuthorityValidationProofStatusReasonMatrix$' -count=1 -v
```

Result: PASS, `25` individually named subrows:

- `7` accepted ready rows: resolved, profile-excluded, meaning-mismatch, no-lib, config-invalid, config-topology, config-unreadable.
- `6` rejected invalid-ready rows: resolved/profile-exclusion reason swap, profile/meaning swaps, unavailable empty reason, unavailable unknown reason, and not-found.
- `6` accepted missing/rejected rows: empty-missing, schema, version, input-manifest, logical-hash, trailing-artifact-integrity.
- `6` rejected ready-plus-catalog-failure rows with otherwise complete ready hashes, one for each of those exact class identities.

Additional patch-invalidated evidence:

- `go test ./internal/resolution -run '^TestP6B' -count=1` -> PASS (`0.593s`).
- `go test ./internal/tsstdlib ./internal/resolution ./internal/analyze -count=1` -> PASS (`6.333s / 1.534s / 4.547s`).
- Verbose six-class authority gate -> PASS (`1.055s`).
- Verbose six-class direct-resolver gate -> PASS (`0.785s`).
- Verbose six-class built-`analyze.Run` gate -> PASS (`1.318s`).

The six generated failure classes continue to preserve exact reason/proof absence, unique site records, counter equality, deterministic replay, zero generic repository gaps, and conflicting-payload rejection. The broader full regression was not mechanically rerun because this patch changes only the final authority validator and its direct test; the affected authority/resolution/analyze packages are the real invalidated boundary. Historical C#/Dart provider baselines are not relabeled.

## Build, Package, And Runtime Evidence

Rule 13 preflight:

- `doctor locks --repo E:\Anvien --json`: analyze lock free/nonexistent.
- The exact destination process query became empty after the read-only doctor process exited.
- Go build cache/temp and all temporary outputs were repo-local under `E:\Anvien\.tmp`.

Build results:

- Broad `go build ./...`: exit `1` only on the existing intentional fixture imports (`models`, `animal`), mixed `animal/main` fixture packages, and standalone C fixture source. It is not labeled PASS; no P6-B compile error appeared.
- `go build ./cmd/... ./internal/...`: PASS.
- `anvien-launcher/src` module build: PASS to repo-local temporary output.
- `anvien-launcher/server-wrapper` module build: PASS to repo-local temporary output.

Destination lifecycle:

- The first real destination build compiled source but failed replacement when the old destination became held during the build.
- A fresh exact process check was empty. Both source and backup endpoints were resolved and verified inside `E:\Anvien`; the old destination was moved to lane-local cleanup storage so the real destination could be rebuilt without substituting a temp artifact.
- Hidden Go lifecycle helper then PASS and wrote `E:\Anvien\anvien\bin\anvien.exe`.
- Restart Manager later identified exact moved-backup holders PIDs `4172`, `16908`, and `18224`, all `anvien.exe mcp` console processes. Only these PIDs were terminated; all were verified gone and Restart Manager returned holder count `0` before deletion. No broad kill occurred.
- `package ensure-runtime`: PASS.
- Packaged `--help`: PASS.
- Packaged real fixture analyze: PASS, `2` scanned / `1` parsed / `0` failed, graph `8` nodes / `6` relationships.
- Fixture index was removed after evidence collection.

No official lifecycle script ran because it performs prohibited install/package-script work. Mandatory process metadata was observed read-only; no file read/write target outside `E:\Anvien` occurred.

## Artifact And Benchmark Evidence

- Final destination binary: `73,613,824` bytes.
- SHA-256: `87E8B2696C3851F58BD389B8E0C2A6EFF0D87E41A3A1399D59A6781361CB9BF6`.
- Build metadata: Go `1.26.3`, `ladybugdb`, trimpath, VCS revision `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`, modified=true.
- Delta from accepted P5-D `71,509,504`: `+2,104,320` bytes (`+2.9427%`).
- Delta from prior catalog-failure residual binary `73,610,752`: `+3,072` bytes (`+0.0042%`).

Generator/catalog/profile/lookup owners did not change. The accepted catalog `2,003,050` bytes, logical/artifact hashes, `100`-file input manifest, `14,587/14,587` semantic identity inventory, two-run byte equality, and three-run load/global/member lookup measurements therefore remain valid and were not rerun. This turn remeasured the changed packaged artifact and graph inventory only; no threshold was invented.

## Final Graph And Detect Evidence

After source, tests, all four ledger updates, this report path, and the concurrent protected handoff were present, the final packaged self-analyze PASS was:

- `2,028` scanned;
- `752` parsed code;
- `0` failed;
- `116,193` nodes;
- `161,365` relationships;
- `19,426` dependency edges.

Required explicit-path command:

```text
anvien detect-changes --repo E:\Anvien --scope all --json
```

Result: exit `0`, risk CRITICAL:

- affected: `35` symbols / `8` files;
- changed: `328` symbols / `8` files;
- affected layers: `backend=31`, `mixed=4`;
- affected functional areas: `analyzer=8`, `mixed=9`, `resolution=18`;
- changed layers: `backend=309`, `docs=19`;
- changed functional areas: `analyzer=28`, `documentation=19`, `resolution=281`.

Resolution-gap delta is `198` entities / `201` occurrences:

- actionability: `analyzer_gap=137`, `non_actionable=61`;
- classification: `builtin=32`, `in_repo_unresolved=137`, `standard_library=29`;
- families: `access=106`, `call=80`, `type-reference=12`;
- functional areas: `analyzer=14`, `resolution=184`.

Semantic app-layer and functional-area fields are complete for `116,193/116,193` nodes. Current Resolution Health total is `0`. Counts and the CRITICAL warning are preserved without normalization.

## Cleanup And Protected Boundary

Turn-created cleanup is `9/9` absent:

1. `.tmp/p6b-proof-matrix` cache/launcher/backup tree;
2. `.tmp/p6b-rm-holder.ps1`;
3. `.tmp/p6b-rm-compile`;
4. `internal/analyze/testdata/p6b-tsstdlib-runtime/.anvien`;
5. `internal/analyze/testdata/p6b-tsstdlib-runtime-default/.anvien`;
6. `internal/analyze/testdata/p6b-tsstdlib-runtime-no-lib/.anvien`;
7. `anvien-launcher/src/src.exe`;
8. `anvien-launcher/src/anvien-launcher.exe`;
9. `anvien-launcher/server-wrapper/server-wrapper.exe`.

Root `.anvien` remains intentionally present for the mandatory final graph. The existing repo-local Ladybug native cache was used, not recreated as evidence or deleted as unrelated state.

Exactly `25` protected Main handoffs were observed by pathname only and preserved after the concurrent `11:50` rotation added `reports/Investigation/rp_main_260822_115036_orchestration_rotation_handoff.md`; none was read, edited, removed, staged, or committed. All three Supervisor reports and all three prior coder reports are immutable. The external HEAD commit/path and all user work remain preserved.

## Exact Candidate Inventory After This Report

Eight tracked modifications:

1. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`
2. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
3. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
4. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
5. `internal/analyze/analyze.go`
6. `internal/resolution/indexes.go`
7. `internal/resolution/resolve.go`
8. `internal/resolution/types.go`

Twenty-four new P6-B source/test/fixture assets:

9. `anvien-web/scripts/generate-tsstdlib-catalog.mjs`
10. `internal/analyze/p6b_tsstdlib_test.go`
11. `internal/analyze/testdata/p6b-tsstdlib-runtime-default/main.ts`
12. `internal/analyze/testdata/p6b-tsstdlib-runtime-no-lib/main.ts`
13. `internal/analyze/testdata/p6b-tsstdlib-runtime-no-lib/tsconfig.json`
14. `internal/analyze/testdata/p6b-tsstdlib-runtime/main.ts`
15. `internal/analyze/testdata/p6b-tsstdlib-runtime/tsconfig.json`
16. `internal/resolution/p6b_tsstdlib_test.go`
17. `internal/tsstdlib/catalog.go`
18. `internal/tsstdlib/catalog.v1.json`
19. `internal/tsstdlib/catalog_test.go`
20. `internal/tsstdlib/profile.go`
21. `internal/tsstdlib/testdata/catalog-failures/input-manifest.json`
22. `internal/tsstdlib/testdata/catalog-failures/logical-hash.json`
23. `internal/tsstdlib/testdata/catalog-failures/schema.json`
24. `internal/tsstdlib/testdata/catalog-failures/trailing-artifact-integrity.json`
25. `internal/tsstdlib/testdata/catalog-failures/version.json`
26. `internal/tsstdlib/testdata/profiles/es2022/tsconfig.json`
27. `internal/tsstdlib/testdata/profiles/es5-promise/tsconfig.json`
28. `internal/tsstdlib/testdata/profiles/es5/tsconfig.json`
29. `internal/tsstdlib/testdata/profiles/invalid-lib/tsconfig.json`
30. `internal/tsstdlib/testdata/profiles/jsonc/tsconfig.json`
31. `internal/tsstdlib/testdata/profiles/no-lib/tsconfig.json`
32. `internal/tsstdlib/testdata/profiles/topology/tsconfig.json`

Two coder reports in the active candidate:

33. `reports/coder/rp_coder_260822_101757_by_gpt-5_p6b_typescript_stdlib_authority_catalog_failure_resubmission.md` — immutable prior candidate report.
34. `reports/coder/rp_coder_260822_115544_by_gpt-5_p6b_typescript_stdlib_authority_proof_matrix_resubmission.md` — this report.

The two older coder reports, all three Supervisor reports, and all `25` protected Main handoffs are outside the `34`-path candidate. Total status after this report is `64` paths. Index remains empty; no stage/commit/push is authorized.

## Exact Owned Diff In This Repair Turn

- `internal/resolution/resolve.go`: exact accepted ready/missing/rejected proof-state/status/reason/hash matrix and explicit profile/mismatch output validation.
- `internal/resolution/p6b_tsstdlib_test.go`: direct `25`-subrow positive/negative matrix and complete-hash ready rejection rows.
- Four Child 06 ledgers: third REJECT truth, R10 refresh, current evidence/measurements, P6-B still unchecked, and later slices locked.
- This single immutable coder report.

No other candidate file was edited in this repair turn.

## Remaining Blocker And Handoff

Implementation blocker inside the assigned residual: **none observed after current source/build/test/runtime/graph/detect/cleanup evidence**.

Acceptance blocker: a new independent Supervisor verdict is required. Main must not stage, commit, open P6-C1/C2/C3/D, or access the target from this coder report alone.

Final coder disposition: **READY_FOR_INDEPENDENT_SUPERVISOR**.
