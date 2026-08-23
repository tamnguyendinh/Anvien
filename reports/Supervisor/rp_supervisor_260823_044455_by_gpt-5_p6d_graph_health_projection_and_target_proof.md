# Independent Supervisor Review — P6-D Graph-Health Projection and Target Proof

## Review identity and boundary

- Role: independent visible Supervisor, review/acceptance only; no repair.
- Active slice: P6-D “Project outcomes and prove the target sites”. P6-A, P6-B, P6-C1, P6-C2, and P6-C3 remain closed; Pn-A/Pn-B/Pn-C and Child 07 remain locked.
- Repository: saved project `E:\Anvien`, branch `master`.
- Git anchor: HEAD `0d9803777880d8ecca09cae90592d4728a231160`; parent `2acd357db32ac0c2bd3c3968a950c773acf3000d`.
- Review claim: assess the exact seven-path P6-D candidate for full P6-D acceptance, independently test current local source/test correctness, and decide whether the coder's blocked disposition is truthful. A truthful blocked handoff is not acceptance.
- Forbidden boundary preserved: no alternate worktree; no target repository access; no network, install, or package script; no stage/commit/push/branch/stash/reset/checkout; no source, test, or ledger edit; no cleanup or reuse of quarantined roots.
- Next owner: Visible Main task `01a02b25-9dbc-78a1-8bd3-bbba39c10a31`.

The official Main handoff was independently sealed before relying on it: `11,569` bytes / `174` LF / `0` CR / strict UTF-8 without BOM / SHA-256 `A8953B70F5A6866FEFBB96D972E25662CC069ECDDFBAB2B4B85F2F13CB5D84BC`. All four current Child 06 ledgers and the full coder report were reread. Production diff was read before test claims were evaluated. Current code was compared with the committed P6-B authority validator and P6-C3 final-outcome contract.

## Exact candidate identity

| # | Relative path | Bytes | SHA-256 |
|---:|---|---:|---|
| 1 | `internal/graphhealth/diagnostics.go` | 26,410 | `796ADF6B3D7B0D643F34E114408EAFE9C003B3B2551AAD328287204C649CBA03` |
| 2 | `internal/graphhealth/p6d_outcome_projection_test.go` | 18,197 | `2222BFC2F769C394A47CA840BBF5B27038C27A746FB908A259F19D88D4CC3CBB` |
| 3 | `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md` | 81,475 | `E59DCC9488D6FF13354AE83271F73F2CB7D4D119292811541C78F66E884FFF0E` |
| 4 | `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-actual-status.md` | 60,884 | `5EED3B37FCC046CD7577D582B55B3505B69096E1485C5DCB980293CDCC8AD815` |
| 5 | `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-evidence.md` | 127,776 | `E91CDCE068FC14C938ACA349EBAEA9C25268139C50384C5081C57871A495F48D` |
| 6 | `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-benchmark.md` | 21,072 | `65694ECA30E88935C56AD011DF424CA098099030644B0C1FE9FB2FC789F10028` |
| 7 | `reports/coder/rp_coder_260823_032449_by_gpt-5_p6d_graph_health_projection_blocked.md` | 15,101 | `2F2E09ED46DE482696949C0BBC4FE71DDD28BFD26AF829135515084083AB2DAA` |

All seven current byte identities match the supplied authority and total exactly `350,915` bytes. The supplied aggregate `E8B515231097313E657BF3EE77EC6A8C01EBF1C236687E98D5FEA5B181EC9D70` is independently reproducible by ordinal/case-sensitive relative-path sorting, rows encoded as `relative-path|bytes|SHA256`, and an LF after every row, including the last. For convention clarity, the same path-sorted rows without the terminal LF produce `9CDCC292A27DE73A5068B0B88547E4151C5C06F36FEFE8B505D383E66D67C29F`; individual file seals are unaffected.

Pre-report Git truth was exact: status `58 = 5 tracked + 53 untracked`; the untracked set contained the P6-D test plus `52` reports split as `42 Investigation + 6 Supervisor + 4 coder`; index entries `0`; `git diff --check` exited `0`. The seven candidate paths had exactly the expected five tracked modifications plus two untracked files.

## Blocking local correctness finding

### P6D-AUTHORITY-IDENTITY-1 — nested authority can drift from its outer final outcome

The structured carrier is intended to contain one final outcome and the immutable P6-B authority for the same handled source site. Current production construction makes that relationship explicit:

- `internal/resolution/types.go:50-71` defines `TypeScriptAuthorityResult` as the lossless result for one handled TypeScript source site.
- `internal/resolution/outcome.go:52-71` defines `ResolutionOutcome` as the final decision for one stable source site, with the P6-B result nested as `Authority`.
- `internal/resolution/outcome.go:352-368` clones the P6-B authority and copies its `sourceSiteId`, stage, `siteKind`, `filePath`, `fileHash`, range, `requestedName`, and `requestedMeaning` directly into the outer outcome.
- `internal/resolution/source_site.go:31-41` derives the stable `sourceSiteId` from normalized file path, fact family/site kind, target/requested name, and all four range coordinates. A carrier retaining the old ID while changing one of those identity components is internally contradictory.

The current P6-D consumer does not enforce that contract. `internal/graphhealth/diagnostics.go:410-428` validates the nested object internally, then compares only nested `sourceSiteId`, stage, and the compatible status family to the outer outcome. It does not compare:

1. `filePath`;
2. `fileHash`;
3. all four range coordinates;
4. `siteKind`;
5. `requestedName`;
6. `requestedMeaning`.

The committed P6-B validator at `internal/resolution/resolve.go:1037-1079` proves that a `TypeScriptAuthorityResult` is internally well-formed; it cannot prove equality to a separate outer object. The committed P6-C3 validator at `internal/resolution/outcome.go:195-215` also compares only source-site ID, stage, and status family. Therefore neither sibling validator closes the cross-object identity hole for graph-health ingestion.

A bounded hostile probe against the current bytes mutated each of the six fields above only inside the nested authority while leaving the outer outcome and flat diagnostic unchanged. Every mutation was still classified as:

```text
standard_library / non_actionable
```

The required fail-closed result is:

```text
unclassified / review
```

The probe deliberately exited nonzero after observing all six false acceptances. Its temporary source was deleted; no candidate file changed. The only residual executable is an unused Go debug-cache entry under the fresh Supervisor root, `gocache\01\...\probe.exe`, `2,311,168` bytes / SHA-256 `01FC1010DA153EC3322C39308B52A94314F368DCB61E64B6313674E88A428DBD`; it is not `anvien.exe`, was not used as runtime evidence, and cannot satisfy any built-current-runtime gate.

Existing hostile rows at `internal/graphhealth/p6d_outcome_projection_test.go:113-174` cover malformed JSON, schema drift, outer/flat identity drift, nested source-site/stage/status drift, incomplete authority, proof-state failures, and resolved-diagnostic rejection. They do not mutate any of the six nested-versus-outer fields above. The passing matrix therefore does not detect this defect.

This is a blocking P6-D source/test correctness defect, independently of all missing runtime and target evidence. Accepting the carrier as standard-library/non-actionable can suppress an analyzer-gap signal using authority proof that belongs to a different file, range, site kind, requested symbol, or meaning.

Minimum repair evidence for this invariant is exact cross-object equality in `validStructuredResolutionAuthority` for all six fields, including every range coordinate, plus six durable hostile regression rows that each assert `unclassified/review`. Because the committed P6-C3 validator shares the same cross-object omission, Main must route that same-invariant contract concern explicitly; this Supervisor did not expand scope or edit the closed slice.

## Source behavior otherwise confirmed

Subject to the blocker above, the candidate's mechanical mapping is source-backed:

- A valid structured repository `unresolved` outcome maps to `in_repo_unresolved/analyzer_gap` at `diagnostics.go:311-318`.
- A valid TypeScript standard-library `unresolved` or `capability_unavailable` outcome maps to `standard_library/non_actionable` at `diagnostics.go:311-322`.
- A structured carrier that is malformed, incomplete, drifted by currently checked fields, unsupported, or a resolved outcome attached as a diagnostic maps to `unclassified/review` at `diagnostics.go:303-325`.
- `decodeStructuredResolutionOutcome` treats an outcome status marker, `schemaVersion`, or `status` field as structured presence; malformed-present carriers cannot fall through to name heuristics. Only a note with no structured marker retains legacy fallback at `diagnostics.go:327-349` and `549-589`.
- `rawJSONObjectPresent` distinguishes absent/`null` from a real JSON object and rejects arrays/strings/invalid objects at `diagnostics.go:535-541`.
- Catalog proof validation matches the P6-B ready/missing/rejected shape and reason families at `diagnostics.go:468-520`: ready requires the three hashes; missing requires zero hashes and exact missing reason; rejected requires artifact hash, no authority/catalog hash, and one accepted rejection reason.
- A resolved structured outcome carried as a diagnostic is fail-closed because only unresolved and capability-unavailable statuses receive a non-review projection; the explicit hostile test row exercises this.
- All diagnostic ingestion forms normalize through `normalizeDiagnosticMetadata` (`diagnostics.go:39`, `94`, `158`, `163`, `169`, `219`). `SourceBackedResolutionGapInputs` reads those normalized diagnostics before `resolutionGapInputFromDiagnostic` projects them, so no separate sibling adapter bypass was found.
- No target-name literal or target-specific branch exists in the candidate production or P6-D test. Target names appear only in governance/evidence/report text describing the still-unrun gate.

## Fresh Anvien graph evidence and blast radius

`anvien --help` was run before graph use. A new E-only Supervisor root was used at `E:\Anvien\.tmp\p6d-supervisor-review-01a02b4f-260823-01`; no prior cache/root was reused. Forced-current `anvien analyze --force` completed with:

- `2,068` files scanned;
- `763` code files parsed;
- `0` parse failures;
- `117,865` nodes;
- `164,424` relationships;
- `20,310` dependency edges;
- `469` unresolved projections.

Current file detail preserved the broad review surface:

| File | Risk | Symbols | Inbound / outbound / local | Unresolved | Flows | Tests |
|---|---|---:|---:|---:|---:|---:|
| `internal/graphhealth/diagnostics.go` | HIGH | 154 | 48 / 70 / 133 | 186 | 1 | 19 |
| `internal/resolution/outcome.go` | HIGH | 119 | 86 / 124 / 117 | 147 | 11 | 23 |
| `internal/resolution/resolve.go` | HIGH | 200 | 132 / 289 / 87 | 403 | 26 | 40 |
| `internal/graphhealth/policy.go` | HIGH | 138 | 294 / 3 / 81 | 78 | 5 | 19 |
| `internal/graphhealth/resolution_gap_inputs.go` | HIGH | 70 | 131 / 44 / 99 | 98 | 0 | 22 |

Current impact results were not normalized away:

- `diagnostics.go` file impact: CRITICAL; `154` contained symbols, `31` affected files, `34` direct dependents, `1` flow, `19` tests.
- `normalizeDiagnosticMetadata`: CRITICAL; `22` impacted symbols, `10` files, `3` modules, `28` processes.
- `structuredResolutionDiagnosticPolicy`: HIGH; `17` impacted symbols, `7` files, `3` modules.
- `classifyDiagnostic`: HIGH; `17` impacted symbols across `3` modules.
- `SourceBackedResolutionGapInputs`: CRITICAL; `15` impacted symbols, `2` modules, `15` processes.
- `resolutionGapInputFromDiagnostic`: HIGH; `8` impacted symbols and `3` processes.
- Committed contract comparison: `validateResolutionOutcome` CRITICAL with `12` impacted symbols / `33` processes; `validateTypeScriptAuthorityResult` CRITICAL with `17` impacted symbols / `33` processes.

HIGH/CRITICAL are blast-radius warnings, not edit bans. They make the missing hostile equality checks acceptance-blocking and require a narrowly scoped repair with the complete downstream regression boundary.

Fresh current `anvien detect-changes --repo E:\Anvien --scope all` completed after the forced graph:

- file layer HIGH;
- `5` tracked changed files / `198` changed symbols;
- `diagnostics.go` contributes `179` symbols and four ledgers contribute `19` documentation symbols;
- separate summary remains `risk=low`, `affected_count=0`;
- ResolutionGap change inventory `95` entities / `99` occurrences;
- actionability `38` analyzer-gap / `57` non-actionable;
- classification `13` builtin / `38` in-repository / `44` standard-library;
- runtime Resolution Health impact `0`;
- semantic completeness `117,865 / 117,865`.

The file-layer HIGH and separate low summary describe different detector layers and are both retained.

## Build, test, and parity evidence

Before building, the current repo analyze lock was absent/free and process inspection found no Go build, compiler, linker, cgo, GCC, or Clang holder. No process required termination. All build/test commands used fresh E-only `HOME`, temp, `GOCACHE`, `GOMODCACHE`, and `GOPATH` under the Supervisor root with `GOPROXY=off`, `GOSUMDB=off`, and `-buildvcs=false` where applicable.

The full build attempt occurred before validation:

| Boundary | Command | Exit / evidence |
|---|---|---|
| Full repository | `go build -buildvcs=false ./...` | exit `1`, `6,820 ms`; offline module lookups for Cobra, Hugot, and tree-sitter were unavailable, one test fixture has mixed Go packages, and one standalone C fixture is not a Go package |
| Affected package build | `go build -buildvcs=false ./internal/graphhealth` | exit `0`, `11,171 ms` |
| Focused P6-D matrix | `go test ./internal/graphhealth -run '^TestP6D' -count=1 -v` | exit `0`; exact subtest groups `3 + 2 + 20 + 1 + 1` |
| Full affected regression | `go test ./internal/graphhealth -count=1` | exit `0` |

The broad build is a truthful blocked/failed attempt, not a successful full build. The affected package compiles and its current tests pass, but that does not override either the hostile source defect or the missing product-runtime gate.

The parity test was also read rather than inferred from its name. It performs Diagnostic → Graph JSON → ResolutionGap node/relationship → Graph JSON → summary and asserts exactly `3` `standard_library/non_actionable`, `0` `in_repo_unresolved/analyzer_gap`, stable source-site identity, and stable gap relationships. That generic in-memory/JSON boundary passes. It is not native Ladybug or built-reader evidence.

## Blocked disposition versus full P6-D acceptance

The coder report's `BLOCKED_BUILT_CURRENT_RUNTIME / REJECT-READY` disposition is materially truthful:

- the permitted offline full build did not produce a current CLI/runtime;
- the one disclosed network-enabled deviation exited `1` after about `81.199 s`, was followed by Main's STOP correction and coder acknowledgement, and produced no binary;
- no stale or current packaged binary was relabeled as evidence;
- native Ladybug and built-reader parity were not run;
- the exact three target sites remain `0/3` unrun/unproved;
- the independent oracle comparison was not run;
- the target pre/post boundary was not established.

Those facts mean the blocked handoff is honest. They also independently prevent full P6-D acceptance. In particular, local in-memory parity cannot substitute for a built-current-runtime analyze, native Ladybug readback, built CLI/MCP/HTTP reader parity, exact target-site outcomes, an independent oracle, or target worktree/index/graph pre/post evidence.

## Quarantine, runtime, and target-access audit

Forbidden roots were inventoried read-only and were never used, written, cleaned, or treated as evidence:

| Root | Files / bytes / zero-byte files | Relevant subinventory |
|---|---:|---|
| `.tmp\p6c3-repair` | `1,673 / 76,978,267 / 19` | quarantined predecessor |
| `.tmp\p6c3-supervisor-rereview-01a02a84` | `1,662 / 93,162,861 / 18` | quarantined predecessor |
| `.tmp\p6c3-main-closure-01a02abb` | `1 / 343 / 0` | Main-only closure root |
| `.tmp\p6d-coder-01a02abb` | `10,053 / 971,554,982 / 84` | gomodcache `5,687 / 469,406,586`; gocache `3,397 / 281,902,402`; runtime `0 / 0` |
| `.tmp\p6d-main-verification-01a02b25` | `1,512 / 64,175,973 / 19` | Main-only debug root; never runtime proof |

The fresh Supervisor root contains no `anvien.exe`. Its sole `.exe` is the debug-cache probe described above and was not executed as a product runtime.

This Supervisor lane did not access the locked target repository. The official authority records Main's inspection of `153` coder command executions with zero target-path command. Within the available evidence boundary, Git and current filesystem review show no target-produced runtime, oracle, pre/post snapshot, or 3/3 result. The closed coder task was not reopened and the target was not accessed merely to strengthen that claim; therefore the claim is confirmed only to the limit of the sealed Main handoff plus this lane's current command/Git/filesystem evidence.

## Required resubmission evidence

1. Close `P6D-AUTHORITY-IDENTITY-1`: require nested authority equality to outer outcome for `filePath`, `fileHash`, the complete range, `siteKind`, `requestedName`, and `requestedMeaning`; add one hostile regression per drift field; prove each becomes `unclassified/review`.
2. Explicitly route the same-invariant omission in the committed P6-C3 `validateResolutionOutcome` contract without silently expanding this P6-D review.
3. Rerun holder inspection and the full build attempt first, then affected build, focused P6-D matrix, full `internal/graphhealth` regression, and exact JSON/gap/summary parity on the repaired current bytes.
4. Run a fresh forced-current Anvien graph, repeat file-detail/impact, and run current detect-changes while preserving file-layer and summary results separately.
5. Under explicit permitted dependency/runtime authority, build a current product runtime and prove native Ladybug plus built-reader parity. A Go cache executable or stale packaged CLI is not acceptable.
6. Establish the locked target pre-state, exact three-site result (`3/3` correct external/capability outcomes and `0/3` in-repository analyzer gaps), independent oracle comparison, and target post-state/boundary. Until Main/Owner explicitly authorizes target access, these gates remain unavailable rather than assumed.
7. Reseal the exact candidate, Git/index/diff boundary, report split, quarantine/runtime inventory, and target-access provenance.

## Verdict

REJECT

P6-D is not acceptable on the current candidate. There are two independent blocker classes: a local fail-closed correctness hole in nested authority identity validation, and the acknowledged absence of built-current-runtime/native Ladybug/built-reader/target/oracle/pre-post evidence. The coder's blocked disposition is truthful but does not satisfy the P6-D acceptance gate. No repair was performed. Raw report bytes, LF/CR counts, strict UTF-8/no-BOM result, and SHA-256 are measured externally after this content freeze and returned in the visible handoff to Main.
