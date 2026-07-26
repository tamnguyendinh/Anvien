# Historical `p0-rc-c-fixture` Boundary Note

Date: 2026-07-26 15:40 +07
Status: historical artifact; removed by owner before this note was written

## What it was

`E:\Anvien\p0-rc-c-fixture` was a small synthetic TypeScript repository used by the earlier P0-RC-C ambient/builtin-resolution investigation. Its source was not `cheapapp.org`. The fixture contained:

- `src/main.ts` using `Promise`, `Math.max`, `Math.min`, and a declared ambient symbol;
- `src/ambient.d.ts` with repository-local ambient declarations;
- `tsconfig.json` selecting ES2022 libraries;
- `oracle.cjs` using the TypeScript compiler API;
- `fixture-benchmark*.json` outputs; and
- a fixture-local Anvien index plus generated `.agents`, `.claude`, `AGENTS.md`, and `CLAUDE.md`.

The purpose was to isolate whether Anvien's resolver/index differed from the TypeScript compiler's ambient declaration universe. That purpose is reflected independently in `reports/Supervisor/rp_supervisor_260726_060300_root_cause_ambient_resolution.md`, which labels the slice `P0-RC-C` and describes the same five probe facts.

## Why it was wrong here

The directory was created at the Anvien repository root at `09:49:18` and last written around `10:08:03`. A temporary fixture is required by the repository rules to live under `E:\Anvien\.tmp\` (normally a slice-specific directory such as `.tmp\p0-rc-c\fixture`). The earlier agent instead used the root path, so the fixture appeared as an untracked root repository item. Anvien analysis then generated its local index and guidance files inside that misplaced fixture.

This was a process/artifact-boundary violation by the earlier context, not a requirement of the current `E:\cheapapp.org` investigation. It was not a copy of the target repository and it was not used as current target evidence.

## Current state

The owner removed `E:\Anvien\p0-rc-c-fixture` at approximately 15:39. A read-only check returned `FIXTURE_REMOVED`; no delete, move, or rewrite command was issued by this investigation. Current work continues with probes only under `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\` and reports under `E:\Anvien\reports\`.

## Evidence limits

The directory's original filesystem timestamps, file inventory, Git untracked status, fixture source excerpt, benchmark labels, and fixture graph metadata were observed before removal. The exact creator process/agent is not recoverable from the remaining repository files; assigning it to the current subagents would be unsupported. The historical report is a pointer only and is not accepted as proof for the fresh target graph.
