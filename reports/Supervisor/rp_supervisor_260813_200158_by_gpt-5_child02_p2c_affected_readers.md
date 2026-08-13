# Supervisor Report: Child 02 P2-C affected readers

Verdict: REJECT

Disposition: RETURN_P2C_WITH_EXACT_BLOCKER

## Metadata and reviewed basis

- Report file: `reports/Supervisor/rp_supervisor_260813_200158_by_gpt-5_child02_p2c_affected_readers.md`
- Review time: `2026-08-13 20:01:58 +07:00` (Asia/Bangkok)
- Reviewer: `gpt-5`
- Repository: `E:\Anvien`
- Branch: `master`
- Baseline/current HEAD: `4d456446fcc49aed0c6d489aa9c63e00d030b53c`
- Staged paths at start and before report write: `0`
- Open implementation slice: Child 02 P2-C only
- Closed boundaries preserved: P2-A and P2-B; P2-D/P2-E and later Children were not opened
- Claim reviewed: C09-C16 are complete at the exact source-proven affected-reader denominator and may be accepted despite the Ladybug VECTOR extension outage.
- Final report SHA-256: computed only after the final write because a file cannot contain its own stable cryptographic hash; the final task handoff reports the exact hash.
- Next owner: Orchestration main

## Executive decision

P2-C cannot be accepted on the current dirty boundary, but the Ladybug VECTOR extension outage is **not** the blocker.

The production changes for C09/C10/C11/C16 are source-correct and contained. C09-C15 pass their direct or nearest-real boundaries. C16 consumes the explicit persisted `CodeEmbedding.label` from the vector result through `ChunkSearchRow`, nearest-chunk dedup, `BestChunkMatch`, hydration grouping, and metadata-table selection; opaque NodeID label parsing is gone; missing label and metadata-query failure return clear errors. Native Ladybug persistence/readback and metadata hydration preserve two same-name Function identities and a same-name Method identity.

The exact blocking evidence is instead a fresh, reproducible sibling regression at the public HTTP consumer of `SemanticSearch`:

```text
go test ./internal/httpapi -count=1

--- FAIL: TestSearchServiceUsesSemanticRuntime
Search() error = semantic search match "Function:alpha" has no persisted label

--- FAIL: TestSearchServiceHybridMergesBM25AndSemanticAndFallsBack
results len = 2, want 3
```

`internal/httpapi/search.go` calls `embeddings.SemanticSearch` directly. Its current test runner at `internal/httpapi/search_test.go:440-451` returns vector rows without the now-required persisted `label`; the explicit hybrid fixture at lines 325-327 does the same. Semantic mode therefore fails closed and hybrid mode falls back to BM25. This does not establish a production data defect—accepted P2-B persistence and the native C16 test prove real rows carry labels—but it does establish that the required blast-radius sibling regression gate is not closed. The plan requires focused tests plus every sibling regression required by impact before unconditional acceptance.

Exact return condition: update only the C16 HTTP semantic/hybrid test fixture contract so every simulated `CodeEmbedding` vector row supplies its explicit persisted label (for these rows, `Function`), then rerun `go test ./internal/httpapi -count=1` and the existing C16/package/native/failure-path gates. Do not add opaque-ID parsing or weaken the fail-closed production behavior. Submit the corrected current boundary for fresh Supervisor review.

## Authority and files inspected

Read in full:

- `AGENTS.md`
- `.agents/skills/working-rules/SKILL.md`
- `.agents/skills/supervisor/SKILL.md`
- `.agents/skills/frontend-development/SKILL.md`
- `.agents/skills/backend-development/SKILL.md`
- `.agents/skills/Data-Integrity/SKILL.md`
- `.agents/skills/Edge-Case/SKILL.md`
- campaign roadmap
- all four Child 02 ledgers: plan, evidence, benchmark, and actual status
- `docs/contracts/graph-accuracy-contract.md`
- accepted P2-A Supervisor report
- accepted P2-B Supervisor report
- current P2-C coder report
- current P2-C QA report
- all three changed production files
- all five focused changed/new test files
- reusable Playwright script
- paired official Playwright JSON/Markdown evidence

Inspected as exact source anchors or same-invariant siblings:

- C12-C15 validate-only source files
- `anvien-web/src/services/backend-client.ts::readFile`
- `internal/httpapi/file.go`
- `internal/httpapi/search.go`
- `internal/httpapi/search_test.go`
- embedding persistence/vector-index lifecycle source
- all six official screenshots at original resolution
- current Git status/diff, untracked inventory, hashes, ports, and shared temp inventory

Verified report hashes:

- coder report: `DC36748AE81F4209F65D03E4F5E3FBE69099E2B4EC51E9E28CD8CE5C7C6B1F4A`
- QA report: `A0462C21BCEB45D87F9919E33EE03EDC575EDEE97EEE626847B577AC8BCDF144`

## Command and evidence matrix

| Command/evidence | Result | What it proves |
|---|---|---|
| `anvien --help` | PASS | Repository Anvien command-surface precondition was observed. |
| Sequential full authority/skill reads | PASS | Review authority, exact P2-C denominator, source-before-test rule, integrity rule, and failure-path bar were established without relying on report summaries. |
| `git status --short`, branch, HEAD, staged/unstaged/untracked inventories | PASS | `master`, exact HEAD `4d45644...`, staged `0`; current dirty boundary is limited to P2-C source/tests/reports and official QA artifacts before this report. |
| SHA-256 inventory | PASS | Both delegated reports and official Playwright script/JSON/Markdown exactly match the expected hashes. |
| Complete production diff and full source reads | PASS | Edits are confined to C09/C10/C11/C16 owners; no alternate suffix/first-match or opaque-label reconstruction remains in the P2-C source surface. |
| `git diff HEAD -- internal/filecontext/context.go internal/mcp/context.go internal/mcp/detect_changes.go internal/mcp/rename.go` | PASS, empty | C12-C15 remain validate-only and unchanged. |
| Residual source searches for `labelFromNodeID`, suffix matching, vector-label propagation, and line conversions | PASS with one sibling lead | No active opaque-ID label parsing or suffix matching remains. Search identified the public HTTP `SemanticSearch` consumer and its stale vector-row fixtures. |
| `npm test -- --run ...` for the three focused frontend files | PASS: `3` files, `11/11` tests | C09 exact identity/suffix-negative, C10 ambiguous/unique grounding, and C11 one-based/exclusive display/read/highlight behavior. |
| `go test ./internal/embeddings ./internal/filecontext ./internal/mcp -count=1` | PASS | C16 package behavior and C12-C15 nearest validate-only regressions pass fresh on the current worktree. |
| `go test ./internal/httpapi -count=1` | **FAIL**, two tests | Public C16 sibling regression is not contract-current: test vector rows omit persisted `label`, causing semantic fail-closed and hybrid BM25 fallback. This is the acceptance blocker. |
| Current tagged native test source plus coder/QA execution evidence | PASS within stated seam | Real Ladybug persists/reads labels; public production `SemanticSearch` issues the vector query and uses production dedup/metadata hydration for opaque same-name rows. Only vector row production by the unavailable extension is seamed. |
| Official Playwright source/evidence inspection | PASS | Script mounts the built production UI, fixtures only external HTTP responses, exercises C09-C11 unchanged, and records console/page/network failures. |
| Six original-resolution screenshot inspections | PASS | Visible state is consistent with the assertions and preserves the approved UI. |
| `git diff --check` and `git diff --cached --check` | PASS | No whitespace error and no staged change. |
| Port/process inventory | PASS | Listener counts on `4848` and `5228` are both zero at final containment check. |
| `.tmp` inventory | PASS/preserved | `.tmp/ladybug-home`, `.tmp/ladybug-native`, and `.tmp/runtime-p2c` exist and were not deleted or modified by the Supervisor. |

Not run:

- No new `anvien analyze --force`, file-detail, impact, or detect-changes command was run. Graph refresh would mutate repository index artifacts despite the delegation allowing exactly one repository write (this report). The current QA graph/impact evidence was produced at the same HEAD/worktree and was checked against current source; the HIGH/CRITICAL blast warnings are preserved below.
- No full build was rerun. QA already performed the required clean-holder Restart Manager/exclusive-open sequence and a successful full build on this exact HEAD/worktree. No source changed after that evidence. The fresh failing HTTP package regression independently prevents PASS, so another build would not close the blocker.
- Official Playwright was not rerun. Hashes, source, paired evidence, screenshots, and current dirty boundary are unchanged; all six screenshots were independently inspected.
- The VECTOR extension was not retried. Coder and QA already supplied independent bounded outage evidence, and the delegation prohibits availability loops.
- No stage, commit, main final detect-changes, cleanup, deletion, production/test/plan/ledger/evidence edit, or P2-D action was performed.

## C09-C16 verdict matrix

| Row | Supervisor verdict | Evidence and reasoning |
|---|---|---|
| C09 `resolveNodeIds` | PASS | Exact graph-set membership only at `useAppState.local-runtime.tsx:667`; graph absence and suffix/fragment inputs yield an empty set. Unit and mounted production evidence show only the full opaque ID is selected. |
| C10 `handleNodeGroundingReference` | PASS | `matchingNodes.length !== 1` fails closed at line 595; unique match retains exact node ID/path and all four numeric coordinates. Same-label/same-name ambiguity creates no citation or panel. |
| C11 `CodeReferencesPanel` | PASS | Graph lines remain one-based; `/api/file` receives the single zero-based conversion; backend file slicing is zero-based inclusive. Exclusive `endLine` with `endCol == 0` excludes the terminal line. Display `L10–11`, scroll target 10, highlights 10-11, and unhighlighted 12 agree. No UI redesign is present. |
| C12 `nodeRange` | PASS, validate-only | Source unchanged; `nodeRange` directly copies `startLine/startCol/endLine/endCol`. Fresh `internal/filecontext` package regression passes. |
| C13 context payloads | PASS, validate-only | Source unchanged; UID/path/line and ambiguity payload behavior remain direct. Fresh `internal/mcp` regression passes. |
| C14 `detectChangedSymbols` | PASS, validate-only | Source unchanged; persisted one-based start/end lines are intersected directly with Git hunk lines. Fresh `internal/mcp` regression passes. |
| C15 `collectRenameChanges` | PASS, validate-only | Source unchanged; target identity is UID-based and the one-based line is converted only at file-array indexing. Fresh `internal/mcp` regression passes. |
| C16 semantic-search hydration | **REJECT for current acceptance boundary** | Production implementation, package behavior, native label persistence/readback, public `SemanticSearch` propagation/dedup/native metadata hydration, same-name identities, and clear missing-label/query-error behavior pass. VECTOR outage is non-blocking. Acceptance nevertheless fails because the real public HTTP sibling suite is red at two semantic/hybrid tests whose vector fixtures omit the new explicit label contract. |

Exact denominator disposition: C09-C15 are `7/7 PASS`; C16 production reader invariant is source/runtime supported, but its required sibling regression gate is not closed. Therefore the current P2-C deliverable is not an unconditional `8/8 PASS` boundary.

## VECTOR-outage acceptance decision

Decision: the official Ladybug VECTOR extension outage does **not** invalidate the actual P2-C reader contract and is not a production defect finding.

Plan/source reasoning:

1. The plan defines P2-C as correction of the exact P2-A readers so they consume P2-B corrected persisted fields. It requires the actual backend/reader boundary and explicitly rejects invented storage or reader contracts.
2. The graph-accuracy contract requires a reader to use an explicit field supplied by its real source contract and forbids reconstructing meaning from an opaque identifier when that field is available.
3. C16's owned defect was `search.go` deriving a metadata table label from NodeID. The corrected code selects `emb.label`, propagates it through vector row/chunk/dedup, and uses it for metadata hydration. That entire interpretation contract is source-proven and exercised.
4. Native Ladybug evidence proves the P2-B persisted label is physically written/read and that production metadata hydration preserves opaque same-name identities across Function and Method tables.
5. Fully native vector-index creation/query is an additional evidence surface whose fresh installation depends on Ladybug's official extension service. The observed `INSTALL VECTOR` failure/HTTP 522 is external availability, not evidence that the C16 source reader violates its explicit-label contract.
6. No incompatible extension substitution, fallback interpretation, or retry loop is allowed or needed for this decision.

Therefore QA's stated VECTOR-only reason for `BLOCKED` is overturned. If the HTTP sibling regression were green, the external outage alone would not prevent P2-C acceptance. Current rejection is solely the reproducible sibling-regression failure above.

## Blast-radius treatment

The accepted/current evidence reports HIGH host-file warnings for the two frontend owners and `internal/embeddings/search.go`; exact C16 types/helpers range from LOW to MEDIUM, while C12-C15 include HIGH/CRITICAL exact impacts. These are scope warnings, not edit prohibitions.

Treatment:

- Production edits remain bounded to C09/C10, C11, and C16 in three files.
- C12-C15 are unchanged and covered by nearest package regression.
- C09-C11 are covered by focused unit and mounted visible runtime evidence.
- C16 is covered by package/native tests and a same-invariant sweep to its public HTTP route.
- That sweep found the one blocking regression in `internal/httpapi/search_test.go`; no additional production owner is inferred.
- No P2-D analyze/read boundary, graph transport, scanner, target repository, or later-Child surface was opened.

## Visual review of all six screenshots

1. `01-mounted-production-fixture.png`: mounted Anvien production UI is coherent; five nodes/four edges and repository state are visible; no missing asset or corruption.
2. `02-before-c09-tool-result.png`: My AI panel is mounted before the tool result; graph and layout are stable.
3. `03-c09-exact-id-only.png`: only `uniqueTarget` is visibly cyan-active; suffix/near-match nodes remain dim; no red impact node is present.
4. `04-c10-ambiguous-fails-closed.png`: ambiguous `Function:sharedName` appears in chat, while no Code Inspector/citation opens.
5. `05-c10-unique-persisted-reference.png`: Code Inspector shows `FUNCTION`, `uniqueTarget`, `src/unique.ts`, and `L10–11`.
6. `06-c11-lines-10-11-highlighted-line-12-excluded.png`: selected viewer targets line 10; lines 10 and 11 are cyan-highlighted; line 12 is visibly unhighlighted; citation remains `L10–11`.

Across the six images there is no observed layout redesign, overlap, clipped critical content, missing font, missing icon, or missing asset. The screenshots are internally consistent with the JSON and Markdown evidence.

## Scope, artifact, temp, and process containment

- Changed production files: exactly three expected files.
- Changed tracked focused tests: exactly four expected files.
- New native test: exactly `internal/embeddings/search_ladybugdb_test.go`.
- C12-C15 source diff: zero.
- Forbidden plan/roadmap/ledger/accepted P2-A/P2-B source diff: zero.
- Staged paths: zero.
- Script SHA-256: `276B4D6EA54E97ED6E99A945BB13368685FE4FDC3CB4C5A6540BB82F8F3DA058`.
- Official JSON SHA-256: `1D46E4DA2689BD91DC228E40636ABEAA4ECF56C867EA288779E4F6FE1348BF69`.
- Official Markdown SHA-256: `0B917711B905F15CAB07712ECC8F2BE43F34C0874725C1C717E4B152F67E3FD3`.
- Official QA directory contains exactly paired JSON/Markdown plus six screenshots. Windows path casing (`Reports/qa` versus `reports/QA`) resolves to the same directory, not duplicate evidence.
- No dead/superseded official QA artifact remains in the current directory. Prior failed attempts described by QA were replaced at deterministic final paths; only the final eight official files are present.
- Shared/ownership-unknown temp directories were inventoried and preserved: `.tmp/ladybug-home` (five directories), `.tmp/ladybug-native` (ten files/five directories), `.tmp/runtime-p2c` (four files). Nothing was deleted.
- Final listener counts: port 4848 = 0; port 5228 = 0.

## Findings

### Blocking finding — C16 public HTTP sibling regression is not label-contract current

- Severity: HIGH process/acceptance blocker; no production corruption reproduced.
- Files: `internal/httpapi/search_test.go:88`, `internal/httpapi/search_test.go:317`, `internal/httpapi/search_test.go:325`, `internal/httpapi/search_test.go:440`.
- Broken acceptance invariant: after correcting production behavior, all affected-reader acceptance plus sibling regression required by impact must pass; reader outputs must agree with the explicit persisted field contract.
- Actual behavior: the HTTP test runner returns vector rows without `label`. `SemanticSearch` correctly fails closed rather than parsing `Function:*`; semantic route test fails, and hybrid route silently exercises its intentional BM25 fallback instead of the expected semantic merge.
- Why this blocks: the failing suite is the real public consumer of C16, not an unrelated or pass-by-default test. A required action remains before unconditional P2-C acceptance.
- Exact fix direction: update the test vector rows only to include their explicit persisted `label` (`Function`) and keep production fail-closed behavior unchanged.
- Re-review evidence required: `go test ./internal/httpapi -count=1` PASS, plus fresh confirmation that `go test ./internal/embeddings ./internal/filecontext ./internal/mcp -count=1` and the focused frontend tests remain PASS. Run the existing tagged native C16 test when its repo-bundled native prerequisites are available; do not retry/install the remote VECTOR extension merely for this correction.

## Non-blocking findings

- The Ladybug official VECTOR extension endpoint outage limits one fully native vector-index evidence surface. It is external availability and does not contradict the C16 source-proven reader contract.
- The native test uses a truthful vector-row seam. This is acceptable evidence for persisted-label propagation/dedup/metadata hydration when paired with production query-shape assertions and real native persistence/readback; it is not represented as fully native vector-index evidence.

## Residual risk

- Fully native vector-index creation/query remains unobserved in this environment until the compatible official extension is available. This is an evidence limitation, not the current blocker and not authorization to substitute an incompatible extension.
- After the exact HTTP fixture correction, the fresh review should ensure no other same-invariant consumer fixture returns a vector row without `label`. Current bounded search found no other active C16 production reconstruction surface.
- P2-D repeated-analyze/failure behavior and P2-E final parity remain intentionally unreviewed and closed.

## Final evaluation

Final verdict: **REJECT**.

Final disposition: **RETURN_P2C_WITH_EXACT_BLOCKER**.

P2-C is returned only because the real HTTP semantic/hybrid sibling regression is currently failing on the new explicit-label row contract. The external VECTOR extension outage is explicitly non-blocking for P2-C. C09-C15 and the C16 production implementation evidence are retained as technically cleared evidence pointers, but a fresh Supervisor review must re-verify the corrected whole boundary; no acceptance or next-slice opening is authorized by this report.

Next owner: **Orchestration main**.
