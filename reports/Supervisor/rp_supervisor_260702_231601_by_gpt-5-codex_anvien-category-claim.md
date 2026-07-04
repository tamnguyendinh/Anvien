# Supervisor Report: Anvien Category Claim

Verdict: REJECT

## Metadata
- Report file: reports/Supervisor/rp_supervisor_260702_231601_by_gpt-5-codex_anvien-category-claim.md
- Review time: 260702 231601 Asia/Bangkok
- Reviewer: gpt-5-codex
- Repo/project: Anvien
- Scope reviewed: External model claim that Anvien is a "code intelligence tool"
- Claim reviewed: "Code intelligence tool" is an adequate product-level description of Anvien
- Authority used: current repo source, README, ARCHITECTURE, AGENTS-managed rules, fresh Anvien graph evidence
- Related artifacts: none

## Executive Summary
- Problem: The external model's category label is directionally true but too narrow for Anvien as implemented in this repo.
- Decision: Reject the claim as a complete assessment. Anvien includes a code intelligence graph core, but the repo also implements agent context generation, skill packaging/sync, MCP/CLI/HTTP command surfaces, local Web/runtime surfaces, prompts/resources, and workflow guardrails.
- Required outcome: Describe Anvien as a local-first AI coding system with code intelligence graphs plus agent skills/workflow/runtime surfaces.

## Blocking Findings

### [HIGH] Claim omits Anvien's agent workflow and skill system
File: E:/Anvien/README.md:3
Issue: The top-level README defines Anvien as a "2-in-1 tool for AI coding: code intelligence graphs + Powerful agent skills", not just a code intelligence tool.
Evidence: README lines 3, 14, and 18 state the product has both graph intelligence and AI agent skills, and exposes the graph to AI coding agents, CLI commands, and a local Web UI.
Why this blocks acceptance: A product-level label that only says "code intelligence tool" drops a first-class half of the stated product.
Fix direction: Use "local-first AI coding system: code intelligence graph + agent skills/workflow/runtime" or equivalent.
Re-review evidence required: Product description acknowledges both graph intelligence and agent-facing workflow surfaces.

### [HIGH] Claim omits source-owned AGENTS/CLAUDE and skill rendering
File: E:/Anvien/internal/aicontext/aicontext.go:61
Issue: Anvien has a real generator for `AGENTS.md`, `CLAUDE.md`, command selection, skill selection, master rules, and skill installation.
Evidence: `GenerateAIContextFiles` writes managed agent context files; `renderAnvienBlock` emits command/resource/prompt/skill guidance; `installBaseSkills` installs skill packages. `skill_packages.go` embeds and discovers skill packages, resolves runtime skill source, and syncs installs.
Why this blocks acceptance: This is agent orchestration/instruction infrastructure, not merely code intelligence.
Fix direction: Mention "agent context generator" and "skill package manager/sync" as part of Anvien.
Re-review evidence required: The description covers `internal/aicontext` and skill package behavior.

### [MEDIUM] Claim omits multiple command/runtime surfaces
File: E:/Anvien/ARCHITECTURE.md:3
Issue: Architecture describes backend, CLI, MCP server, analyzer, storage, contracts, session bridge, and Web UI as part of the system.
Evidence: ARCHITECTURE lines 21-29 list HTTP API, MCP, CLI, contracts, aicontext, session bridge, Web UI, and launcher. README lines 43-47 list CLI, MCP stdio, Local HTTP API, Web UI, and Windows launcher.
Why this blocks acceptance: Anvien is implemented as a local runtime with several surfaces, not just a code analysis library.
Fix direction: Mention CLI/MCP/HTTP/Web/launcher surfaces when defining Anvien.
Re-review evidence required: Product description reflects the implemented surfaces.

### [MEDIUM] Claim omits workflow prompts/resources and evidence discipline
File: E:/Anvien/internal/mcp/prompts.go:32
Issue: Anvien exposes MCP resources and prompts that guide agent workflows, including impact review and evidence-backed architecture mapping.
Evidence: README lines 367-410 list MCP tools/resources/prompts. `resources.go` defines repo, setup, context, clusters, processes, schema resources. `prompts.go` defines `detect_impact` and `generate_map`.
Why this blocks acceptance: Workflow templates and resources are part of Anvien's agent-facing operating model.
Fix direction: Include "workflow layer" or "agent-facing evidence system".
Re-review evidence required: The description covers MCP resources/prompts.

## Source-Level Clearance Notes
- `README.md`: blocked for narrow claim; lines 3, 14, 18, 43-47, 315-336, and 367-410 show product scope beyond code intelligence.
- `ARCHITECTURE.md`: blocked for narrow claim; lines 3-29 and 35-63 show local-first system architecture and multiple serving surfaces.
- `internal/aicontext/aicontext.go`: blocked for narrow claim; lines 61, 94, 122-173, 197, 205, 218, and 316 show generated agent context and skill guidance.
- `internal/aicontext/skill_packages.go`: blocked for narrow claim; lines 27, 122, 126, 188, 344, 815, 990, 1120, and 1128 show embedded/runtime skill package discovery, installation, path safety, and guide rendering.
- `internal/cli/command.go`: blocked for narrow claim; lines 38, 66, 91-96, 103, 128, 155, 233, 240, and 247 show the CLI/root/analyze/runtime command surface and AI-context generation after analyze.
- `internal/mcp/server.go`, `internal/mcp/resources.go`, `internal/mcp/prompts.go`, `internal/mcp/tools.go`: blocked for narrow claim; tool dispatch, resources, prompts, and registered tool definitions prove agent-facing command surfaces.
- `internal/httpapi/server.go`: blocked for narrow claim; lines 55-78 show local HTTP/Web/session/MCP endpoints.

## Evidence Checked
Passed:
- `anvien analyze --force` refreshed the graph at commit `e6397f2`; result: 1455 files scanned, 83901 graph nodes, 122737 relationships, file projection built.
- `anvien query "agent context AGENTS CLAUDE skills generator" --repo E:\Anvien` found `internal/aicontext/aicontext.go` and `internal/aicontext/skill_packages.go` as owner-discovery and docs/setup/AI-context surfaces.
- `anvien query "MCP tools resources prompts command surface" --repo E:\Anvien` found MCP route/tool map, API impact, shape check, prompts, and CLI command surfaces.
- `anvien file-detail` on `internal/aicontext/aicontext.go`, `internal/aicontext/skill_packages.go`, `internal/cli/command.go`, and `internal/mcp/server.go` showed fresh, non-stale graph data and linked flows/tests.
- Source and docs were inspected directly via README, ARCHITECTURE, `internal/aicontext`, `internal/cli`, `internal/mcp`, and `internal/httpapi`.

Failed:
- The external model label fails as a complete product assessment because it ignores implemented first-class agent workflow/runtime surfaces.

Not run:
- Full build/test suite was not needed because this is a classification review, not an implementation acceptance review.

## Invariant Closure
- affected invariant: product/category description must match implemented repo authority and source reality.
- sibling surfaces checked: README, ARCHITECTURE, CLI, MCP, HTTP API, aicontext generator, skill packages, graph evidence.
- residual unverified same-invariant surfaces: none required for this classification claim.

## Required Fix List For Resubmission
1. Replace "Anvien is a code intelligence tool" as a complete assessment with a broader description.
2. Include code intelligence graph, agent skills, generated agent context, command/runtime surfaces, and evidence workflow.

## Overall Evaluation
The label "code intelligence tool" is literally true for Anvien's graph core, but incomplete enough to mislead. The repo implements a local-first AI coding system where code intelligence is the substrate, and agent workflow, skill packaging, command routing, MCP resources/prompts, HTTP/Web runtime, and evidence discipline are first-class product surfaces.
