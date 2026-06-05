# Benchmark Ledger

Title: AI Context Master Rules Generation
Date: 2026-06-06
Status: Completed
Plan: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-plan.md
Evidence: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-evidence.md
Benchmark: docs/plans/2026-06-06-aicontext-master-rules-generation/2026-06-06-aicontext-master-rules-generation-benchmark.md

## B0 - Initial Graph Inventory

| Metric | Value | Source |
| --- | ---: | --- |
| Indexed files before implementation plan | 1404 | `anvien analyze . --force` |
| Graph nodes before implementation plan | 84448 | `anvien analyze . --force` |
| Graph relationships before implementation plan | 122937 | `anvien analyze . --force` |
| File projection dependency edges before implementation plan | 16570 | `anvien analyze . --force` |
| Unresolved items before implementation plan | 430 | `anvien analyze . --force` |

## B1 - Generated Rule Inventory

| Metric | Value | Source |
| --- | ---: | --- |
| Generated numbered AGENTS rules per generated file | 11 | `requireGeneratedMasterRules`, rules 0 through 10 |
| Generated target files carrying rule block | 2 | `AGENTS.md` and `CLAUDE.md` generated through `renderAnvienBlock` |
| Full-build analyze indexed files | 1407 | `anvien analyze . --force` during full build |
| Full-build analyze graph nodes | 84478 | `anvien analyze . --force` during full build |
| Full-build analyze graph relationships | 122964 | `anvien analyze . --force` during full build |
| Full-build analyze file projection dependency edges | 16570 | `anvien analyze . --force` during full build |
| Full-build analyze unresolved items | 430 | `anvien analyze . --force` during full build |
| Final indexed files | 1408 | final `anvien analyze . --force` before commit |
| Final graph nodes | 84487 | final `anvien analyze . --force` before commit |
| Final graph relationships | 122973 | final `anvien analyze . --force` before commit |
| Final file projection dependency edges | 16570 | final `anvien analyze . --force` before commit |
| Final unresolved items | 430 | final `anvien analyze . --force` before commit |
