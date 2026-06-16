# Community Singleton MEMBER_OF Evidence Ledger

## Metadata

- Date: `2026-06-16`
- Plan: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-plan.md`
- Evidence: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-evidence.md`
- Benchmark: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-benchmark.md`
- Actual status: `docs/plans/2026-06-16-community-singleton-memberof/2026-06-16-community-singleton-memberof-actual-status.md`

## E0 - P0 Evidence

- `E0-P0A-REPRO1`: `anvien analyze --force --progress --verbose` on HEAD failed after `db_load` with `db load skipped 506 relationships; refusing incomplete graph load`.
- `E0-P0A-GOOD1`: Detached worktree at `7e579821` analyzed successfully with `nodes=82760`, `relationships=120886`.
- `E0-P0A-TRIGGER1`: HEAD `a17f2fcd` analyzed successfully when excluding `internal/aicontext/skills/cyber/**`, proving the new cyber skill tree is the trigger.
- `E0-P0A-DIAG1`: Diagnostic graph pass found all 506 skipped relationships are `MEMBER_OF` with `missing_target_node`; endpoint pairs were `Function-><missing>` 462, `Method-><missing>` 42, `Class-><missing>` 1, `Interface-><missing>` 1.
- `E0-P0A-SRC1`: Source inspection found `communities.Apply` emits `MEMBER_OF` before `if len(members) < 2 { continue }`, so singleton partitions create edges to community nodes that are never added.

## E1 - P1 Evidence

- `E1-P1A-IMPACT1`: Pre-edit Anvien impact/file-detail unavailable. `anvien status` reported `Repository not indexed`; `anvien impact symbol "Apply" --repo Anvien --direction upstream`, `anvien impact file "internal/communities/communities.go" --repo Anvien --direction upstream`, and `anvien file-detail "internal/communities/communities.go" --repo Anvien --json` all failed with `repository "Anvien" is not registered`. This is the same blocked-index condition caused by `E0-P0A-REPRO1`.
- `E1-P1A-SRC1`: `internal/communities/communities.go` now checks `len(members) < 2` before constructing `communityID`, appending memberships, or emitting `MEMBER_OF`; non-singleton communities still create the `Community` node and then membership relationships.
- `E1-P1B-TEST1`: `internal/communities/communities_test.go` extends `TestApplySkipsSingletonCommunities` to fail if any singleton `MEMBER_OF` relationship or `Community` node is emitted.
- `E1-P1B-TEST2`: `go test ./internal/communities -count=1` passed.

## E2 - P2 Evidence

- `E2-P2A-BUILD1`: `go build ./...` failed before validation because repository fixtures under `test/fixtures` are intentionally non-buildable as Go packages (`package models is not in std`, mixed package names, and C source files outside cgo package). This is unrelated to `internal/communities`.
- `E2-P2A-BUILD2`: Product-scope build `go build ./cmd/... ./internal/... ./contracts/...` passed; `./contracts/...` matched no packages.
- `E2-P2A-TEST1`: Broad suite `go test ./cmd/... ./internal/... -count=1` timed out after 600s, so it is not counted as pass evidence.
- `E2-P2A-RUNTIME1`: `go run ..\cmd\anvien package build-runtime` from `anvien/` passed and wrote `anvien\bin\anvien.exe` using LadybugDB native runtime `.tmp\ladybug-native\v0.17.1\windows-x86_64`.
- `E2-P2A-ANALYZE1`: Source-built `.\anvien\bin\anvien.exe analyze . --force --progress --verbose --benchmark-json .\.tmp\community-singleton-memberof-analyze-after.json` passed through `db_load done`; benchmark reported `dbLoad.skippedRelationships = 0`, `nodeRows = 272114`, and `relationshipRows = 321335`.
- `E2-P2A-IMPACT1`: Post-fix impact on `Function:internal/communities/communities.go:Apply#1` reported `risk = CRITICAL`, `impactedCount = 3`, affected files `internal/analyze/analyze.go`, `internal/cli/command.go`, and `internal/graphaccuracy/access_candidate.go`; this is pipeline blast radius, not a blocker.
- `E2-P2A-DETECT1`: `.\anvien\bin\anvien.exe detect-changes --repo Anvien --scope all --json` reported two changed files: `internal/communities/communities.go` high file risk and `internal/communities/communities_test.go` low file risk; affected flows were empty.
- `E2-P2A-FULLBUILD1`: `npm run full-build` completed npm install/package runtime/global install/version and web build, then failed during launcher rebuild because Windows still held `anvien\bin\anvien.exe` after validation. `anvien doctor processes --json` and `Get-Process` confirmed a transient local binary process; repo analyze lock was free.
- `E2-P2A-FULLBUILD2`: After the transient lock exited, `powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1` passed.
- `E2-P2A-ANALYZE2`: Global CLI `anvien analyze . --force --benchmark-json .\.tmp\community-singleton-memberof-analyze-global-after.json` passed; benchmark reported `dbLoad.skippedRelationships = 0`, `nodeRows = 272114`, `relationshipRows = 321335`, `communitiesEmitted = 1675`, and `membershipsEmitted = 5948`.

## Closure Evidence

- `E-CLOSE-1`: The original failure mode `db_load phase: db load skipped 506 relationships; refusing incomplete graph load` no longer reproduces on the current repo with the global CLI.
