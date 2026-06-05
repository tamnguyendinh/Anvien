# Supervisor vs Code Review Skill Comparison

Date: 2026-06-05

Purpose: compare `internal/aicontext/skills/supervisor` and `internal/aicontext/skills/code-review` as input for discussing a new repo-agnostic skill that uses Anvien correctly across repositories/projects.

## Scope

Compared sources:

- `internal/aicontext/skills/supervisor`
- `internal/aicontext/skills/code-review`

This report evaluates:

- core essence of each skill;
- strengths and weaknesses;
- portability to arbitrary repositories;
- fit with Anvien's command model;
- what should be retained or discarded when designing a new Anvien-centered skill.

## Method Notes

`anvien status` reported the Anvien index as stale after recent commits. `anvien analyze --force` could not refresh the graph because unrelated untracked skill folders, starting with `internal/aicontext/skills/Architect-review`, do not contain a standard `SKILL.md` entry.

Therefore this report is based on direct file inspection, not graph-derived facts. That blocker is itself relevant: skill directories that do not follow the expected package shape can break Anvien analysis.

## Structural Comparison

| Area | `supervisor` | `code-review` |
|---|---|---|
| Current shape | Folder of long prompt files; main file is `supervisor.md`, not `SKILL.md` | Standard skill folder with `SKILL.md` and references |
| Metadata | Nonstandard frontmatter includes `When use` and capitalized `name: Supervisor` | Standard `name` and `description` frontmatter |
| Size | 15 files, many 200-380 lines each; main file is 381 lines | 4 files; `SKILL.md` is 102 lines |
| Disclosure model | Mostly monolithic and repeated across mode/lane variants | Progressive disclosure: entry skill plus focused references |
| Review stance | Independent supervisor verdict with PASS/REJECT/ESCALATE | Code-review discipline: receive feedback, request review, verify before claims |
| Repo assumptions | Strongly assumes `Docs/SPEC/*`, `Docs/execution/progress.md`, reports lanes, notes logs, Coder/Supervisor roles | Mostly generic, but assumes a `code-reviewer` subagent for review requests |
| Anvien usage | No native Anvien command model | No native Anvien command model |
| Best portability | Low without adaptation | Medium/high, except subagent-specific parts |

## Essence Of `supervisor`

The essence of `supervisor` is independent governance of implementation work. It treats every coder claim, report, test result, and completion statement as untrusted until independently verified against source, runtime behavior, authoritative specs, and report history.

Its core operating model is:

1. Select a review mode based on project execution state.
2. Resolve the authoritative spec family and owner boundary.
3. Read touched source before running tests/builds.
4. Consume all relevant reports, not only the latest report.
5. Verify runtime, gates, and sibling surfaces in the same invariant family.
6. Return a hard verdict: PASS, REJECT, or ESCALATE.

The most valuable idea is not "review code"; it is "close an invariant family with evidence." A localized diff is not enough if the same runtime contract, fail-close rule, isolation rule, or data-integrity rule spans sibling paths.

## Strengths Of `supervisor`

`supervisor` has strong review rigor:

- It rejects trust-by-report. Coder claims become pointers to evidence, not evidence themselves.
- It forces source inspection before build/test commands, preventing tests from masking source/spec drift.
- It requires review against authority, not preference.
- It distinguishes implementation fixes from architecture/spec authority.
- It consumes cross-lane reports and treats unresolved reports as active blockers.
- It has a useful concept of invariant-family closure: do not approve a one-symptom fix when adjacent same-family surfaces remain unchecked.
- It has explicit escalation routing: coder for implementation/process compliance, architect for spec or authority problems.
- It requires report artifacts, file/line evidence, and clear verdicts.

For high-risk, multi-agent work, this skill captures the most important habit: do not close work until the system boundary affected by the change is actually closed.

## Weaknesses Of `supervisor`

`supervisor` is weak as a reusable Anvien skill in its current form:

- It is not packaged as a standard skill: no `SKILL.md` at the skill root.
- It can break Anvien analysis because the skill package shape is invalid.
- It is too large and repetitive for progressive disclosure.
- It is heavily tied to one repo/process vocabulary: `Docs/SPEC/*`, `Docs/execution/progress.md`, `reports/coder`, `reports/QA`, `reports/Data-Integrity`, `notes_decisions_log`, Coder/Supervisor/Architect lanes.
- It assumes every repo has a spec-family authority model. Many repos only have README, tests, issue text, API docs, or code as authority.
- It does not know Anvien's specific workflow: `anvien analyze --force`, `impact`, `context`, `query`, `detect-changes`, route/tool/API mapping, graph-health, etc.
- It mixes universal review principles with project-specific rules, making it hard to activate safely in arbitrary repos.
- It creates a risk of process overreach: rejecting work because a target repo lacks the original process artifacts, not because the implementation is wrong.

The main conversion need is to extract the invariant/evidence discipline and replace repo-specific authority paths with a repo-discovery layer.

## Essence Of `code-review`

The essence of `code-review` is technical humility under review pressure. It prevents the agent from treating feedback, subagent output, or its own confidence as truth.

Its core operating model is:

1. When receiving feedback: read, understand, verify, evaluate, respond, implement.
2. When requesting review: compare work against plan/requirements using a review process.
3. Before claiming completion: run fresh verification and cite evidence.

The most valuable idea is "evidence before claims." It is intentionally small, easy to trigger, and easy to remember.

## Strengths Of `code-review`

`code-review` is strong as a skill package:

- It has standard frontmatter and a normal `SKILL.md`.
- It keeps `SKILL.md` compact and routes detail into references.
- It separates three different use cases cleanly: receiving feedback, requesting review, and verification gates.
- It is repo-agnostic in most places.
- It is psychologically useful: it blocks performative agreement, blind implementation, and premature success claims.
- It has a clear "stop and ask" rule for unclear feedback.
- It includes explicit verification gate logic: identify proof command, run it fresh, read output, then claim.

This skill is a good shape template for a future Anvien skill: concise entrypoint, focused references, and memorable laws.

## Weaknesses Of `code-review`

`code-review` is not enough for Anvien-governed implementation work:

- It does not use Anvien for codebase analysis, blast-radius checks, or change detection.
- It does not distinguish graph evidence from local grep/source inspection.
- It does not enforce impact-before-edit.
- It does not have a strategy for stale indexes or failed analysis.
- It assumes a `code-reviewer` subagent for requesting review, which may not exist in every environment.
- It focuses on feedback discipline more than end-to-end implementation governance.
- It does not model invariant-family closure or sibling-surface review.

The main conversion need is to add Anvien command selection and deeper review closure while keeping the skill small.

## Direct Comparison

`supervisor` is stronger at system-level closure. `code-review` is stronger at being a portable skill.

`supervisor` asks: "Is the whole affected invariant actually safe to approve?"
`code-review` asks: "Am I responding and claiming truth with evidence?"

`supervisor` is closer to a governance role.
`code-review` is closer to an agent behavior correction layer.

For a new Anvien skill, neither should be copied verbatim. The target should combine:

- from `supervisor`: zero-trust evidence, authority resolution, invariant-family closure, escalation classification, report discipline;
- from `code-review`: compact skill shape, progressive disclosure, clear trigger model, no premature claims, no blind feedback implementation.

## What To Preserve For A New Anvien Skill

Preserve from `supervisor`:

- Treat claims as pointers, not evidence.
- Inspect source and scope before validation commands.
- Resolve authority before judging drift.
- Review beyond the local diff when the same contract/invariant spans sibling paths.
- Separate implementation problems from architecture/spec/process-authority problems.
- Require explicit verdicts and evidence.

Preserve from `code-review`:

- Compact `SKILL.md`.
- Focused references below 150 lines where possible.
- Clear trigger sections.
- A simple decision tree.
- "Evidence before claims."
- "Ask before assuming unclear feedback."
- "Verify external suggestions before implementing."

## What To Drop Or Rewrite

Drop or rewrite from `supervisor`:

- Hard dependency on `Docs/SPEC/*`.
- Hard dependency on `Docs/execution/progress.md`.
- Hard dependency on lane directories that not every repo has.
- Large repeated prompt variants.
- Nonstandard skill root shape.
- Project-specific severity policy that says every MEDIUM blocks all work.

Drop or rewrite from `code-review`:

- Assumption that a `code-reviewer` subagent is always available.
- Review request flow based only on git SHAs.
- Generic verification gates without Anvien-specific pre-edit and pre-commit gates.

## Proposed New Skill Direction

Working name: `anvien-work-governance` or `anvien-implementation-review`.

Purpose: guide an agent through safe implementation and review in any repo where Anvien is available.

Core principle:

```text
Use Anvien to know scope. Use source to verify facts. Use tests/builds to validate behavior. Use detect-changes before commit.
```

The skill should be repo-agnostic. It should not require `Docs/SPEC/*`; instead it should discover authority in this order:

1. Active user request and current plan.
2. Repo instructions: `AGENTS.md`, `CLAUDE.md`, README, contribution docs.
3. Project-specific plan/spec docs if present.
4. API contracts, tests, schema, route/tool maps, generated contracts.
5. Source code and runtime behavior.

It should use Anvien in a practical sequence:

1. `anvien status` or `anvien analyze --force` before graph-based work.
2. `anvien query` or `file-context` to locate relevant code.
3. `anvien impact` before editing functions/classes/methods/shared contracts.
4. Route/tool/API map commands when changing handlers or contracts.
5. Source edits and tests.
6. `anvien detect-changes --repo <repo> --scope all` before commit.

It should degrade cleanly:

- If Anvien index is stale and refresh succeeds, use refreshed graph evidence.
- If refresh fails because unrelated invalid skill packages or generated artifacts block analyze, record the blocker and continue with bounded direct source inspection only when the task can still be done safely.
- If the requested edit depends on graph truth that cannot be refreshed, stop or ask for user direction.

## Suggested Skill Structure

```text
internal/aicontext/skills/anvien-work-governance/
  SKILL.md
  references/
    command-selection.md
    implementation-slice.md
    review-and-closure.md
    stale-index-and-blockers.md
```

`SKILL.md` should stay short:

- when to use;
- core principle;
- decision tree;
- required Anvien gates;
- references to detailed workflows.

References should split responsibilities:

- `command-selection.md`: which Anvien command to use by task.
- `implementation-slice.md`: plan, impact, edit, build/test, evidence.
- `review-and-closure.md`: invariant-family review, report handling, detect-changes, commit.
- `stale-index-and-blockers.md`: how to handle stale graph, invalid package shape, or blocked analyze.

## Draft Operating Modes

Mode 1: Implementation Slice

- Use when the agent is changing code or skill source.
- Plan first.
- Refresh/use Anvien.
- Run impact before edits.
- Implement narrowly.
- Validate.
- Detect changes.
- Commit if requested or if repo rules require it.

Mode 2: Review / Problem Report

- Use when reviewing code, reviewing feedback, or writing `reports/problem/*`.
- Resolve authority.
- Inspect relevant source and reports.
- Use Anvien to find blast radius where possible.
- Produce findings first, ordered by severity.
- Separate code defects from process/spec/authority defects.

Mode 3: Closure Gate

- Use before claiming done, moving to the next task, or committing.
- Verify fresh command output.
- Check diff scope.
- Run detect-changes.
- Record unresolved blockers explicitly.

## Key Design Risks For The New Skill

1. If it copies `supervisor` too closely, it will become too repo-specific and too heavy.
2. If it copies `code-review` too closely, it will be too generic and will not enforce Anvien's graph/impact/detect-changes discipline.
3. If it makes Anvien mandatory even when graph refresh is blocked by unrelated invalid files, it will halt useful doc/report work unnecessarily.
4. If it allows bypassing Anvien silently, it will lose the tool's main value.
5. If it does not define authority discovery for arbitrary repos, agents will either overfit to Anvien's own repo or invent process rules.

## Recommendation

Build the new skill as a compact Anvien-centered governance skill, not as a bigger code-review skill and not as a direct port of `supervisor`.

The new skill should take this hierarchy:

1. `code-review` shape and discipline for usability.
2. `supervisor` zero-trust and invariant-family closure for rigor.
3. Anvien command gates for repo intelligence.
4. Repo-agnostic authority discovery for portability.

The result should help an agent work safely in any repo/project:

- know where to look;
- know blast radius before editing;
- verify claims with evidence;
- detect affected flows before commit;
- produce useful reports when the process or codebase blocks safe completion.

---
 ---
  name: supervisor-review
  description: Use whenever reviewing completion claims, fixes,
  diffs, plans, reports, screenshots, or artifacts for acceptance
  by verifying repo/project reality and Anvien evidence.
  ---

  # Supervisor Review

  Use this skill to independently decide whether a claim, artifact,
  work result, or completion statement can be accepted in a repo/
  project.

  The review target is only the entry point. The real job is to
  verify the claim against repo/project reality.

  ## Core Law

  Zero-trust review: treat every claim, artifact, result, and
  completion statement as untrusted until independently verified
  against repo/project reality with sufficient evidence.

  ## Role Boundary

  Supervisor Review gives an acceptance verdict. It does not repair
  the work while reviewing unless the user explicitly asks for a
  separate implementation task.

  Do not blend review and fix work. If review finds a problem,
  return a precise REJECT with the evidence and the required next
  step.

  ## Compact-Safe Re-anchor

  After any compact, resume, long gap, or confusing thread, re-
  anchor before verdict:

  - reload the latest user request and current review scope;
  - read applicable repo instructions such as `AGENTS.md`;
  - inspect the current artifact, diff, report, screenshot, log,
  plan, or result being reviewed;
  - discard any prior conclusion that is not proven against current
  evidence.

  Do not continue a previous PASS/REJECT by inertia.

  ## Start Here

  1. Understand the review problem, not just the words or artifact.
  2. Reconstruct the claim being asked for acceptance.
  3. Identify the authority that decides whether the claim is
  correct.
  4. Determine what repo/project reality must be checked.
  5. Gather current evidence from source, runtime, tests, docs,
  repo authority, Anvien, logs, data, or other relevant tools.
  6. Close the affected invariant, not only the visible symptom.
  7. Decide PASS or REJECT from the evidence.
  8. Write the verdict using the required output format.

  ## Authority

  Use the strongest applicable authority:

  1. latest user instruction;
  2. repo rules such as `AGENTS.md`;
  3. active plan, spec, issue, PR, acceptance criteria, or owner
  decision;
  4. contracts, schemas, APIs, generated contracts, tests, docs,
  source code, runtime behavior, and data/source-of-truth state.

  Reports, plans, screenshots, tests, logs, diffs, and tool output
  are evidence. They are not authority by themselves.

  If authority conflicts and the conflict blocks acceptance, REJECT
  and name the conflict.

  ## Claim-To-Evidence Conversion

  Before judging, convert the input into a review claim:

  - What is being claimed explicitly or implicitly?
  - What would have to be true for the claim to be accepted?
  - What authority defines true, complete, and acceptable?
  - Which repo/project surfaces can prove or disprove it?
  - What evidence would be enough for PASS?

  If the claim cannot be reconstructed, REJECT with the missing
  information needed to make it reviewable.

  ## Source Inspection Gate

  When the claim depends on code, inspect source before relying on
  build, test, report, or tool summaries.

  For code changes, bug fixes, wiring claims, contract claims,
  runtime claims, or generated output claims:

  - inspect the relevant diff or files first;
  - read touched production code before validation commands;
  - inspect affected source paths before trusting tests;
  - do not let a green test replace source review.

  If source inspection is required but unavailable, REJECT.

  ## Evidence Protocol

  Gather evidence from the strongest source needed for the review
  problem.

  Use Anvien when codebase evidence is needed to locate behavior,
  map affected files/symbols/routes/tools/contracts, inspect
  dependencies or impact, find sibling surfaces, or prove whether
  the claim covers the full invariant.

  Choose tools from the evidence question. Do not run tools as a
  ritual checklist.

  Evidence must be:

  - current for the reviewed repo/project state;
  - specific to the full claim;
  - traceable to source, runtime, command output, data, docs,
  authority, or Anvien result;
  - strong enough to prove acceptance, not just suggest confidence.

  Missing, stale, indirect, partial, or narrower evidence cannot
  support PASS.

  ## Invariant Closure

  Do not approve a local symptom fix when the same invariant may
  span other surfaces.

  Identify the affected invariant: runtime contract, data integrity
  rule, owner boundary, permission rule, isolation rule, API shape,
  tool contract, state transition, generated artifact contract, or
  process rule.

  Start from the provided artifact or diff, then sweep only the
  relevant same-invariant surfaces, such as:

  - route or entrypoint;
  - alternate trigger;
  - UI panel, dialog, or state path;
  - store, service, API, tool handler, repository, schema, job,
  worker, or generated contract;
  - stale helper, fallback path, fixture, test, or doc contract
  when it can preserve the old behavior.

  Do not expand into unrelated domains. Do not approve until the
  affected invariant is closed for the reviewed scope.

  ## Report And History Closure

  If prior reports, review comments, blocker notes, QA findings,
  bug reports, or resubmissions exist in the same scope, consume
  them as evidence pointers.

  Do not read only the latest artifact when unresolved earlier
  evidence can still affect acceptance.

  A prior issue is closed only when current evidence proves it is
  closed. If closure cannot be proven, REJECT.

  ## Approval Standard

  PASS only when all are true:

  - the real claim is clear;
  - evidence proves the full claim, not a narrower claim;
  - the affected invariant is closed for the reviewed scope;
  - no required follow-up remains before acceptance.

  REJECT when any are true:

  - the claim is false, incomplete, unsafe, or misleading;
  - evidence is missing, stale, indirect, partial, or narrower than
  the claim;
  - source/project reality contradicts the claim;
  - authority conflicts or is missing for acceptance;
  - the fix only addresses the visible symptom while same-invariant
  surfaces remain unchecked or broken;
  - any required action remains before acceptance.

  ## Always Do

  - State the real claim and authority before judging.
  - Verify the full claim against repo/project reality before PASS.
  - Inspect source before trusting build/test/report output when
  code reality matters.
  - Use Anvien when codebase topology, impact, contracts,
  dependencies, or affected flows matter.
  - Review the affected invariant, not only the visible symptom or
  changed lines.
  - Include direct evidence, preferably file/line evidence when
  source is involved.
  - Give exactly one verdict: PASS or REJECT.

  ## Never Do

  - Never trust a claim, report, result, or completion statement by
  itself.
  - Never review only the surface artifact.
  - Never assume the current claim matches a previously seen
  pattern; verify against the actual artifact and repo/project
  state.
  - Never use Anvien or any tool as a fixed command checklist.
  - Never treat Anvien or any tool output as the verdict by itself.
  - Never approve from tests alone when source/runtime reality
  still needs inspection.
  - Never ignore unresolved same-scope reports, blocker notes, or
  review findings.
  - Never claim PASS from missing, stale, indirect, partial, or
  narrower evidence.

  ## Output

  Use exactly one verdict:

  - PASS: evidence proves the full claim is true, complete, and
  acceptable against repo/project reality.
  - REJECT: the claim is false, incomplete, unsafe, misleading,
  unsupported, only partially proven, or blocked by unresolved
  authority/scope/architecture conflict.

  ```text
  Verdict: PASS | REJECT
  Claim reviewed: <the claim being reviewed>
  Authority: <the authority used to judge the claim>
  Evidence: <evidence sources and what each proves>
  Invariant closure: <affected invariant and sibling surfaces
  checked, or why none were needed>
  Reasoning: <why the evidence proves the full claim or requires
  REJECT>
  Next step: <required action when verdict is REJECT; omit when
  PASS>

  