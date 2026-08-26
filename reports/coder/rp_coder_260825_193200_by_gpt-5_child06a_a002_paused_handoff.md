# Coder Handoff — Child 06A A002 Run-Scoped Diagnostic Appender

Date: `2026-08-25` (continued `2026-08-26`)
Status: `CODER_A002_A00X_SCRIPT_BUILD_READY_FOR_MAIN`
Attempt: `A002` on active unchecked `B2-P2A-A001-D001 resolve_calls`
Authority: architecture commit `dafcf7a25e254ef8db09d28eca3575d5225c553e`; plan commit `80eb8c3c9a9cf516dd17f8b3bf15802694bb1277`
Git reference: current HEAD `cc420e7ad719d90dc4b2d9991be0249e8d648daa`; A002 candidate remains uncommitted and the staged set is empty.

## Current result

The four authorized A002 source/test paths are written and validated. The definitive canonical full build completed with exit `0` before every test command. All five focused behavior/equivalence tests and the full `internal/graphhealth` package passed. The full `internal/resolution` package remains truthfully `FAIL` only on the known preserve-only `TestProofBasedCallAccessGoldenCorpus`; an exact current-HEAD production overlay reproduced the identical complete diagnostic payload, proving the failure is pre-existing rather than A002-specific.

`E2-P2A-A002SCRIPTBUILD1` is now complete. The reusable `scripts/build-a00x-benchmark.ps1` passed canonical full-build ordering, final PowerShell parse, two fail-closed no-output guards, and exactly one valid A002 build. Its frozen executable, pinned native DLL, and provenance JSON exist only beneath the required repository-local `.tmp` packet root. No target analyze/benchmark, detect, stage, commit, cleanup, or network action ran.

E2E Verification:

- `[PASS]` Compiled: canonical `scripts/full-build.ps1` -> exit `0`.
- `[PASS]` Runtime boundary: canonical runtime `1.2.8` completed maintenance `analyze . --force` with `2221` scanned, `767` parsed, `0` failed.
- `[PASS]` Happy path: repeated run-scoped appends match legacy diagnostics and graph JSON after every write.
- `[PASS]` Edge cases: nil graph, blank node ID, blank kind, absent node, duplicate merge, target fill, earliest line, ordering, and policy/default fields.

## Invariant family

- Family: per-`ResolveBoundInto` resolution diagnostic append semantics.
- Authority: `E2-P2A-A002ARCH1/PLAN1` and `reports/system-architect/rp_system-architect_260825_184751_by_gpt-5_child06a_a002_run_scoped_diagnostic_appender.md`.
- Primary path: `ResolveBoundInto -> newEmitter -> resolveCall -> emitUnresolvedReference|emitTypeScriptOutcomeDiagnostic -> run-scoped write-through appender -> Graph.AddNode`.
- Preserved sibling path: all non-resolution callers continue using `graphhealth.AppendDiagnosticToNode` unchanged at the call surface.
- Preserved rules: false-return/default behavior, supported first-touch representation normalization, diagnostic bucket identity, count merge, target fill, earliest-line behavior, stable ordering, policy fields, `[]Diagnostic` graph representation, and visibility after every append.
- Resource/lifecycle: one emitter-owned map of `nodeID -> []Diagnostic` slice headers; the graph property and cache receive the same merged slice; no second retained diagnostic-object copy, I/O, goroutine, lock, global cache, flush, or finalizer.

## Changed files and exact hunks

1. `internal/graphhealth/diagnostics.go`
   - Adds private `diagnosticAppender` state and exported internal-package factory `NewDiagnosticAppender`, returning the private append closure.
   - First successful node touch normalizes existing supported diagnostics once; every incoming diagnostic is defaulted and normalized once.
   - Every append re-reads the node, preserves sibling properties, writes `DiagnosticPropertyKey` as `[]Diagnostic`, and calls `Graph.AddNode`.
   - Refactors `appendDiagnosticsFromProperties` through private `appendNormalizedDiagnostic`; legacy `AppendDiagnosticToNode` remains available and behavior-compatible.
2. `internal/resolution/emit.go`
   - Adds the emitter-owned appender function field.
   - Initializes it in `newEmitter` from the same graph.
   - Routes `emitUnresolvedReference` through it.
3. `internal/resolution/outcome.go`
   - Routes `emitTypeScriptOutcomeDiagnostic` through the same emitter-owned appender.
4. `internal/graphhealth/diagnostics_test.go` (new)
   - Compares legacy and run-scoped writes after every append for encoded legacy first-touch input, multiple unique diagnostics, duplicate merge/count, target fill, earliest line, stable order, policy/default fields, graph JSON/write-through, and invalid/absent-node false behavior.

No existing test file, ledger, plan-rules file, graph type, policy/contract, `resolve.go`, persistence/reader, schema, public product contract, or timing instrumentation was edited.

## Pre-edit Anvien evidence

- Exact `anvien --help`: PASS.
- Fresh `anvien analyze --force`: exit `0`; scanned `2217`, parsed `765`, failed `0`; graph `123605` nodes / `170322` relationships; index fresh at commit `80eb8c3c9a9cf516dd17f8b3bf15802694bb1277`.
- File risk:
  - `internal/graphhealth/diagnostics.go`: HIGH; `154` symbols; inbound/outbound `50/70`; linked flows/tests `1/19`.
  - `internal/resolution/emit.go`: HIGH; `154`; `81/231`; `13/24`.
  - `internal/resolution/outcome.go`: HIGH; `119`; `87/124`; `11/23`.
- Symbol impact:
  - `AppendDiagnosticToNode`: CRITICAL; impacted/direct `7/2`; modules/processes `1/24`.
  - `appendDiagnosticsFromProperties`: CRITICAL; `7/1`; `2/13`.
  - `emitter`: CRITICAL; `28/25`; `2/49`.
  - `newEmitter`: CRITICAL; `6/1`; `4/32`.
  - `emitUnresolvedReference`: CRITICAL; `7/4`; `2/36`.
  - `emitTypeScriptOutcomeDiagnostic`: LOW `0/0`, while its file remains HIGH.

HIGH/CRITICAL were treated as blast-radius warnings; the candidate remains inside the exact Architect boundary.

## Source inspection completed

- Production-first order: PASS.
- Production diff inspection: PASS; scoped and surgical.
- `gofmt -d` on the four authorized paths: empty output before and after validation.
- Scoped `git diff --check` before build and after all tests: PASS.
- Final production delta remains exactly `+56/-4`: `diagnostics.go +52/-2`, `emit.go +3/-1`, `outcome.go +1/-1`.
- Final hashes match the paused handoff byte-for-byte:
  - `internal/graphhealth/diagnostics.go`: `AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8`.
  - `internal/resolution/emit.go`: `73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060`.
  - `internal/resolution/outcome.go`: `02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E`.
  - `internal/graphhealth/diagnostics_test.go`: `58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA`.
- Final route inspection confirms exactly two resolution routes use the emitter appender, while `AppendDiagnosticToNode` remains defined and available for legacy callers.
- Exact A002 source boundary is the three modified production files plus the new test file. This existing report is the only A002 documentation file updated. All pre-existing unrelated dirty/untracked work, including the interrupted Vite timestamp, was preserved and untouched.

## Build validation

Definitive command:

`pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`

Fresh resume preflight:

- `anvien doctor locks --repo E:\Anvien --json`: exit `0`, analyze lock free/absent.
- `anvien doctor processes --json`: exit `0`, only editor-owned MCP services; no proven build-output holder and nothing terminated.

Definitive result:

- Exit `0` / PASS.
- Vendor/runtime source verification: PASS (`1798` files, `0` missing, `0` extra, `0` mismatched).
- Canonical runtime and launcher/Vite build: PASS; Vite emitted only the known dynamic-import and chunk-size warnings.
- Runtime: `E:\Anvien\anvien\bin\anvien.exe`, version `1.2.8`, `73,767,936` bytes, SHA-256 `11895442CA4A5708FA439A787BC541102BE7144856315BCD17E9A7B772658AED`.
- Maintenance analyze: exit `0`; `2221` scanned, `767` parsed, `0` failed; graph `123811` nodes / `170572` relationships.
- Validation history remains truthful: the pre-pause invocation was Owner-interrupted and exit `1`; an initial resume invocation later lost its stdout host during Vite and had no remaining process/verdict, so neither was counted. The definitive full rerun above is the only PASS used to unlock tests.

## Tests

Tests ran only after the definitive full-build PASS, in the required order:

| Order | Command / test anchor | Result |
|---:|---|---|
| 1 | `go test ./internal/graphhealth -run '^TestDiagnosticAppenderMatchesLegacySemantics$' -count=1 -v`; `internal/graphhealth/diagnostics_test.go:11` | PASS, exit `0`; all subtests PASS |
| 2 | `go test ./internal/graphhealth -run '^TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity$' -count=1 -v`; `internal/graphhealth/p6d_outcome_projection_test.go:208` | PASS, exit `0` |
| 3 | `go test ./internal/resolution -run '^TestResolveAttachesSourceBackedUnresolvedDiagnostics$' -count=1 -v`; `internal/resolution/resolution_test.go:335` | PASS, exit `0` |
| 4 | `go test ./internal/resolution -run '^TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix$' -count=1 -v`; `internal/resolution/p6c3_structured_outcome_test.go:15` | PASS, exit `0`; all four matrix subtests PASS |
| 5 | `go test ./internal/resolution -run '^TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite$' -count=1 -v`; `internal/resolution/p6b_tsstdlib_test.go:266` | PASS, exit `0`; all six catalog subtests PASS |
| 6 | `go test ./internal/graphhealth -count=1` | PASS, exit `0` |
| 7 | `go test ./internal/resolution -count=1` | FAIL, exit `1`, solely `TestProofBasedCallAccessGoldenCorpus` at `internal/resolution/proof_accuracy_golden_test.go:12` / assertion line `166` |

Known-failure proof:

- Candidate failure payload classifies the callback diagnostic as `unclassified/review`.
- Exact baseline authority is current HEAD `cc420e7ad719d90dc4b2d9991be0249e8d648daa`.
- Repo-local overlay root: `E:\Anvien\.tmp\child06a_a002_current_head_baseline_cc420e7`.
- Overlay replacements are exact HEAD blobs and each `git hash-object` equals `git rev-parse HEAD:<path>`:
  - `diagnostics.go`: `5a978bbe287d7384ace314b6ff970192f89fe46a`.
  - `emit.go`: `d33f37f9ebfc71e41128c36e3c42f3d06a391fac`.
  - `outcome.go`: `a5a9ee80dd005318abab807137e7cff60fd97db0`.
- Command: `go test -overlay=E:\Anvien\.tmp\child06a_a002_current_head_baseline_cc420e7\overlay.json ./internal/resolution -run '^TestProofBasedCallAccessGoldenCorpus$' -count=1`.
- Baseline-overlay result: exit `1` with the identical full `graphhealth.Diagnostic` payload, including `Classification:"unclassified"`, `Actionability:"review"`, every source/range/site/proof field, encoded outcome note, and count.
- Classification: pre-existing preserve-only golden failure. The package is not called PASS; this failure is neither A002 regression nor A002 validation proof. No other failure occurred.

## E2-P2A-A002SCRIPTBUILD1 — Reusable A00x Benchmark Build Script

Status: PASS / CODER_A002_A00X_SCRIPT_BUILD_READY_FOR_MAIN.

### Exact implementation boundary

- Added only scripts/build-a00x-benchmark.ps1; final SHA-256 ADA407C7496FCEA988276F03BAD5001ED139A4AEC9A16B9C32947DA440814EC5.
- The script accepts explicit AttemptId, overlay manifest/output path, overlay hash, exact mapped-source hash set, candidate-source hash set, native hash set, and optional absolute Go executable.
- It fail-closes on repository identity, A000-style attempt ID, .tmp containment, reparse traversal, existing output, manifest shape and exact mapping-set equality, every input hash, go.mod/go.sum/vendor/native presence, and toolchain identity before any output/work directory creation.
- It builds only ./cmd/anvien with the canonical vendor, ladybugdb, trim, buildvcs, ldflags, overlay, and explicit output arguments; copies the pinned lbug_shared.dll; queries version; and writes provenance. It never invokes either canonical build script, writes canonical runtime/launcher/server/web outputs, registers protocol, or launches analyze.
- Child-only environment is recorded and does not mutate ambient process-global state: offline Go/CGO/native variables plus GOCACHE, GOMODCACHE, GOTMPDIR, TEMP, TMP, HOME, USERPROFILE, APPDATA, LOCALAPPDATA, XDG_CACHE_HOME, and XDG_CONFIG_HOME. Every cache/temp/home path is beneath E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\.a00x-A002-work.
- Final PowerShell parse: 4924 tokens and 0 errors. Two syntax-only interpolation defects found after the required full-build gate were narrowed to a delimited FilePath interpolation and a single-quoted formatted AttemptId message; no build contract or other file changed.

### Holder/lock preflight and canonical full build

- The first broad Restart Manager preflight transiently misclassified Windows Defender scanner MsMpEng PID 3840 as a termination target. Stop-Process returned Access Denied, no build started, and exact failure/recheck logs were preserved. The narrow persistent-holder recheck then identified the actual canonical-output holders.
- Windows Restart Manager proved PIDs 5564, 8320, 10888, 12552, and 7484 were actual holders of both E:\Anvien\anvien\bin\anvien.exe and lbug_shared.dll. Each was an installed editor-owned anvien mcp process, and only those five proven holders were terminated.
- Post-termination verification: all five PIDs gone; both canonical output files accepted exclusive read/write access; analyze lock free/absent; no repo-scoped Go, Node, esbuild, or launcher build process remained.
- Exact command: pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1.
- Start/end: 2026-08-26T02:52:20.7205211Z / 2026-08-26T02:58:30.0364325Z; exit 0; elapsed 369.3159114 seconds is validation only, not benchmark data.
- Vendor verification: 1798 files, 0 missing, 0 extra, 0 mismatched, 0 reparse points, and 0 incomplete license dispositions.
- Canonical runtime: version 1.2.8, SHA-256 11895442CA4A5708FA439A787BC541102BE7144856315BCD17E9A7B772658AED.
- Maintenance analyze: 2224 scanned, 767 parsed code, 0 failed. It is canonical full-build validation only; no target repository was analyzed.
- Exact build stdout: E:\Anvien\.tmp\child06a_a002_a00x_script_validation\canonical-full-build.stdout.log, SHA-256 ECEB81A4B6193677D365F44C3CDC77BB0EE1C177BD30CE2E980D63817057AED0.
- Exact build stderr: E:\Anvien\.tmp\child06a_a002_a00x_script_validation\canonical-full-build.stderr.log, SHA-256 E466327A2F3735BF6BB8F042070934BE6DF6C0295CC6B7C58FF0166C4C7B2E2D; only known Vite dynamic-import/chunk-size warnings.
- Result, preflight, guard, and packet-verification JSON artifacts are under E:\Anvien\.tmp\child06a_a002_a00x_script_validation.

### Fail-closed negative guards

- Bad overlay SHA-256: exact contract error reported expected all-zero hash versus actual 7B138FBA06B41CBD4C6709E6DF00C80C03E4015B615D1E42F787DE2A65D378F7. Guard root, executable, DLL, provenance, and .a00x-A002-work all remained absent.
- Output outside repository .tmp: exact containment error rejected E:\Anvien\child06a_a002_a00x_outside_guard_ada407c7\anvien-a002-outside.exe. Guard root, executable, DLL, provenance, and work root all remained absent.
- These two guards created no build artifact and do not count as the one valid build.

### Exactly one valid A002 invocation and frozen packet

- Valid invocation count exactly 1.
- Start/end: 2026-08-26T03:05:27.6669857Z / 2026-08-26T03:08:38.2045746Z; exit 0; elapsed 190.5375889 seconds is build validation only.
- Go: C:\Program Files\Go\bin\go.exe; go version go1.26.3 windows/amd64.
- Executable: E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\anvien-a002-benchmark.exe; version 1.2.8; 73,816,576 bytes; SHA-256 0F8B8244ABC80339A73E3A29F38D32F55A7A6A65281A64C291785ACF8F4A241E.
- Native DLL: E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\lbug_shared.dll; 20,230,656 bytes; SHA-256 20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7.
- Provenance: E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\anvien-a002-benchmark.provenance.json; 19,488 bytes; SHA-256 673FCCFEA17C48640C0AC62243479FE5EFD71E172CB1D4136D046C806E5855F1.
- Provenance verification PASS: schema 1, attempt A002, repo root/HEAD/status, UTC start/end, Go path/version, exact argv/flags, all scoped environment values, output identities, version, and exit 0.
- Overlay contract PASS: manifest SHA-256 7B138FBA06B41CBD4C6709E6DF00C80C03E4015B615D1E42F787DE2A65D378F7; exactly 2/2 mappings and no extra mapping.
  - internal/resolution/resolve.go replacement SHA-256 92A89227E1B1B1C159DE8BE77A1F060361C9058C4F83C78BFC32E9AB40DADEA9.
  - internal/resolution/types.go replacement SHA-256 A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8.
- Candidate contract PASS: exactly 3/3 expected hashes for diagnostics.go, emit.go, and outcome.go.
- Native contract PASS: exactly 3/3 expected hashes for lbug.h, lbug_shared.lib, and lbug_shared.dll.

### Boundary preservation

- Preserve-only hashes are unchanged for both canonical scripts, internal/cli/package_command.go, internal/cli/package_runtime.go, internal/cli/package_runtime_p6d_test.go, all four A002 production/test files, plan-rules.md, and all four standard Child 06A ledgers.
- E:\cheapapp.org remains at HEAD a869876ab6262dacde6cd5d432d099a91852a646 with exactly the same pre/post status rows.
- E:\Restaurant_manager remains at HEAD 605c0bda99491789e7f07628ec4b3d39a3ae1c67 with exactly the same pre/post status rows.
- Untracked-script no-index git diff --check output is empty; tracked report git diff --check passes; staged set remains empty.
- No target analyze/benchmark, detect-changes, staging, commit, cleanup, reset, checkout, stash, network/install, or subagent action ran.

## Residuals and boundary

- Residual unverified surfaces inside the assigned source/build/test scope: `none`.
- Known residual: the package-level golden remains FAIL/pre-existing/preserve-only as proven above.
- Performance/alloc/two-target measurement, per-attempt Supervisor, disposition, detect, staging, and commit remain pending by design and belong to Main or later lanes, not this Coder continuation. The four standard ledgers remain preserve-only and unchanged in this scope.
- No additional production owner was needed. D002-D017 and every forbidden owner remain untouched.
- No target benchmark/analyze performance capture, detect, stage, commit, push, reset, checkout, stash, cleanup, network/install, or subagent action ran.
- Staged set remains empty.

## Handoff status

`CODER_A002_A00X_SCRIPT_BUILD_READY_FOR_MAIN`
