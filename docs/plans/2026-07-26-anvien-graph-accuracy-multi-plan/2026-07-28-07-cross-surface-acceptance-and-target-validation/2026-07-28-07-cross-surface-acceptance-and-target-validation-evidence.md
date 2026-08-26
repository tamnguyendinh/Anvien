# Anvien Cross-Surface Acceptance and Target Validation Evidence Ledger

## Metadata

- Date: `2026-07-28`
- Last revised: `2026-08-24`
- Plan: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-plan.md`
- Evidence: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-evidence.md`
- Benchmark: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-benchmark.md`
- Actual status: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-07-cross-surface-acceptance-and-target-validation/2026-07-28-07-cross-surface-acceptance-and-target-validation-actual-status.md`
- Contract: `docs/contracts/graph-accuracy-contract.md`
- Roadmap: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`
- Predecessor: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy/2026-08-24-06a-accelerate-analyze-without-sacrificing-accuracy-plan.md`

## Evidence Rules

- Evidence must identify the exact source, command, artifact, result, and commit it proves.
- The problem-origin report proves measured defects and bounded targets. Its proposed implementation design is DRAFT and cannot be used as implementation proof.
- A predecessor handoff is valid only when it identifies its child, accepted commit, completed evidence IDs, and refreshed successor state. P7-A requires seven handoffs through Child 06A and opens from the Child 06A closure commit, never directly from Child 06.
- A pending evidence ID reserves required proof; it is not a completed result.
- Child 07 records validation evidence only. Correctness repair evidence belongs to the owning Child 01 through Child 06 ledger; analyze-performance or accepted-equivalence repair evidence belongs to Child 06A.
- Build and behavior results belong here. Counts, sizes, durations, latency, throughput, and memory belong in the benchmark ledger.
- Every target numerator must link to exact source-site and graph/outcome records. Empty, unrelated, stale, or pass-by-default results are invalid.
- The affected-surface denominator is derived from accepted diffs, impact evidence, changed fields, and concrete readers. Broad category names are not evidence.
- Final acceptance requires Supervisor review, detect-changes, an isolated validation commit, and a known target/worktree state.

### Evidence ID Naming

Use `E<phase>-<item>-<kind><n>`. The checklist and tables below declare every accepted or reserved ID explicitly for P7-A, P7-B, and P7-C.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-RULE1`: `E:\Anvien\AGENTS.md` was read in full before document work. It requires planner use, fresh graph evidence for graph-based work, evidence-backed decisions, full build before validation, Supervisor review, and per-slice commits.
- `E0-P0A-PLANNER1`: `.agents/skills/planner/SKILL.md` and all four planner templates were read in full.
- `E0-P0A-REPORT1`: `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` was read in full. It explicitly marks its proposed architecture as DRAFT because no approved SPEC authority existed.
- `E0-P0A-REPORT2`: the same report records the target acceptance denominators: same-name identity `2/4 -> 4/4`, binding patterns `0/6 -> 6/6`, direct exports `0/21 -> 21/21`, barrel calls `0/2 -> 2/2`, and ambient sites `0/3 -> 3/3` truthful external/capability outcomes.
- `E0-P0A-REPORT3`: the report distinguishes module export from visibility, module-path resolution from export resolution, and in-repository gaps from external-declaration outcomes. These are problem/acceptance facts; exact implementation remains owned by the corrected contract and owning child evidence.
- `E0-P0A-BOUNDARY1`: the report names `E:\cheapapp.org` as the real bounded target. Child 07 analyzes it in place and keeps Anvien-owned reports and fixtures outside the target.
- `E0-P0A-GRAPH1`: `anvien analyze --force` completed at HEAD `238aec06d28286acc0fca0f4e6a69f9eb4ff6a49` with `1,556` scanned files, `676` parsed code files, `0` failed files, `85,101` nodes, and `123,969` relationships. Graph path: `E:\Anvien\.anvien\graph.json`.
- `E0-P0A-FD1`: fresh `anvien file-detail` calls classified all four Child 07 ledger files as low-risk documentation. Each ledger had one inbound relationship from the then-current roadmap and no source symbols, flows, routes, tools, or tests.
- `E0-P0A-DOC1`: full contextual review confirms that Child 07 owns validation only. Its active contract is limited to three terminal slices and derives consumer coverage from accepted changes and current impact evidence.
- `E0-P0A-STATUS1`: actual status classifies P7-A, P7-B, and P7-C separately; P7 remains dependency-blocked until all seven predecessor handoffs through Child 06A exist.
- `E0-P0A-DIFF1`: this correction changes documentation only. No Child 07 production implementation or target source change is claimed.
- `E0-P0A-ORDER1`: Direct structure is `Child 06 -> Child 06A -> Child 07`. Child 06 closes at P6-D commit `81163e39718b94a509e41114cada224e8f269e36`; Child 06A currently has no detailed timing map, ordered bottleneck list, production attempt, accepted speedup, or closure commit. Therefore P7-A remains blocked.

## E7 - P7 Evidence

Matching plan item(s): `P7-A`, `P7-B`, `P7-C`

### P7-A — Determinism, integrity, and repeated normal analyze

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E7-P7A-INPUT1` | seven accepted handoffs through Child 06A plus the exact source commit, config, analyzer build, command, machine policy, canonical comparison rules, and accepted Child 06A performance/equivalence/resource basis | pending; Child 06A closure commit absent |
| `E7-P7A-BUILD1` | full build completed before validation | pending |
| `E7-P7A-ANALYZE1` | at least five successful normal analyzes with identical inputs; failed invocations are reported as failures and excluded from accepted graph results | pending |
| `E7-P7A-DETERMINISM1` | canonical node and relationship fact sets are equal across accepted runs under the Child 01/02 contract | pending |
| `E7-P7A-INTEGRITY1` | no unexplained lost occurrences, missing endpoints, orphan references, or persisted record drops in affected projections | pending |
| `E7-P7A-REVIEW1` | independent Supervisor PASS | pending |
| `E7-P7A-DETECT1` | Anvien detect-changes before validation-artifact commit | pending |
| `E7-P7A-COMMIT1` | isolated P7-A evidence commit and known worktree | pending |

### P7-B — Five-family terminal target acceptance

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E7-P7B-TARGET1` | target commit/worktree pre-state, supported in-place analyze command, normal graph output, and post-state | pending |
| `E7-P7B-SITES1` | exact source-site inventory for four same-name declarations, six bindings, 21 direct exports, two barrel calls, and three ambient sites | pending |
| `E7-P7B-ORACLE1` | per-site source and graph/outcome proof producing exactly `4/4`, `6/6`, `21/21`, `2/2`, and `3/3` | pending |
| `E7-P7B-BOUNDARY1` | no target source/report/fixture contamination and no scanner-remediation claim | pending |
| `E7-P7B-REVIEW1` | independent Supervisor PASS | pending |
| `E7-P7B-DETECT1` | Anvien detect-changes before validation-artifact commit | pending |
| `E7-P7B-COMMIT1` | isolated Anvien evidence commit with no target artifact | pending |

Terminal table:

| Defect | Baseline | PASS | Evidence |
|--------|---------:|-----:|----------|
| Same-name identity | 2/4 | 4/4 | `E7-P7B-SITES1`, `E7-P7B-ORACLE1` pending |
| Binding patterns | 0/6 | 6/6 | `E7-P7B-SITES1`, `E7-P7B-ORACLE1` pending |
| Direct exports | 0/21 | 21/21 | `E7-P7B-SITES1`, `E7-P7B-ORACLE1` pending |
| Barrel calls | 0/2 | 2/2 | `E7-P7B-SITES1`, `E7-P7B-ORACLE1` pending |
| Ambient sites | 0/3 | 3/3 correct external/capability outcomes | `E7-P7B-SITES1`, `E7-P7B-ORACLE1` pending |

### P7-C — Affected-surface, runtime, and performance acceptance

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E7-P7C-INVENTORY1` | concrete affected-reader inventory derived from accepted diffs, impact evidence, changed fields, and actual call paths; every inclusion/exclusion justified | pending |
| `E7-P7C-BUILD1` | final full build and any required built app/container runtime | pending |
| `E7-P7C-RUNTIME1` | every included non-UI reader exercised at its nearest real boundary | pending |
| `E7-P7C-PARITY1` | each affected projection matches its accepted graph fact/outcome with zero unexplained loss or drift | pending |
| `E7-P7C-PLAY1` | reusable JSON/Markdown/visual evidence for affected browser-visible rows, or explicit not-applicable evidence when inventory contains none | pending |
| `E7-P7C-BENCH1` | same-method final performance measurements against a pre-recorded baseline and budget | pending |
| `E7-P7C-REVIEW1` | independent Supervisor PASS | pending |
| `E7-P7C-DETECT1` | Anvien detect-changes before final validation-artifact commit | pending |
| `E7-P7C-COMMIT1` | isolated P7-C evidence/benchmark/QA commit and known worktree | pending |

The affected-surface inventory must use this record shape for each concrete reader:

| Reader / boundary | Why affected | Field or behavior | Inclusion evidence | Validation result | Evidence ID |
|-------------------|--------------|-------------------|--------------------|-------------------|-------------|
| pending source-derived inventory | pending | pending | `E7-P7C-INVENTORY1` | pending | pending |

## Closure Evidence

| Evidence ID | Required proof | Status |
|-------------|----------------|--------|
| `E7-P7C-TERMINAL1` | P7-A, P7-B, and P7-C accepted; all failed rows returned to and closed by their owning child | pending |
| `E7-P7C-ROADMAP1` | terminal roadmap status refreshed with exact accepted commits and no successor | pending |
| `E7-P7C-SUPERVISOR1` | Supervisor accepts the complete campaign claim from source, diff, runtime, reports, visual evidence when applicable, benchmarks, and target boundary | pending |
| `E7-P7C-WORKTREE1` | final repository and target worktree states recorded; dead validation attempts removed | pending |
