# Supervisor Report: Child 01 Pn-A Acceptance

Verdict: PASS

Review status: CLOSED

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260811_184507_by_gpt-5-codex_child01_pna.md`
- Review time: `260811 184507-193106 +07:00` (`Asia/Bangkok` / `SE Asia Standard Time`)
- Reviewer: `gpt-5-codex`
- Repo/project: `E:\Anvien`
- Branch / review HEAD: `master` / `d03afd7f5ca50754d3448554e3522672ad45cc13`
- Parent P1-D: `10f66ffdd084a4a8710fd347836a1a083d4021bc`
- Scope reviewed: Child 01 Pn-A, complete P1-A through P1-E result
- Claim reviewed: Child 01 closes the bounded identity/range/lexical-owner/occurrence-conservation/canonical-collision invariant family, reaches the normal-runtime `4/4` target, preserves later-Child boundaries, and has consistent evidence/benchmark/artifact/commit lifecycle sufficient for Child-level acceptance.
- Authority used: latest Owner delegation and session rules; `AGENTS.md`; Supervisor skill; graph-accuracy roadmap and contract; all four Child 01 ledgers; bounded problem-origin finding and causal/bounded verification reports; current production source and exact commit diffs; durable coder/QA/Supervisor evidence.
- Explicit authority exclusion: the problem report's DRAFT repair architecture is not implementation authority. The killed hidden P1-E Supervisor lane has no durable verdict and was neither resumed nor used.

## Executive Summary

- Problem: the bounded target carried four distinct `time`/`now` declaration occurrences through ScopeIR but the baseline graph retained only two because graph identity omitted occurrence/lexical inputs and duplicate canonical node IDs could replace payloads silently.
- Decision: Child 01 is accepted. Current source, commit diffs, fresh graph evidence, full build, focused and product regression, normal target runtime, target preservation, and lifecycle/ledger audits jointly close the complete Child 01 claim.
- Required outcome: accepted at Pn-A. The orchestration agent/main session records this report as `E2-PNA-SUPERVISOR1` and owns any subsequent plan transition; this Supervisor does not edit ledgers or open another slice.
- Blocking findings: none.

## Acceptance Matrix

| Invariant | Verified result | Strongest evidence |
|-----------|-----------------|--------------------|
| Range and selection input | TSJS retains construct range plus declaring-token selection range; the contract is one-based lines, zero-based UTF-8 byte columns, exclusive end; LF/CRLF, Unicode, tab, owner/meaning, and JSON round-trip cases pass | current P1-B source/diff; fresh TSJS `3/3`; `E1-P1B-*`; `E1-P1E-INTEGRITY1` |
| Lexical owner and meaning | graph identity consumes lexical scope, provider meaning label, owner, semantic name, arity, occurrence ID, and cleaned repository-relative path | current `buildWorkspace`/`graphIDForDef` source; fresh P1-C `4/4`; `E1-P1C-*` |
| Identity injectivity and determinism | distinct same-name occurrences and supported meaning lanes receive distinct deterministic IDs; nil arity remains distinct from zero; reordered inputs preserve the canonical set | current P1-C source/test; non-invalidated matched runtime `5/5`; current target `4/4` |
| Occurrence conservation | local production denominator remains `defsByFile`; fixture retains `10/10` Definition targets and `10/10` `DEFINES`; target retains `4/4` unique Variable targets and `4/4` `DEFINES` | `E1-P1C-TEST1`, `E1-P1C-ORACLE1`, `E1-P1E-INTEGRITY1`, fresh target query |
| Canonical collision handling | same-invocation canonical payloads require exact equality; compatible pre-existing enrichment is preserved; non-exact collision fails before relationships and propagates through `ResolveBoundInto` with no result graph | current P1-D source/diff; fresh collision/enrichment/endpoints matrix; `E1-P1D-*` |
| Endpoint integrity | local affected relationships have zero missing endpoints; target has four valid `DEFINES` edges to four unique Variable nodes and one valid File endpoint | fresh focused tests, compiled-boundary evidence, target Cypher evidence |
| Scope containment | no persistence/reader or later-Child production owner changed; generic graph mutation and relationship merge remain preserved | exact commit/path inventory and source sibling sweep |

Residual same-invariant unverified surfaces inside Child 01: **none**.

## Source-Level Clearance Notes

- `internal/scopeir/facts.go`: clear. P1-B adds only optional `SelectionRange *Range` to the existing Definition fact; shared `Range` shape and unrelated provider behavior remain unchanged.
- `internal/providers/tsjs/nodes.go`: clear. `nodeRange` records the source-proven TSJS coordinate and interval contract without an unproved encoding conversion.
- `internal/providers/tsjs/definitions.go`: clear. `collector.addDefinition` retains the construct range and records the declaring `nameNode` selection range; the non-identifier binding-pattern branch remains preserved for Child 03.
- `internal/resolution/indexes.go`: clear. P1-C's deterministic length-prefixed tuple uses cleaned repository-relative file, semantic name, nil-distinct arity, provider occurrence ID, lexical scope, and owner, with provider label as the meaning prefix. `defsByFile` remains the occurrence denominator; no old/new compatibility fallback exists.
- `internal/resolution/emit.go`: clear. P1-D uses an invocation-local canonical Definition payload map, requires exact same-run equality, permits only compatible pre-existing enrichment, and returns collision before `DEFINES` or owner relationships are accepted.
- `internal/resolution/resolve.go`: clear. The emission error is propagated through `ResolveBoundInto`; the failed path returns no result graph.
- Failure-boundary consumers: clear. Analyze/CLI/HTTP source checks the resolution error before database/snapshot/meta/result registration, so a canonical collision is not presented as successful output.
- Preserved siblings: clear. `Graph.AddNode`, `Graph.init`, `Graph.AddRelationship`, `emitRelationship`, File enrichment, standalone `BuildDefinitionIndex`, normalization/sort owners, other providers, persistence/readers, and later semantic producers were not repurposed as a proxy fix.

## Exact Commit and Scope Boundary

The reviewed history is linear and each accepted implementation slice has an isolated Git boundary:

| Slice | Commit | Paths | Production / test-or-fixture | Boundary result |
|-------|--------|------:|------------------------------:|-----------------|
| P1-A | `8cbf6006` | 6 | `0 / 0` | contract/ledger-only authority ratification |
| P1-B | `6b3a3fc4480c7f4c2291d8004207f50b52d91d21` | 12 | `3 / 1` | range/selection inputs |
| P1-C | `df0d7b7b753620f56c932f0e13391c1c016a8f72` | 19 | `1 / 9` | identity, fixtures, and affected semantic assertions |
| P1-D | `10f66ffdd084a4a8710fd347836a1a083d4021bc` | 12 | `2 / 1` | canonical Definition conflict handling |
| P1-E | `d03afd7f5ca50754d3448554e3522672ad45cc13` | 9 | `0 / 0` | validation/reports/ledgers only |

The only production files changed across Child 01 are:

- `internal/scopeir/facts.go`
- `internal/providers/tsjs/nodes.go`
- `internal/providers/tsjs/definitions.go`
- `internal/resolution/indexes.go`
- `internal/resolution/emit.go`
- `internal/resolution/resolve.go`

No persistence, reader, binding-pattern implementation, export, module/re-export, ambient/external, graph-health, scanner, or target-source production surface appears in the Child 01 diff.

## Blast Radius and Scope Control

Fresh Anvien evidence reports these current upstream blast radii:

- `DefinitionFact`: CRITICAL, `607` symbols / `83` files / `23` modules / `69` processes.
- `nodeRange`: MEDIUM, `8 / 4 / 1 / 0`; `collector.addDefinition`: LOW with `0` upstream symbols.
- `buildWorkspace`: CRITICAL, `46 / 19 / 7 / 25`.
- `graphIDForDef`: CRITICAL, `32 / 15 / 3 / 19`; `graphIdentityPart`: CRITICAL, `23 / 11 / 2 / 5`.
- `emitDefinitionNodes`: CRITICAL, `26 / 10 / 6 / 33`; `ResolveBoundInto`: CRITICAL, `82 / 36 / 23 / 33`.
- Preserved global siblings remain CRITICAL: `Graph.AddNode` `299 / 71 / 8 / 77`, `Graph.init` `317 / 70 / 11 / 84`, and `Graph.AddRelationship` `275 / 65 / 10 / 93`.

These CRITICAL values are scope warnings, not edit prohibitions. The diffs stay within the exact source-proven owners, and current graph health shows zero changed-source degradation, zero persisted resolution gaps, zero degraded nodes, and complete semantic fields across all `95,772` nodes.

## Evidence Checked

### Fresh/current in this Pn-A review

- Git truth: `master` at exact HEAD `d03afd7f5ca50754d3448554e3522672ad45cc13`; before report creation the worktree was clean. Final worktree audit contains only this untracked Pn-A report.
- Fresh Anvien graph: `1,578` scanned / `680` parsed-code / `0` failed; `95,772` nodes / `134,703` relationships. All reviewed production file-detail rows are `stale=false` and `changedSinceAnalyze=false`.
- Source and exact diffs: all P1-B/P1-C/P1-D production owners and the sibling preserve-only surfaces above were inspected at current HEAD.
- Full build: repo-native build exits `0` in `217.1s`, reports CLI `1.2.8`, and completes built self-analyze at `1,578/680/0` with graph `95,772/134,703`.
- Range/owner/binding boundary: TSJS focused matrix `3/3` passes.
- Identity boundary: all four `TestP1CIdentity*` tests pass.
- Collision/enrichment/endpoints: the named P1-D matrix passes, including all three collision subcases; full `internal/resolution` and `internal/graph` packages pass.
- Sibling regression: all exact 11 source-proven `AddNode` caller-family packages pass.
- Product regression: `72` total packages minus only the measured `11` non-standalone fixture-corpus packages gives `61`; exact regression passes `61/61`.
- Durable QA: P1-E Markdown/JSON and coder report exist; QA JSON parses as `anvien.qa.non_ui.v1` and records the owned matrices and `61/61` result.
- Target preservation: target remains at HEAD/branch `a869876ab6262dacde6cd5d432d099a91852a646` / `master`, with `13` pre-existing entries, `0` staged, unchanged status SHA-256 `D39E4DF20380C0F45F8200CEECD6F862C15FAABDE86ED1B0E5EB386422899351`, `13/13` per-path hashes unchanged, unchanged oracle SHA-256 `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C`, and unchanged settings SHA-256 `6C831C0AC49625D09EE54E76042612B92F47ABC47E4A90D1310F0EA45E824B3C`.
- Target runtime: packaged analyze exits `0` after process-local Git safe-directory injection, reports `1,359/887/0` and graph `86,164/115,888`. Graph JSON is `415,998,487` bytes at SHA-256 `CAEAE4DB27B94B78204DB9EF8660DE1B8D275B637A7B16E4F349E33BC8954F09`.
- Target oracle: exact node query returns `time@207`, `time@214`, `now@262`, and `now@501`; exact relationship query returns four `DEFINES`, four unique Variable targets, and one valid File endpoint.
- Artifact lifecycle: all nine retained P1-D final captures exist, parse as JSON, and have nine unique hashes; all named failed/superseded P1-D paths are absent. Current-slice P1-E residual paths are `0`.
- Ledger/report audit: corrected exact-column audit passes `20/20`; all `37/37` plan-required E1 evidence rows are present exactly once and recorded; P1-A through P1-E are checked, Pn-A/Pn-B/Pn-C remain unchecked, `E2-PNA-SUPERVISOR1` remains pending, Pn-B/Pn-C/Child 02 remain closed, and no durable P1-E Supervisor report exists.

### Reused because not invalidated

- P1-E local determinism remains valid at `5/5`: each run produced `20/19`, `10/10` distinct Definition targets, `10/10` `DEFINES`, zero missing endpoints, canonical signature SHA-256 `09DA4C6741D8D35DEA7B8680121F1B449C90E145D03B287E7917F67DB6ADA6FA`, and byte-identical graph SHA-256/size `C6C942B666931CE1AD505E3FE44B52C3887A6C16FBE0B2DAF94C25FE0AF8C1E9` / `44,199`. Child 01 production/tests are unchanged after that evidence, and the fresh build/tests/target graph independently corroborate it.
- P1-B, P1-C, and P1-D durable build/runtime/QA reports are reused for exact failure history and slice-boundary traceability where current Pn-A evidence did not need to repeat a closed gate.
- P1-D's nine retained Anvien captures are reused only after current existence, JSON parse, and unique-hash verification.
- P1-E target pre-state is reused only together with the current post-state/hash comparison proving no invalidation.

### Failed or excluded attempts not used as acceptance proof

- PowerShell parser failures in early inventory/hash commands, truncated raw file-detail/impact displays, and one wrong impact projection were harness/output issues; corrected bounded projections supplied the evidence used above.
- Initial target Git inspection and analyze stopped before scan because of Git `dubious ownership`; the successful retry used process-local safe-directory configuration and did not persist target Git config.
- The first Cypher query requested nonexistent `qualifiedName`; the accepted queries use schema-backed labels/properties.
- The first combined lifecycle assertion used reserved PowerShell `$Error` and timed out; the corrected lifecycle audit passes.
- The first E1-row consistency selector counted references inside other row descriptions; exact Evidence-ID-column parsing passes all `37/37` rows. The failed selector is not a ledger finding.

### Not run

- The five-run P1-E determinism harness was not repeated after compact because neither production/test state nor its input fixture changed; rerunning it would duplicate non-invalidated evidence.
- No persistence/reader record-parity or Ladybug byte-parity gate was run. Normal Ladybug/meta bytes can change with database/timestamp output; record-level persistence and affected-reader consistency belong to Child 02.
- No binding-pattern implementation, export, module/re-export, ambient/external, graph-health-semantic, or scanner-repair gate was run; those belong to later Children or an explicit non-goal.
- Raw `go test ./...` was not treated as the product oracle because the repository contains 11 measured non-standalone fixture-corpus packages; the exact source-backed `61/61` product set was used.
- No production/test/fixture/plan/ledger/coder/QA file was edited. No detect-changes or commit was run for this report-only Supervisor task.
- The killed hidden P1-E Supervisor lane was not resumed, queried, messaged, or reconstructed. Pn-B, Pn-C, and Child 02 were not opened.

## History Closure

- P1-B first review REJECTed only contradictory living-ledger prose. The history-closure review verifies the ledger-only correction with production/tests unchanged and returns PASS.
- P1-C first review had the same living-ledger contradiction class. The history-closure review verifies `50/50` assertions, unchanged production/affected tests, and returns PASS.
- P1-D first formal review REJECTed living-ledger truthfulness plus dead-artifact lifecycle. Exact correction removed 11 failed/retry/superseded paths, retained nine final unique captures, kept production/test hashes unchanged, and the history-closure review returns PASS with residual P1-D same-invariant surfaces `none`.
- P1-E was accepted explicitly by Owner after the hidden lane was terminated. That decision is recorded only as Owner acceptance and is not represented as Supervisor PASS. This Pn-A report is the new independent Child-level Supervisor gate.

No closed gate was repeated merely because of compact. Each reused result was checked for invalidation against current HEAD, production/test diff, worktree, artifact, and target state.

## Invariant Closure and Boundaries

The affected Child 01 invariant family is closed across provider facts, ScopeIR lexical ownership, deterministic graph identity, the production `defsByFile` occurrence denominator, canonical Definition conflict handling, fail-clear propagation, graph-visible ranges/identity inputs, and affected relationship endpoints.

The following are valid successor boundaries, not residual Child 01 work:

- Graph JSON/Ladybug record persistence and affected readers — Child 02.
- Recursive TypeScript binding-pattern extraction — Child 03; Child 01 only preserved the existing branch.
- TypeScript exports — Child 04.
- Module and re-export resolution — Child 05.
- Ambient/external outcomes and related health projection — Child 06.
- The separately known eight-path scanner repair — outside this campaign.

Current Definition nodes do not project selection columns as a new persisted field. Child 01 proves construct/selection positions at their nearest owned provider/ScopeIR boundary and graph-visible identity/line behavior; source and plan authority assign persistence/reader consumption discovery to Child 02. This is a scoped boundary, not an unverified same-invariant surface.

## Overall Evaluation

The implementation is narrowly owned despite CRITICAL blast radius, closes the measured `2/4 -> 4/4` defect without a compatibility fallback or global mutation rewrite, conserves occurrences and endpoints, fails canonical conflicts clearly, and preserves later-Child semantics. Evidence quality is sufficient and current: fresh source/diff/graph/build/test/target checks corroborate the retained deterministic evidence; rejection history is closed; benchmarks match evidence; artifacts and worktrees are controlled; no authority conflict remains.

## Handoff to Orchestration Agent

- Plan/slice: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-01-graph-identity-contract-and-strict-construction/2026-07-28-01-graph-identity-contract-and-strict-construction-plan.md` / `Pn-A`.
- Report: `reports/Supervisor/rp_supervisor_260811_184507_by_gpt-5-codex_child01_pna.md`.
- Evidence IDs: this report satisfies the proof required for `E2-PNA-SUPERVISOR1`; supporting terminal rows are `E1-P1E-BUILD1`, `E1-P1E-DETERMINISM1`, `E1-P1E-INTEGRITY1`, `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1`, `E1-P1E-REPORT1`, `E1-P1E-REVIEW1`, `E1-P1E-DETECT1`, and `E1-P1E-COMMIT1`.
- Commit/HEAD: `d03afd7f5ca50754d3448554e3522672ad45cc13` on `master`; parent P1-D `10f66ffdd084a4a8710fd347836a1a083d4021bc`.
- Current worktree at handoff: only this untracked Supervisor report; no production/test/plan/ledger change and nothing staged.
- Open blockers: none for Pn-A. Ledger/checklist recording is intentionally left to the orchestration agent/main session under the delegated boundary.
- Next owner: orchestration agent (main session) / Owner. This Supervisor stops here and does not open a coder, architect, Pn-B, Pn-C, or Child 02 task.
