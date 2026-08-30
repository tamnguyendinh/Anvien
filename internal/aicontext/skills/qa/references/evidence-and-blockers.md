# Evidence And Blockers

> This file is part of QA Skill v2. Read when: collecting evidence, handling test results during QA, or processing blockers that prevent downstream verification.

## Tests In QA

Tests are supporting evidence only.

Rules:
- Do not write or edit tests during QA unless the user gives a separate order.
- Do not patch test failures silently.
- Do not use tests to legitimize broken behavior.
- Tests must model real user flows with all required steps.
- If tests mark wrong behavior as expected, report a test defect or test gap.

Examples of wrong expectations:
- A valid account user seeing only `Account surface is locked`.
- A valid admin seeing only `Admin surface is locked`.
- A signed-out user with no clear login/return path.
- A wrong-role user without a clear denied state.

## Evidence

Record evidence as work completes.

Evidence sources may include:
- build output
- runtime health output
- visible-browser proof
- screenshots, videos, traces
- Playwright report from the visible run
- route/control inventory
- visible-surface inventory
- interactive inventory
- click/input action ledger
- no-result/no-op report
- console/network findings
- DB/source readback evidence
- seed or fixture ledger
- source-of-truth mismatch report
- benchmark counts and runtime metrics when benchmarkable

Do not approve if required evidence, benchmarks, inventories, or ledgers are missing.

## Blocker Propagation

If a blocker prevents reaching downstream flows:

- stop if the blocked gate is required for the current run
- report the downstream flows as `Blocked`
- do not mark them passed
- do not imply coverage
- do not invent confidence
- list exactly what remains unverified and why

Example: if the shell never mounts, page-level flows behind the shell are blocked, not verified.
