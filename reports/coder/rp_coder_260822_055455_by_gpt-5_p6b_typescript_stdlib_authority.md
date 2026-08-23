# P6-B TypeScript Standard-Library Authority — Coder Handoff

Status: `READY_FOR_INDEPENDENT_SUPERVISOR`

CreatedAt: `2026-08-22T05:54:55.6069599+07:00`

This report is a coder handoff, not acceptance. P6-B remains unchecked, unstaged, and uncommitted until an independent Supervisor verdict and Main action.

## Authority and boundary

- Authoritative checkout: `E:\Anvien` only.
- Git branch/HEAD: `master` / `b98131e44932a7bcac17b487ecb2914535927d01`.
- Required parent: `ec765debff335540c77d409ebb2c9f45e4a0a77d`.
- Accepted P6-A authority: `reports/system-architect/rp_system-architect_260822_033607_by_gpt-5_p6a_declaration_universe_decision.md`, with independent Supervisor PASS recorded by the two P6-A Supervisor reports.
- Scope implemented: P6-B only. P6-C1 remains `PRESERVE_ONLY`; P6-C2 remains `ACTIVE` but locked; P6-C3/P6-D and `E:\cheapapp.org` were not opened.
- Prohibitions preserved: no network, install, package script, runtime Node/`tsc`, target access, stage, commit, push, alternate worktree, or internal subagent.
- Official protected Main handoff inventory is exactly `18`. The official transfer path `reports/Investigation/rp_main_260822_052426_orchestration_rotation_handoff.md` is included. Only pathnames were observed through Git status; all 18 were preserved and none was read, modified, removed, staged, or committed.

## Invariant Family Map

- Family: deterministic TypeScript standard-library declaration authority, supported root declaration profile, and repository-first resolver eligibility.
- SSOT: accepted P6-A offline TypeScript `5.9.3` compiler-API generation -> checked-in compact catalog -> Go `go:embed` runtime lookup.
- Primary path: exact offline compiler data -> sorted/hash-validated catalog -> supported root profile -> `analyze.Run` injection -> repository/P5 resolution first -> eligible TypeScript global/type/member miss -> external declaration lookup.
- Sibling surfaces closed: global value, global type, namespace/direct member, inherited member, external receiver member, absent-config default, explicit `target`, explicit `lib`, `noLib`, invalid/unreadable/topology config, catalog schema/version/hash/manifest/trailing-data failure, repository shadowing, explicit-import terminality, and non-TypeScript isolation.
- Forbidden fallback status: no project/package `.d.ts`, ambient module, augmentation, `node_modules` discovery, runtime Node/`tsc`, network/install/script, raw runtime declaration parse, target-name allowlist, synthetic `IMPORTS`, repository ownership pollution, `ExternalSymbol`, outcome DTO, persistence/reader projection, or graph-health change.
- Verify matrix: production catalog/profile/lookup, focused package behavior, build/package outputs, packaged CLI fixture, full regression classification, catalog determinism/metadata, microbenchmarks, binary-size delta, cleanup, fresh self-graph, and detect-changes.

## Implementation outcome

The final candidate implements the accepted mechanism without target-name branches:

- `anvien-web/scripts/generate-tsstdlib-catalog.mjs` is developer-only and checks exact TypeScript `5.9.3`, package-lock integrity, official `lib*.d.ts` inputs, stable ordering, meaning lanes, direct members, base types, value-owner links, catalog hash, and artifact hash.
- `internal/tsstdlib/catalog.v1.json` is the checked-in immutable catalog.
- `internal/tsstdlib/catalog.go` embeds and validates the catalog fail-closed, including rejection of a valid object followed by trailing JSON/token data; it provides global/member lookups with value/type/namespace meaning and inherited-member traversal.
- `internal/tsstdlib/profile.go` selects only zero/one root JSONC `tsconfig.json`; only `target`, `lib`, and `noLib` affect the declaration profile. Unsupported topology, invalid/unreadable config, `jsconfig.json`, nested/multiple config, and ownership ambiguity are capability-unavailable.
- `internal/analyze/analyze.go` creates the authority only for a scanned TypeScript repository and passes the complete scanned inventory to profile selection.
- `internal/resolution/types.go`, `indexes.go`, and `resolve.go` attach the authority after repository/P5 behavior. Explicit-import failure is terminal; non-TypeScript sites cannot consult the catalog.
- P6-B records only four authority counters. It deliberately does not materialize P6-C2 `ExternalSymbol` facts or implement P6-C3/P6-D outcome/projection behavior.

## Exact owned diff

Tracked modifications (`8`):

1. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`
2. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
3. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
4. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
5. `internal/analyze/analyze.go`
6. `internal/resolution/indexes.go`
7. `internal/resolution/resolve.go`
8. `internal/resolution/types.go`

New P6-B source/test/fixture assets (`16`):

1. `anvien-web/scripts/generate-tsstdlib-catalog.mjs`
2. `internal/analyze/p6b_tsstdlib_test.go`
3. `internal/analyze/testdata/p6b-tsstdlib-runtime/main.ts`
4. `internal/analyze/testdata/p6b-tsstdlib-runtime/tsconfig.json`
5. `internal/resolution/p6b_tsstdlib_test.go`
6. `internal/tsstdlib/catalog.go`
7. `internal/tsstdlib/catalog.v1.json`
8. `internal/tsstdlib/catalog_test.go`
9. `internal/tsstdlib/profile.go`
10. `internal/tsstdlib/testdata/profiles/es2022/tsconfig.json`
11. `internal/tsstdlib/testdata/profiles/es5-promise/tsconfig.json`
12. `internal/tsstdlib/testdata/profiles/es5/tsconfig.json`
13. `internal/tsstdlib/testdata/profiles/invalid-lib/tsconfig.json`
14. `internal/tsstdlib/testdata/profiles/jsonc/tsconfig.json`
15. `internal/tsstdlib/testdata/profiles/no-lib/tsconfig.json`
16. `internal/tsstdlib/testdata/profiles/topology/tsconfig.json`

This handoff report is the only additional report path. No notes file was created under the final Main boundary correction.

## Graph and blast-radius evidence

Pre-code graph at the exact anchor: `1,990` scanned / `743` parsed / `0` failed / `116,535` nodes / `160,659` relationships.

- `analyze.Run`: CRITICAL `24` impacted symbols / `8` modules / `23` processes.
- `workspace`: CRITICAL `134 / 7 / 69`.
- `BuildCrossFileBinding`: CRITICAL `95 / 24 / 31`.
- `resolveCall`, `resolveAccess`, `resolveTypeAnnotation`: each CRITICAL `28 / 7 / 32`.
- Existing file impacts: analyze CRITICAL `100` / `23` files; types CRITICAL `62` / `14`; indexes CRITICAL `234` / `35`; resolve CRITICAL `112` / `42`.
- Follow-up `catalog.go`: HIGH file detail; CRITICAL file impact `158` / `24` files. `loadCatalog`: CRITICAL `3` impacted symbols / `2` files / `2` modules / `12` processes.
- `catalog_test.go` and the edited validation test: LOW `0` impact.

HIGH/CRITICAL values were treated as blast-radius warnings requiring surgical edits and regression, never as edit bans.

Final source/cleanup graph refresh PASS: `2,006` scanned / `750` parsed code / `0` failed / `111,524` nodes / `156,259` relationships.

## Catalog, oracle, config, and provenance evidence

- Exact authority inputs: `100` official TypeScript `5.9.3` libraries / `3,141,835` bytes.
- Catalog inventory: `2,030` symbols / authoritative `11,802` direct member rows.
- Logical catalog hash: `22ab8f6986bfa8efd01addffac64b4265fff09244391ecb5d385b80dfa187119`.
- Artifact: `1,388,709` bytes / SHA-256 `78C9F6F49896D2F11E7E45B58C9D7CB90E1CA20413F06FFE9FAC95310C1AF054`.
- `anvien-web/package.json` and `package-lock.json`: zero diff.
- Accepted two-run generation equality remains `2/2`; the final loader EOF hardening changed neither generator source nor catalog bytes/hash. The erroneous PowerShell aggregation `15,838` is rejected and is not evidence.
- Generated profile closures match the accepted oracle: ES2022 `63 / 2,389,540`; ES2015 `13 / 310,653`; ES5 `3 / 232,957`.
- Meaning separation is data-derived: Promise type can be present where its runtime value is profile-excluded; Math value members resolve through direct/value-owner semantics.
- Invalid/unreadable/unsupported config never silently falls back to a broader catalog profile.

## Build and packaging evidence

Lock/process preflight on the exact checkout reported the analyze lock free/nonexistent and no lane-owned build/package process. Editor-owned MCP processes were not terminated.

- `go build ./...`: executed and FAIL only on known intentionally non-buildable language fixtures (fixture imports, mixed package, standalone C fixture). No P6-B compile failure.
- `go build ./cmd/... ./internal/...`: PASS.
- `go build ./...` in `anvien-launcher/src`: PASS.
- `go build ./...` in `anvien-launcher/server-wrapper`: PASS.
- `go run ../cmd/anvien package build-runtime` from `anvien/`: PASS without a package script.
- Final packaged runtime: `anvien/bin/anvien.exe`, `72,961,536` bytes, SHA-256 `E6EE2F57BAA0EF57531EA9B45B288174FE23AE7D9591CCDDE40EE27231FEFE41`, Go `1.26.3`, `ladybugdb` build tag.
- Accepted immediate P5-D baseline: `71,509,504` bytes. P6-B delta: `+1,452,032` bytes (`+2.03%`). No Owner threshold was invented.

The official repository lifecycle build was not executed because it performs package installation/package scripts, both explicitly prohibited. This is a disclosed boundary, not a false PASS.

## Behavior and regression evidence

Focused final command:

```text
go test ./internal/tsstdlib ./internal/resolution ./internal/analyze -count=1
ok internal/tsstdlib 1.434s
ok internal/resolution 0.307s
ok internal/analyze 1.594s
```

Direct test anchors:

- `internal/tsstdlib/catalog_test.go:10` — catalog metadata/default meaning lanes.
- `internal/tsstdlib/catalog_test.go:39` — compiler declaration-profile closures.
- `internal/tsstdlib/catalog_test.go:104` — fail-closed unavailable profiles.
- `internal/tsstdlib/catalog_test.go:133` — schema/version/hash/manifest/trailing-data rejection.
- `internal/resolution/p6b_tsstdlib_test.go:12` — eligible TS global/member/type misses resolve externally.
- `internal/resolution/p6b_tsstdlib_test.go:76` — repository result remains authoritative.
- `internal/resolution/p6b_tsstdlib_test.go:106` — explicit-import failure remains terminal.
- `internal/resolution/p6b_tsstdlib_test.go:134` — unavailable/excluded profiles are explicit, not repository gaps.
- `internal/resolution/p6b_tsstdlib_test.go:169` — non-TypeScript isolation.
- `internal/analyze/p6b_tsstdlib_test.go:14` — real analyze injection and graph-gap absence for Promise/Math.
- `internal/analyze/p6b_tsstdlib_test.go:53` — declaration inventory language isolation.

Nearest real boundary: packaged CLI analyze of `internal/analyze/testdata/p6b-tsstdlib-runtime` PASS at `2/1/0` files and graph `8/6`; resolver metrics were `ResolvedExternalDeclarations=6`, `ExternalCapabilityUnavailable=0`, `ExternalProfileExcluded=0`, `ExternalMeaningMismatches=0`. Promise type/value and Math.max were not gaps. One existing callback-local `resolve` repository-binding gap remained and is outside this declaration-authority invariant. Math.min is covered as ordinary catalog data by direct tests. No target repository was accessed.

Full production regression `go test ./cmd/... ./internal/... -count=1` completed. All packages passed except exactly two known out-of-scope baseline failures:

- C#: expected `ACCESSES=2`, got none.
- Dart: expected `ACCESSES=2`, got `1`.

No Child 04/05 baseline was changed to make this slice pass.

## Benchmark evidence

Final post-EOF microbenchmarks, three runs on Windows/amd64, Intel i7-3770:

- Embedded load: `63.645-66.741 ms/op`, `17,623,452-17,623,817 B/op`, `56,709-56,710 allocs/op`.
- Global lookup: `252.1-257.2 ns/op`, `192 B/op`, `1 alloc/op`.
- Inherited-member lookup: `613.6-625.1 ns/op`, `216 B/op`, `4 allocs/op`.
- Packaged self-analyze observation: `227,084.8119 ms`, max observed system bytes `1,563,363,832`; no comparable immediate P6-isolated baseline exists, so no attributable delta is claimed.

## Runtime dependency disclosure

Normal packaged-runtime analyze PASSed. Source/module/package inspection shows no runtime Node/`tsc`, network, install, package-script, `node_modules` discovery, or raw declaration parsing path.

A PATH restricted to the repo-local binary and Ladybug native directory failed before CLI startup with native exit `-1073741515` (`0xC0000135`) because the Ladybug DLL dependency chain was incomplete. This attempt is not counted as a no-Node PASS. It is an environmental evidence limitation disclosed for independent review; it did not invalidate normal packaged-runtime behavior or introduce a Node runtime path.

## Detect-changes evidence

- Exact requested registry-name invocation: `detect-changes --repo Anvien --scope all --json` -> exit `1`, fail-closed because registry name `Anvien` is ambiguous across two entries.
- Authoritative explicit-path invocation: `detect-changes --repo E:\Anvien --scope all --json` -> exit `0` without listing or reading the other repository.
- Summary: CRITICAL warning; `affected_count=23`, `affected_files=7`, `changed_count=143`, `changed_files=7`; affected app layers `backend=19`, `mixed=4`; affected functional areas `analyzer=8`, `mixed=6`, `resolution=9`; semantic app-layer and functional-area status complete.
- Detect's tracked-file view is supplemented by the exact untracked P6-B inventory above. `git diff --check` is clean and the Git index is empty.

## Cleanup and worktree preservation

Dry-run then exact cleanup removed only:

- `.tmp/p6b-graph-evidence/`
- `.tmp/p6b-analyze-benchmark.json`
- `.tmp/p6b-fixture-benchmark.json`
- `internal/analyze/testdata/p6b-tsstdlib-runtime/.anvien/`
- `anvien-launcher/src/anvien-launcher.exe`
- `anvien-launcher/server-wrapper/anvien-server-wrapper.exe`

All six were verified absent. No other `.tmp`, report, fixture, build output, protected handoff, tracked file, or index entry was removed.

## Blockers, risks, and residual surfaces

- Acceptance blocker: independent Supervisor verdict is pending. This report does not self-accept P6-B.
- Commit blocker: this lane is explicitly prohibited from stage/commit/push. `E6-P6B-COMMIT1` remains pending Main action after Supervisor PASS.
- Official full lifecycle build remains prohibited because it executes install/package scripts; the broad Go build was run and its known fixture failures are disclosed, while every production/shipping Go build passed.
- Direct sanitized-PATH no-Node execution remains unproved because the native loader failed first. No P6-B source/runtime dependency on Node exists; Supervisor may repeat this optional environmental isolation with a complete native DLL chain.
- Residual unverified surfaces within the P6-B behavior invariant: none. P6-C1/C2/C3/D, external target materialization, final outcome carriage, graph-health projection, and target validation are deliberately unimplemented/locked future slices, not residual P6-B work.

## Handoff

Independent Supervisor should review the exact 25-path owned candidate (8 tracked modifications + 16 new source/test/fixture assets + this report), compare it with the four living Child 06 ledgers and P6-A authority, verify all 18 protected Main handoffs remain untouched, and return PASS/REJECT for P6-B only.

Report identity is measured only after this file is frozen. Exact path, filesystem createdAt, bytes, LF, CR, strict UTF-8/BOM status, and SHA-256 are supplied in the external handoff; embedding a self-hash would change the bytes being identified. This file must not be edited after seal measurement.
