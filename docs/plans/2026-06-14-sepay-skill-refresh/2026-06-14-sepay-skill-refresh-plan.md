# Plan

## Metadata

- Title: `SePay Skill Refresh`
- Date: `2026-06-14`
- Status: `Complete; final commit executed in this turn`
- Plan: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-plan.md`
- Evidence: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-actual-status.md`

## Goal

Refresh the SePay references in the `payment-integration` skill so future agents choose the correct SePay product surface, use current official API/webhook/gateway guidance, and avoid stale v1, weak webhook, or mixed IPN/webhook assumptions.

## Rules

- Complete P0 actual status before implementation work.
- Update each checklist item immediately when it is completed.
- Record evidence as work completes.
- Record benchmarkable counts or measurements when they are taken.
- Update later phase status assumptions, next actions, and work steps when actual-status evidence changes the repo state.
- After completing a phase or implementation slice and refreshing `actual-status.md`, update the next affected phase's work steps as needed to match the latest repo reality while preserving that phase's original goal, scope, acceptance criteria, and major phase order.
- Do not edit generated `.agents/**`, `.claude/**`, `AGENTS.md`, or `CLAUDE.md` as source of truth.

## Problem

`reports/problem/2026-06-14-sepay-skill-refresh-analysis.md` records that the current SePay references are partially stale and mix SePay API v1, API v2, Webhooks, Payment Gateway/IPN, VietQR, and Order VA concepts into one mental model. That can make an agent implement the wrong integration path, use legacy endpoints for new work, verify webhooks too weakly, or treat Payment Gateway IPN as if it were the same surface as bank-account webhooks.

## Scope

- Primary SePay reference source:
  - `internal/aicontext/skills/payment-integration/references/sepay/overview.md`
  - `internal/aicontext/skills/payment-integration/references/sepay/api.md`
  - `internal/aicontext/skills/payment-integration/references/sepay/webhooks.md`
  - `internal/aicontext/skills/payment-integration/references/sepay/sdk.md`
  - `internal/aicontext/skills/payment-integration/references/sepay/qr-codes.md`
  - `internal/aicontext/skills/payment-integration/references/sepay/best-practices.md`
- New reference source if needed for separation:
  - `internal/aicontext/skills/payment-integration/references/sepay/payment-gateway.md`
- Minimal routing updates only if the new reference file is added:
  - `internal/aicontext/skills/payment-integration/SKILL.md`
  - `internal/aicontext/skills/payment-integration/references/implementation-workflows.md`
- Generated skill output validation:
  - generated `.agents/skills/payment-integration/**` and `.claude/skills/payment-integration/**` are validation outputs, not permanent source edits.

## Non-Goals

- Do not implement a SePay payment integration in an application repo.
- Do not change unrelated payment providers such as Polar, Stripe, Paddle, or Creem.io.
- Do not update `internal/aicontext/skills/payment-integration/scripts/sepay-webhook-verify.js` in this plan unless the user explicitly expands scope; record it as inspect-only/follow-up.
- Do not hardcode SePay public IP addresses as permanent guidance; link to the official page and record access evidence when exact IPs are quoted.
- Do not rewrite historical reports, old plans, or past evidence that mention stale SePay behavior.

## Requirements

- Use the problem report as the implementation authority and re-check official SePay docs before editing if implementation starts in a later session.
- Make API v2 the default for new API integrations and demote legacy `userapi/*` to migration/legacy guidance.
- Separate bank-account Webhooks from Payment Gateway IPN in language, examples, and routing.
- Add HMAC-SHA256 webhook verification guidance with raw body, timestamp, replay window, and timing-safe comparison.
- Correct webhook response/retry guidance and add payment-code recognition, monitoring, replay, incidents, reconciliation, and IP allowlist guidance.
- Add Test Mode quotas and Live/Test differences, including Live HTTPS requirements.
- Refresh QR generation parameters and bank-specific VA/memo rules.
- Refresh SDK guidance, gateway base URLs, Basic Auth, payment methods, form-signature field order, order status, cancel, detail, and void operations.
- Generalize best practices away from project-specific `CLAUDEKIT` examples.
- If adding `payment-gateway.md`, update only the routing surfaces needed for agents to discover it.

## Acceptance Criteria

- SePay reference docs make the correct first decision between API v2, Webhooks, Payment Gateway/IPN, QR utility, and Order VA.
- No SePay reference presents v1 `https://my.sepay.vn/userapi/`, 2 req/s, `sandbox.pay.sepay.vn/v1/init`, or `pay.sepay.vn/v1/init` as the default current path for new integrations.
- Webhook guidance includes HMAC-SHA256, payment-code recognition, retry timing, replay-safe idempotency, monitoring, reconciliation, and official IP allowlist guidance.
- Payment Gateway/IPN guidance is clearly separate from bank-account webhook guidance and discoverable from the skill entrypoint.
- QR and Order VA guidance includes the missing official parameters and bank-specific rules from the report.
- `best-practices.md` is repo-agnostic and does not teach one project's order/memo format as the general SePay pattern.
- Source skill changes are validated by readback, `git diff --check`, full build before final validation, generated skill smoke where applicable, and Anvien detect-changes before commit.

## Checklist

- [x] P0-A: Complete actual status before implementation work.
  - Goal: establish the real current state for the SePay reference refresh.
  - Work Steps:
    1. Read the problem report and current SePay reference files.
    2. Run `anvien analyze --force`.
    3. Run `anvien file-detail <path> --repo E:\Anvien --json` for each target reference file and direct routing/helper files.
    4. Classify each target as correct, partial, wrong, missing, or inspect-only.
    5. Update later phase status assumptions, next actions, and work steps from the P0 evidence.
  - Implementation Gate: no SePay reference editing starts until `2026-06-14-sepay-skill-refresh-actual-status.md` has a final P0 decision.
  - Acceptance: actual status identifies stale, partial, missing, generated-output, and inspect-only surfaces for this scope.
- [x] P1-A: Update SePay surface map and reference routing.
  - Goal: make the skill route agents to the right SePay surface before they read detailed API examples.
  - Work Steps:
    1. Update `overview.md` to define SePay product surfaces: API v2, bank-account Webhooks, Payment Gateway/IPN, QR utility, Order VAs, and optional Bank Hub/OAuth2/eInvoice/SoundBox.
    2. Replace stale environment/rate-limit summary with current API v2/gateway/Test Mode/Live guidance at summary level only.
    3. Add a concise "choose this surface when..." table.
    4. Add `payment-gateway.md` under `references/sepay` for hosted checkout and IPN rather than hiding IPN inside webhooks.
    5. Update `SKILL.md` and `references/implementation-workflows.md` only enough to make `payment-gateway.md` discoverable.
  - Implementation Gate: P0 status confirms `overview.md`, `SKILL.md`, and `implementation-workflows.md` are editable/routing surfaces and generated output is not edited as source.
  - Acceptance: a future agent can choose API v2 vs Webhooks vs Payment Gateway/IPN vs QR/Order VA before reading implementation details.
- [x] P2-A: Refresh API v2 and Order VA guidance.
  - Goal: make `api.md` current for proactive lookup, reconciliation, bank account/VA/order management, and legacy migration.
  - Work Steps:
    1. Rewrite the API base/auth/rate-limit section around API v2 production/sandbox base URLs and Bearer token auth.
    2. Document API v2 response envelope, pagination, UUID identifiers, integer money fields, and HTTP status/error behavior.
    3. Move legacy v1 `userapi/*` into a clearly labeled legacy/migration section.
    4. Add current Order VA coverage: supported banks, bank-specific fields, exact-amount/partial-payment rules, and production-only constraints where applicable.
    5. Cross-link to `webhooks.md`, `qr-codes.md`, and `payment-gateway.md` only where the integration surface changes.
  - Implementation Gate: P1-A routing is in place or actual-status has been refreshed with an alternate route that preserves the phase goal.
  - Acceptance: `api.md` no longer teaches legacy v1 as the default and includes the Order VA bank-specific constraints from the report.
- [x] P3-A: Refresh bank-account webhook security and operations.
  - Goal: make `webhooks.md` safe for production bank-transaction notifications without mixing it with gateway IPN.
  - Work Steps:
    1. Update webhook setup and auth methods to include None for testing, API Key, HMAC-SHA256 recommended, and OAuth2.
    2. Add HMAC verification requirements: raw body, `X-SePay-Signature`, `X-SePay-Timestamp`, `{timestamp}.{raw_body}`, replay window, and timing-safe compare.
    3. Correct success response, timeout, retry attempts, retry duration, and reconciliation guidance.
    4. Add payment-code structure guidance and explain when to prefer webhook `code` over custom memo parsing.
    5. Add monitoring, delivery logs, dashboard metrics, alerts, incidents, manual replay, and replay-safe idempotency guidance.
    6. Add official IP allowlist guidance as defense-in-depth, not as a replacement for authentication.
    7. Keep `scripts/sepay-webhook-verify.js` inspect-only/follow-up unless the user expands scope.
  - Implementation Gate: P2-A has not changed the webhook payload assumptions; if it has, refresh actual-status and update only this phase's stale work steps.
  - Acceptance: `webhooks.md` documents current bank-account webhook behavior and clearly says Payment Gateway IPN is a separate surface.
- [x] P4-A: Refresh Payment Gateway, IPN, and SDK guidance.
  - Goal: make hosted checkout/IPN and SDK usage current and separate from bank-account webhook reconciliation.
  - Work Steps:
    1. Fill `payment-gateway.md` with current gateway base URLs, Basic Auth, checkout init, IPN URL/acknowledgement, `X-Secret-Key`, payload shape, notification types, idempotency, and order reconciliation.
    2. Update `sdk.md` with `npm i sepay-pg-node`, PHP SDK status, supported `payment_method` values, one-time payment form generation, field-order signature warning, order list/detail, card void, QR cancel, and order/authentication statuses.
    3. Label Laravel package guidance accurately as official only if implementation evidence confirms it; otherwise mark it as separate/community or remove from official SDK path.
    4. Cross-link gateway/IPN guidance from `overview.md` and `implementation-workflows.md`.
  - Implementation Gate: P1-A has created or confirmed the gateway reference location; if not, update this phase's work steps from actual-status without changing the phase goal.
  - Acceptance: hosted checkout agents no longer need to infer gateway/IPN rules from webhook docs or stale SDK endpoint examples.
- [x] P5-A: Refresh QR generation and repo-agnostic best practices.
  - Goal: make QR, bank-specific matching, and production patterns current without project-specific examples.
  - Work Steps:
    1. Update `qr-codes.md` with `showinfo`, `fullacc`, `holder`, `store`, `standee`, bank identifier options, and cache/error guidance.
    2. Add examples for official VA, memo-based VA, no VA, and VietinBank `SEVQR` rules without implying every bank works the same way.
    3. Rewrite `best-practices.md` around general SePay patterns: surface selection, HMAC-first webhooks, payment-code recognition, idempotency, replay safety, API v2 reconciliation, Test Mode/Live separation, IP allowlisting, amount policy, and bank memo transformations.
    4. Remove or rewrite `CLAUDEKIT`, Polar, GitHub, coupon/referral, and other project-specific examples unless they are explicitly framed as non-general examples.
  - Implementation Gate: P3-A and P4-A have established the correct webhook/IPN separation so best practices can reference the right surface names.
  - Acceptance: `qr-codes.md` and `best-practices.md` are current, repo-agnostic SePay guidance and do not teach a single application's memo/order model as the default.
- [x] P6-A: Validate source skill, generated outputs, and graph change evidence.
  - Goal: prove the SePay skill refresh is internally consistent and generated agent surfaces can consume it.
  - Work Steps:
    1. Read back every changed SePay reference file and route file.
    2. Run `git diff --check`.
    3. Run `npm run full-build` before accepting final validation evidence for source skill changes.
    4. Run `anvien analyze --force` after source changes to refresh graph evidence and generated AI context output.
    5. Smoke-check generated `.agents/skills/payment-integration` and `.claude/skills/payment-integration` output if those directories are regenerated in the working tree.
    6. Run `anvien detect-changes --repo Anvien --scope all` before commit.
  - Implementation Gate: implementation phases P1-A through P5-A are complete or explicitly blocked.
  - Acceptance: validation evidence records what each command proves, generated output is not used as source of truth, and detect-changes is recorded before commit.
- [x] Pn-A: Call supervisor for the implemented-plan acceptance loop.
  - Goal: verify the completed plan work against the accepted plan, actual-status decisions, evidence, benchmark, changed files, generated output, and validation results before closure.
  - Work Steps:
    1. Call the supervisor skill to review the full completed plan work.
    2. If supervisor fails the work, return to the responsible implementation workflow/skill for the failed scope only.
    3. Re-run supervisor review after the fix.
    4. Repeat until supervisor passes or records a blocker.
  - Implementation Gate: all planned implementation phases must be completed or explicitly blocked before this review.
  - Acceptance: supervisor review passes, or the plan records a blocker with evidence and no closure is performed.
- [x] Pn-B: Remove dead work created during this plan.
  - Goal: ensure the final diff contains only artifacts that still serve the accepted plan.
  - Work Steps:
    1. Review files, sections, generated output, tests, temp files, and plan artifacts created or modified during this plan.
    2. Remove or rewrite any artifact made obsolete by actual-status findings, user corrections, failed approaches, or phase status updates.
    3. Verify no rejected approach, stale placeholder, unused generated output, or dead helper artifact remains in the final diff.
    4. Call supervisor to review the dead-work cleanup.
    5. If supervisor fails the cleanup, return to the responsible implementation workflow/skill for the failed cleanup scope only, then re-run supervisor review.
  - Implementation Gate: only remove artifacts created by this plan unless the user explicitly approves broader cleanup.
  - Acceptance: final `git diff/status` contains no dead plan-created artifacts, supervisor passes the cleanup, and evidence records what was removed or preserved.
- [x] Pn-C: Close the plan.
  - Goal: finish validation, evidence, benchmark, detect-changes, commit, and final status.
  - Work Steps:
    1. Run the required final validation for the accepted scope.
    2. Regenerate generated outputs if source-of-truth changes require it.
    3. Run Anvien detect-changes before commit when implementation work was performed.
    4. Record final validation, detect-changes, benchmark, and commit evidence.
    5. Commit the completed scope and verify the worktree state.
  - Implementation Gate: Pn-A and Pn-B must pass or record blockers.
  - Acceptance: final evidence is recorded, required commits exist, and the worktree state is known.

## Risk Notes

- The six existing SePay reference markdown files are low graph-risk docs, but their content is high product-risk because stale payment guidance can lead to money-handling mistakes.
- Adding a new `payment-gateway.md` without routing updates would create hidden guidance; routing updates are intentionally scoped.
- `scripts/sepay-webhook-verify.js` has high graph risk and a related test file; it is inspect-only in this plan unless the user explicitly expands scope.
- Official SePay docs can change; implementation should re-check official sources if work does not start immediately after this plan.
- Generated skill output must be treated as output evidence, not the permanent edit target.
