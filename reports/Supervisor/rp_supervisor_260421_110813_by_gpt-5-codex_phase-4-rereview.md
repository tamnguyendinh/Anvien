# Phase 4 Re-review

- Plan: `docs/plans/2026-04-20-convert-all-to-local.md`
- Scope reviewed: Phase 4 only (`Hardening local-only`)
- Verdict: `APPROVED`
- Reviewer: `gpt-5-codex`
- Timestamp: `2026-04-21 11:08:13 +07:00`

## Why Phase 4 now passes

The two prior Phase 4 blockers are closed.

1. `setup` no longer writes or suggests the remote npm-latest MCP path.
   - `anvien/src/cli/setup.ts:28`
   - `anvien/src/cli/setup.ts:38`
   - `anvien/src/cli/setup.ts:252`
   - `getMcpEntry()` is now canonically `anvien mcp`.
   - `anvien/test/integration/setup-skills.test.ts:87`
   - `anvien/test/integration/setup-skills.test.ts:97`
   - Behavioral test now asserts Codex config contains `command = "anvien"`, contains `"mcp"`, and does **not** contain `anvien@latest`.

2. Active onboarding/local-only web surfaces now pass their Phase 4 behavioral checks.
   - `anvien-web/test/unit/OnboardingGuide.local-only.test.tsx:1`
   - `anvien-web/test/unit/useBackend.local-only.test.tsx:1`
   - `anvien-web/test/unit/SettingsPanel.local-runtime.test.tsx:1`
   - `anvien-web/test/unit/SettingsPanel.compat-local-runtime.test.tsx:1`
   - `anvien-web/test/unit/server-connection.test.ts:1`

3. Backend hardening expected by Phase 4 is covered and green.
   - `anvien/test/unit/cors.test.ts:1`
   - `anvien/test/unit/serve-command.test.ts:1`

## Plan-bound interpretation

This rereview is constrained to the actual Phase 4 acceptance in the plan:

- remove `npx -y anvien@latest` fallback from `setup`
- harden backend/CORS to local-only
- force web backend usage back to local-only
- remove remote/API-key/cloud wording from active `onboarding/settings/help` surfaces
- keep wording left in legacy compatibility files for Phase 6 cleanup

I re-checked this boundary and did **not** use unrelated repo-wide noise as a blocker.

## Residual observations that do not block Phase 4

1. `anvien/src/cli/analyze.ts:373`
2. `anvien/src/cli/analyze.ts:374`
3. `anvien/src/cli/analyze.ts:383`
4. `anvien/src/cli/analyze.ts:384`

`analyze.ts` still has troubleshooting copy that mentions `npm install -g anvien@latest` / `npx anvien@latest analyze`. I am **not** treating that as a Phase 4 blocker because the plan narrows this phase to `setup`, `serve`, local backend hardening, and active `onboarding/settings/help` surfaces. This copy should be reviewed in a later CLI/help cleanup pass if the team wants full wording convergence.

## Validation run

- `cd anvien && npx tsc --noEmit`
- `cd anvien-web && npx tsc -b --noEmit`
- `cd anvien && npx vitest run test/integration/setup-skills.test.ts test/unit/setup-session-runtime.test.ts test/unit/tools.test.ts`
- `cd anvien && npx vitest run test/unit/cors.test.ts test/unit/serve-command.test.ts`
- `cd anvien-web && npx vitest run test/unit/useBackend.local-only.test.tsx test/unit/OnboardingGuide.local-only.test.tsx test/unit/SettingsPanel.local-runtime.test.tsx test/unit/SettingsPanel.compat-local-runtime.test.tsx test/unit/package-deps.local-runtime.test.ts`
- `cd anvien-web && npx vitest run test/unit/server-connection.test.ts`

Results:

- `anvien` typecheck: pass
- `anvien-web` typecheck: pass
- `setup-skills` / `tools` targeted tests: pass
- `cors` / `serve-command` targeted tests: pass
- web Phase 4 targeted tests: pass

## Scope hygiene

Current worktree also contains unrelated uncommitted changes outside this Phase 4 review:

- `anvien-web/src/hooks/useAppState.local-runtime.tsx`
- `anvien-web/test/unit/useAppState.local-runtime.test.tsx`
- `anvien/src/core/lbug/lbug-adapter.ts`

They are not used to determine this Phase 4 verdict.
