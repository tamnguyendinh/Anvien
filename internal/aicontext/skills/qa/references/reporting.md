# Reporting

> This file is part of QA Skill v2. Read when: classifying findings, writing QA reports, handling handoff, or managing QA artifacts/commits.

## Finding Classification And Severity

Classify each finding:

| Classification | Meaning |
| --- | --- |
| Confirmed runtime bug | Reproduced on a mounted path. |
| Likely runtime bug | Strong signal, but missing one runtime confirmation. |
| Usability friction | Technically works, but confuses, slows, or misleads a real user. |
| Spec drift | Runtime and docs disagree, but user-visible harm is not fully confirmed. |
| Test gap | Coverage is insufficient to conclude behavior safely, or tests encode the wrong behavior. |

Severity:

| Severity | Meaning |
| --- | --- |
| CRITICAL | User cannot complete a core flow, wrong scoped data is shown, or money/protected action is exposed incorrectly. |
| HIGH | Mounted runtime path is wrong, feature is unreachable, or displayed data is materially incorrect. |
| MEDIUM | Degraded but usable. |
| LOW | Polish only. |

For every finding, include:
- user-visible symptom first
- technical path second, if known
- severity
- classification
- confidence: high / medium / low
- reproducibility: always / intermittent / once
- persona/context/viewport
- route/dialog/flow/control
- expected result
- actual result
- repro steps
- evidence references
- blocked downstream scope
- workaround if one exists
- suggested fix direction if useful, without fixing

## Required QA Report Shape

When writing a QA artifact, create a new file. Do not overwrite old QA reports.

Typical QA report path:

```text
reports/QA/rp_qa_<YYMMDD>_<HHMMSS>_by_<model_slug>_<scope>.md
```

Report template:

```text
# QA Report

Scope:
Runtime:
Target (Browser / API / DB):
Build evidence:
Discovery evidence:
No-fix boundary: QA reports only. No code fixes were made.

## Runtime Contexts
- persona / client:
- owner/app/scope:
- role / auth:
- shift/session/subscription:
- dataset/fixture / seed:
- locale / headers:
- viewport / endpoint target:

## Inventory Summary
- visible surfaces / API endpoints:
- interactive elements / payload scenarios:
- actions / mutations:
- data/source-of-truth checks:
- blocked:
- out of scope:
- unverified:

## Coverage Verdict
- Pass:
- Fail:
- Blocked:
- N/A:
- 100% of declared scope: yes/no

## Action Ledger

### Frontend Action Ledger (UI scopes)
| Surface | Route/flow | Control/action | Context | Expected | Actual | Verdict | Evidence (PNG) |
| --- | --- | --- | --- | --- | --- | --- | --- |

### Backend Action Ledger (API & DB scopes)
| Endpoint & Method | Payload Scenario | Expected Status & Schema | Actual Status & Payload | DB Diff Verified | Latency (ms) | Verdict | Evidence/Log |
| --- | --- | --- | --- | --- | --- | --- | --- |

## Findings
### [SEVERITY] Title
Classification:
Confidence:
Reproducibility:
Persona/context/viewport:
Route/dialog/flow/control:
Expected:
Actual:
Repro steps:
Evidence:
Downstream blocked scope:
Suggested fix direction:

## Source-Of-Truth Checks
| Target (UI / Endpoint) | Runtime source / DB Table | Expected source state | Actual source state | Verdict | Evidence |
| --- | --- | --- | --- | --- | --- |

## Console/Network/Runtime Evidence

## Blockers And Unverified Scope

## Handoff
Handoff: <architect|supervisor|coder|edge-case> - <reason>

## Final Decision
Pass / Fail / Blocked for declared scope only.
```

## Handoff Rules

- (MUST) End every QA report with exactly one concise handoff line.
- Handoff to `architect` when the next step needs architecture, source-of-truth, spec, system-flow, or rule decisions.
- Handoff to `supervisor` when the next step needs acceptance, priority, scope, coordination, or QA/fix order decisions.
- Handoff to `coder` when the issue is a confirmed implementation bug with enough evidence to start a fix.
- Handoff to `edge-case` (via CEO routing) when baseline functional verification passes but the scope involves high-contention mutations, critical locks, atomic counters, or requires hostile chaos / stress testing.
- If there is no finding or blocker, use `Handoff to supervisor - QA scope completed; awaiting acceptance decision.`

## Artifact And Commit Rules

- QA artifacts may be written if requested by the user or required by the active QA plan.
- Commit QA artifacts only when active repo/user rules explicitly require it.
- If the user says not to commit, do not commit.
- QA must stop after reporting and wait for a separate user fix order.
