# Lifecycle Plan Review Context

> This file is part of Edge-Case Review Skill v2. Read when: executing a multi-part or multi-cluster edge-case test plan (e.g. `reports/Edge-Case/QA+EDGE_CASE_TEST_PLAN.md`).

---

## Context Overview

Use this context guide when conducting comprehensive edge-case sweeps driven by a dedicated, structured plan file (such as `reports/Edge-Case/QA+EDGE_CASE_TEST_PLAN.md`).

**Core Mission:** Systematically execute structured test clusters, tracking ordinal progress and discovering systemic broken invariants.

---

## Progress Tracking & Cluster Rules

1. **Plan Sourcing**: Read the active plan file (e.g. `reports/Edge-Case/QA+EDGE_CASE_TEST_PLAN.md`) to determine full scope, total parts, and cluster allocations.
2. **Cluster Definition**: A standard cluster typically contains 5 parts (the final cluster may contain fewer if remaining parts < 5).
3. **Ordinal Progress Marker**:
   - Edge-Case MAY read only the latest prior Edge-Case report owned by the Edge-Case lane for that same overall plan scope, strictly to obtain the **ordinal progress marker** (the last completed part/cluster).
   - *Example:* If the previous report stopped at Part 15, the current run MUST start from Part 16 (Cluster 4).
   - If the previous report shows the overall plan is 100% complete, restart from the first cluster as a fresh rerun on the current head.
   - Do NOT use prior reports for content, hints, or templates.
4. **Single-Cluster Run Limit**: A single Edge-Case execution run processes exactly **one cluster**. Do not proceed to the next cluster in the same turn.

---

## Execution Steps

1. Determine the active cluster to execute based on the ordinal progress marker.
2. For each part within the cluster:
   - Resolve the exact SPEC family.
   - Select the matching domain attack guides from `references/`.
   - Construct and execute the perturbation matrix.
3. Record cumulative coverage in the run report:
   - Current cluster number & part range.
   - Cumulative parts completed vs. remaining.
   - Intermediate update vs. final closure indicator.
   - Cumulative broken invariants discovered.
4. Write a new report file and commit per `references/reporting-and-evidence.md`.
