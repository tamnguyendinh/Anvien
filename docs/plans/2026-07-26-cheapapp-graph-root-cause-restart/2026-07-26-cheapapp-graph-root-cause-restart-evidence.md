# Cheapapp Graph Root-Cause Restart Investigation Evidence Ledger

## Metadata

- Date: `2026-07-26`
- Plan: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-plan.md`
- Evidence: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-evidence.md`
- Benchmark: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-benchmark.md`
- Actual status: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/2026-07-26-cheapapp-graph-root-cause-restart-actual-status.md`

## Evidence Rules

This ledger records why each investigation claim is supported. It does not treat prior reports or agent output as authority. Every claim must link to current source, graph, command, runtime, data, or document evidence. Long metric tables belong in the benchmark ledger.

### Evidence ID Naming

Use stable IDs in the form `E<phase>-<item>-<kind><n>`. Exact IDs, not broad section names, must be referenced from the plan and actual-status file. `SRC` denotes independent source evidence, `GRAPH` graph evidence, `CMP` comparison, `CAUSE` first-divergence evidence, `TARGET` target-boundary evidence, `REVIEW` Supervisor evidence, and `REPORT` durable report output.

## E0 - P0 Evidence

Matching plan item(s): `P0-A`

- `E0-P0A-RULE1`: AGENTS.md and selected planner/debugging/supervisor skill instructions read before investigation actions.
- `E0-P0A-HELP1`: `anvien --help` output captured.
- `E0-P0A-HELP2`: `anvien analyze --help`, `anvien index --help`, and `anvien list --help` output captured.
- `E0-P0A-TARGET1`: target baseline HEAD/status/tree metadata and read-only boundary capture.
- `E0-P0A-OUTPUT1`: supported external-repository/index output location proven from command/source/config inspection.
- `E0-P0A-GRAPH1`: fresh target graph identity and metadata after safe analysis.
- `E0-P0A-TARGET2`: post-analysis target status/tree comparison proving no target artifact or source mutation.
- `E0-P0A-BLOCK1`: exact blocker evidence if safe direct analysis cannot be established.

Historical P0 evidence before owner correction (superseded):

- `E0-P0A-RULE1`: rules and selected skills were read before action.
- `E0-P0A-HELP1`: root Anvien help captured.
- `E0-P0A-HELP2`: analyze/index/list help captured.
- `E0-P0A-OUTPUT1`: README/source inspection proves default repo-local `<repo>\\.anvien` storage and no alternate output root in the reviewed CLI surface.
- `E0-P0A-TARGET1`: read-only check proves `E:\\cheapapp` is absent; `E:\\cheapapp.org` is a distinct existing directory and was not substituted.
- `E0-P0A-BLOCK1`: `reports/Investigation/20260726_142153_p0_direct_target_boundary.md` records the path and safety blocker.
- `E0-P0A-TARGET2`: no target analysis command was run; no target write occurred.
- `E0-P0A-REVIEW1`: Supervisor PASS for the narrow P0 blocker claim in `reports/Supervisor/rp_supervisor_260726_142537_by_gpt-5-6-sol_p0_target_boundary.md`.
- `E0-P0A-SIDE1`: source-verified target write surfaces in the normal analyze/postrun path (`.anvien`, `.gitignore`, metadata, AI context/skills).
- `E0-P0A-SIDE2`: source-verified registry normalization; custom external storage paths are not a supported contract.

Historical P0 decision: superseded by `E0-P0A-R2-OWNER1`. The corrected P0 baseline is complete for `E:\cheapapp.org`; its graph/index stays at `E:\cheapapp.org\.anvien`, while investigation artifacts stay under `E:\Anvien`.

## E1 - P1 Evidence

Current P1-A evidence:

- `E1-P1A-GRAPH1`: fresh graph metadata/hash and target commit from `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart`.
- `E1-P1A-GRAPH2`: analyze benchmark with `42.1059032 s` total duration, 1,359 scanned files, 887 parsed code files, 84,807 nodes, and 114,125 relationships.
- `E1-P1A-FILE1`: source Git manifest and extension-filtered comparison (`895` source TS/JS paths, `887` graph File nodes, `8` missing, `0` graph-only).
- `E1-P1A-FILE2`: eight source hashes and no Git ignore matches.
- `E1-P1A-CMP1`: exact eight `file-detail` not-found results.
- `E1-P1A-REPORT1`: `reports/Investigation/20260726_144919_p1a_file_inventory.md`.

P1-A status: bounded file omission finding recorded and scanner first divergence confirmed in P2-A. No global file-coverage claim is made.

Current P1-B evidence:

- `E1-P1B-GRAPH1`: fresh target graph identity/hash/counts and up-to-date status.
- `E1-P1B-SRC1`: independent TypeScript AST walk of the three fixed target files: six array-pattern bindings, two `time` declarations, two `now` declarations, and 21 export declarations.
- `E1-P1B-CMP1`: raw graph comparison: six pattern names absent, one graph node for each same-name pair, all 21 export names present as definitions but zero `exported`/`visibility` properties.
- `E1-P1B-ORACLE1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle.mjs` and `p1b-identity-oracle-output.json`.
- `E1-P1B-REPORT1`: `reports/Investigation/20260726_155807_p1b_ts_identity.md`.
- `E1-P1B-REVIEW1`: independent rerun/source/file-detail Supervisor PASS in `reports/Supervisor/rp_supervisor_260726_161811_by_gpt-5-6-sol_p1b_ts_identity.md`.

P1-B status: bounded extraction/identity/visibility mismatches and their P2-A first divergences are Supervisor-accepted. No global prevalence or remediation claim.

Current P1-C evidence:

- `E1-P1C-GRAPH1`: `anvien status` and graph SHA-256 bind the comparison to target HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, analyzed at `2026-07-26T07:49:43Z`.
- `E1-P1C-SRC1`: exact `Promise` source at `read-admin-commercial-config.ts:13` and `Math.max`/`Math.min` source at `email-operations-observability.ts:191`.
- `E1-P1C-ANVIEN1`: fresh `file-detail` classifies all three ambient sites as `in_repo_unresolved`, `analyzer_gap`, `unresolved_local_binding`.
- `E1-P1C-ORACLE1`: `p1c-typescript-oracle.mjs` and `p1c-typescript-oracle-output.json` resolve all three sites to TypeScript `lib.es*.d.ts` declarations with zero line diagnostics.
- `E1-P1C-SRC2`: exact declaration → barrel line 21 → two consumer import/call chains for `readAdminCommercialConfig`.
- `E1-P1C-ANVIEN2`: `p1c-anvien-observations.json` records that both consumers only import the barrel File, lack `USES`/`CALLS` edges to the declaration, and expose two false gaps per call site.
- `E1-P1C-ORACLE2`: `p1c-barrel-oracle.mjs` and `p1c-barrel-oracle-output.json` resolve both consumer aliases to `read-admin-commercial-config.ts:10` with zero line diagnostics.
- `E1-P1C-CMP1`: five of five fixed compiler-bound sites are exposed as unresolved by Anvien; classification is bounded `wrong`.
- `E1-P1C-REPORT1`: `reports/Investigation/20260726_152726_p1c_resolution_binding.md`.

P1-C status: bounded ambient and barrel-binding mismatches and their P2-B first divergences are Supervisor-accepted. No global resolver claim or remediation is made.

Current P1-D evidence:

- `E1-P1D-KEY1`: canonical counting keys distinguish physical occurrence, exact `sourceSiteId`, raw gap node, and raw relationship.
- `E1-P1D-RAW1`: independent raw graph scan finds one `ResolutionGap` node and one incoming gap edge for each of seven canonical IDs, zero materialized `SourceSite` nodes, zero explicit nulls, and no raw `step` key.
- `E1-P1D-PROJ1`: expanded `file-detail` captures return one matching sample per selected ID.
- `E1-P1D-PROJ2`: supported Cypher gap-node and `CodeRelation` captures return one row per selected ID; scalar/evidence normalization and default `step` are recorded.
- `E1-P1D-CMP1`: raw/projection cardinality and field comparison.
- `E1-P1D-REPORT1`: `reports/Investigation/20260726_154047_p1d_projection_parity.md`.
- `E1-P1D-REVIEW1`: bounded Supervisor PASS in `reports/Supervisor/rp_supervisor_260726_154849_by_gpt-5-6-sol_p1d_projection_parity.md`.

P1-D status: bounded projection parity accepted; no global projection or resolver conclusion.

Current P1-E evidence:

- `E1-P1E-SRC1`: independent source counts for wrapper and bounded-helper callers.
- `E1-P1E-CMD1`: `context` output for `readAdminCommercialConfig`.
- `E1-P1E-CMD2`: upstream `impact` output for the wrapper UID.
- `E1-P1E-CMD3`: `context` output for `boundedReadLimit`.
- `E1-P1E-CAP1`: bounded empty process arrays recorded as observed command output, not completeness evidence.
- `E1-P1E-CMP1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1e-command-captures.json`.
- `E1-P1E-REPORT1`: `reports/Investigation/20260726_154558_p1e_command_projection.md`.

P1-E status: bounded command/projection comparison complete; no global process/semantic claim.

## E0-R2 - Owner Correction and P0 Reopen

Matching plan item(s): `P0-A`

- `E0-P0A-R2-OWNER1`: owner clarification that the target is `E:\cheapapp.org` and its graph/index remains under `E:\cheapapp.org\.anvien`.
- `E0-P0A-R2-BOUNDARY1`: corrected artifact boundary: reports, ledgers, probes, benchmarks, and causal analysis under `E:\Anvien`; no repository copy.
- `E0-P0A-R2-REOPEN1`: P0 path blocker superseded and reopened against the corrected target; fresh baseline still required.

Matching plan item(s): `P1-A`, `P1-B`, `P1-C`, `P1-D`, `P1-E`

- `E1-P1A-GRAPH1..E1-P1A-GRAPH3`: fresh graph metadata, hash, inventory, and benchmark.
- `E1-P1A-FILE1..E1-P1A-FILE2`: independent target file inventory and graph File inventory.
- `E1-P1A-CMP1`: bidirectional file-set comparison and discrepancy classification.
- `E1-P1A-REPORT1`: P1-A Supervisor report.
- `E1-P1B-SRC1..E1-P1B-GRAPH2`: source/AST and graph declaration/identity evidence.
- `E1-P1B-CMP1`: extraction/identity comparison.
- `E1-P1B-REPORT1`: P1-B Supervisor report.
- `E1-P1C-SRC1..E1-P1C-GRAPH2`: source binding and resolver/edge evidence.
- `E1-P1C-CMP1`: resolution comparison.
- `E1-P1C-REPORT1`: P1-C Supervisor report.
- `E1-P1D-KEY1..E1-P1D-PROJ2`: projection comparison keys and read-only raw/projection evidence.
- `E1-P1D-CMP1`: multiplicity/nullability/cardinality comparison.
- `E1-P1D-REPORT1`: P1-D Supervisor report.
- `E1-P1E-CMD1..E1-P1E-DERIVED2`: command and derived-output evidence.
- `E1-P1E-CMP1`: upstream-versus-downstream comparison.
- `E1-P1E-REPORT1`: P1-E Supervisor report.

## E2 - P2 Evidence

Current P2-A evidence:

- `E2-P2A-SRC1`: hardcoded ignore segments and any-segment matching in `internal/ignore/constants.go:5-78,219-235`.
- `E2-P2A-SRC2`: matcher ordering in `internal/ignore/matcher.go:41-56` and directory pruning in `internal/scanner/scan.go:55-90`.
- `E2-P2A-IMPACT1`: `ShouldIgnorePath` exact-symbol impact (`LOW`, one direct affected file; containing file risk `high`).
- `E2-P2A-IMPACT2`: `WalkRepositoryPaths` exact-symbol impact (`CRITICAL`, five impacted symbols, four modules, 39 processes; scanner file risk `high`).
- `E2-P2A-CAUSE1`: real-target ignore probe returns ignored for all eight missing paths before candidacy.
- `E2-P2A-REPORT1`: `reports/Investigation/20260726_145712_p2a_scanner_root_cause.md`.
- `E2-P2A-TSIR1`: direct real-target parser/extractor probe in `.tmp/cheapapp-graph-root-cause-restart/p2a-identity/` proves the six array binding names are absent before graph emission, while both `time` and both `now` definitions remain distinct in ScopeIR.
- `E2-P2A-IDENTITY1`: `internal/resolution/indexes.go:814-824` removes range/scope from graph definition identity and `internal/graph/types.go:96-104` replaces duplicate node IDs.
- `E2-P2A-VIS1`: `internal/providers/tsjs/definitions.go:100-111` leaves `DefinitionFact.Visibility` empty and `internal/resolution/emit.go:171-174` emits only non-empty visibility.
- `E2-P2A-IMPACT3`: fresh self-graph impact: `graphIDForDef` CRITICAL (3 modules/18 processes); `Graph.AddNode` CRITICAL (16 direct symbols/6 modules/82 processes); collector owners return LOW with a documented dispatch-graph limitation.
- `E2-P2A-REPORT2`: `reports/Investigation/20260726_161748_p2a_ts_identity_root_cause.md`.
- `E2-P2A-REVIEW1`: independent Supervisor PASS for the TypeScript identity/extractor trace, including fresh AST/ScopeIR reruns and the graph-ID/visibility source trace, in `reports/Supervisor/rp_supervisor_260726_164714_by_gpt-5-6-sol_p2a_ts_identity.md`; the scanner report remains a separate bounded causal report.

P2-A status: bounded scanner and TS identity/extractor first divergences reported and identity trace Supervisor-accepted; no remediation performed. The scanner and identity findings remain bounded and carry no global accuracy claim.

Matching plan item(s): `P2-A`, `P2-B`, `P2-C`

- `E2-P2A-SRC1..E2-P2A-IMPACT2`: Anvien scanner/extractor/identity source and impact evidence.
- `E2-P2A-CAUSE1`: bounded reproduction of first extraction/identity divergence.
- `E2-P2A-REPORT1`: P2-A causal report.
- `E2-P2B-SRC1..E2-P2B-IMPACT2`: resolver/module source and impact evidence.
- `E2-P2B-CAUSE1`: bounded reproduction of first resolver divergence.
- `E2-P2B-REPORT1`: P2-B causal report.
- `E2-P2C-SRC1..E2-P2C-IMPACT2`: projection/command/derived source and impact evidence.
- `E2-P2C-CAUSE1`: bounded reproduction of first consumer divergence.
- `E2-P2C-REPORT1`: P2-C causal report.

Current P2-B evidence (owner slice; Supervisor PASS; bounded record accepted; global resolver scope unresolved):

- `E2-P2B-BOUNDARY1`: fresh target graph remains bound to HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, SHA-256 `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`, 84,807 nodes / 114,125 relationships; target worktree pre-existing changes preserved and no probe artifact written there.
- `E2-P2B-TARGET1`: fresh target `file-detail` captures for declaration, barrel, both consumers, and email observability; all are parsed and stale=false, with the exact fixed-site unresolved rows.
- `E2-P2B-ORACLE1`: rerun TypeScript 5.9.3 ambient oracle; `Promise`, `Math.max`, and `Math.min` resolve to `lib.es*.d.ts`; target diagnostics count is zero.
- `E2-P2B-ORACLE2`: rerun TypeScript barrel oracle; both consumer aliases resolve to `read-admin-commercial-config.ts:10`; line diagnostics are empty.
- `E2-P2B-IR1`: direct in-memory parse/extract of the real five target files; Promise annotation, Math member calls, barrel `ImportReexport`, consumer imports, and call ranges are present.
- `E2-P2B-RESOLVE1`: actual barrel-chain versus direct-declaration in-memory control: imports `10/10`, import-use edges `1/3`, resolved calls `37/39`, unresolved refs `542/540`; only consumer import targets were changed in memory.
- `E2-P2B-GRAPH1`: raw selected target graph records showing consumer→barrel `IMPORTS`, barrel→declaration `IMPORTS` + `USES`, absent consumer declaration `CALLS`/`USES`, and exact Promise/Math/call ResolutionGap nodes/edges.
- `E2-P2B-CAUSE-A1`: source trace through `parseFiles`, `buildWorkspace`, `resolveTypeAnnotation`, `resolveCall`, `resolveReceiverType`, and unresolved emission; TypeScript ambient declarations are absent from the resolver workspace.
- `E2-P2B-CAUSE-B1`: source trace through `resolveImports` and `resolveImportedDef`; physical target-file lookup does not follow a barrel's re-export binding, so the consumer `TargetDef` is nil and its scope binding is skipped.
- `E2-P2B-IMPACT1`: self-graph `file-detail` and upstream impact for every candidate owner, preserving HIGH/CRITICAL workflow warnings and LOW collector traversal gaps.
- `E2-P2B-REPORT1`: `reports/Investigation/20260726_155230_p2b_resolver_root_cause.md`.
- `E2-P2B-REVIEW1`: fresh compiler-oracle, direct real-source differential, source, raw graph, and impact Supervisor PASS in `reports/Supervisor/rp_supervisor_260726_162059_by_gpt-5-6-sol_p2b_resolver_root_cause.md`.

P2-B status: bounded ambient and barrel first divergences independently reproduced and Supervisor-accepted; no remediation performed.

Current P2-C evidence (owner slice; Supervisor PASS; bounded record accepted; global process/semantic scope unresolved):

- `E2-P2C-BOUNDARY1`: target remains HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, graph SHA-256 `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`, 84,807 nodes / 114,125 relationships, and graph last-write `2026-07-26T14:49:43.8894011+07:00`; no target analyze or write in P2-C.
- `E2-P2C-TARGET1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2c\target_command_boundary_probe.json` records exact selected relationships, 662 total Process nodes, 2,761 total `STEP_IN_PROCESS` edges, and zero selected symbol/file process edges.
- `E2-P2C-CMD1`: current source reconciliation of P1-E context/impact captures proves commands scan present relationships directly and do not synthesize missing barrel-call edges.
- `E2-P2C-PROJ1`: current source reconciliation of P1-D raw/Cypher/file-detail captures preserves seven-of-seven canonical rows and identifies bounded projection shape changes.
- `E2-P2C-SRC1`: source trace through `contextNeighborhood`, `runImpactBFSProfiled`, `impactAffectedProcesses`, `Builder.buildUnresolved`, and compact unresolved-row encoding.
- `E2-P2C-SRC2`: `internal/lbugload/csv.go:372-387` maps nil relationship step to `0`; `internal/lbugnative/native_ladybugdb.go:132-165` converts every returned scalar to string.
- `E2-P2C-DERIVED1`: in-memory process control removes existing derived nodes, reruns process detection, and compares actual 3,771 calls / 662 processes / 2,761 steps with a +2-call control at 3,773 calls / 662 processes / 2,761 steps; only calls considered change, while process/step counts and selected memberships remain unchanged. This exact control supports no broader process property.
- `E2-P2C-DERIVED2`: five fresh Supervisor reruns under `p2c\supervisor\process-run-1.json` through `process-run-5.json` reproduce the same no-change result and target graph hash.
- `E2-P2C-REJECT1`: Supervisor report `reports/Supervisor/rp_supervisor_260726_163219_by_gpt-5-6-sol_p2c_projection_command.md` rejected the stale derived-process interpretation; this revision removes the stale values and unsupported broader-property wording.
- `E2-P2C-IMPACT1`: refreshed Anvien self-graph file-detail and impact captures under `p2c\owner-evidence`; CRITICAL warnings preserved for context/impact/file-detail/CSV/process owners and LOW warnings for native tuple conversion.
- `E2-P2C-TEST1`: targeted existing tests passed for MCP context/impact projection, file-context row/cardinality handling, and relationship CSV projection.
- `E2-P2C-REPORT1`: `reports/Investigation/20260726_161241_p2c_projection_command_root_cause.md`.
- `E2-P2C-REVIEW2`: corrected-report Supervisor PASS after the historical stale-control rejection, in `reports/Supervisor/rp_supervisor_260726_163556_by_gpt-5-6-sol_p2c_projection_command.md`; the five reruns and corrected `3,771/662/2,761/0/0` versus `3,773/662/2,761/0/0` control remain the accepted bounded measurement.

P2-C status: bounded command traversal is faithful to present edges; file-detail cardinality is faithful with an explicit reduced DTO; database/Cypher projection has bounded nil-to-zero and scalar-to-string representation changes. Global derived-process/semantic completeness remains unresolved. Supervisor PASS is recorded; no remediation performed and the accepted-bounded plan does not extend beyond these fixed cases.

## E3 - P3 Evidence

Matching plan item(s): `P3-A`, `P3-B`

- `E3-P3A-SYN1`: cross-slice causal matrix in `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md`, with every included row linked to exact P1/P2 evidence IDs.
- `E3-P3A-SYN2`: classification ledger separating confirmed-wrong bounded causes, bounded-valid downstream behavior, representation non-parity, and unresolved/authority-blocked questions in the same synthesis report.
- `E3-P3A-LEDGER1`: prior zero-trust ledger audit `reports/Supervisor/rp_supervisor_260726_164403_by_gpt-5-6-sol_plan_ledger_reconciliation.md` and its required stale-status corrections.
- `E3-P3A-LEDGER2`: current plan/evidence/benchmark/actual-status refresh R10/R11; active pending-review contradictions and stale process wording were removed while historical R6/R7 snapshots remain explicitly superseded before the final PASS.
- `E3-P3A-REPORT1`: `reports/Supervisor/260726_165317_graph_root_cause_causal_synthesis.md`.
- `E3-P3A-TARGET1`: `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p3a-target-boundary.json`, confirming target HEAD/hash/inventory, expected repo-local graph, no plan-specific target artifact paths, and removed fixture.
- `E3-P3B-REVIEW1`: final zero-trust Supervisor checks: current target/hash/source boundary, all ten synthesis rows, all current P2 PASS records, corrected process/parity reruns, ledger contradiction sweep, production-source diff, and fresh targeted tests.
- `E3-P3B-REPORT1`: `reports/Supervisor/rp_supervisor_260726_170048_by_gpt-5-6-sol_final_bounded_investigation.md` — PASS for the bounded investigation record only.

P3 status: final Supervisor PASS; plan is `accepted-bounded` for evidence purposes only. Global graph accuracy, product semantics, remediation, performance, and implementation closure remain unresolved/out of scope.

## Closure Evidence

Closure evidence for this read-only investigation: `E3-P3B-REVIEW1` and `E3-P3B-REPORT1`. The cleanup audit found no dead plan-created artifact; historical rejection reports and raw probes were retained for traceability. Final `anvien detect-changes` and commit evidence are `N/A` because no production implementation slice changed Anvien source; this does not imply graph correctness or remediation closure.
