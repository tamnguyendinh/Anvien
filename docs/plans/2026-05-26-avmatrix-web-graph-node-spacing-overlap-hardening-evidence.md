# AVmatrix Web Graph Node Spacing And Overlap Hardening Evidence Ledger

Date: 2026-05-26

Status: Active

Companion files:

- Plan: [2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-plan.md](2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-plan.md)
- Benchmark ledger: [2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-benchmark.md](2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-benchmark.md)

## Evidence Rules

Record evidence as each evidenced task is completed. Evidence should include source traces, impacted files, impact output summaries, test commands, build commands, browser screenshots, DOM diagnostic output, and concise observations needed to audit the plan later.

For doc-only commits, do not use AVmatrix.

Do not record inferred runtime behavior as final evidence. Every behavior claim must include source inspection, command output, test output, browser evidence, or exact geometry measurements.

Keep this file separate from the benchmark ledger. This file records what was inspected, what was run, what changed, and what artifacts prove behavior. Quantitative geometry or performance measurements belong in the benchmark ledger.

## E0 - Plan Creation Evidence

Date: 2026-05-26

Status: recorded

Created file set:

- `docs/plans/2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-plan.md`
- `docs/plans/2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-evidence.md`
- `docs/plans/2026-05-26-avmatrix-web-graph-node-spacing-overlap-hardening-benchmark.md`

Plan creation scope:

- Treat dense Web graph node overlap/crowding as a product readability bug.
- Define the default requirement as one rendered node diameter of empty edge-to-edge gap between rendered circular nodes.
- Plan a hard layout invariant instead of only increasing a spacing constant.
- Keep evidence and benchmark records separate.
- Preserve existing graph orientation labels, filter behavior, island/ring spacing, and deterministic layout semantics.

Doc-only note:

- This plan creation is documentation-only, so AVmatrix was not used for this commit slice.

## E1 - Initial Source Trace From Prior Investigation

Date: 2026-05-26

Status: preliminary; implementation must re-verify with fresh AVmatrix graph before code edits

Relevant source owners observed:

| Area | Path | Observed responsibility |
|---|---|---|
| Web graph conversion and layout | `avmatrix-web/src/lib/graph-adapter.ts` | Computes rendered node size caps, cluster node spacing, island radius, deterministic node offsets, island placement, and ring placement. |
| Sigma rendering integration | `avmatrix-web/src/hooks/useSigma.ts` | Applies rendered node size caps and camera/rendering behavior. |
| Graph canvas | `avmatrix-web/src/components/GraphCanvas.tsx` | Hosts graph UI and may expose diagnostics or validation hooks if needed. |
| Geometry tests | `avmatrix-web/test/unit/graph-adapter.edge-geometry.test.ts` | Contains existing island/ring geometry tests, but must be extended with pairwise same-island spacing checks. |
| Label tests | `avmatrix-web/test/unit/graph-orientation-labels.test.ts` | Protects graph orientation label metadata and overlap guardrails. |
| Browser/e2e tests | `avmatrix-web/e2e/graph-orientation-labels.spec.ts` | Validates label visibility and overlap behavior in browser. |

Important source search command used during plan creation:

```powershell
rg -n "getClusterNodeSpacing|getClusterIslandRadius|getIslandOffset|MAX_RENDERED_NODE_SIZE|capRenderedNodeSize|golden|GOLDEN_ANGLE|GraphCanvas|useSigma|edge-geometry" avmatrix-web -g "*.ts" -g "*.tsx"
```

Observed source symbols:

- `MAX_RENDERED_NODE_SIZE`
- `capRenderedNodeSize`
- `getClusterNodeSpacing`
- `getClusterIslandRadius`
- `getIslandOffset`
- `GOLDEN_ANGLE`
- `GraphCanvas`
- `useSigma`

## E2 - Initial Problem Finding From Prior Investigation

Date: 2026-05-26

Status: preliminary; implementation must reproduce in Phase 0 before code edits

Observed current behavior:

- Current layout uses deterministic spiral placement for nodes inside an island.
- The broad layout has island/ring spacing tests, but no hard same-island pairwise spacing invariant was observed.
- A dense island can pass island/ring separation checks while still placing two rendered nodes too close together.

Preliminary conclusion:

- The user's report is credible and aligns with the current layout structure.
- The proposed UX rule is correct as a product default, but the implementation should express it as a minimum center-distance invariant derived from rendered node size semantics.
- Merely increasing a global spacing constant is not enough proof, because perturbation, future tuning, and camera fit behavior can still reintroduce visual crowding.

## E3 - Plan Creation Source Search

Date: 2026-05-26

Status: recorded

Command:

```powershell
rg -n "node spacing|island radius|edge gap|overlap|cluster island|graph label|ring label" docs\plans avmatrix-web -g "*.md" -g "*.ts" -g "*.tsx"
```

Observed related historical plan context:

- The completed skill-system plan already included adaptive island/ring spacing and graph orientation labeling work.
- That plan's acceptance language included readable island spacing, but the follow-up issue is narrower: enforce pairwise node-node clearance inside dense islands.
- The new plan is a follow-up bug hardening plan and must not rewrite prior closed plan evidence.

## E4 - Pending Implementation Evidence

Date: 2026-05-26

Status: pending

Record implementation evidence here as phases complete:

- fresh AVmatrix graph counts before graph-based implementation work;
- impact analysis blast radius for edited graph layout/Sigma/canvas symbols;
- source diffs and touched files;
- geometry test commands and results;
- Web unit test commands and results;
- e2e/browser test commands and results;
- screenshot artifact paths;
- `detect-changes` output before implementation commits;
- commit hashes for completed implementation slices.
