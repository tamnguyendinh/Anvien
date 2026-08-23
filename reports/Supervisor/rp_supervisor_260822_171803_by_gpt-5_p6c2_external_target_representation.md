# Supervisor Report: Child 06 P6-C2 External Target Representation

Verdict: **REJECT**

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260822_171803_by_gpt-5_p6c2_external_target_representation.md`
- Review time: `2026-08-22 17:18:03 +07:00` through external seal.
- Reviewer: `gpt-5`, existing independent Child 06 Supervisor lane.
- Repository: `E:\Anvien`, branch `master`.
- HEAD: `8055f0a6860721e26462572e34469e0d708d4a52`.
- Parent / accepted P6-B commit: `5c1584fef7153a7a331c8fedd1ce64176ddc873d`.
- Open slice reviewed: P6-C2 only. P6-C3, P6-D, graph-health, target validation, project/package declaration lookup, and target access remained locked.
- Next owner: Main task `01a0289a-74dc-78f1-9106-18c4c98821ac`, or its officially transferred successor.

## Claim And Authority

Claim reviewed: final resolved P6-B authority rows materialize referenced-only deterministic `ExternalSymbol` targets without fake unavailable/excluded/mismatch targets or repository ownership pollution; all same-invariant persistence and reader consumers preserve provenance and non-editable semantics.

Authority read in full before judgment:

1. Root `AGENTS.md`.
2. Repo-local `working-rules`, `supervisor`, `Data-Integrity`, and `Edge-Case` skills.
3. All four current Child 06 ledgers.
4. P6-A declaration-universe decision and the accepted P6-B/P6-C1 Supervisor history needed for the inherited contract.
5. Current coder report, current source/diff, Git boundary, packaged runtime, graph, persistence, and reader reality.

The Data-Integrity and Edge-Case disciplines were applied inside this one Supervisor review: exact owner separation, replay/dedupe, collision and malformed-state behavior, persistence parity, and fail-closed reader behavior were inspected. No separate lane, report, commit, or subagent was opened.

Coder report identity was independently reproduced before source trust:

- path: `reports/coder/rp_coder_260822_162635_by_gpt-5_p6c2_external_target_representation.md`;
- `12,480` bytes / `198` LF / `0` CR;
- strict UTF-8, no BOM;
- SHA-256 `08BCEA399B3ECEED4721858176E3845B19A67DB376FD21138E8B9A017AEB05B0`;
- createdAt = lastWrite `2026-08-22T16:28:50.6827037+07:00`.

## Sole Blocking Invariant

The candidate implementation has no source/runtime blocker established by this review. Acceptance is rejected solely because this independent review no longer satisfies its mandatory no-`C:\` evidence boundary.

Concrete evidence:

1. A fresh production compile was launched with repo-local `GOCACHE`/`GOTMPDIR` and `GOPROXY=off`, but the compiler disclosed a dependency source under `C:\Users\TAM NGUYEN\go\pkg\mod\...`. The command was interrupted immediately and exited `1`. It is classified `ABORTED`, not a successful build.
2. On the first packaged fixture analyze, the PowerShell variable name `$home` collided case-insensitively with the read-only automatic `$HOME`. PowerShell emitted `Cannot overwrite variable HOME because it is read-only or constant`, then the remaining command proceeded. Consequently `ANVIEN_HOME` could resolve to the automatic C-home value for that run and could read or write its registry there.
3. This lane is forbidden to inspect or repair `C:\`, so the external registry state cannot be verified or restored inside the same authority boundary.

The later graph/runtime runs used an explicit `E:\Anvien\.tmp\supervisor-p6c2-review\anvien-home`, reproduced all claimed runtime facts, and all E-side transient artifacts were removed. That does not retroactively make the earlier prohibited access unknowable or compliant. A zero-trust acceptance report cannot assert the required boundary after this event.

Required fix direction is procedural, not a coder change:

- Main/Owner must dispose or reconcile any possible external registry side effect under separately granted authority.
- Re-review the unchanged P6-C2 candidate from a clean Supervisor execution whose Go/module/runtime homes are all explicitly E-local before invocation.
- Required re-review evidence: current exact source/diff and 22-path manifest; E-only compile/test or equivalent current artifact proof; E-isolated packaged resolved/noLib/default fixtures; native Ladybug and MCP/process/rename checks; one current root refresh plus detect; final cleanup and Git boundary; no C-path exposure.

No production, test, ledger, or coder repair is requested by this report.

## Preserved Technical Clearances

These source/runtime findings are independently established and should be retained unless candidate bytes change.

### Resolution and materialization: clear

- `internal/resolution/resolve.go:103-108` finalizes and validates all P6-B authority results before invoking external materialization.
- `recordTypeScriptLookup` records the source graph endpoint, scope, reference kind, target role, and target text alongside every handled P6-B site without changing the immutable authority-result DTO.
- `internal/resolution/external_symbol.go:50-104` requires an endpoint for every final result, checks endpoint/result cardinality, groups only `LookupResolved` rows, sorts semantic IDs, emits referenced targets, and emits no target for unavailable, excluded, or mismatch rows.
- `indexTypeScriptExternalSites` rejects incomplete endpoints, coalesces identical endpoint replays, and rejects a conflicting payload for one source-site identity.
- `typeScriptExternalGroup.add` rejects conflicting authority/version/proof/hash/owner payload for one semantic ID. `externalSymbolGraphID` derives one graph identity from the canonical P6-B semantic ID. A conflicting pre-existing graph node fails closed through exact node comparison.
- Node aggregation sorts names, meanings, profile/config hashes, source-site IDs, declaration ranges, and libraries. It preserves `semanticSymbolId`, optional `semanticOwnerId`, TypeScript/catalog/profile/config provenance, declaration library/range, `external=true`, `editable=false`, and `repositoryOwned=false` without `filePath`.
- External relations are only CALLS, ACCESSES, or USES according to the original reference kind. Every row retains stable site ID, status, proof kind, role/text/range, and the exact JSON-encoded P6-B authority result. Coalesced edges retain the sorted site inventory and one immutable authority proof per site.
- ReferenceIndex and resolved reference/call/access/type metrics are updated once per final unique resolved site. P6-B finalization remains the source-site conflict/counter-equality authority.
- No catalog-wide enumeration, target-name allowlist, `Promise`/`Math.max`/`Math.min` branch, runtime Node/`tsc`, network, package script, project/package lookup, synthetic `IMPORTS`, P6-C3 DTO, or graph-health behavior exists in the nine production owners.

### Production owner groups: clear

- `internal/scopeir/kinds.go`: adds only the dedicated `NodeExternalSymbol` label.
- `internal/resolution/external_symbol.go`: isolated referenced-only materializer and proof emitter as described above.
- `internal/resolution/emit.go`: adds transient endpoint inventory to the emitter; generic relationship coalescing remains the existing deterministic proof-preserving path.
- `internal/resolution/resolve.go`: captures endpoint facts at eligible call/access/type sites and invokes materialization only after final P6-B validation.
- `internal/lbugload/csv.go`: adds an explicit fixed column set and lossless row projection for every required external node property; existing relationship proof columns carry external edge proofs.
- `internal/lbugschema/schema.go`: adds the fixed node table, schema, and internal-source-to-external endpoint pairs; no external source ownership pair is introduced.
- `internal/processes/processes.go`: excludes a CALLS relation when either endpoint is external, preventing external process steps and terminals.
- `internal/mcp/context.go`: projects external identity, provenance, declaration, source-site, and non-editable ownership fields without a fake repository path.
- `internal/mcp/rename.go`: rejects the external label with stable reason `external_symbol_non_editable` before old-name comparison or edit collection.
- Untouched `internal/mcp/impact.go` uses the shared context semantic projection for both target and impacted rows; built runtime evidence below confirms it. Generic graph JSON, HTTP/Web, file-detail, graph-health, and shared later-slice DTOs remain untouched.

Current production identities recorded by this review:

| Path | Bytes | SHA-256 |
| --- | ---: | --- |
| `internal/scopeir/kinds.go` | 5,253 | `CAC3AD018B82E7B2950D0F13E24793D3CD84F708CF8E9533198C74C749B93B6E` |
| `internal/resolution/external_symbol.go` | 11,845 | `275683BD680107A19DCB40B782566C087A793A64CAD0341E040A43DBEBBE623C` |
| `internal/resolution/emit.go` | 26,822 | `DE88FA7717B3E1BBE9AB099616B5D43095A91C86C22706F17F0F5A5CD793C078` |
| `internal/resolution/resolve.go` | 38,706 | `7D1394E661E78679602429212E97F29A605147AFD2DD894E6864A3D15A808A69` |
| `internal/lbugload/csv.go` | 24,986 | `1B7D38FDE913A6A0F05033947B069650C01B8D2F55F64FDDD1D9DC406D894BCA` |
| `internal/lbugschema/schema.go` | 20,174 | `AF0B5727C9261FDB542B94A7DC1383022EA4F569F85FDEC8C816B1CC9B4F23FA` |
| `internal/processes/processes.go` | 14,942 | `F2DE1EB156D118075A3E5D6F62DE926F18E4F13B59000807ACF49F1E85DBD1B9` |
| `internal/mcp/context.go` | 22,267 | `3A0397ACF3A9E45771E4E49A1588AE634E331ADAC2DA464E295D8C6424A50528` |
| `internal/mcp/rename.go` | 8,945 | `B4EE673CFC1B01427B7D309BC1173DEB98D32FC7CFEB652B81E285E5822A2B14` |

### Test-owner inspection: clear

All eight changed/new test owners were read. Their coverage includes:

- final resolved-node identity/provenance and one authority proof per site at direct resolver and built `analyze.Run` boundaries;
- unavailable noLib producing no external node;
- identical-site dedupe and replay equality;
- multi-site edge coalescing without proof loss;
- exact CSV columns/rows and supported Function-to-ExternalSymbol pair;
- fixed Ladybug schema/table/pair inventory;
- exclusion of external source and target CALLS endpoints from process topology;
- context/impact non-editable provenance and rename rejection;
- tagged native Ladybug COPY/readback of the external node and relationship.

Fresh Go tests were not rerun after the C-path disclosure. Existing coder test/build results remain evidence inputs, not this verdict's independent acceptance basis.

## Packaged Runtime, Persistence, And Reader Evidence

The packaged binary was independently remeasured unchanged:

- `E:\Anvien\anvien\bin\anvien.exe`;
- `73,653,760` bytes;
- SHA-256 `7B14655DBB65472CE83FBAF0677294CCB235EC09A34DB3AB877F5C5C1B12F568`;
- packaged help exited `0`.

Two later E-isolated force analyses of the existing resolved TypeScript runtime fixture were byte-identical:

- graph `11` nodes / `10` relationships / `32,666` bytes;
- SHA-256 `CB308956ECD89DDB801BC37ED72D01AD73656D2D01EDC9C49C3A8ADB0CC7920A` on both runs;
- `3` unique referenced external nodes;
- `6` resolved source sites;
- `4` coalesced external edges;
- `6` exact authority proof rows;
- node source-site counts `[2,1,3]`;
- File / DEFINES / IMPORTS / CONTAINS ownership pollution `0 / 0 / 0 / 0`;
- invalid external ownership/property rows `0`.

Additional hostile outcomes:

- noLib fixture: `0` external nodes and `0` external edges;
- compiler-default fixture: only the resolved Promise type materialized; excluded/mismatch paths created no fake target;
- stable IDs, property arrays, declaration ranges, source-site rows, and edge proofs were inspected directly from graph JSON.

Native Ladybug query after a fresh fixture analyze returned exactly four external relation rows. Every row retained authority kind, TypeScript version, ready proof state, semantic ID, node/edge site counts, proof kind, resolution source, `external=True`, `editable=False`, and `repositoryOwned=False`. This proves the fixed schema/COPY/native readback path without fallback substitution.

Built packaged reader behavior was also independently exercised:

- context found `Math.max` as `ExternalSymbol`, returned empty repository `filePath`, complete semantic owner/authority/catalog/source-site fields, no processes, and the three explicit non-editable flags;
- upstream impact returned the same external target semantics, one direct repository caller, and zero affected processes/modules;
- rename dry-run returned `status=rejected`, reason `external_symbol_non_editable`, and no edit counts;
- graph relationships involving the external targets were only the four expected CALLS/USES rows, so process/external-source pollution was zero.

## Current Anvien Evidence

Repo-local packaged help was read before graph use. The final root graph refresh used explicit path `E:\Anvien`, `--force`, and an E-isolated Anvien home. It exited `0`:

- `2,041` scanned / `756` parsed code / `0` failed;
- `116,684` nodes / `162,208` relationships;
- `19,658` dependency edges;
- indexed/current commit `8055f0a6860721e26462572e34469e0d708d4a52`, stale=false.

The coder's earlier `2,039 / 756 / 0 / 116,657 / 162,181` snapshot remains truthful for its time boundary. Current `+2 files / +27 nodes / +27 relationships` is later coder-report/Main-handoff document state and was not normalized away.

Current central file-detail for `internal/resolution/external_symbol.go` is HIGH: `85` symbols / `27` inbound / `72` outbound / `76` local relationships / `109` unresolved / `21` linked tests. HIGH is retained as a blast-radius warning.

Current exact `NodeExternalSymbol` impact itself is LOW with zero upstream impacted symbols; the containing shared label file has `698` inbound references and `109` linked tests. The broader shared-label warning remains material even though the new constant's direct impact is narrow.

Fresh explicit-path detect exited `0` at CRITICAL risk:

- affected `44` symbols / `17` files;
- changed `202` symbols / `17` files;
- affected layers `api=16`, `backend=28`;
- affected areas `mcp=16`, `mixed=7`, `resolution=19`, `storage=2`;
- changed layers `api=11`, `backend=119`, `backend_test=53`, `docs=19`;
- changed areas `analyzer=32`, `documentation=19`, `mcp=11`, `providers=38`, `resolution=55`, `storage=47`;
- ResolutionGap delta `33` entities / `33` occurrences;
- actionability `analyzer_gap=20`, `non_actionable=13`;
- classifications `builtin=12`, `in_repo_unresolved=20`, `standard_library=1`;
- fact families `access=14`, `call=14`, `type-reference=5`;
- current Resolution Health total `0`;
- semantic app-layer and functional-area fields/sources complete for all `116,684` nodes.

CRITICAL is a full blast-radius warning, not an automatic defect or an acceptance substitute.

## Build, Regression, And Benchmark Disposition

- The review's fresh production compile is `ABORTED/exit 1` for the boundary reason above. It is not presented as successful evidence.
- No package script, dependency install, network, target access, stage, commit, push, branch action, UI/Playwright, or subagent occurred.
- The existing packaged artifact is current by exact hash and independently executed all affected runtime/persistence/reader paths described above.
- Known broad repository facts remain truthful: broad `go build ./...` is not successful because intentional language fixtures fail; broad `go test ./internal/...` retains seven repository-local Git/temp failures plus the existing C# and Dart ACCESS parity baselines; no such failure is relabeled as a P6-C2 success.
- Package manifests, lockfiles, `go.mod`, and `go.sum` have zero candidate diff.
- Packaged size delta over accepted P6-B is `+39,936` bytes (`+0.05425%`); no Owner threshold exists.
- Product/runtime benchmark facts independently reproduced here are the `3 / 6 / 4 / 6` representation counts, zero ownership/process pollution, byte-identical graph replay, and exact packaged binary size/hash.

## Exact Git, Candidate, Ledger, And Cleanup Boundary

Pre-report boundary after cleanup:

- branch `master`, HEAD `8055f0a...`, parent `5c1584f...`;
- status `57` paths = `17` tracked modifications + `40` untracked;
- exact P6-C2 candidate `22/22`, missing `0`, extra `0`;
- candidate composition: `9` production + `8` tests + `4` Child 06 ledgers + current coder report;
- candidate storage split: `17` tracked modifications + `5` untracked assets/report;
- protected outside candidate `35`: `30` pathname-only Main handoffs + `3` prior Supervisor reports + `2` older coder reports;
- the thirtieth Main handoff `reports/Investigation/rp_main_260822_163519_orchestration_rotation_handoff.md` appeared after the assignment snapshot, was observed only as a pathname, preserved, and is not candidate drift;
- index `0`;
- `git diff --check` exit `0`;
- coder report hash unchanged;
- P6-C2 remains `[ ]`, unstaged, and uncommitted;
- P6-C3/P6-D and target remain locked.

All E-side review artifacts are absent after cleanup:

- `.tmp/supervisor-p6c2-review`;
- `.tmp/p6c2`;
- `.tmp/p6c2-final-home`;
- fixture `.anvien` directories under resolved, noLib, and compiler-default P6-B runtime testdata.

The four ledgers are truthful for the pre-verdict state: they claim worker validation rather than acceptance, keep P6-C2 unchecked, retain pending review/commit evidence, and keep later slices ordered/locked. This report does not edit them.

This new Supervisor report is the only authorized durable repo file added by this review. It is review-only and outside the 22-path implementation candidate.

## Re-review Handoff

Preserved implementation clearances cover every touched production owner, all changed/new test owners, referenced-only identity/materialization, unavailable/excluded/mismatch absence, site proof/counter/replay behavior, Ladybug schema/CSV/native readback, process exclusion, MCP context/impact/rename, generic graph transport, forbidden-scope absence, packaged runtime, current graph, and exact detect counts.

Residual P6-C2 code invariant blockers established by this review: none.

The sole rejection blocker is the invalidated no-C review boundary. A clean E-only independent re-review may reuse the technical clearances above if all candidate bytes and the exact manifest remain unchanged; it must independently establish a compliant build/test/runtime evidence chain and final boundary before acceptance.
