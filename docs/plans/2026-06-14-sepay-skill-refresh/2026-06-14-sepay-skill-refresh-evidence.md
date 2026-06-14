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

- `E1-P1A-OFFICIAL1`: Re-checked official SePay docs before editing source docs: API v2 overview/quick start, webhook authentication/error handling/monitoring/payment-code structure, Payment Gateway API/IPN/Node/PHP SDK/order APIs, QR generation, Test Mode quota/webhook pages, and IP address page.
- `E1-P1A-IMPL1`: Replaced `internal/aicontext/skills/payment-integration/references/sepay/overview.md` with a surface-selection map covering API v2, bank-account Webhooks, Payment Gateway/IPN, QR utility, Order VAs, optional products, Test Mode vs Live, and production safety rules.
- `E1-P1A-IMPL2`: Added `internal/aicontext/skills/payment-integration/references/sepay/payment-gateway.md` as the dedicated hosted checkout / Payment Gateway IPN reference.
- `E1-P1A-ROUTE1`: Updated `internal/aicontext/skills/payment-integration/SKILL.md` quick reference to list `references/sepay/payment-gateway.md`.
- `E1-P1A-ROUTE2`: Updated `internal/aicontext/skills/payment-integration/references/implementation-workflows.md` so SePay agents choose the surface first and route hosted checkout/IPN work to `payment-gateway.md`.
- `E1-P1A-SRC1`: Readback/grep confirmed `overview.md` names the separate surfaces and links to `payment-gateway.md`, `api.md`, `webhooks.md`, and `qr-codes.md`.
- `E1-P1A-SRC2`: Readback/grep confirmed routing files include `payment-gateway.md`.

## E2 - P2 Evidence

Matching plan item(s): `P2-A`

- `E2-P2A-IMPL1`: Replaced `internal/aicontext/skills/payment-integration/references/sepay/api.md` with API v2-first guidance.
- `E2-P2A-SRC1`: Readback/grep confirmed `api.md` lists `https://userapi.sepay.vn/v2`, `https://userapi-sandbox.sepay.vn/v2`, Bearer auth, 3 requests/second per IP, `Retry-After`, `data` / `meta.pagination`, UUID identifiers, and HTTP status/error handling.
- `E2-P2A-SRC2`: Readback/grep confirmed legacy `https://my.sepay.vn/userapi/*` and 2 calls/second guidance appear only under `Legacy v1 Migration Notes`, not as the default path.
- `E2-P2A-SRC3`: Readback confirmed Order VA guidance covers BIDV enterprise, Sacombank personal/household-business, Vietcombank enterprise/household-business, exact-amount/partial-payment constraints, and bank-specific request-field warnings.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`

- `E3-P3A-IMPL1`: Replaced `internal/aicontext/skills/payment-integration/references/sepay/webhooks.md` with bank-account webhook guidance.
- `E3-P3A-SRC1`: Readback/grep confirmed HMAC-SHA256 guidance includes raw body, `X-SePay-Signature`, `X-SePay-Timestamp`, `{timestamp}.{raw_body}`, replay window, and timing-safe comparison.
- `E3-P3A-SRC2`: Readback confirmed success response, timeout, automatic retry, manual replay, idempotency, monitoring, dashboard metrics, alerts, incidents, reconciliation, and IP allowlist guidance.
- `E3-P3A-SRC3`: Readback confirmed `webhooks.md` explicitly excludes Payment Gateway IPN and points gateway work to `payment-gateway.md`.
- `E3-P3A-SCOPE1`: `scripts/sepay-webhook-verify.js` remained inspect-only; no script or test file was edited.

## E4 - P4 Evidence

Matching plan item(s): `P4-A`

- `E4-P4A-IMPL1`: Filled `internal/aicontext/skills/payment-integration/references/sepay/payment-gateway.md` with gateway base URLs, Basic Auth, checkout init, IPN URL/acknowledgement, `X-Secret-Key`, payload shape, notification types, idempotency, order APIs, and reconciliation.
- `E4-P4A-IMPL2`: Replaced `internal/aicontext/skills/payment-integration/references/sepay/sdk.md` with current official Payment Gateway SDK guidance.
- `E4-P4A-SRC1`: Readback/grep confirmed `payment-gateway.md` includes `https://pgapi.sepay.vn`, `https://pgapi-sandbox.sepay.vn`, `/v1/checkout/init`, `X-Secret-Key`, `ORDER_PAID`, `TRANSACTION_VOID`, and separate gateway idempotency guidance.
- `E4-P4A-SRC2`: Readback/grep confirmed `sdk.md` uses `npm i sepay-pg-node`, keeps PHP `composer require sepay/sepay-pg`, documents `CARD`, `BANK_TRANSFER`, `NAPAS_BANK_TRANSFER`, field-order signature warnings, order list/detail, card void, and bank/NAPAS cancel behavior.
- `E4-P4A-SRC3`: Readback confirmed Laravel package guidance is not presented as the official Payment Gateway SDK path; it directs agents to inspect package authority before use.

## E5 - P5 Evidence

Matching plan item(s): `P5-A`

- `E5-P5A-IMPL1`: Replaced `internal/aicontext/skills/payment-integration/references/sepay/qr-codes.md` with current QR utility guidance.
- `E5-P5A-IMPL2`: Replaced `internal/aicontext/skills/payment-integration/references/sepay/best-practices.md` with repo-agnostic production patterns.
- `E5-P5A-SRC1`: Readback/grep confirmed QR guidance includes `showinfo`, `fullacc`, `holder`, `store`, `standee`, bank identifier variants, official VA, memo-based VA, no-VA, `TKP`, and VietinBank `SEVQR`.
- `E5-P5A-SRC2`: Readback/grep confirmed best practices cover surface selection, HMAC-first webhooks, `code` usage, idempotency, replay safety, API v2/gateway reconciliation, Test Mode/Live separation, IP allowlisting, and amount policy.
- `E5-P5A-SRC3`: Readback/grep confirmed previous project-specific generic guidance was removed from SePay references: `CLAUDEKIT`, product pricing examples, GitHub fulfillment, coupon/referral, and Polar-specific workflow text no longer appear under `references/sepay`.

## E6 - P6 Evidence

Matching plan item(s): `P6-A`

- `E6-P6A-VAL1`: `git diff --check` from `E:\Anvien` passed with no whitespace errors before full validation.
- `E6-P6A-VAL2`: `npm run full-build` from `E:\Anvien` passed. It ran npm install, Go runtime package build, global install, `anvien version` (`1.2.6`), launcher build, `anvien-web` TypeScript/Vite production build, and completed successfully.
- `E6-P6A-GRAPH1`: `npm run full-build` also ran `anvien analyze . --force`. Result: analyzed `E:\Anvien`, files scanned 1395, parsed code 673, failed 0, indexed documents 485, graph nodes 82455, relationships 120516, graph path `.anvien/graph.json`, stale false.
- `E6-P6A-FD1`: `anvien file-detail internal/aicontext/skills/payment-integration/references/sepay/payment-gateway.md --repo Anvien --json` returned markdown docs, parsed, 0 local/outbound/inbound relationships, 0 linked flows/tests, low risk, stale false, changedSinceAnalyze false.
- `E6-P6A-GEN1`: Generated repo output smoke for `.agents/skills/payment-integration` found `payment-gateway.md`, HMAC-SHA256 guidance, and `pgapi.sepay.vn` references in the generated payment-integration skill output.
- `E6-P6A-GEN2`: Generated repo output smoke for `.claude/skills/payment-integration` found `payment-gateway.md`, HMAC-SHA256 guidance, and `pgapi.sepay.vn` references in the generated payment-integration skill output.
- `E6-P6A-GEN3`: Home-level `C:\Users\TAM NGUYEN\.agents\skills\payment-integration` was not regenerated by this source build; smoke showed stale old SePay content there. This is outside the repo-generated validation target and should be refreshed by `anvien setup` or the relevant skill sync command if the user wants home-level installed skills updated immediately.
- `E6-P6A-DETECT1`: `anvien detect-changes --repo Anvien --scope all` passed. Summary: changed files 12, affected files 12, changed app layer docs, changed functional area documentation, affected processes none, risk level low.
- `E6-P6A-SCOPE1`: Validation grep found stale SePay gateway URLs in `internal/aicontext/skills/payment-integration/scripts/checkout-helper.js`. This plan is docs-only per user correction and scoped to reference/routing docs, so no script code was edited. Record as follow-up/out-of-scope if a later plan expands from docs to helper scripts.

## Closure Evidence

Use this section for supervisor review, dead-work cleanup, final detect-changes, commit hash, and closure evidence when the plan reaches completion.

- `E7-PNA-SUP1`: Supervisor PASS report written at `reports/Supervisor/rp_supervisor_260614_185121_by_gpt-5-codex_sepay-skill-refresh-docs.md`. Verdict: PASS for P1-P6 docs-only scope. The report clears source references/routing, validation evidence, generated repo output smoke, and records `scripts/checkout-helper.js` as out-of-scope follow-up.
- `E8-PNB-CLEAN1`: Dead-work cleanup review found no plan-created temp files or rejected/dead source artifacts to remove. `git status --short` contained only planned source docs, plan ledgers, new `payment-gateway.md`, and the supervisor report. Existing `.tmp` entries predate this plan and were not created by this scope. Grep found no `TODO`, `PLACEHOLDER`, stale status text, or pending benchmark cells in the active source/reference scope; remaining `Classification:` labels are section headings in actual-status, not incomplete values.
- `E8-PNB-SUP1`: The same supervisor PASS report reviewed residual same-invariant surfaces and found no residual unverified same-invariant surfaces for the docs-only scope.
- `E9-PNC-GRAPH1`: Final `anvien analyze --force` after supervisor/closure docs passed. Result: files scanned 1396, parsed code 673, failed 0, indexed documents 486, graph nodes 82463, relationships 120524, graph path `.anvien/graph.json`.
- `E9-PNC-DETECT1`: Final `anvien detect-changes --repo Anvien --scope all` after the final graph refresh passed. Summary: changed files 12, affected files 12, changed app layer docs, changed functional area documentation, affected processes none, risk level low.
- `E9-PNC-COMMIT1`: Final commit executed after recording this evidence; commit hash is reported in the final assistant response because a commit cannot include its own final hash in the same tree.
