# Child 06 P6-B TypeScript Standard-Library Authority — Independent Reject-Only Re-review

## Verdict

**REJECT**

The resubmission closes the previously observed production path for exact catalog-validation failure reason/provenance and lossless site carriage, but one same-invariant fail-closed hole remains in `validateTypeScriptAuthorityResult`: a record can claim a validated `ready` catalog identity while carrying a catalog-rejection reason. That contradictory payload is accepted when its hashes are non-empty. This violates the required proof-state/status/reason/hash matrix and the candidate's own claim that every combination is enforced. No P6-C1/C2/C3/D concern is involved.

## Report Identity And Review Authority

- Report: `reports/Supervisor/rp_supervisor_260822_104600_by_gpt-5_p6b_typescript_stdlib_authority.md`
- Reviewer role: independent Supervisor for Child 06 P6-B only.
- Review date: 2026-08-22, Asia/Bangkok.
- Claim reviewed: the sole residual from the immutable 09:05:33 Supervisor report—exact catalog-validation failure reason/provenance plus lossless per-source-site capability carriage—has been repaired, while prior canonical-identity, receiver-terminality, exact ten-vector, and valid-catalog clearances remain intact.
- Acceptance standard: root `AGENTS.md`; full `.agents/skills/working-rules/SKILL.md`; full `.agents/skills/supervisor/SKILL.md`; the four current Child 06 ledgers; immutable prior Supervisor report; immutable current coder report; current source/diff/runtime/graph reality. Reports and test/build claims were treated as evidence inputs, not acceptance.
- Prior Supervisor input seal independently reproduced: `20,729` bytes / `168` LF / `0` CR / strict UTF-8 / no BOM / SHA-256 `C373D2413F1D60082904729F5A536B09D588A5D1DABE12181E454BCE2AD3209A`.
- Current coder input seal independently reproduced: `17,230` bytes / `294` LF / `0` CR / strict UTF-8 / no BOM / SHA-256 `AEEA330F1EEBC1ABA35D2FA1EA342C15F9841B858C2C58411BE4DF15154589D6`; createdAt `2026-08-22T10:19:52.2907073+07:00`.
- Exactly 23 protected Main handoffs were observed by pathname only. None was read, edited, deleted, staged, or committed.

## Exact Git And Candidate Boundary Before This Report

- Workspace: only `E:\Anvien`.
- Branch: `master`.
- HEAD: `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`.
- Sole parent and P6-B implementation base: `b98131e44932a7bcac17b487ecb2914535927d01`.
- The HEAD-only external/user-owned commit changes only the pathname `internal/aicontext/skills/orchestration/SKILL.md`; it is excluded from P6-B.
- Pre-report status: exactly `60` paths = `33` current candidate + `23` protected Main handoffs + `4` immutable-history reports.
- Index: empty.
- `git diff --check`: clean, exit `0`.
- P6-B remains unchecked at plan line 221 and independent acceptance remains unchecked at actual-status line 220. It is unstaged, uncommitted, and later slices remain locked.

The exact 33-path candidate is:

1. Eight tracked modifications:
   - the four Child 06 ledgers under `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/`;
   - `internal/analyze/analyze.go`;
   - `internal/resolution/indexes.go`;
   - `internal/resolution/resolve.go`;
   - `internal/resolution/types.go`.
2. Twenty-four new P6-B source/test/fixture assets:
   - `anvien-web/scripts/generate-tsstdlib-catalog.mjs`;
   - `internal/analyze/p6b_tsstdlib_test.go`;
   - `internal/analyze/testdata/p6b-tsstdlib-runtime-default/main.ts`;
   - `internal/analyze/testdata/p6b-tsstdlib-runtime-no-lib/main.ts`;
   - `internal/analyze/testdata/p6b-tsstdlib-runtime-no-lib/tsconfig.json`;
   - `internal/analyze/testdata/p6b-tsstdlib-runtime/main.ts`;
   - `internal/analyze/testdata/p6b-tsstdlib-runtime/tsconfig.json`;
   - `internal/resolution/p6b_tsstdlib_test.go`;
   - `internal/tsstdlib/catalog.go`;
   - `internal/tsstdlib/catalog.v1.json`;
   - `internal/tsstdlib/catalog_test.go`;
   - `internal/tsstdlib/profile.go`;
   - five catalog-failure JSON fixtures under `internal/tsstdlib/testdata/catalog-failures/`;
   - seven profile `tsconfig.json` fixtures under `internal/tsstdlib/testdata/profiles/` (`es2022`, `es5-promise`, `es5`, `invalid-lib`, `jsonc`, `no-lib`, and `topology`).
3. One new immutable coder report: `reports/coder/rp_coder_260822_101757_by_gpt-5_p6b_typescript_stdlib_authority_catalog_failure_resubmission.md`.

The two older coder reports and two older Supervisor reports are immutable history outside the candidate. This Supervisor report is review-only and outside the 33-path implementation candidate.

## Source-First Clearance By Touched Production Group

### Catalog authority and profile selection

- `internal/tsstdlib/catalog.go:63-74` defines explicit `ready`, `missing`, and `rejected` proof states with the intended absence semantics.
- `unavailableCatalogAuthority` at lines 183-198 assigns `missing` only to empty/missing input, assigns `rejected` to non-empty rejected bytes, retains only the attempted artifact SHA for rejection, and commits the exact typed reason plus inventory into profile/config proof.
- `Authority.availability` at lines 370-383 now returns the stored unavailable-profile reason before the nil-catalog fallback. The earlier reason collapse is closed.
- `baseResult` at lines 386-406 carries proof state, attempted artifact hash, profile hash, and config/inventory hash; authority/logical hashes appear only when a validated catalog exists.
- Loader checks at lines 623-658 distinguish empty, decode/trailing/input-manifest, schema, version/provenance, and logical-hash failures. The six named fixture classes map to the accepted typed reason and proof state.
- `internal/tsstdlib/profile.go` remains fail-closed for config discovery/topology and does not add runtime compiler, network, install, project/package, or target-name behavior.

This production group is cleared for the concrete six generated failure paths. Its bytes do not invalidate the earlier provenance, profile closure, or configuration-topology clearances.

### Catalog generation, artifact, and semantic identity

- The generator remains an offline generation owner; the runtime source contains only the checked-in embedded catalog path. No runtime Node/`tsc`, network, install, or package-script fallback was added.
- Current artifact measurement: `2,003,050` bytes; SHA-256 `F188D15A5D91925DF3E724CBAB97964813E3F6DFD9DF7408FDC7B92EA4CEA487`; logical hash `dca7af22ff26510cf9075fcb587ab76649bf169edcb0001e50c234d89c9dbb0`.
- Current catalog inventory: `100` inputs / `3,141,835` input bytes / `2,030` symbols / `11,802` direct member rows / `14,587` semantic IDs / `14,587` unique / `0` duplicate / `0` format mismatch.
- `catalog.go:813-918` recomputes semantic components and rejects duplicate or component-drift IDs. `catalog.go:923-963` commits authority, identity version, logical catalog hash, semantic owner path, name, meaning, and canonical declaration source ranges.

The prior two-phase canonical semantic identity clearance remains valid. The generator was not rerun because its previously accepted bytes and the catalog artifact were not invalidated by this residual repair.

### Resolution precedence, terminality, and site carriage

- `internal/resolution/indexes.go` retains repository/P5 resolution and explicit-import precedence. No project/package declaration lookup or synthetic `IMPORTS` behavior was added.
- `internal/resolution/resolve.go` retains local/repository/import receiver terminality before external member lookup for calls and accesses, with genuine unclaimed external receivers still eligible. No target-name allowlist or fixture-name branch was found.
- `internal/resolution/types.go:50-75` defines a P6-B site DTO separate from graph materialization and later cross-stage outcomes. It carries source-site identity/range, request meaning, status/reason, resolved IDs/declarations, proof state, authority/version/catalog/artifact/profile/config proof.
- `recordTypeScriptLookup` at `resolve.go:889-910` copies those fields without projection loss.
- `finalizeTypeScriptAuthorityResults` at lines 914-960 clones and canonical-sorts records, deduplicates only deeply identical payloads, rejects conflicting payloads at one stable site, derives counters from unique records, and verifies exact record/counter equality.

Valid-catalog carriage, receiver terminality, and the six concrete catalog-failure carriage paths remain cleared. The only uncleared code in this production group is the validator matrix described below.

### Built analyzer boundary

- `internal/analyze/analyze.go` carries `resolution.Result.TypeScriptAuthorityResults` through the built `analyze.Run` boundary. The exact site DTO remains observable at the P6-B owner boundary; no CLI projection was added or required for this slice.
- The six failure fixtures exercise exact site records rather than aggregate-only metrics. Existing source shows no conversion into generic repository gaps.

The built owner-boundary carriage is cleared for records actually emitted by the current authority. This does not cure acceptance of an independently contradictory record by the resolver validator.

### Tests, fixtures, and ledgers

- `internal/tsstdlib/catalog_test.go`, `internal/resolution/p6b_tsstdlib_test.go`, and `internal/analyze/p6b_tsstdlib_test.go` contain the six named rows: empty/missing, schema, version, input-manifest, logical-hash, and trailing/artifact-integrity.
- The direct resolver matrix proves three unique handled sites per class, exact counters, deterministic replay, and conflicting duplicate rejection. The built analyzer matrix proves seven named sites per class, exact counters, no analysis error, and no generic repository gap.
- Current tests inspect accepted `ready`, `missing`, and `rejected` records, but there is no direct test call that supplies `validateTypeScriptAuthorityResult` a contradictory proof-state/reason combination. Existing conflict tests mutate another payload field while keeping the valid proof-state/reason pairing.
- The four ledgers correctly keep P6-B and independent acceptance open, preserve historical limitations, and label the repair pending review. However, evidence-ledger lines 189-190 and the coder report overstate source reality by saying every state/hash/reason combination is enforced. That overstatement is the documentation manifestation of the same single blocker, not a second blocker.

### Forbidden-surface clearance

The production diff contains no CLI command/projection file, `ExternalSymbol`, graph persistence, graph-health owner, shared P6-C3 DTO, project/package authority, synthetic import edge, target access, target-name branch, runtime Node/`tsc`, network, install, or package-script path. Package manifests have zero diff. No P6-C1/C2/C3/D file was opened as a review target or changed by this review.

## Exact Remaining Blocking Invariant

`validateTypeScriptCatalogProof` at `internal/resolution/resolve.go:1005-1030` treats the three states asymmetrically:

- `missing` requires `capability_unavailable/catalog_missing` and forbids authority, logical-catalog, and artifact hashes;
- `rejected` requires `capability_unavailable`, a recognized catalog-rejection reason, absent authority/logical hashes, and a non-empty attempted artifact hash;
- `ready` at lines 1007-1010 checks only that authority, logical-catalog, and artifact hashes are non-empty. It does not reject catalog-failure reasons and does not bind status/reason to a ready-catalog outcome.

The outer validator at lines 968-1003 then accepts any unavailable status when `Reason` is merely non-empty and symbol/declaration output is empty. Therefore a complete record with:

- `CatalogProofState = ready`;
- non-empty `AuthorityHash`, `CatalogHash`, and `CatalogArtifactHash`;
- `Status = capability_unavailable`;
- `Reason = catalog_schema_unsupported` (equally, another catalog-rejection reason);
- empty resolved symbol and declaration ranges;
- all common source, authority, version, profile, and config fields valid;

passes the ready proof branch, passes the generic unavailable branch at lines 995-998, and returns nil at line 1002. It simultaneously claims a validated ready identity and a rejected catalog. That is precisely the fabricated-ready-identity state the P6-B provenance contract forbids.

The fact that current authority constructors do not intentionally emit this combination is insufficient: this validator is the fail-closed acceptance boundary for site payloads, and the assigned review explicitly requires correct reason/proof/hash combinations plus rejection of conflicting, unknown, and incomplete payloads. A validator that accepts the contradiction does not meet that contract.

## Fix Direction And Required Re-review Evidence

Fix only this same P6-B invariant family:

1. Make the validator enforce an explicit accepted state/status/reason/hash matrix. At minimum, `ready` must reject every catalog-missing/rejection reason; resolved, profile-excluded, meaning-mismatch, and ready-catalog capability-unavailable records must each accept only their defined reason rules. Preserve the already-correct missing and rejected absence semantics.
2. Add direct table-driven negative tests through `validateTypeScriptAuthorityResult` or `finalizeTypeScriptAuthorityResults` for cross-state contradictions, including `ready` plus each catalog-failure reason with otherwise complete non-empty hashes. Retain positive rows for legitimate ready-catalog no-lib/config outcomes and all six missing/rejected failure classes.
3. Update the four ledgers and coder handoff so the fail-closed claim matches the exact implemented matrix. Keep P6-B unchecked until a new independent verdict.

Required re-review evidence is the narrow source diff, the accepted/forbidden matrix tests, preservation of the six authority/resolver/analyzer site-carriage rows, a proportionate production build/focused regression, unchanged artifact/boundary measurements unless bytes legitimately change, and current Anvien file-detail/impact/detect evidence. Canonical identity, receiver terminality, the exact ten-vector compiler matrix, valid-catalog carriage, and the six concrete generated failure paths need not be mechanically rerun unless the repair changes their bytes or owners.

## Current Anvien Evidence

- Repo-local packaged help was run before use and succeeded.
- Exactly one fresh explicit-path graph refresh was run: `anvien analyze E:\Anvien --force --json`.
- Current graph: `2,024` scanned / `752` parsed code / `0` failed / `116,113` nodes / `161,240` relationships; indexed/current commit both `5bfdfb3ea66f4a51c3efd44fc325abc80a317077`; graph stale=false.
- File-detail remains HIGH for each touched production owner: `resolve.go` (183 symbols, 110 inbound, 281 outbound, 23 flows, 36 linked tests), `indexes.go` (344, 324, 108, 19, 33), `types.go` (93, 82, 25, 1, 21), `analyze.go` (333, 92, 317, 20, 8), `catalog.go` (343, 151, 3, 1, 3), and `profile.go` (91, 9, 28, 1, 2). HIGH is a scope warning, not the rejection reason.
- `validateTypeScriptAuthorityResult` impact is CRITICAL: `4` impacted / `1` direct / `2` modules / `23` processes, spanning analyzer and resolution. CRITICAL is a scope warning, not an edit prohibition.
- Fresh explicit-path detect is CRITICAL: `35` affected symbols / `8` affected files; `307` changed symbols / `8` changed files. Resolution-gap delta is `177` entities / `180` occurrences (`125` analyzer-gap, `52` non-actionable); current Resolution Health total is `0`; semantic app-layer and functional-area fields are complete for all `116,113` nodes with no missing fields/sources. Counts and warning meaning are preserved without normalization.

## Build, Runtime, Regression, Benchmark, And Cleanup Disposition

- Source invalidity was established before trusting build/test claims. A fresh mechanical rerun of every already-current successful gate would not close the missing negative validation rule.
- Current handoff evidence records successful production `go build ./cmd/... ./internal/...`, both shipping launcher builds, hidden packaged-runtime build with cached Ladybug, runtime ensure, focused authority/resolver/built-analyzer tests, and affected package tests. Those results remain useful non-acceptance evidence.
- The broad root `go build ./...` still fails only on intentional non-buildable fixtures and is not treated as success. The official lifecycle build was prohibited and not run because it performs install/package scripts.
- The full regression command completed with exactly the two preserved accepted baselines: C# provider parity expected two accesses and got none; Dart expected two and got one. The earlier sanitized/isolated PATH probe is not treated as success.
- Packaged `--help` was independently rerun and succeeded. The final packaged binary was independently measured at `73,610,752` bytes, SHA-256 `E7F851FCEC43F7CB740707FF8BBE8CB89E8B741A49FB37E72ADFBF2332182D42`; delta from accepted P5-D `71,509,504` is `+2,101,248` bytes (`+2.9384%`). No packaged candidate was rebuilt or overwritten by this review.
- Existing product benchmarks remain time-bounded evidence: embedded load `156,659,129-164,050,414 ns/op`, `36,350,514-36,453,763 B/op`, `291,808-291,820 allocs/op`; global lookup `247.4-255.0 ns/op`, `192 B/op`, `1 alloc/op`; inherited member lookup `621.3-708.0 ns/op`, `216 B/op`, `4 allocs/op`. They were not rerun because the blocker is a source validation contract and the benchmark owners were not invalidated.
- All 20 disclosed dead/debug/build paths were independently checked and are absent, including all three fixture `.anvien` directories. The authoritative root `.anvien` from the mandatory graph refresh remains intentionally present.

## Invariant Disposition And Handoff

- Exact typed reason/provenance through `Authority.availability`: cleared.
- Explicit missing/rejected proof absence semantics and attempted-artifact carriage: cleared for generated records.
- Six catalog-failure classes through authority, direct resolver, and built `analyze.Run`, with site/counter equality, deterministic replay, and duplicate conflict rejection: cleared.
- Two-phase semantic identity, `14,587/14,587` uniqueness, component-drift/duplicate rejection: prior clearance retained.
- Local/repository/import receiver terminality and genuine external eligibility: prior clearance retained.
- Exact named ten-vector P6-A compiler matrix and valid-catalog carriage: prior clearance retained.
- Residual same-invariant surface: one—the validator's acceptance of `ready` identity paired with a catalog-failure reason.

The candidate must return to the existing P6-B coder for this narrow fail-closed matrix correction. Main must not finalize, stage, or create the isolated P6-B commit from this report. No code, ledger, old report, protected handoff, target, or later-slice owner was repaired or mutated by this Supervisor review.

