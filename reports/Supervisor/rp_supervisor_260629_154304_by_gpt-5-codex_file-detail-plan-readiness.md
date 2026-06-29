# Supervisor Report: File Detail Plan Readiness

Verdict: REJECT

## Metadata
- Report file: reports/Supervisor/rp_supervisor_260629_154304_by_gpt-5-codex_file-detail-plan-readiness.md
- Review time: 260629 154304 Asia/Bangkok
- Reviewer: gpt-5-codex
- Repo/project: Anvien
- Scope reviewed: docs/plans/2026-06-29-file-detail-compact-full-detail plan set
- Claim reviewed: the file-detail compact full-detail plan is ready to implement, or any missing items are known
- Authority used: latest user request, AGENTS.md repo rules, planner/supervisor skills, plan artifacts, source code, README/RUNBOOK command docs
- Related artifacts: docs/plans/2026-06-29-file-detail-compact-full-detail/*

## Executive Summary
- Problem: The plan is meant to make `file-detail` shorter without cutting its full-detail function, and to add richer related-file metadata.
- Decision: REJECT for implementation readiness. The plan is structurally strong, but it does not close all same-invariant Anvien surfaces and it leaves the central "full detail versus limits" behavior under-specified.
- Required outcome: Add the missing MCP/agent surface slice or explicit preservation proof, and make the compact full-detail limit/default semantics unambiguous before implementation starts.

## Blocking Findings

### [HIGH] MCP and agent file-context surfaces are in the same invariant but are not planned
File: docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-plan.md:65
Issue: The plan scope covers compact builder, CLI, HTTP/API, Web contracts, and Web rendering, but it does not include an explicit MCP/agent surface slice even though repo rules treat MCP, CLI, and Web/API as Anvien command surfaces of the same tool.
Evidence: AGENTS.md:8-13 says Anvien has multiple command surfaces and MCP tools are Anvien commands. The plan scope at lines 65-70 lists CLI, HTTP/API, Web contract/types, and Web panel work, but no MCP/context surface. P2 lines 320-329 scopes CLI, HTTP, contracts, and generated Web artifacts only. Actual status lines 86-95 classify builder, CLI, HTTP, contract, generated, Web client, Web panel, and runtime/browser validation only. The dependency table at lines 111-118 has CLI, HTTP, contracts, generated files, Web, and unrelated worktree, but no `internal/mcp/*` row. Source contradicts that omission: `internal/mcp/target_dispatch.go:63-78` calls `filecontext.NewBuilder(g).BuildFileContext`, `internal/mcp/target_dispatch.go:97-115` converts the expanded file context into MCP file-layer payloads, `internal/mcp/context.go:188-204` returns `fileContext` directly for context file payloads, and `internal/mcp/impact.go:270` plus `internal/mcp/impact.go:345` use the same MCP file context path for impact flows.
Why this blocks acceptance: P1 edits the shared builder/model area. Without a planned MCP decision, implementation can pass CLI/HTTP/Web acceptance while leaving `context file`, MCP impact file flows, or agent-facing file detail payloads stale, incompatible, or untested. That violates the same-invariant closure standard and AGENTS.md rule 0.
Fix direction: Add a dedicated MCP/agent surface item, for example P2-D, or explicitly fold MCP into P2 with its own actual-status rows, source scope, impact gate, behavior tests/smoke, and acceptance. The plan must decide whether MCP stays expanded-only, gets compact/expanded selection, or adapts to compact internals while preserving its public payload.
Re-review evidence required: Updated plan and actual-status rows naming the MCP files/surfaces, source evidence for their chosen behavior, impact evidence requirements, and tests/smoke commands for `context file` and file impact MCP/CLI-equivalent flows.

### [HIGH] Compact "full detail" limit semantics are not explicit enough
File: docs/plans/2026-06-29-file-detail-compact-full-detail/2026-06-29-file-detail-compact-full-detail-plan.md:53
Issue: The plan correctly says compact output must not be summary-only and must preserve all detail facts for the requested file-detail scope, but it does not lock down the behavior of defaults, sample limits, and row limits strongly enough to prevent an implementation that only compacts sampled data.
Evidence: The accepted direction is stated at plan lines 53-61 and non-goals at lines 74-75 reject summary/top-N/cut behavior. Requirements at lines 83-90 require relationship, unresolved, symbol, and related-file fact preservation. However acceptance lines 96-101 allow an accepted default or documented compact mode and cover "sample/row limit behavior" without deciding whether compact `--json` is unlimited full-detail by default, bounded only when the caller requests limits, or intentionally preserves current sample-limited semantics. The benchmark ledger uses baselines with `--relationships 1 --unresolved 1 --linked 1` at benchmark lines 47-52, and actual-status line 89 states current CLI sample flags limit samples per group.
Why this blocks acceptance: The user explicitly rejected "cat bo" style summarization. If implementation only compacts existing sampled groups, the output becomes shorter but may still omit relationship/unresolved/linked facts. The plan needs a precise contract so implementers and tests can distinguish compact full-detail from compact sampled-detail.
Fix direction: Add an owner decision and acceptance invariant for limits. For example: compact default machine output includes all rows for the requested file unless the caller explicitly supplies limit flags; when limits are supplied, the payload must expose totals plus omitted counts/ranges so the limitation is visible. Alternatively, if preserving current sample-limited default is intended, record it as an explicit breaking/behavior decision and explain why it still satisfies "file detail" rather than summary.
Re-review evidence required: Updated Requirements/Acceptance/Benchmark rows defining default format, limit flags, full-row parity tests, limited-row tests, and evidence IDs that prove no hidden truncation occurs.

### [MEDIUM] User-facing command docs are not in the implementation scope
File: README.md:300
Issue: The plan requires CLI/API help and contract docs to make compact versus expanded behavior explicit, but it does not scope README/RUNBOOK updates or explicitly decide they are out of scope.
Evidence: Plan line 91 requires compact versus expanded behavior to be explicit in help and contract docs. README.md:286 and README.md:300 document `anvien context file` and `anvien file-detail <path>`, and README.md:479 documents `/api/file-detail`. RUNBOOK.md:221-223, RUNBOOK.md:313-320, and RUNBOOK.md:507-514 contain file-detail command/API examples. P2-A through P2-C cover CLI help, HTTP behavior, contracts, and generated Web artifacts, but no README/RUNBOOK source docs.
Why this blocks acceptance: If defaults or required format params change, stale docs preserve the old behavior and can mislead users/operators. This is lower risk than the MCP and limit-contract gaps, but it should be either planned or explicitly excluded with a reason.
Fix direction: Add a docs check/update step after CLI/API contract decisions, or record that README/RUNBOOK examples remain valid because defaults preserve old invocations.
Re-review evidence required: Updated plan row for docs or a written non-goal/decision explaining why docs do not need changes.

## Source-Level Clearance Notes
- Plan file set: blocked - the four plan artifacts exist and placeholders were not found, but same-invariant MCP coverage is missing from the phase/status structure.
- `internal/mcp/*`: blocked - source shows `BuildFileContext` is used by MCP/context and impact file flows, but the plan does not include MCP as a planned surface.
- CLI/HTTP/Web/contract slices: clear with caveat - plan has scoped slices, gates, tests, generated artifact rules, runtime validation, and commit boundaries for these surfaces.
- README/RUNBOOK command docs: blocked pending decision - existing docs mention file-detail invocations that may need update if format/default behavior changes.

## Evidence Checked
Passed:
- `rg -n "\\{\\{|\\}\\}|SLICE_|PHASE_|TARGET_SCOPE|P0_EVIDENCE|METRIC|EVIDENCE_ID_OR_COMMAND" docs\\plans\\2026-06-29-file-detail-compact-full-detail` returned no matches, so no template placeholders were found.
- `rg --files docs\\plans\\2026-06-29-file-detail-compact-full-detail` returned the expected plan, evidence, benchmark, and actual-status files.
- Plan source lines 45-55 include slice-splitting, hidden-fallback, full-detail, related-file, and compatibility invariants.
- Plan source lines 320-517 include CLI, HTTP, contract, generated artifact, tests, evidence, and commit gates.
- Plan source lines 754-760 correctly identify high/critical blast radius, lossy compact risk, and Web presentation limit risk.

Failed:
- Source scan found same-invariant MCP consumers in `internal/mcp/target_dispatch.go`, `internal/mcp/context.go`, and `internal/mcp/impact.go`, while plan/status scope omits MCP.
- Plan/benchmark/status wording does not decisively define compact full-detail default versus explicit limited output.
- README/RUNBOOK contain file-detail docs/examples but the plan does not include docs update or non-update decision.

Not run:
- Full build/tests: not applicable for a plan-readiness review before implementation.
- Anvien graph refresh/impact: not rerun in this review because source evidence was enough to reject readiness; the plan already records baseline graph/impact evidence, but resubmission should refresh graph evidence if plan edits add MCP impact scope.

## Invariant Closure
- affected invariant: file-detail/file-context detail contract across Anvien command surfaces and consumers.
- sibling surfaces checked: CLI plan, HTTP/API plan, Web contract/types plan, Web panel plan, MCP/context source, impact source, README/RUNBOOK command docs, actual-status ledger.
- residual unverified same-invariant surfaces: MCP/agent payload behavior and docs/default limit behavior remain unclosed.

## Required Fix List For Resubmission
1. Add MCP/agent file-context coverage to the plan and actual-status ledger, with impact gates and validation evidence for `context file` and file-impact related flows.
2. Add a precise compact full-detail limit/default decision, including tests and benchmark rows for unlimited/default and explicitly limited modes.
3. Add README/RUNBOOK docs update or an explicit no-change decision tied to the final CLI/API default behavior.

## Overall Evaluation
The plan is close and has the right main direction: compact representation, related-file metadata, compatibility, tests, build/runtime validation, evidence, benchmarks, and commit boundaries. It is not ready to implement because it misses a real same-invariant MCP surface and leaves the central "gọn nhưng không cắt" contract ambiguous around limits. Fixing those items should be a small plan supplement, not a rewrite.
