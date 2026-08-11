# QA Report

Scope: Child 01 `P1-E` non-UI identity integration, bounded target oracle, and preservation boundary.
Runtime: full-build packaged `E:\Anvien\anvien\bin\anvien.exe` CLI `1.2.8`.
Browser: N/A — declared scope is a CLI/analyzer graph boundary with no browser or desktop UI.
Build evidence: `E1-P1E-BUILD1` — exact full build exited `0` in `114.4s`.
Anvien evidence: `E1-P1E-DETERMINISM1`, `E1-P1E-INTEGRITY1`, `E1-P1E-TARGET1`, `E1-P1E-BOUNDARY1`.
No-fix boundary: QA reports only. No product code, test, fixture, target source, or Git configuration fix was made.

## Runtime Contexts

- persona: CLI developer/operator validating graph accuracy.
- owner/app/scope: `E:\Anvien` packaged analyzer against a repo-local identity fixture and `E:\cheapapp.org` target.
- role: local user with normal filesystem access.
- shift/session/subscription: N/A.
- dataset/fixture: independently authored P1-C identity fixture plus real bounded target file `modules/email/server/operations/email-operations-observability.ts`.
- locale: N/A.
- viewport: N/A.

## Inventory Summary

- visible surfaces: CLI JSON output, persisted Graph JSON, Ladybug Cypher rows, Git/source/hash comparison.
- interactive elements: N/A — command-driven non-UI surface.
- actions: full build; five local analyzes; provider/identity/collision QA; target analyze; two exact Cypher queries; target post-state check; 11-family and 61-package regressions.
- data/source-of-truth checks: four source declarations, four Variable GraphIDs/scopes, four `DEFINES`, target HEAD/status/per-path/source hashes.
- blocked: none in declared QA scope.
- out of scope: persistence/reader redesign, binding/export/module/ambient semantics, scanner repair, UI/Playwright.
- unverified: none in declared QA behavior scope; Supervisor acceptance is a separate gate.

## Coverage Verdict

- Pass: build, local runtime, conflict fault, target `4/4`, target preservation, focused/sibling/product regression.
- Fail: 0.
- Blocked: 0.
- N/A: browser, screenshot, Playwright, DB/UI controls.
- 100% of declared scope: yes.

## Action Ledger

| Surface | Route/flow | Control/action | Context | Expected | Actual | Verdict | Evidence |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Full build | repository build | run `scripts/full-build.ps1` | Anvien worktree at P1-D HEAD + P1-E ledgers | packaged runtime builds before QA | exit `0`, CLI `1.2.8`, self-analyze complete | Pass | `E1-P1E-BUILD1` |
| Local identity runtime | fixture analyze | run packaged analyze five times | hash-equal identity fixture | same canonical identity/edges every run | `5/5`, graph `20/19`, canonical and whole-graph hashes equal | Pass | `E1-P1E-DETERMINISM1` |
| Owned fault boundary | compiled Go test executable | run public `ResolveBoundInto` conflict test | malformed duplicate ScopeIR | clear conflict; no silent replacement | all range/property mismatch directions PASS | Pass | `E1-P1E-INTEGRITY1` |
| Target runtime | normal in-place analyze | analyze `E:\cheapapp.org` | captured dirty target pre-state | zero parse failures and fresh normal output | `1,359/887/0`, graph `86,164/115,888` | Pass | `E1-P1E-TARGET1` |
| Target oracle | Ladybug node query | query label-exact `Variable` time/now rows | oracle file only | exactly four declarations | `time@207`, `time@214`, `now@262`, `now@501` | Pass | `E1-P1E-TARGET1` |
| Target endpoints | Ladybug relation query | query `CodeRelation` where `type='DEFINES'` | exact four Variable rows | four unique targets and valid File endpoint | edge count `4`, unique Variable IDs `4`, unique File endpoint `1` | Pass | `E1-P1E-TARGET1` |
| Target preservation | Git/source/hash readback | compare full post-state | 13 pre-existing entries | no target-source/worktree drift | HEAD/status/hash and `13/13` file hashes unchanged | Pass | `E1-P1E-BOUNDARY1` |
| Focused regression | Go tests | TSJS + identity/collision + graph | accepted P1-B/C/D owners | all exact owner cases pass | `3/3`, `9/9` + 3 subcases, packages `2/2` | Pass | `E1-P1E-INTEGRITY1` |
| Caller sibling regression | Go tests | 11 production caller families | mixed `AddNode` contracts | preserve all siblings | `11/11` packages PASS | Pass | `E1-P1E-INTEGRITY1` |
| Product regression | Go tests | exact product package inventory | 72 total minus 11 non-standalone corpus packages | all 61 product packages pass | 56 `ok` + 5 no-test, 0 failures | Pass | `E1-P1E-INTEGRITY1` |

## Findings

No runtime bug, usability friction, spec drift, or unresolved test gap was found inside the declared P1-E QA scope.

## Source-Of-Truth Checks

| Visible value/action | Runtime source/API/DB | Expected source state | Actual source state | Verdict | Evidence |
| --- | --- | --- | --- | --- | --- |
| `time@207` | target source -> TSJS -> GraphID/Ladybug | distinct declaration in `parseTime` | unique ID, occurrence `207:8`, scope `202:0-210:1` | Pass | `E1-P1E-TARGET1` |
| `time@214` | target source -> TSJS -> GraphID/Ladybug | distinct reducer-callback declaration | unique ID, occurrence `214:10`, scope `213:46-221:3` | Pass | `E1-P1E-TARGET1` |
| `now@262` | target source -> TSJS -> GraphID/Ladybug | distinct report-builder declaration | unique ID, occurrence `262:8`, scope `254:7-417:1` | Pass | `E1-P1E-TARGET1` |
| `now@501` | target source -> TSJS -> GraphID/Ladybug | distinct report-reader declaration | unique ID, occurrence `501:8`, scope `497:7-634:1` | Pass | `E1-P1E-TARGET1` |
| Target dirty state | Git porcelain + 13 file hashes | identical to pre-state | exact status hash and all 13 hashes match | Pass | `E1-P1E-BOUNDARY1` |

## Console/Network/Runtime Evidence

- Final target analyze stderr is empty; exit code is `0`.
- No browser, network, or UI surface exists in the declared scope.
- Screenshot evidence is N/A because no browser/desktop automation or user-visible UI is involved. Runtime truth comes from the packaged CLI, persisted graph query, and Git/source readback.
- Automated tests are supporting regression evidence; the target QA verdict is based on normal runtime analyze/query and post-state evidence.

## Blockers And Unverified Scope

None inside declared QA scope. Artifact lifecycle, Supervisor acceptance, detect-changes, and commit are workflow gates outside this no-fix QA verdict.

## Final Decision

Pass for the declared Child 01 P1-E non-UI QA scope.

Handoff to supervisor - QA scope completed; awaiting acceptance decision.
