# Plan

Title: AI Context Master Rules Generation
Date: 2026-06-06
Status: Completed
Plan: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-plan.md
Evidence: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-evidence.md
Benchmark: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-benchmark.md

## Goal

Move the currently hand-written repository master rules into the `aicontext.go` generator so generated `AGENTS.md` and `CLAUDE.md` contain the same rule block automatically.

## Scope

In scope:

- `internal/aicontext/aicontext.go`
- `internal/aicontext/aicontext_test.go`
- Generated managed section content for `AGENTS.md` and `CLAUDE.md`

Out of scope:

- Changing skill package install/sync behavior
- Changing command selection guide content beyond preserving it after the new rule block
- Editing generated `AGENTS.md` or `CLAUDE.md` as source of truth
- Changing skill metadata or skill content

## Requirements

- The generated managed block must include the full user-provided `# Master iron rules` section.
- The generated managed block must include the full user-provided `# AGENTS Rules` section, including rules 0 through 10.
- Both `AGENTS.md` and `CLAUDE.md` must receive the same rule block through generation.
- Existing generated Anvien command/resource/MCP/skill guide content must remain present.
- Tests must assert generation behavior after code behavior exists.

## Validation Sequence

Full build means run the whole command sequence below from the repository root before testing:

```powershell
cd .\anvien
npm install
npm run build
npm install -g .
Get-Command anvien
anvien version
cd ..
powershell -ExecutionPolicy Bypass -File .\anvien-launcher\build.ps1
anvien version
anvien analyze . --force
```

After the full build succeeds, run targeted Go tests for `internal/aicontext`, then run `anvien detect-changes --repo Anvien --scope all` before commit.

## Phase Checklist

- [x] P0-A: Confirm generator scope and blast radius with Anvien.
  Goal: Identify exact generator/test symbols and risk before code edit.
  Work Steps: Refresh graph, inspect `internal/aicontext/aicontext.go`, inspect linked tests, run impact for the generator symbol that will be edited.
  Acceptance: Evidence records implementation file, test file, symbol, linked tests, and blast radius.

- [x] P1-A: Add the master rules block to generator output.
  Goal: Make generated `AGENTS.md` and `CLAUDE.md` carry the manual rule block.
  Work Steps: Add a single generator helper/string for the rule block and call it from `renderAnvienBlock` before existing Anvien code-intelligence content.
  Acceptance: Generated content contains `# Master iron rules`, `# AGENTS Rules`, and rules 0 through 10.

- [x] P2-A: Update tests after generator behavior is implemented.
  Goal: Lock the generator invariant.
  Work Steps: Update `TestGenerateAIContextFilesCreatesManagedContextAndSkillPackages` or add focused coverage to assert both generated files contain the rule block and existing guide content still exists.
  Acceptance: Test fails on old generator output and passes with the new generated rule block.

- [x] P3-A: Validate, record evidence, and commit.
  Goal: Close with build/test/change-detection evidence.
  Work Steps: Run gofmt, full build, targeted Go tests, detect-changes, update evidence/benchmark, create report/notes if required, and commit.
  Acceptance: Evidence records pass/fail outputs and final changed-file list.

## Risks

- `renderAnvienBlock` is shared by both `AGENTS.md` and `CLAUDE.md`; this is desired but must be verified for both outputs.
- The generated section already contains Anvien rules, so duplicate/conflicting wording must be avoided by placing the master rule block as the generated authority preface and leaving existing Anvien command guidance intact.
