# Supervisor Report: P6-B TypeScript Standard-Library Authority Reject-Only Re-review

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260822_090533_by_gpt-5_p6b_typescript_stdlib_authority.md`
- Review time: `2026-08-22 09:05:33 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`
- Scope reviewed: only Child 06 P6-B “TypeScript standard-library authority,” limited to the four invariants rejected by the immutable prior review and the same P6-B invariant family. P6-C1/P6-C2/P6-C3/P6-D were not opened or reviewed.
- Claim reviewed: the reject-only resubmission closes all four prior blockers, leaves no required same-scope follow-up, and is ready for independent acceptance before Main stages or commits P6-B.
- Current anchor: branch `master`, HEAD `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`, parent and P6-B implementation base `b98131e44932a7bcac17b487ecb2914535927d01`, predecessor `ec765debff335540c77d409ebb2c9f45e4a0a77d`.
- External HEAD commit: the only `5bfdfb3...` path is `internal/aicontext/skills/orchestration/SKILL.md`; it was preserved and excluded from P6-B.
- Candidate before this report: exactly 8 tracked modifications + 19 new P6-B source/test/fixture assets + immutable coder resubmission report = 28 paths. The old coder report and prior Supervisor report are immutable history outside the candidate.
- Protected boundary: exactly 22 `reports/Investigation/rp_main_*` handoffs at final pre-report status. Only their pathnames were observed. None was read, edited, removed, staged, or committed. Handoff #22, `reports/Investigation/rp_main_260822_085540_orchestration_rotation_handoff.md`, arrived after the fresh graph snapshot and is authorized external state, not candidate drift; the graph was not restarted.
- Git boundary before this report: 52 status paths = 28 candidate + 2 immutable history + 22 protected handoffs; index empty; `git diff --check` exit 0.
- Immutable resubmission seal independently reproduced: 20,479 bytes / 257 LF / 0 CR / strict UTF-8 / no BOM / SHA-256 `2483B48C74B0621C59CE57FED8EE0D010507A050ED6726485EEDFB1D5EDE11E1` / createdAt and lastWrite `2026-08-22T08:25:26.8212759+07:00`.
- Report identity protocol: this report is written once and externally sealed after creation. Its final bytes, LF/CR counts, UTF-8/BOM state, SHA-256, and filesystem createdAt are supplied in the final handoff; no self-hash is embedded.

## Executive Summary

The resubmission fully repairs three prior blockers and the normal valid-catalog path of the first blocker. Canonical semantic identities are component-complete and globally unique; local/repository/import receiver claims are terminal before call/access external fallback; and the durable oracle is an exact named ten-row reproduction of the P6-A compiler matrix. For a valid embedded catalog, handled lookup results are retained losslessly on `resolution.Result` and built `analyze.Result`, duplicate identical facts canonicalize to one record/counter, and conflicting payloads fail closed.

One same-invariant P6-B blocker remains. When embedded catalog validation fails, the authority preserves the exact validation reason only in its internal unavailable profile, then public lookup collapses it to `catalog_missing` and returns empty authority/catalog/config provenance. The new P6-B site finalizer rejects that record as incomplete, so the required reasoned per-source-site capability outcome becomes an analyze error. P6-A explicitly requires missing/schema/version/manifest/hash failures to make the authority unavailable and yield a reasoned capability outcome. This is not CLI, graph materialization, persistence, or P6-C3 final-outcome work; it is a lossless P6-B lookup-result path within prior blocker one.

## Blocking Finding

### [CRITICAL] Catalog-validation failures lose their exact reason and cannot survive P6-B site finalization

Files:

- `internal/tsstdlib/catalog.go:138-142`
- `internal/tsstdlib/catalog.go:318-345`
- `internal/resolution/resolve.go:875-911`
- `internal/resolution/resolve.go:913-965`
- `internal/resolution/resolve.go:967-1005`

Authority:

- Immutable P6-A section 6, line 102 requires a missing asset, unsupported schema, version mismatch, input-manifest mismatch, or catalog-hash mismatch to make the authority unavailable and yield a reasoned capability outcome.
- P6-A section 8, line 145 includes distinct minimum reason codes for `catalog_missing`, `catalog_schema_unsupported`, `catalog_version_mismatch`, and `catalog_hash_mismatch` and requires affected-site authority/catalog/config proof.
- The prior Supervisor blocker one requires an observable site record for every handled capability-unavailable class, not aggregate-only disappearance or an analyze error.

Source evidence:

1. `NewAuthority` receives the typed reason from `loadDefaultCatalog` and stores it in `unavailableProfile(reason, "", "")` at `catalog.go:138-142`.
2. `Authority.availability` checks `authority.catalog == nil` first and always returns `ReasonCatalogMissing` at `catalog.go:318-324`; it does not return the stored profile reason. Schema/version/hash/manifest distinctions are therefore lost at lookup.
3. `baseResult` at `catalog.go:328-345` can populate authority/catalog/artifact hashes only when `authority.catalog != nil`. The catalog-error profile also has an empty config hash, so a handled catalog-unavailable lookup has empty `AuthorityHash`, `CatalogHash`, `CatalogArtifactHash`, and `ConfigHash`.
4. `recordTypeScriptLookup` at `resolve.go:875-911` treats that lookup as handled and appends it. `finalizeTypeScriptAuthorityResults` then calls `validateTypeScriptAuthorityResult`; the validator at `resolve.go:967-983` requires all four fields to be non-empty for every status. It returns an error before an immutable site result can leave `ResolveBoundInto`/`analyze.Run`.

Independent runtime probe:

- A repo-local `.tmp` debug probe used a Go build overlay to replace only the embedded catalog input with `{}`; candidate files were not modified. `NewAuthority.Profile()` reported `unavailable/catalog_schema_unsupported`, proving typed loader reason retention. `LookupGlobal("Promise", type)` reported `capability_unavailable/catalog_missing`, with empty authority hash, logical catalog hash, artifact hash, and config hash. The exact probe directory and binary were removed immediately; final `Test-Path` was false.
- This directly confirms both reason collapse and the payload that the source-site validator rejects. The probe did not access a target, network, package script, alternate checkout, or protected report.

Why this blocks acceptance:

- The valid-catalog noLib/config-unavailable tests cannot stand in for catalog validation failure. A catalog integrity failure is an expressly accepted P6-B capability path, and it currently produces neither the exact reason nor an observable site record.
- Fatal fail-closed validation is necessary but not sufficient here: P6-A chose a reasoned capability outcome, and the prior rejection required that handled lookup truth survive at the P6-B boundary.
- Deferring the correction to P6-C2/C3 would be illegal. Those slices cannot consume a result that P6-B discards through an error, and they do not own catalog-reason fidelity.

Fix direction:

- Preserve the exact typed loader reason through `Authority.availability` instead of collapsing all nil-catalog states to `catalog_missing`.
- Define explicit, fail-closed provenance/absence semantics for catalog-unavailable site records so missing or rejected catalog facts can pass P6-B validation without inventing a ready catalog hash. Retain all evidence that is knowable (authority kind, expected TypeScript/runtime contract, attempted artifact/provenance where applicable, profile/config/inventory proof) and make unavailable fields explicit rather than silently substituting a different reason.
- Ensure `ResolveBoundInto` and built `analyze.Run` return one immutable capability-unavailable record per unique affected site, with counter equality and no generic repository gap, for each catalog failure class. Do not add CLI projection, `ExternalSymbol`, persistence, graph-health, or a P6-C3 final DTO in this repair.

Required re-review evidence:

1. Table-driven authority tests for empty/missing, schema, version, input-manifest, logical-hash, and trailing/artifact-integrity catalog failures, asserting the exact accepted reason rather than a collapsed reason.
2. Direct resolver and built `analyze.Run` tests proving each catalog failure yields observable per-site records, no analysis error, no generic repository gap, exact unique-record/counter equality, identical replay canonicalization, and conflict rejection.
3. Source proof of the catalog-unavailable provenance/absence contract and validation rules; no fake ready-catalog identity or target-name branch.
4. Proportionate post-fix production build, packaged runtime, focused/affected regression, artifact measurement, cleanup, fresh Anvien graph evidence, and explicit-path detect evidence. Preserve the known broad-command limitations rather than relabeling them.

## Disposition Of The Four Prior Blockers

1. Lossless site-level result: closed for valid-catalog resolved, noLib/config capability-unavailable, profile-excluded, and meaning-mismatch paths. `TypeScriptAuthorityResult` retains source site, request, status/reason, target/owner, declarations, and provenance; `finalizeTypeScriptAuthorityResults` clones/sorts, deduplicates only `reflect.DeepEqual` payloads, rejects conflicting payloads, and derives counters from unique records. Still blocked only for catalog-validation capability outcomes as described above.
2. Canonical semantic identity: closed. Generator and loader use a two-phase logical catalog identity; lane IDs commit to authority kind, exact TypeScript version, logical catalog hash, canonical declaration library/range, semantic owner path, name, and meaning. Loader recomputes component IDs and rejects duplicates/component drift.
3. Receiver terminality: closed. `resolveCall` and `resolveAccess` establish `externalBlocked` from local/repository/import receiver claims before catalog member lookup; `lookupTypeScriptMember` repeats the repository-claim guard. Genuine unclaimed external receivers remain eligible. The matrix covers call and access forms for local value, repository-owned colliding receiver, explicit-import failure, and genuine external receiver without target-name branches.
4. Exact compiler matrix: closed. `TestP6BExactCompilerVectorMatrix` has an enforced denominator of ten and one named row corresponding exactly to P6-A lines 65-74, including all four ES2015 values, eight noLib global types, Promise type, and both Math members where required.

## Source-Level Clearance

- `anvien-web/scripts/generate-tsstdlib-catalog.mjs`: clear. Offline/developer-only TypeScript `5.9.3` compiler API, exact lock integrity, sorted official inputs, deterministic logical hash before IDs, component-complete lane identities, and no target-name behavior. Node/`node_modules` use is confined to this generator.
- `internal/tsstdlib/catalog.v1.json`: clear. Independently measured facts and all input hashes match; IDs are complete and unique.
- `internal/tsstdlib/catalog.go`: clear on embed, schema/provenance/trailing-data/logical-hash/input/lane/inheritance validation, identity derivation, duplicate/component-drift rejection, lookup meaning separation, and active-profile filtering. Blocked only on the catalog-error reason/carriage path identified above.
- `internal/tsstdlib/profile.go`: clear. TypeScript-only zero/one root JSONC topology, `target`/`lib`/`noLib`, explicit library closure, unreadable/invalid/unsupported fail-closed behavior, and deterministic profile/config hashes remain within P6-A.
- `internal/analyze/analyze.go`: clear for TypeScript-only authority construction from the complete scanned inventory and lossless transfer of `resolutionResult.TypeScriptAuthorityResults` onto `analyze.Result`. The existing CLI adapter at `internal/cli/command.go:251-262` remains summary-only, but CLI projection is not a P6-B owner in the accepted P6-A owner map. Exact records are observable at built `analyze.Run`; no CLI/C2/C3 expansion is required by this review. Catalog-error resolution still bubbles the blocking finalizer error.
- `internal/resolution/types.go` and `indexes.go`: clear. The P6-B-only site result and authority pointer remain separate from repository ScopeIR and P6-C2 graph materialization/P6-C3 final cross-stage outcomes.
- `internal/resolution/resolve.go`: clear on repository/global/type/import precedence, receiver terminality, site identity, lossless record assembly, identical-payload canonicalization, conflicting-payload failure, and counter equality for valid-catalog paths. Blocked only because the validator assumes ready catalog/config hashes for catalog-unavailable status.
- Test/fixture owners: clear on the exact ten-vector matrix, canonical ID inventory, valid-catalog site carriage, dedupe/conflict, receiver call/access matrix, built noLib/profile/mismatch carriage, and fixture cleanup. Missing evidence is limited to end-to-end catalog-validation capability carriage.
- Package surface: clear. `anvien-web/package.json`, `anvien-web/package-lock.json`, and `anvien/package.json` have zero diff. Production Go source contains no runtime Node/`tsc`, command execution, network/install/package-script fallback, package declaration scan, target path, or target-name allowlist.
- Four living ledgers: time-bounded catalog, graph, build, benchmark, regression, cleanup, and lock statements are internally consistent. Their “all four blockers closed / residual none” claim is not accepted because it omits the catalog-validation subpath of blocker one. P6-B correctly remains unchecked/uncommitted and P6-C1/C2/C3/D remain locked.

## Current Anvien Evidence

- Repo-local packaged help was read before use. Exactly one fresh explicit-path `anvien analyze E:\Anvien --force --json` ran in this review before graph queries. It exited 0 with `2,015` scanned / `752` parsed code / `0` failed / `115,911` nodes / `160,947` relationships at `E:\Anvien\.anvien\graph.json` on indexed/current commit `5bfdfb3...`. This matches Main's later-report snapshot; it is intentionally distinct from the coder's pre-report `2,013 / 115,877 / 160,913` snapshot.
- Handoff #22 appeared after the graph analyzedAt timestamp. Owner authority excludes it from candidate drift and forbids reading it, so the graph was not restarted.

Current compact file-detail summaries (all parsed, high risk, stale=false, changedSinceAnalyze=false):

| Owner | Symbols | In / Out / Local | Unresolved | Flows / Tests |
| --- | ---: | ---: | ---: | ---: |
| `internal/analyze/analyze.go` | 333 | 91 / 317 / 174 | 522 | 20 / 8 |
| `internal/resolution/types.go` | 92 | 75 / 24 / 88 | 35 | 1 / 20 |
| `internal/resolution/indexes.go` | 344 | 323 / 108 / 388 | 453 | 19 / 32 |
| `internal/resolution/resolve.go` | 180 | 106 / 274 / 79 | 351 | 23 / 35 |
| `internal/tsstdlib/catalog.go` | 328 | 137 / 2 / 406 | 234 | 1 / 3 |
| `internal/tsstdlib/profile.go` | 91 | 8 / 28 / 10 | 112 | 1 / 2 |
| generator | 130 | 0 / 0 / 45 | 312 | 0 / 0 |

Current upstream symbol impacts with tests preserve their severity and meaning:

- `resolveCall`, `resolveAccess`, and `resolveTypeAnnotation`: each CRITICAL, 30 impacted symbols / 12 affected files / 6 modules / 32 processes.
- `recordTypeScriptLookup`: CRITICAL `8 / 4 / 2 / 35`.
- `finalizeTypeScriptAuthorityResults`: CRITICAL `31 / 13 / 6 / 32`.
- `loadCatalog`: CRITICAL `16 / 4 / 3 / 12`.
- `selectProfile`: CRITICAL `37 / 10 / 7 / 21`.
- Pure `catalogSemanticID`: LOW `4 / 2 / 1 / 0`.
- Every queried impact reported semantic app-layer and functional-area status complete and current health ResolutionGap total 0. CRITICAL/HIGH are blast-radius warnings, not edit bans.

Fresh explicit-path `detect-changes --repo E:\Anvien --scope all --json` exited 0 with CRITICAL risk:

- affected `35` symbols / `8` files; changed `288` symbols / `8` files;
- affected layers `backend=31`, `mixed=4`; areas `analyzer=8`, `mixed=9`, `resolution=18`;
- ResolutionGap delta `162` entities / `165` occurrences; actionability `analyzer_gap=115`, `non_actionable=47`; classifications `builtin=29`, `in_repo_unresolved=115`, `standard_library=18`; families `access=84`, `call=67`, `type-reference=11`;
- semantic app-layer and functional-area fields complete; current health total 0.

## Artifact, Build, Runtime, Regression, Benchmark, And Cleanup Disposition

Independent artifact checks:

- Catalog: 2,003,050 bytes / 1 LF / 0 CR / no BOM / SHA-256 `F188D15A5D91925DF3E724CBAB97964813E3F6DFD9DF7408FDC7B92EA4CEA487`; stored and independently recomputed logical hash `dca7af22ff26510cf9075fcb587ab76649bf169edcb0001e50c234d89c9dbb0b`.
- Catalog inputs: 100 files / 3,141,835 bytes / 0 missing / 0 manifest hash-or-size mismatches.
- Catalog inventory: 2,030 symbols / 11,802 direct member rows / 14,587 IDs / 14,587 unique / 0 duplicate / 0 format mismatch / 0 lane-without-ID / 0 ID-without-lane.
- Packaged runtime: `anvien/bin/anvien.exe`, 73,605,632 bytes / SHA-256 `5BCE56084C58510FE120A8014587EA70E49B5A8C296BC5D038976664C9CEE9AA`; Go `1.26.3`, `ladybugdb`, trimpath, VCS `b98131e...`, modified=true. Packaged help and the fresh self-analyze succeeded.
- Against accepted P5-D baseline 71,509,504 bytes, recorded P6-B delta is +2,096,128 bytes (+2.9313%). Against the rejected P6-B binary 72,961,536 bytes, delta is +644,096 bytes (+0.8828%).

Build and regression evidence retained without relabeling:

- Official lifecycle build remains prohibited and was not run because it executes installation/package scripts.
- Broad `go build ./...` remains a failure on intentional non-buildable fixtures; it is not a successful full-build gate.
- Production `go build ./cmd/... ./internal/...`, both shipping launcher modules, direct package-runtime build, runtime ensure, and packaged help were reported successful after the final repair source patch. They were not mechanically rerun after source inspection established this rejection blocker.
- Full `go test ./cmd/... ./internal/... -count=1` retains the two accepted C#/Dart parity baselines and three disclosed CLI repo-local-temp/root-`.anvien` isolation failures. P6-B packages were reported successful; no unrelated baseline was edited or promoted to success.
- The sanitized E-only PATH/native-loader probe remains not a success and is not used as no-Node proof.

Benchmark evidence retained as measurement, not invented threshold:

- Embedded load `152.147-160.025 ms/op`, `36,393,174-36,451,851 B/op`, `291,812-291,818 allocs/op`.
- Global lookup `250.3-277.2 ns/op`, `192 B/op`, `1 alloc/op`.
- Inherited member `642.8-654.3 ns/op`, `216 B/op`, `4 allocs/op`.
- The prior two clean generation outputs and equal logical/artifact hashes remain valid because generator/catalog bytes did not change during this review; the generator was not rerun mechanically.

Cleanup:

- The coder's grouped 17-path cleanup claim was covered, and all three current P6-B fixture `.anvien` directories were conservatively checked. In total 18 concrete dead/debug/build paths were tested; 0 existed.
- The Supervisor's temporary catalog probe was removed separately; its directory is absent.
- No candidate, protected handoff, immutable report, root `.anvien`, or unrelated user work was removed.

## Invariant Closure And Required Resubmission

- Closed: deterministic/provenanced catalog; exact inputs and hashes; component-complete unique identities; loader duplicate/component drift rejection; profile closures/config topology; meaning lanes/inheritance; TypeScript-only injection; no runtime Node/network/install/script path; repository/import precedence; local/repository/import receiver terminality; genuine external receiver lookup; valid-catalog site carriage/dedupe/conflict/counter equality; exact ten-row oracle; package and cleanup boundary.
- Open: exact reason and lossless per-source-site capability carriage for catalog missing/schema/version/manifest/hash validation failures.
- CLI projection: summary-only by existing adapter, but not an open P6-B invariant because the accepted P6-B owner is built `analyze.Run`, where valid-catalog records are directly observable. The rejection does not request CLI work.
- Later slices: P6-C1/C2/C3/D remain locked and unchanged. Do not use them to repair or hide this P6-B catalog-failure path.

Return only the catalog-failure carriage repair and its exact re-review evidence. Keep P6-B unchecked, unstaged, and uncommitted until a new independent decision and Main action.
