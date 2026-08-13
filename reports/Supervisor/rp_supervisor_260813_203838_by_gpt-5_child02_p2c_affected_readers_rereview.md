# Child 02 P2-C affected readers — Supervisor re-review

- Review time: 2026-08-13 20:38 Asia/Bangkok
- Role: independent Supervisor, review-only
- Repository: `E:\Anvien`
- Branch: `master`
- Current HEAD: `43aff01e882b787c91da84676e23f3ae28c05720`
- P2-C implementation baseline: `4d456446fcc49aed0c6d489aa9c63e00d030b53c`
- Staged paths at final review: `0`
- Final verdict: **PASS**
- Final disposition: **ACCEPT_P2C_AND_HAND_BACK**
- Next owner: Orchestration main

## Executive decision

The exact blocker from the immutable Supervisor REJECT has been corrected without changing production code. `internal/httpapi/search_test.go` now gives all semantic-success and hybrid-merge vector fixtures the explicit persisted `label` required by the corrected C16 reader contract. Fresh HTTP package and targeted route tests pass. Semantic success, semantic-plus-BM25 hybrid merge, and the intentional embedder-failure BM25 fallback remain separate behaviors; the fixture correction does not mask a production error.

Production C16 remains fail-closed when a vector row lacks `label`, continues to select and propagate `CodeEmbedding.label`, and contains no opaque-NodeID label reconstruction. The only remaining missing-label fixture is the intentional negative case that proves this fail-closed behavior. The current P2-C denominator is therefore unconditional `8/8 PASS`.

The Ladybug official VECTOR extension outage remains non-blocking for P2-C. It is an external limitation on one fully native evidence surface, not a production defect and not a failed invariant in the source-proven affected-reader contract. It was not retried, substituted, or assigned silently to P2-D/P2-E.

## Git and chronology basis

`git branch --show-current` and `git rev-parse HEAD` confirmed `master` at `43aff01e882b787c91da84676e23f3ae28c05720`. `git merge-base --is-ancestor 4d456446fcc49aed0c6d489aa9c63e00d030b53c HEAD` returned success.

Commit `43aff01e...` is unrelated concurrent history. `git show --stat --format=fuller 43aff01e...` shows exactly one path, `internal/aicontext/skills/orchestration/SKILL.md`, with one insertion and one deletion under message `docs(skills): update orchestration rules`. It changes no P2-C source, test, dependency, build manifest, plan, contract, or evidence path and is excluded from P2-C ownership.

The current unstaged/untracked boundary contains the expected P2-C production files, focused tests, reusable Playwright script, final QA artifacts, and reports. The follow-up-specific diff against current HEAD is confined to `internal/httpapi/search_test.go`. Staged count is zero. The unrelated committed skill file does not appear in the dirty P2-C diff.

## Authority and files inspected

Authority was re-read and anchored from:

- `AGENTS.md`
- `.agents/skills/working-rules/SKILL.md`
- `.agents/skills/supervisor/SKILL.md`
- the four Child 02 P2-C ledgers: plan, evidence, benchmark, and actual status
- `docs/contracts/graph-accuracy-contract.md`
- immutable prior REJECT report `reports/Supervisor/rp_supervisor_260813_200158_by_gpt-5_child02_p2c_affected_readers.md`
- coder follow-up report `reports/coder/rp_coder_260813_202428_by_gpt-5_child02_p2c_http_label_fixture_followup.md`

The re-review inspected the current Git status/diff and the exact rejected-invariant sources/tests:

- `internal/httpapi/search_test.go`
- `internal/httpapi/search.go`
- `internal/embeddings/search.go`
- `internal/embeddings/search_test.go`
- `internal/embeddings/search_ladybugdb_test.go`

It also revalidated the retained P2-C path and evidence inventory:

- `anvien-web/src/components/CodeReferencesPanel.tsx`
- `anvien-web/src/hooks/useAppState.local-runtime.tsx`
- the three focused frontend tests
- C12-C15 validate-only source paths
- reusable Playwright script, paired JSON/Markdown, and six screenshots
- original coder and QA reports
- shared temp, port, staging, and artifact containment

No private main attachment was requested or read.

Verified authority/report hashes:

| Artifact | SHA-256 | Result |
|---|---|---|
| Immutable REJECT report | `9829F9D2FA483C1DDE97E981CA3AC1B93A4379F5DBEF409506CC66E3F76A60C7` | Exact match |
| Coder follow-up report | `CFEB5B73EB512C9143CEC42D56BD27CC487E5DE193893F1A9EB97AE141A5C251` | Exact match |
| Original P2-C coder report | `DC36748AE81F4209F65D03E4F5E3FBE69099E2B4EC51E9E28CD8CE5C7C6B1F4A` | Exact match |
| Original P2-C QA report | `A0462C21BCEB45D87F9919E33EE03EDC575EDEE97EEE626847B577AC8BCDF144` | Exact match |

## Exact blocker closure

The prior blocker was not a production failure. The public HTTP sibling fixtures supplied vector rows that violated the newly explicit persisted-label row contract, so correct production fail-closed behavior made the semantic test fail and drove the hybrid test into its intentional BM25 fallback.

The follow-up diff contains exactly three value additions in `internal/httpapi/search_test.go`:

1. `Function:shared` gains `"label": "Function"`.
2. `Function:semantic-only` gains `"label": "Function"`.
3. The default `semanticSearchRunner` row for `Function:alpha` gains `"label": "Function"`.

There is no assertion weakening, alternate fallback, fixture deletion, production edit, or restoration of `labelFromNodeID`. These labels model the persisted row field that production `SemanticSearch` now requires. They allow the public HTTP tests to exercise their intended success/merge branches instead of accidentally testing missing-label failure.

The fresh full HTTP package command passed:

```text
go test ./internal/httpapi -count=1
ok github.com/tamnguyendinh/anvien/internal/httpapi 3.215s
```

The Supervisor also ran the two exact public behaviors with verbose selection:

```text
go test ./internal/httpapi -run '^(TestSearchServiceUsesSemanticRuntime|TestSearchServiceHybridMergesBM25AndSemanticAndFallsBack)$' -count=1 -v
--- PASS: TestSearchServiceUsesSemanticRuntime
--- PASS: TestSearchServiceHybridMergesBM25AndSemanticAndFallsBack
PASS
ok github.com/tamnguyendinh/anvien/internal/httpapi 0.240s
```

The hybrid test still proves two distinct outcomes in one case: a successful semantic-plus-BM25 merge returns the expected combined identities/sources, while the separate offline-embedder subcase intentionally falls back to BM25. The semantic-only service test remains a successful semantic route. Explicit labels repair the fixtures at their actual input contract and do not force production success.

## Command and evidence matrix

| Command/evidence | Result | What it proves |
|---|---|---|
| Full authority and relevant source reads | PASS | Exact denominator, immutable rejection reason, follow-up claim, source contract, and review-only limits were established from primary repo evidence. |
| `git branch --show-current`; `git rev-parse HEAD`; ancestor check | PASS | Exact `master`/`43aff01e...` basis and baseline ancestry. |
| `git show --stat --format=fuller 43aff01e...` | PASS | Concurrent commit contains only the unrelated orchestration skill path and excludes cleanly from P2-C. |
| `git diff 43aff01e... -- internal/httpapi/search_test.go` | PASS | Follow-up is exactly the three explicit `Function` labels; no other change in that file. |
| Full reads of HTTP and embeddings production/tests | PASS | Success, hybrid merge, fallback, explicit-label propagation, and fail-closed paths remain distinct and current. |
| Residual searches for vector-row fixtures and opaque-label reconstruction | PASS | Every positive P2-C HTTP/embeddings vector fixture has explicit label; only the intentional missing-label negative omits it; `labelFromNodeID` and equivalent active reconstruction are absent. |
| `go test ./internal/httpapi -count=1` | PASS, package 3.215s | The exact two former failures and their HTTP siblings are green on the final source. |
| Targeted verbose HTTP test command | PASS, package 0.240s | Direct independent closure of semantic success and hybrid merge/fallback. |
| `go test ./internal/embeddings -count=1` | PASS, package 2.949s | C16 propagation/dedup/hydration and clear missing-label/query-error behavior remain green. |
| Tagged local-native C16 seam | PASS, test 0.53s/package 2.491s | Real Ladybug writes/reads explicit labels and production hydration preserves opaque same-name Function/Method identities; no remote VECTOR operation was made. |
| Retained backend packages `./internal/filecontext ./internal/mcp` | PASS in final coder evidence | C12-C15 nearest validate-only regressions remain valid because neither their source nor dependencies in this follow-up/concurrent commit changed. |
| Retained focused frontend run | PASS, 3 files and 11/11 tests | C09-C11 evidence remains valid because follow-up and concurrent commit touched no frontend source/test/dependency. |
| C12-C15 and plan/contract diff checks | PASS, empty | Validate-only owners and forbidden planning/contract authority remain unchanged. |
| Original report and Playwright artifact hashes | PASS, exact matches | Prior source/QA/visible-runtime evidence is current and was not rewritten by the follow-up. |
| `git diff --check`; `git diff --cached --check` | PASS | No whitespace errors and no staged changes. |
| Listener inventory for ports 4848 and 5228 | PASS, both zero | No P2-C app/fixture listener was left running. |
| Shared `.tmp` inventory | PASS/preserved | Ownership-unknown native/runtime directories remain present and were not cleaned or deleted. |

Not run:

- No full build was repeated. The coder's final clean-holder build passed on the corrected P2-C source boundary in 121.4 seconds (`1601/686/0`). The later/current HEAD difference only records the already-present unrelated skill content and changes no P2-C source, test, dependency, or build manifest. Repeating the build would add no source-currentness evidence. Because no build/retry was invoked, the Restart Manager/holder/exclusive-open preflight was not applicable in this review.
- No P2-A/P2-B or unchanged C09-C15 gate was rerun solely to duplicate retained evidence. The follow-up is test-only in `internal/httpapi/search_test.go`; current concurrent history does not invalidate those gates.
- No graph refresh, Anvien detect-changes, stage, commit, cleanup, delete, or next-slice operation was performed. A graph refresh would mutate repo index state beyond the single permitted report write; current graph/blast evidence was retained.
- No VECTOR availability/install/query retry was made. The tagged native seam used only the repository-local Ladybug prerequisites under `.tmp/ladybug-native/v0.19.1/windows-x86_64`.

## C09-C16 verdict matrix

| Row | Verdict | Re-review basis |
|---|---|---|
| C09 `resolveNodeIds` | PASS | Retained exact-membership and suffix-negative source/unit/mounted evidence remains unchanged and uninvalidated. Exact node identity is preserved; absence/near-match fails closed. |
| C10 `handleNodeGroundingReference` | PASS | Retained evidence remains unchanged: only exactly one persisted identity opens grounding; ambiguous same-name/same-label matches fail closed and coordinates remain exact. |
| C11 `CodeReferencesPanel` | PASS | Retained source/unit/visible evidence remains unchanged: one-based lines, zero-based UTF-8 byte columns, exclusive ends, and the single API line-index conversion agree without UI redesign. |
| C12 `nodeRange` | PASS, validate-only | Source diff remains empty; direct coordinate copying and retained nearest package test remain valid. |
| C13 context payloads | PASS, validate-only | Source diff remains empty; exact identity/path/coordinate and ambiguity payload behavior remains valid. |
| C14 `detectChangedSymbols` | PASS, validate-only | Source diff remains empty; persisted one-based ranges remain compared to Git hunk lines without an extra conversion. |
| C15 `collectRenameChanges` | PASS, validate-only | Source diff remains empty; UID identity and the single file-array indexing conversion remain intact. |
| C16 semantic-search hydration | PASS | HTTP blocker is closed; production consumes persisted `emb.label` end-to-end, keeps same-name identities, deduplicates chunks by exact identity, hydrates the label-selected metadata table, and clearly errors for missing label/query failure. Native seam and public package/HTTP tests pass; opaque-ID reconstruction is absent. |

Exact denominator: **C09-C16 = 8/8 PASS**.

## C16 integrity and edge-case conclusions

- The production vector query returns `emb.label AS label` with node ID, chunk bounds, and distance.
- Row parsing requires a non-empty label. `Function:parseable-but-label-missing` remains an intentional negative fixture and still errors instead of inferring `Function` from the NodeID.
- Label travels through the vector row/chunk/dedup path and controls metadata-table hydration. Same-name Function/Method identities remain separate because identity is not reconstructed from display name or suffix.
- Metadata-query errors remain explicit. Missing label cannot silently become an HTTP semantic success.
- Hybrid fallback remains an intentional response to semantic/embedder failure and is not confused with the successful hybrid-merge case after the fixture correction.
- Bounded residual search found no other P2-C HTTP/embeddings success fixture or consumer that supplies a vector row without explicit label.

## VECTOR outage acceptance decision

Decision retained: the official Ladybug VECTOR extension outage does **not** invalidate P2-C and is not a production defect.

The plan defines P2-C as correction of source-proven affected readers so they consume the corrected persisted fields from P2-B. The graph-accuracy contract requires explicit source fields to be used and forbids recovering semantics from opaque IDs. C16's owned reader defect was label reconstruction from NodeID; current production instead selects persisted `CodeEmbedding.label`, propagates it through the vector result/dedup pipeline, and uses it for metadata hydration.

Real native evidence proves that Ladybug persists and reads the explicit label and that public production hydration preserves opaque same-name identities. The unavailable piece is fresh creation/query of the compatible official VECTOR index, whose installation depends on the external Ladybug extension endpoint. External HTTP 522/install unavailability supplies no evidence that current source violates the affected-reader contract. Substituting an incompatible extension or hiding the issue behind retries would reduce evidence quality and is prohibited.

Accordingly:

- P2-C acceptance is unconditional for its actual contract.
- The outage remains a non-blocking environmental evidence limitation.
- It is not silently moved into P2-D or P2-E. If a later explicitly authorized gate requires a fully native vector-index run, its owner may re-establish that evidence when the compatible service is available; that does not create a current P2-C code defect or authorize opening a later slice now.

## Blast-radius treatment

Prior graph evidence reported HIGH host-file warnings for the two frontend owners and `internal/embeddings/search.go`, with HIGH/CRITICAL warnings among validate-only C12-C15 surfaces. These are scope warnings, not edit prohibitions.

The follow-up touched only the HTTP sibling test identified by the previous blast-radius/same-invariant sweep. Production owners, C09-C15 tests, validate-only sources, API shapes, plans, and later slices did not expand. Fresh HTTP and embeddings tests plus the local-native seam cover the corrected contract at package, public consumer, persistence/readback, and hydration boundaries. No new production blast radius was introduced.

## Visual and artifact evidence disposition

The immutable prior Supervisor report records direct original-resolution inspection of all six official screenshots:

1. mounted production fixture and coherent graph/runtime state;
2. stable pre-C09 tool-result state;
3. exact-ID-only C09 selection;
4. ambiguous C10 grounding failing closed;
5. unique persisted reference with `L10–11`;
6. lines 10–11 highlighted while exclusive terminal line 12 is excluded.

The current review verified that the reusable script, paired JSON/Markdown, and the same six screenshot files remain at the deterministic official paths and that the three authoritative hashes are unchanged:

- Script: `276B4D6EA54E97ED6E99A945BB13368685FE4FDC3CB4C5A6540BB82F8F3DA058`
- JSON: `1D46E4DA2689BD91DC228E40636ABEAA4ECF56C867EA288779E4F6FE1348BF69`
- Markdown: `0B917711B905F15CAB07712ECC8F2BE43F34C0874725C1C717E4B152F67E3FD3`

The follow-up changes only backend test fixtures and cannot alter those visible states. The official QA directory contains exactly the paired JSON/Markdown and six screenshots; no dead/superseded official QA artifact is present. Rerunning Playwright would not test the corrected HTTP fixture invariant and was not needed.

## Scope, artifact, process, and temp containment

- Production diff remains confined to the expected C09/C10, C11, and C16 owners.
- Follow-up diff is exactly `internal/httpapi/search_test.go` with three explicit labels.
- C12-C15 source diff is zero.
- Plan, roadmap, ledgers, contracts, accepted P2-A/P2-B artifacts, rules, and old reports were not edited by this Supervisor.
- Current staged path count is zero; `git diff --cached --check` is clean.
- Ports 4848 and 5228 have zero listeners.
- `.tmp/ladybug-home` (0 files/5 directories), `.tmp/ladybug-native` (10 files/5 directories), and `.tmp/runtime-p2c` (4 files/0 directories) were inventoried and preserved.
- No cleanup, deletion, detect-changes, stage, commit, VECTOR retry, target-repository mutation, or P2-D/P2-E opening occurred.
- This new Supervisor report is the only write made by this re-review.

## Findings and residual risk

Blocking findings: **none**.

Non-blocking findings:

- The compatible Ladybug official VECTOR extension remains externally unavailable, limiting a fully native vector-index creation/query evidence surface. This is neither a production defect nor a P2-C acceptance failure.
- The native test truthfully seams only vector-row production while using real Ladybug persistence/readback and production `SemanticSearch` hydration; it is not represented as a fully native vector-index test.

Residual risk:

- A fully native compatible vector-index run remains unobserved in this environment. It may be re-evidenced only under a future authorized gate when the official service is available, with no incompatible substitution.
- P2-D repeated-analyze/failure behavior, P2-E final parity, and later Children remain closed and unreviewed.

No exact P2-C invariant remains failed.

## Final evaluation

Final verdict: **PASS**.

Final disposition: **ACCEPT_P2C_AND_HAND_BACK**.

P2-C receives unconditional `8/8 PASS` for the actual affected-reader contract. The previous HTTP semantic/hybrid regression blocker is closed at its exact fixture contract; retained evidence remains current; no external outage has been converted into a false production defect or silently ignored.

Next owner: **Orchestration main**.
