# Supervisor Report: P6-A Declaration-Universe Decision

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260822_041014_by_gpt-5_p6a_declaration_universe_decision.md`
- Review time: `2026-08-22 04:10:14 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`
- Scope reviewed: P6-A decision candidate at HEAD `ec765debff335540c77d409ebb2c9f45e4a0a77d`, parent `fcc44334c0f75b3b19046dc8f9f4de40eb459fa9`; four Child 06 ledgers, the Architect/Planner report, the tracked contract deviation, current source/config/package/runtime, and fresh Anvien evidence.
- Claim reviewed: offline TypeScript `5.9.3` compiler-API generation to a checked-in versioned compact catalog embedded with Go; narrow fail-closed root-config support; repository/P5 precedence; P6-C1 preserve-only; P6-C2 referenced-only `ExternalSymbol`; P6-B/C/D locked.
- Authority used: latest Main sealed transfer in task `01a02609-4c95-7e61-b3d9-1387ab847a7b`; `AGENTS.md`; full required skill chain; four Child 06 ledgers; `docs/contracts/graph-accuracy-contract.md`; current source/package/runtime; accepted Child 05 closure.
- Related artifact: `reports/system-architect/rp_system-architect_260822_033607_by_gpt-5_p6a_declaration_universe_decision.md`.

## Executive Summary

- Problem: decide whether the P6-A candidate is technically complete and lies wholly inside the latest sealed transfer boundary.
- Decision: the architecture decision itself is source-backed and its graph/package/compiler evidence reproduces independently, but the current transfer contains a required out-of-boundary contract mutation. That mutation also embeds a stale report hash, so current authority and artifact identity disagree.
- Required outcome: remove the contract mutation from the sealed candidate transfer, or obtain a new explicit Main boundary that authorizes the contract and then make its artifact identity truthful. Re-review is required; no P6-B/C/D work opens from this result.

## Blocking Findings

### [HIGH] Out-of-boundary contract mutation carries a false immutable artifact identity

File: `docs/contracts/graph-accuracy-contract.md:68`

Issue: the latest Main transfer permits exactly the four Child 06 ledgers plus one Architect/Planner report. Current tracked reality additionally changes the graph-accuracy contract by `18+/0-`. At line 84 that new authority says the immutable report hash is `56F513DB5A2121E305CD8D4F9B3CD41ACA9471A54842D300289740FDC7C84BE3`; the actual sealed report is SHA-256 `77D5E9AC8D76D98C76D1816C8D6E69265D4AFB30367E3DA50DF3EAA3445D7BA2`.

Evidence:

- `git diff --name-status`, `git diff --numstat`, and full patch inspection show five tracked files: contract `18+/0-`, plan `42+/27-`, actual status `44+/20-`, evidence `23+/7-`, benchmark `13+/5-`; total tracked diff `140+/59-`.
- The permitted four-ledger candidate portion is `122+/59-`; the contract is the only additional tracked path.
- The report is `25,326` bytes, `289` LF, `0` CR, strict UTF-8, no BOM, with final SHA-256 `77D5...D7BA2`.
- The candidate itself discloses Main's boundary intervention at report line 266. Disclosure proves awareness; it does not authorize the mutation.

Why this blocks acceptance: the latest sealed transfer is the governing authority. Accepting a fifth tracked path would accept work outside that authority, and the stale hash would make an active contract point to a non-matching immutable artifact. Both conditions require follow-up, while acceptance requires none.

Fix direction: restore `docs/contracts/graph-accuracy-contract.md` to its HEAD `ec765de...` content for this candidate transfer, without changing the four ledgers or final Architect report. If Main instead wants the contract included, Main must explicitly reseal a new boundary and the owning lane must update the contract to the final report identity before re-review.

Re-review evidence required: contract has no diff under the current sealed boundary; `git diff --name-status`, `--numstat`, and `--check` show only the four authorized ledgers; candidate report identity remains exactly `25,326/289/0/strict UTF-8/no BOM/77D5...D7BA2`; index remains empty; protected handoffs remain untouched.

## Source-Level Clearance Notes

- `internal/analyze/analyze.go:332`, `internal/resolution/types.go:8`, `internal/resolution/indexes.go:75`: clear. Analyze passes repository ScopeIR to a resolver whose options have only two compatibility booleans; there is no declaration profile/catalog input.
- `internal/resolution/resolve.go:360`, `:554`, `:638`, and `internal/resolution/emit.go`: clear. Repository scope/import/member/type paths precede generic unresolved emission; explicit-import claim checks terminate fallback; no external authority/outcome exists today.
- `internal/graphhealth/diagnostics.go:233` and `internal/graphhealth/resolution_gap_inputs.go:224`: clear. Health currently classifies from `TargetText` and Go-oriented maps, so mechanical structured-outcome projection is correctly assigned to P6-D.
- `internal/lbugload/csv.go:96` and `internal/lbugschema/schema.go`: clear. Unknown labels increment `SkippedNodes`; node tables and endpoint pairs are fixed. Dedicated Ladybug support for `ExternalSymbol` is necessary.
- `internal/processes/processes.go:256`: clear. The call graph currently accepts all qualifying `CALLS` edges without endpoint-label filtering; external endpoints must be excluded.
- `internal/mcp/context.go:289`, `internal/mcp/impact.go:710`, `internal/mcp/rename.go:91`: clear. Context/impact are repository-shaped generic projections; rename has no external-label rejection before edit collection. The C2 consumer map is justified.
- `internal/graph/types.go:39`, `internal/httpapi/graph.go:156`, `internal/filecontext/context.go`: clear/preserve-only. Graph JSON and HTTP are generic graph transports; no P6-A evidence requires a speculative shared graph-struct or UI redesign.
- Production/test/fixture/config/package/target files: no diff. No accepted Child 05 implementation or target evidence was reopened.

## Evidence Checked

Passed:

- Identity/boundary: HEAD/parent exact; branch `master`; origin comparison is `53` ahead, `0` behind; index empty; `git diff --check` clean before this report; exactly 16 protected Main handoffs plus one candidate report were untracked. Protected handoffs were named by status only and never read or written.
- Artifact seals: plan `DBDB76364CB2743A3C86C671766E17D211DB50A2FF8985905AD6200CB9B677C6`; actual `A2FA16014A8C24DDA6E2C3C68CFC4FA7782AB9CAC41517F38410B0D9D29B0BFA`; evidence `FB7AAB4EE76E346C410FD1C220F7D180BDB4E35C040DD64D53A990258FA6F634`; benchmark `BC77CFD6066B09E34D553F2E2717C2C1D7B33151C8E6DC663B1815B8E989E3A7`; contract `E198CAD4EF61655DBA42B52AF50DB68D7BB34A4B06DC6CF5D0734DF753412F9A`.
- Four-ledger consistency: P6-A is decision-recorded but unaccepted; P6-B/C1/C2/C3/D remain locked; P6-C1 is preserve-only; P6-C2 is active; review/commit IDs remain pending; benchmark contains only decision measurements and later gates.
- Fresh Anvien: `anvien analyze --force` at `03:59:24 +07` scanned `1,987`, parsed `743`, failed `0`, produced `116,502` nodes and `160,626` relationships, indexed current commit, `stale=false`. The `+2` files and `+35/+35` facts versus the candidate's earlier `1,985/116,467/160,591` are later report/doc state; parsed-code count is unchanged.
- File-detail: current `stale=false` related-file counts reproduce all claimed values: `187, 34, 58, 60, 50, 256, 250, 23, 22, 31, 30, 41, 14, 28, 41, 9, 22, 46` for the named analyze/resolution/graph/Ladybug/health/process/MCP/HTTP/filecontext owners.
- Blast radius with `--include-tests`: CRITICAL `analyze.Run 24/9/8/23`; `buildWorkspace 59/24/8/23`; `resolveCall`, `resolveAccess`, `resolveTypeAnnotation` each `28/11/7/32`; `emitUnresolvedReference 9/4/2/34`; `graph.Node 1717/273/48/428`; `scopeir.NodeLabel 1846/291/47/422`; `graphhealth.Diagnostic 70/17/5/77`; `nodeCSVRow 25/11/1/12`; `SourceBackedResolutionGapInputs 13/7/1/15`; `processes.Apply 29/8/7/31`; `buildCallsGraph 28/8/6/25`; `contextSymbolPayload 1/1/1/7`; `impactItemPayload 11/5/2/17`; `collectRenameChanges 1/1/1/8`. HIGH `classifyDiagnostic 10/4/3/0`. MEDIUM `NodeSchema 8/6/2/0`.
- TypeScript/package: declared range `^5.4.5`; lock and installed package exact `5.9.3`; integrity `sha512-jl1vZzPDinLr9eUt3J/t7V6FgNEw9QjvBPdysz9KfQDD41fQrC2Y4vKQdiaUpFT4bXlb1RHhLpp8wtm6M5TgSw==`; license Apache-2.0; runtime package files only `bin,go-src`; lifecycle scripts confirmed.
- Independent compiler oracle from E-local dependency: corpus `100/3,141,835`; ES2022 `63/2,389,540`; ES2015 `13/310,653`; ES5 `3/232,957`; final differential `10/10`. The first debug attempt was invalid (`5/10`) because the harness introduced TS18003; after declaring its virtual input, the exact rerun exited 0 at `10/10`. No network, install, package script, target, or package mutation occurred. Repo-local `.tmp/p6a-supervisor-260822` was removed after the run.
- Architecture: offline generation/check-in/embed is the minimum deterministic/package-safe mechanism; full provenance/hash gates, supported config and fail-closed topology, repository/explicit-import precedence, meaning separation, one-final-outcome contract, P6-C1/C2 choices, and later-slice ownership are coherent with current source.
- Build disposition: full build N/A/not run is correct for this decision-only, no-code slice under the explicit package-script prohibition. It does not waive the full build for P6-B or later production slices.
- Deadline: the `03:30 +07` miss is accurately disclosed. It is a process miss but does not invalidate the technical decision; it is not an independent blocker after durable evidence exists.

Failed:

- Sealed-boundary/contract integrity: contract `18+/0-` is outside current authority and contains the stale report hash at line 84.

Not run:

- Full build, Go tests, package scripts, installation, network access, and performance benchmarks: not applicable to this no-code decision review or explicitly forbidden.
- `E:\cheapapp.org` target validation: forbidden in P6-A and remains P6-D work.
- P6-B/C/D runtime behavior: locked and unimplemented.
- `anvien detect-changes`: no implementation commit is being made in this review-only lane.

## Invariant Closure

- Affected invariant: declaration-authority provenance/config/precedence/external identity plus sealed transfer authority and artifact referential integrity.
- Sibling surfaces checked: resolver/analyze, graph transport, Ladybug schema/export, process adjacency, context/impact/rename, graph-health/ResolutionGap, package manifests/lock/install, TypeScript corpus/profile/compiler semantics, all four ledgers, final report, exact diff.
- Residual unverified same-invariant surfaces: none for the architecture decision itself. One verified live blocker remains: the out-of-boundary contract mutation with stale artifact identity.

## Required Fix List For Resubmission

1. Remove `docs/contracts/graph-accuracy-contract.md` from the current sealed P6-A transfer by restoring its HEAD content; do not alter the accepted candidate report or four-ledger content.
2. Supply fresh exact status/diff/hash/encoding evidence showing only the four authorized ledger modifications plus the unchanged final candidate report, with no production/test/fixture/target drift and no protected-handoff mutation.

## Overall Evaluation

The declaration-universe decision is technically well supported: source, graph topology, package reality, TypeScript profile closures, compiler semantics, fail-closed boundaries, consumer adaptations, and later ownership all reproduce. The candidate cannot be accepted in the current repo state because the governing sealed boundary is violated and the additional contract authority is referentially false. Review-only authority forbids repairing that blocker here.
