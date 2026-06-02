# Anvien Deadcode Agent Token Benchmark Evidence Ledger

Date: 2026-06-02

Status: complete

Companion files:

- Plan: [2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md](2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md)
- Benchmark ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md](2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md)

## Reset Notice

All prior evidence from the invalid run is discarded.

Reason: the old run did not correctly measure AI-agent token spend in Anvien mode. Anvien local output volume is not the same thing as agent token usage.

Do not reuse old command logs, old candidate lists, old token totals, old verification tables, or old conclusions.

## Evidence Rules

1. Evidence explains why the benchmark result is valid.
2. Keep quantitative metric tables in the benchmark ledger.
3. Native mode evidence must not include Anvien commands, outputs, resources, generated context, graph artifacts, or prior Anvien reports.
4. Anvien mode evidence must record graph freshness before graph-based commands.
5. Token evidence must prove the AI agent's token usage measurement mechanism, not Anvien output size.
6. Record delivered tool results and file reads only when they enter the AI agent session or exact transcript proxy.
7. Record unobserved local tool artifacts as unobserved; do not turn them into token usage.
8. Record failures, retries, and abandoned paths because they can consume agent tokens when delivered to the agent.
9. Record candidate verification from source-backed facts, not intuition.
10. Do not record deadcode deletion or cleanup patches because this benchmark only finds candidates.

## Evidence Template

Use this template for each phase:

```text
## E<n> - <Phase>

Date:

Status:

Scope:

- ...

Commands / reads:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| ... | ... | ... | ... | ... | ... |

Candidate evidence:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| ... | ... | ... | ... | ... | ... |

Failures / retries:

- ...

Completion:

| Item | Result |
|---|---|
| Declared procedure recorded before discovery | pending |
| Token measurement active | pending |
| Completion condition met | pending |
| Open leads remaining | pending |
| Blocker or incomplete reason | pending |
| Confidence | pending |
```

## E0 - Baseline

Date: 2026-06-02T13:28:02.5794035+07:00

Status: complete

Required evidence:

| Check | Result |
|---|---|
| Baseline commit | `6516020c323b54c74583fbaa2caf81dad4475036` |
| Branch | `master` |
| Worktree status | Clean before final baseline evidence was written; `git status --porcelain=v1` produced no output. |
| Source-code dirty state | Clean. |
| Benchmark docs/reports dirty state | Clean before final baseline evidence was written. |
| Shell | PowerShell 7.6.2 |
| OS / CPU / RAM | Microsoft Windows 10 Pro 10.0.19045 64-bit; Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz; 4 cores / 8 logical processors; 32,459,440 KiB visible memory; 16,246,240 KiB free physical memory. |
| Go / Node / npm / Git versions if used | `go version go1.26.3 windows/amd64`; `node v24.15.0`; `npm 11.12.1`; `git version 2.54.0.windows.1`; `codex-cli 0.136.0`; `tiktoken 0.13.0` |

Commands / reads:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| E0-1 | `git rev-parse HEAD` | Record baseline commit. | yes | Final benchmark baseline evidence; discovery not started. | `6516020c323b54c74583fbaa2caf81dad4475036` |
| E0-2 | `git branch --show-current` | Record branch. | yes | Baseline evidence only; discovery not started. | `master` |
| E0-3 | `git status --porcelain=v1` | Record dirty state. | yes | Final benchmark baseline evidence; discovery not started. | Clean; no output. |
| E0-4 | `git --version` | Record tool version. | yes | Baseline evidence only; discovery not started. | `git version 2.54.0.windows.1` |
| E0-5 | `go version` | Record tool version. | yes | Baseline evidence only; discovery not started. | `go version go1.26.3 windows/amd64` |
| E0-6 | `node --version` | Record tool version. | yes | Baseline evidence only; discovery not started. | `v24.15.0` |
| E0-7 | `npm --version` | Record tool version. | yes | Baseline evidence only; discovery not started. | `11.12.1` |
| E0-8 | `$PSVersionTable.PSVersion.ToString()` | Record shell version. | yes | Baseline evidence only; discovery not started. | `7.6.2` |
| E0-9 | `Get-Date -Format o` | Record local timestamp. | yes | Baseline evidence only; discovery not started. | `2026-06-02T13:28:02.5794035+07:00` |
| E0-10 | `Get-CimInstance Win32_OperatingSystem ...` | Record OS/RAM. | yes | Baseline evidence only; discovery not started. | Windows 10 Pro 10.0.19045 64-bit; 32,459,440 KiB visible memory; 16,246,240 KiB free physical memory. |
| E0-11 | `Get-CimInstance Win32_Processor ...` | Record CPU. | yes | Baseline evidence only; discovery not started. | Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz; 4 cores / 8 logical processors. |
| E0-12 | `codex --version` | Record agent runtime version. | yes | Baseline evidence only; discovery not started. | `codex-cli 0.136.0` |
| E0-13 | Python `import tiktoken` | Record fallback tokenizer version. | yes | Baseline evidence only; discovery not started. | `tiktoken 0.13.0` |

## E1 - Token Measurement Setup

Date: 2026-06-02

Status: complete

Required evidence:

| Item | Result |
|---|---|
| Measurement mechanism | Primary: isolated `codex exec --json` sessions, using `turn.completed.usage.input_tokens + turn.completed.usage.output_tokens` as `agent_session_tokens`. Fallback/audit: `scripts/measure-agent-token-proxy.ps1` transcript proxy with Python `tiktoken` using `o200k_base` unless overridden. |
| Exact model/runtime token telemetry available | Yes for Codex CLI sessions. Probe output included `usage.input_tokens`, `usage.cached_input_tokens`, `usage.output_tokens`, and `usage.reasoning_output_tokens`. |
| Exact transcript proxy available if telemetry is not available | Yes as a fallback for harnessed commands/files/responses: `scripts/measure-agent-token-proxy.ps1` records NDJSON events and counts delivered content through `tiktoken`. It is not used to infer hidden model telemetry. |
| Can measure native mode with this mechanism | Yes, by running native discovery in its own `codex exec --json` session and parsing the usage event. |
| Can measure Anvien mode with this mechanism | Yes, by running Anvien-guided discovery in a separate `codex exec --json` session and parsing the usage event. |
| Can distinguish agent-session tokens from Anvien local output volume | Yes. Codex CLI usage is runtime model usage, and the proxy script marks undelivered/local output separately from delivered transcript content. |
| Can exclude hidden graph/cache/index/output not delivered to agent | Yes. Hidden local artifacts are not part of Codex CLI model usage unless the measured agent reads/receives them; the proxy script records undelivered output under `local_tool_output_volume`. |
| Blocker if token measurement is unavailable | No current blocker. Discovery has not started because the harness must be committed and the final benchmark baseline recaptured first. |

Commands / reads:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| E1-1 | `get_goal` tool | Check whether the active runtime exposes goal/token usage telemetry. | yes | This confirmed no usable runtime token telemetry for the benchmark phases. | `goal: null`; `remainingTokens: null`; `completionBudgetReport: null` |
| E1-2 | `codex exec --json --cd E:\Anvien --sandbox danger-full-access "Reply with exactly: TOKEN_PROBE_OK"` | Check isolated Codex CLI usage telemetry. | yes | Primary benchmark mechanism; measured usage belongs to the isolated Codex CLI session. | `input_tokens=17510`, `cached_input_tokens=3456`, `output_tokens=26`, `reasoning_output_tokens=16`. |
| E1-3 | `python -m pip install --target .tmp\tokenizer-python tiktoken` | Install tokenizer for transcript proxy fallback. | yes | Enables exact proxy counts for delivered transcript text. | Installed `tiktoken 0.13.0` plus dependencies under `.tmp\tokenizer-python`. |
| E1-4 | `scripts\measure-agent-token-proxy.ps1` smoke sequence | Verify transcript proxy with `tiktoken`. | yes | Fallback/audit mechanism produced exact token counts. | Smoke summary reported `agent_session_token_proxy=56` and `token_measurement_valid=true`. |

Completion:

| Item | Result |
|---|---|
| Declared procedure recorded before discovery | yes |
| Token measurement active | yes, for future isolated Codex CLI discovery sessions |
| Completion condition met | yes |
| Open leads remaining | not applicable; discovery did not start |
| Blocker or incomplete reason | none for token measurement; final benchmark baseline still pending after harness commit |
| Confidence | high |

## E2 - Native Discovery Without Anvien

Date: 2026-06-02

Status: complete

Rules for this section:

- Do not use Anvien.
- Record every delivered command result and source file read that enters the agent session/proxy.
- Record candidates before any Anvien-mode work starts.

Native command/read log:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| N-1 | Isolated `codex exec --json` native session | Run native discovery with Anvien surfaces prohibited by mode envelope. | yes, inside measured Codex CLI session | Runtime telemetry source for native `agent_session_tokens`. | Completed with `input_tokens=850171`, `output_tokens=6819`, `agent_session_tokens=856990`. |
| N-2 | Native session command log | Inventory, static scans, `knip`, `deadcode`, and targeted `rg` / `Get-Content` checks. | yes, inside measured Codex CLI session | Command output is included in Codex CLI runtime usage. | 39 completed command executions; 9 nonzero exits from failed tools/no-match searches. |
| N-3 | Targeted source/config reads | Inspect leads after native searches. | yes, inside measured Codex CLI session | File content is included in Codex CLI runtime usage. | 10 targeted files read, 55,619 bytes. |
| N-4 | Native Anvien-use audit | Check completed command strings for Anvien executable/graph usage. | yes | Audit evidence; not part of native discovery cost. | No completed command invoked `anvien`, `anvien.exe`, `anvien\bin`, or read `.anvien`; several paths contain the product name as repo/package text. |

Native candidates:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| N1 | `anvien-web/src/vendor/leiden/` | module/file group | frontend vendored module | `knip --production` reported `index.d.ts`, `index.js`, and `utils.js` unused; `rg` found no imports outside vendored files. | Low/medium risk: may be intentionally parked vendored code. |
| N2 | `anvien-web/package.json` | `graphology-indices`, `graphology-utils`, `mnemonist`, `pandemonium` | dependencies | Imports found only inside unused `src/vendor/leiden/*`; `knip --production` reported unused dependencies. | Tied to N1. |
| N3 | `anvien-web/package.json` | `@sigma/edge-curve`, `axios`, `d3`, `graphology-layout-force`, `graphology-layout-forceatlas2`, `graphology-layout-noverlap`, `lru-cache`, `react-zoom-pan-pinch`, `uuid`, `zod` | dependencies | `knip --production` reported unused; native `rg` found no source imports outside package/lockfile for checked set. | Medium dependency-indirect-use risk. |
| N4 | `anvien-web/src/components/ToolCallCard.tsx` | `default` export | frontend export | `knip --production` reported duplicate export; `rg` showed named imports of `ToolCallCard`, not default imports. | Component is live; default export appears dead. |
| N5 | `anvien-web/src/config/ui-constants.ts` | `DEFAULT_OLLAMA_BASE_URL`, `DEFAULT_OPENROUTER_BASE_URL`, `REQUIRED_NODE_VERSION` | frontend constants | `rg` found `DEFAULT_BACKEND_URL` usage but no imports for these three constants. | `REQUIRED_NODE_VERSION` has stale-intent risk because Vite defines `__REQUIRED_NODE_VERSION__`. |
| N6 | `anvien-web/src/core/llm/settings-service-local-runtime.ts` | legacy provider setting helpers | frontend API helpers | Native `rg` over `anvien-web/src` matched declarations only; tests reference several helpers. | Medium legacy/migration risk. |
| N7 | `anvien-web/src/services/backend-client.ts` | `fetchServerInfo`, `grep`, `fetchProcesses`, `fetchProcessDetail`, `fetchClusters`, `fetchClusterDetail`, `getAnalyzeStatus`, `getEmbedStatus`, `cancelEmbeddings` | frontend API client helpers | Native `rg` over app/test found no call sites. | Medium risk: reserved wrappers for backend HTTP surfaces. |
| N8 | `internal/cli/lazy_action.go` | `createLazyAction` | Go helper | `go run golang.org/x/tools/cmd/deadcode@latest ./cmd/... ./internal/...` reported unreachable; `rg` showed declaration and tests only. | Low risk. |
| N9 | `internal/cobol/copy_expander.go` | `CopyReplacing`, `parseReplacingClause`, `replacementScanner`, `isReplacementWordChar` | Go COBOL helper/parser | Go deadcode reported unreachable; `rg` showed declaration and tests only. | Medium planned-feature risk. |
| N10 | `internal/cobol/cobol.go` | `extractProgram` | Go helper | Go deadcode reported unreachable; `Apply` uses `extractPrograms`; `rg` found declaration and legacy tests. | Low/medium legacy wrapper risk. |
| N11 | `internal/communities/enrichment.go` | LLM enrichment API and helpers | Go internal exported/private API | Go deadcode reported unreachable; `rg` found declarations/tests only; current community detection uses heuristic labels. | Medium planned-capability risk. |
| N12 | `internal/communities/communities.go` | `CommunityColors`, `CommunityColor` | Go constants/helper | Go deadcode reported `CommunityColor` unreachable; `rg` found declarations/tests only. | Low/medium stale UI-color helper risk. |

Native completion:

| Item | Result |
|---|---|
| Declared native procedure recorded before first search | yes |
| Token measurement closed native phase | yes |
| Completion condition met | yes |
| Open native leads remaining | none reported |
| Blocker or incomplete reason | none |
| Confidence | valid native discovery; no Anvien command/resource/graph use found in completed command audit |

## E3 - Native Discovery Report

Date: 2026-06-02

Status: complete

Required evidence:

| Item | Result |
|---|---|
| Native candidate count | 12 |
| Native unique files read | 10 targeted source/config files |
| Native command/search count | 39 completed command executions |
| Native `agent_session_tokens` or exact proxy tokens | `agent_session_tokens=856990` from Codex CLI runtime usage |
| Native token validity | valid |
| Native completion status | complete |
| Native unresolved questions | dependency indirect-use risk; generated/contract/public API risk for frontend/backend helper exports; planned-capability risk for COBOL and communities helpers |

## E4 - Anvien-Guided Discovery

Date: 2026-06-02

Status: complete

Rules for this section:

- Record graph freshness before graph-based work.
- Do not seed discovery from the native candidate list.
- Record delivered Anvien command output separately from Anvien local output volume.
- Do not count Anvien graph/cache/index/output unless delivered into the agent session/proxy.
- Record source reads after Anvien narrows candidate leads.

Graph freshness:

| Check | Result |
|---|---|
| Analyze command | `E:\Anvien\anvien\bin\anvien.exe analyze --force` from clean worktree `E:\Anvien-benchmark-anvien` |
| Analyze local-output token status | local tool work; not agent tokens unless delivered |
| Delivered analyze output to agent | yes, inside measured Codex CLI session; included in runtime usage |
| Indexed commit | clean worktree baseline `6516020c323b54c74583fbaa2caf81dad4475036` |
| Current commit | `6516020c323b54c74583fbaa2caf81dad4475036` in `E:\Anvien-benchmark-anvien` |
| Fresh/stale result | graph refreshed successfully before graph-based discovery |

Anvien command/read log:

| Step | Command or file | Purpose | Delivered to agent? | Token-measurement note | Result |
|---|---|---|---|---|---|
| A-0 | Failed initial isolated Anvien-mode `codex exec` attempt | Start Anvien discovery after native report. | yes | Retry/error evidence only; no discovery work started and no candidates produced. | Codex CLI usage limit error; retry window after 13:50 local. |
| A-1 | Clean worktree `E:\Anvien-benchmark-anvien` | Avoid native-result contamination from current docs. | yes | Worktree setup outside Anvien discovery token session. | Detached at baseline commit `6516020c323b54c74583fbaa2caf81dad4475036`. |
| A-2 | Isolated `codex exec --json` Anvien session | Run Anvien-guided discovery. | yes, inside measured Codex CLI session | Runtime telemetry source for Anvien `agent_session_tokens`. | Completed with `input_tokens=1519088`, `output_tokens=5905`, `agent_session_tokens=1524993`. |
| A-3 | Anvien command log | Graph refresh, query, graph-health, context commands before targeted source checks. | yes, inside measured Codex CLI session | Anvien command output delivered to agent is included in Codex CLI runtime usage; hidden local graph/cache volume is not separately counted. | 41 completed commands total; 30 Anvien commands; 6 nonzero failed/retry commands. |
| A-4 | Targeted source/config reads | Inspect graph-surfaced leads and dynamic/public risk. | yes, inside measured Codex CLI session | File content is included in runtime usage. | 11 targeted repo files read, 69,934 bytes. |
| A-5 | Native-contamination audit | Search Anvien-mode event log for native-report paths/text. | yes | Audit evidence; not part of Anvien discovery cost. | No event log match for `E:\Anvien\.tmp\deadcode-agent-token-benchmark`, native candidate list text, or native reports. |

Anvien candidates:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| A1 | `internal/httpapi/phase_timer.go` | `phaseTimer`, `newPhaseTimer`, related methods | private Go helper suite | Anvien context showed only test call relationships; exact `rg` found `newPhaseTimer` only in `phase_timer_test.go`. | Low risk; package/file imports exist, so verdict depends on symbol references. |
| A2 | `internal/parser/ast_cache.go` | `ASTCache`, `NewASTCache`, `CachedTree`, methods | Go cache type/helper file | Anvien context showed test calls through `ast_cache_test.go`; exact lookup found `NewASTCache` only in tests and `ASTCache` in declarations/method receivers. | Medium risk: exported under `internal`. |
| A3 | `internal/providers/golang/definitions.go` | `goTypeSpecLabel` | unexported Go helper | Anvien context empty incoming relationships; exact lookup found only function definition; active code calls `goTypeSpecLabelForKind`. | Low risk. |
| A4 | `internal/scopeir/kinds.go` | `CallIndex` | exported Go enum constant | Anvien context empty incoming relationships; exact `rg` found only declaration. | Medium risk: exported enum/string value may be used dynamically. |
| A5 | `anvien-web/src/core/llm/types.local-runtime.ts` | `AgentStep` | exported TypeScript interface | Anvien showed no direct incoming uses; exact lookup found only declaration. | Medium risk: exported type may be compatibility surface. |
| A6 | `anvien-web/src/core/llm/types.local-runtime.ts` | `GRAPH_SCHEMA_DESCRIPTION` | exported TypeScript constant | Exact lookup found declaration plus test-only import/assertion; no production consumer. | High risk: inline comment says retained to preserve old export surface. |

Anvien completion:

| Item | Result |
|---|---|
| Declared Anvien procedure recorded before first graph command | yes |
| Token measurement closed Anvien phase | yes |
| Completion condition met | yes |
| Open Anvien leads remaining | none reported |
| Blocker or incomplete reason | initial retry error from Codex CLI usage limit before work started; successful retry completed |
| Confidence | valid Anvien discovery; graph refreshed first; clean worktree avoided native candidate contamination |

## E5 - Anvien Discovery Report

Date: 2026-06-02

Status: complete

Required evidence:

| Item | Result |
|---|---|
| Anvien candidate count | 6 |
| Anvien unique files read | 11 targeted repo files |
| Anvien command count | 30 Anvien commands inside 41 completed total commands |
| Anvien follow-up native search count | included in 11 non-Anvien completed commands |
| Anvien `agent_session_tokens` or exact proxy tokens | `agent_session_tokens=1524993` from Codex CLI runtime usage |
| Anvien token validity | valid |
| Anvien completion status | complete |
| Anvien unresolved questions | exported internal/type compatibility risk for `ASTCache`, `CallIndex`, `AgentStep`, and explicitly retained `GRAPH_SCHEMA_DESCRIPTION` |

## E6 - Candidate Union And Verification

Date: 2026-06-02

Status: complete

Verification rules:

- Verify the union of candidates from both modes.
- Check static references, dynamic/public entrypoint risk, generated-code status, test/build/runtime hooks, and external contract hints.
- Do not delete or edit candidate code.

Candidate verdicts:

| Candidate id | Found by native | Found by Anvien | Path | Symbol/name | Verdict | Verification evidence | Dynamic/public risk |
|---|---|---|---|---|---|---|---|
| D01 | yes | no | `anvien-web/src/vendor/leiden/` | module/file group | `confirmed_deadcode` | `rg` found no `vendor/leiden` imports outside package/lockfile and vendor files; candidate is not wired into app runtime. | low/medium parked-vendor risk |
| D02 | yes | no | `anvien-web/package.json` | `graphology-indices`, `graphology-utils`, `mnemonist`, `pandemonium` | `likely_deadcode` | Imports found only inside unused `src/vendor/leiden/*`; package/lockfile references remain. | tied to D01 |
| D03 | yes | no | `anvien-web/package.json` | `@sigma/edge-curve`, `axios`, `d3`, `graphology-layout-force`, `graphology-layout-forceatlas2`, `graphology-layout-noverlap`, `lru-cache`, `react-zoom-pan-pinch`, `uuid`, `zod` | `likely_deadcode` | `rg` found package/lockfile references and no source imports for checked direct dependencies. | medium indirect dependency risk |
| D04 | yes | no | `anvien-web/src/components/ToolCallCard.tsx` | `default` export | `confirmed_deadcode` | Source imports use named `ToolCallCard`; declaration exports both named component and default export; no default import found. | low |
| D05 | yes | no | `anvien-web/src/config/ui-constants.ts` | `DEFAULT_OLLAMA_BASE_URL`, `DEFAULT_OPENROUTER_BASE_URL`, `REQUIRED_NODE_VERSION` | `confirmed_deadcode` | `rg` over app/test found declarations only for these exports; no imports/call sites. | low/medium stale-intent risk for `REQUIRED_NODE_VERSION` |
| D06 | yes | no | `anvien-web/src/core/llm/settings-service-local-runtime.ts` | legacy provider setting helpers | `likely_deadcode` | App source search matched declarations only; tests/mocks reference several helpers. | medium legacy compatibility risk |
| D07 | yes | no | `anvien-web/src/services/backend-client.ts` | unused API client helpers | `likely_deadcode` | `rg` over app/test found no call sites for listed helper functions. | medium reserved HTTP wrapper risk |
| D08 | yes | no | `internal/cli/lazy_action.go` | `createLazyAction` | `confirmed_deadcode` | Go `deadcode` reported `createLazyAction` unreachable; `rg` showed declaration and tests only. | low |
| D09 | yes | no | `internal/cobol/copy_expander.go` | `CopyReplacing`, `parseReplacingClause`, `replacementScanner`, `isReplacementWordChar` | `likely_deadcode` | Go `deadcode` reported `parseReplacingClause` unreachable; `rg` showed declarations and tests only. | medium planned COBOL feature risk |
| D10 | yes | no | `internal/cobol/cobol.go` | `extractProgram` | `confirmed_deadcode` | Go `deadcode` reported unreachable; `rg` found declaration and legacy tests only. | low/medium legacy test helper risk |
| D11 | yes | no | `internal/communities/enrichment.go` | LLM enrichment API and helpers | `likely_deadcode` | Go `deadcode` reported `EnrichClusters` and `EnrichClustersBatch` unreachable; `rg` found declarations/tests only. | medium planned capability risk |
| D12 | yes | no | `internal/communities/communities.go` | `CommunityColors`, `CommunityColor` | `likely_deadcode` | Go `deadcode` reported `CommunityColor` unreachable; `rg` found declarations/tests only. | low/medium exported internal helper risk |
| D13 | no | yes | `internal/httpapi/phase_timer.go` | `phaseTimer`, `newPhaseTimer`, methods | `confirmed_deadcode` | Go `deadcode` reported `newPhaseTimer` unreachable; `rg` found test call sites only. | low |
| D14 | no | yes | `internal/parser/ast_cache.go` | `ASTCache`, `NewASTCache`, methods | `likely_deadcode` | Go `deadcode` reported `NewASTCache` and methods unreachable; `rg` found tests plus declarations/method receivers only. | medium exported internal API risk |
| D15 | no | yes | `internal/providers/golang/definitions.go` | `goTypeSpecLabel` | `confirmed_deadcode` | Go `deadcode` reported unreachable; `rg` found declaration plus live callers of `goTypeSpecLabelForKind`, not `goTypeSpecLabel`. | low |
| D16 | no | yes | `internal/scopeir/kinds.go` | `CallIndex` | `likely_deadcode` | Exact `rg` found only declaration. | medium exported enum/string compatibility risk |
| D17 | no | yes | `anvien-web/src/core/llm/types.local-runtime.ts` | `AgentStep` | `likely_deadcode` | Exact `rg` found only exported interface declaration. | medium exported type compatibility risk |
| D18 | no | yes | `anvien-web/src/core/llm/types.local-runtime.ts` | `GRAPH_SCHEMA_DESCRIPTION` | `false_positive` | Source comment says it is retained to preserve old export surface for compatibility imports; test asserts it. | intentional compatibility export |

False positives:

| Candidate id | Method source | Reason |
|---|---|---|
| D18 | Anvien | Intentionally retained compatibility export per inline source comment and test coverage. |

Uncertain candidates:

| Candidate id | Method source | Uncertainty reason | Follow-up needed |
|---|---|---|---|
| none | both | no candidate required `uncertain`; public/generated/planned-surface risks were classified as `likely_deadcode` or `false_positive`. | none for this benchmark |

## E7 - Final Comparison Evidence

Date: 2026-06-02

Status: complete

Required comparison facts:

| Question | Evidence |
|---|---|
| How many tokens did the agent spend without Anvien? | Native `agent_session_tokens=856990` from isolated Codex CLI runtime usage. |
| How many tokens did the agent spend with Anvien? | Anvien `agent_session_tokens=1524993` from isolated Codex CLI runtime usage. |
| Which mode used fewer agent-session tokens? | Native mode used fewer tokens by 668,003; Anvien spent 77.95% more than native. |
| Which mode read fewer files? | Native read fewer targeted files: 10 vs Anvien 11. |
| Which mode used fewer search/tool calls? | Anvien used fewer native follow-up searches, 11 vs native 39, but more total commands when 30 Anvien commands are included. |
| Which mode found more confirmed/likely deadcode? | Native found more: 12 confirmed/likely vs Anvien 5 confirmed/likely. |
| Which mode produced fewer false positives? | Native produced fewer: native 0 false positives vs Anvien 1. |
| Which candidates were found by both/native-only/Anvien-only? | Found by both: 0; native-only: 12; Anvien-only: 6. |
| Was token measurement valid for both modes? | Yes. Both discovery modes used isolated `codex exec --json` sessions and `turn.completed.usage` telemetry. |

Required summary shape:

```text
Native mode:
- agent_session_tokens or proxy: 856990
- token validity: valid
- search/tool calls: 39 completed commands, 0 Anvien commands
- file reads: 10 targeted files, 55,619 bytes
- candidates: 12

Anvien mode:
- agent_session_tokens or proxy: 1524993
- token validity: valid
- Anvien calls: 30
- follow-up search/tool calls: 11 non-Anvien commands
- file reads: 11 targeted repo files, 69,934 bytes
- candidates: 6

Shared verification:
- candidates verified: 18
- confirmed/likely/uncertain/false-positive: 7 / 10 / 0 / 1
```

## E8 - Closure

Date: 2026-06-02

Status: complete

Closure checks:

| Check | Result |
|---|---|
| No deadcode deletion/edit was made | yes |
| Token measurement valid for native mode | yes |
| Token measurement valid for Anvien mode | yes |
| Plan checklist updated | yes |
| Benchmark ledger complete | yes, except final validation commit hash |
| Final comparison written | yes |
| Commit hash for documentation update, if committed | recorded in final response because a commit cannot contain its own final hash |
