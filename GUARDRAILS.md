# Guardrails - Anvien

Rules for **human contributors** and **AI agents**. Complements `AGENTS.md` (workflows) and `CONTRIBUTING.md` (PR process).

## Scope (least privilege)

- **Read:** Source, tests, docs, public config as needed.
- **Write:** Only files required for the fix or feature; no unrelated formatting or refactors.
- **Execute:** Tests, typecheck, documented CLI commands. No destructive commands on user data without approval.
- **Off-limits:** Other people's machines, production deployments you don't own, credentials you lack permission to use.

Maintainer may widen scope per task.

---

## Non-negotiables

1. **Never commit secrets** — API keys, tokens, real `.env` values, private URLs, session cookies. Use `.env.example` with placeholders.
2. **Never rename with find-and-replace** in anvien-indexed projects — use the `rename` MCP tool or `anvien rename` CLI with a dry-run/review path first, then review graph-guided edits and text-search fallbacks.
3. **Run impact analysis before editing shared symbols when graph tools are available** — `impact` (upstream) for functions/classes/methods others call. Do not ignore HIGH/CRITICAL without maintainer sign-off.
4. **Run `detect_changes` before commit when graph tools are available** — confirm diffs map to expected symbols/processes.
5. **Preserve embeddings** — if `.anvien/meta.json` shows embeddings, use `anvien analyze --embeddings`; plain `analyze` drops them.
6. **Keep Anvien local-first** — do not add an Anvien-hosted cloud service, required daemon, required Docker path, or managed workspace requirement to the default runtime.
7. **Keep repo context explicit** — Web/HTTP/MCP/CLI code should pass repo name/path/session explicitly. Do not reintroduce graph reads that depend on one mutable process-global active repo.
8. **Keep the launcher optional** — `anvien serve` must remain the direct local backend entry point.
9. **Keep generated AI context source-owned** — `AGENTS.md`, `CLAUDE.md`, and `.claude/skills/anvien/**` are generated outputs. Update `internal/aicontext/aicontext.go` or `internal/aicontext/skills/*.md`, then regenerate.
10. **Do not generate self-referential AI-context skills** — the generated Anvien skill set is task-facing only. It must not include `anvien-ai-context` or a skill that tells agents to modify its own generator.

---

## Local Runtime Invariants

- `anvien analyze` writes repo-local index data under `<repo>/.anvien/`.
- `anvien mcp` exposes indexed repos over stdio for local agents.
- `anvien serve` exposes the local HTTP backend on loopback for the Web UI.
- `anvien-web/` is a thin client over the local backend; it must not become the owner of graph storage.
- File Map/File Detail relationships are derived projections from symbol and source-site graph facts; the canonical graph remains symbol/source-site based.
- `anvien-launcher/` is a Windows convenience layer around the same backend and Web UI.
- Anvien must not store AI provider API keys in browser storage or route chat through an Anvien cloud proxy.
- Current chat execution is through the Codex CLI adapter. `claude-code` is a shared provider identifier/UI slot until a backend adapter is implemented.

---

## Signs (recurring failure patterns)

Format: **Trigger → Instruction → Reason**. Append new Signs when the same mistake repeats.

### Stale graph after edits

- **Trigger:** MCP warns index is behind `HEAD`, or search doesn't match latest commit.
- **Do:** `anvien analyze` (plus `--embeddings` if used).
- **Why:** Tools query LadybugDB from last analyze; git changes are invisible until re-indexed.

### Embeddings vanished after analyze

- **Trigger:** Semantic search quality drops; `stats.embeddings` in `meta.json` is 0 after refresh.
- **Do:** `anvien analyze --embeddings`, confirm `meta.json` reflects stored embeddings.
- **Why:** Embedding generation is opt-in; analyze without the flag does not preserve prior vectors.

### MCP lists no repos

- **Trigger:** MCP stderr says no indexed repos.
- **Do:** `anvien analyze` in the target repo; verify `anvien list` shows it.
- **Why:** MCP discovers repos via `~/.anvien/registry.json`, populated by analyze.

### Wrong repo in multi-repo setups

- **Trigger:** Query/impact results belong to another project.
- **Do:** Call `list_repos`, then pass `repo` on subsequent tools.
- **Why:** Default target is ambiguous when multiple repos are registered.

### Web repo switch regresses

- **Trigger:** Switching repos in the Web UI hangs, falls back to the previous repo, or graph nodes/links do not render.
- **Do:** Verify `anvien serve` directly, then switch repo A -> B -> A with at least two indexed repos. Check `/api/graph?repo=...&stream=true`.
- **Why:** Web graph loading must stay repo-scoped. Reintroducing ambient active-repo state can make one repo's runtime affect another repo.

### Graph selection paths drift

- **Trigger:** Clicking a graph node works, but clicking a dashboard file or search result does not show the expected graph/link context.
- **Do:** Reuse the same visible graph selection path for graph nodes, dashboard files, and search results.
- **Why:** Separate selection paths can hide graph links or apply stale/invisible filters.

### LadybugDB lock / "database busy"

- **Trigger:** Errors opening `.anvien/lbug`, active job lock errors, or WAL/checksum errors.
- **Do:** Stop launcher/backend/MCP sessions that may hold the repo, wait for analyze/embed/delete jobs to finish, then retry or rebuild with `anvien analyze . --force`.
- **Why:** Analyze/embed/delete are repo write paths. Overlapping writers or killing a writer mid-flight can corrupt the repo-local LadybugDB store.

### Generated agent context drift

- **Trigger:** `AGENTS.md`, `CLAUDE.md`, or `.claude/skills/anvien/**` disagrees with `internal/aicontext/` source, or a generated skill tries to own AI-context generator behavior.
- **Do:** Fix `internal/aicontext/aicontext.go` or `internal/aicontext/skills/*.md`, regenerate via analyze/setup, and keep `anvien-ai-context` out of generated skills.
- **Why:** Generated root context and skills are outputs. A self-referential AI-context skill makes generated guidance the source of truth for its own generator, which reverses ownership.

### Launcher behavior diverges from `anvien serve`

- **Trigger:** Packaged launcher works differently from direct `anvien serve`, or reset/stop leaves stale runtime processes.
- **Do:** Build with `anvien-launcher\build.ps1` and smoke-test start, reset, stop, and protocol registration.
- **Why:** The launcher is only a convenience wrapper. It must not change backend semantics.

### Chat settings imply cloud/provider config

- **Trigger:** UI copy suggests Anvien stores API keys, configures model temperature/tokens for a cloud proxy, or fully supports Claude Code chat execution.
- **Do:** Describe the current local session bridge accurately: Codex CLI adapter implemented; `claude-code` reserved until its backend adapter exists.
- **Why:** Misleading provider UI can make users think Anvien is an AI provider management layer instead of a local code-intelligence runtime.

---

## Validation Guardrails

- Docs-only changes: `git diff --check` is usually enough.
- Generated contract changes: run the Go contract generator/tests, then build affected CLI/Web packages.
- CLI/MCP/backend/LadybugDB changes: build, typecheck, and run relevant `anvien/` tests.
- Web graph/repo switching changes: run Web build/tests and manually verify repo A -> B -> A through `anvien serve`.
- Launcher changes: run `anvien-launcher\build.ps1` and smoke-test packaged start/reset/stop.
- Do not run full suites by habit when the change is docs-only or narrowly scoped; broaden validation when the touched code crosses runtime boundaries.

---

## Publishing & supply chain

- **npm:** Do not publish from unreviewed automation. Bump version intentionally; tag releases to match `package.json`.
- **Dependencies:** Minimal, auditable `package.json` changes; run tests and CI after lockfile updates.
- **License:** PolyForm Noncommercial 1.0.0 — do not relicense without maintainer approval.

---

## Escalation

Stop and ask a **human maintainer** when:

- Impact analysis shows HIGH/CRITICAL risk and the task still requires the change.
- You need to alter CI, release, or security-sensitive config.
- Requirements conflict (e.g. "speed up analyze" vs "must keep all embeddings on huge repo").
- You are unsure whether data loss is acceptable (`clean`, forced migrations, schema changes).

---

## Related docs

- [ARCHITECTURE.md](ARCHITECTURE.md) — components and data flow
- [RUNBOOK.md](RUNBOOK.md) — commands for recovery
- [CONTRIBUTING.md](CONTRIBUTING.md) — PR and commit expectations
