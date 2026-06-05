---
name: debugging
description: Use when the user asks to debug.
when_to_use: when debugging bugs, test failures, unexpected behavior, regressions, broken flows, production failures, diagnostics, graph/index/query issues, or any fix that must be traced and verified in a repository with Anvien rules
version: 2.0.0
languages: all
---

# Debugging With Anvien

Use this skill to debug from symptom to verified fix in repositories where Anvien rules or Anvien evidence apply. The workflow is continuous: capture the failure, reproduce it, locate the owner, trace the root cause, write a plan, check impact, fix the source, add defenses when needed, and verify with fresh evidence.

## Iron Laws

- MUST understand the root cause before fixing.
- MUST write a plan before fixing.
- MUST run a full build before testing when active repository rules require it.
- MUST verify with fresh evidence before claiming completion.
- MUST follow active repository rules, including Anvien rules when present.

## Phase 1: Capture The Symptom

Record the failing command, log, stack trace, screenshot, test, route, input, output, or runtime behavior.

Do not propose a fix yet.

## Phase 2: Reproduce

Reproduce the failure consistently.

If the failure is not reproducible, gather more evidence instead of guessing:
- run the failing command again
- isolate the input
- capture logs
- add temporary diagnostic instrumentation when appropriate
- compare local, CI, browser, runtime, and environment differences

## Phase 3: Locate Candidate Owner

Use source inspection, logs, tests, and repository tools to identify the likely owner.

When active repository rules require Anvien:
- Run `anvien analyze --force` before graph-based investigation.
- Unknown owner: `anvien query "<symptom>" --repo <repo>`
- Candidate files: `anvien query files "<symptom>" --repo <repo>`
- Suspect file: `anvien file-context <path> --repo <repo> --json`
- Suspect symbol: `anvien context symbol "<symbol>" --repo <repo>`
- Ambiguous symbol: `anvien context symbol "<symbol>" --file <path> --repo <repo>`
- Route, tool, consumer, or contract behavior: `anvien query api`, `anvien api route-map`, `anvien api tool-map`, `anvien api shape-check`, or `anvien api impact`
- Graph, index, resolution, source-site, or retrieval quality issue: `anvien graph-health summary --repo <repo> --json`, `anvien graph-health files --repo <repo> --json`, `anvien graph-health explain "File:<path>" --repo <repo> --json`, `anvien resolution-inventory --graph .anvien/graph.json`, `anvien source-site-accuracy --graph .anvien/graph.json`, or `anvien query-health --repo <repo>`
- Stuck analyze, stale MCP server, lock, or runtime process issue: `anvien doctor locks --repo <repo> --json` or `anvien doctor processes --json`

Anvien evidence is discovery and impact evidence. It does not prove runtime behavior by itself. Confirm the owner with source inspection, reproduction, logs, tests, or runtime traces.

## Phase 4: Trace Root Cause

Trace backward from the failure point until the original trigger is found.

Ask:
- What failed?
- What value, state, input, assumption, dependency, or environment was wrong?
- What code directly caused the failure?
- What called that code?
- Where did the wrong value or state originate?
- Which layer failed to validate or transform it correctly?
- Is this a single bug, repeated pattern, or architectural problem?

For deep stack failures:
- follow stack traces completely
- inspect callers and callees
- trace data flow across boundaries
- add diagnostic instrumentation before dangerous operations
- compare working examples against the broken path

When active repository rules require Anvien, use graph context to support tracing:
- known symbol: `anvien context symbol "<symbol>" --repo <repo>`
- known file: `anvien file-context <path> --repo <repo> --json`
- known execution flow: read `anvien://repo/<repo>/process/{name}`

Runtime stack traces, logs, and tests remain the proof. If graph evidence conflicts with runtime evidence, trust runtime evidence and record the graph-quality gap if relevant.

## Phase 5: Plan Before Fixing

MUST write a plan before fixing.

## Phase 6: Impact Before Edit

Before editing, identify the files, symbols, contracts, routes, or flows likely to change.

When active repository rules require Anvien:
- File change: `anvien impact file <path> --repo <repo> --direction upstream`
- Symbol change: `anvien impact symbol "<symbol>" --repo <repo> --direction upstream`
- Route/API/tool/contract change: `anvien api impact [route] --repo <repo>` or the matching API impact command

Report HIGH or CRITICAL blast radius clearly. It is warning evidence, not an automatic blocker.

Do not edit generated files as the permanent source of truth. Update the generator or source input instead.

## Phase 7: Fix The Root Cause

Fix the source, not the symptom.

Rules:
- make one coherent fix at a time
- avoid unrelated refactors
- do not bundle cleanup with the fix
- preserve existing behavior outside the bug
- add or update focused regression coverage when practical
- if the fix fails, stop and return to root-cause investigation

If multiple fix attempts fail or each fix reveals a different coupling problem, stop and question the architecture before continuing.

## Phase 8: Defense In Depth

When invalid data, unsafe state, or missing validation caused the bug, add defenses at the layers where they naturally belong.

Consider:
- entry point validation
- business logic validation
- environment guards
- debug instrumentation
- tests that prove bypass paths are blocked

When active repository rules require Anvien, use `file-context`, `context symbol`, graph-health, or resolution inventory to find missed paths, callers, callees, isolation, or unresolved-reference signals. Use that evidence to choose validation checkpoints; do not replace tests with graph evidence.

## Phase 9: Verify

Run the command that proves the claim.

When active repository rules require a full build before testing, run the full build first, then run the test suite or targeted tests that prove the fix.

Verification may include:
- original failing command
- focused regression test
- relevant unit/integration/e2e tests
- full build
- browser/runtime validation
- logs or diagnostics proving the failure path is fixed

When active repository rules require Anvien:
- Before committing implementation work: `anvien detect-changes --repo <repo> --scope all`
- Route, tool, consumer, or contract changes may require `anvien api shape-check` or `anvien api impact`
- Graph builder, resolver, analyzer, query, source-site, or resolution behavior changes may require `anvien graph-health`, `anvien resolution-inventory`, `anvien source-site-accuracy`, or `anvien query-health`

Choose the command that proves the claim. Do not run unrelated graph commands mechanically.

Do not claim completion without fresh verification evidence.

## Red Flags

Stop and return to investigation if:
- proposing fixes before reproducing
- proposing fixes before root cause is known
- editing before writing a plan
- skipping impact checks required by repository rules
- using graph evidence as runtime proof
- adding multiple fixes at once
- relying on "should work"
- claiming success without fresh verification
- continuing after repeated failed fixes without questioning architecture

## Evidence To Report

When reporting the result, include:
- symptom and reproduction evidence
- confirmed root cause
- owner evidence
- impact evidence when required
- fix summary
- verification commands and outcomes
- detect-changes result when required
- remaining risks or test gaps
