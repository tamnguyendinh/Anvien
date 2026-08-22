# Child 06 P6-C2 External Target Representation — Coder Handoff

## Terminal state

`READY_FOR_INDEPENDENT_SUPERVISOR`

This is a worker-validation handoff, not acceptance. P6-C2 remains unchecked, unstaged, and uncommitted until independent Supervisor review and Main closure. P6-C3 and P6-D remain locked.

## Authority and boundary

- Authoritative checkout: `E:\Anvien`
- Branch: `master`
- HEAD: `8055f0a6860721e26462572e34469e0d708d4a52`
- Parent / accepted P6-B commit: `5c1584fef7153a7a331c8fedd1ce64176ddc873d`
- Accepted P6-C1 state: preserve-only, independently accepted and committed in the current HEAD.
- Open slice: only P6-C2, `Add the necessary external target representation`.
- No access to the target repository, alternate worktree, project/package declaration lookup, P6-C3 final outcome DTO, P6-D graph-health work, public traversal options, network, dependency installation, or package scripts.
- No stage, commit, push, branch action, or internal subagent occurred.
- The existing four Child 06 post-commit ledger edits were preserved and extended; P6-C2 remains `[ ]`, while `E6-P6C2-REVIEW1` and `E6-P6C2-COMMIT1` remain pending.

The pathname-only protected inventory observed before this report is exactly `34` untracked paths: `29` Main handoffs, `3` prior P6-B Supervisor reports, and `2` older coder reports. Their contents were not read and none was edited, deleted, staged, or committed.

## Implemented invariant

The implementation materializes only final P6-B `TypeScriptAuthorityResult` rows whose status is `resolved`:

- One deterministic `ExternalSymbol` node is emitted per referenced `ResolvedSymbolID`; the catalog is never materialized wholesale.
- Node identity and provenance retain the semantic owner, TypeScript/catalog/profile/config authority, declaration library and declaration range, name, meaning, and source-site inventory required by current consumers.
- CALLS, ACCESSES, and USES edges retain every immutable per-site proof row and stable source-site ID. Normal graph coalescing unions and sorts those proofs instead of dropping duplicate endpoint facts.
- Duplicate identical facts canonicalize deterministically. Conflicting payloads for one semantic identity fail closed.
- `capability_unavailable`, profile-excluded, and meaning-mismatch results do not create a synthetic external target. Capability-unavailable remains an outcome/gap for later P6-C3/P6-D work.
- External nodes carry no repository `filePath` ownership and receive no File, DEFINES, IMPORTS, or CONTAINS relationship. No synthetic IMPORTS relationship was introduced.
- The P6-B lookup/precedence contract remains terminal before this materializer. No target-name special case was added.

Affected consumers were adapted only where source and impact evidence proved a semantic requirement:

- Ladybug receives a fixed `ExternalSymbol` node table, fixed columns, valid endpoint pairs, COPY loading, and native readback without fallback or skips.
- Process extraction excludes external CALLS endpoints from repository-owned process membership.
- MCP context and impact expose explicit external provenance with `external=true`, `editable=false`, and `repositoryOwned=false`.
- MCP rename rejects an external target as `external_symbol_non_editable` before edit collection.
- Generic graph JSON already transports arbitrary node facts and required no special-case change.
- `internal/mcp/impact.go`, HTTP/Web, file-detail, graph-health, and the P6-C3 shared final DTO were validate/preserve-only and remain unchanged.

## Exact owned candidate

Production (`9` paths):

1. `internal/scopeir/kinds.go`
2. `internal/resolution/external_symbol.go` (new)
3. `internal/resolution/emit.go`
4. `internal/resolution/resolve.go`
5. `internal/lbugload/csv.go`
6. `internal/lbugschema/schema.go`
7. `internal/processes/processes.go`
8. `internal/mcp/context.go`
9. `internal/mcp/rename.go`

Tests (`8` paths):

1. `internal/resolution/p6b_tsstdlib_test.go`
2. `internal/resolution/p6c2_external_symbol_test.go` (new)
3. `internal/analyze/p6b_tsstdlib_test.go`
4. `internal/lbugload/csv_test.go`
5. `internal/lbugschema/schema_test.go`
6. `internal/processes/processes_test.go`
7. `internal/mcp/p6c2_external_symbol_test.go` (new)
8. `internal/lbugnative/p6c2_external_symbol_integration_test.go` (new)

Ledgers (`4` paths):

1. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md`
2. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md`
3. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md`
4. `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md`

Handoff (`1` path):

1. `reports/coder/rp_coder_260822_162635_by_gpt-5_p6c2_external_target_representation.md` (this report)

The expected post-report owned boundary is therefore exactly `22` paths: `9` production + `8` tests + `4` ledgers + `1` report. Before report creation, Git reported `17` modified tracked paths and `4` owned untracked paths; the index was empty.

## Graph and impact evidence

The mandatory pre-production explicit-path refresh at HEAD `8055f0a6860721e26462572e34469e0d708d4a52` completed with:

- `2,034` files scanned
- `752` code files parsed
- `0` failures
- `116,282` nodes
- `161,454` relationships
- `19,426` dependency edges

Full file-detail and upstream file/symbol impact were run before each existing production owner was edited. HIGH/CRITICAL results were treated as blast-radius warnings and drove affected-boundary regression; they were not treated as edit prohibitions.

The most expansive shared label impact was `NodeLabel`: CRITICAL, `1,850` impacted symbols, `168` direct dependents, `47` modules, and `423` processes. File-level impact summaries reported affected-symbol / affected-file / affected-process counts of:

- `internal/resolution/emit.go`: CRITICAL, `57 / 26 / 13`
- `internal/resolution/resolve.go`: CRITICAL, `133 / 106 / 44`
- `internal/lbugload/csv.go`: CRITICAL, `94 / 32 / 31`
- `internal/lbugschema/schema.go`: CRITICAL, `86 / 38 / 45`
- `internal/processes/processes.go`: CRITICAL, `71 / 22 / 23`
- `internal/mcp/context.go`: CRITICAL, `65 / 41 / 21`
- `internal/mcp/rename.go`: MEDIUM, `17 / 7 / 2`; the exact rename symbol impact was LOW.

Resolver/emitter/finalizer, CSV/schema, process, and context symbol impacts were CRITICAL. The implementation was kept to the verified hook and consumer surfaces; the new isolated materializer had no pre-existing symbol to impact.

## Build and packaged artifact evidence

Rule 13 preflight found the analyze lock free and no process executable or command under `E:\Anvien`; no broad process kill occurred.

- Broad `go build ./...`: completed with exit failure only for intentionally non-buildable repository language fixtures (`models`, `animal`, mixed `animal/main`, and non-cgo `simple.c`). It is not labeled PASS and showed no P6-C2 compile failure.
- Production command build, `go build -o .tmp/p6c2/anvien-production.exe ./cmd/anvien`: PASS.
- Local packaged destination build from `E:\Anvien`, using the already pinned Ladybug `v0.19.1` input: PASS and wrote `E:\Anvien\anvien\bin\anvien.exe`.
- Packaged binary size: `73,653,760` bytes.
- Packaged binary SHA-256: `7B14655DBB65472CE83FBAF0677294CCB235EC09A34DB3AB877F5C5C1B12F568`.
- Accepted P6-B binary baseline: `73,613,824` bytes.
- Delta: `+39,936` bytes (`+0.05425%`). No Owner threshold exists, so this is a measurement rather than an acceptance threshold.

## Behavior, persistence, and reader validation

Production behavior was implemented before tests were changed or added.

- Focused P6-C2 rows in resolution, Ladybug CSV/schema, processes, and MCP: PASS.
- Built `analyze.Run` external-materialization boundary: PASS.
- Full affected packages `resolution`, `analyze`, `lbugload`, `lbugschema`, `processes`, and `mcp`: PASS.
- Targeted and full `-tags ladybugdb ./internal/lbugnative`: PASS.
- Native Ladybug readback: `2` node COPY operations, `1` relationship COPY operation, `0` fallback, and `0` skips.
- Built MCP context and impact returned the exact external provenance/non-editable facts.
- External process membership count: `0`.
- Dry-run rename returned `external_symbol_non_editable` before edit collection.

The packaged two-file real fixture produced:

- graph: `11` nodes / `10` relationships / `32,666` bytes
- referenced external nodes: `3`
- resolved source sites: `6`
- coalesced external edges: `4`
- immutable authority proof rows: `6`
- per-node source-site counts: `[2,1,3]`
- File / DEFINES / IMPORTS / CONTAINS repository-ownership pollution: `0 / 0 / 0 / 0`

Two clean force runs were byte-identical with graph SHA-256 `CB308956ECD89DDB801BC37ED72D01AD73656D2D01EDC9C49C3A8ADB0CC7920A`.

The broad `go test ./internal/...` command ran to completion and retained exactly nine classified failures outside the P6-C2 invariant:

- four CLI and three repository tests fail because mandatory repo-local temp directories inherit the parent Git repository;
- the pre-existing C# ACCESS parity baseline expected `2` and observed `0`;
- the pre-existing Dart ACCESS parity baseline expected `2` and observed `1`.

All other internal packages passed. These failures are not relabeled as P6-C2 passes, and passed focused/affected boundaries are not used to erase them.

## Final graph and change detection

After implementation, affected validation, packaged fixture cleanup, and ledger refresh, the final packaged destination analyze exited `0` with:

- `2,039` files scanned
- `756` code files parsed
- `0` failures
- `116,657` nodes
- `162,181` relationships
- `19,658` dependency edges

The root graph has `0` referenced external nodes under the current root config/topology; active representation proof comes from the packaged exact fixture above. This is a truthful configuration observation, not a missing-fixture substitution.

Explicit `anvien detect-changes --repo E:\Anvien --scope all --json` exited `0` at CRITICAL risk:

- affected: `44` symbols / `17` files
- changed: `202` symbols / `17` files
- affected layers: `api=16`, `backend=28`
- affected areas: `mcp=16`, `mixed=7`, `resolution=19`, `storage=2`
- changed layers: `api=11`, `backend=119`, `backend_test=53`, `docs=19`
- changed areas: `analyzer=32`, `documentation=19`, `mcp=11`, `providers=38`, `resolution=55`, `storage=47`
- ResolutionGap delta: `33` entities / `33` occurrences
- gap actionability: `analyzer_gap=20`, `non_actionable=13`
- gap classification: `builtin=12`, `in_repo_unresolved=20`, `standard_library=1`
- gap fact families: `access=14`, `call=14`, `type-reference=5`
- Resolution Health total: `0`
- semantic app-layer/functional-area completeness: `116,657 / 116,657`

CRITICAL is retained as the complete blast-radius warning. The final graph/detect evidence describes the source bytes at invocation time; writing ledgers and this report afterward does not trigger a self-referential rerun.

## Cleanup and worktree integrity

- Exact runtime/fixture tree `.tmp\p6c2` was removed after validation.
- Final isolated home tree `E:\Anvien\.tmp\p6c2-final-home` contained only its lane-created `registry.json`. An Anvien lock check was free, the process inventory contained no command line referencing that tree, the file and empty directory were deleted, and a final existence check returned `false`.
- No unrelated cleanup occurred.
- Pre-report `git diff --check`: PASS.
- Pre-report index entries: `0`.
- Pre-report candidate: exactly `21` owned paths; protected history remained a separate pathname-only inventory of `34` paths.

## Residual gates and next owner

There is no known P6-C2 implementation blocker. Required remaining authority gates are external to this coder lane:

1. Independent Supervisor reviews the exact 22-path candidate and returns PASS or a bounded rejection.
2. Only after PASS, Main updates the closure ledger, checks P6-C2, and creates the isolated commit.

P6-C3 and P6-D must remain locked until Main closes P6-C2. This report makes no acceptance or commit claim.
