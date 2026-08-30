# Reporting & Evidence Standards

> This file is part of Edge-Case Review Skill v2. Read when: classifying findings, formatting reports, applying timestamp naming rules, or verifying artifact commits.

---

## Evidence Bar

Every finding recorded by Edge-Case MUST be classified into one of two categories:

- **`Confirmed Failure`**: A concrete breakage surfaced by executing a real attack path, or an execution-backed proof from a runtime-authentic perturbation exercising the relevant code chain under stress.
- **`Risky Gap`**: No executed breakage confirmed yet, but the current code path presents an attackable fail-open, data corruption, stale-state, or replay vulnerability that is not safely guarded.

> **Rules:**
> - Missing tests alone is NOT a finding.
> - Passive code reading alone is NOT sufficient to close a runnable scope.
> - NEVER present a `Risky Gap` as a confirmed failure without execution proof.

---

## Severity Classification

| Severity | Criteria / Impact Scope |
|---|---|
| **CRITICAL** | Scope / tenant leakage, auth bypass, money / shift bypass, duplicate financial mutations, broken distributed lock semantics |
| **HIGH** | State corruption, replay drift, stale-store action enablement, fail-open permission checks |
| **MEDIUM** | Recoverable state inconsistency, unhandled UI error state without data corruption |
| **LOW** | Noisy log output, minor cosmetic retry glitch |

---

## Required Issue Format

Every issue in the report must follow this standardized format:

```text
[CRITICAL] Payment can be triggered with no active shift after reconnect
Perturbation: reconnect after stale store restore
Expected: payment button remains blocked until active shift is revalidated
Actual: stale client state re-enables payment action
Broken invariant: money functions require active shift
Files:
- electron/renderer/src/features/orders/store/useOrderStore.ts:120
- backend/internal/service/payment_service.go:77
```

---

## Report Naming Rules

Write all Edge-Case reports to:
```text
reports/Edge-Case/rp_edge_<YYMMDD>_<HHMMSS>_by_<model_slug>_<scope>.md
```

- `<YYMMDD>_<HHMMSS>` MUST reflect the realtime creation timestamp.
- `model_slug`: stable lowercase ASCII slug (e.g. `claude-3-7-sonnet`).
- `scope`: lowercase `snake_case` summary.
- NEVER overwrite or append to older Edge-Case reports — each run creates a fresh report file.
- If a finding represents a shared blocker for other lanes, also write:
  ```text
  reports/problem/pb_edge_yymmdd_hhmmss_<scope>.md
  ```

---

## Artifact Commit Rules

- When this lane produces reports or blocker files, it MUST stage and commit those files before finishing.
- **Commit Scope:**
  - `reports/Edge-Case/*`
  - `reports/problem/*` (when created by Edge-Case)
- Do NOT commit transient test artifacts, `.tmp/`, or screenshots unless explicitly requested.
- Verify git status is clean for owned artifacts and report the commit hash.
