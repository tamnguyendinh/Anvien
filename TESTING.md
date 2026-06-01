# Testing — Anvien

How we structure tests and which commands to run locally and in CI.

## Packages

| Package | Path | Runner | Notes |
| ------- | ---- | ------ | ----- |
| Root tooling | `/` | Prettier / ESLint / lint-staged | Formatting and repository automation checks. No `.husky/` hook is currently committed. |
| CLI + MCP + HTTP backend | `cmd/`, `internal/`, `anvien/` | Go test + package build | Primary core/runtime test surface. Includes ingestion, MCP, LadybugDB, repo runtime, file projection, API contracts, embeddings, and server helpers. |
| Web UI unit/component | `anvien-web/` | Vitest + jsdom | Graph UI, File Map, File Detail, repo picker/analyze flows, local runtime settings, chat UI/client behavior. |
| Web UI E2E | `anvien-web/` | Playwright | Browser flows against real `anvien serve` and Vite dev server. |
| Windows launcher | `anvien-launcher/` | Build/manual smoke | Build script produces launcher exe, backend wrapper, bundled web dist, and protocol registration. |

## Commands (local)

From repository root, unless noted:

**Root tooling**

```bash
npm install
npm run format:check
npm run lint
```

**`anvien` (CLI / library)**

```bash
go test ./cmd/... ./internal/... -count=1

cd anvien
npm install
npm run build               # package runtime build
npm test                    # reminder: Go tests own the runtime now
```

**File projection / contract focus**

```bash
go test ./internal/filecontext ./internal/cli ./internal/httpapi ./internal/mcp -count=1
go run ./cmd/generate-web-contracts --check
```

**`anvien-web`**

```bash
cd anvien-web
npm install
npm run build
npm test                    # unit tests (vitest)
npx tsc -b --noEmit         # typecheck (matches CI)
npm run test:coverage
```

**Web E2E**

Run these in separate terminals:

```bash
go run ./cmd/anvien serve --host 127.0.0.1 --port 4848
```

```bash
cd anvien-web
npm run dev
```

Then:

```bash
cd anvien-web
npm run test:e2e
```

Playwright ignores `manual-record.spec.ts` and `debug-issues.spec.ts` by default. Use `PLAYWRIGHT_INSECURE=1` only for explicit browser-security experiments.

**Windows launcher package**

```powershell
powershell -ExecutionPolicy Bypass -File anvien-launcher\build.ps1
```

Use this when launcher code, bundled backend/web output, protocol behavior, or user-facing packaged runtime behavior changed.
The launcher package must build the Go backend binary and must not depend on a bundled Node backend runtime.

## What To Run

Run the smallest useful validation for the change.

| Change area | Minimum useful validation |
| ----------- | ------------------------- |
| Docs only | `git diff --check` |
| Docs describing current behavior | Source or Anvien evidence for the claim, then `git diff --check` |
| CLI command, MCP tool, graph query, ingestion, LadybugDB | `go test ./cmd/... ./internal/... -count=1`, plus package build when command packaging changed. |
| File projection, File Map/File Detail API, file-aware MCP/CLI | `go test ./internal/filecontext ./internal/cli ./internal/httpapi ./internal/mcp -count=1`, contract check, and Web tests if browser behavior changed. |
| Narrow core logic | Targeted Go package test, for example `go test ./internal/<pkg> -run <TestName> -count=1`; broaden if storage/MCP/API wiring is involved. |
| Web UI component/state only | `cd anvien-web && npm run build && npx tsc -b --noEmit && npm test` |
| Repo switching, graph loading, analyze from Web UI | Web build/tests plus manual or Playwright E2E against `anvien serve`. |
| Session chat runtime | Web unit tests plus manual `/api/session/status` and one chat request with the local Codex CLI session. |
| Launcher | Full launcher build plus start/reset/stop smoke check. |

Avoid running the full test matrix for docs-only or copy-only changes. Prefer targeted validation first, then broaden when the touched code crosses runtime boundaries.

## Pre-commit hook

This repository currently does not contain a committed `.husky/pre-commit` hook.
The root package includes Husky and lint-staged dependencies for optional local
hook work, but hooks are not the validation source of truth.

If a hook is added later, the intended behavior for the current package layout is:

1. **Formatting** — `lint-staged` runs prettier on staged files
2. **`anvien-web/` files staged** → `tsc -b --noEmit`
3. **Go runtime files staged** → run targeted Go validation before commit

Maintenance note: if `.husky/pre-commit` is edited, keep it aligned with the current package paths: `anvien/`, `anvien-web/`, and Go runtime paths under `cmd/` and `internal/`.

## Test categories

- **Unit** — Pure logic, parsers, graph/query helpers; fast; no network.
- **Integration** — Real combinations such as filesystem, MCP wiring, HTTP handlers, graph load/query, and larger pipelines as Go package tests under `internal/**` and browser E2E under `anvien-web/e2e/`.
- **Eval-style / golden sets** — For agent- or classification-style behavior, keep labeled inputs and expected outputs (JSON or table-driven tests) and run them in CI when relevant.
- **E2E (web)** — Critical user paths only; prefer `data-testid` attributes for stable selectors. Tests run against the real Go backend (`anvien serve`) and Vite dev server.
- **Manual smoke** — Required for packaged launcher behavior, OS folder picker behavior, and any browser flow that depends on real local machine state.

## Performance metrics (targets)

Set targets to match team expectations, then tune to this repo’s CI reality:

| Metric | Target (initial) | Notes |
| ------ | ---------------- | ----- |
| Unit coverage | Align with CI | Use Go package coverage where useful and Web Vitest coverage when browser behavior changes. |
| Unit wall time | Fast PR feedback | Use targeted Go package tests or focused Web Vitest files for tight loops. |
| Integration duration | &lt; few minutes | Guard heavy tests with env flags if needed. |
| Web graph interaction | No visible UI stall after graph load | Manually verify when graph loading, selection, or filtering changes. |
| Repo switching | Stable across repeated repo A -> B -> A switches | Required for backend/Web repo-runtime changes. |

## Regression testing

Re-run the full relevant suite when:

- Prompt or agent-behavior documentation changes (if tests encode behavior)
- Model or embedding-related code paths change
- Graph schema, query contracts, or MCP tool shapes change
- File projection, file-aware command output, File Map, or File Detail changes
- Dependencies with parsing or runtime impact upgrade
- Repo registry, repo resolver, graph streaming, or Web repo switching changes
- Session bridge, Codex adapter, or chat cancellation changes
- Launcher process management, reset, stop, or protocol registration changes

Manual regression matrix for local runtime changes:

1. Index at least two local repos with `anvien analyze .`.
2. Start `anvien serve` directly, without the launcher.
3. Start `anvien-web` and switch repo A -> B -> A from the dropdown.
4. Confirm graph loads through `/api/graph?repo=...&stream=true` and the UI does not fall back to the previous repo after a successful load.
5. Open File Map and File Detail; confirm `/api/file-hotspots` and `/api/file-context` return file projection data for the selected repo.
6. Click a graph node, a dashboard file, and a search result; each should use the same visible graph selection path.
7. Start analyze from the Web UI and confirm the repo list/dropdown refreshes after success.
8. If chat changed, confirm `/api/session/status` reflects the local Codex CLI session and no Anvien API key is required.
9. If launcher changed, build the launcher and smoke-test start, reset, stop, and protocol registration.

## CI integration

GitHub Actions (`.github/workflows/ci.yml`) orchestrate:

- **`ci-quality.yml`** — prettier format check, eslint lint, package typecheck, Web typecheck, and workflow-convention checks
- **`ci-tests.yml`** — package/Web tests, coverage artifacts where configured, Docker nginx validation, and cross-platform package sanity
- **`ci-e2e.yml`** — Playwright E2E tests, gated on `anvien-web/**` changes

Local checks before pushing:

```bash
go test ./cmd/... ./internal/... -count=1
cd anvien && npm run build && npm test
cd ../anvien-web && npx tsc -b --noEmit && npm test
```

The pre-commit hook is useful for formatting/typecheck feedback, but do not treat it as a replacement for package-specific validation.

## User acceptance / beta (optional)

For staged releases or UI betas, use the packaged local runtime rather than a remote staging service:

1. Build with `anvien-launcher\build.ps1`.
2. Run `.\anvien-launcher\AnvienLauncher.exe`.
3. Confirm the browser opens the exe-served Web UI start screen.
4. Confirm the old root start HTML launcher file is absent from repository root, `anvien-web\dist\`, and `anvien-launcher\web-dist\`.
5. Verify backend health at `http://127.0.0.1:4848/api/info`.
6. Verify start screen, Back to start screen, repo picking, repo switching, graph selection, analyze, reset runtime, and chat status on the target Windows machine.
7. Collect runtime logs from `anvien-launcher/logs/` when diagnosing failures.
