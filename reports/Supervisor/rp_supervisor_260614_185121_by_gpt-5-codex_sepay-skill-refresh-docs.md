# Supervisor Report: SePay Skill Refresh Docs

Verdict: PASS

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260614_185121_by_gpt-5-codex_sepay-skill-refresh-docs.md`
- Review time: `260614 185121 +07:00`
- Reviewer: `gpt-5-codex`
- Repo/project: `Anvien`
- Scope reviewed: docs-only implementation of `docs/plans/2026-06-14-sepay-skill-refresh`
- Claim reviewed: P1-P6 source documentation refresh is complete for the SePay payment-integration references and direct routing surfaces, with validation evidence recorded.
- Authority used: latest user correction limiting the work to docs-only and no coder workflow; plan `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-plan.md`; official SePay docs checked on 2026-06-14; repo rules in `AGENTS.md`.
- Related artifacts: plan/evidence/benchmark/actual-status files under `docs/plans/2026-06-14-sepay-skill-refresh`; source skill docs under `internal/aicontext/skills/payment-integration`; generated repo skill outputs under `.agents/skills/payment-integration` and `.claude/skills/payment-integration`.

## Executive Summary

- Problem: SePay guidance mixed API v1, API v2, Webhooks, Payment Gateway/IPN, QR, and Order VA concepts, creating stale or unsafe integration guidance.
- Decision: PASS for the docs-only plan scope. Source references and routing now separate the SePay surfaces, make API v2 the default, add HMAC webhook guidance, add dedicated Payment Gateway/IPN guidance, refresh SDK/QR docs, and remove project-specific best-practice examples.
- Required outcome: accepted for P1-P6 docs scope. The stale `scripts/checkout-helper.js` URLs are recorded as out-of-scope follow-up because the user explicitly constrained this turn to docs-only.

## Source-Level Clearance Notes

- `internal/aicontext/skills/payment-integration/references/sepay/overview.md`: clear. It now routes agents between API v2, Webhooks, Payment Gateway/IPN, QR utility, Order VAs, optional products, Test Mode, and production safety.
- `internal/aicontext/skills/payment-integration/references/sepay/api.md`: clear. API v2 is the default; legacy `userapi/*` and 2 calls/second guidance are migration-only, not the default current path.
- `internal/aicontext/skills/payment-integration/references/sepay/webhooks.md`: clear. It documents bank-account webhooks, HMAC raw-body verification, retry/replay, payment-code recognition, monitoring, reconciliation, IP allowlisting, and explicitly excludes Payment Gateway IPN.
- `internal/aicontext/skills/payment-integration/references/sepay/payment-gateway.md`: clear. New reference covers gateway base URLs, Basic Auth, checkout init, IPN HTTPS/200 acknowledgement, `X-Secret-Key`, payload shape, notification types, idempotency, order lifecycle APIs, and reconciliation.
- `internal/aicontext/skills/payment-integration/references/sepay/sdk.md`: clear. It uses official Node/PHP SDK install paths, current payment methods, field-order signature warning, order detail/status/cancel/void topics, and does not present Laravel package guidance as the official gateway SDK path.
- `internal/aicontext/skills/payment-integration/references/sepay/qr-codes.md`: clear. It covers `showinfo`, `fullacc`, `holder`, `store`, `standee`, bank identifier variants, official VA, memo-based VA, no-VA, `TKP`, and VietinBank `SEVQR`.
- `internal/aicontext/skills/payment-integration/references/sepay/best-practices.md`: clear. It is repo-agnostic and no longer contains the previous CLAUDEKIT/product/POLAR/GitHub/coupon/referral examples as general SePay guidance.
- `internal/aicontext/skills/payment-integration/SKILL.md` and `references/implementation-workflows.md`: clear. Both make `payment-gateway.md` discoverable and route SePay work by surface first.
- `internal/aicontext/skills/payment-integration/scripts/checkout-helper.js`: not cleared for code behavior and not part of this PASS. It still contains stale SePay gateway URLs and is correctly recorded as out-of-scope follow-up in `E6-P6A-SCOPE1`.

## Evidence Checked

Passed:

- `git diff --check`: passed with no whitespace errors.
- `npm run full-build`: passed; included Go runtime packaging, web TypeScript/Vite production build, global install, and `anvien version` `1.2.6`.
- `anvien analyze . --force`: executed during full build; graph refreshed to 1395 files, 82455 nodes, 120516 relationships, stale false.
- `anvien file-detail internal/aicontext/skills/payment-integration/references/sepay/payment-gateway.md --repo Anvien --json`: new file parsed as markdown docs, 0 relationships, low risk, stale false.
- Generated repo output smoke: `.agents/skills/payment-integration` and `.claude/skills/payment-integration` contain `payment-gateway.md`, HMAC-SHA256 guidance, and `pgapi.sepay.vn`.
- `anvien detect-changes --repo Anvien --scope all`: passed; 12 changed files, 12 affected files, docs layer only, documentation functional area, affected processes none, risk low.
- Verification freshness: fresh in this review turn.

Failed:

- None for accepted docs-only scope.

Not run:

- Runtime browser/E2E: not required for docs-only skill reference changes.
- Helper script tests: not required because no script/code changes were in scope.
- Home-level `~/.agents/skills/payment-integration` regeneration: not performed by this repo source build; recorded as follow-up if installed home skills must be refreshed immediately.

## Invariant Closure

- affected invariant: SePay skill documentation must route agents to the correct SePay product surface and avoid stale/default v1, weak webhook, mixed IPN/webhook, and project-specific general guidance.
- sibling surfaces checked: source SePay reference docs, direct skill entrypoint, implementation workflow routing, generated repo `.agents` output, generated repo `.claude` output, plan/evidence/benchmark/actual-status ledgers.
- residual unverified same-invariant surfaces: none for the docs-only scope. `scripts/checkout-helper.js` is a related code/helper surface with stale URLs, but it is outside the user-corrected docs-only scope and must be handled by a separate code/script plan if expanded.

## Overall Evaluation

The implemented docs satisfy the plan acceptance criteria for P1-P6. The new guidance makes API v2 the default for new API work, separates bank-account Webhooks from Payment Gateway IPN, adds HMAC/replay/monitoring/reconciliation guidance, adds a dedicated gateway reference, refreshes SDK and QR guidance, and generalizes best practices. Validation evidence is current and sufficient for the docs-only claim.
