# Anvien Deadcode Agent Token Benchmark Evidence Ledger

Date: 2026-06-02

Status: complete

Companion files:

- Plan: [2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md](2026-06-02-anvien-deadcode-agent-token-benchmark-plan.md)
- Benchmark ledger: [2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md](2026-06-02-anvien-deadcode-agent-token-benchmark-benchmark.md)

## Reset Notice

Old evidence from the invalidated benchmark run has been removed. This ledger starts clean for the rerun.

## Evidence Rules

1. Record commands, source reads, token-accountant observations, and candidate evidence as they happen.
2. Keep quantitative metric tables in the benchmark ledger.
3. Keep native-search and Anvien-guided discovery separate until union verification.
4. Native-search evidence must not include Anvien commands, outputs, resources, generated context, graph artifacts, or prior Anvien reports.
5. Anvien-guided evidence must record graph freshness before graph-based commands.
6. Token evidence must record only context the token accountant observes the main agent receive, read, or emit.
7. Record every Anvien command output, artifact read, summary, and error that the accountant observes during Anvien-guided work.
8. Record unobserved, redirected, capture-only, and truncated content with the accountant's observation status instead of assuming it is counted or uncounted by type.
9. Record failures, retries, and abandoned paths because they consume agent work.
10. Record candidate verification from source-backed facts, not intuition.
11. Do not record deadcode deletion or cleanup patches because this benchmark only finds candidates.

## Evidence Template

Use this template for each phase:

```text
## E<n> - <Phase>

Date:

Status:

Scope:

- ...

Commands / reads:

| Step | Command or file | Purpose | Accountant observation | Token-accountant note | Result |
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
| Token accountant active | pending |
| Completion condition met | pending |
| Open leads remaining | pending |
| Blocker or incomplete reason | pending |
| Confidence | pending |
```

## E0 - Baseline

Date:

Date: 2026-06-02 11:19:24 +07:00

Status: complete

Required evidence:

| Check | Result |
|---|---|
| Baseline commit | `6564d7d5f5f7d53767a4afbc1028cda26535b977` |
| Discovery start commit | `e0b2f336b37a7d65eec6ddf5d6f20ac7dfd40900`; source code unchanged from baseline, P0 docs only were committed before discovery |
| Branch | `master` |
| Worktree status | clean at baseline (`git status --short` produced no output) |
| Source-code dirty state | 0 dirty source files at baseline |
| Benchmark docs/reports dirty state | 0 dirty benchmark docs/reports at baseline |
| Shell | PowerShell 7.6.2 |
| OS / CPU / RAM | Microsoft Windows 10 Pro 10.0.19045 64-bit; Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz; 4 cores / 8 logical processors; 33,238,466,560 bytes visible memory |
| Go / Node / npm / Git versions if used | Go 1.26.3 windows/amd64; Node v24.15.0; npm 11.12.1; Git 2.54.0.windows.1 |

## E1 - Token Accountant Setup

Date:

Date: 2026-06-02 11:19:24 +07:00

Status: complete

Required evidence:

| Item | Result |
|---|---|
| Accountant identity/mechanism | Main-agent self-observation ledger. `tool_search` exposed sub-agent tools, but no passive observer can automatically see future main-agent tool results unless the main agent forwards them, so the benchmark uses explicit self-tracking of observed context. |
| Can observe main-agent tool results | Yes, for tool outputs that appear in this transcript. Each phase records command/tool text and the observed stdout/stderr/result text the main agent actually receives. |
| Can distinguish observed context from unobserved full-output proxy | Yes. A proxy line such as `full_stdout_proxy_tokens: 228504` is counted only as that observed line, not as the hidden body it describes. Hidden redirected bodies are unobserved until explicitly read. |
| Can record tool-call argument text | Yes. Every shell command, Anvien command, and other tool-call argument used during discovery is recorded as observed emitted context for its phase. |
| Can record source/file reads | Yes. File content is counted only when opened into the main-agent transcript; separate byte/character summaries are not treated as the hidden file body unless the body is also read. |
| Can record agent response text | Yes. Phase reports written into evidence/benchmark ledgers are counted as agent-emitted context for the phase. |
| Truncation handling rule | If a tool output is truncated and the observed portion cannot be measured, token comparison for that phase is invalid. If the output is summarized intentionally, only the summary that entered context is counted. |
| Blocker if exact observed-context accounting is unavailable | No current blocker for P1 start; validity remains conditional on recording each phase without unmeasured truncation. |

## E2 - Native-Search Discovery Without Anvien

Date: 2026-06-02

Status: complete

Rules for this section:

- Do not use Anvien.
- Record every command output and source file read observed by the token accountant.
- Record candidates before any Anvien-guided work starts.

Native command/read log:

| Step | Command or file | Purpose | Accountant observation | Token-accountant note | Result |
|---|---|---|---|---|---|
| N1 | `rg --files` inventory and manifest reads | identify source surface without Anvien | observed summaries and selected manifest content | no Anvien surface used | complete |
| N2 | keyword/declaration/reference scans | find unused/deadcode leads | 29 native command/read groups, 108,350 observed chars total | one timeout/retries counted; no truncation reported by native agent | complete |
| N3 | targeted source reads | classify native leads | 10 files visibly read | large internal scans were not counted as file content unless printed/read | complete |

Native candidates:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| N1 | `internal/cli/setup_command.go` | `setupMCPEntry` | Go private type | exact native search found only declaration | low dynamic risk |
| N2 | `internal/session/controller.go` | `NewControllerWithResolver` | Go exported constructor under `internal` | exact native search found only declaration | medium public/internal API risk |
| N3 | `internal/session/controller.go` | `(*Controller).GetSession` | Go exported method under `internal` | exact native search found only declaration | medium public/internal API risk |
| N4 | `anvien-web/src/components/DropZone.tsx` | `data` | unused callback parameter | targeted ESLint reported unused parameter at line 228 | low risk |
| N5 | `anvien-web/e2e/graph-orientation-labels.spec.ts` | `_error` | unused catch binding | targeted ESLint reported unused binding at line 443 | low risk |
| N6 | `anvien-web/test/unit/GraphCanvas.selection-performance.test.tsx` | `sigmaOnSpy`, `sigmaOffSpy`, `cameraOnSpy`, `cameraOffSpy` | unused test locals | targeted ESLint reported all four locals unused at lines 14-17 | may indicate missing assertions |
| N7 | `anvien-web/src/generated/anvien-contracts.ts` | file-level `/* eslint-disable */` | stale generated directive | targeted ESLint reported unused eslint-disable at line 1 | generated source-of-truth risk |
| N8 | `anvien-web/src/lib/lucide-icons.tsx` | `EyeOff`, `FileArchive`, `FlaskConical`, `Github`, `Heart`, `Upload` | unused local icon re-exports | exact search found each icon name only in the barrel export block | medium public barrel risk |

Native completion:

| Item | Result |
|---|---|
| Declared native procedure recorded before first search | yes |
| Token accountant closed native phase | yes |
| Completion condition met | yes |
| Open native leads remaining | none reported |
| Blocker or incomplete reason | none |
| Confidence | valid native discovery ledger; no Anvien use |

## E3 - Native-Search Discovery Report

Date: 2026-06-02

Status: complete

Required evidence:

| Item | Result |
|---|---|
| Native candidate count | 8 |
| Native unique files read | 10 |
| Native command/search count | 29 shell/read/static-analysis groups |
| Native observed-context token total | 27,088 estimated tokens from 108,350 observed chars |
| Native completion status | complete; no Anvien surface used |
| Native unresolved questions | icon barrel public-use risk; generated directive source-of-truth risk; exported `internal/session` API intent |

## E4 - Anvien-Guided Discovery

Date: 2026-06-02

Status: complete; token accounting invalid

Rules for this section:

- Record graph freshness before graph-based work.
- Do not seed discovery from the native candidate list.
- Record every Anvien command and every Anvien-related context observed by the token accountant.
- Record source reads after Anvien narrows candidate leads.

Graph freshness:

| Check | Result |
|---|---|
| Analyze command | `anvien\bin\anvien.exe analyze --force` |
| Analyze output accountant observation | 1,063 observed stdout/stderr chars |
| Indexed commit | graph refreshed during phase; command succeeded |
| Current commit | benchmark baseline was `6564d7d5f5f7d53767a4afbc1028cda26535b977`; P0 doc commit existed before discovery docs were written |
| Fresh/stale result | fresh after analyze; no worktree changes left by analyze |

Anvien command/read log:

| Step | Command or file | Purpose | Accountant observation | Token-accountant note | Result |
|---|---|---|---|---|---|
| A1 | `anvien analyze --force`, `query`, `cypher`, `context`, `graph-health` | surface graph leads before broad source reads | 24 Anvien calls; reconstructed 777,694 Anvien stdout/stderr chars | multiple large `context` outputs were transcript-truncated | complete but token invalid |
| A2 | follow-up exact native searches | confirm graph-surfaced leads and reject dynamic/public leads | 17 shell invocations / 40 `rg` executions; 77,869 observed chars | one broad native search output was transcript-truncated | complete but token invalid |
| A3 | targeted source snippets | classify graph leads | 12 snippets, 11 unique files, 6,160 chars | source reads were small and targeted | complete |

Anvien candidates:

| Candidate id | Path | Symbol/name | Kind | Discovery evidence | Notes |
|---|---|---|---|---|---|
| A1 | `internal/mcp/context.go` | `contextCategorizedRefs` | Go private function | Anvien zero-inbound lead; exact search found only declaration in `internal/mcp` | high confidence |
| A2 | `internal/mcp/context.go` | `contextClassLikeIncomingRefs` | Go private function | Anvien zero-inbound lead; exact search found only declaration in `internal/mcp` | superseded by set-based helper |
| A3 | `internal/mcp/context.go` | `contextProcessParticipation` | Go private function | Anvien zero-inbound lead; exact search found only declaration in `internal/mcp` | high confidence |
| A4 | `internal/mcp/tools.go` | `containsScore` | Go private function | Anvien zero-inbound lead; exact search found only declaration | high confidence |
| A5 | `internal/mcp/route_shape_impact.go` | `mcpRouteAnalysisRecords` | Go private function | Anvien zero-inbound lead; exact search found only declaration in `internal/mcp` | high confidence |
| A6 | `internal/mcp/server.go` | `writeMessage` | Go private function | Anvien zero-inbound lead; exact search found only declaration | high confidence |
| A7 | `internal/providers/tsjs/nodes.go` | `isFunctionScopeNode` | Go private function | Anvien zero-inbound lead; exact search found only declaration in TSJS provider | high confidence |
| A8 | `internal/providers/golang/definitions.go` | `goTypeSpecLabel` | Go private function | Anvien zero-inbound lead; exact search found only declaration in Go provider | high confidence |
| A9 | `internal/providers/golang/nodes.go` | `descendantsOfType` | Go private function | Anvien zero-inbound lead; scoped package search found only declaration | high confidence |
| A10 | `internal/providers/rust/nodes.go` | `directChildOfKind` | Go private function | Anvien zero-inbound lead; scoped Rust provider search found only declaration | high confidence |
| A11 | `internal/providers/rust/nodes.go` | `directChildrenOfKind` | Go private function | Anvien zero-inbound lead; scoped Rust provider search found only declaration | high confidence |
| A12 | `internal/providers/java/nodes.go` | `firstNamedChildOfType` | Go private function | Anvien zero-inbound lead; scoped Java provider search found only declaration | high confidence |
| A13 | `internal/providers/python/nodes.go` | `namedChildrenOfType` | Go private function | Anvien zero-inbound lead; scoped Python provider search found only declaration | high confidence |
| A14 | `internal/session/controller.go` | `NewControllerWithResolver` | Go exported constructor under `internal` | Anvien zero-inbound lead; exact search found only declaration | also found by native |

Anvien completion:

| Item | Result |
|---|---|
| Declared Anvien procedure recorded before first graph command | yes |
| Token accountant closed Anvien phase | yes |
| Completion condition met | yes for candidate discovery |
| Open Anvien leads remaining | none reported |
| Blocker or incomplete reason | token comparison invalid because observed output was truncated |
| Confidence | candidate discovery usable; token-cost comparison unusable |

## E5 - Anvien-Guided Discovery Report

Date: 2026-06-02

Status: complete; token accounting invalid

Required evidence:

| Item | Result |
|---|---|
| Anvien candidate count | 14 |
| Anvien unique files read | 11 |
| Anvien command count | 24 |
| Anvien follow-up native search count | 17 shell invocations / 40 `rg` executions |
| Anvien observed-context token total | invalid; reconstructed observed categories were 861,723 chars / 215,431 estimated tokens, but transcript truncation prevents exact accounting |
| Anvien completion status | complete for candidate discovery; invalid for token winner |
| Anvien unresolved questions | large Anvien `context` output must be controlled in a rerun before token reduction can be claimed |

## E6 - Candidate Union And Verification

Date: 2026-06-02

Status: complete

Verification rules:

- Verify the union of candidates from both methods.
- Check static references, dynamic/public entrypoint risk, generated-code status, test/build/runtime hooks, and external contract hints.
- Do not delete or edit candidate code.

Candidate verdicts:

| Candidate id | Found by native | Found by Anvien | Path | Symbol/name | Verdict | Verification evidence | Dynamic/public risk |
|---|---|---|---|---|---|---|---|
| D01 | yes | no | `internal/cli/setup_command.go` | `setupMCPEntry` | `confirmed_deadcode` | exact search found only declaration; snippet shows private struct declaration | low |
| D02 | yes | yes | `internal/session/controller.go` | `NewControllerWithResolver` | `likely_deadcode` | exact search found only declaration | medium: exported under Go `internal` package |
| D03 | yes | no | `internal/session/controller.go` | `GetSession` | `likely_deadcode` | exact search found only declaration | medium: exported method under Go `internal` package |
| D04 | yes | no | `anvien-web/src/components/DropZone.tsx` | `data` | `confirmed_deadcode` | targeted ESLint reported unused parameter at line 228; snippet shows callback body ignores it | low |
| D05 | yes | no | `anvien-web/e2e/graph-orientation-labels.spec.ts` | `_error` | `confirmed_deadcode` | targeted ESLint reported unused catch binding at line 443 | low |
| D06 | yes | no | `anvien-web/test/unit/GraphCanvas.selection-performance.test.tsx` | `sigmaOnSpy`, `sigmaOffSpy`, `cameraOnSpy`, `cameraOffSpy` | `confirmed_deadcode` | targeted ESLint reported all four assigned but unused at lines 14-17 | low for liveness; medium test-quality risk |
| D07 | yes | no | `anvien-web/src/generated/anvien-contracts.ts` | `/* eslint-disable */` | `likely_deadcode` | targeted ESLint reported unused directive at line 1 | medium: generated file; generator is source of truth |
| D08 | yes | no | `anvien-web/src/lib/lucide-icons.tsx` | `EyeOff`, `FileArchive`, `FlaskConical`, `Github`, `Heart`, `Upload` | `likely_deadcode` | exact search found each icon only in barrel export block | medium: public local barrel risk |
| D09 | no | yes | `internal/mcp/context.go` | `contextCategorizedRefs` | `confirmed_deadcode` | scoped search in `internal/mcp` found only declaration | low |
| D10 | no | yes | `internal/mcp/context.go` | `contextClassLikeIncomingRefs` | `confirmed_deadcode` | scoped search in `internal/mcp` found only declaration; newer set-based helper is used instead | low |
| D11 | no | yes | `internal/mcp/context.go` | `contextProcessParticipation` | `confirmed_deadcode` | scoped search in `internal/mcp` found only declaration | low |
| D12 | no | yes | `internal/mcp/tools.go` | `containsScore` | `confirmed_deadcode` | exact search found only declaration | low |
| D13 | no | yes | `internal/mcp/route_shape_impact.go` | `mcpRouteAnalysisRecords` | `confirmed_deadcode` | scoped search in `internal/mcp` found only declaration | low |
| D14 | no | yes | `internal/mcp/server.go` | `writeMessage` | `confirmed_deadcode` | exact search found only declaration; lower-level `writeRawMessage` path remains | low |
| D15 | no | yes | `internal/providers/tsjs/nodes.go` | `isFunctionScopeNode` | `confirmed_deadcode` | scoped TSJS provider search found only declaration | low |
| D16 | no | yes | `internal/providers/golang/definitions.go` | `goTypeSpecLabel` | `confirmed_deadcode` | scoped Go provider search found only declaration; callers use `goTypeSpecLabelForKind` | low |
| D17 | no | yes | `internal/providers/golang/nodes.go` | `descendantsOfType` | `confirmed_deadcode` | scoped Go provider search found only declaration | low |
| D18 | no | yes | `internal/providers/rust/nodes.go` | `directChildOfKind` | `confirmed_deadcode` | scoped Rust provider search found only declaration | low |
| D19 | no | yes | `internal/providers/rust/nodes.go` | `directChildrenOfKind` | `confirmed_deadcode` | scoped Rust provider search found only declaration | low |
| D20 | no | yes | `internal/providers/java/nodes.go` | `firstNamedChildOfType` | `confirmed_deadcode` | scoped Java provider search found only declaration | low |
| D21 | no | yes | `internal/providers/python/nodes.go` | `namedChildrenOfType` | `confirmed_deadcode` | scoped Python provider search found only declaration | low |

False positives:

| Candidate id | Method source | Reason |
|---|---|---|
| none | both | no false positives found in this static verification pass |

Uncertain candidates:

| Candidate id | Method source | Uncertainty reason | Follow-up needed |
|---|---|---|---|
| none | both | no candidate required `uncertain`; public/generated risks were classified as `likely_deadcode` instead of confirmed | if cleanup is planned, inspect intended public/generator contract first |

## E7 - Final Comparison Evidence

Date: 2026-06-02

Status: complete

Required comparison facts:

| Question | Evidence |
|---|---|
| Which method used fewer observed-context tokens? | invalid. Native was measured at 27,088 estimated tokens; Anvien phase had at least reconstructed 215,431 estimated tokens in observed categories, but exact accounting is invalid due transcript truncation. |
| Which method read fewer files? | native read fewer files: 10 vs Anvien 11. |
| Which method used fewer search/tool calls? | Anvien used fewer native follow-up shell invocations, 17 vs 29, but more total discovery calls when 24 Anvien calls are included. |
| Which method found more confirmed/likely deadcode? | Anvien-guided found 14 confirmed/likely candidates; native found 8. |
| Which method produced fewer false positives? | tie: 0 false positives in this static verification pass. |
| Which candidates were found by both/native-only/Anvien-only? | both: D02; native-only: D01, D03-D08; Anvien-only: D09-D21. |
| Was the token accountant able to measure exact observed context? | native yes; Anvien no because several observed outputs were truncated. |
| Were unobserved/redirected/capture-only/truncated contents handled by observation status? | yes. Truncated Anvien outputs caused invalid token result instead of a false winner. |

Required summary shape:

```text
Native search:
- observed-context total tokens: 27,088 valid estimated tokens
- task prompt: included in native self-ledger
- tool call arguments: 3,975 estimated tokens
- search output: 5,950 estimated tokens
- file reads: 2,575 estimated tokens
- agent response: 1,550 estimated tokens
- retry/error: 663 estimated tokens
- files read: 10
- candidates: 8

Anvien-guided:
- observed-context total tokens: invalid; reconstructed category floor 215,431 estimated tokens
- task prompt: included in Anvien phase self-ledger
- tool call arguments: not exact after truncation
- Anvien observed context: reconstructed 777,694 chars, but exact token accounting invalid
- follow-up search output: reconstructed 77,869 chars
- file reads: 6,160 chars
- agent response: not exact after truncation
- retry/error: 6 failed/retry calls
- files read: 11
- candidates: 14

Shared verification:
- observed-context total tokens: shared verification cost recorded separately and not used as either discovery method's discovery cost
- candidates verified: 21
- confirmed/likely/uncertain/false-positive: 17 / 4 / 0 / 0
```

## E8 - Closure

Date: 2026-06-02

Status: complete

Closure checks:

| Check | Result |
|---|---|
| No deadcode deletion/edit was made | yes |
| Token accountant ledger complete | complete enough to invalidate token axis honestly; native exact, Anvien invalid due truncation |
| Plan checklist updated | yes |
| Benchmark ledger complete | yes |
| Final comparison written | yes |
| Commit hash for documentation update, if committed | not embedded because a commit cannot contain its own final hash; final response records the documentation commit hash |
