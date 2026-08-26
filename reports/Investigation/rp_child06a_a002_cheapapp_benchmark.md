# Child 06A A002 Cheapapp Candidate Measurement

Status: `MEASUREMENT_A002_CHEAPAPP_INCOMPLETE_MISSING_17_CHILD_INSTRUMENTATION`.

This is a measurement handoff for `E:\cheapapp.org` only. The sole A002 candidate process completed successfully and emitted all `30/30` top-level operation rows, but `candidate\benchmark.json` has no root child-measurement collection and its `resolution` object contains semantic counters only. Candidate resolution-child coverage is therefore exactly `0/17`, not `1/17`. D001 `resolve_calls` and D002-D017 were not measured. No child duration, denominator, rank, share, delta, sum, residual, exclusive-interval count, overlap count, or CPU profile is inferred from parent time, semantic counters, or the accepted A001 capture.

All A002 candidate timings below are raw, unaccepted measurements because the required same-work 17-child packet is incomplete. This report makes no `KEEP`, rejection, rollback, Supervisor, or baseline-promotion decision.

## Run and identity

| Field | Exact value |
|---|---|
| Target | `E:\cheapapp.org` |
| Target HEAD | `a869876ab6262dacde6cd5d432d099a91852a646` |
| Source HEAD authority | `cc420e7ad719d90dc4b2d9991be0249e8d648daa` plus the exact authorized uncommitted A002 diff |
| Runtime version | `1.2.8` |
| Source executable | `E:\Anvien\anvien\bin\anvien.exe` |
| Frozen executable | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\anvien-a002-cheapapp-candidate.exe` |
| Executable bytes / SHA-256 | `73,767,936` / `11895442CA4A5708FA439A787BC541102BE7144856315BCD17E9A7B772658AED` |
| Frozen native DLL | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\lbug_shared.dll` |
| Native DLL bytes / SHA-256 | `20,230,656` / `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |
| Exact command | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\anvien-a002-cheapapp-candidate.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\benchmark.json --benchmark-label child06a-a002-cheapapp-candidate` |
| Launch count | `1` |
| PID | `12756` |
| Process start | `2026-08-25T21:40:32.8136848+07:00` |
| Process end | `2026-08-25T21:43:09.4112295+07:00` |
| Exit code | `0` |
| Process wall | `156.946718300 s` |
| Process CPU | `110.906250000 s` |
| User / kernel CPU | `96.718750000 s` / `14.187500000 s` |
| Graph location | Target-local `E:\cheapapp.org\.anvien` |
| Environment/cache identity | Real target-local graph with `--force`; no separate machine/environment fingerprint was instrumented, so that subfield is unavailable |

Candidate identity matched the assigned runtime before launch. The exact four source identities also matched the Coder authority:

| File | Bytes | SHA-256 |
|---|---:|---|
| `E:\Anvien\internal\graphhealth\diagnostics.go` | 28,325 | `AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8` |
| `E:\Anvien\internal\resolution\emit.go` | 27,508 | `73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060` |
| `E:\Anvien\internal\resolution\outcome.go` | 19,921 | `02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E` |
| `E:\Anvien\internal\graphhealth\diagnostics_test.go` | 10,030 | `58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA` |

The following ten pre-existing target status rows were present before the run and remained unchanged after it:

```text
 M .dockerignore
 M reports/QA/2026-06-15-admin-artifact-delete-confirmation/docker/real-chrome-desktop-admin-artifact-delete-confirmation.png
 M reports/QA/2026-06-15-admin-artifact-delete-confirmation/docker/real-chrome-desktop-admin-artifact-delete-disabled.png
 M reports/QA/2026-06-15-runtime-landing-header-footer-i18n/docker/real-chrome-desktop-apps-empty.png
 M reports/QA/2026-06-15-runtime-landing-header-footer-i18n/docker/real-chrome-desktop-download-empty.png
 M reports/QA/2026-06-15-runtime-landing-header-footer-i18n/docker/real-chrome-desktop-landing-en-review.png
 M reports/QA/2026-06-15-runtime-landing-header-footer-i18n/docker/real-chrome-mobile-landing-en-review.png
?? reports/QA/rp_qa_260619_145255_by_gpt5_user_admin_lifecycle.md
?? reports/QA/rp_qa_260619_152926_by_gpt5_release_catalog_mutation.md
?? reports/problem/images/
```

## Available boundary measurements

The before column is the accepted A001 optimized Cheapapp artifact. It was not rerun or replaced. Candidate deltas are arithmetic on available raw boundaries only and are not acceptance evidence.

| Boundary | Accepted A001 before | Raw A002 candidate | Candidate - before |
|---|---:|---:|---:|
| D001 `resolve_calls` | `25.045225300 s`; calls `27,890`; files `887`; rank `3`; parent share `13.576041%` | **unavailable / not measured** | **unavailable** |
| B1-P1A-OP001 resolution parent | `184.481061700 s` | `22.664132600 s` | `-161.816929100 s (-87.714656%)` |
| Analyzer total | `274.474620900 s` | `120.335153100 s` | `-154.139467800 s (-56.158004%)` |
| Process wall | `279.105934600 s` | `156.946718300 s` | `-122.159216300 s (-43.768047%)` |
| Process CPU | `279.968750000 s` | `110.906250000 s` | `-169.062500000 s (-60.386204%)` |

## Complete 30/30 top-level operation packet

Shares use each same-run process wall as denominator: A001 `279.105934600 s`, A002 `156.946718300 s`. Ranks are descending duration within the same 30-row run; zero-duration ties retain the established row order. All denominators match exactly between A001 and A002. Every A002 value in this table is raw and unaccepted because the child packet is missing.

| Row | Boundary | Operation | A001 before | A002 candidate | Delta | Before share | Candidate share | Before rank | Candidate rank | Same-run denominator |
|---|---|---|---:|---:|---:|---:|---:|---:|---:|---|
| B1-P1A-OP001 | analyzer_internal | resolution | 184.481061700 s | 22.664132600 s | -161.816929100 s (-87.714656%) | 66.097148% | 14.440654% | 1 | 3 | {"runs":1} |
| B1-P1A-OP002 | analyzer_internal | db_load | 40.950769400 s | 51.353203700 s | +10.402434300 s (+25.402293%) | 14.672124% | 32.720151% | 2 | 1 | {"runs":1} |
| B1-P1A-OP003 | analyzer_internal | parse | 12.840209000 s | 10.240434800 s | -2.599774200 s (-20.247133%) | 4.600479% | 6.524784% | 4 | 5 | {"runs":1} |
| B1-P1A-OP004 | analyzer_internal | graph_snapshot | 18.124932500 s | 15.363476600 s | -2.761455900 s (-15.235675%) | 6.493926% | 9.788976% | 3 | 4 | {"nodes":95762,"relationships":129716} |
| B1-P1A-OP005 | analyzer_internal | semantic_enrichment | 5.503521100 s | 5.168669300 s | -0.334851800 s (-6.084319%) | 1.971840% | 3.293264% | 6 | 7 | {"runs":1} |
| B1-P1A-OP006 | analyzer_internal | db_runner_resolve | 1.488161300 s | 1.080884300 s | -0.407277000 s (-27.367799%) | 0.533189% | 0.688695% | 9 | 9 | {"runners":1} |
| B1-P1A-OP007 | analyzer_internal | cross_file_binding | 0.809683000 s | 0.624347600 s | -0.185335400 s (-22.889872%) | 0.290099% | 0.397809% | 10 | 11 | {"runs":1} |
| B1-P1A-OP008 | cli_outer | ai_context | 2.066776600 s | 35.585123000 s | +33.518346400 s (+1621.769203%) | 0.740499% | 22.673378% | 7 | 2 | {"baseSkills":49,"generatedFiles":4} |
| B1-P1A-OP009 | analyzer_internal | analyzer_orchestration | 0.783394400 s | 0.540456600 s | -0.242937800 s (-31.010919%) | 0.280680% | 0.344357% | 11 | 13 | {"analyzeRuns":1} |
| B1-P1A-OP010 | analyzer_internal | processes | 0.626072300 s | 0.760794600 s | +0.134722300 s (+21.518649%) | 0.224314% | 0.484747% | 12 | 10 | {"runs":1} |
| B1-P1A-OP011 | analyzer_internal | scan | 5.869840300 s | 10.032886200 s | +4.163045900 s (+70.922643%) | 2.103087% | 6.392543% | 5 | 6 | {"runs":1} |
| B1-P1A-OP012 | cli_outer | file_projection | 0.473783400 s | 0.279681500 s | -0.194101900 s (-40.968489%) | 0.169750% | 0.178202% | 14 | 14 | {"files":1359,"nodes":95762,"relationships":129716} |
| B1-P1A-OP013 | analyzer_internal | routes | 1.867265600 s | 1.355724400 s | -0.511541200 s (-27.395203%) | 0.669017% | 0.863812% | 8 | 8 | {"runs":1} |
| B1-P1A-OP014 | analyzer_internal | documents | 0.128885400 s | 0.097070400 s | -0.031815000 s (-24.684720%) | 0.046178% | 0.061849% | 19 | 18 | {"runs":1} |
| B1-P1A-OP015 | analyzer_internal | db_runner_close | 0.253016400 s | 0.218863800 s | -0.034152600 s (-13.498176%) | 0.090652% | 0.139451% | 16 | 15 | {"runners":1} |
| B1-P1A-OP016 | analyzer_outer | analyze_setup | 0.216351500 s | 0.052904800 s | -0.163446700 s (-75.546830%) | 0.077516% | 0.033709% | 17 | 20 | {"analyzeRuns":1} |
| B1-P1A-OP017 | analyzer_internal | tools | 0.535691500 s | 0.599759600 s | +0.064068100 s (+11.959887%) | 0.191931% | 0.382142% | 13 | 12 | {"runs":1} |
| B1-P1A-OP018 | analyzer_internal | communities | 0.039965300 s | 0.034361500 s | -0.005603800 s (-14.021664%) | 0.014319% | 0.021894% | 21 | 22 | {"runs":1} |
| B1-P1A-OP019 | cli_outer | registry_meta | 0.355614000 s | 0.099024900 s | -0.256589100 s (-72.153824%) | 0.127412% | 0.063095% | 15 | 16 | {"repositories":1} |
| B1-P1A-OP020 | cli_outer | cli_preparation | 0.056014100 s | 0.046499100 s | -0.009515000 s (-16.986794%) | 0.020069% | 0.029627% | 20 | 21 | {"commands":1} |
| B1-P1A-OP021 | cli_outer | output_publication | 0.036380300 s | 0.001993200 s | -0.034387100 s (-94.521211%) | 0.013035% | 0.001270% | 22 | 26 | {"outputs":1} |
| B1-P1A-OP022 | analyzer_internal | orm | 0.138500100 s | 0.090508100 s | -0.047992000 s (-34.651239%) | 0.049623% | 0.057668% | 18 | 19 | {"runs":1} |
| B1-P1A-OP023 | cli_outer | cli_startup | 0.022272100 s | 0.018961900 s | -0.003310200 s (-14.862541%) | 0.007980% | 0.012082% | 23 | 23 | {"commands":1} |
| B1-P1A-OP024 | analyzer_internal | structure | 0.012747400 s | 0.008391000 s | -0.004356400 s (-34.174812%) | 0.004567% | 0.005346% | 24 | 24 | {"runs":1} |
| B1-P1A-OP025 | analyzer_internal | graph_compact | 0.012060400 s | 0.099020500 s | +0.086960100 s (+721.038274%) | 0.004321% | 0.063092% | 25 | 17 | {"inputNodes":95762,"inputRelationships":129716} |
| B1-P1A-OP026 | analyzer_internal | mro | 0.000019700 s | 0.002167500 s | +0.002147800 s (+10902.538071%) | 0.000007% | 0.001381% | 27 | 25 | {"runs":1} |
| B1-P1A-OP027 | analyzer_internal | benchmark_write | 0.008824100 s | 0.000000000 s | -0.008824100 s (-100.000000%) | 0.003162% | 0.000000% | 26 | 27 | {"artifacts":1} |
| B1-P1A-OP028 | analyzer_internal | cobol | 0.000000000 s | 0.000000000 s | 0.000000000 s (0.000000%) | 0.000000% | 0.000000% | 28 | 28 | {"runs":1} |
| B1-P1A-OP029 | cli_outer | memory_profile | 0.000000000 s | 0.000000000 s | 0.000000000 s (0.000000%) | 0.000000% | 0.000000% | 29 | 29 | {"profiles":0} |
| B1-P1A-OP030 | cli_outer | cpu_profile_completion | 0.000000000 s | 0.000000000 s | 0.000000000 s (0.000000%) | 0.000000% | 0.000000% | 30 | 30 | {"profiles":0} |

## Missing 17/17 resolution-child packet

The A001 columns below identify the accepted before authority. Every A002 child field is `unavailable / not measured`; even when a similar count appears in semantic counters, it is not treated as a child denominator. Consequently no child delta is valid.

| Child row | Child | Accepted A001 duration / share / rank | Accepted A001 denominator | A002 duration / share / rank / denominator / delta | Exact child owner and full call path |
|---|---|---|---|---|---|
| B2-P2A-A001-D001 | resolve_calls | 25.045225300 s / 13.576041% / 3 | {"calls":27890,"files":887} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:resolveCall:92-94`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each call in ir.Calls -> resolution.resolveCall` |
| B2-P2A-A001-D002 | resolve_accesses | 91.241633200 s / 49.458536% / 1 | {"accesses":26042,"files":887} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:resolveAccess:95-97`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each access in ir.Accesses -> resolution.resolveAccess` |
| B2-P2A-A001-D003 | resolve_type_annotations | 61.585096500 s / 33.382883% / 2 | {"files":887,"typeAnnotations":35003} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:resolveTypeAnnotation:98-100`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each annotation in ir.TypeAnnotations -> resolution.resolveTypeAnnotation` |
| B2-P2A-A001-D004 | project_resolution_outcomes | 5.375049800 s / 2.913605% / 4 | {"graphNodes":36575,"graphRelationships":65045,"outcomes":86742,"referencesBySourceScope":25062,"referencesByTargetDef":25062,"workUnits":238486} | **unavailable / not measured** | `E:\Anvien\internal\resolution\outcome.go:projectResolutionOutcomes:114-116`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.projectResolutionOutcomes` |
| B2-P2A-A001-D005 | emit_definition_nodes | 0.452126400 s / 0.245080% / 5 | {"definitions":25811,"exports":3539,"files":887,"runs":1,"workUnits":30237} | **unavailable / not measured** | `E:\Anvien\internal\resolution\emit.go:emitDefinitionNodes:77-81`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitDefinitionNodes` |
| B2-P2A-A001-D006 | finalize_resolution_outcomes | 0.224157600 s / 0.121507% / 7 | {"outcomeMapEntries":86742} | **unavailable / not measured** | `E:\Anvien\internal\resolution\outcome.go:(*resolutionOutcomeCollector).finalize:110-113`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> (*resolutionOutcomeCollector).finalize` |
| B2-P2A-A001-D007 | build_binding_occurrence_index | 0.075494100 s / 0.040922% / 8 | {"bindingLeaves":6922,"definitionsVisited":25811,"filePasses":1774,"files":887,"ownedDefIDsVisited":25811,"ownerBindingsInspected":30749,"runs":1,"scopesVisited":6502,"workUnits":97569} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:buildBindingOccurrenceIndex:62-65`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.buildBindingOccurrenceIndex` |
| B2-P2A-A001-D008 | finalize_typescript_authority_results | 0.040710200 s / 0.022067% / 9 | {"authorityResults":23110} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:finalizeTypeScriptAuthorityResults:103-106`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.finalizeTypeScriptAuthorityResults` |
| B2-P2A-A001-D009 | emit_import_edges | 0.348401700 s / 0.188855% / 6 | {"imports":6141,"targetDefinitions":4717,"targetFiles":4985,"workUnits":15843} | **unavailable / not measured** | `E:\Anvien\internal\resolution\emit.go:emitImportEdges:83`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitImportEdges` |
| B2-P2A-A001-D010 | emit_typescript_external_symbols | 0.028817500 s / 0.015621% / 10 | {"authorityResults":21018,"resolvedRecords":0,"uniqueResolvedSymbols":0,"workUnits":21018} | **unavailable / not measured** | `E:\Anvien\internal\resolution\external_symbol.go:emitTypeScriptExternalSymbols:107-109`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitTypeScriptExternalSymbols` |
| B2-P2A-A001-D011 | emit_method_dispatch_edges | 0.004458600 s / 0.002417% / 11 | {"heritageFacts":0,"memberEntries":7485,"ownerMemberOwners":1548} | **unavailable / not measured** | `E:\Anvien\internal\resolution\emit.go:emitMethodDispatchEdges:102`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitMethodDispatchEdges` |
| B2-P2A-A001-D012 | assemble_resolution_result | 0.000000000 s / 0.000000% / 13 | {"authorityResults":21018,"graphPointers":1,"metricsValues":1,"outcomes":86742,"referenceIndexes":1,"resultAssemblies":1} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:121-127`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> construct resolution.Result return value` |
| B2-P2A-A001-D013 | binding_accumulator_dispose | 0.000000000 s / 0.000000% / 14 | {"accumulatedEntries":7330,"deferredExecutions":1,"fileEntryBuckets":508,"fileScopeBuckets":508,"workUnits":8346} | **unavailable / not measured** | `E:\Anvien\internal\resolution\binding_accumulator.go:(*bindingAccumulator).dispose:71-74`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto deferred closure -> (*bindingAccumulator).dispose` |
| B2-P2A-A001-D014 | emit_heritage_compatibility_edges | 0.000000000 s / 0.000000% / 15 | {"heritageFacts":0,"runs":1} | **unavailable / not measured** | `E:\Anvien\internal\resolution\emit.go:emitHeritageCompatibilityEdges:85-89`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each w.heritage -> resolution.emitHeritageCompatibilityEdges` |
| B2-P2A-A001-D015 | emit_unresolved_heritage_diagnostics | 0.000999500 s / 0.000542% / 12 | {"runs":1,"unresolvedHeritageFacts":7} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:emitUnresolvedHeritageDiagnostics:82`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitUnresolvedHeritageDiagnostics` |
| B2-P2A-A001-D016 | finalize_resolution_metadata | 0.000000000 s / 0.000000% / 16 | {"graphNodes":36575,"graphRelationships":65045,"metadataUpdates":1,"workUnits":101621} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:118-120`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> graphhealth.SetResolutionMetadata` |
| B2-P2A-A001-D017 | runtime_setup | 0.000000000 s / 0.000000% / 17 | {"graphAllocations":0,"newEmitterInvocations":1,"resolveBoundIntoInvocations":1} | **unavailable / not measured** | `E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:66-75`; `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> conditional graph.New -> resolution.newEmitter` |

| Conservation field | Accepted A001 before | Raw A002 candidate |
|---|---:|---:|
| 17-child sum | `184.422170400 s` | **unavailable / not measured** |
| Resolution parent | `184.481061700 s` | `22.664132600 s` |
| Parent residual | `0.058891300 s` | **unavailable / not measured** |
| Exclusive intervals | `2675` | **unavailable / not measured** |
| Overlap count | `0` | **unavailable / not measured** |

No candidate child conservation arithmetic exists. In particular, `22.664132600 s - unknown child sum` is not a measured residual.

## Workload, semantic, graph, output, and resource evidence

The following same-work counts match the accepted A001 optimized capture exactly:

| Field | A001 before | A002 candidate | Result |
|---|---:|---:|---|
| Files scanned / parsed / failed | 1,359 / 887 / 0 | 1,359 / 887 / 0 | equal |
| Parser bytes | 5,430,581 | 5,430,581 | equal |
| Graph nodes / relationships | 95,762 / 129,716 | 95,762 / 129,716 | equal |
| Dependency edges / projection unresolved | 13,360 / 735 | 13,360 / 735 | equal |
| DB readback node / relationship rows | 95,762 / 129,716 | 95,762 / 129,716 | equal |
| Top-level operation coverage | 30/30 | 30/30 | equal |
| Resolution-child coverage | 17/17 | **0/17** | **inside-boundary mismatch; packet incomplete** |
| Go startAllocBytes | 1,309,264 | 1,299,272 | -9,992 |
| Go endAllocBytes | 863,629,352 | 666,783,664 | -196,845,688 |
| Go maxObservedSys | 1,989,634,552 | 1,972,447,736 | -17,186,816 |

Candidate resolution semantic counters, which match A001 exactly, are:

```json
{"DefinitionsIndexed":25811,"ImportsResolved":4985,"ImportUsesEmitted":4717,"ResolvedReferences":25062,"UnresolvedReferences":36845,"UnresolvedReferenceDiagnostics":57683,"UnattributedUnresolvedReferences":0,"ResolvedCalls":6461,"ResolvedAccesses":7615,"ResolvedTypeReferences":10986,"ResolvedExternalDeclarations":0,"ExternalCapabilityUnavailable":21018,"ExternalProfileExcluded":0,"ExternalMeaningMismatches":0,"HeritageFactsIndexed":7,"ResolvedInheritance":0,"UnresolvedInheritance":7,"DuplicateEdgesMerged":14039,"MethodOverridesEmitted":0,"MethodImplementsEmitted":0,"FinalizedImportsEmitted":4985,"CrossFileFilesReprocessed":0,"CrossFileSkipped":true,"CrossFileSkipReason":"covered-by-scopeir-single-pass-resolution","BindingAccumulatorFiles":508,"BindingAccumulatorEntries":7330,"BindingAccumulatorFinalized":true,"BindingAccumulatorDisposed":true,"GraphNodesEmitted":36575,"GraphRelationshipsEmitted":65045}
```

Semantic-counter equality does not supply the absent child timers or denominators.

| Artifact / invariant | Accepted A001 before | Raw A002 candidate | Result / boundary classification |
|---|---|---|---|
| Canonical `graph.json` | 840,614,023 bytes; `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920` | 840,614,023 bytes; `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920` | exact hash match |
| Public `stdout.json` | `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` | `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` | exact hash match |
| `stderr.log` | `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` | `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` | exact hash match |
| Ladybug `lbug` | 597,950,464 bytes; `DC88B68B26BE75954A8BEE85F4CD426F0074171E03AAAC6D55B49C18EDDE1AB7` | 597,946,368 bytes; `8E45E92F142E0BA02D13926F15ABC2FD5C77530F9E1DE185A513A44EDF7135F5` | raw-byte mismatch; outside semantic equivalence boundary by the governing Architect evidence; no acceptance inference |
| `meta.json` | 219 bytes; `5D5227EB91053443D51A51E9EC15D85532C1EF49E7F701B45B24F81076D9F1D9` | 219 bytes; `ACE0CECA1A0320CE42D50E807AE92B2833AB16247A4D28D228536F4DBC4DB991` | raw-byte mismatch; outside semantic equivalence boundary; no cause or acceptance inference |
| Target status | ten preserved rows | same ten rows | equal |
| Child instrumentation | present | absent | **inside packet-completeness boundary; blocking mismatch** |

Because child instrumentation is absent, the full equivalence/resource packet required by A002 is incomplete even though the canonical graph, public output, workload, graph counts, DB counts, semantic counters, and top-level denominators match.

## Structural retained-state confirmation

Read-only source/Coder evidence shows the A002 diagnostic appender retains `O(T)` map entries and slice headers:

- `diagnosticAppender` owns `map[string][]Diagnostic`.
- In `E:\Anvien\internal\graphhealth\diagnostics.go:84-85`, the same `diagnostics` slice is assigned to the map and to the node property; a second retained diagnostic-object copy is not maintained.
- The appender is created through `newEmitter` and becomes unreachable after the resolution run.
- The two retained-state routes are `E:\Anvien\internal\resolution\emit.go:220` and `E:\Anvien\internal\resolution\outcome.go:392`.

This is structural source evidence only, not a substitute for the missing 17-child runtime measurements.

## Raw artifact authority

| Artifact | Exact path | Bytes / SHA-256 where recorded |
|---|---|---|
| Raw root | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark` | directory |
| Preflight identity | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\preflight.json` | authoritative identity/status record |
| One-shot wrapper | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\run-candidate.ps1` | exact one-shot launcher |
| Frozen candidate executable | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\anvien-a002-cheapapp-candidate.exe` | 73,767,936 / `11895442CA4A5708FA439A787BC541102BE7144856315BCD17E9A7B772658AED` |
| Benchmark JSON | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\benchmark.json` | 15,845 / `4CDF0F9813D977EE9C5B303EB94E1376800C4FAED98B7FC61CDB82EBC07747B3` |
| Process JSON | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\process.json` | 1,140 / `E8D026F49B30A45503C0718EC0A611EB6021D77AC34538056B144DA238A2704E` |
| Public stdout | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\stdout.json` | 8,824 / `F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4` |
| Stderr | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\stderr.log` | 213 / `0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC` |
| Preserved canonical graph | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\target-output\graph.json` | 840,614,023 / `8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920` |
| Preserved Ladybug DB | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\target-output\lbug` | 597,946,368 / `8E45E92F142E0BA02D13926F15ABC2FD5C77530F9E1DE185A513A44EDF7135F5` |
| Preserved metadata | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\target-output\meta.json` | 219 / `ACE0CECA1A0320CE42D50E807AE92B2833AB16247A4D28D228536F4DBC4DB991` |
| Candidate resolution CPU profile | expected `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\resolution.cpu.pprof` | **not produced / unavailable** |
| Accepted A001 authority | `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\owner-final-run\optimized` | immutable before packet |
| Accepted A001 report | `E:\Anvien\reports\Investigation\rp_child06a_a001_cheapapp_benchmark.md` | before-table authority |

## Exact blocker and handoff

`E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark\candidate\benchmark.json` contains exactly 30 top-level `operations`, but has no root `resolutionChildren` or equivalent child-measurement property. Its `resolution` member contains only semantic counters. Therefore the following mandatory candidate fields are absent:

- D001 candidate duration, calls, files, rank, parent share, and delta;
- D002-D017 candidate duration, denominator, rank, share, and delta;
- candidate 17-child sum, parent residual, exclusive-interval count, and overlap count;
- candidate `resolution.cpu.pprof`.

No retry was launched. No baseline was rerun. No production source, test, plan, ledger, target source, or Restaurant Manager artifact was edited or accessed. The only valid completion state is `MEASUREMENT_A002_CHEAPAPP_INCOMPLETE_MISSING_17_CHILD_INSTRUMENTATION`.
