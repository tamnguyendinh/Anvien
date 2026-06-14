# Evidence Ledger

## Metadata

- Title: `SePay Skill Refresh`
- Date: `2026-06-14`
- Plan: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-plan.md`
- Evidence: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-evidence.md`
- Benchmark: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-benchmark.md`
- Actual status: `docs/plans/2026-06-14-sepay-skill-refresh/2026-06-14-sepay-skill-refresh-actual-status.md`

## Evidence Rules

The evidence file explains why the work is known to be correct.

Evidence IDs use stable, phase-scoped names so `plan.md`, `actual-status.md`, `benchmark.md`, and later agents can reference exact proof without ambiguity.

Format:

```text
E<phase>-<item>-<kind><n>
```

Rules:

- `E<phase>` matches the plan phase number: `E0` for `P0`, `E1` for `P1`, `E2` for `P2`, and so on.
- `<item>` matches the checklist item without the dash: `P0A`, `P1A`, `P2A`.
- `<kind>` is plan-local. This plan uses `GRAPH`, `FD`, `SRC`, `REPORT`, `ROUTE`, `VAL`, `IMPL`, `GEN`, `DETECT`, and `COMMIT` as needed.
- `<n>` is a 1-based sequence number within that phase item and kind.
- Do not reuse an evidence ID for different facts.
- Reference exact IDs from `actual-status.md` and `benchmark.md`.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-GRAPH1`: Ran `anvien analyze --force` from `E:\Anvien`. Result: graph refreshed successfully; scanned 1390 files, parsed 673 code files, indexed 480 documents, graph nodes 82421, relationships 120482, graph path `.anvien/graph.json`, stale false.
- `E0-P0A-REPORT1`: Read `reports/problem/2026-06-14-sepay-skill-refresh-analysis.md`. The report records 16 current-state findings: API v2 default, rate limit correction, HMAC webhook auth, retry correction, webhook/IPN separation, gateway endpoint/auth refresh, Test Mode, QR gaps, Order VA bank rules, SDK refresh, payment-code recognition, webhook operations, IP allowlisting, Test Mode quotas/live differences, Payment Gateway IPN contract, and SDK/order lifecycle.
- `E0-P0A-SRC1`: Read `internal/aicontext/skills/payment-integration/references/sepay/overview.md`. Current content presents stale sandbox/production gateway endpoints and 2 calls/second rate limit, and does not separate current SePay product surfaces enough for routing.
- `E0-P0A-SRC2`: Read `internal/aicontext/skills/payment-integration/references/sepay/api.md`. Current content uses legacy `https://my.sepay.vn/userapi/` as the base URL and does not make API v2 the default.
- `E0-P0A-SRC3`: Read `internal/aicontext/skills/payment-integration/references/sepay/webhooks.md`. Current content lacks HMAC-SHA256, payment-code structure, monitoring/replay/incidents, precise retry behavior, and IP allowlist guidance; it also risks conflating webhooks with gateway callback behavior.
- `E0-P0A-SRC4`: Read `internal/aicontext/skills/payment-integration/references/sepay/sdk.md`. Current content uses old Node install form and stale gateway endpoints and does not cover current payment methods, field-order signature warning, order lifecycle, IPN, or gateway API details.
- `E0-P0A-SRC5`: Read `internal/aicontext/skills/payment-integration/references/sepay/qr-codes.md`. Current content covers basic QR parameters but misses `showinfo`, `fullacc`, `holder`, `store`, `standee`, bank identifier variants, and bank-specific VA/memo rules.
- `E0-P0A-SRC6`: Read `internal/aicontext/skills/payment-integration/references/sepay/best-practices.md`. Current content is project-specific, including `CLAUDEKIT`, Polar, GitHub, coupon/referral, and product-specific examples, and should be generalized.
- `E0-P0A-ROUTE1`: Read `internal/aicontext/skills/payment-integration/SKILL.md` and `internal/aicontext/skills/payment-integration/references/implementation-workflows.md`. Current routing lists six SePay references and no dedicated Payment Gateway/IPN reference.
- `E0-P0A-SRC7`: Read `internal/aicontext/skills/payment-integration/scripts/sepay-webhook-verify.js`. Current helper supports API Key/OAuth2/none and payload validation but no HMAC raw-body signature verification. It is inspect-only/follow-up for this plan.
- `E0-P0A-FD1`: `anvien file-detail internal/aicontext/skills/payment-integration/references/sepay/overview.md --repo E:\Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD2`: `anvien file-detail internal/aicontext/skills/payment-integration/references/sepay/api.md --repo E:\Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD3`: `anvien file-detail internal/aicontext/skills/payment-integration/references/sepay/webhooks.md --repo E:\Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD4`: `anvien file-detail internal/aicontext/skills/payment-integration/references/sepay/sdk.md --repo E:\Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD5`: `anvien file-detail internal/aicontext/skills/payment-integration/references/sepay/qr-codes.md --repo E:\Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD6`: `anvien file-detail internal/aicontext/skills/payment-integration/references/sepay/best-practices.md --repo E:\Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD7`: `anvien file-detail internal/aicontext/skills/payment-integration/SKILL.md --repo E:\Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD8`: `anvien file-detail internal/aicontext/skills/payment-integration/references/implementation-workflows.md --repo E:\Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false.
- `E0-P0A-FD9`: `anvien file-detail internal/aicontext/skills/payment-integration/scripts/sepay-webhook-verify.js --repo E:\Anvien --json` returned JavaScript source, 18 symbols, 15 local relationships, 4 inbound references from `scripts/test-scripts.js`, 89 unresolved entries, high risk, stale false.
- `E0-P0A-VAL1`: `Test-Path docs/plans/2026-06-14-sepay-skill-refresh` returned false before plan creation, so the selected standard plan directory did not collide with an existing plan.

## E1 - P1 Evidence

Matching plan item(s): `P1-A`

Record implementation evidence for SePay surface map, `overview.md`, `payment-gateway.md` creation, and minimal routing edits here.

Planned evidence IDs:

- `E1-P1A-IMPL1`: changed SePay surface map and routing.
- `E1-P1A-SRC1`: readback of changed `overview.md`.
- `E1-P1A-SRC2`: readback of new `payment-gateway.md` skeleton or completed surface if created in P1-A.
- `E1-P1A-ROUTE1`: readback of changed `SKILL.md` and `implementation-workflows.md`.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`

Record implementation evidence for API v2, legacy migration, rate limit, pagination, response shape, and Order VA bank rules here.

Planned evidence IDs:

- `E2-P2A-IMPL1`: changed `api.md`.
- `E2-P2A-SRC1`: readback proving v2 is default and v1 is legacy.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`

Record implementation evidence for webhook HMAC, payment code, retry, monitoring/replay, IP allowlist, and webhook/IPN separation here.

Planned evidence IDs:

- `E3-P3A-IMPL1`: changed `webhooks.md`.
- `E3-P3A-SRC1`: readback proving HMAC and replay-safe guidance.

## E4 - P4 Evidence

Matching plan item(s): `P4-A`

Record implementation evidence for Payment Gateway, IPN, SDK, order lifecycle, and Laravel package classification here.

Planned evidence IDs:

- `E4-P4A-IMPL1`: changed `payment-gateway.md` and `sdk.md`.
- `E4-P4A-SRC1`: readback proving gateway/IPN is separated from bank webhooks.

## E5 - P5 Evidence

Matching plan item(s): `P5-A`

Record implementation evidence for QR parameter/bank rule updates and generalized best practices here.

Planned evidence IDs:

- `E5-P5A-IMPL1`: changed `qr-codes.md` and `best-practices.md`.
- `E5-P5A-SRC1`: readback proving project-specific guidance was removed or reframed.

## E6 - P6 Evidence

Matching plan item(s): `P6-A`

Record validation and generated-output evidence here.

Planned evidence IDs:

- `E6-P6A-VAL1`: `git diff --check`.
- `E6-P6A-VAL2`: `npm run full-build`.
- `E6-P6A-GRAPH1`: final `anvien analyze --force`.
- `E6-P6A-GEN1`: generated skill output smoke, if generated output changes.
- `E6-P6A-DETECT1`: `anvien detect-changes --repo Anvien --scope all`.

## Closure Evidence

Use this section for supervisor review, dead-work cleanup, final detect-changes, commit hash, and closure evidence when the plan reaches completion.

Planned evidence IDs:

- `E7-PNA-SUP1`: supervisor acceptance review.
- `E8-PNB-CLEAN1`: dead-work cleanup evidence.
- `E9-PNC-COMMIT1`: final commit hash and worktree state.
