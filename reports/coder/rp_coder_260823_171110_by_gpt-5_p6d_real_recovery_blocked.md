# Coder Report: P6-D Real Recovery

Status: `BLOCKED`

## Metadata

- Time: `2026-08-23 17:11:10 +07:00`
- Role: Coder
- Repository: `E:\Anvien`
- Branch / HEAD / parent: `master` / `741335f7efa5ad215ec28361d88c462f431565ab` / `838f379ab00c93dfe7507603f6608bb08d61f038`
- Active slice: P6-D only; P6-E remains locked.
- Target: not accessed. This lane stopped before every `E:\cheapapp.org` action.
- Git authority: read-only; no stage, commit, branch, stash, reset, checkout, clean, push, or tracked-file restoration.
- Result: no production, test, build-script, package, fixture, plan, ledger, or protected report byte was changed by this lane. This report is its sole durable artifact.

## Outcome

The independently cleared six-field graph-health repair is preserved unchanged. A policy-compliant current product runtime cannot be produced from current bytes inside the assigned authority because the repository has neither a complete clean E-only Go dependency closure nor a lane-owned Ladybug native build bundle, and the canonical/package/launcher sources route native build input through the forbidden shared root `E:\Anvien\.tmp\ladybug-native`.

The canonical full build was therefore **not started**. Running it would have crossed the explicit stop boundary before it could produce a current runtime. No raw Go build, stale packaged binary, global runtime, shared/quarantined cache, unit test, or partial command is substituted for this missing gate.

## Exact Git / Worktree Boundary

Pre-report `git status --porcelain=v2 --branch` showed an empty index and exactly these current dirty paths:

- Owner-owned preserve-only deletion: `CONTRIBUTING.md`.
- Main/Planner-owned read-only paths: roadmap plus the four Child 06 ledgers.
- Independently cleared P6-D paths: `internal/graphhealth/diagnostics.go` and untracked `internal/graphhealth/p6d_outcome_projection_test.go`.
- Protected untracked handoffs: `reports/Investigation/rp_main_260823_164621_orchestration_rotation_handoff.md` and `reports/Supervisor/rp_supervisor_260823_164652_by_gpt-5_p6e_current_plan_acceptance_rereview.md`.

Preserved P6-D identities:

- `internal/graphhealth/diagnostics.go`: `26,884` bytes, SHA-256 `518BD06FF5F859BD42AD7B3B9FF8E3F9F4E7235720BC5E568E2C4B2DA72BD2AE`.
- `internal/graphhealth/p6d_outcome_projection_test.go`: `19,336` bytes, SHA-256 `6E90E655FFCCF5A5BD3D398D95EB5B6B7F77754CAAF35DAAF4D9609248A4E713`.

No lane action rewrote or re-audited their already-cleared semantic invariant.

## Invariant Family Map

- Family: current P6-D outcome projection through current product runtime, native Ladybug persistence, and actually affected built readers.
- Authority / SSOT: current Child 06 plan/evidence/benchmark/actual-status set; current P6-D Supervisor REJECT; current P6-E plan-acceptance PASS only for its still-locked plan.
- Preserved production path: P6-C3 structured outcome -> `internal/graphhealth/diagnostics.go` normalization/projection -> Graph JSON / ResolutionGap -> graph-health summary/readers.
- Sibling surfaces requiring current built evidence: canonical full-build final analyze; native Ladybug; CLI `graph-health summary/report/components/explain/files`; CLI `resolution-inventory`.
- Forbidden fallbacks: existing/global/stale runtime; shared/protected/quarantined module or Ladybug cache; target-text inference; target access; raw Go build as full-build substitute; new build workflow or package/build-script edit.
- Verify matrix disposition: local six-field source/test invariant preserved; canonical build/current runtime BLOCKED; native Ladybug BLOCKED; affected built readers BLOCKED; target/oracle/boundary not attempted and remain future QA gates after current-runtime recovery.

## Canonical Build Source Proof

`scripts/full-build.ps1` is the only repository canonical full-build entrypoint. It performs, in order:

1. `npm install` in `anvien`.
2. `npm run build`.
3. `npm install -g .`.
4. `Get-Command anvien`.
5. `anvien version`.
6. `anvien-launcher\build.ps1`.
7. `anvien version`.
8. `anvien analyze . --force`.

Root `package.json` maps `full-build` only to that script. Repository search found no second full-build implementation. `scripts\qa-child02-p2d-repeated-analyze.ps1` calls `npm run full-build`; it is a QA consumer, not an alternative.

The partial repository-owned helpers cannot satisfy the assigned E-only boundary:

- `anvien/package.json` maps `build` to `go run ../cmd/anvien package build-runtime` and its lifecycle scripts invoke that same runtime builder.
- `internal/cli/package_runtime.go` builds the CLI with `-tags ladybugdb`, but `resolvePackageNativeDir` fixes the native output root to `<sourceRoot>\.tmp\ladybug-native`.
- `anvien-launcher/build.ps1` independently fixes native output to `E:\Anvien\.tmp\ladybug-native`, writes canonical binary/launcher/server/web outputs, and registers the launcher protocol.
- `scripts/ensure-ladybug-native.ps1` accepts an output root, but it downloads the Ladybug release whenever the full `{lbug.h, lbug_shared.lib, lbug_shared.dll}` set is absent. Network is forbidden, and composing this partial helper into a new full-build workflow is outside this lane.

Thus there is no existing repository-owned full-build equivalent that both performs all canonical product steps and keeps all new native/cache/runtime output inside the required unique lane root.

## Current Runtime Provenance

The existing `E:\Anvien\anvien\bin\anvien.exe` is inspect-only and rejected as evidence:

- Size: `73,756,672` bytes.
- SHA-256: `6BB431D00C20A6E9954039A54563855561E4A77B5D0AF323A4A38A54C7070B80`.
- Exact `version` execution: exit `0`, `79 ms`, product version `1.2.8`.
- Embedded Go/VCS identity: revision `838f379ab00c93dfe7507603f6608bb08d61f038`, `vcs.modified=true`; current HEAD is `741335f7efa5ad215ec28361d88c462f431565ab`.
- Embedded CGO inputs: `-I E:\Anvien\.tmp\ladybug-native\v0.19.1\windows-x86_64` and the matching shared-root linker path.
- Adjacent `lbug_shared.dll`: `20,230,656` bytes, SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`.

This binary is stale relative to current HEAD and explicitly derived from a forbidden shared native root. It was not used for analyze, native parity, built-reader parity, or target proof.

## Clean E-only Dependency Evidence

Repository inventory found:

- no root `vendor` directory;
- no root or `anvien` `node_modules` directory;
- `anvien-web/node_modules` exists but does not supply the Go or Ladybug native closure;
- existing `anvien/bin` contains only runtime metadata, the stale executable, and the DLL—no `lbug.h` or `lbug_shared.lib`.

A single clean probe used only the lane-owned root `E:\Anvien\.tmp\p6d-real-recovery-01a02e0d-260823-170629`, with new `HOME`, `TEMP`, `GOCACHE`, `GOMODCACHE`, and `GOPATH`, plus `GOPROXY=off`, `GOSUMDB=off`, and `GOTOOLCHAIN=local`:

`go list -mod=mod -deps ./cmd/anvien`

The probe emitted `module lookup disabled by GOPROXY=off` for the exact required module classes: Cobra `v1.10.2`, Hugot `v0.7.2`, go-tree-sitter `v0.25.0`, Dart and Swift grammars, Kotlin, and the C/C#/C++/Go/Java/JavaScript/PHP/Python/Ruby/Rust/TypeScript grammar modules. The tool yield ended before the wrapper printed its planned final exit JSON; afterward no process named `go` remained. The output is used only as missing-dependency evidence, never as a build result.

## Rule-13 / Build Execution Disposition

- Builds started by this lane: `0`.
- Canonical build command: `NOT RUN / BLOCKED BEFORE START`.
- Rule-13 preflight attached to a build: N/A because no build was started. No holder or lock was terminated.
- Reason: canonical execution would use the forbidden shared native root and local/global install/output paths before a policy-compliant current runtime could exist; a fresh clean module cache also lacks the required dependency closure.
- Full-build step exits/timings/runtime/final analyze: unavailable because the command was correctly not started. No PASS/FAIL substitution is claimed.

## Native Ladybug And Built-Reader Parity

- Native Ladybug: `BLOCKED / NOT RUN`. A clean lane-owned header/import-library/DLL bundle does not exist, and the stale/shared bundle is forbidden.
- `graph-health summary/report/components/explain/files`: `BLOCKED / NOT RUN` on a current built runtime.
- `resolution-inventory`: `BLOCKED / NOT RUN` on a current built runtime.
- Generic unit-level Graph JSON/ResolutionGap/summary evidence remains prior accepted local evidence only; it is not promoted to native or built-reader parity here.
- Target, independent TypeScript oracle, and target pre/post boundary: `NOT RUN` by explicit stop boundary.

## Code / Anvien / Validation Disposition

- Same-slice production defect newly proved: none.
- Production edit: none.
- Test edit: none.
- Fresh file-detail/impact: N/A because this lane edited no function/class/method/exported/shared-contract owner and deliberately did not reopen the cleared invariant.
- Full build before validation: blocked before start; therefore no post-build validation was permitted.
- Fresh `anvien detect-changes`: N/A for this lane because it made no implementation change and has no stage/commit authority. The required implementation-commit detect remains Main-owned after a future complete P6-D gate.
- Blast radius: preserved prior P6-D warning—`diagnostics.go` containing-file impact is CRITICAL even though the six-field helper itself was LOW. No new blast-radius claim is made from stale graph evidence.

## Cleanup Disposition

The lane created only `E:\Anvien\.tmp\p6d-real-recovery-01a02e0d-260823-170629` for the clean dependency probe. One recursive cleanup attempt targeted exactly that verified lane-owned root beneath `E:\Anvien`; execution policy rejected it before process creation. No retry, alternate deletion, broad cleanup, or workaround occurred.

Current exact root inventory is `638` files / `1,025,772` bytes / `17` zero-byte files / `0` reparse points. It is debug/cache output only, not Git candidate, product runtime, dependency authority, or evidence input. Every prior protected/shared/quarantined root was preserved and not reused.

## Remaining Gates And Smallest Next Responsible Action

Remaining P6-D gates are exact:

1. A policy-compliant canonical full build or an already repository-owned E-only equivalent covering all eight product steps.
2. A fresh current runtime with direct current-byte provenance and final analyze completion/liveness evidence.
3. Native Ladybug plus all actually affected built-reader parity.
4. Later, in a fresh QA lane with target authority: target pre-state, exact `Promise` / `Math.max` / `Math.min` `3/3`, independent oracle, and target post-state.
5. Fresh clean independent Supervisor review, current detect, cleanup disposition, and isolated Main-owned commit.

Smallest next responsible action for Main/Owner: provide a clean non-quarantined E-only dependency closure for the exact `go.mod/go.sum` set and a clean Ladybug `v0.19.1` `{lbug.h, lbug_shared.lib, lbug_shared.dll}` bundle under a new lane-owned root, **and** provide an existing repository-owned same-step invocation that accepts that unique native/cache/runtime root. If no such invocation exists, explicitly route the hard-coded shared-output limitation to the owner of the canonical build mechanism; it cannot be repaired opportunistically inside P6-D.

No P6-D code change is requested. The six-field repair remains preserved and independently cleared.

## E2E Verification

```text
E2E Verification:
  [BLOCKED] Compiled: canonical scripts/full-build.ps1 or same-step E-only equivalent -> not started; clean dependencies/native inputs and unique-root execution contract absent
  [BLOCKED] Runtime: fresh current anviens runtime -> absent
  [BLOCKED] Happy path: native Ladybug + affected built readers -> not run without current runtime
  [N/A] Edge case: target/oracle/pre-post -> intentionally not accessed in this Coder lane
```

Final handoff disposition: `BLOCKED`.
